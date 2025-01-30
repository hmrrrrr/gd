// Package ENetMultiplayerPeer provides methods for working with ENetMultiplayerPeer object instances.
package ENetMultiplayerPeer

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
A MultiplayerPeer implementation that should be passed to [member MultiplayerAPI.multiplayer_peer] after being initialized as either a client, server, or mesh. Events can then be handled by connecting to [MultiplayerAPI] signals. See [ENetConnection] for more information on the ENet library wrapper.
[b]Note:[/b] ENet only uses UDP, not TCP. When forwarding the server port to make your server accessible on the public Internet, you only need to forward the server port in UDP. You can use the [UPNP] class to try to forward the server port automatically when starting the server.
*/
type Instance [1]gdclass.ENetMultiplayerPeer

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsENetMultiplayerPeer() Instance
}

/*
Create server that listens to connections via [param port]. The port needs to be an available, unused port between 0 and 65535. Note that ports below 1024 are privileged and may require elevated permissions depending on the platform. To change the interface the server listens on, use [method set_bind_ip]. The default IP is the wildcard [code]"*"[/code], which listens on all available interfaces. [param max_clients] is the maximum number of clients that are allowed at once, any number up to 4095 may be used, although the achievable number of simultaneous clients may be far lower and depends on the application. For additional details on the bandwidth parameters, see [method create_client]. Returns [constant OK] if a server was created, [constant ERR_ALREADY_IN_USE] if this ENetMultiplayerPeer instance already has an open connection (in which case you need to call [method MultiplayerPeer.close] first) or [constant ERR_CANT_CREATE] if the server could not be created.
*/
func (self Instance) CreateServer(port int) error { //gd:ENetMultiplayerPeer.create_server
	return error(gd.ToError(class(self).CreateServer(int64(port), int64(32), int64(0), int64(0), int64(0))))
}

/*
Create client that connects to a server at [param address] using specified [param port]. The given address needs to be either a fully qualified domain name (e.g. [code]"www.example.com"[/code]) or an IP address in IPv4 or IPv6 format (e.g. [code]"192.168.1.1"[/code]). The [param port] is the port the server is listening on. The [param channel_count] parameter can be used to specify the number of ENet channels allocated for the connection. The [param in_bandwidth] and [param out_bandwidth] parameters can be used to limit the incoming and outgoing bandwidth to the given number of bytes per second. The default of 0 means unlimited bandwidth. Note that ENet will strategically drop packets on specific sides of a connection between peers to ensure the peer's bandwidth is not overwhelmed. The bandwidth parameters also determine the window size of a connection which limits the amount of reliable packets that may be in transit at any given time. Returns [constant OK] if a client was created, [constant ERR_ALREADY_IN_USE] if this ENetMultiplayerPeer instance already has an open connection (in which case you need to call [method MultiplayerPeer.close] first) or [constant ERR_CANT_CREATE] if the client could not be created. If [param local_port] is specified, the client will also listen to the given port; this is useful for some NAT traversal techniques.
*/
func (self Instance) CreateClient(address string, port int) error { //gd:ENetMultiplayerPeer.create_client
	return error(gd.ToError(class(self).CreateClient(String.New(address), int64(port), int64(0), int64(0), int64(0), int64(0))))
}

/*
Initialize this [MultiplayerPeer] in mesh mode. The provided [param unique_id] will be used as the local peer network unique ID once assigned as the [member MultiplayerAPI.multiplayer_peer]. In the mesh configuration you will need to set up each new peer manually using [ENetConnection] before calling [method add_mesh_peer]. While this technique is more advanced, it allows for better control over the connection process (e.g. when dealing with NAT punch-through) and for better distribution of the network load (which would otherwise be more taxing on the server).
*/
func (self Instance) CreateMesh(unique_id int) error { //gd:ENetMultiplayerPeer.create_mesh
	return error(gd.ToError(class(self).CreateMesh(int64(unique_id))))
}

/*
Add a new remote peer with the given [param peer_id] connected to the given [param host].
[b]Note:[/b] The [param host] must have exactly one peer in the [constant ENetPacketPeer.STATE_CONNECTED] state.
*/
func (self Instance) AddMeshPeer(peer_id int, host [1]gdclass.ENetConnection) error { //gd:ENetMultiplayerPeer.add_mesh_peer
	return error(gd.ToError(class(self).AddMeshPeer(int64(peer_id), host)))
}

/*
The IP used when creating a server. This is set to the wildcard [code]"*"[/code] by default, which binds to all available interfaces. The given IP needs to be in IPv4 or IPv6 address format, for example: [code]"192.168.1.1"[/code].
*/
func (self Instance) SetBindIp(ip string) { //gd:ENetMultiplayerPeer.set_bind_ip
	class(self).SetBindIp(String.New(ip))
}

