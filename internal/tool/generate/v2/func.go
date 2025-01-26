package main

import (
	"fmt"
	"io"
	"strings"

	"graphics.gd/internal/gdjson"
	"graphics.gd/internal/gdtype"
)

type callType int

const (
	callDefault callType = iota
	callBuiltin
	callUtility
)

func (classDB ClassDB) signalCall(w io.Writer, class gdjson.Class, signal gdjson.Signal, singleton bool) {
	if singleton {
		fmt.Fprintf(w, "\nfunc On%v(cb func(", convertName(signal.Name))
	} else {
		fmt.Fprintf(w, "\nfunc (self Instance) On%v(cb func(", convertName(signal.Name))
	}
	for i, arg := range signal.Arguments {
		if i > 0 {
			fmt.Fprint(w, ", ")
		}
		fmt.Fprintf(w, "%v %v", fixReserved(arg.Name), classDB.convertTypeSimple(class, "", arg.Meta, arg.Type))
	}
	fmt.Fprint(w, ")) {\n\t")
	fmt.Fprintf(w, `self[0].AsObject()[0].Connect(gd.NewStringName("%s"), gd.NewCallable(cb), 0)`, signal.Name)
	fmt.Fprint(w, "\n}\n\n")
}

func (classDB ClassDB) simpleCall(w io.Writer, class gdjson.Class, method gdjson.Method, singleton bool) {
	if method.IsVararg {
		return
	}
	switch class.Name {
	case "Float", "Int", "Vector2", "Vector2i", "Rect2", "Rect2i", "Vector3", "Vector3i",
		"Transform2D", "Vector4", "Vector4i", "Plane", "Quaternion", "AABB", "Basis", "Transform3D",
		"RID", "Projection", "Color":
		return
	}
	if method.Description != "" {
		fmt.Fprintln(w, "\n/*")
		fmt.Fprint(w, method.Description)
		fmt.Fprintln(w, "\n*/")
	}
	if method.IsVirtual {
		classDB.simpleVirtualCall(w, class, method)
		return
	}
	resultSimple := classDB.convertTypeSimple(class, class.Name+"."+method.Name+".", method.ReturnValue.Meta, method.ReturnValue.Type)
	resultExpert := classDB.convertType(class.Name, method.ReturnValue.Meta, method.ReturnValue.Type)
	if singleton || method.IsStatic {
		fmt.Fprintf(w, "func %v(", convertName(method.Name))
	} else {
		fmt.Fprintf(w, "func (self Instance) %v(", convertName(method.Name))
	}
	for i, arg := range method.Arguments {
		if arg.DefaultValue == nil {
			if i > 0 {
				fmt.Fprint(w, ", ")
			}
			fmt.Fprintf(w, "%v %v", fixReserved(arg.Name), classDB.convertTypeSimple(class, class.Name+"."+method.Name+"."+arg.Name, arg.Meta, arg.Type))
		}
	}
	fmt.Fprint(w, ") ")
	if method.ReturnValue.Type != "" {
		fmt.Fprintf(w, "%v ", resultSimple)
	}
	fmt.Fprintf(w, "{ //gd:%s.%s\n\t", class.Name, method.Name)
	if singleton {
		fmt.Fprintf(w, "once.Do(singleton)\n\t")
	}
	if method.IsStatic {
		fmt.Fprintf(w, "self := Instance{}\n")
	}
	if method.ReturnValue.Type != "" {
		fmt.Fprintf(w, "return %s(", resultSimple)
	}
	var call strings.Builder
	fmt.Fprintf(&call, "class(self).%v(", convertName(method.Name))
	for i, arg := range method.Arguments {
		if i > 0 {
			fmt.Fprint(&call, ", ")
		}
		val := fixReserved(arg.Name)
		if arg.DefaultValue != nil {
			switch arg.Type {
			case "Array":
				val = "Array.Nil"
			case "Callable":
				val = "Callable.Nil"
			case "Dictionary":
				val = "Dictionary.Nil"
			default:
				val = *arg.DefaultValue
				val = strings.TrimPrefix(val, "&")
				if val == "null" || val == "[]" || val == "{}" || strings.HasSuffix(val, "()") || strings.HasSuffix(val, "[])") {
					if arg.Type == "Callable" {
						val = "nil"
					} else {
						val = "([1]" + classDB.convertTypeSimple(class, "", arg.Meta, arg.Type) + "{}[0])"
					}
				} else {
					if strings.Contains(val, "(") {
						switch {
						case strings.HasPrefix(val, "Rect2("), strings.HasPrefix(val, "Rect2i("), strings.HasPrefix(val, "Transform2D("),
							strings.HasPrefix(val, "Transform3D("):
							val = "gd.New" + val
						case strings.HasPrefix(val, "StringName(\""):
							val = strings.TrimSuffix(strings.TrimPrefix(val, "StringName(\""), "\")")
						case strings.HasPrefix(val, "NodePath(\""):
							val = strings.TrimSuffix(strings.TrimPrefix(val, "NodePath(\""), "\")")
							if val == "" {
								val = `""`
							}
						default:
							val = "gd." + strings.ReplaceAll(strings.ReplaceAll(val, "(", "{"), ")", "}")
						}
					}
				}
			}
			if gdtype.Name(classDB.convertType(class.Name, arg.Meta, arg.Type)) == "gd.Variant" {
				val = `gd.NewVariant(` + val + `)`
			}
		}
		fmt.Fprint(&call, gdtype.Name(classDB.convertType(class.Name, arg.Meta, arg.Type)).ConvertToSimple(val))
	}
	fmt.Fprint(&call, ")")
	fmt.Fprint(w, gdtype.Name(resultExpert).ConvertToGo(call.String(), resultSimple))
	if method.ReturnValue.Type != "" {
		fmt.Fprint(w, ")")
	}
	fmt.Fprintf(w, "\n}\n")
}

