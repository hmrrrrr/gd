package PolygonOccluder3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Occluder3D"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
[PolygonOccluder3D] stores a polygon shape that can be used by the engine's occlusion culling system. When an [OccluderInstance3D] with a [PolygonOccluder3D] is selected in the editor, an editor will appear at the top of the 3D viewport so you can add/remove points. All points must be placed on the same 2D plane, which means it is not possible to create arbitrary 3D shapes with a single [PolygonOccluder3D]. To use arbitrary 3D shapes as occluders, use [ArrayOccluder3D] or [OccluderInstance3D]'s baking feature instead.
See [OccluderInstance3D]'s documentation for instructions on setting up occlusion culling.

*/
type Go [1]classdb.PolygonOccluder3D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.PolygonOccluder3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PolygonOccluder3D"))
	return Go{classdb.PolygonOccluder3D(object)}
}

func (self Go) Polygon() []gd.Vector2 {
		return []gd.Vector2(class(self).GetPolygon().AsSlice())
}

func (self Go) SetPolygon(value []gd.Vector2) {
	class(self).SetPolygon(gd.NewPackedVector2Slice(value))
}

//go:nosplit
func (self class) SetPolygon(polygon gd.PackedVector2Array)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(polygon))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PolygonOccluder3D.Bind_set_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPolygon() gd.PackedVector2Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PolygonOccluder3D.Bind_get_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsPolygonOccluder3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsPolygonOccluder3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsOccluder3D() Occluder3D.GD { return *((*Occluder3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsOccluder3D() Occluder3D.Go { return *((*Occluder3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsOccluder3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsOccluder3D(), name)
	}
}
func init() {classdb.Register("PolygonOccluder3D", func(ptr gd.Object) any { return classdb.PolygonOccluder3D(ptr) })}
