package CSGSphere3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/CSGPrimitive3D"
import "grow.graphics/gd/gdclass/CSGShape3D"
import "grow.graphics/gd/gdclass/GeometryInstance3D"
import "grow.graphics/gd/gdclass/VisualInstance3D"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
This node allows you to create a sphere for use with the CSG system.
[b]Note:[/b] CSG nodes are intended to be used for level prototyping. Creating CSG nodes has a significant CPU cost compared to creating a [MeshInstance3D] with a [PrimitiveMesh]. Moving a CSG node within another CSG node also has a significant CPU cost, so it should be avoided during gameplay.
*/
type Instance [1]classdb.CSGSphere3D

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.CSGSphere3D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CSGSphere3D"))
	return Instance{classdb.CSGSphere3D(object)}
}

func (self Instance) Radius() float64 {
	return float64(float64(class(self).GetRadius()))
}

func (self Instance) SetRadius(value float64) {
	class(self).SetRadius(gd.Float(value))
}

func (self Instance) RadialSegments() int {
	return int(int(class(self).GetRadialSegments()))
}

func (self Instance) SetRadialSegments(value int) {
	class(self).SetRadialSegments(gd.Int(value))
}

func (self Instance) Rings() int {
	return int(int(class(self).GetRings()))
}

func (self Instance) SetRings(value int) {
	class(self).SetRings(gd.Int(value))
}

func (self Instance) SmoothFaces() bool {
	return bool(class(self).GetSmoothFaces())
}

func (self Instance) SetSmoothFaces(value bool) {
	class(self).SetSmoothFaces(value)
}

func (self Instance) Material() gdclass.Material {
	return gdclass.Material(class(self).GetMaterial())
}

func (self Instance) SetMaterial(value gdclass.Material) {
	class(self).SetMaterial(value)
}

//go:nosplit
func (self class) SetRadius(radius gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGSphere3D.Bind_set_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRadius() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGSphere3D.Bind_get_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRadialSegments(radial_segments gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, radial_segments)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGSphere3D.Bind_set_radial_segments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRadialSegments() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGSphere3D.Bind_get_radial_segments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRings(rings gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, rings)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGSphere3D.Bind_set_rings, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRings() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGSphere3D.Bind_get_rings, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSmoothFaces(smooth_faces bool) {
	var frame = callframe.New()
	callframe.Arg(frame, smooth_faces)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGSphere3D.Bind_set_smooth_faces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSmoothFaces() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGSphere3D.Bind_get_smooth_faces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaterial(material gdclass.Material) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(material[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGSphere3D.Bind_set_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaterial() gdclass.Material {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CSGSphere3D.Bind_get_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Material{classdb.Material(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self class) AsCSGSphere3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCSGSphere3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsCSGPrimitive3D() CSGPrimitive3D.Advanced {
	return *((*CSGPrimitive3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCSGPrimitive3D() CSGPrimitive3D.Instance {
	return *((*CSGPrimitive3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsCSGShape3D() CSGShape3D.Advanced {
	return *((*CSGShape3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCSGShape3D() CSGShape3D.Instance {
	return *((*CSGShape3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsGeometryInstance3D() GeometryInstance3D.Advanced {
	return *((*GeometryInstance3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsGeometryInstance3D() GeometryInstance3D.Instance {
	return *((*GeometryInstance3D.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(self.AsCSGPrimitive3D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsCSGPrimitive3D(), name)
	}
}
func init() {
	classdb.Register("CSGSphere3D", func(ptr gd.Object) any { return classdb.CSGSphere3D(ptr) })
}
