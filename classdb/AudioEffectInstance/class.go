// Package AudioEffectInstance provides methods for working with AudioEffectInstance object instances.
package AudioEffectInstance

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
import "graphics.gd/variant/Path"

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

/*
An audio effect instance manipulates the audio it receives for a given effect. This instance is automatically created by an [AudioEffect] when it is added to a bus, and should usually not be created directly. If necessary, it can be fetched at run-time with [method AudioServer.get_bus_effect_instance].

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=AudioEffectInstance)
*/
type Instance [1]gdclass.AudioEffectInstance

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAudioEffectInstance() Instance
}
type Interface interface {
	//Called by the [AudioServer] to process this effect. When [method _process_silence] is not overridden or it returns [code]false[/code], this method is called only when the bus is active.
	//[b]Note:[/b] It is not useful to override this method in GDScript or C#. Only GDExtension can take advantage of it.
	Process(src_buffer unsafe.Pointer, dst_buffer *AudioFrame, frame_count int)
	//Override this method to customize the processing behavior of this effect instance.
	//Should return [code]true[/code] to force the [AudioServer] to always call [method _process], even if the bus has been muted or cannot otherwise be heard.
	ProcessSilence() bool
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) Process(src_buffer unsafe.Pointer, dst_buffer *AudioFrame, frame_count int) {
	return
}
func (self implementation) ProcessSilence() (_ bool) { return }

/*
Called by the [AudioServer] to process this effect. When [method _process_silence] is not overridden or it returns [code]false[/code], this method is called only when the bus is active.
[b]Note:[/b] It is not useful to override this method in GDScript or C#. Only GDExtension can take advantage of it.
*/
func (Instance) _process(impl func(ptr unsafe.Pointer, src_buffer unsafe.Pointer, dst_buffer *AudioFrame, frame_count int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var src_buffer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		var dst_buffer = gd.UnsafeGet[*AudioFrame](p_args, 1)

		var frame_count = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, src_buffer, dst_buffer, int(frame_count))
	}
}

/*
Override this method to customize the processing behavior of this effect instance.
Should return [code]true[/code] to force the [AudioServer] to always call [method _process], even if the bus has been muted or cannot otherwise be heard.
*/
func (Instance) _process_silence(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AudioEffectInstance

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioEffectInstance"))
	casted := Instance{*(*gdclass.AudioEffectInstance)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Called by the [AudioServer] to process this effect. When [method _process_silence] is not overridden or it returns [code]false[/code], this method is called only when the bus is active.
[b]Note:[/b] It is not useful to override this method in GDScript or C#. Only GDExtension can take advantage of it.
*/
func (class) _process(impl func(ptr unsafe.Pointer, src_buffer unsafe.Pointer, dst_buffer *AudioFrame, frame_count gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var src_buffer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		var dst_buffer = gd.UnsafeGet[*AudioFrame](p_args, 1)

		var frame_count = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, src_buffer, dst_buffer, frame_count)
	}
}

/*
Override this method to customize the processing behavior of this effect instance.
Should return [code]true[/code] to force the [AudioServer] to always call [method _process], even if the bus has been muted or cannot otherwise be heard.
*/
func (class) _process_silence(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (self class) AsAudioEffectInstance() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAudioEffectInstance() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_process":
		return reflect.ValueOf(self._process)
	case "_process_silence":
		return reflect.ValueOf(self._process_silence)
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_process":
		return reflect.ValueOf(self._process)
	case "_process_silence":
		return reflect.ValueOf(self._process_silence)
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("AudioEffectInstance", func(ptr gd.Object) any {
		return [1]gdclass.AudioEffectInstance{*(*gdclass.AudioEffectInstance)(unsafe.Pointer(&ptr))}
	})
}
