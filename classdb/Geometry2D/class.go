// Package Geometry2D provides methods for working with Geometry2D object instances.
package Geometry2D

import "unsafe"
import "sync"
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
import "graphics.gd/variant/Vector2"
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
Provides a set of helper functions to create geometric shapes, compute intersections between shapes, and process various other geometric operations in 2D.
*/
var self [1]gdclass.Geometry2D
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.Geometry2D)
	self = *(*[1]gdclass.Geometry2D)(unsafe.Pointer(&obj))
}

/*
Returns [code]true[/code] if [param point] is inside the circle or if it's located exactly [i]on[/i] the circle's boundary, otherwise returns [code]false[/code].
*/
func IsPointInCircle(point Vector2.XY, circle_position Vector2.XY, circle_radius Float.X) bool { //gd:Geometry2D.is_point_in_circle
	once.Do(singleton)
	return bool(class(self).IsPointInCircle(gd.Vector2(point), gd.Vector2(circle_position), gd.Float(circle_radius)))
}

/*
Given the 2D segment ([param segment_from], [param segment_to]), returns the position on the segment (as a number between 0 and 1) at which the segment hits the circle that is located at position [param circle_position] and has radius [param circle_radius]. If the segment does not intersect the circle, -1 is returned (this is also the case if the line extending the segment would intersect the circle, but the segment does not).
*/
func SegmentIntersectsCircle(segment_from Vector2.XY, segment_to Vector2.XY, circle_position Vector2.XY, circle_radius Float.X) Float.X { //gd:Geometry2D.segment_intersects_circle
	once.Do(singleton)
	return Float.X(Float.X(class(self).SegmentIntersectsCircle(gd.Vector2(segment_from), gd.Vector2(segment_to), gd.Vector2(circle_position), gd.Float(circle_radius))))
}

/*
Checks if the two segments ([param from_a], [param to_a]) and ([param from_b], [param to_b]) intersect. If yes, return the point of intersection as [Vector2]. If no intersection takes place, returns [code]null[/code].
*/
func SegmentIntersectsSegment(from_a Vector2.XY, to_a Vector2.XY, from_b Vector2.XY, to_b Vector2.XY) any { //gd:Geometry2D.segment_intersects_segment
	once.Do(singleton)
	return any(class(self).SegmentIntersectsSegment(gd.Vector2(from_a), gd.Vector2(to_a), gd.Vector2(from_b), gd.Vector2(to_b)).Interface())
}

/*
Checks if the two lines ([param from_a], [param dir_a]) and ([param from_b], [param dir_b]) intersect. If yes, return the point of intersection as [Vector2]. If no intersection takes place, returns [code]null[/code].
[b]Note:[/b] The lines are specified using direction vectors, not end points.
*/
func LineIntersectsLine(from_a Vector2.XY, dir_a Vector2.XY, from_b Vector2.XY, dir_b Vector2.XY) any { //gd:Geometry2D.line_intersects_line
	once.Do(singleton)
	return any(class(self).LineIntersectsLine(gd.Vector2(from_a), gd.Vector2(dir_a), gd.Vector2(from_b), gd.Vector2(dir_b)).Interface())
}

/*
Given the two 2D segments ([param p1], [param q1]) and ([param p2], [param q2]), finds those two points on the two segments that are closest to each other. Returns a [PackedVector2Array] that contains this point on ([param p1], [param q1]) as well the accompanying point on ([param p2], [param q2]).
*/
func GetClosestPointsBetweenSegments(p1 Vector2.XY, q1 Vector2.XY, p2 Vector2.XY, q2 Vector2.XY) []Vector2.XY { //gd:Geometry2D.get_closest_points_between_segments
	once.Do(singleton)
	return []Vector2.XY(slices.Collect(class(self).GetClosestPointsBetweenSegments(gd.Vector2(p1), gd.Vector2(q1), gd.Vector2(p2), gd.Vector2(q2)).Values()))
}

/*
Returns the 2D point on the 2D segment ([param s1], [param s2]) that is closest to [param point]. The returned point will always be inside the specified segment.
*/
func GetClosestPointToSegment(point Vector2.XY, s1 Vector2.XY, s2 Vector2.XY) Vector2.XY { //gd:Geometry2D.get_closest_point_to_segment
	once.Do(singleton)
	return Vector2.XY(class(self).GetClosestPointToSegment(gd.Vector2(point), gd.Vector2(s1), gd.Vector2(s2)))
}

