// Package SkeletonModification2DJiggle provides methods for working with SkeletonModification2DJiggle object instances.
package SkeletonModification2DJiggle

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Resource"
import "graphics.gd/classdb/SkeletonModification2D"
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
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
This modification moves a series of bones, typically called a bone chain, towards a target. What makes this modification special is that it calculates the velocity and acceleration for each bone in the bone chain, and runs a very light physics-like calculation using the inputted values. This allows the bones to overshoot the target and "jiggle" around. It can be configured to act more like a spring, or sway around like cloth might.
This modification is useful for adding additional motion to things like hair, the edges of clothing, and more. It has several settings to that allow control over how the joint moves when the target moves.
[b]Note:[/b] The Jiggle modifier has [code]jiggle_joints[/code], which are the data objects that hold the data for each joint in the Jiggle chain. This is different from than [Bone2D] nodes! Jiggle joints hold the data needed for each [Bone2D] in the bone chain used by the Jiggle modification.
*/
type Instance [1]gdclass.SkeletonModification2DJiggle

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsSkeletonModification2DJiggle() Instance
}

/*
If [code]true[/code], the Jiggle modifier will take colliders into account, keeping them from entering into these collision objects.
*/
func (self Instance) SetUseColliders(use_colliders bool) { //gd:SkeletonModification2DJiggle.set_use_colliders
	class(self).SetUseColliders(use_colliders)
}

/*
Returns whether the jiggle modifier is taking physics colliders into account when solving.
*/
func (self Instance) GetUseColliders() bool { //gd:SkeletonModification2DJiggle.get_use_colliders
	return bool(class(self).GetUseColliders())
}

/*
Sets the collision mask that the Jiggle modifier will use when reacting to colliders, if the Jiggle modifier is set to take colliders into account.
*/
func (self Instance) SetCollisionMask(collision_mask int) { //gd:SkeletonModification2DJiggle.set_collision_mask
	class(self).SetCollisionMask(int64(collision_mask))
}

/*
Returns the collision mask used by the Jiggle modifier when collisions are enabled.
*/
func (self Instance) GetCollisionMask() int { //gd:SkeletonModification2DJiggle.get_collision_mask
	return int(int(class(self).GetCollisionMask()))
}

/*
Sets the [Bone2D] node assigned to the Jiggle joint at [param joint_idx].
*/
func (self Instance) SetJiggleJointBone2dNode(joint_idx int, bone2d_node string) { //gd:SkeletonModification2DJiggle.set_jiggle_joint_bone2d_node
	class(self).SetJiggleJointBone2dNode(int64(joint_idx), Path.ToNode(String.New(bone2d_node)))
}

/*
Returns the [Bone2D] node assigned to the Jiggle joint at [param joint_idx].
*/
func (self Instance) GetJiggleJointBone2dNode(joint_idx int) string { //gd:SkeletonModification2DJiggle.get_jiggle_joint_bone2d_node
	return string(class(self).GetJiggleJointBone2dNode(int64(joint_idx)).String())
}

/*
Sets the bone index, [param bone_idx], of the Jiggle joint at [param joint_idx]. When possible, this will also update the [code]bone2d_node[/code] of the Jiggle joint based on data provided by the linked skeleton.
*/
func (self Instance) SetJiggleJointBoneIndex(joint_idx int, bone_idx int) { //gd:SkeletonModification2DJiggle.set_jiggle_joint_bone_index
	class(self).SetJiggleJointBoneIndex(int64(joint_idx), int64(bone_idx))
}

/*
Returns the index of the [Bone2D] node assigned to the Jiggle joint at [param joint_idx].
*/
func (self Instance) GetJiggleJointBoneIndex(joint_idx int) int { //gd:SkeletonModification2DJiggle.get_jiggle_joint_bone_index
	return int(int(class(self).GetJiggleJointBoneIndex(int64(joint_idx))))
}

/*
Sets whether the Jiggle joint at [param joint_idx] should override the default Jiggle joint settings. Setting this to [code]true[/code] will make the joint use its own settings rather than the default ones attached to the modification.
*/
func (self Instance) SetJiggleJointOverride(joint_idx int, override bool) { //gd:SkeletonModification2DJiggle.set_jiggle_joint_override
	class(self).SetJiggleJointOverride(int64(joint_idx), override)
}

/*
Returns a boolean that indicates whether the joint at [param joint_idx] is overriding the default Jiggle joint data defined in the modification.
*/
func (self Instance) GetJiggleJointOverride(joint_idx int) bool { //gd:SkeletonModification2DJiggle.get_jiggle_joint_override
	return bool(class(self).GetJiggleJointOverride(int64(joint_idx)))
}

