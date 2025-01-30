// Package CryptoKey provides methods for working with CryptoKey object instances.
package CryptoKey

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Resource"
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
The CryptoKey class represents a cryptographic key. Keys can be loaded and saved like any other [Resource].
They can be used to generate a self-signed [X509Certificate] via [method Crypto.generate_self_signed_certificate] and as private key in [method StreamPeerTLS.accept_stream] along with the appropriate certificate.
*/
type Instance [1]gdclass.CryptoKey

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCryptoKey() Instance
}

/*
Saves a key to the given [param path]. If [param public_only] is [code]true[/code], only the public key will be saved.
[b]Note:[/b] [param path] should be a "*.pub" file if [param public_only] is [code]true[/code], a "*.key" file otherwise.
*/
func (self Instance) Save(path string) error { //gd:CryptoKey.save
	return error(gd.ToError(class(self).Save(String.New(path), false)))
}

/*
Loads a key from [param path]. If [param public_only] is [code]true[/code], only the public key will be loaded.
[b]Note:[/b] [param path] should be a "*.pub" file if [param public_only] is [code]true[/code], a "*.key" file otherwise.
*/
func (self Instance) Load(path string) error { //gd:CryptoKey.load
	return error(gd.ToError(class(self).Load(String.New(path), false)))
}

/*
Returns [code]true[/code] if this CryptoKey only has the public part, and not the private one.
*/
func (self Instance) IsPublicOnly() bool { //gd:CryptoKey.is_public_only
	return bool(class(self).IsPublicOnly())
}

/*
Returns a string containing the key in PEM format. If [param public_only] is [code]true[/code], only the public key will be included.
*/
func (self Instance) SaveToString() string { //gd:CryptoKey.save_to_string
	return string(class(self).SaveToString(false).String())
}

/*
Loads a key from the given [param string_key]. If [param public_only] is [code]true[/code], only the public key will be loaded.
*/
func (self Instance) LoadFromString(string_key string) error { //gd:CryptoKey.load_from_string
	return error(gd.ToError(class(self).LoadFromString(String.New(string_key), false)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.CryptoKey

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CryptoKey"))
	casted := Instance{*(*gdclass.CryptoKey)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Saves a key to the given [param path]. If [param public_only] is [code]true[/code], only the public key will be saved.
[b]Note:[/b] [param path] should be a "*.pub" file if [param public_only] is [code]true[/code], a "*.key" file otherwise.
*/
//go:nosplit
func (self class) Save(path String.Readable, public_only bool) Error.Code { //gd:CryptoKey.save
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	callframe.Arg(frame, public_only)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CryptoKey.Bind_save, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Loads a key from [param path]. If [param public_only] is [code]true[/code], only the public key will be loaded.
[b]Note:[/b] [param path] should be a "*.pub" file if [param public_only] is [code]true[/code], a "*.key" file otherwise.
*/
//go:nosplit
func (self class) Load(path String.Readable, public_only bool) Error.Code { //gd:CryptoKey.load
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	callframe.Arg(frame, public_only)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CryptoKey.Bind_load, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if this CryptoKey only has the public part, and not the private one.
*/
//go:nosplit
func (self class) IsPublicOnly() bool { //gd:CryptoKey.is_public_only
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CryptoKey.Bind_is_public_only, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a string containing the key in PEM format. If [param public_only] is [code]true[/code], only the public key will be included.
*/
//go:nosplit
func (self class) SaveToString(public_only bool) String.Readable { //gd:CryptoKey.save_to_string
	var frame = callframe.New()
	callframe.Arg(frame, public_only)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CryptoKey.Bind_save_to_string, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Loads a key from the given [param string_key]. If [param public_only] is [code]true[/code], only the public key will be loaded.
*/
//go:nosplit
func (self class) LoadFromString(string_key String.Readable, public_only bool) Error.Code { //gd:CryptoKey.load_from_string
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(string_key)))
	callframe.Arg(frame, public_only)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CryptoKey.Bind_load_from_string, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsCryptoKey() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCryptoKey() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("CryptoKey", func(ptr gd.Object) any { return [1]gdclass.CryptoKey{*(*gdclass.CryptoKey)(unsafe.Pointer(&ptr))} })
}
