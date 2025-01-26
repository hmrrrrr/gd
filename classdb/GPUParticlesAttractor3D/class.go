// Package GPUParticlesAttractor3D provides methods for working with GPUParticlesAttractor3D object instances.
package GPUParticlesAttractor3D

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
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/classdb/VisualInstance3D"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any

/*
Particle attractors can be used to attract particles towards the attractor's origin, or to push them away from the attractor's origin.
Particle attractors work in real-time and can be moved, rotated and scaled during gameplay. Unlike collision shapes, non-uniform scaling of attractors is also supported.
Attractors can be temporarily disabled by hiding them, or by setting their [member strength] to [code]0.0[/code].
[b]Note:[/b] Particle attractors only affect [GPUParticles3D], not [CPUParticles3D].
*/
type Instance [1]gdclass.GPUParticlesAttractor3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGPUParticlesAttractor3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GPUParticlesAttractor3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GPUParticlesAttractor3D"))
	casted := Instance{*(*gdclass.GPUParticlesAttractor3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Strength() Float.X {
	return Float.X(Float.X(class(self).GetStrength()))
}

func (self Instance) SetStrength(value Float.X) {
	class(self).SetStrength(gd.Float(value))
}

func (self Instance) Attenuation() Float.X {
	return Float.X(Float.X(class(self).GetAttenuation()))
}

func (self Instance) SetAttenuation(value Float.X) {
	class(self).SetAttenuation(gd.Float(value))
}

func (self Instance) Directionality() Float.X {
	return Float.X(Float.X(class(self).GetDirectionality()))
}

func (self Instance) SetDirectionality(value Float.X) {
	class(self).SetDirectionality(gd.Float(value))
}

func (self Instance) CullMask() int {
	return int(int(class(self).GetCullMask()))
}

func (self Instance) SetCullMask(value int) {
	class(self).SetCullMask(gd.Int(value))
}

//go:nosplit
func (self class) SetCullMask(mask gd.Int) { //gd:GPUParticlesAttractor3D.set_cull_mask
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesAttractor3D.Bind_set_cull_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCullMask() gd.Int { //gd:GPUParticlesAttractor3D.get_cull_mask
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesAttractor3D.Bind_get_cull_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStrength(strength gd.Float) { //gd:GPUParticlesAttractor3D.set_strength
	var frame = callframe.New()
	callframe.Arg(frame, strength)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesAttractor3D.Bind_set_strength, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStrength() gd.Float { //gd:GPUParticlesAttractor3D.get_strength
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesAttractor3D.Bind_get_strength, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAttenuation(attenuation gd.Float) { //gd:GPUParticlesAttractor3D.set_attenuation
	var frame = callframe.New()
	callframe.Arg(frame, attenuation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesAttractor3D.Bind_set_attenuation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAttenuation() gd.Float { //gd:GPUParticlesAttractor3D.get_attenuation
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesAttractor3D.Bind_get_attenuation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDirectionality(amount gd.Float) { //gd:GPUParticlesAttractor3D.set_directionality
	var frame = callframe.New()
	callframe.Arg(frame, amount)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesAttractor3D.Bind_set_directionality, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDirectionality() gd.Float { //gd:GPUParticlesAttractor3D.get_directionality
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesAttractor3D.Bind_get_directionality, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsGPUParticlesAttractor3D() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGPUParticlesAttractor3D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualInstance3D() VisualInstance3D.Advanced {
	return *((*VisualInstance3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualInstance3D() VisualInstance3D.Instance {
	return *((*VisualInstance3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualInstance3D.Advanced(self.AsVisualInstance3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualInstance3D.Instance(self.AsVisualInstance3D()), name)
	}
}
func init() {
	gdclass.Register("GPUParticlesAttractor3D", func(ptr gd.Object) any {
		return [1]gdclass.GPUParticlesAttractor3D{*(*gdclass.GPUParticlesAttractor3D)(unsafe.Pointer(&ptr))}
	})
}
