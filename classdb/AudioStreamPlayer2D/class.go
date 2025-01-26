// Package AudioStreamPlayer2D provides methods for working with AudioStreamPlayer2D object instances.
package AudioStreamPlayer2D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/classdb/Node2D"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Float"

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

/*
Plays audio that is attenuated with distance to the listener.
By default, audio is heard from the screen center. This can be changed by adding an [AudioListener2D] node to the scene and enabling it by calling [method AudioListener2D.make_current] on it.
See also [AudioStreamPlayer] to play a sound non-positionally.
[b]Note:[/b] Hiding an [AudioStreamPlayer2D] node does not disable its audio output. To temporarily disable an [AudioStreamPlayer2D]'s audio output, set [member volume_db] to a very low value like [code]-100[/code] (which isn't audible to human hearing).
*/
type Instance [1]gdclass.AudioStreamPlayer2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioStreamPlayer2D() Instance
}

/*
Queues the audio to play on the next physics frame, from the given position [param from_position], in seconds.
*/
func (self Instance) Play() { //gd:AudioStreamPlayer2D.play
	class(self).Play(gd.Float(0.0))
}

/*
Sets the position from which audio will be played, in seconds.
*/
func (self Instance) SeekTo(to_position Float.X) { //gd:AudioStreamPlayer2D.seek
	class(self).SeekTo(gd.Float(to_position))
}

/*
Stops the audio.
*/
func (self Instance) Stop() { //gd:AudioStreamPlayer2D.stop
	class(self).Stop()
}

/*
Returns the position in the [AudioStream].
*/
func (self Instance) GetPlaybackPosition() Float.X { //gd:AudioStreamPlayer2D.get_playback_position
	return Float.X(Float.X(class(self).GetPlaybackPosition()))
}

/*
Returns whether the [AudioStreamPlayer] can return the [AudioStreamPlayback] object or not.
*/
func (self Instance) HasStreamPlayback() bool { //gd:AudioStreamPlayer2D.has_stream_playback
	return bool(class(self).HasStreamPlayback())
}

