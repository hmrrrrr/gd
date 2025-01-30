// Package TabContainer provides methods for working with TabContainer object instances.
package TabContainer

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Container"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/Node"
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
import "graphics.gd/variant/Vector2"

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
Arranges child controls into a tabbed view, creating a tab for each one. The active tab's corresponding control is made visible, while all other child controls are hidden. Ignores non-control children.
[b]Note:[/b] The drawing of the clickable tabs is handled by this node; [TabBar] is not needed.
*/
type Instance [1]gdclass.TabContainer

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsTabContainer() Instance
}

/*
Returns the number of tabs.
*/
func (self Instance) GetTabCount() int { //gd:TabContainer.get_tab_count
	return int(int(class(self).GetTabCount()))
}

/*
Returns the previously active tab index.
*/
func (self Instance) GetPreviousTab() int { //gd:TabContainer.get_previous_tab
	return int(int(class(self).GetPreviousTab()))
}

/*
Selects the first available tab with lower index than the currently selected. Returns [code]true[/code] if tab selection changed.
*/
func (self Instance) SelectPreviousAvailable() bool { //gd:TabContainer.select_previous_available
	return bool(class(self).SelectPreviousAvailable())
}

/*
Selects the first available tab with greater index than the currently selected. Returns [code]true[/code] if tab selection changed.
*/
func (self Instance) SelectNextAvailable() bool { //gd:TabContainer.select_next_available
	return bool(class(self).SelectNextAvailable())
}

/*
Returns the child [Control] node located at the active tab index.
*/
func (self Instance) GetCurrentTabControl() [1]gdclass.Control { //gd:TabContainer.get_current_tab_control
	return [1]gdclass.Control(class(self).GetCurrentTabControl())
}

/*
Returns the [TabBar] contained in this container.
[b]Warning:[/b] This is a required internal node, removing and freeing it or editing its tabs may cause a crash. If you wish to edit the tabs, use the methods provided in [TabContainer].
*/
func (self Instance) GetTabBar() [1]gdclass.TabBar { //gd:TabContainer.get_tab_bar
	return [1]gdclass.TabBar(class(self).GetTabBar())
}

/*
Returns the [Control] node from the tab at index [param tab_idx].
*/
func (self Instance) GetTabControl(tab_idx int) [1]gdclass.Control { //gd:TabContainer.get_tab_control
	return [1]gdclass.Control(class(self).GetTabControl(int64(tab_idx)))
}

/*
Sets a custom title for the tab at index [param tab_idx] (tab titles default to the name of the indexed child node). Set it back to the child's name to make the tab default to it again.
*/
func (self Instance) SetTabTitle(tab_idx int, title string) { //gd:TabContainer.set_tab_title
	class(self).SetTabTitle(int64(tab_idx), String.New(title))
}

/*
Returns the title of the tab at index [param tab_idx]. Tab titles default to the name of the indexed child node, but this can be overridden with [method set_tab_title].
*/
func (self Instance) GetTabTitle(tab_idx int) string { //gd:TabContainer.get_tab_title
	return string(class(self).GetTabTitle(int64(tab_idx)).String())
}

/*
Sets a custom tooltip text for tab at index [param tab_idx].
[b]Note:[/b] By default, if the [param tooltip] is empty and the tab text is truncated (not all characters fit into the tab), the title will be displayed as a tooltip. To hide the tooltip, assign [code]" "[/code] as the [param tooltip] text.
*/
func (self Instance) SetTabTooltip(tab_idx int, tooltip string) { //gd:TabContainer.set_tab_tooltip
	class(self).SetTabTooltip(int64(tab_idx), String.New(tooltip))
}

/*
Returns the tooltip text of the tab at index [param tab_idx].
*/
func (self Instance) GetTabTooltip(tab_idx int) string { //gd:TabContainer.get_tab_tooltip
	return string(class(self).GetTabTooltip(int64(tab_idx)).String())
}

/*
Sets an icon for the tab at index [param tab_idx].
*/
func (self Instance) SetTabIcon(tab_idx int, icon [1]gdclass.Texture2D) { //gd:TabContainer.set_tab_icon
	class(self).SetTabIcon(int64(tab_idx), icon)
}

