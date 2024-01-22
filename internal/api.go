//go:build !generate

package gd

import (
	"unsafe"

	"runtime.link/api"
	"runtime.link/api/call"
	"runtime.link/mmm"
)

// API specification for Godot's GDExtension.
type API struct {
	api.Specification

	GetGodotVersion     func() Version
	GetNativeStructSize func(StringName) uintptr

	Memory struct {
		Allocate   func(uintptr) unsafe.Pointer
		Reallocate func(unsafe.Pointer, uintptr) unsafe.Pointer
		Free       func(unsafe.Pointer)
	}

	PrintError              func(code, function, file string, line int32, notifyEditor bool)
	PrintErrorMessage       func(code, message, function, file string, line int32, notifyEditor bool)
	PrintWarning            func(code, function, file string, line int32, notifyEditor bool)
	PrintWarningMessage     func(code, message, function, file string, line int32, notifyEditor bool)
	PrintScriptError        func(code, function, file string, line int32, notifyEditor bool)
	PrintScriptErrorMessage func(code, message, function, file string, line int32, notifyEditor bool)

	Variants struct {
		NewCopy                   func(ctx Context, src Variant) Variant
		NewNil                    func(ctx Context) Variant
		Destroy                   func(self Variant)
		Call                      func(ctx Context, self Variant, method StringName, args ...Variant) (Variant, error)
		CallStatic                func(ctx Context, vtype VariantType, method StringName, args ...Variant) (Variant, error)
		Evaluate                  func(ctx Context, operator Operator, a, b Variant) (ret Variant, ok bool)
		Set                       func(self, key, val Variant) bool
		SetNamed                  func(self Variant, key StringName, val Variant) bool
		SetKeyed                  func(self, key, val Variant) bool
		SetIndexed                func(self Variant, index int64, val Variant) (ok, oob bool)
		Get                       func(ctx Context, self, key Variant) (Variant, bool)
		GetNamed                  func(ctx Context, self Variant, key StringName) (Variant, bool)
		GetKeyed                  func(ctx Context, self, key Variant) (Variant, bool)
		GetIndexed                func(ctx Context, self Variant, index int64) (val Variant, ok, oob bool)
		IteratorInitialize        func(ctx Context, self Variant) (Variant, bool)
		IteratorNext              func(self Variant, iterator Variant) bool
		IteratorGet               func(ctx Context, self Variant, iterator Variant) (Variant, bool)
		Hash                      func(self Variant) int64
		RecursiveHash             func(self Variant, count Int) int64
		HashCompare               func(self, variant Variant) bool
		Booleanize                func(self Variant) bool
		Duplicate                 func(ctx Context, self Variant, deep bool) Variant
		Stringify                 func(ctx Context, self Variant) String
		GetType                   func(self Variant) VariantType
		HasMethod                 func(self Variant, method StringName) bool
		HasMember                 func(self Variant, member StringName) bool
		HasKey                    func(self Variant, key Variant) (hasKey, valid bool)
		GetTypeName               func(ctx Context, self VariantType) String
		CanConvert                func(self Variant, to VariantType) bool
		CanConvertStrict          func(self Variant, to VariantType) bool
		FromTypeConstructor       func(VariantType) func(ret call.Ptr[[3]uintptr], arg uintptr)
		ToTypeConstructor         func(VariantType) func(ret uintptr, arg call.Ptr[[3]uintptr])
		PointerOperatorEvaluator  func(op Operator, a, b VariantType) func(a, b, ret uintptr)
		GetPointerBuiltinMethod   func(VariantType, StringName, Int) func(base uintptr, args call.Args, ret uintptr, c int32)
		GetPointerConstructor     func(vtype VariantType, index int32) func(base uintptr, args call.Args)
		GetPointerDestructor      func(VariantType) func(base uintptr)
		Construct                 func(ctx Context, t VariantType, args ...Variant) (Variant, error)
		GetPointerSetter          func(VariantType, StringName) func(base, arg uintptr)
		GetPointerGetter          func(VariantType, StringName) func(base, ret uintptr)
		GetPointerIndexedSetter   func(VariantType) func(base uintptr, index Int, arg uintptr)
		GetPointerIndexedGetter   func(VariantType) func(base uintptr, index Int, ret uintptr)
		GetPointerKeyedSetter     func(VariantType) func(base uintptr, key uintptr, arg uintptr)
		GetPointerKeyedGetter     func(VariantType) func(base uintptr, key uintptr, ret uintptr)
		GetPointerKeyedChecker    func(VariantType) func(base uintptr, key uintptr) uint32
		GetConstantValue          func(ctx Context, t VariantType, name StringName) Variant
		GetPointerUtilityFunction func(name StringName, hash Int) func(ret uintptr, args call.Args, c int32)
	}
	Strings struct {
		New        func(Context, string) String
		Get        func(String) string
		SetIndex   func(String, Int, rune)
		Index      func(String, Int) rune
		Append     func(Context, String, String) String
		AppendRune func(String, rune)
		Resize     func(String, Int)
	}
	StringNames struct {
		New func(Context, string) StringName
	}
	XMLParser struct {
		OpenBuffer func(XMLParser, []byte) error
	}
	FileAccess struct {
		StoreBuffer func(FileAccess, []byte)
		GetBuffer   func(FileAccess, []byte) int
	}
	PackedByteArray    PackedFunctionsFor[PackedByteArray, byte]
	PackedColorArray   PackedFunctionsFor[PackedColorArray, Color]
	PackedFloat32Array PackedFunctionsFor[PackedFloat32Array, float32]
	PackedFloat64Array PackedFunctionsFor[PackedFloat64Array, float64]
	PackedInt32Array   PackedFunctionsFor[PackedInt32Array, int32]
	PackedInt64Array   PackedFunctionsFor[PackedInt64Array, int64]
	PackedStringArray  struct {
		Index       func(Context, PackedStringArray, Int) String
		SetIndex    func(PackedStringArray, Int, String)
		CopyAsSlice func(Context, PackedStringArray) []String
	}
	PackedVector2Array PackedFunctionsFor[PackedVector2Array, Vector2]
	PackedVector3Array PackedFunctionsFor[PackedVector3Array, Vector3]
	Array              struct {
		Index    func(Context, Array, Int) Variant
		Set      func(self, from Array)
		SetIndex func(Array, Int, Variant)
		SetTyped func(self Array, t VariantType, className StringName, script Script)
	}
	Dictionary struct {
		Index    func(ctx Context, dict Dictionary, key Variant) Variant
		SetIndex func(dict Dictionary, key, val Variant)
	}
	Object struct {
		MethodBindCall        func(ctx Context, method MethodBind, obj Object, arg ...Variant) (Variant, error)
		MethodBindPointerCall func(method MethodBind, obj Object, arg call.Args, ret uintptr)
		Destroy               func(Object)
		GetSingleton          func(ctx Context, name StringName) Object
		GetInstanceBinding    func(Object, ExtensionToken, InstanceBindingType) any
		SetInstanceBinding    func(Object, ExtensionToken, any, InstanceBindingType)
		FreeInstanceBinding   func(Object, ExtensionToken)
		SetInstance           func(Object, StringName, ObjectInterface)
		GetClassName          func(Context, Object, ExtensionToken) String
		CastTo                func(Context, Object, ClassTag) Object
		GetInstanceID         func(Object) ObjectID
	}
	RefCounted struct {
		GetObject func(Context, RefCounted) Object
		SetObject func(RefCounted, Object)
	}
	Callables struct {
		Create func(ctx Context, fn func(...Variant) (Variant, error)) Callable
		Get    func(Callable) (func(...Variant) (Variant, error), bool)
	}
	ClassDB struct {
		ConstructObject func(Context, StringName) Object
		GetClassTag     func(StringName) ClassTag
		GetMethodBind   func(class, method StringName, hash Int) MethodBind

		RegisterClass                 func(library ExtensionToken, name, extends StringName, info ClassInterface)
		RegisterClassMethod           func(ctx Context, library ExtensionToken, class StringName, info Method)
		RegisterClassIntegerConstant  func(library ExtensionToken, class, enum, name StringName, value int64, bitfield bool)
		RegisterClassProperty         func(library ExtensionToken, class StringName, info PropertyInfo, getter, setter StringName)
		RegisterClassPropertyIndexed  func(library ExtensionToken, class StringName, info PropertyInfo, getter, setter StringName, index int64)
		RegisterClassPropertyGroup    func(library ExtensionToken, class StringName, group, prefix String)
		RegisterClassPropertySubGroup func(library ExtensionToken, class StringName, subGroup, prefix String)
		RegisterClassSignal           func(library ExtensionToken, class, signal StringName, args []PropertyInfo)
		UnregisterClass               func(library ExtensionToken, class StringName)
	}
	EditorPlugins struct {
		Add    func(plugin StringName)
		Remove func(plugin StringName)
	}

	GetLibraryPath func(Context, ExtensionToken) String

	ExtensionToken
	cache
}

