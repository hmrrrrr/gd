// Package TreeItem provides methods for working with TreeItem object instances.
package TreeItem

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
import "graphics.gd/variant/Rect2"
import "graphics.gd/variant/Color"
import "graphics.gd/variant/Float"

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
A single item of a [Tree] control. It can contain other [TreeItem]s as children, which allows it to create a hierarchy. It can also contain text and buttons. [TreeItem] is not a [Node], it is internal to the [Tree].
To create a [TreeItem], use [method Tree.create_item] or [method TreeItem.create_child]. To remove a [TreeItem], use [method Object.free].
[b]Note:[/b] The ID values used for buttons are 32-bit, unlike [int] which is always 64-bit. They go from [code]-2147483648[/code] to [code]2147483647[/code].
*/
type Instance [1]gdclass.TreeItem

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsTreeItem() Instance
}

/*
Sets the given column's cell mode to [param mode]. This determines how the cell is displayed and edited. See [enum TreeCellMode] constants for details.
*/
func (self Instance) SetCellMode(column int, mode gdclass.TreeItemTreeCellMode) { //gd:TreeItem.set_cell_mode
	class(self).SetCellMode(gd.Int(column), mode)
}

/*
Returns the column's cell mode.
*/
func (self Instance) GetCellMode(column int) gdclass.TreeItemTreeCellMode { //gd:TreeItem.get_cell_mode
	return gdclass.TreeItemTreeCellMode(class(self).GetCellMode(gd.Int(column)))
}

/*
If [param multiline] is [code]true[/code], the given [param column] is multiline editable.
[b]Note:[/b] This option only affects the type of control ([LineEdit] or [TextEdit]) that appears when editing the column. You can set multiline values with [method set_text] even if the column is not multiline editable.
*/
func (self Instance) SetEditMultiline(column int, multiline bool) { //gd:TreeItem.set_edit_multiline
	class(self).SetEditMultiline(gd.Int(column), multiline)
}

/*
Returns [code]true[/code] if the given [param column] is multiline editable.
*/
func (self Instance) IsEditMultiline(column int) bool { //gd:TreeItem.is_edit_multiline
	return bool(class(self).IsEditMultiline(gd.Int(column)))
}

/*
If [param checked] is [code]true[/code], the given [param column] is checked. Clears column's indeterminate status.
*/
func (self Instance) SetChecked(column int, checked bool) { //gd:TreeItem.set_checked
	class(self).SetChecked(gd.Int(column), checked)
}

/*
If [param indeterminate] is [code]true[/code], the given [param column] is marked indeterminate.
[b]Note:[/b] If set [code]true[/code] from [code]false[/code], then column is cleared of checked status.
*/
func (self Instance) SetIndeterminate(column int, indeterminate bool) { //gd:TreeItem.set_indeterminate
	class(self).SetIndeterminate(gd.Int(column), indeterminate)
}

/*
Returns [code]true[/code] if the given [param column] is checked.
*/
func (self Instance) IsChecked(column int) bool { //gd:TreeItem.is_checked
	return bool(class(self).IsChecked(gd.Int(column)))
}

/*
Returns [code]true[/code] if the given [param column] is indeterminate.
*/
func (self Instance) IsIndeterminate(column int) bool { //gd:TreeItem.is_indeterminate
	return bool(class(self).IsIndeterminate(gd.Int(column)))
}

/*
Propagates this item's checked status to its children and parents for the given [param column]. It is possible to process the items affected by this method call by connecting to [signal Tree.check_propagated_to_item]. The order that the items affected will be processed is as follows: the item invoking this method, children of that item, and finally parents of that item. If [param emit_signal] is [code]false[/code], then [signal Tree.check_propagated_to_item] will not be emitted.
*/
func (self Instance) PropagateCheck(column int) { //gd:TreeItem.propagate_check
	class(self).PropagateCheck(gd.Int(column), true)
}

/*
Sets the given column's text value.
*/
func (self Instance) SetText(column int, text string) { //gd:TreeItem.set_text
	class(self).SetText(gd.Int(column), String.New(text))
}

/*
Returns the given column's text.
*/
func (self Instance) GetText(column int) string { //gd:TreeItem.get_text
	return string(class(self).GetText(gd.Int(column)).String())
}

/*
Sets item's text base writing direction.
*/
func (self Instance) SetTextDirection(column int, direction gdclass.ControlTextDirection) { //gd:TreeItem.set_text_direction
	class(self).SetTextDirection(gd.Int(column), direction)
}

/*
Returns item's text base writing direction.
*/
func (self Instance) GetTextDirection(column int) gdclass.ControlTextDirection { //gd:TreeItem.get_text_direction
	return gdclass.ControlTextDirection(class(self).GetTextDirection(gd.Int(column)))
}

/*
Sets the autowrap mode in the given [param column]. If set to something other than [constant TextServer.AUTOWRAP_OFF], the text gets wrapped inside the cell's bounding rectangle.
*/
func (self Instance) SetAutowrapMode(column int, autowrap_mode gdclass.TextServerAutowrapMode) { //gd:TreeItem.set_autowrap_mode
	class(self).SetAutowrapMode(gd.Int(column), autowrap_mode)
}

/*
Returns the text autowrap mode in the given [param column]. By default it is [constant TextServer.AUTOWRAP_OFF].
*/
func (self Instance) GetAutowrapMode(column int) gdclass.TextServerAutowrapMode { //gd:TreeItem.get_autowrap_mode
	return gdclass.TextServerAutowrapMode(class(self).GetAutowrapMode(gd.Int(column)))
}

/*
Sets the clipping behavior when the text exceeds the item's bounding rectangle in the given [param column].
*/
func (self Instance) SetTextOverrunBehavior(column int, overrun_behavior gdclass.TextServerOverrunBehavior) { //gd:TreeItem.set_text_overrun_behavior
	class(self).SetTextOverrunBehavior(gd.Int(column), overrun_behavior)
}

/*
Returns the clipping behavior when the text exceeds the item's bounding rectangle in the given [param column]. By default it is [constant TextServer.OVERRUN_TRIM_ELLIPSIS].
*/
func (self Instance) GetTextOverrunBehavior(column int) gdclass.TextServerOverrunBehavior { //gd:TreeItem.get_text_overrun_behavior
	return gdclass.TextServerOverrunBehavior(class(self).GetTextOverrunBehavior(gd.Int(column)))
}

/*
Set BiDi algorithm override for the structured text. Has effect for cells that display text.
*/
func (self Instance) SetStructuredTextBidiOverride(column int, parser gdclass.TextServerStructuredTextParser) { //gd:TreeItem.set_structured_text_bidi_override
	class(self).SetStructuredTextBidiOverride(gd.Int(column), parser)
}

/*
Returns the BiDi algorithm override set for this cell.
*/
func (self Instance) GetStructuredTextBidiOverride(column int) gdclass.TextServerStructuredTextParser { //gd:TreeItem.get_structured_text_bidi_override
	return gdclass.TextServerStructuredTextParser(class(self).GetStructuredTextBidiOverride(gd.Int(column)))
}

/*
Set additional options for BiDi override. Has effect for cells that display text.
*/
func (self Instance) SetStructuredTextBidiOverrideOptions(column int, args []any) { //gd:TreeItem.set_structured_text_bidi_override_options
	class(self).SetStructuredTextBidiOverrideOptions(gd.Int(column), gd.EngineArrayFromSlice(args))
}

/*
Returns the additional BiDi options set for this cell.
*/
func (self Instance) GetStructuredTextBidiOverrideOptions(column int) []any { //gd:TreeItem.get_structured_text_bidi_override_options
	return []any(gd.ArrayAs[[]any](gd.InternalArray(class(self).GetStructuredTextBidiOverrideOptions(gd.Int(column)))))
}

/*
Sets language code of item's text used for line-breaking and text shaping algorithms, if left empty current locale is used instead.
*/
func (self Instance) SetLanguage(column int, language string) { //gd:TreeItem.set_language
	class(self).SetLanguage(gd.Int(column), String.New(language))
}

