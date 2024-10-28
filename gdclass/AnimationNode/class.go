package AnimationNode

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

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
		GetChildNodes() gd.Dictionary
		//When inheriting from [AnimationRootNode], implement this virtual method to return a list of the properties on this animation node. Parameters are custom local memory used for your animation nodes, given a resource can be reused in multiple trees. Format is similar to [method Object.get_property_list].
		GetParameterList() gd.Array
		//When inheriting from [AnimationRootNode], implement this virtual method to return a child animation node by its [param name].
		GetChildByName(name string) gdclass.AnimationNode
		//When inheriting from [AnimationRootNode], implement this virtual method to return the default value of a [param parameter]. Parameters are custom local memory used for your animation nodes, given a resource can be reused in multiple trees.
		GetParameterDefaultValue(parameter string) gd.Variant
		//When inheriting from [AnimationRootNode], implement this virtual method to return whether the [param parameter] is read-only. Parameters are custom local memory used for your animation nodes, given a resource can be reused in multiple trees.
		IsParameterReadOnly(parameter string) bool
		//When inheriting from [AnimationRootNode], implement this virtual method to run some code when this animation node is processed. The [param time] parameter is a relative delta, unless [param seek] is [code]true[/code], in which case it is absolute.
		//Here, call the [method blend_input], [method blend_node] or [method blend_animation] functions. You can also use [method get_parameter] and [method set_parameter] to modify local memory.
		//This function should return the delta.
		Process(time float64, seek bool, is_external_seeking bool, test_only bool) float64
		//When inheriting from [AnimationRootNode], implement this virtual method to override the text caption for this animation node.
		GetCaption() string
		//When inheriting from [AnimationRootNode], implement this virtual method to return whether the blend tree editor should display filter editing on this animation node.
		HasFilter() bool
	}

*/
type Go [1]classdb.AnimationNode

/*
When inheriting from [AnimationRootNode], implement this virtual method to return all child animation nodes in order as a [code]name: node[/code] dictionary.
*/
func (Go) _get_child_nodes(impl func(ptr unsafe.Pointer) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
ptr, ok := discreet.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to return a list of the properties on this animation node. Parameters are custom local memory used for your animation nodes, given a resource can be reused in multiple trees. Format is similar to [method Object.get_property_list].
*/
func (Go) _get_parameter_list(impl func(ptr unsafe.Pointer) gd.Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
ptr, ok := discreet.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to return a child animation node by its [param name].
*/
func (Go) _get_child_by_name(impl func(ptr unsafe.Pointer, name string) gdclass.AnimationNode, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var name = discreet.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args,0))
		defer discreet.End(name)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, name.String())
ptr, ok := discreet.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to return the default value of a [param parameter]. Parameters are custom local memory used for your animation nodes, given a resource can be reused in multiple trees.
*/
func (Go) _get_parameter_default_value(impl func(ptr unsafe.Pointer, parameter string) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var parameter = discreet.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args,0))
		defer discreet.End(parameter)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, parameter.String())
