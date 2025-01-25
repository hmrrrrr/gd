// Package MultiplayerPeer provides methods for working with MultiplayerPeer object instances.
package MultiplayerPeer

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
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

/*
Manages the connection with one or more remote peers acting as server or client and assigning unique IDs to each of them. See also [MultiplayerAPI].
[b]Note:[/b] The [MultiplayerAPI] protocol is an implementation detail and isn't meant to be used by non-Godot servers. It may change without notice.
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.
*/
type Instance [1]gdclass.MultiplayerPeer

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsMultiplayerPeer() Instance
}

/*
Sets the peer to which packets will be sent.
The [param id] can be one of: [constant TARGET_PEER_BROADCAST] to send to all connected peers, [constant TARGET_PEER_SERVER] to send to the peer acting as server, a valid peer ID to send to that specific peer, a negative peer ID to send to all peers except that one. By default, the target peer is [constant TARGET_PEER_BROADCAST].
*/
func (self Instance) SetTargetPeer(id int) { //gd:MultiplayerPeer.set_target_peer
	class(self).SetTargetPeer(gd.Int(id))
}

/*
Returns the ID of the [MultiplayerPeer] who sent the next available packet. See [method PacketPeer.get_available_packet_count].
*/
func (self Instance) GetPacketPeer() int { //gd:MultiplayerPeer.get_packet_peer
	return int(int(class(self).GetPacketPeer()))
}

/*
Returns the channel over which the next available packet was received. See [method PacketPeer.get_available_packet_count].
*/
func (self Instance) GetPacketChannel() int { //gd:MultiplayerPeer.get_packet_channel
	return int(int(class(self).GetPacketChannel()))
}

/*
Returns the transfer mode the remote peer used to send the next available packet. See [method PacketPeer.get_available_packet_count].
*/
func (self Instance) GetPacketMode() gdclass.MultiplayerPeerTransferMode { //gd:MultiplayerPeer.get_packet_mode
	return gdclass.MultiplayerPeerTransferMode(class(self).GetPacketMode())
}

/*
Waits up to 1 second to receive a new network event.
*/
func (self Instance) Poll() { //gd:MultiplayerPeer.poll
	class(self).Poll()
}

/*
Immediately close the multiplayer peer returning to the state [constant CONNECTION_DISCONNECTED]. Connected peers will be dropped without emitting [signal peer_disconnected].
*/
func (self Instance) Close() { //gd:MultiplayerPeer.close
	class(self).Close()
}

/*
Disconnects the given [param peer] from this host. If [param force] is [code]true[/code] the [signal peer_disconnected] signal will not be emitted for this peer.
*/
func (self Instance) DisconnectPeer(peer int) { //gd:MultiplayerPeer.disconnect_peer
	class(self).DisconnectPeer(gd.Int(peer), false)
}

/*
Returns the current state of the connection. See [enum ConnectionStatus].
*/
func (self Instance) GetConnectionStatus() gdclass.MultiplayerPeerConnectionStatus { //gd:MultiplayerPeer.get_connection_status
	return gdclass.MultiplayerPeerConnectionStatus(class(self).GetConnectionStatus())
}

/*
Returns the ID of this [MultiplayerPeer].
*/
func (self Instance) GetUniqueId() int { //gd:MultiplayerPeer.get_unique_id
	return int(int(class(self).GetUniqueId()))
}

/*
Returns a randomly generated integer that can be used as a network unique ID.
*/
func (self Instance) GenerateUniqueId() int { //gd:MultiplayerPeer.generate_unique_id
	return int(int(class(self).GenerateUniqueId()))
}

