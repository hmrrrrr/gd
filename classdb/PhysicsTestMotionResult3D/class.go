// Package PhysicsTestMotionResult3D provides methods for working with PhysicsTestMotionResult3D object instances.
package PhysicsTestMotionResult3D

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
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
import "graphics.gd/variant/Vector3"

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
Describes the motion and collision result from [method PhysicsServer3D.body_test_motion].
*/
type Instance [1]gdclass.PhysicsTestMotionResult3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPhysicsTestMotionResult3D() Instance
}

/*
Returns the moving object's travel before collision.
*/
func (self Instance) GetTravel() Vector3.XYZ { //gd:PhysicsTestMotionResult3D.get_travel
	return Vector3.XYZ(class(self).GetTravel())
}

/*
Returns the moving object's remaining movement vector.
*/
func (self Instance) GetRemainder() Vector3.XYZ { //gd:PhysicsTestMotionResult3D.get_remainder
	return Vector3.XYZ(class(self).GetRemainder())
}

/*
Returns the maximum fraction of the motion that can occur without a collision, between [code]0[/code] and [code]1[/code].
*/
func (self Instance) GetCollisionSafeFraction() Float.X { //gd:PhysicsTestMotionResult3D.get_collision_safe_fraction
	return Float.X(Float.X(class(self).GetCollisionSafeFraction()))
}

/*
Returns the minimum fraction of the motion needed to collide, if a collision occurred, between [code]0[/code] and [code]1[/code].
*/
func (self Instance) GetCollisionUnsafeFraction() Float.X { //gd:PhysicsTestMotionResult3D.get_collision_unsafe_fraction
	return Float.X(Float.X(class(self).GetCollisionUnsafeFraction()))
}

/*
Returns the number of detected collisions.
*/
func (self Instance) GetCollisionCount() int { //gd:PhysicsTestMotionResult3D.get_collision_count
	return int(int(class(self).GetCollisionCount()))
}

/*
Returns the point of collision in global coordinates given a collision index (the deepest collision by default), if a collision occurred.
*/
func (self Instance) GetCollisionPoint() Vector3.XYZ { //gd:PhysicsTestMotionResult3D.get_collision_point
	return Vector3.XYZ(class(self).GetCollisionPoint(int64(0)))
}

/*
Returns the colliding body's shape's normal at the point of collision given a collision index (the deepest collision by default), if a collision occurred.
*/
func (self Instance) GetCollisionNormal() Vector3.XYZ { //gd:PhysicsTestMotionResult3D.get_collision_normal
	return Vector3.XYZ(class(self).GetCollisionNormal(int64(0)))
}

/*
Returns the colliding body's velocity given a collision index (the deepest collision by default), if a collision occurred.
*/
func (self Instance) GetColliderVelocity() Vector3.XYZ { //gd:PhysicsTestMotionResult3D.get_collider_velocity
	return Vector3.XYZ(class(self).GetColliderVelocity(int64(0)))
}

/*
Returns the unique instance ID of the colliding body's attached [Object] given a collision index (the deepest collision by default), if a collision occurred. See [method Object.get_instance_id].
*/
func (self Instance) GetColliderId() int { //gd:PhysicsTestMotionResult3D.get_collider_id
	return int(int(class(self).GetColliderId(int64(0))))
}

/*
Returns the colliding body's [RID] used by the [PhysicsServer3D] given a collision index (the deepest collision by default), if a collision occurred.
*/
func (self Instance) GetColliderRid() RID.Body3D { //gd:PhysicsTestMotionResult3D.get_collider_rid
	return RID.Body3D(class(self).GetColliderRid(int64(0)))
}

/*
Returns the colliding body's attached [Object] given a collision index (the deepest collision by default), if a collision occurred.
*/
func (self Instance) GetCollider() Object.Instance { //gd:PhysicsTestMotionResult3D.get_collider
	return Object.Instance(class(self).GetCollider(int64(0)))
}

/*
Returns the colliding body's shape index given a collision index (the deepest collision by default), if a collision occurred. See [CollisionObject3D].
*/
func (self Instance) GetColliderShape() int { //gd:PhysicsTestMotionResult3D.get_collider_shape
	return int(int(class(self).GetColliderShape(int64(0))))
}

/*
Returns the moving object's colliding shape given a collision index (the deepest collision by default), if a collision occurred.
*/
func (self Instance) GetCollisionLocalShape() int { //gd:PhysicsTestMotionResult3D.get_collision_local_shape
	return int(int(class(self).GetCollisionLocalShape(int64(0))))
}

