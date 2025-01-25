// Package Decal provides methods for working with Decal object instances.
package Decal

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
import "graphics.gd/classdb/VisualInstance3D"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Vector3"
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
[Decal]s are used to project a texture onto a [Mesh] in the scene. Use Decals to add detail to a scene without affecting the underlying [Mesh]. They are often used to add weathering to building, add dirt or mud to the ground, or add variety to props. Decals can be moved at any time, making them suitable for things like blob shadows or laser sight dots.
They are made of an [AABB] and a group of [Texture2D]s specifying [Color], normal, ORM (ambient occlusion, roughness, metallic), and emission. Decals are projected within their [AABB] so altering the orientation of the Decal affects the direction in which they are projected. By default, Decals are projected down (i.e. from positive Y to negative Y).
The [Texture2D]s associated with the Decal are automatically stored in a texture atlas which is used for drawing the decals so all decals can be drawn at once. Godot uses clustered decals, meaning they are stored in cluster data and drawn when the mesh is drawn, they are not drawn as a post-processing effect after.
[b]Note:[/b] Decals cannot affect an underlying material's transparency, regardless of its transparency mode (alpha blend, alpha scissor, alpha hash, opaque pre-pass). This means translucent or transparent areas of a material will remain translucent or transparent even if an opaque decal is applied on them.
[b]Note:[/b] Decals are only supported in the Forward+ and Mobile rendering methods, not Compatibility. When using the Mobile rendering method, only 8 decals can be displayed on each mesh resource. Attempting to display more than 8 decals on a single mesh resource will result in decals flickering in and out as the camera moves.
[b]Note:[/b] When using the Mobile rendering method, decals will only correctly affect meshes whose visibility AABB intersects with the decal's AABB. If using a shader to deform the mesh in a way that makes it go outside its AABB, [member GeometryInstance3D.extra_cull_margin] must be increased on the mesh. Otherwise, the decal may not be visible on the mesh.
*/
type Instance [1]gdclass.Decal

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsDecal() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Decal

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Decal"))
	casted := Instance{*(*gdclass.Decal)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Size() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetSize())
}

func (self Instance) SetSize(value Vector3.XYZ) {
	class(self).SetSize(gd.Vector3(value))
}

func (self Instance) TextureAlbedo() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexture(0))
}

func (self Instance) SetTextureAlbedo(value [1]gdclass.Texture2D) {
	class(self).SetTexture(0, value)
}

func (self Instance) TextureNormal() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexture(1))
}

func (self Instance) SetTextureNormal(value [1]gdclass.Texture2D) {
	class(self).SetTexture(1, value)
}

func (self Instance) TextureOrm() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexture(2))
}

func (self Instance) SetTextureOrm(value [1]gdclass.Texture2D) {
	class(self).SetTexture(2, value)
}

func (self Instance) TextureEmission() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexture(3))
}

func (self Instance) SetTextureEmission(value [1]gdclass.Texture2D) {
	class(self).SetTexture(3, value)
}

func (self Instance) EmissionEnergy() Float.X {
	return Float.X(Float.X(class(self).GetEmissionEnergy()))
}

func (self Instance) SetEmissionEnergy(value Float.X) {
	class(self).SetEmissionEnergy(gd.Float(value))
}

func (self Instance) Modulate() Color.RGBA {
	return Color.RGBA(class(self).GetModulate())
}

func (self Instance) SetModulate(value Color.RGBA) {
	class(self).SetModulate(gd.Color(value))
}

func (self Instance) AlbedoMix() Float.X {
	return Float.X(Float.X(class(self).GetAlbedoMix()))
}

func (self Instance) SetAlbedoMix(value Float.X) {
	class(self).SetAlbedoMix(gd.Float(value))
}

func (self Instance) NormalFade() Float.X {
	return Float.X(Float.X(class(self).GetNormalFade()))
}

func (self Instance) SetNormalFade(value Float.X) {
	class(self).SetNormalFade(gd.Float(value))
}

func (self Instance) UpperFade() Float.X {
	return Float.X(Float.X(class(self).GetUpperFade()))
}

func (self Instance) SetUpperFade(value Float.X) {
	class(self).SetUpperFade(gd.Float(value))
}

func (self Instance) LowerFade() Float.X {
	return Float.X(Float.X(class(self).GetLowerFade()))
}

func (self Instance) SetLowerFade(value Float.X) {
	class(self).SetLowerFade(gd.Float(value))
}

func (self Instance) DistanceFadeEnabled() bool {
	return bool(class(self).IsDistanceFadeEnabled())
}

func (self Instance) SetDistanceFadeEnabled(value bool) {
	class(self).SetEnableDistanceFade(value)
}

func (self Instance) DistanceFadeBegin() Float.X {
	return Float.X(Float.X(class(self).GetDistanceFadeBegin()))
}

func (self Instance) SetDistanceFadeBegin(value Float.X) {
	class(self).SetDistanceFadeBegin(gd.Float(value))
}

func (self Instance) DistanceFadeLength() Float.X {
	return Float.X(Float.X(class(self).GetDistanceFadeLength()))
}

func (self Instance) SetDistanceFadeLength(value Float.X) {
	class(self).SetDistanceFadeLength(gd.Float(value))
}

func (self Instance) CullMask() int {
	return int(int(class(self).GetCullMask()))
}

func (self Instance) SetCullMask(value int) {
	class(self).SetCullMask(gd.Int(value))
}

