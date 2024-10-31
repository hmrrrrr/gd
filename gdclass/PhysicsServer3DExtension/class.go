package PhysicsServer3DExtension

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
This class extends [PhysicsServer3D] by providing additional virtual methods that can be overridden. When these methods are overridden, they will be called instead of the internal methods of the physics server.
Intended for use with GDExtension to create custom implementations of [PhysicsServer3D].

	// PhysicsServer3DExtension methods that can be overridden by a [Class] that extends it.
	type PhysicsServer3DExtension interface {
		WorldBoundaryShapeCreate() gd.RID
		SeparationRayShapeCreate() gd.RID
		SphereShapeCreate() gd.RID
		BoxShapeCreate() gd.RID
		CapsuleShapeCreate() gd.RID
		CylinderShapeCreate() gd.RID
		ConvexPolygonShapeCreate() gd.RID
		ConcavePolygonShapeCreate() gd.RID
		HeightmapShapeCreate() gd.RID
		CustomShapeCreate() gd.RID
		ShapeSetData(shape gd.RID, data gd.Variant)
		ShapeSetCustomSolverBias(shape gd.RID, bias float64)
		ShapeSetMargin(shape gd.RID, margin float64)
		ShapeGetMargin(shape gd.RID) float64
		ShapeGetType(shape gd.RID) classdb.PhysicsServer3DShapeType
		ShapeGetData(shape gd.RID) gd.Variant
		ShapeGetCustomSolverBias(shape gd.RID) float64
		SpaceCreate() gd.RID
		SpaceSetActive(space gd.RID, active bool)
		SpaceIsActive(space gd.RID) bool
		SpaceSetParam(space gd.RID, param classdb.PhysicsServer3DSpaceParameter, value float64)
		SpaceGetParam(space gd.RID, param classdb.PhysicsServer3DSpaceParameter) float64
		SpaceGetDirectState(space gd.RID) gdclass.PhysicsDirectSpaceState3D
		SpaceSetDebugContacts(space gd.RID, max_contacts int)
		SpaceGetContacts(space gd.RID) []gd.Vector3
		SpaceGetContactCount(space gd.RID) int
		AreaCreate() gd.RID
		AreaSetSpace(area gd.RID, space gd.RID)
		AreaGetSpace(area gd.RID) gd.RID
		AreaAddShape(area gd.RID, shape gd.RID, transform gd.Transform3D, disabled bool)
		AreaSetShape(area gd.RID, shape_idx int, shape gd.RID)
		AreaSetShapeTransform(area gd.RID, shape_idx int, transform gd.Transform3D)
		AreaSetShapeDisabled(area gd.RID, shape_idx int, disabled bool)
		AreaGetShapeCount(area gd.RID) int
		AreaGetShape(area gd.RID, shape_idx int) gd.RID
		AreaGetShapeTransform(area gd.RID, shape_idx int) gd.Transform3D
		AreaRemoveShape(area gd.RID, shape_idx int)
		AreaClearShapes(area gd.RID)
		AreaAttachObjectInstanceId(area gd.RID, id int)
		AreaGetObjectInstanceId(area gd.RID) int
		AreaSetParam(area gd.RID, param classdb.PhysicsServer3DAreaParameter, value gd.Variant)
		AreaSetTransform(area gd.RID, transform gd.Transform3D)
		AreaGetParam(area gd.RID, param classdb.PhysicsServer3DAreaParameter) gd.Variant
		AreaGetTransform(area gd.RID) gd.Transform3D
		AreaSetCollisionLayer(area gd.RID, layer int)
		AreaGetCollisionLayer(area gd.RID) int
		AreaSetCollisionMask(area gd.RID, mask int)
		AreaGetCollisionMask(area gd.RID) int
		AreaSetMonitorable(area gd.RID, monitorable bool)
		AreaSetRayPickable(area gd.RID, enable bool)
		AreaSetMonitorCallback(area gd.RID, callback gd.Callable)
		AreaSetAreaMonitorCallback(area gd.RID, callback gd.Callable)
		BodyCreate() gd.RID
		BodySetSpace(body gd.RID, space gd.RID)
		BodyGetSpace(body gd.RID) gd.RID
		BodySetMode(body gd.RID, mode classdb.PhysicsServer3DBodyMode)
		BodyGetMode(body gd.RID) classdb.PhysicsServer3DBodyMode
		BodyAddShape(body gd.RID, shape gd.RID, transform gd.Transform3D, disabled bool)
		BodySetShape(body gd.RID, shape_idx int, shape gd.RID)
		BodySetShapeTransform(body gd.RID, shape_idx int, transform gd.Transform3D)
		BodySetShapeDisabled(body gd.RID, shape_idx int, disabled bool)
		BodyGetShapeCount(body gd.RID) int
		BodyGetShape(body gd.RID, shape_idx int) gd.RID
		BodyGetShapeTransform(body gd.RID, shape_idx int) gd.Transform3D
		BodyRemoveShape(body gd.RID, shape_idx int)
		BodyClearShapes(body gd.RID)
		BodyAttachObjectInstanceId(body gd.RID, id int)
		BodyGetObjectInstanceId(body gd.RID) int
		BodySetEnableContinuousCollisionDetection(body gd.RID, enable bool)
		BodyIsContinuousCollisionDetectionEnabled(body gd.RID) bool
		BodySetCollisionLayer(body gd.RID, layer int)
		BodyGetCollisionLayer(body gd.RID) int
		BodySetCollisionMask(body gd.RID, mask int)
		BodyGetCollisionMask(body gd.RID) int
		BodySetCollisionPriority(body gd.RID, priority float64)
		BodyGetCollisionPriority(body gd.RID) float64
		BodySetUserFlags(body gd.RID, flags int)
		BodyGetUserFlags(body gd.RID) int
		BodySetParam(body gd.RID, param classdb.PhysicsServer3DBodyParameter, value gd.Variant)
		BodyGetParam(body gd.RID, param classdb.PhysicsServer3DBodyParameter) gd.Variant
		BodyResetMassProperties(body gd.RID)
		BodySetState(body gd.RID, state classdb.PhysicsServer3DBodyState, value gd.Variant)
		BodyGetState(body gd.RID, state classdb.PhysicsServer3DBodyState) gd.Variant
		BodyApplyCentralImpulse(body gd.RID, impulse gd.Vector3)
		BodyApplyImpulse(body gd.RID, impulse gd.Vector3, position gd.Vector3)
		BodyApplyTorqueImpulse(body gd.RID, impulse gd.Vector3)
		BodyApplyCentralForce(body gd.RID, force gd.Vector3)
		BodyApplyForce(body gd.RID, force gd.Vector3, position gd.Vector3)
		BodyApplyTorque(body gd.RID, torque gd.Vector3)
		BodyAddConstantCentralForce(body gd.RID, force gd.Vector3)
		BodyAddConstantForce(body gd.RID, force gd.Vector3, position gd.Vector3)
		BodyAddConstantTorque(body gd.RID, torque gd.Vector3)
		BodySetConstantForce(body gd.RID, force gd.Vector3)
		BodyGetConstantForce(body gd.RID) gd.Vector3
		BodySetConstantTorque(body gd.RID, torque gd.Vector3)
		BodyGetConstantTorque(body gd.RID) gd.Vector3
		BodySetAxisVelocity(body gd.RID, axis_velocity gd.Vector3)
		BodySetAxisLock(body gd.RID, axis classdb.PhysicsServer3DBodyAxis, lock bool)
		BodyIsAxisLocked(body gd.RID, axis classdb.PhysicsServer3DBodyAxis) bool
		BodyAddCollisionException(body gd.RID, excepted_body gd.RID)
		BodyRemoveCollisionException(body gd.RID, excepted_body gd.RID)
		BodyGetCollisionExceptions(body gd.RID) gd.Array
		BodySetMaxContactsReported(body gd.RID, amount int)
		BodyGetMaxContactsReported(body gd.RID) int
		BodySetContactsReportedDepthThreshold(body gd.RID, threshold float64)
		BodyGetContactsReportedDepthThreshold(body gd.RID) float64
		BodySetOmitForceIntegration(body gd.RID, enable bool)
		BodyIsOmittingForceIntegration(body gd.RID) bool
		BodySetStateSyncCallback(body gd.RID, callable gd.Callable)
		BodySetForceIntegrationCallback(body gd.RID, callable gd.Callable, userdata gd.Variant)
		BodySetRayPickable(body gd.RID, enable bool)
		BodyTestMotion(body gd.RID, from gd.Transform3D, motion gd.Vector3, margin float64, max_collisions int, collide_separation_ray bool, recovery_as_collision bool, result *classdb.PhysicsServer3DExtensionMotionResult) bool
		BodyGetDirectState(body gd.RID) gdclass.PhysicsDirectBodyState3D
		SoftBodyCreate() gd.RID
		SoftBodyUpdateRenderingServer(body gd.RID, rendering_server_handler gdclass.PhysicsServer3DRenderingServerHandler)
		SoftBodySetSpace(body gd.RID, space gd.RID)
		SoftBodyGetSpace(body gd.RID) gd.RID
		SoftBodySetRayPickable(body gd.RID, enable bool)
		SoftBodySetCollisionLayer(body gd.RID, layer int)
		SoftBodyGetCollisionLayer(body gd.RID) int
		SoftBodySetCollisionMask(body gd.RID, mask int)
		SoftBodyGetCollisionMask(body gd.RID) int
		SoftBodyAddCollisionException(body gd.RID, body_b gd.RID)
		SoftBodyRemoveCollisionException(body gd.RID, body_b gd.RID)
		SoftBodyGetCollisionExceptions(body gd.RID) gd.Array
		SoftBodySetState(body gd.RID, state classdb.PhysicsServer3DBodyState, variant gd.Variant)
		SoftBodyGetState(body gd.RID, state classdb.PhysicsServer3DBodyState) gd.Variant
		SoftBodySetTransform(body gd.RID, transform gd.Transform3D)
		SoftBodySetSimulationPrecision(body gd.RID, simulation_precision int)
		SoftBodyGetSimulationPrecision(body gd.RID) int
		SoftBodySetTotalMass(body gd.RID, total_mass float64)
		SoftBodyGetTotalMass(body gd.RID) float64
		SoftBodySetLinearStiffness(body gd.RID, linear_stiffness float64)
		SoftBodyGetLinearStiffness(body gd.RID) float64
		SoftBodySetPressureCoefficient(body gd.RID, pressure_coefficient float64)
		SoftBodyGetPressureCoefficient(body gd.RID) float64
		SoftBodySetDampingCoefficient(body gd.RID, damping_coefficient float64)
		SoftBodyGetDampingCoefficient(body gd.RID) float64
		SoftBodySetDragCoefficient(body gd.RID, drag_coefficient float64)
		SoftBodyGetDragCoefficient(body gd.RID) float64
		SoftBodySetMesh(body gd.RID, mesh gd.RID)
		SoftBodyGetBounds(body gd.RID) gd.AABB
		SoftBodyMovePoint(body gd.RID, point_index int, global_position gd.Vector3)
		SoftBodyGetPointGlobalPosition(body gd.RID, point_index int) gd.Vector3
		SoftBodyRemoveAllPinnedPoints(body gd.RID)
		SoftBodyPinPoint(body gd.RID, point_index int, pin bool)
		SoftBodyIsPointPinned(body gd.RID, point_index int) bool
		JointCreate() gd.RID
		JointClear(joint gd.RID)
		JointMakePin(joint gd.RID, body_A gd.RID, local_A gd.Vector3, body_B gd.RID, local_B gd.Vector3)
		PinJointSetParam(joint gd.RID, param classdb.PhysicsServer3DPinJointParam, value float64)
		PinJointGetParam(joint gd.RID, param classdb.PhysicsServer3DPinJointParam) float64
		PinJointSetLocalA(joint gd.RID, local_A gd.Vector3)
		PinJointGetLocalA(joint gd.RID) gd.Vector3
		PinJointSetLocalB(joint gd.RID, local_B gd.Vector3)
		PinJointGetLocalB(joint gd.RID) gd.Vector3
		JointMakeHinge(joint gd.RID, body_A gd.RID, hinge_A gd.Transform3D, body_B gd.RID, hinge_B gd.Transform3D)
		JointMakeHingeSimple(joint gd.RID, body_A gd.RID, pivot_A gd.Vector3, axis_A gd.Vector3, body_B gd.RID, pivot_B gd.Vector3, axis_B gd.Vector3)
		HingeJointSetParam(joint gd.RID, param classdb.PhysicsServer3DHingeJointParam, value float64)
		HingeJointGetParam(joint gd.RID, param classdb.PhysicsServer3DHingeJointParam) float64
		HingeJointSetFlag(joint gd.RID, flag classdb.PhysicsServer3DHingeJointFlag, enabled bool)
		HingeJointGetFlag(joint gd.RID, flag classdb.PhysicsServer3DHingeJointFlag) bool
		JointMakeSlider(joint gd.RID, body_A gd.RID, local_ref_A gd.Transform3D, body_B gd.RID, local_ref_B gd.Transform3D)
		SliderJointSetParam(joint gd.RID, param classdb.PhysicsServer3DSliderJointParam, value float64)
		SliderJointGetParam(joint gd.RID, param classdb.PhysicsServer3DSliderJointParam) float64
		JointMakeConeTwist(joint gd.RID, body_A gd.RID, local_ref_A gd.Transform3D, body_B gd.RID, local_ref_B gd.Transform3D)
		ConeTwistJointSetParam(joint gd.RID, param classdb.PhysicsServer3DConeTwistJointParam, value float64)
		ConeTwistJointGetParam(joint gd.RID, param classdb.PhysicsServer3DConeTwistJointParam) float64
		JointMakeGeneric6dof(joint gd.RID, body_A gd.RID, local_ref_A gd.Transform3D, body_B gd.RID, local_ref_B gd.Transform3D)
		Generic6dofJointSetParam(joint gd.RID, axis gd.Vector3Axis, param classdb.PhysicsServer3DG6DOFJointAxisParam, value float64)
		Generic6dofJointGetParam(joint gd.RID, axis gd.Vector3Axis, param classdb.PhysicsServer3DG6DOFJointAxisParam) float64
		Generic6dofJointSetFlag(joint gd.RID, axis gd.Vector3Axis, flag classdb.PhysicsServer3DG6DOFJointAxisFlag, enable bool)
		Generic6dofJointGetFlag(joint gd.RID, axis gd.Vector3Axis, flag classdb.PhysicsServer3DG6DOFJointAxisFlag) bool
		JointGetType(joint gd.RID) classdb.PhysicsServer3DJointType
		JointSetSolverPriority(joint gd.RID, priority int)
		JointGetSolverPriority(joint gd.RID) int
		JointDisableCollisionsBetweenBodies(joint gd.RID, disable bool)
		JointIsDisabledCollisionsBetweenBodies(joint gd.RID) bool
		FreeRid(rid gd.RID)
		SetActive(active bool)
		Init()
		Step(step float64)
		Sync()
		FlushQueries()
		EndSync()
		Finish()
		IsFlushingQueries() bool
		GetProcessInfo(process_info classdb.PhysicsServer3DProcessInfo) int
	}
