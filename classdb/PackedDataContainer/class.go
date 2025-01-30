// Package PackedDataContainer provides methods for working with PackedDataContainer object instances.
package PackedDataContainer

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Resource"
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
[PackedDataContainer] can be used to efficiently store data from untyped containers. The data is packed into raw bytes and can be saved to file. Only [Array] and [Dictionary] can be stored this way.
You can retrieve the data by iterating on the container, which will work as if iterating on the packed data itself. If the packed container is a [Dictionary], the data can be retrieved by key names ([String]/[StringName] only).
[codeblock]
var data = { "key": "value", "another_key": 123, "lock": Vector2() }
var packed = PackedDataContainer.new()
packed.pack(data)
ResourceSaver.save(packed, "packed_data.res")
[/codeblock]
[codeblock]
var container = load("packed_data.res")
for key in container:

	prints(key, container[key])

# Prints:
# key value
# lock (0, 0)
# another_key 123
[/codeblock]
Nested containers will be packed recursively. While iterating, they will be returned as [PackedDataContainerRef].
*/
type Instance [1]gdclass.PackedDataContainer

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsPackedDataContainer() Instance
}

/*
Packs the given container into a binary representation. The [param value] must be either [Array] or [Dictionary], any other type will result in invalid data error.
[b]Note:[/b] Subsequent calls to this method will overwrite the existing data.
*/
func (self Instance) Pack(value any) error { //gd:PackedDataContainer.pack
	return error(gd.ToError(class(self).Pack(variant.New(value))))
}

/*
Returns the size of the packed container (see [method Array.size] and [method Dictionary.size]).
*/
func (self Instance) Size() int { //gd:PackedDataContainer.size
	return int(int(class(self).Size()))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.PackedDataContainer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("PackedDataContainer"))
	casted := Instance{*(*gdclass.PackedDataContainer)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Packs the given container into a binary representation. The [param value] must be either [Array] or [Dictionary], any other type will result in invalid data error.
[b]Note:[/b] Subsequent calls to this method will overwrite the existing data.
*/
//go:nosplit
func (self class) Pack(value variant.Any) Error.Code { //gd:PackedDataContainer.pack
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(value)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PackedDataContainer.Bind_pack, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the size of the packed container (see [method Array.size] and [method Dictionary.size]).
*/
//go:nosplit
func (self class) Size() int64 { //gd:PackedDataContainer.size
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.PackedDataContainer.Bind_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsPackedDataContainer() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsPackedDataContainer() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
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
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("PackedDataContainer", func(ptr gd.Object) any {
		return [1]gdclass.PackedDataContainer{*(*gdclass.PackedDataContainer)(unsafe.Pointer(&ptr))}
	})
}
