package AnimationPlayer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/AnimationMixer"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
An animation player is used for general-purpose playback of animations. It contains a dictionary of [AnimationLibrary] resources and custom blend times between animation transitions.
Some methods and properties use a single key to reference an animation directly. These keys are formatted as the key for the library, followed by a forward slash, then the key for the animation within the library, for example [code]"movement/run"[/code]. If the library's key is an empty string (known as the default library), the forward slash is omitted, being the same key used by the library.
[AnimationPlayer] is better-suited than [Tween] for more complex animations, for example ones with non-trivial timings. It can also be used over [Tween] if the animation track editor is more convenient than doing it in code.
Updating the target properties of animations occurs at the process frame.
*/
type Instance [1]classdb.AnimationPlayer

/*
Triggers the [param animation_to] animation when the [param animation_from] animation completes.
*/
func (self Instance) AnimationSetNext(animation_from string, animation_to string) {
	class(self).AnimationSetNext(gd.NewStringName(animation_from), gd.NewStringName(animation_to))
}

/*
Returns the key of the animation which is queued to play after the [param animation_from] animation.
*/
func (self Instance) AnimationGetNext(animation_from string) string {
	return string(class(self).AnimationGetNext(gd.NewStringName(animation_from)).String())
}

/*
Specifies a blend time (in seconds) between two animations, referenced by their keys.
*/
func (self Instance) SetBlendTime(animation_from string, animation_to string, sec float64) {
	class(self).SetBlendTime(gd.NewStringName(animation_from), gd.NewStringName(animation_to), gd.Float(sec))
}

/*
Returns the blend time (in seconds) between two animations, referenced by their keys.
*/
func (self Instance) GetBlendTime(animation_from string, animation_to string) float64 {
	return float64(float64(class(self).GetBlendTime(gd.NewStringName(animation_from), gd.NewStringName(animation_to))))
}

/*
Plays the animation with key [param name]. Custom blend times and speed can be set.
The [param from_end] option only affects when switching to a new animation track, or if the same track but at the start or end. It does not affect resuming playback that was paused in the middle of an animation. If [param custom_speed] is negative and [param from_end] is [code]true[/code], the animation will play backwards (which is equivalent to calling [method play_backwards]).
The [AnimationPlayer] keeps track of its current or last played animation with [member assigned_animation]. If this method is called with that same animation [param name], or with no [param name] parameter, the assigned animation will resume playing if it was paused.
[b]Note:[/b] The animation will be updated the next time the [AnimationPlayer] is processed. If other variables are updated at the same time this is called, they may be updated too early. To perform the update immediately, call [code]advance(0)[/code].
*/
func (self Instance) Play() {
	class(self).Play(gd.NewStringName(""), gd.Float(-1), gd.Float(1.0), false)
}

/*
Plays the animation with key [param name] in reverse.
This method is a shorthand for [method play] with [code]custom_speed = -1.0[/code] and [code]from_end = true[/code], so see its description for more information.
*/
func (self Instance) PlayBackwards() {
	class(self).PlayBackwards(gd.NewStringName(""), gd.Float(-1))
}

/*
See also [method AnimationMixer.capture].
You can use this method to use more detailed options for capture than those performed by [member playback_auto_capture]. When [member playback_auto_capture] is [code]false[/code], this method is almost the same as the following:
[codeblock]
capture(name, duration, trans_type, ease_type)
play(name, custom_blend, custom_speed, from_end)
[/codeblock]
If [param name] is blank, it specifies [member assigned_animation].
If [param duration] is a negative value, the duration is set to the interval between the current position and the first key, when [param from_end] is [code]true[/code], uses the interval between the current position and the last key instead.
[b]Note:[/b] The [param duration] takes [member speed_scale] into account, but [param custom_speed] does not, because the capture cache is interpolated with the blend result and the result may contain multiple animations.
*/
func (self Instance) PlayWithCapture() {
	class(self).PlayWithCapture(gd.NewStringName(""), gd.Float(-1.0), gd.Float(-1), gd.Float(1.0), false, 0, 0)
}

/*
Pauses the currently playing animation. The [member current_animation_position] will be kept and calling [method play] or [method play_backwards] without arguments or with the same animation name as [member assigned_animation] will resume the animation.
See also [method stop].
*/
func (self Instance) Pause() {
	class(self).Pause()
}

