package gdextension

/*
#include <stdlib.h>
#include "gdextension_interface.h"

void get_godot_version(uintptr_t fn, GDExtensionGodotVersion *r_version) {
	((GDExtensionInterfaceGetGodotVersion)fn)(r_version);
}
void *mem_alloc(uintptr_t fn, size_t size) {
	return ((GDExtensionInterfaceMemAlloc)fn)(size);
}
void *mem_realloc(uintptr_t fn, void *ptr, size_t size) {
	return ((GDExtensionInterfaceMemRealloc)fn)(ptr, size);
}
void mem_free(uintptr_t fn, void *ptr) {
	((GDExtensionInterfaceMemFree)fn)(ptr);
}
void print_error(uintptr_t fn, const char *p_description, const char *p_function, const char *p_file, int32_t p_line, GDExtensionBool p_notify_editor) {
	((GDExtensionInterfacePrintError)fn)(p_description, p_function, p_file, p_line, p_notify_editor);
}
void print_error_with_message(uintptr_t fn, const char *p_description, const char *p_message, const char *p_function, const char *p_file, int32_t p_line, GDExtensionBool p_notify_editor) {
	((GDExtensionInterfacePrintErrorWithMessage)fn)(p_description, p_message, p_function, p_file, p_line, p_notify_editor);
}
void print_warning(uintptr_t fn, const char *p_description, const char *p_function, const char *p_file, int32_t p_line, GDExtensionBool p_notify_editor) {
	((GDExtensionInterfacePrintWarning)fn)(p_description, p_function, p_file, p_line, p_notify_editor);
}
void variant_new_copy(uintptr_t fn, uintptr_t r_dest, uintptr_t p_self) {
	((GDExtensionInterfaceVariantNewCopy)fn)((GDExtensionUninitializedVariantPtr)r_dest, (GDExtensionConstVariantPtr)p_self);
}
void variant_new_nil(uintptr_t fn, uintptr_t r_dest) {
	((GDExtensionInterfaceVariantNewNil)fn)((GDExtensionUninitializedVariantPtr)r_dest);
}
void variant_destroy(uintptr_t fn, uintptr_t p_self) {
	((GDExtensionInterfaceVariantDestroy)fn)((GDExtensionVariantPtr)p_self);
}
void variant_call(uintptr_t fn, uintptr_t p_self, uintptr_t p_method, uintptr_t p_args, GDExtensionInt p_argument_count, uintptr_t r_ret, GDExtensionCallError *r_error) {
	((GDExtensionInterfaceVariantCall)fn)((GDExtensionVariantPtr)p_self, (GDExtensionConstStringNamePtr)p_method, (GDExtensionConstVariantPtr)p_args, (GDExtensionInt)p_argument_count, (GDExtensionUninitializedVariantPtr)r_ret, r_error);
}
void variant_call_static(uintptr_t fn, GDExtensionVariantType p_type, uintptr_t p_method, uintptr_t p_args, GDExtensionInt p_argument_count, uintptr_t r_ret, GDExtensionCallError *r_error) {
	((GDExtensionInterfaceVariantCallStatic)fn)((GDExtensionVariantType)p_type, (GDExtensionConstStringNamePtr)p_method, (GDExtensionConstVariantPtr)p_args, (GDExtensionInt)p_argument_count, (GDExtensionUninitializedVariantPtr)r_ret, r_error);
}

uint64_t get_native_struct_size(uintptr_t fn, uintptr_t p_classname) {
	return ((GDExtensionInterfaceGetNativeStructSize)fn)((GDExtensionConstStringNamePtr)p_classname);
}

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
	"unsafe"

	gd "grow.graphics/gd/internal"
	"runtime.link/api/call"
	"runtime.link/mmm"
)

// linkCGO implements the Godot GDExtension API via CGO.
func linkCGO(API *gd.API) {
	get_godot_version := dlsymGD("get_godot_version")
	API.GetGodotVersion = func() gd.Version {
		var version = new(C.GDExtensionGodotVersion)
		C.get_godot_version(C.uintptr_t(uintptr(get_godot_version)), version)
		return gd.Version{
			Major: uint32(version.major),
			Minor: uint32(version.minor),
			Patch: uint32(version.patch),
			Value: C.GoString(version.string),
		}
	}
	mem_alloc := dlsymGD("mem_alloc")
	API.Memory.Allocate = func(size uintptr) unsafe.Pointer {
		return unsafe.Pointer(C.mem_alloc(C.uintptr_t(uintptr(mem_alloc)), C.size_t(size)))
	}
	mem_realloc := dlsymGD("mem_realloc")
	API.Memory.Reallocate = func(ptr unsafe.Pointer, size uintptr) unsafe.Pointer {
		return unsafe.Pointer(C.mem_realloc(C.uintptr_t(uintptr(mem_realloc)), ptr, C.size_t(size)))
	}
	mem_free := dlsymGD("mem_free")
	API.Memory.Free = func(ptr unsafe.Pointer) {
		C.mem_free(C.uintptr_t(uintptr(mem_free)), ptr)
	}
	print_error := dlsymGD("print_error")
	API.PrintError = func(code, function, file string, line int32, notifyEditor bool) {
		p_description := C.CString(code)
		p_function := C.CString(function)
		p_file := C.CString(file)
		p_editor_notify := C.GDExtensionBool(0)
		if notifyEditor {
			p_editor_notify = 1
		}
		C.print_error(
			C.uintptr_t(uintptr(print_error)),
			p_description,
			p_function,
			p_file,
			C.int32_t(line),
			p_editor_notify,
		)
		C.free(unsafe.Pointer(p_description))
		C.free(unsafe.Pointer(p_function))
		C.free(unsafe.Pointer(p_file))
	}
	print_error_with_message := dlsymGD("print_error_with_message")
	API.PrintErrorMessage = func(code, message, function, file string, line int32, notifyEditor bool) {
		p_description := C.CString(code)
		p_message := C.CString(message)
		p_function := C.CString(function)
		p_file := C.CString(file)
		p_editor_notify := C.GDExtensionBool(0)
		if notifyEditor {
			p_editor_notify = 1
		}
		C.print_error_with_message(
			C.uintptr_t(uintptr(print_error_with_message)),
			p_description,
			p_message,
			p_function,
			p_file,
			C.int32_t(line),
			p_editor_notify,
		)
		C.free(unsafe.Pointer(p_description))
		C.free(unsafe.Pointer(p_message))
		C.free(unsafe.Pointer(p_function))
		C.free(unsafe.Pointer(p_file))
	}
	print_warning := dlsymGD("print_warning")
	API.PrintWarning = func(code, function, file string, line int32, notifyEditor bool) {
		p_description := C.CString(code)
		p_function := C.CString(function)
		p_file := C.CString(file)
		p_editor_notify := C.GDExtensionBool(0)
		if notifyEditor {
			p_editor_notify = 1
		}
		C.print_warning(
			C.uintptr_t(uintptr(print_warning)),
			p_description,
			p_function,
			p_file,
			C.int32_t(line),
			p_editor_notify,
		)
		C.free(unsafe.Pointer(p_description))
		C.free(unsafe.Pointer(p_function))
		C.free(unsafe.Pointer(p_file))
	}
	print_script_error := dlsymGD("print_script_error")
	API.PrintScriptError = func(code, function, file string, line int32, notifyEditor bool) {
		p_description := C.CString(code)
		p_function := C.CString(function)
		p_file := C.CString(file)
		p_editor_notify := C.GDExtensionBool(0)
		if notifyEditor {
			p_editor_notify = 1
		}
		C.print_error(
			C.uintptr_t(uintptr(print_script_error)),
			p_description,
			p_function,
			p_file,
			C.int32_t(line),
			p_editor_notify,
		)
		C.free(unsafe.Pointer(p_description))
		C.free(unsafe.Pointer(p_function))
		C.free(unsafe.Pointer(p_file))
	}
	print_script_error_with_message := dlsymGD("print_script_error_with_message")
	API.PrintScriptErrorMessage = func(code, message, function, file string, line int32, notifyEditor bool) {
		p_description := C.CString(code)
		p_message := C.CString(message)
		p_function := C.CString(function)
		p_file := C.CString(file)
		p_editor_notify := C.GDExtensionBool(0)
		if notifyEditor {
			p_editor_notify = 1
		}
		C.print_error_with_message(
			C.uintptr_t(uintptr(print_script_error_with_message)),
			p_description,
			p_message,
			p_function,
			p_file,
			C.int32_t(line),
			p_editor_notify,
		)
		C.free(unsafe.Pointer(p_description))
		C.free(unsafe.Pointer(p_message))
		C.free(unsafe.Pointer(p_function))
		C.free(unsafe.Pointer(p_file))
	}
	variant_new_copy := dlsymGD("variant_new_copy")
	API.Variants.NewCopy = func(ctx gd.Context, self gd.Variant) gd.Variant {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var r_dest = call.Ret[[3]uintptr](frame)
		C.variant_new_copy(
			C.uintptr_t(uintptr(variant_new_copy)),
			C.uintptr_t(r_dest.Uintptr()),
			C.uintptr_t(p_self.Uintptr()),
		)
		var ret = mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_dest.Get())
		frame.Free()
		return ret
	}
	variant_new_nil := dlsymGD("variant_new_nil")
	API.Variants.NewNil = func(ctx gd.Context) gd.Variant {
		var frame = call.New()
		var r_dest = call.Ret[[3]uintptr](frame)
		C.variant_new_nil(
			C.uintptr_t(uintptr(variant_new_nil)),
			C.uintptr_t(r_dest.Uintptr()),
		)
		var ret = mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_dest.Get())
		frame.Free()
		return ret
	}
	variant_destroy := dlsymGD("variant_destroy")
	API.Variants.Destroy = func(self gd.Variant) {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		C.variant_destroy(
			C.uintptr_t(uintptr(variant_destroy)),
			C.uintptr_t(p_self.Uintptr()),
		)
		frame.Free()
	}
	variant_call := dlsymGD("variant_call")
	API.Variants.Call = func(ctx gd.Context, self gd.Variant, method gd.StringName, args ...gd.Variant) (gd.Variant, error) {
		var frame = call.New()
		var p_self = call.Arg(frame, self.Pointer())
		var p_method = call.Arg(frame, method.Pointer())
		for _, arg := range args {
			call.Arg(frame, arg.Pointer())
		}
		var r_ret = call.Ret[[3]uintptr](frame)
		var r_error = new(C.GDExtensionCallError)
		C.variant_call(
			C.uintptr_t(uintptr(variant_call)),
			C.uintptr_t(p_self.Uintptr()),
			C.uintptr_t(p_method.Uintptr()),
			C.uintptr_t(frame.Array(2).Uintptr()),
			C.GDExtensionInt(len(args)),
			C.uintptr_t(r_ret.Uintptr()),
			r_error,
		)
		if r_error.error != 0 {
			frame.Free()
			return gd.Variant{}, &gd.CallError{
				ErrorType: gd.CallErrorType(r_error.error),
				Argument:  int32(r_error.argument),
				Expected:  int32(r_error.expected),
			}
		}
		var ret = mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_ret.Get())
		frame.Free()
		return ret, nil
	}
	API.Variants.CallStatic = func(ctx gd.Context, vtype gd.VariantType, method gd.StringName, args ...gd.Variant) (gd.Variant, error) {
		var frame = call.New()
		var p_method = call.Arg(frame, method.Pointer())
		for _, arg := range args {
			call.Arg(frame, arg.Pointer())
		}
		var r_ret = call.Ret[[3]uintptr](frame)
		var r_error = new(C.GDExtensionCallError)
		C.variant_call_static(
			C.uintptr_t(uintptr(variant_call)),
			C.GDExtensionVariantType(vtype),
			C.uintptr_t(p_method.Uintptr()),
			C.uintptr_t(frame.Array(1).Uintptr()),
			C.GDExtensionInt(len(args)),
			C.uintptr_t(r_ret.Uintptr()),
			r_error,
		)
		if r_error.error != 0 {
			frame.Free()
			return gd.Variant{}, &gd.CallError{
				ErrorType: gd.CallErrorType(r_error.error),
				Argument:  int32(r_error.argument),
				Expected:  int32(r_error.expected),
			}
		}
		var ret = mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_ret.Get())
		frame.Free()
		return ret, nil
	}

	get_native_struct_size := dlsymGD("get_native_struct_size")
	API.GetNativeStructSize = func(name gd.StringName) uint64 {
		var frame = call.New()
		defer frame.Free()
		return uint64(C.get_native_struct_size(
			C.uintptr_t(uintptr(get_native_struct_size)),
			C.uintptr_t(call.Arg(frame, name.Pointer()).Uintptr()),
		))
	}

	classdb_construct_object := dlsymGD("classdb_construct_object")
	API.ClassDB.ConstructObject = func(ctx gd.Context, name gd.StringName) gd.Object {
		var frame = call.New()
		var obj gd.Object
		obj.SetPointer(mmm.Make[gd.API, gd.Pointer, uintptr](ctx, ctx.API(), uintptr(C.classdb_construct_object(
			C.uintptr_t(uintptr(classdb_construct_object)),
			C.uintptr_t(call.Arg(frame, name.Pointer()).Uintptr()),
		))))
		frame.Free()
		return obj
	}

	variant_get := dlsymGD("variant_get")
	API.Variants.Get = func(ctx gd.Context, self, key gd.Variant) (gd.Variant, bool) {
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
		frame.Free()
		return mmm.Make[gd.API, gd.Variant, [3]uintptr](ctx, ctx.API(), r_ret.Get()), r_valid.Get()
	}
	variant_set := dlsymGD("variant_set")
	API.Variants.Set = func(ctx gd.Context, self, key, val gd.Variant) bool {
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
		frame.Free()
		return r_valid.Get()
	}
}
