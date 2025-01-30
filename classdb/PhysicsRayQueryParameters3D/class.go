// Package PhysicsRayQueryParameters3D provides methods for working with PhysicsRayQueryParameters3D object instances.
package PhysicsRayQueryParameters3D

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
By changing various properties of this object, such as the ray position, you can configure the parameters for [method PhysicsDirectSpaceState3D.intersect_ray].
*/
type Instance [1]gdclass.PhysicsRayQueryParameters3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPhysicsRayQueryParameters3D() Instance
}

/*
Returns a new, pre-configured [PhysicsRayQueryParameters3D] object. Use it to quickly create query parameters using the most common options.
[codeblock]
var query = PhysicsRayQueryParameters3D.create(position, position + Vector3(0, -10, 0))
var collision = get_world_3d().direct_space_state.intersect_ray(query)
[/codeblock]
*/
func Create(from Vector3.XYZ, to Vector3.XYZ) [1]gdclass.PhysicsRayQueryParameters3D { //gd:PhysicsRayQueryParameters3D.create
	self := Instance{}
	return [1]gdclass.PhysicsRayQueryParameters3D(class(self).Create(Vector3.XYZ(from), Vector3.XYZ(to), int64(4294967295), gd.ArrayFromSlice[Array.Contains[RID.Any]]([1][]RID.Any{}[0])))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PhysicsRayQueryParameters3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicsRayQueryParameters3D"))
	casted := Instance{*(*gdclass.PhysicsRayQueryParameters3D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) From() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetFrom())
}

func (self Instance) SetFrom(value Vector3.XYZ) {
	class(self).SetFrom(Vector3.XYZ(value))
}

func (self Instance) To() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetTo())
}

func (self Instance) SetTo(value Vector3.XYZ) {
	class(self).SetTo(Vector3.XYZ(value))
}

func (self Instance) CollisionMask() int {
	return int(int(class(self).GetCollisionMask()))
}

func (self Instance) SetCollisionMask(value int) {
	class(self).SetCollisionMask(int64(value))
}

func (self Instance) Exclude() []RID.Any {
	return []RID.Any(gd.ArrayAs[[]RID.Any](gd.InternalArray(class(self).GetExclude())))
}

func (self Instance) SetExclude(value []RID.Any) {
	class(self).SetExclude(gd.ArrayFromSlice[Array.Contains[RID.Any]](value))
}

func (self Instance) CollideWithBodies() bool {
	return bool(class(self).IsCollideWithBodiesEnabled())
}

func (self Instance) SetCollideWithBodies(value bool) {
	class(self).SetCollideWithBodies(value)
}

func (self Instance) CollideWithAreas() bool {
	return bool(class(self).IsCollideWithAreasEnabled())
}

func (self Instance) SetCollideWithAreas(value bool) {
	class(self).SetCollideWithAreas(value)
}

func (self Instance) HitFromInside() bool {
	return bool(class(self).IsHitFromInsideEnabled())
}

func (self Instance) SetHitFromInside(value bool) {
	class(self).SetHitFromInside(value)
}

func (self Instance) HitBackFaces() bool {
	return bool(class(self).IsHitBackFacesEnabled())
}

func (self Instance) SetHitBackFaces(value bool) {
	class(self).SetHitBackFaces(value)
}

/*
Returns a new, pre-configured [PhysicsRayQueryParameters3D] object. Use it to quickly create query parameters using the most common options.
[codeblock]
var query = PhysicsRayQueryParameters3D.create(position, position + Vector3(0, -10, 0))
var collision = get_world_3d().direct_space_state.intersect_ray(query)
[/codeblock]
*/
//go:nosplit
func (self class) Create(from Vector3.XYZ, to Vector3.XYZ, collision_mask int64, exclude Array.Contains[RID.Any]) [1]gdclass.PhysicsRayQueryParameters3D { //gd:PhysicsRayQueryParameters3D.create
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	callframe.Arg(frame, collision_mask)
	callframe.Arg(frame, pointers.Get(gd.InternalArray(exclude)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters3D.Bind_create, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.PhysicsRayQueryParameters3D{gd.PointerWithOwnershipTransferredToGo[gdclass.PhysicsRayQueryParameters3D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFrom(from Vector3.XYZ) { //gd:PhysicsRayQueryParameters3D.set_from
	var frame = callframe.New()
	callframe.Arg(frame, from)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters3D.Bind_set_from, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrom() Vector3.XYZ { //gd:PhysicsRayQueryParameters3D.get_from
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters3D.Bind_get_from, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTo(to Vector3.XYZ) { //gd:PhysicsRayQueryParameters3D.set_to
	var frame = callframe.New()
	callframe.Arg(frame, to)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters3D.Bind_set_to, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTo() Vector3.XYZ { //gd:PhysicsRayQueryParameters3D.get_to
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters3D.Bind_get_to, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollisionMask(collision_mask int64) { //gd:PhysicsRayQueryParameters3D.set_collision_mask
	var frame = callframe.New()
	callframe.Arg(frame, collision_mask)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters3D.Bind_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCollisionMask() int64 { //gd:PhysicsRayQueryParameters3D.get_collision_mask
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters3D.Bind_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetExclude(exclude Array.Contains[RID.Any]) { //gd:PhysicsRayQueryParameters3D.set_exclude
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(exclude)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters3D.Bind_set_exclude, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetExclude() Array.Contains[RID.Any] { //gd:PhysicsRayQueryParameters3D.get_exclude
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters3D.Bind_get_exclude, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[RID.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollideWithBodies(enable bool) { //gd:PhysicsRayQueryParameters3D.set_collide_with_bodies
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters3D.Bind_set_collide_with_bodies, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsCollideWithBodiesEnabled() bool { //gd:PhysicsRayQueryParameters3D.is_collide_with_bodies_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters3D.Bind_is_collide_with_bodies_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollideWithAreas(enable bool) { //gd:PhysicsRayQueryParameters3D.set_collide_with_areas
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters3D.Bind_set_collide_with_areas, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsCollideWithAreasEnabled() bool { //gd:PhysicsRayQueryParameters3D.is_collide_with_areas_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters3D.Bind_is_collide_with_areas_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHitFromInside(enable bool) { //gd:PhysicsRayQueryParameters3D.set_hit_from_inside
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters3D.Bind_set_hit_from_inside, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsHitFromInsideEnabled() bool { //gd:PhysicsRayQueryParameters3D.is_hit_from_inside_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters3D.Bind_is_hit_from_inside_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHitBackFaces(enable bool) { //gd:PhysicsRayQueryParameters3D.set_hit_back_faces
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters3D.Bind_set_hit_back_faces, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsHitBackFacesEnabled() bool { //gd:PhysicsRayQueryParameters3D.is_hit_back_faces_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsRayQueryParameters3D.Bind_is_hit_back_faces_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPhysicsRayQueryParameters3D() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPhysicsRayQueryParameters3D() Instance {
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
	gdclass.Register("PhysicsRayQueryParameters3D", func(ptr gd.Object) any {
		return [1]gdclass.PhysicsRayQueryParameters3D{*(*gdclass.PhysicsRayQueryParameters3D)(unsafe.Pointer(&ptr))}
	})
}
