package NavigationRegion3D

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Node3D"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A traversable 3D region based on a [NavigationMesh] that [NavigationAgent3D]s can use for pathfinding.
Two regions can be connected to each other if they share a similar edge. You can set the minimum distance between two vertices required to connect two edges by using [method NavigationServer3D.map_set_edge_connection_margin].
[b]Note:[/b] Overlapping two regions' navigation meshes is not enough for connecting two regions. They must share a similar edge.
The cost of entering this region from another region can be controlled with the [member enter_cost] value.
[b]Note:[/b] This value is not added to the path cost when the start position is already inside this region.
The cost of traveling distances inside this region can be controlled with the [member travel_cost] multiplier.
[b]Note:[/b] This node caches changes to its properties, so if you make changes to the underlying region [RID] in [NavigationServer3D], they will not be reflected in this node's properties.

*/
type Simple [1]classdb.NavigationRegion3D
func (self Simple) GetRid() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).GetRid())
}
func (self Simple) SetNavigationMesh(navigation_mesh [1]classdb.NavigationMesh) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNavigationMesh(navigation_mesh)
}
func (self Simple) GetNavigationMesh() [1]classdb.NavigationMesh {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.NavigationMesh(Expert(self).GetNavigationMesh(gc))
}
func (self Simple) SetEnabled(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnabled(enabled)
}
func (self Simple) IsEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsEnabled())
}
func (self Simple) SetNavigationMap(navigation_map gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNavigationMap(navigation_map)
}
func (self Simple) GetNavigationMap() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).GetNavigationMap())
}
func (self Simple) SetUseEdgeConnections(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetUseEdgeConnections(enabled)
}
func (self Simple) GetUseEdgeConnections() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetUseEdgeConnections())
}
func (self Simple) SetNavigationLayers(navigation_layers int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNavigationLayers(gd.Int(navigation_layers))
}
func (self Simple) GetNavigationLayers() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetNavigationLayers()))
}
func (self Simple) SetNavigationLayerValue(layer_number int, value bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNavigationLayerValue(gd.Int(layer_number), value)
}
func (self Simple) GetNavigationLayerValue(layer_number int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).GetNavigationLayerValue(gd.Int(layer_number)))
}
func (self Simple) GetRegionRid() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).GetRegionRid())
}
func (self Simple) SetEnterCost(enter_cost float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnterCost(gd.Float(enter_cost))
}
func (self Simple) GetEnterCost() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetEnterCost()))
}
func (self Simple) SetTravelCost(travel_cost float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTravelCost(gd.Float(travel_cost))
}
func (self Simple) GetTravelCost() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetTravelCost()))
}
func (self Simple) BakeNavigationMesh(on_thread bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BakeNavigationMesh(on_thread)
}
func (self Simple) IsBaking() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsBaking())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.NavigationRegion3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns the [RID] of this region on the [NavigationServer3D]. Combined with [method NavigationServer3D.map_get_closest_point_owner] can be used to identify the [NavigationRegion3D] closest to a point on the merged navigation map.
*/
//go:nosplit
func (self class) GetRid() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion3D.Bind_get_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetNavigationMesh(navigation_mesh object.NavigationMesh)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(navigation_mesh[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion3D.Bind_set_navigation_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetNavigationMesh(ctx gd.Lifetime) object.NavigationMesh {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion3D.Bind_get_navigation_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.NavigationMesh
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnabled(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion3D.Bind_set_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion3D.Bind_is_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [RID] of the navigation map this region should use. By default the region will automatically join the [World3D] default navigation map so this function is only required to override the default map.
*/
//go:nosplit
func (self class) SetNavigationMap(navigation_map gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, navigation_map)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion3D.Bind_set_navigation_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the current navigation map [RID] used by this region.
*/
//go:nosplit
func (self class) GetNavigationMap() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion3D.Bind_get_navigation_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseEdgeConnections(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion3D.Bind_set_use_edge_connections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetUseEdgeConnections() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion3D.Bind_get_use_edge_connections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetNavigationLayers(navigation_layers gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, navigation_layers)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion3D.Bind_set_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetNavigationLayers() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion3D.Bind_get_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Based on [param value], enables or disables the specified layer in the [member navigation_layers] bitmask, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetNavigationLayerValue(layer_number gd.Int, value bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion3D.Bind_set_navigation_layer_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether or not the specified layer of the [member navigation_layers] bitmask is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetNavigationLayerValue(layer_number gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion3D.Bind_get_navigation_layer_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [RID] of this region on the [NavigationServer3D].
*/
//go:nosplit
func (self class) GetRegionRid() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion3D.Bind_get_region_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnterCost(enter_cost gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enter_cost)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion3D.Bind_set_enter_cost, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnterCost() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion3D.Bind_get_enter_cost, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTravelCost(travel_cost gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, travel_cost)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion3D.Bind_set_travel_cost, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTravelCost() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion3D.Bind_get_travel_cost, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Bakes the [NavigationMesh]. If [param on_thread] is set to [code]true[/code] (default), the baking is done on a separate thread. Baking on separate thread is useful because navigation baking is not a cheap operation. When it is completed, it automatically sets the new [NavigationMesh]. Please note that baking on separate thread may be very slow if geometry is parsed from meshes as async access to each mesh involves heavy synchronization. Also, please note that baking on a separate thread is automatically disabled on operating systems that cannot use threads (such as Web with threads disabled).
*/
//go:nosplit
func (self class) BakeNavigationMesh(on_thread bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, on_thread)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion3D.Bind_bake_navigation_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] when the [NavigationMesh] is being baked on a background thread.
*/
//go:nosplit
func (self class) IsBaking() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.NavigationRegion3D.Bind_is_baking, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsNavigationRegion3D() Expert { return self[0].AsNavigationRegion3D() }


//go:nosplit
func (self Simple) AsNavigationRegion3D() Simple { return self[0].AsNavigationRegion3D() }


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
func init() {classdb.Register("NavigationRegion3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
