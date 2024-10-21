package VisualShaderNodeStep

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/VisualShaderNode"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Translates to [code]step(edge, x)[/code] in the shader language.
Returns [code]0.0[/code] if [code]x[/code] is smaller than [code]edge[/code] and [code]1.0[/code] otherwise.

*/
type Simple [1]classdb.VisualShaderNodeStep
func (self Simple) SetOpType(op_type classdb.VisualShaderNodeStepOpType) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOpType(op_type)
}
func (self Simple) GetOpType() classdb.VisualShaderNodeStepOpType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.VisualShaderNodeStepOpType(Expert(self).GetOpType())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VisualShaderNodeStep
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetOpType(op_type classdb.VisualShaderNodeStepOpType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, op_type)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeStep.Bind_set_op_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOpType() classdb.VisualShaderNodeStepOpType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeStepOpType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeStep.Bind_get_op_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsVisualShaderNodeStep() Expert { return self[0].AsVisualShaderNodeStep() }


//go:nosplit
func (self Simple) AsVisualShaderNodeStep() Simple { return self[0].AsVisualShaderNodeStep() }


//go:nosplit
func (self class) AsVisualShaderNode() VisualShaderNode.Expert { return self[0].AsVisualShaderNode() }


//go:nosplit
func (self Simple) AsVisualShaderNode() VisualShaderNode.Simple { return self[0].AsVisualShaderNode() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("VisualShaderNodeStep", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type OpType = classdb.VisualShaderNodeStepOpType

const (
/*A floating-point scalar type.*/
	OpTypeScalar OpType = 0
/*A 2D vector type.*/
	OpTypeVector2d OpType = 1
/*The [code]x[/code] port uses a 2D vector type, while the [code]edge[/code] port uses a floating-point scalar type.*/
	OpTypeVector2dScalar OpType = 2
/*A 3D vector type.*/
	OpTypeVector3d OpType = 3
/*The [code]x[/code] port uses a 3D vector type, while the [code]edge[/code] port uses a floating-point scalar type.*/
	OpTypeVector3dScalar OpType = 4
/*A 4D vector type.*/
	OpTypeVector4d OpType = 5
/*The [code]a[/code] and [code]b[/code] ports use a 4D vector type. The [code]weight[/code] port uses a scalar type.*/
	OpTypeVector4dScalar OpType = 6
/*Represents the size of the [enum OpType] enum.*/
	OpTypeMax OpType = 7
)
