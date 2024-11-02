package DirectionalLight3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Light3D"
import "grow.graphics/gd/objects/VisualInstance3D"
import "grow.graphics/gd/objects/Node3D"
import "grow.graphics/gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A directional light is a type of [Light3D] node that models an infinite number of parallel rays covering the entire scene. It is used for lights with strong intensity that are located far away from the scene to model sunlight or moonlight. The worldspace location of the DirectionalLight3D transform (origin) is ignored. Only the basis is used to determine light direction.
*/
type Instance [1]classdb.DirectionalLight3D

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.DirectionalLight3D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("DirectionalLight3D"))
	return Instance{classdb.DirectionalLight3D(object)}
}

func (self Instance) DirectionalShadowMode() classdb.DirectionalLight3DShadowMode {
	return classdb.DirectionalLight3DShadowMode(class(self).GetShadowMode())
}

func (self Instance) SetDirectionalShadowMode(value classdb.DirectionalLight3DShadowMode) {
	class(self).SetShadowMode(value)
}

func (self Instance) DirectionalShadowBlendSplits() bool {
	return bool(class(self).IsBlendSplitsEnabled())
}

func (self Instance) SetDirectionalShadowBlendSplits(value bool) {
	class(self).SetBlendSplits(value)
}

func (self Instance) SkyMode() classdb.DirectionalLight3DSkyMode {
	return classdb.DirectionalLight3DSkyMode(class(self).GetSkyMode())
}

func (self Instance) SetSkyMode(value classdb.DirectionalLight3DSkyMode) {
	class(self).SetSkyMode(value)
}

//go:nosplit
func (self class) SetShadowMode(mode classdb.DirectionalLight3DShadowMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirectionalLight3D.Bind_set_shadow_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetShadowMode() classdb.DirectionalLight3DShadowMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.DirectionalLight3DShadowMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirectionalLight3D.Bind_get_shadow_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBlendSplits(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirectionalLight3D.Bind_set_blend_splits, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsBlendSplitsEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirectionalLight3D.Bind_is_blend_splits_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSkyMode(mode classdb.DirectionalLight3DSkyMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirectionalLight3D.Bind_set_sky_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSkyMode() classdb.DirectionalLight3DSkyMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.DirectionalLight3DSkyMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirectionalLight3D.Bind_get_sky_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsDirectionalLight3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsDirectionalLight3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsLight3D() Light3D.Advanced       { return *((*Light3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsLight3D() Light3D.Instance {
	return *((*Light3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualInstance3D() VisualInstance3D.Advanced {
	return *((*VisualInstance3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualInstance3D() VisualInstance3D.Instance {
	return *((*VisualInstance3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsLight3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsLight3D(), name)
	}
}
func init() {
	classdb.Register("DirectionalLight3D", func(ptr gd.Object) any { return classdb.DirectionalLight3D(ptr) })
}

type ShadowMode = classdb.DirectionalLight3DShadowMode

const (
	/*Renders the entire scene's shadow map from an orthogonal point of view. This is the fastest directional shadow mode. May result in blurrier shadows on close objects.*/
	ShadowOrthogonal ShadowMode = 0
	/*Splits the view frustum in 2 areas, each with its own shadow map. This shadow mode is a compromise between [constant SHADOW_ORTHOGONAL] and [constant SHADOW_PARALLEL_4_SPLITS] in terms of performance.*/
	ShadowParallel2Splits ShadowMode = 1
	/*Splits the view frustum in 4 areas, each with its own shadow map. This is the slowest directional shadow mode.*/
	ShadowParallel4Splits ShadowMode = 2
)

type SkyMode = classdb.DirectionalLight3DSkyMode

const (
	/*Makes the light visible in both scene lighting and sky rendering.*/
	SkyModeLightAndSky SkyMode = 0
	/*Makes the light visible in scene lighting only (including direct lighting and global illumination). When using this mode, the light will not be visible from sky shaders.*/
	SkyModeLightOnly SkyMode = 1
	/*Makes the light visible to sky shaders only. When using this mode the light will not cast light into the scene (either through direct lighting or through global illumination), but can be accessed through sky shaders. This can be useful, for example, when you want to control sky effects without illuminating the scene (during a night cycle, for example).*/
	SkyModeSkyOnly SkyMode = 2
)
