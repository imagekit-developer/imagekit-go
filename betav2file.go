// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/imagekit-developer/imagekit-go/internal/apiform"
	"github.com/imagekit-developer/imagekit-go/internal/apijson"
	"github.com/imagekit-developer/imagekit-go/internal/requestconfig"
	"github.com/imagekit-developer/imagekit-go/option"
	"github.com/imagekit-developer/imagekit-go/packages/param"
	"github.com/imagekit-developer/imagekit-go/packages/respjson"
	"github.com/imagekit-developer/imagekit-go/shared/constant"
)

// BetaV2FileService contains methods and other services that help with interacting
// with the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaV2FileService] method instead.
type BetaV2FileService struct {
	Options []option.RequestOption
}

// NewBetaV2FileService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBetaV2FileService(opts ...option.RequestOption) (r BetaV2FileService) {
	r = BetaV2FileService{}
	r.Options = opts
	return
}

// The V2 API enhances security by verifying the entire payload using JWT. This API
// is in beta.
//
// ImageKit.io allows you to upload files directly from both the server and client
// sides. For server-side uploads, private API key authentication is used. For
// client-side uploads, generate a one-time `token` from your secure backend using
// private API.
// [Learn more](/docs/api-reference/upload-file/upload-file-v2#how-to-implement-secure-client-side-file-upload)
// about how to implement secure client-side file upload.
//
// **File size limit** \
// On the free plan, the maximum upload file sizes are 20MB for images, audio, and raw
// files, and 100MB for videos. On the paid plan, these limits increase to 40MB for
// images, audio, and raw files, and 2GB for videos. These limits can be further increased
// with higher-tier plans.
//
// **Version limit** \
// A file can have a maximum of 100 versions.
//
// **Demo applications**
//
//   - A full-fledged
//     [upload widget using Uppy](https://github.com/imagekit-samples/uppy-uploader),
//     supporting file selections from local storage, URL, Dropbox, Google Drive,
//     Instagram, and more.
//   - [Quick start guides](/docs/quick-start-guides) for various frameworks and
//     technologies.
func (r *BetaV2FileService) Upload(ctx context.Context, body BetaV2FileUploadParams, opts ...option.RequestOption) (res *BetaV2FileUploadResponse, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://upload.imagekit.io/")}, opts...)
	path := "api/v2/files/upload"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Object containing details of a successful upload.
type BetaV2FileUploadResponse struct {
	// An array of tags assigned to the uploaded file by auto tagging.
	AITags []BetaV2FileUploadResponseAITag `json:"AITags,nullable"`
	// The audio codec used in the video (only for video).
	AudioCodec string `json:"audioCodec"`
	// The bit rate of the video in kbps (only for video).
	BitRate int64 `json:"bitRate"`
	// Value of custom coordinates associated with the image in the format
	// `x,y,width,height`. If `customCoordinates` are not defined, then it is `null`.
	// Send `customCoordinates` in `responseFields` in API request to get the value of
	// this field.
	CustomCoordinates string `json:"customCoordinates,nullable"`
	// A key-value data associated with the asset. Use `responseField` in API request
	// to get `customMetadata` in the upload API response. Before setting any custom
	// metadata on an asset, you have to create the field using custom metadata fields
	// API. Send `customMetadata` in `responseFields` in API request to get the value
	// of this field.
	CustomMetadata map[string]any `json:"customMetadata"`
	// Optional text to describe the contents of the file. Can be set by the user or
	// the ai-auto-description extension.
	Description string `json:"description"`
	// The duration of the video in seconds (only for video).
	Duration int64 `json:"duration"`
	// Consolidated embedded metadata associated with the file. It includes exif, iptc,
	// and xmp data. Send `embeddedMetadata` in `responseFields` in API request to get
	// embeddedMetadata in the upload API response.
	EmbeddedMetadata map[string]any `json:"embeddedMetadata"`
	// Extension names with their processing status at the time of completion of the
	// request. It could have one of the following status values:
	//
	// `success`: The extension has been successfully applied. `failed`: The extension
	// has failed and will not be retried. `pending`: The extension will finish
	// processing in some time. On completion, the final status (success / failed) will
	// be sent to the `webhookUrl` provided.
	//
	// If no extension was requested, then this parameter is not returned.
	ExtensionStatus BetaV2FileUploadResponseExtensionStatus `json:"extensionStatus"`
	// Unique fileId. Store this fileld in your database, as this will be used to
	// perform update action on this file.
	FileID string `json:"fileId"`
	// The relative path of the file in the media library e.g.
	// `/marketing-assets/new-banner.jpg`.
	FilePath string `json:"filePath"`
	// Type of the uploaded file. Possible values are `image`, `non-image`.
	FileType string `json:"fileType"`
	// Height of the image in pixels (Only for images)
	Height float64 `json:"height"`
	// Is the file marked as private. It can be either `true` or `false`. Send
	// `isPrivateFile` in `responseFields` in API request to get the value of this
	// field.
	IsPrivateFile bool `json:"isPrivateFile"`
	// Is the file published or in draft state. It can be either `true` or `false`.
	// Send `isPublished` in `responseFields` in API request to get the value of this
	// field.
	IsPublished bool `json:"isPublished"`
	// Legacy metadata. Send `metadata` in `responseFields` in API request to get
	// metadata in the upload API response.
	Metadata Metadata `json:"metadata"`
	// Name of the asset.
	Name string `json:"name"`
	// Size of the image file in Bytes.
	Size float64 `json:"size"`
	// The array of tags associated with the asset. If no tags are set, it will be
	// `null`. Send `tags` in `responseFields` in API request to get the value of this
	// field.
	Tags []string `json:"tags,nullable"`
	// In the case of an image, a small thumbnail URL.
	ThumbnailURL string `json:"thumbnailUrl"`
	// A publicly accessible URL of the file.
	URL string `json:"url"`
	// An object containing the file or file version's `id` (versionId) and `name`.
	VersionInfo BetaV2FileUploadResponseVersionInfo `json:"versionInfo"`
	// The video codec used in the video (only for video).
	VideoCodec string `json:"videoCodec"`
	// Width of the image in pixels (Only for Images)
	Width float64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AITags            respjson.Field
		AudioCodec        respjson.Field
		BitRate           respjson.Field
		CustomCoordinates respjson.Field
		CustomMetadata    respjson.Field
		Description       respjson.Field
		Duration          respjson.Field
		EmbeddedMetadata  respjson.Field
		ExtensionStatus   respjson.Field
		FileID            respjson.Field
		FilePath          respjson.Field
		FileType          respjson.Field
		Height            respjson.Field
		IsPrivateFile     respjson.Field
		IsPublished       respjson.Field
		Metadata          respjson.Field
		Name              respjson.Field
		Size              respjson.Field
		Tags              respjson.Field
		ThumbnailURL      respjson.Field
		URL               respjson.Field
		VersionInfo       respjson.Field
		VideoCodec        respjson.Field
		Width             respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaV2FileUploadResponse) RawJSON() string { return r.JSON.raw }
func (r *BetaV2FileUploadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaV2FileUploadResponseAITag struct {
	// Confidence score of the tag.
	Confidence float64 `json:"confidence"`
	// Name of the tag.
	Name string `json:"name"`
	// Array of `AITags` associated with the image. If no `AITags` are set, it will be
	// null. These tags can be added using the `google-auto-tagging` or
	// `aws-auto-tagging` extensions.
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
func (r BetaV2FileUploadResponseAITag) RawJSON() string { return r.JSON.raw }
func (r *BetaV2FileUploadResponseAITag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Extension names with their processing status at the time of completion of the
// request. It could have one of the following status values:
//
// `success`: The extension has been successfully applied. `failed`: The extension
// has failed and will not be retried. `pending`: The extension will finish
// processing in some time. On completion, the final status (success / failed) will
// be sent to the `webhookUrl` provided.
//
// If no extension was requested, then this parameter is not returned.
type BetaV2FileUploadResponseExtensionStatus struct {
	// Any of "success", "pending", "failed".
	AIAutoDescription string `json:"ai-auto-description"`
	// Any of "success", "pending", "failed".
	AwsAutoTagging string `json:"aws-auto-tagging"`
	// Any of "success", "pending", "failed".
	GoogleAutoTagging string `json:"google-auto-tagging"`
	// Any of "success", "pending", "failed".
	RemoveBg string `json:"remove-bg"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AIAutoDescription respjson.Field
		AwsAutoTagging    respjson.Field
		GoogleAutoTagging respjson.Field
		RemoveBg          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BetaV2FileUploadResponseExtensionStatus) RawJSON() string { return r.JSON.raw }
func (r *BetaV2FileUploadResponseExtensionStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object containing the file or file version's `id` (versionId) and `name`.
type BetaV2FileUploadResponseVersionInfo struct {
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
func (r BetaV2FileUploadResponseVersionInfo) RawJSON() string { return r.JSON.raw }
func (r *BetaV2FileUploadResponseVersionInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaV2FileUploadParams struct {
	// The API accepts any of the following:
	//
	//   - **Binary data** – send the raw bytes as `multipart/form-data`.
	//   - **HTTP / HTTPS URL** – a publicly reachable URL that ImageKit’s servers can
	//     fetch.
	//   - **Base64 string** – the file encoded as a Base64 data URI or plain Base64.
	//
	// When supplying a URL, the server must receive the response headers within 8
	// seconds; otherwise the request fails with 400 Bad Request.
	File io.Reader `json:"file,omitzero,required" format:"binary"`
	// The name with which the file has to be uploaded.
	FileName string `json:"fileName,required"`
	// This is the client-generated JSON Web Token (JWT). The ImageKit.io server uses
	// it to authenticate and check that the upload request parameters have not been
	// tampered with after the token has been generated. Learn how to create the token
	// on the page below. This field is only required for authentication when uploading
	// a file from the client side.
	//
	// **Note**: Sending a JWT that has been used in the past will result in a
	// validation error. Even if your previous request resulted in an error, you should
	// always send a new token.
	//
	// **⚠️Warning**: JWT must be generated on the server-side because it is generated
	// using your account's private API key. This field is required for authentication
	// when uploading a file from the client-side.
	Token param.Opt[string] `json:"token,omitzero"`
	// Server-side checks to run on the asset. Read more about
	// [Upload API checks](/docs/api-reference/upload-file/upload-file-v2#upload-api-checks).
	Checks param.Opt[string] `json:"checks,omitzero"`
	// Define an important area in the image. This is only relevant for image type
	// files.
	//
	//   - To be passed as a string with the x and y coordinates of the top-left corner,
	//     and width and height of the area of interest in the format `x,y,width,height`.
	//     For example - `10,10,100,100`
	//   - Can be used with fo-customtransformation.
	//   - If this field is not specified and the file is overwritten, then
	//     customCoordinates will be removed.
	CustomCoordinates param.Opt[string] `json:"customCoordinates,omitzero"`
	// Optional text to describe the contents of the file.
	Description param.Opt[string] `json:"description,omitzero"`
	// The folder path in which the image has to be uploaded. If the folder(s) didn't
	// exist before, a new folder(s) is created. Using multiple `/` creates a nested
	// folder.
	Folder param.Opt[string] `json:"folder,omitzero"`
	// Whether to mark the file as private or not.
	//
	// If `true`, the file is marked as private and is accessible only using named
	// transformation or signed URL.
	IsPrivateFile param.Opt[bool] `json:"isPrivateFile,omitzero"`
	// Whether to upload file as published or not.
	//
	// If `false`, the file is marked as unpublished, which restricts access to the
	// file only via the media library. Files in draft or unpublished state can only be
	// publicly accessed after being published.
	//
	// The option to upload in draft state is only available in custom enterprise
	// pricing plans.
	IsPublished param.Opt[bool] `json:"isPublished,omitzero"`
	// If set to `true` and a file already exists at the exact location, its AITags
	// will be removed. Set `overwriteAITags` to `false` to preserve AITags.
	OverwriteAITags param.Opt[bool] `json:"overwriteAITags,omitzero"`
	// If the request does not have `customMetadata`, and a file already exists at the
	// exact location, existing customMetadata will be removed.
	OverwriteCustomMetadata param.Opt[bool] `json:"overwriteCustomMetadata,omitzero"`
	// If `false` and `useUniqueFileName` is also `false`, and a file already exists at
	// the exact location, upload API will return an error immediately.
	OverwriteFile param.Opt[bool] `json:"overwriteFile,omitzero"`
	// If the request does not have `tags`, and a file already exists at the exact
	// location, existing tags will be removed.
	OverwriteTags param.Opt[bool] `json:"overwriteTags,omitzero"`
	// Whether to use a unique filename for this file or not.
	//
	// If `true`, ImageKit.io will add a unique suffix to the filename parameter to get
	// a unique filename.
	//
	// If `false`, then the image is uploaded with the provided filename parameter, and
	// any existing file with the same name is replaced.
	UseUniqueFileName param.Opt[bool] `json:"useUniqueFileName,omitzero"`
	// The final status of extensions after they have completed execution will be
	// delivered to this endpoint as a POST request.
	// [Learn more](/docs/api-reference/digital-asset-management-dam/managing-assets/update-file-details#webhook-payload-structure)
	// about the webhook payload structure.
	WebhookURL param.Opt[string] `json:"webhookUrl,omitzero" format:"uri"`
	// JSON key-value pairs to associate with the asset. Create the custom metadata
	// fields before setting these values.
	CustomMetadata map[string]any `json:"customMetadata,omitzero"`
	// Array of extensions to be applied to the image. Each extension can be configured
	// with specific parameters based on the extension type.
	Extensions []BetaV2FileUploadParamsExtensionUnion `json:"extensions,omitzero"`
	// Array of response field keys to include in the API response body.
	//
	// Any of "tags", "customCoordinates", "isPrivateFile", "embeddedMetadata",
	// "isPublished", "customMetadata", "metadata".
	ResponseFields []string `json:"responseFields,omitzero"`
	// Set the tags while uploading the file. Provide an array of tag strings (e.g.
	// `["tag1", "tag2", "tag3"]`). The combined length of all tag characters must not
	// exceed 500, and the `%` character is not allowed. If this field is not specified
	// and the file is overwritten, the existing tags will be removed.
	Tags []string `json:"tags,omitzero"`
	// Configure pre-processing (`pre`) and post-processing (`post`) transformations.
	//
	//   - `pre` — applied before the file is uploaded to the Media Library.
	//     Useful for reducing file size or applying basic optimizations upfront (e.g.,
	//     resize, compress).
	//
	//   - `post` — applied immediately after upload.
	//     Ideal for generating transformed versions (like video encodes or thumbnails)
	//     in advance, so they're ready for delivery without delay.
	//
	// You can mix and match any combination of post-processing types.
	Transformation BetaV2FileUploadParamsTransformation `json:"transformation,omitzero"`
	paramObj
}

func (r BetaV2FileUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaV2FileUploadParamsExtensionUnion struct {
	OfRemoveBg          *BetaV2FileUploadParamsExtensionRemoveBg          `json:",omitzero,inline"`
	OfAutoTagging       *BetaV2FileUploadParamsExtensionAutoTagging       `json:",omitzero,inline"`
	OfAIAutoDescription *BetaV2FileUploadParamsExtensionAIAutoDescription `json:",omitzero,inline"`
	paramUnion
}

func (u BetaV2FileUploadParamsExtensionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRemoveBg, u.OfAutoTagging, u.OfAIAutoDescription)
}
func (u *BetaV2FileUploadParamsExtensionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BetaV2FileUploadParamsExtensionUnion) asAny() any {
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
func (u BetaV2FileUploadParamsExtensionUnion) GetOptions() *BetaV2FileUploadParamsExtensionRemoveBgOptions {
	if vt := u.OfRemoveBg; vt != nil {
		return &vt.Options
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BetaV2FileUploadParamsExtensionUnion) GetMaxTags() *int64 {
	if vt := u.OfAutoTagging; vt != nil {
		return &vt.MaxTags
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BetaV2FileUploadParamsExtensionUnion) GetMinConfidence() *int64 {
	if vt := u.OfAutoTagging; vt != nil {
		return &vt.MinConfidence
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BetaV2FileUploadParamsExtensionUnion) GetName() *string {
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
	apijson.RegisterUnion[BetaV2FileUploadParamsExtensionUnion](
		"name",
		apijson.Discriminator[BetaV2FileUploadParamsExtensionRemoveBg]("remove-bg"),
		apijson.Discriminator[BetaV2FileUploadParamsExtensionAutoTagging]("google-auto-tagging"),
		apijson.Discriminator[BetaV2FileUploadParamsExtensionAutoTagging]("aws-auto-tagging"),
		apijson.Discriminator[BetaV2FileUploadParamsExtensionAIAutoDescription]("ai-auto-description"),
	)
}

// The property Name is required.
type BetaV2FileUploadParamsExtensionRemoveBg struct {
	Options BetaV2FileUploadParamsExtensionRemoveBgOptions `json:"options,omitzero"`
	// Specifies the background removal extension.
	//
	// This field can be elided, and will marshal its zero value as "remove-bg".
	Name constant.RemoveBg `json:"name,required"`
	paramObj
}

func (r BetaV2FileUploadParamsExtensionRemoveBg) MarshalJSON() (data []byte, err error) {
	type shadow BetaV2FileUploadParamsExtensionRemoveBg
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaV2FileUploadParamsExtensionRemoveBg) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BetaV2FileUploadParamsExtensionRemoveBgOptions struct {
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

func (r BetaV2FileUploadParamsExtensionRemoveBgOptions) MarshalJSON() (data []byte, err error) {
	type shadow BetaV2FileUploadParamsExtensionRemoveBgOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaV2FileUploadParamsExtensionRemoveBgOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MaxTags, MinConfidence, Name are required.
type BetaV2FileUploadParamsExtensionAutoTagging struct {
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

func (r BetaV2FileUploadParamsExtensionAutoTagging) MarshalJSON() (data []byte, err error) {
	type shadow BetaV2FileUploadParamsExtensionAutoTagging
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaV2FileUploadParamsExtensionAutoTagging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BetaV2FileUploadParamsExtensionAutoTagging](
		"name", "google-auto-tagging", "aws-auto-tagging",
	)
}

func NewBetaV2FileUploadParamsExtensionAIAutoDescription() BetaV2FileUploadParamsExtensionAIAutoDescription {
	return BetaV2FileUploadParamsExtensionAIAutoDescription{
		Name: "ai-auto-description",
	}
}

// This struct has a constant value, construct it with
// [NewBetaV2FileUploadParamsExtensionAIAutoDescription].
type BetaV2FileUploadParamsExtensionAIAutoDescription struct {
	// Specifies the auto description extension.
	Name constant.AIAutoDescription `json:"name,required"`
	paramObj
}

func (r BetaV2FileUploadParamsExtensionAIAutoDescription) MarshalJSON() (data []byte, err error) {
	type shadow BetaV2FileUploadParamsExtensionAIAutoDescription
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaV2FileUploadParamsExtensionAIAutoDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configure pre-processing (`pre`) and post-processing (`post`) transformations.
//
//   - `pre` — applied before the file is uploaded to the Media Library.
//     Useful for reducing file size or applying basic optimizations upfront (e.g.,
//     resize, compress).
//
//   - `post` — applied immediately after upload.
//     Ideal for generating transformed versions (like video encodes or thumbnails)
//     in advance, so they're ready for delivery without delay.
//
// You can mix and match any combination of post-processing types.
type BetaV2FileUploadParamsTransformation struct {
	// Transformation string to apply before uploading the file to the Media Library.
	// Useful for optimizing files at ingestion.
	Pre param.Opt[string] `json:"pre,omitzero"`
	// List of transformations to apply _after_ the file is uploaded.
	// Each item must match one of the following types: `transformation`,
	// `gif-to-video`, `thumbnail`, `abs`.
	Post []BetaV2FileUploadParamsTransformationPostUnion `json:"post,omitzero"`
	paramObj
}

func (r BetaV2FileUploadParamsTransformation) MarshalJSON() (data []byte, err error) {
	type shadow BetaV2FileUploadParamsTransformation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaV2FileUploadParamsTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type BetaV2FileUploadParamsTransformationPostUnion struct {
	OfTransformation *BetaV2FileUploadParamsTransformationPostTransformation `json:",omitzero,inline"`
	OfGifToVideo     *BetaV2FileUploadParamsTransformationPostGifToVideo     `json:",omitzero,inline"`
	OfThumbnail      *BetaV2FileUploadParamsTransformationPostThumbnail      `json:",omitzero,inline"`
	OfAbs            *BetaV2FileUploadParamsTransformationPostAbs            `json:",omitzero,inline"`
	paramUnion
}

func (u BetaV2FileUploadParamsTransformationPostUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfTransformation, u.OfGifToVideo, u.OfThumbnail, u.OfAbs)
}
func (u *BetaV2FileUploadParamsTransformationPostUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *BetaV2FileUploadParamsTransformationPostUnion) asAny() any {
	if !param.IsOmitted(u.OfTransformation) {
		return u.OfTransformation
	} else if !param.IsOmitted(u.OfGifToVideo) {
		return u.OfGifToVideo
	} else if !param.IsOmitted(u.OfThumbnail) {
		return u.OfThumbnail
	} else if !param.IsOmitted(u.OfAbs) {
		return u.OfAbs
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BetaV2FileUploadParamsTransformationPostUnion) GetProtocol() *string {
	if vt := u.OfAbs; vt != nil {
		return &vt.Protocol
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BetaV2FileUploadParamsTransformationPostUnion) GetType() *string {
	if vt := u.OfTransformation; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfGifToVideo; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfThumbnail; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAbs; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u BetaV2FileUploadParamsTransformationPostUnion) GetValue() *string {
	if vt := u.OfTransformation; vt != nil {
		return (*string)(&vt.Value)
	} else if vt := u.OfGifToVideo; vt != nil && vt.Value.Valid() {
		return &vt.Value.Value
	} else if vt := u.OfThumbnail; vt != nil && vt.Value.Valid() {
		return &vt.Value.Value
	} else if vt := u.OfAbs; vt != nil {
		return (*string)(&vt.Value)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[BetaV2FileUploadParamsTransformationPostUnion](
		"type",
		apijson.Discriminator[BetaV2FileUploadParamsTransformationPostTransformation]("transformation"),
		apijson.Discriminator[BetaV2FileUploadParamsTransformationPostGifToVideo]("gif-to-video"),
		apijson.Discriminator[BetaV2FileUploadParamsTransformationPostThumbnail]("thumbnail"),
		apijson.Discriminator[BetaV2FileUploadParamsTransformationPostAbs]("abs"),
	)
}

// The properties Type, Value are required.
type BetaV2FileUploadParamsTransformationPostTransformation struct {
	// Transformation string (e.g. `w-200,h-200`).
	// Same syntax as ImageKit URL-based transformations.
	Value string `json:"value,required"`
	// Transformation type.
	//
	// This field can be elided, and will marshal its zero value as "transformation".
	Type constant.Transformation `json:"type,required"`
	paramObj
}

func (r BetaV2FileUploadParamsTransformationPostTransformation) MarshalJSON() (data []byte, err error) {
	type shadow BetaV2FileUploadParamsTransformationPostTransformation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaV2FileUploadParamsTransformationPostTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type BetaV2FileUploadParamsTransformationPostGifToVideo struct {
	// Optional transformation string to apply to the output video.
	// **Example**: `q-80`
	Value param.Opt[string] `json:"value,omitzero"`
	// Converts an animated GIF into an MP4.
	//
	// This field can be elided, and will marshal its zero value as "gif-to-video".
	Type constant.GifToVideo `json:"type,required"`
	paramObj
}

func (r BetaV2FileUploadParamsTransformationPostGifToVideo) MarshalJSON() (data []byte, err error) {
	type shadow BetaV2FileUploadParamsTransformationPostGifToVideo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaV2FileUploadParamsTransformationPostGifToVideo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type BetaV2FileUploadParamsTransformationPostThumbnail struct {
	// Optional transformation string.
	// **Example**: `w-150,h-150`
	Value param.Opt[string] `json:"value,omitzero"`
	// Generates a thumbnail image.
	//
	// This field can be elided, and will marshal its zero value as "thumbnail".
	Type constant.Thumbnail `json:"type,required"`
	paramObj
}

func (r BetaV2FileUploadParamsTransformationPostThumbnail) MarshalJSON() (data []byte, err error) {
	type shadow BetaV2FileUploadParamsTransformationPostThumbnail
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaV2FileUploadParamsTransformationPostThumbnail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Protocol, Type, Value are required.
type BetaV2FileUploadParamsTransformationPostAbs struct {
	// Streaming protocol to use (`hls` or `dash`).
	//
	// Any of "hls", "dash".
	Protocol string `json:"protocol,omitzero,required"`
	// List of different representations you want to create separated by an underscore.
	Value string `json:"value,required"`
	// Adaptive Bitrate Streaming (ABS) setup.
	//
	// This field can be elided, and will marshal its zero value as "abs".
	Type constant.Abs `json:"type,required"`
	paramObj
}

func (r BetaV2FileUploadParamsTransformationPostAbs) MarshalJSON() (data []byte, err error) {
	type shadow BetaV2FileUploadParamsTransformationPostAbs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BetaV2FileUploadParamsTransformationPostAbs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[BetaV2FileUploadParamsTransformationPostAbs](
		"protocol", "hls", "dash",
	)
}
