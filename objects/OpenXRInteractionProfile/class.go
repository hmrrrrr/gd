package OpenXRInteractionProfile

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Resource"
import "graphics.gd/variant/Array"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
This object stores suggested bindings for an interaction profile. Interaction profiles define the metadata for a tracked XR device such as an XR controller.
For more information see the [url=https://www.khronos.org/registry/OpenXR/specs/1.0/html/xrspec.html#semantic-path-interaction-profiles]interaction profiles info in the OpenXR specification[/url].
*/
type Instance [1]classdb.OpenXRInteractionProfile
type Any interface {
	gd.IsClass
	AsOpenXRInteractionProfile() Instance
}

/*
Get the number of bindings in this interaction profile.
*/
func (self Instance) GetBindingCount() int {
	return int(int(class(self).GetBindingCount()))
}

/*
Retrieve the binding at this index.
*/
func (self Instance) GetBinding(index int) objects.OpenXRIPBinding {
	return objects.OpenXRIPBinding(class(self).GetBinding(gd.Int(index)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.OpenXRInteractionProfile

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OpenXRInteractionProfile"))
	return Instance{classdb.OpenXRInteractionProfile(object)}
}

func (self Instance) InteractionProfilePath() string {
	return string(class(self).GetInteractionProfilePath().String())
}

func (self Instance) SetInteractionProfilePath(value string) {
	class(self).SetInteractionProfilePath(gd.NewString(value))
}

func (self Instance) Bindings() Array.Any {
	return Array.Any(class(self).GetBindings())
}

func (self Instance) SetBindings(value Array.Any) {
	class(self).SetBindings(value)
}

//go:nosplit
func (self class) SetInteractionProfilePath(interaction_profile_path gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(interaction_profile_path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInteractionProfile.Bind_set_interaction_profile_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetInteractionProfilePath() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInteractionProfile.Bind_get_interaction_profile_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Get the number of bindings in this interaction profile.
*/
//go:nosplit
func (self class) GetBindingCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInteractionProfile.Bind_get_binding_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Retrieve the binding at this index.
*/
//go:nosplit
func (self class) GetBinding(index gd.Int) objects.OpenXRIPBinding {
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInteractionProfile.Bind_get_binding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.OpenXRIPBinding{classdb.OpenXRIPBinding(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBindings(bindings gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(bindings))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInteractionProfile.Bind_set_bindings, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBindings() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRInteractionProfile.Bind_get_bindings, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsOpenXRInteractionProfile() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsOpenXRInteractionProfile() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
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
	classdb.Register("OpenXRInteractionProfile", func(ptr gd.Object) any { return classdb.OpenXRInteractionProfile(ptr) })
}
