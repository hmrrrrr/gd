package VisualShaderNodeColorConstant

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VisualShaderNodeConstant"
import "grow.graphics/gd/gdclass/VisualShaderNode"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Has two output ports representing RGB and alpha channels of [Color].
Translated to [code]vec3 rgb[/code] and [code]float alpha[/code] in the shader language.

*/
type Go [1]classdb.VisualShaderNodeColorConstant
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VisualShaderNodeColorConstant
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeColorConstant"))
	return Go{classdb.VisualShaderNodeColorConstant(object)}
}

func (self Go) Constant() gd.Color {
		return gd.Color(class(self).GetConstant())
}

func (self Go) SetConstant(value gd.Color) {
	class(self).SetConstant(value)
}

//go:nosplit
func (self class) SetConstant(constant gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, constant)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeColorConstant.Bind_set_constant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetConstant() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeColorConstant.Bind_get_constant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeColorConstant() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeColorConstant() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNodeConstant() VisualShaderNodeConstant.GD { return *((*VisualShaderNodeConstant.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeConstant() VisualShaderNodeConstant.Go { return *((*VisualShaderNodeConstant.Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNode() VisualShaderNode.GD { return *((*VisualShaderNode.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNode() VisualShaderNode.Go { return *((*VisualShaderNode.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNodeConstant(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNodeConstant(), name)
	}
}
func init() {classdb.Register("VisualShaderNodeColorConstant", func(ptr gd.Object) any { return classdb.VisualShaderNodeColorConstant(ptr) })}
