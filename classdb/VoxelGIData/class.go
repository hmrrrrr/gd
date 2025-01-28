// Package VoxelGIData provides methods for working with VoxelGIData object instances.
package VoxelGIData

import "unsafe"
import "reflect"
import "slices"
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
import "graphics.gd/variant/Packed"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Transform3D"
import "graphics.gd/variant/AABB"
import "graphics.gd/variant/Vector3"
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
var _ Packed.Bytes
var _ = slices.Delete[[]struct{}, struct{}]

/*
[VoxelGIData] contains baked voxel global illumination for use in a [VoxelGI] node. [VoxelGIData] also offers several properties to adjust the final appearance of the global illumination. These properties can be adjusted at run-time without having to bake the [VoxelGI] node again.
[b]Note:[/b] To prevent text-based scene files ([code].tscn[/code]) from growing too much and becoming slow to load and save, always save [VoxelGIData] to an external binary resource file ([code].res[/code]) instead of embedding it within the scene. This can be done by clicking the dropdown arrow next to the [VoxelGIData] resource, choosing [b]Edit[/b], clicking the floppy disk icon at the top of the Inspector then choosing [b]Save As...[/b].
*/
type Instance [1]gdclass.VoxelGIData

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVoxelGIData() Instance
}

func (self Instance) Allocate(to_cell_xform Transform3D.BasisOrigin, aabb AABB.PositionSize, octree_size Vector3.XYZ, octree_cells []byte, data_cells []byte, distance_field []byte, level_counts []int32) { //gd:VoxelGIData.allocate
	class(self).Allocate(gd.Transform3D(to_cell_xform), gd.AABB(aabb), gd.Vector3(octree_size), Packed.Bytes(Packed.New(octree_cells...)), Packed.Bytes(Packed.New(data_cells...)), Packed.Bytes(Packed.New(distance_field...)), Packed.New(level_counts...))
}

