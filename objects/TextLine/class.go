package TextLine

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/variant/Array"
import "grow.graphics/gd/variant/Vector2"
import "grow.graphics/gd/variant/Float"
import "grow.graphics/gd/variant/Rect2"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Abstraction over [TextServer] for handling a single line of text.
*/
type Instance [1]classdb.TextLine

/*
Clears text line (removes text and inline objects).
*/
func (self Instance) Clear() {
	class(self).Clear()
}

/*
Overrides BiDi for the structured text.
Override ranges should cover full source text without overlaps. BiDi algorithm will be used on each range separately.
*/
func (self Instance) SetBidiOverride(override Array.Any) {
	class(self).SetBidiOverride(override)
}

/*
Adds text span and font to draw it.
*/
func (self Instance) AddString(text string, font objects.Font, font_size int) bool {
	return bool(class(self).AddString(gd.NewString(text), font, gd.Int(font_size), gd.NewString(""), gd.NewVariant(gd.NewVariant(([1]any{}[0])))))
}

/*
Adds inline object to the text buffer, [param key] must be unique. In the text, object is represented as [param length] object replacement characters.
*/
func (self Instance) AddObject(key any, size Vector2.XY) bool {
	return bool(class(self).AddObject(gd.NewVariant(key), gd.Vector2(size), 5, gd.Int(1), gd.Float(0.0)))
}

/*
Sets new size and alignment of embedded object.
*/
func (self Instance) ResizeObject(key any, size Vector2.XY) bool {
	return bool(class(self).ResizeObject(gd.NewVariant(key), gd.Vector2(size), 5, gd.Float(0.0)))
}

/*
Aligns text to the given tab-stops.
*/
func (self Instance) TabAlign(tab_stops []float32) {
	class(self).TabAlign(gd.NewPackedFloat32Slice(tab_stops))
}

/*
Returns array of inline objects.
*/
func (self Instance) GetObjects() Array.Any {
	return Array.Any(class(self).GetObjects())
}

/*
Returns bounding rectangle of the inline object.
*/
func (self Instance) GetObjectRect(key any) Rect2.PositionSize {
	return Rect2.PositionSize(class(self).GetObjectRect(gd.NewVariant(key)))
}

/*
Returns size of the bounding box of the text.
*/
func (self Instance) GetSize() Vector2.XY {
	return Vector2.XY(class(self).GetSize())
}

/*
Returns TextServer buffer RID.
*/
func (self Instance) GetRid() Resource.ID {
	return Resource.ID(class(self).GetRid())
}

/*
Returns the text ascent (number of pixels above the baseline for horizontal layout or to the left of baseline for vertical).
*/
func (self Instance) GetLineAscent() Float.X {
	return Float.X(Float.X(class(self).GetLineAscent()))
}

/*
Returns the text descent (number of pixels below the baseline for horizontal layout or to the right of baseline for vertical).
*/
func (self Instance) GetLineDescent() Float.X {
	return Float.X(Float.X(class(self).GetLineDescent()))
}

/*
Returns width (for horizontal layout) or height (for vertical) of the text.
*/
func (self Instance) GetLineWidth() Float.X {
	return Float.X(Float.X(class(self).GetLineWidth()))
}

/*
Returns pixel offset of the underline below the baseline.
*/
func (self Instance) GetLineUnderlinePosition() Float.X {
	return Float.X(Float.X(class(self).GetLineUnderlinePosition()))
}

/*
Returns thickness of the underline.
*/
func (self Instance) GetLineUnderlineThickness() Float.X {
	return Float.X(Float.X(class(self).GetLineUnderlineThickness()))
}

