package ResourceImporterTextureAtlas

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/ResourceImporter"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This imports a collection of textures from a PNG image into an [AtlasTexture] or 2D [ArrayMesh]. This can be used to save memory when importing 2D animations from spritesheets. Texture atlases are only supported in 2D rendering, not 3D. See also [ResourceImporterTexture] and [ResourceImporterLayeredTexture].
[b]Note:[/b] [ResourceImporterTextureAtlas] does not handle importing [TileSetAtlasSource], which is created using the [TileSet] editor instead.

*/
type Simple [1]classdb.ResourceImporterTextureAtlas
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ResourceImporterTextureAtlas
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsResourceImporterTextureAtlas() Expert { return self[0].AsResourceImporterTextureAtlas() }


//go:nosplit
func (self Simple) AsResourceImporterTextureAtlas() Simple { return self[0].AsResourceImporterTextureAtlas() }


//go:nosplit
func (self class) AsResourceImporter() ResourceImporter.Expert { return self[0].AsResourceImporter() }


//go:nosplit
func (self Simple) AsResourceImporter() ResourceImporter.Simple { return self[0].AsResourceImporter() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("ResourceImporterTextureAtlas", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
