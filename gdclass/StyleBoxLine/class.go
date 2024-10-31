package StyleBoxLine

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/StyleBox"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
A [StyleBox] that displays a single line of a given color and thickness. The line can be either horizontal or vertical. Useful for separators.
*/
type Instance [1]classdb.StyleBoxLine

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.StyleBoxLine

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("StyleBoxLine"))
	return Instance{classdb.StyleBoxLine(object)}
}

func (self Instance) Color() gd.Color {
	return gd.Color(class(self).GetColor())
}

func (self Instance) SetColor(value gd.Color) {
	class(self).SetColor(value)
}

func (self Instance) GrowBegin() float64 {
	return float64(float64(class(self).GetGrowBegin()))
}

func (self Instance) SetGrowBegin(value float64) {
	class(self).SetGrowBegin(gd.Float(value))
}

func (self Instance) GrowEnd() float64 {
	return float64(float64(class(self).GetGrowEnd()))
}

func (self Instance) SetGrowEnd(value float64) {
	class(self).SetGrowEnd(gd.Float(value))
}

func (self Instance) Thickness() int {
	return int(int(class(self).GetThickness()))
}

func (self Instance) SetThickness(value int) {
	class(self).SetThickness(gd.Int(value))
}

func (self Instance) Vertical() bool {
	return bool(class(self).IsVertical())
}

func (self Instance) SetVertical(value bool) {
	class(self).SetVertical(value)
}

//go:nosplit
func (self class) SetColor(color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxLine.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxLine.Bind_get_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetThickness(thickness gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, thickness)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxLine.Bind_set_thickness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetThickness() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxLine.Bind_get_thickness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGrowBegin(offset gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxLine.Bind_set_grow_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGrowBegin() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxLine.Bind_get_grow_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGrowEnd(offset gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxLine.Bind_set_grow_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGrowEnd() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxLine.Bind_get_grow_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVertical(vertical bool) {
	var frame = callframe.New()
	callframe.Arg(frame, vertical)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxLine.Bind_set_vertical, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsVertical() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StyleBoxLine.Bind_is_vertical, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsStyleBoxLine() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsStyleBoxLine() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsStyleBox() StyleBox.Advanced {
	return *((*StyleBox.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsStyleBox() StyleBox.Instance {
	return *((*StyleBox.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsStyleBox(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsStyleBox(), name)
	}
}
func init() {
	classdb.Register("StyleBoxLine", func(ptr gd.Object) any { return classdb.StyleBoxLine(ptr) })
}
