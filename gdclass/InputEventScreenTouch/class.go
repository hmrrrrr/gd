package InputEventScreenTouch

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/InputEventFromWindow"
import "grow.graphics/gd/gdclass/InputEvent"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
Stores information about multi-touch press/release input events. Supports touch press, touch release and [member index] for multi-touch count and order.
*/
type Instance [1]classdb.InputEventScreenTouch

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.InputEventScreenTouch

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("InputEventScreenTouch"))
	return Instance{classdb.InputEventScreenTouch(object)}
}

func (self Instance) Index() int {
	return int(int(class(self).GetIndex()))
}

func (self Instance) SetIndex(value int) {
	class(self).SetIndex(gd.Int(value))
}

func (self Instance) Position() gd.Vector2 {
	return gd.Vector2(class(self).GetPosition())
}

func (self Instance) SetPosition(value gd.Vector2) {
	class(self).SetPosition(value)
}

func (self Instance) DoubleTap() bool {
	return bool(class(self).IsDoubleTap())
}

func (self Instance) SetDoubleTap(value bool) {
	class(self).SetDoubleTap(value)
}

//go:nosplit
func (self class) SetIndex(index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenTouch.Bind_set_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetIndex() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenTouch.Bind_get_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPosition(position gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenTouch.Bind_set_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPosition() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenTouch.Bind_get_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPressed(pressed bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pressed)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenTouch.Bind_set_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetCanceled(canceled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, canceled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenTouch.Bind_set_canceled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetDoubleTap(double_tap bool) {
	var frame = callframe.New()
	callframe.Arg(frame, double_tap)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenTouch.Bind_set_double_tap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsDoubleTap() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenTouch.Bind_is_double_tap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsInputEventScreenTouch() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsInputEventScreenTouch() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsInputEventFromWindow(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsInputEventFromWindow(), name)
	}
}
func init() {
	classdb.Register("InputEventScreenTouch", func(ptr gd.Object) any { return classdb.InputEventScreenTouch(ptr) })
}