//go:nosplit
func (self class) SetSize(size gd.Vector3) { //gd:Decal.set_size
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Decal.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSize() gd.Vector3 { //gd:Decal.get_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Decal.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [Texture2D] associated with the specified [enum DecalTexture]. This is a convenience method, in most cases you should access the texture directly.
For example, instead of [code]$Decal.set_texture(Decal.TEXTURE_ALBEDO, albedo_tex)[/code], use [code]$Decal.texture_albedo = albedo_tex[/code].
One case where this is better than accessing the texture directly is when you want to copy one Decal's textures to another. For example:
[codeblocks]
[gdscript]
for i in Decal.TEXTURE_MAX:
    $NewDecal.set_texture(i, $OldDecal.get_texture(i))
[/gdscript]
[csharp]
for (int i = 0; i < (int)Decal.DecalTexture.Max; i++)
{
    GetNode<Decal>("NewDecal").SetTexture(i, GetNode<Decal>("OldDecal").GetTexture(i));
}
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) SetTexture(atype gdclass.DecalDecalTexture, texture [1]gdclass.Texture2D) { //gd:Decal.set_texture
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Decal.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [Texture2D] associated with the specified [enum DecalTexture]. This is a convenience method, in most cases you should access the texture directly.
For example, instead of [code]albedo_tex = $Decal.get_texture(Decal.TEXTURE_ALBEDO)[/code], use [code]albedo_tex = $Decal.texture_albedo[/code].
One case where this is better than accessing the texture directly is when you want to copy one Decal's textures to another. For example:
[codeblocks]
[gdscript]
for i in Decal.TEXTURE_MAX:
    $NewDecal.set_texture(i, $OldDecal.get_texture(i))
[/gdscript]
[csharp]
for (int i = 0; i < (int)Decal.DecalTexture.Max; i++)
{
    GetNode<Decal>("NewDecal").SetTexture(i, GetNode<Decal>("OldDecal").GetTexture(i));
}
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) GetTexture(atype gdclass.DecalDecalTexture) [1]gdclass.Texture2D { //gd:Decal.get_texture
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Decal.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmissionEnergy(energy gd.Float) { //gd:Decal.set_emission_energy
	var frame = callframe.New()
	callframe.Arg(frame, energy)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Decal.Bind_set_emission_energy, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEmissionEnergy() gd.Float { //gd:Decal.get_emission_energy
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Decal.Bind_get_emission_energy, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAlbedoMix(energy gd.Float) { //gd:Decal.set_albedo_mix
	var frame = callframe.New()
	callframe.Arg(frame, energy)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Decal.Bind_set_albedo_mix, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAlbedoMix() gd.Float { //gd:Decal.get_albedo_mix
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Decal.Bind_get_albedo_mix, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetModulate(color gd.Color) { //gd:Decal.set_modulate
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Decal.Bind_set_modulate, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetModulate() gd.Color { //gd:Decal.get_modulate
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Decal.Bind_get_modulate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUpperFade(fade gd.Float) { //gd:Decal.set_upper_fade
	var frame = callframe.New()
	callframe.Arg(frame, fade)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Decal.Bind_set_upper_fade, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUpperFade() gd.Float { //gd:Decal.get_upper_fade
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Decal.Bind_get_upper_fade, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLowerFade(fade gd.Float) { //gd:Decal.set_lower_fade
	var frame = callframe.New()
	callframe.Arg(frame, fade)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Decal.Bind_set_lower_fade, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLowerFade() gd.Float { //gd:Decal.get_lower_fade
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Decal.Bind_get_lower_fade, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNormalFade(fade gd.Float) { //gd:Decal.set_normal_fade
	var frame = callframe.New()
	callframe.Arg(frame, fade)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Decal.Bind_set_normal_fade, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetNormalFade() gd.Float { //gd:Decal.get_normal_fade
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Decal.Bind_get_normal_fade, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnableDistanceFade(enable bool) { //gd:Decal.set_enable_distance_fade
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Decal.Bind_set_enable_distance_fade, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsDistanceFadeEnabled() bool { //gd:Decal.is_distance_fade_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Decal.Bind_is_distance_fade_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDistanceFadeBegin(distance gd.Float) { //gd:Decal.set_distance_fade_begin
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Decal.Bind_set_distance_fade_begin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDistanceFadeBegin() gd.Float { //gd:Decal.get_distance_fade_begin
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Decal.Bind_get_distance_fade_begin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDistanceFadeLength(distance gd.Float) { //gd:Decal.set_distance_fade_length
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Decal.Bind_set_distance_fade_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDistanceFadeLength() gd.Float { //gd:Decal.get_distance_fade_length
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Decal.Bind_get_distance_fade_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCullMask(mask gd.Int) { //gd:Decal.set_cull_mask
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Decal.Bind_set_cull_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCullMask() gd.Int { //gd:Decal.get_cull_mask
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Decal.Bind_get_cull_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsDecal() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsDecal() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(VisualInstance3D.Advanced(self.AsVisualInstance3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualInstance3D.Instance(self.AsVisualInstance3D()), name)
	}
}
func init() {
	gdclass.Register("Decal", func(ptr gd.Object) any { return [1]gdclass.Decal{*(*gdclass.Decal)(unsafe.Pointer(&ptr))} })
}

type DecalTexture = gdclass.DecalDecalTexture //gd:Decal.DecalTexture

const (
	/*[Texture2D] corresponding to [member texture_albedo].*/
	TextureAlbedo DecalTexture = 0
	/*[Texture2D] corresponding to [member texture_normal].*/
	TextureNormal DecalTexture = 1
	/*[Texture2D] corresponding to [member texture_orm].*/
	TextureOrm DecalTexture = 2
	/*[Texture2D] corresponding to [member texture_emission].*/
	TextureEmission DecalTexture = 3
	/*Max size of [enum DecalTexture] enum.*/
	TextureMax DecalTexture = 4
)
