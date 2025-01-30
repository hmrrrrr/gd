// Package PhysicsServer3DExtension provides methods for working with PhysicsServer3DExtension object instances.
package PhysicsServer3DExtension

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/AABB"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Transform3D"
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
This class extends [PhysicsServer3D] by providing additional virtual methods that can be overridden. When these methods are overridden, they will be called instead of the internal methods of the physics server.
Intended for use with GDExtension to create custom implementations of [PhysicsServer3D].

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=PhysicsServer3DExtension)
*/
type Instance [1]gdclass.PhysicsServer3DExtension

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPhysicsServer3DExtension() Instance
}
type Interface interface {
	WorldBoundaryShapeCreate() RID.Any
	SeparationRayShapeCreate() RID.Any
	SphereShapeCreate() RID.Any
	BoxShapeCreate() RID.Any
	CapsuleShapeCreate() RID.Any
	CylinderShapeCreate() RID.Any
	ConvexPolygonShapeCreate() RID.Any
	ConcavePolygonShapeCreate() RID.Any
	HeightmapShapeCreate() RID.Any
	CustomShapeCreate() RID.Any
	ShapeSetData(shape RID.Any, data any)
	ShapeSetCustomSolverBias(shape RID.Any, bias Float.X)
	ShapeSetMargin(shape RID.Any, margin Float.X)
	ShapeGetMargin(shape RID.Any) Float.X
	ShapeGetType(shape RID.Any) gdclass.PhysicsServer3DShapeType
	ShapeGetData(shape RID.Any) any
	ShapeGetCustomSolverBias(shape RID.Any) Float.X
	SpaceCreate() RID.Any
	SpaceSetActive(space RID.Any, active bool)
	SpaceIsActive(space RID.Any) bool
	SpaceSetParam(space RID.Any, param gdclass.PhysicsServer3DSpaceParameter, value Float.X)
	SpaceGetParam(space RID.Any, param gdclass.PhysicsServer3DSpaceParameter) Float.X
	SpaceGetDirectState(space RID.Any) [1]gdclass.PhysicsDirectSpaceState3D
	SpaceSetDebugContacts(space RID.Any, max_contacts int)
	SpaceGetContacts(space RID.Any) []Vector3.XYZ
	SpaceGetContactCount(space RID.Any) int
	AreaCreate() RID.Any
	AreaSetSpace(area RID.Any, space RID.Any)
	AreaGetSpace(area RID.Any) RID.Any
	AreaAddShape(area RID.Any, shape RID.Any, transform Transform3D.BasisOrigin, disabled bool)
	AreaSetShape(area RID.Any, shape_idx int, shape RID.Any)
	AreaSetShapeTransform(area RID.Any, shape_idx int, transform Transform3D.BasisOrigin)
	AreaSetShapeDisabled(area RID.Any, shape_idx int, disabled bool)
	AreaGetShapeCount(area RID.Any) int
	AreaGetShape(area RID.Any, shape_idx int) RID.Any
	AreaGetShapeTransform(area RID.Any, shape_idx int) Transform3D.BasisOrigin
	AreaRemoveShape(area RID.Any, shape_idx int)
	AreaClearShapes(area RID.Any)
	AreaAttachObjectInstanceId(area RID.Any, id int)
	AreaGetObjectInstanceId(area RID.Any) int
	AreaSetParam(area RID.Any, param gdclass.PhysicsServer3DAreaParameter, value any)
	AreaSetTransform(area RID.Any, transform Transform3D.BasisOrigin)
	AreaGetParam(area RID.Any, param gdclass.PhysicsServer3DAreaParameter) any
	AreaGetTransform(area RID.Any) Transform3D.BasisOrigin
	AreaSetCollisionLayer(area RID.Any, layer int)
	AreaGetCollisionLayer(area RID.Any) int
	AreaSetCollisionMask(area RID.Any, mask int)
	AreaGetCollisionMask(area RID.Any) int
	AreaSetMonitorable(area RID.Any, monitorable bool)
	AreaSetRayPickable(area RID.Any, enable bool)
	AreaSetMonitorCallback(area RID.Any, callback Callable.Function)
	AreaSetAreaMonitorCallback(area RID.Any, callback Callable.Function)
	BodyCreate() RID.Any
	BodySetSpace(body RID.Any, space RID.Any)
	BodyGetSpace(body RID.Any) RID.Any
	BodySetMode(body RID.Any, mode gdclass.PhysicsServer3DBodyMode)
	BodyGetMode(body RID.Any) gdclass.PhysicsServer3DBodyMode
	BodyAddShape(body RID.Any, shape RID.Any, transform Transform3D.BasisOrigin, disabled bool)
	BodySetShape(body RID.Any, shape_idx int, shape RID.Any)
	BodySetShapeTransform(body RID.Any, shape_idx int, transform Transform3D.BasisOrigin)
	BodySetShapeDisabled(body RID.Any, shape_idx int, disabled bool)
	BodyGetShapeCount(body RID.Any) int
	BodyGetShape(body RID.Any, shape_idx int) RID.Any
	BodyGetShapeTransform(body RID.Any, shape_idx int) Transform3D.BasisOrigin
	BodyRemoveShape(body RID.Any, shape_idx int)
	BodyClearShapes(body RID.Any)
	BodyAttachObjectInstanceId(body RID.Any, id int)
	BodyGetObjectInstanceId(body RID.Any) int
	BodySetEnableContinuousCollisionDetection(body RID.Any, enable bool)
	BodyIsContinuousCollisionDetectionEnabled(body RID.Any) bool
	BodySetCollisionLayer(body RID.Any, layer int)
	BodyGetCollisionLayer(body RID.Any) int
	BodySetCollisionMask(body RID.Any, mask int)
	BodyGetCollisionMask(body RID.Any) int
	BodySetCollisionPriority(body RID.Any, priority Float.X)
	BodyGetCollisionPriority(body RID.Any) Float.X
	BodySetUserFlags(body RID.Any, flags int)
	BodyGetUserFlags(body RID.Any) int
	BodySetParam(body RID.Any, param gdclass.PhysicsServer3DBodyParameter, value any)
	BodyGetParam(body RID.Any, param gdclass.PhysicsServer3DBodyParameter) any
	BodyResetMassProperties(body RID.Any)
	BodySetState(body RID.Any, state gdclass.PhysicsServer3DBodyState, value any)
	BodyGetState(body RID.Any, state gdclass.PhysicsServer3DBodyState) any
	BodyApplyCentralImpulse(body RID.Any, impulse Vector3.XYZ)
	BodyApplyImpulse(body RID.Any, impulse Vector3.XYZ, position Vector3.XYZ)
	BodyApplyTorqueImpulse(body RID.Any, impulse Vector3.XYZ)
	BodyApplyCentralForce(body RID.Any, force Vector3.XYZ)
	BodyApplyForce(body RID.Any, force Vector3.XYZ, position Vector3.XYZ)
	BodyApplyTorque(body RID.Any, torque Vector3.XYZ)
	BodyAddConstantCentralForce(body RID.Any, force Vector3.XYZ)
	BodyAddConstantForce(body RID.Any, force Vector3.XYZ, position Vector3.XYZ)
	BodyAddConstantTorque(body RID.Any, torque Vector3.XYZ)
	BodySetConstantForce(body RID.Any, force Vector3.XYZ)
	BodyGetConstantForce(body RID.Any) Vector3.XYZ
	BodySetConstantTorque(body RID.Any, torque Vector3.XYZ)
	BodyGetConstantTorque(body RID.Any) Vector3.XYZ
	BodySetAxisVelocity(body RID.Any, axis_velocity Vector3.XYZ)
	BodySetAxisLock(body RID.Any, axis gdclass.PhysicsServer3DBodyAxis, lock bool)
	BodyIsAxisLocked(body RID.Any, axis gdclass.PhysicsServer3DBodyAxis) bool
	BodyAddCollisionException(body RID.Any, excepted_body RID.Any)
	BodyRemoveCollisionException(body RID.Any, excepted_body RID.Any)
	BodyGetCollisionExceptions(body RID.Any) []RID.Any
	BodySetMaxContactsReported(body RID.Any, amount int)
	BodyGetMaxContactsReported(body RID.Any) int
	BodySetContactsReportedDepthThreshold(body RID.Any, threshold Float.X)
	BodyGetContactsReportedDepthThreshold(body RID.Any) Float.X
	BodySetOmitForceIntegration(body RID.Any, enable bool)
	BodyIsOmittingForceIntegration(body RID.Any) bool
	BodySetStateSyncCallback(body RID.Any, callable Callable.Function)
	BodySetForceIntegrationCallback(body RID.Any, callable Callable.Function, userdata any)
	BodySetRayPickable(body RID.Any, enable bool)
	BodyTestMotion(body RID.Any, from Transform3D.BasisOrigin, motion Vector3.XYZ, margin Float.X, max_collisions int, collide_separation_ray bool, recovery_as_collision bool, result *MotionResult) bool
	BodyGetDirectState(body RID.Any) [1]gdclass.PhysicsDirectBodyState3D
	SoftBodyCreate() RID.Any
	SoftBodyUpdateRenderingServer(body RID.Any, rendering_server_handler [1]gdclass.PhysicsServer3DRenderingServerHandler)
	SoftBodySetSpace(body RID.Any, space RID.Any)
	SoftBodyGetSpace(body RID.Any) RID.Any
	SoftBodySetRayPickable(body RID.Any, enable bool)
	SoftBodySetCollisionLayer(body RID.Any, layer int)
	SoftBodyGetCollisionLayer(body RID.Any) int
	SoftBodySetCollisionMask(body RID.Any, mask int)
	SoftBodyGetCollisionMask(body RID.Any) int
	SoftBodyAddCollisionException(body RID.Any, body_b RID.Any)
	SoftBodyRemoveCollisionException(body RID.Any, body_b RID.Any)
	SoftBodyGetCollisionExceptions(body RID.Any) []RID.Any
	SoftBodySetState(body RID.Any, state gdclass.PhysicsServer3DBodyState, v any)
	SoftBodyGetState(body RID.Any, state gdclass.PhysicsServer3DBodyState) any
	SoftBodySetTransform(body RID.Any, transform Transform3D.BasisOrigin)
	SoftBodySetSimulationPrecision(body RID.Any, simulation_precision int)
	SoftBodyGetSimulationPrecision(body RID.Any) int
	SoftBodySetTotalMass(body RID.Any, total_mass Float.X)
	SoftBodyGetTotalMass(body RID.Any) Float.X
	SoftBodySetLinearStiffness(body RID.Any, linear_stiffness Float.X)
	SoftBodyGetLinearStiffness(body RID.Any) Float.X
	SoftBodySetPressureCoefficient(body RID.Any, pressure_coefficient Float.X)
	SoftBodyGetPressureCoefficient(body RID.Any) Float.X
	SoftBodySetDampingCoefficient(body RID.Any, damping_coefficient Float.X)
	SoftBodyGetDampingCoefficient(body RID.Any) Float.X
	SoftBodySetDragCoefficient(body RID.Any, drag_coefficient Float.X)
	SoftBodyGetDragCoefficient(body RID.Any) Float.X
	SoftBodySetMesh(body RID.Any, mesh RID.Any)
	SoftBodyGetBounds(body RID.Any) AABB.PositionSize
	SoftBodyMovePoint(body RID.Any, point_index int, global_position Vector3.XYZ)
	SoftBodyGetPointGlobalPosition(body RID.Any, point_index int) Vector3.XYZ
	SoftBodyRemoveAllPinnedPoints(body RID.Any)
	SoftBodyPinPoint(body RID.Any, point_index int, pin bool)
	SoftBodyIsPointPinned(body RID.Any, point_index int) bool
	JointCreate() RID.Any
	JointClear(joint RID.Any)
	JointMakePin(joint RID.Any, body_A RID.Any, local_A Vector3.XYZ, body_B RID.Any, local_B Vector3.XYZ)
	PinJointSetParam(joint RID.Any, param gdclass.PhysicsServer3DPinJointParam, value Float.X)
	PinJointGetParam(joint RID.Any, param gdclass.PhysicsServer3DPinJointParam) Float.X
	PinJointSetLocalA(joint RID.Any, local_A Vector3.XYZ)
	PinJointGetLocalA(joint RID.Any) Vector3.XYZ
	PinJointSetLocalB(joint RID.Any, local_B Vector3.XYZ)
	PinJointGetLocalB(joint RID.Any) Vector3.XYZ
	JointMakeHinge(joint RID.Any, body_A RID.Any, hinge_A Transform3D.BasisOrigin, body_B RID.Any, hinge_B Transform3D.BasisOrigin)
	JointMakeHingeSimple(joint RID.Any, body_A RID.Any, pivot_A Vector3.XYZ, axis_A Vector3.XYZ, body_B RID.Any, pivot_B Vector3.XYZ, axis_B Vector3.XYZ)
	HingeJointSetParam(joint RID.Any, param gdclass.PhysicsServer3DHingeJointParam, value Float.X)
	HingeJointGetParam(joint RID.Any, param gdclass.PhysicsServer3DHingeJointParam) Float.X
	HingeJointSetFlag(joint RID.Any, flag gdclass.PhysicsServer3DHingeJointFlag, enabled bool)
	HingeJointGetFlag(joint RID.Any, flag gdclass.PhysicsServer3DHingeJointFlag) bool
	JointMakeSlider(joint RID.Any, body_A RID.Any, local_ref_A Transform3D.BasisOrigin, body_B RID.Any, local_ref_B Transform3D.BasisOrigin)
	SliderJointSetParam(joint RID.Any, param gdclass.PhysicsServer3DSliderJointParam, value Float.X)
	SliderJointGetParam(joint RID.Any, param gdclass.PhysicsServer3DSliderJointParam) Float.X
	JointMakeConeTwist(joint RID.Any, body_A RID.Any, local_ref_A Transform3D.BasisOrigin, body_B RID.Any, local_ref_B Transform3D.BasisOrigin)
	ConeTwistJointSetParam(joint RID.Any, param gdclass.PhysicsServer3DConeTwistJointParam, value Float.X)
	ConeTwistJointGetParam(joint RID.Any, param gdclass.PhysicsServer3DConeTwistJointParam) Float.X
	JointMakeGeneric6dof(joint RID.Any, body_A RID.Any, local_ref_A Transform3D.BasisOrigin, body_B RID.Any, local_ref_B Transform3D.BasisOrigin)
	Generic6dofJointSetParam(joint RID.Any, axis gd.Vector3Axis, param gdclass.PhysicsServer3DG6DOFJointAxisParam, value Float.X)
	Generic6dofJointGetParam(joint RID.Any, axis gd.Vector3Axis, param gdclass.PhysicsServer3DG6DOFJointAxisParam) Float.X
	Generic6dofJointSetFlag(joint RID.Any, axis gd.Vector3Axis, flag gdclass.PhysicsServer3DG6DOFJointAxisFlag, enable bool)
	Generic6dofJointGetFlag(joint RID.Any, axis gd.Vector3Axis, flag gdclass.PhysicsServer3DG6DOFJointAxisFlag) bool
	JointGetType(joint RID.Any) gdclass.PhysicsServer3DJointType
	JointSetSolverPriority(joint RID.Any, priority int)
	JointGetSolverPriority(joint RID.Any) int
	JointDisableCollisionsBetweenBodies(joint RID.Any, disable bool)
	JointIsDisabledCollisionsBetweenBodies(joint RID.Any) bool
	FreeRid(rid RID.Any)
	SetActive(active bool)
	Init()
	Step(step Float.X)
	Sync()
	FlushQueries()
	EndSync()
	Finish()
	IsFlushingQueries() bool
	GetProcessInfo(process_info gdclass.PhysicsServer3DProcessInfo) int
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) WorldBoundaryShapeCreate() (_ RID.Any)                           { return }
func (self implementation) SeparationRayShapeCreate() (_ RID.Any)                           { return }
func (self implementation) SphereShapeCreate() (_ RID.Any)                                  { return }
func (self implementation) BoxShapeCreate() (_ RID.Any)                                     { return }
func (self implementation) CapsuleShapeCreate() (_ RID.Any)                                 { return }
func (self implementation) CylinderShapeCreate() (_ RID.Any)                                { return }
func (self implementation) ConvexPolygonShapeCreate() (_ RID.Any)                           { return }
func (self implementation) ConcavePolygonShapeCreate() (_ RID.Any)                          { return }
func (self implementation) HeightmapShapeCreate() (_ RID.Any)                               { return }
func (self implementation) CustomShapeCreate() (_ RID.Any)                                  { return }
func (self implementation) ShapeSetData(shape RID.Any, data any)                            { return }
func (self implementation) ShapeSetCustomSolverBias(shape RID.Any, bias Float.X)            { return }
func (self implementation) ShapeSetMargin(shape RID.Any, margin Float.X)                    { return }
func (self implementation) ShapeGetMargin(shape RID.Any) (_ Float.X)                        { return }
func (self implementation) ShapeGetType(shape RID.Any) (_ gdclass.PhysicsServer3DShapeType) { return }
func (self implementation) ShapeGetData(shape RID.Any) (_ any)                              { return }
func (self implementation) ShapeGetCustomSolverBias(shape RID.Any) (_ Float.X)              { return }
func (self implementation) SpaceCreate() (_ RID.Any)                                        { return }
func (self implementation) SpaceSetActive(space RID.Any, active bool)                       { return }
func (self implementation) SpaceIsActive(space RID.Any) (_ bool)                            { return }
func (self implementation) SpaceSetParam(space RID.Any, param gdclass.PhysicsServer3DSpaceParameter, value Float.X) {
	return
}
func (self implementation) SpaceGetParam(space RID.Any, param gdclass.PhysicsServer3DSpaceParameter) (_ Float.X) {
	return
}
func (self implementation) SpaceGetDirectState(space RID.Any) (_ [1]gdclass.PhysicsDirectSpaceState3D) {
	return
}
func (self implementation) SpaceSetDebugContacts(space RID.Any, max_contacts int) { return }
func (self implementation) SpaceGetContacts(space RID.Any) (_ []Vector3.XYZ)      { return }
func (self implementation) SpaceGetContactCount(space RID.Any) (_ int)            { return }
func (self implementation) AreaCreate() (_ RID.Any)                               { return }
func (self implementation) AreaSetSpace(area RID.Any, space RID.Any)              { return }
func (self implementation) AreaGetSpace(area RID.Any) (_ RID.Any)                 { return }
func (self implementation) AreaAddShape(area RID.Any, shape RID.Any, transform Transform3D.BasisOrigin, disabled bool) {
	return
}
func (self implementation) AreaSetShape(area RID.Any, shape_idx int, shape RID.Any) { return }
func (self implementation) AreaSetShapeTransform(area RID.Any, shape_idx int, transform Transform3D.BasisOrigin) {
	return
}
func (self implementation) AreaSetShapeDisabled(area RID.Any, shape_idx int, disabled bool) { return }
func (self implementation) AreaGetShapeCount(area RID.Any) (_ int)                          { return }
func (self implementation) AreaGetShape(area RID.Any, shape_idx int) (_ RID.Any)            { return }
func (self implementation) AreaGetShapeTransform(area RID.Any, shape_idx int) (_ Transform3D.BasisOrigin) {
	return
}
func (self implementation) AreaRemoveShape(area RID.Any, shape_idx int)     { return }
func (self implementation) AreaClearShapes(area RID.Any)                    { return }
func (self implementation) AreaAttachObjectInstanceId(area RID.Any, id int) { return }
func (self implementation) AreaGetObjectInstanceId(area RID.Any) (_ int)    { return }
func (self implementation) AreaSetParam(area RID.Any, param gdclass.PhysicsServer3DAreaParameter, value any) {
	return
}
func (self implementation) AreaSetTransform(area RID.Any, transform Transform3D.BasisOrigin) { return }
func (self implementation) AreaGetParam(area RID.Any, param gdclass.PhysicsServer3DAreaParameter) (_ any) {
	return
}
func (self implementation) AreaGetTransform(area RID.Any) (_ Transform3D.BasisOrigin)       { return }
func (self implementation) AreaSetCollisionLayer(area RID.Any, layer int)                   { return }
func (self implementation) AreaGetCollisionLayer(area RID.Any) (_ int)                      { return }
func (self implementation) AreaSetCollisionMask(area RID.Any, mask int)                     { return }
func (self implementation) AreaGetCollisionMask(area RID.Any) (_ int)                       { return }
func (self implementation) AreaSetMonitorable(area RID.Any, monitorable bool)               { return }
func (self implementation) AreaSetRayPickable(area RID.Any, enable bool)                    { return }
func (self implementation) AreaSetMonitorCallback(area RID.Any, callback Callable.Function) { return }
func (self implementation) AreaSetAreaMonitorCallback(area RID.Any, callback Callable.Function) {
	return
}
func (self implementation) BodyCreate() (_ RID.Any)                                        { return }
func (self implementation) BodySetSpace(body RID.Any, space RID.Any)                       { return }
func (self implementation) BodyGetSpace(body RID.Any) (_ RID.Any)                          { return }
func (self implementation) BodySetMode(body RID.Any, mode gdclass.PhysicsServer3DBodyMode) { return }
func (self implementation) BodyGetMode(body RID.Any) (_ gdclass.PhysicsServer3DBodyMode)   { return }
func (self implementation) BodyAddShape(body RID.Any, shape RID.Any, transform Transform3D.BasisOrigin, disabled bool) {
	return
}
func (self implementation) BodySetShape(body RID.Any, shape_idx int, shape RID.Any) { return }
func (self implementation) BodySetShapeTransform(body RID.Any, shape_idx int, transform Transform3D.BasisOrigin) {
	return
}
func (self implementation) BodySetShapeDisabled(body RID.Any, shape_idx int, disabled bool) { return }
func (self implementation) BodyGetShapeCount(body RID.Any) (_ int)                          { return }
func (self implementation) BodyGetShape(body RID.Any, shape_idx int) (_ RID.Any)            { return }
func (self implementation) BodyGetShapeTransform(body RID.Any, shape_idx int) (_ Transform3D.BasisOrigin) {
	return
}
func (self implementation) BodyRemoveShape(body RID.Any, shape_idx int)     { return }
func (self implementation) BodyClearShapes(body RID.Any)                    { return }
func (self implementation) BodyAttachObjectInstanceId(body RID.Any, id int) { return }
func (self implementation) BodyGetObjectInstanceId(body RID.Any) (_ int)    { return }
func (self implementation) BodySetEnableContinuousCollisionDetection(body RID.Any, enable bool) {
	return
}
func (self implementation) BodyIsContinuousCollisionDetectionEnabled(body RID.Any) (_ bool) { return }
func (self implementation) BodySetCollisionLayer(body RID.Any, layer int)                   { return }
func (self implementation) BodyGetCollisionLayer(body RID.Any) (_ int)                      { return }
func (self implementation) BodySetCollisionMask(body RID.Any, mask int)                     { return }
func (self implementation) BodyGetCollisionMask(body RID.Any) (_ int)                       { return }
func (self implementation) BodySetCollisionPriority(body RID.Any, priority Float.X)         { return }
func (self implementation) BodyGetCollisionPriority(body RID.Any) (_ Float.X)               { return }
func (self implementation) BodySetUserFlags(body RID.Any, flags int)                        { return }
func (self implementation) BodyGetUserFlags(body RID.Any) (_ int)                           { return }
func (self implementation) BodySetParam(body RID.Any, param gdclass.PhysicsServer3DBodyParameter, value any) {
	return
}
func (self implementation) BodyGetParam(body RID.Any, param gdclass.PhysicsServer3DBodyParameter) (_ any) {
	return
}
func (self implementation) BodyResetMassProperties(body RID.Any) { return }
func (self implementation) BodySetState(body RID.Any, state gdclass.PhysicsServer3DBodyState, value any) {
	return
}
func (self implementation) BodyGetState(body RID.Any, state gdclass.PhysicsServer3DBodyState) (_ any) {
	return
}
func (self implementation) BodyApplyCentralImpulse(body RID.Any, impulse Vector3.XYZ) { return }
func (self implementation) BodyApplyImpulse(body RID.Any, impulse Vector3.XYZ, position Vector3.XYZ) {
	return
}
func (self implementation) BodyApplyTorqueImpulse(body RID.Any, impulse Vector3.XYZ) { return }
func (self implementation) BodyApplyCentralForce(body RID.Any, force Vector3.XYZ)    { return }
func (self implementation) BodyApplyForce(body RID.Any, force Vector3.XYZ, position Vector3.XYZ) {
	return
}
func (self implementation) BodyApplyTorque(body RID.Any, torque Vector3.XYZ)            { return }
func (self implementation) BodyAddConstantCentralForce(body RID.Any, force Vector3.XYZ) { return }
func (self implementation) BodyAddConstantForce(body RID.Any, force Vector3.XYZ, position Vector3.XYZ) {
	return
}
func (self implementation) BodyAddConstantTorque(body RID.Any, torque Vector3.XYZ)      { return }
func (self implementation) BodySetConstantForce(body RID.Any, force Vector3.XYZ)        { return }
func (self implementation) BodyGetConstantForce(body RID.Any) (_ Vector3.XYZ)           { return }
func (self implementation) BodySetConstantTorque(body RID.Any, torque Vector3.XYZ)      { return }
func (self implementation) BodyGetConstantTorque(body RID.Any) (_ Vector3.XYZ)          { return }
func (self implementation) BodySetAxisVelocity(body RID.Any, axis_velocity Vector3.XYZ) { return }
func (self implementation) BodySetAxisLock(body RID.Any, axis gdclass.PhysicsServer3DBodyAxis, lock bool) {
	return
}
func (self implementation) BodyIsAxisLocked(body RID.Any, axis gdclass.PhysicsServer3DBodyAxis) (_ bool) {
	return
}
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
func (self implementation) BodySetRayPickable(body RID.Any, enable bool) { return }
func (self implementation) BodyTestMotion(body RID.Any, from Transform3D.BasisOrigin, motion Vector3.XYZ, margin Float.X, max_collisions int, collide_separation_ray bool, recovery_as_collision bool, result *MotionResult) (_ bool) {
	return
}
func (self implementation) BodyGetDirectState(body RID.Any) (_ [1]gdclass.PhysicsDirectBodyState3D) {
	return
}
func (self implementation) SoftBodyCreate() (_ RID.Any) { return }
func (self implementation) SoftBodyUpdateRenderingServer(body RID.Any, rendering_server_handler [1]gdclass.PhysicsServer3DRenderingServerHandler) {
	return
}
func (self implementation) SoftBodySetSpace(body RID.Any, space RID.Any)                  { return }
func (self implementation) SoftBodyGetSpace(body RID.Any) (_ RID.Any)                     { return }
func (self implementation) SoftBodySetRayPickable(body RID.Any, enable bool)              { return }
func (self implementation) SoftBodySetCollisionLayer(body RID.Any, layer int)             { return }
func (self implementation) SoftBodyGetCollisionLayer(body RID.Any) (_ int)                { return }
func (self implementation) SoftBodySetCollisionMask(body RID.Any, mask int)               { return }
func (self implementation) SoftBodyGetCollisionMask(body RID.Any) (_ int)                 { return }
func (self implementation) SoftBodyAddCollisionException(body RID.Any, body_b RID.Any)    { return }
func (self implementation) SoftBodyRemoveCollisionException(body RID.Any, body_b RID.Any) { return }
func (self implementation) SoftBodyGetCollisionExceptions(body RID.Any) (_ []RID.Any)     { return }
func (self implementation) SoftBodySetState(body RID.Any, state gdclass.PhysicsServer3DBodyState, v any) {
	return
}
func (self implementation) SoftBodyGetState(body RID.Any, state gdclass.PhysicsServer3DBodyState) (_ any) {
	return
}
func (self implementation) SoftBodySetTransform(body RID.Any, transform Transform3D.BasisOrigin) {
	return
}
func (self implementation) SoftBodySetSimulationPrecision(body RID.Any, simulation_precision int) {
	return
}
func (self implementation) SoftBodyGetSimulationPrecision(body RID.Any) (_ int)               { return }
func (self implementation) SoftBodySetTotalMass(body RID.Any, total_mass Float.X)             { return }
func (self implementation) SoftBodyGetTotalMass(body RID.Any) (_ Float.X)                     { return }
func (self implementation) SoftBodySetLinearStiffness(body RID.Any, linear_stiffness Float.X) { return }
func (self implementation) SoftBodyGetLinearStiffness(body RID.Any) (_ Float.X)               { return }
func (self implementation) SoftBodySetPressureCoefficient(body RID.Any, pressure_coefficient Float.X) {
	return
}
func (self implementation) SoftBodyGetPressureCoefficient(body RID.Any) (_ Float.X) { return }
func (self implementation) SoftBodySetDampingCoefficient(body RID.Any, damping_coefficient Float.X) {
	return
}
func (self implementation) SoftBodyGetDampingCoefficient(body RID.Any) (_ Float.X)            { return }
func (self implementation) SoftBodySetDragCoefficient(body RID.Any, drag_coefficient Float.X) { return }
func (self implementation) SoftBodyGetDragCoefficient(body RID.Any) (_ Float.X)               { return }
func (self implementation) SoftBodySetMesh(body RID.Any, mesh RID.Any)                        { return }
func (self implementation) SoftBodyGetBounds(body RID.Any) (_ AABB.PositionSize)              { return }
func (self implementation) SoftBodyMovePoint(body RID.Any, point_index int, global_position Vector3.XYZ) {
	return
}
func (self implementation) SoftBodyGetPointGlobalPosition(body RID.Any, point_index int) (_ Vector3.XYZ) {
	return
}
func (self implementation) SoftBodyRemoveAllPinnedPoints(body RID.Any)                   { return }
func (self implementation) SoftBodyPinPoint(body RID.Any, point_index int, pin bool)     { return }
func (self implementation) SoftBodyIsPointPinned(body RID.Any, point_index int) (_ bool) { return }
func (self implementation) JointCreate() (_ RID.Any)                                     { return }
func (self implementation) JointClear(joint RID.Any)                                     { return }
func (self implementation) JointMakePin(joint RID.Any, body_A RID.Any, local_A Vector3.XYZ, body_B RID.Any, local_B Vector3.XYZ) {
	return
}
func (self implementation) PinJointSetParam(joint RID.Any, param gdclass.PhysicsServer3DPinJointParam, value Float.X) {
	return
}
func (self implementation) PinJointGetParam(joint RID.Any, param gdclass.PhysicsServer3DPinJointParam) (_ Float.X) {
	return
}
func (self implementation) PinJointSetLocalA(joint RID.Any, local_A Vector3.XYZ) { return }
func (self implementation) PinJointGetLocalA(joint RID.Any) (_ Vector3.XYZ)      { return }
func (self implementation) PinJointSetLocalB(joint RID.Any, local_B Vector3.XYZ) { return }
func (self implementation) PinJointGetLocalB(joint RID.Any) (_ Vector3.XYZ)      { return }
func (self implementation) JointMakeHinge(joint RID.Any, body_A RID.Any, hinge_A Transform3D.BasisOrigin, body_B RID.Any, hinge_B Transform3D.BasisOrigin) {
	return
}
func (self implementation) JointMakeHingeSimple(joint RID.Any, body_A RID.Any, pivot_A Vector3.XYZ, axis_A Vector3.XYZ, body_B RID.Any, pivot_B Vector3.XYZ, axis_B Vector3.XYZ) {
	return
}
func (self implementation) HingeJointSetParam(joint RID.Any, param gdclass.PhysicsServer3DHingeJointParam, value Float.X) {
	return
}
func (self implementation) HingeJointGetParam(joint RID.Any, param gdclass.PhysicsServer3DHingeJointParam) (_ Float.X) {
	return
}
func (self implementation) HingeJointSetFlag(joint RID.Any, flag gdclass.PhysicsServer3DHingeJointFlag, enabled bool) {
	return
}
func (self implementation) HingeJointGetFlag(joint RID.Any, flag gdclass.PhysicsServer3DHingeJointFlag) (_ bool) {
	return
}
func (self implementation) JointMakeSlider(joint RID.Any, body_A RID.Any, local_ref_A Transform3D.BasisOrigin, body_B RID.Any, local_ref_B Transform3D.BasisOrigin) {
	return
}
func (self implementation) SliderJointSetParam(joint RID.Any, param gdclass.PhysicsServer3DSliderJointParam, value Float.X) {
	return
}
func (self implementation) SliderJointGetParam(joint RID.Any, param gdclass.PhysicsServer3DSliderJointParam) (_ Float.X) {
	return
}
func (self implementation) JointMakeConeTwist(joint RID.Any, body_A RID.Any, local_ref_A Transform3D.BasisOrigin, body_B RID.Any, local_ref_B Transform3D.BasisOrigin) {
	return
}
func (self implementation) ConeTwistJointSetParam(joint RID.Any, param gdclass.PhysicsServer3DConeTwistJointParam, value Float.X) {
	return
}
func (self implementation) ConeTwistJointGetParam(joint RID.Any, param gdclass.PhysicsServer3DConeTwistJointParam) (_ Float.X) {
	return
}
func (self implementation) JointMakeGeneric6dof(joint RID.Any, body_A RID.Any, local_ref_A Transform3D.BasisOrigin, body_B RID.Any, local_ref_B Transform3D.BasisOrigin) {
	return
}
func (self implementation) Generic6dofJointSetParam(joint RID.Any, axis gd.Vector3Axis, param gdclass.PhysicsServer3DG6DOFJointAxisParam, value Float.X) {
	return
}
func (self implementation) Generic6dofJointGetParam(joint RID.Any, axis gd.Vector3Axis, param gdclass.PhysicsServer3DG6DOFJointAxisParam) (_ Float.X) {
	return
}
func (self implementation) Generic6dofJointSetFlag(joint RID.Any, axis gd.Vector3Axis, flag gdclass.PhysicsServer3DG6DOFJointAxisFlag, enable bool) {
	return
}
func (self implementation) Generic6dofJointGetFlag(joint RID.Any, axis gd.Vector3Axis, flag gdclass.PhysicsServer3DG6DOFJointAxisFlag) (_ bool) {
	return
}
func (self implementation) JointGetType(joint RID.Any) (_ gdclass.PhysicsServer3DJointType) { return }
func (self implementation) JointSetSolverPriority(joint RID.Any, priority int)              { return }
func (self implementation) JointGetSolverPriority(joint RID.Any) (_ int)                    { return }
func (self implementation) JointDisableCollisionsBetweenBodies(joint RID.Any, disable bool) { return }
func (self implementation) JointIsDisabledCollisionsBetweenBodies(joint RID.Any) (_ bool)   { return }
func (self implementation) FreeRid(rid RID.Any)                                             { return }
func (self implementation) SetActive(active bool)                                           { return }
func (self implementation) Init()                                                           { return }
func (self implementation) Step(step Float.X)                                               { return }
func (self implementation) Sync()                                                           { return }
func (self implementation) FlushQueries()                                                   { return }
func (self implementation) EndSync()                                                        { return }
func (self implementation) Finish()                                                         { return }
func (self implementation) IsFlushingQueries() (_ bool)                                     { return }
func (self implementation) GetProcessInfo(process_info gdclass.PhysicsServer3DProcessInfo) (_ int) {
	return
}
func (Instance) _world_boundary_shape_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, RID.Any(ret))
	}
}
func (Instance) _separation_ray_shape_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, RID.Any(ret))
	}
}
func (Instance) _sphere_shape_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, RID.Any(ret))
	}
}
func (Instance) _box_shape_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, RID.Any(ret))
	}
}
func (Instance) _capsule_shape_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, RID.Any(ret))
	}
}
func (Instance) _cylinder_shape_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, RID.Any(ret))
	}
}
func (Instance) _convex_polygon_shape_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, RID.Any(ret))
	}
}
func (Instance) _concave_polygon_shape_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, RID.Any(ret))
	}
}
func (Instance) _heightmap_shape_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, RID.Any(ret))
	}
}
func (Instance) _custom_shape_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, RID.Any(ret))
	}
}
func (Instance) _shape_set_data(impl func(ptr unsafe.Pointer, shape RID.Any, data any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape = gd.UnsafeGet[RID.Any](p_args, 0)

		var data = variant.Through(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 1))))
		defer pointers.End(gd.InternalVariant(data))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shape, data.Interface())
	}
}
func (Instance) _shape_set_custom_solver_bias(impl func(ptr unsafe.Pointer, shape RID.Any, bias Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape = gd.UnsafeGet[RID.Any](p_args, 0)

		var bias = gd.UnsafeGet[float64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shape, Float.X(bias))
	}
}
func (Instance) _shape_set_margin(impl func(ptr unsafe.Pointer, shape RID.Any, margin Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape = gd.UnsafeGet[RID.Any](p_args, 0)

		var margin = gd.UnsafeGet[float64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shape, Float.X(margin))
	}
}
func (Instance) _shape_get_margin(impl func(ptr unsafe.Pointer, shape RID.Any) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		gd.UnsafeSet(p_back, float64(ret))
	}
}
func (Instance) _shape_get_type(impl func(ptr unsafe.Pointer, shape RID.Any) gdclass.PhysicsServer3DShapeType) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _shape_get_data(impl func(ptr unsafe.Pointer, shape RID.Any) any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		ptr, ok := pointers.End(gd.InternalVariant(variant.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _shape_get_custom_solver_bias(impl func(ptr unsafe.Pointer, shape RID.Any) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		gd.UnsafeSet(p_back, float64(ret))
	}
}
func (Instance) _space_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, RID.Any(ret))
	}
}
func (Instance) _space_set_active(impl func(ptr unsafe.Pointer, space RID.Any, active bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[RID.Any](p_args, 0)

		var active = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, space, active)
	}
}
func (Instance) _space_is_active(impl func(ptr unsafe.Pointer, space RID.Any) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _space_set_param(impl func(ptr unsafe.Pointer, space RID.Any, param gdclass.PhysicsServer3DSpaceParameter, value Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[RID.Any](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DSpaceParameter](p_args, 1)

		var value = gd.UnsafeGet[float64](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, space, param, Float.X(value))
	}
}
func (Instance) _space_get_param(impl func(ptr unsafe.Pointer, space RID.Any, param gdclass.PhysicsServer3DSpaceParameter) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[RID.Any](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DSpaceParameter](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space, param)
		gd.UnsafeSet(p_back, float64(ret))
	}
}
func (Instance) _space_get_direct_state(impl func(ptr unsafe.Pointer, space RID.Any) [1]gdclass.PhysicsDirectSpaceState3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _space_set_debug_contacts(impl func(ptr unsafe.Pointer, space RID.Any, max_contacts int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[RID.Any](p_args, 0)

		var max_contacts = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, space, int(max_contacts))
	}
}
func (Instance) _space_get_contacts(impl func(ptr unsafe.Pointer, space RID.Any) []Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		ptr, ok := pointers.End(gd.InternalPacked[gd.PackedVector3Array, Vector3.XYZ](Packed.New(ret...)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _space_get_contact_count(impl func(ptr unsafe.Pointer, space RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		gd.UnsafeSet(p_back, int64(ret))
	}
}
func (Instance) _area_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, RID.Any(ret))
	}
}
func (Instance) _area_set_space(impl func(ptr unsafe.Pointer, area RID.Any, space RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var space = gd.UnsafeGet[RID.Any](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, space)
	}
}
func (Instance) _area_get_space(impl func(ptr unsafe.Pointer, area RID.Any) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, RID.Any(ret))
	}
}
func (Instance) _area_add_shape(impl func(ptr unsafe.Pointer, area RID.Any, shape RID.Any, transform Transform3D.BasisOrigin, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var shape = gd.UnsafeGet[RID.Any](p_args, 1)

		var transform = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 2)

		var disabled = gd.UnsafeGet[bool](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, shape, transform, disabled)
	}
}
func (Instance) _area_set_shape(impl func(ptr unsafe.Pointer, area RID.Any, shape_idx int, shape RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var shape_idx = gd.UnsafeGet[int64](p_args, 1)

		var shape = gd.UnsafeGet[RID.Any](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, int(shape_idx), shape)
	}
}
func (Instance) _area_set_shape_transform(impl func(ptr unsafe.Pointer, area RID.Any, shape_idx int, transform Transform3D.BasisOrigin)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var shape_idx = gd.UnsafeGet[int64](p_args, 1)

		var transform = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, int(shape_idx), transform)
	}
}
func (Instance) _area_set_shape_disabled(impl func(ptr unsafe.Pointer, area RID.Any, shape_idx int, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var shape_idx = gd.UnsafeGet[int64](p_args, 1)

		var disabled = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, int(shape_idx), disabled)
	}
}
func (Instance) _area_get_shape_count(impl func(ptr unsafe.Pointer, area RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, int64(ret))
	}
}
func (Instance) _area_get_shape(impl func(ptr unsafe.Pointer, area RID.Any, shape_idx int) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var shape_idx = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, int(shape_idx))
		gd.UnsafeSet(p_back, RID.Any(ret))
	}
}
func (Instance) _area_get_shape_transform(impl func(ptr unsafe.Pointer, area RID.Any, shape_idx int) Transform3D.BasisOrigin) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var shape_idx = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, int(shape_idx))
		gd.UnsafeSet(p_back, Transform3D.BasisOrigin(ret))
	}
}
func (Instance) _area_remove_shape(impl func(ptr unsafe.Pointer, area RID.Any, shape_idx int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var shape_idx = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, int(shape_idx))
	}
}
func (Instance) _area_clear_shapes(impl func(ptr unsafe.Pointer, area RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area)
	}
}
func (Instance) _area_attach_object_instance_id(impl func(ptr unsafe.Pointer, area RID.Any, id int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var id = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, int(id))
	}
}
func (Instance) _area_get_object_instance_id(impl func(ptr unsafe.Pointer, area RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, int64(ret))
	}
}
func (Instance) _area_set_param(impl func(ptr unsafe.Pointer, area RID.Any, param gdclass.PhysicsServer3DAreaParameter, value any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DAreaParameter](p_args, 1)

		var value = variant.Through(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 2))))
		defer pointers.End(gd.InternalVariant(value))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, param, value.Interface())
	}
}
func (Instance) _area_set_transform(impl func(ptr unsafe.Pointer, area RID.Any, transform Transform3D.BasisOrigin)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var transform = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, transform)
	}
}
func (Instance) _area_get_param(impl func(ptr unsafe.Pointer, area RID.Any, param gdclass.PhysicsServer3DAreaParameter) any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DAreaParameter](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, param)
		ptr, ok := pointers.End(gd.InternalVariant(variant.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _area_get_transform(impl func(ptr unsafe.Pointer, area RID.Any) Transform3D.BasisOrigin) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, Transform3D.BasisOrigin(ret))
	}
}
func (Instance) _area_set_collision_layer(impl func(ptr unsafe.Pointer, area RID.Any, layer int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var layer = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, int(layer))
	}
}
func (Instance) _area_get_collision_layer(impl func(ptr unsafe.Pointer, area RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, int64(ret))
	}
}
func (Instance) _area_set_collision_mask(impl func(ptr unsafe.Pointer, area RID.Any, mask int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var mask = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, int(mask))
	}
}
func (Instance) _area_get_collision_mask(impl func(ptr unsafe.Pointer, area RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, int64(ret))
	}
}
func (Instance) _area_set_monitorable(impl func(ptr unsafe.Pointer, area RID.Any, monitorable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var monitorable = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, monitorable)
	}
}
func (Instance) _area_set_ray_pickable(impl func(ptr unsafe.Pointer, area RID.Any, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var enable = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, enable)
	}
}
func (Instance) _area_set_monitor_callback(impl func(ptr unsafe.Pointer, area RID.Any, callback Callable.Function)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var callback = Callable.Through(gd.CallableProxy{}, pointers.Pack(pointers.New[gd.Callable](gd.UnsafeGet[[2]uint64](p_args, 1))))
		defer pointers.End(gd.InternalCallable(callback))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, callback)
	}
}
func (Instance) _area_set_area_monitor_callback(impl func(ptr unsafe.Pointer, area RID.Any, callback Callable.Function)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var callback = Callable.Through(gd.CallableProxy{}, pointers.Pack(pointers.New[gd.Callable](gd.UnsafeGet[[2]uint64](p_args, 1))))
		defer pointers.End(gd.InternalCallable(callback))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, callback)
	}
}
func (Instance) _body_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, RID.Any(ret))
	}
}
func (Instance) _body_set_space(impl func(ptr unsafe.Pointer, body RID.Any, space RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var space = gd.UnsafeGet[RID.Any](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, space)
	}
}
func (Instance) _body_get_space(impl func(ptr unsafe.Pointer, body RID.Any) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, RID.Any(ret))
	}
}
func (Instance) _body_set_mode(impl func(ptr unsafe.Pointer, body RID.Any, mode gdclass.PhysicsServer3DBodyMode)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var mode = gd.UnsafeGet[gdclass.PhysicsServer3DBodyMode](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, mode)
	}
}
func (Instance) _body_get_mode(impl func(ptr unsafe.Pointer, body RID.Any) gdclass.PhysicsServer3DBodyMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _body_add_shape(impl func(ptr unsafe.Pointer, body RID.Any, shape RID.Any, transform Transform3D.BasisOrigin, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var shape = gd.UnsafeGet[RID.Any](p_args, 1)

		var transform = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 2)

		var disabled = gd.UnsafeGet[bool](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, shape, transform, disabled)
	}
}
func (Instance) _body_set_shape(impl func(ptr unsafe.Pointer, body RID.Any, shape_idx int, shape RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var shape_idx = gd.UnsafeGet[int64](p_args, 1)

		var shape = gd.UnsafeGet[RID.Any](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(shape_idx), shape)
	}
}
func (Instance) _body_set_shape_transform(impl func(ptr unsafe.Pointer, body RID.Any, shape_idx int, transform Transform3D.BasisOrigin)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var shape_idx = gd.UnsafeGet[int64](p_args, 1)

		var transform = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(shape_idx), transform)
	}
}
func (Instance) _body_set_shape_disabled(impl func(ptr unsafe.Pointer, body RID.Any, shape_idx int, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var shape_idx = gd.UnsafeGet[int64](p_args, 1)

		var disabled = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(shape_idx), disabled)
	}
}
func (Instance) _body_get_shape_count(impl func(ptr unsafe.Pointer, body RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, int64(ret))
	}
}
func (Instance) _body_get_shape(impl func(ptr unsafe.Pointer, body RID.Any, shape_idx int) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var shape_idx = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, int(shape_idx))
		gd.UnsafeSet(p_back, RID.Any(ret))
	}
}
func (Instance) _body_get_shape_transform(impl func(ptr unsafe.Pointer, body RID.Any, shape_idx int) Transform3D.BasisOrigin) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var shape_idx = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, int(shape_idx))
		gd.UnsafeSet(p_back, Transform3D.BasisOrigin(ret))
	}
}
func (Instance) _body_remove_shape(impl func(ptr unsafe.Pointer, body RID.Any, shape_idx int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var shape_idx = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(shape_idx))
	}
}
func (Instance) _body_clear_shapes(impl func(ptr unsafe.Pointer, body RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body)
	}
}
func (Instance) _body_attach_object_instance_id(impl func(ptr unsafe.Pointer, body RID.Any, id int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var id = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(id))
	}
}
func (Instance) _body_get_object_instance_id(impl func(ptr unsafe.Pointer, body RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, int64(ret))
	}
}
func (Instance) _body_set_enable_continuous_collision_detection(impl func(ptr unsafe.Pointer, body RID.Any, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var enable = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, enable)
	}
}
func (Instance) _body_is_continuous_collision_detection_enabled(impl func(ptr unsafe.Pointer, body RID.Any) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _body_set_collision_layer(impl func(ptr unsafe.Pointer, body RID.Any, layer int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var layer = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(layer))
	}
}
func (Instance) _body_get_collision_layer(impl func(ptr unsafe.Pointer, body RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, int64(ret))
	}
}
func (Instance) _body_set_collision_mask(impl func(ptr unsafe.Pointer, body RID.Any, mask int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var mask = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(mask))
	}
}
func (Instance) _body_get_collision_mask(impl func(ptr unsafe.Pointer, body RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, int64(ret))
	}
}
func (Instance) _body_set_collision_priority(impl func(ptr unsafe.Pointer, body RID.Any, priority Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var priority = gd.UnsafeGet[float64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, Float.X(priority))
	}
}
func (Instance) _body_get_collision_priority(impl func(ptr unsafe.Pointer, body RID.Any) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, float64(ret))
	}
}
func (Instance) _body_set_user_flags(impl func(ptr unsafe.Pointer, body RID.Any, flags int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var flags = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(flags))
	}
}
func (Instance) _body_get_user_flags(impl func(ptr unsafe.Pointer, body RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, int64(ret))
	}
}
func (Instance) _body_set_param(impl func(ptr unsafe.Pointer, body RID.Any, param gdclass.PhysicsServer3DBodyParameter, value any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DBodyParameter](p_args, 1)

		var value = variant.Through(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 2))))
		defer pointers.End(gd.InternalVariant(value))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, param, value.Interface())
	}
}
func (Instance) _body_get_param(impl func(ptr unsafe.Pointer, body RID.Any, param gdclass.PhysicsServer3DBodyParameter) any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DBodyParameter](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, param)
		ptr, ok := pointers.End(gd.InternalVariant(variant.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _body_reset_mass_properties(impl func(ptr unsafe.Pointer, body RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body)
	}
}
func (Instance) _body_set_state(impl func(ptr unsafe.Pointer, body RID.Any, state gdclass.PhysicsServer3DBodyState, value any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var state = gd.UnsafeGet[gdclass.PhysicsServer3DBodyState](p_args, 1)

		var value = variant.Through(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 2))))
		defer pointers.End(gd.InternalVariant(value))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, state, value.Interface())
	}
}
func (Instance) _body_get_state(impl func(ptr unsafe.Pointer, body RID.Any, state gdclass.PhysicsServer3DBodyState) any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var state = gd.UnsafeGet[gdclass.PhysicsServer3DBodyState](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, state)
		ptr, ok := pointers.End(gd.InternalVariant(variant.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _body_apply_central_impulse(impl func(ptr unsafe.Pointer, body RID.Any, impulse Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var impulse = gd.UnsafeGet[Vector3.XYZ](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, impulse)
	}
}
func (Instance) _body_apply_impulse(impl func(ptr unsafe.Pointer, body RID.Any, impulse Vector3.XYZ, position Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var impulse = gd.UnsafeGet[Vector3.XYZ](p_args, 1)

		var position = gd.UnsafeGet[Vector3.XYZ](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, impulse, position)
	}
}
func (Instance) _body_apply_torque_impulse(impl func(ptr unsafe.Pointer, body RID.Any, impulse Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var impulse = gd.UnsafeGet[Vector3.XYZ](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, impulse)
	}
}
func (Instance) _body_apply_central_force(impl func(ptr unsafe.Pointer, body RID.Any, force Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var force = gd.UnsafeGet[Vector3.XYZ](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force)
	}
}
func (Instance) _body_apply_force(impl func(ptr unsafe.Pointer, body RID.Any, force Vector3.XYZ, position Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var force = gd.UnsafeGet[Vector3.XYZ](p_args, 1)

		var position = gd.UnsafeGet[Vector3.XYZ](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force, position)
	}
}
func (Instance) _body_apply_torque(impl func(ptr unsafe.Pointer, body RID.Any, torque Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var torque = gd.UnsafeGet[Vector3.XYZ](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, torque)
	}
}
func (Instance) _body_add_constant_central_force(impl func(ptr unsafe.Pointer, body RID.Any, force Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var force = gd.UnsafeGet[Vector3.XYZ](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force)
	}
}
func (Instance) _body_add_constant_force(impl func(ptr unsafe.Pointer, body RID.Any, force Vector3.XYZ, position Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var force = gd.UnsafeGet[Vector3.XYZ](p_args, 1)

		var position = gd.UnsafeGet[Vector3.XYZ](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force, position)
	}
}
func (Instance) _body_add_constant_torque(impl func(ptr unsafe.Pointer, body RID.Any, torque Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var torque = gd.UnsafeGet[Vector3.XYZ](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, torque)
	}
}
func (Instance) _body_set_constant_force(impl func(ptr unsafe.Pointer, body RID.Any, force Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var force = gd.UnsafeGet[Vector3.XYZ](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force)
	}
}
func (Instance) _body_get_constant_force(impl func(ptr unsafe.Pointer, body RID.Any) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, Vector3.XYZ(ret))
	}
}
func (Instance) _body_set_constant_torque(impl func(ptr unsafe.Pointer, body RID.Any, torque Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var torque = gd.UnsafeGet[Vector3.XYZ](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, torque)
	}
}
func (Instance) _body_get_constant_torque(impl func(ptr unsafe.Pointer, body RID.Any) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, Vector3.XYZ(ret))
	}
}
func (Instance) _body_set_axis_velocity(impl func(ptr unsafe.Pointer, body RID.Any, axis_velocity Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var axis_velocity = gd.UnsafeGet[Vector3.XYZ](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, axis_velocity)
	}
}
func (Instance) _body_set_axis_lock(impl func(ptr unsafe.Pointer, body RID.Any, axis gdclass.PhysicsServer3DBodyAxis, lock bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var axis = gd.UnsafeGet[gdclass.PhysicsServer3DBodyAxis](p_args, 1)

		var lock = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, axis, lock)
	}
}
func (Instance) _body_is_axis_locked(impl func(ptr unsafe.Pointer, body RID.Any, axis gdclass.PhysicsServer3DBodyAxis) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var axis = gd.UnsafeGet[gdclass.PhysicsServer3DBodyAxis](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, axis)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _body_add_collision_exception(impl func(ptr unsafe.Pointer, body RID.Any, excepted_body RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var excepted_body = gd.UnsafeGet[RID.Any](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, excepted_body)
	}
}
func (Instance) _body_remove_collision_exception(impl func(ptr unsafe.Pointer, body RID.Any, excepted_body RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var excepted_body = gd.UnsafeGet[RID.Any](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, excepted_body)
	}
}
func (Instance) _body_get_collision_exceptions(impl func(ptr unsafe.Pointer, body RID.Any) []RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		ptr, ok := pointers.End(gd.InternalArray(gd.ArrayFromSlice[Array.Contains[RID.Any]](ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _body_set_max_contacts_reported(impl func(ptr unsafe.Pointer, body RID.Any, amount int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var amount = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(amount))
	}
}
func (Instance) _body_get_max_contacts_reported(impl func(ptr unsafe.Pointer, body RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, int64(ret))
	}
}
func (Instance) _body_set_contacts_reported_depth_threshold(impl func(ptr unsafe.Pointer, body RID.Any, threshold Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var threshold = gd.UnsafeGet[float64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, Float.X(threshold))
	}
}
func (Instance) _body_get_contacts_reported_depth_threshold(impl func(ptr unsafe.Pointer, body RID.Any) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, float64(ret))
	}
}
func (Instance) _body_set_omit_force_integration(impl func(ptr unsafe.Pointer, body RID.Any, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var enable = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, enable)
	}
}
func (Instance) _body_is_omitting_force_integration(impl func(ptr unsafe.Pointer, body RID.Any) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _body_set_state_sync_callback(impl func(ptr unsafe.Pointer, body RID.Any, callable Callable.Function)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var callable = Callable.Through(gd.CallableProxy{}, pointers.Pack(pointers.New[gd.Callable](gd.UnsafeGet[[2]uint64](p_args, 1))))
		defer pointers.End(gd.InternalCallable(callable))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, callable)
	}
}
func (Instance) _body_set_force_integration_callback(impl func(ptr unsafe.Pointer, body RID.Any, callable Callable.Function, userdata any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var callable = Callable.Through(gd.CallableProxy{}, pointers.Pack(pointers.New[gd.Callable](gd.UnsafeGet[[2]uint64](p_args, 1))))
		defer pointers.End(gd.InternalCallable(callable))
		var userdata = variant.Through(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 2))))
		defer pointers.End(gd.InternalVariant(userdata))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, callable, userdata.Interface())
	}
}
func (Instance) _body_set_ray_pickable(impl func(ptr unsafe.Pointer, body RID.Any, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var enable = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, enable)
	}
}
func (Instance) _body_test_motion(impl func(ptr unsafe.Pointer, body RID.Any, from Transform3D.BasisOrigin, motion Vector3.XYZ, margin Float.X, max_collisions int, collide_separation_ray bool, recovery_as_collision bool, result *MotionResult) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var from = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 1)

		var motion = gd.UnsafeGet[Vector3.XYZ](p_args, 2)

		var margin = gd.UnsafeGet[float64](p_args, 3)

		var max_collisions = gd.UnsafeGet[int64](p_args, 4)

		var collide_separation_ray = gd.UnsafeGet[bool](p_args, 5)

		var recovery_as_collision = gd.UnsafeGet[bool](p_args, 6)

		var result = gd.UnsafeGet[*MotionResult](p_args, 7)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, from, motion, Float.X(margin), int(max_collisions), collide_separation_ray, recovery_as_collision, result)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _body_get_direct_state(impl func(ptr unsafe.Pointer, body RID.Any) [1]gdclass.PhysicsDirectBodyState3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _soft_body_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, RID.Any(ret))
	}
}
func (Instance) _soft_body_update_rendering_server(impl func(ptr unsafe.Pointer, body RID.Any, rendering_server_handler [1]gdclass.PhysicsServer3DRenderingServerHandler)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var rendering_server_handler = [1]gdclass.PhysicsServer3DRenderingServerHandler{pointers.New[gdclass.PhysicsServer3DRenderingServerHandler]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}

		defer pointers.End(rendering_server_handler[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, rendering_server_handler)
	}
}
func (Instance) _soft_body_set_space(impl func(ptr unsafe.Pointer, body RID.Any, space RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var space = gd.UnsafeGet[RID.Any](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, space)
	}
}
func (Instance) _soft_body_get_space(impl func(ptr unsafe.Pointer, body RID.Any) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, RID.Any(ret))
	}
}
func (Instance) _soft_body_set_ray_pickable(impl func(ptr unsafe.Pointer, body RID.Any, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var enable = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, enable)
	}
}
func (Instance) _soft_body_set_collision_layer(impl func(ptr unsafe.Pointer, body RID.Any, layer int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var layer = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(layer))
	}
}
func (Instance) _soft_body_get_collision_layer(impl func(ptr unsafe.Pointer, body RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, int64(ret))
	}
}
func (Instance) _soft_body_set_collision_mask(impl func(ptr unsafe.Pointer, body RID.Any, mask int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var mask = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(mask))
	}
}
func (Instance) _soft_body_get_collision_mask(impl func(ptr unsafe.Pointer, body RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, int64(ret))
	}
}
func (Instance) _soft_body_add_collision_exception(impl func(ptr unsafe.Pointer, body RID.Any, body_b RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var body_b = gd.UnsafeGet[RID.Any](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, body_b)
	}
}
func (Instance) _soft_body_remove_collision_exception(impl func(ptr unsafe.Pointer, body RID.Any, body_b RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var body_b = gd.UnsafeGet[RID.Any](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, body_b)
	}
}
func (Instance) _soft_body_get_collision_exceptions(impl func(ptr unsafe.Pointer, body RID.Any) []RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		ptr, ok := pointers.End(gd.InternalArray(gd.ArrayFromSlice[Array.Contains[RID.Any]](ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _soft_body_set_state(impl func(ptr unsafe.Pointer, body RID.Any, state gdclass.PhysicsServer3DBodyState, v any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var state = gd.UnsafeGet[gdclass.PhysicsServer3DBodyState](p_args, 1)

		var v = variant.Through(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 2))))
		defer pointers.End(gd.InternalVariant(v))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, state, v.Interface())
	}
}
func (Instance) _soft_body_get_state(impl func(ptr unsafe.Pointer, body RID.Any, state gdclass.PhysicsServer3DBodyState) any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var state = gd.UnsafeGet[gdclass.PhysicsServer3DBodyState](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, state)
		ptr, ok := pointers.End(gd.InternalVariant(variant.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _soft_body_set_transform(impl func(ptr unsafe.Pointer, body RID.Any, transform Transform3D.BasisOrigin)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var transform = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, transform)
	}
}
func (Instance) _soft_body_set_simulation_precision(impl func(ptr unsafe.Pointer, body RID.Any, simulation_precision int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var simulation_precision = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(simulation_precision))
	}
}
func (Instance) _soft_body_get_simulation_precision(impl func(ptr unsafe.Pointer, body RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, int64(ret))
	}
}
func (Instance) _soft_body_set_total_mass(impl func(ptr unsafe.Pointer, body RID.Any, total_mass Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var total_mass = gd.UnsafeGet[float64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, Float.X(total_mass))
	}
}
func (Instance) _soft_body_get_total_mass(impl func(ptr unsafe.Pointer, body RID.Any) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, float64(ret))
	}
}
func (Instance) _soft_body_set_linear_stiffness(impl func(ptr unsafe.Pointer, body RID.Any, linear_stiffness Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var linear_stiffness = gd.UnsafeGet[float64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, Float.X(linear_stiffness))
	}
}
func (Instance) _soft_body_get_linear_stiffness(impl func(ptr unsafe.Pointer, body RID.Any) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, float64(ret))
	}
}
func (Instance) _soft_body_set_pressure_coefficient(impl func(ptr unsafe.Pointer, body RID.Any, pressure_coefficient Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var pressure_coefficient = gd.UnsafeGet[float64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, Float.X(pressure_coefficient))
	}
}
func (Instance) _soft_body_get_pressure_coefficient(impl func(ptr unsafe.Pointer, body RID.Any) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, float64(ret))
	}
}
func (Instance) _soft_body_set_damping_coefficient(impl func(ptr unsafe.Pointer, body RID.Any, damping_coefficient Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var damping_coefficient = gd.UnsafeGet[float64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, Float.X(damping_coefficient))
	}
}
func (Instance) _soft_body_get_damping_coefficient(impl func(ptr unsafe.Pointer, body RID.Any) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, float64(ret))
	}
}
func (Instance) _soft_body_set_drag_coefficient(impl func(ptr unsafe.Pointer, body RID.Any, drag_coefficient Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var drag_coefficient = gd.UnsafeGet[float64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, Float.X(drag_coefficient))
	}
}
func (Instance) _soft_body_get_drag_coefficient(impl func(ptr unsafe.Pointer, body RID.Any) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, float64(ret))
	}
}
func (Instance) _soft_body_set_mesh(impl func(ptr unsafe.Pointer, body RID.Any, mesh RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var mesh = gd.UnsafeGet[RID.Any](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, mesh)
	}
}
func (Instance) _soft_body_get_bounds(impl func(ptr unsafe.Pointer, body RID.Any) AABB.PositionSize) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, AABB.PositionSize(ret))
	}
}
func (Instance) _soft_body_move_point(impl func(ptr unsafe.Pointer, body RID.Any, point_index int, global_position Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var point_index = gd.UnsafeGet[int64](p_args, 1)

		var global_position = gd.UnsafeGet[Vector3.XYZ](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(point_index), global_position)
	}
}
func (Instance) _soft_body_get_point_global_position(impl func(ptr unsafe.Pointer, body RID.Any, point_index int) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var point_index = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, int(point_index))
		gd.UnsafeSet(p_back, Vector3.XYZ(ret))
	}
}
func (Instance) _soft_body_remove_all_pinned_points(impl func(ptr unsafe.Pointer, body RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body)
	}
}
func (Instance) _soft_body_pin_point(impl func(ptr unsafe.Pointer, body RID.Any, point_index int, pin bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var point_index = gd.UnsafeGet[int64](p_args, 1)

		var pin = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, int(point_index), pin)
	}
}
func (Instance) _soft_body_is_point_pinned(impl func(ptr unsafe.Pointer, body RID.Any, point_index int) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var point_index = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, int(point_index))
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _joint_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, RID.Any(ret))
	}
}
func (Instance) _joint_clear(impl func(ptr unsafe.Pointer, joint RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint)
	}
}
func (Instance) _joint_make_pin(impl func(ptr unsafe.Pointer, joint RID.Any, body_A RID.Any, local_A Vector3.XYZ, body_B RID.Any, local_B Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var body_A = gd.UnsafeGet[RID.Any](p_args, 1)

		var local_A = gd.UnsafeGet[Vector3.XYZ](p_args, 2)

		var body_B = gd.UnsafeGet[RID.Any](p_args, 3)

		var local_B = gd.UnsafeGet[Vector3.XYZ](p_args, 4)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, local_A, body_B, local_B)
	}
}
func (Instance) _pin_joint_set_param(impl func(ptr unsafe.Pointer, joint RID.Any, param gdclass.PhysicsServer3DPinJointParam, value Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DPinJointParam](p_args, 1)

		var value = gd.UnsafeGet[float64](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, Float.X(value))
	}
}
func (Instance) _pin_joint_get_param(impl func(ptr unsafe.Pointer, joint RID.Any, param gdclass.PhysicsServer3DPinJointParam) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DPinJointParam](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, float64(ret))
	}
}
func (Instance) _pin_joint_set_local_a(impl func(ptr unsafe.Pointer, joint RID.Any, local_A Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var local_A = gd.UnsafeGet[Vector3.XYZ](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, local_A)
	}
}
func (Instance) _pin_joint_get_local_a(impl func(ptr unsafe.Pointer, joint RID.Any) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, Vector3.XYZ(ret))
	}
}
func (Instance) _pin_joint_set_local_b(impl func(ptr unsafe.Pointer, joint RID.Any, local_B Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var local_B = gd.UnsafeGet[Vector3.XYZ](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, local_B)
	}
}
func (Instance) _pin_joint_get_local_b(impl func(ptr unsafe.Pointer, joint RID.Any) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, Vector3.XYZ(ret))
	}
}
func (Instance) _joint_make_hinge(impl func(ptr unsafe.Pointer, joint RID.Any, body_A RID.Any, hinge_A Transform3D.BasisOrigin, body_B RID.Any, hinge_B Transform3D.BasisOrigin)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var body_A = gd.UnsafeGet[RID.Any](p_args, 1)

		var hinge_A = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 2)

		var body_B = gd.UnsafeGet[RID.Any](p_args, 3)

		var hinge_B = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 4)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, hinge_A, body_B, hinge_B)
	}
}
func (Instance) _joint_make_hinge_simple(impl func(ptr unsafe.Pointer, joint RID.Any, body_A RID.Any, pivot_A Vector3.XYZ, axis_A Vector3.XYZ, body_B RID.Any, pivot_B Vector3.XYZ, axis_B Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var body_A = gd.UnsafeGet[RID.Any](p_args, 1)

		var pivot_A = gd.UnsafeGet[Vector3.XYZ](p_args, 2)

		var axis_A = gd.UnsafeGet[Vector3.XYZ](p_args, 3)

		var body_B = gd.UnsafeGet[RID.Any](p_args, 4)

		var pivot_B = gd.UnsafeGet[Vector3.XYZ](p_args, 5)

		var axis_B = gd.UnsafeGet[Vector3.XYZ](p_args, 6)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, pivot_A, axis_A, body_B, pivot_B, axis_B)
	}
}
func (Instance) _hinge_joint_set_param(impl func(ptr unsafe.Pointer, joint RID.Any, param gdclass.PhysicsServer3DHingeJointParam, value Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DHingeJointParam](p_args, 1)

		var value = gd.UnsafeGet[float64](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, Float.X(value))
	}
}
func (Instance) _hinge_joint_get_param(impl func(ptr unsafe.Pointer, joint RID.Any, param gdclass.PhysicsServer3DHingeJointParam) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DHingeJointParam](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, float64(ret))
	}
}
func (Instance) _hinge_joint_set_flag(impl func(ptr unsafe.Pointer, joint RID.Any, flag gdclass.PhysicsServer3DHingeJointFlag, enabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var flag = gd.UnsafeGet[gdclass.PhysicsServer3DHingeJointFlag](p_args, 1)

		var enabled = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, flag, enabled)
	}
}
func (Instance) _hinge_joint_get_flag(impl func(ptr unsafe.Pointer, joint RID.Any, flag gdclass.PhysicsServer3DHingeJointFlag) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var flag = gd.UnsafeGet[gdclass.PhysicsServer3DHingeJointFlag](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, flag)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _joint_make_slider(impl func(ptr unsafe.Pointer, joint RID.Any, body_A RID.Any, local_ref_A Transform3D.BasisOrigin, body_B RID.Any, local_ref_B Transform3D.BasisOrigin)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var body_A = gd.UnsafeGet[RID.Any](p_args, 1)

		var local_ref_A = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 2)

		var body_B = gd.UnsafeGet[RID.Any](p_args, 3)

		var local_ref_B = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 4)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, local_ref_A, body_B, local_ref_B)
	}
}
func (Instance) _slider_joint_set_param(impl func(ptr unsafe.Pointer, joint RID.Any, param gdclass.PhysicsServer3DSliderJointParam, value Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DSliderJointParam](p_args, 1)

		var value = gd.UnsafeGet[float64](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, Float.X(value))
	}
}
func (Instance) _slider_joint_get_param(impl func(ptr unsafe.Pointer, joint RID.Any, param gdclass.PhysicsServer3DSliderJointParam) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DSliderJointParam](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, float64(ret))
	}
}
func (Instance) _joint_make_cone_twist(impl func(ptr unsafe.Pointer, joint RID.Any, body_A RID.Any, local_ref_A Transform3D.BasisOrigin, body_B RID.Any, local_ref_B Transform3D.BasisOrigin)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var body_A = gd.UnsafeGet[RID.Any](p_args, 1)

		var local_ref_A = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 2)

		var body_B = gd.UnsafeGet[RID.Any](p_args, 3)

		var local_ref_B = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 4)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, local_ref_A, body_B, local_ref_B)
	}
}
func (Instance) _cone_twist_joint_set_param(impl func(ptr unsafe.Pointer, joint RID.Any, param gdclass.PhysicsServer3DConeTwistJointParam, value Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DConeTwistJointParam](p_args, 1)

		var value = gd.UnsafeGet[float64](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, Float.X(value))
	}
}
func (Instance) _cone_twist_joint_get_param(impl func(ptr unsafe.Pointer, joint RID.Any, param gdclass.PhysicsServer3DConeTwistJointParam) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DConeTwistJointParam](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, float64(ret))
	}
}
func (Instance) _joint_make_generic_6dof(impl func(ptr unsafe.Pointer, joint RID.Any, body_A RID.Any, local_ref_A Transform3D.BasisOrigin, body_B RID.Any, local_ref_B Transform3D.BasisOrigin)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var body_A = gd.UnsafeGet[RID.Any](p_args, 1)

		var local_ref_A = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 2)

		var body_B = gd.UnsafeGet[RID.Any](p_args, 3)

		var local_ref_B = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 4)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, local_ref_A, body_B, local_ref_B)
	}
}
func (Instance) _generic_6dof_joint_set_param(impl func(ptr unsafe.Pointer, joint RID.Any, axis gd.Vector3Axis, param gdclass.PhysicsServer3DG6DOFJointAxisParam, value Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var axis = gd.UnsafeGet[gd.Vector3Axis](p_args, 1)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DG6DOFJointAxisParam](p_args, 2)

		var value = gd.UnsafeGet[float64](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, axis, param, Float.X(value))
	}
}
func (Instance) _generic_6dof_joint_get_param(impl func(ptr unsafe.Pointer, joint RID.Any, axis gd.Vector3Axis, param gdclass.PhysicsServer3DG6DOFJointAxisParam) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var axis = gd.UnsafeGet[gd.Vector3Axis](p_args, 1)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DG6DOFJointAxisParam](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, axis, param)
		gd.UnsafeSet(p_back, float64(ret))
	}
}
func (Instance) _generic_6dof_joint_set_flag(impl func(ptr unsafe.Pointer, joint RID.Any, axis gd.Vector3Axis, flag gdclass.PhysicsServer3DG6DOFJointAxisFlag, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var axis = gd.UnsafeGet[gd.Vector3Axis](p_args, 1)

		var flag = gd.UnsafeGet[gdclass.PhysicsServer3DG6DOFJointAxisFlag](p_args, 2)

		var enable = gd.UnsafeGet[bool](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, axis, flag, enable)
	}
}
func (Instance) _generic_6dof_joint_get_flag(impl func(ptr unsafe.Pointer, joint RID.Any, axis gd.Vector3Axis, flag gdclass.PhysicsServer3DG6DOFJointAxisFlag) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var axis = gd.UnsafeGet[gd.Vector3Axis](p_args, 1)

		var flag = gd.UnsafeGet[gdclass.PhysicsServer3DG6DOFJointAxisFlag](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, axis, flag)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _joint_get_type(impl func(ptr unsafe.Pointer, joint RID.Any) gdclass.PhysicsServer3DJointType) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _joint_set_solver_priority(impl func(ptr unsafe.Pointer, joint RID.Any, priority int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var priority = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, int(priority))
	}
}
func (Instance) _joint_get_solver_priority(impl func(ptr unsafe.Pointer, joint RID.Any) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, int64(ret))
	}
}
func (Instance) _joint_disable_collisions_between_bodies(impl func(ptr unsafe.Pointer, joint RID.Any, disable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var disable = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, disable)
	}
}
func (Instance) _joint_is_disabled_collisions_between_bodies(impl func(ptr unsafe.Pointer, joint RID.Any) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _free_rid(impl func(ptr unsafe.Pointer, rid RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var rid = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, rid)
	}
}
func (Instance) _set_active(impl func(ptr unsafe.Pointer, active bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var active = gd.UnsafeGet[bool](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, active)
	}
}
func (Instance) _init(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}
func (Instance) _step(impl func(ptr unsafe.Pointer, step Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var step = gd.UnsafeGet[float64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, Float.X(step))
	}
}
func (Instance) _sync(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}
func (Instance) _flush_queries(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}
func (Instance) _end_sync(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}
func (Instance) _finish(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}
func (Instance) _is_flushing_queries(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_process_info(impl func(ptr unsafe.Pointer, process_info gdclass.PhysicsServer3DProcessInfo) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var process_info = gd.UnsafeGet[gdclass.PhysicsServer3DProcessInfo](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, process_info)
		gd.UnsafeSet(p_back, int64(ret))
	}
}
func (self Instance) BodyTestMotionIsExcludingBody(body RID.Body3D) bool { //gd:PhysicsServer3DExtension.body_test_motion_is_excluding_body
	return bool(class(self).BodyTestMotionIsExcludingBody(RID.Any(body)))
}
func (self Instance) BodyTestMotionIsExcludingObject(obj int) bool { //gd:PhysicsServer3DExtension.body_test_motion_is_excluding_object
	return bool(class(self).BodyTestMotionIsExcludingObject(int64(obj)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PhysicsServer3DExtension

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicsServer3DExtension"))
	casted := Instance{*(*gdclass.PhysicsServer3DExtension)(unsafe.Pointer(&object))}
	return casted
}

func (class) _world_boundary_shape_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _separation_ray_shape_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _sphere_shape_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _box_shape_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _capsule_shape_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _cylinder_shape_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _convex_polygon_shape_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _concave_polygon_shape_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _heightmap_shape_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _custom_shape_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _shape_set_data(impl func(ptr unsafe.Pointer, shape RID.Any, data variant.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape = gd.UnsafeGet[RID.Any](p_args, 0)

		var data = variant.Through(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 1))))
		defer pointers.End(gd.InternalVariant(data))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shape, data)
	}
}

