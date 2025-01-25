// Package VisualShaderNodeTransformConstant provides methods for working with VisualShaderNodeTransformConstant object instances.
package VisualShaderNodeTransformConstant

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/classdb/VisualShaderNodeConstant"
import "graphics.gd/classdb/VisualShaderNode"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Transform3D"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any

/*
A constant [Transform3D], which can be used as an input node.
*/
type Instance [1]gdclass.VisualShaderNodeTransformConstant

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVisualShaderNodeTransformConstant() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeTransformConstant

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeTransformConstant"))
	casted := Instance{*(*gdclass.VisualShaderNodeTransformConstant)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Constant() Transform3D.BasisOrigin {
	return Transform3D.BasisOrigin(class(self).GetConstant())
}

func (self Instance) SetConstant(value Transform3D.BasisOrigin) {
	class(self).SetConstant(gd.Transform3D(value))
}

//go:nosplit
func (self class) SetConstant(constant gd.Transform3D) {
	var frame = callframe.New()
	callframe.Arg(frame, constant)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTransformConstant.Bind_set_constant, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetConstant() gd.Transform3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeTransformConstant.Bind_get_constant, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeTransformConstant() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeTransformConstant() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualShaderNodeConstant() VisualShaderNodeConstant.Advanced {
	return *((*VisualShaderNodeConstant.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeConstant() VisualShaderNodeConstant.Instance {
	return *((*VisualShaderNodeConstant.Instance)(unsafe.Pointer(&self)))
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualShaderNodeConstant.Advanced(self.AsVisualShaderNodeConstant()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualShaderNodeConstant.Instance(self.AsVisualShaderNodeConstant()), name)
	}
}
func init() {
	gdclass.Register("VisualShaderNodeTransformConstant", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeTransformConstant{*(*gdclass.VisualShaderNodeTransformConstant)(unsafe.Pointer(&ptr))}
	})
}
