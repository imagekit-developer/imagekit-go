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
func (r *AccountOriginService) New(ctx context.Context, body AccountOriginNewParams, opts ...option.RequestOption) (res *AccountOriginNewResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/accounts/origins"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// **Note:** This API is currently in beta.
// Updates the origin identified by `id` and returns the updated origin object.
func (r *AccountOriginService) Update(ctx context.Context, id string, body AccountOriginUpdateParams, opts ...option.RequestOption) (res *AccountOriginUpdateResponseUnion, err error) {
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
func (r *AccountOriginService) List(ctx context.Context, opts ...option.RequestOption) (res *[]AccountOriginListResponseUnion, err error) {
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
func (r *AccountOriginService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AccountOriginGetResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/accounts/origins/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// AccountOriginNewResponseUnion contains all possible properties and values from
// [AccountOriginNewResponseS3], [AccountOriginNewResponseS3Compatible],
// [AccountOriginNewResponseCloudinaryBackup], [AccountOriginNewResponseWebFolder],
// [AccountOriginNewResponseWebProxy],
// [AccountOriginNewResponseGoogleCloudStorageGcs],
// [AccountOriginNewResponseAzureBlobStorage], [AccountOriginNewResponseAkeneoPim].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AccountOriginNewResponseUnion struct {
	ID                        string `json:"id"`
	Bucket                    string `json:"bucket"`
	IncludeCanonicalHeader    bool   `json:"includeCanonicalHeader"`
	Name                      string `json:"name"`
	Prefix                    string `json:"prefix"`
	Type                      string `json:"type"`
	BaseURLForCanonicalHeader string `json:"baseUrlForCanonicalHeader"`
	// This field is from variant [AccountOriginNewResponseS3Compatible].
	Endpoint string `json:"endpoint"`
	// This field is from variant [AccountOriginNewResponseS3Compatible].
	S3ForcePathStyle bool   `json:"s3ForcePathStyle"`
	BaseURL          string `json:"baseUrl"`
	// This field is from variant [AccountOriginNewResponseWebFolder].
	ForwardHostHeaderToOrigin bool `json:"forwardHostHeaderToOrigin"`
	// This field is from variant [AccountOriginNewResponseGoogleCloudStorageGcs].
	ClientEmail string `json:"clientEmail"`
	// This field is from variant [AccountOriginNewResponseAzureBlobStorage].
	AccountName string `json:"accountName"`
	// This field is from variant [AccountOriginNewResponseAzureBlobStorage].
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

func (u AccountOriginNewResponseUnion) AsS3() (v AccountOriginNewResponseS3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginNewResponseUnion) AsS3Compatible() (v AccountOriginNewResponseS3Compatible) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginNewResponseUnion) AsCloudinaryBackup() (v AccountOriginNewResponseCloudinaryBackup) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginNewResponseUnion) AsWebFolder() (v AccountOriginNewResponseWebFolder) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginNewResponseUnion) AsWebProxy() (v AccountOriginNewResponseWebProxy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginNewResponseUnion) AsGoogleCloudStorageGcs() (v AccountOriginNewResponseGoogleCloudStorageGcs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginNewResponseUnion) AsAzureBlobStorage() (v AccountOriginNewResponseAzureBlobStorage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginNewResponseUnion) AsAkeneoPim() (v AccountOriginNewResponseAkeneoPim) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountOriginNewResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *AccountOriginNewResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginNewResponseS3 struct {
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
func (r AccountOriginNewResponseS3) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginNewResponseS3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginNewResponseS3Compatible struct {
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
func (r AccountOriginNewResponseS3Compatible) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginNewResponseS3Compatible) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginNewResponseCloudinaryBackup struct {
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
func (r AccountOriginNewResponseCloudinaryBackup) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginNewResponseCloudinaryBackup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginNewResponseWebFolder struct {
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
func (r AccountOriginNewResponseWebFolder) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginNewResponseWebFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginNewResponseWebProxy struct {
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
func (r AccountOriginNewResponseWebProxy) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginNewResponseWebProxy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginNewResponseGoogleCloudStorageGcs struct {
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
func (r AccountOriginNewResponseGoogleCloudStorageGcs) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginNewResponseGoogleCloudStorageGcs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginNewResponseAzureBlobStorage struct {
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
func (r AccountOriginNewResponseAzureBlobStorage) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginNewResponseAzureBlobStorage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginNewResponseAkeneoPim struct {
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
func (r AccountOriginNewResponseAkeneoPim) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginNewResponseAkeneoPim) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountOriginUpdateResponseUnion contains all possible properties and values
// from [AccountOriginUpdateResponseS3], [AccountOriginUpdateResponseS3Compatible],
// [AccountOriginUpdateResponseCloudinaryBackup],
// [AccountOriginUpdateResponseWebFolder], [AccountOriginUpdateResponseWebProxy],
// [AccountOriginUpdateResponseGoogleCloudStorageGcs],
// [AccountOriginUpdateResponseAzureBlobStorage],
// [AccountOriginUpdateResponseAkeneoPim].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AccountOriginUpdateResponseUnion struct {
	ID                        string `json:"id"`
	Bucket                    string `json:"bucket"`
	IncludeCanonicalHeader    bool   `json:"includeCanonicalHeader"`
	Name                      string `json:"name"`
	Prefix                    string `json:"prefix"`
	Type                      string `json:"type"`
	BaseURLForCanonicalHeader string `json:"baseUrlForCanonicalHeader"`
	// This field is from variant [AccountOriginUpdateResponseS3Compatible].
	Endpoint string `json:"endpoint"`
	// This field is from variant [AccountOriginUpdateResponseS3Compatible].
	S3ForcePathStyle bool   `json:"s3ForcePathStyle"`
	BaseURL          string `json:"baseUrl"`
	// This field is from variant [AccountOriginUpdateResponseWebFolder].
	ForwardHostHeaderToOrigin bool `json:"forwardHostHeaderToOrigin"`
	// This field is from variant [AccountOriginUpdateResponseGoogleCloudStorageGcs].
	ClientEmail string `json:"clientEmail"`
	// This field is from variant [AccountOriginUpdateResponseAzureBlobStorage].
	AccountName string `json:"accountName"`
	// This field is from variant [AccountOriginUpdateResponseAzureBlobStorage].
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

func (u AccountOriginUpdateResponseUnion) AsS3() (v AccountOriginUpdateResponseS3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginUpdateResponseUnion) AsS3Compatible() (v AccountOriginUpdateResponseS3Compatible) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginUpdateResponseUnion) AsCloudinaryBackup() (v AccountOriginUpdateResponseCloudinaryBackup) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginUpdateResponseUnion) AsWebFolder() (v AccountOriginUpdateResponseWebFolder) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginUpdateResponseUnion) AsWebProxy() (v AccountOriginUpdateResponseWebProxy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginUpdateResponseUnion) AsGoogleCloudStorageGcs() (v AccountOriginUpdateResponseGoogleCloudStorageGcs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginUpdateResponseUnion) AsAzureBlobStorage() (v AccountOriginUpdateResponseAzureBlobStorage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginUpdateResponseUnion) AsAkeneoPim() (v AccountOriginUpdateResponseAkeneoPim) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountOriginUpdateResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *AccountOriginUpdateResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginUpdateResponseS3 struct {
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
func (r AccountOriginUpdateResponseS3) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginUpdateResponseS3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginUpdateResponseS3Compatible struct {
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
func (r AccountOriginUpdateResponseS3Compatible) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginUpdateResponseS3Compatible) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginUpdateResponseCloudinaryBackup struct {
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
func (r AccountOriginUpdateResponseCloudinaryBackup) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginUpdateResponseCloudinaryBackup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginUpdateResponseWebFolder struct {
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
func (r AccountOriginUpdateResponseWebFolder) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginUpdateResponseWebFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginUpdateResponseWebProxy struct {
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
func (r AccountOriginUpdateResponseWebProxy) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginUpdateResponseWebProxy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginUpdateResponseGoogleCloudStorageGcs struct {
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
func (r AccountOriginUpdateResponseGoogleCloudStorageGcs) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginUpdateResponseGoogleCloudStorageGcs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginUpdateResponseAzureBlobStorage struct {
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
func (r AccountOriginUpdateResponseAzureBlobStorage) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginUpdateResponseAzureBlobStorage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginUpdateResponseAkeneoPim struct {
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
func (r AccountOriginUpdateResponseAkeneoPim) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginUpdateResponseAkeneoPim) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountOriginListResponseUnion contains all possible properties and values from
// [AccountOriginListResponseS3], [AccountOriginListResponseS3Compatible],
// [AccountOriginListResponseCloudinaryBackup],
// [AccountOriginListResponseWebFolder], [AccountOriginListResponseWebProxy],
// [AccountOriginListResponseGoogleCloudStorageGcs],
// [AccountOriginListResponseAzureBlobStorage],
// [AccountOriginListResponseAkeneoPim].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AccountOriginListResponseUnion struct {
	ID                        string `json:"id"`
	Bucket                    string `json:"bucket"`
	IncludeCanonicalHeader    bool   `json:"includeCanonicalHeader"`
	Name                      string `json:"name"`
	Prefix                    string `json:"prefix"`
	Type                      string `json:"type"`
	BaseURLForCanonicalHeader string `json:"baseUrlForCanonicalHeader"`
	// This field is from variant [AccountOriginListResponseS3Compatible].
	Endpoint string `json:"endpoint"`
	// This field is from variant [AccountOriginListResponseS3Compatible].
	S3ForcePathStyle bool   `json:"s3ForcePathStyle"`
	BaseURL          string `json:"baseUrl"`
	// This field is from variant [AccountOriginListResponseWebFolder].
	ForwardHostHeaderToOrigin bool `json:"forwardHostHeaderToOrigin"`
	// This field is from variant [AccountOriginListResponseGoogleCloudStorageGcs].
	ClientEmail string `json:"clientEmail"`
	// This field is from variant [AccountOriginListResponseAzureBlobStorage].
	AccountName string `json:"accountName"`
	// This field is from variant [AccountOriginListResponseAzureBlobStorage].
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

func (u AccountOriginListResponseUnion) AsS3() (v AccountOriginListResponseS3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginListResponseUnion) AsS3Compatible() (v AccountOriginListResponseS3Compatible) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginListResponseUnion) AsCloudinaryBackup() (v AccountOriginListResponseCloudinaryBackup) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginListResponseUnion) AsWebFolder() (v AccountOriginListResponseWebFolder) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginListResponseUnion) AsWebProxy() (v AccountOriginListResponseWebProxy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginListResponseUnion) AsGoogleCloudStorageGcs() (v AccountOriginListResponseGoogleCloudStorageGcs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginListResponseUnion) AsAzureBlobStorage() (v AccountOriginListResponseAzureBlobStorage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginListResponseUnion) AsAkeneoPim() (v AccountOriginListResponseAkeneoPim) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountOriginListResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *AccountOriginListResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginListResponseS3 struct {
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
func (r AccountOriginListResponseS3) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginListResponseS3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginListResponseS3Compatible struct {
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
func (r AccountOriginListResponseS3Compatible) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginListResponseS3Compatible) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginListResponseCloudinaryBackup struct {
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
func (r AccountOriginListResponseCloudinaryBackup) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginListResponseCloudinaryBackup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginListResponseWebFolder struct {
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
func (r AccountOriginListResponseWebFolder) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginListResponseWebFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginListResponseWebProxy struct {
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
func (r AccountOriginListResponseWebProxy) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginListResponseWebProxy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginListResponseGoogleCloudStorageGcs struct {
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
func (r AccountOriginListResponseGoogleCloudStorageGcs) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginListResponseGoogleCloudStorageGcs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginListResponseAzureBlobStorage struct {
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
func (r AccountOriginListResponseAzureBlobStorage) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginListResponseAzureBlobStorage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginListResponseAkeneoPim struct {
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
func (r AccountOriginListResponseAkeneoPim) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginListResponseAkeneoPim) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccountOriginGetResponseUnion contains all possible properties and values from
// [AccountOriginGetResponseS3], [AccountOriginGetResponseS3Compatible],
// [AccountOriginGetResponseCloudinaryBackup], [AccountOriginGetResponseWebFolder],
// [AccountOriginGetResponseWebProxy],
// [AccountOriginGetResponseGoogleCloudStorageGcs],
// [AccountOriginGetResponseAzureBlobStorage], [AccountOriginGetResponseAkeneoPim].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AccountOriginGetResponseUnion struct {
	ID                        string `json:"id"`
	Bucket                    string `json:"bucket"`
	IncludeCanonicalHeader    bool   `json:"includeCanonicalHeader"`
	Name                      string `json:"name"`
	Prefix                    string `json:"prefix"`
	Type                      string `json:"type"`
	BaseURLForCanonicalHeader string `json:"baseUrlForCanonicalHeader"`
	// This field is from variant [AccountOriginGetResponseS3Compatible].
	Endpoint string `json:"endpoint"`
	// This field is from variant [AccountOriginGetResponseS3Compatible].
	S3ForcePathStyle bool   `json:"s3ForcePathStyle"`
	BaseURL          string `json:"baseUrl"`
	// This field is from variant [AccountOriginGetResponseWebFolder].
	ForwardHostHeaderToOrigin bool `json:"forwardHostHeaderToOrigin"`
	// This field is from variant [AccountOriginGetResponseGoogleCloudStorageGcs].
	ClientEmail string `json:"clientEmail"`
	// This field is from variant [AccountOriginGetResponseAzureBlobStorage].
	AccountName string `json:"accountName"`
	// This field is from variant [AccountOriginGetResponseAzureBlobStorage].
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

func (u AccountOriginGetResponseUnion) AsS3() (v AccountOriginGetResponseS3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginGetResponseUnion) AsS3Compatible() (v AccountOriginGetResponseS3Compatible) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginGetResponseUnion) AsCloudinaryBackup() (v AccountOriginGetResponseCloudinaryBackup) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginGetResponseUnion) AsWebFolder() (v AccountOriginGetResponseWebFolder) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginGetResponseUnion) AsWebProxy() (v AccountOriginGetResponseWebProxy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginGetResponseUnion) AsGoogleCloudStorageGcs() (v AccountOriginGetResponseGoogleCloudStorageGcs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginGetResponseUnion) AsAzureBlobStorage() (v AccountOriginGetResponseAzureBlobStorage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginGetResponseUnion) AsAkeneoPim() (v AccountOriginGetResponseAkeneoPim) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccountOriginGetResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *AccountOriginGetResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginGetResponseS3 struct {
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
func (r AccountOriginGetResponseS3) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginGetResponseS3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginGetResponseS3Compatible struct {
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
func (r AccountOriginGetResponseS3Compatible) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginGetResponseS3Compatible) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginGetResponseCloudinaryBackup struct {
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
func (r AccountOriginGetResponseCloudinaryBackup) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginGetResponseCloudinaryBackup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginGetResponseWebFolder struct {
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
func (r AccountOriginGetResponseWebFolder) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginGetResponseWebFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginGetResponseWebProxy struct {
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
func (r AccountOriginGetResponseWebProxy) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginGetResponseWebProxy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginGetResponseGoogleCloudStorageGcs struct {
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
func (r AccountOriginGetResponseGoogleCloudStorageGcs) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginGetResponseGoogleCloudStorageGcs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginGetResponseAzureBlobStorage struct {
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
func (r AccountOriginGetResponseAzureBlobStorage) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginGetResponseAzureBlobStorage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginGetResponseAkeneoPim struct {
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
func (r AccountOriginGetResponseAkeneoPim) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginGetResponseAkeneoPim) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfS3 *AccountOriginNewParamsBodyS3 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfS3Compatible *AccountOriginNewParamsBodyS3Compatible `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfCloudinaryBackup *AccountOriginNewParamsBodyCloudinaryBackup `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfWebFolder *AccountOriginNewParamsBodyWebFolder `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfWebProxy *AccountOriginNewParamsBodyWebProxy `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfGcs *AccountOriginNewParamsBodyGcs `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfAzureBlob *AccountOriginNewParamsBodyAzureBlob `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfAkeneoPim *AccountOriginNewParamsBodyAkeneoPim `json:",inline"`

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
type AccountOriginNewParamsBodyS3 struct {
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

func (r AccountOriginNewParamsBodyS3) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginNewParamsBodyS3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginNewParamsBodyS3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessKey, Bucket, Endpoint, Name, SecretKey, Type are required.
type AccountOriginNewParamsBodyS3Compatible struct {
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

func (r AccountOriginNewParamsBodyS3Compatible) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginNewParamsBodyS3Compatible
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginNewParamsBodyS3Compatible) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessKey, Bucket, Name, SecretKey, Type are required.
type AccountOriginNewParamsBodyCloudinaryBackup struct {
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

func (r AccountOriginNewParamsBodyCloudinaryBackup) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginNewParamsBodyCloudinaryBackup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginNewParamsBodyCloudinaryBackup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties BaseURL, Name, Type are required.
type AccountOriginNewParamsBodyWebFolder struct {
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

func (r AccountOriginNewParamsBodyWebFolder) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginNewParamsBodyWebFolder
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginNewParamsBodyWebFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Type are required.
type AccountOriginNewParamsBodyWebProxy struct {
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

func (r AccountOriginNewParamsBodyWebProxy) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginNewParamsBodyWebProxy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginNewParamsBodyWebProxy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Bucket, ClientEmail, Name, PrivateKey, Type are required.
type AccountOriginNewParamsBodyGcs struct {
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

func (r AccountOriginNewParamsBodyGcs) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginNewParamsBodyGcs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginNewParamsBodyGcs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccountName, Container, Name, SasToken, Type are required.
type AccountOriginNewParamsBodyAzureBlob struct {
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

func (r AccountOriginNewParamsBodyAzureBlob) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginNewParamsBodyAzureBlob
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginNewParamsBodyAzureBlob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties BaseURL, ClientID, ClientSecret, Name, Password, Type, Username
// are required.
type AccountOriginNewParamsBodyAkeneoPim struct {
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

func (r AccountOriginNewParamsBodyAkeneoPim) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginNewParamsBodyAkeneoPim
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginNewParamsBodyAkeneoPim) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginUpdateParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfS3 *AccountOriginUpdateParamsBodyS3 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfS3Compatible *AccountOriginUpdateParamsBodyS3Compatible `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfCloudinaryBackup *AccountOriginUpdateParamsBodyCloudinaryBackup `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfWebFolder *AccountOriginUpdateParamsBodyWebFolder `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfWebProxy *AccountOriginUpdateParamsBodyWebProxy `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfGcs *AccountOriginUpdateParamsBodyGcs `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfAzureBlob *AccountOriginUpdateParamsBodyAzureBlob `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfAkeneoPim *AccountOriginUpdateParamsBodyAkeneoPim `json:",inline"`

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
type AccountOriginUpdateParamsBodyS3 struct {
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

func (r AccountOriginUpdateParamsBodyS3) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginUpdateParamsBodyS3
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginUpdateParamsBodyS3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessKey, Bucket, Endpoint, Name, SecretKey, Type are required.
type AccountOriginUpdateParamsBodyS3Compatible struct {
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

func (r AccountOriginUpdateParamsBodyS3Compatible) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginUpdateParamsBodyS3Compatible
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginUpdateParamsBodyS3Compatible) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessKey, Bucket, Name, SecretKey, Type are required.
type AccountOriginUpdateParamsBodyCloudinaryBackup struct {
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

func (r AccountOriginUpdateParamsBodyCloudinaryBackup) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginUpdateParamsBodyCloudinaryBackup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginUpdateParamsBodyCloudinaryBackup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties BaseURL, Name, Type are required.
type AccountOriginUpdateParamsBodyWebFolder struct {
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

func (r AccountOriginUpdateParamsBodyWebFolder) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginUpdateParamsBodyWebFolder
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginUpdateParamsBodyWebFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Type are required.
type AccountOriginUpdateParamsBodyWebProxy struct {
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

func (r AccountOriginUpdateParamsBodyWebProxy) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginUpdateParamsBodyWebProxy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginUpdateParamsBodyWebProxy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Bucket, ClientEmail, Name, PrivateKey, Type are required.
type AccountOriginUpdateParamsBodyGcs struct {
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

func (r AccountOriginUpdateParamsBodyGcs) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginUpdateParamsBodyGcs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginUpdateParamsBodyGcs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccountName, Container, Name, SasToken, Type are required.
type AccountOriginUpdateParamsBodyAzureBlob struct {
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

func (r AccountOriginUpdateParamsBodyAzureBlob) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginUpdateParamsBodyAzureBlob
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginUpdateParamsBodyAzureBlob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties BaseURL, ClientID, ClientSecret, Name, Password, Type, Username
// are required.
type AccountOriginUpdateParamsBodyAkeneoPim struct {
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

func (r AccountOriginUpdateParamsBodyAkeneoPim) MarshalJSON() (data []byte, err error) {
	type shadow AccountOriginUpdateParamsBodyAkeneoPim
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountOriginUpdateParamsBodyAkeneoPim) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
