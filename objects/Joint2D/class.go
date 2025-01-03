package Joint2D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Node2D"
import "graphics.gd/objects/CanvasItem"
import "graphics.gd/objects/Node"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/Float"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Abstract base class for all joints in 2D physics. 2D joints bind together two physics bodies ([member node_a] and [member node_b]) and apply a constraint.
*/
type Instance [1]classdb.Joint2D
type Any interface {
	gd.IsClass
	AsJoint2D() Instance
}

/*
Returns the joint's internal [RID] from the [PhysicsServer2D].
*/
func (self Instance) GetRid() Resource.ID {
	return Resource.ID(class(self).GetRid())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Joint2D

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Joint2D"))
	return Instance{*(*classdb.Joint2D)(unsafe.Pointer(&object))}
}

func (self Instance) NodeA() Path.String {
	return Path.String(class(self).GetNodeA().String())
}

func (self Instance) SetNodeA(value Path.String) {
	class(self).SetNodeA(gd.NewString(string(value)).NodePath())
}

func (self Instance) NodeB() Path.String {
	return Path.String(class(self).GetNodeB().String())
}

func (self Instance) SetNodeB(value Path.String) {
	class(self).SetNodeB(gd.NewString(string(value)).NodePath())
}

func (self Instance) Bias() Float.X {
	return Float.X(Float.X(class(self).GetBias()))
}

func (self Instance) SetBias(value Float.X) {
	class(self).SetBias(gd.Float(value))
}

func (self Instance) DisableCollision() bool {
	return bool(class(self).GetExcludeNodesFromCollision())
}

func (self Instance) SetDisableCollision(value bool) {
	class(self).SetExcludeNodesFromCollision(value)
}

//go:nosplit
func (self class) SetNodeA(node gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(node))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint2D.Bind_set_node_a, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetNodeA() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint2D.Bind_get_node_a, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNodeB(node gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(node))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint2D.Bind_set_node_b, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetNodeB() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint2D.Bind_get_node_b, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBias(bias gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, bias)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint2D.Bind_set_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBias() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint2D.Bind_get_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetExcludeNodesFromCollision(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint2D.Bind_set_exclude_nodes_from_collision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetExcludeNodesFromCollision() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint2D.Bind_get_exclude_nodes_from_collision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the joint's internal [RID] from the [PhysicsServer2D].
*/
//go:nosplit
func (self class) GetRid() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Joint2D.Bind_get_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
		return gd.VirtualByName(self.AsNode2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode2D(), name)
	}
}
func init() {
	classdb.Register("Joint2D", func(ptr gd.Object) any { return [1]classdb.Joint2D{*(*classdb.Joint2D)(unsafe.Pointer(&ptr))} })
}
