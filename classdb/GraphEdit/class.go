// Package GraphEdit provides methods for working with GraphEdit object instances.
package GraphEdit

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/Rect2"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
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
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
[GraphEdit] provides tools for creation, manipulation, and display of various graphs. Its main purpose in the engine is to power the visual programming systems, such as visual shaders, but it is also available for use in user projects.
[GraphEdit] by itself is only an empty container, representing an infinite grid where [GraphNode]s can be placed. Each [GraphNode] represents a node in the graph, a single unit of data in the connected scheme. [GraphEdit], in turn, helps to control various interactions with nodes and between nodes. When the user attempts to connect, disconnect, or delete a [GraphNode], a signal is emitted in the [GraphEdit], but no action is taken by default. It is the responsibility of the programmer utilizing this control to implement the necessary logic to determine how each request should be handled.
[b]Performance:[/b] It is greatly advised to enable low-processor usage mode (see [member OS.low_processor_usage_mode]) when using GraphEdits.
[b]Note:[/b] Keep in mind that [method Node.get_children] will also return the connection layer node named [code]_connection_layer[/code] due to technical limitations. This behavior may change in future releases.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=GraphEdit)
*/
type Instance [1]gdclass.GraphEdit

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGraphEdit() Instance
}
type Interface interface {
	//Returns whether the [param mouse_position] is in the input hot zone.
	//By default, a hot zone is a [Rect2] positioned such that its center is at [param in_node].[method GraphNode.get_input_port_position]([param in_port]) (For output's case, call [method GraphNode.get_output_port_position] instead). The hot zone's width is twice the Theme Property [code]port_grab_distance_horizontal[/code], and its height is twice the [code]port_grab_distance_vertical[/code].
	//Below is a sample code to help get started:
	//[codeblock]
	//func _is_in_input_hotzone(in_node, in_port, mouse_position):
	//    var port_size: Vector2 = Vector2(get_theme_constant("port_grab_distance_horizontal"), get_theme_constant("port_grab_distance_vertical"))
	//    var port_pos: Vector2 = in_node.get_position() + in_node.get_input_port_position(in_port) - port_size / 2
	//    var rect = Rect2(port_pos, port_size)
	//
	//    return rect.has_point(mouse_position)
	//[/codeblock]
	IsInInputHotzone(in_node Object.Instance, in_port int, mouse_position Vector2.XY) bool
	//Returns whether the [param mouse_position] is in the output hot zone. For more information on hot zones, see [method _is_in_input_hotzone].
	//Below is a sample code to help get started:
	//[codeblock]
	//func _is_in_output_hotzone(in_node, in_port, mouse_position):
	//    var port_size: Vector2 = Vector2(get_theme_constant("port_grab_distance_horizontal"), get_theme_constant("port_grab_distance_vertical"))
	//    var port_pos: Vector2 = in_node.get_position() + in_node.get_output_port_position(in_port) - port_size / 2
	//    var rect = Rect2(port_pos, port_size)
	//
	//    return rect.has_point(mouse_position)
	//[/codeblock]
	IsInOutputHotzone(in_node Object.Instance, in_port int, mouse_position Vector2.XY) bool
	//Virtual method which can be overridden to customize how connections are drawn.
	GetConnectionLine(from_position Vector2.XY, to_position Vector2.XY) []Vector2.XY
	//This virtual method can be used to insert additional error detection while the user is dragging a connection over a valid port.
	//Return [code]true[/code] if the connection is indeed valid or return [code]false[/code] if the connection is impossible. If the connection is impossible, no snapping to the port and thus no connection request to that port will happen.
	//In this example a connection to same node is suppressed:
	//[codeblocks]
	//[gdscript]
	//func _is_node_hover_valid(from, from_port, to, to_port):
	//    return from != to
	//[/gdscript]
	//[csharp]
	//public override bool _IsNodeHoverValid(StringName fromNode, int fromPort, StringName toNode, int toPort)
	//{
	//    return fromNode != toNode;
	//}
	//[/csharp]
	//[/codeblocks]
	IsNodeHoverValid(from_node string, from_port int, to_node string, to_port int) bool
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) IsInInputHotzone(in_node Object.Instance, in_port int, mouse_position Vector2.XY) (_ bool) {
	return
}
func (self implementation) IsInOutputHotzone(in_node Object.Instance, in_port int, mouse_position Vector2.XY) (_ bool) {
	return
}
func (self implementation) GetConnectionLine(from_position Vector2.XY, to_position Vector2.XY) (_ []Vector2.XY) {
	return
}
func (self implementation) IsNodeHoverValid(from_node string, from_port int, to_node string, to_port int) (_ bool) {
	return
}

