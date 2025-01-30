// Package TileMap provides methods for working with TileMap object instances.
package TileMap

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/classdb/Node2D"
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
import "graphics.gd/variant/Rect2i"
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
Node for 2D tile-based maps. Tilemaps use a [TileSet] which contain a list of tiles which are used to create grid-based maps. A TileMap may have several layers, layouting tiles on top of each other.
For performance reasons, all TileMap updates are batched at the end of a frame. Notably, this means that scene tiles from a [TileSetScenesCollectionSource] may be initialized after their parent. This is only queued when inside the scene tree.
To force an update earlier on, call [method update_internals].

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=TileMap)
*/
type Instance [1]gdclass.TileMap

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsTileMap() Instance
}
type Interface interface {
	//Should return [code]true[/code] if the tile at coordinates [param coords] on layer [param layer] requires a runtime update.
	//[b]Warning:[/b] Make sure this function only return [code]true[/code] when needed. Any tile processed at runtime without a need for it will imply a significant performance penalty.
	//[b]Note:[/b] If the result of this function should changed, use [method notify_runtime_tile_data_update] to notify the TileMap it needs an update.
	UseTileDataRuntimeUpdate(layer int, coords Vector2i.XY) bool
	//Called with a TileData object about to be used internally by the TileMap, allowing its modification at runtime.
	//This method is only called if [method _use_tile_data_runtime_update] is implemented and returns [code]true[/code] for the given tile [param coords] and [param layer].
	//[b]Warning:[/b] The [param tile_data] object's sub-resources are the same as the one in the TileSet. Modifying them might impact the whole TileSet. Instead, make sure to duplicate those resources.
	//[b]Note:[/b] If the properties of [param tile_data] object should change over time, use [method notify_runtime_tile_data_update] to notify the TileMap it needs an update.
	TileDataRuntimeUpdate(layer int, coords Vector2i.XY, tile_data [1]gdclass.TileData)
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) UseTileDataRuntimeUpdate(layer int, coords Vector2i.XY) (_ bool) { return }
func (self implementation) TileDataRuntimeUpdate(layer int, coords Vector2i.XY, tile_data [1]gdclass.TileData) {
	return
}

/*
Should return [code]true[/code] if the tile at coordinates [param coords] on layer [param layer] requires a runtime update.
[b]Warning:[/b] Make sure this function only return [code]true[/code] when needed. Any tile processed at runtime without a need for it will imply a significant performance penalty.
[b]Note:[/b] If the result of this function should changed, use [method notify_runtime_tile_data_update] to notify the TileMap it needs an update.
*/
func (Instance) _use_tile_data_runtime_update(impl func(ptr unsafe.Pointer, layer int, coords Vector2i.XY) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var layer = gd.UnsafeGet[int64](p_args, 0)

		var coords = gd.UnsafeGet[Vector2i.XY](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(layer), coords)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called with a TileData object about to be used internally by the TileMap, allowing its modification at runtime.
This method is only called if [method _use_tile_data_runtime_update] is implemented and returns [code]true[/code] for the given tile [param coords] and [param layer].
[b]Warning:[/b] The [param tile_data] object's sub-resources are the same as the one in the TileSet. Modifying them might impact the whole TileSet. Instead, make sure to duplicate those resources.
[b]Note:[/b] If the properties of [param tile_data] object should change over time, use [method notify_runtime_tile_data_update] to notify the TileMap it needs an update.
*/
func (Instance) _tile_data_runtime_update(impl func(ptr unsafe.Pointer, layer int, coords Vector2i.XY, tile_data [1]gdclass.TileData)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var layer = gd.UnsafeGet[int64](p_args, 0)

		var coords = gd.UnsafeGet[Vector2i.XY](p_args, 1)

		var tile_data = [1]gdclass.TileData{pointers.New[gdclass.TileData]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 2))})}

		defer pointers.End(tile_data[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(layer), coords, tile_data)
	}
}

/*
Assigns [param map] as a [NavigationServer2D] navigation map for the specified TileMap layer [param layer].
*/
func (self Instance) SetNavigationMap(layer int, mapping RID.NavigationMap2D) { //gd:TileMap.set_navigation_map
	class(self).SetNavigationMap(int64(layer), RID.Any(mapping))
}

/*
Returns the [RID] of the [NavigationServer2D] navigation map assigned to the specified TileMap layer [param layer].
*/
func (self Instance) GetNavigationMap(layer int) RID.NavigationMap2D { //gd:TileMap.get_navigation_map
	return RID.NavigationMap2D(class(self).GetNavigationMap(int64(layer)))
}

/*
Forces the TileMap and the layer [param layer] to update.
*/
func (self Instance) ForceUpdate() { //gd:TileMap.force_update
	class(self).ForceUpdate(int64(-1))
}

/*
Returns the number of layers in the TileMap.
*/
func (self Instance) GetLayersCount() int { //gd:TileMap.get_layers_count
	return int(int(class(self).GetLayersCount()))
}

/*
Adds a layer at the given position [param to_position] in the array. If [param to_position] is negative, the position is counted from the end, with [code]-1[/code] adding the layer at the end of the array.
*/
func (self Instance) AddLayer(to_position int) { //gd:TileMap.add_layer
	class(self).AddLayer(int64(to_position))
}

/*
Moves the layer at index [param layer] to the given position [param to_position] in the array.
*/
func (self Instance) MoveLayer(layer int, to_position int) { //gd:TileMap.move_layer
	class(self).MoveLayer(int64(layer), int64(to_position))
}

/*
Removes the layer at index [param layer].
*/
func (self Instance) RemoveLayer(layer int) { //gd:TileMap.remove_layer
	class(self).RemoveLayer(int64(layer))
}

/*
Sets a layer's name. This is mostly useful in the editor.
If [param layer] is negative, the layers are accessed from the last one.
*/
func (self Instance) SetLayerName(layer int, name string) { //gd:TileMap.set_layer_name
	class(self).SetLayerName(int64(layer), String.New(name))
}

/*
Returns a TileMap layer's name.
If [param layer] is negative, the layers are accessed from the last one.
*/
func (self Instance) GetLayerName(layer int) string { //gd:TileMap.get_layer_name
	return string(class(self).GetLayerName(int64(layer)).String())
}

