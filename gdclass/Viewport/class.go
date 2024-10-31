package Viewport

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import "grow.graphics/gd/gdconst"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root
var _ gdconst.Side

/*
A [Viewport] creates a different view into the screen, or a sub-view inside another viewport. Child 2D nodes will display on it, and child Camera3D 3D nodes will render on it too.
Optionally, a viewport can have its own 2D or 3D world, so it doesn't share what it draws with other viewports.
Viewports can also choose to be audio listeners, so they generate positional audio depending on a 2D or 3D camera child of it.
Also, viewports can be assigned to different screens in case the devices have multiple screens.
Finally, viewports can also behave as render targets, in which case they will not be visible unless the associated texture is used to draw.
*/
type Instance [1]classdb.Viewport

/*
Returns the first valid [World2D] for this viewport, searching the [member world_2d] property of itself and any Viewport ancestor.
*/
func (self Instance) FindWorld2d() gdclass.World2D {
	return gdclass.World2D(class(self).FindWorld2d())
}

/*
Returns the transform from the viewport's coordinate system to the embedder's coordinate system.
*/
func (self Instance) GetFinalTransform() gd.Transform2D {
	return gd.Transform2D(class(self).GetFinalTransform())
}

/*
Returns the transform from the Viewport's coordinates to the screen coordinates of the containing window manager window.
*/
func (self Instance) GetScreenTransform() gd.Transform2D {
	return gd.Transform2D(class(self).GetScreenTransform())
}

/*
Returns the visible rectangle in global screen coordinates.
*/
func (self Instance) GetVisibleRect() gd.Rect2 {
	return gd.Rect2(class(self).GetVisibleRect())
}

/*
Returns rendering statistics of the given type. See [enum RenderInfoType] and [enum RenderInfo] for options.
*/
func (self Instance) GetRenderInfo(atype classdb.ViewportRenderInfoType, info classdb.ViewportRenderInfo) int {
	return int(int(class(self).GetRenderInfo(atype, info)))
}

/*
Returns the viewport's texture.
[b]Note:[/b] When trying to store the current texture (e.g. in a file), it might be completely black or outdated if used too early, especially when used in e.g. [method Node._ready]. To make sure the texture you get is correct, you can await [signal RenderingServer.frame_post_draw] signal.
[codeblock]
func _ready():

	await RenderingServer.frame_post_draw
	$Viewport.get_texture().get_image().save_png("user://Screenshot.png")

[/codeblock]
*/
func (self Instance) GetTexture() gdclass.ViewportTexture {
	return gdclass.ViewportTexture(class(self).GetTexture())
}

/*
Returns the viewport's RID from the [RenderingServer].
*/
func (self Instance) GetViewportRid() gd.RID {
	return gd.RID(class(self).GetViewportRid())
}

/*
Helper method which calls the [code]set_text()[/code] method on the currently focused [Control], provided that it is defined (e.g. if the focused Control is [Button] or [LineEdit]).
*/
func (self Instance) PushTextInput(text string) {
	class(self).PushTextInput(gd.NewString(text))
}

/*
Triggers the given [param event] in this [Viewport]. This can be used to pass an [InputEvent] between viewports, or to locally apply inputs that were sent over the network or saved to a file.
If [param in_local_coords] is [code]false[/code], the event's position is in the embedder's coordinates and will be converted to viewport coordinates. If [param in_local_coords] is [code]true[/code], the event's position is in viewport coordinates.
While this method serves a similar purpose as [method Input.parse_input_event], it does not remap the specified [param event] based on project settings like [member ProjectSettings.input_devices/pointing/emulate_touch_from_mouse].
Calling this method will propagate calls to child nodes for following methods in the given order:
- [method Node._input]
- [method Control._gui_input] for [Control] nodes
- [method Node._shortcut_input]
- [method Node._unhandled_key_input]
- [method Node._unhandled_input]
If an earlier method marks the input as handled via [method set_input_as_handled], any later method in this list will not be called.
If none of the methods handle the event and [member physics_object_picking] is [code]true[/code], the event is used for physics object picking.
*/
func (self Instance) PushInput(event gdclass.InputEvent) {
	class(self).PushInput(event, false)
}

/*
Triggers the given [param event] in this [Viewport]. This can be used to pass an [InputEvent] between viewports, or to locally apply inputs that were sent over the network or saved to a file.
If [param in_local_coords] is [code]false[/code], the event's position is in the embedder's coordinates and will be converted to viewport coordinates. If [param in_local_coords] is [code]true[/code], the event's position is in viewport coordinates.
Calling this method will propagate calls to child nodes for following methods in the given order:
- [method Node._shortcut_input]
- [method Node._unhandled_key_input]
- [method Node._unhandled_input]
If an earlier method marks the input as handled via [method set_input_as_handled], any later method in this list will not be called.
If none of the methods handle the event and [member physics_object_picking] is [code]true[/code], the event is used for physics object picking.
[b]Note:[/b] This method doesn't propagate input events to embedded [Window]s or [SubViewport]s.
*/
func (self Instance) PushUnhandledInput(event gdclass.InputEvent) {
	class(self).PushUnhandledInput(event, false)
}

/*
Returns the mouse's position in this [Viewport] using the coordinate system of this [Viewport].
*/
func (self Instance) GetMousePosition() gd.Vector2 {
	return gd.Vector2(class(self).GetMousePosition())
}

/*
Moves the mouse pointer to the specified position in this [Viewport] using the coordinate system of this [Viewport].
[b]Note:[/b] [method warp_mouse] is only supported on Windows, macOS and Linux. It has no effect on Android, iOS and Web.
*/
func (self Instance) WarpMouse(position gd.Vector2) {
	class(self).WarpMouse(position)
}

/*
Force instantly updating the display based on the current mouse cursor position. This includes updating the mouse cursor shape and sending necessary [signal Control.mouse_entered], [signal CollisionObject2D.mouse_entered], [signal CollisionObject3D.mouse_entered] and [signal Window.mouse_entered] signals and their respective [code]mouse_exited[/code] counterparts.
*/
func (self Instance) UpdateMouseCursorState() {
	class(self).UpdateMouseCursorState()
}

/*
Returns the drag data from the GUI, that was previously returned by [method Control._get_drag_data].
*/
func (self Instance) GuiGetDragData() gd.Variant {
	return gd.Variant(class(self).GuiGetDragData())
}

/*
Returns [code]true[/code] if the viewport is currently performing a drag operation.
Alternative to [constant Node.NOTIFICATION_DRAG_BEGIN] and [constant Node.NOTIFICATION_DRAG_END] when you prefer polling the value.
*/
func (self Instance) GuiIsDragging() bool {
	return bool(class(self).GuiIsDragging())
}

/*
Returns [code]true[/code] if the drag operation is successful.
*/
func (self Instance) GuiIsDragSuccessful() bool {
	return bool(class(self).GuiIsDragSuccessful())
}

/*
Removes the focus from the currently focused [Control] within this viewport. If no [Control] has the focus, does nothing.
*/
func (self Instance) GuiReleaseFocus() {
	class(self).GuiReleaseFocus()
}

