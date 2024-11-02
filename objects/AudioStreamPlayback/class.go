package AudioStreamPlayback

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Can play, loop, pause a scroll through audio. See [AudioStream] and [AudioStreamOggVorbis] for usage.

	// AudioStreamPlayback methods that can be overridden by a [Class] that extends it.
	type AudioStreamPlayback interface {
		//Override this method to customize what happens when the playback starts at the given position, such as by calling [method AudioStreamPlayer.play].
		Start(from_pos Float.X)
		//Override this method to customize what happens when the playback is stopped, such as by calling [method AudioStreamPlayer.stop].
		Stop()
		//Overridable method. Should return [code]true[/code] if this playback is active and playing its audio stream.
		IsPlaying() bool
		//Overridable method. Should return how many times this audio stream has looped. Most built-in playbacks always return [code]0[/code].
		GetLoopCount() int
		//Overridable method. Should return the current progress along the audio stream, in seconds.
		GetPlaybackPosition() Float.X
		//Override this method to customize what happens when seeking this audio stream at the given [param position], such as by calling [method AudioStreamPlayer.seek].
		Seek(position Float.X)
		//Override this method to customize how the audio stream is mixed. This method is called even if the playback is not active.
		//[b]Note:[/b] It is not useful to override this method in GDScript or C#. Only GDExtension can take advantage of it.
		Mix(buffer *classdb.AudioFrame, rate_scale Float.X, frames int) int
		//Overridable method. Called whenever the audio stream is mixed if the playback is active and [method AudioServer.set_enable_tagging_used_audio_streams] has been set to [code]true[/code]. Editor plugins may use this method to "tag" the current position along the audio stream and display it in a preview.
		TagUsedStreams()
		//Set the current value of a playback parameter by name (see [method AudioStream._get_parameter_list]).
		SetParameter(name string, value any)
		//Return the current value of a playback parameter by name (see [method AudioStream._get_parameter_list]).
		GetParameter(name string) any
	}
*/
type Instance [1]classdb.AudioStreamPlayback

/*
Override this method to customize what happens when the playback starts at the given position, such as by calling [method AudioStreamPlayer.play].
*/
func (Instance) _start(impl func(ptr unsafe.Pointer, from_pos Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var from_pos = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, Float.X(from_pos))
	}
}

/*
Override this method to customize what happens when the playback is stopped, such as by calling [method AudioStreamPlayer.stop].
*/
func (Instance) _stop(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Overridable method. Should return [code]true[/code] if this playback is active and playing its audio stream.
*/
func (Instance) _is_playing(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable method. Should return how many times this audio stream has looped. Most built-in playbacks always return [code]0[/code].
*/
func (Instance) _get_loop_count(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Overridable method. Should return the current progress along the audio stream, in seconds.
*/
func (Instance) _get_playback_position(impl func(ptr unsafe.Pointer) Float.X) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Float(ret))
	}
}

/*
Override this method to customize what happens when seeking this audio stream at the given [param position], such as by calling [method AudioStreamPlayer.seek].
*/
func (Instance) _seek(impl func(ptr unsafe.Pointer, position Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var position = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, Float.X(position))
	}
}

