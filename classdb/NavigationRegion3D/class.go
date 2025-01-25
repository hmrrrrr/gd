// Package NavigationRegion3D provides methods for working with NavigationRegion3D object instances.
package NavigationRegion3D

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
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
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

/*
A traversable 3D region based on a [NavigationMesh] that [NavigationAgent3D]s can use for pathfinding.
Two regions can be connected to each other if they share a similar edge. You can set the minimum distance between two vertices required to connect two edges by using [method NavigationServer3D.map_set_edge_connection_margin].
[b]Note:[/b] Overlapping two regions' navigation meshes is not enough for connecting two regions. They must share a similar edge.
The cost of entering this region from another region can be controlled with the [member enter_cost] value.
[b]Note:[/b] This value is not added to the path cost when the start position is already inside this region.
The cost of traveling distances inside this region can be controlled with the [member travel_cost] multiplier.
[b]Note:[/b] This node caches changes to its properties, so if you make changes to the underlying region [RID] in [NavigationServer3D], they will not be reflected in this node's properties.
*/
type Instance [1]gdclass.NavigationRegion3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsNavigationRegion3D() Instance
}

/*
Returns the [RID] of this region on the [NavigationServer3D]. Combined with [method NavigationServer3D.map_get_closest_point_owner] can be used to identify the [NavigationRegion3D] closest to a point on the merged navigation map.
*/
func (self Instance) GetRid() Resource.ID { //gd:NavigationRegion3D.get_rid
	return Resource.ID(class(self).GetRid())
}

/*
Sets the [RID] of the navigation map this region should use. By default the region will automatically join the [World3D] default navigation map so this function is only required to override the default map.
*/
func (self Instance) SetNavigationMap(navigation_map Resource.ID) { //gd:NavigationRegion3D.set_navigation_map
	class(self).SetNavigationMap(navigation_map)
}

/*
Returns the current navigation map [RID] used by this region.
*/
func (self Instance) GetNavigationMap() Resource.ID { //gd:NavigationRegion3D.get_navigation_map
	return Resource.ID(class(self).GetNavigationMap())
}

/*
Based on [param value], enables or disables the specified layer in the [member navigation_layers] bitmask, given a [param layer_number] between 1 and 32.
*/
func (self Instance) SetNavigationLayerValue(layer_number int, value bool) { //gd:NavigationRegion3D.set_navigation_layer_value
	class(self).SetNavigationLayerValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member navigation_layers] bitmask is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Instance) GetNavigationLayerValue(layer_number int) bool { //gd:NavigationRegion3D.get_navigation_layer_value
	return bool(class(self).GetNavigationLayerValue(gd.Int(layer_number)))
}

/*
Returns the [RID] of this region on the [NavigationServer3D].
*/
func (self Instance) GetRegionRid() Resource.ID { //gd:NavigationRegion3D.get_region_rid
	return Resource.ID(class(self).GetRegionRid())
}

/*
Bakes the [NavigationMesh]. If [param on_thread] is set to [code]true[/code] (default), the baking is done on a separate thread. Baking on separate thread is useful because navigation baking is not a cheap operation. When it is completed, it automatically sets the new [NavigationMesh]. Please note that baking on separate thread may be very slow if geometry is parsed from meshes as async access to each mesh involves heavy synchronization. Also, please note that baking on a separate thread is automatically disabled on operating systems that cannot use threads (such as Web with threads disabled).
*/
func (self Instance) BakeNavigationMesh() { //gd:NavigationRegion3D.bake_navigation_mesh
	class(self).BakeNavigationMesh(true)
}

