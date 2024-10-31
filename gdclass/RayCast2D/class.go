package RayCast2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node2D"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
A raycast represents a ray from its origin to its [member target_position] that finds the closest [CollisionObject2D] along its path, if it intersects any.
[RayCast2D] can ignore some objects by adding them to an exception list, by making its detection reporting ignore [Area2D]s ([member collide_with_areas]) or [PhysicsBody2D]s ([member collide_with_bodies]), or by configuring physics layers.
[RayCast2D] calculates intersection every physics frame, and it holds the result until the next physics frame. For an immediate raycast, or if you want to configure a [RayCast2D] multiple times within the same physics frame, use [method force_raycast_update].
To sweep over a region of 2D space, you can approximate the region with multiple [RayCast2D]s or use [ShapeCast2D].
*/
type Instance [1]classdb.RayCast2D

/*
Returns whether any object is intersecting with the ray's vector (considering the vector length).
*/
func (self Instance) IsColliding() bool {
	return bool(class(self).IsColliding())
}

/*
Updates the collision information for the ray immediately, without waiting for the next [code]_physics_process[/code] call. Use this method, for example, when the ray or its parent has changed state.
[b]Note:[/b] [member enabled] does not need to be [code]true[/code] for this to work.
*/
func (self Instance) ForceRaycastUpdate() {
	class(self).ForceRaycastUpdate()
}

/*
Returns the first object that the ray intersects, or [code]null[/code] if no object is intersecting the ray (i.e. [method is_colliding] returns [code]false[/code]).
*/
func (self Instance) GetCollider() gd.Object {
	return gd.Object(class(self).GetCollider())
}

/*
Returns the [RID] of the first object that the ray intersects, or an empty [RID] if no object is intersecting the ray (i.e. [method is_colliding] returns [code]false[/code]).
*/
func (self Instance) GetColliderRid() gd.RID {
	return gd.RID(class(self).GetColliderRid())
}

/*
Returns the shape ID of the first object that the ray intersects, or [code]0[/code] if no object is intersecting the ray (i.e. [method is_colliding] returns [code]false[/code]).
To get the intersected shape node, for a [CollisionObject2D] target, use:
[codeblocks]
[gdscript]
var target = get_collider() # A CollisionObject2D.
var shape_id = get_collider_shape() # The shape index in the collider.
var owner_id = target.shape_find_owner(shape_id) # The owner ID in the collider.
var shape = target.shape_owner_get_owner(owner_id)
[/gdscript]
[csharp]
var target = (CollisionObject2D)GetCollider(); // A CollisionObject2D.
var shapeId = GetColliderShape(); // The shape index in the collider.
var ownerId = target.ShapeFindOwner(shapeId); // The owner ID in the collider.
var shape = target.ShapeOwnerGetOwner(ownerId);
[/csharp]
[/codeblocks]
*/
func (self Instance) GetColliderShape() int {
	return int(int(class(self).GetColliderShape()))
}

/*
Returns the collision point at which the ray intersects the closest object, in the global coordinate system. If [member hit_from_inside] is [code]true[/code] and the ray starts inside of a collision shape, this function will return the origin point of the ray.
[b]Note:[/b] Check that [method is_colliding] returns [code]true[/code] before calling this method to ensure the returned point is valid and up-to-date.
*/
func (self Instance) GetCollisionPoint() gd.Vector2 {
	return gd.Vector2(class(self).GetCollisionPoint())
}

/*
Returns the normal of the intersecting object's shape at the collision point, or [code]Vector2(0, 0)[/code] if the ray starts inside the shape and [member hit_from_inside] is [code]true[/code].
[b]Note:[/b] Check that [method is_colliding] returns [code]true[/code] before calling this method to ensure the returned normal is valid and up-to-date.
*/
func (self Instance) GetCollisionNormal() gd.Vector2 {
	return gd.Vector2(class(self).GetCollisionNormal())
}

/*
Adds a collision exception so the ray does not report collisions with the specified [RID].
*/
func (self Instance) AddExceptionRid(rid gd.RID) {
	class(self).AddExceptionRid(rid)
}