/*
Returns the [Control] having the focus within this viewport. If no [Control] has the focus, returns null.
*/
func (self Instance) GuiGetFocusOwner() gdclass.Control {
	return gdclass.Control(class(self).GuiGetFocusOwner())
}

/*
Returns the [Control] that the mouse is currently hovering over in this viewport. If no [Control] has the cursor, returns null.
Typically the leaf [Control] node or deepest level of the subtree which claims hover. This is very useful when used together with [method Node.is_ancestor_of] to find if the mouse is within a control tree.
*/
func (self Instance) GuiGetHoveredControl() gdclass.Control {
	return gdclass.Control(class(self).GuiGetHoveredControl())
}

/*
Stops the input from propagating further down the [SceneTree].
[b]Note:[/b] This does not affect the methods in [Input], only the way events are propagated.
*/
func (self Instance) SetInputAsHandled() {
	class(self).SetInputAsHandled()
}

/*
Returns whether the current [InputEvent] has been handled. Input events are not handled until [method set_input_as_handled] has been called during the lifetime of an [InputEvent].
This is usually done as part of input handling methods like [method Node._input], [method Control._gui_input] or others, as well as in corresponding signal handlers.
If [member handle_input_locally] is set to [code]false[/code], this method will try finding the first parent viewport that is set to handle input locally, and return its value for [method is_input_handled] instead.
*/
func (self Instance) IsInputHandled() bool {
	return bool(class(self).IsInputHandled())
}

/*
Returns a list of the visible embedded [Window]s inside the viewport.
[b]Note:[/b] [Window]s inside other viewports will not be listed.
*/
func (self Instance) GetEmbeddedSubwindows() gd.Array {
	return gd.Array(class(self).GetEmbeddedSubwindows())
}

/*
Set/clear individual bits on the rendering layer mask. This simplifies editing this [Viewport]'s layers.
*/
func (self Instance) SetCanvasCullMaskBit(layer int, enable bool) {
	class(self).SetCanvasCullMaskBit(gd.Int(layer), enable)
}

/*
Returns an individual bit on the rendering layer mask.
*/
func (self Instance) GetCanvasCullMaskBit(layer int) bool {
	return bool(class(self).GetCanvasCullMaskBit(gd.Int(layer)))
}

/*
Returns the currently active 2D camera. Returns null if there are no active cameras.
*/
func (self Instance) GetCamera2d() gdclass.Camera2D {
	return gdclass.Camera2D(class(self).GetCamera2d())
}

/*
Returns the first valid [World3D] for this viewport, searching the [member world_3d] property of itself and any Viewport ancestor.
*/
func (self Instance) FindWorld3d() gdclass.World3D {
	return gdclass.World3D(class(self).FindWorld3d())
}

/*
Returns the currently active 3D camera.
*/
func (self Instance) GetCamera3d() gdclass.Camera3D {
	return gdclass.Camera3D(class(self).GetCamera3d())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.Viewport

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Viewport"))
	return Instance{classdb.Viewport(object)}
}

func (self Instance) Disable3d() bool {
	return bool(class(self).Is3dDisabled())
}

func (self Instance) SetDisable3d(value bool) {
	class(self).SetDisable3d(value)
}

func (self Instance) UseXr() bool {
	return bool(class(self).IsUsingXr())
}

func (self Instance) SetUseXr(value bool) {
	class(self).SetUseXr(value)
}

func (self Instance) OwnWorld3d() bool {
	return bool(class(self).IsUsingOwnWorld3d())
}

func (self Instance) SetOwnWorld3d(value bool) {
	class(self).SetUseOwnWorld3d(value)
}

func (self Instance) World3d() gdclass.World3D {
	return gdclass.World3D(class(self).GetWorld3d())
}

func (self Instance) SetWorld3d(value gdclass.World3D) {
	class(self).SetWorld3d(value)
}

func (self Instance) World2d() gdclass.World2D {
	return gdclass.World2D(class(self).GetWorld2d())
}

func (self Instance) SetWorld2d(value gdclass.World2D) {
	class(self).SetWorld2d(value)
}

func (self Instance) TransparentBg() bool {
	return bool(class(self).HasTransparentBackground())
}

func (self Instance) SetTransparentBg(value bool) {
	class(self).SetTransparentBackground(value)
}

func (self Instance) HandleInputLocally() bool {
	return bool(class(self).IsHandlingInputLocally())
}

func (self Instance) SetHandleInputLocally(value bool) {
	class(self).SetHandleInputLocally(value)
}

func (self Instance) Snap2dTransformsToPixel() bool {
	return bool(class(self).IsSnap2dTransformsToPixelEnabled())
}

func (self Instance) SetSnap2dTransformsToPixel(value bool) {
	class(self).SetSnap2dTransformsToPixel(value)
}

func (self Instance) Snap2dVerticesToPixel() bool {
	return bool(class(self).IsSnap2dVerticesToPixelEnabled())
}

func (self Instance) SetSnap2dVerticesToPixel(value bool) {
	class(self).SetSnap2dVerticesToPixel(value)
}

func (self Instance) Msaa2d() classdb.ViewportMSAA {
	return classdb.ViewportMSAA(class(self).GetMsaa2d())
}

func (self Instance) SetMsaa2d(value classdb.ViewportMSAA) {
	class(self).SetMsaa2d(value)
}

func (self Instance) Msaa3d() classdb.ViewportMSAA {
	return classdb.ViewportMSAA(class(self).GetMsaa3d())
}

func (self Instance) SetMsaa3d(value classdb.ViewportMSAA) {
	class(self).SetMsaa3d(value)
}

func (self Instance) ScreenSpaceAa() classdb.ViewportScreenSpaceAA {
	return classdb.ViewportScreenSpaceAA(class(self).GetScreenSpaceAa())
}

func (self Instance) SetScreenSpaceAa(value classdb.ViewportScreenSpaceAA) {
	class(self).SetScreenSpaceAa(value)
}

func (self Instance) UseTaa() bool {
	return bool(class(self).IsUsingTaa())
}

func (self Instance) SetUseTaa(value bool) {
	class(self).SetUseTaa(value)
}

func (self Instance) UseDebanding() bool {
	return bool(class(self).IsUsingDebanding())
}

func (self Instance) SetUseDebanding(value bool) {
	class(self).SetUseDebanding(value)
}

func (self Instance) UseOcclusionCulling() bool {
	return bool(class(self).IsUsingOcclusionCulling())
}

func (self Instance) SetUseOcclusionCulling(value bool) {
	class(self).SetUseOcclusionCulling(value)
}

func (self Instance) MeshLodThreshold() float64 {
	return float64(float64(class(self).GetMeshLodThreshold()))
}

func (self Instance) SetMeshLodThreshold(value float64) {
	class(self).SetMeshLodThreshold(gd.Float(value))
}

func (self Instance) DebugDraw() classdb.ViewportDebugDraw {
	return classdb.ViewportDebugDraw(class(self).GetDebugDraw())
}

func (self Instance) SetDebugDraw(value classdb.ViewportDebugDraw) {
	class(self).SetDebugDraw(value)
}

func (self Instance) UseHdr2d() bool {
	return bool(class(self).IsUsingHdr2d())
}

