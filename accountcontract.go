// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/imagekit-developer/imagekit-go/v2/internal/apijson"
	"github.com/imagekit-developer/imagekit-go/v2/internal/requestconfig"
	"github.com/imagekit-developer/imagekit-go/v2/option"
	"github.com/imagekit-developer/imagekit-go/v2/packages/respjson"
)

// AccountContractService contains methods and other services that help with
// interacting with the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountContractService] method instead.
type AccountContractService struct {
	Options []option.RequestOption
}

// NewAccountContractService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountContractService(opts ...option.RequestOption) (r AccountContractService) {
	r = AccountContractService{}
	r.Options = opts
	return
}

// Get the account's contract terms and the usage consumed against them in the
// current contract term. Bandwidth is in bytes; video processing units, extension
// units and purge cache requests are counts.
//
// For an agency account, consumption includes the child accounts billed to it. A
// child account billed to its parent cannot call this endpoint; call it with the
// parent account's credentials instead.
//
// Contract terms are read live; consumption is refreshed periodically, not on
// every request. Use `asOf` to check how fresh it is. Contracts of one month or
// less are reported as `hasContract: false`. The response is cached for 5 minutes
// per account.
//
// Requires an API key with account management access.
func (r *AccountContractService) Get(ctx context.Context, opts ...option.RequestOption) (res *AccountContractGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/accounts/contract"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ContractUnits struct {
	// Bandwidth in bytes.
	BandwidthBytes int64 `json:"bandwidthBytes" api:"required"`
	// Number of extension units.
	ExtensionUnitsCount int64 `json:"extensionUnitsCount" api:"required"`
	// Number of purge cache requests.
	PurgeCacheRequests int64 `json:"purgeCacheRequests" api:"required"`
	// Number of video processing units.
	VideoProcessingUnitsCount int64 `json:"videoProcessingUnitsCount" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BandwidthBytes            respjson.Field
		ExtensionUnitsCount       respjson.Field
		PurgeCacheRequests        respjson.Field
		VideoProcessingUnitsCount respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContractUnits) RawJSON() string { return r.JSON.raw }
func (r *ContractUnits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountContractGetResponse struct {
	// Whether the account has an active contract. The other fields are present only
	// when it does.
	HasContract bool `json:"hasContract" api:"required"`
	// When `consumed` was last computed. `null` if it has not been computed for the
	// contract yet; `consumed` is then all zeros.
	AsOf time.Time `json:"asOf" api:"nullable" format:"date-time"`
	// Usage consumed in the current contract term. For an agency account, this
	// includes the child accounts billed to it.
	Consumed ContractUnits `json:"consumed"`
	// Length of each contract term in months.
	ContractMonths int64 `json:"contractMonths"`
	// End date of the current contract term.
	EffectiveContractEndDate time.Time `json:"effectiveContractEndDate" format:"date"`
	// Start date of the current contract term. It moves forward when the contract
	// renews.
	EffectiveContractStartDate time.Time `json:"effectiveContractStartDate" format:"date"`
	// Start date of the first contract term.
	OriginalContractStartDate time.Time `json:"originalContractStartDate" format:"date"`
	// Usage included in the current contract term.
	Quota ContractUnits `json:"quota"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasContract                respjson.Field
		AsOf                       respjson.Field
		Consumed                   respjson.Field
		ContractMonths             respjson.Field
		EffectiveContractEndDate   respjson.Field
		EffectiveContractStartDate respjson.Field
		OriginalContractStartDate  respjson.Field
		Quota                      respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccountContractGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AccountContractGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
