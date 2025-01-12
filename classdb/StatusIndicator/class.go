// Package StatusIndicator provides methods for working with StatusIndicator object instances.
package StatusIndicator

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/NodePath"
import "graphics.gd/variant/Rect2"
import "graphics.gd/variant/Vector2i"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type Instance [1]gdclass.StatusIndicator

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsStatusIndicator() Instance
}

/*
Returns the status indicator rectangle in screen coordinates. If this status indicator is not visible, returns an empty [Rect2].
*/
func (self Instance) GetRect() Rect2.PositionSize {
	return Rect2.PositionSize(class(self).GetRect())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.StatusIndicator

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("StatusIndicator"))
	casted := Instance{*(*gdclass.StatusIndicator)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Tooltip() string {
	return string(class(self).GetTooltip().String())
}

func (self Instance) SetTooltip(value string) {
	class(self).SetTooltip(gd.NewString(value))
}

func (self Instance) Icon() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetIcon())
}

func (self Instance) SetIcon(value [1]gdclass.Texture2D) {
	class(self).SetIcon(value)
}

func (self Instance) Menu() NodePath.String {
	return NodePath.String(class(self).GetMenu().String())
}

func (self Instance) SetMenu(value NodePath.String) {
	class(self).SetMenu(gd.NewString(string(value)).NodePath())
}

func (self Instance) Visible() bool {
	return bool(class(self).IsVisible())
}

func (self Instance) SetVisible(value bool) {
	class(self).SetVisible(value)
}

//go:nosplit
func (self class) SetTooltip(tooltip gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(tooltip))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StatusIndicator.Bind_set_tooltip, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTooltip() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StatusIndicator.Bind_get_tooltip, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetIcon(texture [1]gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StatusIndicator.Bind_set_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetIcon() [1]gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StatusIndicator.Bind_get_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisible(visible bool) {
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StatusIndicator.Bind_set_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsVisible() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StatusIndicator.Bind_is_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMenu(menu gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(menu))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StatusIndicator.Bind_set_menu, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMenu() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StatusIndicator.Bind_get_menu, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the status indicator rectangle in screen coordinates. If this status indicator is not visible, returns an empty [Rect2].
*/
//go:nosplit
func (self class) GetRect() gd.Rect2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StatusIndicator.Bind_get_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnPressed(cb func(mouse_button int, mouse_position Vector2i.XY)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("pressed"), gd.NewCallable(cb), 0)
}

func (self class) AsStatusIndicator() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsStatusIndicator() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced          { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance       { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node.Advanced(self.AsNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node.Instance(self.AsNode()), name)
	}
}
func init() {
	gdclass.Register("StatusIndicator", func(ptr gd.Object) any {
		return [1]gdclass.StatusIndicator{*(*gdclass.StatusIndicator)(unsafe.Pointer(&ptr))}
	})
}
