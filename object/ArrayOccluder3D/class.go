package ArrayOccluder3D

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Occluder3D"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
[ArrayOccluder3D] stores an arbitrary 3D polygon shape that can be used by the engine's occlusion culling system. This is analogous to [ArrayMesh], but for occluders.
See [OccluderInstance3D]'s documentation for instructions on setting up occlusion culling.

*/
type Simple [1]classdb.ArrayOccluder3D
func (self Simple) SetArrays(vertices gd.PackedVector3Array, indices gd.PackedInt32Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetArrays(vertices, indices)
}
func (self Simple) SetVertices(vertices gd.PackedVector3Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVertices(vertices)
}
func (self Simple) SetIndices(indices gd.PackedInt32Array) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetIndices(indices)
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.ArrayOccluder3D
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Sets [member indices] and [member vertices], while updating the final occluder only once after both values are set.
*/
//go:nosplit
func (self class) SetArrays(vertices gd.PackedVector3Array, indices gd.PackedInt32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(vertices))
	callframe.Arg(frame, mmm.Get(indices))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ArrayOccluder3D.Bind_set_arrays, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetVertices(vertices gd.PackedVector3Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(vertices))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ArrayOccluder3D.Bind_set_vertices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetIndices(indices gd.PackedInt32Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(indices))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ArrayOccluder3D.Bind_set_indices, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsArrayOccluder3D() Expert { return self[0].AsArrayOccluder3D() }


//go:nosplit
func (self Simple) AsArrayOccluder3D() Simple { return self[0].AsArrayOccluder3D() }


//go:nosplit
func (self class) AsOccluder3D() Occluder3D.Expert { return self[0].AsOccluder3D() }


//go:nosplit
func (self Simple) AsOccluder3D() Occluder3D.Simple { return self[0].AsOccluder3D() }


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
func init() {classdb.Register("ArrayOccluder3D", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
