package DTLSServer

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
This class is used to store the state of a DTLS server. Upon [method setup] it converts connected [PacketPeerUDP] to [PacketPeerDTLS] accepting them via [method take_connection] as DTLS clients. Under the hood, this class is used to store the DTLS state and cookies of the server. The reason of why the state and cookies are needed is outside of the scope of this documentation.
Below a small example of how to use it:
[codeblocks]
[gdscript]
# server_node.gd
extends Node

var dtls := DTLSServer.new()
var server := UDPServer.new()
var peers = []

func _ready():
    server.listen(4242)
    var key = load("key.key") # Your private key.
    var cert = load("cert.crt") # Your X509 certificate.
    dtls.setup(key, cert)

func _process(delta):
    while server.is_connection_available():
        var peer: PacketPeerUDP = server.take_connection()
        var dtls_peer: PacketPeerDTLS = dtls.take_connection(peer)
        if dtls_peer.get_status() != PacketPeerDTLS.STATUS_HANDSHAKING:
            continue # It is normal that 50% of the connections fails due to cookie exchange.
        print("Peer connected!")
        peers.append(dtls_peer)

    for p in peers:
        p.poll() # Must poll to update the state.
        if p.get_status() == PacketPeerDTLS.STATUS_CONNECTED:
            while p.get_available_packet_count() > 0:
                print("Received message from client: %s" % p.get_packet().get_string_from_utf8())
                p.put_packet("Hello DTLS client".to_utf8_buffer())
[/gdscript]
[csharp]
// ServerNode.cs
using Godot;

public partial class ServerNode : Node
{
    private DtlsServer _dtls = new DtlsServer();
    private UdpServer _server = new UdpServer();
    private Godot.Collections.Array<PacketPeerDtls> _peers = new Godot.Collections.Array<PacketPeerDtls>();

    public override void _Ready()
    {
        _server.Listen(4242);
        var key = GD.Load<CryptoKey>("key.key"); // Your private key.
        var cert = GD.Load<X509Certificate>("cert.crt"); // Your X509 certificate.
        _dtls.Setup(key, cert);
    }

    public override void _Process(double delta)
    {
        while (Server.IsConnectionAvailable())
        {
            PacketPeerUdp peer = _server.TakeConnection();
            PacketPeerDtls dtlsPeer = _dtls.TakeConnection(peer);
            if (dtlsPeer.GetStatus() != PacketPeerDtls.Status.Handshaking)
            {
                continue; // It is normal that 50% of the connections fails due to cookie exchange.
            }
            GD.Print("Peer connected!");
            _peers.Add(dtlsPeer);
        }

        foreach (var p in _peers)
        {
            p.Poll(); // Must poll to update the state.
            if (p.GetStatus() == PacketPeerDtls.Status.Connected)
            {
                while (p.GetAvailablePacketCount() > 0)
                {
                    GD.Print($"Received Message From Client: {p.GetPacket().GetStringFromUtf8()}");
                    p.PutPacket("Hello DTLS Client".ToUtf8Buffer());
                }
            }
        }
    }
}
[/csharp]
[/codeblocks]
[codeblocks]
[gdscript]
# client_node.gd
extends Node

var dtls := PacketPeerDTLS.new()
var udp := PacketPeerUDP.new()
var connected = false

func _ready():
    udp.connect_to_host("127.0.0.1", 4242)
    dtls.connect_to_peer(udp, false) # Use true in production for certificate validation!

func _process(delta):
    dtls.poll()
    if dtls.get_status() == PacketPeerDTLS.STATUS_CONNECTED:
        if !connected:
            # Try to contact server
            dtls.put_packet("The answer is... 42!".to_utf8_buffer())
        while dtls.get_available_packet_count() > 0:
            print("Connected: %s" % dtls.get_packet().get_string_from_utf8())
            connected = true
[/gdscript]
[csharp]
// ClientNode.cs
using Godot;
using System.Text;

public partial class ClientNode : Node
{
    private PacketPeerDtls _dtls = new PacketPeerDtls();
    private PacketPeerUdp _udp = new PacketPeerUdp();
    private bool _connected = false;

    public override void _Ready()
    {
        _udp.ConnectToHost("127.0.0.1", 4242);
        _dtls.ConnectToPeer(_udp, validateCerts: false); // Use true in production for certificate validation!
    }

    public override void _Process(double delta)
    {
        _dtls.Poll();
        if (_dtls.GetStatus() == PacketPeerDtls.Status.Connected)
        {
            if (!_connected)
            {
                // Try to contact server
                _dtls.PutPacket("The Answer Is..42!".ToUtf8Buffer());
            }
            while (_dtls.GetAvailablePacketCount() > 0)
            {
                GD.Print($"Connected: {_dtls.GetPacket().GetStringFromUtf8()}");
                _connected = true;
            }
        }
    }
}
[/csharp]
[/codeblocks]

*/
type Go [1]classdb.DTLSServer

/*
Setup the DTLS server to use the given [param server_options]. See [method TLSOptions.server].
*/
func (self Go) Setup(server_options gdclass.TLSOptions) gd.Error {
	return gd.Error(class(self).Setup(server_options))
}

/*
Try to initiate the DTLS handshake with the given [param udp_peer] which must be already connected (see [method PacketPeerUDP.connect_to_host]).
[b]Note:[/b] You must check that the state of the return PacketPeerUDP is [constant PacketPeerDTLS.STATUS_HANDSHAKING], as it is normal that 50% of the new connections will be invalid due to cookie exchange.
*/
func (self Go) TakeConnection(udp_peer gdclass.PacketPeerUDP) gdclass.PacketPeerDTLS {
	return gdclass.PacketPeerDTLS(class(self).TakeConnection(udp_peer))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.DTLSServer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("DTLSServer"))
	return Go{classdb.DTLSServer(object)}
}

/*
Setup the DTLS server to use the given [param server_options]. See [method TLSOptions.server].
*/
//go:nosplit
func (self class) Setup(server_options gdclass.TLSOptions) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(server_options[0])[0])
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DTLSServer.Bind_setup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Try to initiate the DTLS handshake with the given [param udp_peer] which must be already connected (see [method PacketPeerUDP.connect_to_host]).
[b]Note:[/b] You must check that the state of the return PacketPeerUDP is [constant PacketPeerDTLS.STATUS_HANDSHAKING], as it is normal that 50% of the new connections will be invalid due to cookie exchange.
*/
//go:nosplit
func (self class) TakeConnection(udp_peer gdclass.PacketPeerUDP) gdclass.PacketPeerDTLS {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(udp_peer[0])[0])
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.DTLSServer.Bind_take_connection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.PacketPeerDTLS{classdb.PacketPeerDTLS(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self class) AsDTLSServer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsDTLSServer() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("DTLSServer", func(ptr gd.Object) any { return classdb.DTLSServer(ptr) })}
