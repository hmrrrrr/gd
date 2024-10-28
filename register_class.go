//go:build !generate

package gd

import (
	"fmt"
	"reflect"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"grow.graphics/gd/gdclass"
	"grow.graphics/gd/gdclass/EditorPlugin"
	"grow.graphics/gd/gdclass/Engine"
	"grow.graphics/gd/gdclass/MainLoop"
	"grow.graphics/gd/gdclass/Node"
	"grow.graphics/gd/gdclass/ProjectSettings"
	"grow.graphics/gd/gdclass/Script"
	"grow.graphics/gd/gdclass/ScriptLanguage"
	gd "grow.graphics/gd/internal"
	classdb "grow.graphics/gd/internal/classdb"
	"grow.graphics/gd/internal/discreet"
	"grow.graphics/gd/internal/mmm"
)

type onFree mmm.Pointer[func(), onFree, [0]uintptr]

func (cb onFree) Free() {
	(*mmm.API(cb))()
	mmm.End(cb)
}

// Tool can be embedded inside a struct to make it run in the editor.
type Tool interface{ tool() }

/*
RegisterClass registers a struct available for use inside Godot
extending the given 'Parent' Godot class. The 'Struct' type must
be a named struct that embeds a [Class] field specifying the
parent class to extend.

	type MyClass struct {
		Class[MyClass, Node2D] `gd:"MyClass"`
	}

The tag can be adjusted in order to change the name of the class
within Godot.

Use this in a main or init function to register your Go structs
and they will become available within the Godot engine for use
in the editor and/or within scripts.

All exported fields and methods will be exposed to Godot, so
take caution when embedding types, as their fields and methods
will be promoted.

If the Struct extends [EditorPlugin] then it will be added to
the editor as a plugin.

If the Struct extends [MainLoop] or [SceneTree] then it will be
used as the main loop for the application.

If the Struct implements an OnRegister(Lifetime) method, it will
be called on a temporary instance when the class is registered.
*/
func Register[Struct gd.Extends[Parent], Parent gd.IsClass]() {
	var classType = reflect.TypeOf([0]Struct{}).Elem()
	var superType = reflect.TypeOf([0]Parent{}).Elem()
	if classType.Kind() != reflect.Struct || classType.Name() == "" {
		panic("gdextension.RegisterClass: Class type must be a named struct")
	}

	// Support 'gd' tag for renaming the class within Godot.
	var rename = classNameOf(classType)

	var tool = false
	switch {
	case superType.Implements(reflect.TypeOf([0]interface{ AsScript() Script.Go }{}).Elem()),
		superType.Implements(reflect.TypeOf([0]interface{ AsEditorPlugin() EditorPlugin.Go }{}).Elem()),
		superType.Implements(reflect.TypeOf([0]interface{ AsScriptLanguage() ScriptLanguage.Go }{}).Elem()),
		classType.Implements(reflect.TypeOf([0]Tool{}).Elem()):
		tool = true
	}

	var className = gd.NewStringName(rename)
	var superName = gd.NewStringName(classNameOf(superType))

	gd.Global.ClassDB.RegisterClass(gd.Global.ExtensionToken, className, superName,
		&classImplementation{
			Name:  className,
			Super: superName,
			Type:  classType,
			Tool:  tool,
			Value: reflect.New(classType).Interface().(gd.ExtensionClass),
		})
	_ = func() {
		gd.Global.ClassDB.UnregisterClass(gd.Global.ExtensionToken, className)
	}
	//mmm.Pin[onFree](godot.Lifetime, &unregister, [0]uintptr{})

	registerSignals(className, classType)
	registerMethods(className, classType)

	if superType.Implements(reflect.TypeOf([0]interface {
		AsMainLoop() MainLoop.Go
	}{}).Elem()) {
		main_loop_type := gd.NewString("application/run/main_loop_type")
		ProjectSettings.GD().SetInitialValue(main_loop_type, gd.NewVariant(className))
		ProjectSettings.GD().SetSetting(main_loop_type, gd.NewVariant(className))
	}

	type onRegister interface {
		OnRegister()
	}
	if reflect.PointerTo(classType).Implements(reflect.TypeOf([0]onRegister{}).Elem()) {
		impl := reflect.New(classType).Interface().(onRegister)
		impl.OnRegister()
	}
}

