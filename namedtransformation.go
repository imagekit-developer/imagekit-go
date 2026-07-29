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

// NamedTransformationService contains methods and other services that help with
// interacting with the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNamedTransformationService] method instead.
type NamedTransformationService struct {
	Options []option.RequestOption
}

// NewNamedTransformationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNamedTransformationService(opts ...option.RequestOption) (r NamedTransformationService) {
	r = NamedTransformationService{}
	r.Options = opts
	return
}

// Creates a new named transformation and returns the created object.
//
// Named transformations let you assign a short, reusable name to a complex
// transformation string, so it can be applied in image and video URLs as
// `tr:n-<name>` and later updated without changing any existing URLs.
//
// Learn more about
// [named transformations](https://imagekit.io/docs/transformations#named-transformations).
func (r *NamedTransformationService) New(ctx context.Context, body NamedTransformationNewParams, opts ...option.RequestOption) (res *shared.NamedTransformation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/named-transformations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Updates the named transformation identified by `id` and returns the updated
// object. Only the fields present in the request body are updated; omitted fields
// are left unchanged.
func (r *NamedTransformationService) Update(ctx context.Context, id string, body NamedTransformationUpdateParams, opts ...option.RequestOption) (res *shared.NamedTransformation, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/named-transformations/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns an array of all named transformations configured for your account.
func (r *NamedTransformationService) List(ctx context.Context, opts ...option.RequestOption) (res *[]shared.NamedTransformation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/named-transformations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Permanently deletes the named transformation identified by `id` and returns the
// deleted object.
//
// **Note:**
//
//   - If another named transformation, or your account's upload
//     pre-transformation/post-transformation settings, reference this named
//     transformation (via the `n-<name>` token), the request fails with a `409`
//     error and the response body includes a `references` array describing where it
//     is used. Remove those references first, then retry the deletion. This is a
//     best-effort check and cannot detect references baked into your own application
//     code or previously generated URLs.
func (r *NamedTransformationService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *shared.NamedTransformation, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/named-transformations/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Retrieves the named transformation identified by `id`.
func (r *NamedTransformationService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *shared.NamedTransformation, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/named-transformations/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type NamedTransformationNewParams struct {
	// Name of the named transformation. This is the alias used to refer to the
	// transformation string in image and video URLs, for example `tr:n-<name>`. Can
	// only contain alphanumeric characters, `_` and `-`, and must be unique for your
	// account (case-insensitive).
	Name string `json:"name" api:"required"`
	// The transformation string this name refers to. It must start with `tr:` followed
	// by one or more transformation parameters, for example
	// `tr:w-150,h-150,fo-center,cm-resize`. Learn more about the
	// [transformation syntax](https://imagekit.io/docs/transformations).
	Transformation string `json:"transformation" api:"required"`
	// Whether this named transformation is disabled. Set to `true` to temporarily
	// disable it without deleting it — requests using a disabled named transformation
	// fail at delivery time.
	Disabled param.Opt[bool] `json:"disabled,omitzero"`
	paramObj
}

func (r NamedTransformationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow NamedTransformationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NamedTransformationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NamedTransformationUpdateParams struct {
	// Whether this named transformation is disabled.
	Disabled param.Opt[bool] `json:"disabled,omitzero"`
	// Updated name of the named transformation. Can only contain alphanumeric
	// characters, `_` and `-`, and must be unique for your account (case-insensitive).
	Name param.Opt[string] `json:"name,omitzero"`
	// Updated transformation string. It must start with `tr:` followed by one or more
	// transformation parameters.
	Transformation param.Opt[string] `json:"transformation,omitzero"`
	paramObj
}

func (r NamedTransformationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow NamedTransformationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NamedTransformationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
