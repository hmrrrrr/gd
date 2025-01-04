package shaders

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Shader"
	"graphics.gd/classdb/ShaderMaterial"
	gd "graphics.gd/internal"
	"graphics.gd/shaders/internal/dsl"
	"graphics.gd/shaders/mat2"
	"graphics.gd/shaders/mat3"
	"graphics.gd/shaders/mat4"
	"graphics.gd/shaders/vec1"
	"graphics.gd/shaders/vec2"
	"graphics.gd/shaders/vec3"
	"graphics.gd/shaders/vec4"
	"runtime.link/xyz"
)

type goShader struct {
	classdb.Extension[goShader, ShaderMaterial.Instance]
}

func init() {
	gd.StartupFunctions = append(gd.StartupFunctions, func() {
		classdb.Register[goShader]()
	})
}

type (
	Type2D struct {
		classdb.Extension[goShader, ShaderMaterial.Instance]
	}
	Type3D struct {
		classdb.Extension[goShader, ShaderMaterial.Instance]
	}
)

func (Type2D) shaderType() string { return "canvas_item" }
func (Type3D) shaderType() string { return "spatial" }

type RenderingOption2D xyz.Switch[string, struct {
	BlendingModeMix            RenderingOption2D `json:"blend_mix"`             // Mix blend mode (alpha is transparency), default.
	BlendingModeAdd            RenderingOption2D `json:"blend_add"`             // Additive blend mode.
	BlendingModeSub            RenderingOption2D `json:"blend_sub"`             // Subtractive blend mode.
	BlendingModeMul            RenderingOption2D `json:"blend_mul"`             // Multiplicative blend mode.
	BlendingPremultipliedAlpha RenderingOption2D `json:"blend_premul_alpha"`    // Pre-multiplied alpha blend mode.
	BlendingDisabled           RenderingOption2D `json:"blend_disabled"`        // Disable blending, values (including alpha) are written as-is.
	Unshaded                   RenderingOption2D `json:"unshaded"`              // Result is just albedo. No lighting/shading happens in material.
	LightOnly                  RenderingOption2D `json:"light_only"`            // Only draw on light pass.
	SkipVertexTransform        RenderingOption2D `json:"skip_vertex_transform"` // VERTEX needs to be transformed manually in vertex function.
	WorldVertexCoordinates     RenderingOption2D `json:"world_vertex_coords"`   // VERTEX is modified in world coordinates instead of local.
}]

var RenderingOptions2D = xyz.AccessorFor(RenderingOption2D.Values)

type Program[V, F, M any] interface {
	Super() ShaderMaterial.Instance

	shaderType() string

	Fragment(V) F
	Material(F) M
	Lighting(M) vec4.RGBA
}

type Global struct {
	// Global time since the engine has started, in seconds. It repeats after every 3,600 seconds (which can be changed with
	// the rollover setting). It's not affected by time_scale or pausing. If you need a TIME variable that can be scaled or
	// paused, add your own global shader uniform and update it each frame.
	Time vec1.X `gd:"TIME"`
}

func Compile[V, F, M any](prog Program[V, F, M]) {
	gd.NewCallable(func() {
		material := prog.Super()
		shader := Shader.New()
		writer := strings.Builder{}
		fmt.Fprintf(&writer, "// Code generated by graphics.gd/shaders DO NOT EDIT!\n")
		fmt.Fprintf(&writer, "shader_type %s;\n\n", prog.shaderType())

		linkup(prog)
		var vertices V
		linkup(&vertices)
		var fragment F
		linkup(&fragment)

		compileUniforms(&writer, prog)
		compileFragmentShader(&writer, prog.Fragment(vertices))
		compileMaterialShader(&writer, prog.Material(fragment))
		//compileLightingShader(&writer, prog.Lighting([1]M{}[0]))
		fmt.Println(writer.String())
		shader.SetCode(writer.String())
		material.SetShader(shader)
	}).CallDeferred()
}

func linkup(in any) {
	value := reflect.ValueOf(in).Elem()
	rtype := value.Type()
	for i := 0; i < rtype.NumField(); i++ {
		if value.Field(i).Kind() == reflect.Struct && rtype.Field(i).IsExported() {
			linkup(value.Field(i).Addr().Interface())
		}
		if tag := rtype.Field(i).Tag.Get("gd"); tag != "" {
			dsl.Set(value.Field(i).Addr().Interface().(dsl.Pointer), dsl.Identifier(tag))
		}
	}
}

func compileUniforms(w io.Writer, uniforms any) {
	value := reflect.ValueOf(uniforms).Elem()
	rtype := value.Type()
	for i := 0; i < rtype.NumField(); i++ {
		if tag := rtype.Field(i).Tag.Get("gd"); tag != "" {
			fmt.Fprintf(w, "uniform ")
			switch field := rtype.Field(i); reflect.Zero(field.Type).Interface().(type) {
			case vec2.XY:
				fmt.Fprintf(w, "vec2 %s;\n", tag)
			case vec4.RGBA:
				fmt.Fprintf(w, "vec4 %s;\n", tag)
			case vec1.X:
				fmt.Fprintf(w, "float %s;\n", tag)
			default:
				panic(fmt.Sprintf("unsupported uniform type %s", field.Type))
			}
		}
	}
	fmt.Fprintln(w)
}

