// Package Gradient provides methods for working with Gradient object instances.
package Gradient

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Color"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

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
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
This resource describes a color transition by defining a set of colored points and how to interpolate between them.
See also [Curve] which supports more complex easing methods, but does not support colors.
*/
type Instance [1]gdclass.Gradient

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGradient() Instance
}

/*
Adds the specified color to the gradient, with the specified offset.
*/
func (self Instance) AddPoint(offset Float.X, color Color.RGBA) { //gd:Gradient.add_point
	class(self).AddPoint(float64(offset), Color.RGBA(color))
}

/*
Removes the color at index [param point].
*/
func (self Instance) RemovePoint(point int) { //gd:Gradient.remove_point
	class(self).RemovePoint(int64(point))
}

/*
Sets the offset for the gradient color at index [param point].
*/
func (self Instance) SetOffset(point int, offset Float.X) { //gd:Gradient.set_offset
	class(self).SetOffset(int64(point), float64(offset))
}

/*
Returns the offset of the gradient color at index [param point].
*/
func (self Instance) GetOffset(point int) Float.X { //gd:Gradient.get_offset
	return Float.X(Float.X(class(self).GetOffset(int64(point))))
}

/*
Reverses/mirrors the gradient.
[b]Note:[/b] This method mirrors all points around the middle of the gradient, which may produce unexpected results when [member interpolation_mode] is set to [constant GRADIENT_INTERPOLATE_CONSTANT].
*/
func (self Instance) Reverse() { //gd:Gradient.reverse
	class(self).Reverse()
}

/*
Sets the color of the gradient color at index [param point].
*/
func (self Instance) SetColor(point int, color Color.RGBA) { //gd:Gradient.set_color
	class(self).SetColor(int64(point), Color.RGBA(color))
}

/*
Returns the color of the gradient color at index [param point].
*/
func (self Instance) GetColor(point int) Color.RGBA { //gd:Gradient.get_color
	return Color.RGBA(class(self).GetColor(int64(point)))
}

/*
Returns the interpolated color specified by [param offset].
*/
func (self Instance) Sample(offset Float.X) Color.RGBA { //gd:Gradient.sample
	return Color.RGBA(class(self).Sample(float64(offset)))
}

