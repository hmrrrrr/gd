// Package PhysicsServer2DExtension provides methods for working with PhysicsServer2DExtension object instances.
package PhysicsServer2DExtension

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
import "graphics.gd/variant/RID"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Transform2D"
import "graphics.gd/variant/Vector2"

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

/*
This class extends [PhysicsServer2D] by providing additional virtual methods that can be overridden. When these methods are overridden, they will be called instead of the internal methods of the physics server.
Intended for use with GDExtension to create custom implementations of [PhysicsServer2D].

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=PhysicsServer2DExtension)
*/
type Instance [1]gdclass.PhysicsServer2DExtension

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPhysicsServer2DExtension() Instance
}
type Interface interface {
	//Overridable version of [method PhysicsServer2D.world_boundary_shape_create].
	WorldBoundaryShapeCreate() RID.Any
	//Overridable version of [method PhysicsServer2D.separation_ray_shape_create].
	SeparationRayShapeCreate() RID.Any
	//Overridable version of [method PhysicsServer2D.segment_shape_create].
	SegmentShapeCreate() RID.Any
	//Overridable version of [method PhysicsServer2D.circle_shape_create].
	CircleShapeCreate() RID.Any
	//Overridable version of [method PhysicsServer2D.rectangle_shape_create].
	RectangleShapeCreate() RID.Any
	//Overridable version of [method PhysicsServer2D.capsule_shape_create].
	CapsuleShapeCreate() RID.Any
	//Overridable version of [method PhysicsServer2D.convex_polygon_shape_create].
	ConvexPolygonShapeCreate() RID.Any
	//Overridable version of [method PhysicsServer2D.concave_polygon_shape_create].
	ConcavePolygonShapeCreate() RID.Any
	//Overridable version of [method PhysicsServer2D.shape_set_data].
	ShapeSetData(shape RID.Any, data any)
	//Should set the custom solver bias for the given [param shape]. It defines how much bodies are forced to separate on contact.
	//Overridable version of [PhysicsServer2D]'s internal [code]shape_get_custom_solver_bias[/code] method. Corresponds to [member Shape2D.custom_solver_bias].
	ShapeSetCustomSolverBias(shape RID.Any, bias Float.X)
	//Overridable version of [method PhysicsServer2D.shape_get_type].
	ShapeGetType(shape RID.Any) gdclass.PhysicsServer2DShapeType
	//Overridable version of [method PhysicsServer2D.shape_get_data].
	ShapeGetData(shape RID.Any) any
	//Should return the custom solver bias of the given [param shape], which defines how much bodies are forced to separate on contact when this shape is involved.
	//Overridable version of [PhysicsServer2D]'s internal [code]shape_get_custom_solver_bias[/code] method. Corresponds to [member Shape2D.custom_solver_bias].
	ShapeGetCustomSolverBias(shape RID.Any) Float.X
	//Given two shapes and their parameters, should return [code]true[/code] if a collision between the two would occur, with additional details passed in [param results].
	//Overridable version of [PhysicsServer2D]'s internal [code]shape_collide[/code] method. Corresponds to [method PhysicsDirectSpaceState2D.collide_shape].
	ShapeCollide(shape_A RID.Any, xform_A Transform2D.OriginXY, motion_A Vector2.XY, shape_B RID.Any, xform_B Transform2D.OriginXY, motion_B Vector2.XY, results unsafe.Pointer, result_max int, result_count *int32) bool
	//Overridable version of [method PhysicsServer2D.space_create].
	SpaceCreate() RID.Any
	//Overridable version of [method PhysicsServer2D.space_set_active].
	SpaceSetActive(space RID.Any, active bool)
	//Overridable version of [method PhysicsServer2D.space_is_active].
	SpaceIsActive(space RID.Any) bool
	//Overridable version of [method PhysicsServer2D.space_set_param].
	SpaceSetParam(space RID.Any, param gdclass.PhysicsServer2DSpaceParameter, value Float.X)
	//Overridable version of [method PhysicsServer2D.space_get_param].
	SpaceGetParam(space RID.Any, param gdclass.PhysicsServer2DSpaceParameter) Float.X
	//Overridable version of [method PhysicsServer2D.space_get_direct_state].
	SpaceGetDirectState(space RID.Any) [1]gdclass.PhysicsDirectSpaceState2D
	//Used internally to allow the given [param space] to store contact points, up to [param max_contacts]. This is automatically set for the main [World2D]'s space when [member SceneTree.debug_collisions_hint] is [code]true[/code], or by checking "Visible Collision Shapes" in the editor. Only works in debug builds.
	//Overridable version of [PhysicsServer2D]'s internal [code]space_set_debug_contacts[/code] method.
	SpaceSetDebugContacts(space RID.Any, max_contacts int)
	//Should return the positions of all contacts that have occurred during the last physics step in the given [param space]. See also [method _space_get_contact_count] and [method _space_set_debug_contacts].
	//Overridable version of [PhysicsServer2D]'s internal [code]space_get_contacts[/code] method.
	SpaceGetContacts(space RID.Any) []Vector2.XY
	//Should return how many contacts have occurred during the last physics step in the given [param space]. See also [method _space_get_contacts] and [method _space_set_debug_contacts].
	//Overridable version of [PhysicsServer2D]'s internal [code]space_get_contact_count[/code] method.
	SpaceGetContactCount(space RID.Any) int
	//Overridable version of [method PhysicsServer2D.area_create].
	AreaCreate() RID.Any
	//Overridable version of [method PhysicsServer2D.area_set_space].
	AreaSetSpace(area RID.Any, space RID.Any)
	//Overridable version of [method PhysicsServer2D.area_get_space].
	AreaGetSpace(area RID.Any) RID.Any
	//Overridable version of [method PhysicsServer2D.area_add_shape].
	AreaAddShape(area RID.Any, shape RID.Any, transform Transform2D.OriginXY, disabled bool)
	//Overridable version of [method PhysicsServer2D.area_set_shape].
	AreaSetShape(area RID.Any, shape_idx int, shape RID.Any)
	//Overridable version of [method PhysicsServer2D.area_set_shape_transform].
	AreaSetShapeTransform(area RID.Any, shape_idx int, transform Transform2D.OriginXY)
	//Overridable version of [method PhysicsServer2D.area_set_shape_disabled].
	AreaSetShapeDisabled(area RID.Any, shape_idx int, disabled bool)
	//Overridable version of [method PhysicsServer2D.area_get_shape_count].
	AreaGetShapeCount(area RID.Any) int
	//Overridable version of [method PhysicsServer2D.area_get_shape].
	AreaGetShape(area RID.Any, shape_idx int) RID.Any
	//Overridable version of [method PhysicsServer2D.area_get_shape_transform].
	AreaGetShapeTransform(area RID.Any, shape_idx int) Transform2D.OriginXY
	//Overridable version of [method PhysicsServer2D.area_remove_shape].
	AreaRemoveShape(area RID.Any, shape_idx int)
	//Overridable version of [method PhysicsServer2D.area_clear_shapes].
	AreaClearShapes(area RID.Any)
	//Overridable version of [method PhysicsServer2D.area_attach_object_instance_id].
	AreaAttachObjectInstanceId(area RID.Any, id int)
	//Overridable version of [method PhysicsServer2D.area_get_object_instance_id].
	AreaGetObjectInstanceId(area RID.Any) int
	//Overridable version of [method PhysicsServer2D.area_attach_canvas_instance_id].
	AreaAttachCanvasInstanceId(area RID.Any, id int)
	//Overridable version of [method PhysicsServer2D.area_get_canvas_instance_id].
	AreaGetCanvasInstanceId(area RID.Any) int
	//Overridable version of [method PhysicsServer2D.area_set_param].
	AreaSetParam(area RID.Any, param gdclass.PhysicsServer2DAreaParameter, value any)
	//Overridable version of [method PhysicsServer2D.area_set_transform].
	AreaSetTransform(area RID.Any, transform Transform2D.OriginXY)
	//Overridable version of [method PhysicsServer2D.area_get_param].
	AreaGetParam(area RID.Any, param gdclass.PhysicsServer2DAreaParameter) any
	//Overridable version of [method PhysicsServer2D.area_get_transform].
	AreaGetTransform(area RID.Any) Transform2D.OriginXY
	//Overridable version of [method PhysicsServer2D.area_set_collision_layer].
	AreaSetCollisionLayer(area RID.Any, layer int)
	//Overridable version of [method PhysicsServer2D.area_get_collision_layer].
	AreaGetCollisionLayer(area RID.Any) int
	//Overridable version of [method PhysicsServer2D.area_set_collision_mask].
	AreaSetCollisionMask(area RID.Any, mask int)
	//Overridable version of [method PhysicsServer2D.area_get_collision_mask].
	AreaGetCollisionMask(area RID.Any) int
	//Overridable version of [method PhysicsServer2D.area_set_monitorable].
	AreaSetMonitorable(area RID.Any, monitorable bool)
	//If set to [code]true[/code], allows the area with the given [RID] to detect mouse inputs when the mouse cursor is hovering on it.
	//Overridable version of [PhysicsServer2D]'s internal [code]area_set_pickable[/code] method. Corresponds to [member CollisionObject2D.input_pickable].
	AreaSetPickable(area RID.Any, pickable bool)
	//Overridable version of [method PhysicsServer2D.area_set_monitor_callback].
	AreaSetMonitorCallback(area RID.Any, callback Callable.Function)
	//Overridable version of [method PhysicsServer2D.area_set_area_monitor_callback].
	AreaSetAreaMonitorCallback(area RID.Any, callback Callable.Function)
	//Overridable version of [method PhysicsServer2D.body_create].
	BodyCreate() RID.Any
	//Overridable version of [method PhysicsServer2D.body_set_space].
	BodySetSpace(body RID.Any, space RID.Any)
	//Overridable version of [method PhysicsServer2D.body_get_space].
	BodyGetSpace(body RID.Any) RID.Any
	//Overridable version of [method PhysicsServer2D.body_set_mode].
	BodySetMode(body RID.Any, mode gdclass.PhysicsServer2DBodyMode)
	//Overridable version of [method PhysicsServer2D.body_get_mode].
	BodyGetMode(body RID.Any) gdclass.PhysicsServer2DBodyMode
	//Overridable version of [method PhysicsServer2D.body_add_shape].
	BodyAddShape(body RID.Any, shape RID.Any, transform Transform2D.OriginXY, disabled bool)
	//Overridable version of [method PhysicsServer2D.body_set_shape].
	BodySetShape(body RID.Any, shape_idx int, shape RID.Any)
	//Overridable version of [method PhysicsServer2D.body_set_shape_transform].
	BodySetShapeTransform(body RID.Any, shape_idx int, transform Transform2D.OriginXY)
	//Overridable version of [method PhysicsServer2D.body_get_shape_count].
	BodyGetShapeCount(body RID.Any) int
	//Overridable version of [method PhysicsServer2D.body_get_shape].
	BodyGetShape(body RID.Any, shape_idx int) RID.Any
	//Overridable version of [method PhysicsServer2D.body_get_shape_transform].
	BodyGetShapeTransform(body RID.Any, shape_idx int) Transform2D.OriginXY
	//Overridable version of [method PhysicsServer2D.body_set_shape_disabled].
	BodySetShapeDisabled(body RID.Any, shape_idx int, disabled bool)
	//Overridable version of [method PhysicsServer2D.body_set_shape_as_one_way_collision].
	BodySetShapeAsOneWayCollision(body RID.Any, shape_idx int, enable bool, margin Float.X)
	//Overridable version of [method PhysicsServer2D.body_remove_shape].
	BodyRemoveShape(body RID.Any, shape_idx int)
	//Overridable version of [method PhysicsServer2D.body_clear_shapes].
	BodyClearShapes(body RID.Any)
	//Overridable version of [method PhysicsServer2D.body_attach_object_instance_id].
	BodyAttachObjectInstanceId(body RID.Any, id int)
	//Overridable version of [method PhysicsServer2D.body_get_object_instance_id].
	BodyGetObjectInstanceId(body RID.Any) int
	//Overridable version of [method PhysicsServer2D.body_attach_canvas_instance_id].
	BodyAttachCanvasInstanceId(body RID.Any, id int)
	//Overridable version of [method PhysicsServer2D.body_get_canvas_instance_id].
	BodyGetCanvasInstanceId(body RID.Any) int
	//Overridable version of [method PhysicsServer2D.body_set_continuous_collision_detection_mode].
	BodySetContinuousCollisionDetectionMode(body RID.Any, mode gdclass.PhysicsServer2DCCDMode)
	//Overridable version of [method PhysicsServer2D.body_get_continuous_collision_detection_mode].
	BodyGetContinuousCollisionDetectionMode(body RID.Any) gdclass.PhysicsServer2DCCDMode
	//Overridable version of [method PhysicsServer2D.body_set_collision_layer].
	BodySetCollisionLayer(body RID.Any, layer int)
	//Overridable version of [method PhysicsServer2D.body_get_collision_layer].
	BodyGetCollisionLayer(body RID.Any) int
	//Overridable version of [method PhysicsServer2D.body_set_collision_mask].
	BodySetCollisionMask(body RID.Any, mask int)
	//Overridable version of [method PhysicsServer2D.body_get_collision_mask].
	BodyGetCollisionMask(body RID.Any) int
	//Overridable version of [method PhysicsServer2D.body_set_collision_priority].
	BodySetCollisionPriority(body RID.Any, priority Float.X)
	//Overridable version of [method PhysicsServer2D.body_get_collision_priority].
	BodyGetCollisionPriority(body RID.Any) Float.X
	//Overridable version of [method PhysicsServer2D.body_set_param].
	BodySetParam(body RID.Any, param gdclass.PhysicsServer2DBodyParameter, value any)
	//Overridable version of [method PhysicsServer2D.body_get_param].
	BodyGetParam(body RID.Any, param gdclass.PhysicsServer2DBodyParameter) any
	//Overridable version of [method PhysicsServer2D.body_reset_mass_properties].
	BodyResetMassProperties(body RID.Any)
	//Overridable version of [method PhysicsServer2D.body_set_state].
	BodySetState(body RID.Any, state gdclass.PhysicsServer2DBodyState, value any)
	//Overridable version of [method PhysicsServer2D.body_get_state].
	BodyGetState(body RID.Any, state gdclass.PhysicsServer2DBodyState) any
	//Overridable version of [method PhysicsServer2D.body_apply_central_impulse].
	BodyApplyCentralImpulse(body RID.Any, impulse Vector2.XY)
	//Overridable version of [method PhysicsServer2D.body_apply_torque_impulse].
	BodyApplyTorqueImpulse(body RID.Any, impulse Float.X)
	//Overridable version of [method PhysicsServer2D.body_apply_impulse].
	BodyApplyImpulse(body RID.Any, impulse Vector2.XY, position Vector2.XY)
	//Overridable version of [method PhysicsServer2D.body_apply_central_force].
	BodyApplyCentralForce(body RID.Any, force Vector2.XY)
	//Overridable version of [method PhysicsServer2D.body_apply_force].
	BodyApplyForce(body RID.Any, force Vector2.XY, position Vector2.XY)
	//Overridable version of [method PhysicsServer2D.body_apply_torque].
	BodyApplyTorque(body RID.Any, torque Float.X)
	//Overridable version of [method PhysicsServer2D.body_add_constant_central_force].
	BodyAddConstantCentralForce(body RID.Any, force Vector2.XY)
	//Overridable version of [method PhysicsServer2D.body_add_constant_force].
	BodyAddConstantForce(body RID.Any, force Vector2.XY, position Vector2.XY)
	//Overridable version of [method PhysicsServer2D.body_add_constant_torque].
	BodyAddConstantTorque(body RID.Any, torque Float.X)
	//Overridable version of [method PhysicsServer2D.body_set_constant_force].
	BodySetConstantForce(body RID.Any, force Vector2.XY)
	//Overridable version of [method PhysicsServer2D.body_get_constant_force].
	BodyGetConstantForce(body RID.Any) Vector2.XY
	//Overridable version of [method PhysicsServer2D.body_set_constant_torque].
	BodySetConstantTorque(body RID.Any, torque Float.X)
	//Overridable version of [method PhysicsServer2D.body_get_constant_torque].
	BodyGetConstantTorque(body RID.Any) Float.X
	//Overridable version of [method PhysicsServer2D.body_set_axis_velocity].
	BodySetAxisVelocity(body RID.Any, axis_velocity Vector2.XY)
	//Overridable version of [method PhysicsServer2D.body_add_collision_exception].
	BodyAddCollisionException(body RID.Any, excepted_body RID.Any)
	//Overridable version of [method PhysicsServer2D.body_remove_collision_exception].
	BodyRemoveCollisionException(body RID.Any, excepted_body RID.Any)
	//Returns the [RID]s of all bodies added as collision exceptions for the given [param body]. See also [method _body_add_collision_exception] and [method _body_remove_collision_exception].
	//Overridable version of [PhysicsServer2D]'s internal [code]body_get_collision_exceptions[/code] method. Corresponds to [method PhysicsBody2D.get_collision_exceptions].
	BodyGetCollisionExceptions(body RID.Any) []RID.Any
	//Overridable version of [method PhysicsServer2D.body_set_max_contacts_reported].
	BodySetMaxContactsReported(body RID.Any, amount int)
	//Overridable version of [method PhysicsServer2D.body_get_max_contacts_reported].
	BodyGetMaxContactsReported(body RID.Any) int
	//Overridable version of [PhysicsServer2D]'s internal [code]body_set_contacts_reported_depth_threshold[/code] method.
	//[b]Note:[/b] This method is currently unused by Godot's default physics implementation.
	BodySetContactsReportedDepthThreshold(body RID.Any, threshold Float.X)
	//Overridable version of [PhysicsServer2D]'s internal [code]body_get_contacts_reported_depth_threshold[/code] method.
	//[b]Note:[/b] This method is currently unused by Godot's default physics implementation.
	BodyGetContactsReportedDepthThreshold(body RID.Any) Float.X
	//Overridable version of [method PhysicsServer2D.body_set_omit_force_integration].
	BodySetOmitForceIntegration(body RID.Any, enable bool)
	//Overridable version of [method PhysicsServer2D.body_is_omitting_force_integration].
	BodyIsOmittingForceIntegration(body RID.Any) bool
	//Assigns the [param body] to call the given [param callable] during the synchronization phase of the loop, before [method _step] is called. See also [method _sync].
	//Overridable version of [method PhysicsServer2D.body_set_state_sync_callback].
	BodySetStateSyncCallback(body RID.Any, callable Callable.Function)
	//Overridable version of [method PhysicsServer2D.body_set_force_integration_callback].
	BodySetForceIntegrationCallback(body RID.Any, callable Callable.Function, userdata any)
	//Given a [param body], a [param shape], and their respective parameters, this method should return [code]true[/code] if a collision between the two would occur, with additional details passed in [param results].
	//Overridable version of [PhysicsServer2D]'s internal [code]shape_collide[/code] method. Corresponds to [method PhysicsDirectSpaceState2D.collide_shape].
	BodyCollideShape(body RID.Any, body_shape int, shape RID.Any, shape_xform Transform2D.OriginXY, motion Vector2.XY, results unsafe.Pointer, result_max int, result_count *int32) bool
	//If set to [code]true[/code], allows the body with the given [RID] to detect mouse inputs when the mouse cursor is hovering on it.
	//Overridable version of [PhysicsServer2D]'s internal [code]body_set_pickable[/code] method. Corresponds to [member CollisionObject2D.input_pickable].
	BodySetPickable(body RID.Any, pickable bool)
	//Overridable version of [method PhysicsServer2D.body_get_direct_state].
	BodyGetDirectState(body RID.Any) [1]gdclass.PhysicsDirectBodyState2D
	//Overridable version of [method PhysicsServer2D.body_test_motion]. Unlike the exposed implementation, this method does not receive all of the arguments inside a [PhysicsTestMotionParameters2D].
	BodyTestMotion(body RID.Any, from Transform2D.OriginXY, motion Vector2.XY, margin Float.X, collide_separation_ray bool, recovery_as_collision bool, result *MotionResult) bool
	//Overridable version of [method PhysicsServer2D.joint_create].
	JointCreate() RID.Any
	//Overridable version of [method PhysicsServer2D.joint_clear].
	JointClear(joint RID.Any)
	//Overridable version of [method PhysicsServer2D.joint_set_param].
	JointSetParam(joint RID.Any, param gdclass.PhysicsServer2DJointParam, value Float.X)
	//Overridable version of [method PhysicsServer2D.joint_get_param].
	JointGetParam(joint RID.Any, param gdclass.PhysicsServer2DJointParam) Float.X
	//Overridable version of [method PhysicsServer2D.joint_disable_collisions_between_bodies].
	JointDisableCollisionsBetweenBodies(joint RID.Any, disable bool)
	//Overridable version of [method PhysicsServer2D.joint_is_disabled_collisions_between_bodies].
	JointIsDisabledCollisionsBetweenBodies(joint RID.Any) bool
	//Overridable version of [method PhysicsServer2D.joint_make_pin].
	JointMakePin(joint RID.Any, anchor Vector2.XY, body_a RID.Any, body_b RID.Any)
	//Overridable version of [method PhysicsServer2D.joint_make_groove].
	JointMakeGroove(joint RID.Any, a_groove1 Vector2.XY, a_groove2 Vector2.XY, b_anchor Vector2.XY, body_a RID.Any, body_b RID.Any)
	//Overridable version of [method PhysicsServer2D.joint_make_damped_spring].
	JointMakeDampedSpring(joint RID.Any, anchor_a Vector2.XY, anchor_b Vector2.XY, body_a RID.Any, body_b RID.Any)
	//Overridable version of [method PhysicsServer2D.pin_joint_set_flag].
	PinJointSetFlag(joint RID.Any, flag gdclass.PhysicsServer2DPinJointFlag, enabled bool)
	//Overridable version of [method PhysicsServer2D.pin_joint_get_flag].
	PinJointGetFlag(joint RID.Any, flag gdclass.PhysicsServer2DPinJointFlag) bool
	//Overridable version of [method PhysicsServer2D.pin_joint_set_param].
	PinJointSetParam(joint RID.Any, param gdclass.PhysicsServer2DPinJointParam, value Float.X)
	//Overridable version of [method PhysicsServer2D.pin_joint_get_param].
	PinJointGetParam(joint RID.Any, param gdclass.PhysicsServer2DPinJointParam) Float.X
	//Overridable version of [method PhysicsServer2D.damped_spring_joint_set_param].
	DampedSpringJointSetParam(joint RID.Any, param gdclass.PhysicsServer2DDampedSpringParam, value Float.X)
	//Overridable version of [method PhysicsServer2D.damped_spring_joint_get_param].
	DampedSpringJointGetParam(joint RID.Any, param gdclass.PhysicsServer2DDampedSpringParam) Float.X
	//Overridable version of [method PhysicsServer2D.joint_get_type].
	JointGetType(joint RID.Any) gdclass.PhysicsServer2DJointType
	//Overridable version of [method PhysicsServer2D.free_rid].
	FreeRid(rid RID.Any)
	//Overridable version of [method PhysicsServer2D.set_active].
	SetActive(active bool)
	//Called when the main loop is initialized and creates a new instance of this physics server. See also [method MainLoop._initialize] and [method _finish].
	//Overridable version of [PhysicsServer2D]'s internal [code]init[/code] method.
	Init()
	//Called every physics step to process the physics simulation. [param step] is the time elapsed since the last physics step, in seconds. It is usually the same as [method Node.get_physics_process_delta_time].
	//Overridable version of [PhysicsServer2D]'s internal [code skip-lint]step[/code] method.
	Step(step Float.X)
	//Called to indicate that the physics server is synchronizing and cannot access physics states if running on a separate thread. See also [method _end_sync].
	//Overridable version of [PhysicsServer2D]'s internal [code]sync[/code] method.
	Sync()
	//Called every physics step before [method _step] to process all remaining queries.
	//Overridable version of [PhysicsServer2D]'s internal [code]flush_queries[/code] method.
	FlushQueries()
	//Called to indicate that the physics server has stopped synchronizing. It is in the loop's iteration/physics phase, and can access physics objects even if running on a separate thread. See also [method _sync].
	//Overridable version of [PhysicsServer2D]'s internal [code]end_sync[/code] method.
	EndSync()
	//Called when the main loop finalizes to shut down the physics server. See also [method MainLoop._finalize] and [method _init].
	//Overridable version of [PhysicsServer2D]'s internal [code]finish[/code] method.
	Finish()
	//Overridable method that should return [code]true[/code] when the physics server is processing queries. See also [method _flush_queries].
	//Overridable version of [PhysicsServer2D]'s internal [code]is_flushing_queries[/code] method.
	IsFlushingQueries() bool
	//Overridable version of [method PhysicsServer2D.get_process_info].
	GetProcessInfo(process_info gdclass.PhysicsServer2DProcessInfo) int
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) WorldBoundaryShapeCreate() (_ RID.Any)                           { return }
func (self implementation) SeparationRayShapeCreate() (_ RID.Any)                           { return }
func (self implementation) SegmentShapeCreate() (_ RID.Any)                                 { return }
func (self implementation) CircleShapeCreate() (_ RID.Any)                                  { return }
func (self implementation) RectangleShapeCreate() (_ RID.Any)                               { return }
func (self implementation) CapsuleShapeCreate() (_ RID.Any)                                 { return }
func (self implementation) ConvexPolygonShapeCreate() (_ RID.Any)                           { return }
func (self implementation) ConcavePolygonShapeCreate() (_ RID.Any)                          { return }
func (self implementation) ShapeSetData(shape RID.Any, data any)                            { return }
func (self implementation) ShapeSetCustomSolverBias(shape RID.Any, bias Float.X)            { return }
func (self implementation) ShapeGetType(shape RID.Any) (_ gdclass.PhysicsServer2DShapeType) { return }
func (self implementation) ShapeGetData(shape RID.Any) (_ any)                              { return }
func (self implementation) ShapeGetCustomSolverBias(shape RID.Any) (_ Float.X)              { return }
func (self implementation) ShapeCollide(shape_A RID.Any, xform_A Transform2D.OriginXY, motion_A Vector2.XY, shape_B RID.Any, xform_B Transform2D.OriginXY, motion_B Vector2.XY, results unsafe.Pointer, result_max int, result_count *int32) (_ bool) {
	return
}
func (self implementation) SpaceCreate() (_ RID.Any)                  { return }
func (self implementation) SpaceSetActive(space RID.Any, active bool) { return }
func (self implementation) SpaceIsActive(space RID.Any) (_ bool)      { return }
func (self implementation) SpaceSetParam(space RID.Any, param gdclass.PhysicsServer2DSpaceParameter, value Float.X) {
	return
}
func (self implementation) SpaceGetParam(space RID.Any, param gdclass.PhysicsServer2DSpaceParameter) (_ Float.X) {
	return
}
func (self implementation) SpaceGetDirectState(space RID.Any) (_ [1]gdclass.PhysicsDirectSpaceState2D) {
	return
}
func (self implementation) SpaceSetDebugContacts(space RID.Any, max_contacts int) { return }
func (self implementation) SpaceGetContacts(space RID.Any) (_ []Vector2.XY)       { return }
func (self implementation) SpaceGetContactCount(space RID.Any) (_ int)            { return }
func (self implementation) AreaCreate() (_ RID.Any)                               { return }
func (self implementation) AreaSetSpace(area RID.Any, space RID.Any)              { return }
func (self implementation) AreaGetSpace(area RID.Any) (_ RID.Any)                 { return }
func (self implementation) AreaAddShape(area RID.Any, shape RID.Any, transform Transform2D.OriginXY, disabled bool) {
	return
}
func (self implementation) AreaSetShape(area RID.Any, shape_idx int, shape RID.Any) { return }
func (self implementation) AreaSetShapeTransform(area RID.Any, shape_idx int, transform Transform2D.OriginXY) {
	return
}
func (self implementation) AreaSetShapeDisabled(area RID.Any, shape_idx int, disabled bool) { return }
func (self implementation) AreaGetShapeCount(area RID.Any) (_ int)                          { return }
func (self implementation) AreaGetShape(area RID.Any, shape_idx int) (_ RID.Any)            { return }
func (self implementation) AreaGetShapeTransform(area RID.Any, shape_idx int) (_ Transform2D.OriginXY) {
	return
}
func (self implementation) AreaRemoveShape(area RID.Any, shape_idx int)     { return }
func (self implementation) AreaClearShapes(area RID.Any)                    { return }
func (self implementation) AreaAttachObjectInstanceId(area RID.Any, id int) { return }
func (self implementation) AreaGetObjectInstanceId(area RID.Any) (_ int)    { return }
func (self implementation) AreaAttachCanvasInstanceId(area RID.Any, id int) { return }
func (self implementation) AreaGetCanvasInstanceId(area RID.Any) (_ int)    { return }
func (self implementation) AreaSetParam(area RID.Any, param gdclass.PhysicsServer2DAreaParameter, value any) {
	return
}
func (self implementation) AreaSetTransform(area RID.Any, transform Transform2D.OriginXY) { return }
func (self implementation) AreaGetParam(area RID.Any, param gdclass.PhysicsServer2DAreaParameter) (_ any) {
	return
}
func (self implementation) AreaGetTransform(area RID.Any) (_ Transform2D.OriginXY)          { return }
func (self implementation) AreaSetCollisionLayer(area RID.Any, layer int)                   { return }
func (self implementation) AreaGetCollisionLayer(area RID.Any) (_ int)                      { return }
func (self implementation) AreaSetCollisionMask(area RID.Any, mask int)                     { return }
func (self implementation) AreaGetCollisionMask(area RID.Any) (_ int)                       { return }
func (self implementation) AreaSetMonitorable(area RID.Any, monitorable bool)               { return }
func (self implementation) AreaSetPickable(area RID.Any, pickable bool)                     { return }
func (self implementation) AreaSetMonitorCallback(area RID.Any, callback Callable.Function) { return }
func (self implementation) AreaSetAreaMonitorCallback(area RID.Any, callback Callable.Function) {
	return
}
func (self implementation) BodyCreate() (_ RID.Any)                                        { return }
func (self implementation) BodySetSpace(body RID.Any, space RID.Any)                       { return }
func (self implementation) BodyGetSpace(body RID.Any) (_ RID.Any)                          { return }
func (self implementation) BodySetMode(body RID.Any, mode gdclass.PhysicsServer2DBodyMode) { return }
func (self implementation) BodyGetMode(body RID.Any) (_ gdclass.PhysicsServer2DBodyMode)   { return }
func (self implementation) BodyAddShape(body RID.Any, shape RID.Any, transform Transform2D.OriginXY, disabled bool) {
	return
}
func (self implementation) BodySetShape(body RID.Any, shape_idx int, shape RID.Any) { return }
func (self implementation) BodySetShapeTransform(body RID.Any, shape_idx int, transform Transform2D.OriginXY) {
	return
}
func (self implementation) BodyGetShapeCount(body RID.Any) (_ int)               { return }
func (self implementation) BodyGetShape(body RID.Any, shape_idx int) (_ RID.Any) { return }
func (self implementation) BodyGetShapeTransform(body RID.Any, shape_idx int) (_ Transform2D.OriginXY) {
	return
}
func (self implementation) BodySetShapeDisabled(body RID.Any, shape_idx int, disabled bool) { return }
func (self implementation) BodySetShapeAsOneWayCollision(body RID.Any, shape_idx int, enable bool, margin Float.X) {
	return
}
func (self implementation) BodyRemoveShape(body RID.Any, shape_idx int)     { return }
func (self implementation) BodyClearShapes(body RID.Any)                    { return }
func (self implementation) BodyAttachObjectInstanceId(body RID.Any, id int) { return }
func (self implementation) BodyGetObjectInstanceId(body RID.Any) (_ int)    { return }
func (self implementation) BodyAttachCanvasInstanceId(body RID.Any, id int) { return }
func (self implementation) BodyGetCanvasInstanceId(body RID.Any) (_ int)    { return }
func (self implementation) BodySetContinuousCollisionDetectionMode(body RID.Any, mode gdclass.PhysicsServer2DCCDMode) {
	return
}
func (self implementation) BodyGetContinuousCollisionDetectionMode(body RID.Any) (_ gdclass.PhysicsServer2DCCDMode) {
	return
}
func (self implementation) BodySetCollisionLayer(body RID.Any, layer int)           { return }
func (self implementation) BodyGetCollisionLayer(body RID.Any) (_ int)              { return }
func (self implementation) BodySetCollisionMask(body RID.Any, mask int)             { return }
func (self implementation) BodyGetCollisionMask(body RID.Any) (_ int)               { return }
func (self implementation) BodySetCollisionPriority(body RID.Any, priority Float.X) { return }
func (self implementation) BodyGetCollisionPriority(body RID.Any) (_ Float.X)       { return }
func (self implementation) BodySetParam(body RID.Any, param gdclass.PhysicsServer2DBodyParameter, value any) {
	return
}
func (self implementation) BodyGetParam(body RID.Any, param gdclass.PhysicsServer2DBodyParameter) (_ any) {
	return
}
func (self implementation) BodyResetMassProperties(body RID.Any) { return }
func (self implementation) BodySetState(body RID.Any, state gdclass.PhysicsServer2DBodyState, value any) {
	return
}
func (self implementation) BodyGetState(body RID.Any, state gdclass.PhysicsServer2DBodyState) (_ any) {
	return
}
func (self implementation) BodyApplyCentralImpulse(body RID.Any, impulse Vector2.XY) { return }
func (self implementation) BodyApplyTorqueImpulse(body RID.Any, impulse Float.X)     { return }
func (self implementation) BodyApplyImpulse(body RID.Any, impulse Vector2.XY, position Vector2.XY) {
	return
}
func (self implementation) BodyApplyCentralForce(body RID.Any, force Vector2.XY) { return }
func (self implementation) BodyApplyForce(body RID.Any, force Vector2.XY, position Vector2.XY) {
	return
}
func (self implementation) BodyApplyTorque(body RID.Any, torque Float.X)               { return }
func (self implementation) BodyAddConstantCentralForce(body RID.Any, force Vector2.XY) { return }
func (self implementation) BodyAddConstantForce(body RID.Any, force Vector2.XY, position Vector2.XY) {
	return
}
func (self implementation) BodyAddConstantTorque(body RID.Any, torque Float.X)               { return }
func (self implementation) BodySetConstantForce(body RID.Any, force Vector2.XY)              { return }
func (self implementation) BodyGetConstantForce(body RID.Any) (_ Vector2.XY)                 { return }
func (self implementation) BodySetConstantTorque(body RID.Any, torque Float.X)               { return }
func (self implementation) BodyGetConstantTorque(body RID.Any) (_ Float.X)                   { return }
func (self implementation) BodySetAxisVelocity(body RID.Any, axis_velocity Vector2.XY)       { return }
func (self implementation) BodyAddCollisionException(body RID.Any, excepted_body RID.Any)    { return }
func (self implementation) BodyRemoveCollisionException(body RID.Any, excepted_body RID.Any) { return }
func (self implementation) BodyGetCollisionExceptions(body RID.Any) (_ []RID.Any)            { return }
func (self implementation) BodySetMaxContactsReported(body RID.Any, amount int)              { return }
func (self implementation) BodyGetMaxContactsReported(body RID.Any) (_ int)                  { return }
func (self implementation) BodySetContactsReportedDepthThreshold(body RID.Any, threshold Float.X) {
	return
}
func (self implementation) BodyGetContactsReportedDepthThreshold(body RID.Any) (_ Float.X)    { return }
func (self implementation) BodySetOmitForceIntegration(body RID.Any, enable bool)             { return }
func (self implementation) BodyIsOmittingForceIntegration(body RID.Any) (_ bool)              { return }
func (self implementation) BodySetStateSyncCallback(body RID.Any, callable Callable.Function) { return }
func (self implementation) BodySetForceIntegrationCallback(body RID.Any, callable Callable.Function, userdata any) {
	return
}
func (self implementation) BodyCollideShape(body RID.Any, body_shape int, shape RID.Any, shape_xform Transform2D.OriginXY, motion Vector2.XY, results unsafe.Pointer, result_max int, result_count *int32) (_ bool) {
	return
}
func (self implementation) BodySetPickable(body RID.Any, pickable bool) { return }
func (self implementation) BodyGetDirectState(body RID.Any) (_ [1]gdclass.PhysicsDirectBodyState2D) {
	return
}
func (self implementation) BodyTestMotion(body RID.Any, from Transform2D.OriginXY, motion Vector2.XY, margin Float.X, collide_separation_ray bool, recovery_as_collision bool, result *MotionResult) (_ bool) {
	return
}
func (self implementation) JointCreate() (_ RID.Any) { return }
func (self implementation) JointClear(joint RID.Any) { return }
func (self implementation) JointSetParam(joint RID.Any, param gdclass.PhysicsServer2DJointParam, value Float.X) {
	return
}
func (self implementation) JointGetParam(joint RID.Any, param gdclass.PhysicsServer2DJointParam) (_ Float.X) {
	return
}
func (self implementation) JointDisableCollisionsBetweenBodies(joint RID.Any, disable bool) { return }
func (self implementation) JointIsDisabledCollisionsBetweenBodies(joint RID.Any) (_ bool)   { return }
func (self implementation) JointMakePin(joint RID.Any, anchor Vector2.XY, body_a RID.Any, body_b RID.Any) {
	return
}
func (self implementation) JointMakeGroove(joint RID.Any, a_groove1 Vector2.XY, a_groove2 Vector2.XY, b_anchor Vector2.XY, body_a RID.Any, body_b RID.Any) {
	return
}
func (self implementation) JointMakeDampedSpring(joint RID.Any, anchor_a Vector2.XY, anchor_b Vector2.XY, body_a RID.Any, body_b RID.Any) {
	return
}
func (self implementation) PinJointSetFlag(joint RID.Any, flag gdclass.PhysicsServer2DPinJointFlag, enabled bool) {
	return
}
func (self implementation) PinJointGetFlag(joint RID.Any, flag gdclass.PhysicsServer2DPinJointFlag) (_ bool) {
	return
}
func (self implementation) PinJointSetParam(joint RID.Any, param gdclass.PhysicsServer2DPinJointParam, value Float.X) {
	return
}
func (self implementation) PinJointGetParam(joint RID.Any, param gdclass.PhysicsServer2DPinJointParam) (_ Float.X) {
	return
}
func (self implementation) DampedSpringJointSetParam(joint RID.Any, param gdclass.PhysicsServer2DDampedSpringParam, value Float.X) {
	return
}
func (self implementation) DampedSpringJointGetParam(joint RID.Any, param gdclass.PhysicsServer2DDampedSpringParam) (_ Float.X) {
	return
}
func (self implementation) JointGetType(joint RID.Any) (_ gdclass.PhysicsServer2DJointType) { return }
func (self implementation) FreeRid(rid RID.Any)                                             { return }
func (self implementation) SetActive(active bool)                                           { return }
func (self implementation) Init()                                                           { return }
func (self implementation) Step(step Float.X)                                               { return }
func (self implementation) Sync()                                                           { return }
func (self implementation) FlushQueries()                                                   { return }
func (self implementation) EndSync()                                                        { return }
func (self implementation) Finish()                                                         { return }
func (self implementation) IsFlushingQueries() (_ bool)                                     { return }
func (self implementation) GetProcessInfo(process_info gdclass.PhysicsServer2DProcessInfo) (_ int) {
	return
}

