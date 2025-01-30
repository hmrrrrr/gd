// Package FontFile provides methods for working with FontFile object instances.
package FontFile

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Font"
import "graphics.gd/classdb/Resource"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/Rect2"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Transform2D"
import "graphics.gd/variant/Vector2"
import "graphics.gd/variant/Vector2i"

var _ Object.ID
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ = slices.Delete[[]struct{}, struct{}]

/*
[FontFile] contains a set of glyphs to represent Unicode characters imported from a font file, as well as a cache of rasterized glyphs, and a set of fallback [Font]s to use.
Use [FontVariation] to access specific OpenType variation of the font, create simulated bold / slanted version, and draw lines of text.
For more complex text processing, use [FontVariation] in conjunction with [TextLine] or [TextParagraph].
Supported font formats:
- Dynamic font importer: TrueType (.ttf), TrueType collection (.ttc), OpenType (.otf), OpenType collection (.otc), WOFF (.woff), WOFF2 (.woff2), Type 1 (.pfb, .pfm).
- Bitmap font importer: AngelCode BMFont (.fnt, .font), text and binary (version 3) format variants.
- Monospace image font importer: All supported image formats.
[b]Note:[/b] A character is a symbol that represents an item (letter, digit etc.) in an abstract way.
[b]Note:[/b] A glyph is a bitmap or a shape used to draw one or more characters in a context-dependent manner. Glyph indices are bound to the specific font data source.
[b]Note:[/b] If none of the font data sources contain glyphs for a character used in a string, the character in question will be replaced with a box displaying its hexadecimal code.
[codeblocks]
[gdscript]
var f = load("res://BarlowCondensed-Bold.ttf")
$Label.add_theme_font_override("font", f)
$Label.add_theme_font_size_override("font_size", 64)
[/gdscript]
[csharp]
var f = ResourceLoader.Load<FontFile>("res://BarlowCondensed-Bold.ttf");
GetNode("Label").AddThemeFontOverride("font", f);
GetNode("Label").AddThemeFontSizeOverride("font_size", 64);
[/csharp]
[/codeblocks]
*/
type Instance [1]gdclass.FontFile

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsFontFile() Instance
}

/*
Loads an AngelCode BMFont (.fnt, .font) bitmap font from file [param path].
[b]Warning:[/b] This method should only be used in the editor or in cases when you need to load external fonts at run-time, such as fonts located at the [code]user://[/code] directory.
*/
func (self Instance) LoadBitmapFont(path string) error { //gd:FontFile.load_bitmap_font
	return error(gd.ToError(class(self).LoadBitmapFont(String.New(path))))
}

/*
Loads a TrueType (.ttf), OpenType (.otf), WOFF (.woff), WOFF2 (.woff2) or Type 1 (.pfb, .pfm) dynamic font from file [param path].
[b]Warning:[/b] This method should only be used in the editor or in cases when you need to load external fonts at run-time, such as fonts located at the [code]user://[/code] directory.
*/
func (self Instance) LoadDynamicFont(path string) error { //gd:FontFile.load_dynamic_font
	return error(gd.ToError(class(self).LoadDynamicFont(String.New(path))))
}

/*
Returns number of the font cache entries.
*/
func (self Instance) GetCacheCount() int { //gd:FontFile.get_cache_count
	return int(int(class(self).GetCacheCount()))
}

/*
Removes all font cache entries.
*/
func (self Instance) ClearCache() { //gd:FontFile.clear_cache
	class(self).ClearCache()
}

/*
Removes specified font cache entry.
*/
func (self Instance) RemoveCache(cache_index int) { //gd:FontFile.remove_cache
	class(self).RemoveCache(int64(cache_index))
}

/*
Returns list of the font sizes in the cache. Each size is [Vector2i] with font size and outline size.
*/
func (self Instance) GetSizeCacheList(cache_index int) []Vector2i.XY { //gd:FontFile.get_size_cache_list
	return []Vector2i.XY(gd.ArrayAs[[]Vector2i.XY](gd.InternalArray(class(self).GetSizeCacheList(int64(cache_index)))))
}

/*
Removes all font sizes from the cache entry
*/
func (self Instance) ClearSizeCache(cache_index int) { //gd:FontFile.clear_size_cache
	class(self).ClearSizeCache(int64(cache_index))
}

/*
Removes specified font size from the cache entry.
*/
func (self Instance) RemoveSizeCache(cache_index int, size Vector2i.XY) { //gd:FontFile.remove_size_cache
	class(self).RemoveSizeCache(int64(cache_index), Vector2i.XY(size))
}

/*
Sets variation coordinates for the specified font cache entry. See [method Font.get_supported_variation_list] for more info.
*/
func (self Instance) SetVariationCoordinates(cache_index int, variation_coordinates map[string]float32) { //gd:FontFile.set_variation_coordinates
	class(self).SetVariationCoordinates(int64(cache_index), gd.DictionaryFromMap(variation_coordinates))
}

/*
Returns variation coordinates for the specified font cache entry. See [method Font.get_supported_variation_list] for more info.
*/
func (self Instance) GetVariationCoordinates(cache_index int) map[string]float32 { //gd:FontFile.get_variation_coordinates
	return map[string]float32(gd.DictionaryAs[map[string]float32](class(self).GetVariationCoordinates(int64(cache_index))))
}

/*
Sets embolden strength, if is not equal to zero, emboldens the font outlines. Negative values reduce the outline thickness.
*/
func (self Instance) SetEmbolden(cache_index int, strength Float.X) { //gd:FontFile.set_embolden
	class(self).SetEmbolden(int64(cache_index), float64(strength))
}

/*
Returns embolden strength, if is not equal to zero, emboldens the font outlines. Negative values reduce the outline thickness.
*/
func (self Instance) GetEmbolden(cache_index int) Float.X { //gd:FontFile.get_embolden
	return Float.X(Float.X(class(self).GetEmbolden(int64(cache_index))))
}

/*
Sets 2D transform, applied to the font outlines, can be used for slanting, flipping, and rotating glyphs.
*/
func (self Instance) SetTransform(cache_index int, transform Transform2D.OriginXY) { //gd:FontFile.set_transform
	class(self).SetTransform(int64(cache_index), Transform2D.OriginXY(transform))
}

/*
Returns 2D transform, applied to the font outlines, can be used for slanting, flipping and rotating glyphs.
*/
func (self Instance) GetTransform(cache_index int) Transform2D.OriginXY { //gd:FontFile.get_transform
	return Transform2D.OriginXY(class(self).GetTransform(int64(cache_index)))
}

