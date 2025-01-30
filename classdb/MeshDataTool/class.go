// Package MeshDataTool provides methods for working with MeshDataTool object instances.
package MeshDataTool

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
import "graphics.gd/variant/Color"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/Plane"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Vector2"
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
MeshDataTool provides access to individual vertices in a [Mesh]. It allows users to read and edit vertex data of meshes. It also creates an array of faces and edges.
To use MeshDataTool, load a mesh with [method create_from_surface]. When you are finished editing the data commit the data to a mesh with [method commit_to_surface].
Below is an example of how MeshDataTool may be used.
[codeblocks]
[gdscript]
var mesh = ArrayMesh.new()
mesh.add_surface_from_arrays(Mesh.PRIMITIVE_TRIANGLES, BoxMesh.new().get_mesh_arrays())
var mdt = MeshDataTool.new()
mdt.create_from_surface(mesh, 0)
for i in range(mdt.get_vertex_count()):

	var vertex = mdt.get_vertex(i)
	# In this example we extend the mesh by one unit, which results in separated faces as it is flat shaded.
	vertex += mdt.get_vertex_normal(i)
	# Save your change.
	mdt.set_vertex(i, vertex)

mesh.clear_surfaces()
mdt.commit_to_surface(mesh)
var mi = MeshInstance.new()
mi.mesh = mesh
add_child(mi)
[/gdscript]
[csharp]
var mesh = new ArrayMesh();
mesh.AddSurfaceFromArrays(Mesh.PrimitiveType.Triangles, new BoxMesh().GetMeshArrays());
var mdt = new MeshDataTool();
mdt.CreateFromSurface(mesh, 0);
for (var i = 0; i < mdt.GetVertexCount(); i++)

	{
	    Vector3 vertex = mdt.GetVertex(i);
	    // In this example we extend the mesh by one unit, which results in separated faces as it is flat shaded.
	    vertex += mdt.GetVertexNormal(i);
	    // Save your change.
	    mdt.SetVertex(i, vertex);
	}

mesh.ClearSurfaces();
mdt.CommitToSurface(mesh);
var mi = new MeshInstance();
mi.Mesh = mesh;
AddChild(mi);
[/csharp]
[/codeblocks]
See also [ArrayMesh], [ImmediateMesh] and [SurfaceTool] for procedural geometry generation.
[b]Note:[/b] Godot uses clockwise [url=https://learnopengl.com/Advanced-OpenGL/Face-culling]winding order[/url] for front faces of triangle primitive modes.
*/
type Instance [1]gdclass.MeshDataTool

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsMeshDataTool() Instance
}

/*
Clears all data currently in MeshDataTool.
*/
func (self Instance) Clear() { //gd:MeshDataTool.clear
	class(self).Clear()
}

/*
Uses specified surface of given [Mesh] to populate data for MeshDataTool.
Requires [Mesh] with primitive type [constant Mesh.PRIMITIVE_TRIANGLES].
*/
func (self Instance) CreateFromSurface(mesh [1]gdclass.ArrayMesh, surface int) error { //gd:MeshDataTool.create_from_surface
	return error(gd.ToError(class(self).CreateFromSurface(mesh, int64(surface))))
}

/*
Adds a new surface to specified [Mesh] with edited data.
*/
func (self Instance) CommitToSurface(mesh [1]gdclass.ArrayMesh) error { //gd:MeshDataTool.commit_to_surface
	return error(gd.ToError(class(self).CommitToSurface(mesh, int64(0))))
}

/*
Returns the [Mesh]'s format as a combination of the [enum Mesh.ArrayFormat] flags. For example, a mesh containing both vertices and normals would return a format of [code]3[/code] because [constant Mesh.ARRAY_FORMAT_VERTEX] is [code]1[/code] and [constant Mesh.ARRAY_FORMAT_NORMAL] is [code]2[/code].
*/
func (self Instance) GetFormat() int { //gd:MeshDataTool.get_format
	return int(int(class(self).GetFormat()))
}

