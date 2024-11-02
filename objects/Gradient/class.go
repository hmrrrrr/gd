package Gradient

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
This resource describes a color transition by defining a set of colored points and how to interpolate between them.
See also [Curve] which supports more complex easing methods, but does not support colors.
*/
type Instance [1]classdb.Gradient

/*
Adds the specified color to the gradient, with the specified offset.
*/
func (self Instance) AddPoint(offset float64, color gd.Color) {
	class(self).AddPoint(gd.Float(offset), color)
}

/*
Removes the color at index [param point].
*/
func (self Instance) RemovePoint(point int) {
	class(self).RemovePoint(gd.Int(point))
}

/*
Sets the offset for the gradient color at index [param point].
*/
func (self Instance) SetOffset(point int, offset float64) {
	class(self).SetOffset(gd.Int(point), gd.Float(offset))
}

/*
Returns the offset of the gradient color at index [param point].
*/
func (self Instance) GetOffset(point int) float64 {
	return float64(float64(class(self).GetOffset(gd.Int(point))))
}

/*
Reverses/mirrors the gradient.
[b]Note:[/b] This method mirrors all points around the middle of the gradient, which may produce unexpected results when [member interpolation_mode] is set to [constant GRADIENT_INTERPOLATE_CONSTANT].
*/
func (self Instance) Reverse() {
	class(self).Reverse()
}

/*
Sets the color of the gradient color at index [param point].
*/
func (self Instance) SetColor(point int, color gd.Color) {
	class(self).SetColor(gd.Int(point), color)
}

/*
Returns the color of the gradient color at index [param point].
*/
func (self Instance) GetColor(point int) gd.Color {
	return gd.Color(class(self).GetColor(gd.Int(point)))
}

/*
Returns the interpolated color specified by [param offset].
*/
func (self Instance) Sample(offset float64) gd.Color {
	return gd.Color(class(self).Sample(gd.Float(offset)))
}

/*
Returns the number of colors in the gradient.
*/
func (self Instance) GetPointCount() int {
	return int(int(class(self).GetPointCount()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Gradient

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Gradient"))
	return Instance{classdb.Gradient(object)}
}

func (self Instance) InterpolationMode() classdb.GradientInterpolationMode {
	return classdb.GradientInterpolationMode(class(self).GetInterpolationMode())
}

func (self Instance) SetInterpolationMode(value classdb.GradientInterpolationMode) {
	class(self).SetInterpolationMode(value)
}

func (self Instance) InterpolationColorSpace() classdb.GradientColorSpace {
	return classdb.GradientColorSpace(class(self).GetInterpolationColorSpace())
}

func (self Instance) SetInterpolationColorSpace(value classdb.GradientColorSpace) {
	class(self).SetInterpolationColorSpace(value)
}

func (self Instance) Offsets() []float32 {
	return []float32(class(self).GetOffsets().AsSlice())
}

func (self Instance) SetOffsets(value []float32) {
	class(self).SetOffsets(gd.NewPackedFloat32Slice(value))
}

func (self Instance) Colors() []gd.Color {
	return []gd.Color(class(self).GetColors().AsSlice())
}

func (self Instance) SetColors(value []gd.Color) {
	class(self).SetColors(gd.NewPackedColorSlice(value))
}

/*
Adds the specified color to the gradient, with the specified offset.
*/
//go:nosplit
func (self class) AddPoint(offset gd.Float, color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_add_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the color at index [param point].
*/
//go:nosplit
func (self class) RemovePoint(point gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_remove_point, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the offset for the gradient color at index [param point].
*/
//go:nosplit
func (self class) SetOffset(point gd.Int, offset gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_set_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the offset of the gradient color at index [param point].
*/
//go:nosplit
func (self class) GetOffset(point gd.Int) gd.Float {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_get_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Reverses/mirrors the gradient.
[b]Note:[/b] This method mirrors all points around the middle of the gradient, which may produce unexpected results when [member interpolation_mode] is set to [constant GRADIENT_INTERPOLATE_CONSTANT].
*/
//go:nosplit
func (self class) Reverse() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_reverse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the color of the gradient color at index [param point].
*/
//go:nosplit
func (self class) SetColor(point gd.Int, color gd.Color) {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the color of the gradient color at index [param point].
*/
//go:nosplit
func (self class) GetColor(point gd.Int) gd.Color {
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_get_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the interpolated color specified by [param offset].
*/
//go:nosplit
func (self class) Sample(offset gd.Float) gd.Color {
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_sample, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of colors in the gradient.
*/
//go:nosplit
func (self class) GetPointCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_get_point_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOffsets(offsets gd.PackedFloat32Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(offsets))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_set_offsets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetOffsets() gd.PackedFloat32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_get_offsets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedFloat32Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetColors(colors gd.PackedColorArray) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(colors))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_set_colors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetColors() gd.PackedColorArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_get_colors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedColorArray](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInterpolationMode(interpolation_mode classdb.GradientInterpolationMode) {
	var frame = callframe.New()
	callframe.Arg(frame, interpolation_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_set_interpolation_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetInterpolationMode() classdb.GradientInterpolationMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GradientInterpolationMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_get_interpolation_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInterpolationColorSpace(interpolation_color_space classdb.GradientColorSpace) {
	var frame = callframe.New()
	callframe.Arg(frame, interpolation_color_space)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_set_interpolation_color_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetInterpolationColorSpace() classdb.GradientColorSpace {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.GradientColorSpace](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_get_interpolation_color_space, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsGradient() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsGradient() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
func init() { classdb.Register("Gradient", func(ptr gd.Object) any { return classdb.Gradient(ptr) }) }

type InterpolationMode = classdb.GradientInterpolationMode

const (
	/*Linear interpolation.*/
	GradientInterpolateLinear InterpolationMode = 0
	/*Constant interpolation, color changes abruptly at each point and stays uniform between. This might cause visible aliasing when used for a gradient texture in some cases.*/
	GradientInterpolateConstant InterpolationMode = 1
	/*Cubic interpolation.*/
	GradientInterpolateCubic InterpolationMode = 2
)

type ColorSpace = classdb.GradientColorSpace

const (
	/*sRGB color space.*/
	GradientColorSpaceSrgb ColorSpace = 0
	/*Linear sRGB color space.*/
	GradientColorSpaceLinearSrgb ColorSpace = 1
	/*[url=https://bottosson.github.io/posts/oklab/]Oklab[/url] color space. This color space provides a smooth and uniform-looking transition between colors.*/
	GradientColorSpaceOklab ColorSpace = 2
)
