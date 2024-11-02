package VideoStreamTheora

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/VideoStream"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
[VideoStream] resource handling the [url=https://www.theora.org/]Ogg Theora[/url] video format with [code].ogv[/code] extension. The Theora codec is decoded on the CPU.
[b]Note:[/b] While Ogg Theora videos can also have an [code].ogg[/code] extension, you will have to rename the extension to [code].ogv[/code] to use those videos within Godot.
*/
type Instance [1]classdb.VideoStreamTheora

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.VideoStreamTheora

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VideoStreamTheora"))
	return Instance{classdb.VideoStreamTheora(object)}
}

func (self class) AsVideoStreamTheora() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVideoStreamTheora() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsVideoStream() VideoStream.Advanced {
	return *((*VideoStream.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVideoStream() VideoStream.Instance {
	return *((*VideoStream.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsVideoStream(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVideoStream(), name)
	}
}
func init() {
	classdb.Register("VideoStreamTheora", func(ptr gd.Object) any { return classdb.VideoStreamTheora(ptr) })
}
