// Package VisualShaderNodeFloatOp provides methods for working with VisualShaderNodeFloatOp object instances.
package VisualShaderNodeFloatOp

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/VisualShaderNode"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Applies [member operator] to two floating-point inputs: [code]a[/code] and [code]b[/code].
*/
type Instance [1]gdclass.VisualShaderNodeFloatOp
type Any interface {
	gd.IsClass
	AsVisualShaderNodeFloatOp() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeFloatOp

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeFloatOp"))
	return Instance{*(*gdclass.VisualShaderNodeFloatOp)(unsafe.Pointer(&object))}
}

func (self Instance) Operator() gdclass.VisualShaderNodeFloatOpOperator {
	return gdclass.VisualShaderNodeFloatOpOperator(class(self).GetOperator())
}

func (self Instance) SetOperator(value gdclass.VisualShaderNodeFloatOpOperator) {
	class(self).SetOperator(value)
}

//go:nosplit
func (self class) SetOperator(op gdclass.VisualShaderNodeFloatOpOperator) {
	var frame = callframe.New()
	callframe.Arg(frame, op)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeFloatOp.Bind_set_operator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetOperator() gdclass.VisualShaderNodeFloatOpOperator {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeFloatOpOperator](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeFloatOp.Bind_get_operator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeFloatOp() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVisualShaderNodeFloatOp() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualShaderNode() VisualShaderNode.Advanced {
	return *((*VisualShaderNode.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNode() VisualShaderNode.Instance {
	return *((*VisualShaderNode.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualShaderNode(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualShaderNode(), name)
	}
}
func init() {
	gdclass.Register("VisualShaderNodeFloatOp", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeFloatOp{*(*gdclass.VisualShaderNodeFloatOp)(unsafe.Pointer(&ptr))}
	})
}

type Operator = gdclass.VisualShaderNodeFloatOpOperator

const (
	/*Sums two numbers using [code]a + b[/code].*/
	OpAdd Operator = 0
	/*Subtracts two numbers using [code]a - b[/code].*/
	OpSub Operator = 1
	/*Multiplies two numbers using [code]a * b[/code].*/
	OpMul Operator = 2
	/*Divides two numbers using [code]a / b[/code].*/
	OpDiv Operator = 3
	/*Calculates the remainder of two numbers. Translates to [code]mod(a, b)[/code] in the Godot Shader Language.*/
	OpMod Operator = 4
	/*Raises the [code]a[/code] to the power of [code]b[/code]. Translates to [code]pow(a, b)[/code] in the Godot Shader Language.*/
	OpPow Operator = 5
	/*Returns the greater of two numbers. Translates to [code]max(a, b)[/code] in the Godot Shader Language.*/
	OpMax Operator = 6
	/*Returns the lesser of two numbers. Translates to [code]min(a, b)[/code] in the Godot Shader Language.*/
	OpMin Operator = 7
	/*Returns the arc-tangent of the parameters. Translates to [code]atan(a, b)[/code] in the Godot Shader Language.*/
	OpAtan2 Operator = 8
	/*Generates a step function by comparing [code]b[/code](x) to [code]a[/code](edge). Returns 0.0 if [code]x[/code] is smaller than [code]edge[/code] and otherwise 1.0. Translates to [code]step(a, b)[/code] in the Godot Shader Language.*/
	OpStep Operator = 9
	/*Represents the size of the [enum Operator] enum.*/
	OpEnumSize Operator = 10
)
