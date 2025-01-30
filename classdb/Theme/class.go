// Package Theme provides methods for working with Theme object instances.
package Theme

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Resource"
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
A resource used for styling/skinning [Control] and [Window] nodes. While individual controls can be styled using their local theme overrides (see [method Control.add_theme_color_override]), theme resources allow you to store and apply the same settings across all controls sharing the same type (e.g. style all [Button]s the same). One theme resource can be used for the entire project, but you can also set a separate theme resource to a branch of control nodes. A theme resource assigned to a control applies to the control itself, as well as all of its direct and indirect children (as long as a chain of controls is uninterrupted).
Use [member ProjectSettings.gui/theme/custom] to set up a project-scope theme that will be available to every control in your project.
Use [member Control.theme] of any control node to set up a theme that will be available to that control and all of its direct and indirect children.
*/
type Instance [1]gdclass.Theme

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsTheme() Instance
}

/*
Creates or changes the value of the icon property defined by [param name] and [param theme_type]. Use [method clear_icon] to remove the property.
*/
func (self Instance) SetIcon(name string, theme_type string, texture [1]gdclass.Texture2D) { //gd:Theme.set_icon
	class(self).SetIcon(String.Name(String.New(name)), String.Name(String.New(theme_type)), texture)
}

/*
Returns the icon property defined by [param name] and [param theme_type], if it exists.
Returns the engine fallback icon value if the property doesn't exist (see [member ThemeDB.fallback_icon]). Use [method has_icon] to check for existence.
*/
func (self Instance) GetIcon(name string, theme_type string) [1]gdclass.Texture2D { //gd:Theme.get_icon
	return [1]gdclass.Texture2D(class(self).GetIcon(String.Name(String.New(name)), String.Name(String.New(theme_type))))
}

/*
Returns [code]true[/code] if the icon property defined by [param name] and [param theme_type] exists.
Returns [code]false[/code] if it doesn't exist. Use [method set_icon] to define it.
*/
func (self Instance) HasIcon(name string, theme_type string) bool { //gd:Theme.has_icon
	return bool(class(self).HasIcon(String.Name(String.New(name)), String.Name(String.New(theme_type))))
}

/*
Renames the icon property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_icon] to check for existence, and [method clear_icon] to remove the existing property.
*/
func (self Instance) RenameIcon(old_name string, name string, theme_type string) { //gd:Theme.rename_icon
	class(self).RenameIcon(String.Name(String.New(old_name)), String.Name(String.New(name)), String.Name(String.New(theme_type)))
}

/*
Removes the icon property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_icon] to check for existence.
*/
func (self Instance) ClearIcon(name string, theme_type string) { //gd:Theme.clear_icon
	class(self).ClearIcon(String.Name(String.New(name)), String.Name(String.New(theme_type)))
}

/*
Returns a list of names for icon properties defined with [param theme_type]. Use [method get_icon_type_list] to get a list of possible theme type names.
*/
func (self Instance) GetIconList(theme_type string) []string { //gd:Theme.get_icon_list
	return []string(class(self).GetIconList(String.New(theme_type)).Strings())
}

/*
Returns a list of all unique theme type names for icon properties. Use [method get_type_list] to get a list of all unique theme types.
*/
func (self Instance) GetIconTypeList() []string { //gd:Theme.get_icon_type_list
	return []string(class(self).GetIconTypeList().Strings())
}

/*
Creates or changes the value of the [StyleBox] property defined by [param name] and [param theme_type]. Use [method clear_stylebox] to remove the property.
*/
func (self Instance) SetStylebox(name string, theme_type string, texture [1]gdclass.StyleBox) { //gd:Theme.set_stylebox
	class(self).SetStylebox(String.Name(String.New(name)), String.Name(String.New(theme_type)), texture)
}

/*
Returns the [StyleBox] property defined by [param name] and [param theme_type], if it exists.
Returns the engine fallback stylebox value if the property doesn't exist (see [member ThemeDB.fallback_stylebox]). Use [method has_stylebox] to check for existence.
*/
func (self Instance) GetStylebox(name string, theme_type string) [1]gdclass.StyleBox { //gd:Theme.get_stylebox
	return [1]gdclass.StyleBox(class(self).GetStylebox(String.Name(String.New(name)), String.Name(String.New(theme_type))))
}

/*
Returns [code]true[/code] if the [StyleBox] property defined by [param name] and [param theme_type] exists.
Returns [code]false[/code] if it doesn't exist. Use [method set_stylebox] to define it.
*/
func (self Instance) HasStylebox(name string, theme_type string) bool { //gd:Theme.has_stylebox
	return bool(class(self).HasStylebox(String.Name(String.New(name)), String.Name(String.New(theme_type))))
}

/*
Renames the [StyleBox] property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_stylebox] to check for existence, and [method clear_stylebox] to remove the existing property.
*/
func (self Instance) RenameStylebox(old_name string, name string, theme_type string) { //gd:Theme.rename_stylebox
	class(self).RenameStylebox(String.Name(String.New(old_name)), String.Name(String.New(name)), String.Name(String.New(theme_type)))
}

/*
Removes the [StyleBox] property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_stylebox] to check for existence.
*/
func (self Instance) ClearStylebox(name string, theme_type string) { //gd:Theme.clear_stylebox
	class(self).ClearStylebox(String.Name(String.New(name)), String.Name(String.New(theme_type)))
}

/*
Returns a list of names for [StyleBox] properties defined with [param theme_type]. Use [method get_stylebox_type_list] to get a list of possible theme type names.
*/
func (self Instance) GetStyleboxList(theme_type string) []string { //gd:Theme.get_stylebox_list
	return []string(class(self).GetStyleboxList(String.New(theme_type)).Strings())
}

