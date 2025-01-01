package objects

import (
	"reflect"
	"strings"
	"unsafe"

	gd "graphics.gd/internal"
	"graphics.gd/internal/pointers"
)

// Instance of an object.
type Instance = gd.Object

// ID that uniquely identifies an object.
type ID gd.ObjectID

// IsValid returns true if the Object that corresponds to id is a valid object
// (e.g. has not been deleted from memory). All Objects have a unique instance ID.
func (id ID) IsValid() bool { //gd:is_instance_id_valid is_instance_valid
	return bool(gd.IsInstanceIdValid(gd.Int(id)))
}

// Get returns the Object that corresponds to id. All Objects have a unique instance ID.
func Get(id ID) gd.Object { //gd:instance_from_id
	return gd.InstanceFromId(gd.Int(id))
}

// New returns a new empty object instance.
func New() gd.Object {
	return gd.Global.ClassDB.ConstructObject(gd.NewStringName("Object"))
}

// As attempts to cast the given class to T, returning true
// if the cast was successful.
func As[T gd.IsClass](value gd.IsClass) (T, bool) {
	ext, ok := gd.ExtensionInstances.Load(pointers.Get(value.AsObject())[0])
	if ok {
		if ref, ok := ext.(T); ok {
			return ref, true
		}
	}
	var zero T
	var rtype = reflect.TypeFor[T]()
	if rtype.Kind() == reflect.Pointer {
		return zero, false
	}
	var classtag = gd.Global.ClassDB.GetClassTag(gd.NewStringName(NameOf(rtype)))
	casted := gd.Global.Object.CastTo(value.AsObject(), classtag)
	if casted != (gd.Object{}) && pointers.Get(casted) != ([3]uintptr{}) {
		return (*(*T)(unsafe.Pointer(&casted))), true
	}
	return zero, false
}

func NameOf(rtype reflect.Type) string {
	if rtype.Kind() == reflect.Ptr || rtype.Kind() == reflect.Array {
		return NameOf(rtype.Elem())
	}
	if rtype.Kind() == reflect.Struct && rtype.NumField() > 0 {
		if rtype.Field(0).Anonymous {
			if rename, ok := rtype.Field(0).Tag.Lookup("gd"); ok {
				return rename
			}
			if rtype.Name() == "" {
				return NameOf(rtype.Field(0).Type)
			}
		}
		return strings.TrimPrefix(rtype.Name(), "class")
	}
	return ""
}