/*
Returns the total number of vertices in [Mesh].
*/
func (self Instance) GetVertexCount() int { //gd:MeshDataTool.get_vertex_count
	return int(int(class(self).GetVertexCount()))
}

/*
Returns the number of edges in this [Mesh].
*/
func (self Instance) GetEdgeCount() int { //gd:MeshDataTool.get_edge_count
	return int(int(class(self).GetEdgeCount()))
}

/*
Returns the number of faces in this [Mesh].
*/
func (self Instance) GetFaceCount() int { //gd:MeshDataTool.get_face_count
	return int(int(class(self).GetFaceCount()))
}

/*
Sets the position of the given vertex.
*/
func (self Instance) SetVertex(idx int, vertex Vector3.XYZ) { //gd:MeshDataTool.set_vertex
	class(self).SetVertex(int64(idx), Vector3.XYZ(vertex))
}

/*
Returns the position of the given vertex.
*/
func (self Instance) GetVertex(idx int) Vector3.XYZ { //gd:MeshDataTool.get_vertex
	return Vector3.XYZ(class(self).GetVertex(int64(idx)))
}

/*
Sets the normal of the given vertex.
*/
func (self Instance) SetVertexNormal(idx int, normal Vector3.XYZ) { //gd:MeshDataTool.set_vertex_normal
	class(self).SetVertexNormal(int64(idx), Vector3.XYZ(normal))
}

/*
Returns the normal of the given vertex.
*/
func (self Instance) GetVertexNormal(idx int) Vector3.XYZ { //gd:MeshDataTool.get_vertex_normal
	return Vector3.XYZ(class(self).GetVertexNormal(int64(idx)))
}

/*
Sets the tangent of the given vertex.
*/
func (self Instance) SetVertexTangent(idx int, tangent Plane.NormalD) { //gd:MeshDataTool.set_vertex_tangent
	class(self).SetVertexTangent(int64(idx), Plane.NormalD(tangent))
}

/*
Returns the tangent of the given vertex.
*/
func (self Instance) GetVertexTangent(idx int) Plane.NormalD { //gd:MeshDataTool.get_vertex_tangent
	return Plane.NormalD(class(self).GetVertexTangent(int64(idx)))
}

/*
Sets the UV of the given vertex.
*/
func (self Instance) SetVertexUv(idx int, uv Vector2.XY) { //gd:MeshDataTool.set_vertex_uv
	class(self).SetVertexUv(int64(idx), Vector2.XY(uv))
}

/*
Returns the UV of the given vertex.
*/
func (self Instance) GetVertexUv(idx int) Vector2.XY { //gd:MeshDataTool.get_vertex_uv
	return Vector2.XY(class(self).GetVertexUv(int64(idx)))
}

/*
Sets the UV2 of the given vertex.
*/
func (self Instance) SetVertexUv2(idx int, uv2 Vector2.XY) { //gd:MeshDataTool.set_vertex_uv2
	class(self).SetVertexUv2(int64(idx), Vector2.XY(uv2))
}

/*
Returns the UV2 of the given vertex.
*/
func (self Instance) GetVertexUv2(idx int) Vector2.XY { //gd:MeshDataTool.get_vertex_uv2
	return Vector2.XY(class(self).GetVertexUv2(int64(idx)))
}

/*
Sets the color of the given vertex.
*/
func (self Instance) SetVertexColor(idx int, color Color.RGBA) { //gd:MeshDataTool.set_vertex_color
	class(self).SetVertexColor(int64(idx), Color.RGBA(color))
}

/*
Returns the color of the given vertex.
*/
func (self Instance) GetVertexColor(idx int) Color.RGBA { //gd:MeshDataTool.get_vertex_color
	return Color.RGBA(class(self).GetVertexColor(int64(idx)))
}

/*
Sets the bones of the given vertex.
*/
func (self Instance) SetVertexBones(idx int, bones []int32) { //gd:MeshDataTool.set_vertex_bones
	class(self).SetVertexBones(int64(idx), Packed.New(bones...))
}

