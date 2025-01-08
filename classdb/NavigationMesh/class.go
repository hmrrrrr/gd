// Package NavigationMesh provides methods for working with NavigationMesh object instances.
package NavigationMesh

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/AABB"
import "graphics.gd/variant/Vector3"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A navigation mesh is a collection of polygons that define which areas of an environment are traversable to aid agents in pathfinding through complicated spaces.
*/
type Instance [1]gdclass.NavigationMesh

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsNavigationMesh() Instance
}

/*
Based on [param value], enables or disables the specified layer in the [member geometry_collision_mask], given a [param layer_number] between 1 and 32.
*/
func (self Instance) SetCollisionMaskValue(layer_number int, value bool) {
	class(self).SetCollisionMaskValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member geometry_collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Instance) GetCollisionMaskValue(layer_number int) bool {
	return bool(class(self).GetCollisionMaskValue(gd.Int(layer_number)))
}

/*
Adds a polygon using the indices of the vertices you get when calling [method get_vertices].
*/
func (self Instance) AddPolygon(polygon []int32) {
	class(self).AddPolygon(gd.NewPackedInt32Slice(polygon))
}

/*
Returns the number of polygons in the navigation mesh.
*/
func (self Instance) GetPolygonCount() int {
	return int(int(class(self).GetPolygonCount()))
}

/*
Returns a [PackedInt32Array] containing the indices of the vertices of a created polygon.
*/
func (self Instance) GetPolygon(idx int) []int32 {
	return []int32(class(self).GetPolygon(gd.Int(idx)).AsSlice())
}

/*
Clears the array of polygons, but it doesn't clear the array of vertices.
*/
func (self Instance) ClearPolygons() {
	class(self).ClearPolygons()
}

/*
Initializes the navigation mesh by setting the vertices and indices according to a [Mesh].
[b]Note:[/b] The given [param mesh] must be of type [constant Mesh.PRIMITIVE_TRIANGLES] and have an index array.
*/
func (self Instance) CreateFromMesh(mesh [1]gdclass.Mesh) {
	class(self).CreateFromMesh(mesh)
}

