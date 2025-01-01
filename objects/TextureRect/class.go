package TextureRect

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Control"
import "graphics.gd/objects/CanvasItem"
import "graphics.gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A control that displays a texture, for example an icon inside a GUI. The texture's placement can be controlled with the [member stretch_mode] property. It can scale, tile, or stay centered inside its bounding rectangle.
*/
type Instance [1]classdb.TextureRect
type Any interface {
	gd.IsClass
	AsTextureRect() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.TextureRect

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TextureRect"))
	return Instance{classdb.TextureRect(object)}
}

func (self Instance) Texture() objects.Texture2D {
	return objects.Texture2D(class(self).GetTexture())
}

func (self Instance) SetTexture(value objects.Texture2D) {
	class(self).SetTexture(value)
}

func (self Instance) ExpandMode() classdb.TextureRectExpandMode {
	return classdb.TextureRectExpandMode(class(self).GetExpandMode())
}

func (self Instance) SetExpandMode(value classdb.TextureRectExpandMode) {
	class(self).SetExpandMode(value)
}

func (self Instance) StretchMode() classdb.TextureRectStretchMode {
	return classdb.TextureRectStretchMode(class(self).GetStretchMode())
}

func (self Instance) SetStretchMode(value classdb.TextureRectStretchMode) {
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
func (self class) SetTexture(texture objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureRect.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTexture() objects.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureRect.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetExpandMode(expand_mode classdb.TextureRectExpandMode) {
	var frame = callframe.New()
	callframe.Arg(frame, expand_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureRect.Bind_set_expand_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetExpandMode() classdb.TextureRectExpandMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextureRectExpandMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureRect.Bind_get_expand_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFlipH(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureRect.Bind_set_flip_h, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsFlippedH() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureRect.Bind_is_flipped_h, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFlipV(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureRect.Bind_set_flip_v, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsFlippedV() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureRect.Bind_is_flipped_v, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStretchMode(stretch_mode classdb.TextureRectStretchMode) {
	var frame = callframe.New()
	callframe.Arg(frame, stretch_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureRect.Bind_set_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetStretchMode() classdb.TextureRectStretchMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.TextureRectStretchMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TextureRect.Bind_get_stretch_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsTextureRect() Advanced     { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTextureRect() Instance  { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(self.AsControl(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsControl(), name)
	}
}
func init() {
	classdb.Register("TextureRect", func(ptr gd.Object) any { return [1]classdb.TextureRect{classdb.TextureRect(ptr)} })
}

type ExpandMode = classdb.TextureRectExpandMode

const (
	/*The minimum size will be equal to texture size, i.e. [TextureRect] can't be smaller than the texture.*/
	ExpandKeepSize ExpandMode = 0
	/*The size of the texture won't be considered for minimum size calculation, so the [TextureRect] can be shrunk down past the texture size.*/
	ExpandIgnoreSize ExpandMode = 1
	/*The height of the texture will be ignored. Minimum width will be equal to the current height. Useful for horizontal layouts, e.g. inside [HBoxContainer].*/
	ExpandFitWidth ExpandMode = 2
	/*Same as [constant EXPAND_FIT_WIDTH], but keeps texture's aspect ratio.*/
	ExpandFitWidthProportional ExpandMode = 3
	/*The width of the texture will be ignored. Minimum height will be equal to the current width. Useful for vertical layouts, e.g. inside [VBoxContainer].*/
	ExpandFitHeight ExpandMode = 4
	/*Same as [constant EXPAND_FIT_HEIGHT], but keeps texture's aspect ratio.*/
	ExpandFitHeightProportional ExpandMode = 5
)

type StretchMode = classdb.TextureRectStretchMode

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
	/*Scale the texture to fit the node's bounding rectangle, center it and maintain its aspect ratio.*/
	StretchKeepAspectCentered StretchMode = 5
	/*Scale the texture so that the shorter side fits the bounding rectangle. The other side clips to the node's limits.*/
	StretchKeepAspectCovered StretchMode = 6
)
