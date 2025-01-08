// Package EditorPaths provides methods for working with EditorPaths object instances.
package EditorPaths

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This editor-only singleton returns OS-specific paths to various data folders and files. It can be used in editor plugins to ensure files are saved in the correct location on each operating system.
[b]Note:[/b] This singleton is not accessible in exported projects. Attempting to access it in an exported project will result in a script error as the singleton won't be declared. To prevent script errors in exported projects, use [method Engine.has_singleton] to check whether the singleton is available before using it.
[b]Note:[/b] On the Linux/BSD platform, Godot complies with the [url=https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html]XDG Base Directory Specification[/url]. You can override environment variables following the specification to change the editor and project data paths.
*/
type Instance [1]gdclass.EditorPaths

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorPaths() Instance
}

/*
Returns the absolute path to the user's data folder. This folder should be used for [i]persistent[/i] user data files such as installed export templates.
[b]Default paths per platform:[/b]
[codeblock lang=text]
- Windows: %APPDATA%\Godot\                    (same as `get_config_dir()`)
- macOS: ~/Library/Application Support/Godot/  (same as `get_config_dir()`)
- Linux: ~/.local/share/godot/
[/codeblock]
*/
func (self Instance) GetDataDir() string {
	return string(class(self).GetDataDir().String())
}

/*
Returns the absolute path to the user's configuration folder. This folder should be used for [i]persistent[/i] user configuration files.
[b]Default paths per platform:[/b]
[codeblock lang=text]
- Windows: %APPDATA%\Godot\                    (same as `get_data_dir()`)
- macOS: ~/Library/Application Support/Godot/  (same as `get_data_dir()`)
- Linux: ~/.config/godot/
[/codeblock]
*/
func (self Instance) GetConfigDir() string {
	return string(class(self).GetConfigDir().String())
}

/*
Returns the absolute path to the user's cache folder. This folder should be used for temporary data that can be removed safely whenever the editor is closed (such as generated resource thumbnails).
[b]Default paths per platform:[/b]
[codeblock lang=text]
- Windows: %LOCALAPPDATA%\Godot\
- macOS: ~/Library/Caches/Godot/
- Linux: ~/.cache/godot/
[/codeblock]
*/
func (self Instance) GetCacheDir() string {
	return string(class(self).GetCacheDir().String())
}

/*
Returns [code]true[/code] if the editor is marked as self-contained, [code]false[/code] otherwise. When self-contained mode is enabled, user configuration, data and cache files are saved in an [code]editor_data/[/code] folder next to the editor binary. This makes portable usage easier and ensures the Godot editor minimizes file writes outside its own folder. Self-contained mode is not available for exported projects.
Self-contained mode can be enabled by creating a file named [code]._sc_[/code] or [code]_sc_[/code] in the same folder as the editor binary or macOS .app bundle while the editor is not running. See also [method get_self_contained_file].
[b]Note:[/b] On macOS, quarantine flag should be manually removed before using self-contained mode, see [url=https://docs.godotengine.org/en/stable/tutorials/export/running_on_macos.html]Running on macOS[/url].
[b]Note:[/b] On macOS, placing [code]_sc_[/code] or any other file inside .app bundle will break digital signature and make it non-portable, consider placing it in the same folder as the .app bundle instead.
[b]Note:[/b] The Steam release of Godot uses self-contained mode by default.
*/
func (self Instance) IsSelfContained() bool {
	return bool(class(self).IsSelfContained())
}

/*
Returns the absolute path to the self-contained file that makes the current Godot editor instance be considered as self-contained. Returns an empty string if the current Godot editor instance isn't self-contained. See also [method is_self_contained].
*/
func (self Instance) GetSelfContainedFile() string {
	return string(class(self).GetSelfContainedFile().String())
}

/*
Returns the project-specific editor settings path. Projects all have a unique subdirectory inside the settings path where project-specific editor settings are saved.
*/
func (self Instance) GetProjectSettingsDir() string {
	return string(class(self).GetProjectSettingsDir().String())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorPaths

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorPaths"))
	casted := Instance{*(*gdclass.EditorPaths)(unsafe.Pointer(&object))}
	return casted
}

/*
Returns the absolute path to the user's data folder. This folder should be used for [i]persistent[/i] user data files such as installed export templates.
[b]Default paths per platform:[/b]
[codeblock lang=text]
- Windows: %APPDATA%\Godot\                    (same as `get_config_dir()`)
- macOS: ~/Library/Application Support/Godot/  (same as `get_config_dir()`)
- Linux: ~/.local/share/godot/
[/codeblock]
*/
//go:nosplit
func (self class) GetDataDir() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPaths.Bind_get_data_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the absolute path to the user's configuration folder. This folder should be used for [i]persistent[/i] user configuration files.
[b]Default paths per platform:[/b]
[codeblock lang=text]
- Windows: %APPDATA%\Godot\                    (same as `get_data_dir()`)
- macOS: ~/Library/Application Support/Godot/  (same as `get_data_dir()`)
- Linux: ~/.config/godot/
[/codeblock]
*/
//go:nosplit
func (self class) GetConfigDir() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPaths.Bind_get_config_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the absolute path to the user's cache folder. This folder should be used for temporary data that can be removed safely whenever the editor is closed (such as generated resource thumbnails).
[b]Default paths per platform:[/b]
[codeblock lang=text]
- Windows: %LOCALAPPDATA%\Godot\
- macOS: ~/Library/Caches/Godot/
- Linux: ~/.cache/godot/
[/codeblock]
*/
//go:nosplit
func (self class) GetCacheDir() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPaths.Bind_get_cache_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the editor is marked as self-contained, [code]false[/code] otherwise. When self-contained mode is enabled, user configuration, data and cache files are saved in an [code]editor_data/[/code] folder next to the editor binary. This makes portable usage easier and ensures the Godot editor minimizes file writes outside its own folder. Self-contained mode is not available for exported projects.
Self-contained mode can be enabled by creating a file named [code]._sc_[/code] or [code]_sc_[/code] in the same folder as the editor binary or macOS .app bundle while the editor is not running. See also [method get_self_contained_file].
[b]Note:[/b] On macOS, quarantine flag should be manually removed before using self-contained mode, see [url=https://docs.godotengine.org/en/stable/tutorials/export/running_on_macos.html]Running on macOS[/url].
[b]Note:[/b] On macOS, placing [code]_sc_[/code] or any other file inside .app bundle will break digital signature and make it non-portable, consider placing it in the same folder as the .app bundle instead.
[b]Note:[/b] The Steam release of Godot uses self-contained mode by default.
*/
//go:nosplit
func (self class) IsSelfContained() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPaths.Bind_is_self_contained, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the absolute path to the self-contained file that makes the current Godot editor instance be considered as self-contained. Returns an empty string if the current Godot editor instance isn't self-contained. See also [method is_self_contained].
*/
//go:nosplit
func (self class) GetSelfContainedFile() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPaths.Bind_get_self_contained_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the project-specific editor settings path. Projects all have a unique subdirectory inside the settings path where project-specific editor settings are saved.
*/
//go:nosplit
func (self class) GetProjectSettingsDir() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPaths.Bind_get_project_settings_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsEditorPaths() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorPaths() Instance { return *((*Instance)(unsafe.Pointer(&self))) }

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
	gdclass.Register("EditorPaths", func(ptr gd.Object) any { return [1]gdclass.EditorPaths{*(*gdclass.EditorPaths)(unsafe.Pointer(&ptr))} })
}
