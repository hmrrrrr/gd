// Package SceneState provides methods for working with SceneState object instances.
package SceneState

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

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
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
Maintains a list of resources, nodes, exported and overridden properties, and built-in scripts associated with a scene. They cannot be modified from a [SceneState], only accessed. Useful for peeking into what a [PackedScene] contains without instantiating it.
This class cannot be instantiated directly, it is retrieved for a given scene as the result of [method PackedScene.get_state].
*/
type Instance [1]gdclass.SceneState

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsSceneState() Instance
}

/*
Returns the number of nodes in the scene.
The [code]idx[/code] argument used to query node data in other [code]get_node_*[/code] methods in the interval [code][0, get_node_count() - 1][/code].
*/
func (self Instance) GetNodeCount() int { //gd:SceneState.get_node_count
	return int(int(class(self).GetNodeCount()))
}

/*
Returns the type of the node at [param idx].
*/
func (self Instance) GetNodeType(idx int) string { //gd:SceneState.get_node_type
	return string(class(self).GetNodeType(int64(idx)).String())
}

/*
Returns the name of the node at [param idx].
*/
func (self Instance) GetNodeName(idx int) string { //gd:SceneState.get_node_name
	return string(class(self).GetNodeName(int64(idx)).String())
}

/*
Returns the path to the node at [param idx].
If [param for_parent] is [code]true[/code], returns the path of the [param idx] node's parent instead.
*/
func (self Instance) GetNodePath(idx int) string { //gd:SceneState.get_node_path
	return string(class(self).GetNodePath(int64(idx), false).String())
}

/*
Returns the path to the owner of the node at [param idx], relative to the root node.
*/
func (self Instance) GetNodeOwnerPath(idx int) string { //gd:SceneState.get_node_owner_path
	return string(class(self).GetNodeOwnerPath(int64(idx)).String())
}

/*
Returns [code]true[/code] if the node at [param idx] is an [InstancePlaceholder].
*/
func (self Instance) IsNodeInstancePlaceholder(idx int) bool { //gd:SceneState.is_node_instance_placeholder
	return bool(class(self).IsNodeInstancePlaceholder(int64(idx)))
}

/*
Returns the path to the represented scene file if the node at [param idx] is an [InstancePlaceholder].
*/
func (self Instance) GetNodeInstancePlaceholder(idx int) string { //gd:SceneState.get_node_instance_placeholder
	return string(class(self).GetNodeInstancePlaceholder(int64(idx)).String())
}

/*
Returns a [PackedScene] for the node at [param idx] (i.e. the whole branch starting at this node, with its child nodes and resources), or [code]null[/code] if the node is not an instance.
*/
func (self Instance) GetNodeInstance(idx int) [1]gdclass.PackedScene { //gd:SceneState.get_node_instance
	return [1]gdclass.PackedScene(class(self).GetNodeInstance(int64(idx)))
}

/*
Returns the list of group names associated with the node at [param idx].
*/
func (self Instance) GetNodeGroups(idx int) []string { //gd:SceneState.get_node_groups
	return []string(class(self).GetNodeGroups(int64(idx)).Strings())
}

/*
Returns the node's index, which is its position relative to its siblings. This is only relevant and saved in scenes for cases where new nodes are added to an instantiated or inherited scene among siblings from the base scene. Despite the name, this index is not related to the [param idx] argument used here and in other methods.
*/
func (self Instance) GetNodeIndex(idx int) int { //gd:SceneState.get_node_index
	return int(int(class(self).GetNodeIndex(int64(idx))))
}

/*
Returns the number of exported or overridden properties for the node at [param idx].
The [code]prop_idx[/code] argument used to query node property data in other [code]get_node_property_*[/code] methods in the interval [code][0, get_node_property_count() - 1][/code].
*/
func (self Instance) GetNodePropertyCount(idx int) int { //gd:SceneState.get_node_property_count
	return int(int(class(self).GetNodePropertyCount(int64(idx))))
}

/*
Returns the name of the property at [param prop_idx] for the node at [param idx].
*/
func (self Instance) GetNodePropertyName(idx int, prop_idx int) string { //gd:SceneState.get_node_property_name
	return string(class(self).GetNodePropertyName(int64(idx), int64(prop_idx)).String())
}

/*
Returns the value of the property at [param prop_idx] for the node at [param idx].
*/
func (self Instance) GetNodePropertyValue(idx int, prop_idx int) any { //gd:SceneState.get_node_property_value
	return any(class(self).GetNodePropertyValue(int64(idx), int64(prop_idx)).Interface())
}

/*
Returns the number of signal connections in the scene.
The [code]idx[/code] argument used to query connection metadata in other [code]get_connection_*[/code] methods in the interval [code][0, get_connection_count() - 1][/code].
*/
func (self Instance) GetConnectionCount() int { //gd:SceneState.get_connection_count
	return int(int(class(self).GetConnectionCount()))
}

