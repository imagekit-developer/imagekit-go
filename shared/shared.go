// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"
	"time"

	"github.com/imagekit-developer/imagekit-go/v2/internal/apijson"
	"github.com/imagekit-developer/imagekit-go/v2/packages/param"
	"github.com/imagekit-developer/imagekit-go/v2/packages/respjson"
	"github.com/imagekit-developer/imagekit-go/v2/shared/constant"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type BaseOverlayParam struct {
	// Controls how the layer blends with the base image or underlying content. Maps to
	// `lm` in the URL. By default, layers completely cover the base image beneath
	// them. Layer modes change this behavior:
	//
	//   - `multiply`: Multiplies the pixel values of the layer with the base image. The
	//     result is always darker than the original images. This is ideal for applying
	//     shadows or color tints.
	//   - `displace`: Uses the layer as a displacement map to distort pixels in the base
	//     image. The red channel controls horizontal displacement, and the green channel
	//     controls vertical displacement. Requires `x` or `y` parameter to control
	//     displacement magnitude.
	//   - `cutout`: Acts as an inverse mask where opaque areas of the layer turn the
	//     base image transparent, while transparent areas leave the base image
	//     unchanged. This mode functions like a hole-punch, effectively cutting the
	//     shape of the layer out of the underlying image.
	//   - `cutter`: Acts as a shape mask where only the parts of the base image that
	//     fall inside the opaque area of the layer are preserved. This mode functions
	//     like a cookie-cutter, trimming the base image to match the specific dimensions
	//     and shape of the layer. See
	//     [Layer modes](https://imagekit.io/docs/add-overlays-on-images#layer-modes).
	//
	// Any of "multiply", "cutter", "cutout", "displace".
	LayerMode BaseOverlayLayerMode `json:"layerMode,omitzero"`
	// Specifies the overlay's position relative to the parent asset. See
	// [Position of Layer](https://imagekit.io/docs/transformations#position-of-layer).
	Position OverlayPositionParam `json:"position,omitzero"`
	// Specifies timing information for the overlay (only applicable if the base asset
	// is a video). See
	// [Position of Layer](https://imagekit.io/docs/transformations#position-of-layer).
	Timing OverlayTimingParam `json:"timing,omitzero"`
	paramObj
}

func (r BaseOverlayParam) MarshalJSON() (data []byte, err error) {
	type shadow BaseOverlayParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BaseOverlayParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Controls how the layer blends with the base image or underlying content. Maps to
// `lm` in the URL. By default, layers completely cover the base image beneath
// them. Layer modes change this behavior:
//
//   - `multiply`: Multiplies the pixel values of the layer with the base image. The
//     result is always darker than the original images. This is ideal for applying
//     shadows or color tints.
//   - `displace`: Uses the layer as a displacement map to distort pixels in the base
//     image. The red channel controls horizontal displacement, and the green channel
//     controls vertical displacement. Requires `x` or `y` parameter to control
//     displacement magnitude.
//   - `cutout`: Acts as an inverse mask where opaque areas of the layer turn the
//     base image transparent, while transparent areas leave the base image
//     unchanged. This mode functions like a hole-punch, effectively cutting the
//     shape of the layer out of the underlying image.
//   - `cutter`: Acts as a shape mask where only the parts of the base image that
//     fall inside the opaque area of the layer are preserved. This mode functions
//     like a cookie-cutter, trimming the base image to match the specific dimensions
//     and shape of the layer. See
//     [Layer modes](https://imagekit.io/docs/add-overlays-on-images#layer-modes).
type BaseOverlayLayerMode string

const (
	BaseOverlayLayerModeMultiply BaseOverlayLayerMode = "multiply"
	BaseOverlayLayerModeCutter   BaseOverlayLayerMode = "cutter"
	BaseOverlayLayerModeCutout   BaseOverlayLayerMode = "cutout"
	BaseOverlayLayerModeDisplace BaseOverlayLayerMode = "displace"
)

// ExtensionConfigUnion contains all possible properties and values from
// [ExtensionConfigRemoveBg], [ExtensionConfigAutoTagging],
// [ExtensionConfigAIAutoDescription], [ExtensionConfigAITasks].
//
// Use the [ExtensionConfigUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ExtensionConfigUnion struct {
	// Any of "remove-bg", nil, "ai-auto-description", "ai-tasks".
	Name string `json:"name"`
	// This field is from variant [ExtensionConfigRemoveBg].
	Options ExtensionConfigRemoveBgOptions `json:"options"`
	// This field is from variant [ExtensionConfigAutoTagging].
	MaxTags int64 `json:"maxTags"`
	// This field is from variant [ExtensionConfigAutoTagging].
	MinConfidence int64 `json:"minConfidence"`
	// This field is from variant [ExtensionConfigAITasks].
	Tasks []ExtensionConfigAITasksTaskUnion `json:"tasks"`
	JSON  struct {
		Name          respjson.Field
		Options       respjson.Field
		MaxTags       respjson.Field
		MinConfidence respjson.Field
		Tasks         respjson.Field
		raw           string
	} `json:"-"`
}

func (u ExtensionConfigUnion) AsRemoveBg() (v ExtensionConfigRemoveBg) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtensionConfigUnion) AsAutoTagging() (v ExtensionConfigAutoTagging) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtensionConfigUnion) AsAIAutoDescription() (v ExtensionConfigAIAutoDescription) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtensionConfigUnion) AsAITasks() (v ExtensionConfigAITasks) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtensionConfigUnion) RawJSON() string { return u.JSON.raw }

