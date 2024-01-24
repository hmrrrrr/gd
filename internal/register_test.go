//go:build !generate

package gd_test

import (
	"testing"

	gd "grow.graphics/gd/internal"
)

func TestRegister(t *testing.T) {
	godot := gd.NewContext(API)
	defer godot.End()

	type SimpleClass struct {
		gd.Class[SimpleClass, gd.Node2D]
	}

	gd.Register[SimpleClass](godot)

	if tag := godot.API.ClassDB.GetClassTag(godot.StringName("Node2D")); tag == 0 {
		t.Fail()
	}
	if tag := godot.API.ClassDB.GetClassTag(godot.StringName("SimpleClass")); tag == 0 {
		t.Fail()
	}
}
