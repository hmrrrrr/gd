package AnimationNode

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Resource"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Path"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Base resource for [AnimationTree] nodes. In general, it's not used directly, but you can create custom ones with custom blending formulas.
Inherit this when creating animation nodes mainly for use in [AnimationNodeBlendTree], otherwise [AnimationRootNode] should be used instead.
You can access the time information as read-only parameter which is processed and stored in the previous frame for all nodes except [AnimationNodeOutput].
[b]Note:[/b] If multiple inputs exist in the [AnimationNode], which time information takes precedence depends on the type of [AnimationNode].
[codeblock]
var current_length = $AnimationTree[parameters/AnimationNodeName/current_length]
var current_position = $AnimationTree[parameters/AnimationNodeName/current_position]
var current_delta = $AnimationTree[parameters/AnimationNodeName/current_delta]
[/codeblock]

	// AnimationNode methods that can be overridden by a [Class] that extends it.
	type AnimationNode interface {
		//When inheriting from [AnimationRootNode], implement this virtual method to return all child animation nodes in order as a [code]name: node[/code] dictionary.
		GetChildNodes() Dictionary.Any
		//When inheriting from [AnimationRootNode], implement this virtual method to return a list of the properties on this animation node. Parameters are custom local memory used for your animation nodes, given a resource can be reused in multiple trees. Format is similar to [method Object.get_property_list].
		GetParameterList() Array.Any
		//When inheriting from [AnimationRootNode], implement this virtual method to return a child animation node by its [param name].
		GetChildByName(name string) objects.AnimationNode
		//When inheriting from [AnimationRootNode], implement this virtual method to return the default value of a [param parameter]. Parameters are custom local memory used for your animation nodes, given a resource can be reused in multiple trees.
		GetParameterDefaultValue(parameter string) any
		//When inheriting from [AnimationRootNode], implement this virtual method to return whether the [param parameter] is read-only. Parameters are custom local memory used for your animation nodes, given a resource can be reused in multiple trees.
		IsParameterReadOnly(parameter string) bool
		//When inheriting from [AnimationRootNode], implement this virtual method to run some code when this animation node is processed. The [param time] parameter is a relative delta, unless [param seek] is [code]true[/code], in which case it is absolute.
		//Here, call the [method blend_input], [method blend_node] or [method blend_animation] functions. You can also use [method get_parameter] and [method set_parameter] to modify local memory.
		//This function should return the delta.
		Process(time Float.X, seek bool, is_external_seeking bool, test_only bool) Float.X
		//When inheriting from [AnimationRootNode], implement this virtual method to override the text caption for this animation node.
		GetCaption() string
		//When inheriting from [AnimationRootNode], implement this virtual method to return whether the blend tree editor should display filter editing on this animation node.
		HasFilter() bool
	}
