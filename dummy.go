// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"context"
	"net/http"
	"slices"

	"github.com/imagekit-developer/imagekit-go/internal/apijson"
	"github.com/imagekit-developer/imagekit-go/internal/requestconfig"
	"github.com/imagekit-developer/imagekit-go/option"
	"github.com/imagekit-developer/imagekit-go/packages/param"
	"github.com/imagekit-developer/imagekit-go/shared"
)

// DummyService contains methods and other services that help with interacting with
// the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDummyService] method instead.
type DummyService struct {
	Options []option.RequestOption
}

// NewDummyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDummyService(opts ...option.RequestOption) (r DummyService) {
	r = DummyService{}
	r.Options = opts
	return
}

// Internal test endpoint for SDK generation purposes only. This endpoint
// demonstrates usage of all shared models defined in the Stainless configuration
// and is not intended for public consumption.
func (r *DummyService) New(ctx context.Context, body DummyNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/dummy/test"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type DummyNewParams struct {
	BaseOverlay shared.BaseOverlayParam `json:"baseOverlay,omitzero"`
	// Array of extensions to be applied to the asset. Each extension can be configured
	// with specific parameters based on the extension type.
	Extensions   shared.ExtensionsParam   `json:"extensions,omitzero"`
	ImageOverlay shared.ImageOverlayParam `json:"imageOverlay,omitzero"`
	// Specifies an overlay to be applied on the parent image or video. ImageKit
	// supports overlays including images, text, videos, subtitles, and solid colors.
	// See
	// [Overlay using layers](https://imagekit.io/docs/transformations#overlay-using-layers).
	Overlay                         shared.OverlayUnionParam                    `json:"overlay,omitzero"`
	OverlayPosition                 shared.OverlayPositionParam                 `json:"overlayPosition,omitzero"`
	OverlayTiming                   shared.OverlayTimingParam                   `json:"overlayTiming,omitzero"`
	SolidColorOverlay               shared.SolidColorOverlayParam               `json:"solidColorOverlay,omitzero"`
	SolidColorOverlayTransformation shared.SolidColorOverlayTransformationParam `json:"solidColorOverlayTransformation,omitzero"`
	// Options for generating ImageKit URLs with transformations. See the
	// [Transformations guide](https://imagekit.io/docs/transformations).
	SrcOptions shared.SrcOptionsParam `json:"srcOptions,omitzero"`
	// Available streaming resolutions for
	// [adaptive bitrate streaming](https://imagekit.io/docs/adaptive-bitrate-streaming)
	//
	// Any of "240", "360", "480", "720", "1080", "1440", "2160".
	StreamingResolution shared.StreamingResolution  `json:"streamingResolution,omitzero"`
	SubtitleOverlay     shared.SubtitleOverlayParam `json:"subtitleOverlay,omitzero"`
	// Subtitle styling options.
	// [Learn more](https://imagekit.io/docs/add-overlays-on-videos#styling-controls-for-subtitles-layer)
	// from the docs.
	SubtitleOverlayTransformation shared.SubtitleOverlayTransformationParam `json:"subtitleOverlayTransformation,omitzero"`
	TextOverlay                   shared.TextOverlayParam                   `json:"textOverlay,omitzero"`
	TextOverlayTransformation     shared.TextOverlayTransformationParam     `json:"textOverlayTransformation,omitzero"`
	// The SDK provides easy-to-use names for transformations. These names are
	// converted to the corresponding transformation string before being added to the
	// URL. SDKs are updated regularly to support new transformations. If you want to
	// use a transformation that is not supported by the SDK, You can use the `raw`
	// parameter to pass the transformation string directly. See the
	// [Transformations documentation](https://imagekit.io/docs/transformations).
	Transformation shared.TransformationParam `json:"transformation,omitzero"`
	// By default, the transformation string is added as a query parameter in the URL,
	// e.g., `?tr=w-100,h-100`. If you want to add the transformation string in the
	// path of the URL, set this to `path`. Learn more in the
	// [Transformations guide](https://imagekit.io/docs/transformations).
	//
	// Any of "path", "query".
	TransformationPosition shared.TransformationPosition `json:"transformationPosition,omitzero"`
	VideoOverlay           shared.VideoOverlayParam      `json:"videoOverlay,omitzero"`
	paramObj
}

func (r DummyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DummyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DummyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
