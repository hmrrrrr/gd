// Package Geometry3D provides methods for working with Geometry3D object instances.
package Geometry3D

import "unsafe"
import "sync"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/Plane"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Vector3"

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
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
Provides a set of helper functions to create geometric shapes, compute intersections between shapes, and process various other geometric operations in 3D.
*/
var self [1]gdclass.Geometry3D
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.Geometry3D)
	self = *(*[1]gdclass.Geometry3D)(unsafe.Pointer(&obj))
}

/*
Calculates and returns all the vertex points of a convex shape defined by an array of [param planes].
*/
func ComputeConvexMeshPoints(planes []Plane.NormalD) []Vector3.XYZ { //gd:Geometry3D.compute_convex_mesh_points
	once.Do(singleton)
	return []Vector3.XYZ(slices.Collect(class(self).ComputeConvexMeshPoints(gd.ArrayFromSlice[Array.Contains[Plane.NormalD]](planes)).Values()))
}

/*
Returns an array with 6 [Plane]s that describe the sides of a box centered at the origin. The box size is defined by [param extents], which represents one (positive) corner of the box (i.e. half its actual size).
*/
func BuildBoxPlanes(extents Vector3.XYZ) []Plane.NormalD { //gd:Geometry3D.build_box_planes
	once.Do(singleton)
	return []Plane.NormalD(gd.ArrayAs[[]Plane.NormalD](gd.InternalArray(class(self).BuildBoxPlanes(Vector3.XYZ(extents)))))
}

/*
Returns an array of [Plane]s closely bounding a faceted cylinder centered at the origin with radius [param radius] and height [param height]. The parameter [param sides] defines how many planes will be generated for the round part of the cylinder. The parameter [param axis] describes the axis along which the cylinder is oriented (0 for X, 1 for Y, 2 for Z).
*/
func BuildCylinderPlanes(radius Float.X, height Float.X, sides int) []Plane.NormalD { //gd:Geometry3D.build_cylinder_planes
	once.Do(singleton)
	return []Plane.NormalD(gd.ArrayAs[[]Plane.NormalD](gd.InternalArray(class(self).BuildCylinderPlanes(float64(radius), float64(height), int64(sides), 2))))
}

/*
Returns an array of [Plane]s closely bounding a faceted capsule centered at the origin with radius [param radius] and height [param height]. The parameter [param sides] defines how many planes will be generated for the side part of the capsule, whereas [param lats] gives the number of latitudinal steps at the bottom and top of the capsule. The parameter [param axis] describes the axis along which the capsule is oriented (0 for X, 1 for Y, 2 for Z).
*/
func BuildCapsulePlanes(radius Float.X, height Float.X, sides int, lats int) []Plane.NormalD { //gd:Geometry3D.build_capsule_planes
	once.Do(singleton)
	return []Plane.NormalD(gd.ArrayAs[[]Plane.NormalD](gd.InternalArray(class(self).BuildCapsulePlanes(float64(radius), float64(height), int64(sides), int64(lats), 2))))
}

/*
Given the two 3D segments ([param p1], [param p2]) and ([param q1], [param q2]), finds those two points on the two segments that are closest to each other. Returns a [PackedVector3Array] that contains this point on ([param p1], [param p2]) as well the accompanying point on ([param q1], [param q2]).
*/
func GetClosestPointsBetweenSegments(p1 Vector3.XYZ, p2 Vector3.XYZ, q1 Vector3.XYZ, q2 Vector3.XYZ) []Vector3.XYZ { //gd:Geometry3D.get_closest_points_between_segments
	once.Do(singleton)
	return []Vector3.XYZ(slices.Collect(class(self).GetClosestPointsBetweenSegments(Vector3.XYZ(p1), Vector3.XYZ(p2), Vector3.XYZ(q1), Vector3.XYZ(q2)).Values()))
}

/*
Returns the 3D point on the 3D segment ([param s1], [param s2]) that is closest to [param point]. The returned point will always be inside the specified segment.
*/
func GetClosestPointToSegment(point Vector3.XYZ, s1 Vector3.XYZ, s2 Vector3.XYZ) Vector3.XYZ { //gd:Geometry3D.get_closest_point_to_segment
	once.Do(singleton)
	return Vector3.XYZ(class(self).GetClosestPointToSegment(Vector3.XYZ(point), Vector3.XYZ(s1), Vector3.XYZ(s2)))
}

