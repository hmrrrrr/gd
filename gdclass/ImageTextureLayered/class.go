package ImageTextureLayered

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/TextureLayered"
import "grow.graphics/gd/gdclass/Texture"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Base class for [Texture2DArray], [Cubemap] and [CubemapArray]. Cannot be used directly, but contains all the functions necessary for accessing the derived resource types. See also [Texture3D].

*/
type Go [1]classdb.ImageTextureLayered

/*
Creates an [ImageTextureLayered] from an array of [Image]s. See [method Image.create] for the expected data format. The first image decides the width, height, image format and mipmapping setting. The other images [i]must[/i] have the same width, height, image format and mipmapping setting.
Each [Image] represents one [code]layer[/code].
*/
func (self Go) CreateFromImages(images gd.ArrayOf[gdclass.Image]) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).CreateFromImages(images))
}

/*
Replaces the existing [Image] data at the given [param layer] with this new image.
The given [Image] must have the same width, height, image format, and mipmapping flag as the rest of the referenced images.
If the image format is unsupported, it will be decompressed and converted to a similar and supported [enum Image.Format].
The update is immediate: it's synchronized with drawing.
*/
func (self Go) UpdateLayer(image gdclass.Image, layer int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).UpdateLayer(image, gd.Int(layer))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ImageTextureLayered
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("ImageTextureLayered"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Creates an [ImageTextureLayered] from an array of [Image]s. See [method Image.create] for the expected data format. The first image decides the width, height, image format and mipmapping setting. The other images [i]must[/i] have the same width, height, image format and mipmapping setting.
Each [Image] represents one [code]layer[/code].
*/
//go:nosplit
func (self class) CreateFromImages(images gd.ArrayOf[gdclass.Image]) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(images))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImageTextureLayered.Bind_create_from_images, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Replaces the existing [Image] data at the given [param layer] with this new image.
The given [Image] must have the same width, height, image format, and mipmapping flag as the rest of the referenced images.
If the image format is unsupported, it will be decompressed and converted to a similar and supported [enum Image.Format].
The update is immediate: it's synchronized with drawing.
*/
//go:nosplit
func (self class) UpdateLayer(image gdclass.Image, layer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(image[0].AsPointer())[0])
	callframe.Arg(frame, layer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ImageTextureLayered.Bind_update_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsImageTextureLayered() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsImageTextureLayered() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
	default: return gd.VirtualByName(self.AsTextureLayered(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTextureLayered(), name)
	}
}
func init() {classdb.Register("ImageTextureLayered", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
