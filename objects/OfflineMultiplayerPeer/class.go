package OfflineMultiplayerPeer

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/MultiplayerPeer"
import "graphics.gd/objects/PacketPeer"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
This is the default [member MultiplayerAPI.multiplayer_peer] for the [member Node.multiplayer]. It mimics the behavior of a server with no peers connected.
This means that the [SceneTree] will act as the multiplayer authority by default. Calls to [method MultiplayerAPI.is_server] will return [code]true[/code], and calls to [method MultiplayerAPI.get_unique_id] will return [constant MultiplayerPeer.TARGET_PEER_SERVER].
*/
type Instance [1]classdb.OfflineMultiplayerPeer
type Any interface {
	gd.IsClass
	AsOfflineMultiplayerPeer() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.OfflineMultiplayerPeer

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OfflineMultiplayerPeer"))
	return Instance{classdb.OfflineMultiplayerPeer(object)}
}

func (self class) AsOfflineMultiplayerPeer() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsOfflineMultiplayerPeer() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
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
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsMultiplayerPeer(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsMultiplayerPeer(), name)
	}
}
func init() {
	classdb.Register("OfflineMultiplayerPeer", func(ptr gd.Object) any { return [1]classdb.OfflineMultiplayerPeer{classdb.OfflineMultiplayerPeer(ptr)} })
}
