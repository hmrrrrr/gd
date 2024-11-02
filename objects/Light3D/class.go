package Light3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/VisualInstance3D"
import "grow.graphics/gd/objects/Node3D"
import "grow.graphics/gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Light3D is the [i]abstract[/i] base class for light nodes. As it can't be instantiated, it shouldn't be used directly. Other types of light nodes inherit from it. Light3D contains the common variables and parameters used for lighting.
*/
type Instance [1]classdb.Light3D

/*
Returns the [Color] of an idealized blackbody at the given [member light_temperature]. This value is calculated internally based on the [member light_temperature]. This [Color] is multiplied by [member light_color] before being sent to the [RenderingServer].
*/
func (self Instance) GetCorrelatedColor() gd.Color {
	return gd.Color(class(self).GetCorrelatedColor())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Light3D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Light3D"))
	return Instance{classdb.Light3D(object)}
}

func (self Instance) LightIntensityLumens() float64 {
	return float64(float64(class(self).GetParam(20)))
}

func (self Instance) SetLightIntensityLumens(value float64) {
	class(self).SetParam(20, gd.Float(value))
}

func (self Instance) LightIntensityLux() float64 {
	return float64(float64(class(self).GetParam(20)))
}

func (self Instance) SetLightIntensityLux(value float64) {
	class(self).SetParam(20, gd.Float(value))
}

func (self Instance) LightTemperature() float64 {
	return float64(float64(class(self).GetTemperature()))
}

func (self Instance) SetLightTemperature(value float64) {
	class(self).SetTemperature(gd.Float(value))
}

func (self Instance) LightColor() gd.Color {
	return gd.Color(class(self).GetColor())
}

func (self Instance) SetLightColor(value gd.Color) {
	class(self).SetColor(value)
}

func (self Instance) LightEnergy() float64 {
	return float64(float64(class(self).GetParam(0)))
}

func (self Instance) SetLightEnergy(value float64) {
	class(self).SetParam(0, gd.Float(value))
}

func (self Instance) LightIndirectEnergy() float64 {
	return float64(float64(class(self).GetParam(1)))
}

func (self Instance) SetLightIndirectEnergy(value float64) {
	class(self).SetParam(1, gd.Float(value))
}

func (self Instance) LightVolumetricFogEnergy() float64 {
	return float64(float64(class(self).GetParam(2)))
}

func (self Instance) SetLightVolumetricFogEnergy(value float64) {
	class(self).SetParam(2, gd.Float(value))
}

func (self Instance) LightProjector() objects.Texture2D {
	return objects.Texture2D(class(self).GetProjector())
}

func (self Instance) SetLightProjector(value objects.Texture2D) {
	class(self).SetProjector(value)
}

func (self Instance) LightSize() float64 {
	return float64(float64(class(self).GetParam(5)))
}

func (self Instance) SetLightSize(value float64) {
	class(self).SetParam(5, gd.Float(value))
}

func (self Instance) LightAngularDistance() float64 {
	return float64(float64(class(self).GetParam(5)))
}

func (self Instance) SetLightAngularDistance(value float64) {
	class(self).SetParam(5, gd.Float(value))
}

func (self Instance) LightNegative() bool {
	return bool(class(self).IsNegative())
}

func (self Instance) SetLightNegative(value bool) {
	class(self).SetNegative(value)
}

func (self Instance) LightSpecular() float64 {
	return float64(float64(class(self).GetParam(3)))
}

func (self Instance) SetLightSpecular(value float64) {
	class(self).SetParam(3, gd.Float(value))
}

func (self Instance) LightBakeMode() classdb.Light3DBakeMode {
	return classdb.Light3DBakeMode(class(self).GetBakeMode())
}

func (self Instance) SetLightBakeMode(value classdb.Light3DBakeMode) {
	class(self).SetBakeMode(value)
}

func (self Instance) LightCullMask() int {
	return int(int(class(self).GetCullMask()))
}

func (self Instance) SetLightCullMask(value int) {
	class(self).SetCullMask(gd.Int(value))
}

func (self Instance) ShadowEnabled() bool {
	return bool(class(self).HasShadow())
}

func (self Instance) SetShadowEnabled(value bool) {
	class(self).SetShadow(value)
}

func (self Instance) ShadowBias() float64 {
	return float64(float64(class(self).GetParam(15)))
}

func (self Instance) SetShadowBias(value float64) {
	class(self).SetParam(15, gd.Float(value))
}

func (self Instance) ShadowNormalBias() float64 {
	return float64(float64(class(self).GetParam(14)))
}

