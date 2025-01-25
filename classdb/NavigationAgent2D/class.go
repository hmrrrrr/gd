// Package NavigationAgent2D provides methods for working with NavigationAgent2D object instances.
package NavigationAgent2D

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
import "graphics.gd/classdb/Node"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Color"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any

/*
A 2D agent used to pathfind to a position while avoiding static and dynamic obstacles. The calculation can be used by the parent node to dynamically move it along the path. Requires navigation data to work correctly.
Dynamic obstacles are avoided using RVO collision avoidance. Avoidance is computed before physics, so the pathfinding information can be used safely in the physics step.
[b]Note:[/b] After setting the [member target_position] property, the [method get_next_path_position] method must be used once every physics frame to update the internal path logic of the navigation agent. The vector position it returns should be used as the next movement position for the agent's parent node.
*/
type Instance [1]gdclass.NavigationAgent2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsNavigationAgent2D() Instance
}

/*
Returns the [RID] of this agent on the [NavigationServer2D].
*/
func (self Instance) GetRid() Resource.ID {
	return Resource.ID(class(self).GetRid())
}

/*
Based on [param value], enables or disables the specified layer in the [member navigation_layers] bitmask, given a [param layer_number] between 1 and 32.
*/
func (self Instance) SetNavigationLayerValue(layer_number int, value bool) {
	class(self).SetNavigationLayerValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member navigation_layers] bitmask is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Instance) GetNavigationLayerValue(layer_number int) bool {
	return bool(class(self).GetNavigationLayerValue(gd.Int(layer_number)))
}

/*
Sets the [RID] of the navigation map this NavigationAgent node should use and also updates the [code]agent[/code] on the NavigationServer.
*/
func (self Instance) SetNavigationMap(navigation_map Resource.ID) {
	class(self).SetNavigationMap(navigation_map)
}

/*
Returns the [RID] of the navigation map for this NavigationAgent node. This function returns always the map set on the NavigationAgent node and not the map of the abstract agent on the NavigationServer. If the agent map is changed directly with the NavigationServer API the NavigationAgent node will not be aware of the map change. Use [method set_navigation_map] to change the navigation map for the NavigationAgent and also update the agent on the NavigationServer.
*/
func (self Instance) GetNavigationMap() Resource.ID {
	return Resource.ID(class(self).GetNavigationMap())
}

/*
Returns the next position in global coordinates that can be moved to, making sure that there are no static objects in the way. If the agent does not have a navigation path, it will return the position of the agent's parent. The use of this function once every physics frame is required to update the internal path logic of the NavigationAgent.
*/
func (self Instance) GetNextPathPosition() Vector2.XY {
	return Vector2.XY(class(self).GetNextPathPosition())
}

/*
Replaces the internal velocity in the collision avoidance simulation with [param velocity]. When an agent is teleported to a new position this function should be used in the same frame. If called frequently this function can get agents stuck.
*/
func (self Instance) SetVelocityForced(velocity Vector2.XY) {
	class(self).SetVelocityForced(gd.Vector2(velocity))
}

/*
Returns the distance to the target position, using the agent's global position. The user must set [member target_position] in order for this to be accurate.
*/
func (self Instance) DistanceToTarget() Float.X {
	return Float.X(Float.X(class(self).DistanceToTarget()))
}

/*
Returns the path query result for the path the agent is currently following.
*/
func (self Instance) GetCurrentNavigationResult() [1]gdclass.NavigationPathQueryResult2D {
	return [1]gdclass.NavigationPathQueryResult2D(class(self).GetCurrentNavigationResult())
}

/*
Returns this agent's current path from start to finish in global coordinates. The path only updates when the target position is changed or the agent requires a repath. The path array is not intended to be used in direct path movement as the agent has its own internal path logic that would get corrupted by changing the path array manually. Use the intended [method get_next_path_position] once every physics frame to receive the next path point for the agents movement as this function also updates the internal path logic.
*/
func (self Instance) GetCurrentNavigationPath() []Vector2.XY {
	return []Vector2.XY(class(self).GetCurrentNavigationPath().AsSlice())
}