/*
Returns the 2D point on the 2D line defined by ([param s1], [param s2]) that is closest to [param point]. The returned point can be inside the segment ([param s1], [param s2]) or outside of it, i.e. somewhere on the line extending from the segment.
*/
func GetClosestPointToSegmentUncapped(point Vector2.XY, s1 Vector2.XY, s2 Vector2.XY) Vector2.XY { //gd:Geometry2D.get_closest_point_to_segment_uncapped
	once.Do(singleton)
	return Vector2.XY(class(self).GetClosestPointToSegmentUncapped(gd.Vector2(point), gd.Vector2(s1), gd.Vector2(s2)))
}

/*
Returns if [param point] is inside the triangle specified by [param a], [param b] and [param c].
*/
func PointIsInsideTriangle(point Vector2.XY, a Vector2.XY, b Vector2.XY, c Vector2.XY) bool { //gd:Geometry2D.point_is_inside_triangle
	once.Do(singleton)
	return bool(class(self).PointIsInsideTriangle(gd.Vector2(point), gd.Vector2(a), gd.Vector2(b), gd.Vector2(c)))
}

/*
Returns [code]true[/code] if [param polygon]'s vertices are ordered in clockwise order, otherwise returns [code]false[/code].
[b]Note:[/b] Assumes a Cartesian coordinate system where [code]+x[/code] is right and [code]+y[/code] is up. If using screen coordinates ([code]+y[/code] is down), the result will need to be flipped (i.e. a [code]true[/code] result will indicate counter-clockwise).
*/
func IsPolygonClockwise(polygon []Vector2.XY) bool { //gd:Geometry2D.is_polygon_clockwise
	once.Do(singleton)
	return bool(class(self).IsPolygonClockwise(Packed.New(polygon...)))
}

/*
Returns [code]true[/code] if [param point] is inside [param polygon] or if it's located exactly [i]on[/i] polygon's boundary, otherwise returns [code]false[/code].
*/
func IsPointInPolygon(point Vector2.XY, polygon []Vector2.XY) bool { //gd:Geometry2D.is_point_in_polygon
	once.Do(singleton)
	return bool(class(self).IsPointInPolygon(gd.Vector2(point), Packed.New(polygon...)))
}

/*
Triangulates the polygon specified by the points in [param polygon]. Returns a [PackedInt32Array] where each triangle consists of three consecutive point indices into [param polygon] (i.e. the returned array will have [code]n * 3[/code] elements, with [code]n[/code] being the number of found triangles). Output triangles will always be counter clockwise, and the contour will be flipped if it's clockwise. If the triangulation did not succeed, an empty [PackedInt32Array] is returned.
*/
func TriangulatePolygon(polygon []Vector2.XY) []int32 { //gd:Geometry2D.triangulate_polygon
	once.Do(singleton)
	return []int32(slices.Collect(class(self).TriangulatePolygon(Packed.New(polygon...)).Values()))
}

/*
Triangulates the area specified by discrete set of [param points] such that no point is inside the circumcircle of any resulting triangle. Returns a [PackedInt32Array] where each triangle consists of three consecutive point indices into [param points] (i.e. the returned array will have [code]n * 3[/code] elements, with [code]n[/code] being the number of found triangles). If the triangulation did not succeed, an empty [PackedInt32Array] is returned.
*/
func TriangulateDelaunay(points []Vector2.XY) []int32 { //gd:Geometry2D.triangulate_delaunay
	once.Do(singleton)
	return []int32(slices.Collect(class(self).TriangulateDelaunay(Packed.New(points...)).Values()))
}

/*
Given an array of [Vector2]s, returns the convex hull as a list of points in counterclockwise order. The last point is the same as the first one.
*/
func ConvexHull(points []Vector2.XY) []Vector2.XY { //gd:Geometry2D.convex_hull
	once.Do(singleton)
	return []Vector2.XY(slices.Collect(class(self).ConvexHull(Packed.New(points...)).Values()))
}

/*
Decomposes the [param polygon] into multiple convex hulls and returns an array of [PackedVector2Array].
*/
func DecomposePolygonInConvex(polygon []Vector2.XY) [][]Vector2.XY { //gd:Geometry2D.decompose_polygon_in_convex
	once.Do(singleton)
	return [][]Vector2.XY(gd.ArrayAs[[][]Vector2.XY](gd.InternalArray(class(self).DecomposePolygonInConvex(Packed.New(polygon...)))))
}

