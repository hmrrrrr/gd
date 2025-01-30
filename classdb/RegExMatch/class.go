// Package RegExMatch provides methods for working with RegExMatch object instances.
package RegExMatch

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
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
Contains the results of a single [RegEx] match returned by [method RegEx.search] and [method RegEx.search_all]. It can be used to find the position and range of the match and its capturing groups, and it can extract its substring for you.
*/
type Instance [1]gdclass.RegExMatch

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRegExMatch() Instance
}

/*
Returns the number of capturing groups.
*/
func (self Instance) GetGroupCount() int { //gd:RegExMatch.get_group_count
	return int(int(class(self).GetGroupCount()))
}

/*
Returns the substring of the match from the source string. Capturing groups can be retrieved by providing its group number as an integer or its string name (if it's a named group). The default value of 0 refers to the whole pattern.
Returns an empty string if the group did not match or doesn't exist.
*/
func (self Instance) GetString() string { //gd:RegExMatch.get_string
	return string(class(self).GetString(variant.New(0)).String())
}

/*
Returns the starting position of the match within the source string. The starting position of capturing groups can be retrieved by providing its group number as an integer or its string name (if it's a named group). The default value of 0 refers to the whole pattern.
Returns -1 if the group did not match or doesn't exist.
*/
func (self Instance) GetStart() int { //gd:RegExMatch.get_start
	return int(int(class(self).GetStart(variant.New(0))))
}

/*
Returns the end position of the match within the source string. The end position of capturing groups can be retrieved by providing its group number as an integer or its string name (if it's a named group). The default value of 0 refers to the whole pattern.
Returns -1 if the group did not match or doesn't exist.
*/
func (self Instance) GetEnd() int { //gd:RegExMatch.get_end
	return int(int(class(self).GetEnd(variant.New(0))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RegExMatch

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RegExMatch"))
	casted := Instance{*(*gdclass.RegExMatch)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Subject() string {
	return string(class(self).GetSubject().String())
}

func (self Instance) Names() map[any]any {
	return map[any]any(gd.DictionaryAs[map[any]any](class(self).GetNames()))
}

func (self Instance) Strings() []string {
	return []string(class(self).GetStrings().Strings())
}

//go:nosplit
func (self class) GetSubject() String.Readable { //gd:RegExMatch.get_subject
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RegExMatch.Bind_get_subject, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the number of capturing groups.
*/
//go:nosplit
func (self class) GetGroupCount() int64 { //gd:RegExMatch.get_group_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RegExMatch.Bind_get_group_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetNames() Dictionary.Any { //gd:RegExMatch.get_names
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RegExMatch.Bind_get_names, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetStrings() Packed.Strings { //gd:RegExMatch.get_strings
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RegExMatch.Bind_get_strings, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the substring of the match from the source string. Capturing groups can be retrieved by providing its group number as an integer or its string name (if it's a named group). The default value of 0 refers to the whole pattern.
Returns an empty string if the group did not match or doesn't exist.
*/
//go:nosplit
func (self class) GetString(name variant.Any) String.Readable { //gd:RegExMatch.get_string
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(name)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RegExMatch.Bind_get_string, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the starting position of the match within the source string. The starting position of capturing groups can be retrieved by providing its group number as an integer or its string name (if it's a named group). The default value of 0 refers to the whole pattern.
Returns -1 if the group did not match or doesn't exist.
*/
//go:nosplit
func (self class) GetStart(name variant.Any) int64 { //gd:RegExMatch.get_start
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(name)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RegExMatch.Bind_get_start, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the end position of the match within the source string. The end position of capturing groups can be retrieved by providing its group number as an integer or its string name (if it's a named group). The default value of 0 refers to the whole pattern.
Returns -1 if the group did not match or doesn't exist.
*/
//go:nosplit
func (self class) GetEnd(name variant.Any) int64 { //gd:RegExMatch.get_end
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(name)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RegExMatch.Bind_get_end, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRegExMatch() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRegExMatch() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("RegExMatch", func(ptr gd.Object) any { return [1]gdclass.RegExMatch{*(*gdclass.RegExMatch)(unsafe.Pointer(&ptr))} })
}
