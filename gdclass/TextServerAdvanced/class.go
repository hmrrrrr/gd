package TextServerAdvanced

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/TextServerExtension"
import "grow.graphics/gd/gdclass/TextServer"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
An implementation of [TextServer] that uses HarfBuzz, ICU and SIL Graphite to support BiDi, complex text layouts and contextual OpenType features. This is Godot's default primary [TextServer] interface.

*/
type Go [1]classdb.TextServerAdvanced
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.TextServerAdvanced
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("TextServerAdvanced"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self class) AsTextServerAdvanced() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsTextServerAdvanced() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsTextServerExtension() TextServerExtension.GD { return *((*TextServerExtension.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTextServerExtension() TextServerExtension.Go { return *((*TextServerExtension.Go)(unsafe.Pointer(&self))) }
func (self class) AsTextServer() TextServer.GD { return *((*TextServer.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTextServer() TextServer.Go { return *((*TextServer.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTextServerExtension(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTextServerExtension(), name)
	}
}
func init() {classdb.Register("TextServerAdvanced", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
