package BoxMesh

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/PrimitiveMesh"
import "grow.graphics/gd/objects/Mesh"
import "grow.graphics/gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Generate an axis-aligned box [PrimitiveMesh].
The box's UV layout is arranged in a 3×2 layout that allows texturing each face individually. To apply the same texture on all faces, change the material's UV property to [code]Vector3(3, 2, 1)[/code]. This is equivalent to adding [code]UV *= vec2(3.0, 2.0)[/code] in a vertex shader.
[b]Note:[/b] When using a large textured [BoxMesh] (e.g. as a floor), you may stumble upon UV jittering issues depending on the camera angle. To solve this, increase [member subdivide_depth], [member subdivide_height] and [member subdivide_width] until you no longer notice UV jittering.
*/
type Instance [1]classdb.BoxMesh

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.BoxMesh

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("BoxMesh"))
	return Instance{classdb.BoxMesh(object)}
}

func (self Instance) Size() gd.Vector3 {
	return gd.Vector3(class(self).GetSize())
}

func (self Instance) SetSize(value gd.Vector3) {
	class(self).SetSize(value)
}

func (self Instance) SubdivideWidth() int {
	return int(int(class(self).GetSubdivideWidth()))
}

func (self Instance) SetSubdivideWidth(value int) {
	class(self).SetSubdivideWidth(gd.Int(value))
}

func (self Instance) SubdivideHeight() int {
	return int(int(class(self).GetSubdivideHeight()))
}

func (self Instance) SetSubdivideHeight(value int) {
	class(self).SetSubdivideHeight(gd.Int(value))
}

func (self Instance) SubdivideDepth() int {
	return int(int(class(self).GetSubdivideDepth()))
}

func (self Instance) SetSubdivideDepth(value int) {
	class(self).SetSubdivideDepth(gd.Int(value))
}

//go:nosplit
func (self class) SetSize(size gd.Vector3) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoxMesh.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSize() gd.Vector3 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector3](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoxMesh.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSubdivideWidth(subdivide gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, subdivide)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoxMesh.Bind_set_subdivide_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSubdivideWidth() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoxMesh.Bind_get_subdivide_width, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSubdivideHeight(divisions gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, divisions)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoxMesh.Bind_set_subdivide_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSubdivideHeight() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoxMesh.Bind_get_subdivide_height, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSubdivideDepth(divisions gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, divisions)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoxMesh.Bind_set_subdivide_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSubdivideDepth() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.BoxMesh.Bind_get_subdivide_depth, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsBoxMesh() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsBoxMesh() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsPrimitiveMesh() PrimitiveMesh.Advanced {
	return *((*PrimitiveMesh.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsPrimitiveMesh() PrimitiveMesh.Instance {
	return *((*PrimitiveMesh.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsMesh() Mesh.Advanced    { return *((*Mesh.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMesh() Mesh.Instance { return *((*Mesh.Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() gd.RefCounted    { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Instance) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsPrimitiveMesh(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsPrimitiveMesh(), name)
	}
}
func init() { classdb.Register("BoxMesh", func(ptr gd.Object) any { return classdb.BoxMesh(ptr) }) }
