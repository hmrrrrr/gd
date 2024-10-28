package CSGShape3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/GeometryInstance3D"
import "grow.graphics/gd/gdclass/VisualInstance3D"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
This is the CSG base class that provides CSG operation support to the various CSG nodes in Godot.
[b]Note:[/b] CSG nodes are intended to be used for level prototyping. Creating CSG nodes has a significant CPU cost compared to creating a [MeshInstance3D] with a [PrimitiveMesh]. Moving a CSG node within another CSG node also has a significant CPU cost, so it should be avoided during gameplay.

*/
type Go [1]classdb.CSGShape3D

/*
Returns [code]true[/code] if this is a root shape and is thus the object that is rendered.
*/
func (self Go) IsRootShape() bool {
	return bool(class(self).IsRootShape())
}

/*
Based on [param value], enables or disables the specified layer in the [member collision_mask], given a [param layer_number] between 1 and 32.
*/
func (self Go) SetCollisionMaskValue(layer_number int, value bool) {
	class(self).SetCollisionMaskValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Go) GetCollisionMaskValue(layer_number int) bool {
	return bool(class(self).GetCollisionMaskValue(gd.Int(layer_number)))
}

/*
Based on [param value], enables or disables the specified layer in the [member collision_layer], given a [param layer_number] between 1 and 32.
*/
func (self Go) SetCollisionLayerValue(layer_number int, value bool) {
	class(self).SetCollisionLayerValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member collision_layer] is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Go) GetCollisionLayerValue(layer_number int) bool {
	return bool(class(self).GetCollisionLayerValue(gd.Int(layer_number)))
}

/*
Returns an [Array] with two elements, the first is the [Transform3D] of this node and the second is the root [Mesh] of this node. Only works when this node is the root shape.
*/
func (self Go) GetMeshes() gd.Array {
	return gd.Array(class(self).GetMeshes())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.CSGShape3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CSGShape3D"))
	return Go{classdb.CSGShape3D(object)}
}

func (self Go) Operation() classdb.CSGShape3DOperation {
		return classdb.CSGShape3DOperation(class(self).GetOperation())
}

func (self Go) SetOperation(value classdb.CSGShape3DOperation) {
	class(self).SetOperation(value)
}

func (self Go) Snap() float64 {
		return float64(float64(class(self).GetSnap()))
}

func (self Go) SetSnap(value float64) {
	class(self).SetSnap(gd.Float(value))
}

func (self Go) CalculateTangents() bool {
		return bool(class(self).IsCalculatingTangents())
}

func (self Go) SetCalculateTangents(value bool) {
	class(self).SetCalculateTangents(value)
}

func (self Go) UseCollision() bool {
		return bool(class(self).IsUsingCollision())
}

func (self Go) SetUseCollision(value bool) {
	class(self).SetUseCollision(value)
}

func (self Go) CollisionLayer() int {
		return int(int(class(self).GetCollisionLayer()))
}

func (self Go) SetCollisionLayer(value int) {
	class(self).SetCollisionLayer(gd.Int(value))
}

func (self Go) CollisionMask() int {
		return int(int(class(self).GetCollisionMask()))
}

func (self Go) SetCollisionMask(value int) {
	class(self).SetCollisionMask(gd.Int(value))
}

func (self Go) CollisionPriority() float64 {
		return float64(float64(class(self).GetCollisionPriority()))
}

func (self Go) SetCollisionPriority(value float64) {
	class(self).SetCollisionPriority(gd.Float(value))
}

/*
Returns [code]true[/code] if this is a root shape and is thus the object that is rendered.
*/
//go:nosplit
func (self class) IsRootShape() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_is_root_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOperation(operation classdb.CSGShape3DOperation)  {
	var frame = callframe.New()
	callframe.Arg(frame, operation)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_set_operation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOperation() classdb.CSGShape3DOperation {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CSGShape3DOperation](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_get_operation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSnap(snap gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, snap)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_set_snap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSnap() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_get_snap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseCollision(operation bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, operation)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_set_use_collision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingCollision() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_is_using_collision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollisionLayer(layer gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_set_collision_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCollisionLayer() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_get_collision_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollisionMask(mask gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCollisionMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Based on [param value], enables or disables the specified layer in the [member collision_mask], given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetCollisionMaskValue(layer_number gd.Int, value bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_set_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether or not the specified layer of the [member collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetCollisionMaskValue(layer_number gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_get_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Based on [param value], enables or disables the specified layer in the [member collision_layer], given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetCollisionLayerValue(layer_number gd.Int, value bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_set_collision_layer_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether or not the specified layer of the [member collision_layer] is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetCollisionLayerValue(layer_number gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_get_collision_layer_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollisionPriority(priority gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_set_collision_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCollisionPriority() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_get_collision_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCalculateTangents(enabled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_set_calculate_tangents, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCalculatingTangents() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_is_calculating_tangents, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an [Array] with two elements, the first is the [Transform3D] of this node and the second is the root [Mesh] of this node. Only works when this node is the root shape.
*/
//go:nosplit
func (self class) GetMeshes() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGShape3D.Bind_get_meshes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsCSGShape3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsCSGShape3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsGeometryInstance3D() GeometryInstance3D.GD { return *((*GeometryInstance3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsGeometryInstance3D() GeometryInstance3D.Go { return *((*GeometryInstance3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualInstance3D() VisualInstance3D.GD { return *((*VisualInstance3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualInstance3D() VisualInstance3D.Go { return *((*VisualInstance3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsGeometryInstance3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsGeometryInstance3D(), name)
	}
}
func init() {classdb.Register("CSGShape3D", func(ptr gd.Object) any { return classdb.CSGShape3D(ptr) })}
type Operation = classdb.CSGShape3DOperation

const (
/*Geometry of both primitives is merged, intersecting geometry is removed.*/
	OperationUnion Operation = 0
/*Only intersecting geometry remains, the rest is removed.*/
	OperationIntersection Operation = 1
/*The second shape is subtracted from the first, leaving a dent with its shape.*/
	OperationSubtraction Operation = 2
)
