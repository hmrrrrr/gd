// Package AudioEffectLowShelfFilter provides methods for working with AudioEffectLowShelfFilter object instances.
package AudioEffectLowShelfFilter

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/AudioEffectFilter"
import "graphics.gd/classdb/AudioEffect"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Reduces all frequencies below the [member AudioEffectFilter.cutoff_hz].
*/
type Instance [1]gdclass.AudioEffectLowShelfFilter
type Any interface {
	gd.IsClass
	AsAudioEffectLowShelfFilter() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioEffectLowShelfFilter

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectLowShelfFilter"))
	return Instance{*(*gdclass.AudioEffectLowShelfFilter)(unsafe.Pointer(&object))}
}

func (self class) AsAudioEffectLowShelfFilter() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioEffectLowShelfFilter() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAudioEffectFilter() AudioEffectFilter.Advanced {
	return *((*AudioEffectFilter.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioEffectFilter() AudioEffectFilter.Instance {
	return *((*AudioEffectFilter.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(self.AsAudioEffectFilter(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsAudioEffectFilter(), name)
	}
}
func init() {
	gdclass.Register("AudioEffectLowShelfFilter", func(ptr gd.Object) any {
		return [1]gdclass.AudioEffectLowShelfFilter{*(*gdclass.AudioEffectLowShelfFilter)(unsafe.Pointer(&ptr))}
	})
}
