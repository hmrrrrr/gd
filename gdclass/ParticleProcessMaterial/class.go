package ParticleProcessMaterial

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Material"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
[ParticleProcessMaterial] defines particle properties and behavior. It is used in the [code]process_material[/code] of the [GPUParticles2D] and [GPUParticles3D] nodes. Some of this material's properties are applied to each particle when emitted, while others can have a [CurveTexture] or a [GradientTexture1D] applied to vary numerical or color values over the lifetime of the particle.
*/
type Instance [1]classdb.ParticleProcessMaterial

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ParticleProcessMaterial

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ParticleProcessMaterial"))
	return Instance{classdb.ParticleProcessMaterial(object)}
}

func (self Instance) LifetimeRandomness() float64 {
	return float64(float64(class(self).GetLifetimeRandomness()))
}

func (self Instance) SetLifetimeRandomness(value float64) {
	class(self).SetLifetimeRandomness(gd.Float(value))
}

func (self Instance) ParticleFlagAlignY() bool {
	return bool(class(self).GetParticleFlag(0))
}

func (self Instance) SetParticleFlagAlignY(value bool) {
	class(self).SetParticleFlag(0, value)
}

func (self Instance) ParticleFlagRotateY() bool {
	return bool(class(self).GetParticleFlag(1))
}

func (self Instance) SetParticleFlagRotateY(value bool) {
	class(self).SetParticleFlag(1, value)
}

func (self Instance) ParticleFlagDisableZ() bool {
	return bool(class(self).GetParticleFlag(2))
}

func (self Instance) SetParticleFlagDisableZ(value bool) {
	class(self).SetParticleFlag(2, value)
}

func (self Instance) ParticleFlagDampingAsFriction() bool {
	return bool(class(self).GetParticleFlag(3))
}

func (self Instance) SetParticleFlagDampingAsFriction(value bool) {
	class(self).SetParticleFlag(3, value)
}

func (self Instance) EmissionShapeOffset() gd.Vector3 {
	return gd.Vector3(class(self).GetEmissionShapeOffset())
}

func (self Instance) SetEmissionShapeOffset(value gd.Vector3) {
	class(self).SetEmissionShapeOffset(value)
}

func (self Instance) EmissionShapeScale() gd.Vector3 {
	return gd.Vector3(class(self).GetEmissionShapeScale())
}

func (self Instance) SetEmissionShapeScale(value gd.Vector3) {
	class(self).SetEmissionShapeScale(value)
}

func (self Instance) EmissionShape() classdb.ParticleProcessMaterialEmissionShape {
	return classdb.ParticleProcessMaterialEmissionShape(class(self).GetEmissionShape())
}

func (self Instance) SetEmissionShape(value classdb.ParticleProcessMaterialEmissionShape) {
	class(self).SetEmissionShape(value)
}

func (self Instance) EmissionSphereRadius() float64 {
	return float64(float64(class(self).GetEmissionSphereRadius()))
}

func (self Instance) SetEmissionSphereRadius(value float64) {
	class(self).SetEmissionSphereRadius(gd.Float(value))
}

func (self Instance) EmissionBoxExtents() gd.Vector3 {
	return gd.Vector3(class(self).GetEmissionBoxExtents())
}

func (self Instance) SetEmissionBoxExtents(value gd.Vector3) {
	class(self).SetEmissionBoxExtents(value)
}

func (self Instance) EmissionPointTexture() gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetEmissionPointTexture())
}

func (self Instance) SetEmissionPointTexture(value gdclass.Texture2D) {
	class(self).SetEmissionPointTexture(value)
}

func (self Instance) EmissionNormalTexture() gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetEmissionNormalTexture())
}

func (self Instance) SetEmissionNormalTexture(value gdclass.Texture2D) {
	class(self).SetEmissionNormalTexture(value)
}

func (self Instance) EmissionColorTexture() gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetEmissionColorTexture())
}

func (self Instance) SetEmissionColorTexture(value gdclass.Texture2D) {
	class(self).SetEmissionColorTexture(value)
}

func (self Instance) EmissionPointCount() int {
	return int(int(class(self).GetEmissionPointCount()))
}

func (self Instance) SetEmissionPointCount(value int) {
	class(self).SetEmissionPointCount(gd.Int(value))
}

func (self Instance) EmissionRingAxis() gd.Vector3 {
	return gd.Vector3(class(self).GetEmissionRingAxis())
}

func (self Instance) SetEmissionRingAxis(value gd.Vector3) {
	class(self).SetEmissionRingAxis(value)
}

func (self Instance) EmissionRingHeight() float64 {
	return float64(float64(class(self).GetEmissionRingHeight()))
}

func (self Instance) SetEmissionRingHeight(value float64) {
	class(self).SetEmissionRingHeight(gd.Float(value))
}

func (self Instance) EmissionRingRadius() float64 {
	return float64(float64(class(self).GetEmissionRingRadius()))
}

func (self Instance) SetEmissionRingRadius(value float64) {
	class(self).SetEmissionRingRadius(gd.Float(value))
}

