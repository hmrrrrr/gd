package VisualShaderNodeCustom

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/VisualShaderNode"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
By inheriting this class you can create a custom [VisualShader] script addon which will be automatically added to the Visual Shader Editor. The [VisualShaderNode]'s behavior is defined by overriding the provided virtual methods.
In order for the node to be registered as an editor addon, you must use the [code]@tool[/code] annotation and provide a [code]class_name[/code] for your custom script. For example:
[codeblock]
@tool
extends VisualShaderNodeCustom
class_name VisualShaderNodeNoise
[/codeblock]

	// VisualShaderNodeCustom methods that can be overridden by a [Class] that extends it.
	type VisualShaderNodeCustom interface {
		//Override this method to define the name of the associated custom node in the Visual Shader Editor's members dialog and graph.
		//Defining this method is [b]optional[/b], but recommended. If not overridden, the node will be named as "Unnamed".
		GetName() string
		//Override this method to define the description of the associated custom node in the Visual Shader Editor's members dialog.
		//Defining this method is [b]optional[/b].
		GetDescription() string
		//Override this method to define the path to the associated custom node in the Visual Shader Editor's members dialog. The path may look like [code]"MyGame/MyFunctions/Noise"[/code].
		//Defining this method is [b]optional[/b]. If not overridden, the node will be filed under the "Addons" category.
		GetCategory() string
		//Override this method to define the return icon of the associated custom node in the Visual Shader Editor's members dialog.
		//Defining this method is [b]optional[/b]. If not overridden, no return icon is shown.
		GetReturnIconType() classdb.VisualShaderNodePortType
		//Override this method to define the number of input ports of the associated custom node.
		//Defining this method is [b]required[/b]. If not overridden, the node has no input ports.
		GetInputPortCount() int
		//Override this method to define the returned type of each input port of the associated custom node (see [enum VisualShaderNode.PortType] for possible types).
		//Defining this method is [b]optional[/b], but recommended. If not overridden, input ports will return the [constant VisualShaderNode.PORT_TYPE_SCALAR] type.
		GetInputPortType(port int) classdb.VisualShaderNodePortType
		//Override this method to define the names of input ports of the associated custom node. The names are used both for the input slots in the editor and as identifiers in the shader code, and are passed in the [code]input_vars[/code] array in [method _get_code].
		//Defining this method is [b]optional[/b], but recommended. If not overridden, input ports are named as [code]"in" + str(port)[/code].
		GetInputPortName(port int) string
		//Override this method to define the default value for the specified input port. Prefer use this over [method VisualShaderNode.set_input_port_default_value].
		//Defining this method is [b]required[/b]. If not overridden, the node has no default values for their input ports.
		GetInputPortDefaultValue(port int) any
		//Override this method to define the input port which should be connected by default when this node is created as a result of dragging a connection from an existing node to the empty space on the graph.
		//Defining this method is [b]optional[/b]. If not overridden, the connection will be created to the first valid port.
		GetDefaultInputPort(atype classdb.VisualShaderNodePortType) int
		//Override this method to define the number of output ports of the associated custom node.
		//Defining this method is [b]required[/b]. If not overridden, the node has no output ports.
		GetOutputPortCount() int
		//Override this method to define the returned type of each output port of the associated custom node (see [enum VisualShaderNode.PortType] for possible types).
		//Defining this method is [b]optional[/b], but recommended. If not overridden, output ports will return the [constant VisualShaderNode.PORT_TYPE_SCALAR] type.
		GetOutputPortType(port int) classdb.VisualShaderNodePortType
		//Override this method to define the names of output ports of the associated custom node. The names are used both for the output slots in the editor and as identifiers in the shader code, and are passed in the [code]output_vars[/code] array in [method _get_code].
		//Defining this method is [b]optional[/b], but recommended. If not overridden, output ports are named as [code]"out" + str(port)[/code].
		GetOutputPortName(port int) string
		//Override this method to define the number of the properties.
		//Defining this method is [b]optional[/b].
		GetPropertyCount() int
		//Override this method to define the names of the property of the associated custom node.
		//Defining this method is [b]optional[/b].
		GetPropertyName(index int) string
		//Override this method to define the default index of the property of the associated custom node.
		//Defining this method is [b]optional[/b].
		GetPropertyDefaultIndex(index int) int
		//Override this method to define the options inside the drop-down list property of the associated custom node.
		//Defining this method is [b]optional[/b].
		GetPropertyOptions(index int) []string
		//Override this method to define the actual shader code of the associated custom node. The shader code should be returned as a string, which can have multiple lines (the [code]"""[/code] multiline string construct can be used for convenience).
		//The [param input_vars] and [param output_vars] arrays contain the string names of the various input and output variables, as defined by [code]_get_input_*[/code] and [code]_get_output_*[/code] virtual methods in this class.
		//The output ports can be assigned values in the shader code. For example, [code]return output_vars[0] + " = " + input_vars[0] + ";"[/code].
		//You can customize the generated code based on the shader [param mode] (see [enum Shader.Mode]) and/or [param type] (see [enum VisualShader.Type]).
		//Defining this method is [b]required[/b].
		GetCode(input_vars gd.Array, output_vars gd.Array, mode classdb.ShaderMode, atype classdb.VisualShaderType) string
		//Override this method to add a shader code to the beginning of each shader function (once). The shader code should be returned as a string, which can have multiple lines (the [code]"""[/code] multiline string construct can be used for convenience).
		//If there are multiple custom nodes of different types which use this feature the order of each insertion is undefined.
		//You can customize the generated code based on the shader [param mode] (see [enum Shader.Mode]) and/or [param type] (see [enum VisualShader.Type]).
		//Defining this method is [b]optional[/b].
		GetFuncCode(mode classdb.ShaderMode, atype classdb.VisualShaderType) string
		//Override this method to add shader code on top of the global shader, to define your own standard library of reusable methods, varyings, constants, uniforms, etc. The shader code should be returned as a string, which can have multiple lines (the [code]"""[/code] multiline string construct can be used for convenience).
		//Be careful with this functionality as it can cause name conflicts with other custom nodes, so be sure to give the defined entities unique names.
		//You can customize the generated code based on the shader [param mode] (see [enum Shader.Mode]).
		//Defining this method is [b]optional[/b].
		GetGlobalCode(mode classdb.ShaderMode) string
		//Override this method to enable high-end mark in the Visual Shader Editor's members dialog.
		//Defining this method is [b]optional[/b]. If not overridden, it's [code]false[/code].
		IsHighend() bool
		//Override this method to prevent the node to be visible in the member dialog for the certain [param mode] (see [enum Shader.Mode]) and/or [param type] (see [enum VisualShader.Type]).
		//Defining this method is [b]optional[/b]. If not overridden, it's [code]true[/code].
		IsAvailable(mode classdb.ShaderMode, atype classdb.VisualShaderType) bool
	}