/*
Returns item's text language code.
*/
func (self Instance) GetLanguage(column int) string { //gd:TreeItem.get_language
	return string(class(self).GetLanguage(gd.Int(column)).String())
}

/*
Sets a string to be shown after a column's value (for example, a unit abbreviation).
*/
func (self Instance) SetSuffix(column int, text string) { //gd:TreeItem.set_suffix
	class(self).SetSuffix(gd.Int(column), String.New(text))
}

/*
Gets the suffix string shown after the column value.
*/
func (self Instance) GetSuffix(column int) string { //gd:TreeItem.get_suffix
	return string(class(self).GetSuffix(gd.Int(column)).String())
}

/*
Sets the given cell's icon [Texture2D]. The cell has to be in [constant CELL_MODE_ICON] mode.
*/
func (self Instance) SetIcon(column int, texture [1]gdclass.Texture2D) { //gd:TreeItem.set_icon
	class(self).SetIcon(gd.Int(column), texture)
}

/*
Returns the given column's icon [Texture2D]. Error if no icon is set.
*/
func (self Instance) GetIcon(column int) [1]gdclass.Texture2D { //gd:TreeItem.get_icon
	return [1]gdclass.Texture2D(class(self).GetIcon(gd.Int(column)))
}

/*
Sets the given column's icon's texture region.
*/
func (self Instance) SetIconRegion(column int, region Rect2.PositionSize) { //gd:TreeItem.set_icon_region
	class(self).SetIconRegion(gd.Int(column), gd.Rect2(region))
}

/*
Returns the icon [Texture2D] region as [Rect2].
*/
func (self Instance) GetIconRegion(column int) Rect2.PositionSize { //gd:TreeItem.get_icon_region
	return Rect2.PositionSize(class(self).GetIconRegion(gd.Int(column)))
}

/*
Sets the maximum allowed width of the icon in the given [param column]. This limit is applied on top of the default size of the icon and on top of [theme_item Tree.icon_max_width]. The height is adjusted according to the icon's ratio.
*/
func (self Instance) SetIconMaxWidth(column int, width int) { //gd:TreeItem.set_icon_max_width
	class(self).SetIconMaxWidth(gd.Int(column), gd.Int(width))
}

/*
Returns the maximum allowed width of the icon in the given [param column].
*/
func (self Instance) GetIconMaxWidth(column int) int { //gd:TreeItem.get_icon_max_width
	return int(int(class(self).GetIconMaxWidth(gd.Int(column))))
}

/*
Modulates the given column's icon with [param modulate].
*/
func (self Instance) SetIconModulate(column int, modulate Color.RGBA) { //gd:TreeItem.set_icon_modulate
	class(self).SetIconModulate(gd.Int(column), gd.Color(modulate))
}

/*
Returns the [Color] modulating the column's icon.
*/
func (self Instance) GetIconModulate(column int) Color.RGBA { //gd:TreeItem.get_icon_modulate
	return Color.RGBA(class(self).GetIconModulate(gd.Int(column)))
}

/*
Sets the value of a [constant CELL_MODE_RANGE] column.
*/
func (self Instance) SetRange(column int, value Float.X) { //gd:TreeItem.set_range
	class(self).SetRange(gd.Int(column), gd.Float(value))
}

/*
Returns the value of a [constant CELL_MODE_RANGE] column.
*/
func (self Instance) GetRange(column int) Float.X { //gd:TreeItem.get_range
	return Float.X(Float.X(class(self).GetRange(gd.Int(column))))
}

/*
Sets the range of accepted values for a column. The column must be in the [constant CELL_MODE_RANGE] mode.
If [param expr] is [code]true[/code], the edit mode slider will use an exponential scale as with [member Range.exp_edit].
*/
func (self Instance) SetRangeConfig(column int, min Float.X, max Float.X, step Float.X) { //gd:TreeItem.set_range_config
	class(self).SetRangeConfig(gd.Int(column), gd.Float(min), gd.Float(max), gd.Float(step), false)
}

/*
Returns a dictionary containing the range parameters for a given column. The keys are "min", "max", "step", and "expr".
*/
func (self Instance) GetRangeConfig(column int) RangeConfig { //gd:TreeItem.get_range_config
	return RangeConfig(gd.DictionaryAs[RangeConfig](class(self).GetRangeConfig(gd.Int(column))))
}

/*
Sets the metadata value for the given column, which can be retrieved later using [method get_metadata]. This can be used, for example, to store a reference to the original data.
*/
func (self Instance) SetMetadata(column int, meta any) { //gd:TreeItem.set_metadata
	class(self).SetMetadata(gd.Int(column), gd.NewVariant(meta))
}

/*
Returns the metadata value that was set for the given column using [method set_metadata].
*/
func (self Instance) GetMetadata(column int) any { //gd:TreeItem.get_metadata
	return any(class(self).GetMetadata(gd.Int(column)).Interface())
}

/*
Sets the given column's custom draw callback to the [param callback] method on [param object].
The method named [param callback] should accept two arguments: the [TreeItem] that is drawn and its position and size as a [Rect2].
*/
func (self Instance) SetCustomDraw(column int, obj Object.Instance, callback string) { //gd:TreeItem.set_custom_draw
	class(self).SetCustomDraw(gd.Int(column), obj, gd.NewStringName(callback))
}

/*
Sets the given column's custom draw callback. Use an empty [Callable] ([code skip-lint]Callable()[/code]) to clear the custom callback. The cell has to be in [constant CELL_MODE_CUSTOM] to use this feature.
The [param callback] should accept two arguments: the [TreeItem] that is drawn and its position and size as a [Rect2].
*/
func (self Instance) SetCustomDrawCallback(column int, callback func(item [1]gdclass.TreeItem, rect Rect2.PositionSize)) { //gd:TreeItem.set_custom_draw_callback
	class(self).SetCustomDrawCallback(gd.Int(column), Callable.New(callback))
}

/*
Returns the custom callback of column [param column].
*/
func (self Instance) GetCustomDrawCallback(column int) Callable.Function { //gd:TreeItem.get_custom_draw_callback
	return Callable.Function(class(self).GetCustomDrawCallback(gd.Int(column)))
}

/*
Collapses or uncollapses this [TreeItem] and all the descendants of this item.
*/
func (self Instance) SetCollapsedRecursive(enable bool) { //gd:TreeItem.set_collapsed_recursive
	class(self).SetCollapsedRecursive(enable)
}

/*
Returns [code]true[/code] if this [TreeItem], or any of its descendants, is collapsed.
If [param only_visible] is [code]true[/code] it ignores non-visible [TreeItem]s.
*/
func (self Instance) IsAnyCollapsed() bool { //gd:TreeItem.is_any_collapsed
	return bool(class(self).IsAnyCollapsed(false))
}

/*
Returns [code]true[/code] if [member visible] is [code]true[/code] and all its ancestors are also visible.
*/
func (self Instance) IsVisibleInTree() bool { //gd:TreeItem.is_visible_in_tree
	return bool(class(self).IsVisibleInTree())
}

/*
Uncollapses all [TreeItem]s necessary to reveal this [TreeItem], i.e. all ancestor [TreeItem]s.
*/
func (self Instance) UncollapseTree() { //gd:TreeItem.uncollapse_tree
	class(self).UncollapseTree()
}

/*
If [param selectable] is [code]true[/code], the given [param column] is selectable.
*/
func (self Instance) SetSelectable(column int, selectable bool) { //gd:TreeItem.set_selectable
	class(self).SetSelectable(gd.Int(column), selectable)
}

/*
Returns [code]true[/code] if the given [param column] is selectable.
*/
func (self Instance) IsSelectable(column int) bool { //gd:TreeItem.is_selectable
	return bool(class(self).IsSelectable(gd.Int(column)))
}

