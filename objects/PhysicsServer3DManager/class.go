package PhysicsServer3DManager

import "unsafe"
import "sync"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
[PhysicsServer3DManager] is the API for registering [PhysicsServer3D] implementations and for setting the default implementation.
[b]Note:[/b] It is not possible to switch physics servers at runtime. This class is only used on startup at the server initialization level, by Godot itself and possibly by GDExtensions.
*/
var self objects.PhysicsServer3DManager
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.PhysicsServer3DManager)
	self = *(*objects.PhysicsServer3DManager)(unsafe.Pointer(&obj))
}

/*
Register a [PhysicsServer3D] implementation by passing a [param name] and a [Callable] that returns a [PhysicsServer3D] object.
*/
func RegisterServer(name string, create_callback func() objects.PhysicsServer3D) {
	once.Do(singleton)
	class(self).RegisterServer(gd.NewString(name), gd.NewCallable(create_callback))
}

/*
Set the default [PhysicsServer3D] implementation to the one identified by [param name], if [param priority] is greater than the priority of the current default implementation.
*/
func SetDefaultServer(name string, priority int) {
	once.Do(singleton)
	class(self).SetDefaultServer(gd.NewString(name), gd.Int(priority))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]classdb.PhysicsServer3DManager

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Register a [PhysicsServer3D] implementation by passing a [param name] and a [Callable] that returns a [PhysicsServer3D] object.
*/
//go:nosplit
func (self class) RegisterServer(name gd.String, create_callback gd.Callable) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(create_callback))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsServer3DManager.Bind_register_server, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Set the default [PhysicsServer3D] implementation to the one identified by [param name], if [param priority] is greater than the priority of the current default implementation.
*/
//go:nosplit
func (self class) SetDefaultServer(name gd.String, priority gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, priority)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsServer3DManager.Bind_set_default_server, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {
	classdb.Register("PhysicsServer3DManager", func(ptr gd.Object) any { return [1]classdb.PhysicsServer3DManager{classdb.PhysicsServer3DManager(ptr)} })
}