/*
Overridable version of [method PhysicsServer2D.world_boundary_shape_create].
*/
func (Instance) _world_boundary_shape_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.RID(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.separation_ray_shape_create].
*/
func (Instance) _separation_ray_shape_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.RID(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.segment_shape_create].
*/
func (Instance) _segment_shape_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.RID(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.circle_shape_create].
*/
func (Instance) _circle_shape_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.RID(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.rectangle_shape_create].
*/
func (Instance) _rectangle_shape_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.RID(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.capsule_shape_create].
*/
func (Instance) _capsule_shape_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.RID(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.convex_polygon_shape_create].
*/
func (Instance) _convex_polygon_shape_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.RID(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.concave_polygon_shape_create].
*/
func (Instance) _concave_polygon_shape_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.RID(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.shape_set_data].
*/
func (Instance) _shape_set_data(impl func(ptr unsafe.Pointer, shape RID.Any, data any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)

		var data = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 1))
		defer pointers.End(data)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shape, data.Interface())
	}
}

/*
Should set the custom solver bias for the given [param shape]. It defines how much bodies are forced to separate on contact.
Overridable version of [PhysicsServer2D]'s internal [code]shape_get_custom_solver_bias[/code] method. Corresponds to [member Shape2D.custom_solver_bias].
*/
func (Instance) _shape_set_custom_solver_bias(impl func(ptr unsafe.Pointer, shape RID.Any, bias Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)

		var bias = gd.UnsafeGet[gd.Float](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shape, Float.X(bias))
	}
}

