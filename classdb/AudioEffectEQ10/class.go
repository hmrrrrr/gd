// Package AudioEffectEQ10 provides methods for working with AudioEffectEQ10 object instances.
package AudioEffectEQ10

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/AudioEffectEQ"
import "graphics.gd/classdb/AudioEffect"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
Frequency bands:
Band 1: 31 Hz
Band 2: 62 Hz
Band 3: 125 Hz
Band 4: 250 Hz
Band 5: 500 Hz
Band 6: 1000 Hz
Band 7: 2000 Hz
Band 8: 4000 Hz
Band 9: 8000 Hz
Band 10: 16000 Hz
See also [AudioEffectEQ], [AudioEffectEQ6], [AudioEffectEQ21].
*/
type Instance [1]gdclass.AudioEffectEQ10

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioEffectEQ10() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioEffectEQ10

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectEQ10"))
	casted := Instance{*(*gdclass.AudioEffectEQ10)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self class) AsAudioEffectEQ10() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioEffectEQ10() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsAudioEffectEQ() AudioEffectEQ.Advanced {
	return *((*AudioEffectEQ.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioEffectEQ() AudioEffectEQ.Instance {
	return *((*AudioEffectEQ.Instance)(unsafe.Pointer(&self)))
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AudioEffectEQ.Advanced(self.AsAudioEffectEQ()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AudioEffectEQ.Instance(self.AsAudioEffectEQ()), name)
	}
}
func init() {
	gdclass.Register("AudioEffectEQ10", func(ptr gd.Object) any {
		return [1]gdclass.AudioEffectEQ10{*(*gdclass.AudioEffectEQ10)(unsafe.Pointer(&ptr))}
	})
}