func (class) _shape_set_custom_solver_bias(impl func(ptr unsafe.Pointer, shape RID.Any, bias float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape = gd.UnsafeGet[RID.Any](p_args, 0)

		var bias = gd.UnsafeGet[float64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shape, bias)
	}
}

func (class) _shape_set_margin(impl func(ptr unsafe.Pointer, shape RID.Any, margin float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape = gd.UnsafeGet[RID.Any](p_args, 0)

		var margin = gd.UnsafeGet[float64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, shape, margin)
	}
}

func (class) _shape_get_margin(impl func(ptr unsafe.Pointer, shape RID.Any) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _shape_get_type(impl func(ptr unsafe.Pointer, shape RID.Any) gdclass.PhysicsServer3DShapeType) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _shape_get_data(impl func(ptr unsafe.Pointer, shape RID.Any) variant.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		ptr, ok := pointers.End(gd.InternalVariant(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _shape_get_custom_solver_bias(impl func(ptr unsafe.Pointer, shape RID.Any) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var shape = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, shape)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _space_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _space_set_active(impl func(ptr unsafe.Pointer, space RID.Any, active bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[RID.Any](p_args, 0)

		var active = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, space, active)
	}
}

func (class) _space_is_active(impl func(ptr unsafe.Pointer, space RID.Any) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _space_set_param(impl func(ptr unsafe.Pointer, space RID.Any, param gdclass.PhysicsServer3DSpaceParameter, value float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[RID.Any](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DSpaceParameter](p_args, 1)

		var value = gd.UnsafeGet[float64](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, space, param, value)
	}
}

func (class) _space_get_param(impl func(ptr unsafe.Pointer, space RID.Any, param gdclass.PhysicsServer3DSpaceParameter) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[RID.Any](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DSpaceParameter](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space, param)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _space_get_direct_state(impl func(ptr unsafe.Pointer, space RID.Any) [1]gdclass.PhysicsDirectSpaceState3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _space_set_debug_contacts(impl func(ptr unsafe.Pointer, space RID.Any, max_contacts int64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[RID.Any](p_args, 0)

		var max_contacts = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, space, max_contacts)
	}
}

func (class) _space_get_contacts(impl func(ptr unsafe.Pointer, space RID.Any) Packed.Array[Vector3.XYZ]) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		ptr, ok := pointers.End(gd.InternalPacked[gd.PackedVector3Array, Vector3.XYZ](ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _space_get_contact_count(impl func(ptr unsafe.Pointer, space RID.Any) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var space = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, space)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_set_space(impl func(ptr unsafe.Pointer, area RID.Any, space RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var space = gd.UnsafeGet[RID.Any](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, space)
	}
}

func (class) _area_get_space(impl func(ptr unsafe.Pointer, area RID.Any) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_add_shape(impl func(ptr unsafe.Pointer, area RID.Any, shape RID.Any, transform Transform3D.BasisOrigin, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var shape = gd.UnsafeGet[RID.Any](p_args, 1)

		var transform = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 2)

		var disabled = gd.UnsafeGet[bool](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, shape, transform, disabled)
	}
}

func (class) _area_set_shape(impl func(ptr unsafe.Pointer, area RID.Any, shape_idx int64, shape RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var shape_idx = gd.UnsafeGet[int64](p_args, 1)

		var shape = gd.UnsafeGet[RID.Any](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, shape_idx, shape)
	}
}

func (class) _area_set_shape_transform(impl func(ptr unsafe.Pointer, area RID.Any, shape_idx int64, transform Transform3D.BasisOrigin)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var shape_idx = gd.UnsafeGet[int64](p_args, 1)

		var transform = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, shape_idx, transform)
	}
}

func (class) _area_set_shape_disabled(impl func(ptr unsafe.Pointer, area RID.Any, shape_idx int64, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var shape_idx = gd.UnsafeGet[int64](p_args, 1)

		var disabled = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, shape_idx, disabled)
	}
}

func (class) _area_get_shape_count(impl func(ptr unsafe.Pointer, area RID.Any) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_get_shape(impl func(ptr unsafe.Pointer, area RID.Any, shape_idx int64) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var shape_idx = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, shape_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_get_shape_transform(impl func(ptr unsafe.Pointer, area RID.Any, shape_idx int64) Transform3D.BasisOrigin) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var shape_idx = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, shape_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_remove_shape(impl func(ptr unsafe.Pointer, area RID.Any, shape_idx int64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var shape_idx = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, shape_idx)
	}
}

func (class) _area_clear_shapes(impl func(ptr unsafe.Pointer, area RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area)
	}
}

func (class) _area_attach_object_instance_id(impl func(ptr unsafe.Pointer, area RID.Any, id int64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var id = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, id)
	}
}

func (class) _area_get_object_instance_id(impl func(ptr unsafe.Pointer, area RID.Any) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_set_param(impl func(ptr unsafe.Pointer, area RID.Any, param gdclass.PhysicsServer3DAreaParameter, value variant.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DAreaParameter](p_args, 1)

		var value = variant.Through(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 2))))
		defer pointers.End(gd.InternalVariant(value))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, param, value)
	}
}