/*
Clears the internal arrays for vertices and polygon indices.
*/
func (self Instance) Clear() {
	class(self).Clear()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.NavigationMesh

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("NavigationMesh"))
	casted := Instance{*(*gdclass.NavigationMesh)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Vertices() []Vector3.XYZ {
	return []Vector3.XYZ(class(self).GetVertices().AsSlice())
}

func (self Instance) SetVertices(value []Vector3.XYZ) {
	class(self).SetVertices(gd.NewPackedVector3Slice(*(*[]gd.Vector3)(unsafe.Pointer(&value))))
}

func (self Instance) SamplePartitionType() gdclass.NavigationMeshSamplePartitionType {
	return gdclass.NavigationMeshSamplePartitionType(class(self).GetSamplePartitionType())
}

func (self Instance) SetSamplePartitionType(value gdclass.NavigationMeshSamplePartitionType) {
	class(self).SetSamplePartitionType(value)
}

func (self Instance) GeometryParsedGeometryType() gdclass.NavigationMeshParsedGeometryType {
	return gdclass.NavigationMeshParsedGeometryType(class(self).GetParsedGeometryType())
}

func (self Instance) SetGeometryParsedGeometryType(value gdclass.NavigationMeshParsedGeometryType) {
	class(self).SetParsedGeometryType(value)
}

func (self Instance) GeometryCollisionMask() int {
	return int(int(class(self).GetCollisionMask()))
}

func (self Instance) SetGeometryCollisionMask(value int) {
	class(self).SetCollisionMask(gd.Int(value))
}

func (self Instance) GeometrySourceGeometryMode() gdclass.NavigationMeshSourceGeometryMode {
	return gdclass.NavigationMeshSourceGeometryMode(class(self).GetSourceGeometryMode())
}

func (self Instance) SetGeometrySourceGeometryMode(value gdclass.NavigationMeshSourceGeometryMode) {
	class(self).SetSourceGeometryMode(value)
}

func (self Instance) GeometrySourceGroupName() string {
	return string(class(self).GetSourceGroupName().String())
}

func (self Instance) SetGeometrySourceGroupName(value string) {
	class(self).SetSourceGroupName(gd.NewStringName(value))
}

func (self Instance) CellSize() Float.X {
	return Float.X(Float.X(class(self).GetCellSize()))
}

func (self Instance) SetCellSize(value Float.X) {
	class(self).SetCellSize(gd.Float(value))
}

func (self Instance) CellHeight() Float.X {
	return Float.X(Float.X(class(self).GetCellHeight()))
}

func (self Instance) SetCellHeight(value Float.X) {
	class(self).SetCellHeight(gd.Float(value))
}

func (self Instance) BorderSize() Float.X {
	return Float.X(Float.X(class(self).GetBorderSize()))
}

func (self Instance) SetBorderSize(value Float.X) {
	class(self).SetBorderSize(gd.Float(value))
}

func (self Instance) AgentHeight() Float.X {
	return Float.X(Float.X(class(self).GetAgentHeight()))
}

func (self Instance) SetAgentHeight(value Float.X) {
	class(self).SetAgentHeight(gd.Float(value))
}

func (self Instance) AgentRadius() Float.X {
	return Float.X(Float.X(class(self).GetAgentRadius()))
}

func (self Instance) SetAgentRadius(value Float.X) {
	class(self).SetAgentRadius(gd.Float(value))
}

func (self Instance) AgentMaxClimb() Float.X {
	return Float.X(Float.X(class(self).GetAgentMaxClimb()))
}

func (self Instance) SetAgentMaxClimb(value Float.X) {
	class(self).SetAgentMaxClimb(gd.Float(value))
}

func (self Instance) AgentMaxSlope() Float.X {
	return Float.X(Float.X(class(self).GetAgentMaxSlope()))
}

func (self Instance) SetAgentMaxSlope(value Float.X) {
	class(self).SetAgentMaxSlope(gd.Float(value))
}

func (self Instance) RegionMinSize() Float.X {
	return Float.X(Float.X(class(self).GetRegionMinSize()))
}

func (self Instance) SetRegionMinSize(value Float.X) {
	class(self).SetRegionMinSize(gd.Float(value))
}

func (self Instance) RegionMergeSize() Float.X {
	return Float.X(Float.X(class(self).GetRegionMergeSize()))
}

func (self Instance) SetRegionMergeSize(value Float.X) {
	class(self).SetRegionMergeSize(gd.Float(value))
}

func (self Instance) EdgeMaxLength() Float.X {
	return Float.X(Float.X(class(self).GetEdgeMaxLength()))
}

func (self Instance) SetEdgeMaxLength(value Float.X) {
	class(self).SetEdgeMaxLength(gd.Float(value))
}

func (self Instance) EdgeMaxError() Float.X {
	return Float.X(Float.X(class(self).GetEdgeMaxError()))
}

func (self Instance) SetEdgeMaxError(value Float.X) {
	class(self).SetEdgeMaxError(gd.Float(value))
}

func (self Instance) VerticesPerPolygon() Float.X {
	return Float.X(Float.X(class(self).GetVerticesPerPolygon()))
}

func (self Instance) SetVerticesPerPolygon(value Float.X) {
	class(self).SetVerticesPerPolygon(gd.Float(value))
}

func (self Instance) DetailSampleDistance() Float.X {
	return Float.X(Float.X(class(self).GetDetailSampleDistance()))
}

func (self Instance) SetDetailSampleDistance(value Float.X) {
	class(self).SetDetailSampleDistance(gd.Float(value))
}

func (self Instance) DetailSampleMaxError() Float.X {
	return Float.X(Float.X(class(self).GetDetailSampleMaxError()))
}

func (self Instance) SetDetailSampleMaxError(value Float.X) {
	class(self).SetDetailSampleMaxError(gd.Float(value))
}

func (self Instance) FilterLowHangingObstacles() bool {
	return bool(class(self).GetFilterLowHangingObstacles())
}

func (self Instance) SetFilterLowHangingObstacles(value bool) {
	class(self).SetFilterLowHangingObstacles(value)
}

func (self Instance) FilterLedgeSpans() bool {
	return bool(class(self).GetFilterLedgeSpans())
}

func (self Instance) SetFilterLedgeSpans(value bool) {
	class(self).SetFilterLedgeSpans(value)
}

func (self Instance) FilterWalkableLowHeightSpans() bool {
	return bool(class(self).GetFilterWalkableLowHeightSpans())
}

func (self Instance) SetFilterWalkableLowHeightSpans(value bool) {
	class(self).SetFilterWalkableLowHeightSpans(value)
}

func (self Instance) FilterBakingAabb() AABB.PositionSize {
	return AABB.PositionSize(class(self).GetFilterBakingAabb())
}

func (self Instance) SetFilterBakingAabb(value AABB.PositionSize) {
	class(self).SetFilterBakingAabb(gd.AABB(value))
}

func (self Instance) FilterBakingAabbOffset() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetFilterBakingAabbOffset())
}

