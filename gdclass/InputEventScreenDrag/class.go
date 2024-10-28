package InputEventScreenDrag

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/InputEventFromWindow"
import "grow.graphics/gd/gdclass/InputEvent"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Stores information about screen drag events. See [method Node._input].

*/
type Go [1]classdb.InputEventScreenDrag
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.InputEventScreenDrag
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("InputEventScreenDrag"))
	return Go{classdb.InputEventScreenDrag(object)}
}

func (self Go) Index() int {
		return int(int(class(self).GetIndex()))
}

func (self Go) SetIndex(value int) {
	class(self).SetIndex(gd.Int(value))
}

func (self Go) Tilt() gd.Vector2 {
		return gd.Vector2(class(self).GetTilt())
}

func (self Go) SetTilt(value gd.Vector2) {
	class(self).SetTilt(value)
}

func (self Go) Pressure() float64 {
		return float64(float64(class(self).GetPressure()))
}

func (self Go) SetPressure(value float64) {
	class(self).SetPressure(gd.Float(value))
}

func (self Go) PenInverted() bool {
		return bool(class(self).GetPenInverted())
}

func (self Go) SetPenInverted(value bool) {
	class(self).SetPenInverted(value)
}

func (self Go) Position() gd.Vector2 {
		return gd.Vector2(class(self).GetPosition())
}

func (self Go) SetPosition(value gd.Vector2) {
	class(self).SetPosition(value)
}

func (self Go) Relative() gd.Vector2 {
		return gd.Vector2(class(self).GetRelative())
}

func (self Go) SetRelative(value gd.Vector2) {
	class(self).SetRelative(value)
}

func (self Go) ScreenRelative() gd.Vector2 {
		return gd.Vector2(class(self).GetScreenRelative())
}

func (self Go) SetScreenRelative(value gd.Vector2) {
	class(self).SetScreenRelative(value)
}

func (self Go) Velocity() gd.Vector2 {
		return gd.Vector2(class(self).GetVelocity())
}

func (self Go) SetVelocity(value gd.Vector2) {
	class(self).SetVelocity(value)
}

func (self Go) ScreenVelocity() gd.Vector2 {
		return gd.Vector2(class(self).GetScreenVelocity())
}

func (self Go) SetScreenVelocity(value gd.Vector2) {
	class(self).SetScreenVelocity(value)
}

//go:nosplit
func (self class) SetIndex(index gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_set_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetIndex() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_get_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTilt(tilt gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, tilt)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_set_tilt, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTilt() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_get_tilt, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPressure(pressure gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, pressure)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_set_pressure, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPressure() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_get_pressure, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPenInverted(pen_inverted bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, pen_inverted)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_set_pen_inverted, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPenInverted() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_get_pen_inverted, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPosition(position gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_set_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPosition() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_get_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRelative(relative gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, relative)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_set_relative, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRelative() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_get_relative, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetScreenRelative(relative gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, relative)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_set_screen_relative, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetScreenRelative() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_get_screen_relative, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVelocity(velocity gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, velocity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_set_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVelocity() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_get_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetScreenVelocity(velocity gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, velocity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_set_screen_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetScreenVelocity() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.InputEventScreenDrag.Bind_get_screen_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsInputEventScreenDrag() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsInputEventScreenDrag() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsInputEventFromWindow() InputEventFromWindow.GD { return *((*InputEventFromWindow.GD)(unsafe.Pointer(&self))) }
func (self Go) AsInputEventFromWindow() InputEventFromWindow.Go { return *((*InputEventFromWindow.Go)(unsafe.Pointer(&self))) }
func (self class) AsInputEvent() InputEvent.GD { return *((*InputEvent.GD)(unsafe.Pointer(&self))) }
func (self Go) AsInputEvent() InputEvent.Go { return *((*InputEvent.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsInputEventFromWindow(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsInputEventFromWindow(), name)
	}
}
func init() {classdb.Register("InputEventScreenDrag", func(ptr gd.Object) any { return classdb.InputEventScreenDrag(ptr) })}
