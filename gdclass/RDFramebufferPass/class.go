package RDFramebufferPass

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
This class contains the list of attachment descriptions for a framebuffer pass. Each points with an index to a previously supplied list of texture attachments.
Multipass framebuffers can optimize some configurations in mobile. On desktop, they provide little to no advantage.
This object is used by [RenderingDevice].

*/
type Go [1]classdb.RDFramebufferPass
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.RDFramebufferPass
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RDFramebufferPass"))
	return Go{classdb.RDFramebufferPass(object)}
}

func (self Go) ColorAttachments() []int32 {
		return []int32(class(self).GetColorAttachments().AsSlice())
}

func (self Go) SetColorAttachments(value []int32) {
	class(self).SetColorAttachments(gd.NewPackedInt32Slice(value))
}

func (self Go) InputAttachments() []int32 {
		return []int32(class(self).GetInputAttachments().AsSlice())
}

func (self Go) SetInputAttachments(value []int32) {
	class(self).SetInputAttachments(gd.NewPackedInt32Slice(value))
}

func (self Go) ResolveAttachments() []int32 {
		return []int32(class(self).GetResolveAttachments().AsSlice())
}

func (self Go) SetResolveAttachments(value []int32) {
	class(self).SetResolveAttachments(gd.NewPackedInt32Slice(value))
}

func (self Go) PreserveAttachments() []int32 {
		return []int32(class(self).GetPreserveAttachments().AsSlice())
}

func (self Go) SetPreserveAttachments(value []int32) {
	class(self).SetPreserveAttachments(gd.NewPackedInt32Slice(value))
}

func (self Go) DepthAttachment() int {
		return int(int(class(self).GetDepthAttachment()))
}

func (self Go) SetDepthAttachment(value int) {
	class(self).SetDepthAttachment(gd.Int(value))
}

//go:nosplit
func (self class) SetColorAttachments(p_member gd.PackedInt32Array)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(p_member))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDFramebufferPass.Bind_set_color_attachments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetColorAttachments() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDFramebufferPass.Bind_get_color_attachments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetInputAttachments(p_member gd.PackedInt32Array)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(p_member))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDFramebufferPass.Bind_set_input_attachments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInputAttachments() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDFramebufferPass.Bind_get_input_attachments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetResolveAttachments(p_member gd.PackedInt32Array)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(p_member))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDFramebufferPass.Bind_set_resolve_attachments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetResolveAttachments() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDFramebufferPass.Bind_get_resolve_attachments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetPreserveAttachments(p_member gd.PackedInt32Array)  {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(p_member))
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDFramebufferPass.Bind_set_preserve_attachments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetPreserveAttachments() gd.PackedInt32Array {
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDFramebufferPass.Bind_get_preserve_attachments, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = discreet.New[gd.PackedInt32Array](r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetDepthAttachment(p_member gd.Int)  {
	var frame = callframe.New()
	callframe.Arg(frame, p_member)
	var r_ret callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDFramebufferPass.Bind_set_depth_attachment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetDepthAttachment() gd.Int {
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RDFramebufferPass.Bind_get_depth_attachment, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRDFramebufferPass() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsRDFramebufferPass() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("RDFramebufferPass", func(ptr gd.Object) any { return classdb.RDFramebufferPass(ptr) })}