*/
type Instance [1]classdb.PhysicsServer3DExtension

func (Instance) _world_boundary_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _separation_ray_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _sphere_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _box_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _capsule_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _cylinder_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _convex_polygon_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _concave_polygon_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _heightmap_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _custom_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _shape_set_data(impl func(ptr unsafe.Pointer, shape gd.RID, data gd.Variant)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)
		var data = pointers.New[gd.Variant](gd.UnsafeGet[[3]uintptr](p_args, 1))
		defer pointers.End(data)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shape, data)
	}
}
func (Instance) _shape_set_custom_solver_bias(impl func(ptr unsafe.Pointer, shape gd.RID, bias float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)
		var bias = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shape, float64(bias))
	}
}
func (Instance) _shape_set_margin(impl func(ptr unsafe.Pointer, shape gd.RID, margin float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)
		var margin = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shape, float64(margin))
	}
}
func (Instance) _shape_get_margin(impl func(ptr unsafe.Pointer, shape gd.RID) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _shape_get_type(impl func(ptr unsafe.Pointer, shape gd.RID) classdb.PhysicsServer3DShapeType) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _shape_get_data(impl func(ptr unsafe.Pointer, shape gd.RID) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
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
func (Instance) _shape_get_custom_solver_bias(impl func(ptr unsafe.Pointer, shape gd.RID) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _space_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _space_set_active(impl func(ptr unsafe.Pointer, space gd.RID, active bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)
		var active = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, space, active)
	}
}
func (Instance) _space_is_active(impl func(ptr unsafe.Pointer, space gd.RID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _space_set_param(impl func(ptr unsafe.Pointer, space gd.RID, param classdb.PhysicsServer3DSpaceParameter, value float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DSpaceParameter](p_args, 1)
		var value = gd.UnsafeGet[gd.Float](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, space, param, float64(value))
	}
}
func (Instance) _space_get_param(impl func(ptr unsafe.Pointer, space gd.RID, param classdb.PhysicsServer3DSpaceParameter) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DSpaceParameter](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space, param)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _space_get_direct_state(impl func(ptr unsafe.Pointer, space gd.RID) gdclass.PhysicsDirectSpaceState3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
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
func (Instance) _space_set_debug_contacts(impl func(ptr unsafe.Pointer, space gd.RID, max_contacts int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)
		var max_contacts = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, space, int(max_contacts))
	}
}
func (Instance) _space_get_contacts(impl func(ptr unsafe.Pointer, space gd.RID) []gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		ptr, ok := pointers.End(gd.NewPackedVector3Slice(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _space_get_contact_count(impl func(ptr unsafe.Pointer, space gd.RID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _area_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _area_set_space(impl func(ptr unsafe.Pointer, area gd.RID, space gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var space = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, space)
	}
}
func (Instance) _area_get_space(impl func(ptr unsafe.Pointer, area gd.RID) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _area_add_shape(impl func(ptr unsafe.Pointer, area gd.RID, shape gd.RID, transform gd.Transform3D, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape = gd.UnsafeGet[gd.RID](p_args, 1)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		var disabled = gd.UnsafeGet[bool](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, shape, transform, disabled)
	}
}
func (Instance) _area_set_shape(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx int, shape gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		var shape = gd.UnsafeGet[gd.RID](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, int(shape_idx), shape)
	}
}
func (Instance) _area_set_shape_transform(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx int, transform gd.Transform3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, int(shape_idx), transform)
	}
}
func (Instance) _area_set_shape_disabled(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx int, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		var disabled = gd.UnsafeGet[bool](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, int(shape_idx), disabled)
	}
}
func (Instance) _area_get_shape_count(impl func(ptr unsafe.Pointer, area gd.RID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _area_get_shape(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx int) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, int(shape_idx))
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _area_get_shape_transform(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx int) gd.Transform3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, int(shape_idx))
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _area_remove_shape(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, int(shape_idx))
	}
}
func (Instance) _area_clear_shapes(impl func(ptr unsafe.Pointer, area gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area)
	}
}
func (Instance) _area_attach_object_instance_id(impl func(ptr unsafe.Pointer, area gd.RID, id int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var id = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, int(id))
	}
}
func (Instance) _area_get_object_instance_id(impl func(ptr unsafe.Pointer, area gd.RID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _area_set_param(impl func(ptr unsafe.Pointer, area gd.RID, param classdb.PhysicsServer3DAreaParameter, value gd.Variant)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DAreaParameter](p_args, 1)
		var value = pointers.New[gd.Variant](gd.UnsafeGet[[3]uintptr](p_args, 2))
		defer pointers.End(value)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, param, value)
	}
}
func (Instance) _area_set_transform(impl func(ptr unsafe.Pointer, area gd.RID, transform gd.Transform3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, transform)
	}
}
func (Instance) _area_get_param(impl func(ptr unsafe.Pointer, area gd.RID, param classdb.PhysicsServer3DAreaParameter) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DAreaParameter](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, param)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _area_get_transform(impl func(ptr unsafe.Pointer, area gd.RID) gd.Transform3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _area_set_collision_layer(impl func(ptr unsafe.Pointer, area gd.RID, layer int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var layer = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, int(layer))
	}
}
func (Instance) _area_get_collision_layer(impl func(ptr unsafe.Pointer, area gd.RID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _area_set_collision_mask(impl func(ptr unsafe.Pointer, area gd.RID, mask int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var mask = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, int(mask))
	}
}
func (Instance) _area_get_collision_mask(impl func(ptr unsafe.Pointer, area gd.RID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _area_set_monitorable(impl func(ptr unsafe.Pointer, area gd.RID, monitorable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var monitorable = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, monitorable)
	}
}
func (Instance) _area_set_ray_pickable(impl func(ptr unsafe.Pointer, area gd.RID, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var enable = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, enable)
	}
}
func (Instance) _area_set_monitor_callback(impl func(ptr unsafe.Pointer, area gd.RID, callback gd.Callable)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var callback = pointers.New[gd.Callable](gd.UnsafeGet[[2]uintptr](p_args, 1))
		defer pointers.End(callback)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, callback)
	}
}
func (Instance) _area_set_area_monitor_callback(impl func(ptr unsafe.Pointer, area gd.RID, callback gd.Callable)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var callback = pointers.New[gd.Callable](gd.UnsafeGet[[2]uintptr](p_args, 1))
		defer pointers.End(callback)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, callback)
	}
}
func (Instance) _body_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _body_set_space(impl func(ptr unsafe.Pointer, body gd.RID, space gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var space = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, space)
	}
}
func (Instance) _body_get_space(impl func(ptr unsafe.Pointer, body gd.RID) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _body_set_mode(impl func(ptr unsafe.Pointer, body gd.RID, mode classdb.PhysicsServer3DBodyMode)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var mode = gd.UnsafeGet[classdb.PhysicsServer3DBodyMode](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, mode)
	}
}
func (Instance) _body_get_mode(impl func(ptr unsafe.Pointer, body gd.RID) classdb.PhysicsServer3DBodyMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _body_add_shape(impl func(ptr unsafe.Pointer, body gd.RID, shape gd.RID, transform gd.Transform3D, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape = gd.UnsafeGet[gd.RID](p_args, 1)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		var disabled = gd.UnsafeGet[bool](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, shape, transform, disabled)
	}
}
func (Instance) _body_set_shape(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx int, shape gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		var shape = gd.UnsafeGet[gd.RID](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(shape_idx), shape)
	}
}
func (Instance) _body_set_shape_transform(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx int, transform gd.Transform3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(shape_idx), transform)
	}
}
func (Instance) _body_set_shape_disabled(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx int, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		var disabled = gd.UnsafeGet[bool](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(shape_idx), disabled)
	}
}
func (Instance) _body_get_shape_count(impl func(ptr unsafe.Pointer, body gd.RID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _body_get_shape(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx int) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, int(shape_idx))
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _body_get_shape_transform(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx int) gd.Transform3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, int(shape_idx))
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _body_remove_shape(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(shape_idx))
	}
}
func (Instance) _body_clear_shapes(impl func(ptr unsafe.Pointer, body gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body)
	}
}
func (Instance) _body_attach_object_instance_id(impl func(ptr unsafe.Pointer, body gd.RID, id int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var id = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(id))
	}
}
func (Instance) _body_get_object_instance_id(impl func(ptr unsafe.Pointer, body gd.RID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _body_set_enable_continuous_collision_detection(impl func(ptr unsafe.Pointer, body gd.RID, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var enable = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, enable)
	}
}
func (Instance) _body_is_continuous_collision_detection_enabled(impl func(ptr unsafe.Pointer, body gd.RID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _body_set_collision_layer(impl func(ptr unsafe.Pointer, body gd.RID, layer int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var layer = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(layer))
	}
}
func (Instance) _body_get_collision_layer(impl func(ptr unsafe.Pointer, body gd.RID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _body_set_collision_mask(impl func(ptr unsafe.Pointer, body gd.RID, mask int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var mask = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(mask))
	}
}
func (Instance) _body_get_collision_mask(impl func(ptr unsafe.Pointer, body gd.RID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _body_set_collision_priority(impl func(ptr unsafe.Pointer, body gd.RID, priority float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var priority = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, float64(priority))
	}
}
func (Instance) _body_get_collision_priority(impl func(ptr unsafe.Pointer, body gd.RID) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _body_set_user_flags(impl func(ptr unsafe.Pointer, body gd.RID, flags int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var flags = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(flags))
	}
}
func (Instance) _body_get_user_flags(impl func(ptr unsafe.Pointer, body gd.RID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _body_set_param(impl func(ptr unsafe.Pointer, body gd.RID, param classdb.PhysicsServer3DBodyParameter, value gd.Variant)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DBodyParameter](p_args, 1)
		var value = pointers.New[gd.Variant](gd.UnsafeGet[[3]uintptr](p_args, 2))
		defer pointers.End(value)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, param, value)
	}
}
func (Instance) _body_get_param(impl func(ptr unsafe.Pointer, body gd.RID, param classdb.PhysicsServer3DBodyParameter) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DBodyParameter](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, param)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _body_reset_mass_properties(impl func(ptr unsafe.Pointer, body gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body)
	}
}
func (Instance) _body_set_state(impl func(ptr unsafe.Pointer, body gd.RID, state classdb.PhysicsServer3DBodyState, value gd.Variant)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var state = gd.UnsafeGet[classdb.PhysicsServer3DBodyState](p_args, 1)
		var value = pointers.New[gd.Variant](gd.UnsafeGet[[3]uintptr](p_args, 2))
		defer pointers.End(value)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, state, value)
	}
}
func (Instance) _body_get_state(impl func(ptr unsafe.Pointer, body gd.RID, state classdb.PhysicsServer3DBodyState) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var state = gd.UnsafeGet[classdb.PhysicsServer3DBodyState](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, state)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _body_apply_central_impulse(impl func(ptr unsafe.Pointer, body gd.RID, impulse gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var impulse = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, impulse)
	}
}
func (Instance) _body_apply_impulse(impl func(ptr unsafe.Pointer, body gd.RID, impulse gd.Vector3, position gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var impulse = gd.UnsafeGet[gd.Vector3](p_args, 1)
		var position = gd.UnsafeGet[gd.Vector3](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, impulse, position)
	}
}
func (Instance) _body_apply_torque_impulse(impl func(ptr unsafe.Pointer, body gd.RID, impulse gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var impulse = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, impulse)
	}
}
func (Instance) _body_apply_central_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var force = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force)
	}
}
func (Instance) _body_apply_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector3, position gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var force = gd.UnsafeGet[gd.Vector3](p_args, 1)
		var position = gd.UnsafeGet[gd.Vector3](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force, position)
	}
}
func (Instance) _body_apply_torque(impl func(ptr unsafe.Pointer, body gd.RID, torque gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var torque = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, torque)
	}
}
func (Instance) _body_add_constant_central_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var force = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force)
	}
}
func (Instance) _body_add_constant_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector3, position gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var force = gd.UnsafeGet[gd.Vector3](p_args, 1)
		var position = gd.UnsafeGet[gd.Vector3](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force, position)
	}
}
func (Instance) _body_add_constant_torque(impl func(ptr unsafe.Pointer, body gd.RID, torque gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var torque = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, torque)
	}
}
func (Instance) _body_set_constant_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var force = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force)
	}
}
func (Instance) _body_get_constant_force(impl func(ptr unsafe.Pointer, body gd.RID) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _body_set_constant_torque(impl func(ptr unsafe.Pointer, body gd.RID, torque gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var torque = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, torque)
	}
}
func (Instance) _body_get_constant_torque(impl func(ptr unsafe.Pointer, body gd.RID) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _body_set_axis_velocity(impl func(ptr unsafe.Pointer, body gd.RID, axis_velocity gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var axis_velocity = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, axis_velocity)
	}
}
func (Instance) _body_set_axis_lock(impl func(ptr unsafe.Pointer, body gd.RID, axis classdb.PhysicsServer3DBodyAxis, lock bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var axis = gd.UnsafeGet[classdb.PhysicsServer3DBodyAxis](p_args, 1)
		var lock = gd.UnsafeGet[bool](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, axis, lock)
	}
}
func (Instance) _body_is_axis_locked(impl func(ptr unsafe.Pointer, body gd.RID, axis classdb.PhysicsServer3DBodyAxis) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var axis = gd.UnsafeGet[classdb.PhysicsServer3DBodyAxis](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, axis)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _body_add_collision_exception(impl func(ptr unsafe.Pointer, body gd.RID, excepted_body gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var excepted_body = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, excepted_body)
	}
}
func (Instance) _body_remove_collision_exception(impl func(ptr unsafe.Pointer, body gd.RID, excepted_body gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var excepted_body = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, excepted_body)
	}
}
func (Instance) _body_get_collision_exceptions(impl func(ptr unsafe.Pointer, body gd.RID) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _body_set_max_contacts_reported(impl func(ptr unsafe.Pointer, body gd.RID, amount int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var amount = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(amount))
	}
}
func (Instance) _body_get_max_contacts_reported(impl func(ptr unsafe.Pointer, body gd.RID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _body_set_contacts_reported_depth_threshold(impl func(ptr unsafe.Pointer, body gd.RID, threshold float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var threshold = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, float64(threshold))
	}
}
func (Instance) _body_get_contacts_reported_depth_threshold(impl func(ptr unsafe.Pointer, body gd.RID) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _body_set_omit_force_integration(impl func(ptr unsafe.Pointer, body gd.RID, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var enable = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, enable)
	}
}
func (Instance) _body_is_omitting_force_integration(impl func(ptr unsafe.Pointer, body gd.RID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _body_set_state_sync_callback(impl func(ptr unsafe.Pointer, body gd.RID, callable gd.Callable)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var callable = pointers.New[gd.Callable](gd.UnsafeGet[[2]uintptr](p_args, 1))
		defer pointers.End(callable)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, callable)
	}
}
func (Instance) _body_set_force_integration_callback(impl func(ptr unsafe.Pointer, body gd.RID, callable gd.Callable, userdata gd.Variant)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var callable = pointers.New[gd.Callable](gd.UnsafeGet[[2]uintptr](p_args, 1))
		defer pointers.End(callable)
		var userdata = pointers.New[gd.Variant](gd.UnsafeGet[[3]uintptr](p_args, 2))
		defer pointers.End(userdata)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, callable, userdata)
	}
}
func (Instance) _body_set_ray_pickable(impl func(ptr unsafe.Pointer, body gd.RID, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var enable = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, enable)
	}
}
func (Instance) _body_test_motion(impl func(ptr unsafe.Pointer, body gd.RID, from gd.Transform3D, motion gd.Vector3, margin float64, max_collisions int, collide_separation_ray bool, recovery_as_collision bool, result *classdb.PhysicsServer3DExtensionMotionResult) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var from = gd.UnsafeGet[gd.Transform3D](p_args, 1)
		var motion = gd.UnsafeGet[gd.Vector3](p_args, 2)
		var margin = gd.UnsafeGet[gd.Float](p_args, 3)
		var max_collisions = gd.UnsafeGet[gd.Int](p_args, 4)
		var collide_separation_ray = gd.UnsafeGet[bool](p_args, 5)
		var recovery_as_collision = gd.UnsafeGet[bool](p_args, 6)
		var result = gd.UnsafeGet[*classdb.PhysicsServer3DExtensionMotionResult](p_args, 7)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, from, motion, float64(margin), int(max_collisions), collide_separation_ray, recovery_as_collision, result)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _body_get_direct_state(impl func(ptr unsafe.Pointer, body gd.RID) gdclass.PhysicsDirectBodyState3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
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
func (Instance) _soft_body_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _soft_body_update_rendering_server(impl func(ptr unsafe.Pointer, body gd.RID, rendering_server_handler gdclass.PhysicsServer3DRenderingServerHandler)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var rendering_server_handler = gdclass.PhysicsServer3DRenderingServerHandler{pointers.New[classdb.PhysicsServer3DRenderingServerHandler]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
		defer pointers.End(rendering_server_handler[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, rendering_server_handler)
	}
}
func (Instance) _soft_body_set_space(impl func(ptr unsafe.Pointer, body gd.RID, space gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var space = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, space)
	}
}
func (Instance) _soft_body_get_space(impl func(ptr unsafe.Pointer, body gd.RID) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _soft_body_set_ray_pickable(impl func(ptr unsafe.Pointer, body gd.RID, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var enable = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, enable)
	}
}
func (Instance) _soft_body_set_collision_layer(impl func(ptr unsafe.Pointer, body gd.RID, layer int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var layer = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(layer))
	}
}
func (Instance) _soft_body_get_collision_layer(impl func(ptr unsafe.Pointer, body gd.RID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _soft_body_set_collision_mask(impl func(ptr unsafe.Pointer, body gd.RID, mask int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var mask = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(mask))
	}
}
func (Instance) _soft_body_get_collision_mask(impl func(ptr unsafe.Pointer, body gd.RID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _soft_body_add_collision_exception(impl func(ptr unsafe.Pointer, body gd.RID, body_b gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_b = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, body_b)
	}
}
func (Instance) _soft_body_remove_collision_exception(impl func(ptr unsafe.Pointer, body gd.RID, body_b gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_b = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, body_b)
	}
}
func (Instance) _soft_body_get_collision_exceptions(impl func(ptr unsafe.Pointer, body gd.RID) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _soft_body_set_state(impl func(ptr unsafe.Pointer, body gd.RID, state classdb.PhysicsServer3DBodyState, variant gd.Variant)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var state = gd.UnsafeGet[classdb.PhysicsServer3DBodyState](p_args, 1)
		var variant = pointers.New[gd.Variant](gd.UnsafeGet[[3]uintptr](p_args, 2))
		defer pointers.End(variant)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, state, variant)
	}
}
func (Instance) _soft_body_get_state(impl func(ptr unsafe.Pointer, body gd.RID, state classdb.PhysicsServer3DBodyState) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var state = gd.UnsafeGet[classdb.PhysicsServer3DBodyState](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, state)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _soft_body_set_transform(impl func(ptr unsafe.Pointer, body gd.RID, transform gd.Transform3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, transform)
	}
}
func (Instance) _soft_body_set_simulation_precision(impl func(ptr unsafe.Pointer, body gd.RID, simulation_precision int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var simulation_precision = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(simulation_precision))
	}
}
func (Instance) _soft_body_get_simulation_precision(impl func(ptr unsafe.Pointer, body gd.RID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _soft_body_set_total_mass(impl func(ptr unsafe.Pointer, body gd.RID, total_mass float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var total_mass = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, float64(total_mass))
	}
}
func (Instance) _soft_body_get_total_mass(impl func(ptr unsafe.Pointer, body gd.RID) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _soft_body_set_linear_stiffness(impl func(ptr unsafe.Pointer, body gd.RID, linear_stiffness float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var linear_stiffness = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, float64(linear_stiffness))
	}
}
func (Instance) _soft_body_get_linear_stiffness(impl func(ptr unsafe.Pointer, body gd.RID) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _soft_body_set_pressure_coefficient(impl func(ptr unsafe.Pointer, body gd.RID, pressure_coefficient float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var pressure_coefficient = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, float64(pressure_coefficient))
	}
}
func (Instance) _soft_body_get_pressure_coefficient(impl func(ptr unsafe.Pointer, body gd.RID) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _soft_body_set_damping_coefficient(impl func(ptr unsafe.Pointer, body gd.RID, damping_coefficient float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var damping_coefficient = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, float64(damping_coefficient))
	}
}
func (Instance) _soft_body_get_damping_coefficient(impl func(ptr unsafe.Pointer, body gd.RID) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _soft_body_set_drag_coefficient(impl func(ptr unsafe.Pointer, body gd.RID, drag_coefficient float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var drag_coefficient = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, float64(drag_coefficient))
	}
}
func (Instance) _soft_body_get_drag_coefficient(impl func(ptr unsafe.Pointer, body gd.RID) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _soft_body_set_mesh(impl func(ptr unsafe.Pointer, body gd.RID, mesh gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var mesh = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, mesh)
	}
}
func (Instance) _soft_body_get_bounds(impl func(ptr unsafe.Pointer, body gd.RID) gd.AABB) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _soft_body_move_point(impl func(ptr unsafe.Pointer, body gd.RID, point_index int, global_position gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var point_index = gd.UnsafeGet[gd.Int](p_args, 1)
		var global_position = gd.UnsafeGet[gd.Vector3](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(point_index), global_position)
	}
}
func (Instance) _soft_body_get_point_global_position(impl func(ptr unsafe.Pointer, body gd.RID, point_index int) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var point_index = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, int(point_index))
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _soft_body_remove_all_pinned_points(impl func(ptr unsafe.Pointer, body gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body)
	}
}
func (Instance) _soft_body_pin_point(impl func(ptr unsafe.Pointer, body gd.RID, point_index int, pin bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var point_index = gd.UnsafeGet[gd.Int](p_args, 1)
		var pin = gd.UnsafeGet[bool](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(point_index), pin)
	}
}
func (Instance) _soft_body_is_point_pinned(impl func(ptr unsafe.Pointer, body gd.RID, point_index int) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var point_index = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, int(point_index))
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _joint_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _joint_clear(impl func(ptr unsafe.Pointer, joint gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint)
	}
}
func (Instance) _joint_make_pin(impl func(ptr unsafe.Pointer, joint gd.RID, body_A gd.RID, local_A gd.Vector3, body_B gd.RID, local_B gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_A = gd.UnsafeGet[gd.RID](p_args, 1)
		var local_A = gd.UnsafeGet[gd.Vector3](p_args, 2)
		var body_B = gd.UnsafeGet[gd.RID](p_args, 3)
		var local_B = gd.UnsafeGet[gd.Vector3](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, local_A, body_B, local_B)
	}
}
func (Instance) _pin_joint_set_param(impl func(ptr unsafe.Pointer, joint gd.RID, param classdb.PhysicsServer3DPinJointParam, value float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DPinJointParam](p_args, 1)
		var value = gd.UnsafeGet[gd.Float](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, float64(value))
	}
}
func (Instance) _pin_joint_get_param(impl func(ptr unsafe.Pointer, joint gd.RID, param classdb.PhysicsServer3DPinJointParam) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DPinJointParam](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _pin_joint_set_local_a(impl func(ptr unsafe.Pointer, joint gd.RID, local_A gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var local_A = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, local_A)
	}
}
func (Instance) _pin_joint_get_local_a(impl func(ptr unsafe.Pointer, joint gd.RID) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _pin_joint_set_local_b(impl func(ptr unsafe.Pointer, joint gd.RID, local_B gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var local_B = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, local_B)
	}
}
func (Instance) _pin_joint_get_local_b(impl func(ptr unsafe.Pointer, joint gd.RID) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _joint_make_hinge(impl func(ptr unsafe.Pointer, joint gd.RID, body_A gd.RID, hinge_A gd.Transform3D, body_B gd.RID, hinge_B gd.Transform3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_A = gd.UnsafeGet[gd.RID](p_args, 1)
		var hinge_A = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		var body_B = gd.UnsafeGet[gd.RID](p_args, 3)
		var hinge_B = gd.UnsafeGet[gd.Transform3D](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, hinge_A, body_B, hinge_B)
	}
}
func (Instance) _joint_make_hinge_simple(impl func(ptr unsafe.Pointer, joint gd.RID, body_A gd.RID, pivot_A gd.Vector3, axis_A gd.Vector3, body_B gd.RID, pivot_B gd.Vector3, axis_B gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_A = gd.UnsafeGet[gd.RID](p_args, 1)
		var pivot_A = gd.UnsafeGet[gd.Vector3](p_args, 2)
		var axis_A = gd.UnsafeGet[gd.Vector3](p_args, 3)
		var body_B = gd.UnsafeGet[gd.RID](p_args, 4)
		var pivot_B = gd.UnsafeGet[gd.Vector3](p_args, 5)
		var axis_B = gd.UnsafeGet[gd.Vector3](p_args, 6)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, pivot_A, axis_A, body_B, pivot_B, axis_B)
	}
}
func (Instance) _hinge_joint_set_param(impl func(ptr unsafe.Pointer, joint gd.RID, param classdb.PhysicsServer3DHingeJointParam, value float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DHingeJointParam](p_args, 1)
		var value = gd.UnsafeGet[gd.Float](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, float64(value))
	}
}
func (Instance) _hinge_joint_get_param(impl func(ptr unsafe.Pointer, joint gd.RID, param classdb.PhysicsServer3DHingeJointParam) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DHingeJointParam](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _hinge_joint_set_flag(impl func(ptr unsafe.Pointer, joint gd.RID, flag classdb.PhysicsServer3DHingeJointFlag, enabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var flag = gd.UnsafeGet[classdb.PhysicsServer3DHingeJointFlag](p_args, 1)
		var enabled = gd.UnsafeGet[bool](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, flag, enabled)
	}
}
func (Instance) _hinge_joint_get_flag(impl func(ptr unsafe.Pointer, joint gd.RID, flag classdb.PhysicsServer3DHingeJointFlag) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var flag = gd.UnsafeGet[classdb.PhysicsServer3DHingeJointFlag](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, flag)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _joint_make_slider(impl func(ptr unsafe.Pointer, joint gd.RID, body_A gd.RID, local_ref_A gd.Transform3D, body_B gd.RID, local_ref_B gd.Transform3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_A = gd.UnsafeGet[gd.RID](p_args, 1)
		var local_ref_A = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		var body_B = gd.UnsafeGet[gd.RID](p_args, 3)
		var local_ref_B = gd.UnsafeGet[gd.Transform3D](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, local_ref_A, body_B, local_ref_B)
	}
}
func (Instance) _slider_joint_set_param(impl func(ptr unsafe.Pointer, joint gd.RID, param classdb.PhysicsServer3DSliderJointParam, value float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DSliderJointParam](p_args, 1)
		var value = gd.UnsafeGet[gd.Float](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, float64(value))
	}
}
func (Instance) _slider_joint_get_param(impl func(ptr unsafe.Pointer, joint gd.RID, param classdb.PhysicsServer3DSliderJointParam) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DSliderJointParam](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _joint_make_cone_twist(impl func(ptr unsafe.Pointer, joint gd.RID, body_A gd.RID, local_ref_A gd.Transform3D, body_B gd.RID, local_ref_B gd.Transform3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_A = gd.UnsafeGet[gd.RID](p_args, 1)
		var local_ref_A = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		var body_B = gd.UnsafeGet[gd.RID](p_args, 3)
		var local_ref_B = gd.UnsafeGet[gd.Transform3D](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, local_ref_A, body_B, local_ref_B)
	}
}
func (Instance) _cone_twist_joint_set_param(impl func(ptr unsafe.Pointer, joint gd.RID, param classdb.PhysicsServer3DConeTwistJointParam, value float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DConeTwistJointParam](p_args, 1)
		var value = gd.UnsafeGet[gd.Float](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, float64(value))
	}
}
func (Instance) _cone_twist_joint_get_param(impl func(ptr unsafe.Pointer, joint gd.RID, param classdb.PhysicsServer3DConeTwistJointParam) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DConeTwistJointParam](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _joint_make_generic_6dof(impl func(ptr unsafe.Pointer, joint gd.RID, body_A gd.RID, local_ref_A gd.Transform3D, body_B gd.RID, local_ref_B gd.Transform3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_A = gd.UnsafeGet[gd.RID](p_args, 1)
		var local_ref_A = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		var body_B = gd.UnsafeGet[gd.RID](p_args, 3)
		var local_ref_B = gd.UnsafeGet[gd.Transform3D](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, local_ref_A, body_B, local_ref_B)
	}
}
func (Instance) _generic_6dof_joint_set_param(impl func(ptr unsafe.Pointer, joint gd.RID, axis gd.Vector3Axis, param classdb.PhysicsServer3DG6DOFJointAxisParam, value float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var axis = gd.UnsafeGet[gd.Vector3Axis](p_args, 1)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DG6DOFJointAxisParam](p_args, 2)
		var value = gd.UnsafeGet[gd.Float](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, axis, param, float64(value))
	}
}
func (Instance) _generic_6dof_joint_get_param(impl func(ptr unsafe.Pointer, joint gd.RID, axis gd.Vector3Axis, param classdb.PhysicsServer3DG6DOFJointAxisParam) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var axis = gd.UnsafeGet[gd.Vector3Axis](p_args, 1)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DG6DOFJointAxisParam](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, axis, param)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}
func (Instance) _generic_6dof_joint_set_flag(impl func(ptr unsafe.Pointer, joint gd.RID, axis gd.Vector3Axis, flag classdb.PhysicsServer3DG6DOFJointAxisFlag, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var axis = gd.UnsafeGet[gd.Vector3Axis](p_args, 1)
		var flag = gd.UnsafeGet[classdb.PhysicsServer3DG6DOFJointAxisFlag](p_args, 2)
		var enable = gd.UnsafeGet[bool](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, axis, flag, enable)
	}
}
func (Instance) _generic_6dof_joint_get_flag(impl func(ptr unsafe.Pointer, joint gd.RID, axis gd.Vector3Axis, flag classdb.PhysicsServer3DG6DOFJointAxisFlag) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var axis = gd.UnsafeGet[gd.Vector3Axis](p_args, 1)
		var flag = gd.UnsafeGet[classdb.PhysicsServer3DG6DOFJointAxisFlag](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, axis, flag)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _joint_get_type(impl func(ptr unsafe.Pointer, joint gd.RID) classdb.PhysicsServer3DJointType) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _joint_set_solver_priority(impl func(ptr unsafe.Pointer, joint gd.RID, priority int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var priority = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, int(priority))
	}
}
func (Instance) _joint_get_solver_priority(impl func(ptr unsafe.Pointer, joint gd.RID) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _joint_disable_collisions_between_bodies(impl func(ptr unsafe.Pointer, joint gd.RID, disable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var disable = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, disable)
	}
}
func (Instance) _joint_is_disabled_collisions_between_bodies(impl func(ptr unsafe.Pointer, joint gd.RID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _free_rid(impl func(ptr unsafe.Pointer, rid gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var rid = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, rid)
	}
}
func (Instance) _set_active(impl func(ptr unsafe.Pointer, active bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var active = gd.UnsafeGet[bool](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, active)
	}
}
func (Instance) _init(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}
func (Instance) _step(impl func(ptr unsafe.Pointer, step float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var step = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, float64(step))
	}
}
func (Instance) _sync(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}
func (Instance) _flush_queries(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}
func (Instance) _end_sync(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}
func (Instance) _finish(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}
func (Instance) _is_flushing_queries(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_process_info(impl func(ptr unsafe.Pointer, process_info classdb.PhysicsServer3DProcessInfo) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var process_info = gd.UnsafeGet[classdb.PhysicsServer3DProcessInfo](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, process_info)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (self Instance) BodyTestMotionIsExcludingBody(body gd.RID) bool {
	return bool(class(self).BodyTestMotionIsExcludingBody(body))
}
func (self Instance) BodyTestMotionIsExcludingObject(obj int) bool {
	return bool(class(self).BodyTestMotionIsExcludingObject(gd.Int(obj)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.PhysicsServer3DExtension

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicsServer3DExtension"))
	return Instance{classdb.PhysicsServer3DExtension(object)}
}

func (class) _world_boundary_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _separation_ray_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _sphere_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _box_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _capsule_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _cylinder_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _convex_polygon_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _concave_polygon_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _heightmap_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _custom_shape_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _shape_set_data(impl func(ptr unsafe.Pointer, shape gd.RID, data gd.Variant)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)
		var data = pointers.New[gd.Variant](gd.UnsafeGet[[3]uintptr](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shape, data)
	}
}

func (class) _shape_set_custom_solver_bias(impl func(ptr unsafe.Pointer, shape gd.RID, bias gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)
		var bias = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shape, bias)
	}
}

func (class) _shape_set_margin(impl func(ptr unsafe.Pointer, shape gd.RID, margin gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)
		var margin = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shape, margin)
	}
}