/*
Returns the [Texture2D] for the tab at index [param tab_idx] or [code]null[/code] if the tab has no [Texture2D].
*/
func (self Instance) GetTabIcon(tab_idx int) [1]gdclass.Texture2D { //gd:TabContainer.get_tab_icon
	return [1]gdclass.Texture2D(class(self).GetTabIcon(int64(tab_idx)))
}

/*
Sets the maximum allowed width of the icon for the tab at index [param tab_idx]. This limit is applied on top of the default size of the icon and on top of [theme_item icon_max_width]. The height is adjusted according to the icon's ratio.
*/
func (self Instance) SetTabIconMaxWidth(tab_idx int, width int) { //gd:TabContainer.set_tab_icon_max_width
	class(self).SetTabIconMaxWidth(int64(tab_idx), int64(width))
}

/*
Returns the maximum allowed width of the icon for the tab at index [param tab_idx].
*/
func (self Instance) GetTabIconMaxWidth(tab_idx int) int { //gd:TabContainer.get_tab_icon_max_width
	return int(int(class(self).GetTabIconMaxWidth(int64(tab_idx))))
}

/*
If [param disabled] is [code]true[/code], disables the tab at index [param tab_idx], making it non-interactable.
*/
func (self Instance) SetTabDisabled(tab_idx int, disabled bool) { //gd:TabContainer.set_tab_disabled
	class(self).SetTabDisabled(int64(tab_idx), disabled)
}

/*
Returns [code]true[/code] if the tab at index [param tab_idx] is disabled.
*/
func (self Instance) IsTabDisabled(tab_idx int) bool { //gd:TabContainer.is_tab_disabled
	return bool(class(self).IsTabDisabled(int64(tab_idx)))
}

/*
If [param hidden] is [code]true[/code], hides the tab at index [param tab_idx], making it disappear from the tab area.
*/
func (self Instance) SetTabHidden(tab_idx int, hidden bool) { //gd:TabContainer.set_tab_hidden
	class(self).SetTabHidden(int64(tab_idx), hidden)
}

/*
Returns [code]true[/code] if the tab at index [param tab_idx] is hidden.
*/
func (self Instance) IsTabHidden(tab_idx int) bool { //gd:TabContainer.is_tab_hidden
	return bool(class(self).IsTabHidden(int64(tab_idx)))
}

/*
Sets the metadata value for the tab at index [param tab_idx], which can be retrieved later using [method get_tab_metadata].
*/
func (self Instance) SetTabMetadata(tab_idx int, metadata any) { //gd:TabContainer.set_tab_metadata
	class(self).SetTabMetadata(int64(tab_idx), variant.New(metadata))
}

/*
Returns the metadata value set to the tab at index [param tab_idx] using [method set_tab_metadata]. If no metadata was previously set, returns [code]null[/code] by default.
*/
func (self Instance) GetTabMetadata(tab_idx int) any { //gd:TabContainer.get_tab_metadata
	return any(class(self).GetTabMetadata(int64(tab_idx)).Interface())
}

/*
Sets the button icon from the tab at index [param tab_idx].
*/
func (self Instance) SetTabButtonIcon(tab_idx int, icon [1]gdclass.Texture2D) { //gd:TabContainer.set_tab_button_icon
	class(self).SetTabButtonIcon(int64(tab_idx), icon)
}

/*
Returns the button icon from the tab at index [param tab_idx].
*/
func (self Instance) GetTabButtonIcon(tab_idx int) [1]gdclass.Texture2D { //gd:TabContainer.get_tab_button_icon
	return [1]gdclass.Texture2D(class(self).GetTabButtonIcon(int64(tab_idx)))
}

/*
Returns the index of the tab at local coordinates [param point]. Returns [code]-1[/code] if the point is outside the control boundaries or if there's no tab at the queried position.
*/
func (self Instance) GetTabIdxAtPoint(point Vector2.XY) int { //gd:TabContainer.get_tab_idx_at_point
	return int(int(class(self).GetTabIdxAtPoint(Vector2.XY(point))))
}

