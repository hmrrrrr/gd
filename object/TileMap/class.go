package TileMap

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Node2D"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Node for 2D tile-based maps. Tilemaps use a [TileSet] which contain a list of tiles which are used to create grid-based maps. A TileMap may have several layers, layouting tiles on top of each other.
For performance reasons, all TileMap updates are batched at the end of a frame. Notably, this means that scene tiles from a [TileSetScenesCollectionSource] may be initialized after their parent. This is only queued when inside the scene tree.
To force an update earlier on, call [method update_internals].
	// TileMap methods that can be overridden by a [Class] that extends it.
	type TileMap interface {
		//Should return [code]true[/code] if the tile at coordinates [param coords] on layer [param layer] requires a runtime update.
		//[b]Warning:[/b] Make sure this function only return [code]true[/code] when needed. Any tile processed at runtime without a need for it will imply a significant performance penalty.
		//[b]Note:[/b] If the result of this function should changed, use [method notify_runtime_tile_data_update] to notify the TileMap it needs an update.
		UseTileDataRuntimeUpdate(layer gd.Int, coords gd.Vector2i) bool
		//Called with a TileData object about to be used internally by the TileMap, allowing its modification at runtime.
		//This method is only called if [method _use_tile_data_runtime_update] is implemented and returns [code]true[/code] for the given tile [param coords] and [param layer].
		//[b]Warning:[/b] The [param tile_data] object's sub-resources are the same as the one in the TileSet. Modifying them might impact the whole TileSet. Instead, make sure to duplicate those resources.
		//[b]Note:[/b] If the properties of [param tile_data] object should change over time, use [method notify_runtime_tile_data_update] to notify the TileMap it needs an update.
		TileDataRuntimeUpdate(layer gd.Int, coords gd.Vector2i, tile_data object.TileData) 
	}