/*
Draw text into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
func (self Instance) Draw(canvas Resource.ID, pos Vector2.XY) {
	class(self).Draw(canvas, gd.Vector2(pos), gd.Color(gd.Color{1, 1, 1, 1}))
}

/*
Draw text into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
func (self Instance) DrawOutline(canvas Resource.ID, pos Vector2.XY) {
	class(self).DrawOutline(canvas, gd.Vector2(pos), gd.Int(1), gd.Color(gd.Color{1, 1, 1, 1}))
}

/*
Returns caret character offset at the specified pixel offset at the baseline. This function always returns a valid position.
*/
func (self Instance) HitTest(coords Float.X) int {
	return int(int(class(self).HitTest(gd.Float(coords))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.TextLine

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TextLine"))
	return Instance{classdb.TextLine(object)}
}

func (self Instance) Direction() classdb.TextServerDirection {
	return classdb.TextServerDirection(class(self).GetDirection())
}

func (self Instance) SetDirection(value classdb.TextServerDirection) {
	class(self).SetDirection(value)
}

func (self Instance) Orientation() classdb.TextServerOrientation {
	return classdb.TextServerOrientation(class(self).GetOrientation())
}

func (self Instance) SetOrientation(value classdb.TextServerOrientation) {
	class(self).SetOrientation(value)
}

func (self Instance) PreserveInvalid() bool {
	return bool(class(self).GetPreserveInvalid())
}

func (self Instance) SetPreserveInvalid(value bool) {
	class(self).SetPreserveInvalid(value)
}

func (self Instance) PreserveControl() bool {
	return bool(class(self).GetPreserveControl())
}

func (self Instance) SetPreserveControl(value bool) {
	class(self).SetPreserveControl(value)
}

func (self Instance) Width() Float.X {
	return Float.X(Float.X(class(self).GetWidth()))
}

func (self Instance) SetWidth(value Float.X) {
	class(self).SetWidth(gd.Float(value))
}

func (self Instance) Alignment() HorizontalAlignment {
	return HorizontalAlignment(class(self).GetHorizontalAlignment())
}

func (self Instance) SetAlignment(value HorizontalAlignment) {
	class(self).SetHorizontalAlignment(value)
}

func (self Instance) Flags() classdb.TextServerJustificationFlag {
	return classdb.TextServerJustificationFlag(class(self).GetFlags())
}

func (self Instance) SetFlags(value classdb.TextServerJustificationFlag) {
	class(self).SetFlags(value)
}

func (self Instance) TextOverrunBehavior() classdb.TextServerOverrunBehavior {
	return classdb.TextServerOverrunBehavior(class(self).GetTextOverrunBehavior())
}

func (self Instance) SetTextOverrunBehavior(value classdb.TextServerOverrunBehavior) {
	class(self).SetTextOverrunBehavior(value)
}

func (self Instance) EllipsisChar() string {
	return string(class(self).GetEllipsisChar().String())
}

func (self Instance) SetEllipsisChar(value string) {
	class(self).SetEllipsisChar(gd.NewString(value))
}

/*
Clears text line (removes text and inline objects).
*/
//go:nosplit
func (self class) Clear() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetDirection(direction classdb.TextServerDirection) {
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_set_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDirection() classdb.TextServerDirection {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerDirection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOrientation(orientation classdb.TextServerOrientation) {
	var frame = callframe.New()
	callframe.Arg(frame, orientation)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_set_orientation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetOrientation() classdb.TextServerOrientation {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerOrientation](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_orientation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPreserveInvalid(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_set_preserve_invalid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPreserveInvalid() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_preserve_invalid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPreserveControl(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_set_preserve_control, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPreserveControl() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_preserve_control, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Overrides BiDi for the structured text.
Override ranges should cover full source text without overlaps. BiDi algorithm will be used on each range separately.
*/
//go:nosplit
func (self class) SetBidiOverride(override gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(override))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_set_bidi_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds text span and font to draw it.
*/
//go:nosplit
func (self class) AddString(text gd.String, font objects.Font, font_size gd.Int, language gd.String, meta gd.Variant) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	callframe.Arg(frame, pointers.Get(font[0])[0])
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, pointers.Get(language))
	callframe.Arg(frame, pointers.Get(meta))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_add_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds inline object to the text buffer, [param key] must be unique. In the text, object is represented as [param length] object replacement characters.
*/
//go:nosplit
func (self class) AddObject(key gd.Variant, size gd.Vector2, inline_align InlineAlignment, length gd.Int, baseline gd.Float) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(key))
	callframe.Arg(frame, size)
	callframe.Arg(frame, inline_align)
	callframe.Arg(frame, length)
	callframe.Arg(frame, baseline)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_add_object, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets new size and alignment of embedded object.
*/
//go:nosplit
func (self class) ResizeObject(key gd.Variant, size gd.Vector2, inline_align InlineAlignment, baseline gd.Float) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(key))
	callframe.Arg(frame, size)
	callframe.Arg(frame, inline_align)
	callframe.Arg(frame, baseline)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_resize_object, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWidth(width gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_set_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetWidth() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHorizontalAlignment(alignment HorizontalAlignment) {
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_set_horizontal_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetHorizontalAlignment() HorizontalAlignment {
	var frame = callframe.New()
	var r_ret = callframe.Ret[HorizontalAlignment](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_horizontal_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Aligns text to the given tab-stops.
*/
//go:nosplit
func (self class) TabAlign(tab_stops gd.PackedFloat32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(tab_stops))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_tab_align, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetFlags(flags classdb.TextServerJustificationFlag) {
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_set_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFlags() classdb.TextServerJustificationFlag {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerJustificationFlag](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextOverrunBehavior(overrun_behavior classdb.TextServerOverrunBehavior) {
	var frame = callframe.New()
	callframe.Arg(frame, overrun_behavior)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_set_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextOverrunBehavior() classdb.TextServerOverrunBehavior {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerOverrunBehavior](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEllipsisChar(char gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(char))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_set_ellipsis_char, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEllipsisChar() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_ellipsis_char, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns array of inline objects.
*/
//go:nosplit
func (self class) GetObjects() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_objects, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns bounding rectangle of the inline object.
*/
//go:nosplit
func (self class) GetObjectRect(key gd.Variant) gd.Rect2 {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(key))
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_object_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns size of the bounding box of the text.
*/
//go:nosplit
func (self class) GetSize() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns TextServer buffer RID.
*/
//go:nosplit
func (self class) GetRid() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the text ascent (number of pixels above the baseline for horizontal layout or to the left of baseline for vertical).
*/
//go:nosplit
func (self class) GetLineAscent() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_line_ascent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the text descent (number of pixels below the baseline for horizontal layout or to the right of baseline for vertical).
*/
//go:nosplit
func (self class) GetLineDescent() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_line_descent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns width (for horizontal layout) or height (for vertical) of the text.
*/
//go:nosplit
func (self class) GetLineWidth() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_line_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns pixel offset of the underline below the baseline.
*/
//go:nosplit
func (self class) GetLineUnderlinePosition() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_line_underline_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns thickness of the underline.
*/
//go:nosplit
func (self class) GetLineUnderlineThickness() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_get_line_underline_thickness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Draw text into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
//go:nosplit
func (self class) Draw(canvas gd.RID, pos gd.Vector2, color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_draw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Draw text into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
//go:nosplit
func (self class) DrawOutline(canvas gd.RID, pos gd.Vector2, outline_size gd.Int, color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, outline_size)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_draw_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns caret character offset at the specified pixel offset at the baseline. This function always returns a valid position.
*/
//go:nosplit
func (self class) HitTest(coords gd.Float) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextLine.Bind_hit_test, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsTextLine() Advanced           { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTextLine() Instance        { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() { classdb.Register("TextLine", func(ptr gd.Object) any { return classdb.TextLine(ptr) }) }

type HorizontalAlignment int

const (
	/*Horizontal left alignment, usually for text-derived classes.*/
	HorizontalAlignmentLeft HorizontalAlignment = 0
	/*Horizontal center alignment, usually for text-derived classes.*/
	HorizontalAlignmentCenter HorizontalAlignment = 1
	/*Horizontal right alignment, usually for text-derived classes.*/
	HorizontalAlignmentRight HorizontalAlignment = 2
	/*Expand row to fit width, usually for text-derived classes.*/
	HorizontalAlignmentFill HorizontalAlignment = 3
)

type InlineAlignment int

const (
	/*Aligns the top of the inline object (e.g. image, table) to the position of the text specified by [code]INLINE_ALIGNMENT_TO_*[/code] constant.*/
	InlineAlignmentTopTo InlineAlignment = 0
	/*Aligns the center of the inline object (e.g. image, table) to the position of the text specified by [code]INLINE_ALIGNMENT_TO_*[/code] constant.*/
	InlineAlignmentCenterTo InlineAlignment = 1
	/*Aligns the baseline (user defined) of the inline object (e.g. image, table) to the position of the text specified by [code]INLINE_ALIGNMENT_TO_*[/code] constant.*/
	InlineAlignmentBaselineTo InlineAlignment = 3
	/*Aligns the bottom of the inline object (e.g. image, table) to the position of the text specified by [code]INLINE_ALIGNMENT_TO_*[/code] constant.*/
	InlineAlignmentBottomTo InlineAlignment = 2
	/*Aligns the position of the inline object (e.g. image, table) specified by [code]INLINE_ALIGNMENT_*_TO[/code] constant to the top of the text.*/
	InlineAlignmentToTop InlineAlignment = 0
	/*Aligns the position of the inline object (e.g. image, table) specified by [code]INLINE_ALIGNMENT_*_TO[/code] constant to the center of the text.*/
	InlineAlignmentToCenter InlineAlignment = 4
	/*Aligns the position of the inline object (e.g. image, table) specified by [code]INLINE_ALIGNMENT_*_TO[/code] constant to the baseline of the text.*/
	InlineAlignmentToBaseline InlineAlignment = 8
	/*Aligns inline object (e.g. image, table) to the bottom of the text.*/
	InlineAlignmentToBottom InlineAlignment = 12
	/*Aligns top of the inline object (e.g. image, table) to the top of the text. Equivalent to [code]INLINE_ALIGNMENT_TOP_TO | INLINE_ALIGNMENT_TO_TOP[/code].*/
	InlineAlignmentTop InlineAlignment = 0
	/*Aligns center of the inline object (e.g. image, table) to the center of the text. Equivalent to [code]INLINE_ALIGNMENT_CENTER_TO | INLINE_ALIGNMENT_TO_CENTER[/code].*/
	InlineAlignmentCenter InlineAlignment = 5
	/*Aligns bottom of the inline object (e.g. image, table) to the bottom of the text. Equivalent to [code]INLINE_ALIGNMENT_BOTTOM_TO | INLINE_ALIGNMENT_TO_BOTTOM[/code].*/
	InlineAlignmentBottom InlineAlignment = 14
	/*A bit mask for [code]INLINE_ALIGNMENT_*_TO[/code] alignment constants.*/
	InlineAlignmentImageMask InlineAlignment = 3
	/*A bit mask for [code]INLINE_ALIGNMENT_TO_*[/code] alignment constants.*/
	InlineAlignmentTextMask InlineAlignment = 12
)
