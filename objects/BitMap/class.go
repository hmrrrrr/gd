package BitMap

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Resource"
import "graphics.gd/variant/Vector2i"
import "graphics.gd/variant/Rect2i"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A two-dimensional array of boolean values, can be used to efficiently store a binary matrix (every matrix element takes only one bit) and query the values using natural cartesian coordinates.
*/
type Instance [1]classdb.BitMap
type Any interface {
	gd.IsClass
	AsBitMap() Instance
}

/*
Creates a bitmap with the specified size, filled with [code]false[/code].
*/
func (self Instance) Create(size Vector2i.XY) {
	class(self).Create(gd.Vector2i(size))
}

/*
Creates a bitmap that matches the given image dimensions, every element of the bitmap is set to [code]false[/code] if the alpha value of the image at that position is equal to [param threshold] or less, and [code]true[/code] in other case.
*/
func (self Instance) CreateFromImageAlpha(image objects.Image) {
	class(self).CreateFromImageAlpha(image, gd.Float(0.1))
}

/*
Sets the bitmap's element at the specified position, to the specified value.
*/
func (self Instance) SetBitv(position Vector2i.XY, bit bool) {
	class(self).SetBitv(gd.Vector2i(position), bit)
}

/*
Sets the bitmap's element at the specified position, to the specified value.
*/
func (self Instance) SetBit(x int, y int, bit bool) {
	class(self).SetBit(gd.Int(x), gd.Int(y), bit)
}

/*
Returns bitmap's value at the specified position.
*/
func (self Instance) GetBitv(position Vector2i.XY) bool {
	return bool(class(self).GetBitv(gd.Vector2i(position)))
}

/*
Returns bitmap's value at the specified position.
*/
func (self Instance) GetBit(x int, y int) bool {
	return bool(class(self).GetBit(gd.Int(x), gd.Int(y)))
}

/*
Sets a rectangular portion of the bitmap to the specified value.
*/
func (self Instance) SetBitRect(rect Rect2i.PositionSize, bit bool) {
	class(self).SetBitRect(gd.Rect2i(rect), bit)
}

/*
Returns the number of bitmap elements that are set to [code]true[/code].
*/
func (self Instance) GetTrueBitCount() int {
	return int(int(class(self).GetTrueBitCount()))
}

/*
Returns bitmap's dimensions.
*/
func (self Instance) GetSize() Vector2i.XY {
	return Vector2i.XY(class(self).GetSize())
}

/*
Resizes the image to [param new_size].
*/
func (self Instance) Resize(new_size Vector2i.XY) {
	class(self).Resize(gd.Vector2i(new_size))
}

/*
Applies morphological dilation or erosion to the bitmap. If [param pixels] is positive, dilation is applied to the bitmap. If [param pixels] is negative, erosion is applied to the bitmap. [param rect] defines the area where the morphological operation is applied. Pixels located outside the [param rect] are unaffected by [method grow_mask].
*/
func (self Instance) GrowMask(pixels int, rect Rect2i.PositionSize) {
	class(self).GrowMask(gd.Int(pixels), gd.Rect2i(rect))
}

/*
Returns an image of the same size as the bitmap and with a [enum Image.Format] of type [constant Image.FORMAT_L8]. [code]true[/code] bits of the bitmap are being converted into white pixels, and [code]false[/code] bits into black.
*/
func (self Instance) ConvertToImage() objects.Image {
	return objects.Image(class(self).ConvertToImage())
}