func (self Instance) SetFilterBakingAabbOffset(value Vector3.XYZ) {
	class(self).SetFilterBakingAabbOffset(gd.Vector3(value))
}

//go:nosplit
func (self class) SetSamplePartitionType(sample_partition_type gdclass.NavigationMeshSamplePartitionType) {
	var frame = callframe.New()
	callframe.Arg(frame, sample_partition_type)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_set_sample_partition_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSamplePartitionType() gdclass.NavigationMeshSamplePartitionType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.NavigationMeshSamplePartitionType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_get_sample_partition_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetParsedGeometryType(geometry_type gdclass.NavigationMeshParsedGeometryType) {
	var frame = callframe.New()
	callframe.Arg(frame, geometry_type)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_set_parsed_geometry_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetParsedGeometryType() gdclass.NavigationMeshParsedGeometryType {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.NavigationMeshParsedGeometryType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_get_parsed_geometry_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollisionMask(mask gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCollisionMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Based on [param value], enables or disables the specified layer in the [member geometry_collision_mask], given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetCollisionMaskValue(layer_number gd.Int, value bool) {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_set_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether or not the specified layer of the [member geometry_collision_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetCollisionMaskValue(layer_number gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_get_collision_mask_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSourceGeometryMode(mask gdclass.NavigationMeshSourceGeometryMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_set_source_geometry_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSourceGeometryMode() gdclass.NavigationMeshSourceGeometryMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.NavigationMeshSourceGeometryMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_get_source_geometry_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSourceGroupName(mask gd.StringName) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(mask))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_set_source_group_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSourceGroupName() gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_get_source_group_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCellSize(cell_size gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, cell_size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_set_cell_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCellSize() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_get_cell_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCellHeight(cell_height gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, cell_height)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_set_cell_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCellHeight() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_get_cell_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBorderSize(border_size gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, border_size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_set_border_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBorderSize() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_get_border_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAgentHeight(agent_height gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, agent_height)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_set_agent_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAgentHeight() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_get_agent_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAgentRadius(agent_radius gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, agent_radius)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_set_agent_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAgentRadius() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_get_agent_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAgentMaxClimb(agent_max_climb gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, agent_max_climb)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_set_agent_max_climb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAgentMaxClimb() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_get_agent_max_climb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAgentMaxSlope(agent_max_slope gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, agent_max_slope)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_set_agent_max_slope, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAgentMaxSlope() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_get_agent_max_slope, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRegionMinSize(region_min_size gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, region_min_size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_set_region_min_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRegionMinSize() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_get_region_min_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRegionMergeSize(region_merge_size gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, region_merge_size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_set_region_merge_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRegionMergeSize() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_get_region_merge_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEdgeMaxLength(edge_max_length gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, edge_max_length)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_set_edge_max_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEdgeMaxLength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_get_edge_max_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEdgeMaxError(edge_max_error gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, edge_max_error)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_set_edge_max_error, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEdgeMaxError() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_get_edge_max_error, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVerticesPerPolygon(vertices_per_polygon gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, vertices_per_polygon)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_set_vertices_per_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVerticesPerPolygon() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_get_vertices_per_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDetailSampleDistance(detail_sample_dist gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, detail_sample_dist)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_set_detail_sample_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDetailSampleDistance() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_get_detail_sample_distance, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDetailSampleMaxError(detail_sample_max_error gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, detail_sample_max_error)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_set_detail_sample_max_error, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDetailSampleMaxError() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_get_detail_sample_max_error, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFilterLowHangingObstacles(filter_low_hanging_obstacles bool) {
	var frame = callframe.New()
	callframe.Arg(frame, filter_low_hanging_obstacles)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_set_filter_low_hanging_obstacles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFilterLowHangingObstacles() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_get_filter_low_hanging_obstacles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFilterLedgeSpans(filter_ledge_spans bool) {
	var frame = callframe.New()
	callframe.Arg(frame, filter_ledge_spans)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_set_filter_ledge_spans, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFilterLedgeSpans() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_get_filter_ledge_spans, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFilterWalkableLowHeightSpans(filter_walkable_low_height_spans bool) {
	var frame = callframe.New()
	callframe.Arg(frame, filter_walkable_low_height_spans)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_set_filter_walkable_low_height_spans, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFilterWalkableLowHeightSpans() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_get_filter_walkable_low_height_spans, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFilterBakingAabb(baking_aabb gd.AABB) {
	var frame = callframe.New()
	callframe.Arg(frame, baking_aabb)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_set_filter_baking_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFilterBakingAabb() gd.AABB {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.AABB](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_get_filter_baking_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFilterBakingAabbOffset(baking_aabb_offset gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, baking_aabb_offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_set_filter_baking_aabb_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFilterBakingAabbOffset() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_get_filter_baking_aabb_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the vertices that can be then indexed to create polygons with the [method add_polygon] method.
*/
//go:nosplit
func (self class) SetVertices(vertices gd.PackedVector3Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(vertices))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_set_vertices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a [PackedVector3Array] containing all the vertices being used to create the polygons.
*/
//go:nosplit
func (self class) GetVertices() gd.PackedVector3Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_get_vertices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedVector3Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Adds a polygon using the indices of the vertices you get when calling [method get_vertices].
*/
//go:nosplit
func (self class) AddPolygon(polygon gd.PackedInt32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(polygon))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_add_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the number of polygons in the navigation mesh.
*/
//go:nosplit
func (self class) GetPolygonCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_get_polygon_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a [PackedInt32Array] containing the indices of the vertices of a created polygon.
*/
//go:nosplit
func (self class) GetPolygon(idx gd.Int) gd.PackedInt32Array {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_get_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Clears the array of polygons, but it doesn't clear the array of vertices.
*/
//go:nosplit
func (self class) ClearPolygons() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_clear_polygons, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Initializes the navigation mesh by setting the vertices and indices according to a [Mesh].
[b]Note:[/b] The given [param mesh] must be of type [constant Mesh.PRIMITIVE_TRIANGLES] and have an index array.
*/
//go:nosplit
func (self class) CreateFromMesh(mesh [1]gdclass.Mesh) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(mesh[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_create_from_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Clears the internal arrays for vertices and polygon indices.
*/
//go:nosplit
func (self class) Clear() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationMesh.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsNavigationMesh() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNavigationMesh() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("NavigationMesh", func(ptr gd.Object) any {
		return [1]gdclass.NavigationMesh{*(*gdclass.NavigationMesh)(unsafe.Pointer(&ptr))}
	})
}

type SamplePartitionType = gdclass.NavigationMeshSamplePartitionType

const (
	/*Watershed partitioning. Generally the best choice if you precompute the navigation mesh, use this if you have large open areas.*/
	SamplePartitionWatershed SamplePartitionType = 0
	/*Monotone partitioning. Use this if you want fast navigation mesh generation.*/
	SamplePartitionMonotone SamplePartitionType = 1
	/*Layer partitioning. Good choice to use for tiled navigation mesh with medium and small sized tiles.*/
	SamplePartitionLayers SamplePartitionType = 2
	/*Represents the size of the [enum SamplePartitionType] enum.*/
	SamplePartitionMax SamplePartitionType = 3
)

type ParsedGeometryType = gdclass.NavigationMeshParsedGeometryType

const (
	/*Parses mesh instances as geometry. This includes [MeshInstance3D], [CSGShape3D], and [GridMap] nodes.*/
	ParsedGeometryMeshInstances ParsedGeometryType = 0
	/*Parses [StaticBody3D] colliders as geometry. The collider should be in any of the layers specified by [member geometry_collision_mask].*/
	ParsedGeometryStaticColliders ParsedGeometryType = 1
	/*Both [constant PARSED_GEOMETRY_MESH_INSTANCES] and [constant PARSED_GEOMETRY_STATIC_COLLIDERS].*/
	ParsedGeometryBoth ParsedGeometryType = 2
	/*Represents the size of the [enum ParsedGeometryType] enum.*/
	ParsedGeometryMax ParsedGeometryType = 3
)

type SourceGeometryMode = gdclass.NavigationMeshSourceGeometryMode

const (
	/*Scans the child nodes of the root node recursively for geometry.*/
	SourceGeometryRootNodeChildren SourceGeometryMode = 0
	/*Scans nodes in a group and their child nodes recursively for geometry. The group is specified by [member geometry_source_group_name].*/
	SourceGeometryGroupsWithChildren SourceGeometryMode = 1
	/*Uses nodes in a group for geometry. The group is specified by [member geometry_source_group_name].*/
	SourceGeometryGroupsExplicit SourceGeometryMode = 2
	/*Represents the size of the [enum SourceGeometryMode] enum.*/
	SourceGeometryMax SourceGeometryMode = 3
)
