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

// BulkJobService contains methods and other services that help with interacting
// with the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBulkJobService] method instead.
type BulkJobService struct {
	Options []option.RequestOption
}

// NewBulkJobService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBulkJobService(opts ...option.RequestOption) (r BulkJobService) {
	r = BulkJobService{}
	r.Options = opts
	return
}

// This will copy one folder into another. The selected folder, its nested folders,
// files, and their versions (in `includeVersions` is set to true) are copied in
// this operation. Note: If any file at the destination has the same name as the
// source file, then the source file and its versions will be appended to the
// destination file version history.
func (r *BulkJobService) CopyFolder(ctx context.Context, body BulkJobCopyFolderParams, opts ...option.RequestOption) (res *BulkJobCopyFolderResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/bulkJobs/copyFolder"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This will move one folder into another. The selected folder, its nested folders,
// files, and their versions are moved in this operation. Note: If any file at the
// destination has the same name as the source file, then the source file and its
// versions will be appended to the destination file version history.
func (r *BulkJobService) MoveFolder(ctx context.Context, body BulkJobMoveFolderParams, opts ...option.RequestOption) (res *BulkJobMoveFolderResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/bulkJobs/moveFolder"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This API returns the status of a bulk job like copy and move folder operations.
func (r *BulkJobService) GetStatus(ctx context.Context, jobID string, opts ...option.RequestOption) (res *BulkJobGetStatusResponse, err error) {
	opts = append(r.Options[:], opts...)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return
	}
	path := fmt.Sprintf("v1/bulkJobs/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type BulkJobCopyFolderResponse struct {
	// Unique identifier of the bulk job. This can be used to check the status of the
	// bulk job.
	JobID string `json:"jobId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JobID       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BulkJobCopyFolderResponse) RawJSON() string { return r.JSON.raw }
func (r *BulkJobCopyFolderResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BulkJobMoveFolderResponse struct {
	// Unique identifier of the bulk job. This can be used to check the status of the
	// bulk job.
	JobID string `json:"jobId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JobID       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BulkJobMoveFolderResponse) RawJSON() string { return r.JSON.raw }
func (r *BulkJobMoveFolderResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BulkJobGetStatusResponse struct {
	// Unique identifier of the bulk job.
	JobID string `json:"jobId"`
	// Status of the bulk job. Possible values - `Pending`, `Completed`.
	Status string `json:"status"`
	// Type of the bulk job. Possible values - `COPY_FOLDER`, `MOVE_FOLDER`.
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JobID       respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BulkJobGetStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *BulkJobGetStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BulkJobCopyFolderParams struct {
	// Full path to the destination folder where you want to copy the source folder
	// into.
	DestinationPath string `json:"destinationPath,required"`
	// The full path to the source folder you want to copy.
	SourceFolderPath string `json:"sourceFolderPath,required"`
	// Option to copy all versions of files that are nested inside the selected folder.
	// By default, only the current version of each file will be copied. When set to
	// true, all versions of each file will be copied. Default value - `false`.
	IncludeVersions param.Opt[bool] `json:"includeVersions,omitzero"`
	paramObj
}

func (r BulkJobCopyFolderParams) MarshalJSON() (data []byte, err error) {
	type shadow BulkJobCopyFolderParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BulkJobCopyFolderParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BulkJobMoveFolderParams struct {
	// Full path to the destination folder where you want to move the source folder
	// into.
	DestinationPath string `json:"destinationPath,required"`
	// The full path to the source folder you want to move.
	SourceFolderPath string `json:"sourceFolderPath,required"`
	paramObj
}

func (r BulkJobMoveFolderParams) MarshalJSON() (data []byte, err error) {
	type shadow BulkJobMoveFolderParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BulkJobMoveFolderParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
