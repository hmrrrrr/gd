package AudioEffectCompressor

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/AudioEffect"
import "graphics.gd/objects/Resource"
import "graphics.gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Dynamic range compressor reduces the level of the sound when the amplitude goes over a certain threshold in Decibels. One of the main uses of a compressor is to increase the dynamic range by clipping as little as possible (when sound goes over 0dB).
Compressor has many uses in the mix:
- In the Master bus to compress the whole output (although an [AudioEffectLimiter] is probably better).
- In voice channels to ensure they sound as balanced as possible.
- Sidechained. This can reduce the sound level sidechained with another audio bus for threshold detection. This technique is common in video game mixing to the level of music and SFX while voices are being heard.
- Accentuates transients by using a wider attack, making effects sound more punchy.
*/
type Instance [1]classdb.AudioEffectCompressor
type Any interface {
	gd.IsClass
	AsAudioEffectCompressor() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AudioEffectCompressor

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectCompressor"))
	return Instance{classdb.AudioEffectCompressor(object)}
}

func (self Instance) Threshold() Float.X {
	return Float.X(Float.X(class(self).GetThreshold()))
}

func (self Instance) SetThreshold(value Float.X) {
	class(self).SetThreshold(gd.Float(value))
}

func (self Instance) Ratio() Float.X {
	return Float.X(Float.X(class(self).GetRatio()))
}

func (self Instance) SetRatio(value Float.X) {
	class(self).SetRatio(gd.Float(value))
}

func (self Instance) Gain() Float.X {
	return Float.X(Float.X(class(self).GetGain()))
}

func (self Instance) SetGain(value Float.X) {
	class(self).SetGain(gd.Float(value))
}

func (self Instance) AttackUs() Float.X {
	return Float.X(Float.X(class(self).GetAttackUs()))
}

func (self Instance) SetAttackUs(value Float.X) {
	class(self).SetAttackUs(gd.Float(value))
}

func (self Instance) ReleaseMs() Float.X {
	return Float.X(Float.X(class(self).GetReleaseMs()))
}

func (self Instance) SetReleaseMs(value Float.X) {
	class(self).SetReleaseMs(gd.Float(value))
}

func (self Instance) Mix() Float.X {
	return Float.X(Float.X(class(self).GetMix()))
}

func (self Instance) SetMix(value Float.X) {
	class(self).SetMix(gd.Float(value))
}

func (self Instance) Sidechain() string {
	return string(class(self).GetSidechain().String())
}

func (self Instance) SetSidechain(value string) {
	class(self).SetSidechain(gd.NewStringName(value))
}

//go:nosplit
func (self class) SetThreshold(threshold gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, threshold)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCompressor.Bind_set_threshold, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetThreshold() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCompressor.Bind_get_threshold, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRatio(ratio gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCompressor.Bind_set_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRatio() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCompressor.Bind_get_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGain(gain gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, gain)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCompressor.Bind_set_gain, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGain() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCompressor.Bind_get_gain, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAttackUs(attack_us gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, attack_us)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCompressor.Bind_set_attack_us, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAttackUs() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCompressor.Bind_get_attack_us, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetReleaseMs(release_ms gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, release_ms)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCompressor.Bind_set_release_ms, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetReleaseMs() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCompressor.Bind_get_release_ms, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMix(mix gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, mix)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCompressor.Bind_set_mix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMix() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCompressor.Bind_get_mix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSidechain(sidechain gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(sidechain))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCompressor.Bind_set_sidechain, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSidechain() gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectCompressor.Bind_get_sidechain, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsAudioEffectCompressor() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioEffectCompressor() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsAudioEffect(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsAudioEffect(), name)
	}
}
func init() {
	classdb.Register("AudioEffectCompressor", func(ptr gd.Object) any { return classdb.AudioEffectCompressor(ptr) })
}
