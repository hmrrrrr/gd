package ProjectSettings

import "unsafe"
import "sync"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
Stores variables that can be accessed from everywhere. Use [method get_setting], [method set_setting] or [method has_setting] to access them. Variables stored in [code]project.godot[/code] are also loaded into [ProjectSettings], making this object very useful for reading custom game configuration options.
When naming a Project Settings property, use the full path to the setting including the category. For example, [code]"application/config/name"[/code] for the project name. Category and property names can be viewed in the Project Settings dialog.
[b]Feature tags:[/b] Project settings can be overridden for specific platforms and configurations (debug, release, ...) using [url=$DOCS_URL/tutorials/export/feature_tags.html]feature tags[/url].
[b]Overriding:[/b] Any project setting can be overridden by creating a file named [code]override.cfg[/code] in the project's root directory. This can also be used in exported projects by placing this file in the same directory as the project binary. Overriding will still take the base project settings' [url=$DOCS_URL/tutorials/export/feature_tags.html]feature tags[/url] in account. Therefore, make sure to [i]also[/i] override the setting with the desired feature tags if you want them to override base project settings on all platforms and configurations.
*/
var self gdclass.ProjectSettings
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.ProjectSettings)
	self = *(*gdclass.ProjectSettings)(unsafe.Pointer(&obj))
}

/*
Returns [code]true[/code] if a configuration value is present.
*/
func HasSetting(name string) bool {
	once.Do(singleton)
	return bool(class(self).HasSetting(gd.NewString(name)))
}

/*
Sets the value of a setting.
[b]Example:[/b]
[codeblocks]
[gdscript]
ProjectSettings.set_setting("application/config/name", "Example")
[/gdscript]
[csharp]
ProjectSettings.SetSetting("application/config/name", "Example");
[/csharp]
[/codeblocks]
This can also be used to erase custom project settings. To do this change the setting value to [code]null[/code].
*/
func SetSetting(name string, value gd.Variant) {
	once.Do(singleton)
	class(self).SetSetting(gd.NewString(name), value)
}

/*
Returns the value of the setting identified by [param name]. If the setting doesn't exist and [param default_value] is specified, the value of [param default_value] is returned. Otherwise, [code]null[/code] is returned.
[b]Example:[/b]
[codeblocks]
[gdscript]
print(ProjectSettings.get_setting("application/config/name"))
print(ProjectSettings.get_setting("application/config/custom_description", "No description specified."))
[/gdscript]
[csharp]
GD.Print(ProjectSettings.GetSetting("application/config/name"));
GD.Print(ProjectSettings.GetSetting("application/config/custom_description", "No description specified."));
[/csharp]
[/codeblocks]
[b]Note:[/b] This method doesn't take potential feature overrides into account automatically. Use [method get_setting_with_override] to handle seamlessly.
*/
func GetSetting(name string) gd.Variant {
	once.Do(singleton)
	return gd.Variant(class(self).GetSetting(gd.NewString(name), gd.NewVariant(([1]gd.Variant{}[0]))))
}

/*
Similar to [method get_setting], but applies feature tag overrides if any exists and is valid.
[b]Example:[/b]
If the following setting override exists "application/config/name.windows", and the following code is executed:
[codeblocks]
[gdscript]
print(ProjectSettings.get_setting_with_override("application/config/name"))
[/gdscript]
[csharp]
GD.Print(ProjectSettings.GetSettingWithOverride("application/config/name"));
[/csharp]
[/codeblocks]
Then the overridden setting will be returned instead if the project is running on the [i]Windows[/i] operating system.
*/
func GetSettingWithOverride(name string) gd.Variant {
	once.Do(singleton)
	return gd.Variant(class(self).GetSettingWithOverride(gd.NewStringName(name)))
}

/*
Returns an [Array] of registered global classes. Each global class is represented as a [Dictionary] that contains the following entries:
- [code]base[/code] is a name of the base class;
- [code]class[/code] is a name of the registered global class;
- [code]icon[/code] is a path to a custom icon of the global class, if it has any;
- [code]language[/code] is a name of a programming language in which the global class is written;
- [code]path[/code] is a path to a file containing the global class.
[b]Note:[/b] Both the script and the icon paths are local to the project filesystem, i.e. they start with [code]res://[/code].
*/
func GetGlobalClassList() gd.Array {
	once.Do(singleton)
	return gd.Array(class(self).GetGlobalClassList())
}

/*
Sets the order of a configuration value (influences when saved to the config file).
*/
func SetOrder(name string, position int) {
	once.Do(singleton)
	class(self).SetOrder(gd.NewString(name), gd.Int(position))
}

/*
Returns the order of a configuration value (influences when saved to the config file).
*/
func GetOrder(name string) int {
	once.Do(singleton)
	return int(int(class(self).GetOrder(gd.NewString(name))))
}