func (class) _shape_get_margin(impl func(ptr unsafe.Pointer, shape gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _shape_get_type(impl func(ptr unsafe.Pointer, shape gd.RID) classdb.PhysicsServer3DShapeType) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _shape_get_data(impl func(ptr unsafe.Pointer, shape gd.RID) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
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

func (class) _shape_get_custom_solver_bias(impl func(ptr unsafe.Pointer, shape gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var shape = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _space_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _space_set_active(impl func(ptr unsafe.Pointer, space gd.RID, active bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)
		var active = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, space, active)
	}
}

func (class) _space_is_active(impl func(ptr unsafe.Pointer, space gd.RID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _space_set_param(impl func(ptr unsafe.Pointer, space gd.RID, param classdb.PhysicsServer3DSpaceParameter, value gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DSpaceParameter](p_args, 1)
		var value = gd.UnsafeGet[gd.Float](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, space, param, value)
	}
}

func (class) _space_get_param(impl func(ptr unsafe.Pointer, space gd.RID, param classdb.PhysicsServer3DSpaceParameter) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DSpaceParameter](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space, param)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _space_get_direct_state(impl func(ptr unsafe.Pointer, space gd.RID) gdclass.PhysicsDirectSpaceState3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
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

func (class) _space_set_debug_contacts(impl func(ptr unsafe.Pointer, space gd.RID, max_contacts gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)
		var max_contacts = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, space, max_contacts)
	}
}

func (class) _space_get_contacts(impl func(ptr unsafe.Pointer, space gd.RID) gd.PackedVector3Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
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

func (class) _space_get_contact_count(impl func(ptr unsafe.Pointer, space gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var space = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_set_space(impl func(ptr unsafe.Pointer, area gd.RID, space gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var space = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, space)
	}
}

func (class) _area_get_space(impl func(ptr unsafe.Pointer, area gd.RID) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_add_shape(impl func(ptr unsafe.Pointer, area gd.RID, shape gd.RID, transform gd.Transform3D, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape = gd.UnsafeGet[gd.RID](p_args, 1)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		var disabled = gd.UnsafeGet[bool](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, shape, transform, disabled)
	}
}

func (class) _area_set_shape(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx gd.Int, shape gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		var shape = gd.UnsafeGet[gd.RID](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, shape_idx, shape)
	}
}

func (class) _area_set_shape_transform(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx gd.Int, transform gd.Transform3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, shape_idx, transform)
	}
}

func (class) _area_set_shape_disabled(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx gd.Int, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		var disabled = gd.UnsafeGet[bool](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, shape_idx, disabled)
	}
}

func (class) _area_get_shape_count(impl func(ptr unsafe.Pointer, area gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_get_shape(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx gd.Int) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, shape_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_get_shape_transform(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx gd.Int) gd.Transform3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, shape_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_remove_shape(impl func(ptr unsafe.Pointer, area gd.RID, shape_idx gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, shape_idx)
	}
}

func (class) _area_clear_shapes(impl func(ptr unsafe.Pointer, area gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area)
	}
}

func (class) _area_attach_object_instance_id(impl func(ptr unsafe.Pointer, area gd.RID, id gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var id = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, id)
	}
}

func (class) _area_get_object_instance_id(impl func(ptr unsafe.Pointer, area gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_set_param(impl func(ptr unsafe.Pointer, area gd.RID, param classdb.PhysicsServer3DAreaParameter, value gd.Variant)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DAreaParameter](p_args, 1)
		var value = pointers.New[gd.Variant](gd.UnsafeGet[[3]uintptr](p_args, 2))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, param, value)
	}
}