func (class) _area_set_transform(impl func(ptr unsafe.Pointer, area RID.Any, transform Transform3D.BasisOrigin)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var transform = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, transform)
	}
}

func (class) _area_get_param(impl func(ptr unsafe.Pointer, area RID.Any, param gdclass.PhysicsServer3DAreaParameter) variant.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DAreaParameter](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area, param)
		ptr, ok := pointers.End(gd.InternalVariant(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _area_get_transform(impl func(ptr unsafe.Pointer, area RID.Any) Transform3D.BasisOrigin) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_set_collision_layer(impl func(ptr unsafe.Pointer, area RID.Any, layer int64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var layer = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, layer)
	}
}

func (class) _area_get_collision_layer(impl func(ptr unsafe.Pointer, area RID.Any) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_set_collision_mask(impl func(ptr unsafe.Pointer, area RID.Any, mask int64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var mask = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, mask)
	}
}

func (class) _area_get_collision_mask(impl func(ptr unsafe.Pointer, area RID.Any) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, area)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _area_set_monitorable(impl func(ptr unsafe.Pointer, area RID.Any, monitorable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var monitorable = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, monitorable)
	}
}

func (class) _area_set_ray_pickable(impl func(ptr unsafe.Pointer, area RID.Any, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var enable = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, enable)
	}
}

