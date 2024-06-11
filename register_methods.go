//go:build !generate

package gd

import (
	"fmt"
	"reflect"

	gd "grow.graphics/gd/internal"
	"runtime.link/mmm"
)

func registerMethods(godot Context, class StringName, rtype reflect.Type) {
	classTypePtr := reflect.PointerTo(rtype)
	for i := 0; i < classTypePtr.NumMethod(); i++ {
		i := i
		method := classTypePtr.Method(i)
		if !method.IsExported() || method.Type.NumIn() < 2 || method.Type.In(1) != reflect.TypeOf(Context{}) {
			continue
		}
		var arguments = make([]gd.PropertyInfo, 0, method.Type.NumIn()-2)
		var metadatas = make([]gd.ClassMethodArgumentMetadata, 0, method.Type.NumIn()-2)
		for i := 2; i < method.Type.NumIn(); i++ {
			arguments = append(arguments,
				propertyOf(godot, reflect.StructField{Name: "arg" + fmt.Sprint(i), Type: method.Type.In(i)}))
			metadatas = append(metadatas, 0)
		}
		var returns *gd.PropertyInfo
		var returnMetadata gd.ClassMethodArgumentMetadata
		if method.Type.NumOut() > 0 {
			property := propertyOf(godot, reflect.StructField{Name: "result", Type: method.Type.Out(0)})
			returns = &property
			returnMetadata = 0
		}
		godot.API.ClassDB.RegisterClassMethod(godot, godot.API.ExtensionToken, class, gd.Method{
			Name: godot.StringName(method.Name),
			// FIXME type check and return an error if arguments are invalid.
			Call: func(godot gd.Context, instance any, v ...gd.Variant) (gd.Variant, error) {
				var args = make([]reflect.Value, len(v)+1)
				if method.Type.In(1) == reflect.TypeOf(Context{}) {
					args[0] = reflect.ValueOf(godot)
					for i := 0; i < len(v); i++ {
						if method.Type.In(i + 2).Implements(reflect.TypeOf([0]gd.IsClass{}).Elem()) {
							var obj = reflect.New(method.Type.In(i + 2))
							obj.Interface().(gd.PointerToClass).SetPointer(gd.LetVariantAsPointerType[gd.Pointer](godot, v[i], TypeObject))
							args[i+1] = obj.Elem()
						} else {
							args[i+1] = reflect.ValueOf(v[i].Interface(godot))
						}
					}
				} else {
					args = args[:len(v)]
					for i := 0; i < len(v); i++ {
						if method.Type.In(i + 1).Implements(reflect.TypeOf([0]gd.IsClass{}).Elem()) {
							var obj = reflect.New(method.Type.In(i + 2))
							obj.Interface().(gd.PointerToClass).SetPointer(gd.LetVariantAsPointerType[gd.Pointer](godot, v[i], TypeObject))
							args[i] = obj.Elem()
						} else {
							args[i] = reflect.ValueOf(v[i].Interface(godot))
						}
					}
				}
				rets := reflect.ValueOf(instance.(*instanceImplementation).Value).Method(i).Call(args)
				if len(rets) > 0 {
					return godot.Variant(rets[0].Interface()), nil
				}
				return gd.Variant{}, nil
			},
			PointerCall: func(instance any, args gd.UnsafeArgs, ret gd.UnsafeBack) {
				slowCall(godot.API, reflect.ValueOf(instance.(*instanceImplementation).Value).Method(i), args, ret)
			},
			Arguments:           arguments,
			ArgumentsMetadata:   metadatas,
			ReturnValueInfo:     returns,
			ReturnValueMetadata: returnMetadata,
		})
	}
}

