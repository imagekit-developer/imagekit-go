// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/stainless-sdks/imagekit-go/internal/apiform"
	"github.com/stainless-sdks/imagekit-go/internal/apijson"
	"github.com/stainless-sdks/imagekit-go/internal/requestconfig"
	"github.com/stainless-sdks/imagekit-go/option"
	"github.com/stainless-sdks/imagekit-go/packages/param"
	"github.com/stainless-sdks/imagekit-go/packages/respjson"
	"github.com/stainless-sdks/imagekit-go/shared"
)

// FileService contains methods and other services that help with interacting with
// the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileService] method instead.
type FileService struct {
	Options  []option.RequestOption
	Bulk     FileBulkService
	Versions FileVersionService
	Metadata FileMetadataService
}

// NewFileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFileService(opts ...option.RequestOption) (r FileService) {
	r = FileService{}
	r.Options = opts
	r.Bulk = NewFileBulkService(opts...)
	r.Versions = NewFileVersionService(opts...)
	r.Metadata = NewFileMetadataService(opts...)
	return
}

// This API updates the details or attributes of the current version of the file.
// You can update `tags`, `customCoordinates`, `customMetadata`, publication
// status, remove existing `AITags` and apply extensions using this API.
func (r *FileService) Update(ctx context.Context, fileID string, body FileUpdateParams, opts ...option.RequestOption) (res *FileUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return
	}
	path := fmt.Sprintf("v1/files/%s/details", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// This API deletes the file and all its file versions permanently.
//
// Note: If a file or specific transformation has been requested in the past, then
// the response is cached. Deleting a file does not purge the cache. You can purge
// the cache using purge cache API.
func (r *FileService) Delete(ctx context.Context, fileID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return
	}
	path := fmt.Sprintf("v1/files/%s", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// This will copy a file from one folder to another.
//
// Note: If any file at the destination has the same name as the source file, then
// the source file and its versions (if `includeFileVersions` is set to true) will
// be appended to the destination file version history.
func (r *FileService) Copy(ctx context.Context, body FileCopyParams, opts ...option.RequestOption) (res *FileCopyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/files/copy"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This API returns an object with details or attributes about the current version
// of the file.
func (r *FileService) Get(ctx context.Context, fileID string, opts ...option.RequestOption) (res *FileGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return
	}
	path := fmt.Sprintf("v1/files/%s/details", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This will move a file and all its versions from one folder to another.
//
// Note: If any file at the destination has the same name as the source file, then
// the source file and its versions will be appended to the destination file.
func (r *FileService) Move(ctx context.Context, body FileMoveParams, opts ...option.RequestOption) (res *FileMoveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/files/move"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// You can rename an already existing file in the media library using rename file
// API. This operation would rename all file versions of the file.
//
// Note: The old URLs will stop working. The file/file version URLs cached on CDN
// will continue to work unless a purge is requested.
func (r *FileService) Rename(ctx context.Context, body FileRenameParams, opts ...option.RequestOption) (res *FileRenameResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/files/rename"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// ImageKit.io allows you to upload files directly from both the server and client
// sides. For server-side uploads, private API key authentication is used. For
// client-side uploads, generate a one-time `token`, `signature`, and `expiration`
// from your secure backend using private API.
// [Learn more](/docs/api-reference/upload-file/upload-file#how-to-implement-client-side-file-upload)
// about how to implement client-side file upload.
//
// The [V2 API](/docs/api-reference/upload-file/upload-file-v2) enhances security
// by verifying the entire payload using JWT.
//
// **File size limit** \
// On the free plan, the maximum upload file sizes are 20MB for images, audio, and raw
// files and 100MB for videos. On the paid plan, these limits increase to 40MB for images,
// audio, and raw files and 2GB for videos. These limits can be further increased with
// higher-tier plans.
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
func (r *FileService) Upload(ctx context.Context, body FileUploadParams, opts ...option.RequestOption) (res *FileUploadResponse, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://upload.imagekit.io/")}, opts...)
	path := "api/v1/files/upload"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type FileUpdateResponse struct {
	// An array of tags assigned to the file by auto tagging.
	AITags []FileUpdateResponseAITag `json:"AITags,nullable"`
	// Date and time when the file was uploaded. The date and time is in ISO8601
	// format.
	CreatedAt string `json:"createdAt"`
	// An string with custom coordinates of the file.
	CustomCoordinates string `json:"customCoordinates,nullable"`
	// An object with custom metadata for the file.
	CustomMetadata  any                               `json:"customMetadata"`
	ExtensionStatus FileUpdateResponseExtensionStatus `json:"extensionStatus"`
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
	VersionInfo FileUpdateResponseVersionInfo `json:"versionInfo"`
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
func (r FileUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *FileUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileUpdateResponseAITag struct {
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
func (r FileUpdateResponseAITag) RawJSON() string { return r.JSON.raw }
func (r *FileUpdateResponseAITag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileUpdateResponseExtensionStatus struct {
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
func (r FileUpdateResponseExtensionStatus) RawJSON() string { return r.JSON.raw }
func (r *FileUpdateResponseExtensionStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object with details of the file version.
type FileUpdateResponseVersionInfo struct {
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
func (r FileUpdateResponseVersionInfo) RawJSON() string { return r.JSON.raw }
func (r *FileUpdateResponseVersionInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileCopyResponse = any

// Object containing details of a file or file version.
type FileGetResponse struct {
	// An array of tags assigned to the file by auto tagging.
	AITags []FileGetResponseAITag `json:"AITags,nullable"`
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
	VersionInfo FileGetResponseVersionInfo `json:"versionInfo"`
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
func (r FileGetResponse) RawJSON() string { return r.JSON.raw }
func (r *FileGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileGetResponseAITag struct {
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
func (r FileGetResponseAITag) RawJSON() string { return r.JSON.raw }
func (r *FileGetResponseAITag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object with details of the file version.
type FileGetResponseVersionInfo struct {
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
func (r FileGetResponseVersionInfo) RawJSON() string { return r.JSON.raw }
func (r *FileGetResponseVersionInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileMoveResponse = any

type FileRenameResponse struct {
	// Unique identifier of the purge request. This can be used to check the status of
	// the purge request.
	PurgeRequestID string `json:"purgeRequestId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PurgeRequestID respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileRenameResponse) RawJSON() string { return r.JSON.raw }
func (r *FileRenameResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Object containing details of a successful upload.
type FileUploadResponse struct {
	// An array of tags assigned to the uploaded file by auto tagging.
	AITags []FileUploadResponseAITag `json:"AITags,nullable"`
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
	CustomMetadata any `json:"customMetadata"`
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
	ExtensionStatus FileUploadResponseExtensionStatus `json:"extensionStatus"`
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
	Metadata FileUploadResponseMetadata `json:"metadata"`
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
	VersionInfo FileUploadResponseVersionInfo `json:"versionInfo"`
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
func (r FileUploadResponse) RawJSON() string { return r.JSON.raw }
func (r *FileUploadResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileUploadResponseAITag struct {
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
func (r FileUploadResponseAITag) RawJSON() string { return r.JSON.raw }
func (r *FileUploadResponseAITag) UnmarshalJSON(data []byte) error {
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
type FileUploadResponseExtensionStatus struct {
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
func (r FileUploadResponseExtensionStatus) RawJSON() string { return r.JSON.raw }
func (r *FileUploadResponseExtensionStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Legacy metadata. Send `metadata` in `responseFields` in API request to get
// metadata in the upload API response.
type FileUploadResponseMetadata struct {
	// The audio codec used in the video (only for video).
	AudioCodec string `json:"audioCodec"`
	// The bit rate of the video in kbps (only for video).
	BitRate int64 `json:"bitRate"`
	// The density of the image in DPI.
	Density int64 `json:"density"`
	// The duration of the video in seconds (only for video).
	Duration int64                          `json:"duration"`
	Exif     FileUploadResponseMetadataExif `json:"exif"`
	// The format of the file (e.g., 'jpg', 'mp4').
	Format string `json:"format"`
	// Indicates if the image has a color profile.
	HasColorProfile bool `json:"hasColorProfile"`
	// Indicates if the image contains transparent areas.
	HasTransparency bool `json:"hasTransparency"`
	// The height of the image or video in pixels.
	Height int64 `json:"height"`
	// Perceptual hash of the image.
	PHash string `json:"pHash"`
	// The quality indicator of the image.
	Quality int64 `json:"quality"`
	// The file size in bytes.
	Size int64 `json:"size"`
	// The video codec used in the video (only for video).
	VideoCodec string `json:"videoCodec"`
	// The width of the image or video in pixels.
	Width int64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudioCodec      respjson.Field
		BitRate         respjson.Field
		Density         respjson.Field
		Duration        respjson.Field
		Exif            respjson.Field
		Format          respjson.Field
		HasColorProfile respjson.Field
		HasTransparency respjson.Field
		Height          respjson.Field
		PHash           respjson.Field
		Quality         respjson.Field
		Size            respjson.Field
		VideoCodec      respjson.Field
		Width           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileUploadResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *FileUploadResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileUploadResponseMetadataExif struct {
	// Object containing Exif details.
	Exif shared.ExifDetails `json:"exif"`
	// Object containing GPS information.
	Gps shared.Gps `json:"gps"`
	// Object containing EXIF image information.
	Image shared.ExifImage `json:"image"`
	// JSON object.
	Interoperability shared.Interoperability `json:"interoperability"`
	Makernote        map[string]any          `json:"makernote"`
	// Object containing Thumbnail information.
	Thumbnail shared.Thumbnail `json:"thumbnail"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Exif             respjson.Field
		Gps              respjson.Field
		Image            respjson.Field
		Interoperability respjson.Field
		Makernote        respjson.Field
		Thumbnail        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileUploadResponseMetadataExif) RawJSON() string { return r.JSON.raw }
func (r *FileUploadResponseMetadataExif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object containing the file or file version's `id` (versionId) and `name`.
type FileUploadResponseVersionInfo struct {
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
func (r FileUploadResponseVersionInfo) RawJSON() string { return r.JSON.raw }
func (r *FileUploadResponseVersionInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileUpdateParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfUpdateFileDetails *FileUpdateParamsBodyUpdateFileDetails `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfChangePublicationStatus *FileUpdateParamsBodyChangePublicationStatus `json:",inline"`

	paramObj
}

func (u FileUpdateParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfUpdateFileDetails, u.OfChangePublicationStatus)
}
func (r *FileUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileUpdateParamsBodyUpdateFileDetails struct {
	// Define an important area in the image in the format `x,y,width,height` e.g.
	// `10,10,100,100`. Send `null` to unset this value.
	CustomCoordinates param.Opt[string] `json:"customCoordinates,omitzero"`
	// Optional text to describe the contents of the file.
	Description param.Opt[string] `json:"description,omitzero"`
	// The final status of extensions after they have completed execution will be
	// delivered to this endpoint as a POST request.
	// [Learn more](/docs/api-reference/digital-asset-management-dam/managing-assets/update-file-details#webhook-payload-structure)
	// about the webhook payload structure.
	WebhookURL param.Opt[string] `json:"webhookUrl,omitzero" format:"uri"`
	// An array of tags associated with the file, such as `["tag1", "tag2"]`. Send
	// `null` to unset all tags associated with the file.
	Tags []string `json:"tags,omitzero"`
	// A key-value data to be associated with the asset. To unset a key, send `null`
	// value for that key. Before setting any custom metadata on an asset you have to
	// create the field using custom metadata fields API.
	CustomMetadata any `json:"customMetadata,omitzero"`
	// Array of extensions to be applied to the asset. Each extension can be configured
	// with specific parameters based on the extension type.
	Extensions []FileUpdateParamsBodyUpdateFileDetailsExtensionUnion `json:"extensions,omitzero"`
	// An array of AITags associated with the file that you want to remove, e.g.
	// `["car", "vehicle", "motorsports"]`.
	//
	// If you want to remove all AITags associated with the file, send a string -
	// "all".
	//
	// Note: The remove operation for `AITags` executes before any of the `extensions`
	// are processed.
	RemoveAITags FileUpdateParamsBodyUpdateFileDetailsRemoveAITagsUnion `json:"removeAITags,omitzero"`
	paramObj
}

func (r FileUpdateParamsBodyUpdateFileDetails) MarshalJSON() (data []byte, err error) {
	type shadow FileUpdateParamsBodyUpdateFileDetails
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUpdateParamsBodyUpdateFileDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FileUpdateParamsBodyUpdateFileDetailsExtensionUnion struct {
	OfRemoveBackground *FileUpdateParamsBodyUpdateFileDetailsExtensionRemoveBackground `json:",omitzero,inline"`
	OfAutoTagging      *FileUpdateParamsBodyUpdateFileDetailsExtensionAutoTagging      `json:",omitzero,inline"`
	OfAutoDescription  *FileUpdateParamsBodyUpdateFileDetailsExtensionAutoDescription  `json:",omitzero,inline"`
	paramUnion
}

func (u FileUpdateParamsBodyUpdateFileDetailsExtensionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRemoveBackground, u.OfAutoTagging, u.OfAutoDescription)
}
func (u *FileUpdateParamsBodyUpdateFileDetailsExtensionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FileUpdateParamsBodyUpdateFileDetailsExtensionUnion) asAny() any {
	if !param.IsOmitted(u.OfRemoveBackground) {
		return u.OfRemoveBackground
	} else if !param.IsOmitted(u.OfAutoTagging) {
		return u.OfAutoTagging
	} else if !param.IsOmitted(u.OfAutoDescription) {
		return u.OfAutoDescription
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FileUpdateParamsBodyUpdateFileDetailsExtensionUnion) GetOptions() *FileUpdateParamsBodyUpdateFileDetailsExtensionRemoveBackgroundOptions {
	if vt := u.OfRemoveBackground; vt != nil {
		return &vt.Options
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FileUpdateParamsBodyUpdateFileDetailsExtensionUnion) GetMaxTags() *int64 {
	if vt := u.OfAutoTagging; vt != nil {
		return &vt.MaxTags
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FileUpdateParamsBodyUpdateFileDetailsExtensionUnion) GetMinConfidence() *int64 {
	if vt := u.OfAutoTagging; vt != nil {
		return &vt.MinConfidence
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FileUpdateParamsBodyUpdateFileDetailsExtensionUnion) GetName() *string {
	if vt := u.OfRemoveBackground; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfAutoTagging; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfAutoDescription; vt != nil {
		return (*string)(&vt.Name)
	}
	return nil
}

// The property Name is required.
type FileUpdateParamsBodyUpdateFileDetailsExtensionRemoveBackground struct {
	// Specifies the background removal extension.
	//
	// Any of "remove-bg".
	Name    string                                                                `json:"name,omitzero,required"`
	Options FileUpdateParamsBodyUpdateFileDetailsExtensionRemoveBackgroundOptions `json:"options,omitzero"`
	paramObj
}

func (r FileUpdateParamsBodyUpdateFileDetailsExtensionRemoveBackground) MarshalJSON() (data []byte, err error) {
	type shadow FileUpdateParamsBodyUpdateFileDetailsExtensionRemoveBackground
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUpdateParamsBodyUpdateFileDetailsExtensionRemoveBackground) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FileUpdateParamsBodyUpdateFileDetailsExtensionRemoveBackground](
		"name", "remove-bg",
	)
}

type FileUpdateParamsBodyUpdateFileDetailsExtensionRemoveBackgroundOptions struct {
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

func (r FileUpdateParamsBodyUpdateFileDetailsExtensionRemoveBackgroundOptions) MarshalJSON() (data []byte, err error) {
	type shadow FileUpdateParamsBodyUpdateFileDetailsExtensionRemoveBackgroundOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUpdateParamsBodyUpdateFileDetailsExtensionRemoveBackgroundOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MaxTags, MinConfidence, Name are required.
type FileUpdateParamsBodyUpdateFileDetailsExtensionAutoTagging struct {
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

func (r FileUpdateParamsBodyUpdateFileDetailsExtensionAutoTagging) MarshalJSON() (data []byte, err error) {
	type shadow FileUpdateParamsBodyUpdateFileDetailsExtensionAutoTagging
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUpdateParamsBodyUpdateFileDetailsExtensionAutoTagging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FileUpdateParamsBodyUpdateFileDetailsExtensionAutoTagging](
		"name", "google-auto-tagging", "aws-auto-tagging",
	)
}

// The property Name is required.
type FileUpdateParamsBodyUpdateFileDetailsExtensionAutoDescription struct {
	// Specifies the auto description extension.
	//
	// Any of "ai-auto-description".
	Name string `json:"name,omitzero,required"`
	paramObj
}

func (r FileUpdateParamsBodyUpdateFileDetailsExtensionAutoDescription) MarshalJSON() (data []byte, err error) {
	type shadow FileUpdateParamsBodyUpdateFileDetailsExtensionAutoDescription
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUpdateParamsBodyUpdateFileDetailsExtensionAutoDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FileUpdateParamsBodyUpdateFileDetailsExtensionAutoDescription](
		"name", "ai-auto-description",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FileUpdateParamsBodyUpdateFileDetailsRemoveAITagsUnion struct {
	OfStringArray []string `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfFileUpdatesBodyUpdateFileDetailsRemoveAITagsString)
	OfFileUpdatesBodyUpdateFileDetailsRemoveAITagsString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u FileUpdateParamsBodyUpdateFileDetailsRemoveAITagsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfStringArray, u.OfFileUpdatesBodyUpdateFileDetailsRemoveAITagsString)
}
func (u *FileUpdateParamsBodyUpdateFileDetailsRemoveAITagsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FileUpdateParamsBodyUpdateFileDetailsRemoveAITagsUnion) asAny() any {
	if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	} else if !param.IsOmitted(u.OfFileUpdatesBodyUpdateFileDetailsRemoveAITagsString) {
		return &u.OfFileUpdatesBodyUpdateFileDetailsRemoveAITagsString
	}
	return nil
}

type FileUpdateParamsBodyUpdateFileDetailsRemoveAITagsString string

const (
	FileUpdateParamsBodyUpdateFileDetailsRemoveAITagsStringAll FileUpdateParamsBodyUpdateFileDetailsRemoveAITagsString = "all"
)

type FileUpdateParamsBodyChangePublicationStatus struct {
	// Configure the publication status of a file and its versions.
	Publish FileUpdateParamsBodyChangePublicationStatusPublish `json:"publish,omitzero"`
	paramObj
}

func (r FileUpdateParamsBodyChangePublicationStatus) MarshalJSON() (data []byte, err error) {
	type shadow FileUpdateParamsBodyChangePublicationStatus
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUpdateParamsBodyChangePublicationStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configure the publication status of a file and its versions.
//
// The property IsPublished is required.
type FileUpdateParamsBodyChangePublicationStatusPublish struct {
	// Set to `true` to publish the file. Set to `false` to unpublish the file.
	IsPublished bool `json:"isPublished,required"`
	// Set to `true` to publish/unpublish all versions of the file. Set to `false` to
	// publish/unpublish only the current version of the file.
	IncludeFileVersions param.Opt[bool] `json:"includeFileVersions,omitzero"`
	paramObj
}

func (r FileUpdateParamsBodyChangePublicationStatusPublish) MarshalJSON() (data []byte, err error) {
	type shadow FileUpdateParamsBodyChangePublicationStatusPublish
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUpdateParamsBodyChangePublicationStatusPublish) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileCopyParams struct {
	// Full path to the folder you want to copy the above file into.
	DestinationPath string `json:"destinationPath,required"`
	// The full path of the file you want to copy.
	SourceFilePath string `json:"sourceFilePath,required"`
	// Option to copy all versions of a file. By default, only the current version of
	// the file is copied. When set to true, all versions of the file will be copied.
	// Default value - `false`.
	IncludeFileVersions param.Opt[bool] `json:"includeFileVersions,omitzero"`
	paramObj
}

func (r FileCopyParams) MarshalJSON() (data []byte, err error) {
	type shadow FileCopyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileCopyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileMoveParams struct {
	// Full path to the folder you want to move the above file into.
	DestinationPath string `json:"destinationPath,required"`
	// The full path of the file you want to move.
	SourceFilePath string `json:"sourceFilePath,required"`
	paramObj
}

func (r FileMoveParams) MarshalJSON() (data []byte, err error) {
	type shadow FileMoveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileMoveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileRenameParams struct {
	// The full path of the file you want to rename.
	FilePath string `json:"filePath,required"`
	// The new name of the file. A filename can contain:
	//
	// Alphanumeric Characters: `a-z`, `A-Z`, `0-9` (including Unicode letters, marks,
	// and numerals in other languages). Special Characters: `.`, `_`, and `-`.
	//
	// Any other character, including space, will be replaced by `_`.
	NewFileName string `json:"newFileName,required"`
	// Option to purge cache for the old file and its versions' URLs.
	//
	// When set to true, it will internally issue a purge cache request on CDN to
	// remove cached content of old file and its versions. This purge request is
	// counted against your monthly purge quota.
	//
	// Note: If the old file were accessible at
	// `https://ik.imagekit.io/demo/old-filename.jpg`, a purge cache request would be
	// issued against `https://ik.imagekit.io/demo/old-filename.jpg*` (with a wildcard
	// at the end). It will remove the file and its versions' URLs and any
	// transformations made using query parameters on this file or its versions.
	// However, the cache for file transformations made using path parameters will
	// persist. You can purge them using the purge API. For more details, refer to the
	// purge API documentation.
	//
	// Default value - `false`
	PurgeCache param.Opt[bool] `json:"purgeCache,omitzero"`
	paramObj
}

func (r FileRenameParams) MarshalJSON() (data []byte, err error) {
	type shadow FileRenameParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileRenameParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileUploadParams struct {
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
	// The name with which the file has to be uploaded. The file name can contain:
	//
	// - Alphanumeric Characters: `a-z`, `A-Z`, `0-9`.
	// - Special Characters: `.`, `-`
	//
	// Any other character including space will be replaced by `_`
	FileName string `json:"fileName,required"`
	// A unique value that the ImageKit.io server will use to recognize and prevent
	// subsequent retries for the same request. We suggest using V4 UUIDs, or another
	// random string with enough entropy to avoid collisions. This field is only
	// required for authentication when uploading a file from the client side.
	//
	// **Note**: Sending a value that has been used in the past will result in a
	// validation error. Even if your previous request resulted in an error, you should
	// always send a new value for this field.
	Token param.Opt[string] `json:"token,omitzero"`
	// Server-side checks to run on the asset. Read more about
	// [Upload API checks](/docs/api-reference/upload-file/upload-file#upload-api-checks).
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
	// The time until your signature is valid. It must be a
	// [Unix time](https://en.wikipedia.org/wiki/Unix_time) in less than 1 hour into
	// the future. It should be in seconds. This field is only required for
	// authentication when uploading a file from the client side.
	Expire param.Opt[int64] `json:"expire,omitzero"`
	// The folder path in which the image has to be uploaded. If the folder(s) didn't
	// exist before, a new folder(s) is created.
	//
	// The folder name can contain:
	//
	// - Alphanumeric Characters: `a-z` , `A-Z` , `0-9`
	// - Special Characters: `/` , `_` , `-`
	//
	// Using multiple `/` creates a nested folder.
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
	// Your ImageKit.io public key. This field is only required for authentication when
	// uploading a file from the client side.
	PublicKey param.Opt[string] `json:"publicKey,omitzero"`
	// HMAC-SHA1 digest of the token+expire using your ImageKit.io private API key as a
	// key. Learn how to create a signature on the page below. This should be in
	// lowercase.
	//
	// Signature must be calculated on the server-side. This field is only required for
	// authentication when uploading a file from the client side.
	Signature param.Opt[string] `json:"signature,omitzero"`
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
	Extensions []FileUploadParamsExtensionUnion `json:"extensions,omitzero"`
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
	Transformation FileUploadParamsTransformation `json:"transformation,omitzero"`
	paramObj
}

func (r FileUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
type FileUploadParamsExtensionUnion struct {
	OfRemoveBackground *FileUploadParamsExtensionRemoveBackground `json:",omitzero,inline"`
	OfAutoTagging      *FileUploadParamsExtensionAutoTagging      `json:",omitzero,inline"`
	OfAutoDescription  *FileUploadParamsExtensionAutoDescription  `json:",omitzero,inline"`
	paramUnion
}

func (u FileUploadParamsExtensionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRemoveBackground, u.OfAutoTagging, u.OfAutoDescription)
}
func (u *FileUploadParamsExtensionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FileUploadParamsExtensionUnion) asAny() any {
	if !param.IsOmitted(u.OfRemoveBackground) {
		return u.OfRemoveBackground
	} else if !param.IsOmitted(u.OfAutoTagging) {
		return u.OfAutoTagging
	} else if !param.IsOmitted(u.OfAutoDescription) {
		return u.OfAutoDescription
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FileUploadParamsExtensionUnion) GetOptions() *FileUploadParamsExtensionRemoveBackgroundOptions {
	if vt := u.OfRemoveBackground; vt != nil {
		return &vt.Options
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FileUploadParamsExtensionUnion) GetMaxTags() *int64 {
	if vt := u.OfAutoTagging; vt != nil {
		return &vt.MaxTags
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FileUploadParamsExtensionUnion) GetMinConfidence() *int64 {
	if vt := u.OfAutoTagging; vt != nil {
		return &vt.MinConfidence
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FileUploadParamsExtensionUnion) GetName() *string {
	if vt := u.OfRemoveBackground; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfAutoTagging; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfAutoDescription; vt != nil {
		return (*string)(&vt.Name)
	}
	return nil
}

// The property Name is required.
type FileUploadParamsExtensionRemoveBackground struct {
	// Specifies the background removal extension.
	//
	// Any of "remove-bg".
	Name    string                                           `json:"name,omitzero,required"`
	Options FileUploadParamsExtensionRemoveBackgroundOptions `json:"options,omitzero"`
	paramObj
}

func (r FileUploadParamsExtensionRemoveBackground) MarshalJSON() (data []byte, err error) {
	type shadow FileUploadParamsExtensionRemoveBackground
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUploadParamsExtensionRemoveBackground) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FileUploadParamsExtensionRemoveBackground](
		"name", "remove-bg",
	)
}

type FileUploadParamsExtensionRemoveBackgroundOptions struct {
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

func (r FileUploadParamsExtensionRemoveBackgroundOptions) MarshalJSON() (data []byte, err error) {
	type shadow FileUploadParamsExtensionRemoveBackgroundOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUploadParamsExtensionRemoveBackgroundOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MaxTags, MinConfidence, Name are required.
type FileUploadParamsExtensionAutoTagging struct {
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

func (r FileUploadParamsExtensionAutoTagging) MarshalJSON() (data []byte, err error) {
	type shadow FileUploadParamsExtensionAutoTagging
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUploadParamsExtensionAutoTagging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FileUploadParamsExtensionAutoTagging](
		"name", "google-auto-tagging", "aws-auto-tagging",
	)
}

// The property Name is required.
type FileUploadParamsExtensionAutoDescription struct {
	// Specifies the auto description extension.
	//
	// Any of "ai-auto-description".
	Name string `json:"name,omitzero,required"`
	paramObj
}

func (r FileUploadParamsExtensionAutoDescription) MarshalJSON() (data []byte, err error) {
	type shadow FileUploadParamsExtensionAutoDescription
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUploadParamsExtensionAutoDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FileUploadParamsExtensionAutoDescription](
		"name", "ai-auto-description",
	)
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
type FileUploadParamsTransformation struct {
	// Transformation string to apply before uploading the file to the Media Library.
	// Useful for optimizing files at ingestion.
	Pre param.Opt[string] `json:"pre,omitzero"`
	// List of transformations to apply _after_ the file is uploaded.
	// Each item must match one of the following types: `transformation`,
	// `gif-to-video`, `thumbnail`, `abs`.
	Post []FileUploadParamsTransformationPostUnion `json:"post,omitzero"`
	paramObj
}

func (r FileUploadParamsTransformation) MarshalJSON() (data []byte, err error) {
	type shadow FileUploadParamsTransformation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUploadParamsTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FileUploadParamsTransformationPostUnion struct {
	OfSimplePostTransformation *FileUploadParamsTransformationPostSimplePostTransformation `json:",omitzero,inline"`
	OfConvertGifToVideo        *FileUploadParamsTransformationPostConvertGifToVideo        `json:",omitzero,inline"`
	OfGenerateAThumbnail       *FileUploadParamsTransformationPostGenerateAThumbnail       `json:",omitzero,inline"`
	OfAdaptiveBitrateStreaming *FileUploadParamsTransformationPostAdaptiveBitrateStreaming `json:",omitzero,inline"`
	paramUnion
}

func (u FileUploadParamsTransformationPostUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSimplePostTransformation, u.OfConvertGifToVideo, u.OfGenerateAThumbnail, u.OfAdaptiveBitrateStreaming)
}
func (u *FileUploadParamsTransformationPostUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FileUploadParamsTransformationPostUnion) asAny() any {
	if !param.IsOmitted(u.OfSimplePostTransformation) {
		return u.OfSimplePostTransformation
	} else if !param.IsOmitted(u.OfConvertGifToVideo) {
		return u.OfConvertGifToVideo
	} else if !param.IsOmitted(u.OfGenerateAThumbnail) {
		return u.OfGenerateAThumbnail
	} else if !param.IsOmitted(u.OfAdaptiveBitrateStreaming) {
		return u.OfAdaptiveBitrateStreaming
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FileUploadParamsTransformationPostUnion) GetProtocol() *string {
	if vt := u.OfAdaptiveBitrateStreaming; vt != nil {
		return &vt.Protocol
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FileUploadParamsTransformationPostUnion) GetType() *string {
	if vt := u.OfSimplePostTransformation; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfConvertGifToVideo; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfGenerateAThumbnail; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAdaptiveBitrateStreaming; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FileUploadParamsTransformationPostUnion) GetValue() *string {
	if vt := u.OfSimplePostTransformation; vt != nil {
		return (*string)(&vt.Value)
	} else if vt := u.OfConvertGifToVideo; vt != nil && vt.Value.Valid() {
		return &vt.Value.Value
	} else if vt := u.OfGenerateAThumbnail; vt != nil && vt.Value.Valid() {
		return &vt.Value.Value
	} else if vt := u.OfAdaptiveBitrateStreaming; vt != nil {
		return (*string)(&vt.Value)
	}
	return nil
}

// The properties Type, Value are required.
type FileUploadParamsTransformationPostSimplePostTransformation struct {
	// Transformation type.
	//
	// Any of "transformation".
	Type string `json:"type,omitzero,required"`
	// Transformation string (e.g. `w-200,h-200`).
	// Same syntax as ImageKit URL-based transformations.
	Value string `json:"value,required"`
	paramObj
}

func (r FileUploadParamsTransformationPostSimplePostTransformation) MarshalJSON() (data []byte, err error) {
	type shadow FileUploadParamsTransformationPostSimplePostTransformation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUploadParamsTransformationPostSimplePostTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FileUploadParamsTransformationPostSimplePostTransformation](
		"type", "transformation",
	)
}

// The property Type is required.
type FileUploadParamsTransformationPostConvertGifToVideo struct {
	// Converts an animated GIF into an MP4.
	//
	// Any of "gif-to-video".
	Type string `json:"type,omitzero,required"`
	// Optional transformation string to apply to the output video.
	// **Example**: `q-80`
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r FileUploadParamsTransformationPostConvertGifToVideo) MarshalJSON() (data []byte, err error) {
	type shadow FileUploadParamsTransformationPostConvertGifToVideo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUploadParamsTransformationPostConvertGifToVideo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FileUploadParamsTransformationPostConvertGifToVideo](
		"type", "gif-to-video",
	)
}

// The property Type is required.
type FileUploadParamsTransformationPostGenerateAThumbnail struct {
	// Generates a thumbnail image.
	//
	// Any of "thumbnail".
	Type string `json:"type,omitzero,required"`
	// Optional transformation string.
	// **Example**: `w-150,h-150`
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r FileUploadParamsTransformationPostGenerateAThumbnail) MarshalJSON() (data []byte, err error) {
	type shadow FileUploadParamsTransformationPostGenerateAThumbnail
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUploadParamsTransformationPostGenerateAThumbnail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FileUploadParamsTransformationPostGenerateAThumbnail](
		"type", "thumbnail",
	)
}

// The properties Protocol, Type, Value are required.
type FileUploadParamsTransformationPostAdaptiveBitrateStreaming struct {
	// Streaming protocol to use (`hls` or `dash`).
	//
	// Any of "hls", "dash".
	Protocol string `json:"protocol,omitzero,required"`
	// Adaptive Bitrate Streaming (ABS) setup.
	//
	// Any of "abs".
	Type string `json:"type,omitzero,required"`
	// List of different representations you want to create separated by an underscore.
	Value string `json:"value,required"`
	paramObj
}

func (r FileUploadParamsTransformationPostAdaptiveBitrateStreaming) MarshalJSON() (data []byte, err error) {
	type shadow FileUploadParamsTransformationPostAdaptiveBitrateStreaming
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUploadParamsTransformationPostAdaptiveBitrateStreaming) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FileUploadParamsTransformationPostAdaptiveBitrateStreaming](
		"protocol", "hls", "dash",
	)
	apijson.RegisterFieldValidator[FileUploadParamsTransformationPostAdaptiveBitrateStreaming](
		"type", "abs",
	)
}
