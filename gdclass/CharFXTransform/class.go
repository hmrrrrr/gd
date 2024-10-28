package CharFXTransform

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
By setting various properties on this object, you can control how individual characters will be displayed in a [RichTextEffect].

*/
type Go [1]classdb.CharFXTransform
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.CharFXTransform
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CharFXTransform"))
	return Go{classdb.CharFXTransform(object)}
}

func (self Go) Transform() gd.Transform2D {
		return gd.Transform2D(class(self).GetTransform())
}

func (self Go) SetTransform(value gd.Transform2D) {
	class(self).SetTransform(value)
}

func (self Go) Range() gd.Vector2i {
		return gd.Vector2i(class(self).GetRange())
}

func (self Go) SetRange(value gd.Vector2i) {
	class(self).SetRange(value)
}

func (self Go) ElapsedTime() float64 {
		return float64(float64(class(self).GetElapsedTime()))
}

func (self Go) SetElapsedTime(value float64) {
	class(self).SetElapsedTime(gd.Float(value))
}

func (self Go) Visible() bool {
		return bool(class(self).IsVisible())
}

func (self Go) SetVisible(value bool) {
	class(self).SetVisibility(value)
}

func (self Go) Outline() bool {
		return bool(class(self).IsOutline())
}

func (self Go) SetOutline(value bool) {
	class(self).SetOutline(value)
}

func (self Go) Offset() gd.Vector2 {
		return gd.Vector2(class(self).GetOffset())
}

func (self Go) SetOffset(value gd.Vector2) {
	class(self).SetOffset(value)
}

func (self Go) Color() gd.Color {
		return gd.Color(class(self).GetColor())
}

func (self Go) SetColor(value gd.Color) {
	class(self).SetColor(value)
}

func (self Go) Env() gd.Dictionary {
		return gd.Dictionary(class(self).GetEnvironment())
}

func (self Go) SetEnv(value gd.Dictionary) {
	class(self).SetEnvironment(value)
}

func (self Go) GlyphIndex() int {
		return int(int(class(self).GetGlyphIndex()))
}

func (self Go) SetGlyphIndex(value int) {
	class(self).SetGlyphIndex(gd.Int(value))
}

func (self Go) GlyphCount() int {
		return int(int(class(self).GetGlyphCount()))
}

func (self Go) SetGlyphCount(value int) {
	class(self).SetGlyphCount(gd.Int(value))
}

func (self Go) GlyphFlags() int {
		return int(int(class(self).GetGlyphFlags()))
}

func (self Go) SetGlyphFlags(value int) {
	class(self).SetGlyphFlags(gd.Int(value))
}

func (self Go) RelativeIndex() int {
		return int(int(class(self).GetRelativeIndex()))
}

func (self Go) SetRelativeIndex(value int) {
	class(self).SetRelativeIndex(gd.Int(value))
}

func (self Go) Font() gd.RID {
		return gd.RID(class(self).GetFont())
}

func (self Go) SetFont(value gd.RID) {
	class(self).SetFont(value)
}

//go:nosplit
func (self class) GetTransform() gd.Transform2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTransform(transform gd.Transform2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRange() gd.Vector2i {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRange(arange gd.Vector2i)  {
	var frame = callframe.New()
	callframe.Arg(frame, arange)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetElapsedTime() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_elapsed_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetElapsedTime(time gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, time)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_elapsed_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsVisible() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_is_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVisibility(visibility bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, visibility)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_visibility, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsOutline() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_is_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOutline(outline bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, outline)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOffset() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOffset(offset gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetColor(color gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnvironment() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnvironment(environment gd.Dictionary)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(environment))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlyphIndex() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_glyph_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlyphIndex(glyph_index gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, glyph_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_glyph_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRelativeIndex() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_relative_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRelativeIndex(relative_index gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, relative_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_relative_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlyphCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_glyph_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlyphCount(glyph_count gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, glyph_count)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_glyph_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGlyphFlags() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_glyph_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetGlyphFlags(glyph_flags gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, glyph_flags)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_glyph_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFont() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFont(font gd.RID)  {
	var frame = callframe.New()
	callframe.Arg(frame, font)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsCharFXTransform() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsCharFXTransform() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("CharFXTransform", func(ptr gd.Object) any { return classdb.CharFXTransform(ptr) })}