/*
Returns whether the [param mouse_position] is in the input hot zone.
By default, a hot zone is a [Rect2] positioned such that its center is at [param in_node].[method GraphNode.get_input_port_position]([param in_port]) (For output's case, call [method GraphNode.get_output_port_position] instead). The hot zone's width is twice the Theme Property [code]port_grab_distance_horizontal[/code], and its height is twice the [code]port_grab_distance_vertical[/code].
Below is a sample code to help get started:
[codeblock]
func _is_in_input_hotzone(in_node, in_port, mouse_position):

	var port_size: Vector2 = Vector2(get_theme_constant("port_grab_distance_horizontal"), get_theme_constant("port_grab_distance_vertical"))
	var port_pos: Vector2 = in_node.get_position() + in_node.get_input_port_position(in_port) - port_size / 2
	var rect = Rect2(port_pos, port_size)

	return rect.has_point(mouse_position)

[/codeblock]
*/
func (Instance) _is_in_input_hotzone(impl func(ptr unsafe.Pointer, in_node Object.Instance, in_port int, mouse_position Vector2.XY) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var in_node = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(in_node[0])
		var in_port = gd.UnsafeGet[int64](p_args, 1)

		var mouse_position = gd.UnsafeGet[Vector2.XY](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, in_node, int(in_port), mouse_position)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns whether the [param mouse_position] is in the output hot zone. For more information on hot zones, see [method _is_in_input_hotzone].
Below is a sample code to help get started:
[codeblock]
func _is_in_output_hotzone(in_node, in_port, mouse_position):

	var port_size: Vector2 = Vector2(get_theme_constant("port_grab_distance_horizontal"), get_theme_constant("port_grab_distance_vertical"))
	var port_pos: Vector2 = in_node.get_position() + in_node.get_output_port_position(in_port) - port_size / 2
	var rect = Rect2(port_pos, port_size)

	return rect.has_point(mouse_position)

[/codeblock]
*/
func (Instance) _is_in_output_hotzone(impl func(ptr unsafe.Pointer, in_node Object.Instance, in_port int, mouse_position Vector2.XY) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var in_node = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(in_node[0])
		var in_port = gd.UnsafeGet[int64](p_args, 1)

		var mouse_position = gd.UnsafeGet[Vector2.XY](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, in_node, int(in_port), mouse_position)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Virtual method which can be overridden to customize how connections are drawn.
*/
func (Instance) _get_connection_line(impl func(ptr unsafe.Pointer, from_position Vector2.XY, to_position Vector2.XY) []Vector2.XY) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var from_position = gd.UnsafeGet[Vector2.XY](p_args, 0)

		var to_position = gd.UnsafeGet[Vector2.XY](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from_position, to_position)
		ptr, ok := pointers.End(gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](Packed.New(ret...)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
This virtual method can be used to insert additional error detection while the user is dragging a connection over a valid port.
Return [code]true[/code] if the connection is indeed valid or return [code]false[/code] if the connection is impossible. If the connection is impossible, no snapping to the port and thus no connection request to that port will happen.
In this example a connection to same node is suppressed:
[codeblocks]
[gdscript]
func _is_node_hover_valid(from, from_port, to, to_port):

	return from != to

[/gdscript]
[csharp]
public override bool _IsNodeHoverValid(StringName fromNode, int fromPort, StringName toNode, int toPort)

	{
	    return fromNode != toNode;
	}

[/csharp]
[/codeblocks]
*/
func (Instance) _is_node_hover_valid(impl func(ptr unsafe.Pointer, from_node string, from_port int, to_node string, to_port int) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var from_node = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0)))))
		defer pointers.End(gd.InternalStringName(from_node))
		var from_port = gd.UnsafeGet[int64](p_args, 1)

		var to_node = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2)))))
		defer pointers.End(gd.InternalStringName(to_node))
		var to_port = gd.UnsafeGet[int64](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from_node.String(), int(from_port), to_node.String(), int(to_port))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Create a connection between the [param from_port] of the [param from_node] [GraphNode] and the [param to_port] of the [param to_node] [GraphNode]. If the connection already exists, no connection is created.
*/
func (self Instance) ConnectNode(from_node string, from_port int, to_node string, to_port int) error { //gd:GraphEdit.connect_node
	return error(gd.ToError(class(self).ConnectNode(String.Name(String.New(from_node)), int64(from_port), String.Name(String.New(to_node)), int64(to_port))))
}

/*
Returns [code]true[/code] if the [param from_port] of the [param from_node] [GraphNode] is connected to the [param to_port] of the [param to_node] [GraphNode].
*/
func (self Instance) IsNodeConnected(from_node string, from_port int, to_node string, to_port int) bool { //gd:GraphEdit.is_node_connected
	return bool(class(self).IsNodeConnected(String.Name(String.New(from_node)), int64(from_port), String.Name(String.New(to_node)), int64(to_port)))
}

/*
Removes the connection between the [param from_port] of the [param from_node] [GraphNode] and the [param to_port] of the [param to_node] [GraphNode]. If the connection does not exist, no connection is removed.
*/
func (self Instance) DisconnectNode(from_node string, from_port int, to_node string, to_port int) { //gd:GraphEdit.disconnect_node
	class(self).DisconnectNode(String.Name(String.New(from_node)), int64(from_port), String.Name(String.New(to_node)), int64(to_port))
}

/*
Sets the coloration of the connection between [param from_node]'s [param from_port] and [param to_node]'s [param to_port] with the color provided in the [theme_item activity] theme property. The color is linearly interpolated between the connection color and the activity color using [param amount] as weight.
*/
func (self Instance) SetConnectionActivity(from_node string, from_port int, to_node string, to_port int, amount Float.X) { //gd:GraphEdit.set_connection_activity
	class(self).SetConnectionActivity(String.Name(String.New(from_node)), int64(from_port), String.Name(String.New(to_node)), int64(to_port), float64(amount))
}

/*
Returns an [Array] containing the list of connections. A connection consists in a structure of the form [code]{ from_port: 0, from_node: "GraphNode name 0", to_port: 1, to_node: "GraphNode name 1" }[/code].
*/
func (self Instance) GetConnectionList() []Connection { //gd:GraphEdit.get_connection_list
	return []Connection(gd.ArrayAs[[]Connection](gd.InternalArray(class(self).GetConnectionList())))
}

/*
Returns the closest connection to the given point in screen space. If no connection is found within [param max_distance] pixels, an empty [Dictionary] is returned.
A connection consists in a structure of the form [code]{ from_port: 0, from_node: "GraphNode name 0", to_port: 1, to_node: "GraphNode name 1" }[/code].
For example, getting a connection at a given mouse position can be achieved like this:
[codeblocks]
[gdscript]
var connection = get_closest_connection_at_point(mouse_event.get_position())
[/gdscript]
[/codeblocks]
*/
func (self Instance) GetClosestConnectionAtPoint(point Vector2.XY) Connection { //gd:GraphEdit.get_closest_connection_at_point
	return Connection(gd.DictionaryAs[Connection](class(self).GetClosestConnectionAtPoint(Vector2.XY(point), float64(4.0))))
}

/*
Returns an [Array] containing the list of connections that intersect with the given [Rect2]. A connection consists in a structure of the form [code]{ from_port: 0, from_node: "GraphNode name 0", to_port: 1, to_node: "GraphNode name 1" }[/code].
*/
func (self Instance) GetConnectionsIntersectingWithRect(rect Rect2.PositionSize) []Connection { //gd:GraphEdit.get_connections_intersecting_with_rect
	return []Connection(gd.ArrayAs[[]Connection](gd.InternalArray(class(self).GetConnectionsIntersectingWithRect(Rect2.PositionSize(rect)))))
}

