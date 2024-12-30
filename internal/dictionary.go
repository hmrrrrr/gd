package gd

import (
	"graphics.gd/internal/callframe"
	"graphics.gd/internal/pointers"
)

func (d Dictionary) Index(key Variant) Variant {
	return Global.Dictionary.Index(d, key)
}

func (d Dictionary) SetIndex(key Variant, value Variant) {
	Global.Dictionary.SetIndex(d, key, value)
}

func (d Dictionary) Free() {
	if ptr, ok := pointers.End(d); ok {
		var frame = callframe.New()
		Global.typeset.destruct.Dictionary(callframe.Arg(frame, ptr).Uintptr())
		frame.Free()
	}
}

func NewDictionary() Dictionary {
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	Global.typeset.creation.Dictionary[0](r_ret.Uintptr(), callframe.Args{})
	var raw = r_ret.Get()
	frame.Free()
	return pointers.New[Dictionary]([1]uintptr{raw})
}
