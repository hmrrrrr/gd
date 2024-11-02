package Mesh

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Resource"
import "grow.graphics/gd/variant/Array"
import "grow.graphics/gd/variant/Dictionary"
import "grow.graphics/gd/variant/AABB"
import "grow.graphics/gd/variant/Vector2i"
import "grow.graphics/gd/variant/Vector3"
import "grow.graphics/gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Mesh is a type of [Resource] that contains vertex array-based geometry, divided in [i]surfaces[/i]. Each surface contains a completely separate array and a material used to draw it. Design wise, a mesh with multiple surfaces is preferred to a single surface, because objects created in 3D editing software commonly contain multiple materials. The maximum number of surfaces per mesh is [constant RenderingServer.MAX_MESH_SURFACES].

	// Mesh methods that can be overridden by a [Class] that extends it.
	type Mesh interface {
		//Virtual method to override the surface count for a custom class extending [Mesh].
		GetSurfaceCount() int
		//Virtual method to override the surface array length for a custom class extending [Mesh].
		SurfaceGetArrayLen(index int) int
		//Virtual method to override the surface array index length for a custom class extending [Mesh].
		SurfaceGetArrayIndexLen(index int) int
		//Virtual method to override the surface arrays for a custom class extending [Mesh].
		SurfaceGetArrays(index int) Array.Any
		//Virtual method to override the blend shape arrays for a custom class extending [Mesh].
		SurfaceGetBlendShapeArrays(index int) gd.Array
		//Virtual method to override the surface LODs for a custom class extending [Mesh].
		SurfaceGetLods(index int) Dictionary.Any
		//Virtual method to override the surface format for a custom class extending [Mesh].
		SurfaceGetFormat(index int) int
		//Virtual method to override the surface primitive type for a custom class extending [Mesh].
		SurfaceGetPrimitiveType(index int) int
		//Virtual method to override the setting of a [param material] at the given [param index] for a custom class extending [Mesh].
		SurfaceSetMaterial(index int, material objects.Material)
		//Virtual method to override the surface material for a custom class extending [Mesh].
		SurfaceGetMaterial(index int) objects.Material
		//Virtual method to override the number of blend shapes for a custom class extending [Mesh].
		GetBlendShapeCount() int
		//Virtual method to override the retrieval of blend shape names for a custom class extending [Mesh].
		GetBlendShapeName(index int) string
		//Virtual method to override the names of blend shapes for a custom class extending [Mesh].
		SetBlendShapeName(index int, name string)
		//Virtual method to override the [AABB] for a custom class extending [Mesh].
		GetAabb() AABB.PositionSize
	}
*/
type Instance [1]classdb.Mesh

