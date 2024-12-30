package StreamPeer

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Float"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
StreamPeer is an abstract base class mostly used for stream-based protocols (such as TCP). It provides an API for sending and receiving data through streams as raw data or strings.
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.
*/
type Instance [1]classdb.StreamPeer
type Any interface {
	gd.IsClass
	AsStreamPeer() Instance
}

/*
Sends a chunk of data through the connection, blocking if necessary until the data is done sending. This function returns an [enum Error] code.
*/
func (self Instance) PutData(data []byte) error {
	return error(class(self).PutData(gd.NewPackedByteSlice(data)))
}

/*
Sends a chunk of data through the connection. If all the data could not be sent at once, only part of it will. This function returns two values, an [enum Error] code and an integer, describing how much data was actually sent.
*/
func (self Instance) PutPartialData(data []byte) Array.Any {
	return Array.Any(class(self).PutPartialData(gd.NewPackedByteSlice(data)))
}

/*
Returns a chunk data with the received bytes. The number of bytes to be received can be requested in the [param bytes] argument. If not enough bytes are available, the function will block until the desired amount is received. This function returns two values, an [enum Error] code and a data array.
*/
func (self Instance) GetData(bytes int) Array.Any {
	return Array.Any(class(self).GetData(gd.Int(bytes)))
}

/*
Returns a chunk data with the received bytes. The number of bytes to be received can be requested in the "bytes" argument. If not enough bytes are available, the function will return how many were actually received. This function returns two values, an [enum Error] code, and a data array.
*/
func (self Instance) GetPartialData(bytes int) Array.Any {
	return Array.Any(class(self).GetPartialData(gd.Int(bytes)))
}

/*
Returns the number of bytes this [StreamPeer] has available.
*/
func (self Instance) GetAvailableBytes() int {
	return int(int(class(self).GetAvailableBytes()))
}

/*
Puts a signed byte into the stream.
*/
func (self Instance) Put8(value int) {
	class(self).Put8(gd.Int(value))
}

/*
Puts an unsigned byte into the stream.
*/
func (self Instance) PutU8(value int) {
	class(self).PutU8(gd.Int(value))
}

/*
Puts a signed 16-bit value into the stream.
*/
func (self Instance) Put16(value int) {
	class(self).Put16(gd.Int(value))
}

/*
Puts an unsigned 16-bit value into the stream.
*/
func (self Instance) PutU16(value int) {
	class(self).PutU16(gd.Int(value))
}

/*
Puts a signed 32-bit value into the stream.
*/
func (self Instance) Put32(value int) {
	class(self).Put32(gd.Int(value))
}

/*
Puts an unsigned 32-bit value into the stream.
*/
func (self Instance) PutU32(value int) {
	class(self).PutU32(gd.Int(value))
}

/*
Puts a signed 64-bit value into the stream.
*/
func (self Instance) Put64(value int) {
	class(self).Put64(gd.Int(value))
}

/*
Puts an unsigned 64-bit value into the stream.
*/
func (self Instance) PutU64(value int) {
	class(self).PutU64(gd.Int(value))
}

/*
Puts a single-precision float into the stream.
*/
func (self Instance) PutFloat(value Float.X) {
	class(self).PutFloat(gd.Float(value))
}

/*
Puts a double-precision float into the stream.
*/
func (self Instance) PutDouble(value Float.X) {
	class(self).PutDouble(gd.Float(value))
}

/*
Puts a zero-terminated ASCII string into the stream prepended by a 32-bit unsigned integer representing its size.
[b]Note:[/b] To put an ASCII string without prepending its size, you can use [method put_data]:
[codeblocks]
[gdscript]
put_data("Hello world".to_ascii_buffer())
[/gdscript]
[csharp]
PutData("Hello World".ToAsciiBuffer());
[/csharp]
[/codeblocks]
*/
func (self Instance) PutString(value string) {
	class(self).PutString(gd.NewString(value))
}

/*
Puts a zero-terminated UTF-8 string into the stream prepended by a 32 bits unsigned integer representing its size.
[b]Note:[/b] To put a UTF-8 string without prepending its size, you can use [method put_data]:
[codeblocks]
[gdscript]
put_data("Hello world".to_utf8_buffer())
[/gdscript]
[csharp]
PutData("Hello World".ToUtf8Buffer());
[/csharp]
[/codeblocks]
*/
func (self Instance) PutUtf8String(value string) {
	class(self).PutUtf8String(gd.NewString(value))
}

/*
Puts a Variant into the stream. If [param full_objects] is [code]true[/code] encoding objects is allowed (and can potentially include code).
Internally, this uses the same encoding mechanism as the [method @GlobalScope.var_to_bytes] method.
*/
func (self Instance) PutVar(value any) {
	class(self).PutVar(gd.NewVariant(value), false)
}