/*
Enables or disables the layer [param layer]. A disabled layer is not processed at all (no rendering, no physics, etc.).
If [param layer] is negative, the layers are accessed from the last one.
*/
func (self Instance) SetLayerEnabled(layer int, enabled bool) { //gd:TileMap.set_layer_enabled
	class(self).SetLayerEnabled(int64(layer), enabled)
}

/*
Returns if a layer is enabled.
If [param layer] is negative, the layers are accessed from the last one.
*/
func (self Instance) IsLayerEnabled(layer int) bool { //gd:TileMap.is_layer_enabled
	return bool(class(self).IsLayerEnabled(int64(layer)))
}

/*
Sets a layer's color. It will be multiplied by tile's color and TileMap's modulate.
If [param layer] is negative, the layers are accessed from the last one.
*/
func (self Instance) SetLayerModulate(layer int, modulate Color.RGBA) { //gd:TileMap.set_layer_modulate
	class(self).SetLayerModulate(int64(layer), Color.RGBA(modulate))
}

/*
Returns a TileMap layer's modulate.
If [param layer] is negative, the layers are accessed from the last one.
*/
func (self Instance) GetLayerModulate(layer int) Color.RGBA { //gd:TileMap.get_layer_modulate
	return Color.RGBA(class(self).GetLayerModulate(int64(layer)))
}

/*
Enables or disables a layer's Y-sorting. If a layer is Y-sorted, the layer will behave as a CanvasItem node where each of its tile gets Y-sorted.
Y-sorted layers should usually be on different Z-index values than not Y-sorted layers, otherwise, each of those layer will be Y-sorted as whole with the Y-sorted one. This is usually an undesired behavior.
If [param layer] is negative, the layers are accessed from the last one.
*/
func (self Instance) SetLayerYSortEnabled(layer int, y_sort_enabled bool) { //gd:TileMap.set_layer_y_sort_enabled
	class(self).SetLayerYSortEnabled(int64(layer), y_sort_enabled)
}

/*
Returns if a layer Y-sorts its tiles.
If [param layer] is negative, the layers are accessed from the last one.
*/
func (self Instance) IsLayerYSortEnabled(layer int) bool { //gd:TileMap.is_layer_y_sort_enabled
	return bool(class(self).IsLayerYSortEnabled(int64(layer)))
}

/*
Sets a layer's Y-sort origin value. This Y-sort origin value is added to each tile's Y-sort origin value.
This allows, for example, to fake a different height level on each layer. This can be useful for top-down view games.
If [param layer] is negative, the layers are accessed from the last one.
*/
func (self Instance) SetLayerYSortOrigin(layer int, y_sort_origin int) { //gd:TileMap.set_layer_y_sort_origin
	class(self).SetLayerYSortOrigin(int64(layer), int64(y_sort_origin))
}

/*
Returns a TileMap layer's Y sort origin.
If [param layer] is negative, the layers are accessed from the last one.
*/
func (self Instance) GetLayerYSortOrigin(layer int) int { //gd:TileMap.get_layer_y_sort_origin
	return int(int(class(self).GetLayerYSortOrigin(int64(layer))))
}

/*
Sets a layers Z-index value. This Z-index is added to each tile's Z-index value.
If [param layer] is negative, the layers are accessed from the last one.
*/
func (self Instance) SetLayerZIndex(layer int, z_index int) { //gd:TileMap.set_layer_z_index
	class(self).SetLayerZIndex(int64(layer), int64(z_index))
}

/*
Returns a TileMap layer's Z-index value.
If [param layer] is negative, the layers are accessed from the last one.
*/
func (self Instance) GetLayerZIndex(layer int) int { //gd:TileMap.get_layer_z_index
	return int(int(class(self).GetLayerZIndex(int64(layer))))
}

/*
Enables or disables a layer's built-in navigation regions generation. Disable this if you need to bake navigation regions from a TileMap using a [NavigationRegion2D] node.
*/
func (self Instance) SetLayerNavigationEnabled(layer int, enabled bool) { //gd:TileMap.set_layer_navigation_enabled
	class(self).SetLayerNavigationEnabled(int64(layer), enabled)
}

/*
Returns if a layer's built-in navigation regions generation is enabled.
*/
func (self Instance) IsLayerNavigationEnabled(layer int) bool { //gd:TileMap.is_layer_navigation_enabled
	return bool(class(self).IsLayerNavigationEnabled(int64(layer)))
}

/*
Assigns [param map] as a [NavigationServer2D] navigation map for the specified TileMap layer [param layer].
By default the TileMap uses the default [World2D] navigation map for the first TileMap layer. For each additional TileMap layer a new navigation map is created for the additional layer.
In order to make [NavigationAgent2D] switch between TileMap layer navigation maps use [method NavigationAgent2D.set_navigation_map] with the navigation map received from [method get_layer_navigation_map].
If [param layer] is negative, the layers are accessed from the last one.
*/
func (self Instance) SetLayerNavigationMap(layer int, mapping RID.NavigationMap2D) { //gd:TileMap.set_layer_navigation_map
	class(self).SetLayerNavigationMap(int64(layer), RID.Any(mapping))
}

/*
Returns the [RID] of the [NavigationServer2D] navigation map assigned to the specified TileMap layer [param layer].
By default the TileMap uses the default [World2D] navigation map for the first TileMap layer. For each additional TileMap layer a new navigation map is created for the additional layer.
In order to make [NavigationAgent2D] switch between TileMap layer navigation maps use [method NavigationAgent2D.set_navigation_map] with the navigation map received from [method get_layer_navigation_map].
If [param layer] is negative, the layers are accessed from the last one.
*/
func (self Instance) GetLayerNavigationMap(layer int) RID.NavigationMap2D { //gd:TileMap.get_layer_navigation_map
	return RID.NavigationMap2D(class(self).GetLayerNavigationMap(int64(layer)))
}

