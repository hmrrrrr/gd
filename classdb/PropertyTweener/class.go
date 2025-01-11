// Package PropertyTweener provides methods for working with PropertyTweener object instances.
package PropertyTweener

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Tweener"
import "graphics.gd/variant/Float"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type variantPointers = gd.VariantPointers
type signalPointers = gd.SignalPointers
type callablePointers = gd.CallablePointers

/*
[PropertyTweener] is used to interpolate a property in an object. See [method Tween.tween_property] for more usage information.
[b]Note:[/b] [method Tween.tween_property] is the only correct way to create [PropertyTweener]. Any [PropertyTweener] created manually will not function correctly.
*/
type Instance [1]gdclass.PropertyTweener

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPropertyTweener() Instance
}

/*
Sets a custom initial value to the [PropertyTweener].
[b]Example:[/b]
[codeblock]
var tween = get_tree().create_tween()
tween.tween_property(self, "position", Vector2(200, 100), 1).from(Vector2(100, 100)) #this will move the node from position (100, 100) to (200, 100)
[/codeblock]
*/
func (self Instance) From(value any) [1]gdclass.PropertyTweener {
	return [1]gdclass.PropertyTweener(class(self).From(gd.NewVariant(value)))
}

/*
Makes the [PropertyTweener] use the current property value (i.e. at the time of creating this [PropertyTweener]) as a starting point. This is equivalent of using [method from] with the current value. These two calls will do the same:
[codeblock]
tween.tween_property(self, "position", Vector2(200, 100), 1).from(position)
tween.tween_property(self, "position", Vector2(200, 100), 1).from_current()
[/codeblock]
*/
func (self Instance) FromCurrent() [1]gdclass.PropertyTweener {
	return [1]gdclass.PropertyTweener(class(self).FromCurrent())
}

/*
When called, the final value will be used as a relative value instead.
[b]Example:[/b]
[codeblock]
var tween = get_tree().create_tween()
tween.tween_property(self, "position", Vector2.RIGHT * 100, 1).as_relative() #the node will move by 100 pixels to the right
[/codeblock]
*/
func (self Instance) AsRelative() [1]gdclass.PropertyTweener {
	return [1]gdclass.PropertyTweener(class(self).AsRelative())
}

/*
Sets the type of used transition from [enum Tween.TransitionType]. If not set, the default transition is used from the [Tween] that contains this Tweener.
*/
func (self Instance) SetTrans(trans gdclass.TweenTransitionType) [1]gdclass.PropertyTweener {
	return [1]gdclass.PropertyTweener(class(self).SetTrans(trans))
}

/*
Sets the type of used easing from [enum Tween.EaseType]. If not set, the default easing is used from the [Tween] that contains this Tweener.
*/
func (self Instance) SetEase(ease gdclass.TweenEaseType) [1]gdclass.PropertyTweener {
	return [1]gdclass.PropertyTweener(class(self).SetEase(ease))
}

/*
Allows interpolating the value with a custom easing function. The provided [param interpolator_method] will be called with a value ranging from [code]0.0[/code] to [code]1.0[/code] and is expected to return a value within the same range (values outside the range can be used for overshoot). The return value of the method is then used for interpolation between initial and final value. Note that the parameter passed to the method is still subject to the tweener's own easing.
[b]Example:[/b]
[codeblock]
@export var curve: Curve

func _ready():

	var tween = create_tween()
	# Interpolate the value using a custom curve.
	tween.tween_property(self, "position:x", 300, 1).as_relative().set_custom_interpolator(tween_curve)

func tween_curve(v):

	return curve.sample_baked(v)

[/codeblock]
*/
func (self Instance) SetCustomInterpolator(interpolator_method func(Float.X) Float.X) [1]gdclass.PropertyTweener {
	return [1]gdclass.PropertyTweener(class(self).SetCustomInterpolator(gd.NewCallable(interpolator_method)))
}

