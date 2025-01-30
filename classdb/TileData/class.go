// Package TileData provides methods for working with TileData object instances.
package TileData

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
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Vector2i"

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
[TileData] object represents a single tile in a [TileSet]. It is usually edited using the tileset editor, but it can be modified at runtime using [method TileMap._tile_data_runtime_update].
*/
type Instance [1]gdclass.TileData

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsTileData() Instance
}

/*
Sets the occluder for the TileSet occlusion layer with index [param layer_id].
*/
func (self Instance) SetOccluder(layer_id int, occluder_polygon [1]gdclass.OccluderPolygon2D) { //gd:TileData.set_occluder
	class(self).SetOccluder(int64(layer_id), occluder_polygon)
}

/*
Returns the occluder polygon of the tile for the TileSet occlusion layer with index [param layer_id].
[param flip_h], [param flip_v], and [param transpose] allow transforming the returned polygon.
*/
func (self Instance) GetOccluder(layer_id int) [1]gdclass.OccluderPolygon2D { //gd:TileData.get_occluder
	return [1]gdclass.OccluderPolygon2D(class(self).GetOccluder(int64(layer_id), false, false, false))
}

/*
Sets the constant linear velocity. This does not move the tile. This linear velocity is applied to objects colliding with this tile. This is useful to create conveyor belts.
*/
func (self Instance) SetConstantLinearVelocity(layer_id int, velocity Vector2.XY) { //gd:TileData.set_constant_linear_velocity
	class(self).SetConstantLinearVelocity(int64(layer_id), Vector2.XY(velocity))
}

/*
Returns the constant linear velocity applied to objects colliding with this tile.
*/
func (self Instance) GetConstantLinearVelocity(layer_id int) Vector2.XY { //gd:TileData.get_constant_linear_velocity
	return Vector2.XY(class(self).GetConstantLinearVelocity(int64(layer_id)))
}

/*
Sets the constant angular velocity. This does not rotate the tile. This angular velocity is applied to objects colliding with this tile.
*/
func (self Instance) SetConstantAngularVelocity(layer_id int, velocity Float.X) { //gd:TileData.set_constant_angular_velocity
	class(self).SetConstantAngularVelocity(int64(layer_id), float64(velocity))
}

/*
Returns the constant angular velocity applied to objects colliding with this tile.
*/
func (self Instance) GetConstantAngularVelocity(layer_id int) Float.X { //gd:TileData.get_constant_angular_velocity
	return Float.X(Float.X(class(self).GetConstantAngularVelocity(int64(layer_id))))
}

/*
Sets the polygons count for TileSet physics layer with index [param layer_id].
*/
func (self Instance) SetCollisionPolygonsCount(layer_id int, polygons_count int) { //gd:TileData.set_collision_polygons_count
	class(self).SetCollisionPolygonsCount(int64(layer_id), int64(polygons_count))
}

/*
Returns how many polygons the tile has for TileSet physics layer with index [param layer_id].
*/
func (self Instance) GetCollisionPolygonsCount(layer_id int) int { //gd:TileData.get_collision_polygons_count
	return int(int(class(self).GetCollisionPolygonsCount(int64(layer_id))))
}

/*
Adds a collision polygon to the tile on the given TileSet physics layer.
*/
func (self Instance) AddCollisionPolygon(layer_id int) { //gd:TileData.add_collision_polygon
	class(self).AddCollisionPolygon(int64(layer_id))
}

