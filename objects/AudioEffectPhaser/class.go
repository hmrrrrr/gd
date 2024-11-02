package AudioEffectPhaser

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/AudioEffect"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Combines phase-shifted signals with the original signal. The movement of the phase-shifted signals is controlled using a low-frequency oscillator.
*/
type Instance [1]classdb.AudioEffectPhaser

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AudioEffectPhaser

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectPhaser"))
	return Instance{classdb.AudioEffectPhaser(object)}
}

func (self Instance) RangeMinHz() float64 {
	return float64(float64(class(self).GetRangeMinHz()))
}

func (self Instance) SetRangeMinHz(value float64) {
	class(self).SetRangeMinHz(gd.Float(value))
}

func (self Instance) RangeMaxHz() float64 {
	return float64(float64(class(self).GetRangeMaxHz()))
}

func (self Instance) SetRangeMaxHz(value float64) {
	class(self).SetRangeMaxHz(gd.Float(value))
}

func (self Instance) RateHz() float64 {
	return float64(float64(class(self).GetRateHz()))
}

func (self Instance) SetRateHz(value float64) {
	class(self).SetRateHz(gd.Float(value))
}

func (self Instance) Feedback() float64 {
	return float64(float64(class(self).GetFeedback()))
}

func (self Instance) SetFeedback(value float64) {
	class(self).SetFeedback(gd.Float(value))
}

func (self Instance) Depth() float64 {
	return float64(float64(class(self).GetDepth()))
}

func (self Instance) SetDepth(value float64) {
	class(self).SetDepth(gd.Float(value))
}

//go:nosplit
func (self class) SetRangeMinHz(hz gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, hz)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectPhaser.Bind_set_range_min_hz, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRangeMinHz() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectPhaser.Bind_get_range_min_hz, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRangeMaxHz(hz gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, hz)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectPhaser.Bind_set_range_max_hz, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRangeMaxHz() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectPhaser.Bind_get_range_max_hz, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRateHz(hz gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, hz)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectPhaser.Bind_set_rate_hz, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRateHz() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectPhaser.Bind_get_rate_hz, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFeedback(fbk gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, fbk)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectPhaser.Bind_set_feedback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFeedback() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectPhaser.Bind_get_feedback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDepth(depth gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, depth)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectPhaser.Bind_set_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDepth() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectPhaser.Bind_get_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioEffectPhaser() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioEffectPhaser() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	classdb.Register("AudioEffectPhaser", func(ptr gd.Object) any { return classdb.AudioEffectPhaser(ptr) })
}
