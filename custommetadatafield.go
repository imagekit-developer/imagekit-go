// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

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
func (r *CustomMetadataFieldService) New(ctx context.Context, body CustomMetadataFieldNewParams, opts ...option.RequestOption) (res *CustomMetadataFieldNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/customMetadataFields"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This API updates the label or schema of an existing custom metadata field.
func (r *CustomMetadataFieldService) Update(ctx context.Context, id string, body CustomMetadataFieldUpdateParams, opts ...option.RequestOption) (res *CustomMetadataFieldUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *CustomMetadataFieldService) List(ctx context.Context, query CustomMetadataFieldListParams, opts ...option.RequestOption) (res *[]CustomMetadataFieldListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/customMetadataFields"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// This API deletes a custom metadata field. Even after deleting a custom metadata
// field, you cannot create any new custom metadata field with the same name.
func (r *CustomMetadataFieldService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *CustomMetadataFieldDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/customMetadataFields/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Object containing details of a custom metadata field.
type CustomMetadataFieldNewResponse struct {
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
	Schema CustomMetadataFieldNewResponseSchema `json:"schema,required"`
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
func (r CustomMetadataFieldNewResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomMetadataFieldNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object that describes the rules for the custom metadata field value.
type CustomMetadataFieldNewResponseSchema struct {
	// Type of the custom metadata field.
	//
	// Any of "Text", "Textarea", "Number", "Date", "Boolean", "SingleSelect",
	// "MultiSelect".
	Type string `json:"type,required"`
	// The default value for this custom metadata field. Date type of default value
	// depends on the field type.
	DefaultValue CustomMetadataFieldNewResponseSchemaDefaultValueUnion `json:"defaultValue"`
	// Specifies if the this custom metadata field is required or not.
	IsValueRequired bool `json:"isValueRequired"`
	// Maximum length of string. Only set if `type` is set to `Text` or `Textarea`.
	MaxLength float64 `json:"maxLength"`
	// Maximum value of the field. Only set if field type is `Date` or `Number`. For
	// `Date` type field, the value will be in ISO8601 string format. For `Number` type
	// field, it will be a numeric value.
	MaxValue CustomMetadataFieldNewResponseSchemaMaxValueUnion `json:"maxValue"`
	// Minimum length of string. Only set if `type` is set to `Text` or `Textarea`.
	MinLength float64 `json:"minLength"`
	// Minimum value of the field. Only set if field type is `Date` or `Number`. For
	// `Date` type field, the value will be in ISO8601 string format. For `Number` type
	// field, it will be a numeric value.
	MinValue CustomMetadataFieldNewResponseSchemaMinValueUnion `json:"minValue"`
	// An array of allowed values when field type is `SingleSelect` or `MultiSelect`.
	SelectOptions []CustomMetadataFieldNewResponseSchemaSelectOptionUnion `json:"selectOptions"`
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
func (r CustomMetadataFieldNewResponseSchema) RawJSON() string { return r.JSON.raw }
func (r *CustomMetadataFieldNewResponseSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CustomMetadataFieldNewResponseSchemaDefaultValueUnion contains all possible
// properties and values from
// [CustomMetadataFieldNewResponseSchemaDefaultValueUnionUnion],
// [[]CustomMetadataFieldNewResponseSchemaDefaultValueJsonScalarArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfCustomMetadataFieldNewResponseSchemaDefaultValueJsonScalarArray]
type CustomMetadataFieldNewResponseSchemaDefaultValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]CustomMetadataFieldNewResponseSchemaDefaultValueJsonScalarArrayItemUnion]
	// instead of an object.
	OfCustomMetadataFieldNewResponseSchemaDefaultValueJsonScalarArray []CustomMetadataFieldNewResponseSchemaDefaultValueJsonScalarArrayItemUnion `json:",inline"`
	JSON                                                              struct {
		OfString                                                          respjson.Field
		OfFloat                                                           respjson.Field
		OfBool                                                            respjson.Field
		OfCustomMetadataFieldNewResponseSchemaDefaultValueJsonScalarArray respjson.Field
		raw                                                               string
	} `json:"-"`
}

func (u CustomMetadataFieldNewResponseSchemaDefaultValueUnion) AsCustomMetadataFieldNewResponseSchemaDefaultValueUnion() (v CustomMetadataFieldNewResponseSchemaDefaultValueUnionUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldNewResponseSchemaDefaultValueUnion) AsCustomMetadataFieldNewResponseSchemaDefaultValueJsonScalarArray() (v []CustomMetadataFieldNewResponseSchemaDefaultValueJsonScalarArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CustomMetadataFieldNewResponseSchemaDefaultValueUnion) RawJSON() string { return u.JSON.raw }

func (r *CustomMetadataFieldNewResponseSchemaDefaultValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CustomMetadataFieldNewResponseSchemaDefaultValueUnionUnion contains all possible
// properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type CustomMetadataFieldNewResponseSchemaDefaultValueUnionUnion struct {
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

func (u CustomMetadataFieldNewResponseSchemaDefaultValueUnionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldNewResponseSchemaDefaultValueUnionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldNewResponseSchemaDefaultValueUnionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CustomMetadataFieldNewResponseSchemaDefaultValueUnionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *CustomMetadataFieldNewResponseSchemaDefaultValueUnionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CustomMetadataFieldNewResponseSchemaDefaultValueJsonScalarArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type CustomMetadataFieldNewResponseSchemaDefaultValueJsonScalarArrayItemUnion struct {
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

func (u CustomMetadataFieldNewResponseSchemaDefaultValueJsonScalarArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldNewResponseSchemaDefaultValueJsonScalarArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldNewResponseSchemaDefaultValueJsonScalarArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CustomMetadataFieldNewResponseSchemaDefaultValueJsonScalarArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *CustomMetadataFieldNewResponseSchemaDefaultValueJsonScalarArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CustomMetadataFieldNewResponseSchemaMaxValueUnion contains all possible
// properties and values from [string], [float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat]
type CustomMetadataFieldNewResponseSchemaMaxValueUnion struct {
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

func (u CustomMetadataFieldNewResponseSchemaMaxValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldNewResponseSchemaMaxValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CustomMetadataFieldNewResponseSchemaMaxValueUnion) RawJSON() string { return u.JSON.raw }

func (r *CustomMetadataFieldNewResponseSchemaMaxValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CustomMetadataFieldNewResponseSchemaMinValueUnion contains all possible
// properties and values from [string], [float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat]
type CustomMetadataFieldNewResponseSchemaMinValueUnion struct {
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

func (u CustomMetadataFieldNewResponseSchemaMinValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldNewResponseSchemaMinValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CustomMetadataFieldNewResponseSchemaMinValueUnion) RawJSON() string { return u.JSON.raw }

func (r *CustomMetadataFieldNewResponseSchemaMinValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CustomMetadataFieldNewResponseSchemaSelectOptionUnion contains all possible
// properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type CustomMetadataFieldNewResponseSchemaSelectOptionUnion struct {
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

func (u CustomMetadataFieldNewResponseSchemaSelectOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldNewResponseSchemaSelectOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldNewResponseSchemaSelectOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CustomMetadataFieldNewResponseSchemaSelectOptionUnion) RawJSON() string { return u.JSON.raw }

func (r *CustomMetadataFieldNewResponseSchemaSelectOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Object containing details of a custom metadata field.
type CustomMetadataFieldUpdateResponse struct {
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
	Schema CustomMetadataFieldUpdateResponseSchema `json:"schema,required"`
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
func (r CustomMetadataFieldUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomMetadataFieldUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object that describes the rules for the custom metadata field value.
type CustomMetadataFieldUpdateResponseSchema struct {
	// Type of the custom metadata field.
	//
	// Any of "Text", "Textarea", "Number", "Date", "Boolean", "SingleSelect",
	// "MultiSelect".
	Type string `json:"type,required"`
	// The default value for this custom metadata field. Date type of default value
	// depends on the field type.
	DefaultValue CustomMetadataFieldUpdateResponseSchemaDefaultValueUnion `json:"defaultValue"`
	// Specifies if the this custom metadata field is required or not.
	IsValueRequired bool `json:"isValueRequired"`
	// Maximum length of string. Only set if `type` is set to `Text` or `Textarea`.
	MaxLength float64 `json:"maxLength"`
	// Maximum value of the field. Only set if field type is `Date` or `Number`. For
	// `Date` type field, the value will be in ISO8601 string format. For `Number` type
	// field, it will be a numeric value.
	MaxValue CustomMetadataFieldUpdateResponseSchemaMaxValueUnion `json:"maxValue"`
	// Minimum length of string. Only set if `type` is set to `Text` or `Textarea`.
	MinLength float64 `json:"minLength"`
	// Minimum value of the field. Only set if field type is `Date` or `Number`. For
	// `Date` type field, the value will be in ISO8601 string format. For `Number` type
	// field, it will be a numeric value.
	MinValue CustomMetadataFieldUpdateResponseSchemaMinValueUnion `json:"minValue"`
	// An array of allowed values when field type is `SingleSelect` or `MultiSelect`.
	SelectOptions []CustomMetadataFieldUpdateResponseSchemaSelectOptionUnion `json:"selectOptions"`
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
func (r CustomMetadataFieldUpdateResponseSchema) RawJSON() string { return r.JSON.raw }
func (r *CustomMetadataFieldUpdateResponseSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CustomMetadataFieldUpdateResponseSchemaDefaultValueUnion contains all possible
// properties and values from
// [CustomMetadataFieldUpdateResponseSchemaDefaultValueUnionUnion],
// [[]CustomMetadataFieldUpdateResponseSchemaDefaultValueJsonScalarArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfCustomMetadataFieldUpdateResponseSchemaDefaultValueJsonScalarArray]
type CustomMetadataFieldUpdateResponseSchemaDefaultValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]CustomMetadataFieldUpdateResponseSchemaDefaultValueJsonScalarArrayItemUnion]
	// instead of an object.
	OfCustomMetadataFieldUpdateResponseSchemaDefaultValueJsonScalarArray []CustomMetadataFieldUpdateResponseSchemaDefaultValueJsonScalarArrayItemUnion `json:",inline"`
	JSON                                                                 struct {
		OfString                                                             respjson.Field
		OfFloat                                                              respjson.Field
		OfBool                                                               respjson.Field
		OfCustomMetadataFieldUpdateResponseSchemaDefaultValueJsonScalarArray respjson.Field
		raw                                                                  string
	} `json:"-"`
}

func (u CustomMetadataFieldUpdateResponseSchemaDefaultValueUnion) AsCustomMetadataFieldUpdateResponseSchemaDefaultValueUnion() (v CustomMetadataFieldUpdateResponseSchemaDefaultValueUnionUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldUpdateResponseSchemaDefaultValueUnion) AsCustomMetadataFieldUpdateResponseSchemaDefaultValueJsonScalarArray() (v []CustomMetadataFieldUpdateResponseSchemaDefaultValueJsonScalarArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CustomMetadataFieldUpdateResponseSchemaDefaultValueUnion) RawJSON() string { return u.JSON.raw }

func (r *CustomMetadataFieldUpdateResponseSchemaDefaultValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CustomMetadataFieldUpdateResponseSchemaDefaultValueUnionUnion contains all
// possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type CustomMetadataFieldUpdateResponseSchemaDefaultValueUnionUnion struct {
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

func (u CustomMetadataFieldUpdateResponseSchemaDefaultValueUnionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldUpdateResponseSchemaDefaultValueUnionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldUpdateResponseSchemaDefaultValueUnionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CustomMetadataFieldUpdateResponseSchemaDefaultValueUnionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *CustomMetadataFieldUpdateResponseSchemaDefaultValueUnionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CustomMetadataFieldUpdateResponseSchemaDefaultValueJsonScalarArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type CustomMetadataFieldUpdateResponseSchemaDefaultValueJsonScalarArrayItemUnion struct {
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

func (u CustomMetadataFieldUpdateResponseSchemaDefaultValueJsonScalarArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldUpdateResponseSchemaDefaultValueJsonScalarArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldUpdateResponseSchemaDefaultValueJsonScalarArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CustomMetadataFieldUpdateResponseSchemaDefaultValueJsonScalarArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *CustomMetadataFieldUpdateResponseSchemaDefaultValueJsonScalarArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CustomMetadataFieldUpdateResponseSchemaMaxValueUnion contains all possible
// properties and values from [string], [float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat]
type CustomMetadataFieldUpdateResponseSchemaMaxValueUnion struct {
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

func (u CustomMetadataFieldUpdateResponseSchemaMaxValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldUpdateResponseSchemaMaxValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CustomMetadataFieldUpdateResponseSchemaMaxValueUnion) RawJSON() string { return u.JSON.raw }

func (r *CustomMetadataFieldUpdateResponseSchemaMaxValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CustomMetadataFieldUpdateResponseSchemaMinValueUnion contains all possible
// properties and values from [string], [float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat]
type CustomMetadataFieldUpdateResponseSchemaMinValueUnion struct {
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

func (u CustomMetadataFieldUpdateResponseSchemaMinValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldUpdateResponseSchemaMinValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CustomMetadataFieldUpdateResponseSchemaMinValueUnion) RawJSON() string { return u.JSON.raw }

func (r *CustomMetadataFieldUpdateResponseSchemaMinValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CustomMetadataFieldUpdateResponseSchemaSelectOptionUnion contains all possible
// properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type CustomMetadataFieldUpdateResponseSchemaSelectOptionUnion struct {
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

func (u CustomMetadataFieldUpdateResponseSchemaSelectOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldUpdateResponseSchemaSelectOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldUpdateResponseSchemaSelectOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CustomMetadataFieldUpdateResponseSchemaSelectOptionUnion) RawJSON() string { return u.JSON.raw }

func (r *CustomMetadataFieldUpdateResponseSchemaSelectOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Object containing details of a custom metadata field.
type CustomMetadataFieldListResponse struct {
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
	Schema CustomMetadataFieldListResponseSchema `json:"schema,required"`
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
func (r CustomMetadataFieldListResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomMetadataFieldListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object that describes the rules for the custom metadata field value.
type CustomMetadataFieldListResponseSchema struct {
	// Type of the custom metadata field.
	//
	// Any of "Text", "Textarea", "Number", "Date", "Boolean", "SingleSelect",
	// "MultiSelect".
	Type string `json:"type,required"`
	// The default value for this custom metadata field. Date type of default value
	// depends on the field type.
	DefaultValue CustomMetadataFieldListResponseSchemaDefaultValueUnion `json:"defaultValue"`
	// Specifies if the this custom metadata field is required or not.
	IsValueRequired bool `json:"isValueRequired"`
	// Maximum length of string. Only set if `type` is set to `Text` or `Textarea`.
	MaxLength float64 `json:"maxLength"`
	// Maximum value of the field. Only set if field type is `Date` or `Number`. For
	// `Date` type field, the value will be in ISO8601 string format. For `Number` type
	// field, it will be a numeric value.
	MaxValue CustomMetadataFieldListResponseSchemaMaxValueUnion `json:"maxValue"`
	// Minimum length of string. Only set if `type` is set to `Text` or `Textarea`.
	MinLength float64 `json:"minLength"`
	// Minimum value of the field. Only set if field type is `Date` or `Number`. For
	// `Date` type field, the value will be in ISO8601 string format. For `Number` type
	// field, it will be a numeric value.
	MinValue CustomMetadataFieldListResponseSchemaMinValueUnion `json:"minValue"`
	// An array of allowed values when field type is `SingleSelect` or `MultiSelect`.
	SelectOptions []CustomMetadataFieldListResponseSchemaSelectOptionUnion `json:"selectOptions"`
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
func (r CustomMetadataFieldListResponseSchema) RawJSON() string { return r.JSON.raw }
func (r *CustomMetadataFieldListResponseSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CustomMetadataFieldListResponseSchemaDefaultValueUnion contains all possible
// properties and values from
// [CustomMetadataFieldListResponseSchemaDefaultValueUnionUnion],
// [[]CustomMetadataFieldListResponseSchemaDefaultValueJsonScalarArrayItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfCustomMetadataFieldListResponseSchemaDefaultValueJsonScalarArray]
type CustomMetadataFieldListResponseSchemaDefaultValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]CustomMetadataFieldListResponseSchemaDefaultValueJsonScalarArrayItemUnion]
	// instead of an object.
	OfCustomMetadataFieldListResponseSchemaDefaultValueJsonScalarArray []CustomMetadataFieldListResponseSchemaDefaultValueJsonScalarArrayItemUnion `json:",inline"`
	JSON                                                               struct {
		OfString                                                           respjson.Field
		OfFloat                                                            respjson.Field
		OfBool                                                             respjson.Field
		OfCustomMetadataFieldListResponseSchemaDefaultValueJsonScalarArray respjson.Field
		raw                                                                string
	} `json:"-"`
}

func (u CustomMetadataFieldListResponseSchemaDefaultValueUnion) AsCustomMetadataFieldListResponseSchemaDefaultValueUnion() (v CustomMetadataFieldListResponseSchemaDefaultValueUnionUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldListResponseSchemaDefaultValueUnion) AsCustomMetadataFieldListResponseSchemaDefaultValueJsonScalarArray() (v []CustomMetadataFieldListResponseSchemaDefaultValueJsonScalarArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CustomMetadataFieldListResponseSchemaDefaultValueUnion) RawJSON() string { return u.JSON.raw }

func (r *CustomMetadataFieldListResponseSchemaDefaultValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CustomMetadataFieldListResponseSchemaDefaultValueUnionUnion contains all
// possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type CustomMetadataFieldListResponseSchemaDefaultValueUnionUnion struct {
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

func (u CustomMetadataFieldListResponseSchemaDefaultValueUnionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldListResponseSchemaDefaultValueUnionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldListResponseSchemaDefaultValueUnionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CustomMetadataFieldListResponseSchemaDefaultValueUnionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *CustomMetadataFieldListResponseSchemaDefaultValueUnionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CustomMetadataFieldListResponseSchemaDefaultValueJsonScalarArrayItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type CustomMetadataFieldListResponseSchemaDefaultValueJsonScalarArrayItemUnion struct {
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

func (u CustomMetadataFieldListResponseSchemaDefaultValueJsonScalarArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldListResponseSchemaDefaultValueJsonScalarArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldListResponseSchemaDefaultValueJsonScalarArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CustomMetadataFieldListResponseSchemaDefaultValueJsonScalarArrayItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *CustomMetadataFieldListResponseSchemaDefaultValueJsonScalarArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CustomMetadataFieldListResponseSchemaMaxValueUnion contains all possible
// properties and values from [string], [float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat]
type CustomMetadataFieldListResponseSchemaMaxValueUnion struct {
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

func (u CustomMetadataFieldListResponseSchemaMaxValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldListResponseSchemaMaxValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CustomMetadataFieldListResponseSchemaMaxValueUnion) RawJSON() string { return u.JSON.raw }

func (r *CustomMetadataFieldListResponseSchemaMaxValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CustomMetadataFieldListResponseSchemaMinValueUnion contains all possible
// properties and values from [string], [float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat]
type CustomMetadataFieldListResponseSchemaMinValueUnion struct {
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

func (u CustomMetadataFieldListResponseSchemaMinValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldListResponseSchemaMinValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CustomMetadataFieldListResponseSchemaMinValueUnion) RawJSON() string { return u.JSON.raw }

func (r *CustomMetadataFieldListResponseSchemaMinValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CustomMetadataFieldListResponseSchemaSelectOptionUnion contains all possible
// properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type CustomMetadataFieldListResponseSchemaSelectOptionUnion struct {
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

func (u CustomMetadataFieldListResponseSchemaSelectOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldListResponseSchemaSelectOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomMetadataFieldListResponseSchemaSelectOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CustomMetadataFieldListResponseSchemaSelectOptionUnion) RawJSON() string { return u.JSON.raw }

func (r *CustomMetadataFieldListResponseSchemaSelectOptionUnion) UnmarshalJSON(data []byte) error {
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
	OfCustomMetadataFieldNewsSchemaDefaultValueUnion           *CustomMetadataFieldNewParamsSchemaDefaultValueUnionUnion                `json:",omitzero,inline"`
	OfCustomMetadataFieldNewsSchemaDefaultValueJsonScalarArray []CustomMetadataFieldNewParamsSchemaDefaultValueJsonScalarArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u CustomMetadataFieldNewParamsSchemaDefaultValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCustomMetadataFieldNewsSchemaDefaultValueUnion, u.OfCustomMetadataFieldNewsSchemaDefaultValueJsonScalarArray)
}
func (u *CustomMetadataFieldNewParamsSchemaDefaultValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CustomMetadataFieldNewParamsSchemaDefaultValueUnion) asAny() any {
	if !param.IsOmitted(u.OfCustomMetadataFieldNewsSchemaDefaultValueUnion) {
		return u.OfCustomMetadataFieldNewsSchemaDefaultValueUnion.asAny()
	} else if !param.IsOmitted(u.OfCustomMetadataFieldNewsSchemaDefaultValueJsonScalarArray) {
		return &u.OfCustomMetadataFieldNewsSchemaDefaultValueJsonScalarArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CustomMetadataFieldNewParamsSchemaDefaultValueUnionUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u CustomMetadataFieldNewParamsSchemaDefaultValueUnionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *CustomMetadataFieldNewParamsSchemaDefaultValueUnionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CustomMetadataFieldNewParamsSchemaDefaultValueUnionUnion) asAny() any {
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
type CustomMetadataFieldNewParamsSchemaDefaultValueJsonScalarArrayItemUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u CustomMetadataFieldNewParamsSchemaDefaultValueJsonScalarArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *CustomMetadataFieldNewParamsSchemaDefaultValueJsonScalarArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CustomMetadataFieldNewParamsSchemaDefaultValueJsonScalarArrayItemUnion) asAny() any {
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
	OfCustomMetadataFieldUpdatesSchemaDefaultValueUnion           *CustomMetadataFieldUpdateParamsSchemaDefaultValueUnionUnion                `json:",omitzero,inline"`
	OfCustomMetadataFieldUpdatesSchemaDefaultValueJsonScalarArray []CustomMetadataFieldUpdateParamsSchemaDefaultValueJsonScalarArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u CustomMetadataFieldUpdateParamsSchemaDefaultValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCustomMetadataFieldUpdatesSchemaDefaultValueUnion, u.OfCustomMetadataFieldUpdatesSchemaDefaultValueJsonScalarArray)
}
func (u *CustomMetadataFieldUpdateParamsSchemaDefaultValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CustomMetadataFieldUpdateParamsSchemaDefaultValueUnion) asAny() any {
	if !param.IsOmitted(u.OfCustomMetadataFieldUpdatesSchemaDefaultValueUnion) {
		return u.OfCustomMetadataFieldUpdatesSchemaDefaultValueUnion.asAny()
	} else if !param.IsOmitted(u.OfCustomMetadataFieldUpdatesSchemaDefaultValueJsonScalarArray) {
		return &u.OfCustomMetadataFieldUpdatesSchemaDefaultValueJsonScalarArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CustomMetadataFieldUpdateParamsSchemaDefaultValueUnionUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u CustomMetadataFieldUpdateParamsSchemaDefaultValueUnionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *CustomMetadataFieldUpdateParamsSchemaDefaultValueUnionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CustomMetadataFieldUpdateParamsSchemaDefaultValueUnionUnion) asAny() any {
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
type CustomMetadataFieldUpdateParamsSchemaDefaultValueJsonScalarArrayItemUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u CustomMetadataFieldUpdateParamsSchemaDefaultValueJsonScalarArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *CustomMetadataFieldUpdateParamsSchemaDefaultValueJsonScalarArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CustomMetadataFieldUpdateParamsSchemaDefaultValueJsonScalarArrayItemUnion) asAny() any {
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