/*
Returns the path to the node that owns the signal at [param idx], relative to the root node.
*/
func (self Instance) GetConnectionSource(idx int) string { //gd:SceneState.get_connection_source
	return string(class(self).GetConnectionSource(int64(idx)).String())
}

/*
Returns the name of the signal at [param idx].
*/
func (self Instance) GetConnectionSignal(idx int) string { //gd:SceneState.get_connection_signal
	return string(class(self).GetConnectionSignal(int64(idx)).String())
}

/*
Returns the path to the node that owns the method connected to the signal at [param idx], relative to the root node.
*/
func (self Instance) GetConnectionTarget(idx int) string { //gd:SceneState.get_connection_target
	return string(class(self).GetConnectionTarget(int64(idx)).String())
}

/*
Returns the method connected to the signal at [param idx].
*/
func (self Instance) GetConnectionMethod(idx int) string { //gd:SceneState.get_connection_method
	return string(class(self).GetConnectionMethod(int64(idx)).String())
}

/*
Returns the connection flags for the signal at [param idx]. See [enum Object.ConnectFlags] constants.
*/
func (self Instance) GetConnectionFlags(idx int) int { //gd:SceneState.get_connection_flags
	return int(int(class(self).GetConnectionFlags(int64(idx))))
}

/*
Returns the list of bound parameters for the signal at [param idx].
*/
func (self Instance) GetConnectionBinds(idx int) []any { //gd:SceneState.get_connection_binds
	return []any(gd.ArrayAs[[]any](gd.InternalArray(class(self).GetConnectionBinds(int64(idx)))))
}

