package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func convertType(gdType string) string {
	switch gdType {
	case "int":
		return "int64"
	case "float":
		return "float64"
	case "String":
		return "string"
	case "enum::Error":
		return "error"
	case "const uint8_t **":
		return "*[]byte"
	case "const void*", "const uint8_t*", "const uint8_t *", "PackedByteArray":
		return "[]byte"
	case "PackedStringArray":
		return "[]string"
	case "PackedInt32Array":
		return "[]int32"
	case "PackedInt64Array":
		return "[]int64"
	case "PackedFloat32Array":
		return "[]float32"
	case "PackedFloat64Array":
		return "[]float64"
	case "PackedVector2Array":
		return "[]Vector2"
	case "PackedVector3Array":
		return "[]Vector3"
	case "PackedColorArray":
		return "[]Color"
	case "StringName":
		return "string"
	case "Variant":
		return "any"
	case "float*":
		return "*float64"
	case "int32_t*":
		return "*int32"
	case "void*", "uint8_t*":
		return "[]byte"
	default:
		if strings.HasPrefix(gdType, "const") {
			gdType = strings.TrimPrefix(gdType, "const")
		}
		if strings.HasSuffix(gdType, "*") {
			return "*" + gdType[:len(gdType)-1]
		}
		if strings.HasPrefix(gdType, "enum::") {
			return strings.TrimPrefix(strings.Replace(gdType, ".", "", -1), "enum::")
		}
		if strings.HasPrefix(gdType, "bitfield::") {
			return strings.TrimPrefix(strings.Replace(gdType, ".", "", -1), "bitfield::")
		}
		return gdType
	}
}

func convertName(fnName string) string {
	fnName = strings.ToLower(fnName)

	joins := []string{}
	split := strings.Split(fnName, "_")
	for _, word := range split {
		joins = append(joins, strings.Title(word))
	}
	/*if joins[0] == "Get" {
		backup := joins
		joins = joins[1:]

		if len(joins) == 0 {
			joins = backup
		} else {
			if _, err := strconv.Atoi(joins[0]); err == nil {
				joins = backup
			}
		}
	}*/
	return strings.Join(joins, "")
}

func fixReserved(name string) string {
	switch name {
	case "type":
		return "atype"
	case "range":
		return "arange"
	case "default":
		return "def"
	case "func":
		return "fn"
	case "interface":
		return "intf"
	case "map":
		return "mapping"
	case "var":
		return "v"
	case "string":
		return "s"
	default:
		return name
	}
}

func genEnum(code io.Writer, prefix string, enum Enum) {
	name := prefix + strings.Replace(enum.Name, ".", "", -1)

	fmt.Fprintln(code)
	fmt.Fprintf(code, "type %v int64\n", name)
	if len(enum.Values) > 0 {
		fmt.Fprintln(code)
		fmt.Fprintf(code, "const (\n")
		for _, value := range enum.Values {
			n := prefix + convertName(value.Name)
			if n == name {
				n += "Default"
			}
			fmt.Fprintf(code, "\t%v %v = %v\n", n, name, value.Value)
		}
		fmt.Fprintf(code, ")\n")
	}
}