func (r *ExtensionConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ExtensionConfigUnion to a ExtensionConfigUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ExtensionConfigUnionParam.Overrides()
func (r ExtensionConfigUnion) ToParam() ExtensionConfigUnionParam {
	return param.Override[ExtensionConfigUnionParam](json.RawMessage(r.RawJSON()))
}

type ExtensionConfigRemoveBg struct {
	// Specifies the background removal extension.
	Name    constant.RemoveBg              `json:"name,required"`
	Options ExtensionConfigRemoveBgOptions `json:"options"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Options     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtensionConfigRemoveBg) RawJSON() string { return r.JSON.raw }
func (r *ExtensionConfigRemoveBg) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionConfigRemoveBgOptions struct {
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
func (r ExtensionConfigRemoveBgOptions) RawJSON() string { return r.JSON.raw }
func (r *ExtensionConfigRemoveBgOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionConfigAutoTagging struct {
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
func (r ExtensionConfigAutoTagging) RawJSON() string { return r.JSON.raw }
func (r *ExtensionConfigAutoTagging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionConfigAIAutoDescription struct {
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
func (r ExtensionConfigAIAutoDescription) RawJSON() string { return r.JSON.raw }
func (r *ExtensionConfigAIAutoDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionConfigAITasks struct {
	// Specifies the AI tasks extension for automated image analysis using AI models.
	Name constant.AITasks `json:"name,required"`
	// Array of task objects defining AI operations to perform on the asset.
	Tasks []ExtensionConfigAITasksTaskUnion `json:"tasks,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Tasks       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtensionConfigAITasks) RawJSON() string { return r.JSON.raw }
func (r *ExtensionConfigAITasks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExtensionConfigAITasksTaskUnion contains all possible properties and values from
// [ExtensionConfigAITasksTaskSelectTags],
// [ExtensionConfigAITasksTaskSelectMetadata], [ExtensionConfigAITasksTaskYesNo].
//
// Use the [ExtensionConfigAITasksTaskUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ExtensionConfigAITasksTaskUnion struct {
	Instruction string `json:"instruction"`
	// Any of "select_tags", "select_metadata", "yes_no".
	Type          string `json:"type"`
	MaxSelections int64  `json:"max_selections"`
	MinSelections int64  `json:"min_selections"`
	// This field is a union of [[]string],
	// [[]ExtensionConfigAITasksTaskSelectMetadataVocabularyUnion]
	Vocabulary ExtensionConfigAITasksTaskUnionVocabulary `json:"vocabulary"`
	// This field is from variant [ExtensionConfigAITasksTaskSelectMetadata].
	Field string `json:"field"`
	// This field is from variant [ExtensionConfigAITasksTaskYesNo].
	OnNo ExtensionConfigAITasksTaskYesNoOnNo `json:"on_no"`
	// This field is from variant [ExtensionConfigAITasksTaskYesNo].
	OnUnknown ExtensionConfigAITasksTaskYesNoOnUnknown `json:"on_unknown"`
	// This field is from variant [ExtensionConfigAITasksTaskYesNo].
	OnYes ExtensionConfigAITasksTaskYesNoOnYes `json:"on_yes"`
	JSON  struct {
		Instruction   respjson.Field
		Type          respjson.Field
		MaxSelections respjson.Field
		MinSelections respjson.Field
		Vocabulary    respjson.Field
		Field         respjson.Field
		OnNo          respjson.Field
		OnUnknown     respjson.Field
		OnYes         respjson.Field
		raw           string
	} `json:"-"`
}

// anyExtensionConfigAITasksTask is implemented by each variant of
// [ExtensionConfigAITasksTaskUnion] to add type safety for the return type of
// [ExtensionConfigAITasksTaskUnion.AsAny]
type anyExtensionConfigAITasksTask interface {
	implExtensionConfigAITasksTaskUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ExtensionConfigAITasksTaskUnion.AsAny().(type) {
//	case shared.ExtensionConfigAITasksTaskSelectTags:
//	case shared.ExtensionConfigAITasksTaskSelectMetadata:
//	case shared.ExtensionConfigAITasksTaskYesNo:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ExtensionConfigAITasksTaskUnion) AsAny() anyExtensionConfigAITasksTask {
	switch u.Type {
	case "select_tags":
		return u.AsSelectTags()
	case "select_metadata":
		return u.AsSelectMetadata()
	case "yes_no":
		return u.AsYesNo()
	}
	return nil
}

func (u ExtensionConfigAITasksTaskUnion) AsSelectTags() (v ExtensionConfigAITasksTaskSelectTags) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtensionConfigAITasksTaskUnion) AsSelectMetadata() (v ExtensionConfigAITasksTaskSelectMetadata) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtensionConfigAITasksTaskUnion) AsYesNo() (v ExtensionConfigAITasksTaskYesNo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtensionConfigAITasksTaskUnion) RawJSON() string { return u.JSON.raw }

func (r *ExtensionConfigAITasksTaskUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExtensionConfigAITasksTaskUnionVocabulary is an implicit subunion of
// [ExtensionConfigAITasksTaskUnion]. ExtensionConfigAITasksTaskUnionVocabulary
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ExtensionConfigAITasksTaskUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfStringArray
// OfExtensionConfigAITasksTaskSelectMetadataVocabularyArray]
type ExtensionConfigAITasksTaskUnionVocabulary struct {
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	// This field will be present if the value is a
	// [[]ExtensionConfigAITasksTaskSelectMetadataVocabularyUnion] instead of an
	// object.
	OfExtensionConfigAITasksTaskSelectMetadataVocabularyArray []ExtensionConfigAITasksTaskSelectMetadataVocabularyUnion `json:",inline"`
	JSON                                                      struct {
		OfStringArray                                             respjson.Field
		OfExtensionConfigAITasksTaskSelectMetadataVocabularyArray respjson.Field
		raw                                                       string
	} `json:"-"`
}

func (r *ExtensionConfigAITasksTaskUnionVocabulary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionConfigAITasksTaskSelectTags struct {
	// The question or instruction for the AI to analyze the image.
	Instruction string `json:"instruction,required"`
	// Task type that analyzes the image and adds matching tags from a vocabulary.
	Type constant.SelectTags `json:"type,required"`
	// Maximum number of tags to select from the vocabulary.
	MaxSelections int64 `json:"max_selections"`
	// Minimum number of tags to select from the vocabulary.
	MinSelections int64 `json:"min_selections"`
	// Array of possible tag values. Combined length of all strings must not exceed 500
	// characters. Cannot contain the `%` character.
	Vocabulary []string `json:"vocabulary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Instruction   respjson.Field
		Type          respjson.Field
		MaxSelections respjson.Field
		MinSelections respjson.Field
		Vocabulary    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtensionConfigAITasksTaskSelectTags) RawJSON() string { return r.JSON.raw }
func (r *ExtensionConfigAITasksTaskSelectTags) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (ExtensionConfigAITasksTaskSelectTags) implExtensionConfigAITasksTaskUnion() {}

type ExtensionConfigAITasksTaskSelectMetadata struct {
	// Name of the custom metadata field to set. The field must exist in your account.
	Field string `json:"field,required"`
	// The question or instruction for the AI to analyze the image.
	Instruction string `json:"instruction,required"`
	// Task type that analyzes the image and sets a custom metadata field value from a
	// vocabulary.
	Type constant.SelectMetadata `json:"type,required"`
	// Maximum number of values to select from the vocabulary.
	MaxSelections int64 `json:"max_selections"`
	// Minimum number of values to select from the vocabulary.
	MinSelections int64 `json:"min_selections"`
	// Array of possible values matching the custom metadata field type.
	Vocabulary []ExtensionConfigAITasksTaskSelectMetadataVocabularyUnion `json:"vocabulary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field         respjson.Field
		Instruction   respjson.Field
		Type          respjson.Field
		MaxSelections respjson.Field
		MinSelections respjson.Field
		Vocabulary    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtensionConfigAITasksTaskSelectMetadata) RawJSON() string { return r.JSON.raw }
func (r *ExtensionConfigAITasksTaskSelectMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (ExtensionConfigAITasksTaskSelectMetadata) implExtensionConfigAITasksTaskUnion() {}

// ExtensionConfigAITasksTaskSelectMetadataVocabularyUnion contains all possible
// properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type ExtensionConfigAITasksTaskSelectMetadataVocabularyUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u ExtensionConfigAITasksTaskSelectMetadataVocabularyUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtensionConfigAITasksTaskSelectMetadataVocabularyUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtensionConfigAITasksTaskSelectMetadataVocabularyUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtensionConfigAITasksTaskSelectMetadataVocabularyUnion) RawJSON() string { return u.JSON.raw }

func (r *ExtensionConfigAITasksTaskSelectMetadataVocabularyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionConfigAITasksTaskYesNo struct {
	// The yes/no question for the AI to answer about the image.
	Instruction string `json:"instruction,required"`
	// Task type that asks a yes/no question and executes actions based on the answer.
	Type constant.YesNo `json:"type,required"`
	// Actions to execute if the AI answers no.
	OnNo ExtensionConfigAITasksTaskYesNoOnNo `json:"on_no"`
	// Actions to execute if the AI cannot determine the answer.
	OnUnknown ExtensionConfigAITasksTaskYesNoOnUnknown `json:"on_unknown"`
	// Actions to execute if the AI answers yes.
	OnYes ExtensionConfigAITasksTaskYesNoOnYes `json:"on_yes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Instruction respjson.Field
		Type        respjson.Field
		OnNo        respjson.Field
		OnUnknown   respjson.Field
		OnYes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtensionConfigAITasksTaskYesNo) RawJSON() string { return r.JSON.raw }
func (r *ExtensionConfigAITasksTaskYesNo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (ExtensionConfigAITasksTaskYesNo) implExtensionConfigAITasksTaskUnion() {}

// Actions to execute if the AI answers no.
type ExtensionConfigAITasksTaskYesNoOnNo struct {
	// Array of tag strings to add to the asset.
	AddTags []string `json:"add_tags"`
	// Array of tag strings to remove from the asset.
	RemoveTags []string `json:"remove_tags"`
	// Array of custom metadata field updates.
	SetMetadata []ExtensionConfigAITasksTaskYesNoOnNoSetMetadata `json:"set_metadata"`
	// Array of custom metadata fields to remove.
	UnsetMetadata []ExtensionConfigAITasksTaskYesNoOnNoUnsetMetadata `json:"unset_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddTags       respjson.Field
		RemoveTags    respjson.Field
		SetMetadata   respjson.Field
		UnsetMetadata respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtensionConfigAITasksTaskYesNoOnNo) RawJSON() string { return r.JSON.raw }
func (r *ExtensionConfigAITasksTaskYesNoOnNo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionConfigAITasksTaskYesNoOnNoSetMetadata struct {
	// Name of the custom metadata field to set.
	Field string `json:"field,required"`
	// Value to set for the custom metadata field. The value type should match the
	// custom metadata field type.
	Value ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueUnion `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtensionConfigAITasksTaskYesNoOnNoSetMetadata) RawJSON() string { return r.JSON.raw }
func (r *ExtensionConfigAITasksTaskYesNoOnNoSetMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueUnion contains all possible
// properties and values from [string], [float64], [bool],
// [[]ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueMixedItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfMixed]
type ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueMixedItemUnion] instead of
	// an object.
	OfMixed []ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueMixedItemUnion `json:",inline"`
	JSON    struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		OfMixed  respjson.Field
		raw      string
	} `json:"-"`
}

func (u ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueUnion) AsMixed() (v []ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueMixedItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueUnion) RawJSON() string { return u.JSON.raw }

func (r *ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueMixedItemUnion contains all
// possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueMixedItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueMixedItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueMixedItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueMixedItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueMixedItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueMixedItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionConfigAITasksTaskYesNoOnNoUnsetMetadata struct {
	// Name of the custom metadata field to remove.
	Field string `json:"field,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtensionConfigAITasksTaskYesNoOnNoUnsetMetadata) RawJSON() string { return r.JSON.raw }
func (r *ExtensionConfigAITasksTaskYesNoOnNoUnsetMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Actions to execute if the AI cannot determine the answer.
type ExtensionConfigAITasksTaskYesNoOnUnknown struct {
	// Array of tag strings to add to the asset.
	AddTags []string `json:"add_tags"`
	// Array of tag strings to remove from the asset.
	RemoveTags []string `json:"remove_tags"`
	// Array of custom metadata field updates.
	SetMetadata []ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadata `json:"set_metadata"`
	// Array of custom metadata fields to remove.
	UnsetMetadata []ExtensionConfigAITasksTaskYesNoOnUnknownUnsetMetadata `json:"unset_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddTags       respjson.Field
		RemoveTags    respjson.Field
		SetMetadata   respjson.Field
		UnsetMetadata respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtensionConfigAITasksTaskYesNoOnUnknown) RawJSON() string { return r.JSON.raw }
func (r *ExtensionConfigAITasksTaskYesNoOnUnknown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadata struct {
	// Name of the custom metadata field to set.
	Field string `json:"field,required"`
	// Value to set for the custom metadata field. The value type should match the
	// custom metadata field type.
	Value ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueUnion `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadata) RawJSON() string { return r.JSON.raw }
func (r *ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueUnion contains all
// possible properties and values from [string], [float64], [bool],
// [[]ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueMixedItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfMixed]
type ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueMixedItemUnion]
	// instead of an object.
	OfMixed []ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueMixedItemUnion `json:",inline"`
	JSON    struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		OfMixed  respjson.Field
		raw      string
	} `json:"-"`
}

func (u ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueUnion) AsMixed() (v []ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueMixedItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueMixedItemUnion contains
// all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueMixedItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueMixedItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueMixedItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueMixedItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueMixedItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueMixedItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionConfigAITasksTaskYesNoOnUnknownUnsetMetadata struct {
	// Name of the custom metadata field to remove.
	Field string `json:"field,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtensionConfigAITasksTaskYesNoOnUnknownUnsetMetadata) RawJSON() string { return r.JSON.raw }
func (r *ExtensionConfigAITasksTaskYesNoOnUnknownUnsetMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Actions to execute if the AI answers yes.
type ExtensionConfigAITasksTaskYesNoOnYes struct {
	// Array of tag strings to add to the asset.
	AddTags []string `json:"add_tags"`
	// Array of tag strings to remove from the asset.
	RemoveTags []string `json:"remove_tags"`
	// Array of custom metadata field updates.
	SetMetadata []ExtensionConfigAITasksTaskYesNoOnYesSetMetadata `json:"set_metadata"`
	// Array of custom metadata fields to remove.
	UnsetMetadata []ExtensionConfigAITasksTaskYesNoOnYesUnsetMetadata `json:"unset_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddTags       respjson.Field
		RemoveTags    respjson.Field
		SetMetadata   respjson.Field
		UnsetMetadata respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtensionConfigAITasksTaskYesNoOnYes) RawJSON() string { return r.JSON.raw }
func (r *ExtensionConfigAITasksTaskYesNoOnYes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionConfigAITasksTaskYesNoOnYesSetMetadata struct {
	// Name of the custom metadata field to set.
	Field string `json:"field,required"`
	// Value to set for the custom metadata field. The value type should match the
	// custom metadata field type.
	Value ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueUnion `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtensionConfigAITasksTaskYesNoOnYesSetMetadata) RawJSON() string { return r.JSON.raw }
func (r *ExtensionConfigAITasksTaskYesNoOnYesSetMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueUnion contains all possible
// properties and values from [string], [float64], [bool],
// [[]ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueMixedItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfMixed]
type ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueMixedItemUnion] instead
	// of an object.
	OfMixed []ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueMixedItemUnion `json:",inline"`
	JSON    struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		OfMixed  respjson.Field
		raw      string
	} `json:"-"`
}

func (u ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueUnion) AsMixed() (v []ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueMixedItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueMixedItemUnion contains all
// possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueMixedItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueMixedItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueMixedItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueMixedItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueMixedItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueMixedItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionConfigAITasksTaskYesNoOnYesUnsetMetadata struct {
	// Name of the custom metadata field to remove.
	Field string `json:"field,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtensionConfigAITasksTaskYesNoOnYesUnsetMetadata) RawJSON() string { return r.JSON.raw }
func (r *ExtensionConfigAITasksTaskYesNoOnYesUnsetMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func ExtensionConfigParamOfAutoTagging(maxTags int64, minConfidence int64, name string) ExtensionConfigUnionParam {
	var variant ExtensionConfigAutoTaggingParam
	variant.MaxTags = maxTags
	variant.MinConfidence = minConfidence
	variant.Name = name
	return ExtensionConfigUnionParam{OfAutoTagging: &variant}
}

func ExtensionConfigParamOfAITasks(tasks []ExtensionConfigAITasksTaskUnionParam) ExtensionConfigUnionParam {
	var aiTasks ExtensionConfigAITasksParam
	aiTasks.Tasks = tasks
	return ExtensionConfigUnionParam{OfAITasks: &aiTasks}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtensionConfigUnionParam struct {
	OfRemoveBg          *ExtensionConfigRemoveBgParam          `json:",omitzero,inline"`
	OfAutoTagging       *ExtensionConfigAutoTaggingParam       `json:",omitzero,inline"`
	OfAIAutoDescription *ExtensionConfigAIAutoDescriptionParam `json:",omitzero,inline"`
	OfAITasks           *ExtensionConfigAITasksParam           `json:",omitzero,inline"`
	paramUnion
}

func (u ExtensionConfigUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRemoveBg, u.OfAutoTagging, u.OfAIAutoDescription, u.OfAITasks)
}
func (u *ExtensionConfigUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtensionConfigUnionParam) asAny() any {
	if !param.IsOmitted(u.OfRemoveBg) {
		return u.OfRemoveBg
	} else if !param.IsOmitted(u.OfAutoTagging) {
		return u.OfAutoTagging
	} else if !param.IsOmitted(u.OfAIAutoDescription) {
		return u.OfAIAutoDescription
	} else if !param.IsOmitted(u.OfAITasks) {
		return u.OfAITasks
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtensionConfigUnionParam) GetOptions() *ExtensionConfigRemoveBgOptionsParam {
	if vt := u.OfRemoveBg; vt != nil {
		return &vt.Options
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtensionConfigUnionParam) GetMaxTags() *int64 {
	if vt := u.OfAutoTagging; vt != nil {
		return &vt.MaxTags
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtensionConfigUnionParam) GetMinConfidence() *int64 {
	if vt := u.OfAutoTagging; vt != nil {
		return &vt.MinConfidence
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtensionConfigUnionParam) GetTasks() []ExtensionConfigAITasksTaskUnionParam {
	if vt := u.OfAITasks; vt != nil {
		return vt.Tasks
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtensionConfigUnionParam) GetName() *string {
	if vt := u.OfRemoveBg; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfAutoTagging; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfAIAutoDescription; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfAITasks; vt != nil {
		return (*string)(&vt.Name)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ExtensionConfigUnionParam](
		"name",
		apijson.Discriminator[ExtensionConfigRemoveBgParam]("remove-bg"),
		apijson.Discriminator[ExtensionConfigAutoTaggingParam]("google-auto-tagging"),
		apijson.Discriminator[ExtensionConfigAutoTaggingParam]("aws-auto-tagging"),
		apijson.Discriminator[ExtensionConfigAIAutoDescriptionParam]("ai-auto-description"),
		apijson.Discriminator[ExtensionConfigAITasksParam]("ai-tasks"),
	)
}

// The property Name is required.
type ExtensionConfigRemoveBgParam struct {
	Options ExtensionConfigRemoveBgOptionsParam `json:"options,omitzero"`
	// Specifies the background removal extension.
	//
	// This field can be elided, and will marshal its zero value as "remove-bg".
	Name constant.RemoveBg `json:"name,required"`
	paramObj
}

func (r ExtensionConfigRemoveBgParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionConfigRemoveBgParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionConfigRemoveBgParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtensionConfigRemoveBgOptionsParam struct {
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

func (r ExtensionConfigRemoveBgOptionsParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionConfigRemoveBgOptionsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionConfigRemoveBgOptionsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MaxTags, MinConfidence, Name are required.
type ExtensionConfigAutoTaggingParam struct {
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

func (r ExtensionConfigAutoTaggingParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionConfigAutoTaggingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionConfigAutoTaggingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtensionConfigAutoTaggingParam](
		"name", "google-auto-tagging", "aws-auto-tagging",
	)
}

func NewExtensionConfigAIAutoDescriptionParam() ExtensionConfigAIAutoDescriptionParam {
	return ExtensionConfigAIAutoDescriptionParam{
		Name: "ai-auto-description",
	}
}

// This struct has a constant value, construct it with
// [NewExtensionConfigAIAutoDescriptionParam].
type ExtensionConfigAIAutoDescriptionParam struct {
	// Specifies the auto description extension.
	Name constant.AIAutoDescription `json:"name,required"`
	paramObj
}

func (r ExtensionConfigAIAutoDescriptionParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionConfigAIAutoDescriptionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionConfigAIAutoDescriptionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Tasks are required.
type ExtensionConfigAITasksParam struct {
	// Array of task objects defining AI operations to perform on the asset.
	Tasks []ExtensionConfigAITasksTaskUnionParam `json:"tasks,omitzero,required"`
	// Specifies the AI tasks extension for automated image analysis using AI models.
	//
	// This field can be elided, and will marshal its zero value as "ai-tasks".
	Name constant.AITasks `json:"name,required"`
	paramObj
}

func (r ExtensionConfigAITasksParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionConfigAITasksParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionConfigAITasksParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtensionConfigAITasksTaskUnionParam struct {
	OfSelectTags     *ExtensionConfigAITasksTaskSelectTagsParam     `json:",omitzero,inline"`
	OfSelectMetadata *ExtensionConfigAITasksTaskSelectMetadataParam `json:",omitzero,inline"`
	OfYesNo          *ExtensionConfigAITasksTaskYesNoParam          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtensionConfigAITasksTaskUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSelectTags, u.OfSelectMetadata, u.OfYesNo)
}
func (u *ExtensionConfigAITasksTaskUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtensionConfigAITasksTaskUnionParam) asAny() any {
	if !param.IsOmitted(u.OfSelectTags) {
		return u.OfSelectTags
	} else if !param.IsOmitted(u.OfSelectMetadata) {
		return u.OfSelectMetadata
	} else if !param.IsOmitted(u.OfYesNo) {
		return u.OfYesNo
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtensionConfigAITasksTaskUnionParam) GetField() *string {
	if vt := u.OfSelectMetadata; vt != nil {
		return &vt.Field
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtensionConfigAITasksTaskUnionParam) GetOnNo() *ExtensionConfigAITasksTaskYesNoOnNoParam {
	if vt := u.OfYesNo; vt != nil {
		return &vt.OnNo
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtensionConfigAITasksTaskUnionParam) GetOnUnknown() *ExtensionConfigAITasksTaskYesNoOnUnknownParam {
	if vt := u.OfYesNo; vt != nil {
		return &vt.OnUnknown
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtensionConfigAITasksTaskUnionParam) GetOnYes() *ExtensionConfigAITasksTaskYesNoOnYesParam {
	if vt := u.OfYesNo; vt != nil {
		return &vt.OnYes
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtensionConfigAITasksTaskUnionParam) GetInstruction() *string {
	if vt := u.OfSelectTags; vt != nil {
		return (*string)(&vt.Instruction)
	} else if vt := u.OfSelectMetadata; vt != nil {
		return (*string)(&vt.Instruction)
	} else if vt := u.OfYesNo; vt != nil {
		return (*string)(&vt.Instruction)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtensionConfigAITasksTaskUnionParam) GetType() *string {
	if vt := u.OfSelectTags; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSelectMetadata; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfYesNo; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtensionConfigAITasksTaskUnionParam) GetMaxSelections() *int64 {
	if vt := u.OfSelectTags; vt != nil && vt.MaxSelections.Valid() {
		return &vt.MaxSelections.Value
	} else if vt := u.OfSelectMetadata; vt != nil && vt.MaxSelections.Valid() {
		return &vt.MaxSelections.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtensionConfigAITasksTaskUnionParam) GetMinSelections() *int64 {
	if vt := u.OfSelectTags; vt != nil && vt.MinSelections.Valid() {
		return &vt.MinSelections.Value
	} else if vt := u.OfSelectMetadata; vt != nil && vt.MinSelections.Valid() {
		return &vt.MinSelections.Value
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ExtensionConfigAITasksTaskUnionParam) GetVocabulary() (res extensionConfigAITasksTaskUnionParamVocabulary) {
	if vt := u.OfSelectTags; vt != nil {
		res.any = &vt.Vocabulary
	} else if vt := u.OfSelectMetadata; vt != nil {
		res.any = &vt.Vocabulary
	}
	return
}

// Can have the runtime types [_[]string],
// [_[]ExtensionConfigAITasksTaskSelectMetadataVocabularyUnionParam]
type extensionConfigAITasksTaskUnionParamVocabulary struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]string:
//	case *[]shared.ExtensionConfigAITasksTaskSelectMetadataVocabularyUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u extensionConfigAITasksTaskUnionParamVocabulary) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[ExtensionConfigAITasksTaskUnionParam](
		"type",
		apijson.Discriminator[ExtensionConfigAITasksTaskSelectTagsParam]("select_tags"),
		apijson.Discriminator[ExtensionConfigAITasksTaskSelectMetadataParam]("select_metadata"),
		apijson.Discriminator[ExtensionConfigAITasksTaskYesNoParam]("yes_no"),
	)
}

// The properties Instruction, Type are required.
type ExtensionConfigAITasksTaskSelectTagsParam struct {
	// The question or instruction for the AI to analyze the image.
	Instruction string `json:"instruction,required"`
	// Maximum number of tags to select from the vocabulary.
	MaxSelections param.Opt[int64] `json:"max_selections,omitzero"`
	// Minimum number of tags to select from the vocabulary.
	MinSelections param.Opt[int64] `json:"min_selections,omitzero"`
	// Array of possible tag values. Combined length of all strings must not exceed 500
	// characters. Cannot contain the `%` character.
	Vocabulary []string `json:"vocabulary,omitzero"`
	// Task type that analyzes the image and adds matching tags from a vocabulary.
	//
	// This field can be elided, and will marshal its zero value as "select_tags".
	Type constant.SelectTags `json:"type,required"`
	paramObj
}

func (r ExtensionConfigAITasksTaskSelectTagsParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionConfigAITasksTaskSelectTagsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionConfigAITasksTaskSelectTagsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Field, Instruction, Type are required.
type ExtensionConfigAITasksTaskSelectMetadataParam struct {
	// Name of the custom metadata field to set. The field must exist in your account.
	Field string `json:"field,required"`
	// The question or instruction for the AI to analyze the image.
	Instruction string `json:"instruction,required"`
	// Maximum number of values to select from the vocabulary.
	MaxSelections param.Opt[int64] `json:"max_selections,omitzero"`
	// Minimum number of values to select from the vocabulary.
	MinSelections param.Opt[int64] `json:"min_selections,omitzero"`
	// Array of possible values matching the custom metadata field type.
	Vocabulary []ExtensionConfigAITasksTaskSelectMetadataVocabularyUnionParam `json:"vocabulary,omitzero"`
	// Task type that analyzes the image and sets a custom metadata field value from a
	// vocabulary.
	//
	// This field can be elided, and will marshal its zero value as "select_metadata".
	Type constant.SelectMetadata `json:"type,required"`
	paramObj
}

func (r ExtensionConfigAITasksTaskSelectMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionConfigAITasksTaskSelectMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionConfigAITasksTaskSelectMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtensionConfigAITasksTaskSelectMetadataVocabularyUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u ExtensionConfigAITasksTaskSelectMetadataVocabularyUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *ExtensionConfigAITasksTaskSelectMetadataVocabularyUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtensionConfigAITasksTaskSelectMetadataVocabularyUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// The properties Instruction, Type are required.
type ExtensionConfigAITasksTaskYesNoParam struct {
	// The yes/no question for the AI to answer about the image.
	Instruction string `json:"instruction,required"`
	// Actions to execute if the AI answers no.
	OnNo ExtensionConfigAITasksTaskYesNoOnNoParam `json:"on_no,omitzero"`
	// Actions to execute if the AI cannot determine the answer.
	OnUnknown ExtensionConfigAITasksTaskYesNoOnUnknownParam `json:"on_unknown,omitzero"`
	// Actions to execute if the AI answers yes.
	OnYes ExtensionConfigAITasksTaskYesNoOnYesParam `json:"on_yes,omitzero"`
	// Task type that asks a yes/no question and executes actions based on the answer.
	//
	// This field can be elided, and will marshal its zero value as "yes_no".
	Type constant.YesNo `json:"type,required"`
	paramObj
}

func (r ExtensionConfigAITasksTaskYesNoParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionConfigAITasksTaskYesNoParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionConfigAITasksTaskYesNoParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Actions to execute if the AI answers no.
type ExtensionConfigAITasksTaskYesNoOnNoParam struct {
	// Array of tag strings to add to the asset.
	AddTags []string `json:"add_tags,omitzero"`
	// Array of tag strings to remove from the asset.
	RemoveTags []string `json:"remove_tags,omitzero"`
	// Array of custom metadata field updates.
	SetMetadata []ExtensionConfigAITasksTaskYesNoOnNoSetMetadataParam `json:"set_metadata,omitzero"`
	// Array of custom metadata fields to remove.
	UnsetMetadata []ExtensionConfigAITasksTaskYesNoOnNoUnsetMetadataParam `json:"unset_metadata,omitzero"`
	paramObj
}

func (r ExtensionConfigAITasksTaskYesNoOnNoParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionConfigAITasksTaskYesNoOnNoParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionConfigAITasksTaskYesNoOnNoParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Field, Value are required.
type ExtensionConfigAITasksTaskYesNoOnNoSetMetadataParam struct {
	// Name of the custom metadata field to set.
	Field string `json:"field,required"`
	// Value to set for the custom metadata field. The value type should match the
	// custom metadata field type.
	Value ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueUnionParam `json:"value,omitzero,required"`
	paramObj
}

func (r ExtensionConfigAITasksTaskYesNoOnNoSetMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionConfigAITasksTaskYesNoOnNoSetMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionConfigAITasksTaskYesNoOnNoSetMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueUnionParam struct {
	OfString param.Opt[string]                                                        `json:",omitzero,inline"`
	OfFloat  param.Opt[float64]                                                       `json:",omitzero,inline"`
	OfBool   param.Opt[bool]                                                          `json:",omitzero,inline"`
	OfMixed  []ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueMixedItemUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfMixed)
}
func (u *ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfMixed) {
		return &u.OfMixed
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueMixedItemUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueMixedItemUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueMixedItemUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtensionConfigAITasksTaskYesNoOnNoSetMetadataValueMixedItemUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// The property Field is required.
type ExtensionConfigAITasksTaskYesNoOnNoUnsetMetadataParam struct {
	// Name of the custom metadata field to remove.
	Field string `json:"field,required"`
	paramObj
}

func (r ExtensionConfigAITasksTaskYesNoOnNoUnsetMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionConfigAITasksTaskYesNoOnNoUnsetMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionConfigAITasksTaskYesNoOnNoUnsetMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Actions to execute if the AI cannot determine the answer.
type ExtensionConfigAITasksTaskYesNoOnUnknownParam struct {
	// Array of tag strings to add to the asset.
	AddTags []string `json:"add_tags,omitzero"`
	// Array of tag strings to remove from the asset.
	RemoveTags []string `json:"remove_tags,omitzero"`
	// Array of custom metadata field updates.
	SetMetadata []ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataParam `json:"set_metadata,omitzero"`
	// Array of custom metadata fields to remove.
	UnsetMetadata []ExtensionConfigAITasksTaskYesNoOnUnknownUnsetMetadataParam `json:"unset_metadata,omitzero"`
	paramObj
}

func (r ExtensionConfigAITasksTaskYesNoOnUnknownParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionConfigAITasksTaskYesNoOnUnknownParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionConfigAITasksTaskYesNoOnUnknownParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Field, Value are required.
type ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataParam struct {
	// Name of the custom metadata field to set.
	Field string `json:"field,required"`
	// Value to set for the custom metadata field. The value type should match the
	// custom metadata field type.
	Value ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueUnionParam `json:"value,omitzero,required"`
	paramObj
}

func (r ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueUnionParam struct {
	OfString param.Opt[string]                                                             `json:",omitzero,inline"`
	OfFloat  param.Opt[float64]                                                            `json:",omitzero,inline"`
	OfBool   param.Opt[bool]                                                               `json:",omitzero,inline"`
	OfMixed  []ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueMixedItemUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfMixed)
}
func (u *ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfMixed) {
		return &u.OfMixed
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueMixedItemUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueMixedItemUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueMixedItemUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtensionConfigAITasksTaskYesNoOnUnknownSetMetadataValueMixedItemUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// The property Field is required.
type ExtensionConfigAITasksTaskYesNoOnUnknownUnsetMetadataParam struct {
	// Name of the custom metadata field to remove.
	Field string `json:"field,required"`
	paramObj
}

func (r ExtensionConfigAITasksTaskYesNoOnUnknownUnsetMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionConfigAITasksTaskYesNoOnUnknownUnsetMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionConfigAITasksTaskYesNoOnUnknownUnsetMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Actions to execute if the AI answers yes.
type ExtensionConfigAITasksTaskYesNoOnYesParam struct {
	// Array of tag strings to add to the asset.
	AddTags []string `json:"add_tags,omitzero"`
	// Array of tag strings to remove from the asset.
	RemoveTags []string `json:"remove_tags,omitzero"`
	// Array of custom metadata field updates.
	SetMetadata []ExtensionConfigAITasksTaskYesNoOnYesSetMetadataParam `json:"set_metadata,omitzero"`
	// Array of custom metadata fields to remove.
	UnsetMetadata []ExtensionConfigAITasksTaskYesNoOnYesUnsetMetadataParam `json:"unset_metadata,omitzero"`
	paramObj
}

func (r ExtensionConfigAITasksTaskYesNoOnYesParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionConfigAITasksTaskYesNoOnYesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionConfigAITasksTaskYesNoOnYesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Field, Value are required.
type ExtensionConfigAITasksTaskYesNoOnYesSetMetadataParam struct {
	// Name of the custom metadata field to set.
	Field string `json:"field,required"`
	// Value to set for the custom metadata field. The value type should match the
	// custom metadata field type.
	Value ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueUnionParam `json:"value,omitzero,required"`
	paramObj
}

func (r ExtensionConfigAITasksTaskYesNoOnYesSetMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionConfigAITasksTaskYesNoOnYesSetMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionConfigAITasksTaskYesNoOnYesSetMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueUnionParam struct {
	OfString param.Opt[string]                                                         `json:",omitzero,inline"`
	OfFloat  param.Opt[float64]                                                        `json:",omitzero,inline"`
	OfBool   param.Opt[bool]                                                           `json:",omitzero,inline"`
	OfMixed  []ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueMixedItemUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfMixed)
}
func (u *ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfMixed) {
		return &u.OfMixed
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueMixedItemUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueMixedItemUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueMixedItemUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtensionConfigAITasksTaskYesNoOnYesSetMetadataValueMixedItemUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// The property Field is required.
type ExtensionConfigAITasksTaskYesNoOnYesUnsetMetadataParam struct {
	// Name of the custom metadata field to remove.
	Field string `json:"field,required"`
	paramObj
}

func (r ExtensionConfigAITasksTaskYesNoOnYesUnsetMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionConfigAITasksTaskYesNoOnYesUnsetMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionConfigAITasksTaskYesNoOnYesUnsetMetadataParam) UnmarshalJSON(data []byte) error {
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
	OfAITasks           *ExtensionAITasksParam           `json:",omitzero,inline"`
	OfSavedExtension    *ExtensionSavedExtensionParam    `json:",omitzero,inline"`
	paramUnion
}

func (u ExtensionUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRemoveBg,
		u.OfAutoTagging,
		u.OfAIAutoDescription,
		u.OfAITasks,
		u.OfSavedExtension)
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
	} else if !param.IsOmitted(u.OfAITasks) {
		return u.OfAITasks
	} else if !param.IsOmitted(u.OfSavedExtension) {
		return u.OfSavedExtension
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
func (u ExtensionUnionParam) GetTasks() []ExtensionAITasksTaskUnionParam {
	if vt := u.OfAITasks; vt != nil {
		return vt.Tasks
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtensionUnionParam) GetID() *string {
	if vt := u.OfSavedExtension; vt != nil {
		return &vt.ID
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
	} else if vt := u.OfAITasks; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfSavedExtension; vt != nil {
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
		apijson.Discriminator[ExtensionAITasksParam]("ai-tasks"),
		apijson.Discriminator[ExtensionSavedExtensionParam]("saved-extension"),
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

// The properties Name, Tasks are required.
type ExtensionAITasksParam struct {
	// Array of task objects defining AI operations to perform on the asset.
	Tasks []ExtensionAITasksTaskUnionParam `json:"tasks,omitzero,required"`
	// Specifies the AI tasks extension for automated image analysis using AI models.
	//
	// This field can be elided, and will marshal its zero value as "ai-tasks".
	Name constant.AITasks `json:"name,required"`
	paramObj
}

func (r ExtensionAITasksParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionAITasksParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionAITasksParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtensionAITasksTaskUnionParam struct {
	OfSelectTags     *ExtensionAITasksTaskSelectTagsParam     `json:",omitzero,inline"`
	OfSelectMetadata *ExtensionAITasksTaskSelectMetadataParam `json:",omitzero,inline"`
	OfYesNo          *ExtensionAITasksTaskYesNoParam          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtensionAITasksTaskUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSelectTags, u.OfSelectMetadata, u.OfYesNo)
}
func (u *ExtensionAITasksTaskUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtensionAITasksTaskUnionParam) asAny() any {
	if !param.IsOmitted(u.OfSelectTags) {
		return u.OfSelectTags
	} else if !param.IsOmitted(u.OfSelectMetadata) {
		return u.OfSelectMetadata
	} else if !param.IsOmitted(u.OfYesNo) {
		return u.OfYesNo
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtensionAITasksTaskUnionParam) GetField() *string {
	if vt := u.OfSelectMetadata; vt != nil {
		return &vt.Field
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtensionAITasksTaskUnionParam) GetOnNo() *ExtensionAITasksTaskYesNoOnNoParam {
	if vt := u.OfYesNo; vt != nil {
		return &vt.OnNo
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtensionAITasksTaskUnionParam) GetOnUnknown() *ExtensionAITasksTaskYesNoOnUnknownParam {
	if vt := u.OfYesNo; vt != nil {
		return &vt.OnUnknown
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtensionAITasksTaskUnionParam) GetOnYes() *ExtensionAITasksTaskYesNoOnYesParam {
	if vt := u.OfYesNo; vt != nil {
		return &vt.OnYes
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtensionAITasksTaskUnionParam) GetInstruction() *string {
	if vt := u.OfSelectTags; vt != nil {
		return (*string)(&vt.Instruction)
	} else if vt := u.OfSelectMetadata; vt != nil {
		return (*string)(&vt.Instruction)
	} else if vt := u.OfYesNo; vt != nil {
		return (*string)(&vt.Instruction)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtensionAITasksTaskUnionParam) GetType() *string {
	if vt := u.OfSelectTags; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSelectMetadata; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfYesNo; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtensionAITasksTaskUnionParam) GetMaxSelections() *int64 {
	if vt := u.OfSelectTags; vt != nil && vt.MaxSelections.Valid() {
		return &vt.MaxSelections.Value
	} else if vt := u.OfSelectMetadata; vt != nil && vt.MaxSelections.Valid() {
		return &vt.MaxSelections.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtensionAITasksTaskUnionParam) GetMinSelections() *int64 {
	if vt := u.OfSelectTags; vt != nil && vt.MinSelections.Valid() {
		return &vt.MinSelections.Value
	} else if vt := u.OfSelectMetadata; vt != nil && vt.MinSelections.Valid() {
		return &vt.MinSelections.Value
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ExtensionAITasksTaskUnionParam) GetVocabulary() (res extensionAITasksTaskUnionParamVocabulary) {
	if vt := u.OfSelectTags; vt != nil {
		res.any = &vt.Vocabulary
	} else if vt := u.OfSelectMetadata; vt != nil {
		res.any = &vt.Vocabulary
	}
	return
}

// Can have the runtime types [_[]string],
// [_[]ExtensionAITasksTaskSelectMetadataVocabularyUnionParam]
type extensionAITasksTaskUnionParamVocabulary struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]string:
//	case *[]shared.ExtensionAITasksTaskSelectMetadataVocabularyUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u extensionAITasksTaskUnionParamVocabulary) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[ExtensionAITasksTaskUnionParam](
		"type",
		apijson.Discriminator[ExtensionAITasksTaskSelectTagsParam]("select_tags"),
		apijson.Discriminator[ExtensionAITasksTaskSelectMetadataParam]("select_metadata"),
		apijson.Discriminator[ExtensionAITasksTaskYesNoParam]("yes_no"),
	)
}

// The properties Instruction, Type are required.
type ExtensionAITasksTaskSelectTagsParam struct {
	// The question or instruction for the AI to analyze the image.
	Instruction string `json:"instruction,required"`
	// Maximum number of tags to select from the vocabulary.
	MaxSelections param.Opt[int64] `json:"max_selections,omitzero"`
	// Minimum number of tags to select from the vocabulary.
	MinSelections param.Opt[int64] `json:"min_selections,omitzero"`
	// Array of possible tag values. Combined length of all strings must not exceed 500
	// characters. Cannot contain the `%` character.
	Vocabulary []string `json:"vocabulary,omitzero"`
	// Task type that analyzes the image and adds matching tags from a vocabulary.
	//
	// This field can be elided, and will marshal its zero value as "select_tags".
	Type constant.SelectTags `json:"type,required"`
	paramObj
}

func (r ExtensionAITasksTaskSelectTagsParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionAITasksTaskSelectTagsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionAITasksTaskSelectTagsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Field, Instruction, Type are required.
type ExtensionAITasksTaskSelectMetadataParam struct {
	// Name of the custom metadata field to set. The field must exist in your account.
	Field string `json:"field,required"`
	// The question or instruction for the AI to analyze the image.
	Instruction string `json:"instruction,required"`
	// Maximum number of values to select from the vocabulary.
	MaxSelections param.Opt[int64] `json:"max_selections,omitzero"`
	// Minimum number of values to select from the vocabulary.
	MinSelections param.Opt[int64] `json:"min_selections,omitzero"`
	// Array of possible values matching the custom metadata field type.
	Vocabulary []ExtensionAITasksTaskSelectMetadataVocabularyUnionParam `json:"vocabulary,omitzero"`
	// Task type that analyzes the image and sets a custom metadata field value from a
	// vocabulary.
	//
	// This field can be elided, and will marshal its zero value as "select_metadata".
	Type constant.SelectMetadata `json:"type,required"`
	paramObj
}

func (r ExtensionAITasksTaskSelectMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionAITasksTaskSelectMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionAITasksTaskSelectMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtensionAITasksTaskSelectMetadataVocabularyUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u ExtensionAITasksTaskSelectMetadataVocabularyUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *ExtensionAITasksTaskSelectMetadataVocabularyUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtensionAITasksTaskSelectMetadataVocabularyUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// The properties Instruction, Type are required.
type ExtensionAITasksTaskYesNoParam struct {
	// The yes/no question for the AI to answer about the image.
	Instruction string `json:"instruction,required"`
	// Actions to execute if the AI answers no.
	OnNo ExtensionAITasksTaskYesNoOnNoParam `json:"on_no,omitzero"`
	// Actions to execute if the AI cannot determine the answer.
	OnUnknown ExtensionAITasksTaskYesNoOnUnknownParam `json:"on_unknown,omitzero"`
	// Actions to execute if the AI answers yes.
	OnYes ExtensionAITasksTaskYesNoOnYesParam `json:"on_yes,omitzero"`
	// Task type that asks a yes/no question and executes actions based on the answer.
	//
	// This field can be elided, and will marshal its zero value as "yes_no".
	Type constant.YesNo `json:"type,required"`
	paramObj
}

func (r ExtensionAITasksTaskYesNoParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionAITasksTaskYesNoParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionAITasksTaskYesNoParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Actions to execute if the AI answers no.
type ExtensionAITasksTaskYesNoOnNoParam struct {
	// Array of tag strings to add to the asset.
	AddTags []string `json:"add_tags,omitzero"`
	// Array of tag strings to remove from the asset.
	RemoveTags []string `json:"remove_tags,omitzero"`
	// Array of custom metadata field updates.
	SetMetadata []ExtensionAITasksTaskYesNoOnNoSetMetadataParam `json:"set_metadata,omitzero"`
	// Array of custom metadata fields to remove.
	UnsetMetadata []ExtensionAITasksTaskYesNoOnNoUnsetMetadataParam `json:"unset_metadata,omitzero"`
	paramObj
}

func (r ExtensionAITasksTaskYesNoOnNoParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionAITasksTaskYesNoOnNoParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionAITasksTaskYesNoOnNoParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Field, Value are required.
type ExtensionAITasksTaskYesNoOnNoSetMetadataParam struct {
	// Name of the custom metadata field to set.
	Field string `json:"field,required"`
	// Value to set for the custom metadata field. The value type should match the
	// custom metadata field type.
	Value ExtensionAITasksTaskYesNoOnNoSetMetadataValueUnionParam `json:"value,omitzero,required"`
	paramObj
}

func (r ExtensionAITasksTaskYesNoOnNoSetMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionAITasksTaskYesNoOnNoSetMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionAITasksTaskYesNoOnNoSetMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtensionAITasksTaskYesNoOnNoSetMetadataValueUnionParam struct {
	OfString param.Opt[string]                                                  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64]                                                 `json:",omitzero,inline"`
	OfBool   param.Opt[bool]                                                    `json:",omitzero,inline"`
	OfMixed  []ExtensionAITasksTaskYesNoOnNoSetMetadataValueMixedItemUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u ExtensionAITasksTaskYesNoOnNoSetMetadataValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfMixed)
}
func (u *ExtensionAITasksTaskYesNoOnNoSetMetadataValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtensionAITasksTaskYesNoOnNoSetMetadataValueUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfMixed) {
		return &u.OfMixed
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtensionAITasksTaskYesNoOnNoSetMetadataValueMixedItemUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u ExtensionAITasksTaskYesNoOnNoSetMetadataValueMixedItemUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *ExtensionAITasksTaskYesNoOnNoSetMetadataValueMixedItemUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtensionAITasksTaskYesNoOnNoSetMetadataValueMixedItemUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// The property Field is required.
type ExtensionAITasksTaskYesNoOnNoUnsetMetadataParam struct {
	// Name of the custom metadata field to remove.
	Field string `json:"field,required"`
	paramObj
}

func (r ExtensionAITasksTaskYesNoOnNoUnsetMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionAITasksTaskYesNoOnNoUnsetMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionAITasksTaskYesNoOnNoUnsetMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Actions to execute if the AI cannot determine the answer.
type ExtensionAITasksTaskYesNoOnUnknownParam struct {
	// Array of tag strings to add to the asset.
	AddTags []string `json:"add_tags,omitzero"`
	// Array of tag strings to remove from the asset.
	RemoveTags []string `json:"remove_tags,omitzero"`
	// Array of custom metadata field updates.
	SetMetadata []ExtensionAITasksTaskYesNoOnUnknownSetMetadataParam `json:"set_metadata,omitzero"`
	// Array of custom metadata fields to remove.
	UnsetMetadata []ExtensionAITasksTaskYesNoOnUnknownUnsetMetadataParam `json:"unset_metadata,omitzero"`
	paramObj
}

func (r ExtensionAITasksTaskYesNoOnUnknownParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionAITasksTaskYesNoOnUnknownParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionAITasksTaskYesNoOnUnknownParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Field, Value are required.
type ExtensionAITasksTaskYesNoOnUnknownSetMetadataParam struct {
	// Name of the custom metadata field to set.
	Field string `json:"field,required"`
	// Value to set for the custom metadata field. The value type should match the
	// custom metadata field type.
	Value ExtensionAITasksTaskYesNoOnUnknownSetMetadataValueUnionParam `json:"value,omitzero,required"`
	paramObj
}

func (r ExtensionAITasksTaskYesNoOnUnknownSetMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionAITasksTaskYesNoOnUnknownSetMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionAITasksTaskYesNoOnUnknownSetMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtensionAITasksTaskYesNoOnUnknownSetMetadataValueUnionParam struct {
	OfString param.Opt[string]                                                       `json:",omitzero,inline"`
	OfFloat  param.Opt[float64]                                                      `json:",omitzero,inline"`
	OfBool   param.Opt[bool]                                                         `json:",omitzero,inline"`
	OfMixed  []ExtensionAITasksTaskYesNoOnUnknownSetMetadataValueMixedItemUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u ExtensionAITasksTaskYesNoOnUnknownSetMetadataValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfMixed)
}
func (u *ExtensionAITasksTaskYesNoOnUnknownSetMetadataValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtensionAITasksTaskYesNoOnUnknownSetMetadataValueUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfMixed) {
		return &u.OfMixed
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtensionAITasksTaskYesNoOnUnknownSetMetadataValueMixedItemUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u ExtensionAITasksTaskYesNoOnUnknownSetMetadataValueMixedItemUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *ExtensionAITasksTaskYesNoOnUnknownSetMetadataValueMixedItemUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtensionAITasksTaskYesNoOnUnknownSetMetadataValueMixedItemUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// The property Field is required.
type ExtensionAITasksTaskYesNoOnUnknownUnsetMetadataParam struct {
	// Name of the custom metadata field to remove.
	Field string `json:"field,required"`
	paramObj
}

func (r ExtensionAITasksTaskYesNoOnUnknownUnsetMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionAITasksTaskYesNoOnUnknownUnsetMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionAITasksTaskYesNoOnUnknownUnsetMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Actions to execute if the AI answers yes.
type ExtensionAITasksTaskYesNoOnYesParam struct {
	// Array of tag strings to add to the asset.
	AddTags []string `json:"add_tags,omitzero"`
	// Array of tag strings to remove from the asset.
	RemoveTags []string `json:"remove_tags,omitzero"`
	// Array of custom metadata field updates.
	SetMetadata []ExtensionAITasksTaskYesNoOnYesSetMetadataParam `json:"set_metadata,omitzero"`
	// Array of custom metadata fields to remove.
	UnsetMetadata []ExtensionAITasksTaskYesNoOnYesUnsetMetadataParam `json:"unset_metadata,omitzero"`
	paramObj
}

func (r ExtensionAITasksTaskYesNoOnYesParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionAITasksTaskYesNoOnYesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionAITasksTaskYesNoOnYesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Field, Value are required.
type ExtensionAITasksTaskYesNoOnYesSetMetadataParam struct {
	// Name of the custom metadata field to set.
	Field string `json:"field,required"`
	// Value to set for the custom metadata field. The value type should match the
	// custom metadata field type.
	Value ExtensionAITasksTaskYesNoOnYesSetMetadataValueUnionParam `json:"value,omitzero,required"`
	paramObj
}

func (r ExtensionAITasksTaskYesNoOnYesSetMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionAITasksTaskYesNoOnYesSetMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionAITasksTaskYesNoOnYesSetMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtensionAITasksTaskYesNoOnYesSetMetadataValueUnionParam struct {
	OfString param.Opt[string]                                                   `json:",omitzero,inline"`
	OfFloat  param.Opt[float64]                                                  `json:",omitzero,inline"`
	OfBool   param.Opt[bool]                                                     `json:",omitzero,inline"`
	OfMixed  []ExtensionAITasksTaskYesNoOnYesSetMetadataValueMixedItemUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u ExtensionAITasksTaskYesNoOnYesSetMetadataValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfMixed)
}
func (u *ExtensionAITasksTaskYesNoOnYesSetMetadataValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtensionAITasksTaskYesNoOnYesSetMetadataValueUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfMixed) {
		return &u.OfMixed
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtensionAITasksTaskYesNoOnYesSetMetadataValueMixedItemUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u ExtensionAITasksTaskYesNoOnYesSetMetadataValueMixedItemUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *ExtensionAITasksTaskYesNoOnYesSetMetadataValueMixedItemUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtensionAITasksTaskYesNoOnYesSetMetadataValueMixedItemUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// The property Field is required.
type ExtensionAITasksTaskYesNoOnYesUnsetMetadataParam struct {
	// Name of the custom metadata field to remove.
	Field string `json:"field,required"`
	paramObj
}

func (r ExtensionAITasksTaskYesNoOnYesUnsetMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionAITasksTaskYesNoOnYesUnsetMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionAITasksTaskYesNoOnYesUnsetMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Name are required.
type ExtensionSavedExtensionParam struct {
	// The unique ID of the saved extension to apply.
	ID string `json:"id,required"`
	// Indicates this is a reference to a saved extension.
	//
	// This field can be elided, and will marshal its zero value as "saved-extension".
	Name constant.SavedExtension `json:"name,required"`
	paramObj
}

func (r ExtensionSavedExtensionParam) MarshalJSON() (data []byte, err error) {
	type shadow ExtensionSavedExtensionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtensionSavedExtensionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Options for generating responsive image attributes including `src`, `srcSet`,
// and `sizes` for HTML `<img>` elements. This schema extends `SrcOptions` to add
// support for responsive image generation with breakpoints.
type GetImageAttributesOptionsParam struct {
	// Custom list of **device-width breakpoints** in pixels. These define common
	// screen widths for responsive image generation.
	//
	// Defaults to `[640, 750, 828, 1080, 1200, 1920, 2048, 3840]`. Sorted
	// automatically.
	DeviceBreakpoints []float64 `json:"deviceBreakpoints,omitzero"`
	// Custom list of **image-specific breakpoints** in pixels. Useful for generating
	// small variants (e.g., placeholders or thumbnails).
	//
	// Merged with `deviceBreakpoints` before calculating `srcSet`. Defaults to
	// `[16, 32, 48, 64, 96, 128, 256, 384]`. Sorted automatically.
	ImageBreakpoints []float64 `json:"imageBreakpoints,omitzero"`
	// The value for the HTML `sizes` attribute (e.g., `"100vw"` or
	// `"(min-width:768px) 50vw, 100vw"`).
	//
	//   - If it includes one or more `vw` units, breakpoints smaller than the
	//     corresponding percentage of the smallest device width are excluded.
	//   - If it contains no `vw` units, the full breakpoint list is used.
	//
	// Enables a width-based strategy and generates `w` descriptors in `srcSet`.
	Sizes param.Opt[string] `json:"sizes,omitzero"`
	// The intended display width of the image in pixels, used **only when the `sizes`
	// attribute is not provided**.
	//
	// Triggers a DPR-based strategy (1x and 2x variants) and generates `x` descriptors
	// in `srcSet`.
	//
	// Ignored if `sizes` is present.
	Width param.Opt[float64] `json:"width,omitzero"`
	SrcOptionsParam
}

func (r GetImageAttributesOptionsParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*GetImageAttributesOptionsParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
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
	//
	// Regardless of the encoding method:
	//
	//   - Leading and trailing slashes are removed.
	//   - Remaining slashes within the path are replaced with `@@` when using plain
	//     text.
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
	type shadow struct {
		*ImageOverlayParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
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
func (u OverlayUnionParam) GetLayerMode() *string {
	if vt := u.OfText; vt != nil {
		return (*string)(&vt.LayerMode)
	} else if vt := u.OfImage; vt != nil {
		return (*string)(&vt.LayerMode)
	} else if vt := u.OfVideo; vt != nil {
		return (*string)(&vt.LayerMode)
	} else if vt := u.OfSubtitle; vt != nil {
		return (*string)(&vt.LayerMode)
	} else if vt := u.OfSolidColor; vt != nil {
		return (*string)(&vt.LayerMode)
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

// Resulting set of attributes suitable for an HTML `<img>` element. Useful for
// enabling responsive image loading with `srcSet` and `sizes`.
//
// The property Src is required.
type ResponsiveImageAttributesParam struct {
	// URL for the _largest_ candidate (assigned to plain `src`).
	Src string `json:"src,required" format:"uri"`
	// `sizes` returned (or synthesised as `100vw`). The value for the HTML `sizes`
	// attribute.
	Sizes param.Opt[string] `json:"sizes,omitzero"`
	// Candidate set with `w` or `x` descriptors. Multiple image URLs separated by
	// commas, each with a descriptor.
	SrcSet param.Opt[string] `json:"srcSet,omitzero"`
	// Width as a number (if `width` was provided in the input options).
	Width param.Opt[float64] `json:"width,omitzero"`
	paramObj
}

func (r ResponsiveImageAttributesParam) MarshalJSON() (data []byte, err error) {
	type shadow ResponsiveImageAttributesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponsiveImageAttributesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Saved extension object containing extension configuration.
type SavedExtension struct {
	// Unique identifier of the saved extension.
	ID string `json:"id"`
	// Configuration object for an extension (base extensions only, not saved extension
	// references).
	Config ExtensionConfigUnion `json:"config"`
	// Timestamp when the saved extension was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Description of the saved extension.
	Description string `json:"description"`
	// Name of the saved extension.
	Name string `json:"name"`
	// Timestamp when the saved extension was last updated.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Config      respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SavedExtension) RawJSON() string { return r.JSON.raw }
func (r *SavedExtension) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SavedExtension to a SavedExtensionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SavedExtensionParam.Overrides()
func (r SavedExtension) ToParam() SavedExtensionParam {
	return param.Override[SavedExtensionParam](json.RawMessage(r.RawJSON()))
}

// Saved extension object containing extension configuration.
type SavedExtensionParam struct {
	// Unique identifier of the saved extension.
	ID param.Opt[string] `json:"id,omitzero"`
	// Timestamp when the saved extension was created.
	CreatedAt param.Opt[time.Time] `json:"createdAt,omitzero" format:"date-time"`
	// Description of the saved extension.
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the saved extension.
	Name param.Opt[string] `json:"name,omitzero"`
	// Timestamp when the saved extension was last updated.
	UpdatedAt param.Opt[time.Time] `json:"updatedAt,omitzero" format:"date-time"`
	// Configuration object for an extension (base extensions only, not saved extension
	// references).
	Config ExtensionConfigUnionParam `json:"config,omitzero"`
	paramObj
}

func (r SavedExtensionParam) MarshalJSON() (data []byte, err error) {
	type shadow SavedExtensionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SavedExtensionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	type shadow struct {
		*SolidColorOverlayParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

type SolidColorOverlayTransformationParam struct {
	// Specifies the transparency level of the overlaid solid color layer. Supports
	// integers from `1` to `9`.
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
	// Specifies the corner radius of the solid color overlay.
	//
	//   - Single value (positive integer): Applied to all corners (e.g., `20`).
	//   - `max`: Creates a circular or oval shape.
	//   - Per-corner array: Provide four underscore-separated values representing
	//     top-left, top-right, bottom-right, and bottom-left corners respectively (e.g.,
	//     `10_20_30_40`). See
	//     [Radius](https://imagekit.io/docs/effects-and-enhancements#radius---r).
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
	OfMax    constant.Max      `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u SolidColorOverlayTransformationRadiusUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfMax, u.OfString)
}
func (u *SolidColorOverlayTransformationRadiusUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SolidColorOverlayTransformationRadiusUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfMax) {
		return &u.OfMax
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
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
	//
	// Regardless of the encoding method:
	//
	//   - Leading and trailing slashes are removed.
	//   - Remaining slashes within the path are replaced with `@@` when using plain
	//     text.
	Encoding string `json:"encoding,omitzero"`
	// Control styling of the subtitle. See
	// [Styling subtitles](https://imagekit.io/docs/add-overlays-on-videos#styling-controls-for-subtitles-layer).
	Transformation []SubtitleOverlayTransformationParam `json:"transformation,omitzero"`
	BaseOverlayParam
}

func (r SubtitleOverlayParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*SubtitleOverlayParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
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
	// Sets the font family of subtitle text. Refer to the
	// [supported fonts documented](https://imagekit.io/docs/add-overlays-on-images#supported-text-font-list)
	// in the ImageKit transformations guide.
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
	//
	// Regardless of the encoding method, the input text is always percent-encoded to
	// ensure it is URL-safe.
	Encoding string `json:"encoding,omitzero"`
	// Control styling of the text overlay. See
	// [Text overlays](https://imagekit.io/docs/add-overlays-on-images#text-overlay).
	Transformation []TextOverlayTransformationParam `json:"transformation,omitzero"`
	BaseOverlayParam
}

func (r TextOverlayParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*TextOverlayParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
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
	// Flip/mirror the text horizontally, vertically, or in both directions. Acceptable
	// values: `h` (horizontal), `v` (vertical), `h_v` (horizontal and vertical), or
	// `v_h`.
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
	// Specifies the line height for multi-line text overlays. It will come into effect
	// only if the text wraps over multiple lines. Accepts either an integer value or
	// an arithmetic expression.
	LineHeight TextOverlayTransformationLineHeightUnionParam `json:"lineHeight,omitzero"`
	// Specifies the padding around the overlaid text. Can be provided as a single
	// positive integer or multiple values separated by underscores (following CSS
	// shorthand order). Arithmetic expressions are also accepted.
	Padding TextOverlayTransformationPaddingUnionParam `json:"padding,omitzero"`
	// Specifies the corner radius:
	//
	//   - Single value (positive integer): Applied to all corners (e.g., `20`).
	//   - `max`: Creates a circular or oval shape.
	//   - Per-corner array: Provide four underscore-separated values representing
	//     top-left, top-right, bottom-right, and bottom-left corners respectively (e.g.,
	//     `10_20_30_40`). See
	//     [Radius](https://imagekit.io/docs/effects-and-enhancements#radius---r).
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

// Flip/mirror the text horizontally, vertically, or in both directions. Acceptable
// values: `h` (horizontal), `v` (vertical), `h_v` (horizontal and vertical), or
// `v_h`.
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
	OfMax    constant.Max      `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u TextOverlayTransformationRadiusUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfMax, u.OfString)
}
func (u *TextOverlayTransformationRadiusUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TextOverlayTransformationRadiusUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfMax) {
		return &u.OfMax
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
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
	//   - Dominant color: `dominant` extracts the dominant color from the image. See
	//     [Dominant color background](https://imagekit.io/docs/effects-and-enhancements#dominant-color-background).
	//   - Gradient: `gradient_dominant` or `gradient_dominant_2` creates a gradient
	//     using the dominant colors. Optionally specify palette size (2 or 4), e.g.,
	//     `gradient_dominant_4`. See
	//     [Gradient background](https://imagekit.io/docs/effects-and-enhancements#gradient-background).
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
	// Replaces colors in the image. Supports three formats:
	//
	//   - `toColor` - Replace dominant color with the specified color.
	//   - `toColor_tolerance` - Replace dominant color with specified tolerance (0-100).
	//   - `toColor_tolerance_fromColor` - Replace a specific color with another within
	//     tolerance range. Colors can be hex codes (e.g., `FF0022`) or names (e.g.,
	//     `red`, `blue`). See
	//     [Color replacement](https://imagekit.io/docs/effects-and-enhancements#color-replace---cr).
	ColorReplace param.Opt[string] `json:"colorReplace,omitzero"`
	// Specifies a fallback image if the resource is not found, e.g., a URL or file
	// path. See
	// [Default image](https://imagekit.io/docs/image-transformation#default-image---di).
	DefaultImage param.Opt[string] `json:"defaultImage,omitzero"`
	// Distorts the shape of an image. Supports two modes:
	//
	//   - Perspective distortion: `p-x1_y1_x2_y2_x3_y3_x4_y4` changes the position of
	//     the four corners starting clockwise from top-left.
	//   - Arc distortion: `a-degrees` curves the image upwards (positive values) or
	//     downwards (negative values). See
	//     [Distort effect](https://imagekit.io/docs/effects-and-enhancements#distort---e-distort).
	Distort param.Opt[string] `json:"distort,omitzero"`
	// Accepts values between 0.1 and 5, or `auto` for automatic device pixel ratio
	// (DPR) calculation. Also accepts arithmetic expressions.
	//
	//   - Learn about
	//     [Arithmetic expressions](https://imagekit.io/docs/arithmetic-expressions-in-transformations).
	//   - See [DPR](https://imagekit.io/docs/image-resize-and-crop#dpr---dpr).
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
	// Specifies the corner radius for rounded corners.
	//
	//   - Single value (positive integer): Applied to all corners (e.g., `20`).
	//   - `max`: Creates a circular or oval shape.
	//   - Per-corner array: Provide four underscore-separated values representing
	//     top-left, top-right, bottom-right, and bottom-left corners respectively (e.g.,
	//     `10_20_30_40`). See
	//     [Radius](https://imagekit.io/docs/effects-and-enhancements#radius---r).
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
	OfMax    constant.Max      `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u TransformationRadiusUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfMax, u.OfString)
}
func (u *TransformationRadiusUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TransformationRadiusUnionParam) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfMax) {
		return &u.OfMax
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
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
	//
	// Regardless of the encoding method:
	//
	//   - Leading and trailing slashes are removed.
	//   - Remaining slashes within the path are replaced with `@@` when using plain
	//     text.
	Encoding string `json:"encoding,omitzero"`
	// Array of transformation to be applied to the overlay video. Except
	// `streamingResolutions`, all other video transformations are supported. See
	// [Video transformations](https://imagekit.io/docs/video-transformation).
	Transformation []TransformationParam `json:"transformation,omitzero"`
	BaseOverlayParam
}

func (r VideoOverlayParam) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*VideoOverlayParam
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}
