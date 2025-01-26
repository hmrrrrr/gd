// Package RenderSceneBuffersExtension provides methods for working with RenderSceneBuffersExtension object instances.
package RenderSceneBuffersExtension

import "unsafe"
import "reflect"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/classdb/RenderSceneBuffers"
import "graphics.gd/variant/Float"

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

/*
This class allows for a RenderSceneBuffer implementation to be made in GDExtension.

	See [Interface] for methods that can be overridden by a [Class] that extends it.

%!(EXTRA string=RenderSceneBuffersExtension)
*/
type Instance [1]gdclass.RenderSceneBuffersExtension

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRenderSceneBuffersExtension() Instance
}
type Interface interface {
	//Implement this in GDExtension to handle the (re)sizing of a viewport.
	Configure(config [1]gdclass.RenderSceneBuffersConfiguration)
	//Implement this in GDExtension to record a new FSR sharpness value.
	SetFsrSharpness(fsr_sharpness Float.X)
	//Implement this in GDExtension to change the texture mipmap bias.
	SetTextureMipmapBias(texture_mipmap_bias Float.X)
	//Implement this in GDExtension to react to the debanding flag changing.
	SetUseDebanding(use_debanding bool)
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) Configure(config [1]gdclass.RenderSceneBuffersConfiguration) { return }
func (self implementation) SetFsrSharpness(fsr_sharpness Float.X)                       { return }
func (self implementation) SetTextureMipmapBias(texture_mipmap_bias Float.X)            { return }
func (self implementation) SetUseDebanding(use_debanding bool)                          { return }

/*
Implement this in GDExtension to handle the (re)sizing of a viewport.
*/
func (Instance) _configure(impl func(ptr unsafe.Pointer, config [1]gdclass.RenderSceneBuffersConfiguration)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var config = [1]gdclass.RenderSceneBuffersConfiguration{pointers.New[gdclass.RenderSceneBuffersConfiguration]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(config[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, config)
	}
}

/*
Implement this in GDExtension to record a new FSR sharpness value.
*/
func (Instance) _set_fsr_sharpness(impl func(ptr unsafe.Pointer, fsr_sharpness Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var fsr_sharpness = gd.UnsafeGet[gd.Float](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, Float.X(fsr_sharpness))
	}
}

/*
Implement this in GDExtension to change the texture mipmap bias.
*/
func (Instance) _set_texture_mipmap_bias(impl func(ptr unsafe.Pointer, texture_mipmap_bias Float.X)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var texture_mipmap_bias = gd.UnsafeGet[gd.Float](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, Float.X(texture_mipmap_bias))
	}
}

/*
Implement this in GDExtension to react to the debanding flag changing.
*/
func (Instance) _set_use_debanding(impl func(ptr unsafe.Pointer, use_debanding bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var use_debanding = gd.UnsafeGet[bool](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, use_debanding)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RenderSceneBuffersExtension

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RenderSceneBuffersExtension"))
	casted := Instance{*(*gdclass.RenderSceneBuffersExtension)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Implement this in GDExtension to handle the (re)sizing of a viewport.
*/
func (class) _configure(impl func(ptr unsafe.Pointer, config [1]gdclass.RenderSceneBuffersConfiguration)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var config = [1]gdclass.RenderSceneBuffersConfiguration{pointers.New[gdclass.RenderSceneBuffersConfiguration]([3]uint64{uint64(gd.UnsafeGet[gd.EnginePointer](p_args, 0))})}

		defer pointers.End(config[0])
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, config)
	}
}

/*
Implement this in GDExtension to record a new FSR sharpness value.
*/
func (class) _set_fsr_sharpness(impl func(ptr unsafe.Pointer, fsr_sharpness gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var fsr_sharpness = gd.UnsafeGet[gd.Float](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, fsr_sharpness)
	}
}

/*
Implement this in GDExtension to change the texture mipmap bias.
*/
func (class) _set_texture_mipmap_bias(impl func(ptr unsafe.Pointer, texture_mipmap_bias gd.Float)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var texture_mipmap_bias = gd.UnsafeGet[gd.Float](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, texture_mipmap_bias)
	}
}

/*
Implement this in GDExtension to react to the debanding flag changing.
*/
func (class) _set_use_debanding(impl func(ptr unsafe.Pointer, use_debanding bool)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var use_debanding = gd.UnsafeGet[bool](p_args, 0)

		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, use_debanding)
	}
}

func (self class) AsRenderSceneBuffersExtension() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsRenderSceneBuffersExtension() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRenderSceneBuffers() RenderSceneBuffers.Advanced {
	return *((*RenderSceneBuffers.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsRenderSceneBuffers() RenderSceneBuffers.Instance {
	return *((*RenderSceneBuffers.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_configure":
		return reflect.ValueOf(self._configure)
	case "_set_fsr_sharpness":
		return reflect.ValueOf(self._set_fsr_sharpness)
	case "_set_texture_mipmap_bias":
		return reflect.ValueOf(self._set_texture_mipmap_bias)
	case "_set_use_debanding":
		return reflect.ValueOf(self._set_use_debanding)
	default:
		return gd.VirtualByName(RenderSceneBuffers.Advanced(self.AsRenderSceneBuffers()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_configure":
		return reflect.ValueOf(self._configure)
	case "_set_fsr_sharpness":
		return reflect.ValueOf(self._set_fsr_sharpness)
	case "_set_texture_mipmap_bias":
		return reflect.ValueOf(self._set_texture_mipmap_bias)
	case "_set_use_debanding":
		return reflect.ValueOf(self._set_use_debanding)
	default:
		return gd.VirtualByName(RenderSceneBuffers.Instance(self.AsRenderSceneBuffers()), name)
	}
}
func init() {
	gdclass.Register("RenderSceneBuffersExtension", func(ptr gd.Object) any {
		return [1]gdclass.RenderSceneBuffersExtension{*(*gdclass.RenderSceneBuffersExtension)(unsafe.Pointer(&ptr))}
	})
}
