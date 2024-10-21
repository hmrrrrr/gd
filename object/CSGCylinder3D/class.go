package CSGCylinder3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/CSGPrimitive3D"
import "grow.graphics/gd/object/CSGShape3D"
import "grow.graphics/gd/object/GeometryInstance3D"
import "grow.graphics/gd/object/VisualInstance3D"
import "grow.graphics/gd/object/Node3D"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This node allows you to create a cylinder (or cone) for use with the CSG system.
[b]Note:[/b] CSG nodes are intended to be used for level prototyping. Creating CSG nodes has a significant CPU cost compared to creating a [MeshInstance3D] with a [PrimitiveMesh]. Moving a CSG node within another CSG node also has a significant CPU cost, so it should be avoided during gameplay.

*/
type Simple [1]classdb.CSGCylinder3D
func (self Simple) SetRadius(radius float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRadius(gd.Float(radius))
}
func (self Simple) GetRadius() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetRadius()))
}
func (self Simple) SetHeight(height float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHeight(gd.Float(height))
}
func (self Simple) GetHeight() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetHeight()))
}
func (self Simple) SetSides(sides int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSides(gd.Int(sides))
}
func (self Simple) GetSides() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetSides()))
}
func (self Simple) SetCone(cone bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCone(cone)
}
func (self Simple) IsCone() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsCone())
}
func (self Simple) SetMaterial(material [1]classdb.Material) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMaterial(material)
}
func (self Simple) GetMaterial() [1]classdb.Material {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Material(Expert(self).GetMaterial(gc))
}
func (self Simple) SetSmoothFaces(smooth_faces bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSmoothFaces(smooth_faces)
}
func (self Simple) GetSmoothFaces() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetSmoothFaces())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.CSGCylinder3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetRadius(radius gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGCylinder3D.Bind_set_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRadius() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGCylinder3D.Bind_get_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHeight(height gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGCylinder3D.Bind_set_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHeight() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGCylinder3D.Bind_get_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSides(sides gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, sides)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGCylinder3D.Bind_set_sides, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSides() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGCylinder3D.Bind_get_sides, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCone(cone bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, cone)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGCylinder3D.Bind_set_cone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCone() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGCylinder3D.Bind_is_cone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaterial(material object.Material)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(material[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGCylinder3D.Bind_set_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaterial(ctx gd.Lifetime) object.Material {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGCylinder3D.Bind_get_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Material
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSmoothFaces(smooth_faces bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, smooth_faces)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGCylinder3D.Bind_set_smooth_faces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSmoothFaces() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CSGCylinder3D.Bind_get_smooth_faces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsCSGCylinder3D() Expert { return self[0].AsCSGCylinder3D() }


//go:nosplit
func (self Simple) AsCSGCylinder3D() Simple { return self[0].AsCSGCylinder3D() }


//go:nosplit
func (self class) AsCSGPrimitive3D() CSGPrimitive3D.Expert { return self[0].AsCSGPrimitive3D() }


//go:nosplit
func (self Simple) AsCSGPrimitive3D() CSGPrimitive3D.Simple { return self[0].AsCSGPrimitive3D() }


//go:nosplit
func (self class) AsCSGShape3D() CSGShape3D.Expert { return self[0].AsCSGShape3D() }


//go:nosplit
func (self Simple) AsCSGShape3D() CSGShape3D.Simple { return self[0].AsCSGShape3D() }


//go:nosplit
func (self class) AsGeometryInstance3D() GeometryInstance3D.Expert { return self[0].AsGeometryInstance3D() }


//go:nosplit
func (self Simple) AsGeometryInstance3D() GeometryInstance3D.Simple { return self[0].AsGeometryInstance3D() }


//go:nosplit
func (self class) AsVisualInstance3D() VisualInstance3D.Expert { return self[0].AsVisualInstance3D() }


//go:nosplit
func (self Simple) AsVisualInstance3D() VisualInstance3D.Simple { return self[0].AsVisualInstance3D() }


//go:nosplit
func (self class) AsNode3D() Node3D.Expert { return self[0].AsNode3D() }


//go:nosplit
func (self Simple) AsNode3D() Node3D.Simple { return self[0].AsNode3D() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("CSGCylinder3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
