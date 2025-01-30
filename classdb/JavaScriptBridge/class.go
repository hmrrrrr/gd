// Package JavaScriptBridge provides methods for working with JavaScriptBridge object instances.
package JavaScriptBridge

import "unsafe"
import "sync"
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
The JavaScriptBridge singleton is implemented only in the Web export. It's used to access the browser's JavaScript context. This allows interaction with embedding pages or calling third-party JavaScript APIs.
[b]Note:[/b] This singleton can be disabled at build-time to improve security. By default, the JavaScriptBridge singleton is enabled. Official export templates also have the JavaScriptBridge singleton enabled. See [url=$DOCS_URL/contributing/development/compiling/compiling_for_web.html]Compiling for the Web[/url] in the documentation for more information.
*/
var self [1]gdclass.JavaScriptBridge
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.JavaScriptBridge)
	self = *(*[1]gdclass.JavaScriptBridge)(unsafe.Pointer(&obj))
}

/*
Execute the string [param code] as JavaScript code within the browser window. This is a call to the actual global JavaScript function [code skip-lint]eval()[/code].
If [param use_global_execution_context] is [code]true[/code], the code will be evaluated in the global execution context. Otherwise, it is evaluated in the execution context of a function within the engine's runtime environment.
*/
func Eval(code string) any { //gd:JavaScriptBridge.eval
	once.Do(singleton)
	return any(class(self).Eval(String.New(code), false).Interface())
}

/*
Returns an interface to a JavaScript object that can be used by scripts. The [param interface] must be a valid property of the JavaScript [code]window[/code]. The callback must accept a single [Array] argument, which will contain the JavaScript [code]arguments[/code]. See [JavaScriptObject] for usage.
*/
func GetInterface(intf string) [1]gdclass.JavaScriptObject { //gd:JavaScriptBridge.get_interface
	once.Do(singleton)
	return [1]gdclass.JavaScriptObject(class(self).GetInterface(String.New(intf)))
}

/*
Creates a reference to a [Callable] that can be used as a callback by JavaScript. The reference must be kept until the callback happens, or it won't be called at all. See [JavaScriptObject] for usage.
*/
func CreateCallback(callable Callable.Function) [1]gdclass.JavaScriptObject { //gd:JavaScriptBridge.create_callback
	once.Do(singleton)
	return [1]gdclass.JavaScriptObject(class(self).CreateCallback(Callable.New(callable)))
}

/*
Prompts the user to download a file containing the specified [param buffer]. The file will have the given [param name] and [param mime] type.
[b]Note:[/b] The browser may override the [url=https://en.wikipedia.org/wiki/Media_type]MIME type[/url] provided based on the file [param name]'s extension.
[b]Note:[/b] Browsers might block the download if [method download_buffer] is not being called from a user interaction (e.g. button click).
[b]Note:[/b] Browsers might ask the user for permission or block the download if multiple download requests are made in a quick succession.
*/
func DownloadBuffer(buffer []byte, name string) { //gd:JavaScriptBridge.download_buffer
	once.Do(singleton)
	class(self).DownloadBuffer(Packed.Bytes(Packed.New(buffer...)), String.New(name), String.New("application/octet-stream"))
}

/*
Returns [code]true[/code] if a new version of the progressive web app is waiting to be activated.
[b]Note:[/b] Only relevant when exported as a Progressive Web App.
*/
func PwaNeedsUpdate() bool { //gd:JavaScriptBridge.pwa_needs_update
	once.Do(singleton)
	return bool(class(self).PwaNeedsUpdate())
}

/*
Performs the live update of the progressive web app. Forcing the new version to be installed and the page to be reloaded.
[b]Note:[/b] Your application will be [b]reloaded in all browser tabs[/b].
[b]Note:[/b] Only relevant when exported as a Progressive Web App and [method pwa_needs_update] returns [code]true[/code].
*/
func PwaUpdate() error { //gd:JavaScriptBridge.pwa_update
	once.Do(singleton)
	return error(gd.ToError(class(self).PwaUpdate()))
}

