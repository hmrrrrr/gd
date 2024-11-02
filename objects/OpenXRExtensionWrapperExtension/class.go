package OpenXRExtensionWrapperExtension

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/pointers"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/objects"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
[OpenXRExtensionWrapperExtension] allows clients to implement OpenXR extensions with GDExtension. The extension should be registered with [method register_extension_wrapper].

	// OpenXRExtensionWrapperExtension methods that can be overridden by a [Class] that extends it.
	type OpenXRExtensionWrapperExtension interface {
		//Returns a [Dictionary] of OpenXR extensions related to this extension. The [Dictionary] should contain the name of the extension, mapped to a [code]bool *[/code] cast to an integer:
		//- If the [code]bool *[/code] is a [code]nullptr[/code] this extension is mandatory.
		//- If the [code]bool *[/code] points to a boolean, the boolean will be updated to [code]true[/code] if the extension is enabled.
		GetRequestedExtensions() gd.Dictionary
		//Adds additional data structures when interogating OpenXR system abilities.
		SetSystemPropertiesAndGetNextPointer(next_pointer unsafe.Pointer) int
		//Adds additional data structures when the OpenXR instance is created.
		SetInstanceCreateInfoAndGetNextPointer(next_pointer unsafe.Pointer) int
		//Adds additional data structures when the OpenXR session is created.
		SetSessionCreateAndGetNextPointer(next_pointer unsafe.Pointer) int
		//Adds additional data structures when creating OpenXR swapchains.
		SetSwapchainCreateInfoAndGetNextPointer(next_pointer unsafe.Pointer) int
		//Adds additional data structures when each hand tracker is created.
		SetHandJointLocationsAndGetNextPointer(hand_index int, next_pointer unsafe.Pointer) int
		//Returns the number of composition layers this extension wrapper provides via [method _get_composition_layer].
		//This will only be called if the extension previously registered itself with [method OpenXRAPIExtension.register_composition_layer_provider].
		GetCompositionLayerCount() int
		//Returns a pointer to an [code]XrCompositionLayerBaseHeader[/code] struct to provide the given composition layer.
		//This will only be called if the extension previously registered itself with [method OpenXRAPIExtension.register_composition_layer_provider].
		GetCompositionLayer(index int) int
		//Returns an integer that will be used to sort the given composition layer provided via [method _get_composition_layer]. Lower numbers will move the layer to the front of the list, and higher numbers to the end. The default projection layer has an order of [code]0[/code], so layers provided by this method should probably be above or below (but not exactly) [code]0[/code].
		//This will only be called if the extension previously registered itself with [method OpenXRAPIExtension.register_composition_layer_provider].
		GetCompositionLayerOrder(index int) int
		//Returns a [PackedStringArray] of positional tracker names that are used within the extension wrapper.
		GetSuggestedTrackerNames() []string
		//Allows extensions to register additional controller metadata. This function is called even when the OpenXR API is not constructed as the metadata needs to be available to the editor.
		//Extensions should also provide metadata regardless of whether they are supported on the host system. The controller data is used to setup action maps for users who may have access to the relevant hardware.
		OnRegisterMetadata()
		//Called before the OpenXR instance is created.
		OnBeforeInstanceCreated()
		//Called right after the OpenXR instance is created.
		OnInstanceCreated(instance int)
		//Called right before the OpenXR instance is destroyed.
		OnInstanceDestroyed()
		//Called right after the OpenXR session is created.
		OnSessionCreated(session int)
		//Called as part of the OpenXR process handling. This happens right before general and physics processing steps of the main loop. During this step controller data is queried and made available to game logic.
		OnProcess()
		//Called right before the XR viewports begin their rendering step.
		OnPreRender()
		//Called right after the main swapchains are (re)created.
		OnMainSwapchainsCreated()
		//Called right before the OpenXR session is destroyed.
		OnSessionDestroyed()
		//Called when the OpenXR session state is changed to idle.
		OnStateIdle()
		//Called when the OpenXR session state is changed to ready. This means OpenXR is ready to set up the session.
		OnStateReady()
		//Called when the OpenXR session state is changed to synchronized. OpenXR also returns to this state when the application loses focus.
		OnStateSynchronized()
		//Called when the OpenXR session state is changed to visible. This means OpenXR is now ready to receive frames.
		OnStateVisible()
		//Called when the OpenXR session state is changed to focused. This state is the active state when the game runs.
		OnStateFocused()
		//Called when the OpenXR session state is changed to stopping.
		OnStateStopping()
		//Called when the OpenXR session state is changed to loss pending.
		OnStateLossPending()
		//Called when the OpenXR session state is changed to exiting.
		OnStateExiting()
		//Called when there is an OpenXR event to process. When implementing, return [code]true[/code] if the event was handled, return [code]false[/code] otherwise.
		OnEventPolled(event unsafe.Pointer) bool
		//Adds additional data structures to composition layers created by [OpenXRCompositionLayer].
		//[param property_values] contains the values of the properties returned by [method _get_viewport_composition_layer_extension_properties].
		//[param layer] is a pointer to an [code]XrCompositionLayerBaseHeader[/code] struct.
		SetViewportCompositionLayerAndGetNextPointer(layer unsafe.Pointer, property_values gd.Dictionary, next_pointer unsafe.Pointer) int
		//Gets an array of [Dictionary]s that represent properties, just like [method Object._get_property_list], that will be added to [OpenXRCompositionLayer] nodes.
		GetViewportCompositionLayerExtensionProperties() gd.Array
		//Gets a [Dictionary] containing the default values for the properties returned by [method _get_viewport_composition_layer_extension_properties].
		GetViewportCompositionLayerExtensionPropertyDefaults() gd.Dictionary
		//Called when a composition layer created via [OpenXRCompositionLayer] is destroyed.
		//[param layer] is a pointer to an [code]XrCompositionLayerBaseHeader[/code] struct.
		OnViewportCompositionLayerDestroyed(layer unsafe.Pointer)
	}
