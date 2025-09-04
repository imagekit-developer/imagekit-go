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
