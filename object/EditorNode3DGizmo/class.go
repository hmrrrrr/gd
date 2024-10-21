package EditorNode3DGizmo

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Node3DGizmo"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Gizmo that is used for providing custom visualization and editing (handles and subgizmos) for [Node3D] objects. Can be overridden to create custom gizmos, but for simple gizmos creating a [EditorNode3DGizmoPlugin] is usually recommended.
	// EditorNode3DGizmo methods that can be overridden by a [Class] that extends it.
	type EditorNode3DGizmo interface {
		//Override this method to add all the gizmo elements whenever a gizmo update is requested. It's common to call [method clear] at the beginning of this method and then add visual elements depending on the node's properties.
		Redraw() 
		//Override this method to return the name of an edited handle (handles must have been previously added by [method add_handles]). Handles can be named for reference to the user when editing.
		//The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method add_handles] for more information).
		GetHandleName(id gd.Int, secondary bool) gd.String
		//Override this method to return [code]true[/code] whenever the given handle should be highlighted in the editor.
		//The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method add_handles] for more information).
		IsHandleHighlighted(id gd.Int, secondary bool) bool
		//Override this method to return the current value of a handle. This value will be requested at the start of an edit and used as the [code]restore[/code] argument in [method _commit_handle].
		//The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method add_handles] for more information).
		GetHandleValue(id gd.Int, secondary bool) gd.Variant
		BeginHandleAction(id gd.Int, secondary bool) 
		//Override this method to update the node properties when the user drags a gizmo handle (previously added with [method add_handles]). The provided [param point] is the mouse position in screen coordinates and the [param camera] can be used to convert it to raycasts.
		//The [param secondary] argument is [code]true[/code] when the edited handle is secondary (see [method add_handles] for more information).
		SetHandle(id gd.Int, secondary bool, camera object.Camera3D, point gd.Vector2) 
		//Override this method to commit a handle being edited (handles must have been previously added by [method add_handles]). This usually means creating an [UndoRedo] action for the change, using the current handle value as "do" and the [param restore] argument as "undo".
		//If the [param cancel] argument is [code]true[/code], the [param restore] value should be directly set, without any [UndoRedo] action.
		//The [param secondary] argument is [code]true[/code] when the committed handle is secondary (see [method add_handles] for more information).
		CommitHandle(id gd.Int, secondary bool, restore gd.Variant, cancel bool) 
		//Override this method to allow selecting subgizmos using mouse clicks. Given a [param camera] and a [param point] in screen coordinates, this method should return which subgizmo should be selected. The returned value should be a unique subgizmo identifier, which can have any non-negative value and will be used in other virtual methods like [method _get_subgizmo_transform] or [method _commit_subgizmos].
		SubgizmosIntersectRay(camera object.Camera3D, point gd.Vector2) gd.Int
		//Override this method to allow selecting subgizmos using mouse drag box selection. Given a [param camera] and a [param frustum], this method should return which subgizmos are contained within the frustum. The [param frustum] argument consists of an array with all the [Plane]s that make up the selection frustum. The returned value should contain a list of unique subgizmo identifiers, which can have any non-negative value and will be used in other virtual methods like [method _get_subgizmo_transform] or [method _commit_subgizmos].
		SubgizmosIntersectFrustum(camera object.Camera3D, frustum gd.ArrayOf[gd.Plane]) gd.PackedInt32Array
		//Override this method to update the node properties during subgizmo editing (see [method _subgizmos_intersect_ray] and [method _subgizmos_intersect_frustum]). The [param transform] is given in the [Node3D]'s local coordinate system.
		SetSubgizmoTransform(id gd.Int, transform gd.Transform3D) 
		//Override this method to return the current transform of a subgizmo. This transform will be requested at the start of an edit and used as the [code]restore[/code] argument in [method _commit_subgizmos].
		GetSubgizmoTransform(id gd.Int) gd.Transform3D
		//Override this method to commit a group of subgizmos being edited (see [method _subgizmos_intersect_ray] and [method _subgizmos_intersect_frustum]). This usually means creating an [UndoRedo] action for the change, using the current transforms as "do" and the [param restores] transforms as "undo".
		//If the [param cancel] argument is [code]true[/code], the [param restores] transforms should be directly set, without any [UndoRedo] action.
		CommitSubgizmos(ids gd.PackedInt32Array, restores gd.ArrayOf[gd.Transform3D], cancel bool) 
	}

