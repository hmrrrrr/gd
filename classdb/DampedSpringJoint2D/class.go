// Package DampedSpringJoint2D provides methods for working with DampedSpringJoint2D object instances.
package DampedSpringJoint2D

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
A physics joint that connects two 2D physics bodies with a spring-like force. This resembles a spring that always wants to stretch to a given length.
*/
type Instance [1]gdclass.DampedSpringJoint2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsDampedSpringJoint2D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.DampedSpringJoint2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("DampedSpringJoint2D"))
	casted := Instance{*(*gdclass.DampedSpringJoint2D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Length() Float.X {
	return Float.X(Float.X(class(self).GetLength()))
}

func (self Instance) SetLength(value Float.X) {
	class(self).SetLength(float64(value))
}

func (self Instance) RestLength() Float.X {
	return Float.X(Float.X(class(self).GetRestLength()))
}

func (self Instance) SetRestLength(value Float.X) {
	class(self).SetRestLength(float64(value))
}

func (self Instance) Stiffness() Float.X {
	return Float.X(Float.X(class(self).GetStiffness()))
}

func (self Instance) SetStiffness(value Float.X) {
	class(self).SetStiffness(float64(value))
}

func (self Instance) Damping() Float.X {
	return Float.X(Float.X(class(self).GetDamping()))
}

func (self Instance) SetDamping(value Float.X) {
	class(self).SetDamping(float64(value))
}

//go:nosplit
func (self class) SetLength(length float64) { //gd:DampedSpringJoint2D.set_length
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DampedSpringJoint2D.Bind_set_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLength() float64 { //gd:DampedSpringJoint2D.get_length
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DampedSpringJoint2D.Bind_get_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRestLength(rest_length float64) { //gd:DampedSpringJoint2D.set_rest_length
	var frame = callframe.New()
	callframe.Arg(frame, rest_length)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DampedSpringJoint2D.Bind_set_rest_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRestLength() float64 { //gd:DampedSpringJoint2D.get_rest_length
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DampedSpringJoint2D.Bind_get_rest_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStiffness(stiffness float64) { //gd:DampedSpringJoint2D.set_stiffness
	var frame = callframe.New()
	callframe.Arg(frame, stiffness)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DampedSpringJoint2D.Bind_set_stiffness, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStiffness() float64 { //gd:DampedSpringJoint2D.get_stiffness
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DampedSpringJoint2D.Bind_get_stiffness, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDamping(damping float64) { //gd:DampedSpringJoint2D.set_damping
	var frame = callframe.New()
	callframe.Arg(frame, damping)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DampedSpringJoint2D.Bind_set_damping, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDamping() float64 { //gd:DampedSpringJoint2D.get_damping
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DampedSpringJoint2D.Bind_get_damping, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsDampedSpringJoint2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsDampedSpringJoint2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsJoint2D() Joint2D.Advanced        { return *((*Joint2D.Advanced)(unsafe.Pointer(&self))) }
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
	gdclass.Register("DampedSpringJoint2D", func(ptr gd.Object) any {
		return [1]gdclass.DampedSpringJoint2D{*(*gdclass.DampedSpringJoint2D)(unsafe.Pointer(&ptr))}
	})
}
