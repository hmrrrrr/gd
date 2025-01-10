// Package LineEdit provides methods for working with LineEdit object instances.
package LineEdit

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
[LineEdit] provides an input field for editing a single line of text. It features many built-in shortcuts that are always available ([kbd]Ctrl[/kbd] here maps to [kbd]Cmd[/kbd] on macOS):
- [kbd]Ctrl + C[/kbd]: Copy
- [kbd]Ctrl + X[/kbd]: Cut
- [kbd]Ctrl + V[/kbd] or [kbd]Ctrl + Y[/kbd]: Paste/"yank"
- [kbd]Ctrl + Z[/kbd]: Undo
- [kbd]Ctrl + ~[/kbd]: Swap input direction.
- [kbd]Ctrl + Shift + Z[/kbd]: Redo
- [kbd]Ctrl + U[/kbd]: Delete text from the caret position to the beginning of the line
- [kbd]Ctrl + K[/kbd]: Delete text from the caret position to the end of the line
- [kbd]Ctrl + A[/kbd]: Select all text
- [kbd]Up Arrow[/kbd]/[kbd]Down Arrow[/kbd]: Move the caret to the beginning/end of the line
On macOS, some extra keyboard shortcuts are available:
- [kbd]Cmd + F[/kbd]: Same as [kbd]Right Arrow[/kbd], move the caret one character right
- [kbd]Cmd + B[/kbd]: Same as [kbd]Left Arrow[/kbd], move the caret one character left
- [kbd]Cmd + P[/kbd]: Same as [kbd]Up Arrow[/kbd], move the caret to the previous line
- [kbd]Cmd + N[/kbd]: Same as [kbd]Down Arrow[/kbd], move the caret to the next line
- [kbd]Cmd + D[/kbd]: Same as [kbd]Delete[/kbd], delete the character on the right side of caret
- [kbd]Cmd + H[/kbd]: Same as [kbd]Backspace[/kbd], delete the character on the left side of the caret
- [kbd]Cmd + A[/kbd]: Same as [kbd]Home[/kbd], move the caret to the beginning of the line
- [kbd]Cmd + E[/kbd]: Same as [kbd]End[/kbd], move the caret to the end of the line
- [kbd]Cmd + Left Arrow[/kbd]: Same as [kbd]Home[/kbd], move the caret to the beginning of the line
- [kbd]Cmd + Right Arrow[/kbd]: Same as [kbd]End[/kbd], move the caret to the end of the line
*/
type Instance [1]gdclass.LineEdit

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsLineEdit() Instance
}

/*
Erases the [LineEdit]'s [member text].
*/
func (self Instance) Clear() {
	class(self).Clear()
}

/*
Selects characters inside [LineEdit] between [param from] and [param to]. By default, [param from] is at the beginning and [param to] at the end.
[codeblocks]
[gdscript]
text = "Welcome"
select() # Will select "Welcome".
select(4) # Will select "ome".
select(2, 5) # Will select "lco".
[/gdscript]
[csharp]
Text = "Welcome";
Select(); // Will select "Welcome".
Select(4); // Will select "ome".
Select(2, 5); // Will select "lco".
[/csharp]
[/codeblocks]
*/
func (self Instance) Select() {
	class(self).Select(gd.Int(0), gd.Int(-1))
}

/*
Selects the whole [String].
*/
func (self Instance) SelectAll() {
	class(self).SelectAll()
}

/*
Clears the current selection.
*/
func (self Instance) Deselect() {
	class(self).Deselect()
}

/*
Returns [code]true[/code] if the user has selected text.
*/
func (self Instance) HasSelection() bool {
	return bool(class(self).HasSelection())
}

/*
Returns the text inside the selection.
*/
func (self Instance) GetSelectedText() string {
	return string(class(self).GetSelectedText().String())
}

/*
Returns the selection begin column.
*/
func (self Instance) GetSelectionFromColumn() int {
	return int(int(class(self).GetSelectionFromColumn()))
}

