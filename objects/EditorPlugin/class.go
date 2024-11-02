package EditorPlugin

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/objects/Node"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
Plugins are used by the editor to extend functionality. The most common types of plugins are those which edit a given node or resource type, import plugins and export plugins. See also [EditorScript] to add functions to the editor.
[b]Note:[/b] Some names in this class contain "left" or "right" (e.g. [constant DOCK_SLOT_LEFT_UL]). These APIs assume left-to-right layout, and would be backwards when using right-to-left layout. These names are kept for compatibility reasons.

	// EditorPlugin methods that can be overridden by a [Class] that extends it.
	type EditorPlugin interface {
		//Called when there is a root node in the current edited scene, [method _handles] is implemented and an [InputEvent] happens in the 2D viewport. Intercepts the [InputEvent], if [code]return true[/code] [EditorPlugin] consumes the [param event], otherwise forwards [param event] to other Editor classes.
		//[b]Example:[/b]
		//[codeblocks]
		//[gdscript]
		//# Prevents the InputEvent from reaching other Editor classes.
		//func _forward_canvas_gui_input(event):
		//    return true
		//[/gdscript]
		//[csharp]
		//// Prevents the InputEvent from reaching other Editor classes.
		//public override bool ForwardCanvasGuiInput(InputEvent @event)
		//{
		//    return true;
		//}
		//[/csharp]
		//[/codeblocks]
		//Must [code]return false[/code] in order to forward the [InputEvent] to other Editor classes.
		//[b]Example:[/b]
		//[codeblocks]
		//[gdscript]
		//# Consumes InputEventMouseMotion and forwards other InputEvent types.
		//func _forward_canvas_gui_input(event):
		//    if (event is InputEventMouseMotion):
		//        return true
		//    return false
		//[/gdscript]
		//[csharp]
		//// Consumes InputEventMouseMotion and forwards other InputEvent types.
		//public override bool _ForwardCanvasGuiInput(InputEvent @event)
		//{
		//    if (@event is InputEventMouseMotion)
		//    {
		//        return true;
		//    }
		//    return false;
		//}
		//[/csharp]
		//[/codeblocks]
		ForwardCanvasGuiInput(event objects.InputEvent) bool
		//Called by the engine when the 2D editor's viewport is updated. Use the [code]overlay[/code] [Control] for drawing. You can update the viewport manually by calling [method update_overlays].
		//[codeblocks]
		//[gdscript]
		//func _forward_canvas_draw_over_viewport(overlay):
		//    # Draw a circle at cursor position.
		//    overlay.draw_circle(overlay.get_local_mouse_position(), 64, Color.WHITE)
		//
		//func _forward_canvas_gui_input(event):
		//    if event is InputEventMouseMotion:
		//        # Redraw viewport when cursor is moved.
		//        update_overlays()
		//        return true
		//    return false
		//[/gdscript]
		//[csharp]
		//public override void _ForwardCanvasDrawOverViewport(Control viewportControl)
		//{
		//    // Draw a circle at cursor position.
		//    viewportControl.DrawCircle(viewportControl.GetLocalMousePosition(), 64, Colors.White);
		//}
		//
		//public override bool _ForwardCanvasGuiInput(InputEvent @event)
		//{
		//    if (@event is InputEventMouseMotion)
		//    {
		//        // Redraw viewport when cursor is moved.
		//        UpdateOverlays();
		//        return true;
		//    }
		//    return false;
		//}
		//[/csharp]
		//[/codeblocks]
		ForwardCanvasDrawOverViewport(viewport_control objects.Control)
		//This method is the same as [method _forward_canvas_draw_over_viewport], except it draws on top of everything. Useful when you need an extra layer that shows over anything else.
		//You need to enable calling of this method by using [method set_force_draw_over_forwarding_enabled].
		ForwardCanvasForceDrawOverViewport(viewport_control objects.Control)
		//Called when there is a root node in the current edited scene, [method _handles] is implemented, and an [InputEvent] happens in the 3D viewport. The return value decides whether the [InputEvent] is consumed or forwarded to other [EditorPlugin]s. See [enum AfterGUIInput] for options.
		//[b]Example:[/b]
		//[codeblocks]
		//[gdscript]
		//# Prevents the InputEvent from reaching other Editor classes.
		//func _forward_3d_gui_input(camera, event):
		//    return EditorPlugin.AFTER_GUI_INPUT_STOP
		//[/gdscript]
		//[csharp]
		//// Prevents the InputEvent from reaching other Editor classes.
		//public override EditorPlugin.AfterGuiInput _Forward3DGuiInput(Camera3D camera, InputEvent @event)
		//{
		//    return EditorPlugin.AfterGuiInput.Stop;
		//}
		//[/csharp]
		//[/codeblocks]
		//Must [code]return EditorPlugin.AFTER_GUI_INPUT_PASS[/code] in order to forward the [InputEvent] to other Editor classes.
		//[b]Example:[/b]
		//[codeblocks]
		//[gdscript]
		//# Consumes InputEventMouseMotion and forwards other InputEvent types.
		//func _forward_3d_gui_input(camera, event):
		//    return EditorPlugin.AFTER_GUI_INPUT_STOP if event is InputEventMouseMotion else EditorPlugin.AFTER_GUI_INPUT_PASS
		//[/gdscript]
		//[csharp]
		//// Consumes InputEventMouseMotion and forwards other InputEvent types.
		//public override EditorPlugin.AfterGuiInput _Forward3DGuiInput(Camera3D camera, InputEvent @event)
		//{
		//    return @event is InputEventMouseMotion ? EditorPlugin.AfterGuiInput.Stop : EditorPlugin.AfterGuiInput.Pass;
		//}
		//[/csharp]
		//[/codeblocks]
		Forward3dGuiInput(viewport_camera objects.Camera3D, event objects.InputEvent) int
		//Called by the engine when the 3D editor's viewport is updated. Use the [code]overlay[/code] [Control] for drawing. You can update the viewport manually by calling [method update_overlays].
		//[codeblocks]
		//[gdscript]
		//func _forward_3d_draw_over_viewport(overlay):
		//    # Draw a circle at cursor position.
		//    overlay.draw_circle(overlay.get_local_mouse_position(), 64, Color.WHITE)
		//
		//func _forward_3d_gui_input(camera, event):
		//    if event is InputEventMouseMotion:
		//        # Redraw viewport when cursor is moved.
		//        update_overlays()
		//        return EditorPlugin.AFTER_GUI_INPUT_STOP
		//    return EditorPlugin.AFTER_GUI_INPUT_PASS
		//[/gdscript]
		//[csharp]
		//public override void _Forward3DDrawOverViewport(Control viewportControl)
		//{
		//    // Draw a circle at cursor position.
		//    viewportControl.DrawCircle(viewportControl.GetLocalMousePosition(), 64, Colors.White);
		//}
		//
		//public override EditorPlugin.AfterGuiInput _Forward3DGuiInput(Camera3D viewportCamera, InputEvent @event)
		//{
		//    if (@event is InputEventMouseMotion)
		//    {
		//        // Redraw viewport when cursor is moved.
		//        UpdateOverlays();
		//        return EditorPlugin.AfterGuiInput.Stop;
		//    }
		//    return EditorPlugin.AfterGuiInput.Pass;
		//}
		//[/csharp]
		//[/codeblocks]
		Forward3dDrawOverViewport(viewport_control objects.Control)
		//This method is the same as [method _forward_3d_draw_over_viewport], except it draws on top of everything. Useful when you need an extra layer that shows over anything else.
		//You need to enable calling of this method by using [method set_force_draw_over_forwarding_enabled].
		Forward3dForceDrawOverViewport(viewport_control objects.Control)
		//Override this method in your plugin to provide the name of the plugin when displayed in the Godot editor.
		//For main screen plugins, this appears at the top of the screen, to the right of the "2D", "3D", "Script", and "AssetLib" buttons.
		GetPluginName() string
		//Override this method in your plugin to return a [Texture2D] in order to give it an icon.
		//For main screen plugins, this appears at the top of the screen, to the right of the "2D", "3D", "Script", and "AssetLib" buttons.
		//Ideally, the plugin icon should be white with a transparent background and 16×16 pixels in size.
		//[codeblocks]
		//[gdscript]
		//func _get_plugin_icon():
		//    # You can use a custom icon:
		//    return preload("res://addons/my_plugin/my_plugin_icon.svg")
		//    # Or use a built-in icon:
		//    return EditorInterface.get_editor_theme().get_icon("Node", "EditorIcons")
		//[/gdscript]
		//[csharp]
		//public override Texture2D _GetPluginIcon()
		//{
		//    // You can use a custom icon:
		//    return ResourceLoader.Load<Texture2D>("res://addons/my_plugin/my_plugin_icon.svg");
		//    // Or use a built-in icon:
		//    return EditorInterface.Singleton.GetEditorTheme().GetIcon("Node", "EditorIcons");
		//}
		//[/csharp]
		//[/codeblocks]
		GetPluginIcon() objects.Texture2D
		//Returns [code]true[/code] if this is a main screen editor plugin (it goes in the workspace selector together with [b]2D[/b], [b]3D[/b], [b]Script[/b] and [b]AssetLib[/b]).
		//When the plugin's workspace is selected, other main screen plugins will be hidden, but your plugin will not appear automatically. It needs to be added as a child of [method EditorInterface.get_editor_main_screen] and made visible inside [method _make_visible].
		//Use [method _get_plugin_name] and [method _get_plugin_icon] to customize the plugin button's appearance.
		//[codeblock]
		//var plugin_control
		//
		//func _enter_tree():
		//    plugin_control = preload("my_plugin_control.tscn").instantiate()
		//    EditorInterface.get_editor_main_screen().add_child(plugin_control)
		//    plugin_control.hide()
		//
		//func _has_main_screen():
		//    return true
		//
		//func _make_visible(visible):
		//    plugin_control.visible = visible
		//
		//func _get_plugin_name():
		//    return "My Super Cool Plugin 3000"
		//
		//func _get_plugin_icon():
		//    return EditorInterface.get_editor_theme().get_icon("Node", "EditorIcons")
		//[/codeblock]
		HasMainScreen() bool
		//This function will be called when the editor is requested to become visible. It is used for plugins that edit a specific object type.
		//Remember that you have to manage the visibility of all your editor controls manually.
		MakeVisible(visible bool)
		//This function is used for plugins that edit specific object types (nodes or resources). It requests the editor to edit the given object.
		//[param object] can be [code]null[/code] if the plugin was editing an object, but there is no longer any selected object handled by this plugin. It can be used to cleanup editing state.
		Edit(obj gd.Object)
		//Implement this function if your plugin edits a specific type of object (Resource or Node). If you return [code]true[/code], then you will get the functions [method _edit] and [method _make_visible] called when the editor requests them. If you have declared the methods [method _forward_canvas_gui_input] and [method _forward_3d_gui_input] these will be called too.
		//[b]Note:[/b] Each plugin should handle only one type of objects at a time. If a plugin handles more types of objects and they are edited at the same time, it will result in errors.
		Handles(obj gd.Object) bool
		//Override this method to provide a state data you want to be saved, like view position, grid settings, folding, etc. This is used when saving the scene (so state is kept when opening it again) and for switching tabs (so state can be restored when the tab returns). This data is automatically saved for each scene in an [code]editstate[/code] file in the editor metadata folder. If you want to store global (scene-independent) editor data for your plugin, you can use [method _get_window_layout] instead.
		//Use [method _set_state] to restore your saved state.
		//[b]Note:[/b] This method should not be used to save important settings that should persist with the project.
		//[b]Note:[/b] You must implement [method _get_plugin_name] for the state to be stored and restored correctly.
		//[codeblock]
		//func _get_state():
		//    var state = {"zoom": zoom, "preferred_color": my_color}
		//    return state
		//[/codeblock]
		GetState() gd.Dictionary
		//Restore the state saved by [method _get_state]. This method is called when the current scene tab is changed in the editor.
		//[b]Note:[/b] Your plugin must implement [method _get_plugin_name], otherwise it will not be recognized and this method will not be called.
		//[codeblock]
		//func _set_state(data):
		//    zoom = data.get("zoom", 1.0)
		//    preferred_color = data.get("my_color", Color.WHITE)
		//[/codeblock]
		SetState(state gd.Dictionary)
		//Clear all the state and reset the object being edited to zero. This ensures your plugin does not keep editing a currently existing node, or a node from the wrong scene.
		Clear()
		//Override this method to provide a custom message that lists unsaved changes. The editor will call this method when exiting or when closing a scene, and display the returned string in a confirmation dialog. Return empty string if the plugin has no unsaved changes.
		//When closing a scene, [param for_scene] is the path to the scene being closed. You can use it to handle built-in resources in that scene.
		//If the user confirms saving, [method _save_external_data] will be called, before closing the editor.
		//[codeblock]
		//func _get_unsaved_status(for_scene):
		//    if not unsaved:
		//        return ""
		//
		//    if for_scene.is_empty():
		//        return "Save changes in MyCustomPlugin before closing?"
		//    else:
		//        return "Scene %s has changes from MyCustomPlugin. Save before closing?" % for_scene.get_file()
		//
		//func _save_external_data():
		//    unsaved = false
		//[/codeblock]
		//If the plugin has no scene-specific changes, you can ignore the calls when closing scenes:
		//[codeblock]
		//func _get_unsaved_status(for_scene):
		//    if not for_scene.is_empty():
		//        return ""
		//[/codeblock]
		GetUnsavedStatus(for_scene string) string
		//This method is called after the editor saves the project or when it's closed. It asks the plugin to save edited external scenes/resources.
		SaveExternalData()
		//This method is called when the editor is about to save the project, switch to another tab, etc. It asks the plugin to apply any pending state changes to ensure consistency.
		//This is used, for example, in shader editors to let the plugin know that it must apply the shader code being written by the user to the object.
		ApplyChanges()
		//This is for editors that edit script-based objects. You can return a list of breakpoints in the format ([code]script:line[/code]), for example: [code]res://path_to_script.gd:25[/code].
		GetBreakpoints() []string
		//Restore the plugin GUI layout and data saved by [method _get_window_layout]. This method is called for every plugin on editor startup. Use the provided [param configuration] file to read your saved data.
		//[codeblock]
		//func _set_window_layout(configuration):
		//    $Window.position = configuration.get_value("MyPlugin", "window_position", Vector2())
		//    $Icon.modulate = configuration.get_value("MyPlugin", "icon_color", Color.WHITE)
		//[/codeblock]
		SetWindowLayout(configuration objects.ConfigFile)
		//Override this method to provide the GUI layout of the plugin or any other data you want to be stored. This is used to save the project's editor layout when [method queue_save_layout] is called or the editor layout was changed (for example changing the position of a dock). The data is stored in the [code]editor_layout.cfg[/code] file in the editor metadata directory.
		//Use [method _set_window_layout] to restore your saved layout.
		//[codeblock]
		//func _get_window_layout(configuration):
		//    configuration.set_value("MyPlugin", "window_position", $Window.position)
		//    configuration.set_value("MyPlugin", "icon_color", $Icon.modulate)
		//[/codeblock]
		GetWindowLayout(configuration objects.ConfigFile)
		//This method is called when the editor is about to run the project. The plugin can then perform required operations before the project runs.
		//This method must return a boolean. If this method returns [code]false[/code], the project will not run. The run is aborted immediately, so this also prevents all other plugins' [method _build] methods from running.
		Build() bool
		//Called by the engine when the user enables the [EditorPlugin] in the Plugin tab of the project settings window.
		EnablePlugin()
		//Called by the engine when the user disables the [EditorPlugin] in the Plugin tab of the project settings window.
		DisablePlugin()
	}