/*
Sets the spacing for [param spacing] (see [enum TextServer.SpacingType]) to [param value] in pixels (not relative to the font size).
*/
func (self Instance) SetExtraSpacing(cache_index int, spacing gdclass.TextServerSpacingType, value int) { //gd:FontFile.set_extra_spacing
	class(self).SetExtraSpacing(int64(cache_index), spacing, int64(value))
}

/*
Returns spacing for [param spacing] (see [enum TextServer.SpacingType]) in pixels (not relative to the font size).
*/
func (self Instance) GetExtraSpacing(cache_index int, spacing gdclass.TextServerSpacingType) int { //gd:FontFile.get_extra_spacing
	return int(int(class(self).GetExtraSpacing(int64(cache_index), spacing)))
}

/*
Sets extra baseline offset (as a fraction of font height).
*/
func (self Instance) SetExtraBaselineOffset(cache_index int, baseline_offset Float.X) { //gd:FontFile.set_extra_baseline_offset
	class(self).SetExtraBaselineOffset(int64(cache_index), float64(baseline_offset))
}

/*
Returns extra baseline offset (as a fraction of font height).
*/
func (self Instance) GetExtraBaselineOffset(cache_index int) Float.X { //gd:FontFile.get_extra_baseline_offset
	return Float.X(Float.X(class(self).GetExtraBaselineOffset(int64(cache_index))))
}

/*
Sets an active face index in the TrueType / OpenType collection.
*/
func (self Instance) SetFaceIndex(cache_index int, face_index int) { //gd:FontFile.set_face_index
	class(self).SetFaceIndex(int64(cache_index), int64(face_index))
}

/*
Returns an active face index in the TrueType / OpenType collection.
*/
func (self Instance) GetFaceIndex(cache_index int) int { //gd:FontFile.get_face_index
	return int(int(class(self).GetFaceIndex(int64(cache_index))))
}

/*
Sets the font ascent (number of pixels above the baseline).
*/
func (self Instance) SetCacheAscent(cache_index int, size int, ascent Float.X) { //gd:FontFile.set_cache_ascent
	class(self).SetCacheAscent(int64(cache_index), int64(size), float64(ascent))
}

/*
Returns the font ascent (number of pixels above the baseline).
*/
func (self Instance) GetCacheAscent(cache_index int, size int) Float.X { //gd:FontFile.get_cache_ascent
	return Float.X(Float.X(class(self).GetCacheAscent(int64(cache_index), int64(size))))
}

/*
Sets the font descent (number of pixels below the baseline).
*/
func (self Instance) SetCacheDescent(cache_index int, size int, descent Float.X) { //gd:FontFile.set_cache_descent
	class(self).SetCacheDescent(int64(cache_index), int64(size), float64(descent))
}

/*
Returns the font descent (number of pixels below the baseline).
*/
func (self Instance) GetCacheDescent(cache_index int, size int) Float.X { //gd:FontFile.get_cache_descent
	return Float.X(Float.X(class(self).GetCacheDescent(int64(cache_index), int64(size))))
}

/*
Sets pixel offset of the underline below the baseline.
*/
func (self Instance) SetCacheUnderlinePosition(cache_index int, size int, underline_position Float.X) { //gd:FontFile.set_cache_underline_position
	class(self).SetCacheUnderlinePosition(int64(cache_index), int64(size), float64(underline_position))
}

/*
Returns pixel offset of the underline below the baseline.
*/
func (self Instance) GetCacheUnderlinePosition(cache_index int, size int) Float.X { //gd:FontFile.get_cache_underline_position
	return Float.X(Float.X(class(self).GetCacheUnderlinePosition(int64(cache_index), int64(size))))
}

/*
Sets thickness of the underline in pixels.
*/
func (self Instance) SetCacheUnderlineThickness(cache_index int, size int, underline_thickness Float.X) { //gd:FontFile.set_cache_underline_thickness
	class(self).SetCacheUnderlineThickness(int64(cache_index), int64(size), float64(underline_thickness))
}

/*
Returns thickness of the underline in pixels.
*/
func (self Instance) GetCacheUnderlineThickness(cache_index int, size int) Float.X { //gd:FontFile.get_cache_underline_thickness
	return Float.X(Float.X(class(self).GetCacheUnderlineThickness(int64(cache_index), int64(size))))
}

/*
Sets scaling factor of the color bitmap font.
*/
func (self Instance) SetCacheScale(cache_index int, size int, scale Float.X) { //gd:FontFile.set_cache_scale
	class(self).SetCacheScale(int64(cache_index), int64(size), float64(scale))
}

/*
Returns scaling factor of the color bitmap font.
*/
func (self Instance) GetCacheScale(cache_index int, size int) Float.X { //gd:FontFile.get_cache_scale
	return Float.X(Float.X(class(self).GetCacheScale(int64(cache_index), int64(size))))
}

/*
Returns number of textures used by font cache entry.
*/
func (self Instance) GetTextureCount(cache_index int, size Vector2i.XY) int { //gd:FontFile.get_texture_count
	return int(int(class(self).GetTextureCount(int64(cache_index), Vector2i.XY(size))))
}

/*
Removes all textures from font cache entry.
[b]Note:[/b] This function will not remove glyphs associated with the texture, use [method remove_glyph] to remove them manually.
*/
func (self Instance) ClearTextures(cache_index int, size Vector2i.XY) { //gd:FontFile.clear_textures
	class(self).ClearTextures(int64(cache_index), Vector2i.XY(size))
}

/*
Removes specified texture from the cache entry.
[b]Note:[/b] This function will not remove glyphs associated with the texture. Remove them manually using [method remove_glyph].
*/
func (self Instance) RemoveTexture(cache_index int, size Vector2i.XY, texture_index int) { //gd:FontFile.remove_texture
	class(self).RemoveTexture(int64(cache_index), Vector2i.XY(size), int64(texture_index))
}

/*
Sets font cache texture image.
*/
func (self Instance) SetTextureImage(cache_index int, size Vector2i.XY, texture_index int, image [1]gdclass.Image) { //gd:FontFile.set_texture_image
	class(self).SetTextureImage(int64(cache_index), Vector2i.XY(size), int64(texture_index), image)
}

/*
Returns a copy of the font cache texture image.
*/
func (self Instance) GetTextureImage(cache_index int, size Vector2i.XY, texture_index int) [1]gdclass.Image { //gd:FontFile.get_texture_image
	return [1]gdclass.Image(class(self).GetTextureImage(int64(cache_index), Vector2i.XY(size), int64(texture_index)))
}