/*
Overridable version of [method PhysicsServer2D.shape_get_type].
*/
func (Instance) _shape_get_type(impl func(ptr unsafe.Pointer, shape RID.Any) gdclass.PhysicsServer2DShapeType) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.shape_get_data].
*/
func (Instance) _shape_get_data(impl func(ptr unsafe.Pointer, shape RID.Any) any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		ptr, ok := pointers.End(gd.NewVariant(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Should return the custom solver bias of the given [param shape], which defines how much bodies are forced to separate on contact when this shape is involved.
Overridable version of [PhysicsServer2D]'s internal [code]shape_get_custom_solver_bias[/code] method. Corresponds to [member Shape2D.custom_solver_bias].
*/
func (Instance) _shape_get_custom_solver_bias(impl func(ptr unsafe.Pointer, shape RID.Any) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Given two shapes and their parameters, should return [code]true[/code] if a collision between the two would occur, with additional details passed in [param results].
Overridable version of [PhysicsServer2D]'s internal [code]shape_collide[/code] method. Corresponds to [method PhysicsDirectSpaceState2D.collide_shape].
*/
func (Instance) _shape_collide(impl func(ptr unsafe.Pointer, shape_A RID.Any, xform_A Transform2D.OriginXY, motion_A Vector2.XY, shape_B RID.Any, xform_B Transform2D.OriginXY, motion_B Vector2.XY, results unsafe.Pointer, result_max int, result_count *int32) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape_A = gd.UnsafeGet[gd.RID](p_args, 0)

		var xform_A = gd.UnsafeGet[gd.Transform2D](p_args, 1)

		var motion_A = gd.UnsafeGet[gd.Vector2](p_args, 2)

		var shape_B = gd.UnsafeGet[gd.RID](p_args, 3)

		var xform_B = gd.UnsafeGet[gd.Transform2D](p_args, 4)

		var motion_B = gd.UnsafeGet[gd.Vector2](p_args, 5)

		var results = gd.UnsafeGet[unsafe.Pointer](p_args, 6)

		var result_max = gd.UnsafeGet[gd.Int](p_args, 7)

		var result_count = gd.UnsafeGet[*int32](p_args, 8)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape_A, xform_A, motion_A, shape_B, xform_B, motion_B, results, int(result_max), result_count)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.space_create].
*/
func (Instance) _space_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.RID(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.space_set_active].
*/
func (Instance) _space_set_active(impl func(ptr unsafe.Pointer, space RID.Any, active bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)

		var active = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, space, active)
	}
}

/*
Overridable version of [method PhysicsServer2D.space_is_active].
*/
func (Instance) _space_is_active(impl func(ptr unsafe.Pointer, space RID.Any) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.space_set_param].
*/
func (Instance) _space_set_param(impl func(ptr unsafe.Pointer, space RID.Any, param gdclass.PhysicsServer2DSpaceParameter, value Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer2DSpaceParameter](p_args, 1)

		var value = gd.UnsafeGet[gd.Float](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, space, param, Float.X(value))
	}
}

/*
Overridable version of [method PhysicsServer2D.space_get_param].
*/
func (Instance) _space_get_param(impl func(ptr unsafe.Pointer, space RID.Any, param gdclass.PhysicsServer2DSpaceParameter) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer2DSpaceParameter](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space, param)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.space_get_direct_state].
*/
func (Instance) _space_get_direct_state(impl func(ptr unsafe.Pointer, space RID.Any) [1]gdclass.PhysicsDirectSpaceState2D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Used internally to allow the given [param space] to store contact points, up to [param max_contacts]. This is automatically set for the main [World2D]'s space when [member SceneTree.debug_collisions_hint] is [code]true[/code], or by checking "Visible Collision Shapes" in the editor. Only works in debug builds.
Overridable version of [PhysicsServer2D]'s internal [code]space_set_debug_contacts[/code] method.
*/
func (Instance) _space_set_debug_contacts(impl func(ptr unsafe.Pointer, space RID.Any, max_contacts int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)

		var max_contacts = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, space, int(max_contacts))
	}
}

