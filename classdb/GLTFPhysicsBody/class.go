// Package GLTFPhysicsBody provides methods for working with GLTFPhysicsBody object instances.
package GLTFPhysicsBody

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
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Vector3"
import "graphics.gd/variant/Quaternion"
import "graphics.gd/variant/Basis"

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

/*
Represents a physics body as an intermediary between the [code]OMI_physics_body[/code] GLTF data and Godot's nodes, and it's abstracted in a way that allows adding support for different GLTF physics extensions in the future.
*/
type Instance [1]gdclass.GLTFPhysicsBody

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGLTFPhysicsBody() Instance
}

/*
Creates a new GLTFPhysicsBody instance from the given Godot [CollisionObject3D] node.
*/
func FromNode(body_node [1]gdclass.CollisionObject3D) [1]gdclass.GLTFPhysicsBody { //gd:GLTFPhysicsBody.from_node
	self := Instance{}
	return [1]gdclass.GLTFPhysicsBody(class(self).FromNode(body_node))
}

/*
Converts this GLTFPhysicsBody instance into a Godot [CollisionObject3D] node.
*/
func (self Instance) ToNode() [1]gdclass.CollisionObject3D { //gd:GLTFPhysicsBody.to_node
	return [1]gdclass.CollisionObject3D(class(self).ToNode())
}

/*
Creates a new GLTFPhysicsBody instance by parsing the given [Dictionary] in the [code]OMI_physics_body[/code] GLTF extension format.
*/
func FromDictionary(dictionary Structure) [1]gdclass.GLTFPhysicsBody { //gd:GLTFPhysicsBody.from_dictionary
	self := Instance{}
	return [1]gdclass.GLTFPhysicsBody(class(self).FromDictionary(gd.DictionaryFromMap(dictionary)))
}

/*
Serializes this GLTFPhysicsBody instance into a [Dictionary]. It will be in the format expected by the [code]OMI_physics_body[/code] GLTF extension.
*/
func (self Instance) ToDictionary() Structure { //gd:GLTFPhysicsBody.to_dictionary
	return Structure(gd.DictionaryAs[Structure](class(self).ToDictionary()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GLTFPhysicsBody

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFPhysicsBody"))
	casted := Instance{*(*gdclass.GLTFPhysicsBody)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) BodyType() string {
	return string(class(self).GetBodyType().String())
}

func (self Instance) SetBodyType(value string) {
	class(self).SetBodyType(String.New(value))
}

func (self Instance) Mass() Float.X {
	return Float.X(Float.X(class(self).GetMass()))
}

func (self Instance) SetMass(value Float.X) {
	class(self).SetMass(gd.Float(value))
}

func (self Instance) LinearVelocity() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetLinearVelocity())
}

func (self Instance) SetLinearVelocity(value Vector3.XYZ) {
	class(self).SetLinearVelocity(gd.Vector3(value))
}

func (self Instance) AngularVelocity() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetAngularVelocity())
}

func (self Instance) SetAngularVelocity(value Vector3.XYZ) {
	class(self).SetAngularVelocity(gd.Vector3(value))
}

func (self Instance) CenterOfMass() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetCenterOfMass())
}

func (self Instance) SetCenterOfMass(value Vector3.XYZ) {
	class(self).SetCenterOfMass(gd.Vector3(value))
}

func (self Instance) InertiaDiagonal() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetInertiaDiagonal())
}

func (self Instance) SetInertiaDiagonal(value Vector3.XYZ) {
	class(self).SetInertiaDiagonal(gd.Vector3(value))
}

func (self Instance) InertiaOrientation() Quaternion.IJKX {
	return Quaternion.IJKX(class(self).GetInertiaOrientation())
}

func (self Instance) SetInertiaOrientation(value Quaternion.IJKX) {
	class(self).SetInertiaOrientation(gd.Quaternion(value))
}

func (self Instance) InertiaTensor() Basis.XYZ {
	return Basis.XYZ(class(self).GetInertiaTensor())
}

func (self Instance) SetInertiaTensor(value Basis.XYZ) {
	class(self).SetInertiaTensor(gd.Basis(value))
}

