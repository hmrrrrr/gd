package RenderSceneDataExtension

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/RenderSceneData"
import "graphics.gd/variant/Transform3D"
import "graphics.gd/variant/Projection"
import "graphics.gd/variant/Vector3"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Root

/*
This class allows for a RenderSceneData implementation to be made in GDExtension.

	// RenderSceneDataExtension methods that can be overridden by a [Class] that extends it.
	type RenderSceneDataExtension interface {
		//Implement this in GDExtension to return the camera [Transform3D].
		GetCamTransform() Transform3D.BasisOrigin
		//Implement this in GDExtension to return the camera [Projection].
		GetCamProjection() Projection.XYZW
		//Implement this in GDExtension to return the view count.
		GetViewCount() int
		//Implement this in GDExtension to return the eye offset for the given [param view].
		GetViewEyeOffset(view int) Vector3.XYZ
		//Implement this in GDExtension to return the view [Projection] for the given [param view].
		GetViewProjection(view int) Projection.XYZW
		//Implement this in GDExtension to return the [RID] of the uniform buffer containing the scene data as a UBO.
		GetUniformBuffer() Resource.ID
	}
*/
type Instance [1]classdb.RenderSceneDataExtension
type Any interface {
	gd.IsClass
	AsRenderSceneDataExtension() Instance
}

/*
Implement this in GDExtension to return the camera [Transform3D].
*/
func (Instance) _get_cam_transform(impl func(ptr unsafe.Pointer) Transform3D.BasisOrigin) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Transform3D(ret))
	}
}

/*
Implement this in GDExtension to return the camera [Projection].
*/
func (Instance) _get_cam_projection(impl func(ptr unsafe.Pointer) Projection.XYZW) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement this in GDExtension to return the view count.
*/
func (Instance) _get_view_count(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
	}
}

/*
Implement this in GDExtension to return the eye offset for the given [param view].
*/
func (Instance) _get_view_eye_offset(impl func(ptr unsafe.Pointer, view int) Vector3.XYZ) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var view = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(view))
		gd.UnsafeSet(p_back, gd.Vector3(ret))
	}
}

/*
Implement this in GDExtension to return the view [Projection] for the given [param view].
*/
func (Instance) _get_view_projection(impl func(ptr unsafe.Pointer, view int) Projection.XYZW) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var view = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, int(view))
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement this in GDExtension to return the [RID] of the uniform buffer containing the scene data as a UBO.
*/
func (Instance) _get_uniform_buffer(impl func(ptr unsafe.Pointer) Resource.ID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.RenderSceneDataExtension

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RenderSceneDataExtension"))
	return Instance{classdb.RenderSceneDataExtension(object)}
}

/*
Implement this in GDExtension to return the camera [Transform3D].
*/
func (class) _get_cam_transform(impl func(ptr unsafe.Pointer) gd.Transform3D) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement this in GDExtension to return the camera [Projection].
*/
func (class) _get_cam_projection(impl func(ptr unsafe.Pointer) gd.Projection) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement this in GDExtension to return the view count.
*/
func (class) _get_view_count(impl func(ptr unsafe.Pointer) gd.Int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement this in GDExtension to return the eye offset for the given [param view].
*/
func (class) _get_view_eye_offset(impl func(ptr unsafe.Pointer, view gd.Int) gd.Vector3) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var view = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, view)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement this in GDExtension to return the view [Projection] for the given [param view].
*/
func (class) _get_view_projection(impl func(ptr unsafe.Pointer, view gd.Int) gd.Projection) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		var view = gd.UnsafeGet[gd.Int](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, view)
		gd.UnsafeSet(p_back, ret)
	}
}

/*
Implement this in GDExtension to return the [RID] of the uniform buffer containing the scene data as a UBO.
*/
func (class) _get_uniform_buffer(impl func(ptr unsafe.Pointer) gd.RID) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (self class) AsRenderSceneDataExtension() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRenderSceneDataExtension() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRenderSceneData() RenderSceneData.Advanced {
	return *((*RenderSceneData.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsRenderSceneData() RenderSceneData.Instance {
	return *((*RenderSceneData.Instance)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_cam_transform":
		return reflect.ValueOf(self._get_cam_transform)
	case "_get_cam_projection":
		return reflect.ValueOf(self._get_cam_projection)
	case "_get_view_count":
		return reflect.ValueOf(self._get_view_count)
	case "_get_view_eye_offset":
		return reflect.ValueOf(self._get_view_eye_offset)
	case "_get_view_projection":
		return reflect.ValueOf(self._get_view_projection)
	case "_get_uniform_buffer":
		return reflect.ValueOf(self._get_uniform_buffer)
	default:
		return gd.VirtualByName(self.AsRenderSceneData(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_cam_transform":
		return reflect.ValueOf(self._get_cam_transform)
	case "_get_cam_projection":
		return reflect.ValueOf(self._get_cam_projection)
	case "_get_view_count":
		return reflect.ValueOf(self._get_view_count)
	case "_get_view_eye_offset":
		return reflect.ValueOf(self._get_view_eye_offset)
	case "_get_view_projection":
		return reflect.ValueOf(self._get_view_projection)
	case "_get_uniform_buffer":
		return reflect.ValueOf(self._get_uniform_buffer)
	default:
		return gd.VirtualByName(self.AsRenderSceneData(), name)
	}
}
func init() {
	classdb.Register("RenderSceneDataExtension", func(ptr gd.Object) any {
		return [1]classdb.RenderSceneDataExtension{classdb.RenderSceneDataExtension(ptr)}
	})
}
