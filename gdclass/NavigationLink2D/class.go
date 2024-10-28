package NavigationLink2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node2D"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
A link between two positions on [NavigationRegion2D]s that agents can be routed through. These positions can be on the same [NavigationRegion2D] or on two different ones. Links are useful to express navigation methods other than traveling along the surface of the navigation polygon, such as ziplines, teleporters, or gaps that can be jumped across.

*/
type Go [1]classdb.NavigationLink2D

/*
Returns the [RID] of this link on the [NavigationServer2D].
*/
func (self Go) GetRid() gd.RID {
	return gd.RID(class(self).GetRid())
}

/*
Based on [param value], enables or disables the specified layer in the [member navigation_layers] bitmask, given a [param layer_number] between 1 and 32.
*/
func (self Go) SetNavigationLayerValue(layer_number int, value bool) {
	class(self).SetNavigationLayerValue(gd.Int(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member navigation_layers] bitmask is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Go) GetNavigationLayerValue(layer_number int) bool {
	return bool(class(self).GetNavigationLayerValue(gd.Int(layer_number)))
}

/*
Sets the [member start_position] that is relative to the link from a global [param position].
*/
func (self Go) SetGlobalStartPosition(position gd.Vector2) {
	class(self).SetGlobalStartPosition(position)
}

/*
Returns the [member start_position] that is relative to the link as a global position.
*/
func (self Go) GetGlobalStartPosition() gd.Vector2 {
	return gd.Vector2(class(self).GetGlobalStartPosition())
}

/*
Sets the [member end_position] that is relative to the link from a global [param position].
*/
func (self Go) SetGlobalEndPosition(position gd.Vector2) {
	class(self).SetGlobalEndPosition(position)
}

/*
Returns the [member end_position] that is relative to the link as a global position.
*/
func (self Go) GetGlobalEndPosition() gd.Vector2 {
	return gd.Vector2(class(self).GetGlobalEndPosition())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.NavigationLink2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("NavigationLink2D"))
	return Go{classdb.NavigationLink2D(object)}
}

func (self Go) Enabled() bool {
		return bool(class(self).IsEnabled())
}

func (self Go) SetEnabled(value bool) {
	class(self).SetEnabled(value)
}

func (self Go) Bidirectional() bool {
		return bool(class(self).IsBidirectional())
}

func (self Go) SetBidirectional(value bool) {
	class(self).SetBidirectional(value)
}

func (self Go) NavigationLayers() int {
		return int(int(class(self).GetNavigationLayers()))
}

func (self Go) SetNavigationLayers(value int) {
	class(self).SetNavigationLayers(gd.Int(value))
}

func (self Go) StartPosition() gd.Vector2 {
		return gd.Vector2(class(self).GetStartPosition())
}

func (self Go) SetStartPosition(value gd.Vector2) {
	class(self).SetStartPosition(value)
}

func (self Go) EndPosition() gd.Vector2 {
		return gd.Vector2(class(self).GetEndPosition())
}

func (self Go) SetEndPosition(value gd.Vector2) {
	class(self).SetEndPosition(value)
}

func (self Go) EnterCost() float64 {
		return float64(float64(class(self).GetEnterCost()))
}

func (self Go) SetEnterCost(value float64) {
	class(self).SetEnterCost(gd.Float(value))
}

func (self Go) TravelCost() float64 {
		return float64(float64(class(self).GetTravelCost()))
}

func (self Go) SetTravelCost(value float64) {
	class(self).SetTravelCost(gd.Float(value))
}

/*
Returns the [RID] of this link on the [NavigationServer2D].
*/
//go:nosplit
func (self class) GetRid() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_get_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnabled(enabled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_is_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBidirectional(bidirectional bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, bidirectional)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_bidirectional, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsBidirectional() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_is_bidirectional, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetNavigationLayers(navigation_layers gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, navigation_layers)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetNavigationLayers() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_get_navigation_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Based on [param value], enables or disables the specified layer in the [member navigation_layers] bitmask, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetNavigationLayerValue(layer_number gd.Int, value bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_navigation_layer_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
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
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_get_navigation_layer_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStartPosition(position gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_start_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStartPosition() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_get_start_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEndPosition(position gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_end_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEndPosition() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_get_end_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [member start_position] that is relative to the link from a global [param position].
*/
//go:nosplit
func (self class) SetGlobalStartPosition(position gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_global_start_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [member start_position] that is relative to the link as a global position.
*/
//go:nosplit
func (self class) GetGlobalStartPosition() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_get_global_start_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [member end_position] that is relative to the link from a global [param position].
*/
//go:nosplit
func (self class) SetGlobalEndPosition(position gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_global_end_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [member end_position] that is relative to the link as a global position.
*/
//go:nosplit
func (self class) GetGlobalEndPosition() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_get_global_end_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnterCost(enter_cost gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, enter_cost)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_enter_cost, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnterCost() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_get_enter_cost, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTravelCost(travel_cost gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, travel_cost)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_set_travel_cost, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTravelCost() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.NavigationLink2D.Bind_get_travel_cost, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsNavigationLink2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsNavigationLink2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.GD { return *((*Node2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode2D() Node2D.Go { return *((*Node2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode2D(), name)
	}
}
func init() {classdb.Register("NavigationLink2D", func(ptr gd.Object) any { return classdb.NavigationLink2D(ptr) })}