*/
type Simple [1]classdb.TileMap
func (Simple) _use_tile_data_runtime_update(impl func(ptr unsafe.Pointer, layer int, coords gd.Vector2i) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var layer = gd.UnsafeGet[gd.Int](p_args,0)
		var coords = gd.UnsafeGet[gd.Vector2i](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(layer), coords)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _tile_data_runtime_update(impl func(ptr unsafe.Pointer, layer int, coords gd.Vector2i, tile_data [1]classdb.TileData) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var layer = gd.UnsafeGet[gd.Int](p_args,0)
		var coords = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var tile_data [1]classdb.TileData
		tile_data[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,2)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, int(layer), coords, tile_data)
		gc.End()
	}
}
func (self Simple) SetNavigationMap(layer int, mapping gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNavigationMap(gd.Int(layer), mapping)
}
func (self Simple) GetNavigationMap(layer int) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).GetNavigationMap(gd.Int(layer)))
}
func (self Simple) ForceUpdate(layer int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ForceUpdate(gd.Int(layer))
}
func (self Simple) SetTileset(tileset [1]classdb.TileSet) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTileset(tileset)
}
func (self Simple) GetTileset() [1]classdb.TileSet {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TileSet(Expert(self).GetTileset(gc))
}
func (self Simple) SetRenderingQuadrantSize(size int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRenderingQuadrantSize(gd.Int(size))
}
func (self Simple) GetRenderingQuadrantSize() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetRenderingQuadrantSize()))
}
func (self Simple) GetLayersCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetLayersCount()))
}
func (self Simple) AddLayer(to_position int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddLayer(gd.Int(to_position))
}
func (self Simple) MoveLayer(layer int, to_position int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).MoveLayer(gd.Int(layer), gd.Int(to_position))
}
func (self Simple) RemoveLayer(layer int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveLayer(gd.Int(layer))
}
func (self Simple) SetLayerName(layer int, name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLayerName(gd.Int(layer), gc.String(name))
}
func (self Simple) GetLayerName(layer int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetLayerName(gc, gd.Int(layer)).String())
}
func (self Simple) SetLayerEnabled(layer int, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLayerEnabled(gd.Int(layer), enabled)
}
func (self Simple) IsLayerEnabled(layer int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsLayerEnabled(gd.Int(layer)))
}
func (self Simple) SetLayerModulate(layer int, modulate gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLayerModulate(gd.Int(layer), modulate)
}
func (self Simple) GetLayerModulate(layer int) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetLayerModulate(gd.Int(layer)))
}
func (self Simple) SetLayerYSortEnabled(layer int, y_sort_enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLayerYSortEnabled(gd.Int(layer), y_sort_enabled)
}
func (self Simple) IsLayerYSortEnabled(layer int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsLayerYSortEnabled(gd.Int(layer)))
}
func (self Simple) SetLayerYSortOrigin(layer int, y_sort_origin int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLayerYSortOrigin(gd.Int(layer), gd.Int(y_sort_origin))
}
func (self Simple) GetLayerYSortOrigin(layer int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetLayerYSortOrigin(gd.Int(layer))))
}
func (self Simple) SetLayerZIndex(layer int, z_index int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLayerZIndex(gd.Int(layer), gd.Int(z_index))
}
func (self Simple) GetLayerZIndex(layer int) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetLayerZIndex(gd.Int(layer))))
}
func (self Simple) SetLayerNavigationEnabled(layer int, enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLayerNavigationEnabled(gd.Int(layer), enabled)
}
func (self Simple) IsLayerNavigationEnabled(layer int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsLayerNavigationEnabled(gd.Int(layer)))
}
func (self Simple) SetLayerNavigationMap(layer int, mapping gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLayerNavigationMap(gd.Int(layer), mapping)
}
func (self Simple) GetLayerNavigationMap(layer int) gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).GetLayerNavigationMap(gd.Int(layer)))
}
func (self Simple) SetCollisionAnimatable(enabled bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCollisionAnimatable(enabled)
}
func (self Simple) IsCollisionAnimatable() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsCollisionAnimatable())
}
func (self Simple) SetCollisionVisibilityMode(collision_visibility_mode classdb.TileMapVisibilityMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCollisionVisibilityMode(collision_visibility_mode)
}
func (self Simple) GetCollisionVisibilityMode() classdb.TileMapVisibilityMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TileMapVisibilityMode(Expert(self).GetCollisionVisibilityMode())
}
func (self Simple) SetNavigationVisibilityMode(navigation_visibility_mode classdb.TileMapVisibilityMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNavigationVisibilityMode(navigation_visibility_mode)
}
func (self Simple) GetNavigationVisibilityMode() classdb.TileMapVisibilityMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.TileMapVisibilityMode(Expert(self).GetNavigationVisibilityMode())
}
func (self Simple) SetCell(layer int, coords gd.Vector2i, source_id int, atlas_coords gd.Vector2i, alternative_tile int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCell(gd.Int(layer), coords, gd.Int(source_id), atlas_coords, gd.Int(alternative_tile))
}
func (self Simple) EraseCell(layer int, coords gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EraseCell(gd.Int(layer), coords)
}
func (self Simple) GetCellSourceId(layer int, coords gd.Vector2i, use_proxies bool) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetCellSourceId(gd.Int(layer), coords, use_proxies)))
}
func (self Simple) GetCellAtlasCoords(layer int, coords gd.Vector2i, use_proxies bool) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).GetCellAtlasCoords(gd.Int(layer), coords, use_proxies))
}
func (self Simple) GetCellAlternativeTile(layer int, coords gd.Vector2i, use_proxies bool) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetCellAlternativeTile(gd.Int(layer), coords, use_proxies)))
}
func (self Simple) GetCellTileData(layer int, coords gd.Vector2i, use_proxies bool) [1]classdb.TileData {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TileData(Expert(self).GetCellTileData(gc, gd.Int(layer), coords, use_proxies))
}
func (self Simple) GetCoordsForBodyRid(body gd.RID) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).GetCoordsForBodyRid(body))
}
func (self Simple) GetLayerForBodyRid(body gd.RID) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetLayerForBodyRid(body)))
}
func (self Simple) GetPattern(layer int, coords_array gd.ArrayOf[gd.Vector2i]) [1]classdb.TileMapPattern {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.TileMapPattern(Expert(self).GetPattern(gc, gd.Int(layer), coords_array))
}
func (self Simple) MapPattern(position_in_tilemap gd.Vector2i, coords_in_pattern gd.Vector2i, pattern [1]classdb.TileMapPattern) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).MapPattern(position_in_tilemap, coords_in_pattern, pattern))
}
func (self Simple) SetPattern(layer int, position gd.Vector2i, pattern [1]classdb.TileMapPattern) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPattern(gd.Int(layer), position, pattern)
}
func (self Simple) SetCellsTerrainConnect(layer int, cells gd.ArrayOf[gd.Vector2i], terrain_set int, terrain int, ignore_empty_terrains bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCellsTerrainConnect(gd.Int(layer), cells, gd.Int(terrain_set), gd.Int(terrain), ignore_empty_terrains)
}
func (self Simple) SetCellsTerrainPath(layer int, path gd.ArrayOf[gd.Vector2i], terrain_set int, terrain int, ignore_empty_terrains bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCellsTerrainPath(gd.Int(layer), path, gd.Int(terrain_set), gd.Int(terrain), ignore_empty_terrains)
}
func (self Simple) FixInvalidTiles() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).FixInvalidTiles()
}
func (self Simple) ClearLayer(layer int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ClearLayer(gd.Int(layer))
}
func (self Simple) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Clear()
}
func (self Simple) UpdateInternals() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).UpdateInternals()
}
func (self Simple) NotifyRuntimeTileDataUpdate(layer int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).NotifyRuntimeTileDataUpdate(gd.Int(layer))
}
func (self Simple) GetSurroundingCells(coords gd.Vector2i) gd.ArrayOf[gd.Vector2i] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Vector2i](Expert(self).GetSurroundingCells(gc, coords))
}
func (self Simple) GetUsedCells(layer int) gd.ArrayOf[gd.Vector2i] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Vector2i](Expert(self).GetUsedCells(gc, gd.Int(layer)))
}
func (self Simple) GetUsedCellsById(layer int, source_id int, atlas_coords gd.Vector2i, alternative_tile int) gd.ArrayOf[gd.Vector2i] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gd.Vector2i](Expert(self).GetUsedCellsById(gc, gd.Int(layer), gd.Int(source_id), atlas_coords, gd.Int(alternative_tile)))
}
func (self Simple) GetUsedRect() gd.Rect2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2i(Expert(self).GetUsedRect())
}
func (self Simple) MapToLocal(map_position gd.Vector2i) gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).MapToLocal(map_position))
}
func (self Simple) LocalToMap(local_position gd.Vector2) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).LocalToMap(local_position))
}
func (self Simple) GetNeighborCell(coords gd.Vector2i, neighbor classdb.TileSetCellNeighbor) gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).GetNeighborCell(coords, neighbor))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.TileMap
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Should return [code]true[/code] if the tile at coordinates [param coords] on layer [param layer] requires a runtime update.
[b]Warning:[/b] Make sure this function only return [code]true[/code] when needed. Any tile processed at runtime without a need for it will imply a significant performance penalty.
[b]Note:[/b] If the result of this function should changed, use [method notify_runtime_tile_data_update] to notify the TileMap it needs an update.
*/
func (class) _use_tile_data_runtime_update(impl func(ptr unsafe.Pointer, layer gd.Int, coords gd.Vector2i) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var layer = gd.UnsafeGet[gd.Int](p_args,0)
		var coords = gd.UnsafeGet[gd.Vector2i](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, layer, coords)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Called with a TileData object about to be used internally by the TileMap, allowing its modification at runtime.
This method is only called if [method _use_tile_data_runtime_update] is implemented and returns [code]true[/code] for the given tile [param coords] and [param layer].
[b]Warning:[/b] The [param tile_data] object's sub-resources are the same as the one in the TileSet. Modifying them might impact the whole TileSet. Instead, make sure to duplicate those resources.
[b]Note:[/b] If the properties of [param tile_data] object should change over time, use [method notify_runtime_tile_data_update] to notify the TileMap it needs an update.
*/
func (class) _tile_data_runtime_update(impl func(ptr unsafe.Pointer, layer gd.Int, coords gd.Vector2i, tile_data object.TileData) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var layer = gd.UnsafeGet[gd.Int](p_args,0)
		var coords = gd.UnsafeGet[gd.Vector2i](p_args,1)
		var tile_data object.TileData
		tile_data[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,2)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, layer, coords, tile_data)
		ctx.End()
	}
}

