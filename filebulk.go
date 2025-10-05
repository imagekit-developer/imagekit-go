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
	"github.com/imagekit-developer/imagekit-go/packages/respjson"
)

// FileBulkService contains methods and other services that help with interacting
// with the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileBulkService] method instead.
type FileBulkService struct {
	Options []option.RequestOption
}

// NewFileBulkService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFileBulkService(opts ...option.RequestOption) (r FileBulkService) {
	r = FileBulkService{}
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
func (r *FileBulkService) Delete(ctx context.Context, body FileBulkDeleteParams, opts ...option.RequestOption) (res *FileBulkDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/files/batch/deleteByFileIds"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This API adds tags to multiple files in bulk. A maximum of 50 files can be
// specified at a time.
func (r *FileBulkService) AddTags(ctx context.Context, body FileBulkAddTagsParams, opts ...option.RequestOption) (res *FileBulkAddTagsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/files/addTags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This API removes AITags from multiple files in bulk. A maximum of 50 files can
// be specified at a time.
func (r *FileBulkService) RemoveAITags(ctx context.Context, body FileBulkRemoveAITagsParams, opts ...option.RequestOption) (res *FileBulkRemoveAITagsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/files/removeAITags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This API removes tags from multiple files in bulk. A maximum of 50 files can be
// specified at a time.
func (r *FileBulkService) RemoveTags(ctx context.Context, body FileBulkRemoveTagsParams, opts ...option.RequestOption) (res *FileBulkRemoveTagsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/files/removeTags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type FileBulkDeleteResponse struct {
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
func (r FileBulkDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *FileBulkDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileBulkAddTagsResponse struct {
	// An array of fileIds that in which tags were successfully added.
	SuccessfullyUpdatedFileIDs []string `json:"successfullyUpdatedFileIds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SuccessfullyUpdatedFileIDs respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileBulkAddTagsResponse) RawJSON() string { return r.JSON.raw }
func (r *FileBulkAddTagsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileBulkRemoveAITagsResponse struct {
	// An array of fileIds that in which AITags were successfully removed.
	SuccessfullyUpdatedFileIDs []string `json:"successfullyUpdatedFileIds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SuccessfullyUpdatedFileIDs respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileBulkRemoveAITagsResponse) RawJSON() string { return r.JSON.raw }
func (r *FileBulkRemoveAITagsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileBulkRemoveTagsResponse struct {
	// An array of fileIds that in which tags were successfully removed.
	SuccessfullyUpdatedFileIDs []string `json:"successfullyUpdatedFileIds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SuccessfullyUpdatedFileIDs respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileBulkRemoveTagsResponse) RawJSON() string { return r.JSON.raw }
func (r *FileBulkRemoveTagsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileBulkDeleteParams struct {
	// An array of fileIds which you want to delete.
	FileIDs []string `json:"fileIds,omitzero,required"`
	paramObj
}

func (r FileBulkDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow FileBulkDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileBulkDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileBulkAddTagsParams struct {
	// An array of fileIds to which you want to add tags.
	FileIDs []string `json:"fileIds,omitzero,required"`
	// An array of tags that you want to add to the files.
	Tags []string `json:"tags,omitzero,required"`
	paramObj
}

func (r FileBulkAddTagsParams) MarshalJSON() (data []byte, err error) {
	type shadow FileBulkAddTagsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileBulkAddTagsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileBulkRemoveAITagsParams struct {
	// An array of AITags that you want to remove from the files.
	AITags []string `json:"AITags,omitzero,required"`
	// An array of fileIds from which you want to remove AITags.
	FileIDs []string `json:"fileIds,omitzero,required"`
	paramObj
}

func (r FileBulkRemoveAITagsParams) MarshalJSON() (data []byte, err error) {
	type shadow FileBulkRemoveAITagsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileBulkRemoveAITagsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileBulkRemoveTagsParams struct {
	// An array of fileIds from which you want to remove tags.
	FileIDs []string `json:"fileIds,omitzero,required"`
	// An array of tags that you want to remove from the files.
	Tags []string `json:"tags,omitzero,required"`
	paramObj
}

func (r FileBulkRemoveTagsParams) MarshalJSON() (data []byte, err error) {
	type shadow FileBulkRemoveTagsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileBulkRemoveTagsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
