// Package CharacterBody2D provides methods for working with CharacterBody2D object instances.
package CharacterBody2D

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
import "graphics.gd/classdb/PhysicsBody2D"
import "graphics.gd/classdb/CollisionObject2D"
import "graphics.gd/classdb/Node2D"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
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

/*
[CharacterBody2D] is a specialized class for physics bodies that are meant to be user-controlled. They are not affected by physics at all, but they affect other physics bodies in their path. They are mainly used to provide high-level API to move objects with wall and slope detection ([method move_and_slide] method) in addition to the general collision detection provided by [method PhysicsBody2D.move_and_collide]. This makes it useful for highly configurable physics bodies that must move in specific ways and collide with the world, as is often the case with user-controlled characters.
For game objects that don't require complex movement or collision detection, such as moving platforms, [AnimatableBody2D] is simpler to configure.
*/
type Instance [1]gdclass.CharacterBody2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCharacterBody2D() Instance
}

/*
Moves the body based on [member velocity]. If the body collides with another, it will slide along the other body (by default only on floor) rather than stop immediately. If the other body is a [CharacterBody2D] or [RigidBody2D], it will also be affected by the motion of the other body. You can use this to make moving and rotating platforms, or to make nodes push other nodes.
Modifies [member velocity] if a slide collision occurred. To get the latest collision call [method get_last_slide_collision], for detailed information about collisions that occurred, use [method get_slide_collision].
When the body touches a moving platform, the platform's velocity is automatically added to the body motion. If a collision occurs due to the platform's motion, it will always be first in the slide collisions.
The general behavior and available properties change according to the [member motion_mode].
Returns [code]true[/code] if the body collided, otherwise, returns [code]false[/code].
*/
func (self Instance) MoveAndSlide() bool { //gd:CharacterBody2D.move_and_slide
	return bool(class(self).MoveAndSlide())
}

/*
Allows to manually apply a snap to the floor regardless of the body's velocity. This function does nothing when [method is_on_floor] returns [code]true[/code].
*/
func (self Instance) ApplyFloorSnap() { //gd:CharacterBody2D.apply_floor_snap
	class(self).ApplyFloorSnap()
}

/*
Returns [code]true[/code] if the body collided with the floor on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "floor" or not.
*/
func (self Instance) IsOnFloor() bool { //gd:CharacterBody2D.is_on_floor
	return bool(class(self).IsOnFloor())
}

/*
Returns [code]true[/code] if the body collided only with the floor on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "floor" or not.
*/
func (self Instance) IsOnFloorOnly() bool { //gd:CharacterBody2D.is_on_floor_only
	return bool(class(self).IsOnFloorOnly())
}

/*
Returns [code]true[/code] if the body collided with the ceiling on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "ceiling" or not.
*/
func (self Instance) IsOnCeiling() bool { //gd:CharacterBody2D.is_on_ceiling
	return bool(class(self).IsOnCeiling())
}

/*
Returns [code]true[/code] if the body collided only with the ceiling on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "ceiling" or not.
*/
func (self Instance) IsOnCeilingOnly() bool { //gd:CharacterBody2D.is_on_ceiling_only
	return bool(class(self).IsOnCeilingOnly())
}

/*
Returns [code]true[/code] if the body collided with a wall on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "wall" or not.
*/
func (self Instance) IsOnWall() bool { //gd:CharacterBody2D.is_on_wall
	return bool(class(self).IsOnWall())
}

/*
Returns [code]true[/code] if the body collided only with a wall on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "wall" or not.
*/
func (self Instance) IsOnWallOnly() bool { //gd:CharacterBody2D.is_on_wall_only
	return bool(class(self).IsOnWallOnly())
}

/*
Returns the collision normal of the floor at the last collision point. Only valid after calling [method move_and_slide] and when [method is_on_floor] returns [code]true[/code].
[b]Warning:[/b] The collision normal is not always the same as the surface normal.
*/
func (self Instance) GetFloorNormal() Vector2.XY { //gd:CharacterBody2D.get_floor_normal
	return Vector2.XY(class(self).GetFloorNormal())
}

/*
Returns the collision normal of the wall at the last collision point. Only valid after calling [method move_and_slide] and when [method is_on_wall] returns [code]true[/code].
[b]Warning:[/b] The collision normal is not always the same as the surface normal.
*/
func (self Instance) GetWallNormal() Vector2.XY { //gd:CharacterBody2D.get_wall_normal
	return Vector2.XY(class(self).GetWallNormal())
}