/*
Virtual method to override the surface count for a custom class extending [Mesh].
*/
func (Instance) _get_surface_count(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Virtual method to override the surface array length for a custom class extending [Mesh].
*/
func (Instance) _surface_get_array_len(impl func(ptr unsafe.Pointer, index int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Virtual method to override the surface array index length for a custom class extending [Mesh].
*/
func (Instance) _surface_get_array_index_len(impl func(ptr unsafe.Pointer, index int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Virtual method to override the surface arrays for a custom class extending [Mesh].
*/
func (Instance) _surface_get_arrays(impl func(ptr unsafe.Pointer, index int) Array.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to override the blend shape arrays for a custom class extending [Mesh].
*/
func (Instance) _surface_get_blend_shape_arrays(impl func(ptr unsafe.Pointer, index int) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to override the surface LODs for a custom class extending [Mesh].
*/
func (Instance) _surface_get_lods(impl func(ptr unsafe.Pointer, index int) Dictionary.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to override the surface format for a custom class extending [Mesh].
*/
func (Instance) _surface_get_format(impl func(ptr unsafe.Pointer, index int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Virtual method to override the surface primitive type for a custom class extending [Mesh].
*/
func (Instance) _surface_get_primitive_type(impl func(ptr unsafe.Pointer, index int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Virtual method to override the setting of a [param material] at the given [param index] for a custom class extending [Mesh].
*/
func (Instance) _surface_set_material(impl func(ptr unsafe.Pointer, index int, material objects.Material)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		var material = objects.Material{pointers.New[classdb.Material]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
		defer pointers.End(material[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(index), material)
	}
}

/*
Virtual method to override the surface material for a custom class extending [Mesh].
*/
func (Instance) _surface_get_material(impl func(ptr unsafe.Pointer, index int) objects.Material) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to override the number of blend shapes for a custom class extending [Mesh].
*/
func (Instance) _get_blend_shape_count(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Virtual method to override the retrieval of blend shape names for a custom class extending [Mesh].
*/
func (Instance) _get_blend_shape_name(impl func(ptr unsafe.Pointer, index int) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		ptr, ok := pointers.End(gd.NewStringName(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to override the names of blend shapes for a custom class extending [Mesh].
*/
func (Instance) _set_blend_shape_name(impl func(ptr unsafe.Pointer, index int, name string)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		var name = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 1))
		defer pointers.End(name)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(index), name.String())
	}
}

/*
Virtual method to override the [AABB] for a custom class extending [Mesh].
*/
func (Instance) _get_aabb(impl func(ptr unsafe.Pointer) AABB.PositionSize) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.AABB(ret))
	}
}

/*
Returns the smallest [AABB] enclosing this mesh in local space. Not affected by [code]custom_aabb[/code].
[b]Note:[/b] This is only implemented for [ArrayMesh] and [PrimitiveMesh].
*/
func (self Instance) GetAabb() AABB.PositionSize {
	return AABB.PositionSize(class(self).GetAabb())
}

/*
Returns all the vertices that make up the faces of the mesh. Each three vertices represent one triangle.
*/
func (self Instance) GetFaces() []Vector3.XYZ {
	return []Vector3.XYZ(class(self).GetFaces().AsSlice())
}

/*
Returns the number of surfaces that the [Mesh] holds. This is equivalent to [method MeshInstance3D.get_surface_override_material_count].
*/
func (self Instance) GetSurfaceCount() int {
	return int(int(class(self).GetSurfaceCount()))
}

/*
Returns the arrays for the vertices, normals, UVs, etc. that make up the requested surface (see [method ArrayMesh.add_surface_from_arrays]).
*/
func (self Instance) SurfaceGetArrays(surf_idx int) Array.Any {
	return Array.Any(class(self).SurfaceGetArrays(gd.Int(surf_idx)))
}

/*
Returns the blend shape arrays for the requested surface.
*/
func (self Instance) SurfaceGetBlendShapeArrays(surf_idx int) gd.Array {
	return gd.Array(class(self).SurfaceGetBlendShapeArrays(gd.Int(surf_idx)))
}

/*
Sets a [Material] for a given surface. Surface will be rendered using this material.
[b]Note:[/b] This assigns the material within the [Mesh] resource, not the [Material] associated to the [MeshInstance3D]'s Surface Material Override properties. To set the [Material] associated to the [MeshInstance3D]'s Surface Material Override properties, use [method MeshInstance3D.set_surface_override_material] instead.
*/
func (self Instance) SurfaceSetMaterial(surf_idx int, material objects.Material) {
	class(self).SurfaceSetMaterial(gd.Int(surf_idx), material)
}

/*
Returns a [Material] in a given surface. Surface is rendered using this material.
[b]Note:[/b] This returns the material within the [Mesh] resource, not the [Material] associated to the [MeshInstance3D]'s Surface Material Override properties. To get the [Material] associated to the [MeshInstance3D]'s Surface Material Override properties, use [method MeshInstance3D.get_surface_override_material] instead.
*/
func (self Instance) SurfaceGetMaterial(surf_idx int) objects.Material {
	return objects.Material(class(self).SurfaceGetMaterial(gd.Int(surf_idx)))
}

/*
Creates a placeholder version of this resource ([PlaceholderMesh]).
*/
func (self Instance) CreatePlaceholder() objects.Resource {
	return objects.Resource(class(self).CreatePlaceholder())
}

/*
Calculate a [ConcavePolygonShape3D] from the mesh.
*/
func (self Instance) CreateTrimeshShape() objects.ConcavePolygonShape3D {
	return objects.ConcavePolygonShape3D(class(self).CreateTrimeshShape())
}

/*
Calculate a [ConvexPolygonShape3D] from the mesh.
If [param clean] is [code]true[/code] (default), duplicate and interior vertices are removed automatically. You can set it to [code]false[/code] to make the process faster if not needed.
If [param simplify] is [code]true[/code], the geometry can be further simplified to reduce the number of vertices. Disabled by default.
*/
func (self Instance) CreateConvexShape() objects.ConvexPolygonShape3D {
	return objects.ConvexPolygonShape3D(class(self).CreateConvexShape(true, false))
}

/*
Calculate an outline mesh at a defined offset (margin) from the original mesh.
[b]Note:[/b] This method typically returns the vertices in reverse order (e.g. clockwise to counterclockwise).
*/
func (self Instance) CreateOutline(margin Float.X) objects.Mesh {
	return objects.Mesh(class(self).CreateOutline(gd.Float(margin)))
}

/*
Generate a [TriangleMesh] from the mesh. Considers only surfaces using one of these primitive types: [constant PRIMITIVE_TRIANGLES], [constant PRIMITIVE_TRIANGLE_STRIP].
*/
func (self Instance) GenerateTriangleMesh() objects.TriangleMesh {
	return objects.TriangleMesh(class(self).GenerateTriangleMesh())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Mesh

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Mesh"))
	return Instance{classdb.Mesh(object)}
}

func (self Instance) LightmapSizeHint() Vector2i.XY {
	return Vector2i.XY(class(self).GetLightmapSizeHint())
}

func (self Instance) SetLightmapSizeHint(value Vector2i.XY) {
	class(self).SetLightmapSizeHint(gd.Vector2i(value))
}

/*
Virtual method to override the surface count for a custom class extending [Mesh].
*/
func (class) _get_surface_count(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Virtual method to override the surface array length for a custom class extending [Mesh].
*/
func (class) _surface_get_array_len(impl func(ptr unsafe.Pointer, index gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Virtual method to override the surface array index length for a custom class extending [Mesh].
*/
func (class) _surface_get_array_index_len(impl func(ptr unsafe.Pointer, index gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Virtual method to override the surface arrays for a custom class extending [Mesh].
*/
func (class) _surface_get_arrays(impl func(ptr unsafe.Pointer, index gd.Int) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to override the blend shape arrays for a custom class extending [Mesh].
*/
func (class) _surface_get_blend_shape_arrays(impl func(ptr unsafe.Pointer, index gd.Int) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to override the surface LODs for a custom class extending [Mesh].
*/
func (class) _surface_get_lods(impl func(ptr unsafe.Pointer, index gd.Int) gd.Dictionary) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to override the surface format for a custom class extending [Mesh].
*/
func (class) _surface_get_format(impl func(ptr unsafe.Pointer, index gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Virtual method to override the surface primitive type for a custom class extending [Mesh].
*/
func (class) _surface_get_primitive_type(impl func(ptr unsafe.Pointer, index gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Virtual method to override the setting of a [param material] at the given [param index] for a custom class extending [Mesh].
*/
func (class) _surface_set_material(impl func(ptr unsafe.Pointer, index gd.Int, material objects.Material)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		var material = objects.Material{pointers.New[classdb.Material]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
		defer pointers.End(material[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, index, material)
	}
}

/*
Virtual method to override the surface material for a custom class extending [Mesh].
*/
func (class) _surface_get_material(impl func(ptr unsafe.Pointer, index gd.Int) objects.Material) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to override the number of blend shapes for a custom class extending [Mesh].
*/
func (class) _get_blend_shape_count(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Virtual method to override the retrieval of blend shape names for a custom class extending [Mesh].
*/
func (class) _get_blend_shape_name(impl func(ptr unsafe.Pointer, index gd.Int) gd.StringName) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method to override the names of blend shapes for a custom class extending [Mesh].
*/
func (class) _set_blend_shape_name(impl func(ptr unsafe.Pointer, index gd.Int, name gd.StringName)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		var name = pointers.New[gd.StringName](gd.UnsafeGet[[1]uintptr](p_args, 1))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, index, name)
	}
}

/*
Virtual method to override the [AABB] for a custom class extending [Mesh].
*/
func (class) _get_aabb(impl func(ptr unsafe.Pointer) gd.AABB) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

//go:nosplit
func (self class) SetLightmapSizeHint(size gd.Vector2i) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mesh.Bind_set_lightmap_size_hint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLightmapSizeHint() gd.Vector2i {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mesh.Bind_get_lightmap_size_hint, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the smallest [AABB] enclosing this mesh in local space. Not affected by [code]custom_aabb[/code].
[b]Note:[/b] This is only implemented for [ArrayMesh] and [PrimitiveMesh].
*/
//go:nosplit
func (self class) GetAabb() gd.AABB {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.AABB](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mesh.Bind_get_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns all the vertices that make up the faces of the mesh. Each three vertices represent one triangle.
*/
//go:nosplit
func (self class) GetFaces() gd.PackedVector3Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mesh.Bind_get_faces, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedVector3Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the number of surfaces that the [Mesh] holds. This is equivalent to [method MeshInstance3D.get_surface_override_material_count].
*/
//go:nosplit
func (self class) GetSurfaceCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mesh.Bind_get_surface_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the arrays for the vertices, normals, UVs, etc. that make up the requested surface (see [method ArrayMesh.add_surface_from_arrays]).
*/
//go:nosplit
func (self class) SurfaceGetArrays(surf_idx gd.Int) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, surf_idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mesh.Bind_surface_get_arrays, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the blend shape arrays for the requested surface.
*/
//go:nosplit
func (self class) SurfaceGetBlendShapeArrays(surf_idx gd.Int) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, surf_idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mesh.Bind_surface_get_blend_shape_arrays, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets a [Material] for a given surface. Surface will be rendered using this material.
[b]Note:[/b] This assigns the material within the [Mesh] resource, not the [Material] associated to the [MeshInstance3D]'s Surface Material Override properties. To set the [Material] associated to the [MeshInstance3D]'s Surface Material Override properties, use [method MeshInstance3D.set_surface_override_material] instead.
*/
//go:nosplit
func (self class) SurfaceSetMaterial(surf_idx gd.Int, material objects.Material) {
	var frame = callframe.New()
	callframe.Arg(frame, surf_idx)
	callframe.Arg(frame, pointers.Get(material[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mesh.Bind_surface_set_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a [Material] in a given surface. Surface is rendered using this material.
[b]Note:[/b] This returns the material within the [Mesh] resource, not the [Material] associated to the [MeshInstance3D]'s Surface Material Override properties. To get the [Material] associated to the [MeshInstance3D]'s Surface Material Override properties, use [method MeshInstance3D.get_surface_override_material] instead.
*/
//go:nosplit
func (self class) SurfaceGetMaterial(surf_idx gd.Int) objects.Material {
	var frame = callframe.New()
	callframe.Arg(frame, surf_idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mesh.Bind_surface_get_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Material{classdb.Material(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Creates a placeholder version of this resource ([PlaceholderMesh]).
*/
//go:nosplit
func (self class) CreatePlaceholder() objects.Resource {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mesh.Bind_create_placeholder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Resource{classdb.Resource(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Calculate a [ConcavePolygonShape3D] from the mesh.
*/
//go:nosplit
func (self class) CreateTrimeshShape() objects.ConcavePolygonShape3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mesh.Bind_create_trimesh_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.ConcavePolygonShape3D{classdb.ConcavePolygonShape3D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Calculate a [ConvexPolygonShape3D] from the mesh.
If [param clean] is [code]true[/code] (default), duplicate and interior vertices are removed automatically. You can set it to [code]false[/code] to make the process faster if not needed.
If [param simplify] is [code]true[/code], the geometry can be further simplified to reduce the number of vertices. Disabled by default.
*/
//go:nosplit
func (self class) CreateConvexShape(clean bool, simplify bool) objects.ConvexPolygonShape3D {
	var frame = callframe.New()
	callframe.Arg(frame, clean)
	callframe.Arg(frame, simplify)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mesh.Bind_create_convex_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.ConvexPolygonShape3D{classdb.ConvexPolygonShape3D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Calculate an outline mesh at a defined offset (margin) from the original mesh.
[b]Note:[/b] This method typically returns the vertices in reverse order (e.g. clockwise to counterclockwise).
*/
//go:nosplit
func (self class) CreateOutline(margin gd.Float) objects.Mesh {
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mesh.Bind_create_outline, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Mesh{classdb.Mesh(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Generate a [TriangleMesh] from the mesh. Considers only surfaces using one of these primitive types: [constant PRIMITIVE_TRIANGLES], [constant PRIMITIVE_TRIANGLE_STRIP].
*/
//go:nosplit
func (self class) GenerateTriangleMesh() objects.TriangleMesh {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mesh.Bind_generate_triangle_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.TriangleMesh{classdb.TriangleMesh(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self class) AsMesh() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMesh() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_surface_count":
		return reflect.ValueOf(self._get_surface_count)
	case "_surface_get_array_len":
		return reflect.ValueOf(self._surface_get_array_len)
	case "_surface_get_array_index_len":
		return reflect.ValueOf(self._surface_get_array_index_len)
	case "_surface_get_arrays":
		return reflect.ValueOf(self._surface_get_arrays)
	case "_surface_get_blend_shape_arrays":
		return reflect.ValueOf(self._surface_get_blend_shape_arrays)
	case "_surface_get_lods":
		return reflect.ValueOf(self._surface_get_lods)
	case "_surface_get_format":
		return reflect.ValueOf(self._surface_get_format)
	case "_surface_get_primitive_type":
		return reflect.ValueOf(self._surface_get_primitive_type)
	case "_surface_set_material":
		return reflect.ValueOf(self._surface_set_material)
	case "_surface_get_material":
		return reflect.ValueOf(self._surface_get_material)
	case "_get_blend_shape_count":
		return reflect.ValueOf(self._get_blend_shape_count)
	case "_get_blend_shape_name":
		return reflect.ValueOf(self._get_blend_shape_name)
	case "_set_blend_shape_name":
		return reflect.ValueOf(self._set_blend_shape_name)
	case "_get_aabb":
		return reflect.ValueOf(self._get_aabb)
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_surface_count":
		return reflect.ValueOf(self._get_surface_count)
	case "_surface_get_array_len":
		return reflect.ValueOf(self._surface_get_array_len)
	case "_surface_get_array_index_len":
		return reflect.ValueOf(self._surface_get_array_index_len)
	case "_surface_get_arrays":
		return reflect.ValueOf(self._surface_get_arrays)
	case "_surface_get_blend_shape_arrays":
		return reflect.ValueOf(self._surface_get_blend_shape_arrays)
	case "_surface_get_lods":
		return reflect.ValueOf(self._surface_get_lods)
	case "_surface_get_format":
		return reflect.ValueOf(self._surface_get_format)
	case "_surface_get_primitive_type":
		return reflect.ValueOf(self._surface_get_primitive_type)
	case "_surface_set_material":
		return reflect.ValueOf(self._surface_set_material)
	case "_surface_get_material":
		return reflect.ValueOf(self._surface_get_material)
	case "_get_blend_shape_count":
		return reflect.ValueOf(self._get_blend_shape_count)
	case "_get_blend_shape_name":
		return reflect.ValueOf(self._get_blend_shape_name)
	case "_set_blend_shape_name":
		return reflect.ValueOf(self._set_blend_shape_name)
	case "_get_aabb":
		return reflect.ValueOf(self._get_aabb)
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() { classdb.Register("Mesh", func(ptr gd.Object) any { return classdb.Mesh(ptr) }) }

type PrimitiveType = classdb.MeshPrimitiveType

const (
	/*Render array as points (one vertex equals one point).*/
	PrimitivePoints PrimitiveType = 0
	/*Render array as lines (every two vertices a line is created).*/
	PrimitiveLines PrimitiveType = 1
	/*Render array as line strip.*/
	PrimitiveLineStrip PrimitiveType = 2
	/*Render array as triangles (every three vertices a triangle is created).*/
	PrimitiveTriangles PrimitiveType = 3
	/*Render array as triangle strips.*/
	PrimitiveTriangleStrip PrimitiveType = 4
)

type ArrayType = classdb.MeshArrayType

const (
	/*[PackedVector3Array], [PackedVector2Array], or [Array] of vertex positions.*/
	ArrayVertex ArrayType = 0
	/*[PackedVector3Array] of vertex normals.
	  [b]Note:[/b] The array has to consist of normal vectors, otherwise they will be normalized by the engine, potentially causing visual discrepancies.*/
	ArrayNormal ArrayType = 1
	/*[PackedFloat32Array] of vertex tangents. Each element in groups of 4 floats, first 3 floats determine the tangent, and the last the binormal direction as -1 or 1.*/
	ArrayTangent ArrayType = 2
	/*[PackedColorArray] of vertex colors.*/
	ArrayColor ArrayType = 3
	/*[PackedVector2Array] for UV coordinates.*/
	ArrayTexUv ArrayType = 4
	/*[PackedVector2Array] for second UV coordinates.*/
	ArrayTexUv2 ArrayType = 5
	/*Contains custom color channel 0. [PackedByteArray] if [code](format >> Mesh.ARRAY_FORMAT_CUSTOM0_SHIFT) & Mesh.ARRAY_FORMAT_CUSTOM_MASK[/code] is [constant ARRAY_CUSTOM_RGBA8_UNORM], [constant ARRAY_CUSTOM_RGBA8_SNORM], [constant ARRAY_CUSTOM_RG_HALF], or [constant ARRAY_CUSTOM_RGBA_HALF]. [PackedFloat32Array] otherwise.*/
	ArrayCustom0 ArrayType = 6
	/*Contains custom color channel 1. [PackedByteArray] if [code](format >> Mesh.ARRAY_FORMAT_CUSTOM1_SHIFT) & Mesh.ARRAY_FORMAT_CUSTOM_MASK[/code] is [constant ARRAY_CUSTOM_RGBA8_UNORM], [constant ARRAY_CUSTOM_RGBA8_SNORM], [constant ARRAY_CUSTOM_RG_HALF], or [constant ARRAY_CUSTOM_RGBA_HALF]. [PackedFloat32Array] otherwise.*/
	ArrayCustom1 ArrayType = 7
	/*Contains custom color channel 2. [PackedByteArray] if [code](format >> Mesh.ARRAY_FORMAT_CUSTOM2_SHIFT) & Mesh.ARRAY_FORMAT_CUSTOM_MASK[/code] is [constant ARRAY_CUSTOM_RGBA8_UNORM], [constant ARRAY_CUSTOM_RGBA8_SNORM], [constant ARRAY_CUSTOM_RG_HALF], or [constant ARRAY_CUSTOM_RGBA_HALF]. [PackedFloat32Array] otherwise.*/
	ArrayCustom2 ArrayType = 8
	/*Contains custom color channel 3. [PackedByteArray] if [code](format >> Mesh.ARRAY_FORMAT_CUSTOM3_SHIFT) & Mesh.ARRAY_FORMAT_CUSTOM_MASK[/code] is [constant ARRAY_CUSTOM_RGBA8_UNORM], [constant ARRAY_CUSTOM_RGBA8_SNORM], [constant ARRAY_CUSTOM_RG_HALF], or [constant ARRAY_CUSTOM_RGBA_HALF]. [PackedFloat32Array] otherwise.*/
	ArrayCustom3 ArrayType = 9
	/*[PackedFloat32Array] or [PackedInt32Array] of bone indices. Contains either 4 or 8 numbers per vertex depending on the presence of the [constant ARRAY_FLAG_USE_8_BONE_WEIGHTS] flag.*/
	ArrayBones ArrayType = 10
	/*[PackedFloat32Array] or [PackedFloat64Array] of bone weights in the range [code]0.0[/code] to [code]1.0[/code] (inclusive). Contains either 4 or 8 numbers per vertex depending on the presence of the [constant ARRAY_FLAG_USE_8_BONE_WEIGHTS] flag.*/
	ArrayWeights ArrayType = 11
	/*[PackedInt32Array] of integers used as indices referencing vertices, colors, normals, tangents, and textures. All of those arrays must have the same number of elements as the vertex array. No index can be beyond the vertex array size. When this index array is present, it puts the function into "index mode," where the index selects the [i]i[/i]'th vertex, normal, tangent, color, UV, etc. This means if you want to have different normals or colors along an edge, you have to duplicate the vertices.
	  For triangles, the index array is interpreted as triples, referring to the vertices of each triangle. For lines, the index array is in pairs indicating the start and end of each line.*/
	ArrayIndex ArrayType = 12
	/*Represents the size of the [enum ArrayType] enum.*/
	ArrayMax ArrayType = 13
)

type ArrayCustomFormat = classdb.MeshArrayCustomFormat

const (
	/*Indicates this custom channel contains unsigned normalized byte colors from 0 to 1, encoded as [PackedByteArray].*/
	ArrayCustomRgba8Unorm ArrayCustomFormat = 0
	/*Indicates this custom channel contains signed normalized byte colors from -1 to 1, encoded as [PackedByteArray].*/
	ArrayCustomRgba8Snorm ArrayCustomFormat = 1
	/*Indicates this custom channel contains half precision float colors, encoded as [PackedByteArray]. Only red and green channels are used.*/
	ArrayCustomRgHalf ArrayCustomFormat = 2
	/*Indicates this custom channel contains half precision float colors, encoded as [PackedByteArray].*/
	ArrayCustomRgbaHalf ArrayCustomFormat = 3
	/*Indicates this custom channel contains full float colors, in a [PackedFloat32Array]. Only the red channel is used.*/
	ArrayCustomRFloat ArrayCustomFormat = 4
	/*Indicates this custom channel contains full float colors, in a [PackedFloat32Array]. Only red and green channels are used.*/
	ArrayCustomRgFloat ArrayCustomFormat = 5
	/*Indicates this custom channel contains full float colors, in a [PackedFloat32Array]. Only red, green and blue channels are used.*/
	ArrayCustomRgbFloat ArrayCustomFormat = 6
	/*Indicates this custom channel contains full float colors, in a [PackedFloat32Array].*/
	ArrayCustomRgbaFloat ArrayCustomFormat = 7
	/*Represents the size of the [enum ArrayCustomFormat] enum.*/
	ArrayCustomMax ArrayCustomFormat = 8
)

type ArrayFormat = classdb.MeshArrayFormat

const (
	/*Mesh array contains vertices. All meshes require a vertex array so this should always be present.*/
	ArrayFormatVertex ArrayFormat = 1
	/*Mesh array contains normals.*/
	ArrayFormatNormal ArrayFormat = 2
	/*Mesh array contains tangents.*/
	ArrayFormatTangent ArrayFormat = 4
	/*Mesh array contains colors.*/
	ArrayFormatColor ArrayFormat = 8
	/*Mesh array contains UVs.*/
	ArrayFormatTexUv ArrayFormat = 16
	/*Mesh array contains second UV.*/
	ArrayFormatTexUv2 ArrayFormat = 32
	/*Mesh array contains custom channel index 0.*/
	ArrayFormatCustom0 ArrayFormat = 64
	/*Mesh array contains custom channel index 1.*/
	ArrayFormatCustom1 ArrayFormat = 128
	/*Mesh array contains custom channel index 2.*/
	ArrayFormatCustom2 ArrayFormat = 256
	/*Mesh array contains custom channel index 3.*/
	ArrayFormatCustom3 ArrayFormat = 512
	/*Mesh array contains bones.*/
	ArrayFormatBones ArrayFormat = 1024
	/*Mesh array contains bone weights.*/
	ArrayFormatWeights ArrayFormat = 2048
	/*Mesh array uses indices.*/
	ArrayFormatIndex ArrayFormat = 4096
	/*Mask of mesh channels permitted in blend shapes.*/
	ArrayFormatBlendShapeMask ArrayFormat = 7
	/*Shift of first custom channel.*/
	ArrayFormatCustomBase ArrayFormat = 13
	/*Number of format bits per custom channel. See [enum ArrayCustomFormat].*/
	ArrayFormatCustomBits ArrayFormat = 3
	/*Amount to shift [enum ArrayCustomFormat] for custom channel index 0.*/
	ArrayFormatCustom0Shift ArrayFormat = 13
	/*Amount to shift [enum ArrayCustomFormat] for custom channel index 1.*/
	ArrayFormatCustom1Shift ArrayFormat = 16
	/*Amount to shift [enum ArrayCustomFormat] for custom channel index 2.*/
	ArrayFormatCustom2Shift ArrayFormat = 19
	/*Amount to shift [enum ArrayCustomFormat] for custom channel index 3.*/
	ArrayFormatCustom3Shift ArrayFormat = 22
	/*Mask of custom format bits per custom channel. Must be shifted by one of the SHIFT constants. See [enum ArrayCustomFormat].*/
	ArrayFormatCustomMask ArrayFormat = 7
	/*Shift of first compress flag. Compress flags should be passed to [method ArrayMesh.add_surface_from_arrays] and [method SurfaceTool.commit].*/
	ArrayCompressFlagsBase ArrayFormat = 25
	/*Flag used to mark that the array contains 2D vertices.*/
	ArrayFlagUse2dVertices ArrayFormat = 33554432
	/*Flag indices that the mesh data will use [code]GL_DYNAMIC_DRAW[/code] on GLES. Unused on Vulkan.*/
	ArrayFlagUseDynamicUpdate ArrayFormat = 67108864
	/*Flag used to mark that the mesh contains up to 8 bone influences per vertex. This flag indicates that [constant ARRAY_BONES] and [constant ARRAY_WEIGHTS] elements will have double length.*/
	ArrayFlagUse8BoneWeights ArrayFormat = 134217728
	/*Flag used to mark that the mesh intentionally contains no vertex array.*/
	ArrayFlagUsesEmptyVertexArray ArrayFormat = 268435456
	/*Flag used to mark that a mesh is using compressed attributes (vertices, normals, tangents, UVs). When this form of compression is enabled, vertex positions will be packed into an RGBA16UNORM attribute and scaled in the vertex shader. The normal and tangent will be packed into an RG16UNORM representing an axis, and a 16-bit float stored in the A-channel of the vertex. UVs will use 16-bit normalized floats instead of full 32-bit signed floats. When using this compression mode you must use either vertices, normals, and tangents or only vertices. You cannot use normals without tangents. Importers will automatically enable this compression if they can.*/
	ArrayFlagCompressAttributes ArrayFormat = 536870912
)

type BlendShapeMode = classdb.MeshBlendShapeMode

const (
	/*Blend shapes are normalized.*/
	BlendShapeModeNormalized BlendShapeMode = 0
	/*Blend shapes are relative to base weight.*/
	BlendShapeModeRelative BlendShapeMode = 1
)