func convertName(fnName string) string {
	if fnName == "seek" {
		return "SeekTo"
	}
	if fnName == "type_string" {
		return "TypeToString"
	}
	fnName = strings.ToLower(fnName)
	joins := []string{}
	split := strings.Split(fnName, "_")
	for _, word := range split {
		joins = append(joins, cases.Title(language.English).String(word))
	}
	return strings.Join(joins, "")
}

type classImplementation struct {
	Name  StringName
	Super StringName

	Tool bool

	Type reflect.Type

	Value gd.ExtensionClass
}

var _ gd.ClassInterface = classImplementation{}

func (class classImplementation) IsVirtual() bool {
	return false
}

func (class classImplementation) IsAbstract() bool {
	return class.Type.Kind() == reflect.Interface
}

func (class classImplementation) IsExposed() bool {
	return true // TODO return false if the Go type is not exported.
}

func (class classImplementation) CreateInstance() Object {
	var super = gd.Global.ClassDB.ConstructObject(class.Super)
	instance := class.reloadInstance(super)
	gd.Global.Object.SetInstance(super, class.Name, instance)
	gd.Global.Object.SetInstanceBinding(super, gd.Global.ExtensionToken, nil, nil)
	instance.OnCreate()
	return super
}

func (class classImplementation) ReloadInstance(super Object) gd.ObjectInterface {
	return class.reloadInstance(super)
}

func (class classImplementation) reloadInstance(super Object) gd.ObjectInterface {
	var value = reflect.New(class.Type)
	extensionClass := value.Interface().(gd.ExtensionClass)
	extensionClass.SetPointer(super)

	gd.Global.Instances[discreet.Get(super)[0]] = extensionClass

	value = value.Elem()

	// TODO cache this check
	var signals []signalChan
	for i := 0; i < value.NumField(); i++ {
		var field = value.Type().Field(i)
		if !field.IsExported() || field.Name == "Class" {
			continue
		}
		var (
			rvalue = value.Field(i).Addr()
		)
		name := field.Name
		if tag := field.Tag.Get("gd"); tag != "" {
			name = tag
		}
		// Signal fields need to have their values injected into the field, so that they can be used (emitted).
		if setter, ok := rvalue.Interface().(isSignal); ok {
			signal := gd.NewSignalOf(super, gd.NewStringName(name))
			setter.setSignal(signal)
			emit := rvalue.Elem().FieldByName("Emit")
			fnType := emit.Type()
			emit.Set(reflect.MakeFunc(fnType, func(args []reflect.Value) (results []reflect.Value) {
				var variants = make([]gd.Variant, 0, len(args))
				for _, arg := range args {
					variants = append(variants, gd.NewVariant(arg.Interface()))
				}
				signal.Emit(variants...)
				return nil
			}))
		}
		if field.Type.Kind() == reflect.Chan && field.Type.ChanDir() == reflect.SendDir {
			signal := gd.NewSignalOf(super, gd.NewStringName(name))
			ch := reflect.MakeChan(reflect.ChanOf(reflect.BothDir, field.Type.Elem()), 0)
			rvalue.Elem().Set(ch)
			signals = append(signals, signalChan{
				signal: signal,
				rvalue: ch,
			})
		}
	}
	if len(signals) > 0 {
		go manageSignals(super.AsObject().GetInstanceId(), signals)
	}
	return &instanceImplementation{
		object:   discreet.Get(super)[0],
		Value:    value.Addr().Interface().(gd.ExtensionClass),
		signals:  signals,
		isEditor: !class.Tool && Engine.IsEditorHint(),
	}
}

