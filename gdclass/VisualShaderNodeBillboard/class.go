package VisualShaderNodeBillboard

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VisualShaderNode"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
The output port of this node needs to be connected to [code]Model View Matrix[/code] port of [VisualShaderNodeOutput].

*/
type Go [1]classdb.VisualShaderNodeBillboard
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VisualShaderNodeBillboard
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeBillboard"))
	return Go{classdb.VisualShaderNodeBillboard(object)}
}

func (self Go) BillboardType() classdb.VisualShaderNodeBillboardBillboardType {
		return classdb.VisualShaderNodeBillboardBillboardType(class(self).GetBillboardType())
}

func (self Go) SetBillboardType(value classdb.VisualShaderNodeBillboardBillboardType) {
	class(self).SetBillboardType(value)
}

func (self Go) KeepScale() bool {
		return bool(class(self).IsKeepScaleEnabled())
}

func (self Go) SetKeepScale(value bool) {
	class(self).SetKeepScaleEnabled(value)
}

//go:nosplit
func (self class) SetBillboardType(billboard_type classdb.VisualShaderNodeBillboardBillboardType)  {
	var frame = callframe.New()
	callframe.Arg(frame, billboard_type)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeBillboard.Bind_set_billboard_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBillboardType() classdb.VisualShaderNodeBillboardBillboardType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.VisualShaderNodeBillboardBillboardType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeBillboard.Bind_get_billboard_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetKeepScaleEnabled(enabled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeBillboard.Bind_set_keep_scale_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsKeepScaleEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeBillboard.Bind_is_keep_scale_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeBillboard() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeBillboard() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNode() VisualShaderNode.GD { return *((*VisualShaderNode.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNode() VisualShaderNode.Go { return *((*VisualShaderNode.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNode(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNode(), name)
	}
}
func init() {classdb.Register("VisualShaderNodeBillboard", func(ptr gd.Object) any { return classdb.VisualShaderNodeBillboard(ptr) })}
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