/*
Returns the length of overlap along the collision normal given a collision index (the deepest collision by default), if a collision occurred.
*/
func (self Instance) GetCollisionDepth() Float.X { //gd:PhysicsTestMotionResult3D.get_collision_depth
	return Float.X(Float.X(class(self).GetCollisionDepth(int64(0))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PhysicsTestMotionResult3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicsTestMotionResult3D"))
	casted := Instance{*(*gdclass.PhysicsTestMotionResult3D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Returns the moving object's travel before collision.
*/
//go:nosplit
func (self class) GetTravel() Vector3.XYZ { //gd:PhysicsTestMotionResult3D.get_travel
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionResult3D.Bind_get_travel, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the moving object's remaining movement vector.
*/
//go:nosplit
func (self class) GetRemainder() Vector3.XYZ { //gd:PhysicsTestMotionResult3D.get_remainder
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionResult3D.Bind_get_remainder, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the maximum fraction of the motion that can occur without a collision, between [code]0[/code] and [code]1[/code].
*/
//go:nosplit
func (self class) GetCollisionSafeFraction() float64 { //gd:PhysicsTestMotionResult3D.get_collision_safe_fraction
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionResult3D.Bind_get_collision_safe_fraction, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the minimum fraction of the motion needed to collide, if a collision occurred, between [code]0[/code] and [code]1[/code].
*/
//go:nosplit
func (self class) GetCollisionUnsafeFraction() float64 { //gd:PhysicsTestMotionResult3D.get_collision_unsafe_fraction
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionResult3D.Bind_get_collision_unsafe_fraction, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of detected collisions.
*/
//go:nosplit
func (self class) GetCollisionCount() int64 { //gd:PhysicsTestMotionResult3D.get_collision_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionResult3D.Bind_get_collision_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the point of collision in global coordinates given a collision index (the deepest collision by default), if a collision occurred.
*/
//go:nosplit
func (self class) GetCollisionPoint(collision_index int64) Vector3.XYZ { //gd:PhysicsTestMotionResult3D.get_collision_point
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionResult3D.Bind_get_collision_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the colliding body's shape's normal at the point of collision given a collision index (the deepest collision by default), if a collision occurred.
*/
//go:nosplit
func (self class) GetCollisionNormal(collision_index int64) Vector3.XYZ { //gd:PhysicsTestMotionResult3D.get_collision_normal
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionResult3D.Bind_get_collision_normal, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the colliding body's velocity given a collision index (the deepest collision by default), if a collision occurred.
*/
//go:nosplit
func (self class) GetColliderVelocity(collision_index int64) Vector3.XYZ { //gd:PhysicsTestMotionResult3D.get_collider_velocity
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionResult3D.Bind_get_collider_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the unique instance ID of the colliding body's attached [Object] given a collision index (the deepest collision by default), if a collision occurred. See [method Object.get_instance_id].
*/
//go:nosplit
func (self class) GetColliderId(collision_index int64) int64 { //gd:PhysicsTestMotionResult3D.get_collider_id
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionResult3D.Bind_get_collider_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the colliding body's [RID] used by the [PhysicsServer3D] given a collision index (the deepest collision by default), if a collision occurred.
*/
//go:nosplit
func (self class) GetColliderRid(collision_index int64) RID.Any { //gd:PhysicsTestMotionResult3D.get_collider_rid
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionResult3D.Bind_get_collider_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the colliding body's attached [Object] given a collision index (the deepest collision by default), if a collision occurred.
*/
//go:nosplit
func (self class) GetCollider(collision_index int64) [1]gd.Object { //gd:PhysicsTestMotionResult3D.get_collider
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionResult3D.Bind_get_collider, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(r_ret.Get())})}
	frame.Free()
	return ret
}

/*
Returns the colliding body's shape index given a collision index (the deepest collision by default), if a collision occurred. See [CollisionObject3D].
*/
//go:nosplit
func (self class) GetColliderShape(collision_index int64) int64 { //gd:PhysicsTestMotionResult3D.get_collider_shape
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionResult3D.Bind_get_collider_shape, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the moving object's colliding shape given a collision index (the deepest collision by default), if a collision occurred.
*/
//go:nosplit
func (self class) GetCollisionLocalShape(collision_index int64) int64 { //gd:PhysicsTestMotionResult3D.get_collision_local_shape
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionResult3D.Bind_get_collision_local_shape, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the length of overlap along the collision normal given a collision index (the deepest collision by default), if a collision occurred.
*/
//go:nosplit
func (self class) GetCollisionDepth(collision_index int64) float64 { //gd:PhysicsTestMotionResult3D.get_collision_depth
	var frame = callframe.New()
	callframe.Arg(frame, collision_index)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsTestMotionResult3D.Bind_get_collision_depth, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPhysicsTestMotionResult3D() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPhysicsTestMotionResult3D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("PhysicsTestMotionResult3D", func(ptr gd.Object) any {
		return [1]gdclass.PhysicsTestMotionResult3D{*(*gdclass.PhysicsTestMotionResult3D)(unsafe.Pointer(&ptr))}
	})
}
