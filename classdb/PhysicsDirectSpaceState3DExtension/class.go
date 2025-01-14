// Package PhysicsDirectSpaceState3DExtension provides methods for working with PhysicsDirectSpaceState3DExtension object instances.
package PhysicsDirectSpaceState3DExtension

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/PhysicsDirectSpaceState3D"
import "graphics.gd/variant/Vector3"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Transform3D"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This class extends [PhysicsDirectSpaceState3D] by providing additional virtual methods that can be overridden. When these methods are overridden, they will be called instead of the internal methods of the physics server.
Intended for use with GDExtension to create custom implementations of [PhysicsDirectSpaceState3D].

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=PhysicsDirectSpaceState3DExtension)
*/
type Instance [1]gdclass.PhysicsDirectSpaceState3DExtension

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPhysicsDirectSpaceState3DExtension() Instance
}
type Interface interface {
	IntersectRay(from Vector3.XYZ, to Vector3.XYZ, collision_mask int, collide_with_bodies bool, collide_with_areas bool, hit_from_inside bool, hit_back_faces bool, pick_ray bool, result *PhysicsServer3DExtensionRayResult) bool
	IntersectPoint(position Vector3.XYZ, collision_mask int, collide_with_bodies bool, collide_with_areas bool, results *PhysicsServer3DExtensionShapeResult, max_results int) int
	IntersectShape(shape_rid Resource.ID, transform Transform3D.BasisOrigin, motion Vector3.XYZ, margin Float.X, collision_mask int, collide_with_bodies bool, collide_with_areas bool, result_count *PhysicsServer3DExtensionShapeResult, max_results int) int
	CastMotion(shape_rid Resource.ID, transform Transform3D.BasisOrigin, motion Vector3.XYZ, margin Float.X, collision_mask int, collide_with_bodies bool, collide_with_areas bool, closest_safe *float64, closest_unsafe *float64, info *PhysicsServer3DExtensionShapeRestInfo) bool
	CollideShape(shape_rid Resource.ID, transform Transform3D.BasisOrigin, motion Vector3.XYZ, margin Float.X, collision_mask int, collide_with_bodies bool, collide_with_areas bool, results unsafe.Pointer, max_results int, result_count *int32) bool
	RestInfo(shape_rid Resource.ID, transform Transform3D.BasisOrigin, motion Vector3.XYZ, margin Float.X, collision_mask int, collide_with_bodies bool, collide_with_areas bool, rest_info *PhysicsServer3DExtensionShapeRestInfo) bool
	GetClosestPointToObjectVolume(obj Resource.ID, point Vector3.XYZ) Vector3.XYZ
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) IntersectRay(from Vector3.XYZ, to Vector3.XYZ, collision_mask int, collide_with_bodies bool, collide_with_areas bool, hit_from_inside bool, hit_back_faces bool, pick_ray bool, result *PhysicsServer3DExtensionRayResult) (_ bool) {
	return
}
func (self implementation) IntersectPoint(position Vector3.XYZ, collision_mask int, collide_with_bodies bool, collide_with_areas bool, results *PhysicsServer3DExtensionShapeResult, max_results int) (_ int) {
	return
}
func (self implementation) IntersectShape(shape_rid Resource.ID, transform Transform3D.BasisOrigin, motion Vector3.XYZ, margin Float.X, collision_mask int, collide_with_bodies bool, collide_with_areas bool, result_count *PhysicsServer3DExtensionShapeResult, max_results int) (_ int) {
	return
}
func (self implementation) CastMotion(shape_rid Resource.ID, transform Transform3D.BasisOrigin, motion Vector3.XYZ, margin Float.X, collision_mask int, collide_with_bodies bool, collide_with_areas bool, closest_safe *float64, closest_unsafe *float64, info *PhysicsServer3DExtensionShapeRestInfo) (_ bool) {
	return
}
func (self implementation) CollideShape(shape_rid Resource.ID, transform Transform3D.BasisOrigin, motion Vector3.XYZ, margin Float.X, collision_mask int, collide_with_bodies bool, collide_with_areas bool, results unsafe.Pointer, max_results int, result_count *int32) (_ bool) {
	return
}
func (self implementation) RestInfo(shape_rid Resource.ID, transform Transform3D.BasisOrigin, motion Vector3.XYZ, margin Float.X, collision_mask int, collide_with_bodies bool, collide_with_areas bool, rest_info *PhysicsServer3DExtensionShapeRestInfo) (_ bool) {
	return
}
func (self implementation) GetClosestPointToObjectVolume(obj Resource.ID, point Vector3.XYZ) (_ Vector3.XYZ) {
	return
}
func (Instance) _intersect_ray(impl func(ptr unsafe.Pointer, from Vector3.XYZ, to Vector3.XYZ, collision_mask int, collide_with_bodies bool, collide_with_areas bool, hit_from_inside bool, hit_back_faces bool, pick_ray bool, result *PhysicsServer3DExtensionRayResult) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var from = gd.UnsafeGet[gd.Vector3](p_args, 0)
		var to = gd.UnsafeGet[gd.Vector3](p_args, 1)
		var collision_mask = gd.UnsafeGet[gd.Int](p_args, 2)
		var collide_with_bodies = gd.UnsafeGet[bool](p_args, 3)
		var collide_with_areas = gd.UnsafeGet[bool](p_args, 4)
		var hit_from_inside = gd.UnsafeGet[bool](p_args, 5)
		var hit_back_faces = gd.UnsafeGet[bool](p_args, 6)
		var pick_ray = gd.UnsafeGet[bool](p_args, 7)
		var result = gd.UnsafeGet[*PhysicsServer3DExtensionRayResult](p_args, 8)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from, to, int(collision_mask), collide_with_bodies, collide_with_areas, hit_from_inside, hit_back_faces, pick_ray, result)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _intersect_point(impl func(ptr unsafe.Pointer, position Vector3.XYZ, collision_mask int, collide_with_bodies bool, collide_with_areas bool, results *PhysicsServer3DExtensionShapeResult, max_results int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var position = gd.UnsafeGet[gd.Vector3](p_args, 0)
		var collision_mask = gd.UnsafeGet[gd.Int](p_args, 1)
		var collide_with_bodies = gd.UnsafeGet[bool](p_args, 2)
		var collide_with_areas = gd.UnsafeGet[bool](p_args, 3)
		var results = gd.UnsafeGet[*PhysicsServer3DExtensionShapeResult](p_args, 4)
		var max_results = gd.UnsafeGet[gd.Int](p_args, 5)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, position, int(collision_mask), collide_with_bodies, collide_with_areas, results, int(max_results))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _intersect_shape(impl func(ptr unsafe.Pointer, shape_rid Resource.ID, transform Transform3D.BasisOrigin, motion Vector3.XYZ, margin Float.X, collision_mask int, collide_with_bodies bool, collide_with_areas bool, result_count *PhysicsServer3DExtensionShapeResult, max_results int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape_rid = gd.UnsafeGet[gd.RID](p_args, 0)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 1)
		var motion = gd.UnsafeGet[gd.Vector3](p_args, 2)
		var margin = gd.UnsafeGet[gd.Float](p_args, 3)
		var collision_mask = gd.UnsafeGet[gd.Int](p_args, 4)
		var collide_with_bodies = gd.UnsafeGet[bool](p_args, 5)
		var collide_with_areas = gd.UnsafeGet[bool](p_args, 6)
		var result_count = gd.UnsafeGet[*PhysicsServer3DExtensionShapeResult](p_args, 7)
		var max_results = gd.UnsafeGet[gd.Int](p_args, 8)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape_rid, transform, motion, Float.X(margin), int(collision_mask), collide_with_bodies, collide_with_areas, result_count, int(max_results))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _cast_motion(impl func(ptr unsafe.Pointer, shape_rid Resource.ID, transform Transform3D.BasisOrigin, motion Vector3.XYZ, margin Float.X, collision_mask int, collide_with_bodies bool, collide_with_areas bool, closest_safe *float64, closest_unsafe *float64, info *PhysicsServer3DExtensionShapeRestInfo) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape_rid = gd.UnsafeGet[gd.RID](p_args, 0)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 1)
		var motion = gd.UnsafeGet[gd.Vector3](p_args, 2)
		var margin = gd.UnsafeGet[gd.Float](p_args, 3)
		var collision_mask = gd.UnsafeGet[gd.Int](p_args, 4)
		var collide_with_bodies = gd.UnsafeGet[bool](p_args, 5)
		var collide_with_areas = gd.UnsafeGet[bool](p_args, 6)
		var closest_safe = gd.UnsafeGet[*float64](p_args, 7)
		var closest_unsafe = gd.UnsafeGet[*float64](p_args, 8)
		var info = gd.UnsafeGet[*PhysicsServer3DExtensionShapeRestInfo](p_args, 9)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape_rid, transform, motion, Float.X(margin), int(collision_mask), collide_with_bodies, collide_with_areas, closest_safe, closest_unsafe, info)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _collide_shape(impl func(ptr unsafe.Pointer, shape_rid Resource.ID, transform Transform3D.BasisOrigin, motion Vector3.XYZ, margin Float.X, collision_mask int, collide_with_bodies bool, collide_with_areas bool, results unsafe.Pointer, max_results int, result_count *int32) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape_rid = gd.UnsafeGet[gd.RID](p_args, 0)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 1)
		var motion = gd.UnsafeGet[gd.Vector3](p_args, 2)
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
func (Instance) _rest_info(impl func(ptr unsafe.Pointer, shape_rid Resource.ID, transform Transform3D.BasisOrigin, motion Vector3.XYZ, margin Float.X, collision_mask int, collide_with_bodies bool, collide_with_areas bool, rest_info *PhysicsServer3DExtensionShapeRestInfo) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape_rid = gd.UnsafeGet[gd.RID](p_args, 0)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 1)
		var motion = gd.UnsafeGet[gd.Vector3](p_args, 2)
		var margin = gd.UnsafeGet[gd.Float](p_args, 3)
		var collision_mask = gd.UnsafeGet[gd.Int](p_args, 4)
		var collide_with_bodies = gd.UnsafeGet[bool](p_args, 5)
		var collide_with_areas = gd.UnsafeGet[bool](p_args, 6)
		var rest_info = gd.UnsafeGet[*PhysicsServer3DExtensionShapeRestInfo](p_args, 7)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape_rid, transform, motion, Float.X(margin), int(collision_mask), collide_with_bodies, collide_with_areas, rest_info)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_closest_point_to_object_volume(impl func(ptr unsafe.Pointer, obj Resource.ID, point Vector3.XYZ) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var obj = gd.UnsafeGet[gd.RID](p_args, 0)
		var point = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj, point)
		gd.UnsafeSet(p_back, gd.Vector3(ret))
	}
}
func (self Instance) IsBodyExcludedFromQuery(body Resource.ID) bool {
	return bool(class(self).IsBodyExcludedFromQuery(body))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PhysicsDirectSpaceState3DExtension

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicsDirectSpaceState3DExtension"))
	casted := Instance{*(*gdclass.PhysicsDirectSpaceState3DExtension)(unsafe.Pointer(&object))}
	return casted
}

