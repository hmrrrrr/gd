package AudioStreamSynchronized

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/AudioStream"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
This is a stream that can be fitted with sub-streams, which will be played in-sync. The streams being at exactly the same time when play is pressed, and will end when the last of them ends. If one of the sub-streams loops, then playback will continue.
*/
type Instance [1]classdb.AudioStreamSynchronized

/*
Set one of the synchronized streams, by index.
*/
func (self Instance) SetSyncStream(stream_index int, audio_stream gdclass.AudioStream) {
	class(self).SetSyncStream(gd.Int(stream_index), audio_stream)
}

/*
Get one of the synchronized streams, by index.
*/
func (self Instance) GetSyncStream(stream_index int) gdclass.AudioStream {
	return gdclass.AudioStream(class(self).GetSyncStream(gd.Int(stream_index)))
}

/*
Set the volume of one of the synchronized streams, by index.
*/
func (self Instance) SetSyncStreamVolume(stream_index int, volume_db float64) {
	class(self).SetSyncStreamVolume(gd.Int(stream_index), gd.Float(volume_db))
}

/*
Get the volume of one of the synchronized streams, by index.
*/
func (self Instance) GetSyncStreamVolume(stream_index int) float64 {
	return float64(float64(class(self).GetSyncStreamVolume(gd.Int(stream_index))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AudioStreamSynchronized

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamSynchronized"))
	return Instance{classdb.AudioStreamSynchronized(object)}
}

func (self Instance) StreamCount() int {
	return int(int(class(self).GetStreamCount()))
}

func (self Instance) SetStreamCount(value int) {
	class(self).SetStreamCount(gd.Int(value))
}

//go:nosplit
func (self class) SetStreamCount(stream_count gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, stream_count)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamSynchronized.Bind_set_stream_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetStreamCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamSynchronized.Bind_get_stream_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set one of the synchronized streams, by index.
*/
//go:nosplit
func (self class) SetSyncStream(stream_index gd.Int, audio_stream gdclass.AudioStream) {
	var frame = callframe.New()
	callframe.Arg(frame, stream_index)
	callframe.Arg(frame, pointers.Get(audio_stream[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamSynchronized.Bind_set_sync_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Get one of the synchronized streams, by index.
*/
//go:nosplit
func (self class) GetSyncStream(stream_index gd.Int) gdclass.AudioStream {
	var frame = callframe.New()
	callframe.Arg(frame, stream_index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamSynchronized.Bind_get_sync_stream, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.AudioStream{classdb.AudioStream(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Set the volume of one of the synchronized streams, by index.
*/
//go:nosplit
func (self class) SetSyncStreamVolume(stream_index gd.Int, volume_db gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, stream_index)
	callframe.Arg(frame, volume_db)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamSynchronized.Bind_set_sync_stream_volume, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Get the volume of one of the synchronized streams, by index.
*/
//go:nosplit
func (self class) GetSyncStreamVolume(stream_index gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, stream_index)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamSynchronized.Bind_get_sync_stream_volume, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioStreamSynchronized() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioStreamSynchronized() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAudioStream() AudioStream.Advanced {
	return *((*AudioStream.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioStream() AudioStream.Instance {
	return *((*AudioStream.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsAudioStream(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsAudioStream(), name)
	}
}
func init() {
	classdb.Register("AudioStreamSynchronized", func(ptr gd.Object) any { return classdb.AudioStreamSynchronized(ptr) })
}
