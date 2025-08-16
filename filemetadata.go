// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/imagekit-go/internal/apijson"
	"github.com/stainless-sdks/imagekit-go/internal/apiquery"
	"github.com/stainless-sdks/imagekit-go/internal/requestconfig"
	"github.com/stainless-sdks/imagekit-go/option"
	"github.com/stainless-sdks/imagekit-go/packages/respjson"
)

// FileMetadataService contains methods and other services that help with
// interacting with the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileMetadataService] method instead.
type FileMetadataService struct {
	Options []option.RequestOption
}

// NewFileMetadataService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFileMetadataService(opts ...option.RequestOption) (r FileMetadataService) {
	r = FileMetadataService{}
	r.Options = opts
	return
}

// You can programmatically get image EXIF, pHash, and other metadata for uploaded
// files in the ImageKit.io media library using this API.
//
// You can also get the metadata in upload API response by passing `metadata` in
// `responseFields` parameter.
func (r *FileMetadataService) Get(ctx context.Context, fileID string, opts ...option.RequestOption) (res *FileMetadataGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return
	}
	path := fmt.Sprintf("v1/files/%s/metadata", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get image EXIF, pHash, and other metadata from ImageKit.io powered remote URL
// using this API.
func (r *FileMetadataService) FromURL(ctx context.Context, query FileMetadataFromURLParams, opts ...option.RequestOption) (res *FileMetadataFromURLResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/files/metadata"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// JSON object containing metadata.
type FileMetadataGetResponse struct {
	// The audio codec used in the video (only for video).
	AudioCodec string `json:"audioCodec"`
	// The bit rate of the video in kbps (only for video).
	BitRate int64 `json:"bitRate"`
	// The density of the image in DPI.
	Density int64 `json:"density"`
	// The duration of the video in seconds (only for video).
	Duration int64                       `json:"duration"`
	Exif     FileMetadataGetResponseExif `json:"exif"`
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
func (r FileMetadataGetResponse) RawJSON() string { return r.JSON.raw }
func (r *FileMetadataGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileMetadataGetResponseExif struct {
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
func (r FileMetadataGetResponseExif) RawJSON() string { return r.JSON.raw }
func (r *FileMetadataGetResponseExif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// JSON object containing metadata.
type FileMetadataFromURLResponse struct {
	// The audio codec used in the video (only for video).
	AudioCodec string `json:"audioCodec"`
	// The bit rate of the video in kbps (only for video).
	BitRate int64 `json:"bitRate"`
	// The density of the image in DPI.
	Density int64 `json:"density"`
	// The duration of the video in seconds (only for video).
	Duration int64                           `json:"duration"`
	Exif     FileMetadataFromURLResponseExif `json:"exif"`
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
func (r FileMetadataFromURLResponse) RawJSON() string { return r.JSON.raw }
func (r *FileMetadataFromURLResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileMetadataFromURLResponseExif struct {
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
func (r FileMetadataFromURLResponseExif) RawJSON() string { return r.JSON.raw }
func (r *FileMetadataFromURLResponseExif) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileMetadataFromURLParams struct {
	// Should be a valid file URL. It should be accessible using your ImageKit.io
	// account.
	URL string `query:"url,required" json:"-"`
	paramObj
}

// URLQuery serializes [FileMetadataFromURLParams]'s query parameters as
// `url.Values`.
func (r FileMetadataFromURLParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
