package SkeletonModificationStack2D

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
This resource is used by the Skeleton and holds a stack of [SkeletonModification2D]s.
This controls the order of the modifications and how they are applied. Modification order is especially important for full-body IK setups, as you need to execute the modifications in the correct order to get the desired results. For example, you want to execute a modification on the spine [i]before[/i] the arms on a humanoid skeleton.
This resource also controls how strongly all of the modifications are applied to the [Skeleton2D].

*/
type Go [1]classdb.SkeletonModificationStack2D

/*
Sets up the modification stack so it can execute. This function should be called by [Skeleton2D] and shouldn't be manually called unless you know what you are doing.
*/
func (self Go) Setup() {
	class(self).Setup()
}

/*
Executes all of the [SkeletonModification2D]s in the stack that use the same execution mode as the passed-in [param execution_mode], starting from index [code]0[/code] to [member modification_count].
[b]Note:[/b] The order of the modifications can matter depending on the modifications. For example, modifications on a spine should operate before modifications on the arms in order to get proper results.
*/
func (self Go) Execute(delta float64, execution_mode int) {
	class(self).Execute(gd.Float(delta), gd.Int(execution_mode))
}

/*
Enables all [SkeletonModification2D]s in the stack.
*/
func (self Go) EnableAllModifications(enabled bool) {
	class(self).EnableAllModifications(enabled)
}

/*
Returns the [SkeletonModification2D] at the passed-in index, [param mod_idx].
*/
func (self Go) GetModification(mod_idx int) gdclass.SkeletonModification2D {
	return gdclass.SkeletonModification2D(class(self).GetModification(gd.Int(mod_idx)))
}

/*
Adds the passed-in [SkeletonModification2D] to the stack.
*/
func (self Go) AddModification(modification gdclass.SkeletonModification2D) {
	class(self).AddModification(modification)
}

/*
Deletes the [SkeletonModification2D] at the index position [param mod_idx], if it exists.
*/
func (self Go) DeleteModification(mod_idx int) {
	class(self).DeleteModification(gd.Int(mod_idx))
}

/*
Sets the modification at [param mod_idx] to the passed-in modification, [param modification].
*/
func (self Go) SetModification(mod_idx int, modification gdclass.SkeletonModification2D) {
	class(self).SetModification(gd.Int(mod_idx), modification)
}

/*
Returns a boolean that indicates whether the modification stack is setup and can execute.
*/
func (self Go) GetIsSetup() bool {
	return bool(class(self).GetIsSetup())
}

/*
Returns the [Skeleton2D] node that the SkeletonModificationStack2D is bound to.
*/
func (self Go) GetSkeleton() gdclass.Skeleton2D {
	return gdclass.Skeleton2D(class(self).GetSkeleton())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.SkeletonModificationStack2D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("SkeletonModificationStack2D"))
	return Go{classdb.SkeletonModificationStack2D(object)}
}

func (self Go) Enabled() bool {
		return bool(class(self).GetEnabled())
}

func (self Go) SetEnabled(value bool) {
	class(self).SetEnabled(value)
}

func (self Go) Strength() float64 {
		return float64(float64(class(self).GetStrength()))
}

func (self Go) SetStrength(value float64) {
	class(self).SetStrength(gd.Float(value))
}

func (self Go) ModificationCount() int {
		return int(int(class(self).GetModificationCount()))
}

func (self Go) SetModificationCount(value int) {
	class(self).SetModificationCount(gd.Int(value))
}

/*
Sets up the modification stack so it can execute. This function should be called by [Skeleton2D] and shouldn't be manually called unless you know what you are doing.
*/
//go:nosplit
func (self class) Setup()  {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_setup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Executes all of the [SkeletonModification2D]s in the stack that use the same execution mode as the passed-in [param execution_mode], starting from index [code]0[/code] to [member modification_count].
[b]Note:[/b] The order of the modifications can matter depending on the modifications. For example, modifications on a spine should operate before modifications on the arms in order to get proper results.
*/
//go:nosplit
func (self class) Execute(delta gd.Float, execution_mode gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, delta)
	callframe.Arg(frame, execution_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_execute, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Enables all [SkeletonModification2D]s in the stack.
*/
//go:nosplit
func (self class) EnableAllModifications(enabled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_enable_all_modifications, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [SkeletonModification2D] at the passed-in index, [param mod_idx].
*/
//go:nosplit
func (self class) GetModification(mod_idx gd.Int) gdclass.SkeletonModification2D {
	var frame = callframe.New()
	callframe.Arg(frame, mod_idx)
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_get_modification, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.SkeletonModification2D{classdb.SkeletonModification2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
Adds the passed-in [SkeletonModification2D] to the stack.
*/
//go:nosplit
func (self class) AddModification(modification gdclass.SkeletonModification2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(modification[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_add_modification, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Deletes the [SkeletonModification2D] at the index position [param mod_idx], if it exists.
*/
//go:nosplit
func (self class) DeleteModification(mod_idx gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, mod_idx)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_delete_modification, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the modification at [param mod_idx] to the passed-in modification, [param modification].
*/
//go:nosplit
func (self class) SetModification(mod_idx gd.Int, modification gdclass.SkeletonModification2D)  {
	var frame = callframe.New()
	callframe.Arg(frame, mod_idx)
	callframe.Arg(frame, discreet.Get(modification[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_set_modification, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetModificationCount(count gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, count)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_set_modification_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetModificationCount() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_get_modification_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a boolean that indicates whether the modification stack is setup and can execute.
*/
//go:nosplit
func (self class) GetIsSetup() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_get_is_setup, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetEnabled(enabled bool)  {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_set_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_get_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStrength(strength gd.Float)  {
	var frame = callframe.New()
	callframe.Arg(frame, strength)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_set_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStrength() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_get_strength, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [Skeleton2D] node that the SkeletonModificationStack2D is bound to.
*/
//go:nosplit
func (self class) GetSkeleton() gdclass.Skeleton2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.SkeletonModificationStack2D.Bind_get_skeleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Skeleton2D{classdb.Skeleton2D(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self class) AsSkeletonModificationStack2D() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsSkeletonModificationStack2D() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("SkeletonModificationStack2D", func(ptr gd.Object) any { return classdb.SkeletonModificationStack2D(ptr) })}