func (class) _area_set_monitor_callback(impl func(ptr unsafe.Pointer, area RID.Any, callback Callable.Function)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var callback = Callable.Through(gd.CallableProxy{}, pointers.Pack(pointers.New[gd.Callable](gd.UnsafeGet[[2]uint64](p_args, 1))))
		defer pointers.End(gd.InternalCallable(callback))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, callback)
	}
}

func (class) _area_set_area_monitor_callback(impl func(ptr unsafe.Pointer, area RID.Any, callback Callable.Function)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var area = gd.UnsafeGet[RID.Any](p_args, 0)

		var callback = Callable.Through(gd.CallableProxy{}, pointers.Pack(pointers.New[gd.Callable](gd.UnsafeGet[[2]uint64](p_args, 1))))
		defer pointers.End(gd.InternalCallable(callback))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, area, callback)
	}
}

func (class) _body_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_space(impl func(ptr unsafe.Pointer, body RID.Any, space RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var space = gd.UnsafeGet[RID.Any](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, space)
	}
}

func (class) _body_get_space(impl func(ptr unsafe.Pointer, body RID.Any) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_mode(impl func(ptr unsafe.Pointer, body RID.Any, mode gdclass.PhysicsServer3DBodyMode)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var mode = gd.UnsafeGet[gdclass.PhysicsServer3DBodyMode](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, mode)
	}
}

