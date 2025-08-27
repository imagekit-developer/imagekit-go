// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/stainless-sdks/imagekit-go/internal/apijson"
	"github.com/stainless-sdks/imagekit-go/internal/requestconfig"
	"github.com/stainless-sdks/imagekit-go/option"
	"github.com/stainless-sdks/imagekit-go/packages/respjson"
	"github.com/stainless-sdks/imagekit-go/shared/constant"
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
	opts = append(opts, r.Options...)
	cfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return nil, err
	}
	key := cfg.PrivateAPIKey
	if key == "" {
		return nil, errors.New("The PrivateAPIKey option must be set in order to verify webhook headers")
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

type VideoTransformationAcceptedEvent struct {
	// Unique identifier for the event.
	ID        string                                  `json:"id,required"`
	CreatedAt time.Time                               `json:"created_at,required" format:"date-time"`
	Data      VideoTransformationAcceptedEventData    `json:"data,required"`
	Request   VideoTransformationAcceptedEventRequest `json:"request,required"`
	Type      constant.VideoTransformationAccepted    `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Data        respjson.Field
		Request     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationAcceptedEvent) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationAcceptedEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationAcceptedEventData struct {
	Asset          VideoTransformationAcceptedEventDataAsset          `json:"asset,required"`
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

type VideoTransformationAcceptedEventDataAsset struct {
	// Source asset URL.
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

type VideoTransformationAcceptedEventDataTransformation struct {
	// Any of "video-transformation", "gif-to-video", "video-thumbnail".
	Type    string                                                    `json:"type,required"`
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

type VideoTransformationAcceptedEventDataTransformationOptions struct {
	// Any of "aac", "opus".
	AudioCodec string `json:"audio_codec"`
	AutoRotate bool   `json:"auto_rotate"`
	// Any of "mp4", "webm", "jpg", "png", "webp".
	Format  string `json:"format"`
	Quality int64  `json:"quality"`
	// Any of "HLS", "DASH".
	StreamProtocol string   `json:"stream_protocol"`
	Variants       []string `json:"variants"`
	// Any of "h264", "vp9".
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

type VideoTransformationAcceptedEventRequest struct {
	// URL of the submitted request.
	URL string `json:"url,required" format:"uri"`
	// Unique ID for the originating request.
	XRequestID string `json:"x_request_id,required"`
	// User-Agent header of the originating request.
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

type VideoTransformationErrorEvent struct {
	// Unique identifier for the event.
	ID        string                               `json:"id,required"`
	CreatedAt time.Time                            `json:"created_at,required" format:"date-time"`
	Data      VideoTransformationErrorEventData    `json:"data,required"`
	Request   VideoTransformationErrorEventRequest `json:"request,required"`
	Type      constant.VideoTransformationError    `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Data        respjson.Field
		Request     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationErrorEvent) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationErrorEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationErrorEventData struct {
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

type VideoTransformationErrorEventDataAsset struct {
	// Source asset URL.
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
	// Any of "video-transformation", "gif-to-video", "video-thumbnail".
	Type    string                                                 `json:"type,required"`
	Error   VideoTransformationErrorEventDataTransformationError   `json:"error"`
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

type VideoTransformationErrorEventDataTransformationError struct {
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

type VideoTransformationErrorEventDataTransformationOptions struct {
	// Any of "aac", "opus".
	AudioCodec string `json:"audio_codec"`
	AutoRotate bool   `json:"auto_rotate"`
	// Any of "mp4", "webm", "jpg", "png", "webp".
	Format  string `json:"format"`
	Quality int64  `json:"quality"`
	// Any of "HLS", "DASH".
	StreamProtocol string   `json:"stream_protocol"`
	Variants       []string `json:"variants"`
	// Any of "h264", "vp9".
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

type VideoTransformationErrorEventRequest struct {
	// URL of the submitted request.
	URL string `json:"url,required" format:"uri"`
	// Unique ID for the originating request.
	XRequestID string `json:"x_request_id,required"`
	// User-Agent header of the originating request.
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

type VideoTransformationReadyEvent struct {
	// Unique identifier for the event.
	ID        string                               `json:"id,required"`
	CreatedAt time.Time                            `json:"created_at,required" format:"date-time"`
	Data      VideoTransformationReadyEventData    `json:"data,required"`
	Request   VideoTransformationReadyEventRequest `json:"request,required"`
	Type      constant.VideoTransformationReady    `json:"type,required"`
	Timings   VideoTransformationReadyEventTimings `json:"timings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Data        respjson.Field
		Request     respjson.Field
		Type        respjson.Field
		Timings     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationReadyEvent) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationReadyEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationReadyEventData struct {
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

type VideoTransformationReadyEventDataAsset struct {
	// Source asset URL.
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
	// Any of "video-transformation", "gif-to-video", "video-thumbnail".
	Type    string                                                 `json:"type,required"`
	Options VideoTransformationReadyEventDataTransformationOptions `json:"options"`
	Output  VideoTransformationReadyEventDataTransformationOutput  `json:"output"`
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

type VideoTransformationReadyEventDataTransformationOptions struct {
	// Any of "aac", "opus".
	AudioCodec string `json:"audio_codec"`
	AutoRotate bool   `json:"auto_rotate"`
	// Any of "mp4", "webm", "jpg", "png", "webp".
	Format  string `json:"format"`
	Quality int64  `json:"quality"`
	// Any of "HLS", "DASH".
	StreamProtocol string   `json:"stream_protocol"`
	Variants       []string `json:"variants"`
	// Any of "h264", "vp9".
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

type VideoTransformationReadyEventDataTransformationOutput struct {
	URL           string                                                             `json:"url,required" format:"uri"`
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

type VideoTransformationReadyEventDataTransformationOutputVideoMetadata struct {
	Bitrate  int64   `json:"bitrate,required"`
	Duration float64 `json:"duration,required"`
	Height   int64   `json:"height,required"`
	Width    int64   `json:"width,required"`
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

type VideoTransformationReadyEventRequest struct {
	// URL of the submitted request.
	URL string `json:"url,required" format:"uri"`
	// Unique ID for the originating request.
	XRequestID string `json:"x_request_id,required"`
	// User-Agent header of the originating request.
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

type VideoTransformationReadyEventTimings struct {
	// Milliseconds spent downloading the source.
	DownloadDuration int64 `json:"download_duration"`
	// Milliseconds spent encoding.
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
// [VideoTransformationErrorEvent].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UnsafeUnwrapWebhookEventUnion struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	// This field is a union of [VideoTransformationAcceptedEventData],
	// [VideoTransformationReadyEventData], [VideoTransformationErrorEventData]
	Data UnsafeUnwrapWebhookEventUnionData `json:"data"`
	// This field is a union of [VideoTransformationAcceptedEventRequest],
	// [VideoTransformationReadyEventRequest], [VideoTransformationErrorEventRequest]
	Request UnsafeUnwrapWebhookEventUnionRequest `json:"request"`
	Type    string                               `json:"type"`
	// This field is from variant [VideoTransformationReadyEvent].
	Timings VideoTransformationReadyEventTimings `json:"timings"`
	JSON    struct {
		ID        respjson.Field
		CreatedAt respjson.Field
		Data      respjson.Field
		Request   respjson.Field
		Type      respjson.Field
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
	// [VideoTransformationErrorEventDataTransformation]
	Transformation UnsafeUnwrapWebhookEventUnionDataTransformation `json:"transformation"`
	JSON           struct {
		Asset          respjson.Field
		Transformation respjson.Field
		raw            string
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
	// This field is from variant [VideoTransformationErrorEventDataTransformation].
	Error VideoTransformationErrorEventDataTransformationError `json:"error"`
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
	JSON       struct {
		URL        respjson.Field
		XRequestID respjson.Field
		UserAgent  respjson.Field
		raw        string
	} `json:"-"`
}

func (r *UnsafeUnwrapWebhookEventUnionRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// UnwrapWebhookEventUnion contains all possible properties and values from
// [VideoTransformationAcceptedEvent], [VideoTransformationReadyEvent],
// [VideoTransformationErrorEvent].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UnwrapWebhookEventUnion struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	// This field is a union of [VideoTransformationAcceptedEventData],
	// [VideoTransformationReadyEventData], [VideoTransformationErrorEventData]
	Data UnwrapWebhookEventUnionData `json:"data"`
	// This field is a union of [VideoTransformationAcceptedEventRequest],
	// [VideoTransformationReadyEventRequest], [VideoTransformationErrorEventRequest]
	Request UnwrapWebhookEventUnionRequest `json:"request"`
	Type    string                         `json:"type"`
	// This field is from variant [VideoTransformationReadyEvent].
	Timings VideoTransformationReadyEventTimings `json:"timings"`
	JSON    struct {
		ID        respjson.Field
		CreatedAt respjson.Field
		Data      respjson.Field
		Request   respjson.Field
		Type      respjson.Field
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
	// [VideoTransformationErrorEventDataTransformation]
	Transformation UnwrapWebhookEventUnionDataTransformation `json:"transformation"`
	JSON           struct {
		Asset          respjson.Field
		Transformation respjson.Field
		raw            string
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
	// This field is from variant [VideoTransformationErrorEventDataTransformation].
	Error VideoTransformationErrorEventDataTransformationError `json:"error"`
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
	JSON       struct {
		URL        respjson.Field
		XRequestID respjson.Field
		UserAgent  respjson.Field
		raw        string
	} `json:"-"`
}

func (r *UnwrapWebhookEventUnionRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
