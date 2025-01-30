// Package HashingContext provides methods for working with HashingContext object instances.
package HashingContext

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
The HashingContext class provides an interface for computing cryptographic hashes over multiple iterations. Useful for computing hashes of big files (so you don't have to load them all in memory), network streams, and data streams in general (so you don't have to hold buffers).
The [enum HashType] enum shows the supported hashing algorithms.
[codeblocks]
[gdscript]
const CHUNK_SIZE = 1024

func hash_file(path):

	# Check that file exists.
	if not FileAccess.file_exists(path):
	    return
	# Start an SHA-256 context.
	var ctx = HashingContext.new()
	ctx.start(HashingContext.HASH_SHA256)
	# Open the file to hash.
	var file = FileAccess.open(path, FileAccess.READ)
	# Update the context after reading each chunk.
	while file.get_position() < file.get_length():
	    var remaining = file.get_length() - file.get_position()
	    ctx.update(file.get_buffer(min(remaining, CHUNK_SIZE)))
	# Get the computed hash.
	var res = ctx.finish()
	# Print the result as hex string and array.
	printt(res.hex_encode(), Array(res))

[/gdscript]
[csharp]
public const int ChunkSize = 1024;

public void HashFile(string path)

	{
	    // Check that file exists.
	    if (!FileAccess.FileExists(path))
	    {
	        return;
	    }
	    // Start an SHA-256 context.
	    var ctx = new HashingContext();
	    ctx.Start(HashingContext.HashType.Sha256);
	    // Open the file to hash.
	    using var file = FileAccess.Open(path, FileAccess.ModeFlags.Read);
	    // Update the context after reading each chunk.
	    while (file.GetPosition() < file.GetLength())
	    {
	        int remaining = (int)(file.GetLength() - file.GetPosition());
	        ctx.Update(file.GetBuffer(Mathf.Min(remaining, ChunkSize)));
	    }
	    // Get the computed hash.
	    byte[] res = ctx.Finish();
	    // Print the result as hex string and array.
	    GD.PrintT(res.HexEncode(), (Variant)res);
	}

[/csharp]
[/codeblocks]
*/
type Instance [1]gdclass.HashingContext

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsHashingContext() Instance
}

/*
Starts a new hash computation of the given [param type] (e.g. [constant HASH_SHA256] to start computation of an SHA-256).
*/
func (self Instance) Start(atype gdclass.HashingContextHashType) error { //gd:HashingContext.start
	return error(gd.ToError(class(self).Start(atype)))
}

/*
Updates the computation with the given [param chunk] of data.
*/
func (self Instance) Update(chunk []byte) error { //gd:HashingContext.update
	return error(gd.ToError(class(self).Update(Packed.Bytes(Packed.New(chunk...)))))
}

/*
Closes the current context, and return the computed hash.
*/
func (self Instance) Finish() []byte { //gd:HashingContext.finish
	return []byte(class(self).Finish().Bytes())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.HashingContext

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("HashingContext"))
	casted := Instance{*(*gdclass.HashingContext)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Starts a new hash computation of the given [param type] (e.g. [constant HASH_SHA256] to start computation of an SHA-256).
*/
//go:nosplit
func (self class) Start(atype gdclass.HashingContextHashType) Error.Code { //gd:HashingContext.start
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HashingContext.Bind_start, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Updates the computation with the given [param chunk] of data.
*/
//go:nosplit
func (self class) Update(chunk Packed.Bytes) Error.Code { //gd:HashingContext.update
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedByteArray, byte](Packed.Array[byte](chunk))))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HashingContext.Bind_update, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Closes the current context, and return the computed hash.
*/
//go:nosplit
func (self class) Finish() Packed.Bytes { //gd:HashingContext.finish
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HashingContext.Bind_finish, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Bytes(Array.Through(gd.PackedProxy[gd.PackedByteArray, byte]{}, pointers.Pack(pointers.New[gd.PackedByteArray](r_ret.Get()))))
	frame.Free()
	return ret
}
func (self class) AsHashingContext() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsHashingContext() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("HashingContext", func(ptr gd.Object) any {
		return [1]gdclass.HashingContext{*(*gdclass.HashingContext)(unsafe.Pointer(&ptr))}
	})
}

type HashType = gdclass.HashingContextHashType //gd:HashingContext.HashType

const (
	/*Hashing algorithm: MD5.*/
	HashMd5 HashType = 0
	/*Hashing algorithm: SHA-1.*/
	HashSha1 HashType = 1
	/*Hashing algorithm: SHA-256.*/
	HashSha256 HashType = 2
)