func (class) _body_get_mode(impl func(ptr unsafe.Pointer, body RID.Any) gdclass.PhysicsServer3DBodyMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_add_shape(impl func(ptr unsafe.Pointer, body RID.Any, shape RID.Any, transform Transform3D.BasisOrigin, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var shape = gd.UnsafeGet[RID.Any](p_args, 1)

		var transform = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 2)

		var disabled = gd.UnsafeGet[bool](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, shape, transform, disabled)
	}
}

func (class) _body_set_shape(impl func(ptr unsafe.Pointer, body RID.Any, shape_idx int64, shape RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var shape_idx = gd.UnsafeGet[int64](p_args, 1)

		var shape = gd.UnsafeGet[RID.Any](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, shape_idx, shape)
	}
}

func (class) _body_set_shape_transform(impl func(ptr unsafe.Pointer, body RID.Any, shape_idx int64, transform Transform3D.BasisOrigin)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var shape_idx = gd.UnsafeGet[int64](p_args, 1)

		var transform = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, shape_idx, transform)
	}
}

func (class) _body_set_shape_disabled(impl func(ptr unsafe.Pointer, body RID.Any, shape_idx int64, disabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var shape_idx = gd.UnsafeGet[int64](p_args, 1)

		var disabled = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, shape_idx, disabled)
	}
}

