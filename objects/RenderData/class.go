package RenderData

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/objects"
import classdb "graphics.gd/internal/classdb"
import "graphics.gd/objects/Resource"

var _ unsafe.Pointer
var _ objects.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle

/*
Abstract render data object, exists for the duration of rendering a single viewport.
[b]Note:[/b] This is an internal rendering server object, do not instantiate this from script.
*/
type Instance [1]classdb.RenderData
type Any interface {
	gd.IsClass
	AsRenderData() Instance
}

/*
Returns the [RenderSceneBuffers] object managing the scene buffers for rendering this viewport.
*/
func (self Instance) GetRenderSceneBuffers() objects.RenderSceneBuffers {
	return objects.RenderSceneBuffers(class(self).GetRenderSceneBuffers())
}

/*
Returns the [RenderSceneData] object managing this frames scene data.
*/
func (self Instance) GetRenderSceneData() objects.RenderSceneData {
	return objects.RenderSceneData(class(self).GetRenderSceneData())
}

/*
Returns the [RID] of the environments object in the [RenderingServer] being used to render this viewport.
*/
func (self Instance) GetEnvironment() Resource.ID {
	return Resource.ID(class(self).GetEnvironment())
}

/*
Returns the [RID] of the camera attributes object in the [RenderingServer] being used to render this viewport.
*/
func (self Instance) GetCameraAttributes() Resource.ID {
	return Resource.ID(class(self).GetCameraAttributes())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]classdb.RenderData

func (self class) AsObject() gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() gd.Object         { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RenderData"))
	return Instance{*(*classdb.RenderData)(unsafe.Pointer(&object))}
}

/*
Returns the [RenderSceneBuffers] object managing the scene buffers for rendering this viewport.
*/
//go:nosplit
func (self class) GetRenderSceneBuffers() objects.RenderSceneBuffers {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderData.Bind_get_render_scene_buffers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.RenderSceneBuffers{gd.PointerWithOwnershipTransferredToGo[classdb.RenderSceneBuffers](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the [RenderSceneData] object managing this frames scene data.
*/
//go:nosplit
func (self class) GetRenderSceneData() objects.RenderSceneData {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderData.Bind_get_render_scene_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = objects.RenderSceneData{gd.PointerWithOwnershipTransferredToGo[classdb.RenderSceneData](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns the [RID] of the environments object in the [RenderingServer] being used to render this viewport.
*/
//go:nosplit
func (self class) GetEnvironment() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderData.Bind_get_environment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [RID] of the camera attributes object in the [RenderingServer] being used to render this viewport.
*/
//go:nosplit
func (self class) GetCameraAttributes() gd.RID {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.RID](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RenderData.Bind_get_camera_attributes, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRenderData() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRenderData() Instance { return *((*Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {
	classdb.Register("RenderData", func(ptr gd.Object) any { return [1]classdb.RenderData{*(*classdb.RenderData)(unsafe.Pointer(&ptr))} })
}
