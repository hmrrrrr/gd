package AnimationNodeAdd3

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/AnimationNodeSync"
import "graphics.gd/objects/AnimationNode"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A resource to add to an [AnimationNodeBlendTree]. Blends two animations out of three additively out of three based on the amount value.
This animation node has three inputs:
- The base animation to add to
- A "-add" animation to blend with when the blend amount is negative
- A "+add" animation to blend with when the blend amount is positive
If the absolute value of the amount is greater than [code]1.0[/code], the animation connected to "in" port is blended with the amplified animation connected to "-add"/"+add" port.
*/
type Instance [1]classdb.AnimationNodeAdd3
type Any interface {
	gd.IsClass
	AsAnimationNodeAdd3() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AnimationNodeAdd3

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimationNodeAdd3"))
	return Instance{classdb.AnimationNodeAdd3(object)}
}

func (self class) AsAnimationNodeAdd3() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAnimationNodeAdd3() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsAnimationNodeSync() AnimationNodeSync.Advanced {
	return *((*AnimationNodeSync.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAnimationNodeSync() AnimationNodeSync.Instance {
	return *((*AnimationNodeSync.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsAnimationNode() AnimationNode.Advanced {
	return *((*AnimationNode.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsAnimationNode() AnimationNode.Instance {
	return *((*AnimationNode.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsAnimationNodeSync(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsAnimationNodeSync(), name)
	}
}
func init() {
	classdb.Register("AnimationNodeAdd3", func(ptr gd.Object) any { return [1]classdb.AnimationNodeAdd3{classdb.AnimationNodeAdd3(ptr)} })
}
