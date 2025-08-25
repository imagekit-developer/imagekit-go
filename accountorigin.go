// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/imagekit-go/internal/apijson"
	"github.com/stainless-sdks/imagekit-go/internal/requestconfig"
	"github.com/stainless-sdks/imagekit-go/option"
	"github.com/stainless-sdks/imagekit-go/packages/param"
	"github.com/stainless-sdks/imagekit-go/packages/respjson"
	"github.com/stainless-sdks/imagekit-go/shared/constant"
)

// AccountOriginService contains methods and other services that help with
// interacting with the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountOriginService] method instead.
type AccountOriginService struct {
	Options []option.RequestOption
}

// NewAccountOriginService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountOriginService(opts ...option.RequestOption) (r AccountOriginService) {
	r = AccountOriginService{}
	r.Options = opts
	return
}

// **Note:** This API is currently in beta.
// Creates a new origin and returns the origin object.
func (r *AccountOriginService) New(ctx context.Context, body AccountOriginNewParams, opts ...option.RequestOption) (res *OriginResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/accounts/origins"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// **Note:** This API is currently in beta.
// Updates the origin identified by `id` and returns the updated origin object.
func (r *AccountOriginService) Update(ctx context.Context, id string, body AccountOriginUpdateParams, opts ...option.RequestOption) (res *OriginResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/accounts/origins/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// **Note:** This API is currently in beta.
// Returns an array of all configured origins for the current account.
func (r *AccountOriginService) List(ctx context.Context, opts ...option.RequestOption) (res *[]OriginResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/accounts/origins"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// **Note:** This API is currently in beta.
// Permanently removes the origin identified by `id`. If the origin is in use by
// any URL‑endpoints, the API will return an error.
func (r *AccountOriginService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/accounts/origins/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// **Note:** This API is currently in beta.
// Retrieves the origin identified by `id`.
func (r *AccountOriginService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *OriginResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/accounts/origins/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// OriginResponseUnion contains all possible properties and values from
// [OriginResponseS3], [OriginResponseS3Compatible],
// [OriginResponseCloudinaryBackup], [OriginResponseWebFolder],
// [OriginResponseWebProxy], [OriginResponseGcs], [OriginResponseAzureBlob],
// [OriginResponseAkeneoPim].
//
// Use the [OriginResponseUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type OriginResponseUnion struct {
	ID                     string `json:"id"`
	Bucket                 string `json:"bucket"`
	IncludeCanonicalHeader bool   `json:"includeCanonicalHeader"`
	Name                   string `json:"name"`
	Prefix                 string `json:"prefix"`
	// Any of "S3", "S3_COMPATIBLE", "CLOUDINARY_BACKUP", "WEB_FOLDER", "WEB_PROXY",
	// "GCS", "AZURE_BLOB", "AKENEO_PIM".
	Type                      string `json:"type"`
	BaseURLForCanonicalHeader string `json:"baseUrlForCanonicalHeader"`
	// This field is from variant [OriginResponseS3Compatible].
	Endpoint string `json:"endpoint"`
	// This field is from variant [OriginResponseS3Compatible].
	S3ForcePathStyle bool   `json:"s3ForcePathStyle"`
	BaseURL          string `json:"baseUrl"`
	// This field is from variant [OriginResponseWebFolder].
	ForwardHostHeaderToOrigin bool `json:"forwardHostHeaderToOrigin"`
	// This field is from variant [OriginResponseGcs].
	ClientEmail string `json:"clientEmail"`
	// This field is from variant [OriginResponseAzureBlob].
	AccountName string `json:"accountName"`
	// This field is from variant [OriginResponseAzureBlob].
	Container string `json:"container"`
	JSON      struct {
		ID                        respjson.Field
		Bucket                    respjson.Field
		IncludeCanonicalHeader    respjson.Field
		Name                      respjson.Field
		Prefix                    respjson.Field
		Type                      respjson.Field
		BaseURLForCanonicalHeader respjson.Field
		Endpoint                  respjson.Field
		S3ForcePathStyle          respjson.Field
		BaseURL                   respjson.Field
		ForwardHostHeaderToOrigin respjson.Field
		ClientEmail               respjson.Field
		AccountName               respjson.Field
		Container                 respjson.Field
		raw                       string
	} `json:"-"`
}

// anyOriginResponse is implemented by each variant of [OriginResponseUnion] to add
// type safety for the return type of [OriginResponseUnion.AsAny]
type anyOriginResponse interface {
	implOriginResponseUnion()
}

func (OriginResponseS3) implOriginResponseUnion()               {}
func (OriginResponseS3Compatible) implOriginResponseUnion()     {}
func (OriginResponseCloudinaryBackup) implOriginResponseUnion() {}
func (OriginResponseWebFolder) implOriginResponseUnion()        {}
func (OriginResponseWebProxy) implOriginResponseUnion()         {}
func (OriginResponseGcs) implOriginResponseUnion()              {}
func (OriginResponseAzureBlob) implOriginResponseUnion()        {}
func (OriginResponseAkeneoPim) implOriginResponseUnion()        {}

// Use the following switch statement to find the correct variant
//
//	switch variant := OriginResponseUnion.AsAny().(type) {
//	case imagekit.OriginResponseS3:
//	case imagekit.OriginResponseS3Compatible:
//	case imagekit.OriginResponseCloudinaryBackup:
//	case imagekit.OriginResponseWebFolder:
//	case imagekit.OriginResponseWebProxy:
//	case imagekit.OriginResponseGcs:
//	case imagekit.OriginResponseAzureBlob:
//	case imagekit.OriginResponseAkeneoPim:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u OriginResponseUnion) AsAny() anyOriginResponse {
	switch u.Type {
	case "S3":
		return u.AsS3()
	case "S3_COMPATIBLE":
		return u.AsS3Compatible()
	case "CLOUDINARY_BACKUP":
		return u.AsCloudinaryBackup()
	case "WEB_FOLDER":
		return u.AsWebFolder()
	case "WEB_PROXY":
		return u.AsWebProxy()
	case "GCS":
		return u.AsGcs()
	case "AZURE_BLOB":
		return u.AsAzureBlob()
	case "AKENEO_PIM":
		return u.AsAkeneoPim()
	}
	return nil
}

func (u OriginResponseUnion) AsS3() (v OriginResponseS3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OriginResponseUnion) AsS3Compatible() (v OriginResponseS3Compatible) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OriginResponseUnion) AsCloudinaryBackup() (v OriginResponseCloudinaryBackup) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OriginResponseUnion) AsWebFolder() (v OriginResponseWebFolder) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OriginResponseUnion) AsWebProxy() (v OriginResponseWebProxy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OriginResponseUnion) AsGcs() (v OriginResponseGcs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OriginResponseUnion) AsAzureBlob() (v OriginResponseAzureBlob) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OriginResponseUnion) AsAkeneoPim() (v OriginResponseAkeneoPim) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OriginResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *OriginResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OriginResponseS3 struct {
	// Unique identifier for the origin. This is generated by ImageKit when you create
	// a new origin.
	ID string `json:"id,required"`
	// S3 bucket name.
	Bucket string `json:"bucket,required"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader bool `json:"includeCanonicalHeader,required"`
	// Display name of the origin.
	Name string `json:"name,required"`
	// Path prefix inside the bucket.
	Prefix string      `json:"prefix,required"`
	Type   constant.S3 `json:"type,required"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader string `json:"baseUrlForCanonicalHeader" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		Bucket                    respjson.Field
		IncludeCanonicalHeader    respjson.Field
		Name                      respjson.Field
		Prefix                    respjson.Field
		Type                      respjson.Field
		BaseURLForCanonicalHeader respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OriginResponseS3) RawJSON() string { return r.JSON.raw }
func (r *OriginResponseS3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OriginResponseS3Compatible struct {
	// Unique identifier for the origin. This is generated by ImageKit when you create
	// a new origin.
	ID string `json:"id,required"`
	// S3 bucket name.
	Bucket string `json:"bucket,required"`
	// Custom S3-compatible endpoint.
	Endpoint string `json:"endpoint,required" format:"uri"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader bool `json:"includeCanonicalHeader,required"`
	// Display name of the origin.
	Name string `json:"name,required"`
	// Path prefix inside the bucket.
	Prefix string `json:"prefix,required"`
	// Use path-style S3 URLs?
	S3ForcePathStyle bool                  `json:"s3ForcePathStyle,required"`
	Type             constant.S3Compatible `json:"type,required"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader string `json:"baseUrlForCanonicalHeader" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		Bucket                    respjson.Field
		Endpoint                  respjson.Field
		IncludeCanonicalHeader    respjson.Field
		Name                      respjson.Field
		Prefix                    respjson.Field
		S3ForcePathStyle          respjson.Field
		Type                      respjson.Field
		BaseURLForCanonicalHeader respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OriginResponseS3Compatible) RawJSON() string { return r.JSON.raw }
