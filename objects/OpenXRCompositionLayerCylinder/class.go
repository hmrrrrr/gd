package OpenXRCompositionLayerCylinder

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/OpenXRCompositionLayer"
import "grow.graphics/gd/objects/Node3D"
import "grow.graphics/gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
An OpenXR composition layer that allows rendering a [SubViewport] on an internal slice of a cylinder.
*/
type Instance [1]classdb.OpenXRCompositionLayerCylinder

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.OpenXRCompositionLayerCylinder

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OpenXRCompositionLayerCylinder"))
	return Instance{classdb.OpenXRCompositionLayerCylinder(object)}
}

func (self Instance) Radius() float64 {
	return float64(float64(class(self).GetRadius()))
}

func (self Instance) SetRadius(value float64) {
	class(self).SetRadius(gd.Float(value))
}

func (self Instance) AspectRatio() float64 {
	return float64(float64(class(self).GetAspectRatio()))
}

func (self Instance) SetAspectRatio(value float64) {
	class(self).SetAspectRatio(gd.Float(value))
}

func (self Instance) CentralAngle() float64 {
	return float64(float64(class(self).GetCentralAngle()))
}

func (self Instance) SetCentralAngle(value float64) {
	class(self).SetCentralAngle(gd.Float(value))
}

func (self Instance) FallbackSegments() int {
	return int(int(class(self).GetFallbackSegments()))
}

func (self Instance) SetFallbackSegments(value int) {
	class(self).SetFallbackSegments(gd.Int(value))
}

//go:nosplit
func (self class) SetRadius(radius gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayerCylinder.Bind_set_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRadius() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayerCylinder.Bind_get_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAspectRatio(aspect_ratio gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, aspect_ratio)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayerCylinder.Bind_set_aspect_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAspectRatio() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayerCylinder.Bind_get_aspect_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCentralAngle(angle gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, angle)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayerCylinder.Bind_set_central_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCentralAngle() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayerCylinder.Bind_get_central_angle, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFallbackSegments(segments gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, segments)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayerCylinder.Bind_set_fallback_segments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFallbackSegments() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayerCylinder.Bind_get_fallback_segments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsOpenXRCompositionLayerCylinder() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsOpenXRCompositionLayerCylinder() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsOpenXRCompositionLayer() OpenXRCompositionLayer.Advanced {
	return *((*OpenXRCompositionLayer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsOpenXRCompositionLayer() OpenXRCompositionLayer.Instance {
	return *((*OpenXRCompositionLayer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsOpenXRCompositionLayer(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsOpenXRCompositionLayer(), name)
	}
}
func init() {
	classdb.Register("OpenXRCompositionLayerCylinder", func(ptr gd.Object) any { return classdb.OpenXRCompositionLayerCylinder(ptr) })
}
