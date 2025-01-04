// Package EditorFileSystemImportFormatSupportQuery provides methods for working with EditorFileSystemImportFormatSupportQuery object instances.
package EditorFileSystemImportFormatSupportQuery

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"

var _ Object.ID
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This class is used to query and configure a certain import format. It is used in conjunction with asset format import plugins.

	// EditorFileSystemImportFormatSupportQuery methods that can be overridden by a [Class] that extends it.
	type EditorFileSystemImportFormatSupportQuery interface {
		//Return whether this importer is active.
		IsActive() bool
		//Return the file extensions supported.
		GetFileExtensions() []string
		//Query support. Return false if import must not continue.
		Query() bool
	}
*/
type Instance [1]gdclass.EditorFileSystemImportFormatSupportQuery
type Any interface {
	gd.IsClass
	AsEditorFileSystemImportFormatSupportQuery() Instance
}

/*
Return whether this importer is active.
*/
func (Instance) _is_active(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Return the file extensions supported.
*/
func (Instance) _get_file_extensions(impl func(ptr unsafe.Pointer) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewPackedStringSlice(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Query support. Return false if import must not continue.
*/
func (Instance) _query(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorFileSystemImportFormatSupportQuery

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorFileSystemImportFormatSupportQuery"))
	return Instance{*(*gdclass.EditorFileSystemImportFormatSupportQuery)(unsafe.Pointer(&object))}
}

/*
Return whether this importer is active.
*/
func (class) _is_active(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Return the file extensions supported.
*/
func (class) _get_file_extensions(impl func(ptr unsafe.Pointer) gd.PackedStringArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Query support. Return false if import must not continue.
*/
func (class) _query(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (self class) AsEditorFileSystemImportFormatSupportQuery() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsEditorFileSystemImportFormatSupportQuery() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_is_active":
		return reflect.ValueOf(self._is_active)
	case "_get_file_extensions":
		return reflect.ValueOf(self._get_file_extensions)
	case "_query":
		return reflect.ValueOf(self._query)
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_is_active":
		return reflect.ValueOf(self._is_active)
	case "_get_file_extensions":
		return reflect.ValueOf(self._get_file_extensions)
	case "_query":
		return reflect.ValueOf(self._query)
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	gdclass.Register("EditorFileSystemImportFormatSupportQuery", func(ptr gd.Object) any {
		return [1]gdclass.EditorFileSystemImportFormatSupportQuery{*(*gdclass.EditorFileSystemImportFormatSupportQuery)(unsafe.Pointer(&ptr))}
	})
}
