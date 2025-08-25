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
func (r *AccountOriginService) New(ctx context.Context, body AccountOriginNewParams, opts ...option.RequestOption) (res *OriginUnion, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/accounts/origins"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// **Note:** This API is currently in beta.
// Updates the origin identified by `id` and returns the updated origin object.
func (r *AccountOriginService) Update(ctx context.Context, id string, body AccountOriginUpdateParams, opts ...option.RequestOption) (res *OriginUnion, err error) {
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
func (r *AccountOriginService) List(ctx context.Context, opts ...option.RequestOption) (res *[]OriginUnion, err error) {
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
func (r *AccountOriginService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *OriginUnion, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/accounts/origins/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// OriginUnion contains all possible properties and values from [OriginS3],
// [OriginS3Compatible], [OriginCloudinaryBackup], [OriginWebFolder],
// [OriginWebProxy], [OriginGcs], [OriginAzureBlob], [OriginAkeneoPim].
//
// Use the [OriginUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type OriginUnion struct {
	AccessKey string `json:"accessKey"`
	Bucket    string `json:"bucket"`
	Name      string `json:"name"`
	SecretKey string `json:"secretKey"`
	// Any of "S3", "S3_COMPATIBLE", "CLOUDINARY_BACKUP", "WEB_FOLDER", "WEB_PROXY",
	// "GCS", "AZURE_BLOB", "AKENEO_PIM".
	Type                      string `json:"type"`
	ID                        string `json:"id"`
	BaseURLForCanonicalHeader string `json:"baseUrlForCanonicalHeader"`
	IncludeCanonicalHeader    bool   `json:"includeCanonicalHeader"`
	Prefix                    string `json:"prefix"`
	// This field is from variant [OriginS3Compatible].
	Endpoint string `json:"endpoint"`
	// This field is from variant [OriginS3Compatible].
	S3ForcePathStyle bool   `json:"s3ForcePathStyle"`
	BaseURL          string `json:"baseUrl"`
	// This field is from variant [OriginWebFolder].
	ForwardHostHeaderToOrigin bool `json:"forwardHostHeaderToOrigin"`
	// This field is from variant [OriginGcs].
	ClientEmail string `json:"clientEmail"`
	// This field is from variant [OriginGcs].
	PrivateKey string `json:"privateKey"`
	// This field is from variant [OriginAzureBlob].
	AccountName string `json:"accountName"`
	// This field is from variant [OriginAzureBlob].
	Container string `json:"container"`
	// This field is from variant [OriginAzureBlob].
	SasToken string `json:"sasToken"`
	// This field is from variant [OriginAkeneoPim].
	ClientID string `json:"clientId"`
	// This field is from variant [OriginAkeneoPim].
	ClientSecret string `json:"clientSecret"`
	// This field is from variant [OriginAkeneoPim].
	Password string `json:"password"`
	// This field is from variant [OriginAkeneoPim].
	Username string `json:"username"`
	JSON     struct {
		AccessKey                 respjson.Field
		Bucket                    respjson.Field
		Name                      respjson.Field
		SecretKey                 respjson.Field
		Type                      respjson.Field
		ID                        respjson.Field
		BaseURLForCanonicalHeader respjson.Field
		IncludeCanonicalHeader    respjson.Field
		Prefix                    respjson.Field
		Endpoint                  respjson.Field
		S3ForcePathStyle          respjson.Field
		BaseURL                   respjson.Field
		ForwardHostHeaderToOrigin respjson.Field
		ClientEmail               respjson.Field
		PrivateKey                respjson.Field
		AccountName               respjson.Field
		Container                 respjson.Field
		SasToken                  respjson.Field
		ClientID                  respjson.Field
		ClientSecret              respjson.Field
		Password                  respjson.Field
		Username                  respjson.Field
		raw                       string
	} `json:"-"`
}

// anyOrigin is implemented by each variant of [OriginUnion] to add type safety for
// the return type of [OriginUnion.AsAny]
type anyOrigin interface {
	implOriginUnion()
}

func (OriginS3) implOriginUnion()               {}
func (OriginS3Compatible) implOriginUnion()     {}
func (OriginCloudinaryBackup) implOriginUnion() {}
func (OriginWebFolder) implOriginUnion()        {}
func (OriginWebProxy) implOriginUnion()         {}
func (OriginGcs) implOriginUnion()              {}
func (OriginAzureBlob) implOriginUnion()        {}
func (OriginAkeneoPim) implOriginUnion()        {}

// Use the following switch statement to find the correct variant
//
//	switch variant := OriginUnion.AsAny().(type) {
//	case imagekit.OriginS3:
//	case imagekit.OriginS3Compatible:
//	case imagekit.OriginCloudinaryBackup:
//	case imagekit.OriginWebFolder:
//	case imagekit.OriginWebProxy:
//	case imagekit.OriginGcs:
//	case imagekit.OriginAzureBlob:
//	case imagekit.OriginAkeneoPim:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u OriginUnion) AsAny() anyOrigin {
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

func (u OriginUnion) AsS3() (v OriginS3) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OriginUnion) AsS3Compatible() (v OriginS3Compatible) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OriginUnion) AsCloudinaryBackup() (v OriginCloudinaryBackup) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OriginUnion) AsWebFolder() (v OriginWebFolder) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OriginUnion) AsWebProxy() (v OriginWebProxy) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OriginUnion) AsGcs() (v OriginGcs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OriginUnion) AsAzureBlob() (v OriginAzureBlob) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u OriginUnion) AsAkeneoPim() (v OriginAkeneoPim) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u OriginUnion) RawJSON() string { return u.JSON.raw }