func (self Instance) EmissionRingInnerRadius() float64 {
	return float64(float64(class(self).GetEmissionRingInnerRadius()))
}

func (self Instance) SetEmissionRingInnerRadius(value float64) {
	class(self).SetEmissionRingInnerRadius(gd.Float(value))
}

func (self Instance) Angle() gd.Vector2 {
	return gd.Vector2(class(self).GetParam(7))
}

func (self Instance) SetAngle(value gd.Vector2) {
	class(self).SetParam(7, value)
}

func (self Instance) AngleMin() float64 {
	return float64(float64(class(self).GetParamMin(7)))
}

func (self Instance) SetAngleMin(value float64) {
	class(self).SetParamMin(7, gd.Float(value))
}

func (self Instance) AngleMax() float64 {
	return float64(float64(class(self).GetParamMax(7)))
}

func (self Instance) SetAngleMax(value float64) {
	class(self).SetParamMax(7, gd.Float(value))
}

func (self Instance) AngleCurve() gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetParamTexture(7))
}

func (self Instance) SetAngleCurve(value gdclass.Texture2D) {
	class(self).SetParamTexture(7, value)
}

func (self Instance) InheritVelocityRatio() float64 {
	return float64(float64(class(self).GetInheritVelocityRatio()))
}

func (self Instance) SetInheritVelocityRatio(value float64) {
	class(self).SetInheritVelocityRatio(gd.Float(value))
}

func (self Instance) VelocityPivot() gd.Vector3 {
	return gd.Vector3(class(self).GetVelocityPivot())
}

func (self Instance) SetVelocityPivot(value gd.Vector3) {
	class(self).SetVelocityPivot(value)
}

func (self Instance) Direction() gd.Vector3 {
	return gd.Vector3(class(self).GetDirection())
}

func (self Instance) SetDirection(value gd.Vector3) {
	class(self).SetDirection(value)
}

func (self Instance) Spread() float64 {
	return float64(float64(class(self).GetSpread()))
}

func (self Instance) SetSpread(value float64) {
	class(self).SetSpread(gd.Float(value))
}

func (self Instance) Flatness() float64 {
	return float64(float64(class(self).GetFlatness()))
}

func (self Instance) SetFlatness(value float64) {
	class(self).SetFlatness(gd.Float(value))
}

func (self Instance) InitialVelocity() gd.Vector2 {
	return gd.Vector2(class(self).GetParam(0))
}

func (self Instance) SetInitialVelocity(value gd.Vector2) {
	class(self).SetParam(0, value)
}

func (self Instance) InitialVelocityMin() float64 {
	return float64(float64(class(self).GetParamMin(0)))
}

func (self Instance) SetInitialVelocityMin(value float64) {
	class(self).SetParamMin(0, gd.Float(value))
}

func (self Instance) InitialVelocityMax() float64 {
	return float64(float64(class(self).GetParamMax(0)))
}

func (self Instance) SetInitialVelocityMax(value float64) {
	class(self).SetParamMax(0, gd.Float(value))
}

func (self Instance) AngularVelocity() gd.Vector2 {
	return gd.Vector2(class(self).GetParam(1))
}

func (self Instance) SetAngularVelocity(value gd.Vector2) {
	class(self).SetParam(1, value)
}

func (self Instance) AngularVelocityMin() float64 {
	return float64(float64(class(self).GetParamMin(1)))
}

func (self Instance) SetAngularVelocityMin(value float64) {
	class(self).SetParamMin(1, gd.Float(value))
}

func (self Instance) AngularVelocityMax() float64 {
	return float64(float64(class(self).GetParamMax(1)))
}

func (self Instance) SetAngularVelocityMax(value float64) {
	class(self).SetParamMax(1, gd.Float(value))
}

func (self Instance) AngularVelocityCurve() gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetParamTexture(1))
}

func (self Instance) SetAngularVelocityCurve(value gdclass.Texture2D) {
	class(self).SetParamTexture(1, value)
}

func (self Instance) DirectionalVelocity() gd.Vector2 {
	return gd.Vector2(class(self).GetParam(16))
}

func (self Instance) SetDirectionalVelocity(value gd.Vector2) {
	class(self).SetParam(16, value)
}

func (self Instance) DirectionalVelocityMin() float64 {
	return float64(float64(class(self).GetParamMin(16)))
}

func (self Instance) SetDirectionalVelocityMin(value float64) {
	class(self).SetParamMin(16, gd.Float(value))
}

func (self Instance) DirectionalVelocityMax() float64 {
	return float64(float64(class(self).GetParamMax(16)))
}

func (self Instance) SetDirectionalVelocityMax(value float64) {
	class(self).SetParamMax(16, gd.Float(value))
}

func (self Instance) DirectionalVelocityCurve() gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetParamTexture(16))
}

func (self Instance) SetDirectionalVelocityCurve(value gdclass.Texture2D) {
	class(self).SetParamTexture(16, value)
}

func (self Instance) OrbitVelocity() gd.Vector2 {
	return gd.Vector2(class(self).GetParam(2))
}

func (self Instance) SetOrbitVelocity(value gd.Vector2) {
	class(self).SetParam(2, value)
}

