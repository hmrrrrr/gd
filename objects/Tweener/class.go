package Tweener

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Tweeners are objects that perform a specific animating task, e.g. interpolating a property or calling a method at a given time. A [Tweener] can't be created manually, you need to use a dedicated method from [Tween].
*/
type Instance [1]classdb.Tweener
type Any interface {
	gd.IsClass
	AsTweener() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Tweener

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Tweener"))
	return Instance{classdb.Tweener(object)}
}

func (self Instance) OnFinished(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("finished"), gd.NewCallable(cb), 0)
}

func (self class) AsTweener() Advanced            { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTweener() Instance         { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("Tweener", func(ptr gd.Object) any { return [1]classdb.Tweener{classdb.Tweener(ptr)} })
}