/*
Override this method to customize how the audio stream is mixed. This method is called even if the playback is not active.
[b]Note:[/b] It is not useful to override this method in GDScript or C#. Only GDExtension can take advantage of it.
*/
func (Instance) _mix(impl func(ptr unsafe.Pointer, buffer *classdb.AudioFrame, rate_scale Float.X, frames int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var buffer = gd.UnsafeGet[*classdb.AudioFrame](p_args, 0)
		var rate_scale = gd.UnsafeGet[gd.Float](p_args, 1)
		var frames = gd.UnsafeGet[gd.Int](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, buffer, Float.X(rate_scale), int(frames))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Overridable method. Called whenever the audio stream is mixed if the playback is active and [method AudioServer.set_enable_tagging_used_audio_streams] has been set to [code]true[/code]. Editor plugins may use this method to "tag" the current position along the audio stream and display it in a preview.
*/
func (Instance) _tag_used_streams(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Set the current value of a playback parameter by name (see [method AudioStream._get_parameter_list]).
*/
func (Instance) _set_parameter(impl func(ptr unsafe.Pointer, name string, value any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var name = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(name)
		var value = pointers.New[gd.Variant](gd.UnsafeGet[[3]uintptr](p_args, 1))
		defer pointers.End(value)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, name.String(), value.Interface())
	}
}

/*
Return the current value of a playback parameter by name (see [method AudioStream._get_parameter_list]).
*/
func (Instance) _get_parameter(impl func(ptr unsafe.Pointer, name string) any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var name = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(name)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, name.String())
		ptr, ok := pointers.End(gd.NewVariant(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Associates [AudioSamplePlayback] to this [AudioStreamPlayback] for playing back the audio sample of this stream.
*/
func (self Instance) SetSamplePlayback(playback_sample objects.AudioSamplePlayback) {
	class(self).SetSamplePlayback(playback_sample)
}

/*
Returns the [AudioSamplePlayback] associated with this [AudioStreamPlayback] for playing back the audio sample of this stream.
*/
func (self Instance) GetSamplePlayback() objects.AudioSamplePlayback {
	return objects.AudioSamplePlayback(class(self).GetSamplePlayback())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AudioStreamPlayback

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamPlayback"))
	return Instance{classdb.AudioStreamPlayback(object)}
}

/*
Override this method to customize what happens when the playback starts at the given position, such as by calling [method AudioStreamPlayer.play].
*/
func (class) _start(impl func(ptr unsafe.Pointer, from_pos gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var from_pos = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, from_pos)
	}
}

/*
Override this method to customize what happens when the playback is stopped, such as by calling [method AudioStreamPlayer.stop].
*/
func (class) _stop(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Overridable method. Should return [code]true[/code] if this playback is active and playing its audio stream.
*/
func (class) _is_playing(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable method. Should return how many times this audio stream has looped. Most built-in playbacks always return [code]0[/code].
*/
func (class) _get_loop_count(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable method. Should return the current progress along the audio stream, in seconds.
*/
func (class) _get_playback_position(impl func(ptr unsafe.Pointer) gd.Float) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to customize what happens when seeking this audio stream at the given [param position], such as by calling [method AudioStreamPlayer.seek].
*/
func (class) _seek(impl func(ptr unsafe.Pointer, position gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var position = gd.UnsafeGet[gd.Float](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, position)
	}
}

/*
Override this method to customize how the audio stream is mixed. This method is called even if the playback is not active.
[b]Note:[/b] It is not useful to override this method in GDScript or C#. Only GDExtension can take advantage of it.
*/
func (class) _mix(impl func(ptr unsafe.Pointer, buffer *classdb.AudioFrame, rate_scale gd.Float, frames gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var buffer = gd.UnsafeGet[*classdb.AudioFrame](p_args, 0)
		var rate_scale = gd.UnsafeGet[gd.Float](p_args, 1)
		var frames = gd.UnsafeGet[gd.Int](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, buffer, rate_scale, frames)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Overridable method. Called whenever the audio stream is mixed if the playback is active and [method AudioServer.set_enable_tagging_used_audio_streams] has been set to [code]true[/code]. Editor plugins may use this method to "tag" the current position along the audio stream and display it in a preview.
*/
func (class) _tag_used_streams(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Set the current value of a playback parameter by name (see [method AudioStream._get_parameter_list]).
*/
func (class) _set_parameter(impl func(ptr unsafe.Pointer, name gd.StringName, value gd.Variant)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var name = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		var value = pointers.New[gd.Variant](gd.UnsafeGet[[3]uintptr](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, name, value)
	}
}

/*
Return the current value of a playback parameter by name (see [method AudioStream._get_parameter_list]).
*/
func (class) _get_parameter(impl func(ptr unsafe.Pointer, name gd.StringName) gd.Variant) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var name = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, name)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Associates [AudioSamplePlayback] to this [AudioStreamPlayback] for playing back the audio sample of this stream.
*/
//go:nosplit
func (self class) SetSamplePlayback(playback_sample objects.AudioSamplePlayback) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(playback_sample[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayback.Bind_set_sample_playback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [AudioSamplePlayback] associated with this [AudioStreamPlayback] for playing back the audio sample of this stream.
*/
//go:nosplit
func (self class) GetSamplePlayback() objects.AudioSamplePlayback {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayback.Bind_get_sample_playback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.AudioSamplePlayback{classdb.AudioSamplePlayback(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self class) AsAudioStreamPlayback() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioStreamPlayback() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted        { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted     { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_start":
		return reflect.ValueOf(self._start)
	case "_stop":
		return reflect.ValueOf(self._stop)
	case "_is_playing":
		return reflect.ValueOf(self._is_playing)
	case "_get_loop_count":
		return reflect.ValueOf(self._get_loop_count)
	case "_get_playback_position":
		return reflect.ValueOf(self._get_playback_position)
	case "_seek":
		return reflect.ValueOf(self._seek)
	case "_mix":
		return reflect.ValueOf(self._mix)
	case "_tag_used_streams":
		return reflect.ValueOf(self._tag_used_streams)
	case "_set_parameter":
		return reflect.ValueOf(self._set_parameter)
	case "_get_parameter":
		return reflect.ValueOf(self._get_parameter)
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_start":
		return reflect.ValueOf(self._start)
	case "_stop":
		return reflect.ValueOf(self._stop)
	case "_is_playing":
		return reflect.ValueOf(self._is_playing)
	case "_get_loop_count":
		return reflect.ValueOf(self._get_loop_count)
	case "_get_playback_position":
		return reflect.ValueOf(self._get_playback_position)
	case "_seek":
		return reflect.ValueOf(self._seek)
	case "_mix":
		return reflect.ValueOf(self._mix)
	case "_tag_used_streams":
		return reflect.ValueOf(self._tag_used_streams)
	case "_set_parameter":
		return reflect.ValueOf(self._set_parameter)
	case "_get_parameter":
		return reflect.ValueOf(self._get_parameter)
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("AudioStreamPlayback", func(ptr gd.Object) any { return classdb.AudioStreamPlayback(ptr) })
}
