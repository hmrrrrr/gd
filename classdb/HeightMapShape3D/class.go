// Package HeightMapShape3D provides methods for working with HeightMapShape3D object instances.
package HeightMapShape3D

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Resource"
import "graphics.gd/classdb/Shape3D"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
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
A 3D heightmap shape, intended for use in physics. Usually used to provide a shape for a [CollisionShape3D]. This is useful for terrain, but it is limited as overhangs (such as caves) cannot be stored. Holes in a [HeightMapShape3D] are created by assigning very low values to points in the desired area.
[b]Performance:[/b] [HeightMapShape3D] is faster to check collisions against than [ConcavePolygonShape3D], but it is significantly slower than primitive shapes like [BoxShape3D].
A heightmap collision shape can also be build by using an [Image] reference:
[codeblocks]
[gdscript]
var heightmap_texture: Texture = ResourceLoader.load("res://heightmap_image.exr")
var heightmap_image: Image = heightmap_texture.get_image()
heightmap_image.convert(Image.FORMAT_RF)

var height_min: float = 0.0
var height_max: float = 10.0

update_map_data_from_image(heightmap_image, height_min, height_max)
[/gdscript]
[/codeblocks]
*/
type Instance [1]gdclass.HeightMapShape3D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsHeightMapShape3D() Instance
}

/*
Returns the smallest height value found in [member map_data]. Recalculates only when [member map_data] changes.
*/
func (self Instance) GetMinHeight() Float.X { //gd:HeightMapShape3D.get_min_height
	return Float.X(Float.X(class(self).GetMinHeight()))
}

/*
Returns the largest height value found in [member map_data]. Recalculates only when [member map_data] changes.
*/
func (self Instance) GetMaxHeight() Float.X { //gd:HeightMapShape3D.get_max_height
	return Float.X(Float.X(class(self).GetMaxHeight()))
}

/*
Updates [member map_data] with data read from an [Image] reference. Automatically resizes heightmap [member map_width] and [member map_depth] to fit the full image width and height.
The image needs to be in either [constant Image.FORMAT_RF] (32 bit), [constant Image.FORMAT_RH] (16 bit), or [constant Image.FORMAT_R8] (8 bit).
Each image pixel is read in as a float on the range from [code]0.0[/code] (black pixel) to [code]1.0[/code] (white pixel). This range value gets remapped to [param height_min] and [param height_max] to form the final height value.
*/
func (self Instance) UpdateMapDataFromImage(image [1]gdclass.Image, height_min Float.X, height_max Float.X) { //gd:HeightMapShape3D.update_map_data_from_image
	class(self).UpdateMapDataFromImage(image, float64(height_min), float64(height_max))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.HeightMapShape3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("HeightMapShape3D"))
	casted := Instance{*(*gdclass.HeightMapShape3D)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) MapWidth() int {
	return int(int(class(self).GetMapWidth()))
}

func (self Instance) SetMapWidth(value int) {
	class(self).SetMapWidth(int64(value))
}

func (self Instance) MapDepth() int {
	return int(int(class(self).GetMapDepth()))
}

func (self Instance) SetMapDepth(value int) {
	class(self).SetMapDepth(int64(value))
}

func (self Instance) MapData() []float32 {
	return []float32(slices.Collect(class(self).GetMapData().Values()))
}

func (self Instance) SetMapData(value []float32) {
	class(self).SetMapData(Packed.New(value...))
}

//go:nosplit
func (self class) SetMapWidth(width int64) { //gd:HeightMapShape3D.set_map_width
	var frame = callframe.New()
	callframe.Arg(frame, width)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HeightMapShape3D.Bind_set_map_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMapWidth() int64 { //gd:HeightMapShape3D.get_map_width
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HeightMapShape3D.Bind_get_map_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMapDepth(height int64) { //gd:HeightMapShape3D.set_map_depth
	var frame = callframe.New()
	callframe.Arg(frame, height)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HeightMapShape3D.Bind_set_map_depth, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMapDepth() int64 { //gd:HeightMapShape3D.get_map_depth
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HeightMapShape3D.Bind_get_map_depth, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMapData(data Packed.Array[float32]) { //gd:HeightMapShape3D.set_map_data
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedFloat32Array, float32](data)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HeightMapShape3D.Bind_set_map_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMapData() Packed.Array[float32] { //gd:HeightMapShape3D.get_map_data
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HeightMapShape3D.Bind_get_map_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[float32](Array.Through(gd.PackedProxy[gd.PackedFloat32Array, float32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the smallest height value found in [member map_data]. Recalculates only when [member map_data] changes.
*/
//go:nosplit
func (self class) GetMinHeight() float64 { //gd:HeightMapShape3D.get_min_height
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HeightMapShape3D.Bind_get_min_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the largest height value found in [member map_data]. Recalculates only when [member map_data] changes.
*/
//go:nosplit
func (self class) GetMaxHeight() float64 { //gd:HeightMapShape3D.get_max_height
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HeightMapShape3D.Bind_get_max_height, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Updates [member map_data] with data read from an [Image] reference. Automatically resizes heightmap [member map_width] and [member map_depth] to fit the full image width and height.
The image needs to be in either [constant Image.FORMAT_RF] (32 bit), [constant Image.FORMAT_RH] (16 bit), or [constant Image.FORMAT_R8] (8 bit).
Each image pixel is read in as a float on the range from [code]0.0[/code] (black pixel) to [code]1.0[/code] (white pixel). This range value gets remapped to [param height_min] and [param height_max] to form the final height value.
*/
//go:nosplit
func (self class) UpdateMapDataFromImage(image [1]gdclass.Image, height_min float64, height_max float64) { //gd:HeightMapShape3D.update_map_data_from_image
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(image[0])[0])
	callframe.Arg(frame, height_min)
	callframe.Arg(frame, height_max)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.HeightMapShape3D.Bind_update_map_data_from_image, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsHeightMapShape3D() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsHeightMapShape3D() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsShape3D() Shape3D.Advanced     { return *((*Shape3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsShape3D() Shape3D.Instance {
	return *((*Shape3D.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(Shape3D.Advanced(self.AsShape3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Shape3D.Instance(self.AsShape3D()), name)
	}
}
func init() {
	gdclass.Register("HeightMapShape3D", func(ptr gd.Object) any {
		return [1]gdclass.HeightMapShape3D{*(*gdclass.HeightMapShape3D)(unsafe.Pointer(&ptr))}
	})
}
