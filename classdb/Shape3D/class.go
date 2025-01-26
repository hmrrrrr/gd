// Package Shape3D provides methods for working with Shape3D object instances.
package Shape3D

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
import "graphics.gd/variant/RID"
import "graphics.gd/classdb/Resource"
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
var _ RID.Any

/*
Abstract base class for all 3D shapes, intended for use in physics.
[b]Performance:[/b] Primitive shapes, especially [SphereShape3D], are fast to check collisions against. [ConvexPolygonShape3D] and [HeightMapShape3D] are slower, and [ConcavePolygonShape3D] is the slowest.
*/
type Instance [1]gdclass.Shape3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsShape3D() Instance
}

/*
Returns the [ArrayMesh] used to draw the debug collision for this [Shape3D].
*/
func (self Instance) GetDebugMesh() [1]gdclass.ArrayMesh { //gd:Shape3D.get_debug_mesh
	return [1]gdclass.ArrayMesh(class(self).GetDebugMesh())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Shape3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Shape3D"))
	casted := Instance{*(*gdclass.Shape3D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) CustomSolverBias() Float.X {
	return Float.X(Float.X(class(self).GetCustomSolverBias()))
}

func (self Instance) SetCustomSolverBias(value Float.X) {
	class(self).SetCustomSolverBias(gd.Float(value))
}

func (self Instance) Margin() Float.X {
	return Float.X(Float.X(class(self).GetMargin()))
}

func (self Instance) SetMargin(value Float.X) {
	class(self).SetMargin(gd.Float(value))
}

//go:nosplit
func (self class) SetCustomSolverBias(bias gd.Float) { //gd:Shape3D.set_custom_solver_bias
	var frame = callframe.New()
	callframe.Arg(frame, bias)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shape3D.Bind_set_custom_solver_bias, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCustomSolverBias() gd.Float { //gd:Shape3D.get_custom_solver_bias
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shape3D.Bind_get_custom_solver_bias, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMargin(margin gd.Float) { //gd:Shape3D.set_margin
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shape3D.Bind_set_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMargin() gd.Float { //gd:Shape3D.get_margin
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shape3D.Bind_get_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [ArrayMesh] used to draw the debug collision for this [Shape3D].
*/
//go:nosplit
func (self class) GetDebugMesh() [1]gdclass.ArrayMesh { //gd:Shape3D.get_debug_mesh
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Shape3D.Bind_get_debug_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.ArrayMesh{gd.PointerWithOwnershipTransferredToGo[gdclass.ArrayMesh](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsShape3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsShape3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("Shape3D", func(ptr gd.Object) any { return [1]gdclass.Shape3D{*(*gdclass.Shape3D)(unsafe.Pointer(&ptr))} })
}
