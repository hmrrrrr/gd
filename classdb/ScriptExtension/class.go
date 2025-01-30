// Package ScriptExtension provides methods for working with ScriptExtension object instances.
package ScriptExtension

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Resource"
import "graphics.gd/classdb/Script"
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

type Instance [1]gdclass.ScriptExtension

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsScriptExtension() Instance
}
type Interface interface {
	EditorCanReloadFromFile() bool
	PlaceholderErased(placeholder unsafe.Pointer)
	CanInstantiate() bool
	GetBaseScript() [1]gdclass.Script
	GetGlobalName() string
	InheritsScript(script [1]gdclass.Script) bool
	GetInstanceBaseType() string
	InstanceCreate(for_object Object.Instance) unsafe.Pointer
	PlaceholderInstanceCreate(for_object Object.Instance) unsafe.Pointer
	InstanceHas(obj Object.Instance) bool
	HasSourceCode() bool
	GetSourceCode() string
	SetSourceCode(code string)
	Reload(keep_state bool) error
	GetDocumentation() []map[any]any
	GetClassIconPath() string
	HasMethod(method string) bool
	HasStaticMethod(method string) bool
	//Return the expected argument count for the given [param method], or [code]null[/code] if it can't be determined (which will then fall back to the default behavior).
	GetScriptMethodArgumentCount(method string) any
	GetMethodInfo(method string) map[any]any
	IsTool() bool
	IsValid() bool
	//Returns [code]true[/code] if the script is an abstract script. An abstract script does not have a constructor and cannot be instantiated.
	IsAbstract() bool
	GetLanguage() [1]gdclass.ScriptLanguage
	HasScriptSignal(signal string) bool
	GetScriptSignalList() []map[any]any
	HasPropertyDefaultValue(property string) bool
	GetPropertyDefaultValue(property string) any
	UpdateExports()
	GetScriptMethodList() []map[any]any
	GetScriptPropertyList() []map[any]any
	GetMemberLine(member string) int
	GetConstants() map[any]any
	GetMembers() []string
	IsPlaceholderFallbackEnabled() bool
	GetRpcConfig() any
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) EditorCanReloadFromFile() (_ bool)                            { return }
func (self implementation) PlaceholderErased(placeholder unsafe.Pointer)                 { return }
func (self implementation) CanInstantiate() (_ bool)                                     { return }
func (self implementation) GetBaseScript() (_ [1]gdclass.Script)                         { return }
func (self implementation) GetGlobalName() (_ string)                                    { return }
func (self implementation) InheritsScript(script [1]gdclass.Script) (_ bool)             { return }
func (self implementation) GetInstanceBaseType() (_ string)                              { return }
func (self implementation) InstanceCreate(for_object Object.Instance) (_ unsafe.Pointer) { return }
func (self implementation) PlaceholderInstanceCreate(for_object Object.Instance) (_ unsafe.Pointer) {
	return
}
func (self implementation) InstanceHas(obj Object.Instance) (_ bool)           { return }
func (self implementation) HasSourceCode() (_ bool)                            { return }
func (self implementation) GetSourceCode() (_ string)                          { return }
func (self implementation) SetSourceCode(code string)                          { return }
func (self implementation) Reload(keep_state bool) (_ error)                   { return }
func (self implementation) GetDocumentation() (_ []map[any]any)                { return }
func (self implementation) GetClassIconPath() (_ string)                       { return }
func (self implementation) HasMethod(method string) (_ bool)                   { return }
func (self implementation) HasStaticMethod(method string) (_ bool)             { return }
func (self implementation) GetScriptMethodArgumentCount(method string) (_ any) { return }
func (self implementation) GetMethodInfo(method string) (_ map[any]any)        { return }
func (self implementation) IsTool() (_ bool)                                   { return }
func (self implementation) IsValid() (_ bool)                                  { return }
func (self implementation) IsAbstract() (_ bool)                               { return }
func (self implementation) GetLanguage() (_ [1]gdclass.ScriptLanguage)         { return }
func (self implementation) HasScriptSignal(signal string) (_ bool)             { return }
func (self implementation) GetScriptSignalList() (_ []map[any]any)             { return }
func (self implementation) HasPropertyDefaultValue(property string) (_ bool)   { return }
func (self implementation) GetPropertyDefaultValue(property string) (_ any)    { return }
func (self implementation) UpdateExports()                                     { return }
func (self implementation) GetScriptMethodList() (_ []map[any]any)             { return }
func (self implementation) GetScriptPropertyList() (_ []map[any]any)           { return }
func (self implementation) GetMemberLine(member string) (_ int)                { return }
func (self implementation) GetConstants() (_ map[any]any)                      { return }
func (self implementation) GetMembers() (_ []string)                           { return }
func (self implementation) IsPlaceholderFallbackEnabled() (_ bool)             { return }
func (self implementation) GetRpcConfig() (_ any)                              { return }
func (Instance) _editor_can_reload_from_file(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _placeholder_erased(impl func(ptr unsafe.Pointer, placeholder unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var placeholder = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, placeholder)
	}
}
func (Instance) _can_instantiate(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_base_script(impl func(ptr unsafe.Pointer) [1]gdclass.Script) (cb gd.ExtensionClassCallVirtualFunc) {
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
func (Instance) _get_global_name(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalStringName(String.Name(String.New(ret))))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _inherits_script(impl func(ptr unsafe.Pointer, script [1]gdclass.Script) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var script = [1]gdclass.Script{pointers.New[gdclass.Script]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(script[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, script)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_instance_base_type(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalStringName(String.Name(String.New(ret))))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _instance_create(impl func(ptr unsafe.Pointer, for_object Object.Instance) unsafe.Pointer) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var for_object = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(for_object[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, for_object)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _placeholder_instance_create(impl func(ptr unsafe.Pointer, for_object Object.Instance) unsafe.Pointer) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var for_object = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(for_object[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, for_object)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _instance_has(impl func(ptr unsafe.Pointer, obj Object.Instance) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var obj = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(obj[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _has_source_code(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_source_code(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _set_source_code(impl func(ptr unsafe.Pointer, code string)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var code = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(code))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, code.String())
	}
}
func (Instance) _reload(impl func(ptr unsafe.Pointer, keep_state bool) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var keep_state = gd.UnsafeGet[bool](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, keep_state)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(Error.New(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _get_documentation(impl func(ptr unsafe.Pointer) []map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalArray(gd.ArrayFromSlice[Array.Contains[Dictionary.Any]](ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _get_class_icon_path(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _has_method(impl func(ptr unsafe.Pointer, method string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var method = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0)))))
		defer pointers.End(gd.InternalStringName(method))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, method.String())
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _has_static_method(impl func(ptr unsafe.Pointer, method string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var method = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0)))))
		defer pointers.End(gd.InternalStringName(method))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, method.String())
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Return the expected argument count for the given [param method], or [code]null[/code] if it can't be determined (which will then fall back to the default behavior).
*/
func (Instance) _get_script_method_argument_count(impl func(ptr unsafe.Pointer, method string) any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var method = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0)))))
		defer pointers.End(gd.InternalStringName(method))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, method.String())
		ptr, ok := pointers.End(gd.InternalVariant(variant.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _get_method_info(impl func(ptr unsafe.Pointer, method string) map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var method = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0)))))
		defer pointers.End(gd.InternalStringName(method))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, method.String())
		ptr, ok := pointers.End(gd.InternalDictionary(gd.DictionaryFromMap(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _is_tool(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _is_valid(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns [code]true[/code] if the script is an abstract script. An abstract script does not have a constructor and cannot be instantiated.
*/
func (Instance) _is_abstract(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_language(impl func(ptr unsafe.Pointer) [1]gdclass.ScriptLanguage) (cb gd.ExtensionClassCallVirtualFunc) {
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
func (Instance) _has_script_signal(impl func(ptr unsafe.Pointer, signal string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var signal = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0)))))
		defer pointers.End(gd.InternalStringName(signal))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, signal.String())
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_script_signal_list(impl func(ptr unsafe.Pointer) []map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalArray(gd.ArrayFromSlice[Array.Contains[Dictionary.Any]](ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _has_property_default_value(impl func(ptr unsafe.Pointer, property string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var property = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0)))))
		defer pointers.End(gd.InternalStringName(property))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, property.String())
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_property_default_value(impl func(ptr unsafe.Pointer, property string) any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var property = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0)))))
		defer pointers.End(gd.InternalStringName(property))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, property.String())
		ptr, ok := pointers.End(gd.InternalVariant(variant.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _update_exports(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}
func (Instance) _get_script_method_list(impl func(ptr unsafe.Pointer) []map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalArray(gd.ArrayFromSlice[Array.Contains[Dictionary.Any]](ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _get_script_property_list(impl func(ptr unsafe.Pointer) []map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalArray(gd.ArrayFromSlice[Array.Contains[Dictionary.Any]](ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _get_member_line(impl func(ptr unsafe.Pointer, member string) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var member = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0)))))
		defer pointers.End(gd.InternalStringName(member))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, member.String())
		gd.UnsafeSet(p_back, int64(ret))
	}
}
func (Instance) _get_constants(impl func(ptr unsafe.Pointer) map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalDictionary(gd.DictionaryFromMap(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _get_members(impl func(ptr unsafe.Pointer) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalArray(gd.ArrayFromSlice[Array.Contains[String.Name]](ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _is_placeholder_fallback_enabled(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_rpc_config(impl func(ptr unsafe.Pointer) any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalVariant(variant.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ScriptExtension

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ScriptExtension"))
	casted := Instance{*(*gdclass.ScriptExtension)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (class) _editor_can_reload_from_file(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _placeholder_erased(impl func(ptr unsafe.Pointer, placeholder unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var placeholder = gd.UnsafeGet[unsafe.Pointer](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, placeholder)
	}
}

func (class) _can_instantiate(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_base_script(impl func(ptr unsafe.Pointer) [1]gdclass.Script) (cb gd.ExtensionClassCallVirtualFunc) {
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

func (class) _get_global_name(impl func(ptr unsafe.Pointer) String.Name) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalStringName(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _inherits_script(impl func(ptr unsafe.Pointer, script [1]gdclass.Script) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var script = [1]gdclass.Script{pointers.New[gdclass.Script]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(script[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, script)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_instance_base_type(impl func(ptr unsafe.Pointer) String.Name) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalStringName(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _instance_create(impl func(ptr unsafe.Pointer, for_object [1]gd.Object) unsafe.Pointer) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var for_object = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(for_object[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, for_object)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _placeholder_instance_create(impl func(ptr unsafe.Pointer, for_object [1]gd.Object) unsafe.Pointer) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var for_object = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(for_object[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, for_object)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _instance_has(impl func(ptr unsafe.Pointer, obj [1]gd.Object) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var obj = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}
		defer pointers.End(obj[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _has_source_code(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_source_code(impl func(ptr unsafe.Pointer) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _set_source_code(impl func(ptr unsafe.Pointer, code String.Readable)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var code = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(code))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, code)
	}
}

func (class) _reload(impl func(ptr unsafe.Pointer, keep_state bool) Error.Code) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var keep_state = gd.UnsafeGet[bool](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, keep_state)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _get_documentation(impl func(ptr unsafe.Pointer) Array.Contains[Dictionary.Any]) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalArray(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _get_class_icon_path(impl func(ptr unsafe.Pointer) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _has_method(impl func(ptr unsafe.Pointer, method String.Name) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var method = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0)))))
		defer pointers.End(gd.InternalStringName(method))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, method)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _has_static_method(impl func(ptr unsafe.Pointer, method String.Name) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var method = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0)))))
		defer pointers.End(gd.InternalStringName(method))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, method)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Return the expected argument count for the given [param method], or [code]null[/code] if it can't be determined (which will then fall back to the default behavior).
*/
func (class) _get_script_method_argument_count(impl func(ptr unsafe.Pointer, method String.Name) variant.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var method = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0)))))
		defer pointers.End(gd.InternalStringName(method))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, method)
		ptr, ok := pointers.End(gd.InternalVariant(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _get_method_info(impl func(ptr unsafe.Pointer, method String.Name) Dictionary.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var method = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0)))))
		defer pointers.End(gd.InternalStringName(method))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, method)
		ptr, ok := pointers.End(gd.InternalDictionary(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _is_tool(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _is_valid(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns [code]true[/code] if the script is an abstract script. An abstract script does not have a constructor and cannot be instantiated.
*/
func (class) _is_abstract(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_language(impl func(ptr unsafe.Pointer) [1]gdclass.ScriptLanguage) (cb gd.ExtensionClassCallVirtualFunc) {
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

func (class) _has_script_signal(impl func(ptr unsafe.Pointer, signal String.Name) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var signal = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0)))))
		defer pointers.End(gd.InternalStringName(signal))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, signal)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_script_signal_list(impl func(ptr unsafe.Pointer) Array.Contains[Dictionary.Any]) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalArray(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _has_property_default_value(impl func(ptr unsafe.Pointer, property String.Name) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var property = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0)))))
		defer pointers.End(gd.InternalStringName(property))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, property)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_property_default_value(impl func(ptr unsafe.Pointer, property String.Name) variant.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var property = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0)))))
		defer pointers.End(gd.InternalStringName(property))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, property)
		ptr, ok := pointers.End(gd.InternalVariant(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _update_exports(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (class) _get_script_method_list(impl func(ptr unsafe.Pointer) Array.Contains[Dictionary.Any]) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalArray(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _get_script_property_list(impl func(ptr unsafe.Pointer) Array.Contains[Dictionary.Any]) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalArray(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _get_member_line(impl func(ptr unsafe.Pointer, member String.Name) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var member = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0)))))
		defer pointers.End(gd.InternalStringName(member))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, member)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_constants(impl func(ptr unsafe.Pointer) Dictionary.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalDictionary(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _get_members(impl func(ptr unsafe.Pointer) Array.Contains[String.Name]) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalArray(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _is_placeholder_fallback_enabled(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_rpc_config(impl func(ptr unsafe.Pointer) variant.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalVariant(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (self class) AsScriptExtension() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsScriptExtension() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsScript() Script.Advanced      { return *((*Script.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsScript() Script.Instance   { return *((*Script.Instance)(unsafe.Pointer(&self))) }
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
	case "_editor_can_reload_from_file":
		return reflect.ValueOf(self._editor_can_reload_from_file)
	case "_placeholder_erased":
		return reflect.ValueOf(self._placeholder_erased)
	case "_can_instantiate":
		return reflect.ValueOf(self._can_instantiate)
	case "_get_base_script":
		return reflect.ValueOf(self._get_base_script)
	case "_get_global_name":
		return reflect.ValueOf(self._get_global_name)
	case "_inherits_script":
		return reflect.ValueOf(self._inherits_script)
	case "_get_instance_base_type":
		return reflect.ValueOf(self._get_instance_base_type)
	case "_instance_create":
		return reflect.ValueOf(self._instance_create)
	case "_placeholder_instance_create":
		return reflect.ValueOf(self._placeholder_instance_create)
	case "_instance_has":
		return reflect.ValueOf(self._instance_has)
	case "_has_source_code":
		return reflect.ValueOf(self._has_source_code)
	case "_get_source_code":
		return reflect.ValueOf(self._get_source_code)
	case "_set_source_code":
		return reflect.ValueOf(self._set_source_code)
	case "_reload":
		return reflect.ValueOf(self._reload)
	case "_get_documentation":
		return reflect.ValueOf(self._get_documentation)
	case "_get_class_icon_path":
		return reflect.ValueOf(self._get_class_icon_path)
	case "_has_method":
		return reflect.ValueOf(self._has_method)
	case "_has_static_method":
		return reflect.ValueOf(self._has_static_method)
	case "_get_script_method_argument_count":
		return reflect.ValueOf(self._get_script_method_argument_count)
	case "_get_method_info":
		return reflect.ValueOf(self._get_method_info)
	case "_is_tool":
		return reflect.ValueOf(self._is_tool)
	case "_is_valid":
		return reflect.ValueOf(self._is_valid)
	case "_is_abstract":
		return reflect.ValueOf(self._is_abstract)
	case "_get_language":
		return reflect.ValueOf(self._get_language)
	case "_has_script_signal":
		return reflect.ValueOf(self._has_script_signal)
	case "_get_script_signal_list":
		return reflect.ValueOf(self._get_script_signal_list)
	case "_has_property_default_value":
		return reflect.ValueOf(self._has_property_default_value)
	case "_get_property_default_value":
		return reflect.ValueOf(self._get_property_default_value)
	case "_update_exports":
		return reflect.ValueOf(self._update_exports)
	case "_get_script_method_list":
		return reflect.ValueOf(self._get_script_method_list)
	case "_get_script_property_list":
		return reflect.ValueOf(self._get_script_property_list)
	case "_get_member_line":
		return reflect.ValueOf(self._get_member_line)
	case "_get_constants":
		return reflect.ValueOf(self._get_constants)
	case "_get_members":
		return reflect.ValueOf(self._get_members)
	case "_is_placeholder_fallback_enabled":
		return reflect.ValueOf(self._is_placeholder_fallback_enabled)
	case "_get_rpc_config":
		return reflect.ValueOf(self._get_rpc_config)
	default:
		return gd.VirtualByName(Script.Advanced(self.AsScript()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_editor_can_reload_from_file":
		return reflect.ValueOf(self._editor_can_reload_from_file)
	case "_placeholder_erased":
		return reflect.ValueOf(self._placeholder_erased)
	case "_can_instantiate":
		return reflect.ValueOf(self._can_instantiate)
	case "_get_base_script":
		return reflect.ValueOf(self._get_base_script)
	case "_get_global_name":
		return reflect.ValueOf(self._get_global_name)
	case "_inherits_script":
		return reflect.ValueOf(self._inherits_script)
	case "_get_instance_base_type":
		return reflect.ValueOf(self._get_instance_base_type)
	case "_instance_create":
		return reflect.ValueOf(self._instance_create)
	case "_placeholder_instance_create":
		return reflect.ValueOf(self._placeholder_instance_create)
	case "_instance_has":
		return reflect.ValueOf(self._instance_has)
	case "_has_source_code":
		return reflect.ValueOf(self._has_source_code)
	case "_get_source_code":
		return reflect.ValueOf(self._get_source_code)
	case "_set_source_code":
		return reflect.ValueOf(self._set_source_code)
	case "_reload":
		return reflect.ValueOf(self._reload)
	case "_get_documentation":
		return reflect.ValueOf(self._get_documentation)
	case "_get_class_icon_path":
		return reflect.ValueOf(self._get_class_icon_path)
	case "_has_method":
		return reflect.ValueOf(self._has_method)
	case "_has_static_method":
		return reflect.ValueOf(self._has_static_method)
	case "_get_script_method_argument_count":
		return reflect.ValueOf(self._get_script_method_argument_count)
	case "_get_method_info":
		return reflect.ValueOf(self._get_method_info)
	case "_is_tool":
		return reflect.ValueOf(self._is_tool)
	case "_is_valid":
		return reflect.ValueOf(self._is_valid)
	case "_is_abstract":
		return reflect.ValueOf(self._is_abstract)
	case "_get_language":
		return reflect.ValueOf(self._get_language)
	case "_has_script_signal":
		return reflect.ValueOf(self._has_script_signal)
	case "_get_script_signal_list":
		return reflect.ValueOf(self._get_script_signal_list)
	case "_has_property_default_value":
		return reflect.ValueOf(self._has_property_default_value)
	case "_get_property_default_value":
		return reflect.ValueOf(self._get_property_default_value)
	case "_update_exports":
		return reflect.ValueOf(self._update_exports)
	case "_get_script_method_list":
		return reflect.ValueOf(self._get_script_method_list)
	case "_get_script_property_list":
		return reflect.ValueOf(self._get_script_property_list)
	case "_get_member_line":
		return reflect.ValueOf(self._get_member_line)
	case "_get_constants":
		return reflect.ValueOf(self._get_constants)
	case "_get_members":
		return reflect.ValueOf(self._get_members)
	case "_is_placeholder_fallback_enabled":
		return reflect.ValueOf(self._is_placeholder_fallback_enabled)
	case "_get_rpc_config":
		return reflect.ValueOf(self._get_rpc_config)
	default:
		return gd.VirtualByName(Script.Instance(self.AsScript()), name)
	}
}
func init() {
	gdclass.Register("ScriptExtension", func(ptr gd.Object) any {
		return [1]gdclass.ScriptExtension{*(*gdclass.ScriptExtension)(unsafe.Pointer(&ptr))}
	})
}
