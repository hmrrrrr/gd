package AnimationTree

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/AnimationMixer"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A node used for advanced animation transitions in an [AnimationPlayer].
[b]Note:[/b] When linked with an [AnimationPlayer], several properties and methods of the corresponding [AnimationPlayer] will not function as expected. Playback and transitions should be handled using only the [AnimationTree] and its constituent [AnimationNode](s). The [AnimationPlayer] node should be used solely for adding, deleting, and editing animations.

*/
type Simple [1]classdb.AnimationTree
func (self Simple) SetTreeRoot(animation_node [1]classdb.AnimationRootNode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTreeRoot(animation_node)
}
func (self Simple) GetTreeRoot() [1]classdb.AnimationRootNode {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.AnimationRootNode(Expert(self).GetTreeRoot(gc))
}
func (self Simple) SetAdvanceExpressionBaseNode(path gd.NodePath) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAdvanceExpressionBaseNode(path)
}
func (self Simple) GetAdvanceExpressionBaseNode() gd.NodePath {
	gc := gd.GarbageCollector(); _ = gc
	return gd.NodePath(Expert(self).GetAdvanceExpressionBaseNode(gc))
}
func (self Simple) SetAnimationPlayer(path gd.NodePath) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAnimationPlayer(path)
}
func (self Simple) GetAnimationPlayer() gd.NodePath {
	gc := gd.GarbageCollector(); _ = gc
	return gd.NodePath(Expert(self).GetAnimationPlayer(gc))
}
func (self Simple) SetProcessCallback(mode classdb.AnimationTreeAnimationProcessCallback) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetProcessCallback(mode)
}
func (self Simple) GetProcessCallback() classdb.AnimationTreeAnimationProcessCallback {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.AnimationTreeAnimationProcessCallback(Expert(self).GetProcessCallback())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AnimationTree
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetTreeRoot(animation_node object.AnimationRootNode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(animation_node[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationTree.Bind_set_tree_root, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTreeRoot(ctx gd.Lifetime) object.AnimationRootNode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationTree.Bind_get_tree_root, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.AnimationRootNode
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAdvanceExpressionBaseNode(path gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationTree.Bind_set_advance_expression_base_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAdvanceExpressionBaseNode(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationTree.Bind_get_advance_expression_base_node, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAnimationPlayer(path gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationTree.Bind_set_animation_player, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetAnimationPlayer(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationTree.Bind_get_animation_player, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the process notification in which to update animations.
*/
//go:nosplit
func (self class) SetProcessCallback(mode classdb.AnimationTreeAnimationProcessCallback)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationTree.Bind_set_process_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the process notification in which to update animations.
*/
//go:nosplit
func (self class) GetProcessCallback() classdb.AnimationTreeAnimationProcessCallback {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.AnimationTreeAnimationProcessCallback](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.AnimationTree.Bind_get_process_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsAnimationTree() Expert { return self[0].AsAnimationTree() }


//go:nosplit
func (self Simple) AsAnimationTree() Simple { return self[0].AsAnimationTree() }


//go:nosplit
func (self class) AsAnimationMixer() AnimationMixer.Expert { return self[0].AsAnimationMixer() }


//go:nosplit
func (self Simple) AsAnimationMixer() AnimationMixer.Simple { return self[0].AsAnimationMixer() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

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
func init() {classdb.Register("AnimationTree", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type AnimationProcessCallback = classdb.AnimationTreeAnimationProcessCallback

const (
	AnimationProcessPhysics AnimationProcessCallback = 0
	AnimationProcessIdle AnimationProcessCallback = 1
	AnimationProcessManual AnimationProcessCallback = 2
)
