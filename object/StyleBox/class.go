package StyleBox

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[StyleBox] is an abstract base class for drawing stylized boxes for UI elements. It is used for panels, buttons, [LineEdit] backgrounds, [Tree] backgrounds, etc. and also for testing a transparency mask for pointer signals. If mask test fails on a [StyleBox] assigned as mask to a control, clicks and motion signals will go through it to the one below.
[b]Note:[/b] For control nodes that have [i]Theme Properties[/i], the [code]focus[/code] [StyleBox] is displayed over the [code]normal[/code], [code]hover[/code] or [code]pressed[/code] [StyleBox]. This makes the [code]focus[/code] [StyleBox] more reusable across different nodes.
	// StyleBox methods that can be overridden by a [Class] that extends it.
	type StyleBox interface {
		Draw(to_canvas_item gd.RID, rect gd.Rect2) 
		GetDrawRect(rect gd.Rect2) gd.Rect2
		//Virtual method to be implemented by the user. Returns a custom minimum size that the stylebox must respect when drawing. By default [method get_minimum_size] only takes content margins into account. This method can be overridden to add another size restriction. A combination of the default behavior and the output of this method will be used, to account for both sizes.
		GetMinimumSize() gd.Vector2
		TestMask(point gd.Vector2, rect gd.Rect2) bool
	}

*/
type Simple [1]classdb.StyleBox
func (Simple) _draw(impl func(ptr unsafe.Pointer, to_canvas_item gd.RID, rect gd.Rect2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var to_canvas_item = gd.UnsafeGet[gd.RID](p_args,0)
		var rect = gd.UnsafeGet[gd.Rect2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, to_canvas_item, rect)
		gc.End()
	}
}
func (Simple) _get_draw_rect(impl func(ptr unsafe.Pointer, rect gd.Rect2) gd.Rect2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var rect = gd.UnsafeGet[gd.Rect2](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, rect)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _get_minimum_size(impl func(ptr unsafe.Pointer) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _test_mask(impl func(ptr unsafe.Pointer, point gd.Vector2, rect gd.Rect2) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var point = gd.UnsafeGet[gd.Vector2](p_args,0)
		var rect = gd.UnsafeGet[gd.Rect2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, point, rect)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (self Simple) GetMinimumSize() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetMinimumSize())
}
func (self Simple) SetContentMargin(margin gd.Side, offset float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetContentMargin(margin, gd.Float(offset))
}
func (self Simple) SetContentMarginAll(offset float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetContentMarginAll(gd.Float(offset))
}
func (self Simple) GetContentMargin(margin gd.Side) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetContentMargin(margin)))
}
func (self Simple) GetMargin(margin gd.Side) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetMargin(margin)))
}
func (self Simple) GetOffset() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetOffset())
}
func (self Simple) Draw(canvas_item gd.RID, rect gd.Rect2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Draw(canvas_item, rect)
}
func (self Simple) GetCurrentItemDrawn() [1]classdb.CanvasItem {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.CanvasItem(Expert(self).GetCurrentItemDrawn(gc))
}
func (self Simple) TestMask(point gd.Vector2, rect gd.Rect2) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).TestMask(point, rect))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.StyleBox
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

func (class) _draw(impl func(ptr unsafe.Pointer, to_canvas_item gd.RID, rect gd.Rect2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var to_canvas_item = gd.UnsafeGet[gd.RID](p_args,0)
		var rect = gd.UnsafeGet[gd.Rect2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, to_canvas_item, rect)
		ctx.End()
	}
}

func (class) _get_draw_rect(impl func(ptr unsafe.Pointer, rect gd.Rect2) gd.Rect2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var rect = gd.UnsafeGet[gd.Rect2](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, rect)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Virtual method to be implemented by the user. Returns a custom minimum size that the stylebox must respect when drawing. By default [method get_minimum_size] only takes content margins into account. This method can be overridden to add another size restriction. A combination of the default behavior and the output of this method will be used, to account for both sizes.
*/
func (class) _get_minimum_size(impl func(ptr unsafe.Pointer) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

func (class) _test_mask(impl func(ptr unsafe.Pointer, point gd.Vector2, rect gd.Rect2) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var point = gd.UnsafeGet[gd.Vector2](p_args,0)
		var rect = gd.UnsafeGet[gd.Rect2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, point, rect)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Returns the minimum size that this stylebox can be shrunk to.
*/
//go:nosplit
func (self class) GetMinimumSize() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBox.Bind_get_minimum_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the default value of the specified [enum Side] to [param offset] pixels.
*/
//go:nosplit
func (self class) SetContentMargin(margin gd.Side, offset gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBox.Bind_set_content_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the default margin to [param offset] pixels for all sides.
*/
//go:nosplit
func (self class) SetContentMarginAll(offset gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBox.Bind_set_content_margin_all, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the default margin of the specified [enum Side].
*/
//go:nosplit
func (self class) GetContentMargin(margin gd.Side) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBox.Bind_get_content_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the content margin offset for the specified [enum Side].
Positive values reduce size inwards, unlike [Control]'s margin values.
*/
//go:nosplit
func (self class) GetMargin(margin gd.Side) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBox.Bind_get_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the "offset" of a stylebox. This helper function returns a value equivalent to [code]Vector2(style.get_margin(MARGIN_LEFT), style.get_margin(MARGIN_TOP))[/code].
*/
//go:nosplit
func (self class) GetOffset() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBox.Bind_get_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Draws this stylebox using a canvas item identified by the given [RID].
The [RID] value can either be the result of [method CanvasItem.get_canvas_item] called on an existing [CanvasItem]-derived node, or directly from creating a canvas item in the [RenderingServer] with [method RenderingServer.canvas_item_create].
*/
//go:nosplit
func (self class) Draw(canvas_item gd.RID, rect gd.Rect2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, canvas_item)
	callframe.Arg(frame, rect)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBox.Bind_draw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [CanvasItem] that handles its [constant CanvasItem.NOTIFICATION_DRAW] or [method CanvasItem._draw] callback at this moment.
*/
//go:nosplit
func (self class) GetCurrentItemDrawn(ctx gd.Lifetime) object.CanvasItem {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBox.Bind_get_current_item_drawn, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.CanvasItem
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Test a position in a rectangle, return whether it passes the mask test.
*/
//go:nosplit
func (self class) TestMask(point gd.Vector2, rect gd.Rect2) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, point)
	callframe.Arg(frame, rect)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StyleBox.Bind_test_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsStyleBox() Expert { return self[0].AsStyleBox() }


//go:nosplit
func (self Simple) AsStyleBox() Simple { return self[0].AsStyleBox() }


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
	case "_draw": return reflect.ValueOf(self._draw);
	case "_get_draw_rect": return reflect.ValueOf(self._get_draw_rect);
	case "_get_minimum_size": return reflect.ValueOf(self._get_minimum_size);
	case "_test_mask": return reflect.ValueOf(self._test_mask);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	case "_draw": return reflect.ValueOf(self._draw);
	case "_get_draw_rect": return reflect.ValueOf(self._get_draw_rect);
	case "_get_minimum_size": return reflect.ValueOf(self._get_minimum_size);
	case "_test_mask": return reflect.ValueOf(self._test_mask);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("StyleBox", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