func (class) _intersect_ray(impl func(ptr unsafe.Pointer, from gd.Vector3, to gd.Vector3, collision_mask gd.Int, collide_with_bodies bool, collide_with_areas bool, hit_from_inside bool, hit_back_faces bool, pick_ray bool, result *PhysicsServer3DExtensionRayResult) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var from = gd.UnsafeGet[gd.Vector3](p_args, 0)
		var to = gd.UnsafeGet[gd.Vector3](p_args, 1)
		var collision_mask = gd.UnsafeGet[gd.Int](p_args, 2)
		var collide_with_bodies = gd.UnsafeGet[bool](p_args, 3)
		var collide_with_areas = gd.UnsafeGet[bool](p_args, 4)
		var hit_from_inside = gd.UnsafeGet[bool](p_args, 5)
		var hit_back_faces = gd.UnsafeGet[bool](p_args, 6)
		var pick_ray = gd.UnsafeGet[bool](p_args, 7)
		var result = gd.UnsafeGet[*PhysicsServer3DExtensionRayResult](p_args, 8)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from, to, collision_mask, collide_with_bodies, collide_with_areas, hit_from_inside, hit_back_faces, pick_ray, result)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _intersect_point(impl func(ptr unsafe.Pointer, position gd.Vector3, collision_mask gd.Int, collide_with_bodies bool, collide_with_areas bool, results *PhysicsServer3DExtensionShapeResult, max_results gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var position = gd.UnsafeGet[gd.Vector3](p_args, 0)
		var collision_mask = gd.UnsafeGet[gd.Int](p_args, 1)
		var collide_with_bodies = gd.UnsafeGet[bool](p_args, 2)
		var collide_with_areas = gd.UnsafeGet[bool](p_args, 3)
		var results = gd.UnsafeGet[*PhysicsServer3DExtensionShapeResult](p_args, 4)
		var max_results = gd.UnsafeGet[gd.Int](p_args, 5)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, position, collision_mask, collide_with_bodies, collide_with_areas, results, max_results)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _intersect_shape(impl func(ptr unsafe.Pointer, shape_rid gd.RID, transform gd.Transform3D, motion gd.Vector3, margin gd.Float, collision_mask gd.Int, collide_with_bodies bool, collide_with_areas bool, result_count *PhysicsServer3DExtensionShapeResult, max_results gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape_rid = gd.UnsafeGet[gd.RID](p_args, 0)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 1)
		var motion = gd.UnsafeGet[gd.Vector3](p_args, 2)
		var margin = gd.UnsafeGet[gd.Float](p_args, 3)
		var collision_mask = gd.UnsafeGet[gd.Int](p_args, 4)
		var collide_with_bodies = gd.UnsafeGet[bool](p_args, 5)
		var collide_with_areas = gd.UnsafeGet[bool](p_args, 6)
		var result_count = gd.UnsafeGet[*PhysicsServer3DExtensionShapeResult](p_args, 7)
		var max_results = gd.UnsafeGet[gd.Int](p_args, 8)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape_rid, transform, motion, margin, collision_mask, collide_with_bodies, collide_with_areas, result_count, max_results)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _cast_motion(impl func(ptr unsafe.Pointer, shape_rid gd.RID, transform gd.Transform3D, motion gd.Vector3, margin gd.Float, collision_mask gd.Int, collide_with_bodies bool, collide_with_areas bool, closest_safe *float64, closest_unsafe *float64, info *PhysicsServer3DExtensionShapeRestInfo) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape_rid = gd.UnsafeGet[gd.RID](p_args, 0)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 1)
		var motion = gd.UnsafeGet[gd.Vector3](p_args, 2)
		var margin = gd.UnsafeGet[gd.Float](p_args, 3)
		var collision_mask = gd.UnsafeGet[gd.Int](p_args, 4)
		var collide_with_bodies = gd.UnsafeGet[bool](p_args, 5)
		var collide_with_areas = gd.UnsafeGet[bool](p_args, 6)
		var closest_safe = gd.UnsafeGet[*float64](p_args, 7)
		var closest_unsafe = gd.UnsafeGet[*float64](p_args, 8)
		var info = gd.UnsafeGet[*PhysicsServer3DExtensionShapeRestInfo](p_args, 9)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape_rid, transform, motion, margin, collision_mask, collide_with_bodies, collide_with_areas, closest_safe, closest_unsafe, info)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _collide_shape(impl func(ptr unsafe.Pointer, shape_rid gd.RID, transform gd.Transform3D, motion gd.Vector3, margin gd.Float, collision_mask gd.Int, collide_with_bodies bool, collide_with_areas bool, results unsafe.Pointer, max_results gd.Int, result_count *int32) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape_rid = gd.UnsafeGet[gd.RID](p_args, 0)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 1)
		var motion = gd.UnsafeGet[gd.Vector3](p_args, 2)
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

