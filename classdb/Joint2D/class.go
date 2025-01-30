// Package Joint2D provides methods for working with Joint2D object instances.
package Joint2D

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/CanvasItem"
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
Abstract base class for all joints in 2D physics. 2D joints bind together two physics bodies ([member node_a] and [member node_b]) and apply a constraint.
*/
type Instance [1]gdclass.Joint2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsJoint2D() Instance
}

/*
Returns the joint's internal [RID] from the [PhysicsServer2D].
*/
func (self Instance) GetRid() RID.Joint2D { //gd:Joint2D.get_rid
	return RID.Joint2D(class(self).GetRid())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Joint2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Joint2D"))
	casted := Instance{*(*gdclass.Joint2D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) NodeA() string {
	return string(class(self).GetNodeA().String())
}

func (self Instance) SetNodeA(value string) {
	class(self).SetNodeA(Path.ToNode(String.New(value)))
}

func (self Instance) NodeB() string {
	return string(class(self).GetNodeB().String())
}

func (self Instance) SetNodeB(value string) {
	class(self).SetNodeB(Path.ToNode(String.New(value)))
}

func (self Instance) Bias() Float.X {
	return Float.X(Float.X(class(self).GetBias()))
}

func (self Instance) SetBias(value Float.X) {
	class(self).SetBias(float64(value))
}

func (self Instance) DisableCollision() bool {
	return bool(class(self).GetExcludeNodesFromCollision())
}

func (self Instance) SetDisableCollision(value bool) {
	class(self).SetExcludeNodesFromCollision(value)
}

//go:nosplit
func (self class) SetNodeA(node Path.ToNode) { //gd:Joint2D.set_node_a
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(node)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint2D.Bind_set_node_a, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetNodeA() Path.ToNode { //gd:Joint2D.get_node_a
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint2D.Bind_get_node_a, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Path.ToNode(String.Via(gd.NodePathProxy{}, pointers.Pack(pointers.New[gd.NodePath](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNodeB(node Path.ToNode) { //gd:Joint2D.set_node_b
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(node)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint2D.Bind_set_node_b, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetNodeB() Path.ToNode { //gd:Joint2D.get_node_b
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint2D.Bind_get_node_b, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Path.ToNode(String.Via(gd.NodePathProxy{}, pointers.Pack(pointers.New[gd.NodePath](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBias(bias float64) { //gd:Joint2D.set_bias
	var frame = callframe.New()
	callframe.Arg(frame, bias)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint2D.Bind_set_bias, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBias() float64 { //gd:Joint2D.get_bias
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint2D.Bind_get_bias, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetExcludeNodesFromCollision(enable bool) { //gd:Joint2D.set_exclude_nodes_from_collision
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint2D.Bind_set_exclude_nodes_from_collision, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetExcludeNodesFromCollision() bool { //gd:Joint2D.get_exclude_nodes_from_collision
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint2D.Bind_get_exclude_nodes_from_collision, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the joint's internal [RID] from the [PhysicsServer2D].
*/
//go:nosplit
func (self class) GetRid() RID.Any { //gd:Joint2D.get_rid
	var frame = callframe.New()
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint2D.Bind_get_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsJoint2D() Advanced          { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsJoint2D() Instance       { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(Node2D.Advanced(self.AsNode2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node2D.Instance(self.AsNode2D()), name)
	}
}
func init() {
	gdclass.Register("Joint2D", func(ptr gd.Object) any { return [1]gdclass.Joint2D{*(*gdclass.Joint2D)(unsafe.Pointer(&ptr))} })
}
