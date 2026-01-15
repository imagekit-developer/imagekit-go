// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"encoding/json"
	"errors"
	"net/http"
	"slices"
	"time"

	"github.com/imagekit-developer/imagekit-go/v2/internal/apijson"
	"github.com/imagekit-developer/imagekit-go/v2/internal/requestconfig"
	"github.com/imagekit-developer/imagekit-go/v2/option"
	"github.com/imagekit-developer/imagekit-go/v2/packages/respjson"
	"github.com/imagekit-developer/imagekit-go/v2/shared/constant"
	standardwebhooks "github.com/standard-webhooks/standard-webhooks/libraries/go"
)

// WebhookService contains methods and other services that help with interacting
// with the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebhookService] method instead.
type WebhookService struct {
	Options []option.RequestOption
}

// NewWebhookService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebhookService(opts ...option.RequestOption) (r WebhookService) {
	r = WebhookService{}
	r.Options = opts
	return
}

func (r *WebhookService) UnsafeUnwrap(payload []byte, opts ...option.RequestOption) (*UnsafeUnwrapWebhookEventUnion, error) {
	res := &UnsafeUnwrapWebhookEventUnion{}
	err := res.UnmarshalJSON(payload)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r *WebhookService) Unwrap(payload []byte, headers http.Header, opts ...option.RequestOption) (*UnwrapWebhookEventUnion, error) {
	opts = slices.Concat(r.Options, opts)
	cfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	key := cfg.WebhookSecret
	if key == "" {
		return nil, errors.New("The WebhookSecret option must be set in order to verify webhook headers")
	}
	wh, err := standardwebhooks.NewWebhook(key)
	if err != nil {
		return nil, err
	}
	err = wh.Verify(payload, headers)
	if err != nil {
		return nil, err
	}
	res := &UnwrapWebhookEventUnion{}
	err = res.UnmarshalJSON(payload)
	if err != nil {
		return res, err
	}
	return res, nil
}