/*
Sets the tile identifiers for the cell on layer [param layer] at coordinates [param coords]. Each tile of the [TileSet] is identified using three parts:
- The source identifier [param source_id] identifies a [TileSetSource] identifier. See [method TileSet.set_source_id],
- The atlas coordinates identifier [param atlas_coords] identifies a tile coordinates in the atlas (if the source is a [TileSetAtlasSource]). For [TileSetScenesCollectionSource] it should always be [code]Vector2i(0, 0)[/code]),
- The alternative tile identifier [param alternative_tile] identifies a tile alternative in the atlas (if the source is a [TileSetAtlasSource]), and the scene for a [TileSetScenesCollectionSource].
If [param source_id] is set to [code]-1[/code], [param atlas_coords] to [code]Vector2i(-1, -1)[/code] or [param alternative_tile] to [code]-1[/code], the cell will be erased. An erased cell gets [b]all[/b] its identifiers automatically set to their respective invalid values, namely [code]-1[/code], [code]Vector2i(-1, -1)[/code] and [code]-1[/code].
If [param layer] is negative, the layers are accessed from the last one.
*/
func (self Instance) SetCell(layer int, coords Vector2i.XY) { //gd:TileMap.set_cell
	class(self).SetCell(int64(layer), Vector2i.XY(coords), int64(-1), Vector2i.XY(gd.Vector2i{-1, -1}), int64(0))
}

/*
Erases the cell on layer [param layer] at coordinates [param coords].
If [param layer] is negative, the layers are accessed from the last one.
*/
func (self Instance) EraseCell(layer int, coords Vector2i.XY) { //gd:TileMap.erase_cell
	class(self).EraseCell(int64(layer), Vector2i.XY(coords))
}

/*
Returns the tile source ID of the cell on layer [param layer] at coordinates [param coords]. Returns [code]-1[/code] if the cell does not exist.
If [param use_proxies] is [code]false[/code], ignores the [TileSet]'s tile proxies, returning the raw source identifier. See [method TileSet.map_tile_proxy].
If [param layer] is negative, the layers are accessed from the last one.
*/
func (self Instance) GetCellSourceId(layer int, coords Vector2i.XY) int { //gd:TileMap.get_cell_source_id
	return int(int(class(self).GetCellSourceId(int64(layer), Vector2i.XY(coords), false)))
}

/*
Returns the tile atlas coordinates ID of the cell on layer [param layer] at coordinates [param coords]. Returns [code]Vector2i(-1, -1)[/code] if the cell does not exist.
If [param use_proxies] is [code]false[/code], ignores the [TileSet]'s tile proxies, returning the raw atlas coordinate identifier. See [method TileSet.map_tile_proxy].
If [param layer] is negative, the layers are accessed from the last one.
*/
func (self Instance) GetCellAtlasCoords(layer int, coords Vector2i.XY) Vector2i.XY { //gd:TileMap.get_cell_atlas_coords
	return Vector2i.XY(class(self).GetCellAtlasCoords(int64(layer), Vector2i.XY(coords), false))
}

