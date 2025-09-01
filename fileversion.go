// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/imagekit-developer/imagekit-go/internal/apijson"
	"github.com/imagekit-developer/imagekit-go/internal/requestconfig"
	"github.com/imagekit-developer/imagekit-go/option"
	"github.com/imagekit-developer/imagekit-go/packages/respjson"
)

// FileVersionService contains methods and other services that help with
// interacting with the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileVersionService] method instead.
type FileVersionService struct {
	Options []option.RequestOption
}

// NewFileVersionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFileVersionService(opts ...option.RequestOption) (r FileVersionService) {
	r = FileVersionService{}
	r.Options = opts
	return
}

// This API returns details of all versions of a file.
func (r *FileVersionService) List(ctx context.Context, fileID string, opts ...option.RequestOption) (res *[]File, err error) {
	opts = append(r.Options[:], opts...)
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return
	}
	path := fmt.Sprintf("v1/files/%s/versions", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This API deletes a non-current file version permanently. The API returns an
// empty response.
//
// Note: If you want to delete all versions of a file, use the delete file API.
func (r *FileVersionService) Delete(ctx context.Context, versionID string, body FileVersionDeleteParams, opts ...option.RequestOption) (res *FileVersionDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.FileID == "" {
		err = errors.New("missing required fileId parameter")
		return
	}
	if versionID == "" {
		err = errors.New("missing required versionId parameter")
		return
	}
	path := fmt.Sprintf("v1/files/%s/versions/%s", body.FileID, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// This API returns an object with details or attributes of a file version.
func (r *FileVersionService) Get(ctx context.Context, versionID string, query FileVersionGetParams, opts ...option.RequestOption) (res *File, err error) {
	opts = append(r.Options[:], opts...)
	if query.FileID == "" {
		err = errors.New("missing required fileId parameter")
		return
	}
	if versionID == "" {
		err = errors.New("missing required versionId parameter")
		return
	}
	path := fmt.Sprintf("v1/files/%s/versions/%s", query.FileID, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This API restores a file version as the current file version.
func (r *FileVersionService) Restore(ctx context.Context, versionID string, body FileVersionRestoreParams, opts ...option.RequestOption) (res *File, err error) {
	opts = append(r.Options[:], opts...)
	if body.FileID == "" {
		err = errors.New("missing required fileId parameter")
		return
	}
	if versionID == "" {
		err = errors.New("missing required versionId parameter")
		return
	}
	path := fmt.Sprintf("v1/files/%s/versions/%s/restore", body.FileID, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

type FileVersionDeleteResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileVersionDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *FileVersionDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileVersionDeleteParams struct {
	FileID string `path:"fileId,required" json:"-"`
	paramObj
}

type FileVersionGetParams struct {
	FileID string `path:"fileId,required" json:"-"`
	paramObj
}

type FileVersionRestoreParams struct {
	FileID string `path:"fileId,required" json:"-"`
	paramObj
}
