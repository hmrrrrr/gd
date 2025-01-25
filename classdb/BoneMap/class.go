// Package BoneMap provides methods for working with BoneMap object instances.
package BoneMap

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
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any

/*
This class contains a dictionary that uses a list of bone names in [SkeletonProfile] as key names.
By assigning the actual [Skeleton3D] bone name as the key value, it maps the [Skeleton3D] to the [SkeletonProfile].
*/
type Instance [1]gdclass.BoneMap

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsBoneMap() Instance
}

/*
Returns a skeleton bone name is mapped to [param profile_bone_name].
In the retargeting process, the returned bone name is the bone name of the source skeleton.
*/
func (self Instance) GetSkeletonBoneName(profile_bone_name string) string {
	return string(class(self).GetSkeletonBoneName(gd.NewStringName(profile_bone_name)).String())
}

/*
Maps a skeleton bone name to [param profile_bone_name].
In the retargeting process, the setting bone name is the bone name of the source skeleton.
*/
func (self Instance) SetSkeletonBoneName(profile_bone_name string, skeleton_bone_name string) {
	class(self).SetSkeletonBoneName(gd.NewStringName(profile_bone_name), gd.NewStringName(skeleton_bone_name))
}

/*
Returns a profile bone name having [param skeleton_bone_name]. If not found, an empty [StringName] will be returned.
In the retargeting process, the returned bone name is the bone name of the target skeleton.
*/
func (self Instance) FindProfileBoneName(skeleton_bone_name string) string {
	return string(class(self).FindProfileBoneName(gd.NewStringName(skeleton_bone_name)).String())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.BoneMap

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("BoneMap"))
	casted := Instance{*(*gdclass.BoneMap)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Profile() [1]gdclass.SkeletonProfile {
	return [1]gdclass.SkeletonProfile(class(self).GetProfile())
}

func (self Instance) SetProfile(value [1]gdclass.SkeletonProfile) {
	class(self).SetProfile(value)
}

//go:nosplit
func (self class) GetProfile() [1]gdclass.SkeletonProfile {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoneMap.Bind_get_profile, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.SkeletonProfile{gd.PointerWithOwnershipTransferredToGo[gdclass.SkeletonProfile](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetProfile(profile [1]gdclass.SkeletonProfile) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(profile[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoneMap.Bind_set_profile, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a skeleton bone name is mapped to [param profile_bone_name].
In the retargeting process, the returned bone name is the bone name of the source skeleton.
*/
//go:nosplit
func (self class) GetSkeletonBoneName(profile_bone_name gd.StringName) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(profile_bone_name))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoneMap.Bind_get_skeleton_bone_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

/*
Maps a skeleton bone name to [param profile_bone_name].
In the retargeting process, the setting bone name is the bone name of the source skeleton.
*/
//go:nosplit
func (self class) SetSkeletonBoneName(profile_bone_name gd.StringName, skeleton_bone_name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(profile_bone_name))
	callframe.Arg(frame, pointers.Get(skeleton_bone_name))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoneMap.Bind_set_skeleton_bone_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a profile bone name having [param skeleton_bone_name]. If not found, an empty [StringName] will be returned.
In the retargeting process, the returned bone name is the bone name of the target skeleton.
*/
//go:nosplit
func (self class) FindProfileBoneName(skeleton_bone_name gd.StringName) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(skeleton_bone_name))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoneMap.Bind_find_profile_bone_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}
func (self Instance) OnBoneMapUpdated(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("bone_map_updated"), gd.NewCallable(cb), 0)
}

func (self Instance) OnProfileUpdated(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("profile_updated"), gd.NewCallable(cb), 0)
}

func (self class) AsBoneMap() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsBoneMap() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("BoneMap", func(ptr gd.Object) any { return [1]gdclass.BoneMap{*(*gdclass.BoneMap)(unsafe.Pointer(&ptr))} })
}