func (class) _area_set_transform(impl func(ptr unsafe.Pointer, area gd.RID, transform gd.Transform3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, transform)
	}
}

func (class) _area_get_param(impl func(ptr unsafe.Pointer, area gd.RID, param classdb.PhysicsServer3DAreaParameter) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DAreaParameter](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, param)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _area_get_transform(impl func(ptr unsafe.Pointer, area gd.RID) gd.Transform3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_set_collision_layer(impl func(ptr unsafe.Pointer, area gd.RID, layer gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var layer = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, layer)
	}
}

func (class) _area_get_collision_layer(impl func(ptr unsafe.Pointer, area gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_set_collision_mask(impl func(ptr unsafe.Pointer, area gd.RID, mask gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var mask = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, mask)
	}
}

func (class) _area_get_collision_mask(impl func(ptr unsafe.Pointer, area gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_set_monitorable(impl func(ptr unsafe.Pointer, area gd.RID, monitorable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var monitorable = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, monitorable)
	}
}

func (class) _area_set_ray_pickable(impl func(ptr unsafe.Pointer, area gd.RID, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var enable = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, enable)
	}
}

func (class) _area_set_monitor_callback(impl func(ptr unsafe.Pointer, area gd.RID, callback gd.Callable)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var callback = pointers.New[gd.Callable](gd.UnsafeGet[[2]uintptr](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, callback)
	}
}

func (class) _area_set_area_monitor_callback(impl func(ptr unsafe.Pointer, area gd.RID, callback gd.Callable)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var area = gd.UnsafeGet[gd.RID](p_args, 0)
		var callback = pointers.New[gd.Callable](gd.UnsafeGet[[2]uintptr](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, callback)
	}
}

func (class) _body_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_space(impl func(ptr unsafe.Pointer, body gd.RID, space gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var space = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, space)
	}
}