/*
Sets the time in seconds after which the [PropertyTweener] will start interpolating. By default there's no delay.
*/
func (self Instance) SetDelay(delay Float.X) [1]gdclass.PropertyTweener {
	return [1]gdclass.PropertyTweener(class(self).SetDelay(gd.Float(delay)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PropertyTweener

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PropertyTweener"))
	casted := Instance{*(*gdclass.PropertyTweener)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Sets a custom initial value to the [PropertyTweener].
[b]Example:[/b]
[codeblock]
var tween = get_tree().create_tween()
tween.tween_property(self, "position", Vector2(200, 100), 1).from(Vector2(100, 100)) #this will move the node from position (100, 100) to (200, 100)
[/codeblock]
*/
//go:nosplit
func (self class) From(value gd.Variant) [1]gdclass.PropertyTweener {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(value))
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PropertyTweener.Bind_from, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.PropertyTweener{gd.PointerWithOwnershipTransferredToGo[gdclass.PropertyTweener](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Makes the [PropertyTweener] use the current property value (i.e. at the time of creating this [PropertyTweener]) as a starting point. This is equivalent of using [method from] with the current value. These two calls will do the same:
[codeblock]
tween.tween_property(self, "position", Vector2(200, 100), 1).from(position)
tween.tween_property(self, "position", Vector2(200, 100), 1).from_current()
[/codeblock]
*/
//go:nosplit
func (self class) FromCurrent() [1]gdclass.PropertyTweener {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PropertyTweener.Bind_from_current, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.PropertyTweener{gd.PointerWithOwnershipTransferredToGo[gdclass.PropertyTweener](r_ret.Get())}
	frame.Free()
	return ret
}

/*
When called, the final value will be used as a relative value instead.
[b]Example:[/b]
[codeblock]
var tween = get_tree().create_tween()
tween.tween_property(self, "position", Vector2.RIGHT * 100, 1).as_relative() #the node will move by 100 pixels to the right
[/codeblock]
*/
//go:nosplit
func (self class) AsRelative() [1]gdclass.PropertyTweener {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PropertyTweener.Bind_as_relative, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.PropertyTweener{gd.PointerWithOwnershipTransferredToGo[gdclass.PropertyTweener](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sets the type of used transition from [enum Tween.TransitionType]. If not set, the default transition is used from the [Tween] that contains this Tweener.
*/
//go:nosplit
func (self class) SetTrans(trans gdclass.TweenTransitionType) [1]gdclass.PropertyTweener {
	var frame = callframe.New()
	callframe.Arg(frame, trans)
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PropertyTweener.Bind_set_trans, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.PropertyTweener{gd.PointerWithOwnershipTransferredToGo[gdclass.PropertyTweener](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sets the type of used easing from [enum Tween.EaseType]. If not set, the default easing is used from the [Tween] that contains this Tweener.
*/
//go:nosplit
func (self class) SetEase(ease gdclass.TweenEaseType) [1]gdclass.PropertyTweener {
	var frame = callframe.New()
	callframe.Arg(frame, ease)
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PropertyTweener.Bind_set_ease, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.PropertyTweener{gd.PointerWithOwnershipTransferredToGo[gdclass.PropertyTweener](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Allows interpolating the value with a custom easing function. The provided [param interpolator_method] will be called with a value ranging from [code]0.0[/code] to [code]1.0[/code] and is expected to return a value within the same range (values outside the range can be used for overshoot). The return value of the method is then used for interpolation between initial and final value. Note that the parameter passed to the method is still subject to the tweener's own easing.
[b]Example:[/b]
[codeblock]
@export var curve: Curve

func _ready():
    var tween = create_tween()
    # Interpolate the value using a custom curve.
    tween.tween_property(self, "position:x", 300, 1).as_relative().set_custom_interpolator(tween_curve)

func tween_curve(v):
    return curve.sample_baked(v)
[/codeblock]
*/
//go:nosplit
func (self class) SetCustomInterpolator(interpolator_method gd.Callable) [1]gdclass.PropertyTweener {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(interpolator_method))
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PropertyTweener.Bind_set_custom_interpolator, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.PropertyTweener{gd.PointerWithOwnershipTransferredToGo[gdclass.PropertyTweener](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sets the time in seconds after which the [PropertyTweener] will start interpolating. By default there's no delay.
*/
//go:nosplit
func (self class) SetDelay(delay gd.Float) [1]gdclass.PropertyTweener {
	var frame = callframe.New()
	callframe.Arg(frame, delay)
	var r_ret = callframe.Ret[uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PropertyTweener.Bind_set_delay, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.PropertyTweener{gd.PointerWithOwnershipTransferredToGo[gdclass.PropertyTweener](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsPropertyTweener() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPropertyTweener() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsTweener() Tweener.Advanced    { return *((*Tweener.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTweener() Tweener.Instance {
	return *((*Tweener.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(Tweener.Advanced(self.AsTweener()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Tweener.Instance(self.AsTweener()), name)
	}
}
func init() {
	gdclass.Register("PropertyTweener", func(ptr gd.Object) any {
		return [1]gdclass.PropertyTweener{*(*gdclass.PropertyTweener)(unsafe.Pointer(&ptr))}
	})
}
