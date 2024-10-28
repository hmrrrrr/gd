package RigidBody3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/PhysicsBody3D"
import "grow.graphics/gd/gdclass/CollisionObject3D"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
[RigidBody3D] implements full 3D physics. It cannot be controlled directly, instead, you must apply forces to it (gravity, impulses, etc.), and the physics simulation will calculate the resulting movement, rotation, react to collisions, and affect other physics bodies in its path.
The body's behavior can be adjusted via [member lock_rotation], [member freeze], and [member freeze_mode]. By changing various properties of the object, such as [member mass], you can control how the physics simulation acts on it.
A rigid body will always maintain its shape and size, even when forces are applied to it. It is useful for objects that can be interacted with in an environment, such as a tree that can be knocked over or a stack of crates that can be pushed around.
If you need to override the default physics behavior, you can write a custom force integration function. See [member custom_integrator].
[b]Note:[/b] Changing the 3D transform or [member linear_velocity] of a [RigidBody3D] very often may lead to some unpredictable behaviors. If you need to directly affect the body, prefer [method _integrate_forces] as it allows you to directly access the physics state.
	// RigidBody3D methods that can be overridden by a [Class] that extends it.
	type RigidBody3D interface {
		//Called during physics processing, allowing you to read and safely modify the simulation state for the object. By default, it is called before the standard force integration, but the [member custom_integrator] property allows you to disable the standard force integration and do fully custom force integration for a body.
		IntegrateForces(state gdclass.PhysicsDirectBodyState3D) 
	}

*/
type Go [1]classdb.RigidBody3D

/*
Called during physics processing, allowing you to read and safely modify the simulation state for the object. By default, it is called before the standard force integration, but the [member custom_integrator] property allows you to disable the standard force integration and do fully custom force integration for a body.
*/
func (Go) _integrate_forces(impl func(ptr unsafe.Pointer, state gdclass.PhysicsDirectBodyState3D) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.PhysicsDirectBodyState3D{discreet.New[classdb.PhysicsDirectBodyState3D]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})}
		defer discreet.End(state[0])
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, state)
	}
}

/*
Returns the inverse inertia tensor basis. This is used to calculate the angular acceleration resulting from a torque applied to the [RigidBody3D].
*/
func (self Go) GetInverseInertiaTensor() gd.Basis {
	return gd.Basis(class(self).GetInverseInertiaTensor())
}

/*
Returns the number of contacts this body has with other bodies. By default, this returns 0 unless bodies are configured to monitor contacts (see [member contact_monitor]).
[b]Note:[/b] To retrieve the colliding bodies, use [method get_colliding_bodies].
*/
func (self Go) GetContactCount() int {
	return int(int(class(self).GetContactCount()))
}

/*
Sets an axis velocity. The velocity in the given vector axis will be set as the given vector length. This is useful for jumping behavior.
*/
func (self Go) SetAxisVelocity(axis_velocity gd.Vector3) {
	class(self).SetAxisVelocity(axis_velocity)
}

/*
Applies a directional impulse without affecting rotation.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
This is equivalent to using [method apply_impulse] at the body's center of mass.
*/
func (self Go) ApplyCentralImpulse(impulse gd.Vector3) {
	class(self).ApplyCentralImpulse(impulse)
}

/*
Applies a positioned impulse to the body.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
[param position] is the offset from the body origin in global coordinates.
*/
func (self Go) ApplyImpulse(impulse gd.Vector3) {
	class(self).ApplyImpulse(impulse, gd.Vector3{0, 0, 0})
}

/*
Applies a rotational impulse to the body without affecting the position.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
[b]Note:[/b] [member inertia] is required for this to work. To have [member inertia], an active [CollisionShape3D] must be a child of the node, or you can manually set [member inertia].
*/
func (self Go) ApplyTorqueImpulse(impulse gd.Vector3) {
	class(self).ApplyTorqueImpulse(impulse)
}

/*
Applies a directional force without affecting rotation. A force is time dependent and meant to be applied every physics update.
This is equivalent to using [method apply_force] at the body's center of mass.
*/
func (self Go) ApplyCentralForce(force gd.Vector3) {
	class(self).ApplyCentralForce(force)
}

/*
Applies a positioned force to the body. A force is time dependent and meant to be applied every physics update.
[param position] is the offset from the body origin in global coordinates.
*/
func (self Go) ApplyForce(force gd.Vector3) {
	class(self).ApplyForce(force, gd.Vector3{0, 0, 0})
}