*/
type Instance [1]classdb.OpenXRExtensionWrapperExtension

/*
Returns a [Dictionary] of OpenXR extensions related to this extension. The [Dictionary] should contain the name of the extension, mapped to a [code]bool *[/code] cast to an integer:
- If the [code]bool *[/code] is a [code]nullptr[/code] this extension is mandatory.
- If the [code]bool *[/code] points to a boolean, the boolean will be updated to [code]true[/code] if the extension is enabled.
*/
func (Instance) _get_requested_extensions(impl func(ptr unsafe.Pointer) gd.Dictionary) (cb gd.ExtensionClassCallVirtualFunc) {
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
Adds additional data structures when interogating OpenXR system abilities.
*/
func (Instance) _set_system_properties_and_get_next_pointer(impl func(ptr unsafe.Pointer, next_pointer unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, next_pointer)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Adds additional data structures when the OpenXR instance is created.
*/
func (Instance) _set_instance_create_info_and_get_next_pointer(impl func(ptr unsafe.Pointer, next_pointer unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, next_pointer)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Adds additional data structures when the OpenXR session is created.
*/
func (Instance) _set_session_create_and_get_next_pointer(impl func(ptr unsafe.Pointer, next_pointer unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, next_pointer)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Adds additional data structures when creating OpenXR swapchains.
*/
func (Instance) _set_swapchain_create_info_and_get_next_pointer(impl func(ptr unsafe.Pointer, next_pointer unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, next_pointer)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Adds additional data structures when each hand tracker is created.
*/
func (Instance) _set_hand_joint_locations_and_get_next_pointer(impl func(ptr unsafe.Pointer, hand_index int, next_pointer unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var hand_index = gd.UnsafeGet[gd.Int](p_args, 0)
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(hand_index), next_pointer)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Returns the number of composition layers this extension wrapper provides via [method _get_composition_layer].
This will only be called if the extension previously registered itself with [method OpenXRAPIExtension.register_composition_layer_provider].
*/
func (Instance) _get_composition_layer_count(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Returns a pointer to an [code]XrCompositionLayerBaseHeader[/code] struct to provide the given composition layer.
This will only be called if the extension previously registered itself with [method OpenXRAPIExtension.register_composition_layer_provider].
*/
func (Instance) _get_composition_layer(impl func(ptr unsafe.Pointer, index int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Returns an integer that will be used to sort the given composition layer provided via [method _get_composition_layer]. Lower numbers will move the layer to the front of the list, and higher numbers to the end. The default projection layer has an order of [code]0[/code], so layers provided by this method should probably be above or below (but not exactly) [code]0[/code].
This will only be called if the extension previously registered itself with [method OpenXRAPIExtension.register_composition_layer_provider].
*/
func (Instance) _get_composition_layer_order(impl func(ptr unsafe.Pointer, index int) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(index))
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Returns a [PackedStringArray] of positional tracker names that are used within the extension wrapper.
*/
func (Instance) _get_suggested_tracker_names(impl func(ptr unsafe.Pointer) []string) (cb gd.ExtensionClassCallVirtualFunc) {
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
Allows extensions to register additional controller metadata. This function is called even when the OpenXR API is not constructed as the metadata needs to be available to the editor.
Extensions should also provide metadata regardless of whether they are supported on the host system. The controller data is used to setup action maps for users who may have access to the relevant hardware.
*/
func (Instance) _on_register_metadata(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called before the OpenXR instance is created.
*/
func (Instance) _on_before_instance_created(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called right after the OpenXR instance is created.
*/
func (Instance) _on_instance_created(impl func(ptr unsafe.Pointer, instance int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var instance = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(instance))
	}
}

/*
Called right before the OpenXR instance is destroyed.
*/
func (Instance) _on_instance_destroyed(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called right after the OpenXR session is created.
*/
func (Instance) _on_session_created(impl func(ptr unsafe.Pointer, session int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var session = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, int(session))
	}
}

/*
Called as part of the OpenXR process handling. This happens right before general and physics processing steps of the main loop. During this step controller data is queried and made available to game logic.
*/
func (Instance) _on_process(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called right before the XR viewports begin their rendering step.
*/
func (Instance) _on_pre_render(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called right after the main swapchains are (re)created.
*/
func (Instance) _on_main_swapchains_created(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called right before the OpenXR session is destroyed.
*/
func (Instance) _on_session_destroyed(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to idle.
*/
func (Instance) _on_state_idle(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to ready. This means OpenXR is ready to set up the session.
*/
func (Instance) _on_state_ready(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to synchronized. OpenXR also returns to this state when the application loses focus.
*/
func (Instance) _on_state_synchronized(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to visible. This means OpenXR is now ready to receive frames.
*/
func (Instance) _on_state_visible(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to focused. This state is the active state when the game runs.
*/
func (Instance) _on_state_focused(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to stopping.
*/
func (Instance) _on_state_stopping(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to loss pending.
*/
func (Instance) _on_state_loss_pending(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to exiting.
*/
func (Instance) _on_state_exiting(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when there is an OpenXR event to process. When implementing, return [code]true[/code] if the event was handled, return [code]false[/code] otherwise.
*/
func (Instance) _on_event_polled(impl func(ptr unsafe.Pointer, event unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var event = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, event)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Adds additional data structures to composition layers created by [OpenXRCompositionLayer].
[param property_values] contains the values of the properties returned by [method _get_viewport_composition_layer_extension_properties].
[param layer] is a pointer to an [code]XrCompositionLayerBaseHeader[/code] struct.
*/
func (Instance) _set_viewport_composition_layer_and_get_next_pointer(impl func(ptr unsafe.Pointer, layer unsafe.Pointer, property_values gd.Dictionary, next_pointer unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var layer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		var property_values = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args, 1))
		defer pointers.End(property_values)
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, layer, property_values, next_pointer)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Gets an array of [Dictionary]s that represent properties, just like [method Object._get_property_list], that will be added to [OpenXRCompositionLayer] nodes.
*/
func (Instance) _get_viewport_composition_layer_extension_properties(impl func(ptr unsafe.Pointer) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
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
Gets a [Dictionary] containing the default values for the properties returned by [method _get_viewport_composition_layer_extension_properties].
*/
func (Instance) _get_viewport_composition_layer_extension_property_defaults(impl func(ptr unsafe.Pointer) gd.Dictionary) (cb gd.ExtensionClassCallVirtualFunc) {
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
Called when a composition layer created via [OpenXRCompositionLayer] is destroyed.
[param layer] is a pointer to an [code]XrCompositionLayerBaseHeader[/code] struct.
*/
func (Instance) _on_viewport_composition_layer_destroyed(impl func(ptr unsafe.Pointer, layer unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var layer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, layer)
	}
}

/*
Returns the created [OpenXRAPIExtension], which can be used to access the OpenXR API.
*/
func (self Instance) GetOpenxrApi() objects.OpenXRAPIExtension {
	return objects.OpenXRAPIExtension(class(self).GetOpenxrApi())
}

/*
Registers the extension. This should happen at core module initialization level.
*/
func (self Instance) RegisterExtensionWrapper() {
	class(self).RegisterExtensionWrapper()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.OpenXRExtensionWrapperExtension

func (self class) AsObject() gd.Object    { return self[0].AsObject() }
func (self Instance) AsObject() gd.Object { return self[0].AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("OpenXRExtensionWrapperExtension"))
	return Instance{classdb.OpenXRExtensionWrapperExtension(object)}
}

/*
Returns a [Dictionary] of OpenXR extensions related to this extension. The [Dictionary] should contain the name of the extension, mapped to a [code]bool *[/code] cast to an integer:
- If the [code]bool *[/code] is a [code]nullptr[/code] this extension is mandatory.
- If the [code]bool *[/code] points to a boolean, the boolean will be updated to [code]true[/code] if the extension is enabled.
*/
func (class) _get_requested_extensions(impl func(ptr unsafe.Pointer) gd.Dictionary) (cb gd.ExtensionClassCallVirtualFunc) {
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
Adds additional data structures when interogating OpenXR system abilities.
*/
func (class) _set_system_properties_and_get_next_pointer(impl func(ptr unsafe.Pointer, next_pointer unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, next_pointer)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Adds additional data structures when the OpenXR instance is created.
*/
func (class) _set_instance_create_info_and_get_next_pointer(impl func(ptr unsafe.Pointer, next_pointer unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, next_pointer)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Adds additional data structures when the OpenXR session is created.
*/
func (class) _set_session_create_and_get_next_pointer(impl func(ptr unsafe.Pointer, next_pointer unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, next_pointer)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Adds additional data structures when creating OpenXR swapchains.
*/
func (class) _set_swapchain_create_info_and_get_next_pointer(impl func(ptr unsafe.Pointer, next_pointer unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, next_pointer)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Adds additional data structures when each hand tracker is created.
*/
func (class) _set_hand_joint_locations_and_get_next_pointer(impl func(ptr unsafe.Pointer, hand_index gd.Int, next_pointer unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var hand_index = gd.UnsafeGet[gd.Int](p_args, 0)
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, hand_index, next_pointer)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns the number of composition layers this extension wrapper provides via [method _get_composition_layer].
This will only be called if the extension previously registered itself with [method OpenXRAPIExtension.register_composition_layer_provider].
*/
func (class) _get_composition_layer_count(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns a pointer to an [code]XrCompositionLayerBaseHeader[/code] struct to provide the given composition layer.
This will only be called if the extension previously registered itself with [method OpenXRAPIExtension.register_composition_layer_provider].
*/
func (class) _get_composition_layer(impl func(ptr unsafe.Pointer, index gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns an integer that will be used to sort the given composition layer provided via [method _get_composition_layer]. Lower numbers will move the layer to the front of the list, and higher numbers to the end. The default projection layer has an order of [code]0[/code], so layers provided by this method should probably be above or below (but not exactly) [code]0[/code].
This will only be called if the extension previously registered itself with [method OpenXRAPIExtension.register_composition_layer_provider].
*/
func (class) _get_composition_layer_order(impl func(ptr unsafe.Pointer, index gd.Int) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var index = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, index)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Returns a [PackedStringArray] of positional tracker names that are used within the extension wrapper.
*/
func (class) _get_suggested_tracker_names(impl func(ptr unsafe.Pointer) gd.PackedStringArray) (cb gd.ExtensionClassCallVirtualFunc) {
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
Allows extensions to register additional controller metadata. This function is called even when the OpenXR API is not constructed as the metadata needs to be available to the editor.
Extensions should also provide metadata regardless of whether they are supported on the host system. The controller data is used to setup action maps for users who may have access to the relevant hardware.
*/
func (class) _on_register_metadata(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called before the OpenXR instance is created.
*/
func (class) _on_before_instance_created(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called right after the OpenXR instance is created.
*/
func (class) _on_instance_created(impl func(ptr unsafe.Pointer, instance gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var instance = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, instance)
	}
}

/*
Called right before the OpenXR instance is destroyed.
*/
func (class) _on_instance_destroyed(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called right after the OpenXR session is created.
*/
func (class) _on_session_created(impl func(ptr unsafe.Pointer, session gd.Int)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var session = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, session)
	}
}

/*
Called as part of the OpenXR process handling. This happens right before general and physics processing steps of the main loop. During this step controller data is queried and made available to game logic.
*/
func (class) _on_process(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called right before the XR viewports begin their rendering step.
*/
func (class) _on_pre_render(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called right after the main swapchains are (re)created.
*/
func (class) _on_main_swapchains_created(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called right before the OpenXR session is destroyed.
*/
func (class) _on_session_destroyed(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to idle.
*/
func (class) _on_state_idle(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to ready. This means OpenXR is ready to set up the session.
*/
func (class) _on_state_ready(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to synchronized. OpenXR also returns to this state when the application loses focus.
*/
func (class) _on_state_synchronized(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to visible. This means OpenXR is now ready to receive frames.
*/
func (class) _on_state_visible(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to focused. This state is the active state when the game runs.
*/
func (class) _on_state_focused(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to stopping.
*/
func (class) _on_state_stopping(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to loss pending.
*/
func (class) _on_state_loss_pending(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when the OpenXR session state is changed to exiting.
*/
func (class) _on_state_exiting(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

/*
Called when there is an OpenXR event to process. When implementing, return [code]true[/code] if the event was handled, return [code]false[/code] otherwise.
*/
func (class) _on_event_polled(impl func(ptr unsafe.Pointer, event unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var event = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, event)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Adds additional data structures to composition layers created by [OpenXRCompositionLayer].
[param property_values] contains the values of the properties returned by [method _get_viewport_composition_layer_extension_properties].
[param layer] is a pointer to an [code]XrCompositionLayerBaseHeader[/code] struct.
*/
func (class) _set_viewport_composition_layer_and_get_next_pointer(impl func(ptr unsafe.Pointer, layer unsafe.Pointer, property_values gd.Dictionary, next_pointer unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var layer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		var property_values = pointers.New[gd.Dictionary](gd.UnsafeGet[[1]uintptr](p_args, 1))
		var next_pointer = gd.UnsafeGet[unsafe.Pointer](p_args, 2)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, layer, property_values, next_pointer)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Gets an array of [Dictionary]s that represent properties, just like [method Object._get_property_list], that will be added to [OpenXRCompositionLayer] nodes.
*/
func (class) _get_viewport_composition_layer_extension_properties(impl func(ptr unsafe.Pointer) gd.Array) (cb gd.ExtensionClassCallVirtualFunc) {
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
Gets a [Dictionary] containing the default values for the properties returned by [method _get_viewport_composition_layer_extension_properties].
*/
func (class) _get_viewport_composition_layer_extension_property_defaults(impl func(ptr unsafe.Pointer) gd.Dictionary) (cb gd.ExtensionClassCallVirtualFunc) {
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
Called when a composition layer created via [OpenXRCompositionLayer] is destroyed.
[param layer] is a pointer to an [code]XrCompositionLayerBaseHeader[/code] struct.
*/
func (class) _on_viewport_composition_layer_destroyed(impl func(ptr unsafe.Pointer, layer unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var layer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, layer)
	}
}

/*
Returns the created [OpenXRAPIExtension], which can be used to access the OpenXR API.
*/
//go:nosplit
func (self class) GetOpenxrApi() objects.OpenXRAPIExtension {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRExtensionWrapperExtension.Bind_get_openxr_api, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.OpenXRAPIExtension{classdb.OpenXRAPIExtension(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}

/*
Registers the extension. This should happen at core module initialization level.
*/
//go:nosplit
func (self class) RegisterExtensionWrapper() {
	var frame = callframe.New()
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.OpenXRExtensionWrapperExtension.Bind_register_extension_wrapper, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsOpenXRExtensionWrapperExtension() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsOpenXRExtensionWrapperExtension() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_requested_extensions":
		return reflect.ValueOf(self._get_requested_extensions)
	case "_set_system_properties_and_get_next_pointer":
		return reflect.ValueOf(self._set_system_properties_and_get_next_pointer)
	case "_set_instance_create_info_and_get_next_pointer":
		return reflect.ValueOf(self._set_instance_create_info_and_get_next_pointer)
	case "_set_session_create_and_get_next_pointer":
		return reflect.ValueOf(self._set_session_create_and_get_next_pointer)
	case "_set_swapchain_create_info_and_get_next_pointer":
		return reflect.ValueOf(self._set_swapchain_create_info_and_get_next_pointer)
	case "_set_hand_joint_locations_and_get_next_pointer":
		return reflect.ValueOf(self._set_hand_joint_locations_and_get_next_pointer)
	case "_get_composition_layer_count":
		return reflect.ValueOf(self._get_composition_layer_count)
	case "_get_composition_layer":
		return reflect.ValueOf(self._get_composition_layer)
	case "_get_composition_layer_order":
		return reflect.ValueOf(self._get_composition_layer_order)
	case "_get_suggested_tracker_names":
		return reflect.ValueOf(self._get_suggested_tracker_names)
	case "_on_register_metadata":
		return reflect.ValueOf(self._on_register_metadata)
	case "_on_before_instance_created":
		return reflect.ValueOf(self._on_before_instance_created)
	case "_on_instance_created":
		return reflect.ValueOf(self._on_instance_created)
	case "_on_instance_destroyed":
		return reflect.ValueOf(self._on_instance_destroyed)
	case "_on_session_created":
		return reflect.ValueOf(self._on_session_created)
	case "_on_process":
		return reflect.ValueOf(self._on_process)
	case "_on_pre_render":
		return reflect.ValueOf(self._on_pre_render)
	case "_on_main_swapchains_created":
		return reflect.ValueOf(self._on_main_swapchains_created)
	case "_on_session_destroyed":
		return reflect.ValueOf(self._on_session_destroyed)
	case "_on_state_idle":
		return reflect.ValueOf(self._on_state_idle)
	case "_on_state_ready":
		return reflect.ValueOf(self._on_state_ready)
	case "_on_state_synchronized":
		return reflect.ValueOf(self._on_state_synchronized)
	case "_on_state_visible":
		return reflect.ValueOf(self._on_state_visible)
	case "_on_state_focused":
		return reflect.ValueOf(self._on_state_focused)
	case "_on_state_stopping":
		return reflect.ValueOf(self._on_state_stopping)
	case "_on_state_loss_pending":
		return reflect.ValueOf(self._on_state_loss_pending)
	case "_on_state_exiting":
		return reflect.ValueOf(self._on_state_exiting)
	case "_on_event_polled":
		return reflect.ValueOf(self._on_event_polled)
	case "_set_viewport_composition_layer_and_get_next_pointer":
		return reflect.ValueOf(self._set_viewport_composition_layer_and_get_next_pointer)
	case "_get_viewport_composition_layer_extension_properties":
		return reflect.ValueOf(self._get_viewport_composition_layer_extension_properties)
	case "_get_viewport_composition_layer_extension_property_defaults":
		return reflect.ValueOf(self._get_viewport_composition_layer_extension_property_defaults)
	case "_on_viewport_composition_layer_destroyed":
		return reflect.ValueOf(self._on_viewport_composition_layer_destroyed)
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_requested_extensions":
		return reflect.ValueOf(self._get_requested_extensions)
	case "_set_system_properties_and_get_next_pointer":
		return reflect.ValueOf(self._set_system_properties_and_get_next_pointer)
	case "_set_instance_create_info_and_get_next_pointer":
		return reflect.ValueOf(self._set_instance_create_info_and_get_next_pointer)
	case "_set_session_create_and_get_next_pointer":
		return reflect.ValueOf(self._set_session_create_and_get_next_pointer)
	case "_set_swapchain_create_info_and_get_next_pointer":
		return reflect.ValueOf(self._set_swapchain_create_info_and_get_next_pointer)
	case "_set_hand_joint_locations_and_get_next_pointer":
		return reflect.ValueOf(self._set_hand_joint_locations_and_get_next_pointer)
	case "_get_composition_layer_count":
		return reflect.ValueOf(self._get_composition_layer_count)
	case "_get_composition_layer":
		return reflect.ValueOf(self._get_composition_layer)
	case "_get_composition_layer_order":
		return reflect.ValueOf(self._get_composition_layer_order)
	case "_get_suggested_tracker_names":
		return reflect.ValueOf(self._get_suggested_tracker_names)
	case "_on_register_metadata":
		return reflect.ValueOf(self._on_register_metadata)
	case "_on_before_instance_created":
		return reflect.ValueOf(self._on_before_instance_created)
	case "_on_instance_created":
		return reflect.ValueOf(self._on_instance_created)
	case "_on_instance_destroyed":
		return reflect.ValueOf(self._on_instance_destroyed)
	case "_on_session_created":
		return reflect.ValueOf(self._on_session_created)
	case "_on_process":
		return reflect.ValueOf(self._on_process)
	case "_on_pre_render":
		return reflect.ValueOf(self._on_pre_render)
	case "_on_main_swapchains_created":
		return reflect.ValueOf(self._on_main_swapchains_created)
	case "_on_session_destroyed":
		return reflect.ValueOf(self._on_session_destroyed)
	case "_on_state_idle":
		return reflect.ValueOf(self._on_state_idle)
	case "_on_state_ready":
		return reflect.ValueOf(self._on_state_ready)
	case "_on_state_synchronized":
		return reflect.ValueOf(self._on_state_synchronized)
	case "_on_state_visible":
		return reflect.ValueOf(self._on_state_visible)
	case "_on_state_focused":
		return reflect.ValueOf(self._on_state_focused)
	case "_on_state_stopping":
		return reflect.ValueOf(self._on_state_stopping)
	case "_on_state_loss_pending":
		return reflect.ValueOf(self._on_state_loss_pending)
	case "_on_state_exiting":
		return reflect.ValueOf(self._on_state_exiting)
	case "_on_event_polled":
		return reflect.ValueOf(self._on_event_polled)
	case "_set_viewport_composition_layer_and_get_next_pointer":
		return reflect.ValueOf(self._set_viewport_composition_layer_and_get_next_pointer)
	case "_get_viewport_composition_layer_extension_properties":
		return reflect.ValueOf(self._get_viewport_composition_layer_extension_properties)
	case "_get_viewport_composition_layer_extension_property_defaults":
		return reflect.ValueOf(self._get_viewport_composition_layer_extension_property_defaults)
	case "_on_viewport_composition_layer_destroyed":
		return reflect.ValueOf(self._on_viewport_composition_layer_destroyed)
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {
	classdb.Register("OpenXRExtensionWrapperExtension", func(ptr gd.Object) any { return classdb.OpenXRExtensionWrapperExtension(ptr) })
}