/*
Force synchronization of the persistent file system (when enabled).
[b]Note:[/b] This is only useful for modules or extensions that can't use [FileAccess] to write files.
*/
func ForceFsSync() { //gd:JavaScriptBridge.force_fs_sync
	once.Do(singleton)
	class(self).ForceFsSync()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]gdclass.JavaScriptBridge

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Execute the string [param code] as JavaScript code within the browser window. This is a call to the actual global JavaScript function [code skip-lint]eval()[/code].
If [param use_global_execution_context] is [code]true[/code], the code will be evaluated in the global execution context. Otherwise, it is evaluated in the execution context of a function within the engine's runtime environment.
*/
//go:nosplit
func (self class) Eval(code String.Readable, use_global_execution_context bool) variant.Any { //gd:JavaScriptBridge.eval
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(code)))
	callframe.Arg(frame, use_global_execution_context)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JavaScriptBridge.Bind_eval, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Implementation(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns an interface to a JavaScript object that can be used by scripts. The [param interface] must be a valid property of the JavaScript [code]window[/code]. The callback must accept a single [Array] argument, which will contain the JavaScript [code]arguments[/code]. See [JavaScriptObject] for usage.
*/
//go:nosplit
func (self class) GetInterface(intf String.Readable) [1]gdclass.JavaScriptObject { //gd:JavaScriptBridge.get_interface
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(intf)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JavaScriptBridge.Bind_get_interface, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.JavaScriptObject{gd.PointerWithOwnershipTransferredToGo[gdclass.JavaScriptObject](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Creates a reference to a [Callable] that can be used as a callback by JavaScript. The reference must be kept until the callback happens, or it won't be called at all. See [JavaScriptObject] for usage.
*/
//go:nosplit
func (self class) CreateCallback(callable Callable.Function) [1]gdclass.JavaScriptObject { //gd:JavaScriptBridge.create_callback
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(callable)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JavaScriptBridge.Bind_create_callback, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.JavaScriptObject{gd.PointerWithOwnershipTransferredToGo[gdclass.JavaScriptObject](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Prompts the user to download a file containing the specified [param buffer]. The file will have the given [param name] and [param mime] type.
[b]Note:[/b] The browser may override the [url=https://en.wikipedia.org/wiki/Media_type]MIME type[/url] provided based on the file [param name]'s extension.
[b]Note:[/b] Browsers might block the download if [method download_buffer] is not being called from a user interaction (e.g. button click).
[b]Note:[/b] Browsers might ask the user for permission or block the download if multiple download requests are made in a quick succession.
*/
//go:nosplit
func (self class) DownloadBuffer(buffer Packed.Bytes, name String.Readable, mime String.Readable) { //gd:JavaScriptBridge.download_buffer
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedByteArray, byte](Packed.Array[byte](buffer))))
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(mime)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JavaScriptBridge.Bind_download_buffer, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if a new version of the progressive web app is waiting to be activated.
[b]Note:[/b] Only relevant when exported as a Progressive Web App.
*/
//go:nosplit
func (self class) PwaNeedsUpdate() bool { //gd:JavaScriptBridge.pwa_needs_update
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JavaScriptBridge.Bind_pwa_needs_update, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Performs the live update of the progressive web app. Forcing the new version to be installed and the page to be reloaded.
[b]Note:[/b] Your application will be [b]reloaded in all browser tabs[/b].
[b]Note:[/b] Only relevant when exported as a Progressive Web App and [method pwa_needs_update] returns [code]true[/code].
*/
//go:nosplit
func (self class) PwaUpdate() Error.Code { //gd:JavaScriptBridge.pwa_update
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JavaScriptBridge.Bind_pwa_update, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Force synchronization of the persistent file system (when enabled).
[b]Note:[/b] This is only useful for modules or extensions that can't use [FileAccess] to write files.
*/
//go:nosplit
func (self class) ForceFsSync() { //gd:JavaScriptBridge.force_fs_sync
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JavaScriptBridge.Bind_force_fs_sync, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}
func OnPwaUpdateAvailable(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("pwa_update_available"), gd.NewCallable(cb), 0)
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("JavaScriptBridge", func(ptr gd.Object) any {
		return [1]gdclass.JavaScriptBridge{*(*gdclass.JavaScriptBridge)(unsafe.Pointer(&ptr))}
	})
}
