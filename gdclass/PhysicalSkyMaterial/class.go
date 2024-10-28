package PhysicalSkyMaterial

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Material"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
The [PhysicalSkyMaterial] uses the Preetham analytic daylight model to draw a sky based on physical properties. This results in a substantially more realistic sky than the [ProceduralSkyMaterial], but it is slightly slower and less flexible.
The [PhysicalSkyMaterial] only supports one sun. The color, energy, and direction of the sun are taken from the first [DirectionalLight3D] in the scene tree.

*/
type Go [1]classdb.PhysicalSkyMaterial
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.PhysicalSkyMaterial
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicalSkyMaterial"))
	return Go{classdb.PhysicalSkyMaterial(object)}
}

func (self Go) RayleighCoefficient() float64 {
		return float64(float64(class(self).GetRayleighCoefficient()))
}

func (self Go) SetRayleighCoefficient(value float64) {
	class(self).SetRayleighCoefficient(gd.Float(value))
}

func (self Go) RayleighColor() gd.Color {
		return gd.Color(class(self).GetRayleighColor())
}

func (self Go) SetRayleighColor(value gd.Color) {
	class(self).SetRayleighColor(value)
}

func (self Go) MieCoefficient() float64 {
		return float64(float64(class(self).GetMieCoefficient()))
}

func (self Go) SetMieCoefficient(value float64) {
	class(self).SetMieCoefficient(gd.Float(value))
}

func (self Go) MieEccentricity() float64 {
		return float64(float64(class(self).GetMieEccentricity()))
}

func (self Go) SetMieEccentricity(value float64) {
	class(self).SetMieEccentricity(gd.Float(value))
}

func (self Go) MieColor() gd.Color {
		return gd.Color(class(self).GetMieColor())
}

func (self Go) SetMieColor(value gd.Color) {
	class(self).SetMieColor(value)
}

func (self Go) Turbidity() float64 {
		return float64(float64(class(self).GetTurbidity()))
}

func (self Go) SetTurbidity(value float64) {
	class(self).SetTurbidity(gd.Float(value))
}

func (self Go) SunDiskScale() float64 {
		return float64(float64(class(self).GetSunDiskScale()))
}

func (self Go) SetSunDiskScale(value float64) {
	class(self).SetSunDiskScale(gd.Float(value))
}

func (self Go) GroundColor() gd.Color {
		return gd.Color(class(self).GetGroundColor())
}

func (self Go) SetGroundColor(value gd.Color) {
	class(self).SetGroundColor(value)
}

func (self Go) EnergyMultiplier() float64 {
		return float64(float64(class(self).GetEnergyMultiplier()))
}

func (self Go) SetEnergyMultiplier(value float64) {
	class(self).SetEnergyMultiplier(gd.Float(value))
}

func (self Go) UseDebanding() bool {
		return bool(class(self).GetUseDebanding())
}

func (self Go) SetUseDebanding(value bool) {
	class(self).SetUseDebanding(value)
}

func (self Go) NightSky() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetNightSky())
}

func (self Go) SetNightSky(value gdclass.Texture2D) {
	class(self).SetNightSky(value)
}

//go:nosplit
func (self class) SetRayleighCoefficient(rayleigh gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, rayleigh)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_set_rayleigh_coefficient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRayleighCoefficient() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_get_rayleigh_coefficient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRayleighColor(color gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_set_rayleigh_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRayleighColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_get_rayleigh_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMieCoefficient(mie gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, mie)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_set_mie_coefficient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMieCoefficient() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_get_mie_coefficient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMieEccentricity(eccentricity gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, eccentricity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_set_mie_eccentricity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMieEccentricity() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_get_mie_eccentricity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMieColor(color gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_set_mie_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMieColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_get_mie_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTurbidity(turbidity gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, turbidity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_set_turbidity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTurbidity() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_get_turbidity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSunDiskScale(scale gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_set_sun_disk_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSunDiskScale() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_get_sun_disk_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGroundColor(color gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_set_ground_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGroundColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_get_ground_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnergyMultiplier(multiplier gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, multiplier)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_set_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnergyMultiplier() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_get_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseDebanding(use_debanding bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, use_debanding)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_set_use_debanding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUseDebanding() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_get_use_debanding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetNightSky(night_sky gdclass.Texture2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(night_sky[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_set_night_sky, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetNightSky() gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_get_night_sky, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self class) AsPhysicalSkyMaterial() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsPhysicalSkyMaterial() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsMaterial() Material.GD { return *((*Material.GD)(unsafe.Pointer(&self))) }
func (self Go) AsMaterial() Material.Go { return *((*Material.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsMaterial(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsMaterial(), name)
	}
}
func init() {classdb.Register("PhysicalSkyMaterial", func(ptr gd.Object) any { return classdb.PhysicalSkyMaterial(ptr) })}
