// Package CheckButton provides methods for working with CheckButton object instances.
package CheckButton

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/Packed"
import "graphics.gd/classdb/Button"
import "graphics.gd/classdb/BaseButton"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"

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
var _ = slices.Delete[[]struct{}, struct{}]

/*
[CheckButton] is a toggle button displayed as a check field. It's similar to [CheckBox] in functionality, but it has a different appearance. To follow established UX patterns, it's recommended to use [CheckButton] when toggling it has an [b]immediate[/b] effect on something. For example, it can be used when pressing it shows or hides advanced settings, without asking the user to confirm this action.
See also [BaseButton] which contains common properties and methods associated with this node.
*/
type Instance [1]gdclass.CheckButton

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsCheckButton() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.CheckButton

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("CheckButton"))
	casted := Instance{*(*gdclass.CheckButton)(unsafe.Pointer(&object))}
	return casted
}

func (self class) AsCheckButton() Advanced      { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsCheckButton() Instance   { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsButton() Button.Advanced    { return *((*Button.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsButton() Button.Instance { return *((*Button.Instance)(unsafe.Pointer(&self))) }
func (self class) AsBaseButton() BaseButton.Advanced {
	return *((*BaseButton.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsBaseButton() BaseButton.Instance {
	return *((*BaseButton.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(Button.Advanced(self.AsButton()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Button.Instance(self.AsButton()), name)
	}
}
func init() {
	gdclass.Register("CheckButton", func(ptr gd.Object) any { return [1]gdclass.CheckButton{*(*gdclass.CheckButton)(unsafe.Pointer(&ptr))} })
}
