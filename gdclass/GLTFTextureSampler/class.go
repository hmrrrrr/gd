package GLTFTextureSampler

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Represents a texture sampler as defined by the base GLTF spec. Texture samplers in GLTF specify how to sample data from the texture's base image, when rendering the texture on an object.

*/
type Go [1]classdb.GLTFTextureSampler
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GLTFTextureSampler
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFTextureSampler"))
	return Go{classdb.GLTFTextureSampler(object)}
}

func (self Go) MagFilter() int {
		return int(int(class(self).GetMagFilter()))
}

func (self Go) SetMagFilter(value int) {
	class(self).SetMagFilter(gd.Int(value))
}

func (self Go) MinFilter() int {
		return int(int(class(self).GetMinFilter()))
}

func (self Go) SetMinFilter(value int) {
	class(self).SetMinFilter(gd.Int(value))
}

func (self Go) WrapS() int {
		return int(int(class(self).GetWrapS()))
}

func (self Go) SetWrapS(value int) {
	class(self).SetWrapS(gd.Int(value))
}

func (self Go) WrapT() int {
		return int(int(class(self).GetWrapT()))
}

func (self Go) SetWrapT(value int) {
	class(self).SetWrapT(gd.Int(value))
}

//go:nosplit
func (self class) GetMagFilter() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFTextureSampler.Bind_get_mag_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMagFilter(filter_mode gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, filter_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFTextureSampler.Bind_set_mag_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMinFilter() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFTextureSampler.Bind_get_min_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMinFilter(filter_mode gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, filter_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFTextureSampler.Bind_set_min_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetWrapS() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFTextureSampler.Bind_get_wrap_s, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetWrapS(wrap_mode gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, wrap_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFTextureSampler.Bind_set_wrap_s, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetWrapT() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFTextureSampler.Bind_get_wrap_t, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetWrapT(wrap_mode gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, wrap_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFTextureSampler.Bind_set_wrap_t, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsGLTFTextureSampler() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsGLTFTextureSampler() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("GLTFTextureSampler", func(ptr gd.Object) any { return classdb.GLTFTextureSampler(ptr) })}