func (r *OriginUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OriginUnion to a OriginUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OriginUnionParam.Overrides()
func (r OriginUnion) ToParam() OriginUnionParam {
	return param.Override[OriginUnionParam](json.RawMessage(r.RawJSON()))
}

type OriginS3 struct {
	// Access key for the bucket.
	AccessKey string `json:"accessKey,required"`
	// S3 bucket name.
	Bucket string `json:"bucket,required"`
	// Display name of the origin.
	Name string `json:"name,required"`
	// Secret key for the bucket.
	SecretKey string      `json:"secretKey,required"`
	Type      constant.S3 `json:"type,required"`
	// Unique identifier for the origin. This is generated by ImageKit when you create
	// a new origin.
	ID string `json:"id"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader string `json:"baseUrlForCanonicalHeader" format:"uri"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader bool `json:"includeCanonicalHeader"`
	// Path prefix inside the bucket.
	Prefix string `json:"prefix"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessKey                 respjson.Field
		Bucket                    respjson.Field
		Name                      respjson.Field
		SecretKey                 respjson.Field
		Type                      respjson.Field
		ID                        respjson.Field
		BaseURLForCanonicalHeader respjson.Field
		IncludeCanonicalHeader    respjson.Field
		Prefix                    respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OriginS3) RawJSON() string { return r.JSON.raw }
func (r *OriginS3) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OriginS3Compatible struct {
	// Access key for the bucket.
	AccessKey string `json:"accessKey,required"`
	// S3 bucket name.
	Bucket string `json:"bucket,required"`
	// Custom S3-compatible endpoint.
	Endpoint string `json:"endpoint,required" format:"uri"`
	// Display name of the origin.
	Name string `json:"name,required"`
	// Secret key for the bucket.
	SecretKey string                `json:"secretKey,required"`
	Type      constant.S3Compatible `json:"type,required"`
	// Unique identifier for the origin. This is generated by ImageKit when you create
	// a new origin.
	ID string `json:"id"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader string `json:"baseUrlForCanonicalHeader" format:"uri"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader bool `json:"includeCanonicalHeader"`
	// Path prefix inside the bucket.
	Prefix string `json:"prefix"`
	// Use path-style S3 URLs?
	S3ForcePathStyle bool `json:"s3ForcePathStyle"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessKey                 respjson.Field
		Bucket                    respjson.Field
		Endpoint                  respjson.Field
		Name                      respjson.Field
		SecretKey                 respjson.Field
		Type                      respjson.Field
		ID                        respjson.Field
		BaseURLForCanonicalHeader respjson.Field
		IncludeCanonicalHeader    respjson.Field
		Prefix                    respjson.Field
		S3ForcePathStyle          respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OriginS3Compatible) RawJSON() string { return r.JSON.raw }
func (r *OriginS3Compatible) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OriginCloudinaryBackup struct {
	// Access key for the bucket.
	AccessKey string `json:"accessKey,required"`
	// S3 bucket name.
	Bucket string `json:"bucket,required"`
	// Display name of the origin.
	Name string `json:"name,required"`
	// Secret key for the bucket.
	SecretKey string                    `json:"secretKey,required"`
	Type      constant.CloudinaryBackup `json:"type,required"`
	// Unique identifier for the origin. This is generated by ImageKit when you create
	// a new origin.
	ID string `json:"id"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader string `json:"baseUrlForCanonicalHeader" format:"uri"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader bool `json:"includeCanonicalHeader"`
	// Path prefix inside the bucket.
	Prefix string `json:"prefix"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessKey                 respjson.Field
		Bucket                    respjson.Field
		Name                      respjson.Field
		SecretKey                 respjson.Field
		Type                      respjson.Field
		ID                        respjson.Field
		BaseURLForCanonicalHeader respjson.Field
		IncludeCanonicalHeader    respjson.Field
		Prefix                    respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OriginCloudinaryBackup) RawJSON() string { return r.JSON.raw }
func (r *OriginCloudinaryBackup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OriginWebFolder struct {
	// Root URL for the web folder origin.
	BaseURL string `json:"baseUrl,required" format:"uri"`
	// Display name of the origin.
	Name string             `json:"name,required"`
	Type constant.WebFolder `json:"type,required"`
	// Unique identifier for the origin. This is generated by ImageKit when you create
	// a new origin.
	ID string `json:"id"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader string `json:"baseUrlForCanonicalHeader" format:"uri"`
	// Forward the Host header to origin?
	ForwardHostHeaderToOrigin bool `json:"forwardHostHeaderToOrigin"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader bool `json:"includeCanonicalHeader"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaseURL                   respjson.Field
		Name                      respjson.Field
		Type                      respjson.Field
		ID                        respjson.Field
		BaseURLForCanonicalHeader respjson.Field
		ForwardHostHeaderToOrigin respjson.Field
		IncludeCanonicalHeader    respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OriginWebFolder) RawJSON() string { return r.JSON.raw }
func (r *OriginWebFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OriginWebProxy struct {
	// Display name of the origin.
	Name string            `json:"name,required"`
	Type constant.WebProxy `json:"type,required"`
	// Unique identifier for the origin. This is generated by ImageKit when you create
	// a new origin.
	ID string `json:"id"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader string `json:"baseUrlForCanonicalHeader" format:"uri"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader bool `json:"includeCanonicalHeader"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name                      respjson.Field
		Type                      respjson.Field
		ID                        respjson.Field
		BaseURLForCanonicalHeader respjson.Field
		IncludeCanonicalHeader    respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OriginWebProxy) RawJSON() string { return r.JSON.raw }
func (r *OriginWebProxy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OriginGcs struct {
	Bucket      string `json:"bucket,required"`
	ClientEmail string `json:"clientEmail,required" format:"email"`
	// Display name of the origin.
	Name       string       `json:"name,required"`
	PrivateKey string       `json:"privateKey,required"`
	Type       constant.Gcs `json:"type,required"`
	// Unique identifier for the origin. This is generated by ImageKit when you create
	// a new origin.
	ID string `json:"id"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader string `json:"baseUrlForCanonicalHeader" format:"uri"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader bool   `json:"includeCanonicalHeader"`
	Prefix                 string `json:"prefix"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bucket                    respjson.Field
		ClientEmail               respjson.Field
		Name                      respjson.Field
		PrivateKey                respjson.Field
		Type                      respjson.Field
		ID                        respjson.Field
		BaseURLForCanonicalHeader respjson.Field
		IncludeCanonicalHeader    respjson.Field
		Prefix                    respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OriginGcs) RawJSON() string { return r.JSON.raw }
func (r *OriginGcs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OriginAzureBlob struct {
	AccountName string `json:"accountName,required"`
	Container   string `json:"container,required"`
	// Display name of the origin.
	Name     string             `json:"name,required"`
	SasToken string             `json:"sasToken,required"`
	Type     constant.AzureBlob `json:"type,required"`
	// Unique identifier for the origin. This is generated by ImageKit when you create
	// a new origin.
	ID string `json:"id"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader string `json:"baseUrlForCanonicalHeader" format:"uri"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader bool   `json:"includeCanonicalHeader"`
	Prefix                 string `json:"prefix"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountName               respjson.Field
		Container                 respjson.Field
		Name                      respjson.Field
		SasToken                  respjson.Field
		Type                      respjson.Field
		ID                        respjson.Field
		BaseURLForCanonicalHeader respjson.Field
		IncludeCanonicalHeader    respjson.Field
		Prefix                    respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OriginAzureBlob) RawJSON() string { return r.JSON.raw }
func (r *OriginAzureBlob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OriginAkeneoPim struct {
	// Akeneo instance base URL.
	BaseURL string `json:"baseUrl,required" format:"uri"`
	// Akeneo API client ID.
	ClientID string `json:"clientId,required"`
	// Akeneo API client secret.
	ClientSecret string `json:"clientSecret,required"`
	// Display name of the origin.
	Name string `json:"name,required"`
	// Akeneo API password.
	Password string             `json:"password,required"`
	Type     constant.AkeneoPim `json:"type,required"`
	// Akeneo API username.
	Username string `json:"username,required"`
	// Unique identifier for the origin. This is generated by ImageKit when you create
	// a new origin.
	ID string `json:"id"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader string `json:"baseUrlForCanonicalHeader" format:"uri"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader bool `json:"includeCanonicalHeader"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BaseURL                   respjson.Field
		ClientID                  respjson.Field
		ClientSecret              respjson.Field
		Name                      respjson.Field
		Password                  respjson.Field
		Type                      respjson.Field
		Username                  respjson.Field
		ID                        respjson.Field
		BaseURLForCanonicalHeader respjson.Field
		IncludeCanonicalHeader    respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OriginAkeneoPim) RawJSON() string { return r.JSON.raw }
func (r *OriginAkeneoPim) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func OriginParamOfWebFolder(baseURL string, name string) OriginUnionParam {
	var webFolder OriginWebFolderParam
	webFolder.BaseURL = baseURL
	webFolder.Name = name
	return OriginUnionParam{OfWebFolder: &webFolder}
}

func OriginParamOfWebProxy(name string) OriginUnionParam {
	var webProxy OriginWebProxyParam
	webProxy.Name = name
	return OriginUnionParam{OfWebProxy: &webProxy}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type OriginUnionParam struct {
	OfS3               *OriginS3Param               `json:",omitzero,inline"`
	OfS3Compatible     *OriginS3CompatibleParam     `json:",omitzero,inline"`
	OfCloudinaryBackup *OriginCloudinaryBackupParam `json:",omitzero,inline"`
	OfWebFolder        *OriginWebFolderParam        `json:",omitzero,inline"`
	OfWebProxy         *OriginWebProxyParam         `json:",omitzero,inline"`
	OfGcs              *OriginGcsParam              `json:",omitzero,inline"`
	OfAzureBlob        *OriginAzureBlobParam        `json:",omitzero,inline"`
	OfAkeneoPim        *OriginAkeneoPimParam        `json:",omitzero,inline"`
	paramUnion
}

func (u OriginUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfS3,
		u.OfS3Compatible,
		u.OfCloudinaryBackup,
		u.OfWebFolder,
		u.OfWebProxy,
		u.OfGcs,
		u.OfAzureBlob,
		u.OfAkeneoPim)
}
func (u *OriginUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *OriginUnionParam) asAny() any {
	if !param.IsOmitted(u.OfS3) {
		return u.OfS3
	} else if !param.IsOmitted(u.OfS3Compatible) {
		return u.OfS3Compatible
	} else if !param.IsOmitted(u.OfCloudinaryBackup) {
		return u.OfCloudinaryBackup
	} else if !param.IsOmitted(u.OfWebFolder) {
		return u.OfWebFolder
	} else if !param.IsOmitted(u.OfWebProxy) {
		return u.OfWebProxy
	} else if !param.IsOmitted(u.OfGcs) {
		return u.OfGcs
	} else if !param.IsOmitted(u.OfAzureBlob) {
		return u.OfAzureBlob
	} else if !param.IsOmitted(u.OfAkeneoPim) {
		return u.OfAkeneoPim
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginUnionParam) GetEndpoint() *string {
	if vt := u.OfS3Compatible; vt != nil {
		return &vt.Endpoint
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginUnionParam) GetS3ForcePathStyle() *bool {
	if vt := u.OfS3Compatible; vt != nil && vt.S3ForcePathStyle.Valid() {
		return &vt.S3ForcePathStyle.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginUnionParam) GetForwardHostHeaderToOrigin() *bool {
	if vt := u.OfWebFolder; vt != nil && vt.ForwardHostHeaderToOrigin.Valid() {
		return &vt.ForwardHostHeaderToOrigin.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginUnionParam) GetClientEmail() *string {
	if vt := u.OfGcs; vt != nil {
		return &vt.ClientEmail
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginUnionParam) GetPrivateKey() *string {
	if vt := u.OfGcs; vt != nil {
		return &vt.PrivateKey
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginUnionParam) GetAccountName() *string {
	if vt := u.OfAzureBlob; vt != nil {
		return &vt.AccountName
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginUnionParam) GetContainer() *string {
	if vt := u.OfAzureBlob; vt != nil {
		return &vt.Container
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginUnionParam) GetSasToken() *string {
	if vt := u.OfAzureBlob; vt != nil {
		return &vt.SasToken
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginUnionParam) GetClientID() *string {
	if vt := u.OfAkeneoPim; vt != nil {
		return &vt.ClientID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginUnionParam) GetClientSecret() *string {
	if vt := u.OfAkeneoPim; vt != nil {
		return &vt.ClientSecret
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginUnionParam) GetPassword() *string {
	if vt := u.OfAkeneoPim; vt != nil {
		return &vt.Password
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginUnionParam) GetUsername() *string {
	if vt := u.OfAkeneoPim; vt != nil {
		return &vt.Username
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginUnionParam) GetAccessKey() *string {
	if vt := u.OfS3; vt != nil {
		return (*string)(&vt.AccessKey)
	} else if vt := u.OfS3Compatible; vt != nil {
		return (*string)(&vt.AccessKey)
	} else if vt := u.OfCloudinaryBackup; vt != nil {
		return (*string)(&vt.AccessKey)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginUnionParam) GetBucket() *string {
	if vt := u.OfS3; vt != nil {
		return (*string)(&vt.Bucket)
	} else if vt := u.OfS3Compatible; vt != nil {
		return (*string)(&vt.Bucket)
	} else if vt := u.OfCloudinaryBackup; vt != nil {
		return (*string)(&vt.Bucket)
	} else if vt := u.OfGcs; vt != nil {
		return (*string)(&vt.Bucket)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginUnionParam) GetName() *string {
	if vt := u.OfS3; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfS3Compatible; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfCloudinaryBackup; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfWebFolder; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfWebProxy; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfGcs; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfAzureBlob; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfAkeneoPim; vt != nil {
		return (*string)(&vt.Name)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginUnionParam) GetSecretKey() *string {
	if vt := u.OfS3; vt != nil {
		return (*string)(&vt.SecretKey)
	} else if vt := u.OfS3Compatible; vt != nil {
		return (*string)(&vt.SecretKey)
	} else if vt := u.OfCloudinaryBackup; vt != nil {
		return (*string)(&vt.SecretKey)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginUnionParam) GetType() *string {
	if vt := u.OfS3; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfS3Compatible; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCloudinaryBackup; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWebFolder; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWebProxy; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfGcs; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAzureBlob; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAkeneoPim; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginUnionParam) GetID() *string {
	if vt := u.OfS3; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfS3Compatible; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfCloudinaryBackup; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfWebFolder; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfWebProxy; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfGcs; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfAzureBlob; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	} else if vt := u.OfAkeneoPim; vt != nil && vt.ID.Valid() {
		return &vt.ID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginUnionParam) GetBaseURLForCanonicalHeader() *string {
	if vt := u.OfS3; vt != nil && vt.BaseURLForCanonicalHeader.Valid() {
		return &vt.BaseURLForCanonicalHeader.Value
	} else if vt := u.OfS3Compatible; vt != nil && vt.BaseURLForCanonicalHeader.Valid() {
		return &vt.BaseURLForCanonicalHeader.Value
	} else if vt := u.OfCloudinaryBackup; vt != nil && vt.BaseURLForCanonicalHeader.Valid() {
		return &vt.BaseURLForCanonicalHeader.Value
	} else if vt := u.OfWebFolder; vt != nil && vt.BaseURLForCanonicalHeader.Valid() {
		return &vt.BaseURLForCanonicalHeader.Value
	} else if vt := u.OfWebProxy; vt != nil && vt.BaseURLForCanonicalHeader.Valid() {
		return &vt.BaseURLForCanonicalHeader.Value
	} else if vt := u.OfGcs; vt != nil && vt.BaseURLForCanonicalHeader.Valid() {
		return &vt.BaseURLForCanonicalHeader.Value
	} else if vt := u.OfAzureBlob; vt != nil && vt.BaseURLForCanonicalHeader.Valid() {
		return &vt.BaseURLForCanonicalHeader.Value
	} else if vt := u.OfAkeneoPim; vt != nil && vt.BaseURLForCanonicalHeader.Valid() {
		return &vt.BaseURLForCanonicalHeader.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginUnionParam) GetIncludeCanonicalHeader() *bool {
	if vt := u.OfS3; vt != nil && vt.IncludeCanonicalHeader.Valid() {
		return &vt.IncludeCanonicalHeader.Value
	} else if vt := u.OfS3Compatible; vt != nil && vt.IncludeCanonicalHeader.Valid() {
		return &vt.IncludeCanonicalHeader.Value
	} else if vt := u.OfCloudinaryBackup; vt != nil && vt.IncludeCanonicalHeader.Valid() {
		return &vt.IncludeCanonicalHeader.Value
	} else if vt := u.OfWebFolder; vt != nil && vt.IncludeCanonicalHeader.Valid() {
		return &vt.IncludeCanonicalHeader.Value
	} else if vt := u.OfWebProxy; vt != nil && vt.IncludeCanonicalHeader.Valid() {
		return &vt.IncludeCanonicalHeader.Value
	} else if vt := u.OfGcs; vt != nil && vt.IncludeCanonicalHeader.Valid() {
		return &vt.IncludeCanonicalHeader.Value
	} else if vt := u.OfAzureBlob; vt != nil && vt.IncludeCanonicalHeader.Valid() {
		return &vt.IncludeCanonicalHeader.Value
	} else if vt := u.OfAkeneoPim; vt != nil && vt.IncludeCanonicalHeader.Valid() {
		return &vt.IncludeCanonicalHeader.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginUnionParam) GetPrefix() *string {
	if vt := u.OfS3; vt != nil && vt.Prefix.Valid() {
		return &vt.Prefix.Value
	} else if vt := u.OfS3Compatible; vt != nil && vt.Prefix.Valid() {
		return &vt.Prefix.Value
	} else if vt := u.OfCloudinaryBackup; vt != nil && vt.Prefix.Valid() {
		return &vt.Prefix.Value
	} else if vt := u.OfGcs; vt != nil && vt.Prefix.Valid() {
		return &vt.Prefix.Value
	} else if vt := u.OfAzureBlob; vt != nil && vt.Prefix.Valid() {
		return &vt.Prefix.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u OriginUnionParam) GetBaseURL() *string {
	if vt := u.OfWebFolder; vt != nil {
		return (*string)(&vt.BaseURL)
	} else if vt := u.OfAkeneoPim; vt != nil {
		return (*string)(&vt.BaseURL)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[OriginUnionParam](
		"type",
		apijson.Discriminator[OriginS3Param]("S3"),
		apijson.Discriminator[OriginS3CompatibleParam]("S3_COMPATIBLE"),
		apijson.Discriminator[OriginCloudinaryBackupParam]("CLOUDINARY_BACKUP"),
		apijson.Discriminator[OriginWebFolderParam]("WEB_FOLDER"),
		apijson.Discriminator[OriginWebProxyParam]("WEB_PROXY"),
		apijson.Discriminator[OriginGcsParam]("GCS"),
		apijson.Discriminator[OriginAzureBlobParam]("AZURE_BLOB"),
		apijson.Discriminator[OriginAkeneoPimParam]("AKENEO_PIM"),
	)
}

// The properties AccessKey, Bucket, Name, SecretKey, Type are required.
type OriginS3Param struct {
	// Access key for the bucket.
	AccessKey string `json:"accessKey,required"`
	// S3 bucket name.
	Bucket string `json:"bucket,required"`
	// Display name of the origin.
	Name string `json:"name,required"`
	// Secret key for the bucket.
	SecretKey string `json:"secretKey,required"`
	// Unique identifier for the origin. This is generated by ImageKit when you create
	// a new origin.
	ID param.Opt[string] `json:"id,omitzero"`
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

func (r OriginS3Param) MarshalJSON() (data []byte, err error) {
	type shadow OriginS3Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginS3Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessKey, Bucket, Endpoint, Name, SecretKey, Type are required.
type OriginS3CompatibleParam struct {
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
	// Unique identifier for the origin. This is generated by ImageKit when you create
	// a new origin.
	ID param.Opt[string] `json:"id,omitzero"`
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

func (r OriginS3CompatibleParam) MarshalJSON() (data []byte, err error) {
	type shadow OriginS3CompatibleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginS3CompatibleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccessKey, Bucket, Name, SecretKey, Type are required.
type OriginCloudinaryBackupParam struct {
	// Access key for the bucket.
	AccessKey string `json:"accessKey,required"`
	// S3 bucket name.
	Bucket string `json:"bucket,required"`
	// Display name of the origin.
	Name string `json:"name,required"`
	// Secret key for the bucket.
	SecretKey string `json:"secretKey,required"`
	// Unique identifier for the origin. This is generated by ImageKit when you create
	// a new origin.
	ID param.Opt[string] `json:"id,omitzero"`
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

func (r OriginCloudinaryBackupParam) MarshalJSON() (data []byte, err error) {
	type shadow OriginCloudinaryBackupParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginCloudinaryBackupParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties BaseURL, Name, Type are required.
type OriginWebFolderParam struct {
	// Root URL for the web folder origin.
	BaseURL string `json:"baseUrl,required" format:"uri"`
	// Display name of the origin.
	Name string `json:"name,required"`
	// Unique identifier for the origin. This is generated by ImageKit when you create
	// a new origin.
	ID param.Opt[string] `json:"id,omitzero"`
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

func (r OriginWebFolderParam) MarshalJSON() (data []byte, err error) {
	type shadow OriginWebFolderParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginWebFolderParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Type are required.
type OriginWebProxyParam struct {
	// Display name of the origin.
	Name string `json:"name,required"`
	// Unique identifier for the origin. This is generated by ImageKit when you create
	// a new origin.
	ID param.Opt[string] `json:"id,omitzero"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader param.Opt[string] `json:"baseUrlForCanonicalHeader,omitzero" format:"uri"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader param.Opt[bool] `json:"includeCanonicalHeader,omitzero"`
	// This field can be elided, and will marshal its zero value as "WEB_PROXY".
	Type constant.WebProxy `json:"type,required"`
	paramObj
}

func (r OriginWebProxyParam) MarshalJSON() (data []byte, err error) {
	type shadow OriginWebProxyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginWebProxyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Bucket, ClientEmail, Name, PrivateKey, Type are required.
type OriginGcsParam struct {
	Bucket      string `json:"bucket,required"`
	ClientEmail string `json:"clientEmail,required" format:"email"`
	// Display name of the origin.
	Name       string `json:"name,required"`
	PrivateKey string `json:"privateKey,required"`
	// Unique identifier for the origin. This is generated by ImageKit when you create
	// a new origin.
	ID param.Opt[string] `json:"id,omitzero"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader param.Opt[string] `json:"baseUrlForCanonicalHeader,omitzero" format:"uri"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader param.Opt[bool]   `json:"includeCanonicalHeader,omitzero"`
	Prefix                 param.Opt[string] `json:"prefix,omitzero"`
	// This field can be elided, and will marshal its zero value as "GCS".
	Type constant.Gcs `json:"type,required"`
	paramObj
}

func (r OriginGcsParam) MarshalJSON() (data []byte, err error) {
	type shadow OriginGcsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginGcsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AccountName, Container, Name, SasToken, Type are required.
type OriginAzureBlobParam struct {
	AccountName string `json:"accountName,required"`
	Container   string `json:"container,required"`
	// Display name of the origin.
	Name     string `json:"name,required"`
	SasToken string `json:"sasToken,required"`
	// Unique identifier for the origin. This is generated by ImageKit when you create
	// a new origin.
	ID param.Opt[string] `json:"id,omitzero"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader param.Opt[string] `json:"baseUrlForCanonicalHeader,omitzero" format:"uri"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader param.Opt[bool]   `json:"includeCanonicalHeader,omitzero"`
	Prefix                 param.Opt[string] `json:"prefix,omitzero"`
	// This field can be elided, and will marshal its zero value as "AZURE_BLOB".
	Type constant.AzureBlob `json:"type,required"`
	paramObj
}

func (r OriginAzureBlobParam) MarshalJSON() (data []byte, err error) {
	type shadow OriginAzureBlobParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginAzureBlobParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties BaseURL, ClientID, ClientSecret, Name, Password, Type, Username
// are required.
type OriginAkeneoPimParam struct {
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
	// Unique identifier for the origin. This is generated by ImageKit when you create
	// a new origin.
	ID param.Opt[string] `json:"id,omitzero"`
	// URL used in the Canonical header (if enabled).
	BaseURLForCanonicalHeader param.Opt[string] `json:"baseUrlForCanonicalHeader,omitzero" format:"uri"`
	// Whether to send a Canonical header.
	IncludeCanonicalHeader param.Opt[bool] `json:"includeCanonicalHeader,omitzero"`
	// This field can be elided, and will marshal its zero value as "AKENEO_PIM".
	Type constant.AkeneoPim `json:"type,required"`
	paramObj
}

func (r OriginAkeneoPimParam) MarshalJSON() (data []byte, err error) {
	type shadow OriginAkeneoPimParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OriginAkeneoPimParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginNewParams struct {
	// Schema for origin resources.
	Origin OriginUnionParam
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

type AccountOriginUpdateParams struct {
	// Schema for origin resources.
	Origin OriginUnionParam
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