func (self Instance) OrbitVelocityMin() float64 {
	return float64(float64(class(self).GetParamMin(2)))
}

func (self Instance) SetOrbitVelocityMin(value float64) {
	class(self).SetParamMin(2, gd.Float(value))
}

func (self Instance) OrbitVelocityMax() float64 {
	return float64(float64(class(self).GetParamMax(2)))
}

func (self Instance) SetOrbitVelocityMax(value float64) {
	class(self).SetParamMax(2, gd.Float(value))
}

func (self Instance) OrbitVelocityCurve() gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetParamTexture(2))
}

func (self Instance) SetOrbitVelocityCurve(value gdclass.Texture2D) {
	class(self).SetParamTexture(2, value)
}

func (self Instance) RadialVelocity() gd.Vector2 {
	return gd.Vector2(class(self).GetParam(15))
}

func (self Instance) SetRadialVelocity(value gd.Vector2) {
	class(self).SetParam(15, value)
}

func (self Instance) RadialVelocityMin() float64 {
	return float64(float64(class(self).GetParamMin(15)))
}

func (self Instance) SetRadialVelocityMin(value float64) {
	class(self).SetParamMin(15, gd.Float(value))
}

func (self Instance) RadialVelocityMax() float64 {
	return float64(float64(class(self).GetParamMax(15)))
}

func (self Instance) SetRadialVelocityMax(value float64) {
	class(self).SetParamMax(15, gd.Float(value))
}

func (self Instance) RadialVelocityCurve() gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetParamTexture(15))
}

func (self Instance) SetRadialVelocityCurve(value gdclass.Texture2D) {
	class(self).SetParamTexture(15, value)
}

func (self Instance) VelocityLimitCurve() gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetVelocityLimitCurve())
}

func (self Instance) SetVelocityLimitCurve(value gdclass.Texture2D) {
	class(self).SetVelocityLimitCurve(value)
}

func (self Instance) Gravity() gd.Vector3 {
	return gd.Vector3(class(self).GetGravity())
}

func (self Instance) SetGravity(value gd.Vector3) {
	class(self).SetGravity(value)
}

func (self Instance) LinearAccel() gd.Vector2 {
	return gd.Vector2(class(self).GetParam(3))
}

func (self Instance) SetLinearAccel(value gd.Vector2) {
	class(self).SetParam(3, value)
}

func (self Instance) LinearAccelMin() float64 {
	return float64(float64(class(self).GetParamMin(3)))
}

func (self Instance) SetLinearAccelMin(value float64) {
	class(self).SetParamMin(3, gd.Float(value))
}

func (self Instance) LinearAccelMax() float64 {
	return float64(float64(class(self).GetParamMax(3)))
}

func (self Instance) SetLinearAccelMax(value float64) {
	class(self).SetParamMax(3, gd.Float(value))
}

func (self Instance) LinearAccelCurve() gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetParamTexture(3))
}

func (self Instance) SetLinearAccelCurve(value gdclass.Texture2D) {
	class(self).SetParamTexture(3, value)
}

func (self Instance) RadialAccel() gd.Vector2 {
	return gd.Vector2(class(self).GetParam(4))
}

func (self Instance) SetRadialAccel(value gd.Vector2) {
	class(self).SetParam(4, value)
}

func (self Instance) RadialAccelMin() float64 {
	return float64(float64(class(self).GetParamMin(4)))
}

func (self Instance) SetRadialAccelMin(value float64) {
	class(self).SetParamMin(4, gd.Float(value))
}

func (self Instance) RadialAccelMax() float64 {
	return float64(float64(class(self).GetParamMax(4)))
}

func (self Instance) SetRadialAccelMax(value float64) {
	class(self).SetParamMax(4, gd.Float(value))
}

func (self Instance) RadialAccelCurve() gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetParamTexture(4))
}

func (self Instance) SetRadialAccelCurve(value gdclass.Texture2D) {
	class(self).SetParamTexture(4, value)
}

func (self Instance) TangentialAccel() gd.Vector2 {
	return gd.Vector2(class(self).GetParam(5))
}

func (self Instance) SetTangentialAccel(value gd.Vector2) {
	class(self).SetParam(5, value)
}

func (self Instance) TangentialAccelMin() float64 {
	return float64(float64(class(self).GetParamMin(5)))
}

func (self Instance) SetTangentialAccelMin(value float64) {
	class(self).SetParamMin(5, gd.Float(value))
}

func (self Instance) TangentialAccelMax() float64 {
	return float64(float64(class(self).GetParamMax(5)))
}

func (self Instance) SetTangentialAccelMax(value float64) {
	class(self).SetParamMax(5, gd.Float(value))
}

func (self Instance) TangentialAccelCurve() gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetParamTexture(5))
}

func (self Instance) SetTangentialAccelCurve(value gdclass.Texture2D) {
	class(self).SetParamTexture(5, value)
}

func (self Instance) Damping() gd.Vector2 {
	return gd.Vector2(class(self).GetParam(6))
}

func (self Instance) SetDamping(value gd.Vector2) {
	class(self).SetParam(6, value)
}

