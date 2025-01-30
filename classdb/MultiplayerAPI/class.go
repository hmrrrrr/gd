// Package MultiplayerAPI provides methods for working with MultiplayerAPI object instances.
package MultiplayerAPI

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
Base class for high-level multiplayer API implementations. See also [MultiplayerPeer].
By default, [SceneTree] has a reference to an implementation of this class and uses it to provide multiplayer capabilities (i.e. RPCs) across the whole scene.
It is possible to override the MultiplayerAPI instance used by specific tree branches by calling the [method SceneTree.set_multiplayer] method, effectively allowing to run both client and server in the same scene.
It is also possible to extend or replace the default implementation via scripting or native extensions. See [MultiplayerAPIExtension] for details about extensions, [SceneMultiplayer] for the details about the default implementation.
*/
type Instance [1]gdclass.MultiplayerAPI

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsMultiplayerAPI() Instance
}

/*
Returns [code]true[/code] if there is a [member multiplayer_peer] set.
*/
func (self Instance) HasMultiplayerPeer() bool { //gd:MultiplayerAPI.has_multiplayer_peer
	return bool(class(self).HasMultiplayerPeer())
}

/*
Returns the unique peer ID of this MultiplayerAPI's [member multiplayer_peer].
*/
func (self Instance) GetUniqueId() int { //gd:MultiplayerAPI.get_unique_id
	return int(int(class(self).GetUniqueId()))
}

/*
Returns [code]true[/code] if this MultiplayerAPI's [member multiplayer_peer] is valid and in server mode (listening for connections).
*/
func (self Instance) IsServer() bool { //gd:MultiplayerAPI.is_server
	return bool(class(self).IsServer())
}

/*
Returns the sender's peer ID for the RPC currently being executed.
[b]Note:[/b] This method returns [code]0[/code] when called outside of an RPC. As such, the original peer ID may be lost when code execution is delayed (such as with GDScript's [code]await[/code] keyword).
*/
func (self Instance) GetRemoteSenderId() int { //gd:MultiplayerAPI.get_remote_sender_id
	return int(int(class(self).GetRemoteSenderId()))
}

/*
Method used for polling the MultiplayerAPI. You only need to worry about this if you set [member SceneTree.multiplayer_poll] to [code]false[/code]. By default, [SceneTree] will poll its MultiplayerAPI(s) for you.
[b]Note:[/b] This method results in RPCs being called, so they will be executed in the same context of this function (e.g. [code]_process[/code], [code]physics[/code], [Thread]).
*/
func (self Instance) Poll() error { //gd:MultiplayerAPI.poll
	return error(gd.ToError(class(self).Poll()))
}

/*
Sends an RPC to the target [param peer]. The given [param method] will be called on the remote [param object] with the provided [param arguments]. The RPC may also be called locally depending on the implementation and RPC configuration. See [method Node.rpc] and [method Node.rpc_config].
[b]Note:[/b] Prefer using [method Node.rpc], [method Node.rpc_id], or [code]my_method.rpc(peer, arg1, arg2, ...)[/code] (in GDScript), since they are faster. This method is mostly useful in conjunction with [MultiplayerAPIExtension] when augmenting or replacing the multiplayer capabilities.
*/
func (self Instance) Rpc(peer int, obj Object.Instance, method string) error { //gd:MultiplayerAPI.rpc
	return error(gd.ToError(class(self).Rpc(int64(peer), obj, String.Name(String.New(method)), Array.Nil)))
}

/*
Notifies the MultiplayerAPI of a new [param configuration] for the given [param object]. This method is used internally by [SceneTree] to configure the root path for this MultiplayerAPI (passing [code]null[/code] and a valid [NodePath] as [param configuration]). This method can be further used by MultiplayerAPI implementations to provide additional features, refer to specific implementation (e.g. [SceneMultiplayer]) for details on how they use it.
[b]Note:[/b] This method is mostly relevant when extending or overriding the MultiplayerAPI behavior via [MultiplayerAPIExtension].
*/
func (self Instance) ObjectConfigurationAdd(obj Object.Instance, configuration any) error { //gd:MultiplayerAPI.object_configuration_add
	return error(gd.ToError(class(self).ObjectConfigurationAdd(obj, variant.New(configuration))))
}

/*
Notifies the MultiplayerAPI to remove a [param configuration] for the given [param object]. This method is used internally by [SceneTree] to configure the root path for this MultiplayerAPI (passing [code]null[/code] and an empty [NodePath] as [param configuration]). This method can be further used by MultiplayerAPI implementations to provide additional features, refer to specific implementation (e.g. [SceneMultiplayer]) for details on how they use it.
[b]Note:[/b] This method is mostly relevant when extending or overriding the MultiplayerAPI behavior via [MultiplayerAPIExtension].
*/
func (self Instance) ObjectConfigurationRemove(obj Object.Instance, configuration any) error { //gd:MultiplayerAPI.object_configuration_remove
	return error(gd.ToError(class(self).ObjectConfigurationRemove(obj, variant.New(configuration))))
}

