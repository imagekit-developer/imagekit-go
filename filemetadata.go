// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/imagekit-developer/imagekit-go/internal/apiquery"
	"github.com/imagekit-developer/imagekit-go/internal/requestconfig"
	"github.com/imagekit-developer/imagekit-go/option"
)

// FileMetadataService contains methods and other services that help with
// interacting with the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileMetadataService] method instead.
type FileMetadataService struct {
	Options []option.RequestOption
}

// NewFileMetadataService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFileMetadataService(opts ...option.RequestOption) (r FileMetadataService) {
	r = FileMetadataService{}
	r.Options = opts
	return
}

// You can programmatically get image EXIF, pHash, and other metadata for uploaded
// files in the ImageKit.io media library using this API.
//
// You can also get the metadata in upload API response by passing `metadata` in
// `responseFields` parameter.
func (r *FileMetadataService) Get(ctx context.Context, fileID string, opts ...option.RequestOption) (res *Metadata, err error) {
	opts = append(r.Options[:], opts...)
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return
	}
	path := fmt.Sprintf("v1/files/%s/metadata", fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get image EXIF, pHash, and other metadata from ImageKit.io powered remote URL
// using this API.
func (r *FileMetadataService) GetFromURL(ctx context.Context, query FileMetadataGetFromURLParams, opts ...option.RequestOption) (res *Metadata, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/files/metadata"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type FileMetadataGetFromURLParams struct {
	// Should be a valid file URL. It should be accessible using your ImageKit.io
	// account.
	URL string `query:"url,required" format:"uri" json:"-"`
	paramObj
}

// URLQuery serializes [FileMetadataGetFromURLParams]'s query parameters as
// `url.Values`.
func (r FileMetadataGetFromURLParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
