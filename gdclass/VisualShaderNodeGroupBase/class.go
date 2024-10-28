package VisualShaderNodeGroupBase

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VisualShaderNodeResizableBase"
import "grow.graphics/gd/gdclass/VisualShaderNode"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Currently, has no direct usage, use the derived classes instead.

*/
type Go [1]classdb.VisualShaderNodeGroupBase

/*
Defines all input ports using a [String] formatted as a colon-separated list: [code]id,type,name;[/code] (see [method add_input_port]).
*/
func (self Go) SetInputs(inputs string) {
	class(self).SetInputs(gd.NewString(inputs))
}

/*
Returns a [String] description of the input ports as a colon-separated list using the format [code]id,type,name;[/code] (see [method add_input_port]).
*/
func (self Go) GetInputs() string {
	return string(class(self).GetInputs().String())
}

/*
Defines all output ports using a [String] formatted as a colon-separated list: [code]id,type,name;[/code] (see [method add_output_port]).
*/
func (self Go) SetOutputs(outputs string) {
	class(self).SetOutputs(gd.NewString(outputs))
}

/*
Returns a [String] description of the output ports as a colon-separated list using the format [code]id,type,name;[/code] (see [method add_output_port]).
*/
func (self Go) GetOutputs() string {
	return string(class(self).GetOutputs().String())
}

/*
Returns [code]true[/code] if the specified port name does not override an existed port name and is valid within the shader.
*/
func (self Go) IsValidPortName(name string) bool {
	return bool(class(self).IsValidPortName(gd.NewString(name)))
}

/*
Adds an input port with the specified [param type] (see [enum VisualShaderNode.PortType]) and [param name].
*/
func (self Go) AddInputPort(id int, atype int, name string) {
	class(self).AddInputPort(gd.Int(id), gd.Int(atype), gd.NewString(name))
}

/*
Removes the specified input port.
*/
func (self Go) RemoveInputPort(id int) {
	class(self).RemoveInputPort(gd.Int(id))
}

/*
Returns the number of input ports in use. Alternative for [method get_free_input_port_id].
*/
func (self Go) GetInputPortCount() int {
	return int(int(class(self).GetInputPortCount()))
}

/*
Returns [code]true[/code] if the specified input port exists.
*/
func (self Go) HasInputPort(id int) bool {
	return bool(class(self).HasInputPort(gd.Int(id)))
}

/*
Removes all previously specified input ports.
*/
func (self Go) ClearInputPorts() {
	class(self).ClearInputPorts()
}

/*
Adds an output port with the specified [param type] (see [enum VisualShaderNode.PortType]) and [param name].
*/
func (self Go) AddOutputPort(id int, atype int, name string) {
	class(self).AddOutputPort(gd.Int(id), gd.Int(atype), gd.NewString(name))
}

/*
Removes the specified output port.
*/
func (self Go) RemoveOutputPort(id int) {
	class(self).RemoveOutputPort(gd.Int(id))
}

/*
Returns the number of output ports in use. Alternative for [method get_free_output_port_id].
*/
func (self Go) GetOutputPortCount() int {
	return int(int(class(self).GetOutputPortCount()))
}

/*
Returns [code]true[/code] if the specified output port exists.
*/
func (self Go) HasOutputPort(id int) bool {
	return bool(class(self).HasOutputPort(gd.Int(id)))
}

/*
Removes all previously specified output ports.
*/
func (self Go) ClearOutputPorts() {
	class(self).ClearOutputPorts()
}

/*
Renames the specified input port.
*/
func (self Go) SetInputPortName(id int, name string) {
	class(self).SetInputPortName(gd.Int(id), gd.NewString(name))
}

/*
Sets the specified input port's type (see [enum VisualShaderNode.PortType]).
*/
func (self Go) SetInputPortType(id int, atype int) {
	class(self).SetInputPortType(gd.Int(id), gd.Int(atype))
}

/*
Renames the specified output port.
*/
func (self Go) SetOutputPortName(id int, name string) {
	class(self).SetOutputPortName(gd.Int(id), gd.NewString(name))
}

/*
Sets the specified output port's type (see [enum VisualShaderNode.PortType]).
*/
func (self Go) SetOutputPortType(id int, atype int) {
	class(self).SetOutputPortType(gd.Int(id), gd.Int(atype))
}

/*
Returns a free input port ID which can be used in [method add_input_port].
*/
func (self Go) GetFreeInputPortId() int {
	return int(int(class(self).GetFreeInputPortId()))
}

/*
Returns a free output port ID which can be used in [method add_output_port].
*/
func (self Go) GetFreeOutputPortId() int {
	return int(int(class(self).GetFreeOutputPortId()))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VisualShaderNodeGroupBase
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeGroupBase"))
	return Go{classdb.VisualShaderNodeGroupBase(object)}
}

