// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/imagekit-developer/imagekit-go/v2/internal/apijson"
	"github.com/imagekit-developer/imagekit-go/v2/internal/apiquery"
	"github.com/imagekit-developer/imagekit-go/v2/internal/requestconfig"
	"github.com/imagekit-developer/imagekit-go/v2/option"
	"github.com/imagekit-developer/imagekit-go/v2/packages/respjson"
)

// AccountUsageAnalyticsService contains methods and other services that help with
// interacting with the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountUsageAnalyticsService] method instead.
type AccountUsageAnalyticsService struct {
	Options []option.RequestOption
}

// NewAccountUsageAnalyticsService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountUsageAnalyticsService(opts ...option.RequestOption) (r AccountUsageAnalyticsService) {
	r = AccountUsageAnalyticsService{}
	r.Options = opts
	return
}

// **Note:** This API is currently in beta.
//
// Get the account analytics data between two dates. The response covers the period
// from the start date to the end date, both dates inclusive. Both dates are
// interpreted as UTC calendar days.
//
// The returned data is scoped to the requesting account only. Unlike
// `/v1/accounts/usage`, an agency account's analytics are not aggregated across
// its child accounts.
//
// The response is cached for 5 minutes per account and date range. Use
// `generatedAt` to check how fresh the returned data is.
func (r *AccountUsageAnalyticsService) Get(ctx context.Context, query AccountUsageAnalyticsGetParams, opts ...option.RequestOption) (res *UsageAnalyticsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/accounts/usage-analytics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type RequestBandwidthEntry struct {
	// Total bandwidth used in bytes.
	BandwidthBytes float64 `json:"bandwidthBytes" api:"required"`
	// Number of requests.
	RequestCount float64 `json:"requestCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BandwidthBytes respjson.Field
		RequestCount   respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RequestBandwidthEntry) RawJSON() string { return r.JSON.raw }
func (r *RequestBandwidthEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponse struct {
	// Total bandwidth, in bytes, utilized during the specified date range.
	BandwidthBytes float64 `json:"bandwidthBytes" api:"required"`
	// CDN traffic grouped by browser.
	Browser UsageAnalyticsResponseBrowser `json:"browser" api:"required"`
	// CDN cache hit, miss and error counts for the date range.
	Cache UsageAnalyticsResponseCache `json:"cache" api:"required"`
	// CDN traffic grouped by country.
	Country UsageAnalyticsResponseCountry `json:"country" api:"required"`
	// CDN traffic grouped by device and operating system (e.g. `Desktop - Apple Mac`,
	// `Smartphone - Apple iPhone`).
	Device UsageAnalyticsResponseDevice `json:"device" api:"required"`
	// End date of the computed analytics data.
	EndDate time.Time `json:"endDate" api:"required" format:"date"`
	// Request count grouped by origin error reason. This covers failed origin fetches,
	// such as an asset not found at origin or an origin timeout. It is not the HTTP
	// status code returned to the client, see `statusCodes` for that.
	ErrorReasons []UsageAnalyticsResponseErrorReason `json:"errorReasons" api:"required"`
	// Raw per-extension operation counts for the date range. These are raw operation
	// counts, not billable extension units. For billable usage, use the
	// `/v1/accounts/usage` endpoint.
	Extensions []UsageAnalyticsResponseExtension `json:"extensions" api:"required"`
	// CDN traffic grouped by response `Content-Type`.
	Format UsageAnalyticsResponseFormat `json:"format" api:"required"`
	// Date and time when the analytics data was computed. Use this to gauge how fresh
	// the returned data is. The date and time is in ISO8601 format.
	GeneratedAt time.Time `json:"generatedAt" api:"required" format:"date-time"`
	// Total number of requests made during the specified date range.
	RequestCount float64 `json:"requestCount" api:"required"`
	// Start date of the computed analytics data.
	StartDate time.Time `json:"startDate" api:"required" format:"date"`
	// Request count grouped by HTTP status code.
	StatusCodes []UsageAnalyticsResponseStatusCode `json:"statusCodes" api:"required"`
	// Top URLs that returned a 404 response.
	Top404Assets []UsageAnalyticsResponseTop404Asset `json:"top404Assets" api:"required"`
	// Top image assets by traffic.
	TopImages UsageAnalyticsResponseTopImages `json:"topImages" api:"required"`
	// Top image transformation strings by traffic.
	TopImageTransforms UsageAnalyticsResponseTopImageTransforms `json:"topImageTransforms" api:"required"`
	// Top non-image, non-video assets by traffic.
	TopOtherAssets UsageAnalyticsResponseTopOtherAssets `json:"topOtherAssets" api:"required"`
	// Top HTTP referrers by traffic.
	TopReferrers UsageAnalyticsResponseTopReferrers `json:"topReferrers" api:"required"`
	// Top user agents by traffic.
	TopUserAgents UsageAnalyticsResponseTopUserAgents `json:"topUserAgents" api:"required"`
	// Top video assets by traffic.
	TopVideos UsageAnalyticsResponseTopVideos `json:"topVideos" api:"required"`
	// Top video transformation strings by traffic.
	TopVideoTransforms UsageAnalyticsResponseTopVideoTransforms `json:"topVideoTransforms" api:"required"`
	// CDN traffic grouped by configured URL endpoint. Traffic that does not match any
	// named URL endpoint pattern is grouped under `Default`.
	URLEndpoints UsageAnalyticsResponseURLEndpoints `json:"urlEndpoints" api:"required"`
	// Raw observed video transcode output duration, in seconds, grouped by resolution
	// and codec. These are raw seconds, not billable Video Processing Units (VPU). For
	// billable VPU totals, use the `/v1/accounts/usage` endpoint.
	VideoProcessing []UsageAnalyticsResponseVideoProcessing `json:"videoProcessing" api:"required"`
	ExtraFields     map[string]any                          `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BandwidthBytes     respjson.Field
		Browser            respjson.Field
		Cache              respjson.Field
		Country            respjson.Field
		Device             respjson.Field
		EndDate            respjson.Field
		ErrorReasons       respjson.Field
		Extensions         respjson.Field
		Format             respjson.Field
		GeneratedAt        respjson.Field
		RequestCount       respjson.Field
		StartDate          respjson.Field
		StatusCodes        respjson.Field
		Top404Assets       respjson.Field
		TopImages          respjson.Field
		TopImageTransforms respjson.Field
		TopOtherAssets     respjson.Field
		TopReferrers       respjson.Field
		TopUserAgents      respjson.Field
		TopVideos          respjson.Field
		TopVideoTransforms respjson.Field
		URLEndpoints       respjson.Field
		VideoProcessing    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponse) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CDN traffic grouped by browser.
type UsageAnalyticsResponseBrowser struct {
	// Top browsers sorted by bandwidth utilized.
	ByBandwidth []UsageAnalyticsResponseBrowserByBandwidth `json:"byBandwidth" api:"required"`
	// Top browsers sorted by request count.
	ByRequests []UsageAnalyticsResponseBrowserByRequest `json:"byRequests" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByBandwidth respjson.Field
		ByRequests  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseBrowser) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseBrowser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponseBrowserByBandwidth struct {
	// Browser name (e.g. `Chrome`).
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	RequestBandwidthEntry
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseBrowserByBandwidth) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseBrowserByBandwidth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponseBrowserByRequest struct {
	// Browser name (e.g. `Chrome`).
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	RequestBandwidthEntry
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseBrowserByRequest) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseBrowserByRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CDN cache hit, miss and error counts for the date range.
type UsageAnalyticsResponseCache struct {
	// Number of requests where the CDN encountered a cache error or exceeded capacity
	// while serving the response.
	ErrorCount float64 `json:"errorCount" api:"required"`
	// Number of requests served from cache, including full hits and revalidated hits.
	HitCount float64 `json:"hitCount" api:"required"`
	// Number of requests that were not found in cache and had to be fetched from
	// origin.
	MissCount float64 `json:"missCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ErrorCount  respjson.Field
		HitCount    respjson.Field
		MissCount   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseCache) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseCache) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CDN traffic grouped by country.
