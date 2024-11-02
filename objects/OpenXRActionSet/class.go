package OpenXRActionSet

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Action sets in OpenXR define a collection of actions that can be activated in unison. This allows games to easily change between different states that require different inputs or need to reinterpret inputs. For instance we could have an action set that is active when a menu is open, an action set that is active when the player is freely walking around and an action set that is active when the player is controlling a vehicle.
Action sets can contain the same action with the same name, if such action sets are active at the same time the action set with the highest priority defines which binding is active.
*/
type Instance [1]classdb.OpenXRActionSet

/*
Retrieve the number of actions in our action set.
*/
func (self Instance) GetActionCount() int {
	return int(int(class(self).GetActionCount()))
}

/*
Add an action to this action set.
*/
func (self Instance) AddAction(action objects.OpenXRAction) {
	class(self).AddAction(action)
}

/*
Remove an action from this action set.
*/
func (self Instance) RemoveAction(action objects.OpenXRAction) {
	class(self).RemoveAction(action)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.OpenXRActionSet

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OpenXRActionSet"))
	return Instance{classdb.OpenXRActionSet(object)}
}

func (self Instance) LocalizedName() string {
	return string(class(self).GetLocalizedName().String())
}

func (self Instance) SetLocalizedName(value string) {
	class(self).SetLocalizedName(gd.NewString(value))
}

func (self Instance) Priority() int {
	return int(int(class(self).GetPriority()))
}

func (self Instance) SetPriority(value int) {
	class(self).SetPriority(gd.Int(value))
}

func (self Instance) Actions() gd.Array {
	return gd.Array(class(self).GetActions())
}

func (self Instance) SetActions(value gd.Array) {
	class(self).SetActions(value)
}

//go:nosplit
func (self class) SetLocalizedName(localized_name gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(localized_name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionSet.Bind_set_localized_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetLocalizedName() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionSet.Bind_get_localized_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPriority(priority gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, priority)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionSet.Bind_set_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPriority() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionSet.Bind_get_priority, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Retrieve the number of actions in our action set.
*/
//go:nosplit
func (self class) GetActionCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionSet.Bind_get_action_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetActions(actions gd.Array) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(actions))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionSet.Bind_set_actions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetActions() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionSet.Bind_get_actions, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

/*
Add an action to this action set.
*/
//go:nosplit
func (self class) AddAction(action objects.OpenXRAction) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionSet.Bind_add_action, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Remove an action from this action set.
*/
//go:nosplit
func (self class) RemoveAction(action objects.OpenXRAction) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(action[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRActionSet.Bind_remove_action, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsOpenXRActionSet() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsOpenXRActionSet() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	classdb.Register("OpenXRActionSet", func(ptr gd.Object) any { return classdb.OpenXRActionSet(ptr) })
}
