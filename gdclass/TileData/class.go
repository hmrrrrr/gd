package TileData

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
[TileData] object represents a single tile in a [TileSet]. It is usually edited using the tileset editor, but it can be modified at runtime using [method TileMap._tile_data_runtime_update].
*/
type Instance [1]classdb.TileData

/*
Sets the occluder for the TileSet occlusion layer with index [param layer_id].
*/
func (self Instance) SetOccluder(layer_id int, occluder_polygon gdclass.OccluderPolygon2D) {
	class(self).SetOccluder(gd.Int(layer_id), occluder_polygon)
}

/*
Returns the occluder polygon of the tile for the TileSet occlusion layer with index [param layer_id].
[param flip_h], [param flip_v], and [param transpose] allow transforming the returned polygon.
*/
func (self Instance) GetOccluder(layer_id int) gdclass.OccluderPolygon2D {
	return gdclass.OccluderPolygon2D(class(self).GetOccluder(gd.Int(layer_id), false, false, false))
}

/*
Sets the constant linear velocity. This does not move the tile. This linear velocity is applied to objects colliding with this tile. This is useful to create conveyor belts.
*/
func (self Instance) SetConstantLinearVelocity(layer_id int, velocity gd.Vector2) {
	class(self).SetConstantLinearVelocity(gd.Int(layer_id), velocity)
}

/*
Returns the constant linear velocity applied to objects colliding with this tile.
*/
func (self Instance) GetConstantLinearVelocity(layer_id int) gd.Vector2 {
	return gd.Vector2(class(self).GetConstantLinearVelocity(gd.Int(layer_id)))
}

/*
Sets the constant angular velocity. This does not rotate the tile. This angular velocity is applied to objects colliding with this tile.
*/
func (self Instance) SetConstantAngularVelocity(layer_id int, velocity float64) {
	class(self).SetConstantAngularVelocity(gd.Int(layer_id), gd.Float(velocity))
}

/*
Returns the constant angular velocity applied to objects colliding with this tile.
*/
func (self Instance) GetConstantAngularVelocity(layer_id int) float64 {
	return float64(float64(class(self).GetConstantAngularVelocity(gd.Int(layer_id))))
}

/*
Sets the polygons count for TileSet physics layer with index [param layer_id].
*/
func (self Instance) SetCollisionPolygonsCount(layer_id int, polygons_count int) {
	class(self).SetCollisionPolygonsCount(gd.Int(layer_id), gd.Int(polygons_count))
}

/*
Returns how many polygons the tile has for TileSet physics layer with index [param layer_id].
*/
func (self Instance) GetCollisionPolygonsCount(layer_id int) int {
	return int(int(class(self).GetCollisionPolygonsCount(gd.Int(layer_id))))
}

/*
Adds a collision polygon to the tile on the given TileSet physics layer.
*/
func (self Instance) AddCollisionPolygon(layer_id int) {
	class(self).AddCollisionPolygon(gd.Int(layer_id))
}

/*
Removes the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
func (self Instance) RemoveCollisionPolygon(layer_id int, polygon_index int) {
	class(self).RemoveCollisionPolygon(gd.Int(layer_id), gd.Int(polygon_index))
}

/*
Sets the points of the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
func (self Instance) SetCollisionPolygonPoints(layer_id int, polygon_index int, polygon []gd.Vector2) {
	class(self).SetCollisionPolygonPoints(gd.Int(layer_id), gd.Int(polygon_index), gd.NewPackedVector2Slice(polygon))
}

/*
Returns the points of the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
func (self Instance) GetCollisionPolygonPoints(layer_id int, polygon_index int) []gd.Vector2 {
	return []gd.Vector2(class(self).GetCollisionPolygonPoints(gd.Int(layer_id), gd.Int(polygon_index)).AsSlice())
}

/*
Enables/disables one-way collisions on the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
func (self Instance) SetCollisionPolygonOneWay(layer_id int, polygon_index int, one_way bool) {
	class(self).SetCollisionPolygonOneWay(gd.Int(layer_id), gd.Int(polygon_index), one_way)
}

/*
Returns whether one-way collisions are enabled for the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
func (self Instance) IsCollisionPolygonOneWay(layer_id int, polygon_index int) bool {
	return bool(class(self).IsCollisionPolygonOneWay(gd.Int(layer_id), gd.Int(polygon_index)))
}

/*
Enables/disables one-way collisions on the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
func (self Instance) SetCollisionPolygonOneWayMargin(layer_id int, polygon_index int, one_way_margin float64) {
	class(self).SetCollisionPolygonOneWayMargin(gd.Int(layer_id), gd.Int(polygon_index), gd.Float(one_way_margin))
}

/*
Returns the one-way margin (for one-way platforms) of the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
func (self Instance) GetCollisionPolygonOneWayMargin(layer_id int, polygon_index int) float64 {
	return float64(float64(class(self).GetCollisionPolygonOneWayMargin(gd.Int(layer_id), gd.Int(polygon_index))))
}

/*
Sets the tile's terrain bit for the given [param peering_bit] direction. To check that a direction is valid, use [method is_valid_terrain_peering_bit].
*/
func (self Instance) SetTerrainPeeringBit(peering_bit classdb.TileSetCellNeighbor, terrain int) {
	class(self).SetTerrainPeeringBit(peering_bit, gd.Int(terrain))
}

