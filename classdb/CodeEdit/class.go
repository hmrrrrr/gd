// Package CodeEdit provides methods for working with CodeEdit object instances.
package CodeEdit

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/classdb/TextEdit"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Vector2"

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

/*
CodeEdit is a specialized [TextEdit] designed for editing plain text code files. It has many features commonly found in code editors such as line numbers, line folding, code completion, indent management, and string/comment management.
[b]Note:[/b] Regardless of locale, [CodeEdit] will by default always use left-to-right text direction to correctly display source code.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=CodeEdit)
*/
type Instance [1]gdclass.CodeEdit

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCodeEdit() Instance
}
type Interface interface {
	//Override this method to define how the selected entry should be inserted. If [param replace] is [code]true[/code], any existing text should be replaced.
	ConfirmCodeCompletion(replace bool)
	//Override this method to define what happens when the user requests code completion. If [param force] is [code]true[/code], any checks should be bypassed.
	RequestCodeCompletion(force bool)
	//Override this method to define what items in [param candidates] should be displayed.
	//Both [param candidates] and the return is a [Array] of [Dictionary], see [method get_code_completion_option] for [Dictionary] content.
	FilterCodeCompletionCandidates(candidates []map[any]any) []map[any]any
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) ConfirmCodeCompletion(replace bool) { return }
func (self implementation) RequestCodeCompletion(force bool)   { return }
func (self implementation) FilterCodeCompletionCandidates(candidates []map[any]any) (_ []map[any]any) {
	return
}

/*
Override this method to define how the selected entry should be inserted. If [param replace] is [code]true[/code], any existing text should be replaced.
*/
func (Instance) _confirm_code_completion(impl func(ptr unsafe.Pointer, replace bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var replace = gd.UnsafeGet[bool](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, replace)
	}
}

/*
Override this method to define what happens when the user requests code completion. If [param force] is [code]true[/code], any checks should be bypassed.
*/
func (Instance) _request_code_completion(impl func(ptr unsafe.Pointer, force bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var force = gd.UnsafeGet[bool](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, force)
	}
}

