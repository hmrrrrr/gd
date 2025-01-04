// Package VisualShaderNodeVectorOp provides methods for working with VisualShaderNodeVectorOp object instances.
package VisualShaderNodeVectorOp

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/VisualShaderNodeVectorBase"
import "graphics.gd/classdb/VisualShaderNode"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A visual shader node for use of vector operators. Operates on vector [code]a[/code] and vector [code]b[/code].
*/
type Instance [1]gdclass.VisualShaderNodeVectorOp
type Any interface {
	gd.IsClass
	AsVisualShaderNodeVectorOp() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeVectorOp

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeVectorOp"))
	return Instance{*(*gdclass.VisualShaderNodeVectorOp)(unsafe.Pointer(&object))}
}

func (self Instance) Operator() gdclass.VisualShaderNodeVectorOpOperator {
	return gdclass.VisualShaderNodeVectorOpOperator(class(self).GetOperator())
}

func (self Instance) SetOperator(value gdclass.VisualShaderNodeVectorOpOperator) {
	class(self).SetOperator(value)
}

//go:nosplit
func (self class) SetOperator(op gdclass.VisualShaderNodeVectorOpOperator) {
	var frame = callframe.New()
	callframe.Arg(frame, op)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeVectorOp.Bind_set_operator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetOperator() gdclass.VisualShaderNodeVectorOpOperator {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeVectorOpOperator](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeVectorOp.Bind_get_operator, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeVectorOp() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVisualShaderNodeVectorOp() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualShaderNodeVectorBase() VisualShaderNodeVectorBase.Advanced {
	return *((*VisualShaderNodeVectorBase.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeVectorBase() VisualShaderNodeVectorBase.Instance {
	return *((*VisualShaderNodeVectorBase.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsVisualShaderNodeVectorBase(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualShaderNodeVectorBase(), name)
	}
}
func init() {
	gdclass.Register("VisualShaderNodeVectorOp", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeVectorOp{*(*gdclass.VisualShaderNodeVectorOp)(unsafe.Pointer(&ptr))}
	})
}

type Operator = gdclass.VisualShaderNodeVectorOpOperator

const (
	/*Adds two vectors.*/
	OpAdd Operator = 0
	/*Subtracts a vector from a vector.*/
	OpSub Operator = 1
	/*Multiplies two vectors.*/
	OpMul Operator = 2
	/*Divides vector by vector.*/
	OpDiv Operator = 3
	/*Returns the remainder of the two vectors.*/
	OpMod Operator = 4
	/*Returns the value of the first parameter raised to the power of the second, for each component of the vectors.*/
	OpPow Operator = 5
	/*Returns the greater of two values, for each component of the vectors.*/
	OpMax Operator = 6
	/*Returns the lesser of two values, for each component of the vectors.*/
	OpMin Operator = 7
	/*Calculates the cross product of two vectors.*/
	OpCross Operator = 8
	/*Returns the arc-tangent of the parameters.*/
	OpAtan2 Operator = 9
	/*Returns the vector that points in the direction of reflection. [code]a[/code] is incident vector and [code]b[/code] is the normal vector.*/
	OpReflect Operator = 10
	/*Vector step operator. Returns [code]0.0[/code] if [code]a[/code] is smaller than [code]b[/code] and [code]1.0[/code] otherwise.*/
	OpStep Operator = 11
	/*Represents the size of the [enum Operator] enum.*/
	OpEnumSize Operator = 12
)