/*
Merges (combines) [param polygon_a] and [param polygon_b] and returns an array of merged polygons. This performs [constant OPERATION_UNION] between polygons.
The operation may result in an outer polygon (boundary) and multiple inner polygons (holes) produced which could be distinguished by calling [method is_polygon_clockwise].
*/
func MergePolygons(polygon_a []Vector2.XY, polygon_b []Vector2.XY) [][]Vector2.XY { //gd:Geometry2D.merge_polygons
	once.Do(singleton)
	return [][]Vector2.XY(gd.ArrayAs[[][]Vector2.XY](gd.InternalArray(class(self).MergePolygons(Packed.New(polygon_a...), Packed.New(polygon_b...)))))
}

/*
Clips [param polygon_a] against [param polygon_b] and returns an array of clipped polygons. This performs [constant OPERATION_DIFFERENCE] between polygons. Returns an empty array if [param polygon_b] completely overlaps [param polygon_a].
If [param polygon_b] is enclosed by [param polygon_a], returns an outer polygon (boundary) and inner polygon (hole) which could be distinguished by calling [method is_polygon_clockwise].
*/
func ClipPolygons(polygon_a []Vector2.XY, polygon_b []Vector2.XY) [][]Vector2.XY { //gd:Geometry2D.clip_polygons
	once.Do(singleton)
	return [][]Vector2.XY(gd.ArrayAs[[][]Vector2.XY](gd.InternalArray(class(self).ClipPolygons(Packed.New(polygon_a...), Packed.New(polygon_b...)))))
}

/*
Intersects [param polygon_a] with [param polygon_b] and returns an array of intersected polygons. This performs [constant OPERATION_INTERSECTION] between polygons. In other words, returns common area shared by polygons. Returns an empty array if no intersection occurs.
The operation may result in an outer polygon (boundary) and inner polygon (hole) produced which could be distinguished by calling [method is_polygon_clockwise].
*/
func IntersectPolygons(polygon_a []Vector2.XY, polygon_b []Vector2.XY) [][]Vector2.XY { //gd:Geometry2D.intersect_polygons
	once.Do(singleton)
	return [][]Vector2.XY(gd.ArrayAs[[][]Vector2.XY](gd.InternalArray(class(self).IntersectPolygons(Packed.New(polygon_a...), Packed.New(polygon_b...)))))
}

/*
Mutually excludes common area defined by intersection of [param polygon_a] and [param polygon_b] (see [method intersect_polygons]) and returns an array of excluded polygons. This performs [constant OPERATION_XOR] between polygons. In other words, returns all but common area between polygons.
The operation may result in an outer polygon (boundary) and inner polygon (hole) produced which could be distinguished by calling [method is_polygon_clockwise].
*/
func ExcludePolygons(polygon_a []Vector2.XY, polygon_b []Vector2.XY) [][]Vector2.XY { //gd:Geometry2D.exclude_polygons
	once.Do(singleton)
	return [][]Vector2.XY(gd.ArrayAs[[][]Vector2.XY](gd.InternalArray(class(self).ExcludePolygons(Packed.New(polygon_a...), Packed.New(polygon_b...)))))
}

/*
Clips [param polyline] against [param polygon] and returns an array of clipped polylines. This performs [constant OPERATION_DIFFERENCE] between the polyline and the polygon. This operation can be thought of as cutting a line with a closed shape.
*/
func ClipPolylineWithPolygon(polyline []Vector2.XY, polygon []Vector2.XY) [][]Vector2.XY { //gd:Geometry2D.clip_polyline_with_polygon
	once.Do(singleton)
	return [][]Vector2.XY(gd.ArrayAs[[][]Vector2.XY](gd.InternalArray(class(self).ClipPolylineWithPolygon(Packed.New(polyline...), Packed.New(polygon...)))))
}

/*
Intersects [param polyline] with [param polygon] and returns an array of intersected polylines. This performs [constant OPERATION_INTERSECTION] between the polyline and the polygon. This operation can be thought of as chopping a line with a closed shape.
*/
func IntersectPolylineWithPolygon(polyline []Vector2.XY, polygon []Vector2.XY) [][]Vector2.XY { //gd:Geometry2D.intersect_polyline_with_polygon
	once.Do(singleton)
	return [][]Vector2.XY(gd.ArrayAs[[][]Vector2.XY](gd.InternalArray(class(self).IntersectPolylineWithPolygon(Packed.New(polyline...), Packed.New(polygon...)))))
}

