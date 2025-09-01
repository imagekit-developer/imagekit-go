// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"context"
	"net/http"

	"github.com/imagekit-developer/imagekit-go/internal/apijson"
	"github.com/imagekit-developer/imagekit-go/internal/requestconfig"
	"github.com/imagekit-developer/imagekit-go/option"
	"github.com/imagekit-developer/imagekit-go/packages/param"
	"github.com/imagekit-developer/imagekit-go/packages/respjson"
)

// FolderService contains methods and other services that help with interacting
// with the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFolderService] method instead.
type FolderService struct {
	Options []option.RequestOption
	Job     FolderJobService
}

// NewFolderService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFolderService(opts ...option.RequestOption) (r FolderService) {
	r = FolderService{}
	r.Options = opts
	r.Job = NewFolderJobService(opts...)
	return
}

// This will create a new folder. You can specify the folder name and location of
// the parent folder where this new folder should be created.
func (r *FolderService) New(ctx context.Context, body FolderNewParams, opts ...option.RequestOption) (res *FolderNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/folder"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This will delete a folder and all its contents permanently. The API returns an
// empty response.
func (r *FolderService) Delete(ctx context.Context, body FolderDeleteParams, opts ...option.RequestOption) (res *FolderDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/folder"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// This will copy one folder into another. The selected folder, its nested folders,
// files, and their versions (in `includeVersions` is set to true) are copied in
// this operation. Note: If any file at the destination has the same name as the
// source file, then the source file and its versions will be appended to the
// destination file version history.
func (r *FolderService) Copy(ctx context.Context, body FolderCopyParams, opts ...option.RequestOption) (res *FolderCopyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/bulkJobs/copyFolder"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This will move one folder into another. The selected folder, its nested folders,
// files, and their versions are moved in this operation. Note: If any file at the
// destination has the same name as the source file, then the source file and its
// versions will be appended to the destination file version history.
func (r *FolderService) Move(ctx context.Context, body FolderMoveParams, opts ...option.RequestOption) (res *FolderMoveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/bulkJobs/moveFolder"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This API allows you to rename an existing folder. The folder and all its nested
// assets and sub-folders will remain unchanged, but their paths will be updated to
// reflect the new folder name.
func (r *FolderService) Rename(ctx context.Context, body FolderRenameParams, opts ...option.RequestOption) (res *FolderRenameResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/bulkJobs/renameFolder"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type FolderNewResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FolderNewResponse) RawJSON() string { return r.JSON.raw }
func (r *FolderNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FolderDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *FolderDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Job submitted successfully. A `jobId` will be returned.
type FolderCopyResponse struct {
	// Unique identifier of the bulk job. This can be used to check the status of the
	// bulk job.
	JobID string `json:"jobId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JobID       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FolderCopyResponse) RawJSON() string { return r.JSON.raw }
func (r *FolderCopyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Job submitted successfully. A `jobId` will be returned.
type FolderMoveResponse struct {
	// Unique identifier of the bulk job. This can be used to check the status of the
	// bulk job.
	JobID string `json:"jobId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JobID       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FolderMoveResponse) RawJSON() string { return r.JSON.raw }
func (r *FolderMoveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Job submitted successfully. A `jobId` will be returned.
type FolderRenameResponse struct {
	// Unique identifier of the bulk job. This can be used to check the status of the
	// bulk job.
	JobID string `json:"jobId,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JobID       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FolderRenameResponse) RawJSON() string { return r.JSON.raw }
func (r *FolderRenameResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderNewParams struct {
	// The folder will be created with this name.
	//
	// All characters except alphabets and numbers (inclusive of unicode letters,
	// marks, and numerals in other languages) will be replaced by an underscore i.e.
	// `_`.
	FolderName string `json:"folderName,required"`
	// The folder where the new folder should be created, for root use `/` else the
	// path e.g. `containing/folder/`.
	//
	// Note: If any folder(s) is not present in the parentFolderPath parameter, it will
	// be automatically created. For example, if you pass `/product/images/summer`,
	// then `product`, `images`, and `summer` folders will be created if they don't
	// already exist.
	ParentFolderPath string `json:"parentFolderPath,required"`
	paramObj
}

func (r FolderNewParams) MarshalJSON() (data []byte, err error) {
	type shadow FolderNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FolderNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderDeleteParams struct {
	// Full path to the folder you want to delete. For example `/folder/to/delete/`.
	FolderPath string `json:"folderPath,required"`
	paramObj
}

func (r FolderDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow FolderDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FolderDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderCopyParams struct {
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

func (r FolderCopyParams) MarshalJSON() (data []byte, err error) {
	type shadow FolderCopyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FolderCopyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderMoveParams struct {
	// Full path to the destination folder where you want to move the source folder
	// into.
	DestinationPath string `json:"destinationPath,required"`
	// The full path to the source folder you want to move.
	SourceFolderPath string `json:"sourceFolderPath,required"`
	paramObj
}

func (r FolderMoveParams) MarshalJSON() (data []byte, err error) {
	type shadow FolderMoveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FolderMoveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderRenameParams struct {
	// The full path to the folder you want to rename.
	FolderPath string `json:"folderPath,required"`
	// The new name for the folder.
	//
	// All characters except alphabets and numbers (inclusive of unicode letters,
	// marks, and numerals in other languages) and `-` will be replaced by an
	// underscore i.e. `_`.
	NewFolderName string `json:"newFolderName,required"`
	// Option to purge cache for the old nested files and their versions' URLs.
	//
	// When set to true, it will internally issue a purge cache request on CDN to
	// remove the cached content of the old nested files and their versions. There will
	// only be one purge request for all the nested files, which will be counted
	// against your monthly purge quota.
	//
	// Note: A purge cache request will be issued against
	// `https://ik.imagekit.io/old/folder/path*` (with a wildcard at the end). This
	// will remove all nested files, their versions' URLs, and any transformations made
	// using query parameters on these files or their versions. However, the cache for
	// file transformations made using path parameters will persist. You can purge them
	// using the purge API. For more details, refer to the purge API documentation.
	//
	// Default value - `false`
	PurgeCache param.Opt[bool] `json:"purgeCache,omitzero"`
	paramObj
}

func (r FolderRenameParams) MarshalJSON() (data []byte, err error) {
	type shadow FolderRenameParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FolderRenameParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
