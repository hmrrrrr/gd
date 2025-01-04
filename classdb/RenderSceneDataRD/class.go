// Package RenderSceneDataRD provides methods for working with RenderSceneDataRD object instances.
package RenderSceneDataRD

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/RenderSceneData"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Object holds scene data related to rendering a single frame of a viewport.
[b]Note:[/b] This is an internal rendering server object, do not instantiate this from script.
*/
type Instance [1]gdclass.RenderSceneDataRD
type Any interface {
	gd.IsClass
	AsRenderSceneDataRD() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RenderSceneDataRD

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RenderSceneDataRD"))
	return Instance{*(*gdclass.RenderSceneDataRD)(unsafe.Pointer(&object))}
}

func (self class) AsRenderSceneDataRD() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRenderSceneDataRD() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRenderSceneData() RenderSceneData.Advanced {
	return *((*RenderSceneData.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsRenderSceneData() RenderSceneData.Instance {
	return *((*RenderSceneData.Instance)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRenderSceneData(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRenderSceneData(), name)
	}
}
func init() {
	gdclass.Register("RenderSceneDataRD", func(ptr gd.Object) any {
		return [1]gdclass.RenderSceneDataRD{*(*gdclass.RenderSceneDataRD)(unsafe.Pointer(&ptr))}
	})
}
