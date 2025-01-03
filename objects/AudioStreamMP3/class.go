package AudioStreamMP3

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/AudioStream"
import "graphics.gd/objects/Resource"
import "graphics.gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
MP3 audio stream driver. See [member data] if you want to load an MP3 file at run-time.
*/
type Instance [1]classdb.AudioStreamMP3
type Any interface {
	gd.IsClass
	AsAudioStreamMP3() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AudioStreamMP3

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamMP3"))
	return Instance{*(*classdb.AudioStreamMP3)(unsafe.Pointer(&object))}
}

func (self Instance) Data() []byte {
	return []byte(class(self).GetData().Bytes())
}

func (self Instance) SetData(value []byte) {
	class(self).SetData(gd.NewPackedByteSlice(value))
}

func (self Instance) Bpm() Float.X {
	return Float.X(Float.X(class(self).GetBpm()))
}

func (self Instance) SetBpm(value Float.X) {
	class(self).SetBpm(gd.Float(value))
}

func (self Instance) BeatCount() int {
	return int(int(class(self).GetBeatCount()))
}

func (self Instance) SetBeatCount(value int) {
	class(self).SetBeatCount(gd.Int(value))
}

func (self Instance) BarBeats() int {
	return int(int(class(self).GetBarBeats()))
}

func (self Instance) SetBarBeats(value int) {
	class(self).SetBarBeats(gd.Int(value))
}

func (self Instance) Loop() bool {
	return bool(class(self).HasLoop())
}

func (self Instance) SetLoop(value bool) {
	class(self).SetLoop(value)
}

func (self Instance) LoopOffset() Float.X {
	return Float.X(Float.X(class(self).GetLoopOffset()))
}

func (self Instance) SetLoopOffset(value Float.X) {
	class(self).SetLoopOffset(gd.Float(value))
}

//go:nosplit
func (self class) SetData(data gd.PackedByteArray) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(data))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamMP3.Bind_set_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetData() gd.PackedByteArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamMP3.Bind_get_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLoop(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamMP3.Bind_set_loop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) HasLoop() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamMP3.Bind_has_loop, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLoopOffset(seconds gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, seconds)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamMP3.Bind_set_loop_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLoopOffset() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamMP3.Bind_get_loop_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBpm(bpm gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, bpm)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamMP3.Bind_set_bpm, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBpm() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamMP3.Bind_get_bpm, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBeatCount(count gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamMP3.Bind_set_beat_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBeatCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamMP3.Bind_get_beat_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBarBeats(count gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamMP3.Bind_set_bar_beats, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBarBeats() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamMP3.Bind_get_bar_beats, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioStreamMP3() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioStreamMP3() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	classdb.Register("AudioStreamMP3", func(ptr gd.Object) any {
		return [1]classdb.AudioStreamMP3{*(*classdb.AudioStreamMP3)(unsafe.Pointer(&ptr))}
	})
}
