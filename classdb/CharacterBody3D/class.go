// Package CharacterBody3D provides methods for working with CharacterBody3D object instances.
package CharacterBody3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/PhysicsBody3D"
import "graphics.gd/classdb/CollisionObject3D"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Vector3"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
[CharacterBody3D] is a specialized class for physics bodies that are meant to be user-controlled. They are not affected by physics at all, but they affect other physics bodies in their path. They are mainly used to provide high-level API to move objects with wall and slope detection ([method move_and_slide] method) in addition to the general collision detection provided by [method PhysicsBody3D.move_and_collide]. This makes it useful for highly configurable physics bodies that must move in specific ways and collide with the world, as is often the case with user-controlled characters.
For game objects that don't require complex movement or collision detection, such as moving platforms, [AnimatableBody3D] is simpler to configure.
*/
type Instance [1]gdclass.CharacterBody3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCharacterBody3D() Instance
}

/*
Moves the body based on [member velocity]. If the body collides with another, it will slide along the other body rather than stop immediately. If the other body is a [CharacterBody3D] or [RigidBody3D], it will also be affected by the motion of the other body. You can use this to make moving and rotating platforms, or to make nodes push other nodes.
Modifies [member velocity] if a slide collision occurred. To get the latest collision call [method get_last_slide_collision], for more detailed information about collisions that occurred, use [method get_slide_collision].
When the body touches a moving platform, the platform's velocity is automatically added to the body motion. If a collision occurs due to the platform's motion, it will always be first in the slide collisions.
Returns [code]true[/code] if the body collided, otherwise, returns [code]false[/code].
*/
func (self Instance) MoveAndSlide() bool {
	return bool(class(self).MoveAndSlide())
}

/*
Allows to manually apply a snap to the floor regardless of the body's velocity. This function does nothing when [method is_on_floor] returns [code]true[/code].
*/
func (self Instance) ApplyFloorSnap() {
	class(self).ApplyFloorSnap()
}

/*
Returns [code]true[/code] if the body collided with the floor on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "floor" or not.
*/
func (self Instance) IsOnFloor() bool {
	return bool(class(self).IsOnFloor())
}

/*
Returns [code]true[/code] if the body collided only with the floor on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "floor" or not.
*/
func (self Instance) IsOnFloorOnly() bool {
	return bool(class(self).IsOnFloorOnly())
}

/*
Returns [code]true[/code] if the body collided with the ceiling on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "ceiling" or not.
*/
func (self Instance) IsOnCeiling() bool {
	return bool(class(self).IsOnCeiling())
}

/*
Returns [code]true[/code] if the body collided only with the ceiling on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "ceiling" or not.
*/
func (self Instance) IsOnCeilingOnly() bool {
	return bool(class(self).IsOnCeilingOnly())
}

/*
Returns [code]true[/code] if the body collided with a wall on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "wall" or not.
*/
func (self Instance) IsOnWall() bool {
	return bool(class(self).IsOnWall())
}

/*
Returns [code]true[/code] if the body collided only with a wall on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "wall" or not.
*/
func (self Instance) IsOnWallOnly() bool {
	return bool(class(self).IsOnWallOnly())
}

/*
Returns the collision normal of the floor at the last collision point. Only valid after calling [method move_and_slide] and when [method is_on_floor] returns [code]true[/code].
[b]Warning:[/b] The collision normal is not always the same as the surface normal.
*/
func (self Instance) GetFloorNormal() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetFloorNormal())
}

/*
Returns the collision normal of the wall at the last collision point. Only valid after calling [method move_and_slide] and when [method is_on_wall] returns [code]true[/code].
[b]Warning:[/b] The collision normal is not always the same as the surface normal.
*/
func (self Instance) GetWallNormal() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetWallNormal())
}

/*
Returns the last motion applied to the [CharacterBody3D] during the last call to [method move_and_slide]. The movement can be split into multiple motions when sliding occurs, and this method return the last one, which is useful to retrieve the current direction of the movement.
*/
func (self Instance) GetLastMotion() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetLastMotion())
}

/*
Returns the travel (position delta) that occurred during the last call to [method move_and_slide].
*/
func (self Instance) GetPositionDelta() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetPositionDelta())
}

