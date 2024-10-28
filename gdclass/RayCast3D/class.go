package RayCast3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
A raycast represents a ray from its origin to its [member target_position] that finds the closest [CollisionObject3D] along its path, if it intersects any.
[RayCast3D] can ignore some objects by adding them to an exception list, by making its detection reporting ignore [Area3D]s ([member collide_with_areas]) or [PhysicsBody3D]s ([member collide_with_bodies]), or by configuring physics layers.
[RayCast3D] calculates intersection every physics frame, and it holds the result until the next physics frame. For an immediate raycast, or if you want to configure a [RayCast3D] multiple times within the same physics frame, use [method force_raycast_update].
To sweep over a region of 3D space, you can approximate the region with multiple [RayCast3D]s or use [ShapeCast3D].

*/
type Go [1]classdb.RayCast3D

/*
Returns whether any object is intersecting with the ray's vector (considering the vector length).
*/
func (self Go) IsColliding() bool {
	return bool(class(self).IsColliding())
}

/*
Updates the collision information for the ray immediately, without waiting for the next [code]_physics_process[/code] call. Use this method, for example, when the ray or its parent has changed state.
[b]Note:[/b] [member enabled] does not need to be [code]true[/code] for this to work.
*/
func (self Go) ForceRaycastUpdate() {
	class(self).ForceRaycastUpdate()
}

/*
Returns the first object that the ray intersects, or [code]null[/code] if no object is intersecting the ray (i.e. [method is_colliding] returns [code]false[/code]).
*/
func (self Go) GetCollider() gd.Object {
	return gd.Object(class(self).GetCollider())
}

/*
Returns the [RID] of the first object that the ray intersects, or an empty [RID] if no object is intersecting the ray (i.e. [method is_colliding] returns [code]false[/code]).
*/
func (self Go) GetColliderRid() gd.RID {
	return gd.RID(class(self).GetColliderRid())
}

/*
Returns the shape ID of the first object that the ray intersects, or [code]0[/code] if no object is intersecting the ray (i.e. [method is_colliding] returns [code]false[/code]).
To get the intersected shape node, for a [CollisionObject3D] target, use:
[codeblocks]
[gdscript]
var target = get_collider() # A CollisionObject3D.
var shape_id = get_collider_shape() # The shape index in the collider.
var owner_id = target.shape_find_owner(shape_id) # The owner ID in the collider.
var shape = target.shape_owner_get_owner(owner_id)
[/gdscript]
[csharp]
var target = (CollisionObject3D)GetCollider(); // A CollisionObject3D.
var shapeId = GetColliderShape(); // The shape index in the collider.
var ownerId = target.ShapeFindOwner(shapeId); // The owner ID in the collider.
var shape = target.ShapeOwnerGetOwner(ownerId);
[/csharp]
[/codeblocks]
*/
func (self Go) GetColliderShape() int {
	return int(int(class(self).GetColliderShape()))
}

/*
Returns the collision point at which the ray intersects the closest object, in the global coordinate system. If [member hit_from_inside] is [code]true[/code] and the ray starts inside of a collision shape, this function will return the origin point of the ray.
[b]Note:[/b] Check that [method is_colliding] returns [code]true[/code] before calling this method to ensure the returned point is valid and up-to-date.
*/
func (self Go) GetCollisionPoint() gd.Vector3 {
	return gd.Vector3(class(self).GetCollisionPoint())
}

/*
Returns the normal of the intersecting object's shape at the collision point, or [code]Vector3(0, 0, 0)[/code] if the ray starts inside the shape and [member hit_from_inside] is [code]true[/code].
[b]Note:[/b] Check that [method is_colliding] returns [code]true[/code] before calling this method to ensure the returned normal is valid and up-to-date.
*/
func (self Go) GetCollisionNormal() gd.Vector3 {
	return gd.Vector3(class(self).GetCollisionNormal())
}

/*
Returns the collision object's face index at the collision point, or [code]-1[/code] if the shape intersecting the ray is not a [ConcavePolygonShape3D].
*/
func (self Go) GetCollisionFaceIndex() int {
	return int(int(class(self).GetCollisionFaceIndex()))
}