/*
Inflates or deflates [param polygon] by [param delta] units (pixels). If [param delta] is positive, makes the polygon grow outward. If [param delta] is negative, shrinks the polygon inward. Returns an array of polygons because inflating/deflating may result in multiple discrete polygons. Returns an empty array if [param delta] is negative and the absolute value of it approximately exceeds the minimum bounding rectangle dimensions of the polygon.
Each polygon's vertices will be rounded as determined by [param join_type], see [enum PolyJoinType].
The operation may result in an outer polygon (boundary) and inner polygon (hole) produced which could be distinguished by calling [method is_polygon_clockwise].
[b]Note:[/b] To translate the polygon's vertices specifically, multiply them to a [Transform2D]:
[codeblocks]
[gdscript]
var polygon = PackedVector2Array([Vector2(0, 0), Vector2(100, 0), Vector2(100, 100), Vector2(0, 100)])
var offset = Vector2(50, 50)
polygon = Transform2D(0, offset) * polygon
print(polygon) # prints [(50, 50), (150, 50), (150, 150), (50, 150)]
[/gdscript]
[csharp]
var polygon = new Vector2[] { new Vector2(0, 0), new Vector2(100, 0), new Vector2(100, 100), new Vector2(0, 100) };
var offset = new Vector2(50, 50);
polygon = new Transform2D(0, offset) * polygon;
GD.Print((Variant)polygon); // prints [(50, 50), (150, 50), (150, 150), (50, 150)]
[/csharp]
[/codeblocks]
*/
func OffsetPolygon(polygon []Vector2.XY, delta Float.X) [][]Vector2.XY { //gd:Geometry2D.offset_polygon
	once.Do(singleton)
	return [][]Vector2.XY(gd.ArrayAs[[][]Vector2.XY](gd.InternalArray(class(self).OffsetPolygon(Packed.New(polygon...), gd.Float(delta), 0))))
}

/*
Inflates or deflates [param polyline] by [param delta] units (pixels), producing polygons. If [param delta] is positive, makes the polyline grow outward. Returns an array of polygons because inflating/deflating may result in multiple discrete polygons. If [param delta] is negative, returns an empty array.
Each polygon's vertices will be rounded as determined by [param join_type], see [enum PolyJoinType].
Each polygon's endpoints will be rounded as determined by [param end_type], see [enum PolyEndType].
The operation may result in an outer polygon (boundary) and inner polygon (hole) produced which could be distinguished by calling [method is_polygon_clockwise].
*/
func OffsetPolyline(polyline []Vector2.XY, delta Float.X) [][]Vector2.XY { //gd:Geometry2D.offset_polyline
	once.Do(singleton)
	return [][]Vector2.XY(gd.ArrayAs[[][]Vector2.XY](gd.InternalArray(class(self).OffsetPolyline(Packed.New(polyline...), gd.Float(delta), 0, 3))))
}