/*
Returns the current real velocity since the last call to [method move_and_slide]. For example, when you climb a slope, you will move diagonally even though the velocity is horizontal. This method returns the diagonal movement, as opposed to [member velocity] which returns the requested velocity.
*/
func (self Instance) GetRealVelocity() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetRealVelocity())
}

/*
Returns the floor's collision angle at the last collision point according to [param up_direction], which is [constant Vector3.UP] by default. This value is always positive and only valid after calling [method move_and_slide] and when [method is_on_floor] returns [code]true[/code].
*/
func (self Instance) GetFloorAngle() Float.X {
	return Float.X(Float.X(class(self).GetFloorAngle(gd.Vector3(gd.Vector3{0, 1, 0}))))
}

/*
Returns the linear velocity of the platform at the last collision point. Only valid after calling [method move_and_slide].
*/
func (self Instance) GetPlatformVelocity() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetPlatformVelocity())
}

/*
Returns the angular velocity of the platform at the last collision point. Only valid after calling [method move_and_slide].
*/
func (self Instance) GetPlatformAngularVelocity() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetPlatformAngularVelocity())
}

/*
Returns the number of times the body collided and changed direction during the last call to [method move_and_slide].
*/
func (self Instance) GetSlideCollisionCount() int {
	return int(int(class(self).GetSlideCollisionCount()))
}

/*
Returns a [KinematicCollision3D], which contains information about a collision that occurred during the last call to [method move_and_slide]. Since the body can collide several times in a single call to [method move_and_slide], you must specify the index of the collision in the range 0 to ([method get_slide_collision_count] - 1).
*/
func (self Instance) GetSlideCollision(slide_idx int) [1]gdclass.KinematicCollision3D {
	return [1]gdclass.KinematicCollision3D(class(self).GetSlideCollision(gd.Int(slide_idx)))
}

/*
Returns a [KinematicCollision3D], which contains information about the latest collision that occurred during the last call to [method move_and_slide].
*/
func (self Instance) GetLastSlideCollision() [1]gdclass.KinematicCollision3D {
	return [1]gdclass.KinematicCollision3D(class(self).GetLastSlideCollision())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.CharacterBody3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CharacterBody3D"))
	casted := Instance{*(*gdclass.CharacterBody3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) MotionMode() gdclass.CharacterBody3DMotionMode {
	return gdclass.CharacterBody3DMotionMode(class(self).GetMotionMode())
}

func (self Instance) SetMotionMode(value gdclass.CharacterBody3DMotionMode) {
	class(self).SetMotionMode(value)
}

func (self Instance) UpDirection() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetUpDirection())
}

func (self Instance) SetUpDirection(value Vector3.XYZ) {
	class(self).SetUpDirection(gd.Vector3(value))
}

func (self Instance) SlideOnCeiling() bool {
	return bool(class(self).IsSlideOnCeilingEnabled())
}

func (self Instance) SetSlideOnCeiling(value bool) {
	class(self).SetSlideOnCeilingEnabled(value)
}

func (self Instance) Velocity() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetVelocity())
}