func (self Instance) SetShadowNormalBias(value float64) {
	class(self).SetParam(14, gd.Float(value))
}

func (self Instance) ShadowReverseCullFace() bool {
	return bool(class(self).GetShadowReverseCullFace())
}

func (self Instance) SetShadowReverseCullFace(value bool) {
	class(self).SetShadowReverseCullFace(value)
}

func (self Instance) ShadowTransmittanceBias() float64 {
	return float64(float64(class(self).GetParam(19)))
}

func (self Instance) SetShadowTransmittanceBias(value float64) {
	class(self).SetParam(19, gd.Float(value))
}

func (self Instance) ShadowOpacity() float64 {
	return float64(float64(class(self).GetParam(17)))
}

func (self Instance) SetShadowOpacity(value float64) {
	class(self).SetParam(17, gd.Float(value))
}

func (self Instance) ShadowBlur() float64 {
	return float64(float64(class(self).GetParam(18)))
}

func (self Instance) SetShadowBlur(value float64) {
	class(self).SetParam(18, gd.Float(value))
}

func (self Instance) DistanceFadeEnabled() bool {
	return bool(class(self).IsDistanceFadeEnabled())
}

func (self Instance) SetDistanceFadeEnabled(value bool) {
	class(self).SetEnableDistanceFade(value)
}

func (self Instance) DistanceFadeBegin() float64 {
	return float64(float64(class(self).GetDistanceFadeBegin()))
}

func (self Instance) SetDistanceFadeBegin(value float64) {
	class(self).SetDistanceFadeBegin(gd.Float(value))
}

func (self Instance) DistanceFadeShadow() float64 {
	return float64(float64(class(self).GetDistanceFadeShadow()))
}

func (self Instance) SetDistanceFadeShadow(value float64) {
	class(self).SetDistanceFadeShadow(gd.Float(value))
}

func (self Instance) DistanceFadeLength() float64 {
	return float64(float64(class(self).GetDistanceFadeLength()))
}

func (self Instance) SetDistanceFadeLength(value float64) {
	class(self).SetDistanceFadeLength(gd.Float(value))
}

func (self Instance) EditorOnly() bool {
	return bool(class(self).IsEditorOnly())
}

func (self Instance) SetEditorOnly(value bool) {
	class(self).SetEditorOnly(value)
}