/*
Stops the currently playing animation. The animation position is reset to [code]0[/code] and the [code]custom_speed[/code] is reset to [code]1.0[/code]. See also [method pause].
If [param keep_state] is [code]true[/code], the animation state is not updated visually.
[b]Note:[/b] The method / audio / animation playback tracks will not be processed by this method.
*/
func (self Instance) Stop() {
	class(self).Stop(false)
}

/*
Returns [code]true[/code] if an animation is currently playing (even if [member speed_scale] and/or [code]custom_speed[/code] are [code]0[/code]).
*/
func (self Instance) IsPlaying() bool {
	return bool(class(self).IsPlaying())
}

/*
Queues an animation for playback once the current animation and all previously queued animations are done.
[b]Note:[/b] If a looped animation is currently playing, the queued animation will never play unless the looped animation is stopped somehow.
*/
func (self Instance) Queue(name string) {
	class(self).Queue(gd.NewStringName(name))
}

/*
Returns a list of the animation keys that are currently queued to play.
*/
func (self Instance) GetQueue() []string {
	return []string(class(self).GetQueue().Strings())
}

/*
Clears all queued, unplayed animations.
*/
func (self Instance) ClearQueue() {
	class(self).ClearQueue()
}

/*
Returns the actual playing speed of current animation or [code]0[/code] if not playing. This speed is the [member speed_scale] property multiplied by [code]custom_speed[/code] argument specified when calling the [method play] method.
Returns a negative value if the current animation is playing backwards.
*/
func (self Instance) GetPlayingSpeed() float64 {
	return float64(float64(class(self).GetPlayingSpeed()))
}

/*
Seeks the animation to the [param seconds] point in time (in seconds). If [param update] is [code]true[/code], the animation updates too, otherwise it updates at process time. Events between the current frame and [param seconds] are skipped.
If [param update_only] is [code]true[/code], the method / audio / animation playback tracks will not be processed.
[b]Note:[/b] Seeking to the end of the animation doesn't emit [signal AnimationMixer.animation_finished]. If you want to skip animation and emit the signal, use [method AnimationMixer.advance].
*/
func (self Instance) SeekTo(seconds float64) {
	class(self).SeekTo(gd.Float(seconds), false, false)
}

/*
Sets the process notification in which to update animations.
*/
func (self Instance) SetProcessCallback(mode classdb.AnimationPlayerAnimationProcessCallback) {
	class(self).SetProcessCallback(mode)
}

/*
Returns the process notification in which to update animations.
*/
func (self Instance) GetProcessCallback() classdb.AnimationPlayerAnimationProcessCallback {
	return classdb.AnimationPlayerAnimationProcessCallback(class(self).GetProcessCallback())
}

/*
Sets the call mode used for "Call Method" tracks.
*/
func (self Instance) SetMethodCallMode(mode classdb.AnimationPlayerAnimationMethodCallMode) {
	class(self).SetMethodCallMode(mode)
}

/*
Returns the call mode used for "Call Method" tracks.
*/
func (self Instance) GetMethodCallMode() classdb.AnimationPlayerAnimationMethodCallMode {
	return classdb.AnimationPlayerAnimationMethodCallMode(class(self).GetMethodCallMode())
}

/*
Sets the node which node path references will travel from.
*/
func (self Instance) SetRoot(path string) {
	class(self).SetRoot(gd.NewString(path).NodePath())
}

/*
Returns the node which node path references will travel from.
*/
func (self Instance) GetRoot() string {
	return string(class(self).GetRoot().String())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AnimationPlayer

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimationPlayer"))
	return Instance{classdb.AnimationPlayer(object)}
}

func (self Instance) CurrentAnimation() string {
	return string(class(self).GetCurrentAnimation().String())
}

func (self Instance) SetCurrentAnimation(value string) {
	class(self).SetCurrentAnimation(gd.NewString(value))
}

func (self Instance) AssignedAnimation() string {
	return string(class(self).GetAssignedAnimation().String())
}

func (self Instance) SetAssignedAnimation(value string) {
	class(self).SetAssignedAnimation(gd.NewString(value))
}

func (self Instance) Autoplay() string {
	return string(class(self).GetAutoplay().String())
}

