package NoiseTexture2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Texture2D"
import "grow.graphics/gd/gdclass/Texture"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Uses the [FastNoiseLite] library or other noise generators to fill the texture data of your desired size. [NoiseTexture2D] can also generate normal map textures.
The class uses [Thread]s to generate the texture data internally, so [method Texture2D.get_image] may return [code]null[/code] if the generation process has not completed yet. In that case, you need to wait for the texture to be generated before accessing the image and the generated byte data:
[codeblock]
var texture = NoiseTexture2D.new()
texture.noise = FastNoiseLite.new()
await texture.changed
var image = texture.get_image()
var data = image.get_data()
[/codeblock]

*/
type Go [1]classdb.NoiseTexture2D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.NoiseTexture2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("NoiseTexture2D"))
	return Go{classdb.NoiseTexture2D(object)}
}

func (self Go) Invert() bool {
		return bool(class(self).GetInvert())
}

func (self Go) SetInvert(value bool) {
	class(self).SetInvert(value)
}

func (self Go) In3dSpace() bool {
		return bool(class(self).IsIn3dSpace())
}

func (self Go) SetIn3dSpace(value bool) {
	class(self).SetIn3dSpace(value)
}

func (self Go) GenerateMipmaps() bool {
		return bool(class(self).IsGeneratingMipmaps())
}

func (self Go) SetGenerateMipmaps(value bool) {
	class(self).SetGenerateMipmaps(value)
}

func (self Go) Seamless() bool {
		return bool(class(self).GetSeamless())
}

func (self Go) SetSeamless(value bool) {
	class(self).SetSeamless(value)
}

func (self Go) SeamlessBlendSkirt() float64 {
		return float64(float64(class(self).GetSeamlessBlendSkirt()))
}

func (self Go) SetSeamlessBlendSkirt(value float64) {
	class(self).SetSeamlessBlendSkirt(gd.Float(value))
}

func (self Go) AsNormalMap() bool {
		return bool(class(self).IsNormalMap())
}

func (self Go) SetAsNormalMap(value bool) {
	class(self).SetAsNormalMap(value)
}

func (self Go) BumpStrength() float64 {
		return float64(float64(class(self).GetBumpStrength()))
}

func (self Go) SetBumpStrength(value float64) {
	class(self).SetBumpStrength(gd.Float(value))
}

func (self Go) Normalize() bool {
		return bool(class(self).IsNormalized())
}

func (self Go) SetNormalize(value bool) {
	class(self).SetNormalize(value)
}

func (self Go) ColorRamp() gdclass.Gradient {
		return gdclass.Gradient(class(self).GetColorRamp())
}

func (self Go) SetColorRamp(value gdclass.Gradient) {
	class(self).SetColorRamp(value)
}

func (self Go) Noise() gdclass.Noise {
		return gdclass.Noise(class(self).GetNoise())
}

func (self Go) SetNoise(value gdclass.Noise) {
	class(self).SetNoise(value)
}

//go:nosplit
func (self class) SetWidth(width gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_set_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetHeight(height gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_set_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetInvert(invert bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, invert)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_set_invert, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInvert() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_get_invert, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetIn3dSpace(enable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_set_in_3d_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsIn3dSpace() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_is_in_3d_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGenerateMipmaps(invert bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, invert)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_set_generate_mipmaps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsGeneratingMipmaps() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_is_generating_mipmaps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSeamless(seamless bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, seamless)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_set_seamless, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSeamless() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_get_seamless, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSeamlessBlendSkirt(seamless_blend_skirt gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, seamless_blend_skirt)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_set_seamless_blend_skirt, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSeamlessBlendSkirt() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_get_seamless_blend_skirt, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAsNormalMap(as_normal_map bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, as_normal_map)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_set_as_normal_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsNormalMap() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_is_normal_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBumpStrength(bump_strength gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, bump_strength)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_set_bump_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBumpStrength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_get_bump_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetNormalize(normalize bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, normalize)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_set_normalize, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsNormalized() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_is_normalized, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetColorRamp(gradient gdclass.Gradient)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(gradient[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_set_color_ramp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetColorRamp() gdclass.Gradient {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_get_color_ramp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Gradient{classdb.Gradient(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetNoise(noise gdclass.Noise)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(noise[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_set_noise, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetNoise() gdclass.Noise {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_get_noise, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Noise{classdb.Noise(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self class) AsNoiseTexture2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsNoiseTexture2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsTexture2D() Texture2D.GD { return *((*Texture2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTexture2D() Texture2D.Go { return *((*Texture2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsTexture() Texture.GD { return *((*Texture.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTexture() Texture.Go { return *((*Texture.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTexture2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTexture2D(), name)
	}
}
func init() {classdb.Register("NoiseTexture2D", func(ptr gd.Object) any { return classdb.NoiseTexture2D(ptr) })}