func (class) _rest_info(impl func(ptr unsafe.Pointer, shape_rid gd.RID, transform gd.Transform3D, motion gd.Vector3, margin gd.Float, collision_mask gd.Int, collide_with_bodies bool, collide_with_areas bool, rest_info *PhysicsServer3DExtensionShapeRestInfo) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape_rid = gd.UnsafeGet[gd.RID](p_args, 0)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 1)
		var motion = gd.UnsafeGet[gd.Vector3](p_args, 2)
		var margin = gd.UnsafeGet[gd.Float](p_args, 3)
		var collision_mask = gd.UnsafeGet[gd.Int](p_args, 4)
		var collide_with_bodies = gd.UnsafeGet[bool](p_args, 5)
		var collide_with_areas = gd.UnsafeGet[bool](p_args, 6)
		var rest_info = gd.UnsafeGet[*PhysicsServer3DExtensionShapeRestInfo](p_args, 7)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape_rid, transform, motion, margin, collision_mask, collide_with_bodies, collide_with_areas, rest_info)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_closest_point_to_object_volume(impl func(ptr unsafe.Pointer, obj gd.RID, point gd.Vector3) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var obj = gd.UnsafeGet[gd.RID](p_args, 0)
		var point = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj, point)
		gd.UnsafeSet(p_back, ret)
	}
}

