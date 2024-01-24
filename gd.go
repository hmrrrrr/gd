package gd

import (
	gd "grow.graphics/gd/internal"
	internal "grow.graphics/gd/internal"
)

type Axis int

const (
	X Axis = iota
	Y
	Z
	W
)

// Class can be embedded inside of a struct to represent a new Class type.
// The extended class will be available by calling the [Class.Super] method.
type Class[T, S internal.IsClass] struct {
	internal.Class[T, S]
}

// Context for ownership and a reference to the Godot API, apart from
// its use as an ordinary [context.Context] to signal cancellation, this
// value is not safe to use concurrently. Each goroutine should create
// its own [Context] and use that instead.
//
//	newctx := gd.NewContext(oldctx.API())
//
// When a [Context] is freed, it will free all of the objects that were
// created using it. A [Context] should not be used after free, as it
// will be recycled and will cause values to be unexpectedly freed.
//
// When the context has been passed in as a function argument, always
// assume that the [Context] will be freed when the function returns.
// Classes can be moved between contexts using their [KeepAlive] method.
type Context = internal.Context

// Create a new instance of the given class, which should be an uninitialised
// pointer to a value of that class. T must be a class from this package.
func Create[T internal.PointerToClass](ctx Context, ptr T) T {
	return internal.Create[T](ctx, ptr)
}

// Const can be used to retrieve a 'constant' value from a structured type.
func Const[F func(T) T, T any](constant F) T {
	var zero T
	return constant(zero)
}

// NewVector2 constructs a new Vector2 from the given x and y.
func NewVector2(x, y Float) Vector2 { return Vector2{float32(x), float32(y)} }

/*
Register registers a struct available for use inside Godot
extending the given 'Parent' Godot class. The 'Struct' type must
be a named struct that embeds a [Class] field specifying the
parent class to extend.

	type MyClass struct {
		gd.Class[MyClass, Node2D]
	}

Use this in a main or init function to register your Go structs
and they will become available within the Godot engine for use
in the editor and/or within scripts.

All exported fields and methods will be exposed to Godot, so
take caution when embedding types, as their fields and methods
will be promoted.

If the Struct extends [EditorPlugin] then it will be added to
the editor as a plugin.
*/
func Register[Struct gd.Extends[Parent], Parent gd.IsClass](godot Context) {
	internal.Register[Struct, Parent](godot)
}

// As attempts to cast the given class to T, returning true
// if the cast was successful.
func As[T internal.IsClass](godot Context, class internal.IsClass) (T, bool) {
	return internal.As[T](godot, class)
}

type (
	Radians = internal.Radians
)

type (
	Bool   = internal.Bool
	Int    = internal.Int
	Float  = internal.Float
	String = internal.String

	Vector2     = internal.Vector2
	Vector2i    = internal.Vector2i
	Rect2       = internal.Rect2
	Rect2i      = internal.Rect2i
	Vector3     = internal.Vector3
	Vector3i    = internal.Vector3i
	Transform2D = internal.Transform2D
	Vector4     = internal.Vector4
	Vector4i    = internal.Vector4i
	Plane       = internal.Plane
	Quaternion  = internal.Quaternion
	AABB        = internal.AABB
	Basis       = internal.Basis
	Transform3D = internal.Transform3D
	Projection  = internal.Projection

	Color      = internal.Color
	StringName = internal.StringName
	NodePath   = internal.NodePath
	RID        = internal.RID
	Object     = internal.Object
	Callable   = internal.Callable

	// Signal's T must be a function type.
	Signal[T any] struct {
		internal.Signal

		Emit T
	}

	Dictionary = internal.Dictionary
	Array      = internal.Array

	PackedByteArray    = internal.PackedByteArray
	PackedInt32Array   = internal.PackedInt32Array
	PackedInt64Array   = internal.PackedInt64Array
	PackedFloat32Array = internal.PackedFloat32Array
	PackedFloat64Array = internal.PackedFloat64Array
	PackedStringArray  = internal.PackedStringArray
	PackedVector2Array = internal.PackedVector2Array
	PackedVector3Array = internal.PackedVector3Array
	PackedColorArray   = internal.PackedColorArray
)
