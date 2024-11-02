package ImporterMeshInstance3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Node3D"
import "grow.graphics/gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

type Instance [1]classdb.ImporterMeshInstance3D

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ImporterMeshInstance3D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ImporterMeshInstance3D"))
	return Instance{classdb.ImporterMeshInstance3D(object)}
}

func (self Instance) Mesh() objects.ImporterMesh {
	return objects.ImporterMesh(class(self).GetMesh())
}

func (self Instance) SetMesh(value objects.ImporterMesh) {
	class(self).SetMesh(value)
}

func (self Instance) Skin() objects.Skin {
	return objects.Skin(class(self).GetSkin())
}

func (self Instance) SetSkin(value objects.Skin) {
	class(self).SetSkin(value)
}

func (self Instance) SkeletonPath() string {
	return string(class(self).GetSkeletonPath().String())
}

func (self Instance) SetSkeletonPath(value string) {
	class(self).SetSkeletonPath(gd.NewString(value).NodePath())
}

func (self Instance) LayerMask() int {
	return int(int(class(self).GetLayerMask()))
}

func (self Instance) SetLayerMask(value int) {
	class(self).SetLayerMask(gd.Int(value))
}

func (self Instance) CastShadow() classdb.GeometryInstance3DShadowCastingSetting {
	return classdb.GeometryInstance3DShadowCastingSetting(class(self).GetCastShadowsSetting())
}

func (self Instance) SetCastShadow(value classdb.GeometryInstance3DShadowCastingSetting) {
	class(self).SetCastShadowsSetting(value)
}

func (self Instance) VisibilityRangeBegin() float64 {
	return float64(float64(class(self).GetVisibilityRangeBegin()))
}

func (self Instance) SetVisibilityRangeBegin(value float64) {
	class(self).SetVisibilityRangeBegin(gd.Float(value))
}

func (self Instance) VisibilityRangeBeginMargin() float64 {
	return float64(float64(class(self).GetVisibilityRangeBeginMargin()))
}

func (self Instance) SetVisibilityRangeBeginMargin(value float64) {
	class(self).SetVisibilityRangeBeginMargin(gd.Float(value))
}

func (self Instance) VisibilityRangeEnd() float64 {
	return float64(float64(class(self).GetVisibilityRangeEnd()))
}

func (self Instance) SetVisibilityRangeEnd(value float64) {
	class(self).SetVisibilityRangeEnd(gd.Float(value))
}

func (self Instance) VisibilityRangeEndMargin() float64 {
	return float64(float64(class(self).GetVisibilityRangeEndMargin()))
}

func (self Instance) SetVisibilityRangeEndMargin(value float64) {
	class(self).SetVisibilityRangeEndMargin(gd.Float(value))
}

func (self Instance) VisibilityRangeFadeMode() classdb.GeometryInstance3DVisibilityRangeFadeMode {
	return classdb.GeometryInstance3DVisibilityRangeFadeMode(class(self).GetVisibilityRangeFadeMode())
}

func (self Instance) SetVisibilityRangeFadeMode(value classdb.GeometryInstance3DVisibilityRangeFadeMode) {
	class(self).SetVisibilityRangeFadeMode(value)
}

//go:nosplit
func (self class) SetMesh(mesh objects.ImporterMesh) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(mesh[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_set_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMesh() objects.ImporterMesh {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_get_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.ImporterMesh{classdb.ImporterMesh(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSkin(skin objects.Skin) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(skin[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_set_skin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSkin() objects.Skin {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_get_skin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Skin{classdb.Skin(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSkeletonPath(skeleton_path gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(skeleton_path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_set_skeleton_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSkeletonPath() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_get_skeleton_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLayerMask(layer_mask gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, layer_mask)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_set_layer_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLayerMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_get_layer_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCastShadowsSetting(shadow_casting_setting classdb.GeometryInstance3DShadowCastingSetting) {
	var frame = callframe.New()
	callframe.Arg(frame, shadow_casting_setting)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_set_cast_shadows_setting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCastShadowsSetting() classdb.GeometryInstance3DShadowCastingSetting {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GeometryInstance3DShadowCastingSetting](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_get_cast_shadows_setting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisibilityRangeEndMargin(distance gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_set_visibility_range_end_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVisibilityRangeEndMargin() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_get_visibility_range_end_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisibilityRangeEnd(distance gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_set_visibility_range_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVisibilityRangeEnd() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_get_visibility_range_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisibilityRangeBeginMargin(distance gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_set_visibility_range_begin_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVisibilityRangeBeginMargin() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_get_visibility_range_begin_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisibilityRangeBegin(distance gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_set_visibility_range_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVisibilityRangeBegin() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_get_visibility_range_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisibilityRangeFadeMode(mode classdb.GeometryInstance3DVisibilityRangeFadeMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_set_visibility_range_fade_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVisibilityRangeFadeMode() classdb.GeometryInstance3DVisibilityRangeFadeMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GeometryInstance3DVisibilityRangeFadeMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImporterMeshInstance3D.Bind_get_visibility_range_fade_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsImporterMeshInstance3D() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsImporterMeshInstance3D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode3D(), name)
	}
}
func init() {
	classdb.Register("ImporterMeshInstance3D", func(ptr gd.Object) any { return classdb.ImporterMeshInstance3D(ptr) })
}
