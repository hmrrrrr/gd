// Package SphereShape3D provides methods for working with SphereShape3D object instances.
package SphereShape3D

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Shape3D"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A 3D sphere shape, intended for use in physics. Usually used to provide a shape for a [CollisionShape3D].
[b]Performance:[/b] [SphereShape3D] is fast to check collisions against. It is faster than [BoxShape3D], [CapsuleShape3D], and [CylinderShape3D].
*/
type Instance [1]gdclass.SphereShape3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsSphereShape3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.SphereShape3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SphereShape3D"))
	casted := Instance{*(*gdclass.SphereShape3D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Radius() Float.X {
	return Float.X(Float.X(class(self).GetRadius()))
}

func (self Instance) SetRadius(value Float.X) {
	class(self).SetRadius(gd.Float(value))
}

//go:nosplit
func (self class) SetRadius(radius gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SphereShape3D.Bind_set_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRadius() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SphereShape3D.Bind_get_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsSphereShape3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsSphereShape3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsShape3D() Shape3D.Advanced  { return *((*Shape3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsShape3D() Shape3D.Instance {
	return *((*Shape3D.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(Shape3D.Advanced(self.AsShape3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Shape3D.Instance(self.AsShape3D()), name)
	}
}
func init() {
	gdclass.Register("SphereShape3D", func(ptr gd.Object) any {
		return [1]gdclass.SphereShape3D{*(*gdclass.SphereShape3D)(unsafe.Pointer(&ptr))}
	})
}
