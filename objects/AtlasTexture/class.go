package AtlasTexture

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Texture2D"
import "graphics.gd/objects/Texture"
import "graphics.gd/objects/Resource"
import "graphics.gd/variant/Rect2"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
[Texture2D] resource that draws only part of its [member atlas] texture, as defined by the [member region]. An additional [member margin] can also be set, which is useful for small adjustments.
Multiple [AtlasTexture] resources can be cropped from the same [member atlas]. Packing many smaller textures into a singular large texture helps to optimize video memory costs and render calls.
[b]Note:[/b] [AtlasTexture] cannot be used in an [AnimatedTexture], and may not tile properly in nodes such as [TextureRect], when inside other [AtlasTexture] resources.
*/
type Instance [1]classdb.AtlasTexture
type Any interface {
	gd.IsClass
	AsAtlasTexture() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.AtlasTexture

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AtlasTexture"))
	return Instance{*(*classdb.AtlasTexture)(unsafe.Pointer(&object))}
}

func (self Instance) Atlas() objects.Texture2D {
	return objects.Texture2D(class(self).GetAtlas())
}

func (self Instance) SetAtlas(value objects.Texture2D) {
	class(self).SetAtlas(value)
}

func (self Instance) Region() Rect2.PositionSize {
	return Rect2.PositionSize(class(self).GetRegion())
}

func (self Instance) SetRegion(value Rect2.PositionSize) {
	class(self).SetRegion(gd.Rect2(value))
}

func (self Instance) Margin() Rect2.PositionSize {
	return Rect2.PositionSize(class(self).GetMargin())
}

func (self Instance) SetMargin(value Rect2.PositionSize) {
	class(self).SetMargin(gd.Rect2(value))
}

func (self Instance) FilterClip() bool {
	return bool(class(self).HasFilterClip())
}

func (self Instance) SetFilterClip(value bool) {
	class(self).SetFilterClip(value)
}

//go:nosplit
func (self class) SetAtlas(atlas objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(atlas[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AtlasTexture.Bind_set_atlas, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetAtlas() objects.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AtlasTexture.Bind_get_atlas, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Texture2D{gd.PointerWithOwnershipTransferredToGo[classdb.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRegion(region gd.Rect2) {
	var frame = callframe.New()
	callframe.Arg(frame, region)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AtlasTexture.Bind_set_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetRegion() gd.Rect2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AtlasTexture.Bind_get_region, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMargin(margin gd.Rect2) {
	var frame = callframe.New()
	callframe.Arg(frame, margin)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AtlasTexture.Bind_set_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMargin() gd.Rect2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AtlasTexture.Bind_get_margin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFilterClip(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AtlasTexture.Bind_set_filter_clip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) HasFilterClip() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AtlasTexture.Bind_has_filter_clip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAtlasTexture() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAtlasTexture() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsTexture2D() Texture2D.Advanced {
	return *((*Texture2D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsTexture2D() Texture2D.Instance {
	return *((*Texture2D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsTexture() Texture.Advanced { return *((*Texture.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTexture() Texture.Instance {
	return *((*Texture.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(self.AsTexture2D(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsTexture2D(), name)
	}
}
func init() {
	classdb.Register("AtlasTexture", func(ptr gd.Object) any {
		return [1]classdb.AtlasTexture{*(*classdb.AtlasTexture)(unsafe.Pointer(&ptr))}
	})
}
