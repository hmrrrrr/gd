// Package ImageTexture3D provides methods for working with ImageTexture3D object instances.
package ImageTexture3D

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
import "graphics.gd/classdb/Texture3D"
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
[ImageTexture3D] is a 3-dimensional [ImageTexture] that has a width, height, and depth. See also [ImageTextureLayered].
3D textures are typically used to store density maps for [FogMaterial], color correction LUTs for [Environment], vector fields for [GPUParticlesAttractorVectorField3D] and collision maps for [GPUParticlesCollisionSDF3D]. 3D textures can also be used in custom shaders.
*/
type Instance [1]gdclass.ImageTexture3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsImageTexture3D() Instance
}

/*
Creates the [ImageTexture3D] with specified [param width], [param height], and [param depth]. See [enum Image.Format] for [param format] options. If [param use_mipmaps] is [code]true[/code], then generate mipmaps for the [ImageTexture3D].
*/
func (self Instance) Create(format gdclass.ImageFormat, width int, height int, depth int, use_mipmaps bool, data [][1]gdclass.Image) error { //gd:ImageTexture3D.create
	return error(gd.ToError(class(self).Create(format, int64(width), int64(height), int64(depth), use_mipmaps, gd.ArrayFromSlice[Array.Contains[[1]gdclass.Image]](data))))
}

/*
Replaces the texture's existing data with the layers specified in [param data]. The size of [param data] must match the parameters that were used for [method create]. In other words, the texture cannot be resized or have its format changed by calling [method update].
*/
func (self Instance) Update(data [][1]gdclass.Image) { //gd:ImageTexture3D.update
	class(self).Update(gd.ArrayFromSlice[Array.Contains[[1]gdclass.Image]](data))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ImageTexture3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ImageTexture3D"))
	casted := Instance{*(*gdclass.ImageTexture3D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Creates the [ImageTexture3D] with specified [param width], [param height], and [param depth]. See [enum Image.Format] for [param format] options. If [param use_mipmaps] is [code]true[/code], then generate mipmaps for the [ImageTexture3D].
*/
//go:nosplit
func (self class) Create(format gdclass.ImageFormat, width int64, height int64, depth int64, use_mipmaps bool, data Array.Contains[[1]gdclass.Image]) Error.Code { //gd:ImageTexture3D.create
	var frame = callframe.New()
	callframe.Arg(frame, format)
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, depth)
	callframe.Arg(frame, use_mipmaps)
	callframe.Arg(frame, pointers.Get(gd.InternalArray(data)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImageTexture3D.Bind_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Replaces the texture's existing data with the layers specified in [param data]. The size of [param data] must match the parameters that were used for [method create]. In other words, the texture cannot be resized or have its format changed by calling [method update].
*/
//go:nosplit
func (self class) Update(data Array.Contains[[1]gdclass.Image]) { //gd:ImageTexture3D.update
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(data)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ImageTexture3D.Bind_update, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsImageTexture3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsImageTexture3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsTexture3D() Texture3D.Advanced {
	return *((*Texture3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsTexture3D() Texture3D.Instance {
	return *((*Texture3D.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(Texture3D.Advanced(self.AsTexture3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Texture3D.Instance(self.AsTexture3D()), name)
	}
}
func init() {
	gdclass.Register("ImageTexture3D", func(ptr gd.Object) any {
		return [1]gdclass.ImageTexture3D{*(*gdclass.ImageTexture3D)(unsafe.Pointer(&ptr))}
	})
}
