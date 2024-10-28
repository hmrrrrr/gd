package EditorExportPlatformLinuxBSD

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/EditorExportPlatformPC"
import "grow.graphics/gd/gdclass/EditorExportPlatform"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

type Go [1]classdb.EditorExportPlatformLinuxBSD
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.EditorExportPlatformLinuxBSD
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorExportPlatformLinuxBSD"))
	return Go{classdb.EditorExportPlatformLinuxBSD(object)}
}

func (self class) AsEditorExportPlatformLinuxBSD() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorExportPlatformLinuxBSD() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsEditorExportPlatformPC() EditorExportPlatformPC.GD { return *((*EditorExportPlatformPC.GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorExportPlatformPC() EditorExportPlatformPC.Go { return *((*EditorExportPlatformPC.Go)(unsafe.Pointer(&self))) }
func (self class) AsEditorExportPlatform() EditorExportPlatform.GD { return *((*EditorExportPlatform.GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorExportPlatform() EditorExportPlatform.Go { return *((*EditorExportPlatform.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsEditorExportPlatformPC(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsEditorExportPlatformPC(), name)
	}
}
func init() {classdb.Register("EditorExportPlatformLinuxBSD", func(ptr gd.Object) any { return classdb.EditorExportPlatformLinuxBSD(ptr) })}