/*
Returns which index the agent is currently on in the navigation path's [PackedVector2Array].
*/
func (self Instance) GetCurrentNavigationPathIndex() int {
	return int(int(class(self).GetCurrentNavigationPathIndex()))
}

/*
Returns [code]true[/code] if the agent reached the target, i.e. the agent moved within [member target_desired_distance] of the [member target_position]. It may not always be possible to reach the target but it should always be possible to reach the final position. See [method get_final_position].
*/
func (self Instance) IsTargetReached() bool {
	return bool(class(self).IsTargetReached())
}

/*
Returns [code]true[/code] if [method get_final_position] is within [member target_desired_distance] of the [member target_position].
*/
func (self Instance) IsTargetReachable() bool {
	return bool(class(self).IsTargetReachable())
}

/*
Returns [code]true[/code] if the agent's navigation has finished. If the target is reachable, navigation ends when the target is reached. If the target is unreachable, navigation ends when the last waypoint of the path is reached.
[b]Note:[/b] While [code]true[/code] prefer to stop calling update functions like [method get_next_path_position]. This avoids jittering the standing agent due to calling repeated path updates.
*/
func (self Instance) IsNavigationFinished() bool {
	return bool(class(self).IsNavigationFinished())
}

/*
Returns the reachable final position of the current navigation path in global coordinates. This position can change if the agent needs to update the navigation path which makes the agent emit the [signal path_changed] signal.
*/
func (self Instance) GetFinalPosition() Vector2.XY {
	return Vector2.XY(class(self).GetFinalPosition())
}

/*
Based on [param value], enables or disables the specified layer in the [member avoidance_layers] bitmask, given a [param layer_number] between 1 and 32.
*/
func (self Instance) SetAvoidanceLayerValue(layer_number int, value bool) {
	class(self).SetAvoidanceLayerValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member avoidance_layers] bitmask is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Instance) GetAvoidanceLayerValue(layer_number int) bool {
	return bool(class(self).GetAvoidanceLayerValue(gd.Int(layer_number)))
}

/*
Based on [param value], enables or disables the specified mask in the [member avoidance_mask] bitmask, given a [param mask_number] between 1 and 32.
*/
func (self Instance) SetAvoidanceMaskValue(mask_number int, value bool) {
	class(self).SetAvoidanceMaskValue(gd.Int(mask_number), value)
}

