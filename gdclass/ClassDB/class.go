package ClassDB

import "unsafe"
import "sync"
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
Provides access to metadata stored for every available class.

*/
var self gdclass.ClassDB
var once sync.Once
func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.ClassDB)
	self = *(*gdclass.ClassDB)(unsafe.Pointer(&obj))
}

/*
Returns the names of all the classes available.
*/
func GetClassList() []string {
	once.Do(singleton)
	return []string(class(self).GetClassList().Strings())
}

/*
Returns the names of all the classes that directly or indirectly inherit from [param class].
*/
func GetInheritersFromClass(class_ string) []string {
	once.Do(singleton)
	return []string(class(self).GetInheritersFromClass(gd.NewStringName(class_)).Strings())
}

/*
Returns the parent class of [param class].
*/
func GetParentClass(class_ string) string {
	once.Do(singleton)
	return string(class(self).GetParentClass(gd.NewStringName(class_)).String())
}

/*
Returns whether the specified [param class] is available or not.
*/
func ClassExists(class_ string) bool {
	once.Do(singleton)
	return bool(class(self).ClassExists(gd.NewStringName(class_)))
}

/*
Returns whether [param inherits] is an ancestor of [param class] or not.
*/
func IsParentClass(class_ string, inherits string) bool {
	once.Do(singleton)
	return bool(class(self).IsParentClass(gd.NewStringName(class_), gd.NewStringName(inherits)))
}

/*
Returns [code]true[/code] if objects can be instantiated from the specified [param class], otherwise returns [code]false[/code].
*/
func CanInstantiate(class_ string) bool {
	once.Do(singleton)
	return bool(class(self).CanInstantiate(gd.NewStringName(class_)))
}

/*
Creates an instance of [param class].
*/
func Instantiate(class_ string) gd.Variant {
	once.Do(singleton)
	return gd.Variant(class(self).Instantiate(gd.NewStringName(class_)))
}

/*
Returns whether [param class] or its ancestry has a signal called [param signal] or not.
*/
func ClassHasSignal(class_ string, signal string) bool {
	once.Do(singleton)
	return bool(class(self).ClassHasSignal(gd.NewStringName(class_), gd.NewStringName(signal)))
}

/*
Returns the [param signal] data of [param class] or its ancestry. The returned value is a [Dictionary] with the following keys: [code]args[/code], [code]default_args[/code], [code]flags[/code], [code]id[/code], [code]name[/code], [code]return: (class_name, hint, hint_string, name, type, usage)[/code].
*/
func ClassGetSignal(class_ string, signal string) gd.Dictionary {
	once.Do(singleton)
	return gd.Dictionary(class(self).ClassGetSignal(gd.NewStringName(class_), gd.NewStringName(signal)))
}

/*
Returns an array with all the signals of [param class] or its ancestry if [param no_inheritance] is [code]false[/code]. Every element of the array is a [Dictionary] as described in [method class_get_signal].
*/
func ClassGetSignalList(class_ string) gd.Array {
	once.Do(singleton)
	return gd.Array(class(self).ClassGetSignalList(gd.NewStringName(class_), false))
}

/*
Returns an array with all the properties of [param class] or its ancestry if [param no_inheritance] is [code]false[/code].
*/
func ClassGetPropertyList(class_ string) gd.Array {
	once.Do(singleton)
	return gd.Array(class(self).ClassGetPropertyList(gd.NewStringName(class_), false))
}

/*
Returns the value of [param property] of [param object] or its ancestry.
*/
func ClassGetProperty(obj gd.Object, property string) gd.Variant {
	once.Do(singleton)
	return gd.Variant(class(self).ClassGetProperty(obj, gd.NewStringName(property)))
}

/*
Sets [param property] value of [param object] to [param value].
*/
func ClassSetProperty(obj gd.Object, property string, value gd.Variant) gd.Error {
	once.Do(singleton)
	return gd.Error(class(self).ClassSetProperty(obj, gd.NewStringName(property), value))
}

