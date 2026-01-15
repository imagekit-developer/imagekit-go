// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/imagekit-developer/imagekit-go/v2/internal/apijson"
	"github.com/imagekit-developer/imagekit-go/v2/internal/requestconfig"
	"github.com/imagekit-developer/imagekit-go/v2/option"
	"github.com/imagekit-developer/imagekit-go/v2/packages/param"
	"github.com/imagekit-developer/imagekit-go/v2/shared"
)

// SavedExtensionService contains methods and other services that help with
// interacting with the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSavedExtensionService] method instead.
type SavedExtensionService struct {
	Options []option.RequestOption
}

// NewSavedExtensionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSavedExtensionService(opts ...option.RequestOption) (r SavedExtensionService) {
	r = SavedExtensionService{}
	r.Options = opts
	return
}

// This API creates a new saved extension. Saved extensions allow you to save
// complex extension configurations (like AI tasks) and reuse them by referencing
// the ID in upload or update file APIs.
//
// **Saved extension limit** \
// You can create a maximum of 100 saved extensions per account.
func (r *SavedExtensionService) New(ctx context.Context, body SavedExtensionNewParams, opts ...option.RequestOption) (res *shared.SavedExtension, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/saved-extensions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// This API updates an existing saved extension. You can update the name,
// description, or config.
func (r *SavedExtensionService) Update(ctx context.Context, id string, body SavedExtensionUpdateParams, opts ...option.RequestOption) (res *shared.SavedExtension, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/saved-extensions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// This API returns an array of all saved extensions for your account. Saved
// extensions allow you to save complex extension configurations and reuse them by
// referencing them by ID in upload or update file APIs.
func (r *SavedExtensionService) List(ctx context.Context, opts ...option.RequestOption) (res *[]shared.SavedExtension, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/saved-extensions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This API deletes a saved extension permanently.
func (r *SavedExtensionService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/saved-extensions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// This API returns details of a specific saved extension by ID.
func (r *SavedExtensionService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *shared.SavedExtension, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/saved-extensions/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SavedExtensionNewParams struct {
	// Configuration object for an extension (base extensions only, not saved extension
	// references).
	Config shared.ExtensionConfigUnionParam `json:"config,omitzero,required"`
	// Description of what the saved extension does.
	Description string `json:"description,required"`
	// Name of the saved extension.
	Name string `json:"name,required"`
	paramObj
}

func (r SavedExtensionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SavedExtensionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SavedExtensionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SavedExtensionUpdateParams struct {
	// Updated description of the saved extension.
	Description param.Opt[string] `json:"description,omitzero"`
	// Updated name of the saved extension.
	Name param.Opt[string] `json:"name,omitzero"`
	// Configuration object for an extension (base extensions only, not saved extension
	// references).
	Config shared.ExtensionConfigUnionParam `json:"config,omitzero"`
	paramObj
}

func (r SavedExtensionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow SavedExtensionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SavedExtensionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
