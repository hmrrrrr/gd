// Package PlaceholderMesh provides methods for working with PlaceholderMesh object instances.
package PlaceholderMesh

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/Mesh"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/AABB"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This class is used when loading a project that uses a [Mesh] subclass in 2 conditions:
- When running the project exported in dedicated server mode, only the texture's dimensions are kept (as they may be relied upon for gameplay purposes or positioning of other elements). This allows reducing the exported PCK's size significantly.
- When this subclass is missing due to using a different engine version or build (e.g. modules disabled).
*/
type Instance [1]gdclass.PlaceholderMesh
type Any interface {
	gd.IsClass
	AsPlaceholderMesh() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PlaceholderMesh

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PlaceholderMesh"))
	return Instance{*(*gdclass.PlaceholderMesh)(unsafe.Pointer(&object))}
}

func (self Instance) SetAabb(value AABB.PositionSize) {
	class(self).SetAabb(gd.AABB(value))
}

//go:nosplit
func (self class) SetAabb(aabb gd.AABB) {
	var frame = callframe.New()
	callframe.Arg(frame, aabb)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PlaceholderMesh.Bind_set_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsPlaceholderMesh() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPlaceholderMesh() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsMesh() Mesh.Advanced          { return *((*Mesh.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMesh() Mesh.Instance       { return *((*Mesh.Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(self.AsMesh(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsMesh(), name)
	}
}
func init() {
	gdclass.Register("PlaceholderMesh", func(ptr gd.Object) any {
		return [1]gdclass.PlaceholderMesh{*(*gdclass.PlaceholderMesh)(unsafe.Pointer(&ptr))}
	})
}
