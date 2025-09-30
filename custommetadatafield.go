// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/imagekit-go/internal/apijson"
	"github.com/stainless-sdks/imagekit-go/internal/apiquery"
	"github.com/stainless-sdks/imagekit-go/internal/requestconfig"
	"github.com/stainless-sdks/imagekit-go/option"
	"github.com/stainless-sdks/imagekit-go/packages/param"
	"github.com/stainless-sdks/imagekit-go/packages/respjson"
)

// CustomMetadataFieldService contains methods and other services that help with
// interacting with the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomMetadataFieldService] method instead.
type CustomMetadataFieldService struct {
	Options []option.RequestOption
}

// NewCustomMetadataFieldService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCustomMetadataFieldService(opts ...option.RequestOption) (r CustomMetadataFieldService) {
	r = CustomMetadataFieldService{}
	r.Options = opts
	return
}

// This API creates a new custom metadata field. Once a custom metadata field is
// created either through this API or using the dashboard UI, its value can be set
// on the assets. The value of a field for an asset can be set using the media
// library UI or programmatically through upload or update assets API.
func (r *CustomMetadataFieldService) New(ctx context.Context, body CustomMetadataFieldNewParams, opts ...option.RequestOption) (res *CustomMetadataField, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/customMetadataFields"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This API updates the label or schema of an existing custom metadata field.
func (r *CustomMetadataFieldService) Update(ctx context.Context, id string, body CustomMetadataFieldUpdateParams, opts ...option.RequestOption) (res *CustomMetadataField, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/customMetadataFields/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// This API returns the array of created custom metadata field objects. By default
// the API returns only non deleted field objects, but you can include deleted
// fields in the API response.
//
// You can also filter results by a specific folder path to retrieve custom
// metadata fields applicable at that location. This path-specific filtering is
// useful when using the **Path policy** feature to determine which custom metadata
// fields are selected for a given path.
func (r *CustomMetadataFieldService) List(ctx context.Context, query CustomMetadataFieldListParams, opts ...option.RequestOption) (res *[]CustomMetadataField, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/customMetadataFields"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// This API deletes a custom metadata field. Even after deleting a custom metadata
// field, you cannot create any new custom metadata field with the same name.
func (r *CustomMetadataFieldService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *CustomMetadataFieldDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/customMetadataFields/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Object containing details of a custom metadata field.
type CustomMetadataField struct {
	// Unique identifier for the custom metadata field. Use this to update the field.
	ID string `json:"id,required"`
	// Human readable name of the custom metadata field. This name is displayed as form
	// field label to the users while setting field value on the asset in the media
	// library UI.
	Label string `json:"label,required"`
	// API name of the custom metadata field. This becomes the key while setting
	// `customMetadata` (key-value object) for an asset using upload or update API.
	Name string `json:"name,required"`
	// An object that describes the rules for the custom metadata field value.
	Schema CustomMetadataFieldSchema `json:"schema,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Label       respjson.Field
		Name        respjson.Field
		Schema      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomMetadataField) RawJSON() string { return r.JSON.raw }
func (r *CustomMetadataField) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object that describes the rules for the custom metadata field value.
type CustomMetadataFieldSchema struct {
	// Type of the custom metadata field.
	//
	// Any of "Text", "Textarea", "Number", "Date", "Boolean", "SingleSelect",
	// "MultiSelect".
	Type string `json:"type,required"`
	// The default value for this custom metadata field. Data type of default value
	// depends on the field type.
	DefaultValue CustomMetadataFieldSchemaDefaultValueUnion `json:"defaultValue"`
	// Specifies if the this custom metadata field is required or not.
	IsValueRequired bool `json:"isValueRequired"`
	// Maximum length of string. Only set if `type` is set to `Text` or `Textarea`.
	MaxLength float64 `json:"maxLength"`
	// Maximum value of the field. Only set if field type is `Date` or `Number`. For
	// `Date` type field, the value will be in ISO8601 string format. For `Number` type
	// field, it will be a numeric value.
	MaxValue CustomMetadataFieldSchemaMaxValueUnion `json:"maxValue"`
	// Minimum length of string. Only set if `type` is set to `Text` or `Textarea`.
	MinLength float64 `json:"minLength"`
	// Minimum value of the field. Only set if field type is `Date` or `Number`. For
	// `Date` type field, the value will be in ISO8601 string format. For `Number` type
	// field, it will be a numeric value.
	MinValue CustomMetadataFieldSchemaMinValueUnion `json:"minValue"`
	// An array of allowed values when field type is `SingleSelect` or `MultiSelect`.
	SelectOptions []CustomMetadataFieldSchemaSelectOptionUnion `json:"selectOptions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type            respjson.Field
		DefaultValue    respjson.Field
		IsValueRequired respjson.Field
		MaxLength       respjson.Field
		MaxValue        respjson.Field
		MinLength       respjson.Field
		MinValue        respjson.Field
		SelectOptions   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomMetadataFieldSchema) RawJSON() string { return r.JSON.raw }
func (r *CustomMetadataFieldSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CustomMetadataFieldSchemaDefaultValueUnion contains all possible properties and
// values from [string], [float64], [bool],
// [[]CustomMetadataFieldSchemaDefaultValueMixedItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfMixed]
type CustomMetadataFieldSchemaDefaultValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]CustomMetadataFieldSchemaDefaultValueMixedItemUnion] instead of an object.
	OfMixed []CustomMetadataFieldSchemaDefaultValueMixedItemUnion `json:",inline"`
	JSON    struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		OfMixed  respjson.Field
		raw      string
	} `json:"-"`
}

func (u CustomMetadataFieldSchemaDefaultValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldSchemaDefaultValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldSchemaDefaultValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldSchemaDefaultValueUnion) AsMixed() (v []CustomMetadataFieldSchemaDefaultValueMixedItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CustomMetadataFieldSchemaDefaultValueUnion) RawJSON() string { return u.JSON.raw }

func (r *CustomMetadataFieldSchemaDefaultValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CustomMetadataFieldSchemaDefaultValueMixedItemUnion contains all possible
// properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type CustomMetadataFieldSchemaDefaultValueMixedItemUnion struct {
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

func (u CustomMetadataFieldSchemaDefaultValueMixedItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldSchemaDefaultValueMixedItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldSchemaDefaultValueMixedItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CustomMetadataFieldSchemaDefaultValueMixedItemUnion) RawJSON() string { return u.JSON.raw }

func (r *CustomMetadataFieldSchemaDefaultValueMixedItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CustomMetadataFieldSchemaMaxValueUnion contains all possible properties and
// values from [string], [float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat]
type CustomMetadataFieldSchemaMaxValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	JSON    struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		raw      string
	} `json:"-"`
}

func (u CustomMetadataFieldSchemaMaxValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldSchemaMaxValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CustomMetadataFieldSchemaMaxValueUnion) RawJSON() string { return u.JSON.raw }

func (r *CustomMetadataFieldSchemaMaxValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CustomMetadataFieldSchemaMinValueUnion contains all possible properties and
// values from [string], [float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat]
type CustomMetadataFieldSchemaMinValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	JSON    struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		raw      string
	} `json:"-"`
}

func (u CustomMetadataFieldSchemaMinValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldSchemaMinValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CustomMetadataFieldSchemaMinValueUnion) RawJSON() string { return u.JSON.raw }

func (r *CustomMetadataFieldSchemaMinValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CustomMetadataFieldSchemaSelectOptionUnion contains all possible properties and
// values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type CustomMetadataFieldSchemaSelectOptionUnion struct {
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

func (u CustomMetadataFieldSchemaSelectOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldSchemaSelectOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldSchemaSelectOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CustomMetadataFieldSchemaSelectOptionUnion) RawJSON() string { return u.JSON.raw }

func (r *CustomMetadataFieldSchemaSelectOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomMetadataFieldDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomMetadataFieldDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomMetadataFieldDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomMetadataFieldNewParams struct {
	// Human readable name of the custom metadata field. This should be unique across
	// all non deleted custom metadata fields. This name is displayed as form field
	// label to the users while setting field value on an asset in the media library
	// UI.
	Label string `json:"label,required"`
	// API name of the custom metadata field. This should be unique across all
	// (including deleted) custom metadata fields.
	Name   string                             `json:"name,required"`
	Schema CustomMetadataFieldNewParamsSchema `json:"schema,omitzero,required"`
	paramObj
}

func (r CustomMetadataFieldNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomMetadataFieldNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomMetadataFieldNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type CustomMetadataFieldNewParamsSchema struct {
	// Type of the custom metadata field.
	//
	// Any of "Text", "Textarea", "Number", "Date", "Boolean", "SingleSelect",
	// "MultiSelect".
	Type string `json:"type,omitzero,required"`
	// Sets this custom metadata field as required. Setting custom metadata fields on
	// an asset will throw error if the value for all required fields are not present
	// in upload or update asset API request body.
	IsValueRequired param.Opt[bool] `json:"isValueRequired,omitzero"`
	// Maximum length of string. Only set this property if `type` is set to `Text` or
	// `Textarea`.
	MaxLength param.Opt[float64] `json:"maxLength,omitzero"`
	// Minimum length of string. Only set this property if `type` is set to `Text` or
	// `Textarea`.
	MinLength param.Opt[float64] `json:"minLength,omitzero"`
	// The default value for this custom metadata field. This property is only required
	// if `isValueRequired` property is set to `true`. The value should match the
	// `type` of custom metadata field.
	DefaultValue CustomMetadataFieldNewParamsSchemaDefaultValueUnion `json:"defaultValue,omitzero"`
	// Maximum value of the field. Only set this property if field type is `Date` or
	// `Number`. For `Date` type field, set the minimum date in ISO8601 string format.
	// For `Number` type field, set the minimum numeric value.
	MaxValue CustomMetadataFieldNewParamsSchemaMaxValueUnion `json:"maxValue,omitzero"`
	// Minimum value of the field. Only set this property if field type is `Date` or
	// `Number`. For `Date` type field, set the minimum date in ISO8601 string format.
	// For `Number` type field, set the minimum numeric value.
	MinValue CustomMetadataFieldNewParamsSchemaMinValueUnion `json:"minValue,omitzero"`
	// An array of allowed values. This property is only required if `type` property is
	// set to `SingleSelect` or `MultiSelect`.
	SelectOptions []CustomMetadataFieldNewParamsSchemaSelectOptionUnion `json:"selectOptions,omitzero"`
	paramObj
}

func (r CustomMetadataFieldNewParamsSchema) MarshalJSON() (data []byte, err error) {
	type shadow CustomMetadataFieldNewParamsSchema
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomMetadataFieldNewParamsSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CustomMetadataFieldNewParamsSchema](
		"type", "Text", "Textarea", "Number", "Date", "Boolean", "SingleSelect", "MultiSelect",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CustomMetadataFieldNewParamsSchemaDefaultValueUnion struct {
	OfString param.Opt[string]                                              `json:",omitzero,inline"`
	OfFloat  param.Opt[float64]                                             `json:",omitzero,inline"`
	OfBool   param.Opt[bool]                                                `json:",omitzero,inline"`
	OfMixed  []CustomMetadataFieldNewParamsSchemaDefaultValueMixedItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u CustomMetadataFieldNewParamsSchemaDefaultValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfMixed)
}
func (u *CustomMetadataFieldNewParamsSchemaDefaultValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CustomMetadataFieldNewParamsSchemaDefaultValueUnion) asAny() any {
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
type CustomMetadataFieldNewParamsSchemaDefaultValueMixedItemUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u CustomMetadataFieldNewParamsSchemaDefaultValueMixedItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *CustomMetadataFieldNewParamsSchemaDefaultValueMixedItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CustomMetadataFieldNewParamsSchemaDefaultValueMixedItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CustomMetadataFieldNewParamsSchemaMaxValueUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	paramUnion
}

func (u CustomMetadataFieldNewParamsSchemaMaxValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat)
}
func (u *CustomMetadataFieldNewParamsSchemaMaxValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CustomMetadataFieldNewParamsSchemaMaxValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CustomMetadataFieldNewParamsSchemaMinValueUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	paramUnion
}

func (u CustomMetadataFieldNewParamsSchemaMinValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat)
}
func (u *CustomMetadataFieldNewParamsSchemaMinValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CustomMetadataFieldNewParamsSchemaMinValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CustomMetadataFieldNewParamsSchemaSelectOptionUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u CustomMetadataFieldNewParamsSchemaSelectOptionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *CustomMetadataFieldNewParamsSchemaSelectOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CustomMetadataFieldNewParamsSchemaSelectOptionUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CustomMetadataFieldUpdateParams struct {
	// Human readable name of the custom metadata field. This should be unique across
	// all non deleted custom metadata fields. This name is displayed as form field
	// label to the users while setting field value on an asset in the media library
	// UI. This parameter is required if `schema` is not provided.
	Label param.Opt[string] `json:"label,omitzero"`
	// An object that describes the rules for the custom metadata key. This parameter
	// is required if `label` is not provided. Note: `type` cannot be updated and will
	// be ignored if sent with the `schema`. The schema will be validated as per the
	// existing `type`.
	Schema CustomMetadataFieldUpdateParamsSchema `json:"schema,omitzero"`
	paramObj
}

func (r CustomMetadataFieldUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomMetadataFieldUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomMetadataFieldUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object that describes the rules for the custom metadata key. This parameter
// is required if `label` is not provided. Note: `type` cannot be updated and will
// be ignored if sent with the `schema`. The schema will be validated as per the
// existing `type`.
type CustomMetadataFieldUpdateParamsSchema struct {
	// Sets this custom metadata field as required. Setting custom metadata fields on
	// an asset will throw error if the value for all required fields are not present
	// in upload or update asset API request body.
	IsValueRequired param.Opt[bool] `json:"isValueRequired,omitzero"`
	// Maximum length of string. Only set this property if `type` is set to `Text` or
	// `Textarea`.
	MaxLength param.Opt[float64] `json:"maxLength,omitzero"`
	// Minimum length of string. Only set this property if `type` is set to `Text` or
	// `Textarea`.
	MinLength param.Opt[float64] `json:"minLength,omitzero"`
	// The default value for this custom metadata field. This property is only required
	// if `isValueRequired` property is set to `true`. The value should match the
	// `type` of custom metadata field.
	DefaultValue CustomMetadataFieldUpdateParamsSchemaDefaultValueUnion `json:"defaultValue,omitzero"`
	// Maximum value of the field. Only set this property if field type is `Date` or
	// `Number`. For `Date` type field, set the minimum date in ISO8601 string format.
	// For `Number` type field, set the minimum numeric value.
	MaxValue CustomMetadataFieldUpdateParamsSchemaMaxValueUnion `json:"maxValue,omitzero"`
	// Minimum value of the field. Only set this property if field type is `Date` or
	// `Number`. For `Date` type field, set the minimum date in ISO8601 string format.
	// For `Number` type field, set the minimum numeric value.
	MinValue CustomMetadataFieldUpdateParamsSchemaMinValueUnion `json:"minValue,omitzero"`
	// An array of allowed values. This property is only required if `type` property is
	// set to `SingleSelect` or `MultiSelect`.
	SelectOptions []CustomMetadataFieldUpdateParamsSchemaSelectOptionUnion `json:"selectOptions,omitzero"`
	paramObj
}

func (r CustomMetadataFieldUpdateParamsSchema) MarshalJSON() (data []byte, err error) {
	type shadow CustomMetadataFieldUpdateParamsSchema
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomMetadataFieldUpdateParamsSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CustomMetadataFieldUpdateParamsSchemaDefaultValueUnion struct {
	OfString param.Opt[string]                                                 `json:",omitzero,inline"`
	OfFloat  param.Opt[float64]                                                `json:",omitzero,inline"`
	OfBool   param.Opt[bool]                                                   `json:",omitzero,inline"`
	OfMixed  []CustomMetadataFieldUpdateParamsSchemaDefaultValueMixedItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u CustomMetadataFieldUpdateParamsSchemaDefaultValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool, u.OfMixed)
}
func (u *CustomMetadataFieldUpdateParamsSchemaDefaultValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CustomMetadataFieldUpdateParamsSchemaDefaultValueUnion) asAny() any {
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
type CustomMetadataFieldUpdateParamsSchemaDefaultValueMixedItemUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u CustomMetadataFieldUpdateParamsSchemaDefaultValueMixedItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *CustomMetadataFieldUpdateParamsSchemaDefaultValueMixedItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CustomMetadataFieldUpdateParamsSchemaDefaultValueMixedItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CustomMetadataFieldUpdateParamsSchemaMaxValueUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	paramUnion
}

func (u CustomMetadataFieldUpdateParamsSchemaMaxValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat)
}
func (u *CustomMetadataFieldUpdateParamsSchemaMaxValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CustomMetadataFieldUpdateParamsSchemaMaxValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CustomMetadataFieldUpdateParamsSchemaMinValueUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	paramUnion
}

func (u CustomMetadataFieldUpdateParamsSchemaMinValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat)
}
func (u *CustomMetadataFieldUpdateParamsSchemaMinValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CustomMetadataFieldUpdateParamsSchemaMinValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CustomMetadataFieldUpdateParamsSchemaSelectOptionUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u CustomMetadataFieldUpdateParamsSchemaSelectOptionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *CustomMetadataFieldUpdateParamsSchemaSelectOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CustomMetadataFieldUpdateParamsSchemaSelectOptionUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CustomMetadataFieldListParams struct {
	// The folder path (e.g., `/path/to/folder`) for which to retrieve applicable
	// custom metadata fields. Useful for determining path-specific field selections
	// when the [Path policy](https://imagekit.io/docs/dam/path-policy) feature is in
	// use.
	FolderPath param.Opt[string] `query:"folderPath,omitzero" json:"-"`
	// Set it to `true` to include deleted field objects in the API response.
	IncludeDeleted param.Opt[bool] `query:"includeDeleted,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CustomMetadataFieldListParams]'s query parameters as
// `url.Values`.
func (r CustomMetadataFieldListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
