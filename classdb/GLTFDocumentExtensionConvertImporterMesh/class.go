// Package GLTFDocumentExtensionConvertImporterMesh provides methods for working with GLTFDocumentExtensionConvertImporterMesh object instances.
package GLTFDocumentExtensionConvertImporterMesh

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/GLTFDocumentExtension"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type Instance [1]gdclass.GLTFDocumentExtensionConvertImporterMesh
type Any interface {
	gd.IsClass
	AsGLTFDocumentExtensionConvertImporterMesh() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GLTFDocumentExtensionConvertImporterMesh

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFDocumentExtensionConvertImporterMesh"))
	return Instance{*(*gdclass.GLTFDocumentExtensionConvertImporterMesh)(unsafe.Pointer(&object))}
}

func (self class) AsGLTFDocumentExtensionConvertImporterMesh() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsGLTFDocumentExtensionConvertImporterMesh() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsGLTFDocumentExtension() GLTFDocumentExtension.Advanced {
	return *((*GLTFDocumentExtension.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsGLTFDocumentExtension() GLTFDocumentExtension.Instance {
	return *((*GLTFDocumentExtension.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(GLTFDocumentExtension.Advanced(self.AsGLTFDocumentExtension()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(GLTFDocumentExtension.Instance(self.AsGLTFDocumentExtension()), name)
	}
}
func init() {
	gdclass.Register("GLTFDocumentExtensionConvertImporterMesh", func(ptr gd.Object) any {
		return [1]gdclass.GLTFDocumentExtensionConvertImporterMesh{*(*gdclass.GLTFDocumentExtensionConvertImporterMesh)(unsafe.Pointer(&ptr))}
	})
}
