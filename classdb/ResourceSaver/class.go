// Package ResourceSaver provides methods for working with ResourceSaver object instances.
package ResourceSaver

import "unsafe"
import "sync"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
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
A singleton for saving resource types to the filesystem.
It uses the many [ResourceFormatSaver] classes registered in the engine (either built-in or from a plugin) to save resource data to text-based (e.g. [code].tres[/code] or [code].tscn[/code]) or binary files (e.g. [code].res[/code] or [code].scn[/code]).
*/
var self [1]gdclass.ResourceSaver
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.ResourceSaver)
	self = *(*[1]gdclass.ResourceSaver)(unsafe.Pointer(&obj))
}

/*
Saves a resource to disk to the given path, using a [ResourceFormatSaver] that recognizes the resource object. If [param path] is empty, [ResourceSaver] will try to use [member Resource.resource_path].
The [param flags] bitmask can be specified to customize the save behavior using [enum SaverFlags] flags.
Returns [constant OK] on success.
[b]Note:[/b] When the project is running, any generated UID associated with the resource will not be saved as the required code is only executed in editor mode.
*/
func Save(resource [1]gdclass.Resource) error { //gd:ResourceSaver.save
	once.Do(singleton)
	return error(gd.ToError(class(self).Save(resource, String.New(""), 0)))
}

/*
Returns the list of extensions available for saving a resource of a given type.
*/
func GetRecognizedExtensions(atype [1]gdclass.Resource) []string { //gd:ResourceSaver.get_recognized_extensions
	once.Do(singleton)
	return []string(class(self).GetRecognizedExtensions(atype).Strings())
}

/*
Registers a new [ResourceFormatSaver]. The ResourceSaver will use the ResourceFormatSaver as described in [method save].
This method is performed implicitly for ResourceFormatSavers written in GDScript (see [ResourceFormatSaver] for more information).
*/
func AddResourceFormatSaver(format_saver [1]gdclass.ResourceFormatSaver) { //gd:ResourceSaver.add_resource_format_saver
	once.Do(singleton)
	class(self).AddResourceFormatSaver(format_saver, false)
}

/*
Unregisters the given [ResourceFormatSaver].
*/
func RemoveResourceFormatSaver(format_saver [1]gdclass.ResourceFormatSaver) { //gd:ResourceSaver.remove_resource_format_saver
	once.Do(singleton)
	class(self).RemoveResourceFormatSaver(format_saver)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]gdclass.ResourceSaver

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Saves a resource to disk to the given path, using a [ResourceFormatSaver] that recognizes the resource object. If [param path] is empty, [ResourceSaver] will try to use [member Resource.resource_path].
The [param flags] bitmask can be specified to customize the save behavior using [enum SaverFlags] flags.
Returns [constant OK] on success.
[b]Note:[/b] When the project is running, any generated UID associated with the resource will not be saved as the required code is only executed in editor mode.
*/
//go:nosplit
func (self class) Save(resource [1]gdclass.Resource, path String.Readable, flags gdclass.ResourceSaverSaverFlags) Error.Code { //gd:ResourceSaver.save
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(resource[0])[0])
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	callframe.Arg(frame, flags)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceSaver.Bind_save, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the list of extensions available for saving a resource of a given type.
*/
//go:nosplit
func (self class) GetRecognizedExtensions(atype [1]gdclass.Resource) Packed.Strings { //gd:ResourceSaver.get_recognized_extensions
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(atype[0])[0])
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceSaver.Bind_get_recognized_extensions, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Registers a new [ResourceFormatSaver]. The ResourceSaver will use the ResourceFormatSaver as described in [method save].
This method is performed implicitly for ResourceFormatSavers written in GDScript (see [ResourceFormatSaver] for more information).
*/
//go:nosplit
func (self class) AddResourceFormatSaver(format_saver [1]gdclass.ResourceFormatSaver, at_front bool) { //gd:ResourceSaver.add_resource_format_saver
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(format_saver[0])[0])
	callframe.Arg(frame, at_front)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceSaver.Bind_add_resource_format_saver, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Unregisters the given [ResourceFormatSaver].
*/
//go:nosplit
func (self class) RemoveResourceFormatSaver(format_saver [1]gdclass.ResourceFormatSaver) { //gd:ResourceSaver.remove_resource_format_saver
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(format_saver[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceSaver.Bind_remove_resource_format_saver, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("ResourceSaver", func(ptr gd.Object) any {
		return [1]gdclass.ResourceSaver{*(*gdclass.ResourceSaver)(unsafe.Pointer(&ptr))}
	})
}

type SaverFlags = gdclass.ResourceSaverSaverFlags //gd:ResourceSaver.SaverFlags

const (
	/*No resource saving option.*/
	FlagNone SaverFlags = 0
	/*Save the resource with a path relative to the scene which uses it.*/
	FlagRelativePaths SaverFlags = 1
	/*Bundles external resources.*/
	FlagBundleResources SaverFlags = 2
	/*Changes the [member Resource.resource_path] of the saved resource to match its new location.*/
	FlagChangePath SaverFlags = 4
	/*Do not save editor-specific metadata (identified by their [code]__editor[/code] prefix).*/
	FlagOmitEditorProperties SaverFlags = 8
	/*Save as big endian (see [member FileAccess.big_endian]).*/
	FlagSaveBigEndian SaverFlags = 16
	/*Compress the resource on save using [constant FileAccess.COMPRESSION_ZSTD]. Only available for binary resource types.*/
	FlagCompress SaverFlags = 32
	/*Take over the paths of the saved subresources (see [method Resource.take_over_path]).*/
	FlagReplaceSubresourcePaths SaverFlags = 64
)