/*
Sets the of stiffness of the Jiggle joint at [param joint_idx].
*/
func (self Instance) SetJiggleJointStiffness(joint_idx int, stiffness Float.X) { //gd:SkeletonModification2DJiggle.set_jiggle_joint_stiffness
	class(self).SetJiggleJointStiffness(int64(joint_idx), float64(stiffness))
}

/*
Returns the stiffness of the Jiggle joint at [param joint_idx].
*/
func (self Instance) GetJiggleJointStiffness(joint_idx int) Float.X { //gd:SkeletonModification2DJiggle.get_jiggle_joint_stiffness
	return Float.X(Float.X(class(self).GetJiggleJointStiffness(int64(joint_idx))))
}

/*
Sets the of mass of the Jiggle joint at [param joint_idx].
*/
func (self Instance) SetJiggleJointMass(joint_idx int, mass Float.X) { //gd:SkeletonModification2DJiggle.set_jiggle_joint_mass
	class(self).SetJiggleJointMass(int64(joint_idx), float64(mass))
}

/*
Returns the amount of mass of the jiggle joint at [param joint_idx].
*/
func (self Instance) GetJiggleJointMass(joint_idx int) Float.X { //gd:SkeletonModification2DJiggle.get_jiggle_joint_mass
	return Float.X(Float.X(class(self).GetJiggleJointMass(int64(joint_idx))))
}

/*
Sets the amount of damping of the Jiggle joint at [param joint_idx].
*/
func (self Instance) SetJiggleJointDamping(joint_idx int, damping Float.X) { //gd:SkeletonModification2DJiggle.set_jiggle_joint_damping
	class(self).SetJiggleJointDamping(int64(joint_idx), float64(damping))
}

/*
Returns the amount of damping of the Jiggle joint at [param joint_idx].
*/
func (self Instance) GetJiggleJointDamping(joint_idx int) Float.X { //gd:SkeletonModification2DJiggle.get_jiggle_joint_damping
	return Float.X(Float.X(class(self).GetJiggleJointDamping(int64(joint_idx))))
}

/*
Sets whether the Jiggle joint at [param joint_idx] should use gravity.
*/
func (self Instance) SetJiggleJointUseGravity(joint_idx int, use_gravity bool) { //gd:SkeletonModification2DJiggle.set_jiggle_joint_use_gravity
	class(self).SetJiggleJointUseGravity(int64(joint_idx), use_gravity)
}

/*
Returns a boolean that indicates whether the joint at [param joint_idx] is using gravity or not.
*/
func (self Instance) GetJiggleJointUseGravity(joint_idx int) bool { //gd:SkeletonModification2DJiggle.get_jiggle_joint_use_gravity
	return bool(class(self).GetJiggleJointUseGravity(int64(joint_idx)))
}

/*
Sets the gravity vector of the Jiggle joint at [param joint_idx].
*/
func (self Instance) SetJiggleJointGravity(joint_idx int, gravity Vector2.XY) { //gd:SkeletonModification2DJiggle.set_jiggle_joint_gravity
	class(self).SetJiggleJointGravity(int64(joint_idx), Vector2.XY(gravity))
}

