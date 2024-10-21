package PackedScene

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A simplified interface to a scene file. Provides access to operations and checks that can be performed on the scene resource itself.
Can be used to save a node to a file. When saving, the node as well as all the nodes it owns get saved (see [member Node.owner] property).
[b]Note:[/b] The node doesn't need to own itself.
[b]Example of loading a saved scene:[/b]
[codeblocks]
[gdscript]
# Use load() instead of preload() if the path isn't known at compile-time.
var scene = preload("res://scene.tscn").instantiate()
# Add the node as a child of the node the script is attached to.
add_child(scene)
[/gdscript]
[csharp]
// C# has no preload, so you have to always use ResourceLoader.Load<PackedScene>().
var scene = ResourceLoader.Load<PackedScene>("res://scene.tscn").Instantiate();
// Add the node as a child of the node the script is attached to.
AddChild(scene);
[/csharp]
[/codeblocks]
[b]Example of saving a node with different owners:[/b] The following example creates 3 objects: [Node2D] ([code]node[/code]), [RigidBody2D] ([code]body[/code]) and [CollisionObject2D] ([code]collision[/code]). [code]collision[/code] is a child of [code]body[/code] which is a child of [code]node[/code]. Only [code]body[/code] is owned by [code]node[/code] and [method pack] will therefore only save those two nodes, but not [code]collision[/code].
[codeblocks]
[gdscript]
# Create the objects.
var node = Node2D.new()
var body = RigidBody2D.new()
var collision = CollisionShape2D.new()

# Create the object hierarchy.
body.add_child(collision)
node.add_child(body)

# Change owner of `body`, but not of `collision`.
body.owner = node
var scene = PackedScene.new()

# Only `node` and `body` are now packed.
var result = scene.pack(node)
if result == OK:
    var error = ResourceSaver.save(scene, "res://path/name.tscn")  # Or "user://..."
    if error != OK:
        push_error("An error occurred while saving the scene to disk.")
[/gdscript]
[csharp]
// Create the objects.
var node = new Node2D();
var body = new RigidBody2D();
var collision = new CollisionShape2D();

// Create the object hierarchy.
body.AddChild(collision);
node.AddChild(body);

// Change owner of `body`, but not of `collision`.
body.Owner = node;
var scene = new PackedScene();

// Only `node` and `body` are now packed.
Error result = scene.Pack(node);
if (result == Error.Ok)
{
    Error error = ResourceSaver.Save(scene, "res://path/name.tscn"); // Or "user://..."
    if (error != Error.Ok)
    {
        GD.PushError("An error occurred while saving the scene to disk.");
    }
}
[/csharp]
[/codeblocks]

*/
type Simple [1]classdb.PackedScene
func (self Simple) Pack(path [1]classdb.Node) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Pack(path))
}
func (self Simple) Instantiate(edit_state classdb.PackedSceneGenEditState) [1]classdb.Node {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Node(Expert(self).Instantiate(gc, edit_state))
}
func (self Simple) CanInstantiate() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).CanInstantiate())
}
func (self Simple) GetState() [1]classdb.SceneState {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.SceneState(Expert(self).GetState(gc))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.PackedScene
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Packs the [param path] node, and all owned sub-nodes, into this [PackedScene]. Any existing data will be cleared. See [member Node.owner].
*/
//go:nosplit
func (self class) Pack(path object.Node) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path[0].AsPointer())[0])
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PackedScene.Bind_pack, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Instantiates the scene's node hierarchy. Triggers child scene instantiation(s). Triggers a [constant Node.NOTIFICATION_SCENE_INSTANTIATED] notification on the root node.
*/
//go:nosplit
func (self class) Instantiate(ctx gd.Lifetime, edit_state classdb.PackedSceneGenEditState) object.Node {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, edit_state)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PackedScene.Bind_instantiate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Node
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the scene file has nodes.
*/
//go:nosplit
func (self class) CanInstantiate() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PackedScene.Bind_can_instantiate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the [SceneState] representing the scene file contents.
*/
//go:nosplit
func (self class) GetState(ctx gd.Lifetime) object.SceneState {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PackedScene.Bind_get_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.SceneState
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsPackedScene() Expert { return self[0].AsPackedScene() }


//go:nosplit
func (self Simple) AsPackedScene() Simple { return self[0].AsPackedScene() }


//go:nosplit
func (self class) AsResource() Resource.Expert { return self[0].AsResource() }


//go:nosplit
func (self Simple) AsResource() Resource.Simple { return self[0].AsResource() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("PackedScene", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type GenEditState = classdb.PackedSceneGenEditState

const (
/*If passed to [method instantiate], blocks edits to the scene state.*/
	GenEditStateDisabled GenEditState = 0
/*If passed to [method instantiate], provides local scene resources to the local scene.
[b]Note:[/b] Only available in editor builds.*/
	GenEditStateInstance GenEditState = 1
/*If passed to [method instantiate], provides local scene resources to the local scene. Only the main scene should receive the main edit state.
[b]Note:[/b] Only available in editor builds.*/
	GenEditStateMain GenEditState = 2
/*It's similar to [constant GEN_EDIT_STATE_MAIN], but for the case where the scene is being instantiated to be the base of another one.
[b]Note:[/b] Only available in editor builds.*/
	GenEditStateMainInherited GenEditState = 3
)
