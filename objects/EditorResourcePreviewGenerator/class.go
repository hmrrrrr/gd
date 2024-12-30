package EditorResourcePreviewGenerator

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/variant/Vector2i"
import "graphics.gd/variant/Dictionary"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Custom code to generate previews. Please check [code]file_dialog/thumbnail_size[/code] in [EditorSettings] to find out the right size to do previews at.

	// EditorResourcePreviewGenerator methods that can be overridden by a [Class] that extends it.
	type EditorResourcePreviewGenerator interface {
		//Returns [code]true[/code] if your generator supports the resource of type [param type].
		Handles(atype string) bool
		//Generate a preview from a given resource with the specified size. This must always be implemented.
		//Returning an empty texture is an OK way to fail and let another generator take care.
		//Care must be taken because this function is always called from a thread (not the main thread).
		//[param metadata] dictionary can be modified to store file-specific metadata that can be used in [method EditorResourceTooltipPlugin._make_tooltip_for_path] (like image size, sample length etc.).
		Generate(resource objects.Resource, size Vector2i.XY, metadata Dictionary.Any) objects.Texture2D
		//Generate a preview directly from a path with the specified size. Implementing this is optional, as default code will load and call [method _generate].
		//Returning an empty texture is an OK way to fail and let another generator take care.
		//Care must be taken because this function is always called from a thread (not the main thread).
		//[param metadata] dictionary can be modified to store file-specific metadata that can be used in [method EditorResourceTooltipPlugin._make_tooltip_for_path] (like image size, sample length etc.).
		GenerateFromPath(path string, size Vector2i.XY, metadata Dictionary.Any) objects.Texture2D
		//If this function returns [code]true[/code], the generator will automatically generate the small previews from the normal preview texture generated by the methods [method _generate] or [method _generate_from_path].
		//By default, it returns [code]false[/code].
		GenerateSmallPreviewAutomatically() bool
		//If this function returns [code]true[/code], the generator will call [method _generate] or [method _generate_from_path] for small previews as well.
		//By default, it returns [code]false[/code].
		CanGenerateSmallPreview() bool
	}
*/
type Instance [1]classdb.EditorResourcePreviewGenerator
type Any interface {
	gd.IsClass
	AsEditorResourcePreviewGenerator() Instance
}

/*
Returns [code]true[/code] if your generator supports the resource of type [param type].
*/
func (Instance) _handles(impl func(ptr unsafe.Pointer, atype string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var atype = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(atype)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, atype.String())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Generate a preview from a given resource with the specified size. This must always be implemented.
Returning an empty texture is an OK way to fail and let another generator take care.
Care must be taken because this function is always called from a thread (not the main thread).
[param metadata] dictionary can be modified to store file-specific metadata that can be used in [method EditorResourceTooltipPlugin._make_tooltip_for_path] (like image size, sample length etc.).
*/
func (Instance) _generate(impl func(ptr unsafe.Pointer, resource objects.Resource, size Vector2i.XY, metadata Dictionary.Any) objects.Texture2D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = objects.Resource{pointers.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(resource[0])
		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)
		var metadata = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args, 2))
		defer pointers.End(metadata)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource, size, metadata)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Generate a preview directly from a path with the specified size. Implementing this is optional, as default code will load and call [method _generate].
Returning an empty texture is an OK way to fail and let another generator take care.
Care must be taken because this function is always called from a thread (not the main thread).
[param metadata] dictionary can be modified to store file-specific metadata that can be used in [method EditorResourceTooltipPlugin._make_tooltip_for_path] (like image size, sample length etc.).
*/
func (Instance) _generate_from_path(impl func(ptr unsafe.Pointer, path string, size Vector2i.XY, metadata Dictionary.Any) objects.Texture2D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(path)
		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)
		var metadata = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args, 2))
		defer pointers.End(metadata)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String(), size, metadata)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
If this function returns [code]true[/code], the generator will automatically generate the small previews from the normal preview texture generated by the methods [method _generate] or [method _generate_from_path].
By default, it returns [code]false[/code].
*/
func (Instance) _generate_small_preview_automatically(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
If this function returns [code]true[/code], the generator will call [method _generate] or [method _generate_from_path] for small previews as well.
By default, it returns [code]false[/code].
*/
func (Instance) _can_generate_small_preview(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.EditorResourcePreviewGenerator

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorResourcePreviewGenerator"))
	return Instance{classdb.EditorResourcePreviewGenerator(object)}
}

/*
Returns [code]true[/code] if your generator supports the resource of type [param type].
*/
func (class) _handles(impl func(ptr unsafe.Pointer, atype gd.String) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var atype = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, atype)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Generate a preview from a given resource with the specified size. This must always be implemented.
Returning an empty texture is an OK way to fail and let another generator take care.
Care must be taken because this function is always called from a thread (not the main thread).
[param metadata] dictionary can be modified to store file-specific metadata that can be used in [method EditorResourceTooltipPlugin._make_tooltip_for_path] (like image size, sample length etc.).
*/
func (class) _generate(impl func(ptr unsafe.Pointer, resource objects.Resource, size gd.Vector2i, metadata gd.Dictionary) objects.Texture2D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var resource = objects.Resource{pointers.New[classdb.Resource]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(resource[0])
		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)
		var metadata = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args, 2))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource, size, metadata)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Generate a preview directly from a path with the specified size. Implementing this is optional, as default code will load and call [method _generate].
Returning an empty texture is an OK way to fail and let another generator take care.
Care must be taken because this function is always called from a thread (not the main thread).
[param metadata] dictionary can be modified to store file-specific metadata that can be used in [method EditorResourceTooltipPlugin._make_tooltip_for_path] (like image size, sample length etc.).
*/
func (class) _generate_from_path(impl func(ptr unsafe.Pointer, path gd.String, size gd.Vector2i, metadata gd.Dictionary) objects.Texture2D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var path = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		var size = gd.UnsafeGet[gd.Vector2i](p_args, 1)
		var metadata = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args, 2))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path, size, metadata)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
If this function returns [code]true[/code], the generator will automatically generate the small previews from the normal preview texture generated by the methods [method _generate] or [method _generate_from_path].
By default, it returns [code]false[/code].
*/
func (class) _generate_small_preview_automatically(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
If this function returns [code]true[/code], the generator will call [method _generate] or [method _generate_from_path] for small previews as well.
By default, it returns [code]false[/code].
*/
func (class) _can_generate_small_preview(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (self class) AsEditorResourcePreviewGenerator() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsEditorResourcePreviewGenerator() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_handles":
		return reflect.ValueOf(self._handles)
	case "_generate":
		return reflect.ValueOf(self._generate)
	case "_generate_from_path":
		return reflect.ValueOf(self._generate_from_path)
	case "_generate_small_preview_automatically":
		return reflect.ValueOf(self._generate_small_preview_automatically)
	case "_can_generate_small_preview":
		return reflect.ValueOf(self._can_generate_small_preview)
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_handles":
		return reflect.ValueOf(self._handles)
	case "_generate":
		return reflect.ValueOf(self._generate)
	case "_generate_from_path":
		return reflect.ValueOf(self._generate_from_path)
	case "_generate_small_preview_automatically":
		return reflect.ValueOf(self._generate_small_preview_automatically)
	case "_can_generate_small_preview":
		return reflect.ValueOf(self._can_generate_small_preview)
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("EditorResourcePreviewGenerator", func(ptr gd.Object) any { return classdb.EditorResourcePreviewGenerator(ptr) })
}
