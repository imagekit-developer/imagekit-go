// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/stainless-sdks/imagekit-go/internal/apijson"
	"github.com/stainless-sdks/imagekit-go/packages/param"
	"github.com/stainless-sdks/imagekit-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// The property Name is required.
type AutoDescriptionExtensionParam struct {
	// Specifies the auto description extension.
	//
	// Any of "ai-auto-description".
	Name AutoDescriptionExtensionName `json:"name,omitzero,required"`
	paramObj
}

func (r AutoDescriptionExtensionParam) MarshalJSON() (data []byte, err error) {
	type shadow AutoDescriptionExtensionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutoDescriptionExtensionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the auto description extension.
type AutoDescriptionExtensionName string

const (
	AutoDescriptionExtensionNameAIAutoDescription AutoDescriptionExtensionName = "ai-auto-description"
)

// The properties MaxTags, MinConfidence, Name are required.
type AutoTaggingExtensionParam struct {
	// Maximum number of tags to attach to the asset.
	MaxTags int64 `json:"maxTags,required"`
	// Minimum confidence level for tags to be considered valid.
	MinConfidence int64 `json:"minConfidence,required"`
	// Specifies the auto-tagging extension used.
	//
	// Any of "google-auto-tagging", "aws-auto-tagging".
	Name AutoTaggingExtensionName `json:"name,omitzero,required"`
	paramObj
}

func (r AutoTaggingExtensionParam) MarshalJSON() (data []byte, err error) {
	type shadow AutoTaggingExtensionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AutoTaggingExtensionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the auto-tagging extension used.
type AutoTaggingExtensionName string

const (
	AutoTaggingExtensionNameGoogleAutoTagging AutoTaggingExtensionName = "google-auto-tagging"
	AutoTaggingExtensionNameAwsAutoTagging    AutoTaggingExtensionName = "aws-auto-tagging"
)

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

// The property Name is required.
type RemovedotBgExtensionParam struct {
	// Specifies the background removal extension.
	//
	// Any of "remove-bg".
	Name    RemovedotBgExtensionName         `json:"name,omitzero,required"`
	Options RemovedotBgExtensionOptionsParam `json:"options,omitzero"`
	paramObj
}

func (r RemovedotBgExtensionParam) MarshalJSON() (data []byte, err error) {
	type shadow RemovedotBgExtensionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RemovedotBgExtensionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the background removal extension.
type RemovedotBgExtensionName string

const (
	RemovedotBgExtensionNameRemoveBg RemovedotBgExtensionName = "remove-bg"
)

type RemovedotBgExtensionOptionsParam struct {
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

func (r RemovedotBgExtensionOptionsParam) MarshalJSON() (data []byte, err error) {
	type shadow RemovedotBgExtensionOptionsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RemovedotBgExtensionOptionsParam) UnmarshalJSON(data []byte) error {
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
