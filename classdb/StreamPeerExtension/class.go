// Package StreamPeerExtension provides methods for working with StreamPeerExtension object instances.
package StreamPeerExtension

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/StreamPeer"
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

type Instance [1]gdclass.StreamPeerExtension

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsStreamPeerExtension() Instance
}
type Interface interface {
	GetData(r_buffer unsafe.Pointer, r_bytes int, r_received *int32) error
	GetPartialData(r_buffer unsafe.Pointer, r_bytes int, r_received *int32) error
	PutData(p_data unsafe.Pointer, p_bytes int, r_sent *int32) error
	PutPartialData(p_data unsafe.Pointer, p_bytes int, r_sent *int32) error
	GetAvailableBytes() int
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) GetData(r_buffer unsafe.Pointer, r_bytes int, r_received *int32) (_ error) {
	return
}
func (self implementation) GetPartialData(r_buffer unsafe.Pointer, r_bytes int, r_received *int32) (_ error) {
	return
}
func (self implementation) PutData(p_data unsafe.Pointer, p_bytes int, r_sent *int32) (_ error) {
	return
}
func (self implementation) PutPartialData(p_data unsafe.Pointer, p_bytes int, r_sent *int32) (_ error) {
	return
}
func (self implementation) GetAvailableBytes() (_ int) { return }
func (Instance) _get_data(impl func(ptr unsafe.Pointer, r_buffer unsafe.Pointer, r_bytes int, r_received *int32) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var r_buffer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		var r_bytes = gd.UnsafeGet[int64](p_args, 1)

		var r_received = gd.UnsafeGet[*int32](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, r_buffer, int(r_bytes), r_received)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(Error.New(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _get_partial_data(impl func(ptr unsafe.Pointer, r_buffer unsafe.Pointer, r_bytes int, r_received *int32) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var r_buffer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		var r_bytes = gd.UnsafeGet[int64](p_args, 1)

		var r_received = gd.UnsafeGet[*int32](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, r_buffer, int(r_bytes), r_received)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(Error.New(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _put_data(impl func(ptr unsafe.Pointer, p_data unsafe.Pointer, p_bytes int, r_sent *int32) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_data = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		var p_bytes = gd.UnsafeGet[int64](p_args, 1)

		var r_sent = gd.UnsafeGet[*int32](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_data, int(p_bytes), r_sent)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(Error.New(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _put_partial_data(impl func(ptr unsafe.Pointer, p_data unsafe.Pointer, p_bytes int, r_sent *int32) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_data = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		var p_bytes = gd.UnsafeGet[int64](p_args, 1)

		var r_sent = gd.UnsafeGet[*int32](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_data, int(p_bytes), r_sent)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(Error.New(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _get_available_bytes(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, int64(ret))
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.StreamPeerExtension

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("StreamPeerExtension"))
	casted := Instance{*(*gdclass.StreamPeerExtension)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (class) _get_data(impl func(ptr unsafe.Pointer, r_buffer unsafe.Pointer, r_bytes int64, r_received *int32) Error.Code) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var r_buffer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		var r_bytes = gd.UnsafeGet[int64](p_args, 1)

		var r_received = gd.UnsafeGet[*int32](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, r_buffer, r_bytes, r_received)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _get_partial_data(impl func(ptr unsafe.Pointer, r_buffer unsafe.Pointer, r_bytes int64, r_received *int32) Error.Code) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var r_buffer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		var r_bytes = gd.UnsafeGet[int64](p_args, 1)

		var r_received = gd.UnsafeGet[*int32](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, r_buffer, r_bytes, r_received)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _put_data(impl func(ptr unsafe.Pointer, p_data unsafe.Pointer, p_bytes int64, r_sent *int32) Error.Code) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_data = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		var p_bytes = gd.UnsafeGet[int64](p_args, 1)

		var r_sent = gd.UnsafeGet[*int32](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_data, p_bytes, r_sent)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _put_partial_data(impl func(ptr unsafe.Pointer, p_data unsafe.Pointer, p_bytes int64, r_sent *int32) Error.Code) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_data = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		var p_bytes = gd.UnsafeGet[int64](p_args, 1)

		var r_sent = gd.UnsafeGet[*int32](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_data, p_bytes, r_sent)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _get_available_bytes(impl func(ptr unsafe.Pointer) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (self class) AsStreamPeerExtension() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsStreamPeerExtension() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	case "_get_data":
		return reflect.ValueOf(self._get_data)
	case "_get_partial_data":
		return reflect.ValueOf(self._get_partial_data)
	case "_put_data":
		return reflect.ValueOf(self._put_data)
	case "_put_partial_data":
		return reflect.ValueOf(self._put_partial_data)
	case "_get_available_bytes":
		return reflect.ValueOf(self._get_available_bytes)
	default:
		return gd.VirtualByName(StreamPeer.Advanced(self.AsStreamPeer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_data":
		return reflect.ValueOf(self._get_data)
	case "_get_partial_data":
		return reflect.ValueOf(self._get_partial_data)
	case "_put_data":
		return reflect.ValueOf(self._put_data)
	case "_put_partial_data":
		return reflect.ValueOf(self._put_partial_data)
	case "_get_available_bytes":
		return reflect.ValueOf(self._get_available_bytes)
	default:
		return gd.VirtualByName(StreamPeer.Instance(self.AsStreamPeer()), name)
	}
}
func init() {
	gdclass.Register("StreamPeerExtension", func(ptr gd.Object) any {
		return [1]gdclass.StreamPeerExtension{*(*gdclass.StreamPeerExtension)(unsafe.Pointer(&ptr))}
	})
}
