package VisualShaderNodeVectorLen

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/VisualShaderNodeVectorBase"
import "grow.graphics/gd/object/VisualShaderNode"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Translated to [code]length(p0)[/code] in the shader language.

*/
type Simple [1]classdb.VisualShaderNodeVectorLen
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VisualShaderNodeVectorLen
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsVisualShaderNodeVectorLen() Expert { return self[0].AsVisualShaderNodeVectorLen() }


//go:nosplit
func (self Simple) AsVisualShaderNodeVectorLen() Simple { return self[0].AsVisualShaderNodeVectorLen() }


//go:nosplit
func (self class) AsVisualShaderNodeVectorBase() VisualShaderNodeVectorBase.Expert { return self[0].AsVisualShaderNodeVectorBase() }


//go:nosplit
func (self Simple) AsVisualShaderNodeVectorBase() VisualShaderNodeVectorBase.Simple { return self[0].AsVisualShaderNodeVectorBase() }


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
func init() {classdb.Register("VisualShaderNodeVectorLen", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
