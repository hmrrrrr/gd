// Package AudioEffectSpectrumAnalyzerInstance provides methods for working with AudioEffectSpectrumAnalyzerInstance object instances.
package AudioEffectSpectrumAnalyzerInstance

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
import "graphics.gd/classdb/AudioEffectInstance"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Vector2"

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

/*
The runtime part of an [AudioEffectSpectrumAnalyzer], which can be used to query the magnitude of a frequency range on its host bus.
An instance of this class can be acquired with [method AudioServer.get_bus_effect_instance].
*/
type Instance [1]gdclass.AudioEffectSpectrumAnalyzerInstance

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioEffectSpectrumAnalyzerInstance() Instance
}

/*
Returns the magnitude of the frequencies from [param from_hz] to [param to_hz] in linear energy as a Vector2. The [code]x[/code] component of the return value represents the left stereo channel, and [code]y[/code] represents the right channel.
[param mode] determines how the frequency range will be processed. See [enum MagnitudeMode].
*/
func (self Instance) GetMagnitudeForFrequencyRange(from_hz Float.X, to_hz Float.X) Vector2.XY { //gd:AudioEffectSpectrumAnalyzerInstance.get_magnitude_for_frequency_range
	return Vector2.XY(class(self).GetMagnitudeForFrequencyRange(gd.Float(from_hz), gd.Float(to_hz), 1))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioEffectSpectrumAnalyzerInstance

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectSpectrumAnalyzerInstance"))
	casted := Instance{*(*gdclass.AudioEffectSpectrumAnalyzerInstance)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Returns the magnitude of the frequencies from [param from_hz] to [param to_hz] in linear energy as a Vector2. The [code]x[/code] component of the return value represents the left stereo channel, and [code]y[/code] represents the right channel.
[param mode] determines how the frequency range will be processed. See [enum MagnitudeMode].
*/
//go:nosplit
func (self class) GetMagnitudeForFrequencyRange(from_hz gd.Float, to_hz gd.Float, mode gdclass.AudioEffectSpectrumAnalyzerInstanceMagnitudeMode) gd.Vector2 { //gd:AudioEffectSpectrumAnalyzerInstance.get_magnitude_for_frequency_range
	var frame = callframe.New()
	callframe.Arg(frame, from_hz)
	callframe.Arg(frame, to_hz)
	callframe.Arg(frame, mode)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AudioEffectSpectrumAnalyzerInstance.Bind_get_magnitude_for_frequency_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAudioEffectSpectrumAnalyzerInstance() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioEffectSpectrumAnalyzerInstance() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAudioEffectInstance() AudioEffectInstance.Advanced {
	return *((*AudioEffectInstance.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioEffectInstance() AudioEffectInstance.Instance {
	return *((*AudioEffectInstance.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(AudioEffectInstance.Advanced(self.AsAudioEffectInstance()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(AudioEffectInstance.Instance(self.AsAudioEffectInstance()), name)
	}
}
func init() {
	gdclass.Register("AudioEffectSpectrumAnalyzerInstance", func(ptr gd.Object) any {
		return [1]gdclass.AudioEffectSpectrumAnalyzerInstance{*(*gdclass.AudioEffectSpectrumAnalyzerInstance)(unsafe.Pointer(&ptr))}
	})
}

type MagnitudeMode = gdclass.AudioEffectSpectrumAnalyzerInstanceMagnitudeMode //gd:AudioEffectSpectrumAnalyzerInstance.MagnitudeMode

const (
	/*Use the average value across the frequency range as magnitude.*/
	MagnitudeAverage MagnitudeMode = 0
	/*Use the maximum value of the frequency range as magnitude.*/
	MagnitudeMax MagnitudeMode = 1
)
