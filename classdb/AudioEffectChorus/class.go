// Package AudioEffectChorus provides methods for working with AudioEffectChorus object instances.
package AudioEffectChorus

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/AudioEffect"
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
Adds a chorus audio effect. The effect applies a filter with voices to duplicate the audio source and manipulate it through the filter.
*/
type Instance [1]gdclass.AudioEffectChorus

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioEffectChorus() Instance
}

func (self Instance) SetVoiceDelayMs(voice_idx int, delay_ms Float.X) { //gd:AudioEffectChorus.set_voice_delay_ms
	class(self).SetVoiceDelayMs(int64(voice_idx), float64(delay_ms))
}
func (self Instance) GetVoiceDelayMs(voice_idx int) Float.X { //gd:AudioEffectChorus.get_voice_delay_ms
	return Float.X(Float.X(class(self).GetVoiceDelayMs(int64(voice_idx))))
}
func (self Instance) SetVoiceRateHz(voice_idx int, rate_hz Float.X) { //gd:AudioEffectChorus.set_voice_rate_hz
	class(self).SetVoiceRateHz(int64(voice_idx), float64(rate_hz))
}
func (self Instance) GetVoiceRateHz(voice_idx int) Float.X { //gd:AudioEffectChorus.get_voice_rate_hz
	return Float.X(Float.X(class(self).GetVoiceRateHz(int64(voice_idx))))
}
func (self Instance) SetVoiceDepthMs(voice_idx int, depth_ms Float.X) { //gd:AudioEffectChorus.set_voice_depth_ms
	class(self).SetVoiceDepthMs(int64(voice_idx), float64(depth_ms))
}
func (self Instance) GetVoiceDepthMs(voice_idx int) Float.X { //gd:AudioEffectChorus.get_voice_depth_ms
	return Float.X(Float.X(class(self).GetVoiceDepthMs(int64(voice_idx))))
}
func (self Instance) SetVoiceLevelDb(voice_idx int, level_db Float.X) { //gd:AudioEffectChorus.set_voice_level_db
	class(self).SetVoiceLevelDb(int64(voice_idx), float64(level_db))
}
func (self Instance) GetVoiceLevelDb(voice_idx int) Float.X { //gd:AudioEffectChorus.get_voice_level_db
	return Float.X(Float.X(class(self).GetVoiceLevelDb(int64(voice_idx))))
}
func (self Instance) SetVoiceCutoffHz(voice_idx int, cutoff_hz Float.X) { //gd:AudioEffectChorus.set_voice_cutoff_hz
	class(self).SetVoiceCutoffHz(int64(voice_idx), float64(cutoff_hz))
}
func (self Instance) GetVoiceCutoffHz(voice_idx int) Float.X { //gd:AudioEffectChorus.get_voice_cutoff_hz
	return Float.X(Float.X(class(self).GetVoiceCutoffHz(int64(voice_idx))))
}
func (self Instance) SetVoicePan(voice_idx int, pan Float.X) { //gd:AudioEffectChorus.set_voice_pan
	class(self).SetVoicePan(int64(voice_idx), float64(pan))
}
func (self Instance) GetVoicePan(voice_idx int) Float.X { //gd:AudioEffectChorus.get_voice_pan
	return Float.X(Float.X(class(self).GetVoicePan(int64(voice_idx))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioEffectChorus

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectChorus"))
	casted := Instance{*(*gdclass.AudioEffectChorus)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) VoiceCount() int {
	return int(int(class(self).GetVoiceCount()))
}

func (self Instance) SetVoiceCount(value int) {
	class(self).SetVoiceCount(int64(value))
}

func (self Instance) Dry() Float.X {
	return Float.X(Float.X(class(self).GetDry()))
}

func (self Instance) SetDry(value Float.X) {
	class(self).SetDry(float64(value))
}

func (self Instance) Wet() Float.X {
	return Float.X(Float.X(class(self).GetWet()))
}

func (self Instance) SetWet(value Float.X) {
	class(self).SetWet(float64(value))
}

//go:nosplit
func (self class) SetVoiceCount(voices int64) { //gd:AudioEffectChorus.set_voice_count
	var frame = callframe.New()
	callframe.Arg(frame, voices)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_set_voice_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVoiceCount() int64 { //gd:AudioEffectChorus.get_voice_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_get_voice_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVoiceDelayMs(voice_idx int64, delay_ms float64) { //gd:AudioEffectChorus.set_voice_delay_ms
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	callframe.Arg(frame, delay_ms)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_set_voice_delay_ms, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVoiceDelayMs(voice_idx int64) float64 { //gd:AudioEffectChorus.get_voice_delay_ms
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_get_voice_delay_ms, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVoiceRateHz(voice_idx int64, rate_hz float64) { //gd:AudioEffectChorus.set_voice_rate_hz
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	callframe.Arg(frame, rate_hz)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_set_voice_rate_hz, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVoiceRateHz(voice_idx int64) float64 { //gd:AudioEffectChorus.get_voice_rate_hz
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_get_voice_rate_hz, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVoiceDepthMs(voice_idx int64, depth_ms float64) { //gd:AudioEffectChorus.set_voice_depth_ms
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	callframe.Arg(frame, depth_ms)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_set_voice_depth_ms, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVoiceDepthMs(voice_idx int64) float64 { //gd:AudioEffectChorus.get_voice_depth_ms
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_get_voice_depth_ms, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVoiceLevelDb(voice_idx int64, level_db float64) { //gd:AudioEffectChorus.set_voice_level_db
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	callframe.Arg(frame, level_db)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_set_voice_level_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVoiceLevelDb(voice_idx int64) float64 { //gd:AudioEffectChorus.get_voice_level_db
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_get_voice_level_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVoiceCutoffHz(voice_idx int64, cutoff_hz float64) { //gd:AudioEffectChorus.set_voice_cutoff_hz
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	callframe.Arg(frame, cutoff_hz)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_set_voice_cutoff_hz, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVoiceCutoffHz(voice_idx int64) float64 { //gd:AudioEffectChorus.get_voice_cutoff_hz
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_get_voice_cutoff_hz, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVoicePan(voice_idx int64, pan float64) { //gd:AudioEffectChorus.set_voice_pan
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	callframe.Arg(frame, pan)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_set_voice_pan, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVoicePan(voice_idx int64) float64 { //gd:AudioEffectChorus.get_voice_pan
	var frame = callframe.New()
	callframe.Arg(frame, voice_idx)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_get_voice_pan, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWet(amount float64) { //gd:AudioEffectChorus.set_wet
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_set_wet, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetWet() float64 { //gd:AudioEffectChorus.get_wet
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_get_wet, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDry(amount float64) { //gd:AudioEffectChorus.set_dry
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_set_dry, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDry() float64 { //gd:AudioEffectChorus.get_dry
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectChorus.Bind_get_dry, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioEffectChorus() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioEffectChorus() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsAudioEffect() AudioEffect.Advanced {
	return *((*AudioEffect.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioEffect() AudioEffect.Instance {
	return *((*AudioEffect.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(AudioEffect.Advanced(self.AsAudioEffect()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AudioEffect.Instance(self.AsAudioEffect()), name)
	}
}
func init() {
	gdclass.Register("AudioEffectChorus", func(ptr gd.Object) any {
		return [1]gdclass.AudioEffectChorus{*(*gdclass.AudioEffectChorus)(unsafe.Pointer(&ptr))}
	})
}
