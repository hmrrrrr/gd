package CameraFeed

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A camera feed gives you access to a single physical camera attached to your device. When enabled, Godot will start capturing frames from the camera which can then be used. See also [CameraServer].
[b]Note:[/b] Many cameras will return YCbCr images which are split into two textures and need to be combined in a shader. Godot does this automatically for you if you set the environment to show the camera image in the background.
*/
type Instance [1]classdb.CameraFeed

/*
Returns the unique ID for this feed.
*/
func (self Instance) GetId() int {
	return int(int(class(self).GetId()))
}

/*
Returns the camera's name.
*/
func (self Instance) GetName() string {
	return string(class(self).GetName().String())
}

/*
Returns the position of camera on the device.
*/
func (self Instance) GetPosition() classdb.CameraFeedFeedPosition {
	return classdb.CameraFeedFeedPosition(class(self).GetPosition())
}

/*
Returns feed image data type.
*/
func (self Instance) GetDatatype() classdb.CameraFeedFeedDataType {
	return classdb.CameraFeedFeedDataType(class(self).GetDatatype())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.CameraFeed

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CameraFeed"))
	return Instance{classdb.CameraFeed(object)}
}

func (self Instance) FeedIsActive() bool {
	return bool(class(self).IsActive())
}

func (self Instance) SetFeedIsActive(value bool) {
	class(self).SetActive(value)
}

func (self Instance) FeedTransform() gd.Transform2D {
	return gd.Transform2D(class(self).GetTransform())
}

func (self Instance) SetFeedTransform(value gd.Transform2D) {
	class(self).SetTransform(value)
}

/*
Returns the unique ID for this feed.
*/
//go:nosplit
func (self class) GetId() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraFeed.Bind_get_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) IsActive() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraFeed.Bind_is_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetActive(active bool) {
	var frame = callframe.New()
	callframe.Arg(frame, active)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraFeed.Bind_set_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the camera's name.
*/
//go:nosplit
func (self class) GetName() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraFeed.Bind_get_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the position of camera on the device.
*/
//go:nosplit
func (self class) GetPosition() classdb.CameraFeedFeedPosition {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CameraFeedFeedPosition](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraFeed.Bind_get_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetTransform() gd.Transform2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraFeed.Bind_get_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTransform(transform gd.Transform2D) {
	var frame = callframe.New()
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraFeed.Bind_set_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns feed image data type.
*/
//go:nosplit
func (self class) GetDatatype() classdb.CameraFeedFeedDataType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CameraFeedFeedDataType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CameraFeed.Bind_get_datatype, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCameraFeed() Advanced         { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCameraFeed() Instance      { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("CameraFeed", func(ptr gd.Object) any { return classdb.CameraFeed(ptr) })
}

type FeedDataType = classdb.CameraFeedFeedDataType

const (
	/*No image set for the feed.*/
	FeedNoimage FeedDataType = 0
	/*Feed supplies RGB images.*/
	FeedRgb FeedDataType = 1
	/*Feed supplies YCbCr images that need to be converted to RGB.*/
	FeedYcbcr FeedDataType = 2
	/*Feed supplies separate Y and CbCr images that need to be combined and converted to RGB.*/
	FeedYcbcrSep FeedDataType = 3
)

type FeedPosition = classdb.CameraFeedFeedPosition

const (
	/*Unspecified position.*/
	FeedUnspecified FeedPosition = 0
	/*Camera is mounted at the front of the device.*/
	FeedFront FeedPosition = 1
	/*Camera is mounted at the back of the device.*/
	FeedBack FeedPosition = 2
)