/*
Returns the 3D point on the 3D line defined by ([param s1], [param s2]) that is closest to [param point]. The returned point can be inside the segment ([param s1], [param s2]) or outside of it, i.e. somewhere on the line extending from the segment.
*/
func GetClosestPointToSegmentUncapped(point Vector3.XYZ, s1 Vector3.XYZ, s2 Vector3.XYZ) Vector3.XYZ { //gd:Geometry3D.get_closest_point_to_segment_uncapped
	once.Do(singleton)
	return Vector3.XYZ(class(self).GetClosestPointToSegmentUncapped(Vector3.XYZ(point), Vector3.XYZ(s1), Vector3.XYZ(s2)))
}

/*
Returns a [Vector3] containing weights based on how close a 3D position ([param point]) is to a triangle's different vertices ([param a], [param b] and [param c]). This is useful for interpolating between the data of different vertices in a triangle. One example use case is using this to smoothly rotate over a mesh instead of relying solely on face normals.
[url=https://en.wikipedia.org/wiki/Barycentric_coordinate_system]Here is a more detailed explanation of barycentric coordinates.[/url]
*/
func GetTriangleBarycentricCoords(point Vector3.XYZ, a Vector3.XYZ, b Vector3.XYZ, c Vector3.XYZ) Vector3.XYZ { //gd:Geometry3D.get_triangle_barycentric_coords
	once.Do(singleton)
	return Vector3.XYZ(class(self).GetTriangleBarycentricCoords(Vector3.XYZ(point), Vector3.XYZ(a), Vector3.XYZ(b), Vector3.XYZ(c)))
}

/*
Tests if the 3D ray starting at [param from] with the direction of [param dir] intersects the triangle specified by [param a], [param b] and [param c]. If yes, returns the point of intersection as [Vector3]. If no intersection takes place, returns [code]null[/code].
*/
func RayIntersectsTriangle(from Vector3.XYZ, dir Vector3.XYZ, a Vector3.XYZ, b Vector3.XYZ, c Vector3.XYZ) any { //gd:Geometry3D.ray_intersects_triangle
	once.Do(singleton)
	return any(class(self).RayIntersectsTriangle(Vector3.XYZ(from), Vector3.XYZ(dir), Vector3.XYZ(a), Vector3.XYZ(b), Vector3.XYZ(c)).Interface())
}

/*
Tests if the segment ([param from], [param to]) intersects the triangle [param a], [param b], [param c]. If yes, returns the point of intersection as [Vector3]. If no intersection takes place, returns [code]null[/code].
*/
func SegmentIntersectsTriangle(from Vector3.XYZ, to Vector3.XYZ, a Vector3.XYZ, b Vector3.XYZ, c Vector3.XYZ) any { //gd:Geometry3D.segment_intersects_triangle
	once.Do(singleton)
	return any(class(self).SegmentIntersectsTriangle(Vector3.XYZ(from), Vector3.XYZ(to), Vector3.XYZ(a), Vector3.XYZ(b), Vector3.XYZ(c)).Interface())
}

/*
Checks if the segment ([param from], [param to]) intersects the sphere that is located at [param sphere_position] and has radius [param sphere_radius]. If no, returns an empty [PackedVector3Array]. If yes, returns a [PackedVector3Array] containing the point of intersection and the sphere's normal at the point of intersection.
*/
func SegmentIntersectsSphere(from Vector3.XYZ, to Vector3.XYZ, sphere_position Vector3.XYZ, sphere_radius Float.X) []Vector3.XYZ { //gd:Geometry3D.segment_intersects_sphere
	once.Do(singleton)
	return []Vector3.XYZ(slices.Collect(class(self).SegmentIntersectsSphere(Vector3.XYZ(from), Vector3.XYZ(to), Vector3.XYZ(sphere_position), float64(sphere_radius)).Values()))
}

/*
Checks if the segment ([param from], [param to]) intersects the cylinder with height [param height] that is centered at the origin and has radius [param radius]. If no, returns an empty [PackedVector3Array]. If an intersection takes place, the returned array contains the point of intersection and the cylinder's normal at the point of intersection.
*/
func SegmentIntersectsCylinder(from Vector3.XYZ, to Vector3.XYZ, height Float.X, radius Float.X) []Vector3.XYZ { //gd:Geometry3D.segment_intersects_cylinder
	once.Do(singleton)
	return []Vector3.XYZ(slices.Collect(class(self).SegmentIntersectsCylinder(Vector3.XYZ(from), Vector3.XYZ(to), float64(height), float64(radius)).Values()))
}

