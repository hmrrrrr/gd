package OpenXRCompositionLayerQuad

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/OpenXRCompositionLayer"
import "grow.graphics/gd/gdclass/Node3D"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
An OpenXR composition layer that allows rendering a [SubViewport] on a quad.
*/
type Instance [1]classdb.OpenXRCompositionLayerQuad

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.OpenXRCompositionLayerQuad

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OpenXRCompositionLayerQuad"))
	return Instance{classdb.OpenXRCompositionLayerQuad(object)}
}

func (self Instance) QuadSize() gd.Vector2 {
	return gd.Vector2(class(self).GetQuadSize())
}

func (self Instance) SetQuadSize(value gd.Vector2) {
	class(self).SetQuadSize(value)
}

//go:nosplit
func (self class) SetQuadSize(size gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayerQuad.Bind_set_quad_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetQuadSize() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRCompositionLayerQuad.Bind_get_quad_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsOpenXRCompositionLayerQuad() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsOpenXRCompositionLayerQuad() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsOpenXRCompositionLayer() OpenXRCompositionLayer.Advanced {
	return *((*OpenXRCompositionLayer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsOpenXRCompositionLayer() OpenXRCompositionLayer.Instance {
	return *((*OpenXRCompositionLayer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced    { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode3D() Node3D.Instance { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced        { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance     { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsOpenXRCompositionLayer(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsOpenXRCompositionLayer(), name)
	}
}
func init() {
	classdb.Register("OpenXRCompositionLayerQuad", func(ptr gd.Object) any { return classdb.OpenXRCompositionLayerQuad(ptr) })
}
