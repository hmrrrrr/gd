package ImageTexture3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Texture3D"
import "grow.graphics/gd/gdclass/Texture"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
[ImageTexture3D] is a 3-dimensional [ImageTexture] that has a width, height, and depth. See also [ImageTextureLayered].
3D textures are typically used to store density maps for [FogMaterial], color correction LUTs for [Environment], vector fields for [GPUParticlesAttractorVectorField3D] and collision maps for [GPUParticlesCollisionSDF3D]. 3D textures can also be used in custom shaders.

*/
type Go [1]classdb.ImageTexture3D

/*
Creates the [ImageTexture3D] with specified [param width], [param height], and [param depth]. See [enum Image.Format] for [param format] options. If [param use_mipmaps] is [code]true[/code], then generate mipmaps for the [ImageTexture3D].
*/
func (self Go) Create(format classdb.ImageFormat, width int, height int, depth int, use_mipmaps bool, data gd.Array) gd.Error {
	return gd.Error(class(self).Create(format, gd.Int(width), gd.Int(height), gd.Int(depth), use_mipmaps, data))
}

/*
Replaces the texture's existing data with the layers specified in [param data]. The size of [param data] must match the parameters that were used for [method create]. In other words, the texture cannot be resized or have its format changed by calling [method update].
*/
func (self Go) Update(data gd.Array) {
	class(self).Update(data)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ImageTexture3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ImageTexture3D"))
	return Go{classdb.ImageTexture3D(object)}
}

/*
Creates the [ImageTexture3D] with specified [param width], [param height], and [param depth]. See [enum Image.Format] for [param format] options. If [param use_mipmaps] is [code]true[/code], then generate mipmaps for the [ImageTexture3D].
*/
//go:nosplit
func (self class) Create(format classdb.ImageFormat, width gd.Int, height gd.Int, depth gd.Int, use_mipmaps bool, data gd.Array) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, format)
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, depth)
	callframe.Arg(frame, use_mipmaps)
	callframe.Arg(frame, discreet.Get(data))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImageTexture3D.Bind_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Replaces the texture's existing data with the layers specified in [param data]. The size of [param data] must match the parameters that were used for [method create]. In other words, the texture cannot be resized or have its format changed by calling [method update].
*/
//go:nosplit
func (self class) Update(data gd.Array)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(data))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImageTexture3D.Bind_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsImageTexture3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsImageTexture3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsTexture3D() Texture3D.GD { return *((*Texture3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTexture3D() Texture3D.Go { return *((*Texture3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsTexture() Texture.GD { return *((*Texture.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTexture() Texture.Go { return *((*Texture.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTexture3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTexture3D(), name)
	}
}
func init() {classdb.Register("ImageTexture3D", func(ptr gd.Object) any { return classdb.ImageTexture3D(ptr) })}
