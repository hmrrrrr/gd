package Object

import (
	"reflect"
	"strings"
	"unsafe"

	gd "graphics.gd/internal"
	"graphics.gd/internal/pointers"
)

// As attempts to cast the given class to T, returning the
// casted value if successful, or panicking if not.
func As[T gd.IsClass](value gd.IsClass) T {
	casted, ok := Is[T](value)
	if !ok {
		panic("Object is not of the expected type")
	}
	return casted
}

// Is attempts to cast the given class to T, returning true
// if the cast was successful.
func Is[T gd.IsClass](value gd.IsClass) (T, bool) {
	ext, ok := gd.ExtensionInstances.Load(pointers.Get(value.AsObject()[0])[0])
	if ok {
		if ref, ok := ext.(T); ok {
			pointers.Lay(value.AsObject()[0])
			return ref, true
		}
	}
	var zero T
	var rtype = reflect.TypeFor[T]()
	if rtype.Kind() == reflect.Pointer {
		return zero, false
	}
	var classtag = gd.Global.ClassDB.GetClassTag(gd.NewStringName(nameOf(rtype)))
	casted := gd.Global.Object.CastTo(value.AsObject(), classtag)
	if casted != ([1]gd.Object{}) && pointers.Get(casted[0]) != ([3]uint64{}) {
		return (*(*T)(unsafe.Pointer(&casted))), true
	}
	return zero, false
}

func nameOf(rtype reflect.Type) string {
	if rtype.Kind() == reflect.Ptr || rtype.Kind() == reflect.Array {
		return nameOf(rtype.Elem())
	}
	if rtype.Kind() == reflect.Struct && rtype.NumField() > 0 {
		if rtype.Field(0).Anonymous {
			if rename, ok := rtype.Field(0).Tag.Lookup("gd"); ok {
				return rename
			}
			if rtype.Name() == "" {
				return nameOf(rtype.Field(0).Type)
			}
		}
		return strings.TrimPrefix(rtype.Name(), "class")
	}
	return ""
}
