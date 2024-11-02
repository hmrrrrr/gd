package Shape2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Resource"
import "grow.graphics/gd/variant/Float"
import "grow.graphics/gd/variant/Transform2D"
import "grow.graphics/gd/variant/Vector2"
import "grow.graphics/gd/variant/Color"
import "grow.graphics/gd/variant/Rect2"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Abstract base class for all 2D shapes, intended for use in physics.
[b]Performance:[/b] Primitive shapes, especially [CircleShape2D], are fast to check collisions against. [ConvexPolygonShape2D] is slower, and [ConcavePolygonShape2D] is the slowest.
*/
type Instance [1]classdb.Shape2D

/*
Returns [code]true[/code] if this shape is colliding with another.
This method needs the transformation matrix for this shape ([param local_xform]), the shape to check collisions with ([param with_shape]), and the transformation matrix of that shape ([param shape_xform]).
*/
func (self Instance) Collide(local_xform Transform2D.OriginXY, with_shape objects.Shape2D, shape_xform Transform2D.OriginXY) bool {
	return bool(class(self).Collide(gd.Transform2D(local_xform), with_shape, gd.Transform2D(shape_xform)))
}

/*
Returns whether this shape would collide with another, if a given movement was applied.
This method needs the transformation matrix for this shape ([param local_xform]), the movement to test on this shape ([param local_motion]), the shape to check collisions with ([param with_shape]), the transformation matrix of that shape ([param shape_xform]), and the movement to test onto the other object ([param shape_motion]).
*/
func (self Instance) CollideWithMotion(local_xform Transform2D.OriginXY, local_motion Vector2.XY, with_shape objects.Shape2D, shape_xform Transform2D.OriginXY, shape_motion Vector2.XY) bool {
	return bool(class(self).CollideWithMotion(gd.Transform2D(local_xform), gd.Vector2(local_motion), with_shape, gd.Transform2D(shape_xform), gd.Vector2(shape_motion)))
}

/*
Returns a list of contact point pairs where this shape touches another.
If there are no collisions, the returned list is empty. Otherwise, the returned list contains contact points arranged in pairs, with entries alternating between points on the boundary of this shape and points on the boundary of [param with_shape].
A collision pair A, B can be used to calculate the collision normal with [code](B - A).normalized()[/code], and the collision depth with [code](B - A).length()[/code]. This information is typically used to separate shapes, particularly in collision solvers.
This method needs the transformation matrix for this shape ([param local_xform]), the shape to check collisions with ([param with_shape]), and the transformation matrix of that shape ([param shape_xform]).
*/
func (self Instance) CollideAndGetContacts(local_xform Transform2D.OriginXY, with_shape objects.Shape2D, shape_xform Transform2D.OriginXY) []Vector2.XY {
	return []Vector2.XY(class(self).CollideAndGetContacts(gd.Transform2D(local_xform), with_shape, gd.Transform2D(shape_xform)).AsSlice())
}

/*
Returns a list of contact point pairs where this shape would touch another, if a given movement was applied.
If there would be no collisions, the returned list is empty. Otherwise, the returned list contains contact points arranged in pairs, with entries alternating between points on the boundary of this shape and points on the boundary of [param with_shape].
A collision pair A, B can be used to calculate the collision normal with [code](B - A).normalized()[/code], and the collision depth with [code](B - A).length()[/code]. This information is typically used to separate shapes, particularly in collision solvers.
This method needs the transformation matrix for this shape ([param local_xform]), the movement to test on this shape ([param local_motion]), the shape to check collisions with ([param with_shape]), the transformation matrix of that shape ([param shape_xform]), and the movement to test onto the other object ([param shape_motion]).
*/
func (self Instance) CollideWithMotionAndGetContacts(local_xform Transform2D.OriginXY, local_motion Vector2.XY, with_shape objects.Shape2D, shape_xform Transform2D.OriginXY, shape_motion Vector2.XY) []Vector2.XY {
	return []Vector2.XY(class(self).CollideWithMotionAndGetContacts(gd.Transform2D(local_xform), gd.Vector2(local_motion), with_shape, gd.Transform2D(shape_xform), gd.Vector2(shape_motion)).AsSlice())
}

/*
Draws a solid shape onto a [CanvasItem] with the [RenderingServer] API filled with the specified [param color]. The exact drawing method is specific for each shape and cannot be configured.
*/
func (self Instance) Draw(canvas_item Resource.ID, color Color.RGBA) {
	class(self).Draw(canvas_item, gd.Color(color))
}

/*
Returns a [Rect2] representing the shapes boundary.
*/
func (self Instance) GetRect() Rect2.PositionSize {
	return Rect2.PositionSize(class(self).GetRect())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Shape2D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Shape2D"))
	return Instance{classdb.Shape2D(object)}
}

func (self Instance) CustomSolverBias() Float.X {
	return Float.X(Float.X(class(self).GetCustomSolverBias()))
}

