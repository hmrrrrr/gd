package XRCamera3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Camera3D"
import "grow.graphics/gd/objects/Node3D"
import "grow.graphics/gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
This is a helper spatial node for our camera; note that, if stereoscopic rendering is applicable (VR-HMD), most of the camera properties are ignored, as the HMD information overrides them. The only properties that can be trusted are the near and far planes.
The position and orientation of this node is automatically updated by the XR Server to represent the location of the HMD if such tracking is available and can thus be used by game logic. Note that, in contrast to the XR Controller, the render thread has access to the most up-to-date tracking data of the HMD and the location of the XRCamera3D can lag a few milliseconds behind what is used for rendering as a result.
*/
type Instance [1]classdb.XRCamera3D

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.XRCamera3D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("XRCamera3D"))
	return Instance{classdb.XRCamera3D(object)}
}

func (self class) AsXRCamera3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsXRCamera3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsCamera3D() Camera3D.Advanced {
	return *((*Camera3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCamera3D() Camera3D.Instance {
	return *((*Camera3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsCamera3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsCamera3D(), name)
	}
}
func init() {
	classdb.Register("XRCamera3D", func(ptr gd.Object) any { return classdb.XRCamera3D(ptr) })
}
