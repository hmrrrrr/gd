// Package EditorSyntaxHighlighter provides methods for working with EditorSyntaxHighlighter object instances.
package EditorSyntaxHighlighter

import "unsafe"
import "reflect"
import "slices"
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
import "graphics.gd/variant/Packed"
import "graphics.gd/classdb/SyntaxHighlighter"
import "graphics.gd/classdb/Resource"

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
var _ = slices.Delete[[]struct{}, struct{}]

/*
Base class that all [SyntaxHighlighter]s used by the [ScriptEditor] extend from.
Add a syntax highlighter to an individual script by calling [method ScriptEditorBase.add_syntax_highlighter]. To apply to all scripts on open, call [method ScriptEditor.register_syntax_highlighter].

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=EditorSyntaxHighlighter)
*/
type Instance [1]gdclass.EditorSyntaxHighlighter

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorSyntaxHighlighter() Instance
}
type Interface interface {
	//Virtual method which can be overridden to return the syntax highlighter name.
	GetName() string
	//Virtual method which can be overridden to return the supported language names.
	GetSupportedLanguages() []string
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) GetName() (_ string)                 { return }
func (self implementation) GetSupportedLanguages() (_ []string) { return }

/*
Virtual method which can be overridden to return the syntax highlighter name.
*/
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

/*
Virtual method which can be overridden to return the supported language names.
*/
func (Instance) _get_supported_languages(impl func(ptr unsafe.Pointer) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalPackedStrings(Packed.MakeStrings(ret...)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorSyntaxHighlighter

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorSyntaxHighlighter"))
	casted := Instance{*(*gdclass.EditorSyntaxHighlighter)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Virtual method which can be overridden to return the syntax highlighter name.
*/
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

/*
Virtual method which can be overridden to return the supported language names.
*/
func (class) _get_supported_languages(impl func(ptr unsafe.Pointer) Packed.Strings) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalPackedStrings(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (self class) AsEditorSyntaxHighlighter() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorSyntaxHighlighter() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsSyntaxHighlighter() SyntaxHighlighter.Advanced {
	return *((*SyntaxHighlighter.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsSyntaxHighlighter() SyntaxHighlighter.Instance {
	return *((*SyntaxHighlighter.Instance)(unsafe.Pointer(&self)))
}
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
	case "_get_name":
		return reflect.ValueOf(self._get_name)
	case "_get_supported_languages":
		return reflect.ValueOf(self._get_supported_languages)
	default:
		return gd.VirtualByName(SyntaxHighlighter.Advanced(self.AsSyntaxHighlighter()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_name":
		return reflect.ValueOf(self._get_name)
	case "_get_supported_languages":
		return reflect.ValueOf(self._get_supported_languages)
	default:
		return gd.VirtualByName(SyntaxHighlighter.Instance(self.AsSyntaxHighlighter()), name)
	}
}
func init() {
	gdclass.Register("EditorSyntaxHighlighter", func(ptr gd.Object) any {
		return [1]gdclass.EditorSyntaxHighlighter{*(*gdclass.EditorSyntaxHighlighter)(unsafe.Pointer(&ptr))}
	})
}