/*
Returns true if the server can act as a relay in the current configuration (i.e. if the higher level [MultiplayerAPI] should notify connected clients of other peers, and implement a relay protocol to allow communication between them).
*/
func (self Instance) IsServerRelaySupported() bool { //gd:MultiplayerPeer.is_server_relay_supported
	return bool(class(self).IsServerRelaySupported())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.MultiplayerPeer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MultiplayerPeer"))
	casted := Instance{*(*gdclass.MultiplayerPeer)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) RefuseNewConnections() bool {
	return bool(class(self).IsRefusingNewConnections())
}

func (self Instance) SetRefuseNewConnections(value bool) {
	class(self).SetRefuseNewConnections(value)
}

func (self Instance) TransferMode() gdclass.MultiplayerPeerTransferMode {
	return gdclass.MultiplayerPeerTransferMode(class(self).GetTransferMode())
}

func (self Instance) SetTransferMode(value gdclass.MultiplayerPeerTransferMode) {
	class(self).SetTransferMode(value)
}

func (self Instance) TransferChannel() int {
	return int(int(class(self).GetTransferChannel()))
}

func (self Instance) SetTransferChannel(value int) {
	class(self).SetTransferChannel(gd.Int(value))
}

//go:nosplit
func (self class) SetTransferChannel(channel gd.Int) { //gd:MultiplayerPeer.set_transfer_channel
	var frame = callframe.New()
	callframe.Arg(frame, channel)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerPeer.Bind_set_transfer_channel, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTransferChannel() gd.Int { //gd:MultiplayerPeer.get_transfer_channel
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerPeer.Bind_get_transfer_channel, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTransferMode(mode gdclass.MultiplayerPeerTransferMode) { //gd:MultiplayerPeer.set_transfer_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerPeer.Bind_set_transfer_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTransferMode() gdclass.MultiplayerPeerTransferMode { //gd:MultiplayerPeer.get_transfer_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.MultiplayerPeerTransferMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerPeer.Bind_get_transfer_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the peer to which packets will be sent.
The [param id] can be one of: [constant TARGET_PEER_BROADCAST] to send to all connected peers, [constant TARGET_PEER_SERVER] to send to the peer acting as server, a valid peer ID to send to that specific peer, a negative peer ID to send to all peers except that one. By default, the target peer is [constant TARGET_PEER_BROADCAST].
*/
//go:nosplit
func (self class) SetTargetPeer(id gd.Int) { //gd:MultiplayerPeer.set_target_peer
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerPeer.Bind_set_target_peer, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the ID of the [MultiplayerPeer] who sent the next available packet. See [method PacketPeer.get_available_packet_count].
*/
//go:nosplit
func (self class) GetPacketPeer() gd.Int { //gd:MultiplayerPeer.get_packet_peer
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerPeer.Bind_get_packet_peer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the channel over which the next available packet was received. See [method PacketPeer.get_available_packet_count].
*/
//go:nosplit
func (self class) GetPacketChannel() gd.Int { //gd:MultiplayerPeer.get_packet_channel
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerPeer.Bind_get_packet_channel, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the transfer mode the remote peer used to send the next available packet. See [method PacketPeer.get_available_packet_count].
*/
//go:nosplit
func (self class) GetPacketMode() gdclass.MultiplayerPeerTransferMode { //gd:MultiplayerPeer.get_packet_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.MultiplayerPeerTransferMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerPeer.Bind_get_packet_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Waits up to 1 second to receive a new network event.
*/
//go:nosplit
func (self class) Poll() { //gd:MultiplayerPeer.poll
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerPeer.Bind_poll, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Immediately close the multiplayer peer returning to the state [constant CONNECTION_DISCONNECTED]. Connected peers will be dropped without emitting [signal peer_disconnected].
*/
//go:nosplit
func (self class) Close() { //gd:MultiplayerPeer.close
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerPeer.Bind_close, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Disconnects the given [param peer] from this host. If [param force] is [code]true[/code] the [signal peer_disconnected] signal will not be emitted for this peer.
*/
//go:nosplit
func (self class) DisconnectPeer(peer gd.Int, force bool) { //gd:MultiplayerPeer.disconnect_peer
	var frame = callframe.New()
	callframe.Arg(frame, peer)
	callframe.Arg(frame, force)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerPeer.Bind_disconnect_peer, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the current state of the connection. See [enum ConnectionStatus].
*/
//go:nosplit
func (self class) GetConnectionStatus() gdclass.MultiplayerPeerConnectionStatus { //gd:MultiplayerPeer.get_connection_status
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.MultiplayerPeerConnectionStatus](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerPeer.Bind_get_connection_status, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the ID of this [MultiplayerPeer].
*/
//go:nosplit
func (self class) GetUniqueId() gd.Int { //gd:MultiplayerPeer.get_unique_id
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerPeer.Bind_get_unique_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a randomly generated integer that can be used as a network unique ID.
*/
//go:nosplit
func (self class) GenerateUniqueId() gd.Int { //gd:MultiplayerPeer.generate_unique_id
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerPeer.Bind_generate_unique_id, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRefuseNewConnections(enable bool) { //gd:MultiplayerPeer.set_refuse_new_connections
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerPeer.Bind_set_refuse_new_connections, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsRefusingNewConnections() bool { //gd:MultiplayerPeer.is_refusing_new_connections
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerPeer.Bind_is_refusing_new_connections, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns true if the server can act as a relay in the current configuration (i.e. if the higher level [MultiplayerAPI] should notify connected clients of other peers, and implement a relay protocol to allow communication between them).
*/
//go:nosplit
func (self class) IsServerRelaySupported() bool { //gd:MultiplayerPeer.is_server_relay_supported
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerPeer.Bind_is_server_relay_supported, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnPeerConnected(cb func(id int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("peer_connected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPeerDisconnected(cb func(id int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("peer_disconnected"), gd.NewCallable(cb), 0)
}

func (self class) AsMultiplayerPeer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMultiplayerPeer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("MultiplayerPeer", func(ptr gd.Object) any {
		return [1]gdclass.MultiplayerPeer{*(*gdclass.MultiplayerPeer)(unsafe.Pointer(&ptr))}
	})
}

type ConnectionStatus = gdclass.MultiplayerPeerConnectionStatus //gd:MultiplayerPeer.ConnectionStatus

const (
	/*The MultiplayerPeer is disconnected.*/
	ConnectionDisconnected ConnectionStatus = 0
	/*The MultiplayerPeer is currently connecting to a server.*/
	ConnectionConnecting ConnectionStatus = 1
	/*This MultiplayerPeer is connected.*/
	ConnectionConnected ConnectionStatus = 2
)

type TransferMode = gdclass.MultiplayerPeerTransferMode //gd:MultiplayerPeer.TransferMode

const (
	/*Packets are not acknowledged, no resend attempts are made for lost packets. Packets may arrive in any order. Potentially faster than [constant TRANSFER_MODE_UNRELIABLE_ORDERED]. Use for non-critical data, and always consider whether the order matters.*/
	TransferModeUnreliable TransferMode = 0
	/*Packets are not acknowledged, no resend attempts are made for lost packets. Packets are received in the order they were sent in. Potentially faster than [constant TRANSFER_MODE_RELIABLE]. Use for non-critical data or data that would be outdated if received late due to resend attempt(s) anyway, for example movement and positional data.*/
	TransferModeUnreliableOrdered TransferMode = 1
	/*Packets must be received and resend attempts should be made until the packets are acknowledged. Packets must be received in the order they were sent in. Most reliable transfer mode, but potentially the slowest due to the overhead. Use for critical data that must be transmitted and arrive in order, for example an ability being triggered or a chat message. Consider carefully if the information really is critical, and use sparingly.*/
	TransferModeReliable TransferMode = 2
)
