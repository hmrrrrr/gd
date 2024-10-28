package TileSet

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
A TileSet is a library of tiles for a [TileMap]. A TileSet handles a list of [TileSetSource], each of them storing a set of tiles.
Tiles can either be from a [TileSetAtlasSource], which renders tiles out of a texture with support for physics, navigation, etc., or from a [TileSetScenesCollectionSource], which exposes scene-based tiles.
Tiles are referenced by using three IDs: their source ID, their atlas coordinates ID, and their alternative tile ID.
A TileSet can be configured so that its tiles expose more or fewer properties. To do so, the TileSet resources use property layers, which you can add or remove depending on your needs.
For example, adding a physics layer allows giving collision shapes to your tiles. Each layer has dedicated properties (physics layer and mask), so you may add several TileSet physics layers for each type of collision you need.
See the functions to add new layers for more information.

*/
type Go [1]classdb.TileSet

/*
Returns a new unused source ID. This generated ID is the same that a call to [method add_source] would return.
*/
func (self Go) GetNextSourceId() int {
	return int(int(class(self).GetNextSourceId()))
}

/*
Adds a [TileSetSource] to the TileSet. If [param atlas_source_id_override] is not -1, also set its source ID. Otherwise, a unique identifier is automatically generated.
The function returns the added source ID or -1 if the source could not be added.
[b]Warning:[/b] A source cannot belong to two TileSets at the same time. If the added source was attached to another [TileSet], it will be removed from that one.
*/
func (self Go) AddSource(source gdclass.TileSetSource) int {
	return int(int(class(self).AddSource(source, gd.Int(-1))))
}

/*
Removes the source with the given source ID.
*/
func (self Go) RemoveSource(source_id int) {
	class(self).RemoveSource(gd.Int(source_id))
}

/*
Changes a source's ID.
*/
func (self Go) SetSourceId(source_id int, new_source_id int) {
	class(self).SetSourceId(gd.Int(source_id), gd.Int(new_source_id))
}

/*
Returns the number of [TileSetSource] in this TileSet.
*/
func (self Go) GetSourceCount() int {
	return int(int(class(self).GetSourceCount()))
}

/*
Returns the source ID for source with index [param index].
*/
func (self Go) GetSourceId(index int) int {
	return int(int(class(self).GetSourceId(gd.Int(index))))
}

/*
Returns if this TileSet has a source for the given source ID.
*/
func (self Go) HasSource(source_id int) bool {
	return bool(class(self).HasSource(gd.Int(source_id)))
}

/*
Returns the [TileSetSource] with ID [param source_id].
*/
func (self Go) GetSource(source_id int) gdclass.TileSetSource {
	return gdclass.TileSetSource(class(self).GetSource(gd.Int(source_id)))
}

/*
Returns the occlusion layers count.
*/
func (self Go) GetOcclusionLayersCount() int {
	return int(int(class(self).GetOcclusionLayersCount()))
}

/*
Adds an occlusion layer to the TileSet at the given position [param to_position] in the array. If [param to_position] is -1, adds it at the end of the array.
Occlusion layers allow assigning occlusion polygons to atlas tiles.
*/
func (self Go) AddOcclusionLayer() {
	class(self).AddOcclusionLayer(gd.Int(-1))
}

/*
Moves the occlusion layer at index [param layer_index] to the given position [param to_position] in the array. Also updates the atlas tiles accordingly.
*/
func (self Go) MoveOcclusionLayer(layer_index int, to_position int) {
	class(self).MoveOcclusionLayer(gd.Int(layer_index), gd.Int(to_position))
}

/*
Removes the occlusion layer at index [param layer_index]. Also updates the atlas tiles accordingly.
*/
func (self Go) RemoveOcclusionLayer(layer_index int) {
	class(self).RemoveOcclusionLayer(gd.Int(layer_index))
}

/*
Sets the occlusion layer (as in the rendering server) for occluders in the given TileSet occlusion layer.
*/
func (self Go) SetOcclusionLayerLightMask(layer_index int, light_mask int) {
	class(self).SetOcclusionLayerLightMask(gd.Int(layer_index), gd.Int(light_mask))
}

/*
Returns the light mask of the occlusion layer.
*/
func (self Go) GetOcclusionLayerLightMask(layer_index int) int {
	return int(int(class(self).GetOcclusionLayerLightMask(gd.Int(layer_index))))
}

/*
Enables or disables SDF collision for occluders in the given TileSet occlusion layer.
*/
func (self Go) SetOcclusionLayerSdfCollision(layer_index int, sdf_collision bool) {
	class(self).SetOcclusionLayerSdfCollision(gd.Int(layer_index), sdf_collision)
}

/*
Returns if the occluders from this layer use [code]sdf_collision[/code].
*/
func (self Go) GetOcclusionLayerSdfCollision(layer_index int) bool {
	return bool(class(self).GetOcclusionLayerSdfCollision(gd.Int(layer_index)))
}

/*
Returns the physics layers count.
*/
func (self Go) GetPhysicsLayersCount() int {
	return int(int(class(self).GetPhysicsLayersCount()))
}

/*
Adds a physics layer to the TileSet at the given position [param to_position] in the array. If [param to_position] is -1, adds it at the end of the array.
Physics layers allow assigning collision polygons to atlas tiles.
*/
func (self Go) AddPhysicsLayer() {
	class(self).AddPhysicsLayer(gd.Int(-1))
}

/*
Moves the physics layer at index [param layer_index] to the given position [param to_position] in the array. Also updates the atlas tiles accordingly.
*/
func (self Go) MovePhysicsLayer(layer_index int, to_position int) {
	class(self).MovePhysicsLayer(gd.Int(layer_index), gd.Int(to_position))
}

/*
Removes the physics layer at index [param layer_index]. Also updates the atlas tiles accordingly.
*/
func (self Go) RemovePhysicsLayer(layer_index int) {
	class(self).RemovePhysicsLayer(gd.Int(layer_index))
}

/*
Sets the physics layer (as in the physics server) for bodies in the given TileSet physics layer.
*/
func (self Go) SetPhysicsLayerCollisionLayer(layer_index int, layer int) {
	class(self).SetPhysicsLayerCollisionLayer(gd.Int(layer_index), gd.Int(layer))
}