*/
type Simple [1]classdb.EditorNode3DGizmo
func (Simple) _redraw(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}
func (Simple) _get_handle_name(impl func(ptr unsafe.Pointer, id int, secondary bool) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var id = gd.UnsafeGet[gd.Int](p_args,0)
		var secondary = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(id), secondary)
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}
func (Simple) _is_handle_highlighted(impl func(ptr unsafe.Pointer, id int, secondary bool) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var id = gd.UnsafeGet[gd.Int](p_args,0)
		var secondary = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(id), secondary)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _get_handle_value(impl func(ptr unsafe.Pointer, id int, secondary bool) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var id = gd.UnsafeGet[gd.Int](p_args,0)
		var secondary = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(id), secondary)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _begin_handle_action(impl func(ptr unsafe.Pointer, id int, secondary bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var id = gd.UnsafeGet[gd.Int](p_args,0)
		var secondary = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, int(id), secondary)
		gc.End()
	}
}
func (Simple) _set_handle(impl func(ptr unsafe.Pointer, id int, secondary bool, camera [1]classdb.Camera3D, point gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var id = gd.UnsafeGet[gd.Int](p_args,0)
		var secondary = gd.UnsafeGet[bool](p_args,1)
		var camera [1]classdb.Camera3D
		camera[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,2)}))
		var point = gd.UnsafeGet[gd.Vector2](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, int(id), secondary, camera, point)
		gc.End()
	}
}
func (Simple) _commit_handle(impl func(ptr unsafe.Pointer, id int, secondary bool, restore gd.Variant, cancel bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var id = gd.UnsafeGet[gd.Int](p_args,0)
		var secondary = gd.UnsafeGet[bool](p_args,1)
		var restore = mmm.Let[gd.Variant](gc.Lifetime, gc.API, gd.UnsafeGet[[3]uintptr](p_args,2))
		var cancel = gd.UnsafeGet[bool](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, int(id), secondary, restore, cancel)
		gc.End()
	}
}
func (Simple) _subgizmos_intersect_ray(impl func(ptr unsafe.Pointer, camera [1]classdb.Camera3D, point gd.Vector2) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var camera [1]classdb.Camera3D
		camera[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var point = gd.UnsafeGet[gd.Vector2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, camera, point)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}
func (Simple) _subgizmos_intersect_frustum(impl func(ptr unsafe.Pointer, camera [1]classdb.Camera3D, frustum gd.ArrayOf[gd.Plane]) gd.PackedInt32Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var camera [1]classdb.Camera3D
		camera[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var frustum = gd.TypedArray[gd.Plane](mmm.Let[gd.Array](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1)))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, camera, frustum)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}
