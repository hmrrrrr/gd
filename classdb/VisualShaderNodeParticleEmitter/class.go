// Package VisualShaderNodeParticleEmitter provides methods for working with VisualShaderNodeParticleEmitter object instances.
package VisualShaderNodeParticleEmitter

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/VisualShaderNode"
import "graphics.gd/classdb/Resource"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
Particle emitter nodes can be used in "start" step of particle shaders and they define the starting position of the particles. Connect them to the Position output port.
*/
type Instance [1]gdclass.VisualShaderNodeParticleEmitter

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVisualShaderNodeParticleEmitter() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeParticleEmitter

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeParticleEmitter"))
	casted := Instance{*(*gdclass.VisualShaderNodeParticleEmitter)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Mode2d() bool {
	return bool(class(self).IsMode2d())
}

func (self Instance) SetMode2d(value bool) {
	class(self).SetMode2d(value)
}

//go:nosplit
func (self class) SetMode2d(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeParticleEmitter.Bind_set_mode_2d, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsMode2d() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeParticleEmitter.Bind_is_mode_2d, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeParticleEmitter() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeParticleEmitter() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(VisualShaderNode.Advanced(self.AsVisualShaderNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualShaderNode.Instance(self.AsVisualShaderNode()), name)
	}
}
func init() {
	gdclass.Register("VisualShaderNodeParticleEmitter", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeParticleEmitter{*(*gdclass.VisualShaderNodeParticleEmitter)(unsafe.Pointer(&ptr))}
	})
}
