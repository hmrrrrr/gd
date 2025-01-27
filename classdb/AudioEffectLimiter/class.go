// Package AudioEffectLimiter provides methods for working with AudioEffectLimiter object instances.
package AudioEffectLimiter

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
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/classdb/AudioEffect"
import "graphics.gd/classdb/Resource"
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
var _ RID.Any
var _ String.Readable

/*
A limiter is similar to a compressor, but it's less flexible and designed to disallow sound going over a given dB threshold. Adding one in the Master bus is always recommended to reduce the effects of clipping.
Soft clipping starts to reduce the peaks a little below the threshold level and progressively increases its effect as the input level increases such that the threshold is never exceeded.
*/
type Instance [1]gdclass.AudioEffectLimiter

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioEffectLimiter() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioEffectLimiter

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectLimiter"))
	casted := Instance{*(*gdclass.AudioEffectLimiter)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) CeilingDb() Float.X {
	return Float.X(Float.X(class(self).GetCeilingDb()))
}

func (self Instance) SetCeilingDb(value Float.X) {
	class(self).SetCeilingDb(gd.Float(value))
}

func (self Instance) ThresholdDb() Float.X {
	return Float.X(Float.X(class(self).GetThresholdDb()))
}

func (self Instance) SetThresholdDb(value Float.X) {
	class(self).SetThresholdDb(gd.Float(value))
}

func (self Instance) SoftClipDb() Float.X {
	return Float.X(Float.X(class(self).GetSoftClipDb()))
}

func (self Instance) SetSoftClipDb(value Float.X) {
	class(self).SetSoftClipDb(gd.Float(value))
}

func (self Instance) SoftClipRatio() Float.X {
	return Float.X(Float.X(class(self).GetSoftClipRatio()))
}

func (self Instance) SetSoftClipRatio(value Float.X) {
	class(self).SetSoftClipRatio(gd.Float(value))
}

//go:nosplit
func (self class) SetCeilingDb(ceiling gd.Float) { //gd:AudioEffectLimiter.set_ceiling_db
	var frame = callframe.New()
	callframe.Arg(frame, ceiling)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectLimiter.Bind_set_ceiling_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCeilingDb() gd.Float { //gd:AudioEffectLimiter.get_ceiling_db
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectLimiter.Bind_get_ceiling_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetThresholdDb(threshold gd.Float) { //gd:AudioEffectLimiter.set_threshold_db
	var frame = callframe.New()
	callframe.Arg(frame, threshold)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectLimiter.Bind_set_threshold_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetThresholdDb() gd.Float { //gd:AudioEffectLimiter.get_threshold_db
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectLimiter.Bind_get_threshold_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSoftClipDb(soft_clip gd.Float) { //gd:AudioEffectLimiter.set_soft_clip_db
	var frame = callframe.New()
	callframe.Arg(frame, soft_clip)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectLimiter.Bind_set_soft_clip_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSoftClipDb() gd.Float { //gd:AudioEffectLimiter.get_soft_clip_db
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectLimiter.Bind_get_soft_clip_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSoftClipRatio(soft_clip gd.Float) { //gd:AudioEffectLimiter.set_soft_clip_ratio
	var frame = callframe.New()
	callframe.Arg(frame, soft_clip)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectLimiter.Bind_set_soft_clip_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSoftClipRatio() gd.Float { //gd:AudioEffectLimiter.get_soft_clip_ratio
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectLimiter.Bind_get_soft_clip_ratio, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioEffectLimiter() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioEffectLimiter() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("AudioEffectLimiter", func(ptr gd.Object) any {
		return [1]gdclass.AudioEffectLimiter{*(*gdclass.AudioEffectLimiter)(unsafe.Pointer(&ptr))}
	})
}
