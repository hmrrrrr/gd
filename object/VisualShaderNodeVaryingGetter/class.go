package VisualShaderNodeVaryingGetter

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/VisualShaderNodeVarying"
import "grow.graphics/gd/object/VisualShaderNode"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Outputs a value of a varying defined in the shader. You need to first create a varying that can be used in the given function, e.g. varying getter in Fragment shader requires a varying with mode set to [constant VisualShader.VARYING_MODE_VERTEX_TO_FRAG_LIGHT].

*/
type Simple [1]classdb.VisualShaderNodeVaryingGetter
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.VisualShaderNodeVaryingGetter
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsVisualShaderNodeVaryingGetter() Expert { return self[0].AsVisualShaderNodeVaryingGetter() }


//go:nosplit
func (self Simple) AsVisualShaderNodeVaryingGetter() Simple { return self[0].AsVisualShaderNodeVaryingGetter() }


//go:nosplit
func (self class) AsVisualShaderNodeVarying() VisualShaderNodeVarying.Expert { return self[0].AsVisualShaderNodeVarying() }


//go:nosplit
func (self Simple) AsVisualShaderNodeVarying() VisualShaderNodeVarying.Simple { return self[0].AsVisualShaderNodeVarying() }


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
func init() {classdb.Register("VisualShaderNodeVaryingGetter", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
