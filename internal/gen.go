//go:build generate

//go:generate go run -tags generate .
//go:generate go fmt ./

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
)

type Enum struct {
	Name   string `json:"name"`
	Values []struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	} `json:"values"`
}

type Method struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsConst     bool   `json:"is_const"`
	IsVararg    bool   `json:"is_vararg"`
	IsStatic    bool   `json:"is_static"`
	IsVirtual   bool   `json:"is_virtual"`
	Hash        int64  `json:"hash"`
	ReturnType  string `json:"return_type"` // builtin class
	ReturnValue struct {
		Type string `json:"type"`
		Meta string `json:"meta"`
	} `json:"return_value,omitempty"`
	Arguments []struct {
		Name         string `json:"name"`
		Type         string `json:"type"`
		Meta         string `json:"meta"`
		DefaultValue string `json:"default_value,omitempty"`
	} `json:"arguments,omitempty"`
}

type Class struct {
	Name    string `json:"name"`
	Rename  string `json:"rename,omitempty"`
	Package string `json:"package,omitempty"`
	IsEnum  bool   `json:"is_enum"`

	Description    string `json:"description"`
	IsRefcounted   bool   `json:"is_refcounted"`
	IsInstantiable bool   `json:"is_instantiable"`

	IsKeyed   bool `json:"is_keyed"` // builtin class
	Operators []struct {
		Name       string `json:"name"`
		RightType  string `json:"right_type,omitempty"`
		ReturnType string `json:"return_type"`
	} `json:"operators"` // builtin class
	Constructors []struct {
		Index     int `json:"index"`
		Arguments []struct {
			Name string `json:"name"`
			Type string `json:"type"`
			Meta string `json:"meta"`
		} `json:"arguments,omitempty"`
	} `json:"constructors"` // builtin class
	HasDestructor      bool   `json:"has_destructor"`                 // builtin class
	IndexingReturnType string `json:"indexing_return_type,omitempty"` // builtin class

	Inherits string   `json:"inherits,omitempty"`
	APIType  string   `json:"api_type"`
	Enums    []Enum   `json:"enums,omitempty"`
	Methods  []Method `json:"methods,omitempty"`

	Members []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"members,omitempty"` // builtin class

	Signals []struct {
		Name      string `json:"name"`
		Arguments []struct {
			Name string `json:"name"`
			Type string `json:"type"`
			Meta string `json:"meta"`
		} `json:"arguments,omitempty"`
	} `json:"signals,omitempty"`
	Properties []struct {
		Type   string `json:"type"`
		Name   string `json:"name"`
		Setter string `json:"setter"`
		Getter string `json:"getter"`
		Index  int    `json:"index"`
	} `json:"properties,omitempty"`
	Constants []struct {
		Name  string `json:"name"`
		Type  string `json:"type"` // builtin class
		Value any    `json:"value"`
	} `json:"constants,omitempty"`

	singleton, builtin bool // internal
}

/*
Specification of the Godot Extension API.
Created with https://mholt.github.io/json-to-go/
*/
type Specification struct {
	Header struct {
		VersionMajor    int    `json:"version_major"`
		VersionMinor    int    `json:"version_minor"`
		VersionPatch    int    `json:"version_patch"`
		VersionStatus   string `json:"version_status"`
		VersionBuild    string `json:"version_build"`
		VersionFullName string `json:"version_full_name"`
	} `json:"header"`
	BuiltinClassSizes []struct {
		BuildConfiguration string `json:"build_configuration"`
		Sizes              []struct {
			Name string `json:"name"`
			Size int    `json:"size"`
		} `json:"sizes"`
	} `json:"builtin_class_sizes"`
	BuiltinClassMemberOffsets []struct {
		BuildConfiguration string `json:"build_configuration"`
		Classes            []struct {
			Name    string `json:"name"`
			Members []struct {
				Member string `json:"member"`
				Offset int    `json:"offset"`
			} `json:"members"`
		} `json:"classes"`
	} `json:"builtin_class_member_offsets"`
	GlobalConstants  []interface{} `json:"global_constants"`
	GlobalEnums      []Enum        `json:"global_enums"`
	UtilityFunctions []struct {
		Name       string `json:"name"`
		ReturnType string `json:"return_type,omitempty"`
		Category   string `json:"category"`
		IsVararg   bool   `json:"is_vararg"`
		Hash       int    `json:"hash"`
		Arguments  []struct {
			Name string `json:"name"`
			Type string `json:"type"`
			Meta string `json:"meta"`
		} `json:"arguments,omitempty"`
	} `json:"utility_functions"`
	BuiltinClasses []Class `json:"builtin_classes"`
	Classes        []Class `json:"classes"`
	Singletons     []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"singletons"`
	NativeStructures []struct {
		Name   string `json:"name"`
		Format string `json:"format"`
	} `json:"native_structures"`
}

