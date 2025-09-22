// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/imagekit-go/internal/apijson"
	"github.com/stainless-sdks/imagekit-go/internal/apiquery"
	"github.com/stainless-sdks/imagekit-go/internal/requestconfig"
	"github.com/stainless-sdks/imagekit-go/option"
	"github.com/stainless-sdks/imagekit-go/packages/respjson"
)

// AccountUsageService contains methods and other services that help with
// interacting with the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountUsageService] method instead.
type AccountUsageService struct {
	Options []option.RequestOption
}

// NewAccountUsageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountUsageService(opts ...option.RequestOption) (r AccountUsageService) {
	r = AccountUsageService{}
	r.Options = opts
	return
}

// Get the account usage information between two dates. Note that the API response
// includes data from the start date while excluding data from the end date. In
// other words, the data covers the period starting from the specified start date
// up to, but not including, the end date.
func (r *AccountUsageService) Get(ctx context.Context, query AccountUsageGetParams, opts ...option.RequestOption) (res *AccountUsageGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/accounts/usage"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountUsageGetResponse struct {
	// Amount of bandwidth used in bytes.
	BandwidthBytes int64 `json:"bandwidthBytes"`
	// Number of extension units used.
	ExtensionUnitsCount int64 `json:"extensionUnitsCount"`
	// Storage used by media library in bytes.
	MediaLibraryStorageBytes int64 `json:"mediaLibraryStorageBytes"`
	// Storage used by the original cache in bytes.
	OriginalCacheStorageBytes int64 `json:"originalCacheStorageBytes"`
	// Number of video processing units used.
	VideoProcessingUnitsCount int64 `json:"videoProcessingUnitsCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BandwidthBytes            respjson.Field
		ExtensionUnitsCount       respjson.Field
		MediaLibraryStorageBytes  respjson.Field
		OriginalCacheStorageBytes respjson.Field
		VideoProcessingUnitsCount respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountUsageGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountUsageGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUsageGetParams struct {
	// Specify a `endDate` in `YYYY-MM-DD` format. It should be after the `startDate`.
	// The difference between `startDate` and `endDate` should be less than 90 days.
	EndDate time.Time `query:"endDate,required" format:"date" json:"-"`
	// Specify a `startDate` in `YYYY-MM-DD` format. It should be before the `endDate`.
	// The difference between `startDate` and `endDate` should be less than 90 days.
	StartDate time.Time `query:"startDate,required" format:"date" json:"-"`
	paramObj
}

// URLQuery serializes [AccountUsageGetParams]'s query parameters as `url.Values`.
func (r AccountUsageGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
