package Control

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/CanvasItem"
import "grow.graphics/gd/object/Node"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Base class for all UI-related nodes. [Control] features a bounding rectangle that defines its extents, an anchor position relative to its parent control or the current viewport, and offsets relative to the anchor. The offsets update automatically when the node, any of its parents, or the screen size change.
For more information on Godot's UI system, anchors, offsets, and containers, see the related tutorials in the manual. To build flexible UIs, you'll need a mix of UI elements that inherit from [Control] and [Container] nodes.
[b]User Interface nodes and input[/b]
Godot propagates input events via viewports. Each [Viewport] is responsible for propagating [InputEvent]s to their child nodes. As the [member SceneTree.root] is a [Window], this already happens automatically for all UI elements in your game.
Input events are propagated through the [SceneTree] from the root node to all child nodes by calling [method Node._input]. For UI elements specifically, it makes more sense to override the virtual method [method _gui_input], which filters out unrelated input events, such as by checking z-order, [member mouse_filter], focus, or if the event was inside of the control's bounding box.
Call [method accept_event] so no other node receives the event. Once you accept an input, it becomes handled so [method Node._unhandled_input] will not process it.
Only one [Control] node can be in focus. Only the node in focus will receive events. To get the focus, call [method grab_focus]. [Control] nodes lose focus when another node grabs it, or if you hide the node in focus.
Sets [member mouse_filter] to [constant MOUSE_FILTER_IGNORE] to tell a [Control] node to ignore mouse or touch events. You'll need it if you place an icon on top of a button.
[Theme] resources change the Control's appearance. If you change the [Theme] on a [Control] node, it affects all of its children. To override some of the theme's parameters, call one of the [code]add_theme_*_override[/code] methods, like [method add_theme_font_override]. You can override the theme with the Inspector.
[b]Note:[/b] Theme items are [i]not[/i] [Object] properties. This means you can't access their values using [method Object.get] and [method Object.set]. Instead, use the [code]get_theme_*[/code] and [code]add_theme_*_override[/code] methods provided by this class.
	// Control methods that can be overridden by a [Class] that extends it.
	type Control interface {
		//Virtual method to be implemented by the user. Returns whether the given [param point] is inside this control.
		//If not overridden, default behavior is checking if the point is within control's Rect.
		//[b]Note:[/b] If you want to check if a point is inside the control, you can use [code]Rect2(Vector2.ZERO, size).has_point(point)[/code].
		HasPoint(point gd.Vector2) bool
		//User defined BiDi algorithm override function.
		//Returns an [Array] of [Vector3i] text ranges and text base directions, in the left-to-right order. Ranges should cover full source [param text] without overlaps. BiDi algorithm will be used on each range separately.
		StructuredTextParser(args gd.Array, text gd.String) gd.ArrayOf[gd.Vector3i]
		//Virtual method to be implemented by the user. Returns the minimum size for this control. Alternative to [member custom_minimum_size] for controlling minimum size via code. The actual minimum size will be the max value of these two (in each axis separately).
		//If not overridden, defaults to [constant Vector2.ZERO].
		//[b]Note:[/b] This method will not be called when the script is attached to a [Control] node that already overrides its minimum size (e.g. [Label], [Button], [PanelContainer] etc.). It can only be used with most basic GUI nodes, like [Control], [Container], [Panel] etc.
		GetMinimumSize() gd.Vector2
		//Virtual method to be implemented by the user. Returns the tooltip text for the position [param at_position] in control's local coordinates, which will typically appear when the cursor is resting over this control. See [method get_tooltip].
		//[b]Note:[/b] If this method returns an empty [String], no tooltip is displayed.
		GetTooltip(at_position gd.Vector2) gd.String
		//Godot calls this method to get data that can be dragged and dropped onto controls that expect drop data. Returns [code]null[/code] if there is no data to drag. Controls that want to receive drop data should implement [method _can_drop_data] and [method _drop_data]. [param at_position] is local to this control. Drag may be forced with [method force_drag].
		//A preview that will follow the mouse that should represent the data can be set with [method set_drag_preview]. A good time to set the preview is in this method.
		//[codeblocks]
		//[gdscript]
		//func _get_drag_data(position):
		//    var mydata = make_data() # This is your custom method generating the drag data.
		//    set_drag_preview(make_preview(mydata)) # This is your custom method generating the preview of the drag data.
		//    return mydata
		//[/gdscript]
		//[csharp]
		//public override Variant _GetDragData(Vector2 atPosition)
		//{
		//    var myData = MakeData(); // This is your custom method generating the drag data.
		//    SetDragPreview(MakePreview(myData)); // This is your custom method generating the preview of the drag data.
		//    return myData;
		//}
		//[/csharp]
		//[/codeblocks]
		GetDragData(at_position gd.Vector2) gd.Variant
		//Godot calls this method to test if [param data] from a control's [method _get_drag_data] can be dropped at [param at_position]. [param at_position] is local to this control.
		//This method should only be used to test the data. Process the data in [method _drop_data].
		//[codeblocks]
		//[gdscript]
		//func _can_drop_data(position, data):
		//    # Check position if it is relevant to you
		//    # Otherwise, just check data
		//    return typeof(data) == TYPE_DICTIONARY and data.has("expected")
		//[/gdscript]
		//[csharp]
		//public override bool _CanDropData(Vector2 atPosition, Variant data)
		//{
		//    // Check position if it is relevant to you
		//    // Otherwise, just check data
		//    return data.VariantType == Variant.Type.Dictionary && data.AsGodotDictionary().ContainsKey("expected");
		//}
		//[/csharp]
		//[/codeblocks]
		CanDropData(at_position gd.Vector2, data gd.Variant) bool
		//Godot calls this method to pass you the [param data] from a control's [method _get_drag_data] result. Godot first calls [method _can_drop_data] to test if [param data] is allowed to drop at [param at_position] where [param at_position] is local to this control.
		//[codeblocks]
		//[gdscript]
		//func _can_drop_data(position, data):
		//    return typeof(data) == TYPE_DICTIONARY and data.has("color")
		//
		//func _drop_data(position, data):
		//    var color = data["color"]
		//[/gdscript]
		//[csharp]
		//public override bool _CanDropData(Vector2 atPosition, Variant data)
		//{
		//    return data.VariantType == Variant.Type.Dictionary && dict.AsGodotDictionary().ContainsKey("color");
		//}
		//
		//public override void _DropData(Vector2 atPosition, Variant data)
		//{
		//    Color color = data.AsGodotDictionary()["color"].AsColor();
		//}
		//[/csharp]
		//[/codeblocks]
		DropData(at_position gd.Vector2, data gd.Variant) 
		//Virtual method to be implemented by the user. Returns a [Control] node that should be used as a tooltip instead of the default one. The [param for_text] includes the contents of the [member tooltip_text] property.
		//The returned node must be of type [Control] or Control-derived. It can have child nodes of any type. It is freed when the tooltip disappears, so make sure you always provide a new instance (if you want to use a pre-existing node from your scene tree, you can duplicate it and pass the duplicated instance). When [code]null[/code] or a non-Control node is returned, the default tooltip will be used instead.
		//The returned node will be added as child to a [PopupPanel], so you should only provide the contents of that panel. That [PopupPanel] can be themed using [method Theme.set_stylebox] for the type [code]"TooltipPanel"[/code] (see [member tooltip_text] for an example).
		//[b]Note:[/b] The tooltip is shrunk to minimal size. If you want to ensure it's fully visible, you might want to set its [member custom_minimum_size] to some non-zero value.
		//[b]Note:[/b] The node (and any relevant children) should be [member CanvasItem.visible] when returned, otherwise, the viewport that instantiates it will not be able to calculate its minimum size reliably.
		//[b]Example of usage with a custom-constructed node:[/b]
		//[codeblocks]
		//[gdscript]
		//func _make_custom_tooltip(for_text):
		//    var label = Label.new()
		//    label.text = for_text
		//    return label
		//[/gdscript]
		//[csharp]
		//public override Control _MakeCustomTooltip(string forText)
		//{
		//    var label = new Label();
		//    label.Text = forText;
		//    return label;
		//}
		//[/csharp]
		//[/codeblocks]
		//[b]Example of usage with a custom scene instance:[/b]
		//[codeblocks]
		//[gdscript]
		//func _make_custom_tooltip(for_text):
		//    var tooltip = preload("res://some_tooltip_scene.tscn").instantiate()
		//    tooltip.get_node("Label").text = for_text
		//    return tooltip
		//[/gdscript]
		//[csharp]
		//public override Control _MakeCustomTooltip(string forText)
		//{
		//    Node tooltip = ResourceLoader.Load<PackedScene>("res://some_tooltip_scene.tscn").Instantiate();
		//    tooltip.GetNode<Label>("Label").Text = forText;
		//    return tooltip;
		//}
		//[/csharp]
		//[/codeblocks]
		MakeCustomTooltip(for_text gd.String) gd.Object
		//Virtual method to be implemented by the user. Use this method to process and accept inputs on UI elements. See [method accept_event].
		//[b]Example usage for clicking a control:[/b]
		//[codeblocks]
		//[gdscript]
		//func _gui_input(event):
		//    if event is InputEventMouseButton:
		//        if event.button_index == MOUSE_BUTTON_LEFT and event.pressed:
		//            print("I've been clicked D:")
		//[/gdscript]
		//[csharp]
		//public override void _GuiInput(InputEvent @event)
		//{
		//    if (@event is InputEventMouseButton mb)
		//    {
		//        if (mb.ButtonIndex == MouseButton.Left && mb.Pressed)
		//        {
		//            GD.Print("I've been clicked D:");
		//        }
		//    }
		//}
		//[/csharp]
		//[/codeblocks]
		//The event won't trigger if:
		//* clicking outside the control (see [method _has_point]);
		//* control has [member mouse_filter] set to [constant MOUSE_FILTER_IGNORE];
		//* control is obstructed by another [Control] on top of it, which doesn't have [member mouse_filter] set to [constant MOUSE_FILTER_IGNORE];
		//* control's parent has [member mouse_filter] set to [constant MOUSE_FILTER_STOP] or has accepted the event;
		//* it happens outside the parent's rectangle and the parent has either [member clip_contents] enabled.
		//[b]Note:[/b] Event position is relative to the control origin.
		GuiInput(event object.InputEvent) 
	}

