package VideoStream

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Base resource type for all video streams. Classes that derive from [VideoStream] can all be used as resource types to play back videos in [VideoStreamPlayer].

	// VideoStream methods that can be overridden by a [Class] that extends it.
	type VideoStream interface {
		//Called when the video starts playing, to initialize and return a subclass of [VideoStreamPlayback].
		InstantiatePlayback() objects.VideoStreamPlayback
	}
*/
type Instance [1]classdb.VideoStream

/*
Called when the video starts playing, to initialize and return a subclass of [VideoStreamPlayback].
*/
func (Instance) _instantiate_playback(impl func(ptr unsafe.Pointer) objects.VideoStreamPlayback) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
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
type class [1]classdb.VideoStream

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VideoStream"))
	return Instance{classdb.VideoStream(object)}
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
func (class) _instantiate_playback(impl func(ptr unsafe.Pointer) objects.VideoStreamPlayback) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
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
func (self class) SetFile(file gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(file))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStream.Bind_set_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFile() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStream.Bind_get_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_instantiate_playback":
		return reflect.ValueOf(self._instantiate_playback)
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_instantiate_playback":
		return reflect.ValueOf(self._instantiate_playback)
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {
	classdb.Register("VideoStream", func(ptr gd.Object) any { return classdb.VideoStream(ptr) })
}