/*
Returns the selection end column.
*/
func (self Instance) GetSelectionToColumn() int {
	return int(int(class(self).GetSelectionToColumn()))
}

/*
Returns the scroll offset due to [member caret_column], as a number of characters.
*/
func (self Instance) GetScrollOffset() Float.X {
	return Float.X(Float.X(class(self).GetScrollOffset()))
}

/*
Inserts [param text] at the caret. If the resulting value is longer than [member max_length], nothing happens.
*/
func (self Instance) InsertTextAtCaret(text string) {
	class(self).InsertTextAtCaret(gd.NewString(text))
}

/*
Deletes one character at the caret's current position (equivalent to pressing [kbd]Delete[/kbd]).
*/
func (self Instance) DeleteCharAtCaret() {
	class(self).DeleteCharAtCaret()
}

/*
Deletes a section of the [member text] going from position [param from_column] to [param to_column]. Both parameters should be within the text's length.
*/
func (self Instance) DeleteText(from_column int, to_column int) {
	class(self).DeleteText(gd.Int(from_column), gd.Int(to_column))
}

/*
Executes a given action as defined in the [enum MenuItems] enum.
*/
func (self Instance) MenuOption(option int) {
	class(self).MenuOption(gd.Int(option))
}

/*
Returns the [PopupMenu] of this [LineEdit]. By default, this menu is displayed when right-clicking on the [LineEdit].
You can add custom menu items or remove standard ones. Make sure your IDs don't conflict with the standard ones (see [enum MenuItems]). For example:
[codeblocks]
[gdscript]
func _ready():

	var menu = get_menu()
	# Remove all items after "Redo".
	menu.item_count = menu.get_item_index(MENU_REDO) + 1
	# Add custom items.
	menu.add_separator()
	menu.add_item("Insert Date", MENU_MAX + 1)
	# Connect callback.
	menu.id_pressed.connect(_on_item_pressed)

func _on_item_pressed(id):

	if id == MENU_MAX + 1:
	    insert_text_at_caret(Time.get_date_string_from_system())

[/gdscript]
[csharp]
public override void _Ready()

	{
	    var menu = GetMenu();
	    // Remove all items after "Redo".
	    menu.ItemCount = menu.GetItemIndex(LineEdit.MenuItems.Redo) + 1;
	    // Add custom items.
	    menu.AddSeparator();
	    menu.AddItem("Insert Date", LineEdit.MenuItems.Max + 1);
	    // Add event handler.
	    menu.IdPressed += OnItemPressed;
	}

public void OnItemPressed(int id)

	{
	    if (id == LineEdit.MenuItems.Max + 1)
	    {
	        InsertTextAtCaret(Time.GetDateStringFromSystem());
	    }
	}

[/csharp]
[/codeblocks]
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member Window.visible] property.
*/
func (self Instance) GetMenu() [1]gdclass.PopupMenu {
	return [1]gdclass.PopupMenu(class(self).GetMenu())
}

