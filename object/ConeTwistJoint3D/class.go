package ConeTwistJoint3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Joint3D"
import "grow.graphics/gd/object/Node3D"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A physics joint that connects two 3D physics bodies in a way that simulates a ball-and-socket joint. The twist axis is initiated as the X axis of the [ConeTwistJoint3D]. Once the physics bodies swing, the twist axis is calculated as the middle of the X axes of the joint in the local space of the two physics bodies. Useful for limbs like shoulders and hips, lamps hanging off a ceiling, etc.

*/
type Simple [1]classdb.ConeTwistJoint3D
func (self Simple) SetParam(param classdb.ConeTwistJoint3DParam, value float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetParam(param, gd.Float(value))
}
func (self Simple) GetParam(param classdb.ConeTwistJoint3DParam) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetParam(param)))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ConeTwistJoint3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Sets the value of the specified parameter.
*/
//go:nosplit
func (self class) SetParam(param classdb.ConeTwistJoint3DParam, value gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConeTwistJoint3D.Bind_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the value of the specified parameter.
*/
//go:nosplit
func (self class) GetParam(param classdb.ConeTwistJoint3DParam) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ConeTwistJoint3D.Bind_get_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsConeTwistJoint3D() Expert { return self[0].AsConeTwistJoint3D() }


//go:nosplit
func (self Simple) AsConeTwistJoint3D() Simple { return self[0].AsConeTwistJoint3D() }


//go:nosplit
func (self class) AsJoint3D() Joint3D.Expert { return self[0].AsJoint3D() }


//go:nosplit
func (self Simple) AsJoint3D() Joint3D.Simple { return self[0].AsJoint3D() }


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

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("ConeTwistJoint3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type Param = classdb.ConeTwistJoint3DParam

const (
/*Swing is rotation from side to side, around the axis perpendicular to the twist axis.
The swing span defines, how much rotation will not get corrected along the swing axis.
Could be defined as looseness in the [ConeTwistJoint3D].
If below 0.05, this behavior is locked.*/
	ParamSwingSpan Param = 0
/*Twist is the rotation around the twist axis, this value defined how far the joint can twist.
Twist is locked if below 0.05.*/
	ParamTwistSpan Param = 1
/*The speed with which the swing or twist will take place.
The higher, the faster.*/
	ParamBias Param = 2
/*The ease with which the joint starts to twist. If it's too low, it takes more force to start twisting the joint.*/
	ParamSoftness Param = 3
/*Defines, how fast the swing- and twist-speed-difference on both sides gets synced.*/
	ParamRelaxation Param = 4
/*Represents the size of the [enum Param] enum.*/
	ParamMax Param = 5
)
