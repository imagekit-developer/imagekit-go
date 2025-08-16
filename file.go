// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/imagekit-go/internal/apiform"
	"github.com/stainless-sdks/imagekit-go/internal/apijson"
	"github.com/stainless-sdks/imagekit-go/internal/apiquery"
	"github.com/stainless-sdks/imagekit-go/internal/requestconfig"
	"github.com/stainless-sdks/imagekit-go/option"
	"github.com/stainless-sdks/imagekit-go/packages/param"
	"github.com/stainless-sdks/imagekit-go/packages/respjson"
)

// FileService contains methods and other services that help with interacting with
// the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileService] method instead.
type FileService struct {
	Options  []option.RequestOption
	Details  FileDetailService
	Batch    FileBatchService
	Versions FileVersionService
	Purge    FilePurgeService
	Metadata FileMetadataService
}

// NewFileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFileService(opts ...option.RequestOption) (r FileService) {
	r = FileService{}
	r.Options = opts
	r.Details = NewFileDetailService(opts...)
	r.Batch = NewFileBatchService(opts...)
	r.Versions = NewFileVersionService(opts...)
	r.Purge = NewFilePurgeService(opts...)
	r.Metadata = NewFileMetadataService(opts...)
	return
}

// This API can list all the uploaded files and folders in your ImageKit.io media
// library. In addition, you can fine-tune your query by specifying various filters
// by generating a query string in a Lucene-like syntax and provide this generated
// string as the value of the `searchQuery`.
func (r *FileService) List(ctx context.Context, query FileListParams, opts ...option.RequestOption) (res *[]FileListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
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

// This API adds tags to multiple files in bulk. A maximum of 50 files can be
// specified at a time.
func (r *FileService) AddTags(ctx context.Context, body FileAddTagsParams, opts ...option.RequestOption) (res *FileAddTagsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/files/addTags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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

// This API removes AITags from multiple files in bulk. A maximum of 50 files can
// be specified at a time.
func (r *FileService) RemoveAITags(ctx context.Context, body FileRemoveAITagsParams, opts ...option.RequestOption) (res *FileRemoveAITagsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/files/removeAITags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This API removes tags from multiple files in bulk. A maximum of 50 files can be
// specified at a time.
func (r *FileService) RemoveTags(ctx context.Context, body FileRemoveTagsParams, opts ...option.RequestOption) (res *FileRemoveTagsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/files/removeTags"
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
func (r *FileService) UploadV1(ctx context.Context, body FileUploadV1Params, opts ...option.RequestOption) (res *FileUploadV1Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://upload.imagekit.io/")}, opts...)
	path := "api/v1/files/upload"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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
func (r *FileService) UploadV2(ctx context.Context, body FileUploadV2Params, opts ...option.RequestOption) (res *FileUploadV2Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithBaseURL("https://upload.imagekit.io/")}, opts...)
	path := "api/v2/files/upload"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Object containing Exif details.
type ExifDetails struct {
	ApertureValue            float64 `json:"ApertureValue"`
	ColorSpace               int64   `json:"ColorSpace"`
	CreateDate               string  `json:"CreateDate"`
	CustomRendered           int64   `json:"CustomRendered"`
	DateTimeOriginal         string  `json:"DateTimeOriginal"`
	ExifImageHeight          int64   `json:"ExifImageHeight"`
	ExifImageWidth           int64   `json:"ExifImageWidth"`
	ExifVersion              string  `json:"ExifVersion"`
	ExposureCompensation     float64 `json:"ExposureCompensation"`
	ExposureMode             int64   `json:"ExposureMode"`
	ExposureProgram          int64   `json:"ExposureProgram"`
	ExposureTime             float64 `json:"ExposureTime"`
	Flash                    int64   `json:"Flash"`
	FlashpixVersion          string  `json:"FlashpixVersion"`
	FNumber                  float64 `json:"FNumber"`
	FocalLength              int64   `json:"FocalLength"`
	FocalPlaneResolutionUnit int64   `json:"FocalPlaneResolutionUnit"`
	FocalPlaneXResolution    float64 `json:"FocalPlaneXResolution"`
	FocalPlaneYResolution    float64 `json:"FocalPlaneYResolution"`
	InteropOffset            int64   `json:"InteropOffset"`
	ISO                      int64   `json:"ISO"`
	MeteringMode             int64   `json:"MeteringMode"`
	SceneCaptureType         int64   `json:"SceneCaptureType"`
	ShutterSpeedValue        float64 `json:"ShutterSpeedValue"`
	SubSecTime               string  `json:"SubSecTime"`
	WhiteBalance             int64   `json:"WhiteBalance"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApertureValue            respjson.Field
		ColorSpace               respjson.Field
		CreateDate               respjson.Field
		CustomRendered           respjson.Field
		DateTimeOriginal         respjson.Field
		ExifImageHeight          respjson.Field
		ExifImageWidth           respjson.Field
		ExifVersion              respjson.Field
		ExposureCompensation     respjson.Field
		ExposureMode             respjson.Field
		ExposureProgram          respjson.Field
		ExposureTime             respjson.Field
		Flash                    respjson.Field
		FlashpixVersion          respjson.Field
		FNumber                  respjson.Field
		FocalLength              respjson.Field
		FocalPlaneResolutionUnit respjson.Field
		FocalPlaneXResolution    respjson.Field
		FocalPlaneYResolution    respjson.Field
		InteropOffset            respjson.Field
		ISO                      respjson.Field
		MeteringMode             respjson.Field
		SceneCaptureType         respjson.Field
		ShutterSpeedValue        respjson.Field
		SubSecTime               respjson.Field
		WhiteBalance             respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExifDetails) RawJSON() string { return r.JSON.raw }
func (r *ExifDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Object containing EXIF image information.
type ExifImage struct {
	ExifOffset       int64  `json:"ExifOffset"`
	GpsInfo          int64  `json:"GPSInfo"`
	Make             string `json:"Make"`
	Model            string `json:"Model"`
	ModifyDate       string `json:"ModifyDate"`
	Orientation      int64  `json:"Orientation"`
	ResolutionUnit   int64  `json:"ResolutionUnit"`
	Software         string `json:"Software"`
	XResolution      int64  `json:"XResolution"`
	YCbCrPositioning int64  `json:"YCbCrPositioning"`
	YResolution      int64  `json:"YResolution"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExifOffset       respjson.Field
		GpsInfo          respjson.Field
		Make             respjson.Field
		Model            respjson.Field
		ModifyDate       respjson.Field
		Orientation      respjson.Field
		ResolutionUnit   respjson.Field
		Software         respjson.Field
		XResolution      respjson.Field
		YCbCrPositioning respjson.Field
		YResolution      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExifImage) RawJSON() string { return r.JSON.raw }
func (r *ExifImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Object containing GPS information.
type Gps struct {
	GpsVersionID []int64 `json:"GPSVersionID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GpsVersionID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Gps) RawJSON() string { return r.JSON.raw }
func (r *Gps) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// JSON object.
type Interoperability struct {
	InteropIndex   string `json:"InteropIndex"`
	InteropVersion string `json:"InteropVersion"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InteropIndex   respjson.Field
		InteropVersion respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Interoperability) RawJSON() string { return r.JSON.raw }
func (r *Interoperability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Object containing Thumbnail information.
type Thumbnail struct {
	Compression     int64 `json:"Compression"`
	ResolutionUnit  int64 `json:"ResolutionUnit"`
	ThumbnailLength int64 `json:"ThumbnailLength"`
	ThumbnailOffset int64 `json:"ThumbnailOffset"`
	XResolution     int64 `json:"XResolution"`
	YResolution     int64 `json:"YResolution"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Compression     respjson.Field
		ResolutionUnit  respjson.Field
		ThumbnailLength respjson.Field
		ThumbnailOffset respjson.Field
		XResolution     respjson.Field
		YResolution     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Thumbnail) RawJSON() string { return r.JSON.raw }
func (r *Thumbnail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Object containing details of a file or file version.
type FileListResponse struct {
	// An array of tags assigned to the file by auto tagging.
	AITags []FileListResponseAITag `json:"AITags,nullable"`
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
	VersionInfo FileListResponseVersionInfo `json:"versionInfo"`
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
func (r FileListResponse) RawJSON() string { return r.JSON.raw }
func (r *FileListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileListResponseAITag struct {
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
func (r FileListResponseAITag) RawJSON() string { return r.JSON.raw }
func (r *FileListResponseAITag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object with details of the file version.
type FileListResponseVersionInfo struct {
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
func (r FileListResponseVersionInfo) RawJSON() string { return r.JSON.raw }
func (r *FileListResponseVersionInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileAddTagsResponse struct {
	// An array of fileIds that in which tags were successfully added.
	SuccessfullyUpdatedFileIDs []string `json:"successfullyUpdatedFileIds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SuccessfullyUpdatedFileIDs respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileAddTagsResponse) RawJSON() string { return r.JSON.raw }
func (r *FileAddTagsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileCopyResponse = any

type FileMoveResponse = any

type FileRemoveAITagsResponse struct {
	// An array of fileIds that in which AITags were successfully removed.
	SuccessfullyUpdatedFileIDs []string `json:"successfullyUpdatedFileIds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SuccessfullyUpdatedFileIDs respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileRemoveAITagsResponse) RawJSON() string { return r.JSON.raw }
func (r *FileRemoveAITagsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileRemoveTagsResponse struct {
	// An array of fileIds that in which tags were successfully removed.
	SuccessfullyUpdatedFileIDs []string `json:"successfullyUpdatedFileIds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SuccessfullyUpdatedFileIDs respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileRemoveTagsResponse) RawJSON() string { return r.JSON.raw }
func (r *FileRemoveTagsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
type FileUploadV1Response struct {
	// An array of tags assigned to the uploaded file by auto tagging.
	AITags []FileUploadV1ResponseAITag `json:"AITags,nullable"`
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
	EmbeddedMetadata FileUploadV1ResponseEmbeddedMetadata `json:"embeddedMetadata"`
	// Extension names with their processing status at the time of completion of the
	// request. It could have one of the following status values:
	//
	// `success`: The extension has been successfully applied. `failed`: The extension
	// has failed and will not be retried. `pending`: The extension will finish
	// processing in some time. On completion, the final status (success / failed) will
	// be sent to the `webhookUrl` provided.
	//
	// If no extension was requested, then this parameter is not returned.
	ExtensionStatus FileUploadV1ResponseExtensionStatus `json:"extensionStatus"`
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
	Metadata FileUploadV1ResponseMetadata `json:"metadata"`
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
	VersionInfo FileUploadV1ResponseVersionInfo `json:"versionInfo"`
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
func (r FileUploadV1Response) RawJSON() string { return r.JSON.raw }
func (r *FileUploadV1Response) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileUploadV1ResponseAITag struct {
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
func (r FileUploadV1ResponseAITag) RawJSON() string { return r.JSON.raw }
func (r *FileUploadV1ResponseAITag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Consolidated embedded metadata associated with the file. It includes exif, iptc,
// and xmp data. Send `embeddedMetadata` in `responseFields` in API request to get
// embeddedMetadata in the upload API response.
type FileUploadV1ResponseEmbeddedMetadata struct {
	AboutCvTermCvID                    string    `json:"AboutCvTermCvId"`
	AboutCvTermID                      string    `json:"AboutCvTermId"`
	AboutCvTermName                    string    `json:"AboutCvTermName"`
	AboutCvTermRefinedAbout            string    `json:"AboutCvTermRefinedAbout"`
	AdditionalModelInformation         string    `json:"AdditionalModelInformation"`
	ApplicationRecordVersion           int64     `json:"ApplicationRecordVersion"`
	Artist                             string    `json:"Artist"`
	ArtworkCircaDateCreated            string    `json:"ArtworkCircaDateCreated"`
	ArtworkContentDescription          string    `json:"ArtworkContentDescription"`
	ArtworkContributionDescription     string    `json:"ArtworkContributionDescription"`
	ArtworkCopyrightNotice             string    `json:"ArtworkCopyrightNotice"`
	ArtworkCopyrightOwnerID            string    `json:"ArtworkCopyrightOwnerID"`
	ArtworkCopyrightOwnerName          string    `json:"ArtworkCopyrightOwnerName"`
	ArtworkCreator                     []string  `json:"ArtworkCreator"`
	ArtworkCreatorID                   []string  `json:"ArtworkCreatorID"`
	ArtworkDateCreated                 time.Time `json:"ArtworkDateCreated" format:"date-time"`
	ArtworkLicensorID                  string    `json:"ArtworkLicensorID"`
	ArtworkLicensorName                string    `json:"ArtworkLicensorName"`
	ArtworkPhysicalDescription         string    `json:"ArtworkPhysicalDescription"`
	ArtworkSource                      string    `json:"ArtworkSource"`
	ArtworkSourceInventoryNo           string    `json:"ArtworkSourceInventoryNo"`
	ArtworkSourceInvURL                string    `json:"ArtworkSourceInvURL"`
	ArtworkStylePeriod                 []string  `json:"ArtworkStylePeriod"`
	ArtworkTitle                       string    `json:"ArtworkTitle"`
	AuthorsPosition                    string    `json:"AuthorsPosition"`
	Byline                             string    `json:"Byline"`
	BylineTitle                        string    `json:"BylineTitle"`
	Caption                            string    `json:"Caption"`
	CaptionAbstract                    string    `json:"CaptionAbstract"`
	CaptionWriter                      string    `json:"CaptionWriter"`
	City                               string    `json:"City"`
	ColorSpace                         string    `json:"ColorSpace"`
	ComponentsConfiguration            string    `json:"ComponentsConfiguration"`
	Copyright                          string    `json:"Copyright"`
	CopyrightNotice                    string    `json:"CopyrightNotice"`
	CopyrightOwnerID                   []string  `json:"CopyrightOwnerID"`
	CopyrightOwnerName                 []string  `json:"CopyrightOwnerName"`
	Country                            string    `json:"Country"`
	CountryCode                        string    `json:"CountryCode"`
	CountryPrimaryLocationCode         string    `json:"CountryPrimaryLocationCode"`
	CountryPrimaryLocationName         string    `json:"CountryPrimaryLocationName"`
	Creator                            string    `json:"Creator"`
	CreatorAddress                     string    `json:"CreatorAddress"`
	CreatorCity                        string    `json:"CreatorCity"`
	CreatorCountry                     string    `json:"CreatorCountry"`
	CreatorPostalCode                  string    `json:"CreatorPostalCode"`
	CreatorRegion                      string    `json:"CreatorRegion"`
	CreatorWorkEmail                   string    `json:"CreatorWorkEmail"`
	CreatorWorkTelephone               string    `json:"CreatorWorkTelephone"`
	CreatorWorkURL                     string    `json:"CreatorWorkURL"`
	Credit                             string    `json:"Credit"`
	DateCreated                        time.Time `json:"DateCreated" format:"date-time"`
	DateTimeCreated                    time.Time `json:"DateTimeCreated" format:"date-time"`
	DateTimeOriginal                   time.Time `json:"DateTimeOriginal" format:"date-time"`
	Description                        string    `json:"Description"`
	DigitalImageGuid                   string    `json:"DigitalImageGUID"`
	DigitalSourceType                  string    `json:"DigitalSourceType"`
	EmbeddedEncodedRightsExpr          string    `json:"EmbeddedEncodedRightsExpr"`
	EmbeddedEncodedRightsExprLangID    string    `json:"EmbeddedEncodedRightsExprLangID"`
	EmbeddedEncodedRightsExprType      string    `json:"EmbeddedEncodedRightsExprType"`
	Event                              string    `json:"Event"`
	ExifVersion                        string    `json:"ExifVersion"`
	FlashpixVersion                    string    `json:"FlashpixVersion"`
	GenreCvID                          string    `json:"GenreCvId"`
	GenreCvTermID                      string    `json:"GenreCvTermId"`
	GenreCvTermName                    string    `json:"GenreCvTermName"`
	GenreCvTermRefinedAbout            string    `json:"GenreCvTermRefinedAbout"`
	Headline                           string    `json:"Headline"`
	ImageCreatorID                     string    `json:"ImageCreatorID"`
	ImageCreatorImageID                string    `json:"ImageCreatorImageID"`
	ImageCreatorName                   string    `json:"ImageCreatorName"`
	ImageDescription                   string    `json:"ImageDescription"`
	ImageRegionBoundaryH               []float64 `json:"ImageRegionBoundaryH"`
	ImageRegionBoundaryRx              []float64 `json:"ImageRegionBoundaryRx"`
	ImageRegionBoundaryShape           []string  `json:"ImageRegionBoundaryShape"`
	ImageRegionBoundaryUnit            []string  `json:"ImageRegionBoundaryUnit"`
	ImageRegionBoundaryVerticesX       []float64 `json:"ImageRegionBoundaryVerticesX"`
	ImageRegionBoundaryVerticesY       []float64 `json:"ImageRegionBoundaryVerticesY"`
	ImageRegionBoundaryW               []float64 `json:"ImageRegionBoundaryW"`
	ImageRegionBoundaryX               []float64 `json:"ImageRegionBoundaryX"`
	ImageRegionBoundaryY               []float64 `json:"ImageRegionBoundaryY"`
	ImageRegionCtypeIdentifier         []string  `json:"ImageRegionCtypeIdentifier"`
	ImageRegionCtypeName               []string  `json:"ImageRegionCtypeName"`
	ImageRegionID                      []string  `json:"ImageRegionID"`
	ImageRegionName                    []string  `json:"ImageRegionName"`
	ImageRegionOrganisationInImageName []string  `json:"ImageRegionOrganisationInImageName"`
	ImageRegionPersonInImage           []string  `json:"ImageRegionPersonInImage"`
	ImageRegionRoleIdentifier          []string  `json:"ImageRegionRoleIdentifier"`
	ImageRegionRoleName                []string  `json:"ImageRegionRoleName"`
	ImageSupplierID                    string    `json:"ImageSupplierID"`
	ImageSupplierImageID               string    `json:"ImageSupplierImageID"`
	ImageSupplierName                  string    `json:"ImageSupplierName"`
	Instructions                       string    `json:"Instructions"`
	IntellectualGenre                  string    `json:"IntellectualGenre"`
	Keywords                           []string  `json:"Keywords"`
	LicensorCity                       []string  `json:"LicensorCity"`
	LicensorCountry                    []string  `json:"LicensorCountry"`
	LicensorEmail                      []string  `json:"LicensorEmail"`
	LicensorExtendedAddress            []string  `json:"LicensorExtendedAddress"`
	LicensorID                         []string  `json:"LicensorID"`
	LicensorName                       []string  `json:"LicensorName"`
	LicensorPostalCode                 []string  `json:"LicensorPostalCode"`
	LicensorRegion                     []string  `json:"LicensorRegion"`
	LicensorStreetAddress              []string  `json:"LicensorStreetAddress"`
	LicensorTelephone1                 []string  `json:"LicensorTelephone1"`
	LicensorTelephone2                 []string  `json:"LicensorTelephone2"`
	LicensorURL                        []string  `json:"LicensorURL"`
	LinkedEncodedRightsExpr            string    `json:"LinkedEncodedRightsExpr"`
	LinkedEncodedRightsExprLangID      string    `json:"LinkedEncodedRightsExprLangID"`
	LinkedEncodedRightsExprType        string    `json:"LinkedEncodedRightsExprType"`
	Location                           string    `json:"Location"`
	LocationCreatedCity                string    `json:"LocationCreatedCity"`
	LocationCreatedCountryCode         string    `json:"LocationCreatedCountryCode"`
	LocationCreatedCountryName         string    `json:"LocationCreatedCountryName"`
	LocationCreatedGpsAltitude         string    `json:"LocationCreatedGPSAltitude"`
	LocationCreatedGpsLatitude         string    `json:"LocationCreatedGPSLatitude"`
	LocationCreatedGpsLongitude        string    `json:"LocationCreatedGPSLongitude"`
	LocationCreatedLocationID          string    `json:"LocationCreatedLocationId"`
	LocationCreatedLocationName        string    `json:"LocationCreatedLocationName"`
	LocationCreatedProvinceState       string    `json:"LocationCreatedProvinceState"`
	LocationCreatedSublocation         string    `json:"LocationCreatedSublocation"`
	LocationCreatedWorldRegion         string    `json:"LocationCreatedWorldRegion"`
	LocationShownCity                  []string  `json:"LocationShownCity"`
	LocationShownCountryCode           []string  `json:"LocationShownCountryCode"`
	LocationShownCountryName           []string  `json:"LocationShownCountryName"`
	LocationShownGpsAltitude           []string  `json:"LocationShownGPSAltitude"`
	LocationShownGpsLatitude           []string  `json:"LocationShownGPSLatitude"`
	LocationShownGpsLongitude          []string  `json:"LocationShownGPSLongitude"`
	LocationShownLocationID            []string  `json:"LocationShownLocationId"`
	LocationShownLocationName          []string  `json:"LocationShownLocationName"`
	LocationShownProvinceState         []string  `json:"LocationShownProvinceState"`
	LocationShownSublocation           []string  `json:"LocationShownSublocation"`
	LocationShownWorldRegion           []string  `json:"LocationShownWorldRegion"`
	MaxAvailHeight                     float64   `json:"MaxAvailHeight"`
	MaxAvailWidth                      float64   `json:"MaxAvailWidth"`
	ModelAge                           []float64 `json:"ModelAge"`
	ModelReleaseID                     []string  `json:"ModelReleaseID"`
	ObjectAttributeReference           string    `json:"ObjectAttributeReference"`
	ObjectName                         string    `json:"ObjectName"`
	OffsetTimeOriginal                 string    `json:"OffsetTimeOriginal"`
	OrganisationInImageCode            []string  `json:"OrganisationInImageCode"`
	OrganisationInImageName            []string  `json:"OrganisationInImageName"`
	Orientation                        string    `json:"Orientation"`
	OriginalTransmissionReference      string    `json:"OriginalTransmissionReference"`
	PersonInImage                      []string  `json:"PersonInImage"`
	PersonInImageCvTermCvID            []string  `json:"PersonInImageCvTermCvId"`
	PersonInImageCvTermID              []string  `json:"PersonInImageCvTermId"`
	PersonInImageCvTermName            []string  `json:"PersonInImageCvTermName"`
	PersonInImageCvTermRefinedAbout    []string  `json:"PersonInImageCvTermRefinedAbout"`
	PersonInImageDescription           []string  `json:"PersonInImageDescription"`
	PersonInImageID                    []string  `json:"PersonInImageId"`
	PersonInImageName                  []string  `json:"PersonInImageName"`
	ProductInImageDescription          []string  `json:"ProductInImageDescription"`
	ProductInImageGtin                 []float64 `json:"ProductInImageGTIN"`
	ProductInImageName                 []string  `json:"ProductInImageName"`
	PropertyReleaseID                  []string  `json:"PropertyReleaseID"`
	ProvinceState                      string    `json:"ProvinceState"`
	Rating                             int64     `json:"Rating"`
	RegistryEntryRole                  []string  `json:"RegistryEntryRole"`
	RegistryItemID                     []string  `json:"RegistryItemID"`
	RegistryOrganisationID             []string  `json:"RegistryOrganisationID"`
	ResolutionUnit                     string    `json:"ResolutionUnit"`
	Rights                             string    `json:"Rights"`
	Scene                              []string  `json:"Scene"`
	Source                             string    `json:"Source"`
	SpecialInstructions                string    `json:"SpecialInstructions"`
	State                              string    `json:"State"`
	Subject                            []string  `json:"Subject"`
	SubjectCode                        []string  `json:"SubjectCode"`
	SubjectReference                   []string  `json:"SubjectReference"`
	Sublocation                        string    `json:"Sublocation"`
	TimeCreated                        string    `json:"TimeCreated"`
	Title                              string    `json:"Title"`
	TransmissionReference              string    `json:"TransmissionReference"`
	UsageTerms                         string    `json:"UsageTerms"`
	WebStatement                       string    `json:"WebStatement"`
	Writer                             string    `json:"Writer"`
	WriterEditor                       string    `json:"WriterEditor"`
	XResolution                        float64   `json:"XResolution"`
	YResolution                        float64   `json:"YResolution"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AboutCvTermCvID                    respjson.Field
		AboutCvTermID                      respjson.Field
		AboutCvTermName                    respjson.Field
		AboutCvTermRefinedAbout            respjson.Field
		AdditionalModelInformation         respjson.Field
		ApplicationRecordVersion           respjson.Field
		Artist                             respjson.Field
		ArtworkCircaDateCreated            respjson.Field
		ArtworkContentDescription          respjson.Field
		ArtworkContributionDescription     respjson.Field
		ArtworkCopyrightNotice             respjson.Field
		ArtworkCopyrightOwnerID            respjson.Field
		ArtworkCopyrightOwnerName          respjson.Field
		ArtworkCreator                     respjson.Field
		ArtworkCreatorID                   respjson.Field
		ArtworkDateCreated                 respjson.Field
		ArtworkLicensorID                  respjson.Field
		ArtworkLicensorName                respjson.Field
		ArtworkPhysicalDescription         respjson.Field
		ArtworkSource                      respjson.Field
		ArtworkSourceInventoryNo           respjson.Field
		ArtworkSourceInvURL                respjson.Field
		ArtworkStylePeriod                 respjson.Field
		ArtworkTitle                       respjson.Field
		AuthorsPosition                    respjson.Field
		Byline                             respjson.Field
		BylineTitle                        respjson.Field
		Caption                            respjson.Field
		CaptionAbstract                    respjson.Field
		CaptionWriter                      respjson.Field
		City                               respjson.Field
		ColorSpace                         respjson.Field
		ComponentsConfiguration            respjson.Field
		Copyright                          respjson.Field
		CopyrightNotice                    respjson.Field
		CopyrightOwnerID                   respjson.Field
		CopyrightOwnerName                 respjson.Field
		Country                            respjson.Field
		CountryCode                        respjson.Field
		CountryPrimaryLocationCode         respjson.Field
		CountryPrimaryLocationName         respjson.Field
		Creator                            respjson.Field
		CreatorAddress                     respjson.Field
		CreatorCity                        respjson.Field
		CreatorCountry                     respjson.Field
		CreatorPostalCode                  respjson.Field
		CreatorRegion                      respjson.Field
		CreatorWorkEmail                   respjson.Field
		CreatorWorkTelephone               respjson.Field
		CreatorWorkURL                     respjson.Field
		Credit                             respjson.Field
		DateCreated                        respjson.Field
		DateTimeCreated                    respjson.Field
		DateTimeOriginal                   respjson.Field
		Description                        respjson.Field
		DigitalImageGuid                   respjson.Field
		DigitalSourceType                  respjson.Field
		EmbeddedEncodedRightsExpr          respjson.Field
		EmbeddedEncodedRightsExprLangID    respjson.Field
		EmbeddedEncodedRightsExprType      respjson.Field
		Event                              respjson.Field
		ExifVersion                        respjson.Field
		FlashpixVersion                    respjson.Field
		GenreCvID                          respjson.Field
		GenreCvTermID                      respjson.Field
		GenreCvTermName                    respjson.Field
		GenreCvTermRefinedAbout            respjson.Field
		Headline                           respjson.Field
		ImageCreatorID                     respjson.Field
		ImageCreatorImageID                respjson.Field
		ImageCreatorName                   respjson.Field
		ImageDescription                   respjson.Field
		ImageRegionBoundaryH               respjson.Field
		ImageRegionBoundaryRx              respjson.Field
		ImageRegionBoundaryShape           respjson.Field
		ImageRegionBoundaryUnit            respjson.Field
		ImageRegionBoundaryVerticesX       respjson.Field
		ImageRegionBoundaryVerticesY       respjson.Field
		ImageRegionBoundaryW               respjson.Field
		ImageRegionBoundaryX               respjson.Field
		ImageRegionBoundaryY               respjson.Field
		ImageRegionCtypeIdentifier         respjson.Field
		ImageRegionCtypeName               respjson.Field
		ImageRegionID                      respjson.Field
		ImageRegionName                    respjson.Field
		ImageRegionOrganisationInImageName respjson.Field
		ImageRegionPersonInImage           respjson.Field
		ImageRegionRoleIdentifier          respjson.Field
		ImageRegionRoleName                respjson.Field
		ImageSupplierID                    respjson.Field
		ImageSupplierImageID               respjson.Field
		ImageSupplierName                  respjson.Field
		Instructions                       respjson.Field
		IntellectualGenre                  respjson.Field
		Keywords                           respjson.Field
		LicensorCity                       respjson.Field
		LicensorCountry                    respjson.Field
		LicensorEmail                      respjson.Field
		LicensorExtendedAddress            respjson.Field
		LicensorID                         respjson.Field
		LicensorName                       respjson.Field
		LicensorPostalCode                 respjson.Field
		LicensorRegion                     respjson.Field
		LicensorStreetAddress              respjson.Field
		LicensorTelephone1                 respjson.Field
		LicensorTelephone2                 respjson.Field
		LicensorURL                        respjson.Field
		LinkedEncodedRightsExpr            respjson.Field
		LinkedEncodedRightsExprLangID      respjson.Field
		LinkedEncodedRightsExprType        respjson.Field
		Location                           respjson.Field
		LocationCreatedCity                respjson.Field
		LocationCreatedCountryCode         respjson.Field
		LocationCreatedCountryName         respjson.Field
		LocationCreatedGpsAltitude         respjson.Field
		LocationCreatedGpsLatitude         respjson.Field
		LocationCreatedGpsLongitude        respjson.Field
		LocationCreatedLocationID          respjson.Field
		LocationCreatedLocationName        respjson.Field
		LocationCreatedProvinceState       respjson.Field
		LocationCreatedSublocation         respjson.Field
		LocationCreatedWorldRegion         respjson.Field
		LocationShownCity                  respjson.Field
		LocationShownCountryCode           respjson.Field
		LocationShownCountryName           respjson.Field
		LocationShownGpsAltitude           respjson.Field
		LocationShownGpsLatitude           respjson.Field
		LocationShownGpsLongitude          respjson.Field
		LocationShownLocationID            respjson.Field
		LocationShownLocationName          respjson.Field
		LocationShownProvinceState         respjson.Field
		LocationShownSublocation           respjson.Field
		LocationShownWorldRegion           respjson.Field
		MaxAvailHeight                     respjson.Field
		MaxAvailWidth                      respjson.Field
		ModelAge                           respjson.Field
		ModelReleaseID                     respjson.Field
		ObjectAttributeReference           respjson.Field
		ObjectName                         respjson.Field
		OffsetTimeOriginal                 respjson.Field
		OrganisationInImageCode            respjson.Field
		OrganisationInImageName            respjson.Field
		Orientation                        respjson.Field
		OriginalTransmissionReference      respjson.Field
		PersonInImage                      respjson.Field
		PersonInImageCvTermCvID            respjson.Field
		PersonInImageCvTermID              respjson.Field
		PersonInImageCvTermName            respjson.Field
		PersonInImageCvTermRefinedAbout    respjson.Field
		PersonInImageDescription           respjson.Field
		PersonInImageID                    respjson.Field
		PersonInImageName                  respjson.Field
		ProductInImageDescription          respjson.Field
		ProductInImageGtin                 respjson.Field
		ProductInImageName                 respjson.Field
		PropertyReleaseID                  respjson.Field
		ProvinceState                      respjson.Field
		Rating                             respjson.Field
		RegistryEntryRole                  respjson.Field
		RegistryItemID                     respjson.Field
		RegistryOrganisationID             respjson.Field
		ResolutionUnit                     respjson.Field
		Rights                             respjson.Field
		Scene                              respjson.Field
		Source                             respjson.Field
		SpecialInstructions                respjson.Field
		State                              respjson.Field
		Subject                            respjson.Field
		SubjectCode                        respjson.Field
		SubjectReference                   respjson.Field
		Sublocation                        respjson.Field
		TimeCreated                        respjson.Field
		Title                              respjson.Field
		TransmissionReference              respjson.Field
		UsageTerms                         respjson.Field
		WebStatement                       respjson.Field
		Writer                             respjson.Field
		WriterEditor                       respjson.Field
		XResolution                        respjson.Field
		YResolution                        respjson.Field
		ExtraFields                        map[string]respjson.Field
		raw                                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileUploadV1ResponseEmbeddedMetadata) RawJSON() string { return r.JSON.raw }
func (r *FileUploadV1ResponseEmbeddedMetadata) UnmarshalJSON(data []byte) error {
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
type FileUploadV1ResponseExtensionStatus struct {
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
func (r FileUploadV1ResponseExtensionStatus) RawJSON() string { return r.JSON.raw }
func (r *FileUploadV1ResponseExtensionStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Legacy metadata. Send `metadata` in `responseFields` in API request to get
// metadata in the upload API response.
type FileUploadV1ResponseMetadata struct {
	// The audio codec used in the video (only for video).
	AudioCodec string `json:"audioCodec"`
	// The bit rate of the video in kbps (only for video).
	BitRate int64 `json:"bitRate"`
	// The density of the image in DPI.
	Density int64 `json:"density"`
	// The duration of the video in seconds (only for video).
	Duration int64                            `json:"duration"`
	Exif     FileUploadV1ResponseMetadataExif `json:"exif"`
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
func (r FileUploadV1ResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *FileUploadV1ResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileUploadV1ResponseMetadataExif struct {
	// Object containing Exif details.
	Exif ExifDetails `json:"exif"`
	// Object containing GPS information.
	Gps Gps `json:"gps"`
	// Object containing EXIF image information.
	Image ExifImage `json:"image"`
	// JSON object.
	Interoperability Interoperability `json:"interoperability"`
	Makernote        map[string]any   `json:"makernote"`
	// Object containing Thumbnail information.
	Thumbnail Thumbnail `json:"thumbnail"`
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
func (r FileUploadV1ResponseMetadataExif) RawJSON() string { return r.JSON.raw }
func (r *FileUploadV1ResponseMetadataExif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object containing the file or file version's `id` (versionId) and `name`.
type FileUploadV1ResponseVersionInfo struct {
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
func (r FileUploadV1ResponseVersionInfo) RawJSON() string { return r.JSON.raw }
func (r *FileUploadV1ResponseVersionInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Object containing details of a successful upload.
type FileUploadV2Response struct {
	// An array of tags assigned to the uploaded file by auto tagging.
	AITags []FileUploadV2ResponseAITag `json:"AITags,nullable"`
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
	EmbeddedMetadata FileUploadV2ResponseEmbeddedMetadata `json:"embeddedMetadata"`
	// Extension names with their processing status at the time of completion of the
	// request. It could have one of the following status values:
	//
	// `success`: The extension has been successfully applied. `failed`: The extension
	// has failed and will not be retried. `pending`: The extension will finish
	// processing in some time. On completion, the final status (success / failed) will
	// be sent to the `webhookUrl` provided.
	//
	// If no extension was requested, then this parameter is not returned.
	ExtensionStatus FileUploadV2ResponseExtensionStatus `json:"extensionStatus"`
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
	Metadata FileUploadV2ResponseMetadata `json:"metadata"`
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
	VersionInfo FileUploadV2ResponseVersionInfo `json:"versionInfo"`
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
func (r FileUploadV2Response) RawJSON() string { return r.JSON.raw }
func (r *FileUploadV2Response) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileUploadV2ResponseAITag struct {
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
func (r FileUploadV2ResponseAITag) RawJSON() string { return r.JSON.raw }
func (r *FileUploadV2ResponseAITag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Consolidated embedded metadata associated with the file. It includes exif, iptc,
// and xmp data. Send `embeddedMetadata` in `responseFields` in API request to get
// embeddedMetadata in the upload API response.
type FileUploadV2ResponseEmbeddedMetadata struct {
	AboutCvTermCvID                    string    `json:"AboutCvTermCvId"`
	AboutCvTermID                      string    `json:"AboutCvTermId"`
	AboutCvTermName                    string    `json:"AboutCvTermName"`
	AboutCvTermRefinedAbout            string    `json:"AboutCvTermRefinedAbout"`
	AdditionalModelInformation         string    `json:"AdditionalModelInformation"`
	ApplicationRecordVersion           int64     `json:"ApplicationRecordVersion"`
	Artist                             string    `json:"Artist"`
	ArtworkCircaDateCreated            string    `json:"ArtworkCircaDateCreated"`
	ArtworkContentDescription          string    `json:"ArtworkContentDescription"`
	ArtworkContributionDescription     string    `json:"ArtworkContributionDescription"`
	ArtworkCopyrightNotice             string    `json:"ArtworkCopyrightNotice"`
	ArtworkCopyrightOwnerID            string    `json:"ArtworkCopyrightOwnerID"`
	ArtworkCopyrightOwnerName          string    `json:"ArtworkCopyrightOwnerName"`
	ArtworkCreator                     []string  `json:"ArtworkCreator"`
	ArtworkCreatorID                   []string  `json:"ArtworkCreatorID"`
	ArtworkDateCreated                 time.Time `json:"ArtworkDateCreated" format:"date-time"`
	ArtworkLicensorID                  string    `json:"ArtworkLicensorID"`
	ArtworkLicensorName                string    `json:"ArtworkLicensorName"`
	ArtworkPhysicalDescription         string    `json:"ArtworkPhysicalDescription"`
	ArtworkSource                      string    `json:"ArtworkSource"`
	ArtworkSourceInventoryNo           string    `json:"ArtworkSourceInventoryNo"`
	ArtworkSourceInvURL                string    `json:"ArtworkSourceInvURL"`
	ArtworkStylePeriod                 []string  `json:"ArtworkStylePeriod"`
	ArtworkTitle                       string    `json:"ArtworkTitle"`
	AuthorsPosition                    string    `json:"AuthorsPosition"`
	Byline                             string    `json:"Byline"`
	BylineTitle                        string    `json:"BylineTitle"`
	Caption                            string    `json:"Caption"`
	CaptionAbstract                    string    `json:"CaptionAbstract"`
	CaptionWriter                      string    `json:"CaptionWriter"`
	City                               string    `json:"City"`
	ColorSpace                         string    `json:"ColorSpace"`
	ComponentsConfiguration            string    `json:"ComponentsConfiguration"`
	Copyright                          string    `json:"Copyright"`
	CopyrightNotice                    string    `json:"CopyrightNotice"`
	CopyrightOwnerID                   []string  `json:"CopyrightOwnerID"`
	CopyrightOwnerName                 []string  `json:"CopyrightOwnerName"`
	Country                            string    `json:"Country"`
	CountryCode                        string    `json:"CountryCode"`
	CountryPrimaryLocationCode         string    `json:"CountryPrimaryLocationCode"`
	CountryPrimaryLocationName         string    `json:"CountryPrimaryLocationName"`
	Creator                            string    `json:"Creator"`
	CreatorAddress                     string    `json:"CreatorAddress"`
	CreatorCity                        string    `json:"CreatorCity"`
	CreatorCountry                     string    `json:"CreatorCountry"`
	CreatorPostalCode                  string    `json:"CreatorPostalCode"`
	CreatorRegion                      string    `json:"CreatorRegion"`
	CreatorWorkEmail                   string    `json:"CreatorWorkEmail"`
	CreatorWorkTelephone               string    `json:"CreatorWorkTelephone"`
	CreatorWorkURL                     string    `json:"CreatorWorkURL"`
	Credit                             string    `json:"Credit"`
	DateCreated                        time.Time `json:"DateCreated" format:"date-time"`
	DateTimeCreated                    time.Time `json:"DateTimeCreated" format:"date-time"`
	DateTimeOriginal                   time.Time `json:"DateTimeOriginal" format:"date-time"`
	Description                        string    `json:"Description"`
	DigitalImageGuid                   string    `json:"DigitalImageGUID"`
	DigitalSourceType                  string    `json:"DigitalSourceType"`
	EmbeddedEncodedRightsExpr          string    `json:"EmbeddedEncodedRightsExpr"`
	EmbeddedEncodedRightsExprLangID    string    `json:"EmbeddedEncodedRightsExprLangID"`
	EmbeddedEncodedRightsExprType      string    `json:"EmbeddedEncodedRightsExprType"`
	Event                              string    `json:"Event"`
	ExifVersion                        string    `json:"ExifVersion"`
	FlashpixVersion                    string    `json:"FlashpixVersion"`
	GenreCvID                          string    `json:"GenreCvId"`
	GenreCvTermID                      string    `json:"GenreCvTermId"`
	GenreCvTermName                    string    `json:"GenreCvTermName"`
	GenreCvTermRefinedAbout            string    `json:"GenreCvTermRefinedAbout"`
	Headline                           string    `json:"Headline"`
	ImageCreatorID                     string    `json:"ImageCreatorID"`
	ImageCreatorImageID                string    `json:"ImageCreatorImageID"`
	ImageCreatorName                   string    `json:"ImageCreatorName"`
	ImageDescription                   string    `json:"ImageDescription"`
	ImageRegionBoundaryH               []float64 `json:"ImageRegionBoundaryH"`
	ImageRegionBoundaryRx              []float64 `json:"ImageRegionBoundaryRx"`
	ImageRegionBoundaryShape           []string  `json:"ImageRegionBoundaryShape"`
	ImageRegionBoundaryUnit            []string  `json:"ImageRegionBoundaryUnit"`
	ImageRegionBoundaryVerticesX       []float64 `json:"ImageRegionBoundaryVerticesX"`
	ImageRegionBoundaryVerticesY       []float64 `json:"ImageRegionBoundaryVerticesY"`
	ImageRegionBoundaryW               []float64 `json:"ImageRegionBoundaryW"`
	ImageRegionBoundaryX               []float64 `json:"ImageRegionBoundaryX"`
	ImageRegionBoundaryY               []float64 `json:"ImageRegionBoundaryY"`
	ImageRegionCtypeIdentifier         []string  `json:"ImageRegionCtypeIdentifier"`
	ImageRegionCtypeName               []string  `json:"ImageRegionCtypeName"`
	ImageRegionID                      []string  `json:"ImageRegionID"`
	ImageRegionName                    []string  `json:"ImageRegionName"`
	ImageRegionOrganisationInImageName []string  `json:"ImageRegionOrganisationInImageName"`
	ImageRegionPersonInImage           []string  `json:"ImageRegionPersonInImage"`
	ImageRegionRoleIdentifier          []string  `json:"ImageRegionRoleIdentifier"`
	ImageRegionRoleName                []string  `json:"ImageRegionRoleName"`
	ImageSupplierID                    string    `json:"ImageSupplierID"`
	ImageSupplierImageID               string    `json:"ImageSupplierImageID"`
	ImageSupplierName                  string    `json:"ImageSupplierName"`
	Instructions                       string    `json:"Instructions"`
	IntellectualGenre                  string    `json:"IntellectualGenre"`
	Keywords                           []string  `json:"Keywords"`
	LicensorCity                       []string  `json:"LicensorCity"`
	LicensorCountry                    []string  `json:"LicensorCountry"`
	LicensorEmail                      []string  `json:"LicensorEmail"`
	LicensorExtendedAddress            []string  `json:"LicensorExtendedAddress"`
	LicensorID                         []string  `json:"LicensorID"`
	LicensorName                       []string  `json:"LicensorName"`
	LicensorPostalCode                 []string  `json:"LicensorPostalCode"`
	LicensorRegion                     []string  `json:"LicensorRegion"`
	LicensorStreetAddress              []string  `json:"LicensorStreetAddress"`
	LicensorTelephone1                 []string  `json:"LicensorTelephone1"`
	LicensorTelephone2                 []string  `json:"LicensorTelephone2"`
	LicensorURL                        []string  `json:"LicensorURL"`
	LinkedEncodedRightsExpr            string    `json:"LinkedEncodedRightsExpr"`
	LinkedEncodedRightsExprLangID      string    `json:"LinkedEncodedRightsExprLangID"`
	LinkedEncodedRightsExprType        string    `json:"LinkedEncodedRightsExprType"`
	Location                           string    `json:"Location"`
	LocationCreatedCity                string    `json:"LocationCreatedCity"`
	LocationCreatedCountryCode         string    `json:"LocationCreatedCountryCode"`
	LocationCreatedCountryName         string    `json:"LocationCreatedCountryName"`
	LocationCreatedGpsAltitude         string    `json:"LocationCreatedGPSAltitude"`
	LocationCreatedGpsLatitude         string    `json:"LocationCreatedGPSLatitude"`
	LocationCreatedGpsLongitude        string    `json:"LocationCreatedGPSLongitude"`
	LocationCreatedLocationID          string    `json:"LocationCreatedLocationId"`
	LocationCreatedLocationName        string    `json:"LocationCreatedLocationName"`
	LocationCreatedProvinceState       string    `json:"LocationCreatedProvinceState"`
	LocationCreatedSublocation         string    `json:"LocationCreatedSublocation"`
	LocationCreatedWorldRegion         string    `json:"LocationCreatedWorldRegion"`
	LocationShownCity                  []string  `json:"LocationShownCity"`
	LocationShownCountryCode           []string  `json:"LocationShownCountryCode"`
	LocationShownCountryName           []string  `json:"LocationShownCountryName"`
	LocationShownGpsAltitude           []string  `json:"LocationShownGPSAltitude"`
	LocationShownGpsLatitude           []string  `json:"LocationShownGPSLatitude"`
	LocationShownGpsLongitude          []string  `json:"LocationShownGPSLongitude"`
	LocationShownLocationID            []string  `json:"LocationShownLocationId"`
	LocationShownLocationName          []string  `json:"LocationShownLocationName"`
	LocationShownProvinceState         []string  `json:"LocationShownProvinceState"`
	LocationShownSublocation           []string  `json:"LocationShownSublocation"`
	LocationShownWorldRegion           []string  `json:"LocationShownWorldRegion"`
	MaxAvailHeight                     float64   `json:"MaxAvailHeight"`
	MaxAvailWidth                      float64   `json:"MaxAvailWidth"`
	ModelAge                           []float64 `json:"ModelAge"`
	ModelReleaseID                     []string  `json:"ModelReleaseID"`
	ObjectAttributeReference           string    `json:"ObjectAttributeReference"`
	ObjectName                         string    `json:"ObjectName"`
	OffsetTimeOriginal                 string    `json:"OffsetTimeOriginal"`
	OrganisationInImageCode            []string  `json:"OrganisationInImageCode"`
	OrganisationInImageName            []string  `json:"OrganisationInImageName"`
	Orientation                        string    `json:"Orientation"`
	OriginalTransmissionReference      string    `json:"OriginalTransmissionReference"`
	PersonInImage                      []string  `json:"PersonInImage"`
	PersonInImageCvTermCvID            []string  `json:"PersonInImageCvTermCvId"`
	PersonInImageCvTermID              []string  `json:"PersonInImageCvTermId"`
	PersonInImageCvTermName            []string  `json:"PersonInImageCvTermName"`
	PersonInImageCvTermRefinedAbout    []string  `json:"PersonInImageCvTermRefinedAbout"`
	PersonInImageDescription           []string  `json:"PersonInImageDescription"`
	PersonInImageID                    []string  `json:"PersonInImageId"`
	PersonInImageName                  []string  `json:"PersonInImageName"`
	ProductInImageDescription          []string  `json:"ProductInImageDescription"`
	ProductInImageGtin                 []float64 `json:"ProductInImageGTIN"`
	ProductInImageName                 []string  `json:"ProductInImageName"`
	PropertyReleaseID                  []string  `json:"PropertyReleaseID"`
	ProvinceState                      string    `json:"ProvinceState"`
	Rating                             int64     `json:"Rating"`
	RegistryEntryRole                  []string  `json:"RegistryEntryRole"`
	RegistryItemID                     []string  `json:"RegistryItemID"`
	RegistryOrganisationID             []string  `json:"RegistryOrganisationID"`
	ResolutionUnit                     string    `json:"ResolutionUnit"`
	Rights                             string    `json:"Rights"`
	Scene                              []string  `json:"Scene"`
	Source                             string    `json:"Source"`
	SpecialInstructions                string    `json:"SpecialInstructions"`
	State                              string    `json:"State"`
	Subject                            []string  `json:"Subject"`
	SubjectCode                        []string  `json:"SubjectCode"`
	SubjectReference                   []string  `json:"SubjectReference"`
	Sublocation                        string    `json:"Sublocation"`
	TimeCreated                        string    `json:"TimeCreated"`
	Title                              string    `json:"Title"`
	TransmissionReference              string    `json:"TransmissionReference"`
	UsageTerms                         string    `json:"UsageTerms"`
	WebStatement                       string    `json:"WebStatement"`
	Writer                             string    `json:"Writer"`
	WriterEditor                       string    `json:"WriterEditor"`
	XResolution                        float64   `json:"XResolution"`
	YResolution                        float64   `json:"YResolution"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AboutCvTermCvID                    respjson.Field
		AboutCvTermID                      respjson.Field
		AboutCvTermName                    respjson.Field
		AboutCvTermRefinedAbout            respjson.Field
		AdditionalModelInformation         respjson.Field
		ApplicationRecordVersion           respjson.Field
		Artist                             respjson.Field
		ArtworkCircaDateCreated            respjson.Field
		ArtworkContentDescription          respjson.Field
		ArtworkContributionDescription     respjson.Field
		ArtworkCopyrightNotice             respjson.Field
		ArtworkCopyrightOwnerID            respjson.Field
		ArtworkCopyrightOwnerName          respjson.Field
		ArtworkCreator                     respjson.Field
		ArtworkCreatorID                   respjson.Field
		ArtworkDateCreated                 respjson.Field
		ArtworkLicensorID                  respjson.Field
		ArtworkLicensorName                respjson.Field
		ArtworkPhysicalDescription         respjson.Field
		ArtworkSource                      respjson.Field
		ArtworkSourceInventoryNo           respjson.Field
		ArtworkSourceInvURL                respjson.Field
		ArtworkStylePeriod                 respjson.Field
		ArtworkTitle                       respjson.Field
		AuthorsPosition                    respjson.Field
		Byline                             respjson.Field
		BylineTitle                        respjson.Field
		Caption                            respjson.Field
		CaptionAbstract                    respjson.Field
		CaptionWriter                      respjson.Field
		City                               respjson.Field
		ColorSpace                         respjson.Field
		ComponentsConfiguration            respjson.Field
		Copyright                          respjson.Field
		CopyrightNotice                    respjson.Field
		CopyrightOwnerID                   respjson.Field
		CopyrightOwnerName                 respjson.Field
		Country                            respjson.Field
		CountryCode                        respjson.Field
		CountryPrimaryLocationCode         respjson.Field
		CountryPrimaryLocationName         respjson.Field
		Creator                            respjson.Field
		CreatorAddress                     respjson.Field
		CreatorCity                        respjson.Field
		CreatorCountry                     respjson.Field
		CreatorPostalCode                  respjson.Field
		CreatorRegion                      respjson.Field
		CreatorWorkEmail                   respjson.Field
		CreatorWorkTelephone               respjson.Field
		CreatorWorkURL                     respjson.Field
		Credit                             respjson.Field
		DateCreated                        respjson.Field
		DateTimeCreated                    respjson.Field
		DateTimeOriginal                   respjson.Field
		Description                        respjson.Field
		DigitalImageGuid                   respjson.Field
		DigitalSourceType                  respjson.Field
		EmbeddedEncodedRightsExpr          respjson.Field
		EmbeddedEncodedRightsExprLangID    respjson.Field
		EmbeddedEncodedRightsExprType      respjson.Field
		Event                              respjson.Field
		ExifVersion                        respjson.Field
		FlashpixVersion                    respjson.Field
		GenreCvID                          respjson.Field
		GenreCvTermID                      respjson.Field
		GenreCvTermName                    respjson.Field
		GenreCvTermRefinedAbout            respjson.Field
		Headline                           respjson.Field
		ImageCreatorID                     respjson.Field
		ImageCreatorImageID                respjson.Field
		ImageCreatorName                   respjson.Field
		ImageDescription                   respjson.Field
		ImageRegionBoundaryH               respjson.Field
		ImageRegionBoundaryRx              respjson.Field
		ImageRegionBoundaryShape           respjson.Field
		ImageRegionBoundaryUnit            respjson.Field
		ImageRegionBoundaryVerticesX       respjson.Field
		ImageRegionBoundaryVerticesY       respjson.Field
		ImageRegionBoundaryW               respjson.Field
		ImageRegionBoundaryX               respjson.Field
		ImageRegionBoundaryY               respjson.Field
		ImageRegionCtypeIdentifier         respjson.Field
		ImageRegionCtypeName               respjson.Field
		ImageRegionID                      respjson.Field
		ImageRegionName                    respjson.Field
		ImageRegionOrganisationInImageName respjson.Field
		ImageRegionPersonInImage           respjson.Field
		ImageRegionRoleIdentifier          respjson.Field
		ImageRegionRoleName                respjson.Field
		ImageSupplierID                    respjson.Field
		ImageSupplierImageID               respjson.Field
		ImageSupplierName                  respjson.Field
		Instructions                       respjson.Field
		IntellectualGenre                  respjson.Field
		Keywords                           respjson.Field
		LicensorCity                       respjson.Field
		LicensorCountry                    respjson.Field
		LicensorEmail                      respjson.Field
		LicensorExtendedAddress            respjson.Field
		LicensorID                         respjson.Field
		LicensorName                       respjson.Field
		LicensorPostalCode                 respjson.Field
		LicensorRegion                     respjson.Field
		LicensorStreetAddress              respjson.Field
		LicensorTelephone1                 respjson.Field
		LicensorTelephone2                 respjson.Field
		LicensorURL                        respjson.Field
		LinkedEncodedRightsExpr            respjson.Field
		LinkedEncodedRightsExprLangID      respjson.Field
		LinkedEncodedRightsExprType        respjson.Field
		Location                           respjson.Field
		LocationCreatedCity                respjson.Field
		LocationCreatedCountryCode         respjson.Field
		LocationCreatedCountryName         respjson.Field
		LocationCreatedGpsAltitude         respjson.Field
		LocationCreatedGpsLatitude         respjson.Field
		LocationCreatedGpsLongitude        respjson.Field
		LocationCreatedLocationID          respjson.Field
		LocationCreatedLocationName        respjson.Field
		LocationCreatedProvinceState       respjson.Field
		LocationCreatedSublocation         respjson.Field
		LocationCreatedWorldRegion         respjson.Field
		LocationShownCity                  respjson.Field
		LocationShownCountryCode           respjson.Field
		LocationShownCountryName           respjson.Field
		LocationShownGpsAltitude           respjson.Field
		LocationShownGpsLatitude           respjson.Field
		LocationShownGpsLongitude          respjson.Field
		LocationShownLocationID            respjson.Field
		LocationShownLocationName          respjson.Field
		LocationShownProvinceState         respjson.Field
		LocationShownSublocation           respjson.Field
		LocationShownWorldRegion           respjson.Field
		MaxAvailHeight                     respjson.Field
		MaxAvailWidth                      respjson.Field
		ModelAge                           respjson.Field
		ModelReleaseID                     respjson.Field
		ObjectAttributeReference           respjson.Field
		ObjectName                         respjson.Field
		OffsetTimeOriginal                 respjson.Field
		OrganisationInImageCode            respjson.Field
		OrganisationInImageName            respjson.Field
		Orientation                        respjson.Field
		OriginalTransmissionReference      respjson.Field
		PersonInImage                      respjson.Field
		PersonInImageCvTermCvID            respjson.Field
		PersonInImageCvTermID              respjson.Field
		PersonInImageCvTermName            respjson.Field
		PersonInImageCvTermRefinedAbout    respjson.Field
		PersonInImageDescription           respjson.Field
		PersonInImageID                    respjson.Field
		PersonInImageName                  respjson.Field
		ProductInImageDescription          respjson.Field
		ProductInImageGtin                 respjson.Field
		ProductInImageName                 respjson.Field
		PropertyReleaseID                  respjson.Field
		ProvinceState                      respjson.Field
		Rating                             respjson.Field
		RegistryEntryRole                  respjson.Field
		RegistryItemID                     respjson.Field
		RegistryOrganisationID             respjson.Field
		ResolutionUnit                     respjson.Field
		Rights                             respjson.Field
		Scene                              respjson.Field
		Source                             respjson.Field
		SpecialInstructions                respjson.Field
		State                              respjson.Field
		Subject                            respjson.Field
		SubjectCode                        respjson.Field
		SubjectReference                   respjson.Field
		Sublocation                        respjson.Field
		TimeCreated                        respjson.Field
		Title                              respjson.Field
		TransmissionReference              respjson.Field
		UsageTerms                         respjson.Field
		WebStatement                       respjson.Field
		Writer                             respjson.Field
		WriterEditor                       respjson.Field
		XResolution                        respjson.Field
		YResolution                        respjson.Field
		ExtraFields                        map[string]respjson.Field
		raw                                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileUploadV2ResponseEmbeddedMetadata) RawJSON() string { return r.JSON.raw }
func (r *FileUploadV2ResponseEmbeddedMetadata) UnmarshalJSON(data []byte) error {
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
type FileUploadV2ResponseExtensionStatus struct {
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
func (r FileUploadV2ResponseExtensionStatus) RawJSON() string { return r.JSON.raw }
func (r *FileUploadV2ResponseExtensionStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Legacy metadata. Send `metadata` in `responseFields` in API request to get
// metadata in the upload API response.
type FileUploadV2ResponseMetadata struct {
	// The audio codec used in the video (only for video).
	AudioCodec string `json:"audioCodec"`
	// The bit rate of the video in kbps (only for video).
	BitRate int64 `json:"bitRate"`
	// The density of the image in DPI.
	Density int64 `json:"density"`
	// The duration of the video in seconds (only for video).
	Duration int64                            `json:"duration"`
	Exif     FileUploadV2ResponseMetadataExif `json:"exif"`
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
func (r FileUploadV2ResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *FileUploadV2ResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileUploadV2ResponseMetadataExif struct {
	// Object containing Exif details.
	Exif ExifDetails `json:"exif"`
	// Object containing GPS information.
	Gps Gps `json:"gps"`
	// Object containing EXIF image information.
	Image ExifImage `json:"image"`
	// JSON object.
	Interoperability Interoperability `json:"interoperability"`
	Makernote        map[string]any   `json:"makernote"`
	// Object containing Thumbnail information.
	Thumbnail Thumbnail `json:"thumbnail"`
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
func (r FileUploadV2ResponseMetadataExif) RawJSON() string { return r.JSON.raw }
func (r *FileUploadV2ResponseMetadataExif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object containing the file or file version's `id` (versionId) and `name`.
type FileUploadV2ResponseVersionInfo struct {
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
func (r FileUploadV2ResponseVersionInfo) RawJSON() string { return r.JSON.raw }
func (r *FileUploadV2ResponseVersionInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileListParams struct {
	// Type of files to include in the result set. Accepts three values:
	//
	// `all` - include all types of files in the result set. `image` - only search in
	// image type files. `non-image` - only search in files that are not images, e.g.,
	// JS or CSS or video files.
	//
	// Default value - `all`
	FileType param.Opt[string] `query:"fileType,omitzero" json:"-"`
	// The maximum number of results to return in response:
	//
	// # Minimum value - 1
	//
	// # Maximum value - 1000
	//
	// Default value - 1000
	Limit param.Opt[string] `query:"limit,omitzero" json:"-"`
	// Folder path if you want to limit the search within a specific folder. For
	// example, `/sales-banner/` will only search in folder sales-banner.
	Path param.Opt[string] `query:"path,omitzero" json:"-"`
	// Query string in a Lucene-like query language e.g. `createdAt > "7d"`.
	//
	// Note : When the searchQuery parameter is present, the following query parameters
	// will have no effect on the result:
	//
	// 1. `tags`
	// 2. `type`
	// 3. `name`
	//
	// [Learn more](/docs/api-reference/digital-asset-management-dam/list-and-search-assets#advanced-search-queries)
	// from examples.
	SearchQuery param.Opt[string] `query:"searchQuery,omitzero" json:"-"`
	// The number of results to skip before returning results:
	//
	// # Minimum value - 0
	//
	// Default value - 0
	Skip param.Opt[string] `query:"skip,omitzero" json:"-"`
	// You can sort based on the following fields:
	//
	// 1. name - `ASC_NAME` or `DESC_NAME`
	// 2. createdAt - `ASC_CREATED` or `DESC_CREATED`
	// 3. updatedAt - `ASC_UPDATED` or `DESC_UPDATED`
	// 4. height - `ASC_HEIGHT` or `DESC_HEIGHT`
	// 5. width - `ASC_WIDTH` or `DESC_WIDTH`
	// 6. size - `ASC_SIZE` or `DESC_SIZE`
	//
	// Default value - `ASC_CREATED`
	Sort param.Opt[string] `query:"sort,omitzero" json:"-"`
	// Limit search to one of `file`, `file-version`, or `folder`. Pass `all` to
	// include `files` and `folders` in search results (`file-version` will not be
	// included in this case).
	//
	// Default value - `file`
	//
	// Any of "file", "file-version", "folder", "all".
	Type FileListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FileListParams]'s query parameters as `url.Values`.
func (r FileListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Limit search to one of `file`, `file-version`, or `folder`. Pass `all` to
// include `files` and `folders` in search results (`file-version` will not be
// included in this case).
//
// Default value - `file`
type FileListParamsType string

const (
	FileListParamsTypeFile        FileListParamsType = "file"
	FileListParamsTypeFileVersion FileListParamsType = "file-version"
	FileListParamsTypeFolder      FileListParamsType = "folder"
	FileListParamsTypeAll         FileListParamsType = "all"
)

type FileAddTagsParams struct {
	// An array of fileIds to which you want to add tags.
	FileIDs []string `json:"fileIds,omitzero,required"`
	// An array of tags that you want to add to the files.
	Tags []string `json:"tags,omitzero,required"`
	paramObj
}

func (r FileAddTagsParams) MarshalJSON() (data []byte, err error) {
	type shadow FileAddTagsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileAddTagsParams) UnmarshalJSON(data []byte) error {
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

type FileRemoveAITagsParams struct {
	// An array of AITags that you want to remove from the files.
	AITags []string `json:"AITags,omitzero,required"`
	// An array of fileIds from which you want to remove AITags.
	FileIDs []string `json:"fileIds,omitzero,required"`
	paramObj
}

func (r FileRemoveAITagsParams) MarshalJSON() (data []byte, err error) {
	type shadow FileRemoveAITagsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileRemoveAITagsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileRemoveTagsParams struct {
	// An array of fileIds from which you want to remove tags.
	FileIDs []string `json:"fileIds,omitzero,required"`
	// An array of tags that you want to remove from the files.
	Tags []string `json:"tags,omitzero,required"`
	paramObj
}

func (r FileRemoveTagsParams) MarshalJSON() (data []byte, err error) {
	type shadow FileRemoveTagsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileRemoveTagsParams) UnmarshalJSON(data []byte) error {
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

type FileUploadV1Params struct {
	// Pass the HTTP URL or base64 string. When passing a URL in the file parameter,
	// please ensure that our servers can access the URL. In case ImageKit is unable to
	// download the file from the specified URL, a `400` error response is returned.
	// This will also result in a `400` error if the file download request is aborted
	// if response headers are not received in 8 seconds.
	File string `json:"file,required"`
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
	// Stringified JSON key-value data to be associated with the asset.
	CustomMetadata param.Opt[string] `json:"customMetadata,omitzero"`
	// The time until your signature is valid. It must be a
	// [Unix time](https://en.wikipedia.org/wiki/Unix_time) in less than 1 hour into
	// the future. It should be in seconds. This field is only required for
	// authentication when uploading a file from the client side.
	Expire param.Opt[string] `json:"expire,omitzero"`
	// Stringified JSON object with an array of extensions to be applied to the image.
	// Refer to extensions schema in
	// [update file API request body](/docs/api-reference/digital-asset-management-dam/managing-assets/update-file-details#request-body).
	Extensions param.Opt[string] `json:"extensions,omitzero"`
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
	// If `false` and `useUniqueFileName` is also `false`, and a file already exists at
	// the exact location, upload API will return an error immediately.
	OverwriteFile param.Opt[string] `json:"overwriteFile,omitzero"`
	// Your ImageKit.io public key. This field is only required for authentication when
	// uploading a file from the client side.
	PublicKey param.Opt[string] `json:"publicKey,omitzero"`
	// Comma-separated values of the fields that you want the API to return in the
	// response.
	//
	// For example, set the value of this field to
	// `tags,customCoordinates,isPrivateFile` to get the value of `tags`,
	// `customCoordinates`, and `isPrivateFile` in the response.
	//
	// Accepts combination of `tags`, `customCoordinates`, `isPrivateFile`,
	// `embeddedMetadata`, `isPublished`, `customMetadata`, and `metadata`.
	ResponseFields param.Opt[string] `json:"responseFields,omitzero"`
	// HMAC-SHA1 digest of the token+expire using your ImageKit.io private API key as a
	// key. Learn how to create a signature on the page below. This should be in
	// lowercase.
	//
	// Signature must be calculated on the server-side. This field is only required for
	// authentication when uploading a file from the client side.
	Signature param.Opt[string] `json:"signature,omitzero"`
	// Set the tags while uploading the file.
	//
	// Comma-separated value of tags in the format `tag1,tag2,tag3`. The maximum length
	// of all characters should not exceed 500. `%` is not allowed.
	//
	// If this field is not specified and the file is overwritten then the tags will be
	// removed.
	Tags param.Opt[string] `json:"tags,omitzero"`
	// Stringified JSON object with properties for pre and post transformations:
	//
	// `pre` - Accepts a "string" containing a valid transformation used for requesting
	// a pre-transformation for an image or a video file.
	//
	// `post` - Accepts an array of objects with properties:
	//
	//   - `type`: One of `transformation`, `gif-to-video`, `thumbnail`, or `abs`
	//     (Adaptive bitrate streaming).
	//   - `value`: A "string" corresponding to the required transformation. Required if
	//     `type` is `transformation` or `abs`. Optional if `type` is `gif-to-video` or
	//     `thumbnail`.
	//   - `protocol`: Either `hls` or `dash`, applicable only if `type` is `abs`.
	//
	// Read more about
	// [Adaptive bitrate streaming (ABS)](/docs/adaptive-bitrate-streaming).
	Transformation param.Opt[string] `json:"transformation,omitzero"`
	// The final status of extensions after they have completed execution will be
	// delivered to this endpoint as a POST request.
	// [Learn more](/docs/api-reference/digital-asset-management-dam/managing-assets/update-file-details#webhook-payload-structure)
	// about the webhook payload structure.
	WebhookURL param.Opt[string] `json:"webhookUrl,omitzero"`
	// Whether to mark the file as private or not.
	//
	// If `true`, the file is marked as private and is accessible only using named
	// transformation or signed URL.
	//
	// Any of "true", "false".
	IsPrivateFile FileUploadV1ParamsIsPrivateFile `json:"isPrivateFile,omitzero"`
	// Whether to upload file as published or not.
	//
	// If `false`, the file is marked as unpublished, which restricts access to the
	// file only via the media library. Files in draft or unpublished state can only be
	// publicly accessed after being published.
	//
	// The option to upload in draft state is only available in custom enterprise
	// pricing plans.
	//
	// Any of "true", "false".
	IsPublished FileUploadV1ParamsIsPublished `json:"isPublished,omitzero"`
	// If set to `true` and a file already exists at the exact location, its AITags
	// will be removed. Set `overwriteAITags` to `false` to preserve AITags.
	//
	// Any of "true", "false".
	OverwriteAITags FileUploadV1ParamsOverwriteAITags `json:"overwriteAITags,omitzero"`
	// If the request does not have `customMetadata`, and a file already exists at the
	// exact location, existing customMetadata will be removed.
	//
	// Any of "true", "false".
	OverwriteCustomMetadata FileUploadV1ParamsOverwriteCustomMetadata `json:"overwriteCustomMetadata,omitzero"`
	// If the request does not have `tags`, and a file already exists at the exact
	// location, existing tags will be removed.
	//
	// Any of "true", "false".
	OverwriteTags FileUploadV1ParamsOverwriteTags `json:"overwriteTags,omitzero"`
	// Whether to use a unique filename for this file or not.
	//
	// If `true`, ImageKit.io will add a unique suffix to the filename parameter to get
	// a unique filename.
	//
	// If `false`, then the image is uploaded with the provided filename parameter, and
	// any existing file with the same name is replaced.
	//
	// Any of "true", "false".
	UseUniqueFileName FileUploadV1ParamsUseUniqueFileName `json:"useUniqueFileName,omitzero"`
	paramObj
}

func (r FileUploadV1Params) MarshalMultipart() (data []byte, contentType string, err error) {
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

// Whether to mark the file as private or not.
//
// If `true`, the file is marked as private and is accessible only using named
// transformation or signed URL.
type FileUploadV1ParamsIsPrivateFile string

const (
	FileUploadV1ParamsIsPrivateFileTrue  FileUploadV1ParamsIsPrivateFile = "true"
	FileUploadV1ParamsIsPrivateFileFalse FileUploadV1ParamsIsPrivateFile = "false"
)

// Whether to upload file as published or not.
//
// If `false`, the file is marked as unpublished, which restricts access to the
// file only via the media library. Files in draft or unpublished state can only be
// publicly accessed after being published.
//
// The option to upload in draft state is only available in custom enterprise
// pricing plans.
type FileUploadV1ParamsIsPublished string

const (
	FileUploadV1ParamsIsPublishedTrue  FileUploadV1ParamsIsPublished = "true"
	FileUploadV1ParamsIsPublishedFalse FileUploadV1ParamsIsPublished = "false"
)

// If set to `true` and a file already exists at the exact location, its AITags
// will be removed. Set `overwriteAITags` to `false` to preserve AITags.
type FileUploadV1ParamsOverwriteAITags string

const (
	FileUploadV1ParamsOverwriteAITagsTrue  FileUploadV1ParamsOverwriteAITags = "true"
	FileUploadV1ParamsOverwriteAITagsFalse FileUploadV1ParamsOverwriteAITags = "false"
)

// If the request does not have `customMetadata`, and a file already exists at the
// exact location, existing customMetadata will be removed.
type FileUploadV1ParamsOverwriteCustomMetadata string

const (
	FileUploadV1ParamsOverwriteCustomMetadataTrue  FileUploadV1ParamsOverwriteCustomMetadata = "true"
	FileUploadV1ParamsOverwriteCustomMetadataFalse FileUploadV1ParamsOverwriteCustomMetadata = "false"
)

// If the request does not have `tags`, and a file already exists at the exact
// location, existing tags will be removed.
type FileUploadV1ParamsOverwriteTags string

const (
	FileUploadV1ParamsOverwriteTagsTrue  FileUploadV1ParamsOverwriteTags = "true"
	FileUploadV1ParamsOverwriteTagsFalse FileUploadV1ParamsOverwriteTags = "false"
)

// Whether to use a unique filename for this file or not.
//
// If `true`, ImageKit.io will add a unique suffix to the filename parameter to get
// a unique filename.
//
// If `false`, then the image is uploaded with the provided filename parameter, and
// any existing file with the same name is replaced.
type FileUploadV1ParamsUseUniqueFileName string

const (
	FileUploadV1ParamsUseUniqueFileNameTrue  FileUploadV1ParamsUseUniqueFileName = "true"
	FileUploadV1ParamsUseUniqueFileNameFalse FileUploadV1ParamsUseUniqueFileName = "false"
)

type FileUploadV2Params struct {
	// Pass the HTTP URL or base64 string. When passing a URL in the file parameter,
	// please ensure that our servers can access the URL. In case ImageKit is unable to
	// download the file from the specified URL, a `400` error response is returned.
	// This will also result in a `400` error if the file download request is aborted
	// if response headers are not received in 8 seconds.
	File string `json:"file,required"`
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
	// Stringified JSON key-value data to be associated with the asset.
	CustomMetadata param.Opt[string] `json:"customMetadata,omitzero"`
	// Stringified JSON object with an array of extensions to be applied to the image.
	// Refer to extensions schema in
	// [update file API request body](/docs/api-reference/digital-asset-management-dam/managing-assets/update-file-details#request-body).
	Extensions param.Opt[string] `json:"extensions,omitzero"`
	// The folder path in which the image has to be uploaded. If the folder(s) didn't
	// exist before, a new folder(s) is created. Using multiple `/` creates a nested
	// folder.
	Folder param.Opt[string] `json:"folder,omitzero"`
	// If `false` and `useUniqueFileName` is also `false`, and a file already exists at
	// the exact location, upload API will return an error immediately.
	OverwriteFile param.Opt[string] `json:"overwriteFile,omitzero"`
	// Comma-separated values of the fields that you want the API to return in the
	// response.
	//
	// For example, set the value of this field to
	// `tags,customCoordinates,isPrivateFile` to get the value of `tags`,
	// `customCoordinates`, and `isPrivateFile` in the response.
	//
	// Accepts combination of `tags`, `customCoordinates`, `isPrivateFile`,
	// `embeddedMetadata`, `isPublished`, `customMetadata`, and `metadata`.
	ResponseFields param.Opt[string] `json:"responseFields,omitzero"`
	// Set the tags while uploading the file.
	//
	// Comma-separated value of tags in the format `tag1,tag2,tag3`. The maximum length
	// of all characters should not exceed 500. `%` is not allowed.
	//
	// If this field is not specified and the file is overwritten then the tags will be
	// removed.
	Tags param.Opt[string] `json:"tags,omitzero"`
	// Stringified JSON object with properties for pre and post transformations:
	//
	// `pre` - Accepts a "string" containing a valid transformation used for requesting
	// a pre-transformation for an image or a video file.
	//
	// `post` - Accepts an array of objects with properties:
	//
	//   - `type`: One of `transformation`, `gif-to-video`, `thumbnail`, or `abs`
	//     (Adaptive bitrate streaming).
	//   - `value`: A "string" corresponding to the required transformation. Required if
	//     `type` is `transformation` or `abs`. Optional if `type` is `gif-to-video` or
	//     `thumbnail`.
	//   - `protocol`: Either `hls` or `dash`, applicable only if `type` is `abs`.
	//
	// Read more about
	// [Adaptive bitrate streaming (ABS)](/docs/adaptive-bitrate-streaming).
	Transformation param.Opt[string] `json:"transformation,omitzero"`
	// The final status of extensions after they have completed execution will be
	// delivered to this endpoint as a POST request.
	// [Learn more](/docs/api-reference/digital-asset-management-dam/managing-assets/update-file-details#webhook-payload-structure)
	// about the webhook payload structure.
	WebhookURL param.Opt[string] `json:"webhookUrl,omitzero"`
	// Whether to mark the file as private or not.
	//
	// If `true`, the file is marked as private and is accessible only using named
	// transformation or signed URL.
	//
	// Any of "true", "false".
	IsPrivateFile FileUploadV2ParamsIsPrivateFile `json:"isPrivateFile,omitzero"`
	// Whether to upload file as published or not.
	//
	// If `false`, the file is marked as unpublished, which restricts access to the
	// file only via the media library. Files in draft or unpublished state can only be
	// publicly accessed after being published.
	//
	// The option to upload in draft state is only available in custom enterprise
	// pricing plans.
	//
	// Any of "true", "false".
	IsPublished FileUploadV2ParamsIsPublished `json:"isPublished,omitzero"`
	// If set to `true` and a file already exists at the exact location, its AITags
	// will be removed. Set `overwriteAITags` to `false` to preserve AITags.
	//
	// Any of "true", "false".
	OverwriteAITags FileUploadV2ParamsOverwriteAITags `json:"overwriteAITags,omitzero"`
	// If the request does not have `customMetadata`, and a file already exists at the
	// exact location, existing customMetadata will be removed.
	//
	// Any of "true", "false".
	OverwriteCustomMetadata FileUploadV2ParamsOverwriteCustomMetadata `json:"overwriteCustomMetadata,omitzero"`
	// If the request does not have `tags`, and a file already exists at the exact
	// location, existing tags will be removed.
	//
	// Any of "true", "false".
	OverwriteTags FileUploadV2ParamsOverwriteTags `json:"overwriteTags,omitzero"`
	// Whether to use a unique filename for this file or not.
	//
	// If `true`, ImageKit.io will add a unique suffix to the filename parameter to get
	// a unique filename.
	//
	// If `false`, then the image is uploaded with the provided filename parameter, and
	// any existing file with the same name is replaced.
	//
	// Any of "true", "false".
	UseUniqueFileName FileUploadV2ParamsUseUniqueFileName `json:"useUniqueFileName,omitzero"`
	paramObj
}

func (r FileUploadV2Params) MarshalMultipart() (data []byte, contentType string, err error) {
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

// Whether to mark the file as private or not.
//
// If `true`, the file is marked as private and is accessible only using named
// transformation or signed URL.
type FileUploadV2ParamsIsPrivateFile string

const (
	FileUploadV2ParamsIsPrivateFileTrue  FileUploadV2ParamsIsPrivateFile = "true"
	FileUploadV2ParamsIsPrivateFileFalse FileUploadV2ParamsIsPrivateFile = "false"
)

// Whether to upload file as published or not.
//
// If `false`, the file is marked as unpublished, which restricts access to the
// file only via the media library. Files in draft or unpublished state can only be
// publicly accessed after being published.
//
// The option to upload in draft state is only available in custom enterprise
// pricing plans.
type FileUploadV2ParamsIsPublished string

const (
	FileUploadV2ParamsIsPublishedTrue  FileUploadV2ParamsIsPublished = "true"
	FileUploadV2ParamsIsPublishedFalse FileUploadV2ParamsIsPublished = "false"
)

// If set to `true` and a file already exists at the exact location, its AITags
// will be removed. Set `overwriteAITags` to `false` to preserve AITags.
type FileUploadV2ParamsOverwriteAITags string

const (
	FileUploadV2ParamsOverwriteAITagsTrue  FileUploadV2ParamsOverwriteAITags = "true"
	FileUploadV2ParamsOverwriteAITagsFalse FileUploadV2ParamsOverwriteAITags = "false"
)

// If the request does not have `customMetadata`, and a file already exists at the
// exact location, existing customMetadata will be removed.
type FileUploadV2ParamsOverwriteCustomMetadata string

const (
	FileUploadV2ParamsOverwriteCustomMetadataTrue  FileUploadV2ParamsOverwriteCustomMetadata = "true"
	FileUploadV2ParamsOverwriteCustomMetadataFalse FileUploadV2ParamsOverwriteCustomMetadata = "false"
)

// If the request does not have `tags`, and a file already exists at the exact
// location, existing tags will be removed.
type FileUploadV2ParamsOverwriteTags string

const (
	FileUploadV2ParamsOverwriteTagsTrue  FileUploadV2ParamsOverwriteTags = "true"
	FileUploadV2ParamsOverwriteTagsFalse FileUploadV2ParamsOverwriteTags = "false"
)

// Whether to use a unique filename for this file or not.
//
// If `true`, ImageKit.io will add a unique suffix to the filename parameter to get
// a unique filename.
//
// If `false`, then the image is uploaded with the provided filename parameter, and
// any existing file with the same name is replaced.
type FileUploadV2ParamsUseUniqueFileName string

const (
	FileUploadV2ParamsUseUniqueFileNameTrue  FileUploadV2ParamsUseUniqueFileName = "true"
	FileUploadV2ParamsUseUniqueFileNameFalse FileUploadV2ParamsUseUniqueFileName = "false"
)
