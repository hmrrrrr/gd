// Package ResourceFormatSaver provides methods for working with ResourceFormatSaver object instances.
package ResourceFormatSaver

import "unsafe"
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
The engine can save resources when you do it from the editor, or when you use the [ResourceSaver] singleton. This is accomplished thanks to multiple [ResourceFormatSaver]s, each handling its own format and called automatically by the engine.
By default, Godot saves resources as [code].tres[/code] (text-based), [code].res[/code] (binary) or another built-in format, but you can choose to create your own format by extending this class. Be sure to respect the documented return types and values. You should give it a global class name with [code]class_name[/code] for it to be registered. Like built-in ResourceFormatSavers, it will be called automatically when saving resources of its recognized type(s). You may also implement a [ResourceFormatLoader].

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=ResourceFormatSaver)
*/
type Instance [1]gdclass.ResourceFormatSaver

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsResourceFormatSaver() Instance
}
type Interface interface {
	//Saves the given resource object to a file at the target [param path]. [param flags] is a bitmask composed with [enum ResourceSaver.SaverFlags] constants.
	//Returns [constant OK] on success, or an [enum Error] constant in case of failure.
	Save(resource [1]gdclass.Resource, path string, flags int) error
	//Sets a new UID for the resource at the given [param path]. Returns [constant OK] on success, or an [enum Error] constant in case of failure.
	SetUid(path string, uid int) error
	//Returns whether the given resource object can be saved by this saver.
	Recognize(resource [1]gdclass.Resource) bool
	//Returns the list of extensions available for saving the resource object, provided it is recognized (see [method _recognize]).
	GetRecognizedExtensions(resource [1]gdclass.Resource) []string
	//Returns [code]true[/code] if this saver handles a given save path and [code]false[/code] otherwise.
	//If this method is not implemented, the default behavior returns whether the path's extension is within the ones provided by [method _get_recognized_extensions].
	RecognizePath(resource [1]gdclass.Resource, path string) bool
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) Save(resource [1]gdclass.Resource, path string, flags int) (_ error) {
	return
}
func (self implementation) SetUid(path string, uid int) (_ error)                             { return }
func (self implementation) Recognize(resource [1]gdclass.Resource) (_ bool)                   { return }
func (self implementation) GetRecognizedExtensions(resource [1]gdclass.Resource) (_ []string) { return }
func (self implementation) RecognizePath(resource [1]gdclass.Resource, path string) (_ bool)  { return }

/*
Saves the given resource object to a file at the target [param path]. [param flags] is a bitmask composed with [enum ResourceSaver.SaverFlags] constants.
Returns [constant OK] on success, or an [enum Error] constant in case of failure.
*/
func (Instance) _save(impl func(ptr unsafe.Pointer, resource [1]gdclass.Resource, path string, flags int) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var resource = [1]gdclass.Resource{pointers.New[gdclass.Resource]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(resource[0])
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(path))
		var flags = gd.UnsafeGet[int64](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource, path.String(), int(flags))
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(Error.New(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Sets a new UID for the resource at the given [param path]. Returns [constant OK] on success, or an [enum Error] constant in case of failure.
*/
func (Instance) _set_uid(impl func(ptr unsafe.Pointer, path string, uid int) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(path))
		var uid = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String(), int(uid))
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(Error.New(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns whether the given resource object can be saved by this saver.
*/
func (Instance) _recognize(impl func(ptr unsafe.Pointer, resource [1]gdclass.Resource) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var resource = [1]gdclass.Resource{pointers.New[gdclass.Resource]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(resource[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns the list of extensions available for saving the resource object, provided it is recognized (see [method _recognize]).
*/
func (Instance) _get_recognized_extensions(impl func(ptr unsafe.Pointer, resource [1]gdclass.Resource) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var resource = [1]gdclass.Resource{pointers.New[gdclass.Resource]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(resource[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource)
		ptr, ok := pointers.End(gd.InternalPackedStrings(Packed.MakeStrings(ret...)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns [code]true[/code] if this saver handles a given save path and [code]false[/code] otherwise.
If this method is not implemented, the default behavior returns whether the path's extension is within the ones provided by [method _get_recognized_extensions].
*/
func (Instance) _recognize_path(impl func(ptr unsafe.Pointer, resource [1]gdclass.Resource, path string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var resource = [1]gdclass.Resource{pointers.New[gdclass.Resource]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(resource[0])
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(path))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource, path.String())
		gd.UnsafeSet(p_back, ret)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ResourceFormatSaver

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ResourceFormatSaver"))
	casted := Instance{*(*gdclass.ResourceFormatSaver)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Saves the given resource object to a file at the target [param path]. [param flags] is a bitmask composed with [enum ResourceSaver.SaverFlags] constants.
Returns [constant OK] on success, or an [enum Error] constant in case of failure.
*/
func (class) _save(impl func(ptr unsafe.Pointer, resource [1]gdclass.Resource, path String.Readable, flags int64) Error.Code) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var resource = [1]gdclass.Resource{pointers.New[gdclass.Resource]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(resource[0])
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(path))
		var flags = gd.UnsafeGet[int64](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource, path, flags)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Sets a new UID for the resource at the given [param path]. Returns [constant OK] on success, or an [enum Error] constant in case of failure.
*/
func (class) _set_uid(impl func(ptr unsafe.Pointer, path String.Readable, uid int64) Error.Code) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(path))
		var uid = gd.UnsafeGet[int64](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path, uid)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns whether the given resource object can be saved by this saver.
*/
func (class) _recognize(impl func(ptr unsafe.Pointer, resource [1]gdclass.Resource) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var resource = [1]gdclass.Resource{pointers.New[gdclass.Resource]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(resource[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns the list of extensions available for saving the resource object, provided it is recognized (see [method _recognize]).
*/
func (class) _get_recognized_extensions(impl func(ptr unsafe.Pointer, resource [1]gdclass.Resource) Packed.Strings) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var resource = [1]gdclass.Resource{pointers.New[gdclass.Resource]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(resource[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource)
		ptr, ok := pointers.End(gd.InternalPackedStrings(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns [code]true[/code] if this saver handles a given save path and [code]false[/code] otherwise.
If this method is not implemented, the default behavior returns whether the path's extension is within the ones provided by [method _get_recognized_extensions].
*/
func (class) _recognize_path(impl func(ptr unsafe.Pointer, resource [1]gdclass.Resource, path String.Readable) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var resource = [1]gdclass.Resource{pointers.New[gdclass.Resource]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(resource[0])
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(path))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource, path)
		gd.UnsafeSet(p_back, ret)
	}
}

func (self class) AsResourceFormatSaver() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsResourceFormatSaver() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_save":
		return reflect.ValueOf(self._save)
	case "_set_uid":
		return reflect.ValueOf(self._set_uid)
	case "_recognize":
		return reflect.ValueOf(self._recognize)
	case "_get_recognized_extensions":
		return reflect.ValueOf(self._get_recognized_extensions)
	case "_recognize_path":
		return reflect.ValueOf(self._recognize_path)
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_save":
		return reflect.ValueOf(self._save)
	case "_set_uid":
		return reflect.ValueOf(self._set_uid)
	case "_recognize":
		return reflect.ValueOf(self._recognize)
	case "_get_recognized_extensions":
		return reflect.ValueOf(self._get_recognized_extensions)
	case "_recognize_path":
		return reflect.ValueOf(self._recognize_path)
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("ResourceFormatSaver", func(ptr gd.Object) any {
		return [1]gdclass.ResourceFormatSaver{*(*gdclass.ResourceFormatSaver)(unsafe.Pointer(&ptr))}
	})
}
