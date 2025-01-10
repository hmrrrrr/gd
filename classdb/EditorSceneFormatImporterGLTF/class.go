// Package EditorSceneFormatImporterGLTF provides methods for working with EditorSceneFormatImporterGLTF object instances.
package EditorSceneFormatImporterGLTF

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/EditorSceneFormatImporter"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

type Instance [1]gdclass.EditorSceneFormatImporterGLTF

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorSceneFormatImporterGLTF() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorSceneFormatImporterGLTF

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorSceneFormatImporterGLTF"))
	casted := Instance{*(*gdclass.EditorSceneFormatImporterGLTF)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self class) AsEditorSceneFormatImporterGLTF() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsEditorSceneFormatImporterGLTF() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsEditorSceneFormatImporter() EditorSceneFormatImporter.Advanced {
	return *((*EditorSceneFormatImporter.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsEditorSceneFormatImporter() EditorSceneFormatImporter.Instance {
	return *((*EditorSceneFormatImporter.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(EditorSceneFormatImporter.Advanced(self.AsEditorSceneFormatImporter()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(EditorSceneFormatImporter.Instance(self.AsEditorSceneFormatImporter()), name)
	}
}
func init() {
	gdclass.Register("EditorSceneFormatImporterGLTF", func(ptr gd.Object) any {
		return [1]gdclass.EditorSceneFormatImporterGLTF{*(*gdclass.EditorSceneFormatImporterGLTF)(unsafe.Pointer(&ptr))}
	})
}