/*
Returns whether the menu is visible. Use this instead of [code]get_menu().visible[/code] to improve performance (so the creation of the menu is avoided).
*/
func (self Instance) IsMenuVisible() bool {
	return bool(class(self).IsMenuVisible())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.LineEdit

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("LineEdit"))
	casted := Instance{*(*gdclass.LineEdit)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Text() string {
	return string(class(self).GetText().String())
}

func (self Instance) SetText(value string) {
	class(self).SetText(gd.NewString(value))
}

func (self Instance) PlaceholderText() string {
	return string(class(self).GetPlaceholder().String())
}

func (self Instance) SetPlaceholderText(value string) {
	class(self).SetPlaceholder(gd.NewString(value))
}

func (self Instance) Alignment() HorizontalAlignment {
	return HorizontalAlignment(class(self).GetHorizontalAlignment())
}

func (self Instance) SetAlignment(value HorizontalAlignment) {
	class(self).SetHorizontalAlignment(value)
}

func (self Instance) MaxLength() int {
	return int(int(class(self).GetMaxLength()))
}

func (self Instance) SetMaxLength(value int) {
	class(self).SetMaxLength(gd.Int(value))
}

func (self Instance) Editable() bool {
	return bool(class(self).IsEditable())
}

func (self Instance) SetEditable(value bool) {
	class(self).SetEditable(value)
}

func (self Instance) ExpandToTextLength() bool {
	return bool(class(self).IsExpandToTextLengthEnabled())
}

func (self Instance) SetExpandToTextLength(value bool) {
	class(self).SetExpandToTextLengthEnabled(value)
}

func (self Instance) ContextMenuEnabled() bool {
	return bool(class(self).IsContextMenuEnabled())
}

func (self Instance) SetContextMenuEnabled(value bool) {
	class(self).SetContextMenuEnabled(value)
}

func (self Instance) VirtualKeyboardEnabled() bool {
	return bool(class(self).IsVirtualKeyboardEnabled())
}

func (self Instance) SetVirtualKeyboardEnabled(value bool) {
	class(self).SetVirtualKeyboardEnabled(value)
}

func (self Instance) VirtualKeyboardType() gdclass.LineEditVirtualKeyboardType {
	return gdclass.LineEditVirtualKeyboardType(class(self).GetVirtualKeyboardType())
}

func (self Instance) SetVirtualKeyboardType(value gdclass.LineEditVirtualKeyboardType) {
	class(self).SetVirtualKeyboardType(value)
}

func (self Instance) ClearButtonEnabled() bool {
	return bool(class(self).IsClearButtonEnabled())
}

func (self Instance) SetClearButtonEnabled(value bool) {
	class(self).SetClearButtonEnabled(value)
}

func (self Instance) ShortcutKeysEnabled() bool {
	return bool(class(self).IsShortcutKeysEnabled())
}

func (self Instance) SetShortcutKeysEnabled(value bool) {
	class(self).SetShortcutKeysEnabled(value)
}

func (self Instance) MiddleMousePasteEnabled() bool {
	return bool(class(self).IsMiddleMousePasteEnabled())
}

func (self Instance) SetMiddleMousePasteEnabled(value bool) {
	class(self).SetMiddleMousePasteEnabled(value)
}

func (self Instance) SelectingEnabled() bool {
	return bool(class(self).IsSelectingEnabled())
}

func (self Instance) SetSelectingEnabled(value bool) {
	class(self).SetSelectingEnabled(value)
}

func (self Instance) DeselectOnFocusLossEnabled() bool {
	return bool(class(self).IsDeselectOnFocusLossEnabled())
}

func (self Instance) SetDeselectOnFocusLossEnabled(value bool) {
	class(self).SetDeselectOnFocusLossEnabled(value)
}

func (self Instance) DragAndDropSelectionEnabled() bool {
	return bool(class(self).IsDragAndDropSelectionEnabled())
}

func (self Instance) SetDragAndDropSelectionEnabled(value bool) {
	class(self).SetDragAndDropSelectionEnabled(value)
}

func (self Instance) RightIcon() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetRightIcon())
}

func (self Instance) SetRightIcon(value [1]gdclass.Texture2D) {
	class(self).SetRightIcon(value)
}

func (self Instance) Flat() bool {
	return bool(class(self).IsFlat())
}

func (self Instance) SetFlat(value bool) {
	class(self).SetFlat(value)
}

func (self Instance) DrawControlChars() bool {
	return bool(class(self).GetDrawControlChars())
}

func (self Instance) SetDrawControlChars(value bool) {
	class(self).SetDrawControlChars(value)
}

func (self Instance) SelectAllOnFocus() bool {
	return bool(class(self).IsSelectAllOnFocus())
}

func (self Instance) SetSelectAllOnFocus(value bool) {
	class(self).SetSelectAllOnFocus(value)
}