/*
Returns the default value of [param property] of [param class] or its ancestor classes.
*/
func ClassGetPropertyDefaultValue(class_ string, property string) gd.Variant {
	once.Do(singleton)
	return gd.Variant(class(self).ClassGetPropertyDefaultValue(gd.NewStringName(class_), gd.NewStringName(property)))
}

/*
Returns whether [param class] (or its ancestry if [param no_inheritance] is [code]false[/code]) has a method called [param method] or not.
*/
func ClassHasMethod(class_ string, method string) bool {
	once.Do(singleton)
	return bool(class(self).ClassHasMethod(gd.NewStringName(class_), gd.NewStringName(method), false))
}

/*
Returns the number of arguments of the method [param method] of [param class] or its ancestry if [param no_inheritance] is [code]false[/code].
*/
func ClassGetMethodArgumentCount(class_ string, method string) int {
	once.Do(singleton)
	return int(int(class(self).ClassGetMethodArgumentCount(gd.NewStringName(class_), gd.NewStringName(method), false)))
}

/*
Returns an array with all the methods of [param class] or its ancestry if [param no_inheritance] is [code]false[/code]. Every element of the array is a [Dictionary] with the following keys: [code]args[/code], [code]default_args[/code], [code]flags[/code], [code]id[/code], [code]name[/code], [code]return: (class_name, hint, hint_string, name, type, usage)[/code].
[b]Note:[/b] In exported release builds the debug info is not available, so the returned dictionaries will contain only method names.
*/
func ClassGetMethodList(class_ string) gd.Array {
	once.Do(singleton)
	return gd.Array(class(self).ClassGetMethodList(gd.NewStringName(class_), false))
}

/*
Returns an array with the names all the integer constants of [param class] or its ancestry.
*/
func ClassGetIntegerConstantList(class_ string) []string {
	once.Do(singleton)
	return []string(class(self).ClassGetIntegerConstantList(gd.NewStringName(class_), false).Strings())
}

/*
Returns whether [param class] or its ancestry has an integer constant called [param name] or not.
*/
func ClassHasIntegerConstant(class_ string, name string) bool {
	once.Do(singleton)
	return bool(class(self).ClassHasIntegerConstant(gd.NewStringName(class_), gd.NewStringName(name)))
}

/*
Returns the value of the integer constant [param name] of [param class] or its ancestry. Always returns 0 when the constant could not be found.
*/
func ClassGetIntegerConstant(class_ string, name string) int {
	once.Do(singleton)
	return int(int(class(self).ClassGetIntegerConstant(gd.NewStringName(class_), gd.NewStringName(name))))
}

/*
Returns whether [param class] or its ancestry has an enum called [param name] or not.
*/
func ClassHasEnum(class_ string, name string) bool {
	once.Do(singleton)
	return bool(class(self).ClassHasEnum(gd.NewStringName(class_), gd.NewStringName(name), false))
}

/*
Returns an array with all the enums of [param class] or its ancestry.
*/
func ClassGetEnumList(class_ string) []string {
	once.Do(singleton)
	return []string(class(self).ClassGetEnumList(gd.NewStringName(class_), false).Strings())
}

/*
Returns an array with all the keys in [param enum] of [param class] or its ancestry.
*/
func ClassGetEnumConstants(class_ string, enum string) []string {
	once.Do(singleton)
	return []string(class(self).ClassGetEnumConstants(gd.NewStringName(class_), gd.NewStringName(enum), false).Strings())
}

/*
Returns which enum the integer constant [param name] of [param class] or its ancestry belongs to.
*/
func ClassGetIntegerConstantEnum(class_ string, name string) string {
	once.Do(singleton)
	return string(class(self).ClassGetIntegerConstantEnum(gd.NewStringName(class_), gd.NewStringName(name), false).String())
}

