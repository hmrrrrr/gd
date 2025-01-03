package ShaderInclude

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A shader include file, saved with the [code].gdshaderinc[/code] extension. This class allows you to define a custom shader snippet that can be included in a [Shader] by using the preprocessor directive [code]#include[/code], followed by the file path (e.g. [code]#include "res://shader_lib.gdshaderinc"[/code]). The snippet doesn't have to be a valid shader on its own.
*/
type Instance [1]classdb.ShaderInclude
type Any interface {
	gd.IsClass
	AsShaderInclude() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.ShaderInclude

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ShaderInclude"))
	return Instance{*(*classdb.ShaderInclude)(unsafe.Pointer(&object))}
}

func (self Instance) Code() string {
	return string(class(self).GetCode().String())
}

func (self Instance) SetCode(value string) {
	class(self).SetCode(gd.NewString(value))
}

//go:nosplit
func (self class) SetCode(code gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(code))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShaderInclude.Bind_set_code, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCode() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ShaderInclude.Bind_get_code, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsShaderInclude() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsShaderInclude() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {
	classdb.Register("ShaderInclude", func(ptr gd.Object) any {
		return [1]classdb.ShaderInclude{*(*classdb.ShaderInclude)(unsafe.Pointer(&ptr))}
	})
}