func (self Instance) CaretBlink() bool {
	return bool(class(self).IsCaretBlinkEnabled())
}

func (self Instance) SetCaretBlink(value bool) {
	class(self).SetCaretBlinkEnabled(value)
}

func (self Instance) CaretBlinkInterval() Float.X {
	return Float.X(Float.X(class(self).GetCaretBlinkInterval()))
}

func (self Instance) SetCaretBlinkInterval(value Float.X) {
	class(self).SetCaretBlinkInterval(gd.Float(value))
}

func (self Instance) CaretColumn() int {
	return int(int(class(self).GetCaretColumn()))
}

func (self Instance) SetCaretColumn(value int) {
	class(self).SetCaretColumn(gd.Int(value))
}

func (self Instance) CaretForceDisplayed() bool {
	return bool(class(self).IsCaretForceDisplayed())
}

func (self Instance) SetCaretForceDisplayed(value bool) {
	class(self).SetCaretForceDisplayed(value)
}

func (self Instance) CaretMidGrapheme() bool {
	return bool(class(self).IsCaretMidGraphemeEnabled())
}

func (self Instance) SetCaretMidGrapheme(value bool) {
	class(self).SetCaretMidGraphemeEnabled(value)
}

func (self Instance) Secret() bool {
	return bool(class(self).IsSecret())
}

func (self Instance) SetSecret(value bool) {
	class(self).SetSecret(value)
}

func (self Instance) SecretCharacter() string {
	return string(class(self).GetSecretCharacter().String())
}

func (self Instance) SetSecretCharacter(value string) {
	class(self).SetSecretCharacter(gd.NewString(value))
}

func (self Instance) TextDirection() gdclass.ControlTextDirection {
	return gdclass.ControlTextDirection(class(self).GetTextDirection())
}

func (self Instance) SetTextDirection(value gdclass.ControlTextDirection) {
	class(self).SetTextDirection(value)
}

func (self Instance) Language() string {
	return string(class(self).GetLanguage().String())
}

func (self Instance) SetLanguage(value string) {
	class(self).SetLanguage(gd.NewString(value))
}

func (self Instance) StructuredTextBidiOverride() gdclass.TextServerStructuredTextParser {
	return gdclass.TextServerStructuredTextParser(class(self).GetStructuredTextBidiOverride())
}

func (self Instance) SetStructuredTextBidiOverride(value gdclass.TextServerStructuredTextParser) {
	class(self).SetStructuredTextBidiOverride(value)
}

func (self Instance) StructuredTextBidiOverrideOptions() Array.Any {
	return Array.Any(class(self).GetStructuredTextBidiOverrideOptions())
}

func (self Instance) SetStructuredTextBidiOverrideOptions(value Array.Any) {
	class(self).SetStructuredTextBidiOverrideOptions(value)
}

