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

// AccountOriginNewResponseUnion contains all possible properties and values from
// [AccountOriginNewResponseS3], [AccountOriginNewResponseS3Compatible],
// [AccountOriginNewResponseCloudinaryBackup], [AccountOriginNewResponseWebFolder],
// [AccountOriginNewResponseWebProxy], [AccountOriginNewResponseGcs],
// [AccountOriginNewResponseAzureBlob], [AccountOriginNewResponseAkeneoPim].
//
// Use the [AccountOriginNewResponseUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AccountOriginNewResponseUnion struct {
	ID                     string `json:"id"`
	Bucket                 string `json:"bucket"`
	IncludeCanonicalHeader bool   `json:"includeCanonicalHeader"`
	Name                   string `json:"name"`
	Prefix                 string `json:"prefix"`
	// Any of "S3", "S3_COMPATIBLE", "CLOUDINARY_BACKUP", "WEB_FOLDER", "WEB_PROXY",
	// "GCS", "AZURE_BLOB", "AKENEO_PIM".
	Type                      string `json:"type"`
	BaseURLForCanonicalHeader string `json:"baseUrlForCanonicalHeader"`
	// This field is from variant [AccountOriginNewResponseS3Compatible].
	Endpoint string `json:"endpoint"`
	// This field is from variant [AccountOriginNewResponseS3Compatible].
	S3ForcePathStyle bool   `json:"s3ForcePathStyle"`
	BaseURL          string `json:"baseUrl"`
	// This field is from variant [AccountOriginNewResponseWebFolder].
	ForwardHostHeaderToOrigin bool `json:"forwardHostHeaderToOrigin"`
	// This field is from variant [AccountOriginNewResponseGcs].
	ClientEmail string `json:"clientEmail"`
	// This field is from variant [AccountOriginNewResponseAzureBlob].
	AccountName string `json:"accountName"`
	// This field is from variant [AccountOriginNewResponseAzureBlob].
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

// anyAccountOriginNewResponse is implemented by each variant of
// [AccountOriginNewResponseUnion] to add type safety for the return type of
// [AccountOriginNewResponseUnion.AsAny]
type anyAccountOriginNewResponse interface {
	implAccountOriginNewResponseUnion()
}

func (AccountOriginNewResponseS3) implAccountOriginNewResponseUnion()               {}
func (AccountOriginNewResponseS3Compatible) implAccountOriginNewResponseUnion()     {}
func (AccountOriginNewResponseCloudinaryBackup) implAccountOriginNewResponseUnion() {}
func (AccountOriginNewResponseWebFolder) implAccountOriginNewResponseUnion()        {}
func (AccountOriginNewResponseWebProxy) implAccountOriginNewResponseUnion()         {}
func (AccountOriginNewResponseGcs) implAccountOriginNewResponseUnion()              {}
func (AccountOriginNewResponseAzureBlob) implAccountOriginNewResponseUnion()        {}
func (AccountOriginNewResponseAkeneoPim) implAccountOriginNewResponseUnion()        {}

// Use the following switch statement to find the correct variant
//
//	switch variant := AccountOriginNewResponseUnion.AsAny().(type) {
//	case imagekit.AccountOriginNewResponseS3:
//	case imagekit.AccountOriginNewResponseS3Compatible:
//	case imagekit.AccountOriginNewResponseCloudinaryBackup:
//	case imagekit.AccountOriginNewResponseWebFolder:
//	case imagekit.AccountOriginNewResponseWebProxy:
//	case imagekit.AccountOriginNewResponseGcs:
//	case imagekit.AccountOriginNewResponseAzureBlob:
//	case imagekit.AccountOriginNewResponseAkeneoPim:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u AccountOriginNewResponseUnion) AsAny() anyAccountOriginNewResponse {
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

