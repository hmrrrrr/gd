package OpenXRAPIExtension

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[OpenXRAPIExtension] makes OpenXR available for GDExtension. It provides the OpenXR API to GDExtension through the [method get_instance_proc_addr] method, and the OpenXR instance through [method get_instance].
It also provides methods for querying the status of OpenXR initialization, and helper methods for ease of use of the API with GDExtension.

*/
type Simple [1]classdb.OpenXRAPIExtension
func (self Simple) GetInstance() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetInstance()))
}
func (self Simple) GetSystemId() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSystemId()))
}
func (self Simple) GetSession() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSession()))
}
func (self Simple) TransformFromPose(pose unsafe.Pointer) gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(Expert(self).TransformFromPose(pose))
}
func (self Simple) XrResult(result int, format string, args gd.Array) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).XrResult(gd.Int(result), gc.String(format), args))
}
func (self Simple) OpenxrIsEnabled(check_run_in_editor bool) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).OpenxrIsEnabled(gc, check_run_in_editor))
}
func (self Simple) GetInstanceProcAddr(name string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetInstanceProcAddr(gc.String(name))))
}
func (self Simple) GetErrorString(result int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetErrorString(gc, gd.Int(result)).String())
}
func (self Simple) GetSwapchainFormatName(swapchain_format int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetSwapchainFormatName(gc, gd.Int(swapchain_format)).String())
}
func (self Simple) IsInitialized() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsInitialized())
}
func (self Simple) IsRunning() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsRunning())
}
func (self Simple) GetPlaySpace() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetPlaySpace()))
}
func (self Simple) GetPredictedDisplayTime() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetPredictedDisplayTime()))
}
func (self Simple) GetNextFrameTime() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetNextFrameTime()))
}
func (self Simple) CanRender() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).CanRender())
}
func (self Simple) GetHandTracker(hand_index int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetHandTracker(gd.Int(hand_index))))
}
func (self Simple) RegisterCompositionLayerProvider(extension [1]classdb.OpenXRExtensionWrapperExtension) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RegisterCompositionLayerProvider(extension)
}
func (self Simple) UnregisterCompositionLayerProvider(extension [1]classdb.OpenXRExtensionWrapperExtension) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).UnregisterCompositionLayerProvider(extension)
}
func (self Simple) SetEmulateEnvironmentBlendModeAlphaBlend(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEmulateEnvironmentBlendModeAlphaBlend(enabled)
}
func (self Simple) IsEnvironmentBlendModeAlphaSupported() classdb.OpenXRAPIExtensionOpenXRAlphaBlendModeSupport {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.OpenXRAPIExtensionOpenXRAlphaBlendModeSupport(Expert(self).IsEnvironmentBlendModeAlphaSupported())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.OpenXRAPIExtension
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns the [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrInstance.html]XrInstance[/url] created during the initialization of the OpenXR API.
*/
//go:nosplit
func (self class) GetInstance() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRAPIExtension.Bind_get_instance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the id of the system, which is a [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrSystemId.html]XrSystemId[/url] cast to an integer.
*/
//go:nosplit
func (self class) GetSystemId() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRAPIExtension.Bind_get_system_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the OpenXR session, which is an [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrSession.html]XrSession[/url] cast to an integer.
*/
//go:nosplit
func (self class) GetSession() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRAPIExtension.Bind_get_session, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a [Transform3D] from an [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrPosef.html]XrPosef[/url].
*/
//go:nosplit
func (self class) TransformFromPose(pose unsafe.Pointer) gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pose)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRAPIExtension.Bind_transform_from_pose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the provided [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrResult.html]XrResult[/url] (cast to an integer) is successful. Otherwise returns [code]false[/code] and prints the [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrResult.html]XrResult[/url] converted to a string, with the specified additional information.
*/
//go:nosplit
func (self class) XrResult(result gd.Int, format gd.String, args gd.Array) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, result)
	callframe.Arg(frame, mmm.Get(format))
	callframe.Arg(frame, mmm.Get(args))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRAPIExtension.Bind_xr_result, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if OpenXR is enabled.