func (self Instance) SetUseHdr2d(value bool) {
	class(self).SetUseHdr2d(value)
}

func (self Instance) Scaling3dMode() classdb.ViewportScaling3DMode {
	return classdb.ViewportScaling3DMode(class(self).GetScaling3dMode())
}

func (self Instance) SetScaling3dMode(value classdb.ViewportScaling3DMode) {
	class(self).SetScaling3dMode(value)
}

func (self Instance) Scaling3dScale() float64 {
	return float64(float64(class(self).GetScaling3dScale()))
}

func (self Instance) SetScaling3dScale(value float64) {
	class(self).SetScaling3dScale(gd.Float(value))
}

func (self Instance) TextureMipmapBias() float64 {
	return float64(float64(class(self).GetTextureMipmapBias()))
}

func (self Instance) SetTextureMipmapBias(value float64) {
	class(self).SetTextureMipmapBias(gd.Float(value))
}

func (self Instance) FsrSharpness() float64 {
	return float64(float64(class(self).GetFsrSharpness()))
}

func (self Instance) SetFsrSharpness(value float64) {
	class(self).SetFsrSharpness(gd.Float(value))
}

func (self Instance) VrsMode() classdb.ViewportVRSMode {
	return classdb.ViewportVRSMode(class(self).GetVrsMode())
}

func (self Instance) SetVrsMode(value classdb.ViewportVRSMode) {
	class(self).SetVrsMode(value)
}

func (self Instance) VrsUpdateMode() classdb.ViewportVRSUpdateMode {
	return classdb.ViewportVRSUpdateMode(class(self).GetVrsUpdateMode())
}

func (self Instance) SetVrsUpdateMode(value classdb.ViewportVRSUpdateMode) {
	class(self).SetVrsUpdateMode(value)
}

func (self Instance) VrsTexture() gdclass.Texture2D {
	return gdclass.Texture2D(class(self).GetVrsTexture())
}

func (self Instance) SetVrsTexture(value gdclass.Texture2D) {
	class(self).SetVrsTexture(value)
}

func (self Instance) CanvasItemDefaultTextureFilter() classdb.ViewportDefaultCanvasItemTextureFilter {
	return classdb.ViewportDefaultCanvasItemTextureFilter(class(self).GetDefaultCanvasItemTextureFilter())
}

func (self Instance) SetCanvasItemDefaultTextureFilter(value classdb.ViewportDefaultCanvasItemTextureFilter) {
	class(self).SetDefaultCanvasItemTextureFilter(value)
}

func (self Instance) CanvasItemDefaultTextureRepeat() classdb.ViewportDefaultCanvasItemTextureRepeat {
	return classdb.ViewportDefaultCanvasItemTextureRepeat(class(self).GetDefaultCanvasItemTextureRepeat())
}

func (self Instance) SetCanvasItemDefaultTextureRepeat(value classdb.ViewportDefaultCanvasItemTextureRepeat) {
	class(self).SetDefaultCanvasItemTextureRepeat(value)
}

func (self Instance) AudioListenerEnable2d() bool {
	return bool(class(self).IsAudioListener2d())
}

func (self Instance) SetAudioListenerEnable2d(value bool) {
	class(self).SetAsAudioListener2d(value)
}

func (self Instance) AudioListenerEnable3d() bool {
	return bool(class(self).IsAudioListener3d())
}

func (self Instance) SetAudioListenerEnable3d(value bool) {
	class(self).SetAsAudioListener3d(value)
}

func (self Instance) PhysicsObjectPicking() bool {
	return bool(class(self).GetPhysicsObjectPicking())
}

func (self Instance) SetPhysicsObjectPicking(value bool) {
	class(self).SetPhysicsObjectPicking(value)
}

func (self Instance) PhysicsObjectPickingSort() bool {
	return bool(class(self).GetPhysicsObjectPickingSort())
}

func (self Instance) SetPhysicsObjectPickingSort(value bool) {
	class(self).SetPhysicsObjectPickingSort(value)
}

func (self Instance) PhysicsObjectPickingFirstOnly() bool {
	return bool(class(self).GetPhysicsObjectPickingFirstOnly())
}

func (self Instance) SetPhysicsObjectPickingFirstOnly(value bool) {
	class(self).SetPhysicsObjectPickingFirstOnly(value)
}

func (self Instance) GuiDisableInput() bool {
	return bool(class(self).IsInputDisabled())
}

func (self Instance) SetGuiDisableInput(value bool) {
	class(self).SetDisableInput(value)
}

func (self Instance) GuiSnapControlsToPixels() bool {
	return bool(class(self).IsSnapControlsToPixelsEnabled())
}

func (self Instance) SetGuiSnapControlsToPixels(value bool) {
	class(self).SetSnapControlsToPixels(value)
}

func (self Instance) GuiEmbedSubwindows() bool {
	return bool(class(self).IsEmbeddingSubwindows())
}

func (self Instance) SetGuiEmbedSubwindows(value bool) {
	class(self).SetEmbeddingSubwindows(value)
}

func (self Instance) SdfOversize() classdb.ViewportSDFOversize {
	return classdb.ViewportSDFOversize(class(self).GetSdfOversize())
}

func (self Instance) SetSdfOversize(value classdb.ViewportSDFOversize) {
	class(self).SetSdfOversize(value)
}

func (self Instance) SdfScale() classdb.ViewportSDFScale {
	return classdb.ViewportSDFScale(class(self).GetSdfScale())
}

func (self Instance) SetSdfScale(value classdb.ViewportSDFScale) {
	class(self).SetSdfScale(value)
}

func (self Instance) PositionalShadowAtlasSize() int {
	return int(int(class(self).GetPositionalShadowAtlasSize()))
}

func (self Instance) SetPositionalShadowAtlasSize(value int) {
	class(self).SetPositionalShadowAtlasSize(gd.Int(value))
}

func (self Instance) PositionalShadowAtlas16Bits() bool {
	return bool(class(self).GetPositionalShadowAtlas16Bits())
}

func (self Instance) SetPositionalShadowAtlas16Bits(value bool) {
	class(self).SetPositionalShadowAtlas16Bits(value)
}

func (self Instance) PositionalShadowAtlasQuad0() classdb.ViewportPositionalShadowAtlasQuadrantSubdiv {
	return classdb.ViewportPositionalShadowAtlasQuadrantSubdiv(class(self).GetPositionalShadowAtlasQuadrantSubdiv(0))
}

func (self Instance) SetPositionalShadowAtlasQuad0(value classdb.ViewportPositionalShadowAtlasQuadrantSubdiv) {
	class(self).SetPositionalShadowAtlasQuadrantSubdiv(0, value)
}

func (self Instance) PositionalShadowAtlasQuad1() classdb.ViewportPositionalShadowAtlasQuadrantSubdiv {
	return classdb.ViewportPositionalShadowAtlasQuadrantSubdiv(class(self).GetPositionalShadowAtlasQuadrantSubdiv(1))
}

func (self Instance) SetPositionalShadowAtlasQuad1(value classdb.ViewportPositionalShadowAtlasQuadrantSubdiv) {
	class(self).SetPositionalShadowAtlasQuadrantSubdiv(1, value)
}

