package SeparationRayShape2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Shape2D"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A 2D ray shape, intended for use in physics. Usually used to provide a shape for a [CollisionShape2D]. When a [SeparationRayShape2D] collides with an object, it tries to separate itself from it by moving its endpoint to the collision point. For example, a [SeparationRayShape2D] next to a character can allow it to instantly move up when touching stairs.
*/
type Instance [1]classdb.SeparationRayShape2D

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.SeparationRayShape2D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SeparationRayShape2D"))
	return Instance{classdb.SeparationRayShape2D(object)}
}

func (self Instance) Length() float64 {
	return float64(float64(class(self).GetLength()))
}

func (self Instance) SetLength(value float64) {
	class(self).SetLength(gd.Float(value))
}

func (self Instance) SlideOnSlope() bool {
	return bool(class(self).GetSlideOnSlope())
}

func (self Instance) SetSlideOnSlope(value bool) {
	class(self).SetSlideOnSlope(value)
}

//go:nosplit
func (self class) SetLength(length gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SeparationRayShape2D.Bind_set_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SeparationRayShape2D.Bind_get_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSlideOnSlope(active bool) {
	var frame = callframe.New()
	callframe.Arg(frame, active)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SeparationRayShape2D.Bind_set_slide_on_slope, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSlideOnSlope() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SeparationRayShape2D.Bind_get_slide_on_slope, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsSeparationRayShape2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSeparationRayShape2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsShape2D() Shape2D.Advanced         { return *((*Shape2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsShape2D() Shape2D.Instance {
	return *((*Shape2D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsShape2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsShape2D(), name)
	}
}
func init() {
	classdb.Register("SeparationRayShape2D", func(ptr gd.Object) any { return classdb.SeparationRayShape2D(ptr) })
}