/*
Applies a rotational force without affecting position. A force is time dependent and meant to be applied every physics update.
[b]Note:[/b] [member inertia] is required for this to work. To have [member inertia], an active [CollisionShape3D] must be a child of the node, or you can manually set [member inertia].
*/
func (self Go) ApplyTorque(torque gd.Vector3) {
	class(self).ApplyTorque(torque)
}

/*
Adds a constant directional force without affecting rotation that keeps being applied over time until cleared with [code]constant_force = Vector3(0, 0, 0)[/code].
This is equivalent to using [method add_constant_force] at the body's center of mass.
*/
func (self Go) AddConstantCentralForce(force gd.Vector3) {
	class(self).AddConstantCentralForce(force)
}

/*
Adds a constant positioned force to the body that keeps being applied over time until cleared with [code]constant_force = Vector3(0, 0, 0)[/code].
[param position] is the offset from the body origin in global coordinates.
*/
func (self Go) AddConstantForce(force gd.Vector3) {
	class(self).AddConstantForce(force, gd.Vector3{0, 0, 0})
}

/*
Adds a constant rotational force without affecting position that keeps being applied over time until cleared with [code]constant_torque = Vector3(0, 0, 0)[/code].
*/
func (self Go) AddConstantTorque(torque gd.Vector3) {
	class(self).AddConstantTorque(torque)
}

/*
Returns a list of the bodies colliding with this one. Requires [member contact_monitor] to be set to [code]true[/code] and [member max_contacts_reported] to be set high enough to detect all the collisions.
[b]Note:[/b] The result of this test is not immediate after moving objects. For performance, list of collisions is updated once per frame and before the physics step. Consider using signals instead.
*/
func (self Go) GetCollidingBodies() gd.Array {
	return gd.Array(class(self).GetCollidingBodies())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.RigidBody3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RigidBody3D"))
	return Go{classdb.RigidBody3D(object)}
}

func (self Go) Mass() float64 {
		return float64(float64(class(self).GetMass()))
}

func (self Go) SetMass(value float64) {
	class(self).SetMass(gd.Float(value))
}

func (self Go) PhysicsMaterialOverride() gdclass.PhysicsMaterial {
		return gdclass.PhysicsMaterial(class(self).GetPhysicsMaterialOverride())
}

func (self Go) SetPhysicsMaterialOverride(value gdclass.PhysicsMaterial) {
	class(self).SetPhysicsMaterialOverride(value)
}

func (self Go) GravityScale() float64 {
		return float64(float64(class(self).GetGravityScale()))
}

func (self Go) SetGravityScale(value float64) {
	class(self).SetGravityScale(gd.Float(value))
}

func (self Go) CenterOfMassMode() classdb.RigidBody3DCenterOfMassMode {
		return classdb.RigidBody3DCenterOfMassMode(class(self).GetCenterOfMassMode())
}

func (self Go) SetCenterOfMassMode(value classdb.RigidBody3DCenterOfMassMode) {
	class(self).SetCenterOfMassMode(value)
}

func (self Go) CenterOfMass() gd.Vector3 {
		return gd.Vector3(class(self).GetCenterOfMass())
}

func (self Go) SetCenterOfMass(value gd.Vector3) {
	class(self).SetCenterOfMass(value)
}

func (self Go) Inertia() gd.Vector3 {
		return gd.Vector3(class(self).GetInertia())
}

func (self Go) SetInertia(value gd.Vector3) {
	class(self).SetInertia(value)
}

func (self Go) Sleeping() bool {
		return bool(class(self).IsSleeping())
}

func (self Go) SetSleeping(value bool) {
	class(self).SetSleeping(value)
}

func (self Go) CanSleep() bool {
		return bool(class(self).IsAbleToSleep())
}

func (self Go) SetCanSleep(value bool) {
	class(self).SetCanSleep(value)
}

func (self Go) LockRotation() bool {
		return bool(class(self).IsLockRotationEnabled())
}

func (self Go) SetLockRotation(value bool) {
	class(self).SetLockRotationEnabled(value)
}

func (self Go) Freeze() bool {
		return bool(class(self).IsFreezeEnabled())
}

func (self Go) SetFreeze(value bool) {
	class(self).SetFreezeEnabled(value)
}

func (self Go) FreezeMode() classdb.RigidBody3DFreezeMode {
		return classdb.RigidBody3DFreezeMode(class(self).GetFreezeMode())
}

func (self Go) SetFreezeMode(value classdb.RigidBody3DFreezeMode) {
	class(self).SetFreezeMode(value)
}

func (self Go) CustomIntegrator() bool {
		return bool(class(self).IsUsingCustomIntegrator())
}

func (self Go) SetCustomIntegrator(value bool) {
	class(self).SetUseCustomIntegrator(value)
}

