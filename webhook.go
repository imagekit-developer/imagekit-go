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

type VideoTransformationAcceptedWebhookEvent struct {
	// Unique identifier for the event.
	ID        string                                         `json:"id,required"`
	CreatedAt time.Time                                      `json:"created_at,required" format:"date-time"`
	Data      VideoTransformationAcceptedWebhookEventData    `json:"data,required"`
	Request   VideoTransformationAcceptedWebhookEventRequest `json:"request,required"`
	// Any of "video.transformation.accepted".
	Type VideoTransformationAcceptedWebhookEventType `json:"type,required"`
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
func (r VideoTransformationAcceptedWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationAcceptedWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationAcceptedWebhookEventData struct {
	Asset          VideoTransformationAcceptedWebhookEventDataAsset          `json:"asset,required"`
	Transformation VideoTransformationAcceptedWebhookEventDataTransformation `json:"transformation,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Asset          respjson.Field
		Transformation respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationAcceptedWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationAcceptedWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationAcceptedWebhookEventDataAsset struct {
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
func (r VideoTransformationAcceptedWebhookEventDataAsset) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationAcceptedWebhookEventDataAsset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationAcceptedWebhookEventDataTransformation struct {
	// Any of "video-transformation", "gif-to-video", "video-thumbnail".
	Type    string                                                           `json:"type,required"`
	Options VideoTransformationAcceptedWebhookEventDataTransformationOptions `json:"options"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Options     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationAcceptedWebhookEventDataTransformation) RawJSON() string {
	return r.JSON.raw
}
func (r *VideoTransformationAcceptedWebhookEventDataTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationAcceptedWebhookEventDataTransformationOptions struct {
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
func (r VideoTransformationAcceptedWebhookEventDataTransformationOptions) RawJSON() string {
	return r.JSON.raw
}
func (r *VideoTransformationAcceptedWebhookEventDataTransformationOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationAcceptedWebhookEventRequest struct {
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
func (r VideoTransformationAcceptedWebhookEventRequest) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationAcceptedWebhookEventRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationAcceptedWebhookEventType string

const (
	VideoTransformationAcceptedWebhookEventTypeVideoTransformationAccepted VideoTransformationAcceptedWebhookEventType = "video.transformation.accepted"
)

type VideoTransformationReadyWebhookEvent struct {
	// Unique identifier for the event.
	ID        string                                      `json:"id,required"`
	CreatedAt time.Time                                   `json:"created_at,required" format:"date-time"`
	Data      VideoTransformationReadyWebhookEventData    `json:"data,required"`
	Request   VideoTransformationReadyWebhookEventRequest `json:"request,required"`
	// Any of "video.transformation.ready".
	Type    VideoTransformationReadyWebhookEventType    `json:"type,required"`
	Timings VideoTransformationReadyWebhookEventTimings `json:"timings"`
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
func (r VideoTransformationReadyWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationReadyWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationReadyWebhookEventData struct {
	Asset          VideoTransformationReadyWebhookEventDataAsset          `json:"asset,required"`
	Transformation VideoTransformationReadyWebhookEventDataTransformation `json:"transformation,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Asset          respjson.Field
		Transformation respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationReadyWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationReadyWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationReadyWebhookEventDataAsset struct {
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
func (r VideoTransformationReadyWebhookEventDataAsset) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationReadyWebhookEventDataAsset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationReadyWebhookEventDataTransformation struct {
	// Any of "video-transformation", "gif-to-video", "video-thumbnail".
	Type    string                                                        `json:"type,required"`
	Options VideoTransformationReadyWebhookEventDataTransformationOptions `json:"options"`
	Output  VideoTransformationReadyWebhookEventDataTransformationOutput  `json:"output"`
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
func (r VideoTransformationReadyWebhookEventDataTransformation) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationReadyWebhookEventDataTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationReadyWebhookEventDataTransformationOptions struct {
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
func (r VideoTransformationReadyWebhookEventDataTransformationOptions) RawJSON() string {
	return r.JSON.raw
}
func (r *VideoTransformationReadyWebhookEventDataTransformationOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationReadyWebhookEventDataTransformationOutput struct {
	URL           string                                                                    `json:"url,required" format:"uri"`
	VideoMetadata VideoTransformationReadyWebhookEventDataTransformationOutputVideoMetadata `json:"video_metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL           respjson.Field
		VideoMetadata respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationReadyWebhookEventDataTransformationOutput) RawJSON() string {
	return r.JSON.raw
}
func (r *VideoTransformationReadyWebhookEventDataTransformationOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationReadyWebhookEventDataTransformationOutputVideoMetadata struct {
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
func (r VideoTransformationReadyWebhookEventDataTransformationOutputVideoMetadata) RawJSON() string {
	return r.JSON.raw
}
func (r *VideoTransformationReadyWebhookEventDataTransformationOutputVideoMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationReadyWebhookEventRequest struct {
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
func (r VideoTransformationReadyWebhookEventRequest) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationReadyWebhookEventRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationReadyWebhookEventType string

const (
	VideoTransformationReadyWebhookEventTypeVideoTransformationReady VideoTransformationReadyWebhookEventType = "video.transformation.ready"
)

type VideoTransformationReadyWebhookEventTimings struct {
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
func (r VideoTransformationReadyWebhookEventTimings) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationReadyWebhookEventTimings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationErrorWebhookEvent struct {
	// Unique identifier for the event.
	ID        string                                      `json:"id,required"`
	CreatedAt time.Time                                   `json:"created_at,required" format:"date-time"`
	Data      VideoTransformationErrorWebhookEventData    `json:"data,required"`
	Request   VideoTransformationErrorWebhookEventRequest `json:"request,required"`
	// Any of "video.transformation.error".
	Type VideoTransformationErrorWebhookEventType `json:"type,required"`
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
func (r VideoTransformationErrorWebhookEvent) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationErrorWebhookEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationErrorWebhookEventData struct {
	Asset          VideoTransformationErrorWebhookEventDataAsset          `json:"asset,required"`
	Transformation VideoTransformationErrorWebhookEventDataTransformation `json:"transformation,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Asset          respjson.Field
		Transformation respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VideoTransformationErrorWebhookEventData) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationErrorWebhookEventData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationErrorWebhookEventDataAsset struct {
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
func (r VideoTransformationErrorWebhookEventDataAsset) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationErrorWebhookEventDataAsset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationErrorWebhookEventDataTransformation struct {
	// Any of "video-transformation", "gif-to-video", "video-thumbnail".
	Type    string                                                        `json:"type,required"`
	Error   VideoTransformationErrorWebhookEventDataTransformationError   `json:"error"`
	Options VideoTransformationErrorWebhookEventDataTransformationOptions `json:"options"`
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
func (r VideoTransformationErrorWebhookEventDataTransformation) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationErrorWebhookEventDataTransformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationErrorWebhookEventDataTransformationError struct {
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
func (r VideoTransformationErrorWebhookEventDataTransformationError) RawJSON() string {
	return r.JSON.raw
}
func (r *VideoTransformationErrorWebhookEventDataTransformationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationErrorWebhookEventDataTransformationOptions struct {
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
func (r VideoTransformationErrorWebhookEventDataTransformationOptions) RawJSON() string {
	return r.JSON.raw
}
func (r *VideoTransformationErrorWebhookEventDataTransformationOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationErrorWebhookEventRequest struct {
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
func (r VideoTransformationErrorWebhookEventRequest) RawJSON() string { return r.JSON.raw }
func (r *VideoTransformationErrorWebhookEventRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VideoTransformationErrorWebhookEventType string

const (
	VideoTransformationErrorWebhookEventTypeVideoTransformationError VideoTransformationErrorWebhookEventType = "video.transformation.error"
)

// UnwrapWebhookEventUnion contains all possible properties and values from
// [VideoTransformationAcceptedWebhookEvent],
// [VideoTransformationReadyWebhookEvent], [VideoTransformationErrorWebhookEvent].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UnwrapWebhookEventUnion struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	// This field is a union of [VideoTransformationAcceptedWebhookEventData],
	// [VideoTransformationReadyWebhookEventData],
	// [VideoTransformationErrorWebhookEventData]
	Data UnwrapWebhookEventUnionData `json:"data"`
	// This field is a union of [VideoTransformationAcceptedWebhookEventRequest],
	// [VideoTransformationReadyWebhookEventRequest],
	// [VideoTransformationErrorWebhookEventRequest]
	Request UnwrapWebhookEventUnionRequest `json:"request"`
	Type    string                         `json:"type"`
	// This field is from variant [VideoTransformationReadyWebhookEvent].
	Timings VideoTransformationReadyWebhookEventTimings `json:"timings"`
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

func (u UnwrapWebhookEventUnion) AsVideoTransformationAcceptedWebhookEvent() (v VideoTransformationAcceptedWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsVideoTransformationReadyWebhookEvent() (v VideoTransformationReadyWebhookEvent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnwrapWebhookEventUnion) AsVideoTransformationErrorWebhookEvent() (v VideoTransformationErrorWebhookEvent) {
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
	// This field is a union of [VideoTransformationAcceptedWebhookEventDataAsset],
	// [VideoTransformationReadyWebhookEventDataAsset],
	// [VideoTransformationErrorWebhookEventDataAsset]
	Asset UnwrapWebhookEventUnionDataAsset `json:"asset"`
	// This field is a union of
	// [VideoTransformationAcceptedWebhookEventDataTransformation],
	// [VideoTransformationReadyWebhookEventDataTransformation],
	// [VideoTransformationErrorWebhookEventDataTransformation]
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
	// [VideoTransformationAcceptedWebhookEventDataTransformationOptions],
	// [VideoTransformationReadyWebhookEventDataTransformationOptions],
	// [VideoTransformationErrorWebhookEventDataTransformationOptions]
	Options UnwrapWebhookEventUnionDataTransformationOptions `json:"options"`
	// This field is from variant
	// [VideoTransformationReadyWebhookEventDataTransformation].
	Output VideoTransformationReadyWebhookEventDataTransformationOutput `json:"output"`
	// This field is from variant
	// [VideoTransformationErrorWebhookEventDataTransformation].
	Error VideoTransformationErrorWebhookEventDataTransformationError `json:"error"`
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
