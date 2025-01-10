// Package GLTFSkeleton provides methods for working with GLTFSkeleton object instances.
package GLTFSkeleton

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Dictionary"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

type Instance [1]gdclass.GLTFSkeleton

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGLTFSkeleton() Instance
}

func (self Instance) GetGodotSkeleton() [1]gdclass.Skeleton3D {
	return [1]gdclass.Skeleton3D(class(self).GetGodotSkeleton())
}
func (self Instance) GetBoneAttachmentCount() int {
	return int(int(class(self).GetBoneAttachmentCount()))
}
func (self Instance) GetBoneAttachment(idx int) [1]gdclass.BoneAttachment3D {
	return [1]gdclass.BoneAttachment3D(class(self).GetBoneAttachment(gd.Int(idx)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GLTFSkeleton

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFSkeleton"))
	casted := Instance{*(*gdclass.GLTFSkeleton)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Joints() []int32 {
	return []int32(class(self).GetJoints().AsSlice())
}

func (self Instance) SetJoints(value []int32) {
	class(self).SetJoints(gd.NewPackedInt32Slice(value))
}

func (self Instance) Roots() []int32 {
	return []int32(class(self).GetRoots().AsSlice())
}

func (self Instance) SetRoots(value []int32) {
	class(self).SetRoots(gd.NewPackedInt32Slice(value))
}

func (self Instance) UniqueNames() gd.Array {
	return gd.Array(class(self).GetUniqueNames())
}

func (self Instance) SetUniqueNames(value gd.Array) {
	class(self).SetUniqueNames(value)
}

func (self Instance) GodotBoneNode() Dictionary.Any {
	return Dictionary.Any(class(self).GetGodotBoneNode())
}

func (self Instance) SetGodotBoneNode(value Dictionary.Any) {
	class(self).SetGodotBoneNode(value)
}

//go:nosplit
func (self class) GetJoints() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_get_joints, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetJoints(joints gd.PackedInt32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(joints))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_set_joints, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRoots() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_get_roots, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRoots(roots gd.PackedInt32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(roots))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_set_roots, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGodotSkeleton() [1]gdclass.Skeleton3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_get_godot_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Skeleton3D{gd.PointerWithOwnershipTransferredToGo[gdclass.Skeleton3D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetUniqueNames() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_get_unique_names, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUniqueNames(unique_names gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(unique_names))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_set_unique_names, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a [Dictionary] that maps skeleton bone indices to the indices of GLTF nodes. This property is unused during import, and only set during export. In a GLTF file, a bone is a node, so Godot converts skeleton bones to GLTF nodes.
*/
//go:nosplit
func (self class) GetGodotBoneNode() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_get_godot_bone_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets a [Dictionary] that maps skeleton bone indices to the indices of GLTF nodes. This property is unused during import, and only set during export. In a GLTF file, a bone is a node, so Godot converts skeleton bones to GLTF nodes.
*/
//go:nosplit
func (self class) SetGodotBoneNode(godot_bone_node gd.Dictionary) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(godot_bone_node))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_set_godot_bone_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBoneAttachmentCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_get_bone_attachment_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetBoneAttachment(idx gd.Int) [1]gdclass.BoneAttachment3D {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSkeleton.Bind_get_bone_attachment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.BoneAttachment3D{gd.PointerWithOwnershipTransferredToGo[gdclass.BoneAttachment3D](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsGLTFSkeleton() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGLTFSkeleton() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("GLTFSkeleton", func(ptr gd.Object) any {
		return [1]gdclass.GLTFSkeleton{*(*gdclass.GLTFSkeleton)(unsafe.Pointer(&ptr))}
	})
}
