package MeshInstance3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/GeometryInstance3D"
import "grow.graphics/gd/gdclass/VisualInstance3D"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
MeshInstance3D is a node that takes a [Mesh] resource and adds it to the current scenario by creating an instance of it. This is the class most often used render 3D geometry and can be used to instance a single [Mesh] in many places. This allows reusing geometry, which can save on resources. When a [Mesh] has to be instantiated more than thousands of times at close proximity, consider using a [MultiMesh] in a [MultiMeshInstance3D] instead.

*/
type Go [1]classdb.MeshInstance3D

/*
Returns the internal [SkinReference] containing the skeleton's [RID] attached to this RID. See also [method Resource.get_rid], [method SkinReference.get_skeleton], and [method RenderingServer.instance_attach_skeleton].
*/
func (self Go) GetSkinReference() gdclass.SkinReference {
	return gdclass.SkinReference(class(self).GetSkinReference())
}

/*
Returns the number of surface override materials. This is equivalent to [method Mesh.get_surface_count]. See also [method get_surface_override_material].
*/
func (self Go) GetSurfaceOverrideMaterialCount() int {
	return int(int(class(self).GetSurfaceOverrideMaterialCount()))
}

/*
Sets the override [param material] for the specified [param surface] of the [Mesh] resource. This material is associated with this [MeshInstance3D] rather than with [member mesh].
[b]Note:[/b] This assigns the [Material] associated to the [MeshInstance3D]'s Surface Material Override properties, not the material within the [Mesh] resource. To set the material within the [Mesh] resource, use [method Mesh.surface_get_material] instead.
*/
func (self Go) SetSurfaceOverrideMaterial(surface int, material gdclass.Material) {
	class(self).SetSurfaceOverrideMaterial(gd.Int(surface), material)
}

/*
Returns the override [Material] for the specified [param surface] of the [Mesh] resource. See also [method get_surface_override_material_count].
[b]Note:[/b] This returns the [Material] associated to the [MeshInstance3D]'s Surface Material Override properties, not the material within the [Mesh] resource. To get the material within the [Mesh] resource, use [method Mesh.surface_get_material] instead.
*/
func (self Go) GetSurfaceOverrideMaterial(surface int) gdclass.Material {
	return gdclass.Material(class(self).GetSurfaceOverrideMaterial(gd.Int(surface)))
}

/*
Returns the [Material] that will be used by the [Mesh] when drawing. This can return the [member GeometryInstance3D.material_override], the surface override [Material] defined in this [MeshInstance3D], or the surface [Material] defined in the [member mesh]. For example, if [member GeometryInstance3D.material_override] is used, all surfaces will return the override material.
Returns [code]null[/code] if no material is active, including when [member mesh] is [code]null[/code].
*/
func (self Go) GetActiveMaterial(surface int) gdclass.Material {
	return gdclass.Material(class(self).GetActiveMaterial(gd.Int(surface)))
}

/*
This helper creates a [StaticBody3D] child node with a [ConcavePolygonShape3D] collision shape calculated from the mesh geometry. It's mainly used for testing.
*/
func (self Go) CreateTrimeshCollision() {
	class(self).CreateTrimeshCollision()
}

/*
This helper creates a [StaticBody3D] child node with a [ConvexPolygonShape3D] collision shape calculated from the mesh geometry. It's mainly used for testing.
If [param clean] is [code]true[/code] (default), duplicate and interior vertices are removed automatically. You can set it to [code]false[/code] to make the process faster if not needed.
If [param simplify] is [code]true[/code], the geometry can be further simplified to reduce the number of vertices. Disabled by default.
*/
func (self Go) CreateConvexCollision() {
	class(self).CreateConvexCollision(true, false)
}

/*
This helper creates a [StaticBody3D] child node with multiple [ConvexPolygonShape3D] collision shapes calculated from the mesh geometry via convex decomposition. The convex decomposition operation can be controlled with parameters from the optional [param settings].
*/
func (self Go) CreateMultipleConvexCollisions() {
	class(self).CreateMultipleConvexCollisions(([1]gdclass.MeshConvexDecompositionSettings{}[0]))
}

/*
Returns the number of blend shapes available. Produces an error if [member mesh] is [code]null[/code].
*/
func (self Go) GetBlendShapeCount() int {
	return int(int(class(self).GetBlendShapeCount()))
}

/*
Returns the index of the blend shape with the given [param name]. Returns [code]-1[/code] if no blend shape with this name exists, including when [member mesh] is [code]null[/code].
*/
func (self Go) FindBlendShapeByName(name string) int {
	return int(int(class(self).FindBlendShapeByName(gd.NewStringName(name))))
}

