// Package PhysicalSkyMaterial provides methods for working with PhysicalSkyMaterial object instances.
package PhysicalSkyMaterial

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
import "graphics.gd/classdb/Material"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Color"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
The [PhysicalSkyMaterial] uses the Preetham analytic daylight model to draw a sky based on physical properties. This results in a substantially more realistic sky than the [ProceduralSkyMaterial], but it is slightly slower and less flexible.
The [PhysicalSkyMaterial] only supports one sun. The color, energy, and direction of the sun are taken from the first [DirectionalLight3D] in the scene tree.
*/
type Instance [1]gdclass.PhysicalSkyMaterial

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPhysicalSkyMaterial() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PhysicalSkyMaterial

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicalSkyMaterial"))
	casted := Instance{*(*gdclass.PhysicalSkyMaterial)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) RayleighCoefficient() Float.X {
	return Float.X(Float.X(class(self).GetRayleighCoefficient()))
}

func (self Instance) SetRayleighCoefficient(value Float.X) {
	class(self).SetRayleighCoefficient(gd.Float(value))
}

func (self Instance) RayleighColor() Color.RGBA {
	return Color.RGBA(class(self).GetRayleighColor())
}

func (self Instance) SetRayleighColor(value Color.RGBA) {
	class(self).SetRayleighColor(gd.Color(value))
}

func (self Instance) MieCoefficient() Float.X {
	return Float.X(Float.X(class(self).GetMieCoefficient()))
}

func (self Instance) SetMieCoefficient(value Float.X) {
	class(self).SetMieCoefficient(gd.Float(value))
}

func (self Instance) MieEccentricity() Float.X {
	return Float.X(Float.X(class(self).GetMieEccentricity()))
}

func (self Instance) SetMieEccentricity(value Float.X) {
	class(self).SetMieEccentricity(gd.Float(value))
}

func (self Instance) MieColor() Color.RGBA {
	return Color.RGBA(class(self).GetMieColor())
}

func (self Instance) SetMieColor(value Color.RGBA) {
	class(self).SetMieColor(gd.Color(value))
}

func (self Instance) Turbidity() Float.X {
	return Float.X(Float.X(class(self).GetTurbidity()))
}

func (self Instance) SetTurbidity(value Float.X) {
	class(self).SetTurbidity(gd.Float(value))
}

func (self Instance) SunDiskScale() Float.X {
	return Float.X(Float.X(class(self).GetSunDiskScale()))
}

func (self Instance) SetSunDiskScale(value Float.X) {
	class(self).SetSunDiskScale(gd.Float(value))
}

func (self Instance) GroundColor() Color.RGBA {
	return Color.RGBA(class(self).GetGroundColor())
}

func (self Instance) SetGroundColor(value Color.RGBA) {
	class(self).SetGroundColor(gd.Color(value))
}

func (self Instance) EnergyMultiplier() Float.X {
	return Float.X(Float.X(class(self).GetEnergyMultiplier()))
}

func (self Instance) SetEnergyMultiplier(value Float.X) {
	class(self).SetEnergyMultiplier(gd.Float(value))
}

func (self Instance) UseDebanding() bool {
	return bool(class(self).GetUseDebanding())
}

func (self Instance) SetUseDebanding(value bool) {
	class(self).SetUseDebanding(value)
}

func (self Instance) NightSky() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetNightSky())
}

func (self Instance) SetNightSky(value [1]gdclass.Texture2D) {
	class(self).SetNightSky(value)
}

//go:nosplit
func (self class) SetRayleighCoefficient(rayleigh gd.Float) { //gd:PhysicalSkyMaterial.set_rayleigh_coefficient
	var frame = callframe.New()
	callframe.Arg(frame, rayleigh)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_set_rayleigh_coefficient, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRayleighCoefficient() gd.Float { //gd:PhysicalSkyMaterial.get_rayleigh_coefficient
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_get_rayleigh_coefficient, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRayleighColor(color gd.Color) { //gd:PhysicalSkyMaterial.set_rayleigh_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_set_rayleigh_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRayleighColor() gd.Color { //gd:PhysicalSkyMaterial.get_rayleigh_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_get_rayleigh_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMieCoefficient(mie gd.Float) { //gd:PhysicalSkyMaterial.set_mie_coefficient
	var frame = callframe.New()
	callframe.Arg(frame, mie)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_set_mie_coefficient, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMieCoefficient() gd.Float { //gd:PhysicalSkyMaterial.get_mie_coefficient
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_get_mie_coefficient, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMieEccentricity(eccentricity gd.Float) { //gd:PhysicalSkyMaterial.set_mie_eccentricity
	var frame = callframe.New()
	callframe.Arg(frame, eccentricity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_set_mie_eccentricity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMieEccentricity() gd.Float { //gd:PhysicalSkyMaterial.get_mie_eccentricity
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_get_mie_eccentricity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMieColor(color gd.Color) { //gd:PhysicalSkyMaterial.set_mie_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_set_mie_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMieColor() gd.Color { //gd:PhysicalSkyMaterial.get_mie_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_get_mie_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTurbidity(turbidity gd.Float) { //gd:PhysicalSkyMaterial.set_turbidity
	var frame = callframe.New()
	callframe.Arg(frame, turbidity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_set_turbidity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTurbidity() gd.Float { //gd:PhysicalSkyMaterial.get_turbidity
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_get_turbidity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSunDiskScale(scale gd.Float) { //gd:PhysicalSkyMaterial.set_sun_disk_scale
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_set_sun_disk_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSunDiskScale() gd.Float { //gd:PhysicalSkyMaterial.get_sun_disk_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_get_sun_disk_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGroundColor(color gd.Color) { //gd:PhysicalSkyMaterial.set_ground_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_set_ground_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGroundColor() gd.Color { //gd:PhysicalSkyMaterial.get_ground_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_get_ground_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnergyMultiplier(multiplier gd.Float) { //gd:PhysicalSkyMaterial.set_energy_multiplier
	var frame = callframe.New()
	callframe.Arg(frame, multiplier)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_set_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnergyMultiplier() gd.Float { //gd:PhysicalSkyMaterial.get_energy_multiplier
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_get_energy_multiplier, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseDebanding(use_debanding bool) { //gd:PhysicalSkyMaterial.set_use_debanding
	var frame = callframe.New()
	callframe.Arg(frame, use_debanding)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_set_use_debanding, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUseDebanding() bool { //gd:PhysicalSkyMaterial.get_use_debanding
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_get_use_debanding, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNightSky(night_sky [1]gdclass.Texture2D) { //gd:PhysicalSkyMaterial.set_night_sky
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(night_sky[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_set_night_sky, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetNightSky() [1]gdclass.Texture2D { //gd:PhysicalSkyMaterial.get_night_sky
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalSkyMaterial.Bind_get_night_sky, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsPhysicalSkyMaterial() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPhysicalSkyMaterial() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Material.Advanced(self.AsMaterial()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Material.Instance(self.AsMaterial()), name)
	}
}
func init() {
	gdclass.Register("PhysicalSkyMaterial", func(ptr gd.Object) any {
		return [1]gdclass.PhysicalSkyMaterial{*(*gdclass.PhysicalSkyMaterial)(unsafe.Pointer(&ptr))}
	})
}
