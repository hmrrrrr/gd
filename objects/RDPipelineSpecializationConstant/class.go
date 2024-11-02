package RDPipelineSpecializationConstant

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
A [i]specialization constant[/i] is a way to create additional variants of shaders without actually increasing the number of shader versions that are compiled. This allows improving performance by reducing the number of shader versions and reducing [code]if[/code] branching, while still allowing shaders to be flexible for different use cases.
This object is used by [RenderingDevice].
*/
type Instance [1]classdb.RDPipelineSpecializationConstant

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.RDPipelineSpecializationConstant

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDPipelineSpecializationConstant"))
	return Instance{classdb.RDPipelineSpecializationConstant(object)}
}

func (self Instance) Value() gd.Variant {
	return gd.Variant(class(self).GetValue())
}

func (self Instance) SetValue(value gd.Variant) {
	class(self).SetValue(value)
}

func (self Instance) ConstantId() int {
	return int(int(class(self).GetConstantId()))
}

func (self Instance) SetConstantId(value int) {
	class(self).SetConstantId(gd.Int(value))
}

//go:nosplit
func (self class) SetValue(value gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(value))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineSpecializationConstant.Bind_set_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetValue() gd.Variant {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineSpecializationConstant.Bind_get_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetConstantId(constant_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, constant_id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineSpecializationConstant.Bind_set_constant_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetConstantId() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineSpecializationConstant.Bind_get_constant_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRDPipelineSpecializationConstant() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsRDPipelineSpecializationConstant() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
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
	classdb.Register("RDPipelineSpecializationConstant", func(ptr gd.Object) any { return classdb.RDPipelineSpecializationConstant(ptr) })
}
