package WorldBoundaryShape2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Shape2D"
import "grow.graphics/gd/objects/Resource"
import "grow.graphics/gd/variant/Vector2"
import "grow.graphics/gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A 2D world boundary shape, intended for use in physics. [WorldBoundaryShape2D] works like an infinite straight line that forces all physics bodies to stay above it. The line's normal determines which direction is considered as "above" and in the editor, the smaller line over it represents this direction. It can for example be used for endless flat floors.
*/
type Instance [1]classdb.WorldBoundaryShape2D

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.WorldBoundaryShape2D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("WorldBoundaryShape2D"))
	return Instance{classdb.WorldBoundaryShape2D(object)}
}

func (self Instance) Normal() Vector2.XY {
	return Vector2.XY(class(self).GetNormal())
}

func (self Instance) SetNormal(value Vector2.XY) {
	class(self).SetNormal(gd.Vector2(value))
}

func (self Instance) Distance() Float.X {
	return Float.X(Float.X(class(self).GetDistance()))
}

func (self Instance) SetDistance(value Float.X) {
	class(self).SetDistance(gd.Float(value))
}

//go:nosplit
func (self class) SetNormal(normal gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, normal)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorldBoundaryShape2D.Bind_set_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetNormal() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorldBoundaryShape2D.Bind_get_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDistance(distance gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, distance)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorldBoundaryShape2D.Bind_set_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDistance() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WorldBoundaryShape2D.Bind_get_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsWorldBoundaryShape2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsWorldBoundaryShape2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsShape2D() Shape2D.Advanced         { return *((*Shape2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsShape2D() Shape2D.Instance {
	return *((*Shape2D.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsShape2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsShape2D(), name)
	}
}
func init() {
	classdb.Register("WorldBoundaryShape2D", func(ptr gd.Object) any { return classdb.WorldBoundaryShape2D(ptr) })
}
