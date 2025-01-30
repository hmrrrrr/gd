// Package MultiplayerAPIExtension provides methods for working with MultiplayerAPIExtension object instances.
package MultiplayerAPIExtension

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/MultiplayerAPI"
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
This class can be used to augment or replace the default [MultiplayerAPI] implementation via script or extensions.
The following example augment the default implementation ([SceneMultiplayer]) by logging every RPC being made, and every object being configured for replication.
[codeblocks]
[gdscript]
extends MultiplayerAPIExtension
class_name LogMultiplayer

# We want to augment the default SceneMultiplayer.
var base_multiplayer = SceneMultiplayer.new()

func _init():

	# Just passthrough base signals (copied to var to avoid cyclic reference)
	var cts = connected_to_server
	var cf = connection_failed
	var pc = peer_connected
	var pd = peer_disconnected
	base_multiplayer.connected_to_server.connect(func(): cts.emit())
	base_multiplayer.connection_failed.connect(func(): cf.emit())
	base_multiplayer.peer_connected.connect(func(id): pc.emit(id))
	base_multiplayer.peer_disconnected.connect(func(id): pd.emit(id))

func _poll():

	return base_multiplayer.poll()

# Log RPC being made and forward it to the default multiplayer.
func _rpc(peer: int, object: Object, method: StringName, args: Array) -> Error:

	print("Got RPC for %d: %s::%s(%s)" % [peer, object, method, args])
	return base_multiplayer.rpc(peer, object, method, args)

# Log configuration add. E.g. root path (nullptr, NodePath), replication (Node, Spawner|Synchronizer), custom.
func _object_configuration_add(object, config: Variant) -> Error:

	if config is MultiplayerSynchronizer:
	    print("Adding synchronization configuration for %s. Synchronizer: %s" % [object, config])
	elif config is MultiplayerSpawner:
	    print("Adding node %s to the spawn list. Spawner: %s" % [object, config])
	return base_multiplayer.object_configuration_add(object, config)

# Log configuration remove. E.g. root path (nullptr, NodePath), replication (Node, Spawner|Synchronizer), custom.
func _object_configuration_remove(object, config: Variant) -> Error:

	if config is MultiplayerSynchronizer:
	    print("Removing synchronization configuration for %s. Synchronizer: %s" % [object, config])
	elif config is MultiplayerSpawner:
	    print("Removing node %s from the spawn list. Spawner: %s" % [object, config])
	return base_multiplayer.object_configuration_remove(object, config)

# These can be optional, but in our case we want to augment SceneMultiplayer, so forward everything.
func _set_multiplayer_peer(p_peer: MultiplayerPeer):

	base_multiplayer.multiplayer_peer = p_peer

func _get_multiplayer_peer() -> MultiplayerPeer:

	return base_multiplayer.multiplayer_peer

func _get_unique_id() -> int:

	return base_multiplayer.get_unique_id()

func _get_peer_ids() -> PackedInt32Array:

	return base_multiplayer.get_peers()

[/gdscript]
[/codeblocks]
Then in your main scene or in an autoload call [method SceneTree.set_multiplayer] to start using your custom [MultiplayerAPI]:
[codeblocks]
[gdscript]
# autoload.gd
func _enter_tree():

	# Sets our custom multiplayer as the main one in SceneTree.

get_tree().set_multiplayer(LogMultiplayer.new())
[/gdscript]
[/codeblocks]
Native extensions can alternatively use the [method MultiplayerAPI.set_default_interface] method during initialization to configure themselves as the default implementation.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=MultiplayerAPIExtension)
*/
type Instance [1]gdclass.MultiplayerAPIExtension

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsMultiplayerAPIExtension() Instance
}
type Interface interface {
	//Callback for [method MultiplayerAPI.poll].
	Poll() error
	//Called when the [member MultiplayerAPI.multiplayer_peer] is set.
	SetMultiplayerPeer(multiplayer_peer [1]gdclass.MultiplayerPeer)
	//Called when the [member MultiplayerAPI.multiplayer_peer] is retrieved.
	GetMultiplayerPeer() [1]gdclass.MultiplayerPeer
	//Callback for [method MultiplayerAPI.get_unique_id].
	GetUniqueId() int
	//Callback for [method MultiplayerAPI.get_peers].
	GetPeerIds() []int32
	//Callback for [method MultiplayerAPI.rpc].
	Rpc(peer int, obj Object.Instance, method string, args []any) error
	//Callback for [method MultiplayerAPI.get_remote_sender_id].
	GetRemoteSenderId() int
	//Callback for [method MultiplayerAPI.object_configuration_add].
	ObjectConfigurationAdd(obj Object.Instance, configuration any) error
	//Callback for [method MultiplayerAPI.object_configuration_remove].
	ObjectConfigurationRemove(obj Object.Instance, configuration any) error
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) Poll() (_ error)                                                { return }
func (self implementation) SetMultiplayerPeer(multiplayer_peer [1]gdclass.MultiplayerPeer) { return }
func (self implementation) GetMultiplayerPeer() (_ [1]gdclass.MultiplayerPeer)             { return }
func (self implementation) GetUniqueId() (_ int)                                           { return }
func (self implementation) GetPeerIds() (_ []int32)                                        { return }
func (self implementation) Rpc(peer int, obj Object.Instance, method string, args []any) (_ error) {
	return
}
func (self implementation) GetRemoteSenderId() (_ int) { return }
func (self implementation) ObjectConfigurationAdd(obj Object.Instance, configuration any) (_ error) {
	return
}
func (self implementation) ObjectConfigurationRemove(obj Object.Instance, configuration any) (_ error) {
	return
}