func (class) _body_get_shape_count(impl func(ptr unsafe.Pointer, body RID.Any) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_get_shape(impl func(ptr unsafe.Pointer, body RID.Any, shape_idx int64) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var shape_idx = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, shape_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_get_shape_transform(impl func(ptr unsafe.Pointer, body RID.Any, shape_idx int64) Transform3D.BasisOrigin) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var shape_idx = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, shape_idx)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_remove_shape(impl func(ptr unsafe.Pointer, body RID.Any, shape_idx int64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var shape_idx = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, shape_idx)
	}
}

func (class) _body_clear_shapes(impl func(ptr unsafe.Pointer, body RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body)
	}
}

func (class) _body_attach_object_instance_id(impl func(ptr unsafe.Pointer, body RID.Any, id int64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var id = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, id)
	}
}

func (class) _body_get_object_instance_id(impl func(ptr unsafe.Pointer, body RID.Any) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_enable_continuous_collision_detection(impl func(ptr unsafe.Pointer, body RID.Any, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var enable = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, enable)
	}
}

func (class) _body_is_continuous_collision_detection_enabled(impl func(ptr unsafe.Pointer, body RID.Any) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_collision_layer(impl func(ptr unsafe.Pointer, body RID.Any, layer int64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var layer = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, layer)
	}
}

