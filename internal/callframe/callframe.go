// Package callframe provides a way to create Godot-compatible callframes.
package callframe

import (
	"reflect"
	"runtime"
	"sync"
	"unsafe"
)

var frames sync.Pool

// Frame can be used to reduce the number of allocations required to pass
// pointers across API boundaries. Since dynamic function calls and CGO
// always cause pointers to escape, a frame enables a block of memory to be
// used to copy in const value arguments. The frame is cached using [sync.Pool]
// and recycled across calls. Suitable for use when the function being
// called is only borrowing the pointer for the lifetime of the call. A
// frame has a fixed size, if it runs out of space, arguments will be
// allocated onto the Go heap as usual.
type Frame struct {
	arg uint8
	ret uint8

	pin runtime.Pinner
	ptr [16]*[8]uintptr
	typ [16]reflect.Type
	buf [20][8]uintptr
}

// Nil implements [Any].
type Nil struct{}

// Uintptr always returns 0.
func (Nil) Uintptr() uintptr { return 0 }

// New returns either a new [Frame], or a recycled frame.
func New() *Frame {
	frame, ok := frames.Get().(*Frame)
	if !ok {
		frame = new(Frame)
		frame.ptr = [16]*[8]uintptr{
			&frame.buf[0], &frame.buf[1], &frame.buf[2], &frame.buf[3],
			&frame.buf[4], &frame.buf[5], &frame.buf[6], &frame.buf[7],
			&frame.buf[8], &frame.buf[9], &frame.buf[10], &frame.buf[11],
			&frame.buf[12], &frame.buf[13], &frame.buf[14], &frame.buf[15],
		}
	}
	frame.pin.Pin(frame)
	return frame
}

type array struct{}

// Array of arguments.
type Args struct {
	_    [0]array
	void unsafe.Pointer
}

func (array Args) Uintptr() uintptr {
	return uintptr(array.void)
}

// Array returns a pointer to array of arguments offset by i.
func (frame *Frame) Array(i int) Args {
	return Args{
		void: unsafe.Pointer(&frame.ptr[i]),
	}
}

// Free recycles the [Frame]. Do not reuse the frame after
// calling this method.
func (frame *Frame) Free() {
	frame.arg = 0
	frame.ret = 0
	frame.pin.Unpin()
	frames.Put(frame)
}

// Ptr that does not escape as an allocation
// to the heap, T must be a value type that
// does not contain any pointers. The value
// must not be used after the frame has been
// freed.
type Ptr[T comparable] struct {
	_    [0]T
	void *[8]uintptr
}

// Uintptr returns the uintptr value of the pointer. Useful for passing to C code.
func (ptr Ptr[T]) Uintptr() uintptr {
	return uintptr(unsafe.Pointer(ptr.void))
}

// UnsafePoiner returns the unsafe.Pointer value of the pointer.
func (ptr Ptr[T]) UnsafePointer() unsafe.Pointer {
	return unsafe.Pointer(ptr.void)
}

// Get returns the value of the pointer.
func (ptr Ptr[T]) Get() T {
	return *(*T)(unsafe.Pointer(ptr.void))
}

// Arg adds a new argument to the call frame.
func Arg[T comparable](frame *Frame, arg T) Ptr[T] {
	if unsafe.Sizeof([1]T{}) > 8*unsafe.Sizeof(uintptr(0)) {
		panic("callframe: argument too large")
	}
	*(*T)(unsafe.Pointer(&frame.buf[frame.arg])) = arg
	frame.arg++
	return Ptr[T]{void: &frame.buf[frame.arg-1]}
}

// Align a size to the next multiple of alignment (rounding up).
func Align(size uintptr, alignment uintptr) uintptr {
	return (size + alignment - 1) &^ (alignment - 1)
}

// Ret prepares an expected return value that will be available after the call has been
// made.
func Ret[T comparable](frame *Frame) Ptr[T] {
	if unsafe.Sizeof([1]T{}) > 8*unsafe.Sizeof(uintptr(0)) {
		panic("callframe: return value too large")
	}
	frame.ret++
	return Ptr[T]{void: &frame.buf[frame.ret+15]}
}
