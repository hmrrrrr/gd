package SegmentShape2D

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Shape2D"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A 2D line segment shape, intended for use in physics. Usually used to provide a shape for a [CollisionShape2D].

*/
type Simple [1]classdb.SegmentShape2D
func (self Simple) SetA(a gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetA(a)
}
func (self Simple) GetA() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetA())
}
func (self Simple) SetB(b gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetB(b)
}
func (self Simple) GetB() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetB())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.SegmentShape2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetA(a gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, a)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SegmentShape2D.Bind_set_a, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetA() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SegmentShape2D.Bind_get_a, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetB(b gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, b)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SegmentShape2D.Bind_set_b, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetB() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.SegmentShape2D.Bind_get_b, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsSegmentShape2D() Expert { return self[0].AsSegmentShape2D() }


//go:nosplit
func (self Simple) AsSegmentShape2D() Simple { return self[0].AsSegmentShape2D() }


//go:nosplit
func (self class) AsShape2D() Shape2D.Expert { return self[0].AsShape2D() }


//go:nosplit
func (self Simple) AsShape2D() Shape2D.Simple { return self[0].AsShape2D() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("SegmentShape2D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
