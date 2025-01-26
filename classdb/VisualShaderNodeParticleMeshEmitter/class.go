// Package VisualShaderNodeParticleMeshEmitter provides methods for working with VisualShaderNodeParticleMeshEmitter object instances.
package VisualShaderNodeParticleMeshEmitter

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
import "graphics.gd/classdb/VisualShaderNodeParticleEmitter"
import "graphics.gd/classdb/VisualShaderNode"
import "graphics.gd/classdb/Resource"

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
[VisualShaderNodeParticleEmitter] that makes the particles emitted in a shape of the assigned [member mesh]. It will emit from the mesh's surfaces, either all or only the specified one.
*/
type Instance [1]gdclass.VisualShaderNodeParticleMeshEmitter

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVisualShaderNodeParticleMeshEmitter() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisualShaderNodeParticleMeshEmitter

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisualShaderNodeParticleMeshEmitter"))
	casted := Instance{*(*gdclass.VisualShaderNodeParticleMeshEmitter)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Mesh() [1]gdclass.Mesh {
	return [1]gdclass.Mesh(class(self).GetMesh())
}

func (self Instance) SetMesh(value [1]gdclass.Mesh) {
	class(self).SetMesh(value)
}

func (self Instance) UseAllSurfaces() bool {
	return bool(class(self).IsUseAllSurfaces())
}

func (self Instance) SetUseAllSurfaces(value bool) {
	class(self).SetUseAllSurfaces(value)
}

func (self Instance) SurfaceIndex() int {
	return int(int(class(self).GetSurfaceIndex()))
}

func (self Instance) SetSurfaceIndex(value int) {
	class(self).SetSurfaceIndex(gd.Int(value))
}

//go:nosplit
func (self class) SetMesh(mesh [1]gdclass.Mesh) { //gd:VisualShaderNodeParticleMeshEmitter.set_mesh
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(mesh[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeParticleMeshEmitter.Bind_set_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMesh() [1]gdclass.Mesh { //gd:VisualShaderNodeParticleMeshEmitter.get_mesh
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeParticleMeshEmitter.Bind_get_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Mesh{gd.PointerWithOwnershipTransferredToGo[gdclass.Mesh](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseAllSurfaces(enabled bool) { //gd:VisualShaderNodeParticleMeshEmitter.set_use_all_surfaces
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeParticleMeshEmitter.Bind_set_use_all_surfaces, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsUseAllSurfaces() bool { //gd:VisualShaderNodeParticleMeshEmitter.is_use_all_surfaces
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeParticleMeshEmitter.Bind_is_use_all_surfaces, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSurfaceIndex(surface_index gd.Int) { //gd:VisualShaderNodeParticleMeshEmitter.set_surface_index
	var frame = callframe.New()
	callframe.Arg(frame, surface_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeParticleMeshEmitter.Bind_set_surface_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSurfaceIndex() gd.Int { //gd:VisualShaderNodeParticleMeshEmitter.get_surface_index
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisualShaderNodeParticleMeshEmitter.Bind_get_surface_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVisualShaderNodeParticleMeshEmitter() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualShaderNodeParticleMeshEmitter() Instance {
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualShaderNodeParticleEmitter.Advanced(self.AsVisualShaderNodeParticleEmitter()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisualShaderNodeParticleEmitter.Instance(self.AsVisualShaderNodeParticleEmitter()), name)
	}
}
func init() {
	gdclass.Register("VisualShaderNodeParticleMeshEmitter", func(ptr gd.Object) any {
		return [1]gdclass.VisualShaderNodeParticleMeshEmitter{*(*gdclass.VisualShaderNodeParticleMeshEmitter)(unsafe.Pointer(&ptr))}
	})
}