*/
type Instance [1]classdb.VisualShaderNodeCustom

/*
Override this method to define the name of the associated custom node in the Visual Shader Editor's members dialog and graph.
Defining this method is [b]optional[/b], but recommended. If not overridden, the node will be named as "Unnamed".
*/
func (Instance) _get_name(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewString(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to define the description of the associated custom node in the Visual Shader Editor's members dialog.
Defining this method is [b]optional[/b].
*/
func (Instance) _get_description(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewString(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to define the path to the associated custom node in the Visual Shader Editor's members dialog. The path may look like [code]"MyGame/MyFunctions/Noise"[/code].
Defining this method is [b]optional[/b]. If not overridden, the node will be filed under the "Addons" category.
*/
func (Instance) _get_category(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewString(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to define the return icon of the associated custom node in the Visual Shader Editor's members dialog.
Defining this method is [b]optional[/b]. If not overridden, no return icon is shown.
*/
func (Instance) _get_return_icon_type(impl func(ptr unsafe.Pointer) classdb.VisualShaderNodePortType) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to define the number of input ports of the associated custom node.
Defining this method is [b]required[/b]. If not overridden, the node has no input ports.
*/
func (Instance) _get_input_port_count(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Override this method to define the returned type of each input port of the associated custom node (see [enum VisualShaderNode.PortType] for possible types).
Defining this method is [b]optional[/b], but recommended. If not overridden, input ports will return the [constant VisualShaderNode.PORT_TYPE_SCALAR] type.
*/
func (Instance) _get_input_port_type(impl func(ptr unsafe.Pointer, port int) classdb.VisualShaderNodePortType) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var port = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(port))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to define the names of input ports of the associated custom node. The names are used both for the input slots in the editor and as identifiers in the shader code, and are passed in the [code]input_vars[/code] array in [method _get_code].
Defining this method is [b]optional[/b], but recommended. If not overridden, input ports are named as [code]"in" + str(port)[/code].
*/
func (Instance) _get_input_port_name(impl func(ptr unsafe.Pointer, port int) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var port = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(port))
		ptr, ok := pointers.End(gd.NewString(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to define the default value for the specified input port. Prefer use this over [method VisualShaderNode.set_input_port_default_value].
Defining this method is [b]required[/b]. If not overridden, the node has no default values for their input ports.
*/
func (Instance) _get_input_port_default_value(impl func(ptr unsafe.Pointer, port int) any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var port = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(port))
		ptr, ok := pointers.End(gd.NewVariant(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to define the input port which should be connected by default when this node is created as a result of dragging a connection from an existing node to the empty space on the graph.
Defining this method is [b]optional[/b]. If not overridden, the connection will be created to the first valid port.
*/
func (Instance) _get_default_input_port(impl func(ptr unsafe.Pointer, atype classdb.VisualShaderNodePortType) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var atype = gd.UnsafeGet[classdb.VisualShaderNodePortType](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, atype)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Override this method to define the number of output ports of the associated custom node.
Defining this method is [b]required[/b]. If not overridden, the node has no output ports.
*/
func (Instance) _get_output_port_count(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Override this method to define the returned type of each output port of the associated custom node (see [enum VisualShaderNode.PortType] for possible types).
Defining this method is [b]optional[/b], but recommended. If not overridden, output ports will return the [constant VisualShaderNode.PORT_TYPE_SCALAR] type.
*/
func (Instance) _get_output_port_type(impl func(ptr unsafe.Pointer, port int) classdb.VisualShaderNodePortType) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var port = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(port))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to define the names of output ports of the associated custom node. The names are used both for the output slots in the editor and as identifiers in the shader code, and are passed in the [code]output_vars[/code] array in [method _get_code].
Defining this method is [b]optional[/b], but recommended. If not overridden, output ports are named as [code]"out" + str(port)[/code].
*/
func (Instance) _get_output_port_name(impl func(ptr unsafe.Pointer, port int) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var port = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(port))
		ptr, ok := pointers.End(gd.NewString(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to define the number of the properties.
Defining this method is [b]optional[/b].
*/
func (Instance) _get_property_count(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Override this method to define the names of the property of the associated custom node.
Defining this method is [b]optional[/b].
*/
func (Instance) _get_property_name(impl func(ptr unsafe.Pointer, index int) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		ptr, ok := pointers.End(gd.NewString(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to define the default index of the property of the associated custom node.
Defining this method is [b]optional[/b].
*/
func (Instance) _get_property_default_index(impl func(ptr unsafe.Pointer, index int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Override this method to define the options inside the drop-down list property of the associated custom node.
Defining this method is [b]optional[/b].
*/
func (Instance) _get_property_options(impl func(ptr unsafe.Pointer, index int) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		ptr, ok := pointers.End(gd.NewPackedStringSlice(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to define the actual shader code of the associated custom node. The shader code should be returned as a string, which can have multiple lines (the [code]"""[/code] multiline string construct can be used for convenience).
The [param input_vars] and [param output_vars] arrays contain the string names of the various input and output variables, as defined by [code]_get_input_*[/code] and [code]_get_output_*[/code] virtual methods in this class.
The output ports can be assigned values in the shader code. For example, [code]return output_vars[0] + " = " + input_vars[0] + ";"[/code].
You can customize the generated code based on the shader [param mode] (see [enum Shader.Mode]) and/or [param type] (see [enum VisualShader.Type]).
Defining this method is [b]required[/b].
*/
func (Instance) _get_code(impl func(ptr unsafe.Pointer, input_vars gd.Array, output_vars gd.Array, mode classdb.ShaderMode, atype classdb.VisualShaderType) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var input_vars = pointers.New[gd.Array](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(input_vars)
		var output_vars = pointers.New[gd.Array](gd.UnsafeGet[[1]uintptr](p_args, 1))
		defer pointers.End(output_vars)
		var mode = gd.UnsafeGet[classdb.ShaderMode](p_args, 2)
		var atype = gd.UnsafeGet[classdb.VisualShaderType](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, input_vars, output_vars, mode, atype)
		ptr, ok := pointers.End(gd.NewString(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to add a shader code to the beginning of each shader function (once). The shader code should be returned as a string, which can have multiple lines (the [code]"""[/code] multiline string construct can be used for convenience).
If there are multiple custom nodes of different types which use this feature the order of each insertion is undefined.
You can customize the generated code based on the shader [param mode] (see [enum Shader.Mode]) and/or [param type] (see [enum VisualShader.Type]).
Defining this method is [b]optional[/b].
*/
func (Instance) _get_func_code(impl func(ptr unsafe.Pointer, mode classdb.ShaderMode, atype classdb.VisualShaderType) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var mode = gd.UnsafeGet[classdb.ShaderMode](p_args, 0)
		var atype = gd.UnsafeGet[classdb.VisualShaderType](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, mode, atype)
		ptr, ok := pointers.End(gd.NewString(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to add shader code on top of the global shader, to define your own standard library of reusable methods, varyings, constants, uniforms, etc. The shader code should be returned as a string, which can have multiple lines (the [code]"""[/code] multiline string construct can be used for convenience).
Be careful with this functionality as it can cause name conflicts with other custom nodes, so be sure to give the defined entities unique names.
You can customize the generated code based on the shader [param mode] (see [enum Shader.Mode]).
Defining this method is [b]optional[/b].
*/
func (Instance) _get_global_code(impl func(ptr unsafe.Pointer, mode classdb.ShaderMode) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var mode = gd.UnsafeGet[classdb.ShaderMode](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, mode)
		ptr, ok := pointers.End(gd.NewString(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to enable high-end mark in the Visual Shader Editor's members dialog.
Defining this method is [b]optional[/b]. If not overridden, it's [code]false[/code].
*/
func (Instance) _is_highend(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to prevent the node to be visible in the member dialog for the certain [param mode] (see [enum Shader.Mode]) and/or [param type] (see [enum VisualShader.Type]).
Defining this method is [b]optional[/b]. If not overridden, it's [code]true[/code].
*/
func (Instance) _is_available(impl func(ptr unsafe.Pointer, mode classdb.ShaderMode, atype classdb.VisualShaderType) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var mode = gd.UnsafeGet[classdb.ShaderMode](p_args, 0)
		var atype = gd.UnsafeGet[classdb.VisualShaderType](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, mode, atype)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns the selected index of the drop-down list option within a graph. You may use this function to define the specific behavior in the [method _get_code] or [method _get_global_code].
*/
func (self Instance) GetOptionIndex(option int) int {
	return int(int(class(self).GetOptionIndex(gd.Int(option))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.VisualShaderNodeCustom

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeCustom"))
	return Instance{classdb.VisualShaderNodeCustom(object)}
}

/*
Override this method to define the name of the associated custom node in the Visual Shader Editor's members dialog and graph.
Defining this method is [b]optional[/b], but recommended. If not overridden, the node will be named as "Unnamed".
*/
func (class) _get_name(impl func(ptr unsafe.Pointer) gd.String) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to define the description of the associated custom node in the Visual Shader Editor's members dialog.
Defining this method is [b]optional[/b].
*/
func (class) _get_description(impl func(ptr unsafe.Pointer) gd.String) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to define the path to the associated custom node in the Visual Shader Editor's members dialog. The path may look like [code]"MyGame/MyFunctions/Noise"[/code].
Defining this method is [b]optional[/b]. If not overridden, the node will be filed under the "Addons" category.
*/
func (class) _get_category(impl func(ptr unsafe.Pointer) gd.String) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to define the return icon of the associated custom node in the Visual Shader Editor's members dialog.
Defining this method is [b]optional[/b]. If not overridden, no return icon is shown.
*/
func (class) _get_return_icon_type(impl func(ptr unsafe.Pointer) classdb.VisualShaderNodePortType) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to define the number of input ports of the associated custom node.
Defining this method is [b]required[/b]. If not overridden, the node has no input ports.
*/
func (class) _get_input_port_count(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to define the returned type of each input port of the associated custom node (see [enum VisualShaderNode.PortType] for possible types).
Defining this method is [b]optional[/b], but recommended. If not overridden, input ports will return the [constant VisualShaderNode.PORT_TYPE_SCALAR] type.
*/
func (class) _get_input_port_type(impl func(ptr unsafe.Pointer, port gd.Int) classdb.VisualShaderNodePortType) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var port = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, port)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to define the names of input ports of the associated custom node. The names are used both for the input slots in the editor and as identifiers in the shader code, and are passed in the [code]input_vars[/code] array in [method _get_code].
Defining this method is [b]optional[/b], but recommended. If not overridden, input ports are named as [code]"in" + str(port)[/code].
*/
func (class) _get_input_port_name(impl func(ptr unsafe.Pointer, port gd.Int) gd.String) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var port = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, port)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to define the default value for the specified input port. Prefer use this over [method VisualShaderNode.set_input_port_default_value].
Defining this method is [b]required[/b]. If not overridden, the node has no default values for their input ports.
*/
func (class) _get_input_port_default_value(impl func(ptr unsafe.Pointer, port gd.Int) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var port = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, port)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to define the input port which should be connected by default when this node is created as a result of dragging a connection from an existing node to the empty space on the graph.
Defining this method is [b]optional[/b]. If not overridden, the connection will be created to the first valid port.
*/
func (class) _get_default_input_port(impl func(ptr unsafe.Pointer, atype classdb.VisualShaderNodePortType) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var atype = gd.UnsafeGet[classdb.VisualShaderNodePortType](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, atype)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to define the number of output ports of the associated custom node.
Defining this method is [b]required[/b]. If not overridden, the node has no output ports.
*/
func (class) _get_output_port_count(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to define the returned type of each output port of the associated custom node (see [enum VisualShaderNode.PortType] for possible types).
Defining this method is [b]optional[/b], but recommended. If not overridden, output ports will return the [constant VisualShaderNode.PORT_TYPE_SCALAR] type.
*/
func (class) _get_output_port_type(impl func(ptr unsafe.Pointer, port gd.Int) classdb.VisualShaderNodePortType) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var port = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, port)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to define the names of output ports of the associated custom node. The names are used both for the output slots in the editor and as identifiers in the shader code, and are passed in the [code]output_vars[/code] array in [method _get_code].
Defining this method is [b]optional[/b], but recommended. If not overridden, output ports are named as [code]"out" + str(port)[/code].
*/
func (class) _get_output_port_name(impl func(ptr unsafe.Pointer, port gd.Int) gd.String) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var port = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, port)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to define the number of the properties.
Defining this method is [b]optional[/b].
*/
func (class) _get_property_count(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to define the names of the property of the associated custom node.
Defining this method is [b]optional[/b].
*/
func (class) _get_property_name(impl func(ptr unsafe.Pointer, index gd.Int) gd.String) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to define the default index of the property of the associated custom node.
Defining this method is [b]optional[/b].
*/
func (class) _get_property_default_index(impl func(ptr unsafe.Pointer, index gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to define the options inside the drop-down list property of the associated custom node.
Defining this method is [b]optional[/b].
*/
func (class) _get_property_options(impl func(ptr unsafe.Pointer, index gd.Int) gd.PackedStringArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to define the actual shader code of the associated custom node. The shader code should be returned as a string, which can have multiple lines (the [code]"""[/code] multiline string construct can be used for convenience).
The [param input_vars] and [param output_vars] arrays contain the string names of the various input and output variables, as defined by [code]_get_input_*[/code] and [code]_get_output_*[/code] virtual methods in this class.
The output ports can be assigned values in the shader code. For example, [code]return output_vars[0] + " = " + input_vars[0] + ";"[/code].
You can customize the generated code based on the shader [param mode] (see [enum Shader.Mode]) and/or [param type] (see [enum VisualShader.Type]).
Defining this method is [b]required[/b].
*/
func (class) _get_code(impl func(ptr unsafe.Pointer, input_vars gd.Array, output_vars gd.Array, mode classdb.ShaderMode, atype classdb.VisualShaderType) gd.String) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var input_vars = pointers.New[gd.Array](gd.UnsafeGet[[1]uintptr](p_args, 0))
		var output_vars = pointers.New[gd.Array](gd.UnsafeGet[[1]uintptr](p_args, 1))
		var mode = gd.UnsafeGet[classdb.ShaderMode](p_args, 2)
		var atype = gd.UnsafeGet[classdb.VisualShaderType](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, input_vars, output_vars, mode, atype)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to add a shader code to the beginning of each shader function (once). The shader code should be returned as a string, which can have multiple lines (the [code]"""[/code] multiline string construct can be used for convenience).
If there are multiple custom nodes of different types which use this feature the order of each insertion is undefined.
You can customize the generated code based on the shader [param mode] (see [enum Shader.Mode]) and/or [param type] (see [enum VisualShader.Type]).
Defining this method is [b]optional[/b].
*/
func (class) _get_func_code(impl func(ptr unsafe.Pointer, mode classdb.ShaderMode, atype classdb.VisualShaderType) gd.String) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var mode = gd.UnsafeGet[classdb.ShaderMode](p_args, 0)
		var atype = gd.UnsafeGet[classdb.VisualShaderType](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, mode, atype)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to add shader code on top of the global shader, to define your own standard library of reusable methods, varyings, constants, uniforms, etc. The shader code should be returned as a string, which can have multiple lines (the [code]"""[/code] multiline string construct can be used for convenience).
Be careful with this functionality as it can cause name conflicts with other custom nodes, so be sure to give the defined entities unique names.
You can customize the generated code based on the shader [param mode] (see [enum Shader.Mode]).
Defining this method is [b]optional[/b].
*/
func (class) _get_global_code(impl func(ptr unsafe.Pointer, mode classdb.ShaderMode) gd.String) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var mode = gd.UnsafeGet[classdb.ShaderMode](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, mode)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method to enable high-end mark in the Visual Shader Editor's members dialog.
Defining this method is [b]optional[/b]. If not overridden, it's [code]false[/code].
*/
func (class) _is_highend(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to prevent the node to be visible in the member dialog for the certain [param mode] (see [enum Shader.Mode]) and/or [param type] (see [enum VisualShader.Type]).
Defining this method is [b]optional[/b]. If not overridden, it's [code]true[/code].
*/
func (class) _is_available(impl func(ptr unsafe.Pointer, mode classdb.ShaderMode, atype classdb.VisualShaderType) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var mode = gd.UnsafeGet[classdb.ShaderMode](p_args, 0)
		var atype = gd.UnsafeGet[classdb.VisualShaderType](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, mode, atype)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns the selected index of the drop-down list option within a graph. You may use this function to define the specific behavior in the [method _get_code] or [method _get_global_code].
*/
//go:nosplit
func (self class) GetOptionIndex(option gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, option)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeCustom.Bind_get_option_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeCustom() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVisualShaderNodeCustom() Instance {
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
	case "_get_name":
		return reflect.ValueOf(self._get_name)
	case "_get_description":
		return reflect.ValueOf(self._get_description)
	case "_get_category":
		return reflect.ValueOf(self._get_category)
	case "_get_return_icon_type":
		return reflect.ValueOf(self._get_return_icon_type)
	case "_get_input_port_count":
		return reflect.ValueOf(self._get_input_port_count)
	case "_get_input_port_type":
		return reflect.ValueOf(self._get_input_port_type)
	case "_get_input_port_name":
		return reflect.ValueOf(self._get_input_port_name)
	case "_get_input_port_default_value":
		return reflect.ValueOf(self._get_input_port_default_value)
	case "_get_default_input_port":
		return reflect.ValueOf(self._get_default_input_port)
	case "_get_output_port_count":
		return reflect.ValueOf(self._get_output_port_count)
	case "_get_output_port_type":
		return reflect.ValueOf(self._get_output_port_type)
	case "_get_output_port_name":
		return reflect.ValueOf(self._get_output_port_name)
	case "_get_property_count":
		return reflect.ValueOf(self._get_property_count)
	case "_get_property_name":
		return reflect.ValueOf(self._get_property_name)
	case "_get_property_default_index":
		return reflect.ValueOf(self._get_property_default_index)
	case "_get_property_options":
		return reflect.ValueOf(self._get_property_options)
	case "_get_code":
		return reflect.ValueOf(self._get_code)
	case "_get_func_code":
		return reflect.ValueOf(self._get_func_code)
	case "_get_global_code":
		return reflect.ValueOf(self._get_global_code)
	case "_is_highend":
		return reflect.ValueOf(self._is_highend)
	case "_is_available":
		return reflect.ValueOf(self._is_available)
	default:
		return gd.VirtualByName(self.AsVisualShaderNode(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_name":
		return reflect.ValueOf(self._get_name)
	case "_get_description":
		return reflect.ValueOf(self._get_description)
	case "_get_category":
		return reflect.ValueOf(self._get_category)
	case "_get_return_icon_type":
		return reflect.ValueOf(self._get_return_icon_type)
	case "_get_input_port_count":
		return reflect.ValueOf(self._get_input_port_count)
	case "_get_input_port_type":
		return reflect.ValueOf(self._get_input_port_type)
	case "_get_input_port_name":
		return reflect.ValueOf(self._get_input_port_name)
	case "_get_input_port_default_value":
		return reflect.ValueOf(self._get_input_port_default_value)
	case "_get_default_input_port":
		return reflect.ValueOf(self._get_default_input_port)
	case "_get_output_port_count":
		return reflect.ValueOf(self._get_output_port_count)
	case "_get_output_port_type":
		return reflect.ValueOf(self._get_output_port_type)
	case "_get_output_port_name":
		return reflect.ValueOf(self._get_output_port_name)
	case "_get_property_count":
		return reflect.ValueOf(self._get_property_count)
	case "_get_property_name":
		return reflect.ValueOf(self._get_property_name)
	case "_get_property_default_index":
		return reflect.ValueOf(self._get_property_default_index)
	case "_get_property_options":
		return reflect.ValueOf(self._get_property_options)
	case "_get_code":
		return reflect.ValueOf(self._get_code)
	case "_get_func_code":
		return reflect.ValueOf(self._get_func_code)
	case "_get_global_code":
		return reflect.ValueOf(self._get_global_code)
	case "_is_highend":
		return reflect.ValueOf(self._is_highend)
	case "_is_available":
		return reflect.ValueOf(self._is_available)
	default:
		return gd.VirtualByName(self.AsVisualShaderNode(), name)
	}
}
func init() {
	classdb.Register("VisualShaderNodeCustom", func(ptr gd.Object) any { return classdb.VisualShaderNodeCustom(ptr) })
}