/*
Returns the [AudioStreamPlayback] object associated with this [AudioStreamPlayer2D].
*/
func (self Instance) GetStreamPlayback() [1]gdclass.AudioStreamPlayback { //gd:AudioStreamPlayer2D.get_stream_playback
	return [1]gdclass.AudioStreamPlayback(class(self).GetStreamPlayback())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioStreamPlayer2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamPlayer2D"))
	casted := Instance{*(*gdclass.AudioStreamPlayer2D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Stream() [1]gdclass.AudioStream {
	return [1]gdclass.AudioStream(class(self).GetStream())
}

func (self Instance) SetStream(value [1]gdclass.AudioStream) {
	class(self).SetStream(value)
}

func (self Instance) VolumeDb() Float.X {
	return Float.X(Float.X(class(self).GetVolumeDb()))
}

func (self Instance) SetVolumeDb(value Float.X) {
	class(self).SetVolumeDb(gd.Float(value))
}

func (self Instance) PitchScale() Float.X {
	return Float.X(Float.X(class(self).GetPitchScale()))
}

func (self Instance) SetPitchScale(value Float.X) {
	class(self).SetPitchScale(gd.Float(value))
}

func (self Instance) Playing() bool {
	return bool(class(self).IsPlaying())
}

func (self Instance) Autoplay() bool {
	return bool(class(self).IsAutoplayEnabled())
}

func (self Instance) SetAutoplay(value bool) {
	class(self).SetAutoplay(value)
}

func (self Instance) StreamPaused() bool {
	return bool(class(self).GetStreamPaused())
}

func (self Instance) SetStreamPaused(value bool) {
	class(self).SetStreamPaused(value)
}

func (self Instance) MaxDistance() Float.X {
	return Float.X(Float.X(class(self).GetMaxDistance()))
}

func (self Instance) SetMaxDistance(value Float.X) {
	class(self).SetMaxDistance(gd.Float(value))
}

func (self Instance) Attenuation() Float.X {
	return Float.X(Float.X(class(self).GetAttenuation()))
}

func (self Instance) SetAttenuation(value Float.X) {
	class(self).SetAttenuation(gd.Float(value))
}

func (self Instance) MaxPolyphony() int {
	return int(int(class(self).GetMaxPolyphony()))
}

func (self Instance) SetMaxPolyphony(value int) {
	class(self).SetMaxPolyphony(gd.Int(value))
}

func (self Instance) PanningStrength() Float.X {
	return Float.X(Float.X(class(self).GetPanningStrength()))
}

func (self Instance) SetPanningStrength(value Float.X) {
	class(self).SetPanningStrength(gd.Float(value))
}

func (self Instance) Bus() string {
	return string(class(self).GetBus().String())
}

func (self Instance) SetBus(value string) {
	class(self).SetBus(gd.NewStringName(value))
}

func (self Instance) AreaMask() int {
	return int(int(class(self).GetAreaMask()))
}

func (self Instance) SetAreaMask(value int) {
	class(self).SetAreaMask(gd.Int(value))
}

func (self Instance) PlaybackType() gdclass.AudioServerPlaybackType {
	return gdclass.AudioServerPlaybackType(class(self).GetPlaybackType())
}

func (self Instance) SetPlaybackType(value gdclass.AudioServerPlaybackType) {
	class(self).SetPlaybackType(value)
}

//go:nosplit
func (self class) SetStream(stream [1]gdclass.AudioStream) { //gd:AudioStreamPlayer2D.set_stream
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(stream[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_set_stream, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStream() [1]gdclass.AudioStream { //gd:AudioStreamPlayer2D.get_stream
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_get_stream, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.AudioStream{gd.PointerWithOwnershipTransferredToGo[gdclass.AudioStream](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumeDb(volume_db gd.Float) { //gd:AudioStreamPlayer2D.set_volume_db
	var frame = callframe.New()
	callframe.Arg(frame, volume_db)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_set_volume_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVolumeDb() gd.Float { //gd:AudioStreamPlayer2D.get_volume_db
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_get_volume_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPitchScale(pitch_scale gd.Float) { //gd:AudioStreamPlayer2D.set_pitch_scale
	var frame = callframe.New()
	callframe.Arg(frame, pitch_scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_set_pitch_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPitchScale() gd.Float { //gd:AudioStreamPlayer2D.get_pitch_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_get_pitch_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Queues the audio to play on the next physics frame, from the given position [param from_position], in seconds.
*/
//go:nosplit
func (self class) Play(from_position gd.Float) { //gd:AudioStreamPlayer2D.play
	var frame = callframe.New()
	callframe.Arg(frame, from_position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_play, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the position from which audio will be played, in seconds.
*/
//go:nosplit
func (self class) SeekTo(to_position gd.Float) { //gd:AudioStreamPlayer2D.seek
	var frame = callframe.New()
	callframe.Arg(frame, to_position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_seek, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Stops the audio.
*/
//go:nosplit
func (self class) Stop() { //gd:AudioStreamPlayer2D.stop
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_stop, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsPlaying() bool { //gd:AudioStreamPlayer2D.is_playing
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_is_playing, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the position in the [AudioStream].
*/
//go:nosplit
func (self class) GetPlaybackPosition() gd.Float { //gd:AudioStreamPlayer2D.get_playback_position
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_get_playback_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBus(bus gd.StringName) { //gd:AudioStreamPlayer2D.set_bus
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(bus))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_set_bus, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBus() gd.StringName { //gd:AudioStreamPlayer2D.get_bus
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_get_bus, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoplay(enable bool) { //gd:AudioStreamPlayer2D.set_autoplay
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_set_autoplay, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsAutoplayEnabled() bool { //gd:AudioStreamPlayer2D.is_autoplay_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_is_autoplay_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxDistance(pixels gd.Float) { //gd:AudioStreamPlayer2D.set_max_distance
	var frame = callframe.New()
	callframe.Arg(frame, pixels)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_set_max_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxDistance() gd.Float { //gd:AudioStreamPlayer2D.get_max_distance
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_get_max_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAttenuation(curve gd.Float) { //gd:AudioStreamPlayer2D.set_attenuation
	var frame = callframe.New()
	callframe.Arg(frame, curve)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_set_attenuation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAttenuation() gd.Float { //gd:AudioStreamPlayer2D.get_attenuation
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_get_attenuation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAreaMask(mask gd.Int) { //gd:AudioStreamPlayer2D.set_area_mask
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_set_area_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAreaMask() gd.Int { //gd:AudioStreamPlayer2D.get_area_mask
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_get_area_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStreamPaused(pause bool) { //gd:AudioStreamPlayer2D.set_stream_paused
	var frame = callframe.New()
	callframe.Arg(frame, pause)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_set_stream_paused, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStreamPaused() bool { //gd:AudioStreamPlayer2D.get_stream_paused
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_get_stream_paused, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxPolyphony(max_polyphony gd.Int) { //gd:AudioStreamPlayer2D.set_max_polyphony
	var frame = callframe.New()
	callframe.Arg(frame, max_polyphony)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_set_max_polyphony, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxPolyphony() gd.Int { //gd:AudioStreamPlayer2D.get_max_polyphony
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_get_max_polyphony, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPanningStrength(panning_strength gd.Float) { //gd:AudioStreamPlayer2D.set_panning_strength
	var frame = callframe.New()
	callframe.Arg(frame, panning_strength)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_set_panning_strength, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPanningStrength() gd.Float { //gd:AudioStreamPlayer2D.get_panning_strength
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_get_panning_strength, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns whether the [AudioStreamPlayer] can return the [AudioStreamPlayback] object or not.
*/
//go:nosplit
func (self class) HasStreamPlayback() bool { //gd:AudioStreamPlayer2D.has_stream_playback
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_has_stream_playback, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [AudioStreamPlayback] object associated with this [AudioStreamPlayer2D].
*/
//go:nosplit
func (self class) GetStreamPlayback() [1]gdclass.AudioStreamPlayback { //gd:AudioStreamPlayer2D.get_stream_playback
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_get_stream_playback, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.AudioStreamPlayback{gd.PointerWithOwnershipTransferredToGo[gdclass.AudioStreamPlayback](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPlaybackType(playback_type gdclass.AudioServerPlaybackType) { //gd:AudioStreamPlayer2D.set_playback_type
	var frame = callframe.New()
	callframe.Arg(frame, playback_type)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_set_playback_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPlaybackType() gdclass.AudioServerPlaybackType { //gd:AudioStreamPlayer2D.get_playback_type
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AudioServerPlaybackType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamPlayer2D.Bind_get_playback_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnFinished(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("finished"), gd.NewCallable(cb), 0)
}

func (self class) AsAudioStreamPlayer2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioStreamPlayer2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.Advanced          { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode2D() Node2D.Instance       { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node2D.Advanced(self.AsNode2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node2D.Instance(self.AsNode2D()), name)
	}
}
func init() {
	gdclass.Register("AudioStreamPlayer2D", func(ptr gd.Object) any {
		return [1]gdclass.AudioStreamPlayer2D{*(*gdclass.AudioStreamPlayer2D)(unsafe.Pointer(&ptr))}
	})
}