/*
Sets array containing glyph packing data.
*/
func (self Instance) SetTextureOffsets(cache_index int, size Vector2i.XY, texture_index int, offset []int32) { //gd:FontFile.set_texture_offsets
	class(self).SetTextureOffsets(int64(cache_index), Vector2i.XY(size), int64(texture_index), Packed.New(offset...))
}

/*
Returns a copy of the array containing glyph packing data.
*/
func (self Instance) GetTextureOffsets(cache_index int, size Vector2i.XY, texture_index int) []int32 { //gd:FontFile.get_texture_offsets
	return []int32(slices.Collect(class(self).GetTextureOffsets(int64(cache_index), Vector2i.XY(size), int64(texture_index)).Values()))
}

/*
Returns list of rendered glyphs in the cache entry.
*/
func (self Instance) GetGlyphList(cache_index int, size Vector2i.XY) []int32 { //gd:FontFile.get_glyph_list
	return []int32(slices.Collect(class(self).GetGlyphList(int64(cache_index), Vector2i.XY(size)).Values()))
}

/*
Removes all rendered glyph information from the cache entry.
[b]Note:[/b] This function will not remove textures associated with the glyphs, use [method remove_texture] to remove them manually.
*/
func (self Instance) ClearGlyphs(cache_index int, size Vector2i.XY) { //gd:FontFile.clear_glyphs
	class(self).ClearGlyphs(int64(cache_index), Vector2i.XY(size))
}

/*
Removes specified rendered glyph information from the cache entry.
[b]Note:[/b] This function will not remove textures associated with the glyphs, use [method remove_texture] to remove them manually.
*/
func (self Instance) RemoveGlyph(cache_index int, size Vector2i.XY, glyph int) { //gd:FontFile.remove_glyph
	class(self).RemoveGlyph(int64(cache_index), Vector2i.XY(size), int64(glyph))
}

/*
Sets glyph advance (offset of the next glyph).
[b]Note:[/b] Advance for glyphs outlines is the same as the base glyph advance and is not saved.
*/
func (self Instance) SetGlyphAdvance(cache_index int, size int, glyph int, advance Vector2.XY) { //gd:FontFile.set_glyph_advance
	class(self).SetGlyphAdvance(int64(cache_index), int64(size), int64(glyph), Vector2.XY(advance))
}

/*
Returns glyph advance (offset of the next glyph).
[b]Note:[/b] Advance for glyphs outlines is the same as the base glyph advance and is not saved.
*/
func (self Instance) GetGlyphAdvance(cache_index int, size int, glyph int) Vector2.XY { //gd:FontFile.get_glyph_advance
	return Vector2.XY(class(self).GetGlyphAdvance(int64(cache_index), int64(size), int64(glyph)))
}

/*
Sets glyph offset from the baseline.
*/
func (self Instance) SetGlyphOffset(cache_index int, size Vector2i.XY, glyph int, offset Vector2.XY) { //gd:FontFile.set_glyph_offset
	class(self).SetGlyphOffset(int64(cache_index), Vector2i.XY(size), int64(glyph), Vector2.XY(offset))
}

/*
Returns glyph offset from the baseline.
*/
func (self Instance) GetGlyphOffset(cache_index int, size Vector2i.XY, glyph int) Vector2.XY { //gd:FontFile.get_glyph_offset
	return Vector2.XY(class(self).GetGlyphOffset(int64(cache_index), Vector2i.XY(size), int64(glyph)))
}

/*
Sets glyph size.
*/
func (self Instance) SetGlyphSize(cache_index int, size Vector2i.XY, glyph int, gl_size Vector2.XY) { //gd:FontFile.set_glyph_size
	class(self).SetGlyphSize(int64(cache_index), Vector2i.XY(size), int64(glyph), Vector2.XY(gl_size))
}

/*
Returns glyph size.
*/
func (self Instance) GetGlyphSize(cache_index int, size Vector2i.XY, glyph int) Vector2.XY { //gd:FontFile.get_glyph_size
	return Vector2.XY(class(self).GetGlyphSize(int64(cache_index), Vector2i.XY(size), int64(glyph)))
}

/*
Sets rectangle in the cache texture containing the glyph.
*/
func (self Instance) SetGlyphUvRect(cache_index int, size Vector2i.XY, glyph int, uv_rect Rect2.PositionSize) { //gd:FontFile.set_glyph_uv_rect
	class(self).SetGlyphUvRect(int64(cache_index), Vector2i.XY(size), int64(glyph), Rect2.PositionSize(uv_rect))
}

/*
Returns rectangle in the cache texture containing the glyph.
*/
func (self Instance) GetGlyphUvRect(cache_index int, size Vector2i.XY, glyph int) Rect2.PositionSize { //gd:FontFile.get_glyph_uv_rect
	return Rect2.PositionSize(class(self).GetGlyphUvRect(int64(cache_index), Vector2i.XY(size), int64(glyph)))
}

/*
Sets index of the cache texture containing the glyph.
*/
func (self Instance) SetGlyphTextureIdx(cache_index int, size Vector2i.XY, glyph int, texture_idx int) { //gd:FontFile.set_glyph_texture_idx
	class(self).SetGlyphTextureIdx(int64(cache_index), Vector2i.XY(size), int64(glyph), int64(texture_idx))
}

/*
Returns index of the cache texture containing the glyph.
*/
func (self Instance) GetGlyphTextureIdx(cache_index int, size Vector2i.XY, glyph int) int { //gd:FontFile.get_glyph_texture_idx
	return int(int(class(self).GetGlyphTextureIdx(int64(cache_index), Vector2i.XY(size), int64(glyph))))
}

/*
Returns list of the kerning overrides.
*/
func (self Instance) GetKerningList(cache_index int, size int) []Vector2i.XY { //gd:FontFile.get_kerning_list
	return []Vector2i.XY(gd.ArrayAs[[]Vector2i.XY](gd.InternalArray(class(self).GetKerningList(int64(cache_index), int64(size)))))
}

/*
Removes all kerning overrides.
*/
func (self Instance) ClearKerningMap(cache_index int, size int) { //gd:FontFile.clear_kerning_map
	class(self).ClearKerningMap(int64(cache_index), int64(size))
}

/*
Removes kerning override for the pair of glyphs.
*/
func (self Instance) RemoveKerning(cache_index int, size int, glyph_pair Vector2i.XY) { //gd:FontFile.remove_kerning
	class(self).RemoveKerning(int64(cache_index), int64(size), Vector2i.XY(glyph_pair))
}

