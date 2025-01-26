// Package NavigationPathQueryResult3D provides methods for working with NavigationPathQueryResult3D object instances.
package NavigationPathQueryResult3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Vector3"
import "graphics.gd/variant/RID"

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

/*
This class stores the result of a 3D navigation path query from the [NavigationServer3D].
*/
type Instance [1]gdclass.NavigationPathQueryResult3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsNavigationPathQueryResult3D() Instance
}

/*
Reset the result object to its initial state. This is useful to reuse the object across multiple queries.
*/
func (self Instance) Reset() { //gd:NavigationPathQueryResult3D.reset
	class(self).Reset()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.NavigationPathQueryResult3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("NavigationPathQueryResult3D"))
	casted := Instance{*(*gdclass.NavigationPathQueryResult3D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Path() []Vector3.XYZ {
	return []Vector3.XYZ(class(self).GetPath().AsSlice())
}

func (self Instance) SetPath(value []Vector3.XYZ) {
	class(self).SetPath(gd.NewPackedVector3Slice(*(*[]gd.Vector3)(unsafe.Pointer(&value))))
}

func (self Instance) PathTypes() []int32 {
	return []int32(class(self).GetPathTypes().AsSlice())
}

func (self Instance) SetPathTypes(value []int32) {
	class(self).SetPathTypes(gd.NewPackedInt32Slice(value))
}

func (self Instance) PathRids() []RID.Any {
	return []RID.Any(gd.ArrayAs[[]RID.Any](gd.InternalArray(class(self).GetPathRids())))
}

func (self Instance) SetPathRids(value []RID.Any) {
	class(self).SetPathRids(gd.ArrayFromSlice[Array.Contains[gd.RID]](value))
}

func (self Instance) PathOwnerIds() []int64 {
	return []int64(class(self).GetPathOwnerIds().AsSlice())
}

func (self Instance) SetPathOwnerIds(value []int64) {
	class(self).SetPathOwnerIds(gd.NewPackedInt64Slice(value))
}

//go:nosplit
func (self class) SetPath(path gd.PackedVector3Array) { //gd:NavigationPathQueryResult3D.set_path
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryResult3D.Bind_set_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPath() gd.PackedVector3Array { //gd:NavigationPathQueryResult3D.get_path
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryResult3D.Bind_get_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedVector3Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPathTypes(path_types gd.PackedInt32Array) { //gd:NavigationPathQueryResult3D.set_path_types
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path_types))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryResult3D.Bind_set_path_types, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPathTypes() gd.PackedInt32Array { //gd:NavigationPathQueryResult3D.get_path_types
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryResult3D.Bind_get_path_types, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPathRids(path_rids Array.Contains[gd.RID]) { //gd:NavigationPathQueryResult3D.set_path_rids
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(path_rids)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryResult3D.Bind_set_path_rids, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPathRids() Array.Contains[gd.RID] { //gd:NavigationPathQueryResult3D.get_path_rids
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryResult3D.Bind_get_path_rids, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.RID]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPathOwnerIds(path_owner_ids gd.PackedInt64Array) { //gd:NavigationPathQueryResult3D.set_path_owner_ids
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path_owner_ids))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryResult3D.Bind_set_path_owner_ids, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPathOwnerIds() gd.PackedInt64Array { //gd:NavigationPathQueryResult3D.get_path_owner_ids
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryResult3D.Bind_get_path_owner_ids, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedInt64Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Reset the result object to its initial state. This is useful to reuse the object across multiple queries.
*/
//go:nosplit
func (self class) Reset() { //gd:NavigationPathQueryResult3D.reset
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationPathQueryResult3D.Bind_reset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsNavigationPathQueryResult3D() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsNavigationPathQueryResult3D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("NavigationPathQueryResult3D", func(ptr gd.Object) any {
		return [1]gdclass.NavigationPathQueryResult3D{*(*gdclass.NavigationPathQueryResult3D)(unsafe.Pointer(&ptr))}
	})
}

type PathSegmentType = gdclass.NavigationPathQueryResult3DPathSegmentType //gd:NavigationPathQueryResult3D.PathSegmentType

const (
	/*This segment of the path goes through a region.*/
	PathSegmentTypeRegion PathSegmentType = 0
	/*This segment of the path goes through a link.*/
	PathSegmentTypeLink PathSegmentType = 1
)