type Packed interface {
	PackedByteArray | PackedInt32Array | PackedInt64Array | PackedFloat32Array |
		PackedFloat64Array | PackedStringArray | PackedVector2Array | PackedVector3Array |
		PackedColorArray

	mmm.ManagedPointer[[2]uintptr]
	Size() Int
}

type PackedFunctionsFor[T Packed, V any] struct {
	Index       func(T, Int) V
	SetIndex    func(T, Int, V)
	CopyAsSlice func(T) []V
}

type (
	StringPtr     *uintptr
	StringNamePtr *uintptr
	VariantPtr    *[3]uintptr
)

type CallErrorType int32

const (
	OK CallErrorType = iota
	ErrInvalidMethod
	ErrInvalidArgument
	ErrTooManyArguments
	ErrTooFewArguments
	ErrInstanceIsNil
	ErrMethodNotConst
)

type CallError struct {
	ErrorType CallErrorType
	Argument  int32
	Expected  int32
}

type InstanceBindingType unsafe.Pointer

func (err CallError) Error() string {
	switch err.ErrorType {
	case ErrInvalidMethod:
		return "invalid method"
	case ErrInvalidArgument:
		return "invalid argument"
	case ErrTooManyArguments:
		return "too many arguments"
	case ErrTooFewArguments:
		return "too few arguments"
	case ErrInstanceIsNil:
		return "instance is nil"
	case ErrMethodNotConst:
		return "method not const"
	default:
		return "unknown error"
	}
}

