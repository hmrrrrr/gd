package StreamPeer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
StreamPeer is an abstract base class mostly used for stream-based protocols (such as TCP). It provides an API for sending and receiving data through streams as raw data or strings.
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.

*/
type Go [1]classdb.StreamPeer

/*
Sends a chunk of data through the connection, blocking if necessary until the data is done sending. This function returns an [enum Error] code.
*/
func (self Go) PutData(data []byte) gd.Error {
	return gd.Error(class(self).PutData(gd.NewPackedByteSlice(data)))
}

/*
Sends a chunk of data through the connection. If all the data could not be sent at once, only part of it will. This function returns two values, an [enum Error] code and an integer, describing how much data was actually sent.
*/
func (self Go) PutPartialData(data []byte) gd.Array {
	return gd.Array(class(self).PutPartialData(gd.NewPackedByteSlice(data)))
}

/*
Returns a chunk data with the received bytes. The number of bytes to be received can be requested in the [param bytes] argument. If not enough bytes are available, the function will block until the desired amount is received. This function returns two values, an [enum Error] code and a data array.
*/
func (self Go) GetData(bytes int) gd.Array {
	return gd.Array(class(self).GetData(gd.Int(bytes)))
}

/*
Returns a chunk data with the received bytes. The number of bytes to be received can be requested in the "bytes" argument. If not enough bytes are available, the function will return how many were actually received. This function returns two values, an [enum Error] code, and a data array.
*/
func (self Go) GetPartialData(bytes int) gd.Array {
	return gd.Array(class(self).GetPartialData(gd.Int(bytes)))
}

/*
Returns the number of bytes this [StreamPeer] has available.
*/
func (self Go) GetAvailableBytes() int {
	return int(int(class(self).GetAvailableBytes()))
}

/*
Puts a signed byte into the stream.
*/
func (self Go) Put8(value int) {
	class(self).Put8(gd.Int(value))
}

/*
Puts an unsigned byte into the stream.
*/
func (self Go) PutU8(value int) {
	class(self).PutU8(gd.Int(value))
}

/*
Puts a signed 16-bit value into the stream.
*/
func (self Go) Put16(value int) {
	class(self).Put16(gd.Int(value))
}

/*
Puts an unsigned 16-bit value into the stream.
*/
func (self Go) PutU16(value int) {
	class(self).PutU16(gd.Int(value))
}

/*
Puts a signed 32-bit value into the stream.
*/
func (self Go) Put32(value int) {
	class(self).Put32(gd.Int(value))
}

/*
Puts an unsigned 32-bit value into the stream.
*/
func (self Go) PutU32(value int) {
	class(self).PutU32(gd.Int(value))
}

/*
Puts a signed 64-bit value into the stream.
*/
func (self Go) Put64(value int) {
	class(self).Put64(gd.Int(value))
}

/*
Puts an unsigned 64-bit value into the stream.
*/
func (self Go) PutU64(value int) {
	class(self).PutU64(gd.Int(value))
}

/*
Puts a single-precision float into the stream.
*/
func (self Go) PutFloat(value float64) {
	class(self).PutFloat(gd.Float(value))
}