/*
Assigns [param map] as a [NavigationServer2D] navigation map for the specified TileMap layer [param layer].
*/
//go:nosplit
func (self class) SetNavigationMap(layer gd.Int, mapping gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, mapping)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_set_navigation_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [RID] of the [NavigationServer2D] navigation map assigned to the specified TileMap layer [param layer].
*/
//go:nosplit
func (self class) GetNavigationMap(layer gd.Int) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_get_navigation_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Forces the TileMap and the layer [param layer] to update.
*/
//go:nosplit
func (self class) ForceUpdate(layer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_force_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetTileset(tileset object.TileSet)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(tileset[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_set_tileset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTileset(ctx gd.Lifetime) object.TileSet {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_get_tileset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TileSet
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRenderingQuadrantSize(size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_set_rendering_quadrant_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRenderingQuadrantSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_get_rendering_quadrant_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of layers in the TileMap.
*/
//go:nosplit
func (self class) GetLayersCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_get_layers_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a layer at the given position [param to_position] in the array. If [param to_position] is negative, the position is counted from the end, with [code]-1[/code] adding the layer at the end of the array.
*/
//go:nosplit
func (self class) AddLayer(to_position gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, to_position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_add_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Moves the layer at index [param layer] to the given position [param to_position] in the array.
*/
//go:nosplit
func (self class) MoveLayer(layer gd.Int, to_position gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, to_position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_move_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the layer at index [param layer].
*/
//go:nosplit
func (self class) RemoveLayer(layer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_remove_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a layer's name. This is mostly useful in the editor.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) SetLayerName(layer gd.Int, name gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_set_layer_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a TileMap layer's name.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) GetLayerName(ctx gd.Lifetime, layer gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_get_layer_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Enables or disables the layer [param layer]. A disabled layer is not processed at all (no rendering, no physics, etc.).
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) SetLayerEnabled(layer gd.Int, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_set_layer_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns if a layer is enabled.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) IsLayerEnabled(layer gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_is_layer_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets a layer's color. It will be multiplied by tile's color and TileMap's modulate.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) SetLayerModulate(layer gd.Int, modulate gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, modulate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_set_layer_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a TileMap layer's modulate.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) GetLayerModulate(layer gd.Int) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_get_layer_modulate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Enables or disables a layer's Y-sorting. If a layer is Y-sorted, the layer will behave as a CanvasItem node where each of its tile gets Y-sorted.
Y-sorted layers should usually be on different Z-index values than not Y-sorted layers, otherwise, each of those layer will be Y-sorted as whole with the Y-sorted one. This is usually an undesired behavior.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) SetLayerYSortEnabled(layer gd.Int, y_sort_enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, y_sort_enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_set_layer_y_sort_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns if a layer Y-sorts its tiles.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) IsLayerYSortEnabled(layer gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_is_layer_y_sort_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets a layer's Y-sort origin value. This Y-sort origin value is added to each tile's Y-sort origin value.
This allows, for example, to fake a different height level on each layer. This can be useful for top-down view games.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) SetLayerYSortOrigin(layer gd.Int, y_sort_origin gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, y_sort_origin)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_set_layer_y_sort_origin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a TileMap layer's Y sort origin.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) GetLayerYSortOrigin(layer gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_get_layer_y_sort_origin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets a layers Z-index value. This Z-index is added to each tile's Z-index value.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) SetLayerZIndex(layer gd.Int, z_index gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, z_index)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_set_layer_z_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a TileMap layer's Z-index value.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) GetLayerZIndex(layer gd.Int) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_get_layer_z_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Enables or disables a layer's built-in navigation regions generation. Disable this if you need to bake navigation regions from a TileMap using a [NavigationRegion2D] node.
*/
//go:nosplit
func (self class) SetLayerNavigationEnabled(layer gd.Int, enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_set_layer_navigation_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns if a layer's built-in navigation regions generation is enabled.
*/
//go:nosplit
func (self class) IsLayerNavigationEnabled(layer gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_is_layer_navigation_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Assigns [param map] as a [NavigationServer2D] navigation map for the specified TileMap layer [param layer].
By default the TileMap uses the default [World2D] navigation map for the first TileMap layer. For each additional TileMap layer a new navigation map is created for the additional layer.
In order to make [NavigationAgent2D] switch between TileMap layer navigation maps use [method NavigationAgent2D.set_navigation_map] with the navigation map received from [method get_layer_navigation_map].
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) SetLayerNavigationMap(layer gd.Int, mapping gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, mapping)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_set_layer_navigation_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [RID] of the [NavigationServer2D] navigation map assigned to the specified TileMap layer [param layer].
By default the TileMap uses the default [World2D] navigation map for the first TileMap layer. For each additional TileMap layer a new navigation map is created for the additional layer.
In order to make [NavigationAgent2D] switch between TileMap layer navigation maps use [method NavigationAgent2D.set_navigation_map] with the navigation map received from [method get_layer_navigation_map].
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) GetLayerNavigationMap(layer gd.Int) gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_get_layer_navigation_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollisionAnimatable(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_set_collision_animatable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsCollisionAnimatable() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_is_collision_animatable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetCollisionVisibilityMode(collision_visibility_mode classdb.TileMapVisibilityMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, collision_visibility_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_set_collision_visibility_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCollisionVisibilityMode() classdb.TileMapVisibilityMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TileMapVisibilityMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_get_collision_visibility_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetNavigationVisibilityMode(navigation_visibility_mode classdb.TileMapVisibilityMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, navigation_visibility_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_set_navigation_visibility_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetNavigationVisibilityMode() classdb.TileMapVisibilityMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TileMapVisibilityMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_get_navigation_visibility_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the tile identifiers for the cell on layer [param layer] at coordinates [param coords]. Each tile of the [TileSet] is identified using three parts:
- The source identifier [param source_id] identifies a [TileSetSource] identifier. See [method TileSet.set_source_id],
- The atlas coordinates identifier [param atlas_coords] identifies a tile coordinates in the atlas (if the source is a [TileSetAtlasSource]). For [TileSetScenesCollectionSource] it should always be [code]Vector2i(0, 0)[/code]),
- The alternative tile identifier [param alternative_tile] identifies a tile alternative in the atlas (if the source is a [TileSetAtlasSource]), and the scene for a [TileSetScenesCollectionSource].
If [param source_id] is set to [code]-1[/code], [param atlas_coords] to [code]Vector2i(-1, -1)[/code] or [param alternative_tile] to [code]-1[/code], the cell will be erased. An erased cell gets [b]all[/b] its identifiers automatically set to their respective invalid values, namely [code]-1[/code], [code]Vector2i(-1, -1)[/code] and [code]-1[/code].
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) SetCell(layer gd.Int, coords gd.Vector2i, source_id gd.Int, atlas_coords gd.Vector2i, alternative_tile gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, coords)
	callframe.Arg(frame, source_id)
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, alternative_tile)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_set_cell, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Erases the cell on layer [param layer] at coordinates [param coords].
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) EraseCell(layer gd.Int, coords gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, coords)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_erase_cell, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the tile source ID of the cell on layer [param layer] at coordinates [param coords]. Returns [code]-1[/code] if the cell does not exist.
If [param use_proxies] is [code]false[/code], ignores the [TileSet]'s tile proxies, returning the raw source identifier. See [method TileSet.map_tile_proxy].
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) GetCellSourceId(layer gd.Int, coords gd.Vector2i, use_proxies bool) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, coords)
	callframe.Arg(frame, use_proxies)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_get_cell_source_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the tile atlas coordinates ID of the cell on layer [param layer] at coordinates [param coords]. Returns [code]Vector2i(-1, -1)[/code] if the cell does not exist.
If [param use_proxies] is [code]false[/code], ignores the [TileSet]'s tile proxies, returning the raw atlas coordinate identifier. See [method TileSet.map_tile_proxy].
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) GetCellAtlasCoords(layer gd.Int, coords gd.Vector2i, use_proxies bool) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, coords)
	callframe.Arg(frame, use_proxies)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_get_cell_atlas_coords, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the tile alternative ID of the cell on layer [param layer] at [param coords].
