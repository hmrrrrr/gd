// Package ZIPReader provides methods for working with ZIPReader object instances.
package ZIPReader

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
This class implements a reader that can extract the content of individual files inside a zip archive.
[codeblock]
func read_zip_file():

	var reader := ZIPReader.new()
	var err := reader.open("user://archive.zip")
	if err != OK:
	    return PackedByteArray()
	var res := reader.read_file("hello.txt")
	reader.close()
	return res

[/codeblock]
*/
type Instance [1]gdclass.ZIPReader

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsZIPReader() Instance
}

/*
Opens the zip archive at the given [param path] and reads its file index.
*/
func (self Instance) Open(path string) error { //gd:ZIPReader.open
	return error(gd.ToError(class(self).Open(String.New(path))))
}

/*
Closes the underlying resources used by this instance.
*/
func (self Instance) Close() error { //gd:ZIPReader.close
	return error(gd.ToError(class(self).Close()))
}

/*
Returns the list of names of all files in the loaded archive.
Must be called after [method open].
*/
func (self Instance) GetFiles() []string { //gd:ZIPReader.get_files
	return []string(class(self).GetFiles().Strings())
}

/*
Loads the whole content of a file in the loaded zip archive into memory and returns it.
Must be called after [method open].
*/
func (self Instance) ReadFile(path string) []byte { //gd:ZIPReader.read_file
	return []byte(class(self).ReadFile(String.New(path), true).Bytes())
}

/*
Returns [code]true[/code] if the file exists in the loaded zip archive.
Must be called after [method open].
*/
func (self Instance) FileExists(path string) bool { //gd:ZIPReader.file_exists
	return bool(class(self).FileExists(String.New(path), true))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ZIPReader

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ZIPReader"))
	casted := Instance{*(*gdclass.ZIPReader)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Opens the zip archive at the given [param path] and reads its file index.
*/
//go:nosplit
func (self class) Open(path String.Readable) Error.Code { //gd:ZIPReader.open
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ZIPReader.Bind_open, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Closes the underlying resources used by this instance.
*/
//go:nosplit
func (self class) Close() Error.Code { //gd:ZIPReader.close
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ZIPReader.Bind_close, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the list of names of all files in the loaded archive.
Must be called after [method open].
*/
//go:nosplit
func (self class) GetFiles() Packed.Strings { //gd:ZIPReader.get_files
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ZIPReader.Bind_get_files, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Loads the whole content of a file in the loaded zip archive into memory and returns it.
Must be called after [method open].
*/
//go:nosplit
func (self class) ReadFile(path String.Readable, case_sensitive bool) Packed.Bytes { //gd:ZIPReader.read_file
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	callframe.Arg(frame, case_sensitive)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ZIPReader.Bind_read_file, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Bytes(Array.Through(gd.PackedProxy[gd.PackedByteArray, byte]{}, pointers.Pack(pointers.New[gd.PackedByteArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the file exists in the loaded zip archive.
Must be called after [method open].
*/
//go:nosplit
func (self class) FileExists(path String.Readable, case_sensitive bool) bool { //gd:ZIPReader.file_exists
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	callframe.Arg(frame, case_sensitive)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ZIPReader.Bind_file_exists, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsZIPReader() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsZIPReader() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("ZIPReader", func(ptr gd.Object) any { return [1]gdclass.ZIPReader{*(*gdclass.ZIPReader)(unsafe.Pointer(&ptr))} })
}
