package VisualShader

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Shader"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
This class provides a graph-like visual editor for creating a [Shader]. Although [VisualShader]s do not require coding, they share the same logic with script shaders. They use [VisualShaderNode]s that can be connected to each other to control the flow of the shader. The visual shader graph is converted to a script shader behind the scenes.
*/
type Instance [1]classdb.VisualShader

/*
Sets the mode of this shader.
*/
func (self Instance) SetMode(mode classdb.ShaderMode) {
	class(self).SetMode(mode)
}

/*
Adds the specified [param node] to the shader.
*/
func (self Instance) AddNode(atype classdb.VisualShaderType, node gdclass.VisualShaderNode, position gd.Vector2, id int) {
	class(self).AddNode(atype, node, position, gd.Int(id))
}

/*
Returns the shader node instance with specified [param type] and [param id].
*/
func (self Instance) GetNode(atype classdb.VisualShaderType, id int) gdclass.VisualShaderNode {
	return gdclass.VisualShaderNode(class(self).GetNode(atype, gd.Int(id)))
}

/*
Sets the position of the specified node.
*/
func (self Instance) SetNodePosition(atype classdb.VisualShaderType, id int, position gd.Vector2) {
	class(self).SetNodePosition(atype, gd.Int(id), position)
}

/*
Returns the position of the specified node within the shader graph.
*/
func (self Instance) GetNodePosition(atype classdb.VisualShaderType, id int) gd.Vector2 {
	return gd.Vector2(class(self).GetNodePosition(atype, gd.Int(id)))
}

/*
Returns the list of all nodes in the shader with the specified type.
*/
func (self Instance) GetNodeList(atype classdb.VisualShaderType) []int32 {
	return []int32(class(self).GetNodeList(atype).AsSlice())
}

/*
Returns next valid node ID that can be added to the shader graph.
*/
func (self Instance) GetValidNodeId(atype classdb.VisualShaderType) int {
	return int(int(class(self).GetValidNodeId(atype)))
}

/*
Removes the specified node from the shader.
*/
func (self Instance) RemoveNode(atype classdb.VisualShaderType, id int) {
	class(self).RemoveNode(atype, gd.Int(id))
}

/*
Replaces the specified node with a node of new class type.
*/
func (self Instance) ReplaceNode(atype classdb.VisualShaderType, id int, new_class string) {
	class(self).ReplaceNode(atype, gd.Int(id), gd.NewStringName(new_class))
}

/*
Returns [code]true[/code] if the specified node and port connection exist.
*/
func (self Instance) IsNodeConnection(atype classdb.VisualShaderType, from_node int, from_port int, to_node int, to_port int) bool {
	return bool(class(self).IsNodeConnection(atype, gd.Int(from_node), gd.Int(from_port), gd.Int(to_node), gd.Int(to_port)))
}

/*
Returns [code]true[/code] if the specified nodes and ports can be connected together.
*/
func (self Instance) CanConnectNodes(atype classdb.VisualShaderType, from_node int, from_port int, to_node int, to_port int) bool {
	return bool(class(self).CanConnectNodes(atype, gd.Int(from_node), gd.Int(from_port), gd.Int(to_node), gd.Int(to_port)))
}

/*
Connects the specified nodes and ports.
*/
func (self Instance) ConnectNodes(atype classdb.VisualShaderType, from_node int, from_port int, to_node int, to_port int) gd.Error {
	return gd.Error(class(self).ConnectNodes(atype, gd.Int(from_node), gd.Int(from_port), gd.Int(to_node), gd.Int(to_port)))
}

/*
Connects the specified nodes and ports.
*/
func (self Instance) DisconnectNodes(atype classdb.VisualShaderType, from_node int, from_port int, to_node int, to_port int) {
	class(self).DisconnectNodes(atype, gd.Int(from_node), gd.Int(from_port), gd.Int(to_node), gd.Int(to_port))
}

/*
Connects the specified nodes and ports, even if they can't be connected. Such connection is invalid and will not function properly.
*/
func (self Instance) ConnectNodesForced(atype classdb.VisualShaderType, from_node int, from_port int, to_node int, to_port int) {
	class(self).ConnectNodesForced(atype, gd.Int(from_node), gd.Int(from_port), gd.Int(to_node), gd.Int(to_port))
}

/*
Returns the list of connected nodes with the specified type.
*/
func (self Instance) GetNodeConnections(atype classdb.VisualShaderType) gd.Array {
	return gd.Array(class(self).GetNodeConnections(atype))
}