//go:nosplit
func (self class) SetHorizontalAlignment(alignment HorizontalAlignment) {
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_horizontal_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetHorizontalAlignment() HorizontalAlignment {
	var frame = callframe.New()
	var r_ret = callframe.Ret[HorizontalAlignment](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_get_horizontal_alignment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Erases the [LineEdit]'s [member text].
*/
//go:nosplit
func (self class) Clear() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Selects characters inside [LineEdit] between [param from] and [param to]. By default, [param from] is at the beginning and [param to] at the end.
[codeblocks]
[gdscript]
text = "Welcome"
select() # Will select "Welcome".
select(4) # Will select "ome".
select(2, 5) # Will select "lco".
[/gdscript]
[csharp]
Text = "Welcome";
Select(); // Will select "Welcome".
Select(4); // Will select "ome".
Select(2, 5); // Will select "lco".
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) Select(from gd.Int, to gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_select_, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Selects the whole [String].
*/
//go:nosplit
func (self class) SelectAll() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_select_all, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Clears the current selection.
*/
//go:nosplit
func (self class) Deselect() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_deselect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if the user has selected text.
*/
//go:nosplit
func (self class) HasSelection() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_has_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the text inside the selection.
*/
//go:nosplit
func (self class) GetSelectedText() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_get_selected_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the selection begin column.
*/
//go:nosplit
func (self class) GetSelectionFromColumn() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_get_selection_from_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the selection end column.
*/
//go:nosplit
func (self class) GetSelectionToColumn() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_get_selection_to_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetText(text gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetText() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_get_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetDrawControlChars() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_get_draw_control_chars, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDrawControlChars(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_draw_control_chars, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetTextDirection(direction gdclass.ControlTextDirection) {
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextDirection() gdclass.ControlTextDirection {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.ControlTextDirection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_get_text_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLanguage(language gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(language))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLanguage() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_get_language, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStructuredTextBidiOverride(parser gdclass.TextServerStructuredTextParser) {
	var frame = callframe.New()
	callframe.Arg(frame, parser)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetStructuredTextBidiOverride() gdclass.TextServerStructuredTextParser {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerStructuredTextParser](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_get_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStructuredTextBidiOverrideOptions(args gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(args))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetStructuredTextBidiOverrideOptions() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_get_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPlaceholder(text gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_placeholder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPlaceholder() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_get_placeholder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCaretColumn(position gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_caret_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCaretColumn() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_get_caret_column, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the scroll offset due to [member caret_column], as a number of characters.
*/
//go:nosplit
func (self class) GetScrollOffset() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_get_scroll_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetExpandToTextLengthEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_expand_to_text_length_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsExpandToTextLengthEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_is_expand_to_text_length_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCaretBlinkEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_caret_blink_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsCaretBlinkEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_is_caret_blink_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCaretMidGraphemeEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_caret_mid_grapheme_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsCaretMidGraphemeEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_is_caret_mid_grapheme_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCaretForceDisplayed(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_caret_force_displayed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsCaretForceDisplayed() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_is_caret_force_displayed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCaretBlinkInterval(interval gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, interval)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_caret_blink_interval, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCaretBlinkInterval() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_get_caret_blink_interval, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxLength(chars gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, chars)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_max_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxLength() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_get_max_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Inserts [param text] at the caret. If the resulting value is longer than [member max_length], nothing happens.
*/
//go:nosplit
func (self class) InsertTextAtCaret(text gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_insert_text_at_caret, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Deletes one character at the caret's current position (equivalent to pressing [kbd]Delete[/kbd]).
*/
//go:nosplit
func (self class) DeleteCharAtCaret() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_delete_char_at_caret, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Deletes a section of the [member text] going from position [param from_column] to [param to_column]. Both parameters should be within the text's length.
*/
//go:nosplit
func (self class) DeleteText(from_column gd.Int, to_column gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, from_column)
	callframe.Arg(frame, to_column)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_delete_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetEditable(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_editable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsEditable() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_is_editable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSecret(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_secret, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsSecret() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_is_secret, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSecretCharacter(character gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(character))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_secret_character, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSecretCharacter() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_get_secret_character, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Executes a given action as defined in the [enum MenuItems] enum.
*/
//go:nosplit
func (self class) MenuOption(option gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, option)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_menu_option, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [PopupMenu] of this [LineEdit]. By default, this menu is displayed when right-clicking on the [LineEdit].
You can add custom menu items or remove standard ones. Make sure your IDs don't conflict with the standard ones (see [enum MenuItems]). For example:
[codeblocks]
[gdscript]
func _ready():
    var menu = get_menu()
    # Remove all items after "Redo".
    menu.item_count = menu.get_item_index(MENU_REDO) + 1
    # Add custom items.
    menu.add_separator()
    menu.add_item("Insert Date", MENU_MAX + 1)
    # Connect callback.
    menu.id_pressed.connect(_on_item_pressed)

func _on_item_pressed(id):
    if id == MENU_MAX + 1:
        insert_text_at_caret(Time.get_date_string_from_system())
[/gdscript]
[csharp]
public override void _Ready()
{
    var menu = GetMenu();
    // Remove all items after "Redo".
    menu.ItemCount = menu.GetItemIndex(LineEdit.MenuItems.Redo) + 1;
    // Add custom items.
    menu.AddSeparator();
    menu.AddItem("Insert Date", LineEdit.MenuItems.Max + 1);
    // Add event handler.
    menu.IdPressed += OnItemPressed;
}

public void OnItemPressed(int id)
{
    if (id == LineEdit.MenuItems.Max + 1)
    {
        InsertTextAtCaret(Time.GetDateStringFromSystem());
    }
}
[/csharp]
[/codeblocks]
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member Window.visible] property.
*/
//go:nosplit
func (self class) GetMenu() [1]gdclass.PopupMenu {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_get_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.PopupMenu{gd.PointerLifetimeBoundTo[gdclass.PopupMenu](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns whether the menu is visible. Use this instead of [code]get_menu().visible[/code] to improve performance (so the creation of the menu is avoided).
*/
//go:nosplit
func (self class) IsMenuVisible() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_is_menu_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetContextMenuEnabled(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_context_menu_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsContextMenuEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_is_context_menu_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVirtualKeyboardEnabled(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_virtual_keyboard_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsVirtualKeyboardEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_is_virtual_keyboard_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVirtualKeyboardType(atype gdclass.LineEditVirtualKeyboardType) {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_virtual_keyboard_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVirtualKeyboardType() gdclass.LineEditVirtualKeyboardType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.LineEditVirtualKeyboardType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_get_virtual_keyboard_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetClearButtonEnabled(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_clear_button_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsClearButtonEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_is_clear_button_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShortcutKeysEnabled(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_shortcut_keys_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsShortcutKeysEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_is_shortcut_keys_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMiddleMousePasteEnabled(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_middle_mouse_paste_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsMiddleMousePasteEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_is_middle_mouse_paste_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSelectingEnabled(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_selecting_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsSelectingEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_is_selecting_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDeselectOnFocusLossEnabled(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_deselect_on_focus_loss_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsDeselectOnFocusLossEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_is_deselect_on_focus_loss_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDragAndDropSelectionEnabled(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_drag_and_drop_selection_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsDragAndDropSelectionEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_is_drag_and_drop_selection_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRightIcon(icon [1]gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_right_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRightIcon() [1]gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_get_right_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFlat(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_flat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsFlat() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_is_flat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSelectAllOnFocus(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_set_select_all_on_focus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsSelectAllOnFocus() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.LineEdit.Bind_is_select_all_on_focus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnTextChanged(cb func(new_text string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("text_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnTextChangeRejected(cb func(rejected_substring string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("text_change_rejected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnTextSubmitted(cb func(new_text string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("text_submitted"), gd.NewCallable(cb), 0)
}

func (self class) AsLineEdit() Advanced        { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsLineEdit() Instance     { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.Advanced { return *((*Control.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsControl() Control.Instance {
	return *((*Control.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Control.Advanced(self.AsControl()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Control.Instance(self.AsControl()), name)
	}
}
func init() {
	gdclass.Register("LineEdit", func(ptr gd.Object) any { return [1]gdclass.LineEdit{*(*gdclass.LineEdit)(unsafe.Pointer(&ptr))} })
}

type MenuItems = gdclass.LineEditMenuItems

const (
	/*Cuts (copies and clears) the selected text.*/
	MenuCut MenuItems = 0
	/*Copies the selected text.*/
	MenuCopy MenuItems = 1
	/*Pastes the clipboard text over the selected text (or at the caret's position).
	  Non-printable escape characters are automatically stripped from the OS clipboard via [method String.strip_escapes].*/
	MenuPaste MenuItems = 2
	/*Erases the whole [LineEdit] text.*/
	MenuClear MenuItems = 3
	/*Selects the whole [LineEdit] text.*/
	MenuSelectAll MenuItems = 4
	/*Undoes the previous action.*/
	MenuUndo MenuItems = 5
	/*Reverse the last undo action.*/
	MenuRedo MenuItems = 6
	/*ID of "Text Writing Direction" submenu.*/
	MenuSubmenuTextDir MenuItems = 7
	/*Sets text direction to inherited.*/
	MenuDirInherited MenuItems = 8
	/*Sets text direction to automatic.*/
	MenuDirAuto MenuItems = 9
	/*Sets text direction to left-to-right.*/
	MenuDirLtr MenuItems = 10
	/*Sets text direction to right-to-left.*/
	MenuDirRtl MenuItems = 11
	/*Toggles control character display.*/
	MenuDisplayUcc MenuItems = 12
	/*ID of "Insert Control Character" submenu.*/
	MenuSubmenuInsertUcc MenuItems = 13
	/*Inserts left-to-right mark (LRM) character.*/
	MenuInsertLrm MenuItems = 14
	/*Inserts right-to-left mark (RLM) character.*/
	MenuInsertRlm MenuItems = 15
	/*Inserts start of left-to-right embedding (LRE) character.*/
	MenuInsertLre MenuItems = 16
	/*Inserts start of right-to-left embedding (RLE) character.*/
	MenuInsertRle MenuItems = 17
	/*Inserts start of left-to-right override (LRO) character.*/
	MenuInsertLro MenuItems = 18
	/*Inserts start of right-to-left override (RLO) character.*/
	MenuInsertRlo MenuItems = 19
	/*Inserts pop direction formatting (PDF) character.*/
	MenuInsertPdf MenuItems = 20
	/*Inserts Arabic letter mark (ALM) character.*/
	MenuInsertAlm MenuItems = 21
	/*Inserts left-to-right isolate (LRI) character.*/
	MenuInsertLri MenuItems = 22
	/*Inserts right-to-left isolate (RLI) character.*/
	MenuInsertRli MenuItems = 23
	/*Inserts first strong isolate (FSI) character.*/
	MenuInsertFsi MenuItems = 24
	/*Inserts pop direction isolate (PDI) character.*/
	MenuInsertPdi MenuItems = 25
	/*Inserts zero width joiner (ZWJ) character.*/
	MenuInsertZwj MenuItems = 26
	/*Inserts zero width non-joiner (ZWNJ) character.*/
	MenuInsertZwnj MenuItems = 27
	/*Inserts word joiner (WJ) character.*/
	MenuInsertWj MenuItems = 28
	/*Inserts soft hyphen (SHY) character.*/
	MenuInsertShy MenuItems = 29
	/*Represents the size of the [enum MenuItems] enum.*/
	MenuMax MenuItems = 30
)

type VirtualKeyboardType = gdclass.LineEditVirtualKeyboardType

const (
	/*Default text virtual keyboard.*/
	KeyboardTypeDefault VirtualKeyboardType = 0
	/*Multiline virtual keyboard.*/
	KeyboardTypeMultiline VirtualKeyboardType = 1
	/*Virtual number keypad, useful for PIN entry.*/
	KeyboardTypeNumber VirtualKeyboardType = 2
	/*Virtual number keypad, useful for entering fractional numbers.*/
	KeyboardTypeNumberDecimal VirtualKeyboardType = 3
	/*Virtual phone number keypad.*/
	KeyboardTypePhone VirtualKeyboardType = 4
	/*Virtual keyboard with additional keys to assist with typing email addresses.*/
	KeyboardTypeEmailAddress VirtualKeyboardType = 5
	/*Virtual keyboard for entering a password. On most platforms, this should disable autocomplete and autocapitalization.
	  [b]Note:[/b] This is not supported on Web. Instead, this behaves identically to [constant KEYBOARD_TYPE_DEFAULT].*/
	KeyboardTypePassword VirtualKeyboardType = 6
	/*Virtual keyboard with additional keys to assist with typing URLs.*/
	KeyboardTypeUrl VirtualKeyboardType = 7
)

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
