// Package VisualShaderNodeParticleRingEmitter provides methods for working with VisualShaderNodeParticleRingEmitter object instances.
package VisualShaderNodeParticleRingEmitter

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/classdb/VisualShaderNodeParticleEmitter"
import "graphics.gd/classdb/VisualShaderNode"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
[VisualShaderNodeParticleEmitter] that makes the particles emitted in ring shape with the specified inner and outer radii and height.
*/
type Instance [1]gdclass.VisualShaderNodeParticleRingEmitter
type Any interface {
	gd.IsClass
	AsVisualShaderNodeParticleRingEmitter() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeParticleRingEmitter

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeParticleRingEmitter"))
	return Instance{*(*gdclass.VisualShaderNodeParticleRingEmitter)(unsafe.Pointer(&object))}
}

func (self class) AsVisualShaderNodeParticleRingEmitter() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeParticleRingEmitter() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualShaderNodeParticleEmitter() VisualShaderNodeParticleEmitter.Advanced {
	return *((*VisualShaderNodeParticleEmitter.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeParticleEmitter() VisualShaderNodeParticleEmitter.Instance {
	return *((*VisualShaderNodeParticleEmitter.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsVisualShaderNodeParticleEmitter(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsVisualShaderNodeParticleEmitter(), name)
	}
}
func init() {
	gdclass.Register("VisualShaderNodeParticleRingEmitter", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeParticleRingEmitter{*(*gdclass.VisualShaderNodeParticleRingEmitter)(unsafe.Pointer(&ptr))}
	})
}
