// Package Skeleton2D provides methods for working with Skeleton2D object instances.
package Skeleton2D

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
import "graphics.gd/classdb/Node2D"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Transform2D"

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

/*
[Skeleton2D] parents a hierarchy of [Bone2D] nodes. It holds a reference to each [Bone2D]'s rest pose and acts as a single point of access to its bones.
To set up different types of inverse kinematics for the given Skeleton2D, a [SkeletonModificationStack2D] should be created. The inverse kinematics be applied by increasing [member SkeletonModificationStack2D.modification_count] and creating the desired number of modifications.
*/
type Instance [1]gdclass.Skeleton2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsSkeleton2D() Instance
}

/*
Returns the number of [Bone2D] nodes in the node hierarchy parented by Skeleton2D.
*/
func (self Instance) GetBoneCount() int { //gd:Skeleton2D.get_bone_count
	return int(int(class(self).GetBoneCount()))
}

/*
Returns a [Bone2D] from the node hierarchy parented by Skeleton2D. The object to return is identified by the parameter [param idx]. Bones are indexed by descending the node hierarchy from top to bottom, adding the children of each branch before moving to the next sibling.
*/
func (self Instance) GetBone(idx int) [1]gdclass.Bone2D { //gd:Skeleton2D.get_bone
	return [1]gdclass.Bone2D(class(self).GetBone(gd.Int(idx)))
}

/*
Returns the [RID] of a Skeleton2D instance.
*/
func (self Instance) GetSkeleton() RID.Skeleton { //gd:Skeleton2D.get_skeleton
	return RID.Skeleton(class(self).GetSkeleton())
}

/*
Sets the [SkeletonModificationStack2D] attached to this skeleton.
*/
func (self Instance) SetModificationStack(modification_stack [1]gdclass.SkeletonModificationStack2D) { //gd:Skeleton2D.set_modification_stack
	class(self).SetModificationStack(modification_stack)
}

/*
Returns the [SkeletonModificationStack2D] attached to this skeleton, if one exists.
*/
func (self Instance) GetModificationStack() [1]gdclass.SkeletonModificationStack2D { //gd:Skeleton2D.get_modification_stack
	return [1]gdclass.SkeletonModificationStack2D(class(self).GetModificationStack())
}

/*
Executes all the modifications on the [SkeletonModificationStack2D], if the Skeleton2D has one assigned.
*/
func (self Instance) ExecuteModifications(delta Float.X, execution_mode int) { //gd:Skeleton2D.execute_modifications
	class(self).ExecuteModifications(gd.Float(delta), gd.Int(execution_mode))
}

/*
Sets the local pose transform, [param override_pose], for the bone at [param bone_idx].
[param strength] is the interpolation strength that will be used when applying the pose, and [param persistent] determines if the applied pose will remain.
[b]Note:[/b] The pose transform needs to be a local transform relative to the [Bone2D] node at [param bone_idx]!
*/
func (self Instance) SetBoneLocalPoseOverride(bone_idx int, override_pose Transform2D.OriginXY, strength Float.X, persistent bool) { //gd:Skeleton2D.set_bone_local_pose_override
	class(self).SetBoneLocalPoseOverride(gd.Int(bone_idx), gd.Transform2D(override_pose), gd.Float(strength), persistent)
}

/*
Returns the local pose override transform for [param bone_idx].
*/
func (self Instance) GetBoneLocalPoseOverride(bone_idx int) Transform2D.OriginXY { //gd:Skeleton2D.get_bone_local_pose_override
	return Transform2D.OriginXY(class(self).GetBoneLocalPoseOverride(gd.Int(bone_idx)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Skeleton2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Skeleton2D"))
	casted := Instance{*(*gdclass.Skeleton2D)(unsafe.Pointer(&object))}
	return casted
}

/*
Returns the number of [Bone2D] nodes in the node hierarchy parented by Skeleton2D.
*/
//go:nosplit
func (self class) GetBoneCount() gd.Int { //gd:Skeleton2D.get_bone_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton2D.Bind_get_bone_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a [Bone2D] from the node hierarchy parented by Skeleton2D. The object to return is identified by the parameter [param idx]. Bones are indexed by descending the node hierarchy from top to bottom, adding the children of each branch before moving to the next sibling.
*/
//go:nosplit
func (self class) GetBone(idx gd.Int) [1]gdclass.Bone2D { //gd:Skeleton2D.get_bone
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton2D.Bind_get_bone, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Bone2D{gd.PointerMustAssertInstanceID[gdclass.Bone2D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the [RID] of a Skeleton2D instance.
*/
//go:nosplit
func (self class) GetSkeleton() gd.RID { //gd:Skeleton2D.get_skeleton
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton2D.Bind_get_skeleton, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [SkeletonModificationStack2D] attached to this skeleton.
*/
//go:nosplit
func (self class) SetModificationStack(modification_stack [1]gdclass.SkeletonModificationStack2D) { //gd:Skeleton2D.set_modification_stack
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(modification_stack[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton2D.Bind_set_modification_stack, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [SkeletonModificationStack2D] attached to this skeleton, if one exists.
*/
//go:nosplit
func (self class) GetModificationStack() [1]gdclass.SkeletonModificationStack2D { //gd:Skeleton2D.get_modification_stack
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton2D.Bind_get_modification_stack, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.SkeletonModificationStack2D{gd.PointerWithOwnershipTransferredToGo[gdclass.SkeletonModificationStack2D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Executes all the modifications on the [SkeletonModificationStack2D], if the Skeleton2D has one assigned.
*/
//go:nosplit
func (self class) ExecuteModifications(delta gd.Float, execution_mode gd.Int) { //gd:Skeleton2D.execute_modifications
	var frame = callframe.New()
	callframe.Arg(frame, delta)
	callframe.Arg(frame, execution_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton2D.Bind_execute_modifications, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the local pose transform, [param override_pose], for the bone at [param bone_idx].
[param strength] is the interpolation strength that will be used when applying the pose, and [param persistent] determines if the applied pose will remain.
[b]Note:[/b] The pose transform needs to be a local transform relative to the [Bone2D] node at [param bone_idx]!
*/
//go:nosplit
func (self class) SetBoneLocalPoseOverride(bone_idx gd.Int, override_pose gd.Transform2D, strength gd.Float, persistent bool) { //gd:Skeleton2D.set_bone_local_pose_override
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, override_pose)
	callframe.Arg(frame, strength)
	callframe.Arg(frame, persistent)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton2D.Bind_set_bone_local_pose_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the local pose override transform for [param bone_idx].
*/
//go:nosplit
func (self class) GetBoneLocalPoseOverride(bone_idx gd.Int) gd.Transform2D { //gd:Skeleton2D.get_bone_local_pose_override
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Skeleton2D.Bind_get_bone_local_pose_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnBoneSetupChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("bone_setup_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsSkeleton2D() Advanced       { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSkeleton2D() Instance    { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(Node2D.Advanced(self.AsNode2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node2D.Instance(self.AsNode2D()), name)
	}
}
func init() {
	gdclass.Register("Skeleton2D", func(ptr gd.Object) any { return [1]gdclass.Skeleton2D{*(*gdclass.Skeleton2D)(unsafe.Pointer(&ptr))} })
}
