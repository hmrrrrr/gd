// Package VideoStream provides methods for working with VideoStream object instances.
package VideoStream

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
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function

/*
Base resource type for all video streams. Classes that derive from [VideoStream] can all be used as resource types to play back videos in [VideoStreamPlayer].

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=VideoStream)
*/
type Instance [1]gdclass.VideoStream

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVideoStream() Instance
}
type Interface interface {
	//Called when the video starts playing, to initialize and return a subclass of [VideoStreamPlayback].
	InstantiatePlayback() [1]gdclass.VideoStreamPlayback
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) InstantiatePlayback() (_ [1]gdclass.VideoStreamPlayback) { return }

/*
Called when the video starts playing, to initialize and return a subclass of [VideoStreamPlayback].
*/
func (Instance) _instantiate_playback(impl func(ptr unsafe.Pointer) [1]gdclass.VideoStreamPlayback) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VideoStream

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VideoStream"))
	casted := Instance{*(*gdclass.VideoStream)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) File() string {
	return string(class(self).GetFile().String())
}

func (self Instance) SetFile(value string) {
	class(self).SetFile(gd.NewString(value))
}

/*
Called when the video starts playing, to initialize and return a subclass of [VideoStreamPlayback].
*/
func (class) _instantiate_playback(impl func(ptr unsafe.Pointer) [1]gdclass.VideoStreamPlayback) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

//go:nosplit
func (self class) SetFile(file gd.String) { //gd:VideoStream.set_file
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(file))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStream.Bind_set_file, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFile() gd.String { //gd:VideoStream.get_file
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStream.Bind_get_file, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsVideoStream() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVideoStream() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	case "_instantiate_playback":
		return reflect.ValueOf(self._instantiate_playback)
	default:
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_instantiate_playback":
		return reflect.ValueOf(self._instantiate_playback)
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("VideoStream", func(ptr gd.Object) any { return [1]gdclass.VideoStream{*(*gdclass.VideoStream)(unsafe.Pointer(&ptr))} })
}
