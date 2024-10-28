package OpenXRCompositionLayerCylinder

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/OpenXRCompositionLayer"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
An OpenXR composition layer that allows rendering a [SubViewport] on an internal slice of a cylinder.

*/
type Go [1]classdb.OpenXRCompositionLayerCylinder
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.OpenXRCompositionLayerCylinder
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OpenXRCompositionLayerCylinder"))
	return Go{classdb.OpenXRCompositionLayerCylinder(object)}
}

func (self Go) Radius() float64 {
		return float64(float64(class(self).GetRadius()))
}

func (self Go) SetRadius(value float64) {
	class(self).SetRadius(gd.Float(value))
}

func (self Go) AspectRatio() float64 {
		return float64(float64(class(self).GetAspectRatio()))
}

func (self Go) SetAspectRatio(value float64) {
	class(self).SetAspectRatio(gd.Float(value))
}

func (self Go) CentralAngle() float64 {
		return float64(float64(class(self).GetCentralAngle()))
}

func (self Go) SetCentralAngle(value float64) {
	class(self).SetCentralAngle(gd.Float(value))
}

func (self Go) FallbackSegments() int {
		return int(int(class(self).GetFallbackSegments()))
}

func (self Go) SetFallbackSegments(value int) {
	class(self).SetFallbackSegments(gd.Int(value))
}

//go:nosplit
func (self class) SetRadius(radius gd.Float)  {
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
func (self class) SetAspectRatio(aspect_ratio gd.Float)  {
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
func (self class) SetCentralAngle(angle gd.Float)  {
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
func (self class) SetFallbackSegments(segments gd.Int)  {
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
func (self class) AsOpenXRCompositionLayerCylinder() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsOpenXRCompositionLayerCylinder() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsOpenXRCompositionLayer() OpenXRCompositionLayer.GD { return *((*OpenXRCompositionLayer.GD)(unsafe.Pointer(&self))) }
func (self Go) AsOpenXRCompositionLayer() OpenXRCompositionLayer.Go { return *((*OpenXRCompositionLayer.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsOpenXRCompositionLayer(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsOpenXRCompositionLayer(), name)
	}
}
func init() {classdb.Register("OpenXRCompositionLayerCylinder", func(ptr gd.Object) any { return classdb.OpenXRCompositionLayerCylinder(ptr) })}
