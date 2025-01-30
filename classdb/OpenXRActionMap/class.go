// Package OpenXRActionMap provides methods for working with OpenXRActionMap object instances.
package OpenXRActionMap

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Resource"
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
OpenXR uses an action system similar to Godots Input map system to bind inputs and outputs on various types of XR controllers to named actions. OpenXR specifies more detail on these inputs and outputs than Godot supports.
Another important distinction is that OpenXR offers no control over these bindings. The bindings we register are suggestions, it is up to the XR runtime to offer users the ability to change these bindings. This allows the XR runtime to fill in the gaps if new hardware becomes available.
The action map therefore needs to be loaded at startup and can't be changed afterwards. This resource is a container for the entire action map.
*/
type Instance [1]gdclass.OpenXRActionMap

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsOpenXRActionMap() Instance
}

/*
Retrieve the number of actions sets in our action map.
*/
func (self Instance) GetActionSetCount() int { //gd:OpenXRActionMap.get_action_set_count
	return int(int(class(self).GetActionSetCount()))
}

/*
Retrieve an action set by name.
*/
func (self Instance) FindActionSet(name string) [1]gdclass.OpenXRActionSet { //gd:OpenXRActionMap.find_action_set
	return [1]gdclass.OpenXRActionSet(class(self).FindActionSet(String.New(name)))
}

/*
Retrieve the action set at this index.
*/
func (self Instance) GetActionSet(idx int) [1]gdclass.OpenXRActionSet { //gd:OpenXRActionMap.get_action_set
	return [1]gdclass.OpenXRActionSet(class(self).GetActionSet(int64(idx)))
}

/*
Add an action set.
*/
func (self Instance) AddActionSet(action_set [1]gdclass.OpenXRActionSet) { //gd:OpenXRActionMap.add_action_set
	class(self).AddActionSet(action_set)
}

/*
Remove an action set.
*/
func (self Instance) RemoveActionSet(action_set [1]gdclass.OpenXRActionSet) { //gd:OpenXRActionMap.remove_action_set
	class(self).RemoveActionSet(action_set)
}

/*
Retrieve the number of interaction profiles in our action map.
*/
func (self Instance) GetInteractionProfileCount() int { //gd:OpenXRActionMap.get_interaction_profile_count
	return int(int(class(self).GetInteractionProfileCount()))
}

/*
Find an interaction profile by its name (path).
*/
func (self Instance) FindInteractionProfile(name string) [1]gdclass.OpenXRInteractionProfile { //gd:OpenXRActionMap.find_interaction_profile
	return [1]gdclass.OpenXRInteractionProfile(class(self).FindInteractionProfile(String.New(name)))
}

/*
Get the interaction profile at this index.
*/
func (self Instance) GetInteractionProfile(idx int) [1]gdclass.OpenXRInteractionProfile { //gd:OpenXRActionMap.get_interaction_profile
	return [1]gdclass.OpenXRInteractionProfile(class(self).GetInteractionProfile(int64(idx)))
}

/*
Add an interaction profile.
*/
func (self Instance) AddInteractionProfile(interaction_profile [1]gdclass.OpenXRInteractionProfile) { //gd:OpenXRActionMap.add_interaction_profile
	class(self).AddInteractionProfile(interaction_profile)
}

/*
Remove an interaction profile.
*/
func (self Instance) RemoveInteractionProfile(interaction_profile [1]gdclass.OpenXRInteractionProfile) { //gd:OpenXRActionMap.remove_interaction_profile
	class(self).RemoveInteractionProfile(interaction_profile)
}