/*
Returns a list of all unique theme type names for [StyleBox] properties. Use [method get_type_list] to get a list of all unique theme types.
*/
func (self Instance) GetStyleboxTypeList() []string { //gd:Theme.get_stylebox_type_list
	return []string(class(self).GetStyleboxTypeList().Strings())
}

/*
Creates or changes the value of the [Font] property defined by [param name] and [param theme_type]. Use [method clear_font] to remove the property.
*/
func (self Instance) SetFont(name string, theme_type string, font [1]gdclass.Font) { //gd:Theme.set_font
	class(self).SetFont(String.Name(String.New(name)), String.Name(String.New(theme_type)), font)
}

/*
Returns the [Font] property defined by [param name] and [param theme_type], if it exists.
Returns the default theme font if the property doesn't exist and the default theme font is set up (see [member default_font]). Use [method has_font] to check for existence of the property and [method has_default_font] to check for existence of the default theme font.
Returns the engine fallback font value, if neither exist (see [member ThemeDB.fallback_font]).
*/
func (self Instance) GetFont(name string, theme_type string) [1]gdclass.Font { //gd:Theme.get_font
	return [1]gdclass.Font(class(self).GetFont(String.Name(String.New(name)), String.Name(String.New(theme_type))))
}

/*
Returns [code]true[/code] if the [Font] property defined by [param name] and [param theme_type] exists, or if the default theme font is set up (see [method has_default_font]).
Returns [code]false[/code] if neither exist. Use [method set_font] to define the property.
*/
func (self Instance) HasFont(name string, theme_type string) bool { //gd:Theme.has_font
	return bool(class(self).HasFont(String.Name(String.New(name)), String.Name(String.New(theme_type))))
}

/*
Renames the [Font] property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_font] to check for existence, and [method clear_font] to remove the existing property.
*/
func (self Instance) RenameFont(old_name string, name string, theme_type string) { //gd:Theme.rename_font
	class(self).RenameFont(String.Name(String.New(old_name)), String.Name(String.New(name)), String.Name(String.New(theme_type)))
}

/*
Removes the [Font] property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_font] to check for existence.
*/
func (self Instance) ClearFont(name string, theme_type string) { //gd:Theme.clear_font
	class(self).ClearFont(String.Name(String.New(name)), String.Name(String.New(theme_type)))
}

/*
Returns a list of names for [Font] properties defined with [param theme_type]. Use [method get_font_type_list] to get a list of possible theme type names.
*/
func (self Instance) GetFontList(theme_type string) []string { //gd:Theme.get_font_list
	return []string(class(self).GetFontList(String.New(theme_type)).Strings())
}

/*
Returns a list of all unique theme type names for [Font] properties. Use [method get_type_list] to get a list of all unique theme types.
*/
func (self Instance) GetFontTypeList() []string { //gd:Theme.get_font_type_list
	return []string(class(self).GetFontTypeList().Strings())
}

/*
Creates or changes the value of the font size property defined by [param name] and [param theme_type]. Use [method clear_font_size] to remove the property.
*/
func (self Instance) SetFontSize(name string, theme_type string, font_size int) { //gd:Theme.set_font_size
	class(self).SetFontSize(String.Name(String.New(name)), String.Name(String.New(theme_type)), int64(font_size))
}

/*
Returns the font size property defined by [param name] and [param theme_type], if it exists.
Returns the default theme font size if the property doesn't exist and the default theme font size is set up (see [member default_font_size]). Use [method has_font_size] to check for existence of the property and [method has_default_font_size] to check for existence of the default theme font.
Returns the engine fallback font size value, if neither exist (see [member ThemeDB.fallback_font_size]).
*/
func (self Instance) GetFontSize(name string, theme_type string) int { //gd:Theme.get_font_size
	return int(int(class(self).GetFontSize(String.Name(String.New(name)), String.Name(String.New(theme_type)))))
}

/*
Returns [code]true[/code] if the font size property defined by [param name] and [param theme_type] exists, or if the default theme font size is set up (see [method has_default_font_size]).
Returns [code]false[/code] if neither exist. Use [method set_font_size] to define the property.
*/
func (self Instance) HasFontSize(name string, theme_type string) bool { //gd:Theme.has_font_size
	return bool(class(self).HasFontSize(String.Name(String.New(name)), String.Name(String.New(theme_type))))
}

/*
Renames the font size property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_font_size] to check for existence, and [method clear_font_size] to remove the existing property.
*/
func (self Instance) RenameFontSize(old_name string, name string, theme_type string) { //gd:Theme.rename_font_size
	class(self).RenameFontSize(String.Name(String.New(old_name)), String.Name(String.New(name)), String.Name(String.New(theme_type)))
}

/*
Removes the font size property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_font_size] to check for existence.
*/
func (self Instance) ClearFontSize(name string, theme_type string) { //gd:Theme.clear_font_size
	class(self).ClearFontSize(String.Name(String.New(name)), String.Name(String.New(theme_type)))
}

/*
Returns a list of names for font size properties defined with [param theme_type]. Use [method get_font_size_type_list] to get a list of possible theme type names.
*/
func (self Instance) GetFontSizeList(theme_type string) []string { //gd:Theme.get_font_size_list
	return []string(class(self).GetFontSizeList(String.New(theme_type)).Strings())
}

/*
Returns a list of all unique theme type names for font size properties. Use [method get_type_list] to get a list of all unique theme types.
*/
func (self Instance) GetFontSizeTypeList() []string { //gd:Theme.get_font_size_type_list
	return []string(class(self).GetFontSizeTypeList().Strings())
}