func (r *OriginResponseS3Compatible) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OriginResponseCloudinaryBackup struct {
	// Unique identifier for the origin. This is generated by ImageKit when you create
	// a new origin.
	ID string `json:"id,required"`
	// S3 bucket name.
	Bucket string `json:"bucket,required"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader bool `json:"includeCanonicalHeader,required"`
	// Display name of the origin.
	Name string `json:"name,required"`
	// Path prefix inside the bucket.
	Prefix string                    `json:"prefix,required"`
	Type   constant.CloudinaryBackup `json:"type,required"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader string `json:"baseUrlForCanonicalHeader" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		Bucket                    respjson.Field
		IncludeCanonicalHeader    respjson.Field
		Name                      respjson.Field
		Prefix                    respjson.Field
		Type                      respjson.Field
		BaseURLForCanonicalHeader respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OriginResponseCloudinaryBackup) RawJSON() string { return r.JSON.raw }
func (r *OriginResponseCloudinaryBackup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OriginResponseWebFolder struct {
	// Unique identifier for the origin. This is generated by ImageKit when you create
	// a new origin.
	ID string `json:"id,required"`
	// Root URL for the web folder origin.
	BaseURL string `json:"baseUrl,required" format:"uri"`
	// Forward the Host header to origin?
	ForwardHostHeaderToOrigin bool `json:"forwardHostHeaderToOrigin,required"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader bool `json:"includeCanonicalHeader,required"`
	// Display name of the origin.
	Name string             `json:"name,required"`
	Type constant.WebFolder `json:"type,required"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader string `json:"baseUrlForCanonicalHeader" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		BaseURL                   respjson.Field
		ForwardHostHeaderToOrigin respjson.Field
		IncludeCanonicalHeader    respjson.Field
		Name                      respjson.Field
		Type                      respjson.Field
		BaseURLForCanonicalHeader respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OriginResponseWebFolder) RawJSON() string { return r.JSON.raw }
