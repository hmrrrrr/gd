package AnimationNodeOneShot

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/AnimationNodeSync"
import "grow.graphics/gd/objects/AnimationNode"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A resource to add to an [AnimationNodeBlendTree]. This animation node will execute a sub-animation and return once it finishes. Blend times for fading in and out can be customized, as well as filters.
After setting the request and changing the animation playback, the one-shot node automatically clears the request on the next process frame by setting its [code]request[/code] value to [constant ONE_SHOT_REQUEST_NONE].
[codeblocks]
[gdscript]
# Play child animation connected to "shot" port.
animation_tree.set("parameters/OneShot/request", AnimationNodeOneShot.ONE_SHOT_REQUEST_FIRE)
# Alternative syntax (same result as above).
animation_tree["parameters/OneShot/request"] = AnimationNodeOneShot.ONE_SHOT_REQUEST_FIRE

# Abort child animation connected to "shot" port.
animation_tree.set("parameters/OneShot/request", AnimationNodeOneShot.ONE_SHOT_REQUEST_ABORT)
# Alternative syntax (same result as above).
animation_tree["parameters/OneShot/request"] = AnimationNodeOneShot.ONE_SHOT_REQUEST_ABORT

# Abort child animation with fading out connected to "shot" port.
animation_tree.set("parameters/OneShot/request", AnimationNodeOneShot.ONE_SHOT_REQUEST_FADE_OUT)
# Alternative syntax (same result as above).
animation_tree["parameters/OneShot/request"] = AnimationNodeOneShot.ONE_SHOT_REQUEST_FADE_OUT

# Get current state (read-only).
animation_tree.get("parameters/OneShot/active")
# Alternative syntax (same result as above).
animation_tree["parameters/OneShot/active"]

# Get current internal state (read-only).
animation_tree.get("parameters/OneShot/internal_active")
# Alternative syntax (same result as above).
animation_tree["parameters/OneShot/internal_active"]
[/gdscript]
[csharp]
// Play child animation connected to "shot" port.
animationTree.Set("parameters/OneShot/request", (int)AnimationNodeOneShot.OneShotRequest.Fire);

// Abort child animation connected to "shot" port.
animationTree.Set("parameters/OneShot/request", (int)AnimationNodeOneShot.OneShotRequest.Abort);

// Abort child animation with fading out connected to "shot" port.
animationTree.Set("parameters/OneShot/request", (int)AnimationNodeOneShot.OneShotRequest.FadeOut);

// Get current state (read-only).
animationTree.Get("parameters/OneShot/active");

// Get current internal state (read-only).
animationTree.Get("parameters/OneShot/internal_active");
[/csharp]
[/codeblocks]
*/
type Instance [1]classdb.AnimationNodeOneShot

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AnimationNodeOneShot

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimationNodeOneShot"))
	return Instance{classdb.AnimationNodeOneShot(object)}
}

func (self Instance) MixMode() classdb.AnimationNodeOneShotMixMode {
	return classdb.AnimationNodeOneShotMixMode(class(self).GetMixMode())
}

func (self Instance) SetMixMode(value classdb.AnimationNodeOneShotMixMode) {
	class(self).SetMixMode(value)
}

func (self Instance) FadeinTime() float64 {
	return float64(float64(class(self).GetFadeinTime()))
}

func (self Instance) SetFadeinTime(value float64) {
	class(self).SetFadeinTime(gd.Float(value))
}

func (self Instance) FadeinCurve() objects.Curve {
	return objects.Curve(class(self).GetFadeinCurve())
}

func (self Instance) SetFadeinCurve(value objects.Curve) {
	class(self).SetFadeinCurve(value)
}

func (self Instance) FadeoutTime() float64 {
	return float64(float64(class(self).GetFadeoutTime()))
}

func (self Instance) SetFadeoutTime(value float64) {
	class(self).SetFadeoutTime(gd.Float(value))
}

func (self Instance) FadeoutCurve() objects.Curve {
	return objects.Curve(class(self).GetFadeoutCurve())
}

