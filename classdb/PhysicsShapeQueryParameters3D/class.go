// Package PhysicsShapeQueryParameters3D provides methods for working with PhysicsShapeQueryParameters3D object instances.
package PhysicsShapeQueryParameters3D

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
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Transform3D"
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

/*
By changing various properties of this object, such as the shape, you can configure the parameters for [method PhysicsDirectSpaceState3D.intersect_shape].
*/
type Instance [1]gdclass.PhysicsShapeQueryParameters3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPhysicsShapeQueryParameters3D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PhysicsShapeQueryParameters3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicsShapeQueryParameters3D"))
	casted := Instance{*(*gdclass.PhysicsShapeQueryParameters3D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) CollisionMask() int {
	return int(int(class(self).GetCollisionMask()))
}

func (self Instance) SetCollisionMask(value int) {
	class(self).SetCollisionMask(gd.Int(value))
}

func (self Instance) Exclude() []Resource.ID {
	return []Resource.ID(gd.ArrayAs[[]Resource.ID](gd.InternalArray(class(self).GetExclude())))
}

func (self Instance) SetExclude(value []Resource.ID) {
	class(self).SetExclude(gd.ArrayFromSlice[Array.Contains[gd.RID]](value))
}

func (self Instance) Margin() Float.X {
	return Float.X(Float.X(class(self).GetMargin()))
}

func (self Instance) SetMargin(value Float.X) {
	class(self).SetMargin(gd.Float(value))
}

func (self Instance) Motion() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetMotion())
}

func (self Instance) SetMotion(value Vector3.XYZ) {
	class(self).SetMotion(gd.Vector3(value))
}

func (self Instance) Shape() [1]gdclass.Resource {
	return [1]gdclass.Resource(class(self).GetShape())
}

func (self Instance) SetShape(value [1]gdclass.Resource) {
	class(self).SetShape(value)
}

func (self Instance) ShapeRid() Resource.ID {
	return Resource.ID(class(self).GetShapeRid())
}

func (self Instance) SetShapeRid(value Resource.ID) {
	class(self).SetShapeRid(value)
}

func (self Instance) Transform() Transform3D.BasisOrigin {
	return Transform3D.BasisOrigin(class(self).GetTransform())
}

func (self Instance) SetTransform(value Transform3D.BasisOrigin) {
	class(self).SetTransform(gd.Transform3D(value))
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

//go:nosplit
func (self class) SetShape(shape [1]gdclass.Resource) { //gd:PhysicsShapeQueryParameters3D.set_shape
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(shape[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsShapeQueryParameters3D.Bind_set_shape, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetShape() [1]gdclass.Resource { //gd:PhysicsShapeQueryParameters3D.get_shape
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsShapeQueryParameters3D.Bind_get_shape, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Resource{gd.PointerWithOwnershipTransferredToGo[gdclass.Resource](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShapeRid(shape gd.RID) { //gd:PhysicsShapeQueryParameters3D.set_shape_rid
	var frame = callframe.New()
	callframe.Arg(frame, shape)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsShapeQueryParameters3D.Bind_set_shape_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetShapeRid() gd.RID { //gd:PhysicsShapeQueryParameters3D.get_shape_rid
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsShapeQueryParameters3D.Bind_get_shape_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTransform(transform gd.Transform3D) { //gd:PhysicsShapeQueryParameters3D.set_transform
	var frame = callframe.New()
	callframe.Arg(frame, transform)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsShapeQueryParameters3D.Bind_set_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTransform() gd.Transform3D { //gd:PhysicsShapeQueryParameters3D.get_transform
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsShapeQueryParameters3D.Bind_get_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMotion(motion gd.Vector3) { //gd:PhysicsShapeQueryParameters3D.set_motion
	var frame = callframe.New()
	callframe.Arg(frame, motion)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsShapeQueryParameters3D.Bind_set_motion, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMotion() gd.Vector3 { //gd:PhysicsShapeQueryParameters3D.get_motion
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsShapeQueryParameters3D.Bind_get_motion, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMargin(margin gd.Float) { //gd:PhysicsShapeQueryParameters3D.set_margin
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsShapeQueryParameters3D.Bind_set_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMargin() gd.Float { //gd:PhysicsShapeQueryParameters3D.get_margin
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsShapeQueryParameters3D.Bind_get_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollisionMask(collision_mask gd.Int) { //gd:PhysicsShapeQueryParameters3D.set_collision_mask
	var frame = callframe.New()
	callframe.Arg(frame, collision_mask)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsShapeQueryParameters3D.Bind_set_collision_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCollisionMask() gd.Int { //gd:PhysicsShapeQueryParameters3D.get_collision_mask
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsShapeQueryParameters3D.Bind_get_collision_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetExclude(exclude Array.Contains[gd.RID]) { //gd:PhysicsShapeQueryParameters3D.set_exclude
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(exclude)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsShapeQueryParameters3D.Bind_set_exclude, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetExclude() Array.Contains[gd.RID] { //gd:PhysicsShapeQueryParameters3D.get_exclude
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsShapeQueryParameters3D.Bind_get_exclude, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[gd.RID]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollideWithBodies(enable bool) { //gd:PhysicsShapeQueryParameters3D.set_collide_with_bodies
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsShapeQueryParameters3D.Bind_set_collide_with_bodies, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsCollideWithBodiesEnabled() bool { //gd:PhysicsShapeQueryParameters3D.is_collide_with_bodies_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsShapeQueryParameters3D.Bind_is_collide_with_bodies_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollideWithAreas(enable bool) { //gd:PhysicsShapeQueryParameters3D.set_collide_with_areas
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsShapeQueryParameters3D.Bind_set_collide_with_areas, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsCollideWithAreasEnabled() bool { //gd:PhysicsShapeQueryParameters3D.is_collide_with_areas_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsShapeQueryParameters3D.Bind_is_collide_with_areas_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPhysicsShapeQueryParameters3D() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPhysicsShapeQueryParameters3D() Instance {
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
	gdclass.Register("PhysicsShapeQueryParameters3D", func(ptr gd.Object) any {
		return [1]gdclass.PhysicsShapeQueryParameters3D{*(*gdclass.PhysicsShapeQueryParameters3D)(unsafe.Pointer(&ptr))}
	})
}