func (class) _body_get_space(impl func(ptr unsafe.Pointer, body gd.RID) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_mode(impl func(ptr unsafe.Pointer, body gd.RID, mode classdb.PhysicsServer3DBodyMode)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var mode = gd.UnsafeGet[classdb.PhysicsServer3DBodyMode](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, mode)
	}
}

func (class) _body_get_mode(impl func(ptr unsafe.Pointer, body gd.RID) classdb.PhysicsServer3DBodyMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_add_shape(impl func(ptr unsafe.Pointer, body gd.RID, shape gd.RID, transform gd.Transform3D, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape = gd.UnsafeGet[gd.RID](p_args, 1)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		var disabled = gd.UnsafeGet[bool](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, shape, transform, disabled)
	}
}

func (class) _body_set_shape(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx gd.Int, shape gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		var shape = gd.UnsafeGet[gd.RID](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, shape_idx, shape)
	}
}

func (class) _body_set_shape_transform(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx gd.Int, transform gd.Transform3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, shape_idx, transform)
	}
}

func (class) _body_set_shape_disabled(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx gd.Int, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		var disabled = gd.UnsafeGet[bool](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, shape_idx, disabled)
	}
}

func (class) _body_get_shape_count(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_get_shape(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx gd.Int) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, shape_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_get_shape_transform(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx gd.Int) gd.Transform3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, shape_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_remove_shape(impl func(ptr unsafe.Pointer, body gd.RID, shape_idx gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var shape_idx = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, shape_idx)
	}
}

func (class) _body_clear_shapes(impl func(ptr unsafe.Pointer, body gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body)
	}
}

func (class) _body_attach_object_instance_id(impl func(ptr unsafe.Pointer, body gd.RID, id gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var id = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, id)
	}
}

func (class) _body_get_object_instance_id(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_enable_continuous_collision_detection(impl func(ptr unsafe.Pointer, body gd.RID, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var enable = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, enable)
	}
}

func (class) _body_is_continuous_collision_detection_enabled(impl func(ptr unsafe.Pointer, body gd.RID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_collision_layer(impl func(ptr unsafe.Pointer, body gd.RID, layer gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var layer = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, layer)
	}
}

func (class) _body_get_collision_layer(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_collision_mask(impl func(ptr unsafe.Pointer, body gd.RID, mask gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var mask = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, mask)
	}
}

func (class) _body_get_collision_mask(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_collision_priority(impl func(ptr unsafe.Pointer, body gd.RID, priority gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var priority = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, priority)
	}
}

func (class) _body_get_collision_priority(impl func(ptr unsafe.Pointer, body gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_user_flags(impl func(ptr unsafe.Pointer, body gd.RID, flags gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var flags = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, flags)
	}
}

func (class) _body_get_user_flags(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_param(impl func(ptr unsafe.Pointer, body gd.RID, param classdb.PhysicsServer3DBodyParameter, value gd.Variant)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DBodyParameter](p_args, 1)
		var value = pointers.New[gd.Variant](gd.UnsafeGet[[3]uintptr](p_args, 2))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, param, value)
	}
}

func (class) _body_get_param(impl func(ptr unsafe.Pointer, body gd.RID, param classdb.PhysicsServer3DBodyParameter) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DBodyParameter](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, param)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _body_reset_mass_properties(impl func(ptr unsafe.Pointer, body gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body)
	}
}

func (class) _body_set_state(impl func(ptr unsafe.Pointer, body gd.RID, state classdb.PhysicsServer3DBodyState, value gd.Variant)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var state = gd.UnsafeGet[classdb.PhysicsServer3DBodyState](p_args, 1)
		var value = pointers.New[gd.Variant](gd.UnsafeGet[[3]uintptr](p_args, 2))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, state, value)
	}
}

func (class) _body_get_state(impl func(ptr unsafe.Pointer, body gd.RID, state classdb.PhysicsServer3DBodyState) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var state = gd.UnsafeGet[classdb.PhysicsServer3DBodyState](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, state)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _body_apply_central_impulse(impl func(ptr unsafe.Pointer, body gd.RID, impulse gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var impulse = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, impulse)
	}
}

func (class) _body_apply_impulse(impl func(ptr unsafe.Pointer, body gd.RID, impulse gd.Vector3, position gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var impulse = gd.UnsafeGet[gd.Vector3](p_args, 1)
		var position = gd.UnsafeGet[gd.Vector3](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, impulse, position)
	}
}

func (class) _body_apply_torque_impulse(impl func(ptr unsafe.Pointer, body gd.RID, impulse gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var impulse = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, impulse)
	}
}

func (class) _body_apply_central_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var force = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force)
	}
}

func (class) _body_apply_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector3, position gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var force = gd.UnsafeGet[gd.Vector3](p_args, 1)
		var position = gd.UnsafeGet[gd.Vector3](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force, position)
	}
}

func (class) _body_apply_torque(impl func(ptr unsafe.Pointer, body gd.RID, torque gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var torque = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, torque)
	}
}

func (class) _body_add_constant_central_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var force = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force)
	}
}

func (class) _body_add_constant_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector3, position gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var force = gd.UnsafeGet[gd.Vector3](p_args, 1)
		var position = gd.UnsafeGet[gd.Vector3](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force, position)
	}
}

func (class) _body_add_constant_torque(impl func(ptr unsafe.Pointer, body gd.RID, torque gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var torque = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, torque)
	}
}

func (class) _body_set_constant_force(impl func(ptr unsafe.Pointer, body gd.RID, force gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var force = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force)
	}
}

func (class) _body_get_constant_force(impl func(ptr unsafe.Pointer, body gd.RID) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_constant_torque(impl func(ptr unsafe.Pointer, body gd.RID, torque gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var torque = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, torque)
	}
}

func (class) _body_get_constant_torque(impl func(ptr unsafe.Pointer, body gd.RID) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_axis_velocity(impl func(ptr unsafe.Pointer, body gd.RID, axis_velocity gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var axis_velocity = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, axis_velocity)
	}
}

func (class) _body_set_axis_lock(impl func(ptr unsafe.Pointer, body gd.RID, axis classdb.PhysicsServer3DBodyAxis, lock bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var axis = gd.UnsafeGet[classdb.PhysicsServer3DBodyAxis](p_args, 1)
		var lock = gd.UnsafeGet[bool](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, axis, lock)
	}
}

func (class) _body_is_axis_locked(impl func(ptr unsafe.Pointer, body gd.RID, axis classdb.PhysicsServer3DBodyAxis) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var axis = gd.UnsafeGet[classdb.PhysicsServer3DBodyAxis](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, axis)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_add_collision_exception(impl func(ptr unsafe.Pointer, body gd.RID, excepted_body gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var excepted_body = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, excepted_body)
	}
}

func (class) _body_remove_collision_exception(impl func(ptr unsafe.Pointer, body gd.RID, excepted_body gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var excepted_body = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, excepted_body)
	}
}

func (class) _body_get_collision_exceptions(impl func(ptr unsafe.Pointer, body gd.RID) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _body_set_max_contacts_reported(impl func(ptr unsafe.Pointer, body gd.RID, amount gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var amount = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, amount)
	}
}

func (class) _body_get_max_contacts_reported(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_contacts_reported_depth_threshold(impl func(ptr unsafe.Pointer, body gd.RID, threshold gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var threshold = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, threshold)
	}
}

func (class) _body_get_contacts_reported_depth_threshold(impl func(ptr unsafe.Pointer, body gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_omit_force_integration(impl func(ptr unsafe.Pointer, body gd.RID, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var enable = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, enable)
	}
}

func (class) _body_is_omitting_force_integration(impl func(ptr unsafe.Pointer, body gd.RID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_state_sync_callback(impl func(ptr unsafe.Pointer, body gd.RID, callable gd.Callable)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var callable = pointers.New[gd.Callable](gd.UnsafeGet[[2]uintptr](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, callable)
	}
}

func (class) _body_set_force_integration_callback(impl func(ptr unsafe.Pointer, body gd.RID, callable gd.Callable, userdata gd.Variant)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var callable = pointers.New[gd.Callable](gd.UnsafeGet[[2]uintptr](p_args, 1))
		var userdata = pointers.New[gd.Variant](gd.UnsafeGet[[3]uintptr](p_args, 2))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, callable, userdata)
	}
}

func (class) _body_set_ray_pickable(impl func(ptr unsafe.Pointer, body gd.RID, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var enable = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, enable)
	}
}

