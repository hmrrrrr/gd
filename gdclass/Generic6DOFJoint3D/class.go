package Generic6DOFJoint3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Joint3D"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
The [Generic6DOFJoint3D] (6 Degrees Of Freedom) joint allows for implementing custom types of joints by locking the rotation and translation of certain axes.
The first 3 DOF represent the linear motion of the physics bodies and the last 3 DOF represent the angular motion of the physics bodies. Each axis can be either locked, or limited.

*/
type Go [1]classdb.Generic6DOFJoint3D
func (self Go) SetParamX(param classdb.Generic6DOFJoint3DParam, value float64) {
	class(self).SetParamX(param, gd.Float(value))
}
func (self Go) GetParamX(param classdb.Generic6DOFJoint3DParam) float64 {
	return float64(float64(class(self).GetParamX(param)))
}
func (self Go) SetParamY(param classdb.Generic6DOFJoint3DParam, value float64) {
	class(self).SetParamY(param, gd.Float(value))
}
func (self Go) GetParamY(param classdb.Generic6DOFJoint3DParam) float64 {
	return float64(float64(class(self).GetParamY(param)))
}
func (self Go) SetParamZ(param classdb.Generic6DOFJoint3DParam, value float64) {
	class(self).SetParamZ(param, gd.Float(value))
}
func (self Go) GetParamZ(param classdb.Generic6DOFJoint3DParam) float64 {
	return float64(float64(class(self).GetParamZ(param)))
}
func (self Go) SetFlagX(flag classdb.Generic6DOFJoint3DFlag, value bool) {
	class(self).SetFlagX(flag, value)
}
func (self Go) GetFlagX(flag classdb.Generic6DOFJoint3DFlag) bool {
	return bool(class(self).GetFlagX(flag))
}
func (self Go) SetFlagY(flag classdb.Generic6DOFJoint3DFlag, value bool) {
	class(self).SetFlagY(flag, value)
}
func (self Go) GetFlagY(flag classdb.Generic6DOFJoint3DFlag) bool {
	return bool(class(self).GetFlagY(flag))
}
func (self Go) SetFlagZ(flag classdb.Generic6DOFJoint3DFlag, value bool) {
	class(self).SetFlagZ(flag, value)
}
func (self Go) GetFlagZ(flag classdb.Generic6DOFJoint3DFlag) bool {
	return bool(class(self).GetFlagZ(flag))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Generic6DOFJoint3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Generic6DOFJoint3D"))
	return Go{classdb.Generic6DOFJoint3D(object)}
}

//go:nosplit
func (self class) SetParamX(param classdb.Generic6DOFJoint3DParam, value gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Generic6DOFJoint3D.Bind_set_param_x, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetParamX(param classdb.Generic6DOFJoint3DParam) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Generic6DOFJoint3D.Bind_get_param_x, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetParamY(param classdb.Generic6DOFJoint3DParam, value gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Generic6DOFJoint3D.Bind_set_param_y, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetParamY(param classdb.Generic6DOFJoint3DParam) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Generic6DOFJoint3D.Bind_get_param_y, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetParamZ(param classdb.Generic6DOFJoint3DParam, value gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Generic6DOFJoint3D.Bind_set_param_z, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetParamZ(param classdb.Generic6DOFJoint3DParam) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Generic6DOFJoint3D.Bind_get_param_z, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFlagX(flag classdb.Generic6DOFJoint3DFlag, value bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Generic6DOFJoint3D.Bind_set_flag_x, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFlagX(flag classdb.Generic6DOFJoint3DFlag) bool {
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Generic6DOFJoint3D.Bind_get_flag_x, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFlagY(flag classdb.Generic6DOFJoint3DFlag, value bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Generic6DOFJoint3D.Bind_set_flag_y, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFlagY(flag classdb.Generic6DOFJoint3DFlag) bool {
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Generic6DOFJoint3D.Bind_get_flag_y, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFlagZ(flag classdb.Generic6DOFJoint3DFlag, value bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Generic6DOFJoint3D.Bind_set_flag_z, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFlagZ(flag classdb.Generic6DOFJoint3DFlag) bool {
	var frame = callframe.New()
	callframe.Arg(frame, flag)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Generic6DOFJoint3D.Bind_get_flag_z, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsGeneric6DOFJoint3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsGeneric6DOFJoint3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsJoint3D() Joint3D.GD { return *((*Joint3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsJoint3D() Joint3D.Go { return *((*Joint3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsJoint3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsJoint3D(), name)
	}
}
func init() {classdb.Register("Generic6DOFJoint3D", func(ptr gd.Object) any { return classdb.Generic6DOFJoint3D(ptr) })}
type Param = classdb.Generic6DOFJoint3DParam

const (
/*The minimum difference between the pivot points' axes.*/
	ParamLinearLowerLimit Param = 0
/*The maximum difference between the pivot points' axes.*/
	ParamLinearUpperLimit Param = 1
/*A factor applied to the movement across the axes. The lower, the slower the movement.*/
	ParamLinearLimitSoftness Param = 2
/*The amount of restitution on the axes' movement. The lower, the more momentum gets lost.*/
	ParamLinearRestitution Param = 3
/*The amount of damping that happens at the linear motion across the axes.*/
	ParamLinearDamping Param = 4
/*The velocity the linear motor will try to reach.*/
	ParamLinearMotorTargetVelocity Param = 5
/*The maximum force the linear motor will apply while trying to reach the velocity target.*/
	ParamLinearMotorForceLimit Param = 6
	ParamLinearSpringStiffness Param = 7
	ParamLinearSpringDamping Param = 8
	ParamLinearSpringEquilibriumPoint Param = 9
/*The minimum rotation in negative direction to break loose and rotate around the axes.*/
	ParamAngularLowerLimit Param = 10
/*The minimum rotation in positive direction to break loose and rotate around the axes.*/
	ParamAngularUpperLimit Param = 11
/*The speed of all rotations across the axes.*/
	ParamAngularLimitSoftness Param = 12
/*The amount of rotational damping across the axes. The lower, the more damping occurs.*/
	ParamAngularDamping Param = 13
/*The amount of rotational restitution across the axes. The lower, the more restitution occurs.*/
	ParamAngularRestitution Param = 14
/*The maximum amount of force that can occur, when rotating around the axes.*/
	ParamAngularForceLimit Param = 15
/*When rotating across the axes, this error tolerance factor defines how much the correction gets slowed down. The lower, the slower.*/
	ParamAngularErp Param = 16
/*Target speed for the motor at the axes.*/
	ParamAngularMotorTargetVelocity Param = 17
/*Maximum acceleration for the motor at the axes.*/
	ParamAngularMotorForceLimit Param = 18
	ParamAngularSpringStiffness Param = 19
	ParamAngularSpringDamping Param = 20
	ParamAngularSpringEquilibriumPoint Param = 21
/*Represents the size of the [enum Param] enum.*/
	ParamMax Param = 22
)
type Flag = classdb.Generic6DOFJoint3DFlag

const (
/*If enabled, linear motion is possible within the given limits.*/
	FlagEnableLinearLimit Flag = 0
/*If enabled, rotational motion is possible within the given limits.*/
	FlagEnableAngularLimit Flag = 1
	FlagEnableLinearSpring Flag = 3
	FlagEnableAngularSpring Flag = 2
/*If enabled, there is a rotational motor across these axes.*/
	FlagEnableMotor Flag = 4
/*If enabled, there is a linear motor across these axes.*/
	FlagEnableLinearMotor Flag = 5
/*Represents the size of the [enum Flag] enum.*/
	FlagMax Flag = 6
)
