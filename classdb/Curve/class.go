// Package Curve provides methods for working with Curve object instances.
package Curve

import "unsafe"
import "reflect"
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
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Float"

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

/*
This resource describes a mathematical curve by defining a set of points and tangents at each point. By default, it ranges between [code]0[/code] and [code]1[/code] on the Y axis and positions points relative to the [code]0.5[/code] Y position.
See also [Gradient] which is designed for color interpolation. See also [Curve2D] and [Curve3D].
*/
type Instance [1]gdclass.Curve

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCurve() Instance
}

/*
Adds a point to the curve. For each side, if the [code]*_mode[/code] is [constant TANGENT_LINEAR], the [code]*_tangent[/code] angle (in degrees) uses the slope of the curve halfway to the adjacent point. Allows custom assignments to the [code]*_tangent[/code] angle if [code]*_mode[/code] is set to [constant TANGENT_FREE].
*/
func (self Instance) AddPoint(position Vector2.XY) int { //gd:Curve.add_point
	return int(int(class(self).AddPoint(gd.Vector2(position), gd.Float(0), gd.Float(0), 0, 0)))
}

/*
Removes the point at [param index] from the curve.
*/
func (self Instance) RemovePoint(index int) { //gd:Curve.remove_point
	class(self).RemovePoint(gd.Int(index))
}

/*
Removes all points from the curve.
*/
func (self Instance) ClearPoints() { //gd:Curve.clear_points
	class(self).ClearPoints()
}

/*
Returns the curve coordinates for the point at [param index].
*/
func (self Instance) GetPointPosition(index int) Vector2.XY { //gd:Curve.get_point_position
	return Vector2.XY(class(self).GetPointPosition(gd.Int(index)))
}

/*
Assigns the vertical position [param y] to the point at [param index].
*/
func (self Instance) SetPointValue(index int, y Float.X) { //gd:Curve.set_point_value
	class(self).SetPointValue(gd.Int(index), gd.Float(y))
}

/*
Sets the offset from [code]0.5[/code].
*/
func (self Instance) SetPointOffset(index int, offset Float.X) int { //gd:Curve.set_point_offset
	return int(int(class(self).SetPointOffset(gd.Int(index), gd.Float(offset))))
}

/*
Returns the Y value for the point that would exist at the X position [param offset] along the curve.
*/
func (self Instance) Sample(offset Float.X) Float.X { //gd:Curve.sample
	return Float.X(Float.X(class(self).Sample(gd.Float(offset))))
}

/*
Returns the Y value for the point that would exist at the X position [param offset] along the curve using the baked cache. Bakes the curve's points if not already baked.
*/
func (self Instance) SampleBaked(offset Float.X) Float.X { //gd:Curve.sample_baked
	return Float.X(Float.X(class(self).SampleBaked(gd.Float(offset))))
}

/*
Returns the left tangent angle (in degrees) for the point at [param index].
*/
func (self Instance) GetPointLeftTangent(index int) Float.X { //gd:Curve.get_point_left_tangent
	return Float.X(Float.X(class(self).GetPointLeftTangent(gd.Int(index))))
}

/*
Returns the right tangent angle (in degrees) for the point at [param index].
*/
func (self Instance) GetPointRightTangent(index int) Float.X { //gd:Curve.get_point_right_tangent
	return Float.X(Float.X(class(self).GetPointRightTangent(gd.Int(index))))
}

/*
Returns the left [enum TangentMode] for the point at [param index].
*/
func (self Instance) GetPointLeftMode(index int) gdclass.CurveTangentMode { //gd:Curve.get_point_left_mode
	return gdclass.CurveTangentMode(class(self).GetPointLeftMode(gd.Int(index)))
}

/*
Returns the right [enum TangentMode] for the point at [param index].
*/
func (self Instance) GetPointRightMode(index int) gdclass.CurveTangentMode { //gd:Curve.get_point_right_mode
	return gdclass.CurveTangentMode(class(self).GetPointRightMode(gd.Int(index)))
}

/*
Sets the left tangent angle for the point at [param index] to [param tangent].
*/
func (self Instance) SetPointLeftTangent(index int, tangent Float.X) { //gd:Curve.set_point_left_tangent
	class(self).SetPointLeftTangent(gd.Int(index), gd.Float(tangent))
}

