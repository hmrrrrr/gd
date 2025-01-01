package RenderDataExtension

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/RenderData"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
This class allows for a RenderData implementation to be made in GDExtension.

	// RenderDataExtension methods that can be overridden by a [Class] that extends it.
	type RenderDataExtension interface {
		//Implement this in GDExtension to return the implementation's [RenderSceneBuffers] object.
		GetRenderSceneBuffers() objects.RenderSceneBuffers
		//Implement this in GDExtension to return the implementation's [RenderSceneDataExtension] object.
		GetRenderSceneData() objects.RenderSceneData
		//Implement this in GDExtension to return the [RID] of the implementation's environment object.
		GetEnvironment() Resource.ID
		//Implement this in GDExtension to return the [RID] for the implementation's camera attributes object.
		GetCameraAttributes() Resource.ID
	}
*/
type Instance [1]classdb.RenderDataExtension
type Any interface {
	gd.IsClass
	AsRenderDataExtension() Instance
}

/*
Implement this in GDExtension to return the implementation's [RenderSceneBuffers] object.
*/
func (Instance) _get_render_scene_buffers(impl func(ptr unsafe.Pointer) objects.RenderSceneBuffers) (cb gd.ExtensionClassCallVirtualFunc) {
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
Implement this in GDExtension to return the implementation's [RenderSceneDataExtension] object.
*/
func (Instance) _get_render_scene_data(impl func(ptr unsafe.Pointer) objects.RenderSceneData) (cb gd.ExtensionClassCallVirtualFunc) {
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
Implement this in GDExtension to return the [RID] of the implementation's environment object.
*/
func (Instance) _get_environment(impl func(ptr unsafe.Pointer) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement this in GDExtension to return the [RID] for the implementation's camera attributes object.
*/
func (Instance) _get_camera_attributes(impl func(ptr unsafe.Pointer) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.RenderDataExtension

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RenderDataExtension"))
	return Instance{classdb.RenderDataExtension(object)}
}

/*
Implement this in GDExtension to return the implementation's [RenderSceneBuffers] object.
*/
func (class) _get_render_scene_buffers(impl func(ptr unsafe.Pointer) objects.RenderSceneBuffers) (cb gd.ExtensionClassCallVirtualFunc) {
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
Implement this in GDExtension to return the implementation's [RenderSceneDataExtension] object.
*/
func (class) _get_render_scene_data(impl func(ptr unsafe.Pointer) objects.RenderSceneData) (cb gd.ExtensionClassCallVirtualFunc) {
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
Implement this in GDExtension to return the [RID] of the implementation's environment object.
*/
func (class) _get_environment(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement this in GDExtension to return the [RID] for the implementation's camera attributes object.
*/
func (class) _get_camera_attributes(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (self class) AsRenderDataExtension() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRenderDataExtension() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsRenderData() RenderData.Advanced {
	return *((*RenderData.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsRenderData() RenderData.Instance {
	return *((*RenderData.Instance)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_render_scene_buffers":
		return reflect.ValueOf(self._get_render_scene_buffers)
	case "_get_render_scene_data":
		return reflect.ValueOf(self._get_render_scene_data)
	case "_get_environment":
		return reflect.ValueOf(self._get_environment)
	case "_get_camera_attributes":
		return reflect.ValueOf(self._get_camera_attributes)
	default:
		return gd.VirtualByName(self.AsRenderData(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_render_scene_buffers":
		return reflect.ValueOf(self._get_render_scene_buffers)
	case "_get_render_scene_data":
		return reflect.ValueOf(self._get_render_scene_data)
	case "_get_environment":
		return reflect.ValueOf(self._get_environment)
	case "_get_camera_attributes":
		return reflect.ValueOf(self._get_camera_attributes)
	default:
		return gd.VirtualByName(self.AsRenderData(), name)
	}
}
func init() {
	classdb.Register("RenderDataExtension", func(ptr gd.Object) any { return [1]classdb.RenderDataExtension{classdb.RenderDataExtension(ptr)} })
}