/*
Returns [code]true[/code] if the given [param column] is selected.
*/
func (self Instance) IsSelected(column int) bool { //gd:TreeItem.is_selected
	return bool(class(self).IsSelected(gd.Int(column)))
}

/*
Selects the given [param column].
*/
func (self Instance) Select(column int) { //gd:TreeItem.select
	class(self).Select(gd.Int(column))
}

/*
Deselects the given column.
*/
func (self Instance) Deselect(column int) { //gd:TreeItem.deselect
	class(self).Deselect(gd.Int(column))
}

/*
If [param enabled] is [code]true[/code], the given [param column] is editable.
*/
func (self Instance) SetEditable(column int, enabled bool) { //gd:TreeItem.set_editable
	class(self).SetEditable(gd.Int(column), enabled)
}

/*
Returns [code]true[/code] if the given [param column] is editable.
*/
func (self Instance) IsEditable(column int) bool { //gd:TreeItem.is_editable
	return bool(class(self).IsEditable(gd.Int(column)))
}

/*
Sets the given column's custom color.
*/
func (self Instance) SetCustomColor(column int, color Color.RGBA) { //gd:TreeItem.set_custom_color
	class(self).SetCustomColor(gd.Int(column), gd.Color(color))
}

/*
Returns the custom color of column [param column].
*/
func (self Instance) GetCustomColor(column int) Color.RGBA { //gd:TreeItem.get_custom_color
	return Color.RGBA(class(self).GetCustomColor(gd.Int(column)))
}

/*
Resets the color for the given column to default.
*/
func (self Instance) ClearCustomColor(column int) { //gd:TreeItem.clear_custom_color
	class(self).ClearCustomColor(gd.Int(column))
}

/*
Sets custom font used to draw text in the given [param column].
*/
func (self Instance) SetCustomFont(column int, font [1]gdclass.Font) { //gd:TreeItem.set_custom_font
	class(self).SetCustomFont(gd.Int(column), font)
}

/*
Returns custom font used to draw text in the column [param column].
*/
func (self Instance) GetCustomFont(column int) [1]gdclass.Font { //gd:TreeItem.get_custom_font
	return [1]gdclass.Font(class(self).GetCustomFont(gd.Int(column)))
}

/*
Sets custom font size used to draw text in the given [param column].
*/
func (self Instance) SetCustomFontSize(column int, font_size int) { //gd:TreeItem.set_custom_font_size
	class(self).SetCustomFontSize(gd.Int(column), gd.Int(font_size))
}

/*
Returns custom font size used to draw text in the column [param column].
*/
func (self Instance) GetCustomFontSize(column int) int { //gd:TreeItem.get_custom_font_size
	return int(int(class(self).GetCustomFontSize(gd.Int(column))))
}

/*
Sets the given column's custom background color and whether to just use it as an outline.
*/
func (self Instance) SetCustomBgColor(column int, color Color.RGBA) { //gd:TreeItem.set_custom_bg_color
	class(self).SetCustomBgColor(gd.Int(column), gd.Color(color), false)
}

/*
Resets the background color for the given column to default.
*/
func (self Instance) ClearCustomBgColor(column int) { //gd:TreeItem.clear_custom_bg_color
	class(self).ClearCustomBgColor(gd.Int(column))
}

/*
Returns the custom background color of column [param column].
*/
func (self Instance) GetCustomBgColor(column int) Color.RGBA { //gd:TreeItem.get_custom_bg_color
	return Color.RGBA(class(self).GetCustomBgColor(gd.Int(column)))
}

/*
Makes a cell with [constant CELL_MODE_CUSTOM] display as a non-flat button with a [StyleBox].
*/
func (self Instance) SetCustomAsButton(column int, enable bool) { //gd:TreeItem.set_custom_as_button
	class(self).SetCustomAsButton(gd.Int(column), enable)
}

/*
Returns [code]true[/code] if the cell was made into a button with [method set_custom_as_button].
*/
func (self Instance) IsCustomSetAsButton(column int) bool { //gd:TreeItem.is_custom_set_as_button
	return bool(class(self).IsCustomSetAsButton(gd.Int(column)))
}

/*
Adds a button with [Texture2D] [param button] at column [param column]. The [param id] is used to identify the button in the according [signal Tree.button_clicked] signal and can be different from the buttons index. If not specified, the next available index is used, which may be retrieved by calling [method get_button_count] immediately before this method. Optionally, the button can be [param disabled] and have a [param tooltip_text].
*/
func (self Instance) AddButton(column int, button [1]gdclass.Texture2D) { //gd:TreeItem.add_button
	class(self).AddButton(gd.Int(column), button, gd.Int(-1), false, String.New(""))
}

/*
Returns the number of buttons in column [param column].
*/
func (self Instance) GetButtonCount(column int) int { //gd:TreeItem.get_button_count
	return int(int(class(self).GetButtonCount(gd.Int(column))))
}

/*
Returns the tooltip text for the button at index [param button_index] in column [param column].
*/
func (self Instance) GetButtonTooltipText(column int, button_index int) string { //gd:TreeItem.get_button_tooltip_text
	return string(class(self).GetButtonTooltipText(gd.Int(column), gd.Int(button_index)).String())
}

/*
Returns the ID for the button at index [param button_index] in column [param column].
*/
func (self Instance) GetButtonId(column int, button_index int) int { //gd:TreeItem.get_button_id
	return int(int(class(self).GetButtonId(gd.Int(column), gd.Int(button_index))))
}

/*
Returns the button index if there is a button with ID [param id] in column [param column], otherwise returns -1.
*/
func (self Instance) GetButtonById(column int, id int) int { //gd:TreeItem.get_button_by_id
	return int(int(class(self).GetButtonById(gd.Int(column), gd.Int(id))))
}

/*
Returns the color of the button with ID [param id] in column [param column]. If the specified button does not exist, returns [constant Color.BLACK].
*/
func (self Instance) GetButtonColor(column int, id int) Color.RGBA { //gd:TreeItem.get_button_color
	return Color.RGBA(class(self).GetButtonColor(gd.Int(column), gd.Int(id)))
}

/*
Returns the [Texture2D] of the button at index [param button_index] in column [param column].
*/
func (self Instance) GetButton(column int, button_index int) [1]gdclass.Texture2D { //gd:TreeItem.get_button
	return [1]gdclass.Texture2D(class(self).GetButton(gd.Int(column), gd.Int(button_index)))
}

/*
Sets the tooltip text for the button at index [param button_index] in the given [param column].
*/
func (self Instance) SetButtonTooltipText(column int, button_index int, tooltip string) { //gd:TreeItem.set_button_tooltip_text
	class(self).SetButtonTooltipText(gd.Int(column), gd.Int(button_index), String.New(tooltip))
}

/*
Sets the given column's button [Texture2D] at index [param button_index] to [param button].
*/
func (self Instance) SetButton(column int, button_index int, button [1]gdclass.Texture2D) { //gd:TreeItem.set_button
	class(self).SetButton(gd.Int(column), gd.Int(button_index), button)
}

/*
Removes the button at index [param button_index] in column [param column].
*/
func (self Instance) EraseButton(column int, button_index int) { //gd:TreeItem.erase_button
	class(self).EraseButton(gd.Int(column), gd.Int(button_index))
}

/*
If [code]true[/code], disables the button at index [param button_index] in the given [param column].
*/
func (self Instance) SetButtonDisabled(column int, button_index int, disabled bool) { //gd:TreeItem.set_button_disabled
	class(self).SetButtonDisabled(gd.Int(column), gd.Int(button_index), disabled)
}

/*
Sets the given column's button color at index [param button_index] to [param color].
*/
func (self Instance) SetButtonColor(column int, button_index int, color Color.RGBA) { //gd:TreeItem.set_button_color
	class(self).SetButtonColor(gd.Int(column), gd.Int(button_index), gd.Color(color))
}

