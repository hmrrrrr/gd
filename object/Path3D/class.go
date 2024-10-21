package Path3D

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Node3D"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Can have [PathFollow3D] child nodes moving along the [Curve3D]. See [PathFollow3D] for more information on the usage.
Note that the path is considered as relative to the moved nodes (children of [PathFollow3D]). As such, the curve should usually start with a zero vector [code](0, 0, 0)[/code].

*/
type Simple [1]classdb.Path3D
func (self Simple) SetCurve(curve [1]classdb.Curve3D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCurve(curve)
}
func (self Simple) GetCurve() [1]classdb.Curve3D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Curve3D(Expert(self).GetCurve(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.Path3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetCurve(curve object.Curve3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(curve[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Path3D.Bind_set_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCurve(ctx gd.Lifetime) object.Curve3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Path3D.Bind_get_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Curve3D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsPath3D() Expert { return self[0].AsPath3D() }


//go:nosplit
func (self Simple) AsPath3D() Simple { return self[0].AsPath3D() }


//go:nosplit
func (self class) AsNode3D() Node3D.Expert { return self[0].AsNode3D() }


//go:nosplit
func (self Simple) AsNode3D() Node3D.Simple { return self[0].AsNode3D() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("Path3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