/*
Creates or changes the value of the [Color] property defined by [param name] and [param theme_type]. Use [method clear_color] to remove the property.
*/
func (self Instance) SetColor(name string, theme_type string, color Color.RGBA) { //gd:Theme.set_color
	class(self).SetColor(String.Name(String.New(name)), String.Name(String.New(theme_type)), Color.RGBA(color))
}

/*
Returns the [Color] property defined by [param name] and [param theme_type], if it exists.
Returns the default color value if the property doesn't exist. Use [method has_color] to check for existence.
*/
func (self Instance) GetColor(name string, theme_type string) Color.RGBA { //gd:Theme.get_color
	return Color.RGBA(class(self).GetColor(String.Name(String.New(name)), String.Name(String.New(theme_type))))
}

/*
Returns [code]true[/code] if the [Color] property defined by [param name] and [param theme_type] exists.
Returns [code]false[/code] if it doesn't exist. Use [method set_color] to define it.
*/
func (self Instance) HasColor(name string, theme_type string) bool { //gd:Theme.has_color
	return bool(class(self).HasColor(String.Name(String.New(name)), String.Name(String.New(theme_type))))
}

/*
Renames the [Color] property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_color] to check for existence, and [method clear_color] to remove the existing property.
*/
func (self Instance) RenameColor(old_name string, name string, theme_type string) { //gd:Theme.rename_color
	class(self).RenameColor(String.Name(String.New(old_name)), String.Name(String.New(name)), String.Name(String.New(theme_type)))
}

/*
Removes the [Color] property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_color] to check for existence.
*/
func (self Instance) ClearColor(name string, theme_type string) { //gd:Theme.clear_color
	class(self).ClearColor(String.Name(String.New(name)), String.Name(String.New(theme_type)))
}

/*
Returns a list of names for [Color] properties defined with [param theme_type]. Use [method get_color_type_list] to get a list of possible theme type names.
*/
func (self Instance) GetColorList(theme_type string) []string { //gd:Theme.get_color_list
	return []string(class(self).GetColorList(String.New(theme_type)).Strings())
}

/*
Returns a list of all unique theme type names for [Color] properties. Use [method get_type_list] to get a list of all unique theme types.
*/
func (self Instance) GetColorTypeList() []string { //gd:Theme.get_color_type_list
	return []string(class(self).GetColorTypeList().Strings())
}

/*
Creates or changes the value of the constant property defined by [param name] and [param theme_type]. Use [method clear_constant] to remove the property.
*/
func (self Instance) SetConstant(name string, theme_type string, constant int) { //gd:Theme.set_constant
	class(self).SetConstant(String.Name(String.New(name)), String.Name(String.New(theme_type)), int64(constant))
}

/*
Returns the constant property defined by [param name] and [param theme_type], if it exists.
Returns [code]0[/code] if the property doesn't exist. Use [method has_constant] to check for existence.
*/
func (self Instance) GetConstant(name string, theme_type string) int { //gd:Theme.get_constant
	return int(int(class(self).GetConstant(String.Name(String.New(name)), String.Name(String.New(theme_type)))))
}

/*
Returns [code]true[/code] if the constant property defined by [param name] and [param theme_type] exists.
Returns [code]false[/code] if it doesn't exist. Use [method set_constant] to define it.
*/
func (self Instance) HasConstant(name string, theme_type string) bool { //gd:Theme.has_constant
	return bool(class(self).HasConstant(String.Name(String.New(name)), String.Name(String.New(theme_type))))
}

/*
Renames the constant property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_constant] to check for existence, and [method clear_constant] to remove the existing property.
*/
func (self Instance) RenameConstant(old_name string, name string, theme_type string) { //gd:Theme.rename_constant
	class(self).RenameConstant(String.Name(String.New(old_name)), String.Name(String.New(name)), String.Name(String.New(theme_type)))
}

/*
Removes the constant property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_constant] to check for existence.
*/
func (self Instance) ClearConstant(name string, theme_type string) { //gd:Theme.clear_constant
	class(self).ClearConstant(String.Name(String.New(name)), String.Name(String.New(theme_type)))
}

/*
Returns a list of names for constant properties defined with [param theme_type]. Use [method get_constant_type_list] to get a list of possible theme type names.
*/
func (self Instance) GetConstantList(theme_type string) []string { //gd:Theme.get_constant_list
	return []string(class(self).GetConstantList(String.New(theme_type)).Strings())
}

/*
Returns a list of all unique theme type names for constant properties. Use [method get_type_list] to get a list of all unique theme types.
*/
func (self Instance) GetConstantTypeList() []string { //gd:Theme.get_constant_type_list
	return []string(class(self).GetConstantTypeList().Strings())
}

/*
Returns [code]true[/code] if [member default_base_scale] has a valid value.
Returns [code]false[/code] if it doesn't. The value must be greater than [code]0.0[/code] to be considered valid.
*/
func (self Instance) HasDefaultBaseScale() bool { //gd:Theme.has_default_base_scale
	return bool(class(self).HasDefaultBaseScale())
}

/*
Returns [code]true[/code] if [member default_font] has a valid value.
Returns [code]false[/code] if it doesn't.
*/
func (self Instance) HasDefaultFont() bool { //gd:Theme.has_default_font
	return bool(class(self).HasDefaultFont())
}

/*
Returns [code]true[/code] if [member default_font_size] has a valid value.
Returns [code]false[/code] if it doesn't. The value must be greater than [code]0[/code] to be considered valid.
*/
func (self Instance) HasDefaultFontSize() bool { //gd:Theme.has_default_font_size
	return bool(class(self).HasDefaultFontSize())
}

