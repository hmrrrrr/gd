// Package OpenXRCompositionLayerEquirect provides methods for working with OpenXRCompositionLayerEquirect object instances.
package OpenXRCompositionLayerEquirect

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/classdb/OpenXRCompositionLayer"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Float"

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

/*
An OpenXR composition layer that allows rendering a [SubViewport] on an internal slice of a sphere.
*/
type Instance [1]gdclass.OpenXRCompositionLayerEquirect

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsOpenXRCompositionLayerEquirect() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.OpenXRCompositionLayerEquirect

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OpenXRCompositionLayerEquirect"))
	casted := Instance{*(*gdclass.OpenXRCompositionLayerEquirect)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Radius() Float.X {
	return Float.X(Float.X(class(self).GetRadius()))
}

func (self Instance) SetRadius(value Float.X) {
	class(self).SetRadius(gd.Float(value))
}

func (self Instance) CentralHorizontalAngle() Float.X {
	return Float.X(Float.X(class(self).GetCentralHorizontalAngle()))
}

func (self Instance) SetCentralHorizontalAngle(value Float.X) {
	class(self).SetCentralHorizontalAngle(gd.Float(value))
}

func (self Instance) UpperVerticalAngle() Float.X {
	return Float.X(Float.X(class(self).GetUpperVerticalAngle()))
}

func (self Instance) SetUpperVerticalAngle(value Float.X) {
	class(self).SetUpperVerticalAngle(gd.Float(value))
}

func (self Instance) LowerVerticalAngle() Float.X {
	return Float.X(Float.X(class(self).GetLowerVerticalAngle()))
}

func (self Instance) SetLowerVerticalAngle(value Float.X) {
	class(self).SetLowerVerticalAngle(gd.Float(value))
}

func (self Instance) FallbackSegments() int {
	return int(int(class(self).GetFallbackSegments()))
}

func (self Instance) SetFallbackSegments(value int) {
	class(self).SetFallbackSegments(gd.Int(value))
}

//go:nosplit
func (self class) SetRadius(radius gd.Float) { //gd:OpenXRCompositionLayerEquirect.set_radius
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayerEquirect.Bind_set_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRadius() gd.Float { //gd:OpenXRCompositionLayerEquirect.get_radius
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayerEquirect.Bind_get_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCentralHorizontalAngle(angle gd.Float) { //gd:OpenXRCompositionLayerEquirect.set_central_horizontal_angle
	var frame = callframe.New()
	callframe.Arg(frame, angle)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayerEquirect.Bind_set_central_horizontal_angle, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCentralHorizontalAngle() gd.Float { //gd:OpenXRCompositionLayerEquirect.get_central_horizontal_angle
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayerEquirect.Bind_get_central_horizontal_angle, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUpperVerticalAngle(angle gd.Float) { //gd:OpenXRCompositionLayerEquirect.set_upper_vertical_angle
	var frame = callframe.New()
	callframe.Arg(frame, angle)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayerEquirect.Bind_set_upper_vertical_angle, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUpperVerticalAngle() gd.Float { //gd:OpenXRCompositionLayerEquirect.get_upper_vertical_angle
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayerEquirect.Bind_get_upper_vertical_angle, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLowerVerticalAngle(angle gd.Float) { //gd:OpenXRCompositionLayerEquirect.set_lower_vertical_angle
	var frame = callframe.New()
	callframe.Arg(frame, angle)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayerEquirect.Bind_set_lower_vertical_angle, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLowerVerticalAngle() gd.Float { //gd:OpenXRCompositionLayerEquirect.get_lower_vertical_angle
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayerEquirect.Bind_get_lower_vertical_angle, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFallbackSegments(segments gd.Int) { //gd:OpenXRCompositionLayerEquirect.set_fallback_segments
	var frame = callframe.New()
	callframe.Arg(frame, segments)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayerEquirect.Bind_set_fallback_segments, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFallbackSegments() gd.Int { //gd:OpenXRCompositionLayerEquirect.get_fallback_segments
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayerEquirect.Bind_get_fallback_segments, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsOpenXRCompositionLayerEquirect() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsOpenXRCompositionLayerEquirect() Instance {
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
		return gd.VirtualByName(OpenXRCompositionLayer.Advanced(self.AsOpenXRCompositionLayer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(OpenXRCompositionLayer.Instance(self.AsOpenXRCompositionLayer()), name)
	}
}
func init() {
	gdclass.Register("OpenXRCompositionLayerEquirect", func(ptr gd.Object) any {
		return [1]gdclass.OpenXRCompositionLayerEquirect{*(*gdclass.OpenXRCompositionLayerEquirect)(unsafe.Pointer(&ptr))}
	})
}