func (self Instance) DampingMin() float64 {
	return float64(float64(class(self).GetParamMin(6)))
}

func (self Instance) SetDampingMin(value float64) {
	class(self).SetParamMin(6, gd.Float(value))
}

func (self Instance) DampingMax() float64 {
	return float64(float64(class(self).GetParamMax(6)))
}

func (self Instance) SetDampingMax(value float64) {
	class(self).SetParamMax(6, gd.Float(value))
}

func (self Instance) DampingCurve() gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetParamTexture(6))
}

func (self Instance) SetDampingCurve(value gdclass.Texture2D) {
	class(self).SetParamTexture(6, value)
}

func (self Instance) AttractorInteractionEnabled() bool {
	return bool(class(self).IsAttractorInteractionEnabled())
}

func (self Instance) SetAttractorInteractionEnabled(value bool) {
	class(self).SetAttractorInteractionEnabled(value)
}

func (self Instance) Scale() gd.Vector2 {
	return gd.Vector2(class(self).GetParam(8))
}

func (self Instance) SetScale(value gd.Vector2) {
	class(self).SetParam(8, value)
}

func (self Instance) ScaleMin() float64 {
	return float64(float64(class(self).GetParamMin(8)))
}

func (self Instance) SetScaleMin(value float64) {
	class(self).SetParamMin(8, gd.Float(value))
}

func (self Instance) ScaleMax() float64 {
	return float64(float64(class(self).GetParamMax(8)))
}

func (self Instance) SetScaleMax(value float64) {
	class(self).SetParamMax(8, gd.Float(value))
}

func (self Instance) ScaleCurve() gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetParamTexture(8))
}

func (self Instance) SetScaleCurve(value gdclass.Texture2D) {
	class(self).SetParamTexture(8, value)
}

func (self Instance) ScaleOverVelocity() gd.Vector2 {
	return gd.Vector2(class(self).GetParam(17))
}

func (self Instance) SetScaleOverVelocity(value gd.Vector2) {
	class(self).SetParam(17, value)
}

func (self Instance) ScaleOverVelocityMin() float64 {
	return float64(float64(class(self).GetParamMin(17)))
}

func (self Instance) SetScaleOverVelocityMin(value float64) {
	class(self).SetParamMin(17, gd.Float(value))
}

func (self Instance) ScaleOverVelocityMax() float64 {
	return float64(float64(class(self).GetParamMax(17)))
}

func (self Instance) SetScaleOverVelocityMax(value float64) {
	class(self).SetParamMax(17, gd.Float(value))
}

func (self Instance) ScaleOverVelocityCurve() gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetParamTexture(17))
}

func (self Instance) SetScaleOverVelocityCurve(value gdclass.Texture2D) {
	class(self).SetParamTexture(17, value)
}

func (self Instance) Color() gd.Color {
	return gd.Color(class(self).GetColor())
}

func (self Instance) SetColor(value gd.Color) {
	class(self).SetColor(value)
}

func (self Instance) ColorRamp() gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetColorRamp())
}

func (self Instance) SetColorRamp(value gdclass.Texture2D) {
	class(self).SetColorRamp(value)
}

func (self Instance) ColorInitialRamp() gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetColorInitialRamp())
}

func (self Instance) SetColorInitialRamp(value gdclass.Texture2D) {
	class(self).SetColorInitialRamp(value)
}

func (self Instance) AlphaCurve() gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetAlphaCurve())
}

func (self Instance) SetAlphaCurve(value gdclass.Texture2D) {
	class(self).SetAlphaCurve(value)
}

func (self Instance) EmissionCurve() gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetEmissionCurve())
}

func (self Instance) SetEmissionCurve(value gdclass.Texture2D) {
	class(self).SetEmissionCurve(value)
}

func (self Instance) HueVariation() gd.Vector2 {
	return gd.Vector2(class(self).GetParam(9))
}

func (self Instance) SetHueVariation(value gd.Vector2) {
	class(self).SetParam(9, value)
}

func (self Instance) HueVariationMin() float64 {
	return float64(float64(class(self).GetParamMin(9)))
}

func (self Instance) SetHueVariationMin(value float64) {
	class(self).SetParamMin(9, gd.Float(value))
}

func (self Instance) HueVariationMax() float64 {
	return float64(float64(class(self).GetParamMax(9)))
}

func (self Instance) SetHueVariationMax(value float64) {
	class(self).SetParamMax(9, gd.Float(value))
}

func (self Instance) HueVariationCurve() gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetParamTexture(9))
}

func (self Instance) SetHueVariationCurve(value gdclass.Texture2D) {
	class(self).SetParamTexture(9, value)
}

func (self Instance) AnimSpeed() gd.Vector2 {
	return gd.Vector2(class(self).GetParam(10))
}

func (self Instance) SetAnimSpeed(value gd.Vector2) {
	class(self).SetParam(10, value)
}

func (self Instance) AnimSpeedMin() float64 {
	return float64(float64(class(self).GetParamMin(10)))
}

func (self Instance) SetAnimSpeedMin(value float64) {
	class(self).SetParamMin(10, gd.Float(value))
}