/*
Returns the [ENetPacketPeer] associated to the given [param id].
*/
func (self Instance) GetPeer(id int) [1]gdclass.ENetPacketPeer { //gd:ENetMultiplayerPeer.get_peer
	return [1]gdclass.ENetPacketPeer(class(self).GetPeer(int64(id)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ENetMultiplayerPeer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ENetMultiplayerPeer"))
	casted := Instance{*(*gdclass.ENetMultiplayerPeer)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Host() [1]gdclass.ENetConnection {
	return [1]gdclass.ENetConnection(class(self).GetHost())
}

/*
Create server that listens to connections via [param port]. The port needs to be an available, unused port between 0 and 65535. Note that ports below 1024 are privileged and may require elevated permissions depending on the platform. To change the interface the server listens on, use [method set_bind_ip]. The default IP is the wildcard [code]"*"[/code], which listens on all available interfaces. [param max_clients] is the maximum number of clients that are allowed at once, any number up to 4095 may be used, although the achievable number of simultaneous clients may be far lower and depends on the application. For additional details on the bandwidth parameters, see [method create_client]. Returns [constant OK] if a server was created, [constant ERR_ALREADY_IN_USE] if this ENetMultiplayerPeer instance already has an open connection (in which case you need to call [method MultiplayerPeer.close] first) or [constant ERR_CANT_CREATE] if the server could not be created.
*/
//go:nosplit
func (self class) CreateServer(port int64, max_clients int64, max_channels int64, in_bandwidth int64, out_bandwidth int64) Error.Code { //gd:ENetMultiplayerPeer.create_server
	var frame = callframe.New()
	callframe.Arg(frame, port)
	callframe.Arg(frame, max_clients)
	callframe.Arg(frame, max_channels)
	callframe.Arg(frame, in_bandwidth)
	callframe.Arg(frame, out_bandwidth)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetMultiplayerPeer.Bind_create_server, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Create client that connects to a server at [param address] using specified [param port]. The given address needs to be either a fully qualified domain name (e.g. [code]"www.example.com"[/code]) or an IP address in IPv4 or IPv6 format (e.g. [code]"192.168.1.1"[/code]). The [param port] is the port the server is listening on. The [param channel_count] parameter can be used to specify the number of ENet channels allocated for the connection. The [param in_bandwidth] and [param out_bandwidth] parameters can be used to limit the incoming and outgoing bandwidth to the given number of bytes per second. The default of 0 means unlimited bandwidth. Note that ENet will strategically drop packets on specific sides of a connection between peers to ensure the peer's bandwidth is not overwhelmed. The bandwidth parameters also determine the window size of a connection which limits the amount of reliable packets that may be in transit at any given time. Returns [constant OK] if a client was created, [constant ERR_ALREADY_IN_USE] if this ENetMultiplayerPeer instance already has an open connection (in which case you need to call [method MultiplayerPeer.close] first) or [constant ERR_CANT_CREATE] if the client could not be created. If [param local_port] is specified, the client will also listen to the given port; this is useful for some NAT traversal techniques.
*/
//go:nosplit
func (self class) CreateClient(address String.Readable, port int64, channel_count int64, in_bandwidth int64, out_bandwidth int64, local_port int64) Error.Code { //gd:ENetMultiplayerPeer.create_client
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(address)))
	callframe.Arg(frame, port)
	callframe.Arg(frame, channel_count)
	callframe.Arg(frame, in_bandwidth)
	callframe.Arg(frame, out_bandwidth)
	callframe.Arg(frame, local_port)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetMultiplayerPeer.Bind_create_client, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Initialize this [MultiplayerPeer] in mesh mode. The provided [param unique_id] will be used as the local peer network unique ID once assigned as the [member MultiplayerAPI.multiplayer_peer]. In the mesh configuration you will need to set up each new peer manually using [ENetConnection] before calling [method add_mesh_peer]. While this technique is more advanced, it allows for better control over the connection process (e.g. when dealing with NAT punch-through) and for better distribution of the network load (which would otherwise be more taxing on the server).
*/
//go:nosplit
func (self class) CreateMesh(unique_id int64) Error.Code { //gd:ENetMultiplayerPeer.create_mesh
	var frame = callframe.New()
	callframe.Arg(frame, unique_id)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetMultiplayerPeer.Bind_create_mesh, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Add a new remote peer with the given [param peer_id] connected to the given [param host].
[b]Note:[/b] The [param host] must have exactly one peer in the [constant ENetPacketPeer.STATE_CONNECTED] state.
*/
//go:nosplit
func (self class) AddMeshPeer(peer_id int64, host [1]gdclass.ENetConnection) Error.Code { //gd:ENetMultiplayerPeer.add_mesh_peer
	var frame = callframe.New()
	callframe.Arg(frame, peer_id)
	callframe.Arg(frame, pointers.Get(host[0])[0])
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetMultiplayerPeer.Bind_add_mesh_peer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
The IP used when creating a server. This is set to the wildcard [code]"*"[/code] by default, which binds to all available interfaces. The given IP needs to be in IPv4 or IPv6 address format, for example: [code]"192.168.1.1"[/code].
*/
//go:nosplit
func (self class) SetBindIp(ip String.Readable) { //gd:ENetMultiplayerPeer.set_bind_ip
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(ip)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetMultiplayerPeer.Bind_set_bind_ip, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHost() [1]gdclass.ENetConnection { //gd:ENetMultiplayerPeer.get_host
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetMultiplayerPeer.Bind_get_host, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.ENetConnection{gd.PointerWithOwnershipTransferredToGo[gdclass.ENetConnection](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the [ENetPacketPeer] associated to the given [param id].
*/
//go:nosplit
func (self class) GetPeer(id int64) [1]gdclass.ENetPacketPeer { //gd:ENetMultiplayerPeer.get_peer
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ENetMultiplayerPeer.Bind_get_peer, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.ENetPacketPeer{gd.PointerWithOwnershipTransferredToGo[gdclass.ENetPacketPeer](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsENetMultiplayerPeer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsENetMultiplayerPeer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("ENetMultiplayerPeer", func(ptr gd.Object) any {
		return [1]gdclass.ENetMultiplayerPeer{*(*gdclass.ENetMultiplayerPeer)(unsafe.Pointer(&ptr))}
	})
}