/*
Given a convex hull defined though the [Plane]s in the array [param planes], tests if the segment ([param from], [param to]) intersects with that hull. If an intersection is found, returns a [PackedVector3Array] containing the point the intersection and the hull's normal. Otherwise, returns an empty array.
*/
func SegmentIntersectsConvex(from Vector3.XYZ, to Vector3.XYZ, planes []Plane.NormalD) []Vector3.XYZ { //gd:Geometry3D.segment_intersects_convex
	once.Do(singleton)
	return []Vector3.XYZ(slices.Collect(class(self).SegmentIntersectsConvex(Vector3.XYZ(from), Vector3.XYZ(to), gd.ArrayFromSlice[Array.Contains[Plane.NormalD]](planes)).Values()))
}

/*
Clips the polygon defined by the points in [param points] against the [param plane] and returns the points of the clipped polygon.
*/
func ClipPolygon(points []Vector3.XYZ, plane Plane.NormalD) []Vector3.XYZ { //gd:Geometry3D.clip_polygon
	once.Do(singleton)
	return []Vector3.XYZ(slices.Collect(class(self).ClipPolygon(Packed.New(points...), Plane.NormalD(plane)).Values()))
}

/*
Tetrahedralizes the volume specified by a discrete set of [param points] in 3D space, ensuring that no point lies within the circumsphere of any resulting tetrahedron. The method returns a [PackedInt32Array] where each tetrahedron consists of four consecutive point indices into the [param points] array (resulting in an array with [code]n * 4[/code] elements, where [code]n[/code] is the number of tetrahedra found). If the tetrahedralization is unsuccessful, an empty [PackedInt32Array] is returned.
*/
func TetrahedralizeDelaunay(points []Vector3.XYZ) []int32 { //gd:Geometry3D.tetrahedralize_delaunay
	once.Do(singleton)
	return []int32(slices.Collect(class(self).TetrahedralizeDelaunay(Packed.New(points...)).Values()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]gdclass.Geometry3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Calculates and returns all the vertex points of a convex shape defined by an array of [param planes].
*/
//go:nosplit
func (self class) ComputeConvexMeshPoints(planes Array.Contains[Plane.NormalD]) Packed.Array[Vector3.XYZ] { //gd:Geometry3D.compute_convex_mesh_points
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(planes)))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry3D.Bind_compute_convex_mesh_points, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[Vector3.XYZ](Array.Through(gd.PackedProxy[gd.PackedVector3Array, Vector3.XYZ]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns an array with 6 [Plane]s that describe the sides of a box centered at the origin. The box size is defined by [param extents], which represents one (positive) corner of the box (i.e. half its actual size).
*/
//go:nosplit
func (self class) BuildBoxPlanes(extents Vector3.XYZ) Array.Contains[Plane.NormalD] { //gd:Geometry3D.build_box_planes
	var frame = callframe.New()
	callframe.Arg(frame, extents)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry3D.Bind_build_box_planes, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Plane.NormalD]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns an array of [Plane]s closely bounding a faceted cylinder centered at the origin with radius [param radius] and height [param height]. The parameter [param sides] defines how many planes will be generated for the round part of the cylinder. The parameter [param axis] describes the axis along which the cylinder is oriented (0 for X, 1 for Y, 2 for Z).
*/
//go:nosplit
func (self class) BuildCylinderPlanes(radius float64, height float64, sides int64, axis gd.Vector3Axis) Array.Contains[Plane.NormalD] { //gd:Geometry3D.build_cylinder_planes
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	callframe.Arg(frame, height)
	callframe.Arg(frame, sides)
	callframe.Arg(frame, axis)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry3D.Bind_build_cylinder_planes, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Plane.NormalD]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns an array of [Plane]s closely bounding a faceted capsule centered at the origin with radius [param radius] and height [param height]. The parameter [param sides] defines how many planes will be generated for the side part of the capsule, whereas [param lats] gives the number of latitudinal steps at the bottom and top of the capsule. The parameter [param axis] describes the axis along which the capsule is oriented (0 for X, 1 for Y, 2 for Z).
*/
//go:nosplit
func (self class) BuildCapsulePlanes(radius float64, height float64, sides int64, lats int64, axis gd.Vector3Axis) Array.Contains[Plane.NormalD] { //gd:Geometry3D.build_capsule_planes
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	callframe.Arg(frame, height)
	callframe.Arg(frame, sides)
	callframe.Arg(frame, lats)
	callframe.Arg(frame, axis)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry3D.Bind_build_capsule_planes, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Plane.NormalD]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Given the two 3D segments ([param p1], [param p2]) and ([param q1], [param q2]), finds those two points on the two segments that are closest to each other. Returns a [PackedVector3Array] that contains this point on ([param p1], [param p2]) as well the accompanying point on ([param q1], [param q2]).
*/
//go:nosplit
func (self class) GetClosestPointsBetweenSegments(p1 Vector3.XYZ, p2 Vector3.XYZ, q1 Vector3.XYZ, q2 Vector3.XYZ) Packed.Array[Vector3.XYZ] { //gd:Geometry3D.get_closest_points_between_segments
	var frame = callframe.New()
	callframe.Arg(frame, p1)
	callframe.Arg(frame, p2)
	callframe.Arg(frame, q1)
	callframe.Arg(frame, q2)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry3D.Bind_get_closest_points_between_segments, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[Vector3.XYZ](Array.Through(gd.PackedProxy[gd.PackedVector3Array, Vector3.XYZ]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the 3D point on the 3D segment ([param s1], [param s2]) that is closest to [param point]. The returned point will always be inside the specified segment.
*/
//go:nosplit
func (self class) GetClosestPointToSegment(point Vector3.XYZ, s1 Vector3.XYZ, s2 Vector3.XYZ) Vector3.XYZ { //gd:Geometry3D.get_closest_point_to_segment
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, s1)
	callframe.Arg(frame, s2)
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry3D.Bind_get_closest_point_to_segment, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the 3D point on the 3D line defined by ([param s1], [param s2]) that is closest to [param point]. The returned point can be inside the segment ([param s1], [param s2]) or outside of it, i.e. somewhere on the line extending from the segment.
*/
//go:nosplit
func (self class) GetClosestPointToSegmentUncapped(point Vector3.XYZ, s1 Vector3.XYZ, s2 Vector3.XYZ) Vector3.XYZ { //gd:Geometry3D.get_closest_point_to_segment_uncapped
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, s1)
	callframe.Arg(frame, s2)
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry3D.Bind_get_closest_point_to_segment_uncapped, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a [Vector3] containing weights based on how close a 3D position ([param point]) is to a triangle's different vertices ([param a], [param b] and [param c]). This is useful for interpolating between the data of different vertices in a triangle. One example use case is using this to smoothly rotate over a mesh instead of relying solely on face normals.
[url=https://en.wikipedia.org/wiki/Barycentric_coordinate_system]Here is a more detailed explanation of barycentric coordinates.[/url]
*/
//go:nosplit
func (self class) GetTriangleBarycentricCoords(point Vector3.XYZ, a Vector3.XYZ, b Vector3.XYZ, c Vector3.XYZ) Vector3.XYZ { //gd:Geometry3D.get_triangle_barycentric_coords
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, a)
	callframe.Arg(frame, b)
	callframe.Arg(frame, c)
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry3D.Bind_get_triangle_barycentric_coords, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Tests if the 3D ray starting at [param from] with the direction of [param dir] intersects the triangle specified by [param a], [param b] and [param c]. If yes, returns the point of intersection as [Vector3]. If no intersection takes place, returns [code]null[/code].
*/
//go:nosplit
func (self class) RayIntersectsTriangle(from Vector3.XYZ, dir Vector3.XYZ, a Vector3.XYZ, b Vector3.XYZ, c Vector3.XYZ) variant.Any { //gd:Geometry3D.ray_intersects_triangle
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, dir)
	callframe.Arg(frame, a)
	callframe.Arg(frame, b)
	callframe.Arg(frame, c)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry3D.Bind_ray_intersects_triangle, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Implementation(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Tests if the segment ([param from], [param to]) intersects the triangle [param a], [param b], [param c]. If yes, returns the point of intersection as [Vector3]. If no intersection takes place, returns [code]null[/code].
*/
//go:nosplit
func (self class) SegmentIntersectsTriangle(from Vector3.XYZ, to Vector3.XYZ, a Vector3.XYZ, b Vector3.XYZ, c Vector3.XYZ) variant.Any { //gd:Geometry3D.segment_intersects_triangle
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	callframe.Arg(frame, a)
	callframe.Arg(frame, b)
	callframe.Arg(frame, c)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry3D.Bind_segment_intersects_triangle, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Implementation(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Checks if the segment ([param from], [param to]) intersects the sphere that is located at [param sphere_position] and has radius [param sphere_radius]. If no, returns an empty [PackedVector3Array]. If yes, returns a [PackedVector3Array] containing the point of intersection and the sphere's normal at the point of intersection.
*/
//go:nosplit
func (self class) SegmentIntersectsSphere(from Vector3.XYZ, to Vector3.XYZ, sphere_position Vector3.XYZ, sphere_radius float64) Packed.Array[Vector3.XYZ] { //gd:Geometry3D.segment_intersects_sphere
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	callframe.Arg(frame, sphere_position)
	callframe.Arg(frame, sphere_radius)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry3D.Bind_segment_intersects_sphere, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[Vector3.XYZ](Array.Through(gd.PackedProxy[gd.PackedVector3Array, Vector3.XYZ]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Checks if the segment ([param from], [param to]) intersects the cylinder with height [param height] that is centered at the origin and has radius [param radius]. If no, returns an empty [PackedVector3Array]. If an intersection takes place, the returned array contains the point of intersection and the cylinder's normal at the point of intersection.
*/
//go:nosplit
func (self class) SegmentIntersectsCylinder(from Vector3.XYZ, to Vector3.XYZ, height float64, radius float64) Packed.Array[Vector3.XYZ] { //gd:Geometry3D.segment_intersects_cylinder
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	callframe.Arg(frame, height)
	callframe.Arg(frame, radius)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry3D.Bind_segment_intersects_cylinder, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[Vector3.XYZ](Array.Through(gd.PackedProxy[gd.PackedVector3Array, Vector3.XYZ]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Given a convex hull defined though the [Plane]s in the array [param planes], tests if the segment ([param from], [param to]) intersects with that hull. If an intersection is found, returns a [PackedVector3Array] containing the point the intersection and the hull's normal. Otherwise, returns an empty array.
*/
//go:nosplit
func (self class) SegmentIntersectsConvex(from Vector3.XYZ, to Vector3.XYZ, planes Array.Contains[Plane.NormalD]) Packed.Array[Vector3.XYZ] { //gd:Geometry3D.segment_intersects_convex
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	callframe.Arg(frame, pointers.Get(gd.InternalArray(planes)))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry3D.Bind_segment_intersects_convex, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[Vector3.XYZ](Array.Through(gd.PackedProxy[gd.PackedVector3Array, Vector3.XYZ]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Clips the polygon defined by the points in [param points] against the [param plane] and returns the points of the clipped polygon.
*/
//go:nosplit
func (self class) ClipPolygon(points Packed.Array[Vector3.XYZ], plane Plane.NormalD) Packed.Array[Vector3.XYZ] { //gd:Geometry3D.clip_polygon
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedVector3Array, Vector3.XYZ](points)))
	callframe.Arg(frame, plane)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry3D.Bind_clip_polygon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[Vector3.XYZ](Array.Through(gd.PackedProxy[gd.PackedVector3Array, Vector3.XYZ]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Tetrahedralizes the volume specified by a discrete set of [param points] in 3D space, ensuring that no point lies within the circumsphere of any resulting tetrahedron. The method returns a [PackedInt32Array] where each tetrahedron consists of four consecutive point indices into the [param points] array (resulting in an array with [code]n * 4[/code] elements, where [code]n[/code] is the number of tetrahedra found). If the tetrahedralization is unsuccessful, an empty [PackedInt32Array] is returned.
*/
//go:nosplit
func (self class) TetrahedralizeDelaunay(points Packed.Array[Vector3.XYZ]) Packed.Array[int32] { //gd:Geometry3D.tetrahedralize_delaunay
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedVector3Array, Vector3.XYZ](points)))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry3D.Bind_tetrahedralize_delaunay, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int32](Array.Through(gd.PackedProxy[gd.PackedInt32Array, int32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("Geometry3D", func(ptr gd.Object) any { return [1]gdclass.Geometry3D{*(*gdclass.Geometry3D)(unsafe.Pointer(&ptr))} })
}
