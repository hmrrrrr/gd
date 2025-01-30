// Package EditorFileSystem provides methods for working with EditorFileSystem object instances.
package EditorFileSystem

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Node"
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
This object holds information of all resources in the filesystem, their types, etc.
[b]Note:[/b] This class shouldn't be instantiated directly. Instead, access the singleton using [method EditorInterface.get_resource_filesystem].
*/
type Instance [1]gdclass.EditorFileSystem

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorFileSystem() Instance
}

/*
Gets the root directory object.
*/
func (self Instance) GetFilesystem() [1]gdclass.EditorFileSystemDirectory { //gd:EditorFileSystem.get_filesystem
	return [1]gdclass.EditorFileSystemDirectory(class(self).GetFilesystem())
}

/*
Returns [code]true[/code] if the filesystem is being scanned.
*/
func (self Instance) IsScanning() bool { //gd:EditorFileSystem.is_scanning
	return bool(class(self).IsScanning())
}

/*
Returns the scan progress for 0 to 1 if the FS is being scanned.
*/
func (self Instance) GetScanningProgress() Float.X { //gd:EditorFileSystem.get_scanning_progress
	return Float.X(Float.X(class(self).GetScanningProgress()))
}

/*
Scan the filesystem for changes.
*/
func (self Instance) Scan() { //gd:EditorFileSystem.scan
	class(self).Scan()
}

/*
Check if the source of any imported resource changed.
*/
func (self Instance) ScanSources() { //gd:EditorFileSystem.scan_sources
	class(self).ScanSources()
}

/*
Add a file in an existing directory, or schedule file information to be updated on editor restart. Can be used to update text files saved by an external program.
This will not import the file. To reimport, call [method reimport_files] or [method scan] methods.
*/
func (self Instance) UpdateFile(path string) { //gd:EditorFileSystem.update_file
	class(self).UpdateFile(String.New(path))
}

/*
Returns a view into the filesystem at [param path].
*/
func (self Instance) GetFilesystemPath(path string) [1]gdclass.EditorFileSystemDirectory { //gd:EditorFileSystem.get_filesystem_path
	return [1]gdclass.EditorFileSystemDirectory(class(self).GetFilesystemPath(String.New(path)))
}

/*
Returns the resource type of the file, given the full path. This returns a string such as [code]"Resource"[/code] or [code]"GDScript"[/code], [i]not[/i] a file extension such as [code]".gd"[/code].
*/
func (self Instance) GetFileType(path string) string { //gd:EditorFileSystem.get_file_type
	return string(class(self).GetFileType(String.New(path)).String())
}

/*
Reimports a set of files. Call this if these files or their [code].import[/code] files were directly edited by script or an external program.
If the file type changed or the file was newly created, use [method update_file] or [method scan].
[b]Note:[/b] This function blocks until the import is finished. However, the main loop iteration, including timers and [method Node._process], will occur during the import process due to progress bar updates. Avoid calls to [method reimport_files] or [method scan] while an import is in progress.
*/
func (self Instance) ReimportFiles(files []string) { //gd:EditorFileSystem.reimport_files
	class(self).ReimportFiles(Packed.MakeStrings(files...))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorFileSystem

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorFileSystem"))
	casted := Instance{*(*gdclass.EditorFileSystem)(unsafe.Pointer(&object))}
	return casted
}

/*
Gets the root directory object.
*/
//go:nosplit
func (self class) GetFilesystem() [1]gdclass.EditorFileSystemDirectory { //gd:EditorFileSystem.get_filesystem
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystem.Bind_get_filesystem, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.EditorFileSystemDirectory{gd.PointerLifetimeBoundTo[gdclass.EditorFileSystemDirectory](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the filesystem is being scanned.
*/
//go:nosplit
func (self class) IsScanning() bool { //gd:EditorFileSystem.is_scanning
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystem.Bind_is_scanning, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the scan progress for 0 to 1 if the FS is being scanned.
*/
//go:nosplit
func (self class) GetScanningProgress() float64 { //gd:EditorFileSystem.get_scanning_progress
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystem.Bind_get_scanning_progress, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Scan the filesystem for changes.
*/
//go:nosplit
func (self class) Scan() { //gd:EditorFileSystem.scan
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystem.Bind_scan, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Check if the source of any imported resource changed.
*/
//go:nosplit
func (self class) ScanSources() { //gd:EditorFileSystem.scan_sources
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystem.Bind_scan_sources, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Add a file in an existing directory, or schedule file information to be updated on editor restart. Can be used to update text files saved by an external program.
This will not import the file. To reimport, call [method reimport_files] or [method scan] methods.
*/
//go:nosplit
func (self class) UpdateFile(path String.Readable) { //gd:EditorFileSystem.update_file
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystem.Bind_update_file, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a view into the filesystem at [param path].
*/
//go:nosplit
func (self class) GetFilesystemPath(path String.Readable) [1]gdclass.EditorFileSystemDirectory { //gd:EditorFileSystem.get_filesystem_path
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystem.Bind_get_filesystem_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.EditorFileSystemDirectory{gd.PointerLifetimeBoundTo[gdclass.EditorFileSystemDirectory](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the resource type of the file, given the full path. This returns a string such as [code]"Resource"[/code] or [code]"GDScript"[/code], [i]not[/i] a file extension such as [code]".gd"[/code].
*/
//go:nosplit
func (self class) GetFileType(path String.Readable) String.Readable { //gd:EditorFileSystem.get_file_type
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystem.Bind_get_file_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Reimports a set of files. Call this if these files or their [code].import[/code] files were directly edited by script or an external program.
If the file type changed or the file was newly created, use [method update_file] or [method scan].
[b]Note:[/b] This function blocks until the import is finished. However, the main loop iteration, including timers and [method Node._process], will occur during the import process due to progress bar updates. Avoid calls to [method reimport_files] or [method scan] while an import is in progress.
*/
//go:nosplit
func (self class) ReimportFiles(files Packed.Strings) { //gd:EditorFileSystem.reimport_files
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPackedStrings(files)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorFileSystem.Bind_reimport_files, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self Instance) OnFilesystemChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("filesystem_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnScriptClassesUpdated(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("script_classes_updated"), gd.NewCallable(cb), 0)
}

func (self Instance) OnSourcesChanged(cb func(exist bool)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("sources_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnResourcesReimporting(cb func(resources []string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("resources_reimporting"), gd.NewCallable(cb), 0)
}

func (self Instance) OnResourcesReimported(cb func(resources []string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("resources_reimported"), gd.NewCallable(cb), 0)
}

func (self Instance) OnResourcesReload(cb func(resources []string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("resources_reload"), gd.NewCallable(cb), 0)
}

func (self class) AsEditorFileSystem() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorFileSystem() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced           { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance        { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node.Advanced(self.AsNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node.Instance(self.AsNode()), name)
	}
}
func init() {
	gdclass.Register("EditorFileSystem", func(ptr gd.Object) any {
		return [1]gdclass.EditorFileSystem{*(*gdclass.EditorFileSystem)(unsafe.Pointer(&ptr))}
	})
}