*/
//go:nosplit
func (self class) OpenxrIsEnabled(ctx gd.Lifetime, check_run_in_editor bool) bool {
	var frame = callframe.New()
	callframe.Arg(frame, check_run_in_editor)
	var r_ret = callframe.Ret[bool](frame)
	ctx.API.Object.MethodBindPointerCall(ctx.API.Methods.OpenXRAPIExtension.Bind_openxr_is_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the function pointer of the OpenXR function with the specified name, cast to an integer. If the function with the given name does not exist, the method returns [code]0[/code].
[b]Note:[/b] [code]openxr/util.h[/code] contains utility macros for acquiring OpenXR functions, e.g. [code]GDEXTENSION_INIT_XR_FUNC_V(xrCreateAction)[/code].
*/
//go:nosplit
func (self class) GetInstanceProcAddr(name gd.String) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRAPIExtension.Bind_get_instance_proc_addr, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an error string for the given [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrResult.html]XrResult[/url].
*/
//go:nosplit
func (self class) GetErrorString(ctx gd.Lifetime, result gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, result)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRAPIExtension.Bind_get_error_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the name of the specified swapchain format.
*/
//go:nosplit
func (self class) GetSwapchainFormatName(ctx gd.Lifetime, swapchain_format gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, swapchain_format)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRAPIExtension.Bind_get_swapchain_format_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if OpenXR is initialized.
*/
//go:nosplit
func (self class) IsInitialized() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRAPIExtension.Bind_is_initialized, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if OpenXR is running ([url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/xrBeginSession.html]xrBeginSession[/url] was successfully called and the swapchains were created).
*/
//go:nosplit
func (self class) IsRunning() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRAPIExtension.Bind_is_running, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the play space, which is an [url=https://registry.khronos.org/OpenXR/specs/1.0/man/html/XrSpace.html]XrSpace[/url] cast to an integer.
*/
//go:nosplit
func (self class) GetPlaySpace() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRAPIExtension.Bind_get_play_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the predicted display timing for the current frame.
*/
//go:nosplit
func (self class) GetPredictedDisplayTime() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRAPIExtension.Bind_get_predicted_display_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the predicted display timing for the next frame.
*/
//go:nosplit
func (self class) GetNextFrameTime() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRAPIExtension.Bind_get_next_frame_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if OpenXR is initialized for rendering with an XR viewport.
*/
//go:nosplit
func (self class) CanRender() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRAPIExtension.Bind_can_render, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the corresponding [code]XRHandTrackerEXT[/code] handle for the given hand index value.
*/
//go:nosplit
func (self class) GetHandTracker(hand_index gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hand_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRAPIExtension.Bind_get_hand_tracker, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Registers the given extension as a composition layer provider.
*/
//go:nosplit
func (self class) RegisterCompositionLayerProvider(extension object.OpenXRExtensionWrapperExtension)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(extension[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRAPIExtension.Bind_register_composition_layer_provider, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Unregisters the given extension as a composition layer provider.
*/
//go:nosplit
func (self class) UnregisterCompositionLayerProvider(extension object.OpenXRExtensionWrapperExtension)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(extension[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRAPIExtension.Bind_unregister_composition_layer_provider, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If set to [code]true[/code], an OpenXR extension is loaded which is capable of emulating the [constant XRInterface.XR_ENV_BLEND_MODE_ALPHA_BLEND] blend mode.
*/
//go:nosplit
func (self class) SetEmulateEnvironmentBlendModeAlphaBlend(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRAPIExtension.Bind_set_emulate_environment_blend_mode_alpha_blend, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [enum OpenXRAPIExtension.OpenXRAlphaBlendModeSupport] denoting if [constant XRInterface.XR_ENV_BLEND_MODE_ALPHA_BLEND] is really supported, emulated or not supported at all.
*/
//go:nosplit
func (self class) IsEnvironmentBlendModeAlphaSupported() classdb.OpenXRAPIExtensionOpenXRAlphaBlendModeSupport {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.OpenXRAPIExtensionOpenXRAlphaBlendModeSupport](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRAPIExtension.Bind_is_environment_blend_mode_alpha_supported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsOpenXRAPIExtension() Expert { return self[0].AsOpenXRAPIExtension() }


//go:nosplit
func (self Simple) AsOpenXRAPIExtension() Simple { return self[0].AsOpenXRAPIExtension() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("OpenXRAPIExtension", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type OpenXRAlphaBlendModeSupport = classdb.OpenXRAPIExtensionOpenXRAlphaBlendModeSupport

const (
/*Means that [constant XRInterface.XR_ENV_BLEND_MODE_ALPHA_BLEND] isn't supported at all.*/
	OpenxrAlphaBlendModeSupportNone OpenXRAlphaBlendModeSupport = 0
/*Means that [constant XRInterface.XR_ENV_BLEND_MODE_ALPHA_BLEND] is really supported.*/
	OpenxrAlphaBlendModeSupportReal OpenXRAlphaBlendModeSupport = 1
/*Means that [constant XRInterface.XR_ENV_BLEND_MODE_ALPHA_BLEND] is emulated.*/
	OpenxrAlphaBlendModeSupportEmulating OpenXRAlphaBlendModeSupport = 2
)
