// Package Bone2D provides methods for working with Bone2D object instances.
package Bone2D

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
import "graphics.gd/classdb/Node2D"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Transform2D"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
A hierarchy of [Bone2D]s can be bound to a [Skeleton2D] to control and animate other [Node2D] nodes.
You can use [Bone2D] and [Skeleton2D] nodes to animate 2D meshes created with the [Polygon2D] UV editor.
Each bone has a [member rest] transform that you can reset to with [method apply_rest]. These rest poses are relative to the bone's parent.
If in the editor, you can set the rest pose of an entire skeleton using a menu option, from the code, you need to iterate over the bones to set their individual rest poses.
*/
type Instance [1]gdclass.Bone2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsBone2D() Instance
}

/*
Resets the bone to the rest pose. This is equivalent to setting [member Node2D.transform] to [member rest].
*/
func (self Instance) ApplyRest() { //gd:Bone2D.apply_rest
	class(self).ApplyRest()
}

/*
Returns the node's [member rest] [Transform2D] if it doesn't have a parent, or its rest pose relative to its parent.
*/
func (self Instance) GetSkeletonRest() Transform2D.OriginXY { //gd:Bone2D.get_skeleton_rest
	return Transform2D.OriginXY(class(self).GetSkeletonRest())
}

/*
Returns the node's index as part of the entire skeleton. See [Skeleton2D].
*/
func (self Instance) GetIndexInSkeleton() int { //gd:Bone2D.get_index_in_skeleton
	return int(int(class(self).GetIndexInSkeleton()))
}

/*
When set to [code]true[/code], the [Bone2D] node will attempt to automatically calculate the bone angle and length using the first child [Bone2D] node, if one exists. If none exist, the [Bone2D] cannot automatically calculate these values and will print a warning.
*/
func (self Instance) SetAutocalculateLengthAndAngle(auto_calculate bool) { //gd:Bone2D.set_autocalculate_length_and_angle
	class(self).SetAutocalculateLengthAndAngle(auto_calculate)
}

/*
Returns whether this [Bone2D] is going to autocalculate its length and bone angle using its first [Bone2D] child node, if one exists. If there are no [Bone2D] children, then it cannot autocalculate these values and will print a warning.
*/
func (self Instance) GetAutocalculateLengthAndAngle() bool { //gd:Bone2D.get_autocalculate_length_and_angle
	return bool(class(self).GetAutocalculateLengthAndAngle())
}

/*
Sets the length of the bone in the [Bone2D].
*/
func (self Instance) SetLength(length Float.X) { //gd:Bone2D.set_length
	class(self).SetLength(gd.Float(length))
}

/*
Returns the length of the bone in the [Bone2D] node.
*/
func (self Instance) GetLength() Float.X { //gd:Bone2D.get_length
	return Float.X(Float.X(class(self).GetLength()))
}

/*
Sets the bone angle for the [Bone2D]. This is typically set to the rotation from the [Bone2D] to a child [Bone2D] node.
[b]Note:[/b] This is different from the [Bone2D]'s rotation. The bone's angle is the rotation of the bone shown by the gizmo, which is unaffected by the [Bone2D]'s [member Node2D.transform].
*/
func (self Instance) SetBoneAngle(angle Float.X) { //gd:Bone2D.set_bone_angle
	class(self).SetBoneAngle(gd.Float(angle))
}