/*
Should return the positions of all contacts that have occurred during the last physics step in the given [param space]. See also [method _space_get_contact_count] and [method _space_set_debug_contacts].
Overridable version of [PhysicsServer2D]'s internal [code]space_get_contacts[/code] method.
*/
func (Instance) _space_get_contacts(impl func(ptr unsafe.Pointer, space RID.Any) []Vector2.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		ptr, ok := pointers.End(gd.NewPackedVector2Slice(*(*[]gd.Vector2)(unsafe.Pointer(&ret))))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Should return how many contacts have occurred during the last physics step in the given [param space]. See also [method _space_get_contacts] and [method _space_set_debug_contacts].
Overridable version of [PhysicsServer2D]'s internal [code]space_get_contact_count[/code] method.
*/
func (Instance) _space_get_contact_count(impl func(ptr unsafe.Pointer, space RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.area_create].
*/
func (Instance) _area_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.RID(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_space].
*/
func (Instance) _area_set_space(impl func(ptr unsafe.Pointer, area RID.Any, space RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var space = gd.UnsafeGet[gd.RID](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, space)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_space].
*/
func (Instance) _area_get_space(impl func(ptr unsafe.Pointer, area RID.Any) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, gd.RID(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.area_add_shape].
*/
func (Instance) _area_add_shape(impl func(ptr unsafe.Pointer, area RID.Any, shape RID.Any, transform Transform2D.OriginXY, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape = gd.UnsafeGet[gd.RID](p_args, 1)

		var transform = gd.UnsafeGet[gd.Transform2D](p_args, 2)

		var disabled = gd.UnsafeGet[bool](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, shape, transform, disabled)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_shape].
*/
func (Instance) _area_set_shape(impl func(ptr unsafe.Pointer, area RID.Any, shape_idx int, shape RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)

		var shape = gd.UnsafeGet[gd.RID](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, int(shape_idx), shape)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_shape_transform].
*/
func (Instance) _area_set_shape_transform(impl func(ptr unsafe.Pointer, area RID.Any, shape_idx int, transform Transform2D.OriginXY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)

		var transform = gd.UnsafeGet[gd.Transform2D](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, int(shape_idx), transform)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_shape_disabled].
*/
func (Instance) _area_set_shape_disabled(impl func(ptr unsafe.Pointer, area RID.Any, shape_idx int, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)

		var disabled = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, int(shape_idx), disabled)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_shape_count].
*/
func (Instance) _area_get_shape_count(impl func(ptr unsafe.Pointer, area RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_shape].
*/
func (Instance) _area_get_shape(impl func(ptr unsafe.Pointer, area RID.Any, shape_idx int) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, int(shape_idx))
		gd.UnsafeSet(p_back, gd.RID(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_shape_transform].
*/
func (Instance) _area_get_shape_transform(impl func(ptr unsafe.Pointer, area RID.Any, shape_idx int) Transform2D.OriginXY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, int(shape_idx))
		gd.UnsafeSet(p_back, gd.Transform2D(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.area_remove_shape].
*/
func (Instance) _area_remove_shape(impl func(ptr unsafe.Pointer, area RID.Any, shape_idx int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, int(shape_idx))
	}
}

/*
Overridable version of [method PhysicsServer2D.area_clear_shapes].
*/
func (Instance) _area_clear_shapes(impl func(ptr unsafe.Pointer, area RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_attach_object_instance_id].
*/
func (Instance) _area_attach_object_instance_id(impl func(ptr unsafe.Pointer, area RID.Any, id int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var id = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, int(id))
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_object_instance_id].
*/
func (Instance) _area_get_object_instance_id(impl func(ptr unsafe.Pointer, area RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.area_attach_canvas_instance_id].
*/
func (Instance) _area_attach_canvas_instance_id(impl func(ptr unsafe.Pointer, area RID.Any, id int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var id = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, int(id))
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_canvas_instance_id].
*/
func (Instance) _area_get_canvas_instance_id(impl func(ptr unsafe.Pointer, area RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_param].
*/
func (Instance) _area_set_param(impl func(ptr unsafe.Pointer, area RID.Any, param gdclass.PhysicsServer2DAreaParameter, value any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer2DAreaParameter](p_args, 1)

		var value = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 2))
		defer pointers.End(value)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, param, value.Interface())
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_transform].
*/
func (Instance) _area_set_transform(impl func(ptr unsafe.Pointer, area RID.Any, transform Transform2D.OriginXY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var transform = gd.UnsafeGet[gd.Transform2D](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, transform)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_param].
*/
func (Instance) _area_get_param(impl func(ptr unsafe.Pointer, area RID.Any, param gdclass.PhysicsServer2DAreaParameter) any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer2DAreaParameter](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, param)
		ptr, ok := pointers.End(gd.NewVariant(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_transform].
*/
func (Instance) _area_get_transform(impl func(ptr unsafe.Pointer, area RID.Any) Transform2D.OriginXY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, gd.Transform2D(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_collision_layer].
*/
func (Instance) _area_set_collision_layer(impl func(ptr unsafe.Pointer, area RID.Any, layer int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var layer = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, int(layer))
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_collision_layer].
*/
func (Instance) _area_get_collision_layer(impl func(ptr unsafe.Pointer, area RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_collision_mask].
*/
func (Instance) _area_set_collision_mask(impl func(ptr unsafe.Pointer, area RID.Any, mask int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var mask = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, int(mask))
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_collision_mask].
*/
func (Instance) _area_get_collision_mask(impl func(ptr unsafe.Pointer, area RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_monitorable].
*/
func (Instance) _area_set_monitorable(impl func(ptr unsafe.Pointer, area RID.Any, monitorable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var monitorable = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, monitorable)
	}
}

/*
If set to [code]true[/code], allows the area with the given [RID] to detect mouse inputs when the mouse cursor is hovering on it.
Overridable version of [PhysicsServer2D]'s internal [code]area_set_pickable[/code] method. Corresponds to [member CollisionObject2D.input_pickable].
*/
func (Instance) _area_set_pickable(impl func(ptr unsafe.Pointer, area RID.Any, pickable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var pickable = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, pickable)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_monitor_callback].
*/
func (Instance) _area_set_monitor_callback(impl func(ptr unsafe.Pointer, area RID.Any, callback Callable.Function)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var callback = Callable.Through(gd.CallableProxy{}, pointers.Pack(pointers.New[gd.Callable](gd.UnsafeGet[[2]uint64](p_args, 1))))
		defer pointers.End(gd.InternalCallable(callback))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, callback)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_area_monitor_callback].
*/
func (Instance) _area_set_area_monitor_callback(impl func(ptr unsafe.Pointer, area RID.Any, callback Callable.Function)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var callback = Callable.Through(gd.CallableProxy{}, pointers.Pack(pointers.New[gd.Callable](gd.UnsafeGet[[2]uint64](p_args, 1))))
		defer pointers.End(gd.InternalCallable(callback))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, callback)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_create].
*/
func (Instance) _body_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.RID(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_space].
*/
func (Instance) _body_set_space(impl func(ptr unsafe.Pointer, body RID.Any, space RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var space = gd.UnsafeGet[gd.RID](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, space)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_space].
*/
func (Instance) _body_get_space(impl func(ptr unsafe.Pointer, body RID.Any) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.RID(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_mode].
*/
func (Instance) _body_set_mode(impl func(ptr unsafe.Pointer, body RID.Any, mode gdclass.PhysicsServer2DBodyMode)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var mode = gd.UnsafeGet[gdclass.PhysicsServer2DBodyMode](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, mode)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_mode].
*/
func (Instance) _body_get_mode(impl func(ptr unsafe.Pointer, body RID.Any) gdclass.PhysicsServer2DBodyMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_add_shape].
*/
func (Instance) _body_add_shape(impl func(ptr unsafe.Pointer, body RID.Any, shape RID.Any, transform Transform2D.OriginXY, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape = gd.UnsafeGet[gd.RID](p_args, 1)

		var transform = gd.UnsafeGet[gd.Transform2D](p_args, 2)

		var disabled = gd.UnsafeGet[bool](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, shape, transform, disabled)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_shape].
*/
func (Instance) _body_set_shape(impl func(ptr unsafe.Pointer, body RID.Any, shape_idx int, shape RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)

		var shape = gd.UnsafeGet[gd.RID](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(shape_idx), shape)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_shape_transform].
*/
func (Instance) _body_set_shape_transform(impl func(ptr unsafe.Pointer, body RID.Any, shape_idx int, transform Transform2D.OriginXY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)

		var transform = gd.UnsafeGet[gd.Transform2D](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(shape_idx), transform)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_shape_count].
*/
func (Instance) _body_get_shape_count(impl func(ptr unsafe.Pointer, body RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_shape].
*/
func (Instance) _body_get_shape(impl func(ptr unsafe.Pointer, body RID.Any, shape_idx int) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, int(shape_idx))
		gd.UnsafeSet(p_back, gd.RID(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_shape_transform].
*/
func (Instance) _body_get_shape_transform(impl func(ptr unsafe.Pointer, body RID.Any, shape_idx int) Transform2D.OriginXY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, int(shape_idx))
		gd.UnsafeSet(p_back, gd.Transform2D(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_shape_disabled].
*/
func (Instance) _body_set_shape_disabled(impl func(ptr unsafe.Pointer, body RID.Any, shape_idx int, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)

		var disabled = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(shape_idx), disabled)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_shape_as_one_way_collision].
*/
func (Instance) _body_set_shape_as_one_way_collision(impl func(ptr unsafe.Pointer, body RID.Any, shape_idx int, enable bool, margin Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)

		var enable = gd.UnsafeGet[bool](p_args, 2)

		var margin = gd.UnsafeGet[gd.Float](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(shape_idx), enable, Float.X(margin))
	}
}

/*
Overridable version of [method PhysicsServer2D.body_remove_shape].
*/
func (Instance) _body_remove_shape(impl func(ptr unsafe.Pointer, body RID.Any, shape_idx int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(shape_idx))
	}
}

/*
Overridable version of [method PhysicsServer2D.body_clear_shapes].
*/
func (Instance) _body_clear_shapes(impl func(ptr unsafe.Pointer, body RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_attach_object_instance_id].
*/
func (Instance) _body_attach_object_instance_id(impl func(ptr unsafe.Pointer, body RID.Any, id int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var id = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(id))
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_object_instance_id].
*/
func (Instance) _body_get_object_instance_id(impl func(ptr unsafe.Pointer, body RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.body_attach_canvas_instance_id].
*/
func (Instance) _body_attach_canvas_instance_id(impl func(ptr unsafe.Pointer, body RID.Any, id int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var id = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(id))
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_canvas_instance_id].
*/
func (Instance) _body_get_canvas_instance_id(impl func(ptr unsafe.Pointer, body RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_continuous_collision_detection_mode].
*/
func (Instance) _body_set_continuous_collision_detection_mode(impl func(ptr unsafe.Pointer, body RID.Any, mode gdclass.PhysicsServer2DCCDMode)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var mode = gd.UnsafeGet[gdclass.PhysicsServer2DCCDMode](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, mode)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_continuous_collision_detection_mode].
*/
func (Instance) _body_get_continuous_collision_detection_mode(impl func(ptr unsafe.Pointer, body RID.Any) gdclass.PhysicsServer2DCCDMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_collision_layer].
*/
func (Instance) _body_set_collision_layer(impl func(ptr unsafe.Pointer, body RID.Any, layer int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var layer = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(layer))
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_collision_layer].
*/
func (Instance) _body_get_collision_layer(impl func(ptr unsafe.Pointer, body RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_collision_mask].
*/
func (Instance) _body_set_collision_mask(impl func(ptr unsafe.Pointer, body RID.Any, mask int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var mask = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(mask))
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_collision_mask].
*/
func (Instance) _body_get_collision_mask(impl func(ptr unsafe.Pointer, body RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_collision_priority].
*/
func (Instance) _body_set_collision_priority(impl func(ptr unsafe.Pointer, body RID.Any, priority Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var priority = gd.UnsafeGet[gd.Float](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, Float.X(priority))
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_collision_priority].
*/
func (Instance) _body_get_collision_priority(impl func(ptr unsafe.Pointer, body RID.Any) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_param].
*/
func (Instance) _body_set_param(impl func(ptr unsafe.Pointer, body RID.Any, param gdclass.PhysicsServer2DBodyParameter, value any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer2DBodyParameter](p_args, 1)

		var value = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 2))
		defer pointers.End(value)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, param, value.Interface())
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_param].
*/
func (Instance) _body_get_param(impl func(ptr unsafe.Pointer, body RID.Any, param gdclass.PhysicsServer2DBodyParameter) any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer2DBodyParameter](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, param)
		ptr, ok := pointers.End(gd.NewVariant(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_reset_mass_properties].
*/
func (Instance) _body_reset_mass_properties(impl func(ptr unsafe.Pointer, body RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_state].
*/
func (Instance) _body_set_state(impl func(ptr unsafe.Pointer, body RID.Any, state gdclass.PhysicsServer2DBodyState, value any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var state = gd.UnsafeGet[gdclass.PhysicsServer2DBodyState](p_args, 1)

		var value = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 2))
		defer pointers.End(value)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, state, value.Interface())
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_state].
*/
func (Instance) _body_get_state(impl func(ptr unsafe.Pointer, body RID.Any, state gdclass.PhysicsServer2DBodyState) any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var state = gd.UnsafeGet[gdclass.PhysicsServer2DBodyState](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, state)
		ptr, ok := pointers.End(gd.NewVariant(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_apply_central_impulse].
*/
func (Instance) _body_apply_central_impulse(impl func(ptr unsafe.Pointer, body RID.Any, impulse Vector2.XY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var impulse = gd.UnsafeGet[gd.Vector2](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, impulse)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_apply_torque_impulse].
*/
func (Instance) _body_apply_torque_impulse(impl func(ptr unsafe.Pointer, body RID.Any, impulse Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var impulse = gd.UnsafeGet[gd.Float](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, Float.X(impulse))
	}
}

/*
Overridable version of [method PhysicsServer2D.body_apply_impulse].
*/
func (Instance) _body_apply_impulse(impl func(ptr unsafe.Pointer, body RID.Any, impulse Vector2.XY, position Vector2.XY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var impulse = gd.UnsafeGet[gd.Vector2](p_args, 1)

		var position = gd.UnsafeGet[gd.Vector2](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, impulse, position)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_apply_central_force].
*/
func (Instance) _body_apply_central_force(impl func(ptr unsafe.Pointer, body RID.Any, force Vector2.XY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var force = gd.UnsafeGet[gd.Vector2](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_apply_force].
*/
func (Instance) _body_apply_force(impl func(ptr unsafe.Pointer, body RID.Any, force Vector2.XY, position Vector2.XY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var force = gd.UnsafeGet[gd.Vector2](p_args, 1)

		var position = gd.UnsafeGet[gd.Vector2](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force, position)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_apply_torque].
*/
func (Instance) _body_apply_torque(impl func(ptr unsafe.Pointer, body RID.Any, torque Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var torque = gd.UnsafeGet[gd.Float](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, Float.X(torque))
	}
}

/*
Overridable version of [method PhysicsServer2D.body_add_constant_central_force].
*/
func (Instance) _body_add_constant_central_force(impl func(ptr unsafe.Pointer, body RID.Any, force Vector2.XY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var force = gd.UnsafeGet[gd.Vector2](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_add_constant_force].
*/
func (Instance) _body_add_constant_force(impl func(ptr unsafe.Pointer, body RID.Any, force Vector2.XY, position Vector2.XY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var force = gd.UnsafeGet[gd.Vector2](p_args, 1)

		var position = gd.UnsafeGet[gd.Vector2](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force, position)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_add_constant_torque].
*/
func (Instance) _body_add_constant_torque(impl func(ptr unsafe.Pointer, body RID.Any, torque Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var torque = gd.UnsafeGet[gd.Float](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, Float.X(torque))
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_constant_force].
*/
func (Instance) _body_set_constant_force(impl func(ptr unsafe.Pointer, body RID.Any, force Vector2.XY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var force = gd.UnsafeGet[gd.Vector2](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_constant_force].
*/
func (Instance) _body_get_constant_force(impl func(ptr unsafe.Pointer, body RID.Any) Vector2.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Vector2(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_constant_torque].
*/
func (Instance) _body_set_constant_torque(impl func(ptr unsafe.Pointer, body RID.Any, torque Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var torque = gd.UnsafeGet[gd.Float](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, Float.X(torque))
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_constant_torque].
*/
func (Instance) _body_get_constant_torque(impl func(ptr unsafe.Pointer, body RID.Any) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_axis_velocity].
*/
func (Instance) _body_set_axis_velocity(impl func(ptr unsafe.Pointer, body RID.Any, axis_velocity Vector2.XY)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var axis_velocity = gd.UnsafeGet[gd.Vector2](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, axis_velocity)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_add_collision_exception].
*/
func (Instance) _body_add_collision_exception(impl func(ptr unsafe.Pointer, body RID.Any, excepted_body RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var excepted_body = gd.UnsafeGet[gd.RID](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, excepted_body)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_remove_collision_exception].
*/
func (Instance) _body_remove_collision_exception(impl func(ptr unsafe.Pointer, body RID.Any, excepted_body RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var excepted_body = gd.UnsafeGet[gd.RID](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, excepted_body)
	}
}

/*
Returns the [RID]s of all bodies added as collision exceptions for the given [param body]. See also [method _body_add_collision_exception] and [method _body_remove_collision_exception].
Overridable version of [PhysicsServer2D]'s internal [code]body_get_collision_exceptions[/code] method. Corresponds to [method PhysicsBody2D.get_collision_exceptions].
*/
func (Instance) _body_get_collision_exceptions(impl func(ptr unsafe.Pointer, body RID.Any) []RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		ptr, ok := pointers.End(gd.InternalArray(gd.ArrayFromSlice[Array.Contains[gd.RID]](ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_max_contacts_reported].
*/
func (Instance) _body_set_max_contacts_reported(impl func(ptr unsafe.Pointer, body RID.Any, amount int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var amount = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(amount))
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_max_contacts_reported].
*/
func (Instance) _body_get_max_contacts_reported(impl func(ptr unsafe.Pointer, body RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Overridable version of [PhysicsServer2D]'s internal [code]body_set_contacts_reported_depth_threshold[/code] method.
[b]Note:[/b] This method is currently unused by Godot's default physics implementation.
*/
func (Instance) _body_set_contacts_reported_depth_threshold(impl func(ptr unsafe.Pointer, body RID.Any, threshold Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var threshold = gd.UnsafeGet[gd.Float](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, Float.X(threshold))
	}
}

/*
Overridable version of [PhysicsServer2D]'s internal [code]body_get_contacts_reported_depth_threshold[/code] method.
[b]Note:[/b] This method is currently unused by Godot's default physics implementation.
*/
func (Instance) _body_get_contacts_reported_depth_threshold(impl func(ptr unsafe.Pointer, body RID.Any) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_omit_force_integration].
*/
func (Instance) _body_set_omit_force_integration(impl func(ptr unsafe.Pointer, body RID.Any, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var enable = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, enable)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_is_omitting_force_integration].
*/
func (Instance) _body_is_omitting_force_integration(impl func(ptr unsafe.Pointer, body RID.Any) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Assigns the [param body] to call the given [param callable] during the synchronization phase of the loop, before [method _step] is called. See also [method _sync].
Overridable version of [method PhysicsServer2D.body_set_state_sync_callback].
*/
func (Instance) _body_set_state_sync_callback(impl func(ptr unsafe.Pointer, body RID.Any, callable Callable.Function)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var callable = Callable.Through(gd.CallableProxy{}, pointers.Pack(pointers.New[gd.Callable](gd.UnsafeGet[[2]uint64](p_args, 1))))
		defer pointers.End(gd.InternalCallable(callable))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, callable)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_force_integration_callback].
*/
func (Instance) _body_set_force_integration_callback(impl func(ptr unsafe.Pointer, body RID.Any, callable Callable.Function, userdata any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var callable = Callable.Through(gd.CallableProxy{}, pointers.Pack(pointers.New[gd.Callable](gd.UnsafeGet[[2]uint64](p_args, 1))))
		defer pointers.End(gd.InternalCallable(callable))
		var userdata = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 2))
		defer pointers.End(userdata)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, callable, userdata.Interface())
	}
}

/*
Given a [param body], a [param shape], and their respective parameters, this method should return [code]true[/code] if a collision between the two would occur, with additional details passed in [param results].
Overridable version of [PhysicsServer2D]'s internal [code]shape_collide[/code] method. Corresponds to [method PhysicsDirectSpaceState2D.collide_shape].
*/
func (Instance) _body_collide_shape(impl func(ptr unsafe.Pointer, body RID.Any, body_shape int, shape RID.Any, shape_xform Transform2D.OriginXY, motion Vector2.XY, results unsafe.Pointer, result_max int, result_count *int32) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var body_shape = gd.UnsafeGet[gd.Int](p_args, 1)

		var shape = gd.UnsafeGet[gd.RID](p_args, 2)

		var shape_xform = gd.UnsafeGet[gd.Transform2D](p_args, 3)

		var motion = gd.UnsafeGet[gd.Vector2](p_args, 4)

		var results = gd.UnsafeGet[unsafe.Pointer](p_args, 5)

		var result_max = gd.UnsafeGet[gd.Int](p_args, 6)

		var result_count = gd.UnsafeGet[*int32](p_args, 7)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, int(body_shape), shape, shape_xform, motion, results, int(result_max), result_count)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
If set to [code]true[/code], allows the body with the given [RID] to detect mouse inputs when the mouse cursor is hovering on it.
Overridable version of [PhysicsServer2D]'s internal [code]body_set_pickable[/code] method. Corresponds to [member CollisionObject2D.input_pickable].
*/
func (Instance) _body_set_pickable(impl func(ptr unsafe.Pointer, body RID.Any, pickable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var pickable = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, pickable)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_direct_state].
*/
func (Instance) _body_get_direct_state(impl func(ptr unsafe.Pointer, body RID.Any) [1]gdclass.PhysicsDirectBodyState2D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_test_motion]. Unlike the exposed implementation, this method does not receive all of the arguments inside a [PhysicsTestMotionParameters2D].
*/
func (Instance) _body_test_motion(impl func(ptr unsafe.Pointer, body RID.Any, from Transform2D.OriginXY, motion Vector2.XY, margin Float.X, collide_separation_ray bool, recovery_as_collision bool, result *MotionResult) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var from = gd.UnsafeGet[gd.Transform2D](p_args, 1)

		var motion = gd.UnsafeGet[gd.Vector2](p_args, 2)

		var margin = gd.UnsafeGet[gd.Float](p_args, 3)

		var collide_separation_ray = gd.UnsafeGet[bool](p_args, 4)

		var recovery_as_collision = gd.UnsafeGet[bool](p_args, 5)

		var result = gd.UnsafeGet[*MotionResult](p_args, 6)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, from, motion, Float.X(margin), collide_separation_ray, recovery_as_collision, result)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.joint_create].
*/
func (Instance) _joint_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.RID(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.joint_clear].
*/
func (Instance) _joint_clear(impl func(ptr unsafe.Pointer, joint RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint)
	}
}

/*
Overridable version of [method PhysicsServer2D.joint_set_param].
*/
func (Instance) _joint_set_param(impl func(ptr unsafe.Pointer, joint RID.Any, param gdclass.PhysicsServer2DJointParam, value Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer2DJointParam](p_args, 1)

		var value = gd.UnsafeGet[gd.Float](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, Float.X(value))
	}
}

/*
Overridable version of [method PhysicsServer2D.joint_get_param].
*/
func (Instance) _joint_get_param(impl func(ptr unsafe.Pointer, joint RID.Any, param gdclass.PhysicsServer2DJointParam) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer2DJointParam](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.joint_disable_collisions_between_bodies].
*/
func (Instance) _joint_disable_collisions_between_bodies(impl func(ptr unsafe.Pointer, joint RID.Any, disable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		var disable = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, disable)
	}
}

/*
Overridable version of [method PhysicsServer2D.joint_is_disabled_collisions_between_bodies].
*/
func (Instance) _joint_is_disabled_collisions_between_bodies(impl func(ptr unsafe.Pointer, joint RID.Any) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.joint_make_pin].
*/
func (Instance) _joint_make_pin(impl func(ptr unsafe.Pointer, joint RID.Any, anchor Vector2.XY, body_a RID.Any, body_b RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		var anchor = gd.UnsafeGet[gd.Vector2](p_args, 1)

		var body_a = gd.UnsafeGet[gd.RID](p_args, 2)

		var body_b = gd.UnsafeGet[gd.RID](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, anchor, body_a, body_b)
	}
}

/*
Overridable version of [method PhysicsServer2D.joint_make_groove].
*/
func (Instance) _joint_make_groove(impl func(ptr unsafe.Pointer, joint RID.Any, a_groove1 Vector2.XY, a_groove2 Vector2.XY, b_anchor Vector2.XY, body_a RID.Any, body_b RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		var a_groove1 = gd.UnsafeGet[gd.Vector2](p_args, 1)

		var a_groove2 = gd.UnsafeGet[gd.Vector2](p_args, 2)

		var b_anchor = gd.UnsafeGet[gd.Vector2](p_args, 3)

		var body_a = gd.UnsafeGet[gd.RID](p_args, 4)

		var body_b = gd.UnsafeGet[gd.RID](p_args, 5)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, a_groove1, a_groove2, b_anchor, body_a, body_b)
	}
}

/*
Overridable version of [method PhysicsServer2D.joint_make_damped_spring].
*/
func (Instance) _joint_make_damped_spring(impl func(ptr unsafe.Pointer, joint RID.Any, anchor_a Vector2.XY, anchor_b Vector2.XY, body_a RID.Any, body_b RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		var anchor_a = gd.UnsafeGet[gd.Vector2](p_args, 1)

		var anchor_b = gd.UnsafeGet[gd.Vector2](p_args, 2)

		var body_a = gd.UnsafeGet[gd.RID](p_args, 3)

		var body_b = gd.UnsafeGet[gd.RID](p_args, 4)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, anchor_a, anchor_b, body_a, body_b)
	}
}

/*
Overridable version of [method PhysicsServer2D.pin_joint_set_flag].
*/
func (Instance) _pin_joint_set_flag(impl func(ptr unsafe.Pointer, joint RID.Any, flag gdclass.PhysicsServer2DPinJointFlag, enabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		var flag = gd.UnsafeGet[gdclass.PhysicsServer2DPinJointFlag](p_args, 1)

		var enabled = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, flag, enabled)
	}
}

/*
Overridable version of [method PhysicsServer2D.pin_joint_get_flag].
*/
func (Instance) _pin_joint_get_flag(impl func(ptr unsafe.Pointer, joint RID.Any, flag gdclass.PhysicsServer2DPinJointFlag) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		var flag = gd.UnsafeGet[gdclass.PhysicsServer2DPinJointFlag](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, flag)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.pin_joint_set_param].
*/
func (Instance) _pin_joint_set_param(impl func(ptr unsafe.Pointer, joint RID.Any, param gdclass.PhysicsServer2DPinJointParam, value Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer2DPinJointParam](p_args, 1)

		var value = gd.UnsafeGet[gd.Float](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, Float.X(value))
	}
}

/*
Overridable version of [method PhysicsServer2D.pin_joint_get_param].
*/
func (Instance) _pin_joint_get_param(impl func(ptr unsafe.Pointer, joint RID.Any, param gdclass.PhysicsServer2DPinJointParam) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer2DPinJointParam](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.damped_spring_joint_set_param].
*/
func (Instance) _damped_spring_joint_set_param(impl func(ptr unsafe.Pointer, joint RID.Any, param gdclass.PhysicsServer2DDampedSpringParam, value Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer2DDampedSpringParam](p_args, 1)

		var value = gd.UnsafeGet[gd.Float](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, Float.X(value))
	}
}

/*
Overridable version of [method PhysicsServer2D.damped_spring_joint_get_param].
*/
func (Instance) _damped_spring_joint_get_param(impl func(ptr unsafe.Pointer, joint RID.Any, param gdclass.PhysicsServer2DDampedSpringParam) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer2DDampedSpringParam](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Overridable version of [method PhysicsServer2D.joint_get_type].
*/
func (Instance) _joint_get_type(impl func(ptr unsafe.Pointer, joint RID.Any) gdclass.PhysicsServer2DJointType) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.free_rid].
*/
func (Instance) _free_rid(impl func(ptr unsafe.Pointer, rid RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, rid)
	}
}

/*
Overridable version of [method PhysicsServer2D.set_active].
*/
func (Instance) _set_active(impl func(ptr unsafe.Pointer, active bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var active = gd.UnsafeGet[bool](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, active)
	}
}

/*
Called when the main loop is initialized and creates a new instance of this physics server. See also [method MainLoop._initialize] and [method _finish].
Overridable version of [PhysicsServer2D]'s internal [code]init[/code] method.
*/
func (Instance) _init(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called every physics step to process the physics simulation. [param step] is the time elapsed since the last physics step, in seconds. It is usually the same as [method Node.get_physics_process_delta_time].
Overridable version of [PhysicsServer2D]'s internal [code skip-lint]step[/code] method.
*/
func (Instance) _step(impl func(ptr unsafe.Pointer, step Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var step = gd.UnsafeGet[gd.Float](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, Float.X(step))
	}
}

/*
Called to indicate that the physics server is synchronizing and cannot access physics states if running on a separate thread. See also [method _end_sync].
Overridable version of [PhysicsServer2D]'s internal [code]sync[/code] method.
*/
func (Instance) _sync(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called every physics step before [method _step] to process all remaining queries.
Overridable version of [PhysicsServer2D]'s internal [code]flush_queries[/code] method.
*/
func (Instance) _flush_queries(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called to indicate that the physics server has stopped synchronizing. It is in the loop's iteration/physics phase, and can access physics objects even if running on a separate thread. See also [method _sync].
Overridable version of [PhysicsServer2D]'s internal [code]end_sync[/code] method.
*/
func (Instance) _end_sync(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the main loop finalizes to shut down the physics server. See also [method MainLoop._finalize] and [method _init].
Overridable version of [PhysicsServer2D]'s internal [code]finish[/code] method.
*/
func (Instance) _finish(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Overridable method that should return [code]true[/code] when the physics server is processing queries. See also [method _flush_queries].
Overridable version of [PhysicsServer2D]'s internal [code]is_flushing_queries[/code] method.
*/
func (Instance) _is_flushing_queries(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.get_process_info].
*/
func (Instance) _get_process_info(impl func(ptr unsafe.Pointer, process_info gdclass.PhysicsServer2DProcessInfo) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var process_info = gd.UnsafeGet[gdclass.PhysicsServer2DProcessInfo](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, process_info)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Returns [code]true[/code] if the body with the given [RID] is being excluded from [method _body_test_motion]. See also [method Object.get_instance_id].
*/
func (self Instance) BodyTestMotionIsExcludingBody(body RID.Body2D) bool { //gd:PhysicsServer2DExtension.body_test_motion_is_excluding_body
	return bool(class(self).BodyTestMotionIsExcludingBody(gd.RID(body)))
}

/*
Returns [code]true[/code] if the object with the given instance ID is being excluded from [method _body_test_motion]. See also [method Object.get_instance_id].
*/
func (self Instance) BodyTestMotionIsExcludingObject(obj int) bool { //gd:PhysicsServer2DExtension.body_test_motion_is_excluding_object
	return bool(class(self).BodyTestMotionIsExcludingObject(gd.Int(obj)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PhysicsServer2DExtension

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicsServer2DExtension"))
	casted := Instance{*(*gdclass.PhysicsServer2DExtension)(unsafe.Pointer(&object))}
	return casted
}

/*
Overridable version of [method PhysicsServer2D.world_boundary_shape_create].
*/
func (class) _world_boundary_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.separation_ray_shape_create].
*/
func (class) _separation_ray_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.segment_shape_create].
*/
func (class) _segment_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.circle_shape_create].
*/
func (class) _circle_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.rectangle_shape_create].
*/
func (class) _rectangle_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.capsule_shape_create].
*/
func (class) _capsule_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.convex_polygon_shape_create].
*/
func (class) _convex_polygon_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.concave_polygon_shape_create].
*/
func (class) _concave_polygon_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.shape_set_data].
*/
func (class) _shape_set_data(impl func(ptr unsafe.Pointer, shape gd.RID, data gd.Variant)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)

		var data = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 1))
		defer pointers.End(data)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shape, data)
	}
}

/*
Should set the custom solver bias for the given [param shape]. It defines how much bodies are forced to separate on contact.
Overridable version of [PhysicsServer2D]'s internal [code]shape_get_custom_solver_bias[/code] method. Corresponds to [member Shape2D.custom_solver_bias].
*/
func (class) _shape_set_custom_solver_bias(impl func(ptr unsafe.Pointer, shape gd.RID, bias gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)

		var bias = gd.UnsafeGet[gd.Float](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shape, bias)
	}
}

/*
Overridable version of [method PhysicsServer2D.shape_get_type].
*/
func (class) _shape_get_type(impl func(ptr unsafe.Pointer, shape gd.RID) gdclass.PhysicsServer2DShapeType) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.shape_get_data].
*/
func (class) _shape_get_data(impl func(ptr unsafe.Pointer, shape gd.RID) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Should return the custom solver bias of the given [param shape], which defines how much bodies are forced to separate on contact when this shape is involved.
Overridable version of [PhysicsServer2D]'s internal [code]shape_get_custom_solver_bias[/code] method. Corresponds to [member Shape2D.custom_solver_bias].
*/
func (class) _shape_get_custom_solver_bias(impl func(ptr unsafe.Pointer, shape gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Given two shapes and their parameters, should return [code]true[/code] if a collision between the two would occur, with additional details passed in [param results].
Overridable version of [PhysicsServer2D]'s internal [code]shape_collide[/code] method. Corresponds to [method PhysicsDirectSpaceState2D.collide_shape].
*/
func (class) _shape_collide(impl func(ptr unsafe.Pointer, shape_A gd.RID, xform_A gd.Transform2D, motion_A gd.Vector2, shape_B gd.RID, xform_B gd.Transform2D, motion_B gd.Vector2, results unsafe.Pointer, result_max gd.Int, result_count *int32) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape_A = gd.UnsafeGet[gd.RID](p_args, 0)

		var xform_A = gd.UnsafeGet[gd.Transform2D](p_args, 1)

		var motion_A = gd.UnsafeGet[gd.Vector2](p_args, 2)

		var shape_B = gd.UnsafeGet[gd.RID](p_args, 3)

		var xform_B = gd.UnsafeGet[gd.Transform2D](p_args, 4)

		var motion_B = gd.UnsafeGet[gd.Vector2](p_args, 5)

		var results = gd.UnsafeGet[unsafe.Pointer](p_args, 6)

		var result_max = gd.UnsafeGet[gd.Int](p_args, 7)

		var result_count = gd.UnsafeGet[*int32](p_args, 8)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape_A, xform_A, motion_A, shape_B, xform_B, motion_B, results, result_max, result_count)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.space_create].
*/
func (class) _space_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.space_set_active].
*/
func (class) _space_set_active(impl func(ptr unsafe.Pointer, space gd.RID, active bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)

		var active = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, space, active)
	}
}

/*
Overridable version of [method PhysicsServer2D.space_is_active].
*/
func (class) _space_is_active(impl func(ptr unsafe.Pointer, space gd.RID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.space_set_param].
*/
func (class) _space_set_param(impl func(ptr unsafe.Pointer, space gd.RID, param gdclass.PhysicsServer2DSpaceParameter, value gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer2DSpaceParameter](p_args, 1)

		var value = gd.UnsafeGet[gd.Float](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, space, param, value)
	}
}

/*
Overridable version of [method PhysicsServer2D.space_get_param].
*/
func (class) _space_get_param(impl func(ptr unsafe.Pointer, space gd.RID, param gdclass.PhysicsServer2DSpaceParameter) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer2DSpaceParameter](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space, param)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.space_get_direct_state].
*/
func (class) _space_get_direct_state(impl func(ptr unsafe.Pointer, space gd.RID) [1]gdclass.PhysicsDirectSpaceState2D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Used internally to allow the given [param space] to store contact points, up to [param max_contacts]. This is automatically set for the main [World2D]'s space when [member SceneTree.debug_collisions_hint] is [code]true[/code], or by checking "Visible Collision Shapes" in the editor. Only works in debug builds.
Overridable version of [PhysicsServer2D]'s internal [code]space_set_debug_contacts[/code] method.
*/
func (class) _space_set_debug_contacts(impl func(ptr unsafe.Pointer, space gd.RID, max_contacts gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)

		var max_contacts = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, space, max_contacts)
	}
}

/*
Should return the positions of all contacts that have occurred during the last physics step in the given [param space]. See also [method _space_get_contact_count] and [method _space_set_debug_contacts].
Overridable version of [PhysicsServer2D]'s internal [code]space_get_contacts[/code] method.
*/
func (class) _space_get_contacts(impl func(ptr unsafe.Pointer, space gd.RID) gd.PackedVector2Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Should return how many contacts have occurred during the last physics step in the given [param space]. See also [method _space_get_contacts] and [method _space_set_debug_contacts].
Overridable version of [PhysicsServer2D]'s internal [code]space_get_contact_count[/code] method.
*/
func (class) _space_get_contact_count(impl func(ptr unsafe.Pointer, space gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_create].
*/
func (class) _area_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_space].
*/
func (class) _area_set_space(impl func(ptr unsafe.Pointer, area gd.RID, space gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var space = gd.UnsafeGet[gd.RID](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, space)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_space].
*/
func (class) _area_get_space(impl func(ptr unsafe.Pointer, area gd.RID) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_add_shape].
*/
func (class) _area_add_shape(impl func(ptr unsafe.Pointer, area gd.RID, shape gd.RID, transform gd.Transform2D, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape = gd.UnsafeGet[gd.RID](p_args, 1)

		var transform = gd.UnsafeGet[gd.Transform2D](p_args, 2)

		var disabled = gd.UnsafeGet[bool](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, shape, transform, disabled)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_shape].
*/
func (class) _area_set_shape(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx gd.Int, shape gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)

		var shape = gd.UnsafeGet[gd.RID](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, shape_idx, shape)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_shape_transform].
*/
func (class) _area_set_shape_transform(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx gd.Int, transform gd.Transform2D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)

		var transform = gd.UnsafeGet[gd.Transform2D](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, shape_idx, transform)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_shape_disabled].
*/
func (class) _area_set_shape_disabled(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx gd.Int, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)

		var disabled = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, shape_idx, disabled)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_shape_count].
*/
func (class) _area_get_shape_count(impl func(ptr unsafe.Pointer, area gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_shape].
*/
func (class) _area_get_shape(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx gd.Int) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, shape_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_shape_transform].
*/
func (class) _area_get_shape_transform(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx gd.Int) gd.Transform2D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, shape_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_remove_shape].
*/
func (class) _area_remove_shape(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, shape_idx)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_clear_shapes].
*/
func (class) _area_clear_shapes(impl func(ptr unsafe.Pointer, area gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_attach_object_instance_id].
*/
func (class) _area_attach_object_instance_id(impl func(ptr unsafe.Pointer, area gd.RID, id gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var id = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, id)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_object_instance_id].
*/
func (class) _area_get_object_instance_id(impl func(ptr unsafe.Pointer, area gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_attach_canvas_instance_id].
*/
func (class) _area_attach_canvas_instance_id(impl func(ptr unsafe.Pointer, area gd.RID, id gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var id = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, id)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_canvas_instance_id].
*/
func (class) _area_get_canvas_instance_id(impl func(ptr unsafe.Pointer, area gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_param].
*/
func (class) _area_set_param(impl func(ptr unsafe.Pointer, area gd.RID, param gdclass.PhysicsServer2DAreaParameter, value gd.Variant)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer2DAreaParameter](p_args, 1)

		var value = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 2))
		defer pointers.End(value)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, param, value)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_transform].
*/
func (class) _area_set_transform(impl func(ptr unsafe.Pointer, area gd.RID, transform gd.Transform2D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var transform = gd.UnsafeGet[gd.Transform2D](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, transform)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_param].
*/
func (class) _area_get_param(impl func(ptr unsafe.Pointer, area gd.RID, param gdclass.PhysicsServer2DAreaParameter) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer2DAreaParameter](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, param)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_transform].
*/
func (class) _area_get_transform(impl func(ptr unsafe.Pointer, area gd.RID) gd.Transform2D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_collision_layer].
*/
func (class) _area_set_collision_layer(impl func(ptr unsafe.Pointer, area gd.RID, layer gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var layer = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, layer)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_collision_layer].
*/
func (class) _area_get_collision_layer(impl func(ptr unsafe.Pointer, area gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_collision_mask].
*/
func (class) _area_set_collision_mask(impl func(ptr unsafe.Pointer, area gd.RID, mask gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var mask = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, mask)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_get_collision_mask].
*/
func (class) _area_get_collision_mask(impl func(ptr unsafe.Pointer, area gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_monitorable].
*/
func (class) _area_set_monitorable(impl func(ptr unsafe.Pointer, area gd.RID, monitorable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var monitorable = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, monitorable)
	}
}

/*
If set to [code]true[/code], allows the area with the given [RID] to detect mouse inputs when the mouse cursor is hovering on it.
Overridable version of [PhysicsServer2D]'s internal [code]area_set_pickable[/code] method. Corresponds to [member CollisionObject2D.input_pickable].
*/
func (class) _area_set_pickable(impl func(ptr unsafe.Pointer, area gd.RID, pickable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var pickable = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, pickable)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_monitor_callback].
*/
func (class) _area_set_monitor_callback(impl func(ptr unsafe.Pointer, area gd.RID, callback Callable.Function)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var callback = Callable.Through(gd.CallableProxy{}, pointers.Pack(pointers.New[gd.Callable](gd.UnsafeGet[[2]uint64](p_args, 1))))
		defer pointers.End(gd.InternalCallable(callback))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, callback)
	}
}

/*
Overridable version of [method PhysicsServer2D.area_set_area_monitor_callback].
*/
func (class) _area_set_area_monitor_callback(impl func(ptr unsafe.Pointer, area gd.RID, callback Callable.Function)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)

		var callback = Callable.Through(gd.CallableProxy{}, pointers.Pack(pointers.New[gd.Callable](gd.UnsafeGet[[2]uint64](p_args, 1))))
		defer pointers.End(gd.InternalCallable(callback))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, callback)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_create].
*/
func (class) _body_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_space].
*/
func (class) _body_set_space(impl func(ptr unsafe.Pointer, body gd.RID, space gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var space = gd.UnsafeGet[gd.RID](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, space)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_space].
*/
func (class) _body_get_space(impl func(ptr unsafe.Pointer, body gd.RID) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_mode].
*/
func (class) _body_set_mode(impl func(ptr unsafe.Pointer, body gd.RID, mode gdclass.PhysicsServer2DBodyMode)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var mode = gd.UnsafeGet[gdclass.PhysicsServer2DBodyMode](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, mode)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_mode].
*/
func (class) _body_get_mode(impl func(ptr unsafe.Pointer, body gd.RID) gdclass.PhysicsServer2DBodyMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_add_shape].
*/
func (class) _body_add_shape(impl func(ptr unsafe.Pointer, body gd.RID, shape gd.RID, transform gd.Transform2D, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape = gd.UnsafeGet[gd.RID](p_args, 1)

		var transform = gd.UnsafeGet[gd.Transform2D](p_args, 2)

		var disabled = gd.UnsafeGet[bool](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, shape, transform, disabled)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_shape].
*/
func (class) _body_set_shape(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx gd.Int, shape gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)

		var shape = gd.UnsafeGet[gd.RID](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, shape_idx, shape)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_shape_transform].
*/
func (class) _body_set_shape_transform(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx gd.Int, transform gd.Transform2D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)

		var transform = gd.UnsafeGet[gd.Transform2D](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, shape_idx, transform)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_shape_count].
*/
func (class) _body_get_shape_count(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_shape].
*/
func (class) _body_get_shape(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx gd.Int) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, shape_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_shape_transform].
*/
func (class) _body_get_shape_transform(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx gd.Int) gd.Transform2D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, shape_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_shape_disabled].
*/
func (class) _body_set_shape_disabled(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx gd.Int, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)

		var disabled = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, shape_idx, disabled)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_shape_as_one_way_collision].
