// Package MeshInstance3D provides methods for working with MeshInstance3D object instances.
package MeshInstance3D

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
import "graphics.gd/variant/Dictionary"
import "graphics.gd/classdb/GeometryInstance3D"
import "graphics.gd/classdb/VisualInstance3D"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/NodePath"
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
var _ Dictionary.Any

/*
MeshInstance3D is a node that takes a [Mesh] resource and adds it to the current scenario by creating an instance of it. This is the class most often used render 3D geometry and can be used to instance a single [Mesh] in many places. This allows reusing geometry, which can save on resources. When a [Mesh] has to be instantiated more than thousands of times at close proximity, consider using a [MultiMesh] in a [MultiMeshInstance3D] instead.
*/
type Instance [1]gdclass.MeshInstance3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsMeshInstance3D() Instance
}

/*
Returns the internal [SkinReference] containing the skeleton's [RID] attached to this RID. See also [method Resource.get_rid], [method SkinReference.get_skeleton], and [method RenderingServer.instance_attach_skeleton].
*/
func (self Instance) GetSkinReference() [1]gdclass.SkinReference { //gd:MeshInstance3D.get_skin_reference
	return [1]gdclass.SkinReference(class(self).GetSkinReference())
}

/*
Returns the number of surface override materials. This is equivalent to [method Mesh.get_surface_count]. See also [method get_surface_override_material].
*/
func (self Instance) GetSurfaceOverrideMaterialCount() int { //gd:MeshInstance3D.get_surface_override_material_count
	return int(int(class(self).GetSurfaceOverrideMaterialCount()))
}

/*
Sets the override [param material] for the specified [param surface] of the [Mesh] resource. This material is associated with this [MeshInstance3D] rather than with [member mesh].
[b]Note:[/b] This assigns the [Material] associated to the [MeshInstance3D]'s Surface Material Override properties, not the material within the [Mesh] resource. To set the material within the [Mesh] resource, use [method Mesh.surface_get_material] instead.
*/
func (self Instance) SetSurfaceOverrideMaterial(surface int, material [1]gdclass.Material) { //gd:MeshInstance3D.set_surface_override_material
	class(self).SetSurfaceOverrideMaterial(gd.Int(surface), material)
}

/*
Returns the override [Material] for the specified [param surface] of the [Mesh] resource. See also [method get_surface_override_material_count].
[b]Note:[/b] This returns the [Material] associated to the [MeshInstance3D]'s Surface Material Override properties, not the material within the [Mesh] resource. To get the material within the [Mesh] resource, use [method Mesh.surface_get_material] instead.
*/
func (self Instance) GetSurfaceOverrideMaterial(surface int) [1]gdclass.Material { //gd:MeshInstance3D.get_surface_override_material
	return [1]gdclass.Material(class(self).GetSurfaceOverrideMaterial(gd.Int(surface)))
}

/*
Returns the [Material] that will be used by the [Mesh] when drawing. This can return the [member GeometryInstance3D.material_override], the surface override [Material] defined in this [MeshInstance3D], or the surface [Material] defined in the [member mesh]. For example, if [member GeometryInstance3D.material_override] is used, all surfaces will return the override material.
Returns [code]null[/code] if no material is active, including when [member mesh] is [code]null[/code].
*/
func (self Instance) GetActiveMaterial(surface int) [1]gdclass.Material { //gd:MeshInstance3D.get_active_material
	return [1]gdclass.Material(class(self).GetActiveMaterial(gd.Int(surface)))
}

/*
This helper creates a [StaticBody3D] child node with a [ConcavePolygonShape3D] collision shape calculated from the mesh geometry. It's mainly used for testing.
*/
func (self Instance) CreateTrimeshCollision() { //gd:MeshInstance3D.create_trimesh_collision
	class(self).CreateTrimeshCollision()
}

/*
This helper creates a [StaticBody3D] child node with a [ConvexPolygonShape3D] collision shape calculated from the mesh geometry. It's mainly used for testing.
If [param clean] is [code]true[/code] (default), duplicate and interior vertices are removed automatically. You can set it to [code]false[/code] to make the process faster if not needed.
If [param simplify] is [code]true[/code], the geometry can be further simplified to reduce the number of vertices. Disabled by default.
*/
func (self Instance) CreateConvexCollision() { //gd:MeshInstance3D.create_convex_collision
	class(self).CreateConvexCollision(true, false)
}