/*
Returns the angle of the bone in the [Bone2D].
[b]Note:[/b] This is different from the [Bone2D]'s rotation. The bone's angle is the rotation of the bone shown by the gizmo, which is unaffected by the [Bone2D]'s [member Node2D.transform].
*/
func (self Instance) GetBoneAngle() Float.X { //gd:Bone2D.get_bone_angle
	return Float.X(Float.X(class(self).GetBoneAngle()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Bone2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Bone2D"))
	casted := Instance{*(*gdclass.Bone2D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Rest() Transform2D.OriginXY {
	return Transform2D.OriginXY(class(self).GetRest())
}

func (self Instance) SetRest(value Transform2D.OriginXY) {
	class(self).SetRest(gd.Transform2D(value))
}

//go:nosplit
func (self class) SetRest(rest gd.Transform2D) { //gd:Bone2D.set_rest
	var frame = callframe.New()
	callframe.Arg(frame, rest)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Bone2D.Bind_set_rest, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRest() gd.Transform2D { //gd:Bone2D.get_rest
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Bone2D.Bind_get_rest, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Resets the bone to the rest pose. This is equivalent to setting [member Node2D.transform] to [member rest].
*/
//go:nosplit
func (self class) ApplyRest() { //gd:Bone2D.apply_rest
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Bone2D.Bind_apply_rest, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the node's [member rest] [Transform2D] if it doesn't have a parent, or its rest pose relative to its parent.
*/
//go:nosplit
func (self class) GetSkeletonRest() gd.Transform2D { //gd:Bone2D.get_skeleton_rest
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Bone2D.Bind_get_skeleton_rest, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the node's index as part of the entire skeleton. See [Skeleton2D].
*/
//go:nosplit
func (self class) GetIndexInSkeleton() gd.Int { //gd:Bone2D.get_index_in_skeleton
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Bone2D.Bind_get_index_in_skeleton, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
When set to [code]true[/code], the [Bone2D] node will attempt to automatically calculate the bone angle and length using the first child [Bone2D] node, if one exists. If none exist, the [Bone2D] cannot automatically calculate these values and will print a warning.
*/
//go:nosplit
func (self class) SetAutocalculateLengthAndAngle(auto_calculate bool) { //gd:Bone2D.set_autocalculate_length_and_angle
	var frame = callframe.New()
	callframe.Arg(frame, auto_calculate)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Bone2D.Bind_set_autocalculate_length_and_angle, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether this [Bone2D] is going to autocalculate its length and bone angle using its first [Bone2D] child node, if one exists. If there are no [Bone2D] children, then it cannot autocalculate these values and will print a warning.
*/
//go:nosplit
func (self class) GetAutocalculateLengthAndAngle() bool { //gd:Bone2D.get_autocalculate_length_and_angle
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Bone2D.Bind_get_autocalculate_length_and_angle, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the length of the bone in the [Bone2D].
*/
//go:nosplit
func (self class) SetLength(length gd.Float) { //gd:Bone2D.set_length
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Bone2D.Bind_set_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the length of the bone in the [Bone2D] node.
*/
//go:nosplit
func (self class) GetLength() gd.Float { //gd:Bone2D.get_length
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Bone2D.Bind_get_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the bone angle for the [Bone2D]. This is typically set to the rotation from the [Bone2D] to a child [Bone2D] node.
[b]Note:[/b] This is different from the [Bone2D]'s rotation. The bone's angle is the rotation of the bone shown by the gizmo, which is unaffected by the [Bone2D]'s [member Node2D.transform].
*/
//go:nosplit
func (self class) SetBoneAngle(angle gd.Float) { //gd:Bone2D.set_bone_angle
	var frame = callframe.New()
	callframe.Arg(frame, angle)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Bone2D.Bind_set_bone_angle, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the angle of the bone in the [Bone2D].
[b]Note:[/b] This is different from the [Bone2D]'s rotation. The bone's angle is the rotation of the bone shown by the gizmo, which is unaffected by the [Bone2D]'s [member Node2D.transform].
*/
//go:nosplit
func (self class) GetBoneAngle() gd.Float { //gd:Bone2D.get_bone_angle
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Bone2D.Bind_get_bone_angle, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsBone2D() Advanced           { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsBone2D() Instance        { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("Bone2D", func(ptr gd.Object) any { return [1]gdclass.Bone2D{*(*gdclass.Bone2D)(unsafe.Pointer(&ptr))} })
}
