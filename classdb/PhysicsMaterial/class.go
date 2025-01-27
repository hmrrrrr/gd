// Package PhysicsMaterial provides methods for working with PhysicsMaterial object instances.
package PhysicsMaterial

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
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
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
var _ String.Readable
var _ Path.ToNode

/*
Holds physics-related properties of a surface, namely its roughness and bounciness. This class is used to apply these properties to a physics body.
*/
type Instance [1]gdclass.PhysicsMaterial

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPhysicsMaterial() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PhysicsMaterial

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicsMaterial"))
	casted := Instance{*(*gdclass.PhysicsMaterial)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Friction() Float.X {
	return Float.X(Float.X(class(self).GetFriction()))
}

func (self Instance) SetFriction(value Float.X) {
	class(self).SetFriction(gd.Float(value))
}

func (self Instance) Rough() bool {
	return bool(class(self).IsRough())
}

func (self Instance) SetRough(value bool) {
	class(self).SetRough(value)
}

func (self Instance) Bounce() Float.X {
	return Float.X(Float.X(class(self).GetBounce()))
}

func (self Instance) SetBounce(value Float.X) {
	class(self).SetBounce(gd.Float(value))
}

func (self Instance) Absorbent() bool {
	return bool(class(self).IsAbsorbent())
}

func (self Instance) SetAbsorbent(value bool) {
	class(self).SetAbsorbent(value)
}

//go:nosplit
func (self class) SetFriction(friction gd.Float) { //gd:PhysicsMaterial.set_friction
	var frame = callframe.New()
	callframe.Arg(frame, friction)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsMaterial.Bind_set_friction, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFriction() gd.Float { //gd:PhysicsMaterial.get_friction
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsMaterial.Bind_get_friction, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRough(rough bool) { //gd:PhysicsMaterial.set_rough
	var frame = callframe.New()
	callframe.Arg(frame, rough)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsMaterial.Bind_set_rough, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsRough() bool { //gd:PhysicsMaterial.is_rough
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsMaterial.Bind_is_rough, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBounce(bounce gd.Float) { //gd:PhysicsMaterial.set_bounce
	var frame = callframe.New()
	callframe.Arg(frame, bounce)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsMaterial.Bind_set_bounce, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBounce() gd.Float { //gd:PhysicsMaterial.get_bounce
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsMaterial.Bind_get_bounce, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAbsorbent(absorbent bool) { //gd:PhysicsMaterial.set_absorbent
	var frame = callframe.New()
	callframe.Arg(frame, absorbent)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsMaterial.Bind_set_absorbent, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsAbsorbent() bool { //gd:PhysicsMaterial.is_absorbent
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsMaterial.Bind_is_absorbent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPhysicsMaterial() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPhysicsMaterial() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("PhysicsMaterial", func(ptr gd.Object) any {
		return [1]gdclass.PhysicsMaterial{*(*gdclass.PhysicsMaterial)(unsafe.Pointer(&ptr))}
	})
}