func generate() error {
	file, err := os.Open("./api.json")
	if os.IsNotExist(err) {
		req, err := http.NewRequest("GET", "https://raw.githubusercontent.com/godotengine/godot-headers/master/extension_api.json", nil)
		if err != nil {
			return err
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}

		file, err = os.Create("./api.json")
		if err != nil {
			return err
		}

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return err
		}

		file.Seek(0, 0)
		resp.Body.Close()
	}

	var spec Specification
	if err := json.NewDecoder(file).Decode(&spec); err != nil {
		return err
	}

	var virtuals []Method

	code, err := os.Create("./api.go")
	if err != nil {
		return err
	}

	var exceptions = map[string]bool{
		"GDScriptNativeClass": true,
		"JavaClassWrapper":    true,
		"JavaScript":          true,
	}

	fmt.Fprintln(code, `// Code generated by the generate package DO NOT EDIT`)
	fmt.Fprintln(code, `package gd`)
	fmt.Fprintln(code)
	fmt.Fprintln(code, `import "github.com/readykit/gd/gdnative"`)
	fmt.Fprintln(code, `import "reflect"`)

	for _, enum := range spec.GlobalEnums {
		genEnum(code, "", enum)
	}

	for _, class := range spec.BuiltinClasses {
		for _, enum := range class.Enums {
			genEnum(code, class.Name, enum)
		}
	}

	for _, class := range spec.Classes {
		for _, enum := range class.Enums {
			genEnum(code, class.Name, enum)
		}

		fmt.Fprintln(code)
		fmt.Fprintf(code, "type %v gdnative.Object\n", class.Name)
		fmt.Fprintf(code, `func (%v) class() string { return "%v\000" }`+"\n", class.Name, class.Name)
		fmt.Fprintln(code)
		if class.Inherits != "" {
			fmt.Fprintf(code, `func (gdClass %v) %[2]v() %[2]v { return %[2]v(gdClass) }`+"\n", class.Name, class.Inherits)
			fmt.Fprintln(code)
		}

		var methodCount = 0
		var virtualCount = 0
		for _, method := range class.Methods {
			if !method.IsVirtual {
				methodCount++
			} else {
				virtualCount++
			}
		}

		if methodCount > 0 {
			fmt.Fprintf(code, "var method%v [%d]gdnative.Method\n", class.Name, methodCount)
		}

		// check if a Go type implements that method and returns the method that implements the
		// named virtual.
		fmt.Fprintf(code, "func (gdClass %v) virtual(rtype reflect.Type, name string) (method reflect.Method, ok bool) {\n", class.Name)
		if virtualCount > 0 {
			fmt.Fprintln(code, "\tswitch name {")
			for _, method := range class.Methods {
				if !method.IsVirtual {
					continue
				}

				result := convertType(method.ReturnValue.Type)
				virtuals = append(virtuals, method)

				fmt.Fprintf(code, "\tcase \""+method.Name+"\":\n")
				fmt.Fprintf(code, "\t\tif rtype.Implements(reflect.TypeOf([0]interface{ %v(", convertName(method.Name))
				for i, arg := range method.Arguments {
					fmt.Fprintf(code, "%v %v", fixReserved(arg.Name), convertType(arg.Type))
					if i < len(method.Arguments)-1 {
						fmt.Fprint(code, ", ")
					}
				}
				fmt.Fprintf(code, ") %v }{}).Elem()) {\n", result)
				fmt.Fprintf(code, "\t\t\treturn rtype.MethodByName(`%s`)\n", convertName(method.Name))
				fmt.Fprintln(code, "\t\t}")
				fmt.Fprintln(code, "\t\treturn")
			}
			fmt.Fprintln(code, "\t}")
		}
		if class.Inherits != "" {
			fmt.Fprintf(code, "\treturn gdClass.%v().virtual(rtype, name)\n", class.Inherits)
		} else {
			fmt.Fprintln(code, "\treturn")
		}
		fmt.Fprintln(code, "}")

		var i int
		for _, method := range class.Methods {
			if method.IsVirtual {
				continue
			}

			result := convertType(method.ReturnValue.Type)

			fmt.Fprintf(code, "func (gdClass %v) %v(", class.Name, convertName(method.Name))
			for i, arg := range method.Arguments {
				fmt.Fprintf(code, "%v %v", fixReserved(arg.Name), convertType(arg.Type))
				if i < len(method.Arguments)-1 {
					fmt.Fprint(code, ", ")
				}
			}
			fmt.Fprintf(code, ") %v { ", result)
			if result != "" {
				fmt.Fprintf(code, "return gdnative.Return[%v](gdnative.Object(gdClass), method%v[%d]", result, class.Name, i)
			} else {
				fmt.Fprintf(code, "gdnative.Call(gdnative.Object(gdClass), method%v[%d]", class.Name, i)
			}
			for _, arg := range method.Arguments {
				fmt.Fprint(code, ", ")
				fmt.Fprintf(code, "%v", fixReserved(arg.Name))
			}
			fmt.Fprintf(code, ") }\n")
			i++
		}
	}

	fmt.Fprintln(code)
	fmt.Fprintf(code, `func init() {`)
	fmt.Fprintf(code, "\tgdnative.OnLoad(func() {")
	for _, class := range spec.Classes {
		var i int
		for _, method := range class.Methods {
			if _, ok := exceptions[class.Name]; ok {
				continue
			}

			if method.IsVirtual {
				continue
			}
			fmt.Fprintf(code, "\t\t"+`method%v[%d] = gdnative.MethodOf("%v\000", "%v\000", %v)`+"\n", class.Name, i, class.Name, method.Name, method.Hash)
			i++
		}
	}
	fmt.Fprintf(code, "\t})\n}\n\n")

	code.Close()

	// Create Virtual C functions
	// required because C doesn't support closures and the Godot GDNativeExtensionClassCallVirtual
	// doesn't include a userdata field.

	code, err = os.Create("./gdnative/gdnative_virtuals.h")
	if err != nil {
		return err
	}

	fmt.Fprintln(code, `extern uint8_t goClassGetVirtual(void *p_userdata, const char *p_name);`)
	fmt.Fprintf(code, "extern void goMethodCallDirect(uintptr_t userdata, uintptr_t instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret);\n")

	for i := 1; i < 256; i++ {
		fmt.Fprintf(code, "void goClassCallVirtual%d(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {", i)
		fmt.Fprintf(code, "goMethodCallDirect((uintptr_t)%v, (uintptr_t)p_instance,  p_args, r_ret);", i)
		fmt.Fprintf(code, "}\n")
	}

	fmt.Fprintf(code, "GDNativeExtensionClassCallVirtual get_virtual_func(void *p_userdata, const char *p_name) {\n")
	fmt.Fprintf(code, "switch (goClassGetVirtual(p_userdata, p_name)) {\n")
	for i := 1; i < 256; i++ {
		fmt.Fprintf(code, "\tcase %v: return goClassCallVirtual%d;\n", i, i)
	}
	fmt.Fprintf(code, "\tdefault: return NULL;")
	fmt.Fprintf(code, "}\n}")

	return code.Close()
}

func main() {
	if err := generate(); err != nil {
		log.Fatal(err)
	}
}