// Load the specification, either from a local file or by downloading
// it from the Godot Github repository.
func (spec *Specification) Load() error {
	file, err := os.Open("../extension_api.json")
	if os.IsNotExist(err) {
		req, err := http.NewRequest("GET", "https://raw.githubusercontent.com/godotengine/godot-headers/master/extension_api.json", nil)
		if err != nil {
			return err
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}

		file, err = os.Create("../extension_api.json")
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

	if err := json.NewDecoder(file).Decode(&spec); err != nil {
		return err
	}

	return nil
}

func hasNative(gdType string) bool {
	switch gdType {
	case "int", "float", "String", "bool",
		"PackedStringArray", "PackedInt32Array", "PackedInt64Array",
		"PackedFloat32Array", "PackedFloat64Array", "PackedVector2Array",
		"PackedVector3Array", "PackedColorArray", "Variant":
		return true
	}
	return false
}

func (classDB ClassDB) convertType(pkg, meta string, gdType string) string {
	maybeInternal := func(name string) string {
		if pkg != "internal" {
			return "gd." + name
		}
		return name
	}
	if strings.HasPrefix(gdType, "typedarray::") {
		gdType = strings.TrimPrefix(gdType, "typedarray::")
		return maybeInternal("ArrayOf[" + classDB.convertType(pkg, meta, gdType) + "]")
	}
	switch gdType {
	case "int", "Int":
		return maybeInternal("Int")
	case "float", "Float":
		return maybeInternal("Float")
	case "bool", "Bool":
		return "bool"
	case "String":
		return maybeInternal("String")
	case "StringName":
		return maybeInternal("StringName")
	case "enum::GDExtension.InitializationLevel":
		return maybeInternal("GDExtensionInitializationLevel")
	case "enum::GDExtensionManager.LoadStatus":
		return "GDExtensionManagerLoadStatus"
	case "PackedStringArray", "PackedInt32Array", "PackedInt64Array", "PackedFloat32Array",
		"PackedFloat64Array", "PackedVector2Array", "PackedVector3Array", "PackedColorArray", "PackedByteArray",
		"Vector2", "Vector2i", "Rect2", "Rect2i", "Vector3", "Vector3i", "Transform2D", "Vector4", "Vector4i",
		"Plane", "Quaternion", "AABB", "Basis", "Transform3D", "Projection", "Color", "NodePath", "RID", "Object",
		"Callable", "Signal", "Dictionary", "Array":
		return maybeInternal(gdType)
	case "Variant":
		return maybeInternal("Variant")
	case "enum::Variant.Type":
		return maybeInternal("VariantType")

	// strange C++ cases
	case "enum::Error":
		return "int64"
	case "const uint8_t **":
		return "unsafe.Pointer"
	case "const void*", "const uint8_t*", "const uint8_t *":
		return "unsafe.Pointer"
	case "float*":
		return "*float64"
	case "int32_t*":
		return "*int32"
	case "void*", "uint8_t*":
		return "unsafe.Pointer"

	default:
		gdType = strings.TrimPrefix(gdType, "const")

		if strings.HasSuffix(gdType, "*") {
			return "*" + gdType[:len(gdType)-1]
		}

		gdType = strings.TrimPrefix(strings.Replace(gdType, ".", "", -1), "enum::")
		gdType = strings.TrimPrefix(strings.Replace(gdType, ".", "", -1), "bitfield::")

		gdType = strings.Replace(gdType, ".", "", -1)

		if class, ok := classDB[gdType]; ok {
			if pkg != class.Package {
				if class.Package == "internal" {
					class.Package = "gd"
				}
				return class.Package + "." + class.Name
			}
			return class.Name
		}

		if inCore(gdType) {
			return maybeInternal(gdType)
		}

		return gdType
	}
}

func mapOperator(name string) string {
	switch name {
	case "==":
		return "Equals"
	case "!=":
		return "NotEqual"
	case "<":
		return "Less"
	case "<=":
		return "LessEqual"
	case ">":
		return "Greater"
	case ">=":
		return "GreaterEqual"
	case "+":
		return "Add"
	case "-":
		return "Subtract"
	case "*":
		return "Multiply"
	case "/":
		return "Divide"
	case "unary-":
		return "Negate"
	case "%":
		return "Module"
	case "**":
		return "Power"
	case "<<":
		return "ShiftLeft"
	case ">>":
		return "ShiftRight"
	case "&":
		return "BitAnd"
	case "|":
		return "BitOr"
	case "^":
		return "BitXor"
	case "~":
		return "BitNegate"
	case "and":
		return "And"
	case "or":
		return "Or"
	case "xor":
		return "Xor"
	case "not":
		return "Not"
	case "in":
		return "In"
	default:
		panic("unknown operator: " + name)
	}
}

func convertName(fnName string) string {
	if fnName == "seek" {
		return "SeekTo"
	}
	if fnName == "type_string" {
		return "TypeToString"
	}

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

func genEnum(pkg string, decl, code io.Writer, prefix string, enum Enum) {
	name := prefix + strings.Replace(enum.Name, ".", "", -1)
	if name == "Side" || name == "EulerOrder" {
		return
	}

	if decl != nil {
		fmt.Fprintln(decl)
		fmt.Fprintf(decl, "type %v int64\n", name)

		fmt.Fprintln(code)

		var topLevelPrefix = pkg
		if topLevelPrefix == "internal" {
			topLevelPrefix = "gd"
		}
		fmt.Fprintf(code, "type %v = %v.%[1]v\n", name, topLevelPrefix)

		if name == "GDExtensionInitializationLevel" || name == "VariantType" || name == "Error" {
			genEnum(pkg, nil, decl, prefix, enum)
		}
	}
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

type ClassDB map[string]Class

func (db ClassDB) nameOf(pkg, original string) string {
	class := db[original]
	if pkg == class.Package {
		return class.Name
	}
	if class.Package == "internal" {
		return "gd." + class.Name
	}
	return class.Package + "." + class.Name
}

func (db ClassDB) genArgs(pkg string, method Method) string {
	var args []string
	for _, arg := range method.Arguments {
		if arg.Name == "type" {
			arg.Name = "typ"
		}
		if arg.Name == "func" {
			arg.Name = "fn"
		}
		if arg.Name == "default" {
			arg.Name = "def"
		}
		if arg.Name == "interface" {
			arg.Name = "iface"
		}
		if arg.Name == "range" {
			arg.Name = "rng"
		}
		if arg.Name == "var" {
			arg.Name = "variable"
		}
		if arg.Name == "map" {
			arg.Name = "mapping"
		}
		if arg.Name == "internal" {
			arg.Name = "internal_"
		}
		args = append(args, fmt.Sprintf("%v %v", arg.Name, db.convertType(pkg, arg.Meta, arg.Type)))
	}
	return strings.Join(args, ", ")
}

func pascal(name string) string {
	words := strings.Split(name, "_")
	for i := range words {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, "")
}

func fixReserved(name string) string {
	switch name {
	case "bool":
		return "b"
	case "type":
		return "atype"
	case "range":
		return "arange"
	case "default":
		return "def"
	case "func":
		return "fn"
	case "frame":
		return "frame_"
	case "interface":
		return "intf"
	case "internal":
		return "internal_"
	case "map":
		return "mapping"
	case "var":
		return "v"
	case "string":
		return "s"
	case "RID":
		return "rid"
	default:
		return name
	}
}

func inCore(s string) bool {
	switch s {
	case "Object", "RefCounted":
		return true
	default:
		return false
	}
}

func generate() error {
	var spec Specification
	if err := spec.Load(); err != nil {
		return err
	}

	var classDB = make(ClassDB)
	for _, enum := range spec.GlobalEnums {
		classDB[strings.Replace(enum.Name, ".", "", -1)] = Class{
			IsEnum:  true,
			Name:    strings.Replace(enum.Name, ".", "", -1),
			Package: "internal",
		}
	}
	for _, class := range spec.BuiltinClasses {
		for _, enum := range class.Enums {
			classDB[class.Name+strings.Replace(enum.Name, ".", "", -1)] = Class{
				IsEnum:  true,
				Name:    class.Name + strings.Replace(enum.Name, ".", "", -1),
				Package: "internal",
			}
		}
	}
	for i, class := range spec.Classes {
		var pkg = "internal"
		if !inCore(class.Name) {
			pkg = "classdb"
		}
		class.Package = pkg
		for _, enum := range class.Enums {
			classDB[class.Name+strings.Replace(enum.Name, ".", "", -1)] = Class{
				IsEnum:  true,
				Name:    class.Name + strings.Replace(enum.Name, ".", "", -1),
				Package: pkg,
			}
		}
		spec.Classes[i] = class
		classDB[class.Name] = class
	}

	file, err := os.Open("../extension_api.json")
	if os.IsNotExist(err) {
		req, err := http.NewRequest("GET", "https://raw.githubusercontent.com/godotengine/godot-headers/master/extension_api.json", nil)
		if err != nil {
			return err
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}

		file, err = os.Create("../extension_api.json")
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

	out, err := os.Create("./out.go")
	if err != nil {
		return err
	}
	fmt.Fprintln(out, `//go:build !generate`)
	fmt.Fprintln(out)
	fmt.Fprintln(out, `// Code generated by the generate package DO NOT EDIT`)
	fmt.Fprintln(out, `package gd`)
	fmt.Fprintln(out)
	fmt.Fprintln(out, `import "runtime.link/api/call"`)
	fmt.Fprintln(out)

	core, err := os.Create("./all.go")
	if err != nil {
		return err
	}
	fmt.Fprintln(core, `//go:build !generate`)
	fmt.Fprintln(core)
	fmt.Fprintln(core, `// Code generated by the generate package DO NOT EDIT`)
	fmt.Fprintln(core, `package gd`)
	fmt.Fprintln(core)
	fmt.Fprintln(core, `import "unsafe"`)
	fmt.Fprintln(core, `import "reflect"`)
	fmt.Fprintln(core, `import "runtime.link/mmm"`)
	fmt.Fprintln(core, `import "runtime.link/api/call"`)
	fmt.Fprintln(core)

	all, err := os.Create("./classdb/all.go")
	if err != nil {
		return err
	}
	fmt.Fprintln(all, `//go:build !generate`)
	fmt.Fprintln(all)
	fmt.Fprintln(all, `// Code generated by the generate package DO NOT EDIT`)
	fmt.Fprintln(all, `package gd`)
	fmt.Fprintln(all)
	fmt.Fprintln(all, `import "unsafe"`)
	fmt.Fprintln(all, `import "reflect"`)
	fmt.Fprintln(all, `import "runtime.link/mmm"`)
	fmt.Fprintln(all, `import "runtime.link/api/call"`)
	fmt.Fprintln(all, `import gd "grow.graphics/gd/internal"`)
	fmt.Fprintln(all)

	enums, err := os.Create("../enums.go")
	if err != nil {
		return err
	}
	fmt.Fprintln(enums, `//go:build !generate`)
	fmt.Fprintln(enums)
	fmt.Fprintln(enums, `// Code generated by the generate package DO NOT EDIT`)
	fmt.Fprintln(enums, `package gd`)
	fmt.Fprintln(enums)
	fmt.Fprintln(enums, `import gd "grow.graphics/gd/internal"`)
	fmt.Fprintln(enums, `import classdb "grow.graphics/gd/internal/classdb"`)
	fmt.Fprintln(enums)

	classdb, err := os.Create("../classdb.go")
	if err != nil {
		return err
	}
	fmt.Fprintln(classdb, `//go:build !generate`)
	fmt.Fprintln(classdb)
	fmt.Fprintln(classdb, `// Code generated by the generate package DO NOT EDIT`)
	fmt.Fprintln(classdb, `package gd`)
	fmt.Fprintln(classdb)
	fmt.Fprintln(classdb, `import "unsafe"`)
	fmt.Fprintln(classdb, `import gd "grow.graphics/gd/internal"`)
	fmt.Fprintln(classdb, `import classdb "grow.graphics/gd/internal/classdb"`)
	fmt.Fprintln(classdb)

	for _, enum := range spec.GlobalEnums {
		genEnum("internal", core, enums, "", enum)
	}

	var singletons = make(map[string]bool)
	for _, class := range spec.Singletons {
		singletons[class.Name] = true
	}

	var builtin_sizes = make(map[string]int)
	for _, config := range spec.BuiltinClassSizes {
		if config.BuildConfiguration != "float_64" {
			continue
		}
		for _, size := range config.Sizes {
			builtin_sizes[size.Name] = size.Size
		}
	}

	fmt.Fprintf(out, "type utility struct{\n")
	for _, utility := range spec.UtilityFunctions {
		fmt.Fprintf(out, "\t%v func(ret uintptr, args call.Args, c int32) `hash:\"%v\"`\n", utility.Name, utility.Hash)

		fmt.Fprintf(core, "\nfunc (ctx Context) %v(", convertName(utility.Name))
		for i, arg := range utility.Arguments {
			if i > 0 {
				fmt.Fprintf(core, ", ")
			}
			fmt.Fprintf(core, "%v %v", fixReserved(arg.Name), classDB.convertType("internal", arg.Meta, arg.Type))
		}
		fmt.Fprintf(core, ")")
		result := classDB.convertType("internal", "", utility.ReturnType)
		ptrKind, isPtr := classDB.isPointer(result)
		if result != "" {
			fmt.Fprintf(core, " %v {\n", result)
		} else {
			fmt.Fprintf(core, " {\n")
		}
		fmt.Fprintf(core, "\tvar frame = call.New()\n")
		for _, arg := range utility.Arguments {
			_, ok := classDB[arg.Type]
			if ok {
				fmt.Fprintf(core, "\tcall.Arg(frame, mmm.Get(%v))\n", fixReserved(arg.Name))
				continue
			}
			argType := classDB.convertType("internal", arg.Meta, arg.Type)
			_, argIsPtr := classDB.isPointer(argType)
			if argIsPtr {
				fmt.Fprintf(core, "\tcall.Arg(frame, mmm.Get(%v))\n", fixReserved(arg.Name))
			} else {
				fmt.Fprintf(core, "\tcall.Arg(frame, %v)\n", fixReserved(arg.Name))
			}
		}
		if isPtr {
			fmt.Fprintf(core, "\tvar r_ret = call.Ret[%v](frame)\n", ptrKind)
		} else {
			if result != "" {
				fmt.Fprintf(core, "\tvar r_ret = call.Ret[%v](frame)\n", result)
			} else {
				fmt.Fprintf(core, "\tvar r_ret call.Nil\n")
			}
		}
		fmt.Fprintf(core, "\tctx.API.utility.%v(r_ret.Uintptr(), frame.Array(0), %d)\n", utility.Name, len(utility.Arguments))
		if isPtr {
			_, ok := classDB[result]
			if ok || result == "Object" {
				fmt.Fprintf(core, "\tvar ret %v\n", result)
				fmt.Fprintf(core, "\tret.SetPointer(mmm.New[Pointer](ctx.Lifetime, ctx.API, r_ret.Get()))\n")
			} else {
				fmt.Fprintf(core, "\tvar ret = mmm.New[%v](ctx.Lifetime, ctx.API, r_ret.Get())\n", result)
			}
		} else {
			if result != "" {
				fmt.Fprintf(core, "\tvar ret = r_ret.Get()\n")
			}
		}
		fmt.Fprintf(core, "\tframe.Free()\n")
		if result != "" {
			fmt.Fprintf(core, "\treturn ret\n")
		}
		fmt.Fprintf(core, "}\n")
	}
	fmt.Fprintf(out, "}\n")

	fmt.Fprintf(out, "type builtin struct{\n")
	for _, class := range spec.BuiltinClasses {
		for _, enum := range class.Enums {
			genEnum("internal", core, enums, class.Name, enum)
		}
		fmt.Fprintf(out, "\t%v struct{\n", class.Name)
		for _, method := range class.Methods {
			if method.Name == "map" {
				method.Name = "map_"
			}
			fmt.Fprintf(out, "\t\t%v func(base uintptr, args call.Args, ret uintptr, c int32) `hash:\"%v\"`\n", method.Name, method.Hash)
		}
		fmt.Fprintf(out, "\t}\n")
	}
	fmt.Fprintf(out, "}\n")

	fmt.Fprintf(out, "type typeset struct{\n")
	fmt.Fprintf(out, "\tcreation struct{\n")
	for _, class := range spec.BuiltinClasses {
		fmt.Fprintf(out, "\t\t%v [%d]func(uintptr, call.Args)\n", class.Name, len(class.Constructors))
	}
	fmt.Fprintf(out, "\t}\n")
	fmt.Fprintf(out, "\toperator struct{\n")
	for _, class := range spec.BuiltinClasses {
		fmt.Fprintf(out, "\t\t%v struct {\n", class.Name)
		sort.Slice(class.Operators, func(i, j int) bool {
			return class.Operators[i].Name < class.Operators[j].Name
		})
		var last string
		for _, op := range class.Operators {
			if op.Name == "unary+" || op.RightType == "Variant" {
				continue
			}
			if strings.HasPrefix(op.Name, "unary") || op.RightType == "" {
				if last != "" {
					fmt.Fprintf(out, "\t\t\t}\n")
				}
				fmt.Fprintf(out, "\t\t\t%v func(a, b, ret uintptr)\n", mapOperator(op.Name))
				last = ""
			} else {
				if op.Name != last {
					if last != "" {
						fmt.Fprintf(out, "\t\t\t}\n")
					}
					fmt.Fprintf(out, "\t\t\t%v struct {\n", mapOperator(op.Name))
				}
				fmt.Fprintf(out, "\t\t\t\t%v func(a, b, ret uintptr)\n", op.RightType)
				last = op.Name
			}
		}
		if last != "" {
			fmt.Fprintf(out, "\t\t\t}\n")
		}
		fmt.Fprintf(out, "\t\t}\n")
	}
	fmt.Fprintf(out, "\t}\n")
	fmt.Fprintf(out, "\tdestruct struct{\n")
	for _, class := range spec.BuiltinClasses {
		fmt.Fprintf(out, "\t\t%v func(uintptr)\n", class.Name)
		for _, method := range class.Methods {
			classDB.methodCall(core, "internal", class, method, callBuiltin)
		}
	}
	fmt.Fprintf(out, "\t}\n")
	fmt.Fprintf(out, "}\n")

	fmt.Fprintf(out, "type singletons struct{\n")
	for _, class := range spec.Classes {
		if singletons[class.Name] {
			fmt.Fprintf(out, "\t%v StringName\n", class.Name)
		}
	}
	fmt.Fprintf(out, "}\n")

	fmt.Fprintf(out, "type methods struct{\n")
	for _, class := range spec.Classes {
		if class.Name == "JavaClassWrapper" || class.Name == "JavaScriptBridge" {
			continue
		}
		fmt.Fprintf(out, "\t%v struct{\n", class.Name)
		for _, method := range class.Methods {
			if method.Name == "select" {
				method.Name = "select_"
			}
			if method.IsVirtual {
				continue
			}
			fmt.Fprintf(out, "\t\tBind_%v MethodBind `hash:\"%v\"`\n", method.Name, method.Hash)
		}
		fmt.Fprintf(out, "\t}\n")
	}
	fmt.Fprintf(out, "}\n")

	for _, class := range spec.Classes {
		var w = core
		if class.Package != "internal" {
			w = all
		}
		pkg := class.Package

		prefix := ""
		if pkg != "internal" {
			prefix = "gd."
		}

		for _, enum := range class.Enums {
			genEnum(pkg, w, enums, class.Name, enum)
		}
		if class.Name != "Object" {
			if class.Description != "" {
				fmt.Fprintln(w, "/*")
				fmt.Fprint(w, strings.Replace(class.Description, "*/", "", -1))
				fmt.Fprintln(w, "\n*/")
			}
			if singletons[class.Name] {
				fmt.Fprintf(classdb, "func %v(godot Context) %[2]v { obj := godot.API.Object.GetSingleton(godot, godot.API.Singletons.%[1]v); return *(*%[2]v)(unsafe.Pointer(&obj)) }\n",
					class.Name, classDB.nameOf("gd", class.Name))
			} else {
				fmt.Fprintf(classdb, "type %v = %v \n", class.Name, classDB.nameOf("gd", class.Name))
			}
		}
		fmt.Fprintf(w, "type %[1]v struct {_ [0]*%[1]v; ptr "+prefix+"Pointer}\n", class.Name, classDB.nameOf(pkg, class.Inherits))
		fmt.Fprintf(w, "\n//go:nosplit\nfunc (self %[1]v) AsPointer() "+prefix+"Pointer { return self.ptr }\n", class.Name)
		fmt.Fprintf(w, "\n\n//go:nosplit\nfunc (self *%[1]v) SetPointer(ptr "+prefix+"Pointer) { self.ptr = ptr }\n", class.Name)
		if class.Inherits != "" {
			var i = 1
			super := classDB[class.Inherits]
			fmt.Fprintf(w, "\n\n//go:nosplit\nfunc (self %[1]v) Super() %[3]v { return *(*%[3]v)(unsafe.Pointer(&self)) }\n", class.Name, super.Name, classDB.nameOf(pkg, super.Name))
			for super.Name != "" {
				fmt.Fprintf(w, "\n\n//go:nosplit\nfunc (self %[1]v) As%[2]v() %[3]v { return *(*%[3]v)(unsafe.Pointer(&self)) }\n", class.Name, super.Name, classDB.nameOf(pkg, super.Name))
				i++
				super = classDB[super.Inherits]
			}

		}
		if singletons[class.Name] {
			fmt.Fprintf(w, "\n//go:nosplit\n\nfunc (self %[1]v) IsSingleton() {}\n", class.Name)
		}
		for _, method := range class.Methods {
			classDB.methodCall(w, pkg, class, method, callDefault)
		}
		fmt.Fprintf(w, "\nfunc (self %[1]v) Virtual(name string) reflect.Value {\n", class.Name)
		fmt.Fprintf(w, "\tswitch name {\n")
		for _, method := range class.Methods {
			if method.IsVirtual {
				fmt.Fprintf(w, "\tcase \"%v\": return reflect.ValueOf(self.%v);\n", method.Name, method.Name)
			}
		}
		if class.Inherits != "" {
			fmt.Fprintf(w, "\tdefault: return "+prefix+"VirtualByName(self.Super(), name)\n")
		} else {
			fmt.Fprintf(w, "\tdefault: return reflect.Value{}\n")
		}
		fmt.Fprintf(w, "\t}\n")
		fmt.Fprintf(w, "}\n")
	}

	return nil
}

func (db ClassDB) isPointer(t string) (string, bool) {
	t = strings.TrimPrefix(t, "gd.")
	switch t {
	case "String", "StringName", "NodePath",
		"Dictionary", "Array":
		return "uintptr", true
	case "Callable", "Signal",
		"PackedByteArray", "PackedInt32Array",
		"PackedInt64Array", "PackedFloat32Array",
		"PackedFloat64Array", "PackedStringArray",
		"PackedVector2Array", "PackedVector3Array",
		"PackedColorArray":
		return "[2]uintptr", true
	case "Variant":
		return "[3]uintptr", true
	default:
		if entry, ok := db[t]; ok && !entry.IsEnum {
			return "uintptr", true
		}
		return "", false
	}
}

type callType int

const (
	callDefault callType = iota
	callBuiltin
	callUtility
)

func (classDB ClassDB) methodCall(w io.Writer, pkg string, class Class, method Method, ctype callType) {
	switch class.Name {
	case "Float", "Int", "Vector2", "Vector2i", "Rect2", "Rect2i", "Vector3", "Vector3i",
		"Transform2D", "Vector4", "Vector4i", "Plane", "Quaternion", "AABB", "Basis", "Transform3D",
		"RID", "Projection", "Color":
		return
	}
	result := classDB.convertType(pkg, method.ReturnValue.Meta, method.ReturnValue.Type)
	if method.ReturnType != "" {
		result = classDB.convertType(pkg, "", method.ReturnType)
	}
	ptrKind, isPtr := classDB.isPointer(result)

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
		fmt.Fprintf(w, "func (%s) %s(impl func(ptr unsafe.Pointer, ctx "+prefix+"Context", class.Name, method.Name)
		for _, arg := range method.Arguments {
			fmt.Fprint(w, ", ")
			fmt.Fprintf(w, "%v %v", fixReserved(arg.Name), classDB.convertType(pkg, arg.Meta, arg.Type))
		}
		fmt.Fprintf(w, ") %v, api *"+prefix+"API) (cb "+prefix+"ExtensionClassCallVirtualFunc) {\n", result)
		fmt.Fprintf(w, "\treturn func(class any, p_args "+prefix+"UnsafeArgs, p_back "+prefix+"UnsafeBack) {\n")
		fmt.Fprintf(w, "\t\tctx := %vNewContext(api)\n", prefix)
		for i, arg := range method.Arguments {
			var argType = classDB.convertType(pkg, arg.Meta, arg.Type)

			_, ok := classDB[arg.Type]
			if ok {
				fmt.Fprintf(w, "\t\tvar %v %v\n", fixReserved(arg.Name), argType)
				fmt.Fprintf(w, "\t\t%v.SetPointer(mmm.New["+prefix+"Pointer](ctx.Lifetime, ctx.API, "+prefix+"UnsafeGet[uintptr](p_args,%d)))\n", fixReserved(arg.Name), i)
				continue
			}

			argPtrKind, argIsPtr := classDB.isPointer(argType)
			if argIsPtr {
				fmt.Fprintf(w, "\t\tvar %v = mmm.New[%v](ctx.Lifetime, ctx.API, "+prefix+"UnsafeGet[%v](p_args,%d))\n", fixReserved(arg.Name), argType, argPtrKind, i)
			} else {
				fmt.Fprintf(w, "\t\tvar %v = "+prefix+"UnsafeGet[%v](p_args,%d)\n", fixReserved(arg.Name), argType, i)
			}
		}
		fmt.Fprintf(w, "\t\tself := reflect.ValueOf(class).UnsafePointer()\n")
		if result != "" {
			fmt.Fprintf(w, "\t\tret := ")
		}
		fmt.Fprintf(w, "impl(self,ctx")
		for _, arg := range method.Arguments {
			fmt.Fprint(w, ", ")
			fmt.Fprintf(w, "%v", fixReserved(arg.Name))
		}
		fmt.Fprintf(w, ")\n")
		if isPtr {
			if _, ok := classDB[result]; ok || result == "gd.Object" {
				fmt.Fprintf(w, "\t\tmmm.End(ret.AsPointer())\n")
			} else {
				fmt.Fprintf(w, "\t\tmmm.End(ret)\n")
			}
		}
		fmt.Fprintf(w, "\t\tctx.End()\n")
		if result != "" {
			fmt.Fprintf(w, "\t\t"+prefix+"UnsafeSet[%v](p_back, ret)\n", result)
		}
		fmt.Fprintf(w, "\t}\n")
		fmt.Fprintf(w, "}\n")
		return
	}
	if class.Name == "JavaClassWrapper" || class.Name == "JavaScriptBridge" {
		return
	}
	var context = "ctx " + prefix + "Context"
	if !isPtr {
		context = ""
	}
	fmt.Fprintf(w, "//go:nosplit\nfunc (self %v) %v(%s", class.Name, convertName(method.Name), context)

	if method.Name == "select" {
		method.Name = "select_"
	}
	if method.Name == "map" {
		method.Name = "map_"
	}

	for i, arg := range method.Arguments {
		if i > 0 || isPtr {
			fmt.Fprint(w, ", ")
		}
		fmt.Fprintf(w, "%v %v", fixReserved(arg.Name), classDB.convertType(pkg, arg.Meta, arg.Type))
	}
	fmt.Fprintf(w, ") %v {\n", result)
	if ctype == callBuiltin {
		fmt.Fprintf(w, "\tvar selfPtr = self\n")
	} else {
		fmt.Fprintf(w, "\tvar selfPtr = self.AsPointer()\n")
	}
	fmt.Fprintf(w, "\tvar frame = call.New()\n")
	for _, arg := range method.Arguments {
		_, ok := classDB[arg.Type]
		if ok {
			fmt.Fprintf(w, "\tcall.Arg(frame, mmm.Get(%v.AsPointer()))\n", fixReserved(arg.Name))
			continue
		}

		argType := classDB.convertType(pkg, arg.Meta, arg.Type)
		_, argIsPtr := classDB.isPointer(argType)
		if argIsPtr {
			fmt.Fprintf(w, "\tcall.Arg(frame, mmm.Get(%v))\n", fixReserved(arg.Name))
		} else {
			fmt.Fprintf(w, "\tcall.Arg(frame, %v)\n", fixReserved(arg.Name))
		}
	}
	if isPtr {
		fmt.Fprintf(w, "\tvar r_ret = call.Ret[%v](frame)\n", ptrKind)
	} else {
		if result != "" {
			fmt.Fprintf(w, "\tvar r_ret = call.Ret[%v](frame)\n", result)
		} else {
			fmt.Fprintf(w, "\tvar r_ret call.Nil\n")
		}
	}
	if ctype == callBuiltin {
		fmt.Fprintf(w, "\tvar p_self = call.Arg(frame, mmm.Get(selfPtr))\n")
		fmt.Fprintf(w, "\tmmm.API(selfPtr).builtin.%v.%v(p_self.Uintptr(), frame.Array(0), r_ret.Uintptr(), %d)\n", class.Name, method.Name, len(method.Arguments))
	} else {
		fmt.Fprintf(w, "\tmmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.%v.Bind_%v, self.AsObject(), frame.Array(0), r_ret.Uintptr())\n", class.Name, method.Name)
	}

	if isPtr {
		_, ok := classDB[result]
		if ok || result == "gd.Object" {
			fmt.Fprintf(w, "\tvar ret %v\n", result)
			fmt.Fprintf(w, "\tret.SetPointer(mmm.New["+prefix+"Pointer](ctx.Lifetime, ctx.API, r_ret.Get()))\n")
		} else {
			fmt.Fprintf(w, "\tvar ret = mmm.New[%v](ctx.Lifetime, ctx.API, r_ret.Get())\n", result)
		}
	} else if result != "" {
		fmt.Fprintf(w, "\tvar ret = r_ret.Get()\n")
	}
	fmt.Fprintf(w, "\tframe.Free()\n")
	if result != "" {
		fmt.Fprintf(w, "\treturn ret\n")
	}
	/*if result != "" {
		fmt.Fprintf(w, "\tvar ret %v\n", result)
		fmt.Fprintf(w, "\treturn ret\n")
	}*/
	fmt.Fprintf(w, "}")
}

func main() {
	if err := generate(); err != nil {
		log.Fatal(err)
	}
}