/*
Given an array of [Vector2]s representing tiles, builds an atlas. The returned dictionary has two keys: [code]points[/code] is a [PackedVector2Array] that specifies the positions of each tile, [code]size[/code] contains the overall size of the whole atlas as [Vector2i].
*/
func MakeAtlas(sizes []Vector2.XY) Atlas { //gd:Geometry2D.make_atlas
	once.Do(singleton)
	return Atlas(gd.DictionaryAs[Atlas](class(self).MakeAtlas(Packed.New(sizes...))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]gdclass.Geometry2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Returns [code]true[/code] if [param point] is inside the circle or if it's located exactly [i]on[/i] the circle's boundary, otherwise returns [code]false[/code].
*/
//go:nosplit
func (self class) IsPointInCircle(point gd.Vector2, circle_position gd.Vector2, circle_radius gd.Float) bool { //gd:Geometry2D.is_point_in_circle
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, circle_position)
	callframe.Arg(frame, circle_radius)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_is_point_in_circle, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Given the 2D segment ([param segment_from], [param segment_to]), returns the position on the segment (as a number between 0 and 1) at which the segment hits the circle that is located at position [param circle_position] and has radius [param circle_radius]. If the segment does not intersect the circle, -1 is returned (this is also the case if the line extending the segment would intersect the circle, but the segment does not).
*/
//go:nosplit
func (self class) SegmentIntersectsCircle(segment_from gd.Vector2, segment_to gd.Vector2, circle_position gd.Vector2, circle_radius gd.Float) gd.Float { //gd:Geometry2D.segment_intersects_circle
	var frame = callframe.New()
	callframe.Arg(frame, segment_from)
	callframe.Arg(frame, segment_to)
	callframe.Arg(frame, circle_position)
	callframe.Arg(frame, circle_radius)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_segment_intersects_circle, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Checks if the two segments ([param from_a], [param to_a]) and ([param from_b], [param to_b]) intersect. If yes, return the point of intersection as [Vector2]. If no intersection takes place, returns [code]null[/code].
*/
//go:nosplit
func (self class) SegmentIntersectsSegment(from_a gd.Vector2, to_a gd.Vector2, from_b gd.Vector2, to_b gd.Vector2) gd.Variant { //gd:Geometry2D.segment_intersects_segment
	var frame = callframe.New()
	callframe.Arg(frame, from_a)
	callframe.Arg(frame, to_a)
	callframe.Arg(frame, from_b)
	callframe.Arg(frame, to_b)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_segment_intersects_segment, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Checks if the two lines ([param from_a], [param dir_a]) and ([param from_b], [param dir_b]) intersect. If yes, return the point of intersection as [Vector2]. If no intersection takes place, returns [code]null[/code].
[b]Note:[/b] The lines are specified using direction vectors, not end points.
*/
//go:nosplit
func (self class) LineIntersectsLine(from_a gd.Vector2, dir_a gd.Vector2, from_b gd.Vector2, dir_b gd.Vector2) gd.Variant { //gd:Geometry2D.line_intersects_line
	var frame = callframe.New()
	callframe.Arg(frame, from_a)
	callframe.Arg(frame, dir_a)
	callframe.Arg(frame, from_b)
	callframe.Arg(frame, dir_b)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_line_intersects_line, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Given the two 2D segments ([param p1], [param q1]) and ([param p2], [param q2]), finds those two points on the two segments that are closest to each other. Returns a [PackedVector2Array] that contains this point on ([param p1], [param q1]) as well the accompanying point on ([param p2], [param q2]).
*/
//go:nosplit
func (self class) GetClosestPointsBetweenSegments(p1 gd.Vector2, q1 gd.Vector2, p2 gd.Vector2, q2 gd.Vector2) Packed.Array[Vector2.XY] { //gd:Geometry2D.get_closest_points_between_segments
	var frame = callframe.New()
	callframe.Arg(frame, p1)
	callframe.Arg(frame, q1)
	callframe.Arg(frame, p2)
	callframe.Arg(frame, q2)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_get_closest_points_between_segments, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[Vector2.XY](Array.Through(gd.PackedProxy[gd.PackedVector2Array, Vector2.XY]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the 2D point on the 2D segment ([param s1], [param s2]) that is closest to [param point]. The returned point will always be inside the specified segment.
*/
//go:nosplit
func (self class) GetClosestPointToSegment(point gd.Vector2, s1 gd.Vector2, s2 gd.Vector2) gd.Vector2 { //gd:Geometry2D.get_closest_point_to_segment
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, s1)
	callframe.Arg(frame, s2)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_get_closest_point_to_segment, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the 2D point on the 2D line defined by ([param s1], [param s2]) that is closest to [param point]. The returned point can be inside the segment ([param s1], [param s2]) or outside of it, i.e. somewhere on the line extending from the segment.
*/
//go:nosplit
func (self class) GetClosestPointToSegmentUncapped(point gd.Vector2, s1 gd.Vector2, s2 gd.Vector2) gd.Vector2 { //gd:Geometry2D.get_closest_point_to_segment_uncapped
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, s1)
	callframe.Arg(frame, s2)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_get_closest_point_to_segment_uncapped, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns if [param point] is inside the triangle specified by [param a], [param b] and [param c].
*/
//go:nosplit
func (self class) PointIsInsideTriangle(point gd.Vector2, a gd.Vector2, b gd.Vector2, c gd.Vector2) bool { //gd:Geometry2D.point_is_inside_triangle
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, a)
	callframe.Arg(frame, b)
	callframe.Arg(frame, c)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_point_is_inside_triangle, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if [param polygon]'s vertices are ordered in clockwise order, otherwise returns [code]false[/code].
[b]Note:[/b] Assumes a Cartesian coordinate system where [code]+x[/code] is right and [code]+y[/code] is up. If using screen coordinates ([code]+y[/code] is down), the result will need to be flipped (i.e. a [code]true[/code] result will indicate counter-clockwise).
*/
//go:nosplit
func (self class) IsPolygonClockwise(polygon Packed.Array[Vector2.XY]) bool { //gd:Geometry2D.is_polygon_clockwise
	var frame = callframe.New()
	callframe.Arg(frame, gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](polygon))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_is_polygon_clockwise, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if [param point] is inside [param polygon] or if it's located exactly [i]on[/i] polygon's boundary, otherwise returns [code]false[/code].
*/
//go:nosplit
func (self class) IsPointInPolygon(point gd.Vector2, polygon Packed.Array[Vector2.XY]) bool { //gd:Geometry2D.is_point_in_polygon
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](polygon))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_is_point_in_polygon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Triangulates the polygon specified by the points in [param polygon]. Returns a [PackedInt32Array] where each triangle consists of three consecutive point indices into [param polygon] (i.e. the returned array will have [code]n * 3[/code] elements, with [code]n[/code] being the number of found triangles). Output triangles will always be counter clockwise, and the contour will be flipped if it's clockwise. If the triangulation did not succeed, an empty [PackedInt32Array] is returned.
*/
//go:nosplit
func (self class) TriangulatePolygon(polygon Packed.Array[Vector2.XY]) Packed.Array[int32] { //gd:Geometry2D.triangulate_polygon
	var frame = callframe.New()
	callframe.Arg(frame, gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](polygon))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_triangulate_polygon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int32](Array.Through(gd.PackedProxy[gd.PackedInt32Array, int32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Triangulates the area specified by discrete set of [param points] such that no point is inside the circumcircle of any resulting triangle. Returns a [PackedInt32Array] where each triangle consists of three consecutive point indices into [param points] (i.e. the returned array will have [code]n * 3[/code] elements, with [code]n[/code] being the number of found triangles). If the triangulation did not succeed, an empty [PackedInt32Array] is returned.
*/
//go:nosplit
func (self class) TriangulateDelaunay(points Packed.Array[Vector2.XY]) Packed.Array[int32] { //gd:Geometry2D.triangulate_delaunay
	var frame = callframe.New()
	callframe.Arg(frame, gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](points))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_triangulate_delaunay, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int32](Array.Through(gd.PackedProxy[gd.PackedInt32Array, int32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Given an array of [Vector2]s, returns the convex hull as a list of points in counterclockwise order. The last point is the same as the first one.
*/
//go:nosplit
func (self class) ConvexHull(points Packed.Array[Vector2.XY]) Packed.Array[Vector2.XY] { //gd:Geometry2D.convex_hull
	var frame = callframe.New()
	callframe.Arg(frame, gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](points))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_convex_hull, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[Vector2.XY](Array.Through(gd.PackedProxy[gd.PackedVector2Array, Vector2.XY]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Decomposes the [param polygon] into multiple convex hulls and returns an array of [PackedVector2Array].
*/
//go:nosplit
func (self class) DecomposePolygonInConvex(polygon Packed.Array[Vector2.XY]) Array.Contains[Packed.Array[Vector2.XY]] { //gd:Geometry2D.decompose_polygon_in_convex
	var frame = callframe.New()
	callframe.Arg(frame, gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](polygon))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_decompose_polygon_in_convex, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Packed.Array[Vector2.XY]]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Merges (combines) [param polygon_a] and [param polygon_b] and returns an array of merged polygons. This performs [constant OPERATION_UNION] between polygons.
The operation may result in an outer polygon (boundary) and multiple inner polygons (holes) produced which could be distinguished by calling [method is_polygon_clockwise].
*/
//go:nosplit
func (self class) MergePolygons(polygon_a Packed.Array[Vector2.XY], polygon_b Packed.Array[Vector2.XY]) Array.Contains[Packed.Array[Vector2.XY]] { //gd:Geometry2D.merge_polygons
	var frame = callframe.New()
	callframe.Arg(frame, gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](polygon_a))
	callframe.Arg(frame, gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](polygon_b))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_merge_polygons, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Packed.Array[Vector2.XY]]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Clips [param polygon_a] against [param polygon_b] and returns an array of clipped polygons. This performs [constant OPERATION_DIFFERENCE] between polygons. Returns an empty array if [param polygon_b] completely overlaps [param polygon_a].
If [param polygon_b] is enclosed by [param polygon_a], returns an outer polygon (boundary) and inner polygon (hole) which could be distinguished by calling [method is_polygon_clockwise].
*/
//go:nosplit
func (self class) ClipPolygons(polygon_a Packed.Array[Vector2.XY], polygon_b Packed.Array[Vector2.XY]) Array.Contains[Packed.Array[Vector2.XY]] { //gd:Geometry2D.clip_polygons
	var frame = callframe.New()
	callframe.Arg(frame, gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](polygon_a))
	callframe.Arg(frame, gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](polygon_b))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_clip_polygons, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Packed.Array[Vector2.XY]]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Intersects [param polygon_a] with [param polygon_b] and returns an array of intersected polygons. This performs [constant OPERATION_INTERSECTION] between polygons. In other words, returns common area shared by polygons. Returns an empty array if no intersection occurs.
The operation may result in an outer polygon (boundary) and inner polygon (hole) produced which could be distinguished by calling [method is_polygon_clockwise].
*/
//go:nosplit
func (self class) IntersectPolygons(polygon_a Packed.Array[Vector2.XY], polygon_b Packed.Array[Vector2.XY]) Array.Contains[Packed.Array[Vector2.XY]] { //gd:Geometry2D.intersect_polygons
	var frame = callframe.New()
	callframe.Arg(frame, gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](polygon_a))
	callframe.Arg(frame, gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](polygon_b))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_intersect_polygons, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Packed.Array[Vector2.XY]]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Mutually excludes common area defined by intersection of [param polygon_a] and [param polygon_b] (see [method intersect_polygons]) and returns an array of excluded polygons. This performs [constant OPERATION_XOR] between polygons. In other words, returns all but common area between polygons.
The operation may result in an outer polygon (boundary) and inner polygon (hole) produced which could be distinguished by calling [method is_polygon_clockwise].
*/
//go:nosplit
func (self class) ExcludePolygons(polygon_a Packed.Array[Vector2.XY], polygon_b Packed.Array[Vector2.XY]) Array.Contains[Packed.Array[Vector2.XY]] { //gd:Geometry2D.exclude_polygons
	var frame = callframe.New()
	callframe.Arg(frame, gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](polygon_a))
	callframe.Arg(frame, gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](polygon_b))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_exclude_polygons, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Packed.Array[Vector2.XY]]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Clips [param polyline] against [param polygon] and returns an array of clipped polylines. This performs [constant OPERATION_DIFFERENCE] between the polyline and the polygon. This operation can be thought of as cutting a line with a closed shape.
*/
//go:nosplit
func (self class) ClipPolylineWithPolygon(polyline Packed.Array[Vector2.XY], polygon Packed.Array[Vector2.XY]) Array.Contains[Packed.Array[Vector2.XY]] { //gd:Geometry2D.clip_polyline_with_polygon
	var frame = callframe.New()
	callframe.Arg(frame, gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](polyline))
	callframe.Arg(frame, gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](polygon))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_clip_polyline_with_polygon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Packed.Array[Vector2.XY]]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Intersects [param polyline] with [param polygon] and returns an array of intersected polylines. This performs [constant OPERATION_INTERSECTION] between the polyline and the polygon. This operation can be thought of as chopping a line with a closed shape.
*/
//go:nosplit
func (self class) IntersectPolylineWithPolygon(polyline Packed.Array[Vector2.XY], polygon Packed.Array[Vector2.XY]) Array.Contains[Packed.Array[Vector2.XY]] { //gd:Geometry2D.intersect_polyline_with_polygon
	var frame = callframe.New()
	callframe.Arg(frame, gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](polyline))
	callframe.Arg(frame, gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](polygon))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_intersect_polyline_with_polygon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Packed.Array[Vector2.XY]]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Inflates or deflates [param polygon] by [param delta] units (pixels). If [param delta] is positive, makes the polygon grow outward. If [param delta] is negative, shrinks the polygon inward. Returns an array of polygons because inflating/deflating may result in multiple discrete polygons. Returns an empty array if [param delta] is negative and the absolute value of it approximately exceeds the minimum bounding rectangle dimensions of the polygon.
