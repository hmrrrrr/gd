// Package VisualShaderNodeSwitch provides methods for working with VisualShaderNodeSwitch object instances.
package VisualShaderNodeSwitch

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
Returns an associated value of the [member op_type] type if the provided boolean value is [code]true[/code] or [code]false[/code].
*/
type Instance [1]gdclass.VisualShaderNodeSwitch
type Any interface {
	gd.IsClass
	AsVisualShaderNodeSwitch() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeSwitch

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeSwitch"))
	return Instance{*(*gdclass.VisualShaderNodeSwitch)(unsafe.Pointer(&object))}
}

func (self Instance) OpType() gdclass.VisualShaderNodeSwitchOpType {
	return gdclass.VisualShaderNodeSwitchOpType(class(self).GetOpType())
}

func (self Instance) SetOpType(value gdclass.VisualShaderNodeSwitchOpType) {
	class(self).SetOpType(value)
}

//go:nosplit
func (self class) SetOpType(atype gdclass.VisualShaderNodeSwitchOpType) {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeSwitch.Bind_set_op_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetOpType() gdclass.VisualShaderNodeSwitchOpType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisualShaderNodeSwitchOpType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeSwitch.Bind_get_op_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeSwitch() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVisualShaderNodeSwitch() Instance {
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
	gdclass.Register("VisualShaderNodeSwitch", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeSwitch{*(*gdclass.VisualShaderNodeSwitch)(unsafe.Pointer(&ptr))}
	})
}

type OpType = gdclass.VisualShaderNodeSwitchOpType

const (
	/*A floating-point scalar.*/
	OpTypeFloat OpType = 0
	/*An integer scalar.*/
	OpTypeInt OpType = 1
	/*An unsigned integer scalar.*/
	OpTypeUint OpType = 2
	/*A 2D vector type.*/
	OpTypeVector2d OpType = 3
	/*A 3D vector type.*/
	OpTypeVector3d OpType = 4
	/*A 4D vector type.*/
	OpTypeVector4d OpType = 5
	/*A boolean type.*/
	OpTypeBoolean OpType = 6
	/*A transform type.*/
	OpTypeTransform OpType = 7
	/*Represents the size of the [enum OpType] enum.*/
	OpTypeMax OpType = 8
)
