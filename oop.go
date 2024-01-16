//go:build !generate

package gd

import (
	"reflect"

	"runtime.link/mmm"
)

type Class[T, S IsClass] struct {
	_     [0]*T
	super S
}

func (class *Class[T, S]) SetPointer(ptr Pointer) {
	any(&class.super).(PointerToClass).SetPointer(ptr)
}

func (class Class[T, S]) getPointer() Pointer {
	return class.super.getPointer()
}

func (class Class[T, S]) Context() mmm.Context {
	return class.super.Context()
}

func (class Class[T, S]) Pointer() uintptr {
	return class.super.Pointer()
}

func (class Class[T, S]) class() S { return class.super }

func (class *Class[T, S]) Super() *S { return &class.super }

func (class Class[T, S]) virtual(s string) reflect.Value {
	return class.super.virtual(s)
}

func VirtualByName(class IsClass, name string) reflect.Value {
	return class.virtual(name)
}

type Singleton interface {
	isSingleton()
}

type Pointer mmm.Pointer[API, Pointer, uintptr]

func (ptr Pointer) Free() {
	ptr.API.Object.Destroy(ptr.Pointer())
}

func (ptr Pointer) virtual(string) reflect.Value {
	return reflect.Value{}
}

func (ptr Pointer) getPointer() Pointer    { return ptr }
func (ptr *Pointer) SetPointer(to Pointer) { *ptr = to }

type Extends[T IsClass] interface {
	IsClass
	class() T
}

type PointerToClass interface {
	IsClass
	SetPointer(Pointer)
}

type IsClass interface {
	Context() mmm.Context
	Pointer() uintptr
	getPointer() Pointer
	virtual(string) reflect.Value
}

// MarkFree marks the given class as being freed.
func MarkFree(class IsClass) {
	mmm.MarkFree(class.getPointer())
}