func (classDB ClassDB) simpleVirtualCall(w io.Writer, class gdjson.Class, method gdjson.Method) {
	resultSimple := classDB.convertTypeSimple(class, "", method.ReturnValue.Meta, method.ReturnValue.Type)
	resultExpert := classDB.convertType(class.Name, method.ReturnValue.Meta, method.ReturnValue.Type)
	_, needsLifetime := gdtype.Name(resultExpert).IsPointer()
	if method.IsStatic {
		needsLifetime = true
	}
	fmt.Fprintf(w, "func (Instance) %s(impl func(ptr unsafe.Pointer", method.Name)
	for _, arg := range method.Arguments {
		fmt.Fprint(w, ", ")
		fmt.Fprintf(w, "%v %v", fixReserved(arg.Name), classDB.convertTypeSimple(class, "", arg.Meta, arg.Type))
	}
	fmt.Fprintf(w, ") %v) (cb gd.ExtensionClassCallVirtualFunc) {\n", resultSimple)
	fmt.Fprintf(w, "\treturn func(class any, p_args gd.Address, p_back gd.Address) {\n")
	for i, arg := range method.Arguments {
		var expert = classDB.convertType(class.Name, arg.Meta, arg.Type)
		pointerKind, argIsPtr := gdtype.Name(expert).IsPointer()
		if !argIsPtr {
			pointerKind = expert
		}
		fmt.Fprintf(w, "\t\tvar %v = %v\n", fixReserved(arg.Name), gdtype.Name(expert).LoadFromRawPointerValue(
			fmt.Sprintf("gd.UnsafeGet[%v](p_args,%d)", pointerKind, i),
		))
		if argIsPtr {
			fmt.Fprintf(w, "\t\tdefer %s\n", gdtype.Name(expert).EndPointer(fixReserved(arg.Name)))
		}
	}
	fmt.Fprintf(w, "\t\tself := reflect.ValueOf(class).UnsafePointer()\n")
	if resultSimple != "" {
		fmt.Fprintf(w, "\t\tret := ")
	}
	fmt.Fprintf(w, "impl(self")
	for _, arg := range method.Arguments {
		simple := classDB.convertTypeSimple(class, "", arg.Meta, arg.Type)
		fmt.Fprint(w, ", ")
		fmt.Fprintf(w, "%v", gdtype.Name(classDB.convertType(class.Name, arg.Meta, arg.Type)).ConvertToGo(fixReserved(arg.Name), simple))
	}
	fmt.Fprintf(w, ")\n")
	if resultSimple != "" {
		ret := gdtype.Name(resultExpert).ToUnderlying(gdtype.Name(resultExpert).ConvertToSimple("ret"))
		if needsLifetime {
			fmt.Fprintf(w, "ptr, ok := %s\n", gdtype.Name(resultExpert).EndPointer(ret))
			fmt.Fprintf(w, "\n\t\tif !ok {\n")
			fmt.Fprintf(w, "\t\t\treturn\n")
			fmt.Fprintf(w, "\t\t}\n")
			ret = "ptr"
		}
		fmt.Fprintf(w, "\t\tgd.UnsafeSet(p_back, %s)\n", ret)
	}
	fmt.Fprintf(w, "\t}\n")
	fmt.Fprintf(w, "}\n")
	return
}

