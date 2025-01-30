// Package CylinderMesh provides methods for working with CylinderMesh object instances.
package CylinderMesh

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Mesh"
import "graphics.gd/classdb/PrimitiveMesh"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

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
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
Class representing a cylindrical [PrimitiveMesh]. This class can be used to create cones by setting either the [member top_radius] or [member bottom_radius] properties to [code]0.0[/code].
*/
type Instance [1]gdclass.CylinderMesh

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCylinderMesh() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.CylinderMesh

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CylinderMesh"))
	casted := Instance{*(*gdclass.CylinderMesh)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) TopRadius() Float.X {
	return Float.X(Float.X(class(self).GetTopRadius()))
}

func (self Instance) SetTopRadius(value Float.X) {
	class(self).SetTopRadius(float64(value))
}

func (self Instance) BottomRadius() Float.X {
	return Float.X(Float.X(class(self).GetBottomRadius()))
}

func (self Instance) SetBottomRadius(value Float.X) {
	class(self).SetBottomRadius(float64(value))
}

func (self Instance) Height() Float.X {
	return Float.X(Float.X(class(self).GetHeight()))
}

func (self Instance) SetHeight(value Float.X) {
	class(self).SetHeight(float64(value))
}

func (self Instance) RadialSegments() int {
	return int(int(class(self).GetRadialSegments()))
}

func (self Instance) SetRadialSegments(value int) {
	class(self).SetRadialSegments(int64(value))
}

func (self Instance) Rings() int {
	return int(int(class(self).GetRings()))
}

func (self Instance) SetRings(value int) {
	class(self).SetRings(int64(value))
}

func (self Instance) CapTop() bool {
	return bool(class(self).IsCapTop())
}

func (self Instance) SetCapTop(value bool) {
	class(self).SetCapTop(value)
}

func (self Instance) CapBottom() bool {
	return bool(class(self).IsCapBottom())
}

func (self Instance) SetCapBottom(value bool) {
	class(self).SetCapBottom(value)
}

//go:nosplit
func (self class) SetTopRadius(radius float64) { //gd:CylinderMesh.set_top_radius
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CylinderMesh.Bind_set_top_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTopRadius() float64 { //gd:CylinderMesh.get_top_radius
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CylinderMesh.Bind_get_top_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBottomRadius(radius float64) { //gd:CylinderMesh.set_bottom_radius
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CylinderMesh.Bind_set_bottom_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBottomRadius() float64 { //gd:CylinderMesh.get_bottom_radius
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CylinderMesh.Bind_get_bottom_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHeight(height float64) { //gd:CylinderMesh.set_height
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CylinderMesh.Bind_set_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHeight() float64 { //gd:CylinderMesh.get_height
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CylinderMesh.Bind_get_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRadialSegments(segments int64) { //gd:CylinderMesh.set_radial_segments
	var frame = callframe.New()
	callframe.Arg(frame, segments)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CylinderMesh.Bind_set_radial_segments, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRadialSegments() int64 { //gd:CylinderMesh.get_radial_segments
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CylinderMesh.Bind_get_radial_segments, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRings(rings int64) { //gd:CylinderMesh.set_rings
	var frame = callframe.New()
	callframe.Arg(frame, rings)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CylinderMesh.Bind_set_rings, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRings() int64 { //gd:CylinderMesh.get_rings
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CylinderMesh.Bind_get_rings, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCapTop(cap_top bool) { //gd:CylinderMesh.set_cap_top
	var frame = callframe.New()
	callframe.Arg(frame, cap_top)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CylinderMesh.Bind_set_cap_top, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsCapTop() bool { //gd:CylinderMesh.is_cap_top
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CylinderMesh.Bind_is_cap_top, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCapBottom(cap_bottom bool) { //gd:CylinderMesh.set_cap_bottom
	var frame = callframe.New()
	callframe.Arg(frame, cap_bottom)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CylinderMesh.Bind_set_cap_bottom, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsCapBottom() bool { //gd:CylinderMesh.is_cap_bottom
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CylinderMesh.Bind_is_cap_bottom, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCylinderMesh() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCylinderMesh() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsPrimitiveMesh() PrimitiveMesh.Advanced {
	return *((*PrimitiveMesh.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPrimitiveMesh() PrimitiveMesh.Instance {
	return *((*PrimitiveMesh.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsMesh() Mesh.Advanced    { return *((*Mesh.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMesh() Mesh.Instance { return *((*Mesh.Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(PrimitiveMesh.Advanced(self.AsPrimitiveMesh()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(PrimitiveMesh.Instance(self.AsPrimitiveMesh()), name)
	}
}
func init() {
	gdclass.Register("CylinderMesh", func(ptr gd.Object) any {
		return [1]gdclass.CylinderMesh{*(*gdclass.CylinderMesh)(unsafe.Pointer(&ptr))}
	})
}
