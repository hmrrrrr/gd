// Package AnimationNodeBlendTree provides methods for working with AnimationNodeBlendTree object instances.
package AnimationNodeBlendTree

import "unsafe"
import "reflect"
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
import "graphics.gd/classdb/AnimationRootNode"
import "graphics.gd/classdb/AnimationNode"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Vector2"

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

/*
This animation node may contain a sub-tree of any other type animation nodes, such as [AnimationNodeTransition], [AnimationNodeBlend2], [AnimationNodeBlend3], [AnimationNodeOneShot], etc. This is one of the most commonly used animation node roots.
An [AnimationNodeOutput] node named [code]output[/code] is created by default.
*/
type Instance [1]gdclass.AnimationNodeBlendTree

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAnimationNodeBlendTree() Instance
}

/*
Adds an [AnimationNode] at the given [param position]. The [param name] is used to identify the created sub animation node later.
*/
func (self Instance) AddNode(name string, node [1]gdclass.AnimationNode) { //gd:AnimationNodeBlendTree.add_node
	class(self).AddNode(String.Name(String.New(name)), node, gd.Vector2(gd.Vector2{0, 0}))
}

/*
Returns the sub animation node with the specified [param name].
*/
func (self Instance) GetNode(name string) [1]gdclass.AnimationNode { //gd:AnimationNodeBlendTree.get_node
	return [1]gdclass.AnimationNode(class(self).GetNode(String.Name(String.New(name))))
}

/*
Removes a sub animation node.
*/
func (self Instance) RemoveNode(name string) { //gd:AnimationNodeBlendTree.remove_node
	class(self).RemoveNode(String.Name(String.New(name)))
}

/*
Changes the name of a sub animation node.
*/
func (self Instance) RenameNode(name string, new_name string) { //gd:AnimationNodeBlendTree.rename_node
	class(self).RenameNode(String.Name(String.New(name)), String.Name(String.New(new_name)))
}

/*
Returns [code]true[/code] if a sub animation node with specified [param name] exists.
*/
func (self Instance) HasNode(name string) bool { //gd:AnimationNodeBlendTree.has_node
	return bool(class(self).HasNode(String.Name(String.New(name))))
}

/*
Connects the output of an [AnimationNode] as input for another [AnimationNode], at the input port specified by [param input_index].
*/
func (self Instance) ConnectNode(input_node string, input_index int, output_node string) { //gd:AnimationNodeBlendTree.connect_node
	class(self).ConnectNode(String.Name(String.New(input_node)), gd.Int(input_index), String.Name(String.New(output_node)))
}

/*
Disconnects the animation node connected to the specified input.
*/
func (self Instance) DisconnectNode(input_node string, input_index int) { //gd:AnimationNodeBlendTree.disconnect_node
	class(self).DisconnectNode(String.Name(String.New(input_node)), gd.Int(input_index))
}

/*
Modifies the position of a sub animation node.
*/
func (self Instance) SetNodePosition(name string, position Vector2.XY) { //gd:AnimationNodeBlendTree.set_node_position
	class(self).SetNodePosition(String.Name(String.New(name)), gd.Vector2(position))
}

/*
Returns the position of the sub animation node with the specified [param name].
*/
func (self Instance) GetNodePosition(name string) Vector2.XY { //gd:AnimationNodeBlendTree.get_node_position
	return Vector2.XY(class(self).GetNodePosition(String.Name(String.New(name))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AnimationNodeBlendTree

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimationNodeBlendTree"))
	casted := Instance{*(*gdclass.AnimationNodeBlendTree)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) GraphOffset() Vector2.XY {
	return Vector2.XY(class(self).GetGraphOffset())
}

func (self Instance) SetGraphOffset(value Vector2.XY) {
	class(self).SetGraphOffset(gd.Vector2(value))
}

/*
Adds an [AnimationNode] at the given [param position]. The [param name] is used to identify the created sub animation node later.
*/
//go:nosplit
func (self class) AddNode(name String.Name, node [1]gdclass.AnimationNode, position gd.Vector2) { //gd:AnimationNodeBlendTree.add_node
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(node[0])[0])
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_add_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the sub animation node with the specified [param name].
*/
//go:nosplit
func (self class) GetNode(name String.Name) [1]gdclass.AnimationNode { //gd:AnimationNodeBlendTree.get_node
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_get_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.AnimationNode{gd.PointerWithOwnershipTransferredToGo[gdclass.AnimationNode](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Removes a sub animation node.
*/
//go:nosplit
func (self class) RemoveNode(name String.Name) { //gd:AnimationNodeBlendTree.remove_node
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_remove_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Changes the name of a sub animation node.
*/
//go:nosplit
func (self class) RenameNode(name String.Name, new_name String.Name) { //gd:AnimationNodeBlendTree.rename_node
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(new_name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_rename_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if a sub animation node with specified [param name] exists.
*/
//go:nosplit
func (self class) HasNode(name String.Name) bool { //gd:AnimationNodeBlendTree.has_node
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_has_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Connects the output of an [AnimationNode] as input for another [AnimationNode], at the input port specified by [param input_index].
*/
//go:nosplit
func (self class) ConnectNode(input_node String.Name, input_index gd.Int, output_node String.Name) { //gd:AnimationNodeBlendTree.connect_node
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(input_node)))
	callframe.Arg(frame, input_index)
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(output_node)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_connect_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Disconnects the animation node connected to the specified input.
*/
//go:nosplit
func (self class) DisconnectNode(input_node String.Name, input_index gd.Int) { //gd:AnimationNodeBlendTree.disconnect_node
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(input_node)))
	callframe.Arg(frame, input_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_disconnect_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Modifies the position of a sub animation node.
*/
//go:nosplit
func (self class) SetNodePosition(name String.Name, position gd.Vector2) { //gd:AnimationNodeBlendTree.set_node_position
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_set_node_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the position of the sub animation node with the specified [param name].
*/
//go:nosplit
func (self class) GetNodePosition(name String.Name) gd.Vector2 { //gd:AnimationNodeBlendTree.get_node_position
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_get_node_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGraphOffset(offset gd.Vector2) { //gd:AnimationNodeBlendTree.set_graph_offset
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_set_graph_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGraphOffset() gd.Vector2 { //gd:AnimationNodeBlendTree.get_graph_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeBlendTree.Bind_get_graph_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnNodeChanged(cb func(node_name string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("node_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsAnimationNodeBlendTree() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAnimationNodeBlendTree() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAnimationRootNode() AnimationRootNode.Advanced {
	return *((*AnimationRootNode.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAnimationRootNode() AnimationRootNode.Instance {
	return *((*AnimationRootNode.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAnimationNode() AnimationNode.Advanced {
	return *((*AnimationNode.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAnimationNode() AnimationNode.Instance {
	return *((*AnimationNode.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(AnimationRootNode.Advanced(self.AsAnimationRootNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AnimationRootNode.Instance(self.AsAnimationRootNode()), name)
	}
}
func init() {
	gdclass.Register("AnimationNodeBlendTree", func(ptr gd.Object) any {
		return [1]gdclass.AnimationNodeBlendTree{*(*gdclass.AnimationNodeBlendTree)(unsafe.Pointer(&ptr))}
	})
}