func (class classImplementation) GetVirtual(name StringName) any {
	if !class.Tool && Engine.IsEditorHint() {
		return nil
	}
	var virtual = gd.VirtualByName(class.Value, name.String())
	if !virtual.IsValid() {
		return nil
	}
	var vtype = virtual.Type().In(0)
	GoName := convertName(name.String())
	if GoName == "Ready" {
		return nil // special case, as we override this method for all node types, so that we can assert the scene tree.
	}
	method, ok := reflect.PointerTo(class.Type).MethodByName(GoName)
	if !ok {
		return nil
	}
	if method.Type.NumIn() != vtype.NumIn() {
		panic(fmt.Sprintf("gdextension.RegisterClass: Method %s.%s does not match %s.%s\nis %s want %s", class.Type.Name(), GoName, virtual.Type().Name(), name, method.Type, vtype))
	}
	for i := 1; i < method.Type.NumIn(); i++ {
		if method.Type.In(i) != vtype.In(i) {
			panic(fmt.Sprintf("gdextension.RegisterClass: Method %s.%s does not match %s.%s\nis %s want %s", class.Type.Name(), GoName, virtual.Type().Name(), name, method.Type, vtype))
		}
	}
	var copy = reflect.New(method.Type)
	copy.Elem().Set(method.Func)
	var fn = reflect.NewAt(vtype, copy.UnsafePointer()).Elem()
	return virtual.Call([]reflect.Value{fn})[0].Interface()
}

type instanceImplementation struct {
	object   uintptr
	Value    gd.ExtensionClass
	signals  []signalChan
	isEditor bool
}

var lastGC int

func (instance *instanceImplementation) setupForCall() func() {
	if frame := Engine.GetFramesDrawn(); frame > lastGC {
		discreet.MarkAndSweep()
	}
	return func() {}
}

func (instance *instanceImplementation) OnCreate() {
	if impl, ok := instance.Value.(interface {
		OnCreate()
	}); ok {
		defer instance.setupForCall()()
		impl.OnCreate()
	}
}

func (instance *instanceImplementation) Set(name StringName, value gd.Variant) bool {
	if impl, ok := instance.Value.(interface {
		Set(gd.StringName, gd.Variant) gd.Bool
	}); ok {
		defer instance.setupForCall()()
		ok := bool(impl.Set(name, value))
		if ok {
			if impl, ok := instance.Value.(interface {
				OnSet(gd.StringName, gd.Variant)
			}); ok {
				impl.OnSet(name, value)
			}
		}
		return ok
	}
	sname := name.String()
	rvalue := reflect.ValueOf(instance.Value).Elem()
	field := rvalue.FieldByName(sname)
	if !field.IsValid() {
		for i := 0; i < rvalue.NumField(); i++ {
			if rvalue.Type().Field(i).Tag.Get("gd") == sname {
				field = rvalue.Field(i)
				break
			}
		}
		if !field.IsValid() {
			return false
		}
	}
	if !field.CanSet() {
		return false
	}
	val := value.Interface()
	converted := reflect.ValueOf(val)
	if !converted.IsValid() {
		return false
	}
	if converted.Type().ConvertibleTo(field.Type()) {
		converted = converted.Convert(field.Type())
	}
	if !converted.Type().AssignableTo(field.Type()) {
		return false
	}
	if obj, ok := val.(gd.IsClass); ok {
		ref, ok := gd.As[gd.RefCounted](obj.AsObject())
		if ok {
			ref.Reference()
		}
		//	field.Addr().Interface().(gd.PointerToClass).SetPointer(mmm.Pin[gd.Pointer](instance.Value.GetKeepAlive().Lifetime, instance.Value.GetKeepAlive().API, mmm.End(obj.AsObject().AsPointer())))
	} else {
		field.Set(converted)
	}
	if impl, ok := instance.Value.(interface {
		OnSet(gd.StringName, gd.Variant)
	}); ok {
		defer instance.setupForCall()()
		impl.OnSet(name, value)
	}
	return true
}

func (instance *instanceImplementation) Get(name StringName) (gd.Variant, bool) {
	if impl, ok := instance.Value.(interface {
		Get(StringName) gd.Variant
	}); ok {
		defer instance.setupForCall()()
		return impl.Get(name), true
	}
	sname := name.String()
	rvalue := reflect.ValueOf(instance.Value).Elem()
	field := rvalue.FieldByName(sname)
	if !field.IsValid() {
		for i := 0; i < rvalue.NumField(); i++ {
			rfield := rvalue.Type().Field(i)
			if !rfield.Anonymous && rfield.Tag.Get("gd") == sname {
				field = rvalue.Field(i)
				break
			}
		}
		if !field.IsValid() {
			return gd.Variant{}, false
		}
	}
	if reflect.PointerTo(field.Type()).Implements(reflect.TypeOf([0]isSignal{}).Elem()) {
		return gd.Variant{}, false
	}
	return gd.NewVariant(field.Interface()), true
}

