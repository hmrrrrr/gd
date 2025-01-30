// Package X509Certificate provides methods for working with X509Certificate object instances.
package X509Certificate

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
The X509Certificate class represents an X509 certificate. Certificates can be loaded and saved like any other [Resource].
They can be used as the server certificate in [method StreamPeerTLS.accept_stream] (along with the proper [CryptoKey]), and to specify the only certificate that should be accepted when connecting to a TLS server via [method StreamPeerTLS.connect_to_stream].
*/
type Instance [1]gdclass.X509Certificate

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsX509Certificate() Instance
}

/*
Saves a certificate to the given [param path] (should be a "*.crt" file).
*/
func (self Instance) Save(path string) error { //gd:X509Certificate.save
	return error(gd.ToError(class(self).Save(String.New(path))))
}

/*
Loads a certificate from [param path] ("*.crt" file).
*/
func (self Instance) Load(path string) error { //gd:X509Certificate.load
	return error(gd.ToError(class(self).Load(String.New(path))))
}

/*
Returns a string representation of the certificate, or an empty string if the certificate is invalid.
*/
func (self Instance) SaveToString() string { //gd:X509Certificate.save_to_string
	return string(class(self).SaveToString().String())
}

/*
Loads a certificate from the given [param string].
*/
func (self Instance) LoadFromString(s string) error { //gd:X509Certificate.load_from_string
	return error(gd.ToError(class(self).LoadFromString(String.New(s))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.X509Certificate

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("X509Certificate"))
	casted := Instance{*(*gdclass.X509Certificate)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Saves a certificate to the given [param path] (should be a "*.crt" file).
*/
//go:nosplit
func (self class) Save(path String.Readable) Error.Code { //gd:X509Certificate.save
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.X509Certificate.Bind_save, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Loads a certificate from [param path] ("*.crt" file).
*/
//go:nosplit
func (self class) Load(path String.Readable) Error.Code { //gd:X509Certificate.load
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.X509Certificate.Bind_load, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a string representation of the certificate, or an empty string if the certificate is invalid.
*/
//go:nosplit
func (self class) SaveToString() String.Readable { //gd:X509Certificate.save_to_string
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.X509Certificate.Bind_save_to_string, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Loads a certificate from the given [param string].
*/
//go:nosplit
func (self class) LoadFromString(s String.Readable) Error.Code { //gd:X509Certificate.load_from_string
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(s)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.X509Certificate.Bind_load_from_string, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsX509Certificate() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsX509Certificate() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("X509Certificate", func(ptr gd.Object) any {
		return [1]gdclass.X509Certificate{*(*gdclass.X509Certificate)(unsafe.Pointer(&ptr))}
	})
}