/*
Returns the number of unbound parameters for the signal at [param idx].
*/
func (self Instance) GetConnectionUnbinds(idx int) int { //gd:SceneState.get_connection_unbinds
	return int(int(class(self).GetConnectionUnbinds(int64(idx))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.SceneState

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SceneState"))
	casted := Instance{*(*gdclass.SceneState)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Returns the number of nodes in the scene.
The [code]idx[/code] argument used to query node data in other [code]get_node_*[/code] methods in the interval [code][0, get_node_count() - 1][/code].
*/
//go:nosplit
func (self class) GetNodeCount() int64 { //gd:SceneState.get_node_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneState.Bind_get_node_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the type of the node at [param idx].
*/
//go:nosplit
func (self class) GetNodeType(idx int64) String.Name { //gd:SceneState.get_node_type
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneState.Bind_get_node_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the name of the node at [param idx].
*/
//go:nosplit
func (self class) GetNodeName(idx int64) String.Name { //gd:SceneState.get_node_name
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneState.Bind_get_node_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the path to the node at [param idx].
If [param for_parent] is [code]true[/code], returns the path of the [param idx] node's parent instead.
*/
//go:nosplit
func (self class) GetNodePath(idx int64, for_parent bool) Path.ToNode { //gd:SceneState.get_node_path
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, for_parent)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneState.Bind_get_node_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Path.ToNode(String.Via(gd.NodePathProxy{}, pointers.Pack(pointers.New[gd.NodePath](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the path to the owner of the node at [param idx], relative to the root node.
*/
//go:nosplit
func (self class) GetNodeOwnerPath(idx int64) Path.ToNode { //gd:SceneState.get_node_owner_path
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneState.Bind_get_node_owner_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Path.ToNode(String.Via(gd.NodePathProxy{}, pointers.Pack(pointers.New[gd.NodePath](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the node at [param idx] is an [InstancePlaceholder].
*/
//go:nosplit
func (self class) IsNodeInstancePlaceholder(idx int64) bool { //gd:SceneState.is_node_instance_placeholder
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneState.Bind_is_node_instance_placeholder, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the path to the represented scene file if the node at [param idx] is an [InstancePlaceholder].
*/
//go:nosplit
func (self class) GetNodeInstancePlaceholder(idx int64) String.Readable { //gd:SceneState.get_node_instance_placeholder
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneState.Bind_get_node_instance_placeholder, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns a [PackedScene] for the node at [param idx] (i.e. the whole branch starting at this node, with its child nodes and resources), or [code]null[/code] if the node is not an instance.
*/
//go:nosplit
func (self class) GetNodeInstance(idx int64) [1]gdclass.PackedScene { //gd:SceneState.get_node_instance
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneState.Bind_get_node_instance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.PackedScene{gd.PointerWithOwnershipTransferredToGo[gdclass.PackedScene](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the list of group names associated with the node at [param idx].
*/
//go:nosplit
func (self class) GetNodeGroups(idx int64) Packed.Strings { //gd:SceneState.get_node_groups
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneState.Bind_get_node_groups, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the node's index, which is its position relative to its siblings. This is only relevant and saved in scenes for cases where new nodes are added to an instantiated or inherited scene among siblings from the base scene. Despite the name, this index is not related to the [param idx] argument used here and in other methods.
*/
//go:nosplit
func (self class) GetNodeIndex(idx int64) int64 { //gd:SceneState.get_node_index
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneState.Bind_get_node_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of exported or overridden properties for the node at [param idx].
The [code]prop_idx[/code] argument used to query node property data in other [code]get_node_property_*[/code] methods in the interval [code][0, get_node_property_count() - 1][/code].
*/
//go:nosplit
func (self class) GetNodePropertyCount(idx int64) int64 { //gd:SceneState.get_node_property_count
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneState.Bind_get_node_property_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the name of the property at [param prop_idx] for the node at [param idx].
*/
//go:nosplit
func (self class) GetNodePropertyName(idx int64, prop_idx int64) String.Name { //gd:SceneState.get_node_property_name
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, prop_idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneState.Bind_get_node_property_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the value of the property at [param prop_idx] for the node at [param idx].
*/
//go:nosplit
func (self class) GetNodePropertyValue(idx int64, prop_idx int64) variant.Any { //gd:SceneState.get_node_property_value
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, prop_idx)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneState.Bind_get_node_property_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Implementation(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the number of signal connections in the scene.
The [code]idx[/code] argument used to query connection metadata in other [code]get_connection_*[/code] methods in the interval [code][0, get_connection_count() - 1][/code].
*/
//go:nosplit
func (self class) GetConnectionCount() int64 { //gd:SceneState.get_connection_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneState.Bind_get_connection_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the path to the node that owns the signal at [param idx], relative to the root node.
*/
//go:nosplit
func (self class) GetConnectionSource(idx int64) Path.ToNode { //gd:SceneState.get_connection_source
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneState.Bind_get_connection_source, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Path.ToNode(String.Via(gd.NodePathProxy{}, pointers.Pack(pointers.New[gd.NodePath](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the name of the signal at [param idx].
*/
//go:nosplit
func (self class) GetConnectionSignal(idx int64) String.Name { //gd:SceneState.get_connection_signal
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneState.Bind_get_connection_signal, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the path to the node that owns the method connected to the signal at [param idx], relative to the root node.
*/
//go:nosplit
func (self class) GetConnectionTarget(idx int64) Path.ToNode { //gd:SceneState.get_connection_target
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneState.Bind_get_connection_target, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Path.ToNode(String.Via(gd.NodePathProxy{}, pointers.Pack(pointers.New[gd.NodePath](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the method connected to the signal at [param idx].
*/
//go:nosplit
func (self class) GetConnectionMethod(idx int64) String.Name { //gd:SceneState.get_connection_method
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneState.Bind_get_connection_method, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the connection flags for the signal at [param idx]. See [enum Object.ConnectFlags] constants.
*/
//go:nosplit
func (self class) GetConnectionFlags(idx int64) int64 { //gd:SceneState.get_connection_flags
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneState.Bind_get_connection_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the list of bound parameters for the signal at [param idx].
*/
//go:nosplit
func (self class) GetConnectionBinds(idx int64) Array.Any { //gd:SceneState.get_connection_binds
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneState.Bind_get_connection_binds, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the number of unbound parameters for the signal at [param idx].
*/
//go:nosplit
func (self class) GetConnectionUnbinds(idx int64) int64 { //gd:SceneState.get_connection_unbinds
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SceneState.Bind_get_connection_unbinds, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsSceneState() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSceneState() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("SceneState", func(ptr gd.Object) any { return [1]gdclass.SceneState{*(*gdclass.SceneState)(unsafe.Pointer(&ptr))} })
}

type GenEditState = gdclass.SceneStateGenEditState //gd:SceneState.GenEditState

const (
	/*If passed to [method PackedScene.instantiate], blocks edits to the scene state.*/
	GenEditStateDisabled GenEditState = 0
	/*If passed to [method PackedScene.instantiate], provides inherited scene resources to the local scene.
	  [b]Note:[/b] Only available in editor builds.*/
	GenEditStateInstance GenEditState = 1
	/*If passed to [method PackedScene.instantiate], provides local scene resources to the local scene. Only the main scene should receive the main edit state.
	  [b]Note:[/b] Only available in editor builds.*/
	GenEditStateMain GenEditState = 2
	/*If passed to [method PackedScene.instantiate], it's similar to [constant GEN_EDIT_STATE_MAIN], but for the case where the scene is being instantiated to be the base of another one.
	  [b]Note:[/b] Only available in editor builds.*/
	GenEditStateMainInherited GenEditState = 3
)