func (r *OriginResponseWebFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OriginResponseWebProxy struct {
	// Unique identifier for the origin. This is generated by ImageKit when you create
	// a new origin.
	ID string `json:"id,required"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader bool `json:"includeCanonicalHeader,required"`
	// Display name of the origin.
	Name string            `json:"name,required"`
	Type constant.WebProxy `json:"type,required"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader string `json:"baseUrlForCanonicalHeader" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		IncludeCanonicalHeader    respjson.Field
		Name                      respjson.Field
		Type                      respjson.Field
		BaseURLForCanonicalHeader respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OriginResponseWebProxy) RawJSON() string { return r.JSON.raw }
func (r *OriginResponseWebProxy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OriginResponseGcs struct {
	// Unique identifier for the origin. This is generated by ImageKit when you create
	// a new origin.
	ID          string `json:"id,required"`
	Bucket      string `json:"bucket,required"`
	ClientEmail string `json:"clientEmail,required" format:"email"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader bool `json:"includeCanonicalHeader,required"`
	// Display name of the origin.
	Name   string       `json:"name,required"`
	Prefix string       `json:"prefix,required"`
	Type   constant.Gcs `json:"type,required"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader string `json:"baseUrlForCanonicalHeader" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		Bucket                    respjson.Field
		ClientEmail               respjson.Field
		IncludeCanonicalHeader    respjson.Field
		Name                      respjson.Field
		Prefix                    respjson.Field
		Type                      respjson.Field
		BaseURLForCanonicalHeader respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OriginResponseGcs) RawJSON() string { return r.JSON.raw }
func (r *OriginResponseGcs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OriginResponseAzureBlob struct {
	// Unique identifier for the origin. This is generated by ImageKit when you create
	// a new origin.
	ID          string `json:"id,required"`
	AccountName string `json:"accountName,required"`
	Container   string `json:"container,required"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader bool `json:"includeCanonicalHeader,required"`
	// Display name of the origin.
	Name   string             `json:"name,required"`
	Prefix string             `json:"prefix,required"`
	Type   constant.AzureBlob `json:"type,required"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader string `json:"baseUrlForCanonicalHeader" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		AccountName               respjson.Field
		Container                 respjson.Field
		IncludeCanonicalHeader    respjson.Field
		Name                      respjson.Field
		Prefix                    respjson.Field
		Type                      respjson.Field
		BaseURLForCanonicalHeader respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OriginResponseAzureBlob) RawJSON() string { return r.JSON.raw }
func (r *OriginResponseAzureBlob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OriginResponseAkeneoPim struct {
	// Unique identifier for the origin. This is generated by ImageKit when you create
	// a new origin.
	ID string `json:"id,required"`
	// Akeneo instance base URL.
	BaseURL string `json:"baseUrl,required" format:"uri"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader bool `json:"includeCanonicalHeader,required"`
	// Display name of the origin.
	Name string             `json:"name,required"`
	Type constant.AkeneoPim `json:"type,required"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader string `json:"baseUrlForCanonicalHeader" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		BaseURL                   respjson.Field
		IncludeCanonicalHeader    respjson.Field
		Name                      respjson.Field
		Type                      respjson.Field
		BaseURLForCanonicalHeader respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OriginResponseAkeneoPim) RawJSON() string { return r.JSON.raw }
func (r *OriginResponseAkeneoPim) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfS3 *AccountOriginNewParamsOriginS3 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfS3Compatible *AccountOriginNewParamsOriginS3Compatible `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfCloudinaryBackup *AccountOriginNewParamsOriginCloudinaryBackup `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfWebFolder *AccountOriginNewParamsOriginWebFolder `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfWebProxy *AccountOriginNewParamsOriginWebProxy `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfGcs *AccountOriginNewParamsOriginGcs `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfAzureBlob *AccountOriginNewParamsOriginAzureBlob `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfAkeneoPim *AccountOriginNewParamsOriginAkeneoPim `json:",inline"`

	paramObj
}

func (u AccountOriginNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfS3,
		u.OfS3Compatible,
		u.OfCloudinaryBackup,
		u.OfWebFolder,
		u.OfWebProxy,
		u.OfGcs,
		u.OfAzureBlob,
		u.OfAkeneoPim)
}
func (r *AccountOriginNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessKey, Bucket, Name, SecretKey, Type are required.
type AccountOriginNewParamsOriginS3 struct {
	// Access key for the bucket.
	AccessKey string `json:"accessKey,required"`
	// S3 bucket name.
	Bucket string `json:"bucket,required"`
	// Display name of the origin.
	Name string `json:"name,required"`
	// Secret key for the bucket.
	SecretKey string `json:"secretKey,required"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader param.Opt[string] `json:"baseUrlForCanonicalHeader,omitzero" format:"uri"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader param.Opt[bool] `json:"includeCanonicalHeader,omitzero"`
	// Path prefix inside the bucket.
	Prefix param.Opt[string] `json:"prefix,omitzero"`
	// This field can be elided, and will marshal its zero value as "S3".
	Type constant.S3 `json:"type,required"`
	paramObj
}

func (r AccountOriginNewParamsOriginS3) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginNewParamsOriginS3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginNewParamsOriginS3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessKey, Bucket, Endpoint, Name, SecretKey, Type are required.
type AccountOriginNewParamsOriginS3Compatible struct {
	// Access key for the bucket.
	AccessKey string `json:"accessKey,required"`
	// S3 bucket name.
	Bucket string `json:"bucket,required"`
	// Custom S3-compatible endpoint.
	Endpoint string `json:"endpoint,required" format:"uri"`
	// Display name of the origin.
	Name string `json:"name,required"`
	// Secret key for the bucket.
	SecretKey string `json:"secretKey,required"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader param.Opt[string] `json:"baseUrlForCanonicalHeader,omitzero" format:"uri"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader param.Opt[bool] `json:"includeCanonicalHeader,omitzero"`
	// Path prefix inside the bucket.
	Prefix param.Opt[string] `json:"prefix,omitzero"`
	// Use path-style S3 URLs?
	S3ForcePathStyle param.Opt[bool] `json:"s3ForcePathStyle,omitzero"`
	// This field can be elided, and will marshal its zero value as "S3_COMPATIBLE".
	Type constant.S3Compatible `json:"type,required"`
	paramObj
}

func (r AccountOriginNewParamsOriginS3Compatible) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginNewParamsOriginS3Compatible
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginNewParamsOriginS3Compatible) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessKey, Bucket, Name, SecretKey, Type are required.
type AccountOriginNewParamsOriginCloudinaryBackup struct {
	// Access key for the bucket.
	AccessKey string `json:"accessKey,required"`
	// S3 bucket name.
	Bucket string `json:"bucket,required"`
	// Display name of the origin.
	Name string `json:"name,required"`
	// Secret key for the bucket.
	SecretKey string `json:"secretKey,required"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader param.Opt[string] `json:"baseUrlForCanonicalHeader,omitzero" format:"uri"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader param.Opt[bool] `json:"includeCanonicalHeader,omitzero"`
	// Path prefix inside the bucket.
	Prefix param.Opt[string] `json:"prefix,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "CLOUDINARY_BACKUP".
	Type constant.CloudinaryBackup `json:"type,required"`
	paramObj
}

func (r AccountOriginNewParamsOriginCloudinaryBackup) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginNewParamsOriginCloudinaryBackup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginNewParamsOriginCloudinaryBackup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties BaseURL, Name, Type are required.
type AccountOriginNewParamsOriginWebFolder struct {
	// Root URL for the web folder origin.
	BaseURL string `json:"baseUrl,required" format:"uri"`
	// Display name of the origin.
	Name string `json:"name,required"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader param.Opt[string] `json:"baseUrlForCanonicalHeader,omitzero" format:"uri"`
	// Forward the Host header to origin?
	ForwardHostHeaderToOrigin param.Opt[bool] `json:"forwardHostHeaderToOrigin,omitzero"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader param.Opt[bool] `json:"includeCanonicalHeader,omitzero"`
	// This field can be elided, and will marshal its zero value as "WEB_FOLDER".
	Type constant.WebFolder `json:"type,required"`
	paramObj
}

func (r AccountOriginNewParamsOriginWebFolder) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginNewParamsOriginWebFolder
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginNewParamsOriginWebFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Type are required.
type AccountOriginNewParamsOriginWebProxy struct {
	// Display name of the origin.
	Name string `json:"name,required"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader param.Opt[string] `json:"baseUrlForCanonicalHeader,omitzero" format:"uri"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader param.Opt[bool] `json:"includeCanonicalHeader,omitzero"`
	// This field can be elided, and will marshal its zero value as "WEB_PROXY".
	Type constant.WebProxy `json:"type,required"`
	paramObj
}

func (r AccountOriginNewParamsOriginWebProxy) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginNewParamsOriginWebProxy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginNewParamsOriginWebProxy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Bucket, ClientEmail, Name, PrivateKey, Type are required.
type AccountOriginNewParamsOriginGcs struct {
	Bucket      string `json:"bucket,required"`
	ClientEmail string `json:"clientEmail,required" format:"email"`
	// Display name of the origin.
	Name       string `json:"name,required"`
	PrivateKey string `json:"privateKey,required"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader param.Opt[string] `json:"baseUrlForCanonicalHeader,omitzero" format:"uri"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader param.Opt[bool]   `json:"includeCanonicalHeader,omitzero"`
	Prefix                 param.Opt[string] `json:"prefix,omitzero"`
	// This field can be elided, and will marshal its zero value as "GCS".
	Type constant.Gcs `json:"type,required"`
	paramObj
}

func (r AccountOriginNewParamsOriginGcs) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginNewParamsOriginGcs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginNewParamsOriginGcs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccountName, Container, Name, SasToken, Type are required.
type AccountOriginNewParamsOriginAzureBlob struct {
	AccountName string `json:"accountName,required"`
	Container   string `json:"container,required"`
	// Display name of the origin.
	Name     string `json:"name,required"`
	SasToken string `json:"sasToken,required"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader param.Opt[string] `json:"baseUrlForCanonicalHeader,omitzero" format:"uri"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader param.Opt[bool]   `json:"includeCanonicalHeader,omitzero"`
	Prefix                 param.Opt[string] `json:"prefix,omitzero"`
	// This field can be elided, and will marshal its zero value as "AZURE_BLOB".
	Type constant.AzureBlob `json:"type,required"`
	paramObj
}

func (r AccountOriginNewParamsOriginAzureBlob) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginNewParamsOriginAzureBlob
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginNewParamsOriginAzureBlob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties BaseURL, ClientID, ClientSecret, Name, Password, Type, Username
// are required.
type AccountOriginNewParamsOriginAkeneoPim struct {
	// Akeneo instance base URL.
	BaseURL string `json:"baseUrl,required" format:"uri"`
	// Akeneo API client ID.
	ClientID string `json:"clientId,required"`
	// Akeneo API client secret.
	ClientSecret string `json:"clientSecret,required"`
	// Display name of the origin.
	Name string `json:"name,required"`
	// Akeneo API password.
	Password string `json:"password,required"`
	// Akeneo API username.
	Username string `json:"username,required"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader param.Opt[string] `json:"baseUrlForCanonicalHeader,omitzero" format:"uri"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader param.Opt[bool] `json:"includeCanonicalHeader,omitzero"`
	// This field can be elided, and will marshal its zero value as "AKENEO_PIM".
	Type constant.AkeneoPim `json:"type,required"`
	paramObj
}

func (r AccountOriginNewParamsOriginAkeneoPim) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginNewParamsOriginAkeneoPim
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginNewParamsOriginAkeneoPim) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginUpdateParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfS3 *AccountOriginUpdateParamsOriginS3 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfS3Compatible *AccountOriginUpdateParamsOriginS3Compatible `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfCloudinaryBackup *AccountOriginUpdateParamsOriginCloudinaryBackup `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfWebFolder *AccountOriginUpdateParamsOriginWebFolder `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfWebProxy *AccountOriginUpdateParamsOriginWebProxy `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfGcs *AccountOriginUpdateParamsOriginGcs `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfAzureBlob *AccountOriginUpdateParamsOriginAzureBlob `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfAkeneoPim *AccountOriginUpdateParamsOriginAkeneoPim `json:",inline"`

	paramObj
}

func (u AccountOriginUpdateParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfS3,
		u.OfS3Compatible,
		u.OfCloudinaryBackup,
		u.OfWebFolder,
		u.OfWebProxy,
		u.OfGcs,
		u.OfAzureBlob,
		u.OfAkeneoPim)
}
func (r *AccountOriginUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessKey, Bucket, Name, SecretKey, Type are required.
type AccountOriginUpdateParamsOriginS3 struct {
	// Access key for the bucket.
	AccessKey string `json:"accessKey,required"`
	// S3 bucket name.
	Bucket string `json:"bucket,required"`
	// Display name of the origin.
	Name string `json:"name,required"`
	// Secret key for the bucket.
	SecretKey string `json:"secretKey,required"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader param.Opt[string] `json:"baseUrlForCanonicalHeader,omitzero" format:"uri"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader param.Opt[bool] `json:"includeCanonicalHeader,omitzero"`
	// Path prefix inside the bucket.
	Prefix param.Opt[string] `json:"prefix,omitzero"`
	// This field can be elided, and will marshal its zero value as "S3".
	Type constant.S3 `json:"type,required"`
	paramObj
}

func (r AccountOriginUpdateParamsOriginS3) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginUpdateParamsOriginS3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginUpdateParamsOriginS3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessKey, Bucket, Endpoint, Name, SecretKey, Type are required.
type AccountOriginUpdateParamsOriginS3Compatible struct {
	// Access key for the bucket.
	AccessKey string `json:"accessKey,required"`
	// S3 bucket name.
	Bucket string `json:"bucket,required"`
	// Custom S3-compatible endpoint.
	Endpoint string `json:"endpoint,required" format:"uri"`
	// Display name of the origin.
	Name string `json:"name,required"`
	// Secret key for the bucket.
	SecretKey string `json:"secretKey,required"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader param.Opt[string] `json:"baseUrlForCanonicalHeader,omitzero" format:"uri"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader param.Opt[bool] `json:"includeCanonicalHeader,omitzero"`
	// Path prefix inside the bucket.
	Prefix param.Opt[string] `json:"prefix,omitzero"`
	// Use path-style S3 URLs?
	S3ForcePathStyle param.Opt[bool] `json:"s3ForcePathStyle,omitzero"`
	// This field can be elided, and will marshal its zero value as "S3_COMPATIBLE".
	Type constant.S3Compatible `json:"type,required"`
	paramObj
}

func (r AccountOriginUpdateParamsOriginS3Compatible) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginUpdateParamsOriginS3Compatible
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginUpdateParamsOriginS3Compatible) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessKey, Bucket, Name, SecretKey, Type are required.
type AccountOriginUpdateParamsOriginCloudinaryBackup struct {
	// Access key for the bucket.
	AccessKey string `json:"accessKey,required"`
	// S3 bucket name.
	Bucket string `json:"bucket,required"`
	// Display name of the origin.
	Name string `json:"name,required"`
	// Secret key for the bucket.
	SecretKey string `json:"secretKey,required"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader param.Opt[string] `json:"baseUrlForCanonicalHeader,omitzero" format:"uri"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader param.Opt[bool] `json:"includeCanonicalHeader,omitzero"`
	// Path prefix inside the bucket.
	Prefix param.Opt[string] `json:"prefix,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "CLOUDINARY_BACKUP".
	Type constant.CloudinaryBackup `json:"type,required"`
	paramObj
}