func (Simple) _set_subgizmo_transform(impl func(ptr unsafe.Pointer, id int, transform gd.Transform3D) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var id = gd.UnsafeGet[gd.Int](p_args,0)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, int(id), transform)
		gc.End()
	}
}
func (Simple) _get_subgizmo_transform(impl func(ptr unsafe.Pointer, id int) gd.Transform3D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var id = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(id))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}
func (Simple) _commit_subgizmos(impl func(ptr unsafe.Pointer, ids gd.PackedInt32Array, restores gd.ArrayOf[gd.Transform3D], cancel bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var ids = mmm.Let[gd.PackedInt32Array](gc.Lifetime, gc.API, gd.UnsafeGet[[2]uintptr](p_args,0))
		var restores = gd.TypedArray[gd.Transform3D](mmm.Let[gd.Array](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1)))
		var cancel = gd.UnsafeGet[bool](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, ids, restores, cancel)
		gc.End()
	}
}
func (self Simple) AddLines(lines gd.PackedVector3Array, material [1]classdb.Material, billboard bool, modulate gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddLines(lines, material, billboard, modulate)
}
func (self Simple) AddMesh(mesh [1]classdb.Mesh, material [1]classdb.Material, transform gd.Transform3D, skeleton [1]classdb.SkinReference) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddMesh(mesh, material, transform, skeleton)
}
func (self Simple) AddCollisionSegments(segments gd.PackedVector3Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddCollisionSegments(segments)
}
func (self Simple) AddCollisionTriangles(triangles [1]classdb.TriangleMesh) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddCollisionTriangles(triangles)
}
func (self Simple) AddUnscaledBillboard(material [1]classdb.Material, default_scale float64, modulate gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddUnscaledBillboard(material, gd.Float(default_scale), modulate)
}
func (self Simple) AddHandles(handles gd.PackedVector3Array, material [1]classdb.Material, ids gd.PackedInt32Array, billboard bool, secondary bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddHandles(handles, material, ids, billboard, secondary)
}
func (self Simple) SetNode3d(node [1]classdb.Node) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetNode3d(node)
}
func (self Simple) GetNode3d() [1]classdb.Node3D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Node3D(Expert(self).GetNode3d(gc))
}
func (self Simple) GetPlugin() [1]classdb.EditorNode3DGizmoPlugin {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.EditorNode3DGizmoPlugin(Expert(self).GetPlugin(gc))
}
func (self Simple) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Clear()
}
func (self Simple) SetHidden(hidden bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHidden(hidden)
}
func (self Simple) IsSubgizmoSelected(id int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsSubgizmoSelected(gd.Int(id)))
}
func (self Simple) GetSubgizmoSelection() gd.PackedInt32Array {
	gc := gd.GarbageCollector(); _ = gc
	return gd.PackedInt32Array(Expert(self).GetSubgizmoSelection(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.EditorNode3DGizmo
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Override this method to add all the gizmo elements whenever a gizmo update is requested. It's common to call [method clear] at the beginning of this method and then add visual elements depending on the node's properties.
*/
func (class) _redraw(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Override this method to return the name of an edited handle (handles must have been previously added by [method add_handles]). Handles can be named for reference to the user when editing.
The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method add_handles] for more information).
*/
func (class) _get_handle_name(impl func(ptr unsafe.Pointer, id gd.Int, secondary bool) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var id = gd.UnsafeGet[gd.Int](p_args,0)
		var secondary = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, id, secondary)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Override this method to return [code]true[/code] whenever the given handle should be highlighted in the editor.
The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method add_handles] for more information).
*/
func (class) _is_handle_highlighted(impl func(ptr unsafe.Pointer, id gd.Int, secondary bool) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var id = gd.UnsafeGet[gd.Int](p_args,0)
		var secondary = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, id, secondary)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Override this method to return the current value of a handle. This value will be requested at the start of an edit and used as the [code]restore[/code] argument in [method _commit_handle].
The [param secondary] argument is [code]true[/code] when the requested handle is secondary (see [method add_handles] for more information).
*/
func (class) _get_handle_value(impl func(ptr unsafe.Pointer, id gd.Int, secondary bool) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var id = gd.UnsafeGet[gd.Int](p_args,0)
		var secondary = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, id, secondary)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

func (class) _begin_handle_action(impl func(ptr unsafe.Pointer, id gd.Int, secondary bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var id = gd.UnsafeGet[gd.Int](p_args,0)
		var secondary = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, id, secondary)
		ctx.End()
	}
}