/*
Creates or changes the value of the theme property of [param data_type] defined by [param name] and [param theme_type]. Use [method clear_theme_item] to remove the property.
Fails if the [param value] type is not accepted by [param data_type].
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
func (self Instance) SetThemeItem(data_type gdclass.ThemeDataType, name string, theme_type string, value any) { //gd:Theme.set_theme_item
	class(self).SetThemeItem(data_type, String.Name(String.New(name)), String.Name(String.New(theme_type)), variant.New(value))
}

/*
Returns the theme property of [param data_type] defined by [param name] and [param theme_type], if it exists.
Returns the engine fallback value if the property doesn't exist (see [ThemeDB]). Use [method has_theme_item] to check for existence.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
func (self Instance) GetThemeItem(data_type gdclass.ThemeDataType, name string, theme_type string) any { //gd:Theme.get_theme_item
	return any(class(self).GetThemeItem(data_type, String.Name(String.New(name)), String.Name(String.New(theme_type))).Interface())
}

/*
Returns [code]true[/code] if the theme property of [param data_type] defined by [param name] and [param theme_type] exists.
Returns [code]false[/code] if it doesn't exist. Use [method set_theme_item] to define it.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
func (self Instance) HasThemeItem(data_type gdclass.ThemeDataType, name string, theme_type string) bool { //gd:Theme.has_theme_item
	return bool(class(self).HasThemeItem(data_type, String.Name(String.New(name)), String.Name(String.New(theme_type))))
}

/*
Renames the theme property of [param data_type] defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_theme_item] to check for existence, and [method clear_theme_item] to remove the existing property.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
func (self Instance) RenameThemeItem(data_type gdclass.ThemeDataType, old_name string, name string, theme_type string) { //gd:Theme.rename_theme_item
	class(self).RenameThemeItem(data_type, String.Name(String.New(old_name)), String.Name(String.New(name)), String.Name(String.New(theme_type)))
}

/*
Removes the theme property of [param data_type] defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_theme_item] to check for existence.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
func (self Instance) ClearThemeItem(data_type gdclass.ThemeDataType, name string, theme_type string) { //gd:Theme.clear_theme_item
	class(self).ClearThemeItem(data_type, String.Name(String.New(name)), String.Name(String.New(theme_type)))
}

/*
Returns a list of names for properties of [param data_type] defined with [param theme_type]. Use [method get_theme_item_type_list] to get a list of possible theme type names.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
func (self Instance) GetThemeItemList(data_type gdclass.ThemeDataType, theme_type string) []string { //gd:Theme.get_theme_item_list
	return []string(class(self).GetThemeItemList(data_type, String.New(theme_type)).Strings())
}

/*
Returns a list of all unique theme type names for [param data_type] properties. Use [method get_type_list] to get a list of all unique theme types.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
func (self Instance) GetThemeItemTypeList(data_type gdclass.ThemeDataType) []string { //gd:Theme.get_theme_item_type_list
	return []string(class(self).GetThemeItemTypeList(data_type).Strings())
}

/*
Marks [param theme_type] as a variation of [param base_type].
This adds [param theme_type] as a suggested option for [member Control.theme_type_variation] on a [Control] that is of the [param base_type] class.
Variations can also be nested, i.e. [param base_type] can be another variation. If a chain of variations ends with a [param base_type] matching the class of the [Control], the whole chain is going to be suggested as options.
[b]Note:[/b] Suggestions only show up if this theme resource is set as the project default theme. See [member ProjectSettings.gui/theme/custom].
*/
func (self Instance) SetTypeVariation(theme_type string, base_type string) { //gd:Theme.set_type_variation
	class(self).SetTypeVariation(String.Name(String.New(theme_type)), String.Name(String.New(base_type)))
}

/*
Returns [code]true[/code] if [param theme_type] is marked as a variation of [param base_type].
*/
func (self Instance) IsTypeVariation(theme_type string, base_type string) bool { //gd:Theme.is_type_variation
	return bool(class(self).IsTypeVariation(String.Name(String.New(theme_type)), String.Name(String.New(base_type))))
}

/*
Unmarks [param theme_type] as being a variation of another theme type. See [method set_type_variation].
*/
func (self Instance) ClearTypeVariation(theme_type string) { //gd:Theme.clear_type_variation
	class(self).ClearTypeVariation(String.Name(String.New(theme_type)))
}

/*
Returns the name of the base theme type if [param theme_type] is a valid variation type. Returns an empty string otherwise.
*/
func (self Instance) GetTypeVariationBase(theme_type string) string { //gd:Theme.get_type_variation_base
	return string(class(self).GetTypeVariationBase(String.Name(String.New(theme_type))).String())
}

/*
Returns a list of all type variations for the given [param base_type].
*/
func (self Instance) GetTypeVariationList(base_type string) []string { //gd:Theme.get_type_variation_list
	return []string(class(self).GetTypeVariationList(String.Name(String.New(base_type))).Strings())
}

/*
Adds an empty theme type for every valid data type.
[b]Note:[/b] Empty types are not saved with the theme. This method only exists to perform in-memory changes to the resource. Use available [code]set_*[/code] methods to add theme items.
*/
func (self Instance) AddType(theme_type string) { //gd:Theme.add_type
	class(self).AddType(String.Name(String.New(theme_type)))
}

/*
Removes the theme type, gracefully discarding defined theme items. If the type is a variation, this information is also erased. If the type is a base for type variations, those variations lose their base.
*/
func (self Instance) RemoveType(theme_type string) { //gd:Theme.remove_type
	class(self).RemoveType(String.Name(String.New(theme_type)))
}

/*
Returns a list of all unique theme type names. Use the appropriate [code]get_*_type_list[/code] method to get a list of unique theme types for a single data type.
*/
func (self Instance) GetTypeList() []string { //gd:Theme.get_type_list
	return []string(class(self).GetTypeList().Strings())
}