Each polygon's vertices will be rounded as determined by [param join_type], see [enum PolyJoinType].
The operation may result in an outer polygon (boundary) and inner polygon (hole) produced which could be distinguished by calling [method is_polygon_clockwise].
[b]Note:[/b] To translate the polygon's vertices specifically, multiply them to a [Transform2D]:
[codeblocks]
[gdscript]
var polygon = PackedVector2Array([Vector2(0, 0), Vector2(100, 0), Vector2(100, 100), Vector2(0, 100)])
var offset = Vector2(50, 50)
polygon = Transform2D(0, offset) * polygon
print(polygon) # prints [(50, 50), (150, 50), (150, 150), (50, 150)]
[/gdscript]
[csharp]
var polygon = new Vector2[] { new Vector2(0, 0), new Vector2(100, 0), new Vector2(100, 100), new Vector2(0, 100) };
var offset = new Vector2(50, 50);
polygon = new Transform2D(0, offset) * polygon;
GD.Print((Variant)polygon); // prints [(50, 50), (150, 50), (150, 150), (50, 150)]
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) OffsetPolygon(polygon Packed.Array[Vector2.XY], delta gd.Float, join_type gdclass.Geometry2DPolyJoinType) Array.Contains[Packed.Array[Vector2.XY]] { //gd:Geometry2D.offset_polygon
	var frame = callframe.New()
	callframe.Arg(frame, gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](polygon))
	callframe.Arg(frame, delta)
	callframe.Arg(frame, join_type)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_offset_polygon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Packed.Array[Vector2.XY]]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Inflates or deflates [param polyline] by [param delta] units (pixels), producing polygons. If [param delta] is positive, makes the polyline grow outward. Returns an array of polygons because inflating/deflating may result in multiple discrete polygons. If [param delta] is negative, returns an empty array.
