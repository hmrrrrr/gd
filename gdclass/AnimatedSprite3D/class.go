package AnimatedSprite3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/SpriteBase3D"
import "grow.graphics/gd/gdclass/GeometryInstance3D"
import "grow.graphics/gd/gdclass/VisualInstance3D"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
[AnimatedSprite3D] is similar to the [Sprite3D] node, except it carries multiple textures as animation [member sprite_frames]. Animations are created using a [SpriteFrames] resource, which allows you to import image files (or a folder containing said files) to provide the animation frames for the sprite. The [SpriteFrames] resource can be configured in the editor via the SpriteFrames bottom panel.

*/
type Go [1]classdb.AnimatedSprite3D

/*
Returns [code]true[/code] if an animation is currently playing (even if [member speed_scale] and/or [code]custom_speed[/code] are [code]0[/code]).
*/
func (self Go) IsPlaying() bool {
	return bool(class(self).IsPlaying())
}

/*
Plays the animation with key [param name]. If [param custom_speed] is negative and [param from_end] is [code]true[/code], the animation will play backwards (which is equivalent to calling [method play_backwards]).
If this method is called with that same animation [param name], or with no [param name] parameter, the assigned animation will resume playing if it was paused.
*/
func (self Go) Play() {
	class(self).Play(gd.NewStringName(""), gd.Float(1.0), false)
}

/*
Plays the animation with key [param name] in reverse.
This method is a shorthand for [method play] with [code]custom_speed = -1.0[/code] and [code]from_end = true[/code], so see its description for more information.
*/
func (self Go) PlayBackwards() {
	class(self).PlayBackwards(gd.NewStringName(""))
}

/*
Pauses the currently playing animation. The [member frame] and [member frame_progress] will be kept and calling [method play] or [method play_backwards] without arguments will resume the animation from the current playback position.
See also [method stop].
*/
func (self Go) Pause() {
	class(self).Pause()
}

/*
Stops the currently playing animation. The animation position is reset to [code]0[/code] and the [code]custom_speed[/code] is reset to [code]1.0[/code]. See also [method pause].
*/
func (self Go) Stop() {
	class(self).Stop()
}

/*
The setter of [member frame] resets the [member frame_progress] to [code]0.0[/code] implicitly, but this method avoids that.
This is useful when you want to carry over the current [member frame_progress] to another [member frame].
[b]Example:[/b]
[codeblocks]
[gdscript]
# Change the animation with keeping the frame index and progress.
var current_frame = animated_sprite.get_frame()
var current_progress = animated_sprite.get_frame_progress()
animated_sprite.play("walk_another_skin")
animated_sprite.set_frame_and_progress(current_frame, current_progress)
[/gdscript]
[/codeblocks]
*/
func (self Go) SetFrameAndProgress(frame_ int, progress float64) {
	class(self).SetFrameAndProgress(gd.Int(frame_), gd.Float(progress))
}

/*
Returns the actual playing speed of current animation or [code]0[/code] if not playing. This speed is the [member speed_scale] property multiplied by [code]custom_speed[/code] argument specified when calling the [method play] method.
Returns a negative value if the current animation is playing backwards.
*/
func (self Go) GetPlayingSpeed() float64 {
	return float64(float64(class(self).GetPlayingSpeed()))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AnimatedSprite3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimatedSprite3D"))
	return Go{classdb.AnimatedSprite3D(object)}
}

func (self Go) SpriteFrames() gdclass.SpriteFrames {
		return gdclass.SpriteFrames(class(self).GetSpriteFrames())
}

func (self Go) SetSpriteFrames(value gdclass.SpriteFrames) {
	class(self).SetSpriteFrames(value)
}

func (self Go) Animation() string {
		return string(class(self).GetAnimation().String())
}

func (self Go) SetAnimation(value string) {
	class(self).SetAnimation(gd.NewStringName(value))
}

func (self Go) Autoplay() string {
		return string(class(self).GetAutoplay().String())
}

func (self Go) SetAutoplay(value string) {
	class(self).SetAutoplay(gd.NewString(value))
}

func (self Go) Frame() int {
		return int(int(class(self).GetFrame()))
}

func (self Go) SetFrame(value int) {
	class(self).SetFrame(gd.Int(value))
}

func (self Go) FrameProgress() float64 {
		return float64(float64(class(self).GetFrameProgress()))
}

func (self Go) SetFrameProgress(value float64) {
	class(self).SetFrameProgress(gd.Float(value))
}

