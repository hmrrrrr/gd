// Package ScriptLanguageExtension provides methods for working with ScriptLanguageExtension object instances.
package ScriptLanguageExtension

import "unsafe"
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
import "graphics.gd/classdb/ScriptLanguage"

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

type Instance [1]gdclass.ScriptLanguageExtension

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsScriptLanguageExtension() Instance
}
type Interface interface {
	GetName() string
	Init()
	GetType() string
	GetExtension() string
	Finish()
	GetReservedWords() []string
	IsControlFlowKeyword(keyword string) bool
	GetCommentDelimiters() []string
	GetDocCommentDelimiters() []string
	GetStringDelimiters() []string
	MakeTemplate(template string, class_name string, base_class_name string) [1]gdclass.Script
	GetBuiltInTemplates(obj string) []map[any]any
	IsUsingTemplates() bool
	Validate(script string, path string, validate_functions bool, validate_errors bool, validate_warnings bool, validate_safe_lines bool) map[any]any
	ValidatePath(path string) string
	CreateScript() Object.Instance
	HasNamedClasses() bool
	SupportsBuiltinMode() bool
	SupportsDocumentation() bool
	CanInheritFromFile() bool
	//Returns the line where the function is defined in the code, or [code]-1[/code] if the function is not present.
	FindFunction(function string, code string) int
	MakeFunction(class_name string, function_name string, function_args []string) string
	CanMakeFunction() bool
	OpenInExternalEditor(script [1]gdclass.Script, line int, column int) error
	OverridesExternalEditor() bool
	PreferredFileNameCasing() gdclass.ScriptLanguageScriptNameCasing
	CompleteCode(code string, path string, owner Object.Instance) map[any]any
	LookupCode(code string, symbol string, path string, owner Object.Instance) map[any]any
	AutoIndentCode(code string, from_line int, to_line int) string
	AddGlobalConstant(name string, value any)
	AddNamedGlobalConstant(name string, value any)
	RemoveNamedGlobalConstant(name string)
	ThreadEnter()
	ThreadExit()
	DebugGetError() string
	DebugGetStackLevelCount() int
	DebugGetStackLevelLine(level int) int
	DebugGetStackLevelFunction(level int) string
	//Returns the source associated with a given debug stack position.
	DebugGetStackLevelSource(level int) string
	DebugGetStackLevelLocals(level int, max_subitems int, max_depth int) map[any]any
	DebugGetStackLevelMembers(level int, max_subitems int, max_depth int) map[any]any
	DebugGetStackLevelInstance(level int) unsafe.Pointer
	DebugGetGlobals(max_subitems int, max_depth int) map[any]any
	DebugParseStackLevelExpression(level int, expression string, max_subitems int, max_depth int) string
	DebugGetCurrentStackInfo() []map[any]any
	ReloadAllScripts()
	ReloadToolScript(script [1]gdclass.Script, soft_reload bool)
	GetRecognizedExtensions() []string
	GetPublicFunctions() []map[any]any
	GetPublicConstants() map[any]any
	GetPublicAnnotations() []map[any]any
	ProfilingStart()
	ProfilingStop()
	ProfilingSetSaveNativeCalls(enable bool)
	ProfilingGetAccumulatedData(info_array *ProfilingInfo, info_max int) int
	ProfilingGetFrameData(info_array *ProfilingInfo, info_max int) int
	Frame()
	HandlesGlobalClassType(atype string) bool
	GetGlobalClassName(path string) map[any]any
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) GetName() (_ string)                          { return }
func (self implementation) Init()                                        { return }
func (self implementation) GetType() (_ string)                          { return }
func (self implementation) GetExtension() (_ string)                     { return }
func (self implementation) Finish()                                      { return }
func (self implementation) GetReservedWords() (_ []string)               { return }
func (self implementation) IsControlFlowKeyword(keyword string) (_ bool) { return }
func (self implementation) GetCommentDelimiters() (_ []string)           { return }
func (self implementation) GetDocCommentDelimiters() (_ []string)        { return }
func (self implementation) GetStringDelimiters() (_ []string)            { return }
func (self implementation) MakeTemplate(template string, class_name string, base_class_name string) (_ [1]gdclass.Script) {
	return
}
func (self implementation) GetBuiltInTemplates(obj string) (_ []map[any]any) { return }
func (self implementation) IsUsingTemplates() (_ bool)                       { return }
func (self implementation) Validate(script string, path string, validate_functions bool, validate_errors bool, validate_warnings bool, validate_safe_lines bool) (_ map[any]any) {
	return
}
func (self implementation) ValidatePath(path string) (_ string)               { return }
func (self implementation) CreateScript() (_ Object.Instance)                 { return }
func (self implementation) HasNamedClasses() (_ bool)                         { return }
func (self implementation) SupportsBuiltinMode() (_ bool)                     { return }
func (self implementation) SupportsDocumentation() (_ bool)                   { return }
func (self implementation) CanInheritFromFile() (_ bool)                      { return }
func (self implementation) FindFunction(function string, code string) (_ int) { return }
func (self implementation) MakeFunction(class_name string, function_name string, function_args []string) (_ string) {
	return
}
func (self implementation) CanMakeFunction() (_ bool) { return }
func (self implementation) OpenInExternalEditor(script [1]gdclass.Script, line int, column int) (_ error) {
	return
}
func (self implementation) OverridesExternalEditor() (_ bool) { return }
func (self implementation) PreferredFileNameCasing() (_ gdclass.ScriptLanguageScriptNameCasing) {
	return
}
func (self implementation) CompleteCode(code string, path string, owner Object.Instance) (_ map[any]any) {
	return
}
func (self implementation) LookupCode(code string, symbol string, path string, owner Object.Instance) (_ map[any]any) {
	return
}
func (self implementation) AutoIndentCode(code string, from_line int, to_line int) (_ string) { return }
func (self implementation) AddGlobalConstant(name string, value any)                          { return }
func (self implementation) AddNamedGlobalConstant(name string, value any)                     { return }
func (self implementation) RemoveNamedGlobalConstant(name string)                             { return }
func (self implementation) ThreadEnter()                                                      { return }
func (self implementation) ThreadExit()                                                       { return }
func (self implementation) DebugGetError() (_ string)                                         { return }
func (self implementation) DebugGetStackLevelCount() (_ int)                                  { return }
func (self implementation) DebugGetStackLevelLine(level int) (_ int)                          { return }
func (self implementation) DebugGetStackLevelFunction(level int) (_ string)                   { return }
func (self implementation) DebugGetStackLevelSource(level int) (_ string)                     { return }
func (self implementation) DebugGetStackLevelLocals(level int, max_subitems int, max_depth int) (_ map[any]any) {
	return
}
func (self implementation) DebugGetStackLevelMembers(level int, max_subitems int, max_depth int) (_ map[any]any) {
	return
}
func (self implementation) DebugGetStackLevelInstance(level int) (_ unsafe.Pointer)         { return }
func (self implementation) DebugGetGlobals(max_subitems int, max_depth int) (_ map[any]any) { return }
func (self implementation) DebugParseStackLevelExpression(level int, expression string, max_subitems int, max_depth int) (_ string) {
	return
}
func (self implementation) DebugGetCurrentStackInfo() (_ []map[any]any)                 { return }
func (self implementation) ReloadAllScripts()                                           { return }
func (self implementation) ReloadToolScript(script [1]gdclass.Script, soft_reload bool) { return }
func (self implementation) GetRecognizedExtensions() (_ []string)                       { return }
func (self implementation) GetPublicFunctions() (_ []map[any]any)                       { return }
func (self implementation) GetPublicConstants() (_ map[any]any)                         { return }
func (self implementation) GetPublicAnnotations() (_ []map[any]any)                     { return }
func (self implementation) ProfilingStart()                                             { return }
func (self implementation) ProfilingStop()                                              { return }
func (self implementation) ProfilingSetSaveNativeCalls(enable bool)                     { return }
func (self implementation) ProfilingGetAccumulatedData(info_array *ProfilingInfo, info_max int) (_ int) {
	return
}
func (self implementation) ProfilingGetFrameData(info_array *ProfilingInfo, info_max int) (_ int) {
	return
}
func (self implementation) Frame()                                         { return }
func (self implementation) HandlesGlobalClassType(atype string) (_ bool)   { return }
func (self implementation) GetGlobalClassName(path string) (_ map[any]any) { return }
func (Instance) _get_name(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
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
func (Instance) _init(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}
func (Instance) _get_type(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
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
func (Instance) _get_extension(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
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
func (Instance) _finish(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}
func (Instance) _get_reserved_words(impl func(ptr unsafe.Pointer) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewPackedStringSlice(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _is_control_flow_keyword(impl func(ptr unsafe.Pointer, keyword string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var keyword = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(keyword))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, keyword.String())
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_comment_delimiters(impl func(ptr unsafe.Pointer) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewPackedStringSlice(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _get_doc_comment_delimiters(impl func(ptr unsafe.Pointer) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewPackedStringSlice(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _get_string_delimiters(impl func(ptr unsafe.Pointer) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewPackedStringSlice(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _make_template(impl func(ptr unsafe.Pointer, template string, class_name string, base_class_name string) [1]gdclass.Script) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var template = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(template))
		var class_name = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(class_name))
		var base_class_name = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalString(base_class_name))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, template.String(), class_name.String(), base_class_name.String())
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _get_built_in_templates(impl func(ptr unsafe.Pointer, obj string) []map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var obj = pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		defer pointers.End(obj)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj.String())
		ptr, ok := pointers.End(gd.InternalArray(gd.ArrayFromSlice[Array.Contains[Dictionary.Any]](ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _is_using_templates(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _validate(impl func(ptr unsafe.Pointer, script string, path string, validate_functions bool, validate_errors bool, validate_warnings bool, validate_safe_lines bool) map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var script = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(script))
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(path))
		var validate_functions = gd.UnsafeGet[bool](p_args, 2)

		var validate_errors = gd.UnsafeGet[bool](p_args, 3)

		var validate_warnings = gd.UnsafeGet[bool](p_args, 4)

		var validate_safe_lines = gd.UnsafeGet[bool](p_args, 5)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, script.String(), path.String(), validate_functions, validate_errors, validate_warnings, validate_safe_lines)
		ptr, ok := pointers.End(gd.InternalDictionary(gd.DictionaryFromMap(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _validate_path(impl func(ptr unsafe.Pointer, path string) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(path))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String())
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _create_script(impl func(ptr unsafe.Pointer) Object.Instance) (cb gd.ExtensionClassCallVirtualFunc) {
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
func (Instance) _has_named_classes(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _supports_builtin_mode(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _supports_documentation(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _can_inherit_from_file(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns the line where the function is defined in the code, or [code]-1[/code] if the function is not present.
*/
func (Instance) _find_function(impl func(ptr unsafe.Pointer, function string, code string) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var function = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(function))
		var code = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(code))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, function.String(), code.String())
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _make_function(impl func(ptr unsafe.Pointer, class_name string, function_name string, function_args []string) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var class_name = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(class_name))
		var function_name = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(function_name))
		var function_args = pointers.New[gd.PackedStringArray](gd.UnsafeGet[gd.PackedPointers](p_args, 2))
		defer pointers.End(function_args)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, class_name.String(), function_name.String(), function_args.Strings())
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _can_make_function(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _open_in_external_editor(impl func(ptr unsafe.Pointer, script [1]gdclass.Script, line int, column int) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var script = [1]gdclass.Script{pointers.New[gdclass.Script]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(script[0])
		var line = gd.UnsafeGet[gd.Int](p_args, 1)

		var column = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, script, int(line), int(column))
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _overrides_external_editor(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _preferred_file_name_casing(impl func(ptr unsafe.Pointer) gdclass.ScriptLanguageScriptNameCasing) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _complete_code(impl func(ptr unsafe.Pointer, code string, path string, owner Object.Instance) map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var code = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(code))
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(path))
		var owner = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 2))})}
		defer pointers.End(owner[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, code.String(), path.String(), owner)
		ptr, ok := pointers.End(gd.InternalDictionary(gd.DictionaryFromMap(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _lookup_code(impl func(ptr unsafe.Pointer, code string, symbol string, path string, owner Object.Instance) map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var code = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(code))
		var symbol = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(symbol))
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalString(path))
		var owner = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 3))})}
		defer pointers.End(owner[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, code.String(), symbol.String(), path.String(), owner)
		ptr, ok := pointers.End(gd.InternalDictionary(gd.DictionaryFromMap(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _auto_indent_code(impl func(ptr unsafe.Pointer, code string, from_line int, to_line int) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var code = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(code))
		var from_line = gd.UnsafeGet[gd.Int](p_args, 1)

		var to_line = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, code.String(), int(from_line), int(to_line))
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _add_global_constant(impl func(ptr unsafe.Pointer, name string, value any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var name = pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		defer pointers.End(name)
		var value = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 1))
		defer pointers.End(value)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, name.String(), value.Interface())
	}
}
func (Instance) _add_named_global_constant(impl func(ptr unsafe.Pointer, name string, value any)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var name = pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		defer pointers.End(name)
		var value = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 1))
		defer pointers.End(value)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, name.String(), value.Interface())
	}
}
func (Instance) _remove_named_global_constant(impl func(ptr unsafe.Pointer, name string)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var name = pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		defer pointers.End(name)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, name.String())
	}
}
func (Instance) _thread_enter(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}
func (Instance) _thread_exit(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}
func (Instance) _debug_get_error(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
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
func (Instance) _debug_get_stack_level_count(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _debug_get_stack_level_line(impl func(ptr unsafe.Pointer, level int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var level = gd.UnsafeGet[gd.Int](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(level))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _debug_get_stack_level_function(impl func(ptr unsafe.Pointer, level int) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var level = gd.UnsafeGet[gd.Int](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(level))
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns the source associated with a given debug stack position.
*/
func (Instance) _debug_get_stack_level_source(impl func(ptr unsafe.Pointer, level int) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var level = gd.UnsafeGet[gd.Int](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(level))
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _debug_get_stack_level_locals(impl func(ptr unsafe.Pointer, level int, max_subitems int, max_depth int) map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var level = gd.UnsafeGet[gd.Int](p_args, 0)

		var max_subitems = gd.UnsafeGet[gd.Int](p_args, 1)

		var max_depth = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(level), int(max_subitems), int(max_depth))
		ptr, ok := pointers.End(gd.InternalDictionary(gd.DictionaryFromMap(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _debug_get_stack_level_members(impl func(ptr unsafe.Pointer, level int, max_subitems int, max_depth int) map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var level = gd.UnsafeGet[gd.Int](p_args, 0)

		var max_subitems = gd.UnsafeGet[gd.Int](p_args, 1)

		var max_depth = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(level), int(max_subitems), int(max_depth))
		ptr, ok := pointers.End(gd.InternalDictionary(gd.DictionaryFromMap(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _debug_get_stack_level_instance(impl func(ptr unsafe.Pointer, level int) unsafe.Pointer) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var level = gd.UnsafeGet[gd.Int](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(level))
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _debug_get_globals(impl func(ptr unsafe.Pointer, max_subitems int, max_depth int) map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var max_subitems = gd.UnsafeGet[gd.Int](p_args, 0)

		var max_depth = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(max_subitems), int(max_depth))
		ptr, ok := pointers.End(gd.InternalDictionary(gd.DictionaryFromMap(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _debug_parse_stack_level_expression(impl func(ptr unsafe.Pointer, level int, expression string, max_subitems int, max_depth int) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var level = gd.UnsafeGet[gd.Int](p_args, 0)

		var expression = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(expression))
		var max_subitems = gd.UnsafeGet[gd.Int](p_args, 2)

		var max_depth = gd.UnsafeGet[gd.Int](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(level), expression.String(), int(max_subitems), int(max_depth))
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _debug_get_current_stack_info(impl func(ptr unsafe.Pointer) []map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
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
func (Instance) _reload_all_scripts(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}
func (Instance) _reload_tool_script(impl func(ptr unsafe.Pointer, script [1]gdclass.Script, soft_reload bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var script = [1]gdclass.Script{pointers.New[gdclass.Script]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(script[0])
		var soft_reload = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, script, soft_reload)
	}
}
func (Instance) _get_recognized_extensions(impl func(ptr unsafe.Pointer) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewPackedStringSlice(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _get_public_functions(impl func(ptr unsafe.Pointer) []map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
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
func (Instance) _get_public_constants(impl func(ptr unsafe.Pointer) map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
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
func (Instance) _get_public_annotations(impl func(ptr unsafe.Pointer) []map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
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
func (Instance) _profiling_start(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}
func (Instance) _profiling_stop(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}
func (Instance) _profiling_set_save_native_calls(impl func(ptr unsafe.Pointer, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var enable = gd.UnsafeGet[bool](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, enable)
	}
}
func (Instance) _profiling_get_accumulated_data(impl func(ptr unsafe.Pointer, info_array *ProfilingInfo, info_max int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var info_array = gd.UnsafeGet[*ProfilingInfo](p_args, 0)

		var info_max = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, info_array, int(info_max))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _profiling_get_frame_data(impl func(ptr unsafe.Pointer, info_array *ProfilingInfo, info_max int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var info_array = gd.UnsafeGet[*ProfilingInfo](p_args, 0)

		var info_max = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, info_array, int(info_max))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}
func (Instance) _frame(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}
func (Instance) _handles_global_class_type(impl func(ptr unsafe.Pointer, atype string) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var atype = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(atype))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, atype.String())
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_global_class_name(impl func(ptr unsafe.Pointer, path string) map[any]any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(path))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String())
		ptr, ok := pointers.End(gd.InternalDictionary(gd.DictionaryFromMap(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ScriptLanguageExtension

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ScriptLanguageExtension"))
	casted := Instance{*(*gdclass.ScriptLanguageExtension)(unsafe.Pointer(&object))}
	return casted
}

func (class) _get_name(impl func(ptr unsafe.Pointer) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
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

func (class) _init(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (class) _get_type(impl func(ptr unsafe.Pointer) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
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

func (class) _get_extension(impl func(ptr unsafe.Pointer) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
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

func (class) _finish(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (class) _get_reserved_words(impl func(ptr unsafe.Pointer) gd.PackedStringArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _is_control_flow_keyword(impl func(ptr unsafe.Pointer, keyword String.Readable) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var keyword = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(keyword))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, keyword)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_comment_delimiters(impl func(ptr unsafe.Pointer) gd.PackedStringArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _get_doc_comment_delimiters(impl func(ptr unsafe.Pointer) gd.PackedStringArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _get_string_delimiters(impl func(ptr unsafe.Pointer) gd.PackedStringArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _make_template(impl func(ptr unsafe.Pointer, template String.Readable, class_name String.Readable, base_class_name String.Readable) [1]gdclass.Script) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var template = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(template))
		var class_name = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(class_name))
		var base_class_name = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalString(base_class_name))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, template, class_name, base_class_name)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _get_built_in_templates(impl func(ptr unsafe.Pointer, obj gd.StringName) Array.Contains[Dictionary.Any]) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var obj = pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		defer pointers.End(obj)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj)
		ptr, ok := pointers.End(gd.InternalArray(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _is_using_templates(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _validate(impl func(ptr unsafe.Pointer, script String.Readable, path String.Readable, validate_functions bool, validate_errors bool, validate_warnings bool, validate_safe_lines bool) Dictionary.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var script = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(script))
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(path))
		var validate_functions = gd.UnsafeGet[bool](p_args, 2)

		var validate_errors = gd.UnsafeGet[bool](p_args, 3)

		var validate_warnings = gd.UnsafeGet[bool](p_args, 4)

		var validate_safe_lines = gd.UnsafeGet[bool](p_args, 5)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, script, path, validate_functions, validate_errors, validate_warnings, validate_safe_lines)
		ptr, ok := pointers.End(gd.InternalDictionary(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _validate_path(impl func(ptr unsafe.Pointer, path String.Readable) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(path))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _create_script(impl func(ptr unsafe.Pointer) [1]gd.Object) (cb gd.ExtensionClassCallVirtualFunc) {
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

func (class) _has_named_classes(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _supports_builtin_mode(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _supports_documentation(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _can_inherit_from_file(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns the line where the function is defined in the code, or [code]-1[/code] if the function is not present.
*/
func (class) _find_function(impl func(ptr unsafe.Pointer, function String.Readable, code String.Readable) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var function = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(function))
		var code = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(code))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, function, code)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _make_function(impl func(ptr unsafe.Pointer, class_name String.Readable, function_name String.Readable, function_args gd.PackedStringArray) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var class_name = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(class_name))
		var function_name = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(function_name))
		var function_args = pointers.New[gd.PackedStringArray](gd.UnsafeGet[gd.PackedPointers](p_args, 2))
		defer pointers.End(function_args)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, class_name, function_name, function_args)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _can_make_function(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _open_in_external_editor(impl func(ptr unsafe.Pointer, script [1]gdclass.Script, line gd.Int, column gd.Int) gd.Error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var script = [1]gdclass.Script{pointers.New[gdclass.Script]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(script[0])
		var line = gd.UnsafeGet[gd.Int](p_args, 1)

		var column = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, script, line, column)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _overrides_external_editor(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _preferred_file_name_casing(impl func(ptr unsafe.Pointer) gdclass.ScriptLanguageScriptNameCasing) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _complete_code(impl func(ptr unsafe.Pointer, code String.Readable, path String.Readable, owner [1]gd.Object) Dictionary.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var code = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(code))
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(path))
		var owner = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 2))})}
		defer pointers.End(owner[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, code, path, owner)
		ptr, ok := pointers.End(gd.InternalDictionary(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _lookup_code(impl func(ptr unsafe.Pointer, code String.Readable, symbol String.Readable, path String.Readable, owner [1]gd.Object) Dictionary.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var code = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(code))
		var symbol = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(symbol))
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalString(path))
		var owner = [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 3))})}
		defer pointers.End(owner[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, code, symbol, path, owner)
		ptr, ok := pointers.End(gd.InternalDictionary(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _auto_indent_code(impl func(ptr unsafe.Pointer, code String.Readable, from_line gd.Int, to_line gd.Int) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var code = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(code))
		var from_line = gd.UnsafeGet[gd.Int](p_args, 1)

		var to_line = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, code, from_line, to_line)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _add_global_constant(impl func(ptr unsafe.Pointer, name gd.StringName, value gd.Variant)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var name = pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		defer pointers.End(name)
		var value = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 1))
		defer pointers.End(value)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, name, value)
	}
}

func (class) _add_named_global_constant(impl func(ptr unsafe.Pointer, name gd.StringName, value gd.Variant)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var name = pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		defer pointers.End(name)
		var value = pointers.New[gd.Variant](gd.UnsafeGet[[3]uint64](p_args, 1))
		defer pointers.End(value)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, name, value)
	}
}

func (class) _remove_named_global_constant(impl func(ptr unsafe.Pointer, name gd.StringName)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var name = pointers.New[gd.StringName](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))
		defer pointers.End(name)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, name)
	}
}

func (class) _thread_enter(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (class) _thread_exit(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (class) _debug_get_error(impl func(ptr unsafe.Pointer) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
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

func (class) _debug_get_stack_level_count(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _debug_get_stack_level_line(impl func(ptr unsafe.Pointer, level gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var level = gd.UnsafeGet[gd.Int](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, level)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _debug_get_stack_level_function(impl func(ptr unsafe.Pointer, level gd.Int) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var level = gd.UnsafeGet[gd.Int](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, level)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns the source associated with a given debug stack position.
*/
func (class) _debug_get_stack_level_source(impl func(ptr unsafe.Pointer, level gd.Int) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var level = gd.UnsafeGet[gd.Int](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, level)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _debug_get_stack_level_locals(impl func(ptr unsafe.Pointer, level gd.Int, max_subitems gd.Int, max_depth gd.Int) Dictionary.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var level = gd.UnsafeGet[gd.Int](p_args, 0)

		var max_subitems = gd.UnsafeGet[gd.Int](p_args, 1)

		var max_depth = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, level, max_subitems, max_depth)
		ptr, ok := pointers.End(gd.InternalDictionary(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _debug_get_stack_level_members(impl func(ptr unsafe.Pointer, level gd.Int, max_subitems gd.Int, max_depth gd.Int) Dictionary.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var level = gd.UnsafeGet[gd.Int](p_args, 0)

		var max_subitems = gd.UnsafeGet[gd.Int](p_args, 1)

		var max_depth = gd.UnsafeGet[gd.Int](p_args, 2)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, level, max_subitems, max_depth)
		ptr, ok := pointers.End(gd.InternalDictionary(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _debug_get_stack_level_instance(impl func(ptr unsafe.Pointer, level gd.Int) unsafe.Pointer) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var level = gd.UnsafeGet[gd.Int](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, level)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _debug_get_globals(impl func(ptr unsafe.Pointer, max_subitems gd.Int, max_depth gd.Int) Dictionary.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var max_subitems = gd.UnsafeGet[gd.Int](p_args, 0)

		var max_depth = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, max_subitems, max_depth)
		ptr, ok := pointers.End(gd.InternalDictionary(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _debug_parse_stack_level_expression(impl func(ptr unsafe.Pointer, level gd.Int, expression String.Readable, max_subitems gd.Int, max_depth gd.Int) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var level = gd.UnsafeGet[gd.Int](p_args, 0)

		var expression = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(expression))
		var max_subitems = gd.UnsafeGet[gd.Int](p_args, 2)

		var max_depth = gd.UnsafeGet[gd.Int](p_args, 3)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, level, expression, max_subitems, max_depth)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _debug_get_current_stack_info(impl func(ptr unsafe.Pointer) Array.Contains[Dictionary.Any]) (cb gd.ExtensionClassCallVirtualFunc) {
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

func (class) _reload_all_scripts(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (class) _reload_tool_script(impl func(ptr unsafe.Pointer, script [1]gdclass.Script, soft_reload bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var script = [1]gdclass.Script{pointers.New[gdclass.Script]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(script[0])
		var soft_reload = gd.UnsafeGet[bool](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, script, soft_reload)
	}
}

func (class) _get_recognized_extensions(impl func(ptr unsafe.Pointer) gd.PackedStringArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _get_public_functions(impl func(ptr unsafe.Pointer) Array.Contains[Dictionary.Any]) (cb gd.ExtensionClassCallVirtualFunc) {
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

func (class) _get_public_constants(impl func(ptr unsafe.Pointer) Dictionary.Any) (cb gd.ExtensionClassCallVirtualFunc) {
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

func (class) _get_public_annotations(impl func(ptr unsafe.Pointer) Array.Contains[Dictionary.Any]) (cb gd.ExtensionClassCallVirtualFunc) {
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

func (class) _profiling_start(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (class) _profiling_stop(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (class) _profiling_set_save_native_calls(impl func(ptr unsafe.Pointer, enable bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var enable = gd.UnsafeGet[bool](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, enable)
	}
}

func (class) _profiling_get_accumulated_data(impl func(ptr unsafe.Pointer, info_array *ProfilingInfo, info_max gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var info_array = gd.UnsafeGet[*ProfilingInfo](p_args, 0)

		var info_max = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, info_array, info_max)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _profiling_get_frame_data(impl func(ptr unsafe.Pointer, info_array *ProfilingInfo, info_max gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var info_array = gd.UnsafeGet[*ProfilingInfo](p_args, 0)

		var info_max = gd.UnsafeGet[gd.Int](p_args, 1)

		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, info_array, info_max)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _frame(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (class) _handles_global_class_type(impl func(ptr unsafe.Pointer, atype String.Readable) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var atype = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(atype))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, atype)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_global_class_name(impl func(ptr unsafe.Pointer, path String.Readable) Dictionary.Any) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(path))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path)
		ptr, ok := pointers.End(gd.InternalDictionary(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (self class) AsScriptLanguageExtension() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsScriptLanguageExtension() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsScriptLanguage() ScriptLanguage.Advanced {
	return *((*ScriptLanguage.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsScriptLanguage() ScriptLanguage.Instance {
	return *((*ScriptLanguage.Instance)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_name":
		return reflect.ValueOf(self._get_name)
	case "_init":
		return reflect.ValueOf(self._init)
	case "_get_type":
		return reflect.ValueOf(self._get_type)
	case "_get_extension":
		return reflect.ValueOf(self._get_extension)
	case "_finish":
		return reflect.ValueOf(self._finish)
	case "_get_reserved_words":
		return reflect.ValueOf(self._get_reserved_words)
	case "_is_control_flow_keyword":
		return reflect.ValueOf(self._is_control_flow_keyword)
	case "_get_comment_delimiters":
		return reflect.ValueOf(self._get_comment_delimiters)
	case "_get_doc_comment_delimiters":
		return reflect.ValueOf(self._get_doc_comment_delimiters)
	case "_get_string_delimiters":
		return reflect.ValueOf(self._get_string_delimiters)
	case "_make_template":
		return reflect.ValueOf(self._make_template)
	case "_get_built_in_templates":
		return reflect.ValueOf(self._get_built_in_templates)
	case "_is_using_templates":
		return reflect.ValueOf(self._is_using_templates)
	case "_validate":
		return reflect.ValueOf(self._validate)
	case "_validate_path":
		return reflect.ValueOf(self._validate_path)
	case "_create_script":
		return reflect.ValueOf(self._create_script)
	case "_has_named_classes":
		return reflect.ValueOf(self._has_named_classes)
	case "_supports_builtin_mode":
		return reflect.ValueOf(self._supports_builtin_mode)
	case "_supports_documentation":
		return reflect.ValueOf(self._supports_documentation)
	case "_can_inherit_from_file":
		return reflect.ValueOf(self._can_inherit_from_file)
	case "_find_function":
		return reflect.ValueOf(self._find_function)
	case "_make_function":
		return reflect.ValueOf(self._make_function)
	case "_can_make_function":
		return reflect.ValueOf(self._can_make_function)
	case "_open_in_external_editor":
		return reflect.ValueOf(self._open_in_external_editor)
	case "_overrides_external_editor":
		return reflect.ValueOf(self._overrides_external_editor)
	case "_preferred_file_name_casing":
		return reflect.ValueOf(self._preferred_file_name_casing)
	case "_complete_code":
		return reflect.ValueOf(self._complete_code)
	case "_lookup_code":
		return reflect.ValueOf(self._lookup_code)
	case "_auto_indent_code":
		return reflect.ValueOf(self._auto_indent_code)
	case "_add_global_constant":
		return reflect.ValueOf(self._add_global_constant)
	case "_add_named_global_constant":
		return reflect.ValueOf(self._add_named_global_constant)
	case "_remove_named_global_constant":
		return reflect.ValueOf(self._remove_named_global_constant)
	case "_thread_enter":
		return reflect.ValueOf(self._thread_enter)
	case "_thread_exit":
		return reflect.ValueOf(self._thread_exit)
	case "_debug_get_error":
		return reflect.ValueOf(self._debug_get_error)
	case "_debug_get_stack_level_count":
		return reflect.ValueOf(self._debug_get_stack_level_count)
	case "_debug_get_stack_level_line":
		return reflect.ValueOf(self._debug_get_stack_level_line)
	case "_debug_get_stack_level_function":
		return reflect.ValueOf(self._debug_get_stack_level_function)
	case "_debug_get_stack_level_source":
		return reflect.ValueOf(self._debug_get_stack_level_source)
	case "_debug_get_stack_level_locals":
		return reflect.ValueOf(self._debug_get_stack_level_locals)
	case "_debug_get_stack_level_members":
		return reflect.ValueOf(self._debug_get_stack_level_members)
	case "_debug_get_stack_level_instance":
		return reflect.ValueOf(self._debug_get_stack_level_instance)
	case "_debug_get_globals":
		return reflect.ValueOf(self._debug_get_globals)
	case "_debug_parse_stack_level_expression":
		return reflect.ValueOf(self._debug_parse_stack_level_expression)
	case "_debug_get_current_stack_info":
		return reflect.ValueOf(self._debug_get_current_stack_info)
	case "_reload_all_scripts":
		return reflect.ValueOf(self._reload_all_scripts)
	case "_reload_tool_script":
		return reflect.ValueOf(self._reload_tool_script)
	case "_get_recognized_extensions":
		return reflect.ValueOf(self._get_recognized_extensions)
	case "_get_public_functions":
		return reflect.ValueOf(self._get_public_functions)
	case "_get_public_constants":
		return reflect.ValueOf(self._get_public_constants)
	case "_get_public_annotations":
		return reflect.ValueOf(self._get_public_annotations)
	case "_profiling_start":
		return reflect.ValueOf(self._profiling_start)
	case "_profiling_stop":
		return reflect.ValueOf(self._profiling_stop)
	case "_profiling_set_save_native_calls":
		return reflect.ValueOf(self._profiling_set_save_native_calls)
	case "_profiling_get_accumulated_data":
		return reflect.ValueOf(self._profiling_get_accumulated_data)
	case "_profiling_get_frame_data":
		return reflect.ValueOf(self._profiling_get_frame_data)
	case "_frame":
		return reflect.ValueOf(self._frame)
	case "_handles_global_class_type":
		return reflect.ValueOf(self._handles_global_class_type)
	case "_get_global_class_name":
		return reflect.ValueOf(self._get_global_class_name)
	default:
		return gd.VirtualByName(ScriptLanguage.Advanced(self.AsScriptLanguage()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_name":
		return reflect.ValueOf(self._get_name)
	case "_init":
		return reflect.ValueOf(self._init)
	case "_get_type":
		return reflect.ValueOf(self._get_type)
	case "_get_extension":
		return reflect.ValueOf(self._get_extension)
	case "_finish":
		return reflect.ValueOf(self._finish)
	case "_get_reserved_words":
		return reflect.ValueOf(self._get_reserved_words)
	case "_is_control_flow_keyword":
		return reflect.ValueOf(self._is_control_flow_keyword)
	case "_get_comment_delimiters":
		return reflect.ValueOf(self._get_comment_delimiters)
	case "_get_doc_comment_delimiters":
		return reflect.ValueOf(self._get_doc_comment_delimiters)
	case "_get_string_delimiters":
		return reflect.ValueOf(self._get_string_delimiters)
	case "_make_template":
		return reflect.ValueOf(self._make_template)
	case "_get_built_in_templates":
		return reflect.ValueOf(self._get_built_in_templates)
	case "_is_using_templates":
		return reflect.ValueOf(self._is_using_templates)
	case "_validate":
		return reflect.ValueOf(self._validate)
	case "_validate_path":
		return reflect.ValueOf(self._validate_path)
	case "_create_script":
		return reflect.ValueOf(self._create_script)
	case "_has_named_classes":
		return reflect.ValueOf(self._has_named_classes)
	case "_supports_builtin_mode":
		return reflect.ValueOf(self._supports_builtin_mode)
	case "_supports_documentation":
		return reflect.ValueOf(self._supports_documentation)
	case "_can_inherit_from_file":
		return reflect.ValueOf(self._can_inherit_from_file)
	case "_find_function":
		return reflect.ValueOf(self._find_function)
	case "_make_function":
		return reflect.ValueOf(self._make_function)
	case "_can_make_function":
		return reflect.ValueOf(self._can_make_function)
	case "_open_in_external_editor":
		return reflect.ValueOf(self._open_in_external_editor)
	case "_overrides_external_editor":
		return reflect.ValueOf(self._overrides_external_editor)
	case "_preferred_file_name_casing":
		return reflect.ValueOf(self._preferred_file_name_casing)
	case "_complete_code":
		return reflect.ValueOf(self._complete_code)
	case "_lookup_code":
		return reflect.ValueOf(self._lookup_code)
	case "_auto_indent_code":
		return reflect.ValueOf(self._auto_indent_code)
	case "_add_global_constant":
		return reflect.ValueOf(self._add_global_constant)
	case "_add_named_global_constant":
		return reflect.ValueOf(self._add_named_global_constant)
	case "_remove_named_global_constant":
		return reflect.ValueOf(self._remove_named_global_constant)
	case "_thread_enter":
		return reflect.ValueOf(self._thread_enter)
	case "_thread_exit":
		return reflect.ValueOf(self._thread_exit)
	case "_debug_get_error":
		return reflect.ValueOf(self._debug_get_error)
	case "_debug_get_stack_level_count":
		return reflect.ValueOf(self._debug_get_stack_level_count)
	case "_debug_get_stack_level_line":
		return reflect.ValueOf(self._debug_get_stack_level_line)
	case "_debug_get_stack_level_function":
		return reflect.ValueOf(self._debug_get_stack_level_function)
	case "_debug_get_stack_level_source":
		return reflect.ValueOf(self._debug_get_stack_level_source)
	case "_debug_get_stack_level_locals":
		return reflect.ValueOf(self._debug_get_stack_level_locals)
	case "_debug_get_stack_level_members":
		return reflect.ValueOf(self._debug_get_stack_level_members)
	case "_debug_get_stack_level_instance":
		return reflect.ValueOf(self._debug_get_stack_level_instance)
	case "_debug_get_globals":
		return reflect.ValueOf(self._debug_get_globals)
	case "_debug_parse_stack_level_expression":
		return reflect.ValueOf(self._debug_parse_stack_level_expression)
	case "_debug_get_current_stack_info":
		return reflect.ValueOf(self._debug_get_current_stack_info)
	case "_reload_all_scripts":
		return reflect.ValueOf(self._reload_all_scripts)
	case "_reload_tool_script":
		return reflect.ValueOf(self._reload_tool_script)
	case "_get_recognized_extensions":
		return reflect.ValueOf(self._get_recognized_extensions)
	case "_get_public_functions":
		return reflect.ValueOf(self._get_public_functions)
	case "_get_public_constants":
		return reflect.ValueOf(self._get_public_constants)
	case "_get_public_annotations":
		return reflect.ValueOf(self._get_public_annotations)
	case "_profiling_start":
		return reflect.ValueOf(self._profiling_start)
	case "_profiling_stop":
		return reflect.ValueOf(self._profiling_stop)
	case "_profiling_set_save_native_calls":
		return reflect.ValueOf(self._profiling_set_save_native_calls)
	case "_profiling_get_accumulated_data":
		return reflect.ValueOf(self._profiling_get_accumulated_data)
	case "_profiling_get_frame_data":
		return reflect.ValueOf(self._profiling_get_frame_data)
	case "_frame":
		return reflect.ValueOf(self._frame)
	case "_handles_global_class_type":
		return reflect.ValueOf(self._handles_global_class_type)
	case "_get_global_class_name":
		return reflect.ValueOf(self._get_global_class_name)
	default:
		return gd.VirtualByName(ScriptLanguage.Instance(self.AsScriptLanguage()), name)
	}
}
func init() {
	gdclass.Register("ScriptLanguageExtension", func(ptr gd.Object) any {
		return [1]gdclass.ScriptLanguageExtension{*(*gdclass.ScriptLanguageExtension)(unsafe.Pointer(&ptr))}
	})
}

type LookupResultType = gdclass.ScriptLanguageExtensionLookupResultType //gd:ScriptLanguageExtension.LookupResultType

const (
	LookupResultScriptLocation      LookupResultType = 0
	LookupResultClass               LookupResultType = 1
	LookupResultClassConstant       LookupResultType = 2
	LookupResultClassProperty       LookupResultType = 3
	LookupResultClassMethod         LookupResultType = 4
	LookupResultClassSignal         LookupResultType = 5
	LookupResultClassEnum           LookupResultType = 6
	LookupResultClassTbdGlobalscope LookupResultType = 7
	LookupResultClassAnnotation     LookupResultType = 8
	LookupResultMax                 LookupResultType = 9
)

type CodeCompletionLocation = gdclass.ScriptLanguageExtensionCodeCompletionLocation //gd:ScriptLanguageExtension.CodeCompletionLocation

const (
	/*The option is local to the location of the code completion query - e.g. a local variable. Subsequent value of location represent options from the outer class, the exact value represent how far they are (in terms of inner classes).*/
	LocationLocal CodeCompletionLocation = 0
	/*The option is from the containing class or a parent class, relative to the location of the code completion query. Perform a bitwise OR with the class depth (e.g. [code]0[/code] for the local class, [code]1[/code] for the parent, [code]2[/code] for the grandparent, etc.) to store the depth of an option in the class or a parent class.*/
	LocationParentMask CodeCompletionLocation = 256
	/*The option is from user code which is not local and not in a derived class (e.g. Autoload Singletons).*/
	LocationOtherUserCode CodeCompletionLocation = 512
	/*The option is from other engine code, not covered by the other enum constants - e.g. built-in classes.*/
	LocationOther CodeCompletionLocation = 1024
)

type CodeCompletionKind = gdclass.ScriptLanguageExtensionCodeCompletionKind //gd:ScriptLanguageExtension.CodeCompletionKind

const (
	CodeCompletionKindClass     CodeCompletionKind = 0
	CodeCompletionKindFunction  CodeCompletionKind = 1
	CodeCompletionKindSignal    CodeCompletionKind = 2
	CodeCompletionKindVariable  CodeCompletionKind = 3
	CodeCompletionKindMember    CodeCompletionKind = 4
	CodeCompletionKindEnum      CodeCompletionKind = 5
	CodeCompletionKindConstant  CodeCompletionKind = 6
	CodeCompletionKindNodePath  CodeCompletionKind = 7
	CodeCompletionKindFilePath  CodeCompletionKind = 8
	CodeCompletionKindPlainText CodeCompletionKind = 9
	CodeCompletionKindMax       CodeCompletionKind = 10
)

type Error = gd.Error //gd:Error

const (
	/*Methods that return [enum Error] return [constant OK] when no error occurred.
	  Since [constant OK] has value 0, and all other error constants are positive integers, it can also be used in boolean checks.
	  [b]Example:[/b]
	  [codeblock]
	  var error = method_that_returns_error()
	  if error != OK:
	      printerr("Failure!")

	  # Or, alternatively:
	  if error:
	      printerr("Still failing!")
	  [/codeblock]
	  [b]Note:[/b] Many functions do not return an error code, but will print error messages to standard output.*/
	Ok Error = 0
	/*Generic error.*/
	Failed Error = 1
	/*Unavailable error.*/
	ErrUnavailable Error = 2
	/*Unconfigured error.*/
	ErrUnconfigured Error = 3
	/*Unauthorized error.*/
	ErrUnauthorized Error = 4
	/*Parameter range error.*/
	ErrParameterRangeError Error = 5
	/*Out of memory (OOM) error.*/
	ErrOutOfMemory Error = 6
	/*File: Not found error.*/
	ErrFileNotFound Error = 7
	/*File: Bad drive error.*/
	ErrFileBadDrive Error = 8
	/*File: Bad path error.*/
	ErrFileBadPath Error = 9
	/*File: No permission error.*/
	ErrFileNoPermission Error = 10
	/*File: Already in use error.*/
	ErrFileAlreadyInUse Error = 11
	/*File: Can't open error.*/
	ErrFileCantOpen Error = 12
	/*File: Can't write error.*/
	ErrFileCantWrite Error = 13
	/*File: Can't read error.*/
	ErrFileCantRead Error = 14
	/*File: Unrecognized error.*/
	ErrFileUnrecognized Error = 15
	/*File: Corrupt error.*/
	ErrFileCorrupt Error = 16
	/*File: Missing dependencies error.*/
	ErrFileMissingDependencies Error = 17
	/*File: End of file (EOF) error.*/
	ErrFileEof Error = 18
	/*Can't open error.*/
	ErrCantOpen Error = 19
	/*Can't create error.*/
	ErrCantCreate Error = 20
	/*Query failed error.*/
	ErrQueryFailed Error = 21
	/*Already in use error.*/
	ErrAlreadyInUse Error = 22
	/*Locked error.*/
	ErrLocked Error = 23
	/*Timeout error.*/
	ErrTimeout Error = 24
	/*Can't connect error.*/
	ErrCantConnect Error = 25
	/*Can't resolve error.*/
	ErrCantResolve Error = 26
	/*Connection error.*/
	ErrConnectionError Error = 27
	/*Can't acquire resource error.*/
	ErrCantAcquireResource Error = 28
	/*Can't fork process error.*/
	ErrCantFork Error = 29
	/*Invalid data error.*/
	ErrInvalidData Error = 30
	/*Invalid parameter error.*/
	ErrInvalidParameter Error = 31
	/*Already exists error.*/
	ErrAlreadyExists Error = 32
	/*Does not exist error.*/
	ErrDoesNotExist Error = 33
	/*Database: Read error.*/
	ErrDatabaseCantRead Error = 34
	/*Database: Write error.*/
	ErrDatabaseCantWrite Error = 35
	/*Compilation failed error.*/
	ErrCompilationFailed Error = 36
	/*Method not found error.*/
	ErrMethodNotFound Error = 37
	/*Linking failed error.*/
	ErrLinkFailed Error = 38
	/*Script failed error.*/
	ErrScriptFailed Error = 39
	/*Cycling link (import cycle) error.*/
	ErrCyclicLink Error = 40
	/*Invalid declaration error.*/
	ErrInvalidDeclaration Error = 41
	/*Duplicate symbol error.*/
	ErrDuplicateSymbol Error = 42
	/*Parse error.*/
	ErrParseError Error = 43
	/*Busy error.*/
	ErrBusy Error = 44
	/*Skip error.*/
	ErrSkip Error = 45
	/*Help error. Used internally when passing [code]--version[/code] or [code]--help[/code] as executable options.*/
	ErrHelp Error = 46
	/*Bug error, caused by an implementation issue in the method.
	  [b]Note:[/b] If a built-in method returns this code, please open an issue on [url=https://github.com/godotengine/godot/issues]the GitHub Issue Tracker[/url].*/
	ErrBug Error = 47
	/*Printer on fire error (This is an easter egg, no built-in methods return this error code).*/
	ErrPrinterOnFire Error = 48
)
