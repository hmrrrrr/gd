package StreamPeerGZIP

import "unsafe"
import "reflect"
import "runtime.link/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/object/StreamPeer"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This class allows to compress or decompress data using GZIP/deflate in a streaming fashion. This is particularly useful when compressing or decompressing files that have to be sent through the network without needing to allocate them all in memory.
After starting the stream via [method start_compression] (or [method start_decompression]), calling [method StreamPeer.put_partial_data] on this stream will compress (or decompress) the data, writing it to the internal buffer. Calling [method StreamPeer.get_available_bytes] will return the pending bytes in the internal buffer, and [method StreamPeer.get_partial_data] will retrieve the compressed (or decompressed) bytes from it. When the stream is over, you must call [method finish] to ensure the internal buffer is properly flushed (make sure to call [method StreamPeer.get_available_bytes] on last time to check if more data needs to be read after that).

*/
type Simple [1]classdb.StreamPeerGZIP
func (self Simple) StartCompression(use_deflate bool, buffer_size int) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).StartCompression(use_deflate, gd.Int(buffer_size)))
}
func (self Simple) StartDecompression(use_deflate bool, buffer_size int) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).StartDecompression(use_deflate, gd.Int(buffer_size)))
}
func (self Simple) Finish() gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(Expert(self).Finish())
}
func (self Simple) Clear() {
	gc := gd.GarbageCollector(); _ = gc
	Expert(self).Clear()
}
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.StreamPeerGZIP
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Start the stream in compression mode with the given [param buffer_size], if [param use_deflate] is [code]true[/code] uses deflate instead of GZIP.
*/
//go:nosplit
func (self class) StartCompression(use_deflate bool, buffer_size gd.Int) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, use_deflate)
	callframe.Arg(frame, buffer_size)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeerGZIP.Bind_start_compression, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Start the stream in decompression mode with the given [param buffer_size], if [param use_deflate] is [code]true[/code] uses deflate instead of GZIP.
*/
//go:nosplit
func (self class) StartDecompression(use_deflate bool, buffer_size gd.Int) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, use_deflate)
	callframe.Arg(frame, buffer_size)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeerGZIP.Bind_start_decompression, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Finalizes the stream, compressing or decompressing any buffered chunk left.
*/
//go:nosplit
func (self class) Finish() int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeerGZIP.Bind_finish, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Clears this stream, resetting the internal state.
*/
//go:nosplit
func (self class) Clear()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.StreamPeerGZIP.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}

//go:nosplit
func (self class) AsStreamPeerGZIP() Expert { return self[0].AsStreamPeerGZIP() }


//go:nosplit
func (self Simple) AsStreamPeerGZIP() Simple { return self[0].AsStreamPeerGZIP() }


//go:nosplit
func (self class) AsStreamPeer() StreamPeer.Expert { return self[0].AsStreamPeer() }


//go:nosplit
func (self Simple) AsStreamPeer() StreamPeer.Simple { return self[0].AsStreamPeer() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("StreamPeerGZIP", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