/*
Override this method to update the node properties when the user drags a gizmo handle (previously added with [method add_handles]). The provided [param point] is the mouse position in screen coordinates and the [param camera] can be used to convert it to raycasts.
The [param secondary] argument is [code]true[/code] when the edited handle is secondary (see [method add_handles] for more information).
*/
func (class) _set_handle(impl func(ptr unsafe.Pointer, id gd.Int, secondary bool, camera object.Camera3D, point gd.Vector2) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var id = gd.UnsafeGet[gd.Int](p_args,0)
		var secondary = gd.UnsafeGet[bool](p_args,1)
		var camera object.Camera3D
		camera[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,2)}))
		var point = gd.UnsafeGet[gd.Vector2](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, id, secondary, camera, point)
		ctx.End()
	}
}

/*
Override this method to commit a handle being edited (handles must have been previously added by [method add_handles]). This usually means creating an [UndoRedo] action for the change, using the current handle value as "do" and the [param restore] argument as "undo".
If the [param cancel] argument is [code]true[/code], the [param restore] value should be directly set, without any [UndoRedo] action.
The [param secondary] argument is [code]true[/code] when the committed handle is secondary (see [method add_handles] for more information).
*/
func (class) _commit_handle(impl func(ptr unsafe.Pointer, id gd.Int, secondary bool, restore gd.Variant, cancel bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var id = gd.UnsafeGet[gd.Int](p_args,0)
		var secondary = gd.UnsafeGet[bool](p_args,1)
		var restore = mmm.Let[gd.Variant](ctx.Lifetime, ctx.API, gd.UnsafeGet[[3]uintptr](p_args,2))
		var cancel = gd.UnsafeGet[bool](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, id, secondary, restore, cancel)
		ctx.End()
	}
}

/*
Override this method to allow selecting subgizmos using mouse clicks. Given a [param camera] and a [param point] in screen coordinates, this method should return which subgizmo should be selected. The returned value should be a unique subgizmo identifier, which can have any non-negative value and will be used in other virtual methods like [method _get_subgizmo_transform] or [method _commit_subgizmos].
*/
func (class) _subgizmos_intersect_ray(impl func(ptr unsafe.Pointer, camera object.Camera3D, point gd.Vector2) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var camera object.Camera3D
		camera[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var point = gd.UnsafeGet[gd.Vector2](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, camera, point)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Override this method to allow selecting subgizmos using mouse drag box selection. Given a [param camera] and a [param frustum], this method should return which subgizmos are contained within the frustum. The [param frustum] argument consists of an array with all the [Plane]s that make up the selection frustum. The returned value should contain a list of unique subgizmo identifiers, which can have any non-negative value and will be used in other virtual methods like [method _get_subgizmo_transform] or [method _commit_subgizmos].
*/
func (class) _subgizmos_intersect_frustum(impl func(ptr unsafe.Pointer, camera object.Camera3D, frustum gd.ArrayOf[gd.Plane]) gd.PackedInt32Array, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var camera object.Camera3D
		camera[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var frustum = gd.TypedArray[gd.Plane](mmm.Let[gd.Array](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1)))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, camera, frustum)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Override this method to update the node properties during subgizmo editing (see [method _subgizmos_intersect_ray] and [method _subgizmos_intersect_frustum]). The [param transform] is given in the [Node3D]'s local coordinate system.
*/
func (class) _set_subgizmo_transform(impl func(ptr unsafe.Pointer, id gd.Int, transform gd.Transform3D) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var id = gd.UnsafeGet[gd.Int](p_args,0)
		var transform = gd.UnsafeGet[gd.Transform3D](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, id, transform)
		ctx.End()
	}
}

/*
Override this method to return the current transform of a subgizmo. This transform will be requested at the start of an edit and used as the [code]restore[/code] argument in [method _commit_subgizmos].
*/
func (class) _get_subgizmo_transform(impl func(ptr unsafe.Pointer, id gd.Int) gd.Transform3D, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var id = gd.UnsafeGet[gd.Int](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, id)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Override this method to commit a group of subgizmos being edited (see [method _subgizmos_intersect_ray] and [method _subgizmos_intersect_frustum]). This usually means creating an [UndoRedo] action for the change, using the current transforms as "do" and the [param restores] transforms as "undo".
If the [param cancel] argument is [code]true[/code], the [param restores] transforms should be directly set, without any [UndoRedo] action.
*/
func (class) _commit_subgizmos(impl func(ptr unsafe.Pointer, ids gd.PackedInt32Array, restores gd.ArrayOf[gd.Transform3D], cancel bool) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var ids = mmm.Let[gd.PackedInt32Array](ctx.Lifetime, ctx.API, gd.UnsafeGet[[2]uintptr](p_args,0))
		var restores = gd.TypedArray[gd.Transform3D](mmm.Let[gd.Array](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1)))
		var cancel = gd.UnsafeGet[bool](p_args,2)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, ids, restores, cancel)
		ctx.End()
	}
}