type BaseWebhookEvent struct {
	// Unique identifier for the event.
	ID string `json:"id,required"`
	// The type of webhook event.
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaseWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *BaseWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Triggered when a post-transformation fails. The original file remains available,
// but the requested transformation could not be generated.
type UploadPostTransformErrorEvent struct {
	// Timestamp of when the event occurred in ISO8601 format.
	CreatedAt time.Time                            `json:"created_at,required" format:"date-time"`
	Data      UploadPostTransformErrorEventData    `json:"data,required"`
	Request   UploadPostTransformErrorEventRequest `json:"request,required"`
	Type      constant.UploadPostTransformError    `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Data        respjson.Field
		Request     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BaseWebhookEvent
}

// Returns the unmodified JSON received from the API
func (r UploadPostTransformErrorEvent) RawJSON() string { return r.JSON.raw }
func (r *UploadPostTransformErrorEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UploadPostTransformErrorEventData struct {
	// Unique identifier of the originally uploaded file.
	FileID string `json:"fileId,required"`
	// Name of the file.
	Name string `json:"name,required"`
	// Path of the file.
	Path           string                                          `json:"path,required"`
	Transformation UploadPostTransformErrorEventDataTransformation `json:"transformation,required"`
	// URL of the attempted post-transformation.
	URL string `json:"url,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID         respjson.Field
		Name           respjson.Field
		Path           respjson.Field
		Transformation respjson.Field
		URL            respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UploadPostTransformErrorEventData) RawJSON() string { return r.JSON.raw }
func (r *UploadPostTransformErrorEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UploadPostTransformErrorEventDataTransformation struct {
	Error UploadPostTransformErrorEventDataTransformationError `json:"error,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UploadPostTransformErrorEventDataTransformation) RawJSON() string { return r.JSON.raw }
func (r *UploadPostTransformErrorEventDataTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UploadPostTransformErrorEventDataTransformationError struct {
	// Reason for the post-transformation failure.
	Reason string `json:"reason,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UploadPostTransformErrorEventDataTransformationError) RawJSON() string { return r.JSON.raw }
func (r *UploadPostTransformErrorEventDataTransformationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UploadPostTransformErrorEventRequest struct {
	Transformation UploadPostTransformErrorEventRequestTransformation `json:"transformation,required"`
	// Unique identifier for the originating request.
	XRequestID string `json:"x_request_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Transformation respjson.Field
		XRequestID     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UploadPostTransformErrorEventRequest) RawJSON() string { return r.JSON.raw }
func (r *UploadPostTransformErrorEventRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UploadPostTransformErrorEventRequestTransformation struct {
	// Type of the requested post-transformation.
	//
	// Any of "transformation", "abs", "gif-to-video", "thumbnail".
	Type string `json:"type,required"`
	// Only applicable if transformation type is 'abs'. Streaming protocol used.
	//
	// Any of "hls", "dash".
	Protocol string `json:"protocol"`
	// Value for the requested transformation type.
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Protocol    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UploadPostTransformErrorEventRequestTransformation) RawJSON() string { return r.JSON.raw }
func (r *UploadPostTransformErrorEventRequestTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Triggered when a post-transformation completes successfully. The transformed
// version of the file is now ready and can be accessed via the provided URL. Note
// that each post-transformation generates a separate webhook event.
type UploadPostTransformSuccessEvent struct {
	// Timestamp of when the event occurred in ISO8601 format.
	CreatedAt time.Time                              `json:"created_at,required" format:"date-time"`
	Data      UploadPostTransformSuccessEventData    `json:"data,required"`
	Request   UploadPostTransformSuccessEventRequest `json:"request,required"`
	Type      constant.UploadPostTransformSuccess    `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Data        respjson.Field
		Request     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BaseWebhookEvent
}

// Returns the unmodified JSON received from the API
func (r UploadPostTransformSuccessEvent) RawJSON() string { return r.JSON.raw }
func (r *UploadPostTransformSuccessEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UploadPostTransformSuccessEventData struct {
	// Unique identifier of the originally uploaded file.
	FileID string `json:"fileId,required"`
	// Name of the file.
	Name string `json:"name,required"`
	// URL of the generated post-transformation.
	URL string `json:"url,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID      respjson.Field
		Name        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UploadPostTransformSuccessEventData) RawJSON() string { return r.JSON.raw }
func (r *UploadPostTransformSuccessEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UploadPostTransformSuccessEventRequest struct {
	Transformation UploadPostTransformSuccessEventRequestTransformation `json:"transformation,required"`
	// Unique identifier for the originating request.
	XRequestID string `json:"x_request_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Transformation respjson.Field
		XRequestID     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UploadPostTransformSuccessEventRequest) RawJSON() string { return r.JSON.raw }
func (r *UploadPostTransformSuccessEventRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UploadPostTransformSuccessEventRequestTransformation struct {
	// Type of the requested post-transformation.
	//
	// Any of "transformation", "abs", "gif-to-video", "thumbnail".
	Type string `json:"type,required"`
	// Only applicable if transformation type is 'abs'. Streaming protocol used.
	//
	// Any of "hls", "dash".
	Protocol string `json:"protocol"`
	// Value for the requested transformation type.
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Protocol    respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UploadPostTransformSuccessEventRequestTransformation) RawJSON() string { return r.JSON.raw }
func (r *UploadPostTransformSuccessEventRequestTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Triggered when a pre-transformation fails. The file upload may have been
// accepted, but the requested transformation could not be applied.
type UploadPreTransformErrorEvent struct {
	// Timestamp of when the event occurred in ISO8601 format.
	CreatedAt time.Time                           `json:"created_at,required" format:"date-time"`
	Data      UploadPreTransformErrorEventData    `json:"data,required"`
	Request   UploadPreTransformErrorEventRequest `json:"request,required"`
	Type      constant.UploadPreTransformError    `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Data        respjson.Field
		Request     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BaseWebhookEvent
}

// Returns the unmodified JSON received from the API
func (r UploadPreTransformErrorEvent) RawJSON() string { return r.JSON.raw }
func (r *UploadPreTransformErrorEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UploadPreTransformErrorEventData struct {
	// Name of the file.
	Name string `json:"name,required"`
	// Path of the file.
	Path           string                                         `json:"path,required"`
	Transformation UploadPreTransformErrorEventDataTransformation `json:"transformation,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name           respjson.Field
		Path           respjson.Field
		Transformation respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UploadPreTransformErrorEventData) RawJSON() string { return r.JSON.raw }
func (r *UploadPreTransformErrorEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UploadPreTransformErrorEventDataTransformation struct {
	Error UploadPreTransformErrorEventDataTransformationError `json:"error,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UploadPreTransformErrorEventDataTransformation) RawJSON() string { return r.JSON.raw }
func (r *UploadPreTransformErrorEventDataTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UploadPreTransformErrorEventDataTransformationError struct {
	// Reason for the pre-transformation failure.
	Reason string `json:"reason,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UploadPreTransformErrorEventDataTransformationError) RawJSON() string { return r.JSON.raw }
func (r *UploadPreTransformErrorEventDataTransformationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UploadPreTransformErrorEventRequest struct {
	// The requested pre-transformation string.
	Transformation string `json:"transformation,required"`
	// Unique identifier for the originating request.
	XRequestID string `json:"x_request_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Transformation respjson.Field
		XRequestID     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UploadPreTransformErrorEventRequest) RawJSON() string { return r.JSON.raw }
func (r *UploadPreTransformErrorEventRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Triggered when a pre-transformation completes successfully. The file has been
// processed with the requested transformation and is now available in the Media
// Library.
type UploadPreTransformSuccessEvent struct {
	// Timestamp of when the event occurred in ISO8601 format.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Object containing details of a successful upload.
	Data    UploadPreTransformSuccessEventData    `json:"data,required"`
	Request UploadPreTransformSuccessEventRequest `json:"request,required"`
	Type    constant.UploadPreTransformSuccess    `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Data        respjson.Field
		Request     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BaseWebhookEvent
}

// Returns the unmodified JSON received from the API
func (r UploadPreTransformSuccessEvent) RawJSON() string { return r.JSON.raw }
func (r *UploadPreTransformSuccessEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Object containing details of a successful upload.
type UploadPreTransformSuccessEventData struct {
	// An array of tags assigned to the uploaded file by auto tagging.
	AITags []UploadPreTransformSuccessEventDataAITag `json:"AITags,nullable"`
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
	ExtensionStatus UploadPreTransformSuccessEventDataExtensionStatus `json:"extensionStatus"`
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
	// This field is included in the response only if the Path policy feature is
	// available in the plan. It contains schema definitions for the custom metadata
	// fields selected for the specified file path. Field selection can only be done
	// when the Path policy feature is enabled.
	//
	// Keys are the names of the custom metadata fields; the value object has details
	// about the custom metadata schema.
	SelectedFieldsSchema map[string]UploadPreTransformSuccessEventDataSelectedFieldsSchema `json:"selectedFieldsSchema"`
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
	VersionInfo UploadPreTransformSuccessEventDataVersionInfo `json:"versionInfo"`
	// The video codec used in the video (only for video).
	VideoCodec string `json:"videoCodec"`
	// Width of the image in pixels (Only for Images)
	Width float64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AITags               respjson.Field
		AudioCodec           respjson.Field
		BitRate              respjson.Field
		CustomCoordinates    respjson.Field
		CustomMetadata       respjson.Field
		Description          respjson.Field
		Duration             respjson.Field
		EmbeddedMetadata     respjson.Field
		ExtensionStatus      respjson.Field
		FileID               respjson.Field
		FilePath             respjson.Field
		FileType             respjson.Field
		Height               respjson.Field
		IsPrivateFile        respjson.Field
		IsPublished          respjson.Field
		Metadata             respjson.Field
		Name                 respjson.Field
		SelectedFieldsSchema respjson.Field
		Size                 respjson.Field
		Tags                 respjson.Field
		ThumbnailURL         respjson.Field
		URL                  respjson.Field
		VersionInfo          respjson.Field
		VideoCodec           respjson.Field
		Width                respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UploadPreTransformSuccessEventData) RawJSON() string { return r.JSON.raw }
func (r *UploadPreTransformSuccessEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UploadPreTransformSuccessEventDataAITag struct {
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
func (r UploadPreTransformSuccessEventDataAITag) RawJSON() string { return r.JSON.raw }
func (r *UploadPreTransformSuccessEventDataAITag) UnmarshalJSON(data []byte) error {
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
type UploadPreTransformSuccessEventDataExtensionStatus struct {
	// Any of "success", "pending", "failed".
	AIAutoDescription string `json:"ai-auto-description"`
	// Any of "success", "pending", "failed".
	AITasks string `json:"ai-tasks"`
	// Any of "success", "pending", "failed".
	AwsAutoTagging string `json:"aws-auto-tagging"`
	// Any of "success", "pending", "failed".
	GoogleAutoTagging string `json:"google-auto-tagging"`
	// Any of "success", "pending", "failed".
	RemoveBg string `json:"remove-bg"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AIAutoDescription respjson.Field
		AITasks           respjson.Field
		AwsAutoTagging    respjson.Field
		GoogleAutoTagging respjson.Field
		RemoveBg          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UploadPreTransformSuccessEventDataExtensionStatus) RawJSON() string { return r.JSON.raw }
func (r *UploadPreTransformSuccessEventDataExtensionStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UploadPreTransformSuccessEventDataSelectedFieldsSchema struct {
	// Type of the custom metadata field.
	//
	// Any of "Text", "Textarea", "Number", "Date", "Boolean", "SingleSelect",
	// "MultiSelect".
	Type string `json:"type,required"`
	// The default value for this custom metadata field. The value should match the
	// `type` of custom metadata field.
	DefaultValue UploadPreTransformSuccessEventDataSelectedFieldsSchemaDefaultValueUnion `json:"defaultValue"`
	// Specifies if the custom metadata field is required or not.
	IsValueRequired bool `json:"isValueRequired"`
	// Maximum length of string. Only set if `type` is set to `Text` or `Textarea`.
	MaxLength float64 `json:"maxLength"`
	// Maximum value of the field. Only set if field type is `Date` or `Number`. For
	// `Date` type field, the value will be in ISO8601 string format. For `Number` type
	// field, it will be a numeric value.
	MaxValue UploadPreTransformSuccessEventDataSelectedFieldsSchemaMaxValueUnion `json:"maxValue"`
	// Minimum length of string. Only set if `type` is set to `Text` or `Textarea`.
	MinLength float64 `json:"minLength"`
	// Minimum value of the field. Only set if field type is `Date` or `Number`. For
	// `Date` type field, the value will be in ISO8601 string format. For `Number` type
	// field, it will be a numeric value.
	MinValue UploadPreTransformSuccessEventDataSelectedFieldsSchemaMinValueUnion `json:"minValue"`
	// Indicates whether the custom metadata field is read only. A read only field
	// cannot be modified after being set. This field is configurable only via the
	// **Path policy** feature.
	ReadOnly bool `json:"readOnly"`
	// An array of allowed values when field type is `SingleSelect` or `MultiSelect`.
	SelectOptions []UploadPreTransformSuccessEventDataSelectedFieldsSchemaSelectOptionUnion `json:"selectOptions"`
	// Specifies if the selectOptions array is truncated. It is truncated when number
	// of options are > 100.
	SelectOptionsTruncated bool `json:"selectOptionsTruncated"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type                   respjson.Field
		DefaultValue           respjson.Field
		IsValueRequired        respjson.Field
		MaxLength              respjson.Field
		MaxValue               respjson.Field
		MinLength              respjson.Field
		MinValue               respjson.Field
		ReadOnly               respjson.Field
		SelectOptions          respjson.Field
		SelectOptionsTruncated respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UploadPreTransformSuccessEventDataSelectedFieldsSchema) RawJSON() string { return r.JSON.raw }
func (r *UploadPreTransformSuccessEventDataSelectedFieldsSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UploadPreTransformSuccessEventDataSelectedFieldsSchemaDefaultValueUnion contains
// all possible properties and values from [string], [float64], [bool],
// [[]UploadPreTransformSuccessEventDataSelectedFieldsSchemaDefaultValueMixedItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfMixed]
type UploadPreTransformSuccessEventDataSelectedFieldsSchemaDefaultValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]UploadPreTransformSuccessEventDataSelectedFieldsSchemaDefaultValueMixedItemUnion]
	// instead of an object.
	OfMixed []UploadPreTransformSuccessEventDataSelectedFieldsSchemaDefaultValueMixedItemUnion `json:",inline"`
	JSON    struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		OfMixed  respjson.Field
		raw      string
	} `json:"-"`
}

func (u UploadPreTransformSuccessEventDataSelectedFieldsSchemaDefaultValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UploadPreTransformSuccessEventDataSelectedFieldsSchemaDefaultValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UploadPreTransformSuccessEventDataSelectedFieldsSchemaDefaultValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UploadPreTransformSuccessEventDataSelectedFieldsSchemaDefaultValueUnion) AsMixed() (v []UploadPreTransformSuccessEventDataSelectedFieldsSchemaDefaultValueMixedItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UploadPreTransformSuccessEventDataSelectedFieldsSchemaDefaultValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *UploadPreTransformSuccessEventDataSelectedFieldsSchemaDefaultValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UploadPreTransformSuccessEventDataSelectedFieldsSchemaDefaultValueMixedItemUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type UploadPreTransformSuccessEventDataSelectedFieldsSchemaDefaultValueMixedItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u UploadPreTransformSuccessEventDataSelectedFieldsSchemaDefaultValueMixedItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UploadPreTransformSuccessEventDataSelectedFieldsSchemaDefaultValueMixedItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UploadPreTransformSuccessEventDataSelectedFieldsSchemaDefaultValueMixedItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UploadPreTransformSuccessEventDataSelectedFieldsSchemaDefaultValueMixedItemUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *UploadPreTransformSuccessEventDataSelectedFieldsSchemaDefaultValueMixedItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UploadPreTransformSuccessEventDataSelectedFieldsSchemaMaxValueUnion contains all
// possible properties and values from [string], [float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat]
type UploadPreTransformSuccessEventDataSelectedFieldsSchemaMaxValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	JSON    struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		raw      string
	} `json:"-"`
}

func (u UploadPreTransformSuccessEventDataSelectedFieldsSchemaMaxValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UploadPreTransformSuccessEventDataSelectedFieldsSchemaMaxValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UploadPreTransformSuccessEventDataSelectedFieldsSchemaMaxValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *UploadPreTransformSuccessEventDataSelectedFieldsSchemaMaxValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UploadPreTransformSuccessEventDataSelectedFieldsSchemaMinValueUnion contains all
// possible properties and values from [string], [float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat]
type UploadPreTransformSuccessEventDataSelectedFieldsSchemaMinValueUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	JSON    struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		raw      string
	} `json:"-"`
}

func (u UploadPreTransformSuccessEventDataSelectedFieldsSchemaMinValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UploadPreTransformSuccessEventDataSelectedFieldsSchemaMinValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UploadPreTransformSuccessEventDataSelectedFieldsSchemaMinValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *UploadPreTransformSuccessEventDataSelectedFieldsSchemaMinValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UploadPreTransformSuccessEventDataSelectedFieldsSchemaSelectOptionUnion contains
// all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type UploadPreTransformSuccessEventDataSelectedFieldsSchemaSelectOptionUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u UploadPreTransformSuccessEventDataSelectedFieldsSchemaSelectOptionUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UploadPreTransformSuccessEventDataSelectedFieldsSchemaSelectOptionUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UploadPreTransformSuccessEventDataSelectedFieldsSchemaSelectOptionUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UploadPreTransformSuccessEventDataSelectedFieldsSchemaSelectOptionUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *UploadPreTransformSuccessEventDataSelectedFieldsSchemaSelectOptionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An object containing the file or file version's `id` (versionId) and `name`.
type UploadPreTransformSuccessEventDataVersionInfo struct {
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
func (r UploadPreTransformSuccessEventDataVersionInfo) RawJSON() string { return r.JSON.raw }
func (r *UploadPreTransformSuccessEventDataVersionInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UploadPreTransformSuccessEventRequest struct {
	// The requested pre-transformation string.
	Transformation string `json:"transformation,required"`
	// Unique identifier for the originating request.
	XRequestID string `json:"x_request_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Transformation respjson.Field
		XRequestID     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UploadPreTransformSuccessEventRequest) RawJSON() string { return r.JSON.raw }
func (r *UploadPreTransformSuccessEventRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Triggered when a new video transformation request is accepted for processing.
// This event confirms that ImageKit has received and queued your transformation
// request. Use this for debugging and tracking transformation lifecycle.
type VideoTransformationAcceptedEvent struct {
	// Timestamp when the event was created in ISO8601 format.
	CreatedAt time.Time                            `json:"created_at,required" format:"date-time"`
	Data      VideoTransformationAcceptedEventData `json:"data,required"`
	// Information about the original request that triggered the video transformation.
	Request VideoTransformationAcceptedEventRequest `json:"request,required"`
	Type    constant.VideoTransformationAccepted    `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Data        respjson.Field
		Request     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BaseWebhookEvent
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationAcceptedEvent) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationAcceptedEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationAcceptedEventData struct {
	// Information about the source video asset being transformed.
	Asset VideoTransformationAcceptedEventDataAsset `json:"asset,required"`
	// Base information about a video transformation request.
	Transformation VideoTransformationAcceptedEventDataTransformation `json:"transformation,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Asset          respjson.Field
		Transformation respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationAcceptedEventData) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationAcceptedEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the source video asset being transformed.
type VideoTransformationAcceptedEventDataAsset struct {
	// URL to download or access the source video file.
	URL string `json:"url,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationAcceptedEventDataAsset) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationAcceptedEventDataAsset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Base information about a video transformation request.
type VideoTransformationAcceptedEventDataTransformation struct {
	// Type of video transformation:
	//
	//   - `video-transformation`: Standard video processing (resize, format conversion,
	//     etc.)
	//   - `gif-to-video`: Convert animated GIF to video format
	//   - `video-thumbnail`: Generate thumbnail image from video
	//
	// Any of "video-transformation", "gif-to-video", "video-thumbnail".
	Type string `json:"type,required"`
	// Configuration options for video transformations.
	Options VideoTransformationAcceptedEventDataTransformationOptions `json:"options"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Options     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationAcceptedEventDataTransformation) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationAcceptedEventDataTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration options for video transformations.
type VideoTransformationAcceptedEventDataTransformationOptions struct {
	// Audio codec used for encoding (aac or opus).
	//
	// Any of "aac", "opus".
	AudioCodec string `json:"audio_codec"`
	// Whether to automatically rotate the video based on metadata.
	AutoRotate bool `json:"auto_rotate"`
	// Output format for the transformed video or thumbnail.
	//
	// Any of "mp4", "webm", "jpg", "png", "webp".
	Format string `json:"format"`
	// Quality setting for the output video.
	Quality int64 `json:"quality"`
	// Streaming protocol for adaptive bitrate streaming.
	//
	// Any of "HLS", "DASH".
	StreamProtocol string `json:"stream_protocol"`
	// Array of quality representations for adaptive bitrate streaming.
	Variants []string `json:"variants"`
	// Video codec used for encoding (h264, vp9, or av1).
	//
	// Any of "h264", "vp9", "av1".
	VideoCodec string `json:"video_codec"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudioCodec     respjson.Field
		AutoRotate     respjson.Field
		Format         respjson.Field
		Quality        respjson.Field
		StreamProtocol respjson.Field
		Variants       respjson.Field
		VideoCodec     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationAcceptedEventDataTransformationOptions) RawJSON() string {
	return r.JSON.raw
}
func (r *VideoTransformationAcceptedEventDataTransformationOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the original request that triggered the video transformation.
type VideoTransformationAcceptedEventRequest struct {
	// Full URL of the transformation request that was submitted.
	URL string `json:"url,required" format:"uri"`
	// Unique identifier for the originating transformation request.
	XRequestID string `json:"x_request_id,required"`
	// User-Agent header from the original request that triggered the transformation.
	UserAgent string `json:"user_agent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		XRequestID  respjson.Field
		UserAgent   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationAcceptedEventRequest) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationAcceptedEventRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Triggered when an error occurs during video encoding. Listen to this webhook to
// log error reasons and debug issues. Check your origin and URL endpoint settings
// if the reason is related to download failure. For other errors, contact ImageKit
// support.
type VideoTransformationErrorEvent struct {
	// Timestamp when the event was created in ISO8601 format.
	CreatedAt time.Time                         `json:"created_at,required" format:"date-time"`
	Data      VideoTransformationErrorEventData `json:"data,required"`
	// Information about the original request that triggered the video transformation.
	Request VideoTransformationErrorEventRequest `json:"request,required"`
	Type    constant.VideoTransformationError    `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Data        respjson.Field
		Request     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BaseWebhookEvent
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationErrorEvent) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationErrorEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationErrorEventData struct {
	// Information about the source video asset being transformed.
	Asset          VideoTransformationErrorEventDataAsset          `json:"asset,required"`
	Transformation VideoTransformationErrorEventDataTransformation `json:"transformation,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Asset          respjson.Field
		Transformation respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationErrorEventData) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationErrorEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the source video asset being transformed.
type VideoTransformationErrorEventDataAsset struct {
	// URL to download or access the source video file.
	URL string `json:"url,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationErrorEventDataAsset) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationErrorEventDataAsset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationErrorEventDataTransformation struct {
	// Type of video transformation:
	//
	//   - `video-transformation`: Standard video processing (resize, format conversion,
	//     etc.)
	//   - `gif-to-video`: Convert animated GIF to video format
	//   - `video-thumbnail`: Generate thumbnail image from video
	//
	// Any of "video-transformation", "gif-to-video", "video-thumbnail".
	Type string `json:"type,required"`
	// Details about the transformation error.
	Error VideoTransformationErrorEventDataTransformationError `json:"error"`
	// Configuration options for video transformations.
	Options VideoTransformationErrorEventDataTransformationOptions `json:"options"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Error       respjson.Field
		Options     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationErrorEventDataTransformation) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationErrorEventDataTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details about the transformation error.
type VideoTransformationErrorEventDataTransformationError struct {
	// Specific reason for the transformation failure:
	//
	// - `encoding_failed`: Error during video encoding process
	// - `download_failed`: Could not download source video
	// - `internal_server_error`: Unexpected server error
	//
	// Any of "encoding_failed", "download_failed", "internal_server_error".
	Reason string `json:"reason,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationErrorEventDataTransformationError) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationErrorEventDataTransformationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration options for video transformations.
type VideoTransformationErrorEventDataTransformationOptions struct {
	// Audio codec used for encoding (aac or opus).
	//
	// Any of "aac", "opus".
	AudioCodec string `json:"audio_codec"`
	// Whether to automatically rotate the video based on metadata.
	AutoRotate bool `json:"auto_rotate"`
	// Output format for the transformed video or thumbnail.
	//
	// Any of "mp4", "webm", "jpg", "png", "webp".
	Format string `json:"format"`
	// Quality setting for the output video.
	Quality int64 `json:"quality"`
	// Streaming protocol for adaptive bitrate streaming.
	//
	// Any of "HLS", "DASH".
	StreamProtocol string `json:"stream_protocol"`
	// Array of quality representations for adaptive bitrate streaming.
	Variants []string `json:"variants"`
	// Video codec used for encoding (h264, vp9, or av1).
	//
	// Any of "h264", "vp9", "av1".
	VideoCodec string `json:"video_codec"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudioCodec     respjson.Field
		AutoRotate     respjson.Field
		Format         respjson.Field
		Quality        respjson.Field
		StreamProtocol respjson.Field
		Variants       respjson.Field
		VideoCodec     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationErrorEventDataTransformationOptions) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationErrorEventDataTransformationOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the original request that triggered the video transformation.
type VideoTransformationErrorEventRequest struct {
	// Full URL of the transformation request that was submitted.
	URL string `json:"url,required" format:"uri"`
	// Unique identifier for the originating transformation request.
	XRequestID string `json:"x_request_id,required"`
	// User-Agent header from the original request that triggered the transformation.
	UserAgent string `json:"user_agent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		XRequestID  respjson.Field
		UserAgent   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationErrorEventRequest) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationErrorEventRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Triggered when video encoding is finished and the transformed resource is ready
// to be served. This is the key event to listen for - update your database or CMS
// flags when you receive this so your application can start showing the
// transformed video to users.
type VideoTransformationReadyEvent struct {
	// Timestamp when the event was created in ISO8601 format.
	CreatedAt time.Time                         `json:"created_at,required" format:"date-time"`
	Data      VideoTransformationReadyEventData `json:"data,required"`
	// Information about the original request that triggered the video transformation.
	Request VideoTransformationReadyEventRequest `json:"request,required"`
	Type    constant.VideoTransformationReady    `json:"type,required"`
	// Performance metrics for the transformation process.
	Timings VideoTransformationReadyEventTimings `json:"timings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Data        respjson.Field
		Request     respjson.Field
		Type        respjson.Field
		Timings     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	BaseWebhookEvent
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationReadyEvent) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationReadyEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationReadyEventData struct {
	// Information about the source video asset being transformed.
	Asset          VideoTransformationReadyEventDataAsset          `json:"asset,required"`
	Transformation VideoTransformationReadyEventDataTransformation `json:"transformation,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Asset          respjson.Field
		Transformation respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationReadyEventData) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationReadyEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the source video asset being transformed.
type VideoTransformationReadyEventDataAsset struct {
	// URL to download or access the source video file.
	URL string `json:"url,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationReadyEventDataAsset) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationReadyEventDataAsset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationReadyEventDataTransformation struct {
	// Type of video transformation:
	//
	//   - `video-transformation`: Standard video processing (resize, format conversion,
	//     etc.)
	//   - `gif-to-video`: Convert animated GIF to video format
	//   - `video-thumbnail`: Generate thumbnail image from video
	//
	// Any of "video-transformation", "gif-to-video", "video-thumbnail".
	Type string `json:"type,required"`
	// Configuration options for video transformations.
	Options VideoTransformationReadyEventDataTransformationOptions `json:"options"`
	// Information about the transformed output video.
	Output VideoTransformationReadyEventDataTransformationOutput `json:"output"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Options     respjson.Field
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationReadyEventDataTransformation) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationReadyEventDataTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration options for video transformations.
type VideoTransformationReadyEventDataTransformationOptions struct {
	// Audio codec used for encoding (aac or opus).
	//
	// Any of "aac", "opus".
	AudioCodec string `json:"audio_codec"`
	// Whether to automatically rotate the video based on metadata.
	AutoRotate bool `json:"auto_rotate"`
	// Output format for the transformed video or thumbnail.
	//
	// Any of "mp4", "webm", "jpg", "png", "webp".
	Format string `json:"format"`
	// Quality setting for the output video.
	Quality int64 `json:"quality"`
	// Streaming protocol for adaptive bitrate streaming.
	//
	// Any of "HLS", "DASH".
	StreamProtocol string `json:"stream_protocol"`
	// Array of quality representations for adaptive bitrate streaming.
	Variants []string `json:"variants"`
	// Video codec used for encoding (h264, vp9, or av1).
	//
	// Any of "h264", "vp9", "av1".
	VideoCodec string `json:"video_codec"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudioCodec     respjson.Field
		AutoRotate     respjson.Field
		Format         respjson.Field
		Quality        respjson.Field
		StreamProtocol respjson.Field
		Variants       respjson.Field
		VideoCodec     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationReadyEventDataTransformationOptions) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationReadyEventDataTransformationOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the transformed output video.
type VideoTransformationReadyEventDataTransformationOutput struct {
	// URL to access the transformed video.
	URL string `json:"url,required" format:"uri"`
	// Metadata of the output video file.
	VideoMetadata VideoTransformationReadyEventDataTransformationOutputVideoMetadata `json:"video_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL           respjson.Field
		VideoMetadata respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationReadyEventDataTransformationOutput) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationReadyEventDataTransformationOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata of the output video file.
type VideoTransformationReadyEventDataTransformationOutputVideoMetadata struct {
	// Bitrate of the output video in bits per second.
	Bitrate int64 `json:"bitrate,required"`
	// Duration of the output video in seconds.
	Duration float64 `json:"duration,required"`
	// Height of the output video in pixels.
	Height int64 `json:"height,required"`
	// Width of the output video in pixels.
	Width int64 `json:"width,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bitrate     respjson.Field
		Duration    respjson.Field
		Height      respjson.Field
		Width       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationReadyEventDataTransformationOutputVideoMetadata) RawJSON() string {
	return r.JSON.raw
}
func (r *VideoTransformationReadyEventDataTransformationOutputVideoMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the original request that triggered the video transformation.
type VideoTransformationReadyEventRequest struct {
	// Full URL of the transformation request that was submitted.
	URL string `json:"url,required" format:"uri"`
	// Unique identifier for the originating transformation request.
	XRequestID string `json:"x_request_id,required"`
	// User-Agent header from the original request that triggered the transformation.
	UserAgent string `json:"user_agent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		XRequestID  respjson.Field
		UserAgent   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationReadyEventRequest) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationReadyEventRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Performance metrics for the transformation process.
type VideoTransformationReadyEventTimings struct {
	// Time spent downloading the source video from your origin or media library, in
	// milliseconds.
	DownloadDuration int64 `json:"download_duration"`
	// Time spent encoding the video, in milliseconds.
	EncodingDuration int64 `json:"encoding_duration"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DownloadDuration respjson.Field
		EncodingDuration respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationReadyEventTimings) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationReadyEventTimings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnion contains all possible properties and values from
// [VideoTransformationAcceptedEvent], [VideoTransformationReadyEvent],
// [VideoTransformationErrorEvent], [UploadPreTransformSuccessEvent],
// [UploadPreTransformErrorEvent], [UploadPostTransformSuccessEvent],
// [UploadPostTransformErrorEvent].
//
// Use the [UnsafeUnwrapWebhookEventUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UnsafeUnwrapWebhookEventUnion struct {
	// This field is from variant [VideoTransformationAcceptedEvent],
	// [VideoTransformationReadyEvent], [VideoTransformationErrorEvent],
	// [UploadPreTransformSuccessEvent], [UploadPreTransformErrorEvent],
	// [UploadPostTransformSuccessEvent], [UploadPostTransformErrorEvent].
	ID string `json:"id"`
	// Any of nil, nil, nil, nil, nil, nil, nil.
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	// This field is a union of [VideoTransformationAcceptedEventData],
	// [VideoTransformationReadyEventData], [VideoTransformationErrorEventData],
	// [UploadPreTransformSuccessEventData], [UploadPreTransformErrorEventData],
	// [UploadPostTransformSuccessEventData], [UploadPostTransformErrorEventData]
	Data UnsafeUnwrapWebhookEventUnionData `json:"data"`
	// This field is a union of [VideoTransformationAcceptedEventRequest],
	// [VideoTransformationReadyEventRequest], [VideoTransformationErrorEventRequest],
	// [UploadPreTransformSuccessEventRequest], [UploadPreTransformErrorEventRequest],
	// [UploadPostTransformSuccessEventRequest], [UploadPostTransformErrorEventRequest]
	Request UnsafeUnwrapWebhookEventUnionRequest `json:"request"`
	// This field is from variant [VideoTransformationReadyEvent].
	Timings VideoTransformationReadyEventTimings `json:"timings"`
	JSON    struct {
		ID        respjson.Field
		Type      respjson.Field
		CreatedAt respjson.Field
		Data      respjson.Field
		Request   respjson.Field
		Timings   respjson.Field
		raw       string
	} `json:"-"`
}

func (u UnsafeUnwrapWebhookEventUnion) AsVideoTransformationAcceptedEvent() (v VideoTransformationAcceptedEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsVideoTransformationReadyEvent() (v VideoTransformationReadyEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsVideoTransformationErrorEvent() (v VideoTransformationErrorEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsUploadPreTransformSuccessEvent() (v UploadPreTransformSuccessEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsUploadPreTransformErrorEvent() (v UploadPreTransformErrorEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsUploadPostTransformSuccessEvent() (v UploadPostTransformSuccessEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnsafeUnwrapWebhookEventUnion) AsUploadPostTransformErrorEvent() (v UploadPostTransformErrorEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UnsafeUnwrapWebhookEventUnion) RawJSON() string { return u.JSON.raw }

func (r *UnsafeUnwrapWebhookEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionData is an implicit subunion of
// [UnsafeUnwrapWebhookEventUnion]. UnsafeUnwrapWebhookEventUnionData provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
type UnsafeUnwrapWebhookEventUnionData struct {
	// This field is a union of [VideoTransformationAcceptedEventDataAsset],
	// [VideoTransformationReadyEventDataAsset],
	// [VideoTransformationErrorEventDataAsset]
	Asset UnsafeUnwrapWebhookEventUnionDataAsset `json:"asset"`
	// This field is a union of [VideoTransformationAcceptedEventDataTransformation],
	// [VideoTransformationReadyEventDataTransformation],
	// [VideoTransformationErrorEventDataTransformation],
	// [UploadPreTransformErrorEventDataTransformation],
	// [UploadPostTransformErrorEventDataTransformation]
	Transformation UnsafeUnwrapWebhookEventUnionDataTransformation `json:"transformation"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	AITags []UploadPreTransformSuccessEventDataAITag `json:"AITags"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	AudioCodec string `json:"audioCodec"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	BitRate int64 `json:"bitRate"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	CustomCoordinates string `json:"customCoordinates"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	CustomMetadata map[string]any `json:"customMetadata"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	Description string `json:"description"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	Duration int64 `json:"duration"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	EmbeddedMetadata map[string]any `json:"embeddedMetadata"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	ExtensionStatus UploadPreTransformSuccessEventDataExtensionStatus `json:"extensionStatus"`
	FileID          string                                            `json:"fileId"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	FilePath string `json:"filePath"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	FileType string `json:"fileType"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	Height float64 `json:"height"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	IsPrivateFile bool `json:"isPrivateFile"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	IsPublished bool `json:"isPublished"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	Metadata Metadata `json:"metadata"`
	Name     string   `json:"name"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	SelectedFieldsSchema map[string]UploadPreTransformSuccessEventDataSelectedFieldsSchema `json:"selectedFieldsSchema"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	Size float64 `json:"size"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	Tags []string `json:"tags"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	ThumbnailURL string `json:"thumbnailUrl"`
	URL          string `json:"url"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	VersionInfo UploadPreTransformSuccessEventDataVersionInfo `json:"versionInfo"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	VideoCodec string `json:"videoCodec"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	Width float64 `json:"width"`
	Path  string  `json:"path"`
	JSON  struct {
		Asset                respjson.Field
		Transformation       respjson.Field
		AITags               respjson.Field
		AudioCodec           respjson.Field
		BitRate              respjson.Field
		CustomCoordinates    respjson.Field
		CustomMetadata       respjson.Field
		Description          respjson.Field
		Duration             respjson.Field
		EmbeddedMetadata     respjson.Field
		ExtensionStatus      respjson.Field
		FileID               respjson.Field
		FilePath             respjson.Field
		FileType             respjson.Field
		Height               respjson.Field
		IsPrivateFile        respjson.Field
		IsPublished          respjson.Field
		Metadata             respjson.Field
		Name                 respjson.Field
		SelectedFieldsSchema respjson.Field
		Size                 respjson.Field
		Tags                 respjson.Field
		ThumbnailURL         respjson.Field
		URL                  respjson.Field
		VersionInfo          respjson.Field
		VideoCodec           respjson.Field
		Width                respjson.Field
		Path                 respjson.Field
		raw                  string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionDataAsset is an implicit subunion of
// [UnsafeUnwrapWebhookEventUnion]. UnsafeUnwrapWebhookEventUnionDataAsset provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
type UnsafeUnwrapWebhookEventUnionDataAsset struct {
	URL  string `json:"url"`
	JSON struct {
		URL respjson.Field
		raw string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionDataAsset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionDataTransformation is an implicit subunion of
// [UnsafeUnwrapWebhookEventUnion]. UnsafeUnwrapWebhookEventUnionDataTransformation
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
type UnsafeUnwrapWebhookEventUnionDataTransformation struct {
	Type string `json:"type"`
	// This field is a union of
	// [VideoTransformationAcceptedEventDataTransformationOptions],
	// [VideoTransformationReadyEventDataTransformationOptions],
	// [VideoTransformationErrorEventDataTransformationOptions]
	Options UnsafeUnwrapWebhookEventUnionDataTransformationOptions `json:"options"`
	// This field is from variant [VideoTransformationReadyEventDataTransformation].
	Output VideoTransformationReadyEventDataTransformationOutput `json:"output"`
	// This field is a union of [VideoTransformationErrorEventDataTransformationError],
	// [UploadPreTransformErrorEventDataTransformationError],
	// [UploadPostTransformErrorEventDataTransformationError]
	Error UnsafeUnwrapWebhookEventUnionDataTransformationError `json:"error"`
	JSON  struct {
		Type    respjson.Field
		Options respjson.Field
		Output  respjson.Field
		Error   respjson.Field
		raw     string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionDataTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionDataTransformationOptions is an implicit subunion
// of [UnsafeUnwrapWebhookEventUnion].
// UnsafeUnwrapWebhookEventUnionDataTransformationOptions provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
type UnsafeUnwrapWebhookEventUnionDataTransformationOptions struct {
	AudioCodec     string   `json:"audio_codec"`
	AutoRotate     bool     `json:"auto_rotate"`
	Format         string   `json:"format"`
	Quality        int64    `json:"quality"`
	StreamProtocol string   `json:"stream_protocol"`
	Variants       []string `json:"variants"`
	VideoCodec     string   `json:"video_codec"`
	JSON           struct {
		AudioCodec     respjson.Field
		AutoRotate     respjson.Field
		Format         respjson.Field
		Quality        respjson.Field
		StreamProtocol respjson.Field
		Variants       respjson.Field
		VideoCodec     respjson.Field
		raw            string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionDataTransformationOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionDataTransformationError is an implicit subunion of
// [UnsafeUnwrapWebhookEventUnion].
// UnsafeUnwrapWebhookEventUnionDataTransformationError provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
type UnsafeUnwrapWebhookEventUnionDataTransformationError struct {
	Reason string `json:"reason"`
	JSON   struct {
		Reason respjson.Field
		raw    string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionDataTransformationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionRequest is an implicit subunion of
// [UnsafeUnwrapWebhookEventUnion]. UnsafeUnwrapWebhookEventUnionRequest provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
type UnsafeUnwrapWebhookEventUnionRequest struct {
	URL        string `json:"url"`
	XRequestID string `json:"x_request_id"`
	UserAgent  string `json:"user_agent"`
	// This field is a union of [string], [string],
	// [UploadPostTransformSuccessEventRequestTransformation],
	// [UploadPostTransformErrorEventRequestTransformation]
	Transformation UnsafeUnwrapWebhookEventUnionRequestTransformation `json:"transformation"`
	JSON           struct {
		URL            respjson.Field
		XRequestID     respjson.Field
		UserAgent      respjson.Field
		Transformation respjson.Field
		raw            string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnsafeUnwrapWebhookEventUnionRequestTransformation is an implicit subunion of
// [UnsafeUnwrapWebhookEventUnion].
// UnsafeUnwrapWebhookEventUnionRequestTransformation provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnsafeUnwrapWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type UnsafeUnwrapWebhookEventUnionRequestTransformation struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	Type     string `json:"type"`
	Protocol string `json:"protocol"`
	Value    string `json:"value"`
	JSON     struct {
		OfString respjson.Field
		Type     respjson.Field
		Protocol respjson.Field
		Value    respjson.Field
		raw      string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionRequestTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnion contains all possible properties and values from
// [VideoTransformationAcceptedEvent], [VideoTransformationReadyEvent],
// [VideoTransformationErrorEvent], [UploadPreTransformSuccessEvent],
// [UploadPreTransformErrorEvent], [UploadPostTransformSuccessEvent],
// [UploadPostTransformErrorEvent].
//
// Use the [UnwrapWebhookEventUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UnwrapWebhookEventUnion struct {
	// This field is from variant [VideoTransformationAcceptedEvent],
	// [VideoTransformationReadyEvent], [VideoTransformationErrorEvent],
	// [UploadPreTransformSuccessEvent], [UploadPreTransformErrorEvent],
	// [UploadPostTransformSuccessEvent], [UploadPostTransformErrorEvent].
	ID string `json:"id"`
	// Any of nil, nil, nil, nil, nil, nil, nil.
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	// This field is a union of [VideoTransformationAcceptedEventData],
	// [VideoTransformationReadyEventData], [VideoTransformationErrorEventData],
	// [UploadPreTransformSuccessEventData], [UploadPreTransformErrorEventData],
	// [UploadPostTransformSuccessEventData], [UploadPostTransformErrorEventData]
	Data UnwrapWebhookEventUnionData `json:"data"`
	// This field is a union of [VideoTransformationAcceptedEventRequest],
	// [VideoTransformationReadyEventRequest], [VideoTransformationErrorEventRequest],
	// [UploadPreTransformSuccessEventRequest], [UploadPreTransformErrorEventRequest],
	// [UploadPostTransformSuccessEventRequest], [UploadPostTransformErrorEventRequest]
	Request UnwrapWebhookEventUnionRequest `json:"request"`
	// This field is from variant [VideoTransformationReadyEvent].
	Timings VideoTransformationReadyEventTimings `json:"timings"`
	JSON    struct {
		ID        respjson.Field
		Type      respjson.Field
		CreatedAt respjson.Field
		Data      respjson.Field
		Request   respjson.Field
		Timings   respjson.Field
		raw       string
	} `json:"-"`
}

func (u UnwrapWebhookEventUnion) AsVideoTransformationAcceptedEvent() (v VideoTransformationAcceptedEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsVideoTransformationReadyEvent() (v VideoTransformationReadyEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsVideoTransformationErrorEvent() (v VideoTransformationErrorEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsUploadPreTransformSuccessEvent() (v UploadPreTransformSuccessEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsUploadPreTransformErrorEvent() (v UploadPreTransformErrorEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsUploadPostTransformSuccessEvent() (v UploadPostTransformSuccessEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsUploadPostTransformErrorEvent() (v UploadPostTransformErrorEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UnwrapWebhookEventUnion) RawJSON() string { return u.JSON.raw }

func (r *UnwrapWebhookEventUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionData is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionData provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionData struct {
	// This field is a union of [VideoTransformationAcceptedEventDataAsset],
	// [VideoTransformationReadyEventDataAsset],
	// [VideoTransformationErrorEventDataAsset]
	Asset UnwrapWebhookEventUnionDataAsset `json:"asset"`
	// This field is a union of [VideoTransformationAcceptedEventDataTransformation],
	// [VideoTransformationReadyEventDataTransformation],
	// [VideoTransformationErrorEventDataTransformation],
	// [UploadPreTransformErrorEventDataTransformation],
	// [UploadPostTransformErrorEventDataTransformation]
	Transformation UnwrapWebhookEventUnionDataTransformation `json:"transformation"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	AITags []UploadPreTransformSuccessEventDataAITag `json:"AITags"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	AudioCodec string `json:"audioCodec"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	BitRate int64 `json:"bitRate"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	CustomCoordinates string `json:"customCoordinates"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	CustomMetadata map[string]any `json:"customMetadata"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	Description string `json:"description"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	Duration int64 `json:"duration"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	EmbeddedMetadata map[string]any `json:"embeddedMetadata"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	ExtensionStatus UploadPreTransformSuccessEventDataExtensionStatus `json:"extensionStatus"`
	FileID          string                                            `json:"fileId"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	FilePath string `json:"filePath"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	FileType string `json:"fileType"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	Height float64 `json:"height"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	IsPrivateFile bool `json:"isPrivateFile"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	IsPublished bool `json:"isPublished"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	Metadata Metadata `json:"metadata"`
	Name     string   `json:"name"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	SelectedFieldsSchema map[string]UploadPreTransformSuccessEventDataSelectedFieldsSchema `json:"selectedFieldsSchema"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	Size float64 `json:"size"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	Tags []string `json:"tags"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	ThumbnailURL string `json:"thumbnailUrl"`
	URL          string `json:"url"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	VersionInfo UploadPreTransformSuccessEventDataVersionInfo `json:"versionInfo"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	VideoCodec string `json:"videoCodec"`
	// This field is from variant [UploadPreTransformSuccessEventData].
	Width float64 `json:"width"`
	Path  string  `json:"path"`
	JSON  struct {
		Asset                respjson.Field
		Transformation       respjson.Field
		AITags               respjson.Field
		AudioCodec           respjson.Field
		BitRate              respjson.Field
		CustomCoordinates    respjson.Field
		CustomMetadata       respjson.Field
		Description          respjson.Field
		Duration             respjson.Field
		EmbeddedMetadata     respjson.Field
		ExtensionStatus      respjson.Field
		FileID               respjson.Field
		FilePath             respjson.Field
		FileType             respjson.Field
		Height               respjson.Field
		IsPrivateFile        respjson.Field
		IsPublished          respjson.Field
		Metadata             respjson.Field
		Name                 respjson.Field
		SelectedFieldsSchema respjson.Field
		Size                 respjson.Field
		Tags                 respjson.Field
		ThumbnailURL         respjson.Field
		URL                  respjson.Field
		VersionInfo          respjson.Field
		VideoCodec           respjson.Field
		Width                respjson.Field
		Path                 respjson.Field
		raw                  string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataAsset is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataAsset provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataAsset struct {
	URL  string `json:"url"`
	JSON struct {
		URL respjson.Field
		raw string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataAsset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataTransformation is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataTransformation provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataTransformation struct {
	Type string `json:"type"`
	// This field is a union of
	// [VideoTransformationAcceptedEventDataTransformationOptions],
	// [VideoTransformationReadyEventDataTransformationOptions],
	// [VideoTransformationErrorEventDataTransformationOptions]
	Options UnwrapWebhookEventUnionDataTransformationOptions `json:"options"`
	// This field is from variant [VideoTransformationReadyEventDataTransformation].
	Output VideoTransformationReadyEventDataTransformationOutput `json:"output"`
	// This field is a union of [VideoTransformationErrorEventDataTransformationError],
	// [UploadPreTransformErrorEventDataTransformationError],
	// [UploadPostTransformErrorEventDataTransformationError]
	Error UnwrapWebhookEventUnionDataTransformationError `json:"error"`
	JSON  struct {
		Type    respjson.Field
		Options respjson.Field
		Output  respjson.Field
		Error   respjson.Field
		raw     string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataTransformationOptions is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataTransformationOptions
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataTransformationOptions struct {
	AudioCodec     string   `json:"audio_codec"`
	AutoRotate     bool     `json:"auto_rotate"`
	Format         string   `json:"format"`
	Quality        int64    `json:"quality"`
	StreamProtocol string   `json:"stream_protocol"`
	Variants       []string `json:"variants"`
	VideoCodec     string   `json:"video_codec"`
	JSON           struct {
		AudioCodec     respjson.Field
		AutoRotate     respjson.Field
		Format         respjson.Field
		Quality        respjson.Field
		StreamProtocol respjson.Field
		Variants       respjson.Field
		VideoCodec     respjson.Field
		raw            string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataTransformationOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionDataTransformationError is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionDataTransformationError
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionDataTransformationError struct {
	Reason string `json:"reason"`
	JSON   struct {
		Reason respjson.Field
		raw    string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionDataTransformationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionRequest is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionRequest provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
type UnwrapWebhookEventUnionRequest struct {
	URL        string `json:"url"`
	XRequestID string `json:"x_request_id"`
	UserAgent  string `json:"user_agent"`
	// This field is a union of [string], [string],
	// [UploadPostTransformSuccessEventRequestTransformation],
	// [UploadPostTransformErrorEventRequestTransformation]
	Transformation UnwrapWebhookEventUnionRequestTransformation `json:"transformation"`
	JSON           struct {
		URL            respjson.Field
		XRequestID     respjson.Field
		UserAgent      respjson.Field
		Transformation respjson.Field
		raw            string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnionRequestTransformation is an implicit subunion of
// [UnwrapWebhookEventUnion]. UnwrapWebhookEventUnionRequestTransformation provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [UnwrapWebhookEventUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type UnwrapWebhookEventUnionRequestTransformation struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	Type     string `json:"type"`
	Protocol string `json:"protocol"`
	Value    string `json:"value"`
	JSON     struct {
		OfString respjson.Field
		Type     respjson.Field
		Protocol respjson.Field
		Value    respjson.Field
		raw      string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionRequestTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