If [param use_proxies] is [code]false[/code], ignores the [TileSet]'s tile proxies, returning the raw alternative identifier. See [method TileSet.map_tile_proxy].
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) GetCellAlternativeTile(layer gd.Int, coords gd.Vector2i, use_proxies bool) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, coords)
	callframe.Arg(frame, use_proxies)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_get_cell_alternative_tile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [TileData] object associated with the given cell, or [code]null[/code] if the cell does not exist or is not a [TileSetAtlasSource].
If [param layer] is negative, the layers are accessed from the last one.
[codeblock]
func get_clicked_tile_power():
    var clicked_cell = tile_map.local_to_map(tile_map.get_local_mouse_position())
    var data = tile_map.get_cell_tile_data(0, clicked_cell)
    if data:
        return data.get_custom_data("power")
    else:
        return 0
[/codeblock]
If [param use_proxies] is [code]false[/code], ignores the [TileSet]'s tile proxies. See [method TileSet.map_tile_proxy].
*/
//go:nosplit
func (self class) GetCellTileData(ctx gd.Lifetime, layer gd.Int, coords gd.Vector2i, use_proxies bool) object.TileData {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, coords)
	callframe.Arg(frame, use_proxies)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_get_cell_tile_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TileData
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the coordinates of the tile for given physics body RID. Such RID can be retrieved from [method KinematicCollision2D.get_collider_rid], when colliding with a tile.
*/
//go:nosplit
func (self class) GetCoordsForBodyRid(body gd.RID) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_get_coords_for_body_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the tilemap layer of the tile for given physics body RID. Such RID can be retrieved from [method KinematicCollision2D.get_collider_rid], when colliding with a tile.
*/
//go:nosplit
func (self class) GetLayerForBodyRid(body gd.RID) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_get_layer_for_body_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a new [TileMapPattern] from the given layer and set of cells.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) GetPattern(ctx gd.Lifetime, layer gd.Int, coords_array gd.ArrayOf[gd.Vector2i]) object.TileMapPattern {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, mmm.Get(coords_array))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_get_pattern, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.TileMapPattern
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns for the given coordinate [param coords_in_pattern] in a [TileMapPattern] the corresponding cell coordinates if the pattern was pasted at the [param position_in_tilemap] coordinates (see [method set_pattern]). This mapping is required as in half-offset tile shapes, the mapping might not work by calculating [code]position_in_tile_map + coords_in_pattern[/code].
*/
//go:nosplit
func (self class) MapPattern(position_in_tilemap gd.Vector2i, coords_in_pattern gd.Vector2i, pattern object.TileMapPattern) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position_in_tilemap)
	callframe.Arg(frame, coords_in_pattern)
	callframe.Arg(frame, mmm.Get(pattern[0].AsPointer())[0])
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_map_pattern, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Paste the given [TileMapPattern] at the given [param position] and [param layer] in the tile map.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) SetPattern(layer gd.Int, position gd.Vector2i, pattern object.TileMapPattern)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, position)
	callframe.Arg(frame, mmm.Get(pattern[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_set_pattern, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Update all the cells in the [param cells] coordinates array so that they use the given [param terrain] for the given [param terrain_set]. If an updated cell has the same terrain as one of its neighboring cells, this function tries to join the two. This function might update neighboring tiles if needed to create correct terrain transitions.
If [param ignore_empty_terrains] is true, empty terrains will be ignored when trying to find the best fitting tile for the given terrain constraints.
If [param layer] is negative, the layers are accessed from the last one.
[b]Note:[/b] To work correctly, this method requires the TileMap's TileSet to have terrains set up with all required terrain combinations. Otherwise, it may produce unexpected results.
*/
//go:nosplit
func (self class) SetCellsTerrainConnect(layer gd.Int, cells gd.ArrayOf[gd.Vector2i], terrain_set gd.Int, terrain gd.Int, ignore_empty_terrains bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, mmm.Get(cells))
	callframe.Arg(frame, terrain_set)
	callframe.Arg(frame, terrain)
	callframe.Arg(frame, ignore_empty_terrains)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_set_cells_terrain_connect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Update all the cells in the [param path] coordinates array so that they use the given [param terrain] for the given [param terrain_set]. The function will also connect two successive cell in the path with the same terrain. This function might update neighboring tiles if needed to create correct terrain transitions.
If [param ignore_empty_terrains] is true, empty terrains will be ignored when trying to find the best fitting tile for the given terrain constraints.
If [param layer] is negative, the layers are accessed from the last one.
[b]Note:[/b] To work correctly, this method requires the TileMap's TileSet to have terrains set up with all required terrain combinations. Otherwise, it may produce unexpected results.
*/
//go:nosplit
func (self class) SetCellsTerrainPath(layer gd.Int, path gd.ArrayOf[gd.Vector2i], terrain_set gd.Int, terrain gd.Int, ignore_empty_terrains bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, terrain_set)
	callframe.Arg(frame, terrain)
	callframe.Arg(frame, ignore_empty_terrains)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_set_cells_terrain_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Clears cells that do not exist in the tileset.
*/
//go:nosplit
func (self class) FixInvalidTiles()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_fix_invalid_tiles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Clears all cells on the given layer.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) ClearLayer(layer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_clear_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Clears all cells.
*/
//go:nosplit
func (self class) Clear()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Triggers a direct update of the TileMap. Usually, calling this function is not needed, as TileMap node updates automatically when one of its properties or cells is modified.
However, for performance reasons, those updates are batched and delayed to the end of the frame. Calling this function will force the TileMap to update right away instead.
[b]Warning:[/b] Updating the TileMap is computationally expensive and may impact performance. Try to limit the number of updates and how many tiles they impact.
*/
//go:nosplit
func (self class) UpdateInternals()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_update_internals, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Notifies the TileMap node that calls to [method _use_tile_data_runtime_update] or [method _tile_data_runtime_update] will lead to different results. This will thus trigger a TileMap update.
If [param layer] is provided, only notifies changes for the given layer. Providing the [param layer] argument (when applicable) is usually preferred for performance reasons.
[b]Warning:[/b] Updating the TileMap is computationally expensive and may impact performance. Try to limit the number of calls to this function to avoid unnecessary update.
[b]Note:[/b] This does not trigger a direct update of the TileMap, the update will be done at the end of the frame as usual (unless you call [method update_internals]).
*/
//go:nosplit
func (self class) NotifyRuntimeTileDataUpdate(layer gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_notify_runtime_tile_data_update, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the list of all neighbourings cells to the one at [param coords].
*/
//go:nosplit
func (self class) GetSurroundingCells(ctx gd.Lifetime, coords gd.Vector2i) gd.ArrayOf[gd.Vector2i] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_get_surrounding_cells, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Vector2i](ret)
}
/*
Returns a [Vector2i] array with the positions of all cells containing a tile in the given layer. A cell is considered empty if its source identifier equals -1, its atlas coordinates identifiers is [code]Vector2(-1, -1)[/code] and its alternative identifier is -1.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) GetUsedCells(ctx gd.Lifetime, layer gd.Int) gd.ArrayOf[gd.Vector2i] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_get_used_cells, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Vector2i](ret)
}
/*
Returns a [Vector2i] array with the positions of all cells containing a tile in the given layer. Tiles may be filtered according to their source ([param source_id]), their atlas coordinates ([param atlas_coords]) or alternative id ([param alternative_tile]).
If a parameter has its value set to the default one, this parameter is not used to filter a cell. Thus, if all parameters have their respective default value, this method returns the same result as [method get_used_cells].
A cell is considered empty if its source identifier equals -1, its atlas coordinates identifiers is [code]Vector2(-1, -1)[/code] and its alternative identifier is -1.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) GetUsedCellsById(ctx gd.Lifetime, layer gd.Int, source_id gd.Int, atlas_coords gd.Vector2i, alternative_tile gd.Int) gd.ArrayOf[gd.Vector2i] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, source_id)
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, alternative_tile)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_get_used_cells_by_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gd.Vector2i](ret)
}
/*
Returns a rectangle enclosing the used (non-empty) tiles of the map, including all layers.
*/
//go:nosplit
func (self class) GetUsedRect() gd.Rect2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_get_used_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the centered position of a cell in the TileMap's local coordinate space. To convert the returned value into global coordinates, use [method Node2D.to_global]. See also [method local_to_map].
[b]Note:[/b] This may not correspond to the visual position of the tile, i.e. it ignores the [member TileData.texture_origin] property of individual tiles.
*/
//go:nosplit
func (self class) MapToLocal(map_position gd.Vector2i) gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, map_position)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_map_to_local, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the map coordinates of the cell containing the given [param local_position]. If [param local_position] is in global coordinates, consider using [method Node2D.to_local] before passing it to this method. See also [method map_to_local].
*/
//go:nosplit
func (self class) LocalToMap(local_position gd.Vector2) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, local_position)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_local_to_map, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the neighboring cell to the one at coordinates [param coords], identified by the [param neighbor] direction. This method takes into account the different layouts a TileMap can take.
*/
//go:nosplit
func (self class) GetNeighborCell(coords gd.Vector2i, neighbor classdb.TileSetCellNeighbor) gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	callframe.Arg(frame, neighbor)
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.TileMap.Bind_get_neighbor_cell, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsTileMap() Expert { return self[0].AsTileMap() }