/*
Returns the collision layer (as in the physics server) bodies on the given TileSet's physics layer are in.
*/
func (self Go) GetPhysicsLayerCollisionLayer(layer_index int) int {
	return int(int(class(self).GetPhysicsLayerCollisionLayer(gd.Int(layer_index))))
}

/*
Sets the physics layer (as in the physics server) for bodies in the given TileSet physics layer.
*/
func (self Go) SetPhysicsLayerCollisionMask(layer_index int, mask int) {
	class(self).SetPhysicsLayerCollisionMask(gd.Int(layer_index), gd.Int(mask))
}

/*
Returns the collision mask of bodies on the given TileSet's physics layer.
*/
func (self Go) GetPhysicsLayerCollisionMask(layer_index int) int {
	return int(int(class(self).GetPhysicsLayerCollisionMask(gd.Int(layer_index))))
}

/*
Sets the physics material for bodies in the given TileSet physics layer.
*/
func (self Go) SetPhysicsLayerPhysicsMaterial(layer_index int, physics_material gdclass.PhysicsMaterial) {
	class(self).SetPhysicsLayerPhysicsMaterial(gd.Int(layer_index), physics_material)
}

/*
Returns the physics material of bodies on the given TileSet's physics layer.
*/
func (self Go) GetPhysicsLayerPhysicsMaterial(layer_index int) gdclass.PhysicsMaterial {
	return gdclass.PhysicsMaterial(class(self).GetPhysicsLayerPhysicsMaterial(gd.Int(layer_index)))
}

/*
Returns the terrain sets count.
*/
func (self Go) GetTerrainSetsCount() int {
	return int(int(class(self).GetTerrainSetsCount()))
}

/*
Adds a new terrain set at the given position [param to_position] in the array. If [param to_position] is -1, adds it at the end of the array.
*/
func (self Go) AddTerrainSet() {
	class(self).AddTerrainSet(gd.Int(-1))
}

/*
Moves the terrain set at index [param terrain_set] to the given position [param to_position] in the array. Also updates the atlas tiles accordingly.
*/
func (self Go) MoveTerrainSet(terrain_set int, to_position int) {
	class(self).MoveTerrainSet(gd.Int(terrain_set), gd.Int(to_position))
}

/*
Removes the terrain set at index [param terrain_set]. Also updates the atlas tiles accordingly.
*/
func (self Go) RemoveTerrainSet(terrain_set int) {
	class(self).RemoveTerrainSet(gd.Int(terrain_set))
}

/*
Sets a terrain mode. Each mode determines which bits of a tile shape is used to match the neighboring tiles' terrains.
*/
func (self Go) SetTerrainSetMode(terrain_set int, mode classdb.TileSetTerrainMode) {
	class(self).SetTerrainSetMode(gd.Int(terrain_set), mode)
}

/*
Returns a terrain set mode.
*/
func (self Go) GetTerrainSetMode(terrain_set int) classdb.TileSetTerrainMode {
	return classdb.TileSetTerrainMode(class(self).GetTerrainSetMode(gd.Int(terrain_set)))
}

/*
Returns the number of terrains in the given terrain set.
*/
func (self Go) GetTerrainsCount(terrain_set int) int {
	return int(int(class(self).GetTerrainsCount(gd.Int(terrain_set))))
}

/*
Adds a new terrain to the given terrain set [param terrain_set] at the given position [param to_position] in the array. If [param to_position] is -1, adds it at the end of the array.
*/
func (self Go) AddTerrain(terrain_set int) {
	class(self).AddTerrain(gd.Int(terrain_set), gd.Int(-1))
}

/*
Moves the terrain at index [param terrain_index] for terrain set [param terrain_set] to the given position [param to_position] in the array. Also updates the atlas tiles accordingly.
*/
func (self Go) MoveTerrain(terrain_set int, terrain_index int, to_position int) {
	class(self).MoveTerrain(gd.Int(terrain_set), gd.Int(terrain_index), gd.Int(to_position))
}

/*
Removes the terrain at index [param terrain_index] in the given terrain set [param terrain_set]. Also updates the atlas tiles accordingly.
*/
func (self Go) RemoveTerrain(terrain_set int, terrain_index int) {
	class(self).RemoveTerrain(gd.Int(terrain_set), gd.Int(terrain_index))
}

/*
Sets a terrain's name.
*/
func (self Go) SetTerrainName(terrain_set int, terrain_index int, name string) {
	class(self).SetTerrainName(gd.Int(terrain_set), gd.Int(terrain_index), gd.NewString(name))
}

/*
Returns a terrain's name.
*/
func (self Go) GetTerrainName(terrain_set int, terrain_index int) string {
	return string(class(self).GetTerrainName(gd.Int(terrain_set), gd.Int(terrain_index)).String())
}

/*
Sets a terrain's color. This color is used for identifying the different terrains in the TileSet editor.
*/
func (self Go) SetTerrainColor(terrain_set int, terrain_index int, color gd.Color) {
	class(self).SetTerrainColor(gd.Int(terrain_set), gd.Int(terrain_index), color)
}

/*
Returns a terrain's color.
*/
func (self Go) GetTerrainColor(terrain_set int, terrain_index int) gd.Color {
	return gd.Color(class(self).GetTerrainColor(gd.Int(terrain_set), gd.Int(terrain_index)))
}

/*
Returns the navigation layers count.
*/
func (self Go) GetNavigationLayersCount() int {
	return int(int(class(self).GetNavigationLayersCount()))
}

/*
Adds a navigation layer to the TileSet at the given position [param to_position] in the array. If [param to_position] is -1, adds it at the end of the array.
Navigation layers allow assigning a navigable area to atlas tiles.
*/
func (self Go) AddNavigationLayer() {
	class(self).AddNavigationLayer(gd.Int(-1))
}

/*
Moves the navigation layer at index [param layer_index] to the given position [param to_position] in the array. Also updates the atlas tiles accordingly.
*/
func (self Go) MoveNavigationLayer(layer_index int, to_position int) {
	class(self).MoveNavigationLayer(gd.Int(layer_index), gd.Int(to_position))
}

/*
Removes the navigation layer at index [param layer_index]. Also updates the atlas tiles accordingly.
*/
func (self Go) RemoveNavigationLayer(layer_index int) {
	class(self).RemoveNavigationLayer(gd.Int(layer_index))
}

