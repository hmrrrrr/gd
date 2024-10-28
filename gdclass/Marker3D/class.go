package Marker3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Generic 3D position hint for editing. It's just like a plain [Node3D], but it displays as a cross in the 3D editor at all times.

*/
type Go [1]classdb.Marker3D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.Marker3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Marker3D"))
	return Go{classdb.Marker3D(object)}
}

func (self Go) GizmoExtents() float64 {
		return float64(float64(class(self).GetGizmoExtents()))
}

func (self Go) SetGizmoExtents(value float64) {
	class(self).SetGizmoExtents(gd.Float(value))
}

//go:nosplit
func (self class) SetGizmoExtents(extents gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, extents)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Marker3D.Bind_set_gizmo_extents, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGizmoExtents() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Marker3D.Bind_get_gizmo_extents, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsMarker3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsMarker3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode3D(), name)
	}
}
func init() {classdb.Register("Marker3D", func(ptr gd.Object) any { return classdb.Marker3D(ptr) })}
