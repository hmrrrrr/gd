// Package VehicleBody3D provides methods for working with VehicleBody3D object instances.
package VehicleBody3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/RigidBody3D"
import "graphics.gd/classdb/PhysicsBody3D"
import "graphics.gd/classdb/CollisionObject3D"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This physics body implements all the physics logic needed to simulate a car. It is based on the raycast vehicle system commonly found in physics engines. Aside from a [CollisionShape3D] for the main body of the vehicle, you must also add a [VehicleWheel3D] node for each wheel. You should also add a [MeshInstance3D] to this node for the 3D model of the vehicle, but this model should generally not include meshes for the wheels. You can control the vehicle by using the [member brake], [member engine_force], and [member steering] properties. The position or orientation of this node shouldn't be changed directly.
[b]Note:[/b] The origin point of your VehicleBody3D will determine the center of gravity of your vehicle. To make the vehicle more grounded, the origin point is usually kept low, moving the [CollisionShape3D] and [MeshInstance3D] upwards.
[b]Note:[/b] This class has known issues and isn't designed to provide realistic 3D vehicle physics. If you want advanced vehicle physics, you may have to write your own physics integration using [CharacterBody3D] or [RigidBody3D].
*/
type Instance [1]gdclass.VehicleBody3D
type Any interface {
	gd.IsClass
	AsVehicleBody3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VehicleBody3D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VehicleBody3D"))
	return Instance{*(*gdclass.VehicleBody3D)(unsafe.Pointer(&object))}
}

func (self Instance) EngineForce() Float.X {
	return Float.X(Float.X(class(self).GetEngineForce()))
}

func (self Instance) SetEngineForce(value Float.X) {
	class(self).SetEngineForce(gd.Float(value))
}

func (self Instance) Brake() Float.X {
	return Float.X(Float.X(class(self).GetBrake()))
}

func (self Instance) SetBrake(value Float.X) {
	class(self).SetBrake(gd.Float(value))
}

func (self Instance) Steering() Float.X {
	return Float.X(Float.X(class(self).GetSteering()))
}

func (self Instance) SetSteering(value Float.X) {
	class(self).SetSteering(gd.Float(value))
}

//go:nosplit
func (self class) SetEngineForce(engine_force gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, engine_force)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleBody3D.Bind_set_engine_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEngineForce() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleBody3D.Bind_get_engine_force, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBrake(brake gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, brake)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleBody3D.Bind_set_brake, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBrake() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleBody3D.Bind_get_brake, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSteering(steering gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, steering)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleBody3D.Bind_set_steering, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSteering() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VehicleBody3D.Bind_get_steering, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVehicleBody3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVehicleBody3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRigidBody3D() RigidBody3D.Advanced {
	return *((*RigidBody3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsRigidBody3D() RigidBody3D.Instance {
	return *((*RigidBody3D.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(self.AsRigidBody3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRigidBody3D(), name)
	}
}
func init() {
	gdclass.Register("VehicleBody3D", func(ptr gd.Object) any {
		return [1]gdclass.VehicleBody3D{*(*gdclass.VehicleBody3D)(unsafe.Pointer(&ptr))}
	})
}
