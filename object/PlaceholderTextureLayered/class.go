package PlaceholderTextureLayered

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/TextureLayered"
import "grow.graphics/gd/object/Texture"
import "grow.graphics/gd/object/Resource"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This class is used when loading a project that uses a [TextureLayered] subclass in 2 conditions:
- When running the project exported in dedicated server mode, only the texture's dimensions are kept (as they may be relied upon for gameplay purposes or positioning of other elements). This allows reducing the exported PCK's size significantly.
- When this subclass is missing due to using a different engine version or build (e.g. modules disabled).
[b]Note:[/b] This is not intended to be used as an actual texture for rendering. It is not guaranteed to work like one in shaders or materials (for example when calculating UV).

*/
type Simple [1]classdb.PlaceholderTextureLayered
func (self Simple) SetSize(size gd.Vector2i) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSize(size)
}
func (self Simple) GetSize() gd.Vector2i {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2i(Expert(self).GetSize())
}
func (self Simple) SetLayers(layers int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLayers(gd.Int(layers))
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.PlaceholderTextureLayered
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

//go:nosplit
func (self class) SetSize(size gd.Vector2i)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PlaceholderTextureLayered.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetSize() gd.Vector2i {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2i](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PlaceholderTextureLayered.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLayers(layers gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, layers)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PlaceholderTextureLayered.Bind_set_layers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsPlaceholderTextureLayered() Expert { return self[0].AsPlaceholderTextureLayered() }


//go:nosplit
func (self Simple) AsPlaceholderTextureLayered() Simple { return self[0].AsPlaceholderTextureLayered() }


//go:nosplit
func (self class) AsTextureLayered() TextureLayered.Expert { return self[0].AsTextureLayered() }


//go:nosplit
func (self Simple) AsTextureLayered() TextureLayered.Simple { return self[0].AsTextureLayered() }


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
func init() {classdb.Register("PlaceholderTextureLayered", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
