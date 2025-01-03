package FBXDocument

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/GLTFDocument"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
The FBXDocument handles FBX documents. It provides methods to append data from buffers or files, generate scenes, and register/unregister document extensions.
When exporting FBX from Blender, use the "FBX Units Scale" option. The "FBX Units Scale" option sets the correct scale factor and avoids manual adjustments when re-importing into Blender, such as through glTF export.
*/
type Instance [1]classdb.FBXDocument
type Any interface {
	gd.IsClass
	AsFBXDocument() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.FBXDocument

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("FBXDocument"))
	return Instance{*(*classdb.FBXDocument)(unsafe.Pointer(&object))}
}

func (self class) AsFBXDocument() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsFBXDocument() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsGLTFDocument() GLTFDocument.Advanced {
	return *((*GLTFDocument.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsGLTFDocument() GLTFDocument.Instance {
	return *((*GLTFDocument.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsGLTFDocument(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsGLTFDocument(), name)
	}
}
func init() {
	classdb.Register("FBXDocument", func(ptr gd.Object) any { return [1]classdb.FBXDocument{*(*classdb.FBXDocument)(unsafe.Pointer(&ptr))} })
}
