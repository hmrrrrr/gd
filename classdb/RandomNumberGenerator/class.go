// Package RandomNumberGenerator provides methods for working with RandomNumberGenerator object instances.
package RandomNumberGenerator

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
RandomNumberGenerator is a class for generating pseudo-random numbers. It currently uses [url=https://www.pcg-random.org/]PCG32[/url].
[b]Note:[/b] The underlying algorithm is an implementation detail and should not be depended upon.
To generate a random float number (within a given range) based on a time-dependent seed:
[codeblock]
var rng = RandomNumberGenerator.new()
func _ready():

	var my_random_number = rng.randf_range(-10.0, 10.0)

[/codeblock]
*/
type Instance [1]gdclass.RandomNumberGenerator

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRandomNumberGenerator() Instance
}

/*
Returns a pseudo-random 32-bit unsigned integer between [code]0[/code] and [code]4294967295[/code] (inclusive).
*/
func (self Instance) Randi() int { //gd:RandomNumberGenerator.randi
	return int(int(class(self).Randi()))
}

/*
Returns a pseudo-random float between [code]0.0[/code] and [code]1.0[/code] (inclusive).
*/
func (self Instance) Randf() Float.X { //gd:RandomNumberGenerator.randf
	return Float.X(Float.X(class(self).Randf()))
}

/*
Returns a [url=https://en.wikipedia.org/wiki/Normal_distribution]normally-distributed[/url], pseudo-random floating-point number from the specified [param mean] and a standard [param deviation]. This is also known as a Gaussian distribution.
[b]Note:[/b] This method uses the [url=https://en.wikipedia.org/wiki/Box%E2%80%93Muller_transform]Box-Muller transform[/url] algorithm.
*/
func (self Instance) Randfn() Float.X { //gd:RandomNumberGenerator.randfn
	return Float.X(Float.X(class(self).Randfn(float64(0.0), float64(1.0))))
}

/*
Returns a pseudo-random float between [param from] and [param to] (inclusive).
*/
func (self Instance) RandfRange(from Float.X, to Float.X) Float.X { //gd:RandomNumberGenerator.randf_range
	return Float.X(Float.X(class(self).RandfRange(float64(from), float64(to))))
}

/*
Returns a pseudo-random 32-bit signed integer between [param from] and [param to] (inclusive).
*/
func (self Instance) RandiRange(from int, to int) int { //gd:RandomNumberGenerator.randi_range
	return int(int(class(self).RandiRange(int64(from), int64(to))))
}

/*
Returns a random index with non-uniform weights. Prints an error and returns [code]-1[/code] if the array is empty.
[codeblocks]
[gdscript]
var rng = RandomNumberGenerator.new()

var my_array = ["one", "two", "three", "four"]
var weights = PackedFloat32Array([0.5, 1, 1, 2])

# Prints one of the four elements in `my_array`.
# It is more likely to print "four", and less likely to print "one".
print(my_array[rng.rand_weighted(weights)])
[/gdscript]
[/codeblocks]
*/
func (self Instance) RandWeighted(weights []float32) int { //gd:RandomNumberGenerator.rand_weighted
	return int(int(class(self).RandWeighted(Packed.New(weights...))))
}