*/
func (class) _body_set_shape_as_one_way_collision(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx gd.Int, enable bool, margin gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)

		var enable = gd.UnsafeGet[bool](p_args, 2)

		var margin = gd.UnsafeGet[gd.Float](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, shape_idx, enable, margin)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_remove_shape].
*/
func (class) _body_remove_shape(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, shape_idx)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_clear_shapes].
*/
func (class) _body_clear_shapes(impl func(ptr unsafe.Pointer, body gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_attach_object_instance_id].
*/
func (class) _body_attach_object_instance_id(impl func(ptr unsafe.Pointer, body gd.RID, id gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var id = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, id)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_object_instance_id].
*/
func (class) _body_get_object_instance_id(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_attach_canvas_instance_id].
*/
func (class) _body_attach_canvas_instance_id(impl func(ptr unsafe.Pointer, body gd.RID, id gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var id = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, id)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_canvas_instance_id].
*/
func (class) _body_get_canvas_instance_id(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_continuous_collision_detection_mode].
*/
func (class) _body_set_continuous_collision_detection_mode(impl func(ptr unsafe.Pointer, body gd.RID, mode gdclass.PhysicsServer2DCCDMode)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var mode = gd.UnsafeGet[gdclass.PhysicsServer2DCCDMode](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, mode)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_continuous_collision_detection_mode].
*/
func (class) _body_get_continuous_collision_detection_mode(impl func(ptr unsafe.Pointer, body gd.RID) gdclass.PhysicsServer2DCCDMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_collision_layer].
*/
func (class) _body_set_collision_layer(impl func(ptr unsafe.Pointer, body gd.RID, layer gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var layer = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, layer)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_collision_layer].
*/
func (class) _body_get_collision_layer(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_collision_mask].
*/
func (class) _body_set_collision_mask(impl func(ptr unsafe.Pointer, body gd.RID, mask gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var mask = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, mask)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_collision_mask].
*/
func (class) _body_get_collision_mask(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_collision_priority].
*/
func (class) _body_set_collision_priority(impl func(ptr unsafe.Pointer, body gd.RID, priority gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var priority = gd.UnsafeGet[gd.Float](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, priority)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_collision_priority].
*/
func (class) _body_get_collision_priority(impl func(ptr unsafe.Pointer, body gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_param].
*/
func (class) _body_set_param(impl func(ptr unsafe.Pointer, body gd.RID, param gdclass.PhysicsServer2DBodyParameter, value gd.Variant)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer2DBodyParameter](p_args, 1)

		var value = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 2))
		defer pointers.End(value)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, param, value)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_param].
