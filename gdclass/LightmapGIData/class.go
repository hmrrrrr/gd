package LightmapGIData

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
[LightmapGIData] contains baked lightmap and dynamic object probe data for [LightmapGI]. It is replaced every time lightmaps are baked in [LightmapGI].

*/
type Go [1]classdb.LightmapGIData

/*
Adds an object that is considered baked within this [LightmapGIData].
*/
func (self Go) AddUser(path string, uv_scale gd.Rect2, slice_index int, sub_instance int) {
	class(self).AddUser(gd.NewString(path).NodePath(), uv_scale, gd.Int(slice_index), gd.Int(sub_instance))
}

/*
Returns the number of objects that are considered baked within this [LightmapGIData].
*/
func (self Go) GetUserCount() int {
	return int(int(class(self).GetUserCount()))
}

/*
Returns the [NodePath] of the baked object at index [param user_idx].
*/
func (self Go) GetUserPath(user_idx int) string {
	return string(class(self).GetUserPath(gd.Int(user_idx)).String())
}

/*
Clear all objects that are considered baked within this [LightmapGIData].
*/
func (self Go) ClearUsers() {
	class(self).ClearUsers()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.LightmapGIData
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("LightmapGIData"))
	return Go{classdb.LightmapGIData(object)}
}

func (self Go) LightmapTextures() gd.Array {
		return gd.Array(class(self).GetLightmapTextures())
}

func (self Go) SetLightmapTextures(value gd.Array) {
	class(self).SetLightmapTextures(value)
}

func (self Go) UsesSphericalHarmonics() bool {
		return bool(class(self).IsUsingSphericalHarmonics())
}

func (self Go) SetUsesSphericalHarmonics(value bool) {
	class(self).SetUsesSphericalHarmonics(value)
}

func (self Go) LightTexture() gdclass.TextureLayered {
		return gdclass.TextureLayered(class(self).GetLightTexture())
}

func (self Go) SetLightTexture(value gdclass.TextureLayered) {
	class(self).SetLightTexture(value)
}

//go:nosplit
func (self class) SetLightmapTextures(light_textures gd.Array)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(light_textures))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGIData.Bind_set_lightmap_textures, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLightmapTextures() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGIData.Bind_get_lightmap_textures, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
If [param uses_spherical_harmonics] is [code]true[/code], tells the engine to treat the lightmap data as if it was baked with directional information.
[b]Note:[/b] Changing this value on already baked lightmaps will not cause them to be baked again. This means the material appearance will look incorrect until lightmaps are baked again, in which case the value set here is discarded as the entire [LightmapGIData] resource is replaced by the lightmapper.
*/
//go:nosplit
func (self class) SetUsesSphericalHarmonics(uses_spherical_harmonics bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, uses_spherical_harmonics)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGIData.Bind_set_uses_spherical_harmonics, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], lightmaps were baked with directional information. See also [member LightmapGI.directional].
*/
//go:nosplit
func (self class) IsUsingSphericalHarmonics() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGIData.Bind_is_using_spherical_harmonics, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds an object that is considered baked within this [LightmapGIData].
*/
//go:nosplit
func (self class) AddUser(path gd.NodePath, uv_scale gd.Rect2, slice_index gd.Int, sub_instance gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(path))
	callframe.Arg(frame, uv_scale)
	callframe.Arg(frame, slice_index)
	callframe.Arg(frame, sub_instance)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGIData.Bind_add_user, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of objects that are considered baked within this [LightmapGIData].
*/
//go:nosplit
func (self class) GetUserCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGIData.Bind_get_user_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [NodePath] of the baked object at index [param user_idx].
*/
//go:nosplit
func (self class) GetUserPath(user_idx gd.Int) gd.NodePath {
	var frame = callframe.New()
	callframe.Arg(frame, user_idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGIData.Bind_get_user_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}
/*
Clear all objects that are considered baked within this [LightmapGIData].
*/
//go:nosplit
func (self class) ClearUsers()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGIData.Bind_clear_users, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetLightTexture(light_texture gdclass.TextureLayered)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(light_texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGIData.Bind_set_light_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLightTexture() gdclass.TextureLayered {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LightmapGIData.Bind_get_light_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.TextureLayered{classdb.TextureLayered(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self class) AsLightmapGIData() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsLightmapGIData() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("LightmapGIData", func(ptr gd.Object) any { return classdb.LightmapGIData(ptr) })}
