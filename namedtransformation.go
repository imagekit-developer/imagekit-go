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
// A named transformation is a short, reusable name for a transformation string.
// Use it in image and video URLs as `tr:n-<name>`, and update the underlying
// transformation later without changing existing URLs. Learn more about
// [named transformations](https://imagekit.io/docs/transformations#named-transformations).
//
// You can create up to 250 named transformations per account.
func (r *NamedTransformationService) New(ctx context.Context, body NamedTransformationNewParams, opts ...option.RequestOption) (res *shared.NamedTransformation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/named-transformations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Updates the named transformation identified by `id` and returns the updated
// object. Only the fields present in the request body are updated; other fields
// stay unchanged.
//
// Renaming or disabling a named transformation fails with a `409` error if it is
// still referenced (via the `n-<name>` token) by an upload pre-transformation or
// post-transformation setting. This check is best-effort and can't detect
// references in your own application code or in previously generated URLs.
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

// Permanently deletes the named transformation identified by `id`.
//
// Deletion fails with a `409` error if the named transformation is still
// referenced (via the `n-<name>` token) by an upload pre-transformation or
// post-transformation setting. This check is best-effort and can't detect
// references in your own application code or in previously generated URLs.
func (r *NamedTransformationService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/named-transformations/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
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
	// Alias for the transformation string, used in URLs as `tr:n-<name>`. This is
	// case-sensitive, contains only alphanumeric characters or `_` (underscore), and
	// is unique across all named transformations for your account.
	Name string `json:"name" api:"required"`
	// The transformation string this named transformation refers to. Learn more about
	// the [transformation string syntax](https://imagekit.io/docs/transformations).
	Transformation string `json:"transformation" api:"required"`
	// Whether the named transformation is currently enabled. When set to `false`,
	// requests using this named transformation fail at delivery time.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
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
	// Whether the named transformation is enabled. Omit to leave the current value
	// unchanged.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Alias for the transformation string, used in URLs as `tr:n-<name>`. This is
	// case-sensitive, contains only alphanumeric characters or `_` (underscore), and
	// is unique across all named transformations for your account.
	Name param.Opt[string] `json:"name,omitzero"`
	// The transformation string this named transformation refers to. Learn more about
	// the [transformation string syntax](https://imagekit.io/docs/transformations).
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