/*
Removes all connections between nodes.
*/
func (self Instance) ClearConnections() { //gd:GraphEdit.clear_connections
	class(self).ClearConnections()
}

/*
Ends the creation of the current connection. In other words, if you are dragging a connection you can use this method to abort the process and remove the line that followed your cursor.
This is best used together with [signal connection_drag_started] and [signal connection_drag_ended] to add custom behavior like node addition through shortcuts.
[b]Note:[/b] This method suppresses any other connection request signals apart from [signal connection_drag_ended].
*/
func (self Instance) ForceConnectionDragEnd() { //gd:GraphEdit.force_connection_drag_end
	class(self).ForceConnectionDragEnd()
}

/*
Allows to disconnect nodes when dragging from the right port of the [GraphNode]'s slot if it has the specified type. See also [method remove_valid_right_disconnect_type].
*/
func (self Instance) AddValidRightDisconnectType(atype int) { //gd:GraphEdit.add_valid_right_disconnect_type
	class(self).AddValidRightDisconnectType(int64(atype))
}

/*
Disallows to disconnect nodes when dragging from the right port of the [GraphNode]'s slot if it has the specified type. Use this to disable disconnection previously allowed with [method add_valid_right_disconnect_type].
*/
func (self Instance) RemoveValidRightDisconnectType(atype int) { //gd:GraphEdit.remove_valid_right_disconnect_type
	class(self).RemoveValidRightDisconnectType(int64(atype))
}

/*
Allows to disconnect nodes when dragging from the left port of the [GraphNode]'s slot if it has the specified type. See also [method remove_valid_left_disconnect_type].
*/
func (self Instance) AddValidLeftDisconnectType(atype int) { //gd:GraphEdit.add_valid_left_disconnect_type
	class(self).AddValidLeftDisconnectType(int64(atype))
}

/*
Disallows to disconnect nodes when dragging from the left port of the [GraphNode]'s slot if it has the specified type. Use this to disable disconnection previously allowed with [method add_valid_left_disconnect_type].
*/
func (self Instance) RemoveValidLeftDisconnectType(atype int) { //gd:GraphEdit.remove_valid_left_disconnect_type
	class(self).RemoveValidLeftDisconnectType(int64(atype))
}

/*
Allows the connection between two different port types. The port type is defined individually for the left and the right port of each slot with the [method GraphNode.set_slot] method.
See also [method is_valid_connection_type] and [method remove_valid_connection_type].
*/
func (self Instance) AddValidConnectionType(from_type int, to_type int) { //gd:GraphEdit.add_valid_connection_type
	class(self).AddValidConnectionType(int64(from_type), int64(to_type))
}

/*
Disallows the connection between two different port types previously allowed by [method add_valid_connection_type]. The port type is defined individually for the left and the right port of each slot with the [method GraphNode.set_slot] method.
See also [method is_valid_connection_type].
*/
func (self Instance) RemoveValidConnectionType(from_type int, to_type int) { //gd:GraphEdit.remove_valid_connection_type
	class(self).RemoveValidConnectionType(int64(from_type), int64(to_type))
}

/*
Returns whether it's possible to make a connection between two different port types. The port type is defined individually for the left and the right port of each slot with the [method GraphNode.set_slot] method.
See also [method add_valid_connection_type] and [method remove_valid_connection_type].
*/
func (self Instance) IsValidConnectionType(from_type int, to_type int) bool { //gd:GraphEdit.is_valid_connection_type
	return bool(class(self).IsValidConnectionType(int64(from_type), int64(to_type)))
}

/*
Returns the points which would make up a connection between [param from_node] and [param to_node].
*/
func (self Instance) GetConnectionLine(from_node Vector2.XY, to_node Vector2.XY) []Vector2.XY { //gd:GraphEdit.get_connection_line
	return []Vector2.XY(slices.Collect(class(self).GetConnectionLine(Vector2.XY(from_node), Vector2.XY(to_node)).Values()))
}

/*
Attaches the [param element] [GraphElement] to the [param frame] [GraphFrame].
*/
func (self Instance) AttachGraphElementToFrame(element string, frame_ string) { //gd:GraphEdit.attach_graph_element_to_frame
	class(self).AttachGraphElementToFrame(String.Name(String.New(element)), String.Name(String.New(frame_)))
}

/*
Detaches the [param element] [GraphElement] from the [GraphFrame] it is currently attached to.
*/
func (self Instance) DetachGraphElementFromFrame(element string) { //gd:GraphEdit.detach_graph_element_from_frame
	class(self).DetachGraphElementFromFrame(String.Name(String.New(element)))
}

/*
Returns the [GraphFrame] that contains the [GraphElement] with the given name.
*/
func (self Instance) GetElementFrame(element string) [1]gdclass.GraphFrame { //gd:GraphEdit.get_element_frame
	return [1]gdclass.GraphFrame(class(self).GetElementFrame(String.Name(String.New(element))))
}

/*
Returns an array of node names that are attached to the [GraphFrame] with the given name.
*/
func (self Instance) GetAttachedNodesOfFrame(frame_ string) []string { //gd:GraphEdit.get_attached_nodes_of_frame
	return []string(gd.ArrayAs[[]string](gd.InternalArray(class(self).GetAttachedNodesOfFrame(String.Name(String.New(frame_))))))
}

/*
Gets the [HBoxContainer] that contains the zooming and grid snap controls in the top left of the graph. You can use this method to reposition the toolbar or to add your own custom controls to it.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
func (self Instance) GetMenuHbox() [1]gdclass.HBoxContainer { //gd:GraphEdit.get_menu_hbox
	return [1]gdclass.HBoxContainer(class(self).GetMenuHbox())
}

/*
Rearranges selected nodes in a layout with minimum crossings between connections and uniform horizontal and vertical gap between nodes.
*/
func (self Instance) ArrangeNodes() { //gd:GraphEdit.arrange_nodes
	class(self).ArrangeNodes()
}

