// Package GrooveJoint2D provides methods for working with GrooveJoint2D object instances.
package GrooveJoint2D

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Joint2D"
import "graphics.gd/classdb/Node"
import "graphics.gd/classdb/Node2D"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

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
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
A physics joint that restricts the movement of two 2D physics bodies to a fixed axis. For example, a [StaticBody2D] representing a piston base can be attached to a [RigidBody2D] representing the piston head, moving up and down.
*/
type Instance [1]gdclass.GrooveJoint2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGrooveJoint2D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GrooveJoint2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GrooveJoint2D"))
	casted := Instance{*(*gdclass.GrooveJoint2D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Length() Float.X {
	return Float.X(Float.X(class(self).GetLength()))
}

func (self Instance) SetLength(value Float.X) {
	class(self).SetLength(float64(value))
}

func (self Instance) InitialOffset() Float.X {
	return Float.X(Float.X(class(self).GetInitialOffset()))
}

func (self Instance) SetInitialOffset(value Float.X) {
	class(self).SetInitialOffset(float64(value))
}

//go:nosplit
func (self class) SetLength(length float64) { //gd:GrooveJoint2D.set_length
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GrooveJoint2D.Bind_set_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLength() float64 { //gd:GrooveJoint2D.get_length
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GrooveJoint2D.Bind_get_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInitialOffset(offset float64) { //gd:GrooveJoint2D.set_initial_offset
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GrooveJoint2D.Bind_set_initial_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetInitialOffset() float64 { //gd:GrooveJoint2D.get_initial_offset
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GrooveJoint2D.Bind_get_initial_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsGrooveJoint2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGrooveJoint2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsJoint2D() Joint2D.Advanced  { return *((*Joint2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsJoint2D() Joint2D.Instance {
	return *((*Joint2D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode2D() Node2D.Advanced    { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode2D() Node2D.Instance { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Joint2D.Advanced(self.AsJoint2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Joint2D.Instance(self.AsJoint2D()), name)
	}
}
func init() {
	gdclass.Register("GrooveJoint2D", func(ptr gd.Object) any {
		return [1]gdclass.GrooveJoint2D{*(*gdclass.GrooveJoint2D)(unsafe.Pointer(&ptr))}
	})
}
