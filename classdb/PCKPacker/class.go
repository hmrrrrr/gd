// Package PCKPacker provides methods for working with PCKPacker object instances.
package PCKPacker

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
The [PCKPacker] is used to create packages that can be loaded into a running project using [method ProjectSettings.load_resource_pack].
[codeblocks]
[gdscript]
var packer = PCKPacker.new()
packer.pck_start("test.pck")
packer.add_file("res://text.txt", "text.txt")
packer.flush()
[/gdscript]
[csharp]
var packer = new PckPacker();
packer.PckStart("test.pck");
packer.AddFile("res://text.txt", "text.txt");
packer.Flush();
[/csharp]
[/codeblocks]
The above [PCKPacker] creates package [code]test.pck[/code], then adds a file named [code]text.txt[/code] at the root of the package.
*/
type Instance [1]gdclass.PCKPacker

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPCKPacker() Instance
}

/*
Creates a new PCK file with the name [param pck_name]. The [code].pck[/code] file extension isn't added automatically, so it should be part of [param pck_name] (even though it's not required).
*/
func (self Instance) PckStart(pck_name string) error { //gd:PCKPacker.pck_start
	return error(gd.ToError(class(self).PckStart(String.New(pck_name), int64(32), String.New("0000000000000000000000000000000000000000000000000000000000000000"), false)))
}

/*
Adds the [param source_path] file to the current PCK package at the [param pck_path] internal path (should start with [code]res://[/code]).
*/
func (self Instance) AddFile(pck_path string, source_path string) error { //gd:PCKPacker.add_file
	return error(gd.ToError(class(self).AddFile(String.New(pck_path), String.New(source_path), false)))
}

/*
Writes the files specified using all [method add_file] calls since the last flush. If [param verbose] is [code]true[/code], a list of files added will be printed to the console for easier debugging.
*/
func (self Instance) Flush() error { //gd:PCKPacker.flush
	return error(gd.ToError(class(self).Flush(false)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PCKPacker

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PCKPacker"))
	casted := Instance{*(*gdclass.PCKPacker)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Creates a new PCK file with the name [param pck_name]. The [code].pck[/code] file extension isn't added automatically, so it should be part of [param pck_name] (even though it's not required).
*/
//go:nosplit
func (self class) PckStart(pck_name String.Readable, alignment int64, key String.Readable, encrypt_directory bool) Error.Code { //gd:PCKPacker.pck_start
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(pck_name)))
	callframe.Arg(frame, alignment)
	callframe.Arg(frame, pointers.Get(gd.InternalString(key)))
	callframe.Arg(frame, encrypt_directory)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PCKPacker.Bind_pck_start, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Adds the [param source_path] file to the current PCK package at the [param pck_path] internal path (should start with [code]res://[/code]).
*/
//go:nosplit
func (self class) AddFile(pck_path String.Readable, source_path String.Readable, encrypt bool) Error.Code { //gd:PCKPacker.add_file
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(pck_path)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(source_path)))
	callframe.Arg(frame, encrypt)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PCKPacker.Bind_add_file, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Writes the files specified using all [method add_file] calls since the last flush. If [param verbose] is [code]true[/code], a list of files added will be printed to the console for easier debugging.
*/
//go:nosplit
func (self class) Flush(verbose bool) Error.Code { //gd:PCKPacker.flush
	var frame = callframe.New()
	callframe.Arg(frame, verbose)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PCKPacker.Bind_flush, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsPCKPacker() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPCKPacker() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("PCKPacker", func(ptr gd.Object) any { return [1]gdclass.PCKPacker{*(*gdclass.PCKPacker)(unsafe.Pointer(&ptr))} })
}
