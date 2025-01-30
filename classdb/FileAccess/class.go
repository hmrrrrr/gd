// Package FileAccess provides methods for working with FileAccess object instances.
package FileAccess

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
This class can be used to permanently store data in the user device's file system and to read from it. This is useful for store game save data or player configuration files.
Here's a sample on how to write and read from a file:
[codeblocks]
[gdscript]
func save_to_file(content):

	var file = FileAccess.open("user://save_game.dat", FileAccess.WRITE)
	file.store_string(content)

func load_from_file():

	var file = FileAccess.open("user://save_game.dat", FileAccess.READ)
	var content = file.get_as_text()
	return content

[/gdscript]
[csharp]
public void SaveToFile(string content)

	{
	    using var file = FileAccess.Open("user://save_game.dat", FileAccess.ModeFlags.Write);
	    file.StoreString(content);
	}

public string LoadFromFile()

	{
	    using var file = FileAccess.Open("user://save_game.dat", FileAccess.ModeFlags.Read);
	    string content = file.GetAsText();
	    return content;
	}

[/csharp]
[/codeblocks]
In the example above, the file will be saved in the user data folder as specified in the [url=$DOCS_URL/tutorials/io/data_paths.html]Data paths[/url] documentation.
[FileAccess] will close when it's freed, which happens when it goes out of scope or when it gets assigned with [code]null[/code]. [method close] can be used to close it before then explicitly. In C# the reference must be disposed manually, which can be done with the [code]using[/code] statement or by calling the [code]Dispose[/code] method directly.
[b]Note:[/b] To access project resources once exported, it is recommended to use [ResourceLoader] instead of [FileAccess], as some files are converted to engine-specific formats and their original source files might not be present in the exported PCK package.
[b]Note:[/b] Files are automatically closed only if the process exits "normally" (such as by clicking the window manager's close button or pressing [b]Alt + F4[/b]). If you stop the project execution by pressing [b]F8[/b] while the project is running, the file won't be closed as the game process will be killed. You can work around this by calling [method flush] at regular intervals.
*/
type Instance [1]gdclass.FileAccess

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsFileAccess() Instance
}

/*
Creates a new [FileAccess] object and opens the file for writing or reading, depending on the flags.
Returns [code]null[/code] if opening the file failed. You can use [method get_open_error] to check the error that occurred.
*/
func Open(path string, flags gdclass.FileAccessModeFlags) [1]gdclass.FileAccess { //gd:FileAccess.open
	self := Instance{}
	return [1]gdclass.FileAccess(class(self).Open(String.New(path), flags))
}

/*
Creates a new [FileAccess] object and opens an encrypted file in write or read mode. You need to pass a binary key to encrypt/decrypt it.
[b]Note:[/b] The provided key must be 32 bytes long.
Returns [code]null[/code] if opening the file failed. You can use [method get_open_error] to check the error that occurred.
*/
func OpenEncrypted(path string, mode_flags gdclass.FileAccessModeFlags, key []byte) [1]gdclass.FileAccess { //gd:FileAccess.open_encrypted
	self := Instance{}
	return [1]gdclass.FileAccess(class(self).OpenEncrypted(String.New(path), mode_flags, Packed.Bytes(Packed.New(key...))))
}

/*
Creates a new [FileAccess] object and opens an encrypted file in write or read mode. You need to pass a password to encrypt/decrypt it.
Returns [code]null[/code] if opening the file failed. You can use [method get_open_error] to check the error that occurred.
*/
func OpenEncryptedWithPass(path string, mode_flags gdclass.FileAccessModeFlags, pass string) [1]gdclass.FileAccess { //gd:FileAccess.open_encrypted_with_pass
	self := Instance{}
	return [1]gdclass.FileAccess(class(self).OpenEncryptedWithPass(String.New(path), mode_flags, String.New(pass)))
}

/*
Creates a new [FileAccess] object and opens a compressed file for reading or writing.
[b]Note:[/b] [method open_compressed] can only read files that were saved by Godot, not third-party compression formats. See [url=https://github.com/godotengine/godot/issues/28999]GitHub issue #28999[/url] for a workaround.
Returns [code]null[/code] if opening the file failed. You can use [method get_open_error] to check the error that occurred.
*/
func OpenCompressed(path string, mode_flags gdclass.FileAccessModeFlags) [1]gdclass.FileAccess { //gd:FileAccess.open_compressed
	self := Instance{}
	return [1]gdclass.FileAccess(class(self).OpenCompressed(String.New(path), mode_flags, 0))
}

/*
Returns the result of the last [method open] call in the current thread.
*/
func GetOpenError() error { //gd:FileAccess.get_open_error
	self := Instance{}
	return error(gd.ToError(class(self).GetOpenError()))
}

/*
Returns the whole [param path] file contents as a [PackedByteArray] without any decoding.
Returns an empty [PackedByteArray] if an error occurred while opening the file. You can use [method get_open_error] to check the error that occurred.
*/
func GetFileAsBytes(path string) []byte { //gd:FileAccess.get_file_as_bytes
	self := Instance{}
	return []byte(class(self).GetFileAsBytes(String.New(path)).Bytes())
}

/*
Returns the whole [param path] file contents as a [String]. Text is interpreted as being UTF-8 encoded.
Returns an empty [String] if an error occurred while opening the file. You can use [method get_open_error] to check the error that occurred.
*/
func GetFileAsString(path string) string { //gd:FileAccess.get_file_as_string
	self := Instance{}
	return string(class(self).GetFileAsString(String.New(path)).String())
}

/*
Resizes the file to a specified length. The file must be open in a mode that permits writing. If the file is extended, NUL characters are appended. If the file is truncated, all data from the end file to the original length of the file is lost.
*/
func (self Instance) Resize(length int) error { //gd:FileAccess.resize
	return error(gd.ToError(class(self).Resize(int64(length))))
}