/*
Defines all input ports using a [String] formatted as a colon-separated list: [code]id,type,name;[/code] (see [method add_input_port]).
*/
//go:nosplit
func (self class) SetInputs(inputs gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(inputs))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeGroupBase.Bind_set_inputs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a [String] description of the input ports as a colon-separated list using the format [code]id,type,name;[/code] (see [method add_input_port]).
*/
//go:nosplit
func (self class) GetInputs() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeGroupBase.Bind_get_inputs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
/*
Defines all output ports using a [String] formatted as a colon-separated list: [code]id,type,name;[/code] (see [method add_output_port]).
*/
//go:nosplit
func (self class) SetOutputs(outputs gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(outputs))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeGroupBase.Bind_set_outputs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a [String] description of the output ports as a colon-separated list using the format [code]id,type,name;[/code] (see [method add_output_port]).
*/
//go:nosplit
func (self class) GetOutputs() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeGroupBase.Bind_get_outputs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the specified port name does not override an existed port name and is valid within the shader.
*/
//go:nosplit
func (self class) IsValidPortName(name gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeGroupBase.Bind_is_valid_port_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds an input port with the specified [param type] (see [enum VisualShaderNode.PortType]) and [param name].
*/
//go:nosplit
func (self class) AddInputPort(id gd.Int, atype gd.Int, name gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, atype)
	callframe.Arg(frame, discreet.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeGroupBase.Bind_add_input_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the specified input port.
*/
//go:nosplit
func (self class) RemoveInputPort(id gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeGroupBase.Bind_remove_input_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of input ports in use. Alternative for [method get_free_input_port_id].
*/
//go:nosplit
func (self class) GetInputPortCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeGroupBase.Bind_get_input_port_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the specified input port exists.
*/
//go:nosplit
func (self class) HasInputPort(id gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeGroupBase.Bind_has_input_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes all previously specified input ports.
*/
//go:nosplit
func (self class) ClearInputPorts()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeGroupBase.Bind_clear_input_ports, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds an output port with the specified [param type] (see [enum VisualShaderNode.PortType]) and [param name].
*/
//go:nosplit
func (self class) AddOutputPort(id gd.Int, atype gd.Int, name gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, atype)
	callframe.Arg(frame, discreet.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeGroupBase.Bind_add_output_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the specified output port.
*/
//go:nosplit
func (self class) RemoveOutputPort(id gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeGroupBase.Bind_remove_output_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of output ports in use. Alternative for [method get_free_output_port_id].
*/
//go:nosplit
func (self class) GetOutputPortCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeGroupBase.Bind_get_output_port_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the specified output port exists.
*/
//go:nosplit
func (self class) HasOutputPort(id gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeGroupBase.Bind_has_output_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes all previously specified output ports.
*/
//go:nosplit
func (self class) ClearOutputPorts()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeGroupBase.Bind_clear_output_ports, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Renames the specified input port.
*/
//go:nosplit
func (self class) SetInputPortName(id gd.Int, name gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, discreet.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeGroupBase.Bind_set_input_port_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the specified input port's type (see [enum VisualShaderNode.PortType]).
*/
//go:nosplit
func (self class) SetInputPortType(id gd.Int, atype gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeGroupBase.Bind_set_input_port_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Renames the specified output port.
*/
//go:nosplit
func (self class) SetOutputPortName(id gd.Int, name gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, discreet.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeGroupBase.Bind_set_output_port_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the specified output port's type (see [enum VisualShaderNode.PortType]).
*/
//go:nosplit
func (self class) SetOutputPortType(id gd.Int, atype gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeGroupBase.Bind_set_output_port_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a free input port ID which can be used in [method add_input_port].
*/
//go:nosplit
func (self class) GetFreeInputPortId() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeGroupBase.Bind_get_free_input_port_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a free output port ID which can be used in [method add_output_port].
*/
//go:nosplit
func (self class) GetFreeOutputPortId() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeGroupBase.Bind_get_free_output_port_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeGroupBase() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeGroupBase() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNodeResizableBase() VisualShaderNodeResizableBase.GD { return *((*VisualShaderNodeResizableBase.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeResizableBase() VisualShaderNodeResizableBase.Go { return *((*VisualShaderNodeResizableBase.Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNode() VisualShaderNode.GD { return *((*VisualShaderNode.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNode() VisualShaderNode.Go { return *((*VisualShaderNode.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNodeResizableBase(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNodeResizableBase(), name)
	}
}
func init() {classdb.Register("VisualShaderNodeGroupBase", func(ptr gd.Object) any { return classdb.VisualShaderNodeGroupBase(ptr) })}