/*
Returns the tile alternative ID of the cell on layer [param layer] at [param coords].
If [param use_proxies] is [code]false[/code], ignores the [TileSet]'s tile proxies, returning the raw alternative identifier. See [method TileSet.map_tile_proxy].
If [param layer] is negative, the layers are accessed from the last one.
*/
func (self Instance) GetCellAlternativeTile(layer int, coords Vector2i.XY) int { //gd:TileMap.get_cell_alternative_tile
	return int(int(class(self).GetCellAlternativeTile(int64(layer), Vector2i.XY(coords), false)))
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
func (self Instance) GetCellTileData(layer int, coords Vector2i.XY) [1]gdclass.TileData { //gd:TileMap.get_cell_tile_data
	return [1]gdclass.TileData(class(self).GetCellTileData(int64(layer), Vector2i.XY(coords), false))
}

/*
Returns the coordinates of the tile for given physics body RID. Such RID can be retrieved from [method KinematicCollision2D.get_collider_rid], when colliding with a tile.
*/
func (self Instance) GetCoordsForBodyRid(body RID.Body2D) Vector2i.XY { //gd:TileMap.get_coords_for_body_rid
	return Vector2i.XY(class(self).GetCoordsForBodyRid(RID.Any(body)))
}

/*
Returns the tilemap layer of the tile for given physics body RID. Such RID can be retrieved from [method KinematicCollision2D.get_collider_rid], when colliding with a tile.
*/
func (self Instance) GetLayerForBodyRid(body RID.Body2D) int { //gd:TileMap.get_layer_for_body_rid
	return int(int(class(self).GetLayerForBodyRid(RID.Any(body))))
}

/*
Creates a new [TileMapPattern] from the given layer and set of cells.
If [param layer] is negative, the layers are accessed from the last one.
*/
func (self Instance) GetPattern(layer int, coords_array []Vector2i.XY) [1]gdclass.TileMapPattern { //gd:TileMap.get_pattern
	return [1]gdclass.TileMapPattern(class(self).GetPattern(int64(layer), gd.ArrayFromSlice[Array.Contains[Vector2i.XY]](coords_array)))
}

/*
Returns for the given coordinate [param coords_in_pattern] in a [TileMapPattern] the corresponding cell coordinates if the pattern was pasted at the [param position_in_tilemap] coordinates (see [method set_pattern]). This mapping is required as in half-offset tile shapes, the mapping might not work by calculating [code]position_in_tile_map + coords_in_pattern[/code].
*/
func (self Instance) MapPattern(position_in_tilemap Vector2i.XY, coords_in_pattern Vector2i.XY, pattern [1]gdclass.TileMapPattern) Vector2i.XY { //gd:TileMap.map_pattern
	return Vector2i.XY(class(self).MapPattern(Vector2i.XY(position_in_tilemap), Vector2i.XY(coords_in_pattern), pattern))
}

/*
Paste the given [TileMapPattern] at the given [param position] and [param layer] in the tile map.
If [param layer] is negative, the layers are accessed from the last one.
*/
func (self Instance) SetPattern(layer int, position Vector2i.XY, pattern [1]gdclass.TileMapPattern) { //gd:TileMap.set_pattern
	class(self).SetPattern(int64(layer), Vector2i.XY(position), pattern)
}

/*
Update all the cells in the [param cells] coordinates array so that they use the given [param terrain] for the given [param terrain_set]. If an updated cell has the same terrain as one of its neighboring cells, this function tries to join the two. This function might update neighboring tiles if needed to create correct terrain transitions.
If [param ignore_empty_terrains] is true, empty terrains will be ignored when trying to find the best fitting tile for the given terrain constraints.
If [param layer] is negative, the layers are accessed from the last one.
[b]Note:[/b] To work correctly, this method requires the TileMap's TileSet to have terrains set up with all required terrain combinations. Otherwise, it may produce unexpected results.
*/
func (self Instance) SetCellsTerrainConnect(layer int, cells []Vector2i.XY, terrain_set int, terrain int) { //gd:TileMap.set_cells_terrain_connect
	class(self).SetCellsTerrainConnect(int64(layer), gd.ArrayFromSlice[Array.Contains[Vector2i.XY]](cells), int64(terrain_set), int64(terrain), true)
}

/*
Update all the cells in the [param path] coordinates array so that they use the given [param terrain] for the given [param terrain_set]. The function will also connect two successive cell in the path with the same terrain. This function might update neighboring tiles if needed to create correct terrain transitions.
If [param ignore_empty_terrains] is true, empty terrains will be ignored when trying to find the best fitting tile for the given terrain constraints.
If [param layer] is negative, the layers are accessed from the last one.
[b]Note:[/b] To work correctly, this method requires the TileMap's TileSet to have terrains set up with all required terrain combinations. Otherwise, it may produce unexpected results.
*/
func (self Instance) SetCellsTerrainPath(layer int, path []Vector2i.XY, terrain_set int, terrain int) { //gd:TileMap.set_cells_terrain_path
	class(self).SetCellsTerrainPath(int64(layer), gd.ArrayFromSlice[Array.Contains[Vector2i.XY]](path), int64(terrain_set), int64(terrain), true)
}

/*
Clears cells that do not exist in the tileset.
*/
func (self Instance) FixInvalidTiles() { //gd:TileMap.fix_invalid_tiles
	class(self).FixInvalidTiles()
}

/*
Clears all cells on the given layer.
If [param layer] is negative, the layers are accessed from the last one.
*/
func (self Instance) ClearLayer(layer int) { //gd:TileMap.clear_layer
	class(self).ClearLayer(int64(layer))
}

/*
Clears all cells.
*/
func (self Instance) Clear() { //gd:TileMap.clear
	class(self).Clear()
}

/*
Triggers a direct update of the TileMap. Usually, calling this function is not needed, as TileMap node updates automatically when one of its properties or cells is modified.
However, for performance reasons, those updates are batched and delayed to the end of the frame. Calling this function will force the TileMap to update right away instead.
[b]Warning:[/b] Updating the TileMap is computationally expensive and may impact performance. Try to limit the number of updates and how many tiles they impact.
*/
func (self Instance) UpdateInternals() { //gd:TileMap.update_internals
	class(self).UpdateInternals()
}

/*
Notifies the TileMap node that calls to [method _use_tile_data_runtime_update] or [method _tile_data_runtime_update] will lead to different results. This will thus trigger a TileMap update.
If [param layer] is provided, only notifies changes for the given layer. Providing the [param layer] argument (when applicable) is usually preferred for performance reasons.
[b]Warning:[/b] Updating the TileMap is computationally expensive and may impact performance. Try to limit the number of calls to this function to avoid unnecessary update.
[b]Note:[/b] This does not trigger a direct update of the TileMap, the update will be done at the end of the frame as usual (unless you call [method update_internals]).
*/
func (self Instance) NotifyRuntimeTileDataUpdate() { //gd:TileMap.notify_runtime_tile_data_update
	class(self).NotifyRuntimeTileDataUpdate(int64(-1))
}

/*
Returns the list of all neighbourings cells to the one at [param coords].
*/
func (self Instance) GetSurroundingCells(coords Vector2i.XY) []Vector2i.XY { //gd:TileMap.get_surrounding_cells
	return []Vector2i.XY(gd.ArrayAs[[]Vector2i.XY](gd.InternalArray(class(self).GetSurroundingCells(Vector2i.XY(coords)))))
}

/*
Returns a [Vector2i] array with the positions of all cells containing a tile in the given layer. A cell is considered empty if its source identifier equals -1, its atlas coordinates identifiers is [code]Vector2(-1, -1)[/code] and its alternative identifier is -1.
If [param layer] is negative, the layers are accessed from the last one.
*/
func (self Instance) GetUsedCells(layer int) []Vector2i.XY { //gd:TileMap.get_used_cells
	return []Vector2i.XY(gd.ArrayAs[[]Vector2i.XY](gd.InternalArray(class(self).GetUsedCells(int64(layer)))))
}

/*
Returns a [Vector2i] array with the positions of all cells containing a tile in the given layer. Tiles may be filtered according to their source ([param source_id]), their atlas coordinates ([param atlas_coords]) or alternative id ([param alternative_tile]).
If a parameter has its value set to the default one, this parameter is not used to filter a cell. Thus, if all parameters have their respective default value, this method returns the same result as [method get_used_cells].
A cell is considered empty if its source identifier equals -1, its atlas coordinates identifiers is [code]Vector2(-1, -1)[/code] and its alternative identifier is -1.
If [param layer] is negative, the layers are accessed from the last one.
*/
func (self Instance) GetUsedCellsById(layer int) []Vector2i.XY { //gd:TileMap.get_used_cells_by_id
	return []Vector2i.XY(gd.ArrayAs[[]Vector2i.XY](gd.InternalArray(class(self).GetUsedCellsById(int64(layer), int64(-1), Vector2i.XY(gd.Vector2i{-1, -1}), int64(-1)))))
}

/*
Returns a rectangle enclosing the used (non-empty) tiles of the map, including all layers.
*/
func (self Instance) GetUsedRect() Rect2i.PositionSize { //gd:TileMap.get_used_rect
	return Rect2i.PositionSize(class(self).GetUsedRect())
}

/*
Returns the centered position of a cell in the TileMap's local coordinate space. To convert the returned value into global coordinates, use [method Node2D.to_global]. See also [method local_to_map].
[b]Note:[/b] This may not correspond to the visual position of the tile, i.e. it ignores the [member TileData.texture_origin] property of individual tiles.
*/
func (self Instance) MapToLocal(map_position Vector2i.XY) Vector2.XY { //gd:TileMap.map_to_local
	return Vector2.XY(class(self).MapToLocal(Vector2i.XY(map_position)))
}

/*
Returns the map coordinates of the cell containing the given [param local_position]. If [param local_position] is in global coordinates, consider using [method Node2D.to_local] before passing it to this method. See also [method map_to_local].
*/
func (self Instance) LocalToMap(local_position Vector2.XY) Vector2i.XY { //gd:TileMap.local_to_map
	return Vector2i.XY(class(self).LocalToMap(Vector2.XY(local_position)))
}

/*
Returns the neighboring cell to the one at coordinates [param coords], identified by the [param neighbor] direction. This method takes into account the different layouts a TileMap can take.
*/
func (self Instance) GetNeighborCell(coords Vector2i.XY, neighbor gdclass.TileSetCellNeighbor) Vector2i.XY { //gd:TileMap.get_neighbor_cell
	return Vector2i.XY(class(self).GetNeighborCell(Vector2i.XY(coords), neighbor))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.TileMap

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TileMap"))
	casted := Instance{*(*gdclass.TileMap)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) TileSet() [1]gdclass.TileSet {
	return [1]gdclass.TileSet(class(self).GetTileset())
}

func (self Instance) SetTileSet(value [1]gdclass.TileSet) {
	class(self).SetTileset(value)
}

func (self Instance) RenderingQuadrantSize() int {
	return int(int(class(self).GetRenderingQuadrantSize()))
}

func (self Instance) SetRenderingQuadrantSize(value int) {
	class(self).SetRenderingQuadrantSize(int64(value))
}

func (self Instance) CollisionAnimatable() bool {
	return bool(class(self).IsCollisionAnimatable())
}

func (self Instance) SetCollisionAnimatable(value bool) {
	class(self).SetCollisionAnimatable(value)
}

func (self Instance) CollisionVisibilityMode() gdclass.TileMapVisibilityMode {
	return gdclass.TileMapVisibilityMode(class(self).GetCollisionVisibilityMode())
}

func (self Instance) SetCollisionVisibilityMode(value gdclass.TileMapVisibilityMode) {
	class(self).SetCollisionVisibilityMode(value)
}

func (self Instance) NavigationVisibilityMode() gdclass.TileMapVisibilityMode {
	return gdclass.TileMapVisibilityMode(class(self).GetNavigationVisibilityMode())
}

func (self Instance) SetNavigationVisibilityMode(value gdclass.TileMapVisibilityMode) {
	class(self).SetNavigationVisibilityMode(value)
}

/*
Should return [code]true[/code] if the tile at coordinates [param coords] on layer [param layer] requires a runtime update.
[b]Warning:[/b] Make sure this function only return [code]true[/code] when needed. Any tile processed at runtime without a need for it will imply a significant performance penalty.
[b]Note:[/b] If the result of this function should changed, use [method notify_runtime_tile_data_update] to notify the TileMap it needs an update.
*/
func (class) _use_tile_data_runtime_update(impl func(ptr unsafe.Pointer, layer int64, coords Vector2i.XY) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var layer = gd.UnsafeGet[int64](p_args, 0)

		var coords = gd.UnsafeGet[Vector2i.XY](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, layer, coords)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called with a TileData object about to be used internally by the TileMap, allowing its modification at runtime.
This method is only called if [method _use_tile_data_runtime_update] is implemented and returns [code]true[/code] for the given tile [param coords] and [param layer].
[b]Warning:[/b] The [param tile_data] object's sub-resources are the same as the one in the TileSet. Modifying them might impact the whole TileSet. Instead, make sure to duplicate those resources.
[b]Note:[/b] If the properties of [param tile_data] object should change over time, use [method notify_runtime_tile_data_update] to notify the TileMap it needs an update.
*/
func (class) _tile_data_runtime_update(impl func(ptr unsafe.Pointer, layer int64, coords Vector2i.XY, tile_data [1]gdclass.TileData)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var layer = gd.UnsafeGet[int64](p_args, 0)

		var coords = gd.UnsafeGet[Vector2i.XY](p_args, 1)

		var tile_data = [1]gdclass.TileData{pointers.New[gdclass.TileData]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 2))})}

		defer pointers.End(tile_data[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, layer, coords, tile_data)
	}
}

