// Package PolygonPathFinder provides methods for working with PolygonPathFinder object instances.
package PolygonPathFinder

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Rect2"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type Instance [1]gdclass.PolygonPathFinder
type Any interface {
	gd.IsClass
	AsPolygonPathFinder() Instance
}

func (self Instance) Setup(points []Vector2.XY, connections []int32) {
	class(self).Setup(gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&points))), gd.NewPackedInt32Slice(connections))
}
func (self Instance) FindPath(from Vector2.XY, to Vector2.XY) []Vector2.XY {
	return []Vector2.XY(class(self).FindPath(gd.Vector2(from), gd.Vector2(to)).AsSlice())
}
func (self Instance) GetIntersections(from Vector2.XY, to Vector2.XY) []Vector2.XY {
	return []Vector2.XY(class(self).GetIntersections(gd.Vector2(from), gd.Vector2(to)).AsSlice())
}
func (self Instance) GetClosestPoint(point Vector2.XY) Vector2.XY {
	return Vector2.XY(class(self).GetClosestPoint(gd.Vector2(point)))
}
func (self Instance) IsPointInside(point Vector2.XY) bool {
	return bool(class(self).IsPointInside(gd.Vector2(point)))
}
func (self Instance) SetPointPenalty(idx int, penalty Float.X) {
	class(self).SetPointPenalty(gd.Int(idx), gd.Float(penalty))
}
func (self Instance) GetPointPenalty(idx int) Float.X {
	return Float.X(Float.X(class(self).GetPointPenalty(gd.Int(idx))))
}
func (self Instance) GetBounds() Rect2.PositionSize {
	return Rect2.PositionSize(class(self).GetBounds())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PolygonPathFinder

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PolygonPathFinder"))
	casted := Instance{*(*gdclass.PolygonPathFinder)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

//go:nosplit
func (self class) Setup(points gd.PackedVector2Array, connections gd.PackedInt32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(points))
	callframe.Arg(frame, pointers.Get(connections))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PolygonPathFinder.Bind_setup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) FindPath(from gd.Vector2, to gd.Vector2) gd.PackedVector2Array {
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PolygonPathFinder.Bind_find_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetIntersections(from gd.Vector2, to gd.Vector2) gd.PackedVector2Array {
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PolygonPathFinder.Bind_get_intersections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetClosestPoint(point gd.Vector2) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PolygonPathFinder.Bind_get_closest_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) IsPointInside(point gd.Vector2) bool {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PolygonPathFinder.Bind_is_point_inside, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPointPenalty(idx gd.Int, penalty gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, penalty)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PolygonPathFinder.Bind_set_point_penalty, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPointPenalty(idx gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PolygonPathFinder.Bind_get_point_penalty, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetBounds() gd.Rect2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PolygonPathFinder.Bind_get_bounds, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPolygonPathFinder() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPolygonPathFinder() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("PolygonPathFinder", func(ptr gd.Object) any {
		return [1]gdclass.PolygonPathFinder{*(*gdclass.PolygonPathFinder)(unsafe.Pointer(&ptr))}
	})
}
