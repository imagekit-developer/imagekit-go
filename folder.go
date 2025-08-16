// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/imagekit-go/internal/apijson"
	"github.com/stainless-sdks/imagekit-go/internal/requestconfig"
	"github.com/stainless-sdks/imagekit-go/option"
	"github.com/stainless-sdks/imagekit-go/packages/param"
)

// FolderService contains methods and other services that help with interacting
// with the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFolderService] method instead.
type FolderService struct {
	Options []option.RequestOption
}

// NewFolderService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFolderService(opts ...option.RequestOption) (r FolderService) {
	r = FolderService{}
	r.Options = opts
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

type FolderNewResponse = any

type FolderDeleteResponse = any

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
