// Package AudioStreamWAV provides methods for working with AudioStreamWAV object instances.
package AudioStreamWAV

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/AudioStream"
import "graphics.gd/classdb/Resource"
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
AudioStreamWAV stores sound samples loaded from WAV files. To play the stored sound, use an [AudioStreamPlayer] (for non-positional audio) or [AudioStreamPlayer2D]/[AudioStreamPlayer3D] (for positional audio). The sound can be looped.
This class can also be used to store dynamically-generated PCM audio data. See also [AudioStreamGenerator] for procedural audio generation.
*/
type Instance [1]gdclass.AudioStreamWAV

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioStreamWAV() Instance
}

/*
Saves the AudioStreamWAV as a WAV file to [param path]. Samples with IMA ADPCM or QOA formats can't be saved.
[b]Note:[/b] A [code].wav[/code] extension is automatically appended to [param path] if it is missing.
*/
func (self Instance) SaveToWav(path string) error { //gd:AudioStreamWAV.save_to_wav
	return error(gd.ToError(class(self).SaveToWav(String.New(path))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioStreamWAV

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamWAV"))
	casted := Instance{*(*gdclass.AudioStreamWAV)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Data() []byte {
	return []byte(class(self).GetData().Bytes())
}

func (self Instance) SetData(value []byte) {
	class(self).SetData(Packed.Bytes(Packed.New(value...)))
}

func (self Instance) Format() gdclass.AudioStreamWAVFormat {
	return gdclass.AudioStreamWAVFormat(class(self).GetFormat())
}

func (self Instance) SetFormat(value gdclass.AudioStreamWAVFormat) {
	class(self).SetFormat(value)
}

func (self Instance) LoopMode() gdclass.AudioStreamWAVLoopMode {
	return gdclass.AudioStreamWAVLoopMode(class(self).GetLoopMode())
}

func (self Instance) SetLoopMode(value gdclass.AudioStreamWAVLoopMode) {
	class(self).SetLoopMode(value)
}

func (self Instance) LoopBegin() int {
	return int(int(class(self).GetLoopBegin()))
}

func (self Instance) SetLoopBegin(value int) {
	class(self).SetLoopBegin(int64(value))
}

func (self Instance) LoopEnd() int {
	return int(int(class(self).GetLoopEnd()))
}

func (self Instance) SetLoopEnd(value int) {
	class(self).SetLoopEnd(int64(value))
}

func (self Instance) MixRate() int {
	return int(int(class(self).GetMixRate()))
}

func (self Instance) SetMixRate(value int) {
	class(self).SetMixRate(int64(value))
}

func (self Instance) Stereo() bool {
	return bool(class(self).IsStereo())
}

func (self Instance) SetStereo(value bool) {
	class(self).SetStereo(value)
}

//go:nosplit
func (self class) SetData(data Packed.Bytes) { //gd:AudioStreamWAV.set_data
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedByteArray, byte](Packed.Array[byte](data))))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamWAV.Bind_set_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetData() Packed.Bytes { //gd:AudioStreamWAV.get_data
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamWAV.Bind_get_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Bytes(Array.Through(gd.PackedProxy[gd.PackedByteArray, byte]{}, pointers.Pack(pointers.New[gd.PackedByteArray](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFormat(format gdclass.AudioStreamWAVFormat) { //gd:AudioStreamWAV.set_format
	var frame = callframe.New()
	callframe.Arg(frame, format)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamWAV.Bind_set_format, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFormat() gdclass.AudioStreamWAVFormat { //gd:AudioStreamWAV.get_format
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AudioStreamWAVFormat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamWAV.Bind_get_format, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLoopMode(loop_mode gdclass.AudioStreamWAVLoopMode) { //gd:AudioStreamWAV.set_loop_mode
	var frame = callframe.New()
	callframe.Arg(frame, loop_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamWAV.Bind_set_loop_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLoopMode() gdclass.AudioStreamWAVLoopMode { //gd:AudioStreamWAV.get_loop_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.AudioStreamWAVLoopMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamWAV.Bind_get_loop_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLoopBegin(loop_begin int64) { //gd:AudioStreamWAV.set_loop_begin
	var frame = callframe.New()
	callframe.Arg(frame, loop_begin)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamWAV.Bind_set_loop_begin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLoopBegin() int64 { //gd:AudioStreamWAV.get_loop_begin
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamWAV.Bind_get_loop_begin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLoopEnd(loop_end int64) { //gd:AudioStreamWAV.set_loop_end
	var frame = callframe.New()
	callframe.Arg(frame, loop_end)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamWAV.Bind_set_loop_end, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLoopEnd() int64 { //gd:AudioStreamWAV.get_loop_end
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamWAV.Bind_get_loop_end, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMixRate(mix_rate int64) { //gd:AudioStreamWAV.set_mix_rate
	var frame = callframe.New()
	callframe.Arg(frame, mix_rate)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamWAV.Bind_set_mix_rate, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMixRate() int64 { //gd:AudioStreamWAV.get_mix_rate
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamWAV.Bind_get_mix_rate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStereo(stereo bool) { //gd:AudioStreamWAV.set_stereo
	var frame = callframe.New()
	callframe.Arg(frame, stereo)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamWAV.Bind_set_stereo, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsStereo() bool { //gd:AudioStreamWAV.is_stereo
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamWAV.Bind_is_stereo, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Saves the AudioStreamWAV as a WAV file to [param path]. Samples with IMA ADPCM or QOA formats can't be saved.
[b]Note:[/b] A [code].wav[/code] extension is automatically appended to [param path] if it is missing.
*/
//go:nosplit
func (self class) SaveToWav(path String.Readable) Error.Code { //gd:AudioStreamWAV.save_to_wav
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioStreamWAV.Bind_save_to_wav, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsAudioStreamWAV() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioStreamWAV() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AudioStream.Advanced(self.AsAudioStream()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AudioStream.Instance(self.AsAudioStream()), name)
	}
}
func init() {
	gdclass.Register("AudioStreamWAV", func(ptr gd.Object) any {
		return [1]gdclass.AudioStreamWAV{*(*gdclass.AudioStreamWAV)(unsafe.Pointer(&ptr))}
	})
}

type Format = gdclass.AudioStreamWAVFormat //gd:AudioStreamWAV.Format

const (
	/*8-bit audio codec.*/
	Format8Bits Format = 0
	/*16-bit audio codec.*/
	Format16Bits Format = 1
	/*Audio is compressed using IMA ADPCM.*/
	FormatImaAdpcm Format = 2
	/*Audio is compressed as QOA ([url=https://qoaformat.org/]Quite OK Audio[/url]).*/
	FormatQoa Format = 3
)

type LoopMode = gdclass.AudioStreamWAVLoopMode //gd:AudioStreamWAV.LoopMode

const (
	/*Audio does not loop.*/
	LoopDisabled LoopMode = 0
	/*Audio loops the data between [member loop_begin] and [member loop_end], playing forward only.*/
	LoopForward LoopMode = 1
	/*Audio loops the data between [member loop_begin] and [member loop_end], playing back and forth.*/
	LoopPingpong LoopMode = 2
	/*Audio loops the data between [member loop_begin] and [member loop_end], playing backward only.*/
	LoopBackward LoopMode = 3
)
