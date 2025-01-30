// Package WebRTCMultiplayerPeer provides methods for working with WebRTCMultiplayerPeer object instances.
package WebRTCMultiplayerPeer

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/MultiplayerPeer"
import "graphics.gd/classdb/PacketPeer"
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
This class constructs a full mesh of [WebRTCPeerConnection] (one connection for each peer) that can be used as a [member MultiplayerAPI.multiplayer_peer].
You can add each [WebRTCPeerConnection] via [method add_peer] or remove them via [method remove_peer]. Peers must be added in [constant WebRTCPeerConnection.STATE_NEW] state to allow it to create the appropriate channels. This class will not create offers nor set descriptions, it will only poll them, and notify connections and disconnections.
When creating the peer via [method create_client] or [method create_server] the [method MultiplayerPeer.is_server_relay_supported] method will return [code]true[/code] enabling peer exchange and packet relaying when supported by the [MultiplayerAPI] implementation.
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.
*/
type Instance [1]gdclass.WebRTCMultiplayerPeer

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsWebRTCMultiplayerPeer() Instance
}

/*
Initialize the multiplayer peer as a server (with unique ID of [code]1[/code]). This mode enables [method MultiplayerPeer.is_server_relay_supported], allowing the upper [MultiplayerAPI] layer to perform peer exchange and packet relaying.
You can optionally specify a [param channels_config] array of [enum MultiplayerPeer.TransferMode] which will be used to create extra channels (WebRTC only supports one transfer mode per channel).
*/
func (self Instance) CreateServer() error { //gd:WebRTCMultiplayerPeer.create_server
	return error(gd.ToError(class(self).CreateServer(Array.Nil)))
}

/*
Initialize the multiplayer peer as a client with the given [param peer_id] (must be between 2 and 2147483647). In this mode, you should only call [method add_peer] once and with [param peer_id] of [code]1[/code]. This mode enables [method MultiplayerPeer.is_server_relay_supported], allowing the upper [MultiplayerAPI] layer to perform peer exchange and packet relaying.
You can optionally specify a [param channels_config] array of [enum MultiplayerPeer.TransferMode] which will be used to create extra channels (WebRTC only supports one transfer mode per channel).
*/
func (self Instance) CreateClient(peer_id int) error { //gd:WebRTCMultiplayerPeer.create_client
	return error(gd.ToError(class(self).CreateClient(int64(peer_id), Array.Nil)))
}

/*
Initialize the multiplayer peer as a mesh (i.e. all peers connect to each other) with the given [param peer_id] (must be between 1 and 2147483647).
*/
func (self Instance) CreateMesh(peer_id int) error { //gd:WebRTCMultiplayerPeer.create_mesh
	return error(gd.ToError(class(self).CreateMesh(int64(peer_id), Array.Nil)))
}

/*
Add a new peer to the mesh with the given [param peer_id]. The [WebRTCPeerConnection] must be in state [constant WebRTCPeerConnection.STATE_NEW].
Three channels will be created for reliable, unreliable, and ordered transport. The value of [param unreliable_lifetime] will be passed to the [code]"maxPacketLifetime"[/code] option when creating unreliable and ordered channels (see [method WebRTCPeerConnection.create_data_channel]).
*/
func (self Instance) AddPeer(peer [1]gdclass.WebRTCPeerConnection, peer_id int) error { //gd:WebRTCMultiplayerPeer.add_peer
	return error(gd.ToError(class(self).AddPeer(peer, int64(peer_id), int64(1))))
}

/*
Remove the peer with given [param peer_id] from the mesh. If the peer was connected, and [signal MultiplayerPeer.peer_connected] was emitted for it, then [signal MultiplayerPeer.peer_disconnected] will be emitted.
*/
func (self Instance) RemovePeer(peer_id int) { //gd:WebRTCMultiplayerPeer.remove_peer
	class(self).RemovePeer(int64(peer_id))
}

/*
Returns [code]true[/code] if the given [param peer_id] is in the peers map (it might not be connected though).
*/
func (self Instance) HasPeer(peer_id int) bool { //gd:WebRTCMultiplayerPeer.has_peer
	return bool(class(self).HasPeer(int64(peer_id)))
}

/*
Returns a dictionary representation of the peer with given [param peer_id] with three keys. [code]"connection"[/code] containing the [WebRTCPeerConnection] to this peer, [code]"channels"[/code] an array of three [WebRTCDataChannel], and [code]"connected"[/code] a boolean representing if the peer connection is currently connected (all three channels are open).
*/
func (self Instance) GetPeer(peer_id int) Conn { //gd:WebRTCMultiplayerPeer.get_peer
	return Conn(gd.DictionaryAs[Conn](class(self).GetPeer(int64(peer_id))))
}