/*
Adds a collision exception so the ray does not report collisions with the specified [RID].
*/
func (self Go) AddExceptionRid(rid gd.RID) {
	class(self).AddExceptionRid(rid)
}

/*
Adds a collision exception so the ray does not report collisions with the specified [CollisionObject3D] node.
*/
func (self Go) AddException(node gdclass.CollisionObject3D) {
	class(self).AddException(node)
}

/*
Removes a collision exception so the ray does report collisions with the specified [RID].
*/
func (self Go) RemoveExceptionRid(rid gd.RID) {
	class(self).RemoveExceptionRid(rid)
}

/*
Removes a collision exception so the ray does report collisions with the specified [CollisionObject3D] node.
*/
func (self Go) RemoveException(node gdclass.CollisionObject3D) {
	class(self).RemoveException(node)
}

/*
Removes all collision exceptions for this ray.
*/
func (self Go) ClearExceptions() {
	class(self).ClearExceptions()
}

/*
Based on [param value], enables or disables the specified layer in the [member collision_mask], given a [param layer_number] between 1 and 32.
*/
func (self Go) SetCollisionMaskValue(layer_number int, value bool) {
	class(self).SetCollisionMaskValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Go) GetCollisionMaskValue(layer_number int) bool {
	return bool(class(self).GetCollisionMaskValue(gd.Int(layer_number)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.RayCast3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RayCast3D"))
	return Go{classdb.RayCast3D(object)}
}

func (self Go) Enabled() bool {
		return bool(class(self).IsEnabled())
}

func (self Go) SetEnabled(value bool) {
	class(self).SetEnabled(value)
}

func (self Go) ExcludeParent() bool {
		return bool(class(self).GetExcludeParentBody())
}

func (self Go) SetExcludeParent(value bool) {
	class(self).SetExcludeParentBody(value)
}

func (self Go) TargetPosition() gd.Vector3 {
		return gd.Vector3(class(self).GetTargetPosition())
}

func (self Go) SetTargetPosition(value gd.Vector3) {
	class(self).SetTargetPosition(value)
}

func (self Go) CollisionMask() int {
		return int(int(class(self).GetCollisionMask()))
}

func (self Go) SetCollisionMask(value int) {
	class(self).SetCollisionMask(gd.Int(value))
}

func (self Go) HitFromInside() bool {
		return bool(class(self).IsHitFromInsideEnabled())
}

func (self Go) SetHitFromInside(value bool) {
	class(self).SetHitFromInside(value)
}

func (self Go) HitBackFaces() bool {
		return bool(class(self).IsHitBackFacesEnabled())
}

func (self Go) SetHitBackFaces(value bool) {
	class(self).SetHitBackFaces(value)
}

func (self Go) CollideWithAreas() bool {
		return bool(class(self).IsCollideWithAreasEnabled())
}

func (self Go) SetCollideWithAreas(value bool) {
	class(self).SetCollideWithAreas(value)
}

func (self Go) CollideWithBodies() bool {
		return bool(class(self).IsCollideWithBodiesEnabled())
}

func (self Go) SetCollideWithBodies(value bool) {
	class(self).SetCollideWithBodies(value)
}

func (self Go) DebugShapeCustomColor() gd.Color {
		return gd.Color(class(self).GetDebugShapeCustomColor())
}

func (self Go) SetDebugShapeCustomColor(value gd.Color) {
	class(self).SetDebugShapeCustomColor(value)
}

func (self Go) DebugShapeThickness() int {
		return int(int(class(self).GetDebugShapeThickness()))
}

func (self Go) SetDebugShapeThickness(value int) {
	class(self).SetDebugShapeThickness(gd.Int(value))
}

//go:nosplit
func (self class) SetEnabled(enabled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_set_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_is_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTargetPosition(local_point gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, local_point)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_set_target_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTargetPosition() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_get_target_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_is_colliding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Updates the collision information for the ray immediately, without waiting for the next [code]_physics_process[/code] call. Use this method, for example, when the ray or its parent has changed state.
[b]Note:[/b] [member enabled] does not need to be [code]true[/code] for this to work.
*/
//go:nosplit
func (self class) ForceRaycastUpdate()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_force_raycast_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the first object that the ray intersects, or [code]null[/code] if no object is intersecting the ray (i.e. [method is_colliding] returns [code]false[/code]).
*/
//go:nosplit
func (self class) GetCollider() gd.Object {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_get_collider, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_get_collider_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the shape ID of the first object that the ray intersects, or [code]0[/code] if no object is intersecting the ray (i.e. [method is_colliding] returns [code]false[/code]).
To get the intersected shape node, for a [CollisionObject3D] target, use:
[codeblocks]
[gdscript]
var target = get_collider() # A CollisionObject3D.
var shape_id = get_collider_shape() # The shape index in the collider.
var owner_id = target.shape_find_owner(shape_id) # The owner ID in the collider.
var shape = target.shape_owner_get_owner(owner_id)
[/gdscript]
[csharp]
var target = (CollisionObject3D)GetCollider(); // A CollisionObject3D.
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_get_collider_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the collision point at which the ray intersects the closest object, in the global coordinate system. If [member hit_from_inside] is [code]true[/code] and the ray starts inside of a collision shape, this function will return the origin point of the ray.
[b]Note:[/b] Check that [method is_colliding] returns [code]true[/code] before calling this method to ensure the returned point is valid and up-to-date.
*/
//go:nosplit
func (self class) GetCollisionPoint() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_get_collision_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the normal of the intersecting object's shape at the collision point, or [code]Vector3(0, 0, 0)[/code] if the ray starts inside the shape and [member hit_from_inside] is [code]true[/code].
[b]Note:[/b] Check that [method is_colliding] returns [code]true[/code] before calling this method to ensure the returned normal is valid and up-to-date.
*/
//go:nosplit
func (self class) GetCollisionNormal() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_get_collision_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the collision object's face index at the collision point, or [code]-1[/code] if the shape intersecting the ray is not a [ConcavePolygonShape3D].
*/
//go:nosplit
func (self class) GetCollisionFaceIndex() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_get_collision_face_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a collision exception so the ray does not report collisions with the specified [RID].
*/
//go:nosplit
func (self class) AddExceptionRid(rid gd.RID)  {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_add_exception_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a collision exception so the ray does not report collisions with the specified [CollisionObject3D] node.
*/
//go:nosplit
func (self class) AddException(node gdclass.CollisionObject3D)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(node[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_add_exception, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a collision exception so the ray does report collisions with the specified [RID].
*/
//go:nosplit
func (self class) RemoveExceptionRid(rid gd.RID)  {
	var frame = callframe.New()
	callframe.Arg(frame, rid)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_remove_exception_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a collision exception so the ray does report collisions with the specified [CollisionObject3D] node.
*/
//go:nosplit
func (self class) RemoveException(node gdclass.CollisionObject3D)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(node[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_remove_exception, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes all collision exceptions for this ray.
*/
//go:nosplit
func (self class) ClearExceptions()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_clear_exceptions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetCollisionMask(mask gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCollisionMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Based on [param value], enables or disables the specified layer in the [member collision_mask], given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetCollisionMaskValue(layer_number gd.Int, value bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_set_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_get_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetExcludeParentBody(mask bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_set_exclude_parent_body, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetExcludeParentBody() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_get_exclude_parent_body, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollideWithAreas(enable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_set_collide_with_areas, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCollideWithAreasEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_is_collide_with_areas_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollideWithBodies(enable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_set_collide_with_bodies, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCollideWithBodiesEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_is_collide_with_bodies_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHitFromInside(enable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_set_hit_from_inside, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsHitFromInsideEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_is_hit_from_inside_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHitBackFaces(enable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_set_hit_back_faces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsHitBackFacesEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_is_hit_back_faces_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDebugShapeCustomColor(debug_shape_custom_color gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, debug_shape_custom_color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_set_debug_shape_custom_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDebugShapeCustomColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_get_debug_shape_custom_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDebugShapeThickness(debug_shape_thickness gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, debug_shape_thickness)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_set_debug_shape_thickness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDebugShapeThickness() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RayCast3D.Bind_get_debug_shape_thickness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRayCast3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsRayCast3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}
func init() {classdb.Register("RayCast3D", func(ptr gd.Object) any { return classdb.RayCast3D(ptr) })}
