package XRPositionalTracker

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/XRTracker"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
An instance of this object represents a device that is tracked, such as a controller or anchor point. HMDs aren't represented here as they are handled internally.
As controllers are turned on and the [XRInterface] detects them, instances of this object are automatically added to this list of active tracking objects accessible through the [XRServer].
The [XRNode3D] and [XRAnchor3D] both consume objects of this type and should be used in your project. The positional trackers are just under-the-hood objects that make this all work. These are mostly exposed so that GDExtension-based interfaces can interact with them.
*/
type Instance [1]classdb.XRPositionalTracker

/*
Returns [code]true[/code] if the tracker is available and is currently tracking the bound [param name] pose.
*/
func (self Instance) HasPose(name string) bool {
	return bool(class(self).HasPose(gd.NewStringName(name)))
}

/*
Returns the current [XRPose] state object for the bound [param name] pose.
*/
func (self Instance) GetPose(name string) gdclass.XRPose {
	return gdclass.XRPose(class(self).GetPose(gd.NewStringName(name)))
}

/*
Marks this pose as invalid, we don't clear the last reported state but it allows users to decide if trackers need to be hidden if we lose tracking or just remain at their last known position.
*/
func (self Instance) InvalidatePose(name string) {
	class(self).InvalidatePose(gd.NewStringName(name))
}

/*
Sets the transform, linear velocity, angular velocity and tracking confidence for the given pose. This method is called by a [XRInterface] implementation and should not be used directly.
*/
func (self Instance) SetPose(name string, transform gd.Transform3D, linear_velocity gd.Vector3, angular_velocity gd.Vector3, tracking_confidence classdb.XRPoseTrackingConfidence) {
	class(self).SetPose(gd.NewStringName(name), transform, linear_velocity, angular_velocity, tracking_confidence)
}

/*
Returns an input for this tracker. It can return a boolean, float or [Vector2] value depending on whether the input is a button, trigger or thumbstick/thumbpad.
*/
func (self Instance) GetInput(name string) gd.Variant {
	return gd.Variant(class(self).GetInput(gd.NewStringName(name)))
}

/*
Changes the value for the given input. This method is called by a [XRInterface] implementation and should not be used directly.
*/
func (self Instance) SetInput(name string, value gd.Variant) {
	class(self).SetInput(gd.NewStringName(name), value)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.XRPositionalTracker

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("XRPositionalTracker"))
	return Instance{classdb.XRPositionalTracker(object)}
}

func (self Instance) Profile() string {
	return string(class(self).GetTrackerProfile().String())
}

func (self Instance) SetProfile(value string) {
	class(self).SetTrackerProfile(gd.NewString(value))
}

func (self Instance) Hand() classdb.XRPositionalTrackerTrackerHand {
	return classdb.XRPositionalTrackerTrackerHand(class(self).GetTrackerHand())
}

func (self Instance) SetHand(value classdb.XRPositionalTrackerTrackerHand) {
	class(self).SetTrackerHand(value)
}

//go:nosplit
func (self class) GetTrackerProfile() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPositionalTracker.Bind_get_tracker_profile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTrackerProfile(profile gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(profile))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPositionalTracker.Bind_set_tracker_profile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTrackerHand() classdb.XRPositionalTrackerTrackerHand {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.XRPositionalTrackerTrackerHand](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPositionalTracker.Bind_get_tracker_hand, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTrackerHand(hand classdb.XRPositionalTrackerTrackerHand) {
	var frame = callframe.New()
	callframe.Arg(frame, hand)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPositionalTracker.Bind_set_tracker_hand, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the tracker is available and is currently tracking the bound [param name] pose.
*/
//go:nosplit
func (self class) HasPose(name gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPositionalTracker.Bind_has_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the current [XRPose] state object for the bound [param name] pose.
*/
//go:nosplit
func (self class) GetPose(name gd.StringName) gdclass.XRPose {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPositionalTracker.Bind_get_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.XRPose{classdb.XRPose(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Marks this pose as invalid, we don't clear the last reported state but it allows users to decide if trackers need to be hidden if we lose tracking or just remain at their last known position.
*/
//go:nosplit
func (self class) InvalidatePose(name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPositionalTracker.Bind_invalidate_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the transform, linear velocity, angular velocity and tracking confidence for the given pose. This method is called by a [XRInterface] implementation and should not be used directly.
*/
//go:nosplit
func (self class) SetPose(name gd.StringName, transform gd.Transform3D, linear_velocity gd.Vector3, angular_velocity gd.Vector3, tracking_confidence classdb.XRPoseTrackingConfidence) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, transform)
	callframe.Arg(frame, linear_velocity)
	callframe.Arg(frame, angular_velocity)
	callframe.Arg(frame, tracking_confidence)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPositionalTracker.Bind_set_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns an input for this tracker. It can return a boolean, float or [Vector2] value depending on whether the input is a button, trigger or thumbstick/thumbpad.
*/
//go:nosplit
func (self class) GetInput(name gd.StringName) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPositionalTracker.Bind_get_input, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Changes the value for the given input. This method is called by a [XRInterface] implementation and should not be used directly.
*/
//go:nosplit
func (self class) SetInput(name gd.StringName, value gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(value))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPositionalTracker.Bind_set_input, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Instance) OnPoseChanged(cb func(pose gdclass.XRPose)) {
	self[0].AsObject().Connect(gd.NewStringName("pose_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPoseLostTracking(cb func(pose gdclass.XRPose)) {
	self[0].AsObject().Connect(gd.NewStringName("pose_lost_tracking"), gd.NewCallable(cb), 0)
}

func (self Instance) OnButtonPressed(cb func(name string)) {
	self[0].AsObject().Connect(gd.NewStringName("button_pressed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnButtonReleased(cb func(name string)) {
	self[0].AsObject().Connect(gd.NewStringName("button_released"), gd.NewCallable(cb), 0)
}

func (self Instance) OnInputFloatChanged(cb func(name string, value float64)) {
	self[0].AsObject().Connect(gd.NewStringName("input_float_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnInputVector2Changed(cb func(name string, vector gd.Vector2)) {
	self[0].AsObject().Connect(gd.NewStringName("input_vector2_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnProfileChanged(cb func(role string)) {
	self[0].AsObject().Connect(gd.NewStringName("profile_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsXRPositionalTracker() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsXRPositionalTracker() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsXRTracker() XRTracker.Advanced {
	return *((*XRTracker.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsXRTracker() XRTracker.Instance {
	return *((*XRTracker.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsXRTracker(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsXRTracker(), name)
	}
}
func init() {
	classdb.Register("XRPositionalTracker", func(ptr gd.Object) any { return classdb.XRPositionalTracker(ptr) })
}

type TrackerHand = classdb.XRPositionalTrackerTrackerHand

const (
	/*The hand this tracker is held in is unknown or not applicable.*/
	TrackerHandUnknown TrackerHand = 0
	/*This tracker is the left hand controller.*/
	TrackerHandLeft TrackerHand = 1
	/*This tracker is the right hand controller.*/
	TrackerHandRight TrackerHand = 2
	/*Represents the size of the [enum TrackerHand] enum.*/
	TrackerHandMax TrackerHand = 3
)
