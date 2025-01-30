// Package RDPipelineRasterizationState provides methods for working with RDPipelineRasterizationState object instances.
package RDPipelineRasterizationState

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
This object is used by [RenderingDevice].
*/
type Instance [1]gdclass.RDPipelineRasterizationState

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRDPipelineRasterizationState() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RDPipelineRasterizationState

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDPipelineRasterizationState"))
	casted := Instance{*(*gdclass.RDPipelineRasterizationState)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) EnableDepthClamp() bool {
	return bool(class(self).GetEnableDepthClamp())
}

func (self Instance) SetEnableDepthClamp(value bool) {
	class(self).SetEnableDepthClamp(value)
}

func (self Instance) DiscardPrimitives() bool {
	return bool(class(self).GetDiscardPrimitives())
}

func (self Instance) SetDiscardPrimitives(value bool) {
	class(self).SetDiscardPrimitives(value)
}

func (self Instance) Wireframe() bool {
	return bool(class(self).GetWireframe())
}

func (self Instance) SetWireframe(value bool) {
	class(self).SetWireframe(value)
}

func (self Instance) CullMode() gdclass.RenderingDevicePolygonCullMode {
	return gdclass.RenderingDevicePolygonCullMode(class(self).GetCullMode())
}

func (self Instance) SetCullMode(value gdclass.RenderingDevicePolygonCullMode) {
	class(self).SetCullMode(value)
}

func (self Instance) FrontFace() gdclass.RenderingDevicePolygonFrontFace {
	return gdclass.RenderingDevicePolygonFrontFace(class(self).GetFrontFace())
}

func (self Instance) SetFrontFace(value gdclass.RenderingDevicePolygonFrontFace) {
	class(self).SetFrontFace(value)
}

func (self Instance) DepthBiasEnabled() bool {
	return bool(class(self).GetDepthBiasEnabled())
}

func (self Instance) SetDepthBiasEnabled(value bool) {
	class(self).SetDepthBiasEnabled(value)
}

func (self Instance) DepthBiasConstantFactor() Float.X {
	return Float.X(Float.X(class(self).GetDepthBiasConstantFactor()))
}

func (self Instance) SetDepthBiasConstantFactor(value Float.X) {
	class(self).SetDepthBiasConstantFactor(float64(value))
}

func (self Instance) DepthBiasClamp() Float.X {
	return Float.X(Float.X(class(self).GetDepthBiasClamp()))
}

func (self Instance) SetDepthBiasClamp(value Float.X) {
	class(self).SetDepthBiasClamp(float64(value))
}

func (self Instance) DepthBiasSlopeFactor() Float.X {
	return Float.X(Float.X(class(self).GetDepthBiasSlopeFactor()))
}

func (self Instance) SetDepthBiasSlopeFactor(value Float.X) {
	class(self).SetDepthBiasSlopeFactor(float64(value))
}

func (self Instance) LineWidth() Float.X {
	return Float.X(Float.X(class(self).GetLineWidth()))
}

func (self Instance) SetLineWidth(value Float.X) {
	class(self).SetLineWidth(float64(value))
}

func (self Instance) PatchControlPoints() int {
	return int(int(class(self).GetPatchControlPoints()))
}

func (self Instance) SetPatchControlPoints(value int) {
	class(self).SetPatchControlPoints(int64(value))
}

//go:nosplit
func (self class) SetEnableDepthClamp(p_member bool) { //gd:RDPipelineRasterizationState.set_enable_depth_clamp
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_set_enable_depth_clamp, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnableDepthClamp() bool { //gd:RDPipelineRasterizationState.get_enable_depth_clamp
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_get_enable_depth_clamp, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDiscardPrimitives(p_member bool) { //gd:RDPipelineRasterizationState.set_discard_primitives
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_set_discard_primitives, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDiscardPrimitives() bool { //gd:RDPipelineRasterizationState.get_discard_primitives
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_get_discard_primitives, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWireframe(p_member bool) { //gd:RDPipelineRasterizationState.set_wireframe
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_set_wireframe, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetWireframe() bool { //gd:RDPipelineRasterizationState.get_wireframe
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_get_wireframe, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCullMode(p_member gdclass.RenderingDevicePolygonCullMode) { //gd:RDPipelineRasterizationState.set_cull_mode
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_set_cull_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCullMode() gdclass.RenderingDevicePolygonCullMode { //gd:RDPipelineRasterizationState.get_cull_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDevicePolygonCullMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_get_cull_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrontFace(p_member gdclass.RenderingDevicePolygonFrontFace) { //gd:RDPipelineRasterizationState.set_front_face
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_set_front_face, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrontFace() gdclass.RenderingDevicePolygonFrontFace { //gd:RDPipelineRasterizationState.get_front_face
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDevicePolygonFrontFace](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_get_front_face, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDepthBiasEnabled(p_member bool) { //gd:RDPipelineRasterizationState.set_depth_bias_enabled
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_set_depth_bias_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDepthBiasEnabled() bool { //gd:RDPipelineRasterizationState.get_depth_bias_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_get_depth_bias_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDepthBiasConstantFactor(p_member float64) { //gd:RDPipelineRasterizationState.set_depth_bias_constant_factor
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_set_depth_bias_constant_factor, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDepthBiasConstantFactor() float64 { //gd:RDPipelineRasterizationState.get_depth_bias_constant_factor
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_get_depth_bias_constant_factor, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDepthBiasClamp(p_member float64) { //gd:RDPipelineRasterizationState.set_depth_bias_clamp
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_set_depth_bias_clamp, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDepthBiasClamp() float64 { //gd:RDPipelineRasterizationState.get_depth_bias_clamp
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_get_depth_bias_clamp, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDepthBiasSlopeFactor(p_member float64) { //gd:RDPipelineRasterizationState.set_depth_bias_slope_factor
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_set_depth_bias_slope_factor, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDepthBiasSlopeFactor() float64 { //gd:RDPipelineRasterizationState.get_depth_bias_slope_factor
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_get_depth_bias_slope_factor, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLineWidth(p_member float64) { //gd:RDPipelineRasterizationState.set_line_width
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_set_line_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLineWidth() float64 { //gd:RDPipelineRasterizationState.get_line_width
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_get_line_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPatchControlPoints(p_member int64) { //gd:RDPipelineRasterizationState.set_patch_control_points
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_set_patch_control_points, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPatchControlPoints() int64 { //gd:RDPipelineRasterizationState.get_patch_control_points
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineRasterizationState.Bind_get_patch_control_points, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRDPipelineRasterizationState() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsRDPipelineRasterizationState() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
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
	gdclass.Register("RDPipelineRasterizationState", func(ptr gd.Object) any {
		return [1]gdclass.RDPipelineRasterizationState{*(*gdclass.RDPipelineRasterizationState)(unsafe.Pointer(&ptr))}
	})
}
