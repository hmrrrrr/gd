package gdextension

import (
	"reflect"
	"runtime/cgo"
	"strings"
	"unsafe"

	gd "grow.graphics/gd/internal"
	"runtime.link/api/call"
	"runtime.link/mmm"
)

func registerMethods(godot gd.Context, class gd.StringName, rtype reflect.Type) {
	classTypePtr := reflect.PointerTo(rtype)
	var returnInfo = (*gd.PropertyInfo)(godot.API.Memory.Allocate(unsafe.Sizeof(gd.PropertyInfo{})))

	for i := 0; i < classTypePtr.NumMethod(); i++ {
		method := classTypePtr.Method(i)
		if !method.IsExported() {
			continue
		}
		if method.Type.NumIn() > 2 || method.Type.NumOut() > 0 || method.Type.NumIn() < 2 {
			continue
		}
		if method.Type.In(1) != reflect.TypeOf(gd.Context{}) {
			continue
		}

		var methodName = godot.StringName(strings.ToLower(method.Name))

		var frame = call.New()

		var info = gd.ClassMethodInfo{
			Name:            (gd.StringNamePtr)(unsafe.Pointer(call.Arg(frame, mmm.Get(methodName)).Uintptr())),
			HasReturnValue:  method.Type.NumOut() > 0,
			ReturnValueInfo: returnInfo,
		}
		info.PointerCall.Set(func(userdata unsafe.Pointer, instance cgo.Handle, args, ret unsafe.Pointer, err *gd.CallError) {
			ctx := gd.NewContext(godot.API)
			var in = make([]reflect.Value, method.Type.NumIn())
			in[0] = reflect.ValueOf(instance.Value())
			in[1] = reflect.ValueOf(ctx)
			var out = method.Func.Call(in)
			if len(out) > 0 {
				reflect.NewAt(method.Type.Out(0), ret).Elem().Set(out[0])
			}
			ctx.End()
		})
		godot.API.ClassDB.RegisterClassMethod(godot.API.ExtensionToken,
			(gd.StringNamePtr)(unsafe.Pointer(call.Arg(frame, mmm.Get(class)).Uintptr())), &info)
		frame.Free()
	}
}
