package gd

import (
	"reflect"
	"unsafe"

	NodePathType "graphics.gd/variant/NodePath"
)

func VariantTypeOf(rtype reflect.Type) (vtype VariantType, ok bool) {
	switch rtype.Kind() {
	case reflect.Bool:
		return TypeBool, true
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint:
		return TypeInt, true
	case reflect.Uint64:
		return TypeRID, true
	case reflect.Float32, reflect.Float64:
		return TypeFloat, true
	case reflect.Complex64, reflect.Complex128:
		return TypeVector2, true
	case reflect.Pointer:
		return VariantTypeOf(rtype.Elem())
	case reflect.Func:
		return TypeCallable, true
	case reflect.Array:
		return TypeArray, true
	case reflect.String:
		if rtype == reflect.TypeFor[NodePathType.String]() {
			return TypeNodePath, true
		}
		return TypeString, true
	case reflect.Slice:
		switch rtype.Elem().Kind() {
		case reflect.Uint8:
			return TypePackedByteArray, true
		case reflect.Int32:
			return TypePackedInt32Array, true
		case reflect.Int64:
			return TypePackedInt64Array, true
		case reflect.Float32:
			return TypePackedFloat32Array, true
		case reflect.Float64:
			return TypePackedFloat64Array, true
		case reflect.String:
			return TypePackedStringArray, true
		default:
			switch rtype.Elem() {
			case reflect.TypeFor[Vector2]():
				return TypePackedVector2Array, true
			case reflect.TypeFor[Vector3]():
				return TypePackedVector3Array, true
			case reflect.TypeFor[Color]():
				return TypePackedColorArray, true
			case reflect.TypeFor[Vector4]():
				return TypePackedVector4Array, true
			default:
				return TypeArray, true
			}
		}
	case reflect.Map:
		return TypeDictionary, true
	case reflect.Struct:
		switch rtype {
		case reflect.TypeOf([0]Variant{}).Elem():
			vtype = TypeNil
		case reflect.TypeOf([0]Bool{}).Elem():
			vtype = TypeBool
		case reflect.TypeOf([0]Int{}).Elem():
			vtype = TypeInt
		case reflect.TypeOf([0]Float{}).Elem():
			vtype = TypeFloat
		case reflect.TypeOf([0]String{}).Elem():
			vtype = TypeString
		case reflect.TypeOf([0]Vector2{}).Elem():
			vtype = TypeVector2
		case reflect.TypeOf([0]Vector2i{}).Elem():
			vtype = TypeVector2i
		case reflect.TypeOf([0]Rect2{}).Elem():
			vtype = TypeRect2
		case reflect.TypeOf([0]Rect2i{}).Elem():
			vtype = TypeRect2i
		case reflect.TypeOf([0]Vector3{}).Elem():
			vtype = TypeVector3
		case reflect.TypeOf([0]Vector3i{}).Elem():
			vtype = TypeVector3i
		case reflect.TypeOf([0]Transform2D{}).Elem():
			vtype = TypeTransform2D
		case reflect.TypeOf([0]Vector4{}).Elem():
			vtype = TypeVector4
		case reflect.TypeOf([0]Vector4i{}).Elem():
			vtype = TypeVector4i
		case reflect.TypeOf([0]Plane{}).Elem():
			vtype = TypePlane
		case reflect.TypeOf([0]Quaternion{}).Elem():
			vtype = TypeQuaternion
		case reflect.TypeOf([0]AABB{}).Elem():
			vtype = TypeAABB
		case reflect.TypeOf([0]Basis{}).Elem():
			vtype = TypeBasis
		case reflect.TypeOf([0]Transform3D{}).Elem():
			vtype = TypeTransform3D
		case reflect.TypeOf([0]Projection{}).Elem():
			vtype = TypeProjection
		case reflect.TypeOf([0]Color{}).Elem():
			vtype = TypeColor
		case reflect.TypeOf([0]StringName{}).Elem():
			vtype = TypeStringName
		case reflect.TypeOf([0]NodePath{}).Elem():
			vtype = TypeNodePath
		case reflect.TypeOf([0]RID{}).Elem():
			vtype = TypeRID
		case reflect.TypeOf([0]Object{}).Elem():
			vtype = TypeObject
		case reflect.TypeOf([0]Callable{}).Elem():
			vtype = TypeCallable
		case reflect.TypeOf([0]Dictionary{}).Elem():
			vtype = TypeDictionary
		case reflect.TypeOf([0]Array{}).Elem():
			vtype = TypeArray
		case reflect.TypeOf([0]PackedByteArray{}).Elem():
			vtype = TypePackedByteArray
		case reflect.TypeOf([0]PackedInt32Array{}).Elem():
			vtype = TypePackedInt32Array
		case reflect.TypeOf([0]PackedInt64Array{}).Elem():
			vtype = TypePackedInt64Array
		case reflect.TypeOf([0]PackedFloat32Array{}).Elem():
			vtype = TypePackedFloat32Array
		case reflect.TypeOf([0]PackedFloat64Array{}).Elem():
			vtype = TypePackedFloat64Array
		case reflect.TypeOf([0]PackedStringArray{}).Elem():
			vtype = TypePackedStringArray
		case reflect.TypeOf([0]PackedVector2Array{}).Elem():
			vtype = TypePackedVector2Array
		case reflect.TypeOf([0]PackedVector3Array{}).Elem():
			vtype = TypePackedVector3Array
		case reflect.TypeOf([0]PackedColorArray{}).Elem():
			vtype = TypePackedColorArray
		case reflect.TypeOf([0]unsafe.Pointer{}).Elem():
			vtype = TypeNil
		case reflect.TypeOf([0]*ScriptLanguageExtensionProfilingInfo{}).Elem():
			vtype = TypeNil
		default:
			switch {
			case rtype.Implements(reflect.TypeOf([0]IsClass{}).Elem()):
				vtype = TypeObject
			case rtype.Implements(reflect.TypeOf([0]IsSignal{}).Elem()):
				vtype = TypeSignal
			default:
				vtype = TypeDictionary
			}
		}
		return vtype, true
	default:
		return TypeNil, false
	}
}
