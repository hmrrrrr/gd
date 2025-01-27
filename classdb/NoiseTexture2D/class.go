// Package NoiseTexture2D provides methods for working with NoiseTexture2D object instances.
package NoiseTexture2D

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
import "graphics.gd/classdb/Texture2D"
import "graphics.gd/classdb/Texture"
import "graphics.gd/classdb/Resource"
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
type Instance [1]gdclass.NoiseTexture2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsNoiseTexture2D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.NoiseTexture2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("NoiseTexture2D"))
	casted := Instance{*(*gdclass.NoiseTexture2D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) SetWidth(value int) {
	class(self).SetWidth(gd.Int(value))
}

func (self Instance) SetHeight(value int) {
	class(self).SetHeight(gd.Int(value))
}

func (self Instance) Invert() bool {
	return bool(class(self).GetInvert())
}

func (self Instance) SetInvert(value bool) {
	class(self).SetInvert(value)
}

func (self Instance) In3dSpace() bool {
	return bool(class(self).IsIn3dSpace())
}

func (self Instance) SetIn3dSpace(value bool) {
	class(self).SetIn3dSpace(value)
}

func (self Instance) GenerateMipmaps() bool {
	return bool(class(self).IsGeneratingMipmaps())
}

func (self Instance) SetGenerateMipmaps(value bool) {
	class(self).SetGenerateMipmaps(value)
}

func (self Instance) Seamless() bool {
	return bool(class(self).GetSeamless())
}

func (self Instance) SetSeamless(value bool) {
	class(self).SetSeamless(value)
}

func (self Instance) SeamlessBlendSkirt() Float.X {
	return Float.X(Float.X(class(self).GetSeamlessBlendSkirt()))
}

func (self Instance) SetSeamlessBlendSkirt(value Float.X) {
	class(self).SetSeamlessBlendSkirt(gd.Float(value))
}

func (self Instance) AsNormalMap() bool {
	return bool(class(self).IsNormalMap())
}

func (self Instance) SetAsNormalMap(value bool) {
	class(self).SetAsNormalMap(value)
}

func (self Instance) BumpStrength() Float.X {
	return Float.X(Float.X(class(self).GetBumpStrength()))
}

func (self Instance) SetBumpStrength(value Float.X) {
	class(self).SetBumpStrength(gd.Float(value))
}

func (self Instance) Normalize() bool {
	return bool(class(self).IsNormalized())
}

func (self Instance) SetNormalize(value bool) {
	class(self).SetNormalize(value)
}

func (self Instance) ColorRamp() [1]gdclass.Gradient {
	return [1]gdclass.Gradient(class(self).GetColorRamp())
}

func (self Instance) SetColorRamp(value [1]gdclass.Gradient) {
	class(self).SetColorRamp(value)
}

func (self Instance) Noise() [1]gdclass.Noise {
	return [1]gdclass.Noise(class(self).GetNoise())
}

func (self Instance) SetNoise(value [1]gdclass.Noise) {
	class(self).SetNoise(value)
}

//go:nosplit
func (self class) SetWidth(width gd.Int) { //gd:NoiseTexture2D.set_width
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_set_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetHeight(height gd.Int) { //gd:NoiseTexture2D.set_height
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_set_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetInvert(invert bool) { //gd:NoiseTexture2D.set_invert
	var frame = callframe.New()
	callframe.Arg(frame, invert)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_set_invert, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetInvert() bool { //gd:NoiseTexture2D.get_invert
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_get_invert, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetIn3dSpace(enable bool) { //gd:NoiseTexture2D.set_in_3d_space
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_set_in_3d_space, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsIn3dSpace() bool { //gd:NoiseTexture2D.is_in_3d_space
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_is_in_3d_space, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGenerateMipmaps(invert bool) { //gd:NoiseTexture2D.set_generate_mipmaps
	var frame = callframe.New()
	callframe.Arg(frame, invert)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_set_generate_mipmaps, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsGeneratingMipmaps() bool { //gd:NoiseTexture2D.is_generating_mipmaps
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_is_generating_mipmaps, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSeamless(seamless bool) { //gd:NoiseTexture2D.set_seamless
	var frame = callframe.New()
	callframe.Arg(frame, seamless)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_set_seamless, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSeamless() bool { //gd:NoiseTexture2D.get_seamless
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_get_seamless, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSeamlessBlendSkirt(seamless_blend_skirt gd.Float) { //gd:NoiseTexture2D.set_seamless_blend_skirt
	var frame = callframe.New()
	callframe.Arg(frame, seamless_blend_skirt)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_set_seamless_blend_skirt, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSeamlessBlendSkirt() gd.Float { //gd:NoiseTexture2D.get_seamless_blend_skirt
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_get_seamless_blend_skirt, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAsNormalMap(as_normal_map bool) { //gd:NoiseTexture2D.set_as_normal_map
	var frame = callframe.New()
	callframe.Arg(frame, as_normal_map)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_set_as_normal_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsNormalMap() bool { //gd:NoiseTexture2D.is_normal_map
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_is_normal_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBumpStrength(bump_strength gd.Float) { //gd:NoiseTexture2D.set_bump_strength
	var frame = callframe.New()
	callframe.Arg(frame, bump_strength)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_set_bump_strength, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBumpStrength() gd.Float { //gd:NoiseTexture2D.get_bump_strength
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_get_bump_strength, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNormalize(normalize bool) { //gd:NoiseTexture2D.set_normalize
	var frame = callframe.New()
	callframe.Arg(frame, normalize)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_set_normalize, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsNormalized() bool { //gd:NoiseTexture2D.is_normalized
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_is_normalized, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetColorRamp(gradient [1]gdclass.Gradient) { //gd:NoiseTexture2D.set_color_ramp
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gradient[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_set_color_ramp, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetColorRamp() [1]gdclass.Gradient { //gd:NoiseTexture2D.get_color_ramp
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_get_color_ramp, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Gradient{gd.PointerWithOwnershipTransferredToGo[gdclass.Gradient](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNoise(noise [1]gdclass.Noise) { //gd:NoiseTexture2D.set_noise
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(noise[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_set_noise, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetNoise() [1]gdclass.Noise { //gd:NoiseTexture2D.get_noise
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NoiseTexture2D.Bind_get_noise, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Noise{gd.PointerWithOwnershipTransferredToGo[gdclass.Noise](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsNoiseTexture2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNoiseTexture2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsTexture2D() Texture2D.Advanced {
	return *((*Texture2D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsTexture2D() Texture2D.Instance {
	return *((*Texture2D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsTexture() Texture.Advanced { return *((*Texture.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTexture() Texture.Instance {
	return *((*Texture.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(Texture2D.Advanced(self.AsTexture2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Texture2D.Instance(self.AsTexture2D()), name)
	}
}
func init() {
	gdclass.Register("NoiseTexture2D", func(ptr gd.Object) any {
		return [1]gdclass.NoiseTexture2D{*(*gdclass.NoiseTexture2D)(unsafe.Pointer(&ptr))}
	})
}