/*
Returns the tile's terrain bit for the given [param peering_bit] direction. To check that a direction is valid, use [method is_valid_terrain_peering_bit].
*/
func (self Instance) GetTerrainPeeringBit(peering_bit classdb.TileSetCellNeighbor) int {
	return int(int(class(self).GetTerrainPeeringBit(peering_bit)))
}

/*
Returns whether the given [param peering_bit] direction is valid for this tile.
*/
func (self Instance) IsValidTerrainPeeringBit(peering_bit classdb.TileSetCellNeighbor) bool {
	return bool(class(self).IsValidTerrainPeeringBit(peering_bit))
}

/*
Sets the navigation polygon for the TileSet navigation layer with index [param layer_id].
*/
func (self Instance) SetNavigationPolygon(layer_id int, navigation_polygon gdclass.NavigationPolygon) {
	class(self).SetNavigationPolygon(gd.Int(layer_id), navigation_polygon)
}

/*
Returns the navigation polygon of the tile for the TileSet navigation layer with index [param layer_id].
[param flip_h], [param flip_v], and [param transpose] allow transforming the returned polygon.
*/
func (self Instance) GetNavigationPolygon(layer_id int) gdclass.NavigationPolygon {
	return gdclass.NavigationPolygon(class(self).GetNavigationPolygon(gd.Int(layer_id), false, false, false))
}

/*
Sets the tile's custom data value for the TileSet custom data layer with name [param layer_name].
*/
func (self Instance) SetCustomData(layer_name string, value gd.Variant) {
	class(self).SetCustomData(gd.NewString(layer_name), value)
}

/*
Returns the custom data value for custom data layer named [param layer_name].
*/
func (self Instance) GetCustomData(layer_name string) gd.Variant {
	return gd.Variant(class(self).GetCustomData(gd.NewString(layer_name)))
}

/*
Sets the tile's custom data value for the TileSet custom data layer with index [param layer_id].
*/
func (self Instance) SetCustomDataByLayerId(layer_id int, value gd.Variant) {
	class(self).SetCustomDataByLayerId(gd.Int(layer_id), value)
}

