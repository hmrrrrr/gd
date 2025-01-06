// Package ProgressBar provides methods for working with ProgressBar object instances.
package ProgressBar

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/Range"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Node"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
A control used for visual representation of a percentage. Shows fill percentage from right to left.
*/
type Instance [1]gdclass.ProgressBar
type Any interface {
	gd.IsClass
	AsProgressBar() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ProgressBar

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ProgressBar"))
	casted := Instance{*(*gdclass.ProgressBar)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) FillMode() int {
	return int(int(class(self).GetFillMode()))
}

func (self Instance) SetFillMode(value int) {
	class(self).SetFillMode(gd.Int(value))
}

func (self Instance) ShowPercentage() bool {
	return bool(class(self).IsPercentageShown())
}

func (self Instance) SetShowPercentage(value bool) {
	class(self).SetShowPercentage(value)
}

func (self Instance) Indeterminate() bool {
	return bool(class(self).IsIndeterminate())
}

func (self Instance) SetIndeterminate(value bool) {
	class(self).SetIndeterminate(value)
}

func (self Instance) EditorPreviewIndeterminate() bool {
	return bool(class(self).IsEditorPreviewIndeterminateEnabled())
}

func (self Instance) SetEditorPreviewIndeterminate(value bool) {
	class(self).SetEditorPreviewIndeterminate(value)
}

//go:nosplit
func (self class) SetFillMode(mode gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProgressBar.Bind_set_fill_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFillMode() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProgressBar.Bind_get_fill_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetShowPercentage(visible bool) {
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProgressBar.Bind_set_show_percentage, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsPercentageShown() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProgressBar.Bind_is_percentage_shown, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetIndeterminate(indeterminate bool) {
	var frame = callframe.New()
	callframe.Arg(frame, indeterminate)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProgressBar.Bind_set_indeterminate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsIndeterminate() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProgressBar.Bind_is_indeterminate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEditorPreviewIndeterminate(preview_indeterminate bool) {
	var frame = callframe.New()
	callframe.Arg(frame, preview_indeterminate)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProgressBar.Bind_set_editor_preview_indeterminate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsEditorPreviewIndeterminateEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ProgressBar.Bind_is_editor_preview_indeterminate_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsProgressBar() Advanced     { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsProgressBar() Instance  { return *((*Instance)(unsafe.Pointer(&self))) }
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
		return gd.VirtualByName(Range.Advanced(self.AsRange()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Range.Instance(self.AsRange()), name)
	}
}
func init() {
	gdclass.Register("ProgressBar", func(ptr gd.Object) any { return [1]gdclass.ProgressBar{*(*gdclass.ProgressBar)(unsafe.Pointer(&ptr))} })
}

type FillMode = gdclass.ProgressBarFillMode

const (
	/*The progress bar fills from begin to end horizontally, according to the language direction. If [method Control.is_layout_rtl] returns [code]false[/code], it fills from left to right, and if it returns [code]true[/code], it fills from right to left.*/
	FillBeginToEnd FillMode = 0
	/*The progress bar fills from end to begin horizontally, according to the language direction. If [method Control.is_layout_rtl] returns [code]false[/code], it fills from right to left, and if it returns [code]true[/code], it fills from left to right.*/
	FillEndToBegin FillMode = 1
	/*The progress fills from top to bottom.*/
	FillTopToBottom FillMode = 2
	/*The progress fills from bottom to top.*/
	FillBottomToTop FillMode = 3
)