*/
type Instance [1]classdb.EditorPlugin

/*
Called when there is a root node in the current edited scene, [method _handles] is implemented and an [InputEvent] happens in the 2D viewport. Intercepts the [InputEvent], if [code]return true[/code] [EditorPlugin] consumes the [param event], otherwise forwards [param event] to other Editor classes.
[b]Example:[/b]
[codeblocks]
[gdscript]
# Prevents the InputEvent from reaching other Editor classes.
func _forward_canvas_gui_input(event):

	return true

[/gdscript]
[csharp]
// Prevents the InputEvent from reaching other Editor classes.
public override bool ForwardCanvasGuiInput(InputEvent @event)

	{
	    return true;
	}

[/csharp]
[/codeblocks]
Must [code]return false[/code] in order to forward the [InputEvent] to other Editor classes.
[b]Example:[/b]
[codeblocks]
[gdscript]
# Consumes InputEventMouseMotion and forwards other InputEvent types.
func _forward_canvas_gui_input(event):

	if (event is InputEventMouseMotion):
	    return true
	return false

[/gdscript]
[csharp]
// Consumes InputEventMouseMotion and forwards other InputEvent types.
public override bool _ForwardCanvasGuiInput(InputEvent @event)

	{
	    if (@event is InputEventMouseMotion)
	    {
	        return true;
	    }
	    return false;
	}

[/csharp]
[/codeblocks]
*/
func (Instance) _forward_canvas_gui_input(impl func(ptr unsafe.Pointer, event objects.InputEvent) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var event = objects.InputEvent{pointers.New[classdb.InputEvent]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(event[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, event)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called by the engine when the 2D editor's viewport is updated. Use the [code]overlay[/code] [Control] for drawing. You can update the viewport manually by calling [method update_overlays].
[codeblocks]
[gdscript]
func _forward_canvas_draw_over_viewport(overlay):

	# Draw a circle at cursor position.
	overlay.draw_circle(overlay.get_local_mouse_position(), 64, Color.WHITE)

func _forward_canvas_gui_input(event):

	if event is InputEventMouseMotion:
	    # Redraw viewport when cursor is moved.
	    update_overlays()
	    return true
	return false

[/gdscript]
[csharp]
public override void _ForwardCanvasDrawOverViewport(Control viewportControl)

	{
	    // Draw a circle at cursor position.
	    viewportControl.DrawCircle(viewportControl.GetLocalMousePosition(), 64, Colors.White);
	}

public override bool _ForwardCanvasGuiInput(InputEvent @event)

	{
	    if (@event is InputEventMouseMotion)
	    {
	        // Redraw viewport when cursor is moved.
	        UpdateOverlays();
	        return true;
	    }
	    return false;
	}

[/csharp]
[/codeblocks]
*/
func (Instance) _forward_canvas_draw_over_viewport(impl func(ptr unsafe.Pointer, viewport_control objects.Control)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var viewport_control = objects.Control{pointers.New[classdb.Control]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(viewport_control[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, viewport_control)
	}
}

/*
This method is the same as [method _forward_canvas_draw_over_viewport], except it draws on top of everything. Useful when you need an extra layer that shows over anything else.
You need to enable calling of this method by using [method set_force_draw_over_forwarding_enabled].
*/
func (Instance) _forward_canvas_force_draw_over_viewport(impl func(ptr unsafe.Pointer, viewport_control objects.Control)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var viewport_control = objects.Control{pointers.New[classdb.Control]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(viewport_control[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, viewport_control)
	}
}

/*
Called when there is a root node in the current edited scene, [method _handles] is implemented, and an [InputEvent] happens in the 3D viewport. The return value decides whether the [InputEvent] is consumed or forwarded to other [EditorPlugin]s. See [enum AfterGUIInput] for options.
[b]Example:[/b]
[codeblocks]
[gdscript]
# Prevents the InputEvent from reaching other Editor classes.
func _forward_3d_gui_input(camera, event):

	return EditorPlugin.AFTER_GUI_INPUT_STOP

[/gdscript]
[csharp]
// Prevents the InputEvent from reaching other Editor classes.
public override EditorPlugin.AfterGuiInput _Forward3DGuiInput(Camera3D camera, InputEvent @event)

	{
	    return EditorPlugin.AfterGuiInput.Stop;
	}

[/csharp]
[/codeblocks]
Must [code]return EditorPlugin.AFTER_GUI_INPUT_PASS[/code] in order to forward the [InputEvent] to other Editor classes.
[b]Example:[/b]
[codeblocks]
[gdscript]
# Consumes InputEventMouseMotion and forwards other InputEvent types.
func _forward_3d_gui_input(camera, event):

	return EditorPlugin.AFTER_GUI_INPUT_STOP if event is InputEventMouseMotion else EditorPlugin.AFTER_GUI_INPUT_PASS

[/gdscript]
[csharp]
// Consumes InputEventMouseMotion and forwards other InputEvent types.
public override EditorPlugin.AfterGuiInput _Forward3DGuiInput(Camera3D camera, InputEvent @event)

	{
	    return @event is InputEventMouseMotion ? EditorPlugin.AfterGuiInput.Stop : EditorPlugin.AfterGuiInput.Pass;
	}

[/csharp]
[/codeblocks]
*/
func (Instance) _forward_3d_gui_input(impl func(ptr unsafe.Pointer, viewport_camera objects.Camera3D, event objects.InputEvent) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var viewport_camera = objects.Camera3D{pointers.New[classdb.Camera3D]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(viewport_camera[0])
		var event = objects.InputEvent{pointers.New[classdb.InputEvent]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
		defer pointers.End(event[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, viewport_camera, event)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Called by the engine when the 3D editor's viewport is updated. Use the [code]overlay[/code] [Control] for drawing. You can update the viewport manually by calling [method update_overlays].
[codeblocks]
[gdscript]
func _forward_3d_draw_over_viewport(overlay):

	# Draw a circle at cursor position.
	overlay.draw_circle(overlay.get_local_mouse_position(), 64, Color.WHITE)

func _forward_3d_gui_input(camera, event):

	if event is InputEventMouseMotion:
	    # Redraw viewport when cursor is moved.
	    update_overlays()
	    return EditorPlugin.AFTER_GUI_INPUT_STOP
	return EditorPlugin.AFTER_GUI_INPUT_PASS

[/gdscript]
[csharp]
public override void _Forward3DDrawOverViewport(Control viewportControl)

	{
	    // Draw a circle at cursor position.
	    viewportControl.DrawCircle(viewportControl.GetLocalMousePosition(), 64, Colors.White);
	}

public override EditorPlugin.AfterGuiInput _Forward3DGuiInput(Camera3D viewportCamera, InputEvent @event)

	{
	    if (@event is InputEventMouseMotion)
	    {
	        // Redraw viewport when cursor is moved.
	        UpdateOverlays();
	        return EditorPlugin.AfterGuiInput.Stop;
	    }
	    return EditorPlugin.AfterGuiInput.Pass;
	}

[/csharp]
[/codeblocks]
*/
func (Instance) _forward_3d_draw_over_viewport(impl func(ptr unsafe.Pointer, viewport_control objects.Control)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var viewport_control = objects.Control{pointers.New[classdb.Control]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(viewport_control[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, viewport_control)
	}
}

/*
This method is the same as [method _forward_3d_draw_over_viewport], except it draws on top of everything. Useful when you need an extra layer that shows over anything else.
You need to enable calling of this method by using [method set_force_draw_over_forwarding_enabled].
*/
func (Instance) _forward_3d_force_draw_over_viewport(impl func(ptr unsafe.Pointer, viewport_control objects.Control)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var viewport_control = objects.Control{pointers.New[classdb.Control]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(viewport_control[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, viewport_control)
	}
}

/*
Override this method in your plugin to provide the name of the plugin when displayed in the Godot editor.
For main screen plugins, this appears at the top of the screen, to the right of the "2D", "3D", "Script", and "AssetLib" buttons.
*/
func (Instance) _get_plugin_name(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewString(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method in your plugin to return a [Texture2D] in order to give it an icon.
For main screen plugins, this appears at the top of the screen, to the right of the "2D", "3D", "Script", and "AssetLib" buttons.
Ideally, the plugin icon should be white with a transparent background and 16×16 pixels in size.
[codeblocks]
[gdscript]
func _get_plugin_icon():

	# You can use a custom icon:
	return preload("res://addons/my_plugin/my_plugin_icon.svg")
	# Or use a built-in icon:
	return EditorInterface.get_editor_theme().get_icon("Node", "EditorIcons")

[/gdscript]
[csharp]
public override Texture2D _GetPluginIcon()

	{
	    // You can use a custom icon:
	    return ResourceLoader.Load<Texture2D>("res://addons/my_plugin/my_plugin_icon.svg");
	    // Or use a built-in icon:
	    return EditorInterface.Singleton.GetEditorTheme().GetIcon("Node", "EditorIcons");
	}

[/csharp]
[/codeblocks]
*/
func (Instance) _get_plugin_icon(impl func(ptr unsafe.Pointer) objects.Texture2D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns [code]true[/code] if this is a main screen editor plugin (it goes in the workspace selector together with [b]2D[/b], [b]3D[/b], [b]Script[/b] and [b]AssetLib[/b]).
When the plugin's workspace is selected, other main screen plugins will be hidden, but your plugin will not appear automatically. It needs to be added as a child of [method EditorInterface.get_editor_main_screen] and made visible inside [method _make_visible].
Use [method _get_plugin_name] and [method _get_plugin_icon] to customize the plugin button's appearance.
[codeblock]
var plugin_control

func _enter_tree():

	plugin_control = preload("my_plugin_control.tscn").instantiate()
	EditorInterface.get_editor_main_screen().add_child(plugin_control)
	plugin_control.hide()

func _has_main_screen():

	return true

func _make_visible(visible):

	plugin_control.visible = visible

func _get_plugin_name():

	return "My Super Cool Plugin 3000"

func _get_plugin_icon():

	return EditorInterface.get_editor_theme().get_icon("Node", "EditorIcons")

[/codeblock]
*/
func (Instance) _has_main_screen(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
This function will be called when the editor is requested to become visible. It is used for plugins that edit a specific object type.
Remember that you have to manage the visibility of all your editor controls manually.
*/
func (Instance) _make_visible(impl func(ptr unsafe.Pointer, visible bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var visible = gd.UnsafeGet[bool](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, visible)
	}
}

/*
This function is used for plugins that edit specific object types (nodes or resources). It requests the editor to edit the given object.
[param object] can be [code]null[/code] if the plugin was editing an object, but there is no longer any selected object handled by this plugin. It can be used to cleanup editing state.
*/
func (Instance) _edit(impl func(ptr unsafe.Pointer, obj gd.Object)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})
		defer pointers.End(obj)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, obj)
	}
}

/*
Implement this function if your plugin edits a specific type of object (Resource or Node). If you return [code]true[/code], then you will get the functions [method _edit] and [method _make_visible] called when the editor requests them. If you have declared the methods [method _forward_canvas_gui_input] and [method _forward_3d_gui_input] these will be called too.
[b]Note:[/b] Each plugin should handle only one type of objects at a time. If a plugin handles more types of objects and they are edited at the same time, it will result in errors.
*/
func (Instance) _handles(impl func(ptr unsafe.Pointer, obj gd.Object) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})
		defer pointers.End(obj)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to provide a state data you want to be saved, like view position, grid settings, folding, etc. This is used when saving the scene (so state is kept when opening it again) and for switching tabs (so state can be restored when the tab returns). This data is automatically saved for each scene in an [code]editstate[/code] file in the editor metadata folder. If you want to store global (scene-independent) editor data for your plugin, you can use [method _get_window_layout] instead.
Use [method _set_state] to restore your saved state.
[b]Note:[/b] This method should not be used to save important settings that should persist with the project.
[b]Note:[/b] You must implement [method _get_plugin_name] for the state to be stored and restored correctly.
[codeblock]
func _get_state():

	var state = {"zoom": zoom, "preferred_color": my_color}
	return state

[/codeblock]
*/
func (Instance) _get_state(impl func(ptr unsafe.Pointer) gd.Dictionary) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Restore the state saved by [method _get_state]. This method is called when the current scene tab is changed in the editor.
[b]Note:[/b] Your plugin must implement [method _get_plugin_name], otherwise it will not be recognized and this method will not be called.
[codeblock]
func _set_state(data):

	zoom = data.get("zoom", 1.0)
	preferred_color = data.get("my_color", Color.WHITE)

[/codeblock]
*/
func (Instance) _set_state(impl func(ptr unsafe.Pointer, state gd.Dictionary)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(state)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, state)
	}
}

/*
Clear all the state and reset the object being edited to zero. This ensures your plugin does not keep editing a currently existing node, or a node from the wrong scene.
*/
func (Instance) _clear(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Override this method to provide a custom message that lists unsaved changes. The editor will call this method when exiting or when closing a scene, and display the returned string in a confirmation dialog. Return empty string if the plugin has no unsaved changes.
When closing a scene, [param for_scene] is the path to the scene being closed. You can use it to handle built-in resources in that scene.
If the user confirms saving, [method _save_external_data] will be called, before closing the editor.
[codeblock]
func _get_unsaved_status(for_scene):

	if not unsaved:
	    return ""

	if for_scene.is_empty():
	    return "Save changes in MyCustomPlugin before closing?"
	else:
	    return "Scene %s has changes from MyCustomPlugin. Save before closing?" % for_scene.get_file()

func _save_external_data():

	unsaved = false

[/codeblock]
If the plugin has no scene-specific changes, you can ignore the calls when closing scenes:
[codeblock]
func _get_unsaved_status(for_scene):

	if not for_scene.is_empty():
	    return ""

[/codeblock]
*/
func (Instance) _get_unsaved_status(impl func(ptr unsafe.Pointer, for_scene string) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var for_scene = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		defer pointers.End(for_scene)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, for_scene.String())
		ptr, ok := pointers.End(gd.NewString(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
This method is called after the editor saves the project or when it's closed. It asks the plugin to save edited external scenes/resources.
*/
func (Instance) _save_external_data(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
This method is called when the editor is about to save the project, switch to another tab, etc. It asks the plugin to apply any pending state changes to ensure consistency.
This is used, for example, in shader editors to let the plugin know that it must apply the shader code being written by the user to the object.
*/
func (Instance) _apply_changes(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
This is for editors that edit script-based objects. You can return a list of breakpoints in the format ([code]script:line[/code]), for example: [code]res://path_to_script.gd:25[/code].
*/
func (Instance) _get_breakpoints(impl func(ptr unsafe.Pointer) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.NewPackedStringSlice(ret))
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Restore the plugin GUI layout and data saved by [method _get_window_layout]. This method is called for every plugin on editor startup. Use the provided [param configuration] file to read your saved data.
[codeblock]
func _set_window_layout(configuration):

	$Window.position = configuration.get_value("MyPlugin", "window_position", Vector2())
	$Icon.modulate = configuration.get_value("MyPlugin", "icon_color", Color.WHITE)

[/codeblock]
*/
func (Instance) _set_window_layout(impl func(ptr unsafe.Pointer, configuration objects.ConfigFile)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var configuration = objects.ConfigFile{pointers.New[classdb.ConfigFile]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(configuration[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, configuration)
	}
}

/*
Override this method to provide the GUI layout of the plugin or any other data you want to be stored. This is used to save the project's editor layout when [method queue_save_layout] is called or the editor layout was changed (for example changing the position of a dock). The data is stored in the [code]editor_layout.cfg[/code] file in the editor metadata directory.
Use [method _set_window_layout] to restore your saved layout.
[codeblock]
func _get_window_layout(configuration):

	configuration.set_value("MyPlugin", "window_position", $Window.position)
	configuration.set_value("MyPlugin", "icon_color", $Icon.modulate)

[/codeblock]
*/
func (Instance) _get_window_layout(impl func(ptr unsafe.Pointer, configuration objects.ConfigFile)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var configuration = objects.ConfigFile{pointers.New[classdb.ConfigFile]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(configuration[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, configuration)
	}
}

/*
This method is called when the editor is about to run the project. The plugin can then perform required operations before the project runs.
This method must return a boolean. If this method returns [code]false[/code], the project will not run. The run is aborted immediately, so this also prevents all other plugins' [method _build] methods from running.
*/
func (Instance) _build(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called by the engine when the user enables the [EditorPlugin] in the Plugin tab of the project settings window.
*/
func (Instance) _enable_plugin(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called by the engine when the user disables the [EditorPlugin] in the Plugin tab of the project settings window.
*/
func (Instance) _disable_plugin(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Adds a custom control to a container (see [enum CustomControlContainer]). There are many locations where custom controls can be added in the editor UI.
Please remember that you have to manage the visibility of your custom controls yourself (and likely hide it after adding it).
When your plugin is deactivated, make sure to remove your custom control with [method remove_control_from_container] and free it with [method Node.queue_free].
*/
func (self Instance) AddControlToContainer(container classdb.EditorPluginCustomControlContainer, control objects.Control) {
	class(self).AddControlToContainer(container, control)
}

/*
Adds a control to the bottom panel (together with Output, Debug, Animation, etc). Returns a reference to the button added. It's up to you to hide/show the button when needed. When your plugin is deactivated, make sure to remove your custom control with [method remove_control_from_bottom_panel] and free it with [method Node.queue_free].
Optionally, you can specify a shortcut parameter. When pressed, this shortcut will toggle the bottom panel's visibility. See the default editor bottom panel shortcuts in the Editor Settings for inspiration. Per convention, they all use [kbd]Alt[/kbd] modifier.
*/
func (self Instance) AddControlToBottomPanel(control objects.Control, title string) objects.Button {
	return objects.Button(class(self).AddControlToBottomPanel(control, gd.NewString(title), ([1]objects.Shortcut{}[0])))
}

/*
Adds the control to a specific dock slot (see [enum DockSlot] for options).
If the dock is repositioned and as long as the plugin is active, the editor will save the dock position on further sessions.
When your plugin is deactivated, make sure to remove your custom control with [method remove_control_from_docks] and free it with [method Node.queue_free].
Optionally, you can specify a shortcut parameter. When pressed, this shortcut will toggle the dock's visibility once it's moved to the bottom panel (this shortcut does not affect the dock otherwise). See the default editor bottom panel shortcuts in the Editor Settings for inspiration. Per convention, they all use [kbd]Alt[/kbd] modifier.
*/
func (self Instance) AddControlToDock(slot classdb.EditorPluginDockSlot, control objects.Control) {
	class(self).AddControlToDock(slot, control, ([1]objects.Shortcut{}[0]))
}

/*
Removes the control from the dock. You have to manually [method Node.queue_free] the control.
*/
func (self Instance) RemoveControlFromDocks(control objects.Control) {
	class(self).RemoveControlFromDocks(control)
}

/*
Removes the control from the bottom panel. You have to manually [method Node.queue_free] the control.
*/
func (self Instance) RemoveControlFromBottomPanel(control objects.Control) {
	class(self).RemoveControlFromBottomPanel(control)
}

/*
Removes the control from the specified container. You have to manually [method Node.queue_free] the control.
*/
func (self Instance) RemoveControlFromContainer(container classdb.EditorPluginCustomControlContainer, control objects.Control) {
	class(self).RemoveControlFromContainer(container, control)
}

/*
Sets the tab icon for the given control in a dock slot. Setting to [code]null[/code] removes the icon.
*/
func (self Instance) SetDockTabIcon(control objects.Control, icon objects.Texture2D) {
	class(self).SetDockTabIcon(control, icon)
}

/*
Adds a custom menu item to [b]Project > Tools[/b] named [param name]. When clicked, the provided [param callable] will be called.
*/
func (self Instance) AddToolMenuItem(name string, callable gd.Callable) {
	class(self).AddToolMenuItem(gd.NewString(name), callable)
}

/*
Adds a custom [PopupMenu] submenu under [b]Project > Tools >[/b] [param name]. Use [method remove_tool_menu_item] on plugin clean up to remove the menu.
*/
func (self Instance) AddToolSubmenuItem(name string, submenu objects.PopupMenu) {
	class(self).AddToolSubmenuItem(gd.NewString(name), submenu)
}

/*
Removes a menu [param name] from [b]Project > Tools[/b].
*/
func (self Instance) RemoveToolMenuItem(name string) {
	class(self).RemoveToolMenuItem(gd.NewString(name))
}

/*
Returns the [PopupMenu] under [b]Scene > Export As...[/b].
*/
func (self Instance) GetExportAsMenu() objects.PopupMenu {
	return objects.PopupMenu(class(self).GetExportAsMenu())
}

/*
Adds a custom type, which will appear in the list of nodes or resources.
When a given node or resource is selected, the base type will be instantiated (e.g. "Node3D", "Control", "Resource"), then the script will be loaded and set to this object.
[b]Note:[/b] The base type is the base engine class which this type's class hierarchy inherits, not any custom type parent classes.
You can use the virtual method [method _handles] to check if your custom object is being edited by checking the script or using the [code]is[/code] keyword.
During run-time, this will be a simple object with a script so this function does not need to be called then.
[b]Note:[/b] Custom types added this way are not true classes. They are just a helper to create a node with specific script.
*/
func (self Instance) AddCustomType(atype string, base string, script objects.Script, icon objects.Texture2D) {
	class(self).AddCustomType(gd.NewString(atype), gd.NewString(base), script, icon)
}

/*
Removes a custom type added by [method add_custom_type].
*/
func (self Instance) RemoveCustomType(atype string) {
	class(self).RemoveCustomType(gd.NewString(atype))
}

/*
Adds a script at [param path] to the Autoload list as [param name].
*/
func (self Instance) AddAutoloadSingleton(name string, path string) {
	class(self).AddAutoloadSingleton(gd.NewString(name), gd.NewString(path))
}

/*
Removes an Autoload [param name] from the list.
*/
func (self Instance) RemoveAutoloadSingleton(name string) {
	class(self).RemoveAutoloadSingleton(gd.NewString(name))
}

/*
Updates the overlays of the 2D and 3D editor viewport. Causes methods [method _forward_canvas_draw_over_viewport], [method _forward_canvas_force_draw_over_viewport], [method _forward_3d_draw_over_viewport] and [method _forward_3d_force_draw_over_viewport] to be called.
*/
func (self Instance) UpdateOverlays() int {
	return int(int(class(self).UpdateOverlays()))
}

/*
Makes a specific item in the bottom panel visible.
*/
func (self Instance) MakeBottomPanelItemVisible(item objects.Control) {
	class(self).MakeBottomPanelItemVisible(item)
}

/*
Minimizes the bottom panel.
*/
func (self Instance) HideBottomPanel() {
	class(self).HideBottomPanel()
}

/*
Gets the undo/redo object. Most actions in the editor can be undoable, so use this object to make sure this happens when it's worth it.
*/
func (self Instance) GetUndoRedo() objects.EditorUndoRedoManager {
	return objects.EditorUndoRedoManager(class(self).GetUndoRedo())
}

/*
Hooks a callback into the undo/redo action creation when a property is modified in the inspector. This allows, for example, to save other properties that may be lost when a given property is modified.
The callback should have 4 arguments: [Object] [code]undo_redo[/code], [Object] [code]modified_object[/code], [String] [code]property[/code] and [Variant] [code]new_value[/code]. They are, respectively, the [UndoRedo] object used by the inspector, the currently modified object, the name of the modified property and the new value the property is about to take.
*/
func (self Instance) AddUndoRedoInspectorHookCallback(callable gd.Callable) {
	class(self).AddUndoRedoInspectorHookCallback(callable)
}

/*
Removes a callback previously added by [method add_undo_redo_inspector_hook_callback].
*/
func (self Instance) RemoveUndoRedoInspectorHookCallback(callable gd.Callable) {
	class(self).RemoveUndoRedoInspectorHookCallback(callable)
}

/*
Queue save the project's editor layout.
*/
func (self Instance) QueueSaveLayout() {
	class(self).QueueSaveLayout()
}

/*
Registers a custom translation parser plugin for extracting translatable strings from custom files.
*/
func (self Instance) AddTranslationParserPlugin(parser objects.EditorTranslationParserPlugin) {
	class(self).AddTranslationParserPlugin(parser)
}

/*
Removes a custom translation parser plugin registered by [method add_translation_parser_plugin].
*/
func (self Instance) RemoveTranslationParserPlugin(parser objects.EditorTranslationParserPlugin) {
	class(self).RemoveTranslationParserPlugin(parser)
}

/*
Registers a new [EditorImportPlugin]. Import plugins are used to import custom and unsupported assets as a custom [Resource] type.
If [param first_priority] is [code]true[/code], the new import plugin is inserted first in the list and takes precedence over pre-existing plugins.
[b]Note:[/b] If you want to import custom 3D asset formats use [method add_scene_format_importer_plugin] instead.
See [method add_inspector_plugin] for an example of how to register a plugin.
*/
func (self Instance) AddImportPlugin(importer objects.EditorImportPlugin) {
	class(self).AddImportPlugin(importer, false)
}

/*
Removes an import plugin registered by [method add_import_plugin].
*/
func (self Instance) RemoveImportPlugin(importer objects.EditorImportPlugin) {
	class(self).RemoveImportPlugin(importer)
}

/*
Registers a new [EditorSceneFormatImporter]. Scene importers are used to import custom 3D asset formats as scenes.
If [param first_priority] is [code]true[/code], the new import plugin is inserted first in the list and takes precedence over pre-existing plugins.
*/
func (self Instance) AddSceneFormatImporterPlugin(scene_format_importer objects.EditorSceneFormatImporter) {
	class(self).AddSceneFormatImporterPlugin(scene_format_importer, false)
}

/*
Removes a scene format importer registered by [method add_scene_format_importer_plugin].
*/
func (self Instance) RemoveSceneFormatImporterPlugin(scene_format_importer objects.EditorSceneFormatImporter) {
	class(self).RemoveSceneFormatImporterPlugin(scene_format_importer)
}

/*
Add a [EditorScenePostImportPlugin]. These plugins allow customizing the import process of 3D assets by adding new options to the import dialogs.
If [param first_priority] is [code]true[/code], the new import plugin is inserted first in the list and takes precedence over pre-existing plugins.
*/
func (self Instance) AddScenePostImportPlugin(scene_import_plugin objects.EditorScenePostImportPlugin) {
	class(self).AddScenePostImportPlugin(scene_import_plugin, false)
}

/*
Remove the [EditorScenePostImportPlugin], added with [method add_scene_post_import_plugin].
*/
func (self Instance) RemoveScenePostImportPlugin(scene_import_plugin objects.EditorScenePostImportPlugin) {
	class(self).RemoveScenePostImportPlugin(scene_import_plugin)
}

/*
Registers a new [EditorExportPlugin]. Export plugins are used to perform tasks when the project is being exported.
See [method add_inspector_plugin] for an example of how to register a plugin.
*/
func (self Instance) AddExportPlugin(plugin objects.EditorExportPlugin) {
	class(self).AddExportPlugin(plugin)
}

/*
Removes an export plugin registered by [method add_export_plugin].
*/
func (self Instance) RemoveExportPlugin(plugin objects.EditorExportPlugin) {
	class(self).RemoveExportPlugin(plugin)
}

/*
Registers a new [EditorNode3DGizmoPlugin]. Gizmo plugins are used to add custom gizmos to the 3D preview viewport for a [Node3D].
See [method add_inspector_plugin] for an example of how to register a plugin.
*/
func (self Instance) AddNode3dGizmoPlugin(plugin objects.EditorNode3DGizmoPlugin) {
	class(self).AddNode3dGizmoPlugin(plugin)
}

/*
Removes a gizmo plugin registered by [method add_node_3d_gizmo_plugin].
*/
func (self Instance) RemoveNode3dGizmoPlugin(plugin objects.EditorNode3DGizmoPlugin) {
	class(self).RemoveNode3dGizmoPlugin(plugin)
}

/*
Registers a new [EditorInspectorPlugin]. Inspector plugins are used to extend [EditorInspector] and provide custom configuration tools for your object's properties.
[b]Note:[/b] Always use [method remove_inspector_plugin] to remove the registered [EditorInspectorPlugin] when your [EditorPlugin] is disabled to prevent leaks and an unexpected behavior.
[codeblocks]
[gdscript]
const MyInspectorPlugin = preload("res://addons/your_addon/path/to/your/script.gd")
var inspector_plugin = MyInspectorPlugin.new()

func _enter_tree():

	add_inspector_plugin(inspector_plugin)

func _exit_tree():

	remove_inspector_plugin(inspector_plugin)

[/gdscript]
[/codeblocks]
*/
func (self Instance) AddInspectorPlugin(plugin objects.EditorInspectorPlugin) {
	class(self).AddInspectorPlugin(plugin)
}

/*
Removes an inspector plugin registered by [method add_import_plugin]
*/
func (self Instance) RemoveInspectorPlugin(plugin objects.EditorInspectorPlugin) {
	class(self).RemoveInspectorPlugin(plugin)
}

/*
Registers a new [EditorResourceConversionPlugin]. Resource conversion plugins are used to add custom resource converters to the editor inspector.
See [EditorResourceConversionPlugin] for an example of how to create a resource conversion plugin.
*/
func (self Instance) AddResourceConversionPlugin(plugin objects.EditorResourceConversionPlugin) {
	class(self).AddResourceConversionPlugin(plugin)
}

/*
Removes a resource conversion plugin registered by [method add_resource_conversion_plugin].
*/
func (self Instance) RemoveResourceConversionPlugin(plugin objects.EditorResourceConversionPlugin) {
	class(self).RemoveResourceConversionPlugin(plugin)
}

/*
Use this method if you always want to receive inputs from 3D view screen inside [method _forward_3d_gui_input]. It might be especially usable if your plugin will want to use raycast in the scene.
*/
func (self Instance) SetInputEventForwardingAlwaysEnabled() {
	class(self).SetInputEventForwardingAlwaysEnabled()
}

/*
Enables calling of [method _forward_canvas_force_draw_over_viewport] for the 2D editor and [method _forward_3d_force_draw_over_viewport] for the 3D editor when their viewports are updated. You need to call this method only once and it will work permanently for this plugin.
*/
func (self Instance) SetForceDrawOverForwardingEnabled() {
	class(self).SetForceDrawOverForwardingEnabled()
}

/*
Returns the [EditorInterface] singleton instance.
*/
func (self Instance) GetEditorInterface() objects.EditorInterface {
	return objects.EditorInterface(class(self).GetEditorInterface())
}

/*
Gets the Editor's dialog used for making scripts.
[b]Note:[/b] Users can configure it before use.
[b]Warning:[/b] Removing and freeing this node will render a part of the editor useless and may cause a crash.
*/
func (self Instance) GetScriptCreateDialog() objects.ScriptCreateDialog {
	return objects.ScriptCreateDialog(class(self).GetScriptCreateDialog())
}

/*
Adds a [Script] as debugger plugin to the Debugger. The script must extend [EditorDebuggerPlugin].
*/
func (self Instance) AddDebuggerPlugin(script objects.EditorDebuggerPlugin) {
	class(self).AddDebuggerPlugin(script)
}

/*
Removes the debugger plugin with given script from the Debugger.
*/
func (self Instance) RemoveDebuggerPlugin(script objects.EditorDebuggerPlugin) {
	class(self).RemoveDebuggerPlugin(script)
}

/*
Provide the version of the plugin declared in the [code]plugin.cfg[/code] config file.
*/
func (self Instance) GetPluginVersion() string {
	return string(class(self).GetPluginVersion().String())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.EditorPlugin

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorPlugin"))
	return Instance{classdb.EditorPlugin(object)}
}

/*
Called when there is a root node in the current edited scene, [method _handles] is implemented and an [InputEvent] happens in the 2D viewport. Intercepts the [InputEvent], if [code]return true[/code] [EditorPlugin] consumes the [param event], otherwise forwards [param event] to other Editor classes.
[b]Example:[/b]
[codeblocks]
[gdscript]
# Prevents the InputEvent from reaching other Editor classes.
func _forward_canvas_gui_input(event):

	return true

[/gdscript]
[csharp]
// Prevents the InputEvent from reaching other Editor classes.
public override bool ForwardCanvasGuiInput(InputEvent @event)

	{
	    return true;
	}

[/csharp]
[/codeblocks]
Must [code]return false[/code] in order to forward the [InputEvent] to other Editor classes.
[b]Example:[/b]
[codeblocks]
[gdscript]
# Consumes InputEventMouseMotion and forwards other InputEvent types.
func _forward_canvas_gui_input(event):

	if (event is InputEventMouseMotion):
	    return true
	return false

[/gdscript]
[csharp]
// Consumes InputEventMouseMotion and forwards other InputEvent types.
public override bool _ForwardCanvasGuiInput(InputEvent @event)

	{
	    if (@event is InputEventMouseMotion)
	    {
	        return true;
	    }
	    return false;
	}

[/csharp]
[/codeblocks]
*/
func (class) _forward_canvas_gui_input(impl func(ptr unsafe.Pointer, event objects.InputEvent) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var event = objects.InputEvent{pointers.New[classdb.InputEvent]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(event[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, event)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called by the engine when the 2D editor's viewport is updated. Use the [code]overlay[/code] [Control] for drawing. You can update the viewport manually by calling [method update_overlays].
[codeblocks]
[gdscript]
func _forward_canvas_draw_over_viewport(overlay):

	# Draw a circle at cursor position.
	overlay.draw_circle(overlay.get_local_mouse_position(), 64, Color.WHITE)

func _forward_canvas_gui_input(event):

	if event is InputEventMouseMotion:
	    # Redraw viewport when cursor is moved.
	    update_overlays()
	    return true
	return false

[/gdscript]
[csharp]
public override void _ForwardCanvasDrawOverViewport(Control viewportControl)

	{
	    // Draw a circle at cursor position.
	    viewportControl.DrawCircle(viewportControl.GetLocalMousePosition(), 64, Colors.White);
	}

public override bool _ForwardCanvasGuiInput(InputEvent @event)

	{
	    if (@event is InputEventMouseMotion)
	    {
	        // Redraw viewport when cursor is moved.
	        UpdateOverlays();
	        return true;
	    }
	    return false;
	}

[/csharp]
[/codeblocks]
*/
func (class) _forward_canvas_draw_over_viewport(impl func(ptr unsafe.Pointer, viewport_control objects.Control)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var viewport_control = objects.Control{pointers.New[classdb.Control]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(viewport_control[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, viewport_control)
	}
}

/*
This method is the same as [method _forward_canvas_draw_over_viewport], except it draws on top of everything. Useful when you need an extra layer that shows over anything else.
You need to enable calling of this method by using [method set_force_draw_over_forwarding_enabled].
*/
func (class) _forward_canvas_force_draw_over_viewport(impl func(ptr unsafe.Pointer, viewport_control objects.Control)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var viewport_control = objects.Control{pointers.New[classdb.Control]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(viewport_control[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, viewport_control)
	}
}

/*
Called when there is a root node in the current edited scene, [method _handles] is implemented, and an [InputEvent] happens in the 3D viewport. The return value decides whether the [InputEvent] is consumed or forwarded to other [EditorPlugin]s. See [enum AfterGUIInput] for options.
[b]Example:[/b]
[codeblocks]
[gdscript]
# Prevents the InputEvent from reaching other Editor classes.
func _forward_3d_gui_input(camera, event):

	return EditorPlugin.AFTER_GUI_INPUT_STOP

[/gdscript]
[csharp]
// Prevents the InputEvent from reaching other Editor classes.
public override EditorPlugin.AfterGuiInput _Forward3DGuiInput(Camera3D camera, InputEvent @event)

	{
	    return EditorPlugin.AfterGuiInput.Stop;
	}

[/csharp]
[/codeblocks]
Must [code]return EditorPlugin.AFTER_GUI_INPUT_PASS[/code] in order to forward the [InputEvent] to other Editor classes.
[b]Example:[/b]
[codeblocks]
[gdscript]
# Consumes InputEventMouseMotion and forwards other InputEvent types.
func _forward_3d_gui_input(camera, event):

	return EditorPlugin.AFTER_GUI_INPUT_STOP if event is InputEventMouseMotion else EditorPlugin.AFTER_GUI_INPUT_PASS

[/gdscript]
[csharp]
// Consumes InputEventMouseMotion and forwards other InputEvent types.
public override EditorPlugin.AfterGuiInput _Forward3DGuiInput(Camera3D camera, InputEvent @event)

	{
	    return @event is InputEventMouseMotion ? EditorPlugin.AfterGuiInput.Stop : EditorPlugin.AfterGuiInput.Pass;
	}

[/csharp]
[/codeblocks]
*/
func (class) _forward_3d_gui_input(impl func(ptr unsafe.Pointer, viewport_camera objects.Camera3D, event objects.InputEvent) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var viewport_camera = objects.Camera3D{pointers.New[classdb.Camera3D]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(viewport_camera[0])
		var event = objects.InputEvent{pointers.New[classdb.InputEvent]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 1)})}
		defer pointers.End(event[0])
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, viewport_camera, event)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called by the engine when the 3D editor's viewport is updated. Use the [code]overlay[/code] [Control] for drawing. You can update the viewport manually by calling [method update_overlays].
[codeblocks]
[gdscript]
func _forward_3d_draw_over_viewport(overlay):

	# Draw a circle at cursor position.
	overlay.draw_circle(overlay.get_local_mouse_position(), 64, Color.WHITE)

func _forward_3d_gui_input(camera, event):

	if event is InputEventMouseMotion:
	    # Redraw viewport when cursor is moved.
	    update_overlays()
	    return EditorPlugin.AFTER_GUI_INPUT_STOP
	return EditorPlugin.AFTER_GUI_INPUT_PASS

[/gdscript]
[csharp]
public override void _Forward3DDrawOverViewport(Control viewportControl)

	{
	    // Draw a circle at cursor position.
	    viewportControl.DrawCircle(viewportControl.GetLocalMousePosition(), 64, Colors.White);
	}

public override EditorPlugin.AfterGuiInput _Forward3DGuiInput(Camera3D viewportCamera, InputEvent @event)

	{
	    if (@event is InputEventMouseMotion)
	    {
	        // Redraw viewport when cursor is moved.
	        UpdateOverlays();
	        return EditorPlugin.AfterGuiInput.Stop;
	    }
	    return EditorPlugin.AfterGuiInput.Pass;
	}

[/csharp]
[/codeblocks]
*/
func (class) _forward_3d_draw_over_viewport(impl func(ptr unsafe.Pointer, viewport_control objects.Control)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var viewport_control = objects.Control{pointers.New[classdb.Control]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(viewport_control[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, viewport_control)
	}
}

/*
This method is the same as [method _forward_3d_draw_over_viewport], except it draws on top of everything. Useful when you need an extra layer that shows over anything else.
You need to enable calling of this method by using [method set_force_draw_over_forwarding_enabled].
*/
func (class) _forward_3d_force_draw_over_viewport(impl func(ptr unsafe.Pointer, viewport_control objects.Control)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var viewport_control = objects.Control{pointers.New[classdb.Control]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(viewport_control[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, viewport_control)
	}
}

/*
Override this method in your plugin to provide the name of the plugin when displayed in the Godot editor.
For main screen plugins, this appears at the top of the screen, to the right of the "2D", "3D", "Script", and "AssetLib" buttons.
*/
func (class) _get_plugin_name(impl func(ptr unsafe.Pointer) gd.String) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Override this method in your plugin to return a [Texture2D] in order to give it an icon.
For main screen plugins, this appears at the top of the screen, to the right of the "2D", "3D", "Script", and "AssetLib" buttons.
Ideally, the plugin icon should be white with a transparent background and 16×16 pixels in size.
[codeblocks]
[gdscript]
func _get_plugin_icon():

	# You can use a custom icon:
	return preload("res://addons/my_plugin/my_plugin_icon.svg")
	# Or use a built-in icon:
	return EditorInterface.get_editor_theme().get_icon("Node", "EditorIcons")

[/gdscript]
[csharp]
public override Texture2D _GetPluginIcon()

	{
	    // You can use a custom icon:
	    return ResourceLoader.Load<Texture2D>("res://addons/my_plugin/my_plugin_icon.svg");
	    // Or use a built-in icon:
	    return EditorInterface.Singleton.GetEditorTheme().GetIcon("Node", "EditorIcons");
	}

[/csharp]
[/codeblocks]
*/
func (class) _get_plugin_icon(impl func(ptr unsafe.Pointer) objects.Texture2D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret[0])
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Returns [code]true[/code] if this is a main screen editor plugin (it goes in the workspace selector together with [b]2D[/b], [b]3D[/b], [b]Script[/b] and [b]AssetLib[/b]).
When the plugin's workspace is selected, other main screen plugins will be hidden, but your plugin will not appear automatically. It needs to be added as a child of [method EditorInterface.get_editor_main_screen] and made visible inside [method _make_visible].
Use [method _get_plugin_name] and [method _get_plugin_icon] to customize the plugin button's appearance.
[codeblock]
var plugin_control

func _enter_tree():

	plugin_control = preload("my_plugin_control.tscn").instantiate()
	EditorInterface.get_editor_main_screen().add_child(plugin_control)
	plugin_control.hide()

func _has_main_screen():

	return true

func _make_visible(visible):

	plugin_control.visible = visible

func _get_plugin_name():

	return "My Super Cool Plugin 3000"

func _get_plugin_icon():

	return EditorInterface.get_editor_theme().get_icon("Node", "EditorIcons")

[/codeblock]
*/
func (class) _has_main_screen(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
This function will be called when the editor is requested to become visible. It is used for plugins that edit a specific object type.
Remember that you have to manage the visibility of all your editor controls manually.
*/
func (class) _make_visible(impl func(ptr unsafe.Pointer, visible bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var visible = gd.UnsafeGet[bool](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, visible)
	}
}

/*
This function is used for plugins that edit specific object types (nodes or resources). It requests the editor to edit the given object.
[param object] can be [code]null[/code] if the plugin was editing an object, but there is no longer any selected object handled by this plugin. It can be used to cleanup editing state.
*/
func (class) _edit(impl func(ptr unsafe.Pointer, obj gd.Object)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})
		defer pointers.End(obj)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, obj)
	}
}

/*
Implement this function if your plugin edits a specific type of object (Resource or Node). If you return [code]true[/code], then you will get the functions [method _edit] and [method _make_visible] called when the editor requests them. If you have declared the methods [method _forward_canvas_gui_input] and [method _forward_3d_gui_input] these will be called too.
[b]Note:[/b] Each plugin should handle only one type of objects at a time. If a plugin handles more types of objects and they are edited at the same time, it will result in errors.
*/
func (class) _handles(impl func(ptr unsafe.Pointer, obj gd.Object) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var obj = pointers.New[gd.Object]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})
		defer pointers.End(obj)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, obj)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Override this method to provide a state data you want to be saved, like view position, grid settings, folding, etc. This is used when saving the scene (so state is kept when opening it again) and for switching tabs (so state can be restored when the tab returns). This data is automatically saved for each scene in an [code]editstate[/code] file in the editor metadata folder. If you want to store global (scene-independent) editor data for your plugin, you can use [method _get_window_layout] instead.
Use [method _set_state] to restore your saved state.
[b]Note:[/b] This method should not be used to save important settings that should persist with the project.
[b]Note:[/b] You must implement [method _get_plugin_name] for the state to be stored and restored correctly.
[codeblock]
func _get_state():

	var state = {"zoom": zoom, "preferred_color": my_color}
	return state

[/codeblock]
*/
func (class) _get_state(impl func(ptr unsafe.Pointer) gd.Dictionary) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Restore the state saved by [method _get_state]. This method is called when the current scene tab is changed in the editor.
[b]Note:[/b] Your plugin must implement [method _get_plugin_name], otherwise it will not be recognized and this method will not be called.
[codeblock]
func _set_state(data):

	zoom = data.get("zoom", 1.0)
	preferred_color = data.get("my_color", Color.WHITE)

[/codeblock]
*/
func (class) _set_state(impl func(ptr unsafe.Pointer, state gd.Dictionary)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var state = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, state)
	}
}

/*
Clear all the state and reset the object being edited to zero. This ensures your plugin does not keep editing a currently existing node, or a node from the wrong scene.
*/
func (class) _clear(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Override this method to provide a custom message that lists unsaved changes. The editor will call this method when exiting or when closing a scene, and display the returned string in a confirmation dialog. Return empty string if the plugin has no unsaved changes.
When closing a scene, [param for_scene] is the path to the scene being closed. You can use it to handle built-in resources in that scene.
If the user confirms saving, [method _save_external_data] will be called, before closing the editor.
[codeblock]
func _get_unsaved_status(for_scene):

	if not unsaved:
	    return ""

	if for_scene.is_empty():
	    return "Save changes in MyCustomPlugin before closing?"
	else:
	    return "Scene %s has changes from MyCustomPlugin. Save before closing?" % for_scene.get_file()

func _save_external_data():

	unsaved = false

[/codeblock]
If the plugin has no scene-specific changes, you can ignore the calls when closing scenes:
[codeblock]
func _get_unsaved_status(for_scene):

	if not for_scene.is_empty():
	    return ""

[/codeblock]
*/
func (class) _get_unsaved_status(impl func(ptr unsafe.Pointer, for_scene gd.String) gd.String) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var for_scene = pointers.New[gd.String](gd.UnsafeGet[[1]uintptr](p_args, 0))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, for_scene)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
This method is called after the editor saves the project or when it's closed. It asks the plugin to save edited external scenes/resources.
*/
func (class) _save_external_data(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
This method is called when the editor is about to save the project, switch to another tab, etc. It asks the plugin to apply any pending state changes to ensure consistency.
This is used, for example, in shader editors to let the plugin know that it must apply the shader code being written by the user to the object.
*/
func (class) _apply_changes(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
This is for editors that edit script-based objects. You can return a list of breakpoints in the format ([code]script:line[/code]), for example: [code]res://path_to_script.gd:25[/code].
*/
func (class) _get_breakpoints(impl func(ptr unsafe.Pointer) gd.PackedStringArray) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(ret)
		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Restore the plugin GUI layout and data saved by [method _get_window_layout]. This method is called for every plugin on editor startup. Use the provided [param configuration] file to read your saved data.
[codeblock]
func _set_window_layout(configuration):

	$Window.position = configuration.get_value("MyPlugin", "window_position", Vector2())
	$Icon.modulate = configuration.get_value("MyPlugin", "icon_color", Color.WHITE)

[/codeblock]
*/
func (class) _set_window_layout(impl func(ptr unsafe.Pointer, configuration objects.ConfigFile)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var configuration = objects.ConfigFile{pointers.New[classdb.ConfigFile]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(configuration[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, configuration)
	}
}

/*
Override this method to provide the GUI layout of the plugin or any other data you want to be stored. This is used to save the project's editor layout when [method queue_save_layout] is called or the editor layout was changed (for example changing the position of a dock). The data is stored in the [code]editor_layout.cfg[/code] file in the editor metadata directory.
Use [method _set_window_layout] to restore your saved layout.
[codeblock]
func _get_window_layout(configuration):

	configuration.set_value("MyPlugin", "window_position", $Window.position)
	configuration.set_value("MyPlugin", "icon_color", $Icon.modulate)

[/codeblock]
*/
func (class) _get_window_layout(impl func(ptr unsafe.Pointer, configuration objects.ConfigFile)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var configuration = objects.ConfigFile{pointers.New[classdb.ConfigFile]([3]uintptr{gd.UnsafeGet[uintptr](p_args, 0)})}
		defer pointers.End(configuration[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, configuration)
	}
}

/*
This method is called when the editor is about to run the project. The plugin can then perform required operations before the project runs.
This method must return a boolean. If this method returns [code]false[/code], the project will not run. The run is aborted immediately, so this also prevents all other plugins' [method _build] methods from running.
*/
func (class) _build(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Called by the engine when the user enables the [EditorPlugin] in the Plugin tab of the project settings window.
*/
func (class) _enable_plugin(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called by the engine when the user disables the [EditorPlugin] in the Plugin tab of the project settings window.
*/
func (class) _disable_plugin(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Adds a custom control to a container (see [enum CustomControlContainer]). There are many locations where custom controls can be added in the editor UI.
Please remember that you have to manage the visibility of your custom controls yourself (and likely hide it after adding it).
When your plugin is deactivated, make sure to remove your custom control with [method remove_control_from_container] and free it with [method Node.queue_free].
*/
//go:nosplit
func (self class) AddControlToContainer(container classdb.EditorPluginCustomControlContainer, control objects.Control) {
	var frame = callframe.New()
	callframe.Arg(frame, container)
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(gd.Object(control[0])))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_control_to_container, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a control to the bottom panel (together with Output, Debug, Animation, etc). Returns a reference to the button added. It's up to you to hide/show the button when needed. When your plugin is deactivated, make sure to remove your custom control with [method remove_control_from_bottom_panel] and free it with [method Node.queue_free].
Optionally, you can specify a shortcut parameter. When pressed, this shortcut will toggle the bottom panel's visibility. See the default editor bottom panel shortcuts in the Editor Settings for inspiration. Per convention, they all use [kbd]Alt[/kbd] modifier.
*/
//go:nosplit
func (self class) AddControlToBottomPanel(control objects.Control, title gd.String, shortcut objects.Shortcut) objects.Button {
	var frame = callframe.New()
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(gd.Object(control[0])))
	callframe.Arg(frame, pointers.Get(title))
	callframe.Arg(frame, pointers.Get(shortcut[0])[0])
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_control_to_bottom_panel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.Button{classdb.Button(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Adds the control to a specific dock slot (see [enum DockSlot] for options).
If the dock is repositioned and as long as the plugin is active, the editor will save the dock position on further sessions.
When your plugin is deactivated, make sure to remove your custom control with [method remove_control_from_docks] and free it with [method Node.queue_free].
Optionally, you can specify a shortcut parameter. When pressed, this shortcut will toggle the dock's visibility once it's moved to the bottom panel (this shortcut does not affect the dock otherwise). See the default editor bottom panel shortcuts in the Editor Settings for inspiration. Per convention, they all use [kbd]Alt[/kbd] modifier.
*/
//go:nosplit
func (self class) AddControlToDock(slot classdb.EditorPluginDockSlot, control objects.Control, shortcut objects.Shortcut) {
	var frame = callframe.New()
	callframe.Arg(frame, slot)
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(gd.Object(control[0])))
	callframe.Arg(frame, pointers.Get(shortcut[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_control_to_dock, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the control from the dock. You have to manually [method Node.queue_free] the control.
*/
//go:nosplit
func (self class) RemoveControlFromDocks(control objects.Control) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(control[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_control_from_docks, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the control from the bottom panel. You have to manually [method Node.queue_free] the control.
*/
//go:nosplit
func (self class) RemoveControlFromBottomPanel(control objects.Control) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(control[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_control_from_bottom_panel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the control from the specified container. You have to manually [method Node.queue_free] the control.
*/
//go:nosplit
func (self class) RemoveControlFromContainer(container classdb.EditorPluginCustomControlContainer, control objects.Control) {
	var frame = callframe.New()
	callframe.Arg(frame, container)
	callframe.Arg(frame, pointers.Get(control[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_control_from_container, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Sets the tab icon for the given control in a dock slot. Setting to [code]null[/code] removes the icon.
*/
//go:nosplit
func (self class) SetDockTabIcon(control objects.Control, icon objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(control[0])[0])
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_set_dock_tab_icon, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a custom menu item to [b]Project > Tools[/b] named [param name]. When clicked, the provided [param callable] will be called.
*/
//go:nosplit
func (self class) AddToolMenuItem(name gd.String, callable gd.Callable) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(callable))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_tool_menu_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a custom [PopupMenu] submenu under [b]Project > Tools >[/b] [param name]. Use [method remove_tool_menu_item] on plugin clean up to remove the menu.
*/
//go:nosplit
func (self class) AddToolSubmenuItem(name gd.String, submenu objects.PopupMenu) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(gd.Object(submenu[0])))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_tool_submenu_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes a menu [param name] from [b]Project > Tools[/b].
*/
//go:nosplit
func (self class) RemoveToolMenuItem(name gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_tool_menu_item, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [PopupMenu] under [b]Scene > Export As...[/b].
*/
//go:nosplit
func (self class) GetExportAsMenu() objects.PopupMenu {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_get_export_as_menu, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.PopupMenu{classdb.PopupMenu(gd.PointerLifetimeBoundTo(self.AsObject(), r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Adds a custom type, which will appear in the list of nodes or resources.
When a given node or resource is selected, the base type will be instantiated (e.g. "Node3D", "Control", "Resource"), then the script will be loaded and set to this object.
[b]Note:[/b] The base type is the base engine class which this type's class hierarchy inherits, not any custom type parent classes.
You can use the virtual method [method _handles] to check if your custom object is being edited by checking the script or using the [code]is[/code] keyword.
During run-time, this will be a simple object with a script so this function does not need to be called then.
[b]Note:[/b] Custom types added this way are not true classes. They are just a helper to create a node with specific script.
*/
//go:nosplit
func (self class) AddCustomType(atype gd.String, base gd.String, script objects.Script, icon objects.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(atype))
	callframe.Arg(frame, pointers.Get(base))
	callframe.Arg(frame, pointers.Get(script[0])[0])
	callframe.Arg(frame, pointers.Get(icon[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_custom_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes a custom type added by [method add_custom_type].
*/
//go:nosplit
func (self class) RemoveCustomType(atype gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(atype))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_custom_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Adds a script at [param path] to the Autoload list as [param name].
*/
//go:nosplit
func (self class) AddAutoloadSingleton(name gd.String, path gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	callframe.Arg(frame, pointers.Get(path))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_autoload_singleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes an Autoload [param name] from the list.
*/
//go:nosplit
func (self class) RemoveAutoloadSingleton(name gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(name))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_autoload_singleton, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Updates the overlays of the 2D and 3D editor viewport. Causes methods [method _forward_canvas_draw_over_viewport], [method _forward_canvas_force_draw_over_viewport], [method _forward_3d_draw_over_viewport] and [method _forward_3d_force_draw_over_viewport] to be called.
*/
//go:nosplit
func (self class) UpdateOverlays() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_update_overlays, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Makes a specific item in the bottom panel visible.
*/
//go:nosplit
func (self class) MakeBottomPanelItemVisible(item objects.Control) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(item[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_make_bottom_panel_item_visible, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Minimizes the bottom panel.
*/
//go:nosplit
func (self class) HideBottomPanel() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_hide_bottom_panel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Gets the undo/redo object. Most actions in the editor can be undoable, so use this object to make sure this happens when it's worth it.
*/
//go:nosplit
func (self class) GetUndoRedo() objects.EditorUndoRedoManager {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_get_undo_redo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.EditorUndoRedoManager{classdb.EditorUndoRedoManager(gd.PointerLifetimeBoundTo(self.AsObject(), r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Hooks a callback into the undo/redo action creation when a property is modified in the inspector. This allows, for example, to save other properties that may be lost when a given property is modified.
The callback should have 4 arguments: [Object] [code]undo_redo[/code], [Object] [code]modified_object[/code], [String] [code]property[/code] and [Variant] [code]new_value[/code]. They are, respectively, the [UndoRedo] object used by the inspector, the currently modified object, the name of the modified property and the new value the property is about to take.
*/
//go:nosplit
func (self class) AddUndoRedoInspectorHookCallback(callable gd.Callable) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(callable))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_undo_redo_inspector_hook_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes a callback previously added by [method add_undo_redo_inspector_hook_callback].
*/
//go:nosplit
func (self class) RemoveUndoRedoInspectorHookCallback(callable gd.Callable) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(callable))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_undo_redo_inspector_hook_callback, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Queue save the project's editor layout.
*/
//go:nosplit
func (self class) QueueSaveLayout() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_queue_save_layout, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Registers a custom translation parser plugin for extracting translatable strings from custom files.
*/
//go:nosplit
func (self class) AddTranslationParserPlugin(parser objects.EditorTranslationParserPlugin) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(parser[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_translation_parser_plugin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes a custom translation parser plugin registered by [method add_translation_parser_plugin].
*/
//go:nosplit
func (self class) RemoveTranslationParserPlugin(parser objects.EditorTranslationParserPlugin) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(parser[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_translation_parser_plugin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Registers a new [EditorImportPlugin]. Import plugins are used to import custom and unsupported assets as a custom [Resource] type.
If [param first_priority] is [code]true[/code], the new import plugin is inserted first in the list and takes precedence over pre-existing plugins.
[b]Note:[/b] If you want to import custom 3D asset formats use [method add_scene_format_importer_plugin] instead.
See [method add_inspector_plugin] for an example of how to register a plugin.
*/
//go:nosplit
func (self class) AddImportPlugin(importer objects.EditorImportPlugin, first_priority bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(importer[0])[0])
	callframe.Arg(frame, first_priority)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_import_plugin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes an import plugin registered by [method add_import_plugin].
*/
//go:nosplit
func (self class) RemoveImportPlugin(importer objects.EditorImportPlugin) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(importer[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_import_plugin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Registers a new [EditorSceneFormatImporter]. Scene importers are used to import custom 3D asset formats as scenes.
If [param first_priority] is [code]true[/code], the new import plugin is inserted first in the list and takes precedence over pre-existing plugins.
*/
//go:nosplit
func (self class) AddSceneFormatImporterPlugin(scene_format_importer objects.EditorSceneFormatImporter, first_priority bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(scene_format_importer[0])[0])
	callframe.Arg(frame, first_priority)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_scene_format_importer_plugin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes a scene format importer registered by [method add_scene_format_importer_plugin].
*/
//go:nosplit
func (self class) RemoveSceneFormatImporterPlugin(scene_format_importer objects.EditorSceneFormatImporter) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(scene_format_importer[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_scene_format_importer_plugin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Add a [EditorScenePostImportPlugin]. These plugins allow customizing the import process of 3D assets by adding new options to the import dialogs.
If [param first_priority] is [code]true[/code], the new import plugin is inserted first in the list and takes precedence over pre-existing plugins.
*/
//go:nosplit
func (self class) AddScenePostImportPlugin(scene_import_plugin objects.EditorScenePostImportPlugin, first_priority bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(scene_import_plugin[0])[0])
	callframe.Arg(frame, first_priority)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_scene_post_import_plugin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Remove the [EditorScenePostImportPlugin], added with [method add_scene_post_import_plugin].
*/
//go:nosplit
func (self class) RemoveScenePostImportPlugin(scene_import_plugin objects.EditorScenePostImportPlugin) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(scene_import_plugin[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_scene_post_import_plugin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Registers a new [EditorExportPlugin]. Export plugins are used to perform tasks when the project is being exported.
See [method add_inspector_plugin] for an example of how to register a plugin.
*/
//go:nosplit
func (self class) AddExportPlugin(plugin objects.EditorExportPlugin) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(plugin[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_export_plugin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes an export plugin registered by [method add_export_plugin].
*/
//go:nosplit
func (self class) RemoveExportPlugin(plugin objects.EditorExportPlugin) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(plugin[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_export_plugin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Registers a new [EditorNode3DGizmoPlugin]. Gizmo plugins are used to add custom gizmos to the 3D preview viewport for a [Node3D].
See [method add_inspector_plugin] for an example of how to register a plugin.
*/
//go:nosplit
func (self class) AddNode3dGizmoPlugin(plugin objects.EditorNode3DGizmoPlugin) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(plugin[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_node_3d_gizmo_plugin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes a gizmo plugin registered by [method add_node_3d_gizmo_plugin].
*/
//go:nosplit
func (self class) RemoveNode3dGizmoPlugin(plugin objects.EditorNode3DGizmoPlugin) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(plugin[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_node_3d_gizmo_plugin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Registers a new [EditorInspectorPlugin]. Inspector plugins are used to extend [EditorInspector] and provide custom configuration tools for your object's properties.
[b]Note:[/b] Always use [method remove_inspector_plugin] to remove the registered [EditorInspectorPlugin] when your [EditorPlugin] is disabled to prevent leaks and an unexpected behavior.
[codeblocks]
[gdscript]
const MyInspectorPlugin = preload("res://addons/your_addon/path/to/your/script.gd")
var inspector_plugin = MyInspectorPlugin.new()

func _enter_tree():
    add_inspector_plugin(inspector_plugin)

func _exit_tree():
    remove_inspector_plugin(inspector_plugin)
[/gdscript]
[/codeblocks]
*/
//go:nosplit
func (self class) AddInspectorPlugin(plugin objects.EditorInspectorPlugin) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(plugin[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_inspector_plugin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes an inspector plugin registered by [method add_import_plugin]
*/
//go:nosplit
func (self class) RemoveInspectorPlugin(plugin objects.EditorInspectorPlugin) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(plugin[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_inspector_plugin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Registers a new [EditorResourceConversionPlugin]. Resource conversion plugins are used to add custom resource converters to the editor inspector.
See [EditorResourceConversionPlugin] for an example of how to create a resource conversion plugin.
*/
//go:nosplit
func (self class) AddResourceConversionPlugin(plugin objects.EditorResourceConversionPlugin) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(plugin[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_resource_conversion_plugin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes a resource conversion plugin registered by [method add_resource_conversion_plugin].
*/
//go:nosplit
func (self class) RemoveResourceConversionPlugin(plugin objects.EditorResourceConversionPlugin) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(plugin[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_resource_conversion_plugin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Use this method if you always want to receive inputs from 3D view screen inside [method _forward_3d_gui_input]. It might be especially usable if your plugin will want to use raycast in the scene.
*/
//go:nosplit
func (self class) SetInputEventForwardingAlwaysEnabled() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_set_input_event_forwarding_always_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Enables calling of [method _forward_canvas_force_draw_over_viewport] for the 2D editor and [method _forward_3d_force_draw_over_viewport] for the 3D editor when their viewports are updated. You need to call this method only once and it will work permanently for this plugin.
*/
//go:nosplit
func (self class) SetForceDrawOverForwardingEnabled() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_set_force_draw_over_forwarding_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [EditorInterface] singleton instance.
*/
//go:nosplit
func (self class) GetEditorInterface() objects.EditorInterface {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_get_editor_interface, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.EditorInterface{classdb.EditorInterface(gd.PointerLifetimeBoundTo(self.AsObject(), r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Gets the Editor's dialog used for making scripts.
[b]Note:[/b] Users can configure it before use.
[b]Warning:[/b] Removing and freeing this node will render a part of the editor useless and may cause a crash.
*/
//go:nosplit
func (self class) GetScriptCreateDialog() objects.ScriptCreateDialog {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_get_script_create_dialog, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.ScriptCreateDialog{classdb.ScriptCreateDialog(gd.PointerLifetimeBoundTo(self.AsObject(), r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Adds a [Script] as debugger plugin to the Debugger. The script must extend [EditorDebuggerPlugin].
*/
//go:nosplit
func (self class) AddDebuggerPlugin(script objects.EditorDebuggerPlugin) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(script[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_add_debugger_plugin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Removes the debugger plugin with given script from the Debugger.
*/
//go:nosplit
func (self class) RemoveDebuggerPlugin(script objects.EditorDebuggerPlugin) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(script[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_remove_debugger_plugin, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Provide the version of the plugin declared in the [code]plugin.cfg[/code] config file.
*/
//go:nosplit
func (self class) GetPluginVersion() gd.String {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.EditorPlugin.Bind_get_plugin_version, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.String](r_ret.Get())
	frame.Free()
	return ret
}
func (self Instance) OnSceneChanged(cb func(scene_root objects.Node)) {
	self[0].AsObject().Connect(gd.NewStringName("scene_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnSceneClosed(cb func(filepath string)) {
	self[0].AsObject().Connect(gd.NewStringName("scene_closed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnMainScreenChanged(cb func(screen_name string)) {
	self[0].AsObject().Connect(gd.NewStringName("main_screen_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnResourceSaved(cb func(resource objects.Resource)) {
	self[0].AsObject().Connect(gd.NewStringName("resource_saved"), gd.NewCallable(cb), 0)
}

func (self Instance) OnSceneSaved(cb func(filepath string)) {
	self[0].AsObject().Connect(gd.NewStringName("scene_saved"), gd.NewCallable(cb), 0)
}

func (self Instance) OnProjectSettingsChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("project_settings_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsEditorPlugin() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsEditorPlugin() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced       { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance    { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_forward_canvas_gui_input":
		return reflect.ValueOf(self._forward_canvas_gui_input)
	case "_forward_canvas_draw_over_viewport":
		return reflect.ValueOf(self._forward_canvas_draw_over_viewport)
	case "_forward_canvas_force_draw_over_viewport":
		return reflect.ValueOf(self._forward_canvas_force_draw_over_viewport)
	case "_forward_3d_gui_input":
		return reflect.ValueOf(self._forward_3d_gui_input)
	case "_forward_3d_draw_over_viewport":
		return reflect.ValueOf(self._forward_3d_draw_over_viewport)
	case "_forward_3d_force_draw_over_viewport":
		return reflect.ValueOf(self._forward_3d_force_draw_over_viewport)
	case "_get_plugin_name":
		return reflect.ValueOf(self._get_plugin_name)
	case "_get_plugin_icon":
		return reflect.ValueOf(self._get_plugin_icon)
	case "_has_main_screen":
		return reflect.ValueOf(self._has_main_screen)
	case "_make_visible":
		return reflect.ValueOf(self._make_visible)
	case "_edit":
		return reflect.ValueOf(self._edit)
	case "_handles":
		return reflect.ValueOf(self._handles)
	case "_get_state":
		return reflect.ValueOf(self._get_state)
	case "_set_state":
		return reflect.ValueOf(self._set_state)
	case "_clear":
		return reflect.ValueOf(self._clear)
	case "_get_unsaved_status":
		return reflect.ValueOf(self._get_unsaved_status)
	case "_save_external_data":
		return reflect.ValueOf(self._save_external_data)
	case "_apply_changes":
		return reflect.ValueOf(self._apply_changes)
	case "_get_breakpoints":
		return reflect.ValueOf(self._get_breakpoints)
	case "_set_window_layout":
		return reflect.ValueOf(self._set_window_layout)
	case "_get_window_layout":
		return reflect.ValueOf(self._get_window_layout)
	case "_build":
		return reflect.ValueOf(self._build)
	case "_enable_plugin":
		return reflect.ValueOf(self._enable_plugin)
	case "_disable_plugin":
		return reflect.ValueOf(self._disable_plugin)
	default:
		return gd.VirtualByName(self.AsNode(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_forward_canvas_gui_input":
		return reflect.ValueOf(self._forward_canvas_gui_input)
	case "_forward_canvas_draw_over_viewport":
		return reflect.ValueOf(self._forward_canvas_draw_over_viewport)
	case "_forward_canvas_force_draw_over_viewport":
		return reflect.ValueOf(self._forward_canvas_force_draw_over_viewport)
	case "_forward_3d_gui_input":
		return reflect.ValueOf(self._forward_3d_gui_input)
	case "_forward_3d_draw_over_viewport":
		return reflect.ValueOf(self._forward_3d_draw_over_viewport)
	case "_forward_3d_force_draw_over_viewport":
		return reflect.ValueOf(self._forward_3d_force_draw_over_viewport)
	case "_get_plugin_name":
		return reflect.ValueOf(self._get_plugin_name)
	case "_get_plugin_icon":
		return reflect.ValueOf(self._get_plugin_icon)
	case "_has_main_screen":
		return reflect.ValueOf(self._has_main_screen)
	case "_make_visible":
		return reflect.ValueOf(self._make_visible)
	case "_edit":
		return reflect.ValueOf(self._edit)
	case "_handles":
		return reflect.ValueOf(self._handles)
	case "_get_state":
		return reflect.ValueOf(self._get_state)
	case "_set_state":
		return reflect.ValueOf(self._set_state)
	case "_clear":
		return reflect.ValueOf(self._clear)
	case "_get_unsaved_status":
		return reflect.ValueOf(self._get_unsaved_status)
	case "_save_external_data":
		return reflect.ValueOf(self._save_external_data)
	case "_apply_changes":
		return reflect.ValueOf(self._apply_changes)
	case "_get_breakpoints":
		return reflect.ValueOf(self._get_breakpoints)
	case "_set_window_layout":
		return reflect.ValueOf(self._set_window_layout)
	case "_get_window_layout":
		return reflect.ValueOf(self._get_window_layout)
	case "_build":
		return reflect.ValueOf(self._build)
	case "_enable_plugin":
		return reflect.ValueOf(self._enable_plugin)
	case "_disable_plugin":
		return reflect.ValueOf(self._disable_plugin)
	default:
		return gd.VirtualByName(self.AsNode(), name)
	}
}
func init() {
	classdb.Register("EditorPlugin", func(ptr gd.Object) any { return classdb.EditorPlugin(ptr) })
}

type CustomControlContainer = classdb.EditorPluginCustomControlContainer

const (
	/*Main editor toolbar, next to play buttons.*/
	ContainerToolbar CustomControlContainer = 0
	/*The toolbar that appears when 3D editor is active.*/
	ContainerSpatialEditorMenu CustomControlContainer = 1
	/*Left sidebar of the 3D editor.*/
	ContainerSpatialEditorSideLeft CustomControlContainer = 2
	/*Right sidebar of the 3D editor.*/
	ContainerSpatialEditorSideRight CustomControlContainer = 3
	/*Bottom panel of the 3D editor.*/
	ContainerSpatialEditorBottom CustomControlContainer = 4
	/*The toolbar that appears when 2D editor is active.*/
	ContainerCanvasEditorMenu CustomControlContainer = 5
	/*Left sidebar of the 2D editor.*/
	ContainerCanvasEditorSideLeft CustomControlContainer = 6
	/*Right sidebar of the 2D editor.*/
	ContainerCanvasEditorSideRight CustomControlContainer = 7
	/*Bottom panel of the 2D editor.*/
	ContainerCanvasEditorBottom CustomControlContainer = 8
	/*Bottom section of the inspector.*/
	ContainerInspectorBottom CustomControlContainer = 9
	/*Tab of Project Settings dialog, to the left of other tabs.*/
	ContainerProjectSettingTabLeft CustomControlContainer = 10
	/*Tab of Project Settings dialog, to the right of other tabs.*/
	ContainerProjectSettingTabRight CustomControlContainer = 11
)

type DockSlot = classdb.EditorPluginDockSlot

const (
	/*Dock slot, left side, upper-left (empty in default layout).*/
	DockSlotLeftUl DockSlot = 0
	/*Dock slot, left side, bottom-left (empty in default layout).*/
	DockSlotLeftBl DockSlot = 1
	/*Dock slot, left side, upper-right (in default layout includes Scene and Import docks).*/
	DockSlotLeftUr DockSlot = 2
	/*Dock slot, left side, bottom-right (in default layout includes FileSystem dock).*/
	DockSlotLeftBr DockSlot = 3
	/*Dock slot, right side, upper-left (in default layout includes Inspector, Node, and History docks).*/
	DockSlotRightUl DockSlot = 4
	/*Dock slot, right side, bottom-left (empty in default layout).*/
	DockSlotRightBl DockSlot = 5
	/*Dock slot, right side, upper-right (empty in default layout).*/
	DockSlotRightUr DockSlot = 6
	/*Dock slot, right side, bottom-right (empty in default layout).*/
	DockSlotRightBr DockSlot = 7
	/*Represents the size of the [enum DockSlot] enum.*/
	DockSlotMax DockSlot = 8
)

type AfterGUIInput = classdb.EditorPluginAfterGUIInput

const (
	/*Forwards the [InputEvent] to other EditorPlugins.*/
	AfterGuiInputPass AfterGUIInput = 0
	/*Prevents the [InputEvent] from reaching other Editor classes.*/
	AfterGuiInputStop AfterGUIInput = 1
	/*Pass the [InputEvent] to other editor plugins except the main [Node3D] one. This can be used to prevent node selection changes and work with sub-gizmos instead.*/
	AfterGuiInputCustom AfterGUIInput = 2
)