/*
This helper creates a [StaticBody3D] child node with multiple [ConvexPolygonShape3D] collision shapes calculated from the mesh geometry via convex decomposition. The convex decomposition operation can be controlled with parameters from the optional [param settings].
*/
func (self Instance) CreateMultipleConvexCollisions() { //gd:MeshInstance3D.create_multiple_convex_collisions
	class(self).CreateMultipleConvexCollisions([1][1]gdclass.MeshConvexDecompositionSettings{}[0])
}

/*
Returns the number of blend shapes available. Produces an error if [member mesh] is [code]null[/code].
*/
func (self Instance) GetBlendShapeCount() int { //gd:MeshInstance3D.get_blend_shape_count
	return int(int(class(self).GetBlendShapeCount()))
}

/*
Returns the index of the blend shape with the given [param name]. Returns [code]-1[/code] if no blend shape with this name exists, including when [member mesh] is [code]null[/code].
*/
func (self Instance) FindBlendShapeByName(name string) int { //gd:MeshInstance3D.find_blend_shape_by_name
	return int(int(class(self).FindBlendShapeByName(gd.NewStringName(name))))
}

/*
Returns the value of the blend shape at the given [param blend_shape_idx]. Returns [code]0.0[/code] and produces an error if [member mesh] is [code]null[/code] or doesn't have a blend shape at that index.
*/
func (self Instance) GetBlendShapeValue(blend_shape_idx int) Float.X { //gd:MeshInstance3D.get_blend_shape_value
	return Float.X(Float.X(class(self).GetBlendShapeValue(gd.Int(blend_shape_idx))))
}

/*
Sets the value of the blend shape at [param blend_shape_idx] to [param value]. Produces an error if [member mesh] is [code]null[/code] or doesn't have a blend shape at that index.
*/
func (self Instance) SetBlendShapeValue(blend_shape_idx int, value Float.X) { //gd:MeshInstance3D.set_blend_shape_value
	class(self).SetBlendShapeValue(gd.Int(blend_shape_idx), gd.Float(value))
}

/*
This helper creates a [MeshInstance3D] child node with gizmos at every vertex calculated from the mesh geometry. It's mainly used for testing.
*/
func (self Instance) CreateDebugTangents() { //gd:MeshInstance3D.create_debug_tangents
	class(self).CreateDebugTangents()
}

