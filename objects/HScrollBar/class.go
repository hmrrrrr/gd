package HScrollBar

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/ScrollBar"
import "graphics.gd/objects/Range"
import "graphics.gd/objects/Control"
import "graphics.gd/objects/CanvasItem"
import "graphics.gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
A horizontal scrollbar, typically used to navigate through content that extends beyond the visible width of a control. It is a [Range]-based control and goes from left (min) to right (max).
*/
type Instance [1]classdb.HScrollBar
type Any interface {
	gd.IsClass
	AsHScrollBar() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.HScrollBar

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("HScrollBar"))
	return Instance{classdb.HScrollBar(object)}
}

func (self class) AsHScrollBar() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsHScrollBar() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsScrollBar() ScrollBar.Advanced {
	return *((*ScrollBar.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsScrollBar() ScrollBar.Instance {
	return *((*ScrollBar.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRange() Range.Advanced     { return *((*Range.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRange() Range.Instance  { return *((*Range.Instance)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.Advanced { return *((*Control.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsControl() Control.Instance {
	return *((*Control.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(self.AsScrollBar(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsScrollBar(), name)
	}
}
func init() {
	classdb.Register("HScrollBar", func(ptr gd.Object) any { return classdb.HScrollBar(ptr) })
}
