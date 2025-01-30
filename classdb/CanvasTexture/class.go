// Package CanvasTexture provides methods for working with CanvasTexture object instances.
package CanvasTexture

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Resource"
import "graphics.gd/classdb/Texture"
import "graphics.gd/classdb/Texture2D"
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
[CanvasTexture] is an alternative to [ImageTexture] for 2D rendering. It allows using normal maps and specular maps in any node that inherits from [CanvasItem]. [CanvasTexture] also allows overriding the texture's filter and repeat mode independently of the node's properties (or the project settings).
[b]Note:[/b] [CanvasTexture] cannot be used in 3D. It will not display correctly when applied to any [VisualInstance3D], such as [Sprite3D] or [Decal]. For physically-based materials in 3D, use [BaseMaterial3D] instead.
*/
type Instance [1]gdclass.CanvasTexture

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCanvasTexture() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.CanvasTexture

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CanvasTexture"))
	casted := Instance{*(*gdclass.CanvasTexture)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) DiffuseTexture() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetDiffuseTexture())
}

func (self Instance) SetDiffuseTexture(value [1]gdclass.Texture2D) {
	class(self).SetDiffuseTexture(value)
}

func (self Instance) NormalTexture() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetNormalTexture())
}

func (self Instance) SetNormalTexture(value [1]gdclass.Texture2D) {
	class(self).SetNormalTexture(value)
}

func (self Instance) SpecularTexture() [1]gdclass.Texture2D {
	return [1]gdclass.Texture2D(class(self).GetSpecularTexture())
}

func (self Instance) SetSpecularTexture(value [1]gdclass.Texture2D) {
	class(self).SetSpecularTexture(value)
}

func (self Instance) SpecularColor() Color.RGBA {
	return Color.RGBA(class(self).GetSpecularColor())
}

func (self Instance) SetSpecularColor(value Color.RGBA) {
	class(self).SetSpecularColor(Color.RGBA(value))
}

func (self Instance) SpecularShininess() Float.X {
	return Float.X(Float.X(class(self).GetSpecularShininess()))
}

func (self Instance) SetSpecularShininess(value Float.X) {
	class(self).SetSpecularShininess(float64(value))
}

func (self Instance) TextureFilter() gdclass.CanvasItemTextureFilter {
	return gdclass.CanvasItemTextureFilter(class(self).GetTextureFilter())
}

func (self Instance) SetTextureFilter(value gdclass.CanvasItemTextureFilter) {
	class(self).SetTextureFilter(value)
}

func (self Instance) TextureRepeat() gdclass.CanvasItemTextureRepeat {
	return gdclass.CanvasItemTextureRepeat(class(self).GetTextureRepeat())
}

func (self Instance) SetTextureRepeat(value gdclass.CanvasItemTextureRepeat) {
	class(self).SetTextureRepeat(value)
}

//go:nosplit
func (self class) SetDiffuseTexture(texture [1]gdclass.Texture2D) { //gd:CanvasTexture.set_diffuse_texture
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasTexture.Bind_set_diffuse_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDiffuseTexture() [1]gdclass.Texture2D { //gd:CanvasTexture.get_diffuse_texture
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasTexture.Bind_get_diffuse_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetNormalTexture(texture [1]gdclass.Texture2D) { //gd:CanvasTexture.set_normal_texture
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasTexture.Bind_set_normal_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetNormalTexture() [1]gdclass.Texture2D { //gd:CanvasTexture.get_normal_texture
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasTexture.Bind_get_normal_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSpecularTexture(texture [1]gdclass.Texture2D) { //gd:CanvasTexture.set_specular_texture
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasTexture.Bind_set_specular_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSpecularTexture() [1]gdclass.Texture2D { //gd:CanvasTexture.get_specular_texture
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasTexture.Bind_get_specular_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSpecularColor(color Color.RGBA) { //gd:CanvasTexture.set_specular_color
	var frame = callframe.New()
	callframe.Arg(frame, color)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasTexture.Bind_set_specular_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSpecularColor() Color.RGBA { //gd:CanvasTexture.get_specular_color
	var frame = callframe.New()
	var r_ret = callframe.Ret[Color.RGBA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasTexture.Bind_get_specular_color, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSpecularShininess(shininess float64) { //gd:CanvasTexture.set_specular_shininess
	var frame = callframe.New()
	callframe.Arg(frame, shininess)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasTexture.Bind_set_specular_shininess, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSpecularShininess() float64 { //gd:CanvasTexture.get_specular_shininess
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasTexture.Bind_get_specular_shininess, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureFilter(filter gdclass.CanvasItemTextureFilter) { //gd:CanvasTexture.set_texture_filter
	var frame = callframe.New()
	callframe.Arg(frame, filter)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasTexture.Bind_set_texture_filter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureFilter() gdclass.CanvasItemTextureFilter { //gd:CanvasTexture.get_texture_filter
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.CanvasItemTextureFilter](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasTexture.Bind_get_texture_filter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureRepeat(repeat gdclass.CanvasItemTextureRepeat) { //gd:CanvasTexture.set_texture_repeat
	var frame = callframe.New()
	callframe.Arg(frame, repeat)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasTexture.Bind_set_texture_repeat, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureRepeat() gdclass.CanvasItemTextureRepeat { //gd:CanvasTexture.get_texture_repeat
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.CanvasItemTextureRepeat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.CanvasTexture.Bind_get_texture_repeat, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsCanvasTexture() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCanvasTexture() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Texture2D.Advanced(self.AsTexture2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Texture2D.Instance(self.AsTexture2D()), name)
	}
}
func init() {
	gdclass.Register("CanvasTexture", func(ptr gd.Object) any {
		return [1]gdclass.CanvasTexture{*(*gdclass.CanvasTexture)(unsafe.Pointer(&ptr))}
	})
}
