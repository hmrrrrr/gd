// Package RDPipelineColorBlendState provides methods for working with RDPipelineColorBlendState object instances.
package RDPipelineColorBlendState

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
import "graphics.gd/variant/RID"
import "graphics.gd/variant/Color"

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

/*
This object is used by [RenderingDevice].
*/
type Instance [1]gdclass.RDPipelineColorBlendState

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRDPipelineColorBlendState() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RDPipelineColorBlendState

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDPipelineColorBlendState"))
	casted := Instance{*(*gdclass.RDPipelineColorBlendState)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) EnableLogicOp() bool {
	return bool(class(self).GetEnableLogicOp())
}

func (self Instance) SetEnableLogicOp(value bool) {
	class(self).SetEnableLogicOp(value)
}

func (self Instance) LogicOp() gdclass.RenderingDeviceLogicOperation {
	return gdclass.RenderingDeviceLogicOperation(class(self).GetLogicOp())
}

func (self Instance) SetLogicOp(value gdclass.RenderingDeviceLogicOperation) {
	class(self).SetLogicOp(value)
}

func (self Instance) BlendConstant() Color.RGBA {
	return Color.RGBA(class(self).GetBlendConstant())
}

func (self Instance) SetBlendConstant(value Color.RGBA) {
	class(self).SetBlendConstant(gd.Color(value))
}

func (self Instance) Attachments() [][1]gdclass.RDPipelineColorBlendStateAttachment {
	return [][1]gdclass.RDPipelineColorBlendStateAttachment(gd.ArrayAs[[][1]gdclass.RDPipelineColorBlendStateAttachment](gd.InternalArray(class(self).GetAttachments())))
}

func (self Instance) SetAttachments(value [][1]gdclass.RDPipelineColorBlendStateAttachment) {
	class(self).SetAttachments(gd.ArrayFromSlice[Array.Contains[[1]gdclass.RDPipelineColorBlendStateAttachment]](value))
}

//go:nosplit
func (self class) SetEnableLogicOp(p_member bool) { //gd:RDPipelineColorBlendState.set_enable_logic_op
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendState.Bind_set_enable_logic_op, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetEnableLogicOp() bool { //gd:RDPipelineColorBlendState.get_enable_logic_op
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendState.Bind_get_enable_logic_op, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLogicOp(p_member gdclass.RenderingDeviceLogicOperation) { //gd:RDPipelineColorBlendState.set_logic_op
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendState.Bind_set_logic_op, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetLogicOp() gdclass.RenderingDeviceLogicOperation { //gd:RDPipelineColorBlendState.get_logic_op
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.RenderingDeviceLogicOperation](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendState.Bind_get_logic_op, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBlendConstant(p_member gd.Color) { //gd:RDPipelineColorBlendState.set_blend_constant
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendState.Bind_set_blend_constant, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBlendConstant() gd.Color { //gd:RDPipelineColorBlendState.get_blend_constant
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Color](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendState.Bind_get_blend_constant, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAttachments(attachments Array.Contains[[1]gdclass.RDPipelineColorBlendStateAttachment]) { //gd:RDPipelineColorBlendState.set_attachments
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(attachments)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendState.Bind_set_attachments, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAttachments() Array.Contains[[1]gdclass.RDPipelineColorBlendStateAttachment] { //gd:RDPipelineColorBlendState.get_attachments
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDPipelineColorBlendState.Bind_get_attachments, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.RDPipelineColorBlendStateAttachment]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}
func (self class) AsRDPipelineColorBlendState() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsRDPipelineColorBlendState() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("RDPipelineColorBlendState", func(ptr gd.Object) any {
		return [1]gdclass.RDPipelineColorBlendState{*(*gdclass.RDPipelineColorBlendState)(unsafe.Pointer(&ptr))}
	})
}
