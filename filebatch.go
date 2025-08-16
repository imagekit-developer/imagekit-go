// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/imagekit-go/internal/apijson"
	"github.com/stainless-sdks/imagekit-go/internal/requestconfig"
	"github.com/stainless-sdks/imagekit-go/option"
	"github.com/stainless-sdks/imagekit-go/packages/param"
	"github.com/stainless-sdks/imagekit-go/packages/respjson"
)

// FileBatchService contains methods and other services that help with interacting
// with the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileBatchService] method instead.
type FileBatchService struct {
	Options []option.RequestOption
}

// NewFileBatchService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFileBatchService(opts ...option.RequestOption) (r FileBatchService) {
	r = FileBatchService{}
	r.Options = opts
	return
}

// This API deletes multiple files and all their file versions permanently.
//
// Note: If a file or specific transformation has been requested in the past, then
// the response is cached. Deleting a file does not purge the cache. You can purge
// the cache using purge cache API.
//
// A maximum of 100 files can be deleted at a time.
func (r *FileBatchService) Delete(ctx context.Context, body FileBatchDeleteParams, opts ...option.RequestOption) (res *FileBatchDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/files/batch/deleteByFileIds"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type FileBatchDeleteResponse struct {
	// An array of fileIds that were successfully deleted.
	SuccessfullyDeletedFileIDs []string `json:"successfullyDeletedFileIds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SuccessfullyDeletedFileIDs respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileBatchDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *FileBatchDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileBatchDeleteParams struct {
	// An array of fileIds which you want to delete.
	FileIDs []string `json:"fileIds,omitzero,required"`
	paramObj
}

func (r FileBatchDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow FileBatchDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileBatchDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
