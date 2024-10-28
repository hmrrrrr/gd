package AudioEffectLowPassFilter

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/AudioEffectFilter"
import "grow.graphics/gd/gdclass/AudioEffect"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Cuts frequencies higher than the [member AudioEffectFilter.cutoff_hz] and allows lower frequencies to pass.

*/
type Go [1]classdb.AudioEffectLowPassFilter
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.AudioEffectLowPassFilter
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectLowPassFilter"))
	return Go{classdb.AudioEffectLowPassFilter(object)}
}

func (self class) AsAudioEffectLowPassFilter() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioEffectLowPassFilter() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsAudioEffectFilter() AudioEffectFilter.GD { return *((*AudioEffectFilter.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioEffectFilter() AudioEffectFilter.Go { return *((*AudioEffectFilter.Go)(unsafe.Pointer(&self))) }
func (self class) AsAudioEffect() AudioEffect.GD { return *((*AudioEffect.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAudioEffect() AudioEffect.Go { return *((*AudioEffect.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAudioEffectFilter(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsAudioEffectFilter(), name)
	}
}
func init() {classdb.Register("AudioEffectLowPassFilter", func(ptr gd.Object) any { return classdb.AudioEffectLowPassFilter(ptr) })}