/*
Setup this action set with our default actions.
*/
func (self Instance) CreateDefaultActionSets() { //gd:OpenXRActionMap.create_default_action_sets
	class(self).CreateDefaultActionSets()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.OpenXRActionMap

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OpenXRActionMap"))
	casted := Instance{*(*gdclass.OpenXRActionMap)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) ActionSets() []any {
	return []any(gd.ArrayAs[[]any](gd.InternalArray(class(self).GetActionSets())))
}

func (self Instance) SetActionSets(value []any) {
	class(self).SetActionSets(gd.EngineArrayFromSlice(value))
}

func (self Instance) InteractionProfiles() []any {
	return []any(gd.ArrayAs[[]any](gd.InternalArray(class(self).GetInteractionProfiles())))
}

func (self Instance) SetInteractionProfiles(value []any) {
	class(self).SetInteractionProfiles(gd.EngineArrayFromSlice(value))
}

//go:nosplit
func (self class) SetActionSets(action_sets Array.Any) { //gd:OpenXRActionMap.set_action_sets
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(action_sets)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_set_action_sets, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetActionSets() Array.Any { //gd:OpenXRActionMap.get_action_sets
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_get_action_sets, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Retrieve the number of actions sets in our action map.
*/
//go:nosplit
func (self class) GetActionSetCount() int64 { //gd:OpenXRActionMap.get_action_set_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_get_action_set_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Retrieve an action set by name.
*/
//go:nosplit
func (self class) FindActionSet(name String.Readable) [1]gdclass.OpenXRActionSet { //gd:OpenXRActionMap.find_action_set
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_find_action_set, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.OpenXRActionSet{gd.PointerWithOwnershipTransferredToGo[gdclass.OpenXRActionSet](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Retrieve the action set at this index.
*/
//go:nosplit
func (self class) GetActionSet(idx int64) [1]gdclass.OpenXRActionSet { //gd:OpenXRActionMap.get_action_set
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_get_action_set, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.OpenXRActionSet{gd.PointerWithOwnershipTransferredToGo[gdclass.OpenXRActionSet](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Add an action set.
*/
//go:nosplit
func (self class) AddActionSet(action_set [1]gdclass.OpenXRActionSet) { //gd:OpenXRActionMap.add_action_set
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action_set[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_add_action_set, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Remove an action set.
*/
//go:nosplit
func (self class) RemoveActionSet(action_set [1]gdclass.OpenXRActionSet) { //gd:OpenXRActionMap.remove_action_set
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action_set[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_remove_action_set, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetInteractionProfiles(interaction_profiles Array.Any) { //gd:OpenXRActionMap.set_interaction_profiles
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(interaction_profiles)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_set_interaction_profiles, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetInteractionProfiles() Array.Any { //gd:OpenXRActionMap.get_interaction_profiles
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_get_interaction_profiles, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[variant.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Retrieve the number of interaction profiles in our action map.
*/
//go:nosplit
func (self class) GetInteractionProfileCount() int64 { //gd:OpenXRActionMap.get_interaction_profile_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_get_interaction_profile_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Find an interaction profile by its name (path).
*/
//go:nosplit
func (self class) FindInteractionProfile(name String.Readable) [1]gdclass.OpenXRInteractionProfile { //gd:OpenXRActionMap.find_interaction_profile
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_find_interaction_profile, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.OpenXRInteractionProfile{gd.PointerWithOwnershipTransferredToGo[gdclass.OpenXRInteractionProfile](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Get the interaction profile at this index.
*/
//go:nosplit
func (self class) GetInteractionProfile(idx int64) [1]gdclass.OpenXRInteractionProfile { //gd:OpenXRActionMap.get_interaction_profile
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_get_interaction_profile, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.OpenXRInteractionProfile{gd.PointerWithOwnershipTransferredToGo[gdclass.OpenXRInteractionProfile](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Add an interaction profile.
*/
//go:nosplit
func (self class) AddInteractionProfile(interaction_profile [1]gdclass.OpenXRInteractionProfile) { //gd:OpenXRActionMap.add_interaction_profile
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(interaction_profile[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_add_interaction_profile, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Remove an interaction profile.
*/
//go:nosplit
func (self class) RemoveInteractionProfile(interaction_profile [1]gdclass.OpenXRInteractionProfile) { //gd:OpenXRActionMap.remove_interaction_profile
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(interaction_profile[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_remove_interaction_profile, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Setup this action set with our default actions.
*/
//go:nosplit
func (self class) CreateDefaultActionSets() { //gd:OpenXRActionMap.create_default_action_sets
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionMap.Bind_create_default_action_sets, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsOpenXRActionMap() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsOpenXRActionMap() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("OpenXRActionMap", func(ptr gd.Object) any {
		return [1]gdclass.OpenXRActionMap{*(*gdclass.OpenXRActionMap)(unsafe.Pointer(&ptr))}
	})
}
