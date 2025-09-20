// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/imagekit-developer/imagekit-go/internal/apijson"
	"github.com/imagekit-developer/imagekit-go/internal/requestconfig"
	"github.com/imagekit-developer/imagekit-go/option"
	"github.com/imagekit-developer/imagekit-go/packages/param"
	"github.com/imagekit-developer/imagekit-go/packages/respjson"
)

// CacheInvalidationService contains methods and other services that help with
// interacting with the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCacheInvalidationService] method instead.
type CacheInvalidationService struct {
	Options []option.RequestOption
}

// NewCacheInvalidationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCacheInvalidationService(opts ...option.RequestOption) (r CacheInvalidationService) {
	r = CacheInvalidationService{}
	r.Options = opts
	return
}

// This API will purge CDN cache and ImageKit.io's internal cache for a file. Note:
// Purge cache is an asynchronous process and it may take some time to reflect the
// changes.
func (r *CacheInvalidationService) New(ctx context.Context, body CacheInvalidationNewParams, opts ...option.RequestOption) (res *CacheInvalidationNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/files/purge"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This API returns the status of a purge cache request.
func (r *CacheInvalidationService) Get(ctx context.Context, requestID string, opts ...option.RequestOption) (res *CacheInvalidationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if requestID == "" {
		err = errors.New("missing required requestId parameter")
		return
	}
	path := fmt.Sprintf("v1/files/purge/%s", requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CacheInvalidationNewResponse struct {
	// Unique identifier of the purge request. This can be used to check the status of
	// the purge request.
	RequestID string `json:"requestId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CacheInvalidationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *CacheInvalidationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CacheInvalidationGetResponse struct {
	// Status of the purge request.
	//
	// Any of "Pending", "Completed".
	Status CacheInvalidationGetResponseStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CacheInvalidationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *CacheInvalidationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the purge request.
type CacheInvalidationGetResponseStatus string

const (
	CacheInvalidationGetResponseStatusPending   CacheInvalidationGetResponseStatus = "Pending"
	CacheInvalidationGetResponseStatusCompleted CacheInvalidationGetResponseStatus = "Completed"
)

type CacheInvalidationNewParams struct {
	// The full URL of the file to be purged.
	URL string `json:"url,required" format:"uri"`
	paramObj
}

func (r CacheInvalidationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CacheInvalidationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CacheInvalidationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
