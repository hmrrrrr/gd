package CapsuleShape2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Shape2D"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
A 2D capsule shape, intended for use in physics. Usually used to provide a shape for a [CollisionShape2D].
[b]Performance:[/b] [CapsuleShape2D] is fast to check collisions against, but it is slower than [RectangleShape2D] and [CircleShape2D].

*/
type Go [1]classdb.CapsuleShape2D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.CapsuleShape2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CapsuleShape2D"))
	return Go{classdb.CapsuleShape2D(object)}
}

func (self Go) Radius() float64 {
		return float64(float64(class(self).GetRadius()))
}

func (self Go) SetRadius(value float64) {
	class(self).SetRadius(gd.Float(value))
}

func (self Go) Height() float64 {
		return float64(float64(class(self).GetHeight()))
}

func (self Go) SetHeight(value float64) {
	class(self).SetHeight(gd.Float(value))
}

//go:nosplit
func (self class) SetRadius(radius gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CapsuleShape2D.Bind_set_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRadius() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CapsuleShape2D.Bind_get_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHeight(height gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CapsuleShape2D.Bind_set_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHeight() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CapsuleShape2D.Bind_get_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCapsuleShape2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsCapsuleShape2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsShape2D() Shape2D.GD { return *((*Shape2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsShape2D() Shape2D.Go { return *((*Shape2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsShape2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsShape2D(), name)
	}
}
func init() {classdb.Register("CapsuleShape2D", func(ptr gd.Object) any { return classdb.CapsuleShape2D(ptr) })}