Each polygon's vertices will be rounded as determined by [param join_type], see [enum PolyJoinType].
Each polygon's endpoints will be rounded as determined by [param end_type], see [enum PolyEndType].
The operation may result in an outer polygon (boundary) and inner polygon (hole) produced which could be distinguished by calling [method is_polygon_clockwise].
*/
//go:nosplit
func (self class) OffsetPolyline(polyline Packed.Array[Vector2.XY], delta gd.Float, join_type gdclass.Geometry2DPolyJoinType, end_type gdclass.Geometry2DPolyEndType) Array.Contains[Packed.Array[Vector2.XY]] { //gd:Geometry2D.offset_polyline
	var frame = callframe.New()
	callframe.Arg(frame, gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](polyline))
	callframe.Arg(frame, delta)
	callframe.Arg(frame, join_type)
	callframe.Arg(frame, end_type)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_offset_polyline, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Packed.Array[Vector2.XY]]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Given an array of [Vector2]s representing tiles, builds an atlas. The returned dictionary has two keys: [code]points[/code] is a [PackedVector2Array] that specifies the positions of each tile, [code]size[/code] contains the overall size of the whole atlas as [Vector2i].
*/
//go:nosplit
func (self class) MakeAtlas(sizes Packed.Array[Vector2.XY]) Dictionary.Any { //gd:Geometry2D.make_atlas
	var frame = callframe.New()
	callframe.Arg(frame, gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](sizes))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Geometry2D.Bind_make_atlas, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
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
	gdclass.Register("Geometry2D", func(ptr gd.Object) any { return [1]gdclass.Geometry2D{*(*gdclass.Geometry2D)(unsafe.Pointer(&ptr))} })
}

