// Package TextureButton provides methods for working with TextureButton object instances.
package TextureButton

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/BaseButton"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
[TextureButton] has the same functionality as [Button], except it uses sprites instead of Godot's [Theme] resource. It is faster to create, but it doesn't support localization like more complex [Control]s.
The "normal" state must contain a texture ([member texture_normal]); other textures are optional.
See also [BaseButton] which contains common properties and methods associated with this node.
*/
type Instance [1]gdclass.TextureButton
type Any interface {
	gd.IsClass
	AsTextureButton() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.TextureButton

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TextureButton"))
	casted := Instance{*(*gdclass.TextureButton)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) TextureNormal() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTextureNormal())
}

func (self Instance) SetTextureNormal(value [1]gdclass.Texture2D) {
	class(self).SetTextureNormal(value)
}

func (self Instance) TexturePressed() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTexturePressed())
}

func (self Instance) SetTexturePressed(value [1]gdclass.Texture2D) {
	class(self).SetTexturePressed(value)
}

func (self Instance) TextureHover() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTextureHover())
}

func (self Instance) SetTextureHover(value [1]gdclass.Texture2D) {
	class(self).SetTextureHover(value)
}

func (self Instance) TextureDisabled() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTextureDisabled())
}

func (self Instance) SetTextureDisabled(value [1]gdclass.Texture2D) {
	class(self).SetTextureDisabled(value)
}

func (self Instance) TextureFocused() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetTextureFocused())
}

func (self Instance) SetTextureFocused(value [1]gdclass.Texture2D) {
	class(self).SetTextureFocused(value)
}

func (self Instance) TextureClickMask() [1]gdclass.BitMap {
	return [1]gdclass.BitMap(class(self).GetClickMask())
}

func (self Instance) SetTextureClickMask(value [1]gdclass.BitMap) {
	class(self).SetClickMask(value)
}

func (self Instance) IgnoreTextureSize() bool {
	return bool(class(self).GetIgnoreTextureSize())
}

func (self Instance) SetIgnoreTextureSize(value bool) {
	class(self).SetIgnoreTextureSize(value)
}

func (self Instance) StretchMode() gdclass.TextureButtonStretchMode {
	return gdclass.TextureButtonStretchMode(class(self).GetStretchMode())
}

func (self Instance) SetStretchMode(value gdclass.TextureButtonStretchMode) {
	class(self).SetStretchMode(value)
}

func (self Instance) FlipH() bool {
	return bool(class(self).IsFlippedH())
}

func (self Instance) SetFlipH(value bool) {
	class(self).SetFlipH(value)
}

func (self Instance) FlipV() bool {
	return bool(class(self).IsFlippedV())
}

func (self Instance) SetFlipV(value bool) {
	class(self).SetFlipV(value)
}

//go:nosplit
func (self class) SetTextureNormal(texture [1]gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_texture_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetTexturePressed(texture [1]gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_texture_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetTextureHover(texture [1]gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_texture_hover, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetTextureDisabled(texture [1]gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_texture_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetTextureFocused(texture [1]gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_texture_focused, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetClickMask(mask [1]gdclass.BitMap) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(mask[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_click_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetIgnoreTextureSize(ignore bool) {
	var frame = callframe.New()
	callframe.Arg(frame, ignore)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_ignore_texture_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetStretchMode(mode gdclass.TextureButtonStretchMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetFlipH(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_flip_h, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsFlippedH() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_is_flipped_h, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFlipV(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_set_flip_v, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsFlippedV() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_is_flipped_v, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetTextureNormal() [1]gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_get_texture_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetTexturePressed() [1]gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_get_texture_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetTextureHover() [1]gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_get_texture_hover, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetTextureDisabled() [1]gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_get_texture_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetTextureFocused() [1]gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_get_texture_focused, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetClickMask() [1]gdclass.BitMap {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_get_click_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.BitMap{gd.PointerWithOwnershipTransferredToGo[gdclass.BitMap](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetIgnoreTextureSize() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_get_ignore_texture_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) GetStretchMode() gdclass.TextureButtonStretchMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextureButtonStretchMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureButton.Bind_get_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsTextureButton() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTextureButton() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsBaseButton() BaseButton.Advanced {
	return *((*BaseButton.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsBaseButton() BaseButton.Instance {
	return *((*BaseButton.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsControl() Control.Advanced { return *((*Control.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsControl() Control.Instance {
	return *((*Control.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(BaseButton.Advanced(self.AsBaseButton()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(BaseButton.Instance(self.AsBaseButton()), name)
	}
}
func init() {
	gdclass.Register("TextureButton", func(ptr gd.Object) any {
		return [1]gdclass.TextureButton{*(*gdclass.TextureButton)(unsafe.Pointer(&ptr))}
	})
}

type StretchMode = gdclass.TextureButtonStretchMode

const (
	/*Scale to fit the node's bounding rectangle.*/
	StretchScale StretchMode = 0
	/*Tile inside the node's bounding rectangle.*/
	StretchTile StretchMode = 1
	/*The texture keeps its original size and stays in the bounding rectangle's top-left corner.*/
	StretchKeep StretchMode = 2
	/*The texture keeps its original size and stays centered in the node's bounding rectangle.*/
	StretchKeepCentered StretchMode = 3
	/*Scale the texture to fit the node's bounding rectangle, but maintain the texture's aspect ratio.*/
	StretchKeepAspect StretchMode = 4
	/*Scale the texture to fit the node's bounding rectangle, center it, and maintain its aspect ratio.*/
	StretchKeepAspectCentered StretchMode = 5
	/*Scale the texture so that the shorter side fits the bounding rectangle. The other side clips to the node's limits.*/
	StretchKeepAspectCovered StretchMode = 6
)