func (self Instance) PositionalShadowAtlasQuad2() classdb.ViewportPositionalShadowAtlasQuadrantSubdiv {
	return classdb.ViewportPositionalShadowAtlasQuadrantSubdiv(class(self).GetPositionalShadowAtlasQuadrantSubdiv(2))
}

func (self Instance) SetPositionalShadowAtlasQuad2(value classdb.ViewportPositionalShadowAtlasQuadrantSubdiv) {
	class(self).SetPositionalShadowAtlasQuadrantSubdiv(2, value)
}

func (self Instance) PositionalShadowAtlasQuad3() classdb.ViewportPositionalShadowAtlasQuadrantSubdiv {
	return classdb.ViewportPositionalShadowAtlasQuadrantSubdiv(class(self).GetPositionalShadowAtlasQuadrantSubdiv(3))
}

func (self Instance) SetPositionalShadowAtlasQuad3(value classdb.ViewportPositionalShadowAtlasQuadrantSubdiv) {
	class(self).SetPositionalShadowAtlasQuadrantSubdiv(3, value)
}

func (self Instance) CanvasTransform() gd.Transform2D {
	return gd.Transform2D(class(self).GetCanvasTransform())
}

func (self Instance) SetCanvasTransform(value gd.Transform2D) {
	class(self).SetCanvasTransform(value)
}

func (self Instance) GlobalCanvasTransform() gd.Transform2D {
	return gd.Transform2D(class(self).GetGlobalCanvasTransform())
}

func (self Instance) SetGlobalCanvasTransform(value gd.Transform2D) {
	class(self).SetGlobalCanvasTransform(value)
}

func (self Instance) CanvasCullMask() int {
	return int(int(class(self).GetCanvasCullMask()))
}

func (self Instance) SetCanvasCullMask(value int) {
	class(self).SetCanvasCullMask(gd.Int(value))
}

