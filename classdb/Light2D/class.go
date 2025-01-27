// Package Light2D provides methods for working with Light2D object instances.
package Light2D

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
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/classdb/Node2D"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/variant/Color"
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
var _ RID.Any
var _ String.Readable

/*
Casts light in a 2D environment. A light is defined as a color, an energy value, a mode (see constants), and various other parameters (range and shadows-related).
*/
type Instance [1]gdclass.Light2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsLight2D() Instance
}

/*
Sets the light's height, which is used in 2D normal mapping. See [member PointLight2D.height] and [member DirectionalLight2D.height].
*/
func (self Instance) SetHeight(height Float.X) { //gd:Light2D.set_height
	class(self).SetHeight(gd.Float(height))
}

/*
Returns the light's height, which is used in 2D normal mapping. See [member PointLight2D.height] and [member DirectionalLight2D.height].
*/
func (self Instance) GetHeight() Float.X { //gd:Light2D.get_height
	return Float.X(Float.X(class(self).GetHeight()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Light2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Light2D"))
	casted := Instance{*(*gdclass.Light2D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Enabled() bool {
	return bool(class(self).IsEnabled())
}

func (self Instance) SetEnabled(value bool) {
	class(self).SetEnabled(value)
}

func (self Instance) EditorOnly() bool {
	return bool(class(self).IsEditorOnly())
}

func (self Instance) SetEditorOnly(value bool) {
	class(self).SetEditorOnly(value)
}

func (self Instance) Color() Color.RGBA {
	return Color.RGBA(class(self).GetColor())
}

func (self Instance) SetColor(value Color.RGBA) {
	class(self).SetColor(gd.Color(value))
}

func (self Instance) Energy() Float.X {
	return Float.X(Float.X(class(self).GetEnergy()))
}

func (self Instance) SetEnergy(value Float.X) {
	class(self).SetEnergy(gd.Float(value))
}

func (self Instance) BlendMode() gdclass.Light2DBlendMode {
	return gdclass.Light2DBlendMode(class(self).GetBlendMode())
}

func (self Instance) SetBlendMode(value gdclass.Light2DBlendMode) {
	class(self).SetBlendMode(value)
}

func (self Instance) RangeZMin() int {
	return int(int(class(self).GetZRangeMin()))
}

func (self Instance) SetRangeZMin(value int) {
	class(self).SetZRangeMin(gd.Int(value))
}

func (self Instance) RangeZMax() int {
	return int(int(class(self).GetZRangeMax()))
}

func (self Instance) SetRangeZMax(value int) {
	class(self).SetZRangeMax(gd.Int(value))
}

func (self Instance) RangeLayerMin() int {
	return int(int(class(self).GetLayerRangeMin()))
}

func (self Instance) SetRangeLayerMin(value int) {
	class(self).SetLayerRangeMin(gd.Int(value))
}

func (self Instance) RangeLayerMax() int {
	return int(int(class(self).GetLayerRangeMax()))
}

func (self Instance) SetRangeLayerMax(value int) {
	class(self).SetLayerRangeMax(gd.Int(value))
}

func (self Instance) RangeItemCullMask() int {
	return int(int(class(self).GetItemCullMask()))
}

func (self Instance) SetRangeItemCullMask(value int) {
	class(self).SetItemCullMask(gd.Int(value))
}

func (self Instance) ShadowEnabled() bool {
	return bool(class(self).IsShadowEnabled())
}

func (self Instance) SetShadowEnabled(value bool) {
	class(self).SetShadowEnabled(value)
}

func (self Instance) ShadowColor() Color.RGBA {
	return Color.RGBA(class(self).GetShadowColor())
}

func (self Instance) SetShadowColor(value Color.RGBA) {
	class(self).SetShadowColor(gd.Color(value))
}

func (self Instance) ShadowFilter() gdclass.Light2DShadowFilter {
	return gdclass.Light2DShadowFilter(class(self).GetShadowFilter())
}

func (self Instance) SetShadowFilter(value gdclass.Light2DShadowFilter) {
	class(self).SetShadowFilter(value)
}

func (self Instance) ShadowFilterSmooth() Float.X {
	return Float.X(Float.X(class(self).GetShadowSmooth()))
}

func (self Instance) SetShadowFilterSmooth(value Float.X) {
	class(self).SetShadowSmooth(gd.Float(value))
}

func (self Instance) ShadowItemCullMask() int {
	return int(int(class(self).GetItemShadowCullMask()))
}

func (self Instance) SetShadowItemCullMask(value int) {
	class(self).SetItemShadowCullMask(gd.Int(value))
}

//go:nosplit
func (self class) SetEnabled(enabled bool) { //gd:Light2D.set_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_set_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsEnabled() bool { //gd:Light2D.is_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_is_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEditorOnly(editor_only bool) { //gd:Light2D.set_editor_only
	var frame = callframe.New()
	callframe.Arg(frame, editor_only)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_set_editor_only, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsEditorOnly() bool { //gd:Light2D.is_editor_only
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_is_editor_only, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetColor(color gd.Color) { //gd:Light2D.set_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_set_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetColor() gd.Color { //gd:Light2D.get_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_get_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnergy(energy gd.Float) { //gd:Light2D.set_energy
	var frame = callframe.New()
	callframe.Arg(frame, energy)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_set_energy, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnergy() gd.Float { //gd:Light2D.get_energy
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_get_energy, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetZRangeMin(z gd.Int) { //gd:Light2D.set_z_range_min
	var frame = callframe.New()
	callframe.Arg(frame, z)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_set_z_range_min, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetZRangeMin() gd.Int { //gd:Light2D.get_z_range_min
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_get_z_range_min, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetZRangeMax(z gd.Int) { //gd:Light2D.set_z_range_max
	var frame = callframe.New()
	callframe.Arg(frame, z)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_set_z_range_max, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetZRangeMax() gd.Int { //gd:Light2D.get_z_range_max
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_get_z_range_max, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLayerRangeMin(layer gd.Int) { //gd:Light2D.set_layer_range_min
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_set_layer_range_min, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLayerRangeMin() gd.Int { //gd:Light2D.get_layer_range_min
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_get_layer_range_min, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLayerRangeMax(layer gd.Int) { //gd:Light2D.set_layer_range_max
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_set_layer_range_max, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLayerRangeMax() gd.Int { //gd:Light2D.get_layer_range_max
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_get_layer_range_max, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetItemCullMask(item_cull_mask gd.Int) { //gd:Light2D.set_item_cull_mask
	var frame = callframe.New()
	callframe.Arg(frame, item_cull_mask)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_set_item_cull_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetItemCullMask() gd.Int { //gd:Light2D.get_item_cull_mask
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_get_item_cull_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetItemShadowCullMask(item_shadow_cull_mask gd.Int) { //gd:Light2D.set_item_shadow_cull_mask
	var frame = callframe.New()
	callframe.Arg(frame, item_shadow_cull_mask)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_set_item_shadow_cull_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetItemShadowCullMask() gd.Int { //gd:Light2D.get_item_shadow_cull_mask
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_get_item_shadow_cull_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShadowEnabled(enabled bool) { //gd:Light2D.set_shadow_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_set_shadow_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsShadowEnabled() bool { //gd:Light2D.is_shadow_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_is_shadow_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShadowSmooth(smooth gd.Float) { //gd:Light2D.set_shadow_smooth
	var frame = callframe.New()
	callframe.Arg(frame, smooth)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_set_shadow_smooth, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetShadowSmooth() gd.Float { //gd:Light2D.get_shadow_smooth
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_get_shadow_smooth, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShadowFilter(filter gdclass.Light2DShadowFilter) { //gd:Light2D.set_shadow_filter
	var frame = callframe.New()
	callframe.Arg(frame, filter)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_set_shadow_filter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetShadowFilter() gdclass.Light2DShadowFilter { //gd:Light2D.get_shadow_filter
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.Light2DShadowFilter](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_get_shadow_filter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShadowColor(shadow_color gd.Color) { //gd:Light2D.set_shadow_color
	var frame = callframe.New()
	callframe.Arg(frame, shadow_color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_set_shadow_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetShadowColor() gd.Color { //gd:Light2D.get_shadow_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_get_shadow_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBlendMode(mode gdclass.Light2DBlendMode) { //gd:Light2D.set_blend_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_set_blend_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBlendMode() gdclass.Light2DBlendMode { //gd:Light2D.get_blend_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.Light2DBlendMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_get_blend_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the light's height, which is used in 2D normal mapping. See [member PointLight2D.height] and [member DirectionalLight2D.height].
*/
//go:nosplit
func (self class) SetHeight(height gd.Float) { //gd:Light2D.set_height
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_set_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the light's height, which is used in 2D normal mapping. See [member PointLight2D.height] and [member DirectionalLight2D.height].
*/
//go:nosplit
func (self class) GetHeight() gd.Float { //gd:Light2D.get_height
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Light2D.Bind_get_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsLight2D() Advanced          { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsLight2D() Instance       { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode2D() Node2D.Advanced    { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode2D() Node2D.Instance { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node2D.Advanced(self.AsNode2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node2D.Instance(self.AsNode2D()), name)
	}
}
func init() {
	gdclass.Register("Light2D", func(ptr gd.Object) any { return [1]gdclass.Light2D{*(*gdclass.Light2D)(unsafe.Pointer(&ptr))} })
}

type ShadowFilter = gdclass.Light2DShadowFilter //gd:Light2D.ShadowFilter

const (
	/*No filter applies to the shadow map. This provides hard shadow edges and is the fastest to render. See [member shadow_filter].*/
	ShadowFilterNone ShadowFilter = 0
	/*Percentage closer filtering (5 samples) applies to the shadow map. This is slower compared to hard shadow rendering. See [member shadow_filter].*/
	ShadowFilterPcf5 ShadowFilter = 1
	/*Percentage closer filtering (13 samples) applies to the shadow map. This is the slowest shadow filtering mode, and should be used sparingly. See [member shadow_filter].*/
	ShadowFilterPcf13 ShadowFilter = 2
)

type BlendMode = gdclass.Light2DBlendMode //gd:Light2D.BlendMode

const (
	/*Adds the value of pixels corresponding to the Light2D to the values of pixels under it. This is the common behavior of a light.*/
	BlendModeAdd BlendMode = 0
	/*Subtracts the value of pixels corresponding to the Light2D to the values of pixels under it, resulting in inversed light effect.*/
	BlendModeSub BlendMode = 1
	/*Mix the value of pixels corresponding to the Light2D to the values of pixels under it by linear interpolation.*/
	BlendModeMix BlendMode = 2
)
