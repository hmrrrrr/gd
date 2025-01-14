// Package StreamPeerTLS provides methods for working with StreamPeerTLS object instances.
package StreamPeerTLS

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/StreamPeer"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

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
func (self Instance) Poll() {
	class(self).Poll()
}

/*
Accepts a peer connection as a server using the given [param server_options]. See [method TLSOptions.server].
*/
func (self Instance) AcceptStream(stream [1]gdclass.StreamPeer, server_options [1]gdclass.TLSOptions) error {
	return error(class(self).AcceptStream(stream, server_options))
}

/*
Connects to a peer using an underlying [StreamPeer] [param stream] and verifying the remote certificate is correctly signed for the given [param common_name]. You can pass the optional [param client_options] parameter to customize the trusted certification authorities, or disable the common name verification. See [method TLSOptions.client] and [method TLSOptions.client_unsafe].
*/
func (self Instance) ConnectToStream(stream [1]gdclass.StreamPeer, common_name string) error {
	return error(class(self).ConnectToStream(stream, gd.NewString(common_name), [1][1]gdclass.TLSOptions{}[0]))
}

/*
Returns the status of the connection. See [enum Status] for values.
*/
func (self Instance) GetStatus() gdclass.StreamPeerTLSStatus {
	return gdclass.StreamPeerTLSStatus(class(self).GetStatus())
}

/*
Returns the underlying [StreamPeer] connection, used in [method accept_stream] or [method connect_to_stream].
*/
func (self Instance) GetStream() [1]gdclass.StreamPeer {
	return [1]gdclass.StreamPeer(class(self).GetStream())
}

/*
Disconnects from host.
*/
func (self Instance) DisconnectFromStream() {
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
func (self class) Poll() {
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerTLS.Bind_poll, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Accepts a peer connection as a server using the given [param server_options]. See [method TLSOptions.server].
*/
//go:nosplit
func (self class) AcceptStream(stream [1]gdclass.StreamPeer, server_options [1]gdclass.TLSOptions) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(stream[0])[0])
	callframe.Arg(frame, pointers.Get(server_options[0])[0])
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerTLS.Bind_accept_stream, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Connects to a peer using an underlying [StreamPeer] [param stream] and verifying the remote certificate is correctly signed for the given [param common_name]. You can pass the optional [param client_options] parameter to customize the trusted certification authorities, or disable the common name verification. See [method TLSOptions.client] and [method TLSOptions.client_unsafe].
*/
//go:nosplit
func (self class) ConnectToStream(stream [1]gdclass.StreamPeer, common_name gd.String, client_options [1]gdclass.TLSOptions) gd.Error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(stream[0])[0])
	callframe.Arg(frame, pointers.Get(common_name))
	callframe.Arg(frame, pointers.Get(client_options[0])[0])
	var r_ret = callframe.Ret[gd.Error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerTLS.Bind_connect_to_stream, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the status of the connection. See [enum Status] for values.
*/
//go:nosplit
func (self class) GetStatus() gdclass.StreamPeerTLSStatus {
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
func (self class) GetStream() [1]gdclass.StreamPeer {
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
func (self class) DisconnectFromStream() {
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

type Status = gdclass.StreamPeerTLSStatus

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

type Error = gd.Error

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
