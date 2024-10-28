package TextParagraph

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
Abstraction over [TextServer] for handling a single paragraph of text.

*/
type Go [1]classdb.TextParagraph

/*
Clears text paragraph (removes text and inline objects).
*/
func (self Go) Clear() {
	class(self).Clear()
}

/*
Overrides BiDi for the structured text.
Override ranges should cover full source text without overlaps. BiDi algorithm will be used on each range separately.
*/
func (self Go) SetBidiOverride(override gd.Array) {
	class(self).SetBidiOverride(override)
}

/*
Sets drop cap, overrides previously set drop cap. Drop cap (dropped capital) is a decorative element at the beginning of a paragraph that is larger than the rest of the text.
*/
func (self Go) SetDropcap(text string, font gdclass.Font, font_size int) bool {
	return bool(class(self).SetDropcap(gd.NewString(text), font, gd.Int(font_size), gd.NewRect2(0, 0, 0, 0), gd.NewString("")))
}

/*
Removes dropcap.
*/
func (self Go) ClearDropcap() {
	class(self).ClearDropcap()
}

/*
Adds text span and font to draw it.
*/
func (self Go) AddString(text string, font gdclass.Font, font_size int) bool {
	return bool(class(self).AddString(gd.NewString(text), font, gd.Int(font_size), gd.NewString(""), gd.NewVariant(([1]gd.Variant{}[0]))))
}

/*
Adds inline object to the text buffer, [param key] must be unique. In the text, object is represented as [param length] object replacement characters.
*/
func (self Go) AddObject(key gd.Variant, size gd.Vector2) bool {
	return bool(class(self).AddObject(key, size, 5, gd.Int(1), gd.Float(0.0)))
}

/*
Sets new size and alignment of embedded object.
*/
func (self Go) ResizeObject(key gd.Variant, size gd.Vector2) bool {
	return bool(class(self).ResizeObject(key, size, 5, gd.Float(0.0)))
}

/*
Aligns paragraph to the given tab-stops.
*/
func (self Go) TabAlign(tab_stops []float32) {
	class(self).TabAlign(gd.NewPackedFloat32Slice(tab_stops))
}

/*
Returns the size of the bounding box of the paragraph, without line breaks.
*/
func (self Go) GetNonWrappedSize() gd.Vector2 {
	return gd.Vector2(class(self).GetNonWrappedSize())
}

/*
Returns the size of the bounding box of the paragraph.
*/
func (self Go) GetSize() gd.Vector2 {
	return gd.Vector2(class(self).GetSize())
}

/*
Returns TextServer full string buffer RID.
*/
func (self Go) GetRid() gd.RID {
	return gd.RID(class(self).GetRid())
}

/*
Returns TextServer line buffer RID.
*/
func (self Go) GetLineRid(line int) gd.RID {
	return gd.RID(class(self).GetLineRid(gd.Int(line)))
}

/*
Returns drop cap text buffer RID.
*/
func (self Go) GetDropcapRid() gd.RID {
	return gd.RID(class(self).GetDropcapRid())
}

/*
Returns number of lines in the paragraph.
*/
func (self Go) GetLineCount() int {
	return int(int(class(self).GetLineCount()))
}

/*
Returns array of inline objects in the line.
*/
func (self Go) GetLineObjects(line int) gd.Array {
	return gd.Array(class(self).GetLineObjects(gd.Int(line)))
}

/*
Returns bounding rectangle of the inline object.
*/
func (self Go) GetLineObjectRect(line int, key gd.Variant) gd.Rect2 {
	return gd.Rect2(class(self).GetLineObjectRect(gd.Int(line), key))
}

/*
Returns size of the bounding box of the line of text. Returned size is rounded up.
*/
func (self Go) GetLineSize(line int) gd.Vector2 {
	return gd.Vector2(class(self).GetLineSize(gd.Int(line)))
}

/*
Returns character range of the line.
*/
func (self Go) GetLineRange(line int) gd.Vector2i {
	return gd.Vector2i(class(self).GetLineRange(gd.Int(line)))
}

/*
Returns the text line ascent (number of pixels above the baseline for horizontal layout or to the left of baseline for vertical).
*/
func (self Go) GetLineAscent(line int) float64 {
	return float64(float64(class(self).GetLineAscent(gd.Int(line))))
}

/*
Returns the text line descent (number of pixels below the baseline for horizontal layout or to the right of baseline for vertical).
*/
func (self Go) GetLineDescent(line int) float64 {
	return float64(float64(class(self).GetLineDescent(gd.Int(line))))
}

/*
Returns width (for horizontal layout) or height (for vertical) of the line of text.
*/
func (self Go) GetLineWidth(line int) float64 {
	return float64(float64(class(self).GetLineWidth(gd.Int(line))))
}