func (self Instance) AnimSpeedMax() float64 {
	return float64(float64(class(self).GetParamMax(10)))
}

func (self Instance) SetAnimSpeedMax(value float64) {
	class(self).SetParamMax(10, gd.Float(value))
}

func (self Instance) AnimSpeedCurve() gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetParamTexture(10))
}

func (self Instance) SetAnimSpeedCurve(value gdclass.Texture2D) {
	class(self).SetParamTexture(10, value)
}

func (self Instance) AnimOffset() gd.Vector2 {
	return gd.Vector2(class(self).GetParam(11))
}

func (self Instance) SetAnimOffset(value gd.Vector2) {
	class(self).SetParam(11, value)
}

func (self Instance) AnimOffsetMin() float64 {
	return float64(float64(class(self).GetParamMin(11)))
}

func (self Instance) SetAnimOffsetMin(value float64) {
	class(self).SetParamMin(11, gd.Float(value))
}

func (self Instance) AnimOffsetMax() float64 {
	return float64(float64(class(self).GetParamMax(11)))
}

func (self Instance) SetAnimOffsetMax(value float64) {
	class(self).SetParamMax(11, gd.Float(value))
}

func (self Instance) AnimOffsetCurve() gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetParamTexture(11))
}

func (self Instance) SetAnimOffsetCurve(value gdclass.Texture2D) {
	class(self).SetParamTexture(11, value)
}

func (self Instance) TurbulenceEnabled() bool {
	return bool(class(self).GetTurbulenceEnabled())
}

func (self Instance) SetTurbulenceEnabled(value bool) {
	class(self).SetTurbulenceEnabled(value)
}

func (self Instance) TurbulenceNoiseStrength() float64 {
	return float64(float64(class(self).GetTurbulenceNoiseStrength()))
}

func (self Instance) SetTurbulenceNoiseStrength(value float64) {
	class(self).SetTurbulenceNoiseStrength(gd.Float(value))
}

func (self Instance) TurbulenceNoiseScale() float64 {
	return float64(float64(class(self).GetTurbulenceNoiseScale()))
}

func (self Instance) SetTurbulenceNoiseScale(value float64) {
	class(self).SetTurbulenceNoiseScale(gd.Float(value))
}

func (self Instance) TurbulenceNoiseSpeed() gd.Vector3 {
	return gd.Vector3(class(self).GetTurbulenceNoiseSpeed())
}

func (self Instance) SetTurbulenceNoiseSpeed(value gd.Vector3) {
	class(self).SetTurbulenceNoiseSpeed(value)
}

func (self Instance) TurbulenceNoiseSpeedRandom() float64 {
	return float64(float64(class(self).GetTurbulenceNoiseSpeedRandom()))
}

func (self Instance) SetTurbulenceNoiseSpeedRandom(value float64) {
	class(self).SetTurbulenceNoiseSpeedRandom(gd.Float(value))
}

func (self Instance) TurbulenceInfluence() gd.Vector2 {
	return gd.Vector2(class(self).GetParam(13))
}

func (self Instance) SetTurbulenceInfluence(value gd.Vector2) {
	class(self).SetParam(13, value)
}

func (self Instance) TurbulenceInfluenceMin() float64 {
	return float64(float64(class(self).GetParamMin(13)))
}

func (self Instance) SetTurbulenceInfluenceMin(value float64) {
	class(self).SetParamMin(13, gd.Float(value))
}

func (self Instance) TurbulenceInfluenceMax() float64 {
	return float64(float64(class(self).GetParamMax(13)))
}

func (self Instance) SetTurbulenceInfluenceMax(value float64) {
	class(self).SetParamMax(13, gd.Float(value))
}

func (self Instance) TurbulenceInitialDisplacement() gd.Vector2 {
	return gd.Vector2(class(self).GetParam(14))
}

func (self Instance) SetTurbulenceInitialDisplacement(value gd.Vector2) {
	class(self).SetParam(14, value)
}

func (self Instance) TurbulenceInitialDisplacementMin() float64 {
	return float64(float64(class(self).GetParamMin(14)))
}

func (self Instance) SetTurbulenceInitialDisplacementMin(value float64) {
	class(self).SetParamMin(14, gd.Float(value))
}

func (self Instance) TurbulenceInitialDisplacementMax() float64 {
	return float64(float64(class(self).GetParamMax(14)))
}

func (self Instance) SetTurbulenceInitialDisplacementMax(value float64) {
	class(self).SetParamMax(14, gd.Float(value))
}

func (self Instance) TurbulenceInfluenceOverLife() gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetParamTexture(12))
}

func (self Instance) SetTurbulenceInfluenceOverLife(value gdclass.Texture2D) {
	class(self).SetParamTexture(12, value)
}

func (self Instance) CollisionMode() classdb.ParticleProcessMaterialCollisionMode {
	return classdb.ParticleProcessMaterialCollisionMode(class(self).GetCollisionMode())
}

func (self Instance) SetCollisionMode(value classdb.ParticleProcessMaterialCollisionMode) {
	class(self).SetCollisionMode(value)
}