/*
Adds lines to the gizmo (as sets of 2 points), with a given material. The lines are used for visualizing the gizmo. Call this method during [method _redraw].
*/
//go:nosplit
func (self class) AddLines(lines gd.PackedVector3Array, material object.Material, billboard bool, modulate gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(lines))
	callframe.Arg(frame, mmm.Get(material[0].AsPointer())[0])
	callframe.Arg(frame, billboard)
	callframe.Arg(frame, modulate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorNode3DGizmo.Bind_add_lines, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a mesh to the gizmo with the specified [param material], local [param transform] and [param skeleton]. Call this method during [method _redraw].
*/
//go:nosplit
func (self class) AddMesh(mesh object.Mesh, material object.Material, transform gd.Transform3D, skeleton object.SkinReference)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(mesh[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(material[0].AsPointer())[0])
	callframe.Arg(frame, transform)
	callframe.Arg(frame, mmm.Get(skeleton[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorNode3DGizmo.Bind_add_mesh, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds the specified [param segments] to the gizmo's collision shape for picking. Call this method during [method _redraw].
*/
//go:nosplit
func (self class) AddCollisionSegments(segments gd.PackedVector3Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(segments))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorNode3DGizmo.Bind_add_collision_segments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds collision triangles to the gizmo for picking. A [TriangleMesh] can be generated from a regular [Mesh] too. Call this method during [method _redraw].
*/
//go:nosplit
func (self class) AddCollisionTriangles(triangles object.TriangleMesh)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(triangles[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorNode3DGizmo.Bind_add_collision_triangles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds an unscaled billboard for visualization and selection. Call this method during [method _redraw].
*/
//go:nosplit
func (self class) AddUnscaledBillboard(material object.Material, default_scale gd.Float, modulate gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(material[0].AsPointer())[0])
	callframe.Arg(frame, default_scale)
	callframe.Arg(frame, modulate)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorNode3DGizmo.Bind_add_unscaled_billboard, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a list of handles (points) which can be used to edit the properties of the gizmo's [Node3D]. The [param ids] argument can be used to specify a custom identifier for each handle, if an empty array is passed, the ids will be assigned automatically from the [param handles] argument order.
The [param secondary] argument marks the added handles as secondary, meaning they will normally have lower selection priority than regular handles. When the user is holding the shift key secondary handles will switch to have higher priority than regular handles. This change in priority can be used to place multiple handles at the same point while still giving the user control on their selection.
There are virtual methods which will be called upon editing of these handles. Call this method during [method _redraw].
*/
//go:nosplit
func (self class) AddHandles(handles gd.PackedVector3Array, material object.Material, ids gd.PackedInt32Array, billboard bool, secondary bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(handles))
	callframe.Arg(frame, mmm.Get(material[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(ids))
	callframe.Arg(frame, billboard)
	callframe.Arg(frame, secondary)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorNode3DGizmo.Bind_add_handles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the reference [Node3D] node for the gizmo. [param node] must inherit from [Node3D].
*/
//go:nosplit
func (self class) SetNode3d(node object.Node)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(node[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorNode3DGizmo.Bind_set_node_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the [Node3D] node associated with this gizmo.
*/
//go:nosplit
func (self class) GetNode3d(ctx gd.Lifetime) object.Node3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorNode3DGizmo.Bind_get_node_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Node3D
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the [EditorNode3DGizmoPlugin] that owns this gizmo. It's useful to retrieve materials using [method EditorNode3DGizmoPlugin.get_material].
*/
//go:nosplit
func (self class) GetPlugin(ctx gd.Lifetime) object.EditorNode3DGizmoPlugin {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorNode3DGizmo.Bind_get_plugin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.EditorNode3DGizmoPlugin
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Removes everything in the gizmo including meshes, collisions and handles.
*/
//go:nosplit
func (self class) Clear()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorNode3DGizmo.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the gizmo's hidden state. If [code]true[/code], the gizmo will be hidden. If [code]false[/code], it will be shown.
*/
//go:nosplit
func (self class) SetHidden(hidden bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, hidden)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorNode3DGizmo.Bind_set_hidden, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the given subgizmo is currently selected. Can be used to highlight selected elements during [method _redraw].
*/
//go:nosplit
func (self class) IsSubgizmoSelected(id gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorNode3DGizmo.Bind_is_subgizmo_selected, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a list of the currently selected subgizmos. Can be used to highlight selected elements during [method _redraw].
*/
//go:nosplit
func (self class) GetSubgizmoSelection(ctx gd.Lifetime) gd.PackedInt32Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorNode3DGizmo.Bind_get_subgizmo_selection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedInt32Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsEditorNode3DGizmo() Expert { return self[0].AsEditorNode3DGizmo() }


//go:nosplit
func (self Simple) AsEditorNode3DGizmo() Simple { return self[0].AsEditorNode3DGizmo() }


//go:nosplit
func (self class) AsNode3DGizmo() Node3DGizmo.Expert { return self[0].AsNode3DGizmo() }


//go:nosplit
func (self Simple) AsNode3DGizmo() Node3DGizmo.Simple { return self[0].AsNode3DGizmo() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_redraw": return reflect.ValueOf(self._redraw);
	case "_get_handle_name": return reflect.ValueOf(self._get_handle_name);
	case "_is_handle_highlighted": return reflect.ValueOf(self._is_handle_highlighted);
	case "_get_handle_value": return reflect.ValueOf(self._get_handle_value);
	case "_begin_handle_action": return reflect.ValueOf(self._begin_handle_action);
	case "_set_handle": return reflect.ValueOf(self._set_handle);
	case "_commit_handle": return reflect.ValueOf(self._commit_handle);
	case "_subgizmos_intersect_ray": return reflect.ValueOf(self._subgizmos_intersect_ray);
	case "_subgizmos_intersect_frustum": return reflect.ValueOf(self._subgizmos_intersect_frustum);
	case "_set_subgizmo_transform": return reflect.ValueOf(self._set_subgizmo_transform);
	case "_get_subgizmo_transform": return reflect.ValueOf(self._get_subgizmo_transform);
	case "_commit_subgizmos": return reflect.ValueOf(self._commit_subgizmos);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	case "_redraw": return reflect.ValueOf(self._redraw);
	case "_get_handle_name": return reflect.ValueOf(self._get_handle_name);
	case "_is_handle_highlighted": return reflect.ValueOf(self._is_handle_highlighted);
	case "_get_handle_value": return reflect.ValueOf(self._get_handle_value);
	case "_begin_handle_action": return reflect.ValueOf(self._begin_handle_action);
	case "_set_handle": return reflect.ValueOf(self._set_handle);
	case "_commit_handle": return reflect.ValueOf(self._commit_handle);
	case "_subgizmos_intersect_ray": return reflect.ValueOf(self._subgizmos_intersect_ray);
	case "_subgizmos_intersect_frustum": return reflect.ValueOf(self._subgizmos_intersect_frustum);
	case "_set_subgizmo_transform": return reflect.ValueOf(self._set_subgizmo_transform);
	case "_get_subgizmo_transform": return reflect.ValueOf(self._get_subgizmo_transform);
	case "_commit_subgizmos": return reflect.ValueOf(self._commit_subgizmos);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("EditorNode3DGizmo", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