/*
Returns a [Vector2] representing the amount of gravity the Jiggle joint at [param joint_idx] is influenced by.
*/
func (self Instance) GetJiggleJointGravity(joint_idx int) Vector2.XY { //gd:SkeletonModification2DJiggle.get_jiggle_joint_gravity
	return Vector2.XY(class(self).GetJiggleJointGravity(int64(joint_idx)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.SkeletonModification2DJiggle

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SkeletonModification2DJiggle"))
	casted := Instance{*(*gdclass.SkeletonModification2DJiggle)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) TargetNodepath() string {
	return string(class(self).GetTargetNode().String())
}

func (self Instance) SetTargetNodepath(value string) {
	class(self).SetTargetNode(Path.ToNode(String.New(value)))
}

func (self Instance) JiggleDataChainLength() int {
	return int(int(class(self).GetJiggleDataChainLength()))
}

func (self Instance) SetJiggleDataChainLength(value int) {
	class(self).SetJiggleDataChainLength(int64(value))
}

func (self Instance) Stiffness() Float.X {
	return Float.X(Float.X(class(self).GetStiffness()))
}

func (self Instance) SetStiffness(value Float.X) {
	class(self).SetStiffness(float64(value))
}

func (self Instance) Mass() Float.X {
	return Float.X(Float.X(class(self).GetMass()))
}

func (self Instance) SetMass(value Float.X) {
	class(self).SetMass(float64(value))
}

func (self Instance) Damping() Float.X {
	return Float.X(Float.X(class(self).GetDamping()))
}

func (self Instance) SetDamping(value Float.X) {
	class(self).SetDamping(float64(value))
}

func (self Instance) UseGravity() bool {
	return bool(class(self).GetUseGravity())
}

func (self Instance) SetUseGravity(value bool) {
	class(self).SetUseGravity(value)
}

func (self Instance) Gravity() Vector2.XY {
	return Vector2.XY(class(self).GetGravity())
}

func (self Instance) SetGravity(value Vector2.XY) {
	class(self).SetGravity(Vector2.XY(value))
}

//go:nosplit
func (self class) SetTargetNode(target_nodepath Path.ToNode) { //gd:SkeletonModification2DJiggle.set_target_node
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(target_nodepath)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_target_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTargetNode() Path.ToNode { //gd:SkeletonModification2DJiggle.get_target_node
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_target_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Path.ToNode(String.Via(gd.NodePathProxy{}, pointers.Pack(pointers.New[gd.NodePath](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetJiggleDataChainLength(length int64) { //gd:SkeletonModification2DJiggle.set_jiggle_data_chain_length
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_jiggle_data_chain_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetJiggleDataChainLength() int64 { //gd:SkeletonModification2DJiggle.get_jiggle_data_chain_length
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_jiggle_data_chain_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStiffness(stiffness float64) { //gd:SkeletonModification2DJiggle.set_stiffness
	var frame = callframe.New()
	callframe.Arg(frame, stiffness)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_stiffness, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStiffness() float64 { //gd:SkeletonModification2DJiggle.get_stiffness
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_stiffness, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMass(mass float64) { //gd:SkeletonModification2DJiggle.set_mass
	var frame = callframe.New()
	callframe.Arg(frame, mass)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_mass, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMass() float64 { //gd:SkeletonModification2DJiggle.get_mass
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_mass, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDamping(damping float64) { //gd:SkeletonModification2DJiggle.set_damping
	var frame = callframe.New()
	callframe.Arg(frame, damping)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_damping, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDamping() float64 { //gd:SkeletonModification2DJiggle.get_damping
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_damping, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseGravity(use_gravity bool) { //gd:SkeletonModification2DJiggle.set_use_gravity
	var frame = callframe.New()
	callframe.Arg(frame, use_gravity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_use_gravity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUseGravity() bool { //gd:SkeletonModification2DJiggle.get_use_gravity
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_use_gravity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGravity(gravity Vector2.XY) { //gd:SkeletonModification2DJiggle.set_gravity
	var frame = callframe.New()
	callframe.Arg(frame, gravity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_gravity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGravity() Vector2.XY { //gd:SkeletonModification2DJiggle.get_gravity
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_gravity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [code]true[/code], the Jiggle modifier will take colliders into account, keeping them from entering into these collision objects.
*/
//go:nosplit
func (self class) SetUseColliders(use_colliders bool) { //gd:SkeletonModification2DJiggle.set_use_colliders
	var frame = callframe.New()
	callframe.Arg(frame, use_colliders)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_use_colliders, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether the jiggle modifier is taking physics colliders into account when solving.
*/
//go:nosplit
func (self class) GetUseColliders() bool { //gd:SkeletonModification2DJiggle.get_use_colliders
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_use_colliders, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the collision mask that the Jiggle modifier will use when reacting to colliders, if the Jiggle modifier is set to take colliders into account.
*/
//go:nosplit
func (self class) SetCollisionMask(collision_mask int64) { //gd:SkeletonModification2DJiggle.set_collision_mask
	var frame = callframe.New()
	callframe.Arg(frame, collision_mask)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the collision mask used by the Jiggle modifier when collisions are enabled.
*/
//go:nosplit
func (self class) GetCollisionMask() int64 { //gd:SkeletonModification2DJiggle.get_collision_mask
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [Bone2D] node assigned to the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetJiggleJointBone2dNode(joint_idx int64, bone2d_node Path.ToNode) { //gd:SkeletonModification2DJiggle.set_jiggle_joint_bone2d_node
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(bone2d_node)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [Bone2D] node assigned to the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetJiggleJointBone2dNode(joint_idx int64) Path.ToNode { //gd:SkeletonModification2DJiggle.get_jiggle_joint_bone2d_node
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_bone2d_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Path.ToNode(String.Via(gd.NodePathProxy{}, pointers.Pack(pointers.New[gd.NodePath](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Sets the bone index, [param bone_idx], of the Jiggle joint at [param joint_idx]. When possible, this will also update the [code]bone2d_node[/code] of the Jiggle joint based on data provided by the linked skeleton.
*/
//go:nosplit
func (self class) SetJiggleJointBoneIndex(joint_idx int64, bone_idx int64) { //gd:SkeletonModification2DJiggle.set_jiggle_joint_bone_index
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_bone_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the index of the [Bone2D] node assigned to the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetJiggleJointBoneIndex(joint_idx int64) int64 { //gd:SkeletonModification2DJiggle.get_jiggle_joint_bone_index
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_bone_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets whether the Jiggle joint at [param joint_idx] should override the default Jiggle joint settings. Setting this to [code]true[/code] will make the joint use its own settings rather than the default ones attached to the modification.
*/
//go:nosplit
func (self class) SetJiggleJointOverride(joint_idx int64, override bool) { //gd:SkeletonModification2DJiggle.set_jiggle_joint_override
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, override)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a boolean that indicates whether the joint at [param joint_idx] is overriding the default Jiggle joint data defined in the modification.
*/
//go:nosplit
func (self class) GetJiggleJointOverride(joint_idx int64) bool { //gd:SkeletonModification2DJiggle.get_jiggle_joint_override
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the of stiffness of the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetJiggleJointStiffness(joint_idx int64, stiffness float64) { //gd:SkeletonModification2DJiggle.set_jiggle_joint_stiffness
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, stiffness)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_stiffness, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the stiffness of the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetJiggleJointStiffness(joint_idx int64) float64 { //gd:SkeletonModification2DJiggle.get_jiggle_joint_stiffness
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_stiffness, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the of mass of the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetJiggleJointMass(joint_idx int64, mass float64) { //gd:SkeletonModification2DJiggle.set_jiggle_joint_mass
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, mass)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_mass, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the amount of mass of the jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetJiggleJointMass(joint_idx int64) float64 { //gd:SkeletonModification2DJiggle.get_jiggle_joint_mass
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_mass, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the amount of damping of the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetJiggleJointDamping(joint_idx int64, damping float64) { //gd:SkeletonModification2DJiggle.set_jiggle_joint_damping
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, damping)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_damping, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the amount of damping of the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) GetJiggleJointDamping(joint_idx int64) float64 { //gd:SkeletonModification2DJiggle.get_jiggle_joint_damping
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_damping, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets whether the Jiggle joint at [param joint_idx] should use gravity.
*/
//go:nosplit
func (self class) SetJiggleJointUseGravity(joint_idx int64, use_gravity bool) { //gd:SkeletonModification2DJiggle.set_jiggle_joint_use_gravity
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, use_gravity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_use_gravity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a boolean that indicates whether the joint at [param joint_idx] is using gravity or not.
*/
//go:nosplit
func (self class) GetJiggleJointUseGravity(joint_idx int64) bool { //gd:SkeletonModification2DJiggle.get_jiggle_joint_use_gravity
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_use_gravity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the gravity vector of the Jiggle joint at [param joint_idx].
*/
//go:nosplit
func (self class) SetJiggleJointGravity(joint_idx int64, gravity Vector2.XY) { //gd:SkeletonModification2DJiggle.set_jiggle_joint_gravity
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	callframe.Arg(frame, gravity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_set_jiggle_joint_gravity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a [Vector2] representing the amount of gravity the Jiggle joint at [param joint_idx] is influenced by.
*/
//go:nosplit
func (self class) GetJiggleJointGravity(joint_idx int64) Vector2.XY { //gd:SkeletonModification2DJiggle.get_jiggle_joint_gravity
	var frame = callframe.New()
	callframe.Arg(frame, joint_idx)
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModification2DJiggle.Bind_get_jiggle_joint_gravity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsSkeletonModification2DJiggle() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsSkeletonModification2DJiggle() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsSkeletonModification2D() SkeletonModification2D.Advanced {
	return *((*SkeletonModification2D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsSkeletonModification2D() SkeletonModification2D.Instance {
	return *((*SkeletonModification2D.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(SkeletonModification2D.Advanced(self.AsSkeletonModification2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(SkeletonModification2D.Instance(self.AsSkeletonModification2D()), name)
	}
}
func init() {
	gdclass.Register("SkeletonModification2DJiggle", func(ptr gd.Object) any {
		return [1]gdclass.SkeletonModification2DJiggle{*(*gdclass.SkeletonModification2DJiggle)(unsafe.Pointer(&ptr))}
	})
}
