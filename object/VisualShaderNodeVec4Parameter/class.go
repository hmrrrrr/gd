package VisualShaderNodeVec4Parameter

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/VisualShaderNodeParameter"
import "grow.graphics/gd/object/VisualShaderNode"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Translated to [code]uniform vec4[/code] in the shader language.

*/
type Simple [1]classdb.VisualShaderNodeVec4Parameter
func (self Simple) SetDefaultValueEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDefaultValueEnabled(enabled)
}
func (self Simple) IsDefaultValueEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsDefaultValueEnabled())
}
func (self Simple) SetDefaultValue(value gd.Vector4) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDefaultValue(value)
}
func (self Simple) GetDefaultValue() gd.Vector4 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector4(Expert(self).GetDefaultValue())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VisualShaderNodeVec4Parameter
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetDefaultValueEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeVec4Parameter.Bind_set_default_value_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDefaultValueEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeVec4Parameter.Bind_is_default_value_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDefaultValue(value gd.Vector4)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeVec4Parameter.Bind_set_default_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDefaultValue() gd.Vector4 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector4](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.VisualShaderNodeVec4Parameter.Bind_get_default_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsVisualShaderNodeVec4Parameter() Expert { return self[0].AsVisualShaderNodeVec4Parameter() }


//go:nosplit
func (self Simple) AsVisualShaderNodeVec4Parameter() Simple { return self[0].AsVisualShaderNodeVec4Parameter() }


//go:nosplit
func (self class) AsVisualShaderNodeParameter() VisualShaderNodeParameter.Expert { return self[0].AsVisualShaderNodeParameter() }


//go:nosplit
func (self Simple) AsVisualShaderNodeParameter() VisualShaderNodeParameter.Simple { return self[0].AsVisualShaderNodeParameter() }


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
func init() {classdb.Register("VisualShaderNodeVec4Parameter", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
