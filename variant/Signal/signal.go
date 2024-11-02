// Package Signal provides a type representing an event stream or pub/sub message queue.
package Signal

import (
	"sync"

	"grow.graphics/gd/gdconst"
	gd "grow.graphics/gd/internal"
	"grow.graphics/gd/variant/Callable"
)

// Chan is a multi-producer, multi-consumer channel, when T is a value, it denotes the
// value being sent on the channel, when T is a function, the results of calling that
// function are passed on the channel instead.
type Chan[T any] struct {
	_ [0]sync.Mutex //nocopy

	owner any
	topic string
	funcs []Callable.Func
	proxy *gd.Signal
}

const ErrInvalidParameter = gdconst.ErrInvalidParameter

// Attach connects this signal to the specified [Callable.Func]
// A signal can only be connected once to the same [Callable.Func].
// If the signal is already connected, returns [ErrInvalidParameter]
func (c *Chan[T]) Attach(fn Callable.Func) error { //gd:Signal.connect
	if c.proxy != nil {
		//return c.proxy.C(fn)
	}
	for _, f := range c.funcs {
		if f == fn {
			return ErrInvalidParameter
		}
	}
	c.funcs = append(c.funcs, fn)
	return nil
}

// Remove disconnects this signal from the specified [Callable.Func].
func (c *Chan[T]) Remove(fn Callable.Func) { //gd:Signal.disconnect
	if c.proxy != nil {
		//if c.proxy.Has(fn) {
		//	c.proxy.Remove(fn)
		//}
		return
	}
	for i, f := range c.funcs {
		if f == fn {
			c.funcs = append(c.funcs[:i], c.funcs[i+1:]...)
			return
		}
	}
}

// Emit emits this signal. All Callables connected to this signal will be triggered.
func (c *Chan[T]) Emit(signal T) { //gd:Signal.emit
	if c.proxy != nil {
		//c.proxy.Emit(signal)
		return
	}
	for _, f := range c.funcs {
		f.Call(signal)
	}
}

// Callables returns a slice of callables for this signal.
func (c *Chan[T]) Callables() []Callable.Func { //gd:Signal.get_connections
	if c.proxy != nil {
		//return c.proxy.Callables()
	}
	return c.funcs
}

// Name returns the name of this signal.
func (c *Chan[T]) Name() string { //gd:Signal.get_name
	if c.proxy != nil {
		//return c.proxy.Name()
	}
	return c.topic
}

// Emitter returns the object that emits this signal.
func (c *Chan[T]) Emitter() any { //gd:Signal.get_object Signal.get_object_id
	if c.proxy != nil {
		//return c.proxy.Emitter()
	}
	return c.owner
}

// Has returns true if the specified callable is connected to this signal.
func (c *Chan[T]) Has(fn Callable.Func) bool { //gd:Signal.is_connected
	if c.proxy != nil {
		//return c.proxy.Has(fn)
	}
	for _, f := range c.funcs {
		if f == fn {
			return true
		}
	}
	return false
}

// IsNull returns true if this signal is not connected to any callable.
func (c *Chan[T]) IsNull() bool { //gd:Signal.is_null
	return c.owner == nil && c.proxy == nil && len(c.funcs) == 0
}

// Proxy converts the signal into a proxied signal and returns the proxy.
func (c *Chan[T]) Proxy(freed bool, alloc func(any, string) gd.Signal) *gd.Signal {
	if c.proxy != nil {
		//return c.proxy
	}
	if c.IsNull() {
		return nil
	}
	//proxy := alloc(c.owner, c.topic)
	//for _, f := range c.funcs {
	//	proxy.Attach(f)
	//}
	c.funcs = nil
	//c.proxy = proxy
	//return proxy
	return nil
}

// Reset clears all connections to this signal.
func (c *Chan[T]) Reset() {
	if c.proxy != nil {
		c.proxy.Free()
	}
	*c = Chan[T]{}
}
