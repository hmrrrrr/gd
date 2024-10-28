package MultiplayerAPI

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
Base class for high-level multiplayer API implementations. See also [MultiplayerPeer].
By default, [SceneTree] has a reference to an implementation of this class and uses it to provide multiplayer capabilities (i.e. RPCs) across the whole scene.
It is possible to override the MultiplayerAPI instance used by specific tree branches by calling the [method SceneTree.set_multiplayer] method, effectively allowing to run both client and server in the same scene.
It is also possible to extend or replace the default implementation via scripting or native extensions. See [MultiplayerAPIExtension] for details about extensions, [SceneMultiplayer] for the details about the default implementation.

*/
type Go [1]classdb.MultiplayerAPI

/*
Returns [code]true[/code] if there is a [member multiplayer_peer] set.
*/
func (self Go) HasMultiplayerPeer() bool {
	return bool(class(self).HasMultiplayerPeer())
}

/*
Returns the unique peer ID of this MultiplayerAPI's [member multiplayer_peer].
*/
func (self Go) GetUniqueId() int {
	return int(int(class(self).GetUniqueId()))
}

/*
Returns [code]true[/code] if this MultiplayerAPI's [member multiplayer_peer] is valid and in server mode (listening for connections).
*/
func (self Go) IsServer() bool {
	return bool(class(self).IsServer())
}

/*
Returns the sender's peer ID for the RPC currently being executed.
[b]Note:[/b] This method returns [code]0[/code] when called outside of an RPC. As such, the original peer ID may be lost when code execution is delayed (such as with GDScript's [code]await[/code] keyword).
*/
func (self Go) GetRemoteSenderId() int {
	return int(int(class(self).GetRemoteSenderId()))
}

/*
Method used for polling the MultiplayerAPI. You only need to worry about this if you set [member SceneTree.multiplayer_poll] to [code]false[/code]. By default, [SceneTree] will poll its MultiplayerAPI(s) for you.
[b]Note:[/b] This method results in RPCs being called, so they will be executed in the same context of this function (e.g. [code]_process[/code], [code]physics[/code], [Thread]).
*/
func (self Go) Poll() gd.Error {
	return gd.Error(class(self).Poll())
}

/*
Sends an RPC to the target [param peer]. The given [param method] will be called on the remote [param object] with the provided [param arguments]. The RPC may also be called locally depending on the implementation and RPC configuration. See [method Node.rpc] and [method Node.rpc_config].
[b]Note:[/b] Prefer using [method Node.rpc], [method Node.rpc_id], or [code]my_method.rpc(peer, arg1, arg2, ...)[/code] (in GDScript), since they are faster. This method is mostly useful in conjunction with [MultiplayerAPIExtension] when augmenting or replacing the multiplayer capabilities.
*/
func (self Go) Rpc(peer int, obj gd.Object, method string) gd.Error {
	return gd.Error(class(self).Rpc(gd.Int(peer), obj, gd.NewStringName(method), ([1]gd.Array{}[0])))
}

/*
Notifies the MultiplayerAPI of a new [param configuration] for the given [param object]. This method is used internally by [SceneTree] to configure the root path for this MultiplayerAPI (passing [code]null[/code] and a valid [NodePath] as [param configuration]). This method can be further used by MultiplayerAPI implementations to provide additional features, refer to specific implementation (e.g. [SceneMultiplayer]) for details on how they use it.
[b]Note:[/b] This method is mostly relevant when extending or overriding the MultiplayerAPI behavior via [MultiplayerAPIExtension].
*/
func (self Go) ObjectConfigurationAdd(obj gd.Object, configuration gd.Variant) gd.Error {
	return gd.Error(class(self).ObjectConfigurationAdd(obj, configuration))
}

/*
Notifies the MultiplayerAPI to remove a [param configuration] for the given [param object]. This method is used internally by [SceneTree] to configure the root path for this MultiplayerAPI (passing [code]null[/code] and an empty [NodePath] as [param configuration]). This method can be further used by MultiplayerAPI implementations to provide additional features, refer to specific implementation (e.g. [SceneMultiplayer]) for details on how they use it.
[b]Note:[/b] This method is mostly relevant when extending or overriding the MultiplayerAPI behavior via [MultiplayerAPIExtension].
*/
func (self Go) ObjectConfigurationRemove(obj gd.Object, configuration gd.Variant) gd.Error {
	return gd.Error(class(self).ObjectConfigurationRemove(obj, configuration))
}

