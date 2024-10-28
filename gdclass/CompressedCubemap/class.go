package CompressedCubemap

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/CompressedTextureLayered"
import "grow.graphics/gd/gdclass/TextureLayered"
import "grow.graphics/gd/gdclass/Texture"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
A cubemap that is loaded from a [code].ccube[/code] file. This file format is internal to Godot; it is created by importing other image formats with the import system. [CompressedCubemap] can use one of 4 compression methods:
- Lossless (WebP or PNG, uncompressed on the GPU)
- Lossy (WebP, uncompressed on the GPU)
- VRAM Compressed (compressed on the GPU)
- VRAM Uncompressed (uncompressed on the GPU)
- Basis Universal (compressed on the GPU. Lower file sizes than VRAM Compressed, but slower to compress and lower quality than VRAM Compressed)
Only [b]VRAM Compressed[/b] actually reduces the memory usage on the GPU. The [b]Lossless[/b] and [b]Lossy[/b] compression methods will reduce the required storage on disk, but they will not reduce memory usage on the GPU as the texture is sent to the GPU uncompressed.
Using [b]VRAM Compressed[/b] also improves loading times, as VRAM-compressed textures are faster to load compared to textures using lossless or lossy compression. VRAM compression can exhibit noticeable artifacts and is intended to be used for 3D rendering, not 2D.
See [Cubemap] for a general description of cubemaps.

*/
type Go [1]classdb.CompressedCubemap
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.CompressedCubemap
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CompressedCubemap"))
	return Go{classdb.CompressedCubemap(object)}
}

func (self class) AsCompressedCubemap() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsCompressedCubemap() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsCompressedTextureLayered() CompressedTextureLayered.GD { return *((*CompressedTextureLayered.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCompressedTextureLayered() CompressedTextureLayered.Go { return *((*CompressedTextureLayered.Go)(unsafe.Pointer(&self))) }
func (self class) AsTextureLayered() TextureLayered.GD { return *((*TextureLayered.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTextureLayered() TextureLayered.Go { return *((*TextureLayered.Go)(unsafe.Pointer(&self))) }
func (self class) AsTexture() Texture.GD { return *((*Texture.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTexture() Texture.Go { return *((*Texture.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsCompressedTextureLayered(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsCompressedTextureLayered(), name)
	}
}
func init() {classdb.Register("CompressedCubemap", func(ptr gd.Object) any { return classdb.CompressedCubemap(ptr) })}