func (u AccountOriginNewResponseUnion) AsGcs() (v AccountOriginNewResponseGcs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginNewResponseUnion) AsAzureBlob() (v AccountOriginNewResponseAzureBlob) {
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

type AccountOriginNewResponseGcs struct {
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
func (r AccountOriginNewResponseGcs) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginNewResponseGcs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginNewResponseAzureBlob struct {
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
func (r AccountOriginNewResponseAzureBlob) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginNewResponseAzureBlob) UnmarshalJSON(data []byte) error {
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
// [AccountOriginUpdateResponseGcs], [AccountOriginUpdateResponseAzureBlob],
// [AccountOriginUpdateResponseAkeneoPim].
//
// Use the [AccountOriginUpdateResponseUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AccountOriginUpdateResponseUnion struct {
	ID                     string `json:"id"`
	Bucket                 string `json:"bucket"`
	IncludeCanonicalHeader bool   `json:"includeCanonicalHeader"`
	Name                   string `json:"name"`
	Prefix                 string `json:"prefix"`
	// Any of "S3", "S3_COMPATIBLE", "CLOUDINARY_BACKUP", "WEB_FOLDER", "WEB_PROXY",
	// "GCS", "AZURE_BLOB", "AKENEO_PIM".
	Type                      string `json:"type"`
	BaseURLForCanonicalHeader string `json:"baseUrlForCanonicalHeader"`
	// This field is from variant [AccountOriginUpdateResponseS3Compatible].
	Endpoint string `json:"endpoint"`
	// This field is from variant [AccountOriginUpdateResponseS3Compatible].
	S3ForcePathStyle bool   `json:"s3ForcePathStyle"`
	BaseURL          string `json:"baseUrl"`
	// This field is from variant [AccountOriginUpdateResponseWebFolder].
	ForwardHostHeaderToOrigin bool `json:"forwardHostHeaderToOrigin"`
	// This field is from variant [AccountOriginUpdateResponseGcs].
	ClientEmail string `json:"clientEmail"`
	// This field is from variant [AccountOriginUpdateResponseAzureBlob].
	AccountName string `json:"accountName"`
	// This field is from variant [AccountOriginUpdateResponseAzureBlob].
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

// anyAccountOriginUpdateResponse is implemented by each variant of
// [AccountOriginUpdateResponseUnion] to add type safety for the return type of
// [AccountOriginUpdateResponseUnion.AsAny]
type anyAccountOriginUpdateResponse interface {
	implAccountOriginUpdateResponseUnion()
}

func (AccountOriginUpdateResponseS3) implAccountOriginUpdateResponseUnion()               {}
func (AccountOriginUpdateResponseS3Compatible) implAccountOriginUpdateResponseUnion()     {}
func (AccountOriginUpdateResponseCloudinaryBackup) implAccountOriginUpdateResponseUnion() {}
func (AccountOriginUpdateResponseWebFolder) implAccountOriginUpdateResponseUnion()        {}
func (AccountOriginUpdateResponseWebProxy) implAccountOriginUpdateResponseUnion()         {}
func (AccountOriginUpdateResponseGcs) implAccountOriginUpdateResponseUnion()              {}
func (AccountOriginUpdateResponseAzureBlob) implAccountOriginUpdateResponseUnion()        {}
func (AccountOriginUpdateResponseAkeneoPim) implAccountOriginUpdateResponseUnion()        {}

// Use the following switch statement to find the correct variant
//
//	switch variant := AccountOriginUpdateResponseUnion.AsAny().(type) {
//	case imagekit.AccountOriginUpdateResponseS3:
//	case imagekit.AccountOriginUpdateResponseS3Compatible:
//	case imagekit.AccountOriginUpdateResponseCloudinaryBackup:
//	case imagekit.AccountOriginUpdateResponseWebFolder:
//	case imagekit.AccountOriginUpdateResponseWebProxy:
//	case imagekit.AccountOriginUpdateResponseGcs:
//	case imagekit.AccountOriginUpdateResponseAzureBlob:
//	case imagekit.AccountOriginUpdateResponseAkeneoPim:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u AccountOriginUpdateResponseUnion) AsAny() anyAccountOriginUpdateResponse {
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

func (u AccountOriginUpdateResponseUnion) AsGcs() (v AccountOriginUpdateResponseGcs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginUpdateResponseUnion) AsAzureBlob() (v AccountOriginUpdateResponseAzureBlob) {
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

type AccountOriginUpdateResponseGcs struct {
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
func (r AccountOriginUpdateResponseGcs) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginUpdateResponseGcs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginUpdateResponseAzureBlob struct {
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
func (r AccountOriginUpdateResponseAzureBlob) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginUpdateResponseAzureBlob) UnmarshalJSON(data []byte) error {
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
// [AccountOriginListResponseGcs], [AccountOriginListResponseAzureBlob],
// [AccountOriginListResponseAkeneoPim].
//
// Use the [AccountOriginListResponseUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AccountOriginListResponseUnion struct {
	ID                     string `json:"id"`
	Bucket                 string `json:"bucket"`
	IncludeCanonicalHeader bool   `json:"includeCanonicalHeader"`
	Name                   string `json:"name"`
	Prefix                 string `json:"prefix"`
	// Any of "S3", "S3_COMPATIBLE", "CLOUDINARY_BACKUP", "WEB_FOLDER", "WEB_PROXY",
	// "GCS", "AZURE_BLOB", "AKENEO_PIM".
	Type                      string `json:"type"`
	BaseURLForCanonicalHeader string `json:"baseUrlForCanonicalHeader"`
	// This field is from variant [AccountOriginListResponseS3Compatible].
	Endpoint string `json:"endpoint"`
	// This field is from variant [AccountOriginListResponseS3Compatible].
	S3ForcePathStyle bool   `json:"s3ForcePathStyle"`
	BaseURL          string `json:"baseUrl"`
	// This field is from variant [AccountOriginListResponseWebFolder].
	ForwardHostHeaderToOrigin bool `json:"forwardHostHeaderToOrigin"`
	// This field is from variant [AccountOriginListResponseGcs].
	ClientEmail string `json:"clientEmail"`
	// This field is from variant [AccountOriginListResponseAzureBlob].
	AccountName string `json:"accountName"`
	// This field is from variant [AccountOriginListResponseAzureBlob].
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

// anyAccountOriginListResponse is implemented by each variant of
// [AccountOriginListResponseUnion] to add type safety for the return type of
// [AccountOriginListResponseUnion.AsAny]
type anyAccountOriginListResponse interface {
	implAccountOriginListResponseUnion()
}

func (AccountOriginListResponseS3) implAccountOriginListResponseUnion()               {}
func (AccountOriginListResponseS3Compatible) implAccountOriginListResponseUnion()     {}
func (AccountOriginListResponseCloudinaryBackup) implAccountOriginListResponseUnion() {}
func (AccountOriginListResponseWebFolder) implAccountOriginListResponseUnion()        {}
func (AccountOriginListResponseWebProxy) implAccountOriginListResponseUnion()         {}
func (AccountOriginListResponseGcs) implAccountOriginListResponseUnion()              {}
func (AccountOriginListResponseAzureBlob) implAccountOriginListResponseUnion()        {}
func (AccountOriginListResponseAkeneoPim) implAccountOriginListResponseUnion()        {}

// Use the following switch statement to find the correct variant
//
//	switch variant := AccountOriginListResponseUnion.AsAny().(type) {
//	case imagekit.AccountOriginListResponseS3:
//	case imagekit.AccountOriginListResponseS3Compatible:
//	case imagekit.AccountOriginListResponseCloudinaryBackup:
//	case imagekit.AccountOriginListResponseWebFolder:
//	case imagekit.AccountOriginListResponseWebProxy:
//	case imagekit.AccountOriginListResponseGcs:
//	case imagekit.AccountOriginListResponseAzureBlob:
//	case imagekit.AccountOriginListResponseAkeneoPim:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u AccountOriginListResponseUnion) AsAny() anyAccountOriginListResponse {
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

func (u AccountOriginListResponseUnion) AsGcs() (v AccountOriginListResponseGcs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginListResponseUnion) AsAzureBlob() (v AccountOriginListResponseAzureBlob) {
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

type AccountOriginListResponseGcs struct {
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
func (r AccountOriginListResponseGcs) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginListResponseGcs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginListResponseAzureBlob struct {
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
func (r AccountOriginListResponseAzureBlob) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginListResponseAzureBlob) UnmarshalJSON(data []byte) error {
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
// [AccountOriginGetResponseWebProxy], [AccountOriginGetResponseGcs],
// [AccountOriginGetResponseAzureBlob], [AccountOriginGetResponseAkeneoPim].
//
// Use the [AccountOriginGetResponseUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AccountOriginGetResponseUnion struct {
	ID                     string `json:"id"`
	Bucket                 string `json:"bucket"`
	IncludeCanonicalHeader bool   `json:"includeCanonicalHeader"`
	Name                   string `json:"name"`
	Prefix                 string `json:"prefix"`
	// Any of "S3", "S3_COMPATIBLE", "CLOUDINARY_BACKUP", "WEB_FOLDER", "WEB_PROXY",
	// "GCS", "AZURE_BLOB", "AKENEO_PIM".
	Type                      string `json:"type"`
	BaseURLForCanonicalHeader string `json:"baseUrlForCanonicalHeader"`
	// This field is from variant [AccountOriginGetResponseS3Compatible].
	Endpoint string `json:"endpoint"`
	// This field is from variant [AccountOriginGetResponseS3Compatible].
	S3ForcePathStyle bool   `json:"s3ForcePathStyle"`
	BaseURL          string `json:"baseUrl"`
	// This field is from variant [AccountOriginGetResponseWebFolder].
	ForwardHostHeaderToOrigin bool `json:"forwardHostHeaderToOrigin"`
	// This field is from variant [AccountOriginGetResponseGcs].
	ClientEmail string `json:"clientEmail"`
	// This field is from variant [AccountOriginGetResponseAzureBlob].
	AccountName string `json:"accountName"`
	// This field is from variant [AccountOriginGetResponseAzureBlob].
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

// anyAccountOriginGetResponse is implemented by each variant of
// [AccountOriginGetResponseUnion] to add type safety for the return type of
// [AccountOriginGetResponseUnion.AsAny]
type anyAccountOriginGetResponse interface {
	implAccountOriginGetResponseUnion()
}

func (AccountOriginGetResponseS3) implAccountOriginGetResponseUnion()               {}
func (AccountOriginGetResponseS3Compatible) implAccountOriginGetResponseUnion()     {}
func (AccountOriginGetResponseCloudinaryBackup) implAccountOriginGetResponseUnion() {}
func (AccountOriginGetResponseWebFolder) implAccountOriginGetResponseUnion()        {}
func (AccountOriginGetResponseWebProxy) implAccountOriginGetResponseUnion()         {}
func (AccountOriginGetResponseGcs) implAccountOriginGetResponseUnion()              {}
func (AccountOriginGetResponseAzureBlob) implAccountOriginGetResponseUnion()        {}
func (AccountOriginGetResponseAkeneoPim) implAccountOriginGetResponseUnion()        {}

// Use the following switch statement to find the correct variant
//
//	switch variant := AccountOriginGetResponseUnion.AsAny().(type) {
//	case imagekit.AccountOriginGetResponseS3:
//	case imagekit.AccountOriginGetResponseS3Compatible:
//	case imagekit.AccountOriginGetResponseCloudinaryBackup:
//	case imagekit.AccountOriginGetResponseWebFolder:
//	case imagekit.AccountOriginGetResponseWebProxy:
//	case imagekit.AccountOriginGetResponseGcs:
//	case imagekit.AccountOriginGetResponseAzureBlob:
//	case imagekit.AccountOriginGetResponseAkeneoPim:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u AccountOriginGetResponseUnion) AsAny() anyAccountOriginGetResponse {
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

func (u AccountOriginGetResponseUnion) AsGcs() (v AccountOriginGetResponseGcs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccountOriginGetResponseUnion) AsAzureBlob() (v AccountOriginGetResponseAzureBlob) {
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

type AccountOriginGetResponseGcs struct {
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
func (r AccountOriginGetResponseGcs) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginGetResponseGcs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountOriginGetResponseAzureBlob struct {
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
func (r AccountOriginGetResponseAzureBlob) RawJSON() string { return r.JSON.raw }
func (r *AccountOriginGetResponseAzureBlob) UnmarshalJSON(data []byte) error {
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
