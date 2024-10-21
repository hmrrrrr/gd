package VisualShaderNodeBillboard

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/VisualShaderNode"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The output port of this node needs to be connected to [code]Model View Matrix[/code] port of [VisualShaderNodeOutput].

*/
type Simple [1]classdb.VisualShaderNodeBillboard
func (self Simple) SetBillboardType(billboard_type classdb.VisualShaderNodeBillboardBillboardType) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBillboardType(billboard_type)
}
func (self Simple) GetBillboardType() classdb.VisualShaderNodeBillboardBillboardType {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.VisualShaderNodeBillboardBillboardType(Expert(self).GetBillboardType())
}
func (self Simple) SetKeepScaleEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetKeepScaleEnabled(enabled)
}
func (self Simple) IsKeepScaleEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsKeepScaleEnabled())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VisualShaderNodeBillboard
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetBillboardType(billboard_type classdb.VisualShaderNodeBillboardBillboardType)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, billboard_type)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeBillboard.Bind_set_billboard_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBillboardType() classdb.VisualShaderNodeBillboardBillboardType {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeBillboardBillboardType](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeBillboard.Bind_get_billboard_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetKeepScaleEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeBillboard.Bind_set_keep_scale_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsKeepScaleEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeBillboard.Bind_is_keep_scale_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsVisualShaderNodeBillboard() Expert { return self[0].AsVisualShaderNodeBillboard() }


//go:nosplit
func (self Simple) AsVisualShaderNodeBillboard() Simple { return self[0].AsVisualShaderNodeBillboard() }


//go:nosplit
func (self class) AsVisualShaderNode() VisualShaderNode.Expert { return self[0].AsVisualShaderNode() }


//go:nosplit
func (self Simple) AsVisualShaderNode() VisualShaderNode.Simple { return self[0].AsVisualShaderNode() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("VisualShaderNodeBillboard", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type BillboardType = classdb.VisualShaderNodeBillboardBillboardType

const (
/*Billboarding is disabled and the node does nothing.*/
	BillboardTypeDisabled BillboardType = 0
/*A standard billboarding algorithm is enabled.*/
	BillboardTypeEnabled BillboardType = 1
/*A billboarding algorithm to rotate around Y-axis is enabled.*/
	BillboardTypeFixedY BillboardType = 2
/*A billboarding algorithm designed to use on particles is enabled.*/
	BillboardTypeParticles BillboardType = 3
/*Represents the size of the [enum BillboardType] enum.*/
	BillboardTypeMax BillboardType = 4
)