/*
Writes the file's buffer to disk. Flushing is automatically performed when the file is closed. This means you don't need to call [method flush] manually before closing a file. Still, calling [method flush] can be used to ensure the data is safe even if the project crashes instead of being closed gracefully.
[b]Note:[/b] Only call [method flush] when you actually need it. Otherwise, it will decrease performance due to constant disk writes.
*/
func (self Instance) Flush() { //gd:FileAccess.flush
	class(self).Flush()
}

/*
Returns the path as a [String] for the current open file.
*/
func (self Instance) GetPath() string { //gd:FileAccess.get_path
	return string(class(self).GetPath().String())
}

/*
Returns the absolute path as a [String] for the current open file.
*/
func (self Instance) GetPathAbsolute() string { //gd:FileAccess.get_path_absolute
	return string(class(self).GetPathAbsolute().String())
}

/*
Returns [code]true[/code] if the file is currently opened.
*/
func (self Instance) IsOpen() bool { //gd:FileAccess.is_open
	return bool(class(self).IsOpen())
}

/*
Changes the file reading/writing cursor to the specified position (in bytes from the beginning of the file).
*/
func (self Instance) SeekTo(position int) { //gd:FileAccess.seek
	class(self).SeekTo(int64(position))
}

/*
Changes the file reading/writing cursor to the specified position (in bytes from the end of the file).
[b]Note:[/b] This is an offset, so you should use negative numbers or the cursor will be at the end of the file.
*/
func (self Instance) SeekEnd() { //gd:FileAccess.seek_end
	class(self).SeekEnd(int64(0))
}

/*
Returns the file cursor's position.
*/
func (self Instance) GetPosition() int { //gd:FileAccess.get_position
	return int(int(class(self).GetPosition()))
}

/*
Returns the size of the file in bytes.
*/
func (self Instance) GetLength() int { //gd:FileAccess.get_length
	return int(int(class(self).GetLength()))
}

/*
Returns [code]true[/code] if the file cursor has already read past the end of the file.
[b]Note:[/b] [code]eof_reached() == false[/code] cannot be used to check whether there is more data available. To loop while there is more data available, use:
[codeblocks]
[gdscript]
while file.get_position() < file.get_length():

	# Read data

[/gdscript]
[csharp]
while (file.GetPosition() < file.GetLength())

	{
	    // Read data
	}

[/csharp]
[/codeblocks]
*/
func (self Instance) EofReached() bool { //gd:FileAccess.eof_reached
	return bool(class(self).EofReached())
}

/*
Returns the next 8 bits from the file as an integer. See [method store_8] for details on what values can be stored and retrieved this way.
*/
func (self Instance) Get8() int { //gd:FileAccess.get_8
	return int(int(class(self).Get8()))
}

/*
Returns the next 16 bits from the file as an integer. See [method store_16] for details on what values can be stored and retrieved this way.
*/
func (self Instance) Get16() int { //gd:FileAccess.get_16
	return int(int(class(self).Get16()))
}

/*
Returns the next 32 bits from the file as an integer. See [method store_32] for details on what values can be stored and retrieved this way.
*/
func (self Instance) Get32() int { //gd:FileAccess.get_32
	return int(int(class(self).Get32()))
}

/*
Returns the next 64 bits from the file as an integer. See [method store_64] for details on what values can be stored and retrieved this way.
*/
func (self Instance) Get64() int { //gd:FileAccess.get_64
	return int(int(class(self).Get64()))
}

/*
Returns the next 32 bits from the file as a floating-point number.
*/
func (self Instance) GetFloat() Float.X { //gd:FileAccess.get_float
	return Float.X(Float.X(class(self).GetFloat()))
}

/*
Returns the next 64 bits from the file as a floating-point number.
*/
func (self Instance) GetDouble() Float.X { //gd:FileAccess.get_double
	return Float.X(Float.X(class(self).GetDouble()))
}

/*
Returns the next bits from the file as a floating-point number.
*/
func (self Instance) GetReal() Float.X { //gd:FileAccess.get_real
	return Float.X(Float.X(class(self).GetReal()))
}

/*
Returns next [param length] bytes of the file as a [PackedByteArray].
*/
func (self Instance) GetBuffer(length int) []byte { //gd:FileAccess.get_buffer
	return []byte(class(self).GetBuffer(int64(length)).Bytes())
}

/*
Returns the next line of the file as a [String]. The returned string doesn't include newline ([code]\n[/code]) or carriage return ([code]\r[/code]) characters, but does include any other leading or trailing whitespace.
Text is interpreted as being UTF-8 encoded.
*/
func (self Instance) GetLine() string { //gd:FileAccess.get_line
	return string(class(self).GetLine().String())
}

/*
Returns the next value of the file in CSV (Comma-Separated Values) format. You can pass a different delimiter [param delim] to use other than the default [code]","[/code] (comma). This delimiter must be one-character long, and cannot be a double quotation mark.
Text is interpreted as being UTF-8 encoded. Text values must be enclosed in double quotes if they include the delimiter character. Double quotes within a text value can be escaped by doubling their occurrence.
For example, the following CSV lines are valid and will be properly parsed as two strings each:
[codeblock lang=text]
Alice,"Hello, Bob!"
Bob,Alice! What a surprise!
Alice,"I thought you'd reply with ""Hello, world""."
[/codeblock]
Note how the second line can omit the enclosing quotes as it does not include the delimiter. However it [i]could[/i] very well use quotes, it was only written without for demonstration purposes. The third line must use [code]""[/code] for each quotation mark that needs to be interpreted as such instead of the end of a text value.
*/
func (self Instance) GetCsvLine() []string { //gd:FileAccess.get_csv_line
	return []string(class(self).GetCsvLine(String.New(",")).Strings())
}

/*
Returns the whole file as a [String]. Text is interpreted as being UTF-8 encoded.
If [param skip_cr] is [code]true[/code], carriage return characters ([code]\r[/code], CR) will be ignored when parsing the UTF-8, so that only line feed characters ([code]\n[/code], LF) represent a new line (Unix convention).
*/
func (self Instance) GetAsText() string { //gd:FileAccess.get_as_text
	return string(class(self).GetAsText(false).String())
}

