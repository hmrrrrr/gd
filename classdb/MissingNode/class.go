// Package MissingNode provides methods for working with MissingNode object instances.
package MissingNode

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
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
This is an internal editor class intended for keeping data of nodes of unknown type (most likely this type was supplied by an extension that is no longer loaded). It can't be manually instantiated or placed in a scene.
[b]Warning:[/b] Ignore missing nodes unless you know what you are doing. Existing properties on a missing node can be freely modified in code, regardless of the type they are intended to be.
*/
type Instance [1]gdclass.MissingNode

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsMissingNode() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.MissingNode

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MissingNode"))
	casted := Instance{*(*gdclass.MissingNode)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) OriginalClass() string {
	return string(class(self).GetOriginalClass().String())
}

func (self Instance) SetOriginalClass(value string) {
	class(self).SetOriginalClass(String.New(value))
}

func (self Instance) OriginalScene() string {
	return string(class(self).GetOriginalScene().String())
}

func (self Instance) SetOriginalScene(value string) {
	class(self).SetOriginalScene(String.New(value))
}

func (self Instance) RecordingProperties() bool {
	return bool(class(self).IsRecordingProperties())
}

func (self Instance) SetRecordingProperties(value bool) {
	class(self).SetRecordingProperties(value)
}

//go:nosplit
func (self class) SetOriginalClass(name String.Readable) { //gd:MissingNode.set_original_class
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MissingNode.Bind_set_original_class, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOriginalClass() String.Readable { //gd:MissingNode.get_original_class
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MissingNode.Bind_get_original_class, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOriginalScene(name String.Readable) { //gd:MissingNode.set_original_scene
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MissingNode.Bind_set_original_scene, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOriginalScene() String.Readable { //gd:MissingNode.get_original_scene
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MissingNode.Bind_get_original_scene, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetRecordingProperties(enable bool) { //gd:MissingNode.set_recording_properties
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MissingNode.Bind_set_recording_properties, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsRecordingProperties() bool { //gd:MissingNode.is_recording_properties
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MissingNode.Bind_is_recording_properties, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsMissingNode() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMissingNode() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced      { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance   { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node.Advanced(self.AsNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node.Instance(self.AsNode()), name)
	}
}
func init() {
	gdclass.Register("MissingNode", func(ptr gd.Object) any { return [1]gdclass.MissingNode{*(*gdclass.MissingNode)(unsafe.Pointer(&ptr))} })
}
