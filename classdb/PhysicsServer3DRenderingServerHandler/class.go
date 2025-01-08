// Package PhysicsServer3DRenderingServerHandler provides methods for working with PhysicsServer3DRenderingServerHandler object instances.
package PhysicsServer3DRenderingServerHandler

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Vector3"
import "graphics.gd/variant/AABB"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

type Instance [1]gdclass.PhysicsServer3DRenderingServerHandler

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPhysicsServer3DRenderingServerHandler() Instance
}
type Interface interface {
	//Called by the [PhysicsServer3D] to set the position for the [SoftBody3D] vertex at the index specified by [param vertex_id].
	//[b]Note:[/b] The [param vertex] parameter used to be of type [code]const void*[/code] prior to Godot 4.2.
	SetVertex(vertex_id int, vertex Vector3.XYZ)
	//Called by the [PhysicsServer3D] to set the normal for the [SoftBody3D] vertex at the index specified by [param vertex_id].
	//[b]Note:[/b] The [param normal] parameter used to be of type [code]const void*[/code] prior to Godot 4.2.
	SetNormal(vertex_id int, normal Vector3.XYZ)
	//Called by the [PhysicsServer3D] to set the bounding box for the [SoftBody3D].
	SetAabb(aabb AABB.PositionSize)
}

// Implementation implements [Interface] with empty methods.
type Implementation struct{}

func (self Implementation) SetVertex(vertex_id int, vertex Vector3.XYZ) { return }
func (self Implementation) SetNormal(vertex_id int, normal Vector3.XYZ) { return }
func (self Implementation) SetAabb(aabb AABB.PositionSize)              { return }

/*
Called by the [PhysicsServer3D] to set the position for the [SoftBody3D] vertex at the index specified by [param vertex_id].
[b]Note:[/b] The [param vertex] parameter used to be of type [code]const void*[/code] prior to Godot 4.2.
*/
func (Instance) _set_vertex(impl func(ptr unsafe.Pointer, vertex_id int, vertex Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var vertex_id = gd.UnsafeGet[gd.Int](p_args, 0)
		var vertex = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(vertex_id), vertex)
	}
}

/*
Called by the [PhysicsServer3D] to set the normal for the [SoftBody3D] vertex at the index specified by [param vertex_id].
[b]Note:[/b] The [param normal] parameter used to be of type [code]const void*[/code] prior to Godot 4.2.
*/
func (Instance) _set_normal(impl func(ptr unsafe.Pointer, vertex_id int, normal Vector3.XYZ)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var vertex_id = gd.UnsafeGet[gd.Int](p_args, 0)
		var normal = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(vertex_id), normal)
	}
}

/*
Called by the [PhysicsServer3D] to set the bounding box for the [SoftBody3D].
*/
func (Instance) _set_aabb(impl func(ptr unsafe.Pointer, aabb AABB.PositionSize)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var aabb = gd.UnsafeGet[gd.AABB](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, aabb)
	}
}

/*
Sets the position for the [SoftBody3D] vertex at the index specified by [param vertex_id].
*/
func (self Instance) SetVertex(vertex_id int, vertex Vector3.XYZ) {
	class(self).SetVertex(gd.Int(vertex_id), gd.Vector3(vertex))
}

/*
Sets the normal for the [SoftBody3D] vertex at the index specified by [param vertex_id].
*/
func (self Instance) SetNormal(vertex_id int, normal Vector3.XYZ) {
	class(self).SetNormal(gd.Int(vertex_id), gd.Vector3(normal))
}

/*
Sets the bounding box for the [SoftBody3D].
*/
func (self Instance) SetAabb(aabb AABB.PositionSize) {
	class(self).SetAabb(gd.AABB(aabb))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PhysicsServer3DRenderingServerHandler

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PhysicsServer3DRenderingServerHandler"))
	casted := Instance{*(*gdclass.PhysicsServer3DRenderingServerHandler)(unsafe.Pointer(&object))}
	return casted
}

/*
Called by the [PhysicsServer3D] to set the position for the [SoftBody3D] vertex at the index specified by [param vertex_id].
[b]Note:[/b] The [param vertex] parameter used to be of type [code]const void*[/code] prior to Godot 4.2.
*/
func (class) _set_vertex(impl func(ptr unsafe.Pointer, vertex_id gd.Int, vertex gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var vertex_id = gd.UnsafeGet[gd.Int](p_args, 0)
		var vertex = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, vertex_id, vertex)
	}
}

/*
Called by the [PhysicsServer3D] to set the normal for the [SoftBody3D] vertex at the index specified by [param vertex_id].
[b]Note:[/b] The [param normal] parameter used to be of type [code]const void*[/code] prior to Godot 4.2.
*/
func (class) _set_normal(impl func(ptr unsafe.Pointer, vertex_id gd.Int, normal gd.Vector3)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var vertex_id = gd.UnsafeGet[gd.Int](p_args, 0)
		var normal = gd.UnsafeGet[gd.Vector3](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, vertex_id, normal)
	}
}

/*
Called by the [PhysicsServer3D] to set the bounding box for the [SoftBody3D].
*/
func (class) _set_aabb(impl func(ptr unsafe.Pointer, aabb gd.AABB)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var aabb = gd.UnsafeGet[gd.AABB](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, aabb)
	}
}

/*
Sets the position for the [SoftBody3D] vertex at the index specified by [param vertex_id].
*/
//go:nosplit
func (self class) SetVertex(vertex_id gd.Int, vertex gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, vertex_id)
	callframe.Arg(frame, vertex)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsServer3DRenderingServerHandler.Bind_set_vertex, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the normal for the [SoftBody3D] vertex at the index specified by [param vertex_id].
*/
//go:nosplit
func (self class) SetNormal(vertex_id gd.Int, normal gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, vertex_id)
	callframe.Arg(frame, normal)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsServer3DRenderingServerHandler.Bind_set_normal, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the bounding box for the [SoftBody3D].
*/
//go:nosplit
func (self class) SetAabb(aabb gd.AABB) {
	var frame = callframe.New()
	callframe.Arg(frame, aabb)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PhysicsServer3DRenderingServerHandler.Bind_set_aabb, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsPhysicsServer3DRenderingServerHandler() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPhysicsServer3DRenderingServerHandler() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_set_vertex":
		return reflect.ValueOf(self._set_vertex)
	case "_set_normal":
		return reflect.ValueOf(self._set_normal)
	case "_set_aabb":
		return reflect.ValueOf(self._set_aabb)
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_set_vertex":
		return reflect.ValueOf(self._set_vertex)
	case "_set_normal":
		return reflect.ValueOf(self._set_normal)
	case "_set_aabb":
		return reflect.ValueOf(self._set_aabb)
	default:
		return gd.VirtualByName(Object.Instance(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("PhysicsServer3DRenderingServerHandler", func(ptr gd.Object) any {
		return [1]gdclass.PhysicsServer3DRenderingServerHandler{*(*gdclass.PhysicsServer3DRenderingServerHandler)(unsafe.Pointer(&ptr))}
	})
}
