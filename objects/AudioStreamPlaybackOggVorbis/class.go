package AudioStreamPlaybackOggVorbis

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/AudioStreamPlaybackResampled"
import "grow.graphics/gd/objects/AudioStreamPlayback"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

type Instance [1]classdb.AudioStreamPlaybackOggVorbis

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AudioStreamPlaybackOggVorbis

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AudioStreamPlaybackOggVorbis"))
	return Instance{classdb.AudioStreamPlaybackOggVorbis(object)}
}

func (self class) AsAudioStreamPlaybackOggVorbis() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioStreamPlaybackOggVorbis() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAudioStreamPlaybackResampled() AudioStreamPlaybackResampled.Advanced {
	return *((*AudioStreamPlaybackResampled.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioStreamPlaybackResampled() AudioStreamPlaybackResampled.Instance {
	return *((*AudioStreamPlaybackResampled.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAudioStreamPlayback() AudioStreamPlayback.Advanced {
	return *((*AudioStreamPlayback.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAudioStreamPlayback() AudioStreamPlayback.Instance {
	return *((*AudioStreamPlayback.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsAudioStreamPlaybackResampled(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsAudioStreamPlaybackResampled(), name)
	}
}
func init() {
	classdb.Register("AudioStreamPlaybackOggVorbis", func(ptr gd.Object) any { return classdb.AudioStreamPlaybackOggVorbis(ptr) })
}