/*
Sets kerning for the pair of glyphs.
*/
func (self Instance) SetKerning(cache_index int, size int, glyph_pair Vector2i.XY, kerning Vector2.XY) { //gd:FontFile.set_kerning
	class(self).SetKerning(int64(cache_index), int64(size), Vector2i.XY(glyph_pair), Vector2.XY(kerning))
}

/*
Returns kerning for the pair of glyphs.
*/
func (self Instance) GetKerning(cache_index int, size int, glyph_pair Vector2i.XY) Vector2.XY { //gd:FontFile.get_kerning
	return Vector2.XY(class(self).GetKerning(int64(cache_index), int64(size), Vector2i.XY(glyph_pair)))
}

/*
Renders the range of characters to the font cache texture.
*/
func (self Instance) RenderRange(cache_index int, size Vector2i.XY, start int, end int) { //gd:FontFile.render_range
	class(self).RenderRange(int64(cache_index), Vector2i.XY(size), int64(start), int64(end))
}

/*
Renders specified glyph to the font cache texture.
*/
func (self Instance) RenderGlyph(cache_index int, size Vector2i.XY, index int) { //gd:FontFile.render_glyph
	class(self).RenderGlyph(int64(cache_index), Vector2i.XY(size), int64(index))
}

/*
Adds override for [method Font.is_language_supported].
*/
func (self Instance) SetLanguageSupportOverride(language string, supported bool) { //gd:FontFile.set_language_support_override
	class(self).SetLanguageSupportOverride(String.New(language), supported)
}

/*
Returns [code]true[/code] if support override is enabled for the [param language].
*/
func (self Instance) GetLanguageSupportOverride(language string) bool { //gd:FontFile.get_language_support_override
	return bool(class(self).GetLanguageSupportOverride(String.New(language)))
}

/*
Remove language support override.
*/
func (self Instance) RemoveLanguageSupportOverride(language string) { //gd:FontFile.remove_language_support_override
	class(self).RemoveLanguageSupportOverride(String.New(language))
}

/*
Returns list of language support overrides.
*/
func (self Instance) GetLanguageSupportOverrides() []string { //gd:FontFile.get_language_support_overrides
	return []string(class(self).GetLanguageSupportOverrides().Strings())
}

/*
Adds override for [method Font.is_script_supported].
*/
func (self Instance) SetScriptSupportOverride(script string, supported bool) { //gd:FontFile.set_script_support_override
	class(self).SetScriptSupportOverride(String.New(script), supported)
}

/*
Returns [code]true[/code] if support override is enabled for the [param script].
*/
func (self Instance) GetScriptSupportOverride(script string) bool { //gd:FontFile.get_script_support_override
	return bool(class(self).GetScriptSupportOverride(String.New(script)))
}

/*
Removes script support override.
*/
func (self Instance) RemoveScriptSupportOverride(script string) { //gd:FontFile.remove_script_support_override
	class(self).RemoveScriptSupportOverride(String.New(script))
}

/*
Returns list of script support overrides.
*/
func (self Instance) GetScriptSupportOverrides() []string { //gd:FontFile.get_script_support_overrides
	return []string(class(self).GetScriptSupportOverrides().Strings())
}

/*
Returns the glyph index of a [param char], optionally modified by the [param variation_selector].
*/
func (self Instance) GetGlyphIndex(size int, char int, variation_selector int) int { //gd:FontFile.get_glyph_index
	return int(int(class(self).GetGlyphIndex(int64(size), int64(char), int64(variation_selector))))
}

