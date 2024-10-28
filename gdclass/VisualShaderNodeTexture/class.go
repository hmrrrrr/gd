package VisualShaderNodeTexture

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VisualShaderNode"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Performs a lookup operation on the provided texture, with support for multiple texture sources to choose from.

*/
type Go [1]classdb.VisualShaderNodeTexture
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VisualShaderNodeTexture
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeTexture"))
	return Go{classdb.VisualShaderNodeTexture(object)}
}

func (self Go) Source() classdb.VisualShaderNodeTextureSource {
		return classdb.VisualShaderNodeTextureSource(class(self).GetSource())
}

func (self Go) SetSource(value classdb.VisualShaderNodeTextureSource) {
	class(self).SetSource(value)
}

func (self Go) Texture() gdclass.Texture2D {
		return gdclass.Texture2D(class(self).GetTexture())
}

func (self Go) SetTexture(value gdclass.Texture2D) {
	class(self).SetTexture(value)
}

func (self Go) TextureType() classdb.VisualShaderNodeTextureTextureType {
		return classdb.VisualShaderNodeTextureTextureType(class(self).GetTextureType())
}

func (self Go) SetTextureType(value classdb.VisualShaderNodeTextureTextureType) {
	class(self).SetTextureType(value)
}

//go:nosplit
func (self class) SetSource(value classdb.VisualShaderNodeTextureSource)  {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTexture.Bind_set_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSource() classdb.VisualShaderNodeTextureSource {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeTextureSource](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTexture.Bind_get_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTexture(value gdclass.Texture2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(value[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTexture.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTexture() gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTexture.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTextureType(value classdb.VisualShaderNodeTextureTextureType)  {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTexture.Bind_set_texture_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTextureType() classdb.VisualShaderNodeTextureTextureType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeTextureTextureType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTexture.Bind_get_texture_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeTexture() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeTexture() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNode() VisualShaderNode.GD { return *((*VisualShaderNode.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNode() VisualShaderNode.Go { return *((*VisualShaderNode.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNode(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNode(), name)
	}
}
func init() {classdb.Register("VisualShaderNodeTexture", func(ptr gd.Object) any { return classdb.VisualShaderNodeTexture(ptr) })}
type Source = classdb.VisualShaderNodeTextureSource

const (
/*Use the texture given as an argument for this function.*/
	SourceTexture Source = 0
/*Use the current viewport's texture as the source.*/
	SourceScreen Source = 1
/*Use the texture from this shader's texture built-in (e.g. a texture of a [Sprite2D]).*/
	Source2dTexture Source = 2
/*Use the texture from this shader's normal map built-in.*/
	Source2dNormal Source = 3
/*Use the depth texture captured during the depth prepass. Only available when the depth prepass is used (i.e. in spatial shaders and in the forward_plus or gl_compatibility renderers).*/
	SourceDepth Source = 4
/*Use the texture provided in the input port for this function.*/
	SourcePort Source = 5
/*Use the normal buffer captured during the depth prepass. Only available when the normal-roughness buffer is available (i.e. in spatial shaders and in the forward_plus renderer).*/
	Source3dNormal Source = 6
/*Use the roughness buffer captured during the depth prepass. Only available when the normal-roughness buffer is available (i.e. in spatial shaders and in the forward_plus renderer).*/
	SourceRoughness Source = 7
/*Represents the size of the [enum Source] enum.*/
	SourceMax Source = 8
)
type TextureType = classdb.VisualShaderNodeTextureTextureType

const (
/*No hints are added to the uniform declaration.*/
	TypeData TextureType = 0
/*Adds [code]source_color[/code] as hint to the uniform declaration for proper sRGB to linear conversion.*/
	TypeColor TextureType = 1
/*Adds [code]hint_normal[/code] as hint to the uniform declaration, which internally converts the texture for proper usage as normal map.*/
	TypeNormalMap TextureType = 2
/*Represents the size of the [enum TextureType] enum.*/
	TypeMax TextureType = 3
)