/*
Sets the specified [param node] as the one selected.
*/
func (self Instance) SetSelected(node [1]gdclass.Node) { //gd:GraphEdit.set_selected
	class(self).SetSelected(node)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GraphEdit

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GraphEdit"))
	casted := Instance{*(*gdclass.GraphEdit)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) ScrollOffset() Vector2.XY {
	return Vector2.XY(class(self).GetScrollOffset())
}

func (self Instance) SetScrollOffset(value Vector2.XY) {
	class(self).SetScrollOffset(Vector2.XY(value))
}

func (self Instance) ShowGrid() bool {
	return bool(class(self).IsShowingGrid())
}

func (self Instance) SetShowGrid(value bool) {
	class(self).SetShowGrid(value)
}

func (self Instance) GridPattern() gdclass.GraphEditGridPattern {
	return gdclass.GraphEditGridPattern(class(self).GetGridPattern())
}

func (self Instance) SetGridPattern(value gdclass.GraphEditGridPattern) {
	class(self).SetGridPattern(value)
}

func (self Instance) SnappingEnabled() bool {
	return bool(class(self).IsSnappingEnabled())
}

func (self Instance) SetSnappingEnabled(value bool) {
	class(self).SetSnappingEnabled(value)
}

func (self Instance) SnappingDistance() int {
	return int(int(class(self).GetSnappingDistance()))
}

func (self Instance) SetSnappingDistance(value int) {
	class(self).SetSnappingDistance(int64(value))
}

func (self Instance) PanningScheme() gdclass.GraphEditPanningScheme {
	return gdclass.GraphEditPanningScheme(class(self).GetPanningScheme())
}

func (self Instance) SetPanningScheme(value gdclass.GraphEditPanningScheme) {
	class(self).SetPanningScheme(value)
}

func (self Instance) RightDisconnects() bool {
	return bool(class(self).IsRightDisconnectsEnabled())
}

func (self Instance) SetRightDisconnects(value bool) {
	class(self).SetRightDisconnects(value)
}

func (self Instance) ConnectionLinesCurvature() Float.X {
	return Float.X(Float.X(class(self).GetConnectionLinesCurvature()))
}

func (self Instance) SetConnectionLinesCurvature(value Float.X) {
	class(self).SetConnectionLinesCurvature(float64(value))
}

func (self Instance) ConnectionLinesThickness() Float.X {
	return Float.X(Float.X(class(self).GetConnectionLinesThickness()))
}

func (self Instance) SetConnectionLinesThickness(value Float.X) {
	class(self).SetConnectionLinesThickness(float64(value))
}

func (self Instance) ConnectionLinesAntialiased() bool {
	return bool(class(self).IsConnectionLinesAntialiased())
}

func (self Instance) SetConnectionLinesAntialiased(value bool) {
	class(self).SetConnectionLinesAntialiased(value)
}

func (self Instance) Zoom() Float.X {
	return Float.X(Float.X(class(self).GetZoom()))
}

func (self Instance) SetZoom(value Float.X) {
	class(self).SetZoom(float64(value))
}

func (self Instance) ZoomMin() Float.X {
	return Float.X(Float.X(class(self).GetZoomMin()))
}

func (self Instance) SetZoomMin(value Float.X) {
	class(self).SetZoomMin(float64(value))
}

func (self Instance) ZoomMax() Float.X {
	return Float.X(Float.X(class(self).GetZoomMax()))
}

func (self Instance) SetZoomMax(value Float.X) {
	class(self).SetZoomMax(float64(value))
}

func (self Instance) ZoomStep() Float.X {
	return Float.X(Float.X(class(self).GetZoomStep()))
}

func (self Instance) SetZoomStep(value Float.X) {
	class(self).SetZoomStep(float64(value))
}

func (self Instance) MinimapEnabled() bool {
	return bool(class(self).IsMinimapEnabled())
}

func (self Instance) SetMinimapEnabled(value bool) {
	class(self).SetMinimapEnabled(value)
}

func (self Instance) MinimapSize() Vector2.XY {
	return Vector2.XY(class(self).GetMinimapSize())
}

func (self Instance) SetMinimapSize(value Vector2.XY) {
	class(self).SetMinimapSize(Vector2.XY(value))
}

func (self Instance) MinimapOpacity() Float.X {
	return Float.X(Float.X(class(self).GetMinimapOpacity()))
}

func (self Instance) SetMinimapOpacity(value Float.X) {
	class(self).SetMinimapOpacity(float64(value))
}

func (self Instance) ShowMenu() bool {
	return bool(class(self).IsShowingMenu())
}

func (self Instance) SetShowMenu(value bool) {
	class(self).SetShowMenu(value)
}

func (self Instance) ShowZoomLabel() bool {
	return bool(class(self).IsShowingZoomLabel())
}

func (self Instance) SetShowZoomLabel(value bool) {
	class(self).SetShowZoomLabel(value)
}

func (self Instance) ShowZoomButtons() bool {
	return bool(class(self).IsShowingZoomButtons())
}

func (self Instance) SetShowZoomButtons(value bool) {
	class(self).SetShowZoomButtons(value)
}

func (self Instance) ShowGridButtons() bool {
	return bool(class(self).IsShowingGridButtons())
}

func (self Instance) SetShowGridButtons(value bool) {
	class(self).SetShowGridButtons(value)
}

func (self Instance) ShowMinimapButton() bool {
	return bool(class(self).IsShowingMinimapButton())
}

func (self Instance) SetShowMinimapButton(value bool) {
	class(self).SetShowMinimapButton(value)
}

func (self Instance) ShowArrangeButton() bool {
	return bool(class(self).IsShowingArrangeButton())
}

func (self Instance) SetShowArrangeButton(value bool) {
	class(self).SetShowArrangeButton(value)
}

/*
Returns whether the [param mouse_position] is in the input hot zone.
By default, a hot zone is a [Rect2] positioned such that its center is at [param in_node].[method GraphNode.get_input_port_position]([param in_port]) (For output's case, call [method GraphNode.get_output_port_position] instead). The hot zone's width is twice the Theme Property [code]port_grab_distance_horizontal[/code], and its height is twice the [code]port_grab_distance_vertical[/code].
Below is a sample code to help get started:
[codeblock]
func _is_in_input_hotzone(in_node, in_port, mouse_position):

	var port_size: Vector2 = Vector2(get_theme_constant("port_grab_distance_horizontal"), get_theme_constant("port_grab_distance_vertical"))
	var port_pos: Vector2 = in_node.get_position() + in_node.get_input_port_position(in_port) - port_size / 2
	var rect = Rect2(port_pos, port_size)

	return rect.has_point(mouse_position)

[/codeblock]
*/
func (class) _is_in_input_hotzone(impl func(ptr unsafe.Pointer, in_node [1]gd.Object, in_port int64, mouse_position Vector2.XY) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var in_node = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(in_node[0])
		var in_port = gd.UnsafeGet[int64](p_args, 1)

		var mouse_position = gd.UnsafeGet[Vector2.XY](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, in_node, in_port, mouse_position)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns whether the [param mouse_position] is in the output hot zone. For more information on hot zones, see [method _is_in_input_hotzone].
Below is a sample code to help get started:
[codeblock]
func _is_in_output_hotzone(in_node, in_port, mouse_position):

	var port_size: Vector2 = Vector2(get_theme_constant("port_grab_distance_horizontal"), get_theme_constant("port_grab_distance_vertical"))
	var port_pos: Vector2 = in_node.get_position() + in_node.get_output_port_position(in_port) - port_size / 2
	var rect = Rect2(port_pos, port_size)

	return rect.has_point(mouse_position)

[/codeblock]
*/
func (class) _is_in_output_hotzone(impl func(ptr unsafe.Pointer, in_node [1]gd.Object, in_port int64, mouse_position Vector2.XY) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var in_node = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(in_node[0])
		var in_port = gd.UnsafeGet[int64](p_args, 1)

		var mouse_position = gd.UnsafeGet[Vector2.XY](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, in_node, in_port, mouse_position)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Virtual method which can be overridden to customize how connections are drawn.
*/
func (class) _get_connection_line(impl func(ptr unsafe.Pointer, from_position Vector2.XY, to_position Vector2.XY) Packed.Array[Vector2.XY]) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var from_position = gd.UnsafeGet[Vector2.XY](p_args, 0)

		var to_position = gd.UnsafeGet[Vector2.XY](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from_position, to_position)
		ptr, ok := pointers.End(gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
This virtual method can be used to insert additional error detection while the user is dragging a connection over a valid port.
Return [code]true[/code] if the connection is indeed valid or return [code]false[/code] if the connection is impossible. If the connection is impossible, no snapping to the port and thus no connection request to that port will happen.
In this example a connection to same node is suppressed:
[codeblocks]
[gdscript]
func _is_node_hover_valid(from, from_port, to, to_port):

	return from != to

[/gdscript]
[csharp]
public override bool _IsNodeHoverValid(StringName fromNode, int fromPort, StringName toNode, int toPort)

	{
	    return fromNode != toNode;
	}

[/csharp]
[/codeblocks]
*/
func (class) _is_node_hover_valid(impl func(ptr unsafe.Pointer, from_node String.Name, from_port int64, to_node String.Name, to_port int64) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var from_node = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0)))))
		defer pointers.End(gd.InternalStringName(from_node))
		var from_port = gd.UnsafeGet[int64](p_args, 1)

		var to_node = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2)))))
		defer pointers.End(gd.InternalStringName(to_node))
		var to_port = gd.UnsafeGet[int64](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, from_node, from_port, to_node, to_port)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Create a connection between the [param from_port] of the [param from_node] [GraphNode] and the [param to_port] of the [param to_node] [GraphNode]. If the connection already exists, no connection is created.
