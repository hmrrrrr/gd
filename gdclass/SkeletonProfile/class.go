package SkeletonProfile

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
This resource is used in [EditorScenePostImport]. Some parameters are referring to bones in [Skeleton3D], [Skin], [Animation], and some other nodes are rewritten based on the parameters of [SkeletonProfile].
[b]Note:[/b] These parameters need to be set only when creating a custom profile. In [SkeletonProfileHumanoid], they are defined internally as read-only values.

*/
type Go [1]classdb.SkeletonProfile

/*
Returns the name of the group at [param group_idx] that will be the drawing group in the [BoneMap] editor.
*/
func (self Go) GetGroupName(group_idx int) string {
	return string(class(self).GetGroupName(gd.Int(group_idx)).String())
}

/*
Sets the name of the group at [param group_idx] that will be the drawing group in the [BoneMap] editor.
*/
func (self Go) SetGroupName(group_idx int, group_name string) {
	class(self).SetGroupName(gd.Int(group_idx), gd.NewStringName(group_name))
}

/*
Returns the texture of the group at [param group_idx] that will be the drawing group background image in the [BoneMap] editor.
*/
func (self Go) GetTexture(group_idx int) gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetTexture(gd.Int(group_idx)))
}

/*
Sets the texture of the group at [param group_idx] that will be the drawing group background image in the [BoneMap] editor.
*/
func (self Go) SetTexture(group_idx int, texture gdclass.Texture2D) {
	class(self).SetTexture(gd.Int(group_idx), texture)
}

/*
Returns the bone index that matches [param bone_name] as its name.
*/
func (self Go) FindBone(bone_name string) int {
	return int(int(class(self).FindBone(gd.NewStringName(bone_name))))
}

/*
Returns the name of the bone at [param bone_idx] that will be the key name in the [BoneMap].
In the retargeting process, the returned bone name is the bone name of the target skeleton.
*/
func (self Go) GetBoneName(bone_idx int) string {
	return string(class(self).GetBoneName(gd.Int(bone_idx)).String())
}

/*
Sets the name of the bone at [param bone_idx] that will be the key name in the [BoneMap].
In the retargeting process, the setting bone name is the bone name of the target skeleton.
*/
func (self Go) SetBoneName(bone_idx int, bone_name string) {
	class(self).SetBoneName(gd.Int(bone_idx), gd.NewStringName(bone_name))
}

/*
Returns the name of the bone which is the parent to the bone at [param bone_idx]. The result is empty if the bone has no parent.
*/
func (self Go) GetBoneParent(bone_idx int) string {
	return string(class(self).GetBoneParent(gd.Int(bone_idx)).String())
}

/*
Sets the bone with name [param bone_parent] as the parent of the bone at [param bone_idx]. If an empty string is passed, then the bone has no parent.
*/
func (self Go) SetBoneParent(bone_idx int, bone_parent string) {
	class(self).SetBoneParent(gd.Int(bone_idx), gd.NewStringName(bone_parent))
}

/*
Returns the tail direction of the bone at [param bone_idx].
*/
func (self Go) GetTailDirection(bone_idx int) classdb.SkeletonProfileTailDirection {
	return classdb.SkeletonProfileTailDirection(class(self).GetTailDirection(gd.Int(bone_idx)))
}

/*
Sets the tail direction of the bone at [param bone_idx].
[b]Note:[/b] This only specifies the method of calculation. The actual coordinates required should be stored in an external skeleton, so the calculation itself needs to be done externally.
*/
func (self Go) SetTailDirection(bone_idx int, tail_direction classdb.SkeletonProfileTailDirection) {
	class(self).SetTailDirection(gd.Int(bone_idx), tail_direction)
}

/*
Returns the name of the bone which is the tail of the bone at [param bone_idx].
*/
func (self Go) GetBoneTail(bone_idx int) string {
	return string(class(self).GetBoneTail(gd.Int(bone_idx)).String())
}

/*
Sets the bone with name [param bone_tail] as the tail of the bone at [param bone_idx].
*/
func (self Go) SetBoneTail(bone_idx int, bone_tail string) {
	class(self).SetBoneTail(gd.Int(bone_idx), gd.NewStringName(bone_tail))
}

/*
Returns the reference pose transform for bone [param bone_idx].
*/
func (self Go) GetReferencePose(bone_idx int) gd.Transform3D {
	return gd.Transform3D(class(self).GetReferencePose(gd.Int(bone_idx)))
}

/*
Sets the reference pose transform for bone [param bone_idx].
*/
func (self Go) SetReferencePose(bone_idx int, bone_name gd.Transform3D) {
	class(self).SetReferencePose(gd.Int(bone_idx), bone_name)
}