/*
Returns the bones of the given vertex.
*/
func (self Instance) GetVertexBones(idx int) []int32 { //gd:MeshDataTool.get_vertex_bones
	return []int32(slices.Collect(class(self).GetVertexBones(int64(idx)).Values()))
}

/*
Sets the bone weights of the given vertex.
*/
func (self Instance) SetVertexWeights(idx int, weights []float32) { //gd:MeshDataTool.set_vertex_weights
	class(self).SetVertexWeights(int64(idx), Packed.New(weights...))
}

/*
Returns bone weights of the given vertex.
*/
func (self Instance) GetVertexWeights(idx int) []float32 { //gd:MeshDataTool.get_vertex_weights
	return []float32(slices.Collect(class(self).GetVertexWeights(int64(idx)).Values()))
}

/*
Sets the metadata associated with the given vertex.
*/
func (self Instance) SetVertexMeta(idx int, meta any) { //gd:MeshDataTool.set_vertex_meta
	class(self).SetVertexMeta(int64(idx), variant.New(meta))
}

/*
Returns the metadata associated with the given vertex.
*/
func (self Instance) GetVertexMeta(idx int) any { //gd:MeshDataTool.get_vertex_meta
	return any(class(self).GetVertexMeta(int64(idx)).Interface())
}

/*
Returns an array of edges that share the given vertex.
*/
func (self Instance) GetVertexEdges(idx int) []int32 { //gd:MeshDataTool.get_vertex_edges
	return []int32(slices.Collect(class(self).GetVertexEdges(int64(idx)).Values()))
}

/*
Returns an array of faces that share the given vertex.
*/
func (self Instance) GetVertexFaces(idx int) []int32 { //gd:MeshDataTool.get_vertex_faces
	return []int32(slices.Collect(class(self).GetVertexFaces(int64(idx)).Values()))
}

/*
Returns index of specified vertex connected to given edge.
Vertex argument can only be 0 or 1 because edges are comprised of two vertices.
*/
func (self Instance) GetEdgeVertex(idx int, vertex int) int { //gd:MeshDataTool.get_edge_vertex
	return int(int(class(self).GetEdgeVertex(int64(idx), int64(vertex))))
}

/*
Returns array of faces that touch given edge.
*/
func (self Instance) GetEdgeFaces(idx int) []int32 { //gd:MeshDataTool.get_edge_faces
	return []int32(slices.Collect(class(self).GetEdgeFaces(int64(idx)).Values()))
}

/*
Sets the metadata of the given edge.
*/
func (self Instance) SetEdgeMeta(idx int, meta any) { //gd:MeshDataTool.set_edge_meta
	class(self).SetEdgeMeta(int64(idx), variant.New(meta))
}

/*
Returns meta information assigned to given edge.
*/
func (self Instance) GetEdgeMeta(idx int) any { //gd:MeshDataTool.get_edge_meta
	return any(class(self).GetEdgeMeta(int64(idx)).Interface())
}

/*
Returns the specified vertex index of the given face.
Vertex argument must be either 0, 1, or 2 because faces contain three vertices.
[b]Example:[/b]
[codeblocks]
[gdscript]
var index = mesh_data_tool.get_face_vertex(0, 1) # Gets the index of the second vertex of the first face.
var position = mesh_data_tool.get_vertex(index)
var normal = mesh_data_tool.get_vertex_normal(index)
[/gdscript]
[csharp]
int index = meshDataTool.GetFaceVertex(0, 1); // Gets the index of the second vertex of the first face.
Vector3 position = meshDataTool.GetVertex(index);
Vector3 normal = meshDataTool.GetVertexNormal(index);
[/csharp]
[/codeblocks]
*/
func (self Instance) GetFaceVertex(idx int, vertex int) int { //gd:MeshDataTool.get_face_vertex
	return int(int(class(self).GetFaceVertex(int64(idx), int64(vertex))))
}