//go:nosplit
func (self Simple) AsTileMap() Simple { return self[0].AsTileMap() }


//go:nosplit
func (self class) AsNode2D() Node2D.Expert { return self[0].AsNode2D() }


//go:nosplit
func (self Simple) AsNode2D() Node2D.Simple { return self[0].AsNode2D() }


//go:nosplit
func (self class) AsCanvasItem() CanvasItem.Expert { return self[0].AsCanvasItem() }


//go:nosplit
func (self Simple) AsCanvasItem() CanvasItem.Simple { return self[0].AsCanvasItem() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_use_tile_data_runtime_update": return reflect.ValueOf(self._use_tile_data_runtime_update);
	case "_tile_data_runtime_update": return reflect.ValueOf(self._tile_data_runtime_update);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	case "_use_tile_data_runtime_update": return reflect.ValueOf(self._use_tile_data_runtime_update);
	case "_tile_data_runtime_update": return reflect.ValueOf(self._tile_data_runtime_update);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("TileMap", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type VisibilityMode = classdb.TileMapVisibilityMode

const (
/*Use the debug settings to determine visibility.*/
	VisibilityModeDefault VisibilityMode = 0
/*Always hide.*/
	VisibilityModeForceHide VisibilityMode = 2
/*Always show.*/
	VisibilityModeForceShow VisibilityMode = 1
)