/*
Returns the last motion applied to the [CharacterBody2D] during the last call to [method move_and_slide]. The movement can be split into multiple motions when sliding occurs, and this method return the last one, which is useful to retrieve the current direction of the movement.
*/
func (self Instance) GetLastMotion() Vector2.XY { //gd:CharacterBody2D.get_last_motion
	return Vector2.XY(class(self).GetLastMotion())
}

/*
Returns the travel (position delta) that occurred during the last call to [method move_and_slide].
*/
func (self Instance) GetPositionDelta() Vector2.XY { //gd:CharacterBody2D.get_position_delta
	return Vector2.XY(class(self).GetPositionDelta())
}

/*
Returns the current real velocity since the last call to [method move_and_slide]. For example, when you climb a slope, you will move diagonally even though the velocity is horizontal. This method returns the diagonal movement, as opposed to [member velocity] which returns the requested velocity.
*/
func (self Instance) GetRealVelocity() Vector2.XY { //gd:CharacterBody2D.get_real_velocity
	return Vector2.XY(class(self).GetRealVelocity())
}

/*
Returns the floor's collision angle at the last collision point according to [param up_direction], which is [constant Vector2.UP] by default. This value is always positive and only valid after calling [method move_and_slide] and when [method is_on_floor] returns [code]true[/code].
*/
func (self Instance) GetFloorAngle() Float.X { //gd:CharacterBody2D.get_floor_angle
	return Float.X(Float.X(class(self).GetFloorAngle(gd.Vector2(gd.Vector2{0, -1}))))
}

/*
Returns the linear velocity of the platform at the last collision point. Only valid after calling [method move_and_slide].
*/
func (self Instance) GetPlatformVelocity() Vector2.XY { //gd:CharacterBody2D.get_platform_velocity
	return Vector2.XY(class(self).GetPlatformVelocity())
}

/*
Returns the number of times the body collided and changed direction during the last call to [method move_and_slide].
*/
func (self Instance) GetSlideCollisionCount() int { //gd:CharacterBody2D.get_slide_collision_count
	return int(int(class(self).GetSlideCollisionCount()))
}

/*
Returns a [KinematicCollision2D], which contains information about a collision that occurred during the last call to [method move_and_slide]. Since the body can collide several times in a single call to [method move_and_slide], you must specify the index of the collision in the range 0 to ([method get_slide_collision_count] - 1).
[b]Example usage:[/b]
[codeblocks]
[gdscript]
for i in get_slide_collision_count():

	var collision = get_slide_collision(i)
	print("Collided with: ", collision.get_collider().name)

[/gdscript]
[csharp]
for (int i = 0; i < GetSlideCollisionCount(); i++)

	{
	    KinematicCollision2D collision = GetSlideCollision(i);
	    GD.Print("Collided with: ", (collision.GetCollider() as Node).Name);
	}

[/csharp]
[/codeblocks]
*/
func (self Instance) GetSlideCollision(slide_idx int) [1]gdclass.KinematicCollision2D { //gd:CharacterBody2D.get_slide_collision
	return [1]gdclass.KinematicCollision2D(class(self).GetSlideCollision(gd.Int(slide_idx)))
}