func (r AccountOriginUpdateParamsOriginCloudinaryBackup) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginUpdateParamsOriginCloudinaryBackup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginUpdateParamsOriginCloudinaryBackup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties BaseURL, Name, Type are required.
type AccountOriginUpdateParamsOriginWebFolder struct {
	// Root URL for the web folder origin.
	BaseURL string `json:"baseUrl,required" format:"uri"`
	// Display name of the origin.
	Name string `json:"name,required"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader param.Opt[string] `json:"baseUrlForCanonicalHeader,omitzero" format:"uri"`
	// Forward the Host header to origin?
	ForwardHostHeaderToOrigin param.Opt[bool] `json:"forwardHostHeaderToOrigin,omitzero"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader param.Opt[bool] `json:"includeCanonicalHeader,omitzero"`
	// This field can be elided, and will marshal its zero value as "WEB_FOLDER".
	Type constant.WebFolder `json:"type,required"`
	paramObj
}

func (r AccountOriginUpdateParamsOriginWebFolder) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginUpdateParamsOriginWebFolder
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginUpdateParamsOriginWebFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Type are required.
type AccountOriginUpdateParamsOriginWebProxy struct {
	// Display name of the origin.
	Name string `json:"name,required"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader param.Opt[string] `json:"baseUrlForCanonicalHeader,omitzero" format:"uri"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader param.Opt[bool] `json:"includeCanonicalHeader,omitzero"`
	// This field can be elided, and will marshal its zero value as "WEB_PROXY".
	Type constant.WebProxy `json:"type,required"`
	paramObj
}

