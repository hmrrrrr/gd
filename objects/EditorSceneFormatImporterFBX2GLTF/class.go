package EditorSceneFormatImporterFBX2GLTF

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/EditorSceneFormatImporter"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Imports Autodesk FBX 3D scenes by way of converting them to glTF 2.0 using the FBX2glTF command line tool.
The location of the FBX2glTF binary is set via the [member EditorSettings.filesystem/import/fbx/fbx2gltf_path] editor setting.
This importer is only used if [member ProjectSettings.filesystem/import/fbx2gltf/enabled] is set to [code]true[/code].
*/
type Instance [1]classdb.EditorSceneFormatImporterFBX2GLTF
type Any interface {
	gd.IsClass
	AsEditorSceneFormatImporterFBX2GLTF() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.EditorSceneFormatImporterFBX2GLTF

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorSceneFormatImporterFBX2GLTF"))
	return Instance{classdb.EditorSceneFormatImporterFBX2GLTF(object)}
}

func (self class) AsEditorSceneFormatImporterFBX2GLTF() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsEditorSceneFormatImporterFBX2GLTF() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsEditorSceneFormatImporter() EditorSceneFormatImporter.Advanced {
	return *((*EditorSceneFormatImporter.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsEditorSceneFormatImporter() EditorSceneFormatImporter.Instance {
	return *((*EditorSceneFormatImporter.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsEditorSceneFormatImporter(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsEditorSceneFormatImporter(), name)
	}
}
func init() {
	classdb.Register("EditorSceneFormatImporterFBX2GLTF", func(ptr gd.Object) any {
		return [1]classdb.EditorSceneFormatImporterFBX2GLTF{classdb.EditorSceneFormatImporterFBX2GLTF(ptr)}
	})
}
