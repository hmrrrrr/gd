// Package EditorExportPlatformAndroid provides methods for working with EditorExportPlatformAndroid object instances.
package EditorExportPlatformAndroid

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/EditorExportPlatform"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type Instance [1]gdclass.EditorExportPlatformAndroid
type Any interface {
	gd.IsClass
	AsEditorExportPlatformAndroid() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorExportPlatformAndroid

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorExportPlatformAndroid"))
	return Instance{*(*gdclass.EditorExportPlatformAndroid)(unsafe.Pointer(&object))}
}

func (self class) AsEditorExportPlatformAndroid() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsEditorExportPlatformAndroid() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsEditorExportPlatform() EditorExportPlatform.Advanced {
	return *((*EditorExportPlatform.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsEditorExportPlatform() EditorExportPlatform.Instance {
	return *((*EditorExportPlatform.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsEditorExportPlatform(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsEditorExportPlatform(), name)
	}
}
func init() {
	gdclass.Register("EditorExportPlatformAndroid", func(ptr gd.Object) any {
		return [1]gdclass.EditorExportPlatformAndroid{*(*gdclass.EditorExportPlatformAndroid)(unsafe.Pointer(&ptr))}
	})
}
