// Package XROrigin3D provides methods for working with XROrigin3D object instances.
package XROrigin3D

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
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Float"

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
This is a special node within the AR/VR system that maps the physical location of the center of our tracking space to the virtual location within our game world.
Multiple origin points can be added to the scene tree, but only one can used at a time. All the [XRCamera3D], [XRController3D], and [XRAnchor3D] nodes should be direct children of this node for spatial tracking to work correctly.
It is the position of this node that you update when your character needs to move through your game world while we're not moving in the real world. Movement in the real world is always in relation to this origin point.
For example, if your character is driving a car, the [XROrigin3D] node should be a child node of this car. Or, if you're implementing a teleport system to move your character, you should change the position of this node.
*/
type Instance [1]gdclass.XROrigin3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsXROrigin3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.XROrigin3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("XROrigin3D"))
	casted := Instance{*(*gdclass.XROrigin3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) WorldScale() Float.X {
	return Float.X(Float.X(class(self).GetWorldScale()))
}

func (self Instance) SetWorldScale(value Float.X) {
	class(self).SetWorldScale(gd.Float(value))
}

func (self Instance) Current() bool {
	return bool(class(self).IsCurrent())
}

func (self Instance) SetCurrent(value bool) {
	class(self).SetCurrent(value)
}

//go:nosplit
func (self class) SetWorldScale(world_scale gd.Float) { //gd:XROrigin3D.set_world_scale
	var frame = callframe.New()
	callframe.Arg(frame, world_scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XROrigin3D.Bind_set_world_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetWorldScale() gd.Float { //gd:XROrigin3D.get_world_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XROrigin3D.Bind_get_world_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCurrent(enabled bool) { //gd:XROrigin3D.set_current
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XROrigin3D.Bind_set_current, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsCurrent() bool { //gd:XROrigin3D.is_current
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XROrigin3D.Bind_is_current, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsXROrigin3D() Advanced       { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsXROrigin3D() Instance    { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Advanced(self.AsNode3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Instance(self.AsNode3D()), name)
	}
}
func init() {
	gdclass.Register("XROrigin3D", func(ptr gd.Object) any { return [1]gdclass.XROrigin3D{*(*gdclass.XROrigin3D)(unsafe.Pointer(&ptr))} })
}
