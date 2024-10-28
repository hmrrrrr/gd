package EditorSyntaxHighlighter

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/SyntaxHighlighter"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Base class that all [SyntaxHighlighter]s used by the [ScriptEditor] extend from.
Add a syntax highlighter to an individual script by calling [method ScriptEditorBase.add_syntax_highlighter]. To apply to all scripts on open, call [method ScriptEditor.register_syntax_highlighter].
	// EditorSyntaxHighlighter methods that can be overridden by a [Class] that extends it.
	type EditorSyntaxHighlighter interface {
		//Virtual method which can be overridden to return the syntax highlighter name.
		GetName() string
		//Virtual method which can be overridden to return the supported language names.
		GetSupportedLanguages() []string
	}

*/
type Go [1]classdb.EditorSyntaxHighlighter

/*
Virtual method which can be overridden to return the syntax highlighter name.
*/
func (Go) _get_name(impl func(ptr unsafe.Pointer) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
ptr, ok := discreet.End(gd.NewString(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method which can be overridden to return the supported language names.
*/
func (Go) _get_supported_languages(impl func(ptr unsafe.Pointer) []string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
ptr, ok := discreet.End(gd.NewPackedStringSlice(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.EditorSyntaxHighlighter
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorSyntaxHighlighter"))
	return Go{classdb.EditorSyntaxHighlighter(object)}
}

/*
Virtual method which can be overridden to return the syntax highlighter name.
*/
func (class) _get_name(impl func(ptr unsafe.Pointer) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
ptr, ok := discreet.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Virtual method which can be overridden to return the supported language names.
*/
func (class) _get_supported_languages(impl func(ptr unsafe.Pointer) gd.PackedStringArray, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
ptr, ok := discreet.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (self class) AsEditorSyntaxHighlighter() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorSyntaxHighlighter() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsSyntaxHighlighter() SyntaxHighlighter.GD { return *((*SyntaxHighlighter.GD)(unsafe.Pointer(&self))) }
func (self Go) AsSyntaxHighlighter() SyntaxHighlighter.Go { return *((*SyntaxHighlighter.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_name": return reflect.ValueOf(self._get_name);
	case "_get_supported_languages": return reflect.ValueOf(self._get_supported_languages);
	default: return gd.VirtualByName(self.AsSyntaxHighlighter(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_get_name": return reflect.ValueOf(self._get_name);
	case "_get_supported_languages": return reflect.ValueOf(self._get_supported_languages);
	default: return gd.VirtualByName(self.AsSyntaxHighlighter(), name)
	}
}
func init() {classdb.Register("EditorSyntaxHighlighter", func(ptr gd.Object) any { return classdb.EditorSyntaxHighlighter(ptr) })}