/*
Creates a new GLTFPhysicsBody instance from the given Godot [CollisionObject3D] node.
*/
//go:nosplit
func (self class) FromNode(body_node [1]gdclass.CollisionObject3D) [1]gdclass.GLTFPhysicsBody { //gd:GLTFPhysicsBody.from_node
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(body_node[0])[0])
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_from_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.GLTFPhysicsBody{gd.PointerWithOwnershipTransferredToGo[gdclass.GLTFPhysicsBody](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Converts this GLTFPhysicsBody instance into a Godot [CollisionObject3D] node.
*/
//go:nosplit
func (self class) ToNode() [1]gdclass.CollisionObject3D { //gd:GLTFPhysicsBody.to_node
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_to_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.CollisionObject3D{gd.PointerWithOwnershipTransferredToGo[gdclass.CollisionObject3D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Creates a new GLTFPhysicsBody instance by parsing the given [Dictionary] in the [code]OMI_physics_body[/code] GLTF extension format.
*/
//go:nosplit
func (self class) FromDictionary(dictionary Dictionary.Any) [1]gdclass.GLTFPhysicsBody { //gd:GLTFPhysicsBody.from_dictionary
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalDictionary(dictionary)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_from_dictionary, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.GLTFPhysicsBody{gd.PointerWithOwnershipTransferredToGo[gdclass.GLTFPhysicsBody](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Serializes this GLTFPhysicsBody instance into a [Dictionary]. It will be in the format expected by the [code]OMI_physics_body[/code] GLTF extension.
*/
//go:nosplit
func (self class) ToDictionary() Dictionary.Any { //gd:GLTFPhysicsBody.to_dictionary
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_to_dictionary, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetBodyType() String.Readable { //gd:GLTFPhysicsBody.get_body_type
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_get_body_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBodyType(body_type String.Readable) { //gd:GLTFPhysicsBody.set_body_type
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(body_type)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_set_body_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMass() gd.Float { //gd:GLTFPhysicsBody.get_mass
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_get_mass, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMass(mass gd.Float) { //gd:GLTFPhysicsBody.set_mass
	var frame = callframe.New()
	callframe.Arg(frame, mass)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_set_mass, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLinearVelocity() gd.Vector3 { //gd:GLTFPhysicsBody.get_linear_velocity
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_get_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLinearVelocity(linear_velocity gd.Vector3) { //gd:GLTFPhysicsBody.set_linear_velocity
	var frame = callframe.New()
	callframe.Arg(frame, linear_velocity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_set_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAngularVelocity() gd.Vector3 { //gd:GLTFPhysicsBody.get_angular_velocity
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_get_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAngularVelocity(angular_velocity gd.Vector3) { //gd:GLTFPhysicsBody.set_angular_velocity
	var frame = callframe.New()
	callframe.Arg(frame, angular_velocity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_set_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCenterOfMass() gd.Vector3 { //gd:GLTFPhysicsBody.get_center_of_mass
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_get_center_of_mass, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCenterOfMass(center_of_mass gd.Vector3) { //gd:GLTFPhysicsBody.set_center_of_mass
	var frame = callframe.New()
	callframe.Arg(frame, center_of_mass)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_set_center_of_mass, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetInertiaDiagonal() gd.Vector3 { //gd:GLTFPhysicsBody.get_inertia_diagonal
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_get_inertia_diagonal, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInertiaDiagonal(inertia_diagonal gd.Vector3) { //gd:GLTFPhysicsBody.set_inertia_diagonal
	var frame = callframe.New()
	callframe.Arg(frame, inertia_diagonal)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_set_inertia_diagonal, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetInertiaOrientation() gd.Quaternion { //gd:GLTFPhysicsBody.get_inertia_orientation
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Quaternion](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_get_inertia_orientation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInertiaOrientation(inertia_orientation gd.Quaternion) { //gd:GLTFPhysicsBody.set_inertia_orientation
	var frame = callframe.New()
	callframe.Arg(frame, inertia_orientation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_set_inertia_orientation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetInertiaTensor() gd.Basis { //gd:GLTFPhysicsBody.get_inertia_tensor
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Basis](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_get_inertia_tensor, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInertiaTensor(inertia_tensor gd.Basis) { //gd:GLTFPhysicsBody.set_inertia_tensor
	var frame = callframe.New()
	callframe.Arg(frame, inertia_tensor)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFPhysicsBody.Bind_set_inertia_tensor, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsGLTFPhysicsBody() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGLTFPhysicsBody() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("GLTFPhysicsBody", func(ptr gd.Object) any {
		return [1]gdclass.GLTFPhysicsBody{*(*gdclass.GLTFPhysicsBody)(unsafe.Pointer(&ptr))}
	})
}

type Structure map[interface{}]interface{}