*/
type Simple [1]classdb.Control
func (self Simple) AcceptEvent() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AcceptEvent()
}
func (self Simple) GetMinimumSize() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetMinimumSize())
}
func (self Simple) GetCombinedMinimumSize() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetCombinedMinimumSize())
}
func (self Simple) SetAnchorsPreset(preset classdb.ControlLayoutPreset, keep_offsets bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAnchorsPreset(preset, keep_offsets)
}
func (self Simple) SetOffsetsPreset(preset classdb.ControlLayoutPreset, resize_mode classdb.ControlLayoutPresetMode, margin int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOffsetsPreset(preset, resize_mode, gd.Int(margin))
}
func (self Simple) SetAnchorsAndOffsetsPreset(preset classdb.ControlLayoutPreset, resize_mode classdb.ControlLayoutPresetMode, margin int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAnchorsAndOffsetsPreset(preset, resize_mode, gd.Int(margin))
}
func (self Simple) SetAnchor(side gd.Side, anchor float64, keep_offset bool, push_opposite_anchor bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAnchor(side, gd.Float(anchor), keep_offset, push_opposite_anchor)
}
func (self Simple) GetAnchor(side gd.Side) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetAnchor(side)))
}
func (self Simple) SetOffset(side gd.Side, offset float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetOffset(side, gd.Float(offset))
}
func (self Simple) GetOffset(offset gd.Side) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetOffset(offset)))
}
func (self Simple) SetAnchorAndOffset(side gd.Side, anchor float64, offset float64, push_opposite_anchor bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAnchorAndOffset(side, gd.Float(anchor), gd.Float(offset), push_opposite_anchor)
}
func (self Simple) SetBegin(position gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetBegin(position)
}
func (self Simple) SetEnd(position gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetEnd(position)
}
func (self Simple) SetPosition(position gd.Vector2, keep_offsets bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPosition(position, keep_offsets)
}
func (self Simple) SetSize(size gd.Vector2, keep_offsets bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetSize(size, keep_offsets)
}
func (self Simple) ResetSize() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ResetSize()
}
func (self Simple) SetCustomMinimumSize(size gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetCustomMinimumSize(size)
}
func (self Simple) SetGlobalPosition(position gd.Vector2, keep_offsets bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetGlobalPosition(position, keep_offsets)
}
func (self Simple) SetRotation(radians float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRotation(gd.Float(radians))
}
func (self Simple) SetRotationDegrees(degrees float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetRotationDegrees(gd.Float(degrees))
}
func (self Simple) SetScale(scale gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetScale(scale)
}
func (self Simple) SetPivotOffset(pivot_offset gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetPivotOffset(pivot_offset)
}
func (self Simple) GetBegin() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetBegin())
}
func (self Simple) GetEnd() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetEnd())
}
func (self Simple) GetPosition() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetPosition())
}
func (self Simple) GetSize() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetSize())
}
func (self Simple) GetRotation() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetRotation()))
}
func (self Simple) GetRotationDegrees() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetRotationDegrees()))
}
func (self Simple) GetScale() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetScale())
}
func (self Simple) GetPivotOffset() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetPivotOffset())
}
func (self Simple) GetCustomMinimumSize() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetCustomMinimumSize())
}
func (self Simple) GetParentAreaSize() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetParentAreaSize())
}
func (self Simple) GetGlobalPosition() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetGlobalPosition())
}
func (self Simple) GetScreenPosition() gd.Vector2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector2(Expert(self).GetScreenPosition())
}
func (self Simple) GetRect() gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2(Expert(self).GetRect())
}
func (self Simple) GetGlobalRect() gd.Rect2 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Rect2(Expert(self).GetGlobalRect())
}
func (self Simple) SetFocusMode(mode classdb.ControlFocusMode) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFocusMode(mode)
}
func (self Simple) GetFocusMode() classdb.ControlFocusMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ControlFocusMode(Expert(self).GetFocusMode())
}
func (self Simple) HasFocus() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasFocus())
}
func (self Simple) GrabFocus() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GrabFocus()
}
func (self Simple) ReleaseFocus() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ReleaseFocus()
}
func (self Simple) FindPrevValidFocus() [1]classdb.Control {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Control(Expert(self).FindPrevValidFocus(gc))
}
func (self Simple) FindNextValidFocus() [1]classdb.Control {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Control(Expert(self).FindNextValidFocus(gc))
}
func (self Simple) FindValidFocusNeighbor(side gd.Side) [1]classdb.Control {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Control(Expert(self).FindValidFocusNeighbor(gc, side))
}
func (self Simple) SetHSizeFlags(flags classdb.ControlSizeFlags) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHSizeFlags(flags)
}
func (self Simple) GetHSizeFlags() classdb.ControlSizeFlags {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ControlSizeFlags(Expert(self).GetHSizeFlags())
}
func (self Simple) SetStretchRatio(ratio float64) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetStretchRatio(gd.Float(ratio))
}
func (self Simple) GetStretchRatio() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetStretchRatio()))
}
func (self Simple) SetVSizeFlags(flags classdb.ControlSizeFlags) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVSizeFlags(flags)
}
func (self Simple) GetVSizeFlags() classdb.ControlSizeFlags {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ControlSizeFlags(Expert(self).GetVSizeFlags())
}
func (self Simple) SetTheme(theme [1]classdb.Theme) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTheme(theme)
}
func (self Simple) GetTheme() [1]classdb.Theme {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Theme(Expert(self).GetTheme(gc))
}
func (self Simple) SetThemeTypeVariation(theme_type string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetThemeTypeVariation(gc.StringName(theme_type))
}
func (self Simple) GetThemeTypeVariation() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetThemeTypeVariation(gc).String())
}
func (self Simple) BeginBulkThemeOverride() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).BeginBulkThemeOverride()
}
func (self Simple) EndBulkThemeOverride() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).EndBulkThemeOverride()
}
func (self Simple) AddThemeIconOverride(name string, texture [1]classdb.Texture2D) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddThemeIconOverride(gc.StringName(name), texture)
}
func (self Simple) AddThemeStyleboxOverride(name string, stylebox [1]classdb.StyleBox) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddThemeStyleboxOverride(gc.StringName(name), stylebox)
}
func (self Simple) AddThemeFontOverride(name string, font [1]classdb.Font) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddThemeFontOverride(gc.StringName(name), font)
}
func (self Simple) AddThemeFontSizeOverride(name string, font_size int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddThemeFontSizeOverride(gc.StringName(name), gd.Int(font_size))
}
func (self Simple) AddThemeColorOverride(name string, color gd.Color) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddThemeColorOverride(gc.StringName(name), color)
}
func (self Simple) AddThemeConstantOverride(name string, constant int) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).AddThemeConstantOverride(gc.StringName(name), gd.Int(constant))
}
func (self Simple) RemoveThemeIconOverride(name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveThemeIconOverride(gc.StringName(name))
}
func (self Simple) RemoveThemeStyleboxOverride(name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveThemeStyleboxOverride(gc.StringName(name))
}
func (self Simple) RemoveThemeFontOverride(name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveThemeFontOverride(gc.StringName(name))
}
func (self Simple) RemoveThemeFontSizeOverride(name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveThemeFontSizeOverride(gc.StringName(name))
}
func (self Simple) RemoveThemeColorOverride(name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveThemeColorOverride(gc.StringName(name))
}
func (self Simple) RemoveThemeConstantOverride(name string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).RemoveThemeConstantOverride(gc.StringName(name))
}
func (self Simple) GetThemeIcon(name string, theme_type string) [1]classdb.Texture2D {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Texture2D(Expert(self).GetThemeIcon(gc, gc.StringName(name), gc.StringName(theme_type)))
}
func (self Simple) GetThemeStylebox(name string, theme_type string) [1]classdb.StyleBox {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.StyleBox(Expert(self).GetThemeStylebox(gc, gc.StringName(name), gc.StringName(theme_type)))
}
func (self Simple) GetThemeFont(name string, theme_type string) [1]classdb.Font {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Font(Expert(self).GetThemeFont(gc, gc.StringName(name), gc.StringName(theme_type)))
}
func (self Simple) GetThemeFontSize(name string, theme_type string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetThemeFontSize(gc.StringName(name), gc.StringName(theme_type))))
}
func (self Simple) GetThemeColor(name string, theme_type string) gd.Color {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Color(Expert(self).GetThemeColor(gc.StringName(name), gc.StringName(theme_type)))
}
func (self Simple) GetThemeConstant(name string, theme_type string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetThemeConstant(gc.StringName(name), gc.StringName(theme_type))))
}
func (self Simple) HasThemeIconOverride(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasThemeIconOverride(gc.StringName(name)))
}
func (self Simple) HasThemeStyleboxOverride(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasThemeStyleboxOverride(gc.StringName(name)))
}
func (self Simple) HasThemeFontOverride(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasThemeFontOverride(gc.StringName(name)))
}
func (self Simple) HasThemeFontSizeOverride(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasThemeFontSizeOverride(gc.StringName(name)))
}
func (self Simple) HasThemeColorOverride(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasThemeColorOverride(gc.StringName(name)))
}
func (self Simple) HasThemeConstantOverride(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasThemeConstantOverride(gc.StringName(name)))
}
func (self Simple) HasThemeIcon(name string, theme_type string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasThemeIcon(gc.StringName(name), gc.StringName(theme_type)))
}
func (self Simple) HasThemeStylebox(name string, theme_type string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasThemeStylebox(gc.StringName(name), gc.StringName(theme_type)))
}
func (self Simple) HasThemeFont(name string, theme_type string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasThemeFont(gc.StringName(name), gc.StringName(theme_type)))
}
func (self Simple) HasThemeFontSize(name string, theme_type string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasThemeFontSize(gc.StringName(name), gc.StringName(theme_type)))
}
func (self Simple) HasThemeColor(name string, theme_type string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasThemeColor(gc.StringName(name), gc.StringName(theme_type)))
}
func (self Simple) HasThemeConstant(name string, theme_type string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).HasThemeConstant(gc.StringName(name), gc.StringName(theme_type)))
}
func (self Simple) GetThemeDefaultBaseScale() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(Expert(self).GetThemeDefaultBaseScale()))
}
func (self Simple) GetThemeDefaultFont() [1]classdb.Font {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Font(Expert(self).GetThemeDefaultFont(gc))
}
func (self Simple) GetThemeDefaultFontSize() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(Expert(self).GetThemeDefaultFontSize()))
}
func (self Simple) GetParentControl() [1]classdb.Control {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Control(Expert(self).GetParentControl(gc))
}
func (self Simple) SetHGrowDirection(direction classdb.ControlGrowDirection) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetHGrowDirection(direction)
}
func (self Simple) GetHGrowDirection() classdb.ControlGrowDirection {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ControlGrowDirection(Expert(self).GetHGrowDirection())
}
func (self Simple) SetVGrowDirection(direction classdb.ControlGrowDirection) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetVGrowDirection(direction)
}
func (self Simple) GetVGrowDirection() classdb.ControlGrowDirection {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ControlGrowDirection(Expert(self).GetVGrowDirection())
}
func (self Simple) SetTooltipText(hint string) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetTooltipText(gc.String(hint))
}
func (self Simple) GetTooltipText() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetTooltipText(gc).String())
}
func (self Simple) GetTooltip(at_position gd.Vector2) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(Expert(self).GetTooltip(gc, at_position).String())
}
func (self Simple) SetDefaultCursorShape(shape classdb.ControlCursorShape) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDefaultCursorShape(shape)
}
func (self Simple) GetDefaultCursorShape() classdb.ControlCursorShape {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ControlCursorShape(Expert(self).GetDefaultCursorShape())
}
func (self Simple) GetCursorShape(position gd.Vector2) classdb.ControlCursorShape {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ControlCursorShape(Expert(self).GetCursorShape(position))
}
func (self Simple) SetFocusNeighbor(side gd.Side, neighbor gd.NodePath) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFocusNeighbor(side, neighbor)
}
func (self Simple) GetFocusNeighbor(side gd.Side) gd.NodePath {
	gc := gd.GarbageCollector(); _ = gc
	return gd.NodePath(Expert(self).GetFocusNeighbor(gc, side))
}
func (self Simple) SetFocusNext(next gd.NodePath) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFocusNext(next)
}
func (self Simple) GetFocusNext() gd.NodePath {
	gc := gd.GarbageCollector(); _ = gc
	return gd.NodePath(Expert(self).GetFocusNext(gc))
}
func (self Simple) SetFocusPrevious(previous gd.NodePath) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetFocusPrevious(previous)
}
func (self Simple) GetFocusPrevious() gd.NodePath {
	gc := gd.GarbageCollector(); _ = gc
	return gd.NodePath(Expert(self).GetFocusPrevious(gc))
}
func (self Simple) ForceDrag(data gd.Variant, preview [1]classdb.Control) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).ForceDrag(data, preview)
}
func (self Simple) SetMouseFilter(filter classdb.ControlMouseFilter) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetMouseFilter(filter)
}
func (self Simple) GetMouseFilter() classdb.ControlMouseFilter {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ControlMouseFilter(Expert(self).GetMouseFilter())
}
func (self Simple) SetForcePassScrollEvents(force_pass_scroll_events bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetForcePassScrollEvents(force_pass_scroll_events)
}
func (self Simple) IsForcePassScrollEvents() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsForcePassScrollEvents())
}
func (self Simple) SetClipContents(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetClipContents(enable)
}
func (self Simple) IsClippingContents() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsClippingContents())
}
func (self Simple) GrabClickFocus() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).GrabClickFocus()
}
func (self Simple) SetDragForwarding(drag_func gd.Callable, can_drop_func gd.Callable, drop_func gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDragForwarding(drag_func, can_drop_func, drop_func)
}
func (self Simple) SetDragPreview(control [1]classdb.Control) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetDragPreview(control)
}
func (self Simple) IsDragSuccessful() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsDragSuccessful())
}
func (self Simple) WarpMouse(position gd.Vector2) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).WarpMouse(position)
}
func (self Simple) SetShortcutContext(node [1]classdb.Node) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetShortcutContext(node)
}
func (self Simple) GetShortcutContext() [1]classdb.Node {
	gc := gd.GarbageCollector(); _ = gc
	return [1]classdb.Node(Expert(self).GetShortcutContext(gc))
}
func (self Simple) UpdateMinimumSize() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).UpdateMinimumSize()
}
func (self Simple) SetLayoutDirection(direction classdb.ControlLayoutDirection) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLayoutDirection(direction)
}
func (self Simple) GetLayoutDirection() classdb.ControlLayoutDirection {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ControlLayoutDirection(Expert(self).GetLayoutDirection())
}
func (self Simple) IsLayoutRtl() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsLayoutRtl())
}
func (self Simple) SetAutoTranslate(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetAutoTranslate(enable)
}
func (self Simple) IsAutoTranslating() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsAutoTranslating())
}
func (self Simple) SetLocalizeNumeralSystem(enable bool) {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).SetLocalizeNumeralSystem(enable)
}
func (self Simple) IsLocalizingNumeralSystem() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(Expert(self).IsLocalizingNumeralSystem())
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.Control
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Virtual method to be implemented by the user. Returns whether the given [param point] is inside this control.
If not overridden, default behavior is checking if the point is within control's Rect.
[b]Note:[/b] If you want to check if a point is inside the control, you can use [code]Rect2(Vector2.ZERO, size).has_point(point)[/code].
*/
func (class) _has_point(impl func(ptr unsafe.Pointer, point gd.Vector2) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var point = gd.UnsafeGet[gd.Vector2](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, point)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
User defined BiDi algorithm override function.
Returns an [Array] of [Vector3i] text ranges and text base directions, in the left-to-right order. Ranges should cover full source [param text] without overlaps. BiDi algorithm will be used on each range separately.
*/
func (class) _structured_text_parser(impl func(ptr unsafe.Pointer, args gd.Array, text gd.String) gd.ArrayOf[gd.Vector3i], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var args = mmm.Let[gd.Array](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var text = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, args, text)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		ctx.End()
	}
}