//go:nosplit
func (self class) SetEditorOnly(editor_only bool) {
	var frame = callframe.New()
	callframe.Arg(frame, editor_only)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_set_editor_only, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsEditorOnly() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_is_editor_only, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the value of the specified [enum Light3D.Param] parameter.
*/
//go:nosplit
func (self class) SetParam(param classdb.Light3DParam, value gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the value of the specified [enum Light3D.Param] parameter.
*/
//go:nosplit
func (self class) GetParam(param classdb.Light3DParam) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_get_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShadow(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_set_shadow, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) HasShadow() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_has_shadow, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNegative(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_set_negative, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsNegative() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_is_negative, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCullMask(cull_mask gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, cull_mask)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_set_cull_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCullMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_get_cull_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnableDistanceFade(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_set_enable_distance_fade, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsDistanceFadeEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_is_distance_fade_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDistanceFadeBegin(distance gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_set_distance_fade_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDistanceFadeBegin() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_get_distance_fade_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDistanceFadeShadow(distance gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_set_distance_fade_shadow, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDistanceFadeShadow() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_get_distance_fade_shadow, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDistanceFadeLength(distance gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_set_distance_fade_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDistanceFadeLength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_get_distance_fade_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetColor(color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_get_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShadowReverseCullFace(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_set_shadow_reverse_cull_face, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetShadowReverseCullFace() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_get_shadow_reverse_cull_face, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBakeMode(bake_mode classdb.Light3DBakeMode) {
	var frame = callframe.New()
	callframe.Arg(frame, bake_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_set_bake_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBakeMode() classdb.Light3DBakeMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.Light3DBakeMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_get_bake_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetProjector(projector objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(projector[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_set_projector, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetProjector() objects.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_get_projector, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTemperature(temperature gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, temperature)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_set_temperature, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTemperature() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_get_temperature, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [Color] of an idealized blackbody at the given [member light_temperature]. This value is calculated internally based on the [member light_temperature]. This [Color] is multiplied by [member light_color] before being sent to the [RenderingServer].
*/
//go:nosplit
func (self class) GetCorrelatedColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light3D.Bind_get_correlated_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsLight3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsLight3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(self.AsVisualInstance3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualInstance3D(), name)
	}
}
func init() { classdb.Register("Light3D", func(ptr gd.Object) any { return classdb.Light3D(ptr) }) }

type Param = classdb.Light3DParam

const (
	/*Constant for accessing [member light_energy].*/
	ParamEnergy Param = 0
	/*Constant for accessing [member light_indirect_energy].*/
	ParamIndirectEnergy Param = 1
	/*Constant for accessing [member light_volumetric_fog_energy].*/
	ParamVolumetricFogEnergy Param = 2
	/*Constant for accessing [member light_specular].*/
	ParamSpecular Param = 3
	/*Constant for accessing [member OmniLight3D.omni_range] or [member SpotLight3D.spot_range].*/
	ParamRange Param = 4
	/*Constant for accessing [member light_size].*/
	ParamSize Param = 5
	/*Constant for accessing [member OmniLight3D.omni_attenuation] or [member SpotLight3D.spot_attenuation].*/
	ParamAttenuation Param = 6
	/*Constant for accessing [member SpotLight3D.spot_angle].*/
	ParamSpotAngle Param = 7
	/*Constant for accessing [member SpotLight3D.spot_angle_attenuation].*/
	ParamSpotAttenuation Param = 8
	/*Constant for accessing [member DirectionalLight3D.directional_shadow_max_distance].*/
	ParamShadowMaxDistance Param = 9
	/*Constant for accessing [member DirectionalLight3D.directional_shadow_split_1].*/
	ParamShadowSplit1Offset Param = 10
	/*Constant for accessing [member DirectionalLight3D.directional_shadow_split_2].*/
	ParamShadowSplit2Offset Param = 11
	/*Constant for accessing [member DirectionalLight3D.directional_shadow_split_3].*/
	ParamShadowSplit3Offset Param = 12
	/*Constant for accessing [member DirectionalLight3D.directional_shadow_fade_start].*/
	ParamShadowFadeStart Param = 13
	/*Constant for accessing [member shadow_normal_bias].*/
	ParamShadowNormalBias Param = 14
	/*Constant for accessing [member shadow_bias].*/
	ParamShadowBias Param = 15
	/*Constant for accessing [member DirectionalLight3D.directional_shadow_pancake_size].*/
	ParamShadowPancakeSize Param = 16
	/*Constant for accessing [member shadow_opacity].*/
	ParamShadowOpacity Param = 17
	/*Constant for accessing [member shadow_blur].*/
	ParamShadowBlur Param = 18
	/*Constant for accessing [member shadow_transmittance_bias].*/
	ParamTransmittanceBias Param = 19
	/*Constant for accessing [member light_intensity_lumens] and [member light_intensity_lux]. Only used when [member ProjectSettings.rendering/lights_and_shadows/use_physical_light_units] is [code]true[/code].*/
	ParamIntensity Param = 20
	/*Represents the size of the [enum Param] enum.*/
	ParamMax Param = 21
)

type BakeMode = classdb.Light3DBakeMode

const (
	/*Light is ignored when baking. This is the fastest mode, but the light will be taken into account when baking global illumination. This mode should generally be used for dynamic lights that change quickly, as the effect of global illumination is less noticeable on those lights.
	  [b]Note:[/b] Hiding a light does [i]not[/i] affect baking [LightmapGI]. Hiding a light will still affect baking [VoxelGI] and SDFGI (see [member Environment.sdfgi_enabled]).*/
	BakeDisabled BakeMode = 0
	/*Light is taken into account in static baking ([VoxelGI], [LightmapGI], SDFGI ([member Environment.sdfgi_enabled])). The light can be moved around or modified, but its global illumination will not update in real-time. This is suitable for subtle changes (such as flickering torches), but generally not large changes such as toggling a light on and off.
	  [b]Note:[/b] The light is not baked in [LightmapGI] if [member editor_only] is [code]true[/code].*/
	BakeStatic BakeMode = 1
	/*Light is taken into account in dynamic baking ([VoxelGI] and SDFGI ([member Environment.sdfgi_enabled]) only). The light can be moved around or modified with global illumination updating in real-time. The light's global illumination appearance will be slightly different compared to [constant BAKE_STATIC]. This has a greater performance cost compared to [constant BAKE_STATIC]. When using SDFGI, the update speed of dynamic lights is affected by [member ProjectSettings.rendering/global_illumination/sdfgi/frames_to_update_lights].*/
	BakeDynamic BakeMode = 2
)
