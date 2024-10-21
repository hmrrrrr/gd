package PhysicsDirectBodyState3DExtension

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/PhysicsDirectBodyState3D"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This class extends [PhysicsDirectBodyState3D] by providing additional virtual methods that can be overridden. When these methods are overridden, they will be called instead of the internal methods of the physics server.
Intended for use with GDExtension to create custom implementations of [PhysicsDirectBodyState3D].
	// PhysicsDirectBodyState3DExtension methods that can be overridden by a [Class] that extends it.
	type PhysicsDirectBodyState3DExtension interface {
		GetTotalGravity() gd.Vector3
		GetTotalLinearDamp() gd.Float
		GetTotalAngularDamp() gd.Float
		GetCenterOfMass() gd.Vector3
		GetCenterOfMassLocal() gd.Vector3
		GetPrincipalInertiaAxes() gd.Basis
		GetInverseMass() gd.Float
		GetInverseInertia() gd.Vector3
		GetInverseInertiaTensor() gd.Basis
		SetLinearVelocity(velocity gd.Vector3) 
		GetLinearVelocity() gd.Vector3
		SetAngularVelocity(velocity gd.Vector3) 
		GetAngularVelocity() gd.Vector3
		SetTransform(transform gd.Transform3D) 
		GetTransform() gd.Transform3D
		GetVelocityAtLocalPosition(local_position gd.Vector3) gd.Vector3
		ApplyCentralImpulse(impulse gd.Vector3) 
		ApplyImpulse(impulse gd.Vector3, position gd.Vector3) 
		ApplyTorqueImpulse(impulse gd.Vector3) 
		ApplyCentralForce(force gd.Vector3) 
		ApplyForce(force gd.Vector3, position gd.Vector3) 
		ApplyTorque(torque gd.Vector3) 
		AddConstantCentralForce(force gd.Vector3) 
		AddConstantForce(force gd.Vector3, position gd.Vector3) 
		AddConstantTorque(torque gd.Vector3) 
		SetConstantForce(force gd.Vector3) 
		GetConstantForce() gd.Vector3
		SetConstantTorque(torque gd.Vector3) 
		GetConstantTorque() gd.Vector3
		SetSleepState(enabled bool) 
		IsSleeping() bool
		GetContactCount() gd.Int
		GetContactLocalPosition(contact_idx gd.Int) gd.Vector3
		GetContactLocalNormal(contact_idx gd.Int) gd.Vector3
		GetContactImpulse(contact_idx gd.Int) gd.Vector3
		GetContactLocalShape(contact_idx gd.Int) gd.Int
		GetContactLocalVelocityAtPosition(contact_idx gd.Int) gd.Vector3
		GetContactCollider(contact_idx gd.Int) gd.RID
		GetContactColliderPosition(contact_idx gd.Int) gd.Vector3
		GetContactColliderId(contact_idx gd.Int) gd.Int
		GetContactColliderObject(contact_idx gd.Int) gd.Object
		GetContactColliderShape(contact_idx gd.Int) gd.Int
		GetContactColliderVelocityAtPosition(contact_idx gd.Int) gd.Vector3
		GetStep() gd.Float
		IntegrateForces() 
		GetSpaceState() object.PhysicsDirectSpaceState3D
	}

*/
type Simple [1]classdb.PhysicsDirectBodyState3DExtension
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.PhysicsDirectBodyState3DExtension
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