func (classDB ClassDB) methodCall(w io.Writer, pkg string, class gdjson.Class, method gdjson.Method, ctype callType) {
	if ctype == callDefault && method.IsVararg {
		return
	}
	switch class.Name {
	case "Float", "Int", "Vector2", "Vector2i", "Rect2", "Rect2i", "Vector3", "Vector3i",
		"Transform2D", "Vector4", "Vector4i", "Plane", "Quaternion", "AABB", "Basis", "Transform3D",
		"RID", "Projection", "Color":
		return
	}
	result := classDB.convertType(class.Name, method.ReturnValue.Meta, method.ReturnValue.Type)
	if method.ReturnType != "" {
		result = classDB.convertType(class.Name, "", method.ReturnType)
	}
	ptrKind, isPtr := gdtype.Name(result).IsPointer()

	prefix := ""
	if pkg != "internal" {
		prefix = "gd."
	}

	fmt.Fprintln(w)
	if method.Description != "" {
		fmt.Fprintln(w, "/*")
		fmt.Fprint(w, method.Description)
		fmt.Fprintln(w, "\n*/")
	}

	if method.IsVirtual {
		fmt.Fprintf(w, "func (class) %s(impl func(ptr unsafe.Pointer", method.Name)
		for _, arg := range method.Arguments {
			fmt.Fprint(w, ", ")
			fmt.Fprintf(w, "%v %v", fixReserved(arg.Name), classDB.convertType(class.Name, arg.Meta, arg.Type))
		}
		fmt.Fprintf(w, ") %v) (cb "+prefix+"ExtensionClassCallVirtualFunc) {\n", result)
		fmt.Fprintf(w, "\treturn func(class any, p_args "+prefix+"Address, p_back "+prefix+"Address) {\n")
		for i, arg := range method.Arguments {
			var argType = classDB.convertType(class.Name, arg.Meta, arg.Type)
			pointerKind, argIsPtr := gdtype.Name(argType).IsPointer()
			if !argIsPtr {
				pointerKind = argType
			}
			fmt.Fprintf(w, "\t\tvar %v = %v\n", fixReserved(arg.Name), gdtype.Name(argType).LoadFromRawPointerValue(
				fmt.Sprintf("gd.UnsafeGet[%v](p_args,%d)", pointerKind, i),
			))
			if argIsPtr {
				fmt.Fprintf(w, "\t\tdefer %s\n", gdtype.Name(argType).EndPointer(fixReserved(arg.Name)))
			}
		}
		fmt.Fprintf(w, "\t\tself := reflect.ValueOf(class).UnsafePointer()\n")
		if result != "" {
			fmt.Fprintf(w, "\t\tret := ")
		}
		fmt.Fprintf(w, "impl(self")
		for _, arg := range method.Arguments {
			fmt.Fprint(w, ", ")
			fmt.Fprintf(w, "%v", fixReserved(arg.Name))
		}
		fmt.Fprintf(w, ")\n")
		if result != "" {
			ret := gdtype.Name(result).ToUnderlying("ret")
			if isPtr {
				fmt.Fprintf(w, "ptr, ok := %s\n", gdtype.Name(result).EndPointer(ret))
				fmt.Fprintf(w, "\n\t\tif !ok {\n")
				fmt.Fprintf(w, "\t\t\treturn\n")
				fmt.Fprintf(w, "\t\t}\n")
				ret = "ptr"
			}

			fmt.Fprintf(w, "\t\t"+prefix+"UnsafeSet(p_back, %s)\n", ret)
		}
		fmt.Fprintf(w, "\t}\n")
		fmt.Fprintf(w, "}\n")
		return
	}
	if ctype == callBuiltin && strings.HasPrefix(class.Name, "Packed") {
		fmt.Fprintf(w, "//go:nosplit\nfunc (self *class) %v(", convertName(method.Name))
	} else {
		fmt.Fprintf(w, "//go:nosplit\nfunc (self class) %v(", convertName(method.Name))
	}

	if method.Name == "select" {
		method.Name = "select_"
	}
	if method.Name == "map" {
		method.Name = "map_"
	}

	for i, arg := range method.Arguments {
		if i > 0 {
			fmt.Fprint(w, ", ")
		}
		fmt.Fprintf(w, "%v %v", fixReserved(arg.Name), classDB.convertType(class.Name, arg.Meta, arg.Type))
	}
	if method.IsVararg {
		if len(method.Arguments) > 0 {
			fmt.Fprint(w, ", ")
		}
		fmt.Fprintf(w, "args ..."+prefix+"Variant")
	}
	fmt.Fprintf(w, ") %v { //gd:%s.%s\n", result, class.Name, method.Name)
	fmt.Fprintf(w, "\tvar frame = callframe.New()\n")
	for _, arg := range method.Arguments {
		_, ok := classDB[arg.Type]
		if ok {
			switch semantics := gdjson.ClassMethodOwnership[class.Name][method.Name][arg.Name]; semantics {
			case gdjson.OwnershipTransferred, gdjson.LifetimeBoundToClass:
				fmt.Fprintf(w, "\tcallframe.Arg(frame, gd.PointerWithOwnershipTransferredToGodot(%v[0].AsObject()[0]))\n", fixReserved(arg.Name))
			case gdjson.RefCountedManagement, gdjson.IsTemporaryReference, gdjson.MustAssertInstanceID, gdjson.ReversesTheOwnership:
				fmt.Fprintf(w, "\tcallframe.Arg(frame, pointers.Get(%v[0])[0])\n", fixReserved(arg.Name))
			default:
				panic("unknown ownership: " + fmt.Sprint(semantics))
			}
			continue
		}
		argType := classDB.convertType(class.Name, arg.Meta, arg.Type)
		fmt.Fprint(w, gdtype.Name(argType).LoadOntoCallFrame(fixReserved(arg.Name)))
	}
	if method.IsVararg {
		fmt.Fprintf(w, "\tfor _, arg := range args {\n")
		fmt.Fprintf(w, "\t\tcallframe.Arg(frame, pointers.Get(arg))\n")
		fmt.Fprintf(w, "\t}\n")
	}
	if isPtr {
		fmt.Fprintf(w, "\tvar r_ret = callframe.Ret[%v](frame)\n", ptrKind)
	} else {
		if result != "" {
			fmt.Fprintf(w, "\tvar r_ret = callframe.Ret[%v](frame)\n", result)
		} else {
			fmt.Fprintf(w, "\tvar r_ret = callframe.Nil\n")
		}
	}
	if method.IsVararg {
		fmt.Fprintf(w, "\tif len(args) > 0 { panic(`varargs not supported for class methods yet`); }\n")
	}
	fmt.Fprintf(w, "\tgd.Global.Object.MethodBindPointerCall(gd.Global.Methods.%v.Bind_%v, self.AsObject(), frame.Array(0), r_ret.Addr())\n", class.Name, method.Name)

	if isPtr {
		_, ok := classDB[strings.TrimPrefix(result, "[1]gdclass.")]
		if ok || result == "gd.Object" {
			if result == "gd.Object" {
				switch semantics := gdjson.ClassMethodOwnership[class.Name][method.Name]["return value"]; semantics {
				case gdjson.RefCountedManagement, gdjson.OwnershipTransferred:
					fmt.Fprintf(w, "\tvar ret = "+prefix+"PointerWithOwnershipTransferredToGo[gd.Object](r_ret.Get())\n")
				case gdjson.LifetimeBoundToClass:
					fmt.Fprintf(w, "\tvar ret = "+prefix+"PointerLifetimeBoundTo[gd.Object](self.AsObject(), r_ret.Get())\n")
				case gdjson.MustAssertInstanceID:
					fmt.Fprintf(w, "\tvar ret = "+prefix+"PointerMustAssertInstanceID[gd.Object](r_ret.Get())\n")
				default:
					panic("unknown ownership: " + fmt.Sprint(semantics))
				}
			} else {
				switch semantics := gdjson.ClassMethodOwnership[class.Name][method.Name]["return value"]; semantics {
				case gdjson.RefCountedManagement, gdjson.OwnershipTransferred:
					fmt.Fprintf(w, "\tvar ret = [1]gdclass.%s{"+prefix+"PointerWithOwnershipTransferredToGo[gdclass.%[1]s](r_ret.Get())}\n", method.ReturnValue.Type)
				case gdjson.LifetimeBoundToClass:
					fmt.Fprintf(w, "\tvar ret = [1]gdclass.%s{"+prefix+"PointerLifetimeBoundTo[gdclass.%[1]s](self.AsObject(), r_ret.Get())}\n", method.ReturnValue.Type)
				case gdjson.MustAssertInstanceID:
					fmt.Fprintf(w, "\tvar ret = [1]gdclass.%s{"+prefix+"PointerMustAssertInstanceID[gdclass.%[1]s](r_ret.Get())}\n", method.ReturnValue.Type)
				case gdjson.IsTemporaryReference:
					fmt.Fprintf(w, "\tvar ret = [1]gdclass.%s{"+prefix+"PointerBorrowedTemporarily[gdclass.%[1]s](r_ret.Get())}\n", method.ReturnValue.Type)
				default:
					panic("unknown ownership: " + fmt.Sprint(semantics))
				}
			}
		} else {
			fmt.Fprintf(w, "\tvar ret = %s\n", gdtype.Name(result).LoadFromRawPointerValue("r_ret.Get()"))
		}
	} else if result != "" {
		fmt.Fprintf(w, "\tvar ret = r_ret.Get()\n")
	}
	fmt.Fprintf(w, "\tframe.Free()\n")
	if result != "" {
		if strings.HasPrefix(result, "ArrayOf") || strings.HasPrefix(result, "gd.ArrayOf") {
			result = strings.ReplaceAll(result, "gd.ArrayOf", "gd.TypedArray")
			result = strings.ReplaceAll(result, "ArrayOf", "TypedArray")
			fmt.Fprintf(w, "\treturn %s(ret)\n", result)
		} else {
			fmt.Fprintf(w, "\treturn ret\n")
		}
	}
	/*if result != "" {
		fmt.Fprintf(w, "\tvar ret %v\n", result)
		fmt.Fprintf(w, "\treturn ret\n")
	}*/
	fmt.Fprintf(w, "}")
}