/*
Sets the right tangent angle for the point at [param index] to [param tangent].
*/
func (self Instance) SetPointRightTangent(index int, tangent Float.X) { //gd:Curve.set_point_right_tangent
	class(self).SetPointRightTangent(gd.Int(index), gd.Float(tangent))
}

/*
Sets the left [enum TangentMode] for the point at [param index] to [param mode].
*/
func (self Instance) SetPointLeftMode(index int, mode gdclass.CurveTangentMode) { //gd:Curve.set_point_left_mode
	class(self).SetPointLeftMode(gd.Int(index), mode)
}

/*
Sets the right [enum TangentMode] for the point at [param index] to [param mode].
*/
func (self Instance) SetPointRightMode(index int, mode gdclass.CurveTangentMode) { //gd:Curve.set_point_right_mode
	class(self).SetPointRightMode(gd.Int(index), mode)
}

/*
Removes duplicate points, i.e. points that are less than 0.00001 units (engine epsilon value) away from their neighbor on the curve.
*/
func (self Instance) CleanDupes() { //gd:Curve.clean_dupes
	class(self).CleanDupes()
}

/*
Recomputes the baked cache of points for the curve.
*/
func (self Instance) Bake() { //gd:Curve.bake
	class(self).Bake()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Curve

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Curve"))
	casted := Instance{*(*gdclass.Curve)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) MinValue() Float.X {
	return Float.X(Float.X(class(self).GetMinValue()))
}

func (self Instance) SetMinValue(value Float.X) {
	class(self).SetMinValue(gd.Float(value))
}

func (self Instance) MaxValue() Float.X {
	return Float.X(Float.X(class(self).GetMaxValue()))
}

func (self Instance) SetMaxValue(value Float.X) {
	class(self).SetMaxValue(gd.Float(value))
}

func (self Instance) BakeResolution() int {
	return int(int(class(self).GetBakeResolution()))
}

func (self Instance) SetBakeResolution(value int) {
	class(self).SetBakeResolution(gd.Int(value))
}

func (self Instance) PointCount() int {
	return int(int(class(self).GetPointCount()))
}

func (self Instance) SetPointCount(value int) {
	class(self).SetPointCount(gd.Int(value))
}

//go:nosplit
func (self class) GetPointCount() gd.Int { //gd:Curve.get_point_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_get_point_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPointCount(count gd.Int) { //gd:Curve.set_point_count
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_set_point_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds a point to the curve. For each side, if the [code]*_mode[/code] is [constant TANGENT_LINEAR], the [code]*_tangent[/code] angle (in degrees) uses the slope of the curve halfway to the adjacent point. Allows custom assignments to the [code]*_tangent[/code] angle if [code]*_mode[/code] is set to [constant TANGENT_FREE].
*/
//go:nosplit
func (self class) AddPoint(position gd.Vector2, left_tangent gd.Float, right_tangent gd.Float, left_mode gdclass.CurveTangentMode, right_mode gdclass.CurveTangentMode) gd.Int { //gd:Curve.add_point
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, left_tangent)
	callframe.Arg(frame, right_tangent)
	callframe.Arg(frame, left_mode)
	callframe.Arg(frame, right_mode)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_add_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes the point at [param index] from the curve.
*/
//go:nosplit
func (self class) RemovePoint(index gd.Int) { //gd:Curve.remove_point
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_remove_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes all points from the curve.
*/
//go:nosplit
func (self class) ClearPoints() { //gd:Curve.clear_points
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_clear_points, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the curve coordinates for the point at [param index].
*/
//go:nosplit
func (self class) GetPointPosition(index gd.Int) gd.Vector2 { //gd:Curve.get_point_position
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_get_point_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Assigns the vertical position [param y] to the point at [param index].
*/
//go:nosplit
func (self class) SetPointValue(index gd.Int, y gd.Float) { //gd:Curve.set_point_value
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, y)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_set_point_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the offset from [code]0.5[/code].
*/
//go:nosplit
func (self class) SetPointOffset(index gd.Int, offset gd.Float) gd.Int { //gd:Curve.set_point_offset
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, offset)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_set_point_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the Y value for the point that would exist at the X position [param offset] along the curve.
*/
//go:nosplit
func (self class) Sample(offset gd.Float) gd.Float { //gd:Curve.sample
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_sample, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the Y value for the point that would exist at the X position [param offset] along the curve using the baked cache. Bakes the curve's points if not already baked.
*/
//go:nosplit
func (self class) SampleBaked(offset gd.Float) gd.Float { //gd:Curve.sample_baked
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_sample_baked, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the left tangent angle (in degrees) for the point at [param index].
*/
//go:nosplit
func (self class) GetPointLeftTangent(index gd.Int) gd.Float { //gd:Curve.get_point_left_tangent
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_get_point_left_tangent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the right tangent angle (in degrees) for the point at [param index].
*/
//go:nosplit
func (self class) GetPointRightTangent(index gd.Int) gd.Float { //gd:Curve.get_point_right_tangent
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_get_point_right_tangent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the left [enum TangentMode] for the point at [param index].
*/
//go:nosplit
func (self class) GetPointLeftMode(index gd.Int) gdclass.CurveTangentMode { //gd:Curve.get_point_left_mode
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gdclass.CurveTangentMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_get_point_left_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the right [enum TangentMode] for the point at [param index].
*/
//go:nosplit
func (self class) GetPointRightMode(index gd.Int) gdclass.CurveTangentMode { //gd:Curve.get_point_right_mode
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[gdclass.CurveTangentMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_get_point_right_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the left tangent angle for the point at [param index] to [param tangent].
*/
//go:nosplit
func (self class) SetPointLeftTangent(index gd.Int, tangent gd.Float) { //gd:Curve.set_point_left_tangent
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, tangent)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_set_point_left_tangent, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the right tangent angle for the point at [param index] to [param tangent].
*/
//go:nosplit
func (self class) SetPointRightTangent(index gd.Int, tangent gd.Float) { //gd:Curve.set_point_right_tangent
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, tangent)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_set_point_right_tangent, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the left [enum TangentMode] for the point at [param index] to [param mode].
*/
//go:nosplit
func (self class) SetPointLeftMode(index gd.Int, mode gdclass.CurveTangentMode) { //gd:Curve.set_point_left_mode
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_set_point_left_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the right [enum TangentMode] for the point at [param index] to [param mode].
*/
//go:nosplit
func (self class) SetPointRightMode(index gd.Int, mode gdclass.CurveTangentMode) { //gd:Curve.set_point_right_mode
	var frame = callframe.New()
	callframe.Arg(frame, index)
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_set_point_right_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMinValue() gd.Float { //gd:Curve.get_min_value
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_get_min_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMinValue(min gd.Float) { //gd:Curve.set_min_value
	var frame = callframe.New()
	callframe.Arg(frame, min)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_set_min_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMaxValue() gd.Float { //gd:Curve.get_max_value
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_get_max_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMaxValue(max gd.Float) { //gd:Curve.set_max_value
	var frame = callframe.New()
	callframe.Arg(frame, max)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_set_max_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes duplicate points, i.e. points that are less than 0.00001 units (engine epsilon value) away from their neighbor on the curve.
*/
//go:nosplit
func (self class) CleanDupes() { //gd:Curve.clean_dupes
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_clean_dupes, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Recomputes the baked cache of points for the curve.
*/
//go:nosplit
func (self class) Bake() { //gd:Curve.bake
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_bake, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBakeResolution() gd.Int { //gd:Curve.get_bake_resolution
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_get_bake_resolution, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBakeResolution(resolution gd.Int) { //gd:Curve.set_bake_resolution
	var frame = callframe.New()
	callframe.Arg(frame, resolution)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Curve.Bind_set_bake_resolution, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self Instance) OnRangeChanged(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("range_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsCurve() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCurve() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("Curve", func(ptr gd.Object) any { return [1]gdclass.Curve{*(*gdclass.Curve)(unsafe.Pointer(&ptr))} })
}

type TangentMode = gdclass.CurveTangentMode //gd:Curve.TangentMode

const (
	/*The tangent on this side of the point is user-defined.*/
	TangentFree TangentMode = 0
	/*The curve calculates the tangent on this side of the point as the slope halfway towards the adjacent point.*/
	TangentLinear TangentMode = 1
	/*The total number of available tangent modes.*/
	TangentModeCount TangentMode = 2
)