type PolyBooleanOperation = gdclass.Geometry2DPolyBooleanOperation //gd:Geometry2D.PolyBooleanOperation

const (
	/*Create regions where either subject or clip polygons (or both) are filled.*/
	OperationUnion PolyBooleanOperation = 0
	/*Create regions where subject polygons are filled except where clip polygons are filled.*/
	OperationDifference PolyBooleanOperation = 1
	/*Create regions where both subject and clip polygons are filled.*/
	OperationIntersection PolyBooleanOperation = 2
	/*Create regions where either subject or clip polygons are filled but not where both are filled.*/
	OperationXor PolyBooleanOperation = 3
)

type PolyJoinType = gdclass.Geometry2DPolyJoinType //gd:Geometry2D.PolyJoinType

const (
	/*Squaring is applied uniformally at all convex edge joins at [code]1 * delta[/code].*/
	JoinSquare PolyJoinType = 0
	/*While flattened paths can never perfectly trace an arc, they are approximated by a series of arc chords.*/
	JoinRound PolyJoinType = 1
	/*There's a necessary limit to mitered joins since offsetting edges that join at very acute angles will produce excessively long and narrow "spikes". For any given edge join, when miter offsetting would exceed that maximum distance, "square" joining is applied.*/
	JoinMiter PolyJoinType = 2
)

type PolyEndType = gdclass.Geometry2DPolyEndType //gd:Geometry2D.PolyEndType

const (
	/*Endpoints are joined using the [enum PolyJoinType] value and the path filled as a polygon.*/
	EndPolygon PolyEndType = 0
	/*Endpoints are joined using the [enum PolyJoinType] value and the path filled as a polyline.*/
	EndJoined PolyEndType = 1
	/*Endpoints are squared off with no extension.*/
	EndButt PolyEndType = 2
	/*Endpoints are squared off and extended by [code]delta[/code] units.*/
	EndSquare PolyEndType = 3
	/*Endpoints are rounded off and extended by [code]delta[/code] units.*/
	EndRound PolyEndType = 4
)

type Atlas struct {
	Points []struct {
		X float32
		Y float32
	} `gd:"points"`
	Size struct {
		X int32
		Y int32
	} `gd:"size"`
}