ptr, ok := discreet.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to return whether the [param parameter] is read-only. Parameters are custom local memory used for your animation nodes, given a resource can be reused in multiple trees.
*/
func (Go) _is_parameter_read_only(impl func(ptr unsafe.Pointer, parameter string) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var parameter = discreet.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args,0))
		defer discreet.End(parameter)
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
func (Go) _process(impl func(ptr unsafe.Pointer, time float64, seek bool, is_external_seeking bool, test_only bool) float64, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var time = gd.UnsafeGet[gd.Float](p_args,0)
		var seek = gd.UnsafeGet[bool](p_args,1)
		var is_external_seeking = gd.UnsafeGet[bool](p_args,2)
		var test_only = gd.UnsafeGet[bool](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, float64(time), seek, is_external_seeking, test_only)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to override the text caption for this animation node.
*/
func (Go) _get_caption(impl func(ptr unsafe.Pointer) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
ptr, ok := discreet.End(gd.NewString(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to return whether the blend tree editor should display filter editing on this animation node.
*/
func (Go) _has_filter(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Adds an input to the animation node. This is only useful for animation nodes created for use in an [AnimationNodeBlendTree]. If the addition fails, returns [code]false[/code].
*/
func (self Go) AddInput(name string) bool {
	return bool(class(self).AddInput(gd.NewString(name)))
}

/*
Removes an input, call this only when inactive.
*/
func (self Go) RemoveInput(index int) {
	class(self).RemoveInput(gd.Int(index))
}

/*
Sets the name of the input at the given [param input] index. If the setting fails, returns [code]false[/code].
*/
func (self Go) SetInputName(input int, name string) bool {
	return bool(class(self).SetInputName(gd.Int(input), gd.NewString(name)))
}

/*
Gets the name of an input by index.
*/
func (self Go) GetInputName(input int) string {
	return string(class(self).GetInputName(gd.Int(input)).String())
}

/*
Amount of inputs in this animation node, only useful for animation nodes that go into [AnimationNodeBlendTree].
*/
func (self Go) GetInputCount() int {
	return int(int(class(self).GetInputCount()))
}

/*
Returns the input index which corresponds to [param name]. If not found, returns [code]-1[/code].
*/
func (self Go) FindInput(name string) int {
	return int(int(class(self).FindInput(gd.NewString(name))))
}

/*
Adds or removes a path for the filter.
*/
func (self Go) SetFilterPath(path string, enable bool) {
	class(self).SetFilterPath(gd.NewString(path).NodePath(), enable)
}

/*
Returns whether the given path is filtered.
*/
func (self Go) IsPathFiltered(path string) bool {
	return bool(class(self).IsPathFiltered(gd.NewString(path).NodePath()))
}

/*
Blend an animation by [param blend] amount (name must be valid in the linked [AnimationPlayer]). A [param time] and [param delta] may be passed, as well as whether [param seeked] happened.
A [param looped_flag] is used by internal processing immediately after the loop. See also [enum Animation.LoopedFlag].
*/
func (self Go) BlendAnimation(animation string, time float64, delta float64, seeked bool, is_external_seeking bool, blend float64) {
	class(self).BlendAnimation(gd.NewStringName(animation), gd.Float(time), gd.Float(delta), seeked, is_external_seeking, gd.Float(blend), 0)
}

/*
Blend another animation node (in case this animation node contains child animation nodes). This function is only useful if you inherit from [AnimationRootNode] instead, otherwise editors will not display your animation node for addition.
*/
func (self Go) BlendNode(name string, node gdclass.AnimationNode, time float64, seek bool, is_external_seeking bool, blend float64) float64 {
	return float64(float64(class(self).BlendNode(gd.NewStringName(name), node, gd.Float(time), seek, is_external_seeking, gd.Float(blend), 0, true, false)))
}

/*
Blend an input. This is only useful for animation nodes created for an [AnimationNodeBlendTree]. The [param time] parameter is a relative delta, unless [param seek] is [code]true[/code], in which case it is absolute. A filter mode may be optionally passed (see [enum FilterAction] for options).
*/
func (self Go) BlendInput(input_index int, time float64, seek bool, is_external_seeking bool, blend float64) float64 {
	return float64(float64(class(self).BlendInput(gd.Int(input_index), gd.Float(time), seek, is_external_seeking, gd.Float(blend), 0, true, false)))
}

/*
Sets a custom parameter. These are used as local memory, because resources can be reused across the tree or scenes.
*/
func (self Go) SetParameter(name string, value gd.Variant) {
	class(self).SetParameter(gd.NewStringName(name), value)
}

/*
Gets the value of a parameter. Parameters are custom local memory used for your animation nodes, given a resource can be reused in multiple trees.
*/
func (self Go) GetParameter(name string) gd.Variant {
	return gd.Variant(class(self).GetParameter(gd.NewStringName(name)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AnimationNode
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimationNode"))
	return Go{classdb.AnimationNode(object)}
}

func (self Go) FilterEnabled() bool {
		return bool(class(self).IsFilterEnabled())
}

func (self Go) SetFilterEnabled(value bool) {
	class(self).SetFilterEnabled(value)
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to return all child animation nodes in order as a [code]name: node[/code] dictionary.
*/
func (class) _get_child_nodes(impl func(ptr unsafe.Pointer) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
ptr, ok := discreet.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to return a list of the properties on this animation node. Parameters are custom local memory used for your animation nodes, given a resource can be reused in multiple trees. Format is similar to [method Object.get_property_list].
*/
func (class) _get_parameter_list(impl func(ptr unsafe.Pointer) gd.Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
ptr, ok := discreet.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to return a child animation node by its [param name].
*/
func (class) _get_child_by_name(impl func(ptr unsafe.Pointer, name gd.StringName) gdclass.AnimationNode, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var name = discreet.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, name)
ptr, ok := discreet.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to return the default value of a [param parameter]. Parameters are custom local memory used for your animation nodes, given a resource can be reused in multiple trees.
*/
func (class) _get_parameter_default_value(impl func(ptr unsafe.Pointer, parameter gd.StringName) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var parameter = discreet.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, parameter)
ptr, ok := discreet.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to return whether the [param parameter] is read-only. Parameters are custom local memory used for your animation nodes, given a resource can be reused in multiple trees.
*/
func (class) _is_parameter_read_only(impl func(ptr unsafe.Pointer, parameter gd.StringName) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var parameter = discreet.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args,0))
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
func (class) _process(impl func(ptr unsafe.Pointer, time gd.Float, seek bool, is_external_seeking bool, test_only bool) gd.Float, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var time = gd.UnsafeGet[gd.Float](p_args,0)
		var seek = gd.UnsafeGet[bool](p_args,1)
		var is_external_seeking = gd.UnsafeGet[bool](p_args,2)
		var test_only = gd.UnsafeGet[bool](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, time, seek, is_external_seeking, test_only)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to override the text caption for this animation node.
*/
func (class) _get_caption(impl func(ptr unsafe.Pointer) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
ptr, ok := discreet.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
When inheriting from [AnimationRootNode], implement this virtual method to return whether the blend tree editor should display filter editing on this animation node.
*/
func (class) _has_filter(impl func(ptr unsafe.Pointer) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
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
	callframe.Arg(frame, discreet.Get(name))
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
func (self class) RemoveInput(index gd.Int)  {
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
	callframe.Arg(frame, discreet.Get(name))
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
	var ret = discreet.New[gd.String](r_ret.Get())
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
	callframe.Arg(frame, discreet.Get(name))
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
func (self class) SetFilterPath(path gd.NodePath, enable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(path))
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
	callframe.Arg(frame, discreet.Get(path))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNode.Bind_is_path_filtered, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFilterEnabled(enable bool)  {
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
func (self class) BlendAnimation(animation gd.StringName, time gd.Float, delta gd.Float, seeked bool, is_external_seeking bool, blend gd.Float, looped_flag classdb.AnimationLoopedFlag)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(animation))
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
func (self class) BlendNode(name gd.StringName, node gdclass.AnimationNode, time gd.Float, seek bool, is_external_seeking bool, blend gd.Float, filter classdb.AnimationNodeFilterAction, sync bool, test_only bool) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
	callframe.Arg(frame, discreet.Get(node[0])[0])
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
func (self class) SetParameter(name gd.StringName, value gd.Variant)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
	callframe.Arg(frame, discreet.Get(value))
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
	callframe.Arg(frame, discreet.Get(name))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNode.Bind_get_parameter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}
func (self Go) OnTreeChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("tree_changed"), gd.NewCallable(cb), 0)
}


func (self Go) OnAnimationNodeRenamed(cb func(object_id int, old_name string, new_name string)) {
	self[0].AsObject().Connect(gd.NewStringName("animation_node_renamed"), gd.NewCallable(cb), 0)
}


func (self Go) OnAnimationNodeRemoved(cb func(object_id int, name string)) {
	self[0].AsObject().Connect(gd.NewStringName("animation_node_removed"), gd.NewCallable(cb), 0)
}


func (self class) AsAnimationNode() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAnimationNode() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_child_nodes": return reflect.ValueOf(self._get_child_nodes);
	case "_get_parameter_list": return reflect.ValueOf(self._get_parameter_list);
	case "_get_child_by_name": return reflect.ValueOf(self._get_child_by_name);
	case "_get_parameter_default_value": return reflect.ValueOf(self._get_parameter_default_value);
	case "_is_parameter_read_only": return reflect.ValueOf(self._is_parameter_read_only);
	case "_process": return reflect.ValueOf(self._process);
	case "_get_caption": return reflect.ValueOf(self._get_caption);
	case "_has_filter": return reflect.ValueOf(self._has_filter);
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_get_child_nodes": return reflect.ValueOf(self._get_child_nodes);
	case "_get_parameter_list": return reflect.ValueOf(self._get_parameter_list);
	case "_get_child_by_name": return reflect.ValueOf(self._get_child_by_name);
	case "_get_parameter_default_value": return reflect.ValueOf(self._get_parameter_default_value);
	case "_is_parameter_read_only": return reflect.ValueOf(self._is_parameter_read_only);
	case "_process": return reflect.ValueOf(self._process);
	case "_get_caption": return reflect.ValueOf(self._get_caption);
	case "_has_filter": return reflect.ValueOf(self._has_filter);
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("AnimationNode", func(ptr gd.Object) any { return classdb.AnimationNode(ptr) })}
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
