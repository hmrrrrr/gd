package VisualShaderNodeTextureParameterTriplanar

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/VisualShaderNodeTextureParameter"
import "grow.graphics/gd/gdclass/VisualShaderNodeParameter"
import "grow.graphics/gd/gdclass/VisualShaderNode"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Performs a lookup operation on the texture provided as a uniform for the shader, with support for triplanar mapping.

*/
type Go [1]classdb.VisualShaderNodeTextureParameterTriplanar
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.VisualShaderNodeTextureParameterTriplanar
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeTextureParameterTriplanar"))
	return Go{classdb.VisualShaderNodeTextureParameterTriplanar(object)}
}

func (self class) AsVisualShaderNodeTextureParameterTriplanar() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeTextureParameterTriplanar() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNodeTextureParameter() VisualShaderNodeTextureParameter.GD { return *((*VisualShaderNodeTextureParameter.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeTextureParameter() VisualShaderNodeTextureParameter.Go { return *((*VisualShaderNodeTextureParameter.Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNodeParameter() VisualShaderNodeParameter.GD { return *((*VisualShaderNodeParameter.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNodeParameter() VisualShaderNodeParameter.Go { return *((*VisualShaderNodeParameter.Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualShaderNode() VisualShaderNode.GD { return *((*VisualShaderNode.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualShaderNode() VisualShaderNode.Go { return *((*VisualShaderNode.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNodeTextureParameter(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsVisualShaderNodeTextureParameter(), name)
	}
}
func init() {classdb.Register("VisualShaderNodeTextureParameterTriplanar", func(ptr gd.Object) any { return classdb.VisualShaderNodeTextureParameterTriplanar(ptr) })}
