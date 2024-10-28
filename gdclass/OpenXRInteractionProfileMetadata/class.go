package OpenXRInteractionProfileMetadata

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
This class allows OpenXR core and extensions to register metadata relating to supported interaction devices such as controllers, trackers, haptic devices, etc. It is primarily used by the action map editor and to sanitize any action map by removing extension-dependent entries when applicable.

*/
type Go [1]classdb.OpenXRInteractionProfileMetadata

/*
Allows for renaming old interaction profile paths to new paths to maintain backwards compatibility with older action maps.
*/
func (self Go) RegisterProfileRename(old_name string, new_name string) {
	class(self).RegisterProfileRename(gd.NewString(old_name), gd.NewString(new_name))
}

/*
Registers a top level path to which profiles can be bound. For instance [code]/user/hand/left[/code] refers to the bind point for the player's left hand. Extensions can register additional top level paths, for instance a haptic vest extension might register [code]/user/body/vest[/code].
[param display_name] is the name shown to the user. [param openxr_path] is the top level path being registered. [param openxr_extension_name] is optional and ensures the top level path is only used if the specified extension is available/enabled.
When a top level path ends up being bound by OpenXR, a [XRPositionalTracker] is instantiated to manage the state of the device.
*/
func (self Go) RegisterTopLevelPath(display_name string, openxr_path string, openxr_extension_name string) {
	class(self).RegisterTopLevelPath(gd.NewString(display_name), gd.NewString(openxr_path), gd.NewString(openxr_extension_name))
}

/*
Registers an interaction profile using its OpenXR designation (e.g. [code]/interaction_profiles/khr/simple_controller[/code] is the profile for OpenXR's simple controller profile).
[param display_name] is the description shown to the user. [param openxr_path] is the interaction profile path being registered. [param openxr_extension_name] optionally restricts this profile to the given extension being enabled/available. If the extension is not available, the profile and all related entries used in an action map are filtered out.
*/
func (self Go) RegisterInteractionProfile(display_name string, openxr_path string, openxr_extension_name string) {
	class(self).RegisterInteractionProfile(gd.NewString(display_name), gd.NewString(openxr_path), gd.NewString(openxr_extension_name))
}

/*
Registers an input/output path for the given [param interaction_profile]. The profile should previously have been registered using [method register_interaction_profile]. [param display_name] is the description shown to the user. [param toplevel_path] specifies the bind path this input/output can be bound to (e.g. [code]/user/hand/left[/code] or [code]/user/hand/right[/code]). [param openxr_path] is the action input/output being registered (e.g. [code]/user/hand/left/input/aim/pose[/code]). [param openxr_extension_name] restricts this input/output to an enabled/available extension, this doesn't need to repeat the extension on the profile but relates to overlapping extension (e.g. [code]XR_EXT_palm_pose[/code] that introduces [code]…/input/palm_ext/pose[/code] input paths). [param action_type] defines the type of input or output provided by OpenXR.
*/
func (self Go) RegisterIoPath(interaction_profile string, display_name string, toplevel_path string, openxr_path string, openxr_extension_name string, action_type classdb.OpenXRActionActionType) {
	class(self).RegisterIoPath(gd.NewString(interaction_profile), gd.NewString(display_name), gd.NewString(toplevel_path), gd.NewString(openxr_path), gd.NewString(openxr_extension_name), action_type)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.OpenXRInteractionProfileMetadata
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OpenXRInteractionProfileMetadata"))
	return Go{classdb.OpenXRInteractionProfileMetadata(object)}
}

/*
Allows for renaming old interaction profile paths to new paths to maintain backwards compatibility with older action maps.
*/
//go:nosplit
func (self class) RegisterProfileRename(old_name gd.String, new_name gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(old_name))
	callframe.Arg(frame, discreet.Get(new_name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInteractionProfileMetadata.Bind_register_profile_rename, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Registers a top level path to which profiles can be bound. For instance [code]/user/hand/left[/code] refers to the bind point for the player's left hand. Extensions can register additional top level paths, for instance a haptic vest extension might register [code]/user/body/vest[/code].
[param display_name] is the name shown to the user. [param openxr_path] is the top level path being registered. [param openxr_extension_name] is optional and ensures the top level path is only used if the specified extension is available/enabled.
When a top level path ends up being bound by OpenXR, a [XRPositionalTracker] is instantiated to manage the state of the device.
*/
//go:nosplit
func (self class) RegisterTopLevelPath(display_name gd.String, openxr_path gd.String, openxr_extension_name gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(display_name))
	callframe.Arg(frame, discreet.Get(openxr_path))
	callframe.Arg(frame, discreet.Get(openxr_extension_name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInteractionProfileMetadata.Bind_register_top_level_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Registers an interaction profile using its OpenXR designation (e.g. [code]/interaction_profiles/khr/simple_controller[/code] is the profile for OpenXR's simple controller profile).
[param display_name] is the description shown to the user. [param openxr_path] is the interaction profile path being registered. [param openxr_extension_name] optionally restricts this profile to the given extension being enabled/available. If the extension is not available, the profile and all related entries used in an action map are filtered out.
*/
//go:nosplit
func (self class) RegisterInteractionProfile(display_name gd.String, openxr_path gd.String, openxr_extension_name gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(display_name))
	callframe.Arg(frame, discreet.Get(openxr_path))
	callframe.Arg(frame, discreet.Get(openxr_extension_name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInteractionProfileMetadata.Bind_register_interaction_profile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Registers an input/output path for the given [param interaction_profile]. The profile should previously have been registered using [method register_interaction_profile]. [param display_name] is the description shown to the user. [param toplevel_path] specifies the bind path this input/output can be bound to (e.g. [code]/user/hand/left[/code] or [code]/user/hand/right[/code]). [param openxr_path] is the action input/output being registered (e.g. [code]/user/hand/left/input/aim/pose[/code]). [param openxr_extension_name] restricts this input/output to an enabled/available extension, this doesn't need to repeat the extension on the profile but relates to overlapping extension (e.g. [code]XR_EXT_palm_pose[/code] that introduces [code]…/input/palm_ext/pose[/code] input paths). [param action_type] defines the type of input or output provided by OpenXR.
*/
//go:nosplit
func (self class) RegisterIoPath(interaction_profile gd.String, display_name gd.String, toplevel_path gd.String, openxr_path gd.String, openxr_extension_name gd.String, action_type classdb.OpenXRActionActionType)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(interaction_profile))
	callframe.Arg(frame, discreet.Get(display_name))
	callframe.Arg(frame, discreet.Get(toplevel_path))
	callframe.Arg(frame, discreet.Get(openxr_path))
	callframe.Arg(frame, discreet.Get(openxr_extension_name))
	callframe.Arg(frame, action_type)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInteractionProfileMetadata.Bind_register_io_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsOpenXRInteractionProfileMetadata() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsOpenXRInteractionProfileMetadata() Go { return *((*Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {classdb.Register("OpenXRInteractionProfileMetadata", func(ptr gd.Object) any { return classdb.OpenXRInteractionProfileMetadata(ptr) })}
