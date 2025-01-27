// Package VisualShaderNodeTransformOp provides methods for working with VisualShaderNodeTransformOp object instances.
package VisualShaderNodeTransformOp

import "unsafe"
import "reflect"
import "slices"
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
import "graphics.gd/variant/Path"
import "graphics.gd/variant/Packed"
import "graphics.gd/classdb/VisualShaderNode"
import "graphics.gd/classdb/Resource"

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
var _ = slices.Delete[[]struct{}, struct{}]

/*
Applies [member operator] to two transform (4×4 matrices) inputs.
*/
type Instance [1]gdclass.VisualShaderNodeTransformOp

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVisualShaderNodeTransformOp() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeTransformOp

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeTransformOp"))
	casted := Instance{*(*gdclass.VisualShaderNodeTransformOp)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Operator() gdclass.VisualShaderNodeTransformOpOperator {
	return gdclass.VisualShaderNodeTransformOpOperator(class(self).GetOperator())
}

func (self Instance) SetOperator(value gdclass.VisualShaderNodeTransformOpOperator) {
	class(self).SetOperator(value)
}

//go:nosplit
func (self class) SetOperator(op gdclass.VisualShaderNodeTransformOpOperator) { //gd:VisualShaderNodeTransformOp.set_operator
	var frame = callframe.New()
	callframe.Arg(frame, op)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTransformOp.Bind_set_operator, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOperator() gdclass.VisualShaderNodeTransformOpOperator { //gd:VisualShaderNodeTransformOp.get_operator
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeTransformOpOperator](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTransformOp.Bind_get_operator, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeTransformOp() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeTransformOp() Instance {
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualShaderNode.Advanced(self.AsVisualShaderNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualShaderNode.Instance(self.AsVisualShaderNode()), name)
	}
}
func init() {
	gdclass.Register("VisualShaderNodeTransformOp", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeTransformOp{*(*gdclass.VisualShaderNodeTransformOp)(unsafe.Pointer(&ptr))}
	})
}

type Operator = gdclass.VisualShaderNodeTransformOpOperator //gd:VisualShaderNodeTransformOp.Operator

const (
	/*Multiplies transform [code]a[/code] by the transform [code]b[/code].*/
	OpAxb Operator = 0
	/*Multiplies transform [code]b[/code] by the transform [code]a[/code].*/
	OpBxa Operator = 1
	/*Performs a component-wise multiplication of transform [code]a[/code] by the transform [code]b[/code].*/
	OpAxbComp Operator = 2
	/*Performs a component-wise multiplication of transform [code]b[/code] by the transform [code]a[/code].*/
	OpBxaComp Operator = 3
	/*Adds two transforms.*/
	OpAdd Operator = 4
	/*Subtracts the transform [code]a[/code] from the transform [code]b[/code].*/
	OpAMinusB Operator = 5
	/*Subtracts the transform [code]b[/code] from the transform [code]a[/code].*/
	OpBMinusA Operator = 6
	/*Divides the transform [code]a[/code] by the transform [code]b[/code].*/
	OpADivB Operator = 7
	/*Divides the transform [code]b[/code] by the transform [code]a[/code].*/
	OpBDivA Operator = 8
	/*Represents the size of the [enum Operator] enum.*/
	OpMax Operator = 9
)
