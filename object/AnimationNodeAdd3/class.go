package AnimationNodeAdd3

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/AnimationNodeSync"
import "grow.graphics/gd/object/AnimationNode"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A resource to add to an [AnimationNodeBlendTree]. Blends two animations out of three additively out of three based on the amount value.
This animation node has three inputs:
- The base animation to add to
- A "-add" animation to blend with when the blend amount is negative
- A "+add" animation to blend with when the blend amount is positive
If the absolute value of the amount is greater than [code]1.0[/code], the animation connected to "in" port is blended with the amplified animation connected to "-add"/"+add" port.

*/
type Simple [1]classdb.AnimationNodeAdd3
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.AnimationNodeAdd3
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsAnimationNodeAdd3() Expert { return self[0].AsAnimationNodeAdd3() }


//go:nosplit
func (self Simple) AsAnimationNodeAdd3() Simple { return self[0].AsAnimationNodeAdd3() }


//go:nosplit
func (self class) AsAnimationNodeSync() AnimationNodeSync.Expert { return self[0].AsAnimationNodeSync() }


//go:nosplit
func (self Simple) AsAnimationNodeSync() AnimationNodeSync.Simple { return self[0].AsAnimationNodeSync() }


//go:nosplit
func (self class) AsAnimationNode() AnimationNode.Expert { return self[0].AsAnimationNode() }


//go:nosplit
func (self Simple) AsAnimationNode() AnimationNode.Simple { return self[0].AsAnimationNode() }


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
func init() {classdb.Register("AnimationNodeAdd3", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