*/
//go:nosplit
func (self class) ConnectNode(from_node String.Name, from_port int64, to_node String.Name, to_port int64) Error.Code { //gd:GraphEdit.connect_node
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(from_node)))
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(to_node)))
	callframe.Arg(frame, to_port)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_connect_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the [param from_port] of the [param from_node] [GraphNode] is connected to the [param to_port] of the [param to_node] [GraphNode].
*/
//go:nosplit
func (self class) IsNodeConnected(from_node String.Name, from_port int64, to_node String.Name, to_port int64) bool { //gd:GraphEdit.is_node_connected
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(from_node)))
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(to_node)))
	callframe.Arg(frame, to_port)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_node_connected, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes the connection between the [param from_port] of the [param from_node] [GraphNode] and the [param to_port] of the [param to_node] [GraphNode]. If the connection does not exist, no connection is removed.
*/
//go:nosplit
func (self class) DisconnectNode(from_node String.Name, from_port int64, to_node String.Name, to_port int64) { //gd:GraphEdit.disconnect_node
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(from_node)))
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(to_node)))
	callframe.Arg(frame, to_port)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_disconnect_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the coloration of the connection between [param from_node]'s [param from_port] and [param to_node]'s [param to_port] with the color provided in the [theme_item activity] theme property. The color is linearly interpolated between the connection color and the activity color using [param amount] as weight.