//go:nosplit
func (self class) IsBodyExcludedFromQuery(body gd.RID) bool {
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsDirectSpaceState3DExtension.Bind_is_body_excluded_from_query, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPhysicsDirectSpaceState3DExtension() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPhysicsDirectSpaceState3DExtension() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsPhysicsDirectSpaceState3D() PhysicsDirectSpaceState3D.Advanced {
	return *((*PhysicsDirectSpaceState3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPhysicsDirectSpaceState3D() PhysicsDirectSpaceState3D.Instance {
	return *((*PhysicsDirectSpaceState3D.Instance)(unsafe.Pointer(&self)))
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
	case "_get_closest_point_to_object_volume":
		return reflect.ValueOf(self._get_closest_point_to_object_volume)
	default:
		return gd.VirtualByName(PhysicsDirectSpaceState3D.Advanced(self.AsPhysicsDirectSpaceState3D()), name)
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
	case "_get_closest_point_to_object_volume":
		return reflect.ValueOf(self._get_closest_point_to_object_volume)
	default:
		return gd.VirtualByName(PhysicsDirectSpaceState3D.Instance(self.AsPhysicsDirectSpaceState3D()), name)
	}
}
func init() {
	gdclass.Register("PhysicsDirectSpaceState3DExtension", func(ptr gd.Object) any {
		return [1]gdclass.PhysicsDirectSpaceState3DExtension{*(*gdclass.PhysicsDirectSpaceState3DExtension)(unsafe.Pointer(&ptr))}
	})
}