func (self Instance) SetAutoplay(value string) {
	class(self).SetAutoplay(gd.NewString(value))
}

func (self Instance) CurrentAnimationLength() float64 {
	return float64(float64(class(self).GetCurrentAnimationLength()))
}

func (self Instance) CurrentAnimationPosition() float64 {
	return float64(float64(class(self).GetCurrentAnimationPosition()))
}

func (self Instance) PlaybackAutoCapture() bool {
	return bool(class(self).IsAutoCapture())
}

func (self Instance) SetPlaybackAutoCapture(value bool) {
	class(self).SetAutoCapture(value)
}

func (self Instance) PlaybackAutoCaptureDuration() float64 {
	return float64(float64(class(self).GetAutoCaptureDuration()))
}

func (self Instance) SetPlaybackAutoCaptureDuration(value float64) {
	class(self).SetAutoCaptureDuration(gd.Float(value))
}

func (self Instance) PlaybackAutoCaptureTransitionType() classdb.TweenTransitionType {
	return classdb.TweenTransitionType(class(self).GetAutoCaptureTransitionType())
}

func (self Instance) SetPlaybackAutoCaptureTransitionType(value classdb.TweenTransitionType) {
	class(self).SetAutoCaptureTransitionType(value)
}

func (self Instance) PlaybackAutoCaptureEaseType() classdb.TweenEaseType {
	return classdb.TweenEaseType(class(self).GetAutoCaptureEaseType())
}

func (self Instance) SetPlaybackAutoCaptureEaseType(value classdb.TweenEaseType) {
	class(self).SetAutoCaptureEaseType(value)
}

func (self Instance) PlaybackDefaultBlendTime() float64 {
	return float64(float64(class(self).GetDefaultBlendTime()))
}

func (self Instance) SetPlaybackDefaultBlendTime(value float64) {
	class(self).SetDefaultBlendTime(gd.Float(value))
}

func (self Instance) SpeedScale() float64 {
	return float64(float64(class(self).GetSpeedScale()))
}

func (self Instance) SetSpeedScale(value float64) {
	class(self).SetSpeedScale(gd.Float(value))
}

func (self Instance) MovieQuitOnFinish() bool {
	return bool(class(self).IsMovieQuitOnFinishEnabled())
}

func (self Instance) SetMovieQuitOnFinish(value bool) {
	class(self).SetMovieQuitOnFinishEnabled(value)
}

/*
Triggers the [param animation_to] animation when the [param animation_from] animation completes.
*/
//go:nosplit
func (self class) AnimationSetNext(animation_from gd.StringName, animation_to gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(animation_from))
	callframe.Arg(frame, pointers.Get(animation_to))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_animation_set_next, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the key of the animation which is queued to play after the [param animation_from] animation.
*/
//go:nosplit
func (self class) AnimationGetNext(animation_from gd.StringName) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(animation_from))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_animation_get_next, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

/*
Specifies a blend time (in seconds) between two animations, referenced by their keys.
*/
//go:nosplit
func (self class) SetBlendTime(animation_from gd.StringName, animation_to gd.StringName, sec gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(animation_from))
	callframe.Arg(frame, pointers.Get(animation_to))
	callframe.Arg(frame, sec)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_set_blend_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the blend time (in seconds) between two animations, referenced by their keys.