/*
Removes the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
func (self Instance) RemoveCollisionPolygon(layer_id int, polygon_index int) { //gd:TileData.remove_collision_polygon
	class(self).RemoveCollisionPolygon(int64(layer_id), int64(polygon_index))
}

/*
Sets the points of the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
func (self Instance) SetCollisionPolygonPoints(layer_id int, polygon_index int, polygon []Vector2.XY) { //gd:TileData.set_collision_polygon_points
	class(self).SetCollisionPolygonPoints(int64(layer_id), int64(polygon_index), Packed.New(polygon...))
}

/*
Returns the points of the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
func (self Instance) GetCollisionPolygonPoints(layer_id int, polygon_index int) []Vector2.XY { //gd:TileData.get_collision_polygon_points
	return []Vector2.XY(slices.Collect(class(self).GetCollisionPolygonPoints(int64(layer_id), int64(polygon_index)).Values()))
}

/*
Enables/disables one-way collisions on the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
func (self Instance) SetCollisionPolygonOneWay(layer_id int, polygon_index int, one_way bool) { //gd:TileData.set_collision_polygon_one_way
	class(self).SetCollisionPolygonOneWay(int64(layer_id), int64(polygon_index), one_way)
}

/*
Returns whether one-way collisions are enabled for the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
func (self Instance) IsCollisionPolygonOneWay(layer_id int, polygon_index int) bool { //gd:TileData.is_collision_polygon_one_way
	return bool(class(self).IsCollisionPolygonOneWay(int64(layer_id), int64(polygon_index)))
}

/*
Enables/disables one-way collisions on the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
func (self Instance) SetCollisionPolygonOneWayMargin(layer_id int, polygon_index int, one_way_margin Float.X) { //gd:TileData.set_collision_polygon_one_way_margin
	class(self).SetCollisionPolygonOneWayMargin(int64(layer_id), int64(polygon_index), float64(one_way_margin))
}

/*
Returns the one-way margin (for one-way platforms) of the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
func (self Instance) GetCollisionPolygonOneWayMargin(layer_id int, polygon_index int) Float.X { //gd:TileData.get_collision_polygon_one_way_margin
	return Float.X(Float.X(class(self).GetCollisionPolygonOneWayMargin(int64(layer_id), int64(polygon_index))))
}

/*
Sets the tile's terrain bit for the given [param peering_bit] direction. To check that a direction is valid, use [method is_valid_terrain_peering_bit].
*/
func (self Instance) SetTerrainPeeringBit(peering_bit gdclass.TileSetCellNeighbor, terrain int) { //gd:TileData.set_terrain_peering_bit
	class(self).SetTerrainPeeringBit(peering_bit, int64(terrain))
}

/*
Returns the tile's terrain bit for the given [param peering_bit] direction. To check that a direction is valid, use [method is_valid_terrain_peering_bit].
*/
func (self Instance) GetTerrainPeeringBit(peering_bit gdclass.TileSetCellNeighbor) int { //gd:TileData.get_terrain_peering_bit
	return int(int(class(self).GetTerrainPeeringBit(peering_bit)))
}

/*
Returns whether the given [param peering_bit] direction is valid for this tile.
*/
func (self Instance) IsValidTerrainPeeringBit(peering_bit gdclass.TileSetCellNeighbor) bool { //gd:TileData.is_valid_terrain_peering_bit
	return bool(class(self).IsValidTerrainPeeringBit(peering_bit))
}

/*
Sets the navigation polygon for the TileSet navigation layer with index [param layer_id].
*/
func (self Instance) SetNavigationPolygon(layer_id int, navigation_polygon [1]gdclass.NavigationPolygon) { //gd:TileData.set_navigation_polygon
	class(self).SetNavigationPolygon(int64(layer_id), navigation_polygon)
}

/*
Returns the navigation polygon of the tile for the TileSet navigation layer with index [param layer_id].
[param flip_h], [param flip_v], and [param transpose] allow transforming the returned polygon.
*/
func (self Instance) GetNavigationPolygon(layer_id int) [1]gdclass.NavigationPolygon { //gd:TileData.get_navigation_polygon
	return [1]gdclass.NavigationPolygon(class(self).GetNavigationPolygon(int64(layer_id), false, false, false))
}

/*
Sets the tile's custom data value for the TileSet custom data layer with name [param layer_name].
*/
func (self Instance) SetCustomData(layer_name string, value any) { //gd:TileData.set_custom_data
	class(self).SetCustomData(String.New(layer_name), variant.New(value))
}

/*
Returns the custom data value for custom data layer named [param layer_name].
*/
func (self Instance) GetCustomData(layer_name string) any { //gd:TileData.get_custom_data
	return any(class(self).GetCustomData(String.New(layer_name)).Interface())
}

/*
Sets the tile's custom data value for the TileSet custom data layer with index [param layer_id].
*/
func (self Instance) SetCustomDataByLayerId(layer_id int, value any) { //gd:TileData.set_custom_data_by_layer_id
	class(self).SetCustomDataByLayerId(int64(layer_id), variant.New(value))
}