/*
Puts a double-precision float into the stream.
*/
func (self Go) PutDouble(value float64) {
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
func (self Go) PutString(value string) {
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
func (self Go) PutUtf8String(value string) {
	class(self).PutUtf8String(gd.NewString(value))
}

/*
Puts a Variant into the stream. If [param full_objects] is [code]true[/code] encoding objects is allowed (and can potentially include code).
Internally, this uses the same encoding mechanism as the [method @GlobalScope.var_to_bytes] method.
*/
func (self Go) PutVar(value gd.Variant) {
	class(self).PutVar(value, false)
}

/*
Gets a signed byte from the stream.
*/
func (self Go) Get8() int {
	return int(int(class(self).Get8()))
}

/*
Gets an unsigned byte from the stream.
*/
func (self Go) GetU8() int {
	return int(int(class(self).GetU8()))
}

/*
Gets a signed 16-bit value from the stream.
*/
func (self Go) Get16() int {
	return int(int(class(self).Get16()))
}

/*
Gets an unsigned 16-bit value from the stream.
*/
func (self Go) GetU16() int {
	return int(int(class(self).GetU16()))
}

/*
Gets a signed 32-bit value from the stream.
*/
func (self Go) Get32() int {
	return int(int(class(self).Get32()))
}

/*
Gets an unsigned 32-bit value from the stream.
*/
func (self Go) GetU32() int {
	return int(int(class(self).GetU32()))
}

/*
Gets a signed 64-bit value from the stream.
*/
func (self Go) Get64() int {
	return int(int(class(self).Get64()))
}

/*
Gets an unsigned 64-bit value from the stream.
*/
func (self Go) GetU64() int {
	return int(int(class(self).GetU64()))
}

/*
Gets a single-precision float from the stream.
*/
func (self Go) GetFloat() float64 {
	return float64(float64(class(self).GetFloat()))
}

/*
Gets a double-precision float from the stream.
*/
func (self Go) GetDouble() float64 {
	return float64(float64(class(self).GetDouble()))
}

/*
Gets an ASCII string with byte-length [param bytes] from the stream. If [param bytes] is negative (default) the length will be read from the stream using the reverse process of [method put_string].
*/
func (self Go) GetString() string {
	return string(class(self).GetString(gd.Int(-1)).String())
}

/*
Gets a UTF-8 string with byte-length [param bytes] from the stream (this decodes the string sent as UTF-8). If [param bytes] is negative (default) the length will be read from the stream using the reverse process of [method put_utf8_string].
*/
func (self Go) GetUtf8String() string {
	return string(class(self).GetUtf8String(gd.Int(-1)).String())
}

/*
Gets a Variant from the stream. If [param allow_objects] is [code]true[/code], decoding objects is allowed.
Internally, this uses the same decoding mechanism as the [method @GlobalScope.bytes_to_var] method.
[b]Warning:[/b] Deserialized objects can contain code which gets executed. Do not use this option if the serialized object comes from untrusted sources to avoid potential security threats such as remote code execution.
*/
func (self Go) GetVar() gd.Variant {
	return gd.Variant(class(self).GetVar(false))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.StreamPeer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("StreamPeer"))
	return Go{classdb.StreamPeer(object)}
}

func (self Go) BigEndian() bool {
		return bool(class(self).IsBigEndianEnabled())
}

func (self Go) SetBigEndian(value bool) {
	class(self).SetBigEndian(value)
}

/*
Sends a chunk of data through the connection, blocking if necessary until the data is done sending. This function returns an [enum Error] code.
*/
//go:nosplit
func (self class) PutData(data gd.PackedByteArray) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(data))
	var r_ret = callframe.Ret[int64](frame)
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
	callframe.Arg(frame, discreet.Get(data))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_put_partial_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
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
	var ret = discreet.New[gd.Array](r_ret.Get())
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
	var ret = discreet.New[gd.Array](r_ret.Get())
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
func (self class) SetBigEndian(enable bool)  {
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
func (self class) Put8(value gd.Int)  {
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
func (self class) PutU8(value gd.Int)  {
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
func (self class) Put16(value gd.Int)  {
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
func (self class) PutU16(value gd.Int)  {
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
func (self class) Put32(value gd.Int)  {
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
func (self class) PutU32(value gd.Int)  {
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
func (self class) Put64(value gd.Int)  {
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
func (self class) PutU64(value gd.Int)  {
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
func (self class) PutFloat(value gd.Float)  {
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
func (self class) PutDouble(value gd.Float)  {
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
func (self class) PutString(value gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(value))
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
func (self class) PutUtf8String(value gd.String)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(value))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeer.Bind_put_utf8_string, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Puts a Variant into the stream. If [param full_objects] is [code]true[/code] encoding objects is allowed (and can potentially include code).
Internally, this uses the same encoding mechanism as the [method @GlobalScope.var_to_bytes] method.
*/
//go:nosplit
func (self class) PutVar(value gd.Variant, full_objects bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(value))
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
	var ret = discreet.New[gd.String](r_ret.Get())
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
	var ret = discreet.New[gd.String](r_ret.Get())
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
	var ret = discreet.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsStreamPeer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsStreamPeer() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("StreamPeer", func(ptr gd.Object) any { return classdb.StreamPeer(ptr) })}
