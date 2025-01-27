// Package Marshalls provides methods for working with Marshalls object instances.
package Marshalls

import "unsafe"
import "sync"
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
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"

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

/*
Provides data transformation and encoding utility functions.
*/
var self [1]gdclass.Marshalls
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.Marshalls)
	self = *(*[1]gdclass.Marshalls)(unsafe.Pointer(&obj))
}

/*
Returns a Base64-encoded string of the [Variant] [param variant]. If [param full_objects] is [code]true[/code], encoding objects is allowed (and can potentially include code).
Internally, this uses the same encoding mechanism as the [method @GlobalScope.var_to_bytes] method.
*/
func VariantToBase64(variant any) string { //gd:Marshalls.variant_to_base64
	once.Do(singleton)
	return string(class(self).VariantToBase64(gd.NewVariant(variant), false).String())
}

/*
Returns a decoded [Variant] corresponding to the Base64-encoded string [param base64_str]. If [param allow_objects] is [code]true[/code], decoding objects is allowed.
Internally, this uses the same decoding mechanism as the [method @GlobalScope.bytes_to_var] method.
[b]Warning:[/b] Deserialized objects can contain code which gets executed. Do not use this option if the serialized object comes from untrusted sources to avoid potential security threats such as remote code execution.
*/
func Base64ToVariant(base64_str string) any { //gd:Marshalls.base64_to_variant
	once.Do(singleton)
	return any(class(self).Base64ToVariant(String.New(base64_str), false).Interface())
}

/*
Returns a Base64-encoded string of a given [PackedByteArray].
*/
func RawToBase64(array []byte) string { //gd:Marshalls.raw_to_base64
	once.Do(singleton)
	return string(class(self).RawToBase64(gd.NewPackedByteSlice(array)).String())
}

/*
Returns a decoded [PackedByteArray] corresponding to the Base64-encoded string [param base64_str].
*/
func Base64ToRaw(base64_str string) []byte { //gd:Marshalls.base64_to_raw
	once.Do(singleton)
	return []byte(class(self).Base64ToRaw(String.New(base64_str)).Bytes())
}

/*
Returns a Base64-encoded string of the UTF-8 string [param utf8_str].
*/
func Utf8ToBase64(utf8_str string) string { //gd:Marshalls.utf8_to_base64
	once.Do(singleton)
	return string(class(self).Utf8ToBase64(String.New(utf8_str)).String())
}

/*
Returns a decoded string corresponding to the Base64-encoded string [param base64_str].
*/
func Base64ToUtf8(base64_str string) string { //gd:Marshalls.base64_to_utf8
	once.Do(singleton)
	return string(class(self).Base64ToUtf8(String.New(base64_str)).String())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]gdclass.Marshalls

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Returns a Base64-encoded string of the [Variant] [param variant]. If [param full_objects] is [code]true[/code], encoding objects is allowed (and can potentially include code).
Internally, this uses the same encoding mechanism as the [method @GlobalScope.var_to_bytes] method.
*/
//go:nosplit
func (self class) VariantToBase64(variant gd.Variant, full_objects bool) String.Readable { //gd:Marshalls.variant_to_base64
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(variant))
	callframe.Arg(frame, full_objects)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Marshalls.Bind_variant_to_base64, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns a decoded [Variant] corresponding to the Base64-encoded string [param base64_str]. If [param allow_objects] is [code]true[/code], decoding objects is allowed.
Internally, this uses the same decoding mechanism as the [method @GlobalScope.bytes_to_var] method.
[b]Warning:[/b] Deserialized objects can contain code which gets executed. Do not use this option if the serialized object comes from untrusted sources to avoid potential security threats such as remote code execution.
*/
//go:nosplit
func (self class) Base64ToVariant(base64_str String.Readable, allow_objects bool) gd.Variant { //gd:Marshalls.base64_to_variant
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(base64_str)))
	callframe.Arg(frame, allow_objects)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Marshalls.Bind_base64_to_variant, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a Base64-encoded string of a given [PackedByteArray].
*/
//go:nosplit
func (self class) RawToBase64(array gd.PackedByteArray) String.Readable { //gd:Marshalls.raw_to_base64
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(array))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Marshalls.Bind_raw_to_base64, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns a decoded [PackedByteArray] corresponding to the Base64-encoded string [param base64_str].
*/
//go:nosplit
func (self class) Base64ToRaw(base64_str String.Readable) gd.PackedByteArray { //gd:Marshalls.base64_to_raw
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(base64_str)))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Marshalls.Bind_base64_to_raw, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = pointers.New[gd.PackedByteArray](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a Base64-encoded string of the UTF-8 string [param utf8_str].
*/
//go:nosplit
func (self class) Utf8ToBase64(utf8_str String.Readable) String.Readable { //gd:Marshalls.utf8_to_base64
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(utf8_str)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Marshalls.Bind_utf8_to_base64, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns a decoded string corresponding to the Base64-encoded string [param base64_str].
*/
//go:nosplit
func (self class) Base64ToUtf8(base64_str String.Readable) String.Readable { //gd:Marshalls.base64_to_utf8
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(base64_str)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Marshalls.Bind_base64_to_utf8, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("Marshalls", func(ptr gd.Object) any { return [1]gdclass.Marshalls{*(*gdclass.Marshalls)(unsafe.Pointer(&ptr))} })
}
