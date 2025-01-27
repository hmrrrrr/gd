// Package GDExtension provides methods for working with GDExtension object instances.
package GDExtension

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/classdb/Resource"

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

/*
The [GDExtension] resource type represents a [url=https://en.wikipedia.org/wiki/Shared_library]shared library[/url] which can expand the functionality of the engine. The [GDExtensionManager] singleton is responsible for loading, reloading, and unloading [GDExtension] resources.
[b]Note:[/b] GDExtension itself is not a scripting language and has no relation to [GDScript] resources.
*/
type Instance [1]gdclass.GDExtension

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGDExtension() Instance
}

/*
Returns [code]true[/code] if this extension's library has been opened.
*/
func (self Instance) IsLibraryOpen() bool { //gd:GDExtension.is_library_open
	return bool(class(self).IsLibraryOpen())
}

/*
Returns the lowest level required for this extension to be properly initialized (see the [enum InitializationLevel] enum).
*/
func (self Instance) GetMinimumLibraryInitializationLevel() gd.GDExtensionInitializationLevel { //gd:GDExtension.get_minimum_library_initialization_level
	return gd.GDExtensionInitializationLevel(class(self).GetMinimumLibraryInitializationLevel())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GDExtension

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GDExtension"))
	casted := Instance{*(*gdclass.GDExtension)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Returns [code]true[/code] if this extension's library has been opened.
*/
//go:nosplit
func (self class) IsLibraryOpen() bool { //gd:GDExtension.is_library_open
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GDExtension.Bind_is_library_open, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the lowest level required for this extension to be properly initialized (see the [enum InitializationLevel] enum).
*/
//go:nosplit
func (self class) GetMinimumLibraryInitializationLevel() gd.GDExtensionInitializationLevel { //gd:GDExtension.get_minimum_library_initialization_level
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.GDExtensionInitializationLevel](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GDExtension.Bind_get_minimum_library_initialization_level, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsGDExtension() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGDExtension() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("GDExtension", func(ptr gd.Object) any { return [1]gdclass.GDExtension{*(*gdclass.GDExtension)(unsafe.Pointer(&ptr))} })
}

type InitializationLevel = gdclass.GDExtensionInitializationLevel //gd:GDExtension.InitializationLevel

const (
	/*The library is initialized at the same time as the core features of the engine.*/
	InitializationLevelCore InitializationLevel = 0
	/*The library is initialized at the same time as the engine's servers (such as [RenderingServer] or [PhysicsServer3D]).*/
	InitializationLevelServers InitializationLevel = 1
	/*The library is initialized at the same time as the engine's scene-related classes.*/
	InitializationLevelScene InitializationLevel = 2
	/*The library is initialized at the same time as the engine's editor classes. Only happens when loading the GDExtension in the editor.*/
	InitializationLevelEditor InitializationLevel = 3
)
