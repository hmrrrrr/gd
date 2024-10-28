package GLTFCamera

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Represents a camera as defined by the base GLTF spec.

*/
type Go [1]classdb.GLTFCamera

/*
Create a new GLTFCamera instance from the given Godot [Camera3D] node.
*/
func (self Go) FromNode(camera_node gdclass.Camera3D) gdclass.GLTFCamera {
	return gdclass.GLTFCamera(class(self).FromNode(camera_node))
}

/*
Converts this GLTFCamera instance into a Godot [Camera3D] node.
*/
func (self Go) ToNode() gdclass.Camera3D {
	return gdclass.Camera3D(class(self).ToNode())
}

/*
Creates a new GLTFCamera instance by parsing the given [Dictionary].
*/
func (self Go) FromDictionary(dictionary gd.Dictionary) gdclass.GLTFCamera {
	return gdclass.GLTFCamera(class(self).FromDictionary(dictionary))
}

/*
Serializes this GLTFCamera instance into a [Dictionary].
*/
func (self Go) ToDictionary() gd.Dictionary {
	return gd.Dictionary(class(self).ToDictionary())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.GLTFCamera
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GLTFCamera"))
	return Go{classdb.GLTFCamera(object)}
}

func (self Go) Perspective() bool {
		return bool(class(self).GetPerspective())
}

func (self Go) SetPerspective(value bool) {
	class(self).SetPerspective(value)
}

func (self Go) Fov() float64 {
		return float64(float64(class(self).GetFov()))
}

func (self Go) SetFov(value float64) {
	class(self).SetFov(gd.Float(value))
}

func (self Go) SizeMag() float64 {
		return float64(float64(class(self).GetSizeMag()))
}

func (self Go) SetSizeMag(value float64) {
	class(self).SetSizeMag(gd.Float(value))
}

func (self Go) DepthFar() float64 {
		return float64(float64(class(self).GetDepthFar()))
}

func (self Go) SetDepthFar(value float64) {
	class(self).SetDepthFar(gd.Float(value))
}

func (self Go) DepthNear() float64 {
		return float64(float64(class(self).GetDepthNear()))
}

func (self Go) SetDepthNear(value float64) {
	class(self).SetDepthNear(gd.Float(value))
}

/*
Create a new GLTFCamera instance from the given Godot [Camera3D] node.
*/
//go:nosplit
func (self class) FromNode(camera_node gdclass.Camera3D) gdclass.GLTFCamera {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(camera_node[0])[0])
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFCamera.Bind_from_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.GLTFCamera{classdb.GLTFCamera(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Converts this GLTFCamera instance into a Godot [Camera3D] node.
*/
//go:nosplit
func (self class) ToNode() gdclass.Camera3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFCamera.Bind_to_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Camera3D{classdb.Camera3D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Creates a new GLTFCamera instance by parsing the given [Dictionary].
*/
//go:nosplit
func (self class) FromDictionary(dictionary gd.Dictionary) gdclass.GLTFCamera {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(dictionary))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFCamera.Bind_from_dictionary, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.GLTFCamera{classdb.GLTFCamera(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Serializes this GLTFCamera instance into a [Dictionary].
*/
//go:nosplit
func (self class) ToDictionary() gd.Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFCamera.Bind_to_dictionary, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.Dictionary](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetPerspective() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFCamera.Bind_get_perspective, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPerspective(perspective bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, perspective)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFCamera.Bind_set_perspective, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFov() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFCamera.Bind_get_fov, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFov(fov gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, fov)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFCamera.Bind_set_fov, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSizeMag() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFCamera.Bind_get_size_mag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetSizeMag(size_mag gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, size_mag)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFCamera.Bind_set_size_mag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDepthFar() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFCamera.Bind_get_depth_far, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDepthFar(zdepth_far gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, zdepth_far)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFCamera.Bind_set_depth_far, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDepthNear() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFCamera.Bind_get_depth_near, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDepthNear(zdepth_near gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, zdepth_near)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GLTFCamera.Bind_set_depth_near, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsGLTFCamera() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsGLTFCamera() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("GLTFCamera", func(ptr gd.Object) any { return classdb.GLTFCamera(ptr) })}
