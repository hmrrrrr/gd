package MultiMesh

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
MultiMesh provides low-level mesh instancing. Drawing thousands of [MeshInstance3D] nodes can be slow, since each object is submitted to the GPU then drawn individually.
MultiMesh is much faster as it can draw thousands of instances with a single draw call, resulting in less API overhead.
As a drawback, if the instances are too far away from each other, performance may be reduced as every single instance will always render (they are spatially indexed as one, for the whole object).
Since instances may have any behavior, the AABB used for visibility must be provided by the user.
[b]Note:[/b] A MultiMesh is a single object, therefore the same maximum lights per object restriction applies. This means, that once the maximum lights are consumed by one or more instances, the rest of the MultiMesh instances will [b]not[/b] receive any lighting.
[b]Note:[/b] Blend Shapes will be ignored if used in a MultiMesh.

*/
type Go [1]classdb.MultiMesh

/*
Sets the [Transform3D] for a specific instance.
*/
func (self Go) SetInstanceTransform(instance int, transform gd.Transform3D) {
	class(self).SetInstanceTransform(gd.Int(instance), transform)
}

/*
Sets the [Transform2D] for a specific instance.
*/
func (self Go) SetInstanceTransform2d(instance int, transform gd.Transform2D) {
	class(self).SetInstanceTransform2d(gd.Int(instance), transform)
}

/*
Returns the [Transform3D] of a specific instance.
*/
func (self Go) GetInstanceTransform(instance int) gd.Transform3D {
	return gd.Transform3D(class(self).GetInstanceTransform(gd.Int(instance)))
}

/*
Returns the [Transform2D] of a specific instance.
*/
func (self Go) GetInstanceTransform2d(instance int) gd.Transform2D {
	return gd.Transform2D(class(self).GetInstanceTransform2d(gd.Int(instance)))
}

/*
Sets the color of a specific instance by [i]multiplying[/i] the mesh's existing vertex colors. This allows for different color tinting per instance.
[b]Note:[/b] Each component is stored in 32 bits in the Forward+ and Mobile rendering methods, but is packed into 16 bits in the Compatibility rendering method.
For the color to take effect, ensure that [member use_colors] is [code]true[/code] on the [MultiMesh] and [member BaseMaterial3D.vertex_color_use_as_albedo] is [code]true[/code] on the material. If you intend to set an absolute color instead of tinting, make sure the material's albedo color is set to pure white ([code]Color(1, 1, 1)[/code]).
*/
func (self Go) SetInstanceColor(instance int, color gd.Color) {
	class(self).SetInstanceColor(gd.Int(instance), color)
}

/*
Gets a specific instance's color multiplier.
*/
func (self Go) GetInstanceColor(instance int) gd.Color {
	return gd.Color(class(self).GetInstanceColor(gd.Int(instance)))
}

/*
Sets custom data for a specific instance. [param custom_data] is a [Color] type only to contain 4 floating-point numbers.
[b]Note:[/b] Each number is stored in 32 bits in the Forward+ and Mobile rendering methods, but is packed into 16 bits in the Compatibility rendering method.
For the custom data to be used, ensure that [member use_custom_data] is [code]true[/code].
This custom instance data has to be manually accessed in your custom shader using [code]INSTANCE_CUSTOM[/code].
*/
func (self Go) SetInstanceCustomData(instance int, custom_data gd.Color) {
	class(self).SetInstanceCustomData(gd.Int(instance), custom_data)
}

/*
Returns the custom data that has been set for a specific instance.
*/
func (self Go) GetInstanceCustomData(instance int) gd.Color {
	return gd.Color(class(self).GetInstanceCustomData(gd.Int(instance)))
}

/*
Returns the visibility axis-aligned bounding box in local space.
*/
func (self Go) GetAabb() gd.AABB {
	return gd.AABB(class(self).GetAabb())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.MultiMesh
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MultiMesh"))
	return Go{classdb.MultiMesh(object)}
}

func (self Go) TransformFormat() classdb.MultiMeshTransformFormat {
		return classdb.MultiMeshTransformFormat(class(self).GetTransformFormat())
}

func (self Go) SetTransformFormat(value classdb.MultiMeshTransformFormat) {
	class(self).SetTransformFormat(value)
}

func (self Go) UseColors() bool {
		return bool(class(self).IsUsingColors())
}

func (self Go) SetUseColors(value bool) {
	class(self).SetUseColors(value)
}

func (self Go) UseCustomData() bool {
		return bool(class(self).IsUsingCustomData())
}

func (self Go) SetUseCustomData(value bool) {
	class(self).SetUseCustomData(value)
}

func (self Go) CustomAabb() gd.AABB {
		return gd.AABB(class(self).GetCustomAabb())
}

func (self Go) SetCustomAabb(value gd.AABB) {
	class(self).SetCustomAabb(value)
}

func (self Go) InstanceCount() int {
		return int(int(class(self).GetInstanceCount()))
}