/*
Sets the navigation layers (as in the navigation server) for navigation regions in the given TileSet navigation layer.
*/
func (self Go) SetNavigationLayerLayers(layer_index int, layers int) {
	class(self).SetNavigationLayerLayers(gd.Int(layer_index), gd.Int(layers))
}

/*
Returns the navigation layers (as in the Navigation server) of the given TileSet navigation layer.
*/
func (self Go) GetNavigationLayerLayers(layer_index int) int {
	return int(int(class(self).GetNavigationLayerLayers(gd.Int(layer_index))))
}

/*
Based on [param value], enables or disables the specified navigation layer of the TileSet navigation data layer identified by the given [param layer_index], given a navigation_layers [param layer_number] between 1 and 32.
*/
func (self Go) SetNavigationLayerLayerValue(layer_index int, layer_number int, value bool) {
	class(self).SetNavigationLayerLayerValue(gd.Int(layer_index), gd.Int(layer_number), value)
}

/*
Returns whether or not the specified navigation layer of the TileSet navigation data layer identified by the given [param layer_index] is enabled, given a navigation_layers [param layer_number] between 1 and 32.
*/
func (self Go) GetNavigationLayerLayerValue(layer_index int, layer_number int) bool {
	return bool(class(self).GetNavigationLayerLayerValue(gd.Int(layer_index), gd.Int(layer_number)))
}

/*
Returns the custom data layers count.
*/
func (self Go) GetCustomDataLayersCount() int {
	return int(int(class(self).GetCustomDataLayersCount()))
}

/*
Adds a custom data layer to the TileSet at the given position [param to_position] in the array. If [param to_position] is -1, adds it at the end of the array.
Custom data layers allow assigning custom properties to atlas tiles.
*/
func (self Go) AddCustomDataLayer() {
	class(self).AddCustomDataLayer(gd.Int(-1))
}

/*
Moves the custom data layer at index [param layer_index] to the given position [param to_position] in the array. Also updates the atlas tiles accordingly.
*/
func (self Go) MoveCustomDataLayer(layer_index int, to_position int) {
	class(self).MoveCustomDataLayer(gd.Int(layer_index), gd.Int(to_position))
}

/*
Removes the custom data layer at index [param layer_index]. Also updates the atlas tiles accordingly.
*/
func (self Go) RemoveCustomDataLayer(layer_index int) {
	class(self).RemoveCustomDataLayer(gd.Int(layer_index))
}

/*
Returns the index of the custom data layer identified by the given name.
*/
func (self Go) GetCustomDataLayerByName(layer_name string) int {
	return int(int(class(self).GetCustomDataLayerByName(gd.NewString(layer_name))))
}

/*
Sets the name of the custom data layer identified by the given index. Names are identifiers of the layer therefore if the name is already taken it will fail and raise an error.
*/
func (self Go) SetCustomDataLayerName(layer_index int, layer_name string) {
	class(self).SetCustomDataLayerName(gd.Int(layer_index), gd.NewString(layer_name))
}

/*
Returns the name of the custom data layer identified by the given index.
*/
func (self Go) GetCustomDataLayerName(layer_index int) string {
	return string(class(self).GetCustomDataLayerName(gd.Int(layer_index)).String())
}

/*
Sets the type of the custom data layer identified by the given index.
*/
func (self Go) SetCustomDataLayerType(layer_index int, layer_type gd.VariantType) {
	class(self).SetCustomDataLayerType(gd.Int(layer_index), layer_type)
}

/*
Returns the type of the custom data layer identified by the given index.
*/
func (self Go) GetCustomDataLayerType(layer_index int) gd.VariantType {
	return gd.VariantType(class(self).GetCustomDataLayerType(gd.Int(layer_index)))
}

/*
Creates a source-level proxy for the given source ID. A proxy will map set of tile identifiers to another set of identifiers. Both the atlas coordinates ID and the alternative tile ID are kept the same when using source-level proxies.
This can be used to replace a source in all TileMaps using this TileSet, as TileMap nodes will find and use the proxy's target source when one is available.
Proxied tiles can be automatically replaced in TileMap nodes using the editor.
*/
func (self Go) SetSourceLevelTileProxy(source_from int, source_to int) {
	class(self).SetSourceLevelTileProxy(gd.Int(source_from), gd.Int(source_to))
}

/*
Returns the source-level proxy for the given source identifier.
If the TileSet has no proxy for the given identifier, returns -1.
*/
func (self Go) GetSourceLevelTileProxy(source_from int) int {
	return int(int(class(self).GetSourceLevelTileProxy(gd.Int(source_from))))
}

/*
Returns if there is a source-level proxy for the given source ID.
*/
func (self Go) HasSourceLevelTileProxy(source_from int) bool {
	return bool(class(self).HasSourceLevelTileProxy(gd.Int(source_from)))
}

/*
Removes a source-level tile proxy.
*/
func (self Go) RemoveSourceLevelTileProxy(source_from int) {
	class(self).RemoveSourceLevelTileProxy(gd.Int(source_from))
}

/*
Creates a coordinates-level proxy for the given identifiers. A proxy will map set of tile identifiers to another set of identifiers. The alternative tile ID is kept the same when using coordinates-level proxies.
This can be used to replace a tile in all TileMaps using this TileSet, as TileMap nodes will find and use the proxy's target tile when one is available.
Proxied tiles can be automatically replaced in TileMap nodes using the editor.
*/
func (self Go) SetCoordsLevelTileProxy(p_source_from int, coords_from gd.Vector2i, source_to int, coords_to gd.Vector2i) {
	class(self).SetCoordsLevelTileProxy(gd.Int(p_source_from), coords_from, gd.Int(source_to), coords_to)
}

/*
Returns the coordinate-level proxy for the given identifiers. The returned array contains the two target identifiers of the proxy (source ID and atlas coordinates ID).
If the TileSet has no proxy for the given identifiers, returns an empty Array.
*/
func (self Go) GetCoordsLevelTileProxy(source_from int, coords_from gd.Vector2i) gd.Array {
	return gd.Array(class(self).GetCoordsLevelTileProxy(gd.Int(source_from), coords_from))
}

/*
Returns if there is a coodinates-level proxy for the given identifiers.
*/
func (self Go) HasCoordsLevelTileProxy(source_from int, coords_from gd.Vector2i) bool {
	return bool(class(self).HasCoordsLevelTileProxy(gd.Int(source_from), coords_from))
}