func (self Go) ContinuousCd() bool {
		return bool(class(self).IsUsingContinuousCollisionDetection())
}

func (self Go) SetContinuousCd(value bool) {
	class(self).SetUseContinuousCollisionDetection(value)
}

func (self Go) ContactMonitor() bool {
		return bool(class(self).IsContactMonitorEnabled())
}

func (self Go) SetContactMonitor(value bool) {
	class(self).SetContactMonitor(value)
}

func (self Go) MaxContactsReported() int {
		return int(int(class(self).GetMaxContactsReported()))
}

func (self Go) SetMaxContactsReported(value int) {
	class(self).SetMaxContactsReported(gd.Int(value))
}

func (self Go) LinearVelocity() gd.Vector3 {
		return gd.Vector3(class(self).GetLinearVelocity())
}

func (self Go) SetLinearVelocity(value gd.Vector3) {
	class(self).SetLinearVelocity(value)
}

func (self Go) LinearDampMode() classdb.RigidBody3DDampMode {
		return classdb.RigidBody3DDampMode(class(self).GetLinearDampMode())
}

func (self Go) SetLinearDampMode(value classdb.RigidBody3DDampMode) {
	class(self).SetLinearDampMode(value)
}

func (self Go) LinearDamp() float64 {
		return float64(float64(class(self).GetLinearDamp()))
}

func (self Go) SetLinearDamp(value float64) {
	class(self).SetLinearDamp(gd.Float(value))
}

func (self Go) AngularVelocity() gd.Vector3 {
		return gd.Vector3(class(self).GetAngularVelocity())
}

func (self Go) SetAngularVelocity(value gd.Vector3) {
	class(self).SetAngularVelocity(value)
}

func (self Go) AngularDampMode() classdb.RigidBody3DDampMode {
		return classdb.RigidBody3DDampMode(class(self).GetAngularDampMode())
}

func (self Go) SetAngularDampMode(value classdb.RigidBody3DDampMode) {
	class(self).SetAngularDampMode(value)
}

func (self Go) AngularDamp() float64 {
		return float64(float64(class(self).GetAngularDamp()))
}

func (self Go) SetAngularDamp(value float64) {
	class(self).SetAngularDamp(gd.Float(value))
}

func (self Go) ConstantForce() gd.Vector3 {
		return gd.Vector3(class(self).GetConstantForce())
}

func (self Go) SetConstantForce(value gd.Vector3) {
	class(self).SetConstantForce(value)
}

func (self Go) ConstantTorque() gd.Vector3 {
		return gd.Vector3(class(self).GetConstantTorque())
}

func (self Go) SetConstantTorque(value gd.Vector3) {
	class(self).SetConstantTorque(value)
}

/*
Called during physics processing, allowing you to read and safely modify the simulation state for the object. By default, it is called before the standard force integration, but the [member custom_integrator] property allows you to disable the standard force integration and do fully custom force integration for a body.
*/
func (class) _integrate_forces(impl func(ptr unsafe.Pointer, state gdclass.PhysicsDirectBodyState3D) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = gdclass.PhysicsDirectBodyState3D{discreet.New[classdb.PhysicsDirectBodyState3D]([3]uintptr{gd.UnsafeGet[uintptr](p_args,0)})}
		defer discreet.End(state[0])
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, state)
	}
}

