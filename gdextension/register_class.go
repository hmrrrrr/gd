package gdextension

import (
	"context"
	"fmt"
	"reflect"
	"runtime/cgo"
	"strings"
	"unsafe"

	gd "grow.graphics/gd/internal"
	"runtime.link/mmm"
)

/*
RegisterClass registers a struct available for use inside Godot
extending the given 'Parent' Godot class. The 'Struct' type must
be a named struct that embeds a [gd.Class] field specifying the
parent class to extend.

	type MyClass struct {
		gd.Class[MyClass, gd.Node2D]
	}

Use this in a main or init function to register your Go structs
and they will become available within the Godot engine for use
in the editor and/or within scripts.

All exported fields and methods will be exposed to Godot, so
take caution when embedding types, as their fields and methods
will be promoted.

If the Struct extends [gd.EditorPlugin] then it will be added to
the editor as a plugin.
*/
func Register[Struct gd.Extends[Parent], Parent gd.IsClass](godot gd.Context) {
	var classType = reflect.TypeOf([0]Struct{}).Elem()
	var superType = reflect.TypeOf([0]Parent{}).Elem()
	if classType.Kind() != reflect.Struct || classType.Name() == "" {
		panic("gdextension.RegisterClass: Class type must be a named struct")
	}

	var className = godot.StringName(classType.Name())
	var superName = godot.StringName(strings.TrimPrefix(superType.Name(), "class"))

	var info gd.ClassCreationInfo

	attachProperties(godot, className, &info, classType)

	info.IsExposed = true

	info.CreateInstance.Set(func(userdata unsafe.Pointer) uintptr {
		ctx := gd.NewContext(godot.API())
		var super = godot.API().ClassDB.ConstructObject(ctx, superName)
		var instance = reflect.New(classType)
		instance.Interface().(gd.PointerToClass).SetPointer(
			mmm.Make[gd.API, gd.Pointer](ctx, godot.API(), super.Pointer()))
		mmm.MarkFree(super.AsPointer())
		injectDependenciesInto(ctx, godot.API(), instance.Elem(), superType)
		godot.API().Object.SetInstance(super, className, instance.Interface())
		return super.Pointer()
	})
	info.FreeInstance.Set(func(userdata unsafe.Pointer, handle cgo.Handle) {
		class := handle.Value().(gd.IsClass)
		gd.MarkFree(class) // godot frees the underlying object, so we need to mark it as free.
		class.Context().(interface{ Free() }).Free()
		handle.Delete()
	})

	// Dispatch virtual functions, these are functions that structs can
	// override with their own implementation.
	info.GetVirtual.Set(func(userdata unsafe.Pointer, ptr_name gd.StringNamePtr) gd.ExtensionClassCallVirtualFunc {
		sname := mmm.Make[gd.API, gd.StringName](nil, godot.API(), *ptr_name)
		vname := sname.String()
		var class Struct
		var virtual = gd.VirtualByName(class, vname)
		if !virtual.IsValid() {
			return gd.ExtensionClassCallVirtualFunc{}
		}
		var vtype = virtual.Type().In(0)
		GoName := convertName(vname)
		method, ok := reflect.PtrTo(classType).MethodByName(GoName)
		if !ok {
			return gd.ExtensionClassCallVirtualFunc{}
		}
		if method.Type.NumIn() != vtype.NumIn() {
			panic(fmt.Sprintf("gdextension.RegisterClass: Method %s.%s does not match %s.%s", classType.Name(), GoName, virtual.Type().Name(), vname))
		}
		for i := 1; i < method.Type.NumIn(); i++ {
			if method.Type.In(i) != vtype.In(i) {
				panic(fmt.Sprintf("gdextension.RegisterClass: Method %s.%s does not match %s.%s", classType.Name(), GoName, virtual.Type().Name(), vname))
			}
		}
		var copy = reflect.New(method.Type)
		copy.Elem().Set(method.Func)
		var fn = reflect.NewAt(vtype, copy.UnsafePointer()).Elem()
		return virtual.Call([]reflect.Value{fn, reflect.ValueOf(godot.API())})[0].Interface().(gd.ExtensionClassCallVirtualFunc)
	})

	godot.API().ClassDB.RegisterClass(godot.API().ExtensionToken, (gd.StringNamePtr)(unsafe.Pointer(&className)), (gd.StringNamePtr)(unsafe.Pointer(&superName)), &info)
	godot.Defer(func() {
		godot.API().ClassDB.UnregisterClass(godot.API().ExtensionToken, (gd.StringNamePtr)(unsafe.Pointer(&className)))
	})

	/*if superType.Implements(reflect.TypeOf([0]interface{ AsEditorPlugin() gd.EditorPlugin }{}).Elem()) {
		godot.API().EditorPlugins.Add((gd.StringNamePtr)(unsafe.Pointer(&className)))
		godot.Defer(func() {
			godot.API().EditorPlugins.Remove((gd.StringNamePtr)(unsafe.Pointer(&className)))
		})
	}*/

	registerSignals(godot, className, classType)
	registerMethods(godot, className, classType)
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
		joins = append(joins, strings.Title(word))
	}
	return strings.Join(joins, "")
}

// injectDependenciesInto the given freshly allocated value, this
// function makes sure any [gd.Object] types are instantiated and
// that any referenced singletons are injected.
func injectDependenciesInto(ctx context.Context, Godot *gd.API, value reflect.Value, superType reflect.Type) {
	if value.CanAddr() && value.Kind() != reflect.Struct {
		panic("gdextension.injectDependenciesInto: value must be an addressable struct")
	}
	localCtx := gd.NewContext(Godot)
	defer localCtx.Free()

	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)

		// support private fields.
		fieldValue := reflect.NewAt(field.Type, unsafe.Add(value.Addr().UnsafePointer(), field.Offset)).Interface()

		container, ok := fieldValue.(gd.PointerToClass)
		if ok && container.Pointer() == 0 {
			_, ok := fieldValue.(gd.Singleton)
			if ok {
				var name = localCtx.StringName(strings.TrimPrefix(field.Type.Name(), "class"))
				singleton := Godot.Object.GetSingleton(localCtx, name)
				container.SetPointer(mmm.Make[gd.API, gd.Pointer](nil, Godot, singleton.Pointer()))
			}
		}

	}
}