*/
func (class) _body_get_param(impl func(ptr unsafe.Pointer, body gd.RID, param gdclass.PhysicsServer2DBodyParameter) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer2DBodyParameter](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, param)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_reset_mass_properties].
*/
func (class) _body_reset_mass_properties(impl func(ptr unsafe.Pointer, body gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_state].
*/
func (class) _body_set_state(impl func(ptr unsafe.Pointer, body gd.RID, state gdclass.PhysicsServer2DBodyState, value gd.Variant)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var state = gd.UnsafeGet[gdclass.PhysicsServer2DBodyState](p_args, 1)

		var value = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 2))
		defer pointers.End(value)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, state, value)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_state].
*/
func (class) _body_get_state(impl func(ptr unsafe.Pointer, body gd.RID, state gdclass.PhysicsServer2DBodyState) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var state = gd.UnsafeGet[gdclass.PhysicsServer2DBodyState](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, state)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_apply_central_impulse].
*/
func (class) _body_apply_central_impulse(impl func(ptr unsafe.Pointer, body gd.RID, impulse gd.Vector2)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var impulse = gd.UnsafeGet[gd.Vector2](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, impulse)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_apply_torque_impulse].
*/
func (class) _body_apply_torque_impulse(impl func(ptr unsafe.Pointer, body gd.RID, impulse gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var impulse = gd.UnsafeGet[gd.Float](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, impulse)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_apply_impulse].
*/
func (class) _body_apply_impulse(impl func(ptr unsafe.Pointer, body gd.RID, impulse gd.Vector2, position gd.Vector2)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var impulse = gd.UnsafeGet[gd.Vector2](p_args, 1)

		var position = gd.UnsafeGet[gd.Vector2](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, impulse, position)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_apply_central_force].
*/
func (class) _body_apply_central_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector2)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var force = gd.UnsafeGet[gd.Vector2](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_apply_force].
*/
func (class) _body_apply_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector2, position gd.Vector2)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var force = gd.UnsafeGet[gd.Vector2](p_args, 1)

		var position = gd.UnsafeGet[gd.Vector2](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force, position)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_apply_torque].
*/
func (class) _body_apply_torque(impl func(ptr unsafe.Pointer, body gd.RID, torque gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var torque = gd.UnsafeGet[gd.Float](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, torque)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_add_constant_central_force].
*/
func (class) _body_add_constant_central_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector2)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var force = gd.UnsafeGet[gd.Vector2](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_add_constant_force].
*/
func (class) _body_add_constant_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector2, position gd.Vector2)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var force = gd.UnsafeGet[gd.Vector2](p_args, 1)

		var position = gd.UnsafeGet[gd.Vector2](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force, position)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_add_constant_torque].
*/
func (class) _body_add_constant_torque(impl func(ptr unsafe.Pointer, body gd.RID, torque gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var torque = gd.UnsafeGet[gd.Float](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, torque)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_constant_force].
*/
func (class) _body_set_constant_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector2)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var force = gd.UnsafeGet[gd.Vector2](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_constant_force].
*/
func (class) _body_get_constant_force(impl func(ptr unsafe.Pointer, body gd.RID) gd.Vector2) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_constant_torque].
*/
func (class) _body_set_constant_torque(impl func(ptr unsafe.Pointer, body gd.RID, torque gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var torque = gd.UnsafeGet[gd.Float](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, torque)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_constant_torque].
*/
func (class) _body_get_constant_torque(impl func(ptr unsafe.Pointer, body gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_axis_velocity].
*/
func (class) _body_set_axis_velocity(impl func(ptr unsafe.Pointer, body gd.RID, axis_velocity gd.Vector2)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var axis_velocity = gd.UnsafeGet[gd.Vector2](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, axis_velocity)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_add_collision_exception].
*/
func (class) _body_add_collision_exception(impl func(ptr unsafe.Pointer, body gd.RID, excepted_body gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var excepted_body = gd.UnsafeGet[gd.RID](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, excepted_body)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_remove_collision_exception].
*/
func (class) _body_remove_collision_exception(impl func(ptr unsafe.Pointer, body gd.RID, excepted_body gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var excepted_body = gd.UnsafeGet[gd.RID](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, excepted_body)
	}
}

/*
Returns the [RID]s of all bodies added as collision exceptions for the given [param body]. See also [method _body_add_collision_exception] and [method _body_remove_collision_exception].
Overridable version of [PhysicsServer2D]'s internal [code]body_get_collision_exceptions[/code] method. Corresponds to [method PhysicsBody2D.get_collision_exceptions].
*/
func (class) _body_get_collision_exceptions(impl func(ptr unsafe.Pointer, body gd.RID) Array.Contains[gd.RID]) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		ptr, ok := pointers.End(gd.InternalArray(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_max_contacts_reported].
*/
func (class) _body_set_max_contacts_reported(impl func(ptr unsafe.Pointer, body gd.RID, amount gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var amount = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, amount)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_max_contacts_reported].
*/
func (class) _body_get_max_contacts_reported(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [PhysicsServer2D]'s internal [code]body_set_contacts_reported_depth_threshold[/code] method.
[b]Note:[/b] This method is currently unused by Godot's default physics implementation.
*/
func (class) _body_set_contacts_reported_depth_threshold(impl func(ptr unsafe.Pointer, body gd.RID, threshold gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var threshold = gd.UnsafeGet[gd.Float](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, threshold)
	}
}

/*
Overridable version of [PhysicsServer2D]'s internal [code]body_get_contacts_reported_depth_threshold[/code] method.
[b]Note:[/b] This method is currently unused by Godot's default physics implementation.
*/
func (class) _body_get_contacts_reported_depth_threshold(impl func(ptr unsafe.Pointer, body gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_omit_force_integration].
*/
func (class) _body_set_omit_force_integration(impl func(ptr unsafe.Pointer, body gd.RID, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var enable = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, enable)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_is_omitting_force_integration].
*/
func (class) _body_is_omitting_force_integration(impl func(ptr unsafe.Pointer, body gd.RID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Assigns the [param body] to call the given [param callable] during the synchronization phase of the loop, before [method _step] is called. See also [method _sync].
Overridable version of [method PhysicsServer2D.body_set_state_sync_callback].
*/
func (class) _body_set_state_sync_callback(impl func(ptr unsafe.Pointer, body gd.RID, callable Callable.Function)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var callable = Callable.Through(gd.CallableProxy{}, pointers.Pack(pointers.New[gd.Callable](gd.UnsafeGet[[2]uint64](p_args, 1))))
		defer pointers.End(gd.InternalCallable(callable))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, callable)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_set_force_integration_callback].
*/
func (class) _body_set_force_integration_callback(impl func(ptr unsafe.Pointer, body gd.RID, callable Callable.Function, userdata gd.Variant)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var callable = Callable.Through(gd.CallableProxy{}, pointers.Pack(pointers.New[gd.Callable](gd.UnsafeGet[[2]uint64](p_args, 1))))
		defer pointers.End(gd.InternalCallable(callable))
		var userdata = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 2))
		defer pointers.End(userdata)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, callable, userdata)
	}
}

/*
Given a [param body], a [param shape], and their respective parameters, this method should return [code]true[/code] if a collision between the two would occur, with additional details passed in [param results].
Overridable version of [PhysicsServer2D]'s internal [code]shape_collide[/code] method. Corresponds to [method PhysicsDirectSpaceState2D.collide_shape].
*/
func (class) _body_collide_shape(impl func(ptr unsafe.Pointer, body gd.RID, body_shape gd.Int, shape gd.RID, shape_xform gd.Transform2D, motion gd.Vector2, results unsafe.Pointer, result_max gd.Int, result_count *int32) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var body_shape = gd.UnsafeGet[gd.Int](p_args, 1)

		var shape = gd.UnsafeGet[gd.RID](p_args, 2)

		var shape_xform = gd.UnsafeGet[gd.Transform2D](p_args, 3)

		var motion = gd.UnsafeGet[gd.Vector2](p_args, 4)

		var results = gd.UnsafeGet[unsafe.Pointer](p_args, 5)

		var result_max = gd.UnsafeGet[gd.Int](p_args, 6)

		var result_count = gd.UnsafeGet[*int32](p_args, 7)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, body_shape, shape, shape_xform, motion, results, result_max, result_count)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
If set to [code]true[/code], allows the body with the given [RID] to detect mouse inputs when the mouse cursor is hovering on it.
Overridable version of [PhysicsServer2D]'s internal [code]body_set_pickable[/code] method. Corresponds to [member CollisionObject2D.input_pickable].
*/
func (class) _body_set_pickable(impl func(ptr unsafe.Pointer, body gd.RID, pickable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var pickable = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, pickable)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_get_direct_state].
*/
func (class) _body_get_direct_state(impl func(ptr unsafe.Pointer, body gd.RID) [1]gdclass.PhysicsDirectBodyState2D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Overridable version of [method PhysicsServer2D.body_test_motion]. Unlike the exposed implementation, this method does not receive all of the arguments inside a [PhysicsTestMotionParameters2D].
*/
func (class) _body_test_motion(impl func(ptr unsafe.Pointer, body gd.RID, from gd.Transform2D, motion gd.Vector2, margin gd.Float, collide_separation_ray bool, recovery_as_collision bool, result *MotionResult) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)

		var from = gd.UnsafeGet[gd.Transform2D](p_args, 1)

		var motion = gd.UnsafeGet[gd.Vector2](p_args, 2)

		var margin = gd.UnsafeGet[gd.Float](p_args, 3)

		var collide_separation_ray = gd.UnsafeGet[bool](p_args, 4)

		var recovery_as_collision = gd.UnsafeGet[bool](p_args, 5)

		var result = gd.UnsafeGet[*MotionResult](p_args, 6)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, from, motion, margin, collide_separation_ray, recovery_as_collision, result)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.joint_create].
*/
func (class) _joint_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.joint_clear].
*/
func (class) _joint_clear(impl func(ptr unsafe.Pointer, joint gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint)
	}
}

/*
Overridable version of [method PhysicsServer2D.joint_set_param].
*/
func (class) _joint_set_param(impl func(ptr unsafe.Pointer, joint gd.RID, param gdclass.PhysicsServer2DJointParam, value gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer2DJointParam](p_args, 1)

		var value = gd.UnsafeGet[gd.Float](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, value)
	}
}