/*
Returns a dictionary which keys are the peer ids and values the peer representation as in [method get_peer].
*/
func (self Instance) GetPeers() map[int]Conn { //gd:WebRTCMultiplayerPeer.get_peers
	return map[int]Conn(gd.DictionaryAs[map[int]Conn](class(self).GetPeers()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.WebRTCMultiplayerPeer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("WebRTCMultiplayerPeer"))
	casted := Instance{*(*gdclass.WebRTCMultiplayerPeer)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Initialize the multiplayer peer as a server (with unique ID of [code]1[/code]). This mode enables [method MultiplayerPeer.is_server_relay_supported], allowing the upper [MultiplayerAPI] layer to perform peer exchange and packet relaying.
You can optionally specify a [param channels_config] array of [enum MultiplayerPeer.TransferMode] which will be used to create extra channels (WebRTC only supports one transfer mode per channel).
*/
//go:nosplit
func (self class) CreateServer(channels_config Array.Any) Error.Code { //gd:WebRTCMultiplayerPeer.create_server
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(channels_config)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCMultiplayerPeer.Bind_create_server, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Initialize the multiplayer peer as a client with the given [param peer_id] (must be between 2 and 2147483647). In this mode, you should only call [method add_peer] once and with [param peer_id] of [code]1[/code]. This mode enables [method MultiplayerPeer.is_server_relay_supported], allowing the upper [MultiplayerAPI] layer to perform peer exchange and packet relaying.
You can optionally specify a [param channels_config] array of [enum MultiplayerPeer.TransferMode] which will be used to create extra channels (WebRTC only supports one transfer mode per channel).
*/
//go:nosplit
func (self class) CreateClient(peer_id int64, channels_config Array.Any) Error.Code { //gd:WebRTCMultiplayerPeer.create_client
	var frame = callframe.New()
	callframe.Arg(frame, peer_id)
	callframe.Arg(frame, pointers.Get(gd.InternalArray(channels_config)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCMultiplayerPeer.Bind_create_client, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Initialize the multiplayer peer as a mesh (i.e. all peers connect to each other) with the given [param peer_id] (must be between 1 and 2147483647).
*/
//go:nosplit
func (self class) CreateMesh(peer_id int64, channels_config Array.Any) Error.Code { //gd:WebRTCMultiplayerPeer.create_mesh
	var frame = callframe.New()
	callframe.Arg(frame, peer_id)
	callframe.Arg(frame, pointers.Get(gd.InternalArray(channels_config)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCMultiplayerPeer.Bind_create_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Add a new peer to the mesh with the given [param peer_id]. The [WebRTCPeerConnection] must be in state [constant WebRTCPeerConnection.STATE_NEW].
Three channels will be created for reliable, unreliable, and ordered transport. The value of [param unreliable_lifetime] will be passed to the [code]"maxPacketLifetime"[/code] option when creating unreliable and ordered channels (see [method WebRTCPeerConnection.create_data_channel]).
*/
//go:nosplit
func (self class) AddPeer(peer [1]gdclass.WebRTCPeerConnection, peer_id int64, unreliable_lifetime int64) Error.Code { //gd:WebRTCMultiplayerPeer.add_peer
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(peer[0])[0])
	callframe.Arg(frame, peer_id)
	callframe.Arg(frame, unreliable_lifetime)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCMultiplayerPeer.Bind_add_peer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Remove the peer with given [param peer_id] from the mesh. If the peer was connected, and [signal MultiplayerPeer.peer_connected] was emitted for it, then [signal MultiplayerPeer.peer_disconnected] will be emitted.
*/
//go:nosplit
func (self class) RemovePeer(peer_id int64) { //gd:WebRTCMultiplayerPeer.remove_peer
	var frame = callframe.New()
	callframe.Arg(frame, peer_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCMultiplayerPeer.Bind_remove_peer, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the given [param peer_id] is in the peers map (it might not be connected though).
*/
//go:nosplit
func (self class) HasPeer(peer_id int64) bool { //gd:WebRTCMultiplayerPeer.has_peer
	var frame = callframe.New()
	callframe.Arg(frame, peer_id)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCMultiplayerPeer.Bind_has_peer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a dictionary representation of the peer with given [param peer_id] with three keys. [code]"connection"[/code] containing the [WebRTCPeerConnection] to this peer, [code]"channels"[/code] an array of three [WebRTCDataChannel], and [code]"connected"[/code] a boolean representing if the peer connection is currently connected (all three channels are open).
*/
//go:nosplit
func (self class) GetPeer(peer_id int64) Dictionary.Any { //gd:WebRTCMultiplayerPeer.get_peer
	var frame = callframe.New()
	callframe.Arg(frame, peer_id)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCMultiplayerPeer.Bind_get_peer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns a dictionary which keys are the peer ids and values the peer representation as in [method get_peer].
*/
//go:nosplit
func (self class) GetPeers() Dictionary.Any { //gd:WebRTCMultiplayerPeer.get_peers
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.WebRTCMultiplayerPeer.Bind_get_peers, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}
func (self class) AsWebRTCMultiplayerPeer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsWebRTCMultiplayerPeer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsMultiplayerPeer() MultiplayerPeer.Advanced {
	return *((*MultiplayerPeer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsMultiplayerPeer() MultiplayerPeer.Instance {
	return *((*MultiplayerPeer.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(MultiplayerPeer.Advanced(self.AsMultiplayerPeer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(MultiplayerPeer.Instance(self.AsMultiplayerPeer()), name)
	}
}
func init() {
	gdclass.Register("WebRTCMultiplayerPeer", func(ptr gd.Object) any {
		return [1]gdclass.WebRTCMultiplayerPeer{*(*gdclass.WebRTCMultiplayerPeer)(unsafe.Pointer(&ptr))}
	})
}

type Conn struct {
	Connection [1]gdclass.WebRTCPeerConnection `gd:"connection"`
	Channels   [][1]gdclass.WebRTCDataChannel  `gd:"channels"`
	Connected  bool                            `gd:"connected"`
}
