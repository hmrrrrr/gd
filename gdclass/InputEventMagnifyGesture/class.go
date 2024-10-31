package InputEventMagnifyGesture

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/InputEventGesture"
import "grow.graphics/gd/gdclass/InputEventWithModifiers"
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
Stores the factor of a magnifying touch gesture. This is usually performed when the user pinches the touch screen and used for zooming in/out.
[b]Note:[/b] On Android, this requires the [member ProjectSettings.input_devices/pointing/android/enable_pan_and_scale_gestures] project setting to be enabled.
*/
type Instance [1]classdb.InputEventMagnifyGesture

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.InputEventMagnifyGesture

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("InputEventMagnifyGesture"))
	return Instance{classdb.InputEventMagnifyGesture(object)}
}

func (self Instance) Factor() float64 {
	return float64(float64(class(self).GetFactor()))
}

func (self Instance) SetFactor(value float64) {
	class(self).SetFactor(gd.Float(value))
}

//go:nosplit
func (self class) SetFactor(factor gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, factor)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMagnifyGesture.Bind_set_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFactor() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMagnifyGesture.Bind_get_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsInputEventMagnifyGesture() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsInputEventMagnifyGesture() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsInputEventGesture() InputEventGesture.Advanced {
	return *((*InputEventGesture.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsInputEventGesture() InputEventGesture.Instance {
	return *((*InputEventGesture.Instance)(unsafe.Pointer(&self)))
}
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsInputEventGesture(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsInputEventGesture(), name)
	}
}
func init() {
	classdb.Register("InputEventMagnifyGesture", func(ptr gd.Object) any { return classdb.InputEventMagnifyGesture(ptr) })
}
