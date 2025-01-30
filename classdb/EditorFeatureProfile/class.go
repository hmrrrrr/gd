// Package EditorFeatureProfile provides methods for working with EditorFeatureProfile object instances.
package EditorFeatureProfile

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
An editor feature profile can be used to disable specific features of the Godot editor. When disabled, the features won't appear in the editor, which makes the editor less cluttered. This is useful in education settings to reduce confusion or when working in a team. For example, artists and level designers could use a feature profile that disables the script editor to avoid accidentally making changes to files they aren't supposed to edit.
To manage editor feature profiles visually, use [b]Editor > Manage Feature Profiles...[/b] at the top of the editor window.
*/
type Instance [1]gdclass.EditorFeatureProfile

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorFeatureProfile() Instance
}

/*
If [param disable] is [code]true[/code], disables the class specified by [param class_name]. When disabled, the class won't appear in the Create New Node dialog.
*/
func (self Instance) SetDisableClass(class_name string, disable bool) { //gd:EditorFeatureProfile.set_disable_class
	class(self).SetDisableClass(String.Name(String.New(class_name)), disable)
}

/*
Returns [code]true[/code] if the class specified by [param class_name] is disabled. When disabled, the class won't appear in the Create New Node dialog.
*/
func (self Instance) IsClassDisabled(class_name string) bool { //gd:EditorFeatureProfile.is_class_disabled
	return bool(class(self).IsClassDisabled(String.Name(String.New(class_name))))
}

/*
If [param disable] is [code]true[/code], disables editing for the class specified by [param class_name]. When disabled, the class will still appear in the Create New Node dialog but the Inspector will be read-only when selecting a node that extends the class.
*/
func (self Instance) SetDisableClassEditor(class_name string, disable bool) { //gd:EditorFeatureProfile.set_disable_class_editor
	class(self).SetDisableClassEditor(String.Name(String.New(class_name)), disable)
}

/*
Returns [code]true[/code] if editing for the class specified by [param class_name] is disabled. When disabled, the class will still appear in the Create New Node dialog but the Inspector will be read-only when selecting a node that extends the class.
*/
func (self Instance) IsClassEditorDisabled(class_name string) bool { //gd:EditorFeatureProfile.is_class_editor_disabled
	return bool(class(self).IsClassEditorDisabled(String.Name(String.New(class_name))))
}

/*
If [param disable] is [code]true[/code], disables editing for [param property] in the class specified by [param class_name]. When a property is disabled, it won't appear in the Inspector when selecting a node that extends the class specified by [param class_name].
*/
func (self Instance) SetDisableClassProperty(class_name string, property string, disable bool) { //gd:EditorFeatureProfile.set_disable_class_property
	class(self).SetDisableClassProperty(String.Name(String.New(class_name)), String.Name(String.New(property)), disable)
}

/*
Returns [code]true[/code] if [param property] is disabled in the class specified by [param class_name]. When a property is disabled, it won't appear in the Inspector when selecting a node that extends the class specified by [param class_name].
*/
func (self Instance) IsClassPropertyDisabled(class_name string, property string) bool { //gd:EditorFeatureProfile.is_class_property_disabled
	return bool(class(self).IsClassPropertyDisabled(String.Name(String.New(class_name)), String.Name(String.New(property))))
}

/*
If [param disable] is [code]true[/code], disables the editor feature specified in [param feature]. When a feature is disabled, it will disappear from the editor entirely.
*/
func (self Instance) SetDisableFeature(feature gdclass.EditorFeatureProfileFeature, disable bool) { //gd:EditorFeatureProfile.set_disable_feature
	class(self).SetDisableFeature(feature, disable)
}

/*
Returns [code]true[/code] if the [param feature] is disabled. When a feature is disabled, it will disappear from the editor entirely.
*/
func (self Instance) IsFeatureDisabled(feature gdclass.EditorFeatureProfileFeature) bool { //gd:EditorFeatureProfile.is_feature_disabled
	return bool(class(self).IsFeatureDisabled(feature))
}