func (instance *instanceImplementation) GetPropertyList() []gd.PropertyInfo {
	if impl, ok := instance.Value.(interface {
		GetPropertyList() []gd.PropertyInfo
	}); ok {
		defer instance.setupForCall()()
		return impl.GetPropertyList()
	}
	rtype := reflect.ValueOf(instance.Value).Elem().Type()
	var list = make([]gd.PropertyInfo, 0, rtype.NumField())
	for i := 0; i < rtype.NumField(); i++ {
		field := rtype.Field(i)
		if !field.IsExported() || field.Anonymous {
			continue
		}
		if _, ok := field.Type.MethodByName("AsNode"); ok {
			continue
		}
		list = append(list, propertyOf(field))
	}
	return list
}

func (instance *instanceImplementation) PropertyCanRevert(name StringName) bool {
	if impl, ok := instance.Value.(interface {
		PropertyCanRevert(gd.StringName) gd.Bool
	}); ok {
		defer instance.setupForCall()()
		return bool(impl.PropertyCanRevert(name))
	}
	sname := name.String()
	rtype := reflect.TypeOf(instance.Value).Elem()
	field, ok := rtype.FieldByName(sname)
	if !ok {
		for i := 0; i < rtype.NumField(); i++ {
			if rtype.Field(i).Tag.Get("gd") == sname {
				field = rtype.Field(i)
				ok = true
				break
			}
		}
	}
	if !ok {
		return false
	}
	_, ok = field.Tag.Lookup("default")
	return ok
}
func (instance *instanceImplementation) PropertyGetRevert(name StringName) (gd.Variant, bool) {
	if impl, ok := instance.Value.(interface {
		PropertyGetRevert(gd.StringName) (gd.Variant, bool)
	}); ok {
		defer instance.setupForCall()()
		return impl.PropertyGetRevert(name)
	}
	sname := name.String()
	rtype := reflect.TypeOf(instance.Value).Elem()
	field, ok := rtype.FieldByName(sname)
	if !ok {
		for i := 0; i < rtype.NumField(); i++ {
			if rtype.Field(i).Tag.Get("gd") == sname {
				field = rtype.Field(i)
				ok = true
				break
			}
		}
	}
	if !ok {
		return gd.Variant{}, false
	}
	var value = reflect.New(field.Type)
	if tag := field.Tag.Get("default"); tag != "" {
		_, err := fmt.Sscanf(tag, "%v", value.Interface())
		if err != nil {
			return gd.Variant{}, false
		}
	}
	return gd.NewVariant(value.Elem().Interface()), true
}

func (instance *instanceImplementation) ValidateProperty(name StringName, info *gd.PropertyInfo) bool {
	defer instance.setupForCall()()
	switch validate := instance.Value.(type) {
	case interface {
		ValidateProperty(gd.StringName, *gd.PropertyInfo) gd.Bool
	}:
		return bool(validate.ValidateProperty(name, info))
	}
	return true
}

func (instance *instanceImplementation) Notification(what int32, reversed bool) {
	if what == 13 { // NOTIFICATION_READY
		instance.ready()
	}
	if !instance.isEditor {
		defer instance.setupForCall()()
		switch notify := instance.Value.(type) {
		case interface{ Notification(NotificationType) }:
			notify.Notification(NotificationType(what))
		default:
		}
	}
}

func (instance *instanceImplementation) ToString() (String, bool) {
	defer instance.setupForCall()()
	switch onfree := instance.Value.(type) {
	case interface{ ToString() gd.String }:
		return onfree.ToString(), true
	}
	return String{}, false
}

func (instance *instanceImplementation) Reference()   {}
func (instance *instanceImplementation) Unreference() {}

func (instance *instanceImplementation) CallVirtual(name StringName, virtual any, args gd.UnsafeArgs, back gd.UnsafeBack) {
	defer instance.setupForCall()()
	virtual.(gd.ExtensionClassCallVirtualFunc)(instance.Value, args, back)
}