func (class) _body_test_motion(impl func(ptr unsafe.Pointer, body gd.RID, from gd.Transform3D, motion gd.Vector3, margin gd.Float, max_collisions gd.Int, collide_separation_ray bool, recovery_as_collision bool, result *classdb.PhysicsServer3DExtensionMotionResult) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var from = gd.UnsafeGet[gd.Transform3D](p_args, 1)
		var motion = gd.UnsafeGet[gd.Vector3](p_args, 2)
		var margin = gd.UnsafeGet[gd.Float](p_args, 3)
		var max_collisions = gd.UnsafeGet[gd.Int](p_args, 4)
		var collide_separation_ray = gd.UnsafeGet[bool](p_args, 5)
		var recovery_as_collision = gd.UnsafeGet[bool](p_args, 6)
		var result = gd.UnsafeGet[*classdb.PhysicsServer3DExtensionMotionResult](p_args, 7)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, from, motion, margin, max_collisions, collide_separation_ray, recovery_as_collision, result)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_get_direct_state(impl func(ptr unsafe.Pointer, body gd.RID) gdclass.PhysicsDirectBodyState3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
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

func (class) _soft_body_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_update_rendering_server(impl func(ptr unsafe.Pointer, body gd.RID, rendering_server_handler gdclass.PhysicsServer3DRenderingServerHandler)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var rendering_server_handler = gdclass.PhysicsServer3DRenderingServerHandler{pointers.New[classdb.PhysicsServer3DRenderingServerHandler]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
		defer pointers.End(rendering_server_handler[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, rendering_server_handler)
	}
}

func (class) _soft_body_set_space(impl func(ptr unsafe.Pointer, body gd.RID, space gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var space = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, space)
	}
}

func (class) _soft_body_get_space(impl func(ptr unsafe.Pointer, body gd.RID) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_set_ray_pickable(impl func(ptr unsafe.Pointer, body gd.RID, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var enable = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, enable)
	}
}

func (class) _soft_body_set_collision_layer(impl func(ptr unsafe.Pointer, body gd.RID, layer gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var layer = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, layer)
	}
}

func (class) _soft_body_get_collision_layer(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_set_collision_mask(impl func(ptr unsafe.Pointer, body gd.RID, mask gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var mask = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, mask)
	}
}

func (class) _soft_body_get_collision_mask(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_add_collision_exception(impl func(ptr unsafe.Pointer, body gd.RID, body_b gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_b = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, body_b)
	}
}

func (class) _soft_body_remove_collision_exception(impl func(ptr unsafe.Pointer, body gd.RID, body_b gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_b = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, body_b)
	}
}

func (class) _soft_body_get_collision_exceptions(impl func(ptr unsafe.Pointer, body gd.RID) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _soft_body_set_state(impl func(ptr unsafe.Pointer, body gd.RID, state classdb.PhysicsServer3DBodyState, variant gd.Variant)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var state = gd.UnsafeGet[classdb.PhysicsServer3DBodyState](p_args, 1)
		var variant = pointers.New[gd.Variant](gd.UnsafeGet[[3]uintptr](p_args, 2))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, state, variant)
	}
}

func (class) _soft_body_get_state(impl func(ptr unsafe.Pointer, body gd.RID, state classdb.PhysicsServer3DBodyState) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var state = gd.UnsafeGet[classdb.PhysicsServer3DBodyState](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, state)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _soft_body_set_transform(impl func(ptr unsafe.Pointer, body gd.RID, transform gd.Transform3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, transform)
	}
}

func (class) _soft_body_set_simulation_precision(impl func(ptr unsafe.Pointer, body gd.RID, simulation_precision gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var simulation_precision = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, simulation_precision)
	}
}

