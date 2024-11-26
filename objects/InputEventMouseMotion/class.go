package InputEventMouseMotion

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/InputEventMouse"
import "grow.graphics/gd/objects/InputEventWithModifiers"
import "grow.graphics/gd/objects/InputEventFromWindow"
import "grow.graphics/gd/objects/InputEvent"
import "grow.graphics/gd/objects/Resource"
import "grow.graphics/gd/variant/Vector2"
import "grow.graphics/gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Stores information about a mouse or a pen motion. This includes relative position, absolute position, and velocity. See [method Node._input].
[b]Note:[/b] By default, this event is only emitted once per frame rendered at most. If you need more precise input reporting, set [member Input.use_accumulated_input] to [code]false[/code] to make events emitted as often as possible. If you use InputEventMouseMotion to draw lines, consider implementing [url=https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm]Bresenham's line algorithm[/url] as well to avoid visible gaps in lines if the user is moving the mouse quickly.
*/
type Instance [1]classdb.InputEventMouseMotion

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.InputEventMouseMotion

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("InputEventMouseMotion"))
	return Instance{classdb.InputEventMouseMotion(object)}
}

func (self Instance) Tilt() Vector2.XY {
	return Vector2.XY(class(self).GetTilt())
}

func (self Instance) SetTilt(value Vector2.XY) {
	class(self).SetTilt(gd.Vector2(value))
}

func (self Instance) Pressure() Float.X {
	return Float.X(Float.X(class(self).GetPressure()))
}

func (self Instance) SetPressure(value Float.X) {
	class(self).SetPressure(gd.Float(value))
}

func (self Instance) PenInverted() bool {
	return bool(class(self).GetPenInverted())
}

func (self Instance) SetPenInverted(value bool) {
	class(self).SetPenInverted(value)
}

func (self Instance) Relative() Vector2.XY {
	return Vector2.XY(class(self).GetRelative())
}

func (self Instance) SetRelative(value Vector2.XY) {
	class(self).SetRelative(gd.Vector2(value))
}

func (self Instance) ScreenRelative() Vector2.XY {
	return Vector2.XY(class(self).GetScreenRelative())
}

func (self Instance) SetScreenRelative(value Vector2.XY) {
	class(self).SetScreenRelative(gd.Vector2(value))
}

func (self Instance) Velocity() Vector2.XY {
	return Vector2.XY(class(self).GetVelocity())
}

func (self Instance) SetVelocity(value Vector2.XY) {
	class(self).SetVelocity(gd.Vector2(value))
}

func (self Instance) ScreenVelocity() Vector2.XY {
	return Vector2.XY(class(self).GetScreenVelocity())
}

func (self Instance) SetScreenVelocity(value Vector2.XY) {
	class(self).SetScreenVelocity(gd.Vector2(value))
}

//go:nosplit
func (self class) SetTilt(tilt gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, tilt)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseMotion.Bind_set_tilt, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTilt() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseMotion.Bind_get_tilt, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPressure(pressure gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, pressure)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseMotion.Bind_set_pressure, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPressure() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseMotion.Bind_get_pressure, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPenInverted(pen_inverted bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pen_inverted)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseMotion.Bind_set_pen_inverted, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPenInverted() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseMotion.Bind_get_pen_inverted, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRelative(relative gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, relative)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseMotion.Bind_set_relative, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRelative() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseMotion.Bind_get_relative, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetScreenRelative(relative gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, relative)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseMotion.Bind_set_screen_relative, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetScreenRelative() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseMotion.Bind_get_screen_relative, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVelocity(velocity gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, velocity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseMotion.Bind_set_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVelocity() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseMotion.Bind_get_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetScreenVelocity(velocity gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, velocity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseMotion.Bind_set_screen_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetScreenVelocity() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventMouseMotion.Bind_get_screen_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsInputEventMouseMotion() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsInputEventMouseMotion() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsInputEventMouse() InputEventMouse.Advanced {
	return *((*InputEventMouse.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsInputEventMouse() InputEventMouse.Instance {
	return *((*InputEventMouse.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsInputEventMouse(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsInputEventMouse(), name)
	}
}
func init() {
	classdb.Register("InputEventMouseMotion", func(ptr gd.Object) any { return classdb.InputEventMouseMotion(ptr) })
}