func (instance *instanceImplementation) GetRID() RID {
	return 0
}

func (instance *instanceImplementation) Free() {
	for _, signal := range instance.signals {
		signal.rvalue.Close()
	}
	delete(gd.Global.Instances, instance.object)
	defer instance.setupForCall()()
	switch onfree := instance.Value.(type) {
	case interface{ OnFree() }:
		onfree.OnFree()
	}
}

// ready is responsible for asserting the scene tree for struct members that implement
// Super().AsNode() and asserting that these nodes are added as children to the Super.
//
// TODO this could be partially pre-compiled for a given [Register] type and cached in
// order to avoid any use of reflection at instantiation time.
func (instance *instanceImplementation) ready() {
	var parent gdclass.Node = gdclass.Node{classdb.Node(instance.Value.AsObject())}

	var rvalue = reflect.ValueOf(instance.Value).Elem()
	for i := 0; i < rvalue.NumField(); i++ {
		field := rvalue.Type().Field(i)
		if !field.IsExported() || field.Name == "Class" {
			continue
		}
		instance.assertChild(rvalue.Field(i).Addr().Interface(), field, parent, parent)
	}
	if !instance.isEditor {
		defer instance.setupForCall()()
		switch ready := instance.Value.(type) {
		case interface{ Ready() }:
			ready.Ready()
		}
	}
}

func (instance *instanceImplementation) assertChild(value any, field reflect.StructField, parent, owner gdclass.Node) {
	type isNode interface {
		gd.PointerToClass

		AsNode() Node.Go
	}
	nodeType := reflect.TypeOf([0]isNode{}).Elem()
	if !field.Type.Implements(nodeType) && !reflect.PointerTo(field.Type).Implements(nodeType) {
		return
	}
	var (
		rvalue = reflect.ValueOf(value)
	)
	if rvalue.Elem().Kind() == reflect.Pointer {
		rvalue.Elem().Set(reflect.New(rvalue.Elem().Type().Elem()))
		value = rvalue.Elem().Interface()
	}
	class := value.(isNode)
	if rvalue.Elem().Kind() == reflect.Struct {
		defer func() {
			rvalue := rvalue.Elem()
			for i := 0; i < rvalue.NumField(); i++ {
				field := rvalue.Type().Field(i)
				if !field.IsExported() || field.Name == "Class" || field.Anonymous {
					continue
				}
				instance.assertChild(rvalue.Field(i).Addr().Interface(), field, class.AsNode(), owner)
			}
		}()
	}
	name := field.Name
	if tag := field.Tag.Get("gd"); tag != "" {
		name = tag
	}
	path := gd.NewString(name).NodePath()
	if !Node.GD(parent).HasNode(path) {
		child := gd.Global.ClassDB.ConstructObject(gd.NewStringName(classNameOf(field.Type)))
		native, ok := gd.Global.Instances[discreet.Get(child)[0]]
		if ok {
			rvalue.Elem().Set(reflect.ValueOf(native))
			class = native.(isNode)
		}
		var mode Node.InternalMode = Node.InternalModeDisabled
		if !field.IsExported() {
			mode = Node.InternalModeFront
		}
		Node.GD(class.AsNode()).SetName(gd.NewString(field.Name))
		var adding gdclass.Node = gdclass.Node{classdb.Node(class.AsObject())}
		Node.GD(parent).AddChild(adding, true, mode)
		Node.GD(class.AsNode()).SetOwner(owner)
		return
	}
	var node = Node.GD(parent).GetNode(path)
	if name := node[0].AsObject().GetClass().String(); name != classNameOf(field.Type) {
		panic(fmt.Sprintf("gd.Register: Node %s.%s is not of type %s (%s)", rvalue.Type().Name(), field.Name, field.Type.Name(), name))
	}
	ref, native := gd.Global.Instances[discreet.Get(node[0])[0]]
	if native {
		rvalue.Elem().Set(reflect.ValueOf(ref))
		discreet.End(node[0])
	} else {
		//class.SetPointer(mmm.Let[gd.Pointer](instance.Value.GetKeepAlive().Lifetime, tmp.API, mmm.End(node[0].AsPointer())))
	}
}
