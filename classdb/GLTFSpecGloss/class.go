// Package GLTFSpecGloss provides methods for working with GLTFSpecGloss object instances.
package GLTFSpecGloss

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Color"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
KHR_materials_pbrSpecularGlossiness is an archived GLTF extension. This means that it is deprecated and not recommended for new files. However, it is still supported for loading old files.
*/
type Instance [1]gdclass.GLTFSpecGloss
type Any interface {
	gd.IsClass
	AsGLTFSpecGloss() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GLTFSpecGloss

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFSpecGloss"))
	return Instance{*(*gdclass.GLTFSpecGloss)(unsafe.Pointer(&object))}
}

func (self Instance) DiffuseImg() [1]gdclass.Image {
	return [1]gdclass.Image(class(self).GetDiffuseImg())
}

func (self Instance) SetDiffuseImg(value [1]gdclass.Image) {
	class(self).SetDiffuseImg(value)
}

func (self Instance) DiffuseFactor() Color.RGBA {
	return Color.RGBA(class(self).GetDiffuseFactor())
}

func (self Instance) SetDiffuseFactor(value Color.RGBA) {
	class(self).SetDiffuseFactor(gd.Color(value))
}

func (self Instance) GlossFactor() Float.X {
	return Float.X(Float.X(class(self).GetGlossFactor()))
}

func (self Instance) SetGlossFactor(value Float.X) {
	class(self).SetGlossFactor(gd.Float(value))
}

func (self Instance) SpecularFactor() Color.RGBA {
	return Color.RGBA(class(self).GetSpecularFactor())
}

func (self Instance) SetSpecularFactor(value Color.RGBA) {
	class(self).SetSpecularFactor(gd.Color(value))
}

func (self Instance) SpecGlossImg() [1]gdclass.Image {
	return [1]gdclass.Image(class(self).GetSpecGlossImg())
}

func (self Instance) SetSpecGlossImg(value [1]gdclass.Image) {
	class(self).SetSpecGlossImg(value)
}

//go:nosplit
func (self class) GetDiffuseImg() [1]gdclass.Image {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSpecGloss.Bind_get_diffuse_img, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Image{gd.PointerWithOwnershipTransferredToGo[gdclass.Image](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDiffuseImg(diffuse_img [1]gdclass.Image) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(diffuse_img[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSpecGloss.Bind_set_diffuse_img, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDiffuseFactor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSpecGloss.Bind_get_diffuse_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDiffuseFactor(diffuse_factor gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, diffuse_factor)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSpecGloss.Bind_set_diffuse_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlossFactor() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSpecGloss.Bind_get_gloss_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlossFactor(gloss_factor gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, gloss_factor)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSpecGloss.Bind_set_gloss_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSpecularFactor() gd.Color {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSpecGloss.Bind_get_specular_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSpecularFactor(specular_factor gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, specular_factor)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSpecGloss.Bind_set_specular_factor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSpecGlossImg() [1]gdclass.Image {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSpecGloss.Bind_get_spec_gloss_img, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Image{gd.PointerWithOwnershipTransferredToGo[gdclass.Image](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSpecGlossImg(spec_gloss_img [1]gdclass.Image) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(spec_gloss_img[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSpecGloss.Bind_set_spec_gloss_img, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsGLTFSpecGloss() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGLTFSpecGloss() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("GLTFSpecGloss", func(ptr gd.Object) any {
		return [1]gdclass.GLTFSpecGloss{*(*gdclass.GLTFSpecGloss)(unsafe.Pointer(&ptr))}
	})
}