//go:nosplit
func (self class) SetMass(mass gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, mass)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMass() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetInertia(inertia gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, inertia)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_inertia, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInertia() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_inertia, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCenterOfMassMode(mode classdb.RigidBody3DCenterOfMassMode)  {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_center_of_mass_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCenterOfMassMode() classdb.RigidBody3DCenterOfMassMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RigidBody3DCenterOfMassMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_center_of_mass_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCenterOfMass(center_of_mass gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, center_of_mass)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_center_of_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCenterOfMass() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_center_of_mass, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPhysicsMaterialOverride(physics_material_override gdclass.PhysicsMaterial)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(physics_material_override[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_physics_material_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPhysicsMaterialOverride() gdclass.PhysicsMaterial {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_physics_material_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.PhysicsMaterial{classdb.PhysicsMaterial(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLinearVelocity(linear_velocity gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, linear_velocity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLinearVelocity() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAngularVelocity(angular_velocity gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, angular_velocity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAngularVelocity() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the inverse inertia tensor basis. This is used to calculate the angular acceleration resulting from a torque applied to the [RigidBody3D].
*/
//go:nosplit
func (self class) GetInverseInertiaTensor() gd.Basis {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Basis](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_inverse_inertia_tensor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGravityScale(gravity_scale gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, gravity_scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_gravity_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGravityScale() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_gravity_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLinearDampMode(linear_damp_mode classdb.RigidBody3DDampMode)  {
	var frame = callframe.New()
	callframe.Arg(frame, linear_damp_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_linear_damp_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLinearDampMode() classdb.RigidBody3DDampMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RigidBody3DDampMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_linear_damp_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAngularDampMode(angular_damp_mode classdb.RigidBody3DDampMode)  {
	var frame = callframe.New()
	callframe.Arg(frame, angular_damp_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_angular_damp_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAngularDampMode() classdb.RigidBody3DDampMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RigidBody3DDampMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_angular_damp_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLinearDamp(linear_damp gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, linear_damp)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_linear_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLinearDamp() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_linear_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAngularDamp(angular_damp gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, angular_damp)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_angular_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAngularDamp() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_angular_damp, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxContactsReported(amount gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_max_contacts_reported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxContactsReported() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_max_contacts_reported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of contacts this body has with other bodies. By default, this returns 0 unless bodies are configured to monitor contacts (see [member contact_monitor]).
[b]Note:[/b] To retrieve the colliding bodies, use [method get_colliding_bodies].
*/
//go:nosplit
func (self class) GetContactCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_contact_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseCustomIntegrator(enable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_use_custom_integrator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingCustomIntegrator() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_is_using_custom_integrator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetContactMonitor(enabled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_contact_monitor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsContactMonitorEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_is_contact_monitor_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseContinuousCollisionDetection(enable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_use_continuous_collision_detection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingContinuousCollisionDetection() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_is_using_continuous_collision_detection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets an axis velocity. The velocity in the given vector axis will be set as the given vector length. This is useful for jumping behavior.
*/
//go:nosplit
func (self class) SetAxisVelocity(axis_velocity gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, axis_velocity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_axis_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Applies a directional impulse without affecting rotation.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
This is equivalent to using [method apply_impulse] at the body's center of mass.
*/
//go:nosplit
func (self class) ApplyCentralImpulse(impulse gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, impulse)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_apply_central_impulse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Applies a positioned impulse to the body.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
[param position] is the offset from the body origin in global coordinates.
*/
//go:nosplit
func (self class) ApplyImpulse(impulse gd.Vector3, position gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, impulse)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_apply_impulse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Applies a rotational impulse to the body without affecting the position.
An impulse is time-independent! Applying an impulse every frame would result in a framerate-dependent force. For this reason, it should only be used when simulating one-time impacts (use the "_force" functions otherwise).
[b]Note:[/b] [member inertia] is required for this to work. To have [member inertia], an active [CollisionShape3D] must be a child of the node, or you can manually set [member inertia].
*/
//go:nosplit
func (self class) ApplyTorqueImpulse(impulse gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, impulse)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_apply_torque_impulse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Applies a directional force without affecting rotation. A force is time dependent and meant to be applied every physics update.
This is equivalent to using [method apply_force] at the body's center of mass.
*/
//go:nosplit
func (self class) ApplyCentralForce(force gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, force)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_apply_central_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Applies a positioned force to the body. A force is time dependent and meant to be applied every physics update.
[param position] is the offset from the body origin in global coordinates.
*/
//go:nosplit
func (self class) ApplyForce(force gd.Vector3, position gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, force)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_apply_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Applies a rotational force without affecting position. A force is time dependent and meant to be applied every physics update.
[b]Note:[/b] [member inertia] is required for this to work. To have [member inertia], an active [CollisionShape3D] must be a child of the node, or you can manually set [member inertia].
*/
//go:nosplit
func (self class) ApplyTorque(torque gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, torque)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_apply_torque, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a constant directional force without affecting rotation that keeps being applied over time until cleared with [code]constant_force = Vector3(0, 0, 0)[/code].
This is equivalent to using [method add_constant_force] at the body's center of mass.
*/
//go:nosplit
func (self class) AddConstantCentralForce(force gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, force)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_add_constant_central_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a constant positioned force to the body that keeps being applied over time until cleared with [code]constant_force = Vector3(0, 0, 0)[/code].
[param position] is the offset from the body origin in global coordinates.
*/
//go:nosplit
func (self class) AddConstantForce(force gd.Vector3, position gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, force)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_add_constant_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a constant rotational force without affecting position that keeps being applied over time until cleared with [code]constant_torque = Vector3(0, 0, 0)[/code].
*/
//go:nosplit
func (self class) AddConstantTorque(torque gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, torque)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_add_constant_torque, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetConstantForce(force gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, force)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_constant_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetConstantForce() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_constant_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetConstantTorque(torque gd.Vector3)  {
	var frame = callframe.New()
	callframe.Arg(frame, torque)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_constant_torque, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetConstantTorque() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_constant_torque, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSleeping(sleeping bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, sleeping)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_sleeping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsSleeping() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_is_sleeping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCanSleep(able_to_sleep bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, able_to_sleep)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_can_sleep, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAbleToSleep() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_is_able_to_sleep, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLockRotationEnabled(lock_rotation bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, lock_rotation)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_lock_rotation_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsLockRotationEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_is_lock_rotation_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFreezeEnabled(freeze_mode bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, freeze_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_freeze_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsFreezeEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_is_freeze_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFreezeMode(freeze_mode classdb.RigidBody3DFreezeMode)  {
	var frame = callframe.New()
	callframe.Arg(frame, freeze_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_set_freeze_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFreezeMode() classdb.RigidBody3DFreezeMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.RigidBody3DFreezeMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_freeze_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a list of the bodies colliding with this one. Requires [member contact_monitor] to be set to [code]true[/code] and [member max_contacts_reported] to be set high enough to detect all the collisions.
[b]Note:[/b] The result of this test is not immediate after moving objects. For performance, list of collisions is updated once per frame and before the physics step. Consider using signals instead.
*/
//go:nosplit
func (self class) GetCollidingBodies() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RigidBody3D.Bind_get_colliding_bodies, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self Go) OnBodyShapeEntered(cb func(body_rid gd.RID, body gdclass.Node, body_shape_index int, local_shape_index int)) {
	self[0].AsObject().Connect(gd.NewStringName("body_shape_entered"), gd.NewCallable(cb), 0)
}


func (self Go) OnBodyShapeExited(cb func(body_rid gd.RID, body gdclass.Node, body_shape_index int, local_shape_index int)) {
	self[0].AsObject().Connect(gd.NewStringName("body_shape_exited"), gd.NewCallable(cb), 0)
}


func (self Go) OnBodyEntered(cb func(body gdclass.Node)) {
	self[0].AsObject().Connect(gd.NewStringName("body_entered"), gd.NewCallable(cb), 0)
}


func (self Go) OnBodyExited(cb func(body gdclass.Node)) {
	self[0].AsObject().Connect(gd.NewStringName("body_exited"), gd.NewCallable(cb), 0)
}


func (self Go) OnSleepingStateChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("sleeping_state_changed"), gd.NewCallable(cb), 0)
}


func (self class) AsRigidBody3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsRigidBody3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsPhysicsBody3D() PhysicsBody3D.GD { return *((*PhysicsBody3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsPhysicsBody3D() PhysicsBody3D.Go { return *((*PhysicsBody3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsCollisionObject3D() CollisionObject3D.GD { return *((*CollisionObject3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCollisionObject3D() CollisionObject3D.Go { return *((*CollisionObject3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_integrate_forces": return reflect.ValueOf(self._integrate_forces);
	default: return gd.VirtualByName(self.AsPhysicsBody3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_integrate_forces": return reflect.ValueOf(self._integrate_forces);
	default: return gd.VirtualByName(self.AsPhysicsBody3D(), name)
	}
}
func init() {classdb.Register("RigidBody3D", func(ptr gd.Object) any { return classdb.RigidBody3D(ptr) })}
type FreezeMode = classdb.RigidBody3DFreezeMode

const (
/*Static body freeze mode (default). The body is not affected by gravity and forces. It can be only moved by user code and doesn't collide with other bodies along its path.*/
	FreezeModeStatic FreezeMode = 0
/*Kinematic body freeze mode. Similar to [constant FREEZE_MODE_STATIC], but collides with other bodies along its path when moved. Useful for a frozen body that needs to be animated.*/
	FreezeModeKinematic FreezeMode = 1
)
type CenterOfMassMode = classdb.RigidBody3DCenterOfMassMode

const (
/*In this mode, the body's center of mass is calculated automatically based on its shapes. This assumes that the shapes' origins are also their center of mass.*/
	CenterOfMassModeAuto CenterOfMassMode = 0
/*In this mode, the body's center of mass is set through [member center_of_mass]. Defaults to the body's origin position.*/
	CenterOfMassModeCustom CenterOfMassMode = 1
)
type DampMode = classdb.RigidBody3DDampMode

const (
/*In this mode, the body's damping value is added to any value set in areas or the default value.*/
	DampModeCombine DampMode = 0
/*In this mode, the body's damping value replaces any value set in areas or the default value.*/
	DampModeReplace DampMode = 1
)