/*
Sets the specified setting's initial value. This is the value the setting reverts to.
*/
func SetInitialValue(name string, value gd.Variant) {
	once.Do(singleton)
	class(self).SetInitialValue(gd.NewString(name), value)
}

/*
Defines if the specified setting is considered basic or advanced. Basic settings will always be shown in the project settings. Advanced settings will only be shown if the user enables the "Advanced Settings" option.
*/
func SetAsBasic(name string, basic bool) {
	once.Do(singleton)
	class(self).SetAsBasic(gd.NewString(name), basic)
}

/*
Defines if the specified setting is considered internal. An internal setting won't show up in the Project Settings dialog. This is mostly useful for addons that need to store their own internal settings without exposing them directly to the user.
*/
func SetAsInternal(name string, internal_ bool) {
	once.Do(singleton)
	class(self).SetAsInternal(gd.NewString(name), internal_)
}

/*
Adds a custom property info to a property. The dictionary must contain:
- [code]"name"[/code]: [String] (the property's name)
- [code]"type"[/code]: [int] (see [enum Variant.Type])
- optionally [code]"hint"[/code]: [int] (see [enum PropertyHint]) and [code]"hint_string"[/code]: [String]
[b]Example:[/b]
[codeblocks]
[gdscript]
ProjectSettings.set("category/property_name", 0)

	var property_info = {
	    "name": "category/property_name",
	    "type": TYPE_INT,
	    "hint": PROPERTY_HINT_ENUM,
	    "hint_string": "one,two,three"
	}

ProjectSettings.add_property_info(property_info)
[/gdscript]
[csharp]
ProjectSettings.Singleton.Set("category/property_name", 0);

var propertyInfo = new Godot.Collections.Dictionary

	{
	    {"name", "category/propertyName"},
	    {"type", (int)Variant.Type.Int},
	    {"hint", (int)PropertyHint.Enum},
	    {"hint_string", "one,two,three"},
	};

ProjectSettings.AddPropertyInfo(propertyInfo);
[/csharp]
[/codeblocks]
*/
func AddPropertyInfo(hint gd.Dictionary) {
	once.Do(singleton)
	class(self).AddPropertyInfo(hint)
}

/*
Sets whether a setting requires restarting the editor to properly take effect.
[b]Note:[/b] This is just a hint to display to the user that the editor must be restarted for changes to take effect. Enabling [method set_restart_if_changed] does [i]not[/i] delay the setting being set when changed.
*/
func SetRestartIfChanged(name string, restart bool) {
	once.Do(singleton)
	class(self).SetRestartIfChanged(gd.NewString(name), restart)
}

/*
Clears the whole configuration (not recommended, may break things).
*/
func Clear(name string) {
	once.Do(singleton)
	class(self).Clear(gd.NewString(name))
}

/*
Returns the localized path (starting with [code]res://[/code]) corresponding to the absolute, native OS [param path]. See also [method globalize_path].
*/
func LocalizePath(path string) string {
	once.Do(singleton)
	return string(class(self).LocalizePath(gd.NewString(path)).String())
}

/*
Returns the absolute, native OS path corresponding to the localized [param path] (starting with [code]res://[/code] or [code]user://[/code]). The returned path will vary depending on the operating system and user preferences. See [url=$DOCS_URL/tutorials/io/data_paths.html]File paths in Godot projects[/url] to see what those paths convert to. See also [method localize_path].
[b]Note:[/b] [method globalize_path] with [code]res://[/code] will not work in an exported project. Instead, prepend the executable's base directory to the path when running from an exported project:
[codeblock]
var path = ""
if OS.has_feature("editor"):

	# Running from an editor binary.
	# `path` will contain the absolute path to `hello.txt` located in the project root.
	path = ProjectSettings.globalize_path("res://hello.txt")

else:

	# Running from an exported project.
	# `path` will contain the absolute path to `hello.txt` next to the executable.
	# This is *not* identical to using `ProjectSettings.globalize_path()` with a `res://` path,
	# but is close enough in spirit.
	path = OS.get_executable_path().get_base_dir().path_join("hello.txt")

[/codeblock]
*/
func GlobalizePath(path string) string {
	once.Do(singleton)
	return string(class(self).GlobalizePath(gd.NewString(path)).String())
}

/*
Saves the configuration to the [code]project.godot[/code] file.
[b]Note:[/b] This method is intended to be used by editor plugins, as modified [ProjectSettings] can't be loaded back in the running app. If you want to change project settings in exported projects, use [method save_custom] to save [code]override.cfg[/code] file.
*/
func Save() gd.Error {
	once.Do(singleton)
	return gd.Error(class(self).Save())
}