/*
Returns a [KinematicCollision2D], which contains information about the latest collision that occurred during the last call to [method move_and_slide].
*/
func (self Instance) GetLastSlideCollision() [1]gdclass.KinematicCollision2D { //gd:CharacterBody2D.get_last_slide_collision
	return [1]gdclass.KinematicCollision2D(class(self).GetLastSlideCollision())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.CharacterBody2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CharacterBody2D"))
	casted := Instance{*(*gdclass.CharacterBody2D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) MotionMode() gdclass.CharacterBody2DMotionMode {
	return gdclass.CharacterBody2DMotionMode(class(self).GetMotionMode())
}

func (self Instance) SetMotionMode(value gdclass.CharacterBody2DMotionMode) {
	class(self).SetMotionMode(value)
}

func (self Instance) UpDirection() Vector2.XY {
	return Vector2.XY(class(self).GetUpDirection())
}

func (self Instance) SetUpDirection(value Vector2.XY) {
	class(self).SetUpDirection(gd.Vector2(value))
}

func (self Instance) Velocity() Vector2.XY {
	return Vector2.XY(class(self).GetVelocity())
}

func (self Instance) SetVelocity(value Vector2.XY) {
	class(self).SetVelocity(gd.Vector2(value))
}

func (self Instance) SlideOnCeiling() bool {
	return bool(class(self).IsSlideOnCeilingEnabled())
}

func (self Instance) SetSlideOnCeiling(value bool) {
	class(self).SetSlideOnCeilingEnabled(value)
}

func (self Instance) MaxSlides() int {
	return int(int(class(self).GetMaxSlides()))
}

func (self Instance) SetMaxSlides(value int) {
	class(self).SetMaxSlides(gd.Int(value))
}

func (self Instance) WallMinSlideAngle() Float.X {
	return Float.X(Float.X(class(self).GetWallMinSlideAngle()))
}

func (self Instance) SetWallMinSlideAngle(value Float.X) {
	class(self).SetWallMinSlideAngle(gd.Float(value))
}

func (self Instance) FloorStopOnSlope() bool {
	return bool(class(self).IsFloorStopOnSlopeEnabled())
}

func (self Instance) SetFloorStopOnSlope(value bool) {
	class(self).SetFloorStopOnSlopeEnabled(value)
}

func (self Instance) FloorConstantSpeed() bool {
	return bool(class(self).IsFloorConstantSpeedEnabled())
}

func (self Instance) SetFloorConstantSpeed(value bool) {
	class(self).SetFloorConstantSpeedEnabled(value)
}

func (self Instance) FloorBlockOnWall() bool {
	return bool(class(self).IsFloorBlockOnWallEnabled())
}

func (self Instance) SetFloorBlockOnWall(value bool) {
	class(self).SetFloorBlockOnWallEnabled(value)
}

func (self Instance) FloorMaxAngle() Float.X {
	return Float.X(Float.X(class(self).GetFloorMaxAngle()))
}

func (self Instance) SetFloorMaxAngle(value Float.X) {
	class(self).SetFloorMaxAngle(gd.Float(value))
}

func (self Instance) FloorSnapLength() Float.X {
	return Float.X(Float.X(class(self).GetFloorSnapLength()))
}

func (self Instance) SetFloorSnapLength(value Float.X) {
	class(self).SetFloorSnapLength(gd.Float(value))
}

func (self Instance) PlatformOnLeave() gdclass.CharacterBody2DPlatformOnLeave {
	return gdclass.CharacterBody2DPlatformOnLeave(class(self).GetPlatformOnLeave())
}

func (self Instance) SetPlatformOnLeave(value gdclass.CharacterBody2DPlatformOnLeave) {
	class(self).SetPlatformOnLeave(value)
}

func (self Instance) PlatformFloorLayers() int {
	return int(int(class(self).GetPlatformFloorLayers()))
}

func (self Instance) SetPlatformFloorLayers(value int) {
	class(self).SetPlatformFloorLayers(gd.Int(value))
}

func (self Instance) PlatformWallLayers() int {
	return int(int(class(self).GetPlatformWallLayers()))
}

func (self Instance) SetPlatformWallLayers(value int) {
	class(self).SetPlatformWallLayers(gd.Int(value))
}

func (self Instance) SafeMargin() Float.X {
	return Float.X(Float.X(class(self).GetSafeMargin()))
}

func (self Instance) SetSafeMargin(value Float.X) {
	class(self).SetSafeMargin(gd.Float(value))
}

/*
Moves the body based on [member velocity]. If the body collides with another, it will slide along the other body (by default only on floor) rather than stop immediately. If the other body is a [CharacterBody2D] or [RigidBody2D], it will also be affected by the motion of the other body. You can use this to make moving and rotating platforms, or to make nodes push other nodes.
Modifies [member velocity] if a slide collision occurred. To get the latest collision call [method get_last_slide_collision], for detailed information about collisions that occurred, use [method get_slide_collision].
When the body touches a moving platform, the platform's velocity is automatically added to the body motion. If a collision occurs due to the platform's motion, it will always be first in the slide collisions.
The general behavior and available properties change according to the [member motion_mode].
Returns [code]true[/code] if the body collided, otherwise, returns [code]false[/code].
*/
//go:nosplit
func (self class) MoveAndSlide() bool { //gd:CharacterBody2D.move_and_slide
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_move_and_slide, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Allows to manually apply a snap to the floor regardless of the body's velocity. This function does nothing when [method is_on_floor] returns [code]true[/code].
*/
//go:nosplit
func (self class) ApplyFloorSnap() { //gd:CharacterBody2D.apply_floor_snap
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_apply_floor_snap, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetVelocity(velocity gd.Vector2) { //gd:CharacterBody2D.set_velocity
	var frame = callframe.New()
	callframe.Arg(frame, velocity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_set_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVelocity() gd.Vector2 { //gd:CharacterBody2D.get_velocity
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_get_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSafeMargin(margin gd.Float) { //gd:CharacterBody2D.set_safe_margin
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_set_safe_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSafeMargin() gd.Float { //gd:CharacterBody2D.get_safe_margin
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_get_safe_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) IsFloorStopOnSlopeEnabled() bool { //gd:CharacterBody2D.is_floor_stop_on_slope_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_is_floor_stop_on_slope_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFloorStopOnSlopeEnabled(enabled bool) { //gd:CharacterBody2D.set_floor_stop_on_slope_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_set_floor_stop_on_slope_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetFloorConstantSpeedEnabled(enabled bool) { //gd:CharacterBody2D.set_floor_constant_speed_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_set_floor_constant_speed_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsFloorConstantSpeedEnabled() bool { //gd:CharacterBody2D.is_floor_constant_speed_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_is_floor_constant_speed_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFloorBlockOnWallEnabled(enabled bool) { //gd:CharacterBody2D.set_floor_block_on_wall_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_set_floor_block_on_wall_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsFloorBlockOnWallEnabled() bool { //gd:CharacterBody2D.is_floor_block_on_wall_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_is_floor_block_on_wall_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSlideOnCeilingEnabled(enabled bool) { //gd:CharacterBody2D.set_slide_on_ceiling_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_set_slide_on_ceiling_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsSlideOnCeilingEnabled() bool { //gd:CharacterBody2D.is_slide_on_ceiling_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_is_slide_on_ceiling_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPlatformFloorLayers(exclude_layer gd.Int) { //gd:CharacterBody2D.set_platform_floor_layers
	var frame = callframe.New()
	callframe.Arg(frame, exclude_layer)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_set_platform_floor_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPlatformFloorLayers() gd.Int { //gd:CharacterBody2D.get_platform_floor_layers
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_get_platform_floor_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPlatformWallLayers(exclude_layer gd.Int) { //gd:CharacterBody2D.set_platform_wall_layers
	var frame = callframe.New()
	callframe.Arg(frame, exclude_layer)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_set_platform_wall_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPlatformWallLayers() gd.Int { //gd:CharacterBody2D.get_platform_wall_layers
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_get_platform_wall_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetMaxSlides() gd.Int { //gd:CharacterBody2D.get_max_slides
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_get_max_slides, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxSlides(max_slides gd.Int) { //gd:CharacterBody2D.set_max_slides
	var frame = callframe.New()
	callframe.Arg(frame, max_slides)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_set_max_slides, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFloorMaxAngle() gd.Float { //gd:CharacterBody2D.get_floor_max_angle
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_get_floor_max_angle, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFloorMaxAngle(radians gd.Float) { //gd:CharacterBody2D.set_floor_max_angle
	var frame = callframe.New()
	callframe.Arg(frame, radians)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_set_floor_max_angle, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFloorSnapLength() gd.Float { //gd:CharacterBody2D.get_floor_snap_length
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_get_floor_snap_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFloorSnapLength(floor_snap_length gd.Float) { //gd:CharacterBody2D.set_floor_snap_length
	var frame = callframe.New()
	callframe.Arg(frame, floor_snap_length)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_set_floor_snap_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetWallMinSlideAngle() gd.Float { //gd:CharacterBody2D.get_wall_min_slide_angle
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_get_wall_min_slide_angle, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWallMinSlideAngle(radians gd.Float) { //gd:CharacterBody2D.set_wall_min_slide_angle
	var frame = callframe.New()
	callframe.Arg(frame, radians)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_set_wall_min_slide_angle, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUpDirection() gd.Vector2 { //gd:CharacterBody2D.get_up_direction
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_get_up_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUpDirection(up_direction gd.Vector2) { //gd:CharacterBody2D.set_up_direction
	var frame = callframe.New()
	callframe.Arg(frame, up_direction)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_set_up_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetMotionMode(mode gdclass.CharacterBody2DMotionMode) { //gd:CharacterBody2D.set_motion_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_set_motion_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMotionMode() gdclass.CharacterBody2DMotionMode { //gd:CharacterBody2D.get_motion_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.CharacterBody2DMotionMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_get_motion_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPlatformOnLeave(on_leave_apply_velocity gdclass.CharacterBody2DPlatformOnLeave) { //gd:CharacterBody2D.set_platform_on_leave
	var frame = callframe.New()
	callframe.Arg(frame, on_leave_apply_velocity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_set_platform_on_leave, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPlatformOnLeave() gdclass.CharacterBody2DPlatformOnLeave { //gd:CharacterBody2D.get_platform_on_leave
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.CharacterBody2DPlatformOnLeave](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_get_platform_on_leave, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the body collided with the floor on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "floor" or not.
*/
//go:nosplit
func (self class) IsOnFloor() bool { //gd:CharacterBody2D.is_on_floor
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_is_on_floor, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the body collided only with the floor on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "floor" or not.
*/
//go:nosplit
func (self class) IsOnFloorOnly() bool { //gd:CharacterBody2D.is_on_floor_only
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_is_on_floor_only, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the body collided with the ceiling on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "ceiling" or not.
*/
//go:nosplit
func (self class) IsOnCeiling() bool { //gd:CharacterBody2D.is_on_ceiling
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_is_on_ceiling, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the body collided only with the ceiling on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "ceiling" or not.
*/
//go:nosplit
func (self class) IsOnCeilingOnly() bool { //gd:CharacterBody2D.is_on_ceiling_only
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_is_on_ceiling_only, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the body collided with a wall on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "wall" or not.
*/
//go:nosplit
func (self class) IsOnWall() bool { //gd:CharacterBody2D.is_on_wall
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_is_on_wall, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the body collided only with a wall on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "wall" or not.
*/
//go:nosplit
func (self class) IsOnWallOnly() bool { //gd:CharacterBody2D.is_on_wall_only
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_is_on_wall_only, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the collision normal of the floor at the last collision point. Only valid after calling [method move_and_slide] and when [method is_on_floor] returns [code]true[/code].
[b]Warning:[/b] The collision normal is not always the same as the surface normal.
*/
//go:nosplit
func (self class) GetFloorNormal() gd.Vector2 { //gd:CharacterBody2D.get_floor_normal
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_get_floor_normal, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the collision normal of the wall at the last collision point. Only valid after calling [method move_and_slide] and when [method is_on_wall] returns [code]true[/code].
[b]Warning:[/b] The collision normal is not always the same as the surface normal.
*/
//go:nosplit
func (self class) GetWallNormal() gd.Vector2 { //gd:CharacterBody2D.get_wall_normal
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_get_wall_normal, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the last motion applied to the [CharacterBody2D] during the last call to [method move_and_slide]. The movement can be split into multiple motions when sliding occurs, and this method return the last one, which is useful to retrieve the current direction of the movement.
*/
//go:nosplit
func (self class) GetLastMotion() gd.Vector2 { //gd:CharacterBody2D.get_last_motion
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_get_last_motion, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the travel (position delta) that occurred during the last call to [method move_and_slide].
*/
//go:nosplit
func (self class) GetPositionDelta() gd.Vector2 { //gd:CharacterBody2D.get_position_delta
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_get_position_delta, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the current real velocity since the last call to [method move_and_slide]. For example, when you climb a slope, you will move diagonally even though the velocity is horizontal. This method returns the diagonal movement, as opposed to [member velocity] which returns the requested velocity.
*/
//go:nosplit
func (self class) GetRealVelocity() gd.Vector2 { //gd:CharacterBody2D.get_real_velocity
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_get_real_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the floor's collision angle at the last collision point according to [param up_direction], which is [constant Vector2.UP] by default. This value is always positive and only valid after calling [method move_and_slide] and when [method is_on_floor] returns [code]true[/code].
*/
//go:nosplit
func (self class) GetFloorAngle(up_direction gd.Vector2) gd.Float { //gd:CharacterBody2D.get_floor_angle
	var frame = callframe.New()
	callframe.Arg(frame, up_direction)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_get_floor_angle, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the linear velocity of the platform at the last collision point. Only valid after calling [method move_and_slide].
*/
//go:nosplit
func (self class) GetPlatformVelocity() gd.Vector2 { //gd:CharacterBody2D.get_platform_velocity
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_get_platform_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of times the body collided and changed direction during the last call to [method move_and_slide].
*/
//go:nosplit
func (self class) GetSlideCollisionCount() gd.Int { //gd:CharacterBody2D.get_slide_collision_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_get_slide_collision_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a [KinematicCollision2D], which contains information about a collision that occurred during the last call to [method move_and_slide]. Since the body can collide several times in a single call to [method move_and_slide], you must specify the index of the collision in the range 0 to ([method get_slide_collision_count] - 1).
[b]Example usage:[/b]
[codeblocks]
[gdscript]
for i in get_slide_collision_count():
    var collision = get_slide_collision(i)
    print("Collided with: ", collision.get_collider().name)
[/gdscript]
[csharp]
for (int i = 0; i < GetSlideCollisionCount(); i++)
{
    KinematicCollision2D collision = GetSlideCollision(i);
    GD.Print("Collided with: ", (collision.GetCollider() as Node).Name);
}
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) GetSlideCollision(slide_idx gd.Int) [1]gdclass.KinematicCollision2D { //gd:CharacterBody2D.get_slide_collision
	var frame = callframe.New()
	callframe.Arg(frame, slide_idx)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_get_slide_collision, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.KinematicCollision2D{gd.PointerWithOwnershipTransferredToGo[gdclass.KinematicCollision2D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns a [KinematicCollision2D], which contains information about the latest collision that occurred during the last call to [method move_and_slide].
*/
//go:nosplit
func (self class) GetLastSlideCollision() [1]gdclass.KinematicCollision2D { //gd:CharacterBody2D.get_last_slide_collision
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody2D.Bind_get_last_slide_collision, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.KinematicCollision2D{gd.PointerWithOwnershipTransferredToGo[gdclass.KinematicCollision2D](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsCharacterBody2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCharacterBody2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsPhysicsBody2D() PhysicsBody2D.Advanced {
	return *((*PhysicsBody2D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPhysicsBody2D() PhysicsBody2D.Instance {
	return *((*PhysicsBody2D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsCollisionObject2D() CollisionObject2D.Advanced {
	return *((*CollisionObject2D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCollisionObject2D() CollisionObject2D.Instance {
	return *((*CollisionObject2D.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(PhysicsBody2D.Advanced(self.AsPhysicsBody2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(PhysicsBody2D.Instance(self.AsPhysicsBody2D()), name)
	}
}
func init() {
	gdclass.Register("CharacterBody2D", func(ptr gd.Object) any {
		return [1]gdclass.CharacterBody2D{*(*gdclass.CharacterBody2D)(unsafe.Pointer(&ptr))}
	})
}

type MotionMode = gdclass.CharacterBody2DMotionMode //gd:CharacterBody2D.MotionMode

const (
	/*Apply when notions of walls, ceiling and floor are relevant. In this mode the body motion will react to slopes (acceleration/slowdown). This mode is suitable for sided games like platformers.*/
	MotionModeGrounded MotionMode = 0
	/*Apply when there is no notion of floor or ceiling. All collisions will be reported as [code]on_wall[/code]. In this mode, when you slide, the speed will always be constant. This mode is suitable for top-down games.*/
	MotionModeFloating MotionMode = 1
)

type PlatformOnLeave = gdclass.CharacterBody2DPlatformOnLeave //gd:CharacterBody2D.PlatformOnLeave

const (
	/*Add the last platform velocity to the [member velocity] when you leave a moving platform.*/
	PlatformOnLeaveAddVelocity PlatformOnLeave = 0
	/*Add the last platform velocity to the [member velocity] when you leave a moving platform, but any downward motion is ignored. It's useful to keep full jump height even when the platform is moving down.*/
	PlatformOnLeaveAddUpwardVelocity PlatformOnLeave = 1
	/*Do nothing when leaving a platform.*/
	PlatformOnLeaveDoNothing PlatformOnLeave = 2
)
