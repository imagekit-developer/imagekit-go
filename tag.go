// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/imagekit-developer/imagekit-go/v2/internal/apijson"
	"github.com/imagekit-developer/imagekit-go/v2/internal/apiquery"
	"github.com/imagekit-developer/imagekit-go/v2/internal/requestconfig"
	"github.com/imagekit-developer/imagekit-go/v2/option"
	"github.com/imagekit-developer/imagekit-go/v2/packages/param"
	"github.com/imagekit-developer/imagekit-go/v2/packages/respjson"
)

// TagService contains methods and other services that help with interacting with
// the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTagService] method instead.
type TagService struct {
	Options []option.RequestOption
}

// NewTagService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTagService(opts ...option.RequestOption) (r TagService) {
	r = TagService{}
	r.Options = opts
	return
}

// This API lists the distinct tags in use across your media library, with the
// number of assets each tag is applied to and when it was last used. Tags are
// matched case-insensitively, so `Summer` and `summer` are one entry: `name` is
// the lowercased form and `label` is the spelling that was used most recently.
//
// Use `tagSource` and `fileType` to list only the tags in use on matching assets.
// `lastUsedAt` and the `ASC_LAST_USED` / `DESC_LAST_USED` sorts follow those
// filters, while `assetCount` always counts every asset carrying the tag. Use
// `search` or `prefix` to narrow the listing and `cursor` to page through it.
//
// Tag counts are updated asynchronously. Changes to a file's tags usually appear
// in this list within about a minute, but can take longer during heavy tagging
// activity, so a tag you just added or removed may not be reflected immediately.
func (r *TagService) List(ctx context.Context, query TagListParams, opts ...option.RequestOption) (res *TagListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/tags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A tag in use in your media library, aggregated case-insensitively across the
// assets that carry it.
type Tag struct {
	// The number of assets carrying the tag, counted across both tag sources
	// (user-generated tags and AI-generated tags) and all file types regardless of the
	// request's filters.
	AssetCount int64 `json:"assetCount" api:"required"`
	// The tag as it was spelled when it was last applied to any asset. Use it for
	// display.
	Label string `json:"label" api:"required"`
	// When the tag was last applied to an asset matching the `tagSource` and
	// `fileType` filters of the request, in ISO 8601 format.
	LastUsedAt time.Time `json:"lastUsedAt" api:"required" format:"date-time"`
	// The tag in lowercase. Tags are case-insensitive, so this is the unique key of
	// the tag within your account.
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssetCount  respjson.Field
		Label       respjson.Field
		LastUsedAt  respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Tag) RawJSON() string { return r.JSON.raw }
func (r *Tag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagListResponse struct {
	// Cursor for the next page, or `null` when there are no more tags.
	NextCursor string `json:"nextCursor" api:"required"`
	// Tags matching the request, in the requested order.
	Tags []Tag `json:"tags" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextCursor  respjson.Field
		Tags        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagListResponse) RawJSON() string { return r.JSON.raw }
func (r *TagListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagListParams struct {
	// The `nextCursor` value from the previous response, to fetch the next page. A
	// cursor is bound to the `tagSource`, `fileType`, `sort`, `search` and `prefix` it
	// was issued with; sending it with different values returns a 400.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// The maximum number of tags to return in response.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Only return tags that begin with this value. Matching is case-insensitive.
	Prefix param.Opt[string] `query:"prefix,omitzero" json:"-"`
	// Only return tags that contain this text. Every word in the value must appear in
	// the tag and the last word may be incomplete, which suits search-as-you-type.
	// Matching is case-insensitive.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Filter tags by the type of asset they are applied to.
	//
	//   - `image` — tags in use on images
	//   - `video` — tags in use on videos
	//   - `static` — tags in use on any other file, such as documents, scripts and
	//     stylesheets
	//   - `all` — tags in use on any asset
	//
	// Any of "image", "video", "static", "all".
	FileType TagListParamsFileType `query:"fileType,omitzero" json:"-"`
	// Sort the results.
	//
	//   - `ASC_NAME` / `DESC_NAME` — by `name`
	//   - `ASC_LAST_USED` / `DESC_LAST_USED` — by `lastUsedAt` within the requested
	//     `tagSource` and `fileType`, ties broken by `name`
	//
	// Any of "ASC_NAME", "DESC_NAME", "ASC_LAST_USED", "DESC_LAST_USED".
	Sort TagListParamsSort `query:"sort,omitzero" json:"-"`
	// Filter tags by who added them.
	//
	// - `user` — tags added by users through the dashboard or the API
	// - `ai` — tags added by AI auto-tagging extensions
	// - `all` — both
	//
	// Any of "user", "ai", "all".
	TagSource TagListParamsTagSource `query:"tagSource,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TagListParams]'s query parameters as `url.Values`.
func (r TagListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter tags by the type of asset they are applied to.
//
//   - `image` — tags in use on images
//   - `video` — tags in use on videos
//   - `static` — tags in use on any other file, such as documents, scripts and
//     stylesheets
//   - `all` — tags in use on any asset
type TagListParamsFileType string

const (
	TagListParamsFileTypeImage  TagListParamsFileType = "image"
	TagListParamsFileTypeVideo  TagListParamsFileType = "video"
	TagListParamsFileTypeStatic TagListParamsFileType = "static"
	TagListParamsFileTypeAll    TagListParamsFileType = "all"
)

// Sort the results.
//
//   - `ASC_NAME` / `DESC_NAME` — by `name`
//   - `ASC_LAST_USED` / `DESC_LAST_USED` — by `lastUsedAt` within the requested
//     `tagSource` and `fileType`, ties broken by `name`
type TagListParamsSort string

const (
	TagListParamsSortAscName      TagListParamsSort = "ASC_NAME"
	TagListParamsSortDescName     TagListParamsSort = "DESC_NAME"
	TagListParamsSortAscLastUsed  TagListParamsSort = "ASC_LAST_USED"
	TagListParamsSortDescLastUsed TagListParamsSort = "DESC_LAST_USED"
)

// Filter tags by who added them.
//
// - `user` — tags added by users through the dashboard or the API
// - `ai` — tags added by AI auto-tagging extensions
// - `all` — both
type TagListParamsTagSource string

const (
	TagListParamsTagSourceUser TagListParamsTagSource = "user"
	TagListParamsTagSourceAI   TagListParamsTagSource = "ai"
	TagListParamsTagSourceAll  TagListParamsTagSource = "all"
)
