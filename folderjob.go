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
	"github.com/imagekit-developer/imagekit-go/packages/respjson"
)

// FolderJobService contains methods and other services that help with interacting
// with the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFolderJobService] method instead.
type FolderJobService struct {
	Options []option.RequestOption
}

// NewFolderJobService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFolderJobService(opts ...option.RequestOption) (r FolderJobService) {
	r = FolderJobService{}
	r.Options = opts
	return
}

// This API returns the status of a bulk job like copy and move folder operations.
func (r *FolderJobService) Get(ctx context.Context, jobID string, opts ...option.RequestOption) (res *FolderJobGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return
	}
	path := fmt.Sprintf("v1/bulkJobs/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type FolderJobGetResponse struct {
	// Unique identifier of the bulk job.
	JobID string `json:"jobId"`
	// Unique identifier of the purge request. This will be present only if
	// `purgeCache` is set to `true` in the rename folder API request.
	PurgeRequestID string `json:"purgeRequestId"`
	// Status of the bulk job.
	//
	// Any of "Pending", "Completed".
	Status FolderJobGetResponseStatus `json:"status"`
	// Type of the bulk job.
	//
	// Any of "COPY_FOLDER", "MOVE_FOLDER", "RENAME_FOLDER".
	Type FolderJobGetResponseType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JobID          respjson.Field
		PurgeRequestID respjson.Field
		Status         respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FolderJobGetResponse) RawJSON() string { return r.JSON.raw }
func (r *FolderJobGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the bulk job.
type FolderJobGetResponseStatus string

const (
	FolderJobGetResponseStatusPending   FolderJobGetResponseStatus = "Pending"
	FolderJobGetResponseStatusCompleted FolderJobGetResponseStatus = "Completed"
)

// Type of the bulk job.
type FolderJobGetResponseType string

const (
	FolderJobGetResponseTypeCopyFolder   FolderJobGetResponseType = "COPY_FOLDER"
	FolderJobGetResponseTypeMoveFolder   FolderJobGetResponseType = "MOVE_FOLDER"
	FolderJobGetResponseTypeRenameFolder FolderJobGetResponseType = "RENAME_FOLDER"
)
