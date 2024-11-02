package Polygon2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Node2D"
import "grow.graphics/gd/objects/CanvasItem"
import "grow.graphics/gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A Polygon2D is defined by a set of points. Each point is connected to the next, with the final point being connected to the first, resulting in a closed polygon. Polygon2Ds can be filled with color (solid or gradient) or filled with a given texture.
*/
type Instance [1]classdb.Polygon2D

/*
Adds a bone with the specified [param path] and [param weights].
*/
func (self Instance) AddBone(path string, weights []float32) {
	class(self).AddBone(gd.NewString(path).NodePath(), gd.NewPackedFloat32Slice(weights))
}

/*
Returns the number of bones in this [Polygon2D].
*/
func (self Instance) GetBoneCount() int {
	return int(int(class(self).GetBoneCount()))
}

/*
Returns the path to the node associated with the specified bone.
*/
func (self Instance) GetBonePath(index int) string {
	return string(class(self).GetBonePath(gd.Int(index)).String())
}

/*
Returns the weight values of the specified bone.
*/
func (self Instance) GetBoneWeights(index int) []float32 {
	return []float32(class(self).GetBoneWeights(gd.Int(index)).AsSlice())
}

/*
Removes the specified bone from this [Polygon2D].
*/
func (self Instance) EraseBone(index int) {
	class(self).EraseBone(gd.Int(index))
}

/*
Removes all bones from this [Polygon2D].
*/
func (self Instance) ClearBones() {
	class(self).ClearBones()
}

/*
Sets the path to the node associated with the specified bone.
*/
func (self Instance) SetBonePath(index int, path string) {
	class(self).SetBonePath(gd.Int(index), gd.NewString(path).NodePath())
}

/*
Sets the weight values for the specified bone.
*/
func (self Instance) SetBoneWeights(index int, weights []float32) {
	class(self).SetBoneWeights(gd.Int(index), gd.NewPackedFloat32Slice(weights))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Polygon2D

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Polygon2D"))
	return Instance{classdb.Polygon2D(object)}
}

func (self Instance) Color() gd.Color {
	return gd.Color(class(self).GetColor())
}

func (self Instance) SetColor(value gd.Color) {
	class(self).SetColor(value)
}

func (self Instance) Offset() gd.Vector2 {
	return gd.Vector2(class(self).GetOffset())
}

func (self Instance) SetOffset(value gd.Vector2) {
	class(self).SetOffset(value)
}

func (self Instance) Antialiased() bool {
	return bool(class(self).GetAntialiased())
}

func (self Instance) SetAntialiased(value bool) {
	class(self).SetAntialiased(value)
}

func (self Instance) Texture() objects.Texture2D {
	return objects.Texture2D(class(self).GetTexture())
}

func (self Instance) SetTexture(value objects.Texture2D) {
	class(self).SetTexture(value)
}

func (self Instance) TextureOffset() gd.Vector2 {
	return gd.Vector2(class(self).GetTextureOffset())
}

func (self Instance) SetTextureOffset(value gd.Vector2) {
	class(self).SetTextureOffset(value)
}

func (self Instance) TextureScale() gd.Vector2 {
	return gd.Vector2(class(self).GetTextureScale())
}

func (self Instance) SetTextureScale(value gd.Vector2) {
	class(self).SetTextureScale(value)
}

func (self Instance) TextureRotation() float64 {
	return float64(float64(class(self).GetTextureRotation()))
}

func (self Instance) SetTextureRotation(value float64) {
	class(self).SetTextureRotation(gd.Float(value))
}

func (self Instance) Skeleton() string {
	return string(class(self).GetSkeleton().String())
}

func (self Instance) SetSkeleton(value string) {
	class(self).SetSkeleton(gd.NewString(value).NodePath())
}

func (self Instance) InvertEnabled() bool {
	return bool(class(self).GetInvertEnabled())
}

func (self Instance) SetInvertEnabled(value bool) {
	class(self).SetInvertEnabled(value)
}

func (self Instance) InvertBorder() float64 {
	return float64(float64(class(self).GetInvertBorder()))
}

func (self Instance) SetInvertBorder(value float64) {
	class(self).SetInvertBorder(gd.Float(value))
}

func (self Instance) Polygon() []gd.Vector2 {
	return []gd.Vector2(class(self).GetPolygon().AsSlice())
}

