// Package EditorResourcePicker provides methods for working with EditorResourcePicker object instances.
package EditorResourcePicker

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/classdb/HBoxContainer"
import "graphics.gd/classdb/BoxContainer"
import "graphics.gd/classdb/Container"
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
This [Control] node is used in the editor's Inspector dock to allow editing of [Resource] type properties. It provides options for creating, loading, saving and converting resources. Can be used with [EditorInspectorPlugin] to recreate the same behavior.
[b]Note:[/b] This [Control] does not include any editor for the resource, as editing is controlled by the Inspector dock itself or sub-Inspectors.

	// EditorResourcePicker methods that can be overridden by a [Class] that extends it.
	type EditorResourcePicker interface {
		//This virtual method is called when updating the context menu of [EditorResourcePicker]. Implement this method to override the "New ..." items with your own options. [param menu_node] is a reference to the [PopupMenu] node.
		//[b]Note:[/b] Implement [method _handle_menu_selected] to handle these custom items.
		SetCreateOptions(menu_node Object.Instance)
		//This virtual method can be implemented to handle context menu items not handled by default. See [method _set_create_options].
		HandleMenuSelected(id int) bool
	}
*/
type Instance [1]gdclass.EditorResourcePicker
type Any interface {
	gd.IsClass
	AsEditorResourcePicker() Instance
}

/*
This virtual method is called when updating the context menu of [EditorResourcePicker]. Implement this method to override the "New ..." items with your own options. [param menu_node] is a reference to the [PopupMenu] node.
[b]Note:[/b] Implement [method _handle_menu_selected] to handle these custom items.
*/
func (Instance) _set_create_options(impl func(ptr unsafe.Pointer, menu_node Object.Instance)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var menu_node = [1]gd.Object{pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(menu_node[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, menu_node)
	}
}

/*
This virtual method can be implemented to handle context menu items not handled by default. See [method _set_create_options].
*/
func (Instance) _handle_menu_selected(impl func(ptr unsafe.Pointer, id int) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var id = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(id))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns a list of all allowed types and subtypes corresponding to the [member base_type]. If the [member base_type] is empty, an empty list is returned.
*/
func (self Instance) GetAllowedTypes() []string {
	return []string(class(self).GetAllowedTypes().Strings())
}

/*
Sets the toggle mode state for the main button. Works only if [member toggle_mode] is set to [code]true[/code].
*/
func (self Instance) SetTogglePressed(pressed bool) {
	class(self).SetTogglePressed(pressed)
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorResourcePicker

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorResourcePicker"))
	return Instance{*(*gdclass.EditorResourcePicker)(unsafe.Pointer(&object))}
}

func (self Instance) BaseType() string {
	return string(class(self).GetBaseType().String())
}

func (self Instance) SetBaseType(value string) {
	class(self).SetBaseType(gd.NewString(value))
}

func (self Instance) EditedResource() [1]gdclass.Resource {
	return [1]gdclass.Resource(class(self).GetEditedResource())
}

func (self Instance) SetEditedResource(value [1]gdclass.Resource) {
	class(self).SetEditedResource(value)
}

func (self Instance) Editable() bool {
	return bool(class(self).IsEditable())
}

func (self Instance) SetEditable(value bool) {
	class(self).SetEditable(value)
}

func (self Instance) ToggleMode() bool {
	return bool(class(self).IsToggleMode())
}

func (self Instance) SetToggleMode(value bool) {
	class(self).SetToggleMode(value)
}

/*
This virtual method is called when updating the context menu of [EditorResourcePicker]. Implement this method to override the "New ..." items with your own options. [param menu_node] is a reference to the [PopupMenu] node.
[b]Note:[/b] Implement [method _handle_menu_selected] to handle these custom items.
*/
func (class) _set_create_options(impl func(ptr unsafe.Pointer, menu_node [1]gd.Object)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var menu_node = [1]gd.Object{pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(menu_node[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, menu_node)
	}
}

/*
This virtual method can be implemented to handle context menu items not handled by default. See [method _set_create_options].
*/
func (class) _handle_menu_selected(impl func(ptr unsafe.Pointer, id gd.Int) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var id = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, id)
		gd.UnsafeSet(p_back, ret)
	}
}

//go:nosplit
func (self class) SetBaseType(base_type gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(base_type))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorResourcePicker.Bind_set_base_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetBaseType() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorResourcePicker.Bind_get_base_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns a list of all allowed types and subtypes corresponding to the [member base_type]. If the [member base_type] is empty, an empty list is returned.
*/
//go:nosplit
func (self class) GetAllowedTypes() gd.PackedStringArray {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorResourcePicker.Bind_get_allowed_types, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.PackedStringArray](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEditedResource(resource [1]gdclass.Resource) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(resource[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorResourcePicker.Bind_set_edited_resource, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetEditedResource() [1]gdclass.Resource {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorResourcePicker.Bind_get_edited_resource, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = [1]gdclass.Resource{gd.PointerWithOwnershipTransferredToGo[gdclass.Resource](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetToggleMode(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorResourcePicker.Bind_set_toggle_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsToggleMode() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorResourcePicker.Bind_is_toggle_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the toggle mode state for the main button. Works only if [member toggle_mode] is set to [code]true[/code].
*/
//go:nosplit
func (self class) SetTogglePressed(pressed bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pressed)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorResourcePicker.Bind_set_toggle_pressed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) SetEditable(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorResourcePicker.Bind_set_editable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsEditable() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorResourcePicker.Bind_is_editable, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnResourceSelected(cb func(resource [1]gdclass.Resource, inspect bool)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("resource_selected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnResourceChanged(cb func(resource [1]gdclass.Resource)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("resource_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsEditorResourcePicker() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorResourcePicker() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsHBoxContainer() HBoxContainer.Advanced {
	return *((*HBoxContainer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsHBoxContainer() HBoxContainer.Instance {
	return *((*HBoxContainer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsBoxContainer() BoxContainer.Advanced {
	return *((*BoxContainer.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsBoxContainer() BoxContainer.Instance {
	return *((*BoxContainer.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsContainer() Container.Advanced {
	return *((*Container.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsContainer() Container.Instance {
	return *((*Container.Instance)(unsafe.Pointer(&self)))
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
	case "_set_create_options":
		return reflect.ValueOf(self._set_create_options)
	case "_handle_menu_selected":
		return reflect.ValueOf(self._handle_menu_selected)
	default:
		return gd.VirtualByName(HBoxContainer.Advanced(self.AsHBoxContainer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_set_create_options":
		return reflect.ValueOf(self._set_create_options)
	case "_handle_menu_selected":
		return reflect.ValueOf(self._handle_menu_selected)
	default:
		return gd.VirtualByName(HBoxContainer.Instance(self.AsHBoxContainer()), name)
	}
}
func init() {
	gdclass.Register("EditorResourcePicker", func(ptr gd.Object) any {
		return [1]gdclass.EditorResourcePicker{*(*gdclass.EditorResourcePicker)(unsafe.Pointer(&ptr))}
	})
}
