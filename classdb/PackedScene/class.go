// Package PackedScene provides methods for working with PackedScene object instances.
package PackedScene

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
type Instance [1]gdclass.PackedScene

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPackedScene() Instance
}

/*
Packs the [param path] node, and all owned sub-nodes, into this [PackedScene]. Any existing data will be cleared. See [member Node.owner].
*/
func (self Instance) Pack(path [1]gdclass.Node) error { //gd:PackedScene.pack
	return error(gd.ToError(class(self).Pack(path)))
}

/*
Instantiates the scene's node hierarchy. Triggers child scene instantiation(s). Triggers a [constant Node.NOTIFICATION_SCENE_INSTANTIATED] notification on the root node.
*/
func (self Instance) Instantiate() [1]gdclass.Node { //gd:PackedScene.instantiate
	return [1]gdclass.Node(class(self).Instantiate(0))
}

/*
Returns [code]true[/code] if the scene file has nodes.
*/
func (self Instance) CanInstantiate() bool { //gd:PackedScene.can_instantiate
	return bool(class(self).CanInstantiate())
}

/*
Returns the [SceneState] representing the scene file contents.
*/
func (self Instance) GetState() [1]gdclass.SceneState { //gd:PackedScene.get_state
	return [1]gdclass.SceneState(class(self).GetState())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PackedScene

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PackedScene"))
	casted := Instance{*(*gdclass.PackedScene)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Packs the [param path] node, and all owned sub-nodes, into this [PackedScene]. Any existing data will be cleared. See [member Node.owner].
*/
//go:nosplit
func (self class) Pack(path [1]gdclass.Node) Error.Code { //gd:PackedScene.pack
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(path[0])[0])
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PackedScene.Bind_pack, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Instantiates the scene's node hierarchy. Triggers child scene instantiation(s). Triggers a [constant Node.NOTIFICATION_SCENE_INSTANTIATED] notification on the root node.
*/
//go:nosplit
func (self class) Instantiate(edit_state gdclass.PackedSceneGenEditState) [1]gdclass.Node { //gd:PackedScene.instantiate
	var frame = callframe.New()
	callframe.Arg(frame, edit_state)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PackedScene.Bind_instantiate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Node{gd.PointerWithOwnershipTransferredToGo[gdclass.Node](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the scene file has nodes.
*/
//go:nosplit
func (self class) CanInstantiate() bool { //gd:PackedScene.can_instantiate
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PackedScene.Bind_can_instantiate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [SceneState] representing the scene file contents.
*/
//go:nosplit
func (self class) GetState() [1]gdclass.SceneState { //gd:PackedScene.get_state
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PackedScene.Bind_get_state, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.SceneState{gd.PointerWithOwnershipTransferredToGo[gdclass.SceneState](r_ret.Get())}
	frame.Free()
	return ret
}
func (self class) AsPackedScene() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPackedScene() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("PackedScene", func(ptr gd.Object) any { return [1]gdclass.PackedScene{*(*gdclass.PackedScene)(unsafe.Pointer(&ptr))} })
}

type GenEditState = gdclass.PackedSceneGenEditState //gd:PackedScene.GenEditState

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