/*
Takes a snapshot from the current [ArrayMesh] with all blend shapes applied according to their current weights and bakes it to the provided [param existing] mesh. If no [param existing] mesh is provided a new [ArrayMesh] is created, baked and returned. Mesh surface materials are not copied.
[b]Performance:[/b] [Mesh] data needs to be received from the GPU, stalling the [RenderingServer] in the process.
*/
func (self Instance) BakeMeshFromCurrentBlendShapeMix() [1]gdclass.ArrayMesh { //gd:MeshInstance3D.bake_mesh_from_current_blend_shape_mix
	return [1]gdclass.ArrayMesh(class(self).BakeMeshFromCurrentBlendShapeMix([1][1]gdclass.ArrayMesh{}[0]))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.MeshInstance3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MeshInstance3D"))
	casted := Instance{*(*gdclass.MeshInstance3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Mesh() [1]gdclass.Mesh {
	return [1]gdclass.Mesh(class(self).GetMesh())
}

func (self Instance) SetMesh(value [1]gdclass.Mesh) {
	class(self).SetMesh(value)
}

func (self Instance) Skin() [1]gdclass.Skin {
	return [1]gdclass.Skin(class(self).GetSkin())
}

func (self Instance) SetSkin(value [1]gdclass.Skin) {
	class(self).SetSkin(value)
}

func (self Instance) Skeleton() NodePath.String {
	return NodePath.String(class(self).GetSkeletonPath().String())
}

func (self Instance) SetSkeleton(value NodePath.String) {
	class(self).SetSkeletonPath(gd.NewString(string(value)).NodePath())
}

//go:nosplit
func (self class) SetMesh(mesh [1]gdclass.Mesh) { //gd:MeshInstance3D.set_mesh
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(mesh[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_set_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMesh() [1]gdclass.Mesh { //gd:MeshInstance3D.get_mesh
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_get_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Mesh{gd.PointerWithOwnershipTransferredToGo[gdclass.Mesh](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSkeletonPath(skeleton_path gd.NodePath) { //gd:MeshInstance3D.set_skeleton_path
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(skeleton_path))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_set_skeleton_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSkeletonPath() gd.NodePath { //gd:MeshInstance3D.get_skeleton_path
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_get_skeleton_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSkin(skin [1]gdclass.Skin) { //gd:MeshInstance3D.set_skin
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(skin[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_set_skin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSkin() [1]gdclass.Skin { //gd:MeshInstance3D.get_skin
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_get_skin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Skin{gd.PointerWithOwnershipTransferredToGo[gdclass.Skin](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the internal [SkinReference] containing the skeleton's [RID] attached to this RID. See also [method Resource.get_rid], [method SkinReference.get_skeleton], and [method RenderingServer.instance_attach_skeleton].
*/
//go:nosplit
func (self class) GetSkinReference() [1]gdclass.SkinReference { //gd:MeshInstance3D.get_skin_reference
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_get_skin_reference, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.SkinReference{gd.PointerWithOwnershipTransferredToGo[gdclass.SkinReference](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the number of surface override materials. This is equivalent to [method Mesh.get_surface_count]. See also [method get_surface_override_material].
*/
//go:nosplit
func (self class) GetSurfaceOverrideMaterialCount() gd.Int { //gd:MeshInstance3D.get_surface_override_material_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_get_surface_override_material_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the override [param material] for the specified [param surface] of the [Mesh] resource. This material is associated with this [MeshInstance3D] rather than with [member mesh].
[b]Note:[/b] This assigns the [Material] associated to the [MeshInstance3D]'s Surface Material Override properties, not the material within the [Mesh] resource. To set the material within the [Mesh] resource, use [method Mesh.surface_get_material] instead.
*/
//go:nosplit
func (self class) SetSurfaceOverrideMaterial(surface gd.Int, material [1]gdclass.Material) { //gd:MeshInstance3D.set_surface_override_material
	var frame = callframe.New()
	callframe.Arg(frame, surface)
	callframe.Arg(frame, pointers.Get(material[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_set_surface_override_material, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the override [Material] for the specified [param surface] of the [Mesh] resource. See also [method get_surface_override_material_count].
[b]Note:[/b] This returns the [Material] associated to the [MeshInstance3D]'s Surface Material Override properties, not the material within the [Mesh] resource. To get the material within the [Mesh] resource, use [method Mesh.surface_get_material] instead.
*/
//go:nosplit
func (self class) GetSurfaceOverrideMaterial(surface gd.Int) [1]gdclass.Material { //gd:MeshInstance3D.get_surface_override_material
	var frame = callframe.New()
	callframe.Arg(frame, surface)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_get_surface_override_material, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Material{gd.PointerWithOwnershipTransferredToGo[gdclass.Material](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the [Material] that will be used by the [Mesh] when drawing. This can return the [member GeometryInstance3D.material_override], the surface override [Material] defined in this [MeshInstance3D], or the surface [Material] defined in the [member mesh]. For example, if [member GeometryInstance3D.material_override] is used, all surfaces will return the override material.
Returns [code]null[/code] if no material is active, including when [member mesh] is [code]null[/code].
*/
//go:nosplit
func (self class) GetActiveMaterial(surface gd.Int) [1]gdclass.Material { //gd:MeshInstance3D.get_active_material
	var frame = callframe.New()
	callframe.Arg(frame, surface)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_get_active_material, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Material{gd.PointerWithOwnershipTransferredToGo[gdclass.Material](r_ret.Get())}
	frame.Free()
	return ret
}

/*
This helper creates a [StaticBody3D] child node with a [ConcavePolygonShape3D] collision shape calculated from the mesh geometry. It's mainly used for testing.
*/
//go:nosplit
func (self class) CreateTrimeshCollision() { //gd:MeshInstance3D.create_trimesh_collision
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_create_trimesh_collision, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
This helper creates a [StaticBody3D] child node with a [ConvexPolygonShape3D] collision shape calculated from the mesh geometry. It's mainly used for testing.
If [param clean] is [code]true[/code] (default), duplicate and interior vertices are removed automatically. You can set it to [code]false[/code] to make the process faster if not needed.
If [param simplify] is [code]true[/code], the geometry can be further simplified to reduce the number of vertices. Disabled by default.
*/
//go:nosplit
func (self class) CreateConvexCollision(clean bool, simplify bool) { //gd:MeshInstance3D.create_convex_collision
	var frame = callframe.New()
	callframe.Arg(frame, clean)
	callframe.Arg(frame, simplify)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_create_convex_collision, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
This helper creates a [StaticBody3D] child node with multiple [ConvexPolygonShape3D] collision shapes calculated from the mesh geometry via convex decomposition. The convex decomposition operation can be controlled with parameters from the optional [param settings].
*/
//go:nosplit
func (self class) CreateMultipleConvexCollisions(settings [1]gdclass.MeshConvexDecompositionSettings) { //gd:MeshInstance3D.create_multiple_convex_collisions
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(settings[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_create_multiple_convex_collisions, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the number of blend shapes available. Produces an error if [member mesh] is [code]null[/code].
*/
//go:nosplit
func (self class) GetBlendShapeCount() gd.Int { //gd:MeshInstance3D.get_blend_shape_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_get_blend_shape_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the index of the blend shape with the given [param name]. Returns [code]-1[/code] if no blend shape with this name exists, including when [member mesh] is [code]null[/code].
*/
//go:nosplit
func (self class) FindBlendShapeByName(name gd.StringName) gd.Int { //gd:MeshInstance3D.find_blend_shape_by_name
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_find_blend_shape_by_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the value of the blend shape at the given [param blend_shape_idx]. Returns [code]0.0[/code] and produces an error if [member mesh] is [code]null[/code] or doesn't have a blend shape at that index.
*/
//go:nosplit
func (self class) GetBlendShapeValue(blend_shape_idx gd.Int) gd.Float { //gd:MeshInstance3D.get_blend_shape_value
	var frame = callframe.New()
	callframe.Arg(frame, blend_shape_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_get_blend_shape_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the value of the blend shape at [param blend_shape_idx] to [param value]. Produces an error if [member mesh] is [code]null[/code] or doesn't have a blend shape at that index.
*/
//go:nosplit
func (self class) SetBlendShapeValue(blend_shape_idx gd.Int, value gd.Float) { //gd:MeshInstance3D.set_blend_shape_value
	var frame = callframe.New()
	callframe.Arg(frame, blend_shape_idx)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_set_blend_shape_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
This helper creates a [MeshInstance3D] child node with gizmos at every vertex calculated from the mesh geometry. It's mainly used for testing.
*/
//go:nosplit
func (self class) CreateDebugTangents() { //gd:MeshInstance3D.create_debug_tangents
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_create_debug_tangents, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Takes a snapshot from the current [ArrayMesh] with all blend shapes applied according to their current weights and bakes it to the provided [param existing] mesh. If no [param existing] mesh is provided a new [ArrayMesh] is created, baked and returned. Mesh surface materials are not copied.
[b]Performance:[/b] [Mesh] data needs to be received from the GPU, stalling the [RenderingServer] in the process.
*/
//go:nosplit
func (self class) BakeMeshFromCurrentBlendShapeMix(existing [1]gdclass.ArrayMesh) [1]gdclass.ArrayMesh { //gd:MeshInstance3D.bake_mesh_from_current_blend_shape_mix
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(existing[0])[0])
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_bake_mesh_from_current_blend_shape_mix, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.ArrayMesh{gd.PointerWithOwnershipTransferredToGo[gdclass.ArrayMesh](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsMeshInstance3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMeshInstance3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsGeometryInstance3D() GeometryInstance3D.Advanced {
	return *((*GeometryInstance3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsGeometryInstance3D() GeometryInstance3D.Instance {
	return *((*GeometryInstance3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualInstance3D() VisualInstance3D.Advanced {
	return *((*VisualInstance3D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisualInstance3D() VisualInstance3D.Instance {
	return *((*VisualInstance3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(GeometryInstance3D.Advanced(self.AsGeometryInstance3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(GeometryInstance3D.Instance(self.AsGeometryInstance3D()), name)
	}
}
func init() {
	gdclass.Register("MeshInstance3D", func(ptr gd.Object) any {
		return [1]gdclass.MeshInstance3D{*(*gdclass.MeshInstance3D)(unsafe.Pointer(&ptr))}
	})
}