/*
Adds missing and overrides existing definitions with values from the [param other] theme resource.
[b]Note:[/b] This modifies the current theme. If you want to merge two themes together without modifying either one, create a new empty theme and merge the other two into it one after another.
*/
func (self Instance) MergeWith(other [1]gdclass.Theme) { //gd:Theme.merge_with
	class(self).MergeWith(other)
}

/*
Removes all the theme properties defined on the theme resource.
*/
func (self Instance) Clear() { //gd:Theme.clear
	class(self).Clear()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Theme

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Theme"))
	casted := Instance{*(*gdclass.Theme)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) DefaultBaseScale() Float.X {
	return Float.X(Float.X(class(self).GetDefaultBaseScale()))
}

func (self Instance) SetDefaultBaseScale(value Float.X) {
	class(self).SetDefaultBaseScale(float64(value))
}

func (self Instance) DefaultFont() [1]gdclass.Font {
	return [1]gdclass.Font(class(self).GetDefaultFont())
}

func (self Instance) SetDefaultFont(value [1]gdclass.Font) {
	class(self).SetDefaultFont(value)
}

func (self Instance) DefaultFontSize() int {
	return int(int(class(self).GetDefaultFontSize()))
}

func (self Instance) SetDefaultFontSize(value int) {
	class(self).SetDefaultFontSize(int64(value))
}

/*
Creates or changes the value of the icon property defined by [param name] and [param theme_type]. Use [method clear_icon] to remove the property.
*/
//go:nosplit
func (self class) SetIcon(name String.Name, theme_type String.Name, texture [1]gdclass.Texture2D) { //gd:Theme.set_icon
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_set_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the icon property defined by [param name] and [param theme_type], if it exists.
Returns the engine fallback icon value if the property doesn't exist (see [member ThemeDB.fallback_icon]). Use [method has_icon] to check for existence.
*/
//go:nosplit
func (self class) GetIcon(name String.Name, theme_type String.Name) [1]gdclass.Texture2D { //gd:Theme.get_icon
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the icon property defined by [param name] and [param theme_type] exists.
Returns [code]false[/code] if it doesn't exist. Use [method set_icon] to define it.
*/
//go:nosplit
func (self class) HasIcon(name String.Name, theme_type String.Name) bool { //gd:Theme.has_icon
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_has_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Renames the icon property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_icon] to check for existence, and [method clear_icon] to remove the existing property.
*/
//go:nosplit
func (self class) RenameIcon(old_name String.Name, name String.Name, theme_type String.Name) { //gd:Theme.rename_icon
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(old_name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_rename_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes the icon property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_icon] to check for existence.
*/
//go:nosplit
func (self class) ClearIcon(name String.Name, theme_type String.Name) { //gd:Theme.clear_icon
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_clear_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a list of names for icon properties defined with [param theme_type]. Use [method get_icon_type_list] to get a list of possible theme type names.
*/
//go:nosplit
func (self class) GetIconList(theme_type String.Readable) Packed.Strings { //gd:Theme.get_icon_list
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(theme_type)))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_icon_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns a list of all unique theme type names for icon properties. Use [method get_type_list] to get a list of all unique theme types.
*/
//go:nosplit
func (self class) GetIconTypeList() Packed.Strings { //gd:Theme.get_icon_type_list
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_icon_type_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Creates or changes the value of the [StyleBox] property defined by [param name] and [param theme_type]. Use [method clear_stylebox] to remove the property.
*/
//go:nosplit
func (self class) SetStylebox(name String.Name, theme_type String.Name, texture [1]gdclass.StyleBox) { //gd:Theme.set_stylebox
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_set_stylebox, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [StyleBox] property defined by [param name] and [param theme_type], if it exists.
Returns the engine fallback stylebox value if the property doesn't exist (see [member ThemeDB.fallback_stylebox]). Use [method has_stylebox] to check for existence.
*/
//go:nosplit
func (self class) GetStylebox(name String.Name, theme_type String.Name) [1]gdclass.StyleBox { //gd:Theme.get_stylebox
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_stylebox, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.StyleBox{gd.PointerWithOwnershipTransferredToGo[gdclass.StyleBox](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the [StyleBox] property defined by [param name] and [param theme_type] exists.
Returns [code]false[/code] if it doesn't exist. Use [method set_stylebox] to define it.
*/
//go:nosplit
func (self class) HasStylebox(name String.Name, theme_type String.Name) bool { //gd:Theme.has_stylebox
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_has_stylebox, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Renames the [StyleBox] property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_stylebox] to check for existence, and [method clear_stylebox] to remove the existing property.
*/
//go:nosplit
func (self class) RenameStylebox(old_name String.Name, name String.Name, theme_type String.Name) { //gd:Theme.rename_stylebox
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(old_name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_rename_stylebox, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes the [StyleBox] property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_stylebox] to check for existence.
*/
//go:nosplit
func (self class) ClearStylebox(name String.Name, theme_type String.Name) { //gd:Theme.clear_stylebox
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_clear_stylebox, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a list of names for [StyleBox] properties defined with [param theme_type]. Use [method get_stylebox_type_list] to get a list of possible theme type names.
*/
//go:nosplit
func (self class) GetStyleboxList(theme_type String.Readable) Packed.Strings { //gd:Theme.get_stylebox_list
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(theme_type)))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_stylebox_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns a list of all unique theme type names for [StyleBox] properties. Use [method get_type_list] to get a list of all unique theme types.
*/
//go:nosplit
func (self class) GetStyleboxTypeList() Packed.Strings { //gd:Theme.get_stylebox_type_list
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_stylebox_type_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Creates or changes the value of the [Font] property defined by [param name] and [param theme_type]. Use [method clear_font] to remove the property.
*/
//go:nosplit
func (self class) SetFont(name String.Name, theme_type String.Name, font [1]gdclass.Font) { //gd:Theme.set_font
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	callframe.Arg(frame, pointers.Get(font[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_set_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [Font] property defined by [param name] and [param theme_type], if it exists.
Returns the default theme font if the property doesn't exist and the default theme font is set up (see [member default_font]). Use [method has_font] to check for existence of the property and [method has_default_font] to check for existence of the default theme font.
Returns the engine fallback font value, if neither exist (see [member ThemeDB.fallback_font]).
*/
//go:nosplit
func (self class) GetFont(name String.Name, theme_type String.Name) [1]gdclass.Font { //gd:Theme.get_font
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Font{gd.PointerWithOwnershipTransferredToGo[gdclass.Font](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the [Font] property defined by [param name] and [param theme_type] exists, or if the default theme font is set up (see [method has_default_font]).
Returns [code]false[/code] if neither exist. Use [method set_font] to define the property.
*/
//go:nosplit
func (self class) HasFont(name String.Name, theme_type String.Name) bool { //gd:Theme.has_font
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_has_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Renames the [Font] property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_font] to check for existence, and [method clear_font] to remove the existing property.
*/
//go:nosplit
func (self class) RenameFont(old_name String.Name, name String.Name, theme_type String.Name) { //gd:Theme.rename_font
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(old_name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_rename_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes the [Font] property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_font] to check for existence.
*/
//go:nosplit
func (self class) ClearFont(name String.Name, theme_type String.Name) { //gd:Theme.clear_font
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_clear_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a list of names for [Font] properties defined with [param theme_type]. Use [method get_font_type_list] to get a list of possible theme type names.
*/
//go:nosplit
func (self class) GetFontList(theme_type String.Readable) Packed.Strings { //gd:Theme.get_font_list
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(theme_type)))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_font_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns a list of all unique theme type names for [Font] properties. Use [method get_type_list] to get a list of all unique theme types.
*/
//go:nosplit
func (self class) GetFontTypeList() Packed.Strings { //gd:Theme.get_font_type_list
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_font_type_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Creates or changes the value of the font size property defined by [param name] and [param theme_type]. Use [method clear_font_size] to remove the property.
*/
//go:nosplit
func (self class) SetFontSize(name String.Name, theme_type String.Name, font_size int64) { //gd:Theme.set_font_size
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	callframe.Arg(frame, font_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_set_font_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the font size property defined by [param name] and [param theme_type], if it exists.
Returns the default theme font size if the property doesn't exist and the default theme font size is set up (see [member default_font_size]). Use [method has_font_size] to check for existence of the property and [method has_default_font_size] to check for existence of the default theme font.
Returns the engine fallback font size value, if neither exist (see [member ThemeDB.fallback_font_size]).
*/
//go:nosplit
func (self class) GetFontSize(name String.Name, theme_type String.Name) int64 { //gd:Theme.get_font_size
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_font_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the font size property defined by [param name] and [param theme_type] exists, or if the default theme font size is set up (see [method has_default_font_size]).
Returns [code]false[/code] if neither exist. Use [method set_font_size] to define the property.
*/
//go:nosplit
func (self class) HasFontSize(name String.Name, theme_type String.Name) bool { //gd:Theme.has_font_size
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_has_font_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Renames the font size property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_font_size] to check for existence, and [method clear_font_size] to remove the existing property.
*/
//go:nosplit
func (self class) RenameFontSize(old_name String.Name, name String.Name, theme_type String.Name) { //gd:Theme.rename_font_size
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(old_name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_rename_font_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes the font size property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_font_size] to check for existence.
*/
//go:nosplit
func (self class) ClearFontSize(name String.Name, theme_type String.Name) { //gd:Theme.clear_font_size
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_clear_font_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a list of names for font size properties defined with [param theme_type]. Use [method get_font_size_type_list] to get a list of possible theme type names.
*/
//go:nosplit
func (self class) GetFontSizeList(theme_type String.Readable) Packed.Strings { //gd:Theme.get_font_size_list
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(theme_type)))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_font_size_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns a list of all unique theme type names for font size properties. Use [method get_type_list] to get a list of all unique theme types.
*/
//go:nosplit
func (self class) GetFontSizeTypeList() Packed.Strings { //gd:Theme.get_font_size_type_list
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_font_size_type_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Creates or changes the value of the [Color] property defined by [param name] and [param theme_type]. Use [method clear_color] to remove the property.
*/
//go:nosplit
func (self class) SetColor(name String.Name, theme_type String.Name, color Color.RGBA) { //gd:Theme.set_color
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [Color] property defined by [param name] and [param theme_type], if it exists.
Returns the default color value if the property doesn't exist. Use [method has_color] to check for existence.
*/
//go:nosplit
func (self class) GetColor(name String.Name, theme_type String.Name) Color.RGBA { //gd:Theme.get_color
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the [Color] property defined by [param name] and [param theme_type] exists.
Returns [code]false[/code] if it doesn't exist. Use [method set_color] to define it.
*/
//go:nosplit
func (self class) HasColor(name String.Name, theme_type String.Name) bool { //gd:Theme.has_color
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_has_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Renames the [Color] property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_color] to check for existence, and [method clear_color] to remove the existing property.
*/
//go:nosplit
func (self class) RenameColor(old_name String.Name, name String.Name, theme_type String.Name) { //gd:Theme.rename_color
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(old_name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_rename_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes the [Color] property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_color] to check for existence.
*/
//go:nosplit
func (self class) ClearColor(name String.Name, theme_type String.Name) { //gd:Theme.clear_color
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_clear_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a list of names for [Color] properties defined with [param theme_type]. Use [method get_color_type_list] to get a list of possible theme type names.
*/
//go:nosplit
func (self class) GetColorList(theme_type String.Readable) Packed.Strings { //gd:Theme.get_color_list
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(theme_type)))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_color_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns a list of all unique theme type names for [Color] properties. Use [method get_type_list] to get a list of all unique theme types.
*/
//go:nosplit
func (self class) GetColorTypeList() Packed.Strings { //gd:Theme.get_color_type_list
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_color_type_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Creates or changes the value of the constant property defined by [param name] and [param theme_type]. Use [method clear_constant] to remove the property.
*/
//go:nosplit
func (self class) SetConstant(name String.Name, theme_type String.Name, constant int64) { //gd:Theme.set_constant
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	callframe.Arg(frame, constant)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_set_constant, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the constant property defined by [param name] and [param theme_type], if it exists.
Returns [code]0[/code] if the property doesn't exist. Use [method has_constant] to check for existence.
*/
//go:nosplit
func (self class) GetConstant(name String.Name, theme_type String.Name) int64 { //gd:Theme.get_constant
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_constant, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the constant property defined by [param name] and [param theme_type] exists.
Returns [code]false[/code] if it doesn't exist. Use [method set_constant] to define it.
*/
//go:nosplit
func (self class) HasConstant(name String.Name, theme_type String.Name) bool { //gd:Theme.has_constant
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_has_constant, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Renames the constant property defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_constant] to check for existence, and [method clear_constant] to remove the existing property.
*/
//go:nosplit
func (self class) RenameConstant(old_name String.Name, name String.Name, theme_type String.Name) { //gd:Theme.rename_constant
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(old_name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_rename_constant, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes the constant property defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_constant] to check for existence.
*/
//go:nosplit
func (self class) ClearConstant(name String.Name, theme_type String.Name) { //gd:Theme.clear_constant
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_clear_constant, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a list of names for constant properties defined with [param theme_type]. Use [method get_constant_type_list] to get a list of possible theme type names.
*/
//go:nosplit
func (self class) GetConstantList(theme_type String.Readable) Packed.Strings { //gd:Theme.get_constant_list
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(theme_type)))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_constant_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns a list of all unique theme type names for constant properties. Use [method get_type_list] to get a list of all unique theme types.
*/
//go:nosplit
func (self class) GetConstantTypeList() Packed.Strings { //gd:Theme.get_constant_type_list
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_constant_type_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDefaultBaseScale(base_scale float64) { //gd:Theme.set_default_base_scale
	var frame = callframe.New()
	callframe.Arg(frame, base_scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_set_default_base_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDefaultBaseScale() float64 { //gd:Theme.get_default_base_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_default_base_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if [member default_base_scale] has a valid value.
Returns [code]false[/code] if it doesn't. The value must be greater than [code]0.0[/code] to be considered valid.
*/
//go:nosplit
func (self class) HasDefaultBaseScale() bool { //gd:Theme.has_default_base_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_has_default_base_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDefaultFont(font [1]gdclass.Font) { //gd:Theme.set_default_font
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(font[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_set_default_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDefaultFont() [1]gdclass.Font { //gd:Theme.get_default_font
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_default_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Font{gd.PointerWithOwnershipTransferredToGo[gdclass.Font](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if [member default_font] has a valid value.
Returns [code]false[/code] if it doesn't.
*/
//go:nosplit
func (self class) HasDefaultFont() bool { //gd:Theme.has_default_font
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_has_default_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDefaultFontSize(font_size int64) { //gd:Theme.set_default_font_size
	var frame = callframe.New()
	callframe.Arg(frame, font_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_set_default_font_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDefaultFontSize() int64 { //gd:Theme.get_default_font_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_default_font_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if [member default_font_size] has a valid value.
Returns [code]false[/code] if it doesn't. The value must be greater than [code]0[/code] to be considered valid.
*/
//go:nosplit
func (self class) HasDefaultFontSize() bool { //gd:Theme.has_default_font_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_has_default_font_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates or changes the value of the theme property of [param data_type] defined by [param name] and [param theme_type]. Use [method clear_theme_item] to remove the property.
Fails if the [param value] type is not accepted by [param data_type].
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
//go:nosplit
func (self class) SetThemeItem(data_type gdclass.ThemeDataType, name String.Name, theme_type String.Name, value variant.Any) { //gd:Theme.set_theme_item
	var frame = callframe.New()
	callframe.Arg(frame, data_type)
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(value)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_set_theme_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the theme property of [param data_type] defined by [param name] and [param theme_type], if it exists.
Returns the engine fallback value if the property doesn't exist (see [ThemeDB]). Use [method has_theme_item] to check for existence.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
//go:nosplit
func (self class) GetThemeItem(data_type gdclass.ThemeDataType, name String.Name, theme_type String.Name) variant.Any { //gd:Theme.get_theme_item
	var frame = callframe.New()
	callframe.Arg(frame, data_type)
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_theme_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Through(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the theme property of [param data_type] defined by [param name] and [param theme_type] exists.
Returns [code]false[/code] if it doesn't exist. Use [method set_theme_item] to define it.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
//go:nosplit
func (self class) HasThemeItem(data_type gdclass.ThemeDataType, name String.Name, theme_type String.Name) bool { //gd:Theme.has_theme_item
	var frame = callframe.New()
	callframe.Arg(frame, data_type)
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_has_theme_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Renames the theme property of [param data_type] defined by [param old_name] and [param theme_type] to [param name], if it exists.
Fails if it doesn't exist, or if a similar property with the new name already exists. Use [method has_theme_item] to check for existence, and [method clear_theme_item] to remove the existing property.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
//go:nosplit
func (self class) RenameThemeItem(data_type gdclass.ThemeDataType, old_name String.Name, name String.Name, theme_type String.Name) { //gd:Theme.rename_theme_item
	var frame = callframe.New()
	callframe.Arg(frame, data_type)
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(old_name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_rename_theme_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes the theme property of [param data_type] defined by [param name] and [param theme_type], if it exists.
Fails if it doesn't exist. Use [method has_theme_item] to check for existence.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
//go:nosplit
func (self class) ClearThemeItem(data_type gdclass.ThemeDataType, name String.Name, theme_type String.Name) { //gd:Theme.clear_theme_item
	var frame = callframe.New()
	callframe.Arg(frame, data_type)
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_clear_theme_item, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a list of names for properties of [param data_type] defined with [param theme_type]. Use [method get_theme_item_type_list] to get a list of possible theme type names.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
//go:nosplit
func (self class) GetThemeItemList(data_type gdclass.ThemeDataType, theme_type String.Readable) Packed.Strings { //gd:Theme.get_theme_item_list
	var frame = callframe.New()
	callframe.Arg(frame, data_type)
	callframe.Arg(frame, pointers.Get(gd.InternalString(theme_type)))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_theme_item_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns a list of all unique theme type names for [param data_type] properties. Use [method get_type_list] to get a list of all unique theme types.
[b]Note:[/b] This method is analogous to calling the corresponding data type specific method, but can be used for more generalized logic.
*/
//go:nosplit
func (self class) GetThemeItemTypeList(data_type gdclass.ThemeDataType) Packed.Strings { //gd:Theme.get_theme_item_type_list
	var frame = callframe.New()
	callframe.Arg(frame, data_type)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_theme_item_type_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Marks [param theme_type] as a variation of [param base_type].
This adds [param theme_type] as a suggested option for [member Control.theme_type_variation] on a [Control] that is of the [param base_type] class.
Variations can also be nested, i.e. [param base_type] can be another variation. If a chain of variations ends with a [param base_type] matching the class of the [Control], the whole chain is going to be suggested as options.
[b]Note:[/b] Suggestions only show up if this theme resource is set as the project default theme. See [member ProjectSettings.gui/theme/custom].
*/
//go:nosplit
func (self class) SetTypeVariation(theme_type String.Name, base_type String.Name) { //gd:Theme.set_type_variation
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(base_type)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_set_type_variation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if [param theme_type] is marked as a variation of [param base_type].
*/
//go:nosplit
func (self class) IsTypeVariation(theme_type String.Name, base_type String.Name) bool { //gd:Theme.is_type_variation
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(base_type)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_is_type_variation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Unmarks [param theme_type] as being a variation of another theme type. See [method set_type_variation].
*/
//go:nosplit
func (self class) ClearTypeVariation(theme_type String.Name) { //gd:Theme.clear_type_variation
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_clear_type_variation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the name of the base theme type if [param theme_type] is a valid variation type. Returns an empty string otherwise.
*/
//go:nosplit
func (self class) GetTypeVariationBase(theme_type String.Name) String.Name { //gd:Theme.get_type_variation_base
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_type_variation_base, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns a list of all type variations for the given [param base_type].
*/
//go:nosplit
func (self class) GetTypeVariationList(base_type String.Name) Packed.Strings { //gd:Theme.get_type_variation_list
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(base_type)))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_type_variation_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Adds an empty theme type for every valid data type.
[b]Note:[/b] Empty types are not saved with the theme. This method only exists to perform in-memory changes to the resource. Use available [code]set_*[/code] methods to add theme items.
*/
//go:nosplit
func (self class) AddType(theme_type String.Name) { //gd:Theme.add_type
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_add_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes the theme type, gracefully discarding defined theme items. If the type is a variation, this information is also erased. If the type is a base for type variations, those variations lose their base.
*/
//go:nosplit
func (self class) RemoveType(theme_type String.Name) { //gd:Theme.remove_type
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(theme_type)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_remove_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a list of all unique theme type names. Use the appropriate [code]get_*_type_list[/code] method to get a list of unique theme types for a single data type.
*/
//go:nosplit
func (self class) GetTypeList() Packed.Strings { //gd:Theme.get_type_list
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_get_type_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Adds missing and overrides existing definitions with values from the [param other] theme resource.
[b]Note:[/b] This modifies the current theme. If you want to merge two themes together without modifying either one, create a new empty theme and merge the other two into it one after another.
*/
//go:nosplit
func (self class) MergeWith(other [1]gdclass.Theme) { //gd:Theme.merge_with
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(other[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_merge_with, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes all the theme properties defined on the theme resource.
*/
//go:nosplit
func (self class) Clear() { //gd:Theme.clear
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Theme.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsTheme() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTheme() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("Theme", func(ptr gd.Object) any { return [1]gdclass.Theme{*(*gdclass.Theme)(unsafe.Pointer(&ptr))} })
}

type DataType = gdclass.ThemeDataType //gd:Theme.DataType

const (
	/*Theme's [Color] item type.*/
	DataTypeColor DataType = 0
	/*Theme's constant item type.*/
	DataTypeConstant DataType = 1
	/*Theme's [Font] item type.*/
	DataTypeFont DataType = 2
	/*Theme's font size item type.*/
	DataTypeFontSize DataType = 3
	/*Theme's icon [Texture2D] item type.*/
	DataTypeIcon DataType = 4
	/*Theme's [StyleBox] item type.*/
	DataTypeStylebox DataType = 5
	/*Maximum value for the DataType enum.*/
	DataTypeMax DataType = 6
)