func (self Go) SpeedScale() float64 {
		return float64(float64(class(self).GetSpeedScale()))
}

func (self Go) SetSpeedScale(value float64) {
	class(self).SetSpeedScale(gd.Float(value))
}

//go:nosplit
func (self class) SetSpriteFrames(sprite_frames gdclass.SpriteFrames)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(sprite_frames[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_set_sprite_frames, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSpriteFrames() gdclass.SpriteFrames {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_get_sprite_frames, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.SpriteFrames{classdb.SpriteFrames(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAnimation(name gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_set_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAnimation() gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_get_animation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutoplay(name gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_set_autoplay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAutoplay() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_get_autoplay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if an animation is currently playing (even if [member speed_scale] and/or [code]custom_speed[/code] are [code]0[/code]).
*/
//go:nosplit
func (self class) IsPlaying() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_is_playing, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Plays the animation with key [param name]. If [param custom_speed] is negative and [param from_end] is [code]true[/code], the animation will play backwards (which is equivalent to calling [method play_backwards]).
If this method is called with that same animation [param name], or with no [param name] parameter, the assigned animation will resume playing if it was paused.
*/
//go:nosplit
func (self class) Play(name gd.StringName, custom_speed gd.Float, from_end bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
	callframe.Arg(frame, custom_speed)
	callframe.Arg(frame, from_end)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_play, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Plays the animation with key [param name] in reverse.
This method is a shorthand for [method play] with [code]custom_speed = -1.0[/code] and [code]from_end = true[/code], so see its description for more information.
*/
//go:nosplit
func (self class) PlayBackwards(name gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_play_backwards, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Pauses the currently playing animation. The [member frame] and [member frame_progress] will be kept and calling [method play] or [method play_backwards] without arguments will resume the animation from the current playback position.
See also [method stop].
*/
//go:nosplit
func (self class) Pause()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_pause, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Stops the currently playing animation. The animation position is reset to [code]0[/code] and the [code]custom_speed[/code] is reset to [code]1.0[/code]. See also [method pause].
*/
//go:nosplit
func (self class) Stop()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_stop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetFrame(frame_ gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, frame_)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_set_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFrame() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_get_frame, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFrameProgress(progress gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, progress)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_set_frame_progress, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFrameProgress() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_get_frame_progress, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
The setter of [member frame] resets the [member frame_progress] to [code]0.0[/code] implicitly, but this method avoids that.
This is useful when you want to carry over the current [member frame_progress] to another [member frame].
[b]Example:[/b]
[codeblocks]
[gdscript]
# Change the animation with keeping the frame index and progress.
var current_frame = animated_sprite.get_frame()
var current_progress = animated_sprite.get_frame_progress()
animated_sprite.play("walk_another_skin")
animated_sprite.set_frame_and_progress(current_frame, current_progress)
[/gdscript]
[/codeblocks]
*/
//go:nosplit
func (self class) SetFrameAndProgress(frame_ gd.Int, progress gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, frame_)
	callframe.Arg(frame, progress)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_set_frame_and_progress, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetSpeedScale(speed_scale gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, speed_scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_set_speed_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSpeedScale() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_get_speed_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedSprite3D.Bind_get_playing_speed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnSpriteFramesChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("sprite_frames_changed"), gd.NewCallable(cb), 0)
}


func (self Go) OnAnimationChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("animation_changed"), gd.NewCallable(cb), 0)
}


func (self Go) OnFrameChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("frame_changed"), gd.NewCallable(cb), 0)
}


func (self Go) OnAnimationLooped(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("animation_looped"), gd.NewCallable(cb), 0)
}


func (self Go) OnAnimationFinished(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("animation_finished"), gd.NewCallable(cb), 0)
}


func (self class) AsAnimatedSprite3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAnimatedSprite3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsSpriteBase3D() SpriteBase3D.GD { return *((*SpriteBase3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsSpriteBase3D() SpriteBase3D.Go { return *((*SpriteBase3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsGeometryInstance3D() GeometryInstance3D.GD { return *((*GeometryInstance3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsGeometryInstance3D() GeometryInstance3D.Go { return *((*GeometryInstance3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualInstance3D() VisualInstance3D.GD { return *((*VisualInstance3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualInstance3D() VisualInstance3D.Go { return *((*VisualInstance3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsSpriteBase3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsSpriteBase3D(), name)
	}
}
func init() {classdb.Register("AnimatedSprite3D", func(ptr gd.Object) any { return classdb.AnimatedSprite3D(ptr) })}
