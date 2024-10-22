package AudioEffectAmplify

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/AudioEffect"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Increases or decreases the volume being routed through the audio bus.

*/
type Go [1]classdb.AudioEffectAmplify
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioEffectAmplify
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("AudioEffectAmplify"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) VolumeDb() float64 {
	gc := gd.GarbageCollector(); _ = gc
		return float64(float64(class(self).GetVolumeDb()))
}

func (self Go) SetVolumeDb(value float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetVolumeDb(gd.Float(value))
}

//go:nosplit
func (self class) SetVolumeDb(volume gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, volume)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectAmplify.Bind_set_volume_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVolumeDb() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AudioEffectAmplify.Bind_get_volume_db, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioEffectAmplify() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioEffectAmplify() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsAudioEffect() AudioEffect.GD { return *((*AudioEffect.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioEffect() AudioEffect.Go { return *((*AudioEffect.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAudioEffect(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAudioEffect(), name)
	}
}
func init() {classdb.Register("AudioEffectAmplify", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