/*
Returns the value of the blend shape at the given [param blend_shape_idx]. Returns [code]0.0[/code] and produces an error if [member mesh] is [code]null[/code] or doesn't have a blend shape at that index.
*/
func (self Go) GetBlendShapeValue(blend_shape_idx int) float64 {
	return float64(float64(class(self).GetBlendShapeValue(gd.Int(blend_shape_idx))))
}

/*
Sets the value of the blend shape at [param blend_shape_idx] to [param value]. Produces an error if [member mesh] is [code]null[/code] or doesn't have a blend shape at that index.
*/
func (self Go) SetBlendShapeValue(blend_shape_idx int, value float64) {
	class(self).SetBlendShapeValue(gd.Int(blend_shape_idx), gd.Float(value))
}

/*
This helper creates a [MeshInstance3D] child node with gizmos at every vertex calculated from the mesh geometry. It's mainly used for testing.
*/
func (self Go) CreateDebugTangents() {
	class(self).CreateDebugTangents()
}

/*
Takes a snapshot from the current [ArrayMesh] with all blend shapes applied according to their current weights and bakes it to the provided [param existing] mesh. If no [param existing] mesh is provided a new [ArrayMesh] is created, baked and returned. Mesh surface materials are not copied.
[b]Performance:[/b] [Mesh] data needs to be received from the GPU, stalling the [RenderingServer] in the process.
*/
func (self Go) BakeMeshFromCurrentBlendShapeMix() gdclass.ArrayMesh {
	return gdclass.ArrayMesh(class(self).BakeMeshFromCurrentBlendShapeMix(([1]gdclass.ArrayMesh{}[0])))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.MeshInstance3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MeshInstance3D"))
	return Go{classdb.MeshInstance3D(object)}
}

func (self Go) Mesh() gdclass.Mesh {
		return gdclass.Mesh(class(self).GetMesh())
}

func (self Go) SetMesh(value gdclass.Mesh) {
	class(self).SetMesh(value)
}

func (self Go) Skin() gdclass.Skin {
		return gdclass.Skin(class(self).GetSkin())
}

func (self Go) SetSkin(value gdclass.Skin) {
	class(self).SetSkin(value)
}

func (self Go) Skeleton() string {
		return string(class(self).GetSkeletonPath().String())
}

func (self Go) SetSkeleton(value string) {
	class(self).SetSkeletonPath(gd.NewString(value).NodePath())
}