*/
type Instance [1]classdb.AnimationNode
type Any interface {
	gd.IsClass
	AsAnimationNode() Instance
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to return all child animation nodes in order as a [code]name: node[/code] dictionary.
*/
func (Instance) _get_child_nodes(impl func(ptr unsafe.Pointer) Dictionary.Any) (cb gd.ExtensionClassCallVirtualFunc) {
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
When inheriting from [AnimationRootNode], implement this virtual method to return a list of the properties on this animation node. Parameters are custom local memory used for your animation nodes, given a resource can be reused in multiple trees. Format is similar to [method Object.get_property_list].
*/
func (Instance) _get_parameter_list(impl func(ptr unsafe.Pointer) Array.Any) (cb gd.ExtensionClassCallVirtualFunc) {
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
When inheriting from [AnimationRootNode], implement this virtual method to return a child animation node by its [param name].
*/
func (Instance) _get_child_by_name(impl func(ptr unsafe.Pointer, name string) objects.AnimationNode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var name = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(name)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, name.String())
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to return the default value of a [param parameter]. Parameters are custom local memory used for your animation nodes, given a resource can be reused in multiple trees.
*/
func (Instance) _get_parameter_default_value(impl func(ptr unsafe.Pointer, parameter string) any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var parameter = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(parameter)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, parameter.String())
		ptr, ok := pointers.End(gd.NewVariant(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to return whether the [param parameter] is read-only. Parameters are custom local memory used for your animation nodes, given a resource can be reused in multiple trees.
*/
func (Instance) _is_parameter_read_only(impl func(ptr unsafe.Pointer, parameter string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var parameter = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(parameter)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, parameter.String())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to run some code when this animation node is processed. The [param time] parameter is a relative delta, unless [param seek] is [code]true[/code], in which case it is absolute.
Here, call the [method blend_input], [method blend_node] or [method blend_animation] functions. You can also use [method get_parameter] and [method set_parameter] to modify local memory.
This function should return the delta.
*/
func (Instance) _process(impl func(ptr unsafe.Pointer, time Float.X, seek bool, is_external_seeking bool, test_only bool) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var time = gd.UnsafeGet[gd.Float](p_args, 0)
		var seek = gd.UnsafeGet[bool](p_args, 1)
		var is_external_seeking = gd.UnsafeGet[bool](p_args, 2)
		var test_only = gd.UnsafeGet[bool](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, Float.X(time), seek, is_external_seeking, test_only)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to override the text caption for this animation node.
*/
func (Instance) _get_caption(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
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
When inheriting from [AnimationRootNode], implement this virtual method to return whether the blend tree editor should display filter editing on this animation node.
*/
func (Instance) _has_filter(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Adds an input to the animation node. This is only useful for animation nodes created for use in an [AnimationNodeBlendTree]. If the addition fails, returns [code]false[/code].
*/
func (self Instance) AddInput(name string) bool {
	return bool(class(self).AddInput(gd.NewString(name)))
}

/*
Removes an input, call this only when inactive.
*/
func (self Instance) RemoveInput(index int) {
	class(self).RemoveInput(gd.Int(index))
}

/*
Sets the name of the input at the given [param input] index. If the setting fails, returns [code]false[/code].
*/
func (self Instance) SetInputName(input int, name string) bool {
	return bool(class(self).SetInputName(gd.Int(input), gd.NewString(name)))
}

/*
Gets the name of an input by index.
*/
func (self Instance) GetInputName(input int) string {
	return string(class(self).GetInputName(gd.Int(input)).String())
}

/*
Amount of inputs in this animation node, only useful for animation nodes that go into [AnimationNodeBlendTree].
*/
func (self Instance) GetInputCount() int {
	return int(int(class(self).GetInputCount()))
}

/*
Returns the input index which corresponds to [param name]. If not found, returns [code]-1[/code].
*/
func (self Instance) FindInput(name string) int {
	return int(int(class(self).FindInput(gd.NewString(name))))
}

/*
Adds or removes a path for the filter.
*/
func (self Instance) SetFilterPath(path Path.String, enable bool) {
	class(self).SetFilterPath(gd.NewString(string(path)).NodePath(), enable)
}

/*
Returns whether the given path is filtered.
*/
func (self Instance) IsPathFiltered(path Path.String) bool {
	return bool(class(self).IsPathFiltered(gd.NewString(string(path)).NodePath()))
}

/*
Blend an animation by [param blend] amount (name must be valid in the linked [AnimationPlayer]). A [param time] and [param delta] may be passed, as well as whether [param seeked] happened.
A [param looped_flag] is used by internal processing immediately after the loop. See also [enum Animation.LoopedFlag].
*/
func (self Instance) BlendAnimation(animation string, time Float.X, delta Float.X, seeked bool, is_external_seeking bool, blend Float.X) {
	class(self).BlendAnimation(gd.NewStringName(animation), gd.Float(time), gd.Float(delta), seeked, is_external_seeking, gd.Float(blend), 0)
}

/*
Blend another animation node (in case this animation node contains child animation nodes). This function is only useful if you inherit from [AnimationRootNode] instead, otherwise editors will not display your animation node for addition.
*/
func (self Instance) BlendNode(name string, node objects.AnimationNode, time Float.X, seek bool, is_external_seeking bool, blend Float.X) Float.X {
	return Float.X(Float.X(class(self).BlendNode(gd.NewStringName(name), node, gd.Float(time), seek, is_external_seeking, gd.Float(blend), 0, true, false)))
}

/*
Blend an input. This is only useful for animation nodes created for an [AnimationNodeBlendTree]. The [param time] parameter is a relative delta, unless [param seek] is [code]true[/code], in which case it is absolute. A filter mode may be optionally passed (see [enum FilterAction] for options).
*/
func (self Instance) BlendInput(input_index int, time Float.X, seek bool, is_external_seeking bool, blend Float.X) Float.X {
	return Float.X(Float.X(class(self).BlendInput(gd.Int(input_index), gd.Float(time), seek, is_external_seeking, gd.Float(blend), 0, true, false)))
}

/*
Sets a custom parameter. These are used as local memory, because resources can be reused across the tree or scenes.
*/
func (self Instance) SetParameter(name string, value any) {
	class(self).SetParameter(gd.NewStringName(name), gd.NewVariant(value))
}

/*
Gets the value of a parameter. Parameters are custom local memory used for your animation nodes, given a resource can be reused in multiple trees.
*/
func (self Instance) GetParameter(name string) any {
	return any(class(self).GetParameter(gd.NewStringName(name)).Interface())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AnimationNode

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimationNode"))
	return Instance{*(*classdb.AnimationNode)(unsafe.Pointer(&object))}
}

func (self Instance) FilterEnabled() bool {
	return bool(class(self).IsFilterEnabled())
}

func (self Instance) SetFilterEnabled(value bool) {
	class(self).SetFilterEnabled(value)
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to return all child animation nodes in order as a [code]name: node[/code] dictionary.
*/
func (class) _get_child_nodes(impl func(ptr unsafe.Pointer) gd.Dictionary) (cb gd.ExtensionClassCallVirtualFunc) {
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
When inheriting from [AnimationRootNode], implement this virtual method to return a list of the properties on this animation node. Parameters are custom local memory used for your animation nodes, given a resource can be reused in multiple trees. Format is similar to [method Object.get_property_list].
*/
func (class) _get_parameter_list(impl func(ptr unsafe.Pointer) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
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
When inheriting from [AnimationRootNode], implement this virtual method to return a child animation node by its [param name].
*/
func (class) _get_child_by_name(impl func(ptr unsafe.Pointer, name gd.StringName) objects.AnimationNode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var name = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, name)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to return the default value of a [param parameter]. Parameters are custom local memory used for your animation nodes, given a resource can be reused in multiple trees.
*/
func (class) _get_parameter_default_value(impl func(ptr unsafe.Pointer, parameter gd.StringName) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var parameter = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, parameter)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to return whether the [param parameter] is read-only. Parameters are custom local memory used for your animation nodes, given a resource can be reused in multiple trees.
*/
func (class) _is_parameter_read_only(impl func(ptr unsafe.Pointer, parameter gd.StringName) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var parameter = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, parameter)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to run some code when this animation node is processed. The [param time] parameter is a relative delta, unless [param seek] is [code]true[/code], in which case it is absolute.
Here, call the [method blend_input], [method blend_node] or [method blend_animation] functions. You can also use [method get_parameter] and [method set_parameter] to modify local memory.
This function should return the delta.
*/
func (class) _process(impl func(ptr unsafe.Pointer, time gd.Float, seek bool, is_external_seeking bool, test_only bool) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var time = gd.UnsafeGet[gd.Float](p_args, 0)
		var seek = gd.UnsafeGet[bool](p_args, 1)
		var is_external_seeking = gd.UnsafeGet[bool](p_args, 2)
		var test_only = gd.UnsafeGet[bool](p_args, 3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, time, seek, is_external_seeking, test_only)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to override the text caption for this animation node.
*/
func (class) _get_caption(impl func(ptr unsafe.Pointer) gd.String) (cb gd.ExtensionClassCallVirtualFunc) {
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
When inheriting from [AnimationRootNode], implement this virtual method to return whether the blend tree editor should display filter editing on this animation node.
*/
func (class) _has_filter(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Adds an input to the animation node. This is only useful for animation nodes created for use in an [AnimationNodeBlendTree]. If the addition fails, returns [code]false[/code].
*/
//go:nosplit
func (self class) AddInput(name gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNode.Bind_add_input, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes an input, call this only when inactive.
*/
//go:nosplit
func (self class) RemoveInput(index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNode.Bind_remove_input, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the name of the input at the given [param input] index. If the setting fails, returns [code]false[/code].
*/
//go:nosplit
func (self class) SetInputName(input gd.Int, name gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, input)
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNode.Bind_set_input_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Gets the name of an input by index.
*/
//go:nosplit
func (self class) GetInputName(input gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, input)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNode.Bind_get_input_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Amount of inputs in this animation node, only useful for animation nodes that go into [AnimationNodeBlendTree].
*/
//go:nosplit
func (self class) GetInputCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNode.Bind_get_input_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the input index which corresponds to [param name]. If not found, returns [code]-1[/code].
*/
//go:nosplit
func (self class) FindInput(name gd.String) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNode.Bind_find_input, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds or removes a path for the filter.
*/
//go:nosplit
func (self class) SetFilterPath(path gd.NodePath, enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNode.Bind_set_filter_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether the given path is filtered.
*/
//go:nosplit
func (self class) IsPathFiltered(path gd.NodePath) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNode.Bind_is_path_filtered, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFilterEnabled(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNode.Bind_set_filter_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsFilterEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNode.Bind_is_filter_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Blend an animation by [param blend] amount (name must be valid in the linked [AnimationPlayer]). A [param time] and [param delta] may be passed, as well as whether [param seeked] happened.
A [param looped_flag] is used by internal processing immediately after the loop. See also [enum Animation.LoopedFlag].
*/
//go:nosplit
func (self class) BlendAnimation(animation gd.StringName, time gd.Float, delta gd.Float, seeked bool, is_external_seeking bool, blend gd.Float, looped_flag classdb.AnimationLoopedFlag) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(animation))
	callframe.Arg(frame, time)
	callframe.Arg(frame, delta)
	callframe.Arg(frame, seeked)
	callframe.Arg(frame, is_external_seeking)
	callframe.Arg(frame, blend)
	callframe.Arg(frame, looped_flag)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNode.Bind_blend_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Blend another animation node (in case this animation node contains child animation nodes). This function is only useful if you inherit from [AnimationRootNode] instead, otherwise editors will not display your animation node for addition.
*/
//go:nosplit
func (self class) BlendNode(name gd.StringName, node objects.AnimationNode, time gd.Float, seek bool, is_external_seeking bool, blend gd.Float, filter classdb.AnimationNodeFilterAction, sync bool, test_only bool) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(node[0])[0])
	callframe.Arg(frame, time)
	callframe.Arg(frame, seek)
	callframe.Arg(frame, is_external_seeking)
	callframe.Arg(frame, blend)
	callframe.Arg(frame, filter)
	callframe.Arg(frame, sync)
	callframe.Arg(frame, test_only)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNode.Bind_blend_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Blend an input. This is only useful for animation nodes created for an [AnimationNodeBlendTree]. The [param time] parameter is a relative delta, unless [param seek] is [code]true[/code], in which case it is absolute. A filter mode may be optionally passed (see [enum FilterAction] for options).
*/
//go:nosplit
func (self class) BlendInput(input_index gd.Int, time gd.Float, seek bool, is_external_seeking bool, blend gd.Float, filter classdb.AnimationNodeFilterAction, sync bool, test_only bool) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, input_index)
	callframe.Arg(frame, time)
	callframe.Arg(frame, seek)
	callframe.Arg(frame, is_external_seeking)
	callframe.Arg(frame, blend)
	callframe.Arg(frame, filter)
	callframe.Arg(frame, sync)
	callframe.Arg(frame, test_only)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNode.Bind_blend_input, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets a custom parameter. These are used as local memory, because resources can be reused across the tree or scenes.
*/
//go:nosplit
func (self class) SetParameter(name gd.StringName, value gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(value))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNode.Bind_set_parameter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Gets the value of a parameter. Parameters are custom local memory used for your animation nodes, given a resource can be reused in multiple trees.
*/
//go:nosplit
func (self class) GetParameter(name gd.StringName) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNode.Bind_get_parameter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}
func (self Instance) OnTreeChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("tree_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnAnimationNodeRenamed(cb func(object_id int, old_name string, new_name string)) {
	self[0].AsObject().Connect(gd.NewStringName("animation_node_renamed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnAnimationNodeRemoved(cb func(object_id int, name string)) {
	self[0].AsObject().Connect(gd.NewStringName("animation_node_removed"), gd.NewCallable(cb), 0)
}

func (self class) AsAnimationNode() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAnimationNode() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	case "_get_child_nodes":
		return reflect.ValueOf(self._get_child_nodes)
	case "_get_parameter_list":
		return reflect.ValueOf(self._get_parameter_list)
	case "_get_child_by_name":
		return reflect.ValueOf(self._get_child_by_name)
	case "_get_parameter_default_value":
		return reflect.ValueOf(self._get_parameter_default_value)
	case "_is_parameter_read_only":
		return reflect.ValueOf(self._is_parameter_read_only)
	case "_process":
		return reflect.ValueOf(self._process)
	case "_get_caption":
		return reflect.ValueOf(self._get_caption)
	case "_has_filter":
		return reflect.ValueOf(self._has_filter)
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_child_nodes":
		return reflect.ValueOf(self._get_child_nodes)
	case "_get_parameter_list":
		return reflect.ValueOf(self._get_parameter_list)
	case "_get_child_by_name":
		return reflect.ValueOf(self._get_child_by_name)
	case "_get_parameter_default_value":
		return reflect.ValueOf(self._get_parameter_default_value)
	case "_is_parameter_read_only":
		return reflect.ValueOf(self._is_parameter_read_only)
	case "_process":
		return reflect.ValueOf(self._process)
	case "_get_caption":
		return reflect.ValueOf(self._get_caption)
	case "_has_filter":
		return reflect.ValueOf(self._has_filter)
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {
	classdb.Register("AnimationNode", func(ptr gd.Object) any {
		return [1]classdb.AnimationNode{*(*classdb.AnimationNode)(unsafe.Pointer(&ptr))}
	})
}

type FilterAction = classdb.AnimationNodeFilterAction

const (
	/*Do not use filtering.*/
	FilterIgnore FilterAction = 0
	/*Paths matching the filter will be allowed to pass.*/
	FilterPass FilterAction = 1
	/*Paths matching the filter will be discarded.*/
	FilterStop FilterAction = 2
	/*Paths matching the filter will be blended (by the blend value).*/
	FilterBlend FilterAction = 3
)
