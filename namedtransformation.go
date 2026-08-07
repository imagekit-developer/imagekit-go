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
//
// **Note:** You can create up to 250 named transformations per account. Once this
// limit is reached, the request fails with a `400` error.
func (r *NamedTransformationService) New(ctx context.Context, body NamedTransformationNewParams, opts ...option.RequestOption) (res *shared.NamedTransformation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/named-transformations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Updates the named transformation identified by `id` and returns the updated
// object. Only the fields present in the request body are updated; omitted fields
// are left unchanged.
//
// **Note:**
//
//   - If you rename this named transformation, or set `enabled` to `false`, and
//     another _enabled_ named transformation, or your account's upload
//     pre-transformation/post-transformation settings, reference it (via the
//     `n-<name>` token), the request fails with a `409` error whose `message`
//     describes what it is referenced by. A reference from a named transformation
//     that is itself disabled does not block this request. Remove or disable those
//     references first, then retry. This is a best-effort check and cannot detect
//     references baked into your own application code or previously generated URLs.
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
//   - If another _enabled_ named transformation, or your account's upload
//     pre-transformation/post-transformation settings, reference this named
//     transformation (via the `n-<name>` token), the request fails with a `409`
//     error whose `message` describes what it is referenced by. A reference from a
//     named transformation that is itself disabled does not block this request.
//     Remove or disable those references first, then retry the deletion. This is a
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
	// only contain alphanumeric characters or `_` (hyphens are not allowed), and must
	// be unique for your account. Name matching is case-sensitive, so
	// `Small_Thumbnail` and `small_thumbnail` are treated as different names.
	Name string `json:"name" api:"required"`
	// The transformation this name refers to, expressed as one or more comma-separated
	// transformation parameters, for example `w-150,h-150,fo-center,cm-resize`. You do
	// not need to prefix this with `tr:` — it is added automatically. If you do
	// include it, it must appear in lowercase at the start of the string, or the
	// request is rejected. Learn more about the
	// [transformation syntax](https://imagekit.io/docs/transformations).
	Transformation string `json:"transformation" api:"required"`
	// Whether this named transformation is enabled. Set to `false` to temporarily
	// disable it without deleting it — requests using a disabled named transformation
	// fail at delivery time.
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
	// Whether this named transformation is enabled. If omitted, the existing value is
	// left unchanged.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Updated name of the named transformation. Can only contain alphanumeric
	// characters and `_`, and must be unique for your account. Name matching is
	// case-sensitive, so `Small_Thumbnail` and `small_thumbnail` are treated as
	// different names.
	Name param.Opt[string] `json:"name,omitzero"`
	// Updated transformation, expressed as one or more comma-separated transformation
	// parameters. You do not need to prefix this with `tr:` — it is added
	// automatically. If you do include it, it must appear in lowercase at the start of
	// the string, or the request is rejected.
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