/*
Virtual method to be implemented by the user. Returns the minimum size for this control. Alternative to [member custom_minimum_size] for controlling minimum size via code. The actual minimum size will be the max value of these two (in each axis separately).
If not overridden, defaults to [constant Vector2.ZERO].
[b]Note:[/b] This method will not be called when the script is attached to a [Control] node that already overrides its minimum size (e.g. [Label], [Button], [PanelContainer] etc.). It can only be used with most basic GUI nodes, like [Control], [Container], [Panel] etc.
*/
func (class) _get_minimum_size(impl func(ptr unsafe.Pointer) gd.Vector2, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Virtual method to be implemented by the user. Returns the tooltip text for the position [param at_position] in control's local coordinates, which will typically appear when the cursor is resting over this control. See [method get_tooltip].
[b]Note:[/b] If this method returns an empty [String], no tooltip is displayed.
*/
func (class) _get_tooltip(impl func(ptr unsafe.Pointer, at_position gd.Vector2) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var at_position = gd.UnsafeGet[gd.Vector2](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, at_position)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Godot calls this method to get data that can be dragged and dropped onto controls that expect drop data. Returns [code]null[/code] if there is no data to drag. Controls that want to receive drop data should implement [method _can_drop_data] and [method _drop_data]. [param at_position] is local to this control. Drag may be forced with [method force_drag].
A preview that will follow the mouse that should represent the data can be set with [method set_drag_preview]. A good time to set the preview is in this method.
[codeblocks]
[gdscript]
func _get_drag_data(position):
    var mydata = make_data() # This is your custom method generating the drag data.
    set_drag_preview(make_preview(mydata)) # This is your custom method generating the preview of the drag data.
    return mydata
[/gdscript]
[csharp]
public override Variant _GetDragData(Vector2 atPosition)
{
    var myData = MakeData(); // This is your custom method generating the drag data.
    SetDragPreview(MakePreview(myData)); // This is your custom method generating the preview of the drag data.
    return myData;
}
[/csharp]
[/codeblocks]
*/
func (class) _get_drag_data(impl func(ptr unsafe.Pointer, at_position gd.Vector2) gd.Variant, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var at_position = gd.UnsafeGet[gd.Vector2](p_args,0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, at_position)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Godot calls this method to test if [param data] from a control's [method _get_drag_data] can be dropped at [param at_position]. [param at_position] is local to this control.
This method should only be used to test the data. Process the data in [method _drop_data].
[codeblocks]
[gdscript]
func _can_drop_data(position, data):
    # Check position if it is relevant to you
    # Otherwise, just check data
    return typeof(data) == TYPE_DICTIONARY and data.has("expected")
[/gdscript]
[csharp]
public override bool _CanDropData(Vector2 atPosition, Variant data)
{
    // Check position if it is relevant to you
    // Otherwise, just check data
    return data.VariantType == Variant.Type.Dictionary && data.AsGodotDictionary().ContainsKey("expected");
}
[/csharp]
[/codeblocks]
*/
func (class) _can_drop_data(impl func(ptr unsafe.Pointer, at_position gd.Vector2, data gd.Variant) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var at_position = gd.UnsafeGet[gd.Vector2](p_args,0)
		var data = mmm.Let[gd.Variant](ctx.Lifetime, ctx.API, gd.UnsafeGet[[3]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, at_position, data)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Godot calls this method to pass you the [param data] from a control's [method _get_drag_data] result. Godot first calls [method _can_drop_data] to test if [param data] is allowed to drop at [param at_position] where [param at_position] is local to this control.
[codeblocks]
[gdscript]
func _can_drop_data(position, data):
    return typeof(data) == TYPE_DICTIONARY and data.has("color")

func _drop_data(position, data):
    var color = data["color"]
[/gdscript]
[csharp]
public override bool _CanDropData(Vector2 atPosition, Variant data)
{
    return data.VariantType == Variant.Type.Dictionary && dict.AsGodotDictionary().ContainsKey("color");
}

public override void _DropData(Vector2 atPosition, Variant data)
{
    Color color = data.AsGodotDictionary()["color"].AsColor();
}
[/csharp]
[/codeblocks]
*/
func (class) _drop_data(impl func(ptr unsafe.Pointer, at_position gd.Vector2, data gd.Variant) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var at_position = gd.UnsafeGet[gd.Vector2](p_args,0)
		var data = mmm.Let[gd.Variant](ctx.Lifetime, ctx.API, gd.UnsafeGet[[3]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, at_position, data)
		ctx.End()
	}
}

/*
Virtual method to be implemented by the user. Returns a [Control] node that should be used as a tooltip instead of the default one. The [param for_text] includes the contents of the [member tooltip_text] property.
The returned node must be of type [Control] or Control-derived. It can have child nodes of any type. It is freed when the tooltip disappears, so make sure you always provide a new instance (if you want to use a pre-existing node from your scene tree, you can duplicate it and pass the duplicated instance). When [code]null[/code] or a non-Control node is returned, the default tooltip will be used instead.
The returned node will be added as child to a [PopupPanel], so you should only provide the contents of that panel. That [PopupPanel] can be themed using [method Theme.set_stylebox] for the type [code]"TooltipPanel"[/code] (see [member tooltip_text] for an example).
[b]Note:[/b] The tooltip is shrunk to minimal size. If you want to ensure it's fully visible, you might want to set its [member custom_minimum_size] to some non-zero value.
[b]Note:[/b] The node (and any relevant children) should be [member CanvasItem.visible] when returned, otherwise, the viewport that instantiates it will not be able to calculate its minimum size reliably.
[b]Example of usage with a custom-constructed node:[/b]
[codeblocks]
[gdscript]
func _make_custom_tooltip(for_text):
    var label = Label.new()
    label.text = for_text
    return label
[/gdscript]
[csharp]
public override Control _MakeCustomTooltip(string forText)
{
    var label = new Label();
    label.Text = forText;
    return label;
}
[/csharp]
[/codeblocks]
[b]Example of usage with a custom scene instance:[/b]
[codeblocks]
[gdscript]
func _make_custom_tooltip(for_text):
    var tooltip = preload("res://some_tooltip_scene.tscn").instantiate()
    tooltip.get_node("Label").text = for_text
    return tooltip
[/gdscript]
[csharp]
public override Control _MakeCustomTooltip(string forText)
{
    Node tooltip = ResourceLoader.Load<PackedScene>("res://some_tooltip_scene.tscn").Instantiate();
    tooltip.GetNode<Label>("Label").Text = forText;
    return tooltip;
}
[/csharp]
[/codeblocks]
*/
func (class) _make_custom_tooltip(impl func(ptr unsafe.Pointer, for_text gd.String) gd.Object, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var for_text = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, for_text)
		gd.UnsafeSet(p_back, mmm.End(ret.AsPointer()))
		ctx.End()
	}
}

/*
Virtual method to be implemented by the user. Use this method to process and accept inputs on UI elements. See [method accept_event].
[b]Example usage for clicking a control:[/b]
[codeblocks]
[gdscript]
func _gui_input(event):
    if event is InputEventMouseButton:
        if event.button_index == MOUSE_BUTTON_LEFT and event.pressed:
            print("I've been clicked D:")
[/gdscript]
[csharp]
public override void _GuiInput(InputEvent @event)
{
    if (@event is InputEventMouseButton mb)
    {
        if (mb.ButtonIndex == MouseButton.Left && mb.Pressed)
        {
            GD.Print("I've been clicked D:");
        }
    }
}
[/csharp]
[/codeblocks]
The event won't trigger if:
* clicking outside the control (see [method _has_point]);
* control has [member mouse_filter] set to [constant MOUSE_FILTER_IGNORE];
* control is obstructed by another [Control] on top of it, which doesn't have [member mouse_filter] set to [constant MOUSE_FILTER_IGNORE];
* control's parent has [member mouse_filter] set to [constant MOUSE_FILTER_STOP] or has accepted the event;
* it happens outside the parent's rectangle and the parent has either [member clip_contents] enabled.
[b]Note:[/b] Event position is relative to the control origin.
*/
func (class) _gui_input(impl func(ptr unsafe.Pointer, event object.InputEvent) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var event object.InputEvent
		event[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, event)
		ctx.End()
	}
}

/*
Marks an input event as handled. Once you accept an input event, it stops propagating, even to nodes listening to [method Node._unhandled_input] or [method Node._unhandled_key_input].
[b]Note:[/b] This does not affect the methods in [Input], only the way events are propagated.
*/
//go:nosplit
func (self class) AcceptEvent()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_accept_event, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the minimum size for this control. See [member custom_minimum_size].
*/
//go:nosplit
func (self class) GetMinimumSize() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_minimum_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns combined minimum size from [member custom_minimum_size] and [method get_minimum_size].
*/
//go:nosplit
func (self class) GetCombinedMinimumSize() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_combined_minimum_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the anchors to a [param preset] from [enum Control.LayoutPreset] enum. This is the code equivalent to using the Layout menu in the 2D editor.
If [param keep_offsets] is [code]true[/code], control's position will also be updated.
*/
//go:nosplit
func (self class) SetAnchorsPreset(preset classdb.ControlLayoutPreset, keep_offsets bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, preset)
	callframe.Arg(frame, keep_offsets)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_anchors_preset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the offsets to a [param preset] from [enum Control.LayoutPreset] enum. This is the code equivalent to using the Layout menu in the 2D editor.
Use parameter [param resize_mode] with constants from [enum Control.LayoutPresetMode] to better determine the resulting size of the [Control]. Constant size will be ignored if used with presets that change size, e.g. [constant PRESET_LEFT_WIDE].
Use parameter [param margin] to determine the gap between the [Control] and the edges.
*/
//go:nosplit
func (self class) SetOffsetsPreset(preset classdb.ControlLayoutPreset, resize_mode classdb.ControlLayoutPresetMode, margin gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, preset)
	callframe.Arg(frame, resize_mode)
	callframe.Arg(frame, margin)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_offsets_preset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets both anchor preset and offset preset. See [method set_anchors_preset] and [method set_offsets_preset].
*/
//go:nosplit
func (self class) SetAnchorsAndOffsetsPreset(preset classdb.ControlLayoutPreset, resize_mode classdb.ControlLayoutPresetMode, margin gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, preset)
	callframe.Arg(frame, resize_mode)
	callframe.Arg(frame, margin)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_anchors_and_offsets_preset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the anchor for the specified [enum Side] to [param anchor]. A setter method for [member anchor_bottom], [member anchor_left], [member anchor_right] and [member anchor_top].
If [param keep_offset] is [code]true[/code], offsets aren't updated after this operation.
If [param push_opposite_anchor] is [code]true[/code] and the opposite anchor overlaps this anchor, the opposite one will have its value overridden. For example, when setting left anchor to 1 and the right anchor has value of 0.5, the right anchor will also get value of 1. If [param push_opposite_anchor] was [code]false[/code], the left anchor would get value 0.5.
*/
//go:nosplit
func (self class) SetAnchor(side gd.Side, anchor gd.Float, keep_offset bool, push_opposite_anchor bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, side)
	callframe.Arg(frame, anchor)
	callframe.Arg(frame, keep_offset)
	callframe.Arg(frame, push_opposite_anchor)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_anchor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the anchor for the specified [enum Side]. A getter method for [member anchor_bottom], [member anchor_left], [member anchor_right] and [member anchor_top].
*/
//go:nosplit
func (self class) GetAnchor(side gd.Side) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, side)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_anchor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the offset for the specified [enum Side] to [param offset]. A setter method for [member offset_bottom], [member offset_left], [member offset_right] and [member offset_top].
*/
//go:nosplit
func (self class) SetOffset(side gd.Side, offset gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, side)
	callframe.Arg(frame, offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the offset for the specified [enum Side]. A getter method for [member offset_bottom], [member offset_left], [member offset_right] and [member offset_top].
*/
//go:nosplit
func (self class) GetOffset(offset gd.Side) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, offset)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Works the same as [method set_anchor], but instead of [code]keep_offset[/code] argument and automatic update of offset, it allows to set the offset yourself (see [method set_offset]).
*/
//go:nosplit
func (self class) SetAnchorAndOffset(side gd.Side, anchor gd.Float, offset gd.Float, push_opposite_anchor bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, side)
	callframe.Arg(frame, anchor)
	callframe.Arg(frame, offset)
	callframe.Arg(frame, push_opposite_anchor)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_anchor_and_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets [member offset_left] and [member offset_top] at the same time. Equivalent of changing [member position].
*/
//go:nosplit
func (self class) SetBegin(position gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets [member offset_right] and [member offset_bottom] at the same time.
*/
//go:nosplit
func (self class) SetEnd(position gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [member position] to given [param position].
If [param keep_offsets] is [code]true[/code], control's anchors will be updated instead of offsets.
*/
//go:nosplit
func (self class) SetPosition(position gd.Vector2, keep_offsets bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, keep_offsets)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the size (see [member size]).
If [param keep_offsets] is [code]true[/code], control's anchors will be updated instead of offsets.
*/
//go:nosplit
func (self class) SetSize(size gd.Vector2, keep_offsets bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	callframe.Arg(frame, keep_offsets)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Resets the size to [method get_combined_minimum_size]. This is equivalent to calling [code]set_size(Vector2())[/code] (or any size below the minimum).
*/
//go:nosplit
func (self class) ResetSize()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_reset_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetCustomMinimumSize(size gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_custom_minimum_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [member global_position] to given [param position].
If [param keep_offsets] is [code]true[/code], control's anchors will be updated instead of offsets.
*/
//go:nosplit
func (self class) SetGlobalPosition(position gd.Vector2, keep_offsets bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	callframe.Arg(frame, keep_offsets)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_global_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetRotation(radians gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, radians)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetRotationDegrees(degrees gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, degrees)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_rotation_degrees, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetScale(scale gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetPivotOffset(pivot_offset gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, pivot_offset)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_pivot_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [member offset_left] and [member offset_top]. See also [member position].
*/
//go:nosplit
func (self class) GetBegin() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_begin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [member offset_right] and [member offset_bottom].
*/
//go:nosplit
func (self class) GetEnd() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_end, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetPosition() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetSize() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetRotation() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_rotation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetRotationDegrees() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_rotation_degrees, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetScale() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetPivotOffset() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_pivot_offset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetCustomMinimumSize() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_custom_minimum_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the width/height occupied in the parent control.
*/
//go:nosplit
func (self class) GetParentAreaSize() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_parent_area_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) GetGlobalPosition() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_global_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the position of this [Control] in global screen coordinates (i.e. taking window position into account). Mostly useful for editor plugins.
Equals to [member global_position] if the window is embedded (see [member Viewport.gui_embed_subwindows]).
[b]Example usage for showing a popup:[/b]
[codeblock]
popup_menu.position = get_screen_position() + get_local_mouse_position()
popup_menu.reset_size()
popup_menu.popup()
[/codeblock]
*/
//go:nosplit
func (self class) GetScreenPosition() gd.Vector2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_screen_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the position and size of the control in the coordinate system of the containing node. See [member position], [member scale] and [member size].
[b]Note:[/b] If [member rotation] is not the default rotation, the resulting size is not meaningful.
[b]Note:[/b] Setting [member Viewport.gui_snap_controls_to_pixels] to [code]true[/code] can lead to rounding inaccuracies between the displayed control and the returned [Rect2].
*/
//go:nosplit
func (self class) GetRect() gd.Rect2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the position and size of the control relative to the containing canvas. See [member global_position] and [member size].
[b]Note:[/b] If the node itself or any parent [CanvasItem] between the node and the canvas have a non default rotation or skew, the resulting size is likely not meaningful.
[b]Note:[/b] Setting [member Viewport.gui_snap_controls_to_pixels] to [code]true[/code] can lead to rounding inaccuracies between the displayed control and the returned [Rect2].
*/
//go:nosplit
func (self class) GetGlobalRect() gd.Rect2 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_global_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFocusMode(mode classdb.ControlFocusMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_focus_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFocusMode() classdb.ControlFocusMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ControlFocusMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_focus_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if this is the current focused control. See [member focus_mode].
*/
//go:nosplit
func (self class) HasFocus() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_has_focus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Steal the focus from another control and become the focused control (see [member focus_mode]).
[b]Note:[/b] Using this method together with [method Callable.call_deferred] makes it more reliable, especially when called inside [method Node._ready].
*/
//go:nosplit
func (self class) GrabFocus()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_grab_focus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Give up the focus. No other control will be able to receive input.
*/
//go:nosplit
func (self class) ReleaseFocus()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_release_focus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Finds the previous (above in the tree) [Control] that can receive the focus.
*/
//go:nosplit
func (self class) FindPrevValidFocus(ctx gd.Lifetime) object.Control {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_find_prev_valid_focus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Control
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Finds the next (below in the tree) [Control] that can receive the focus.
*/
//go:nosplit
func (self class) FindNextValidFocus(ctx gd.Lifetime) object.Control {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_find_next_valid_focus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Control
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Finds the next [Control] that can receive the focus on the specified [enum Side].
[b]Note:[/b] This is different from [method get_focus_neighbor], which returns the path of a specified focus neighbor.
*/
//go:nosplit
func (self class) FindValidFocusNeighbor(ctx gd.Lifetime, side gd.Side) object.Control {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, side)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_find_valid_focus_neighbor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Control
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHSizeFlags(flags classdb.ControlSizeFlags)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_h_size_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHSizeFlags() classdb.ControlSizeFlags {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ControlSizeFlags](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_h_size_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetStretchRatio(ratio gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, ratio)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_stretch_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetStretchRatio() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_stretch_ratio, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVSizeFlags(flags classdb.ControlSizeFlags)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_v_size_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVSizeFlags() classdb.ControlSizeFlags {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ControlSizeFlags](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_v_size_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTheme(theme object.Theme)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(theme[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_theme, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTheme(ctx gd.Lifetime) object.Theme {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_theme, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Theme
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetThemeTypeVariation(theme_type gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_theme_type_variation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetThemeTypeVariation(ctx gd.Lifetime) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_theme_type_variation, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Prevents [code]*_theme_*_override[/code] methods from emitting [constant NOTIFICATION_THEME_CHANGED] until [method end_bulk_theme_override] is called.
*/
//go:nosplit
func (self class) BeginBulkThemeOverride()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_begin_bulk_theme_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Ends a bulk theme override update. See [method begin_bulk_theme_override].
*/
//go:nosplit
func (self class) EndBulkThemeOverride()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_end_bulk_theme_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a local override for a theme icon with the specified [param name]. Local overrides always take precedence when fetching theme items for the control. An override can be removed with [method remove_theme_icon_override].
See also [method get_theme_icon].
*/
//go:nosplit
func (self class) AddThemeIconOverride(name gd.StringName, texture object.Texture2D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(texture[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_add_theme_icon_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a local override for a theme [StyleBox] with the specified [param name]. Local overrides always take precedence when fetching theme items for the control. An override can be removed with [method remove_theme_stylebox_override].
See also [method get_theme_stylebox].
[b]Example of modifying a property in a StyleBox by duplicating it:[/b]
[codeblocks]
[gdscript]
# The snippet below assumes the child node MyButton has a StyleBoxFlat assigned.
# Resources are shared across instances, so we need to duplicate it
# to avoid modifying the appearance of all other buttons.
var new_stylebox_normal = $MyButton.get_theme_stylebox("normal").duplicate()
new_stylebox_normal.border_width_top = 3
new_stylebox_normal.border_color = Color(0, 1, 0.5)
$MyButton.add_theme_stylebox_override("normal", new_stylebox_normal)
# Remove the stylebox override.
$MyButton.remove_theme_stylebox_override("normal")
[/gdscript]
[csharp]
// The snippet below assumes the child node MyButton has a StyleBoxFlat assigned.
// Resources are shared across instances, so we need to duplicate it
// to avoid modifying the appearance of all other buttons.
StyleBoxFlat newStyleboxNormal = GetNode<Button>("MyButton").GetThemeStylebox("normal").Duplicate() as StyleBoxFlat;
newStyleboxNormal.BorderWidthTop = 3;
newStyleboxNormal.BorderColor = new Color(0, 1, 0.5f);
GetNode<Button>("MyButton").AddThemeStyleboxOverride("normal", newStyleboxNormal);
// Remove the stylebox override.
GetNode<Button>("MyButton").RemoveThemeStyleboxOverride("normal");
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) AddThemeStyleboxOverride(name gd.StringName, stylebox object.StyleBox)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(stylebox[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_add_theme_stylebox_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a local override for a theme [Font] with the specified [param name]. Local overrides always take precedence when fetching theme items for the control. An override can be removed with [method remove_theme_font_override].
See also [method get_theme_font].
*/
//go:nosplit
func (self class) AddThemeFontOverride(name gd.StringName, font object.Font)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(font[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_add_theme_font_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a local override for a theme font size with the specified [param name]. Local overrides always take precedence when fetching theme items for the control. An override can be removed with [method remove_theme_font_size_override].
See also [method get_theme_font_size].
*/
//go:nosplit
func (self class) AddThemeFontSizeOverride(name gd.StringName, font_size gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, font_size)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_add_theme_font_size_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a local override for a theme [Color] with the specified [param name]. Local overrides always take precedence when fetching theme items for the control. An override can be removed with [method remove_theme_color_override].
See also [method get_theme_color].
[b]Example of overriding a label's color and resetting it later:[/b]
[codeblocks]
[gdscript]
# Given the child Label node "MyLabel", override its font color with a custom value.
$MyLabel.add_theme_color_override("font_color", Color(1, 0.5, 0))
# Reset the font color of the child label.
$MyLabel.remove_theme_color_override("font_color")
# Alternatively it can be overridden with the default value from the Label type.
$MyLabel.add_theme_color_override("font_color", get_theme_color("font_color", "Label"))
[/gdscript]
[csharp]
// Given the child Label node "MyLabel", override its font color with a custom value.
GetNode<Label>("MyLabel").AddThemeColorOverride("font_color", new Color(1, 0.5f, 0));
// Reset the font color of the child label.
GetNode<Label>("MyLabel").RemoveThemeColorOverride("font_color");
// Alternatively it can be overridden with the default value from the Label type.
GetNode<Label>("MyLabel").AddThemeColorOverride("font_color", GetThemeColor("font_color", "Label"));
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) AddThemeColorOverride(name gd.StringName, color gd.Color)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, color)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_add_theme_color_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Creates a local override for a theme constant with the specified [param name]. Local overrides always take precedence when fetching theme items for the control. An override can be removed with [method remove_theme_constant_override].
See also [method get_theme_constant].
*/
//go:nosplit
func (self class) AddThemeConstantOverride(name gd.StringName, constant gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, constant)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_add_theme_constant_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a local override for a theme icon with the specified [param name] previously added by [method add_theme_icon_override] or via the Inspector dock.
*/
//go:nosplit
func (self class) RemoveThemeIconOverride(name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_remove_theme_icon_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a local override for a theme [StyleBox] with the specified [param name] previously added by [method add_theme_stylebox_override] or via the Inspector dock.
*/
//go:nosplit
func (self class) RemoveThemeStyleboxOverride(name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_remove_theme_stylebox_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a local override for a theme [Font] with the specified [param name] previously added by [method add_theme_font_override] or via the Inspector dock.
*/
//go:nosplit
func (self class) RemoveThemeFontOverride(name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_remove_theme_font_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a local override for a theme font size with the specified [param name] previously added by [method add_theme_font_size_override] or via the Inspector dock.
*/
//go:nosplit
func (self class) RemoveThemeFontSizeOverride(name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_remove_theme_font_size_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a local override for a theme [Color] with the specified [param name] previously added by [method add_theme_color_override] or via the Inspector dock.
*/
//go:nosplit
func (self class) RemoveThemeColorOverride(name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_remove_theme_color_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes a local override for a theme constant with the specified [param name] previously added by [method add_theme_constant_override] or via the Inspector dock.
*/
//go:nosplit
func (self class) RemoveThemeConstantOverride(name gd.StringName)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_remove_theme_constant_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns an icon from the first matching [Theme] in the tree if that [Theme] has an icon item with the specified [param name] and [param theme_type].
See [method get_theme_color] for details.
*/
//go:nosplit
func (self class) GetThemeIcon(ctx gd.Lifetime, name gd.StringName, theme_type gd.StringName) object.Texture2D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_theme_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Texture2D
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns a [StyleBox] from the first matching [Theme] in the tree if that [Theme] has a stylebox item with the specified [param name] and [param theme_type].
See [method get_theme_color] for details.
*/
//go:nosplit
func (self class) GetThemeStylebox(ctx gd.Lifetime, name gd.StringName, theme_type gd.StringName) object.StyleBox {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_theme_stylebox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.StyleBox
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns a [Font] from the first matching [Theme] in the tree if that [Theme] has a font item with the specified [param name] and [param theme_type].
See [method get_theme_color] for details.
*/
//go:nosplit
func (self class) GetThemeFont(ctx gd.Lifetime, name gd.StringName, theme_type gd.StringName) object.Font {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_theme_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Font
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns a font size from the first matching [Theme] in the tree if that [Theme] has a font size item with the specified [param name] and [param theme_type].
See [method get_theme_color] for details.
*/
//go:nosplit
func (self class) GetThemeFontSize(name gd.StringName, theme_type gd.StringName) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_theme_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a [Color] from the first matching [Theme] in the tree if that [Theme] has a color item with the specified [param name] and [param theme_type]. If [param theme_type] is omitted the class name of the current control is used as the type, or [member theme_type_variation] if it is defined. If the type is a class name its parent classes are also checked, in order of inheritance. If the type is a variation its base types are checked, in order of dependency, then the control's class name and its parent classes are checked.
For the current control its local overrides are considered first (see [method add_theme_color_override]), then its assigned [member theme]. After the current control, each parent control and its assigned [member theme] are considered; controls without a [member theme] assigned are skipped. If no matching [Theme] is found in the tree, the custom project [Theme] (see [member ProjectSettings.gui/theme/custom]) and the default [Theme] are used (see [ThemeDB]).
[codeblocks]
[gdscript]
func _ready():
    # Get the font color defined for the current Control's class, if it exists.
    modulate = get_theme_color("font_color")
    # Get the font color defined for the Button class.
    modulate = get_theme_color("font_color", "Button")
[/gdscript]
[csharp]
public override void _Ready()
{
    // Get the font color defined for the current Control's class, if it exists.
    Modulate = GetThemeColor("font_color");
    // Get the font color defined for the Button class.
    Modulate = GetThemeColor("font_color", "Button");
}
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) GetThemeColor(name gd.StringName, theme_type gd.StringName) gd.Color {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[gd.Color](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_theme_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a constant from the first matching [Theme] in the tree if that [Theme] has a constant item with the specified [param name] and [param theme_type].
See [method get_theme_color] for details.
*/
//go:nosplit
func (self class) GetThemeConstant(name gd.StringName, theme_type gd.StringName) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_theme_constant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if there is a local override for a theme icon with the specified [param name] in this [Control] node.
See [method add_theme_icon_override].
*/
//go:nosplit
func (self class) HasThemeIconOverride(name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_has_theme_icon_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if there is a local override for a theme [StyleBox] with the specified [param name] in this [Control] node.
See [method add_theme_stylebox_override].
*/
//go:nosplit
func (self class) HasThemeStyleboxOverride(name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_has_theme_stylebox_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if there is a local override for a theme [Font] with the specified [param name] in this [Control] node.
See [method add_theme_font_override].
*/
//go:nosplit
func (self class) HasThemeFontOverride(name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_has_theme_font_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if there is a local override for a theme font size with the specified [param name] in this [Control] node.
See [method add_theme_font_size_override].
*/
//go:nosplit
func (self class) HasThemeFontSizeOverride(name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_has_theme_font_size_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if there is a local override for a theme [Color] with the specified [param name] in this [Control] node.
See [method add_theme_color_override].
*/
//go:nosplit
func (self class) HasThemeColorOverride(name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_has_theme_color_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if there is a local override for a theme constant with the specified [param name] in this [Control] node.
See [method add_theme_constant_override].
*/
//go:nosplit
func (self class) HasThemeConstantOverride(name gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_has_theme_constant_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if there is a matching [Theme] in the tree that has an icon item with the specified [param name] and [param theme_type].
See [method get_theme_color] for details.
*/
//go:nosplit
func (self class) HasThemeIcon(name gd.StringName, theme_type gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_has_theme_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if there is a matching [Theme] in the tree that has a stylebox item with the specified [param name] and [param theme_type].
See [method get_theme_color] for details.
*/
//go:nosplit
func (self class) HasThemeStylebox(name gd.StringName, theme_type gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_has_theme_stylebox, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if there is a matching [Theme] in the tree that has a font item with the specified [param name] and [param theme_type].
See [method get_theme_color] for details.
*/
//go:nosplit
func (self class) HasThemeFont(name gd.StringName, theme_type gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_has_theme_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if there is a matching [Theme] in the tree that has a font size item with the specified [param name] and [param theme_type].
See [method get_theme_color] for details.
*/
//go:nosplit
func (self class) HasThemeFontSize(name gd.StringName, theme_type gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_has_theme_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if there is a matching [Theme] in the tree that has a color item with the specified [param name] and [param theme_type].
See [method get_theme_color] for details.
*/
//go:nosplit
func (self class) HasThemeColor(name gd.StringName, theme_type gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_has_theme_color, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if there is a matching [Theme] in the tree that has a constant item with the specified [param name] and [param theme_type].
See [method get_theme_color] for details.
*/
//go:nosplit
func (self class) HasThemeConstant(name gd.StringName, theme_type gd.StringName) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(theme_type))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_has_theme_constant, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the default base scale value from the first matching [Theme] in the tree if that [Theme] has a valid [member Theme.default_base_scale] value.
See [method get_theme_color] for details.
*/
//go:nosplit
func (self class) GetThemeDefaultBaseScale() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_theme_default_base_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the default font from the first matching [Theme] in the tree if that [Theme] has a valid [member Theme.default_font] value.
See [method get_theme_color] for details.
*/
//go:nosplit
func (self class) GetThemeDefaultFont(ctx gd.Lifetime) object.Font {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_theme_default_font, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Font
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the default font size value from the first matching [Theme] in the tree if that [Theme] has a valid [member Theme.default_font_size] value.
See [method get_theme_color] for details.
*/
//go:nosplit
func (self class) GetThemeDefaultFontSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_theme_default_font_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the parent control node.
*/
//go:nosplit
func (self class) GetParentControl(ctx gd.Lifetime) object.Control {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_parent_control, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Control
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHGrowDirection(direction classdb.ControlGrowDirection)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_h_grow_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHGrowDirection() classdb.ControlGrowDirection {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ControlGrowDirection](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_h_grow_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetVGrowDirection(direction classdb.ControlGrowDirection)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_v_grow_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetVGrowDirection() classdb.ControlGrowDirection {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ControlGrowDirection](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_v_grow_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTooltipText(hint gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(hint))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_tooltip_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTooltipText(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_tooltip_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the tooltip text for the position [param at_position] in control's local coordinates, which will typically appear when the cursor is resting over this control. By default, it returns [member tooltip_text].
This method can be overridden to customize its behavior. See [method _get_tooltip].
[b]Note:[/b] If this method returns an empty [String], no tooltip is displayed.
*/
//go:nosplit
func (self class) GetTooltip(ctx gd.Lifetime, at_position gd.Vector2) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, at_position)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_tooltip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDefaultCursorShape(shape classdb.ControlCursorShape)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, shape)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_default_cursor_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDefaultCursorShape() classdb.ControlCursorShape {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ControlCursorShape](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_default_cursor_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the mouse cursor shape the control displays on mouse hover. See [enum CursorShape].
*/
//go:nosplit
func (self class) GetCursorShape(position gd.Vector2) classdb.ControlCursorShape {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Ret[classdb.ControlCursorShape](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_cursor_shape, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the focus neighbor for the specified [enum Side] to the [Control] at [param neighbor] node path. A setter method for [member focus_neighbor_bottom], [member focus_neighbor_left], [member focus_neighbor_right] and [member focus_neighbor_top].
*/
//go:nosplit
func (self class) SetFocusNeighbor(side gd.Side, neighbor gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, side)
	callframe.Arg(frame, mmm.Get(neighbor))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_focus_neighbor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the focus neighbor for the specified [enum Side]. A getter method for [member focus_neighbor_bottom], [member focus_neighbor_left], [member focus_neighbor_right] and [member focus_neighbor_top].
[b]Note:[/b] To find the next [Control] on the specific [enum Side], even if a neighbor is not assigned, use [method find_valid_focus_neighbor].
*/
//go:nosplit
func (self class) GetFocusNeighbor(ctx gd.Lifetime, side gd.Side) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, side)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_focus_neighbor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFocusNext(next gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(next))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_focus_next, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFocusNext(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_focus_next, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetFocusPrevious(previous gd.NodePath)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(previous))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_focus_previous, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetFocusPrevious(ctx gd.Lifetime) gd.NodePath {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_focus_previous, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.NodePath](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Forces drag and bypasses [method _get_drag_data] and [method set_drag_preview] by passing [param data] and [param preview]. Drag will start even if the mouse is neither over nor pressed on this control.
The methods [method _can_drop_data] and [method _drop_data] must be implemented on controls that want to receive drop data.
*/
//go:nosplit
func (self class) ForceDrag(data gd.Variant, preview object.Control)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(data))
	callframe.Arg(frame, mmm.End(preview[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_force_drag, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetMouseFilter(filter classdb.ControlMouseFilter)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, filter)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_mouse_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMouseFilter() classdb.ControlMouseFilter {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ControlMouseFilter](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_mouse_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetForcePassScrollEvents(force_pass_scroll_events bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, force_pass_scroll_events)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_force_pass_scroll_events, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsForcePassScrollEvents() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_is_force_pass_scroll_events, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetClipContents(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_clip_contents, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsClippingContents() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_is_clipping_contents, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Creates an [InputEventMouseButton] that attempts to click the control. If the event is received, the control acquires focus.
[codeblocks]
[gdscript]
func _process(delta):
    grab_click_focus() # When clicking another Control node, this node will be clicked instead.
[/gdscript]
[csharp]
public override void _Process(double delta)
{
    GrabClickFocus(); // When clicking another Control node, this node will be clicked instead.
}
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) GrabClickFocus()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_grab_click_focus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Forwards the handling of this control's [method _get_drag_data],  [method _can_drop_data] and [method _drop_data] virtual functions to delegate callables.
For each argument, if not empty, the delegate callable is used, otherwise the local (virtual) function is used.
The function format for each callable should be exactly the same as the virtual functions described above.
*/
//go:nosplit
func (self class) SetDragForwarding(drag_func gd.Callable, can_drop_func gd.Callable, drop_func gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(drag_func))
	callframe.Arg(frame, mmm.Get(can_drop_func))
	callframe.Arg(frame, mmm.Get(drop_func))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_drag_forwarding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Shows the given control at the mouse pointer. A good time to call this method is in [method _get_drag_data]. The control must not be in the scene tree. You should not free the control, and you should not keep a reference to the control beyond the duration of the drag. It will be deleted automatically after the drag has ended.
[codeblocks]
[gdscript]
@export var color = Color(1, 0, 0, 1)

func _get_drag_data(position):
    # Use a control that is not in the tree
    var cpb = ColorPickerButton.new()
    cpb.color = color
    cpb.size = Vector2(50, 50)
    set_drag_preview(cpb)
    return color
[/gdscript]
[csharp]
[Export]
private Color _color = new Color(1, 0, 0, 1);

public override Variant _GetDragData(Vector2 atPosition)
{
    // Use a control that is not in the tree
    var cpb = new ColorPickerButton();
    cpb.Color = _color;
    cpb.Size = new Vector2(50, 50);
    SetDragPreview(cpb);
    return _color;
}
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) SetDragPreview(control object.Control)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(control[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_drag_preview, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if a drag operation is successful. Alternative to [method Viewport.gui_is_drag_successful].
Best used with [constant Node.NOTIFICATION_DRAG_END].
*/
//go:nosplit
func (self class) IsDragSuccessful() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_is_drag_successful, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Moves the mouse cursor to [param position], relative to [member position] of this [Control].
[b]Note:[/b] [method warp_mouse] is only supported on Windows, macOS and Linux. It has no effect on Android, iOS and Web.
*/
//go:nosplit
func (self class) WarpMouse(position gd.Vector2)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_warp_mouse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetShortcutContext(node object.Node)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(node[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_shortcut_context, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetShortcutContext(ctx gd.Lifetime) object.Node {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_shortcut_context, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret object.Node
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Invalidates the size cache in this node and in parent nodes up to top level. Intended to be used with [method get_minimum_size] when the return value is changed. Setting [member custom_minimum_size] directly calls this method automatically.
*/
//go:nosplit
func (self class) UpdateMinimumSize()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_update_minimum_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetLayoutDirection(direction classdb.ControlLayoutDirection)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, direction)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_layout_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetLayoutDirection() classdb.ControlLayoutDirection {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ControlLayoutDirection](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_get_layout_direction, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if layout is right-to-left.
*/
//go:nosplit
func (self class) IsLayoutRtl() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_is_layout_rtl, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetAutoTranslate(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_auto_translate, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsAutoTranslating() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_is_auto_translating, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetLocalizeNumeralSystem(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_set_localize_numeral_system, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsLocalizingNumeralSystem() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.Control.Bind_is_localizing_numeral_system, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) AsControl() Expert { return self[0].AsControl() }


//go:nosplit
func (self Simple) AsControl() Simple { return self[0].AsControl() }


//go:nosplit
func (self class) AsCanvasItem() CanvasItem.Expert { return self[0].AsCanvasItem() }


//go:nosplit
func (self Simple) AsCanvasItem() CanvasItem.Simple { return self[0].AsCanvasItem() }


//go:nosplit
func (self class) AsNode() Node.Expert { return self[0].AsNode() }


//go:nosplit
func (self Simple) AsNode() Node.Simple { return self[0].AsNode() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_has_point": return reflect.ValueOf(self._has_point);
	case "_structured_text_parser": return reflect.ValueOf(self._structured_text_parser);
	case "_get_minimum_size": return reflect.ValueOf(self._get_minimum_size);
	case "_get_tooltip": return reflect.ValueOf(self._get_tooltip);
	case "_get_drag_data": return reflect.ValueOf(self._get_drag_data);
	case "_can_drop_data": return reflect.ValueOf(self._can_drop_data);
	case "_drop_data": return reflect.ValueOf(self._drop_data);
	case "_make_custom_tooltip": return reflect.ValueOf(self._make_custom_tooltip);
	case "_gui_input": return reflect.ValueOf(self._gui_input);
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("Control", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
type FocusMode = classdb.ControlFocusMode

const (
/*The node cannot grab focus. Use with [member focus_mode].*/
	FocusNone FocusMode = 0
/*The node can only grab focus on mouse clicks. Use with [member focus_mode].*/
	FocusClick FocusMode = 1
/*The node can grab focus on mouse click, using the arrows and the Tab keys on the keyboard, or using the D-pad buttons on a gamepad. Use with [member focus_mode].*/
	FocusAll FocusMode = 2
)
type CursorShape = classdb.ControlCursorShape

const (
/*Show the system's arrow mouse cursor when the user hovers the node. Use with [member mouse_default_cursor_shape].*/
	CursorArrow CursorShape = 0
/*Show the system's I-beam mouse cursor when the user hovers the node. The I-beam pointer has a shape similar to "I". It tells the user they can highlight or insert text.*/
	CursorIbeam CursorShape = 1
/*Show the system's pointing hand mouse cursor when the user hovers the node.*/
	CursorPointingHand CursorShape = 2
/*Show the system's cross mouse cursor when the user hovers the node.*/
	CursorCross CursorShape = 3
/*Show the system's wait mouse cursor when the user hovers the node. Often an hourglass.*/
	CursorWait CursorShape = 4
/*Show the system's busy mouse cursor when the user hovers the node. Often an arrow with a small hourglass.*/
	CursorBusy CursorShape = 5
/*Show the system's drag mouse cursor, often a closed fist or a cross symbol, when the user hovers the node. It tells the user they're currently dragging an item, like a node in the Scene dock.*/
	CursorDrag CursorShape = 6
/*Show the system's drop mouse cursor when the user hovers the node. It can be an open hand. It tells the user they can drop an item they're currently grabbing, like a node in the Scene dock.*/
	CursorCanDrop CursorShape = 7
/*Show the system's forbidden mouse cursor when the user hovers the node. Often a crossed circle.*/
	CursorForbidden CursorShape = 8
/*Show the system's vertical resize mouse cursor when the user hovers the node. A double-headed vertical arrow. It tells the user they can resize the window or the panel vertically.*/
	CursorVsize CursorShape = 9
/*Show the system's horizontal resize mouse cursor when the user hovers the node. A double-headed horizontal arrow. It tells the user they can resize the window or the panel horizontally.*/
	CursorHsize CursorShape = 10
/*Show the system's window resize mouse cursor when the user hovers the node. The cursor is a double-headed arrow that goes from the bottom left to the top right. It tells the user they can resize the window or the panel both horizontally and vertically.*/
	CursorBdiagsize CursorShape = 11
/*Show the system's window resize mouse cursor when the user hovers the node. The cursor is a double-headed arrow that goes from the top left to the bottom right, the opposite of [constant CURSOR_BDIAGSIZE]. It tells the user they can resize the window or the panel both horizontally and vertically.*/
	CursorFdiagsize CursorShape = 12
/*Show the system's move mouse cursor when the user hovers the node. It shows 2 double-headed arrows at a 90 degree angle. It tells the user they can move a UI element freely.*/
	CursorMove CursorShape = 13
/*Show the system's vertical split mouse cursor when the user hovers the node. On Windows, it's the same as [constant CURSOR_VSIZE].*/
	CursorVsplit CursorShape = 14
/*Show the system's horizontal split mouse cursor when the user hovers the node. On Windows, it's the same as [constant CURSOR_HSIZE].*/
	CursorHsplit CursorShape = 15
/*Show the system's help mouse cursor when the user hovers the node, a question mark.*/
	CursorHelp CursorShape = 16
)
type LayoutPreset = classdb.ControlLayoutPreset

const (
/*Snap all 4 anchors to the top-left of the parent control's bounds. Use with [method set_anchors_preset].*/
	PresetTopLeft LayoutPreset = 0
/*Snap all 4 anchors to the top-right of the parent control's bounds. Use with [method set_anchors_preset].*/
	PresetTopRight LayoutPreset = 1
/*Snap all 4 anchors to the bottom-left of the parent control's bounds. Use with [method set_anchors_preset].*/
	PresetBottomLeft LayoutPreset = 2
/*Snap all 4 anchors to the bottom-right of the parent control's bounds. Use with [method set_anchors_preset].*/
	PresetBottomRight LayoutPreset = 3
/*Snap all 4 anchors to the center of the left edge of the parent control's bounds. Use with [method set_anchors_preset].*/
	PresetCenterLeft LayoutPreset = 4
/*Snap all 4 anchors to the center of the top edge of the parent control's bounds. Use with [method set_anchors_preset].*/
	PresetCenterTop LayoutPreset = 5
/*Snap all 4 anchors to the center of the right edge of the parent control's bounds. Use with [method set_anchors_preset].*/
	PresetCenterRight LayoutPreset = 6
/*Snap all 4 anchors to the center of the bottom edge of the parent control's bounds. Use with [method set_anchors_preset].*/
	PresetCenterBottom LayoutPreset = 7
/*Snap all 4 anchors to the center of the parent control's bounds. Use with [method set_anchors_preset].*/
	PresetCenter LayoutPreset = 8
/*Snap all 4 anchors to the left edge of the parent control. The left offset becomes relative to the left edge and the top offset relative to the top left corner of the node's parent. Use with [method set_anchors_preset].*/
	PresetLeftWide LayoutPreset = 9
/*Snap all 4 anchors to the top edge of the parent control. The left offset becomes relative to the top left corner, the top offset relative to the top edge, and the right offset relative to the top right corner of the node's parent. Use with [method set_anchors_preset].*/
	PresetTopWide LayoutPreset = 10
/*Snap all 4 anchors to the right edge of the parent control. The right offset becomes relative to the right edge and the top offset relative to the top right corner of the node's parent. Use with [method set_anchors_preset].*/
	PresetRightWide LayoutPreset = 11
/*Snap all 4 anchors to the bottom edge of the parent control. The left offset becomes relative to the bottom left corner, the bottom offset relative to the bottom edge, and the right offset relative to the bottom right corner of the node's parent. Use with [method set_anchors_preset].*/
	PresetBottomWide LayoutPreset = 12
/*Snap all 4 anchors to a vertical line that cuts the parent control in half. Use with [method set_anchors_preset].*/
	PresetVcenterWide LayoutPreset = 13
/*Snap all 4 anchors to a horizontal line that cuts the parent control in half. Use with [method set_anchors_preset].*/
	PresetHcenterWide LayoutPreset = 14
/*Snap all 4 anchors to the respective corners of the parent control. Set all 4 offsets to 0 after you applied this preset and the [Control] will fit its parent control. Use with [method set_anchors_preset].*/
	PresetFullRect LayoutPreset = 15
)
type LayoutPresetMode = classdb.ControlLayoutPresetMode

const (
/*The control will be resized to its minimum size.*/
	PresetModeMinsize LayoutPresetMode = 0
/*The control's width will not change.*/
	PresetModeKeepWidth LayoutPresetMode = 1
/*The control's height will not change.*/
	PresetModeKeepHeight LayoutPresetMode = 2
/*The control's size will not change.*/
	PresetModeKeepSize LayoutPresetMode = 3
)
type SizeFlags = classdb.ControlSizeFlags

const (
/*Tells the parent [Container] to align the node with its start, either the top or the left edge. It is mutually exclusive with [constant SIZE_FILL] and other shrink size flags, but can be used with [constant SIZE_EXPAND] in some containers. Use with [member size_flags_horizontal] and [member size_flags_vertical].
[b]Note:[/b] Setting this flag is equal to not having any size flags.*/
	SizeShrinkBegin SizeFlags = 0
/*Tells the parent [Container] to expand the bounds of this node to fill all the available space without pushing any other node. It is mutually exclusive with shrink size flags. Use with [member size_flags_horizontal] and [member size_flags_vertical].*/
	SizeFill SizeFlags = 1
/*Tells the parent [Container] to let this node take all the available space on the axis you flag. If multiple neighboring nodes are set to expand, they'll share the space based on their stretch ratio. See [member size_flags_stretch_ratio]. Use with [member size_flags_horizontal] and [member size_flags_vertical].*/
	SizeExpand SizeFlags = 2
/*Sets the node's size flags to both fill and expand. See [constant SIZE_FILL] and [constant SIZE_EXPAND] for more information.*/
	SizeExpandFill SizeFlags = 3
/*Tells the parent [Container] to center the node in the available space. It is mutually exclusive with [constant SIZE_FILL] and other shrink size flags, but can be used with [constant SIZE_EXPAND] in some containers. Use with [member size_flags_horizontal] and [member size_flags_vertical].*/
	SizeShrinkCenter SizeFlags = 4
/*Tells the parent [Container] to align the node with its end, either the bottom or the right edge. It is mutually exclusive with [constant SIZE_FILL] and other shrink size flags, but can be used with [constant SIZE_EXPAND] in some containers. Use with [member size_flags_horizontal] and [member size_flags_vertical].*/
	SizeShrinkEnd SizeFlags = 8
)
type MouseFilter = classdb.ControlMouseFilter

const (
/*The control will receive mouse movement input events and mouse button input events if clicked on through [method _gui_input]. And the control will receive the [signal mouse_entered] and [signal mouse_exited] signals. These events are automatically marked as handled, and they will not propagate further to other controls. This also results in blocking signals in other controls.*/
	MouseFilterStop MouseFilter = 0
/*The control will receive mouse movement input events and mouse button input events if clicked on through [method _gui_input]. And the control will receive the [signal mouse_entered] and [signal mouse_exited] signals. If this control does not handle the event, the parent control (if any) will be considered, and so on until there is no more parent control to potentially handle it. This also allows signals to fire in other controls. If no control handled it, the event will be passed to [method Node._shortcut_input] for further processing.*/
	MouseFilterPass MouseFilter = 1
/*The control will not receive mouse movement input events and mouse button input events if clicked on through [method _gui_input]. The control will also not receive the [signal mouse_entered] nor [signal mouse_exited] signals. This will not block other controls from receiving these events or firing the signals. Ignored events will not be handled automatically.
[b]Note:[/b] If the control has received [signal mouse_entered] but not [signal mouse_exited], changing the [member mouse_filter] to [constant MOUSE_FILTER_IGNORE] will cause [signal mouse_exited] to be emitted.*/
	MouseFilterIgnore MouseFilter = 2
)
type GrowDirection = classdb.ControlGrowDirection

const (
/*The control will grow to the left or top to make up if its minimum size is changed to be greater than its current size on the respective axis.*/
	GrowDirectionBegin GrowDirection = 0
/*The control will grow to the right or bottom to make up if its minimum size is changed to be greater than its current size on the respective axis.*/
	GrowDirectionEnd GrowDirection = 1
/*The control will grow in both directions equally to make up if its minimum size is changed to be greater than its current size.*/
	GrowDirectionBoth GrowDirection = 2
)
type Anchor = classdb.ControlAnchor

const (
/*Snaps one of the 4 anchor's sides to the origin of the node's [code]Rect[/code], in the top left. Use it with one of the [code]anchor_*[/code] member variables, like [member anchor_left]. To change all 4 anchors at once, use [method set_anchors_preset].*/
	AnchorBegin Anchor = 0
/*Snaps one of the 4 anchor's sides to the end of the node's [code]Rect[/code], in the bottom right. Use it with one of the [code]anchor_*[/code] member variables, like [member anchor_left]. To change all 4 anchors at once, use [method set_anchors_preset].*/
	AnchorEnd Anchor = 1
)
type LayoutDirection = classdb.ControlLayoutDirection

const (
/*Automatic layout direction, determined from the parent control layout direction.*/
	LayoutDirectionInherited LayoutDirection = 0
/*Automatic layout direction, determined from the current locale.*/
	LayoutDirectionLocale LayoutDirection = 1
/*Left-to-right layout direction.*/
	LayoutDirectionLtr LayoutDirection = 2
/*Right-to-left layout direction.*/
	LayoutDirectionRtl LayoutDirection = 3
)
type TextDirection = classdb.ControlTextDirection

const (
/*Text writing direction is the same as layout direction.*/
	TextDirectionInherited TextDirection = 3
/*Automatic text writing direction, determined from the current locale and text content.*/
	TextDirectionAuto TextDirection = 0
/*Left-to-right text writing direction.*/
	TextDirectionLtr TextDirection = 1
/*Right-to-left text writing direction.*/
	TextDirectionRtl TextDirection = 2
)
