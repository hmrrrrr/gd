// Package CharFXTransform provides methods for working with CharFXTransform object instances.
package CharFXTransform

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Color"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Transform2D"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Vector2i"

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
By setting various properties on this object, you can control how individual characters will be displayed in a [RichTextEffect].
*/
type Instance [1]gdclass.CharFXTransform

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCharFXTransform() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.CharFXTransform

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CharFXTransform"))
	casted := Instance{*(*gdclass.CharFXTransform)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Transform() Transform2D.OriginXY {
	return Transform2D.OriginXY(class(self).GetTransform())
}

func (self Instance) SetTransform(value Transform2D.OriginXY) {
	class(self).SetTransform(Transform2D.OriginXY(value))
}

func (self Instance) Range() Vector2i.XY {
	return Vector2i.XY(class(self).GetRange())
}

func (self Instance) SetRange(value Vector2i.XY) {
	class(self).SetRange(Vector2i.XY(value))
}

func (self Instance) ElapsedTime() Float.X {
	return Float.X(Float.X(class(self).GetElapsedTime()))
}

func (self Instance) SetElapsedTime(value Float.X) {
	class(self).SetElapsedTime(float64(value))
}

func (self Instance) Visible() bool {
	return bool(class(self).IsVisible())
}

func (self Instance) SetVisible(value bool) {
	class(self).SetVisibility(value)
}

func (self Instance) Outline() bool {
	return bool(class(self).IsOutline())
}

func (self Instance) SetOutline(value bool) {
	class(self).SetOutline(value)
}

func (self Instance) Offset() Vector2.XY {
	return Vector2.XY(class(self).GetOffset())
}

func (self Instance) SetOffset(value Vector2.XY) {
	class(self).SetOffset(Vector2.XY(value))
}

func (self Instance) Color() Color.RGBA {
	return Color.RGBA(class(self).GetColor())
}

func (self Instance) SetColor(value Color.RGBA) {
	class(self).SetColor(Color.RGBA(value))
}

func (self Instance) Env() map[any]any {
	return map[any]any(gd.DictionaryAs[map[any]any](class(self).GetEnvironment()))
}

func (self Instance) SetEnv(value map[any]any) {
	class(self).SetEnvironment(gd.DictionaryFromMap(value))
}

func (self Instance) GlyphIndex() int {
	return int(int(class(self).GetGlyphIndex()))
}

func (self Instance) SetGlyphIndex(value int) {
	class(self).SetGlyphIndex(int64(value))
}

func (self Instance) GlyphCount() int {
	return int(int(class(self).GetGlyphCount()))
}

func (self Instance) SetGlyphCount(value int) {
	class(self).SetGlyphCount(int64(value))
}

func (self Instance) GlyphFlags() int {
	return int(int(class(self).GetGlyphFlags()))
}

func (self Instance) SetGlyphFlags(value int) {
	class(self).SetGlyphFlags(int64(value))
}

func (self Instance) RelativeIndex() int {
	return int(int(class(self).GetRelativeIndex()))
}

func (self Instance) SetRelativeIndex(value int) {
	class(self).SetRelativeIndex(int64(value))
}

func (self Instance) Font() RID.Any {
	return RID.Any(class(self).GetFont())
}

func (self Instance) SetFont(value RID.Any) {
	class(self).SetFont(RID.Any(value))
}

//go:nosplit
func (self class) GetTransform() Transform2D.OriginXY { //gd:CharFXTransform.get_transform
	var frame = callframe.New()
	var r_ret = callframe.Ret[Transform2D.OriginXY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTransform(transform Transform2D.OriginXY) { //gd:CharFXTransform.set_transform
	var frame = callframe.New()
	callframe.Arg(frame, transform)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRange() Vector2i.XY { //gd:CharFXTransform.get_range
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2i.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRange(arange Vector2i.XY) { //gd:CharFXTransform.set_range
	var frame = callframe.New()
	callframe.Arg(frame, arange)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetElapsedTime() float64 { //gd:CharFXTransform.get_elapsed_time
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_elapsed_time, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetElapsedTime(time float64) { //gd:CharFXTransform.set_elapsed_time
	var frame = callframe.New()
	callframe.Arg(frame, time)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_elapsed_time, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsVisible() bool { //gd:CharFXTransform.is_visible
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_is_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisibility(visibility bool) { //gd:CharFXTransform.set_visibility
	var frame = callframe.New()
	callframe.Arg(frame, visibility)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_visibility, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsOutline() bool { //gd:CharFXTransform.is_outline
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_is_outline, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOutline(outline bool) { //gd:CharFXTransform.set_outline
	var frame = callframe.New()
	callframe.Arg(frame, outline)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_outline, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOffset() Vector2.XY { //gd:CharFXTransform.get_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOffset(offset Vector2.XY) { //gd:CharFXTransform.set_offset
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetColor() Color.RGBA { //gd:CharFXTransform.get_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetColor(color Color.RGBA) { //gd:CharFXTransform.set_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnvironment() Dictionary.Any { //gd:CharFXTransform.get_environment
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_environment, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnvironment(environment Dictionary.Any) { //gd:CharFXTransform.set_environment
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalDictionary(environment)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_environment, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlyphIndex() int64 { //gd:CharFXTransform.get_glyph_index
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_glyph_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlyphIndex(glyph_index int64) { //gd:CharFXTransform.set_glyph_index
	var frame = callframe.New()
	callframe.Arg(frame, glyph_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_glyph_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRelativeIndex() int64 { //gd:CharFXTransform.get_relative_index
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_relative_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRelativeIndex(relative_index int64) { //gd:CharFXTransform.set_relative_index
	var frame = callframe.New()
	callframe.Arg(frame, relative_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_relative_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlyphCount() int64 { //gd:CharFXTransform.get_glyph_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_glyph_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlyphCount(glyph_count int64) { //gd:CharFXTransform.set_glyph_count
	var frame = callframe.New()
	callframe.Arg(frame, glyph_count)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_glyph_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlyphFlags() int64 { //gd:CharFXTransform.get_glyph_flags
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_glyph_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlyphFlags(glyph_flags int64) { //gd:CharFXTransform.set_glyph_flags
	var frame = callframe.New()
	callframe.Arg(frame, glyph_flags)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_glyph_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFont() RID.Any { //gd:CharFXTransform.get_font
	var frame = callframe.New()
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_get_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFont(font RID.Any) { //gd:CharFXTransform.set_font
	var frame = callframe.New()
	callframe.Arg(frame, font)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CharFXTransform.Bind_set_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsCharFXTransform() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCharFXTransform() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("CharFXTransform", func(ptr gd.Object) any {
		return [1]gdclass.CharFXTransform{*(*gdclass.CharFXTransform)(unsafe.Pointer(&ptr))}
	})
}