/*
Returns [code]true[/code] if the button at index [param button_index] for the given [param column] is disabled.
*/
func (self Instance) IsButtonDisabled(column int, button_index int) bool { //gd:TreeItem.is_button_disabled
	return bool(class(self).IsButtonDisabled(gd.Int(column), gd.Int(button_index)))
}

/*
Sets the given column's tooltip text.
*/
func (self Instance) SetTooltipText(column int, tooltip string) { //gd:TreeItem.set_tooltip_text
	class(self).SetTooltipText(gd.Int(column), String.New(tooltip))
}

/*
Returns the given column's tooltip text.
*/
func (self Instance) GetTooltipText(column int) string { //gd:TreeItem.get_tooltip_text
	return string(class(self).GetTooltipText(gd.Int(column)).String())
}

/*
Sets the given column's text alignment. See [enum HorizontalAlignment] for possible values.
*/
func (self Instance) SetTextAlignment(column int, text_alignment HorizontalAlignment) { //gd:TreeItem.set_text_alignment
	class(self).SetTextAlignment(gd.Int(column), text_alignment)
}

/*
Returns the given column's text alignment.
*/
func (self Instance) GetTextAlignment(column int) HorizontalAlignment { //gd:TreeItem.get_text_alignment
	return HorizontalAlignment(class(self).GetTextAlignment(gd.Int(column)))
}

/*
If [param enable] is [code]true[/code], the given [param column] is expanded to the right.
*/
func (self Instance) SetExpandRight(column int, enable bool) { //gd:TreeItem.set_expand_right
	class(self).SetExpandRight(gd.Int(column), enable)
}

/*
Returns [code]true[/code] if [code]expand_right[/code] is set.
*/
func (self Instance) GetExpandRight(column int) bool { //gd:TreeItem.get_expand_right
	return bool(class(self).GetExpandRight(gd.Int(column)))
}

/*
Creates an item and adds it as a child.
The new item will be inserted as position [param index] (the default value [code]-1[/code] means the last position), or it will be the last child if [param index] is higher than the child count.
*/
func (self Instance) CreateChild() [1]gdclass.TreeItem { //gd:TreeItem.create_child
	return [1]gdclass.TreeItem(class(self).CreateChild(gd.Int(-1)))
}

/*
Adds a previously unparented [TreeItem] as a direct child of this one. The [param child] item must not be a part of any [Tree] or parented to any [TreeItem]. See also [method remove_child].
*/
func (self Instance) AddChild(child [1]gdclass.TreeItem) { //gd:TreeItem.add_child
	class(self).AddChild(child)
}

/*
Removes the given child [TreeItem] and all its children from the [Tree]. Note that it doesn't free the item from memory, so it can be reused later (see [method add_child]). To completely remove a [TreeItem] use [method Object.free].
[b]Note:[/b] If you want to move a child from one [Tree] to another, then instead of removing and adding it manually you can use [method move_before] or [method move_after].
*/
func (self Instance) RemoveChild(child [1]gdclass.TreeItem) { //gd:TreeItem.remove_child
	class(self).RemoveChild(child)
}

/*
Returns the [Tree] that owns this TreeItem.
*/
func (self Instance) GetTree() [1]gdclass.Tree { //gd:TreeItem.get_tree
	return [1]gdclass.Tree(class(self).GetTree())
}

/*
Returns the next sibling TreeItem in the tree or a null object if there is none.
*/
func (self Instance) GetNext() [1]gdclass.TreeItem { //gd:TreeItem.get_next
	return [1]gdclass.TreeItem(class(self).GetNext())
}

/*
Returns the previous sibling TreeItem in the tree or a null object if there is none.
*/
func (self Instance) GetPrev() [1]gdclass.TreeItem { //gd:TreeItem.get_prev
	return [1]gdclass.TreeItem(class(self).GetPrev())
}

/*
Returns the parent TreeItem or a null object if there is none.
*/
func (self Instance) GetParent() [1]gdclass.TreeItem { //gd:TreeItem.get_parent
	return [1]gdclass.TreeItem(class(self).GetParent())
}

/*
Returns the TreeItem's first child.
*/
func (self Instance) GetFirstChild() [1]gdclass.TreeItem { //gd:TreeItem.get_first_child
	return [1]gdclass.TreeItem(class(self).GetFirstChild())
}

/*
Returns the next TreeItem in the tree (in the context of a depth-first search) or a [code]null[/code] object if there is none.
If [param wrap] is enabled, the method will wrap around to the first element in the tree when called on the last element, otherwise it returns [code]null[/code].
*/
func (self Instance) GetNextInTree() [1]gdclass.TreeItem { //gd:TreeItem.get_next_in_tree
	return [1]gdclass.TreeItem(class(self).GetNextInTree(false))
}

/*
Returns the previous TreeItem in the tree (in the context of a depth-first search) or a [code]null[/code] object if there is none.
If [param wrap] is enabled, the method will wrap around to the last element in the tree when called on the first visible element, otherwise it returns [code]null[/code].
*/
func (self Instance) GetPrevInTree() [1]gdclass.TreeItem { //gd:TreeItem.get_prev_in_tree
	return [1]gdclass.TreeItem(class(self).GetPrevInTree(false))
}

/*
Returns the next visible TreeItem in the tree (in the context of a depth-first search) or a [code]null[/code] object if there is none.
If [param wrap] is enabled, the method will wrap around to the first visible element in the tree when called on the last visible element, otherwise it returns [code]null[/code].
*/
func (self Instance) GetNextVisible() [1]gdclass.TreeItem { //gd:TreeItem.get_next_visible
	return [1]gdclass.TreeItem(class(self).GetNextVisible(false))
}

/*
Returns the previous visible sibling TreeItem in the tree (in the context of a depth-first search) or a [code]null[/code] object if there is none.
If [param wrap] is enabled, the method will wrap around to the last visible element in the tree when called on the first visible element, otherwise it returns [code]null[/code].
*/
func (self Instance) GetPrevVisible() [1]gdclass.TreeItem { //gd:TreeItem.get_prev_visible
	return [1]gdclass.TreeItem(class(self).GetPrevVisible(false))
}

/*
Returns a child item by its [param index] (see [method get_child_count]). This method is often used for iterating all children of an item.
Negative indices access the children from the last one.
*/
func (self Instance) GetChild(index int) [1]gdclass.TreeItem { //gd:TreeItem.get_child
	return [1]gdclass.TreeItem(class(self).GetChild(gd.Int(index)))
}

/*
Returns the number of child items.
*/
func (self Instance) GetChildCount() int { //gd:TreeItem.get_child_count
	return int(int(class(self).GetChildCount()))
}

/*
Returns an array of references to the item's children.
*/
func (self Instance) GetChildren() [][1]gdclass.TreeItem { //gd:TreeItem.get_children
	return [][1]gdclass.TreeItem(gd.ArrayAs[[][1]gdclass.TreeItem](gd.InternalArray(class(self).GetChildren())))
}

/*
Returns the node's order in the tree. For example, if called on the first child item the position is [code]0[/code].
*/
func (self Instance) GetIndex() int { //gd:TreeItem.get_index
	return int(int(class(self).GetIndex()))
}

/*
Moves this TreeItem right before the given [param item].
[b]Note:[/b] You can't move to the root or move the root.
*/
func (self Instance) MoveBefore(item [1]gdclass.TreeItem) { //gd:TreeItem.move_before
	class(self).MoveBefore(item)
}

