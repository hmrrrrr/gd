package GradientTexture2D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Texture2D"
import "grow.graphics/gd/gdclass/Texture"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
A 2D texture that obtains colors from a [Gradient] to fill the texture data. This texture is able to transform a color transition into different patterns such as a linear or a radial gradient. The gradient is sampled individually for each pixel so it does not necessarily represent an exact copy of the gradient(see [member width] and [member height]). See also [GradientTexture1D], [CurveTexture] and [CurveXYZTexture].

*/
type Go [1]classdb.GradientTexture2D
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GradientTexture2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GradientTexture2D"))
	return Go{classdb.GradientTexture2D(object)}
}

func (self Go) Gradient() gdclass.Gradient {
		return gdclass.Gradient(class(self).GetGradient())
}

func (self Go) SetGradient(value gdclass.Gradient) {
	class(self).SetGradient(value)
}

func (self Go) UseHdr() bool {
		return bool(class(self).IsUsingHdr())
}

func (self Go) SetUseHdr(value bool) {
	class(self).SetUseHdr(value)
}

func (self Go) Fill() classdb.GradientTexture2DFill {
		return classdb.GradientTexture2DFill(class(self).GetFill())
}

func (self Go) SetFill(value classdb.GradientTexture2DFill) {
	class(self).SetFill(value)
}

func (self Go) FillFrom() gd.Vector2 {
		return gd.Vector2(class(self).GetFillFrom())
}

func (self Go) SetFillFrom(value gd.Vector2) {
	class(self).SetFillFrom(value)
}

func (self Go) FillTo() gd.Vector2 {
		return gd.Vector2(class(self).GetFillTo())
}

func (self Go) SetFillTo(value gd.Vector2) {
	class(self).SetFillTo(value)
}

func (self Go) Repeat() classdb.GradientTexture2DRepeat {
		return classdb.GradientTexture2DRepeat(class(self).GetRepeat())
}

func (self Go) SetRepeat(value classdb.GradientTexture2DRepeat) {
	class(self).SetRepeat(value)
}

//go:nosplit
func (self class) SetGradient(gradient gdclass.Gradient)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(gradient[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_set_gradient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetGradient() gdclass.Gradient {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_get_gradient, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Gradient{classdb.Gradient(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetWidth(width gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_set_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetHeight(height gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_set_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetUseHdr(enabled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_set_use_hdr, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsUsingHdr() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_is_using_hdr, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFill(fill classdb.GradientTexture2DFill)  {
	var frame = callframe.New()
	callframe.Arg(frame, fill)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_set_fill, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFill() classdb.GradientTexture2DFill {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GradientTexture2DFill](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_get_fill, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFillFrom(fill_from gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, fill_from)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_set_fill_from, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFillFrom() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_get_fill_from, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFillTo(fill_to gd.Vector2)  {
	var frame = callframe.New()
	callframe.Arg(frame, fill_to)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_set_fill_to, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFillTo() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_get_fill_to, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRepeat(repeat classdb.GradientTexture2DRepeat)  {
	var frame = callframe.New()
	callframe.Arg(frame, repeat)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_set_repeat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetRepeat() classdb.GradientTexture2DRepeat {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GradientTexture2DRepeat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GradientTexture2D.Bind_get_repeat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsGradientTexture2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsGradientTexture2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsTexture2D() Texture2D.GD { return *((*Texture2D.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTexture2D() Texture2D.Go { return *((*Texture2D.Go)(unsafe.Pointer(&self))) }
func (self class) AsTexture() Texture.GD { return *((*Texture.GD)(unsafe.Pointer(&self))) }
func (self Go) AsTexture() Texture.Go { return *((*Texture.Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTexture2D(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsTexture2D(), name)
	}
}
func init() {classdb.Register("GradientTexture2D", func(ptr gd.Object) any { return classdb.GradientTexture2D(ptr) })}
type Fill = classdb.GradientTexture2DFill

const (
/*The colors are linearly interpolated in a straight line.*/
	FillLinear Fill = 0
/*The colors are linearly interpolated in a circular pattern.*/
	FillRadial Fill = 1
/*The colors are linearly interpolated in a square pattern.*/
	FillSquare Fill = 2
)
type Repeat = classdb.GradientTexture2DRepeat

const (
/*The gradient fill is restricted to the range defined by [member fill_from] to [member fill_to] offsets.*/
	RepeatNone Repeat = 0
/*The texture is filled starting from [member fill_from] to [member fill_to] offsets, repeating the same pattern in both directions.*/
	RepeatDefault Repeat = 1
/*The texture is filled starting from [member fill_from] to [member fill_to] offsets, mirroring the pattern in both directions.*/
	RepeatMirror Repeat = 2
)