/*
Returns the index of the tab tied to the given [param control]. The control must be a child of the [TabContainer].
*/
func (self Instance) GetTabIdxFromControl(control [1]gdclass.Control) int { //gd:TabContainer.get_tab_idx_from_control
	return int(int(class(self).GetTabIdxFromControl(control)))
}

/*
If set on a [Popup] node instance, a popup menu icon appears in the top-right corner of the [TabContainer] (setting it to [code]null[/code] will make it go away). Clicking it will expand the [Popup] node.
*/
func (self Instance) SetPopup(popup [1]gdclass.Node) { //gd:TabContainer.set_popup
	class(self).SetPopup(popup)
}

/*
Returns the [Popup] node instance if one has been set already with [method set_popup].
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member Window.visible] property.
*/
func (self Instance) GetPopup() [1]gdclass.Popup { //gd:TabContainer.get_popup
	return [1]gdclass.Popup(class(self).GetPopup())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.TabContainer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("TabContainer"))
	casted := Instance{*(*gdclass.TabContainer)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) TabAlignment() gdclass.TabBarAlignmentMode {
	return gdclass.TabBarAlignmentMode(class(self).GetTabAlignment())
}

func (self Instance) SetTabAlignment(value gdclass.TabBarAlignmentMode) {
	class(self).SetTabAlignment(value)
}

func (self Instance) CurrentTab() int {
	return int(int(class(self).GetCurrentTab()))
}

func (self Instance) SetCurrentTab(value int) {
	class(self).SetCurrentTab(int64(value))
}

func (self Instance) TabsPosition() gdclass.TabContainerTabPosition {
	return gdclass.TabContainerTabPosition(class(self).GetTabsPosition())
}

func (self Instance) SetTabsPosition(value gdclass.TabContainerTabPosition) {
	class(self).SetTabsPosition(value)
}

func (self Instance) ClipTabs() bool {
	return bool(class(self).GetClipTabs())
}

func (self Instance) SetClipTabs(value bool) {
	class(self).SetClipTabs(value)
}

func (self Instance) TabsVisible() bool {
	return bool(class(self).AreTabsVisible())
}

func (self Instance) SetTabsVisible(value bool) {
	class(self).SetTabsVisible(value)
}

func (self Instance) AllTabsInFront() bool {
	return bool(class(self).IsAllTabsInFront())
}

func (self Instance) SetAllTabsInFront(value bool) {
	class(self).SetAllTabsInFront(value)
}

func (self Instance) DragToRearrangeEnabled() bool {
	return bool(class(self).GetDragToRearrangeEnabled())
}

func (self Instance) SetDragToRearrangeEnabled(value bool) {
	class(self).SetDragToRearrangeEnabled(value)
}

func (self Instance) TabsRearrangeGroup() int {
	return int(int(class(self).GetTabsRearrangeGroup()))
}

func (self Instance) SetTabsRearrangeGroup(value int) {
	class(self).SetTabsRearrangeGroup(int64(value))
}

func (self Instance) UseHiddenTabsForMinSize() bool {
	return bool(class(self).GetUseHiddenTabsForMinSize())
}

func (self Instance) SetUseHiddenTabsForMinSize(value bool) {
	class(self).SetUseHiddenTabsForMinSize(value)
}

func (self Instance) TabFocusMode() gdclass.ControlFocusMode {
	return gdclass.ControlFocusMode(class(self).GetTabFocusMode())
}

func (self Instance) SetTabFocusMode(value gdclass.ControlFocusMode) {
	class(self).SetTabFocusMode(value)
}

func (self Instance) DeselectEnabled() bool {
	return bool(class(self).GetDeselectEnabled())
}

func (self Instance) SetDeselectEnabled(value bool) {
	class(self).SetDeselectEnabled(value)
}

