package GLTFSpecGloss

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
KHR_materials_pbrSpecularGlossiness is an archived GLTF extension. This means that it is deprecated and not recommended for new files. However, it is still supported for loading old files.
*/
type Instance [1]classdb.GLTFSpecGloss

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.GLTFSpecGloss

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFSpecGloss"))
	return Instance{classdb.GLTFSpecGloss(object)}
}

func (self Instance) DiffuseImg() gdclass.Image {
	return gdclass.Image(class(self).GetDiffuseImg())
}

func (self Instance) SetDiffuseImg(value gdclass.Image) {
	class(self).SetDiffuseImg(value)
}

func (self Instance) DiffuseFactor() gd.Color {
	return gd.Color(class(self).GetDiffuseFactor())
}

func (self Instance) SetDiffuseFactor(value gd.Color) {
	class(self).SetDiffuseFactor(value)
}

func (self Instance) GlossFactor() float64 {
	return float64(float64(class(self).GetGlossFactor()))
}

func (self Instance) SetGlossFactor(value float64) {
	class(self).SetGlossFactor(gd.Float(value))
}

func (self Instance) SpecularFactor() gd.Color {
	return gd.Color(class(self).GetSpecularFactor())
}

func (self Instance) SetSpecularFactor(value gd.Color) {
	class(self).SetSpecularFactor(value)
}

func (self Instance) SpecGlossImg() gdclass.Image {
	return gdclass.Image(class(self).GetSpecGlossImg())
}

func (self Instance) SetSpecGlossImg(value gdclass.Image) {
	class(self).SetSpecGlossImg(value)
}

//go:nosplit
func (self class) GetDiffuseImg() gdclass.Image {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSpecGloss.Bind_get_diffuse_img, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Image{classdb.Image(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDiffuseImg(diffuse_img gdclass.Image) {
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
func (self class) GetSpecGlossImg() gdclass.Image {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFSpecGloss.Bind_get_spec_gloss_img, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Image{classdb.Image(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSpecGlossImg(spec_gloss_img gdclass.Image) {
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
	classdb.Register("GLTFSpecGloss", func(ptr gd.Object) any { return classdb.GLTFSpecGloss(ptr) })
}