func (class) _body_get_collision_layer(impl func(ptr unsafe.Pointer, body RID.Any) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_collision_mask(impl func(ptr unsafe.Pointer, body RID.Any, mask int64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var mask = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, mask)
	}
}

func (class) _body_get_collision_mask(impl func(ptr unsafe.Pointer, body RID.Any) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_collision_priority(impl func(ptr unsafe.Pointer, body RID.Any, priority float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var priority = gd.UnsafeGet[float64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, priority)
	}
}

func (class) _body_get_collision_priority(impl func(ptr unsafe.Pointer, body RID.Any) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_user_flags(impl func(ptr unsafe.Pointer, body RID.Any, flags int64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var flags = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, flags)
	}
}

func (class) _body_get_user_flags(impl func(ptr unsafe.Pointer, body RID.Any) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_param(impl func(ptr unsafe.Pointer, body RID.Any, param gdclass.PhysicsServer3DBodyParameter, value variant.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DBodyParameter](p_args, 1)

		var value = variant.Through(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 2))))
		defer pointers.End(gd.InternalVariant(value))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, param, value)
	}
}

func (class) _body_get_param(impl func(ptr unsafe.Pointer, body RID.Any, param gdclass.PhysicsServer3DBodyParameter) variant.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DBodyParameter](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, param)
		ptr, ok := pointers.End(gd.InternalVariant(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _body_reset_mass_properties(impl func(ptr unsafe.Pointer, body RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body)
	}
}

func (class) _body_set_state(impl func(ptr unsafe.Pointer, body RID.Any, state gdclass.PhysicsServer3DBodyState, value variant.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var state = gd.UnsafeGet[gdclass.PhysicsServer3DBodyState](p_args, 1)

		var value = variant.Through(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 2))))
		defer pointers.End(gd.InternalVariant(value))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, state, value)
	}
}

func (class) _body_get_state(impl func(ptr unsafe.Pointer, body RID.Any, state gdclass.PhysicsServer3DBodyState) variant.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var state = gd.UnsafeGet[gdclass.PhysicsServer3DBodyState](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, state)
		ptr, ok := pointers.End(gd.InternalVariant(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _body_apply_central_impulse(impl func(ptr unsafe.Pointer, body RID.Any, impulse Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var impulse = gd.UnsafeGet[Vector3.XYZ](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, impulse)
	}
}

func (class) _body_apply_impulse(impl func(ptr unsafe.Pointer, body RID.Any, impulse Vector3.XYZ, position Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var impulse = gd.UnsafeGet[Vector3.XYZ](p_args, 1)

		var position = gd.UnsafeGet[Vector3.XYZ](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, impulse, position)
	}
}

func (class) _body_apply_torque_impulse(impl func(ptr unsafe.Pointer, body RID.Any, impulse Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var impulse = gd.UnsafeGet[Vector3.XYZ](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, impulse)
	}
}

func (class) _body_apply_central_force(impl func(ptr unsafe.Pointer, body RID.Any, force Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var force = gd.UnsafeGet[Vector3.XYZ](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force)
	}
}

func (class) _body_apply_force(impl func(ptr unsafe.Pointer, body RID.Any, force Vector3.XYZ, position Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var force = gd.UnsafeGet[Vector3.XYZ](p_args, 1)

		var position = gd.UnsafeGet[Vector3.XYZ](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force, position)
	}
}

func (class) _body_apply_torque(impl func(ptr unsafe.Pointer, body RID.Any, torque Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var torque = gd.UnsafeGet[Vector3.XYZ](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, torque)
	}
}

func (class) _body_add_constant_central_force(impl func(ptr unsafe.Pointer, body RID.Any, force Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var force = gd.UnsafeGet[Vector3.XYZ](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force)
	}
}

func (class) _body_add_constant_force(impl func(ptr unsafe.Pointer, body RID.Any, force Vector3.XYZ, position Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var force = gd.UnsafeGet[Vector3.XYZ](p_args, 1)

		var position = gd.UnsafeGet[Vector3.XYZ](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force, position)
	}
}

func (class) _body_add_constant_torque(impl func(ptr unsafe.Pointer, body RID.Any, torque Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var torque = gd.UnsafeGet[Vector3.XYZ](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, torque)
	}
}

func (class) _body_set_constant_force(impl func(ptr unsafe.Pointer, body RID.Any, force Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var force = gd.UnsafeGet[Vector3.XYZ](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, force)
	}
}

func (class) _body_get_constant_force(impl func(ptr unsafe.Pointer, body RID.Any) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_constant_torque(impl func(ptr unsafe.Pointer, body RID.Any, torque Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var torque = gd.UnsafeGet[Vector3.XYZ](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, torque)
	}
}

func (class) _body_get_constant_torque(impl func(ptr unsafe.Pointer, body RID.Any) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_axis_velocity(impl func(ptr unsafe.Pointer, body RID.Any, axis_velocity Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var axis_velocity = gd.UnsafeGet[Vector3.XYZ](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, axis_velocity)
	}
}

func (class) _body_set_axis_lock(impl func(ptr unsafe.Pointer, body RID.Any, axis gdclass.PhysicsServer3DBodyAxis, lock bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var axis = gd.UnsafeGet[gdclass.PhysicsServer3DBodyAxis](p_args, 1)

		var lock = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, axis, lock)
	}
}

func (class) _body_is_axis_locked(impl func(ptr unsafe.Pointer, body RID.Any, axis gdclass.PhysicsServer3DBodyAxis) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var axis = gd.UnsafeGet[gdclass.PhysicsServer3DBodyAxis](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, axis)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_add_collision_exception(impl func(ptr unsafe.Pointer, body RID.Any, excepted_body RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var excepted_body = gd.UnsafeGet[RID.Any](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, excepted_body)
	}
}

func (class) _body_remove_collision_exception(impl func(ptr unsafe.Pointer, body RID.Any, excepted_body RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var excepted_body = gd.UnsafeGet[RID.Any](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, excepted_body)
	}
}

func (class) _body_get_collision_exceptions(impl func(ptr unsafe.Pointer, body RID.Any) Array.Contains[RID.Any]) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		ptr, ok := pointers.End(gd.InternalArray(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _body_set_max_contacts_reported(impl func(ptr unsafe.Pointer, body RID.Any, amount int64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var amount = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, amount)
	}
}

func (class) _body_get_max_contacts_reported(impl func(ptr unsafe.Pointer, body RID.Any) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_contacts_reported_depth_threshold(impl func(ptr unsafe.Pointer, body RID.Any, threshold float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var threshold = gd.UnsafeGet[float64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, threshold)
	}
}

func (class) _body_get_contacts_reported_depth_threshold(impl func(ptr unsafe.Pointer, body RID.Any) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_omit_force_integration(impl func(ptr unsafe.Pointer, body RID.Any, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var enable = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, enable)
	}
}

func (class) _body_is_omitting_force_integration(impl func(ptr unsafe.Pointer, body RID.Any) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_set_state_sync_callback(impl func(ptr unsafe.Pointer, body RID.Any, callable Callable.Function)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var callable = Callable.Through(gd.CallableProxy{}, pointers.Pack(pointers.New[gd.Callable](gd.UnsafeGet[[2]uint64](p_args, 1))))
		defer pointers.End(gd.InternalCallable(callable))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, callable)
	}
}

func (class) _body_set_force_integration_callback(impl func(ptr unsafe.Pointer, body RID.Any, callable Callable.Function, userdata variant.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var callable = Callable.Through(gd.CallableProxy{}, pointers.Pack(pointers.New[gd.Callable](gd.UnsafeGet[[2]uint64](p_args, 1))))
		defer pointers.End(gd.InternalCallable(callable))
		var userdata = variant.Through(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 2))))
		defer pointers.End(gd.InternalVariant(userdata))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, callable, userdata)
	}
}

func (class) _body_set_ray_pickable(impl func(ptr unsafe.Pointer, body RID.Any, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var enable = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, enable)
	}
}

func (class) _body_test_motion(impl func(ptr unsafe.Pointer, body RID.Any, from Transform3D.BasisOrigin, motion Vector3.XYZ, margin float64, max_collisions int64, collide_separation_ray bool, recovery_as_collision bool, result *MotionResult) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var from = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 1)

		var motion = gd.UnsafeGet[Vector3.XYZ](p_args, 2)

		var margin = gd.UnsafeGet[float64](p_args, 3)

		var max_collisions = gd.UnsafeGet[int64](p_args, 4)

		var collide_separation_ray = gd.UnsafeGet[bool](p_args, 5)

		var recovery_as_collision = gd.UnsafeGet[bool](p_args, 6)

		var result = gd.UnsafeGet[*MotionResult](p_args, 7)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, from, motion, margin, max_collisions, collide_separation_ray, recovery_as_collision, result)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _body_get_direct_state(impl func(ptr unsafe.Pointer, body RID.Any) [1]gdclass.PhysicsDirectBodyState3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _soft_body_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_update_rendering_server(impl func(ptr unsafe.Pointer, body RID.Any, rendering_server_handler [1]gdclass.PhysicsServer3DRenderingServerHandler)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var rendering_server_handler = [1]gdclass.PhysicsServer3DRenderingServerHandler{pointers.New[gdclass.PhysicsServer3DRenderingServerHandler]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}

		defer pointers.End(rendering_server_handler[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, rendering_server_handler)
	}
}

func (class) _soft_body_set_space(impl func(ptr unsafe.Pointer, body RID.Any, space RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var space = gd.UnsafeGet[RID.Any](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, space)
	}
}

func (class) _soft_body_get_space(impl func(ptr unsafe.Pointer, body RID.Any) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_set_ray_pickable(impl func(ptr unsafe.Pointer, body RID.Any, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var enable = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, enable)
	}
}

func (class) _soft_body_set_collision_layer(impl func(ptr unsafe.Pointer, body RID.Any, layer int64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var layer = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, layer)
	}
}

func (class) _soft_body_get_collision_layer(impl func(ptr unsafe.Pointer, body RID.Any) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_set_collision_mask(impl func(ptr unsafe.Pointer, body RID.Any, mask int64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var mask = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, mask)
	}
}

func (class) _soft_body_get_collision_mask(impl func(ptr unsafe.Pointer, body RID.Any) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_add_collision_exception(impl func(ptr unsafe.Pointer, body RID.Any, body_b RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var body_b = gd.UnsafeGet[RID.Any](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, body_b)
	}
}

func (class) _soft_body_remove_collision_exception(impl func(ptr unsafe.Pointer, body RID.Any, body_b RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var body_b = gd.UnsafeGet[RID.Any](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, body_b)
	}
}

func (class) _soft_body_get_collision_exceptions(impl func(ptr unsafe.Pointer, body RID.Any) Array.Contains[RID.Any]) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		ptr, ok := pointers.End(gd.InternalArray(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _soft_body_set_state(impl func(ptr unsafe.Pointer, body RID.Any, state gdclass.PhysicsServer3DBodyState, v variant.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var state = gd.UnsafeGet[gdclass.PhysicsServer3DBodyState](p_args, 1)

		var v = variant.Through(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 2))))
		defer pointers.End(gd.InternalVariant(v))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, state, v)
	}
}

func (class) _soft_body_get_state(impl func(ptr unsafe.Pointer, body RID.Any, state gdclass.PhysicsServer3DBodyState) variant.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var state = gd.UnsafeGet[gdclass.PhysicsServer3DBodyState](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, state)
		ptr, ok := pointers.End(gd.InternalVariant(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _soft_body_set_transform(impl func(ptr unsafe.Pointer, body RID.Any, transform Transform3D.BasisOrigin)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var transform = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, transform)
	}
}

func (class) _soft_body_set_simulation_precision(impl func(ptr unsafe.Pointer, body RID.Any, simulation_precision int64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var simulation_precision = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, simulation_precision)
	}
}