/*
Attaches the given node to the given frame.
*/
func (self Instance) AttachNodeToFrame(atype classdb.VisualShaderType, id int, frame_ int) {
	class(self).AttachNodeToFrame(atype, gd.Int(id), gd.Int(frame_))
}

/*
Detaches the given node from the frame it is attached to.
*/
func (self Instance) DetachNodeFromFrame(atype classdb.VisualShaderType, id int) {
	class(self).DetachNodeFromFrame(atype, gd.Int(id))
}

/*
Adds a new varying value node to the shader.
*/
func (self Instance) AddVarying(name string, mode classdb.VisualShaderVaryingMode, atype classdb.VisualShaderVaryingType) {
	class(self).AddVarying(gd.NewString(name), mode, atype)
}

/*
Removes a varying value node with the given [param name]. Prints an error if a node with this name is not found.
*/
func (self Instance) RemoveVarying(name string) {
	class(self).RemoveVarying(gd.NewString(name))
}

/*
Returns [code]true[/code] if the shader has a varying with the given [param name].
*/
func (self Instance) HasVarying(name string) bool {
	return bool(class(self).HasVarying(gd.NewString(name)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.VisualShader

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShader"))
	return Instance{classdb.VisualShader(object)}
}

func (self Instance) GraphOffset() gd.Vector2 {
	return gd.Vector2(class(self).GetGraphOffset())
}

func (self Instance) SetGraphOffset(value gd.Vector2) {
	class(self).SetGraphOffset(value)
}

/*
Sets the mode of this shader.
*/
//go:nosplit
func (self class) SetMode(mode classdb.ShaderMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_set_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds the specified [param node] to the shader.
*/
//go:nosplit
func (self class) AddNode(atype classdb.VisualShaderType, node gdclass.VisualShaderNode, position gd.Vector2, id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, pointers.Get(node[0])[0])
	callframe.Arg(frame, position)
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_add_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the shader node instance with specified [param type] and [param id].
*/
//go:nosplit
func (self class) GetNode(atype classdb.VisualShaderType, id gd.Int) gdclass.VisualShaderNode {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_get_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.VisualShaderNode{classdb.VisualShaderNode(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Sets the position of the specified node.
*/
//go:nosplit
func (self class) SetNodePosition(atype classdb.VisualShaderType, id gd.Int, position gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, id)
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_set_node_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the position of the specified node within the shader graph.
*/
//go:nosplit
func (self class) GetNodePosition(atype classdb.VisualShaderType, id gd.Int) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_get_node_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the list of all nodes in the shader with the specified type.
*/
//go:nosplit
func (self class) GetNodeList(atype classdb.VisualShaderType) gd.PackedInt32Array {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_get_node_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns next valid node ID that can be added to the shader graph.
*/
//go:nosplit
func (self class) GetValidNodeId(atype classdb.VisualShaderType) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_get_valid_node_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes the specified node from the shader.
*/
//go:nosplit
func (self class) RemoveNode(atype classdb.VisualShaderType, id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_remove_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Replaces the specified node with a node of new class type.
*/
//go:nosplit
func (self class) ReplaceNode(atype classdb.VisualShaderType, id gd.Int, new_class gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, id)
	callframe.Arg(frame, pointers.Get(new_class))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_replace_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the specified node and port connection exist.
*/
//go:nosplit
func (self class) IsNodeConnection(atype classdb.VisualShaderType, from_node gd.Int, from_port gd.Int, to_node gd.Int, to_port gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, from_node)
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, to_node)
	callframe.Arg(frame, to_port)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_is_node_connection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the specified nodes and ports can be connected together.
*/
//go:nosplit
func (self class) CanConnectNodes(atype classdb.VisualShaderType, from_node gd.Int, from_port gd.Int, to_node gd.Int, to_port gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, from_node)
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, to_node)
	callframe.Arg(frame, to_port)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_can_connect_nodes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Connects the specified nodes and ports.
*/
//go:nosplit
func (self class) ConnectNodes(atype classdb.VisualShaderType, from_node gd.Int, from_port gd.Int, to_node gd.Int, to_port gd.Int) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, from_node)
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, to_node)
	callframe.Arg(frame, to_port)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_connect_nodes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Connects the specified nodes and ports.
*/
//go:nosplit
func (self class) DisconnectNodes(atype classdb.VisualShaderType, from_node gd.Int, from_port gd.Int, to_node gd.Int, to_port gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, from_node)
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, to_node)
	callframe.Arg(frame, to_port)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_disconnect_nodes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Connects the specified nodes and ports, even if they can't be connected. Such connection is invalid and will not function properly.
*/
//go:nosplit
func (self class) ConnectNodesForced(atype classdb.VisualShaderType, from_node gd.Int, from_port gd.Int, to_node gd.Int, to_port gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, from_node)
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, to_node)
	callframe.Arg(frame, to_port)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_connect_nodes_forced, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the list of connected nodes with the specified type.
*/
//go:nosplit
func (self class) GetNodeConnections(atype classdb.VisualShaderType) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_get_node_connections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGraphOffset(offset gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_set_graph_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGraphOffset() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_get_graph_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Attaches the given node to the given frame.
*/
//go:nosplit
func (self class) AttachNodeToFrame(atype classdb.VisualShaderType, id gd.Int, frame_ gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, id)
	callframe.Arg(frame, frame_)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_attach_node_to_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Detaches the given node from the frame it is attached to.
*/
//go:nosplit
func (self class) DetachNodeFromFrame(atype classdb.VisualShaderType, id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_detach_node_from_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a new varying value node to the shader.
*/
//go:nosplit
func (self class) AddVarying(name gd.String, mode classdb.VisualShaderVaryingMode, atype classdb.VisualShaderVaryingType) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, mode)
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_add_varying, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes a varying value node with the given [param name]. Prints an error if a node with this name is not found.
*/
//go:nosplit
func (self class) RemoveVarying(name gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_remove_varying, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the shader has a varying with the given [param name].
*/
//go:nosplit
func (self class) HasVarying(name gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShader.Bind_has_varying, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShader() Advanced     { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVisualShader() Instance  { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsShader() Shader.Advanced    { return *((*Shader.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsShader() Shader.Instance { return *((*Shader.Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(self.AsShader(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsShader(), name)
	}
}
func init() {
	classdb.Register("VisualShader", func(ptr gd.Object) any { return classdb.VisualShader(ptr) })
}

type Type = classdb.VisualShaderType

const (
	/*A vertex shader, operating on vertices.*/
	TypeVertex Type = 0
	/*A fragment shader, operating on fragments (pixels).*/
	TypeFragment Type = 1
	/*A shader for light calculations.*/
	TypeLight Type = 2
	/*A function for the "start" stage of particle shader.*/
	TypeStart Type = 3
	/*A function for the "process" stage of particle shader.*/
	TypeProcess Type = 4
	/*A function for the "collide" stage (particle collision handler) of particle shader.*/
	TypeCollide Type = 5
	/*A function for the "start" stage of particle shader, with customized output.*/
	TypeStartCustom Type = 6
	/*A function for the "process" stage of particle shader, with customized output.*/
	TypeProcessCustom Type = 7
	/*A shader for 3D environment's sky.*/
	TypeSky Type = 8
	/*A compute shader that runs for each froxel of the volumetric fog map.*/
	TypeFog Type = 9
	/*Represents the size of the [enum Type] enum.*/
	TypeMax Type = 10
)

type VaryingMode = classdb.VisualShaderVaryingMode

const (
	/*Varying is passed from [code]Vertex[/code] function to [code]Fragment[/code] and [code]Light[/code] functions.*/
	VaryingModeVertexToFragLight VaryingMode = 0
	/*Varying is passed from [code]Fragment[/code] function to [code]Light[/code] function.*/
	VaryingModeFragToLight VaryingMode = 1
	/*Represents the size of the [enum VaryingMode] enum.*/
	VaryingModeMax VaryingMode = 2
)

type VaryingType = classdb.VisualShaderVaryingType

const (
	/*Varying is of type [float].*/
	VaryingTypeFloat VaryingType = 0
	/*Varying is of type [int].*/
	VaryingTypeInt VaryingType = 1
	/*Varying is of type unsigned [int].*/
	VaryingTypeUint VaryingType = 2
	/*Varying is of type [Vector2].*/
	VaryingTypeVector2d VaryingType = 3
	/*Varying is of type [Vector3].*/
	VaryingTypeVector3d VaryingType = 4
	/*Varying is of type [Vector4].*/
	VaryingTypeVector4d VaryingType = 5
	/*Varying is of type [bool].*/
	VaryingTypeBoolean VaryingType = 6
	/*Varying is of type [Transform3D].*/
	VaryingTypeTransform VaryingType = 7
	/*Represents the size of the [enum VaryingType] enum.*/
	VaryingTypeMax VaryingType = 8
)