/*
Callback for [method MultiplayerAPI.poll].
*/
func (Instance) _poll(impl func(ptr unsafe.Pointer) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(Error.New(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Called when the [member MultiplayerAPI.multiplayer_peer] is set.
*/
func (Instance) _set_multiplayer_peer(impl func(ptr unsafe.Pointer, multiplayer_peer [1]gdclass.MultiplayerPeer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var multiplayer_peer = [1]gdclass.MultiplayerPeer{pointers.New[gdclass.MultiplayerPeer]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(multiplayer_peer[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, multiplayer_peer)
	}
}

/*
Called when the [member MultiplayerAPI.multiplayer_peer] is retrieved.
*/
func (Instance) _get_multiplayer_peer(impl func(ptr unsafe.Pointer) [1]gdclass.MultiplayerPeer) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Callback for [method MultiplayerAPI.get_unique_id].
*/
func (Instance) _get_unique_id(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, int64(ret))
	}
}

/*
Callback for [method MultiplayerAPI.get_peers].
*/
func (Instance) _get_peer_ids(impl func(ptr unsafe.Pointer) []int32) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalPacked[gd.PackedInt32Array, int32](Packed.New(ret...)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Callback for [method MultiplayerAPI.rpc].
*/
func (Instance) _rpc(impl func(ptr unsafe.Pointer, peer int, obj Object.Instance, method string, args []any) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var peer = gd.UnsafeGet[int64](p_args, 0)

		var obj = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}
		defer pointers.End(obj[0])
		var method = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2)))))
		defer pointers.End(gd.InternalStringName(method))
		var args = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 3))))
		defer pointers.End(gd.InternalArray(args))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(peer), obj, method.String(), gd.ArrayAs[[]any](gd.InternalArray(args)))
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(Error.New(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Callback for [method MultiplayerAPI.get_remote_sender_id].
*/
func (Instance) _get_remote_sender_id(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, int64(ret))
	}
}

/*
Callback for [method MultiplayerAPI.object_configuration_add].
*/
func (Instance) _object_configuration_add(impl func(ptr unsafe.Pointer, obj Object.Instance, configuration any) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var obj = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(obj[0])
		var configuration = variant.Through(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 1))))
		defer pointers.End(gd.InternalVariant(configuration))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj, configuration.Interface())
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(Error.New(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Callback for [method MultiplayerAPI.object_configuration_remove].
*/
func (Instance) _object_configuration_remove(impl func(ptr unsafe.Pointer, obj Object.Instance, configuration any) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var obj = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(obj[0])
		var configuration = variant.Through(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 1))))
		defer pointers.End(gd.InternalVariant(configuration))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj, configuration.Interface())
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(Error.New(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.MultiplayerAPIExtension

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MultiplayerAPIExtension"))
	casted := Instance{*(*gdclass.MultiplayerAPIExtension)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Callback for [method MultiplayerAPI.poll].
*/
func (class) _poll(impl func(ptr unsafe.Pointer) Error.Code) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Called when the [member MultiplayerAPI.multiplayer_peer] is set.
*/
func (class) _set_multiplayer_peer(impl func(ptr unsafe.Pointer, multiplayer_peer [1]gdclass.MultiplayerPeer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var multiplayer_peer = [1]gdclass.MultiplayerPeer{pointers.New[gdclass.MultiplayerPeer]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(multiplayer_peer[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, multiplayer_peer)
	}
}

/*
Called when the [member MultiplayerAPI.multiplayer_peer] is retrieved.
*/
func (class) _get_multiplayer_peer(impl func(ptr unsafe.Pointer) [1]gdclass.MultiplayerPeer) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Callback for [method MultiplayerAPI.get_unique_id].
*/
func (class) _get_unique_id(impl func(ptr unsafe.Pointer) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Callback for [method MultiplayerAPI.get_peers].
*/
func (class) _get_peer_ids(impl func(ptr unsafe.Pointer) Packed.Array[int32]) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalPacked[gd.PackedInt32Array, int32](ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Callback for [method MultiplayerAPI.rpc].
*/
func (class) _rpc(impl func(ptr unsafe.Pointer, peer int64, obj [1]gd.Object, method String.Name, args Array.Any) Error.Code) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var peer = gd.UnsafeGet[int64](p_args, 0)

		var obj = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 1))})}
		defer pointers.End(obj[0])
		var method = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2)))))
		defer pointers.End(gd.InternalStringName(method))
		var args = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 3))))
		defer pointers.End(gd.InternalArray(args))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, peer, obj, method, args)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Callback for [method MultiplayerAPI.get_remote_sender_id].