/*
Returns whether [param class] (or its ancestor classes if [param no_inheritance] is [code]false[/code]) has an enum called [param enum] that is a bitfield.
*/
func IsClassEnumBitfield(class_ string, enum string) bool {
	once.Do(singleton)
	return bool(class(self).IsClassEnumBitfield(gd.NewStringName(class_), gd.NewStringName(enum), false))
}

/*
Returns whether this [param class] is enabled or not.
*/
func IsClassEnabled(class_ string) bool {
	once.Do(singleton)
	return bool(class(self).IsClassEnabled(gd.NewStringName(class_)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func GD() class { once.Do(singleton); return self }
type class [1]classdb.ClassDB
func (self class) AsObject() gd.Object { return self[0].AsObject() }

/*
Returns the names of all the classes available.
*/
//go:nosplit
func (self class) GetClassList() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_get_class_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the names of all the classes that directly or indirectly inherit from [param class].
*/
//go:nosplit
func (self class) GetInheritersFromClass(class_ gd.StringName) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(class_))
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_get_inheriters_from_class, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the parent class of [param class].
*/
//go:nosplit
func (self class) GetParentClass(class_ gd.StringName) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(class_))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_get_parent_class, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns whether the specified [param class] is available or not.
*/
//go:nosplit
func (self class) ClassExists(class_ gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(class_))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_exists, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns whether [param inherits] is an ancestor of [param class] or not.
*/
//go:nosplit
func (self class) IsParentClass(class_ gd.StringName, inherits gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(class_))
	callframe.Arg(frame, discreet.Get(inherits))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_is_parent_class, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if objects can be instantiated from the specified [param class], otherwise returns [code]false[/code].
*/
//go:nosplit
func (self class) CanInstantiate(class_ gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(class_))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_can_instantiate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates an instance of [param class].
*/
//go:nosplit
func (self class) Instantiate(class_ gd.StringName) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(class_))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_instantiate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns whether [param class] or its ancestry has a signal called [param signal] or not.
*/
//go:nosplit
func (self class) ClassHasSignal(class_ gd.StringName, signal gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(class_))
	callframe.Arg(frame, discreet.Get(signal))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_has_signal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [param signal] data of [param class] or its ancestry. The returned value is a [Dictionary] with the following keys: [code]args[/code], [code]default_args[/code], [code]flags[/code], [code]id[/code], [code]name[/code], [code]return: (class_name, hint, hint_string, name, type, usage)[/code].
*/
//go:nosplit
func (self class) ClassGetSignal(class_ gd.StringName, signal gd.StringName) gd.Dictionary {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(class_))
	callframe.Arg(frame, discreet.Get(signal))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_signal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an array with all the signals of [param class] or its ancestry if [param no_inheritance] is [code]false[/code]. Every element of the array is a [Dictionary] as described in [method class_get_signal].
*/
//go:nosplit
func (self class) ClassGetSignalList(class_ gd.StringName, no_inheritance bool) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(class_))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_signal_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an array with all the properties of [param class] or its ancestry if [param no_inheritance] is [code]false[/code].
*/
//go:nosplit
func (self class) ClassGetPropertyList(class_ gd.StringName, no_inheritance bool) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(class_))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_property_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the value of [param property] of [param object] or its ancestry.
*/
//go:nosplit
func (self class) ClassGetProperty(obj gd.Object, property gd.StringName) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(obj))
	callframe.Arg(frame, discreet.Get(property))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_property, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets [param property] value of [param object] to [param value].
*/
//go:nosplit
func (self class) ClassSetProperty(obj gd.Object, property gd.StringName, value gd.Variant) int64 {
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(obj))
	callframe.Arg(frame, discreet.Get(property))
	callframe.Arg(frame, discreet.Get(value))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_set_property, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the default value of [param property] of [param class] or its ancestor classes.
*/
//go:nosplit
func (self class) ClassGetPropertyDefaultValue(class_ gd.StringName, property gd.StringName) gd.Variant {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(class_))
	callframe.Arg(frame, discreet.Get(property))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_property_default_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns whether [param class] (or its ancestry if [param no_inheritance] is [code]false[/code]) has a method called [param method] or not.