/*
Loads the contents of the .pck or .zip file specified by [param pack] into the resource filesystem ([code]res://[/code]). Returns [code]true[/code] on success.
[b]Note:[/b] If a file from [param pack] shares the same path as a file already in the resource filesystem, any attempts to load that file will use the file from [param pack] unless [param replace_files] is set to [code]false[/code].
[b]Note:[/b] The optional [param offset] parameter can be used to specify the offset in bytes to the start of the resource pack. This is only supported for .pck files.
*/
func LoadResourcePack(pack string) bool {
	once.Do(singleton)
	return bool(class(self).LoadResourcePack(gd.NewString(pack), true, gd.Int(0)))
}

/*
Saves the configuration to a custom file. The file extension must be [code].godot[/code] (to save in text-based [ConfigFile] format) or [code].binary[/code] (to save in binary format). You can also save [code]override.cfg[/code] file, which is also text, but can be used in exported projects unlike other formats.
*/
func SaveCustom(file string) gd.Error {
	once.Do(singleton)
	return gd.Error(class(self).SaveCustom(gd.NewString(file)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]classdb.ProjectSettings

func (self class) AsObject() gd.Object { return self[0].AsObject() }

/*
Returns [code]true[/code] if a configuration value is present.
*/
//go:nosplit
func (self class) HasSetting(name gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_has_setting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the value of a setting.
[b]Example:[/b]
[codeblocks]
[gdscript]
ProjectSettings.set_setting("application/config/name", "Example")
[/gdscript]
[csharp]
ProjectSettings.SetSetting("application/config/name", "Example");
[/csharp]
[/codeblocks]
This can also be used to erase custom project settings. To do this change the setting value to [code]null[/code].
*/
//go:nosplit
func (self class) SetSetting(name gd.String, value gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(value))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_set_setting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the value of the setting identified by [param name]. If the setting doesn't exist and [param default_value] is specified, the value of [param default_value] is returned. Otherwise, [code]null[/code] is returned.
[b]Example:[/b]
[codeblocks]
[gdscript]
print(ProjectSettings.get_setting("application/config/name"))
print(ProjectSettings.get_setting("application/config/custom_description", "No description specified."))
[/gdscript]
[csharp]
GD.Print(ProjectSettings.GetSetting("application/config/name"));
GD.Print(ProjectSettings.GetSetting("application/config/custom_description", "No description specified."));
[/csharp]
[/codeblocks]
[b]Note:[/b] This method doesn't take potential feature overrides into account automatically. Use [method get_setting_with_override] to handle seamlessly.
*/
//go:nosplit
func (self class) GetSetting(name gd.String, default_value gd.Variant) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(default_value))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_get_setting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Similar to [method get_setting], but applies feature tag overrides if any exists and is valid.
[b]Example:[/b]
If the following setting override exists "application/config/name.windows", and the following code is executed:
[codeblocks]
[gdscript]
print(ProjectSettings.get_setting_with_override("application/config/name"))
[/gdscript]
[csharp]
GD.Print(ProjectSettings.GetSettingWithOverride("application/config/name"));
[/csharp]
[/codeblocks]
Then the overridden setting will be returned instead if the project is running on the [i]Windows[/i] operating system.
*/
//go:nosplit
func (self class) GetSettingWithOverride(name gd.StringName) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_get_setting_with_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns an [Array] of registered global classes. Each global class is represented as a [Dictionary] that contains the following entries:
- [code]base[/code] is a name of the base class;
- [code]class[/code] is a name of the registered global class;
- [code]icon[/code] is a path to a custom icon of the global class, if it has any;
- [code]language[/code] is a name of a programming language in which the global class is written;
- [code]path[/code] is a path to a file containing the global class.
[b]Note:[/b] Both the script and the icon paths are local to the project filesystem, i.e. they start with [code]res://[/code].
*/
//go:nosplit
func (self class) GetGlobalClassList() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_get_global_class_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the order of a configuration value (influences when saved to the config file).
*/
//go:nosplit
func (self class) SetOrder(name gd.String, position gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_set_order, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the order of a configuration value (influences when saved to the config file).
*/
//go:nosplit
func (self class) GetOrder(name gd.String) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_get_order, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the specified setting's initial value. This is the value the setting reverts to.
*/
//go:nosplit
func (self class) SetInitialValue(name gd.String, value gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(value))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_set_initial_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Defines if the specified setting is considered basic or advanced. Basic settings will always be shown in the project settings. Advanced settings will only be shown if the user enables the "Advanced Settings" option.
*/
//go:nosplit
func (self class) SetAsBasic(name gd.String, basic bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, basic)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_set_as_basic, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Defines if the specified setting is considered internal. An internal setting won't show up in the Project Settings dialog. This is mostly useful for addons that need to store their own internal settings without exposing them directly to the user.
*/
//go:nosplit
func (self class) SetAsInternal(name gd.String, internal_ bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, internal_)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_set_as_internal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a custom property info to a property. The dictionary must contain:
- [code]"name"[/code]: [String] (the property's name)
- [code]"type"[/code]: [int] (see [enum Variant.Type])
- optionally [code]"hint"[/code]: [int] (see [enum PropertyHint]) and [code]"hint_string"[/code]: [String]
[b]Example:[/b]
[codeblocks]
[gdscript]
ProjectSettings.set("category/property_name", 0)

var property_info = {
    "name": "category/property_name",
    "type": TYPE_INT,
    "hint": PROPERTY_HINT_ENUM,
    "hint_string": "one,two,three"
}

ProjectSettings.add_property_info(property_info)
[/gdscript]
[csharp]
ProjectSettings.Singleton.Set("category/property_name", 0);

var propertyInfo = new Godot.Collections.Dictionary
{
    {"name", "category/propertyName"},
    {"type", (int)Variant.Type.Int},
    {"hint", (int)PropertyHint.Enum},
    {"hint_string", "one,two,three"},
};

ProjectSettings.AddPropertyInfo(propertyInfo);
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) AddPropertyInfo(hint gd.Dictionary) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(hint))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_add_property_info, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets whether a setting requires restarting the editor to properly take effect.
[b]Note:[/b] This is just a hint to display to the user that the editor must be restarted for changes to take effect. Enabling [method set_restart_if_changed] does [i]not[/i] delay the setting being set when changed.
*/
//go:nosplit
func (self class) SetRestartIfChanged(name gd.String, restart bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, restart)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_set_restart_if_changed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Clears the whole configuration (not recommended, may break things).
*/
//go:nosplit
func (self class) Clear(name gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the localized path (starting with [code]res://[/code]) corresponding to the absolute, native OS [param path]. See also [method globalize_path].
*/
//go:nosplit
func (self class) LocalizePath(path gd.String) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_localize_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the absolute, native OS path corresponding to the localized [param path] (starting with [code]res://[/code] or [code]user://[/code]). The returned path will vary depending on the operating system and user preferences. See [url=$DOCS_URL/tutorials/io/data_paths.html]File paths in Godot projects[/url] to see what those paths convert to. See also [method localize_path].
[b]Note:[/b] [method globalize_path] with [code]res://[/code] will not work in an exported project. Instead, prepend the executable's base directory to the path when running from an exported project:
[codeblock]
var path = ""
if OS.has_feature("editor"):
    # Running from an editor binary.
    # `path` will contain the absolute path to `hello.txt` located in the project root.
    path = ProjectSettings.globalize_path("res://hello.txt")
else:
    # Running from an exported project.
    # `path` will contain the absolute path to `hello.txt` next to the executable.
    # This is *not* identical to using `ProjectSettings.globalize_path()` with a `res://` path,
    # but is close enough in spirit.
    path = OS.get_executable_path().get_base_dir().path_join("hello.txt")
[/codeblock]
*/
//go:nosplit
func (self class) GlobalizePath(path gd.String) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_globalize_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Saves the configuration to the [code]project.godot[/code] file.
[b]Note:[/b] This method is intended to be used by editor plugins, as modified [ProjectSettings] can't be loaded back in the running app. If you want to change project settings in exported projects, use [method save_custom] to save [code]override.cfg[/code] file.
*/
//go:nosplit
func (self class) Save() int64 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_save, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Loads the contents of the .pck or .zip file specified by [param pack] into the resource filesystem ([code]res://[/code]). Returns [code]true[/code] on success.
[b]Note:[/b] If a file from [param pack] shares the same path as a file already in the resource filesystem, any attempts to load that file will use the file from [param pack] unless [param replace_files] is set to [code]false[/code].
[b]Note:[/b] The optional [param offset] parameter can be used to specify the offset in bytes to the start of the resource pack. This is only supported for .pck files.
*/
//go:nosplit
func (self class) LoadResourcePack(pack gd.String, replace_files bool, offset gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(pack))
	callframe.Arg(frame, replace_files)
	callframe.Arg(frame, offset)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_load_resource_pack, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Saves the configuration to a custom file. The file extension must be [code].godot[/code] (to save in text-based [ConfigFile] format) or [code].binary[/code] (to save in binary format). You can also save [code]override.cfg[/code] file, which is also text, but can be used in exported projects unlike other formats.
*/
//go:nosplit
func (self class) SaveCustom(file gd.String) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(file))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProjectSettings.Bind_save_custom, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func OnSettingsChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("settings_changed"), gd.NewCallable(cb), 0)
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {
	classdb.Register("ProjectSettings", func(ptr gd.Object) any { return classdb.ProjectSettings(ptr) })
}
