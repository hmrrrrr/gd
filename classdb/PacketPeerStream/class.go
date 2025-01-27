// Package PacketPeerStream provides methods for working with PacketPeerStream object instances.
package PacketPeerStream

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/Packed"
import "graphics.gd/classdb/PacketPeer"

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
var _ = slices.Delete[[]struct{}, struct{}]

/*
PacketStreamPeer provides a wrapper for working using packets over a stream. This allows for using packet based code with StreamPeers. PacketPeerStream implements a custom protocol over the StreamPeer, so the user should not read or write to the wrapped StreamPeer directly.
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.
*/
type Instance [1]gdclass.PacketPeerStream

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPacketPeerStream() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PacketPeerStream

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PacketPeerStream"))
	casted := Instance{*(*gdclass.PacketPeerStream)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) InputBufferMaxSize() int {
	return int(int(class(self).GetInputBufferMaxSize()))
}

func (self Instance) SetInputBufferMaxSize(value int) {
	class(self).SetInputBufferMaxSize(gd.Int(value))
}

func (self Instance) OutputBufferMaxSize() int {
	return int(int(class(self).GetOutputBufferMaxSize()))
}

func (self Instance) SetOutputBufferMaxSize(value int) {
	class(self).SetOutputBufferMaxSize(gd.Int(value))
}

func (self Instance) StreamPeer() [1]gdclass.StreamPeer {
	return [1]gdclass.StreamPeer(class(self).GetStreamPeer())
}

func (self Instance) SetStreamPeer(value [1]gdclass.StreamPeer) {
	class(self).SetStreamPeer(value)
}

//go:nosplit
func (self class) SetStreamPeer(peer [1]gdclass.StreamPeer) { //gd:PacketPeerStream.set_stream_peer
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(peer[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeerStream.Bind_set_stream_peer, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStreamPeer() [1]gdclass.StreamPeer { //gd:PacketPeerStream.get_stream_peer
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeerStream.Bind_get_stream_peer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.StreamPeer{gd.PointerWithOwnershipTransferredToGo[gdclass.StreamPeer](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInputBufferMaxSize(max_size_bytes gd.Int) { //gd:PacketPeerStream.set_input_buffer_max_size
	var frame = callframe.New()
	callframe.Arg(frame, max_size_bytes)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeerStream.Bind_set_input_buffer_max_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetOutputBufferMaxSize(max_size_bytes gd.Int) { //gd:PacketPeerStream.set_output_buffer_max_size
	var frame = callframe.New()
	callframe.Arg(frame, max_size_bytes)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeerStream.Bind_set_output_buffer_max_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetInputBufferMaxSize() gd.Int { //gd:PacketPeerStream.get_input_buffer_max_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeerStream.Bind_get_input_buffer_max_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetOutputBufferMaxSize() gd.Int { //gd:PacketPeerStream.get_output_buffer_max_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PacketPeerStream.Bind_get_output_buffer_max_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPacketPeerStream() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPacketPeerStream() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsPacketPeer() PacketPeer.Advanced {
	return *((*PacketPeer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPacketPeer() PacketPeer.Instance {
	return *((*PacketPeer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(PacketPeer.Advanced(self.AsPacketPeer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(PacketPeer.Instance(self.AsPacketPeer()), name)
	}
}
func init() {
	gdclass.Register("PacketPeerStream", func(ptr gd.Object) any {
		return [1]gdclass.PacketPeerStream{*(*gdclass.PacketPeerStream)(unsafe.Pointer(&ptr))}
	})
}