*/
//go:nosplit
func (self class) ClassHasMethod(class_ gd.StringName, method gd.StringName, no_inheritance bool) bool {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(class_))
	callframe.Arg(frame, discreet.Get(method))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_has_method, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of arguments of the method [param method] of [param class] or its ancestry if [param no_inheritance] is [code]false[/code].
*/
//go:nosplit
func (self class) ClassGetMethodArgumentCount(class_ gd.StringName, method gd.StringName, no_inheritance bool) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(class_))
	callframe.Arg(frame, discreet.Get(method))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_method_argument_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an array with all the methods of [param class] or its ancestry if [param no_inheritance] is [code]false[/code]. Every element of the array is a [Dictionary] with the following keys: [code]args[/code], [code]default_args[/code], [code]flags[/code], [code]id[/code], [code]name[/code], [code]return: (class_name, hint, hint_string, name, type, usage)[/code].
[b]Note:[/b] In exported release builds the debug info is not available, so the returned dictionaries will contain only method names.
*/
//go:nosplit
func (self class) ClassGetMethodList(class_ gd.StringName, no_inheritance bool) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(class_))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_method_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an array with the names all the integer constants of [param class] or its ancestry.
*/
//go:nosplit
func (self class) ClassGetIntegerConstantList(class_ gd.StringName, no_inheritance bool) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(class_))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_integer_constant_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns whether [param class] or its ancestry has an integer constant called [param name] or not.
*/
//go:nosplit
func (self class) ClassHasIntegerConstant(class_ gd.StringName, name gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(class_))
	callframe.Arg(frame, discreet.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_has_integer_constant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the value of the integer constant [param name] of [param class] or its ancestry. Always returns 0 when the constant could not be found.
*/
//go:nosplit
func (self class) ClassGetIntegerConstant(class_ gd.StringName, name gd.StringName) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(class_))
	callframe.Arg(frame, discreet.Get(name))
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_integer_constant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns whether [param class] or its ancestry has an enum called [param name] or not.
*/
//go:nosplit
func (self class) ClassHasEnum(class_ gd.StringName, name gd.StringName, no_inheritance bool) bool {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(class_))
	callframe.Arg(frame, discreet.Get(name))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_has_enum, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an array with all the enums of [param class] or its ancestry.
*/
//go:nosplit
func (self class) ClassGetEnumList(class_ gd.StringName, no_inheritance bool) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(class_))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_enum_list, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns an array with all the keys in [param enum] of [param class] or its ancestry.
*/
//go:nosplit
func (self class) ClassGetEnumConstants(class_ gd.StringName, enum gd.StringName, no_inheritance bool) gd.PackedStringArray {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(class_))
	callframe.Arg(frame, discreet.Get(enum))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_enum_constants, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns which enum the integer constant [param name] of [param class] or its ancestry belongs to.
*/
//go:nosplit
func (self class) ClassGetIntegerConstantEnum(class_ gd.StringName, name gd.StringName, no_inheritance bool) gd.StringName {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(class_))
	callframe.Arg(frame, discreet.Get(name))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_integer_constant_enum, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.StringName](r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns whether [param class] (or its ancestor classes if [param no_inheritance] is [code]false[/code]) has an enum called [param enum] that is a bitfield.
*/
//go:nosplit
func (self class) IsClassEnumBitfield(class_ gd.StringName, enum gd.StringName, no_inheritance bool) bool {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(class_))
	callframe.Arg(frame, discreet.Get(enum))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_is_class_enum_bitfield, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns whether this [param class] is enabled or not.
*/
//go:nosplit
func (self class) IsClassEnabled(class_ gd.StringName) bool {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(class_))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_is_class_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {classdb.Register("ClassDB", func(ptr gd.Object) any { return classdb.ClassDB(ptr) })}