/*
Creates an [Array] of polygons covering a rectangular portion of the bitmap. It uses a marching squares algorithm, followed by Ramer-Douglas-Peucker (RDP) reduction of the number of vertices. Each polygon is described as a [PackedVector2Array] of its vertices.
To get polygons covering the whole bitmap, pass:
[codeblock]
Rect2(Vector2(), get_size())
[/codeblock]
[param epsilon] is passed to RDP to control how accurately the polygons cover the bitmap: a lower [param epsilon] corresponds to more points in the polygons.
*/
func (self Instance) OpaqueToPolygons(rect Rect2i.PositionSize) gd.Array {
	return gd.Array(class(self).OpaqueToPolygons(gd.Rect2i(rect), gd.Float(2.0)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.BitMap

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("BitMap"))
	return Instance{*(*classdb.BitMap)(unsafe.Pointer(&object))}
}

/*
Creates a bitmap with the specified size, filled with [code]false[/code].
*/
//go:nosplit
func (self class) Create(size gd.Vector2i) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BitMap.Bind_create, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Creates a bitmap that matches the given image dimensions, every element of the bitmap is set to [code]false[/code] if the alpha value of the image at that position is equal to [param threshold] or less, and [code]true[/code] in other case.
*/
//go:nosplit
func (self class) CreateFromImageAlpha(image objects.Image, threshold gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(image[0])[0])
	callframe.Arg(frame, threshold)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BitMap.Bind_create_from_image_alpha, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the bitmap's element at the specified position, to the specified value.
*/
//go:nosplit
func (self class) SetBitv(position gd.Vector2i, bit bool) {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, bit)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BitMap.Bind_set_bitv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the bitmap's element at the specified position, to the specified value.
*/
//go:nosplit
func (self class) SetBit(x gd.Int, y gd.Int, bit bool) {
	var frame = callframe.New()
	callframe.Arg(frame, x)
	callframe.Arg(frame, y)
	callframe.Arg(frame, bit)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BitMap.Bind_set_bit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns bitmap's value at the specified position.
*/
//go:nosplit
func (self class) GetBitv(position gd.Vector2i) bool {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BitMap.Bind_get_bitv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns bitmap's value at the specified position.
*/
//go:nosplit
func (self class) GetBit(x gd.Int, y gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, x)
	callframe.Arg(frame, y)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BitMap.Bind_get_bit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets a rectangular portion of the bitmap to the specified value.
*/
//go:nosplit
func (self class) SetBitRect(rect gd.Rect2i, bit bool) {
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	callframe.Arg(frame, bit)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BitMap.Bind_set_bit_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the number of bitmap elements that are set to [code]true[/code].
*/
//go:nosplit
func (self class) GetTrueBitCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BitMap.Bind_get_true_bit_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns bitmap's dimensions.
*/
//go:nosplit
func (self class) GetSize() gd.Vector2i {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BitMap.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Resizes the image to [param new_size].
*/
//go:nosplit
func (self class) Resize(new_size gd.Vector2i) {
	var frame = callframe.New()
	callframe.Arg(frame, new_size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BitMap.Bind_resize, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Applies morphological dilation or erosion to the bitmap. If [param pixels] is positive, dilation is applied to the bitmap. If [param pixels] is negative, erosion is applied to the bitmap. [param rect] defines the area where the morphological operation is applied. Pixels located outside the [param rect] are unaffected by [method grow_mask].
*/
//go:nosplit
func (self class) GrowMask(pixels gd.Int, rect gd.Rect2i) {
	var frame = callframe.New()
	callframe.Arg(frame, pixels)
	callframe.Arg(frame, rect)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BitMap.Bind_grow_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns an image of the same size as the bitmap and with a [enum Image.Format] of type [constant Image.FORMAT_L8]. [code]true[/code] bits of the bitmap are being converted into white pixels, and [code]false[/code] bits into black.
*/
//go:nosplit
func (self class) ConvertToImage() objects.Image {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BitMap.Bind_convert_to_image, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Image{gd.PointerWithOwnershipTransferredToGo[classdb.Image](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Creates an [Array] of polygons covering a rectangular portion of the bitmap. It uses a marching squares algorithm, followed by Ramer-Douglas-Peucker (RDP) reduction of the number of vertices. Each polygon is described as a [PackedVector2Array] of its vertices.
To get polygons covering the whole bitmap, pass:
[codeblock]
Rect2(Vector2(), get_size())
[/codeblock]
[param epsilon] is passed to RDP to control how accurately the polygons cover the bitmap: a lower [param epsilon] corresponds to more points in the polygons.
*/
//go:nosplit
func (self class) OpaqueToPolygons(rect gd.Rect2i, epsilon gd.Float) gd.Array {
	var frame = callframe.New()
	callframe.Arg(frame, rect)
	callframe.Arg(frame, epsilon)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BitMap.Bind_opaque_to_polygons, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsBitMap() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsBitMap() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
func init() {
	classdb.Register("BitMap", func(ptr gd.Object) any { return [1]classdb.BitMap{*(*classdb.BitMap)(unsafe.Pointer(&ptr))} })
}
