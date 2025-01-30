// Package VisibleOnScreenEnabler2D provides methods for working with VisibleOnScreenEnabler2D object instances.
package VisibleOnScreenEnabler2D

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"
import "graphics.gd/classdb/Node2D"
import "graphics.gd/classdb/VisibleOnScreenNotifier2D"
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
[VisibleOnScreenEnabler2D] contains a rectangular region of 2D space and a target node. The target node will be automatically enabled (via its [member Node.process_mode] property) when any part of this region becomes visible on the screen, and automatically disabled otherwise. This can for example be used to activate enemies only when the player approaches them.
See [VisibleOnScreenNotifier2D] if you only want to be notified when the region is visible on screen.
[b]Note:[/b] [VisibleOnScreenEnabler2D] uses the render culling code to determine whether it's visible on screen, so it won't function unless [member CanvasItem.visible] is set to [code]true[/code].
*/
type Instance [1]gdclass.VisibleOnScreenEnabler2D

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVisibleOnScreenEnabler2D() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VisibleOnScreenEnabler2D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VisibleOnScreenEnabler2D"))
	casted := Instance{*(*gdclass.VisibleOnScreenEnabler2D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) EnableMode() gdclass.VisibleOnScreenEnabler2DEnableMode {
	return gdclass.VisibleOnScreenEnabler2DEnableMode(class(self).GetEnableMode())
}

func (self Instance) SetEnableMode(value gdclass.VisibleOnScreenEnabler2DEnableMode) {
	class(self).SetEnableMode(value)
}

func (self Instance) EnableNodePath() string {
	return string(class(self).GetEnableNodePath().String())
}

func (self Instance) SetEnableNodePath(value string) {
	class(self).SetEnableNodePath(Path.ToNode(String.New(value)))
}

//go:nosplit
func (self class) SetEnableMode(mode gdclass.VisibleOnScreenEnabler2DEnableMode) { //gd:VisibleOnScreenEnabler2D.set_enable_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisibleOnScreenEnabler2D.Bind_set_enable_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnableMode() gdclass.VisibleOnScreenEnabler2DEnableMode { //gd:VisibleOnScreenEnabler2D.get_enable_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.VisibleOnScreenEnabler2DEnableMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisibleOnScreenEnabler2D.Bind_get_enable_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEnableNodePath(path Path.ToNode) { //gd:VisibleOnScreenEnabler2D.set_enable_node_path
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(path)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisibleOnScreenEnabler2D.Bind_set_enable_node_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnableNodePath() Path.ToNode { //gd:VisibleOnScreenEnabler2D.get_enable_node_path
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VisibleOnScreenEnabler2D.Bind_get_enable_node_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Path.ToNode(String.Via(gd.NodePathProxy{}, pointers.Pack(pointers.New[gd.NodePath](r_ret.Get()))))
	frame.Free()
	return ret
}
func (self class) AsVisibleOnScreenEnabler2D() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVisibleOnScreenEnabler2D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisibleOnScreenNotifier2D() VisibleOnScreenNotifier2D.Advanced {
	return *((*VisibleOnScreenNotifier2D.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsVisibleOnScreenNotifier2D() VisibleOnScreenNotifier2D.Instance {
	return *((*VisibleOnScreenNotifier2D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode2D() Node2D.Advanced    { return *((*Node2D.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode2D() Node2D.Instance { return *((*Node2D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisibleOnScreenNotifier2D.Advanced(self.AsVisibleOnScreenNotifier2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(VisibleOnScreenNotifier2D.Instance(self.AsVisibleOnScreenNotifier2D()), name)
	}
}
func init() {
	gdclass.Register("VisibleOnScreenEnabler2D", func(ptr gd.Object) any {
		return [1]gdclass.VisibleOnScreenEnabler2D{*(*gdclass.VisibleOnScreenEnabler2D)(unsafe.Pointer(&ptr))}
	})
}

type EnableMode = gdclass.VisibleOnScreenEnabler2DEnableMode //gd:VisibleOnScreenEnabler2D.EnableMode

const (
	/*Corresponds to [constant Node.PROCESS_MODE_INHERIT].*/
	EnableModeInherit EnableMode = 0
	/*Corresponds to [constant Node.PROCESS_MODE_ALWAYS].*/
	EnableModeAlways EnableMode = 1
	/*Corresponds to [constant Node.PROCESS_MODE_WHEN_PAUSED].*/
	EnableModeWhenPaused EnableMode = 2
)