/*
Returns character code associated with [param glyph_index], or [code]0[/code] if [param glyph_index] is invalid. See [method get_glyph_index].
*/
func (self Instance) GetCharFromGlyphIndex(size int, glyph_index int) int { //gd:FontFile.get_char_from_glyph_index
	return int(int(class(self).GetCharFromGlyphIndex(int64(size), int64(glyph_index))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.FontFile

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("FontFile"))
	casted := Instance{*(*gdclass.FontFile)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Data() []byte {
	return []byte(class(self).GetData().Bytes())
}

func (self Instance) SetData(value []byte) {
	class(self).SetData(Packed.Bytes(Packed.New(value...)))
}

func (self Instance) GenerateMipmaps() bool {
	return bool(class(self).GetGenerateMipmaps())
}

func (self Instance) SetGenerateMipmaps(value bool) {
	class(self).SetGenerateMipmaps(value)
}

func (self Instance) DisableEmbeddedBitmaps() bool {
	return bool(class(self).GetDisableEmbeddedBitmaps())
}

func (self Instance) SetDisableEmbeddedBitmaps(value bool) {
	class(self).SetDisableEmbeddedBitmaps(value)
}

func (self Instance) Antialiasing() gdclass.TextServerFontAntialiasing {
	return gdclass.TextServerFontAntialiasing(class(self).GetAntialiasing())
}

func (self Instance) SetAntialiasing(value gdclass.TextServerFontAntialiasing) {
	class(self).SetAntialiasing(value)
}

func (self Instance) SetFontName(value string) {
	class(self).SetFontName(String.New(value))
}

func (self Instance) SetStyleName(value string) {
	class(self).SetFontStyleName(String.New(value))
}

func (self Instance) SetFontStyle(value gdclass.TextServerFontStyle) {
	class(self).SetFontStyle(value)
}

func (self Instance) SetFontWeight(value int) {
	class(self).SetFontWeight(int64(value))
}

func (self Instance) SetFontStretch(value int) {
	class(self).SetFontStretch(int64(value))
}

func (self Instance) SubpixelPositioning() gdclass.TextServerSubpixelPositioning {
	return gdclass.TextServerSubpixelPositioning(class(self).GetSubpixelPositioning())
}

func (self Instance) SetSubpixelPositioning(value gdclass.TextServerSubpixelPositioning) {
	class(self).SetSubpixelPositioning(value)
}

func (self Instance) MultichannelSignedDistanceField() bool {
	return bool(class(self).IsMultichannelSignedDistanceField())
}

func (self Instance) SetMultichannelSignedDistanceField(value bool) {
	class(self).SetMultichannelSignedDistanceField(value)
}

func (self Instance) MsdfPixelRange() int {
	return int(int(class(self).GetMsdfPixelRange()))
}

func (self Instance) SetMsdfPixelRange(value int) {
	class(self).SetMsdfPixelRange(int64(value))
}

func (self Instance) MsdfSize() int {
	return int(int(class(self).GetMsdfSize()))
}

func (self Instance) SetMsdfSize(value int) {
	class(self).SetMsdfSize(int64(value))
}

func (self Instance) AllowSystemFallback() bool {
	return bool(class(self).IsAllowSystemFallback())
}

func (self Instance) SetAllowSystemFallback(value bool) {
	class(self).SetAllowSystemFallback(value)
}

func (self Instance) ForceAutohinter() bool {
	return bool(class(self).IsForceAutohinter())
}

func (self Instance) SetForceAutohinter(value bool) {
	class(self).SetForceAutohinter(value)
}

func (self Instance) Hinting() gdclass.TextServerHinting {
	return gdclass.TextServerHinting(class(self).GetHinting())
}

func (self Instance) SetHinting(value gdclass.TextServerHinting) {
	class(self).SetHinting(value)
}

func (self Instance) Oversampling() Float.X {
	return Float.X(Float.X(class(self).GetOversampling()))
}

func (self Instance) SetOversampling(value Float.X) {
	class(self).SetOversampling(float64(value))
}

func (self Instance) FixedSize() int {
	return int(int(class(self).GetFixedSize()))
}

func (self Instance) SetFixedSize(value int) {
	class(self).SetFixedSize(int64(value))
}

func (self Instance) FixedSizeScaleMode() gdclass.TextServerFixedSizeScaleMode {
	return gdclass.TextServerFixedSizeScaleMode(class(self).GetFixedSizeScaleMode())
}

func (self Instance) SetFixedSizeScaleMode(value gdclass.TextServerFixedSizeScaleMode) {
	class(self).SetFixedSizeScaleMode(value)
}

func (self Instance) OpentypeFeatureOverrides() map[any]any {
	return map[any]any(gd.DictionaryAs[map[any]any](class(self).GetOpentypeFeatureOverrides()))
}

func (self Instance) SetOpentypeFeatureOverrides(value map[any]any) {
	class(self).SetOpentypeFeatureOverrides(gd.DictionaryFromMap(value))
}

/*
Loads an AngelCode BMFont (.fnt, .font) bitmap font from file [param path].
[b]Warning:[/b] This method should only be used in the editor or in cases when you need to load external fonts at run-time, such as fonts located at the [code]user://[/code] directory.
*/
//go:nosplit
func (self class) LoadBitmapFont(path String.Readable) Error.Code { //gd:FontFile.load_bitmap_font
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_load_bitmap_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Loads a TrueType (.ttf), OpenType (.otf), WOFF (.woff), WOFF2 (.woff2) or Type 1 (.pfb, .pfm) dynamic font from file [param path].
[b]Warning:[/b] This method should only be used in the editor or in cases when you need to load external fonts at run-time, such as fonts located at the [code]user://[/code] directory.
*/
//go:nosplit
func (self class) LoadDynamicFont(path String.Readable) Error.Code { //gd:FontFile.load_dynamic_font
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(path)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_load_dynamic_font, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetData(data Packed.Bytes) { //gd:FontFile.set_data
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedByteArray, byte](Packed.Array[byte](data))))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetData() Packed.Bytes { //gd:FontFile.get_data
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_data, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Bytes(Array.Through(gd.PackedProxy[gd.PackedByteArray, byte]{}, pointers.Pack(pointers.New[gd.PackedByteArray](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFontName(name String.Readable) { //gd:FontFile.set_font_name
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_font_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetFontStyleName(name String.Readable) { //gd:FontFile.set_font_style_name
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(name)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_font_style_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetFontStyle(style gdclass.TextServerFontStyle) { //gd:FontFile.set_font_style
	var frame = callframe.New()
	callframe.Arg(frame, style)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_font_style, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetFontWeight(weight int64) { //gd:FontFile.set_font_weight
	var frame = callframe.New()
	callframe.Arg(frame, weight)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_font_weight, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetFontStretch(stretch int64) { //gd:FontFile.set_font_stretch
	var frame = callframe.New()
	callframe.Arg(frame, stretch)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_font_stretch, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetAntialiasing(antialiasing gdclass.TextServerFontAntialiasing) { //gd:FontFile.set_antialiasing
	var frame = callframe.New()
	callframe.Arg(frame, antialiasing)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_antialiasing, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAntialiasing() gdclass.TextServerFontAntialiasing { //gd:FontFile.get_antialiasing
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerFontAntialiasing](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_antialiasing, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDisableEmbeddedBitmaps(disable_embedded_bitmaps bool) { //gd:FontFile.set_disable_embedded_bitmaps
	var frame = callframe.New()
	callframe.Arg(frame, disable_embedded_bitmaps)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_disable_embedded_bitmaps, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDisableEmbeddedBitmaps() bool { //gd:FontFile.get_disable_embedded_bitmaps
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_disable_embedded_bitmaps, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetGenerateMipmaps(generate_mipmaps bool) { //gd:FontFile.set_generate_mipmaps
	var frame = callframe.New()
	callframe.Arg(frame, generate_mipmaps)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_generate_mipmaps, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetGenerateMipmaps() bool { //gd:FontFile.get_generate_mipmaps
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_generate_mipmaps, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMultichannelSignedDistanceField(msdf bool) { //gd:FontFile.set_multichannel_signed_distance_field
	var frame = callframe.New()
	callframe.Arg(frame, msdf)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_multichannel_signed_distance_field, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsMultichannelSignedDistanceField() bool { //gd:FontFile.is_multichannel_signed_distance_field
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_is_multichannel_signed_distance_field, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMsdfPixelRange(msdf_pixel_range int64) { //gd:FontFile.set_msdf_pixel_range
	var frame = callframe.New()
	callframe.Arg(frame, msdf_pixel_range)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_msdf_pixel_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMsdfPixelRange() int64 { //gd:FontFile.get_msdf_pixel_range
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_msdf_pixel_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetMsdfSize(msdf_size int64) { //gd:FontFile.set_msdf_size
	var frame = callframe.New()
	callframe.Arg(frame, msdf_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_msdf_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetMsdfSize() int64 { //gd:FontFile.get_msdf_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_msdf_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFixedSize(fixed_size int64) { //gd:FontFile.set_fixed_size
	var frame = callframe.New()
	callframe.Arg(frame, fixed_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_fixed_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFixedSize() int64 { //gd:FontFile.get_fixed_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_fixed_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetFixedSizeScaleMode(fixed_size_scale_mode gdclass.TextServerFixedSizeScaleMode) { //gd:FontFile.set_fixed_size_scale_mode
	var frame = callframe.New()
	callframe.Arg(frame, fixed_size_scale_mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_fixed_size_scale_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFixedSizeScaleMode() gdclass.TextServerFixedSizeScaleMode { //gd:FontFile.get_fixed_size_scale_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerFixedSizeScaleMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_fixed_size_scale_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAllowSystemFallback(allow_system_fallback bool) { //gd:FontFile.set_allow_system_fallback
	var frame = callframe.New()
	callframe.Arg(frame, allow_system_fallback)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_allow_system_fallback, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsAllowSystemFallback() bool { //gd:FontFile.is_allow_system_fallback
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_is_allow_system_fallback, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetForceAutohinter(force_autohinter bool) { //gd:FontFile.set_force_autohinter
	var frame = callframe.New()
	callframe.Arg(frame, force_autohinter)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_force_autohinter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsForceAutohinter() bool { //gd:FontFile.is_force_autohinter
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_is_force_autohinter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetHinting(hinting gdclass.TextServerHinting) { //gd:FontFile.set_hinting
	var frame = callframe.New()
	callframe.Arg(frame, hinting)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_hinting, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetHinting() gdclass.TextServerHinting { //gd:FontFile.get_hinting
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerHinting](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_hinting, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSubpixelPositioning(subpixel_positioning gdclass.TextServerSubpixelPositioning) { //gd:FontFile.set_subpixel_positioning
	var frame = callframe.New()
	callframe.Arg(frame, subpixel_positioning)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_subpixel_positioning, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSubpixelPositioning() gdclass.TextServerSubpixelPositioning { //gd:FontFile.get_subpixel_positioning
	var frame = callframe.New()
	var r_ret = callframe.Ret[gdclass.TextServerSubpixelPositioning](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_subpixel_positioning, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOversampling(oversampling float64) { //gd:FontFile.set_oversampling
	var frame = callframe.New()
	callframe.Arg(frame, oversampling)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_oversampling, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOversampling() float64 { //gd:FontFile.get_oversampling
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_oversampling, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns number of the font cache entries.
*/
//go:nosplit
func (self class) GetCacheCount() int64 { //gd:FontFile.get_cache_count
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_cache_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes all font cache entries.
*/
//go:nosplit
func (self class) ClearCache() { //gd:FontFile.clear_cache
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_clear_cache, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes specified font cache entry.
*/
//go:nosplit
func (self class) RemoveCache(cache_index int64) { //gd:FontFile.remove_cache
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_remove_cache, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns list of the font sizes in the cache. Each size is [Vector2i] with font size and outline size.
*/
//go:nosplit
func (self class) GetSizeCacheList(cache_index int64) Array.Contains[Vector2i.XY] { //gd:FontFile.get_size_cache_list
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_size_cache_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Vector2i.XY]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Removes all font sizes from the cache entry
*/
//go:nosplit
func (self class) ClearSizeCache(cache_index int64) { //gd:FontFile.clear_size_cache
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_clear_size_cache, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes specified font size from the cache entry.
*/
//go:nosplit
func (self class) RemoveSizeCache(cache_index int64, size Vector2i.XY) { //gd:FontFile.remove_size_cache
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_remove_size_cache, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets variation coordinates for the specified font cache entry. See [method Font.get_supported_variation_list] for more info.
*/
//go:nosplit
func (self class) SetVariationCoordinates(cache_index int64, variation_coordinates Dictionary.Any) { //gd:FontFile.set_variation_coordinates
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, pointers.Get(gd.InternalDictionary(variation_coordinates)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_variation_coordinates, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns variation coordinates for the specified font cache entry. See [method Font.get_supported_variation_list] for more info.
*/
//go:nosplit
func (self class) GetVariationCoordinates(cache_index int64) Dictionary.Any { //gd:FontFile.get_variation_coordinates
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_variation_coordinates, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets embolden strength, if is not equal to zero, emboldens the font outlines. Negative values reduce the outline thickness.
*/
//go:nosplit
func (self class) SetEmbolden(cache_index int64, strength float64) { //gd:FontFile.set_embolden
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, strength)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_embolden, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns embolden strength, if is not equal to zero, emboldens the font outlines. Negative values reduce the outline thickness.
*/
//go:nosplit
func (self class) GetEmbolden(cache_index int64) float64 { //gd:FontFile.get_embolden
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_embolden, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets 2D transform, applied to the font outlines, can be used for slanting, flipping, and rotating glyphs.
*/
//go:nosplit
func (self class) SetTransform(cache_index int64, transform Transform2D.OriginXY) { //gd:FontFile.set_transform
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, transform)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns 2D transform, applied to the font outlines, can be used for slanting, flipping and rotating glyphs.
*/
//go:nosplit
func (self class) GetTransform(cache_index int64) Transform2D.OriginXY { //gd:FontFile.get_transform
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret = callframe.Ret[Transform2D.OriginXY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_transform, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the spacing for [param spacing] (see [enum TextServer.SpacingType]) to [param value] in pixels (not relative to the font size).
*/
//go:nosplit
func (self class) SetExtraSpacing(cache_index int64, spacing gdclass.TextServerSpacingType, value int64) { //gd:FontFile.set_extra_spacing
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, spacing)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_extra_spacing, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns spacing for [param spacing] (see [enum TextServer.SpacingType]) in pixels (not relative to the font size).
*/
//go:nosplit
func (self class) GetExtraSpacing(cache_index int64, spacing gdclass.TextServerSpacingType) int64 { //gd:FontFile.get_extra_spacing
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, spacing)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_extra_spacing, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets extra baseline offset (as a fraction of font height).
*/
//go:nosplit
func (self class) SetExtraBaselineOffset(cache_index int64, baseline_offset float64) { //gd:FontFile.set_extra_baseline_offset
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, baseline_offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_extra_baseline_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns extra baseline offset (as a fraction of font height).
*/
//go:nosplit
func (self class) GetExtraBaselineOffset(cache_index int64) float64 { //gd:FontFile.get_extra_baseline_offset
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_extra_baseline_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets an active face index in the TrueType / OpenType collection.
*/
//go:nosplit
func (self class) SetFaceIndex(cache_index int64, face_index int64) { //gd:FontFile.set_face_index
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, face_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_face_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an active face index in the TrueType / OpenType collection.
*/
//go:nosplit
func (self class) GetFaceIndex(cache_index int64) int64 { //gd:FontFile.get_face_index
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_face_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the font ascent (number of pixels above the baseline).
*/
//go:nosplit
func (self class) SetCacheAscent(cache_index int64, size int64, ascent float64) { //gd:FontFile.set_cache_ascent
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, ascent)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_cache_ascent, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the font ascent (number of pixels above the baseline).
*/
//go:nosplit
func (self class) GetCacheAscent(cache_index int64, size int64) float64 { //gd:FontFile.get_cache_ascent
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_cache_ascent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets the font descent (number of pixels below the baseline).
*/
//go:nosplit
func (self class) SetCacheDescent(cache_index int64, size int64, descent float64) { //gd:FontFile.set_cache_descent
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, descent)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_cache_descent, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the font descent (number of pixels below the baseline).
*/
//go:nosplit
func (self class) GetCacheDescent(cache_index int64, size int64) float64 { //gd:FontFile.get_cache_descent
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_cache_descent, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets pixel offset of the underline below the baseline.
*/
//go:nosplit
func (self class) SetCacheUnderlinePosition(cache_index int64, size int64, underline_position float64) { //gd:FontFile.set_cache_underline_position
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, underline_position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_cache_underline_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns pixel offset of the underline below the baseline.
*/
//go:nosplit
func (self class) GetCacheUnderlinePosition(cache_index int64, size int64) float64 { //gd:FontFile.get_cache_underline_position
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_cache_underline_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets thickness of the underline in pixels.
*/
//go:nosplit
func (self class) SetCacheUnderlineThickness(cache_index int64, size int64, underline_thickness float64) { //gd:FontFile.set_cache_underline_thickness
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, underline_thickness)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_cache_underline_thickness, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns thickness of the underline in pixels.
*/
//go:nosplit
func (self class) GetCacheUnderlineThickness(cache_index int64, size int64) float64 { //gd:FontFile.get_cache_underline_thickness
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_cache_underline_thickness, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets scaling factor of the color bitmap font.
*/
//go:nosplit
func (self class) SetCacheScale(cache_index int64, size int64, scale float64) { //gd:FontFile.set_cache_scale
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_cache_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns scaling factor of the color bitmap font.
*/
//go:nosplit
func (self class) GetCacheScale(cache_index int64, size int64) float64 { //gd:FontFile.get_cache_scale
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_cache_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns number of textures used by font cache entry.
*/
//go:nosplit
func (self class) GetTextureCount(cache_index int64, size Vector2i.XY) int64 { //gd:FontFile.get_texture_count
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_texture_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes all textures from font cache entry.
[b]Note:[/b] This function will not remove glyphs associated with the texture, use [method remove_glyph] to remove them manually.
*/
//go:nosplit
func (self class) ClearTextures(cache_index int64, size Vector2i.XY) { //gd:FontFile.clear_textures
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_clear_textures, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes specified texture from the cache entry.
[b]Note:[/b] This function will not remove glyphs associated with the texture. Remove them manually using [method remove_glyph].
*/
//go:nosplit
func (self class) RemoveTexture(cache_index int64, size Vector2i.XY, texture_index int64) { //gd:FontFile.remove_texture
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_remove_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets font cache texture image.
*/
//go:nosplit
func (self class) SetTextureImage(cache_index int64, size Vector2i.XY, texture_index int64, image [1]gdclass.Image) { //gd:FontFile.set_texture_image
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	callframe.Arg(frame, pointers.Get(image[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_texture_image, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a copy of the font cache texture image.
*/
//go:nosplit
func (self class) GetTextureImage(cache_index int64, size Vector2i.XY, texture_index int64) [1]gdclass.Image { //gd:FontFile.get_texture_image
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_texture_image, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Image{gd.PointerWithOwnershipTransferredToGo[gdclass.Image](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sets array containing glyph packing data.
*/
//go:nosplit
func (self class) SetTextureOffsets(cache_index int64, size Vector2i.XY, texture_index int64, offset Packed.Array[int32]) { //gd:FontFile.set_texture_offsets
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	callframe.Arg(frame, pointers.Get(gd.InternalPacked[gd.PackedInt32Array, int32](offset)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_texture_offsets, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns a copy of the array containing glyph packing data.
*/
//go:nosplit
func (self class) GetTextureOffsets(cache_index int64, size Vector2i.XY, texture_index int64) Packed.Array[int32] { //gd:FontFile.get_texture_offsets
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, texture_index)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_texture_offsets, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int32](Array.Through(gd.PackedProxy[gd.PackedInt32Array, int32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns list of rendered glyphs in the cache entry.
*/
//go:nosplit
func (self class) GetGlyphList(cache_index int64, size Vector2i.XY) Packed.Array[int32] { //gd:FontFile.get_glyph_list
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_glyph_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Array[int32](Array.Through(gd.PackedProxy[gd.PackedInt32Array, int32]{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Removes all rendered glyph information from the cache entry.
[b]Note:[/b] This function will not remove textures associated with the glyphs, use [method remove_texture] to remove them manually.
*/
//go:nosplit
func (self class) ClearGlyphs(cache_index int64, size Vector2i.XY) { //gd:FontFile.clear_glyphs
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_clear_glyphs, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes specified rendered glyph information from the cache entry.
[b]Note:[/b] This function will not remove textures associated with the glyphs, use [method remove_texture] to remove them manually.
*/
//go:nosplit
func (self class) RemoveGlyph(cache_index int64, size Vector2i.XY, glyph int64) { //gd:FontFile.remove_glyph
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_remove_glyph, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets glyph advance (offset of the next glyph).
[b]Note:[/b] Advance for glyphs outlines is the same as the base glyph advance and is not saved.
*/
//go:nosplit
func (self class) SetGlyphAdvance(cache_index int64, size int64, glyph int64, advance Vector2.XY) { //gd:FontFile.set_glyph_advance
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, advance)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_glyph_advance, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns glyph advance (offset of the next glyph).
[b]Note:[/b] Advance for glyphs outlines is the same as the base glyph advance and is not saved.
*/
//go:nosplit
func (self class) GetGlyphAdvance(cache_index int64, size int64, glyph int64) Vector2.XY { //gd:FontFile.get_glyph_advance
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_glyph_advance, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets glyph offset from the baseline.
*/
//go:nosplit
func (self class) SetGlyphOffset(cache_index int64, size Vector2i.XY, glyph int64, offset Vector2.XY) { //gd:FontFile.set_glyph_offset
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, offset)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_glyph_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns glyph offset from the baseline.
*/
//go:nosplit
func (self class) GetGlyphOffset(cache_index int64, size Vector2i.XY, glyph int64) Vector2.XY { //gd:FontFile.get_glyph_offset
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_glyph_offset, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets glyph size.
*/
//go:nosplit
func (self class) SetGlyphSize(cache_index int64, size Vector2i.XY, glyph int64, gl_size Vector2.XY) { //gd:FontFile.set_glyph_size
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, gl_size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_glyph_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns glyph size.
*/
//go:nosplit
func (self class) GetGlyphSize(cache_index int64, size Vector2i.XY, glyph int64) Vector2.XY { //gd:FontFile.get_glyph_size
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_glyph_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets rectangle in the cache texture containing the glyph.
*/
//go:nosplit
func (self class) SetGlyphUvRect(cache_index int64, size Vector2i.XY, glyph int64, uv_rect Rect2.PositionSize) { //gd:FontFile.set_glyph_uv_rect
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, uv_rect)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_glyph_uv_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns rectangle in the cache texture containing the glyph.
*/
//go:nosplit
func (self class) GetGlyphUvRect(cache_index int64, size Vector2i.XY, glyph int64) Rect2.PositionSize { //gd:FontFile.get_glyph_uv_rect
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[Rect2.PositionSize](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_glyph_uv_rect, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Sets index of the cache texture containing the glyph.
*/
//go:nosplit
func (self class) SetGlyphTextureIdx(cache_index int64, size Vector2i.XY, glyph int64, texture_idx int64) { //gd:FontFile.set_glyph_texture_idx
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	callframe.Arg(frame, texture_idx)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_glyph_texture_idx, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns index of the cache texture containing the glyph.
*/
//go:nosplit
func (self class) GetGlyphTextureIdx(cache_index int64, size Vector2i.XY, glyph int64) int64 { //gd:FontFile.get_glyph_texture_idx
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_glyph_texture_idx, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns list of the kerning overrides.
*/
//go:nosplit
func (self class) GetKerningList(cache_index int64, size int64) Array.Contains[Vector2i.XY] { //gd:FontFile.get_kerning_list
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_kerning_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Vector2i.XY]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Removes all kerning overrides.
*/
//go:nosplit
func (self class) ClearKerningMap(cache_index int64, size int64) { //gd:FontFile.clear_kerning_map
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_clear_kerning_map, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes kerning override for the pair of glyphs.
*/
//go:nosplit
func (self class) RemoveKerning(cache_index int64, size int64, glyph_pair Vector2i.XY) { //gd:FontFile.remove_kerning
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph_pair)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_remove_kerning, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets kerning for the pair of glyphs.
*/
//go:nosplit
func (self class) SetKerning(cache_index int64, size int64, glyph_pair Vector2i.XY, kerning Vector2.XY) { //gd:FontFile.set_kerning
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph_pair)
	callframe.Arg(frame, kerning)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_kerning, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns kerning for the pair of glyphs.
*/
//go:nosplit
func (self class) GetKerning(cache_index int64, size int64, glyph_pair Vector2i.XY) Vector2.XY { //gd:FontFile.get_kerning
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph_pair)
	var r_ret = callframe.Ret[Vector2.XY](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_kerning, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Renders the range of characters to the font cache texture.
*/
//go:nosplit
func (self class) RenderRange(cache_index int64, size Vector2i.XY, start int64, end int64) { //gd:FontFile.render_range
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, start)
	callframe.Arg(frame, end)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_render_range, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Renders specified glyph to the font cache texture.
*/
//go:nosplit
func (self class) RenderGlyph(cache_index int64, size Vector2i.XY, index int64) { //gd:FontFile.render_glyph
	var frame = callframe.New()
	callframe.Arg(frame, cache_index)
	callframe.Arg(frame, size)
	callframe.Arg(frame, index)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_render_glyph, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Adds override for [method Font.is_language_supported].
*/
//go:nosplit
func (self class) SetLanguageSupportOverride(language String.Readable, supported bool) { //gd:FontFile.set_language_support_override
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(language)))
	callframe.Arg(frame, supported)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_language_support_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if support override is enabled for the [param language].
*/
//go:nosplit
func (self class) GetLanguageSupportOverride(language String.Readable) bool { //gd:FontFile.get_language_support_override
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(language)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_language_support_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Remove language support override.
*/
//go:nosplit
func (self class) RemoveLanguageSupportOverride(language String.Readable) { //gd:FontFile.remove_language_support_override
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(language)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_remove_language_support_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns list of language support overrides.
*/
//go:nosplit
func (self class) GetLanguageSupportOverrides() Packed.Strings { //gd:FontFile.get_language_support_overrides
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_language_support_overrides, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Adds override for [method Font.is_script_supported].
*/
//go:nosplit
func (self class) SetScriptSupportOverride(script String.Readable, supported bool) { //gd:FontFile.set_script_support_override
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(script)))
	callframe.Arg(frame, supported)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_script_support_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if support override is enabled for the [param script].
*/
//go:nosplit
func (self class) GetScriptSupportOverride(script String.Readable) bool { //gd:FontFile.get_script_support_override
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(script)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_script_support_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Removes script support override.
*/
//go:nosplit
func (self class) RemoveScriptSupportOverride(script String.Readable) { //gd:FontFile.remove_script_support_override
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(script)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_remove_script_support_override, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns list of script support overrides.
*/
//go:nosplit
func (self class) GetScriptSupportOverrides() Packed.Strings { //gd:FontFile.get_script_support_overrides
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_script_support_overrides, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOpentypeFeatureOverrides(overrides Dictionary.Any) { //gd:FontFile.set_opentype_feature_overrides
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalDictionary(overrides)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_set_opentype_feature_overrides, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOpentypeFeatureOverrides() Dictionary.Any { //gd:FontFile.get_opentype_feature_overrides
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_opentype_feature_overrides, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the glyph index of a [param char], optionally modified by the [param variation_selector].
*/
//go:nosplit
func (self class) GetGlyphIndex(size int64, char int64, variation_selector int64) int64 { //gd:FontFile.get_glyph_index
	var frame = callframe.New()
	callframe.Arg(frame, size)
	callframe.Arg(frame, char)
	callframe.Arg(frame, variation_selector)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_glyph_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns character code associated with [param glyph_index], or [code]0[/code] if [param glyph_index] is invalid. See [method get_glyph_index].
*/
//go:nosplit
func (self class) GetCharFromGlyphIndex(size int64, glyph_index int64) int64 { //gd:FontFile.get_char_from_glyph_index
	var frame = callframe.New()
	callframe.Arg(frame, size)
	callframe.Arg(frame, glyph_index)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.FontFile.Bind_get_char_from_glyph_index, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsFontFile() Advanced     { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsFontFile() Instance  { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsFont() Font.Advanced    { return *((*Font.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsFont() Font.Instance { return *((*Font.Instance)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Font.Advanced(self.AsFont()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Font.Instance(self.AsFont()), name)
	}
}
func init() {
	gdclass.Register("FontFile", func(ptr gd.Object) any { return [1]gdclass.FontFile{*(*gdclass.FontFile)(unsafe.Pointer(&ptr))} })
}