/*
Gets a signed byte from the stream.
*/
func (self Instance) Get8() int {
	return int(int(class(self).Get8()))
}

/*
Gets an unsigned byte from the stream.
*/
func (self Instance) GetU8() int {
	return int(int(class(self).GetU8()))
}

/*
Gets a signed 16-bit value from the stream.
*/
func (self Instance) Get16() int {
	return int(int(class(self).Get16()))
}

/*
Gets an unsigned 16-bit value from the stream.
*/
func (self Instance) GetU16() int {
	return int(int(class(self).GetU16()))
}

/*
Gets a signed 32-bit value from the stream.
*/
func (self Instance) Get32() int {
	return int(int(class(self).Get32()))
}

/*
Gets an unsigned 32-bit value from the stream.
*/
func (self Instance) GetU32() int {
	return int(int(class(self).GetU32()))
}

/*
Gets a signed 64-bit value from the stream.
*/
func (self Instance) Get64() int {
	return int(int(class(self).Get64()))
}

/*
Gets an unsigned 64-bit value from the stream.
*/
func (self Instance) GetU64() int {
	return int(int(class(self).GetU64()))
}

/*
Gets a single-precision float from the stream.
*/
func (self Instance) GetFloat() Float.X {
	return Float.X(Float.X(class(self).GetFloat()))
}

/*
Gets a double-precision float from the stream.
*/
func (self Instance) GetDouble() Float.X {
	return Float.X(Float.X(class(self).GetDouble()))
}

/*
Gets an ASCII string with byte-length [param bytes] from the stream. If [param bytes] is negative (default) the length will be read from the stream using the reverse process of [method put_string].
*/
func (self Instance) GetString() string {
	return string(class(self).GetString(gd.Int(-1)).String())
}

/*
Gets a UTF-8 string with byte-length [param bytes] from the stream (this decodes the string sent as UTF-8). If [param bytes] is negative (default) the length will be read from the stream using the reverse process of [method put_utf8_string].
*/
func (self Instance) GetUtf8String() string {
	return string(class(self).GetUtf8String(gd.Int(-1)).String())
}

/*
Gets a Variant from the stream. If [param allow_objects] is [code]true[/code], decoding objects is allowed.
Internally, this uses the same decoding mechanism as the [method @GlobalScope.bytes_to_var] method.
[b]Warning:[/b] Deserialized objects can contain code which gets executed. Do not use this option if the serialized object comes from untrusted sources to avoid potential security threats such as remote code execution.
*/
func (self Instance) GetVar() any {
	return any(class(self).GetVar(false).Interface())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.StreamPeer

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("StreamPeer"))
	return Instance{classdb.StreamPeer(object)}
}

func (self Instance) BigEndian() bool {
	return bool(class(self).IsBigEndianEnabled())
}

func (self Instance) SetBigEndian(value bool) {
	class(self).SetBigEndian(value)
}

