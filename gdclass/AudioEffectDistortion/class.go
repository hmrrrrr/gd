package AudioEffectDistortion

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/AudioEffect"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
Different types are available: clip, tan, lo-fi (bit crushing), overdrive, or waveshape.
By distorting the waveform the frequency content changes, which will often make the sound "crunchy" or "abrasive". For games, it can simulate sound coming from some saturated device or speaker very efficiently.
*/
type Instance [1]classdb.AudioEffectDistortion

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AudioEffectDistortion

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectDistortion"))
	return Instance{classdb.AudioEffectDistortion(object)}
}

func (self Instance) Mode() classdb.AudioEffectDistortionMode {
	return classdb.AudioEffectDistortionMode(class(self).GetMode())
}

func (self Instance) SetMode(value classdb.AudioEffectDistortionMode) {
	class(self).SetMode(value)
}

func (self Instance) PreGain() float64 {
	return float64(float64(class(self).GetPreGain()))
}

func (self Instance) SetPreGain(value float64) {
	class(self).SetPreGain(gd.Float(value))
}

func (self Instance) KeepHfHz() float64 {
	return float64(float64(class(self).GetKeepHfHz()))
}

func (self Instance) SetKeepHfHz(value float64) {
	class(self).SetKeepHfHz(gd.Float(value))
}

func (self Instance) Drive() float64 {
	return float64(float64(class(self).GetDrive()))
}

func (self Instance) SetDrive(value float64) {
	class(self).SetDrive(gd.Float(value))
}

func (self Instance) PostGain() float64 {
	return float64(float64(class(self).GetPostGain()))
}

func (self Instance) SetPostGain(value float64) {
	class(self).SetPostGain(gd.Float(value))
}

//go:nosplit
func (self class) SetMode(mode classdb.AudioEffectDistortionMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDistortion.Bind_set_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMode() classdb.AudioEffectDistortionMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AudioEffectDistortionMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDistortion.Bind_get_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPreGain(pre_gain gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, pre_gain)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDistortion.Bind_set_pre_gain, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPreGain() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDistortion.Bind_get_pre_gain, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetKeepHfHz(keep_hf_hz gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, keep_hf_hz)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDistortion.Bind_set_keep_hf_hz, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetKeepHfHz() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDistortion.Bind_get_keep_hf_hz, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDrive(drive gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, drive)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDistortion.Bind_set_drive, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDrive() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDistortion.Bind_get_drive, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPostGain(post_gain gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, post_gain)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDistortion.Bind_set_post_gain, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPostGain() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectDistortion.Bind_get_post_gain, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioEffectDistortion() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioEffectDistortion() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	classdb.Register("AudioEffectDistortion", func(ptr gd.Object) any { return classdb.AudioEffectDistortion(ptr) })
}

type Mode = classdb.AudioEffectDistortionMode

const (
	/*Digital distortion effect which cuts off peaks at the top and bottom of the waveform.*/
	ModeClip Mode = 0
	ModeAtan Mode = 1
	/*Low-resolution digital distortion effect (bit depth reduction). You can use it to emulate the sound of early digital audio devices.*/
	ModeLofi Mode = 2
	/*Emulates the warm distortion produced by a field effect transistor, which is commonly used in solid-state musical instrument amplifiers. The [member drive] property has no effect in this mode.*/
	ModeOverdrive Mode = 3
	/*Waveshaper distortions are used mainly by electronic musicians to achieve an extra-abrasive sound.*/
	ModeWaveshape Mode = 4
)
