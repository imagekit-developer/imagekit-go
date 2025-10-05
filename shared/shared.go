// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/imagekit-developer/imagekit-go/internal/apijson"
	"github.com/imagekit-developer/imagekit-go/packages/param"
	"github.com/imagekit-developer/imagekit-go/shared/constant"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type BaseOverlayParam struct {
	Position OverlayPositionParam `json:"position,omitzero"`
	Timing   OverlayTimingParam   `json:"timing,omitzero"`
	paramObj
}

func (r BaseOverlayParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseOverlayParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseOverlayParam) UnmarshalJSON(data []byte) error {
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

type ImageOverlayParam struct {
	// Specifies the relative path to the image used as an overlay.
	Input string         `json:"input,required"`
	Type  constant.Image `json:"type,required"`
	// The input path can be included in the layer as either `i-{input}` or
	// `ie-{base64_encoded_input}`. By default, the SDK determines the appropriate
	// format automatically. To always use base64 encoding (`ie-{base64}`), set this
	// parameter to `base64`. To always use plain text (`i-{input}`), set it to
	// `plain`.
	Encoding string `json:"encoding,omitzero"`
	// Array of transformations to be applied to the overlay image. Supported
	// transformations depends on the base/parent asset. See overlays on
	// [Images](https://imagekit.io/docs/add-overlays-on-images#list-of-supported-image-transformations-in-image-layers)
	// and
	// [Videos](https://imagekit.io/docs/add-overlays-on-videos#list-of-transformations-supported-on-image-overlay).
	Transformation []TransformationParam `json:"transformation,omitzero"`
	BaseOverlayParam
}

func (r ImageOverlayParam) MarshalJSON() (data []byte, err error) {
	type shadow ImageOverlayParam
	return param.MarshalObject(r, (*shadow)(&r))
}

func OverlayParamOfText(text string) OverlayUnionParam {
	var variant TextOverlayParam
	variant.Text = text
	return OverlayUnionParam{OfText: &variant}
}

func OverlayParamOfImage(input string) OverlayUnionParam {
	var image ImageOverlayParam
	image.Input = input
	return OverlayUnionParam{OfImage: &image}
}

func OverlayParamOfVideo(input string) OverlayUnionParam {
	var video VideoOverlayParam
	video.Input = input
	return OverlayUnionParam{OfVideo: &video}
}

func OverlayParamOfSubtitle(input string) OverlayUnionParam {
	var subtitle SubtitleOverlayParam
	subtitle.Input = input
	return OverlayUnionParam{OfSubtitle: &subtitle}
}

func OverlayParamOfSolidColor(color string) OverlayUnionParam {
	var solidColor SolidColorOverlayParam
	solidColor.Color = color
	return OverlayUnionParam{OfSolidColor: &solidColor}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OverlayUnionParam struct {
	OfText       *TextOverlayParam       `json:",omitzero,inline"`
	OfImage      *ImageOverlayParam      `json:",omitzero,inline"`
	OfVideo      *VideoOverlayParam      `json:",omitzero,inline"`
	OfSubtitle   *SubtitleOverlayParam   `json:",omitzero,inline"`
	OfSolidColor *SolidColorOverlayParam `json:",omitzero,inline"`
	paramUnion
}

func (u OverlayUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfText,
		u.OfImage,
		u.OfVideo,
		u.OfSubtitle,
		u.OfSolidColor)
}
func (u *OverlayUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OverlayUnionParam) asAny() any {
	if !param.IsOmitted(u.OfText) {
		return u.OfText
	} else if !param.IsOmitted(u.OfImage) {
		return u.OfImage
	} else if !param.IsOmitted(u.OfVideo) {
		return u.OfVideo
	} else if !param.IsOmitted(u.OfSubtitle) {
		return u.OfSubtitle
	} else if !param.IsOmitted(u.OfSolidColor) {
		return u.OfSolidColor
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OverlayUnionParam) GetText() *string {
	if vt := u.OfText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OverlayUnionParam) GetColor() *string {
	if vt := u.OfSolidColor; vt != nil {
		return &vt.Color
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OverlayUnionParam) GetType() *string {
	if vt := u.OfText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfImage; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfVideo; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSubtitle; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSolidColor; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OverlayUnionParam) GetEncoding() *string {
	if vt := u.OfText; vt != nil {
		return (*string)(&vt.Encoding)
	} else if vt := u.OfImage; vt != nil {
		return (*string)(&vt.Encoding)
	} else if vt := u.OfVideo; vt != nil {
		return (*string)(&vt.Encoding)
	} else if vt := u.OfSubtitle; vt != nil {
		return (*string)(&vt.Encoding)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OverlayUnionParam) GetInput() *string {
	if vt := u.OfImage; vt != nil {
		return (*string)(&vt.Input)
	} else if vt := u.OfVideo; vt != nil {
		return (*string)(&vt.Input)
	} else if vt := u.OfSubtitle; vt != nil {
		return (*string)(&vt.Input)
	}
	return nil
}

// Returns a pointer to the underlying variant's Position property, if present.
func (u OverlayUnionParam) GetPosition() *OverlayPositionParam {
	if vt := u.OfText; vt != nil {
		return &vt.Position
	} else if vt := u.OfImage; vt != nil {
		return &vt.Position
	} else if vt := u.OfVideo; vt != nil {
		return &vt.Position
	} else if vt := u.OfSubtitle; vt != nil {
		return &vt.Position
	} else if vt := u.OfSolidColor; vt != nil {
		return &vt.Position
	}
	return nil
}

// Returns a pointer to the underlying variant's Timing property, if present.
func (u OverlayUnionParam) GetTiming() *OverlayTimingParam {
	if vt := u.OfText; vt != nil {
		return &vt.Timing
	} else if vt := u.OfImage; vt != nil {
		return &vt.Timing
	} else if vt := u.OfVideo; vt != nil {
		return &vt.Timing
	} else if vt := u.OfSubtitle; vt != nil {
		return &vt.Timing
	} else if vt := u.OfSolidColor; vt != nil {
		return &vt.Timing
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u OverlayUnionParam) GetTransformation() (res overlayUnionParamTransformation) {
	if vt := u.OfText; vt != nil {
		res.any = &vt.Transformation
	} else if vt := u.OfImage; vt != nil {
		res.any = &vt.Transformation
	} else if vt := u.OfVideo; vt != nil {
		res.any = &vt.Transformation
	} else if vt := u.OfSubtitle; vt != nil {
		res.any = &vt.Transformation
	} else if vt := u.OfSolidColor; vt != nil {
		res.any = &vt.Transformation
	}
	return
}

// Can have the runtime types [_[]TextOverlayTransformationParam],
// [_[]TransformationParam], [_[]SubtitleOverlayTransformationParam],
// [_[]SolidColorOverlayTransformationParam]
type overlayUnionParamTransformation struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]shared.TextOverlayTransformationParam:
//	case *[]shared.TransformationParam:
//	case *[]shared.SubtitleOverlayTransformationParam:
//	case *[]shared.SolidColorOverlayTransformationParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u overlayUnionParamTransformation) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[OverlayUnionParam](
		"type",
		apijson.Discriminator[TextOverlayParam]("text"),
		apijson.Discriminator[ImageOverlayParam]("image"),
		apijson.Discriminator[VideoOverlayParam]("video"),
		apijson.Discriminator[SubtitleOverlayParam]("subtitle"),
		apijson.Discriminator[SolidColorOverlayParam]("solidColor"),
	)
}

type OverlayPositionParam struct {
	// Specifies the position of the overlay relative to the parent image or video.
	// Maps to `lfo` in the URL.
	//
	// Any of "center", "top", "left", "bottom", "right", "top_left", "top_right",
	// "bottom_left", "bottom_right".
	Focus OverlayPositionFocus `json:"focus,omitzero"`
	// Specifies the x-coordinate of the top-left corner of the base asset where the
	// overlay's top-left corner will be positioned. It also accepts arithmetic
	// expressions such as `bw_mul_0.4` or `bw_sub_cw`. Maps to `lx` in the URL. Learn
	// about
	// [Arithmetic expressions](https://imagekit.io/docs/arithmetic-expressions-in-transformations).
	X OverlayPositionXUnionParam `json:"x,omitzero"`
	// Specifies the y-coordinate of the top-left corner of the base asset where the
	// overlay's top-left corner will be positioned. It also accepts arithmetic
	// expressions such as `bh_mul_0.4` or `bh_sub_ch`. Maps to `ly` in the URL. Learn
	// about
	// [Arithmetic expressions](https://imagekit.io/docs/arithmetic-expressions-in-transformations).
	Y OverlayPositionYUnionParam `json:"y,omitzero"`
	paramObj
}

func (r OverlayPositionParam) MarshalJSON() (data []byte, err error) {
	type shadow OverlayPositionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OverlayPositionParam) UnmarshalJSON(data []byte) error {
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OverlayPositionXUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u OverlayPositionXUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *OverlayPositionXUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OverlayPositionXUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OverlayPositionYUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u OverlayPositionYUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *OverlayPositionYUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OverlayPositionYUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

type OverlayTimingParam struct {
	// Specifies the duration (in seconds) during which the overlay should appear on
	// the base video. Accepts a positive number up to two decimal places (e.g., `20`
	// or `20.50`) and arithmetic expressions such as `bdu_mul_0.4` or `bdu_sub_idu`.
	// Applies only if the base asset is a video. Maps to `ldu` in the URL.
	Duration OverlayTimingDurationUnionParam `json:"duration,omitzero"`
	// Specifies the end time (in seconds) for when the overlay should disappear from
	// the base video. If both end and duration are provided, duration is ignored.
	// Accepts a positive number up to two decimal places (e.g., `20` or `20.50`) and
	// arithmetic expressions such as `bdu_mul_0.4` or `bdu_sub_idu`. Applies only if
	// the base asset is a video. Maps to `leo` in the URL.
	End OverlayTimingEndUnionParam `json:"end,omitzero"`
	// Specifies the start time (in seconds) for when the overlay should appear on the
	// base video. Accepts a positive number up to two decimal places (e.g., `20` or
	// `20.50`) and arithmetic expressions such as `bdu_mul_0.4` or `bdu_sub_idu`.
	// Applies only if the base asset is a video. Maps to `lso` in the URL.
	Start OverlayTimingStartUnionParam `json:"start,omitzero"`
	paramObj
}

func (r OverlayTimingParam) MarshalJSON() (data []byte, err error) {
	type shadow OverlayTimingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OverlayTimingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OverlayTimingDurationUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u OverlayTimingDurationUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *OverlayTimingDurationUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OverlayTimingDurationUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OverlayTimingEndUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u OverlayTimingEndUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *OverlayTimingEndUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OverlayTimingEndUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OverlayTimingStartUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u OverlayTimingStartUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *OverlayTimingStartUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OverlayTimingStartUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

type SolidColorOverlayParam struct {
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
	Transformation []SolidColorOverlayTransformationParam `json:"transformation,omitzero"`
	BaseOverlayParam
}

func (r SolidColorOverlayParam) MarshalJSON() (data []byte, err error) {
	type shadow SolidColorOverlayParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type SolidColorOverlayTransformationParam struct {
	// Specifies the transparency level of the solid color overlay. Accepts integers
	// from `1` to `9`.
	Alpha param.Opt[float64] `json:"alpha,omitzero"`
	// Specifies the background color of the solid color overlay. Accepts an RGB hex
	// code (e.g., `FF0000`), an RGBA code (e.g., `FFAABB50`), or a color name.
	Background param.Opt[string] `json:"background,omitzero"`
	// Creates a linear gradient with two colors. Pass `true` for a default gradient,
	// or provide a string for a custom gradient. Only works if the base asset is an
	// image. See
	// [gradient](https://imagekit.io/docs/effects-and-enhancements#gradient---e-gradient).
	Gradient SolidColorOverlayTransformationGradientUnionParam `json:"gradient,omitzero"`
	// Controls the height of the solid color overlay. Accepts a numeric value or an
	// arithmetic expression. Learn about
	// [arithmetic expressions](https://imagekit.io/docs/arithmetic-expressions-in-transformations).
	Height SolidColorOverlayTransformationHeightUnionParam `json:"height,omitzero"`
	// Specifies the corner radius of the solid color overlay. Set to `max` for
	// circular or oval shape. See
	// [radius](https://imagekit.io/docs/effects-and-enhancements#radius---r).
	Radius SolidColorOverlayTransformationRadiusUnionParam `json:"radius,omitzero"`
	// Controls the width of the solid color overlay. Accepts a numeric value or an
	// arithmetic expression (e.g., `bw_mul_0.2` or `bh_div_2`). Learn about
	// [arithmetic expressions](https://imagekit.io/docs/arithmetic-expressions-in-transformations).
	Width SolidColorOverlayTransformationWidthUnionParam `json:"width,omitzero"`
	paramObj
}

func (r SolidColorOverlayTransformationParam) MarshalJSON() (data []byte, err error) {
	type shadow SolidColorOverlayTransformationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SolidColorOverlayTransformationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SolidColorOverlayTransformationGradientUnionParam struct {
	// Construct this variant with constant.ValueOf[bool]()
	OfSolidColorOverlayTransformationGradientBoolean param.Opt[bool]   `json:",omitzero,inline"`
	OfString                                         param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u SolidColorOverlayTransformationGradientUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSolidColorOverlayTransformationGradientBoolean, u.OfString)
}
func (u *SolidColorOverlayTransformationGradientUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SolidColorOverlayTransformationGradientUnionParam) asAny() any {
	if !param.IsOmitted(u.OfSolidColorOverlayTransformationGradientBoolean) {
		return &u.OfSolidColorOverlayTransformationGradientBoolean
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

type SolidColorOverlayTransformationGradientBoolean bool

const (
	SolidColorOverlayTransformationGradientBooleanTrue SolidColorOverlayTransformationGradientBoolean = true
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SolidColorOverlayTransformationHeightUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u SolidColorOverlayTransformationHeightUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *SolidColorOverlayTransformationHeightUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SolidColorOverlayTransformationHeightUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SolidColorOverlayTransformationRadiusUnionParam struct {
	OfFloat param.Opt[float64] `json:",omitzero,inline"`
	// Construct this variant with constant.ValueOf[constant.Max]()
	OfMax constant.Max `json:",omitzero,inline"`
	paramUnion
}

func (u SolidColorOverlayTransformationRadiusUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfMax)
}
func (u *SolidColorOverlayTransformationRadiusUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SolidColorOverlayTransformationRadiusUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfMax) {
		return &u.OfMax
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SolidColorOverlayTransformationWidthUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u SolidColorOverlayTransformationWidthUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *SolidColorOverlayTransformationWidthUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SolidColorOverlayTransformationWidthUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Options for generating ImageKit URLs with transformations. See the
// [Transformations guide](https://imagekit.io/docs/transformations).
//
// The properties Src, URLEndpoint are required.
type SrcOptionsParam struct {
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
	ExpiresIn param.Opt[float64] `json:"expiresIn,omitzero"`
	// Whether to sign the URL or not. Set this to `true` if you want to generate a
	// signed URL. If `signed` is `true` and `expiresIn` is not specified, the signed
	// URL will not expire (valid indefinitely). Note: If `expiresIn` is set to any
	// value above 0, the URL will always be signed regardless of this setting.
	// [Learn more](https://imagekit.io/docs/media-delivery-basic-security#how-to-generate-signed-urls).
	Signed param.Opt[bool] `json:"signed,omitzero"`
	// These are additional query parameters that you want to add to the final URL.
	// They can be any query parameters and not necessarily related to ImageKit. This
	// is especially useful if you want to add a versioning parameter to your URLs.
	QueryParameters map[string]string `json:"queryParameters,omitzero"`
	// An array of objects specifying the transformations to be applied in the URL. If
	// more than one transformation is specified, they are applied in the order they
	// are specified as chained transformations. See
	// [Chained transformations](https://imagekit.io/docs/transformations#chained-transformations).
	Transformation []TransformationParam `json:"transformation,omitzero"`
	// By default, the transformation string is added as a query parameter in the URL,
	// e.g., `?tr=w-100,h-100`. If you want to add the transformation string in the
	// path of the URL, set this to `path`. Learn more in the
	// [Transformations guide](https://imagekit.io/docs/transformations).
	//
	// Any of "path", "query".
	TransformationPosition TransformationPosition `json:"transformationPosition,omitzero"`
	paramObj
}

func (r SrcOptionsParam) MarshalJSON() (data []byte, err error) {
	type shadow SrcOptionsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SrcOptionsParam) UnmarshalJSON(data []byte) error {
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

type SubtitleOverlayParam struct {
	// Specifies the relative path to the subtitle file used as an overlay.
	Input string            `json:"input,required"`
	Type  constant.Subtitle `json:"type,required"`
	// The input path can be included in the layer as either `i-{input}` or
	// `ie-{base64_encoded_input}`. By default, the SDK determines the appropriate
	// format automatically. To always use base64 encoding (`ie-{base64}`), set this
	// parameter to `base64`. To always use plain text (`i-{input}`), set it to
	// `plain`.
	Encoding string `json:"encoding,omitzero"`
	// Control styling of the subtitle. See
	// [Styling subtitles](https://imagekit.io/docs/add-overlays-on-videos#styling-controls-for-subtitles-layer).
	Transformation []SubtitleOverlayTransformationParam `json:"transformation,omitzero"`
	BaseOverlayParam
}

func (r SubtitleOverlayParam) MarshalJSON() (data []byte, err error) {
	type shadow SubtitleOverlayParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// Subtitle styling options.
// [Learn more](https://imagekit.io/docs/add-overlays-on-videos#styling-controls-for-subtitles-layer)
// from the docs.
type SubtitleOverlayTransformationParam struct {
	// Specifies the subtitle background color using a standard color name, an RGB
	// color code (e.g., FF0000), or an RGBA color code (e.g., FFAABB50).
	//
	// [Subtitle styling options](https://imagekit.io/docs/add-overlays-on-videos#styling-controls-for-subtitles-layer)
	Background param.Opt[string] `json:"background,omitzero"`
	// Sets the font color of the subtitle text using a standard color name, an RGB
	// color code (e.g., FF0000), or an RGBA color code (e.g., FFAABB50).
	//
	// [Subtitle styling options](https://imagekit.io/docs/add-overlays-on-videos#styling-controls-for-subtitles-layer)
	Color param.Opt[string] `json:"color,omitzero"`
	// Font family for subtitles. Refer to the
	// [supported fonts](https://imagekit.io/docs/add-overlays-on-images#supported-text-font-list).
	FontFamily param.Opt[string] `json:"fontFamily,omitzero"`
	// Sets the font outline of the subtitle text. Requires the outline width (an
	// integer) and the outline color (as an RGB color code, RGBA color code, or
	// standard web color name) separated by an underscore. Example: `fol-2_blue`
	// (outline width of 2px and outline color blue), `fol-2_A1CCDD` (outline width of
	// 2px and outline color `#A1CCDD`) and `fol-2_A1CCDD50` (outline width of 2px and
	// outline color `#A1CCDD` at 50% opacity).
	//
	// [Subtitle styling options](https://imagekit.io/docs/add-overlays-on-videos#styling-controls-for-subtitles-layer)
	FontOutline param.Opt[string] `json:"fontOutline,omitzero"`
	// Sets the font shadow for the subtitle text. Requires the shadow color (as an RGB
	// color code, RGBA color code, or standard web color name) and shadow indent (an
	// integer) separated by an underscore. Example: `fsh-blue_2` (shadow color blue,
	// indent of 2px), `fsh-A1CCDD_3` (shadow color `#A1CCDD`, indent of 3px),
	// `fsh-A1CCDD50_3` (shadow color `#A1CCDD` at 50% opacity, indent of 3px).
	//
	// [Subtitle styling options](https://imagekit.io/docs/add-overlays-on-videos#styling-controls-for-subtitles-layer)
	FontShadow param.Opt[string] `json:"fontShadow,omitzero"`
	// Sets the font size of subtitle text.
	//
	// [Subtitle styling options](https://imagekit.io/docs/add-overlays-on-videos#styling-controls-for-subtitles-layer)
	FontSize param.Opt[float64] `json:"fontSize,omitzero"`
	// Sets the typography style of the subtitle text. Supports values are `b` for
	// bold, `i` for italics, and `b_i` for bold with italics.
	//
	// [Subtitle styling options](https://imagekit.io/docs/add-overlays-on-videos#styling-controls-for-subtitles-layer)
	//
	// Any of "b", "i", "b_i".
	Typography SubtitleOverlayTransformationTypography `json:"typography,omitzero"`
	paramObj
}

func (r SubtitleOverlayTransformationParam) MarshalJSON() (data []byte, err error) {
	type shadow SubtitleOverlayTransformationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SubtitleOverlayTransformationParam) UnmarshalJSON(data []byte) error {
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

type TextOverlayParam struct {
	// Specifies the text to be displayed in the overlay. The SDK automatically handles
	// special characters and encoding.
	Text string        `json:"text,required"`
	Type constant.Text `json:"type,required"`
	// Text can be included in the layer as either `i-{input}` (plain text) or
	// `ie-{base64_encoded_input}` (base64). By default, the SDK selects the
	// appropriate format based on the input text. To always use base64
	// (`ie-{base64}`), set this parameter to `base64`. To always use plain text
	// (`i-{input}`), set it to `plain`.
	Encoding string `json:"encoding,omitzero"`
	// Control styling of the text overlay. See
	// [Text overlays](https://imagekit.io/docs/add-overlays-on-images#text-overlay).
	Transformation []TextOverlayTransformationParam `json:"transformation,omitzero"`
	BaseOverlayParam
}

func (r TextOverlayParam) MarshalJSON() (data []byte, err error) {
	type shadow TextOverlayParam
	return param.MarshalObject(r, (*shadow)(&r))
}

type TextOverlayTransformationParam struct {
	// Specifies the transparency level of the text overlay. Accepts integers from `1`
	// to `9`.
	Alpha param.Opt[float64] `json:"alpha,omitzero"`
	// Specifies the background color of the text overlay. Accepts an RGB hex code, an
	// RGBA code, or a color name.
	Background param.Opt[string] `json:"background,omitzero"`
	// Specifies the font color of the overlaid text. Accepts an RGB hex code (e.g.,
	// `FF0000`), an RGBA code (e.g., `FFAABB50`), or a color name.
	FontColor param.Opt[string] `json:"fontColor,omitzero"`
	// Specifies the font family of the overlaid text. Choose from the supported fonts
	// list or use a custom font. See
	// [Supported fonts](https://imagekit.io/docs/add-overlays-on-images#supported-text-font-list)
	// and
	// [Custom font](https://imagekit.io/docs/add-overlays-on-images#change-font-family-in-text-overlay).
	FontFamily param.Opt[string] `json:"fontFamily,omitzero"`
	// Specifies the typography style of the text. Supported values:
	//
	//   - Single styles: `b` (bold), `i` (italic), `strikethrough`.
	//   - Combinations: Any combination separated by underscores, e.g., `b_i`,
	//     `b_i_strikethrough`.
	Typography param.Opt[string] `json:"typography,omitzero"`
	// Flip the text overlay horizontally, vertically, or both.
	//
	// Any of "h", "v", "h_v", "v_h".
	Flip TextOverlayTransformationFlip `json:"flip,omitzero"`
	// Specifies the font size of the overlaid text. Accepts a numeric value or an
	// arithmetic expression.
	FontSize TextOverlayTransformationFontSizeUnionParam `json:"fontSize,omitzero"`
	// Specifies the inner alignment of the text when width is more than the text
	// length.
	//
	// Any of "left", "right", "center".
	InnerAlignment TextOverlayTransformationInnerAlignment `json:"innerAlignment,omitzero"`
	// Specifies the line height of the text overlay. Accepts integer values
	// representing line height in points. It can also accept
	// [arithmetic expressions](https://imagekit.io/docs/arithmetic-expressions-in-transformations)
	// such as `bw_mul_0.2`, or `bh_div_20`.
	LineHeight TextOverlayTransformationLineHeightUnionParam `json:"lineHeight,omitzero"`
	// Specifies the padding around the overlaid text. Can be provided as a single
	// positive integer or multiple values separated by underscores (following CSS
	// shorthand order). Arithmetic expressions are also accepted.
	Padding TextOverlayTransformationPaddingUnionParam `json:"padding,omitzero"`
	// Specifies the corner radius of the text overlay. Set to `max` to achieve a
	// circular or oval shape.
	Radius TextOverlayTransformationRadiusUnionParam `json:"radius,omitzero"`
	// Specifies the rotation angle of the text overlay. Accepts a numeric value for
	// clockwise rotation or a string prefixed with "N" for counter-clockwise rotation.
	Rotation TextOverlayTransformationRotationUnionParam `json:"rotation,omitzero"`
	// Specifies the maximum width (in pixels) of the overlaid text. The text wraps
	// automatically, and arithmetic expressions (e.g., `bw_mul_0.2` or `bh_div_2`) are
	// supported. Useful when used in conjunction with the `background`. Learn about
	// [Arithmetic expressions](https://imagekit.io/docs/arithmetic-expressions-in-transformations).
	Width TextOverlayTransformationWidthUnionParam `json:"width,omitzero"`
	paramObj
}

func (r TextOverlayTransformationParam) MarshalJSON() (data []byte, err error) {
	type shadow TextOverlayTransformationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextOverlayTransformationParam) UnmarshalJSON(data []byte) error {
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TextOverlayTransformationFontSizeUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u TextOverlayTransformationFontSizeUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *TextOverlayTransformationFontSizeUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TextOverlayTransformationFontSizeUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Specifies the inner alignment of the text when width is more than the text
// length.
type TextOverlayTransformationInnerAlignment string

const (
	TextOverlayTransformationInnerAlignmentLeft   TextOverlayTransformationInnerAlignment = "left"
	TextOverlayTransformationInnerAlignmentRight  TextOverlayTransformationInnerAlignment = "right"
	TextOverlayTransformationInnerAlignmentCenter TextOverlayTransformationInnerAlignment = "center"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TextOverlayTransformationLineHeightUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u TextOverlayTransformationLineHeightUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *TextOverlayTransformationLineHeightUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TextOverlayTransformationLineHeightUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TextOverlayTransformationPaddingUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u TextOverlayTransformationPaddingUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *TextOverlayTransformationPaddingUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TextOverlayTransformationPaddingUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TextOverlayTransformationRadiusUnionParam struct {
	OfFloat param.Opt[float64] `json:",omitzero,inline"`
	// Construct this variant with constant.ValueOf[constant.Max]()
	OfMax constant.Max `json:",omitzero,inline"`
	paramUnion
}

func (u TextOverlayTransformationRadiusUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfMax)
}
func (u *TextOverlayTransformationRadiusUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TextOverlayTransformationRadiusUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfMax) {
		return &u.OfMax
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TextOverlayTransformationRotationUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u TextOverlayTransformationRotationUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *TextOverlayTransformationRotationUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TextOverlayTransformationRotationUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TextOverlayTransformationWidthUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u TextOverlayTransformationWidthUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *TextOverlayTransformationWidthUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TextOverlayTransformationWidthUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// The SDK provides easy-to-use names for transformations. These names are
// converted to the corresponding transformation string before being added to the
// URL. SDKs are updated regularly to support new transformations. If you want to
// use a transformation that is not supported by the SDK, You can use the `raw`
// parameter to pass the transformation string directly. See the
// [Transformations documentation](https://imagekit.io/docs/transformations).
type TransformationParam struct {
	// Uses AI to change the background. Provide a text prompt or a base64-encoded
	// prompt, e.g., `prompt-snow road` or `prompte-[urlencoded_base64_encoded_text]`.
	// Not supported inside overlay. See
	// [AI Change Background](https://imagekit.io/docs/ai-transformations#change-background-e-changebg).
	AIChangeBackground param.Opt[string] `json:"aiChangeBackground,omitzero"`
	// Uses AI to edit images based on a text prompt. Provide a text prompt or a
	// base64-encoded prompt, e.g., `prompt-snow road` or
	// `prompte-[urlencoded_base64_encoded_text]`. Not supported inside overlay.
	// See [AI Edit](https://imagekit.io/docs/ai-transformations#edit-image-e-edit).
	AIEdit param.Opt[string] `json:"aiEdit,omitzero"`
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
	Background param.Opt[string] `json:"background,omitzero"`
	// Specifies the Gaussian blur level. Accepts an integer value between 1 and 100,
	// or an expression like `bl-10`. See
	// [Blur](https://imagekit.io/docs/effects-and-enhancements#blur---bl).
	Blur param.Opt[float64] `json:"blur,omitzero"`
	// Adds a border to the output media. Accepts a string in the format
	// `<border-width>_<hex-code>` (e.g., `5_FFF000` for a 5px yellow border), or an
	// expression like `ih_div_20_FF00FF`. See
	// [Border](https://imagekit.io/docs/effects-and-enhancements#border---b).
	Border param.Opt[string] `json:"border,omitzero"`
	// Indicates whether the output image should retain the original color profile. See
	// [Color profile](https://imagekit.io/docs/image-optimization#color-profile---cp).
	ColorProfile param.Opt[bool] `json:"colorProfile,omitzero"`
	// Specifies a fallback image if the resource is not found, e.g., a URL or file
	// path. See
	// [Default image](https://imagekit.io/docs/image-transformation#default-image---di).
	DefaultImage param.Opt[string] `json:"defaultImage,omitzero"`
	// Accepts values between 0.1 and 5, or `auto` for automatic device pixel ratio
	// (DPR) calculation. See
	// [DPR](https://imagekit.io/docs/image-resize-and-crop#dpr---dpr).
	Dpr param.Opt[float64] `json:"dpr,omitzero"`
	// Refines padding and cropping behavior for pad resize, maintain ratio, and
	// extract crop modes. Supports manual positions and coordinate-based focus. With
	// AI-based cropping, you can automatically keep key subjects in frame—such as
	// faces or detected objects (e.g., `fo-face`, `fo-person`, `fo-car`)— while
	// resizing.
	//
	// - See [Focus](https://imagekit.io/docs/image-resize-and-crop#focus---fo).
	// - [Object aware cropping](https://imagekit.io/docs/image-resize-and-crop#object-aware-cropping---fo-object-name)
	Focus param.Opt[string] `json:"focus,omitzero"`
	// Specifies whether the output image (in JPEG or PNG) should be compressed
	// losslessly. See
	// [Lossless compression](https://imagekit.io/docs/image-optimization#lossless-webp-and-png---lo).
	Lossless param.Opt[bool] `json:"lossless,omitzero"`
	// By default, ImageKit removes all metadata during automatic image compression.
	// Set this to true to preserve metadata. See
	// [Image metadata](https://imagekit.io/docs/image-optimization#image-metadata---md).
	Metadata param.Opt[bool] `json:"metadata,omitzero"`
	// Named transformation reference. See
	// [Named transformations](https://imagekit.io/docs/transformations#named-transformations).
	Named param.Opt[string] `json:"named,omitzero"`
	// Specifies the opacity level of the output image. See
	// [Opacity](https://imagekit.io/docs/effects-and-enhancements#opacity---o).
	Opacity param.Opt[float64] `json:"opacity,omitzero"`
	// If set to true, serves the original file without applying any transformations.
	// See
	// [Deliver original file as-is](https://imagekit.io/docs/core-delivery-features#deliver-original-file-as-is---orig-true).
	Original param.Opt[bool] `json:"original,omitzero"`
	// Specifies whether the output JPEG image should be rendered progressively.
	// Progressive loading begins with a low-quality, pixelated version of the full
	// image, which gradually improves to provide a faster perceived load time. See
	// [Progressive images](https://imagekit.io/docs/image-optimization#progressive-image---pr).
	Progressive param.Opt[bool] `json:"progressive,omitzero"`
	// Specifies the quality of the output image for lossy formats such as JPEG, WebP,
	// and AVIF. A higher quality value results in a larger file size with better
	// quality, while a lower value produces a smaller file size with reduced quality.
	// See [Quality](https://imagekit.io/docs/image-optimization#quality---q).
	Quality param.Opt[float64] `json:"quality,omitzero"`
	// Pass any transformation not directly supported by the SDK. This transformation
	// string is appended to the URL as provided.
	Raw param.Opt[string] `json:"raw,omitzero"`
	// Accepts a numeric value that determines how much to zoom in or out of the
	// cropped area. It should be used in conjunction with fo-face or fo-<object_name>.
	// See [Zoom](https://imagekit.io/docs/image-resize-and-crop#zoom---z).
	Zoom param.Opt[float64] `json:"zoom,omitzero"`
	// Adds an AI-based drop shadow around a foreground object on a transparent or
	// removed background. Optionally, control the direction, elevation, and saturation
	// of the light source (e.g., `az-45` to change light direction). Pass `true` for
	// the default drop shadow, or provide a string for a custom drop shadow. Supported
	// inside overlay. See
	// [AI Drop Shadow](https://imagekit.io/docs/ai-transformations#ai-drop-shadow-e-dropshadow).
	AIDropShadow TransformationAIDropShadowUnionParam `json:"aiDropShadow,omitzero"`
	// Applies ImageKit's in-house background removal. Supported inside overlay. See
	// [AI Background Removal](https://imagekit.io/docs/ai-transformations#imagekit-background-removal-e-bgremove).
	//
	// Any of true.
	AIRemoveBackground bool `json:"aiRemoveBackground,omitzero"`
	// Uses third-party background removal. Note: It is recommended to use
	// aiRemoveBackground, ImageKit's in-house solution, which is more cost-effective.
	// Supported inside overlay. See
	// [External Background Removal](https://imagekit.io/docs/ai-transformations#background-removal-e-removedotbg).
	//
	// Any of true.
	AIRemoveBackgroundExternal bool `json:"aiRemoveBackgroundExternal,omitzero"`
	// Performs AI-based retouching to improve faces or product shots. Not supported
	// inside overlay. See
	// [AI Retouch](https://imagekit.io/docs/ai-transformations#retouch-e-retouch).
	//
	// Any of true.
	AIRetouch bool `json:"aiRetouch,omitzero"`
	// Upscales images beyond their original dimensions using AI. Not supported inside
	// overlay. See
	// [AI Upscale](https://imagekit.io/docs/ai-transformations#upscale-e-upscale).
	//
	// Any of true.
	AIUpscale bool `json:"aiUpscale,omitzero"`
	// Generates a variation of an image using AI. This produces a new image with
	// slight variations from the original, such as changes in color, texture, and
	// other visual elements, while preserving the structure and essence of the
	// original image. Not supported inside overlay. See
	// [AI Generate Variations](https://imagekit.io/docs/ai-transformations#generate-variations-of-an-image-e-genvar).
	//
	// Any of true.
	AIVariation bool `json:"aiVariation,omitzero"`
	// Specifies the aspect ratio for the output, e.g., "ar-4-3". Typically used with
	// either width or height (but not both). For example: aspectRatio = `4:3`, `4_3`,
	// or an expression like `iar_div_2`. See
	// [Image resize and crop – Aspect ratio](https://imagekit.io/docs/image-resize-and-crop#aspect-ratio---ar).
	AspectRatio TransformationAspectRatioUnionParam `json:"aspectRatio,omitzero"`
	// Specifies the audio codec, e.g., `aac`, `opus`, or `none`. See
	// [Audio codec](https://imagekit.io/docs/video-optimization#audio-codec---ac).
	//
	// Any of "aac", "opus", "none".
	AudioCodec TransformationAudioCodec `json:"audioCodec,omitzero"`
	// Automatically enhances the contrast of an image (contrast stretch). See
	// [Contrast Stretch](https://imagekit.io/docs/effects-and-enhancements#contrast-stretch---e-contrast).
	//
	// Any of true.
	ContrastStretch bool `json:"contrastStretch,omitzero"`
	// Crop modes for image resizing. See
	// [Crop modes & focus](https://imagekit.io/docs/image-resize-and-crop#crop-crop-modes--focus).
	//
	// Any of "force", "at_max", "at_max_enlarge", "at_least", "maintain_ratio".
	Crop TransformationCrop `json:"crop,omitzero"`
	// Additional crop modes for image resizing. See
	// [Crop modes & focus](https://imagekit.io/docs/image-resize-and-crop#crop-crop-modes--focus).
	//
	// Any of "pad_resize", "extract", "pad_extract".
	CropMode TransformationCropMode `json:"cropMode,omitzero"`
	// Specifies the duration (in seconds) for trimming videos, e.g., `5` or `10.5`.
	// Typically used with startOffset to indicate the length from the start offset.
	// Arithmetic expressions are supported. See
	// [Trim videos – Duration](https://imagekit.io/docs/trim-videos#duration---du).
	Duration TransformationDurationUnionParam `json:"duration,omitzero"`
	// Specifies the end offset (in seconds) for trimming videos, e.g., `5` or `10.5`.
	// Typically used with startOffset to define a time window. Arithmetic expressions
	// are supported. See
	// [Trim videos – End offset](https://imagekit.io/docs/trim-videos#end-offset---eo).
	EndOffset TransformationEndOffsetUnionParam `json:"endOffset,omitzero"`
	// Flips or mirrors an image either horizontally, vertically, or both. Acceptable
	// values: `h` (horizontal), `v` (vertical), `h_v` (horizontal and vertical), or
	// `v_h`. See [Flip](https://imagekit.io/docs/effects-and-enhancements#flip---fl).
	//
	// Any of "h", "v", "h_v", "v_h".
	Flip TransformationFlip `json:"flip,omitzero"`
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
	Format TransformationFormat `json:"format,omitzero"`
	// Creates a linear gradient with two colors. Pass `true` for a default gradient,
	// or provide a string for a custom gradient. See
	// [Gradient](https://imagekit.io/docs/effects-and-enhancements#gradient---e-gradient).
	Gradient TransformationGradientUnionParam `json:"gradient,omitzero"`
	// Enables a grayscale effect for images. See
	// [Grayscale](https://imagekit.io/docs/effects-and-enhancements#grayscale---e-grayscale).
	//
	// Any of true.
	Grayscale bool `json:"grayscale,omitzero"`
	// Specifies the height of the output. If a value between 0 and 1 is provided, it
	// is treated as a percentage (e.g., `0.5` represents 50% of the original height).
	// You can also supply arithmetic expressions (e.g., `ih_mul_0.5`). Height
	// transformation –
	// [Images](https://imagekit.io/docs/image-resize-and-crop#height---h) ·
	// [Videos](https://imagekit.io/docs/video-resize-and-crop#height---h)
	Height TransformationHeightUnionParam `json:"height,omitzero"`
	// Specifies an overlay to be applied on the parent image or video. ImageKit
	// supports overlays including images, text, videos, subtitles, and solid colors.
	// See
	// [Overlay using layers](https://imagekit.io/docs/transformations#overlay-using-layers).
	Overlay OverlayUnionParam `json:"overlay,omitzero"`
	// Extracts a specific page or frame from multi-page or layered files (PDF, PSD,
	// AI). For example, specify by number (e.g., `2`), a range (e.g., `3-4` for the
	// 2nd and 3rd layers), or by name (e.g., `name-layer-4` for a PSD layer). See
	// [Thumbnail extraction](https://imagekit.io/docs/vector-and-animated-images#get-thumbnail-from-psd-pdf-ai-eps-and-animated-files).
	Page TransformationPageUnionParam `json:"page,omitzero"`
	// Specifies the corner radius for rounded corners (e.g., 20) or `max` for circular
	// or oval shape. See
	// [Radius](https://imagekit.io/docs/effects-and-enhancements#radius---r).
	Radius TransformationRadiusUnionParam `json:"radius,omitzero"`
	// Specifies the rotation angle in degrees. Positive values rotate the image
	// clockwise; you can also use, for example, `N40` for counterclockwise rotation or
	// `auto` to use the orientation specified in the image's EXIF data. For videos,
	// only the following values are supported: 0, 90, 180, 270, or 360. See
	// [Rotate](https://imagekit.io/docs/effects-and-enhancements#rotate---rt).
	Rotation TransformationRotationUnionParam `json:"rotation,omitzero"`
	// Adds a shadow beneath solid objects in an image with a transparent background.
	// For AI-based drop shadows, refer to aiDropShadow. Pass `true` for a default
	// shadow, or provide a string for a custom shadow. See
	// [Shadow](https://imagekit.io/docs/effects-and-enhancements#shadow---e-shadow).
	Shadow TransformationShadowUnionParam `json:"shadow,omitzero"`
	// Sharpens the input image, highlighting edges and finer details. Pass `true` for
	// default sharpening, or provide a numeric value for custom sharpening. See
	// [Sharpen](https://imagekit.io/docs/effects-and-enhancements#sharpen---e-sharpen).
	Sharpen TransformationSharpenUnionParam `json:"sharpen,omitzero"`
	// Specifies the start offset (in seconds) for trimming videos, e.g., `5` or
	// `10.5`. Arithmetic expressions are also supported. See
	// [Trim videos – Start offset](https://imagekit.io/docs/trim-videos#start-offset---so).
	StartOffset TransformationStartOffsetUnionParam `json:"startOffset,omitzero"`
	// An array of resolutions for adaptive bitrate streaming, e.g., [`240`, `360`,
	// `480`, `720`, `1080`]. See
	// [Adaptive Bitrate Streaming](https://imagekit.io/docs/adaptive-bitrate-streaming).
	StreamingResolutions []StreamingResolution `json:"streamingResolutions,omitzero"`
	// Useful for images with a solid or nearly solid background and a central object.
	// This parameter trims the background, leaving only the central object in the
	// output image. See
	// [Trim edges](https://imagekit.io/docs/effects-and-enhancements#trim-edges---t).
	Trim TransformationTrimUnionParam `json:"trim,omitzero"`
	// Applies Unsharp Masking (USM), an image sharpening technique. Pass `true` for a
	// default unsharp mask, or provide a string for a custom unsharp mask. See
	// [Unsharp Mask](https://imagekit.io/docs/effects-and-enhancements#unsharp-mask---e-usm).
	UnsharpMask TransformationUnsharpMaskUnionParam `json:"unsharpMask,omitzero"`
	// Specifies the video codec, e.g., `h264`, `vp9`, `av1`, or `none`. See
	// [Video codec](https://imagekit.io/docs/video-optimization#video-codec---vc).
	//
	// Any of "h264", "vp9", "av1", "none".
	VideoCodec TransformationVideoCodec `json:"videoCodec,omitzero"`
	// Specifies the width of the output. If a value between 0 and 1 is provided, it is
	// treated as a percentage (e.g., `0.4` represents 40% of the original width). You
	// can also supply arithmetic expressions (e.g., `iw_div_2`). Width transformation
	// – [Images](https://imagekit.io/docs/image-resize-and-crop#width---w) ·
	// [Videos](https://imagekit.io/docs/video-resize-and-crop#width---w)
	Width TransformationWidthUnionParam `json:"width,omitzero"`
	// Focus using cropped image coordinates - X coordinate. See
	// [Focus using cropped coordinates](https://imagekit.io/docs/image-resize-and-crop#example---focus-using-cropped-image-coordinates).
	X TransformationXUnionParam `json:"x,omitzero"`
	// Focus using cropped image coordinates - X center coordinate. See
	// [Focus using cropped coordinates](https://imagekit.io/docs/image-resize-and-crop#example---focus-using-cropped-image-coordinates).
	XCenter TransformationXCenterUnionParam `json:"xCenter,omitzero"`
	// Focus using cropped image coordinates - Y coordinate. See
	// [Focus using cropped coordinates](https://imagekit.io/docs/image-resize-and-crop#example---focus-using-cropped-image-coordinates).
	Y TransformationYUnionParam `json:"y,omitzero"`
	// Focus using cropped image coordinates - Y center coordinate. See
	// [Focus using cropped coordinates](https://imagekit.io/docs/image-resize-and-crop#example---focus-using-cropped-image-coordinates).
	YCenter TransformationYCenterUnionParam `json:"yCenter,omitzero"`
	paramObj
}

func (r TransformationParam) MarshalJSON() (data []byte, err error) {
	type shadow TransformationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransformationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TransformationParam](
		"aiRemoveBackground", true,
	)
	apijson.RegisterFieldValidator[TransformationParam](
		"aiRemoveBackgroundExternal", true,
	)
	apijson.RegisterFieldValidator[TransformationParam](
		"aiRetouch", true,
	)
	apijson.RegisterFieldValidator[TransformationParam](
		"aiUpscale", true,
	)
	apijson.RegisterFieldValidator[TransformationParam](
		"aiVariation", true,
	)
	apijson.RegisterFieldValidator[TransformationParam](
		"contrastStretch", true,
	)
	apijson.RegisterFieldValidator[TransformationParam](
		"grayscale", true,
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TransformationAIDropShadowUnionParam struct {
	// Construct this variant with constant.ValueOf[bool]()
	OfTransformationAIDropShadowBoolean param.Opt[bool]   `json:",omitzero,inline"`
	OfString                            param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u TransformationAIDropShadowUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfTransformationAIDropShadowBoolean, u.OfString)
}
func (u *TransformationAIDropShadowUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TransformationAIDropShadowUnionParam) asAny() any {
	if !param.IsOmitted(u.OfTransformationAIDropShadowBoolean) {
		return &u.OfTransformationAIDropShadowBoolean
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

type TransformationAIDropShadowBoolean bool

const (
	TransformationAIDropShadowBooleanTrue TransformationAIDropShadowBoolean = true
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TransformationAspectRatioUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u TransformationAspectRatioUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *TransformationAspectRatioUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TransformationAspectRatioUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TransformationDurationUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u TransformationDurationUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *TransformationDurationUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TransformationDurationUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TransformationEndOffsetUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u TransformationEndOffsetUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *TransformationEndOffsetUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TransformationEndOffsetUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TransformationGradientUnionParam struct {
	// Construct this variant with constant.ValueOf[bool]()
	OfTransformationGradientBoolean param.Opt[bool]   `json:",omitzero,inline"`
	OfString                        param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u TransformationGradientUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfTransformationGradientBoolean, u.OfString)
}
func (u *TransformationGradientUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TransformationGradientUnionParam) asAny() any {
	if !param.IsOmitted(u.OfTransformationGradientBoolean) {
		return &u.OfTransformationGradientBoolean
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

type TransformationGradientBoolean bool

const (
	TransformationGradientBooleanTrue TransformationGradientBoolean = true
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TransformationHeightUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u TransformationHeightUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *TransformationHeightUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TransformationHeightUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TransformationPageUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u TransformationPageUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *TransformationPageUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TransformationPageUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TransformationRadiusUnionParam struct {
	OfFloat param.Opt[float64] `json:",omitzero,inline"`
	// Construct this variant with constant.ValueOf[constant.Max]()
	OfMax constant.Max `json:",omitzero,inline"`
	paramUnion
}

func (u TransformationRadiusUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfMax)
}
func (u *TransformationRadiusUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TransformationRadiusUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfMax) {
		return &u.OfMax
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TransformationRotationUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u TransformationRotationUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *TransformationRotationUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TransformationRotationUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TransformationShadowUnionParam struct {
	// Construct this variant with constant.ValueOf[bool]()
	OfTransformationShadowBoolean param.Opt[bool]   `json:",omitzero,inline"`
	OfString                      param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u TransformationShadowUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfTransformationShadowBoolean, u.OfString)
}
func (u *TransformationShadowUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TransformationShadowUnionParam) asAny() any {
	if !param.IsOmitted(u.OfTransformationShadowBoolean) {
		return &u.OfTransformationShadowBoolean
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

type TransformationShadowBoolean bool

const (
	TransformationShadowBooleanTrue TransformationShadowBoolean = true
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TransformationSharpenUnionParam struct {
	// Construct this variant with constant.ValueOf[bool]()
	OfTransformationSharpenBoolean param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat                        param.Opt[float64] `json:",omitzero,inline"`
	paramUnion
}

func (u TransformationSharpenUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfTransformationSharpenBoolean, u.OfFloat)
}
func (u *TransformationSharpenUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TransformationSharpenUnionParam) asAny() any {
	if !param.IsOmitted(u.OfTransformationSharpenBoolean) {
		return &u.OfTransformationSharpenBoolean
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	}
	return nil
}

type TransformationSharpenBoolean bool

const (
	TransformationSharpenBooleanTrue TransformationSharpenBoolean = true
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TransformationStartOffsetUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u TransformationStartOffsetUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *TransformationStartOffsetUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TransformationStartOffsetUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TransformationTrimUnionParam struct {
	// Construct this variant with constant.ValueOf[bool]()
	OfTransformationTrimBoolean param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat                     param.Opt[float64] `json:",omitzero,inline"`
	paramUnion
}

func (u TransformationTrimUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfTransformationTrimBoolean, u.OfFloat)
}
func (u *TransformationTrimUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TransformationTrimUnionParam) asAny() any {
	if !param.IsOmitted(u.OfTransformationTrimBoolean) {
		return &u.OfTransformationTrimBoolean
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	}
	return nil
}

type TransformationTrimBoolean bool

const (
	TransformationTrimBooleanTrue TransformationTrimBoolean = true
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TransformationUnsharpMaskUnionParam struct {
	// Construct this variant with constant.ValueOf[bool]()
	OfTransformationUnsharpMaskBoolean param.Opt[bool]   `json:",omitzero,inline"`
	OfString                           param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u TransformationUnsharpMaskUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfTransformationUnsharpMaskBoolean, u.OfString)
}
func (u *TransformationUnsharpMaskUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TransformationUnsharpMaskUnionParam) asAny() any {
	if !param.IsOmitted(u.OfTransformationUnsharpMaskBoolean) {
		return &u.OfTransformationUnsharpMaskBoolean
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TransformationWidthUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u TransformationWidthUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *TransformationWidthUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TransformationWidthUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TransformationXUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u TransformationXUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *TransformationXUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TransformationXUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TransformationXCenterUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u TransformationXCenterUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *TransformationXCenterUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TransformationXCenterUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TransformationYUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u TransformationYUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *TransformationYUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TransformationYUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TransformationYCenterUnionParam struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u TransformationYCenterUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *TransformationYCenterUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TransformationYCenterUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
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

type VideoOverlayParam struct {
	// Specifies the relative path to the video used as an overlay.
	Input string         `json:"input,required"`
	Type  constant.Video `json:"type,required"`
	// The input path can be included in the layer as either `i-{input}` or
	// `ie-{base64_encoded_input}`. By default, the SDK determines the appropriate
	// format automatically. To always use base64 encoding (`ie-{base64}`), set this
	// parameter to `base64`. To always use plain text (`i-{input}`), set it to
	// `plain`.
	Encoding string `json:"encoding,omitzero"`
	// Array of transformation to be applied to the overlay video. Except
	// `streamingResolutions`, all other video transformations are supported. See
	// [Video transformations](https://imagekit.io/docs/video-transformation).
	Transformation []TransformationParam `json:"transformation,omitzero"`
	BaseOverlayParam
}

func (r VideoOverlayParam) MarshalJSON() (data []byte, err error) {
	type shadow VideoOverlayParam
	return param.MarshalObject(r, (*shadow)(&r))
}