/*
Moves this TreeItem right after the given [param item].
[b]Note:[/b] You can't move to the root or move the root.
*/
func (self Instance) MoveAfter(item [1]gdclass.TreeItem) { //gd:TreeItem.move_after
	class(self).MoveAfter(item)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.TreeItem

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TreeItem"))
	casted := Instance{*(*gdclass.TreeItem)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Collapsed() bool {
	return bool(class(self).IsCollapsed())
}

func (self Instance) SetCollapsed(value bool) {
	class(self).SetCollapsed(value)
}

func (self Instance) Visible() bool {
	return bool(class(self).IsVisible())
}

func (self Instance) SetVisible(value bool) {
	class(self).SetVisible(value)
}

func (self Instance) DisableFolding() bool {
	return bool(class(self).IsFoldingDisabled())
}

func (self Instance) SetDisableFolding(value bool) {
	class(self).SetDisableFolding(value)
}

func (self Instance) CustomMinimumHeight() int {
	return int(int(class(self).GetCustomMinimumHeight()))
}

func (self Instance) SetCustomMinimumHeight(value int) {
	class(self).SetCustomMinimumHeight(gd.Int(value))
}

/*
Sets the given column's cell mode to [param mode]. This determines how the cell is displayed and edited. See [enum TreeCellMode] constants for details.
*/
//go:nosplit
func (self class) SetCellMode(column gd.Int, mode gdclass.TreeItemTreeCellMode) { //gd:TreeItem.set_cell_mode
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_cell_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the column's cell mode.
*/
//go:nosplit
func (self class) GetCellMode(column gd.Int) gdclass.TreeItemTreeCellMode { //gd:TreeItem.get_cell_mode
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gdclass.TreeItemTreeCellMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_cell_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param multiline] is [code]true[/code], the given [param column] is multiline editable.
[b]Note:[/b] This option only affects the type of control ([LineEdit] or [TextEdit]) that appears when editing the column. You can set multiline values with [method set_text] even if the column is not multiline editable.
*/
//go:nosplit
func (self class) SetEditMultiline(column gd.Int, multiline bool) { //gd:TreeItem.set_edit_multiline
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, multiline)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_edit_multiline, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the given [param column] is multiline editable.
*/
//go:nosplit
func (self class) IsEditMultiline(column gd.Int) bool { //gd:TreeItem.is_edit_multiline
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_is_edit_multiline, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param checked] is [code]true[/code], the given [param column] is checked. Clears column's indeterminate status.
*/
//go:nosplit
func (self class) SetChecked(column gd.Int, checked bool) { //gd:TreeItem.set_checked
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, checked)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_checked, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
If [param indeterminate] is [code]true[/code], the given [param column] is marked indeterminate.
[b]Note:[/b] If set [code]true[/code] from [code]false[/code], then column is cleared of checked status.
*/
//go:nosplit
func (self class) SetIndeterminate(column gd.Int, indeterminate bool) { //gd:TreeItem.set_indeterminate
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, indeterminate)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_indeterminate, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the given [param column] is checked.
*/
//go:nosplit
func (self class) IsChecked(column gd.Int) bool { //gd:TreeItem.is_checked
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_is_checked, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the given [param column] is indeterminate.
*/
//go:nosplit
func (self class) IsIndeterminate(column gd.Int) bool { //gd:TreeItem.is_indeterminate
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_is_indeterminate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Propagates this item's checked status to its children and parents for the given [param column]. It is possible to process the items affected by this method call by connecting to [signal Tree.check_propagated_to_item]. The order that the items affected will be processed is as follows: the item invoking this method, children of that item, and finally parents of that item. If [param emit_signal] is [code]false[/code], then [signal Tree.check_propagated_to_item] will not be emitted.
*/
//go:nosplit
func (self class) PropagateCheck(column gd.Int, emit_signal bool) { //gd:TreeItem.propagate_check
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, emit_signal)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_propagate_check, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the given column's text value.
*/
//go:nosplit
func (self class) SetText(column gd.Int, text String.Readable) { //gd:TreeItem.set_text
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, pointers.Get(gd.InternalString(text)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the given column's text.
*/
//go:nosplit
func (self class) GetText(column gd.Int) String.Readable { //gd:TreeItem.get_text
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets item's text base writing direction.
*/
//go:nosplit
func (self class) SetTextDirection(column gd.Int, direction gdclass.ControlTextDirection) { //gd:TreeItem.set_text_direction
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, direction)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_text_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns item's text base writing direction.
*/
//go:nosplit
func (self class) GetTextDirection(column gd.Int) gdclass.ControlTextDirection { //gd:TreeItem.get_text_direction
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gdclass.ControlTextDirection](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_text_direction, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the autowrap mode in the given [param column]. If set to something other than [constant TextServer.AUTOWRAP_OFF], the text gets wrapped inside the cell's bounding rectangle.
*/
//go:nosplit
func (self class) SetAutowrapMode(column gd.Int, autowrap_mode gdclass.TextServerAutowrapMode) { //gd:TreeItem.set_autowrap_mode
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, autowrap_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_autowrap_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the text autowrap mode in the given [param column]. By default it is [constant TextServer.AUTOWRAP_OFF].
*/
//go:nosplit
func (self class) GetAutowrapMode(column gd.Int) gdclass.TextServerAutowrapMode { //gd:TreeItem.get_autowrap_mode
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gdclass.TextServerAutowrapMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_autowrap_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the clipping behavior when the text exceeds the item's bounding rectangle in the given [param column].
*/
//go:nosplit
func (self class) SetTextOverrunBehavior(column gd.Int, overrun_behavior gdclass.TextServerOverrunBehavior) { //gd:TreeItem.set_text_overrun_behavior
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, overrun_behavior)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the clipping behavior when the text exceeds the item's bounding rectangle in the given [param column]. By default it is [constant TextServer.OVERRUN_TRIM_ELLIPSIS].
*/
//go:nosplit
func (self class) GetTextOverrunBehavior(column gd.Int) gdclass.TextServerOverrunBehavior { //gd:TreeItem.get_text_overrun_behavior
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gdclass.TextServerOverrunBehavior](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_text_overrun_behavior, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set BiDi algorithm override for the structured text. Has effect for cells that display text.
*/
//go:nosplit
func (self class) SetStructuredTextBidiOverride(column gd.Int, parser gdclass.TextServerStructuredTextParser) { //gd:TreeItem.set_structured_text_bidi_override
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, parser)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the BiDi algorithm override set for this cell.
*/
//go:nosplit
func (self class) GetStructuredTextBidiOverride(column gd.Int) gdclass.TextServerStructuredTextParser { //gd:TreeItem.get_structured_text_bidi_override
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gdclass.TextServerStructuredTextParser](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_structured_text_bidi_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set additional options for BiDi override. Has effect for cells that display text.
*/
//go:nosplit
func (self class) SetStructuredTextBidiOverrideOptions(column gd.Int, args Array.Any) { //gd:TreeItem.set_structured_text_bidi_override_options
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, pointers.Get(gd.InternalArray(args)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the additional BiDi options set for this cell.
*/
//go:nosplit
func (self class) GetStructuredTextBidiOverrideOptions(column gd.Int) Array.Any { //gd:TreeItem.get_structured_text_bidi_override_options
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_structured_text_bidi_override_options, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets language code of item's text used for line-breaking and text shaping algorithms, if left empty current locale is used instead.
*/
//go:nosplit
func (self class) SetLanguage(column gd.Int, language String.Readable) { //gd:TreeItem.set_language
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, pointers.Get(gd.InternalString(language)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_language, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns item's text language code.
*/
//go:nosplit
func (self class) GetLanguage(column gd.Int) String.Readable { //gd:TreeItem.get_language
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_language, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets a string to be shown after a column's value (for example, a unit abbreviation).
*/
//go:nosplit
func (self class) SetSuffix(column gd.Int, text String.Readable) { //gd:TreeItem.set_suffix
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, pointers.Get(gd.InternalString(text)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_suffix, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Gets the suffix string shown after the column value.
*/
//go:nosplit
func (self class) GetSuffix(column gd.Int) String.Readable { //gd:TreeItem.get_suffix
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_suffix, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets the given cell's icon [Texture2D]. The cell has to be in [constant CELL_MODE_ICON] mode.
*/
//go:nosplit
func (self class) SetIcon(column gd.Int, texture [1]gdclass.Texture2D) { //gd:TreeItem.set_icon
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the given column's icon [Texture2D]. Error if no icon is set.
*/
//go:nosplit
func (self class) GetIcon(column gd.Int) [1]gdclass.Texture2D { //gd:TreeItem.get_icon
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sets the given column's icon's texture region.
*/
//go:nosplit
func (self class) SetIconRegion(column gd.Int, region gd.Rect2) { //gd:TreeItem.set_icon_region
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, region)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_icon_region, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the icon [Texture2D] region as [Rect2].
*/
//go:nosplit
func (self class) GetIconRegion(column gd.Int) gd.Rect2 { //gd:TreeItem.get_icon_region
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_icon_region, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the maximum allowed width of the icon in the given [param column]. This limit is applied on top of the default size of the icon and on top of [theme_item Tree.icon_max_width]. The height is adjusted according to the icon's ratio.
*/
//go:nosplit
func (self class) SetIconMaxWidth(column gd.Int, width gd.Int) { //gd:TreeItem.set_icon_max_width
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, width)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_icon_max_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the maximum allowed width of the icon in the given [param column].
*/
//go:nosplit
func (self class) GetIconMaxWidth(column gd.Int) gd.Int { //gd:TreeItem.get_icon_max_width
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_icon_max_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Modulates the given column's icon with [param modulate].
*/
//go:nosplit
func (self class) SetIconModulate(column gd.Int, modulate gd.Color) { //gd:TreeItem.set_icon_modulate
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, modulate)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_icon_modulate, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [Color] modulating the column's icon.
*/
//go:nosplit
func (self class) GetIconModulate(column gd.Int) gd.Color { //gd:TreeItem.get_icon_modulate
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_icon_modulate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the value of a [constant CELL_MODE_RANGE] column.
*/
//go:nosplit
func (self class) SetRange(column gd.Int, value gd.Float) { //gd:TreeItem.set_range
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the value of a [constant CELL_MODE_RANGE] column.
*/
//go:nosplit
func (self class) GetRange(column gd.Int) gd.Float { //gd:TreeItem.get_range
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the range of accepted values for a column. The column must be in the [constant CELL_MODE_RANGE] mode.
If [param expr] is [code]true[/code], the edit mode slider will use an exponential scale as with [member Range.exp_edit].
*/
//go:nosplit
func (self class) SetRangeConfig(column gd.Int, min gd.Float, max gd.Float, step gd.Float, expr bool) { //gd:TreeItem.set_range_config
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, min)
	callframe.Arg(frame, max)
	callframe.Arg(frame, step)
	callframe.Arg(frame, expr)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_range_config, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a dictionary containing the range parameters for a given column. The keys are "min", "max", "step", and "expr".
*/
//go:nosplit
func (self class) GetRangeConfig(column gd.Int) Dictionary.Any { //gd:TreeItem.get_range_config
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_range_config, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets the metadata value for the given column, which can be retrieved later using [method get_metadata]. This can be used, for example, to store a reference to the original data.
*/
//go:nosplit
func (self class) SetMetadata(column gd.Int, meta gd.Variant) { //gd:TreeItem.set_metadata
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, pointers.Get(meta))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_metadata, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the metadata value that was set for the given column using [method set_metadata].
*/
//go:nosplit
func (self class) GetMetadata(column gd.Int) gd.Variant { //gd:TreeItem.get_metadata
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_metadata, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the given column's custom draw callback to the [param callback] method on [param object].
The method named [param callback] should accept two arguments: the [TreeItem] that is drawn and its position and size as a [Rect2].
*/
//go:nosplit
func (self class) SetCustomDraw(column gd.Int, obj [1]gd.Object, callback gd.StringName) { //gd:TreeItem.set_custom_draw
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, pointers.Get(obj[0])[0])
	callframe.Arg(frame, pointers.Get(callback))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_custom_draw, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the given column's custom draw callback. Use an empty [Callable] ([code skip-lint]Callable()[/code]) to clear the custom callback. The cell has to be in [constant CELL_MODE_CUSTOM] to use this feature.
The [param callback] should accept two arguments: the [TreeItem] that is drawn and its position and size as a [Rect2].
*/
//go:nosplit
func (self class) SetCustomDrawCallback(column gd.Int, callback Callable.Function) { //gd:TreeItem.set_custom_draw_callback
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callback)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_custom_draw_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the custom callback of column [param column].
*/
//go:nosplit
func (self class) GetCustomDrawCallback(column gd.Int) Callable.Function { //gd:TreeItem.get_custom_draw_callback
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[[2]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_custom_draw_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Callable.Through(gd.CallableProxy{}, pointers.Pack(pointers.New[gd.Callable](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollapsed(enable bool) { //gd:TreeItem.set_collapsed
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_collapsed, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsCollapsed() bool { //gd:TreeItem.is_collapsed
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_is_collapsed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Collapses or uncollapses this [TreeItem] and all the descendants of this item.
*/
//go:nosplit
func (self class) SetCollapsedRecursive(enable bool) { //gd:TreeItem.set_collapsed_recursive
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_collapsed_recursive, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if this [TreeItem], or any of its descendants, is collapsed.
If [param only_visible] is [code]true[/code] it ignores non-visible [TreeItem]s.
*/
//go:nosplit
func (self class) IsAnyCollapsed(only_visible bool) bool { //gd:TreeItem.is_any_collapsed
	var frame = callframe.New()
	callframe.Arg(frame, only_visible)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_is_any_collapsed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisible(enable bool) { //gd:TreeItem.set_visible
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsVisible() bool { //gd:TreeItem.is_visible
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_is_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if [member visible] is [code]true[/code] and all its ancestors are also visible.
*/
//go:nosplit
func (self class) IsVisibleInTree() bool { //gd:TreeItem.is_visible_in_tree
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_is_visible_in_tree, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Uncollapses all [TreeItem]s necessary to reveal this [TreeItem], i.e. all ancestor [TreeItem]s.
*/
//go:nosplit
func (self class) UncollapseTree() { //gd:TreeItem.uncollapse_tree
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_uncollapse_tree, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetCustomMinimumHeight(height gd.Int) { //gd:TreeItem.set_custom_minimum_height
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_custom_minimum_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCustomMinimumHeight() gd.Int { //gd:TreeItem.get_custom_minimum_height
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_custom_minimum_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param selectable] is [code]true[/code], the given [param column] is selectable.
*/
//go:nosplit
func (self class) SetSelectable(column gd.Int, selectable bool) { //gd:TreeItem.set_selectable
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, selectable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_selectable, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the given [param column] is selectable.
*/
//go:nosplit
func (self class) IsSelectable(column gd.Int) bool { //gd:TreeItem.is_selectable
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_is_selectable, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the given [param column] is selected.
*/
//go:nosplit
func (self class) IsSelected(column gd.Int) bool { //gd:TreeItem.is_selected
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_is_selected, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Selects the given [param column].
*/
//go:nosplit
func (self class) Select(column gd.Int) { //gd:TreeItem.select_
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_select_, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Deselects the given column.
*/
//go:nosplit
func (self class) Deselect(column gd.Int) { //gd:TreeItem.deselect
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_deselect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
If [param enabled] is [code]true[/code], the given [param column] is editable.
*/
//go:nosplit
func (self class) SetEditable(column gd.Int, enabled bool) { //gd:TreeItem.set_editable
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_editable, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the given [param column] is editable.
*/
//go:nosplit
func (self class) IsEditable(column gd.Int) bool { //gd:TreeItem.is_editable
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_is_editable, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the given column's custom color.
*/
//go:nosplit
func (self class) SetCustomColor(column gd.Int, color gd.Color) { //gd:TreeItem.set_custom_color
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_custom_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the custom color of column [param column].
*/
//go:nosplit
func (self class) GetCustomColor(column gd.Int) gd.Color { //gd:TreeItem.get_custom_color
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_custom_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Resets the color for the given column to default.
*/
//go:nosplit
func (self class) ClearCustomColor(column gd.Int) { //gd:TreeItem.clear_custom_color
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_clear_custom_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets custom font used to draw text in the given [param column].
*/
//go:nosplit
func (self class) SetCustomFont(column gd.Int, font [1]gdclass.Font) { //gd:TreeItem.set_custom_font
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, pointers.Get(font[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_custom_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns custom font used to draw text in the column [param column].
*/
//go:nosplit
func (self class) GetCustomFont(column gd.Int) [1]gdclass.Font { //gd:TreeItem.get_custom_font
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_custom_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Font{gd.PointerWithOwnershipTransferredToGo[gdclass.Font](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sets custom font size used to draw text in the given [param column].
*/
//go:nosplit
func (self class) SetCustomFontSize(column gd.Int, font_size gd.Int) { //gd:TreeItem.set_custom_font_size
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, font_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_custom_font_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns custom font size used to draw text in the column [param column].
*/
//go:nosplit
func (self class) GetCustomFontSize(column gd.Int) gd.Int { //gd:TreeItem.get_custom_font_size
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_custom_font_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the given column's custom background color and whether to just use it as an outline.
*/
//go:nosplit
func (self class) SetCustomBgColor(column gd.Int, color gd.Color, just_outline bool) { //gd:TreeItem.set_custom_bg_color
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, color)
	callframe.Arg(frame, just_outline)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_custom_bg_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Resets the background color for the given column to default.
*/
//go:nosplit
func (self class) ClearCustomBgColor(column gd.Int) { //gd:TreeItem.clear_custom_bg_color
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_clear_custom_bg_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the custom background color of column [param column].
*/
//go:nosplit
func (self class) GetCustomBgColor(column gd.Int) gd.Color { //gd:TreeItem.get_custom_bg_color
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_custom_bg_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Makes a cell with [constant CELL_MODE_CUSTOM] display as a non-flat button with a [StyleBox].
*/
//go:nosplit
func (self class) SetCustomAsButton(column gd.Int, enable bool) { //gd:TreeItem.set_custom_as_button
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_custom_as_button, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the cell was made into a button with [method set_custom_as_button].
*/
//go:nosplit
func (self class) IsCustomSetAsButton(column gd.Int) bool { //gd:TreeItem.is_custom_set_as_button
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_is_custom_set_as_button, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a button with [Texture2D] [param button] at column [param column]. The [param id] is used to identify the button in the according [signal Tree.button_clicked] signal and can be different from the buttons index. If not specified, the next available index is used, which may be retrieved by calling [method get_button_count] immediately before this method. Optionally, the button can be [param disabled] and have a [param tooltip_text].
*/
//go:nosplit
func (self class) AddButton(column gd.Int, button [1]gdclass.Texture2D, id gd.Int, disabled bool, tooltip_text String.Readable) { //gd:TreeItem.add_button
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, pointers.Get(button[0])[0])
	callframe.Arg(frame, id)
	callframe.Arg(frame, disabled)
	callframe.Arg(frame, pointers.Get(gd.InternalString(tooltip_text)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_add_button, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the number of buttons in column [param column].
*/
//go:nosplit
func (self class) GetButtonCount(column gd.Int) gd.Int { //gd:TreeItem.get_button_count
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_button_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the tooltip text for the button at index [param button_index] in column [param column].
*/
//go:nosplit
func (self class) GetButtonTooltipText(column gd.Int, button_index gd.Int) String.Readable { //gd:TreeItem.get_button_tooltip_text
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_button_tooltip_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the ID for the button at index [param button_index] in column [param column].
*/
//go:nosplit
func (self class) GetButtonId(column gd.Int, button_index gd.Int) gd.Int { //gd:TreeItem.get_button_id
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_button_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the button index if there is a button with ID [param id] in column [param column], otherwise returns -1.
*/
//go:nosplit
func (self class) GetButtonById(column gd.Int, id gd.Int) gd.Int { //gd:TreeItem.get_button_by_id
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_button_by_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the color of the button with ID [param id] in column [param column]. If the specified button does not exist, returns [constant Color.BLACK].
*/
//go:nosplit
func (self class) GetButtonColor(column gd.Int, id gd.Int) gd.Color { //gd:TreeItem.get_button_color
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_button_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [Texture2D] of the button at index [param button_index] in column [param column].
*/
//go:nosplit
func (self class) GetButton(column gd.Int, button_index gd.Int) [1]gdclass.Texture2D { //gd:TreeItem.get_button
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_button, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sets the tooltip text for the button at index [param button_index] in the given [param column].
*/
//go:nosplit
func (self class) SetButtonTooltipText(column gd.Int, button_index gd.Int, tooltip String.Readable) { //gd:TreeItem.set_button_tooltip_text
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	callframe.Arg(frame, pointers.Get(gd.InternalString(tooltip)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_button_tooltip_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the given column's button [Texture2D] at index [param button_index] to [param button].
*/
//go:nosplit
func (self class) SetButton(column gd.Int, button_index gd.Int, button [1]gdclass.Texture2D) { //gd:TreeItem.set_button
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	callframe.Arg(frame, pointers.Get(button[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_button, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes the button at index [param button_index] in column [param column].
*/
//go:nosplit
func (self class) EraseButton(column gd.Int, button_index gd.Int) { //gd:TreeItem.erase_button
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_erase_button, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
If [code]true[/code], disables the button at index [param button_index] in the given [param column].
*/
//go:nosplit
func (self class) SetButtonDisabled(column gd.Int, button_index gd.Int, disabled bool) { //gd:TreeItem.set_button_disabled
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	callframe.Arg(frame, disabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_button_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the given column's button color at index [param button_index] to [param color].
*/
//go:nosplit
func (self class) SetButtonColor(column gd.Int, button_index gd.Int, color gd.Color) { //gd:TreeItem.set_button_color
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_button_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the button at index [param button_index] for the given [param column] is disabled.
*/
//go:nosplit
func (self class) IsButtonDisabled(column gd.Int, button_index gd.Int) bool { //gd:TreeItem.is_button_disabled
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, button_index)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_is_button_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the given column's tooltip text.
*/
//go:nosplit
func (self class) SetTooltipText(column gd.Int, tooltip String.Readable) { //gd:TreeItem.set_tooltip_text
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, pointers.Get(gd.InternalString(tooltip)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_tooltip_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the given column's tooltip text.
*/
//go:nosplit
func (self class) GetTooltipText(column gd.Int) String.Readable { //gd:TreeItem.get_tooltip_text
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_tooltip_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets the given column's text alignment. See [enum HorizontalAlignment] for possible values.
*/
//go:nosplit
func (self class) SetTextAlignment(column gd.Int, text_alignment HorizontalAlignment) { //gd:TreeItem.set_text_alignment
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, text_alignment)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_text_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the given column's text alignment.
*/
//go:nosplit
func (self class) GetTextAlignment(column gd.Int) HorizontalAlignment { //gd:TreeItem.get_text_alignment
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[HorizontalAlignment](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_text_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param enable] is [code]true[/code], the given [param column] is expanded to the right.
*/
//go:nosplit
func (self class) SetExpandRight(column gd.Int, enable bool) { //gd:TreeItem.set_expand_right
	var frame = callframe.New()
	callframe.Arg(frame, column)
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_expand_right, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if [code]expand_right[/code] is set.
*/
//go:nosplit
func (self class) GetExpandRight(column gd.Int) bool { //gd:TreeItem.get_expand_right
	var frame = callframe.New()
	callframe.Arg(frame, column)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_expand_right, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDisableFolding(disable bool) { //gd:TreeItem.set_disable_folding
	var frame = callframe.New()
	callframe.Arg(frame, disable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_set_disable_folding, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsFoldingDisabled() bool { //gd:TreeItem.is_folding_disabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_is_folding_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates an item and adds it as a child.
The new item will be inserted as position [param index] (the default value [code]-1[/code] means the last position), or it will be the last child if [param index] is higher than the child count.
*/
//go:nosplit
func (self class) CreateChild(index gd.Int) [1]gdclass.TreeItem { //gd:TreeItem.create_child
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_create_child, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TreeItem{gd.PointerMustAssertInstanceID[gdclass.TreeItem](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Adds a previously unparented [TreeItem] as a direct child of this one. The [param child] item must not be a part of any [Tree] or parented to any [TreeItem]. See also [method remove_child].
*/
//go:nosplit
func (self class) AddChild(child [1]gdclass.TreeItem) { //gd:TreeItem.add_child
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(child[0].AsObject()[0]))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_add_child, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes the given child [TreeItem] and all its children from the [Tree]. Note that it doesn't free the item from memory, so it can be reused later (see [method add_child]). To completely remove a [TreeItem] use [method Object.free].
[b]Note:[/b] If you want to move a child from one [Tree] to another, then instead of removing and adding it manually you can use [method move_before] or [method move_after].
*/
//go:nosplit
func (self class) RemoveChild(child [1]gdclass.TreeItem) { //gd:TreeItem.remove_child
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(child[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_remove_child, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [Tree] that owns this TreeItem.
*/
//go:nosplit
func (self class) GetTree() [1]gdclass.Tree { //gd:TreeItem.get_tree
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_tree, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Tree{gd.PointerMustAssertInstanceID[gdclass.Tree](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the next sibling TreeItem in the tree or a null object if there is none.
*/
//go:nosplit
func (self class) GetNext() [1]gdclass.TreeItem { //gd:TreeItem.get_next
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_next, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TreeItem{gd.PointerMustAssertInstanceID[gdclass.TreeItem](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the previous sibling TreeItem in the tree or a null object if there is none.
*/
//go:nosplit
func (self class) GetPrev() [1]gdclass.TreeItem { //gd:TreeItem.get_prev
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_prev, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TreeItem{gd.PointerMustAssertInstanceID[gdclass.TreeItem](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the parent TreeItem or a null object if there is none.
*/
//go:nosplit
func (self class) GetParent() [1]gdclass.TreeItem { //gd:TreeItem.get_parent
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_parent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TreeItem{gd.PointerMustAssertInstanceID[gdclass.TreeItem](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the TreeItem's first child.
*/
//go:nosplit
func (self class) GetFirstChild() [1]gdclass.TreeItem { //gd:TreeItem.get_first_child
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_first_child, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TreeItem{gd.PointerMustAssertInstanceID[gdclass.TreeItem](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the next TreeItem in the tree (in the context of a depth-first search) or a [code]null[/code] object if there is none.
If [param wrap] is enabled, the method will wrap around to the first element in the tree when called on the last element, otherwise it returns [code]null[/code].
*/
//go:nosplit
func (self class) GetNextInTree(wrap bool) [1]gdclass.TreeItem { //gd:TreeItem.get_next_in_tree
	var frame = callframe.New()
	callframe.Arg(frame, wrap)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_next_in_tree, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TreeItem{gd.PointerMustAssertInstanceID[gdclass.TreeItem](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the previous TreeItem in the tree (in the context of a depth-first search) or a [code]null[/code] object if there is none.
If [param wrap] is enabled, the method will wrap around to the last element in the tree when called on the first visible element, otherwise it returns [code]null[/code].
*/
//go:nosplit
func (self class) GetPrevInTree(wrap bool) [1]gdclass.TreeItem { //gd:TreeItem.get_prev_in_tree
	var frame = callframe.New()
	callframe.Arg(frame, wrap)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_prev_in_tree, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TreeItem{gd.PointerMustAssertInstanceID[gdclass.TreeItem](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the next visible TreeItem in the tree (in the context of a depth-first search) or a [code]null[/code] object if there is none.
If [param wrap] is enabled, the method will wrap around to the first visible element in the tree when called on the last visible element, otherwise it returns [code]null[/code].
*/
//go:nosplit
func (self class) GetNextVisible(wrap bool) [1]gdclass.TreeItem { //gd:TreeItem.get_next_visible
	var frame = callframe.New()
	callframe.Arg(frame, wrap)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_next_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TreeItem{gd.PointerMustAssertInstanceID[gdclass.TreeItem](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the previous visible sibling TreeItem in the tree (in the context of a depth-first search) or a [code]null[/code] object if there is none.
If [param wrap] is enabled, the method will wrap around to the last visible element in the tree when called on the first visible element, otherwise it returns [code]null[/code].
*/
//go:nosplit
func (self class) GetPrevVisible(wrap bool) [1]gdclass.TreeItem { //gd:TreeItem.get_prev_visible
	var frame = callframe.New()
	callframe.Arg(frame, wrap)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_prev_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TreeItem{gd.PointerMustAssertInstanceID[gdclass.TreeItem](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns a child item by its [param index] (see [method get_child_count]). This method is often used for iterating all children of an item.
Negative indices access the children from the last one.
*/
//go:nosplit
func (self class) GetChild(index gd.Int) [1]gdclass.TreeItem { //gd:TreeItem.get_child
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_child, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TreeItem{gd.PointerMustAssertInstanceID[gdclass.TreeItem](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the number of child items.
*/
//go:nosplit
func (self class) GetChildCount() gd.Int { //gd:TreeItem.get_child_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_child_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an array of references to the item's children.
*/
//go:nosplit
func (self class) GetChildren() Array.Contains[[1]gdclass.TreeItem] { //gd:TreeItem.get_children
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_children, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.TreeItem]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the node's order in the tree. For example, if called on the first child item the position is [code]0[/code].
*/
//go:nosplit
func (self class) GetIndex() gd.Int { //gd:TreeItem.get_index
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_get_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Moves this TreeItem right before the given [param item].
[b]Note:[/b] You can't move to the root or move the root.
*/
//go:nosplit
func (self class) MoveBefore(item [1]gdclass.TreeItem) { //gd:TreeItem.move_before
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(item[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_move_before, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Moves this TreeItem right after the given [param item].
[b]Note:[/b] You can't move to the root or move the root.
*/
//go:nosplit
func (self class) MoveAfter(item [1]gdclass.TreeItem) { //gd:TreeItem.move_after
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(item[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TreeItem.Bind_move_after, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsTreeItem() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTreeItem() Instance { return *((*Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Instance(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("TreeItem", func(ptr gd.Object) any { return [1]gdclass.TreeItem{*(*gdclass.TreeItem)(unsafe.Pointer(&ptr))} })
}

type TreeCellMode = gdclass.TreeItemTreeCellMode //gd:TreeItem.TreeCellMode

const (
	/*Cell shows a string label. When editable, the text can be edited using a [LineEdit], or a [TextEdit] popup if [method set_edit_multiline] is used.*/
	CellModeString TreeCellMode = 0
	/*Cell shows a checkbox, optionally with text. The checkbox can be pressed, released, or indeterminate (via [method set_indeterminate]). The checkbox can't be clicked unless the cell is editable.*/
	CellModeCheck TreeCellMode = 1
	/*Cell shows a numeric range. When editable, it can be edited using a range slider. Use [method set_range] to set the value and [method set_range_config] to configure the range.
	  This cell can also be used in a text dropdown mode when you assign a text with [method set_text]. Separate options with a comma, e.g. [code]"Option1,Option2,Option3"[/code].*/
	CellModeRange TreeCellMode = 2
	/*Cell shows an icon. It can't be edited nor display text.*/
	CellModeIcon TreeCellMode = 3
	/*Cell shows as a clickable button. It will display an arrow similar to [OptionButton], but doesn't feature a dropdown (for that you can use [constant CELL_MODE_RANGE]). Clicking the button emits the [signal Tree.item_edited] signal. The button is flat by default, you can use [method set_custom_as_button] to display it with a [StyleBox].
	  This mode also supports custom drawing using [method set_custom_draw_callback].*/
	CellModeCustom TreeCellMode = 4
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

type RangeConfig struct {
	Min  float32 `gd:"min"`
	Max  float32 `gd:"max"`
	Step float32 `gd:"step"`
	Expr string  `gd:"expr"`
}