/*
Removes a coordinates-level proxy for the given identifiers.
*/
func (self Go) RemoveCoordsLevelTileProxy(source_from int, coords_from gd.Vector2i) {
	class(self).RemoveCoordsLevelTileProxy(gd.Int(source_from), coords_from)
}

/*
Create an alternative-level proxy for the given identifiers. A proxy will map set of tile identifiers to another set of identifiers.
This can be used to replace a tile in all TileMaps using this TileSet, as TileMap nodes will find and use the proxy's target tile when one is available.
Proxied tiles can be automatically replaced in TileMap nodes using the editor.
*/
func (self Go) SetAlternativeLevelTileProxy(source_from int, coords_from gd.Vector2i, alternative_from int, source_to int, coords_to gd.Vector2i, alternative_to int) {
	class(self).SetAlternativeLevelTileProxy(gd.Int(source_from), coords_from, gd.Int(alternative_from), gd.Int(source_to), coords_to, gd.Int(alternative_to))
}

/*
Returns the alternative-level proxy for the given identifiers. The returned array contains the three proxie's target identifiers (source ID, atlas coords ID and alternative tile ID).
If the TileSet has no proxy for the given identifiers, returns an empty Array.
*/
func (self Go) GetAlternativeLevelTileProxy(source_from int, coords_from gd.Vector2i, alternative_from int) gd.Array {
	return gd.Array(class(self).GetAlternativeLevelTileProxy(gd.Int(source_from), coords_from, gd.Int(alternative_from)))
}

/*
Returns if there is an alternative-level proxy for the given identifiers.
*/
func (self Go) HasAlternativeLevelTileProxy(source_from int, coords_from gd.Vector2i, alternative_from int) bool {
	return bool(class(self).HasAlternativeLevelTileProxy(gd.Int(source_from), coords_from, gd.Int(alternative_from)))
}

/*
Removes an alternative-level proxy for the given identifiers.
*/
func (self Go) RemoveAlternativeLevelTileProxy(source_from int, coords_from gd.Vector2i, alternative_from int) {
	class(self).RemoveAlternativeLevelTileProxy(gd.Int(source_from), coords_from, gd.Int(alternative_from))
}

/*
According to the configured proxies, maps the provided identifiers to a new set of identifiers. The source ID, atlas coordinates ID and alternative tile ID are returned as a 3 elements Array.
This function first look for matching alternative-level proxies, then coordinates-level proxies, then source-level proxies.
If no proxy corresponding to provided identifiers are found, returns the same values the ones used as arguments.
*/
func (self Go) MapTileProxy(source_from int, coords_from gd.Vector2i, alternative_from int) gd.Array {
	return gd.Array(class(self).MapTileProxy(gd.Int(source_from), coords_from, gd.Int(alternative_from)))
}

/*
Clears tile proxies pointing to invalid tiles.
*/
func (self Go) CleanupInvalidTileProxies() {
	class(self).CleanupInvalidTileProxies()
}

/*
Clears all tile proxies.
*/
func (self Go) ClearTileProxies() {
	class(self).ClearTileProxies()
}

/*
Adds a [TileMapPattern] to be stored in the TileSet resource. If provided, insert it at the given [param index].
*/
func (self Go) AddPattern(pattern gdclass.TileMapPattern) int {
	return int(int(class(self).AddPattern(pattern, gd.Int(-1))))
}

/*
Returns the [TileMapPattern] at the given [param index].
*/
func (self Go) GetPattern() gdclass.TileMapPattern {
	return gdclass.TileMapPattern(class(self).GetPattern(gd.Int(-1)))
}

/*
Remove the [TileMapPattern] at the given index.
*/
func (self Go) RemovePattern(index int) {
	class(self).RemovePattern(gd.Int(index))
}

/*
Returns the number of [TileMapPattern] this tile set handles.
*/
func (self Go) GetPatternsCount() int {
	return int(int(class(self).GetPatternsCount()))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.TileSet
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TileSet"))
	return Go{classdb.TileSet(object)}
}

func (self Go) TileShape() classdb.TileSetTileShape {
		return classdb.TileSetTileShape(class(self).GetTileShape())
}

func (self Go) SetTileShape(value classdb.TileSetTileShape) {
	class(self).SetTileShape(value)
}

func (self Go) TileLayout() classdb.TileSetTileLayout {
		return classdb.TileSetTileLayout(class(self).GetTileLayout())
}

func (self Go) SetTileLayout(value classdb.TileSetTileLayout) {
	class(self).SetTileLayout(value)
}

func (self Go) TileOffsetAxis() classdb.TileSetTileOffsetAxis {
		return classdb.TileSetTileOffsetAxis(class(self).GetTileOffsetAxis())
}

func (self Go) SetTileOffsetAxis(value classdb.TileSetTileOffsetAxis) {
	class(self).SetTileOffsetAxis(value)
}

func (self Go) TileSize() gd.Vector2i {
		return gd.Vector2i(class(self).GetTileSize())
}

func (self Go) SetTileSize(value gd.Vector2i) {
	class(self).SetTileSize(value)
}

func (self Go) UvClipping() bool {
		return bool(class(self).IsUvClipping())
}

func (self Go) SetUvClipping(value bool) {
	class(self).SetUvClipping(value)
}

/*
Returns a new unused source ID. This generated ID is the same that a call to [method add_source] would return.
*/
//go:nosplit
func (self class) GetNextSourceId() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_next_source_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a [TileSetSource] to the TileSet. If [param atlas_source_id_override] is not -1, also set its source ID. Otherwise, a unique identifier is automatically generated.
The function returns the added source ID or -1 if the source could not be added.
[b]Warning:[/b] A source cannot belong to two TileSets at the same time. If the added source was attached to another [TileSet], it will be removed from that one.
*/
//go:nosplit
func (self class) AddSource(source gdclass.TileSetSource, atlas_source_id_override gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(source[0])[0])
	callframe.Arg(frame, atlas_source_id_override)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_add_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes the source with the given source ID.