func (self Instance) CollisionFriction() float64 {
	return float64(float64(class(self).GetCollisionFriction()))
}

func (self Instance) SetCollisionFriction(value float64) {
	class(self).SetCollisionFriction(gd.Float(value))
}

func (self Instance) CollisionBounce() float64 {
	return float64(float64(class(self).GetCollisionBounce()))
}

func (self Instance) SetCollisionBounce(value float64) {
	class(self).SetCollisionBounce(gd.Float(value))
}

func (self Instance) CollisionUseScale() bool {
	return bool(class(self).IsCollisionUsingScale())
}

func (self Instance) SetCollisionUseScale(value bool) {
	class(self).SetCollisionUseScale(value)
}

func (self Instance) SubEmitterMode() classdb.ParticleProcessMaterialSubEmitterMode {
	return classdb.ParticleProcessMaterialSubEmitterMode(class(self).GetSubEmitterMode())
}

func (self Instance) SetSubEmitterMode(value classdb.ParticleProcessMaterialSubEmitterMode) {
	class(self).SetSubEmitterMode(value)
}

func (self Instance) SubEmitterFrequency() float64 {
	return float64(float64(class(self).GetSubEmitterFrequency()))
}

func (self Instance) SetSubEmitterFrequency(value float64) {
	class(self).SetSubEmitterFrequency(gd.Float(value))
}

func (self Instance) SubEmitterAmountAtEnd() int {
	return int(int(class(self).GetSubEmitterAmountAtEnd()))
}

func (self Instance) SetSubEmitterAmountAtEnd(value int) {
	class(self).SetSubEmitterAmountAtEnd(gd.Int(value))
}

func (self Instance) SubEmitterAmountAtCollision() int {
	return int(int(class(self).GetSubEmitterAmountAtCollision()))
}

func (self Instance) SetSubEmitterAmountAtCollision(value int) {
	class(self).SetSubEmitterAmountAtCollision(gd.Int(value))
}

func (self Instance) SubEmitterKeepVelocity() bool {
	return bool(class(self).GetSubEmitterKeepVelocity())
}

func (self Instance) SetSubEmitterKeepVelocity(value bool) {
	class(self).SetSubEmitterKeepVelocity(value)
}