type UsageAnalyticsResponseCountry struct {
	// Top requesting countries sorted by total bandwidth utilized.
	ByBandwidth []UsageAnalyticsResponseCountryByBandwidth `json:"byBandwidth" api:"required"`
	// Top requesting countries sorted by request count.
	ByRequests []UsageAnalyticsResponseCountryByRequest `json:"byRequests" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByBandwidth respjson.Field
		ByRequests  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseCountry) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseCountry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponseCountryByBandwidth struct {
	// ISO country code.
	Code string `json:"code" api:"required"`
	// Country name.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	RequestBandwidthEntry
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseCountryByBandwidth) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseCountryByBandwidth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponseCountryByRequest struct {
	// ISO country code.
	Code string `json:"code" api:"required"`
	// Country name.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	RequestBandwidthEntry
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseCountryByRequest) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseCountryByRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CDN traffic grouped by device and operating system (e.g. `Desktop - Apple Mac`,
// `Smartphone - Apple iPhone`).
type UsageAnalyticsResponseDevice struct {
	// Top device/OS combinations sorted by bandwidth utilized.
	ByBandwidth []UsageAnalyticsResponseDeviceByBandwidth `json:"byBandwidth" api:"required"`
	// Top device/OS combinations sorted by request count.
	ByRequests []UsageAnalyticsResponseDeviceByRequest `json:"byRequests" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByBandwidth respjson.Field
		ByRequests  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseDevice) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseDevice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponseDeviceByBandwidth struct {
	// Device category combined with operating system or vendor (e.g.
	// `Desktop - Windows PC`).
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	RequestBandwidthEntry
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseDeviceByBandwidth) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseDeviceByBandwidth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponseDeviceByRequest struct {
	// Device category combined with operating system or vendor (e.g.
	// `Desktop - Windows PC`).
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	RequestBandwidthEntry
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseDeviceByRequest) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseDeviceByRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponseErrorReason struct {
	// Description of the error reason.
	Name string `json:"name" api:"required"`
	// Number of requests that failed with this error reason.
	RequestCount float64 `json:"requestCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name         respjson.Field
		RequestCount respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseErrorReason) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseErrorReason) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponseExtension struct {
	// Extension identifier.
	Name string `json:"name" api:"required"`
	// Number of times this extension ran during the date range.
	OperationCount float64 `json:"operationCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name           respjson.Field
		OperationCount respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseExtension) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseExtension) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CDN traffic grouped by response `Content-Type`.