/*
Returns the custom data value for custom data layer with index [param layer_id].
*/
func (self Instance) GetCustomDataByLayerId(layer_id int) gd.Variant {
	return gd.Variant(class(self).GetCustomDataByLayerId(gd.Int(layer_id)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.TileData

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TileData"))
	return Instance{classdb.TileData(object)}
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

func (self Instance) TextureOrigin() gd.Vector2i {
	return gd.Vector2i(class(self).GetTextureOrigin())
}

func (self Instance) SetTextureOrigin(value gd.Vector2i) {
	class(self).SetTextureOrigin(value)
}

func (self Instance) Modulate() gd.Color {
	return gd.Color(class(self).GetModulate())
}

func (self Instance) SetModulate(value gd.Color) {
	class(self).SetModulate(value)
}

func (self Instance) Material() gdclass.Material {
	return gdclass.Material(class(self).GetMaterial())
}

func (self Instance) SetMaterial(value gdclass.Material) {
	class(self).SetMaterial(value)
}

func (self Instance) ZIndex() int {
	return int(int(class(self).GetZIndex()))
}

func (self Instance) SetZIndex(value int) {
	class(self).SetZIndex(gd.Int(value))
}

func (self Instance) YSortOrigin() int {
	return int(int(class(self).GetYSortOrigin()))
}

func (self Instance) SetYSortOrigin(value int) {
	class(self).SetYSortOrigin(gd.Int(value))
}

func (self Instance) TerrainSet() int {
	return int(int(class(self).GetTerrainSet()))
}

func (self Instance) SetTerrainSet(value int) {
	class(self).SetTerrainSet(gd.Int(value))
}

func (self Instance) Terrain() int {
	return int(int(class(self).GetTerrain()))
}

func (self Instance) SetTerrain(value int) {
	class(self).SetTerrain(gd.Int(value))
}

func (self Instance) Probability() float64 {
	return float64(float64(class(self).GetProbability()))
}

func (self Instance) SetProbability(value float64) {
	class(self).SetProbability(gd.Float(value))
}

//go:nosplit
func (self class) SetFlipH(flip_h bool) {
	var frame = callframe.New()
	callframe.Arg(frame, flip_h)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_flip_h, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFlipH() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_flip_h, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFlipV(flip_v bool) {
	var frame = callframe.New()
	callframe.Arg(frame, flip_v)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_flip_v, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFlipV() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_flip_v, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTranspose(transpose bool) {
	var frame = callframe.New()
	callframe.Arg(frame, transpose)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_transpose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTranspose() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_transpose, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaterial(material gdclass.Material) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(material[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaterial() gdclass.Material {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Material{classdb.Material(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureOrigin(texture_origin gd.Vector2i) {
	var frame = callframe.New()
	callframe.Arg(frame, texture_origin)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_texture_origin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureOrigin() gd.Vector2i {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_texture_origin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetModulate(modulate gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, modulate)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetModulate() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetZIndex(z_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, z_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_z_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetZIndex() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_z_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetYSortOrigin(y_sort_origin gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, y_sort_origin)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_y_sort_origin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetYSortOrigin() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_y_sort_origin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the occluder for the TileSet occlusion layer with index [param layer_id].
*/
//go:nosplit
func (self class) SetOccluder(layer_id gd.Int, occluder_polygon gdclass.OccluderPolygon2D) {
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, pointers.Get(occluder_polygon[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_occluder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the occluder polygon of the tile for the TileSet occlusion layer with index [param layer_id].
[param flip_h], [param flip_v], and [param transpose] allow transforming the returned polygon.
*/
//go:nosplit
func (self class) GetOccluder(layer_id gd.Int, flip_h bool, flip_v bool, transpose bool) gdclass.OccluderPolygon2D {
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, flip_h)
	callframe.Arg(frame, flip_v)
	callframe.Arg(frame, transpose)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_occluder, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.OccluderPolygon2D{classdb.OccluderPolygon2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Sets the constant linear velocity. This does not move the tile. This linear velocity is applied to objects colliding with this tile. This is useful to create conveyor belts.
*/
//go:nosplit
func (self class) SetConstantLinearVelocity(layer_id gd.Int, velocity gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, velocity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_constant_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the constant linear velocity applied to objects colliding with this tile.
*/
//go:nosplit
func (self class) GetConstantLinearVelocity(layer_id gd.Int) gd.Vector2 {
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_constant_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the constant angular velocity. This does not rotate the tile. This angular velocity is applied to objects colliding with this tile.
*/
//go:nosplit
func (self class) SetConstantAngularVelocity(layer_id gd.Int, velocity gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, velocity)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_constant_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the constant angular velocity applied to objects colliding with this tile.
*/
//go:nosplit
func (self class) GetConstantAngularVelocity(layer_id gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_constant_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the polygons count for TileSet physics layer with index [param layer_id].
*/
//go:nosplit
func (self class) SetCollisionPolygonsCount(layer_id gd.Int, polygons_count gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, polygons_count)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_collision_polygons_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns how many polygons the tile has for TileSet physics layer with index [param layer_id].
*/
//go:nosplit
func (self class) GetCollisionPolygonsCount(layer_id gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_collision_polygons_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a collision polygon to the tile on the given TileSet physics layer.
*/
//go:nosplit
func (self class) AddCollisionPolygon(layer_id gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_add_collision_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
//go:nosplit
func (self class) RemoveCollisionPolygon(layer_id gd.Int, polygon_index gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, polygon_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_remove_collision_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the points of the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
//go:nosplit
func (self class) SetCollisionPolygonPoints(layer_id gd.Int, polygon_index gd.Int, polygon gd.PackedVector2Array) {
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, polygon_index)
	callframe.Arg(frame, pointers.Get(polygon))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_collision_polygon_points, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the points of the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
//go:nosplit
func (self class) GetCollisionPolygonPoints(layer_id gd.Int, polygon_index gd.Int) gd.PackedVector2Array {
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, polygon_index)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_collision_polygon_points, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedVector2Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Enables/disables one-way collisions on the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
//go:nosplit
func (self class) SetCollisionPolygonOneWay(layer_id gd.Int, polygon_index gd.Int, one_way bool) {
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, polygon_index)
	callframe.Arg(frame, one_way)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_collision_polygon_one_way, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether one-way collisions are enabled for the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
//go:nosplit
func (self class) IsCollisionPolygonOneWay(layer_id gd.Int, polygon_index gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, polygon_index)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_is_collision_polygon_one_way, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Enables/disables one-way collisions on the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
//go:nosplit
func (self class) SetCollisionPolygonOneWayMargin(layer_id gd.Int, polygon_index gd.Int, one_way_margin gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, polygon_index)
	callframe.Arg(frame, one_way_margin)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_collision_polygon_one_way_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the one-way margin (for one-way platforms) of the polygon at index [param polygon_index] for TileSet physics layer with index [param layer_id].
*/
//go:nosplit
func (self class) GetCollisionPolygonOneWayMargin(layer_id gd.Int, polygon_index gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, polygon_index)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_collision_polygon_one_way_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTerrainSet(terrain_set gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, terrain_set)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_terrain_set, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTerrainSet() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_terrain_set, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTerrain(terrain gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, terrain)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_terrain, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTerrain() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_terrain, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the tile's terrain bit for the given [param peering_bit] direction. To check that a direction is valid, use [method is_valid_terrain_peering_bit].
*/
//go:nosplit
func (self class) SetTerrainPeeringBit(peering_bit classdb.TileSetCellNeighbor, terrain gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, peering_bit)
	callframe.Arg(frame, terrain)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_terrain_peering_bit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the tile's terrain bit for the given [param peering_bit] direction. To check that a direction is valid, use [method is_valid_terrain_peering_bit].
*/
//go:nosplit
func (self class) GetTerrainPeeringBit(peering_bit classdb.TileSetCellNeighbor) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, peering_bit)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_terrain_peering_bit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns whether the given [param peering_bit] direction is valid for this tile.
*/
//go:nosplit
func (self class) IsValidTerrainPeeringBit(peering_bit classdb.TileSetCellNeighbor) bool {
	var frame = callframe.New()
	callframe.Arg(frame, peering_bit)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_is_valid_terrain_peering_bit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the navigation polygon for the TileSet navigation layer with index [param layer_id].
*/
//go:nosplit
func (self class) SetNavigationPolygon(layer_id gd.Int, navigation_polygon gdclass.NavigationPolygon) {
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, pointers.Get(navigation_polygon[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_navigation_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the navigation polygon of the tile for the TileSet navigation layer with index [param layer_id].
[param flip_h], [param flip_v], and [param transpose] allow transforming the returned polygon.
*/
//go:nosplit
func (self class) GetNavigationPolygon(layer_id gd.Int, flip_h bool, flip_v bool, transpose bool) gdclass.NavigationPolygon {
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, flip_h)
	callframe.Arg(frame, flip_v)
	callframe.Arg(frame, transpose)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_navigation_polygon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.NavigationPolygon{classdb.NavigationPolygon(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetProbability(probability gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, probability)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_probability, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetProbability() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_probability, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the tile's custom data value for the TileSet custom data layer with name [param layer_name].
*/
//go:nosplit
func (self class) SetCustomData(layer_name gd.String, value gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(layer_name))
	callframe.Arg(frame, pointers.Get(value))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_custom_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the custom data value for custom data layer named [param layer_name].
*/
//go:nosplit
func (self class) GetCustomData(layer_name gd.String) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(layer_name))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_custom_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets the tile's custom data value for the TileSet custom data layer with index [param layer_id].
*/
//go:nosplit
func (self class) SetCustomDataByLayerId(layer_id gd.Int, value gd.Variant) {
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	callframe.Arg(frame, pointers.Get(value))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_set_custom_data_by_layer_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the custom data value for custom data layer with index [param layer_id].
*/
//go:nosplit
func (self class) GetCustomDataByLayerId(layer_id gd.Int) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, layer_id)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileData.Bind_get_custom_data_by_layer_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}
func (self Instance) OnChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("changed"), gd.NewCallable(cb), 0)
}

func (self class) AsTileData() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTileData() Instance { return *((*Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() { classdb.Register("TileData", func(ptr gd.Object) any { return classdb.TileData(ptr) }) }