func slowCall(godot *gd.API, method reflect.Value, p_args gd.UnsafeArgs, p_ret gd.UnsafeBack) {
	var ctx = gd.NewContext(godot)
	defer ctx.End()
	var (
		args = make([]reflect.Value, method.Type().NumIn())
	)
	var offset = 0
	if method.Type().In(0) == reflect.TypeOf(Context{}) {
		args[0] = reflect.ValueOf(ctx)
		offset = 1
	}
	for i := offset; i < method.Type().NumIn(); i++ {
		switch method.Type().In(i) {
		case reflect.TypeOf(Bool(false)):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Bool](p_args, i-1))
		case reflect.TypeOf(Int(0)):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Int](p_args, i-1))
		case reflect.TypeOf(Float(0)):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Float](p_args, i-1))
		case reflect.TypeOf(String{}):
			ptr := gd.UnsafeGet[uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[String](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(Vector2{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Vector2](p_args, i-1))
		case reflect.TypeOf(Vector2i{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Vector2i](p_args, i-1))
		case reflect.TypeOf(Rect2{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Rect2](p_args, i-1))
		case reflect.TypeOf(Rect2i{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Rect2i](p_args, i-1))
		case reflect.TypeOf(Vector3{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Vector3](p_args, i-1))
		case reflect.TypeOf(Vector3i{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Vector3i](p_args, i-1))
		case reflect.TypeOf(Transform2D{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Transform2D](p_args, i-1))
		case reflect.TypeOf(Vector4{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Vector4](p_args, i-1))
		case reflect.TypeOf(Vector4i{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Vector4i](p_args, i-1))
		case reflect.TypeOf(Plane{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Plane](p_args, i-1))
		case reflect.TypeOf(Quaternion{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Quaternion](p_args, i-1))
		case reflect.TypeOf(AABB{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[AABB](p_args, i-1))
		case reflect.TypeOf(Basis{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Basis](p_args, i-1))
		case reflect.TypeOf(Transform3D{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Transform3D](p_args, i-1))
		case reflect.TypeOf(Projection{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Projection](p_args, i-1))
		case reflect.TypeOf(Color{}):
			args[i] = reflect.ValueOf(gd.UnsafeGet[Color](p_args, i-1))
		case reflect.TypeOf(StringName{}):
			ptr := gd.UnsafeGet[uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[StringName](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(NodePath{}):
			ptr := gd.UnsafeGet[uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[NodePath](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(RID(0)):
			args[i] = reflect.ValueOf(gd.UnsafeGet[RID](p_args, i-1))
		case reflect.TypeOf(Object{}):
			ptr := gd.UnsafeGet[uintptr](p_args, i-1)
			var obj Object
			obj.SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{ptr}))
			args[i] = reflect.ValueOf(obj)
		case reflect.TypeOf(Callable{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[Callable](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(gd.Signal{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[gd.Signal](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(Dictionary{}):
			ptr := gd.UnsafeGet[uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[Dictionary](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(Array{}):
			ptr := gd.UnsafeGet[uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[Array](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(PackedByteArray{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[PackedByteArray](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(PackedInt32Array{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[PackedInt32Array](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(PackedInt64Array{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[PackedInt64Array](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(PackedFloat32Array{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[PackedFloat32Array](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(PackedFloat64Array{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[PackedFloat64Array](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(PackedStringArray{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[PackedStringArray](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(PackedVector2Array{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[PackedVector2Array](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(PackedVector3Array{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[PackedVector3Array](ctx.Lifetime, ctx.API, ptr))
		case reflect.TypeOf(PackedColorArray{}):
			ptr := gd.UnsafeGet[[2]uintptr](p_args, i-1)
			args[i] = reflect.ValueOf(mmm.Let[PackedColorArray](ctx.Lifetime, ctx.API, ptr))
		default:
			panic(fmt.Sprintf("gdextension: unsupported Godot -> Go type %v", method.Type().In(i)))
		}
	}
	rrets := method.Call(args)
	if len(rrets) > 0 {
		switch val := rrets[0].Interface().(type) {
		case Bool:
			gd.UnsafeSet[Bool](p_ret, val)
		case Int:
			gd.UnsafeSet[Int](p_ret, val)
		case Float:
			gd.UnsafeSet[Float](p_ret, val)
		case String:
			gd.UnsafeSet[uintptr](p_ret, mmm.Get(val))
			mmm.End(val)
		case Vector2:
			gd.UnsafeSet[Vector2](p_ret, val)
		case Vector2i:
			gd.UnsafeSet[Vector2i](p_ret, val)
		case Rect2:
			gd.UnsafeSet[Rect2](p_ret, val)
		case Rect2i:
			gd.UnsafeSet[Rect2i](p_ret, val)
		case Vector3:
			gd.UnsafeSet[Vector3](p_ret, val)
		case Vector3i:
			gd.UnsafeSet[Vector3i](p_ret, val)
		case Transform2D:
			gd.UnsafeSet[Transform2D](p_ret, val)
		case Vector4:
			gd.UnsafeSet[Vector4](p_ret, val)
		case Vector4i:
			gd.UnsafeSet[Vector4i](p_ret, val)
		case Plane:
			gd.UnsafeSet[Plane](p_ret, val)
		case Quaternion:
			gd.UnsafeSet[Quaternion](p_ret, val)
		case AABB:
			gd.UnsafeSet[AABB](p_ret, val)
		case Basis:
			gd.UnsafeSet[Basis](p_ret, val)
		case Transform3D:
			gd.UnsafeSet[Transform3D](p_ret, val)
		case Projection:
			gd.UnsafeSet[Projection](p_ret, val)
		case Color:
			gd.UnsafeSet[Color](p_ret, val)
		case StringName:
			gd.UnsafeSet[uintptr](p_ret, mmm.Get(val))
			mmm.End(val)
		case NodePath:
			gd.UnsafeSet[uintptr](p_ret, mmm.Get(val))
			mmm.End(val)
		case RID:
			gd.UnsafeSet[RID](p_ret, val)
		case Object:
			gd.UnsafeSet[uintptr](p_ret, mmm.Get(val.AsPointer())[0])
			mmm.End(val.AsPointer())
		case Callable:
			gd.UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
			mmm.End(val)
		case gd.Signal:
			gd.UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
			mmm.End(val)
		case Dictionary:
			gd.UnsafeSet[uintptr](p_ret, mmm.Get(val))
			mmm.End(val)
		case Array:
			gd.UnsafeSet[uintptr](p_ret, mmm.Get(val))
			mmm.End(val)
		case PackedByteArray:
			gd.UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		case PackedInt32Array:
			gd.UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		case PackedInt64Array:
			gd.UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		case PackedFloat32Array:
			gd.UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		case PackedFloat64Array:
			gd.UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		case PackedStringArray:
			gd.UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		case PackedVector2Array:
			gd.UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		case PackedVector3Array:
			gd.UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		case PackedColorArray:
			gd.UnsafeSet[[2]uintptr](p_ret, mmm.Get(val))
		default:
			panic(fmt.Sprintf("gdextension: unsupported Go -> Godot type %v", method.Type().Out(0)))
		}
	}
}
