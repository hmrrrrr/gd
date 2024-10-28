package ResourceImporterOggVorbis

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/discreet"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/ResourceImporter"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ = discreet.Root

/*
Ogg Vorbis is a lossy audio format, with better audio quality compared to [ResourceImporterMP3] at a given bitrate.
In most cases, it's recommended to use Ogg Vorbis over MP3. However, if you're using an MP3 sound source with no higher quality source available, then it's recommended to use the MP3 file directly to avoid double lossy compression.
Ogg Vorbis requires more CPU to decode than [ResourceImporterWAV]. If you need to play a lot of simultaneous sounds, it's recommended to use WAV for those sounds instead, especially if targeting low-end devices.

*/
type Go [1]classdb.ResourceImporterOggVorbis

/*
This method loads audio data from a PackedByteArray buffer into an AudioStreamOggVorbis object.
*/
func (self Go) LoadFromBuffer(buffer []byte) gdclass.AudioStreamOggVorbis {
	return gdclass.AudioStreamOggVorbis(class(self).LoadFromBuffer(gd.NewPackedByteSlice(buffer)))
}

/*
This method loads audio data from a file into an AudioStreamOggVorbis object. The file path is provided as a string.
*/
func (self Go) LoadFromFile(path string) gdclass.AudioStreamOggVorbis {
	return gdclass.AudioStreamOggVorbis(class(self).LoadFromFile(gd.NewString(path)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ResourceImporterOggVorbis
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }
func New() Go {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ResourceImporterOggVorbis"))
	return Go{classdb.ResourceImporterOggVorbis(object)}
}

/*
This method loads audio data from a PackedByteArray buffer into an AudioStreamOggVorbis object.
*/
//go:nosplit
func (self class) LoadFromBuffer(buffer gd.PackedByteArray) gdclass.AudioStreamOggVorbis {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(buffer))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceImporterOggVorbis.Bind_load_from_buffer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.AudioStreamOggVorbis{classdb.AudioStreamOggVorbis(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
/*
This method loads audio data from a file into an AudioStreamOggVorbis object. The file path is provided as a string.
*/
//go:nosplit
func (self class) LoadFromFile(path gd.String) gdclass.AudioStreamOggVorbis {
	var frame = callframe.New()
	callframe.Arg(frame, discreet.Get(path))
	var r_ret = callframe.Ret[[1]uintptr](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ResourceImporterOggVorbis.Bind_load_from_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = gdclass.AudioStreamOggVorbis{classdb.AudioStreamOggVorbis(gd.PointerWithOwnershipTransferredToGo(r_ret.Get()))}
	frame.Free()
	return ret
}
func (self class) AsResourceImporterOggVorbis() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsResourceImporterOggVorbis() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResourceImporter() ResourceImporter.GD { return *((*ResourceImporter.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResourceImporter() ResourceImporter.Go { return *((*ResourceImporter.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResourceImporter(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResourceImporter(), name)
	}
}
func init() {classdb.Register("ResourceImporterOggVorbis", func(ptr gd.Object) any { return classdb.ResourceImporterOggVorbis(ptr) })}