*/
func (class) _get_remote_sender_id(impl func(ptr unsafe.Pointer) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Callback for [method MultiplayerAPI.object_configuration_add].
*/
func (class) _object_configuration_add(impl func(ptr unsafe.Pointer, obj [1]gd.Object, configuration variant.Any) Error.Code) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var obj = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(obj[0])
		var configuration = variant.Through(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 1))))
		defer pointers.End(gd.InternalVariant(configuration))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj, configuration)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Callback for [method MultiplayerAPI.object_configuration_remove].
*/
func (class) _object_configuration_remove(impl func(ptr unsafe.Pointer, obj [1]gd.Object, configuration variant.Any) Error.Code) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var obj = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(obj[0])
		var configuration = variant.Through(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 1))))
		defer pointers.End(gd.InternalVariant(configuration))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj, configuration)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (self class) AsMultiplayerAPIExtension() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMultiplayerAPIExtension() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsMultiplayerAPI() MultiplayerAPI.Advanced {
	return *((*MultiplayerAPI.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsMultiplayerAPI() MultiplayerAPI.Instance {
	return *((*MultiplayerAPI.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_poll":
		return reflect.ValueOf(self._poll)
	case "_set_multiplayer_peer":
		return reflect.ValueOf(self._set_multiplayer_peer)
	case "_get_multiplayer_peer":
		return reflect.ValueOf(self._get_multiplayer_peer)
	case "_get_unique_id":
		return reflect.ValueOf(self._get_unique_id)
	case "_get_peer_ids":
		return reflect.ValueOf(self._get_peer_ids)
	case "_rpc":
		return reflect.ValueOf(self._rpc)
	case "_get_remote_sender_id":
		return reflect.ValueOf(self._get_remote_sender_id)
	case "_object_configuration_add":
		return reflect.ValueOf(self._object_configuration_add)
	case "_object_configuration_remove":
		return reflect.ValueOf(self._object_configuration_remove)
	default:
		return gd.VirtualByName(MultiplayerAPI.Advanced(self.AsMultiplayerAPI()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_poll":
		return reflect.ValueOf(self._poll)
	case "_set_multiplayer_peer":
		return reflect.ValueOf(self._set_multiplayer_peer)
	case "_get_multiplayer_peer":
		return reflect.ValueOf(self._get_multiplayer_peer)
	case "_get_unique_id":
		return reflect.ValueOf(self._get_unique_id)
	case "_get_peer_ids":
		return reflect.ValueOf(self._get_peer_ids)
	case "_rpc":
		return reflect.ValueOf(self._rpc)
	case "_get_remote_sender_id":
		return reflect.ValueOf(self._get_remote_sender_id)
	case "_object_configuration_add":
		return reflect.ValueOf(self._object_configuration_add)
	case "_object_configuration_remove":
		return reflect.ValueOf(self._object_configuration_remove)
	default:
		return gd.VirtualByName(MultiplayerAPI.Instance(self.AsMultiplayerAPI()), name)
	}
}
func init() {
	gdclass.Register("MultiplayerAPIExtension", func(ptr gd.Object) any {
		return [1]gdclass.MultiplayerAPIExtension{*(*gdclass.MultiplayerAPIExtension)(unsafe.Pointer(&ptr))}
	})
}
