// Package PhysicalBone2D provides methods for working with PhysicalBone2D object instances.
package PhysicalBone2D

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/CollisionObject2D"
import "graphics.gd/classdb/Node"
import "graphics.gd/classdb/Node2D"
import "graphics.gd/classdb/PhysicsBody2D"
import "graphics.gd/classdb/RigidBody2D"
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
The [PhysicalBone2D] node is a [RigidBody2D]-based node that can be used to make [Bone2D]s in a [Skeleton2D] react to physics.
[b]Note:[/b] To make the [Bone2D]s visually follow the [PhysicalBone2D] node, use a [SkeletonModification2DPhysicalBones] modification on the [Skeleton2D] parent.
[b]Note:[/b] The [PhysicalBone2D] node does not automatically create a [Joint2D] node to keep [PhysicalBone2D] nodes together. They must be created manually. For most cases, you want to use a [PinJoint2D] node. The [PhysicalBone2D] node will automatically configure the [Joint2D] node once it's been added as a child node.
*/
type Instance [1]gdclass.PhysicalBone2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPhysicalBone2D() Instance
}

/*
Returns the first [Joint2D] child node, if one exists. This is mainly a helper function to make it easier to get the [Joint2D] that the [PhysicalBone2D] is autoconfiguring.
*/
func (self Instance) GetJoint() [1]gdclass.Joint2D { //gd:PhysicalBone2D.get_joint
	return [1]gdclass.Joint2D(class(self).GetJoint())
}

/*
Returns a boolean that indicates whether the [PhysicalBone2D] is running and simulating using the Godot 2D physics engine. When [code]true[/code], the PhysicalBone2D node is using physics.
*/
func (self Instance) IsSimulatingPhysics() bool { //gd:PhysicalBone2D.is_simulating_physics
	return bool(class(self).IsSimulatingPhysics())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PhysicalBone2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicalBone2D"))
	casted := Instance{*(*gdclass.PhysicalBone2D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Bone2dNodepath() string {
	return string(class(self).GetBone2dNodepath().String())
}

func (self Instance) SetBone2dNodepath(value string) {
	class(self).SetBone2dNodepath(Path.ToNode(String.New(value)))
}

func (self Instance) Bone2dIndex() int {
	return int(int(class(self).GetBone2dIndex()))
}

func (self Instance) SetBone2dIndex(value int) {
	class(self).SetBone2dIndex(int64(value))
}

func (self Instance) AutoConfigureJoint() bool {
	return bool(class(self).GetAutoConfigureJoint())
}

func (self Instance) SetAutoConfigureJoint(value bool) {
	class(self).SetAutoConfigureJoint(value)
}

func (self Instance) SimulatePhysics() bool {
	return bool(class(self).GetSimulatePhysics())
}

func (self Instance) SetSimulatePhysics(value bool) {
	class(self).SetSimulatePhysics(value)
}

func (self Instance) FollowBoneWhenSimulating() bool {
	return bool(class(self).GetFollowBoneWhenSimulating())
}

func (self Instance) SetFollowBoneWhenSimulating(value bool) {
	class(self).SetFollowBoneWhenSimulating(value)
}

/*
Returns the first [Joint2D] child node, if one exists. This is mainly a helper function to make it easier to get the [Joint2D] that the [PhysicalBone2D] is autoconfiguring.
*/
//go:nosplit
func (self class) GetJoint() [1]gdclass.Joint2D { //gd:PhysicalBone2D.get_joint
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone2D.Bind_get_joint, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Joint2D{gd.PointerMustAssertInstanceID[gdclass.Joint2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetAutoConfigureJoint() bool { //gd:PhysicalBone2D.get_auto_configure_joint
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone2D.Bind_get_auto_configure_joint, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoConfigureJoint(auto_configure_joint bool) { //gd:PhysicalBone2D.set_auto_configure_joint
	var frame = callframe.New()
	callframe.Arg(frame, auto_configure_joint)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone2D.Bind_set_auto_configure_joint, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetSimulatePhysics(simulate_physics bool) { //gd:PhysicalBone2D.set_simulate_physics
	var frame = callframe.New()
	callframe.Arg(frame, simulate_physics)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone2D.Bind_set_simulate_physics, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSimulatePhysics() bool { //gd:PhysicalBone2D.get_simulate_physics
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone2D.Bind_get_simulate_physics, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a boolean that indicates whether the [PhysicalBone2D] is running and simulating using the Godot 2D physics engine. When [code]true[/code], the PhysicalBone2D node is using physics.
*/
//go:nosplit
func (self class) IsSimulatingPhysics() bool { //gd:PhysicalBone2D.is_simulating_physics
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone2D.Bind_is_simulating_physics, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBone2dNodepath(nodepath Path.ToNode) { //gd:PhysicalBone2D.set_bone2d_nodepath
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(nodepath)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone2D.Bind_set_bone2d_nodepath, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBone2dNodepath() Path.ToNode { //gd:PhysicalBone2D.get_bone2d_nodepath
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone2D.Bind_get_bone2d_nodepath, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Path.ToNode(String.Via(gd.NodePathProxy{}, pointers.Pack(pointers.New[gd.NodePath](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBone2dIndex(bone_index int64) { //gd:PhysicalBone2D.set_bone2d_index
	var frame = callframe.New()
	callframe.Arg(frame, bone_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone2D.Bind_set_bone2d_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBone2dIndex() int64 { //gd:PhysicalBone2D.get_bone2d_index
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone2D.Bind_get_bone2d_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFollowBoneWhenSimulating(follow_bone bool) { //gd:PhysicalBone2D.set_follow_bone_when_simulating
	var frame = callframe.New()
	callframe.Arg(frame, follow_bone)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone2D.Bind_set_follow_bone_when_simulating, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFollowBoneWhenSimulating() bool { //gd:PhysicalBone2D.get_follow_bone_when_simulating
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicalBone2D.Bind_get_follow_bone_when_simulating, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPhysicalBone2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPhysicalBone2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRigidBody2D() RigidBody2D.Advanced {
	return *((*RigidBody2D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsRigidBody2D() RigidBody2D.Instance {
	return *((*RigidBody2D.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(RigidBody2D.Advanced(self.AsRigidBody2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RigidBody2D.Instance(self.AsRigidBody2D()), name)
	}
}
func init() {
	gdclass.Register("PhysicalBone2D", func(ptr gd.Object) any {
		return [1]gdclass.PhysicalBone2D{*(*gdclass.PhysicalBone2D)(unsafe.Pointer(&ptr))}
	})
}
