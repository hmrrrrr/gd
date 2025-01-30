// Package PacketPeer provides methods for working with PacketPeer object instances.
package PacketPeer

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
PacketPeer is an abstraction and base class for packet-based protocols (such as UDP). It provides an API for sending and receiving packets both as raw data or variables. This makes it easy to transfer data over a protocol, without having to encode data as low-level bytes or having to worry about network ordering.
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.
*/
type Instance [1]gdclass.PacketPeer

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPacketPeer() Instance
}

/*
Gets a Variant. If [param allow_objects] is [code]true[/code], decoding objects is allowed.
Internally, this uses the same decoding mechanism as the [method @GlobalScope.bytes_to_var] method.
[b]Warning:[/b] Deserialized objects can contain code which gets executed. Do not use this option if the serialized object comes from untrusted sources to avoid potential security threats such as remote code execution.
*/
func (self Instance) GetVar() any { //gd:PacketPeer.get_var
	return any(class(self).GetVar(false).Interface())
}

/*
Sends a [Variant] as a packet. If [param full_objects] is [code]true[/code], encoding objects is allowed (and can potentially include code).
Internally, this uses the same encoding mechanism as the [method @GlobalScope.var_to_bytes] method.
*/
func (self Instance) PutVar(v any) error { //gd:PacketPeer.put_var
	return error(gd.ToError(class(self).PutVar(variant.New(v), false)))
}

/*
Gets a raw packet.
*/
func (self Instance) GetPacket() []byte { //gd:PacketPeer.get_packet
	return []byte(class(self).GetPacket().Bytes())
}

/*
Sends a raw packet.
*/
func (self Instance) PutPacket(buffer []byte) error { //gd:PacketPeer.put_packet
	return error(gd.ToError(class(self).PutPacket(Packed.Bytes(Packed.New(buffer...)))))
}

/*
Returns the error state of the last packet received (via [method get_packet] and [method get_var]).
*/
func (self Instance) GetPacketError() error { //gd:PacketPeer.get_packet_error
	return error(gd.ToError(class(self).GetPacketError()))
}

/*
Returns the number of packets currently available in the ring-buffer.
*/
func (self Instance) GetAvailablePacketCount() int { //gd:PacketPeer.get_available_packet_count
	return int(int(class(self).GetAvailablePacketCount()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PacketPeer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PacketPeer"))
	casted := Instance{*(*gdclass.PacketPeer)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) EncodeBufferMaxSize() int {
	return int(int(class(self).GetEncodeBufferMaxSize()))
}

func (self Instance) SetEncodeBufferMaxSize(value int) {
	class(self).SetEncodeBufferMaxSize(int64(value))
}

/*
Gets a Variant. If [param allow_objects] is [code]true[/code], decoding objects is allowed.
Internally, this uses the same decoding mechanism as the [method @GlobalScope.bytes_to_var] method.
[b]Warning:[/b] Deserialized objects can contain code which gets executed. Do not use this option if the serialized object comes from untrusted sources to avoid potential security threats such as remote code execution.
*/
//go:nosplit
func (self class) GetVar(allow_objects bool) variant.Any { //gd:PacketPeer.get_var
	var frame = callframe.New()
	callframe.Arg(frame, allow_objects)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeer.Bind_get_var, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Implementation(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sends a [Variant] as a packet. If [param full_objects] is [code]true[/code], encoding objects is allowed (and can potentially include code).
Internally, this uses the same encoding mechanism as the [method @GlobalScope.var_to_bytes] method.
*/
//go:nosplit
func (self class) PutVar(v variant.Any, full_objects bool) Error.Code { //gd:PacketPeer.put_var
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(v)))
	callframe.Arg(frame, full_objects)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeer.Bind_put_var, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Gets a raw packet.
*/
//go:nosplit
func (self class) GetPacket() Packed.Bytes { //gd:PacketPeer.get_packet
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeer.Bind_get_packet, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Bytes(Array.Through(gd.PackedProxy[gd.PackedByteArray, byte]{}, pointers.Pack(pointers.New[gd.PackedByteArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Sends a raw packet.
*/
//go:nosplit
func (self class) PutPacket(buffer Packed.Bytes) Error.Code { //gd:PacketPeer.put_packet
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedByteArray, byte](Packed.Array[byte](buffer))))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeer.Bind_put_packet, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the error state of the last packet received (via [method get_packet] and [method get_var]).
*/
//go:nosplit
func (self class) GetPacketError() Error.Code { //gd:PacketPeer.get_packet_error
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeer.Bind_get_packet_error, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the number of packets currently available in the ring-buffer.
*/
//go:nosplit
func (self class) GetAvailablePacketCount() int64 { //gd:PacketPeer.get_available_packet_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeer.Bind_get_available_packet_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetEncodeBufferMaxSize() int64 { //gd:PacketPeer.get_encode_buffer_max_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeer.Bind_get_encode_buffer_max_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEncodeBufferMaxSize(max_size int64) { //gd:PacketPeer.set_encode_buffer_max_size
	var frame = callframe.New()
	callframe.Arg(frame, max_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeer.Bind_set_encode_buffer_max_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsPacketPeer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPacketPeer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("PacketPeer", func(ptr gd.Object) any { return [1]gdclass.PacketPeer{*(*gdclass.PacketPeer)(unsafe.Pointer(&ptr))} })
}