/*
Returns the custom data value for custom data layer with index [param layer_id].
*/
func (self Instance) GetCustomDataByLayerId(layer_id int) any { //gd:TileData.get_custom_data_by_layer_id
	return any(class(self).GetCustomDataByLayerId(int64(layer_id)).Interface())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.TileData

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TileData"))
	casted := Instance{*(*gdclass.TileData)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) FlipH() bool {
	return bool(class(self).GetFlipH())
}

func (self Instance) SetFlipH(value bool) {
	class(self).SetFlipH(value)
}

func (self Instance) FlipV() bool {
	return bool(class(self).GetFlipV())
}

func (self Instance) SetFlipV(value bool) {
	class(self).SetFlipV(value)
}

func (self Instance) Transpose() bool {
	return bool(class(self).GetTranspose())
}

func (self Instance) SetTranspose(value bool) {
	class(self).SetTranspose(value)
}

func (self Instance) TextureOrigin() Vector2i.XY {
	return Vector2i.XY(class(self).GetTextureOrigin())
}

func (self Instance) SetTextureOrigin(value Vector2i.XY) {
	class(self).SetTextureOrigin(Vector2i.XY(value))
}

func (self Instance) Modulate() Color.RGBA {
	return Color.RGBA(class(self).GetModulate())
}

func (self Instance) SetModulate(value Color.RGBA) {
	class(self).SetModulate(Color.RGBA(value))
}

func (self Instance) Material() [1]gdclass.Material {
	return [1]gdclass.Material(class(self).GetMaterial())
}

func (self Instance) SetMaterial(value [1]gdclass.Material) {
	class(self).SetMaterial(value)
}

func (self Instance) ZIndex() int {
	return int(int(class(self).GetZIndex()))
}

func (self Instance) SetZIndex(value int) {
	class(self).SetZIndex(int64(value))
}

func (self Instance) YSortOrigin() int {
	return int(int(class(self).GetYSortOrigin()))
}

func (self Instance) SetYSortOrigin(value int) {
	class(self).SetYSortOrigin(int64(value))
}

func (self Instance) TerrainSet() int {
	return int(int(class(self).GetTerrainSet()))
}

func (self Instance) SetTerrainSet(value int) {
	class(self).SetTerrainSet(int64(value))
}

func (self Instance) Terrain() int {
	return int(int(class(self).GetTerrain()))
}

func (self Instance) SetTerrain(value int) {
	class(self).SetTerrain(int64(value))
}

func (self Instance) Probability() Float.X {
	return Float.X(Float.X(class(self).GetProbability()))
}

func (self Instance) SetProbability(value Float.X) {
	class(self).SetProbability(float64(value))
}

//go:nosplit
func (self class) SetFlipH(flip_h bool) { //gd:TileData.set_flip_h
	var frame = callframe.New()
	callframe.Arg(frame, flip_h)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_flip_h, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFlipH() bool { //gd:TileData.get_flip_h
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_flip_h, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFlipV(flip_v bool) { //gd:TileData.set_flip_v
	var frame = callframe.New()
	callframe.Arg(frame, flip_v)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_flip_v, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFlipV() bool { //gd:TileData.get_flip_v
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_flip_v, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTranspose(transpose bool) { //gd:TileData.set_transpose
	var frame = callframe.New()
	callframe.Arg(frame, transpose)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_transpose, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTranspose() bool { //gd:TileData.get_transpose
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_transpose, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaterial(material [1]gdclass.Material) { //gd:TileData.set_material
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(material[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_material, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaterial() [1]gdclass.Material { //gd:TileData.get_material
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_material, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Material{gd.PointerWithOwnershipTransferredToGo[gdclass.Material](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureOrigin(texture_origin Vector2i.XY) { //gd:TileData.set_texture_origin
	var frame = callframe.New()
	callframe.Arg(frame, texture_origin)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_texture_origin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureOrigin() Vector2i.XY { //gd:TileData.get_texture_origin
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector2i.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_texture_origin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetModulate(modulate Color.RGBA) { //gd:TileData.set_modulate
	var frame = callframe.New()
	callframe.Arg(frame, modulate)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_modulate, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetModulate() Color.RGBA { //gd:TileData.get_modulate
	var frame = callframe.New()
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_modulate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetZIndex(z_index int64) { //gd:TileData.set_z_index
	var frame = callframe.New()
	callframe.Arg(frame, z_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_z_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetZIndex() int64 { //gd:TileData.get_z_index
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_z_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetYSortOrigin(y_sort_origin int64) { //gd:TileData.set_y_sort_origin
	var frame = callframe.New()
	callframe.Arg(frame, y_sort_origin)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_y_sort_origin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetYSortOrigin() int64 { //gd:TileData.get_y_sort_origin
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_y_sort_origin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the occluder for the TileSet occlusion layer with index [param layer_id].
*/
//go:nosplit
func (self class) SetOccluder(layer_id int64, occluder_polygon [1]gdclass.OccluderPolygon2D) { //gd:TileData.set_occluder
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, pointers.Get(occluder_polygon[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_occluder, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the occluder polygon of the tile for the TileSet occlusion layer with index [param layer_id].
[param flip_h], [param flip_v], and [param transpose] allow transforming the returned polygon.
*/
//go:nosplit
func (self class) GetOccluder(layer_id int64, flip_h bool, flip_v bool, transpose bool) [1]gdclass.OccluderPolygon2D { //gd:TileData.get_occluder
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, flip_h)
	callframe.Arg(frame, flip_v)
	callframe.Arg(frame, transpose)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_occluder, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.OccluderPolygon2D{gd.PointerWithOwnershipTransferredToGo[gdclass.OccluderPolygon2D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sets the constant linear velocity. This does not move the tile. This linear velocity is applied to objects colliding with this tile. This is useful to create conveyor belts.
*/
//go:nosplit
func (self class) SetConstantLinearVelocity(layer_id int64, velocity Vector2.XY) { //gd:TileData.set_constant_linear_velocity
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, velocity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_constant_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the constant linear velocity applied to objects colliding with this tile.
*/
//go:nosplit
func (self class) GetConstantLinearVelocity(layer_id int64) Vector2.XY { //gd:TileData.get_constant_linear_velocity
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_constant_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the constant angular velocity. This does not rotate the tile. This angular velocity is applied to objects colliding with this tile.
*/
//go:nosplit
func (self class) SetConstantAngularVelocity(layer_id int64, velocity float64) { //gd:TileData.set_constant_angular_velocity
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, velocity)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_constant_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the constant angular velocity applied to objects colliding with this tile.
*/
//go:nosplit
func (self class) GetConstantAngularVelocity(layer_id int64) float64 { //gd:TileData.get_constant_angular_velocity
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_constant_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the polygons count for TileSet physics layer with index [param layer_id].
*/
//go:nosplit
func (self class) SetCollisionPolygonsCount(layer_id int64, polygons_count int64) { //gd:TileData.set_collision_polygons_count
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, polygons_count)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_collision_polygons_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns how many polygons the tile has for TileSet physics layer with index [param layer_id].
*/
//go:nosplit
func (self class) GetCollisionPolygonsCount(layer_id int64) int64 { //gd:TileData.get_collision_polygons_count
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_collision_polygons_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a collision polygon to the tile on the given TileSet physics layer.
*/
//go:nosplit
func (self class) AddCollisionPolygon(layer_id int64) { //gd:TileData.add_collision_polygon
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_add_collision_polygon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
//go:nosplit
func (self class) RemoveCollisionPolygon(layer_id int64, polygon_index int64) { //gd:TileData.remove_collision_polygon
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, polygon_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_remove_collision_polygon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the points of the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
//go:nosplit
func (self class) SetCollisionPolygonPoints(layer_id int64, polygon_index int64, polygon Packed.Array[Vector2.XY]) { //gd:TileData.set_collision_polygon_points
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, polygon_index)
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](polygon)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_collision_polygon_points, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the points of the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
//go:nosplit
func (self class) GetCollisionPolygonPoints(layer_id int64, polygon_index int64) Packed.Array[Vector2.XY] { //gd:TileData.get_collision_polygon_points
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, polygon_index)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_collision_polygon_points, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[Vector2.XY](Array.Through(gd.PackedProxy[gd.PackedVector2Array, Vector2.XY]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Enables/disables one-way collisions on the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
//go:nosplit
func (self class) SetCollisionPolygonOneWay(layer_id int64, polygon_index int64, one_way bool) { //gd:TileData.set_collision_polygon_one_way
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, polygon_index)
	callframe.Arg(frame, one_way)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_collision_polygon_one_way, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether one-way collisions are enabled for the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
//go:nosplit
func (self class) IsCollisionPolygonOneWay(layer_id int64, polygon_index int64) bool { //gd:TileData.is_collision_polygon_one_way
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, polygon_index)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_is_collision_polygon_one_way, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Enables/disables one-way collisions on the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
//go:nosplit
func (self class) SetCollisionPolygonOneWayMargin(layer_id int64, polygon_index int64, one_way_margin float64) { //gd:TileData.set_collision_polygon_one_way_margin
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, polygon_index)
	callframe.Arg(frame, one_way_margin)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_collision_polygon_one_way_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the one-way margin (for one-way platforms) of the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
//go:nosplit
func (self class) GetCollisionPolygonOneWayMargin(layer_id int64, polygon_index int64) float64 { //gd:TileData.get_collision_polygon_one_way_margin
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, polygon_index)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_collision_polygon_one_way_margin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTerrainSet(terrain_set int64) { //gd:TileData.set_terrain_set
	var frame = callframe.New()
	callframe.Arg(frame, terrain_set)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_terrain_set, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTerrainSet() int64 { //gd:TileData.get_terrain_set
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_terrain_set, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTerrain(terrain int64) { //gd:TileData.set_terrain
	var frame = callframe.New()
	callframe.Arg(frame, terrain)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_terrain, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTerrain() int64 { //gd:TileData.get_terrain
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_terrain, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the tile's terrain bit for the given [param peering_bit] direction. To check that a direction is valid, use [method is_valid_terrain_peering_bit].
*/
//go:nosplit
func (self class) SetTerrainPeeringBit(peering_bit gdclass.TileSetCellNeighbor, terrain int64) { //gd:TileData.set_terrain_peering_bit
	var frame = callframe.New()
	callframe.Arg(frame, peering_bit)
	callframe.Arg(frame, terrain)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_terrain_peering_bit, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the tile's terrain bit for the given [param peering_bit] direction. To check that a direction is valid, use [method is_valid_terrain_peering_bit].
*/
//go:nosplit
func (self class) GetTerrainPeeringBit(peering_bit gdclass.TileSetCellNeighbor) int64 { //gd:TileData.get_terrain_peering_bit
	var frame = callframe.New()
	callframe.Arg(frame, peering_bit)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_terrain_peering_bit, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns whether the given [param peering_bit] direction is valid for this tile.
*/
//go:nosplit
func (self class) IsValidTerrainPeeringBit(peering_bit gdclass.TileSetCellNeighbor) bool { //gd:TileData.is_valid_terrain_peering_bit
	var frame = callframe.New()
	callframe.Arg(frame, peering_bit)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_is_valid_terrain_peering_bit, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the navigation polygon for the TileSet navigation layer with index [param layer_id].
*/
//go:nosplit
func (self class) SetNavigationPolygon(layer_id int64, navigation_polygon [1]gdclass.NavigationPolygon) { //gd:TileData.set_navigation_polygon
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, pointers.Get(navigation_polygon[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_navigation_polygon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the navigation polygon of the tile for the TileSet navigation layer with index [param layer_id].
[param flip_h], [param flip_v], and [param transpose] allow transforming the returned polygon.
*/
//go:nosplit
func (self class) GetNavigationPolygon(layer_id int64, flip_h bool, flip_v bool, transpose bool) [1]gdclass.NavigationPolygon { //gd:TileData.get_navigation_polygon
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, flip_h)
	callframe.Arg(frame, flip_v)
	callframe.Arg(frame, transpose)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_navigation_polygon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.NavigationPolygon{gd.PointerWithOwnershipTransferredToGo[gdclass.NavigationPolygon](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetProbability(probability float64) { //gd:TileData.set_probability
	var frame = callframe.New()
	callframe.Arg(frame, probability)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_probability, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetProbability() float64 { //gd:TileData.get_probability
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_probability, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the tile's custom data value for the TileSet custom data layer with name [param layer_name].
*/
//go:nosplit
func (self class) SetCustomData(layer_name String.Readable, value variant.Any) { //gd:TileData.set_custom_data
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(layer_name)))
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(value)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_custom_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the custom data value for custom data layer named [param layer_name].
*/
//go:nosplit
func (self class) GetCustomData(layer_name String.Readable) variant.Any { //gd:TileData.get_custom_data
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(layer_name)))
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_custom_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Through(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets the tile's custom data value for the TileSet custom data layer with index [param layer_id].
*/
//go:nosplit
func (self class) SetCustomDataByLayerId(layer_id int64, value variant.Any) { //gd:TileData.set_custom_data_by_layer_id
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(value)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_custom_data_by_layer_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the custom data value for custom data layer with index [param layer_id].
*/
//go:nosplit
func (self class) GetCustomDataByLayerId(layer_id int64) variant.Any { //gd:TileData.get_custom_data_by_layer_id
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_custom_data_by_layer_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Through(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}
func (self Instance) OnChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("changed"), gd.NewCallable(cb), 0)
}

func (self class) AsTileData() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTileData() Instance { return *((*Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Instance(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("TileData", func(ptr gd.Object) any { return [1]gdclass.TileData{*(*gdclass.TileData)(unsafe.Pointer(&ptr))} })
}
