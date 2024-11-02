package StreamPeerTCP

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/StreamPeer"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A stream peer that handles TCP connections. This object can be used to connect to TCP servers, or also is returned by a TCP server.
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.
*/
type Instance [1]classdb.StreamPeerTCP

/*
Opens the TCP socket, and binds it to the specified local address.
This method is generally not needed, and only used to force the subsequent call to [method connect_to_host] to use the specified [param host] and [param port] as source address. This can be desired in some NAT punchthrough techniques, or when forcing the source network interface.
*/
func (self Instance) Bind(port int) error {
	return error(class(self).Bind(gd.Int(port), gd.NewString("*")))
}

/*
Connects to the specified [code]host:port[/code] pair. A hostname will be resolved if valid. Returns [constant OK] on success.
*/
func (self Instance) ConnectToHost(host string, port int) error {
	return error(class(self).ConnectToHost(gd.NewString(host), gd.Int(port)))
}

/*
Poll the socket, updating its state. See [method get_status].
*/
func (self Instance) Poll() error {
	return error(class(self).Poll())
}

/*
Returns the status of the connection, see [enum Status].
*/
func (self Instance) GetStatus() classdb.StreamPeerTCPStatus {
	return classdb.StreamPeerTCPStatus(class(self).GetStatus())
}

/*
Returns the IP of this peer.
*/
func (self Instance) GetConnectedHost() string {
	return string(class(self).GetConnectedHost().String())
}

/*
Returns the port of this peer.
*/
func (self Instance) GetConnectedPort() int {
	return int(int(class(self).GetConnectedPort()))
}

/*
Returns the local port to which this peer is bound.
*/
func (self Instance) GetLocalPort() int {
	return int(int(class(self).GetLocalPort()))
}

/*
Disconnects from host.
*/
func (self Instance) DisconnectFromHost() {
	class(self).DisconnectFromHost()
}

/*
If [param enabled] is [code]true[/code], packets will be sent immediately. If [param enabled] is [code]false[/code] (the default), packet transfers will be delayed and combined using [url=https://en.wikipedia.org/wiki/Nagle%27s_algorithm]Nagle's algorithm[/url].
[b]Note:[/b] It's recommended to leave this disabled for applications that send large packets or need to transfer a lot of data, as enabling this can decrease the total available bandwidth.
*/
func (self Instance) SetNoDelay(enabled bool) {
	class(self).SetNoDelay(enabled)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.StreamPeerTCP

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("StreamPeerTCP"))
	return Instance{classdb.StreamPeerTCP(object)}
}

/*
Opens the TCP socket, and binds it to the specified local address.
This method is generally not needed, and only used to force the subsequent call to [method connect_to_host] to use the specified [param host] and [param port] as source address. This can be desired in some NAT punchthrough techniques, or when forcing the source network interface.
*/
//go:nosplit
func (self class) Bind(port gd.Int, host gd.String) error {
	var frame = callframe.New()
	callframe.Arg(frame, port)
	callframe.Arg(frame, pointers.Get(host))
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerTCP.Bind_bind, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Connects to the specified [code]host:port[/code] pair. A hostname will be resolved if valid. Returns [constant OK] on success.
*/
//go:nosplit
func (self class) ConnectToHost(host gd.String, port gd.Int) error {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(host))
	callframe.Arg(frame, port)
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerTCP.Bind_connect_to_host, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Poll the socket, updating its state. See [method get_status].
*/
//go:nosplit
func (self class) Poll() error {
	var frame = callframe.New()
	var r_ret = callframe.Ret[error](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerTCP.Bind_poll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the status of the connection, see [enum Status].
*/
//go:nosplit
func (self class) GetStatus() classdb.StreamPeerTCPStatus {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.StreamPeerTCPStatus](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerTCP.Bind_get_status, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the IP of this peer.
*/
//go:nosplit
func (self class) GetConnectedHost() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerTCP.Bind_get_connected_host, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the port of this peer.
*/
//go:nosplit
func (self class) GetConnectedPort() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerTCP.Bind_get_connected_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the local port to which this peer is bound.
*/
//go:nosplit
func (self class) GetLocalPort() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerTCP.Bind_get_local_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Disconnects from host.
*/
//go:nosplit
func (self class) DisconnectFromHost() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerTCP.Bind_disconnect_from_host, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
If [param enabled] is [code]true[/code], packets will be sent immediately. If [param enabled] is [code]false[/code] (the default), packet transfers will be delayed and combined using [url=https://en.wikipedia.org/wiki/Nagle%27s_algorithm]Nagle's algorithm[/url].
[b]Note:[/b] It's recommended to leave this disabled for applications that send large packets or need to transfer a lot of data, as enabling this can decrease the total available bandwidth.
*/
//go:nosplit
func (self class) SetNoDelay(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.StreamPeerTCP.Bind_set_no_delay, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsStreamPeerTCP() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsStreamPeerTCP() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsStreamPeer() StreamPeer.Advanced {
	return *((*StreamPeer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsStreamPeer() StreamPeer.Instance {
	return *((*StreamPeer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsStreamPeer(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsStreamPeer(), name)
	}
}
func init() {
	classdb.Register("StreamPeerTCP", func(ptr gd.Object) any { return classdb.StreamPeerTCP(ptr) })
}

type Status = classdb.StreamPeerTCPStatus

const (
	/*The initial status of the [StreamPeerTCP]. This is also the status after disconnecting.*/
	StatusNone Status = 0
	/*A status representing a [StreamPeerTCP] that is connecting to a host.*/
	StatusConnecting Status = 1
	/*A status representing a [StreamPeerTCP] that is connected to a host.*/
	StatusConnected Status = 2
	/*A status representing a [StreamPeerTCP] in error state.*/
	StatusError Status = 3
)

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
