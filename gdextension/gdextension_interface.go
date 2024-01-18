package gdextension

/*
#include "gdextension_interface.h"

uintptr_t classdb_construct_object(uintptr_t fn, uintptr_t p_classname) {
	return (uintptr_t)((GDExtensionInterfaceClassdbConstructObject)fn)((GDExtensionConstStringNamePtr)p_classname);
}
void variant_get(uintptr_t fn, uintptr_t p_self, uintptr_t p_key, uintptr_t r_ret, uintptr_t r_valid) {
	((GDExtensionInterfaceVariantGet)fn)((GDExtensionConstVariantPtr)p_self, (GDExtensionConstVariantPtr)p_key, (GDExtensionUninitializedVariantPtr)r_ret, (GDExtensionBool*)r_valid);
}
void variant_set(uintptr_t fn, uintptr_t p_self, uintptr_t p_key, uintptr_t p_value, uintptr_t r_valid) {
	((GDExtensionInterfaceVariantSet)fn)((GDExtensionVariantPtr)p_self, (GDExtensionConstVariantPtr)p_key, (GDExtensionConstVariantPtr)p_value, (GDExtensionBool*)r_valid);
}

*/
import "C"

import (
	"grow.graphics/gd"
	internal "grow.graphics/gd/internal"
	"runtime.link/api/call"
	"runtime.link/mmm"
)

func linkCGO(API *internal.API) {
	classdb_construct_object := dlsymGD("classdb_construct_object")
	API.ClassDB.ConstructObject = func(ctx gd.Context, name gd.StringName) gd.Object {
		var frame = call.New()
		var obj gd.Object
		obj.SetPointer(mmm.Make[internal.API, internal.Pointer, uintptr](ctx, ctx.API(), uintptr(C.classdb_construct_object(
			C.uintptr_t(uintptr(classdb_construct_object)),
			C.uintptr_t(call.Arg(frame, name.Pointer()).Uintptr()),
		))))
		frame.Free()
		return obj
	}

	variant_get := dlsymGD("variant_get")
	API.Variants.Get = func(ctx internal.Context, self, key internal.Variant) (internal.Variant, bool) {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_key = call.Arg(frame, key.Pointer())
		var r_ret = call.Ret[[3]uintptr](frame)
		var r_valid = call.Ret[bool](frame)
		C.variant_get(
			C.uintptr_t(uintptr(variant_get)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_key.Uintptr()),
			C.uintptr_t(r_ret.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		return mmm.Make[internal.API, internal.Variant, [3]uintptr](ctx, ctx.API(), r_ret.Get()), r_valid.Get()
	}
	variant_set := dlsymGD("variant_set")
	API.Variants.Set = func(ctx internal.Context, self, key, val internal.Variant) bool {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_key = call.Arg(frame, key.Pointer())
		var p_value = call.Arg(frame, val.Pointer())
		var r_valid = call.Ret[bool](frame)
		C.variant_set(
			C.uintptr_t(uintptr(variant_set)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_key.Uintptr()),
			C.uintptr_t(p_value.Uintptr()),
			C.uintptr_t(r_valid.Uintptr()),
		)
		return r_valid.Get()
	}
}