/*
Returns pixel offset of the underline below the baseline.
*/
func (self Go) GetLineUnderlinePosition(line int) float64 {
	return float64(float64(class(self).GetLineUnderlinePosition(gd.Int(line))))
}

/*
Returns thickness of the underline.
*/
func (self Go) GetLineUnderlineThickness(line int) float64 {
	return float64(float64(class(self).GetLineUnderlineThickness(gd.Int(line))))
}

/*
Returns drop cap bounding box size.
*/
func (self Go) GetDropcapSize() gd.Vector2 {
	return gd.Vector2(class(self).GetDropcapSize())
}

/*
Returns number of lines used by dropcap.
*/
func (self Go) GetDropcapLines() int {
	return int(int(class(self).GetDropcapLines()))
}

/*
Draw all lines of the text and drop cap into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
func (self Go) Draw(canvas gd.RID, pos gd.Vector2) {
	class(self).Draw(canvas, pos, gd.Color{1, 1, 1, 1}, gd.Color{1, 1, 1, 1})
}

/*
Draw outlines of all lines of the text and drop cap into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
func (self Go) DrawOutline(canvas gd.RID, pos gd.Vector2) {
	class(self).DrawOutline(canvas, pos, gd.Int(1), gd.Color{1, 1, 1, 1}, gd.Color{1, 1, 1, 1})
}

/*
Draw single line of text into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
func (self Go) DrawLine(canvas gd.RID, pos gd.Vector2, line int) {
	class(self).DrawLine(canvas, pos, gd.Int(line), gd.Color{1, 1, 1, 1})
}

/*
Draw outline of the single line of text into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
func (self Go) DrawLineOutline(canvas gd.RID, pos gd.Vector2, line int) {
	class(self).DrawLineOutline(canvas, pos, gd.Int(line), gd.Int(1), gd.Color{1, 1, 1, 1})
}

/*
Draw drop cap into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
func (self Go) DrawDropcap(canvas gd.RID, pos gd.Vector2) {
	class(self).DrawDropcap(canvas, pos, gd.Color{1, 1, 1, 1})
}

/*
Draw drop cap outline into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
func (self Go) DrawDropcapOutline(canvas gd.RID, pos gd.Vector2) {
	class(self).DrawDropcapOutline(canvas, pos, gd.Int(1), gd.Color{1, 1, 1, 1})
}

/*
Returns caret character offset at the specified coordinates. This function always returns a valid position.
*/
func (self Go) HitTest(coords gd.Vector2) int {
	return int(int(class(self).HitTest(coords)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.TextParagraph
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TextParagraph"))
	return Go{classdb.TextParagraph(object)}
}

func (self Go) Direction() classdb.TextServerDirection {
		return classdb.TextServerDirection(class(self).GetDirection())
}

func (self Go) SetDirection(value classdb.TextServerDirection) {
	class(self).SetDirection(value)
}

func (self Go) CustomPunctuation() string {
		return string(class(self).GetCustomPunctuation().String())
}

func (self Go) SetCustomPunctuation(value string) {
	class(self).SetCustomPunctuation(gd.NewString(value))
}

func (self Go) Orientation() classdb.TextServerOrientation {
		return classdb.TextServerOrientation(class(self).GetOrientation())
}

func (self Go) SetOrientation(value classdb.TextServerOrientation) {
	class(self).SetOrientation(value)
}

func (self Go) PreserveInvalid() bool {
		return bool(class(self).GetPreserveInvalid())
}

func (self Go) SetPreserveInvalid(value bool) {
	class(self).SetPreserveInvalid(value)
}

func (self Go) PreserveControl() bool {
		return bool(class(self).GetPreserveControl())
}

func (self Go) SetPreserveControl(value bool) {
	class(self).SetPreserveControl(value)
}

func (self Go) Alignment() gd.HorizontalAlignment {
		return gd.HorizontalAlignment(class(self).GetAlignment())
}

func (self Go) SetAlignment(value gd.HorizontalAlignment) {
	class(self).SetAlignment(value)
}

func (self Go) BreakFlags() classdb.TextServerLineBreakFlag {
		return classdb.TextServerLineBreakFlag(class(self).GetBreakFlags())
}

func (self Go) SetBreakFlags(value classdb.TextServerLineBreakFlag) {
	class(self).SetBreakFlags(value)
}

func (self Go) JustificationFlags() classdb.TextServerJustificationFlag {
		return classdb.TextServerJustificationFlag(class(self).GetJustificationFlags())
}

func (self Go) SetJustificationFlags(value classdb.TextServerJustificationFlag) {
	class(self).SetJustificationFlags(value)
}

func (self Go) TextOverrunBehavior() classdb.TextServerOverrunBehavior {
		return classdb.TextServerOverrunBehavior(class(self).GetTextOverrunBehavior())
}

func (self Go) SetTextOverrunBehavior(value classdb.TextServerOverrunBehavior) {
	class(self).SetTextOverrunBehavior(value)
}

func (self Go) EllipsisChar() string {
		return string(class(self).GetEllipsisChar().String())
}

func (self Go) SetEllipsisChar(value string) {
	class(self).SetEllipsisChar(gd.NewString(value))
}

func (self Go) Width() float64 {
		return float64(float64(class(self).GetWidth()))
}

func (self Go) SetWidth(value float64) {
	class(self).SetWidth(gd.Float(value))
}

func (self Go) MaxLinesVisible() int {
		return int(int(class(self).GetMaxLinesVisible()))
}

func (self Go) SetMaxLinesVisible(value int) {
	class(self).SetMaxLinesVisible(gd.Int(value))
}

/*
Clears text paragraph (removes text and inline objects).
*/
//go:nosplit
func (self class) Clear()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetDirection(direction classdb.TextServerDirection)  {
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_set_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDirection() classdb.TextServerDirection {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerDirection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCustomPunctuation(custom_punctuation gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(custom_punctuation))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_set_custom_punctuation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCustomPunctuation() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_custom_punctuation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOrientation(orientation classdb.TextServerOrientation)  {
	var frame = callframe.New()
	callframe.Arg(frame, orientation)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_set_orientation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOrientation() classdb.TextServerOrientation {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerOrientation](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_orientation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPreserveInvalid(enabled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_set_preserve_invalid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPreserveInvalid() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_preserve_invalid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPreserveControl(enabled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_set_preserve_control, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPreserveControl() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_preserve_control, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Overrides BiDi for the structured text.
Override ranges should cover full source text without overlaps. BiDi algorithm will be used on each range separately.
*/
//go:nosplit
func (self class) SetBidiOverride(override gd.Array)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(override))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_set_bidi_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets drop cap, overrides previously set drop cap. Drop cap (dropped capital) is a decorative element at the beginning of a paragraph that is larger than the rest of the text.
*/
//go:nosplit
func (self class) SetDropcap(text gd.String, font gdclass.Font, font_size gd.Int, dropcap_margins gd.Rect2, language gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(text))
	callframe.Arg(frame, discreet.Get(font[0])[0])
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, dropcap_margins)
	callframe.Arg(frame, discreet.Get(language))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_set_dropcap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes dropcap.
*/
//go:nosplit
func (self class) ClearDropcap()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_clear_dropcap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds text span and font to draw it.
*/
//go:nosplit
func (self class) AddString(text gd.String, font gdclass.Font, font_size gd.Int, language gd.String, meta gd.Variant) bool {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(text))
	callframe.Arg(frame, discreet.Get(font[0])[0])
	callframe.Arg(frame, font_size)
	callframe.Arg(frame, discreet.Get(language))
	callframe.Arg(frame, discreet.Get(meta))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_add_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds inline object to the text buffer, [param key] must be unique. In the text, object is represented as [param length] object replacement characters.
*/
//go:nosplit
func (self class) AddObject(key gd.Variant, size gd.Vector2, inline_align gd.InlineAlignment, length gd.Int, baseline gd.Float) bool {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(key))
	callframe.Arg(frame, size)
	callframe.Arg(frame, inline_align)
	callframe.Arg(frame, length)
	callframe.Arg(frame, baseline)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_add_object, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets new size and alignment of embedded object.
*/
//go:nosplit
func (self class) ResizeObject(key gd.Variant, size gd.Vector2, inline_align gd.InlineAlignment, baseline gd.Float) bool {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(key))
	callframe.Arg(frame, size)
	callframe.Arg(frame, inline_align)
	callframe.Arg(frame, baseline)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_resize_object, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAlignment(alignment gd.HorizontalAlignment)  {
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_set_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAlignment() gd.HorizontalAlignment {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.HorizontalAlignment](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Aligns paragraph to the given tab-stops.
*/
//go:nosplit
func (self class) TabAlign(tab_stops gd.PackedFloat32Array)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(tab_stops))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_tab_align, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetBreakFlags(flags classdb.TextServerLineBreakFlag)  {
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_set_break_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBreakFlags() classdb.TextServerLineBreakFlag {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerLineBreakFlag](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_break_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetJustificationFlags(flags classdb.TextServerJustificationFlag)  {
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_set_justification_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetJustificationFlags() classdb.TextServerJustificationFlag {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerJustificationFlag](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_justification_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTextOverrunBehavior(overrun_behavior classdb.TextServerOverrunBehavior)  {
	var frame = callframe.New()
	callframe.Arg(frame, overrun_behavior)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_set_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTextOverrunBehavior() classdb.TextServerOverrunBehavior {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextServerOverrunBehavior](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEllipsisChar(char gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(char))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_set_ellipsis_char, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEllipsisChar() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_ellipsis_char, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetWidth(width gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_set_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetWidth() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the size of the bounding box of the paragraph, without line breaks.
*/
//go:nosplit
func (self class) GetNonWrappedSize() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_non_wrapped_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the size of the bounding box of the paragraph.
*/
//go:nosplit
func (self class) GetSize() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns TextServer full string buffer RID.
*/
//go:nosplit
func (self class) GetRid() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns TextServer line buffer RID.
*/
//go:nosplit
func (self class) GetLineRid(line gd.Int) gd.RID {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_line_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns drop cap text buffer RID.
*/
//go:nosplit
func (self class) GetDropcapRid() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_dropcap_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns number of lines in the paragraph.
*/
//go:nosplit
func (self class) GetLineCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_line_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxLinesVisible(max_lines_visible gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, max_lines_visible)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_set_max_lines_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxLinesVisible() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_max_lines_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns array of inline objects in the line.
*/
//go:nosplit
func (self class) GetLineObjects(line gd.Int) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_line_objects, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns bounding rectangle of the inline object.
*/
//go:nosplit
func (self class) GetLineObjectRect(line gd.Int, key gd.Variant) gd.Rect2 {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, discreet.Get(key))
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_line_object_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns size of the bounding box of the line of text. Returned size is rounded up.
*/
//go:nosplit
func (self class) GetLineSize(line gd.Int) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_line_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns character range of the line.
*/
//go:nosplit
func (self class) GetLineRange(line gd.Int) gd.Vector2i {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_line_range, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the text line ascent (number of pixels above the baseline for horizontal layout or to the left of baseline for vertical).
*/
//go:nosplit
func (self class) GetLineAscent(line gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_line_ascent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the text line descent (number of pixels below the baseline for horizontal layout or to the right of baseline for vertical).
*/
//go:nosplit
func (self class) GetLineDescent(line gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_line_descent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns width (for horizontal layout) or height (for vertical) of the line of text.
*/
//go:nosplit
func (self class) GetLineWidth(line gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_line_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns pixel offset of the underline below the baseline.
*/
//go:nosplit
func (self class) GetLineUnderlinePosition(line gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_line_underline_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns thickness of the underline.
*/
//go:nosplit
func (self class) GetLineUnderlineThickness(line gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_line_underline_thickness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns drop cap bounding box size.
*/
//go:nosplit
func (self class) GetDropcapSize() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_dropcap_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns number of lines used by dropcap.
*/
//go:nosplit
func (self class) GetDropcapLines() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_get_dropcap_lines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Draw all lines of the text and drop cap into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
//go:nosplit
func (self class) Draw(canvas gd.RID, pos gd.Vector2, color gd.Color, dc_color gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, color)
	callframe.Arg(frame, dc_color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_draw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draw outlines of all lines of the text and drop cap into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
//go:nosplit
func (self class) DrawOutline(canvas gd.RID, pos gd.Vector2, outline_size gd.Int, color gd.Color, dc_color gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, outline_size)
	callframe.Arg(frame, color)
	callframe.Arg(frame, dc_color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_draw_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draw single line of text into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
//go:nosplit
func (self class) DrawLine(canvas gd.RID, pos gd.Vector2, line gd.Int, color gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, line)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_draw_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draw outline of the single line of text into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
//go:nosplit
func (self class) DrawLineOutline(canvas gd.RID, pos gd.Vector2, line gd.Int, outline_size gd.Int, color gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, line)
	callframe.Arg(frame, outline_size)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_draw_line_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draw drop cap into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
//go:nosplit
func (self class) DrawDropcap(canvas gd.RID, pos gd.Vector2, color gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_draw_dropcap, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Draw drop cap outline into a canvas item at a given position, with [param color]. [param pos] specifies the top left corner of the bounding box.
*/
//go:nosplit
func (self class) DrawDropcapOutline(canvas gd.RID, pos gd.Vector2, outline_size gd.Int, color gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, canvas)
	callframe.Arg(frame, pos)
	callframe.Arg(frame, outline_size)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_draw_dropcap_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns caret character offset at the specified coordinates. This function always returns a valid position.
*/
//go:nosplit
func (self class) HitTest(coords gd.Vector2) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextParagraph.Bind_hit_test, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsTextParagraph() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsTextParagraph() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("TextParagraph", func(ptr gd.Object) any { return classdb.TextParagraph(ptr) })}
