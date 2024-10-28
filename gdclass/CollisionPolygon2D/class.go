package CollisionPolygon2D

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
A node that provides a polygon shape to a [CollisionObject2D] parent and allows to edit it. The polygon can be concave or convex. This can give a detection shape to an [Area2D], turn [PhysicsBody2D] into a solid object, or give a hollow shape to a [StaticBody2D].
[b]Warning:[/b] A non-uniformly scaled [CollisionShape2D] will likely not behave as expected. Make sure to keep its scale the same on all axes and adjust its shape resource instead.

*/
type Go [1]classdb.CollisionPolygon2D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.CollisionPolygon2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CollisionPolygon2D"))
	return Go{classdb.CollisionPolygon2D(object)}
}

func (self Go) BuildMode() classdb.CollisionPolygon2DBuildMode {
		return classdb.CollisionPolygon2DBuildMode(class(self).GetBuildMode())
}

func (self Go) SetBuildMode(value classdb.CollisionPolygon2DBuildMode) {
	class(self).SetBuildMode(value)
}

func (self Go) Polygon() []gd.Vector2 {
		return []gd.Vector2(class(self).GetPolygon().AsSlice())
}

func (self Go) SetPolygon(value []gd.Vector2) {
	class(self).SetPolygon(gd.NewPackedVector2Slice(value))
}

func (self Go) Disabled() bool {
		return bool(class(self).IsDisabled())
}

func (self Go) SetDisabled(value bool) {
	class(self).SetDisabled(value)
}

func (self Go) OneWayCollision() bool {
		return bool(class(self).IsOneWayCollisionEnabled())
}

func (self Go) SetOneWayCollision(value bool) {
	class(self).SetOneWayCollision(value)
}

func (self Go) OneWayCollisionMargin() float64 {
		return float64(float64(class(self).GetOneWayCollisionMargin()))
}

func (self Go) SetOneWayCollisionMargin(value float64) {
	class(self).SetOneWayCollisionMargin(gd.Float(value))
}

//go:nosplit
func (self class) SetPolygon(polygon gd.PackedVector2Array)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(polygon))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionPolygon2D.Bind_set_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPolygon() gd.PackedVector2Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionPolygon2D.Bind_get_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBuildMode(build_mode classdb.CollisionPolygon2DBuildMode)  {
	var frame = callframe.New()
	callframe.Arg(frame, build_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionPolygon2D.Bind_set_build_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetBuildMode() classdb.CollisionPolygon2DBuildMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CollisionPolygon2DBuildMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionPolygon2D.Bind_get_build_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDisabled(disabled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, disabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionPolygon2D.Bind_set_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsDisabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionPolygon2D.Bind_is_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOneWayCollision(enabled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionPolygon2D.Bind_set_one_way_collision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsOneWayCollisionEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionPolygon2D.Bind_is_one_way_collision_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetOneWayCollisionMargin(margin gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionPolygon2D.Bind_set_one_way_collision_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetOneWayCollisionMargin() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CollisionPolygon2D.Bind_get_one_way_collision_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCollisionPolygon2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsCollisionPolygon2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("CollisionPolygon2D", func(ptr gd.Object) any { return classdb.CollisionPolygon2D(ptr) })}
type BuildMode = classdb.CollisionPolygon2DBuildMode

const (
/*Collisions will include the polygon and its contained area. In this mode the node has the same effect as several [ConvexPolygonShape2D] nodes, one for each convex shape in the convex decomposition of the polygon (but without the overhead of multiple nodes).*/
	BuildSolids BuildMode = 0
/*Collisions will only include the polygon edges. In this mode the node has the same effect as a single [ConcavePolygonShape2D] made of segments, with the restriction that each segment (after the first one) starts where the previous one ends, and the last one ends where the first one starts (forming a closed but hollow polygon).*/
	BuildSegments BuildMode = 1
)
