package CubemapArray

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/ImageTextureLayered"
import "grow.graphics/gd/gdclass/TextureLayered"
import "grow.graphics/gd/gdclass/Texture"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
[CubemapArray]s are made of an array of [Cubemap]s. Like [Cubemap]s, they are made of multiple textures, the amount of which must be divisible by 6 (one for each face of the cube).
The primary benefit of [CubemapArray]s is that they can be accessed in shader code using a single texture reference. In other words, you can pass multiple [Cubemap]s into a shader using a single [CubemapArray]. [Cubemap]s are allocated in adjacent cache regions on the GPU, which makes [CubemapArray]s the most efficient way to store multiple [Cubemap]s.
[b]Note:[/b] Godot uses [CubemapArray]s internally for many effects, including the [Sky] if you set [member ProjectSettings.rendering/reflections/sky_reflections/texture_array_reflections] to [code]true[/code]. To create such a texture file yourself, reimport your image files using the import presets of the File System dock.
[b]Note:[/b] [CubemapArray] is not supported in the OpenGL 3 rendering backend.
*/
type Instance [1]classdb.CubemapArray

/*
Creates a placeholder version of this resource ([PlaceholderCubemapArray]).
*/
func (self Instance) CreatePlaceholder() gdclass.Resource {
	return gdclass.Resource(class(self).CreatePlaceholder())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.CubemapArray

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CubemapArray"))
	return Instance{classdb.CubemapArray(object)}
}

/*
Creates a placeholder version of this resource ([PlaceholderCubemapArray]).
*/
//go:nosplit
func (self class) CreatePlaceholder() gdclass.Resource {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CubemapArray.Bind_create_placeholder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Resource{classdb.Resource(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self class) AsCubemapArray() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCubemapArray() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsImageTextureLayered() ImageTextureLayered.Advanced {
	return *((*ImageTextureLayered.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsImageTextureLayered() ImageTextureLayered.Instance {
	return *((*ImageTextureLayered.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsTextureLayered() TextureLayered.Advanced {
	return *((*TextureLayered.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsTextureLayered() TextureLayered.Instance {
	return *((*TextureLayered.Instance)(unsafe.Pointer(&self)))
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsImageTextureLayered(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsImageTextureLayered(), name)
	}
}
func init() {
	classdb.Register("CubemapArray", func(ptr gd.Object) any { return classdb.CubemapArray(ptr) })
}