/*
Sets up a time-based seed for this [RandomNumberGenerator] instance. Unlike the [@GlobalScope] random number generation functions, different [RandomNumberGenerator] instances can use different seeds.
*/
func (self Instance) Randomize() { //gd:RandomNumberGenerator.randomize
	class(self).Randomize()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RandomNumberGenerator

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RandomNumberGenerator"))
	casted := Instance{*(*gdclass.RandomNumberGenerator)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Seed() int {
	return int(int(class(self).GetSeed()))
}

func (self Instance) SetSeed(value int) {
	class(self).SetSeed(int64(value))
}

func (self Instance) State() int {
	return int(int(class(self).GetState()))
}

func (self Instance) SetState(value int) {
	class(self).SetState(int64(value))
}

//go:nosplit
func (self class) SetSeed(seed int64) { //gd:RandomNumberGenerator.set_seed
	var frame = callframe.New()
	callframe.Arg(frame, seed)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RandomNumberGenerator.Bind_set_seed, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSeed() int64 { //gd:RandomNumberGenerator.get_seed
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RandomNumberGenerator.Bind_get_seed, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetState(state int64) { //gd:RandomNumberGenerator.set_state
	var frame = callframe.New()
	callframe.Arg(frame, state)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RandomNumberGenerator.Bind_set_state, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetState() int64 { //gd:RandomNumberGenerator.get_state
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RandomNumberGenerator.Bind_get_state, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a pseudo-random 32-bit unsigned integer between [code]0[/code] and [code]4294967295[/code] (inclusive).
*/
//go:nosplit
func (self class) Randi() int64 { //gd:RandomNumberGenerator.randi
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RandomNumberGenerator.Bind_randi, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a pseudo-random float between [code]0.0[/code] and [code]1.0[/code] (inclusive).
*/
//go:nosplit
func (self class) Randf() float64 { //gd:RandomNumberGenerator.randf
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RandomNumberGenerator.Bind_randf, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a [url=https://en.wikipedia.org/wiki/Normal_distribution]normally-distributed[/url], pseudo-random floating-point number from the specified [param mean] and a standard [param deviation]. This is also known as a Gaussian distribution.
[b]Note:[/b] This method uses the [url=https://en.wikipedia.org/wiki/Box%E2%80%93Muller_transform]Box-Muller transform[/url] algorithm.
*/
//go:nosplit
func (self class) Randfn(mean float64, deviation float64) float64 { //gd:RandomNumberGenerator.randfn
	var frame = callframe.New()
	callframe.Arg(frame, mean)
	callframe.Arg(frame, deviation)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RandomNumberGenerator.Bind_randfn, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a pseudo-random float between [param from] and [param to] (inclusive).
*/
//go:nosplit
func (self class) RandfRange(from float64, to float64) float64 { //gd:RandomNumberGenerator.randf_range
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RandomNumberGenerator.Bind_randf_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a pseudo-random 32-bit signed integer between [param from] and [param to] (inclusive).
*/
//go:nosplit
func (self class) RandiRange(from int64, to int64) int64 { //gd:RandomNumberGenerator.randi_range
	var frame = callframe.New()
	callframe.Arg(frame, from)
	callframe.Arg(frame, to)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RandomNumberGenerator.Bind_randi_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a random index with non-uniform weights. Prints an error and returns [code]-1[/code] if the array is empty.
[codeblocks]
[gdscript]
var rng = RandomNumberGenerator.new()

var my_array = ["one", "two", "three", "four"]
var weights = PackedFloat32Array([0.5, 1, 1, 2])

# Prints one of the four elements in `my_array`.
# It is more likely to print "four", and less likely to print "one".
print(my_array[rng.rand_weighted(weights)])
[/gdscript]
[/codeblocks]
*/
//go:nosplit
func (self class) RandWeighted(weights Packed.Array[float32]) int64 { //gd:RandomNumberGenerator.rand_weighted
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedFloat32Array, float32](weights)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RandomNumberGenerator.Bind_rand_weighted, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets up a time-based seed for this [RandomNumberGenerator] instance. Unlike the [@GlobalScope] random number generation functions, different [RandomNumberGenerator] instances can use different seeds.
*/
//go:nosplit
func (self class) Randomize() { //gd:RandomNumberGenerator.randomize
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RandomNumberGenerator.Bind_randomize, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func (self class) AsRandomNumberGenerator() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRandomNumberGenerator() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
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
	gdclass.Register("RandomNumberGenerator", func(ptr gd.Object) any {
		return [1]gdclass.RandomNumberGenerator{*(*gdclass.RandomNumberGenerator)(unsafe.Pointer(&ptr))}
	})
}