func (self Go) SetInstanceCount(value int) {
	class(self).SetInstanceCount(gd.Int(value))
}

func (self Go) VisibleInstanceCount() int {
		return int(int(class(self).GetVisibleInstanceCount()))
}

func (self Go) SetVisibleInstanceCount(value int) {
	class(self).SetVisibleInstanceCount(gd.Int(value))
}

func (self Go) Mesh() gdclass.Mesh {
		return gdclass.Mesh(class(self).GetMesh())
}

func (self Go) SetMesh(value gdclass.Mesh) {
	class(self).SetMesh(value)
}

func (self Go) Buffer() []float32 {
		return []float32(class(self).GetBuffer().AsSlice())
}

func (self Go) SetBuffer(value []float32) {
	class(self).SetBuffer(gd.NewPackedFloat32Slice(value))
}

//go:nosplit
func (self class) SetMesh(mesh gdclass.Mesh)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(mesh[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMesh() gdclass.Mesh {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_get_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Mesh{classdb.Mesh(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseColors(enable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_use_colors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingColors() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_is_using_colors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUseCustomData(enable bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_use_custom_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingCustomData() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_is_using_custom_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTransformFormat(format classdb.MultiMeshTransformFormat)  {
	var frame = callframe.New()
	callframe.Arg(frame, format)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_transform_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTransformFormat() classdb.MultiMeshTransformFormat {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.MultiMeshTransformFormat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_get_transform_format, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetInstanceCount(count gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_instance_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInstanceCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_get_instance_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVisibleInstanceCount(count gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_visible_instance_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVisibleInstanceCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_get_visible_instance_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [Transform3D] for a specific instance.
*/
//go:nosplit
func (self class) SetInstanceTransform(instance gd.Int, transform gd.Transform3D)  {
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_instance_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [Transform2D] for a specific instance.
*/
//go:nosplit
func (self class) SetInstanceTransform2d(instance gd.Int, transform gd.Transform2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_instance_transform_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [Transform3D] of a specific instance.
*/
//go:nosplit
func (self class) GetInstanceTransform(instance gd.Int) gd.Transform3D {
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_get_instance_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [Transform2D] of a specific instance.
*/
//go:nosplit
func (self class) GetInstanceTransform2d(instance gd.Int) gd.Transform2D {
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_get_instance_transform_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the color of a specific instance by [i]multiplying[/i] the mesh's existing vertex colors. This allows for different color tinting per instance.
[b]Note:[/b] Each component is stored in 32 bits in the Forward+ and Mobile rendering methods, but is packed into 16 bits in the Compatibility rendering method.
For the color to take effect, ensure that [member use_colors] is [code]true[/code] on the [MultiMesh] and [member BaseMaterial3D.vertex_color_use_as_albedo] is [code]true[/code] on the material. If you intend to set an absolute color instead of tinting, make sure the material's albedo color is set to pure white ([code]Color(1, 1, 1)[/code]).
*/
//go:nosplit
func (self class) SetInstanceColor(instance gd.Int, color gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_instance_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Gets a specific instance's color multiplier.
*/
//go:nosplit
func (self class) GetInstanceColor(instance gd.Int) gd.Color {
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_get_instance_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets custom data for a specific instance. [param custom_data] is a [Color] type only to contain 4 floating-point numbers.
[b]Note:[/b] Each number is stored in 32 bits in the Forward+ and Mobile rendering methods, but is packed into 16 bits in the Compatibility rendering method.
For the custom data to be used, ensure that [member use_custom_data] is [code]true[/code].
This custom instance data has to be manually accessed in your custom shader using [code]INSTANCE_CUSTOM[/code].
*/
//go:nosplit
func (self class) SetInstanceCustomData(instance gd.Int, custom_data gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	callframe.Arg(frame, custom_data)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_instance_custom_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the custom data that has been set for a specific instance.
*/
//go:nosplit
func (self class) GetInstanceCustomData(instance gd.Int) gd.Color {
	var frame = callframe.New()
	callframe.Arg(frame, instance)
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_get_instance_custom_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCustomAabb(aabb gd.AABB)  {
	var frame = callframe.New()
	callframe.Arg(frame, aabb)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_custom_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCustomAabb() gd.AABB {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.AABB](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_get_custom_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the visibility axis-aligned bounding box in local space.
*/
//go:nosplit
func (self class) GetAabb() gd.AABB {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.AABB](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_get_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetBuffer() gd.PackedFloat32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_get_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedFloat32Array](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBuffer(buffer gd.PackedFloat32Array)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(buffer))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiMesh.Bind_set_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsMultiMesh() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsMultiMesh() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("MultiMesh", func(ptr gd.Object) any { return classdb.MultiMesh(ptr) })}
type TransformFormat = classdb.MultiMeshTransformFormat

const (
/*Use this when using 2D transforms.*/
	Transform2d TransformFormat = 0
/*Use this when using 3D transforms.*/
	Transform3d TransformFormat = 1
)
