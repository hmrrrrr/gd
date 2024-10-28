package XRVRS

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
This class is used by various XR interfaces to generate VRS textures that can be used to speed up rendering.

*/
type Go [1]classdb.XRVRS

/*
Generates the VRS texture based on a render [param target_size] adjusted by our VRS tile size. For each eyes focal point passed in [param eye_foci] a layer is created. Focal point should be in NDC.
The result will be cached, requesting a VRS texture with unchanged parameters and settings will return the cached RID.
*/
func (self Go) MakeVrsTexture(target_size gd.Vector2, eye_foci []gd.Vector2) gd.RID {
	return gd.RID(class(self).MakeVrsTexture(target_size, gd.NewPackedVector2Slice(eye_foci)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.XRVRS
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("XRVRS"))
	return Go{classdb.XRVRS(object)}
}

func (self Go) VrsMinRadius() float64 {
		return float64(float64(class(self).GetVrsMinRadius()))
}

func (self Go) SetVrsMinRadius(value float64) {
	class(self).SetVrsMinRadius(gd.Float(value))
}

func (self Go) VrsStrength() float64 {
		return float64(float64(class(self).GetVrsStrength()))
}

func (self Go) SetVrsStrength(value float64) {
	class(self).SetVrsStrength(gd.Float(value))
}

//go:nosplit
func (self class) GetVrsMinRadius() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRVRS.Bind_get_vrs_min_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVrsMinRadius(radius gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRVRS.Bind_set_vrs_min_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVrsStrength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRVRS.Bind_get_vrs_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVrsStrength(strength gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, strength)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRVRS.Bind_set_vrs_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Generates the VRS texture based on a render [param target_size] adjusted by our VRS tile size. For each eyes focal point passed in [param eye_foci] a layer is created. Focal point should be in NDC.
The result will be cached, requesting a VRS texture with unchanged parameters and settings will return the cached RID.
*/
//go:nosplit
func (self class) MakeVrsTexture(target_size gd.Vector2, eye_foci gd.PackedVector2Array) gd.RID {
	var frame = callframe.New()
	callframe.Arg(frame, target_size)
	callframe.Arg(frame, discreet.Get(eye_foci))
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.XRVRS.Bind_make_vrs_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsXRVRS() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsXRVRS() Go { return *((*Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {classdb.Register("XRVRS", func(ptr gd.Object) any { return classdb.XRVRS(ptr) })}
