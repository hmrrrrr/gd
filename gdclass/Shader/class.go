package Shader

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
A custom shader program implemented in the Godot shading language, saved with the [code].gdshader[/code] extension.
This class is used by a [ShaderMaterial] and allows you to write your own custom behavior for rendering visual items or updating particle information. For a detailed explanation and usage, please see the tutorials linked below.

*/
type Go [1]classdb.Shader

/*
Returns the shader mode for the shader.
*/
func (self Go) GetMode() classdb.ShaderMode {
	return classdb.ShaderMode(class(self).GetMode())
}

/*
Sets the default texture to be used with a texture uniform. The default is used if a texture is not set in the [ShaderMaterial].
[b]Note:[/b] [param name] must match the name of the uniform in the code exactly.
[b]Note:[/b] If the sampler array is used use [param index] to access the specified texture.
*/
func (self Go) SetDefaultTextureParameter(name string, texture gdclass.Texture2D) {
	class(self).SetDefaultTextureParameter(gd.NewStringName(name), texture, gd.Int(0))
}

/*
Returns the texture that is set as default for the specified parameter.
[b]Note:[/b] [param name] must match the name of the uniform in the code exactly.
[b]Note:[/b] If the sampler array is used use [param index] to access the specified texture.
*/
func (self Go) GetDefaultTextureParameter(name string) gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetDefaultTextureParameter(gd.NewStringName(name), gd.Int(0)))
}

/*
Get the list of shader uniforms that can be assigned to a [ShaderMaterial], for use with [method ShaderMaterial.set_shader_parameter] and [method ShaderMaterial.get_shader_parameter]. The parameters returned are contained in dictionaries in a similar format to the ones returned by [method Object.get_property_list].
If argument [param get_groups] is true, parameter grouping hints will be provided.
*/
func (self Go) GetShaderUniformList() gd.Array {
	return gd.Array(class(self).GetShaderUniformList(false))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Shader
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Shader"))
	return Go{classdb.Shader(object)}
}

func (self Go) Code() string {
		return string(class(self).GetCode().String())
}

func (self Go) SetCode(value string) {
	class(self).SetCode(gd.NewString(value))
}

/*
Returns the shader mode for the shader.
*/
//go:nosplit
func (self class) GetMode() classdb.ShaderMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ShaderMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shader.Bind_get_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCode(code gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(code))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shader.Bind_set_code, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCode() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shader.Bind_get_code, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the default texture to be used with a texture uniform. The default is used if a texture is not set in the [ShaderMaterial].
[b]Note:[/b] [param name] must match the name of the uniform in the code exactly.
[b]Note:[/b] If the sampler array is used use [param index] to access the specified texture.
*/
//go:nosplit
func (self class) SetDefaultTextureParameter(name gd.StringName, texture gdclass.Texture2D, index gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
	callframe.Arg(frame, discreet.Get(texture[0])[0])
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shader.Bind_set_default_texture_parameter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the texture that is set as default for the specified parameter.
[b]Note:[/b] [param name] must match the name of the uniform in the code exactly.
[b]Note:[/b] If the sampler array is used use [param index] to access the specified texture.
*/
//go:nosplit
func (self class) GetDefaultTextureParameter(name gd.StringName, index gd.Int) gdclass.Texture2D {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shader.Bind_get_default_texture_parameter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Get the list of shader uniforms that can be assigned to a [ShaderMaterial], for use with [method ShaderMaterial.set_shader_parameter] and [method ShaderMaterial.get_shader_parameter]. The parameters returned are contained in dictionaries in a similar format to the ones returned by [method Object.get_property_list].
If argument [param get_groups] is true, parameter grouping hints will be provided.
*/
//go:nosplit
func (self class) GetShaderUniformList(get_groups bool) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, get_groups)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shader.Bind_get_shader_uniform_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsShader() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsShader() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("Shader", func(ptr gd.Object) any { return classdb.Shader(ptr) })}
type Mode = classdb.ShaderMode

const (
/*Mode used to draw all 3D objects.*/
	ModeSpatial Mode = 0
/*Mode used to draw all 2D objects.*/
	ModeCanvasItem Mode = 1
/*Mode used to calculate particle information on a per-particle basis. Not used for drawing.*/
	ModeParticles Mode = 2
/*Mode used for drawing skies. Only works with shaders attached to [Sky] objects.*/
	ModeSky Mode = 3
/*Mode used for setting the color and density of volumetric fog effect.*/
	ModeFog Mode = 4
)