func (self Instance) SetCustomSolverBias(value Float.X) {
	class(self).SetCustomSolverBias(gd.Float(value))
}

//go:nosplit
func (self class) SetCustomSolverBias(bias gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, bias)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shape2D.Bind_set_custom_solver_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCustomSolverBias() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shape2D.Bind_get_custom_solver_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if this shape is colliding with another.
This method needs the transformation matrix for this shape ([param local_xform]), the shape to check collisions with ([param with_shape]), and the transformation matrix of that shape ([param shape_xform]).
*/
//go:nosplit
func (self class) Collide(local_xform gd.Transform2D, with_shape objects.Shape2D, shape_xform gd.Transform2D) bool {
	var frame = callframe.New()
	callframe.Arg(frame, local_xform)
	callframe.Arg(frame, pointers.Get(with_shape[0])[0])
	callframe.Arg(frame, shape_xform)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shape2D.Bind_collide, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns whether this shape would collide with another, if a given movement was applied.
This method needs the transformation matrix for this shape ([param local_xform]), the movement to test on this shape ([param local_motion]), the shape to check collisions with ([param with_shape]), the transformation matrix of that shape ([param shape_xform]), and the movement to test onto the other object ([param shape_motion]).
*/
//go:nosplit
func (self class) CollideWithMotion(local_xform gd.Transform2D, local_motion gd.Vector2, with_shape objects.Shape2D, shape_xform gd.Transform2D, shape_motion gd.Vector2) bool {
	var frame = callframe.New()
	callframe.Arg(frame, local_xform)
	callframe.Arg(frame, local_motion)
	callframe.Arg(frame, pointers.Get(with_shape[0])[0])
	callframe.Arg(frame, shape_xform)
	callframe.Arg(frame, shape_motion)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shape2D.Bind_collide_with_motion, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a list of contact point pairs where this shape touches another.
If there are no collisions, the returned list is empty. Otherwise, the returned list contains contact points arranged in pairs, with entries alternating between points on the boundary of this shape and points on the boundary of [param with_shape].
A collision pair A, B can be used to calculate the collision normal with [code](B - A).normalized()[/code], and the collision depth with [code](B - A).length()[/code]. This information is typically used to separate shapes, particularly in collision solvers.
This method needs the transformation matrix for this shape ([param local_xform]), the shape to check collisions with ([param with_shape]), and the transformation matrix of that shape ([param shape_xform]).
*/
//go:nosplit
func (self class) CollideAndGetContacts(local_xform gd.Transform2D, with_shape objects.Shape2D, shape_xform gd.Transform2D) gd.PackedVector2Array {
	var frame = callframe.New()
	callframe.Arg(frame, local_xform)
	callframe.Arg(frame, pointers.Get(with_shape[0])[0])
	callframe.Arg(frame, shape_xform)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shape2D.Bind_collide_and_get_contacts, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a list of contact point pairs where this shape would touch another, if a given movement was applied.
If there would be no collisions, the returned list is empty. Otherwise, the returned list contains contact points arranged in pairs, with entries alternating between points on the boundary of this shape and points on the boundary of [param with_shape].
A collision pair A, B can be used to calculate the collision normal with [code](B - A).normalized()[/code], and the collision depth with [code](B - A).length()[/code]. This information is typically used to separate shapes, particularly in collision solvers.
This method needs the transformation matrix for this shape ([param local_xform]), the movement to test on this shape ([param local_motion]), the shape to check collisions with ([param with_shape]), the transformation matrix of that shape ([param shape_xform]), and the movement to test onto the other object ([param shape_motion]).
*/
//go:nosplit
func (self class) CollideWithMotionAndGetContacts(local_xform gd.Transform2D, local_motion gd.Vector2, with_shape objects.Shape2D, shape_xform gd.Transform2D, shape_motion gd.Vector2) gd.PackedVector2Array {
	var frame = callframe.New()
	callframe.Arg(frame, local_xform)
	callframe.Arg(frame, local_motion)
	callframe.Arg(frame, pointers.Get(with_shape[0])[0])
	callframe.Arg(frame, shape_xform)
	callframe.Arg(frame, shape_motion)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shape2D.Bind_collide_with_motion_and_get_contacts, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Draws a solid shape onto a [CanvasItem] with the [RenderingServer] API filled with the specified [param color]. The exact drawing method is specific for each shape and cannot be configured.
*/
//go:nosplit
func (self class) Draw(canvas_item gd.RID, color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, canvas_item)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shape2D.Bind_draw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a [Rect2] representing the shapes boundary.
*/
//go:nosplit
func (self class) GetRect() gd.Rect2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shape2D.Bind_get_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsShape2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsShape2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() { classdb.Register("Shape2D", func(ptr gd.Object) any { return classdb.Shape2D(ptr) }) }