/*
Returns the bounds of the baked voxel data as an [AABB], which should match [member VoxelGI.size] after being baked (which only contains the size as a [Vector3]).
[b]Note:[/b] If the size was modified without baking the VoxelGI data, then the value of [method get_bounds] and [member VoxelGI.size] will not match.
*/
func (self Instance) GetBounds() AABB.PositionSize { //gd:VoxelGIData.get_bounds
	return AABB.PositionSize(class(self).GetBounds())
}
func (self Instance) GetOctreeSize() Vector3.XYZ { //gd:VoxelGIData.get_octree_size
	return Vector3.XYZ(class(self).GetOctreeSize())
}
func (self Instance) GetToCellXform() Transform3D.BasisOrigin { //gd:VoxelGIData.get_to_cell_xform
	return Transform3D.BasisOrigin(class(self).GetToCellXform())
}
func (self Instance) GetOctreeCells() []byte { //gd:VoxelGIData.get_octree_cells
	return []byte(class(self).GetOctreeCells().Bytes())
}
func (self Instance) GetDataCells() []byte { //gd:VoxelGIData.get_data_cells
	return []byte(class(self).GetDataCells().Bytes())
}
func (self Instance) GetLevelCounts() []int32 { //gd:VoxelGIData.get_level_counts
	return []int32(slices.Collect(class(self).GetLevelCounts().Values()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VoxelGIData

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VoxelGIData"))
	casted := Instance{*(*gdclass.VoxelGIData)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) DynamicRange() Float.X {
	return Float.X(Float.X(class(self).GetDynamicRange()))
}

func (self Instance) SetDynamicRange(value Float.X) {
	class(self).SetDynamicRange(gd.Float(value))
}

func (self Instance) Energy() Float.X {
	return Float.X(Float.X(class(self).GetEnergy()))
}

func (self Instance) SetEnergy(value Float.X) {
	class(self).SetEnergy(gd.Float(value))
}

func (self Instance) Bias() Float.X {
	return Float.X(Float.X(class(self).GetBias()))
}

func (self Instance) SetBias(value Float.X) {
	class(self).SetBias(gd.Float(value))
}

func (self Instance) NormalBias() Float.X {
	return Float.X(Float.X(class(self).GetNormalBias()))
}

func (self Instance) SetNormalBias(value Float.X) {
	class(self).SetNormalBias(gd.Float(value))
}

func (self Instance) Propagation() Float.X {
	return Float.X(Float.X(class(self).GetPropagation()))
}

func (self Instance) SetPropagation(value Float.X) {
	class(self).SetPropagation(gd.Float(value))
}

func (self Instance) UseTwoBounces() bool {
	return bool(class(self).IsUsingTwoBounces())
}

func (self Instance) SetUseTwoBounces(value bool) {
	class(self).SetUseTwoBounces(value)
}

func (self Instance) Interior() bool {
	return bool(class(self).IsInterior())
}

func (self Instance) SetInterior(value bool) {
	class(self).SetInterior(value)
}

//go:nosplit
func (self class) Allocate(to_cell_xform gd.Transform3D, aabb gd.AABB, octree_size gd.Vector3, octree_cells Packed.Bytes, data_cells Packed.Bytes, distance_field Packed.Bytes, level_counts Packed.Array[int32]) { //gd:VoxelGIData.allocate
	var frame = callframe.New()
	callframe.Arg(frame, to_cell_xform)
	callframe.Arg(frame, aabb)
	callframe.Arg(frame, octree_size)
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedByteArray, byte](Packed.Array[byte](octree_cells))))
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedByteArray, byte](Packed.Array[byte](data_cells))))
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedByteArray, byte](Packed.Array[byte](distance_field))))
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedInt32Array, int32](level_counts)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGIData.Bind_allocate, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the bounds of the baked voxel data as an [AABB], which should match [member VoxelGI.size] after being baked (which only contains the size as a [Vector3]).
[b]Note:[/b] If the size was modified without baking the VoxelGI data, then the value of [method get_bounds] and [member VoxelGI.size] will not match.
*/
//go:nosplit
func (self class) GetBounds() gd.AABB { //gd:VoxelGIData.get_bounds
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.AABB](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGIData.Bind_get_bounds, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetOctreeSize() gd.Vector3 { //gd:VoxelGIData.get_octree_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGIData.Bind_get_octree_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetToCellXform() gd.Transform3D { //gd:VoxelGIData.get_to_cell_xform
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGIData.Bind_get_to_cell_xform, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetOctreeCells() Packed.Bytes { //gd:VoxelGIData.get_octree_cells
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGIData.Bind_get_octree_cells, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Bytes(Array.Through(gd.PackedProxy[gd.PackedByteArray, byte]{}, pointers.Pack(pointers.New[gd.PackedByteArray](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetDataCells() Packed.Bytes { //gd:VoxelGIData.get_data_cells
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGIData.Bind_get_data_cells, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Bytes(Array.Through(gd.PackedProxy[gd.PackedByteArray, byte]{}, pointers.Pack(pointers.New[gd.PackedByteArray](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetLevelCounts() Packed.Array[int32] { //gd:VoxelGIData.get_level_counts
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGIData.Bind_get_level_counts, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int32](Array.Through(gd.PackedProxy[gd.PackedInt32Array, int32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDynamicRange(dynamic_range gd.Float) { //gd:VoxelGIData.set_dynamic_range
	var frame = callframe.New()
	callframe.Arg(frame, dynamic_range)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGIData.Bind_set_dynamic_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDynamicRange() gd.Float { //gd:VoxelGIData.get_dynamic_range
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGIData.Bind_get_dynamic_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnergy(energy gd.Float) { //gd:VoxelGIData.set_energy
	var frame = callframe.New()
	callframe.Arg(frame, energy)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGIData.Bind_set_energy, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnergy() gd.Float { //gd:VoxelGIData.get_energy
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGIData.Bind_get_energy, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBias(bias gd.Float) { //gd:VoxelGIData.set_bias
	var frame = callframe.New()
	callframe.Arg(frame, bias)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGIData.Bind_set_bias, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBias() gd.Float { //gd:VoxelGIData.get_bias
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGIData.Bind_get_bias, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNormalBias(bias gd.Float) { //gd:VoxelGIData.set_normal_bias
	var frame = callframe.New()
	callframe.Arg(frame, bias)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGIData.Bind_set_normal_bias, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetNormalBias() gd.Float { //gd:VoxelGIData.get_normal_bias
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGIData.Bind_get_normal_bias, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPropagation(propagation gd.Float) { //gd:VoxelGIData.set_propagation
	var frame = callframe.New()
	callframe.Arg(frame, propagation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGIData.Bind_set_propagation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPropagation() gd.Float { //gd:VoxelGIData.get_propagation
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGIData.Bind_get_propagation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInterior(interior bool) { //gd:VoxelGIData.set_interior
	var frame = callframe.New()
	callframe.Arg(frame, interior)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGIData.Bind_set_interior, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsInterior() bool { //gd:VoxelGIData.is_interior
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGIData.Bind_is_interior, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseTwoBounces(enable bool) { //gd:VoxelGIData.set_use_two_bounces
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGIData.Bind_set_use_two_bounces, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingTwoBounces() bool { //gd:VoxelGIData.is_using_two_bounces
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VoxelGIData.Bind_is_using_two_bounces, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsVoxelGIData() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVoxelGIData() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("VoxelGIData", func(ptr gd.Object) any { return [1]gdclass.VoxelGIData{*(*gdclass.VoxelGIData)(unsafe.Pointer(&ptr))} })
}
