package Texture3DRD

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/Texture3D"
import "grow.graphics/gd/object/Texture"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This texture class allows you to use a 3D texture created directly on the [RenderingDevice] as a texture for materials, meshes, etc.

*/
type Simple [1]classdb.Texture3DRD
func (self Simple) SetTextureRdRid(texture_rd_rid gd.RID) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTextureRdRid(texture_rd_rid)
}
func (self Simple) GetTextureRdRid() gd.RID {
	gc := gd.GarbageCollector(); _ = gc
	return gd.RID(Expert(self).GetTextureRdRid())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.Texture3DRD
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetTextureRdRid(texture_rd_rid gd.RID)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, texture_rd_rid)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Texture3DRD.Bind_set_texture_rd_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTextureRdRid() gd.RID {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Texture3DRD.Bind_get_texture_rd_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsTexture3DRD() Expert { return self[0].AsTexture3DRD() }


//go:nosplit
func (self Simple) AsTexture3DRD() Simple { return self[0].AsTexture3DRD() }


//go:nosplit
func (self class) AsTexture3D() Texture3D.Expert { return self[0].AsTexture3D() }


//go:nosplit
func (self Simple) AsTexture3D() Texture3D.Simple { return self[0].AsTexture3D() }


//go:nosplit
func (self class) AsTexture() Texture.Expert { return self[0].AsTexture() }


//go:nosplit
func (self Simple) AsTexture() Texture.Simple { return self[0].AsTexture() }


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
func init() {classdb.Register("Texture3DRD", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
