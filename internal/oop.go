//go:build !generate

package gd

import (
	"reflect"
	"strings"
	"sync"
	"unsafe"

	"graphics.gd/internal/pointers"
)

var ExtensionInstances sync.Map

type NotificationType int32

func PointerWithOwnershipTransferredToGo[T pointers.Generic[T, [3]uintptr]](ptr uintptr) T {
	return pointers.New[T]([3]uintptr{ptr})
}

func PointerBorrowedTemporarily[T pointers.Generic[T, [3]uintptr]](ptr uintptr) T {
	return pointers.Let[T]([3]uintptr{ptr})
}

func PointerWithOwnershipTransferredToGodot[T pointers.Generic[T, [3]uintptr]](ptr T) uintptr {
	raw := pointers.Get(ptr)
	pointers.Set(ptr, [3]uintptr{raw[0], uintptr(Global.Object.GetInstanceID([1]Object{pointers.Raw[Object]([3]uintptr{raw[0]})}))})
	pointers.Lay(ptr)
	if raw[1] != 0 {
		panic("illegal transfer of ownership from Go -> Godot")
	}
	return raw[0]
}

func PointerMustAssertInstanceID[T pointers.Generic[T, [3]uintptr]](ptr uintptr) T {
	if ptr == 0 {
		return T{}
	}
	return pointers.Let[T]([3]uintptr{ptr, uintptr(Global.Object.GetInstanceID([1]Object{pointers.Raw[Object]([3]uintptr{ptr})}))})
}

func PointerLifetimeBoundTo[T pointers.Generic[T, [3]uintptr]](obj [1]Object, ptr uintptr) T {
	return pointers.Let[T]([3]uintptr{ptr, 0})
}

func (self Object) AsObject() [1]Object {
	return [1]Object{self}
}

func (self RefCounted) Free() {
	_, ok := pointers.End(Object(self))
	if !ok {
		return
	}
}

func (self Object) Free() {
	raw, ok := pointers.End(self)
	if !ok {
		return
	}
	this := [1]Object{pointers.Raw[Object](raw)}
	//fmt.Println("FREE ", pointers.Trio[Object](self), this[0].GetClass().String())
	//	fmt.Println(runtime.Caller(2))
	// Important that we don't destroy RefCounted objects, instead
	// they should be unreferenced instead.
	ref := Global.Object.CastTo(this, Global.refCountedClassTag)
	if ref != ([1]Object{}) {
		if (*(*RefCounted)(unsafe.Pointer(&ref))).Unreference() {
			Global.Object.Destroy(this)
		}
	} else {
		Global.Object.Destroy(this)
	}
}

type Class[T any, S IsClass] struct {
	_     [0]*T
	super S
}

func (class *Class[T, S]) AsObject() [1]Object {
	return class.super.AsObject()
}

func (class Class[T, S]) class() S { return class.super } //lint:ignore U1000 false positive.

func (class *Class[T, S]) Super() *S { return &class.super }

func (class Class[T, S]) Virtual(s string) reflect.Value {
	return class.super.Virtual(s)
}

func VirtualByName(class IsClass, name string) reflect.Value {
	return class.Virtual(name)
}

func classNameOf(rtype reflect.Type) string {
	if rtype.Kind() == reflect.Ptr || rtype.Kind() == reflect.Array {
		return classNameOf(rtype.Elem())
	}
	if rtype.Implements(reflect.TypeOf([0]IsClass{}).Elem()) {
		if rtype.Field(0).Anonymous {
			if rename, ok := rtype.Field(0).Tag.Lookup("gd"); ok {
				return rename
			}
		}
		if rtype.Name() == "" && rtype.Field(0).Anonymous {
			return classNameOf(rtype.Field(0).Type)
		}
		return strings.TrimPrefix(rtype.Name(), "class")
	}
	return ""
}

// As attempts to cast the given class to T, returning true
// if the cast was successful.
func As[T IsClass](class IsClass) (T, bool) {
	/*if ref, ok := Global.Instances[pointers.Get(class.AsObject())[0]].(T); ok {
	extension := any(ref).(ExtensionClass)
	extension.SetPointer(class.AsObject())
	return ref, true
	}*/
	var zero T
	var rtype = reflect.TypeOf([0]T{}).Elem()
	if rtype.Kind() == reflect.Pointer {
		return zero, false
	}
	var classtag = Global.ClassDB.GetClassTag(NewStringName(classNameOf(rtype)))
	casted := Global.Object.CastTo(class.AsObject(), classtag)
	if casted != ([1]Object{}) && pointers.Get(casted[0]) != ([3]uintptr{}) {
		return (*(*T)(unsafe.Pointer(&casted))), true
	}
	return zero, false
}

func as[T any](v Variant) T {
	var zero T
	val := v.Interface()
	if val == nil {
		return zero
	}
	if obj, ok := val.(IsClass); ok {
		var classtag = Global.ClassDB.GetClassTag(obj.AsObject()[0].GetClass().StringName())
		casted := Global.Object.CastTo(obj.AsObject(), classtag)
		if casted != ([1]Object{}) && pointers.Get(casted[0]) != ([3]uintptr{}) {
			any(&zero).(PointerToClass).SetPointer(casted)
		}
		return zero
	}
	return val.(T)
}

type Singleton interface {
	IsSingleton()
}

type Extends[T IsClass] interface {
	class() T

	Virtual(string) reflect.Value
}

type PointerToClass interface {
	IsClass
	SetPointer([1]Object)
}

type IsClass interface {
	Virtual(string) reflect.Value
	AsObject() [1]Object
}