func (class) _soft_body_get_simulation_precision(impl func(ptr unsafe.Pointer, body gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_set_total_mass(impl func(ptr unsafe.Pointer, body gd.RID, total_mass gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var total_mass = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, total_mass)
	}
}

func (class) _soft_body_get_total_mass(impl func(ptr unsafe.Pointer, body gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_set_linear_stiffness(impl func(ptr unsafe.Pointer, body gd.RID, linear_stiffness gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var linear_stiffness = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, linear_stiffness)
	}
}

func (class) _soft_body_get_linear_stiffness(impl func(ptr unsafe.Pointer, body gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_set_pressure_coefficient(impl func(ptr unsafe.Pointer, body gd.RID, pressure_coefficient gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var pressure_coefficient = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, pressure_coefficient)
	}
}

func (class) _soft_body_get_pressure_coefficient(impl func(ptr unsafe.Pointer, body gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_set_damping_coefficient(impl func(ptr unsafe.Pointer, body gd.RID, damping_coefficient gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var damping_coefficient = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, damping_coefficient)
	}
}

func (class) _soft_body_get_damping_coefficient(impl func(ptr unsafe.Pointer, body gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_set_drag_coefficient(impl func(ptr unsafe.Pointer, body gd.RID, drag_coefficient gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var drag_coefficient = gd.UnsafeGet[gd.Float](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, drag_coefficient)
	}
}

func (class) _soft_body_get_drag_coefficient(impl func(ptr unsafe.Pointer, body gd.RID) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_set_mesh(impl func(ptr unsafe.Pointer, body gd.RID, mesh gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var mesh = gd.UnsafeGet[gd.RID](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, mesh)
	}
}

func (class) _soft_body_get_bounds(impl func(ptr unsafe.Pointer, body gd.RID) gd.AABB) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_move_point(impl func(ptr unsafe.Pointer, body gd.RID, point_index gd.Int, global_position gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var point_index = gd.UnsafeGet[gd.Int](p_args, 1)
		var global_position = gd.UnsafeGet[gd.Vector3](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, point_index, global_position)
	}
}

func (class) _soft_body_get_point_global_position(impl func(ptr unsafe.Pointer, body gd.RID, point_index gd.Int) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var point_index = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, point_index)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_remove_all_pinned_points(impl func(ptr unsafe.Pointer, body gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body)
	}
}

func (class) _soft_body_pin_point(impl func(ptr unsafe.Pointer, body gd.RID, point_index gd.Int, pin bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var point_index = gd.UnsafeGet[gd.Int](p_args, 1)
		var pin = gd.UnsafeGet[bool](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, point_index, pin)
	}
}

func (class) _soft_body_is_point_pinned(impl func(ptr unsafe.Pointer, body gd.RID, point_index gd.Int) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var body = gd.UnsafeGet[gd.RID](p_args, 0)
		var point_index = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, point_index)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _joint_create(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _joint_clear(impl func(ptr unsafe.Pointer, joint gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint)
	}
}

func (class) _joint_make_pin(impl func(ptr unsafe.Pointer, joint gd.RID, body_A gd.RID, local_A gd.Vector3, body_B gd.RID, local_B gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_A = gd.UnsafeGet[gd.RID](p_args, 1)
		var local_A = gd.UnsafeGet[gd.Vector3](p_args, 2)
		var body_B = gd.UnsafeGet[gd.RID](p_args, 3)
		var local_B = gd.UnsafeGet[gd.Vector3](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, local_A, body_B, local_B)
	}
}

func (class) _pin_joint_set_param(impl func(ptr unsafe.Pointer, joint gd.RID, param classdb.PhysicsServer3DPinJointParam, value gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DPinJointParam](p_args, 1)
		var value = gd.UnsafeGet[gd.Float](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, value)
	}
}

func (class) _pin_joint_get_param(impl func(ptr unsafe.Pointer, joint gd.RID, param classdb.PhysicsServer3DPinJointParam) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DPinJointParam](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _pin_joint_set_local_a(impl func(ptr unsafe.Pointer, joint gd.RID, local_A gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var local_A = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, local_A)
	}
}

func (class) _pin_joint_get_local_a(impl func(ptr unsafe.Pointer, joint gd.RID) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _pin_joint_set_local_b(impl func(ptr unsafe.Pointer, joint gd.RID, local_B gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var local_B = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, local_B)
	}
}

func (class) _pin_joint_get_local_b(impl func(ptr unsafe.Pointer, joint gd.RID) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _joint_make_hinge(impl func(ptr unsafe.Pointer, joint gd.RID, body_A gd.RID, hinge_A gd.Transform3D, body_B gd.RID, hinge_B gd.Transform3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_A = gd.UnsafeGet[gd.RID](p_args, 1)
		var hinge_A = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		var body_B = gd.UnsafeGet[gd.RID](p_args, 3)
		var hinge_B = gd.UnsafeGet[gd.Transform3D](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, hinge_A, body_B, hinge_B)
	}
}

func (class) _joint_make_hinge_simple(impl func(ptr unsafe.Pointer, joint gd.RID, body_A gd.RID, pivot_A gd.Vector3, axis_A gd.Vector3, body_B gd.RID, pivot_B gd.Vector3, axis_B gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_A = gd.UnsafeGet[gd.RID](p_args, 1)
		var pivot_A = gd.UnsafeGet[gd.Vector3](p_args, 2)
		var axis_A = gd.UnsafeGet[gd.Vector3](p_args, 3)
		var body_B = gd.UnsafeGet[gd.RID](p_args, 4)
		var pivot_B = gd.UnsafeGet[gd.Vector3](p_args, 5)
		var axis_B = gd.UnsafeGet[gd.Vector3](p_args, 6)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, pivot_A, axis_A, body_B, pivot_B, axis_B)
	}
}

func (class) _hinge_joint_set_param(impl func(ptr unsafe.Pointer, joint gd.RID, param classdb.PhysicsServer3DHingeJointParam, value gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DHingeJointParam](p_args, 1)
		var value = gd.UnsafeGet[gd.Float](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, value)
	}
}

func (class) _hinge_joint_get_param(impl func(ptr unsafe.Pointer, joint gd.RID, param classdb.PhysicsServer3DHingeJointParam) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DHingeJointParam](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _hinge_joint_set_flag(impl func(ptr unsafe.Pointer, joint gd.RID, flag classdb.PhysicsServer3DHingeJointFlag, enabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var flag = gd.UnsafeGet[classdb.PhysicsServer3DHingeJointFlag](p_args, 1)
		var enabled = gd.UnsafeGet[bool](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, flag, enabled)
	}
}

func (class) _hinge_joint_get_flag(impl func(ptr unsafe.Pointer, joint gd.RID, flag classdb.PhysicsServer3DHingeJointFlag) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var flag = gd.UnsafeGet[classdb.PhysicsServer3DHingeJointFlag](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, flag)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _joint_make_slider(impl func(ptr unsafe.Pointer, joint gd.RID, body_A gd.RID, local_ref_A gd.Transform3D, body_B gd.RID, local_ref_B gd.Transform3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_A = gd.UnsafeGet[gd.RID](p_args, 1)
		var local_ref_A = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		var body_B = gd.UnsafeGet[gd.RID](p_args, 3)
		var local_ref_B = gd.UnsafeGet[gd.Transform3D](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, local_ref_A, body_B, local_ref_B)
	}
}

func (class) _slider_joint_set_param(impl func(ptr unsafe.Pointer, joint gd.RID, param classdb.PhysicsServer3DSliderJointParam, value gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DSliderJointParam](p_args, 1)
		var value = gd.UnsafeGet[gd.Float](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, value)
	}
}

func (class) _slider_joint_get_param(impl func(ptr unsafe.Pointer, joint gd.RID, param classdb.PhysicsServer3DSliderJointParam) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DSliderJointParam](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _joint_make_cone_twist(impl func(ptr unsafe.Pointer, joint gd.RID, body_A gd.RID, local_ref_A gd.Transform3D, body_B gd.RID, local_ref_B gd.Transform3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_A = gd.UnsafeGet[gd.RID](p_args, 1)
		var local_ref_A = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		var body_B = gd.UnsafeGet[gd.RID](p_args, 3)
		var local_ref_B = gd.UnsafeGet[gd.Transform3D](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, local_ref_A, body_B, local_ref_B)
	}
}

func (class) _cone_twist_joint_set_param(impl func(ptr unsafe.Pointer, joint gd.RID, param classdb.PhysicsServer3DConeTwistJointParam, value gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DConeTwistJointParam](p_args, 1)
		var value = gd.UnsafeGet[gd.Float](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, value)
	}
}

func (class) _cone_twist_joint_get_param(impl func(ptr unsafe.Pointer, joint gd.RID, param classdb.PhysicsServer3DConeTwistJointParam) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DConeTwistJointParam](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _joint_make_generic_6dof(impl func(ptr unsafe.Pointer, joint gd.RID, body_A gd.RID, local_ref_A gd.Transform3D, body_B gd.RID, local_ref_B gd.Transform3D)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var body_A = gd.UnsafeGet[gd.RID](p_args, 1)
		var local_ref_A = gd.UnsafeGet[gd.Transform3D](p_args, 2)
		var body_B = gd.UnsafeGet[gd.RID](p_args, 3)
		var local_ref_B = gd.UnsafeGet[gd.Transform3D](p_args, 4)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, local_ref_A, body_B, local_ref_B)
	}
}

func (class) _generic_6dof_joint_set_param(impl func(ptr unsafe.Pointer, joint gd.RID, axis gd.Vector3Axis, param classdb.PhysicsServer3DG6DOFJointAxisParam, value gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var axis = gd.UnsafeGet[gd.Vector3Axis](p_args, 1)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DG6DOFJointAxisParam](p_args, 2)
		var value = gd.UnsafeGet[gd.Float](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, axis, param, value)
	}
}

func (class) _generic_6dof_joint_get_param(impl func(ptr unsafe.Pointer, joint gd.RID, axis gd.Vector3Axis, param classdb.PhysicsServer3DG6DOFJointAxisParam) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var axis = gd.UnsafeGet[gd.Vector3Axis](p_args, 1)
		var param = gd.UnsafeGet[classdb.PhysicsServer3DG6DOFJointAxisParam](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, axis, param)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _generic_6dof_joint_set_flag(impl func(ptr unsafe.Pointer, joint gd.RID, axis gd.Vector3Axis, flag classdb.PhysicsServer3DG6DOFJointAxisFlag, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var axis = gd.UnsafeGet[gd.Vector3Axis](p_args, 1)
		var flag = gd.UnsafeGet[classdb.PhysicsServer3DG6DOFJointAxisFlag](p_args, 2)
		var enable = gd.UnsafeGet[bool](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, axis, flag, enable)
	}
}

func (class) _generic_6dof_joint_get_flag(impl func(ptr unsafe.Pointer, joint gd.RID, axis gd.Vector3Axis, flag classdb.PhysicsServer3DG6DOFJointAxisFlag) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var axis = gd.UnsafeGet[gd.Vector3Axis](p_args, 1)
		var flag = gd.UnsafeGet[classdb.PhysicsServer3DG6DOFJointAxisFlag](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, axis, flag)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _joint_get_type(impl func(ptr unsafe.Pointer, joint gd.RID) classdb.PhysicsServer3DJointType) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _joint_set_solver_priority(impl func(ptr unsafe.Pointer, joint gd.RID, priority gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var priority = gd.UnsafeGet[gd.Int](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, priority)
	}
}

func (class) _joint_get_solver_priority(impl func(ptr unsafe.Pointer, joint gd.RID) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _joint_disable_collisions_between_bodies(impl func(ptr unsafe.Pointer, joint gd.RID, disable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		var disable = gd.UnsafeGet[bool](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, disable)
	}
}

func (class) _joint_is_disabled_collisions_between_bodies(impl func(ptr unsafe.Pointer, joint gd.RID) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var joint = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _free_rid(impl func(ptr unsafe.Pointer, rid gd.RID)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var rid = gd.UnsafeGet[gd.RID](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, rid)
	}
}

func (class) _set_active(impl func(ptr unsafe.Pointer, active bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var active = gd.UnsafeGet[bool](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, active)
	}
}

func (class) _init(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (class) _step(impl func(ptr unsafe.Pointer, step gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var step = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, step)
	}
}

func (class) _sync(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (class) _flush_queries(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (class) _end_sync(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (class) _finish(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (class) _is_flushing_queries(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_process_info(impl func(ptr unsafe.Pointer, process_info classdb.PhysicsServer3DProcessInfo) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var process_info = gd.UnsafeGet[classdb.PhysicsServer3DProcessInfo](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, process_info)
		gd.UnsafeSet(p_back, ret)
	}
}

//go:nosplit
func (self class) BodyTestMotionIsExcludingBody(body gd.RID) bool {
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsServer3DExtension.Bind_body_test_motion_is_excluding_body, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) BodyTestMotionIsExcludingObject(obj gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, obj)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsServer3DExtension.Bind_body_test_motion_is_excluding_object, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPhysicsServer3DExtension() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPhysicsServer3DExtension() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_world_boundary_shape_create":
		return reflect.ValueOf(self._world_boundary_shape_create)
	case "_separation_ray_shape_create":
		return reflect.ValueOf(self._separation_ray_shape_create)
	case "_sphere_shape_create":
		return reflect.ValueOf(self._sphere_shape_create)
	case "_box_shape_create":
		return reflect.ValueOf(self._box_shape_create)
	case "_capsule_shape_create":
		return reflect.ValueOf(self._capsule_shape_create)
	case "_cylinder_shape_create":
		return reflect.ValueOf(self._cylinder_shape_create)
	case "_convex_polygon_shape_create":
		return reflect.ValueOf(self._convex_polygon_shape_create)
	case "_concave_polygon_shape_create":
		return reflect.ValueOf(self._concave_polygon_shape_create)
	case "_heightmap_shape_create":
		return reflect.ValueOf(self._heightmap_shape_create)
	case "_custom_shape_create":
		return reflect.ValueOf(self._custom_shape_create)
	case "_shape_set_data":
		return reflect.ValueOf(self._shape_set_data)
	case "_shape_set_custom_solver_bias":
		return reflect.ValueOf(self._shape_set_custom_solver_bias)
	case "_shape_set_margin":
		return reflect.ValueOf(self._shape_set_margin)
	case "_shape_get_margin":
		return reflect.ValueOf(self._shape_get_margin)
	case "_shape_get_type":
		return reflect.ValueOf(self._shape_get_type)
	case "_shape_get_data":
		return reflect.ValueOf(self._shape_get_data)
	case "_shape_get_custom_solver_bias":
		return reflect.ValueOf(self._shape_get_custom_solver_bias)
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
	case "_area_set_ray_pickable":
		return reflect.ValueOf(self._area_set_ray_pickable)
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
	case "_body_set_shape_disabled":
		return reflect.ValueOf(self._body_set_shape_disabled)
	case "_body_get_shape_count":
		return reflect.ValueOf(self._body_get_shape_count)
	case "_body_get_shape":
		return reflect.ValueOf(self._body_get_shape)
	case "_body_get_shape_transform":
		return reflect.ValueOf(self._body_get_shape_transform)
	case "_body_remove_shape":
		return reflect.ValueOf(self._body_remove_shape)
	case "_body_clear_shapes":
		return reflect.ValueOf(self._body_clear_shapes)
	case "_body_attach_object_instance_id":
		return reflect.ValueOf(self._body_attach_object_instance_id)
	case "_body_get_object_instance_id":
		return reflect.ValueOf(self._body_get_object_instance_id)
	case "_body_set_enable_continuous_collision_detection":
		return reflect.ValueOf(self._body_set_enable_continuous_collision_detection)
	case "_body_is_continuous_collision_detection_enabled":
		return reflect.ValueOf(self._body_is_continuous_collision_detection_enabled)
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
	case "_body_set_user_flags":
		return reflect.ValueOf(self._body_set_user_flags)
	case "_body_get_user_flags":
		return reflect.ValueOf(self._body_get_user_flags)
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
	case "_body_apply_impulse":
		return reflect.ValueOf(self._body_apply_impulse)
	case "_body_apply_torque_impulse":
		return reflect.ValueOf(self._body_apply_torque_impulse)
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
	case "_body_set_axis_lock":
		return reflect.ValueOf(self._body_set_axis_lock)
	case "_body_is_axis_locked":
		return reflect.ValueOf(self._body_is_axis_locked)
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
	case "_body_set_ray_pickable":
		return reflect.ValueOf(self._body_set_ray_pickable)
	case "_body_test_motion":
		return reflect.ValueOf(self._body_test_motion)
	case "_body_get_direct_state":
		return reflect.ValueOf(self._body_get_direct_state)
	case "_soft_body_create":
		return reflect.ValueOf(self._soft_body_create)
	case "_soft_body_update_rendering_server":
		return reflect.ValueOf(self._soft_body_update_rendering_server)
	case "_soft_body_set_space":
		return reflect.ValueOf(self._soft_body_set_space)
	case "_soft_body_get_space":
		return reflect.ValueOf(self._soft_body_get_space)
	case "_soft_body_set_ray_pickable":
		return reflect.ValueOf(self._soft_body_set_ray_pickable)
	case "_soft_body_set_collision_layer":
		return reflect.ValueOf(self._soft_body_set_collision_layer)
	case "_soft_body_get_collision_layer":
		return reflect.ValueOf(self._soft_body_get_collision_layer)
	case "_soft_body_set_collision_mask":
		return reflect.ValueOf(self._soft_body_set_collision_mask)
	case "_soft_body_get_collision_mask":
		return reflect.ValueOf(self._soft_body_get_collision_mask)
	case "_soft_body_add_collision_exception":
		return reflect.ValueOf(self._soft_body_add_collision_exception)
	case "_soft_body_remove_collision_exception":
		return reflect.ValueOf(self._soft_body_remove_collision_exception)
	case "_soft_body_get_collision_exceptions":
		return reflect.ValueOf(self._soft_body_get_collision_exceptions)
	case "_soft_body_set_state":
		return reflect.ValueOf(self._soft_body_set_state)
	case "_soft_body_get_state":
		return reflect.ValueOf(self._soft_body_get_state)
	case "_soft_body_set_transform":
		return reflect.ValueOf(self._soft_body_set_transform)
	case "_soft_body_set_simulation_precision":
		return reflect.ValueOf(self._soft_body_set_simulation_precision)
	case "_soft_body_get_simulation_precision":
		return reflect.ValueOf(self._soft_body_get_simulation_precision)
	case "_soft_body_set_total_mass":
		return reflect.ValueOf(self._soft_body_set_total_mass)
	case "_soft_body_get_total_mass":
		return reflect.ValueOf(self._soft_body_get_total_mass)
	case "_soft_body_set_linear_stiffness":
		return reflect.ValueOf(self._soft_body_set_linear_stiffness)
	case "_soft_body_get_linear_stiffness":
		return reflect.ValueOf(self._soft_body_get_linear_stiffness)
	case "_soft_body_set_pressure_coefficient":
		return reflect.ValueOf(self._soft_body_set_pressure_coefficient)
	case "_soft_body_get_pressure_coefficient":
		return reflect.ValueOf(self._soft_body_get_pressure_coefficient)
	case "_soft_body_set_damping_coefficient":
		return reflect.ValueOf(self._soft_body_set_damping_coefficient)
	case "_soft_body_get_damping_coefficient":
		return reflect.ValueOf(self._soft_body_get_damping_coefficient)
	case "_soft_body_set_drag_coefficient":
		return reflect.ValueOf(self._soft_body_set_drag_coefficient)
	case "_soft_body_get_drag_coefficient":
		return reflect.ValueOf(self._soft_body_get_drag_coefficient)
	case "_soft_body_set_mesh":
		return reflect.ValueOf(self._soft_body_set_mesh)
	case "_soft_body_get_bounds":
		return reflect.ValueOf(self._soft_body_get_bounds)
	case "_soft_body_move_point":
		return reflect.ValueOf(self._soft_body_move_point)
	case "_soft_body_get_point_global_position":
		return reflect.ValueOf(self._soft_body_get_point_global_position)
	case "_soft_body_remove_all_pinned_points":
		return reflect.ValueOf(self._soft_body_remove_all_pinned_points)
	case "_soft_body_pin_point":
		return reflect.ValueOf(self._soft_body_pin_point)
	case "_soft_body_is_point_pinned":
		return reflect.ValueOf(self._soft_body_is_point_pinned)
	case "_joint_create":
		return reflect.ValueOf(self._joint_create)
	case "_joint_clear":
		return reflect.ValueOf(self._joint_clear)
	case "_joint_make_pin":
		return reflect.ValueOf(self._joint_make_pin)
	case "_pin_joint_set_param":
		return reflect.ValueOf(self._pin_joint_set_param)
	case "_pin_joint_get_param":
		return reflect.ValueOf(self._pin_joint_get_param)
	case "_pin_joint_set_local_a":
		return reflect.ValueOf(self._pin_joint_set_local_a)
	case "_pin_joint_get_local_a":
		return reflect.ValueOf(self._pin_joint_get_local_a)
	case "_pin_joint_set_local_b":
		return reflect.ValueOf(self._pin_joint_set_local_b)
	case "_pin_joint_get_local_b":
		return reflect.ValueOf(self._pin_joint_get_local_b)
	case "_joint_make_hinge":
		return reflect.ValueOf(self._joint_make_hinge)
	case "_joint_make_hinge_simple":
		return reflect.ValueOf(self._joint_make_hinge_simple)
	case "_hinge_joint_set_param":
		return reflect.ValueOf(self._hinge_joint_set_param)
	case "_hinge_joint_get_param":
		return reflect.ValueOf(self._hinge_joint_get_param)
	case "_hinge_joint_set_flag":
		return reflect.ValueOf(self._hinge_joint_set_flag)
	case "_hinge_joint_get_flag":
		return reflect.ValueOf(self._hinge_joint_get_flag)
	case "_joint_make_slider":
		return reflect.ValueOf(self._joint_make_slider)
	case "_slider_joint_set_param":
		return reflect.ValueOf(self._slider_joint_set_param)
	case "_slider_joint_get_param":
		return reflect.ValueOf(self._slider_joint_get_param)
	case "_joint_make_cone_twist":
		return reflect.ValueOf(self._joint_make_cone_twist)
	case "_cone_twist_joint_set_param":
		return reflect.ValueOf(self._cone_twist_joint_set_param)
	case "_cone_twist_joint_get_param":
		return reflect.ValueOf(self._cone_twist_joint_get_param)
	case "_joint_make_generic_6dof":
		return reflect.ValueOf(self._joint_make_generic_6dof)
	case "_generic_6dof_joint_set_param":
		return reflect.ValueOf(self._generic_6dof_joint_set_param)
	case "_generic_6dof_joint_get_param":
		return reflect.ValueOf(self._generic_6dof_joint_get_param)
	case "_generic_6dof_joint_set_flag":
		return reflect.ValueOf(self._generic_6dof_joint_set_flag)
	case "_generic_6dof_joint_get_flag":
		return reflect.ValueOf(self._generic_6dof_joint_get_flag)
	case "_joint_get_type":
		return reflect.ValueOf(self._joint_get_type)
	case "_joint_set_solver_priority":
		return reflect.ValueOf(self._joint_set_solver_priority)
	case "_joint_get_solver_priority":
		return reflect.ValueOf(self._joint_get_solver_priority)
	case "_joint_disable_collisions_between_bodies":
		return reflect.ValueOf(self._joint_disable_collisions_between_bodies)
	case "_joint_is_disabled_collisions_between_bodies":
		return reflect.ValueOf(self._joint_is_disabled_collisions_between_bodies)
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
	case "_sphere_shape_create":
		return reflect.ValueOf(self._sphere_shape_create)
	case "_box_shape_create":
		return reflect.ValueOf(self._box_shape_create)
	case "_capsule_shape_create":
		return reflect.ValueOf(self._capsule_shape_create)
	case "_cylinder_shape_create":
		return reflect.ValueOf(self._cylinder_shape_create)
	case "_convex_polygon_shape_create":
		return reflect.ValueOf(self._convex_polygon_shape_create)
	case "_concave_polygon_shape_create":
		return reflect.ValueOf(self._concave_polygon_shape_create)
	case "_heightmap_shape_create":
		return reflect.ValueOf(self._heightmap_shape_create)
	case "_custom_shape_create":
		return reflect.ValueOf(self._custom_shape_create)
	case "_shape_set_data":
		return reflect.ValueOf(self._shape_set_data)
	case "_shape_set_custom_solver_bias":
		return reflect.ValueOf(self._shape_set_custom_solver_bias)
	case "_shape_set_margin":
		return reflect.ValueOf(self._shape_set_margin)
	case "_shape_get_margin":
		return reflect.ValueOf(self._shape_get_margin)
	case "_shape_get_type":
		return reflect.ValueOf(self._shape_get_type)
	case "_shape_get_data":
		return reflect.ValueOf(self._shape_get_data)
	case "_shape_get_custom_solver_bias":
		return reflect.ValueOf(self._shape_get_custom_solver_bias)
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
	case "_area_set_ray_pickable":
		return reflect.ValueOf(self._area_set_ray_pickable)
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
	case "_body_set_shape_disabled":
		return reflect.ValueOf(self._body_set_shape_disabled)
	case "_body_get_shape_count":
		return reflect.ValueOf(self._body_get_shape_count)
	case "_body_get_shape":
		return reflect.ValueOf(self._body_get_shape)
	case "_body_get_shape_transform":
		return reflect.ValueOf(self._body_get_shape_transform)
	case "_body_remove_shape":
		return reflect.ValueOf(self._body_remove_shape)
	case "_body_clear_shapes":
		return reflect.ValueOf(self._body_clear_shapes)
	case "_body_attach_object_instance_id":
		return reflect.ValueOf(self._body_attach_object_instance_id)
	case "_body_get_object_instance_id":
		return reflect.ValueOf(self._body_get_object_instance_id)
	case "_body_set_enable_continuous_collision_detection":
		return reflect.ValueOf(self._body_set_enable_continuous_collision_detection)
	case "_body_is_continuous_collision_detection_enabled":
		return reflect.ValueOf(self._body_is_continuous_collision_detection_enabled)
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
	case "_body_set_user_flags":
		return reflect.ValueOf(self._body_set_user_flags)
	case "_body_get_user_flags":
		return reflect.ValueOf(self._body_get_user_flags)
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
	case "_body_apply_impulse":
		return reflect.ValueOf(self._body_apply_impulse)
	case "_body_apply_torque_impulse":
		return reflect.ValueOf(self._body_apply_torque_impulse)
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
	case "_body_set_axis_lock":
		return reflect.ValueOf(self._body_set_axis_lock)
	case "_body_is_axis_locked":
		return reflect.ValueOf(self._body_is_axis_locked)
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
	case "_body_set_ray_pickable":
		return reflect.ValueOf(self._body_set_ray_pickable)
	case "_body_test_motion":
		return reflect.ValueOf(self._body_test_motion)
	case "_body_get_direct_state":
		return reflect.ValueOf(self._body_get_direct_state)
	case "_soft_body_create":
		return reflect.ValueOf(self._soft_body_create)
	case "_soft_body_update_rendering_server":
		return reflect.ValueOf(self._soft_body_update_rendering_server)
	case "_soft_body_set_space":
		return reflect.ValueOf(self._soft_body_set_space)
	case "_soft_body_get_space":
		return reflect.ValueOf(self._soft_body_get_space)
	case "_soft_body_set_ray_pickable":
		return reflect.ValueOf(self._soft_body_set_ray_pickable)
	case "_soft_body_set_collision_layer":
		return reflect.ValueOf(self._soft_body_set_collision_layer)
	case "_soft_body_get_collision_layer":
		return reflect.ValueOf(self._soft_body_get_collision_layer)
	case "_soft_body_set_collision_mask":
		return reflect.ValueOf(self._soft_body_set_collision_mask)
	case "_soft_body_get_collision_mask":
		return reflect.ValueOf(self._soft_body_get_collision_mask)
	case "_soft_body_add_collision_exception":
		return reflect.ValueOf(self._soft_body_add_collision_exception)
	case "_soft_body_remove_collision_exception":
		return reflect.ValueOf(self._soft_body_remove_collision_exception)
	case "_soft_body_get_collision_exceptions":
		return reflect.ValueOf(self._soft_body_get_collision_exceptions)
	case "_soft_body_set_state":
		return reflect.ValueOf(self._soft_body_set_state)
	case "_soft_body_get_state":
		return reflect.ValueOf(self._soft_body_get_state)
	case "_soft_body_set_transform":
		return reflect.ValueOf(self._soft_body_set_transform)
	case "_soft_body_set_simulation_precision":
		return reflect.ValueOf(self._soft_body_set_simulation_precision)
	case "_soft_body_get_simulation_precision":
		return reflect.ValueOf(self._soft_body_get_simulation_precision)
	case "_soft_body_set_total_mass":
		return reflect.ValueOf(self._soft_body_set_total_mass)
	case "_soft_body_get_total_mass":
		return reflect.ValueOf(self._soft_body_get_total_mass)
	case "_soft_body_set_linear_stiffness":
		return reflect.ValueOf(self._soft_body_set_linear_stiffness)
	case "_soft_body_get_linear_stiffness":
		return reflect.ValueOf(self._soft_body_get_linear_stiffness)
	case "_soft_body_set_pressure_coefficient":
		return reflect.ValueOf(self._soft_body_set_pressure_coefficient)
	case "_soft_body_get_pressure_coefficient":
		return reflect.ValueOf(self._soft_body_get_pressure_coefficient)
	case "_soft_body_set_damping_coefficient":
		return reflect.ValueOf(self._soft_body_set_damping_coefficient)
	case "_soft_body_get_damping_coefficient":
		return reflect.ValueOf(self._soft_body_get_damping_coefficient)
	case "_soft_body_set_drag_coefficient":
		return reflect.ValueOf(self._soft_body_set_drag_coefficient)
	case "_soft_body_get_drag_coefficient":
		return reflect.ValueOf(self._soft_body_get_drag_coefficient)
	case "_soft_body_set_mesh":
		return reflect.ValueOf(self._soft_body_set_mesh)
	case "_soft_body_get_bounds":
		return reflect.ValueOf(self._soft_body_get_bounds)
	case "_soft_body_move_point":
		return reflect.ValueOf(self._soft_body_move_point)
	case "_soft_body_get_point_global_position":
		return reflect.ValueOf(self._soft_body_get_point_global_position)
	case "_soft_body_remove_all_pinned_points":
		return reflect.ValueOf(self._soft_body_remove_all_pinned_points)
	case "_soft_body_pin_point":
		return reflect.ValueOf(self._soft_body_pin_point)
	case "_soft_body_is_point_pinned":
		return reflect.ValueOf(self._soft_body_is_point_pinned)
	case "_joint_create":
		return reflect.ValueOf(self._joint_create)
	case "_joint_clear":
		return reflect.ValueOf(self._joint_clear)
	case "_joint_make_pin":
		return reflect.ValueOf(self._joint_make_pin)
	case "_pin_joint_set_param":
		return reflect.ValueOf(self._pin_joint_set_param)
	case "_pin_joint_get_param":
		return reflect.ValueOf(self._pin_joint_get_param)
	case "_pin_joint_set_local_a":
		return reflect.ValueOf(self._pin_joint_set_local_a)
	case "_pin_joint_get_local_a":
		return reflect.ValueOf(self._pin_joint_get_local_a)
	case "_pin_joint_set_local_b":
		return reflect.ValueOf(self._pin_joint_set_local_b)
	case "_pin_joint_get_local_b":
		return reflect.ValueOf(self._pin_joint_get_local_b)
	case "_joint_make_hinge":
		return reflect.ValueOf(self._joint_make_hinge)
	case "_joint_make_hinge_simple":
		return reflect.ValueOf(self._joint_make_hinge_simple)
	case "_hinge_joint_set_param":
		return reflect.ValueOf(self._hinge_joint_set_param)
	case "_hinge_joint_get_param":
		return reflect.ValueOf(self._hinge_joint_get_param)
	case "_hinge_joint_set_flag":
		return reflect.ValueOf(self._hinge_joint_set_flag)
	case "_hinge_joint_get_flag":
		return reflect.ValueOf(self._hinge_joint_get_flag)
	case "_joint_make_slider":
		return reflect.ValueOf(self._joint_make_slider)
	case "_slider_joint_set_param":
		return reflect.ValueOf(self._slider_joint_set_param)
	case "_slider_joint_get_param":
		return reflect.ValueOf(self._slider_joint_get_param)
	case "_joint_make_cone_twist":
		return reflect.ValueOf(self._joint_make_cone_twist)
	case "_cone_twist_joint_set_param":
		return reflect.ValueOf(self._cone_twist_joint_set_param)
	case "_cone_twist_joint_get_param":
		return reflect.ValueOf(self._cone_twist_joint_get_param)
	case "_joint_make_generic_6dof":
		return reflect.ValueOf(self._joint_make_generic_6dof)
	case "_generic_6dof_joint_set_param":
		return reflect.ValueOf(self._generic_6dof_joint_set_param)
	case "_generic_6dof_joint_get_param":
		return reflect.ValueOf(self._generic_6dof_joint_get_param)
	case "_generic_6dof_joint_set_flag":
		return reflect.ValueOf(self._generic_6dof_joint_set_flag)
	case "_generic_6dof_joint_get_flag":
		return reflect.ValueOf(self._generic_6dof_joint_get_flag)
	case "_joint_get_type":
		return reflect.ValueOf(self._joint_get_type)
	case "_joint_set_solver_priority":
		return reflect.ValueOf(self._joint_set_solver_priority)
	case "_joint_get_solver_priority":
		return reflect.ValueOf(self._joint_get_solver_priority)
	case "_joint_disable_collisions_between_bodies":
		return reflect.ValueOf(self._joint_disable_collisions_between_bodies)
	case "_joint_is_disabled_collisions_between_bodies":
		return reflect.ValueOf(self._joint_is_disabled_collisions_between_bodies)
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
	classdb.Register("PhysicsServer3DExtension", func(ptr gd.Object) any { return classdb.PhysicsServer3DExtension(ptr) })
}