*/
//go:nosplit
func (self class) GetBlendTime(animation_from gd.StringName, animation_to gd.StringName) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(animation_from))
	callframe.Arg(frame, pointers.Get(animation_to))
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_get_blend_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDefaultBlendTime(sec gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, sec)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_set_default_blend_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDefaultBlendTime() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_get_default_blend_time, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoCapture(auto_capture bool) {
	var frame = callframe.New()
	callframe.Arg(frame, auto_capture)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_set_auto_capture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsAutoCapture() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_is_auto_capture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoCaptureDuration(auto_capture_duration gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, auto_capture_duration)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_set_auto_capture_duration, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutoCaptureDuration() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_get_auto_capture_duration, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoCaptureTransitionType(auto_capture_transition_type classdb.TweenTransitionType) {
	var frame = callframe.New()
	callframe.Arg(frame, auto_capture_transition_type)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_set_auto_capture_transition_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutoCaptureTransitionType() classdb.TweenTransitionType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TweenTransitionType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_get_auto_capture_transition_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoCaptureEaseType(auto_capture_ease_type classdb.TweenEaseType) {
	var frame = callframe.New()
	callframe.Arg(frame, auto_capture_ease_type)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_set_auto_capture_ease_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutoCaptureEaseType() classdb.TweenEaseType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TweenEaseType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_get_auto_capture_ease_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Plays the animation with key [param name]. Custom blend times and speed can be set.
The [param from_end] option only affects when switching to a new animation track, or if the same track but at the start or end. It does not affect resuming playback that was paused in the middle of an animation. If [param custom_speed] is negative and [param from_end] is [code]true[/code], the animation will play backwards (which is equivalent to calling [method play_backwards]).
The [AnimationPlayer] keeps track of its current or last played animation with [member assigned_animation]. If this method is called with that same animation [param name], or with no [param name] parameter, the assigned animation will resume playing if it was paused.
[b]Note:[/b] The animation will be updated the next time the [AnimationPlayer] is processed. If other variables are updated at the same time this is called, they may be updated too early. To perform the update immediately, call [code]advance(0)[/code].
*/
//go:nosplit
func (self class) Play(name gd.StringName, custom_blend gd.Float, custom_speed gd.Float, from_end bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, custom_blend)
	callframe.Arg(frame, custom_speed)
	callframe.Arg(frame, from_end)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_play, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Plays the animation with key [param name] in reverse.
This method is a shorthand for [method play] with [code]custom_speed = -1.0[/code] and [code]from_end = true[/code], so see its description for more information.
*/
//go:nosplit
func (self class) PlayBackwards(name gd.StringName, custom_blend gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, custom_blend)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_play_backwards, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
See also [method AnimationMixer.capture].
You can use this method to use more detailed options for capture than those performed by [member playback_auto_capture]. When [member playback_auto_capture] is [code]false[/code], this method is almost the same as the following:
[codeblock]
capture(name, duration, trans_type, ease_type)
play(name, custom_blend, custom_speed, from_end)
[/codeblock]
If [param name] is blank, it specifies [member assigned_animation].
If [param duration] is a negative value, the duration is set to the interval between the current position and the first key, when [param from_end] is [code]true[/code], uses the interval between the current position and the last key instead.
[b]Note:[/b] The [param duration] takes [member speed_scale] into account, but [param custom_speed] does not, because the capture cache is interpolated with the blend result and the result may contain multiple animations.
*/
//go:nosplit
func (self class) PlayWithCapture(name gd.StringName, duration gd.Float, custom_blend gd.Float, custom_speed gd.Float, from_end bool, trans_type classdb.TweenTransitionType, ease_type classdb.TweenEaseType) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, duration)
	callframe.Arg(frame, custom_blend)
	callframe.Arg(frame, custom_speed)
	callframe.Arg(frame, from_end)
	callframe.Arg(frame, trans_type)
	callframe.Arg(frame, ease_type)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_play_with_capture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Pauses the currently playing animation. The [member current_animation_position] will be kept and calling [method play] or [method play_backwards] without arguments or with the same animation name as [member assigned_animation] will resume the animation.
See also [method stop].
*/
//go:nosplit
func (self class) Pause() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_pause, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Stops the currently playing animation. The animation position is reset to [code]0[/code] and the [code]custom_speed[/code] is reset to [code]1.0[/code]. See also [method pause].
If [param keep_state] is [code]true[/code], the animation state is not updated visually.
[b]Note:[/b] The method / audio / animation playback tracks will not be processed by this method.
*/
//go:nosplit
func (self class) Stop(keep_state bool) {
	var frame = callframe.New()
	callframe.Arg(frame, keep_state)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_stop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns [code]true[/code] if an animation is currently playing (even if [member speed_scale] and/or [code]custom_speed[/code] are [code]0[/code]).
*/
//go:nosplit
func (self class) IsPlaying() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_is_playing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCurrentAnimation(animation gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(animation))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_set_current_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCurrentAnimation() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_get_current_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAssignedAnimation(animation gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(animation))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_set_assigned_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAssignedAnimation() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_get_assigned_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Queues an animation for playback once the current animation and all previously queued animations are done.
[b]Note:[/b] If a looped animation is currently playing, the queued animation will never play unless the looped animation is stopped somehow.
*/
//go:nosplit
func (self class) Queue(name gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_queue, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a list of the animation keys that are currently queued to play.
*/
//go:nosplit
func (self class) GetQueue() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_get_queue, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Clears all queued, unplayed animations.
*/
//go:nosplit
func (self class) ClearQueue() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_clear_queue, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetSpeedScale(speed gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, speed)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_set_speed_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSpeedScale() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_get_speed_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the actual playing speed of current animation or [code]0[/code] if not playing. This speed is the [member speed_scale] property multiplied by [code]custom_speed[/code] argument specified when calling the [method play] method.
Returns a negative value if the current animation is playing backwards.
*/
//go:nosplit
func (self class) GetPlayingSpeed() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_get_playing_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoplay(name gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_set_autoplay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAutoplay() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_get_autoplay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMovieQuitOnFinishEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_set_movie_quit_on_finish_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsMovieQuitOnFinishEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_is_movie_quit_on_finish_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetCurrentAnimationPosition() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_get_current_animation_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetCurrentAnimationLength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_get_current_animation_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Seeks the animation to the [param seconds] point in time (in seconds). If [param update] is [code]true[/code], the animation updates too, otherwise it updates at process time. Events between the current frame and [param seconds] are skipped.
If [param update_only] is [code]true[/code], the method / audio / animation playback tracks will not be processed.
[b]Note:[/b] Seeking to the end of the animation doesn't emit [signal AnimationMixer.animation_finished]. If you want to skip animation and emit the signal, use [method AnimationMixer.advance].
*/
//go:nosplit
func (self class) SeekTo(seconds gd.Float, update bool, update_only bool) {
	var frame = callframe.New()
	callframe.Arg(frame, seconds)
	callframe.Arg(frame, update)
	callframe.Arg(frame, update_only)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_seek, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the process notification in which to update animations.
*/
//go:nosplit
func (self class) SetProcessCallback(mode classdb.AnimationPlayerAnimationProcessCallback) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_set_process_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the process notification in which to update animations.
*/
//go:nosplit
func (self class) GetProcessCallback() classdb.AnimationPlayerAnimationProcessCallback {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AnimationPlayerAnimationProcessCallback](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_get_process_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the call mode used for "Call Method" tracks.
*/
//go:nosplit
func (self class) SetMethodCallMode(mode classdb.AnimationPlayerAnimationMethodCallMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_set_method_call_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the call mode used for "Call Method" tracks.
*/
//go:nosplit
func (self class) GetMethodCallMode() classdb.AnimationPlayerAnimationMethodCallMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AnimationPlayerAnimationMethodCallMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_get_method_call_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the node which node path references will travel from.
*/
//go:nosplit
func (self class) SetRoot(path gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_set_root, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the node which node path references will travel from.
*/
//go:nosplit
func (self class) GetRoot() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationPlayer.Bind_get_root, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}
func (self Instance) OnCurrentAnimationChanged(cb func(name string)) {
	self[0].AsObject().Connect(gd.NewStringName("current_animation_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnAnimationChanged(cb func(old_name string, new_name string)) {
	self[0].AsObject().Connect(gd.NewStringName("animation_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsAnimationPlayer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAnimationPlayer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsAnimationMixer() AnimationMixer.Advanced {
	return *((*AnimationMixer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAnimationMixer() AnimationMixer.Instance {
	return *((*AnimationMixer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsAnimationMixer(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsAnimationMixer(), name)
	}
}
func init() {
	classdb.Register("AnimationPlayer", func(ptr gd.Object) any { return classdb.AnimationPlayer(ptr) })
}

type AnimationProcessCallback = classdb.AnimationPlayerAnimationProcessCallback

const (
	AnimationProcessPhysics AnimationProcessCallback = 0
	AnimationProcessIdle    AnimationProcessCallback = 1
	AnimationProcessManual  AnimationProcessCallback = 2
)

type AnimationMethodCallMode = classdb.AnimationPlayerAnimationMethodCallMode

const (
	AnimationMethodCallDeferred  AnimationMethodCallMode = 0
	AnimationMethodCallImmediate AnimationMethodCallMode = 1
)