*/
//go:nosplit
func (self class) SetConnectionActivity(from_node String.Name, from_port int64, to_node String.Name, to_port int64, amount float64) { //gd:GraphEdit.set_connection_activity
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(from_node)))
	callframe.Arg(frame, from_port)
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(to_node)))
	callframe.Arg(frame, to_port)
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_connection_activity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an [Array] containing the list of connections. A connection consists in a structure of the form [code]{ from_port: 0, from_node: "GraphNode name 0", to_port: 1, to_node: "GraphNode name 1" }[/code].
*/
//go:nosplit
func (self class) GetConnectionList() Array.Contains[Dictionary.Any] { //gd:GraphEdit.get_connection_list
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_connection_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Dictionary.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the closest connection to the given point in screen space. If no connection is found within [param max_distance] pixels, an empty [Dictionary] is returned.
A connection consists in a structure of the form [code]{ from_port: 0, from_node: "GraphNode name 0", to_port: 1, to_node: "GraphNode name 1" }[/code].
For example, getting a connection at a given mouse position can be achieved like this:
[codeblocks]
[gdscript]
var connection = get_closest_connection_at_point(mouse_event.get_position())
[/gdscript]
[/codeblocks]
*/
//go:nosplit
func (self class) GetClosestConnectionAtPoint(point Vector2.XY, max_distance float64) Dictionary.Any { //gd:GraphEdit.get_closest_connection_at_point
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, max_distance)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_closest_connection_at_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns an [Array] containing the list of connections that intersect with the given [Rect2]. A connection consists in a structure of the form [code]{ from_port: 0, from_node: "GraphNode name 0", to_port: 1, to_node: "GraphNode name 1" }[/code].
*/
//go:nosplit
func (self class) GetConnectionsIntersectingWithRect(rect Rect2.PositionSize) Array.Contains[Dictionary.Any] { //gd:GraphEdit.get_connections_intersecting_with_rect
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_connections_intersecting_with_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Dictionary.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Removes all connections between nodes.
*/
//go:nosplit
func (self class) ClearConnections() { //gd:GraphEdit.clear_connections
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_clear_connections, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Ends the creation of the current connection. In other words, if you are dragging a connection you can use this method to abort the process and remove the line that followed your cursor.
This is best used together with [signal connection_drag_started] and [signal connection_drag_ended] to add custom behavior like node addition through shortcuts.
[b]Note:[/b] This method suppresses any other connection request signals apart from [signal connection_drag_ended].
*/
//go:nosplit
func (self class) ForceConnectionDragEnd() { //gd:GraphEdit.force_connection_drag_end
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_force_connection_drag_end, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetScrollOffset() Vector2.XY { //gd:GraphEdit.get_scroll_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_scroll_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetScrollOffset(offset Vector2.XY) { //gd:GraphEdit.set_scroll_offset
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_scroll_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Allows to disconnect nodes when dragging from the right port of the [GraphNode]'s slot if it has the specified type. See also [method remove_valid_right_disconnect_type].
*/
//go:nosplit
func (self class) AddValidRightDisconnectType(atype int64) { //gd:GraphEdit.add_valid_right_disconnect_type
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_add_valid_right_disconnect_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Disallows to disconnect nodes when dragging from the right port of the [GraphNode]'s slot if it has the specified type. Use this to disable disconnection previously allowed with [method add_valid_right_disconnect_type].
*/
//go:nosplit
func (self class) RemoveValidRightDisconnectType(atype int64) { //gd:GraphEdit.remove_valid_right_disconnect_type
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_remove_valid_right_disconnect_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Allows to disconnect nodes when dragging from the left port of the [GraphNode]'s slot if it has the specified type. See also [method remove_valid_left_disconnect_type].
*/
//go:nosplit
func (self class) AddValidLeftDisconnectType(atype int64) { //gd:GraphEdit.add_valid_left_disconnect_type
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_add_valid_left_disconnect_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Disallows to disconnect nodes when dragging from the left port of the [GraphNode]'s slot if it has the specified type. Use this to disable disconnection previously allowed with [method add_valid_left_disconnect_type].
*/
//go:nosplit
func (self class) RemoveValidLeftDisconnectType(atype int64) { //gd:GraphEdit.remove_valid_left_disconnect_type
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_remove_valid_left_disconnect_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Allows the connection between two different port types. The port type is defined individually for the left and the right port of each slot with the [method GraphNode.set_slot] method.
See also [method is_valid_connection_type] and [method remove_valid_connection_type].
*/
//go:nosplit
func (self class) AddValidConnectionType(from_type int64, to_type int64) { //gd:GraphEdit.add_valid_connection_type
	var frame = callframe.New()
	callframe.Arg(frame, from_type)
	callframe.Arg(frame, to_type)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_add_valid_connection_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Disallows the connection between two different port types previously allowed by [method add_valid_connection_type]. The port type is defined individually for the left and the right port of each slot with the [method GraphNode.set_slot] method.
See also [method is_valid_connection_type].
*/
//go:nosplit
func (self class) RemoveValidConnectionType(from_type int64, to_type int64) { //gd:GraphEdit.remove_valid_connection_type
	var frame = callframe.New()
	callframe.Arg(frame, from_type)
	callframe.Arg(frame, to_type)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_remove_valid_connection_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether it's possible to make a connection between two different port types. The port type is defined individually for the left and the right port of each slot with the [method GraphNode.set_slot] method.
See also [method add_valid_connection_type] and [method remove_valid_connection_type].
*/
//go:nosplit
func (self class) IsValidConnectionType(from_type int64, to_type int64) bool { //gd:GraphEdit.is_valid_connection_type
	var frame = callframe.New()
	callframe.Arg(frame, from_type)
	callframe.Arg(frame, to_type)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_valid_connection_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the points which would make up a connection between [param from_node] and [param to_node].
*/
//go:nosplit
func (self class) GetConnectionLine(from_node Vector2.XY, to_node Vector2.XY) Packed.Array[Vector2.XY] { //gd:GraphEdit.get_connection_line
	var frame = callframe.New()
	callframe.Arg(frame, from_node)
	callframe.Arg(frame, to_node)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_connection_line, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[Vector2.XY](Array.Through(gd.PackedProxy[gd.PackedVector2Array, Vector2.XY]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Attaches the [param element] [GraphElement] to the [param frame] [GraphFrame].
*/
//go:nosplit
func (self class) AttachGraphElementToFrame(element String.Name, frame_ String.Name) { //gd:GraphEdit.attach_graph_element_to_frame
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(element)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(frame_)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_attach_graph_element_to_frame, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Detaches the [param element] [GraphElement] from the [GraphFrame] it is currently attached to.
*/
//go:nosplit
func (self class) DetachGraphElementFromFrame(element String.Name) { //gd:GraphEdit.detach_graph_element_from_frame
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(element)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_detach_graph_element_from_frame, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [GraphFrame] that contains the [GraphElement] with the given name.
*/
//go:nosplit
func (self class) GetElementFrame(element String.Name) [1]gdclass.GraphFrame { //gd:GraphEdit.get_element_frame
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(element)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_element_frame, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.GraphFrame{gd.PointerLifetimeBoundTo[gdclass.GraphFrame](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns an array of node names that are attached to the [GraphFrame] with the given name.
*/
//go:nosplit
func (self class) GetAttachedNodesOfFrame(frame_ String.Name) Array.Contains[String.Name] { //gd:GraphEdit.get_attached_nodes_of_frame
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(frame_)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_attached_nodes_of_frame, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[String.Name]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPanningScheme(scheme gdclass.GraphEditPanningScheme) { //gd:GraphEdit.set_panning_scheme
	var frame = callframe.New()
	callframe.Arg(frame, scheme)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_panning_scheme, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPanningScheme() gdclass.GraphEditPanningScheme { //gd:GraphEdit.get_panning_scheme
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.GraphEditPanningScheme](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_panning_scheme, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetZoom(zoom float64) { //gd:GraphEdit.set_zoom
	var frame = callframe.New()
	callframe.Arg(frame, zoom)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_zoom, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetZoom() float64 { //gd:GraphEdit.get_zoom
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_zoom, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetZoomMin(zoom_min float64) { //gd:GraphEdit.set_zoom_min
	var frame = callframe.New()
	callframe.Arg(frame, zoom_min)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_zoom_min, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetZoomMin() float64 { //gd:GraphEdit.get_zoom_min
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_zoom_min, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetZoomMax(zoom_max float64) { //gd:GraphEdit.set_zoom_max
	var frame = callframe.New()
	callframe.Arg(frame, zoom_max)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_zoom_max, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetZoomMax() float64 { //gd:GraphEdit.get_zoom_max
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_zoom_max, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetZoomStep(zoom_step float64) { //gd:GraphEdit.set_zoom_step
	var frame = callframe.New()
	callframe.Arg(frame, zoom_step)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_zoom_step, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetZoomStep() float64 { //gd:GraphEdit.get_zoom_step
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_zoom_step, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShowGrid(enable bool) { //gd:GraphEdit.set_show_grid
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_show_grid, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsShowingGrid() bool { //gd:GraphEdit.is_showing_grid
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_showing_grid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGridPattern(pattern gdclass.GraphEditGridPattern) { //gd:GraphEdit.set_grid_pattern
	var frame = callframe.New()
	callframe.Arg(frame, pattern)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_grid_pattern, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGridPattern() gdclass.GraphEditGridPattern { //gd:GraphEdit.get_grid_pattern
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.GraphEditGridPattern](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_grid_pattern, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSnappingEnabled(enable bool) { //gd:GraphEdit.set_snapping_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_snapping_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsSnappingEnabled() bool { //gd:GraphEdit.is_snapping_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_snapping_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSnappingDistance(pixels int64) { //gd:GraphEdit.set_snapping_distance
	var frame = callframe.New()
	callframe.Arg(frame, pixels)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_snapping_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSnappingDistance() int64 { //gd:GraphEdit.get_snapping_distance
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_snapping_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetConnectionLinesCurvature(curvature float64) { //gd:GraphEdit.set_connection_lines_curvature
	var frame = callframe.New()
	callframe.Arg(frame, curvature)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_connection_lines_curvature, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetConnectionLinesCurvature() float64 { //gd:GraphEdit.get_connection_lines_curvature
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_connection_lines_curvature, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetConnectionLinesThickness(pixels float64) { //gd:GraphEdit.set_connection_lines_thickness
	var frame = callframe.New()
	callframe.Arg(frame, pixels)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_connection_lines_thickness, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetConnectionLinesThickness() float64 { //gd:GraphEdit.get_connection_lines_thickness
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_connection_lines_thickness, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetConnectionLinesAntialiased(pixels bool) { //gd:GraphEdit.set_connection_lines_antialiased
	var frame = callframe.New()
	callframe.Arg(frame, pixels)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_connection_lines_antialiased, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsConnectionLinesAntialiased() bool { //gd:GraphEdit.is_connection_lines_antialiased
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_connection_lines_antialiased, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMinimapSize(size Vector2.XY) { //gd:GraphEdit.set_minimap_size
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_minimap_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMinimapSize() Vector2.XY { //gd:GraphEdit.get_minimap_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_minimap_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMinimapOpacity(opacity float64) { //gd:GraphEdit.set_minimap_opacity
	var frame = callframe.New()
	callframe.Arg(frame, opacity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_minimap_opacity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMinimapOpacity() float64 { //gd:GraphEdit.get_minimap_opacity
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_minimap_opacity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMinimapEnabled(enable bool) { //gd:GraphEdit.set_minimap_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_minimap_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsMinimapEnabled() bool { //gd:GraphEdit.is_minimap_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_minimap_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShowMenu(hidden bool) { //gd:GraphEdit.set_show_menu
	var frame = callframe.New()
	callframe.Arg(frame, hidden)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_show_menu, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsShowingMenu() bool { //gd:GraphEdit.is_showing_menu
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_showing_menu, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShowZoomLabel(enable bool) { //gd:GraphEdit.set_show_zoom_label
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_show_zoom_label, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsShowingZoomLabel() bool { //gd:GraphEdit.is_showing_zoom_label
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_showing_zoom_label, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShowGridButtons(hidden bool) { //gd:GraphEdit.set_show_grid_buttons
	var frame = callframe.New()
	callframe.Arg(frame, hidden)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_show_grid_buttons, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsShowingGridButtons() bool { //gd:GraphEdit.is_showing_grid_buttons
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_showing_grid_buttons, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShowZoomButtons(hidden bool) { //gd:GraphEdit.set_show_zoom_buttons
	var frame = callframe.New()
	callframe.Arg(frame, hidden)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_show_zoom_buttons, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsShowingZoomButtons() bool { //gd:GraphEdit.is_showing_zoom_buttons
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_showing_zoom_buttons, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShowMinimapButton(hidden bool) { //gd:GraphEdit.set_show_minimap_button
	var frame = callframe.New()
	callframe.Arg(frame, hidden)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_show_minimap_button, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsShowingMinimapButton() bool { //gd:GraphEdit.is_showing_minimap_button
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_showing_minimap_button, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShowArrangeButton(hidden bool) { //gd:GraphEdit.set_show_arrange_button
	var frame = callframe.New()
	callframe.Arg(frame, hidden)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_show_arrange_button, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsShowingArrangeButton() bool { //gd:GraphEdit.is_showing_arrange_button
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_showing_arrange_button, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRightDisconnects(enable bool) { //gd:GraphEdit.set_right_disconnects
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_right_disconnects, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsRightDisconnectsEnabled() bool { //gd:GraphEdit.is_right_disconnects_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_is_right_disconnects_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Gets the [HBoxContainer] that contains the zooming and grid snap controls in the top left of the graph. You can use this method to reposition the toolbar or to add your own custom controls to it.
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member CanvasItem.visible] property.
*/
//go:nosplit
func (self class) GetMenuHbox() [1]gdclass.HBoxContainer { //gd:GraphEdit.get_menu_hbox
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_get_menu_hbox, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.HBoxContainer{gd.PointerLifetimeBoundTo[gdclass.HBoxContainer](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Rearranges selected nodes in a layout with minimum crossings between connections and uniform horizontal and vertical gap between nodes.
*/
//go:nosplit
func (self class) ArrangeNodes() { //gd:GraphEdit.arrange_nodes
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_arrange_nodes, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the specified [param node] as the one selected.
*/
//go:nosplit
func (self class) SetSelected(node [1]gdclass.Node) { //gd:GraphEdit.set_selected
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(node[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphEdit.Bind_set_selected, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self Instance) OnConnectionRequest(cb func(from_node string, from_port int, to_node string, to_port int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("connection_request"), gd.NewCallable(cb), 0)
}

func (self Instance) OnDisconnectionRequest(cb func(from_node string, from_port int, to_node string, to_port int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("disconnection_request"), gd.NewCallable(cb), 0)
}

func (self Instance) OnConnectionToEmpty(cb func(from_node string, from_port int, release_position Vector2.XY)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("connection_to_empty"), gd.NewCallable(cb), 0)
}

func (self Instance) OnConnectionFromEmpty(cb func(to_node string, to_port int, release_position Vector2.XY)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("connection_from_empty"), gd.NewCallable(cb), 0)
}

func (self Instance) OnConnectionDragStarted(cb func(from_node string, from_port int, is_output bool)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("connection_drag_started"), gd.NewCallable(cb), 0)
}

func (self Instance) OnConnectionDragEnded(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("connection_drag_ended"), gd.NewCallable(cb), 0)
}

func (self Instance) OnCopyNodesRequest(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("copy_nodes_request"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPasteNodesRequest(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("paste_nodes_request"), gd.NewCallable(cb), 0)
}

func (self Instance) OnDuplicateNodesRequest(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("duplicate_nodes_request"), gd.NewCallable(cb), 0)
}

func (self Instance) OnDeleteNodesRequest(cb func(nodes []string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("delete_nodes_request"), gd.NewCallable(cb), 0)
}

func (self Instance) OnNodeSelected(cb func(node [1]gdclass.Node)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("node_selected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnNodeDeselected(cb func(node [1]gdclass.Node)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("node_deselected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnFrameRectChanged(cb func(frame_ [1]gdclass.GraphFrame, new_rect Vector2.XY)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("frame_rect_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPopupRequest(cb func(at_position Vector2.XY)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("popup_request"), gd.NewCallable(cb), 0)
}

func (self Instance) OnBeginNodeMove(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("begin_node_move"), gd.NewCallable(cb), 0)
}

func (self Instance) OnEndNodeMove(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("end_node_move"), gd.NewCallable(cb), 0)
}

func (self Instance) OnGraphElementsLinkedToFrameRequest(cb func(elements []any, frame_ string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("graph_elements_linked_to_frame_request"), gd.NewCallable(cb), 0)
}

func (self Instance) OnScrollOffsetChanged(cb func(offset Vector2.XY)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("scroll_offset_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsGraphEdit() Advanced       { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGraphEdit() Instance    { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.Advanced { return *((*Control.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsControl() Control.Instance {
	return *((*Control.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_is_in_input_hotzone":
		return reflect.ValueOf(self._is_in_input_hotzone)
	case "_is_in_output_hotzone":
		return reflect.ValueOf(self._is_in_output_hotzone)
	case "_get_connection_line":
		return reflect.ValueOf(self._get_connection_line)
	case "_is_node_hover_valid":
		return reflect.ValueOf(self._is_node_hover_valid)
	default:
		return gd.VirtualByName(Control.Advanced(self.AsControl()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_is_in_input_hotzone":
		return reflect.ValueOf(self._is_in_input_hotzone)
	case "_is_in_output_hotzone":
		return reflect.ValueOf(self._is_in_output_hotzone)
	case "_get_connection_line":
		return reflect.ValueOf(self._get_connection_line)
	case "_is_node_hover_valid":
		return reflect.ValueOf(self._is_node_hover_valid)
	default:
		return gd.VirtualByName(Control.Instance(self.AsControl()), name)
	}
}
func init() {
	gdclass.Register("GraphEdit", func(ptr gd.Object) any { return [1]gdclass.GraphEdit{*(*gdclass.GraphEdit)(unsafe.Pointer(&ptr))} })
}

type PanningScheme = gdclass.GraphEditPanningScheme //gd:GraphEdit.PanningScheme

const (
	/*[kbd]Mouse Wheel[/kbd] will zoom, [kbd]Ctrl + Mouse Wheel[/kbd] will move the view.*/
	ScrollZooms PanningScheme = 0
	/*[kbd]Mouse Wheel[/kbd] will move the view, [kbd]Ctrl + Mouse Wheel[/kbd] will zoom.*/
	ScrollPans PanningScheme = 1
)

type GridPattern = gdclass.GraphEditGridPattern //gd:GraphEdit.GridPattern

const (
	/*Draw the grid using solid lines.*/
	GridPatternLines GridPattern = 0
	/*Draw the grid using dots.*/
	GridPatternDots GridPattern = 1
)

type Connection struct {
	FromPort int    `gd:"from_port"`
	FromNode string `gd:"from_node"`
	ToPort   int    `gd:"to_port"`
	ToNode   string `gd:"to_node"`
}