/*
Overridable version of [method PhysicsServer2D.joint_get_param].
*/
func (class) _joint_get_param(impl func(ptr unsafe.Pointer, joint gd.RID, param gdclass.PhysicsServer2DJointParam) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer2DJointParam](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.joint_disable_collisions_between_bodies].
*/
func (class) _joint_disable_collisions_between_bodies(impl func(ptr unsafe.Pointer, joint gd.RID, disable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		var disable = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, disable)
	}
}

/*
Overridable version of [method PhysicsServer2D.joint_is_disabled_collisions_between_bodies].
*/
func (class) _joint_is_disabled_collisions_between_bodies(impl func(ptr unsafe.Pointer, joint gd.RID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.joint_make_pin].
*/
func (class) _joint_make_pin(impl func(ptr unsafe.Pointer, joint gd.RID, anchor gd.Vector2, body_a gd.RID, body_b gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		var anchor = gd.UnsafeGet[gd.Vector2](p_args, 1)

		var body_a = gd.UnsafeGet[gd.RID](p_args, 2)

		var body_b = gd.UnsafeGet[gd.RID](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, anchor, body_a, body_b)
	}
}

/*
Overridable version of [method PhysicsServer2D.joint_make_groove].
*/
func (class) _joint_make_groove(impl func(ptr unsafe.Pointer, joint gd.RID, a_groove1 gd.Vector2, a_groove2 gd.Vector2, b_anchor gd.Vector2, body_a gd.RID, body_b gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		var a_groove1 = gd.UnsafeGet[gd.Vector2](p_args, 1)

		var a_groove2 = gd.UnsafeGet[gd.Vector2](p_args, 2)

		var b_anchor = gd.UnsafeGet[gd.Vector2](p_args, 3)

		var body_a = gd.UnsafeGet[gd.RID](p_args, 4)

		var body_b = gd.UnsafeGet[gd.RID](p_args, 5)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, a_groove1, a_groove2, b_anchor, body_a, body_b)
	}
}

/*
Overridable version of [method PhysicsServer2D.joint_make_damped_spring].
*/
func (class) _joint_make_damped_spring(impl func(ptr unsafe.Pointer, joint gd.RID, anchor_a gd.Vector2, anchor_b gd.Vector2, body_a gd.RID, body_b gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		var anchor_a = gd.UnsafeGet[gd.Vector2](p_args, 1)

		var anchor_b = gd.UnsafeGet[gd.Vector2](p_args, 2)

		var body_a = gd.UnsafeGet[gd.RID](p_args, 3)

		var body_b = gd.UnsafeGet[gd.RID](p_args, 4)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, anchor_a, anchor_b, body_a, body_b)
	}
}

/*
Overridable version of [method PhysicsServer2D.pin_joint_set_flag].
*/
func (class) _pin_joint_set_flag(impl func(ptr unsafe.Pointer, joint gd.RID, flag gdclass.PhysicsServer2DPinJointFlag, enabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		var flag = gd.UnsafeGet[gdclass.PhysicsServer2DPinJointFlag](p_args, 1)

		var enabled = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, flag, enabled)
	}
}

/*
Overridable version of [method PhysicsServer2D.pin_joint_get_flag].
*/
func (class) _pin_joint_get_flag(impl func(ptr unsafe.Pointer, joint gd.RID, flag gdclass.PhysicsServer2DPinJointFlag) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		var flag = gd.UnsafeGet[gdclass.PhysicsServer2DPinJointFlag](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, flag)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.pin_joint_set_param].
*/
func (class) _pin_joint_set_param(impl func(ptr unsafe.Pointer, joint gd.RID, param gdclass.PhysicsServer2DPinJointParam, value gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer2DPinJointParam](p_args, 1)

		var value = gd.UnsafeGet[gd.Float](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, value)
	}
}

/*
Overridable version of [method PhysicsServer2D.pin_joint_get_param].
*/
func (class) _pin_joint_get_param(impl func(ptr unsafe.Pointer, joint gd.RID, param gdclass.PhysicsServer2DPinJointParam) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer2DPinJointParam](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.damped_spring_joint_set_param].
*/
func (class) _damped_spring_joint_set_param(impl func(ptr unsafe.Pointer, joint gd.RID, param gdclass.PhysicsServer2DDampedSpringParam, value gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer2DDampedSpringParam](p_args, 1)

		var value = gd.UnsafeGet[gd.Float](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, value)
	}
}

/*
Overridable version of [method PhysicsServer2D.damped_spring_joint_get_param].
*/
func (class) _damped_spring_joint_get_param(impl func(ptr unsafe.Pointer, joint gd.RID, param gdclass.PhysicsServer2DDampedSpringParam) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer2DDampedSpringParam](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.joint_get_type].
*/
func (class) _joint_get_type(impl func(ptr unsafe.Pointer, joint gd.RID) gdclass.PhysicsServer2DJointType) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.free_rid].
*/
func (class) _free_rid(impl func(ptr unsafe.Pointer, rid gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var rid = gd.UnsafeGet[gd.RID](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, rid)
	}
}

/*
Overridable version of [method PhysicsServer2D.set_active].
*/
func (class) _set_active(impl func(ptr unsafe.Pointer, active bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var active = gd.UnsafeGet[bool](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, active)
	}
}

