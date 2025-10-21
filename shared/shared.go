// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"

	"github.com/imagekit-developer/imagekit-go/v2/internal/apijson"
	"github.com/imagekit-developer/imagekit-go/v2/packages/param"
	"github.com/imagekit-developer/imagekit-go/v2/packages/respjson"
	"github.com/imagekit-developer/imagekit-go/v2/shared/constant"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type BaseOverlay struct {
	Position OverlayPosition `json:"position"`
	Timing   OverlayTiming   `json:"timing"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Position    respjson.Field
		Timing      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaseOverlay) RawJSON() string { return r.JSON.raw }
func (r *BaseOverlay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Extensions []ExtensionUnion

// ExtensionUnion contains all possible properties and values from
// [ExtensionRemoveBg], [ExtensionAutoTagging], [ExtensionAIAutoDescription].
//
// Use the [ExtensionUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ExtensionUnion struct {
	// Any of "remove-bg", nil, "ai-auto-description".
	Name string `json:"name"`
	// This field is from variant [ExtensionRemoveBg].
	Options ExtensionRemoveBgOptions `json:"options"`
	// This field is from variant [ExtensionAutoTagging].
	MaxTags int64 `json:"maxTags"`
	// This field is from variant [ExtensionAutoTagging].
	MinConfidence int64 `json:"minConfidence"`
	JSON          struct {
		Name          respjson.Field
		Options       respjson.Field
		MaxTags       respjson.Field
		MinConfidence respjson.Field
		raw           string
	} `json:"-"`
}

func (u ExtensionUnion) AsRemoveBg() (v ExtensionRemoveBg) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtensionUnion) AsAutoTagging() (v ExtensionAutoTagging) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtensionUnion) AsAIAutoDescription() (v ExtensionAIAutoDescription) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtensionUnion) RawJSON() string { return u.JSON.raw }

func (r *ExtensionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionRemoveBg struct {
	// Specifies the background removal extension.
	Name    constant.RemoveBg        `json:"name,required"`
	Options ExtensionRemoveBgOptions `json:"options"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Options     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtensionRemoveBg) RawJSON() string { return r.JSON.raw }
func (r *ExtensionRemoveBg) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionRemoveBgOptions struct {
	// Whether to add an artificial shadow to the result. Default is false. Note:
	// Adding shadows is currently only supported for car photos.
	AddShadow bool `json:"add_shadow"`
	// Specifies a solid color background using hex code (e.g., "81d4fa", "fff") or
	// color name (e.g., "green"). If this parameter is set, `bg_image_url` must be
	// empty.
	BgColor string `json:"bg_color"`
	// Sets a background image from a URL. If this parameter is set, `bg_color` must be
	// empty.
	BgImageURL string `json:"bg_image_url"`
	// Allows semi-transparent regions in the result. Default is true. Note:
	// Semitransparency is currently only supported for car windows.
	Semitransparency bool `json:"semitransparency"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddShadow        respjson.Field
		BgColor          respjson.Field
		BgImageURL       respjson.Field
		Semitransparency respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtensionRemoveBgOptions) RawJSON() string { return r.JSON.raw }
func (r *ExtensionRemoveBgOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionAutoTagging struct {
	// Maximum number of tags to attach to the asset.
	MaxTags int64 `json:"maxTags,required"`
	// Minimum confidence level for tags to be considered valid.
	MinConfidence int64 `json:"minConfidence,required"`
	// Specifies the auto-tagging extension used.
	//
	// Any of "google-auto-tagging", "aws-auto-tagging".
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxTags       respjson.Field
		MinConfidence respjson.Field
		Name          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtensionAutoTagging) RawJSON() string { return r.JSON.raw }
func (r *ExtensionAutoTagging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionAIAutoDescription struct {
	// Specifies the auto description extension.
	Name constant.AIAutoDescription `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtensionAIAutoDescription) RawJSON() string { return r.JSON.raw }
