package CurveTexture

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Texture2D"
import "grow.graphics/gd/object/Texture"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A 1D texture where pixel brightness corresponds to points on a [Curve] resource, either in grayscale or in red. This visual representation simplifies the task of saving curves as image files.
If you need to store up to 3 curves within a single texture, use [CurveXYZTexture] instead. See also [GradientTexture1D] and [GradientTexture2D].

*/
type Simple [1]classdb.CurveTexture
func (self Simple) SetWidth(width int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetWidth(gd.Int(width))
}
func (self Simple) SetCurve(curve [1]classdb.Curve) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCurve(curve)
}
func (self Simple) GetCurve() [1]classdb.Curve {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Curve(Expert(self).GetCurve(gc))
}
func (self Simple) SetTextureMode(texture_mode classdb.CurveTextureTextureMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTextureMode(texture_mode)
}
func (self Simple) GetTextureMode() classdb.CurveTextureTextureMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.CurveTextureTextureMode(Expert(self).GetTextureMode())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.CurveTexture
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetWidth(width gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CurveTexture.Bind_set_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetCurve(curve object.Curve)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(curve[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CurveTexture.Bind_set_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetCurve(ctx gd.Lifetime) object.Curve {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CurveTexture.Bind_get_curve, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Curve
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTextureMode(texture_mode classdb.CurveTextureTextureMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texture_mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CurveTexture.Bind_set_texture_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTextureMode() classdb.CurveTextureTextureMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.CurveTextureTextureMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CurveTexture.Bind_get_texture_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsCurveTexture() Expert { return self[0].AsCurveTexture() }


//go:nosplit
func (self Simple) AsCurveTexture() Simple { return self[0].AsCurveTexture() }


//go:nosplit
func (self class) AsTexture2D() Texture2D.Expert { return self[0].AsTexture2D() }


//go:nosplit
func (self Simple) AsTexture2D() Texture2D.Simple { return self[0].AsTexture2D() }


//go:nosplit
func (self class) AsTexture() Texture.Expert { return self[0].AsTexture() }


//go:nosplit
func (self Simple) AsTexture() Texture.Simple { return self[0].AsTexture() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("CurveTexture", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type TextureMode = classdb.CurveTextureTextureMode

const (
/*Store the curve equally across the red, green and blue channels. This uses more video memory, but is more compatible with shaders that only read the green and blue values.*/
	TextureModeRgb TextureMode = 0
/*Store the curve only in the red channel. This saves video memory, but some custom shaders may not be able to work with this.*/
	TextureModeRed TextureMode = 1
)