func (self Instance) SetPolygon(value []gd.Vector2) {
	class(self).SetPolygon(gd.NewPackedVector2Slice(value))
}

func (self Instance) Uv() []gd.Vector2 {
	return []gd.Vector2(class(self).GetUv().AsSlice())
}

func (self Instance) SetUv(value []gd.Vector2) {
	class(self).SetUv(gd.NewPackedVector2Slice(value))
}

func (self Instance) VertexColors() []gd.Color {
	return []gd.Color(class(self).GetVertexColors().AsSlice())
}

func (self Instance) SetVertexColors(value []gd.Color) {
	class(self).SetVertexColors(gd.NewPackedColorSlice(value))
}

func (self Instance) Polygons() gd.Array {
	return gd.Array(class(self).GetPolygons())
}

func (self Instance) SetPolygons(value gd.Array) {
	class(self).SetPolygons(value)
}

func (self Instance) InternalVertexCount() int {
	return int(int(class(self).GetInternalVertexCount()))
}

func (self Instance) SetInternalVertexCount(value int) {
	class(self).SetInternalVertexCount(gd.Int(value))
}

//go:nosplit
func (self class) SetPolygon(polygon gd.PackedVector2Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(polygon))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPolygon() gd.PackedVector2Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUv(uv gd.PackedVector2Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(uv))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_uv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetUv() gd.PackedVector2Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_uv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetColor(color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetColor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPolygons(polygons gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(polygons))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_polygons, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPolygons() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_polygons, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVertexColors(vertex_colors gd.PackedColorArray) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(vertex_colors))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_vertex_colors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVertexColors() gd.PackedColorArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_vertex_colors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedColorArray](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTexture(texture objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTexture() objects.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureOffset(texture_offset gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, texture_offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_texture_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureOffset() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_texture_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureRotation(texture_rotation gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, texture_rotation)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_texture_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureRotation() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_texture_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureScale(texture_scale gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, texture_scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_texture_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureScale() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_texture_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInvertEnabled(invert bool) {
	var frame = callframe.New()
	callframe.Arg(frame, invert)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_invert_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetInvertEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_invert_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAntialiased(antialiased bool) {
	var frame = callframe.New()
	callframe.Arg(frame, antialiased)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_antialiased, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAntialiased() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_antialiased, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInvertBorder(invert_border gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, invert_border)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_invert_border, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetInvertBorder() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_invert_border, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOffset(offset gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetOffset() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a bone with the specified [param path] and [param weights].
*/
//go:nosplit
func (self class) AddBone(path gd.NodePath, weights gd.PackedFloat32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	callframe.Arg(frame, pointers.Get(weights))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_add_bone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the number of bones in this [Polygon2D].
*/
//go:nosplit
func (self class) GetBoneCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_bone_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the path to the node associated with the specified bone.
*/
//go:nosplit
func (self class) GetBonePath(index gd.Int) gd.NodePath {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_bone_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the weight values of the specified bone.
*/
//go:nosplit
func (self class) GetBoneWeights(index gd.Int) gd.PackedFloat32Array {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_bone_weights, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedFloat32Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Removes the specified bone from this [Polygon2D].
*/
//go:nosplit
func (self class) EraseBone(index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_erase_bone, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes all bones from this [Polygon2D].
*/
//go:nosplit
func (self class) ClearBones() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_clear_bones, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the path to the node associated with the specified bone.
*/
//go:nosplit
func (self class) SetBonePath(index gd.Int, path gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, pointers.Get(path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_bone_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the weight values for the specified bone.
*/
//go:nosplit
func (self class) SetBoneWeights(index gd.Int, weights gd.PackedFloat32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, pointers.Get(weights))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_bone_weights, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetSkeleton(skeleton gd.NodePath) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(skeleton))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSkeleton() gd.NodePath {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.NodePath](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInternalVertexCount(internal_vertex_count gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, internal_vertex_count)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_set_internal_vertex_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetInternalVertexCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Polygon2D.Bind_get_internal_vertex_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPolygon2D() Advanced        { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPolygon2D() Instance     { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.Advanced    { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode2D() Node2D.Instance { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode2D(), name)
	}
}
func init() { classdb.Register("Polygon2D", func(ptr gd.Object) any { return classdb.Polygon2D(ptr) }) }