/*
Assigns [param map] as a [NavigationServer2D] navigation map for the specified TileMap layer [param layer].
*/
//go:nosplit
func (self class) SetNavigationMap(layer int64, mapping RID.Any) { //gd:TileMap.set_navigation_map
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_set_navigation_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [RID] of the [NavigationServer2D] navigation map assigned to the specified TileMap layer [param layer].
*/
//go:nosplit
func (self class) GetNavigationMap(layer int64) RID.Any { //gd:TileMap.get_navigation_map
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_get_navigation_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Forces the TileMap and the layer [param layer] to update.
*/
//go:nosplit
func (self class) ForceUpdate(layer int64) { //gd:TileMap.force_update
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_force_update, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetTileset(tileset [1]gdclass.TileSet) { //gd:TileMap.set_tileset
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(tileset[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_set_tileset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTileset() [1]gdclass.TileSet { //gd:TileMap.get_tileset
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_get_tileset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TileSet{gd.PointerWithOwnershipTransferredToGo[gdclass.TileSet](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRenderingQuadrantSize(size int64) { //gd:TileMap.set_rendering_quadrant_size
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_set_rendering_quadrant_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRenderingQuadrantSize() int64 { //gd:TileMap.get_rendering_quadrant_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_get_rendering_quadrant_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of layers in the TileMap.
*/
//go:nosplit
func (self class) GetLayersCount() int64 { //gd:TileMap.get_layers_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_get_layers_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a layer at the given position [param to_position] in the array. If [param to_position] is negative, the position is counted from the end, with [code]-1[/code] adding the layer at the end of the array.
*/
//go:nosplit
func (self class) AddLayer(to_position int64) { //gd:TileMap.add_layer
	var frame = callframe.New()
	callframe.Arg(frame, to_position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_add_layer, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Moves the layer at index [param layer] to the given position [param to_position] in the array.
*/
//go:nosplit
func (self class) MoveLayer(layer int64, to_position int64) { //gd:TileMap.move_layer
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, to_position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_move_layer, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes the layer at index [param layer].
*/
//go:nosplit
func (self class) RemoveLayer(layer int64) { //gd:TileMap.remove_layer
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_remove_layer, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets a layer's name. This is mostly useful in the editor.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) SetLayerName(layer int64, name String.Readable) { //gd:TileMap.set_layer_name
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_set_layer_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a TileMap layer's name.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) GetLayerName(layer int64) String.Readable { //gd:TileMap.get_layer_name
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_get_layer_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Enables or disables the layer [param layer]. A disabled layer is not processed at all (no rendering, no physics, etc.).
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) SetLayerEnabled(layer int64, enabled bool) { //gd:TileMap.set_layer_enabled
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_set_layer_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns if a layer is enabled.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) IsLayerEnabled(layer int64) bool { //gd:TileMap.is_layer_enabled
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_is_layer_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets a layer's color. It will be multiplied by tile's color and TileMap's modulate.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) SetLayerModulate(layer int64, modulate Color.RGBA) { //gd:TileMap.set_layer_modulate
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, modulate)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_set_layer_modulate, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a TileMap layer's modulate.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) GetLayerModulate(layer int64) Color.RGBA { //gd:TileMap.get_layer_modulate
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_get_layer_modulate, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) SetLayerYSortEnabled(layer int64, y_sort_enabled bool) { //gd:TileMap.set_layer_y_sort_enabled
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, y_sort_enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_set_layer_y_sort_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns if a layer Y-sorts its tiles.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) IsLayerYSortEnabled(layer int64) bool { //gd:TileMap.is_layer_y_sort_enabled
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_is_layer_y_sort_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) SetLayerYSortOrigin(layer int64, y_sort_origin int64) { //gd:TileMap.set_layer_y_sort_origin
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, y_sort_origin)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_set_layer_y_sort_origin, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a TileMap layer's Y sort origin.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) GetLayerYSortOrigin(layer int64) int64 { //gd:TileMap.get_layer_y_sort_origin
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_get_layer_y_sort_origin, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets a layers Z-index value. This Z-index is added to each tile's Z-index value.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) SetLayerZIndex(layer int64, z_index int64) { //gd:TileMap.set_layer_z_index
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, z_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_set_layer_z_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a TileMap layer's Z-index value.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) GetLayerZIndex(layer int64) int64 { //gd:TileMap.get_layer_z_index
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_get_layer_z_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Enables or disables a layer's built-in navigation regions generation. Disable this if you need to bake navigation regions from a TileMap using a [NavigationRegion2D] node.
*/
//go:nosplit
func (self class) SetLayerNavigationEnabled(layer int64, enabled bool) { //gd:TileMap.set_layer_navigation_enabled
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_set_layer_navigation_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns if a layer's built-in navigation regions generation is enabled.
*/
//go:nosplit
func (self class) IsLayerNavigationEnabled(layer int64) bool { //gd:TileMap.is_layer_navigation_enabled
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_is_layer_navigation_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) SetLayerNavigationMap(layer int64, mapping RID.Any) { //gd:TileMap.set_layer_navigation_map
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, mapping)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_set_layer_navigation_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [RID] of the [NavigationServer2D] navigation map assigned to the specified TileMap layer [param layer].
By default the TileMap uses the default [World2D] navigation map for the first TileMap layer. For each additional TileMap layer a new navigation map is created for the additional layer.
In order to make [NavigationAgent2D] switch between TileMap layer navigation maps use [method NavigationAgent2D.set_navigation_map] with the navigation map received from [method get_layer_navigation_map].
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) GetLayerNavigationMap(layer int64) RID.Any { //gd:TileMap.get_layer_navigation_map
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Ret[RID.Any](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_get_layer_navigation_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollisionAnimatable(enabled bool) { //gd:TileMap.set_collision_animatable
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_set_collision_animatable, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsCollisionAnimatable() bool { //gd:TileMap.is_collision_animatable
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_is_collision_animatable, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCollisionVisibilityMode(collision_visibility_mode gdclass.TileMapVisibilityMode) { //gd:TileMap.set_collision_visibility_mode
	var frame = callframe.New()
	callframe.Arg(frame, collision_visibility_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_set_collision_visibility_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCollisionVisibilityMode() gdclass.TileMapVisibilityMode { //gd:TileMap.get_collision_visibility_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TileMapVisibilityMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_get_collision_visibility_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNavigationVisibilityMode(navigation_visibility_mode gdclass.TileMapVisibilityMode) { //gd:TileMap.set_navigation_visibility_mode
	var frame = callframe.New()
	callframe.Arg(frame, navigation_visibility_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_set_navigation_visibility_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetNavigationVisibilityMode() gdclass.TileMapVisibilityMode { //gd:TileMap.get_navigation_visibility_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TileMapVisibilityMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_get_navigation_visibility_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) SetCell(layer int64, coords Vector2i.XY, source_id int64, atlas_coords Vector2i.XY, alternative_tile int64) { //gd:TileMap.set_cell
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, coords)
	callframe.Arg(frame, source_id)
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, alternative_tile)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_set_cell, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Erases the cell on layer [param layer] at coordinates [param coords].
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) EraseCell(layer int64, coords Vector2i.XY) { //gd:TileMap.erase_cell
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, coords)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_erase_cell, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the tile source ID of the cell on layer [param layer] at coordinates [param coords]. Returns [code]-1[/code] if the cell does not exist.
If [param use_proxies] is [code]false[/code], ignores the [TileSet]'s tile proxies, returning the raw source identifier. See [method TileSet.map_tile_proxy].
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) GetCellSourceId(layer int64, coords Vector2i.XY, use_proxies bool) int64 { //gd:TileMap.get_cell_source_id
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, coords)
	callframe.Arg(frame, use_proxies)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_get_cell_source_id, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) GetCellAtlasCoords(layer int64, coords Vector2i.XY, use_proxies bool) Vector2i.XY { //gd:TileMap.get_cell_atlas_coords
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, coords)
	callframe.Arg(frame, use_proxies)
	var r_ret = callframe.Ret[Vector2i.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_get_cell_atlas_coords, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) GetCellAlternativeTile(layer int64, coords Vector2i.XY, use_proxies bool) int64 { //gd:TileMap.get_cell_alternative_tile
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, coords)
	callframe.Arg(frame, use_proxies)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_get_cell_alternative_tile, self.AsObject(), frame.Array(0), r_ret.Addr())
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
func (self class) GetCellTileData(layer int64, coords Vector2i.XY, use_proxies bool) [1]gdclass.TileData { //gd:TileMap.get_cell_tile_data
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, coords)
	callframe.Arg(frame, use_proxies)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_get_cell_tile_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TileData{gd.PointerMustAssertInstanceID[gdclass.TileData](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the coordinates of the tile for given physics body RID. Such RID can be retrieved from [method KinematicCollision2D.get_collider_rid], when colliding with a tile.
*/
//go:nosplit
func (self class) GetCoordsForBodyRid(body RID.Any) Vector2i.XY { //gd:TileMap.get_coords_for_body_rid
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[Vector2i.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_get_coords_for_body_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the tilemap layer of the tile for given physics body RID. Such RID can be retrieved from [method KinematicCollision2D.get_collider_rid], when colliding with a tile.
*/
//go:nosplit
func (self class) GetLayerForBodyRid(body RID.Any) int64 { //gd:TileMap.get_layer_for_body_rid
	var frame = callframe.New()
	callframe.Arg(frame, body)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_get_layer_for_body_rid, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a new [TileMapPattern] from the given layer and set of cells.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) GetPattern(layer int64, coords_array Array.Contains[Vector2i.XY]) [1]gdclass.TileMapPattern { //gd:TileMap.get_pattern
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, pointers.Get(gd.InternalArray(coords_array)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_get_pattern, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TileMapPattern{gd.PointerWithOwnershipTransferredToGo[gdclass.TileMapPattern](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns for the given coordinate [param coords_in_pattern] in a [TileMapPattern] the corresponding cell coordinates if the pattern was pasted at the [param position_in_tilemap] coordinates (see [method set_pattern]). This mapping is required as in half-offset tile shapes, the mapping might not work by calculating [code]position_in_tile_map + coords_in_pattern[/code].
*/
//go:nosplit
func (self class) MapPattern(position_in_tilemap Vector2i.XY, coords_in_pattern Vector2i.XY, pattern [1]gdclass.TileMapPattern) Vector2i.XY { //gd:TileMap.map_pattern
	var frame = callframe.New()
	callframe.Arg(frame, position_in_tilemap)
	callframe.Arg(frame, coords_in_pattern)
	callframe.Arg(frame, pointers.Get(pattern[0])[0])
	var r_ret = callframe.Ret[Vector2i.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_map_pattern, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Paste the given [TileMapPattern] at the given [param position] and [param layer] in the tile map.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) SetPattern(layer int64, position Vector2i.XY, pattern [1]gdclass.TileMapPattern) { //gd:TileMap.set_pattern
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, position)
	callframe.Arg(frame, pointers.Get(pattern[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_set_pattern, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Update all the cells in the [param cells] coordinates array so that they use the given [param terrain] for the given [param terrain_set]. If an updated cell has the same terrain as one of its neighboring cells, this function tries to join the two. This function might update neighboring tiles if needed to create correct terrain transitions.
If [param ignore_empty_terrains] is true, empty terrains will be ignored when trying to find the best fitting tile for the given terrain constraints.
If [param layer] is negative, the layers are accessed from the last one.
[b]Note:[/b] To work correctly, this method requires the TileMap's TileSet to have terrains set up with all required terrain combinations. Otherwise, it may produce unexpected results.
*/
//go:nosplit
func (self class) SetCellsTerrainConnect(layer int64, cells Array.Contains[Vector2i.XY], terrain_set int64, terrain int64, ignore_empty_terrains bool) { //gd:TileMap.set_cells_terrain_connect
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, pointers.Get(gd.InternalArray(cells)))
	callframe.Arg(frame, terrain_set)
	callframe.Arg(frame, terrain)
	callframe.Arg(frame, ignore_empty_terrains)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_set_cells_terrain_connect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Update all the cells in the [param path] coordinates array so that they use the given [param terrain] for the given [param terrain_set]. The function will also connect two successive cell in the path with the same terrain. This function might update neighboring tiles if needed to create correct terrain transitions.
If [param ignore_empty_terrains] is true, empty terrains will be ignored when trying to find the best fitting tile for the given terrain constraints.
If [param layer] is negative, the layers are accessed from the last one.
[b]Note:[/b] To work correctly, this method requires the TileMap's TileSet to have terrains set up with all required terrain combinations. Otherwise, it may produce unexpected results.
*/
//go:nosplit
func (self class) SetCellsTerrainPath(layer int64, path Array.Contains[Vector2i.XY], terrain_set int64, terrain int64, ignore_empty_terrains bool) { //gd:TileMap.set_cells_terrain_path
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, pointers.Get(gd.InternalArray(path)))
	callframe.Arg(frame, terrain_set)
	callframe.Arg(frame, terrain)
	callframe.Arg(frame, ignore_empty_terrains)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_set_cells_terrain_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Clears cells that do not exist in the tileset.
*/
//go:nosplit
func (self class) FixInvalidTiles() { //gd:TileMap.fix_invalid_tiles
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_fix_invalid_tiles, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Clears all cells on the given layer.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) ClearLayer(layer int64) { //gd:TileMap.clear_layer
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_clear_layer, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Clears all cells.
*/
//go:nosplit
func (self class) Clear() { //gd:TileMap.clear
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Triggers a direct update of the TileMap. Usually, calling this function is not needed, as TileMap node updates automatically when one of its properties or cells is modified.
However, for performance reasons, those updates are batched and delayed to the end of the frame. Calling this function will force the TileMap to update right away instead.
[b]Warning:[/b] Updating the TileMap is computationally expensive and may impact performance. Try to limit the number of updates and how many tiles they impact.
*/
//go:nosplit
func (self class) UpdateInternals() { //gd:TileMap.update_internals
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_update_internals, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Notifies the TileMap node that calls to [method _use_tile_data_runtime_update] or [method _tile_data_runtime_update] will lead to different results. This will thus trigger a TileMap update.
If [param layer] is provided, only notifies changes for the given layer. Providing the [param layer] argument (when applicable) is usually preferred for performance reasons.
[b]Warning:[/b] Updating the TileMap is computationally expensive and may impact performance. Try to limit the number of calls to this function to avoid unnecessary update.
[b]Note:[/b] This does not trigger a direct update of the TileMap, the update will be done at the end of the frame as usual (unless you call [method update_internals]).
*/
//go:nosplit
func (self class) NotifyRuntimeTileDataUpdate(layer int64) { //gd:TileMap.notify_runtime_tile_data_update
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_notify_runtime_tile_data_update, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the list of all neighbourings cells to the one at [param coords].
*/
//go:nosplit
func (self class) GetSurroundingCells(coords Vector2i.XY) Array.Contains[Vector2i.XY] { //gd:TileMap.get_surrounding_cells
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_get_surrounding_cells, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Vector2i.XY]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns a [Vector2i] array with the positions of all cells containing a tile in the given layer. A cell is considered empty if its source identifier equals -1, its atlas coordinates identifiers is [code]Vector2(-1, -1)[/code] and its alternative identifier is -1.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) GetUsedCells(layer int64) Array.Contains[Vector2i.XY] { //gd:TileMap.get_used_cells
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_get_used_cells, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Vector2i.XY]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns a [Vector2i] array with the positions of all cells containing a tile in the given layer. Tiles may be filtered according to their source ([param source_id]), their atlas coordinates ([param atlas_coords]) or alternative id ([param alternative_tile]).
If a parameter has its value set to the default one, this parameter is not used to filter a cell. Thus, if all parameters have their respective default value, this method returns the same result as [method get_used_cells].
A cell is considered empty if its source identifier equals -1, its atlas coordinates identifiers is [code]Vector2(-1, -1)[/code] and its alternative identifier is -1.
If [param layer] is negative, the layers are accessed from the last one.
*/
//go:nosplit
func (self class) GetUsedCellsById(layer int64, source_id int64, atlas_coords Vector2i.XY, alternative_tile int64) Array.Contains[Vector2i.XY] { //gd:TileMap.get_used_cells_by_id
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, source_id)
	callframe.Arg(frame, atlas_coords)
	callframe.Arg(frame, alternative_tile)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_get_used_cells_by_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Vector2i.XY]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns a rectangle enclosing the used (non-empty) tiles of the map, including all layers.
*/
//go:nosplit
func (self class) GetUsedRect() Rect2i.PositionSize { //gd:TileMap.get_used_rect
	var frame = callframe.New()
	var r_ret = callframe.Ret[Rect2i.PositionSize](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_get_used_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the centered position of a cell in the TileMap's local coordinate space. To convert the returned value into global coordinates, use [method Node2D.to_global]. See also [method local_to_map].
[b]Note:[/b] This may not correspond to the visual position of the tile, i.e. it ignores the [member TileData.texture_origin] property of individual tiles.
*/
//go:nosplit
func (self class) MapToLocal(map_position Vector2i.XY) Vector2.XY { //gd:TileMap.map_to_local
	var frame = callframe.New()
	callframe.Arg(frame, map_position)
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_map_to_local, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the map coordinates of the cell containing the given [param local_position]. If [param local_position] is in global coordinates, consider using [method Node2D.to_local] before passing it to this method. See also [method map_to_local].
*/
//go:nosplit
func (self class) LocalToMap(local_position Vector2.XY) Vector2i.XY { //gd:TileMap.local_to_map
	var frame = callframe.New()
	callframe.Arg(frame, local_position)
	var r_ret = callframe.Ret[Vector2i.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_local_to_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the neighboring cell to the one at coordinates [param coords], identified by the [param neighbor] direction. This method takes into account the different layouts a TileMap can take.
*/
//go:nosplit
func (self class) GetNeighborCell(coords Vector2i.XY, neighbor gdclass.TileSetCellNeighbor) Vector2i.XY { //gd:TileMap.get_neighbor_cell
	var frame = callframe.New()
	callframe.Arg(frame, coords)
	callframe.Arg(frame, neighbor)
	var r_ret = callframe.Ret[Vector2i.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileMap.Bind_get_neighbor_cell, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("changed"), gd.NewCallable(cb), 0)
}

func (self class) AsTileMap() Advanced          { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTileMap() Instance       { return *((*Instance)(unsafe.Pointer(&self))) }
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
	case "_use_tile_data_runtime_update":
		return reflect.ValueOf(self._use_tile_data_runtime_update)
	case "_tile_data_runtime_update":
		return reflect.ValueOf(self._tile_data_runtime_update)
	default:
		return gd.VirtualByName(Node2D.Advanced(self.AsNode2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_use_tile_data_runtime_update":
		return reflect.ValueOf(self._use_tile_data_runtime_update)
	case "_tile_data_runtime_update":
		return reflect.ValueOf(self._tile_data_runtime_update)
	default:
		return gd.VirtualByName(Node2D.Instance(self.AsNode2D()), name)
	}
}
func init() {
	gdclass.Register("TileMap", func(ptr gd.Object) any { return [1]gdclass.TileMap{*(*gdclass.TileMap)(unsafe.Pointer(&ptr))} })
}

type VisibilityMode = gdclass.TileMapVisibilityMode //gd:TileMap.VisibilityMode

const (
	/*Use the debug settings to determine visibility.*/
	VisibilityModeDefault VisibilityMode = 0
	/*Always hide.*/
	VisibilityModeForceHide VisibilityMode = 2
	/*Always show.*/
	VisibilityModeForceShow VisibilityMode = 1
)
