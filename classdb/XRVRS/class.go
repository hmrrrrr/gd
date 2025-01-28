// Package XRVRS provides methods for working with XRVRS object instances.
package XRVRS

import "unsafe"
import "reflect"
import "slices"
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
import "graphics.gd/variant/Path"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Vector2"

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
var _ = slices.Delete[[]struct{}, struct{}]

/*
This class is used by various XR interfaces to generate VRS textures that can be used to speed up rendering.
*/
type Instance [1]gdclass.XRVRS

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsXRVRS() Instance
}

/*
Generates the VRS texture based on a render [param target_size] adjusted by our VRS tile size. For each eyes focal point passed in [param eye_foci] a layer is created. Focal point should be in NDC.
The result will be cached, requesting a VRS texture with unchanged parameters and settings will return the cached RID.
*/
func (self Instance) MakeVrsTexture(target_size Vector2.XY, eye_foci []Vector2.XY) RID.Texture { //gd:XRVRS.make_vrs_texture
	return RID.Texture(class(self).MakeVrsTexture(gd.Vector2(target_size), Packed.New(eye_foci...)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.XRVRS

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("XRVRS"))
	casted := Instance{*(*gdclass.XRVRS)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) VrsMinRadius() Float.X {
	return Float.X(Float.X(class(self).GetVrsMinRadius()))
}

func (self Instance) SetVrsMinRadius(value Float.X) {
	class(self).SetVrsMinRadius(gd.Float(value))
}

func (self Instance) VrsStrength() Float.X {
	return Float.X(Float.X(class(self).GetVrsStrength()))
}

func (self Instance) SetVrsStrength(value Float.X) {
	class(self).SetVrsStrength(gd.Float(value))
}

//go:nosplit
func (self class) GetVrsMinRadius() gd.Float { //gd:XRVRS.get_vrs_min_radius
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRVRS.Bind_get_vrs_min_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVrsMinRadius(radius gd.Float) { //gd:XRVRS.set_vrs_min_radius
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRVRS.Bind_set_vrs_min_radius, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVrsStrength() gd.Float { //gd:XRVRS.get_vrs_strength
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRVRS.Bind_get_vrs_strength, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVrsStrength(strength gd.Float) { //gd:XRVRS.set_vrs_strength
	var frame = callframe.New()
	callframe.Arg(frame, strength)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRVRS.Bind_set_vrs_strength, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Generates the VRS texture based on a render [param target_size] adjusted by our VRS tile size. For each eyes focal point passed in [param eye_foci] a layer is created. Focal point should be in NDC.
The result will be cached, requesting a VRS texture with unchanged parameters and settings will return the cached RID.
*/
//go:nosplit
func (self class) MakeVrsTexture(target_size gd.Vector2, eye_foci Packed.Array[Vector2.XY]) gd.RID { //gd:XRVRS.make_vrs_texture
	var frame = callframe.New()
	callframe.Arg(frame, target_size)
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedVector2Array, Vector2.XY](eye_foci)))
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRVRS.Bind_make_vrs_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsXRVRS() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsXRVRS() Instance { return *((*Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Instance(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("XRVRS", func(ptr gd.Object) any { return [1]gdclass.XRVRS{*(*gdclass.XRVRS)(unsafe.Pointer(&ptr))} })
}
