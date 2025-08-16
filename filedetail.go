// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/imagekit-go/internal/apijson"
	"github.com/stainless-sdks/imagekit-go/internal/requestconfig"
	"github.com/stainless-sdks/imagekit-go/option"
	"github.com/stainless-sdks/imagekit-go/packages/param"
	"github.com/stainless-sdks/imagekit-go/packages/respjson"
)

// FileDetailService contains methods and other services that help with interacting
// with the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileDetailService] method instead.
type FileDetailService struct {
	Options []option.RequestOption
}

// NewFileDetailService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFileDetailService(opts ...option.RequestOption) (r FileDetailService) {
	r = FileDetailService{}
	r.Options = opts
	return
}

// This API returns an object with details or attributes about the current version
// of the file.
func (r *FileDetailService) Get(ctx context.Context, fileID string, opts ...option.RequestOption) (res *FileDetailGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return
	}
	path := fmt.Sprintf("v1/files/%s/details", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This API updates the details or attributes of the current version of the file.
// You can update `tags`, `customCoordinates`, `customMetadata`, publication
// status, remove existing `AITags` and apply extensions using this API.
func (r *FileDetailService) Update(ctx context.Context, fileID string, body FileDetailUpdateParams, opts ...option.RequestOption) (res *FileDetailUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return
	}
	path := fmt.Sprintf("v1/files/%s/details", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Object containing details of a file or file version.
type FileDetailGetResponse struct {
	// An array of tags assigned to the file by auto tagging.
	AITags []FileDetailGetResponseAITag `json:"AITags,nullable"`
	// Date and time when the file was uploaded. The date and time is in ISO8601
	// format.
	CreatedAt string `json:"createdAt"`
	// An string with custom coordinates of the file.
	CustomCoordinates string `json:"customCoordinates,nullable"`
	// An object with custom metadata for the file.
	CustomMetadata any `json:"customMetadata"`
	// Unique identifier of the asset.
	FileID string `json:"fileId"`
	// Path of the file. This is the path you would use in the URL to access the file.
	// For example, if the file is at the root of the media library, the path will be
	// `/file.jpg`. If the file is inside a folder named `images`, the path will be
	// `/images/file.jpg`.
	FilePath string `json:"filePath"`
	// Type of the file. Possible values are `image`, `non-image`.
	FileType string `json:"fileType"`
	// Specifies if the image has an alpha channel.
	HasAlpha bool `json:"hasAlpha"`
	// Height of the file.
	Height float64 `json:"height"`
	// Specifies if the file is private or not.
	IsPrivateFile bool `json:"isPrivateFile"`
	// Specifies if the file is published or not.
	IsPublished bool `json:"isPublished"`
	// MIME type of the file.
	Mime string `json:"mime"`
	// Name of the asset.
	Name string `json:"name"`
	// Size of the file in bytes.
	Size float64 `json:"size"`
	// An array of tags assigned to the file. Tags are used to search files in the
	// media library.
	Tags []string `json:"tags,nullable"`
	// URL of the thumbnail image. This URL is used to access the thumbnail image of
	// the file in the media library.
	Thumbnail string `json:"thumbnail"`
	// Type of the asset.
	Type string `json:"type"`
	// Date and time when the file was last updated. The date and time is in ISO8601
	// format.
	UpdatedAt string `json:"updatedAt"`
	// URL of the file.
	URL string `json:"url"`
	// An object with details of the file version.
	VersionInfo FileDetailGetResponseVersionInfo `json:"versionInfo"`
	// Width of the file.
	Width float64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AITags            respjson.Field
		CreatedAt         respjson.Field
		CustomCoordinates respjson.Field
		CustomMetadata    respjson.Field
		FileID            respjson.Field
		FilePath          respjson.Field
		FileType          respjson.Field
		HasAlpha          respjson.Field
		Height            respjson.Field
		IsPrivateFile     respjson.Field
		IsPublished       respjson.Field
		Mime              respjson.Field
		Name              respjson.Field
		Size              respjson.Field
		Tags              respjson.Field
		Thumbnail         respjson.Field
		Type              respjson.Field
		UpdatedAt         respjson.Field
		URL               respjson.Field
		VersionInfo       respjson.Field
		Width             respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileDetailGetResponse) RawJSON() string { return r.JSON.raw }
func (r *FileDetailGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileDetailGetResponseAITag struct {
	// Confidence score of the tag.
	Confidence float64 `json:"confidence"`
	// Name of the tag.
	Name string `json:"name"`
	// Source of the tag. Possible values are `google-auto-tagging` and
	// `aws-auto-tagging`.
	Source string `json:"source"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Confidence  respjson.Field
		Name        respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileDetailGetResponseAITag) RawJSON() string { return r.JSON.raw }
func (r *FileDetailGetResponseAITag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object with details of the file version.
type FileDetailGetResponseVersionInfo struct {
	// Unique identifier of the file version.
	ID string `json:"id"`
	// Name of the file version.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileDetailGetResponseVersionInfo) RawJSON() string { return r.JSON.raw }
func (r *FileDetailGetResponseVersionInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileDetailUpdateResponse struct {
	// An array of tags assigned to the file by auto tagging.
	AITags []FileDetailUpdateResponseAITag `json:"AITags,nullable"`
	// Date and time when the file was uploaded. The date and time is in ISO8601
	// format.
	CreatedAt string `json:"createdAt"`
	// An string with custom coordinates of the file.
	CustomCoordinates string `json:"customCoordinates,nullable"`
	// An object with custom metadata for the file.
	CustomMetadata  any                                     `json:"customMetadata"`
	ExtensionStatus FileDetailUpdateResponseExtensionStatus `json:"extensionStatus"`
	// Unique identifier of the asset.
	FileID string `json:"fileId"`
	// Path of the file. This is the path you would use in the URL to access the file.
	// For example, if the file is at the root of the media library, the path will be
	// `/file.jpg`. If the file is inside a folder named `images`, the path will be
	// `/images/file.jpg`.
	FilePath string `json:"filePath"`
	// Type of the file. Possible values are `image`, `non-image`.
	FileType string `json:"fileType"`
	// Specifies if the image has an alpha channel.
	HasAlpha bool `json:"hasAlpha"`
	// Height of the file.
	Height float64 `json:"height"`
	// Specifies if the file is private or not.
	IsPrivateFile bool `json:"isPrivateFile"`
	// Specifies if the file is published or not.
	IsPublished bool `json:"isPublished"`
	// MIME type of the file.
	Mime string `json:"mime"`
	// Name of the asset.
	Name string `json:"name"`
	// Size of the file in bytes.
	Size float64 `json:"size"`
	// An array of tags assigned to the file. Tags are used to search files in the
	// media library.
	Tags []string `json:"tags,nullable"`
	// URL of the thumbnail image. This URL is used to access the thumbnail image of
	// the file in the media library.
	Thumbnail string `json:"thumbnail"`
	// Type of the asset.
	Type string `json:"type"`
	// Date and time when the file was last updated. The date and time is in ISO8601
	// format.
	UpdatedAt string `json:"updatedAt"`
	// URL of the file.
	URL string `json:"url"`
	// An object with details of the file version.
	VersionInfo FileDetailUpdateResponseVersionInfo `json:"versionInfo"`
	// Width of the file.
	Width float64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AITags            respjson.Field
		CreatedAt         respjson.Field
		CustomCoordinates respjson.Field
		CustomMetadata    respjson.Field
		ExtensionStatus   respjson.Field
		FileID            respjson.Field
		FilePath          respjson.Field
		FileType          respjson.Field
		HasAlpha          respjson.Field
		Height            respjson.Field
		IsPrivateFile     respjson.Field
		IsPublished       respjson.Field
		Mime              respjson.Field
		Name              respjson.Field
		Size              respjson.Field
		Tags              respjson.Field
		Thumbnail         respjson.Field
		Type              respjson.Field
		UpdatedAt         respjson.Field
		URL               respjson.Field
		VersionInfo       respjson.Field
		Width             respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileDetailUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *FileDetailUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileDetailUpdateResponseAITag struct {
	// Confidence score of the tag.
	Confidence float64 `json:"confidence"`
	// Name of the tag.
	Name string `json:"name"`
	// Source of the tag. Possible values are `google-auto-tagging` and
	// `aws-auto-tagging`.
	Source string `json:"source"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Confidence  respjson.Field
		Name        respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileDetailUpdateResponseAITag) RawJSON() string { return r.JSON.raw }
func (r *FileDetailUpdateResponseAITag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileDetailUpdateResponseExtensionStatus struct {
	// Any of "success", "pending", "failed".
	AwsAutoTagging string `json:"aws-auto-tagging"`
	// Any of "success", "pending", "failed".
	GoogleAutoTagging string `json:"google-auto-tagging"`
	// Any of "success", "pending", "failed".
	RemoveBg string `json:"remove-bg"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AwsAutoTagging    respjson.Field
		GoogleAutoTagging respjson.Field
		RemoveBg          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileDetailUpdateResponseExtensionStatus) RawJSON() string { return r.JSON.raw }
func (r *FileDetailUpdateResponseExtensionStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object with details of the file version.
type FileDetailUpdateResponseVersionInfo struct {
	// Unique identifier of the file version.
	ID string `json:"id"`
	// Name of the file version.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileDetailUpdateResponseVersionInfo) RawJSON() string { return r.JSON.raw }
func (r *FileDetailUpdateResponseVersionInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileDetailUpdateParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfUpdateFileDetails *FileDetailUpdateParamsBodyUpdateFileDetails `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfChangePublicationStatus *FileDetailUpdateParamsBodyChangePublicationStatus `json:",inline"`

	paramObj
}

func (u FileDetailUpdateParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfUpdateFileDetails, u.OfChangePublicationStatus)
}
func (r *FileDetailUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileDetailUpdateParamsBodyUpdateFileDetails struct {
	// Define an important area in the image in the format `x,y,width,height` e.g.
	// `10,10,100,100`. Send `null` to unset this value.
	CustomCoordinates param.Opt[string] `json:"customCoordinates,omitzero"`
	// The final status of extensions after they have completed execution will be
	// delivered to this endpoint as a POST request.
	// [Learn more](/docs/api-reference/digital-asset-management-dam/managing-assets/update-file-details#webhook-payload-structure)
	// about the webhook payload structure.
	WebhookURL param.Opt[string] `json:"webhookUrl,omitzero"`
	// An array of tags associated with the file, such as `["tag1", "tag2"]`. Send
	// `null` to unset all tags associated with the file.
	Tags []string `json:"tags,omitzero"`
	// A key-value data to be associated with the asset. To unset a key, send `null`
	// value for that key. Before setting any custom metadata on an asset you have to
	// create the field using custom metadata fields API.
	CustomMetadata any `json:"customMetadata,omitzero"`
	// Array of extensions to be applied to the asset. Each extension can be configured
	// with specific parameters based on the extension type.
	Extensions []FileDetailUpdateParamsBodyUpdateFileDetailsExtensionUnion `json:"extensions,omitzero"`
	// An array of AITags associated with the file that you want to remove, e.g.
	// `["car", "vehicle", "motorsports"]`.
	//
	// If you want to remove all AITags associated with the file, send a string -
	// "all".
	//
	// Note: The remove operation for `AITags` executes before any of the `extensions`
	// are processed.
	RemoveAITags FileDetailUpdateParamsBodyUpdateFileDetailsRemoveAITagsUnion `json:"removeAITags,omitzero"`
	paramObj
}

func (r FileDetailUpdateParamsBodyUpdateFileDetails) MarshalJSON() (data []byte, err error) {
	type shadow FileDetailUpdateParamsBodyUpdateFileDetails
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileDetailUpdateParamsBodyUpdateFileDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FileDetailUpdateParamsBodyUpdateFileDetailsExtensionUnion struct {
	OfRemoveBackground *FileDetailUpdateParamsBodyUpdateFileDetailsExtensionRemoveBackground `json:",omitzero,inline"`
	OfAutoTagging      *FileDetailUpdateParamsBodyUpdateFileDetailsExtensionAutoTagging      `json:",omitzero,inline"`
	paramUnion
}

func (u FileDetailUpdateParamsBodyUpdateFileDetailsExtensionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRemoveBackground, u.OfAutoTagging)
}
func (u *FileDetailUpdateParamsBodyUpdateFileDetailsExtensionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FileDetailUpdateParamsBodyUpdateFileDetailsExtensionUnion) asAny() any {
	if !param.IsOmitted(u.OfRemoveBackground) {
		return u.OfRemoveBackground
	} else if !param.IsOmitted(u.OfAutoTagging) {
		return u.OfAutoTagging
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FileDetailUpdateParamsBodyUpdateFileDetailsExtensionUnion) GetOptions() *FileDetailUpdateParamsBodyUpdateFileDetailsExtensionRemoveBackgroundOptions {
	if vt := u.OfRemoveBackground; vt != nil {
		return &vt.Options
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FileDetailUpdateParamsBodyUpdateFileDetailsExtensionUnion) GetMaxTags() *int64 {
	if vt := u.OfAutoTagging; vt != nil {
		return &vt.MaxTags
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FileDetailUpdateParamsBodyUpdateFileDetailsExtensionUnion) GetMinConfidence() *int64 {
	if vt := u.OfAutoTagging; vt != nil {
		return &vt.MinConfidence
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FileDetailUpdateParamsBodyUpdateFileDetailsExtensionUnion) GetName() *string {
	if vt := u.OfRemoveBackground; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfAutoTagging; vt != nil {
		return (*string)(&vt.Name)
	}
	return nil
}

// The property Name is required.
type FileDetailUpdateParamsBodyUpdateFileDetailsExtensionRemoveBackground struct {
	// Specifies the background removal extension.
	//
	// Any of "remove-bg".
	Name    string                                                                      `json:"name,omitzero,required"`
	Options FileDetailUpdateParamsBodyUpdateFileDetailsExtensionRemoveBackgroundOptions `json:"options,omitzero"`
	paramObj
}

func (r FileDetailUpdateParamsBodyUpdateFileDetailsExtensionRemoveBackground) MarshalJSON() (data []byte, err error) {
	type shadow FileDetailUpdateParamsBodyUpdateFileDetailsExtensionRemoveBackground
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileDetailUpdateParamsBodyUpdateFileDetailsExtensionRemoveBackground) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FileDetailUpdateParamsBodyUpdateFileDetailsExtensionRemoveBackground](
		"name", "remove-bg",
	)
}

type FileDetailUpdateParamsBodyUpdateFileDetailsExtensionRemoveBackgroundOptions struct {
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

func (r FileDetailUpdateParamsBodyUpdateFileDetailsExtensionRemoveBackgroundOptions) MarshalJSON() (data []byte, err error) {
	type shadow FileDetailUpdateParamsBodyUpdateFileDetailsExtensionRemoveBackgroundOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileDetailUpdateParamsBodyUpdateFileDetailsExtensionRemoveBackgroundOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MaxTags, MinConfidence, Name are required.
type FileDetailUpdateParamsBodyUpdateFileDetailsExtensionAutoTagging struct {
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

func (r FileDetailUpdateParamsBodyUpdateFileDetailsExtensionAutoTagging) MarshalJSON() (data []byte, err error) {
	type shadow FileDetailUpdateParamsBodyUpdateFileDetailsExtensionAutoTagging
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileDetailUpdateParamsBodyUpdateFileDetailsExtensionAutoTagging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FileDetailUpdateParamsBodyUpdateFileDetailsExtensionAutoTagging](
		"name", "google-auto-tagging", "aws-auto-tagging",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FileDetailUpdateParamsBodyUpdateFileDetailsRemoveAITagsUnion struct {
	OfStringArray []string `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfFileDetailUpdatesBodyUpdateFileDetailsRemoveAITagsString)
	OfFileDetailUpdatesBodyUpdateFileDetailsRemoveAITagsString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u FileDetailUpdateParamsBodyUpdateFileDetailsRemoveAITagsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfStringArray, u.OfFileDetailUpdatesBodyUpdateFileDetailsRemoveAITagsString)
}
func (u *FileDetailUpdateParamsBodyUpdateFileDetailsRemoveAITagsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FileDetailUpdateParamsBodyUpdateFileDetailsRemoveAITagsUnion) asAny() any {
	if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	} else if !param.IsOmitted(u.OfFileDetailUpdatesBodyUpdateFileDetailsRemoveAITagsString) {
		return &u.OfFileDetailUpdatesBodyUpdateFileDetailsRemoveAITagsString
	}
	return nil
}

type FileDetailUpdateParamsBodyUpdateFileDetailsRemoveAITagsString string

const (
	FileDetailUpdateParamsBodyUpdateFileDetailsRemoveAITagsStringAll FileDetailUpdateParamsBodyUpdateFileDetailsRemoveAITagsString = "all"
)

type FileDetailUpdateParamsBodyChangePublicationStatus struct {
	// Configure the publication status of a file and its versions.
	Publish FileDetailUpdateParamsBodyChangePublicationStatusPublish `json:"publish,omitzero"`
	paramObj
}

func (r FileDetailUpdateParamsBodyChangePublicationStatus) MarshalJSON() (data []byte, err error) {
	type shadow FileDetailUpdateParamsBodyChangePublicationStatus
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileDetailUpdateParamsBodyChangePublicationStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configure the publication status of a file and its versions.
//
// The property IsPublished is required.
type FileDetailUpdateParamsBodyChangePublicationStatusPublish struct {
	// Set to `true` to publish the file. Set to `false` to unpublish the file.
	IsPublished bool `json:"isPublished,required"`
	// Set to `true` to publish/unpublish all versions of the file. Set to `false` to
	// publish/unpublish only the current version of the file.
	IncludeFileVersions param.Opt[bool] `json:"includeFileVersions,omitzero"`
	paramObj
}

func (r FileDetailUpdateParamsBodyChangePublicationStatusPublish) MarshalJSON() (data []byte, err error) {
	type shadow FileDetailUpdateParamsBodyChangePublicationStatusPublish
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileDetailUpdateParamsBodyChangePublicationStatusPublish) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