func (r AccountOriginUpdateParamsOriginWebProxy) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginUpdateParamsOriginWebProxy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginUpdateParamsOriginWebProxy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Bucket, ClientEmail, Name, PrivateKey, Type are required.
type AccountOriginUpdateParamsOriginGcs struct {
	Bucket      string `json:"bucket,required"`
	ClientEmail string `json:"clientEmail,required" format:"email"`
	// Display name of the origin.
	Name       string `json:"name,required"`
	PrivateKey string `json:"privateKey,required"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader param.Opt[string] `json:"baseUrlForCanonicalHeader,omitzero" format:"uri"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader param.Opt[bool]   `json:"includeCanonicalHeader,omitzero"`
	Prefix                 param.Opt[string] `json:"prefix,omitzero"`
	// This field can be elided, and will marshal its zero value as "GCS".
	Type constant.Gcs `json:"type,required"`
	paramObj
}

func (r AccountOriginUpdateParamsOriginGcs) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginUpdateParamsOriginGcs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginUpdateParamsOriginGcs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccountName, Container, Name, SasToken, Type are required.
type AccountOriginUpdateParamsOriginAzureBlob struct {
	AccountName string `json:"accountName,required"`
	Container   string `json:"container,required"`
	// Display name of the origin.
	Name     string `json:"name,required"`
	SasToken string `json:"sasToken,required"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader param.Opt[string] `json:"baseUrlForCanonicalHeader,omitzero" format:"uri"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader param.Opt[bool]   `json:"includeCanonicalHeader,omitzero"`
	Prefix                 param.Opt[string] `json:"prefix,omitzero"`
	// This field can be elided, and will marshal its zero value as "AZURE_BLOB".
	Type constant.AzureBlob `json:"type,required"`
	paramObj
}