func (self Instance) SetVelocity(value Vector3.XYZ) {
	class(self).SetVelocity(gd.Vector3(value))
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

func (self Instance) PlatformOnLeave() gdclass.CharacterBody3DPlatformOnLeave {
	return gdclass.CharacterBody3DPlatformOnLeave(class(self).GetPlatformOnLeave())
}

func (self Instance) SetPlatformOnLeave(value gdclass.CharacterBody3DPlatformOnLeave) {
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
Moves the body based on [member velocity]. If the body collides with another, it will slide along the other body rather than stop immediately. If the other body is a [CharacterBody3D] or [RigidBody3D], it will also be affected by the motion of the other body. You can use this to make moving and rotating platforms, or to make nodes push other nodes.
Modifies [member velocity] if a slide collision occurred. To get the latest collision call [method get_last_slide_collision], for more detailed information about collisions that occurred, use [method get_slide_collision].
When the body touches a moving platform, the platform's velocity is automatically added to the body motion. If a collision occurs due to the platform's motion, it will always be first in the slide collisions.
Returns [code]true[/code] if the body collided, otherwise, returns [code]false[/code].
*/
//go:nosplit
func (self class) MoveAndSlide() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_move_and_slide, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Allows to manually apply a snap to the floor regardless of the body's velocity. This function does nothing when [method is_on_floor] returns [code]true[/code].
*/
//go:nosplit
func (self class) ApplyFloorSnap() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_apply_floor_snap, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetVelocity(velocity gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, velocity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_set_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVelocity() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_get_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSafeMargin(margin gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_set_safe_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSafeMargin() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_get_safe_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) IsFloorStopOnSlopeEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_is_floor_stop_on_slope_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFloorStopOnSlopeEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_set_floor_stop_on_slope_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetFloorConstantSpeedEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_set_floor_constant_speed_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsFloorConstantSpeedEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_is_floor_constant_speed_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFloorBlockOnWallEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_set_floor_block_on_wall_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsFloorBlockOnWallEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_is_floor_block_on_wall_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSlideOnCeilingEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_set_slide_on_ceiling_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsSlideOnCeilingEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_is_slide_on_ceiling_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPlatformFloorLayers(exclude_layer gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, exclude_layer)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_set_platform_floor_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPlatformFloorLayers() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_get_platform_floor_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPlatformWallLayers(exclude_layer gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, exclude_layer)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_set_platform_wall_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPlatformWallLayers() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_get_platform_wall_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetMaxSlides() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_get_max_slides, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxSlides(max_slides gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, max_slides)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_set_max_slides, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFloorMaxAngle() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_get_floor_max_angle, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFloorMaxAngle(radians gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, radians)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_set_floor_max_angle, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFloorSnapLength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_get_floor_snap_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFloorSnapLength(floor_snap_length gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, floor_snap_length)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_set_floor_snap_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetWallMinSlideAngle() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_get_wall_min_slide_angle, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWallMinSlideAngle(radians gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, radians)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_set_wall_min_slide_angle, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUpDirection() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_get_up_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUpDirection(up_direction gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, up_direction)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_set_up_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetMotionMode(mode gdclass.CharacterBody3DMotionMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_set_motion_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMotionMode() gdclass.CharacterBody3DMotionMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.CharacterBody3DMotionMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_get_motion_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPlatformOnLeave(on_leave_apply_velocity gdclass.CharacterBody3DPlatformOnLeave) {
	var frame = callframe.New()
	callframe.Arg(frame, on_leave_apply_velocity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_set_platform_on_leave, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPlatformOnLeave() gdclass.CharacterBody3DPlatformOnLeave {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.CharacterBody3DPlatformOnLeave](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_get_platform_on_leave, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the body collided with the floor on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "floor" or not.
*/
//go:nosplit
func (self class) IsOnFloor() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_is_on_floor, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the body collided only with the floor on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "floor" or not.
*/
//go:nosplit
func (self class) IsOnFloorOnly() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_is_on_floor_only, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the body collided with the ceiling on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "ceiling" or not.
*/
//go:nosplit
func (self class) IsOnCeiling() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_is_on_ceiling, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the body collided only with the ceiling on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "ceiling" or not.
*/
//go:nosplit
func (self class) IsOnCeilingOnly() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_is_on_ceiling_only, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the body collided with a wall on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "wall" or not.
*/
//go:nosplit
func (self class) IsOnWall() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_is_on_wall, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the body collided only with a wall on the last call of [method move_and_slide]. Otherwise, returns [code]false[/code]. The [member up_direction] and [member floor_max_angle] are used to determine whether a surface is "wall" or not.
*/
//go:nosplit
func (self class) IsOnWallOnly() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_is_on_wall_only, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the collision normal of the floor at the last collision point. Only valid after calling [method move_and_slide] and when [method is_on_floor] returns [code]true[/code].
[b]Warning:[/b] The collision normal is not always the same as the surface normal.
*/
//go:nosplit
func (self class) GetFloorNormal() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_get_floor_normal, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the collision normal of the wall at the last collision point. Only valid after calling [method move_and_slide] and when [method is_on_wall] returns [code]true[/code].
[b]Warning:[/b] The collision normal is not always the same as the surface normal.
*/
//go:nosplit
func (self class) GetWallNormal() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_get_wall_normal, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the last motion applied to the [CharacterBody3D] during the last call to [method move_and_slide]. The movement can be split into multiple motions when sliding occurs, and this method return the last one, which is useful to retrieve the current direction of the movement.
*/
//go:nosplit
func (self class) GetLastMotion() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_get_last_motion, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the travel (position delta) that occurred during the last call to [method move_and_slide].
*/
//go:nosplit
func (self class) GetPositionDelta() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_get_position_delta, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the current real velocity since the last call to [method move_and_slide]. For example, when you climb a slope, you will move diagonally even though the velocity is horizontal. This method returns the diagonal movement, as opposed to [member velocity] which returns the requested velocity.
*/
//go:nosplit
func (self class) GetRealVelocity() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_get_real_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the floor's collision angle at the last collision point according to [param up_direction], which is [constant Vector3.UP] by default. This value is always positive and only valid after calling [method move_and_slide] and when [method is_on_floor] returns [code]true[/code].
*/
//go:nosplit
func (self class) GetFloorAngle(up_direction gd.Vector3) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, up_direction)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_get_floor_angle, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the linear velocity of the platform at the last collision point. Only valid after calling [method move_and_slide].
*/
//go:nosplit
func (self class) GetPlatformVelocity() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_get_platform_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the angular velocity of the platform at the last collision point. Only valid after calling [method move_and_slide].
*/
//go:nosplit
func (self class) GetPlatformAngularVelocity() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_get_platform_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of times the body collided and changed direction during the last call to [method move_and_slide].
*/
//go:nosplit
func (self class) GetSlideCollisionCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_get_slide_collision_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a [KinematicCollision3D], which contains information about a collision that occurred during the last call to [method move_and_slide]. Since the body can collide several times in a single call to [method move_and_slide], you must specify the index of the collision in the range 0 to ([method get_slide_collision_count] - 1).
*/
//go:nosplit
func (self class) GetSlideCollision(slide_idx gd.Int) [1]gdclass.KinematicCollision3D {
	var frame = callframe.New()
	callframe.Arg(frame, slide_idx)
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_get_slide_collision, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.KinematicCollision3D{gd.PointerWithOwnershipTransferredToGo[gdclass.KinematicCollision3D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns a [KinematicCollision3D], which contains information about the latest collision that occurred during the last call to [method move_and_slide].
*/
//go:nosplit
func (self class) GetLastSlideCollision() [1]gdclass.KinematicCollision3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharacterBody3D.Bind_get_last_slide_collision, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.KinematicCollision3D{gd.PointerWithOwnershipTransferredToGo[gdclass.KinematicCollision3D](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsCharacterBody3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCharacterBody3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsPhysicsBody3D() PhysicsBody3D.Advanced {
	return *((*PhysicsBody3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPhysicsBody3D() PhysicsBody3D.Instance {
	return *((*PhysicsBody3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsCollisionObject3D() CollisionObject3D.Advanced {
	return *((*CollisionObject3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCollisionObject3D() CollisionObject3D.Instance {
	return *((*CollisionObject3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(PhysicsBody3D.Advanced(self.AsPhysicsBody3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(PhysicsBody3D.Instance(self.AsPhysicsBody3D()), name)
	}
}
func init() {
	gdclass.Register("CharacterBody3D", func(ptr gd.Object) any {
		return [1]gdclass.CharacterBody3D{*(*gdclass.CharacterBody3D)(unsafe.Pointer(&ptr))}
	})
}

type MotionMode = gdclass.CharacterBody3DMotionMode

const (
	/*Apply when notions of walls, ceiling and floor are relevant. In this mode the body motion will react to slopes (acceleration/slowdown). This mode is suitable for grounded games like platformers.*/
	MotionModeGrounded MotionMode = 0
	/*Apply when there is no notion of floor or ceiling. All collisions will be reported as [code]on_wall[/code]. In this mode, when you slide, the speed will always be constant. This mode is suitable for games without ground like space games.*/
	MotionModeFloating MotionMode = 1
)

type PlatformOnLeave = gdclass.CharacterBody3DPlatformOnLeave

const (
	/*Add the last platform velocity to the [member velocity] when you leave a moving platform.*/
	PlatformOnLeaveAddVelocity PlatformOnLeave = 0
	/*Add the last platform velocity to the [member velocity] when you leave a moving platform, but any downward motion is ignored. It's useful to keep full jump height even when the platform is moving down.*/
	PlatformOnLeaveAddUpwardVelocity PlatformOnLeave = 1
	/*Do nothing when leaving a platform.*/
	PlatformOnLeaveDoNothing PlatformOnLeave = 2
)
