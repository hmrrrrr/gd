// Package XRPose provides methods for working with XRPose object instances.
package XRPose

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Transform3D"
import "graphics.gd/variant/Vector3"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
XR runtimes often identify multiple locations on devices such as controllers that are spatially tracked.
Orientation, location, linear velocity and angular velocity are all provided for each pose by the XR runtime. This object contains this state of a pose.
*/
type Instance [1]gdclass.XRPose

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsXRPose() Instance
}

/*
Returns the [member transform] with world scale and our reference frame applied. This is the transform used to position [XRNode3D] objects.
*/
func (self Instance) GetAdjustedTransform() Transform3D.BasisOrigin { //gd:XRPose.get_adjusted_transform
	return Transform3D.BasisOrigin(class(self).GetAdjustedTransform())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.XRPose

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("XRPose"))
	casted := Instance{*(*gdclass.XRPose)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) HasTrackingData() bool {
	return bool(class(self).GetHasTrackingData())
}

func (self Instance) SetHasTrackingData(value bool) {
	class(self).SetHasTrackingData(value)
}

func (self Instance) Name() string {
	return string(class(self).GetName().String())
}

func (self Instance) SetName(value string) {
	class(self).SetName(String.Name(String.New(value)))
}

func (self Instance) Transform() Transform3D.BasisOrigin {
	return Transform3D.BasisOrigin(class(self).GetTransform())
}

func (self Instance) SetTransform(value Transform3D.BasisOrigin) {
	class(self).SetTransform(Transform3D.BasisOrigin(value))
}

func (self Instance) LinearVelocity() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetLinearVelocity())
}

func (self Instance) SetLinearVelocity(value Vector3.XYZ) {
	class(self).SetLinearVelocity(Vector3.XYZ(value))
}

func (self Instance) AngularVelocity() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetAngularVelocity())
}

func (self Instance) SetAngularVelocity(value Vector3.XYZ) {
	class(self).SetAngularVelocity(Vector3.XYZ(value))
}

func (self Instance) TrackingConfidence() gdclass.XRPoseTrackingConfidence {
	return gdclass.XRPoseTrackingConfidence(class(self).GetTrackingConfidence())
}

func (self Instance) SetTrackingConfidence(value gdclass.XRPoseTrackingConfidence) {
	class(self).SetTrackingConfidence(value)
}

//go:nosplit
func (self class) SetHasTrackingData(has_tracking_data bool) { //gd:XRPose.set_has_tracking_data
	var frame = callframe.New()
	callframe.Arg(frame, has_tracking_data)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPose.Bind_set_has_tracking_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHasTrackingData() bool { //gd:XRPose.get_has_tracking_data
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPose.Bind_get_has_tracking_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetName(name String.Name) { //gd:XRPose.set_name
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPose.Bind_set_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetName() String.Name { //gd:XRPose.get_name
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPose.Bind_get_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTransform(transform Transform3D.BasisOrigin) { //gd:XRPose.set_transform
	var frame = callframe.New()
	callframe.Arg(frame, transform)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPose.Bind_set_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTransform() Transform3D.BasisOrigin { //gd:XRPose.get_transform
	var frame = callframe.New()
	var r_ret = callframe.Ret[Transform3D.BasisOrigin](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPose.Bind_get_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [member transform] with world scale and our reference frame applied. This is the transform used to position [XRNode3D] objects.
*/
//go:nosplit
func (self class) GetAdjustedTransform() Transform3D.BasisOrigin { //gd:XRPose.get_adjusted_transform
	var frame = callframe.New()
	var r_ret = callframe.Ret[Transform3D.BasisOrigin](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPose.Bind_get_adjusted_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLinearVelocity(velocity Vector3.XYZ) { //gd:XRPose.set_linear_velocity
	var frame = callframe.New()
	callframe.Arg(frame, velocity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPose.Bind_set_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLinearVelocity() Vector3.XYZ { //gd:XRPose.get_linear_velocity
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPose.Bind_get_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAngularVelocity(velocity Vector3.XYZ) { //gd:XRPose.set_angular_velocity
	var frame = callframe.New()
	callframe.Arg(frame, velocity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPose.Bind_set_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAngularVelocity() Vector3.XYZ { //gd:XRPose.get_angular_velocity
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPose.Bind_get_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTrackingConfidence(tracking_confidence gdclass.XRPoseTrackingConfidence) { //gd:XRPose.set_tracking_confidence
	var frame = callframe.New()
	callframe.Arg(frame, tracking_confidence)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPose.Bind_set_tracking_confidence, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTrackingConfidence() gdclass.XRPoseTrackingConfidence { //gd:XRPose.get_tracking_confidence
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.XRPoseTrackingConfidence](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRPose.Bind_get_tracking_confidence, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsXRPose() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsXRPose() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("XRPose", func(ptr gd.Object) any { return [1]gdclass.XRPose{*(*gdclass.XRPose)(unsafe.Pointer(&ptr))} })
}

type TrackingConfidence = gdclass.XRPoseTrackingConfidence //gd:XRPose.TrackingConfidence

const (
	/*No tracking information is available for this pose.*/
	XrTrackingConfidenceNone TrackingConfidence = 0
	/*Tracking information may be inaccurate or estimated. For example, with inside out tracking this would indicate a controller may be (partially) obscured.*/
	XrTrackingConfidenceLow TrackingConfidence = 1
	/*Tracking information is considered accurate and up to date.*/
	XrTrackingConfidenceHigh TrackingConfidence = 2
)