//go:nosplit
func (self class) SetDirection(degrees gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, degrees)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDirection() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInheritVelocityRatio(ratio gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_inherit_velocity_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetInheritVelocityRatio() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_inherit_velocity_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSpread(degrees gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, degrees)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_spread, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSpread() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_spread, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFlatness(amount gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_flatness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFlatness() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_flatness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the minimum and maximum values of the given [param param].
The [code]x[/code] component of the argument vector corresponds to minimum and the [code]y[/code] component corresponds to maximum.
*/
//go:nosplit
func (self class) SetParam(param classdb.ParticleProcessMaterialParameter, value gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the minimum and maximum values of the given [param param] as a vector.
The [code]x[/code] component of the returned vector corresponds to minimum and the [code]y[/code] component corresponds to maximum.
*/
//go:nosplit
func (self class) GetParam(param classdb.ParticleProcessMaterialParameter) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_param, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the minimum value range for the given parameter.
*/
//go:nosplit
func (self class) SetParamMin(param classdb.ParticleProcessMaterialParameter, value gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_param_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the minimum value range for the given parameter.
*/
//go:nosplit
func (self class) GetParamMin(param classdb.ParticleProcessMaterialParameter) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_param_min, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the maximum value range for the given parameter.
*/
//go:nosplit
func (self class) SetParamMax(param classdb.ParticleProcessMaterialParameter, value gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_param_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the maximum value range for the given parameter.
*/
//go:nosplit
func (self class) GetParamMax(param classdb.ParticleProcessMaterialParameter) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_param_max, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [Texture2D] for the specified [enum Parameter].
*/
//go:nosplit
func (self class) SetParamTexture(param classdb.ParticleProcessMaterialParameter, texture gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_param_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [Texture2D] used by the specified parameter.
*/
//go:nosplit
func (self class) GetParamTexture(param classdb.ParticleProcessMaterialParameter) gdclass.Texture2D {
	var frame = callframe.New()
	callframe.Arg(frame, param)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_param_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetColor(color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetColorRamp(ramp gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(ramp[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_color_ramp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetColorRamp() gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_color_ramp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAlphaCurve(curve gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(curve[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_alpha_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAlphaCurve() gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_alpha_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionCurve(curve gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(curve[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_emission_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionCurve() gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_emission_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetColorInitialRamp(ramp gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(ramp[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_color_initial_ramp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetColorInitialRamp() gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_color_initial_ramp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVelocityLimitCurve(curve gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(curve[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_velocity_limit_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVelocityLimitCurve() gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_velocity_limit_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
If [code]true[/code], enables the specified particle flag. See [enum ParticleFlags] for options.
*/
//go:nosplit
func (self class) SetParticleFlag(particle_flag classdb.ParticleProcessMaterialParticleFlags, enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, particle_flag)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_particle_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the specified particle flag is enabled. See [enum ParticleFlags] for options.
*/
//go:nosplit
func (self class) GetParticleFlag(particle_flag classdb.ParticleProcessMaterialParticleFlags) bool {
	var frame = callframe.New()
	callframe.Arg(frame, particle_flag)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_particle_flag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVelocityPivot(pivot gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, pivot)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_velocity_pivot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVelocityPivot() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_velocity_pivot, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionShape(shape classdb.ParticleProcessMaterialEmissionShape) {
	var frame = callframe.New()
	callframe.Arg(frame, shape)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_emission_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionShape() classdb.ParticleProcessMaterialEmissionShape {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ParticleProcessMaterialEmissionShape](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_emission_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionSphereRadius(radius gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_emission_sphere_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionSphereRadius() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_emission_sphere_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionBoxExtents(extents gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, extents)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_emission_box_extents, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionBoxExtents() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_emission_box_extents, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionPointTexture(texture gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_emission_point_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionPointTexture() gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_emission_point_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionNormalTexture(texture gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_emission_normal_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionNormalTexture() gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_emission_normal_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionColorTexture(texture gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_emission_color_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionColorTexture() gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_emission_color_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionPointCount(point_count gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, point_count)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_emission_point_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionPointCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_emission_point_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionRingAxis(axis gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, axis)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_emission_ring_axis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionRingAxis() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_emission_ring_axis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionRingHeight(height gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_emission_ring_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionRingHeight() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_emission_ring_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionRingRadius(radius gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_emission_ring_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionRingRadius() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_emission_ring_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionRingInnerRadius(inner_radius gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, inner_radius)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_emission_ring_inner_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionRingInnerRadius() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_emission_ring_inner_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionShapeOffset(emission_shape_offset gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, emission_shape_offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_emission_shape_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionShapeOffset() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_emission_shape_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionShapeScale(emission_shape_scale gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, emission_shape_scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_emission_shape_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionShapeScale() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_emission_shape_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetTurbulenceEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_turbulence_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTurbulenceEnabled(turbulence_enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, turbulence_enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_turbulence_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTurbulenceNoiseStrength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_turbulence_noise_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTurbulenceNoiseStrength(turbulence_noise_strength gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, turbulence_noise_strength)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_turbulence_noise_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTurbulenceNoiseScale() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_turbulence_noise_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTurbulenceNoiseScale(turbulence_noise_scale gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, turbulence_noise_scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_turbulence_noise_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTurbulenceNoiseSpeedRandom() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_turbulence_noise_speed_random, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTurbulenceNoiseSpeedRandom(turbulence_noise_speed_random gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, turbulence_noise_speed_random)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_turbulence_noise_speed_random, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTurbulenceNoiseSpeed() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_turbulence_noise_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTurbulenceNoiseSpeed(turbulence_noise_speed gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, turbulence_noise_speed)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_turbulence_noise_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGravity() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGravity(accel_vec gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, accel_vec)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_gravity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetLifetimeRandomness(randomness gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, randomness)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_lifetime_randomness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLifetimeRandomness() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_lifetime_randomness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetSubEmitterMode() classdb.ParticleProcessMaterialSubEmitterMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ParticleProcessMaterialSubEmitterMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_sub_emitter_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSubEmitterMode(mode classdb.ParticleProcessMaterialSubEmitterMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_sub_emitter_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSubEmitterFrequency() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_sub_emitter_frequency, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSubEmitterFrequency(hz gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, hz)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_sub_emitter_frequency, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSubEmitterAmountAtEnd() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_sub_emitter_amount_at_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSubEmitterAmountAtEnd(amount gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_sub_emitter_amount_at_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSubEmitterAmountAtCollision() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_sub_emitter_amount_at_collision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSubEmitterAmountAtCollision(amount gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_sub_emitter_amount_at_collision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSubEmitterKeepVelocity() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_sub_emitter_keep_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSubEmitterKeepVelocity(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_sub_emitter_keep_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetAttractorInteractionEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_attractor_interaction_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsAttractorInteractionEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_is_attractor_interaction_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollisionMode(mode classdb.ParticleProcessMaterialCollisionMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_collision_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCollisionMode() classdb.ParticleProcessMaterialCollisionMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ParticleProcessMaterialCollisionMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_collision_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollisionUseScale(radius bool) {
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_collision_use_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsCollisionUsingScale() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_is_collision_using_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollisionFriction(friction gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, friction)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_collision_friction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCollisionFriction() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_collision_friction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollisionBounce(bounce gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, bounce)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_set_collision_bounce, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCollisionBounce() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ParticleProcessMaterial.Bind_get_collision_bounce, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsParticleProcessMaterial() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsParticleProcessMaterial() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsMaterial() Material.Advanced {
	return *((*Material.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsMaterial() Material.Instance {
	return *((*Material.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsMaterial(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsMaterial(), name)
	}
}
func init() {
	classdb.Register("ParticleProcessMaterial", func(ptr gd.Object) any { return classdb.ParticleProcessMaterial(ptr) })
}

type Parameter = classdb.ParticleProcessMaterialParameter

const (
	/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set initial velocity properties.*/
	ParamInitialLinearVelocity Parameter = 0
	/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set angular velocity properties.*/
	ParamAngularVelocity Parameter = 1
	/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set orbital velocity properties.*/
	ParamOrbitVelocity Parameter = 2
	/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set linear acceleration properties.*/
	ParamLinearAccel Parameter = 3
	/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set radial acceleration properties.*/
	ParamRadialAccel Parameter = 4
	/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set tangential acceleration properties.*/
	ParamTangentialAccel Parameter = 5
	/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set damping properties.*/
	ParamDamping Parameter = 6
	/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set angle properties.*/
	ParamAngle Parameter = 7
	/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set scale properties.*/
	ParamScale Parameter = 8
	/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set hue variation properties.*/
	ParamHueVariation Parameter = 9
	/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set animation speed properties.*/
	ParamAnimSpeed Parameter = 10
	/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set animation offset properties.*/
	ParamAnimOffset Parameter = 11
	/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set radial velocity properties.*/
	ParamRadialVelocity Parameter = 15
	/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set directional velocity properties.*/
	ParamDirectionalVelocity Parameter = 16
	/*Use with [method set_param_min], [method set_param_max], and [method set_param_texture] to set scale over velocity properties.*/
	ParamScaleOverVelocity Parameter = 17
	/*Represents the size of the [enum Parameter] enum.*/
	ParamMax Parameter = 18
	/*Use with [method set_param_min] and [method set_param_max] to set the turbulence minimum und maximum influence on each particles velocity.*/
	ParamTurbVelInfluence Parameter = 13
	/*Use with [method set_param_min] and [method set_param_max] to set the turbulence minimum and maximum displacement of the particles spawn position.*/
	ParamTurbInitDisplacement Parameter = 14
	/*Use with [method set_param_texture] to set the turbulence influence over the particles life time.*/
	ParamTurbInfluenceOverLife Parameter = 12
)

type ParticleFlags = classdb.ParticleProcessMaterialParticleFlags

const (
	/*Use with [method set_particle_flag] to set [member particle_flag_align_y].*/
	ParticleFlagAlignYToVelocity ParticleFlags = 0
	/*Use with [method set_particle_flag] to set [member particle_flag_rotate_y].*/
	ParticleFlagRotateY ParticleFlags = 1
	/*Use with [method set_particle_flag] to set [member particle_flag_disable_z].*/
	ParticleFlagDisableZ          ParticleFlags = 2
	ParticleFlagDampingAsFriction ParticleFlags = 3
	/*Represents the size of the [enum ParticleFlags] enum.*/
	ParticleFlagMax ParticleFlags = 4
)

type EmissionShape = classdb.ParticleProcessMaterialEmissionShape

const (
	/*All particles will be emitted from a single point.*/
	EmissionShapePoint EmissionShape = 0
	/*Particles will be emitted in the volume of a sphere.*/
	EmissionShapeSphere EmissionShape = 1
	/*Particles will be emitted on the surface of a sphere.*/
	EmissionShapeSphereSurface EmissionShape = 2
	/*Particles will be emitted in the volume of a box.*/
	EmissionShapeBox EmissionShape = 3
	/*Particles will be emitted at a position determined by sampling a random point on the [member emission_point_texture]. Particle color will be modulated by [member emission_color_texture].*/
	EmissionShapePoints EmissionShape = 4
	/*Particles will be emitted at a position determined by sampling a random point on the [member emission_point_texture]. Particle velocity and rotation will be set based on [member emission_normal_texture]. Particle color will be modulated by [member emission_color_texture].*/
	EmissionShapeDirectedPoints EmissionShape = 5
	/*Particles will be emitted in a ring or cylinder.*/
	EmissionShapeRing EmissionShape = 6
	/*Represents the size of the [enum EmissionShape] enum.*/
	EmissionShapeMax EmissionShape = 7
)

type SubEmitterMode = classdb.ParticleProcessMaterialSubEmitterMode

const (
	SubEmitterDisabled    SubEmitterMode = 0
	SubEmitterConstant    SubEmitterMode = 1
	SubEmitterAtEnd       SubEmitterMode = 2
	SubEmitterAtCollision SubEmitterMode = 3
	/*Represents the size of the [enum SubEmitterMode] enum.*/
	SubEmitterMax SubEmitterMode = 4
)

type CollisionMode = classdb.ParticleProcessMaterialCollisionMode

const (
	/*No collision for particles. Particles will go through [GPUParticlesCollision3D] nodes.*/
	CollisionDisabled CollisionMode = 0
	/*[RigidBody3D]-style collision for particles using [GPUParticlesCollision3D] nodes.*/
	CollisionRigid CollisionMode = 1
	/*Hide particles instantly when colliding with a [GPUParticlesCollision3D] node. This can be combined with a subemitter that uses the [constant COLLISION_RIGID] collision mode to "replace" the parent particle with the subemitter on impact.*/
	CollisionHideOnContact CollisionMode = 2
	/*Represents the size of the [enum CollisionMode] enum.*/
	CollisionMax CollisionMode = 3
)