func (r *ExtensionAIAutoDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionsParam []ExtensionUnionParam

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtensionUnionParam struct {
	OfRemoveBg          *ExtensionRemoveBgParam          `json:",omitzero,inline"`
	OfAutoTagging       *ExtensionAutoTaggingParam       `json:",omitzero,inline"`
	OfAIAutoDescription *ExtensionAIAutoDescriptionParam `json:",omitzero,inline"`
	paramUnion
}

func (u ExtensionUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRemoveBg, u.OfAutoTagging, u.OfAIAutoDescription)
}
func (u *ExtensionUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtensionUnionParam) asAny() any {
	if !param.IsOmitted(u.OfRemoveBg) {
		return u.OfRemoveBg
	} else if !param.IsOmitted(u.OfAutoTagging) {
		return u.OfAutoTagging
	} else if !param.IsOmitted(u.OfAIAutoDescription) {
		return u.OfAIAutoDescription
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtensionUnionParam) GetOptions() *ExtensionRemoveBgOptionsParam {
	if vt := u.OfRemoveBg; vt != nil {
		return &vt.Options
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtensionUnionParam) GetMaxTags() *int64 {
	if vt := u.OfAutoTagging; vt != nil {
		return &vt.MaxTags
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtensionUnionParam) GetMinConfidence() *int64 {
	if vt := u.OfAutoTagging; vt != nil {
		return &vt.MinConfidence
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtensionUnionParam) GetName() *string {
	if vt := u.OfRemoveBg; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfAutoTagging; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfAIAutoDescription; vt != nil {
		return (*string)(&vt.Name)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ExtensionUnionParam](
		"name",
		apijson.Discriminator[ExtensionRemoveBgParam]("remove-bg"),
		apijson.Discriminator[ExtensionAutoTaggingParam]("google-auto-tagging"),
		apijson.Discriminator[ExtensionAutoTaggingParam]("aws-auto-tagging"),
		apijson.Discriminator[ExtensionAIAutoDescriptionParam]("ai-auto-description"),
	)
}

// The property Name is required.
type ExtensionRemoveBgParam struct {
	Options ExtensionRemoveBgOptionsParam `json:"options,omitzero"`
	// Specifies the background removal extension.
	//
	// This field can be elided, and will marshal its zero value as "remove-bg".
	Name constant.RemoveBg `json:"name,required"`
	paramObj
}

func (r ExtensionRemoveBgParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionRemoveBgParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionRemoveBgParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionRemoveBgOptionsParam struct {
	// Whether to add an artificial shadow to the result. Default is false. Note:
	// Adding shadows is currently only supported for car photos.
	AddShadow param.Opt[bool] `json:"add_shadow,omitzero"`
	// Specifies a solid color background using hex code (e.g., "81d4fa", "fff") or
	// color name (e.g., "green"). If this parameter is set, `bg_image_url` must be
	// empty.
	BgColor param.Opt[string] `json:"bg_color,omitzero"`
	// Sets a background image from a URL. If this parameter is set, `bg_color` must be
	// empty.
	BgImageURL param.Opt[string] `json:"bg_image_url,omitzero"`
	// Allows semi-transparent regions in the result. Default is true. Note:
	// Semitransparency is currently only supported for car windows.
	Semitransparency param.Opt[bool] `json:"semitransparency,omitzero"`
	paramObj
}

func (r ExtensionRemoveBgOptionsParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionRemoveBgOptionsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionRemoveBgOptionsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MaxTags, MinConfidence, Name are required.
type ExtensionAutoTaggingParam struct {
	// Maximum number of tags to attach to the asset.
	MaxTags int64 `json:"maxTags,required"`
	// Minimum confidence level for tags to be considered valid.
	MinConfidence int64 `json:"minConfidence,required"`
	// Specifies the auto-tagging extension used.
	//
	// Any of "google-auto-tagging", "aws-auto-tagging".
	Name string `json:"name,omitzero,required"`
	paramObj
}

func (r ExtensionAutoTaggingParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionAutoTaggingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionAutoTaggingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtensionAutoTaggingParam](
		"name", "google-auto-tagging", "aws-auto-tagging",
	)
}

func NewExtensionAIAutoDescriptionParam() ExtensionAIAutoDescriptionParam {
	return ExtensionAIAutoDescriptionParam{
		Name: "ai-auto-description",
	}
}

// This struct has a constant value, construct it with
// [NewExtensionAIAutoDescriptionParam].
type ExtensionAIAutoDescriptionParam struct {
	// Specifies the auto description extension.
	Name constant.AIAutoDescription `json:"name,required"`
	paramObj
}

func (r ExtensionAIAutoDescriptionParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionAIAutoDescriptionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionAIAutoDescriptionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Options for generating responsive image attributes including `src`, `srcSet`,
// and `sizes` for HTML `<img>` elements. This schema extends `SrcOptions` to add
// support for responsive image generation with breakpoints.
type GetImageAttributesOptions struct {
	// Custom list of **device-width breakpoints** in pixels. These define common
	// screen widths for responsive image generation.
	//
	// Defaults to `[640, 750, 828, 1080, 1200, 1920, 2048, 3840]`. Sorted
	// automatically.
	DeviceBreakpoints []float64 `json:"deviceBreakpoints"`
	// Custom list of **image-specific breakpoints** in pixels. Useful for generating
	// small variants (e.g., placeholders or thumbnails).
	//
	// Merged with `deviceBreakpoints` before calculating `srcSet`. Defaults to
	// `[16, 32, 48, 64, 96, 128, 256, 384]`. Sorted automatically.
	ImageBreakpoints []float64 `json:"imageBreakpoints"`
	// The value for the HTML `sizes` attribute (e.g., `"100vw"` or
	// `"(min-width:768px) 50vw, 100vw"`).
	//
	//   - If it includes one or more `vw` units, breakpoints smaller than the
	//     corresponding percentage of the smallest device width are excluded.
	//   - If it contains no `vw` units, the full breakpoint list is used.
	//
	// Enables a width-based strategy and generates `w` descriptors in `srcSet`.
	Sizes string `json:"sizes"`
	// The intended display width of the image in pixels, used **only when the `sizes`
	// attribute is not provided**.
	//
	// Triggers a DPR-based strategy (1x and 2x variants) and generates `x` descriptors
	// in `srcSet`.
	//
	// Ignored if `sizes` is present.
	Width float64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeviceBreakpoints respjson.Field
		ImageBreakpoints  respjson.Field
		Sizes             respjson.Field
		Width             respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
	SrcOptions
}

// Returns the unmodified JSON received from the API
func (r GetImageAttributesOptions) RawJSON() string { return r.JSON.raw }
func (r *GetImageAttributesOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ImageOverlay struct {
	// Specifies the relative path to the image used as an overlay.
	Input string         `json:"input,required"`
	Type  constant.Image `json:"type,required"`
	// The input path can be included in the layer as either `i-{input}` or
	// `ie-{base64_encoded_input}`. By default, the SDK determines the appropriate
	// format automatically. To always use base64 encoding (`ie-{base64}`), set this
	// parameter to `base64`. To always use plain text (`i-{input}`), set it to
	// `plain`.
	//
	// Any of "auto", "plain", "base64".
	Encoding string `json:"encoding"`
	// Array of transformations to be applied to the overlay image. Supported
	// transformations depends on the base/parent asset. See overlays on
	// [Images](https://imagekit.io/docs/add-overlays-on-images#list-of-supported-image-transformations-in-image-layers)
	// and
	// [Videos](https://imagekit.io/docs/add-overlays-on-videos#list-of-transformations-supported-on-image-overlay).
	Transformation []Transformation `json:"transformation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Input          respjson.Field
		Type           respjson.Field
		Encoding       respjson.Field
		Transformation respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
	BaseOverlay
}

// Returns the unmodified JSON received from the API
func (r ImageOverlay) RawJSON() string { return r.JSON.raw }
func (r *ImageOverlay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (ImageOverlay) ImplOverlayUnion() {}

// OverlayUnion contains all possible properties and values from [TextOverlay],
// [ImageOverlay], [VideoOverlay], [SubtitleOverlay], [SolidColorOverlay].
//
// Use the [OverlayUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type OverlayUnion struct {
	// This field is from variant [TextOverlay], [ImageOverlay], [VideoOverlay],
	// [SubtitleOverlay], [SolidColorOverlay].
	Position OverlayPosition `json:"position"`
	// This field is from variant [TextOverlay], [ImageOverlay], [VideoOverlay],
	// [SubtitleOverlay], [SolidColorOverlay].
	Timing OverlayTiming `json:"timing"`
	// This field is from variant [TextOverlay].
	Text string `json:"text"`
	// Any of "text", "image", "video", "subtitle", "solidColor".
	Type     string `json:"type"`
	Encoding string `json:"encoding"`
	// This field is a union of [[]TextOverlayTransformation], [[]Transformation],
	// [[]Transformation], [[]SubtitleOverlayTransformation],
	// [[]SolidColorOverlayTransformation]
	Transformation OverlayUnionTransformation `json:"transformation"`
	Input          string                     `json:"input"`
	// This field is from variant [SolidColorOverlay].
	Color string `json:"color"`
	JSON  struct {
		Position       respjson.Field
		Timing         respjson.Field
		Text           respjson.Field
		Type           respjson.Field
		Encoding       respjson.Field
		Transformation respjson.Field
		Input          respjson.Field
		Color          respjson.Field
		raw            string
	} `json:"-"`
}

// anyOverlay is implemented by each variant of [OverlayUnion] to add type safety
// for the return type of [OverlayUnion.AsAny]
type anyOverlay interface {
	ImplOverlayUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := OverlayUnion.AsAny().(type) {
//	case shared.TextOverlay:
//	case shared.ImageOverlay:
//	case shared.VideoOverlay:
//	case shared.SubtitleOverlay:
//	case shared.SolidColorOverlay:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u OverlayUnion) AsAny() anyOverlay {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image":
		return u.AsImage()
	case "video":
		return u.AsVideo()
	case "subtitle":
		return u.AsSubtitle()
	case "solidColor":
		return u.AsSolidColor()
	}
	return nil
}

func (u OverlayUnion) AsText() (v TextOverlay) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OverlayUnion) AsImage() (v ImageOverlay) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OverlayUnion) AsVideo() (v VideoOverlay) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OverlayUnion) AsSubtitle() (v SubtitleOverlay) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OverlayUnion) AsSolidColor() (v SolidColorOverlay) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OverlayUnion) RawJSON() string { return u.JSON.raw }

func (r *OverlayUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OverlayUnionTransformation is an implicit subunion of [OverlayUnion].
// OverlayUnionTransformation provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [OverlayUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfTextOverlayTransformationArray OfTransformationArray
// OfSubtitleOverlayTransformationArray OfSolidColorOverlayTransformationArray]
type OverlayUnionTransformation struct {
	// This field will be present if the value is a [[]TextOverlayTransformation]
	// instead of an object.
	OfTextOverlayTransformationArray []TextOverlayTransformation `json:",inline"`
	// This field will be present if the value is a [[]Transformation] instead of an
	// object.
	OfTransformationArray []Transformation `json:",inline"`
	// This field will be present if the value is a [[]SubtitleOverlayTransformation]
	// instead of an object.
	OfSubtitleOverlayTransformationArray []SubtitleOverlayTransformation `json:",inline"`
	// This field will be present if the value is a [[]SolidColorOverlayTransformation]
	// instead of an object.
	OfSolidColorOverlayTransformationArray []SolidColorOverlayTransformation `json:",inline"`
	JSON                                   struct {
		OfTextOverlayTransformationArray       respjson.Field
		OfTransformationArray                  respjson.Field
		OfSubtitleOverlayTransformationArray   respjson.Field
		OfSolidColorOverlayTransformationArray respjson.Field
		raw                                    string
	} `json:"-"`
}

func (r *OverlayUnionTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OverlayPosition struct {
	// Specifies the position of the overlay relative to the parent image or video.
	// Maps to `lfo` in the URL.
	//
	// Any of "center", "top", "left", "bottom", "right", "top_left", "top_right",
	// "bottom_left", "bottom_right".
	Focus OverlayPositionFocus `json:"focus"`
	// Specifies the x-coordinate of the top-left corner of the base asset where the
	// overlay's top-left corner will be positioned. It also accepts arithmetic
	// expressions such as `bw_mul_0.4` or `bw_sub_cw`. Maps to `lx` in the URL. Learn
	// about
	// [Arithmetic expressions](https://imagekit.io/docs/arithmetic-expressions-in-transformations).
	X OverlayPositionXUnion `json:"x"`
	// Specifies the y-coordinate of the top-left corner of the base asset where the
	// overlay's top-left corner will be positioned. It also accepts arithmetic
	// expressions such as `bh_mul_0.4` or `bh_sub_ch`. Maps to `ly` in the URL. Learn
	// about
	// [Arithmetic expressions](https://imagekit.io/docs/arithmetic-expressions-in-transformations).
	Y OverlayPositionYUnion `json:"y"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Focus       respjson.Field
		X           respjson.Field
		Y           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OverlayPosition) RawJSON() string { return r.JSON.raw }
func (r *OverlayPosition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the position of the overlay relative to the parent image or video.
// Maps to `lfo` in the URL.
type OverlayPositionFocus string

const (
	OverlayPositionFocusCenter      OverlayPositionFocus = "center"
	OverlayPositionFocusTop         OverlayPositionFocus = "top"
	OverlayPositionFocusLeft        OverlayPositionFocus = "left"
	OverlayPositionFocusBottom      OverlayPositionFocus = "bottom"
	OverlayPositionFocusRight       OverlayPositionFocus = "right"
	OverlayPositionFocusTopLeft     OverlayPositionFocus = "top_left"
	OverlayPositionFocusTopRight    OverlayPositionFocus = "top_right"
	OverlayPositionFocusBottomLeft  OverlayPositionFocus = "bottom_left"
	OverlayPositionFocusBottomRight OverlayPositionFocus = "bottom_right"
)

// OverlayPositionXUnion contains all possible properties and values from
// [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type OverlayPositionXUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u OverlayPositionXUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OverlayPositionXUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OverlayPositionXUnion) RawJSON() string { return u.JSON.raw }

func (r *OverlayPositionXUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OverlayPositionYUnion contains all possible properties and values from
// [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type OverlayPositionYUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u OverlayPositionYUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OverlayPositionYUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OverlayPositionYUnion) RawJSON() string { return u.JSON.raw }

func (r *OverlayPositionYUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OverlayTiming struct {
	// Specifies the duration (in seconds) during which the overlay should appear on
	// the base video. Accepts a positive number up to two decimal places (e.g., `20`
	// or `20.50`) and arithmetic expressions such as `bdu_mul_0.4` or `bdu_sub_idu`.
	// Applies only if the base asset is a video. Maps to `ldu` in the URL.
	Duration OverlayTimingDurationUnion `json:"duration"`
	// Specifies the end time (in seconds) for when the overlay should disappear from
	// the base video. If both end and duration are provided, duration is ignored.
	// Accepts a positive number up to two decimal places (e.g., `20` or `20.50`) and
	// arithmetic expressions such as `bdu_mul_0.4` or `bdu_sub_idu`. Applies only if
	// the base asset is a video. Maps to `leo` in the URL.
	End OverlayTimingEndUnion `json:"end"`
	// Specifies the start time (in seconds) for when the overlay should appear on the
	// base video. Accepts a positive number up to two decimal places (e.g., `20` or
	// `20.50`) and arithmetic expressions such as `bdu_mul_0.4` or `bdu_sub_idu`.
	// Applies only if the base asset is a video. Maps to `lso` in the URL.
	Start OverlayTimingStartUnion `json:"start"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duration    respjson.Field
		End         respjson.Field
		Start       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OverlayTiming) RawJSON() string { return r.JSON.raw }
func (r *OverlayTiming) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OverlayTimingDurationUnion contains all possible properties and values from
// [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type OverlayTimingDurationUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u OverlayTimingDurationUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OverlayTimingDurationUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OverlayTimingDurationUnion) RawJSON() string { return u.JSON.raw }

func (r *OverlayTimingDurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OverlayTimingEndUnion contains all possible properties and values from
// [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type OverlayTimingEndUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u OverlayTimingEndUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OverlayTimingEndUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OverlayTimingEndUnion) RawJSON() string { return u.JSON.raw }

func (r *OverlayTimingEndUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OverlayTimingStartUnion contains all possible properties and values from
// [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type OverlayTimingStartUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u OverlayTimingStartUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OverlayTimingStartUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OverlayTimingStartUnion) RawJSON() string { return u.JSON.raw }

func (r *OverlayTimingStartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resulting set of attributes suitable for an HTML `<img>` element. Useful for
// enabling responsive image loading with `srcSet` and `sizes`.
type ResponsiveImageAttributes struct {
	// URL for the _largest_ candidate (assigned to plain `src`).
	Src string `json:"src,required" format:"uri"`
	// `sizes` returned (or synthesised as `100vw`). The value for the HTML `sizes`
	// attribute.
	Sizes string `json:"sizes"`
	// Candidate set with `w` or `x` descriptors. Multiple image URLs separated by
	// commas, each with a descriptor.
	SrcSet string `json:"srcSet"`
	// Width as a number (if `width` was provided in the input options).
	Width float64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Src         respjson.Field
		Sizes       respjson.Field
		SrcSet      respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponsiveImageAttributes) RawJSON() string { return r.JSON.raw }
func (r *ResponsiveImageAttributes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SolidColorOverlay struct {
	// Specifies the color of the block using an RGB hex code (e.g., `FF0000`), an RGBA
	// code (e.g., `FFAABB50`), or a color name (e.g., `red`). If an 8-character value
	// is provided, the last two characters represent the opacity level (from `00` for
	// 0.00 to `99` for 0.99).
	Color string              `json:"color,required"`
	Type  constant.SolidColor `json:"type,required"`
	// Control width and height of the solid color overlay. Supported transformations
	// depend on the base/parent asset. See overlays on
	// [Images](https://imagekit.io/docs/add-overlays-on-images#apply-transformation-on-solid-color-overlay)
	// and
	// [Videos](https://imagekit.io/docs/add-overlays-on-videos#apply-transformations-on-solid-color-block-overlay).
	Transformation []SolidColorOverlayTransformation `json:"transformation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Color          respjson.Field
		Type           respjson.Field
		Transformation respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
	BaseOverlay
}

// Returns the unmodified JSON received from the API
func (r SolidColorOverlay) RawJSON() string { return r.JSON.raw }
func (r *SolidColorOverlay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (SolidColorOverlay) ImplOverlayUnion() {}

type SolidColorOverlayTransformation struct {
	// Specifies the transparency level of the solid color overlay. Accepts integers
	// from `1` to `9`.
	Alpha float64 `json:"alpha"`
	// Specifies the background color of the solid color overlay. Accepts an RGB hex
	// code (e.g., `FF0000`), an RGBA code (e.g., `FFAABB50`), or a color name.
	Background string `json:"background"`
	// Creates a linear gradient with two colors. Pass `true` for a default gradient,
	// or provide a string for a custom gradient. Only works if the base asset is an
	// image. See
	// [gradient](https://imagekit.io/docs/effects-and-enhancements#gradient---e-gradient).
	Gradient SolidColorOverlayTransformationGradientUnion `json:"gradient"`
	// Controls the height of the solid color overlay. Accepts a numeric value or an
	// arithmetic expression. Learn about
	// [arithmetic expressions](https://imagekit.io/docs/arithmetic-expressions-in-transformations).
	Height SolidColorOverlayTransformationHeightUnion `json:"height"`
	// Specifies the corner radius of the solid color overlay. Set to `max` for
	// circular or oval shape. See
	// [radius](https://imagekit.io/docs/effects-and-enhancements#radius---r).
	Radius SolidColorOverlayTransformationRadiusUnion `json:"radius"`
	// Controls the width of the solid color overlay. Accepts a numeric value or an
	// arithmetic expression (e.g., `bw_mul_0.2` or `bh_div_2`). Learn about
	// [arithmetic expressions](https://imagekit.io/docs/arithmetic-expressions-in-transformations).
	Width SolidColorOverlayTransformationWidthUnion `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Alpha       respjson.Field
		Background  respjson.Field
		Gradient    respjson.Field
		Height      respjson.Field
		Radius      respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SolidColorOverlayTransformation) RawJSON() string { return r.JSON.raw }
func (r *SolidColorOverlayTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SolidColorOverlayTransformationGradientUnion contains all possible properties
// and values from [bool], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfSolidColorOverlayTransformationGradientBoolean OfString]
type SolidColorOverlayTransformationGradientUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfSolidColorOverlayTransformationGradientBoolean bool `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfSolidColorOverlayTransformationGradientBoolean respjson.Field
		OfString                                         respjson.Field
		raw                                              string
	} `json:"-"`
}

func (u SolidColorOverlayTransformationGradientUnion) AsSolidColorOverlayTransformationGradientBoolean() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SolidColorOverlayTransformationGradientUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SolidColorOverlayTransformationGradientUnion) RawJSON() string { return u.JSON.raw }

func (r *SolidColorOverlayTransformationGradientUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SolidColorOverlayTransformationGradientBoolean bool

const (
	SolidColorOverlayTransformationGradientBooleanTrue SolidColorOverlayTransformationGradientBoolean = true
)

// SolidColorOverlayTransformationHeightUnion contains all possible properties and
// values from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type SolidColorOverlayTransformationHeightUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u SolidColorOverlayTransformationHeightUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SolidColorOverlayTransformationHeightUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SolidColorOverlayTransformationHeightUnion) RawJSON() string { return u.JSON.raw }

func (r *SolidColorOverlayTransformationHeightUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SolidColorOverlayTransformationRadiusUnion contains all possible properties and
// values from [float64], [constant.Max].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfMax]
type SolidColorOverlayTransformationRadiusUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [constant.Max] instead of an
	// object.
	OfMax constant.Max `json:",inline"`
	JSON  struct {
		OfFloat respjson.Field
		OfMax   respjson.Field
		raw     string
	} `json:"-"`
}

func (u SolidColorOverlayTransformationRadiusUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SolidColorOverlayTransformationRadiusUnion) AsMax() (v constant.Max) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SolidColorOverlayTransformationRadiusUnion) RawJSON() string { return u.JSON.raw }

func (r *SolidColorOverlayTransformationRadiusUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SolidColorOverlayTransformationWidthUnion contains all possible properties and
// values from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type SolidColorOverlayTransformationWidthUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u SolidColorOverlayTransformationWidthUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SolidColorOverlayTransformationWidthUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SolidColorOverlayTransformationWidthUnion) RawJSON() string { return u.JSON.raw }

func (r *SolidColorOverlayTransformationWidthUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Options for generating ImageKit URLs with transformations. See the
// [Transformations guide](https://imagekit.io/docs/transformations).
type SrcOptions struct {
	// Accepts a relative or absolute path of the resource. If a relative path is
	// provided, it is appended to the `urlEndpoint`. If an absolute path is provided,
	// `urlEndpoint` is ignored.
	Src string `json:"src,required"`
	// Get your urlEndpoint from the
	// [ImageKit dashboard](https://imagekit.io/dashboard/url-endpoints).
	URLEndpoint string `json:"urlEndpoint,required" format:"uri"`
	// When you want the signed URL to expire, specified in seconds. If `expiresIn` is
	// anything above 0, the URL will always be signed even if `signed` is set to
	// false. If not specified and `signed` is `true`, the signed URL will not expire
	// (valid indefinitely).
	//
	// Example: Setting `expiresIn: 3600` will make the URL expire 1 hour from
	// generation time. After the expiry time, the signed URL will no longer be valid
	// and ImageKit will return a 401 Unauthorized status code.
	//
	// [Learn more](https://imagekit.io/docs/media-delivery-basic-security#how-to-generate-signed-urls).
	ExpiresIn float64 `json:"expiresIn"`
	// These are additional query parameters that you want to add to the final URL.
	// They can be any query parameters and not necessarily related to ImageKit. This
	// is especially useful if you want to add a versioning parameter to your URLs.
	QueryParameters map[string]string `json:"queryParameters"`
	// Whether to sign the URL or not. Set this to `true` if you want to generate a
	// signed URL. If `signed` is `true` and `expiresIn` is not specified, the signed
	// URL will not expire (valid indefinitely). Note: If `expiresIn` is set to any
	// value above 0, the URL will always be signed regardless of this setting.
	// [Learn more](https://imagekit.io/docs/media-delivery-basic-security#how-to-generate-signed-urls).
	Signed bool `json:"signed"`
	// An array of objects specifying the transformations to be applied in the URL. If
	// more than one transformation is specified, they are applied in the order they
	// are specified as chained transformations. See
	// [Chained transformations](https://imagekit.io/docs/transformations#chained-transformations).
	Transformation []Transformation `json:"transformation"`
	// By default, the transformation string is added as a query parameter in the URL,
	// e.g., `?tr=w-100,h-100`. If you want to add the transformation string in the
	// path of the URL, set this to `path`. Learn more in the
	// [Transformations guide](https://imagekit.io/docs/transformations).
	//
	// Any of "path", "query".
	TransformationPosition TransformationPosition `json:"transformationPosition"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Src                    respjson.Field
		URLEndpoint            respjson.Field
		ExpiresIn              respjson.Field
		QueryParameters        respjson.Field
		Signed                 respjson.Field
		Transformation         respjson.Field
		TransformationPosition respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SrcOptions) RawJSON() string { return r.JSON.raw }
func (r *SrcOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Available streaming resolutions for
// [adaptive bitrate streaming](https://imagekit.io/docs/adaptive-bitrate-streaming)
type StreamingResolution string

const (
	StreamingResolution240  StreamingResolution = "240"
	StreamingResolution360  StreamingResolution = "360"
	StreamingResolution480  StreamingResolution = "480"
	StreamingResolution720  StreamingResolution = "720"
	StreamingResolution1080 StreamingResolution = "1080"
	StreamingResolution1440 StreamingResolution = "1440"
	StreamingResolution2160 StreamingResolution = "2160"
)

type SubtitleOverlay struct {
	// Specifies the relative path to the subtitle file used as an overlay.
	Input string            `json:"input,required"`
	Type  constant.Subtitle `json:"type,required"`
	// The input path can be included in the layer as either `i-{input}` or
	// `ie-{base64_encoded_input}`. By default, the SDK determines the appropriate
	// format automatically. To always use base64 encoding (`ie-{base64}`), set this
	// parameter to `base64`. To always use plain text (`i-{input}`), set it to
	// `plain`.
	//
	// Any of "auto", "plain", "base64".
	Encoding string `json:"encoding"`
	// Control styling of the subtitle. See
	// [Styling subtitles](https://imagekit.io/docs/add-overlays-on-videos#styling-controls-for-subtitles-layer).
	Transformation []SubtitleOverlayTransformation `json:"transformation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Input          respjson.Field
		Type           respjson.Field
		Encoding       respjson.Field
		Transformation respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
	BaseOverlay
}

// Returns the unmodified JSON received from the API
func (r SubtitleOverlay) RawJSON() string { return r.JSON.raw }
func (r *SubtitleOverlay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (SubtitleOverlay) ImplOverlayUnion() {}

// Subtitle styling options.
// [Learn more](https://imagekit.io/docs/add-overlays-on-videos#styling-controls-for-subtitles-layer)
// from the docs.
type SubtitleOverlayTransformation struct {
	// Specifies the subtitle background color using a standard color name, an RGB
	// color code (e.g., FF0000), or an RGBA color code (e.g., FFAABB50).
	//
	// [Subtitle styling options](https://imagekit.io/docs/add-overlays-on-videos#styling-controls-for-subtitles-layer)
	Background string `json:"background"`
	// Sets the font color of the subtitle text using a standard color name, an RGB
	// color code (e.g., FF0000), or an RGBA color code (e.g., FFAABB50).
	//
	// [Subtitle styling options](https://imagekit.io/docs/add-overlays-on-videos#styling-controls-for-subtitles-layer)
	Color string `json:"color"`
	// Font family for subtitles. Refer to the
	// [supported fonts](https://imagekit.io/docs/add-overlays-on-images#supported-text-font-list).
	FontFamily string `json:"fontFamily"`
	// Sets the font outline of the subtitle text. Requires the outline width (an
	// integer) and the outline color (as an RGB color code, RGBA color code, or
	// standard web color name) separated by an underscore. Example: `fol-2_blue`
	// (outline width of 2px and outline color blue), `fol-2_A1CCDD` (outline width of
	// 2px and outline color `#A1CCDD`) and `fol-2_A1CCDD50` (outline width of 2px and
	// outline color `#A1CCDD` at 50% opacity).
	//
	// [Subtitle styling options](https://imagekit.io/docs/add-overlays-on-videos#styling-controls-for-subtitles-layer)
	FontOutline string `json:"fontOutline"`
	// Sets the font shadow for the subtitle text. Requires the shadow color (as an RGB
	// color code, RGBA color code, or standard web color name) and shadow indent (an
	// integer) separated by an underscore. Example: `fsh-blue_2` (shadow color blue,
	// indent of 2px), `fsh-A1CCDD_3` (shadow color `#A1CCDD`, indent of 3px),
	// `fsh-A1CCDD50_3` (shadow color `#A1CCDD` at 50% opacity, indent of 3px).
	//
	// [Subtitle styling options](https://imagekit.io/docs/add-overlays-on-videos#styling-controls-for-subtitles-layer)
	FontShadow string `json:"fontShadow"`
	// Sets the font size of subtitle text.
	//
	// [Subtitle styling options](https://imagekit.io/docs/add-overlays-on-videos#styling-controls-for-subtitles-layer)
	FontSize float64 `json:"fontSize"`
	// Sets the typography style of the subtitle text. Supports values are `b` for
	// bold, `i` for italics, and `b_i` for bold with italics.
	//
	// [Subtitle styling options](https://imagekit.io/docs/add-overlays-on-videos#styling-controls-for-subtitles-layer)
	//
	// Any of "b", "i", "b_i".
	Typography SubtitleOverlayTransformationTypography `json:"typography"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Background  respjson.Field
		Color       respjson.Field
		FontFamily  respjson.Field
		FontOutline respjson.Field
		FontShadow  respjson.Field
		FontSize    respjson.Field
		Typography  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SubtitleOverlayTransformation) RawJSON() string { return r.JSON.raw }
func (r *SubtitleOverlayTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sets the typography style of the subtitle text. Supports values are `b` for
// bold, `i` for italics, and `b_i` for bold with italics.
//
// [Subtitle styling options](https://imagekit.io/docs/add-overlays-on-videos#styling-controls-for-subtitles-layer)
type SubtitleOverlayTransformationTypography string

const (
	SubtitleOverlayTransformationTypographyB  SubtitleOverlayTransformationTypography = "b"
	SubtitleOverlayTransformationTypographyI  SubtitleOverlayTransformationTypography = "i"
	SubtitleOverlayTransformationTypographyBI SubtitleOverlayTransformationTypography = "b_i"
)

type TextOverlay struct {
	// Specifies the text to be displayed in the overlay. The SDK automatically handles
	// special characters and encoding.
	Text string        `json:"text,required"`
	Type constant.Text `json:"type,required"`
	// Text can be included in the layer as either `i-{input}` (plain text) or
	// `ie-{base64_encoded_input}` (base64). By default, the SDK selects the
	// appropriate format based on the input text. To always use base64
	// (`ie-{base64}`), set this parameter to `base64`. To always use plain text
	// (`i-{input}`), set it to `plain`.
	//
	// Any of "auto", "plain", "base64".
	Encoding string `json:"encoding"`
	// Control styling of the text overlay. See
	// [Text overlays](https://imagekit.io/docs/add-overlays-on-images#text-overlay).
	Transformation []TextOverlayTransformation `json:"transformation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text           respjson.Field
		Type           respjson.Field
		Encoding       respjson.Field
		Transformation respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
	BaseOverlay
}

// Returns the unmodified JSON received from the API
func (r TextOverlay) RawJSON() string { return r.JSON.raw }
func (r *TextOverlay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (TextOverlay) ImplOverlayUnion() {}

type TextOverlayTransformation struct {
	// Specifies the transparency level of the text overlay. Accepts integers from `1`
	// to `9`.
	Alpha float64 `json:"alpha"`
	// Specifies the background color of the text overlay. Accepts an RGB hex code, an
	// RGBA code, or a color name.
	Background string `json:"background"`
	// Flip the text overlay horizontally, vertically, or both.
	//
	// Any of "h", "v", "h_v", "v_h".
	Flip TextOverlayTransformationFlip `json:"flip"`
	// Specifies the font color of the overlaid text. Accepts an RGB hex code (e.g.,
	// `FF0000`), an RGBA code (e.g., `FFAABB50`), or a color name.
	FontColor string `json:"fontColor"`
	// Specifies the font family of the overlaid text. Choose from the supported fonts
	// list or use a custom font. See
	// [Supported fonts](https://imagekit.io/docs/add-overlays-on-images#supported-text-font-list)
	// and
	// [Custom font](https://imagekit.io/docs/add-overlays-on-images#change-font-family-in-text-overlay).
	FontFamily string `json:"fontFamily"`
	// Specifies the font size of the overlaid text. Accepts a numeric value or an
	// arithmetic expression.
	FontSize TextOverlayTransformationFontSizeUnion `json:"fontSize"`
	// Specifies the inner alignment of the text when width is more than the text
	// length.
	//
	// Any of "left", "right", "center".
	InnerAlignment TextOverlayTransformationInnerAlignment `json:"innerAlignment"`
	// Specifies the line height of the text overlay. Accepts integer values
	// representing line height in points. It can also accept
	// [arithmetic expressions](https://imagekit.io/docs/arithmetic-expressions-in-transformations)
	// such as `bw_mul_0.2`, or `bh_div_20`.
	LineHeight TextOverlayTransformationLineHeightUnion `json:"lineHeight"`
	// Specifies the padding around the overlaid text. Can be provided as a single
	// positive integer or multiple values separated by underscores (following CSS
	// shorthand order). Arithmetic expressions are also accepted.
	Padding TextOverlayTransformationPaddingUnion `json:"padding"`
	// Specifies the corner radius of the text overlay. Set to `max` to achieve a
	// circular or oval shape.
	Radius TextOverlayTransformationRadiusUnion `json:"radius"`
	// Specifies the rotation angle of the text overlay. Accepts a numeric value for
	// clockwise rotation or a string prefixed with "N" for counter-clockwise rotation.
	Rotation TextOverlayTransformationRotationUnion `json:"rotation"`
	// Specifies the typography style of the text. Supported values:
	//
	//   - Single styles: `b` (bold), `i` (italic), `strikethrough`.
	//   - Combinations: Any combination separated by underscores, e.g., `b_i`,
	//     `b_i_strikethrough`.
	Typography string `json:"typography"`
	// Specifies the maximum width (in pixels) of the overlaid text. The text wraps
	// automatically, and arithmetic expressions (e.g., `bw_mul_0.2` or `bh_div_2`) are
	// supported. Useful when used in conjunction with the `background`. Learn about
	// [Arithmetic expressions](https://imagekit.io/docs/arithmetic-expressions-in-transformations).
	Width TextOverlayTransformationWidthUnion `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Alpha          respjson.Field
		Background     respjson.Field
		Flip           respjson.Field
		FontColor      respjson.Field
		FontFamily     respjson.Field
		FontSize       respjson.Field
		InnerAlignment respjson.Field
		LineHeight     respjson.Field
		Padding        respjson.Field
		Radius         respjson.Field
		Rotation       respjson.Field
		Typography     respjson.Field
		Width          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TextOverlayTransformation) RawJSON() string { return r.JSON.raw }
func (r *TextOverlayTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Flip the text overlay horizontally, vertically, or both.
type TextOverlayTransformationFlip string

const (
	TextOverlayTransformationFlipH  TextOverlayTransformationFlip = "h"
	TextOverlayTransformationFlipV  TextOverlayTransformationFlip = "v"
	TextOverlayTransformationFlipHV TextOverlayTransformationFlip = "h_v"
	TextOverlayTransformationFlipVH TextOverlayTransformationFlip = "v_h"
)

// TextOverlayTransformationFontSizeUnion contains all possible properties and
// values from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type TextOverlayTransformationFontSizeUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u TextOverlayTransformationFontSizeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TextOverlayTransformationFontSizeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TextOverlayTransformationFontSizeUnion) RawJSON() string { return u.JSON.raw }

func (r *TextOverlayTransformationFontSizeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the inner alignment of the text when width is more than the text
// length.
type TextOverlayTransformationInnerAlignment string

const (
	TextOverlayTransformationInnerAlignmentLeft   TextOverlayTransformationInnerAlignment = "left"
	TextOverlayTransformationInnerAlignmentRight  TextOverlayTransformationInnerAlignment = "right"
	TextOverlayTransformationInnerAlignmentCenter TextOverlayTransformationInnerAlignment = "center"
)

// TextOverlayTransformationLineHeightUnion contains all possible properties and
// values from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type TextOverlayTransformationLineHeightUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u TextOverlayTransformationLineHeightUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TextOverlayTransformationLineHeightUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TextOverlayTransformationLineHeightUnion) RawJSON() string { return u.JSON.raw }

func (r *TextOverlayTransformationLineHeightUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TextOverlayTransformationPaddingUnion contains all possible properties and
// values from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type TextOverlayTransformationPaddingUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u TextOverlayTransformationPaddingUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TextOverlayTransformationPaddingUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TextOverlayTransformationPaddingUnion) RawJSON() string { return u.JSON.raw }

func (r *TextOverlayTransformationPaddingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TextOverlayTransformationRadiusUnion contains all possible properties and values
// from [float64], [constant.Max].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfMax]
type TextOverlayTransformationRadiusUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [constant.Max] instead of an
	// object.
	OfMax constant.Max `json:",inline"`
	JSON  struct {
		OfFloat respjson.Field
		OfMax   respjson.Field
		raw     string
	} `json:"-"`
}

func (u TextOverlayTransformationRadiusUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TextOverlayTransformationRadiusUnion) AsMax() (v constant.Max) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TextOverlayTransformationRadiusUnion) RawJSON() string { return u.JSON.raw }

func (r *TextOverlayTransformationRadiusUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TextOverlayTransformationRotationUnion contains all possible properties and
// values from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type TextOverlayTransformationRotationUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u TextOverlayTransformationRotationUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TextOverlayTransformationRotationUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TextOverlayTransformationRotationUnion) RawJSON() string { return u.JSON.raw }

func (r *TextOverlayTransformationRotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TextOverlayTransformationWidthUnion contains all possible properties and values
// from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type TextOverlayTransformationWidthUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u TextOverlayTransformationWidthUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TextOverlayTransformationWidthUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TextOverlayTransformationWidthUnion) RawJSON() string { return u.JSON.raw }

func (r *TextOverlayTransformationWidthUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The SDK provides easy-to-use names for transformations. These names are
// converted to the corresponding transformation string before being added to the
// URL. SDKs are updated regularly to support new transformations. If you want to
// use a transformation that is not supported by the SDK, You can use the `raw`
// parameter to pass the transformation string directly. See the
// [Transformations documentation](https://imagekit.io/docs/transformations).
type Transformation struct {
	// Uses AI to change the background. Provide a text prompt or a base64-encoded
	// prompt, e.g., `prompt-snow road` or `prompte-[urlencoded_base64_encoded_text]`.
	// Not supported inside overlay. See
	// [AI Change Background](https://imagekit.io/docs/ai-transformations#change-background-e-changebg).
	AIChangeBackground string `json:"aiChangeBackground"`
	// Adds an AI-based drop shadow around a foreground object on a transparent or
	// removed background. Optionally, control the direction, elevation, and saturation
	// of the light source (e.g., `az-45` to change light direction). Pass `true` for
	// the default drop shadow, or provide a string for a custom drop shadow. Supported
	// inside overlay. See
	// [AI Drop Shadow](https://imagekit.io/docs/ai-transformations#ai-drop-shadow-e-dropshadow).
	AIDropShadow TransformationAIDropShadowUnion `json:"aiDropShadow"`
	// Uses AI to edit images based on a text prompt. Provide a text prompt or a
	// base64-encoded prompt, e.g., `prompt-snow road` or
	// `prompte-[urlencoded_base64_encoded_text]`. Not supported inside overlay.
	// See [AI Edit](https://imagekit.io/docs/ai-transformations#edit-image-e-edit).
	AIEdit string `json:"aiEdit"`
	// Applies ImageKit's in-house background removal. Supported inside overlay. See
	// [AI Background Removal](https://imagekit.io/docs/ai-transformations#imagekit-background-removal-e-bgremove).
	//
	// Any of true.
	AIRemoveBackground bool `json:"aiRemoveBackground"`
	// Uses third-party background removal. Note: It is recommended to use
	// aiRemoveBackground, ImageKit's in-house solution, which is more cost-effective.
	// Supported inside overlay. See
	// [External Background Removal](https://imagekit.io/docs/ai-transformations#background-removal-e-removedotbg).
	//
	// Any of true.
	AIRemoveBackgroundExternal bool `json:"aiRemoveBackgroundExternal"`
	// Performs AI-based retouching to improve faces or product shots. Not supported
	// inside overlay. See
	// [AI Retouch](https://imagekit.io/docs/ai-transformations#retouch-e-retouch).
	//
	// Any of true.
	AIRetouch bool `json:"aiRetouch"`
	// Upscales images beyond their original dimensions using AI. Not supported inside
	// overlay. See
	// [AI Upscale](https://imagekit.io/docs/ai-transformations#upscale-e-upscale).
	//
	// Any of true.
	AIUpscale bool `json:"aiUpscale"`
	// Generates a variation of an image using AI. This produces a new image with
	// slight variations from the original, such as changes in color, texture, and
	// other visual elements, while preserving the structure and essence of the
	// original image. Not supported inside overlay. See
	// [AI Generate Variations](https://imagekit.io/docs/ai-transformations#generate-variations-of-an-image-e-genvar).
	//
	// Any of true.
	AIVariation bool `json:"aiVariation"`
	// Specifies the aspect ratio for the output, e.g., "ar-4-3". Typically used with
	// either width or height (but not both). For example: aspectRatio = `4:3`, `4_3`,
	// or an expression like `iar_div_2`. See
	// [Image resize and crop – Aspect ratio](https://imagekit.io/docs/image-resize-and-crop#aspect-ratio---ar).
	AspectRatio TransformationAspectRatioUnion `json:"aspectRatio"`
	// Specifies the audio codec, e.g., `aac`, `opus`, or `none`. See
	// [Audio codec](https://imagekit.io/docs/video-optimization#audio-codec---ac).
	//
	// Any of "aac", "opus", "none".
	AudioCodec TransformationAudioCodec `json:"audioCodec"`
	// Specifies the background to be used in conjunction with certain cropping
	// strategies when resizing an image.
	//
	//   - A solid color: e.g., `red`, `F3F3F3`, `AAFF0010`. See
	//     [Solid color background](https://imagekit.io/docs/effects-and-enhancements#solid-color-background).
	//   - A blurred background: e.g., `blurred`, `blurred_25_N15`, etc. See
	//     [Blurred background](https://imagekit.io/docs/effects-and-enhancements#blurred-background).
	//   - Expand the image boundaries using generative fill: `genfill`. Not supported
	//     inside overlay. Optionally, control the background scene by passing a text
	//     prompt: `genfill[:-prompt-${text}]` or
	//     `genfill[:-prompte-${urlencoded_base64_encoded_text}]`. See
	//     [Generative fill background](https://imagekit.io/docs/ai-transformations#generative-fill-bg-genfill).
	Background string `json:"background"`
	// Specifies the Gaussian blur level. Accepts an integer value between 1 and 100,
	// or an expression like `bl-10`. See
	// [Blur](https://imagekit.io/docs/effects-and-enhancements#blur---bl).
	Blur float64 `json:"blur"`
	// Adds a border to the output media. Accepts a string in the format
	// `<border-width>_<hex-code>` (e.g., `5_FFF000` for a 5px yellow border), or an
	// expression like `ih_div_20_FF00FF`. See
	// [Border](https://imagekit.io/docs/effects-and-enhancements#border---b).
	Border string `json:"border"`
	// Indicates whether the output image should retain the original color profile. See
	// [Color profile](https://imagekit.io/docs/image-optimization#color-profile---cp).
	ColorProfile bool `json:"colorProfile"`
	// Automatically enhances the contrast of an image (contrast stretch). See
	// [Contrast Stretch](https://imagekit.io/docs/effects-and-enhancements#contrast-stretch---e-contrast).
	//
	// Any of true.
	ContrastStretch bool `json:"contrastStretch"`
	// Crop modes for image resizing. See
	// [Crop modes & focus](https://imagekit.io/docs/image-resize-and-crop#crop-crop-modes--focus).
	//
	// Any of "force", "at_max", "at_max_enlarge", "at_least", "maintain_ratio".
	Crop TransformationCrop `json:"crop"`
	// Additional crop modes for image resizing. See
	// [Crop modes & focus](https://imagekit.io/docs/image-resize-and-crop#crop-crop-modes--focus).
	//
	// Any of "pad_resize", "extract", "pad_extract".
	CropMode TransformationCropMode `json:"cropMode"`
	// Specifies a fallback image if the resource is not found, e.g., a URL or file
	// path. See
	// [Default image](https://imagekit.io/docs/image-transformation#default-image---di).
	DefaultImage string `json:"defaultImage"`
	// Accepts values between 0.1 and 5, or `auto` for automatic device pixel ratio
	// (DPR) calculation. See
	// [DPR](https://imagekit.io/docs/image-resize-and-crop#dpr---dpr).
	Dpr float64 `json:"dpr"`
	// Specifies the duration (in seconds) for trimming videos, e.g., `5` or `10.5`.
	// Typically used with startOffset to indicate the length from the start offset.
	// Arithmetic expressions are supported. See
	// [Trim videos – Duration](https://imagekit.io/docs/trim-videos#duration---du).
	Duration TransformationDurationUnion `json:"duration"`
	// Specifies the end offset (in seconds) for trimming videos, e.g., `5` or `10.5`.
	// Typically used with startOffset to define a time window. Arithmetic expressions
	// are supported. See
	// [Trim videos – End offset](https://imagekit.io/docs/trim-videos#end-offset---eo).
	EndOffset TransformationEndOffsetUnion `json:"endOffset"`
	// Flips or mirrors an image either horizontally, vertically, or both. Acceptable
	// values: `h` (horizontal), `v` (vertical), `h_v` (horizontal and vertical), or
	// `v_h`. See [Flip](https://imagekit.io/docs/effects-and-enhancements#flip---fl).
	//
	// Any of "h", "v", "h_v", "v_h".
	Flip TransformationFlip `json:"flip"`
	// Refines padding and cropping behavior for pad resize, maintain ratio, and
	// extract crop modes. Supports manual positions and coordinate-based focus. With
	// AI-based cropping, you can automatically keep key subjects in frame—such as
	// faces or detected objects (e.g., `fo-face`, `fo-person`, `fo-car`)— while
	// resizing.
	//
	// - See [Focus](https://imagekit.io/docs/image-resize-and-crop#focus---fo).
	// - [Object aware cropping](https://imagekit.io/docs/image-resize-and-crop#object-aware-cropping---fo-object-name)
	Focus string `json:"focus"`
	// Specifies the output format for images or videos, e.g., `jpg`, `png`, `webp`,
	// `mp4`, or `auto`. You can also pass `orig` for images to return the original
	// format. ImageKit automatically delivers images and videos in the optimal format
	// based on device support unless overridden by the dashboard settings or the
	// format parameter. See
	// [Image format](https://imagekit.io/docs/image-optimization#format---f) and
	// [Video format](https://imagekit.io/docs/video-optimization#format---f).
	//
	// Any of "auto", "webp", "jpg", "jpeg", "png", "gif", "svg", "mp4", "webm",
	// "avif", "orig".
	Format TransformationFormat `json:"format"`
	// Creates a linear gradient with two colors. Pass `true` for a default gradient,
	// or provide a string for a custom gradient. See
	// [Gradient](https://imagekit.io/docs/effects-and-enhancements#gradient---e-gradient).
	Gradient TransformationGradientUnion `json:"gradient"`
	// Enables a grayscale effect for images. See
	// [Grayscale](https://imagekit.io/docs/effects-and-enhancements#grayscale---e-grayscale).
	//
	// Any of true.
	Grayscale bool `json:"grayscale"`
	// Specifies the height of the output. If a value between 0 and 1 is provided, it
	// is treated as a percentage (e.g., `0.5` represents 50% of the original height).
	// You can also supply arithmetic expressions (e.g., `ih_mul_0.5`). Height
	// transformation –
	// [Images](https://imagekit.io/docs/image-resize-and-crop#height---h) ·
	// [Videos](https://imagekit.io/docs/video-resize-and-crop#height---h)
	Height TransformationHeightUnion `json:"height"`
	// Specifies whether the output image (in JPEG or PNG) should be compressed
	// losslessly. See
	// [Lossless compression](https://imagekit.io/docs/image-optimization#lossless-webp-and-png---lo).
	Lossless bool `json:"lossless"`
	// By default, ImageKit removes all metadata during automatic image compression.
	// Set this to true to preserve metadata. See
	// [Image metadata](https://imagekit.io/docs/image-optimization#image-metadata---md).
	Metadata bool `json:"metadata"`
	// Named transformation reference. See
	// [Named transformations](https://imagekit.io/docs/transformations#named-transformations).
	Named string `json:"named"`
	// Specifies the opacity level of the output image. See
	// [Opacity](https://imagekit.io/docs/effects-and-enhancements#opacity---o).
	Opacity float64 `json:"opacity"`
	// If set to true, serves the original file without applying any transformations.
	// See
	// [Deliver original file as-is](https://imagekit.io/docs/core-delivery-features#deliver-original-file-as-is---orig-true).
	Original bool `json:"original"`
	// Specifies an overlay to be applied on the parent image or video. ImageKit
	// supports overlays including images, text, videos, subtitles, and solid colors.
	// See
	// [Overlay using layers](https://imagekit.io/docs/transformations#overlay-using-layers).
	Overlay OverlayUnion `json:"overlay"`
	// Extracts a specific page or frame from multi-page or layered files (PDF, PSD,
	// AI). For example, specify by number (e.g., `2`), a range (e.g., `3-4` for the
	// 2nd and 3rd layers), or by name (e.g., `name-layer-4` for a PSD layer). See
	// [Thumbnail extraction](https://imagekit.io/docs/vector-and-animated-images#get-thumbnail-from-psd-pdf-ai-eps-and-animated-files).
	Page TransformationPageUnion `json:"page"`
	// Specifies whether the output JPEG image should be rendered progressively.
	// Progressive loading begins with a low-quality, pixelated version of the full
	// image, which gradually improves to provide a faster perceived load time. See
	// [Progressive images](https://imagekit.io/docs/image-optimization#progressive-image---pr).
	Progressive bool `json:"progressive"`
	// Specifies the quality of the output image for lossy formats such as JPEG, WebP,
	// and AVIF. A higher quality value results in a larger file size with better
	// quality, while a lower value produces a smaller file size with reduced quality.
	// See [Quality](https://imagekit.io/docs/image-optimization#quality---q).
	Quality float64 `json:"quality"`
	// Specifies the corner radius for rounded corners (e.g., 20) or `max` for circular
	// or oval shape. See
	// [Radius](https://imagekit.io/docs/effects-and-enhancements#radius---r).
	Radius TransformationRadiusUnion `json:"radius"`
	// Pass any transformation not directly supported by the SDK. This transformation
	// string is appended to the URL as provided.
	Raw string `json:"raw"`
	// Specifies the rotation angle in degrees. Positive values rotate the image
	// clockwise; you can also use, for example, `N40` for counterclockwise rotation or
	// `auto` to use the orientation specified in the image's EXIF data. For videos,
	// only the following values are supported: 0, 90, 180, 270, or 360. See
	// [Rotate](https://imagekit.io/docs/effects-and-enhancements#rotate---rt).
	Rotation TransformationRotationUnion `json:"rotation"`
	// Adds a shadow beneath solid objects in an image with a transparent background.
	// For AI-based drop shadows, refer to aiDropShadow. Pass `true` for a default
	// shadow, or provide a string for a custom shadow. See
	// [Shadow](https://imagekit.io/docs/effects-and-enhancements#shadow---e-shadow).
	Shadow TransformationShadowUnion `json:"shadow"`
	// Sharpens the input image, highlighting edges and finer details. Pass `true` for
	// default sharpening, or provide a numeric value for custom sharpening. See
	// [Sharpen](https://imagekit.io/docs/effects-and-enhancements#sharpen---e-sharpen).
	Sharpen TransformationSharpenUnion `json:"sharpen"`
	// Specifies the start offset (in seconds) for trimming videos, e.g., `5` or
	// `10.5`. Arithmetic expressions are also supported. See
	// [Trim videos – Start offset](https://imagekit.io/docs/trim-videos#start-offset---so).
	StartOffset TransformationStartOffsetUnion `json:"startOffset"`
	// An array of resolutions for adaptive bitrate streaming, e.g., [`240`, `360`,
	// `480`, `720`, `1080`]. See
	// [Adaptive Bitrate Streaming](https://imagekit.io/docs/adaptive-bitrate-streaming).
	StreamingResolutions []StreamingResolution `json:"streamingResolutions"`
	// Useful for images with a solid or nearly solid background and a central object.
	// This parameter trims the background, leaving only the central object in the
	// output image. See
	// [Trim edges](https://imagekit.io/docs/effects-and-enhancements#trim-edges---t).
	Trim TransformationTrimUnion `json:"trim"`
	// Applies Unsharp Masking (USM), an image sharpening technique. Pass `true` for a
	// default unsharp mask, or provide a string for a custom unsharp mask. See
	// [Unsharp Mask](https://imagekit.io/docs/effects-and-enhancements#unsharp-mask---e-usm).
	UnsharpMask TransformationUnsharpMaskUnion `json:"unsharpMask"`
	// Specifies the video codec, e.g., `h264`, `vp9`, `av1`, or `none`. See
	// [Video codec](https://imagekit.io/docs/video-optimization#video-codec---vc).
	//
	// Any of "h264", "vp9", "av1", "none".
	VideoCodec TransformationVideoCodec `json:"videoCodec"`
	// Specifies the width of the output. If a value between 0 and 1 is provided, it is
	// treated as a percentage (e.g., `0.4` represents 40% of the original width). You
	// can also supply arithmetic expressions (e.g., `iw_div_2`). Width transformation
	// – [Images](https://imagekit.io/docs/image-resize-and-crop#width---w) ·
	// [Videos](https://imagekit.io/docs/video-resize-and-crop#width---w)
	Width TransformationWidthUnion `json:"width"`
	// Focus using cropped image coordinates - X coordinate. See
	// [Focus using cropped coordinates](https://imagekit.io/docs/image-resize-and-crop#example---focus-using-cropped-image-coordinates).
	X TransformationXUnion `json:"x"`
	// Focus using cropped image coordinates - X center coordinate. See
	// [Focus using cropped coordinates](https://imagekit.io/docs/image-resize-and-crop#example---focus-using-cropped-image-coordinates).
	XCenter TransformationXCenterUnion `json:"xCenter"`
	// Focus using cropped image coordinates - Y coordinate. See
	// [Focus using cropped coordinates](https://imagekit.io/docs/image-resize-and-crop#example---focus-using-cropped-image-coordinates).
	Y TransformationYUnion `json:"y"`
	// Focus using cropped image coordinates - Y center coordinate. See
	// [Focus using cropped coordinates](https://imagekit.io/docs/image-resize-and-crop#example---focus-using-cropped-image-coordinates).
	YCenter TransformationYCenterUnion `json:"yCenter"`
	// Accepts a numeric value that determines how much to zoom in or out of the
	// cropped area. It should be used in conjunction with fo-face or fo-<object_name>.
	// See [Zoom](https://imagekit.io/docs/image-resize-and-crop#zoom---z).
	Zoom float64 `json:"zoom"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AIChangeBackground         respjson.Field
		AIDropShadow               respjson.Field
		AIEdit                     respjson.Field
		AIRemoveBackground         respjson.Field
		AIRemoveBackgroundExternal respjson.Field
		AIRetouch                  respjson.Field
		AIUpscale                  respjson.Field
		AIVariation                respjson.Field
		AspectRatio                respjson.Field
		AudioCodec                 respjson.Field
		Background                 respjson.Field
		Blur                       respjson.Field
		Border                     respjson.Field
		ColorProfile               respjson.Field
		ContrastStretch            respjson.Field
		Crop                       respjson.Field
		CropMode                   respjson.Field
		DefaultImage               respjson.Field
		Dpr                        respjson.Field
		Duration                   respjson.Field
		EndOffset                  respjson.Field
		Flip                       respjson.Field
		Focus                      respjson.Field
		Format                     respjson.Field
		Gradient                   respjson.Field
		Grayscale                  respjson.Field
		Height                     respjson.Field
		Lossless                   respjson.Field
		Metadata                   respjson.Field
		Named                      respjson.Field
		Opacity                    respjson.Field
		Original                   respjson.Field
		Overlay                    respjson.Field
		Page                       respjson.Field
		Progressive                respjson.Field
		Quality                    respjson.Field
		Radius                     respjson.Field
		Raw                        respjson.Field
		Rotation                   respjson.Field
		Shadow                     respjson.Field
		Sharpen                    respjson.Field
		StartOffset                respjson.Field
		StreamingResolutions       respjson.Field
		Trim                       respjson.Field
		UnsharpMask                respjson.Field
		VideoCodec                 respjson.Field
		Width                      respjson.Field
		X                          respjson.Field
		XCenter                    respjson.Field
		Y                          respjson.Field
		YCenter                    respjson.Field
		Zoom                       respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Transformation) RawJSON() string { return r.JSON.raw }
func (r *Transformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TransformationAIDropShadowUnion contains all possible properties and values from
// [bool], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfTransformationAIDropShadowBoolean OfString]
type TransformationAIDropShadowUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfTransformationAIDropShadowBoolean bool `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfTransformationAIDropShadowBoolean respjson.Field
		OfString                            respjson.Field
		raw                                 string
	} `json:"-"`
}

func (u TransformationAIDropShadowUnion) AsTransformationAIDropShadowBoolean() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TransformationAIDropShadowUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TransformationAIDropShadowUnion) RawJSON() string { return u.JSON.raw }

func (r *TransformationAIDropShadowUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransformationAIDropShadowBoolean bool

const (
	TransformationAIDropShadowBooleanTrue TransformationAIDropShadowBoolean = true
)

// TransformationAspectRatioUnion contains all possible properties and values from
// [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type TransformationAspectRatioUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u TransformationAspectRatioUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TransformationAspectRatioUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TransformationAspectRatioUnion) RawJSON() string { return u.JSON.raw }

func (r *TransformationAspectRatioUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the audio codec, e.g., `aac`, `opus`, or `none`. See
// [Audio codec](https://imagekit.io/docs/video-optimization#audio-codec---ac).
type TransformationAudioCodec string

const (
	TransformationAudioCodecAac  TransformationAudioCodec = "aac"
	TransformationAudioCodecOpus TransformationAudioCodec = "opus"
	TransformationAudioCodecNone TransformationAudioCodec = "none"
)

// Crop modes for image resizing. See
// [Crop modes & focus](https://imagekit.io/docs/image-resize-and-crop#crop-crop-modes--focus).
type TransformationCrop string

const (
	TransformationCropForce         TransformationCrop = "force"
	TransformationCropAtMax         TransformationCrop = "at_max"
	TransformationCropAtMaxEnlarge  TransformationCrop = "at_max_enlarge"
	TransformationCropAtLeast       TransformationCrop = "at_least"
	TransformationCropMaintainRatio TransformationCrop = "maintain_ratio"
)

// Additional crop modes for image resizing. See
// [Crop modes & focus](https://imagekit.io/docs/image-resize-and-crop#crop-crop-modes--focus).
type TransformationCropMode string

const (
	TransformationCropModePadResize  TransformationCropMode = "pad_resize"
	TransformationCropModeExtract    TransformationCropMode = "extract"
	TransformationCropModePadExtract TransformationCropMode = "pad_extract"
)

// TransformationDurationUnion contains all possible properties and values from
// [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type TransformationDurationUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u TransformationDurationUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TransformationDurationUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TransformationDurationUnion) RawJSON() string { return u.JSON.raw }

func (r *TransformationDurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TransformationEndOffsetUnion contains all possible properties and values from
// [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type TransformationEndOffsetUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u TransformationEndOffsetUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TransformationEndOffsetUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TransformationEndOffsetUnion) RawJSON() string { return u.JSON.raw }

func (r *TransformationEndOffsetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Flips or mirrors an image either horizontally, vertically, or both. Acceptable
// values: `h` (horizontal), `v` (vertical), `h_v` (horizontal and vertical), or
// `v_h`. See [Flip](https://imagekit.io/docs/effects-and-enhancements#flip---fl).
type TransformationFlip string

const (
	TransformationFlipH  TransformationFlip = "h"
	TransformationFlipV  TransformationFlip = "v"
	TransformationFlipHV TransformationFlip = "h_v"
	TransformationFlipVH TransformationFlip = "v_h"
)

// Specifies the output format for images or videos, e.g., `jpg`, `png`, `webp`,
// `mp4`, or `auto`. You can also pass `orig` for images to return the original
// format. ImageKit automatically delivers images and videos in the optimal format
// based on device support unless overridden by the dashboard settings or the
// format parameter. See
// [Image format](https://imagekit.io/docs/image-optimization#format---f) and
// [Video format](https://imagekit.io/docs/video-optimization#format---f).
type TransformationFormat string

const (
	TransformationFormatAuto TransformationFormat = "auto"
	TransformationFormatWebp TransformationFormat = "webp"
	TransformationFormatJpg  TransformationFormat = "jpg"
	TransformationFormatJpeg TransformationFormat = "jpeg"
	TransformationFormatPng  TransformationFormat = "png"
	TransformationFormatGif  TransformationFormat = "gif"
	TransformationFormatSvg  TransformationFormat = "svg"
	TransformationFormatMP4  TransformationFormat = "mp4"
	TransformationFormatWebm TransformationFormat = "webm"
	TransformationFormatAvif TransformationFormat = "avif"
	TransformationFormatOrig TransformationFormat = "orig"
)

// TransformationGradientUnion contains all possible properties and values from
// [bool], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfTransformationGradientBoolean OfString]
type TransformationGradientUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfTransformationGradientBoolean bool `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfTransformationGradientBoolean respjson.Field
		OfString                        respjson.Field
		raw                             string
	} `json:"-"`
}

func (u TransformationGradientUnion) AsTransformationGradientBoolean() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TransformationGradientUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TransformationGradientUnion) RawJSON() string { return u.JSON.raw }

func (r *TransformationGradientUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransformationGradientBoolean bool

const (
	TransformationGradientBooleanTrue TransformationGradientBoolean = true
)

// TransformationHeightUnion contains all possible properties and values from
// [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type TransformationHeightUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u TransformationHeightUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TransformationHeightUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TransformationHeightUnion) RawJSON() string { return u.JSON.raw }

func (r *TransformationHeightUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TransformationPageUnion contains all possible properties and values from
// [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type TransformationPageUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u TransformationPageUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TransformationPageUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TransformationPageUnion) RawJSON() string { return u.JSON.raw }

func (r *TransformationPageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TransformationRadiusUnion contains all possible properties and values from
// [float64], [constant.Max].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfMax]
type TransformationRadiusUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [constant.Max] instead of an
	// object.
	OfMax constant.Max `json:",inline"`
	JSON  struct {
		OfFloat respjson.Field
		OfMax   respjson.Field
		raw     string
	} `json:"-"`
}

func (u TransformationRadiusUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TransformationRadiusUnion) AsMax() (v constant.Max) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TransformationRadiusUnion) RawJSON() string { return u.JSON.raw }

func (r *TransformationRadiusUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TransformationRotationUnion contains all possible properties and values from
// [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type TransformationRotationUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u TransformationRotationUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TransformationRotationUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TransformationRotationUnion) RawJSON() string { return u.JSON.raw }

func (r *TransformationRotationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TransformationShadowUnion contains all possible properties and values from
// [bool], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfTransformationShadowBoolean OfString]
type TransformationShadowUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfTransformationShadowBoolean bool `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfTransformationShadowBoolean respjson.Field
		OfString                      respjson.Field
		raw                           string
	} `json:"-"`
}

func (u TransformationShadowUnion) AsTransformationShadowBoolean() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TransformationShadowUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TransformationShadowUnion) RawJSON() string { return u.JSON.raw }

func (r *TransformationShadowUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransformationShadowBoolean bool

const (
	TransformationShadowBooleanTrue TransformationShadowBoolean = true
)

// TransformationSharpenUnion contains all possible properties and values from
// [bool], [float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfTransformationSharpenBoolean OfFloat]
type TransformationSharpenUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfTransformationSharpenBoolean bool `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	JSON    struct {
		OfTransformationSharpenBoolean respjson.Field
		OfFloat                        respjson.Field
		raw                            string
	} `json:"-"`
}

func (u TransformationSharpenUnion) AsTransformationSharpenBoolean() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TransformationSharpenUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TransformationSharpenUnion) RawJSON() string { return u.JSON.raw }

func (r *TransformationSharpenUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransformationSharpenBoolean bool

const (
	TransformationSharpenBooleanTrue TransformationSharpenBoolean = true
)

// TransformationStartOffsetUnion contains all possible properties and values from
// [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type TransformationStartOffsetUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u TransformationStartOffsetUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TransformationStartOffsetUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TransformationStartOffsetUnion) RawJSON() string { return u.JSON.raw }

func (r *TransformationStartOffsetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TransformationTrimUnion contains all possible properties and values from [bool],
// [float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfTransformationTrimBoolean OfFloat]
type TransformationTrimUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfTransformationTrimBoolean bool `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	JSON    struct {
		OfTransformationTrimBoolean respjson.Field
		OfFloat                     respjson.Field
		raw                         string
	} `json:"-"`
}

func (u TransformationTrimUnion) AsTransformationTrimBoolean() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TransformationTrimUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TransformationTrimUnion) RawJSON() string { return u.JSON.raw }

func (r *TransformationTrimUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransformationTrimBoolean bool

const (
	TransformationTrimBooleanTrue TransformationTrimBoolean = true
)

// TransformationUnsharpMaskUnion contains all possible properties and values from
// [bool], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfTransformationUnsharpMaskBoolean OfString]
type TransformationUnsharpMaskUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfTransformationUnsharpMaskBoolean bool `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfTransformationUnsharpMaskBoolean respjson.Field
		OfString                           respjson.Field
		raw                                string
	} `json:"-"`
}

func (u TransformationUnsharpMaskUnion) AsTransformationUnsharpMaskBoolean() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TransformationUnsharpMaskUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TransformationUnsharpMaskUnion) RawJSON() string { return u.JSON.raw }

func (r *TransformationUnsharpMaskUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransformationUnsharpMaskBoolean bool

const (
	TransformationUnsharpMaskBooleanTrue TransformationUnsharpMaskBoolean = true
)

// Specifies the video codec, e.g., `h264`, `vp9`, `av1`, or `none`. See
// [Video codec](https://imagekit.io/docs/video-optimization#video-codec---vc).
type TransformationVideoCodec string

const (
	TransformationVideoCodecH264 TransformationVideoCodec = "h264"
	TransformationVideoCodecVp9  TransformationVideoCodec = "vp9"
	TransformationVideoCodecAv1  TransformationVideoCodec = "av1"
	TransformationVideoCodecNone TransformationVideoCodec = "none"
)

// TransformationWidthUnion contains all possible properties and values from
// [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type TransformationWidthUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u TransformationWidthUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TransformationWidthUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TransformationWidthUnion) RawJSON() string { return u.JSON.raw }

func (r *TransformationWidthUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TransformationXUnion contains all possible properties and values from [float64],
// [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type TransformationXUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u TransformationXUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TransformationXUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TransformationXUnion) RawJSON() string { return u.JSON.raw }

func (r *TransformationXUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TransformationXCenterUnion contains all possible properties and values from
// [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type TransformationXCenterUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u TransformationXCenterUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TransformationXCenterUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TransformationXCenterUnion) RawJSON() string { return u.JSON.raw }

func (r *TransformationXCenterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TransformationYUnion contains all possible properties and values from [float64],
// [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type TransformationYUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u TransformationYUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TransformationYUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TransformationYUnion) RawJSON() string { return u.JSON.raw }

func (r *TransformationYUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TransformationYCenterUnion contains all possible properties and values from
// [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type TransformationYCenterUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u TransformationYCenterUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TransformationYCenterUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TransformationYCenterUnion) RawJSON() string { return u.JSON.raw }

func (r *TransformationYCenterUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// By default, the transformation string is added as a query parameter in the URL,
// e.g., `?tr=w-100,h-100`. If you want to add the transformation string in the
// path of the URL, set this to `path`. Learn more in the
// [Transformations guide](https://imagekit.io/docs/transformations).
type TransformationPosition string

const (
	TransformationPositionPath  TransformationPosition = "path"
	TransformationPositionQuery TransformationPosition = "query"
)

type VideoOverlay struct {
	// Specifies the relative path to the video used as an overlay.
	Input string         `json:"input,required"`
	Type  constant.Video `json:"type,required"`
	// The input path can be included in the layer as either `i-{input}` or
	// `ie-{base64_encoded_input}`. By default, the SDK determines the appropriate
	// format automatically. To always use base64 encoding (`ie-{base64}`), set this
	// parameter to `base64`. To always use plain text (`i-{input}`), set it to
	// `plain`.
	//
	// Any of "auto", "plain", "base64".
	Encoding string `json:"encoding"`
	// Array of transformation to be applied to the overlay video. Except
	// `streamingResolutions`, all other video transformations are supported. See
	// [Video transformations](https://imagekit.io/docs/video-transformation).
	Transformation []Transformation `json:"transformation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Input          respjson.Field
		Type           respjson.Field
		Encoding       respjson.Field
		Transformation respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
	BaseOverlay
}

// Returns the unmodified JSON received from the API
func (r VideoOverlay) RawJSON() string { return r.JSON.raw }
func (r *VideoOverlay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (VideoOverlay) ImplOverlayUnion() {}
