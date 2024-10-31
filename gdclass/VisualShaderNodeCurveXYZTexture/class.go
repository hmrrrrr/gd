package VisualShaderNodeCurveXYZTexture

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VisualShaderNodeResizableBase"
import "grow.graphics/gd/gdclass/VisualShaderNode"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
Comes with a built-in editor for texture's curves.
*/
type Instance [1]classdb.VisualShaderNodeCurveXYZTexture

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.VisualShaderNodeCurveXYZTexture

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeCurveXYZTexture"))
	return Instance{classdb.VisualShaderNodeCurveXYZTexture(object)}
}

func (self Instance) Texture() gdclass.CurveXYZTexture {
	return gdclass.CurveXYZTexture(class(self).GetTexture())
}

func (self Instance) SetTexture(value gdclass.CurveXYZTexture) {
	class(self).SetTexture(value)
}

//go:nosplit
func (self class) SetTexture(texture gdclass.CurveXYZTexture) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeCurveXYZTexture.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTexture() gdclass.CurveXYZTexture {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeCurveXYZTexture.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.CurveXYZTexture{classdb.CurveXYZTexture(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeCurveXYZTexture() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeCurveXYZTexture() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualShaderNodeResizableBase() VisualShaderNodeResizableBase.Advanced {
	return *((*VisualShaderNodeResizableBase.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeResizableBase() VisualShaderNodeResizableBase.Instance {
	return *((*VisualShaderNodeResizableBase.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualShaderNode() VisualShaderNode.Advanced {
	return *((*VisualShaderNode.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNode() VisualShaderNode.Instance {
	return *((*VisualShaderNode.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualShaderNodeResizableBase(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualShaderNodeResizableBase(), name)
	}
}
func init() {
	classdb.Register("VisualShaderNodeCurveXYZTexture", func(ptr gd.Object) any { return classdb.VisualShaderNodeCurveXYZTexture(ptr) })
}
