// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/stainless-sdks/imagekit-go/internal/apijson"
	"github.com/stainless-sdks/imagekit-go/internal/apiquery"
	"github.com/stainless-sdks/imagekit-go/internal/requestconfig"
	"github.com/stainless-sdks/imagekit-go/option"
	"github.com/stainless-sdks/imagekit-go/packages/param"
	"github.com/stainless-sdks/imagekit-go/packages/respjson"
)

// AssetService contains methods and other services that help with interacting with
// the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssetService] method instead.
type AssetService struct {
	Options []option.RequestOption
}

// NewAssetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAssetService(opts ...option.RequestOption) (r AssetService) {
	r = AssetService{}
	r.Options = opts
	return
}

// This API can list all the uploaded files and folders in your ImageKit.io media
// library. In addition, you can fine-tune your query by specifying various filters
// by generating a query string in a Lucene-like syntax and provide this generated
// string as the value of the `searchQuery`.
func (r *AssetService) List(ctx context.Context, query AssetListParams, opts ...option.RequestOption) (res *[]AssetListResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/files"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// AssetListResponseUnion contains all possible properties and values from [File],
// [Folder].
//
// Use the [AssetListResponseUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AssetListResponseUnion struct {
	// This field is from variant [File].
	AITags    []FileAITag `json:"AITags"`
	CreatedAt time.Time   `json:"createdAt"`
	// This field is from variant [File].
	CustomCoordinates string `json:"customCoordinates"`
	// This field is from variant [File].
	CustomMetadata map[string]any `json:"customMetadata"`
	// This field is from variant [File].
	Description string `json:"description"`
	// This field is from variant [File].
	FileID string `json:"fileId"`
	// This field is from variant [File].
	FilePath string `json:"filePath"`
	// This field is from variant [File].
	FileType string `json:"fileType"`
	// This field is from variant [File].
	HasAlpha bool `json:"hasAlpha"`
	// This field is from variant [File].
	Height float64 `json:"height"`
	// This field is from variant [File].
	IsPrivateFile bool `json:"isPrivateFile"`
	// This field is from variant [File].
	IsPublished bool `json:"isPublished"`
	// This field is from variant [File].
	Mime string `json:"mime"`
	Name string `json:"name"`
	// This field is from variant [File].
	Size float64 `json:"size"`
	// This field is from variant [File].
	Tags []string `json:"tags"`
	// This field is from variant [File].
	Thumbnail string `json:"thumbnail"`
	// Any of nil, "folder".
	Type      string    `json:"type"`
	UpdatedAt time.Time `json:"updatedAt"`
	// This field is from variant [File].
	URL string `json:"url"`
	// This field is from variant [File].
	VersionInfo FileVersionInfo `json:"versionInfo"`
	// This field is from variant [File].
	Width float64 `json:"width"`
	// This field is from variant [Folder].
	FolderID string `json:"folderId"`
	// This field is from variant [Folder].
	FolderPath string `json:"folderPath"`
	JSON       struct {
		AITags            respjson.Field
		CreatedAt         respjson.Field
		CustomCoordinates respjson.Field
		CustomMetadata    respjson.Field
		Description       respjson.Field
		FileID            respjson.Field
		FilePath          respjson.Field
		FileType          respjson.Field
		HasAlpha          respjson.Field
		Height            respjson.Field
		IsPrivateFile     respjson.Field
		IsPublished       respjson.Field
		Mime              respjson.Field
		Name              respjson.Field
		Size              respjson.Field
		Tags              respjson.Field
		Thumbnail         respjson.Field
		Type              respjson.Field
		UpdatedAt         respjson.Field
		URL               respjson.Field
		VersionInfo       respjson.Field
		Width             respjson.Field
		FolderID          respjson.Field
		FolderPath        respjson.Field
		raw               string
	} `json:"-"`
}