/*
Returns the peer IDs of all connected peers of this MultiplayerAPI's [member multiplayer_peer].
*/
func (self Go) GetPeers() []int32 {
	return []int32(class(self).GetPeers().AsSlice())
}

/*
Sets the default MultiplayerAPI implementation class. This method can be used by modules and extensions to configure which implementation will be used by [SceneTree] when the engine starts.
*/
func (self Go) SetDefaultInterface(interface_name string) {
	class(self).SetDefaultInterface(gd.NewStringName(interface_name))
}

/*
Returns the default MultiplayerAPI implementation class name. This is usually [code]"SceneMultiplayer"[/code] when [SceneMultiplayer] is available. See [method set_default_interface].
*/
func (self Go) GetDefaultInterface() string {
	return string(class(self).GetDefaultInterface().String())
}

/*
Returns a new instance of the default MultiplayerAPI.
*/
func (self Go) CreateDefaultInterface() gdclass.MultiplayerAPI {
	return gdclass.MultiplayerAPI(class(self).CreateDefaultInterface())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.MultiplayerAPI
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MultiplayerAPI"))
	return Go{classdb.MultiplayerAPI(object)}
}

func (self Go) MultiplayerPeer() gdclass.MultiplayerPeer {
		return gdclass.MultiplayerPeer(class(self).GetMultiplayerPeer())
}

func (self Go) SetMultiplayerPeer(value gdclass.MultiplayerPeer) {
	class(self).SetMultiplayerPeer(value)
}

