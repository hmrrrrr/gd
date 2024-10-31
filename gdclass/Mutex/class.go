package Mutex

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
A synchronization mutex (mutual exclusion). This is used to synchronize multiple [Thread]s, and is equivalent to a binary [Semaphore]. It guarantees that only one thread can access a critical section at a time.
This is a reentrant mutex, meaning that it can be locked multiple times by one thread, provided it also unlocks it as many times.
[b]Warning:[/b] Mutexes must be used carefully to avoid deadlocks.
[b]Warning:[/b] To ensure proper cleanup without crashes or deadlocks, the following conditions must be met:
- When a [Mutex]'s reference count reaches zero and it is therefore destroyed, no threads (including the one on which the destruction will happen) must have it locked.
- When a [Thread]'s reference count reaches zero and it is therefore destroyed, it must not have any mutex locked.
*/
type Instance [1]classdb.Mutex

/*
Locks this [Mutex], blocks until it is unlocked by the current owner.
[b]Note:[/b] This function returns without blocking if the thread already has ownership of the mutex.
*/
func (self Instance) Lock() {
	class(self).Lock()
}

/*
Tries locking this [Mutex], but does not block. Returns [code]true[/code] on success, [code]false[/code] otherwise.
[b]Note:[/b] This function returns [code]true[/code] if the thread already has ownership of the mutex.
*/
func (self Instance) TryLock() bool {
	return bool(class(self).TryLock())
}

/*
Unlocks this [Mutex], leaving it to other threads.
[b]Note:[/b] If a thread called [method lock] or [method try_lock] multiple times while already having ownership of the mutex, it must also call [method unlock] the same number of times in order to unlock it correctly.
[b]Warning:[/b] Calling [method unlock] more times that [method lock] on a given thread, thus ending up trying to unlock a non-locked mutex, is wrong and may causes crashes or deadlocks.
*/
func (self Instance) Unlock() {
	class(self).Unlock()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Mutex

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Mutex"))
	return Instance{classdb.Mutex(object)}
}

/*
Locks this [Mutex], blocks until it is unlocked by the current owner.
[b]Note:[/b] This function returns without blocking if the thread already has ownership of the mutex.
*/
//go:nosplit
func (self class) Lock() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mutex.Bind_lock, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Tries locking this [Mutex], but does not block. Returns [code]true[/code] on success, [code]false[/code] otherwise.
[b]Note:[/b] This function returns [code]true[/code] if the thread already has ownership of the mutex.
*/
//go:nosplit
func (self class) TryLock() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mutex.Bind_try_lock, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Unlocks this [Mutex], leaving it to other threads.
[b]Note:[/b] If a thread called [method lock] or [method try_lock] multiple times while already having ownership of the mutex, it must also call [method unlock] the same number of times in order to unlock it correctly.
[b]Warning:[/b] Calling [method unlock] more times that [method lock] on a given thread, thus ending up trying to unlock a non-locked mutex, is wrong and may causes crashes or deadlocks.
*/
//go:nosplit
func (self class) Unlock() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Mutex.Bind_unlock, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsMutex() Advanced              { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMutex() Instance           { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() { classdb.Register("Mutex", func(ptr gd.Object) any { return classdb.Mutex(ptr) }) }