/*
Returns the specified [param feature]'s human-readable name.
*/
func (self Instance) GetFeatureName(feature gdclass.EditorFeatureProfileFeature) string { //gd:EditorFeatureProfile.get_feature_name
	return string(class(self).GetFeatureName(feature).String())
}

/*
Saves the editor feature profile to a file in JSON format. It can then be imported using the feature profile manager's [b]Import[/b] button or the [method load_from_file] method.
[b]Note:[/b] Feature profiles created via the user interface are saved in the [code]feature_profiles[/code] directory, as a file with the [code].profile[/code] extension. The editor configuration folder can be found by using [method EditorPaths.get_config_dir].
*/
func (self Instance) SaveToFile(path string) error { //gd:EditorFeatureProfile.save_to_file
	return error(gd.ToError(class(self).SaveToFile(String.New(path))))
}

/*
Loads an editor feature profile from a file. The file must follow the JSON format obtained by using the feature profile manager's [b]Export[/b] button or the [method save_to_file] method.
[b]Note:[/b] Feature profiles created via the user interface are loaded from the [code]feature_profiles[/code] directory, as a file with the [code].profile[/code] extension. The editor configuration folder can be found by using [method EditorPaths.get_config_dir].
*/
func (self Instance) LoadFromFile(path string) error { //gd:EditorFeatureProfile.load_from_file
	return error(gd.ToError(class(self).LoadFromFile(String.New(path))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorFeatureProfile

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorFeatureProfile"))
	casted := Instance{*(*gdclass.EditorFeatureProfile)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
If [param disable] is [code]true[/code], disables the class specified by [param class_name]. When disabled, the class won't appear in the Create New Node dialog.
*/
//go:nosplit
func (self class) SetDisableClass(class_name String.Name, disable bool) { //gd:EditorFeatureProfile.set_disable_class
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_name)))
	callframe.Arg(frame, disable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFeatureProfile.Bind_set_disable_class, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the class specified by [param class_name] is disabled. When disabled, the class won't appear in the Create New Node dialog.
*/
//go:nosplit
func (self class) IsClassDisabled(class_name String.Name) bool { //gd:EditorFeatureProfile.is_class_disabled
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_name)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFeatureProfile.Bind_is_class_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param disable] is [code]true[/code], disables editing for the class specified by [param class_name]. When disabled, the class will still appear in the Create New Node dialog but the Inspector will be read-only when selecting a node that extends the class.
*/
//go:nosplit
func (self class) SetDisableClassEditor(class_name String.Name, disable bool) { //gd:EditorFeatureProfile.set_disable_class_editor
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_name)))
	callframe.Arg(frame, disable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFeatureProfile.Bind_set_disable_class_editor, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if editing for the class specified by [param class_name] is disabled. When disabled, the class will still appear in the Create New Node dialog but the Inspector will be read-only when selecting a node that extends the class.
*/
//go:nosplit
func (self class) IsClassEditorDisabled(class_name String.Name) bool { //gd:EditorFeatureProfile.is_class_editor_disabled
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_name)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFeatureProfile.Bind_is_class_editor_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param disable] is [code]true[/code], disables editing for [param property] in the class specified by [param class_name]. When a property is disabled, it won't appear in the Inspector when selecting a node that extends the class specified by [param class_name].
*/
//go:nosplit
func (self class) SetDisableClassProperty(class_name String.Name, property String.Name, disable bool) { //gd:EditorFeatureProfile.set_disable_class_property
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(property)))
	callframe.Arg(frame, disable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFeatureProfile.Bind_set_disable_class_property, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if [param property] is disabled in the class specified by [param class_name]. When a property is disabled, it won't appear in the Inspector when selecting a node that extends the class specified by [param class_name].
*/
//go:nosplit
func (self class) IsClassPropertyDisabled(class_name String.Name, property String.Name) bool { //gd:EditorFeatureProfile.is_class_property_disabled
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(property)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFeatureProfile.Bind_is_class_property_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param disable] is [code]true[/code], disables the editor feature specified in [param feature]. When a feature is disabled, it will disappear from the editor entirely.
*/
//go:nosplit
func (self class) SetDisableFeature(feature gdclass.EditorFeatureProfileFeature, disable bool) { //gd:EditorFeatureProfile.set_disable_feature
	var frame = callframe.New()
	callframe.Arg(frame, feature)
	callframe.Arg(frame, disable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFeatureProfile.Bind_set_disable_feature, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the [param feature] is disabled. When a feature is disabled, it will disappear from the editor entirely.
*/
//go:nosplit
func (self class) IsFeatureDisabled(feature gdclass.EditorFeatureProfileFeature) bool { //gd:EditorFeatureProfile.is_feature_disabled
	var frame = callframe.New()
	callframe.Arg(frame, feature)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFeatureProfile.Bind_is_feature_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the specified [param feature]'s human-readable name.
*/
//go:nosplit
func (self class) GetFeatureName(feature gdclass.EditorFeatureProfileFeature) String.Readable { //gd:EditorFeatureProfile.get_feature_name
	var frame = callframe.New()
	callframe.Arg(frame, feature)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFeatureProfile.Bind_get_feature_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Saves the editor feature profile to a file in JSON format. It can then be imported using the feature profile manager's [b]Import[/b] button or the [method load_from_file] method.
[b]Note:[/b] Feature profiles created via the user interface are saved in the [code]feature_profiles[/code] directory, as a file with the [code].profile[/code] extension. The editor configuration folder can be found by using [method EditorPaths.get_config_dir].
*/
//go:nosplit
func (self class) SaveToFile(path String.Readable) Error.Code { //gd:EditorFeatureProfile.save_to_file
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFeatureProfile.Bind_save_to_file, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Loads an editor feature profile from a file. The file must follow the JSON format obtained by using the feature profile manager's [b]Export[/b] button or the [method save_to_file] method.
[b]Note:[/b] Feature profiles created via the user interface are loaded from the [code]feature_profiles[/code] directory, as a file with the [code].profile[/code] extension. The editor configuration folder can be found by using [method EditorPaths.get_config_dir].
*/
//go:nosplit
func (self class) LoadFromFile(path String.Readable) Error.Code { //gd:EditorFeatureProfile.load_from_file
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFeatureProfile.Bind_load_from_file, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsEditorFeatureProfile() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorFeatureProfile() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("EditorFeatureProfile", func(ptr gd.Object) any {
		return [1]gdclass.EditorFeatureProfile{*(*gdclass.EditorFeatureProfile)(unsafe.Pointer(&ptr))}
	})
}

type Feature = gdclass.EditorFeatureProfileFeature //gd:EditorFeatureProfile.Feature

const (
	/*The 3D editor. If this feature is disabled, the 3D editor won't display but 3D nodes will still display in the Create New Node dialog.*/
	Feature3d Feature = 0
	/*The Script tab, which contains the script editor and class reference browser. If this feature is disabled, the Script tab won't display.*/
	FeatureScript Feature = 1
	/*The AssetLib tab. If this feature is disabled, the AssetLib tab won't display.*/
	FeatureAssetLib Feature = 2
	/*Scene tree editing. If this feature is disabled, the Scene tree dock will still be visible but will be read-only.*/
	FeatureSceneTree Feature = 3
	/*The Node dock. If this feature is disabled, signals and groups won't be visible and modifiable from the editor.*/
	FeatureNodeDock Feature = 4
	/*The FileSystem dock. If this feature is disabled, the FileSystem dock won't be visible.*/
	FeatureFilesystemDock Feature = 5
	/*The Import dock. If this feature is disabled, the Import dock won't be visible.*/
	FeatureImportDock Feature = 6
	/*The History dock. If this feature is disabled, the History dock won't be visible.*/
	FeatureHistoryDock Feature = 7
	/*Represents the size of the [enum Feature] enum.*/
	FeatureMax Feature = 8
)