type UsageAnalyticsResponseFormat struct {
	// Top content types sorted by bandwidth utilized.
	ByBandwidth []UsageAnalyticsResponseFormatByBandwidth `json:"byBandwidth" api:"required"`
	// Top content types sorted by request count.
	ByRequests []UsageAnalyticsResponseFormatByRequest `json:"byRequests" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByBandwidth respjson.Field
		ByRequests  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseFormat) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseFormat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponseFormatByBandwidth struct {
	// MIME type (e.g. `image/webp`).
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	RequestBandwidthEntry
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseFormatByBandwidth) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseFormatByBandwidth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponseFormatByRequest struct {
	// MIME type (e.g. `image/webp`).
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	RequestBandwidthEntry
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseFormatByRequest) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseFormatByRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponseStatusCode struct {
	// HTTP status code.
	Name string `json:"name" api:"required"`
	// Number of requests that received this status code.
	RequestCount float64 `json:"requestCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name         respjson.Field
		RequestCount respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseStatusCode) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseStatusCode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponseTop404Asset struct {
	// URL that returned a 404 response.
	Name string `json:"name" api:"required"`
	// Number of requests to this URL that returned a 404 response.
	RequestCount float64 `json:"requestCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name         respjson.Field
		RequestCount respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseTop404Asset) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseTop404Asset) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Top image assets by traffic.
type UsageAnalyticsResponseTopImages struct {
	// Top image assets sorted by bandwidth utilized.
	ByBandwidth []UsageAnalyticsResponseTopImagesByBandwidth `json:"byBandwidth" api:"required"`
	// Top image assets sorted by request count.
	ByRequests []UsageAnalyticsResponseTopImagesByRequest `json:"byRequests" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByBandwidth respjson.Field
		ByRequests  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseTopImages) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseTopImages) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponseTopImagesByBandwidth struct {
	// URL of the image asset.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	RequestBandwidthEntry
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseTopImagesByBandwidth) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseTopImagesByBandwidth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponseTopImagesByRequest struct {
	// URL of the image asset.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	RequestBandwidthEntry
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseTopImagesByRequest) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseTopImagesByRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Top image transformation strings by traffic.
type UsageAnalyticsResponseTopImageTransforms struct {
	// Top image transformation strings sorted by bandwidth utilized.
	ByBandwidth []UsageAnalyticsResponseTopImageTransformsByBandwidth `json:"byBandwidth" api:"required"`
	// Top image transformation strings sorted by request count.
	ByRequests []UsageAnalyticsResponseTopImageTransformsByRequest `json:"byRequests" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByBandwidth respjson.Field
		ByRequests  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseTopImageTransforms) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseTopImageTransforms) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponseTopImageTransformsByBandwidth struct {
	// Image transformation string (e.g. `tr:w-400,h-400`).
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	RequestBandwidthEntry
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseTopImageTransformsByBandwidth) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseTopImageTransformsByBandwidth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponseTopImageTransformsByRequest struct {
	// Image transformation string (e.g. `tr:w-400,h-400`).
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	RequestBandwidthEntry
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseTopImageTransformsByRequest) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseTopImageTransformsByRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Top non-image, non-video assets by traffic.
type UsageAnalyticsResponseTopOtherAssets struct {
	// Top non-image, non-video assets sorted by bandwidth utilized.
	ByBandwidth []UsageAnalyticsResponseTopOtherAssetsByBandwidth `json:"byBandwidth" api:"required"`
	// Top non-image, non-video assets sorted by request count.
	ByRequests []UsageAnalyticsResponseTopOtherAssetsByRequest `json:"byRequests" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByBandwidth respjson.Field
		ByRequests  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseTopOtherAssets) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseTopOtherAssets) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponseTopOtherAssetsByBandwidth struct {
	// URL of the non-image, non-video asset.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	RequestBandwidthEntry
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseTopOtherAssetsByBandwidth) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseTopOtherAssetsByBandwidth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponseTopOtherAssetsByRequest struct {
	// URL of the non-image, non-video asset.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	RequestBandwidthEntry
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseTopOtherAssetsByRequest) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseTopOtherAssetsByRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Top HTTP referrers by traffic.
type UsageAnalyticsResponseTopReferrers struct {
	// Top HTTP referrers sorted by bandwidth utilized.
	ByBandwidth []UsageAnalyticsResponseTopReferrersByBandwidth `json:"byBandwidth" api:"required"`
	// Top HTTP referrers sorted by request count.
	ByRequests []UsageAnalyticsResponseTopReferrersByRequest `json:"byRequests" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByBandwidth respjson.Field
		ByRequests  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseTopReferrers) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseTopReferrers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponseTopReferrersByBandwidth struct {
	// Referrer URL.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	RequestBandwidthEntry
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseTopReferrersByBandwidth) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseTopReferrersByBandwidth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponseTopReferrersByRequest struct {
	// Referrer URL.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	RequestBandwidthEntry
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseTopReferrersByRequest) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseTopReferrersByRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Top user agents by traffic.
type UsageAnalyticsResponseTopUserAgents struct {
	// Top user agents sorted by bandwidth utilized.
	ByBandwidth []UsageAnalyticsResponseTopUserAgentsByBandwidth `json:"byBandwidth" api:"required"`
	// Top user agents sorted by request count.
	ByRequests []UsageAnalyticsResponseTopUserAgentsByRequest `json:"byRequests" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByBandwidth respjson.Field
		ByRequests  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseTopUserAgents) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseTopUserAgents) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponseTopUserAgentsByBandwidth struct {
	// User agent string.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	RequestBandwidthEntry
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseTopUserAgentsByBandwidth) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseTopUserAgentsByBandwidth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponseTopUserAgentsByRequest struct {
	// User agent string.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	RequestBandwidthEntry
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseTopUserAgentsByRequest) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseTopUserAgentsByRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Top video assets by traffic.
type UsageAnalyticsResponseTopVideos struct {
	// Top video assets sorted by bandwidth utilized.
	ByBandwidth []UsageAnalyticsResponseTopVideosByBandwidth `json:"byBandwidth" api:"required"`
	// Top video assets sorted by request count.
	ByRequests []UsageAnalyticsResponseTopVideosByRequest `json:"byRequests" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByBandwidth respjson.Field
		ByRequests  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseTopVideos) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseTopVideos) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponseTopVideosByBandwidth struct {
	// URL of the video asset.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	RequestBandwidthEntry
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseTopVideosByBandwidth) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseTopVideosByBandwidth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponseTopVideosByRequest struct {
	// Full URL of the video asset (e.g. `https://ik.imagekit.io/demo/clip.mp4`).
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	RequestBandwidthEntry
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseTopVideosByRequest) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseTopVideosByRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Top video transformation strings by traffic.
type UsageAnalyticsResponseTopVideoTransforms struct {
	// Top video transformation strings sorted by bandwidth utilized.
	ByBandwidth []UsageAnalyticsResponseTopVideoTransformsByBandwidth `json:"byBandwidth" api:"required"`
	// Top video transformation strings sorted by request count.
	ByRequests []UsageAnalyticsResponseTopVideoTransformsByRequest `json:"byRequests" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByBandwidth respjson.Field
		ByRequests  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseTopVideoTransforms) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseTopVideoTransforms) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponseTopVideoTransformsByBandwidth struct {
	// Video transformation string (e.g. `tr:h-720,f-mp4`).
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	RequestBandwidthEntry
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseTopVideoTransformsByBandwidth) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseTopVideoTransformsByBandwidth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponseTopVideoTransformsByRequest struct {
	// Video transformation string (e.g. `tr:h-720,f-mp4`).
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	RequestBandwidthEntry
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseTopVideoTransformsByRequest) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseTopVideoTransformsByRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CDN traffic grouped by configured URL endpoint. Traffic that does not match any
// named URL endpoint pattern is grouped under `Default`.
type UsageAnalyticsResponseURLEndpoints struct {
	// Top URL endpoints sorted by bandwidth utilized.
	ByBandwidth []UsageAnalyticsResponseURLEndpointsByBandwidth `json:"byBandwidth" api:"required"`
	// Top URL endpoints sorted by request count.
	ByRequests []UsageAnalyticsResponseURLEndpointsByRequest `json:"byRequests" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ByBandwidth respjson.Field
		ByRequests  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseURLEndpoints) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseURLEndpoints) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponseURLEndpointsByBandwidth struct {
	// URL endpoint name, or `Default` for traffic that does not match a named
	// endpoint.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	RequestBandwidthEntry
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseURLEndpointsByBandwidth) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseURLEndpointsByBandwidth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponseURLEndpointsByRequest struct {
	// URL endpoint name, or `Default` for traffic that does not match a named
	// endpoint.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	RequestBandwidthEntry
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseURLEndpointsByRequest) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseURLEndpointsByRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageAnalyticsResponseVideoProcessing struct {
	// Video codec used for the output (e.g. `h264`, `av1`).
	Codec string `json:"codec" api:"required"`
	// Total output duration, in seconds, for this resolution and codec combination.
	DurationSeconds float64 `json:"durationSeconds" api:"required"`
	// Output resolution tier (e.g. `SD`, `HD`, `4K`).
	Resolution string `json:"resolution" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Codec           respjson.Field
		DurationSeconds respjson.Field
		Resolution      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageAnalyticsResponseVideoProcessing) RawJSON() string { return r.JSON.raw }
func (r *UsageAnalyticsResponseVideoProcessing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUsageAnalyticsGetParams struct {
	// Specify an `endDate` in `YYYY-MM-DD` format, interpreted as a UTC calendar day.
	// It should be after the `startDate`. The difference between `startDate` and
	// `endDate` should be less than 90 days.
	EndDate time.Time `query:"endDate" api:"required" format:"date" json:"-"`
	// Specify a `startDate` in `YYYY-MM-DD` format, interpreted as a UTC calendar day.
	// It should be before the `endDate`. The difference between `startDate` and
	// `endDate` should be less than 90 days.
	StartDate time.Time `query:"startDate" api:"required" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [AccountUsageAnalyticsGetParams]'s query parameters as
// `url.Values`.
func (r AccountUsageAnalyticsGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
