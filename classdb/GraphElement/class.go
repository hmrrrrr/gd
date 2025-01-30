// Package GraphElement provides methods for working with GraphElement object instances.
package GraphElement

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Container"
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
[GraphElement] allows to create custom elements for a [GraphEdit] graph. By default such elements can be selected, resized, and repositioned, but they cannot be connected. For a graph element that allows for connections see [GraphNode].
*/
type Instance [1]gdclass.GraphElement

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGraphElement() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GraphElement

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GraphElement"))
	casted := Instance{*(*gdclass.GraphElement)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) PositionOffset() Vector2.XY {
	return Vector2.XY(class(self).GetPositionOffset())
}

func (self Instance) SetPositionOffset(value Vector2.XY) {
	class(self).SetPositionOffset(Vector2.XY(value))
}

func (self Instance) Resizable() bool {
	return bool(class(self).IsResizable())
}

func (self Instance) SetResizable(value bool) {
	class(self).SetResizable(value)
}

func (self Instance) Draggable() bool {
	return bool(class(self).IsDraggable())
}

func (self Instance) SetDraggable(value bool) {
	class(self).SetDraggable(value)
}

func (self Instance) Selectable() bool {
	return bool(class(self).IsSelectable())
}

func (self Instance) SetSelectable(value bool) {
	class(self).SetSelectable(value)
}

func (self Instance) Selected() bool {
	return bool(class(self).IsSelected())
}

func (self Instance) SetSelected(value bool) {
	class(self).SetSelected(value)
}

//go:nosplit
func (self class) SetResizable(resizable bool) { //gd:GraphElement.set_resizable
	var frame = callframe.New()
	callframe.Arg(frame, resizable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphElement.Bind_set_resizable, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsResizable() bool { //gd:GraphElement.is_resizable
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphElement.Bind_is_resizable, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDraggable(draggable bool) { //gd:GraphElement.set_draggable
	var frame = callframe.New()
	callframe.Arg(frame, draggable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphElement.Bind_set_draggable, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsDraggable() bool { //gd:GraphElement.is_draggable
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphElement.Bind_is_draggable, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSelectable(selectable bool) { //gd:GraphElement.set_selectable
	var frame = callframe.New()
	callframe.Arg(frame, selectable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphElement.Bind_set_selectable, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsSelectable() bool { //gd:GraphElement.is_selectable
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphElement.Bind_is_selectable, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSelected(selected bool) { //gd:GraphElement.set_selected
	var frame = callframe.New()
	callframe.Arg(frame, selected)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphElement.Bind_set_selected, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsSelected() bool { //gd:GraphElement.is_selected
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphElement.Bind_is_selected, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPositionOffset(offset Vector2.XY) { //gd:GraphElement.set_position_offset
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphElement.Bind_set_position_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPositionOffset() Vector2.XY { //gd:GraphElement.get_position_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GraphElement.Bind_get_position_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnNodeSelected(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("node_selected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnNodeDeselected(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("node_deselected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnRaiseRequest(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("raise_request"), gd.NewCallable(cb), 0)
}

func (self Instance) OnDeleteRequest(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("delete_request"), gd.NewCallable(cb), 0)
}

func (self Instance) OnResizeRequest(cb func(new_size Vector2.XY)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("resize_request"), gd.NewCallable(cb), 0)
}

func (self Instance) OnResizeEnd(cb func(new_size Vector2.XY)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("resize_end"), gd.NewCallable(cb), 0)
}

func (self Instance) OnDragged(cb func(from Vector2.XY, to Vector2.XY)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("dragged"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPositionOffsetChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("position_offset_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsGraphElement() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGraphElement() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsContainer() Container.Advanced {
	return *((*Container.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsContainer() Container.Instance {
	return *((*Container.Instance)(unsafe.Pointer(&self)))
}
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
	default:
		return gd.VirtualByName(Container.Advanced(self.AsContainer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Container.Instance(self.AsContainer()), name)
	}
}
func init() {
	gdclass.Register("GraphElement", func(ptr gd.Object) any {
		return [1]gdclass.GraphElement{*(*gdclass.GraphElement)(unsafe.Pointer(&ptr))}
	})
}