/*
Returns an MD5 String representing the file at the given path or an empty [String] on failure.
*/
func GetMd5(path string) string { //gd:FileAccess.get_md5
	self := Instance{}
	return string(class(self).GetMd5(String.New(path)).String())
}

/*
Returns an SHA-256 [String] representing the file at the given path or an empty [String] on failure.
*/
func GetSha256(path string) string { //gd:FileAccess.get_sha256
	self := Instance{}
	return string(class(self).GetSha256(String.New(path)).String())
}

/*
Returns the last error that happened when trying to perform operations. Compare with the [code]ERR_FILE_*[/code] constants from [enum Error].
*/
func (self Instance) GetError() error { //gd:FileAccess.get_error
	return error(gd.ToError(class(self).GetError()))
}

/*
Returns the next [Variant] value from the file. If [param allow_objects] is [code]true[/code], decoding objects is allowed.
Internally, this uses the same decoding mechanism as the [method @GlobalScope.bytes_to_var] method.
[b]Warning:[/b] Deserialized objects can contain code which gets executed. Do not use this option if the serialized object comes from untrusted sources to avoid potential security threats such as remote code execution.
*/
func (self Instance) GetVar() any { //gd:FileAccess.get_var
	return any(class(self).GetVar(false).Interface())
}

/*
Stores an integer as 8 bits in the file.
[b]Note:[/b] The [param value] should lie in the interval [code][0, 255][/code]. Any other value will overflow and wrap around.
To store a signed integer, use [method store_64], or convert it manually (see [method store_16] for an example).
*/
func (self Instance) Store8(value int) { //gd:FileAccess.store_8
	class(self).Store8(int64(value))
}

/*
Stores an integer as 16 bits in the file.
[b]Note:[/b] The [param value] should lie in the interval [code][0, 2^16 - 1][/code]. Any other value will overflow and wrap around.
To store a signed integer, use [method store_64] or store a signed integer from the interval [code][-2^15, 2^15 - 1][/code] (i.e. keeping one bit for the signedness) and compute its sign manually when reading. For example:
[codeblocks]
[gdscript]
const MAX_15B = 1 << 15
const MAX_16B = 1 << 16

func unsigned16_to_signed(unsigned):

	return (unsigned + MAX_15B) % MAX_16B - MAX_15B

func _ready():

	var f = FileAccess.open("user://file.dat", FileAccess.WRITE_READ)
	f.store_16(-42) # This wraps around and stores 65494 (2^16 - 42).
	f.store_16(121) # In bounds, will store 121.
	f.seek(0) # Go back to start to read the stored value.
	var read1 = f.get_16() # 65494
	var read2 = f.get_16() # 121
	var converted1 = unsigned16_to_signed(read1) # -42
	var converted2 = unsigned16_to_signed(read2) # 121

[/gdscript]
[csharp]
public override void _Ready()

	{
	    using var f = FileAccess.Open("user://file.dat", FileAccess.ModeFlags.WriteRead);
	    f.Store16(unchecked((ushort)-42)); // This wraps around and stores 65494 (2^16 - 42).
	    f.Store16(121); // In bounds, will store 121.
	    f.Seek(0); // Go back to start to read the stored value.
	    ushort read1 = f.Get16(); // 65494
	    ushort read2 = f.Get16(); // 121
	    short converted1 = (short)read1; // -42
	    short converted2 = (short)read2; // 121
	}

[/csharp]
[/codeblocks]
*/
func (self Instance) Store16(value int) { //gd:FileAccess.store_16
	class(self).Store16(int64(value))
}

/*
Stores an integer as 32 bits in the file.
[b]Note:[/b] The [param value] should lie in the interval [code][0, 2^32 - 1][/code]. Any other value will overflow and wrap around.
To store a signed integer, use [method store_64], or convert it manually (see [method store_16] for an example).
*/
func (self Instance) Store32(value int) { //gd:FileAccess.store_32
	class(self).Store32(int64(value))
}

/*
Stores an integer as 64 bits in the file.
[b]Note:[/b] The [param value] must lie in the interval [code][-2^63, 2^63 - 1][/code] (i.e. be a valid [int] value).
*/
func (self Instance) Store64(value int) { //gd:FileAccess.store_64
	class(self).Store64(int64(value))
}

/*
Stores a floating-point number as 32 bits in the file.
*/
func (self Instance) StoreFloat(value Float.X) { //gd:FileAccess.store_float
	class(self).StoreFloat(float64(value))
}

/*
Stores a floating-point number as 64 bits in the file.
*/
func (self Instance) StoreDouble(value Float.X) { //gd:FileAccess.store_double
	class(self).StoreDouble(float64(value))
}

/*
Stores a floating-point number in the file.
*/
func (self Instance) StoreReal(value Float.X) { //gd:FileAccess.store_real
	class(self).StoreReal(float64(value))
}

/*
Stores the given array of bytes in the file.
*/
func (self Instance) StoreBuffer(buffer []byte) { //gd:FileAccess.store_buffer
	class(self).StoreBuffer(Packed.Bytes(Packed.New(buffer...)))
}

/*
Appends [param line] to the file followed by a line return character ([code]\n[/code]), encoding the text as UTF-8.
*/
func (self Instance) StoreLine(line string) { //gd:FileAccess.store_line
	class(self).StoreLine(String.New(line))
}

/*
Store the given [PackedStringArray] in the file as a line formatted in the CSV (Comma-Separated Values) format. You can pass a different delimiter [param delim] to use other than the default [code]","[/code] (comma). This delimiter must be one-character long.
Text will be encoded as UTF-8.
*/
func (self Instance) StoreCsvLine(values []string) { //gd:FileAccess.store_csv_line
	class(self).StoreCsvLine(Packed.MakeStrings(values...), String.New(","))
}

