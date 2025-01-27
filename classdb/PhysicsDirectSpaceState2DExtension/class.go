// Package PhysicsDirectSpaceState2DExtension provides methods for working with PhysicsDirectSpaceState2DExtension object instances.
package PhysicsDirectSpaceState2DExtension

import "unsafe"
import "reflect"
import "slices"
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
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/Packed"
import "graphics.gd/classdb/PhysicsDirectSpaceState2D"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Transform2D"
import "graphics.gd/variant/Float"

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
var _ = slices.Delete[[]struct{}, struct{}]

/*
This class extends [PhysicsDirectSpaceState2D] by providing additional virtual methods that can be overridden. When these methods are overridden, they will be called instead of the internal methods of the physics server.
Intended for use with GDExtension to create custom implementations of [PhysicsDirectSpaceState2D].

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=PhysicsDirectSpaceState2DExtension)
*/
type Instance [1]gdclass.PhysicsDirectSpaceState2DExtension

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPhysicsDirectSpaceState2DExtension() Instance
}
type Interface interface {
	IntersectRay(from Vector2.XY, to Vector2.XY, collision_mask int, collide_with_bodies bool, collide_with_areas bool, hit_from_inside bool, result *PhysicsServer2DExtensionRayResult) bool
	IntersectPoint(position Vector2.XY, canvas_instance_id int, collision_mask int, collide_with_bodies bool, collide_with_areas bool, results *PhysicsServer2DExtensionShapeResult, max_results int) int
	IntersectShape(shape_rid RID.Any, transform Transform2D.OriginXY, motion Vector2.XY, margin Float.X, collision_mask int, collide_with_bodies bool, collide_with_areas bool, result *PhysicsServer2DExtensionShapeResult, max_results int) int
	CastMotion(shape_rid RID.Any, transform Transform2D.OriginXY, motion Vector2.XY, margin Float.X, collision_mask int, collide_with_bodies bool, collide_with_areas bool, closest_safe *float64, closest_unsafe *float64) bool
	CollideShape(shape_rid RID.Any, transform Transform2D.OriginXY, motion Vector2.XY, margin Float.X, collision_mask int, collide_with_bodies bool, collide_with_areas bool, results unsafe.Pointer, max_results int, result_count *int32) bool
	RestInfo(shape_rid RID.Any, transform Transform2D.OriginXY, motion Vector2.XY, margin Float.X, collision_mask int, collide_with_bodies bool, collide_with_areas bool, rest_info *PhysicsServer2DExtensionShapeRestInfo) bool
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) IntersectRay(from Vector2.XY, to Vector2.XY, collision_mask int, collide_with_bodies bool, collide_with_areas bool, hit_from_inside bool, result *PhysicsServer2DExtensionRayResult) (_ bool) {
	return
}
func (self implementation) IntersectPoint(position Vector2.XY, canvas_instance_id int, collision_mask int, collide_with_bodies bool, collide_with_areas bool, results *PhysicsServer2DExtensionShapeResult, max_results int) (_ int) {
	return
}
func (self implementation) IntersectShape(shape_rid RID.Any, transform Transform2D.OriginXY, motion Vector2.XY, margin Float.X, collision_mask int, collide_with_bodies bool, collide_with_areas bool, result *PhysicsServer2DExtensionShapeResult, max_results int) (_ int) {
	return
}
func (self implementation) CastMotion(shape_rid RID.Any, transform Transform2D.OriginXY, motion Vector2.XY, margin Float.X, collision_mask int, collide_with_bodies bool, collide_with_areas bool, closest_safe *float64, closest_unsafe *float64) (_ bool) {
	return
}
func (self implementation) CollideShape(shape_rid RID.Any, transform Transform2D.OriginXY, motion Vector2.XY, margin Float.X, collision_mask int, collide_with_bodies bool, collide_with_areas bool, results unsafe.Pointer, max_results int, result_count *int32) (_ bool) {
	return
}
func (self implementation) RestInfo(shape_rid RID.Any, transform Transform2D.OriginXY, motion Vector2.XY, margin Float.X, collision_mask int, collide_with_bodies bool, collide_with_areas bool, rest_info *PhysicsServer2DExtensionShapeRestInfo) (_ bool) {
	return
}
func (Instance) _intersect_ray(impl func(ptr unsafe.Pointer, from Vector2.XY, to Vector2.XY, collision_mask int, collide_with_bodies bool, collide_with_areas bool, hit_from_inside bool, result *PhysicsServer2DExtensionRayResult) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var from = gd.UnsafeGet[gd.Vector2](p_args, 0)

		var to = gd.UnsafeGet[gd.Vector2](p_args, 1)

		var collision_mask = gd.UnsafeGet[gd.Int](p_args, 2)

		var collide_with_bodies = gd.UnsafeGet[bool](p_args, 3)

		var collide_with_areas = gd.UnsafeGet[bool](p_args, 4)

		var hit_from_inside = gd.UnsafeGet[bool](p_args, 5)

		var result = gd.UnsafeGet[*PhysicsServer2DExtensionRayResult](p_args, 6)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from, to, int(collision_mask), collide_with_bodies, collide_with_areas, hit_from_inside, result)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _intersect_point(impl func(ptr unsafe.Pointer, position Vector2.XY, canvas_instance_id int, collision_mask int, collide_with_bodies bool, collide_with_areas bool, results *PhysicsServer2DExtensionShapeResult, max_results int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var position = gd.UnsafeGet[gd.Vector2](p_args, 0)

		var canvas_instance_id = gd.UnsafeGet[gd.Int](p_args, 1)

		var collision_mask = gd.UnsafeGet[gd.Int](p_args, 2)

		var collide_with_bodies = gd.UnsafeGet[bool](p_args, 3)

		var collide_with_areas = gd.UnsafeGet[bool](p_args, 4)

		var results = gd.UnsafeGet[*PhysicsServer2DExtensionShapeResult](p_args, 5)

		var max_results = gd.UnsafeGet[gd.Int](p_args, 6)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, position, int(canvas_instance_id), int(collision_mask), collide_with_bodies, collide_with_areas, results, int(max_results))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _intersect_shape(impl func(ptr unsafe.Pointer, shape_rid RID.Any, transform Transform2D.OriginXY, motion Vector2.XY, margin Float.X, collision_mask int, collide_with_bodies bool, collide_with_areas bool, result *PhysicsServer2DExtensionShapeResult, max_results int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var transform = gd.UnsafeGet[gd.Transform2D](p_args, 1)

		var motion = gd.UnsafeGet[gd.Vector2](p_args, 2)

		var margin = gd.UnsafeGet[gd.Float](p_args, 3)

		var collision_mask = gd.UnsafeGet[gd.Int](p_args, 4)

		var collide_with_bodies = gd.UnsafeGet[bool](p_args, 5)

		var collide_with_areas = gd.UnsafeGet[bool](p_args, 6)

		var result = gd.UnsafeGet[*PhysicsServer2DExtensionShapeResult](p_args, 7)

		var max_results = gd.UnsafeGet[gd.Int](p_args, 8)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape_rid, transform, motion, Float.X(margin), int(collision_mask), collide_with_bodies, collide_with_areas, result, int(max_results))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _cast_motion(impl func(ptr unsafe.Pointer, shape_rid RID.Any, transform Transform2D.OriginXY, motion Vector2.XY, margin Float.X, collision_mask int, collide_with_bodies bool, collide_with_areas bool, closest_safe *float64, closest_unsafe *float64) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var transform = gd.UnsafeGet[gd.Transform2D](p_args, 1)

		var motion = gd.UnsafeGet[gd.Vector2](p_args, 2)

		var margin = gd.UnsafeGet[gd.Float](p_args, 3)

		var collision_mask = gd.UnsafeGet[gd.Int](p_args, 4)

		var collide_with_bodies = gd.UnsafeGet[bool](p_args, 5)

		var collide_with_areas = gd.UnsafeGet[bool](p_args, 6)

		var closest_safe = gd.UnsafeGet[*float64](p_args, 7)

		var closest_unsafe = gd.UnsafeGet[*float64](p_args, 8)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape_rid, transform, motion, Float.X(margin), int(collision_mask), collide_with_bodies, collide_with_areas, closest_safe, closest_unsafe)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _collide_shape(impl func(ptr unsafe.Pointer, shape_rid RID.Any, transform Transform2D.OriginXY, motion Vector2.XY, margin Float.X, collision_mask int, collide_with_bodies bool, collide_with_areas bool, results unsafe.Pointer, max_results int, result_count *int32) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var transform = gd.UnsafeGet[gd.Transform2D](p_args, 1)

		var motion = gd.UnsafeGet[gd.Vector2](p_args, 2)

		var margin = gd.UnsafeGet[gd.Float](p_args, 3)

		var collision_mask = gd.UnsafeGet[gd.Int](p_args, 4)

		var collide_with_bodies = gd.UnsafeGet[bool](p_args, 5)

		var collide_with_areas = gd.UnsafeGet[bool](p_args, 6)

		var results = gd.UnsafeGet[unsafe.Pointer](p_args, 7)

		var max_results = gd.UnsafeGet[gd.Int](p_args, 8)

		var result_count = gd.UnsafeGet[*int32](p_args, 9)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape_rid, transform, motion, Float.X(margin), int(collision_mask), collide_with_bodies, collide_with_areas, results, int(max_results), result_count)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _rest_info(impl func(ptr unsafe.Pointer, shape_rid RID.Any, transform Transform2D.OriginXY, motion Vector2.XY, margin Float.X, collision_mask int, collide_with_bodies bool, collide_with_areas bool, rest_info *PhysicsServer2DExtensionShapeRestInfo) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var transform = gd.UnsafeGet[gd.Transform2D](p_args, 1)

		var motion = gd.UnsafeGet[gd.Vector2](p_args, 2)

		var margin = gd.UnsafeGet[gd.Float](p_args, 3)

		var collision_mask = gd.UnsafeGet[gd.Int](p_args, 4)

		var collide_with_bodies = gd.UnsafeGet[bool](p_args, 5)

		var collide_with_areas = gd.UnsafeGet[bool](p_args, 6)

		var rest_info = gd.UnsafeGet[*PhysicsServer2DExtensionShapeRestInfo](p_args, 7)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape_rid, transform, motion, Float.X(margin), int(collision_mask), collide_with_bodies, collide_with_areas, rest_info)
		gd.UnsafeSet(p_back, ret)
	}
}
func (self Instance) IsBodyExcludedFromQuery(body RID.Body2D) bool { //gd:PhysicsDirectSpaceState2DExtension.is_body_excluded_from_query
	return bool(class(self).IsBodyExcludedFromQuery(gd.RID(body)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PhysicsDirectSpaceState2DExtension

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicsDirectSpaceState2DExtension"))
	casted := Instance{*(*gdclass.PhysicsDirectSpaceState2DExtension)(unsafe.Pointer(&object))}
	return casted
}

func (class) _intersect_ray(impl func(ptr unsafe.Pointer, from gd.Vector2, to gd.Vector2, collision_mask gd.Int, collide_with_bodies bool, collide_with_areas bool, hit_from_inside bool, result *PhysicsServer2DExtensionRayResult) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var from = gd.UnsafeGet[gd.Vector2](p_args, 0)

		var to = gd.UnsafeGet[gd.Vector2](p_args, 1)

		var collision_mask = gd.UnsafeGet[gd.Int](p_args, 2)

		var collide_with_bodies = gd.UnsafeGet[bool](p_args, 3)

		var collide_with_areas = gd.UnsafeGet[bool](p_args, 4)

		var hit_from_inside = gd.UnsafeGet[bool](p_args, 5)

		var result = gd.UnsafeGet[*PhysicsServer2DExtensionRayResult](p_args, 6)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from, to, collision_mask, collide_with_bodies, collide_with_areas, hit_from_inside, result)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _intersect_point(impl func(ptr unsafe.Pointer, position gd.Vector2, canvas_instance_id gd.Int, collision_mask gd.Int, collide_with_bodies bool, collide_with_areas bool, results *PhysicsServer2DExtensionShapeResult, max_results gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var position = gd.UnsafeGet[gd.Vector2](p_args, 0)

		var canvas_instance_id = gd.UnsafeGet[gd.Int](p_args, 1)

		var collision_mask = gd.UnsafeGet[gd.Int](p_args, 2)

		var collide_with_bodies = gd.UnsafeGet[bool](p_args, 3)

		var collide_with_areas = gd.UnsafeGet[bool](p_args, 4)

		var results = gd.UnsafeGet[*PhysicsServer2DExtensionShapeResult](p_args, 5)

		var max_results = gd.UnsafeGet[gd.Int](p_args, 6)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, position, canvas_instance_id, collision_mask, collide_with_bodies, collide_with_areas, results, max_results)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _intersect_shape(impl func(ptr unsafe.Pointer, shape_rid gd.RID, transform gd.Transform2D, motion gd.Vector2, margin gd.Float, collision_mask gd.Int, collide_with_bodies bool, collide_with_areas bool, result *PhysicsServer2DExtensionShapeResult, max_results gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var transform = gd.UnsafeGet[gd.Transform2D](p_args, 1)

		var motion = gd.UnsafeGet[gd.Vector2](p_args, 2)

		var margin = gd.UnsafeGet[gd.Float](p_args, 3)

		var collision_mask = gd.UnsafeGet[gd.Int](p_args, 4)

		var collide_with_bodies = gd.UnsafeGet[bool](p_args, 5)

		var collide_with_areas = gd.UnsafeGet[bool](p_args, 6)

		var result = gd.UnsafeGet[*PhysicsServer2DExtensionShapeResult](p_args, 7)

		var max_results = gd.UnsafeGet[gd.Int](p_args, 8)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape_rid, transform, motion, margin, collision_mask, collide_with_bodies, collide_with_areas, result, max_results)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _cast_motion(impl func(ptr unsafe.Pointer, shape_rid gd.RID, transform gd.Transform2D, motion gd.Vector2, margin gd.Float, collision_mask gd.Int, collide_with_bodies bool, collide_with_areas bool, closest_safe *float64, closest_unsafe *float64) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var transform = gd.UnsafeGet[gd.Transform2D](p_args, 1)

		var motion = gd.UnsafeGet[gd.Vector2](p_args, 2)

		var margin = gd.UnsafeGet[gd.Float](p_args, 3)

		var collision_mask = gd.UnsafeGet[gd.Int](p_args, 4)

		var collide_with_bodies = gd.UnsafeGet[bool](p_args, 5)

		var collide_with_areas = gd.UnsafeGet[bool](p_args, 6)

		var closest_safe = gd.UnsafeGet[*float64](p_args, 7)

		var closest_unsafe = gd.UnsafeGet[*float64](p_args, 8)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape_rid, transform, motion, margin, collision_mask, collide_with_bodies, collide_with_areas, closest_safe, closest_unsafe)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _collide_shape(impl func(ptr unsafe.Pointer, shape_rid gd.RID, transform gd.Transform2D, motion gd.Vector2, margin gd.Float, collision_mask gd.Int, collide_with_bodies bool, collide_with_areas bool, results unsafe.Pointer, max_results gd.Int, result_count *int32) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var transform = gd.UnsafeGet[gd.Transform2D](p_args, 1)

		var motion = gd.UnsafeGet[gd.Vector2](p_args, 2)

		var margin = gd.UnsafeGet[gd.Float](p_args, 3)

		var collision_mask = gd.UnsafeGet[gd.Int](p_args, 4)

		var collide_with_bodies = gd.UnsafeGet[bool](p_args, 5)

		var collide_with_areas = gd.UnsafeGet[bool](p_args, 6)

		var results = gd.UnsafeGet[unsafe.Pointer](p_args, 7)

		var max_results = gd.UnsafeGet[gd.Int](p_args, 8)

		var result_count = gd.UnsafeGet[*int32](p_args, 9)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape_rid, transform, motion, margin, collision_mask, collide_with_bodies, collide_with_areas, results, max_results, result_count)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _rest_info(impl func(ptr unsafe.Pointer, shape_rid gd.RID, transform gd.Transform2D, motion gd.Vector2, margin gd.Float, collision_mask gd.Int, collide_with_bodies bool, collide_with_areas bool, rest_info *PhysicsServer2DExtensionShapeRestInfo) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape_rid = gd.UnsafeGet[gd.RID](p_args, 0)

		var transform = gd.UnsafeGet[gd.Transform2D](p_args, 1)

		var motion = gd.UnsafeGet[gd.Vector2](p_args, 2)

		var margin = gd.UnsafeGet[gd.Float](p_args, 3)

		var collision_mask = gd.UnsafeGet[gd.Int](p_args, 4)

		var collide_with_bodies = gd.UnsafeGet[bool](p_args, 5)

		var collide_with_areas = gd.UnsafeGet[bool](p_args, 6)

		var rest_info = gd.UnsafeGet[*PhysicsServer2DExtensionShapeRestInfo](p_args, 7)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape_rid, transform, motion, margin, collision_mask, collide_with_bodies, collide_with_areas, rest_info)
		gd.UnsafeSet(p_back, ret)
	}
}

//go:nosplit
func (self class) IsBodyExcludedFromQuery(body gd.RID) bool { //gd:PhysicsDirectSpaceState2DExtension.is_body_excluded_from_query
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectSpaceState2DExtension.Bind_is_body_excluded_from_query, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPhysicsDirectSpaceState2DExtension() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPhysicsDirectSpaceState2DExtension() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsPhysicsDirectSpaceState2D() PhysicsDirectSpaceState2D.Advanced {
	return *((*PhysicsDirectSpaceState2D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPhysicsDirectSpaceState2D() PhysicsDirectSpaceState2D.Instance {
	return *((*PhysicsDirectSpaceState2D.Instance)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_intersect_ray":
		return reflect.ValueOf(self._intersect_ray)
	case "_intersect_point":
		return reflect.ValueOf(self._intersect_point)
	case "_intersect_shape":
		return reflect.ValueOf(self._intersect_shape)
	case "_cast_motion":
		return reflect.ValueOf(self._cast_motion)
	case "_collide_shape":
		return reflect.ValueOf(self._collide_shape)
	case "_rest_info":
		return reflect.ValueOf(self._rest_info)
	default:
		return gd.VirtualByName(PhysicsDirectSpaceState2D.Advanced(self.AsPhysicsDirectSpaceState2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_intersect_ray":
		return reflect.ValueOf(self._intersect_ray)
	case "_intersect_point":
		return reflect.ValueOf(self._intersect_point)
	case "_intersect_shape":
		return reflect.ValueOf(self._intersect_shape)
	case "_cast_motion":
		return reflect.ValueOf(self._cast_motion)
	case "_collide_shape":
		return reflect.ValueOf(self._collide_shape)
	case "_rest_info":
		return reflect.ValueOf(self._rest_info)
	default:
		return gd.VirtualByName(PhysicsDirectSpaceState2D.Instance(self.AsPhysicsDirectSpaceState2D()), name)
	}
}
func init() {
	gdclass.Register("PhysicsDirectSpaceState2DExtension", func(ptr gd.Object) any {
		return [1]gdclass.PhysicsDirectSpaceState2DExtension{*(*gdclass.PhysicsDirectSpaceState2DExtension)(unsafe.Pointer(&ptr))}
	})
}