/*
Sends a chunk of data through the connection, blocking if necessary until the data is done sending. This function returns an [enum Error] code.
*/
//go:nosplit
func (self class) PutData(data gd.PackedByteArray) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(data))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_put_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sends a chunk of data through the connection. If all the data could not be sent at once, only part of it will. This function returns two values, an [enum Error] code and an integer, describing how much data was actually sent.
*/
//go:nosplit
func (self class) PutPartialData(data gd.PackedByteArray) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(data))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_put_partial_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a chunk data with the received bytes. The number of bytes to be received can be requested in the [param bytes] argument. If not enough bytes are available, the function will block until the desired amount is received. This function returns two values, an [enum Error] code and a data array.
*/
//go:nosplit
func (self class) GetData(bytes gd.Int) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, bytes)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_get_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a chunk data with the received bytes. The number of bytes to be received can be requested in the "bytes" argument. If not enough bytes are available, the function will return how many were actually received. This function returns two values, an [enum Error] code, and a data array.
*/
//go:nosplit
func (self class) GetPartialData(bytes gd.Int) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, bytes)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_get_partial_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the number of bytes this [StreamPeer] has available.
*/
//go:nosplit
func (self class) GetAvailableBytes() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_get_available_bytes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBigEndian(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_set_big_endian, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsBigEndianEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_is_big_endian_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Puts a signed byte into the stream.
*/
//go:nosplit
func (self class) Put8(value gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_put_8, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Puts an unsigned byte into the stream.
*/
//go:nosplit
func (self class) PutU8(value gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_put_u8, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Puts a signed 16-bit value into the stream.
*/
//go:nosplit
func (self class) Put16(value gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_put_16, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Puts an unsigned 16-bit value into the stream.
*/
//go:nosplit
func (self class) PutU16(value gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_put_u16, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Puts a signed 32-bit value into the stream.
*/
//go:nosplit
func (self class) Put32(value gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_put_32, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Puts an unsigned 32-bit value into the stream.
*/
//go:nosplit
func (self class) PutU32(value gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_put_u32, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Puts a signed 64-bit value into the stream.
*/
//go:nosplit
func (self class) Put64(value gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_put_64, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Puts an unsigned 64-bit value into the stream.
*/
//go:nosplit
func (self class) PutU64(value gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_put_u64, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Puts a single-precision float into the stream.
*/
//go:nosplit
func (self class) PutFloat(value gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_put_float, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Puts a double-precision float into the stream.
*/
//go:nosplit
func (self class) PutDouble(value gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, value)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_put_double, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Puts a zero-terminated ASCII string into the stream prepended by a 32-bit unsigned integer representing its size.
[b]Note:[/b] To put an ASCII string without prepending its size, you can use [method put_data]:
[codeblocks]
[gdscript]
put_data("Hello world".to_ascii_buffer())
[/gdscript]
[csharp]
PutData("Hello World".ToAsciiBuffer());
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) PutString(value gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(value))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_put_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Puts a zero-terminated UTF-8 string into the stream prepended by a 32 bits unsigned integer representing its size.
[b]Note:[/b] To put a UTF-8 string without prepending its size, you can use [method put_data]:
[codeblocks]
[gdscript]
put_data("Hello world".to_utf8_buffer())
[/gdscript]
[csharp]
PutData("Hello World".ToUtf8Buffer());
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) PutUtf8String(value gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(value))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_put_utf8_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Puts a Variant into the stream. If [param full_objects] is [code]true[/code] encoding objects is allowed (and can potentially include code).
Internally, this uses the same encoding mechanism as the [method @GlobalScope.var_to_bytes] method.
*/
//go:nosplit
func (self class) PutVar(value gd.Variant, full_objects bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(value))
	callframe.Arg(frame, full_objects)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_put_var, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Gets a signed byte from the stream.
*/
//go:nosplit
func (self class) Get8() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_get_8, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Gets an unsigned byte from the stream.
*/
//go:nosplit
func (self class) GetU8() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_get_u8, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Gets a signed 16-bit value from the stream.
*/
//go:nosplit
func (self class) Get16() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_get_16, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Gets an unsigned 16-bit value from the stream.
*/
//go:nosplit
func (self class) GetU16() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_get_u16, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Gets a signed 32-bit value from the stream.
*/
//go:nosplit
func (self class) Get32() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_get_32, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Gets an unsigned 32-bit value from the stream.
*/
//go:nosplit
func (self class) GetU32() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_get_u32, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Gets a signed 64-bit value from the stream.
*/
//go:nosplit
func (self class) Get64() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_get_64, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Gets an unsigned 64-bit value from the stream.
*/
//go:nosplit
func (self class) GetU64() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_get_u64, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Gets a single-precision float from the stream.
*/
//go:nosplit
func (self class) GetFloat() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_get_float, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Gets a double-precision float from the stream.
*/
//go:nosplit
func (self class) GetDouble() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_get_double, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Gets an ASCII string with byte-length [param bytes] from the stream. If [param bytes] is negative (default) the length will be read from the stream using the reverse process of [method put_string].
*/
//go:nosplit
func (self class) GetString(bytes gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, bytes)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_get_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Gets a UTF-8 string with byte-length [param bytes] from the stream (this decodes the string sent as UTF-8). If [param bytes] is negative (default) the length will be read from the stream using the reverse process of [method put_utf8_string].
*/
//go:nosplit
func (self class) GetUtf8String(bytes gd.Int) gd.String {
	var frame = callframe.New()
	callframe.Arg(frame, bytes)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_get_utf8_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Gets a Variant from the stream. If [param allow_objects] is [code]true[/code], decoding objects is allowed.
Internally, this uses the same decoding mechanism as the [method @GlobalScope.bytes_to_var] method.
[b]Warning:[/b] Deserialized objects can contain code which gets executed. Do not use this option if the serialized object comes from untrusted sources to avoid potential security threats such as remote code execution.
*/
//go:nosplit
func (self class) GetVar(allow_objects bool) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, allow_objects)
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_get_var, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsStreamPeer() Advanced         { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsStreamPeer() Instance      { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {
	classdb.Register("StreamPeer", func(ptr gd.Object) any { return classdb.StreamPeer(ptr) })
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