/*
Returns the number of tabs.
*/
//go:nosplit
func (self class) GetTabCount() int64 { //gd:TabContainer.get_tab_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tab_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCurrentTab(tab_idx int64) { //gd:TabContainer.set_current_tab
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_current_tab, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCurrentTab() int64 { //gd:TabContainer.get_current_tab
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_current_tab, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the previously active tab index.
*/
//go:nosplit
func (self class) GetPreviousTab() int64 { //gd:TabContainer.get_previous_tab
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_previous_tab, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Selects the first available tab with lower index than the currently selected. Returns [code]true[/code] if tab selection changed.
*/
//go:nosplit
func (self class) SelectPreviousAvailable() bool { //gd:TabContainer.select_previous_available
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_select_previous_available, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Selects the first available tab with greater index than the currently selected. Returns [code]true[/code] if tab selection changed.
*/
//go:nosplit
func (self class) SelectNextAvailable() bool { //gd:TabContainer.select_next_available
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_select_next_available, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the child [Control] node located at the active tab index.
*/
//go:nosplit
func (self class) GetCurrentTabControl() [1]gdclass.Control { //gd:TabContainer.get_current_tab_control
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_current_tab_control, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Control{gd.PointerMustAssertInstanceID[gdclass.Control](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the [TabBar] contained in this container.
[b]Warning:[/b] This is a required internal node, removing and freeing it or editing its tabs may cause a crash. If you wish to edit the tabs, use the methods provided in [TabContainer].
*/
//go:nosplit
func (self class) GetTabBar() [1]gdclass.TabBar { //gd:TabContainer.get_tab_bar
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tab_bar, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TabBar{gd.PointerLifetimeBoundTo[gdclass.TabBar](self.AsObject(), r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the [Control] node from the tab at index [param tab_idx].
*/
//go:nosplit
func (self class) GetTabControl(tab_idx int64) [1]gdclass.Control { //gd:TabContainer.get_tab_control
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tab_control, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Control{gd.PointerMustAssertInstanceID[gdclass.Control](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTabAlignment(alignment gdclass.TabBarAlignmentMode) { //gd:TabContainer.set_tab_alignment
	var frame = callframe.New()
	callframe.Arg(frame, alignment)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_tab_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTabAlignment() gdclass.TabBarAlignmentMode { //gd:TabContainer.get_tab_alignment
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TabBarAlignmentMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tab_alignment, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTabsPosition(tabs_position gdclass.TabContainerTabPosition) { //gd:TabContainer.set_tabs_position
	var frame = callframe.New()
	callframe.Arg(frame, tabs_position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_tabs_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTabsPosition() gdclass.TabContainerTabPosition { //gd:TabContainer.get_tabs_position
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TabContainerTabPosition](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tabs_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetClipTabs(clip_tabs bool) { //gd:TabContainer.set_clip_tabs
	var frame = callframe.New()
	callframe.Arg(frame, clip_tabs)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_clip_tabs, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetClipTabs() bool { //gd:TabContainer.get_clip_tabs
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_clip_tabs, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTabsVisible(visible bool) { //gd:TabContainer.set_tabs_visible
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_tabs_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) AreTabsVisible() bool { //gd:TabContainer.are_tabs_visible
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_are_tabs_visible, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAllTabsInFront(is_front bool) { //gd:TabContainer.set_all_tabs_in_front
	var frame = callframe.New()
	callframe.Arg(frame, is_front)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_all_tabs_in_front, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsAllTabsInFront() bool { //gd:TabContainer.is_all_tabs_in_front
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_is_all_tabs_in_front, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets a custom title for the tab at index [param tab_idx] (tab titles default to the name of the indexed child node). Set it back to the child's name to make the tab default to it again.
*/
//go:nosplit
func (self class) SetTabTitle(tab_idx int64, title String.Readable) { //gd:TabContainer.set_tab_title
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, pointers.Get(gd.InternalString(title)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_tab_title, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the title of the tab at index [param tab_idx]. Tab titles default to the name of the indexed child node, but this can be overridden with [method set_tab_title].
*/
//go:nosplit
func (self class) GetTabTitle(tab_idx int64) String.Readable { //gd:TabContainer.get_tab_title
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tab_title, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets a custom tooltip text for tab at index [param tab_idx].
[b]Note:[/b] By default, if the [param tooltip] is empty and the tab text is truncated (not all characters fit into the tab), the title will be displayed as a tooltip. To hide the tooltip, assign [code]" "[/code] as the [param tooltip] text.
*/
//go:nosplit
func (self class) SetTabTooltip(tab_idx int64, tooltip String.Readable) { //gd:TabContainer.set_tab_tooltip
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, pointers.Get(gd.InternalString(tooltip)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_tab_tooltip, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the tooltip text of the tab at index [param tab_idx].
*/
//go:nosplit
func (self class) GetTabTooltip(tab_idx int64) String.Readable { //gd:TabContainer.get_tab_tooltip
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tab_tooltip, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets an icon for the tab at index [param tab_idx].
*/
//go:nosplit
func (self class) SetTabIcon(tab_idx int64, icon [1]gdclass.Texture2D) { //gd:TabContainer.set_tab_icon
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_tab_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [Texture2D] for the tab at index [param tab_idx] or [code]null[/code] if the tab has no [Texture2D].
*/
//go:nosplit
func (self class) GetTabIcon(tab_idx int64) [1]gdclass.Texture2D { //gd:TabContainer.get_tab_icon
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tab_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sets the maximum allowed width of the icon for the tab at index [param tab_idx]. This limit is applied on top of the default size of the icon and on top of [theme_item icon_max_width]. The height is adjusted according to the icon's ratio.
*/
//go:nosplit
func (self class) SetTabIconMaxWidth(tab_idx int64, width int64) { //gd:TabContainer.set_tab_icon_max_width
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, width)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_tab_icon_max_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the maximum allowed width of the icon for the tab at index [param tab_idx].
*/
//go:nosplit
func (self class) GetTabIconMaxWidth(tab_idx int64) int64 { //gd:TabContainer.get_tab_icon_max_width
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tab_icon_max_width, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param disabled] is [code]true[/code], disables the tab at index [param tab_idx], making it non-interactable.
*/
//go:nosplit
func (self class) SetTabDisabled(tab_idx int64, disabled bool) { //gd:TabContainer.set_tab_disabled
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, disabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_tab_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the tab at index [param tab_idx] is disabled.
*/
//go:nosplit
func (self class) IsTabDisabled(tab_idx int64) bool { //gd:TabContainer.is_tab_disabled
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_is_tab_disabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If [param hidden] is [code]true[/code], hides the tab at index [param tab_idx], making it disappear from the tab area.
*/
//go:nosplit
func (self class) SetTabHidden(tab_idx int64, hidden bool) { //gd:TabContainer.set_tab_hidden
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, hidden)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_tab_hidden, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the tab at index [param tab_idx] is hidden.
*/
//go:nosplit
func (self class) IsTabHidden(tab_idx int64) bool { //gd:TabContainer.is_tab_hidden
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_is_tab_hidden, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the metadata value for the tab at index [param tab_idx], which can be retrieved later using [method get_tab_metadata].
*/
//go:nosplit
func (self class) SetTabMetadata(tab_idx int64, metadata variant.Any) { //gd:TabContainer.set_tab_metadata
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(metadata)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_tab_metadata, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the metadata value set to the tab at index [param tab_idx] using [method set_tab_metadata]. If no metadata was previously set, returns [code]null[/code] by default.
*/
//go:nosplit
func (self class) GetTabMetadata(tab_idx int64) variant.Any { //gd:TabContainer.get_tab_metadata
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tab_metadata, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Implementation(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets the button icon from the tab at index [param tab_idx].
*/
//go:nosplit
func (self class) SetTabButtonIcon(tab_idx int64, icon [1]gdclass.Texture2D) { //gd:TabContainer.set_tab_button_icon
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_tab_button_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the button icon from the tab at index [param tab_idx].
*/
//go:nosplit
func (self class) GetTabButtonIcon(tab_idx int64) [1]gdclass.Texture2D { //gd:TabContainer.get_tab_button_icon
	var frame = callframe.New()
	callframe.Arg(frame, tab_idx)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tab_button_icon, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the index of the tab at local coordinates [param point]. Returns [code]-1[/code] if the point is outside the control boundaries or if there's no tab at the queried position.
*/
//go:nosplit
func (self class) GetTabIdxAtPoint(point Vector2.XY) int64 { //gd:TabContainer.get_tab_idx_at_point
	var frame = callframe.New()
	callframe.Arg(frame, point)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tab_idx_at_point, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the index of the tab tied to the given [param control]. The control must be a child of the [TabContainer].
*/
//go:nosplit
func (self class) GetTabIdxFromControl(control [1]gdclass.Control) int64 { //gd:TabContainer.get_tab_idx_from_control
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(control[0])[0])
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tab_idx_from_control, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
If set on a [Popup] node instance, a popup menu icon appears in the top-right corner of the [TabContainer] (setting it to [code]null[/code] will make it go away). Clicking it will expand the [Popup] node.
*/
//go:nosplit
func (self class) SetPopup(popup [1]gdclass.Node) { //gd:TabContainer.set_popup
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(popup[0].AsObject()[0]))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_popup, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [Popup] node instance if one has been set already with [method set_popup].
[b]Warning:[/b] This is a required internal node, removing and freeing it may cause a crash. If you wish to hide it or any of its children, use their [member Window.visible] property.
*/
//go:nosplit
func (self class) GetPopup() [1]gdclass.Popup { //gd:TabContainer.get_popup
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_popup, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Popup{gd.PointerMustAssertInstanceID[gdclass.Popup](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDragToRearrangeEnabled(enabled bool) { //gd:TabContainer.set_drag_to_rearrange_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_drag_to_rearrange_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDragToRearrangeEnabled() bool { //gd:TabContainer.get_drag_to_rearrange_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_drag_to_rearrange_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTabsRearrangeGroup(group_id int64) { //gd:TabContainer.set_tabs_rearrange_group
	var frame = callframe.New()
	callframe.Arg(frame, group_id)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_tabs_rearrange_group, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTabsRearrangeGroup() int64 { //gd:TabContainer.get_tabs_rearrange_group
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tabs_rearrange_group, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseHiddenTabsForMinSize(enabled bool) { //gd:TabContainer.set_use_hidden_tabs_for_min_size
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_use_hidden_tabs_for_min_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUseHiddenTabsForMinSize() bool { //gd:TabContainer.get_use_hidden_tabs_for_min_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_use_hidden_tabs_for_min_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTabFocusMode(focus_mode gdclass.ControlFocusMode) { //gd:TabContainer.set_tab_focus_mode
	var frame = callframe.New()
	callframe.Arg(frame, focus_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_tab_focus_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTabFocusMode() gdclass.ControlFocusMode { //gd:TabContainer.get_tab_focus_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.ControlFocusMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_tab_focus_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDeselectEnabled(enabled bool) { //gd:TabContainer.set_deselect_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_set_deselect_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDeselectEnabled() bool { //gd:TabContainer.get_deselect_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TabContainer.Bind_get_deselect_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnActiveTabRearranged(cb func(idx_to int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("active_tab_rearranged"), gd.NewCallable(cb), 0)
}

func (self Instance) OnTabChanged(cb func(tab int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("tab_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnTabClicked(cb func(tab int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("tab_clicked"), gd.NewCallable(cb), 0)
}

func (self Instance) OnTabHovered(cb func(tab int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("tab_hovered"), gd.NewCallable(cb), 0)
}

func (self Instance) OnTabSelected(cb func(tab int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("tab_selected"), gd.NewCallable(cb), 0)
}

func (self Instance) OnTabButtonPressed(cb func(tab int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("tab_button_pressed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnPrePopupPressed(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("pre_popup_pressed"), gd.NewCallable(cb), 0)
}

func (self class) AsTabContainer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsTabContainer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	default:
		return gd.VirtualByName(Container.Advanced(self.AsContainer()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Container.Instance(self.AsContainer()), name)
	}
}
func init() {
	gdclass.Register("TabContainer", func(ptr gd.Object) any {
		return [1]gdclass.TabContainer{*(*gdclass.TabContainer)(unsafe.Pointer(&ptr))}
	})
}

type TabPosition = gdclass.TabContainerTabPosition //gd:TabContainer.TabPosition

const (
	/*Places the tab bar at the top.*/
	PositionTop TabPosition = 0
	/*Places the tab bar at the bottom. The tab bar's [StyleBox] will be flipped vertically.*/
	PositionBottom TabPosition = 1
	/*Represents the size of the [enum TabPosition] enum.*/
	PositionMax TabPosition = 2
)