type Version struct {
	Major uint32
	Minor uint32
	Patch uint32
	Value string
}

func (v Version) String() string {
	return v.Value
}

type MethodBind uintptr

type ClassTag uintptr
type ExtensionToken uintptr

type InstanceID uint64

type Operator int32

const (
	Equal Operator = iota
	NotEqual
	Less
	LessEqual
	Greater
	GreaterEqual
	Add
	Subtract
	Multiply
	Divide
	Negate
	Positive
	Module
	Power
	ShiftLeft
	ShiftRight
	BitAnd
	BitOr
	BitXor
	BitNegate
	LogicalAnd
	LogicalOr
	LogicalXor
	LogicalNegate
	In
)

type ExtensionInitialization[T any] struct {
	MinimumInitializationLevel GDExtensionInitializationLevel
	Userdata                   T
	Initialize                 call.Back[func(userdata T, level GDExtensionInitializationLevel)]
	Deinitialize               call.Back[func(userdata T, level GDExtensionInitializationLevel)]
}

type InstanceBindingCallbacks struct {
	Create    call.Back[func(token, instance unsafe.Pointer) unsafe.Pointer]
	Free      call.Back[func(token, instance, binding unsafe.Pointer)]
	Reference call.Back[func(token, binding unsafe.Pointer, reference bool) bool]
}

type PropertyInfo struct {
	Type       VariantType
	Name       StringName
	ClassName  StringName
	Hint       PropertyHint
	HintString String
	Usage      PropertyUsageFlags
}

type MethodInfo struct {
	Name             StringName
	ReturnValue      PropertyInfo
	Flags            MethodFlags
	ID               int32
	Arguments        []PropertyInfo
	DefaultArguments []Variant
}

type ClassInterface interface {
	IsVirtual() bool
	IsAbstract() bool
	IsExposed() bool

	CreateInstance() Object
	GetVirtual(StringName) any
}

type ObjectInterface interface {
	Set(StringName, Variant) bool
	Get(StringName) (Variant, bool)
	GetPropertyList(Context) []PropertyInfo
	PropertyCanRevert(StringName) bool
	PropertyGetRevert(StringName) Variant
	ValidateProperty(StringName, PropertyInfo) bool
	Notification(int32, bool)
	ToString() (String, bool)
	Reference()
	Unreference()
	CallVirtual(StringName, any, UnsafeArgs, UnsafeBack)
	GetRID() RID
	Free()
}

type ExtensionClassCallVirtualFunc func(any, UnsafeArgs, UnsafeBack)

type ClassMethodArgumentMetadata uint32

const (
	ArgumentMetadataNone ClassMethodArgumentMetadata = iota
	ArgumentMetadataIntIsInt8
	ArgumentMetadataIntIsInt16
	ArgumentMetadataIntIsInt32
	ArgumentMetadataIntIsInt64
	ArgumentMetadataIntIsUint8
	ArgumentMetadataIntIsUint16
	ArgumentMetadataIntIsUint32
	ArgumentMetadataIntIsUint64
	ArgumentMetadataRealIsFloat32
	ArgumentMetadataRealIsFloat64
)

type Method struct {
	Name                StringName
	Call                func(any, ...Variant) (Variant, error)
	PointerCall         func(any, UnsafeArgs, UnsafeBack)
	MethodFlags         MethodFlags
	ReturnValueInfo     *PropertyInfo
	ReturnValueMetadata ClassMethodArgumentMetadata

	Arguments         []PropertyInfo
	ArgumentsMetadata []ClassMethodArgumentMetadata

	DefaultArguments []Variant
}

type ScriptInstance interface {
	Set(name StringName, value Variant)
	Get(name StringName) Variant
	GetPropertyList() []PropertyInfo
	PropertyCanRevert(name StringName) bool
	PropertyGetRevert(name StringName) Variant
	GetOwner() Object
	GetPropertyState(add func(name StringName, value Variant))
	GetMethodList() []MethodInfo
	GetPropertyType(name StringName, valid bool) VariantType
	ValidateProperty(name StringName, property PropertyInfo) bool
	HasMethod(name StringName) bool
	Call(name StringName, args ...Variant) (Variant, error)
	Notification(notification int32, reversed bool)
	ToString(valid bool) String
	RefcountIncremented()
	RefcountDecremented()
	GetScript() Script
	IsPlaceholder() Bool
	SetFallback(name StringName, value Variant)
	GetFallback(name StringName) Variant
	GetLanguage() ScriptLanguage
}
