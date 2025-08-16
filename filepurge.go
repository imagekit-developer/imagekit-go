// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/imagekit-go/internal/apijson"
	"github.com/stainless-sdks/imagekit-go/internal/requestconfig"
	"github.com/stainless-sdks/imagekit-go/option"
	"github.com/stainless-sdks/imagekit-go/packages/param"
	"github.com/stainless-sdks/imagekit-go/packages/respjson"
)

// FilePurgeService contains methods and other services that help with interacting
// with the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFilePurgeService] method instead.
type FilePurgeService struct {
	Options []option.RequestOption
}

// NewFilePurgeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFilePurgeService(opts ...option.RequestOption) (r FilePurgeService) {
	r = FilePurgeService{}
	r.Options = opts
	return
}

// This API will purge CDN cache and ImageKit.io's internal cache for a file. Note:
// Purge cache is an asynchronous process and it may take some time to reflect the
// changes.
func (r *FilePurgeService) Execute(ctx context.Context, body FilePurgeExecuteParams, opts ...option.RequestOption) (res *FilePurgeExecuteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/files/purge"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This API returns the status of a purge cache request.
func (r *FilePurgeService) Status(ctx context.Context, requestID string, opts ...option.RequestOption) (res *FilePurgeStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	if requestID == "" {
		err = errors.New("missing required requestId parameter")
		return
	}
	path := fmt.Sprintf("v1/files/purge/%s", requestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type FilePurgeExecuteResponse struct {
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
func (r FilePurgeExecuteResponse) RawJSON() string { return r.JSON.raw }
func (r *FilePurgeExecuteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FilePurgeStatusResponse struct {
	// Status of the purge request.
	//
	// Any of "Pending", "Completed".
	Status FilePurgeStatusResponseStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FilePurgeStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *FilePurgeStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the purge request.
type FilePurgeStatusResponseStatus string

const (
	FilePurgeStatusResponseStatusPending   FilePurgeStatusResponseStatus = "Pending"
	FilePurgeStatusResponseStatusCompleted FilePurgeStatusResponseStatus = "Completed"
)

type FilePurgeExecuteParams struct {
	// The full URL of the file to be purged.
	URL string `json:"url,required"`
	paramObj
}

func (r FilePurgeExecuteParams) MarshalJSON() (data []byte, err error) {
	type shadow FilePurgeExecuteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FilePurgeExecuteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
