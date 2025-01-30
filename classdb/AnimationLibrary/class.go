// Package AnimationLibrary provides methods for working with AnimationLibrary object instances.
package AnimationLibrary

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
An animation library stores a set of animations accessible through [StringName] keys, for use with [AnimationPlayer] nodes.
*/
type Instance [1]gdclass.AnimationLibrary

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAnimationLibrary() Instance
}

/*
Adds the [param animation] to the library, accessible by the key [param name].
*/
func (self Instance) AddAnimation(name string, animation [1]gdclass.Animation) error { //gd:AnimationLibrary.add_animation
	return error(gd.ToError(class(self).AddAnimation(String.Name(String.New(name)), animation)))
}

/*
Removes the [Animation] with the key [param name].
*/
func (self Instance) RemoveAnimation(name string) { //gd:AnimationLibrary.remove_animation
	class(self).RemoveAnimation(String.Name(String.New(name)))
}

/*
Changes the key of the [Animation] associated with the key [param name] to [param newname].
*/
func (self Instance) RenameAnimation(name string, newname string) { //gd:AnimationLibrary.rename_animation
	class(self).RenameAnimation(String.Name(String.New(name)), String.Name(String.New(newname)))
}

/*
Returns [code]true[/code] if the library stores an [Animation] with [param name] as the key.
*/
func (self Instance) HasAnimation(name string) bool { //gd:AnimationLibrary.has_animation
	return bool(class(self).HasAnimation(String.Name(String.New(name))))
}

/*
Returns the [Animation] with the key [param name]. If the animation does not exist, [code]null[/code] is returned and an error is logged.
*/
func (self Instance) GetAnimation(name string) [1]gdclass.Animation { //gd:AnimationLibrary.get_animation
	return [1]gdclass.Animation(class(self).GetAnimation(String.Name(String.New(name))))
}

/*
Returns the keys for the [Animation]s stored in the library.
*/
func (self Instance) GetAnimationList() []string { //gd:AnimationLibrary.get_animation_list
	return []string(gd.ArrayAs[[]string](gd.InternalArray(class(self).GetAnimationList())))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AnimationLibrary

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimationLibrary"))
	casted := Instance{*(*gdclass.AnimationLibrary)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Adds the [param animation] to the library, accessible by the key [param name].
*/
//go:nosplit
func (self class) AddAnimation(name String.Name, animation [1]gdclass.Animation) Error.Code { //gd:AnimationLibrary.add_animation
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(animation[0])[0])
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationLibrary.Bind_add_animation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Removes the [Animation] with the key [param name].
*/
//go:nosplit
func (self class) RemoveAnimation(name String.Name) { //gd:AnimationLibrary.remove_animation
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationLibrary.Bind_remove_animation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Changes the key of the [Animation] associated with the key [param name] to [param newname].
*/
//go:nosplit
func (self class) RenameAnimation(name String.Name, newname String.Name) { //gd:AnimationLibrary.rename_animation
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(newname)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationLibrary.Bind_rename_animation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the library stores an [Animation] with [param name] as the key.
*/
//go:nosplit
func (self class) HasAnimation(name String.Name) bool { //gd:AnimationLibrary.has_animation
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationLibrary.Bind_has_animation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [Animation] with the key [param name]. If the animation does not exist, [code]null[/code] is returned and an error is logged.
*/
//go:nosplit
func (self class) GetAnimation(name String.Name) [1]gdclass.Animation { //gd:AnimationLibrary.get_animation
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationLibrary.Bind_get_animation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Animation{gd.PointerWithOwnershipTransferredToGo[gdclass.Animation](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the keys for the [Animation]s stored in the library.
*/
//go:nosplit
func (self class) GetAnimationList() Array.Contains[String.Name] { //gd:AnimationLibrary.get_animation_list
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimationLibrary.Bind_get_animation_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[String.Name]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}
func (self Instance) OnAnimationAdded(cb func(name string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("animation_added"), gd.NewCallable(cb), 0)
}

func (self Instance) OnAnimationRemoved(cb func(name string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("animation_removed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnAnimationRenamed(cb func(name string, to_name string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("animation_renamed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnAnimationChanged(cb func(name string)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("animation_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsAnimationLibrary() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAnimationLibrary() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("AnimationLibrary", func(ptr gd.Object) any {
		return [1]gdclass.AnimationLibrary{*(*gdclass.AnimationLibrary)(unsafe.Pointer(&ptr))}
	})
}
