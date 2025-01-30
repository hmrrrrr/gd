// Package ImageTextureLayered provides methods for working with ImageTextureLayered object instances.
package ImageTextureLayered

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Resource"
import "graphics.gd/classdb/Texture"
import "graphics.gd/classdb/TextureLayered"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

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
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
Base class for [Texture2DArray], [Cubemap] and [CubemapArray]. Cannot be used directly, but contains all the functions necessary for accessing the derived resource types. See also [Texture3D].
*/
type Instance [1]gdclass.ImageTextureLayered

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsImageTextureLayered() Instance
}

/*
Creates an [ImageTextureLayered] from an array of [Image]s. See [method Image.create] for the expected data format. The first image decides the width, height, image format and mipmapping setting. The other images [i]must[/i] have the same width, height, image format and mipmapping setting.
Each [Image] represents one [code]layer[/code].
*/
func (self Instance) CreateFromImages(images [][1]gdclass.Image) error { //gd:ImageTextureLayered.create_from_images
	return error(gd.ToError(class(self).CreateFromImages(gd.ArrayFromSlice[Array.Contains[[1]gdclass.Image]](images))))
}

/*
Replaces the existing [Image] data at the given [param layer] with this new image.
The given [Image] must have the same width, height, image format, and mipmapping flag as the rest of the referenced images.
If the image format is unsupported, it will be decompressed and converted to a similar and supported [enum Image.Format].
The update is immediate: it's synchronized with drawing.
*/
func (self Instance) UpdateLayer(image [1]gdclass.Image, layer int) { //gd:ImageTextureLayered.update_layer
	class(self).UpdateLayer(image, int64(layer))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ImageTextureLayered

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ImageTextureLayered"))
	casted := Instance{*(*gdclass.ImageTextureLayered)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Creates an [ImageTextureLayered] from an array of [Image]s. See [method Image.create] for the expected data format. The first image decides the width, height, image format and mipmapping setting. The other images [i]must[/i] have the same width, height, image format and mipmapping setting.
Each [Image] represents one [code]layer[/code].
*/
//go:nosplit
func (self class) CreateFromImages(images Array.Contains[[1]gdclass.Image]) Error.Code { //gd:ImageTextureLayered.create_from_images
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(images)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImageTextureLayered.Bind_create_from_images, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
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
func (self class) UpdateLayer(image [1]gdclass.Image, layer int64) { //gd:ImageTextureLayered.update_layer
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(image[0])[0])
	callframe.Arg(frame, layer)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImageTextureLayered.Bind_update_layer, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsImageTextureLayered() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsImageTextureLayered() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(TextureLayered.Advanced(self.AsTextureLayered()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(TextureLayered.Instance(self.AsTextureLayered()), name)
	}
}
func init() {
	gdclass.Register("ImageTextureLayered", func(ptr gd.Object) any {
		return [1]gdclass.ImageTextureLayered{*(*gdclass.ImageTextureLayered)(unsafe.Pointer(&ptr))}
	})
}