func (class) _get_total_gravity(impl func(ptr unsafe.Pointer) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_total_linear_damp(impl func(ptr unsafe.Pointer) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_total_angular_damp(impl func(ptr unsafe.Pointer) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_center_of_mass(impl func(ptr unsafe.Pointer) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_center_of_mass_local(impl func(ptr unsafe.Pointer) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_principal_inertia_axes(impl func(ptr unsafe.Pointer) gd.Basis, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_inverse_mass(impl func(ptr unsafe.Pointer) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_inverse_inertia(impl func(ptr unsafe.Pointer) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_inverse_inertia_tensor(impl func(ptr unsafe.Pointer) gd.Basis, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _set_linear_velocity(impl func(ptr unsafe.Pointer, velocity gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var velocity = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, velocity)
		ctx.End()
	}
}

func (class) _get_linear_velocity(impl func(ptr unsafe.Pointer) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _set_angular_velocity(impl func(ptr unsafe.Pointer, velocity gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var velocity = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, velocity)
		ctx.End()
	}
}

func (class) _get_angular_velocity(impl func(ptr unsafe.Pointer) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _set_transform(impl func(ptr unsafe.Pointer, transform gd.Transform3D) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, transform)
		ctx.End()
	}
}

func (class) _get_transform(impl func(ptr unsafe.Pointer) gd.Transform3D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_velocity_at_local_position(impl func(ptr unsafe.Pointer, local_position gd.Vector3) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var local_position = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, local_position)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _apply_central_impulse(impl func(ptr unsafe.Pointer, impulse gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var impulse = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, impulse)
		ctx.End()
	}
}

func (class) _apply_impulse(impl func(ptr unsafe.Pointer, impulse gd.Vector3, position gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var impulse = gd.UnsafeGet[gd.Vector3](p_args,0)
		var position = gd.UnsafeGet[gd.Vector3](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, impulse, position)
		ctx.End()
	}
}

func (class) _apply_torque_impulse(impl func(ptr unsafe.Pointer, impulse gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var impulse = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, impulse)
		ctx.End()
	}
}

func (class) _apply_central_force(impl func(ptr unsafe.Pointer, force gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var force = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force)
		ctx.End()
	}
}

func (class) _apply_force(impl func(ptr unsafe.Pointer, force gd.Vector3, position gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var force = gd.UnsafeGet[gd.Vector3](p_args,0)
		var position = gd.UnsafeGet[gd.Vector3](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force, position)
		ctx.End()
	}
}

func (class) _apply_torque(impl func(ptr unsafe.Pointer, torque gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var torque = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, torque)
		ctx.End()
	}
}

func (class) _add_constant_central_force(impl func(ptr unsafe.Pointer, force gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var force = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force)
		ctx.End()
	}
}

func (class) _add_constant_force(impl func(ptr unsafe.Pointer, force gd.Vector3, position gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var force = gd.UnsafeGet[gd.Vector3](p_args,0)
		var position = gd.UnsafeGet[gd.Vector3](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force, position)
		ctx.End()
	}
}

func (class) _add_constant_torque(impl func(ptr unsafe.Pointer, torque gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var torque = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, torque)
		ctx.End()
	}
}

func (class) _set_constant_force(impl func(ptr unsafe.Pointer, force gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var force = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, force)
		ctx.End()
	}
}

func (class) _get_constant_force(impl func(ptr unsafe.Pointer) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _set_constant_torque(impl func(ptr unsafe.Pointer, torque gd.Vector3) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var torque = gd.UnsafeGet[gd.Vector3](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, torque)
		ctx.End()
	}
}

func (class) _get_constant_torque(impl func(ptr unsafe.Pointer) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _set_sleep_state(impl func(ptr unsafe.Pointer, enabled bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var enabled = gd.UnsafeGet[bool](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, enabled)
		ctx.End()
	}
}

func (class) _is_sleeping(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_contact_count(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_contact_local_position(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_contact_local_normal(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_contact_impulse(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_contact_local_shape(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_contact_local_velocity_at_position(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_contact_collider(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.RID, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_contact_collider_position(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_contact_collider_id(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_contact_collider_object(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Object, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, mmm.End(ret.AsPointer()))
		ctx.End()
	}
}

func (class) _get_contact_collider_shape(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_contact_collider_velocity_at_position(impl func(ptr unsafe.Pointer, contact_idx gd.Int) gd.Vector3, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var contact_idx = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, contact_idx)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _get_step(impl func(ptr unsafe.Pointer) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _integrate_forces(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

func (class) _get_space_state(impl func(ptr unsafe.Pointer) object.PhysicsDirectSpaceState3D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		ctx.End()
	}
}


//go:nosplit
func (self class) AsPhysicsDirectBodyState3DExtension() Expert { return self[0].AsPhysicsDirectBodyState3DExtension() }


//go:nosplit
func (self Simple) AsPhysicsDirectBodyState3DExtension() Simple { return self[0].AsPhysicsDirectBodyState3DExtension() }


//go:nosplit
func (self class) AsPhysicsDirectBodyState3D() PhysicsDirectBodyState3D.Expert { return self[0].AsPhysicsDirectBodyState3D() }


//go:nosplit
func (self Simple) AsPhysicsDirectBodyState3D() PhysicsDirectBodyState3D.Simple { return self[0].AsPhysicsDirectBodyState3D() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_total_gravity": return reflect.ValueOf(self._get_total_gravity);
	case "_get_total_linear_damp": return reflect.ValueOf(self._get_total_linear_damp);
	case "_get_total_angular_damp": return reflect.ValueOf(self._get_total_angular_damp);
	case "_get_center_of_mass": return reflect.ValueOf(self._get_center_of_mass);
	case "_get_center_of_mass_local": return reflect.ValueOf(self._get_center_of_mass_local);
	case "_get_principal_inertia_axes": return reflect.ValueOf(self._get_principal_inertia_axes);
	case "_get_inverse_mass": return reflect.ValueOf(self._get_inverse_mass);
	case "_get_inverse_inertia": return reflect.ValueOf(self._get_inverse_inertia);
	case "_get_inverse_inertia_tensor": return reflect.ValueOf(self._get_inverse_inertia_tensor);
	case "_set_linear_velocity": return reflect.ValueOf(self._set_linear_velocity);
	case "_get_linear_velocity": return reflect.ValueOf(self._get_linear_velocity);
	case "_set_angular_velocity": return reflect.ValueOf(self._set_angular_velocity);
	case "_get_angular_velocity": return reflect.ValueOf(self._get_angular_velocity);
	case "_set_transform": return reflect.ValueOf(self._set_transform);
	case "_get_transform": return reflect.ValueOf(self._get_transform);
	case "_get_velocity_at_local_position": return reflect.ValueOf(self._get_velocity_at_local_position);
	case "_apply_central_impulse": return reflect.ValueOf(self._apply_central_impulse);
	case "_apply_impulse": return reflect.ValueOf(self._apply_impulse);
	case "_apply_torque_impulse": return reflect.ValueOf(self._apply_torque_impulse);
	case "_apply_central_force": return reflect.ValueOf(self._apply_central_force);
	case "_apply_force": return reflect.ValueOf(self._apply_force);
	case "_apply_torque": return reflect.ValueOf(self._apply_torque);
	case "_add_constant_central_force": return reflect.ValueOf(self._add_constant_central_force);
	case "_add_constant_force": return reflect.ValueOf(self._add_constant_force);
	case "_add_constant_torque": return reflect.ValueOf(self._add_constant_torque);
	case "_set_constant_force": return reflect.ValueOf(self._set_constant_force);
	case "_get_constant_force": return reflect.ValueOf(self._get_constant_force);
	case "_set_constant_torque": return reflect.ValueOf(self._set_constant_torque);
	case "_get_constant_torque": return reflect.ValueOf(self._get_constant_torque);
	case "_set_sleep_state": return reflect.ValueOf(self._set_sleep_state);
	case "_is_sleeping": return reflect.ValueOf(self._is_sleeping);
	case "_get_contact_count": return reflect.ValueOf(self._get_contact_count);
	case "_get_contact_local_position": return reflect.ValueOf(self._get_contact_local_position);
	case "_get_contact_local_normal": return reflect.ValueOf(self._get_contact_local_normal);
	case "_get_contact_impulse": return reflect.ValueOf(self._get_contact_impulse);
	case "_get_contact_local_shape": return reflect.ValueOf(self._get_contact_local_shape);
	case "_get_contact_local_velocity_at_position": return reflect.ValueOf(self._get_contact_local_velocity_at_position);
	case "_get_contact_collider": return reflect.ValueOf(self._get_contact_collider);
	case "_get_contact_collider_position": return reflect.ValueOf(self._get_contact_collider_position);
	case "_get_contact_collider_id": return reflect.ValueOf(self._get_contact_collider_id);
	case "_get_contact_collider_object": return reflect.ValueOf(self._get_contact_collider_object);
	case "_get_contact_collider_shape": return reflect.ValueOf(self._get_contact_collider_shape);
	case "_get_contact_collider_velocity_at_position": return reflect.ValueOf(self._get_contact_collider_velocity_at_position);
	case "_get_step": return reflect.ValueOf(self._get_step);
	case "_integrate_forces": return reflect.ValueOf(self._integrate_forces);
	case "_get_space_state": return reflect.ValueOf(self._get_space_state);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("PhysicsDirectBodyState3DExtension", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