func (u AssetListResponseUnion) AsFileFileVersion() (v File) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AssetListResponseUnion) AsFolder() (v Folder) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AssetListResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *AssetListResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AssetListParams struct {
	// The maximum number of results to return in response.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Folder path if you want to limit the search within a specific folder. For
	// example, `/sales-banner/` will only search in folder sales-banner.
	//
	// Note : If your use case involves searching within a folder as well as its
	// subfolders, you can use `path` parameter in `searchQuery` with appropriate
	// operator. Checkout
	// [Supported parameters](/docs/api-reference/digital-asset-management-dam/list-and-search-assets#supported-parameters)
	// for more information.
	Path param.Opt[string] `query:"path,omitzero" json:"-"`
	// Query string in a Lucene-like query language e.g. `createdAt > "7d"`.
	//
	// Note : When the searchQuery parameter is present, the following query parameters
	// will have no effect on the result:
	//
	// 1. `tags`
	// 2. `type`
	// 3. `name`
	//
	// [Learn more](/docs/api-reference/digital-asset-management-dam/list-and-search-assets#advanced-search-queries)
	// from examples.
	SearchQuery param.Opt[string] `query:"searchQuery,omitzero" json:"-"`
	// The number of results to skip before returning results.
	Skip param.Opt[int64] `query:"skip,omitzero" json:"-"`
	// Filter results by file type.
	//
	// - `all` — include all file types
	// - `image` — include only image files
	// - `non-image` — include only non-image files (e.g., JS, CSS, video)
	//
	// Any of "all", "image", "non-image".
	FileType AssetListParamsFileType `query:"fileType,omitzero" json:"-"`
	// Sort the results by one of the supported fields in ascending or descending
	// order.
	//
	// Any of "ASC_NAME", "DESC_NAME", "ASC_CREATED", "DESC_CREATED", "ASC_UPDATED",
	// "DESC_UPDATED", "ASC_HEIGHT", "DESC_HEIGHT", "ASC_WIDTH", "DESC_WIDTH",
	// "ASC_SIZE", "DESC_SIZE", "ASC_RELEVANCE", "DESC_RELEVANCE".
	Sort AssetListParamsSort `query:"sort,omitzero" json:"-"`
	// Filter results by asset type.
	//
	// - `file` — returns only files
	// - `file-version` — returns specific file versions
	// - `folder` — returns only folders
	// - `all` — returns both files and folders (excludes `file-version`)
	//
	// Any of "file", "file-version", "folder", "all".
	Type AssetListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AssetListParams]'s query parameters as `url.Values`.
func (r AssetListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter results by file type.
//
// - `all` — include all file types
// - `image` — include only image files
// - `non-image` — include only non-image files (e.g., JS, CSS, video)
type AssetListParamsFileType string

const (
	AssetListParamsFileTypeAll      AssetListParamsFileType = "all"
	AssetListParamsFileTypeImage    AssetListParamsFileType = "image"
	AssetListParamsFileTypeNonImage AssetListParamsFileType = "non-image"
)

// Sort the results by one of the supported fields in ascending or descending
// order.
type AssetListParamsSort string

const (
	AssetListParamsSortAscName       AssetListParamsSort = "ASC_NAME"
	AssetListParamsSortDescName      AssetListParamsSort = "DESC_NAME"
	AssetListParamsSortAscCreated    AssetListParamsSort = "ASC_CREATED"
	AssetListParamsSortDescCreated   AssetListParamsSort = "DESC_CREATED"
	AssetListParamsSortAscUpdated    AssetListParamsSort = "ASC_UPDATED"
	AssetListParamsSortDescUpdated   AssetListParamsSort = "DESC_UPDATED"
	AssetListParamsSortAscHeight     AssetListParamsSort = "ASC_HEIGHT"
	AssetListParamsSortDescHeight    AssetListParamsSort = "DESC_HEIGHT"
	AssetListParamsSortAscWidth      AssetListParamsSort = "ASC_WIDTH"
	AssetListParamsSortDescWidth     AssetListParamsSort = "DESC_WIDTH"
	AssetListParamsSortAscSize       AssetListParamsSort = "ASC_SIZE"
	AssetListParamsSortDescSize      AssetListParamsSort = "DESC_SIZE"
	AssetListParamsSortAscRelevance  AssetListParamsSort = "ASC_RELEVANCE"
	AssetListParamsSortDescRelevance AssetListParamsSort = "DESC_RELEVANCE"
)

// Filter results by asset type.
//
// - `file` — returns only files
// - `file-version` — returns specific file versions
// - `folder` — returns only folders
// - `all` — returns both files and folders (excludes `file-version`)
type AssetListParamsType string

const (
	AssetListParamsTypeFile        AssetListParamsType = "file"
	AssetListParamsTypeFileVersion AssetListParamsType = "file-version"
	AssetListParamsTypeFolder      AssetListParamsType = "folder"
	AssetListParamsTypeAll         AssetListParamsType = "all"
)
