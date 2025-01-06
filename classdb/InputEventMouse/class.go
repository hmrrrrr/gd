// Package InputEventMouse provides methods for working with InputEventMouse object instances.
package InputEventMouse

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/InputEventWithModifiers"
import "graphics.gd/classdb/InputEventFromWindow"
import "graphics.gd/classdb/InputEvent"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Vector2"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Stores general information about mouse events.
*/
type Instance [1]gdclass.InputEventMouse
type Any interface {
	gd.IsClass
	AsInputEventMouse() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.InputEventMouse

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("InputEventMouse"))
	casted := Instance{*(*gdclass.InputEventMouse)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) ButtonMask() MouseButtonMask {
	return MouseButtonMask(class(self).GetButtonMask())
}

func (self Instance) SetButtonMask(value MouseButtonMask) {
	class(self).SetButtonMask(value)
}

func (self Instance) Position() Vector2.XY {
	return Vector2.XY(class(self).GetPosition())
}

func (self Instance) SetPosition(value Vector2.XY) {
	class(self).SetPosition(gd.Vector2(value))
}

func (self Instance) GlobalPosition() Vector2.XY {
	return Vector2.XY(class(self).GetGlobalPosition())
}

func (self Instance) SetGlobalPosition(value Vector2.XY) {
	class(self).SetGlobalPosition(gd.Vector2(value))
}

//go:nosplit
func (self class) SetButtonMask(button_mask MouseButtonMask) {
	var frame = callframe.New()
	callframe.Arg(frame, button_mask)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouse.Bind_set_button_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetButtonMask() MouseButtonMask {
	var frame = callframe.New()
	var r_ret = callframe.Ret[MouseButtonMask](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouse.Bind_get_button_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPosition(position gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouse.Bind_set_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPosition() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouse.Bind_get_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlobalPosition(global_position gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, global_position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouse.Bind_set_global_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlobalPosition() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouse.Bind_get_global_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsInputEventMouse() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsInputEventMouse() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsInputEventWithModifiers() InputEventWithModifiers.Advanced {
	return *((*InputEventWithModifiers.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsInputEventWithModifiers() InputEventWithModifiers.Instance {
	return *((*InputEventWithModifiers.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsInputEventFromWindow() InputEventFromWindow.Advanced {
	return *((*InputEventFromWindow.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsInputEventFromWindow() InputEventFromWindow.Instance {
	return *((*InputEventFromWindow.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsInputEvent() InputEvent.Advanced {
	return *((*InputEvent.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsInputEvent() InputEvent.Instance {
	return *((*InputEvent.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(InputEventWithModifiers.Advanced(self.AsInputEventWithModifiers()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(InputEventWithModifiers.Instance(self.AsInputEventWithModifiers()), name)
	}
}
func init() {
	gdclass.Register("InputEventMouse", func(ptr gd.Object) any {
		return [1]gdclass.InputEventMouse{*(*gdclass.InputEventMouse)(unsafe.Pointer(&ptr))}
	})
}

type MouseButtonMask int

const (
	/*Primary mouse button mask, usually for the left button.*/
	MouseButtonMaskLeft MouseButtonMask = 1
	/*Secondary mouse button mask, usually for the right button.*/
	MouseButtonMaskRight MouseButtonMask = 2
	/*Middle mouse button mask.*/
	MouseButtonMaskMiddle MouseButtonMask = 4
	/*Extra mouse button 1 mask.*/
	MouseButtonMaskMbXbutton1 MouseButtonMask = 128
	/*Extra mouse button 2 mask.*/
	MouseButtonMaskMbXbutton2 MouseButtonMask = 256
)
