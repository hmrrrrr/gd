// Package Expression provides methods for working with Expression object instances.
package Expression

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
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
An expression can be made of any arithmetic operation, built-in math function call, method call of a passed instance, or built-in type construction call.
An example expression text using the built-in math functions could be [code]sqrt(pow(3, 2) + pow(4, 2))[/code].
In the following example we use a [LineEdit] node to write our expression and show the result.
[codeblocks]
[gdscript]
var expression = Expression.new()

func _ready():

	$LineEdit.text_submitted.connect(self._on_text_submitted)

func _on_text_submitted(command):

	var error = expression.parse(command)
	if error != OK:
	    print(expression.get_error_text())
	    return
	var result = expression.execute()
	if not expression.has_execute_failed():
	    $LineEdit.text = str(result)

[/gdscript]
[csharp]
private Expression _expression = new Expression();

public override void _Ready()

	{
	    GetNode<LineEdit>("LineEdit").TextSubmitted += OnTextEntered;
	}

private void OnTextEntered(string command)

	{
	    Error error = _expression.Parse(command);
	    if (error != Error.Ok)
	    {
	        GD.Print(_expression.GetErrorText());
	        return;
	    }
	    Variant result = _expression.Execute();
	    if (!_expression.HasExecuteFailed())
	    {
	        GetNode<LineEdit>("LineEdit").Text = result.ToString();
	    }
	}

[/csharp]
[/codeblocks]
*/
type Instance [1]gdclass.Expression

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsExpression() Instance
}

/*
Parses the expression and returns an [enum Error] code.
You can optionally specify names of variables that may appear in the expression with [param input_names], so that you can bind them when it gets executed.
*/
func (self Instance) Parse(expression string) error { //gd:Expression.parse
	return error(gd.ToError(class(self).Parse(String.New(expression), Packed.MakeStrings([1][]string{}[0]...))))
}

/*
Executes the expression that was previously parsed by [method parse] and returns the result. Before you use the returned object, you should check if the method failed by calling [method has_execute_failed].
If you defined input variables in [method parse], you can specify their values in the inputs array, in the same order.
*/
func (self Instance) Execute() any { //gd:Expression.execute
	return any(class(self).Execute(Array.Nil, [1]Object.Instance{}[0], true, false).Interface())
}

/*
Returns [code]true[/code] if [method execute] has failed.
*/
func (self Instance) HasExecuteFailed() bool { //gd:Expression.has_execute_failed
	return bool(class(self).HasExecuteFailed())
}

/*
Returns the error text if [method parse] or [method execute] has failed.
*/
func (self Instance) GetErrorText() string { //gd:Expression.get_error_text
	return string(class(self).GetErrorText().String())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Expression

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Expression"))
	casted := Instance{*(*gdclass.Expression)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Parses the expression and returns an [enum Error] code.
You can optionally specify names of variables that may appear in the expression with [param input_names], so that you can bind them when it gets executed.
*/
//go:nosplit
func (self class) Parse(expression String.Readable, input_names Packed.Strings) Error.Code { //gd:Expression.parse
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(expression)))
	callframe.Arg(frame, pointers.Get(gd.InternalPackedStrings(input_names)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Expression.Bind_parse, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Executes the expression that was previously parsed by [method parse] and returns the result. Before you use the returned object, you should check if the method failed by calling [method has_execute_failed].
If you defined input variables in [method parse], you can specify their values in the inputs array, in the same order.
*/
//go:nosplit
func (self class) Execute(inputs Array.Any, base_instance [1]gd.Object, show_error bool, const_calls_only bool) variant.Any { //gd:Expression.execute
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalArray(inputs)))
	callframe.Arg(frame, pointers.Get(base_instance[0])[0])
	callframe.Arg(frame, show_error)
	callframe.Arg(frame, const_calls_only)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Expression.Bind_execute, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Through(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if [method execute] has failed.
*/
//go:nosplit
func (self class) HasExecuteFailed() bool { //gd:Expression.has_execute_failed
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Expression.Bind_has_execute_failed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the error text if [method parse] or [method execute] has failed.
*/
//go:nosplit
func (self class) GetErrorText() String.Readable { //gd:Expression.get_error_text
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Expression.Bind_get_error_text, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}
func (self class) AsExpression() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsExpression() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("Expression", func(ptr gd.Object) any { return [1]gdclass.Expression{*(*gdclass.Expression)(unsafe.Pointer(&ptr))} })
}