/*
Called when the main loop is initialized and creates a new instance of this physics server. See also [method MainLoop._initialize] and [method _finish].
Overridable version of [PhysicsServer2D]'s internal [code]init[/code] method.
*/
func (class) _init(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called every physics step to process the physics simulation. [param step] is the time elapsed since the last physics step, in seconds. It is usually the same as [method Node.get_physics_process_delta_time].
Overridable version of [PhysicsServer2D]'s internal [code skip-lint]step[/code] method.
*/
func (class) _step(impl func(ptr unsafe.Pointer, step gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var step = gd.UnsafeGet[gd.Float](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, step)
	}
}

/*
Called to indicate that the physics server is synchronizing and cannot access physics states if running on a separate thread. See also [method _end_sync].
Overridable version of [PhysicsServer2D]'s internal [code]sync[/code] method.
*/
func (class) _sync(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called every physics step before [method _step] to process all remaining queries.
Overridable version of [PhysicsServer2D]'s internal [code]flush_queries[/code] method.
*/
func (class) _flush_queries(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called to indicate that the physics server has stopped synchronizing. It is in the loop's iteration/physics phase, and can access physics objects even if running on a separate thread. See also [method _sync].
Overridable version of [PhysicsServer2D]'s internal [code]end_sync[/code] method.
*/
func (class) _end_sync(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the main loop finalizes to shut down the physics server. See also [method MainLoop._finalize] and [method _init].
Overridable version of [PhysicsServer2D]'s internal [code]finish[/code] method.
*/
func (class) _finish(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Overridable method that should return [code]true[/code] when the physics server is processing queries. See also [method _flush_queries].
Overridable version of [PhysicsServer2D]'s internal [code]is_flushing_queries[/code] method.
*/
func (class) _is_flushing_queries(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable version of [method PhysicsServer2D.get_process_info].
*/
func (class) _get_process_info(impl func(ptr unsafe.Pointer, process_info gdclass.PhysicsServer2DProcessInfo) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var process_info = gd.UnsafeGet[gdclass.PhysicsServer2DProcessInfo](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, process_info)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns [code]true[/code] if the body with the given [RID] is being excluded from [method _body_test_motion]. See also [method Object.get_instance_id].
*/
//go:nosplit
func (self class) BodyTestMotionIsExcludingBody(body gd.RID) bool { //gd:PhysicsServer2DExtension.body_test_motion_is_excluding_body
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsServer2DExtension.Bind_body_test_motion_is_excluding_body, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the object with the given instance ID is being excluded from [method _body_test_motion]. See also [method Object.get_instance_id].
*/
//go:nosplit
func (self class) BodyTestMotionIsExcludingObject(obj gd.Int) bool { //gd:PhysicsServer2DExtension.body_test_motion_is_excluding_object
	var frame = callframe.New()
	callframe.Arg(frame, obj)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsServer2DExtension.Bind_body_test_motion_is_excluding_object, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPhysicsServer2DExtension() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPhysicsServer2DExtension() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_world_boundary_shape_create":
		return reflect.ValueOf(self._world_boundary_shape_create)
	case "_separation_ray_shape_create":
		return reflect.ValueOf(self._separation_ray_shape_create)
	case "_segment_shape_create":
		return reflect.ValueOf(self._segment_shape_create)
	case "_circle_shape_create":
		return reflect.ValueOf(self._circle_shape_create)
	case "_rectangle_shape_create":
		return reflect.ValueOf(self._rectangle_shape_create)
	case "_capsule_shape_create":
		return reflect.ValueOf(self._capsule_shape_create)
	case "_convex_polygon_shape_create":
		return reflect.ValueOf(self._convex_polygon_shape_create)
	case "_concave_polygon_shape_create":
		return reflect.ValueOf(self._concave_polygon_shape_create)
	case "_shape_set_data":
		return reflect.ValueOf(self._shape_set_data)
	case "_shape_set_custom_solver_bias":
		return reflect.ValueOf(self._shape_set_custom_solver_bias)
	case "_shape_get_type":
		return reflect.ValueOf(self._shape_get_type)
	case "_shape_get_data":
		return reflect.ValueOf(self._shape_get_data)
	case "_shape_get_custom_solver_bias":
		return reflect.ValueOf(self._shape_get_custom_solver_bias)
	case "_shape_collide":
		return reflect.ValueOf(self._shape_collide)
	case "_space_create":
		return reflect.ValueOf(self._space_create)
	case "_space_set_active":
		return reflect.ValueOf(self._space_set_active)
	case "_space_is_active":
		return reflect.ValueOf(self._space_is_active)
	case "_space_set_param":
		return reflect.ValueOf(self._space_set_param)
	case "_space_get_param":
		return reflect.ValueOf(self._space_get_param)
	case "_space_get_direct_state":
		return reflect.ValueOf(self._space_get_direct_state)
	case "_space_set_debug_contacts":
		return reflect.ValueOf(self._space_set_debug_contacts)
	case "_space_get_contacts":
		return reflect.ValueOf(self._space_get_contacts)
	case "_space_get_contact_count":
		return reflect.ValueOf(self._space_get_contact_count)
	case "_area_create":
		return reflect.ValueOf(self._area_create)
	case "_area_set_space":
		return reflect.ValueOf(self._area_set_space)
	case "_area_get_space":
		return reflect.ValueOf(self._area_get_space)
	case "_area_add_shape":
		return reflect.ValueOf(self._area_add_shape)
	case "_area_set_shape":
		return reflect.ValueOf(self._area_set_shape)
	case "_area_set_shape_transform":
		return reflect.ValueOf(self._area_set_shape_transform)
	case "_area_set_shape_disabled":
		return reflect.ValueOf(self._area_set_shape_disabled)
	case "_area_get_shape_count":
		return reflect.ValueOf(self._area_get_shape_count)
	case "_area_get_shape":
		return reflect.ValueOf(self._area_get_shape)
	case "_area_get_shape_transform":
		return reflect.ValueOf(self._area_get_shape_transform)
	case "_area_remove_shape":
		return reflect.ValueOf(self._area_remove_shape)
	case "_area_clear_shapes":
		return reflect.ValueOf(self._area_clear_shapes)
	case "_area_attach_object_instance_id":
		return reflect.ValueOf(self._area_attach_object_instance_id)
	case "_area_get_object_instance_id":
		return reflect.ValueOf(self._area_get_object_instance_id)
	case "_area_attach_canvas_instance_id":
		return reflect.ValueOf(self._area_attach_canvas_instance_id)
	case "_area_get_canvas_instance_id":
		return reflect.ValueOf(self._area_get_canvas_instance_id)
	case "_area_set_param":
		return reflect.ValueOf(self._area_set_param)
	case "_area_set_transform":
		return reflect.ValueOf(self._area_set_transform)
	case "_area_get_param":
		return reflect.ValueOf(self._area_get_param)
	case "_area_get_transform":
		return reflect.ValueOf(self._area_get_transform)
	case "_area_set_collision_layer":
		return reflect.ValueOf(self._area_set_collision_layer)
	case "_area_get_collision_layer":
		return reflect.ValueOf(self._area_get_collision_layer)
	case "_area_set_collision_mask":
		return reflect.ValueOf(self._area_set_collision_mask)
	case "_area_get_collision_mask":
		return reflect.ValueOf(self._area_get_collision_mask)
	case "_area_set_monitorable":
		return reflect.ValueOf(self._area_set_monitorable)
	case "_area_set_pickable":
		return reflect.ValueOf(self._area_set_pickable)
	case "_area_set_monitor_callback":
		return reflect.ValueOf(self._area_set_monitor_callback)
	case "_area_set_area_monitor_callback":
		return reflect.ValueOf(self._area_set_area_monitor_callback)
	case "_body_create":
		return reflect.ValueOf(self._body_create)
	case "_body_set_space":
		return reflect.ValueOf(self._body_set_space)
	case "_body_get_space":
		return reflect.ValueOf(self._body_get_space)
	case "_body_set_mode":
		return reflect.ValueOf(self._body_set_mode)
	case "_body_get_mode":
		return reflect.ValueOf(self._body_get_mode)
	case "_body_add_shape":
		return reflect.ValueOf(self._body_add_shape)
	case "_body_set_shape":
		return reflect.ValueOf(self._body_set_shape)
	case "_body_set_shape_transform":
		return reflect.ValueOf(self._body_set_shape_transform)
	case "_body_get_shape_count":
		return reflect.ValueOf(self._body_get_shape_count)
	case "_body_get_shape":
		return reflect.ValueOf(self._body_get_shape)
	case "_body_get_shape_transform":
		return reflect.ValueOf(self._body_get_shape_transform)
	case "_body_set_shape_disabled":
		return reflect.ValueOf(self._body_set_shape_disabled)
	case "_body_set_shape_as_one_way_collision":
		return reflect.ValueOf(self._body_set_shape_as_one_way_collision)
	case "_body_remove_shape":
		return reflect.ValueOf(self._body_remove_shape)
	case "_body_clear_shapes":
		return reflect.ValueOf(self._body_clear_shapes)
	case "_body_attach_object_instance_id":
		return reflect.ValueOf(self._body_attach_object_instance_id)
	case "_body_get_object_instance_id":
		return reflect.ValueOf(self._body_get_object_instance_id)
	case "_body_attach_canvas_instance_id":
		return reflect.ValueOf(self._body_attach_canvas_instance_id)
	case "_body_get_canvas_instance_id":
		return reflect.ValueOf(self._body_get_canvas_instance_id)
	case "_body_set_continuous_collision_detection_mode":
		return reflect.ValueOf(self._body_set_continuous_collision_detection_mode)
	case "_body_get_continuous_collision_detection_mode":
		return reflect.ValueOf(self._body_get_continuous_collision_detection_mode)
	case "_body_set_collision_layer":
		return reflect.ValueOf(self._body_set_collision_layer)
	case "_body_get_collision_layer":
		return reflect.ValueOf(self._body_get_collision_layer)
	case "_body_set_collision_mask":
		return reflect.ValueOf(self._body_set_collision_mask)
	case "_body_get_collision_mask":
		return reflect.ValueOf(self._body_get_collision_mask)
	case "_body_set_collision_priority":
		return reflect.ValueOf(self._body_set_collision_priority)
	case "_body_get_collision_priority":
		return reflect.ValueOf(self._body_get_collision_priority)
	case "_body_set_param":
		return reflect.ValueOf(self._body_set_param)
	case "_body_get_param":
		return reflect.ValueOf(self._body_get_param)
	case "_body_reset_mass_properties":
		return reflect.ValueOf(self._body_reset_mass_properties)
	case "_body_set_state":
		return reflect.ValueOf(self._body_set_state)
	case "_body_get_state":
		return reflect.ValueOf(self._body_get_state)
	case "_body_apply_central_impulse":
		return reflect.ValueOf(self._body_apply_central_impulse)
	case "_body_apply_torque_impulse":
		return reflect.ValueOf(self._body_apply_torque_impulse)
	case "_body_apply_impulse":
		return reflect.ValueOf(self._body_apply_impulse)
	case "_body_apply_central_force":
		return reflect.ValueOf(self._body_apply_central_force)
	case "_body_apply_force":
		return reflect.ValueOf(self._body_apply_force)
	case "_body_apply_torque":
		return reflect.ValueOf(self._body_apply_torque)
	case "_body_add_constant_central_force":
		return reflect.ValueOf(self._body_add_constant_central_force)
	case "_body_add_constant_force":
		return reflect.ValueOf(self._body_add_constant_force)
	case "_body_add_constant_torque":
		return reflect.ValueOf(self._body_add_constant_torque)
	case "_body_set_constant_force":
		return reflect.ValueOf(self._body_set_constant_force)
	case "_body_get_constant_force":
		return reflect.ValueOf(self._body_get_constant_force)
	case "_body_set_constant_torque":
		return reflect.ValueOf(self._body_set_constant_torque)
	case "_body_get_constant_torque":
		return reflect.ValueOf(self._body_get_constant_torque)
	case "_body_set_axis_velocity":
		return reflect.ValueOf(self._body_set_axis_velocity)
	case "_body_add_collision_exception":
		return reflect.ValueOf(self._body_add_collision_exception)
	case "_body_remove_collision_exception":
		return reflect.ValueOf(self._body_remove_collision_exception)
	case "_body_get_collision_exceptions":
		return reflect.ValueOf(self._body_get_collision_exceptions)
	case "_body_set_max_contacts_reported":
		return reflect.ValueOf(self._body_set_max_contacts_reported)
	case "_body_get_max_contacts_reported":
		return reflect.ValueOf(self._body_get_max_contacts_reported)
	case "_body_set_contacts_reported_depth_threshold":
		return reflect.ValueOf(self._body_set_contacts_reported_depth_threshold)
	case "_body_get_contacts_reported_depth_threshold":
		return reflect.ValueOf(self._body_get_contacts_reported_depth_threshold)
	case "_body_set_omit_force_integration":
		return reflect.ValueOf(self._body_set_omit_force_integration)
	case "_body_is_omitting_force_integration":
		return reflect.ValueOf(self._body_is_omitting_force_integration)
	case "_body_set_state_sync_callback":
		return reflect.ValueOf(self._body_set_state_sync_callback)
	case "_body_set_force_integration_callback":
		return reflect.ValueOf(self._body_set_force_integration_callback)
	case "_body_collide_shape":
		return reflect.ValueOf(self._body_collide_shape)
	case "_body_set_pickable":
		return reflect.ValueOf(self._body_set_pickable)
	case "_body_get_direct_state":
		return reflect.ValueOf(self._body_get_direct_state)
	case "_body_test_motion":
		return reflect.ValueOf(self._body_test_motion)
	case "_joint_create":
		return reflect.ValueOf(self._joint_create)
	case "_joint_clear":
		return reflect.ValueOf(self._joint_clear)
	case "_joint_set_param":
		return reflect.ValueOf(self._joint_set_param)
	case "_joint_get_param":
		return reflect.ValueOf(self._joint_get_param)
	case "_joint_disable_collisions_between_bodies":
		return reflect.ValueOf(self._joint_disable_collisions_between_bodies)
	case "_joint_is_disabled_collisions_between_bodies":
		return reflect.ValueOf(self._joint_is_disabled_collisions_between_bodies)
	case "_joint_make_pin":
		return reflect.ValueOf(self._joint_make_pin)
	case "_joint_make_groove":
		return reflect.ValueOf(self._joint_make_groove)
	case "_joint_make_damped_spring":
		return reflect.ValueOf(self._joint_make_damped_spring)
	case "_pin_joint_set_flag":
		return reflect.ValueOf(self._pin_joint_set_flag)
	case "_pin_joint_get_flag":
		return reflect.ValueOf(self._pin_joint_get_flag)
	case "_pin_joint_set_param":
		return reflect.ValueOf(self._pin_joint_set_param)
	case "_pin_joint_get_param":
		return reflect.ValueOf(self._pin_joint_get_param)
	case "_damped_spring_joint_set_param":
		return reflect.ValueOf(self._damped_spring_joint_set_param)
	case "_damped_spring_joint_get_param":
		return reflect.ValueOf(self._damped_spring_joint_get_param)
	case "_joint_get_type":
		return reflect.ValueOf(self._joint_get_type)
	case "_free_rid":
		return reflect.ValueOf(self._free_rid)
	case "_set_active":
		return reflect.ValueOf(self._set_active)
	case "_init":
		return reflect.ValueOf(self._init)
	case "_step":
		return reflect.ValueOf(self._step)
	case "_sync":
		return reflect.ValueOf(self._sync)
	case "_flush_queries":
		return reflect.ValueOf(self._flush_queries)
	case "_end_sync":
		return reflect.ValueOf(self._end_sync)
	case "_finish":
		return reflect.ValueOf(self._finish)
	case "_is_flushing_queries":
		return reflect.ValueOf(self._is_flushing_queries)
	case "_get_process_info":
		return reflect.ValueOf(self._get_process_info)
	default:
		return reflect.Value{}
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_world_boundary_shape_create":
		return reflect.ValueOf(self._world_boundary_shape_create)
	case "_separation_ray_shape_create":
		return reflect.ValueOf(self._separation_ray_shape_create)
	case "_segment_shape_create":
		return reflect.ValueOf(self._segment_shape_create)
	case "_circle_shape_create":
		return reflect.ValueOf(self._circle_shape_create)
	case "_rectangle_shape_create":
		return reflect.ValueOf(self._rectangle_shape_create)
	case "_capsule_shape_create":
		return reflect.ValueOf(self._capsule_shape_create)
	case "_convex_polygon_shape_create":
		return reflect.ValueOf(self._convex_polygon_shape_create)
	case "_concave_polygon_shape_create":
		return reflect.ValueOf(self._concave_polygon_shape_create)
	case "_shape_set_data":
		return reflect.ValueOf(self._shape_set_data)
	case "_shape_set_custom_solver_bias":
		return reflect.ValueOf(self._shape_set_custom_solver_bias)
	case "_shape_get_type":
		return reflect.ValueOf(self._shape_get_type)
	case "_shape_get_data":
		return reflect.ValueOf(self._shape_get_data)
	case "_shape_get_custom_solver_bias":
		return reflect.ValueOf(self._shape_get_custom_solver_bias)
	case "_shape_collide":
		return reflect.ValueOf(self._shape_collide)
	case "_space_create":
		return reflect.ValueOf(self._space_create)
	case "_space_set_active":
		return reflect.ValueOf(self._space_set_active)
	case "_space_is_active":
		return reflect.ValueOf(self._space_is_active)
	case "_space_set_param":
		return reflect.ValueOf(self._space_set_param)
	case "_space_get_param":
		return reflect.ValueOf(self._space_get_param)
	case "_space_get_direct_state":
		return reflect.ValueOf(self._space_get_direct_state)
	case "_space_set_debug_contacts":
		return reflect.ValueOf(self._space_set_debug_contacts)
	case "_space_get_contacts":
		return reflect.ValueOf(self._space_get_contacts)
	case "_space_get_contact_count":
		return reflect.ValueOf(self._space_get_contact_count)
	case "_area_create":
		return reflect.ValueOf(self._area_create)
	case "_area_set_space":
		return reflect.ValueOf(self._area_set_space)
	case "_area_get_space":
		return reflect.ValueOf(self._area_get_space)
	case "_area_add_shape":
		return reflect.ValueOf(self._area_add_shape)
	case "_area_set_shape":
		return reflect.ValueOf(self._area_set_shape)
	case "_area_set_shape_transform":
		return reflect.ValueOf(self._area_set_shape_transform)
	case "_area_set_shape_disabled":
		return reflect.ValueOf(self._area_set_shape_disabled)
	case "_area_get_shape_count":
		return reflect.ValueOf(self._area_get_shape_count)
	case "_area_get_shape":
		return reflect.ValueOf(self._area_get_shape)
	case "_area_get_shape_transform":
		return reflect.ValueOf(self._area_get_shape_transform)
	case "_area_remove_shape":
		return reflect.ValueOf(self._area_remove_shape)
	case "_area_clear_shapes":
		return reflect.ValueOf(self._area_clear_shapes)
	case "_area_attach_object_instance_id":
		return reflect.ValueOf(self._area_attach_object_instance_id)
	case "_area_get_object_instance_id":
		return reflect.ValueOf(self._area_get_object_instance_id)
	case "_area_attach_canvas_instance_id":
		return reflect.ValueOf(self._area_attach_canvas_instance_id)
	case "_area_get_canvas_instance_id":
		return reflect.ValueOf(self._area_get_canvas_instance_id)
	case "_area_set_param":
		return reflect.ValueOf(self._area_set_param)
	case "_area_set_transform":
		return reflect.ValueOf(self._area_set_transform)
	case "_area_get_param":
		return reflect.ValueOf(self._area_get_param)
	case "_area_get_transform":
		return reflect.ValueOf(self._area_get_transform)
	case "_area_set_collision_layer":
		return reflect.ValueOf(self._area_set_collision_layer)
	case "_area_get_collision_layer":
		return reflect.ValueOf(self._area_get_collision_layer)
	case "_area_set_collision_mask":
		return reflect.ValueOf(self._area_set_collision_mask)
	case "_area_get_collision_mask":
		return reflect.ValueOf(self._area_get_collision_mask)
	case "_area_set_monitorable":
		return reflect.ValueOf(self._area_set_monitorable)
	case "_area_set_pickable":
		return reflect.ValueOf(self._area_set_pickable)
	case "_area_set_monitor_callback":
		return reflect.ValueOf(self._area_set_monitor_callback)
	case "_area_set_area_monitor_callback":
		return reflect.ValueOf(self._area_set_area_monitor_callback)
	case "_body_create":
		return reflect.ValueOf(self._body_create)
	case "_body_set_space":
		return reflect.ValueOf(self._body_set_space)
	case "_body_get_space":
		return reflect.ValueOf(self._body_get_space)
	case "_body_set_mode":
		return reflect.ValueOf(self._body_set_mode)
	case "_body_get_mode":
		return reflect.ValueOf(self._body_get_mode)
	case "_body_add_shape":
		return reflect.ValueOf(self._body_add_shape)
	case "_body_set_shape":
		return reflect.ValueOf(self._body_set_shape)
	case "_body_set_shape_transform":
		return reflect.ValueOf(self._body_set_shape_transform)
	case "_body_get_shape_count":
		return reflect.ValueOf(self._body_get_shape_count)
	case "_body_get_shape":
		return reflect.ValueOf(self._body_get_shape)
	case "_body_get_shape_transform":
		return reflect.ValueOf(self._body_get_shape_transform)
	case "_body_set_shape_disabled":
		return reflect.ValueOf(self._body_set_shape_disabled)
	case "_body_set_shape_as_one_way_collision":
		return reflect.ValueOf(self._body_set_shape_as_one_way_collision)
	case "_body_remove_shape":
		return reflect.ValueOf(self._body_remove_shape)
	case "_body_clear_shapes":
		return reflect.ValueOf(self._body_clear_shapes)
	case "_body_attach_object_instance_id":
		return reflect.ValueOf(self._body_attach_object_instance_id)
	case "_body_get_object_instance_id":
		return reflect.ValueOf(self._body_get_object_instance_id)
	case "_body_attach_canvas_instance_id":
		return reflect.ValueOf(self._body_attach_canvas_instance_id)
	case "_body_get_canvas_instance_id":
		return reflect.ValueOf(self._body_get_canvas_instance_id)
	case "_body_set_continuous_collision_detection_mode":
		return reflect.ValueOf(self._body_set_continuous_collision_detection_mode)
	case "_body_get_continuous_collision_detection_mode":
		return reflect.ValueOf(self._body_get_continuous_collision_detection_mode)
	case "_body_set_collision_layer":
		return reflect.ValueOf(self._body_set_collision_layer)
	case "_body_get_collision_layer":
		return reflect.ValueOf(self._body_get_collision_layer)
	case "_body_set_collision_mask":
		return reflect.ValueOf(self._body_set_collision_mask)
	case "_body_get_collision_mask":
		return reflect.ValueOf(self._body_get_collision_mask)
	case "_body_set_collision_priority":
		return reflect.ValueOf(self._body_set_collision_priority)
	case "_body_get_collision_priority":
		return reflect.ValueOf(self._body_get_collision_priority)
	case "_body_set_param":
		return reflect.ValueOf(self._body_set_param)
	case "_body_get_param":
		return reflect.ValueOf(self._body_get_param)
	case "_body_reset_mass_properties":
		return reflect.ValueOf(self._body_reset_mass_properties)
	case "_body_set_state":
		return reflect.ValueOf(self._body_set_state)
	case "_body_get_state":
		return reflect.ValueOf(self._body_get_state)
	case "_body_apply_central_impulse":
		return reflect.ValueOf(self._body_apply_central_impulse)
	case "_body_apply_torque_impulse":
		return reflect.ValueOf(self._body_apply_torque_impulse)
	case "_body_apply_impulse":
		return reflect.ValueOf(self._body_apply_impulse)
	case "_body_apply_central_force":
		return reflect.ValueOf(self._body_apply_central_force)
	case "_body_apply_force":
		return reflect.ValueOf(self._body_apply_force)
	case "_body_apply_torque":
		return reflect.ValueOf(self._body_apply_torque)
	case "_body_add_constant_central_force":
		return reflect.ValueOf(self._body_add_constant_central_force)
	case "_body_add_constant_force":
		return reflect.ValueOf(self._body_add_constant_force)
	case "_body_add_constant_torque":
		return reflect.ValueOf(self._body_add_constant_torque)
	case "_body_set_constant_force":
		return reflect.ValueOf(self._body_set_constant_force)
	case "_body_get_constant_force":
		return reflect.ValueOf(self._body_get_constant_force)
	case "_body_set_constant_torque":
		return reflect.ValueOf(self._body_set_constant_torque)
	case "_body_get_constant_torque":
		return reflect.ValueOf(self._body_get_constant_torque)
	case "_body_set_axis_velocity":
		return reflect.ValueOf(self._body_set_axis_velocity)
	case "_body_add_collision_exception":
		return reflect.ValueOf(self._body_add_collision_exception)
	case "_body_remove_collision_exception":
		return reflect.ValueOf(self._body_remove_collision_exception)
	case "_body_get_collision_exceptions":
		return reflect.ValueOf(self._body_get_collision_exceptions)
	case "_body_set_max_contacts_reported":
		return reflect.ValueOf(self._body_set_max_contacts_reported)
	case "_body_get_max_contacts_reported":
		return reflect.ValueOf(self._body_get_max_contacts_reported)
	case "_body_set_contacts_reported_depth_threshold":
		return reflect.ValueOf(self._body_set_contacts_reported_depth_threshold)
	case "_body_get_contacts_reported_depth_threshold":
		return reflect.ValueOf(self._body_get_contacts_reported_depth_threshold)
	case "_body_set_omit_force_integration":
		return reflect.ValueOf(self._body_set_omit_force_integration)
	case "_body_is_omitting_force_integration":
		return reflect.ValueOf(self._body_is_omitting_force_integration)
	case "_body_set_state_sync_callback":
		return reflect.ValueOf(self._body_set_state_sync_callback)
	case "_body_set_force_integration_callback":
		return reflect.ValueOf(self._body_set_force_integration_callback)
	case "_body_collide_shape":
		return reflect.ValueOf(self._body_collide_shape)
	case "_body_set_pickable":
		return reflect.ValueOf(self._body_set_pickable)
	case "_body_get_direct_state":
		return reflect.ValueOf(self._body_get_direct_state)
	case "_body_test_motion":
		return reflect.ValueOf(self._body_test_motion)
	case "_joint_create":
		return reflect.ValueOf(self._joint_create)
	case "_joint_clear":
		return reflect.ValueOf(self._joint_clear)
	case "_joint_set_param":
		return reflect.ValueOf(self._joint_set_param)
	case "_joint_get_param":
		return reflect.ValueOf(self._joint_get_param)
	case "_joint_disable_collisions_between_bodies":
		return reflect.ValueOf(self._joint_disable_collisions_between_bodies)
	case "_joint_is_disabled_collisions_between_bodies":
		return reflect.ValueOf(self._joint_is_disabled_collisions_between_bodies)
	case "_joint_make_pin":
		return reflect.ValueOf(self._joint_make_pin)
	case "_joint_make_groove":
		return reflect.ValueOf(self._joint_make_groove)
	case "_joint_make_damped_spring":
		return reflect.ValueOf(self._joint_make_damped_spring)
	case "_pin_joint_set_flag":
		return reflect.ValueOf(self._pin_joint_set_flag)
	case "_pin_joint_get_flag":
		return reflect.ValueOf(self._pin_joint_get_flag)
	case "_pin_joint_set_param":
		return reflect.ValueOf(self._pin_joint_set_param)
	case "_pin_joint_get_param":
		return reflect.ValueOf(self._pin_joint_get_param)
	case "_damped_spring_joint_set_param":
		return reflect.ValueOf(self._damped_spring_joint_set_param)
	case "_damped_spring_joint_get_param":
		return reflect.ValueOf(self._damped_spring_joint_get_param)
	case "_joint_get_type":
		return reflect.ValueOf(self._joint_get_type)
	case "_free_rid":
		return reflect.ValueOf(self._free_rid)
	case "_set_active":
		return reflect.ValueOf(self._set_active)
	case "_init":
		return reflect.ValueOf(self._init)
	case "_step":
		return reflect.ValueOf(self._step)
	case "_sync":
		return reflect.ValueOf(self._sync)
	case "_flush_queries":
		return reflect.ValueOf(self._flush_queries)
	case "_end_sync":
		return reflect.ValueOf(self._end_sync)
	case "_finish":
		return reflect.ValueOf(self._finish)
	case "_is_flushing_queries":
		return reflect.ValueOf(self._is_flushing_queries)
	case "_get_process_info":
		return reflect.ValueOf(self._get_process_info)
	default:
		return reflect.Value{}
	}
}
func init() {
	gdclass.Register("PhysicsServer2DExtension", func(ptr gd.Object) any {
		return [1]gdclass.PhysicsServer2DExtension{*(*gdclass.PhysicsServer2DExtension)(unsafe.Pointer(&ptr))}
	})
}