/*
Returns the peer IDs of all connected peers of this MultiplayerAPI's [member multiplayer_peer].
*/
func (self Instance) GetPeers() []int32 { //gd:MultiplayerAPI.get_peers
	return []int32(slices.Collect(class(self).GetPeers().Values()))
}

/*
Sets the default MultiplayerAPI implementation class. This method can be used by modules and extensions to configure which implementation will be used by [SceneTree] when the engine starts.
*/
func SetDefaultInterface(interface_name string) { //gd:MultiplayerAPI.set_default_interface
	self := Instance{}
	class(self).SetDefaultInterface(String.Name(String.New(interface_name)))
}

/*
Returns the default MultiplayerAPI implementation class name. This is usually [code]"SceneMultiplayer"[/code] when [SceneMultiplayer] is available. See [method set_default_interface].
*/
func GetDefaultInterface() string { //gd:MultiplayerAPI.get_default_interface
	self := Instance{}
	return string(class(self).GetDefaultInterface().String())
}

/*
Returns a new instance of the default MultiplayerAPI.
*/
func CreateDefaultInterface() [1]gdclass.MultiplayerAPI { //gd:MultiplayerAPI.create_default_interface
	self := Instance{}
	return [1]gdclass.MultiplayerAPI(class(self).CreateDefaultInterface())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.MultiplayerAPI

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MultiplayerAPI"))
	casted := Instance{*(*gdclass.MultiplayerAPI)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) MultiplayerPeer() [1]gdclass.MultiplayerPeer {
	return [1]gdclass.MultiplayerPeer(class(self).GetMultiplayerPeer())
}

func (self Instance) SetMultiplayerPeer(value [1]gdclass.MultiplayerPeer) {
	class(self).SetMultiplayerPeer(value)
}

/*
Returns [code]true[/code] if there is a [member multiplayer_peer] set.
*/
//go:nosplit
func (self class) HasMultiplayerPeer() bool { //gd:MultiplayerAPI.has_multiplayer_peer
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_has_multiplayer_peer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetMultiplayerPeer() [1]gdclass.MultiplayerPeer { //gd:MultiplayerAPI.get_multiplayer_peer
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_get_multiplayer_peer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.MultiplayerPeer{gd.PointerWithOwnershipTransferredToGo[gdclass.MultiplayerPeer](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMultiplayerPeer(peer [1]gdclass.MultiplayerPeer) { //gd:MultiplayerAPI.set_multiplayer_peer
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(peer[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_set_multiplayer_peer, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the unique peer ID of this MultiplayerAPI's [member multiplayer_peer].
*/
//go:nosplit
func (self class) GetUniqueId() int64 { //gd:MultiplayerAPI.get_unique_id
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_get_unique_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if this MultiplayerAPI's [member multiplayer_peer] is valid and in server mode (listening for connections).
*/
//go:nosplit
func (self class) IsServer() bool { //gd:MultiplayerAPI.is_server
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_is_server, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the sender's peer ID for the RPC currently being executed.
[b]Note:[/b] This method returns [code]0[/code] when called outside of an RPC. As such, the original peer ID may be lost when code execution is delayed (such as with GDScript's [code]await[/code] keyword).
*/
//go:nosplit
func (self class) GetRemoteSenderId() int64 { //gd:MultiplayerAPI.get_remote_sender_id
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_get_remote_sender_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Method used for polling the MultiplayerAPI. You only need to worry about this if you set [member SceneTree.multiplayer_poll] to [code]false[/code]. By default, [SceneTree] will poll its MultiplayerAPI(s) for you.
[b]Note:[/b] This method results in RPCs being called, so they will be executed in the same context of this function (e.g. [code]_process[/code], [code]physics[/code], [Thread]).
*/
//go:nosplit
func (self class) Poll() Error.Code { //gd:MultiplayerAPI.poll
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_poll, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Sends an RPC to the target [param peer]. The given [param method] will be called on the remote [param object] with the provided [param arguments]. The RPC may also be called locally depending on the implementation and RPC configuration. See [method Node.rpc] and [method Node.rpc_config].
[b]Note:[/b] Prefer using [method Node.rpc], [method Node.rpc_id], or [code]my_method.rpc(peer, arg1, arg2, ...)[/code] (in GDScript), since they are faster. This method is mostly useful in conjunction with [MultiplayerAPIExtension] when augmenting or replacing the multiplayer capabilities.
*/
//go:nosplit
func (self class) Rpc(peer int64, obj [1]gd.Object, method String.Name, arguments Array.Any) Error.Code { //gd:MultiplayerAPI.rpc
	var frame = callframe.New()
	callframe.Arg(frame, peer)
	callframe.Arg(frame, pointers.Get(obj[0])[0])
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(method)))
	callframe.Arg(frame, pointers.Get(gd.InternalArray(arguments)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_rpc, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Notifies the MultiplayerAPI of a new [param configuration] for the given [param object]. This method is used internally by [SceneTree] to configure the root path for this MultiplayerAPI (passing [code]null[/code] and a valid [NodePath] as [param configuration]). This method can be further used by MultiplayerAPI implementations to provide additional features, refer to specific implementation (e.g. [SceneMultiplayer]) for details on how they use it.
[b]Note:[/b] This method is mostly relevant when extending or overriding the MultiplayerAPI behavior via [MultiplayerAPIExtension].
*/
//go:nosplit
func (self class) ObjectConfigurationAdd(obj [1]gd.Object, configuration variant.Any) Error.Code { //gd:MultiplayerAPI.object_configuration_add
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(obj[0].AsObject()[0]))
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(configuration)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_object_configuration_add, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Notifies the MultiplayerAPI to remove a [param configuration] for the given [param object]. This method is used internally by [SceneTree] to configure the root path for this MultiplayerAPI (passing [code]null[/code] and an empty [NodePath] as [param configuration]). This method can be further used by MultiplayerAPI implementations to provide additional features, refer to specific implementation (e.g. [SceneMultiplayer]) for details on how they use it.
[b]Note:[/b] This method is mostly relevant when extending or overriding the MultiplayerAPI behavior via [MultiplayerAPIExtension].
*/
//go:nosplit
func (self class) ObjectConfigurationRemove(obj [1]gd.Object, configuration variant.Any) Error.Code { //gd:MultiplayerAPI.object_configuration_remove
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(obj[0])[0])
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(configuration)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_object_configuration_remove, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the peer IDs of all connected peers of this MultiplayerAPI's [member multiplayer_peer].
*/
//go:nosplit
func (self class) GetPeers() Packed.Array[int32] { //gd:MultiplayerAPI.get_peers
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_get_peers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int32](Array.Through(gd.PackedProxy[gd.PackedInt32Array, int32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Sets the default MultiplayerAPI implementation class. This method can be used by modules and extensions to configure which implementation will be used by [SceneTree] when the engine starts.
*/
//go:nosplit
func (self class) SetDefaultInterface(interface_name String.Name) { //gd:MultiplayerAPI.set_default_interface
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(interface_name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_set_default_interface, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the default MultiplayerAPI implementation class name. This is usually [code]"SceneMultiplayer"[/code] when [SceneMultiplayer] is available. See [method set_default_interface].
*/
//go:nosplit
func (self class) GetDefaultInterface() String.Name { //gd:MultiplayerAPI.get_default_interface
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_get_default_interface, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns a new instance of the default MultiplayerAPI.
*/
//go:nosplit
func (self class) CreateDefaultInterface() [1]gdclass.MultiplayerAPI { //gd:MultiplayerAPI.create_default_interface
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_create_default_interface, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.MultiplayerAPI{gd.PointerWithOwnershipTransferredToGo[gdclass.MultiplayerAPI](r_ret.Get())}
	frame.Free()
	return ret
}
func (self Instance) OnPeerConnected(cb func(id int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("peer_connected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPeerDisconnected(cb func(id int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("peer_disconnected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnConnectedToServer(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("connected_to_server"), gd.NewCallable(cb), 0)
}

func (self Instance) OnConnectionFailed(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("connection_failed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnServerDisconnected(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("server_disconnected"), gd.NewCallable(cb), 0)
}

func (self class) AsMultiplayerAPI() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMultiplayerAPI() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("MultiplayerAPI", func(ptr gd.Object) any {
		return [1]gdclass.MultiplayerAPI{*(*gdclass.MultiplayerAPI)(unsafe.Pointer(&ptr))}
	})
}

type RPCMode = gdclass.MultiplayerAPIRPCMode //gd:MultiplayerAPI.RPCMode

const (
	/*Used with [method Node.rpc_config] to disable a method or property for all RPC calls, making it unavailable. Default for all methods.*/
	RpcModeDisabled RPCMode = 0
	/*Used with [method Node.rpc_config] to set a method to be callable remotely by any peer. Analogous to the [code]@rpc("any_peer")[/code] annotation. Calls are accepted from all remote peers, no matter if they are node's authority or not.*/
	RpcModeAnyPeer RPCMode = 1
	/*Used with [method Node.rpc_config] to set a method to be callable remotely only by the current multiplayer authority (which is the server by default). Analogous to the [code]@rpc("authority")[/code] annotation. See [method Node.set_multiplayer_authority].*/
	RpcModeAuthority RPCMode = 2
)