/*
Appends [param string] to the file without a line return, encoding the text as UTF-8.
[b]Note:[/b] This method is intended to be used to write text files. The string is stored as a UTF-8 encoded buffer without string length or terminating zero, which means that it can't be loaded back easily. If you want to store a retrievable string in a binary file, consider using [method store_pascal_string] instead. For retrieving strings from a text file, you can use [code]get_buffer(length).get_string_from_utf8()[/code] (if you know the length) or [method get_as_text].
*/
func (self Instance) StoreString(s string) { //gd:FileAccess.store_string
	class(self).StoreString(String.New(s))
}

/*
Stores any Variant value in the file. If [param full_objects] is [code]true[/code], encoding objects is allowed (and can potentially include code).
Internally, this uses the same encoding mechanism as the [method @GlobalScope.var_to_bytes] method.
[b]Note:[/b] Not all properties are included. Only properties that are configured with the [constant PROPERTY_USAGE_STORAGE] flag set will be serialized. You can add a new usage flag to a property by overriding the [method Object._get_property_list] method in your class. You can also check how property usage is configured by calling [method Object._get_property_list]. See [enum PropertyUsageFlags] for the possible usage flags.
*/
func (self Instance) StoreVar(value any) { //gd:FileAccess.store_var
	class(self).StoreVar(variant.New(value), false)
}

/*
Stores the given [String] as a line in the file in Pascal format (i.e. also store the length of the string).
Text will be encoded as UTF-8.
*/
func (self Instance) StorePascalString(s string) { //gd:FileAccess.store_pascal_string
	class(self).StorePascalString(String.New(s))
}

/*
Returns a [String] saved in Pascal format from the file.
Text is interpreted as being UTF-8 encoded.
*/
func (self Instance) GetPascalString() string { //gd:FileAccess.get_pascal_string
	return string(class(self).GetPascalString().String())
}

/*
Closes the currently opened file and prevents subsequent read/write operations. Use [method flush] to persist the data to disk without closing the file.
[b]Note:[/b] [FileAccess] will automatically close when it's freed, which happens when it goes out of scope or when it gets assigned with [code]null[/code]. In C# the reference must be disposed after we are done using it, this can be done with the [code]using[/code] statement or calling the [code]Dispose[/code] method directly.
*/
func (self Instance) Close() { //gd:FileAccess.close
	class(self).Close()
}

/*
Returns [code]true[/code] if the file exists in the given path.
[b]Note:[/b] Many resources types are imported (e.g. textures or sound files), and their source asset will not be included in the exported game, as only the imported version is used. See [method ResourceLoader.exists] for an alternative approach that takes resource remapping into account.
For a non-static, relative equivalent, use [method DirAccess.file_exists].
*/
func FileExists(path string) bool { //gd:FileAccess.file_exists
	self := Instance{}
	return bool(class(self).FileExists(String.New(path)))
}

/*
Returns the last time the [param file] was modified in Unix timestamp format, or [code]0[/code] on error. This Unix timestamp can be converted to another format using the [Time] singleton.
*/
func GetModifiedTime(file string) int { //gd:FileAccess.get_modified_time
	self := Instance{}
	return int(int(class(self).GetModifiedTime(String.New(file))))
}

/*
Returns file UNIX permissions.
[b]Note:[/b] This method is implemented on iOS, Linux/BSD, and macOS.
*/
func GetUnixPermissions(file string) gdclass.FileAccessUnixPermissionFlags { //gd:FileAccess.get_unix_permissions
	self := Instance{}
	return gdclass.FileAccessUnixPermissionFlags(class(self).GetUnixPermissions(String.New(file)))
}

/*
Sets file UNIX permissions.
[b]Note:[/b] This method is implemented on iOS, Linux/BSD, and macOS.
*/
func SetUnixPermissions(file string, permissions gdclass.FileAccessUnixPermissionFlags) error { //gd:FileAccess.set_unix_permissions
	self := Instance{}
	return error(gd.ToError(class(self).SetUnixPermissions(String.New(file), permissions)))
}

/*
Returns [code]true[/code], if file [code]hidden[/code] attribute is set.
[b]Note:[/b] This method is implemented on iOS, BSD, macOS, and Windows.
*/
func GetHiddenAttribute(file string) bool { //gd:FileAccess.get_hidden_attribute
	self := Instance{}
	return bool(class(self).GetHiddenAttribute(String.New(file)))
}

/*
Sets file [b]hidden[/b] attribute.
[b]Note:[/b] This method is implemented on iOS, BSD, macOS, and Windows.
*/
func SetHiddenAttribute(file string, hidden bool) error { //gd:FileAccess.set_hidden_attribute
	self := Instance{}
	return error(gd.ToError(class(self).SetHiddenAttribute(String.New(file), hidden)))
}

/*
Sets file [b]read only[/b] attribute.
[b]Note:[/b] This method is implemented on iOS, BSD, macOS, and Windows.
*/
func SetReadOnlyAttribute(file string, ro bool) error { //gd:FileAccess.set_read_only_attribute
	self := Instance{}
	return error(gd.ToError(class(self).SetReadOnlyAttribute(String.New(file), ro)))
}