//go:nosplit
func (self class) SetWorld2d(world_2d gdclass.World2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(world_2d[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_world_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetWorld2d() gdclass.World2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_world_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.World2D{classdb.World2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the first valid [World2D] for this viewport, searching the [member world_2d] property of itself and any Viewport ancestor.
*/
//go:nosplit
func (self class) FindWorld2d() gdclass.World2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_find_world_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.World2D{classdb.World2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCanvasTransform(xform gd.Transform2D) {
	var frame = callframe.New()
	callframe.Arg(frame, xform)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_canvas_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCanvasTransform() gd.Transform2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_canvas_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGlobalCanvasTransform(xform gd.Transform2D) {
	var frame = callframe.New()
	callframe.Arg(frame, xform)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_global_canvas_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetGlobalCanvasTransform() gd.Transform2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_global_canvas_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the transform from the viewport's coordinate system to the embedder's coordinate system.
*/
//go:nosplit
func (self class) GetFinalTransform() gd.Transform2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_final_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the transform from the Viewport's coordinates to the screen coordinates of the containing window manager window.
*/
//go:nosplit
func (self class) GetScreenTransform() gd.Transform2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Transform2D](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_screen_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the visible rectangle in global screen coordinates.
*/
//go:nosplit
func (self class) GetVisibleRect() gd.Rect2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Rect2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_visible_rect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTransparentBackground(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_transparent_background, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) HasTransparentBackground() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_has_transparent_background, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseHdr2d(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_use_hdr_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingHdr2d() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_is_using_hdr_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMsaa2d(msaa classdb.ViewportMSAA) {
	var frame = callframe.New()
	callframe.Arg(frame, msaa)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_msaa_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMsaa2d() classdb.ViewportMSAA {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ViewportMSAA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_msaa_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMsaa3d(msaa classdb.ViewportMSAA) {
	var frame = callframe.New()
	callframe.Arg(frame, msaa)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_msaa_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMsaa3d() classdb.ViewportMSAA {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ViewportMSAA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_msaa_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetScreenSpaceAa(screen_space_aa classdb.ViewportScreenSpaceAA) {
	var frame = callframe.New()
	callframe.Arg(frame, screen_space_aa)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_screen_space_aa, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetScreenSpaceAa() classdb.ViewportScreenSpaceAA {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ViewportScreenSpaceAA](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_screen_space_aa, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseTaa(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_use_taa, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingTaa() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_is_using_taa, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseDebanding(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_use_debanding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingDebanding() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_is_using_debanding, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseOcclusionCulling(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_use_occlusion_culling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingOcclusionCulling() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_is_using_occlusion_culling, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDebugDraw(debug_draw classdb.ViewportDebugDraw) {
	var frame = callframe.New()
	callframe.Arg(frame, debug_draw)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_debug_draw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDebugDraw() classdb.ViewportDebugDraw {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ViewportDebugDraw](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_debug_draw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns rendering statistics of the given type. See [enum RenderInfoType] and [enum RenderInfo] for options.
*/
//go:nosplit
func (self class) GetRenderInfo(atype classdb.ViewportRenderInfoType, info classdb.ViewportRenderInfo) gd.Int {
	var frame = callframe.New()
	callframe.Arg(frame, atype)
	callframe.Arg(frame, info)
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_render_info, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the viewport's texture.
[b]Note:[/b] When trying to store the current texture (e.g. in a file), it might be completely black or outdated if used too early, especially when used in e.g. [method Node._ready]. To make sure the texture you get is correct, you can await [signal RenderingServer.frame_post_draw] signal.
[codeblock]
func _ready():
    await RenderingServer.frame_post_draw
    $Viewport.get_texture().get_image().save_png("user://Screenshot.png")
[/codeblock]
*/
//go:nosplit
func (self class) GetTexture() gdclass.ViewportTexture {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.ViewportTexture{classdb.ViewportTexture(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPhysicsObjectPicking(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_physics_object_picking, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPhysicsObjectPicking() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_physics_object_picking, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPhysicsObjectPickingSort(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_physics_object_picking_sort, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPhysicsObjectPickingSort() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_physics_object_picking_sort, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPhysicsObjectPickingFirstOnly(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_physics_object_picking_first_only, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPhysicsObjectPickingFirstOnly() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_physics_object_picking_first_only, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the viewport's RID from the [RenderingServer].
*/
//go:nosplit
func (self class) GetViewportRid() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_viewport_rid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Helper method which calls the [code]set_text()[/code] method on the currently focused [Control], provided that it is defined (e.g. if the focused Control is [Button] or [LineEdit]).
*/
//go:nosplit
func (self class) PushTextInput(text gd.String) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(text))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_push_text_input, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Triggers the given [param event] in this [Viewport]. This can be used to pass an [InputEvent] between viewports, or to locally apply inputs that were sent over the network or saved to a file.
If [param in_local_coords] is [code]false[/code], the event's position is in the embedder's coordinates and will be converted to viewport coordinates. If [param in_local_coords] is [code]true[/code], the event's position is in viewport coordinates.
While this method serves a similar purpose as [method Input.parse_input_event], it does not remap the specified [param event] based on project settings like [member ProjectSettings.input_devices/pointing/emulate_touch_from_mouse].
Calling this method will propagate calls to child nodes for following methods in the given order:
- [method Node._input]
- [method Control._gui_input] for [Control] nodes
- [method Node._shortcut_input]
- [method Node._unhandled_key_input]
- [method Node._unhandled_input]
If an earlier method marks the input as handled via [method set_input_as_handled], any later method in this list will not be called.
If none of the methods handle the event and [member physics_object_picking] is [code]true[/code], the event is used for physics object picking.
*/
//go:nosplit
func (self class) PushInput(event gdclass.InputEvent, in_local_coords bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(event[0])[0])
	callframe.Arg(frame, in_local_coords)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_push_input, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Triggers the given [param event] in this [Viewport]. This can be used to pass an [InputEvent] between viewports, or to locally apply inputs that were sent over the network or saved to a file.
If [param in_local_coords] is [code]false[/code], the event's position is in the embedder's coordinates and will be converted to viewport coordinates. If [param in_local_coords] is [code]true[/code], the event's position is in viewport coordinates.
Calling this method will propagate calls to child nodes for following methods in the given order:
- [method Node._shortcut_input]
- [method Node._unhandled_key_input]
- [method Node._unhandled_input]
If an earlier method marks the input as handled via [method set_input_as_handled], any later method in this list will not be called.
If none of the methods handle the event and [member physics_object_picking] is [code]true[/code], the event is used for physics object picking.
[b]Note:[/b] This method doesn't propagate input events to embedded [Window]s or [SubViewport]s.
*/
//go:nosplit
func (self class) PushUnhandledInput(event gdclass.InputEvent, in_local_coords bool) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(event[0])[0])
	callframe.Arg(frame, in_local_coords)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_push_unhandled_input, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the mouse's position in this [Viewport] using the coordinate system of this [Viewport].
*/
//go:nosplit
func (self class) GetMousePosition() gd.Vector2 {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Vector2](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_mouse_position, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Moves the mouse pointer to the specified position in this [Viewport] using the coordinate system of this [Viewport].
[b]Note:[/b] [method warp_mouse] is only supported on Windows, macOS and Linux. It has no effect on Android, iOS and Web.
*/
//go:nosplit
func (self class) WarpMouse(position gd.Vector2) {
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_warp_mouse, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Force instantly updating the display based on the current mouse cursor position. This includes updating the mouse cursor shape and sending necessary [signal Control.mouse_entered], [signal CollisionObject2D.mouse_entered], [signal CollisionObject3D.mouse_entered] and [signal Window.mouse_entered] signals and their respective [code]mouse_exited[/code] counterparts.
*/
//go:nosplit
func (self class) UpdateMouseCursorState() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_update_mouse_cursor_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the drag data from the GUI, that was previously returned by [method Control._get_drag_data].
*/
//go:nosplit
func (self class) GuiGetDragData() gd.Variant {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[3]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_gui_get_drag_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Variant](r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the viewport is currently performing a drag operation.
Alternative to [constant Node.NOTIFICATION_DRAG_BEGIN] and [constant Node.NOTIFICATION_DRAG_END] when you prefer polling the value.
*/
//go:nosplit
func (self class) GuiIsDragging() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_gui_is_dragging, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if the drag operation is successful.
*/
//go:nosplit
func (self class) GuiIsDragSuccessful() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_gui_is_drag_successful, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes the focus from the currently focused [Control] within this viewport. If no [Control] has the focus, does nothing.
*/
//go:nosplit
func (self class) GuiReleaseFocus() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_gui_release_focus, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the [Control] having the focus within this viewport. If no [Control] has the focus, returns null.
*/
//go:nosplit
func (self class) GuiGetFocusOwner() gdclass.Control {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_gui_get_focus_owner, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Control{classdb.Control(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the [Control] that the mouse is currently hovering over in this viewport. If no [Control] has the cursor, returns null.
Typically the leaf [Control] node or deepest level of the subtree which claims hover. This is very useful when used together with [method Node.is_ancestor_of] to find if the mouse is within a control tree.
*/
//go:nosplit
func (self class) GuiGetHoveredControl() gdclass.Control {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_gui_get_hovered_control, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Control{classdb.Control(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDisableInput(disable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, disable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_disable_input, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsInputDisabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_is_input_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPositionalShadowAtlasSize(size gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_positional_shadow_atlas_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPositionalShadowAtlasSize() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_positional_shadow_atlas_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPositionalShadowAtlas16Bits(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_positional_shadow_atlas_16_bits, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetPositionalShadowAtlas16Bits() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_positional_shadow_atlas_16_bits, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSnapControlsToPixels(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_snap_controls_to_pixels, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsSnapControlsToPixelsEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_is_snap_controls_to_pixels_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSnap2dTransformsToPixel(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_snap_2d_transforms_to_pixel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsSnap2dTransformsToPixelEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_is_snap_2d_transforms_to_pixel_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSnap2dVerticesToPixel(enabled bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_snap_2d_vertices_to_pixel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsSnap2dVerticesToPixelEnabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_is_snap_2d_vertices_to_pixel_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the number of subdivisions to use in the specified quadrant. A higher number of subdivisions allows you to have more shadows in the scene at once, but reduces the quality of the shadows. A good practice is to have quadrants with a varying number of subdivisions and to have as few subdivisions as possible.
*/
//go:nosplit
func (self class) SetPositionalShadowAtlasQuadrantSubdiv(quadrant gd.Int, subdiv classdb.ViewportPositionalShadowAtlasQuadrantSubdiv) {
	var frame = callframe.New()
	callframe.Arg(frame, quadrant)
	callframe.Arg(frame, subdiv)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_positional_shadow_atlas_quadrant_subdiv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns the positional shadow atlas quadrant subdivision of the specified quadrant.
*/
//go:nosplit
func (self class) GetPositionalShadowAtlasQuadrantSubdiv(quadrant gd.Int) classdb.ViewportPositionalShadowAtlasQuadrantSubdiv {
	var frame = callframe.New()
	callframe.Arg(frame, quadrant)
	var r_ret = callframe.Ret[classdb.ViewportPositionalShadowAtlasQuadrantSubdiv](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_positional_shadow_atlas_quadrant_subdiv, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Stops the input from propagating further down the [SceneTree].
[b]Note:[/b] This does not affect the methods in [Input], only the way events are propagated.
*/
//go:nosplit
func (self class) SetInputAsHandled() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_input_as_handled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns whether the current [InputEvent] has been handled. Input events are not handled until [method set_input_as_handled] has been called during the lifetime of an [InputEvent].
This is usually done as part of input handling methods like [method Node._input], [method Control._gui_input] or others, as well as in corresponding signal handlers.
If [member handle_input_locally] is set to [code]false[/code], this method will try finding the first parent viewport that is set to handle input locally, and return its value for [method is_input_handled] instead.
*/
//go:nosplit
func (self class) IsInputHandled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_is_input_handled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHandleInputLocally(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_handle_input_locally, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsHandlingInputLocally() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_is_handling_input_locally, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDefaultCanvasItemTextureFilter(mode classdb.ViewportDefaultCanvasItemTextureFilter) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_default_canvas_item_texture_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDefaultCanvasItemTextureFilter() classdb.ViewportDefaultCanvasItemTextureFilter {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ViewportDefaultCanvasItemTextureFilter](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_default_canvas_item_texture_filter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetEmbeddingSubwindows(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_embedding_subwindows, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsEmbeddingSubwindows() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_is_embedding_subwindows, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a list of the visible embedded [Window]s inside the viewport.
[b]Note:[/b] [Window]s inside other viewports will not be listed.
*/
//go:nosplit
func (self class) GetEmbeddedSubwindows() gd.Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_embedded_subwindows, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = pointers.New[gd.Array](r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCanvasCullMask(mask gd.Int) {
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_canvas_cull_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetCanvasCullMask() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_canvas_cull_mask, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Set/clear individual bits on the rendering layer mask. This simplifies editing this [Viewport]'s layers.
*/
//go:nosplit
func (self class) SetCanvasCullMaskBit(layer gd.Int, enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_canvas_cull_mask_bit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

/*
Returns an individual bit on the rendering layer mask.
*/
//go:nosplit
func (self class) GetCanvasCullMaskBit(layer gd.Int) bool {
	var frame = callframe.New()
	callframe.Arg(frame, layer)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_canvas_cull_mask_bit, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDefaultCanvasItemTextureRepeat(mode classdb.ViewportDefaultCanvasItemTextureRepeat) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_default_canvas_item_texture_repeat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetDefaultCanvasItemTextureRepeat() classdb.ViewportDefaultCanvasItemTextureRepeat {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ViewportDefaultCanvasItemTextureRepeat](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_default_canvas_item_texture_repeat, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSdfOversize(oversize classdb.ViewportSDFOversize) {
	var frame = callframe.New()
	callframe.Arg(frame, oversize)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_sdf_oversize, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSdfOversize() classdb.ViewportSDFOversize {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ViewportSDFOversize](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_sdf_oversize, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSdfScale(scale classdb.ViewportSDFScale) {
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_sdf_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetSdfScale() classdb.ViewportSDFScale {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ViewportSDFScale](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_sdf_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMeshLodThreshold(pixels gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, pixels)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_mesh_lod_threshold, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetMeshLodThreshold() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_mesh_lod_threshold, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAsAudioListener2d(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_as_audio_listener_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsAudioListener2d() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_is_audio_listener_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the currently active 2D camera. Returns null if there are no active cameras.
*/
//go:nosplit
func (self class) GetCamera2d() gdclass.Camera2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_camera_2d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Camera2D{classdb.Camera2D(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetWorld3d(world_3d gdclass.World3D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(world_3d[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_world_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetWorld3d() gdclass.World3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_world_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.World3D{classdb.World3D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Returns the first valid [World3D] for this viewport, searching the [member world_3d] property of itself and any Viewport ancestor.
*/
//go:nosplit
func (self class) FindWorld3d() gdclass.World3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_find_world_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.World3D{classdb.World3D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseOwnWorld3d(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_use_own_world_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingOwnWorld3d() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_is_using_own_world_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the currently active 3D camera.
*/
//go:nosplit
func (self class) GetCamera3d() gdclass.Camera3D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_camera_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Camera3D{classdb.Camera3D(gd.PointerMustAssertInstanceID(r_ret.Get()))}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAsAudioListener3d(enable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_as_audio_listener_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsAudioListener3d() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_is_audio_listener_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDisable3d(disable bool) {
	var frame = callframe.New()
	callframe.Arg(frame, disable)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_disable_3d, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) Is3dDisabled() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_is_3d_disabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUseXr(use bool) {
	var frame = callframe.New()
	callframe.Arg(frame, use)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_use_xr, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) IsUsingXr() bool {
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_is_using_xr, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetScaling3dMode(scaling_3d_mode classdb.ViewportScaling3DMode) {
	var frame = callframe.New()
	callframe.Arg(frame, scaling_3d_mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_scaling_3d_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetScaling3dMode() classdb.ViewportScaling3DMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ViewportScaling3DMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_scaling_3d_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetScaling3dScale(scale gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_scaling_3d_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetScaling3dScale() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_scaling_3d_scale, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFsrSharpness(fsr_sharpness gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, fsr_sharpness)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_fsr_sharpness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetFsrSharpness() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_fsr_sharpness, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTextureMipmapBias(texture_mipmap_bias gd.Float) {
	var frame = callframe.New()
	callframe.Arg(frame, texture_mipmap_bias)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_texture_mipmap_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetTextureMipmapBias() gd.Float {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_texture_mipmap_bias, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVrsMode(mode classdb.ViewportVRSMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_vrs_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVrsMode() classdb.ViewportVRSMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ViewportVRSMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_vrs_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVrsUpdateMode(mode classdb.ViewportVRSUpdateMode) {
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_vrs_update_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVrsUpdateMode() classdb.ViewportVRSUpdateMode {
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ViewportVRSUpdateMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_vrs_update_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVrsTexture(texture gdclass.Texture2D) {
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_set_vrs_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) GetVrsTexture() gdclass.Texture2D {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Viewport.Bind_get_vrs_texture, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.Texture2D{classdb.Texture2D(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self Instance) OnSizeChanged(cb func()) {
	self[0].AsObject().Connect(gd.NewStringName("size_changed"), gd.NewCallable(cb), 0)
}

func (self Instance) OnGuiFocusChanged(cb func(node gdclass.Control)) {
	self[0].AsObject().Connect(gd.NewStringName("gui_focus_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsViewport() Advanced     { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsViewport() Instance  { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced    { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsNode(), name)
	}
}
func init() { classdb.Register("Viewport", func(ptr gd.Object) any { return classdb.Viewport(ptr) }) }

type PositionalShadowAtlasQuadrantSubdiv = classdb.ViewportPositionalShadowAtlasQuadrantSubdiv

const (
	/*This quadrant will not be used.*/
	ShadowAtlasQuadrantSubdivDisabled PositionalShadowAtlasQuadrantSubdiv = 0
	/*This quadrant will only be used by one shadow map.*/
	ShadowAtlasQuadrantSubdiv1 PositionalShadowAtlasQuadrantSubdiv = 1
	/*This quadrant will be split in 4 and used by up to 4 shadow maps.*/
	ShadowAtlasQuadrantSubdiv4 PositionalShadowAtlasQuadrantSubdiv = 2
	/*This quadrant will be split 16 ways and used by up to 16 shadow maps.*/
	ShadowAtlasQuadrantSubdiv16 PositionalShadowAtlasQuadrantSubdiv = 3
	/*This quadrant will be split 64 ways and used by up to 64 shadow maps.*/
	ShadowAtlasQuadrantSubdiv64 PositionalShadowAtlasQuadrantSubdiv = 4
	/*This quadrant will be split 256 ways and used by up to 256 shadow maps. Unless the [member positional_shadow_atlas_size] is very high, the shadows in this quadrant will be very low resolution.*/
	ShadowAtlasQuadrantSubdiv256 PositionalShadowAtlasQuadrantSubdiv = 5
	/*This quadrant will be split 1024 ways and used by up to 1024 shadow maps. Unless the [member positional_shadow_atlas_size] is very high, the shadows in this quadrant will be very low resolution.*/
	ShadowAtlasQuadrantSubdiv1024 PositionalShadowAtlasQuadrantSubdiv = 6
	/*Represents the size of the [enum PositionalShadowAtlasQuadrantSubdiv] enum.*/
	ShadowAtlasQuadrantSubdivMax PositionalShadowAtlasQuadrantSubdiv = 7
)

type Scaling3DMode = classdb.ViewportScaling3DMode

const (
	/*Use bilinear scaling for the viewport's 3D buffer. The amount of scaling can be set using [member scaling_3d_scale]. Values less than [code]1.0[/code] will result in undersampling while values greater than [code]1.0[/code] will result in supersampling. A value of [code]1.0[/code] disables scaling.*/
	Scaling3dModeBilinear Scaling3DMode = 0
	/*Use AMD FidelityFX Super Resolution 1.0 upscaling for the viewport's 3D buffer. The amount of scaling can be set using [member scaling_3d_scale]. Values less than [code]1.0[/code] will be result in the viewport being upscaled using FSR. Values greater than [code]1.0[/code] are not supported and bilinear downsampling will be used instead. A value of [code]1.0[/code] disables scaling.*/
	Scaling3dModeFsr Scaling3DMode = 1
	/*Use AMD FidelityFX Super Resolution 2.2 upscaling for the viewport's 3D buffer. The amount of scaling can be set using [member Viewport.scaling_3d_scale]. Values less than [code]1.0[/code] will be result in the viewport being upscaled using FSR2. Values greater than [code]1.0[/code] are not supported and bilinear downsampling will be used instead. A value of [code]1.0[/code] will use FSR2 at native resolution as a TAA solution.*/
	Scaling3dModeFsr2 Scaling3DMode = 2
	/*Represents the size of the [enum Scaling3DMode] enum.*/
	Scaling3dModeMax Scaling3DMode = 3
)

type MSAA = classdb.ViewportMSAA

const (
	/*Multisample antialiasing mode disabled. This is the default value, and is also the fastest setting.*/
	MsaaDisabled MSAA = 0
	/*Use 2× Multisample Antialiasing. This has a moderate performance cost. It helps reduce aliasing noticeably, but 4× MSAA still looks substantially better.*/
	Msaa2x MSAA = 1
	/*Use 4× Multisample Antialiasing. This has a significant performance cost, and is generally a good compromise between performance and quality.*/
	Msaa4x MSAA = 2
	/*Use 8× Multisample Antialiasing. This has a very high performance cost. The difference between 4× and 8× MSAA may not always be visible in real gameplay conditions. Likely unsupported on low-end and older hardware.*/
	Msaa8x MSAA = 3
	/*Represents the size of the [enum MSAA] enum.*/
	MsaaMax MSAA = 4
)

type ScreenSpaceAA = classdb.ViewportScreenSpaceAA

const (
	/*Do not perform any antialiasing in the full screen post-process.*/
	ScreenSpaceAaDisabled ScreenSpaceAA = 0
	/*Use fast approximate antialiasing. FXAA is a popular screen-space antialiasing method, which is fast but will make the image look blurry, especially at lower resolutions. It can still work relatively well at large resolutions such as 1440p and 4K.*/
	ScreenSpaceAaFxaa ScreenSpaceAA = 1
	/*Represents the size of the [enum ScreenSpaceAA] enum.*/
	ScreenSpaceAaMax ScreenSpaceAA = 2
)

type RenderInfo = classdb.ViewportRenderInfo

const (
	/*Amount of objects in frame.*/
	RenderInfoObjectsInFrame RenderInfo = 0
	/*Amount of vertices in frame.*/
	RenderInfoPrimitivesInFrame RenderInfo = 1
	/*Amount of draw calls in frame.*/
	RenderInfoDrawCallsInFrame RenderInfo = 2
	/*Represents the size of the [enum RenderInfo] enum.*/
	RenderInfoMax RenderInfo = 3
)

type RenderInfoType = classdb.ViewportRenderInfoType

const (
	/*Visible render pass (excluding shadows).*/
	RenderInfoTypeVisible RenderInfoType = 0
	/*Shadow render pass. Objects will be rendered several times depending on the number of amounts of lights with shadows and the number of directional shadow splits.*/
	RenderInfoTypeShadow RenderInfoType = 1
	/*Canvas item rendering. This includes all 2D rendering.*/
	RenderInfoTypeCanvas RenderInfoType = 2
	/*Represents the size of the [enum RenderInfoType] enum.*/
	RenderInfoTypeMax RenderInfoType = 3
)

type DebugDraw = classdb.ViewportDebugDraw

const (
	/*Objects are displayed normally.*/
	DebugDrawDisabled DebugDraw = 0
	/*Objects are displayed without light information.*/
	DebugDrawUnshaded DebugDraw = 1
	/*Objects are displayed without textures and only with lighting information.*/
	DebugDrawLighting DebugDraw = 2
	/*Objects are displayed semi-transparent with additive blending so you can see where they are drawing over top of one another. A higher overdraw means you are wasting performance on drawing pixels that are being hidden behind others.*/
	DebugDrawOverdraw DebugDraw = 3
	/*Objects are displayed as wireframe models.*/
	DebugDrawWireframe DebugDraw = 4
	/*Objects are displayed without lighting information and their textures replaced by normal mapping.*/
	DebugDrawNormalBuffer DebugDraw = 5
	/*Objects are displayed with only the albedo value from [VoxelGI]s.*/
	DebugDrawVoxelGiAlbedo DebugDraw = 6
	/*Objects are displayed with only the lighting value from [VoxelGI]s.*/
	DebugDrawVoxelGiLighting DebugDraw = 7
	/*Objects are displayed with only the emission color from [VoxelGI]s.*/
	DebugDrawVoxelGiEmission DebugDraw = 8
	/*Draws the shadow atlas that stores shadows from [OmniLight3D]s and [SpotLight3D]s in the upper left quadrant of the [Viewport].*/
	DebugDrawShadowAtlas DebugDraw = 9
	/*Draws the shadow atlas that stores shadows from [DirectionalLight3D]s in the upper left quadrant of the [Viewport].*/
	DebugDrawDirectionalShadowAtlas DebugDraw = 10
	/*Draws the scene luminance buffer (if available) in the upper left quadrant of the [Viewport].*/
	DebugDrawSceneLuminance DebugDraw = 11
	/*Draws the screen-space ambient occlusion texture instead of the scene so that you can clearly see how it is affecting objects. In order for this display mode to work, you must have [member Environment.ssao_enabled] set in your [WorldEnvironment].*/
	DebugDrawSsao DebugDraw = 12
	/*Draws the screen-space indirect lighting texture instead of the scene so that you can clearly see how it is affecting objects. In order for this display mode to work, you must have [member Environment.ssil_enabled] set in your [WorldEnvironment].*/
	DebugDrawSsil DebugDraw = 13
	/*Colors each PSSM split for the [DirectionalLight3D]s in the scene a different color so you can see where the splits are. In order, they will be colored red, green, blue, and yellow.*/
	DebugDrawPssmSplits DebugDraw = 14
	/*Draws the decal atlas used by [Decal]s and light projector textures in the upper left quadrant of the [Viewport].*/
	DebugDrawDecalAtlas DebugDraw = 15
	/*Draws the cascades used to render signed distance field global illumination (SDFGI).
	  Does nothing if the current environment's [member Environment.sdfgi_enabled] is [code]false[/code] or SDFGI is not supported on the platform.*/
	DebugDrawSdfgi DebugDraw = 16
	/*Draws the probes used for signed distance field global illumination (SDFGI).
	  Does nothing if the current environment's [member Environment.sdfgi_enabled] is [code]false[/code] or SDFGI is not supported on the platform.*/
	DebugDrawSdfgiProbes DebugDraw = 17
	/*Draws the buffer used for global illumination (GI).*/
	DebugDrawGiBuffer DebugDraw = 18
	/*Draws all of the objects at their highest polycount, without low level of detail (LOD).*/
	DebugDrawDisableLod DebugDraw = 19
	/*Draws the cluster used by [OmniLight3D] nodes to optimize light rendering.*/
	DebugDrawClusterOmniLights DebugDraw = 20
	/*Draws the cluster used by [SpotLight3D] nodes to optimize light rendering.*/
	DebugDrawClusterSpotLights DebugDraw = 21
	/*Draws the cluster used by [Decal] nodes to optimize decal rendering.*/
	DebugDrawClusterDecals DebugDraw = 22
	/*Draws the cluster used by [ReflectionProbe] nodes to optimize decal rendering.*/
	DebugDrawClusterReflectionProbes DebugDraw = 23
	/*Draws the buffer used for occlusion culling.*/
	DebugDrawOccluders DebugDraw = 24
	/*Draws vector lines over the viewport to indicate the movement of pixels between frames.*/
	DebugDrawMotionVectors DebugDraw = 25
	/*Draws the internal resolution buffer of the scene before post-processing is applied.*/
	DebugDrawInternalBuffer DebugDraw = 26
)

type DefaultCanvasItemTextureFilter = classdb.ViewportDefaultCanvasItemTextureFilter

const (
	/*The texture filter reads from the nearest pixel only. This makes the texture look pixelated from up close, and grainy from a distance (due to mipmaps not being sampled).*/
	DefaultCanvasItemTextureFilterNearest DefaultCanvasItemTextureFilter = 0
	/*The texture filter blends between the nearest 4 pixels. This makes the texture look smooth from up close, and grainy from a distance (due to mipmaps not being sampled).*/
	DefaultCanvasItemTextureFilterLinear DefaultCanvasItemTextureFilter = 1
	/*The texture filter blends between the nearest 4 pixels and between the nearest 2 mipmaps (or uses the nearest mipmap if [member ProjectSettings.rendering/textures/default_filters/use_nearest_mipmap_filter] is [code]true[/code]). This makes the texture look smooth from up close, and smooth from a distance.
	  Use this for non-pixel art textures that may be viewed at a low scale (e.g. due to [Camera2D] zoom or sprite scaling), as mipmaps are important to smooth out pixels that are smaller than on-screen pixels.*/
	DefaultCanvasItemTextureFilterLinearWithMipmaps DefaultCanvasItemTextureFilter = 2
	/*The texture filter reads from the nearest pixel and blends between the nearest 2 mipmaps (or uses the nearest mipmap if [member ProjectSettings.rendering/textures/default_filters/use_nearest_mipmap_filter] is [code]true[/code]). This makes the texture look pixelated from up close, and smooth from a distance.
	  Use this for non-pixel art textures that may be viewed at a low scale (e.g. due to [Camera2D] zoom or sprite scaling), as mipmaps are important to smooth out pixels that are smaller than on-screen pixels.*/
	DefaultCanvasItemTextureFilterNearestWithMipmaps DefaultCanvasItemTextureFilter = 3
	/*Represents the size of the [enum DefaultCanvasItemTextureFilter] enum.*/
	DefaultCanvasItemTextureFilterMax DefaultCanvasItemTextureFilter = 4
)

type DefaultCanvasItemTextureRepeat = classdb.ViewportDefaultCanvasItemTextureRepeat

const (
	/*Disables textures repeating. Instead, when reading UVs outside the 0-1 range, the value will be clamped to the edge of the texture, resulting in a stretched out look at the borders of the texture.*/
	DefaultCanvasItemTextureRepeatDisabled DefaultCanvasItemTextureRepeat = 0
	/*Enables the texture to repeat when UV coordinates are outside the 0-1 range. If using one of the linear filtering modes, this can result in artifacts at the edges of a texture when the sampler filters across the edges of the texture.*/
	DefaultCanvasItemTextureRepeatEnabled DefaultCanvasItemTextureRepeat = 1
	/*Flip the texture when repeating so that the edge lines up instead of abruptly changing.*/
	DefaultCanvasItemTextureRepeatMirror DefaultCanvasItemTextureRepeat = 2
	/*Represents the size of the [enum DefaultCanvasItemTextureRepeat] enum.*/
	DefaultCanvasItemTextureRepeatMax DefaultCanvasItemTextureRepeat = 3
)

type SDFOversize = classdb.ViewportSDFOversize

const (
	/*The signed distance field only covers the viewport's own rectangle.*/
	SdfOversize100Percent SDFOversize = 0
	/*The signed distance field is expanded to cover 20% of the viewport's size around the borders.*/
	SdfOversize120Percent SDFOversize = 1
	/*The signed distance field is expanded to cover 50% of the viewport's size around the borders.*/
	SdfOversize150Percent SDFOversize = 2
	/*The signed distance field is expanded to cover 100% (double) of the viewport's size around the borders.*/
	SdfOversize200Percent SDFOversize = 3
	/*Represents the size of the [enum SDFOversize] enum.*/
	SdfOversizeMax SDFOversize = 4
)

type SDFScale = classdb.ViewportSDFScale

const (
	/*The signed distance field is rendered at full resolution.*/
	SdfScale100Percent SDFScale = 0
	/*The signed distance field is rendered at half the resolution of this viewport.*/
	SdfScale50Percent SDFScale = 1
	/*The signed distance field is rendered at a quarter the resolution of this viewport.*/
	SdfScale25Percent SDFScale = 2
	/*Represents the size of the [enum SDFScale] enum.*/
	SdfScaleMax SDFScale = 3
)

type VRSMode = classdb.ViewportVRSMode

const (
	/*Variable Rate Shading is disabled.*/
	VrsDisabled VRSMode = 0
	/*Variable Rate Shading uses a texture. Note, for stereoscopic use a texture atlas with a texture for each view.*/
	VrsTexture VRSMode = 1
	/*Variable Rate Shading's texture is supplied by the primary [XRInterface].*/
	VrsXr VRSMode = 2
	/*Represents the size of the [enum VRSMode] enum.*/
	VrsMax VRSMode = 3
)

type VRSUpdateMode = classdb.ViewportVRSUpdateMode

const (
	/*The input texture for variable rate shading will not be processed.*/
	VrsUpdateDisabled VRSUpdateMode = 0
	/*The input texture for variable rate shading will be processed once.*/
	VrsUpdateOnce VRSUpdateMode = 1
	/*The input texture for variable rate shading will be processed each frame.*/
	VrsUpdateAlways VRSUpdateMode = 2
	/*Represents the size of the [enum VRSUpdateMode] enum.*/
	VrsUpdateMax VRSUpdateMode = 3
)