/*
Returns the offset of the bone at [param bone_idx] that will be the button position in the [BoneMap] editor.
This is the offset with origin at the top left corner of the square.
*/
func (self Go) GetHandleOffset(bone_idx int) gd.Vector2 {
	return gd.Vector2(class(self).GetHandleOffset(gd.Int(bone_idx)))
}

/*
Sets the offset of the bone at [param bone_idx] that will be the button position in the [BoneMap] editor.
This is the offset with origin at the top left corner of the square.
*/
func (self Go) SetHandleOffset(bone_idx int, handle_offset gd.Vector2) {
	class(self).SetHandleOffset(gd.Int(bone_idx), handle_offset)
}

/*
Returns the group of the bone at [param bone_idx].
*/
func (self Go) GetGroup(bone_idx int) string {
	return string(class(self).GetGroup(gd.Int(bone_idx)).String())
}

/*
Sets the group of the bone at [param bone_idx].
*/
func (self Go) SetGroup(bone_idx int, group string) {
	class(self).SetGroup(gd.Int(bone_idx), gd.NewStringName(group))
}

/*
Returns whether the bone at [param bone_idx] is required for retargeting.
This value is used by the bone map editor. If this method returns [code]true[/code], and no bone is assigned, the handle color will be red on the bone map editor.
*/
func (self Go) IsRequired(bone_idx int) bool {
	return bool(class(self).IsRequired(gd.Int(bone_idx)))
}

/*
Sets the required status for bone [param bone_idx] to [param required].
*/
func (self Go) SetRequired(bone_idx int, required bool) {
	class(self).SetRequired(gd.Int(bone_idx), required)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.SkeletonProfile
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SkeletonProfile"))
	return Go{classdb.SkeletonProfile(object)}
}

func (self Go) RootBone() string {
		return string(class(self).GetRootBone().String())
}

func (self Go) SetRootBone(value string) {
	class(self).SetRootBone(gd.NewStringName(value))
}

func (self Go) ScaleBaseBone() string {
		return string(class(self).GetScaleBaseBone().String())
}

func (self Go) SetScaleBaseBone(value string) {
	class(self).SetScaleBaseBone(gd.NewStringName(value))
}

func (self Go) GroupSize() int {
		return int(int(class(self).GetGroupSize()))
}

func (self Go) SetGroupSize(value int) {
	class(self).SetGroupSize(gd.Int(value))
}

func (self Go) BoneSize() int {
		return int(int(class(self).GetBoneSize()))
}

func (self Go) SetBoneSize(value int) {
	class(self).SetBoneSize(gd.Int(value))
}

