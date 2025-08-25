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
	"github.com/stainless-sdks/imagekit-go/shared/constant"
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
func (r *FileService) Get(ctx context.Context, fileID string, opts ...option.RequestOption) (res *shared.File, err error) {
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

// JSON object containing metadata.
type Metadata struct {
	// The audio codec used in the video (only for video).
	AudioCodec string `json:"audioCodec"`
	// The bit rate of the video in kbps (only for video).
	BitRate int64 `json:"bitRate"`
	// The density of the image in DPI.
	Density int64 `json:"density"`
	// The duration of the video in seconds (only for video).
	Duration int64        `json:"duration"`
	Exif     MetadataExif `json:"exif"`
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
func (r Metadata) RawJSON() string { return r.JSON.raw }
func (r *Metadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MetadataExif struct {
	// Object containing Exif details.
	Exif MetadataExifExif `json:"exif"`
	// Object containing GPS information.
	Gps MetadataExifGps `json:"gps"`
	// Object containing EXIF image information.
	Image MetadataExifImage `json:"image"`
	// JSON object.
	Interoperability MetadataExifInteroperability `json:"interoperability"`
	Makernote        map[string]any               `json:"makernote"`
	// Object containing Thumbnail information.
	Thumbnail MetadataExifThumbnail `json:"thumbnail"`
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
func (r MetadataExif) RawJSON() string { return r.JSON.raw }
func (r *MetadataExif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Object containing Exif details.
type MetadataExifExif struct {
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
func (r MetadataExifExif) RawJSON() string { return r.JSON.raw }
func (r *MetadataExifExif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Object containing GPS information.
type MetadataExifGps struct {
	GpsVersionID []int64 `json:"GPSVersionID"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GpsVersionID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetadataExifGps) RawJSON() string { return r.JSON.raw }
func (r *MetadataExifGps) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Object containing EXIF image information.
type MetadataExifImage struct {
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
func (r MetadataExifImage) RawJSON() string { return r.JSON.raw }
func (r *MetadataExifImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// JSON object.
type MetadataExifInteroperability struct {
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
func (r MetadataExifInteroperability) RawJSON() string { return r.JSON.raw }
func (r *MetadataExifInteroperability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Object containing Thumbnail information.
type MetadataExifThumbnail struct {
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
func (r MetadataExifThumbnail) RawJSON() string { return r.JSON.raw }
func (r *MetadataExifThumbnail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Object containing details of a file or file version.
type FileUpdateResponse struct {
	ExtensionStatus FileUpdateResponseExtensionStatus `json:"extensionStatus"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtensionStatus respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
	shared.File
}

// Returns the unmodified JSON received from the API
func (r FileUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *FileUpdateResponse) UnmarshalJSON(data []byte) error {
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

type FileCopyResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileCopyResponse) RawJSON() string { return r.JSON.raw }
func (r *FileCopyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileMoveResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileMoveResponse) RawJSON() string { return r.JSON.raw }
func (r *FileMoveResponse) UnmarshalJSON(data []byte) error {
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
	OfUpdateFileDetails *FileUpdateParamsUpdateUpdateFileDetails `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfChangePublicationStatus *FileUpdateParamsUpdateChangePublicationStatus `json:",inline"`

	paramObj
}

func (u FileUpdateParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfUpdateFileDetails, u.OfChangePublicationStatus)
}
func (r *FileUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileUpdateParamsUpdateUpdateFileDetails struct {
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
	CustomMetadata map[string]any `json:"customMetadata,omitzero"`
	// Array of extensions to be applied to the asset. Each extension can be configured
	// with specific parameters based on the extension type.
	Extensions []FileUpdateParamsUpdateUpdateFileDetailsExtensionUnion `json:"extensions,omitzero"`
	// An array of AITags associated with the file that you want to remove, e.g.
	// `["car", "vehicle", "motorsports"]`.
	//
	// If you want to remove all AITags associated with the file, send a string -
	// "all".
	//
	// Note: The remove operation for `AITags` executes before any of the `extensions`
	// are processed.
	RemoveAITags FileUpdateParamsUpdateUpdateFileDetailsRemoveAITagsUnion `json:"removeAITags,omitzero"`
	paramObj
}

func (r FileUpdateParamsUpdateUpdateFileDetails) MarshalJSON() (data []byte, err error) {
	type shadow FileUpdateParamsUpdateUpdateFileDetails
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUpdateParamsUpdateUpdateFileDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FileUpdateParamsUpdateUpdateFileDetailsExtensionUnion struct {
	OfRemoveBg          *FileUpdateParamsUpdateUpdateFileDetailsExtensionRemoveBg          `json:",omitzero,inline"`
	OfAutoTagging       *FileUpdateParamsUpdateUpdateFileDetailsExtensionAutoTagging       `json:",omitzero,inline"`
	OfAIAutoDescription *FileUpdateParamsUpdateUpdateFileDetailsExtensionAIAutoDescription `json:",omitzero,inline"`
	paramUnion
}

func (u FileUpdateParamsUpdateUpdateFileDetailsExtensionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRemoveBg, u.OfAutoTagging, u.OfAIAutoDescription)
}
func (u *FileUpdateParamsUpdateUpdateFileDetailsExtensionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FileUpdateParamsUpdateUpdateFileDetailsExtensionUnion) asAny() any {
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
func (u FileUpdateParamsUpdateUpdateFileDetailsExtensionUnion) GetOptions() *FileUpdateParamsUpdateUpdateFileDetailsExtensionRemoveBgOptions {
	if vt := u.OfRemoveBg; vt != nil {
		return &vt.Options
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FileUpdateParamsUpdateUpdateFileDetailsExtensionUnion) GetMaxTags() *int64 {
	if vt := u.OfAutoTagging; vt != nil {
		return &vt.MaxTags
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FileUpdateParamsUpdateUpdateFileDetailsExtensionUnion) GetMinConfidence() *int64 {
	if vt := u.OfAutoTagging; vt != nil {
		return &vt.MinConfidence
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FileUpdateParamsUpdateUpdateFileDetailsExtensionUnion) GetName() *string {
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
	apijson.RegisterUnion[FileUpdateParamsUpdateUpdateFileDetailsExtensionUnion](
		"name",
		apijson.Discriminator[FileUpdateParamsUpdateUpdateFileDetailsExtensionRemoveBg]("remove-bg"),
		apijson.Discriminator[FileUpdateParamsUpdateUpdateFileDetailsExtensionAutoTagging]("google-auto-tagging"),
		apijson.Discriminator[FileUpdateParamsUpdateUpdateFileDetailsExtensionAutoTagging]("aws-auto-tagging"),
		apijson.Discriminator[FileUpdateParamsUpdateUpdateFileDetailsExtensionAIAutoDescription]("ai-auto-description"),
	)
}

// The property Name is required.
type FileUpdateParamsUpdateUpdateFileDetailsExtensionRemoveBg struct {
	Options FileUpdateParamsUpdateUpdateFileDetailsExtensionRemoveBgOptions `json:"options,omitzero"`
	// Specifies the background removal extension.
	//
	// This field can be elided, and will marshal its zero value as "remove-bg".
	Name constant.RemoveBg `json:"name,required"`
	paramObj
}

func (r FileUpdateParamsUpdateUpdateFileDetailsExtensionRemoveBg) MarshalJSON() (data []byte, err error) {
	type shadow FileUpdateParamsUpdateUpdateFileDetailsExtensionRemoveBg
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUpdateParamsUpdateUpdateFileDetailsExtensionRemoveBg) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileUpdateParamsUpdateUpdateFileDetailsExtensionRemoveBgOptions struct {
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

func (r FileUpdateParamsUpdateUpdateFileDetailsExtensionRemoveBgOptions) MarshalJSON() (data []byte, err error) {
	type shadow FileUpdateParamsUpdateUpdateFileDetailsExtensionRemoveBgOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUpdateParamsUpdateUpdateFileDetailsExtensionRemoveBgOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties MaxTags, MinConfidence, Name are required.
type FileUpdateParamsUpdateUpdateFileDetailsExtensionAutoTagging struct {
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

func (r FileUpdateParamsUpdateUpdateFileDetailsExtensionAutoTagging) MarshalJSON() (data []byte, err error) {
	type shadow FileUpdateParamsUpdateUpdateFileDetailsExtensionAutoTagging
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUpdateParamsUpdateUpdateFileDetailsExtensionAutoTagging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FileUpdateParamsUpdateUpdateFileDetailsExtensionAutoTagging](
		"name", "google-auto-tagging", "aws-auto-tagging",
	)
}

func NewFileUpdateParamsUpdateUpdateFileDetailsExtensionAIAutoDescription() FileUpdateParamsUpdateUpdateFileDetailsExtensionAIAutoDescription {
	return FileUpdateParamsUpdateUpdateFileDetailsExtensionAIAutoDescription{
		Name: "ai-auto-description",
	}
}

// This struct has a constant value, construct it with
// [NewFileUpdateParamsUpdateUpdateFileDetailsExtensionAIAutoDescription].
type FileUpdateParamsUpdateUpdateFileDetailsExtensionAIAutoDescription struct {
	// Specifies the auto description extension.
	Name constant.AIAutoDescription `json:"name,required"`
	paramObj
}

func (r FileUpdateParamsUpdateUpdateFileDetailsExtensionAIAutoDescription) MarshalJSON() (data []byte, err error) {
	type shadow FileUpdateParamsUpdateUpdateFileDetailsExtensionAIAutoDescription
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUpdateParamsUpdateUpdateFileDetailsExtensionAIAutoDescription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type FileUpdateParamsUpdateUpdateFileDetailsRemoveAITagsUnion struct {
	OfStringArray []string `json:",omitzero,inline"`
	// Construct this variant with constant.ValueOf[constant.All]()
	OfAll constant.All `json:",omitzero,inline"`
	paramUnion
}

func (u FileUpdateParamsUpdateUpdateFileDetailsRemoveAITagsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfStringArray, u.OfAll)
}
func (u *FileUpdateParamsUpdateUpdateFileDetailsRemoveAITagsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FileUpdateParamsUpdateUpdateFileDetailsRemoveAITagsUnion) asAny() any {
	if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	} else if !param.IsOmitted(u.OfAll) {
		return &u.OfAll
	}
	return nil
}

type FileUpdateParamsUpdateChangePublicationStatus struct {
	// Configure the publication status of a file and its versions.
	Publish FileUpdateParamsUpdateChangePublicationStatusPublish `json:"publish,omitzero"`
	paramObj
}

func (r FileUpdateParamsUpdateChangePublicationStatus) MarshalJSON() (data []byte, err error) {
	type shadow FileUpdateParamsUpdateChangePublicationStatus
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUpdateParamsUpdateChangePublicationStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configure the publication status of a file and its versions.
//
// The property IsPublished is required.
type FileUpdateParamsUpdateChangePublicationStatusPublish struct {
	// Set to `true` to publish the file. Set to `false` to unpublish the file.
	IsPublished bool `json:"isPublished,required"`
	// Set to `true` to publish/unpublish all versions of the file. Set to `false` to
	// publish/unpublish only the current version of the file.
	IncludeFileVersions param.Opt[bool] `json:"includeFileVersions,omitzero"`
	paramObj
}

func (r FileUpdateParamsUpdateChangePublicationStatusPublish) MarshalJSON() (data []byte, err error) {
	type shadow FileUpdateParamsUpdateChangePublicationStatusPublish
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUpdateParamsUpdateChangePublicationStatusPublish) UnmarshalJSON(data []byte) error {
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
	OfRemoveBg          *FileUploadParamsExtensionRemoveBg          `json:",omitzero,inline"`
	OfAutoTagging       *FileUploadParamsExtensionAutoTagging       `json:",omitzero,inline"`
	OfAIAutoDescription *FileUploadParamsExtensionAIAutoDescription `json:",omitzero,inline"`
	paramUnion
}

func (u FileUploadParamsExtensionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRemoveBg, u.OfAutoTagging, u.OfAIAutoDescription)
}
func (u *FileUploadParamsExtensionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FileUploadParamsExtensionUnion) asAny() any {
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
func (u FileUploadParamsExtensionUnion) GetOptions() *FileUploadParamsExtensionRemoveBgOptions {
	if vt := u.OfRemoveBg; vt != nil {
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
	apijson.RegisterUnion[FileUploadParamsExtensionUnion](
		"name",
		apijson.Discriminator[FileUploadParamsExtensionRemoveBg]("remove-bg"),
		apijson.Discriminator[FileUploadParamsExtensionAutoTagging]("google-auto-tagging"),
		apijson.Discriminator[FileUploadParamsExtensionAutoTagging]("aws-auto-tagging"),
		apijson.Discriminator[FileUploadParamsExtensionAIAutoDescription]("ai-auto-description"),
	)
}

// The property Name is required.
type FileUploadParamsExtensionRemoveBg struct {
	Options FileUploadParamsExtensionRemoveBgOptions `json:"options,omitzero"`
	// Specifies the background removal extension.
	//
	// This field can be elided, and will marshal its zero value as "remove-bg".
	Name constant.RemoveBg `json:"name,required"`
	paramObj
}

func (r FileUploadParamsExtensionRemoveBg) MarshalJSON() (data []byte, err error) {
	type shadow FileUploadParamsExtensionRemoveBg
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUploadParamsExtensionRemoveBg) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileUploadParamsExtensionRemoveBgOptions struct {
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

func (r FileUploadParamsExtensionRemoveBgOptions) MarshalJSON() (data []byte, err error) {
	type shadow FileUploadParamsExtensionRemoveBgOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUploadParamsExtensionRemoveBgOptions) UnmarshalJSON(data []byte) error {
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

func NewFileUploadParamsExtensionAIAutoDescription() FileUploadParamsExtensionAIAutoDescription {
	return FileUploadParamsExtensionAIAutoDescription{
		Name: "ai-auto-description",
	}
}

// This struct has a constant value, construct it with
// [NewFileUploadParamsExtensionAIAutoDescription].
type FileUploadParamsExtensionAIAutoDescription struct {
	// Specifies the auto description extension.
	Name constant.AIAutoDescription `json:"name,required"`
	paramObj
}

func (r FileUploadParamsExtensionAIAutoDescription) MarshalJSON() (data []byte, err error) {
	type shadow FileUploadParamsExtensionAIAutoDescription
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUploadParamsExtensionAIAutoDescription) UnmarshalJSON(data []byte) error {
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
	OfTransformation *FileUploadParamsTransformationPostTransformation `json:",omitzero,inline"`
	OfGifToVideo     *FileUploadParamsTransformationPostGifToVideo     `json:",omitzero,inline"`
	OfThumbnail      *FileUploadParamsTransformationPostThumbnail      `json:",omitzero,inline"`
	OfAbs            *FileUploadParamsTransformationPostAbs            `json:",omitzero,inline"`
	paramUnion
}

func (u FileUploadParamsTransformationPostUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfTransformation, u.OfGifToVideo, u.OfThumbnail, u.OfAbs)
}
func (u *FileUploadParamsTransformationPostUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *FileUploadParamsTransformationPostUnion) asAny() any {
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
func (u FileUploadParamsTransformationPostUnion) GetProtocol() *string {
	if vt := u.OfAbs; vt != nil {
		return &vt.Protocol
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u FileUploadParamsTransformationPostUnion) GetType() *string {
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
func (u FileUploadParamsTransformationPostUnion) GetValue() *string {
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
	apijson.RegisterUnion[FileUploadParamsTransformationPostUnion](
		"type",
		apijson.Discriminator[FileUploadParamsTransformationPostTransformation]("transformation"),
		apijson.Discriminator[FileUploadParamsTransformationPostGifToVideo]("gif-to-video"),
		apijson.Discriminator[FileUploadParamsTransformationPostThumbnail]("thumbnail"),
		apijson.Discriminator[FileUploadParamsTransformationPostAbs]("abs"),
	)
}

// The properties Type, Value are required.
type FileUploadParamsTransformationPostTransformation struct {
	// Transformation string (e.g. `w-200,h-200`).
	// Same syntax as ImageKit URL-based transformations.
	Value string `json:"value,required"`
	// Transformation type.
	//
	// This field can be elided, and will marshal its zero value as "transformation".
	Type constant.Transformation `json:"type,required"`
	paramObj
}

func (r FileUploadParamsTransformationPostTransformation) MarshalJSON() (data []byte, err error) {
	type shadow FileUploadParamsTransformationPostTransformation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUploadParamsTransformationPostTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type FileUploadParamsTransformationPostGifToVideo struct {
	// Optional transformation string to apply to the output video.
	// **Example**: `q-80`
	Value param.Opt[string] `json:"value,omitzero"`
	// Converts an animated GIF into an MP4.
	//
	// This field can be elided, and will marshal its zero value as "gif-to-video".
	Type constant.GifToVideo `json:"type,required"`
	paramObj
}

func (r FileUploadParamsTransformationPostGifToVideo) MarshalJSON() (data []byte, err error) {
	type shadow FileUploadParamsTransformationPostGifToVideo
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUploadParamsTransformationPostGifToVideo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type FileUploadParamsTransformationPostThumbnail struct {
	// Optional transformation string.
	// **Example**: `w-150,h-150`
	Value param.Opt[string] `json:"value,omitzero"`
	// Generates a thumbnail image.
	//
	// This field can be elided, and will marshal its zero value as "thumbnail".
	Type constant.Thumbnail `json:"type,required"`
	paramObj
}

func (r FileUploadParamsTransformationPostThumbnail) MarshalJSON() (data []byte, err error) {
	type shadow FileUploadParamsTransformationPostThumbnail
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUploadParamsTransformationPostThumbnail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Protocol, Type, Value are required.
type FileUploadParamsTransformationPostAbs struct {
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

func (r FileUploadParamsTransformationPostAbs) MarshalJSON() (data []byte, err error) {
	type shadow FileUploadParamsTransformationPostAbs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileUploadParamsTransformationPostAbs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[FileUploadParamsTransformationPostAbs](
		"protocol", "hls", "dash",
	)
}