/*
Returns [code]true[/code], if file [code]read only[/code] attribute is set.
[b]Note:[/b] This method is implemented on iOS, BSD, macOS, and Windows.
*/
func GetReadOnlyAttribute(file string) bool { //gd:FileAccess.get_read_only_attribute
	self := Instance{}
	return bool(class(self).GetReadOnlyAttribute(String.New(file)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.FileAccess

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("FileAccess"))
	casted := Instance{*(*gdclass.FileAccess)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) BigEndian() bool {
	return bool(class(self).IsBigEndian())
}

func (self Instance) SetBigEndian(value bool) {
	class(self).SetBigEndian(value)
}

/*
Creates a new [FileAccess] object and opens the file for writing or reading, depending on the flags.
Returns [code]null[/code] if opening the file failed. You can use [method get_open_error] to check the error that occurred.
*/
//go:nosplit
func (self class) Open(path String.Readable, flags gdclass.FileAccessModeFlags) [1]gdclass.FileAccess { //gd:FileAccess.open
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	callframe.Arg(frame, flags)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_open, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.FileAccess{gd.PointerWithOwnershipTransferredToGo[gdclass.FileAccess](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Creates a new [FileAccess] object and opens an encrypted file in write or read mode. You need to pass a binary key to encrypt/decrypt it.
[b]Note:[/b] The provided key must be 32 bytes long.
Returns [code]null[/code] if opening the file failed. You can use [method get_open_error] to check the error that occurred.
*/
//go:nosplit
func (self class) OpenEncrypted(path String.Readable, mode_flags gdclass.FileAccessModeFlags, key Packed.Bytes) [1]gdclass.FileAccess { //gd:FileAccess.open_encrypted
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	callframe.Arg(frame, mode_flags)
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedByteArray, byte](Packed.Array[byte](key))))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_open_encrypted, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.FileAccess{gd.PointerWithOwnershipTransferredToGo[gdclass.FileAccess](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Creates a new [FileAccess] object and opens an encrypted file in write or read mode. You need to pass a password to encrypt/decrypt it.
Returns [code]null[/code] if opening the file failed. You can use [method get_open_error] to check the error that occurred.
*/
//go:nosplit
func (self class) OpenEncryptedWithPass(path String.Readable, mode_flags gdclass.FileAccessModeFlags, pass String.Readable) [1]gdclass.FileAccess { //gd:FileAccess.open_encrypted_with_pass
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	callframe.Arg(frame, mode_flags)
	callframe.Arg(frame, pointers.Get(gd.InternalString(pass)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_open_encrypted_with_pass, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.FileAccess{gd.PointerWithOwnershipTransferredToGo[gdclass.FileAccess](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Creates a new [FileAccess] object and opens a compressed file for reading or writing.
[b]Note:[/b] [method open_compressed] can only read files that were saved by Godot, not third-party compression formats. See [url=https://github.com/godotengine/godot/issues/28999]GitHub issue #28999[/url] for a workaround.
Returns [code]null[/code] if opening the file failed. You can use [method get_open_error] to check the error that occurred.
*/
//go:nosplit
func (self class) OpenCompressed(path String.Readable, mode_flags gdclass.FileAccessModeFlags, compression_mode gdclass.FileAccessCompressionMode) [1]gdclass.FileAccess { //gd:FileAccess.open_compressed
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	callframe.Arg(frame, mode_flags)
	callframe.Arg(frame, compression_mode)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_open_compressed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.FileAccess{gd.PointerWithOwnershipTransferredToGo[gdclass.FileAccess](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the result of the last [method open] call in the current thread.
*/
//go:nosplit
func (self class) GetOpenError() Error.Code { //gd:FileAccess.get_open_error
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_get_open_error, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the whole [param path] file contents as a [PackedByteArray] without any decoding.
Returns an empty [PackedByteArray] if an error occurred while opening the file. You can use [method get_open_error] to check the error that occurred.
*/
//go:nosplit
func (self class) GetFileAsBytes(path String.Readable) Packed.Bytes { //gd:FileAccess.get_file_as_bytes
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_get_file_as_bytes, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Bytes(Array.Through(gd.PackedProxy[gd.PackedByteArray, byte]{}, pointers.Pack(pointers.New[gd.PackedByteArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the whole [param path] file contents as a [String]. Text is interpreted as being UTF-8 encoded.
Returns an empty [String] if an error occurred while opening the file. You can use [method get_open_error] to check the error that occurred.
*/
//go:nosplit
func (self class) GetFileAsString(path String.Readable) String.Readable { //gd:FileAccess.get_file_as_string
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_get_file_as_string, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Resizes the file to a specified length. The file must be open in a mode that permits writing. If the file is extended, NUL characters are appended. If the file is truncated, all data from the end file to the original length of the file is lost.
*/
//go:nosplit
func (self class) Resize(length int64) Error.Code { //gd:FileAccess.resize
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_resize, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Writes the file's buffer to disk. Flushing is automatically performed when the file is closed. This means you don't need to call [method flush] manually before closing a file. Still, calling [method flush] can be used to ensure the data is safe even if the project crashes instead of being closed gracefully.
[b]Note:[/b] Only call [method flush] when you actually need it. Otherwise, it will decrease performance due to constant disk writes.
*/
//go:nosplit
func (self class) Flush() { //gd:FileAccess.flush
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_flush, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the path as a [String] for the current open file.
*/
//go:nosplit
func (self class) GetPath() String.Readable { //gd:FileAccess.get_path
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_get_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the absolute path as a [String] for the current open file.
*/
//go:nosplit
func (self class) GetPathAbsolute() String.Readable { //gd:FileAccess.get_path_absolute
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_get_path_absolute, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the file is currently opened.
*/
//go:nosplit
func (self class) IsOpen() bool { //gd:FileAccess.is_open
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_is_open, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Changes the file reading/writing cursor to the specified position (in bytes from the beginning of the file).
*/
//go:nosplit
func (self class) SeekTo(position int64) { //gd:FileAccess.seek
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_seek, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Changes the file reading/writing cursor to the specified position (in bytes from the end of the file).
[b]Note:[/b] This is an offset, so you should use negative numbers or the cursor will be at the end of the file.
*/
//go:nosplit
func (self class) SeekEnd(position int64) { //gd:FileAccess.seek_end
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_seek_end, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the file cursor's position.
*/
//go:nosplit
func (self class) GetPosition() int64 { //gd:FileAccess.get_position
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_get_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the size of the file in bytes.
*/
//go:nosplit
func (self class) GetLength() int64 { //gd:FileAccess.get_length
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_get_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the file cursor has already read past the end of the file.
[b]Note:[/b] [code]eof_reached() == false[/code] cannot be used to check whether there is more data available. To loop while there is more data available, use:
[codeblocks]
[gdscript]
while file.get_position() < file.get_length():
    # Read data
[/gdscript]
[csharp]
while (file.GetPosition() < file.GetLength())
{
    // Read data
}
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) EofReached() bool { //gd:FileAccess.eof_reached
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_eof_reached, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the next 8 bits from the file as an integer. See [method store_8] for details on what values can be stored and retrieved this way.
*/
//go:nosplit
func (self class) Get8() int64 { //gd:FileAccess.get_8
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_get_8, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the next 16 bits from the file as an integer. See [method store_16] for details on what values can be stored and retrieved this way.
*/
//go:nosplit
func (self class) Get16() int64 { //gd:FileAccess.get_16
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_get_16, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the next 32 bits from the file as an integer. See [method store_32] for details on what values can be stored and retrieved this way.
*/
//go:nosplit
func (self class) Get32() int64 { //gd:FileAccess.get_32
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_get_32, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the next 64 bits from the file as an integer. See [method store_64] for details on what values can be stored and retrieved this way.
*/
//go:nosplit
func (self class) Get64() int64 { //gd:FileAccess.get_64
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_get_64, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the next 32 bits from the file as a floating-point number.
*/
//go:nosplit
func (self class) GetFloat() float64 { //gd:FileAccess.get_float
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_get_float, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the next 64 bits from the file as a floating-point number.
*/
//go:nosplit
func (self class) GetDouble() float64 { //gd:FileAccess.get_double
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_get_double, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the next bits from the file as a floating-point number.
*/
//go:nosplit
func (self class) GetReal() float64 { //gd:FileAccess.get_real
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_get_real, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns next [param length] bytes of the file as a [PackedByteArray].
*/
//go:nosplit
func (self class) GetBuffer(length int64) Packed.Bytes { //gd:FileAccess.get_buffer
	var frame = callframe.New()
	callframe.Arg(frame, length)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_get_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Bytes(Array.Through(gd.PackedProxy[gd.PackedByteArray, byte]{}, pointers.Pack(pointers.New[gd.PackedByteArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the next line of the file as a [String]. The returned string doesn't include newline ([code]\n[/code]) or carriage return ([code]\r[/code]) characters, but does include any other leading or trailing whitespace.
Text is interpreted as being UTF-8 encoded.
*/
//go:nosplit
func (self class) GetLine() String.Readable { //gd:FileAccess.get_line
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_get_line, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the next value of the file in CSV (Comma-Separated Values) format. You can pass a different delimiter [param delim] to use other than the default [code]","[/code] (comma). This delimiter must be one-character long, and cannot be a double quotation mark.
Text is interpreted as being UTF-8 encoded. Text values must be enclosed in double quotes if they include the delimiter character. Double quotes within a text value can be escaped by doubling their occurrence.
For example, the following CSV lines are valid and will be properly parsed as two strings each:
[codeblock lang=text]
Alice,"Hello, Bob!"
Bob,Alice! What a surprise!
Alice,"I thought you'd reply with ""Hello, world""."
[/codeblock]
Note how the second line can omit the enclosing quotes as it does not include the delimiter. However it [i]could[/i] very well use quotes, it was only written without for demonstration purposes. The third line must use [code]""[/code] for each quotation mark that needs to be interpreted as such instead of the end of a text value.
*/
//go:nosplit
func (self class) GetCsvLine(delim String.Readable) Packed.Strings { //gd:FileAccess.get_csv_line
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(delim)))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_get_csv_line, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the whole file as a [String]. Text is interpreted as being UTF-8 encoded.
If [param skip_cr] is [code]true[/code], carriage return characters ([code]\r[/code], CR) will be ignored when parsing the UTF-8, so that only line feed characters ([code]\n[/code], LF) represent a new line (Unix convention).
*/
//go:nosplit
func (self class) GetAsText(skip_cr bool) String.Readable { //gd:FileAccess.get_as_text
	var frame = callframe.New()
	callframe.Arg(frame, skip_cr)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_get_as_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns an MD5 String representing the file at the given path or an empty [String] on failure.
*/
//go:nosplit
func (self class) GetMd5(path String.Readable) String.Readable { //gd:FileAccess.get_md5
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_get_md5, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns an SHA-256 [String] representing the file at the given path or an empty [String] on failure.
*/
//go:nosplit
func (self class) GetSha256(path String.Readable) String.Readable { //gd:FileAccess.get_sha256
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_get_sha256, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) IsBigEndian() bool { //gd:FileAccess.is_big_endian
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_is_big_endian, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBigEndian(big_endian bool) { //gd:FileAccess.set_big_endian
	var frame = callframe.New()
	callframe.Arg(frame, big_endian)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_set_big_endian, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the last error that happened when trying to perform operations. Compare with the [code]ERR_FILE_*[/code] constants from [enum Error].
*/
//go:nosplit
func (self class) GetError() Error.Code { //gd:FileAccess.get_error
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_get_error, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the next [Variant] value from the file. If [param allow_objects] is [code]true[/code], decoding objects is allowed.
Internally, this uses the same decoding mechanism as the [method @GlobalScope.bytes_to_var] method.
[b]Warning:[/b] Deserialized objects can contain code which gets executed. Do not use this option if the serialized object comes from untrusted sources to avoid potential security threats such as remote code execution.
*/
//go:nosplit
func (self class) GetVar(allow_objects bool) variant.Any { //gd:FileAccess.get_var
	var frame = callframe.New()
	callframe.Arg(frame, allow_objects)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_get_var, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Implementation(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Stores an integer as 8 bits in the file.
[b]Note:[/b] The [param value] should lie in the interval [code][0, 255][/code]. Any other value will overflow and wrap around.
To store a signed integer, use [method store_64], or convert it manually (see [method store_16] for an example).
*/
//go:nosplit
func (self class) Store8(value int64) { //gd:FileAccess.store_8
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_store_8, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Stores an integer as 16 bits in the file.
[b]Note:[/b] The [param value] should lie in the interval [code][0, 2^16 - 1][/code]. Any other value will overflow and wrap around.
To store a signed integer, use [method store_64] or store a signed integer from the interval [code][-2^15, 2^15 - 1][/code] (i.e. keeping one bit for the signedness) and compute its sign manually when reading. For example:
[codeblocks]
[gdscript]
const MAX_15B = 1 << 15
const MAX_16B = 1 << 16

func unsigned16_to_signed(unsigned):
    return (unsigned + MAX_15B) % MAX_16B - MAX_15B

func _ready():
    var f = FileAccess.open("user://file.dat", FileAccess.WRITE_READ)
    f.store_16(-42) # This wraps around and stores 65494 (2^16 - 42).
    f.store_16(121) # In bounds, will store 121.
    f.seek(0) # Go back to start to read the stored value.
    var read1 = f.get_16() # 65494
    var read2 = f.get_16() # 121
    var converted1 = unsigned16_to_signed(read1) # -42
    var converted2 = unsigned16_to_signed(read2) # 121
[/gdscript]
[csharp]
public override void _Ready()
{
    using var f = FileAccess.Open("user://file.dat", FileAccess.ModeFlags.WriteRead);
    f.Store16(unchecked((ushort)-42)); // This wraps around and stores 65494 (2^16 - 42).
    f.Store16(121); // In bounds, will store 121.
    f.Seek(0); // Go back to start to read the stored value.
    ushort read1 = f.Get16(); // 65494
    ushort read2 = f.Get16(); // 121
    short converted1 = (short)read1; // -42
    short converted2 = (short)read2; // 121
}
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) Store16(value int64) { //gd:FileAccess.store_16
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_store_16, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Stores an integer as 32 bits in the file.
[b]Note:[/b] The [param value] should lie in the interval [code][0, 2^32 - 1][/code]. Any other value will overflow and wrap around.
To store a signed integer, use [method store_64], or convert it manually (see [method store_16] for an example).
*/
//go:nosplit
func (self class) Store32(value int64) { //gd:FileAccess.store_32
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_store_32, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Stores an integer as 64 bits in the file.
[b]Note:[/b] The [param value] must lie in the interval [code][-2^63, 2^63 - 1][/code] (i.e. be a valid [int] value).
*/
//go:nosplit
func (self class) Store64(value int64) { //gd:FileAccess.store_64
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_store_64, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Stores a floating-point number as 32 bits in the file.
*/
//go:nosplit
func (self class) StoreFloat(value float64) { //gd:FileAccess.store_float
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_store_float, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Stores a floating-point number as 64 bits in the file.
*/
//go:nosplit
func (self class) StoreDouble(value float64) { //gd:FileAccess.store_double
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_store_double, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Stores a floating-point number in the file.
*/
//go:nosplit
func (self class) StoreReal(value float64) { //gd:FileAccess.store_real
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_store_real, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Stores the given array of bytes in the file.
*/
//go:nosplit
func (self class) StoreBuffer(buffer Packed.Bytes) { //gd:FileAccess.store_buffer
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedByteArray, byte](Packed.Array[byte](buffer))))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_store_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Appends [param line] to the file followed by a line return character ([code]\n[/code]), encoding the text as UTF-8.
*/
//go:nosplit
func (self class) StoreLine(line String.Readable) { //gd:FileAccess.store_line
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(line)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_store_line, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Store the given [PackedStringArray] in the file as a line formatted in the CSV (Comma-Separated Values) format. You can pass a different delimiter [param delim] to use other than the default [code]","[/code] (comma). This delimiter must be one-character long.
Text will be encoded as UTF-8.
*/
//go:nosplit
func (self class) StoreCsvLine(values Packed.Strings, delim String.Readable) { //gd:FileAccess.store_csv_line
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPackedStrings(values)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(delim)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_store_csv_line, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Appends [param string] to the file without a line return, encoding the text as UTF-8.
[b]Note:[/b] This method is intended to be used to write text files. The string is stored as a UTF-8 encoded buffer without string length or terminating zero, which means that it can't be loaded back easily. If you want to store a retrievable string in a binary file, consider using [method store_pascal_string] instead. For retrieving strings from a text file, you can use [code]get_buffer(length).get_string_from_utf8()[/code] (if you know the length) or [method get_as_text].
*/
//go:nosplit
func (self class) StoreString(s String.Readable) { //gd:FileAccess.store_string
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(s)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_store_string, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Stores any Variant value in the file. If [param full_objects] is [code]true[/code], encoding objects is allowed (and can potentially include code).
Internally, this uses the same encoding mechanism as the [method @GlobalScope.var_to_bytes] method.
[b]Note:[/b] Not all properties are included. Only properties that are configured with the [constant PROPERTY_USAGE_STORAGE] flag set will be serialized. You can add a new usage flag to a property by overriding the [method Object._get_property_list] method in your class. You can also check how property usage is configured by calling [method Object._get_property_list]. See [enum PropertyUsageFlags] for the possible usage flags.
*/
//go:nosplit
func (self class) StoreVar(value variant.Any, full_objects bool) { //gd:FileAccess.store_var
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(value)))
	callframe.Arg(frame, full_objects)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_store_var, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Stores the given [String] as a line in the file in Pascal format (i.e. also store the length of the string).
Text will be encoded as UTF-8.
*/
//go:nosplit
func (self class) StorePascalString(s String.Readable) { //gd:FileAccess.store_pascal_string
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(s)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_store_pascal_string, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a [String] saved in Pascal format from the file.
Text is interpreted as being UTF-8 encoded.
*/
//go:nosplit
func (self class) GetPascalString() String.Readable { //gd:FileAccess.get_pascal_string
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_get_pascal_string, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Closes the currently opened file and prevents subsequent read/write operations. Use [method flush] to persist the data to disk without closing the file.
[b]Note:[/b] [FileAccess] will automatically close when it's freed, which happens when it goes out of scope or when it gets assigned with [code]null[/code]. In C# the reference must be disposed after we are done using it, this can be done with the [code]using[/code] statement or calling the [code]Dispose[/code] method directly.
*/
//go:nosplit
func (self class) Close() { //gd:FileAccess.close
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_close, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the file exists in the given path.
[b]Note:[/b] Many resources types are imported (e.g. textures or sound files), and their source asset will not be included in the exported game, as only the imported version is used. See [method ResourceLoader.exists] for an alternative approach that takes resource remapping into account.
For a non-static, relative equivalent, use [method DirAccess.file_exists].
*/
//go:nosplit
func (self class) FileExists(path String.Readable) bool { //gd:FileAccess.file_exists
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_file_exists, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the last time the [param file] was modified in Unix timestamp format, or [code]0[/code] on error. This Unix timestamp can be converted to another format using the [Time] singleton.
*/
//go:nosplit
func (self class) GetModifiedTime(file String.Readable) int64 { //gd:FileAccess.get_modified_time
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(file)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_get_modified_time, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns file UNIX permissions.
[b]Note:[/b] This method is implemented on iOS, Linux/BSD, and macOS.
*/
//go:nosplit
func (self class) GetUnixPermissions(file String.Readable) gdclass.FileAccessUnixPermissionFlags { //gd:FileAccess.get_unix_permissions
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(file)))
	var r_ret = callframe.Ret[gdclass.FileAccessUnixPermissionFlags](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_get_unix_permissions, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets file UNIX permissions.
[b]Note:[/b] This method is implemented on iOS, Linux/BSD, and macOS.
*/
//go:nosplit
func (self class) SetUnixPermissions(file String.Readable, permissions gdclass.FileAccessUnixPermissionFlags) Error.Code { //gd:FileAccess.set_unix_permissions
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(file)))
	callframe.Arg(frame, permissions)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_set_unix_permissions, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code], if file [code]hidden[/code] attribute is set.
[b]Note:[/b] This method is implemented on iOS, BSD, macOS, and Windows.
*/
//go:nosplit
func (self class) GetHiddenAttribute(file String.Readable) bool { //gd:FileAccess.get_hidden_attribute
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(file)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_get_hidden_attribute, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets file [b]hidden[/b] attribute.
[b]Note:[/b] This method is implemented on iOS, BSD, macOS, and Windows.
*/
//go:nosplit
func (self class) SetHiddenAttribute(file String.Readable, hidden bool) Error.Code { //gd:FileAccess.set_hidden_attribute
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(file)))
	callframe.Arg(frame, hidden)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_set_hidden_attribute, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Sets file [b]read only[/b] attribute.
[b]Note:[/b] This method is implemented on iOS, BSD, macOS, and Windows.
*/
//go:nosplit
func (self class) SetReadOnlyAttribute(file String.Readable, ro bool) Error.Code { //gd:FileAccess.set_read_only_attribute
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(file)))
	callframe.Arg(frame, ro)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_set_read_only_attribute, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code], if file [code]read only[/code] attribute is set.
[b]Note:[/b] This method is implemented on iOS, BSD, macOS, and Windows.
*/
//go:nosplit
func (self class) GetReadOnlyAttribute(file String.Readable) bool { //gd:FileAccess.get_read_only_attribute
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(file)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FileAccess.Bind_get_read_only_attribute, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsFileAccess() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsFileAccess() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("FileAccess", func(ptr gd.Object) any { return [1]gdclass.FileAccess{*(*gdclass.FileAccess)(unsafe.Pointer(&ptr))} })
}

type ModeFlags = gdclass.FileAccessModeFlags //gd:FileAccess.ModeFlags

const (
	/*Opens the file for read operations. The cursor is positioned at the beginning of the file.*/
	Read ModeFlags = 1
	/*Opens the file for write operations. The file is created if it does not exist, and truncated if it does.
	  [b]Note:[/b] When creating a file it must be in an already existing directory. To recursively create directories for a file path, see [method DirAccess.make_dir_recursive].*/
	Write ModeFlags = 2
	/*Opens the file for read and write operations. Does not truncate the file. The cursor is positioned at the beginning of the file.*/
	ReadWrite ModeFlags = 3
	/*Opens the file for read and write operations. The file is created if it does not exist, and truncated if it does. The cursor is positioned at the beginning of the file.
	  [b]Note:[/b] When creating a file it must be in an already existing directory. To recursively create directories for a file path, see [method DirAccess.make_dir_recursive].*/
	WriteRead ModeFlags = 7
)

type CompressionMode = gdclass.FileAccessCompressionMode //gd:FileAccess.CompressionMode

const (
	/*Uses the [url=https://fastlz.org/]FastLZ[/url] compression method.*/
	CompressionFastlz CompressionMode = 0
	/*Uses the [url=https://en.wikipedia.org/wiki/DEFLATE]DEFLATE[/url] compression method.*/
	CompressionDeflate CompressionMode = 1
	/*Uses the [url=https://facebook.github.io/zstd/]Zstandard[/url] compression method.*/
	CompressionZstd CompressionMode = 2
	/*Uses the [url=https://www.gzip.org/]gzip[/url] compression method.*/
	CompressionGzip CompressionMode = 3
	/*Uses the [url=https://github.com/google/brotli]brotli[/url] compression method (only decompression is supported).*/
	CompressionBrotli CompressionMode = 4
)

type UnixPermissionFlags = gdclass.FileAccessUnixPermissionFlags //gd:FileAccess.UnixPermissionFlags

const (
	/*Read for owner bit.*/
	UnixReadOwner UnixPermissionFlags = 256
	/*Write for owner bit.*/
	UnixWriteOwner UnixPermissionFlags = 128
	/*Execute for owner bit.*/
	UnixExecuteOwner UnixPermissionFlags = 64
	/*Read for group bit.*/
	UnixReadGroup UnixPermissionFlags = 32
	/*Write for group bit.*/
	UnixWriteGroup UnixPermissionFlags = 16
	/*Execute for group bit.*/
	UnixExecuteGroup UnixPermissionFlags = 8
	/*Read for other bit.*/
	UnixReadOther UnixPermissionFlags = 4
	/*Write for other bit.*/
	UnixWriteOther UnixPermissionFlags = 2
	/*Execute for other bit.*/
	UnixExecuteOther UnixPermissionFlags = 1
	/*Set user id on execution bit.*/
	UnixSetUserId UnixPermissionFlags = 2048
	/*Set group id on execution bit.*/
	UnixSetGroupId UnixPermissionFlags = 1024
	/*Restricted deletion (sticky) bit.*/
	UnixRestrictedDelete UnixPermissionFlags = 512
)
