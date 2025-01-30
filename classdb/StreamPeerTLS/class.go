// Package StreamPeerTLS provides methods for working with StreamPeerTLS object instances.
package StreamPeerTLS

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/StreamPeer"
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
A stream peer that handles TLS connections. This object can be used to connect to a TLS server or accept a single TLS client connection.
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.
*/
type Instance [1]gdclass.StreamPeerTLS

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsStreamPeerTLS() Instance
}

/*
Poll the connection to check for incoming bytes. Call this right before [method StreamPeer.get_available_bytes] for it to work properly.
*/
func (self Instance) Poll() { //gd:StreamPeerTLS.poll
	class(self).Poll()
}

/*
Accepts a peer connection as a server using the given [param server_options]. See [method TLSOptions.server].
*/
func (self Instance) AcceptStream(stream [1]gdclass.StreamPeer, server_options [1]gdclass.TLSOptions) error { //gd:StreamPeerTLS.accept_stream
	return error(gd.ToError(class(self).AcceptStream(stream, server_options)))
}

/*
Connects to a peer using an underlying [StreamPeer] [param stream] and verifying the remote certificate is correctly signed for the given [param common_name]. You can pass the optional [param client_options] parameter to customize the trusted certification authorities, or disable the common name verification. See [method TLSOptions.client] and [method TLSOptions.client_unsafe].
*/
func (self Instance) ConnectToStream(stream [1]gdclass.StreamPeer, common_name string) error { //gd:StreamPeerTLS.connect_to_stream
	return error(gd.ToError(class(self).ConnectToStream(stream, String.New(common_name), [1][1]gdclass.TLSOptions{}[0])))
}

/*
Returns the status of the connection. See [enum Status] for values.
*/
func (self Instance) GetStatus() gdclass.StreamPeerTLSStatus { //gd:StreamPeerTLS.get_status
	return gdclass.StreamPeerTLSStatus(class(self).GetStatus())
}

/*
Returns the underlying [StreamPeer] connection, used in [method accept_stream] or [method connect_to_stream].
*/
func (self Instance) GetStream() [1]gdclass.StreamPeer { //gd:StreamPeerTLS.get_stream
	return [1]gdclass.StreamPeer(class(self).GetStream())
}

/*
Disconnects from host.
*/
func (self Instance) DisconnectFromStream() { //gd:StreamPeerTLS.disconnect_from_stream
	class(self).DisconnectFromStream()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.StreamPeerTLS

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("StreamPeerTLS"))
	casted := Instance{*(*gdclass.StreamPeerTLS)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Poll the connection to check for incoming bytes. Call this right before [method StreamPeer.get_available_bytes] for it to work properly.
*/
//go:nosplit
func (self class) Poll() { //gd:StreamPeerTLS.poll
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerTLS.Bind_poll, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Accepts a peer connection as a server using the given [param server_options]. See [method TLSOptions.server].
*/
//go:nosplit
func (self class) AcceptStream(stream [1]gdclass.StreamPeer, server_options [1]gdclass.TLSOptions) Error.Code { //gd:StreamPeerTLS.accept_stream
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(stream[0])[0])
	callframe.Arg(frame, pointers.Get(server_options[0])[0])
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerTLS.Bind_accept_stream, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Connects to a peer using an underlying [StreamPeer] [param stream] and verifying the remote certificate is correctly signed for the given [param common_name]. You can pass the optional [param client_options] parameter to customize the trusted certification authorities, or disable the common name verification. See [method TLSOptions.client] and [method TLSOptions.client_unsafe].
*/
//go:nosplit
func (self class) ConnectToStream(stream [1]gdclass.StreamPeer, common_name String.Readable, client_options [1]gdclass.TLSOptions) Error.Code { //gd:StreamPeerTLS.connect_to_stream
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(stream[0])[0])
	callframe.Arg(frame, pointers.Get(gd.InternalString(common_name)))
	callframe.Arg(frame, pointers.Get(client_options[0])[0])
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerTLS.Bind_connect_to_stream, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the status of the connection. See [enum Status] for values.
*/
//go:nosplit
func (self class) GetStatus() gdclass.StreamPeerTLSStatus { //gd:StreamPeerTLS.get_status
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.StreamPeerTLSStatus](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerTLS.Bind_get_status, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the underlying [StreamPeer] connection, used in [method accept_stream] or [method connect_to_stream].
*/
//go:nosplit
func (self class) GetStream() [1]gdclass.StreamPeer { //gd:StreamPeerTLS.get_stream
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerTLS.Bind_get_stream, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.StreamPeer{gd.PointerWithOwnershipTransferredToGo[gdclass.StreamPeer](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Disconnects from host.
*/
//go:nosplit
func (self class) DisconnectFromStream() { //gd:StreamPeerTLS.disconnect_from_stream
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerTLS.Bind_disconnect_from_stream, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsStreamPeerTLS() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsStreamPeerTLS() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsStreamPeer() StreamPeer.Advanced {
	return *((*StreamPeer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsStreamPeer() StreamPeer.Instance {
	return *((*StreamPeer.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(StreamPeer.Advanced(self.AsStreamPeer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(StreamPeer.Instance(self.AsStreamPeer()), name)
	}
}
func init() {
	gdclass.Register("StreamPeerTLS", func(ptr gd.Object) any {
		return [1]gdclass.StreamPeerTLS{*(*gdclass.StreamPeerTLS)(unsafe.Pointer(&ptr))}
	})
}

type Status = gdclass.StreamPeerTLSStatus //gd:StreamPeerTLS.Status

const (
	/*A status representing a [StreamPeerTLS] that is disconnected.*/
	StatusDisconnected Status = 0
	/*A status representing a [StreamPeerTLS] during handshaking.*/
	StatusHandshaking Status = 1
	/*A status representing a [StreamPeerTLS] that is connected to a host.*/
	StatusConnected Status = 2
	/*A status representing a [StreamPeerTLS] in error state.*/
	StatusError Status = 3
	/*An error status that shows a mismatch in the TLS certificate domain presented by the host and the domain requested for validation.*/
	StatusErrorHostnameMismatch Status = 4
)