/*
Returns specified edge associated with given face.
Edge argument must be either 0, 1, or 2 because a face only has three edges.
*/
func (self Instance) GetFaceEdge(idx int, edge int) int { //gd:MeshDataTool.get_face_edge
	return int(int(class(self).GetFaceEdge(int64(idx), int64(edge))))
}

/*
Sets the metadata of the given face.
*/
func (self Instance) SetFaceMeta(idx int, meta any) { //gd:MeshDataTool.set_face_meta
	class(self).SetFaceMeta(int64(idx), variant.New(meta))
}

/*
Returns the metadata associated with the given face.
*/
func (self Instance) GetFaceMeta(idx int) any { //gd:MeshDataTool.get_face_meta
	return any(class(self).GetFaceMeta(int64(idx)).Interface())
}

/*
Calculates and returns the face normal of the given face.
*/
func (self Instance) GetFaceNormal(idx int) Vector3.XYZ { //gd:MeshDataTool.get_face_normal
	return Vector3.XYZ(class(self).GetFaceNormal(int64(idx)))
}

/*
Sets the material to be used by newly-constructed [Mesh].
*/
func (self Instance) SetMaterial(material [1]gdclass.Material) { //gd:MeshDataTool.set_material
	class(self).SetMaterial(material)
}

/*
Returns the material assigned to the [Mesh].
*/
func (self Instance) GetMaterial() [1]gdclass.Material { //gd:MeshDataTool.get_material
	return [1]gdclass.Material(class(self).GetMaterial())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.MeshDataTool

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MeshDataTool"))
	casted := Instance{*(*gdclass.MeshDataTool)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Clears all data currently in MeshDataTool.
*/
//go:nosplit
func (self class) Clear() { //gd:MeshDataTool.clear
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Uses specified surface of given [Mesh] to populate data for MeshDataTool.
Requires [Mesh] with primitive type [constant Mesh.PRIMITIVE_TRIANGLES].
*/
//go:nosplit
func (self class) CreateFromSurface(mesh [1]gdclass.ArrayMesh, surface int64) Error.Code { //gd:MeshDataTool.create_from_surface
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(mesh[0])[0])
	callframe.Arg(frame, surface)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_create_from_surface, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Adds a new surface to specified [Mesh] with edited data.
*/
//go:nosplit
func (self class) CommitToSurface(mesh [1]gdclass.ArrayMesh, compression_flags int64) Error.Code { //gd:MeshDataTool.commit_to_surface
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(mesh[0])[0])
	callframe.Arg(frame, compression_flags)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_commit_to_surface, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the [Mesh]'s format as a combination of the [enum Mesh.ArrayFormat] flags. For example, a mesh containing both vertices and normals would return a format of [code]3[/code] because [constant Mesh.ARRAY_FORMAT_VERTEX] is [code]1[/code] and [constant Mesh.ARRAY_FORMAT_NORMAL] is [code]2[/code].
*/
//go:nosplit
func (self class) GetFormat() int64 { //gd:MeshDataTool.get_format
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_format, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the total number of vertices in [Mesh].
*/
//go:nosplit
func (self class) GetVertexCount() int64 { //gd:MeshDataTool.get_vertex_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_vertex_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of edges in this [Mesh].
*/
//go:nosplit
func (self class) GetEdgeCount() int64 { //gd:MeshDataTool.get_edge_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_edge_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of faces in this [Mesh].
*/
//go:nosplit
func (self class) GetFaceCount() int64 { //gd:MeshDataTool.get_face_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_face_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the position of the given vertex.
*/
//go:nosplit
func (self class) SetVertex(idx int64, vertex Vector3.XYZ) { //gd:MeshDataTool.set_vertex
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, vertex)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_set_vertex, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the position of the given vertex.
*/
//go:nosplit
func (self class) GetVertex(idx int64) Vector3.XYZ { //gd:MeshDataTool.get_vertex
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_vertex, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the normal of the given vertex.
*/
//go:nosplit
func (self class) SetVertexNormal(idx int64, normal Vector3.XYZ) { //gd:MeshDataTool.set_vertex_normal
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, normal)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_set_vertex_normal, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the normal of the given vertex.
*/
//go:nosplit
func (self class) GetVertexNormal(idx int64) Vector3.XYZ { //gd:MeshDataTool.get_vertex_normal
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_vertex_normal, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the tangent of the given vertex.
*/
//go:nosplit
func (self class) SetVertexTangent(idx int64, tangent Plane.NormalD) { //gd:MeshDataTool.set_vertex_tangent
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, tangent)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_set_vertex_tangent, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the tangent of the given vertex.
*/
//go:nosplit
func (self class) GetVertexTangent(idx int64) Plane.NormalD { //gd:MeshDataTool.get_vertex_tangent
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[Plane.NormalD](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_vertex_tangent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the UV of the given vertex.
*/
//go:nosplit
func (self class) SetVertexUv(idx int64, uv Vector2.XY) { //gd:MeshDataTool.set_vertex_uv
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, uv)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_set_vertex_uv, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the UV of the given vertex.
*/
//go:nosplit
func (self class) GetVertexUv(idx int64) Vector2.XY { //gd:MeshDataTool.get_vertex_uv
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_vertex_uv, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the UV2 of the given vertex.
*/
//go:nosplit
func (self class) SetVertexUv2(idx int64, uv2 Vector2.XY) { //gd:MeshDataTool.set_vertex_uv2
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, uv2)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_set_vertex_uv2, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the UV2 of the given vertex.
*/
//go:nosplit
func (self class) GetVertexUv2(idx int64) Vector2.XY { //gd:MeshDataTool.get_vertex_uv2
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_vertex_uv2, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the color of the given vertex.
*/
//go:nosplit
func (self class) SetVertexColor(idx int64, color Color.RGBA) { //gd:MeshDataTool.set_vertex_color
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_set_vertex_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the color of the given vertex.
*/
//go:nosplit
func (self class) GetVertexColor(idx int64) Color.RGBA { //gd:MeshDataTool.get_vertex_color
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_vertex_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the bones of the given vertex.
*/
//go:nosplit
func (self class) SetVertexBones(idx int64, bones Packed.Array[int32]) { //gd:MeshDataTool.set_vertex_bones
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedInt32Array, int32](bones)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_set_vertex_bones, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the bones of the given vertex.
*/
//go:nosplit
func (self class) GetVertexBones(idx int64) Packed.Array[int32] { //gd:MeshDataTool.get_vertex_bones
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_vertex_bones, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int32](Array.Through(gd.PackedProxy[gd.PackedInt32Array, int32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Sets the bone weights of the given vertex.
*/
//go:nosplit
func (self class) SetVertexWeights(idx int64, weights Packed.Array[float32]) { //gd:MeshDataTool.set_vertex_weights
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedFloat32Array, float32](weights)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_set_vertex_weights, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns bone weights of the given vertex.
*/
//go:nosplit
func (self class) GetVertexWeights(idx int64) Packed.Array[float32] { //gd:MeshDataTool.get_vertex_weights
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_vertex_weights, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[float32](Array.Through(gd.PackedProxy[gd.PackedFloat32Array, float32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Sets the metadata associated with the given vertex.
*/
//go:nosplit
func (self class) SetVertexMeta(idx int64, meta variant.Any) { //gd:MeshDataTool.set_vertex_meta
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(meta)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_set_vertex_meta, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the metadata associated with the given vertex.
*/
//go:nosplit
func (self class) GetVertexMeta(idx int64) variant.Any { //gd:MeshDataTool.get_vertex_meta
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_vertex_meta, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Through(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns an array of edges that share the given vertex.
*/
//go:nosplit
func (self class) GetVertexEdges(idx int64) Packed.Array[int32] { //gd:MeshDataTool.get_vertex_edges
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_vertex_edges, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int32](Array.Through(gd.PackedProxy[gd.PackedInt32Array, int32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns an array of faces that share the given vertex.
*/
//go:nosplit
func (self class) GetVertexFaces(idx int64) Packed.Array[int32] { //gd:MeshDataTool.get_vertex_faces
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_vertex_faces, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int32](Array.Through(gd.PackedProxy[gd.PackedInt32Array, int32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns index of specified vertex connected to given edge.
Vertex argument can only be 0 or 1 because edges are comprised of two vertices.
*/
//go:nosplit
func (self class) GetEdgeVertex(idx int64, vertex int64) int64 { //gd:MeshDataTool.get_edge_vertex
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, vertex)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_edge_vertex, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns array of faces that touch given edge.
*/
//go:nosplit
func (self class) GetEdgeFaces(idx int64) Packed.Array[int32] { //gd:MeshDataTool.get_edge_faces
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_edge_faces, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int32](Array.Through(gd.PackedProxy[gd.PackedInt32Array, int32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Sets the metadata of the given edge.
*/
//go:nosplit
func (self class) SetEdgeMeta(idx int64, meta variant.Any) { //gd:MeshDataTool.set_edge_meta
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(meta)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_set_edge_meta, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns meta information assigned to given edge.
*/
//go:nosplit
func (self class) GetEdgeMeta(idx int64) variant.Any { //gd:MeshDataTool.get_edge_meta
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_edge_meta, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Through(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the specified vertex index of the given face.
Vertex argument must be either 0, 1, or 2 because faces contain three vertices.
[b]Example:[/b]
[codeblocks]
[gdscript]
var index = mesh_data_tool.get_face_vertex(0, 1) # Gets the index of the second vertex of the first face.
var position = mesh_data_tool.get_vertex(index)
var normal = mesh_data_tool.get_vertex_normal(index)
[/gdscript]
[csharp]
int index = meshDataTool.GetFaceVertex(0, 1); // Gets the index of the second vertex of the first face.
Vector3 position = meshDataTool.GetVertex(index);
Vector3 normal = meshDataTool.GetVertexNormal(index);
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) GetFaceVertex(idx int64, vertex int64) int64 { //gd:MeshDataTool.get_face_vertex
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, vertex)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_face_vertex, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns specified edge associated with given face.
Edge argument must be either 0, 1, or 2 because a face only has three edges.
*/
//go:nosplit
func (self class) GetFaceEdge(idx int64, edge int64) int64 { //gd:MeshDataTool.get_face_edge
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, edge)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_face_edge, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the metadata of the given face.
*/
//go:nosplit
func (self class) SetFaceMeta(idx int64, meta variant.Any) { //gd:MeshDataTool.set_face_meta
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(meta)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_set_face_meta, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the metadata associated with the given face.
*/
//go:nosplit
func (self class) GetFaceMeta(idx int64) variant.Any { //gd:MeshDataTool.get_face_meta
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_face_meta, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Through(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Calculates and returns the face normal of the given face.
*/
//go:nosplit
func (self class) GetFaceNormal(idx int64) Vector3.XYZ { //gd:MeshDataTool.get_face_normal
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_face_normal, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the material to be used by newly-constructed [Mesh].
*/
//go:nosplit
func (self class) SetMaterial(material [1]gdclass.Material) { //gd:MeshDataTool.set_material
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(material[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_set_material, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the material assigned to the [Mesh].
*/
//go:nosplit
func (self class) GetMaterial() [1]gdclass.Material { //gd:MeshDataTool.get_material
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MeshDataTool.Bind_get_material, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Material{gd.PointerWithOwnershipTransferredToGo[gdclass.Material](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsMeshDataTool() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMeshDataTool() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("MeshDataTool", func(ptr gd.Object) any {
		return [1]gdclass.MeshDataTool{*(*gdclass.MeshDataTool)(unsafe.Pointer(&ptr))}
	})
}