/*
Returns [code]true[/code] when the [NavigationMesh] is being baked on a background thread.
*/
func (self Instance) IsBaking() bool { //gd:NavigationRegion3D.is_baking
	return bool(class(self).IsBaking())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.NavigationRegion3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("NavigationRegion3D"))
	casted := Instance{*(*gdclass.NavigationRegion3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) NavigationMesh() [1]gdclass.NavigationMesh {
	return [1]gdclass.NavigationMesh(class(self).GetNavigationMesh())
}

func (self Instance) SetNavigationMesh(value [1]gdclass.NavigationMesh) {
	class(self).SetNavigationMesh(value)
}

func (self Instance) Enabled() bool {
	return bool(class(self).IsEnabled())
}

func (self Instance) SetEnabled(value bool) {
	class(self).SetEnabled(value)
}

func (self Instance) UseEdgeConnections() bool {
	return bool(class(self).GetUseEdgeConnections())
}

func (self Instance) SetUseEdgeConnections(value bool) {
	class(self).SetUseEdgeConnections(value)
}

func (self Instance) NavigationLayers() int {
	return int(int(class(self).GetNavigationLayers()))
}

func (self Instance) SetNavigationLayers(value int) {
	class(self).SetNavigationLayers(gd.Int(value))
}

func (self Instance) EnterCost() Float.X {
	return Float.X(Float.X(class(self).GetEnterCost()))
}

func (self Instance) SetEnterCost(value Float.X) {
	class(self).SetEnterCost(gd.Float(value))
}

func (self Instance) TravelCost() Float.X {
	return Float.X(Float.X(class(self).GetTravelCost()))
}

func (self Instance) SetTravelCost(value Float.X) {
	class(self).SetTravelCost(gd.Float(value))
}

/*
Returns the [RID] of this region on the [NavigationServer3D]. Combined with [method NavigationServer3D.map_get_closest_point_owner] can be used to identify the [NavigationRegion3D] closest to a point on the merged navigation map.
*/
//go:nosplit
func (self class) GetRid() gd.RID { //gd:NavigationRegion3D.get_rid
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationRegion3D.Bind_get_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNavigationMesh(navigation_mesh [1]gdclass.NavigationMesh) { //gd:NavigationRegion3D.set_navigation_mesh
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(navigation_mesh[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationRegion3D.Bind_set_navigation_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetNavigationMesh() [1]gdclass.NavigationMesh { //gd:NavigationRegion3D.get_navigation_mesh
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationRegion3D.Bind_get_navigation_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.NavigationMesh{gd.PointerWithOwnershipTransferredToGo[gdclass.NavigationMesh](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnabled(enabled bool) { //gd:NavigationRegion3D.set_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationRegion3D.Bind_set_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsEnabled() bool { //gd:NavigationRegion3D.is_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationRegion3D.Bind_is_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [RID] of the navigation map this region should use. By default the region will automatically join the [World3D] default navigation map so this function is only required to override the default map.
*/
//go:nosplit
func (self class) SetNavigationMap(navigation_map gd.RID) { //gd:NavigationRegion3D.set_navigation_map
	var frame = callframe.New()
	callframe.Arg(frame, navigation_map)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationRegion3D.Bind_set_navigation_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the current navigation map [RID] used by this region.
*/
//go:nosplit
func (self class) GetNavigationMap() gd.RID { //gd:NavigationRegion3D.get_navigation_map
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationRegion3D.Bind_get_navigation_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseEdgeConnections(enabled bool) { //gd:NavigationRegion3D.set_use_edge_connections
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationRegion3D.Bind_set_use_edge_connections, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUseEdgeConnections() bool { //gd:NavigationRegion3D.get_use_edge_connections
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationRegion3D.Bind_get_use_edge_connections, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNavigationLayers(navigation_layers gd.Int) { //gd:NavigationRegion3D.set_navigation_layers
	var frame = callframe.New()
	callframe.Arg(frame, navigation_layers)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationRegion3D.Bind_set_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetNavigationLayers() gd.Int { //gd:NavigationRegion3D.get_navigation_layers
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationRegion3D.Bind_get_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Based on [param value], enables or disables the specified layer in the [member navigation_layers] bitmask, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetNavigationLayerValue(layer_number gd.Int, value bool) { //gd:NavigationRegion3D.set_navigation_layer_value
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationRegion3D.Bind_set_navigation_layer_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether or not the specified layer of the [member navigation_layers] bitmask is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetNavigationLayerValue(layer_number gd.Int) bool { //gd:NavigationRegion3D.get_navigation_layer_value
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationRegion3D.Bind_get_navigation_layer_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [RID] of this region on the [NavigationServer3D].
*/
//go:nosplit
func (self class) GetRegionRid() gd.RID { //gd:NavigationRegion3D.get_region_rid
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationRegion3D.Bind_get_region_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnterCost(enter_cost gd.Float) { //gd:NavigationRegion3D.set_enter_cost
	var frame = callframe.New()
	callframe.Arg(frame, enter_cost)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationRegion3D.Bind_set_enter_cost, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnterCost() gd.Float { //gd:NavigationRegion3D.get_enter_cost
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationRegion3D.Bind_get_enter_cost, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTravelCost(travel_cost gd.Float) { //gd:NavigationRegion3D.set_travel_cost
	var frame = callframe.New()
	callframe.Arg(frame, travel_cost)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationRegion3D.Bind_set_travel_cost, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTravelCost() gd.Float { //gd:NavigationRegion3D.get_travel_cost
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationRegion3D.Bind_get_travel_cost, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Bakes the [NavigationMesh]. If [param on_thread] is set to [code]true[/code] (default), the baking is done on a separate thread. Baking on separate thread is useful because navigation baking is not a cheap operation. When it is completed, it automatically sets the new [NavigationMesh]. Please note that baking on separate thread may be very slow if geometry is parsed from meshes as async access to each mesh involves heavy synchronization. Also, please note that baking on a separate thread is automatically disabled on operating systems that cannot use threads (such as Web with threads disabled).
*/
//go:nosplit
func (self class) BakeNavigationMesh(on_thread bool) { //gd:NavigationRegion3D.bake_navigation_mesh
	var frame = callframe.New()
	callframe.Arg(frame, on_thread)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationRegion3D.Bind_bake_navigation_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] when the [NavigationMesh] is being baked on a background thread.
*/
//go:nosplit
func (self class) IsBaking() bool { //gd:NavigationRegion3D.is_baking
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationRegion3D.Bind_is_baking, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnNavigationMeshChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("navigation_mesh_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnBakeFinished(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("bake_finished"), gd.NewCallable(cb), 0)
}

func (self class) AsNavigationRegion3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNavigationRegion3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.Advanced         { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance      { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced             { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance          { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Advanced(self.AsNode3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Instance(self.AsNode3D()), name)
	}
}
func init() {
	gdclass.Register("NavigationRegion3D", func(ptr gd.Object) any {
		return [1]gdclass.NavigationRegion3D{*(*gdclass.NavigationRegion3D)(unsafe.Pointer(&ptr))}
	})
}