func compileFragmentShader(w io.Writer, vertices any) {
	fmt.Fprintf(w, "void vertex() {\n")
	value := reflect.ValueOf(vertices)
	rtype := value.Type()
	for i := 0; i < rtype.NumField(); i++ {
		field := rtype.Field(i)
		switch field.Tag.Get("gd") {
		case "VERTEX":
			fmt.Fprintf(w, "\tVERTEX = ")
			compileExpression(w, value.Field(i).Interface().(dsl.Evaluator))
			fmt.Fprintf(w, ";\n")
		}
	}
	fmt.Fprintf(w, "}\n")
}

func compileMaterialShader(w io.Writer, material any) {
	fmt.Fprintf(w, "void fragment() {\n")
	value := reflect.ValueOf(material)
	rtype := value.Type()
	for i := 0; i < rtype.NumField(); i++ {
		field := rtype.Field(i)
		switch field.Tag.Get("gd") {
		case "COLOR":
			fmt.Fprintf(w, "\tCOLOR = ")
			compileExpression(w, value.Field(i).Interface().(dsl.Evaluator))
			fmt.Fprintf(w, ";\n")
		}
	}
	fmt.Fprintf(w, "}\n")
}

func compileExpression(w io.Writer, expression dsl.Evaluator) {
	if expr := dsl.Evaluate(expression); expr != nil {
		expression = expr
	}
	fmt.Println(reflect.TypeOf(expression))
	switch value := expression.(type) {
	case vec4.RGBA:
		compileCall(w, "vec4", value.R, value.G, value.B, value.A)
	case vec4.XYZW:
		compileCall(w, "vec4", value.X, value.Y, value.Z, value.W)
	case vec4.T[int]:
		compileCall(w, "ivec4", value.X, value.Y, value.Z, value.W)
	case vec4.T[uint]:
		compileCall(w, "uvec4", value.X, value.Y, value.Z, value.W)
	case vec4.T[bool]:
		compileCall(w, "bvec4", value.X, value.Y, value.Z, value.W)
	case vec3.XYZ:
		compileCall(w, "vec3", value.X, value.Y, value.Z)
	case vec3.RGB:
		compileCall(w, "vec3", value.R, value.G, value.B)
	case vec3.T[int]:
		compileCall(w, "ivec3", value.X, value.Y, value.Z)
	case vec3.T[uint]:
		compileCall(w, "uvec3", value.X, value.Y, value.Z)
	case vec3.T[bool]:
		compileCall(w, "bvec3", value.X, value.Y, value.Z)
	case vec2.XY:
		compileCall(w, "vec2", value.X, value.Y)
	case vec2.T[int]:
		compileCall(w, "ivec2", value.X, value.Y)
	case vec2.T[uint]:
		compileCall(w, "uvec2", value.X, value.Y)
	case vec2.T[bool]:
		compileCall(w, "bvec2", value.X, value.Y)
	case vec1.X:
		fmt.Fprintf(w, "%f", value.X)
	case vec1.T[int]:
		fmt.Fprintf(w, "%d", value.X)
	case vec1.T[uint]:
		fmt.Fprintf(w, "%d", value.X)
	case vec1.T[bool]:
		fmt.Fprintf(w, "%t", value.X)
	case mat2.ColumnMajor:
		compileCall(w, "mat2", value.Columns[0][0], value.Columns[0][1], value.Columns[1][0], value.Columns[1][1])
	case mat3.ColumnMajor:
		compileCall(w, "mat3", value.Columns[0][0], value.Columns[0][1], value.Columns[0][2], value.Columns[1][0], value.Columns[1][1], value.Columns[1][2], value.Columns[2][0], value.Columns[2][1], value.Columns[2][2])
	case mat4.ColumnMajor:
		compileCall(w, "mat4", value.Columns[0][0], value.Columns[0][1], value.Columns[0][2], value.Columns[0][3], value.Columns[1][0], value.Columns[1][1], value.Columns[1][2], value.Columns[1][3], value.Columns[2][0], value.Columns[2][1], value.Columns[2][2], value.Columns[2][3], value.Columns[3][0], value.Columns[3][1], value.Columns[3][2], value.Columns[3][3])
	case dsl.Operation:
		fmt.Fprintf(w, "(")
		compileExpression(w, value.A)
		fmt.Fprintf(w, " %s ", value.Op)
		compileExpression(w, value.B)
		fmt.Fprintf(w, ")")
	case dsl.Identifier:
		fmt.Fprintf(w, "%s", value)
	case dsl.Ternary:
		fmt.Fprintf(w, "(")
		compileExpression(w, value.If)
		fmt.Fprintf(w, " ? ")
		compileExpression(w, value.A)
		fmt.Fprintf(w, " : ")
		compileExpression(w, value.B)
		fmt.Fprintf(w, ")")
	case dsl.FunctionCall:
		compileCall(w, string(value.Name), value.Args...)
	default:
		panic(fmt.Sprintf("unsupported expression type %T", expression))
	}
}

func compileCall(w io.Writer, name string, args ...dsl.Evaluator) {
	fmt.Fprintf(w, "%s(", name)
	for i, arg := range args {
		if i > 0 {
			fmt.Fprintf(w, ", ")
		}
		compileExpression(w, arg)
	}
	fmt.Fprintf(w, ")")
}
