// Package TextureCubemapRD provides methods for working with TextureCubemapRD object instances.
package TextureCubemapRD

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/TextureLayeredRD"
import "graphics.gd/classdb/TextureLayered"
import "graphics.gd/classdb/Texture"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This texture class allows you to use a cubemap texture created directly on the [RenderingDevice] as a texture for materials, meshes, etc.
*/
type Instance [1]gdclass.TextureCubemapRD
type Any interface {
	gd.IsClass
	AsTextureCubemapRD() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.TextureCubemapRD

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TextureCubemapRD"))
	return Instance{*(*gdclass.TextureCubemapRD)(unsafe.Pointer(&object))}
}

func (self class) AsTextureCubemapRD() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTextureCubemapRD() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsTextureLayeredRD() TextureLayeredRD.Advanced {
	return *((*TextureLayeredRD.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsTextureLayeredRD() TextureLayeredRD.Instance {
	return *((*TextureLayeredRD.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsTextureLayeredRD(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsTextureLayeredRD(), name)
	}
}
func init() {
	gdclass.Register("TextureCubemapRD", func(ptr gd.Object) any {
		return [1]gdclass.TextureCubemapRD{*(*gdclass.TextureCubemapRD)(unsafe.Pointer(&ptr))}
	})
}
