// Package DirAccess provides methods for working with DirAccess object instances.
package DirAccess

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
This class is used to manage directories and their content, even outside of the project folder.
[DirAccess] can't be instantiated directly. Instead it is created with a static method that takes a path for which it will be opened.
Most of the methods have a static alternative that can be used without creating a [DirAccess]. Static methods only support absolute paths (including [code]res://[/code] and [code]user://[/code]).
[codeblock]
# Standard
var dir = DirAccess.open("user://levels")
dir.make_dir("world1")
# Static
DirAccess.make_dir_absolute("user://levels/world1")
[/codeblock]
[b]Note:[/b] Many resources types are imported (e.g. textures or sound files), and their source asset will not be included in the exported game, as only the imported version is used. Use [ResourceLoader] to access imported resources.
Here is an example on how to iterate through the files of a directory:
[codeblocks]
[gdscript]
func dir_contents(path):

	var dir = DirAccess.open(path)
	if dir:
	    dir.list_dir_begin()
	    var file_name = dir.get_next()
	    while file_name != "":
	        if dir.current_is_dir():
	            print("Found directory: " + file_name)
	        else:
	            print("Found file: " + file_name)
	        file_name = dir.get_next()
	else:
	    print("An error occurred when trying to access the path.")

[/gdscript]
[csharp]
public void DirContents(string path)

	{
	    using var dir = DirAccess.Open(path);
	    if (dir != null)
	    {
	        dir.ListDirBegin();
	        string fileName = dir.GetNext();
	        while (fileName != "")
	        {
	            if (dir.CurrentIsDir())
	            {
	                GD.Print($"Found directory: {fileName}");
	            }
	            else
	            {
	                GD.Print($"Found file: {fileName}");
	            }
	            fileName = dir.GetNext();
	        }
	    }
	    else
	    {
	        GD.Print("An error occurred when trying to access the path.");
	    }
	}

[/csharp]
[/codeblocks]
*/
type Instance [1]gdclass.DirAccess
type Any interface {
	gd.IsClass
	AsDirAccess() Instance
}

/*
Creates a new [DirAccess] object and opens an existing directory of the filesystem. The [param path] argument can be within the project tree ([code]res://folder[/code]), the user directory ([code]user://folder[/code]) or an absolute path of the user filesystem (e.g. [code]/tmp/folder[/code] or [code]C:\tmp\folder[/code]).
Returns [code]null[/code] if opening the directory failed. You can use [method get_open_error] to check the error that occurred.
*/
func Open(path string) [1]gdclass.DirAccess {
	self := Instance{}
	return [1]gdclass.DirAccess(class(self).Open(gd.NewString(path)))
}

/*
Returns the result of the last [method open] call in the current thread.
*/
func GetOpenError() error {
	self := Instance{}
	return error(class(self).GetOpenError())
}

/*
Initializes the stream used to list all files and directories using the [method get_next] function, closing the currently opened stream if needed. Once the stream has been processed, it should typically be closed with [method list_dir_end].
Affected by [member include_hidden] and [member include_navigational].
[b]Note:[/b] The order of files and directories returned by this method is not deterministic, and can vary between operating systems. If you want a list of all files or folders sorted alphabetically, use [method get_files] or [method get_directories].
*/
func (self Instance) ListDirBegin() error {
	return error(class(self).ListDirBegin())
}

/*
Returns the next element (file or directory) in the current directory.
The name of the file or directory is returned (and not its full path). Once the stream has been fully processed, the method returns an empty [String] and closes the stream automatically (i.e. [method list_dir_end] would not be mandatory in such a case).
*/
func (self Instance) GetNext() string {
	return string(class(self).GetNext().String())
}

/*
Returns whether the current item processed with the last [method get_next] call is a directory ([code].[/code] and [code]..[/code] are considered directories).
*/
func (self Instance) CurrentIsDir() bool {
	return bool(class(self).CurrentIsDir())
}

/*
Closes the current stream opened with [method list_dir_begin] (whether it has been fully processed with [method get_next] does not matter).
*/
func (self Instance) ListDirEnd() {
	class(self).ListDirEnd()
}

/*
Returns a [PackedStringArray] containing filenames of the directory contents, excluding directories. The array is sorted alphabetically.
Affected by [member include_hidden].
[b]Note:[/b] When used on a [code]res://[/code] path in an exported project, only the files actually included in the PCK at the given folder level are returned. In practice, this means that since imported resources are stored in a top-level [code].godot/[/code] folder, only paths to [code]*.gd[/code] and [code]*.import[/code] files are returned (plus a few files such as [code]project.godot[/code] or [code]project.binary[/code] and the project icon). In an exported project, the list of returned files will also vary depending on whether [member ProjectSettings.editor/export/convert_text_resources_to_binary] is [code]true[/code].
*/
func (self Instance) GetFiles() []string {
	return []string(class(self).GetFiles().Strings())
}

/*
Returns a [PackedStringArray] containing filenames of the directory contents, excluding directories, at the given [param path]. The array is sorted alphabetically.
Use [method get_files] if you want more control of what gets included.
*/
func GetFilesAt(path string) []string {
	self := Instance{}
	return []string(class(self).GetFilesAt(gd.NewString(path)).Strings())
}

/*
Returns a [PackedStringArray] containing filenames of the directory contents, excluding files. The array is sorted alphabetically.
Affected by [member include_hidden] and [member include_navigational].
*/
func (self Instance) GetDirectories() []string {
	return []string(class(self).GetDirectories().Strings())
}

/*
Returns a [PackedStringArray] containing filenames of the directory contents, excluding files, at the given [param path]. The array is sorted alphabetically.
Use [method get_directories] if you want more control of what gets included.
*/
func GetDirectoriesAt(path string) []string {
	self := Instance{}
	return []string(class(self).GetDirectoriesAt(gd.NewString(path)).Strings())
}

/*
On Windows, returns the number of drives (partitions) mounted on the current filesystem.
On macOS, returns the number of mounted volumes.
On Linux, returns the number of mounted volumes and GTK 3 bookmarks.
On other platforms, the method returns 0.
*/
func GetDriveCount() int {
	self := Instance{}
	return int(int(class(self).GetDriveCount()))
}

/*
On Windows, returns the name of the drive (partition) passed as an argument (e.g. [code]C:[/code]).
On macOS, returns the path to the mounted volume passed as an argument.
On Linux, returns the path to the mounted volume or GTK 3 bookmark passed as an argument.
On other platforms, or if the requested drive does not exist, the method returns an empty String.
*/
func GetDriveName(idx int) string {
	self := Instance{}
	return string(class(self).GetDriveName(gd.Int(idx)).String())
}

/*
Returns the currently opened directory's drive index. See [method get_drive_name] to convert returned index to the name of the drive.
*/
func (self Instance) GetCurrentDrive() int {
	return int(int(class(self).GetCurrentDrive()))
}

/*
Changes the currently opened directory to the one passed as an argument. The argument can be relative to the current directory (e.g. [code]newdir[/code] or [code]../newdir[/code]), or an absolute path (e.g. [code]/tmp/newdir[/code] or [code]res://somedir/newdir[/code]).
Returns one of the [enum Error] code constants ([constant OK] on success).
[b]Note:[/b] The new directory must be within the same scope, e.g. when you had opened a directory inside [code]res://[/code], you can't change it to [code]user://[/code] directory. If you need to open a directory in another access scope, use [method open] to create a new instance instead.
*/
func (self Instance) ChangeDir(to_dir string) error {
	return error(class(self).ChangeDir(gd.NewString(to_dir)))
}

/*
Returns the absolute path to the currently opened directory (e.g. [code]res://folder[/code] or [code]C:\tmp\folder[/code]).
*/
func (self Instance) GetCurrentDir() string {
	return string(class(self).GetCurrentDir(true).String())
}

/*
Creates a directory. The argument can be relative to the current directory, or an absolute path. The target directory should be placed in an already existing directory (to create the full path recursively, see [method make_dir_recursive]).
Returns one of the [enum Error] code constants ([constant OK] on success).
*/
func (self Instance) MakeDir(path string) error {
	return error(class(self).MakeDir(gd.NewString(path)))
}

/*
Static version of [method make_dir]. Supports only absolute paths.
*/
func MakeDirAbsolute(path string) error {
	self := Instance{}
	return error(class(self).MakeDirAbsolute(gd.NewString(path)))
}

/*
Creates a target directory and all necessary intermediate directories in its path, by calling [method make_dir] recursively. The argument can be relative to the current directory, or an absolute path.
Returns one of the [enum Error] code constants ([constant OK] on success).
*/
func (self Instance) MakeDirRecursive(path string) error {
	return error(class(self).MakeDirRecursive(gd.NewString(path)))
}

/*
Static version of [method make_dir_recursive]. Supports only absolute paths.
*/
func MakeDirRecursiveAbsolute(path string) error {
	self := Instance{}
	return error(class(self).MakeDirRecursiveAbsolute(gd.NewString(path)))
}

/*
Returns whether the target file exists. The argument can be relative to the current directory, or an absolute path.
For a static equivalent, use [method FileAccess.file_exists].
*/
func (self Instance) FileExists(path string) bool {
	return bool(class(self).FileExists(gd.NewString(path)))
}

/*
Returns whether the target directory exists. The argument can be relative to the current directory, or an absolute path.
*/
func (self Instance) DirExists(path string) bool {
	return bool(class(self).DirExists(gd.NewString(path)))
}

/*
Static version of [method dir_exists]. Supports only absolute paths.
*/
func DirExistsAbsolute(path string) bool {
	self := Instance{}
	return bool(class(self).DirExistsAbsolute(gd.NewString(path)))
}

/*
Returns the available space on the current directory's disk, in bytes. Returns [code]0[/code] if the platform-specific method to query the available space fails.
*/
func (self Instance) GetSpaceLeft() int {
	return int(int(class(self).GetSpaceLeft()))
}

/*
Copies the [param from] file to the [param to] destination. Both arguments should be paths to files, either relative or absolute. If the destination file exists and is not access-protected, it will be overwritten.
If [param chmod_flags] is different than [code]-1[/code], the Unix permissions for the destination path will be set to the provided value, if available on the current operating system.
Returns one of the [enum Error] code constants ([constant OK] on success).
*/
func (self Instance) Copy(from string, to string) error {
	return error(class(self).Copy(gd.NewString(from), gd.NewString(to), gd.Int(-1)))
}

/*
Static version of [method copy]. Supports only absolute paths.
*/
func CopyAbsolute(from string, to string) error {
	self := Instance{}
	return error(class(self).CopyAbsolute(gd.NewString(from), gd.NewString(to), gd.Int(-1)))
}

/*
Renames (move) the [param from] file or directory to the [param to] destination. Both arguments should be paths to files or directories, either relative or absolute. If the destination file or directory exists and is not access-protected, it will be overwritten.
Returns one of the [enum Error] code constants ([constant OK] on success).
*/
func (self Instance) Rename(from string, to string) error {
	return error(class(self).Rename(gd.NewString(from), gd.NewString(to)))
}

/*
Static version of [method rename]. Supports only absolute paths.
*/
func RenameAbsolute(from string, to string) error {
	self := Instance{}
	return error(class(self).RenameAbsolute(gd.NewString(from), gd.NewString(to)))
}

/*
Permanently deletes the target file or an empty directory. The argument can be relative to the current directory, or an absolute path. If the target directory is not empty, the operation will fail.
If you don't want to delete the file/directory permanently, use [method OS.move_to_trash] instead.
Returns one of the [enum Error] code constants ([constant OK] on success).
*/
func (self Instance) Remove(path string) error {
	return error(class(self).Remove(gd.NewString(path)))
}

/*
Static version of [method remove]. Supports only absolute paths.
*/
func RemoveAbsolute(path string) error {
	self := Instance{}
	return error(class(self).RemoveAbsolute(gd.NewString(path)))
}

/*
Returns [code]true[/code] if the file or directory is a symbolic link, directory junction, or other reparse point.
[b]Note:[/b] This method is implemented on macOS, Linux, and Windows.
*/
func (self Instance) IsLink(path string) bool {
	return bool(class(self).IsLink(gd.NewString(path)))
}

/*
Returns target of the symbolic link.
[b]Note:[/b] This method is implemented on macOS, Linux, and Windows.
*/
func (self Instance) ReadLink(path string) string {
	return string(class(self).ReadLink(gd.NewString(path)).String())
}

/*
Creates symbolic link between files or folders.
[b]Note:[/b] On Windows, this method works only if the application is running with elevated privileges or Developer Mode is enabled.
[b]Note:[/b] This method is implemented on macOS, Linux, and Windows.
*/
func (self Instance) CreateLink(source string, target string) error {
	return error(class(self).CreateLink(gd.NewString(source), gd.NewString(target)))
}

/*
Returns [code]true[/code] if the file system or directory use case sensitive file names.
[b]Note:[/b] This method is implemented on macOS, Linux (for EXT4 and F2FS filesystems only) and Windows. On other platforms, it always returns [code]true[/code].
*/
func (self Instance) IsCaseSensitive(path string) bool {
	return bool(class(self).IsCaseSensitive(gd.NewString(path)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.DirAccess

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("DirAccess"))
	return Instance{*(*gdclass.DirAccess)(unsafe.Pointer(&object))}
}

func (self Instance) IncludeNavigational() bool {
	return bool(class(self).GetIncludeNavigational())
}

func (self Instance) SetIncludeNavigational(value bool) {
	class(self).SetIncludeNavigational(value)
}

func (self Instance) IncludeHidden() bool {
	return bool(class(self).GetIncludeHidden())
}

func (self Instance) SetIncludeHidden(value bool) {
	class(self).SetIncludeHidden(value)
}

/*
Creates a new [DirAccess] object and opens an existing directory of the filesystem. The [param path] argument can be within the project tree ([code]res://folder[/code]), the user directory ([code]user://folder[/code]) or an absolute path of the user filesystem (e.g. [code]/tmp/folder[/code] or [code]C:\tmp\folder[/code]).
Returns [code]null[/code] if opening the directory failed. You can use [method get_open_error] to check the error that occurred.
*/
//go:nosplit
func (self class) Open(path gd.String) [1]gdclass.DirAccess {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_open, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.DirAccess{gd.PointerWithOwnershipTransferredToGo[gdclass.DirAccess](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the result of the last [method open] call in the current thread.
*/
//go:nosplit
func (self class) GetOpenError() error {
	var frame = callframe.New()
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_get_open_error, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Initializes the stream used to list all files and directories using the [method get_next] function, closing the currently opened stream if needed. Once the stream has been processed, it should typically be closed with [method list_dir_end].
Affected by [member include_hidden] and [member include_navigational].
[b]Note:[/b] The order of files and directories returned by this method is not deterministic, and can vary between operating systems. If you want a list of all files or folders sorted alphabetically, use [method get_files] or [method get_directories].
*/
//go:nosplit
func (self class) ListDirBegin() error {
	var frame = callframe.New()
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_list_dir_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the next element (file or directory) in the current directory.
The name of the file or directory is returned (and not its full path). Once the stream has been fully processed, the method returns an empty [String] and closes the stream automatically (i.e. [method list_dir_end] would not be mandatory in such a case).
*/
//go:nosplit
func (self class) GetNext() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_get_next, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns whether the current item processed with the last [method get_next] call is a directory ([code].[/code] and [code]..[/code] are considered directories).
*/
//go:nosplit
func (self class) CurrentIsDir() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_current_is_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Closes the current stream opened with [method list_dir_begin] (whether it has been fully processed with [method get_next] does not matter).
*/
//go:nosplit
func (self class) ListDirEnd() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_list_dir_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns a [PackedStringArray] containing filenames of the directory contents, excluding directories. The array is sorted alphabetically.
Affected by [member include_hidden].
[b]Note:[/b] When used on a [code]res://[/code] path in an exported project, only the files actually included in the PCK at the given folder level are returned. In practice, this means that since imported resources are stored in a top-level [code].godot/[/code] folder, only paths to [code]*.gd[/code] and [code]*.import[/code] files are returned (plus a few files such as [code]project.godot[/code] or [code]project.binary[/code] and the project icon). In an exported project, the list of returned files will also vary depending on whether [member ProjectSettings.editor/export/convert_text_resources_to_binary] is [code]true[/code].
*/
//go:nosplit
func (self class) GetFiles() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_get_files, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a [PackedStringArray] containing filenames of the directory contents, excluding directories, at the given [param path]. The array is sorted alphabetically.
Use [method get_files] if you want more control of what gets included.
*/
//go:nosplit
func (self class) GetFilesAt(path gd.String) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_get_files_at, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a [PackedStringArray] containing filenames of the directory contents, excluding files. The array is sorted alphabetically.
Affected by [member include_hidden] and [member include_navigational].
*/
//go:nosplit
func (self class) GetDirectories() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_get_directories, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a [PackedStringArray] containing filenames of the directory contents, excluding files, at the given [param path]. The array is sorted alphabetically.
Use [method get_directories] if you want more control of what gets included.
*/
//go:nosplit
func (self class) GetDirectoriesAt(path gd.String) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_get_directories_at, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
On Windows, returns the number of drives (partitions) mounted on the current filesystem.
On macOS, returns the number of mounted volumes.
On Linux, returns the number of mounted volumes and GTK 3 bookmarks.
On other platforms, the method returns 0.
*/
//go:nosplit
func (self class) GetDriveCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_get_drive_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
On Windows, returns the name of the drive (partition) passed as an argument (e.g. [code]C:[/code]).
On macOS, returns the path to the mounted volume passed as an argument.
On Linux, returns the path to the mounted volume or GTK 3 bookmark passed as an argument.
On other platforms, or if the requested drive does not exist, the method returns an empty String.
*/
//go:nosplit
func (self class) GetDriveName(idx gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_get_drive_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the currently opened directory's drive index. See [method get_drive_name] to convert returned index to the name of the drive.
*/
//go:nosplit
func (self class) GetCurrentDrive() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_get_current_drive, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Changes the currently opened directory to the one passed as an argument. The argument can be relative to the current directory (e.g. [code]newdir[/code] or [code]../newdir[/code]), or an absolute path (e.g. [code]/tmp/newdir[/code] or [code]res://somedir/newdir[/code]).
Returns one of the [enum Error] code constants ([constant OK] on success).
[b]Note:[/b] The new directory must be within the same scope, e.g. when you had opened a directory inside [code]res://[/code], you can't change it to [code]user://[/code] directory. If you need to open a directory in another access scope, use [method open] to create a new instance instead.
*/
//go:nosplit
func (self class) ChangeDir(to_dir gd.String) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(to_dir))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_change_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the absolute path to the currently opened directory (e.g. [code]res://folder[/code] or [code]C:\tmp\folder[/code]).
*/
//go:nosplit
func (self class) GetCurrentDir(include_drive bool) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, include_drive)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_get_current_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Creates a directory. The argument can be relative to the current directory, or an absolute path. The target directory should be placed in an already existing directory (to create the full path recursively, see [method make_dir_recursive]).
Returns one of the [enum Error] code constants ([constant OK] on success).
*/
//go:nosplit
func (self class) MakeDir(path gd.String) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_make_dir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Static version of [method make_dir]. Supports only absolute paths.
*/
//go:nosplit
func (self class) MakeDirAbsolute(path gd.String) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_make_dir_absolute, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates a target directory and all necessary intermediate directories in its path, by calling [method make_dir] recursively. The argument can be relative to the current directory, or an absolute path.
Returns one of the [enum Error] code constants ([constant OK] on success).
*/
//go:nosplit
func (self class) MakeDirRecursive(path gd.String) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_make_dir_recursive, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Static version of [method make_dir_recursive]. Supports only absolute paths.
*/
//go:nosplit
func (self class) MakeDirRecursiveAbsolute(path gd.String) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_make_dir_recursive_absolute, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns whether the target file exists. The argument can be relative to the current directory, or an absolute path.
For a static equivalent, use [method FileAccess.file_exists].
*/
//go:nosplit
func (self class) FileExists(path gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_file_exists, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns whether the target directory exists. The argument can be relative to the current directory, or an absolute path.
*/
//go:nosplit
func (self class) DirExists(path gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_dir_exists, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Static version of [method dir_exists]. Supports only absolute paths.
*/
//go:nosplit
func (self class) DirExistsAbsolute(path gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_dir_exists_absolute, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the available space on the current directory's disk, in bytes. Returns [code]0[/code] if the platform-specific method to query the available space fails.
*/
//go:nosplit
func (self class) GetSpaceLeft() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_get_space_left, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Copies the [param from] file to the [param to] destination. Both arguments should be paths to files, either relative or absolute. If the destination file exists and is not access-protected, it will be overwritten.
If [param chmod_flags] is different than [code]-1[/code], the Unix permissions for the destination path will be set to the provided value, if available on the current operating system.
Returns one of the [enum Error] code constants ([constant OK] on success).
*/
//go:nosplit
func (self class) Copy(from gd.String, to gd.String, chmod_flags gd.Int) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(from))
	callframe.Arg(frame, pointers.Get(to))
	callframe.Arg(frame, chmod_flags)
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_copy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Static version of [method copy]. Supports only absolute paths.
*/
//go:nosplit
func (self class) CopyAbsolute(from gd.String, to gd.String, chmod_flags gd.Int) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(from))
	callframe.Arg(frame, pointers.Get(to))
	callframe.Arg(frame, chmod_flags)
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_copy_absolute, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Renames (move) the [param from] file or directory to the [param to] destination. Both arguments should be paths to files or directories, either relative or absolute. If the destination file or directory exists and is not access-protected, it will be overwritten.
Returns one of the [enum Error] code constants ([constant OK] on success).
*/
//go:nosplit
func (self class) Rename(from gd.String, to gd.String) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(from))
	callframe.Arg(frame, pointers.Get(to))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_rename, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Static version of [method rename]. Supports only absolute paths.
*/
//go:nosplit
func (self class) RenameAbsolute(from gd.String, to gd.String) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(from))
	callframe.Arg(frame, pointers.Get(to))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_rename_absolute, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Permanently deletes the target file or an empty directory. The argument can be relative to the current directory, or an absolute path. If the target directory is not empty, the operation will fail.
If you don't want to delete the file/directory permanently, use [method OS.move_to_trash] instead.
Returns one of the [enum Error] code constants ([constant OK] on success).
*/
//go:nosplit
func (self class) Remove(path gd.String) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_remove, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Static version of [method remove]. Supports only absolute paths.
*/
//go:nosplit
func (self class) RemoveAbsolute(path gd.String) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_remove_absolute, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the file or directory is a symbolic link, directory junction, or other reparse point.
[b]Note:[/b] This method is implemented on macOS, Linux, and Windows.
*/
//go:nosplit
func (self class) IsLink(path gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_is_link, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns target of the symbolic link.
[b]Note:[/b] This method is implemented on macOS, Linux, and Windows.
*/
//go:nosplit
func (self class) ReadLink(path gd.String) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_read_link, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Creates symbolic link between files or folders.
[b]Note:[/b] On Windows, this method works only if the application is running with elevated privileges or Developer Mode is enabled.
[b]Note:[/b] This method is implemented on macOS, Linux, and Windows.
*/
//go:nosplit
func (self class) CreateLink(source gd.String, target gd.String) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(source))
	callframe.Arg(frame, pointers.Get(target))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_create_link, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetIncludeNavigational(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_set_include_navigational, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetIncludeNavigational() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_get_include_navigational, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetIncludeHidden(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_set_include_hidden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetIncludeHidden() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_get_include_hidden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the file system or directory use case sensitive file names.
[b]Note:[/b] This method is implemented on macOS, Linux (for EXT4 and F2FS filesystems only) and Windows. On other platforms, it always returns [code]true[/code].
*/
//go:nosplit
func (self class) IsCaseSensitive(path gd.String) bool {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DirAccess.Bind_is_case_sensitive, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsDirAccess() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsDirAccess() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("DirAccess", func(ptr gd.Object) any { return [1]gdclass.DirAccess{*(*gdclass.DirAccess)(unsafe.Pointer(&ptr))} })
}

type Error int

const (
	/*Methods that return [enum Error] return [constant OK] when no error occurred.
	  Since [constant OK] has value 0, and all other error constants are positive integers, it can also be used in boolean checks.
	  [b]Example:[/b]
	  [codeblock]
	  var error = method_that_returns_error()
	  if error != OK:
	      printerr("Failure!")

	  # Or, alternatively:
	  if error:
	      printerr("Still failing!")
	  [/codeblock]
	  [b]Note:[/b] Many functions do not return an error code, but will print error messages to standard output.*/
	Ok Error = 0
	/*Generic error.*/
	Failed Error = 1
	/*Unavailable error.*/
	ErrUnavailable Error = 2
	/*Unconfigured error.*/
	ErrUnconfigured Error = 3
	/*Unauthorized error.*/
	ErrUnauthorized Error = 4
	/*Parameter range error.*/
	ErrParameterRangeError Error = 5
	/*Out of memory (OOM) error.*/
	ErrOutOfMemory Error = 6
	/*File: Not found error.*/
	ErrFileNotFound Error = 7
	/*File: Bad drive error.*/
	ErrFileBadDrive Error = 8
	/*File: Bad path error.*/
	ErrFileBadPath Error = 9
	/*File: No permission error.*/
	ErrFileNoPermission Error = 10
	/*File: Already in use error.*/
	ErrFileAlreadyInUse Error = 11
	/*File: Can't open error.*/
	ErrFileCantOpen Error = 12
	/*File: Can't write error.*/
	ErrFileCantWrite Error = 13
	/*File: Can't read error.*/
	ErrFileCantRead Error = 14
	/*File: Unrecognized error.*/
	ErrFileUnrecognized Error = 15
	/*File: Corrupt error.*/
	ErrFileCorrupt Error = 16
	/*File: Missing dependencies error.*/
	ErrFileMissingDependencies Error = 17
	/*File: End of file (EOF) error.*/
	ErrFileEof Error = 18
	/*Can't open error.*/
	ErrCantOpen Error = 19
	/*Can't create error.*/
	ErrCantCreate Error = 20
	/*Query failed error.*/
	ErrQueryFailed Error = 21
	/*Already in use error.*/
	ErrAlreadyInUse Error = 22
	/*Locked error.*/
	ErrLocked Error = 23
	/*Timeout error.*/
	ErrTimeout Error = 24
	/*Can't connect error.*/
	ErrCantConnect Error = 25
	/*Can't resolve error.*/
	ErrCantResolve Error = 26
	/*Connection error.*/
	ErrConnectionError Error = 27
	/*Can't acquire resource error.*/
	ErrCantAcquireResource Error = 28
	/*Can't fork process error.*/
	ErrCantFork Error = 29
	/*Invalid data error.*/
	ErrInvalidData Error = 30
	/*Invalid parameter error.*/
	ErrInvalidParameter Error = 31
	/*Already exists error.*/
	ErrAlreadyExists Error = 32
	/*Does not exist error.*/
	ErrDoesNotExist Error = 33
	/*Database: Read error.*/
	ErrDatabaseCantRead Error = 34
	/*Database: Write error.*/
	ErrDatabaseCantWrite Error = 35
	/*Compilation failed error.*/
	ErrCompilationFailed Error = 36
	/*Method not found error.*/
	ErrMethodNotFound Error = 37
	/*Linking failed error.*/
	ErrLinkFailed Error = 38
	/*Script failed error.*/
	ErrScriptFailed Error = 39
	/*Cycling link (import cycle) error.*/
	ErrCyclicLink Error = 40
	/*Invalid declaration error.*/
	ErrInvalidDeclaration Error = 41
	/*Duplicate symbol error.*/
	ErrDuplicateSymbol Error = 42
	/*Parse error.*/
	ErrParseError Error = 43
	/*Busy error.*/
	ErrBusy Error = 44
	/*Skip error.*/
	ErrSkip Error = 45
	/*Help error. Used internally when passing [code]--version[/code] or [code]--help[/code] as executable options.*/
	ErrHelp Error = 46
	/*Bug error, caused by an implementation issue in the method.
	  [b]Note:[/b] If a built-in method returns this code, please open an issue on [url=https://github.com/godotengine/godot/issues]the GitHub Issue Tracker[/url].*/
	ErrBug Error = 47
	/*Printer on fire error (This is an easter egg, no built-in methods return this error code).*/
	ErrPrinterOnFire Error = 48
)