func (class) _soft_body_get_simulation_precision(impl func(ptr unsafe.Pointer, body RID.Any) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_set_total_mass(impl func(ptr unsafe.Pointer, body RID.Any, total_mass float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var total_mass = gd.UnsafeGet[float64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, total_mass)
	}
}

func (class) _soft_body_get_total_mass(impl func(ptr unsafe.Pointer, body RID.Any) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_set_linear_stiffness(impl func(ptr unsafe.Pointer, body RID.Any, linear_stiffness float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var linear_stiffness = gd.UnsafeGet[float64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, linear_stiffness)
	}
}

func (class) _soft_body_get_linear_stiffness(impl func(ptr unsafe.Pointer, body RID.Any) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_set_pressure_coefficient(impl func(ptr unsafe.Pointer, body RID.Any, pressure_coefficient float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var pressure_coefficient = gd.UnsafeGet[float64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, pressure_coefficient)
	}
}

func (class) _soft_body_get_pressure_coefficient(impl func(ptr unsafe.Pointer, body RID.Any) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_set_damping_coefficient(impl func(ptr unsafe.Pointer, body RID.Any, damping_coefficient float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var damping_coefficient = gd.UnsafeGet[float64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, damping_coefficient)
	}
}

func (class) _soft_body_get_damping_coefficient(impl func(ptr unsafe.Pointer, body RID.Any) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_set_drag_coefficient(impl func(ptr unsafe.Pointer, body RID.Any, drag_coefficient float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var drag_coefficient = gd.UnsafeGet[float64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, drag_coefficient)
	}
}

func (class) _soft_body_get_drag_coefficient(impl func(ptr unsafe.Pointer, body RID.Any) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_set_mesh(impl func(ptr unsafe.Pointer, body RID.Any, mesh RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var mesh = gd.UnsafeGet[RID.Any](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, mesh)
	}
}

func (class) _soft_body_get_bounds(impl func(ptr unsafe.Pointer, body RID.Any) AABB.PositionSize) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_move_point(impl func(ptr unsafe.Pointer, body RID.Any, point_index int64, global_position Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var point_index = gd.UnsafeGet[int64](p_args, 1)

		var global_position = gd.UnsafeGet[Vector3.XYZ](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, point_index, global_position)
	}
}

func (class) _soft_body_get_point_global_position(impl func(ptr unsafe.Pointer, body RID.Any, point_index int64) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var point_index = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, point_index)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _soft_body_remove_all_pinned_points(impl func(ptr unsafe.Pointer, body RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body)
	}
}

func (class) _soft_body_pin_point(impl func(ptr unsafe.Pointer, body RID.Any, point_index int64, pin bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var point_index = gd.UnsafeGet[int64](p_args, 1)

		var pin = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, body, point_index, pin)
	}
}

func (class) _soft_body_is_point_pinned(impl func(ptr unsafe.Pointer, body RID.Any, point_index int64) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var body = gd.UnsafeGet[RID.Any](p_args, 0)

		var point_index = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, body, point_index)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _joint_create(impl func(ptr unsafe.Pointer) RID.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _joint_clear(impl func(ptr unsafe.Pointer, joint RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint)
	}
}

func (class) _joint_make_pin(impl func(ptr unsafe.Pointer, joint RID.Any, body_A RID.Any, local_A Vector3.XYZ, body_B RID.Any, local_B Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var body_A = gd.UnsafeGet[RID.Any](p_args, 1)

		var local_A = gd.UnsafeGet[Vector3.XYZ](p_args, 2)

		var body_B = gd.UnsafeGet[RID.Any](p_args, 3)

		var local_B = gd.UnsafeGet[Vector3.XYZ](p_args, 4)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, local_A, body_B, local_B)
	}
}

func (class) _pin_joint_set_param(impl func(ptr unsafe.Pointer, joint RID.Any, param gdclass.PhysicsServer3DPinJointParam, value float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DPinJointParam](p_args, 1)

		var value = gd.UnsafeGet[float64](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, value)
	}
}

func (class) _pin_joint_get_param(impl func(ptr unsafe.Pointer, joint RID.Any, param gdclass.PhysicsServer3DPinJointParam) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DPinJointParam](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _pin_joint_set_local_a(impl func(ptr unsafe.Pointer, joint RID.Any, local_A Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var local_A = gd.UnsafeGet[Vector3.XYZ](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, local_A)
	}
}

func (class) _pin_joint_get_local_a(impl func(ptr unsafe.Pointer, joint RID.Any) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _pin_joint_set_local_b(impl func(ptr unsafe.Pointer, joint RID.Any, local_B Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var local_B = gd.UnsafeGet[Vector3.XYZ](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, local_B)
	}
}

func (class) _pin_joint_get_local_b(impl func(ptr unsafe.Pointer, joint RID.Any) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _joint_make_hinge(impl func(ptr unsafe.Pointer, joint RID.Any, body_A RID.Any, hinge_A Transform3D.BasisOrigin, body_B RID.Any, hinge_B Transform3D.BasisOrigin)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var body_A = gd.UnsafeGet[RID.Any](p_args, 1)

		var hinge_A = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 2)

		var body_B = gd.UnsafeGet[RID.Any](p_args, 3)

		var hinge_B = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 4)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, hinge_A, body_B, hinge_B)
	}
}

func (class) _joint_make_hinge_simple(impl func(ptr unsafe.Pointer, joint RID.Any, body_A RID.Any, pivot_A Vector3.XYZ, axis_A Vector3.XYZ, body_B RID.Any, pivot_B Vector3.XYZ, axis_B Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var body_A = gd.UnsafeGet[RID.Any](p_args, 1)

		var pivot_A = gd.UnsafeGet[Vector3.XYZ](p_args, 2)

		var axis_A = gd.UnsafeGet[Vector3.XYZ](p_args, 3)

		var body_B = gd.UnsafeGet[RID.Any](p_args, 4)

		var pivot_B = gd.UnsafeGet[Vector3.XYZ](p_args, 5)

		var axis_B = gd.UnsafeGet[Vector3.XYZ](p_args, 6)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, pivot_A, axis_A, body_B, pivot_B, axis_B)
	}
}

func (class) _hinge_joint_set_param(impl func(ptr unsafe.Pointer, joint RID.Any, param gdclass.PhysicsServer3DHingeJointParam, value float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DHingeJointParam](p_args, 1)

		var value = gd.UnsafeGet[float64](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, value)
	}
}

func (class) _hinge_joint_get_param(impl func(ptr unsafe.Pointer, joint RID.Any, param gdclass.PhysicsServer3DHingeJointParam) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DHingeJointParam](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _hinge_joint_set_flag(impl func(ptr unsafe.Pointer, joint RID.Any, flag gdclass.PhysicsServer3DHingeJointFlag, enabled bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var flag = gd.UnsafeGet[gdclass.PhysicsServer3DHingeJointFlag](p_args, 1)

		var enabled = gd.UnsafeGet[bool](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, flag, enabled)
	}
}

func (class) _hinge_joint_get_flag(impl func(ptr unsafe.Pointer, joint RID.Any, flag gdclass.PhysicsServer3DHingeJointFlag) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var flag = gd.UnsafeGet[gdclass.PhysicsServer3DHingeJointFlag](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, flag)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _joint_make_slider(impl func(ptr unsafe.Pointer, joint RID.Any, body_A RID.Any, local_ref_A Transform3D.BasisOrigin, body_B RID.Any, local_ref_B Transform3D.BasisOrigin)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var body_A = gd.UnsafeGet[RID.Any](p_args, 1)

		var local_ref_A = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 2)

		var body_B = gd.UnsafeGet[RID.Any](p_args, 3)

		var local_ref_B = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 4)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, local_ref_A, body_B, local_ref_B)
	}
}

func (class) _slider_joint_set_param(impl func(ptr unsafe.Pointer, joint RID.Any, param gdclass.PhysicsServer3DSliderJointParam, value float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DSliderJointParam](p_args, 1)

		var value = gd.UnsafeGet[float64](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, value)
	}
}

func (class) _slider_joint_get_param(impl func(ptr unsafe.Pointer, joint RID.Any, param gdclass.PhysicsServer3DSliderJointParam) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DSliderJointParam](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _joint_make_cone_twist(impl func(ptr unsafe.Pointer, joint RID.Any, body_A RID.Any, local_ref_A Transform3D.BasisOrigin, body_B RID.Any, local_ref_B Transform3D.BasisOrigin)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var body_A = gd.UnsafeGet[RID.Any](p_args, 1)

		var local_ref_A = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 2)

		var body_B = gd.UnsafeGet[RID.Any](p_args, 3)

		var local_ref_B = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 4)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, local_ref_A, body_B, local_ref_B)
	}
}

func (class) _cone_twist_joint_set_param(impl func(ptr unsafe.Pointer, joint RID.Any, param gdclass.PhysicsServer3DConeTwistJointParam, value float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DConeTwistJointParam](p_args, 1)

		var value = gd.UnsafeGet[float64](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, param, value)
	}
}

func (class) _cone_twist_joint_get_param(impl func(ptr unsafe.Pointer, joint RID.Any, param gdclass.PhysicsServer3DConeTwistJointParam) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DConeTwistJointParam](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, param)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _joint_make_generic_6dof(impl func(ptr unsafe.Pointer, joint RID.Any, body_A RID.Any, local_ref_A Transform3D.BasisOrigin, body_B RID.Any, local_ref_B Transform3D.BasisOrigin)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var body_A = gd.UnsafeGet[RID.Any](p_args, 1)

		var local_ref_A = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 2)

		var body_B = gd.UnsafeGet[RID.Any](p_args, 3)

		var local_ref_B = gd.UnsafeGet[Transform3D.BasisOrigin](p_args, 4)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, body_A, local_ref_A, body_B, local_ref_B)
	}
}

func (class) _generic_6dof_joint_set_param(impl func(ptr unsafe.Pointer, joint RID.Any, axis gd.Vector3Axis, param gdclass.PhysicsServer3DG6DOFJointAxisParam, value float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var axis = gd.UnsafeGet[gd.Vector3Axis](p_args, 1)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DG6DOFJointAxisParam](p_args, 2)

		var value = gd.UnsafeGet[float64](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, axis, param, value)
	}
}

func (class) _generic_6dof_joint_get_param(impl func(ptr unsafe.Pointer, joint RID.Any, axis gd.Vector3Axis, param gdclass.PhysicsServer3DG6DOFJointAxisParam) float64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var axis = gd.UnsafeGet[gd.Vector3Axis](p_args, 1)

		var param = gd.UnsafeGet[gdclass.PhysicsServer3DG6DOFJointAxisParam](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, axis, param)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _generic_6dof_joint_set_flag(impl func(ptr unsafe.Pointer, joint RID.Any, axis gd.Vector3Axis, flag gdclass.PhysicsServer3DG6DOFJointAxisFlag, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var axis = gd.UnsafeGet[gd.Vector3Axis](p_args, 1)

		var flag = gd.UnsafeGet[gdclass.PhysicsServer3DG6DOFJointAxisFlag](p_args, 2)

		var enable = gd.UnsafeGet[bool](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, axis, flag, enable)
	}
}

func (class) _generic_6dof_joint_get_flag(impl func(ptr unsafe.Pointer, joint RID.Any, axis gd.Vector3Axis, flag gdclass.PhysicsServer3DG6DOFJointAxisFlag) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var axis = gd.UnsafeGet[gd.Vector3Axis](p_args, 1)

		var flag = gd.UnsafeGet[gdclass.PhysicsServer3DG6DOFJointAxisFlag](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint, axis, flag)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _joint_get_type(impl func(ptr unsafe.Pointer, joint RID.Any) gdclass.PhysicsServer3DJointType) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _joint_set_solver_priority(impl func(ptr unsafe.Pointer, joint RID.Any, priority int64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var priority = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, priority)
	}
}

func (class) _joint_get_solver_priority(impl func(ptr unsafe.Pointer, joint RID.Any) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _joint_disable_collisions_between_bodies(impl func(ptr unsafe.Pointer, joint RID.Any, disable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		var disable = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, joint, disable)
	}
}

func (class) _joint_is_disabled_collisions_between_bodies(impl func(ptr unsafe.Pointer, joint RID.Any) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var joint = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, joint)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _free_rid(impl func(ptr unsafe.Pointer, rid RID.Any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var rid = gd.UnsafeGet[RID.Any](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, rid)
	}
}

func (class) _set_active(impl func(ptr unsafe.Pointer, active bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var active = gd.UnsafeGet[bool](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, active)
	}
}

func (class) _init(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (class) _step(impl func(ptr unsafe.Pointer, step float64)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var step = gd.UnsafeGet[float64](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, step)
	}
}

func (class) _sync(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (class) _flush_queries(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (class) _end_sync(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (class) _finish(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (class) _is_flushing_queries(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_process_info(impl func(ptr unsafe.Pointer, process_info gdclass.PhysicsServer3DProcessInfo) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var process_info = gd.UnsafeGet[gdclass.PhysicsServer3DProcessInfo](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, process_info)
		gd.UnsafeSet(p_back, ret)
	}
}

//go:nosplit
func (self class) BodyTestMotionIsExcludingBody(body RID.Any) bool { //gd:PhysicsServer3DExtension.body_test_motion_is_excluding_body
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsServer3DExtension.Bind_body_test_motion_is_excluding_body, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) BodyTestMotionIsExcludingObject(obj int64) bool { //gd:PhysicsServer3DExtension.body_test_motion_is_excluding_object
	var frame = callframe.New()
	callframe.Arg(frame, obj)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsServer3DExtension.Bind_body_test_motion_is_excluding_object, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gdclass.Register("PhysicsServer3DExtension", func(ptr gd.Object) any {
		return [1]gdclass.PhysicsServer3DExtension{*(*gdclass.PhysicsServer3DExtension)(unsafe.Pointer(&ptr))}
	})
}