*/
//go:nosplit
func (self class) RemoveSource(source_id gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, source_id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_remove_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Changes a source's ID.
*/
//go:nosplit
func (self class) SetSourceId(source_id gd.Int, new_source_id gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, source_id)
	callframe.Arg(frame, new_source_id)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_set_source_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of [TileSetSource] in this TileSet.
*/
//go:nosplit
func (self class) GetSourceCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_source_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the source ID for source with index [param index].
*/
//go:nosplit
func (self class) GetSourceId(index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_source_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns if this TileSet has a source for the given source ID.
*/
//go:nosplit
func (self class) HasSource(source_id gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, source_id)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_has_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [TileSetSource] with ID [param source_id].
*/
//go:nosplit
func (self class) GetSource(source_id gd.Int) gdclass.TileSetSource {
	var frame = callframe.New()
	callframe.Arg(frame, source_id)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.TileSetSource{classdb.TileSetSource(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTileShape(shape classdb.TileSetTileShape)  {
	var frame = callframe.New()
	callframe.Arg(frame, shape)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_set_tile_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTileShape() classdb.TileSetTileShape {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TileSetTileShape](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_tile_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTileLayout(layout classdb.TileSetTileLayout)  {
	var frame = callframe.New()
	callframe.Arg(frame, layout)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_set_tile_layout, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTileLayout() classdb.TileSetTileLayout {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TileSetTileLayout](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_tile_layout, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTileOffsetAxis(alignment classdb.TileSetTileOffsetAxis)  {
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_set_tile_offset_axis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTileOffsetAxis() classdb.TileSetTileOffsetAxis {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TileSetTileOffsetAxis](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_tile_offset_axis, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTileSize(size gd.Vector2i)  {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_set_tile_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTileSize() gd.Vector2i {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_tile_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetUvClipping(uv_clipping bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, uv_clipping)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_set_uv_clipping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUvClipping() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_is_uv_clipping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the occlusion layers count.
*/
//go:nosplit
func (self class) GetOcclusionLayersCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_occlusion_layers_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds an occlusion layer to the TileSet at the given position [param to_position] in the array. If [param to_position] is -1, adds it at the end of the array.
Occlusion layers allow assigning occlusion polygons to atlas tiles.
*/
//go:nosplit
func (self class) AddOcclusionLayer(to_position gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, to_position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_add_occlusion_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Moves the occlusion layer at index [param layer_index] to the given position [param to_position] in the array. Also updates the atlas tiles accordingly.
*/
//go:nosplit
func (self class) MoveOcclusionLayer(layer_index gd.Int, to_position gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, layer_index)
	callframe.Arg(frame, to_position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_move_occlusion_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the occlusion layer at index [param layer_index]. Also updates the atlas tiles accordingly.
*/
//go:nosplit
func (self class) RemoveOcclusionLayer(layer_index gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, layer_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_remove_occlusion_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the occlusion layer (as in the rendering server) for occluders in the given TileSet occlusion layer.
*/
//go:nosplit
func (self class) SetOcclusionLayerLightMask(layer_index gd.Int, light_mask gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, layer_index)
	callframe.Arg(frame, light_mask)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_set_occlusion_layer_light_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the light mask of the occlusion layer.
*/
//go:nosplit
func (self class) GetOcclusionLayerLightMask(layer_index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, layer_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_occlusion_layer_light_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Enables or disables SDF collision for occluders in the given TileSet occlusion layer.
*/
//go:nosplit
func (self class) SetOcclusionLayerSdfCollision(layer_index gd.Int, sdf_collision bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, layer_index)
	callframe.Arg(frame, sdf_collision)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_set_occlusion_layer_sdf_collision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns if the occluders from this layer use [code]sdf_collision[/code].
*/
//go:nosplit
func (self class) GetOcclusionLayerSdfCollision(layer_index gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, layer_index)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_occlusion_layer_sdf_collision, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the physics layers count.
*/
//go:nosplit
func (self class) GetPhysicsLayersCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_physics_layers_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a physics layer to the TileSet at the given position [param to_position] in the array. If [param to_position] is -1, adds it at the end of the array.
Physics layers allow assigning collision polygons to atlas tiles.
*/
//go:nosplit
func (self class) AddPhysicsLayer(to_position gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, to_position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_add_physics_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Moves the physics layer at index [param layer_index] to the given position [param to_position] in the array. Also updates the atlas tiles accordingly.
*/
//go:nosplit
func (self class) MovePhysicsLayer(layer_index gd.Int, to_position gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, layer_index)
	callframe.Arg(frame, to_position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_move_physics_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the physics layer at index [param layer_index]. Also updates the atlas tiles accordingly.
*/
//go:nosplit
func (self class) RemovePhysicsLayer(layer_index gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, layer_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_remove_physics_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the physics layer (as in the physics server) for bodies in the given TileSet physics layer.
*/
//go:nosplit
func (self class) SetPhysicsLayerCollisionLayer(layer_index gd.Int, layer gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, layer_index)
	callframe.Arg(frame, layer)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_set_physics_layer_collision_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the collision layer (as in the physics server) bodies on the given TileSet's physics layer are in.
*/
//go:nosplit
func (self class) GetPhysicsLayerCollisionLayer(layer_index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, layer_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_physics_layer_collision_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the physics layer (as in the physics server) for bodies in the given TileSet physics layer.
*/
//go:nosplit
func (self class) SetPhysicsLayerCollisionMask(layer_index gd.Int, mask gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, layer_index)
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_set_physics_layer_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the collision mask of bodies on the given TileSet's physics layer.
*/
//go:nosplit
func (self class) GetPhysicsLayerCollisionMask(layer_index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, layer_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_physics_layer_collision_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the physics material for bodies in the given TileSet physics layer.
*/
//go:nosplit
func (self class) SetPhysicsLayerPhysicsMaterial(layer_index gd.Int, physics_material gdclass.PhysicsMaterial)  {
	var frame = callframe.New()
	callframe.Arg(frame, layer_index)
	callframe.Arg(frame, discreet.Get(physics_material[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_set_physics_layer_physics_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the physics material of bodies on the given TileSet's physics layer.
*/
//go:nosplit
func (self class) GetPhysicsLayerPhysicsMaterial(layer_index gd.Int) gdclass.PhysicsMaterial {
	var frame = callframe.New()
	callframe.Arg(frame, layer_index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_physics_layer_physics_material, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.PhysicsMaterial{classdb.PhysicsMaterial(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Returns the terrain sets count.
*/
//go:nosplit
func (self class) GetTerrainSetsCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_terrain_sets_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a new terrain set at the given position [param to_position] in the array. If [param to_position] is -1, adds it at the end of the array.
*/
//go:nosplit
func (self class) AddTerrainSet(to_position gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, to_position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_add_terrain_set, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Moves the terrain set at index [param terrain_set] to the given position [param to_position] in the array. Also updates the atlas tiles accordingly.
*/
//go:nosplit
func (self class) MoveTerrainSet(terrain_set gd.Int, to_position gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, terrain_set)
	callframe.Arg(frame, to_position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_move_terrain_set, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the terrain set at index [param terrain_set]. Also updates the atlas tiles accordingly.
*/
//go:nosplit
func (self class) RemoveTerrainSet(terrain_set gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, terrain_set)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_remove_terrain_set, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a terrain mode. Each mode determines which bits of a tile shape is used to match the neighboring tiles' terrains.
*/
//go:nosplit
func (self class) SetTerrainSetMode(terrain_set gd.Int, mode classdb.TileSetTerrainMode)  {
	var frame = callframe.New()
	callframe.Arg(frame, terrain_set)
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_set_terrain_set_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a terrain set mode.
*/
//go:nosplit
func (self class) GetTerrainSetMode(terrain_set gd.Int) classdb.TileSetTerrainMode {
	var frame = callframe.New()
	callframe.Arg(frame, terrain_set)
	var r_ret = callframe.Ret[classdb.TileSetTerrainMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_terrain_set_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of terrains in the given terrain set.
*/
//go:nosplit
func (self class) GetTerrainsCount(terrain_set gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, terrain_set)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_terrains_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a new terrain to the given terrain set [param terrain_set] at the given position [param to_position] in the array. If [param to_position] is -1, adds it at the end of the array.
*/
//go:nosplit
func (self class) AddTerrain(terrain_set gd.Int, to_position gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, terrain_set)
	callframe.Arg(frame, to_position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_add_terrain, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Moves the terrain at index [param terrain_index] for terrain set [param terrain_set] to the given position [param to_position] in the array. Also updates the atlas tiles accordingly.
*/
//go:nosplit
func (self class) MoveTerrain(terrain_set gd.Int, terrain_index gd.Int, to_position gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, terrain_set)
	callframe.Arg(frame, terrain_index)
	callframe.Arg(frame, to_position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_move_terrain, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the terrain at index [param terrain_index] in the given terrain set [param terrain_set]. Also updates the atlas tiles accordingly.
*/
//go:nosplit
func (self class) RemoveTerrain(terrain_set gd.Int, terrain_index gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, terrain_set)
	callframe.Arg(frame, terrain_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_remove_terrain, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets a terrain's name.
*/
//go:nosplit
func (self class) SetTerrainName(terrain_set gd.Int, terrain_index gd.Int, name gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, terrain_set)
	callframe.Arg(frame, terrain_index)
	callframe.Arg(frame, discreet.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_set_terrain_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a terrain's name.
*/
//go:nosplit
func (self class) GetTerrainName(terrain_set gd.Int, terrain_index gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, terrain_set)
	callframe.Arg(frame, terrain_index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_terrain_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets a terrain's color. This color is used for identifying the different terrains in the TileSet editor.
*/
//go:nosplit
func (self class) SetTerrainColor(terrain_set gd.Int, terrain_index gd.Int, color gd.Color)  {
	var frame = callframe.New()
	callframe.Arg(frame, terrain_set)
	callframe.Arg(frame, terrain_index)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_set_terrain_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a terrain's color.
*/
//go:nosplit
func (self class) GetTerrainColor(terrain_set gd.Int, terrain_index gd.Int) gd.Color {
	var frame = callframe.New()
	callframe.Arg(frame, terrain_set)
	callframe.Arg(frame, terrain_index)
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_terrain_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the navigation layers count.
*/
//go:nosplit
func (self class) GetNavigationLayersCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_navigation_layers_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a navigation layer to the TileSet at the given position [param to_position] in the array. If [param to_position] is -1, adds it at the end of the array.
Navigation layers allow assigning a navigable area to atlas tiles.
*/
//go:nosplit
func (self class) AddNavigationLayer(to_position gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, to_position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_add_navigation_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Moves the navigation layer at index [param layer_index] to the given position [param to_position] in the array. Also updates the atlas tiles accordingly.
*/
//go:nosplit
func (self class) MoveNavigationLayer(layer_index gd.Int, to_position gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, layer_index)
	callframe.Arg(frame, to_position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_move_navigation_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the navigation layer at index [param layer_index]. Also updates the atlas tiles accordingly.
*/
//go:nosplit
func (self class) RemoveNavigationLayer(layer_index gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, layer_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_remove_navigation_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the navigation layers (as in the navigation server) for navigation regions in the given TileSet navigation layer.
*/
//go:nosplit
func (self class) SetNavigationLayerLayers(layer_index gd.Int, layers gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, layer_index)
	callframe.Arg(frame, layers)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_set_navigation_layer_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the navigation layers (as in the Navigation server) of the given TileSet navigation layer.
*/
//go:nosplit
func (self class) GetNavigationLayerLayers(layer_index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, layer_index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_navigation_layer_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Based on [param value], enables or disables the specified navigation layer of the TileSet navigation data layer identified by the given [param layer_index], given a navigation_layers [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetNavigationLayerLayerValue(layer_index gd.Int, layer_number gd.Int, value bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, layer_index)
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_set_navigation_layer_layer_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns whether or not the specified navigation layer of the TileSet navigation data layer identified by the given [param layer_index] is enabled, given a navigation_layers [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetNavigationLayerLayerValue(layer_index gd.Int, layer_number gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, layer_index)
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_navigation_layer_layer_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the custom data layers count.
*/
//go:nosplit
func (self class) GetCustomDataLayersCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_custom_data_layers_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a custom data layer to the TileSet at the given position [param to_position] in the array. If [param to_position] is -1, adds it at the end of the array.
Custom data layers allow assigning custom properties to atlas tiles.
*/
//go:nosplit
func (self class) AddCustomDataLayer(to_position gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, to_position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_add_custom_data_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Moves the custom data layer at index [param layer_index] to the given position [param to_position] in the array. Also updates the atlas tiles accordingly.
*/
//go:nosplit
func (self class) MoveCustomDataLayer(layer_index gd.Int, to_position gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, layer_index)
	callframe.Arg(frame, to_position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_move_custom_data_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the custom data layer at index [param layer_index]. Also updates the atlas tiles accordingly.
*/
//go:nosplit
func (self class) RemoveCustomDataLayer(layer_index gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, layer_index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_remove_custom_data_layer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the index of the custom data layer identified by the given name.
*/
//go:nosplit
func (self class) GetCustomDataLayerByName(layer_name gd.String) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(layer_name))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_custom_data_layer_by_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the name of the custom data layer identified by the given index. Names are identifiers of the layer therefore if the name is already taken it will fail and raise an error.
*/
//go:nosplit
func (self class) SetCustomDataLayerName(layer_index gd.Int, layer_name gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, layer_index)
	callframe.Arg(frame, discreet.Get(layer_name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_set_custom_data_layer_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the name of the custom data layer identified by the given index.
*/
//go:nosplit
func (self class) GetCustomDataLayerName(layer_index gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, layer_index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_custom_data_layer_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the type of the custom data layer identified by the given index.
*/
//go:nosplit
func (self class) SetCustomDataLayerType(layer_index gd.Int, layer_type gd.VariantType)  {
	var frame = callframe.New()
	callframe.Arg(frame, layer_index)
	callframe.Arg(frame, layer_type)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_set_custom_data_layer_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the type of the custom data layer identified by the given index.
*/
//go:nosplit
func (self class) GetCustomDataLayerType(layer_index gd.Int) gd.VariantType {
	var frame = callframe.New()
	callframe.Arg(frame, layer_index)
	var r_ret = callframe.Ret[gd.VariantType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_custom_data_layer_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates a source-level proxy for the given source ID. A proxy will map set of tile identifiers to another set of identifiers. Both the atlas coordinates ID and the alternative tile ID are kept the same when using source-level proxies.
This can be used to replace a source in all TileMaps using this TileSet, as TileMap nodes will find and use the proxy's target source when one is available.
Proxied tiles can be automatically replaced in TileMap nodes using the editor.
*/
//go:nosplit
func (self class) SetSourceLevelTileProxy(source_from gd.Int, source_to gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, source_from)
	callframe.Arg(frame, source_to)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_set_source_level_tile_proxy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the source-level proxy for the given source identifier.
If the TileSet has no proxy for the given identifier, returns -1.
*/
//go:nosplit
func (self class) GetSourceLevelTileProxy(source_from gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, source_from)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_source_level_tile_proxy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns if there is a source-level proxy for the given source ID.
*/
//go:nosplit
func (self class) HasSourceLevelTileProxy(source_from gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, source_from)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_has_source_level_tile_proxy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes a source-level tile proxy.
*/
//go:nosplit
func (self class) RemoveSourceLevelTileProxy(source_from gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, source_from)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_remove_source_level_tile_proxy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a coordinates-level proxy for the given identifiers. A proxy will map set of tile identifiers to another set of identifiers. The alternative tile ID is kept the same when using coordinates-level proxies.
This can be used to replace a tile in all TileMaps using this TileSet, as TileMap nodes will find and use the proxy's target tile when one is available.
Proxied tiles can be automatically replaced in TileMap nodes using the editor.
*/
//go:nosplit
func (self class) SetCoordsLevelTileProxy(p_source_from gd.Int, coords_from gd.Vector2i, source_to gd.Int, coords_to gd.Vector2i)  {
	var frame = callframe.New()
	callframe.Arg(frame, p_source_from)
	callframe.Arg(frame, coords_from)
	callframe.Arg(frame, source_to)
	callframe.Arg(frame, coords_to)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_set_coords_level_tile_proxy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the coordinate-level proxy for the given identifiers. The returned array contains the two target identifiers of the proxy (source ID and atlas coordinates ID).
If the TileSet has no proxy for the given identifiers, returns an empty Array.
*/
//go:nosplit
func (self class) GetCoordsLevelTileProxy(source_from gd.Int, coords_from gd.Vector2i) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, source_from)
	callframe.Arg(frame, coords_from)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_coords_level_tile_proxy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns if there is a coodinates-level proxy for the given identifiers.
*/
//go:nosplit
func (self class) HasCoordsLevelTileProxy(source_from gd.Int, coords_from gd.Vector2i) bool {
	var frame = callframe.New()
	callframe.Arg(frame, source_from)
	callframe.Arg(frame, coords_from)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_has_coords_level_tile_proxy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes a coordinates-level proxy for the given identifiers.
*/
//go:nosplit
func (self class) RemoveCoordsLevelTileProxy(source_from gd.Int, coords_from gd.Vector2i)  {
	var frame = callframe.New()
	callframe.Arg(frame, source_from)
	callframe.Arg(frame, coords_from)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_remove_coords_level_tile_proxy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Create an alternative-level proxy for the given identifiers. A proxy will map set of tile identifiers to another set of identifiers.
This can be used to replace a tile in all TileMaps using this TileSet, as TileMap nodes will find and use the proxy's target tile when one is available.
Proxied tiles can be automatically replaced in TileMap nodes using the editor.
*/
//go:nosplit
func (self class) SetAlternativeLevelTileProxy(source_from gd.Int, coords_from gd.Vector2i, alternative_from gd.Int, source_to gd.Int, coords_to gd.Vector2i, alternative_to gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, source_from)
	callframe.Arg(frame, coords_from)
	callframe.Arg(frame, alternative_from)
	callframe.Arg(frame, source_to)
	callframe.Arg(frame, coords_to)
	callframe.Arg(frame, alternative_to)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_set_alternative_level_tile_proxy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the alternative-level proxy for the given identifiers. The returned array contains the three proxie's target identifiers (source ID, atlas coords ID and alternative tile ID).
If the TileSet has no proxy for the given identifiers, returns an empty Array.
*/
//go:nosplit
func (self class) GetAlternativeLevelTileProxy(source_from gd.Int, coords_from gd.Vector2i, alternative_from gd.Int) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, source_from)
	callframe.Arg(frame, coords_from)
	callframe.Arg(frame, alternative_from)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_alternative_level_tile_proxy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns if there is an alternative-level proxy for the given identifiers.
*/
//go:nosplit
func (self class) HasAlternativeLevelTileProxy(source_from gd.Int, coords_from gd.Vector2i, alternative_from gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, source_from)
	callframe.Arg(frame, coords_from)
	callframe.Arg(frame, alternative_from)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_has_alternative_level_tile_proxy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Removes an alternative-level proxy for the given identifiers.
*/
//go:nosplit
func (self class) RemoveAlternativeLevelTileProxy(source_from gd.Int, coords_from gd.Vector2i, alternative_from gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, source_from)
	callframe.Arg(frame, coords_from)
	callframe.Arg(frame, alternative_from)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_remove_alternative_level_tile_proxy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
According to the configured proxies, maps the provided identifiers to a new set of identifiers. The source ID, atlas coordinates ID and alternative tile ID are returned as a 3 elements Array.
This function first look for matching alternative-level proxies, then coordinates-level proxies, then source-level proxies.
If no proxy corresponding to provided identifiers are found, returns the same values the ones used as arguments.
*/
//go:nosplit
func (self class) MapTileProxy(source_from gd.Int, coords_from gd.Vector2i, alternative_from gd.Int) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, source_from)
	callframe.Arg(frame, coords_from)
	callframe.Arg(frame, alternative_from)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_map_tile_proxy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Clears tile proxies pointing to invalid tiles.
*/
//go:nosplit
func (self class) CleanupInvalidTileProxies()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_cleanup_invalid_tile_proxies, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Clears all tile proxies.
*/
//go:nosplit
func (self class) ClearTileProxies()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_clear_tile_proxies, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a [TileMapPattern] to be stored in the TileSet resource. If provided, insert it at the given [param index].
*/
//go:nosplit
func (self class) AddPattern(pattern gdclass.TileMapPattern, index gd.Int) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(pattern[0])[0])
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_add_pattern, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [TileMapPattern] at the given [param index].
*/
//go:nosplit
func (self class) GetPattern(index gd.Int) gdclass.TileMapPattern {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_pattern, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.TileMapPattern{classdb.TileMapPattern(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Remove the [TileMapPattern] at the given index.
*/
//go:nosplit
func (self class) RemovePattern(index gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_remove_pattern, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the number of [TileMapPattern] this tile set handles.
*/
//go:nosplit
func (self class) GetPatternsCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSet.Bind_get_patterns_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsTileSet() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsTileSet() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("TileSet", func(ptr gd.Object) any { return classdb.TileSet(ptr) })}
type TileShape = classdb.TileSetTileShape

const (
/*Rectangular tile shape.*/
	TileShapeSquare TileShape = 0
/*Diamond tile shape (for isometric look).
[b]Note:[/b] Isometric [TileSet] works best if [TileMap] and all its layers have Y-sort enabled.*/
	TileShapeIsometric TileShape = 1
/*Rectangular tile shape with one row/column out of two offset by half a tile.*/
	TileShapeHalfOffsetSquare TileShape = 2
/*Hexagonal tile shape.*/
	TileShapeHexagon TileShape = 3
)
type TileLayout = classdb.TileSetTileLayout

const (
/*Tile coordinates layout where both axis stay consistent with their respective local horizontal and vertical axis.*/
	TileLayoutStacked TileLayout = 0
/*Same as [constant TILE_LAYOUT_STACKED], but the first half-offset is negative instead of positive.*/
	TileLayoutStackedOffset TileLayout = 1
/*Tile coordinates layout where the horizontal axis stay horizontal, and the vertical one goes down-right.*/
	TileLayoutStairsRight TileLayout = 2
/*Tile coordinates layout where the vertical axis stay vertical, and the horizontal one goes down-right.*/
	TileLayoutStairsDown TileLayout = 3
/*Tile coordinates layout where the horizontal axis goes up-right, and the vertical one goes down-right.*/
	TileLayoutDiamondRight TileLayout = 4
/*Tile coordinates layout where the horizontal axis goes down-right, and the vertical one goes down-left.*/
	TileLayoutDiamondDown TileLayout = 5
)
type TileOffsetAxis = classdb.TileSetTileOffsetAxis

const (
/*Horizontal half-offset.*/
	TileOffsetAxisHorizontal TileOffsetAxis = 0
/*Vertical half-offset.*/
	TileOffsetAxisVertical TileOffsetAxis = 1
)
type CellNeighbor = classdb.TileSetCellNeighbor

const (
/*Neighbor on the right side.*/
	CellNeighborRightSide CellNeighbor = 0
/*Neighbor in the right corner.*/
	CellNeighborRightCorner CellNeighbor = 1
/*Neighbor on the bottom right side.*/
	CellNeighborBottomRightSide CellNeighbor = 2
/*Neighbor in the bottom right corner.*/
	CellNeighborBottomRightCorner CellNeighbor = 3
/*Neighbor on the bottom side.*/
	CellNeighborBottomSide CellNeighbor = 4
/*Neighbor in the bottom corner.*/
	CellNeighborBottomCorner CellNeighbor = 5
/*Neighbor on the bottom left side.*/
	CellNeighborBottomLeftSide CellNeighbor = 6
/*Neighbor in the bottom left corner.*/
	CellNeighborBottomLeftCorner CellNeighbor = 7
/*Neighbor on the left side.*/
	CellNeighborLeftSide CellNeighbor = 8
/*Neighbor in the left corner.*/
	CellNeighborLeftCorner CellNeighbor = 9
/*Neighbor on the top left side.*/
	CellNeighborTopLeftSide CellNeighbor = 10
/*Neighbor in the top left corner.*/
	CellNeighborTopLeftCorner CellNeighbor = 11
/*Neighbor on the top side.*/
	CellNeighborTopSide CellNeighbor = 12
/*Neighbor in the top corner.*/
	CellNeighborTopCorner CellNeighbor = 13
/*Neighbor on the top right side.*/
	CellNeighborTopRightSide CellNeighbor = 14
/*Neighbor in the top right corner.*/
	CellNeighborTopRightCorner CellNeighbor = 15
)
type TerrainMode = classdb.TileSetTerrainMode

const (
/*Requires both corners and side to match with neighboring tiles' terrains.*/
	TerrainModeMatchCornersAndSides TerrainMode = 0
/*Requires corners to match with neighboring tiles' terrains.*/
	TerrainModeMatchCorners TerrainMode = 1
/*Requires sides to match with neighboring tiles' terrains.*/
	TerrainModeMatchSides TerrainMode = 2
)