func (self Instance) SetFadeoutCurve(value objects.Curve) {
	class(self).SetFadeoutCurve(value)
}

func (self Instance) BreakLoopAtEnd() bool {
	return bool(class(self).IsLoopBrokenAtEnd())
}

func (self Instance) SetBreakLoopAtEnd(value bool) {
	class(self).SetBreakLoopAtEnd(value)
}

func (self Instance) Autorestart() bool {
	return bool(class(self).HasAutorestart())
}

func (self Instance) SetAutorestart(value bool) {
	class(self).SetAutorestart(value)
}

func (self Instance) AutorestartDelay() float64 {
	return float64(float64(class(self).GetAutorestartDelay()))
}

func (self Instance) SetAutorestartDelay(value float64) {
	class(self).SetAutorestartDelay(gd.Float(value))
}

func (self Instance) AutorestartRandomDelay() float64 {
	return float64(float64(class(self).GetAutorestartRandomDelay()))
}

func (self Instance) SetAutorestartRandomDelay(value float64) {
	class(self).SetAutorestartRandomDelay(gd.Float(value))
}

//go:nosplit
func (self class) SetFadeinTime(time gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, time)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_set_fadein_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFadeinTime() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_get_fadein_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFadeinCurve(curve objects.Curve) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(curve[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_set_fadein_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFadeinCurve() objects.Curve {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_get_fadein_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Curve{classdb.Curve(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFadeoutTime(time gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, time)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_set_fadeout_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFadeoutTime() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_get_fadeout_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFadeoutCurve(curve objects.Curve) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(curve[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_set_fadeout_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFadeoutCurve() objects.Curve {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_get_fadeout_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Curve{classdb.Curve(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBreakLoopAtEnd(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_set_break_loop_at_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsLoopBrokenAtEnd() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_is_loop_broken_at_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutorestart(active bool) {
	var frame = callframe.New()
	callframe.Arg(frame, active)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_set_autorestart, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) HasAutorestart() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_has_autorestart, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutorestartDelay(time gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, time)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_set_autorestart_delay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutorestartDelay() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_get_autorestart_delay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutorestartRandomDelay(time gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, time)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_set_autorestart_random_delay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutorestartRandomDelay() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_get_autorestart_random_delay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMixMode(mode classdb.AnimationNodeOneShotMixMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_set_mix_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMixMode() classdb.AnimationNodeOneShotMixMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AnimationNodeOneShotMixMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationNodeOneShot.Bind_get_mix_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAnimationNodeOneShot() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAnimationNodeOneShot() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsAnimationNodeSync() AnimationNodeSync.Advanced {
	return *((*AnimationNodeSync.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAnimationNodeSync() AnimationNodeSync.Instance {
	return *((*AnimationNodeSync.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAnimationNode() AnimationNode.Advanced {
	return *((*AnimationNode.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAnimationNode() AnimationNode.Instance {
	return *((*AnimationNode.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsAnimationNodeSync(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsAnimationNodeSync(), name)
	}
}
func init() {
	classdb.Register("AnimationNodeOneShot", func(ptr gd.Object) any { return classdb.AnimationNodeOneShot(ptr) })
}

type OneShotRequest = classdb.AnimationNodeOneShotOneShotRequest

const (
	/*The default state of the request. Nothing is done.*/
	OneShotRequestNone OneShotRequest = 0
	/*The request to play the animation connected to "shot" port.*/
	OneShotRequestFire OneShotRequest = 1
	/*The request to stop the animation connected to "shot" port.*/
	OneShotRequestAbort OneShotRequest = 2
	/*The request to fade out the animation connected to "shot" port.*/
	OneShotRequestFadeOut OneShotRequest = 3
)

type MixMode = classdb.AnimationNodeOneShotMixMode

const (
	/*Blends two animations. See also [AnimationNodeBlend2].*/
	MixModeBlend MixMode = 0
	/*Blends two animations additively. See also [AnimationNodeAdd2].*/
	MixModeAdd MixMode = 1
)
