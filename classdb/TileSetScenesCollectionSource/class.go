// Package TileSetScenesCollectionSource provides methods for working with TileSetScenesCollectionSource object instances.
package TileSetScenesCollectionSource

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Resource"
import "graphics.gd/classdb/TileSetSource"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

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
When placed on a [TileMap], tiles from [TileSetScenesCollectionSource] will automatically instantiate an associated scene at the cell's position in the TileMap.
Scenes are instantiated as children of the [TileMap] when it enters the tree. If you add/remove a scene tile in the [TileMap] that is already inside the tree, the [TileMap] will automatically instantiate/free the scene accordingly.
[b]Note:[/b] Scene tiles all occupy one tile slot and instead use alternate tile ID to identify scene index. [method TileSetSource.get_tiles_count] will always return [code]1[/code]. Use [method get_scene_tiles_count] to get a number of scenes in a [TileSetScenesCollectionSource].
Use this code if you want to find the scene path at a given tile in [TileMapLayer]:
[codeblocks]
[gdscript]
var source_id = tile_map_layer.get_cell_source_id(Vector2i(x, y))
if source_id > -1:

	var scene_source = tile_map_layer.tile_set.get_source(source_id)
	if scene_source is TileSetScenesCollectionSource:
	    var alt_id = tile_map_layer.get_cell_alternative_tile(Vector2i(x, y))
	    # The assigned PackedScene.
	    var scene = scene_source.get_scene_tile_scene(alt_id)

[/gdscript]
[csharp]
int sourceId = tileMapLayer.GetCellSourceId(new Vector2I(x, y));
if (sourceId > -1)

	{
	    TileSetSource source = tileMapLayer.TileSet.GetSource(sourceId);
	    if (source is TileSetScenesCollectionSource sceneSource)
	    {
	        int altId = tileMapLayer.GetCellAlternativeTile(new Vector2I(x, y));
	        // The assigned PackedScene.
	        PackedScene scene = sceneSource.GetSceneTileScene(altId);
	    }
	}

[/csharp]
[/codeblocks]
*/
type Instance [1]gdclass.TileSetScenesCollectionSource

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsTileSetScenesCollectionSource() Instance
}

/*
Returns the number or scene tiles this TileSet source has.
*/
func (self Instance) GetSceneTilesCount() int { //gd:TileSetScenesCollectionSource.get_scene_tiles_count
	return int(int(class(self).GetSceneTilesCount()))
}

/*
Returns the scene tile ID of the scene tile at [param index].
*/
func (self Instance) GetSceneTileId(index int) int { //gd:TileSetScenesCollectionSource.get_scene_tile_id
	return int(int(class(self).GetSceneTileId(int64(index))))
}

/*
Returns whether this TileSet source has a scene tile with [param id].
*/
func (self Instance) HasSceneTileId(id int) bool { //gd:TileSetScenesCollectionSource.has_scene_tile_id
	return bool(class(self).HasSceneTileId(int64(id)))
}

/*
Creates a scene-based tile out of the given scene.
Returns a newly generated unique ID.
*/
func (self Instance) CreateSceneTile(packed_scene [1]gdclass.PackedScene) int { //gd:TileSetScenesCollectionSource.create_scene_tile
	return int(int(class(self).CreateSceneTile(packed_scene, int64(-1))))
}

/*
Changes a scene tile's ID from [param id] to [param new_id]. This will fail if there is already a tile with an ID equal to [param new_id].
*/
func (self Instance) SetSceneTileId(id int, new_id int) { //gd:TileSetScenesCollectionSource.set_scene_tile_id
	class(self).SetSceneTileId(int64(id), int64(new_id))
}

/*
Assigns a [PackedScene] resource to the scene tile with [param id]. This will fail if the scene does not extend CanvasItem, as positioning properties are needed to place the scene on the TileMap.
*/
func (self Instance) SetSceneTileScene(id int, packed_scene [1]gdclass.PackedScene) { //gd:TileSetScenesCollectionSource.set_scene_tile_scene
	class(self).SetSceneTileScene(int64(id), packed_scene)
}

/*
Returns the [PackedScene] resource of scene tile with [param id].
*/
func (self Instance) GetSceneTileScene(id int) [1]gdclass.PackedScene { //gd:TileSetScenesCollectionSource.get_scene_tile_scene
	return [1]gdclass.PackedScene(class(self).GetSceneTileScene(int64(id)))
}

/*
Sets whether or not the scene tile with [param id] should display a placeholder in the editor. This might be useful for scenes that are not visible.
*/
func (self Instance) SetSceneTileDisplayPlaceholder(id int, display_placeholder bool) { //gd:TileSetScenesCollectionSource.set_scene_tile_display_placeholder
	class(self).SetSceneTileDisplayPlaceholder(int64(id), display_placeholder)
}

/*
Returns whether the scene tile with [param id] displays a placeholder in the editor.
*/
func (self Instance) GetSceneTileDisplayPlaceholder(id int) bool { //gd:TileSetScenesCollectionSource.get_scene_tile_display_placeholder
	return bool(class(self).GetSceneTileDisplayPlaceholder(int64(id)))
}