/*
Returns [code]true[/code] if there is a [member multiplayer_peer] set.
*/
//go:nosplit
func (self class) HasMultiplayerPeer() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_has_multiplayer_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetMultiplayerPeer() gdclass.MultiplayerPeer {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_get_multiplayer_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.MultiplayerPeer{classdb.MultiplayerPeer(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMultiplayerPeer(peer gdclass.MultiplayerPeer)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(peer[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_set_multiplayer_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the unique peer ID of this MultiplayerAPI's [member multiplayer_peer].
*/
//go:nosplit
func (self class) GetUniqueId() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_get_unique_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if this MultiplayerAPI's [member multiplayer_peer] is valid and in server mode (listening for connections).
*/
//go:nosplit
func (self class) IsServer() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_is_server, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the sender's peer ID for the RPC currently being executed.
[b]Note:[/b] This method returns [code]0[/code] when called outside of an RPC. As such, the original peer ID may be lost when code execution is delayed (such as with GDScript's [code]await[/code] keyword).
*/
//go:nosplit
func (self class) GetRemoteSenderId() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_get_remote_sender_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Method used for polling the MultiplayerAPI. You only need to worry about this if you set [member SceneTree.multiplayer_poll] to [code]false[/code]. By default, [SceneTree] will poll its MultiplayerAPI(s) for you.
[b]Note:[/b] This method results in RPCs being called, so they will be executed in the same context of this function (e.g. [code]_process[/code], [code]physics[/code], [Thread]).
*/
//go:nosplit
func (self class) Poll() int64 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_poll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sends an RPC to the target [param peer]. The given [param method] will be called on the remote [param object] with the provided [param arguments]. The RPC may also be called locally depending on the implementation and RPC configuration. See [method Node.rpc] and [method Node.rpc_config].
[b]Note:[/b] Prefer using [method Node.rpc], [method Node.rpc_id], or [code]my_method.rpc(peer, arg1, arg2, ...)[/code] (in GDScript), since they are faster. This method is mostly useful in conjunction with [MultiplayerAPIExtension] when augmenting or replacing the multiplayer capabilities.
*/
//go:nosplit
func (self class) Rpc(peer gd.Int, obj gd.Object, method gd.StringName, arguments gd.Array) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, peer)
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(obj))
	callframe.Arg(frame, discreet.Get(method))
	callframe.Arg(frame, discreet.Get(arguments))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_rpc, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Notifies the MultiplayerAPI of a new [param configuration] for the given [param object]. This method is used internally by [SceneTree] to configure the root path for this MultiplayerAPI (passing [code]null[/code] and a valid [NodePath] as [param configuration]). This method can be further used by MultiplayerAPI implementations to provide additional features, refer to specific implementation (e.g. [SceneMultiplayer]) for details on how they use it.
[b]Note:[/b] This method is mostly relevant when extending or overriding the MultiplayerAPI behavior via [MultiplayerAPIExtension].
*/
//go:nosplit
func (self class) ObjectConfigurationAdd(obj gd.Object, configuration gd.Variant) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(obj))
	callframe.Arg(frame, discreet.Get(configuration))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_object_configuration_add, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Notifies the MultiplayerAPI to remove a [param configuration] for the given [param object]. This method is used internally by [SceneTree] to configure the root path for this MultiplayerAPI (passing [code]null[/code] and an empty [NodePath] as [param configuration]). This method can be further used by MultiplayerAPI implementations to provide additional features, refer to specific implementation (e.g. [SceneMultiplayer]) for details on how they use it.
[b]Note:[/b] This method is mostly relevant when extending or overriding the MultiplayerAPI behavior via [MultiplayerAPIExtension].
*/
//go:nosplit
func (self class) ObjectConfigurationRemove(obj gd.Object, configuration gd.Variant) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(obj))
	callframe.Arg(frame, discreet.Get(configuration))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_object_configuration_remove, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the peer IDs of all connected peers of this MultiplayerAPI's [member multiplayer_peer].
*/
//go:nosplit
func (self class) GetPeers() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_get_peers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the default MultiplayerAPI implementation class. This method can be used by modules and extensions to configure which implementation will be used by [SceneTree] when the engine starts.
*/
//go:nosplit
func (self class) SetDefaultInterface(interface_name gd.StringName)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(interface_name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_set_default_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the default MultiplayerAPI implementation class name. This is usually [code]"SceneMultiplayer"[/code] when [SceneMultiplayer] is available. See [method set_default_interface].
*/
//go:nosplit
func (self class) GetDefaultInterface() gd.StringName {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_get_default_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns a new instance of the default MultiplayerAPI.
*/
//go:nosplit
func (self class) CreateDefaultInterface() gdclass.MultiplayerAPI {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerAPI.Bind_create_default_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.MultiplayerAPI{classdb.MultiplayerAPI(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self Go) OnPeerConnected(cb func(id int)) {
	self[0].AsObject().Connect(gd.NewStringName("peer_connected"), gd.NewCallable(cb), 0)
}


func (self Go) OnPeerDisconnected(cb func(id int)) {
	self[0].AsObject().Connect(gd.NewStringName("peer_disconnected"), gd.NewCallable(cb), 0)
}


func (self Go) OnConnectedToServer(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("connected_to_server"), gd.NewCallable(cb), 0)
}


func (self Go) OnConnectionFailed(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("connection_failed"), gd.NewCallable(cb), 0)
}


func (self Go) OnServerDisconnected(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("server_disconnected"), gd.NewCallable(cb), 0)
}


func (self class) AsMultiplayerAPI() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsMultiplayerAPI() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("MultiplayerAPI", func(ptr gd.Object) any { return classdb.MultiplayerAPI(ptr) })}
type RPCMode = classdb.MultiplayerAPIRPCMode

const (
/*Used with [method Node.rpc_config] to disable a method or property for all RPC calls, making it unavailable. Default for all methods.*/
	RpcModeDisabled RPCMode = 0
/*Used with [method Node.rpc_config] to set a method to be callable remotely by any peer. Analogous to the [code]@rpc("any_peer")[/code] annotation. Calls are accepted from all remote peers, no matter if they are node's authority or not.*/
	RpcModeAnyPeer RPCMode = 1
/*Used with [method Node.rpc_config] to set a method to be callable remotely only by the current multiplayer authority (which is the server by default). Analogous to the [code]@rpc("authority")[/code] annotation. See [method Node.set_multiplayer_authority].*/
	RpcModeAuthority RPCMode = 2
)