func (r AccountOriginUpdateParamsOriginAzureBlob) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginUpdateParamsOriginAzureBlob
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginUpdateParamsOriginAzureBlob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties BaseURL, ClientID, ClientSecret, Name, Password, Type, Username
// are required.
type AccountOriginUpdateParamsOriginAkeneoPim struct {
	// Akeneo instance base URL.
	BaseURL string `json:"baseUrl,required" format:"uri"`
	// Akeneo API client ID.
	ClientID string `json:"clientId,required"`
	// Akeneo API client secret.
	ClientSecret string `json:"clientSecret,required"`
	// Display name of the origin.
	Name string `json:"name,required"`
	// Akeneo API password.
	Password string `json:"password,required"`
	// Akeneo API username.
	Username string `json:"username,required"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader param.Opt[string] `json:"baseUrlForCanonicalHeader,omitzero" format:"uri"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader param.Opt[bool] `json:"includeCanonicalHeader,omitzero"`
	// This field can be elided, and will marshal its zero value as "AKENEO_PIM".
	Type constant.AkeneoPim `json:"type,required"`
	paramObj
}

func (r AccountOriginUpdateParamsOriginAkeneoPim) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginUpdateParamsOriginAkeneoPim
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginUpdateParamsOriginAkeneoPim) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