/*
Adds a collision exception so the ray does not report collisions with the specified [CollisionObject2D] node.
*/
func (self Instance) AddException(node gdclass.CollisionObject2D) {
	class(self).AddException(node)
}

/*
Removes a collision exception so the ray does report collisions with the specified [RID].
*/
func (self Instance) RemoveExceptionRid(rid gd.RID) {
	class(self).RemoveExceptionRid(rid)
}

/*
Removes a collision exception so the ray does report collisions with the specified [CollisionObject2D] node.
*/
func (self Instance) RemoveException(node gdclass.CollisionObject2D) {
	class(self).RemoveException(node)
}

/*
Removes all collision exceptions for this ray.
*/
func (self Instance) ClearExceptions() {
	class(self).ClearExceptions()
}

/*
Based on [param value], enables or disables the specified layer in the [member collision_mask], given a [param layer_number] between 1 and 32.
*/
func (self Instance) SetCollisionMaskValue(layer_number int, value bool) {
	class(self).SetCollisionMaskValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Instance) GetCollisionMaskValue(layer_number int) bool {
	return bool(class(self).GetCollisionMaskValue(gd.Int(layer_number)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.RayCast2D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RayCast2D"))
	return Instance{classdb.RayCast2D(object)}
}

func (self Instance) Enabled() bool {
	return bool(class(self).IsEnabled())
}

func (self Instance) SetEnabled(value bool) {
	class(self).SetEnabled(value)
}

func (self Instance) ExcludeParent() bool {
	return bool(class(self).GetExcludeParentBody())
}

func (self Instance) SetExcludeParent(value bool) {
	class(self).SetExcludeParentBody(value)
}

func (self Instance) TargetPosition() gd.Vector2 {
	return gd.Vector2(class(self).GetTargetPosition())
}

func (self Instance) SetTargetPosition(value gd.Vector2) {
	class(self).SetTargetPosition(value)
}

func (self Instance) CollisionMask() int {
	return int(int(class(self).GetCollisionMask()))
}

func (self Instance) SetCollisionMask(value int) {
	class(self).SetCollisionMask(gd.Int(value))
}

func (self Instance) HitFromInside() bool {
	return bool(class(self).IsHitFromInsideEnabled())
}

func (self Instance) SetHitFromInside(value bool) {
	class(self).SetHitFromInside(value)
}

func (self Instance) CollideWithAreas() bool {
	return bool(class(self).IsCollideWithAreasEnabled())
}

func (self Instance) SetCollideWithAreas(value bool) {
	class(self).SetCollideWithAreas(value)
}

func (self Instance) CollideWithBodies() bool {
	return bool(class(self).IsCollideWithBodiesEnabled())
}

func (self Instance) SetCollideWithBodies(value bool) {
	class(self).SetCollideWithBodies(value)
}

//go:nosplit
func (self class) SetEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast2D.Bind_set_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast2D.Bind_is_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTargetPosition(local_point gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, local_point)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast2D.Bind_set_target_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTargetPosition() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast2D.Bind_get_target_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns whether any object is intersecting with the ray's vector (considering the vector length).
*/
//go:nosplit
func (self class) IsColliding() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast2D.Bind_is_colliding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Updates the collision information for the ray immediately, without waiting for the next [code]_physics_process[/code] call. Use this method, for example, when the ray or its parent has changed state.
[b]Note:[/b] [member enabled] does not need to be [code]true[/code] for this to work.
*/
//go:nosplit
func (self class) ForceRaycastUpdate() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast2D.Bind_force_raycast_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the first object that the ray intersects, or [code]null[/code] if no object is intersecting the ray (i.e. [method is_colliding] returns [code]false[/code]).
*/
//go:nosplit
func (self class) GetCollider() gd.Object {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast2D.Bind_get_collider, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gd.PointerWithOwnershipTransferredToGo(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the [RID] of the first object that the ray intersects, or an empty [RID] if no object is intersecting the ray (i.e. [method is_colliding] returns [code]false[/code]).
*/
//go:nosplit
func (self class) GetColliderRid() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast2D.Bind_get_collider_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the shape ID of the first object that the ray intersects, or [code]0[/code] if no object is intersecting the ray (i.e. [method is_colliding] returns [code]false[/code]).
To get the intersected shape node, for a [CollisionObject2D] target, use:
[codeblocks]
[gdscript]
var target = get_collider() # A CollisionObject2D.
var shape_id = get_collider_shape() # The shape index in the collider.
var owner_id = target.shape_find_owner(shape_id) # The owner ID in the collider.
var shape = target.shape_owner_get_owner(owner_id)
[/gdscript]
[csharp]
var target = (CollisionObject2D)GetCollider(); // A CollisionObject2D.
var shapeId = GetColliderShape(); // The shape index in the collider.
var ownerId = target.ShapeFindOwner(shapeId); // The owner ID in the collider.
var shape = target.ShapeOwnerGetOwner(ownerId);
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) GetColliderShape() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast2D.Bind_get_collider_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the collision point at which the ray intersects the closest object, in the global coordinate system. If [member hit_from_inside] is [code]true[/code] and the ray starts inside of a collision shape, this function will return the origin point of the ray.
[b]Note:[/b] Check that [method is_colliding] returns [code]true[/code] before calling this method to ensure the returned point is valid and up-to-date.
*/
//go:nosplit
func (self class) GetCollisionPoint() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast2D.Bind_get_collision_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the normal of the intersecting object's shape at the collision point, or [code]Vector2(0, 0)[/code] if the ray starts inside the shape and [member hit_from_inside] is [code]true[/code].
[b]Note:[/b] Check that [method is_colliding] returns [code]true[/code] before calling this method to ensure the returned normal is valid and up-to-date.
*/
//go:nosplit
func (self class) GetCollisionNormal() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast2D.Bind_get_collision_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a collision exception so the ray does not report collisions with the specified [RID].
*/
//go:nosplit
func (self class) AddExceptionRid(rid gd.RID) {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast2D.Bind_add_exception_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a collision exception so the ray does not report collisions with the specified [CollisionObject2D] node.
*/
//go:nosplit
func (self class) AddException(node gdclass.CollisionObject2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(node[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast2D.Bind_add_exception, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes a collision exception so the ray does report collisions with the specified [RID].
*/
//go:nosplit
func (self class) RemoveExceptionRid(rid gd.RID) {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast2D.Bind_remove_exception_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes a collision exception so the ray does report collisions with the specified [CollisionObject2D] node.
*/
//go:nosplit
func (self class) RemoveException(node gdclass.CollisionObject2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(node[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast2D.Bind_remove_exception, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes all collision exceptions for this ray.
*/
//go:nosplit
func (self class) ClearExceptions() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast2D.Bind_clear_exceptions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetCollisionMask(mask gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast2D.Bind_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCollisionMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast2D.Bind_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Based on [param value], enables or disables the specified layer in the [member collision_mask], given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetCollisionMaskValue(layer_number gd.Int, value bool) {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast2D.Bind_set_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether or not the specified layer of the [member collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetCollisionMaskValue(layer_number gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast2D.Bind_get_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetExcludeParentBody(mask bool) {
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast2D.Bind_set_exclude_parent_body, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetExcludeParentBody() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast2D.Bind_get_exclude_parent_body, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollideWithAreas(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast2D.Bind_set_collide_with_areas, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsCollideWithAreasEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast2D.Bind_is_collide_with_areas_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollideWithBodies(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast2D.Bind_set_collide_with_bodies, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsCollideWithBodiesEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast2D.Bind_is_collide_with_bodies_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHitFromInside(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast2D.Bind_set_hit_from_inside, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsHitFromInsideEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast2D.Bind_is_hit_from_inside_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRayCast2D() Advanced        { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRayCast2D() Instance     { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.Advanced    { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode2D() Node2D.Instance { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode2D(), name)
	}
}
func init() { classdb.Register("RayCast2D", func(ptr gd.Object) any { return classdb.RayCast2D(ptr) }) }