//go:nosplit
func (self class) SetRootBone(bone_name gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(bone_name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonProfile.Bind_set_root_bone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRootBone() gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonProfile.Bind_get_root_bone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetScaleBaseBone(bone_name gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(bone_name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonProfile.Bind_set_scale_base_bone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetScaleBaseBone() gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonProfile.Bind_get_scale_base_bone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGroupSize(size gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonProfile.Bind_set_group_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGroupSize() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonProfile.Bind_get_group_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the name of the group at [param group_idx] that will be the drawing group in the [BoneMap] editor.
*/
//go:nosplit
func (self class) GetGroupName(group_idx gd.Int) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, group_idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonProfile.Bind_get_group_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the name of the group at [param group_idx] that will be the drawing group in the [BoneMap] editor.
*/
//go:nosplit
func (self class) SetGroupName(group_idx gd.Int, group_name gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, group_idx)
	callframe.Arg(frame, discreet.Get(group_name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonProfile.Bind_set_group_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the texture of the group at [param group_idx] that will be the drawing group background image in the [BoneMap] editor.
*/
//go:nosplit
func (self class) GetTexture(group_idx gd.Int) gdclass.Texture2D {
	var frame = callframe.New()
	callframe.Arg(frame, group_idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonProfile.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Sets the texture of the group at [param group_idx] that will be the drawing group background image in the [BoneMap] editor.
*/
//go:nosplit
func (self class) SetTexture(group_idx gd.Int, texture gdclass.Texture2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, group_idx)
	callframe.Arg(frame, discreet.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonProfile.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetBoneSize(size gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonProfile.Bind_set_bone_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBoneSize() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonProfile.Bind_get_bone_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the bone index that matches [param bone_name] as its name.
*/
//go:nosplit
func (self class) FindBone(bone_name gd.StringName) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(bone_name))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonProfile.Bind_find_bone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the name of the bone at [param bone_idx] that will be the key name in the [BoneMap].
In the retargeting process, the returned bone name is the bone name of the target skeleton.
*/
//go:nosplit
func (self class) GetBoneName(bone_idx gd.Int) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonProfile.Bind_get_bone_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the name of the bone at [param bone_idx] that will be the key name in the [BoneMap].
In the retargeting process, the setting bone name is the bone name of the target skeleton.
*/
//go:nosplit
func (self class) SetBoneName(bone_idx gd.Int, bone_name gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, discreet.Get(bone_name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonProfile.Bind_set_bone_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the name of the bone which is the parent to the bone at [param bone_idx]. The result is empty if the bone has no parent.
*/
//go:nosplit
func (self class) GetBoneParent(bone_idx gd.Int) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonProfile.Bind_get_bone_parent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the bone with name [param bone_parent] as the parent of the bone at [param bone_idx]. If an empty string is passed, then the bone has no parent.
*/
//go:nosplit
func (self class) SetBoneParent(bone_idx gd.Int, bone_parent gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, discreet.Get(bone_parent))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonProfile.Bind_set_bone_parent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the tail direction of the bone at [param bone_idx].
*/
//go:nosplit
func (self class) GetTailDirection(bone_idx gd.Int) classdb.SkeletonProfileTailDirection {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[classdb.SkeletonProfileTailDirection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonProfile.Bind_get_tail_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the tail direction of the bone at [param bone_idx].
[b]Note:[/b] This only specifies the method of calculation. The actual coordinates required should be stored in an external skeleton, so the calculation itself needs to be done externally.
*/
//go:nosplit
func (self class) SetTailDirection(bone_idx gd.Int, tail_direction classdb.SkeletonProfileTailDirection)  {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, tail_direction)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonProfile.Bind_set_tail_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the name of the bone which is the tail of the bone at [param bone_idx].
*/
//go:nosplit
func (self class) GetBoneTail(bone_idx gd.Int) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonProfile.Bind_get_bone_tail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the bone with name [param bone_tail] as the tail of the bone at [param bone_idx].
*/
//go:nosplit
func (self class) SetBoneTail(bone_idx gd.Int, bone_tail gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, discreet.Get(bone_tail))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonProfile.Bind_set_bone_tail, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the reference pose transform for bone [param bone_idx].
*/
//go:nosplit
func (self class) GetReferencePose(bone_idx gd.Int) gd.Transform3D {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonProfile.Bind_get_reference_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the reference pose transform for bone [param bone_idx].
*/
//go:nosplit
func (self class) SetReferencePose(bone_idx gd.Int, bone_name gd.Transform3D)  {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, bone_name)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonProfile.Bind_set_reference_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the offset of the bone at [param bone_idx] that will be the button position in the [BoneMap] editor.
This is the offset with origin at the top left corner of the square.
*/
//go:nosplit
func (self class) GetHandleOffset(bone_idx gd.Int) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonProfile.Bind_get_handle_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the offset of the bone at [param bone_idx] that will be the button position in the [BoneMap] editor.
This is the offset with origin at the top left corner of the square.
*/
//go:nosplit
func (self class) SetHandleOffset(bone_idx gd.Int, handle_offset gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, handle_offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonProfile.Bind_set_handle_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the group of the bone at [param bone_idx].
*/
//go:nosplit
func (self class) GetGroup(bone_idx gd.Int) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonProfile.Bind_get_group, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the group of the bone at [param bone_idx].
*/
//go:nosplit
func (self class) SetGroup(bone_idx gd.Int, group gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, discreet.Get(group))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonProfile.Bind_set_group, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether the bone at [param bone_idx] is required for retargeting.
This value is used by the bone map editor. If this method returns [code]true[/code], and no bone is assigned, the handle color will be red on the bone map editor.
*/
//go:nosplit
func (self class) IsRequired(bone_idx gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonProfile.Bind_is_required, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the required status for bone [param bone_idx] to [param required].
*/
//go:nosplit
func (self class) SetRequired(bone_idx gd.Int, required bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, bone_idx)
	callframe.Arg(frame, required)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonProfile.Bind_set_required, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Go) OnProfileUpdated(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("profile_updated"), gd.NewCallable(cb), 0)
}


func (self class) AsSkeletonProfile() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSkeletonProfile() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("SkeletonProfile", func(ptr gd.Object) any { return classdb.SkeletonProfile(ptr) })}
type TailDirection = classdb.SkeletonProfileTailDirection

const (
/*Direction to the average coordinates of bone children.*/
	TailDirectionAverageChildren TailDirection = 0
/*Direction to the coordinates of specified bone child.*/
	TailDirectionSpecificChild TailDirection = 1
/*Direction is not calculated.*/
	TailDirectionEnd TailDirection = 2
)