/*
Override this method to define what items in [param candidates] should be displayed.
Both [param candidates] and the return is a [Array] of [Dictionary], see [method get_code_completion_option] for [Dictionary] content.
*/
func (Instance) _filter_code_completion_candidates(impl func(ptr unsafe.Pointer, candidates []map[any]any) []map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var candidates = Array.Through(gd.ArrayProxy[Dictionary.Any]{}, pointers.Pack(pointers.New[gd.Array](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalArray(candidates))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, gd.ArrayAs[[]map[any]any](gd.InternalArray(candidates)))
		ptr, ok := pointers.End(gd.InternalArray(gd.ArrayFromSlice[Array.Contains[Dictionary.Any]](ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Perform an indent as if the user activated the "ui_text_indent" action.
*/
func (self Instance) DoIndent() { //gd:CodeEdit.do_indent
	class(self).DoIndent()
}

/*
Indents selected lines, or in the case of no selection the caret line by one.
*/
func (self Instance) IndentLines() { //gd:CodeEdit.indent_lines
	class(self).IndentLines()
}

/*
Unindents selected lines, or in the case of no selection the caret line by one. Same as performing "ui_text_unindent" action.
*/
func (self Instance) UnindentLines() { //gd:CodeEdit.unindent_lines
	class(self).UnindentLines()
}

/*
Converts the indents of lines between [param from_line] and [param to_line] to tabs or spaces as set by [member indent_use_spaces].
Values of [code]-1[/code] convert the entire text.
*/
func (self Instance) ConvertIndent() { //gd:CodeEdit.convert_indent
	class(self).ConvertIndent(gd.Int(-1), gd.Int(-1))
}

/*
Adds a brace pair.
Both the start and end keys must be symbols. Only the start key has to be unique.
*/
func (self Instance) AddAutoBraceCompletionPair(start_key string, end_key string) { //gd:CodeEdit.add_auto_brace_completion_pair
	class(self).AddAutoBraceCompletionPair(String.New(start_key), String.New(end_key))
}

/*
Returns [code]true[/code] if open key [param open_key] exists.
*/
func (self Instance) HasAutoBraceCompletionOpenKey(open_key string) bool { //gd:CodeEdit.has_auto_brace_completion_open_key
	return bool(class(self).HasAutoBraceCompletionOpenKey(String.New(open_key)))
}

/*
Returns [code]true[/code] if close key [param close_key] exists.
*/
func (self Instance) HasAutoBraceCompletionCloseKey(close_key string) bool { //gd:CodeEdit.has_auto_brace_completion_close_key
	return bool(class(self).HasAutoBraceCompletionCloseKey(String.New(close_key)))
}

/*
Gets the matching auto brace close key for [param open_key].
*/
func (self Instance) GetAutoBraceCompletionCloseKey(open_key string) string { //gd:CodeEdit.get_auto_brace_completion_close_key
	return string(class(self).GetAutoBraceCompletionCloseKey(String.New(open_key)).String())
}

/*
Sets the line as breakpointed.
*/
func (self Instance) SetLineAsBreakpoint(line int, breakpointed bool) { //gd:CodeEdit.set_line_as_breakpoint
	class(self).SetLineAsBreakpoint(gd.Int(line), breakpointed)
}

/*
Returns whether the line at the specified index is breakpointed or not.
*/
func (self Instance) IsLineBreakpointed(line int) bool { //gd:CodeEdit.is_line_breakpointed
	return bool(class(self).IsLineBreakpointed(gd.Int(line)))
}

/*
Clears all breakpointed lines.
*/
func (self Instance) ClearBreakpointedLines() { //gd:CodeEdit.clear_breakpointed_lines
	class(self).ClearBreakpointedLines()
}

/*
Gets all breakpointed lines.
*/
func (self Instance) GetBreakpointedLines() []int32 { //gd:CodeEdit.get_breakpointed_lines
	return []int32(class(self).GetBreakpointedLines().AsSlice())
}

/*
Sets the line as bookmarked.
*/
func (self Instance) SetLineAsBookmarked(line int, bookmarked bool) { //gd:CodeEdit.set_line_as_bookmarked
	class(self).SetLineAsBookmarked(gd.Int(line), bookmarked)
}

/*
Returns whether the line at the specified index is bookmarked or not.
*/
func (self Instance) IsLineBookmarked(line int) bool { //gd:CodeEdit.is_line_bookmarked
	return bool(class(self).IsLineBookmarked(gd.Int(line)))
}

/*
Clears all bookmarked lines.
*/
func (self Instance) ClearBookmarkedLines() { //gd:CodeEdit.clear_bookmarked_lines
	class(self).ClearBookmarkedLines()
}

/*
Gets all bookmarked lines.
*/
func (self Instance) GetBookmarkedLines() []int32 { //gd:CodeEdit.get_bookmarked_lines
	return []int32(class(self).GetBookmarkedLines().AsSlice())
}

/*
Sets the line as executing.
*/
func (self Instance) SetLineAsExecuting(line int, executing bool) { //gd:CodeEdit.set_line_as_executing
	class(self).SetLineAsExecuting(gd.Int(line), executing)
}

/*
Returns whether the line at the specified index is marked as executing or not.
*/
func (self Instance) IsLineExecuting(line int) bool { //gd:CodeEdit.is_line_executing
	return bool(class(self).IsLineExecuting(gd.Int(line)))
}

/*
Clears all executed lines.
*/
func (self Instance) ClearExecutingLines() { //gd:CodeEdit.clear_executing_lines
	class(self).ClearExecutingLines()
}

/*
Gets all executing lines.
*/
func (self Instance) GetExecutingLines() []int32 { //gd:CodeEdit.get_executing_lines
	return []int32(class(self).GetExecutingLines().AsSlice())
}

/*
Returns if the given line is foldable, that is, it has indented lines right below it or a comment / string block.
*/
func (self Instance) CanFoldLine(line int) bool { //gd:CodeEdit.can_fold_line
	return bool(class(self).CanFoldLine(gd.Int(line)))
}

/*
Folds the given line, if possible (see [method can_fold_line]).
*/
func (self Instance) FoldLine(line int) { //gd:CodeEdit.fold_line
	class(self).FoldLine(gd.Int(line))
}

/*
Unfolds all lines that were previously folded.
*/
func (self Instance) UnfoldLine(line int) { //gd:CodeEdit.unfold_line
	class(self).UnfoldLine(gd.Int(line))
}

/*
Folds all lines that are possible to be folded (see [method can_fold_line]).
*/
func (self Instance) FoldAllLines() { //gd:CodeEdit.fold_all_lines
	class(self).FoldAllLines()
}

/*
Unfolds all lines, folded or not.
*/
func (self Instance) UnfoldAllLines() { //gd:CodeEdit.unfold_all_lines
	class(self).UnfoldAllLines()
}

/*
Toggle the folding of the code block at the given line.
*/
func (self Instance) ToggleFoldableLine(line int) { //gd:CodeEdit.toggle_foldable_line
	class(self).ToggleFoldableLine(gd.Int(line))
}

/*
Toggle the folding of the code block on all lines with a caret on them.
*/
func (self Instance) ToggleFoldableLinesAtCarets() { //gd:CodeEdit.toggle_foldable_lines_at_carets
	class(self).ToggleFoldableLinesAtCarets()
}

/*
Returns whether the line at the specified index is folded or not.
*/
func (self Instance) IsLineFolded(line int) bool { //gd:CodeEdit.is_line_folded
	return bool(class(self).IsLineFolded(gd.Int(line)))
}

/*
Returns all lines that are current folded.
*/
func (self Instance) GetFoldedLines() []int { //gd:CodeEdit.get_folded_lines
	return []int(gd.ArrayAs[[]int](gd.InternalArray(class(self).GetFoldedLines())))
}

/*
Creates a new code region with the selection. At least one single line comment delimiter have to be defined (see [method add_comment_delimiter]).
A code region is a part of code that is highlighted when folded and can help organize your script.
Code region start and end tags can be customized (see [method set_code_region_tags]).
Code regions are delimited using start and end tags (respectively [code]region[/code] and [code]endregion[/code] by default) preceded by one line comment delimiter. (eg. [code]#region[/code] and [code]#endregion[/code])
*/
func (self Instance) CreateCodeRegion() { //gd:CodeEdit.create_code_region
	class(self).CreateCodeRegion()
}

/*
Returns the code region start tag (without comment delimiter).
*/
func (self Instance) GetCodeRegionStartTag() string { //gd:CodeEdit.get_code_region_start_tag
	return string(class(self).GetCodeRegionStartTag().String())
}

/*
Returns the code region end tag (without comment delimiter).
*/
func (self Instance) GetCodeRegionEndTag() string { //gd:CodeEdit.get_code_region_end_tag
	return string(class(self).GetCodeRegionEndTag().String())
}

/*
Sets the code region start and end tags (without comment delimiter).
*/
func (self Instance) SetCodeRegionTags() { //gd:CodeEdit.set_code_region_tags
	class(self).SetCodeRegionTags(String.New("region"), String.New("endregion"))
}

/*
Returns whether the line at the specified index is a code region start.
*/
func (self Instance) IsLineCodeRegionStart(line int) bool { //gd:CodeEdit.is_line_code_region_start
	return bool(class(self).IsLineCodeRegionStart(gd.Int(line)))
}

/*
Returns whether the line at the specified index is a code region end.
*/
func (self Instance) IsLineCodeRegionEnd(line int) bool { //gd:CodeEdit.is_line_code_region_end
	return bool(class(self).IsLineCodeRegionEnd(gd.Int(line)))
}

/*
Defines a string delimiter from [param start_key] to [param end_key]. Both keys should be symbols, and [param start_key] must not be shared with other delimiters.
If [param line_only] is [code]true[/code] or [param end_key] is an empty [String], the region does not carry over to the next line.
*/
func (self Instance) AddStringDelimiter(start_key string, end_key string) { //gd:CodeEdit.add_string_delimiter
	class(self).AddStringDelimiter(String.New(start_key), String.New(end_key), false)
}

/*
Removes the string delimiter with [param start_key].
*/
func (self Instance) RemoveStringDelimiter(start_key string) { //gd:CodeEdit.remove_string_delimiter
	class(self).RemoveStringDelimiter(String.New(start_key))
}

/*
Returns [code]true[/code] if string [param start_key] exists.
*/
func (self Instance) HasStringDelimiter(start_key string) bool { //gd:CodeEdit.has_string_delimiter
	return bool(class(self).HasStringDelimiter(String.New(start_key)))
}

/*
Removes all string delimiters.
*/
func (self Instance) ClearStringDelimiters() { //gd:CodeEdit.clear_string_delimiters
	class(self).ClearStringDelimiters()
}

/*
Returns the delimiter index if [param line] [param column] is in a string. If [param column] is not provided, will return the delimiter index if the entire [param line] is a string. Otherwise [code]-1[/code].
*/
func (self Instance) IsInString(line int) int { //gd:CodeEdit.is_in_string
	return int(int(class(self).IsInString(gd.Int(line), gd.Int(-1))))
}

/*
Adds a comment delimiter from [param start_key] to [param end_key]. Both keys should be symbols, and [param start_key] must not be shared with other delimiters.
If [param line_only] is [code]true[/code] or [param end_key] is an empty [String], the region does not carry over to the next line.
*/
func (self Instance) AddCommentDelimiter(start_key string, end_key string) { //gd:CodeEdit.add_comment_delimiter
	class(self).AddCommentDelimiter(String.New(start_key), String.New(end_key), false)
}

/*
Removes the comment delimiter with [param start_key].
*/
func (self Instance) RemoveCommentDelimiter(start_key string) { //gd:CodeEdit.remove_comment_delimiter
	class(self).RemoveCommentDelimiter(String.New(start_key))
}

/*
Returns [code]true[/code] if comment [param start_key] exists.
*/
func (self Instance) HasCommentDelimiter(start_key string) bool { //gd:CodeEdit.has_comment_delimiter
	return bool(class(self).HasCommentDelimiter(String.New(start_key)))
}

/*
Removes all comment delimiters.
*/
func (self Instance) ClearCommentDelimiters() { //gd:CodeEdit.clear_comment_delimiters
	class(self).ClearCommentDelimiters()
}

/*
Returns delimiter index if [param line] [param column] is in a comment. If [param column] is not provided, will return delimiter index if the entire [param line] is a comment. Otherwise [code]-1[/code].
*/
func (self Instance) IsInComment(line int) int { //gd:CodeEdit.is_in_comment
	return int(int(class(self).IsInComment(gd.Int(line), gd.Int(-1))))
}

/*
Gets the start key for a string or comment region index.
*/
func (self Instance) GetDelimiterStartKey(delimiter_index int) string { //gd:CodeEdit.get_delimiter_start_key
	return string(class(self).GetDelimiterStartKey(gd.Int(delimiter_index)).String())
}

/*
Gets the end key for a string or comment region index.
*/
func (self Instance) GetDelimiterEndKey(delimiter_index int) string { //gd:CodeEdit.get_delimiter_end_key
	return string(class(self).GetDelimiterEndKey(gd.Int(delimiter_index)).String())
}

/*
If [param line] [param column] is in a string or comment, returns the start position of the region. If not or no start could be found, both [Vector2] values will be [code]-1[/code].
*/
func (self Instance) GetDelimiterStartPosition(line int, column int) Vector2.XY { //gd:CodeEdit.get_delimiter_start_position
	return Vector2.XY(class(self).GetDelimiterStartPosition(gd.Int(line), gd.Int(column)))
}

/*
If [param line] [param column] is in a string or comment, returns the end position of the region. If not or no end could be found, both [Vector2] values will be [code]-1[/code].
*/
func (self Instance) GetDelimiterEndPosition(line int, column int) Vector2.XY { //gd:CodeEdit.get_delimiter_end_position
	return Vector2.XY(class(self).GetDelimiterEndPosition(gd.Int(line), gd.Int(column)))
}

/*
Sets the code hint text. Pass an empty string to clear.
*/
func (self Instance) SetCodeHint(code_hint string) { //gd:CodeEdit.set_code_hint
	class(self).SetCodeHint(String.New(code_hint))
}

/*
Sets if the code hint should draw below the text.
*/
func (self Instance) SetCodeHintDrawBelow(draw_below bool) { //gd:CodeEdit.set_code_hint_draw_below
	class(self).SetCodeHintDrawBelow(draw_below)
}

/*
Returns the full text with char [code]0xFFFF[/code] at the caret location.
*/
func (self Instance) GetTextForCodeCompletion() string { //gd:CodeEdit.get_text_for_code_completion
	return string(class(self).GetTextForCodeCompletion().String())
}

/*
Emits [signal code_completion_requested], if [param force] is [code]true[/code] will bypass all checks. Otherwise will check that the caret is in a word or in front of a prefix. Will ignore the request if all current options are of type file path, node path, or signal.
*/
func (self Instance) RequestCodeCompletion() { //gd:CodeEdit.request_code_completion
	class(self).RequestCodeCompletion(false)
}

/*
Submits an item to the queue of potential candidates for the autocomplete menu. Call [method update_code_completion_options] to update the list.
[param location] indicates location of the option relative to the location of the code completion query. See [enum CodeEdit.CodeCompletionLocation] for how to set this value.
[b]Note:[/b] This list will replace all current candidates.
*/
func (self Instance) AddCodeCompletionOption(atype gdclass.CodeEditCodeCompletionKind, display_text string, insert_text string) { //gd:CodeEdit.add_code_completion_option
	class(self).AddCodeCompletionOption(atype, String.New(display_text), String.New(insert_text), gd.Color(gd.Color{1, 1, 1, 1}), [1][1]gdclass.Resource{}[0], gd.NewVariant(gd.NewVariant(([1]any{}[0]))), gd.Int(1024))
}

/*
Submits all completion options added with [method add_code_completion_option]. Will try to force the autocomplete menu to popup, if [param force] is [code]true[/code].
[b]Note:[/b] This will replace all current candidates.
*/
func (self Instance) UpdateCodeCompletionOptions(force bool) { //gd:CodeEdit.update_code_completion_options
	class(self).UpdateCodeCompletionOptions(force)
}

/*
Gets all completion options, see [method get_code_completion_option] for return content.
*/
func (self Instance) GetCodeCompletionOptions() []CompletionInfo { //gd:CodeEdit.get_code_completion_options
	return []CompletionInfo(gd.ArrayAs[[]CompletionInfo](gd.InternalArray(class(self).GetCodeCompletionOptions())))
}

/*
Gets the completion option at [param index]. The return [Dictionary] has the following key-values:
[code]kind[/code]: [enum CodeCompletionKind]
[code]display_text[/code]: Text that is shown on the autocomplete menu.
[code]insert_text[/code]: Text that is to be inserted when this item is selected.
[code]font_color[/code]: Color of the text on the autocomplete menu.
[code]icon[/code]: Icon to draw on the autocomplete menu.
[code]default_value[/code]: Value of the symbol.
*/
func (self Instance) GetCodeCompletionOption(index int) CompletionInfo { //gd:CodeEdit.get_code_completion_option
	return CompletionInfo(gd.DictionaryAs[CompletionInfo](class(self).GetCodeCompletionOption(gd.Int(index))))
}

/*
Gets the index of the current selected completion option.
*/
func (self Instance) GetCodeCompletionSelectedIndex() int { //gd:CodeEdit.get_code_completion_selected_index
	return int(int(class(self).GetCodeCompletionSelectedIndex()))
}

/*
Sets the current selected completion option.
*/
func (self Instance) SetCodeCompletionSelectedIndex(index int) { //gd:CodeEdit.set_code_completion_selected_index
	class(self).SetCodeCompletionSelectedIndex(gd.Int(index))
}

/*
Inserts the selected entry into the text. If [param replace] is [code]true[/code], any existing text is replaced rather than merged.
*/
func (self Instance) ConfirmCodeCompletion() { //gd:CodeEdit.confirm_code_completion
	class(self).ConfirmCodeCompletion(false)
}

/*
Cancels the autocomplete menu.
*/
func (self Instance) CancelCodeCompletion() { //gd:CodeEdit.cancel_code_completion
	class(self).CancelCodeCompletion()
}

/*
Returns the full text with char [code]0xFFFF[/code] at the cursor location.
*/
func (self Instance) GetTextForSymbolLookup() string { //gd:CodeEdit.get_text_for_symbol_lookup
	return string(class(self).GetTextForSymbolLookup().String())
}

/*
Returns the full text with char [code]0xFFFF[/code] at the specified location.
*/
func (self Instance) GetTextWithCursorChar(line int, column int) string { //gd:CodeEdit.get_text_with_cursor_char
	return string(class(self).GetTextWithCursorChar(gd.Int(line), gd.Int(column)).String())
}

/*
Sets the symbol emitted by [signal symbol_validate] as a valid lookup.
*/
func (self Instance) SetSymbolLookupWordAsValid(valid bool) { //gd:CodeEdit.set_symbol_lookup_word_as_valid
	class(self).SetSymbolLookupWordAsValid(valid)
}

/*
Moves all lines up that are selected or have a caret on them.
*/
func (self Instance) MoveLinesUp() { //gd:CodeEdit.move_lines_up
	class(self).MoveLinesUp()
}

/*
Moves all lines down that are selected or have a caret on them.
*/
func (self Instance) MoveLinesDown() { //gd:CodeEdit.move_lines_down
	class(self).MoveLinesDown()
}

/*
Deletes all lines that are selected or have a caret on them.
*/
func (self Instance) DeleteLines() { //gd:CodeEdit.delete_lines
	class(self).DeleteLines()
}

/*
Duplicates all selected text and duplicates all lines with a caret on them.
*/
func (self Instance) DuplicateSelection() { //gd:CodeEdit.duplicate_selection
	class(self).DuplicateSelection()
}

/*
Duplicates all lines currently selected with any caret. Duplicates the entire line beneath the current one no matter where the caret is within the line.
*/
func (self Instance) DuplicateLines() { //gd:CodeEdit.duplicate_lines
	class(self).DuplicateLines()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.CodeEdit

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CodeEdit"))
	casted := Instance{*(*gdclass.CodeEdit)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) SymbolLookupOnClick() bool {
	return bool(class(self).IsSymbolLookupOnClickEnabled())
}

func (self Instance) SetSymbolLookupOnClick(value bool) {
	class(self).SetSymbolLookupOnClickEnabled(value)
}

func (self Instance) LineFolding() bool {
	return bool(class(self).IsLineFoldingEnabled())
}

func (self Instance) SetLineFolding(value bool) {
	class(self).SetLineFoldingEnabled(value)
}

func (self Instance) LineLengthGuidelines() []int {
	return []int(gd.ArrayAs[[]int](gd.InternalArray(class(self).GetLineLengthGuidelines())))
}

func (self Instance) SetLineLengthGuidelines(value []int) {
	class(self).SetLineLengthGuidelines(gd.ArrayFromSlice[Array.Contains[gd.Int]](value))
}

func (self Instance) GuttersDrawBreakpointsGutter() bool {
	return bool(class(self).IsDrawingBreakpointsGutter())
}

func (self Instance) SetGuttersDrawBreakpointsGutter(value bool) {
	class(self).SetDrawBreakpointsGutter(value)
}

func (self Instance) GuttersDrawBookmarks() bool {
	return bool(class(self).IsDrawingBookmarksGutter())
}

func (self Instance) SetGuttersDrawBookmarks(value bool) {
	class(self).SetDrawBookmarksGutter(value)
}

func (self Instance) GuttersDrawExecutingLines() bool {
	return bool(class(self).IsDrawingExecutingLinesGutter())
}

func (self Instance) SetGuttersDrawExecutingLines(value bool) {
	class(self).SetDrawExecutingLinesGutter(value)
}

func (self Instance) GuttersDrawLineNumbers() bool {
	return bool(class(self).IsDrawLineNumbersEnabled())
}

func (self Instance) SetGuttersDrawLineNumbers(value bool) {
	class(self).SetDrawLineNumbers(value)
}

func (self Instance) GuttersZeroPadLineNumbers() bool {
	return bool(class(self).IsLineNumbersZeroPadded())
}

func (self Instance) SetGuttersZeroPadLineNumbers(value bool) {
	class(self).SetLineNumbersZeroPadded(value)
}

func (self Instance) GuttersDrawFoldGutter() bool {
	return bool(class(self).IsDrawingFoldGutter())
}

func (self Instance) SetGuttersDrawFoldGutter(value bool) {
	class(self).SetDrawFoldGutter(value)
}

func (self Instance) DelimiterStrings() []string {
	return []string(gd.ArrayAs[[]string](gd.InternalArray(class(self).GetStringDelimiters())))
}

func (self Instance) SetDelimiterStrings(value []string) {
	class(self).SetStringDelimiters(gd.ArrayFromSlice[Array.Contains[String.Readable]](value))
}

func (self Instance) DelimiterComments() []string {
	return []string(gd.ArrayAs[[]string](gd.InternalArray(class(self).GetCommentDelimiters())))
}

func (self Instance) SetDelimiterComments(value []string) {
	class(self).SetCommentDelimiters(gd.ArrayFromSlice[Array.Contains[String.Readable]](value))
}

func (self Instance) CodeCompletionEnabled() bool {
	return bool(class(self).IsCodeCompletionEnabled())
}

func (self Instance) SetCodeCompletionEnabled(value bool) {
	class(self).SetCodeCompletionEnabled(value)
}

func (self Instance) CodeCompletionPrefixes() []string {
	return []string(gd.ArrayAs[[]string](gd.InternalArray(class(self).GetCodeCompletionPrefixes())))
}

func (self Instance) SetCodeCompletionPrefixes(value []string) {
	class(self).SetCodeCompletionPrefixes(gd.ArrayFromSlice[Array.Contains[String.Readable]](value))
}

func (self Instance) IndentSize() int {
	return int(int(class(self).GetIndentSize()))
}

func (self Instance) SetIndentSize(value int) {
	class(self).SetIndentSize(gd.Int(value))
}

func (self Instance) IndentUseSpaces() bool {
	return bool(class(self).IsIndentUsingSpaces())
}

func (self Instance) SetIndentUseSpaces(value bool) {
	class(self).SetIndentUsingSpaces(value)
}

func (self Instance) IndentAutomatic() bool {
	return bool(class(self).IsAutoIndentEnabled())
}

func (self Instance) SetIndentAutomatic(value bool) {
	class(self).SetAutoIndentEnabled(value)
}

func (self Instance) IndentAutomaticPrefixes() []string {
	return []string(gd.ArrayAs[[]string](gd.InternalArray(class(self).GetAutoIndentPrefixes())))
}

func (self Instance) SetIndentAutomaticPrefixes(value []string) {
	class(self).SetAutoIndentPrefixes(gd.ArrayFromSlice[Array.Contains[String.Readable]](value))
}

func (self Instance) AutoBraceCompletionEnabled() bool {
	return bool(class(self).IsAutoBraceCompletionEnabled())
}

func (self Instance) SetAutoBraceCompletionEnabled(value bool) {
	class(self).SetAutoBraceCompletionEnabled(value)
}

func (self Instance) AutoBraceCompletionHighlightMatching() bool {
	return bool(class(self).IsHighlightMatchingBracesEnabled())
}

func (self Instance) SetAutoBraceCompletionHighlightMatching(value bool) {
	class(self).SetHighlightMatchingBracesEnabled(value)
}

func (self Instance) AutoBraceCompletionPairs() map[any]any {
	return map[any]any(gd.DictionaryAs[map[any]any](class(self).GetAutoBraceCompletionPairs()))
}

func (self Instance) SetAutoBraceCompletionPairs(value map[any]any) {
	class(self).SetAutoBraceCompletionPairs(gd.DictionaryFromMap(value))
}

/*
Override this method to define how the selected entry should be inserted. If [param replace] is [code]true[/code], any existing text should be replaced.
*/
func (class) _confirm_code_completion(impl func(ptr unsafe.Pointer, replace bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var replace = gd.UnsafeGet[bool](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, replace)
	}
}

/*
Override this method to define what happens when the user requests code completion. If [param force] is [code]true[/code], any checks should be bypassed.
*/
func (class) _request_code_completion(impl func(ptr unsafe.Pointer, force bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var force = gd.UnsafeGet[bool](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, force)
	}
}

/*
Override this method to define what items in [param candidates] should be displayed.
Both [param candidates] and the return is a [Array] of [Dictionary], see [method get_code_completion_option] for [Dictionary] content.
*/
func (class) _filter_code_completion_candidates(impl func(ptr unsafe.Pointer, candidates Array.Contains[Dictionary.Any]) Array.Contains[Dictionary.Any]) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var candidates = Array.Through(gd.ArrayProxy[Dictionary.Any]{}, pointers.Pack(pointers.New[gd.Array](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalArray(candidates))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, candidates)
		ptr, ok := pointers.End(gd.InternalArray(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

//go:nosplit
func (self class) SetIndentSize(size gd.Int) { //gd:CodeEdit.set_indent_size
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_set_indent_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetIndentSize() gd.Int { //gd:CodeEdit.get_indent_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_get_indent_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetIndentUsingSpaces(use_spaces bool) { //gd:CodeEdit.set_indent_using_spaces
	var frame = callframe.New()
	callframe.Arg(frame, use_spaces)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_set_indent_using_spaces, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsIndentUsingSpaces() bool { //gd:CodeEdit.is_indent_using_spaces
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_is_indent_using_spaces, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoIndentEnabled(enable bool) { //gd:CodeEdit.set_auto_indent_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_set_auto_indent_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsAutoIndentEnabled() bool { //gd:CodeEdit.is_auto_indent_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_is_auto_indent_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoIndentPrefixes(prefixes Array.Contains[String.Readable]) { //gd:CodeEdit.set_auto_indent_prefixes
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(prefixes)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_set_auto_indent_prefixes, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutoIndentPrefixes() Array.Contains[String.Readable] { //gd:CodeEdit.get_auto_indent_prefixes
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_get_auto_indent_prefixes, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[String.Readable]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Perform an indent as if the user activated the "ui_text_indent" action.
*/
//go:nosplit
func (self class) DoIndent() { //gd:CodeEdit.do_indent
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_do_indent, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Indents selected lines, or in the case of no selection the caret line by one.
*/
//go:nosplit
func (self class) IndentLines() { //gd:CodeEdit.indent_lines
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_indent_lines, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Unindents selected lines, or in the case of no selection the caret line by one. Same as performing "ui_text_unindent" action.
*/
//go:nosplit
func (self class) UnindentLines() { //gd:CodeEdit.unindent_lines
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_unindent_lines, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Converts the indents of lines between [param from_line] and [param to_line] to tabs or spaces as set by [member indent_use_spaces].
Values of [code]-1[/code] convert the entire text.
*/
//go:nosplit
func (self class) ConvertIndent(from_line gd.Int, to_line gd.Int) { //gd:CodeEdit.convert_indent
	var frame = callframe.New()
	callframe.Arg(frame, from_line)
	callframe.Arg(frame, to_line)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_convert_indent, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetAutoBraceCompletionEnabled(enable bool) { //gd:CodeEdit.set_auto_brace_completion_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_set_auto_brace_completion_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsAutoBraceCompletionEnabled() bool { //gd:CodeEdit.is_auto_brace_completion_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_is_auto_brace_completion_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHighlightMatchingBracesEnabled(enable bool) { //gd:CodeEdit.set_highlight_matching_braces_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_set_highlight_matching_braces_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsHighlightMatchingBracesEnabled() bool { //gd:CodeEdit.is_highlight_matching_braces_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_is_highlight_matching_braces_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a brace pair.
Both the start and end keys must be symbols. Only the start key has to be unique.
*/
//go:nosplit
func (self class) AddAutoBraceCompletionPair(start_key String.Readable, end_key String.Readable) { //gd:CodeEdit.add_auto_brace_completion_pair
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(start_key)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(end_key)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_add_auto_brace_completion_pair, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetAutoBraceCompletionPairs(pairs Dictionary.Any) { //gd:CodeEdit.set_auto_brace_completion_pairs
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalDictionary(pairs)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_set_auto_brace_completion_pairs, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutoBraceCompletionPairs() Dictionary.Any { //gd:CodeEdit.get_auto_brace_completion_pairs
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_get_auto_brace_completion_pairs, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if open key [param open_key] exists.
*/
//go:nosplit
func (self class) HasAutoBraceCompletionOpenKey(open_key String.Readable) bool { //gd:CodeEdit.has_auto_brace_completion_open_key
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(open_key)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_has_auto_brace_completion_open_key, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if close key [param close_key] exists.
*/
//go:nosplit
func (self class) HasAutoBraceCompletionCloseKey(close_key String.Readable) bool { //gd:CodeEdit.has_auto_brace_completion_close_key
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(close_key)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_has_auto_brace_completion_close_key, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Gets the matching auto brace close key for [param open_key].
*/
//go:nosplit
func (self class) GetAutoBraceCompletionCloseKey(open_key String.Readable) String.Readable { //gd:CodeEdit.get_auto_brace_completion_close_key
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(open_key)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_get_auto_brace_completion_close_key, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDrawBreakpointsGutter(enable bool) { //gd:CodeEdit.set_draw_breakpoints_gutter
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_set_draw_breakpoints_gutter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsDrawingBreakpointsGutter() bool { //gd:CodeEdit.is_drawing_breakpoints_gutter
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_is_drawing_breakpoints_gutter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDrawBookmarksGutter(enable bool) { //gd:CodeEdit.set_draw_bookmarks_gutter
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_set_draw_bookmarks_gutter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsDrawingBookmarksGutter() bool { //gd:CodeEdit.is_drawing_bookmarks_gutter
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_is_drawing_bookmarks_gutter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDrawExecutingLinesGutter(enable bool) { //gd:CodeEdit.set_draw_executing_lines_gutter
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_set_draw_executing_lines_gutter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsDrawingExecutingLinesGutter() bool { //gd:CodeEdit.is_drawing_executing_lines_gutter
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_is_drawing_executing_lines_gutter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the line as breakpointed.
*/
//go:nosplit
func (self class) SetLineAsBreakpoint(line gd.Int, breakpointed bool) { //gd:CodeEdit.set_line_as_breakpoint
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, breakpointed)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_set_line_as_breakpoint, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether the line at the specified index is breakpointed or not.
*/
//go:nosplit
func (self class) IsLineBreakpointed(line gd.Int) bool { //gd:CodeEdit.is_line_breakpointed
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_is_line_breakpointed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Clears all breakpointed lines.
*/
//go:nosplit
func (self class) ClearBreakpointedLines() { //gd:CodeEdit.clear_breakpointed_lines
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_clear_breakpointed_lines, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Gets all breakpointed lines.
*/
//go:nosplit
func (self class) GetBreakpointedLines() gd.PackedInt32Array { //gd:CodeEdit.get_breakpointed_lines
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_get_breakpointed_lines, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the line as bookmarked.
*/
//go:nosplit
func (self class) SetLineAsBookmarked(line gd.Int, bookmarked bool) { //gd:CodeEdit.set_line_as_bookmarked
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, bookmarked)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_set_line_as_bookmarked, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether the line at the specified index is bookmarked or not.
*/
//go:nosplit
func (self class) IsLineBookmarked(line gd.Int) bool { //gd:CodeEdit.is_line_bookmarked
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_is_line_bookmarked, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Clears all bookmarked lines.
*/
//go:nosplit
func (self class) ClearBookmarkedLines() { //gd:CodeEdit.clear_bookmarked_lines
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_clear_bookmarked_lines, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Gets all bookmarked lines.
*/
//go:nosplit
func (self class) GetBookmarkedLines() gd.PackedInt32Array { //gd:CodeEdit.get_bookmarked_lines
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_get_bookmarked_lines, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the line as executing.
*/
//go:nosplit
func (self class) SetLineAsExecuting(line gd.Int, executing bool) { //gd:CodeEdit.set_line_as_executing
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, executing)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_set_line_as_executing, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether the line at the specified index is marked as executing or not.
*/
//go:nosplit
func (self class) IsLineExecuting(line gd.Int) bool { //gd:CodeEdit.is_line_executing
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_is_line_executing, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Clears all executed lines.
*/
//go:nosplit
func (self class) ClearExecutingLines() { //gd:CodeEdit.clear_executing_lines
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_clear_executing_lines, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Gets all executing lines.
*/
//go:nosplit
func (self class) GetExecutingLines() gd.PackedInt32Array { //gd:CodeEdit.get_executing_lines
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_get_executing_lines, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDrawLineNumbers(enable bool) { //gd:CodeEdit.set_draw_line_numbers
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_set_draw_line_numbers, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsDrawLineNumbersEnabled() bool { //gd:CodeEdit.is_draw_line_numbers_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_is_draw_line_numbers_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLineNumbersZeroPadded(enable bool) { //gd:CodeEdit.set_line_numbers_zero_padded
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_set_line_numbers_zero_padded, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsLineNumbersZeroPadded() bool { //gd:CodeEdit.is_line_numbers_zero_padded
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_is_line_numbers_zero_padded, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDrawFoldGutter(enable bool) { //gd:CodeEdit.set_draw_fold_gutter
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_set_draw_fold_gutter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsDrawingFoldGutter() bool { //gd:CodeEdit.is_drawing_fold_gutter
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_is_drawing_fold_gutter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLineFoldingEnabled(enabled bool) { //gd:CodeEdit.set_line_folding_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_set_line_folding_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsLineFoldingEnabled() bool { //gd:CodeEdit.is_line_folding_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_is_line_folding_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns if the given line is foldable, that is, it has indented lines right below it or a comment / string block.
*/
//go:nosplit
func (self class) CanFoldLine(line gd.Int) bool { //gd:CodeEdit.can_fold_line
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_can_fold_line, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Folds the given line, if possible (see [method can_fold_line]).
*/
//go:nosplit
func (self class) FoldLine(line gd.Int) { //gd:CodeEdit.fold_line
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_fold_line, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Unfolds all lines that were previously folded.
*/
//go:nosplit
func (self class) UnfoldLine(line gd.Int) { //gd:CodeEdit.unfold_line
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_unfold_line, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Folds all lines that are possible to be folded (see [method can_fold_line]).
*/
//go:nosplit
func (self class) FoldAllLines() { //gd:CodeEdit.fold_all_lines
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_fold_all_lines, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Unfolds all lines, folded or not.
*/
//go:nosplit
func (self class) UnfoldAllLines() { //gd:CodeEdit.unfold_all_lines
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_unfold_all_lines, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Toggle the folding of the code block at the given line.
*/
//go:nosplit
func (self class) ToggleFoldableLine(line gd.Int) { //gd:CodeEdit.toggle_foldable_line
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_toggle_foldable_line, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Toggle the folding of the code block on all lines with a caret on them.
*/
//go:nosplit
func (self class) ToggleFoldableLinesAtCarets() { //gd:CodeEdit.toggle_foldable_lines_at_carets
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_toggle_foldable_lines_at_carets, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether the line at the specified index is folded or not.
*/
//go:nosplit
func (self class) IsLineFolded(line gd.Int) bool { //gd:CodeEdit.is_line_folded
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_is_line_folded, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns all lines that are current folded.
*/
//go:nosplit
func (self class) GetFoldedLines() Array.Contains[gd.Int] { //gd:CodeEdit.get_folded_lines
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_get_folded_lines, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.Int]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Creates a new code region with the selection. At least one single line comment delimiter have to be defined (see [method add_comment_delimiter]).
A code region is a part of code that is highlighted when folded and can help organize your script.
Code region start and end tags can be customized (see [method set_code_region_tags]).
Code regions are delimited using start and end tags (respectively [code]region[/code] and [code]endregion[/code] by default) preceded by one line comment delimiter. (eg. [code]#region[/code] and [code]#endregion[/code])
*/
//go:nosplit
func (self class) CreateCodeRegion() { //gd:CodeEdit.create_code_region
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_create_code_region, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the code region start tag (without comment delimiter).
*/
//go:nosplit
func (self class) GetCodeRegionStartTag() String.Readable { //gd:CodeEdit.get_code_region_start_tag
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_get_code_region_start_tag, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the code region end tag (without comment delimiter).
*/
//go:nosplit
func (self class) GetCodeRegionEndTag() String.Readable { //gd:CodeEdit.get_code_region_end_tag
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_get_code_region_end_tag, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets the code region start and end tags (without comment delimiter).
*/
//go:nosplit
func (self class) SetCodeRegionTags(start String.Readable, end String.Readable) { //gd:CodeEdit.set_code_region_tags
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(start)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(end)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_set_code_region_tags, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether the line at the specified index is a code region start.
*/
//go:nosplit
func (self class) IsLineCodeRegionStart(line gd.Int) bool { //gd:CodeEdit.is_line_code_region_start
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_is_line_code_region_start, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns whether the line at the specified index is a code region end.
*/
//go:nosplit
func (self class) IsLineCodeRegionEnd(line gd.Int) bool { //gd:CodeEdit.is_line_code_region_end
	var frame = callframe.New()
	callframe.Arg(frame, line)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_is_line_code_region_end, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Defines a string delimiter from [param start_key] to [param end_key]. Both keys should be symbols, and [param start_key] must not be shared with other delimiters.
If [param line_only] is [code]true[/code] or [param end_key] is an empty [String], the region does not carry over to the next line.
*/
//go:nosplit
func (self class) AddStringDelimiter(start_key String.Readable, end_key String.Readable, line_only bool) { //gd:CodeEdit.add_string_delimiter
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(start_key)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(end_key)))
	callframe.Arg(frame, line_only)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_add_string_delimiter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes the string delimiter with [param start_key].
*/
//go:nosplit
func (self class) RemoveStringDelimiter(start_key String.Readable) { //gd:CodeEdit.remove_string_delimiter
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(start_key)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_remove_string_delimiter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if string [param start_key] exists.
*/
//go:nosplit
func (self class) HasStringDelimiter(start_key String.Readable) bool { //gd:CodeEdit.has_string_delimiter
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(start_key)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_has_string_delimiter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStringDelimiters(string_delimiters Array.Contains[String.Readable]) { //gd:CodeEdit.set_string_delimiters
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(string_delimiters)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_set_string_delimiters, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes all string delimiters.
*/
//go:nosplit
func (self class) ClearStringDelimiters() { //gd:CodeEdit.clear_string_delimiters
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_clear_string_delimiters, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStringDelimiters() Array.Contains[String.Readable] { //gd:CodeEdit.get_string_delimiters
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_get_string_delimiters, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[String.Readable]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the delimiter index if [param line] [param column] is in a string. If [param column] is not provided, will return the delimiter index if the entire [param line] is a string. Otherwise [code]-1[/code].
*/
//go:nosplit
func (self class) IsInString(line gd.Int, column gd.Int) gd.Int { //gd:CodeEdit.is_in_string
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_is_in_string, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a comment delimiter from [param start_key] to [param end_key]. Both keys should be symbols, and [param start_key] must not be shared with other delimiters.
If [param line_only] is [code]true[/code] or [param end_key] is an empty [String], the region does not carry over to the next line.
*/
//go:nosplit
func (self class) AddCommentDelimiter(start_key String.Readable, end_key String.Readable, line_only bool) { //gd:CodeEdit.add_comment_delimiter
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(start_key)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(end_key)))
	callframe.Arg(frame, line_only)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_add_comment_delimiter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes the comment delimiter with [param start_key].
*/
//go:nosplit
func (self class) RemoveCommentDelimiter(start_key String.Readable) { //gd:CodeEdit.remove_comment_delimiter
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(start_key)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_remove_comment_delimiter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if comment [param start_key] exists.
*/
//go:nosplit
func (self class) HasCommentDelimiter(start_key String.Readable) bool { //gd:CodeEdit.has_comment_delimiter
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(start_key)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_has_comment_delimiter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCommentDelimiters(comment_delimiters Array.Contains[String.Readable]) { //gd:CodeEdit.set_comment_delimiters
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(comment_delimiters)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_set_comment_delimiters, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes all comment delimiters.
*/
//go:nosplit
func (self class) ClearCommentDelimiters() { //gd:CodeEdit.clear_comment_delimiters
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_clear_comment_delimiters, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCommentDelimiters() Array.Contains[String.Readable] { //gd:CodeEdit.get_comment_delimiters
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_get_comment_delimiters, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[String.Readable]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns delimiter index if [param line] [param column] is in a comment. If [param column] is not provided, will return delimiter index if the entire [param line] is a comment. Otherwise [code]-1[/code].
*/
//go:nosplit
func (self class) IsInComment(line gd.Int, column gd.Int) gd.Int { //gd:CodeEdit.is_in_comment
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_is_in_comment, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Gets the start key for a string or comment region index.
*/
//go:nosplit
func (self class) GetDelimiterStartKey(delimiter_index gd.Int) String.Readable { //gd:CodeEdit.get_delimiter_start_key
	var frame = callframe.New()
	callframe.Arg(frame, delimiter_index)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_get_delimiter_start_key, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Gets the end key for a string or comment region index.
*/
//go:nosplit
func (self class) GetDelimiterEndKey(delimiter_index gd.Int) String.Readable { //gd:CodeEdit.get_delimiter_end_key
	var frame = callframe.New()
	callframe.Arg(frame, delimiter_index)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_get_delimiter_end_key, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
If [param line] [param column] is in a string or comment, returns the start position of the region. If not or no start could be found, both [Vector2] values will be [code]-1[/code].
*/
//go:nosplit
func (self class) GetDelimiterStartPosition(line gd.Int, column gd.Int) gd.Vector2 { //gd:CodeEdit.get_delimiter_start_position
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_get_delimiter_start_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param line] [param column] is in a string or comment, returns the end position of the region. If not or no end could be found, both [Vector2] values will be [code]-1[/code].
*/
//go:nosplit
func (self class) GetDelimiterEndPosition(line gd.Int, column gd.Int) gd.Vector2 { //gd:CodeEdit.get_delimiter_end_position
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_get_delimiter_end_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the code hint text. Pass an empty string to clear.
*/
//go:nosplit
func (self class) SetCodeHint(code_hint String.Readable) { //gd:CodeEdit.set_code_hint
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(code_hint)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_set_code_hint, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets if the code hint should draw below the text.
*/
//go:nosplit
func (self class) SetCodeHintDrawBelow(draw_below bool) { //gd:CodeEdit.set_code_hint_draw_below
	var frame = callframe.New()
	callframe.Arg(frame, draw_below)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_set_code_hint_draw_below, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the full text with char [code]0xFFFF[/code] at the caret location.
*/
//go:nosplit
func (self class) GetTextForCodeCompletion() String.Readable { //gd:CodeEdit.get_text_for_code_completion
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_get_text_for_code_completion, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Emits [signal code_completion_requested], if [param force] is [code]true[/code] will bypass all checks. Otherwise will check that the caret is in a word or in front of a prefix. Will ignore the request if all current options are of type file path, node path, or signal.
*/
//go:nosplit
func (self class) RequestCodeCompletion(force bool) { //gd:CodeEdit.request_code_completion
	var frame = callframe.New()
	callframe.Arg(frame, force)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_request_code_completion, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Submits an item to the queue of potential candidates for the autocomplete menu. Call [method update_code_completion_options] to update the list.
[param location] indicates location of the option relative to the location of the code completion query. See [enum CodeEdit.CodeCompletionLocation] for how to set this value.
[b]Note:[/b] This list will replace all current candidates.
*/
//go:nosplit
func (self class) AddCodeCompletionOption(atype gdclass.CodeEditCodeCompletionKind, display_text String.Readable, insert_text String.Readable, text_color gd.Color, icon [1]gdclass.Resource, value gd.Variant, location gd.Int) { //gd:CodeEdit.add_code_completion_option
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, pointers.Get(gd.InternalString(display_text)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(insert_text)))
	callframe.Arg(frame, text_color)
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	callframe.Arg(frame, pointers.Get(value))
	callframe.Arg(frame, location)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_add_code_completion_option, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Submits all completion options added with [method add_code_completion_option]. Will try to force the autocomplete menu to popup, if [param force] is [code]true[/code].
[b]Note:[/b] This will replace all current candidates.
*/
//go:nosplit
func (self class) UpdateCodeCompletionOptions(force bool) { //gd:CodeEdit.update_code_completion_options
	var frame = callframe.New()
	callframe.Arg(frame, force)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_update_code_completion_options, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Gets all completion options, see [method get_code_completion_option] for return content.
*/
//go:nosplit
func (self class) GetCodeCompletionOptions() Array.Contains[Dictionary.Any] { //gd:CodeEdit.get_code_completion_options
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_get_code_completion_options, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Dictionary.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Gets the completion option at [param index]. The return [Dictionary] has the following key-values:
[code]kind[/code]: [enum CodeCompletionKind]
[code]display_text[/code]: Text that is shown on the autocomplete menu.
[code]insert_text[/code]: Text that is to be inserted when this item is selected.
[code]font_color[/code]: Color of the text on the autocomplete menu.
[code]icon[/code]: Icon to draw on the autocomplete menu.
[code]default_value[/code]: Value of the symbol.
*/
//go:nosplit
func (self class) GetCodeCompletionOption(index gd.Int) Dictionary.Any { //gd:CodeEdit.get_code_completion_option
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_get_code_completion_option, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Gets the index of the current selected completion option.
*/
//go:nosplit
func (self class) GetCodeCompletionSelectedIndex() gd.Int { //gd:CodeEdit.get_code_completion_selected_index
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_get_code_completion_selected_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the current selected completion option.
*/
//go:nosplit
func (self class) SetCodeCompletionSelectedIndex(index gd.Int) { //gd:CodeEdit.set_code_completion_selected_index
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_set_code_completion_selected_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Inserts the selected entry into the text. If [param replace] is [code]true[/code], any existing text is replaced rather than merged.
*/
//go:nosplit
func (self class) ConfirmCodeCompletion(replace bool) { //gd:CodeEdit.confirm_code_completion
	var frame = callframe.New()
	callframe.Arg(frame, replace)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_confirm_code_completion, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Cancels the autocomplete menu.
*/
//go:nosplit
func (self class) CancelCodeCompletion() { //gd:CodeEdit.cancel_code_completion
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_cancel_code_completion, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetCodeCompletionEnabled(enable bool) { //gd:CodeEdit.set_code_completion_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_set_code_completion_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsCodeCompletionEnabled() bool { //gd:CodeEdit.is_code_completion_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_is_code_completion_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCodeCompletionPrefixes(prefixes Array.Contains[String.Readable]) { //gd:CodeEdit.set_code_completion_prefixes
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(prefixes)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_set_code_completion_prefixes, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCodeCompletionPrefixes() Array.Contains[String.Readable] { //gd:CodeEdit.get_code_completion_prefixes
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_get_code_completion_prefixes, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[String.Readable]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLineLengthGuidelines(guideline_columns Array.Contains[gd.Int]) { //gd:CodeEdit.set_line_length_guidelines
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(guideline_columns)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_set_line_length_guidelines, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLineLengthGuidelines() Array.Contains[gd.Int] { //gd:CodeEdit.get_line_length_guidelines
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_get_line_length_guidelines, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.Int]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSymbolLookupOnClickEnabled(enable bool) { //gd:CodeEdit.set_symbol_lookup_on_click_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_set_symbol_lookup_on_click_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsSymbolLookupOnClickEnabled() bool { //gd:CodeEdit.is_symbol_lookup_on_click_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_is_symbol_lookup_on_click_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the full text with char [code]0xFFFF[/code] at the cursor location.
*/
//go:nosplit
func (self class) GetTextForSymbolLookup() String.Readable { //gd:CodeEdit.get_text_for_symbol_lookup
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_get_text_for_symbol_lookup, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the full text with char [code]0xFFFF[/code] at the specified location.
*/
//go:nosplit
func (self class) GetTextWithCursorChar(line gd.Int, column gd.Int) String.Readable { //gd:CodeEdit.get_text_with_cursor_char
	var frame = callframe.New()
	callframe.Arg(frame, line)
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_get_text_with_cursor_char, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets the symbol emitted by [signal symbol_validate] as a valid lookup.
*/
//go:nosplit
func (self class) SetSymbolLookupWordAsValid(valid bool) { //gd:CodeEdit.set_symbol_lookup_word_as_valid
	var frame = callframe.New()
	callframe.Arg(frame, valid)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_set_symbol_lookup_word_as_valid, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Moves all lines up that are selected or have a caret on them.
*/
//go:nosplit
func (self class) MoveLinesUp() { //gd:CodeEdit.move_lines_up
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_move_lines_up, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Moves all lines down that are selected or have a caret on them.
*/
//go:nosplit
func (self class) MoveLinesDown() { //gd:CodeEdit.move_lines_down
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_move_lines_down, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Deletes all lines that are selected or have a caret on them.
*/
//go:nosplit
func (self class) DeleteLines() { //gd:CodeEdit.delete_lines
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_delete_lines, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Duplicates all selected text and duplicates all lines with a caret on them.
*/
//go:nosplit
func (self class) DuplicateSelection() { //gd:CodeEdit.duplicate_selection
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_duplicate_selection, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Duplicates all lines currently selected with any caret. Duplicates the entire line beneath the current one no matter where the caret is within the line.
*/
//go:nosplit
func (self class) DuplicateLines() { //gd:CodeEdit.duplicate_lines
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CodeEdit.Bind_duplicate_lines, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self Instance) OnBreakpointToggled(cb func(line int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("breakpoint_toggled"), gd.NewCallable(cb), 0)
}

func (self Instance) OnCodeCompletionRequested(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("code_completion_requested"), gd.NewCallable(cb), 0)
}

func (self Instance) OnSymbolLookup(cb func(symbol string, line int, column int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("symbol_lookup"), gd.NewCallable(cb), 0)
}

func (self Instance) OnSymbolValidate(cb func(symbol string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("symbol_validate"), gd.NewCallable(cb), 0)
}

func (self class) AsCodeEdit() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCodeEdit() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsTextEdit() TextEdit.Advanced {
	return *((*TextEdit.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsTextEdit() TextEdit.Instance {
	return *((*TextEdit.Instance)(unsafe.Pointer(&self)))
}
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
	case "_confirm_code_completion":
		return reflect.ValueOf(self._confirm_code_completion)
	case "_request_code_completion":
		return reflect.ValueOf(self._request_code_completion)
	case "_filter_code_completion_candidates":
		return reflect.ValueOf(self._filter_code_completion_candidates)
	default:
		return gd.VirtualByName(TextEdit.Advanced(self.AsTextEdit()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_confirm_code_completion":
		return reflect.ValueOf(self._confirm_code_completion)
	case "_request_code_completion":
		return reflect.ValueOf(self._request_code_completion)
	case "_filter_code_completion_candidates":
		return reflect.ValueOf(self._filter_code_completion_candidates)
	default:
		return gd.VirtualByName(TextEdit.Instance(self.AsTextEdit()), name)
	}
}
func init() {
	gdclass.Register("CodeEdit", func(ptr gd.Object) any { return [1]gdclass.CodeEdit{*(*gdclass.CodeEdit)(unsafe.Pointer(&ptr))} })
}

type CodeCompletionKind = gdclass.CodeEditCodeCompletionKind //gd:CodeEdit.CodeCompletionKind

const (
	/*Marks the option as a class.*/
	KindClass CodeCompletionKind = 0
	/*Marks the option as a function.*/
	KindFunction CodeCompletionKind = 1
	/*Marks the option as a Godot signal.*/
	KindSignal CodeCompletionKind = 2
	/*Marks the option as a variable.*/
	KindVariable CodeCompletionKind = 3
	/*Marks the option as a member.*/
	KindMember CodeCompletionKind = 4
	/*Marks the option as an enum entry.*/
	KindEnum CodeCompletionKind = 5
	/*Marks the option as a constant.*/
	KindConstant CodeCompletionKind = 6
	/*Marks the option as a Godot node path.*/
	KindNodePath CodeCompletionKind = 7
	/*Marks the option as a file path.*/
	KindFilePath CodeCompletionKind = 8
	/*Marks the option as unclassified or plain text.*/
	KindPlainText CodeCompletionKind = 9
)

type CodeCompletionLocation = gdclass.CodeEditCodeCompletionLocation //gd:CodeEdit.CodeCompletionLocation

const (
	/*The option is local to the location of the code completion query - e.g. a local variable. Subsequent value of location represent options from the outer class, the exact value represent how far they are (in terms of inner classes).*/
	LocationLocal CodeCompletionLocation = 0
	/*The option is from the containing class or a parent class, relative to the location of the code completion query. Perform a bitwise OR with the class depth (e.g. [code]0[/code] for the local class, [code]1[/code] for the parent, [code]2[/code] for the grandparent, etc.) to store the depth of an option in the class or a parent class.*/
	LocationParentMask CodeCompletionLocation = 256
	/*The option is from user code which is not local and not in a derived class (e.g. Autoload Singletons).*/
	LocationOtherUserCode CodeCompletionLocation = 512
	/*The option is from other engine code, not covered by the other enum constants - e.g. built-in classes.*/
	LocationOther CodeCompletionLocation = 1024
)

type CompletionInfo struct {
	Kind        [1]gdclass.CodeEditCodeCompletionKind `gd:"kind"`
	DisplayText string                                `gd:"display_text"`
	InsertText  string                                `gd:"insert_text"`
	FontColor   struct {
		R float32
		G float32
		B float32
		A float32
	} `gd:"font_color"`
	Icon         Resource.Path `gd:"icon"`
	DefaultValue string        `gd:"default_value"`
}