/*
Returns whether or not the specified mask of the [member avoidance_mask] bitmask is enabled, given a [param mask_number] between 1 and 32.
*/
func (self Instance) GetAvoidanceMaskValue(mask_number int) bool {
	return bool(class(self).GetAvoidanceMaskValue(gd.Int(mask_number)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.NavigationAgent2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("NavigationAgent2D"))
	casted := Instance{*(*gdclass.NavigationAgent2D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) TargetPosition() Vector2.XY {
	return Vector2.XY(class(self).GetTargetPosition())
}

func (self Instance) SetTargetPosition(value Vector2.XY) {
	class(self).SetTargetPosition(gd.Vector2(value))
}

func (self Instance) PathDesiredDistance() Float.X {
	return Float.X(Float.X(class(self).GetPathDesiredDistance()))
}

func (self Instance) SetPathDesiredDistance(value Float.X) {
	class(self).SetPathDesiredDistance(gd.Float(value))
}

func (self Instance) TargetDesiredDistance() Float.X {
	return Float.X(Float.X(class(self).GetTargetDesiredDistance()))
}

func (self Instance) SetTargetDesiredDistance(value Float.X) {
	class(self).SetTargetDesiredDistance(gd.Float(value))
}

func (self Instance) PathMaxDistance() Float.X {
	return Float.X(Float.X(class(self).GetPathMaxDistance()))
}

func (self Instance) SetPathMaxDistance(value Float.X) {
	class(self).SetPathMaxDistance(gd.Float(value))
}

func (self Instance) NavigationLayers() int {
	return int(int(class(self).GetNavigationLayers()))
}

func (self Instance) SetNavigationLayers(value int) {
	class(self).SetNavigationLayers(gd.Int(value))
}

func (self Instance) PathfindingAlgorithm() gdclass.NavigationPathQueryParameters2DPathfindingAlgorithm {
	return gdclass.NavigationPathQueryParameters2DPathfindingAlgorithm(class(self).GetPathfindingAlgorithm())
}

func (self Instance) SetPathfindingAlgorithm(value gdclass.NavigationPathQueryParameters2DPathfindingAlgorithm) {
	class(self).SetPathfindingAlgorithm(value)
}

func (self Instance) PathPostprocessing() gdclass.NavigationPathQueryParameters2DPathPostProcessing {
	return gdclass.NavigationPathQueryParameters2DPathPostProcessing(class(self).GetPathPostprocessing())
}

func (self Instance) SetPathPostprocessing(value gdclass.NavigationPathQueryParameters2DPathPostProcessing) {
	class(self).SetPathPostprocessing(value)
}

func (self Instance) PathMetadataFlags() gdclass.NavigationPathQueryParameters2DPathMetadataFlags {
	return gdclass.NavigationPathQueryParameters2DPathMetadataFlags(class(self).GetPathMetadataFlags())
}

func (self Instance) SetPathMetadataFlags(value gdclass.NavigationPathQueryParameters2DPathMetadataFlags) {
	class(self).SetPathMetadataFlags(value)
}

func (self Instance) SimplifyPath() bool {
	return bool(class(self).GetSimplifyPath())
}

func (self Instance) SetSimplifyPath(value bool) {
	class(self).SetSimplifyPath(value)
}

func (self Instance) SimplifyEpsilon() Float.X {
	return Float.X(Float.X(class(self).GetSimplifyEpsilon()))
}

func (self Instance) SetSimplifyEpsilon(value Float.X) {
	class(self).SetSimplifyEpsilon(gd.Float(value))
}

func (self Instance) AvoidanceEnabled() bool {
	return bool(class(self).GetAvoidanceEnabled())
}

func (self Instance) SetAvoidanceEnabled(value bool) {
	class(self).SetAvoidanceEnabled(value)
}

func (self Instance) Velocity() Vector2.XY {
	return Vector2.XY(class(self).GetVelocity())
}

func (self Instance) SetVelocity(value Vector2.XY) {
	class(self).SetVelocity(gd.Vector2(value))
}

func (self Instance) Radius() Float.X {
	return Float.X(Float.X(class(self).GetRadius()))
}

func (self Instance) SetRadius(value Float.X) {
	class(self).SetRadius(gd.Float(value))
}

func (self Instance) NeighborDistance() Float.X {
	return Float.X(Float.X(class(self).GetNeighborDistance()))
}

func (self Instance) SetNeighborDistance(value Float.X) {
	class(self).SetNeighborDistance(gd.Float(value))
}

func (self Instance) MaxNeighbors() int {
	return int(int(class(self).GetMaxNeighbors()))
}

func (self Instance) SetMaxNeighbors(value int) {
	class(self).SetMaxNeighbors(gd.Int(value))
}

func (self Instance) TimeHorizonAgents() Float.X {
	return Float.X(Float.X(class(self).GetTimeHorizonAgents()))
}

func (self Instance) SetTimeHorizonAgents(value Float.X) {
	class(self).SetTimeHorizonAgents(gd.Float(value))
}

func (self Instance) TimeHorizonObstacles() Float.X {
	return Float.X(Float.X(class(self).GetTimeHorizonObstacles()))
}

func (self Instance) SetTimeHorizonObstacles(value Float.X) {
	class(self).SetTimeHorizonObstacles(gd.Float(value))
}

func (self Instance) MaxSpeed() Float.X {
	return Float.X(Float.X(class(self).GetMaxSpeed()))
}

func (self Instance) SetMaxSpeed(value Float.X) {
	class(self).SetMaxSpeed(gd.Float(value))
}

func (self Instance) AvoidanceLayers() int {
	return int(int(class(self).GetAvoidanceLayers()))
}

func (self Instance) SetAvoidanceLayers(value int) {
	class(self).SetAvoidanceLayers(gd.Int(value))
}

func (self Instance) AvoidanceMask() int {
	return int(int(class(self).GetAvoidanceMask()))
}

func (self Instance) SetAvoidanceMask(value int) {
	class(self).SetAvoidanceMask(gd.Int(value))
}

func (self Instance) AvoidancePriority() Float.X {
	return Float.X(Float.X(class(self).GetAvoidancePriority()))
}

func (self Instance) SetAvoidancePriority(value Float.X) {
	class(self).SetAvoidancePriority(gd.Float(value))
}

func (self Instance) DebugEnabled() bool {
	return bool(class(self).GetDebugEnabled())
}

func (self Instance) SetDebugEnabled(value bool) {
	class(self).SetDebugEnabled(value)
}

func (self Instance) DebugUseCustom() bool {
	return bool(class(self).GetDebugUseCustom())
}

func (self Instance) SetDebugUseCustom(value bool) {
	class(self).SetDebugUseCustom(value)
}

func (self Instance) DebugPathCustomColor() Color.RGBA {
	return Color.RGBA(class(self).GetDebugPathCustomColor())
}

func (self Instance) SetDebugPathCustomColor(value Color.RGBA) {
	class(self).SetDebugPathCustomColor(gd.Color(value))
}

func (self Instance) DebugPathCustomPointSize() Float.X {
	return Float.X(Float.X(class(self).GetDebugPathCustomPointSize()))
}

func (self Instance) SetDebugPathCustomPointSize(value Float.X) {
	class(self).SetDebugPathCustomPointSize(gd.Float(value))
}

func (self Instance) DebugPathCustomLineWidth() Float.X {
	return Float.X(Float.X(class(self).GetDebugPathCustomLineWidth()))
}

func (self Instance) SetDebugPathCustomLineWidth(value Float.X) {
	class(self).SetDebugPathCustomLineWidth(gd.Float(value))
}

/*
Returns the [RID] of this agent on the [NavigationServer2D].
*/
//go:nosplit
func (self class) GetRid() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAvoidanceEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_avoidance_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAvoidanceEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_avoidance_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPathDesiredDistance(desired_distance gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, desired_distance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_path_desired_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPathDesiredDistance() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_path_desired_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTargetDesiredDistance(desired_distance gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, desired_distance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_target_desired_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTargetDesiredDistance() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_target_desired_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRadius(radius gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRadius() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNeighborDistance(neighbor_distance gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, neighbor_distance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_neighbor_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetNeighborDistance() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_neighbor_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxNeighbors(max_neighbors gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, max_neighbors)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_max_neighbors, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxNeighbors() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_max_neighbors, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTimeHorizonAgents(time_horizon gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, time_horizon)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_time_horizon_agents, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTimeHorizonAgents() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_time_horizon_agents, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTimeHorizonObstacles(time_horizon gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, time_horizon)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_time_horizon_obstacles, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTimeHorizonObstacles() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_time_horizon_obstacles, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxSpeed(max_speed gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, max_speed)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_max_speed, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxSpeed() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_max_speed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPathMaxDistance(max_speed gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, max_speed)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_path_max_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPathMaxDistance() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_path_max_distance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNavigationLayers(navigation_layers gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, navigation_layers)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetNavigationLayers() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Based on [param value], enables or disables the specified layer in the [member navigation_layers] bitmask, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetNavigationLayerValue(layer_number gd.Int, value bool) {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_navigation_layer_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether or not the specified layer of the [member navigation_layers] bitmask is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetNavigationLayerValue(layer_number gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_navigation_layer_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPathfindingAlgorithm(pathfinding_algorithm gdclass.NavigationPathQueryParameters2DPathfindingAlgorithm) {
	var frame = callframe.New()
	callframe.Arg(frame, pathfinding_algorithm)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_pathfinding_algorithm, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPathfindingAlgorithm() gdclass.NavigationPathQueryParameters2DPathfindingAlgorithm {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.NavigationPathQueryParameters2DPathfindingAlgorithm](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_pathfinding_algorithm, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPathPostprocessing(path_postprocessing gdclass.NavigationPathQueryParameters2DPathPostProcessing) {
	var frame = callframe.New()
	callframe.Arg(frame, path_postprocessing)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_path_postprocessing, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPathPostprocessing() gdclass.NavigationPathQueryParameters2DPathPostProcessing {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.NavigationPathQueryParameters2DPathPostProcessing](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_path_postprocessing, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPathMetadataFlags(flags gdclass.NavigationPathQueryParameters2DPathMetadataFlags) {
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_path_metadata_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPathMetadataFlags() gdclass.NavigationPathQueryParameters2DPathMetadataFlags {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.NavigationPathQueryParameters2DPathMetadataFlags](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_path_metadata_flags, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the [RID] of the navigation map this NavigationAgent node should use and also updates the [code]agent[/code] on the NavigationServer.
*/
//go:nosplit
func (self class) SetNavigationMap(navigation_map gd.RID) {
	var frame = callframe.New()
	callframe.Arg(frame, navigation_map)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_navigation_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [RID] of the navigation map for this NavigationAgent node. This function returns always the map set on the NavigationAgent node and not the map of the abstract agent on the NavigationServer. If the agent map is changed directly with the NavigationServer API the NavigationAgent node will not be aware of the map change. Use [method set_navigation_map] to change the navigation map for the NavigationAgent and also update the agent on the NavigationServer.
*/
//go:nosplit
func (self class) GetNavigationMap() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_navigation_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTargetPosition(position gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_target_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTargetPosition() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_target_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSimplifyPath(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_simplify_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSimplifyPath() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_simplify_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSimplifyEpsilon(epsilon gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, epsilon)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_simplify_epsilon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSimplifyEpsilon() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_simplify_epsilon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the next position in global coordinates that can be moved to, making sure that there are no static objects in the way. If the agent does not have a navigation path, it will return the position of the agent's parent. The use of this function once every physics frame is required to update the internal path logic of the NavigationAgent.
*/
//go:nosplit
func (self class) GetNextPathPosition() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_next_path_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Replaces the internal velocity in the collision avoidance simulation with [param velocity]. When an agent is teleported to a new position this function should be used in the same frame. If called frequently this function can get agents stuck.
*/
//go:nosplit
func (self class) SetVelocityForced(velocity gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, velocity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_velocity_forced, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetVelocity(velocity gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, velocity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVelocity() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the distance to the target position, using the agent's global position. The user must set [member target_position] in order for this to be accurate.
*/
//go:nosplit
func (self class) DistanceToTarget() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_distance_to_target, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the path query result for the path the agent is currently following.
*/
//go:nosplit
func (self class) GetCurrentNavigationResult() [1]gdclass.NavigationPathQueryResult2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_current_navigation_result, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.NavigationPathQueryResult2D{gd.PointerWithOwnershipTransferredToGo[gdclass.NavigationPathQueryResult2D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns this agent's current path from start to finish in global coordinates. The path only updates when the target position is changed or the agent requires a repath. The path array is not intended to be used in direct path movement as the agent has its own internal path logic that would get corrupted by changing the path array manually. Use the intended [method get_next_path_position] once every physics frame to receive the next path point for the agents movement as this function also updates the internal path logic.
*/
//go:nosplit
func (self class) GetCurrentNavigationPath() gd.PackedVector2Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_current_navigation_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns which index the agent is currently on in the navigation path's [PackedVector2Array].
*/
//go:nosplit
func (self class) GetCurrentNavigationPathIndex() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_current_navigation_path_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the agent reached the target, i.e. the agent moved within [member target_desired_distance] of the [member target_position]. It may not always be possible to reach the target but it should always be possible to reach the final position. See [method get_final_position].
*/
//go:nosplit
func (self class) IsTargetReached() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_is_target_reached, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if [method get_final_position] is within [member target_desired_distance] of the [member target_position].
*/
//go:nosplit
func (self class) IsTargetReachable() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_is_target_reachable, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the agent's navigation has finished. If the target is reachable, navigation ends when the target is reached. If the target is unreachable, navigation ends when the last waypoint of the path is reached.
[b]Note:[/b] While [code]true[/code] prefer to stop calling update functions like [method get_next_path_position]. This avoids jittering the standing agent due to calling repeated path updates.
*/
//go:nosplit
func (self class) IsNavigationFinished() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_is_navigation_finished, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the reachable final position of the current navigation path in global coordinates. This position can change if the agent needs to update the navigation path which makes the agent emit the [signal path_changed] signal.
*/
//go:nosplit
func (self class) GetFinalPosition() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_final_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAvoidanceLayers(layers gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, layers)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_avoidance_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAvoidanceLayers() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_avoidance_layers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAvoidanceMask(mask gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_avoidance_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAvoidanceMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_avoidance_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Based on [param value], enables or disables the specified layer in the [member avoidance_layers] bitmask, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetAvoidanceLayerValue(layer_number gd.Int, value bool) {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_avoidance_layer_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether or not the specified layer of the [member avoidance_layers] bitmask is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetAvoidanceLayerValue(layer_number gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_avoidance_layer_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Based on [param value], enables or disables the specified mask in the [member avoidance_mask] bitmask, given a [param mask_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetAvoidanceMaskValue(mask_number gd.Int, value bool) {
	var frame = callframe.New()
	callframe.Arg(frame, mask_number)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_avoidance_mask_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether or not the specified mask of the [member avoidance_mask] bitmask is enabled, given a [param mask_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetAvoidanceMaskValue(mask_number gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, mask_number)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_avoidance_mask_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAvoidancePriority(priority gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_avoidance_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAvoidancePriority() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_avoidance_priority, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDebugEnabled(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_debug_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDebugEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_debug_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDebugUseCustom(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_debug_use_custom, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDebugUseCustom() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_debug_use_custom, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDebugPathCustomColor(color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_debug_path_custom_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDebugPathCustomColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_debug_path_custom_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDebugPathCustomPointSize(point_size gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, point_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_debug_path_custom_point_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDebugPathCustomPointSize() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_debug_path_custom_point_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDebugPathCustomLineWidth(line_width gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, line_width)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_set_debug_path_custom_line_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDebugPathCustomLineWidth() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationAgent2D.Bind_get_debug_path_custom_line_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnPathChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("path_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnTargetReached(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("target_reached"), gd.NewCallable(cb), 0)
}

func (self Instance) OnWaypointReached(cb func(details map[any]any)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("waypoint_reached"), gd.NewCallable(cb), 0)
}

func (self Instance) OnLinkReached(cb func(details map[any]any)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("link_reached"), gd.NewCallable(cb), 0)
}

func (self Instance) OnNavigationFinished(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("navigation_finished"), gd.NewCallable(cb), 0)
}

func (self Instance) OnVelocityComputed(cb func(safe_velocity Vector2.XY)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("velocity_computed"), gd.NewCallable(cb), 0)
}

func (self class) AsNavigationAgent2D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNavigationAgent2D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced            { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance         { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node.Advanced(self.AsNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node.Instance(self.AsNode()), name)
	}
}
func init() {
	gdclass.Register("NavigationAgent2D", func(ptr gd.Object) any {
		return [1]gdclass.NavigationAgent2D{*(*gdclass.NavigationAgent2D)(unsafe.Pointer(&ptr))}
	})
}