//go:nosplit
func (self class) SetMesh(mesh gdclass.Mesh)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(mesh[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_set_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMesh() gdclass.Mesh {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_get_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Mesh{classdb.Mesh(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSkeletonPath(skeleton_path gd.NodePath)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(skeleton_path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_set_skeleton_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSkeletonPath() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_get_skeleton_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSkin(skin gdclass.Skin)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(skin[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_set_skin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSkin() gdclass.Skin {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_get_skin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Skin{classdb.Skin(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Returns the internal [SkinReference] containing the skeleton's [RID] attached to this RID. See also [method Resource.get_rid], [method SkinReference.get_skeleton], and [method RenderingServer.instance_attach_skeleton].
*/
//go:nosplit
func (self class) GetSkinReference() gdclass.SkinReference {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_get_skin_reference, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.SkinReference{classdb.SkinReference(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Returns the number of surface override materials. This is equivalent to [method Mesh.get_surface_count]. See also [method get_surface_override_material].
*/
//go:nosplit
func (self class) GetSurfaceOverrideMaterialCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_get_surface_override_material_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the override [param material] for the specified [param surface] of the [Mesh] resource. This material is associated with this [MeshInstance3D] rather than with [member mesh].
[b]Note:[/b] This assigns the [Material] associated to the [MeshInstance3D]'s Surface Material Override properties, not the material within the [Mesh] resource. To set the material within the [Mesh] resource, use [method Mesh.surface_get_material] instead.
*/
//go:nosplit
func (self class) SetSurfaceOverrideMaterial(surface gd.Int, material gdclass.Material)  {
	var frame = callframe.New()
	callframe.Arg(frame, surface)
	callframe.Arg(frame, discreet.Get(material[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_set_surface_override_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the override [Material] for the specified [param surface] of the [Mesh] resource. See also [method get_surface_override_material_count].
[b]Note:[/b] This returns the [Material] associated to the [MeshInstance3D]'s Surface Material Override properties, not the material within the [Mesh] resource. To get the material within the [Mesh] resource, use [method Mesh.surface_get_material] instead.
*/
//go:nosplit
func (self class) GetSurfaceOverrideMaterial(surface gd.Int) gdclass.Material {
	var frame = callframe.New()
	callframe.Arg(frame, surface)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_get_surface_override_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Material{classdb.Material(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Returns the [Material] that will be used by the [Mesh] when drawing. This can return the [member GeometryInstance3D.material_override], the surface override [Material] defined in this [MeshInstance3D], or the surface [Material] defined in the [member mesh]. For example, if [member GeometryInstance3D.material_override] is used, all surfaces will return the override material.
Returns [code]null[/code] if no material is active, including when [member mesh] is [code]null[/code].
*/
//go:nosplit
func (self class) GetActiveMaterial(surface gd.Int) gdclass.Material {
	var frame = callframe.New()
	callframe.Arg(frame, surface)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_get_active_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Material{classdb.Material(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
This helper creates a [StaticBody3D] child node with a [ConcavePolygonShape3D] collision shape calculated from the mesh geometry. It's mainly used for testing.
*/
//go:nosplit
func (self class) CreateTrimeshCollision()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_create_trimesh_collision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
This helper creates a [StaticBody3D] child node with a [ConvexPolygonShape3D] collision shape calculated from the mesh geometry. It's mainly used for testing.
If [param clean] is [code]true[/code] (default), duplicate and interior vertices are removed automatically. You can set it to [code]false[/code] to make the process faster if not needed.
If [param simplify] is [code]true[/code], the geometry can be further simplified to reduce the number of vertices. Disabled by default.
*/
//go:nosplit
func (self class) CreateConvexCollision(clean bool, simplify bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, clean)
	callframe.Arg(frame, simplify)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_create_convex_collision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
This helper creates a [StaticBody3D] child node with multiple [ConvexPolygonShape3D] collision shapes calculated from the mesh geometry via convex decomposition. The convex decomposition operation can be controlled with parameters from the optional [param settings].
*/
//go:nosplit
func (self class) CreateMultipleConvexCollisions(settings gdclass.MeshConvexDecompositionSettings)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(settings[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_create_multiple_convex_collisions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of blend shapes available. Produces an error if [member mesh] is [code]null[/code].
*/
//go:nosplit
func (self class) GetBlendShapeCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_get_blend_shape_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the index of the blend shape with the given [param name]. Returns [code]-1[/code] if no blend shape with this name exists, including when [member mesh] is [code]null[/code].
*/
//go:nosplit
func (self class) FindBlendShapeByName(name gd.StringName) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(name))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_find_blend_shape_by_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the value of the blend shape at the given [param blend_shape_idx]. Returns [code]0.0[/code] and produces an error if [member mesh] is [code]null[/code] or doesn't have a blend shape at that index.
*/
//go:nosplit
func (self class) GetBlendShapeValue(blend_shape_idx gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, blend_shape_idx)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_get_blend_shape_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the value of the blend shape at [param blend_shape_idx] to [param value]. Produces an error if [member mesh] is [code]null[/code] or doesn't have a blend shape at that index.
*/
//go:nosplit
func (self class) SetBlendShapeValue(blend_shape_idx gd.Int, value gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, blend_shape_idx)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_set_blend_shape_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
This helper creates a [MeshInstance3D] child node with gizmos at every vertex calculated from the mesh geometry. It's mainly used for testing.
*/
//go:nosplit
func (self class) CreateDebugTangents()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_create_debug_tangents, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Takes a snapshot from the current [ArrayMesh] with all blend shapes applied according to their current weights and bakes it to the provided [param existing] mesh. If no [param existing] mesh is provided a new [ArrayMesh] is created, baked and returned. Mesh surface materials are not copied.
[b]Performance:[/b] [Mesh] data needs to be received from the GPU, stalling the [RenderingServer] in the process.
*/
//go:nosplit
func (self class) BakeMeshFromCurrentBlendShapeMix(existing gdclass.ArrayMesh) gdclass.ArrayMesh {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(existing[0])[0])
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshInstance3D.Bind_bake_mesh_from_current_blend_shape_mix, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.ArrayMesh{classdb.ArrayMesh(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self class) AsMeshInstance3D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsMeshInstance3D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsGeometryInstance3D() GeometryInstance3D.GD { return *((*GeometryInstance3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsGeometryInstance3D() GeometryInstance3D.Go { return *((*GeometryInstance3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsVisualInstance3D() VisualInstance3D.GD { return *((*VisualInstance3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsVisualInstance3D() VisualInstance3D.Go { return *((*VisualInstance3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode3D() Node3D.GD { return *((*Node3D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode3D() Node3D.Go { return *((*Node3D.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsGeometryInstance3D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsGeometryInstance3D(), name)
	}
}
func init() {classdb.Register("MeshInstance3D", func(ptr gd.Object) any { return classdb.MeshInstance3D(ptr) })}