/*
Remove the scene tile with [param id].
*/
func (self Instance) RemoveSceneTile(id int) { //gd:TileSetScenesCollectionSource.remove_scene_tile
	class(self).RemoveSceneTile(int64(id))
}

/*
Returns the scene ID a following call to [method create_scene_tile] would return.
*/
func (self Instance) GetNextSceneTileId() int { //gd:TileSetScenesCollectionSource.get_next_scene_tile_id
	return int(int(class(self).GetNextSceneTileId()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.TileSetScenesCollectionSource

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TileSetScenesCollectionSource"))
	casted := Instance{*(*gdclass.TileSetScenesCollectionSource)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Returns the number or scene tiles this TileSet source has.
*/
//go:nosplit
func (self class) GetSceneTilesCount() int64 { //gd:TileSetScenesCollectionSource.get_scene_tiles_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetScenesCollectionSource.Bind_get_scene_tiles_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the scene tile ID of the scene tile at [param index].
*/
//go:nosplit
func (self class) GetSceneTileId(index int64) int64 { //gd:TileSetScenesCollectionSource.get_scene_tile_id
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetScenesCollectionSource.Bind_get_scene_tile_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns whether this TileSet source has a scene tile with [param id].
*/
//go:nosplit
func (self class) HasSceneTileId(id int64) bool { //gd:TileSetScenesCollectionSource.has_scene_tile_id
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetScenesCollectionSource.Bind_has_scene_tile_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a scene-based tile out of the given scene.
Returns a newly generated unique ID.
*/
//go:nosplit
func (self class) CreateSceneTile(packed_scene [1]gdclass.PackedScene, id_override int64) int64 { //gd:TileSetScenesCollectionSource.create_scene_tile
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(packed_scene[0])[0])
	callframe.Arg(frame, id_override)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetScenesCollectionSource.Bind_create_scene_tile, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Changes a scene tile's ID from [param id] to [param new_id]. This will fail if there is already a tile with an ID equal to [param new_id].
*/
//go:nosplit
func (self class) SetSceneTileId(id int64, new_id int64) { //gd:TileSetScenesCollectionSource.set_scene_tile_id
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, new_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetScenesCollectionSource.Bind_set_scene_tile_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Assigns a [PackedScene] resource to the scene tile with [param id]. This will fail if the scene does not extend CanvasItem, as positioning properties are needed to place the scene on the TileMap.
*/
//go:nosplit
func (self class) SetSceneTileScene(id int64, packed_scene [1]gdclass.PackedScene) { //gd:TileSetScenesCollectionSource.set_scene_tile_scene
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, pointers.Get(packed_scene[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetScenesCollectionSource.Bind_set_scene_tile_scene, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [PackedScene] resource of scene tile with [param id].
*/
//go:nosplit
func (self class) GetSceneTileScene(id int64) [1]gdclass.PackedScene { //gd:TileSetScenesCollectionSource.get_scene_tile_scene
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetScenesCollectionSource.Bind_get_scene_tile_scene, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.PackedScene{gd.PointerWithOwnershipTransferredToGo[gdclass.PackedScene](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sets whether or not the scene tile with [param id] should display a placeholder in the editor. This might be useful for scenes that are not visible.
*/
//go:nosplit
func (self class) SetSceneTileDisplayPlaceholder(id int64, display_placeholder bool) { //gd:TileSetScenesCollectionSource.set_scene_tile_display_placeholder
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, display_placeholder)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetScenesCollectionSource.Bind_set_scene_tile_display_placeholder, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether the scene tile with [param id] displays a placeholder in the editor.
*/
//go:nosplit
func (self class) GetSceneTileDisplayPlaceholder(id int64) bool { //gd:TileSetScenesCollectionSource.get_scene_tile_display_placeholder
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetScenesCollectionSource.Bind_get_scene_tile_display_placeholder, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Remove the scene tile with [param id].
*/
//go:nosplit
func (self class) RemoveSceneTile(id int64) { //gd:TileSetScenesCollectionSource.remove_scene_tile
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetScenesCollectionSource.Bind_remove_scene_tile, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the scene ID a following call to [method create_scene_tile] would return.
*/
//go:nosplit
func (self class) GetNextSceneTileId() int64 { //gd:TileSetScenesCollectionSource.get_next_scene_tile_id
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TileSetScenesCollectionSource.Bind_get_next_scene_tile_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsTileSetScenesCollectionSource() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsTileSetScenesCollectionSource() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsTileSetSource() TileSetSource.Advanced {
	return *((*TileSetSource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsTileSetSource() TileSetSource.Instance {
	return *((*TileSetSource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(TileSetSource.Advanced(self.AsTileSetSource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(TileSetSource.Instance(self.AsTileSetSource()), name)
	}
}
func init() {
	gdclass.Register("TileSetScenesCollectionSource", func(ptr gd.Object) any {
		return [1]gdclass.TileSetScenesCollectionSource{*(*gdclass.TileSetScenesCollectionSource)(unsafe.Pointer(&ptr))}
	})
}
