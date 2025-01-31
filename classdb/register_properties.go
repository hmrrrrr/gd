package classdb

import (
	"fmt"
	"reflect"

	ResourceClass "graphics.gd/classdb/Resource"
	gd "graphics.gd/internal"
	"graphics.gd/internal/pointers"
)

func propertyOf(field reflect.StructField) (gd.PropertyInfo, bool) {
	var name = field.Name
	tag, ok := field.Tag.Lookup("gd")
	if ok {
		name = tag
	}
	var vtype gd.VariantType
	var hint PropertyHint
	var hintString = nameOf(field.Type)
	switch {
	case field.Type.Kind() == reflect.Pointer && field.Type.Implements(reflect.TypeOf([0]interface{ Super() ResourceClass.Instance }{}).Elem()):
		vtype = gd.TypeObject
		hint |= PropertyHintResourceType
		hintString = nameOf(field.Type.Elem())
	default:
		vtype, ok = gd.VariantTypeOf(field.Type)
		if !ok {
			return gd.PropertyInfo{}, false
		}
		if vtype == gd.TypeArray && (field.Type.Kind() == reflect.Array || field.Type.Kind() == reflect.Slice) {
			elem := field.Type.Elem()
			etype, ok := gd.VariantTypeOf(elem)
			if !ok {
				return gd.PropertyInfo{}, false
			}
			if elem.Implements(reflect.TypeFor[ResourceClass.Any]()) {
				hintString = fmt.Sprintf("%d/%d:%s", gd.TypeObject, PropertyHintResourceType, nameOf(elem)) // MAKE_RESOURCE_TYPE_HINT
			} else {
				hint |= PropertyHintArrayType
				hintString = etype.String()
			}
		}
	}
	if field.Type.Implements(reflect.TypeOf([0]interface{ AsResource() ResourceClass.Instance }{}).Elem()) {
		hint |= PropertyHintResourceType
	}
	var usage = PropertyUsageStorage | PropertyUsageEditor
	if vtype == gd.TypeNil {
		usage |= PropertyUsageNilIsVariant
	}
	if rangeHint, ok := field.Tag.Lookup("range"); ok {
		hint |= PropertyHintRange
		hintString = rangeHint
	}
	return gd.PropertyInfo{
		Type:       vtype,
		Name:       gd.NewStringName(name),
		ClassName:  gd.NewStringName(nameOf(field.Type)),
		Hint:       int64(hint),
		HintString: gd.NewString(hintString),
		Usage:      int64(usage),
	}, true
}

// Set needs to reference++ any resources that are sucessfully set.
func (instance *instanceImplementation) Set(name gd.StringName, value gd.Variant) bool {
	if impl, ok := instance.Value.(interface {
		Set(string, any) bool
	}); ok {
		ok := bool(impl.Set(name.String(), value.Interface()))
		if ok {
			if impl, ok := instance.Value.(interface {
				OnSet(string, any)
			}); ok {
				impl.OnSet(name.String(), value.Interface())
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
	if value.Type() == gd.TypeNil {
		field.Set(reflect.Zero(field.Type()))
		return true
	}
	var isExtensionClass bool
	var converted reflect.Value
	if value.Type() == gd.TypeObject {
		obj := gd.LetVariantAsPointerType[gd.Object](value, gd.TypeObject)
		ext, ok := gd.ExtensionInstances.Load(pointers.Get(obj)[0])
		if ok {
			converted = reflect.ValueOf(ext)
			isExtensionClass = true
		}
	}
	if !converted.IsValid() {
		var err error
		converted, err = gd.ConvertToDesiredGoType(value, field.Type())
		if err != nil {
			return false
		}
	}
	if converted.Kind() == reflect.Array || isExtensionClass {
		if !field.IsZero() {
			freeable, ok := field.Interface().(interface{ Free() })
			if ok {
				freeable.Free()
			}
		}
		obj, ok := converted.Interface().(interface {
			AsObject() [1]gd.Object
		})
		if !ok {
			return false
		}
		ref, ok := gd.As[gd.RefCounted](obj.AsObject()[0])
		if ok {
			ref.Reference()
		}
		pointers.Pin(obj.AsObject()[0])
	}
	field.Set(converted)
	if impl, ok := instance.Value.(interface {
		OnSet(gd.StringName, gd.Variant)
	}); ok {
		impl.OnSet(name, value)
	}
	return true
}

func (instance *instanceImplementation) Get(name gd.StringName) (gd.Variant, bool) {
	if impl, ok := instance.Value.(interface {
		Get(string) any
	}); ok {
		return gd.NewVariant(impl.Get(name.String())), true
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
		if field.Type().Kind() == reflect.Chan || reflect.PointerTo(field.Type()).Implements(reflect.TypeOf([0]gd.IsSignal{}).Elem()) {
			return gd.Variant{}, false
		}
	}
	if field.Type().Implements(reflect.TypeFor[interface{ superType() reflect.Type }]()) {
		if field.IsZero() {
			return gd.Global.Variants.NewNil(), false
		}
		obj := field.Interface().(interface{ AsObject() [1]gd.Object }).AsObject()[0]
		vary := gd.NewVariant(obj)
		return vary, true
	}
	return gd.NewVariant(field.Interface()), true
}

func (instance *instanceImplementation) GetPropertyList() []gd.PropertyInfo {
	if impl, ok := instance.Value.(interface {
		GetPropertyList() []gd.PropertyInfo
	}); ok {
		return impl.GetPropertyList()
	}
	return nil
}

func (instance *instanceImplementation) PropertyCanRevert(name gd.StringName) bool {
	if impl, ok := instance.Value.(interface {
		PropertyCanRevert(string) bool
	}); ok {
		return bool(impl.PropertyCanRevert(name.String()))
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
func (instance *instanceImplementation) PropertyGetRevert(name gd.StringName) (gd.Variant, bool) {
	if impl, ok := instance.Value.(interface {
		PropertyGetRevert(string) (any, bool)
	}); ok {
		val, ok := impl.PropertyGetRevert(name.String())
		return gd.NewVariant(val), ok
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

func (instance *instanceImplementation) ValidateProperty(info *gd.PropertyInfo) bool {
	switch validate := instance.Value.(type) {
	case interface {
		ValidateProperty(*gd.PropertyInfo) bool
	}:
		return bool(validate.ValidateProperty(info))
	}
	return true
}
