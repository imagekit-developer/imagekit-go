// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"context"
	"net/http"
	"slices"

	"github.com/imagekit-developer/imagekit-go/v2/internal/apijson"
	"github.com/imagekit-developer/imagekit-go/v2/internal/requestconfig"
	"github.com/imagekit-developer/imagekit-go/v2/option"
	"github.com/imagekit-developer/imagekit-go/v2/packages/param"
	"github.com/imagekit-developer/imagekit-go/v2/packages/respjson"
)

// AIFilterSearchService contains methods and other services that help with
// interacting with the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIFilterSearchService] method instead.
type AIFilterSearchService struct {
	Options []option.RequestOption
}

// NewAIFilterSearchService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIFilterSearchService(opts ...option.RequestOption) (r AIFilterSearchService) {
	r = AIFilterSearchService{}
	r.Options = opts
	return
}

// Convert a natural-language prompt into a structured ImageKit media-library
// search query. The response returns a `searchQuery` string (the same Lucene-like
// syntax accepted by the list and search assets API) plus suggested filter
// parameters. This endpoint only generates the query; it does not execute the
// search.
func (r *AIFilterSearchService) New(ctx context.Context, body AIFilterSearchNewParams, opts ...option.RequestOption) (res *AIFilterSearchNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/ai-filter-search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type AIFilterSearchNewResponse struct {
	// Suggested asset-type filter derived from the prompt. Empty string means no
	// file-type restriction.
	//
	// Any of "", "images", "videos", "cssJs", "others".
	FileType AIFilterSearchNewResponseFileType `json:"fileType"`
	// Whether previous file versions should be included in the search results.
	IsVersionIncludedInSearch bool `json:"isVersionIncludedInSearch"`
	// Generated query in ImageKit's Lucene-like syntax. Pass this as the `searchQuery`
	// parameter to the list and search assets API. Empty string when no filters could
	// be derived from the prompt.
	SearchQuery string `json:"searchQuery"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileType                  respjson.Field
		IsVersionIncludedInSearch respjson.Field
		SearchQuery               respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIFilterSearchNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AIFilterSearchNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Suggested asset-type filter derived from the prompt. Empty string means no
// file-type restriction.
type AIFilterSearchNewResponseFileType string

const (
	AIFilterSearchNewResponseFileTypeEmpty  AIFilterSearchNewResponseFileType = ""
	AIFilterSearchNewResponseFileTypeImages AIFilterSearchNewResponseFileType = "images"
	AIFilterSearchNewResponseFileTypeVideos AIFilterSearchNewResponseFileType = "videos"
	AIFilterSearchNewResponseFileTypeCssJs  AIFilterSearchNewResponseFileType = "cssJs"
	AIFilterSearchNewResponseFileTypeOthers AIFilterSearchNewResponseFileType = "others"
)

type AIFilterSearchNewParams struct {
	// Natural-language description of what to search for, e.g. "red dresses tagged
	// summer uploaded last month".
	Prompt string `json:"prompt" api:"required"`
	// Absolute path of the folder the user is currently in. Used to resolve relative
	// references like "this folder" in the prompt.
	CurrentFolder param.Opt[string] `json:"currentFolder,omitzero"`
	// IANA timezone (e.g. `Asia/Kolkata`) used to resolve relative date references in
	// the prompt. Defaults to UTC when omitted.
	Timezone param.Opt[string] `json:"timezone,omitzero"`
	paramObj
}

func (r AIFilterSearchNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AIFilterSearchNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIFilterSearchNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