/*
Returns the number of colors in the gradient.
*/
func (self Instance) GetPointCount() int { //gd:Gradient.get_point_count
	return int(int(class(self).GetPointCount()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Gradient

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Gradient"))
	casted := Instance{*(*gdclass.Gradient)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) InterpolationMode() gdclass.GradientInterpolationMode {
	return gdclass.GradientInterpolationMode(class(self).GetInterpolationMode())
}

func (self Instance) SetInterpolationMode(value gdclass.GradientInterpolationMode) {
	class(self).SetInterpolationMode(value)
}

func (self Instance) InterpolationColorSpace() gdclass.GradientColorSpace {
	return gdclass.GradientColorSpace(class(self).GetInterpolationColorSpace())
}

func (self Instance) SetInterpolationColorSpace(value gdclass.GradientColorSpace) {
	class(self).SetInterpolationColorSpace(value)
}

func (self Instance) Offsets() []float32 {
	return []float32(slices.Collect(class(self).GetOffsets().Values()))
}

func (self Instance) SetOffsets(value []float32) {
	class(self).SetOffsets(Packed.New(value...))
}

func (self Instance) Colors() []Color.RGBA {
	return []Color.RGBA(slices.Collect(class(self).GetColors().Values()))
}

func (self Instance) SetColors(value []Color.RGBA) {
	class(self).SetColors(Packed.New(value...))
}

/*
Adds the specified color to the gradient, with the specified offset.
*/
//go:nosplit
func (self class) AddPoint(offset float64, color Color.RGBA) { //gd:Gradient.add_point
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_add_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes the color at index [param point].
*/
//go:nosplit
func (self class) RemovePoint(point int64) { //gd:Gradient.remove_point
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_remove_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the offset for the gradient color at index [param point].
*/
//go:nosplit
func (self class) SetOffset(point int64, offset float64) { //gd:Gradient.set_offset
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_set_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the offset of the gradient color at index [param point].
*/
//go:nosplit
func (self class) GetOffset(point int64) float64 { //gd:Gradient.get_offset
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_get_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Reverses/mirrors the gradient.
[b]Note:[/b] This method mirrors all points around the middle of the gradient, which may produce unexpected results when [member interpolation_mode] is set to [constant GRADIENT_INTERPOLATE_CONSTANT].
*/
//go:nosplit
func (self class) Reverse() { //gd:Gradient.reverse
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_reverse, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the color of the gradient color at index [param point].
*/
//go:nosplit
func (self class) SetColor(point int64, color Color.RGBA) { //gd:Gradient.set_color
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the color of the gradient color at index [param point].
*/
//go:nosplit
func (self class) GetColor(point int64) Color.RGBA { //gd:Gradient.get_color
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_get_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the interpolated color specified by [param offset].
*/
//go:nosplit
func (self class) Sample(offset float64) Color.RGBA { //gd:Gradient.sample
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_sample, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of colors in the gradient.
*/
//go:nosplit
func (self class) GetPointCount() int64 { //gd:Gradient.get_point_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_get_point_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOffsets(offsets Packed.Array[float32]) { //gd:Gradient.set_offsets
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedFloat32Array, float32](offsets)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_set_offsets, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOffsets() Packed.Array[float32] { //gd:Gradient.get_offsets
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_get_offsets, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[float32](Array.Through(gd.PackedProxy[gd.PackedFloat32Array, float32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetColors(colors Packed.Array[Color.RGBA]) { //gd:Gradient.set_colors
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedColorArray, Color.RGBA](colors)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_set_colors, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetColors() Packed.Array[Color.RGBA] { //gd:Gradient.get_colors
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_get_colors, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[Color.RGBA](Array.Through(gd.PackedProxy[gd.PackedColorArray, Color.RGBA]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInterpolationMode(interpolation_mode gdclass.GradientInterpolationMode) { //gd:Gradient.set_interpolation_mode
	var frame = callframe.New()
	callframe.Arg(frame, interpolation_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_set_interpolation_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetInterpolationMode() gdclass.GradientInterpolationMode { //gd:Gradient.get_interpolation_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.GradientInterpolationMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_get_interpolation_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetInterpolationColorSpace(interpolation_color_space gdclass.GradientColorSpace) { //gd:Gradient.set_interpolation_color_space
	var frame = callframe.New()
	callframe.Arg(frame, interpolation_color_space)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_set_interpolation_color_space, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetInterpolationColorSpace() gdclass.GradientColorSpace { //gd:Gradient.get_interpolation_color_space
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.GradientColorSpace](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Gradient.Bind_get_interpolation_color_space, self.AsObject(), frame.Array(0), r_ret.Addr())
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
	gdclass.Register("Gradient", func(ptr gd.Object) any { return [1]gdclass.Gradient{*(*gdclass.Gradient)(unsafe.Pointer(&ptr))} })
}

type InterpolationMode = gdclass.GradientInterpolationMode //gd:Gradient.InterpolationMode

const (
	/*Linear interpolation.*/
	GradientInterpolateLinear InterpolationMode = 0
	/*Constant interpolation, color changes abruptly at each point and stays uniform between. This might cause visible aliasing when used for a gradient texture in some cases.*/
	GradientInterpolateConstant InterpolationMode = 1
	/*Cubic interpolation.*/
	GradientInterpolateCubic InterpolationMode = 2
)

type ColorSpace = gdclass.GradientColorSpace //gd:Gradient.ColorSpace

const (
	/*sRGB color space.*/
	GradientColorSpaceSrgb ColorSpace = 0
	/*Linear sRGB color space.*/
	GradientColorSpaceLinearSrgb ColorSpace = 1
	/*[url=https://bottosson.github.io/posts/oklab/]Oklab[/url] color space. This color space provides a smooth and uniform-looking transition between colors.*/
	GradientColorSpaceOklab ColorSpace = 2
)
