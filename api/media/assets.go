package media

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/creasty/defaults"
	"github.com/imagekit-developer/imagekit-go/api"
	"github.com/imagekit-developer/imagekit-go/extension"
	"gopkg.in/validator.v2"
)

// AssetType represents type of media library asset in request filter.
type AssetType string

const (
	File        AssetType = "file"
	FileVersion AssetType = "file-version"
	Folder      AssetType = "folder"
)

// Sort specifies sort order for Assets results data.
type Sort string

const (
	AscName     Sort = "ASC_NAME"
	DescName    Sort = "DESC_NAME"
	AscCreated  Sort = "ASC_CREATED"
	DescCreated Sort = "ASC_CREATED"
	AscHeight   Sort = "ASC_HEIGHT"
	DescHeight  Sort = "DESC_HEIGHT"
	AscWidth    Sort = "ASC_WIDTH"
	DescWidth   Sort = "DESC_WIDTH"
	AscSize     Sort = "ASC_SIZE"
	DescSize    Sort = "DESC_SIZE"
)

// FileType represents all, image or non-image etc type in request filter.
type FileType string

const (
	All      FileType = "all"
	Image    FileType = "image"
	NonImage FileType = "non-image"
)

// AssetParam struct is a parameter type to Assets() function to search / list media library assets.
type AssetsParam struct {
	Type        AssetType `json:"type,omitempty"`
	Sort        Sort      `json:"sort,omitempty"`
	Path        string    `json:"path,omitempty"`
	SearchQuery string    `json:"searchQuery,omitempty"`
	FileType    FileType  `json:"fileType,omitempty"`
	Tags        string    `json:"tags,omitempty"`
	Limit       int       `json:"limit,omitempty"`
	Skip        int       `json:"skip,omitempty"`
}

// AssetVersionsParam represents filter for getting asset version
type AssetVersionsParam struct {
	FileId    string `validate:"nonzero" json:"fileId"`
	VersionId string `json:"versionId,omitempty"`
}

// Asset represents media library asset details.
type Asset struct {
	FileId            string `json:"fileId"`
	Name              string `json:"name"`
	FilePath          string `json:"filePath"`
	Tags              []string
	AITags            []map[string]any  `json:"AITags"`
	VersionInfo       map[string]string `json:"versionInfo"`
	IsPrivateFile     *bool             `json:"isPrivateFile"`
	CustomCoordinates *string           `json:"customCoordinates"`
	Url               string            `json:"url"`
	Thumbnail         string            `json:"thumbnail"`
	FileType          FileType          `json:"fileType"`
	Mime              string            `json:"mime"`
	Height            int               `json:"height"`
	Width             int               `json:"Width"`
	Size              uint64            `json:"size"`
	HasAlpha          bool              `json:"hasAlpha"`
	CustomMetadata    map[string]any    `json:"customMetadata,omitempty"`
	EmbeddedMetadata  map[string]any    `json:"embeddedMetadata"`
	CreatedAt         time.Time         `json:"createdAt"`
	UpdatedAt         time.Time         `json:"updatedAt"`
}

// AssetsResponse represents response type of Assets().
type AssetsResponse struct {
	Data []Asset
	api.Response
}

// AssetResponse represents response type of AssetById().
type AssetResponse struct {
	Data Asset
	api.Response
}

// UpdateAssetParam represents asset attributes to update
type UpdateAssetParam struct {
	RemoveAITags      []string               `json:"removeAITags,omitempty"`
	WebhookUrl        string                 `json:"webhookUrl,omitempty"`
	Extensions        []extension.IExtension `json:"extensions,omitempty"`
	Tags              []string               `json:"tags,omitempty"`
	CustomCoordinates string                 `json:"customCoordinates,omitempty"`
	CustomMetadata    map[string]any         `json:"customMetadata,omitempty"`
}

// TagsParam represents parameters to add tags to bulk assets
type TagsParam struct {
	FileIds []string `json:"fileIds"`
	Tags    []string `json:"tags"`
}

// AITagsParam represents parameters to add AI tags to bulk assets
type AITagsParam struct {
	FileIds []string `json:"fileIds"`
	AITags  []string `json:"AITags"`
}

// UpdatedIds represents response to tags update calls
type UpdatedIds struct {
	FileIds []string `json:"successfullyUpdatedFileIds"`
}

// TagsResponse represents response to add tags to bulk assets. Contains fileIds in Data
type TagsResponse struct {
	Data UpdatedIds
	api.Response
}

// FileIdsParam is a struct to hold slice of file ids to pass as a parameter.
type FileIdsParam struct {
	FileIds []string `validate:"nonzero" json:"fileIds"`
}

// DeletedIds is a struct to hold slice of successfully deleted assets ids.
type DeletedIds struct {
	FileIds []string            `json:"successfullyDeletedFileIds"`
	Errors  []map[string]string `json:"errors"`
}

// DeleteAssetsResponse represents response to delete assets api which includes ids of deleted assets.
type DeleteAssetsResponse struct {
	Data DeletedIds
	api.Response
}

// CopyAssetParam represents parameters to copy asset api
type CopyAssetParam struct {
	SourcePath          string `validate:"nonzero" json:"sourceFilePath"`
	DestinationPath     string `validate:"nonzero" json:"destinationPath"`
	IncludeFileVersions bool   `json:"includeFileVersions"`
}

// MoveAssetParam represents parameters to move asset api
type MoveAssetParam struct {
	SourcePath      string `validate:"nonzero" json:"sourceFilePath"`
	DestinationPath string `validate:"nonzero" json:"destinationPath"`
}

// RenameAssetParam represents parameter to rename asset api
type RenameAssetParam struct {
	FilePath    string `validate:"nonzero" json:"filePath"`
	NewFileName string `validate:"nonzero" json:"newFileName"`
	PurgeCache  bool   `json:"purgeCache"`
}

// PurgeRequestId contains purge request ids
type PurgeRequestId struct {
	RequestId string `json:"purgeRequestId"`
}

// RenameAssetResponse represents response struct of rename asset api
type RenameAssetResponse struct {
	Data PurgeRequestId
	api.Response
}

// JobStatus represents response Data to job status api
type JobStatus struct {
	JobId  string `json:"jobId"`
	Type   string `json:"type"`
	Status string `json:"status"`
}

// JobStatusResponse represents response to job status api
type JobStatusResponse struct {
	Data JobStatus
	api.Response
}

type ErrorMissingFileIds struct {
	Message        string   `json:"message"`
	MissingFileIds []string `json:"missingFileIds"`
	err            error
}

func (e *ErrorMissingFileIds) Error() string {
	return fmt.Sprintf("%s, %s", e.Message, strings.Join(e.MissingFileIds, ","))
}
func (e *ErrorMissingFileIds) Unwrap() error {
	return e.err
}

type ErrorPartialSuccess struct {
	UpdatedFileIds []string            `json:"successfullyUpdatedFileIds"`
	Errors         []map[string]string `json:"errors"`
}

func (e *ErrorPartialSuccess) Error() string {
	return fmt.Sprintf("%v", e.Errors)
}

// Assets retrieves media library assets. Filter options can be supplied as AssetsParams.
func (m *API) Assets(ctx context.Context, params AssetsParam) (*AssetsResponse, error) {
	if err := defaults.Set(&params); err != nil {
		return nil, err
	}

	values, err := api.StructToParams(params)
	if err != nil {
		return nil, err
	}

	var query = values.Encode()

	if query != "" {
		query = "?" + query
	}

	resp, err := m.get(ctx, "files"+query)
	defer api.DeferredBodyClose(resp)

	response := &AssetsResponse{}
	api.SetResponseMeta(resp, response)

	if err != nil {
		return response, err
	}

	if resp.StatusCode != 200 {
		err = response.ParseError()
	} else {
		err = json.Unmarshal(response.Body(), &response.Data)
	}

	return response, err
}

// AssetById returns details of single asset by provided id
func (m *API) AssetById(ctx context.Context, fileId string) (*AssetResponse, error) {
	resp, err := m.get(ctx, fmt.Sprintf("files/%s/details", fileId))

	defer api.DeferredBodyClose(resp)

	response := &AssetResponse{}

	api.SetResponseMeta(resp, response)

	if err != nil {
		return response, err
	}

	if resp.StatusCode != 200 {
		err = response.ParseError()
	} else {
		err = json.Unmarshal(response.Body(), &response.Data)
	}
	return response, err
}

// AssetVersions fetches given file version specified by version id or all versions if versionId not supplied
func (m *API) AssetVersions(ctx context.Context, params AssetVersionsParam) (*AssetsResponse, error) {
	parts := []string{"files", params.FileId, "versions"}
	if params.VersionId != "" {
		parts = append(parts, params.VersionId)
	}

	if err := validator.Validate(&params); err != nil {
		return nil, err
	}

	resp, err := m.get(ctx, strings.Join(parts, "/"))
	defer api.DeferredBodyClose(resp)

	response := &AssetsResponse{}
	api.SetResponseMeta(resp, response)

	if err != nil {
		return response, err
	}

	if resp.StatusCode != 200 {
		err = response.ParseError()
	} else {
		if params.VersionId == "" {
			err = json.Unmarshal(response.Body(), &response.Data)
		} else {
			var asset = Asset{}
			if err = json.Unmarshal(response.Body(), &asset); err == nil {
				response.Data = []Asset{asset}
			}
		}
	}
	return response, err
}

// UpdateAsset updates single asset properties specified by UpdateAssetParam
func (m *API) UpdateAsset(ctx context.Context, fileId string, params UpdateAssetParam) (*AssetResponse, error) {
	response := &AssetResponse{}
	var err error

	if fileId == "" {
		return nil, errors.New("fileId can not be empty")
	}

	resp, err := m.patch(ctx, fmt.Sprintf("files/%s/details", fileId), params)

	defer api.DeferredBodyClose(resp)

	api.SetResponseMeta(resp, response)

	if err != nil {
		return response, err
	}

	if resp.StatusCode != 200 {
		err = response.ParseError()
	} else {
		err = json.Unmarshal(response.Body(), &response.Data)
	}
	return response, err
}

// AddTags assigns tags to bulk files specified by FileIds
func (m *API) AddTags(ctx context.Context, params TagsParam) (*TagsResponse, error) {
	response := &TagsResponse{}
	var err error

	resp, err := m.post(ctx, "files/addTags", params)
	defer api.DeferredBodyClose(resp)

	api.SetResponseMeta(resp, response)

	if err != nil {
		return response, err
	}

	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(response.Body(), &response.Data)
		break
	case 404:
		var errMissing = &ErrorMissingFileIds{err: api.ErrNotFound}
		var err = json.Unmarshal(response.Body(), errMissing)
		if err != nil {
			return nil, err
		}
		//errMissing.err = api.ErrNotFound
		return nil, errMissing
	case 207:
		var errPartial = &ErrorPartialSuccess{}
		var err = json.Unmarshal(response.Body(), errPartial)
		if err != nil {
			return nil, err
		}

		return nil, errPartial
	default:
		err = response.ParseError()
	}
	return response, err
}

// RemoveTags removes tags from bulk files specified by FileIds
func (m *API) RemoveTags(ctx context.Context, params TagsParam) (*TagsResponse, error) {
	response := &TagsResponse{}
	var err error

	resp, err := m.post(ctx, "files/removeTags", params)
	defer api.DeferredBodyClose(resp)

	api.SetResponseMeta(resp, response)

	if err != nil {
		return response, err
	}

	switch resp.StatusCode {
	case 200:
		err = json.Unmarshal(response.Body(), &response.Data)
		break
	case 404:
		var errMissing = &ErrorMissingFileIds{err: api.ErrNotFound}
		var err = json.Unmarshal(response.Body(), errMissing)
		if err != nil {
			return nil, err
		}
		return nil, errMissing
	case 207:
		var errPartial = &ErrorPartialSuccess{}
		var err = json.Unmarshal(response.Body(), errPartial)
		if err != nil {
			return nil, err
		}

		return nil, errPartial
	default:
		err = response.ParseError()
	}
	return response, err
}

// RemoveAITags removes tags from bulk files specified by FileIds
func (m *API) RemoveAITags(ctx context.Context, params AITagsParam) (*TagsResponse, error) {
	response := &TagsResponse{}
	var err error

	resp, err := m.post(ctx, "files/removeAITags", params)
	defer api.DeferredBodyClose(resp)

	api.SetResponseMeta(resp, response)

	if err != nil {
		return response, err
	}

	if resp.StatusCode != 200 {
		err = response.ParseError()
	} else {
		err = json.Unmarshal(response.Body(), &response.Data)
	}

	return response, err
}

// DeleteAsset removes an asset by FileId from media library
func (m *API) DeleteAsset(ctx context.Context, fileId string) (*api.Response, error) {
	var err error
	response := &api.Response{}

	if fileId == "" {
		return nil, errors.New("fileId can not be empty")
	}

	resp, err := m.delete(ctx, "files/"+fileId, nil)
	defer api.DeferredBodyClose(resp)

	api.SetResponseMeta(resp, response)

	if err != nil {
		return response, err
	}

	if resp.StatusCode != 204 {
		err = response.ParseError()
	}
	return response, err
}

// DeleteAssetVersion removes given asset version
func (m *API) DeleteAssetVersion(ctx context.Context, fileId string, versionId string) (*api.Response, error) {
	var err error
	response := &api.Response{}

	if fileId == "" {
		return nil, errors.New("fileId can not be empty")
	}

	if versionId == "" {
		return nil, errors.New("versionId can not be empty")
	}

	resp, err := m.delete(ctx, fmt.Sprintf("files/%s/versions/%s", fileId, versionId), nil)
	defer api.DeferredBodyClose(resp)

	api.SetResponseMeta(resp, response)

	if err != nil {
		return response, err
	}

	if resp.StatusCode != 204 {
		err = response.ParseError()
	}
	return response, err
}

// DeleteBulkAssets deletes multiple assets from media library
func (m *API) DeleteBulkAssets(ctx context.Context, param FileIdsParam) (*DeleteAssetsResponse, error) {
	var err error
	response := &DeleteAssetsResponse{}

	if err = validator.Validate(&param); err != nil {
		return nil, err
	}

	resp, err := m.post(ctx, "files/batch/deleteByFileIds", param)
	defer api.DeferredBodyClose(resp)

	api.SetResponseMeta(resp, response)

	if err != nil {
		return response, err
	}

	if resp.StatusCode == 200 || resp.StatusCode == 207 {
		err = json.Unmarshal(response.Body(), &response.Data)
	} else {
		err = response.ParseError()
	}
	return response, err
}

// CopyAsset copies an asset to target path
func (m *API) CopyAsset(ctx context.Context, param CopyAssetParam) (*api.Response, error) {
	var err error

	response := &api.Response{}

	if err = validator.Validate(&param); err != nil {
		return nil, err
	}

	resp, err := m.post(ctx, "files/copy", &param)
	defer api.DeferredBodyClose(resp)

	api.SetResponseMeta(resp, response)

	if err != nil {
		return response, err
	}

	if resp.StatusCode != 204 {
		err = response.ParseError()
	}
	return response, err
}

// MoveAsset moves an asset to target path
func (m *API) MoveAsset(ctx context.Context, param MoveAssetParam) (*api.Response, error) {
	var err error

	response := &api.Response{}

	if err = validator.Validate(&param); err != nil {
		return nil, err
	}

	resp, err := m.post(ctx, "files/move", &param)
	defer api.DeferredBodyClose(resp)

	api.SetResponseMeta(resp, response)

	if err != nil {
		return response, err
	}

	if resp.StatusCode != 204 {
		err = response.ParseError()
	}
	return response, err
}

// RenameAsset renames an asset to new name as specified in RenameAssetParam struct and optionally includes purge request id
func (m *API) RenameAsset(ctx context.Context, param RenameAssetParam) (*RenameAssetResponse, error) {
	var err error
	var response = &RenameAssetResponse{}

	if err = validator.Validate(&param); err != nil {
		return nil, err
	}

	resp, err := m.put(ctx, "files/rename", &param)
	defer api.DeferredBodyClose(resp)

	api.SetResponseMeta(resp, response)

	if err != nil {
		return response, err
	}

	if resp.StatusCode != 200 {
		err = response.ParseError()
	} else {
		err = json.Unmarshal(response.Body(), &response.Data)
	}

	return response, err
}

// RestoreVersion sets specified verison of the asset as current version
func (m *API) RestoreVersion(ctx context.Context, param AssetVersionsParam) (*AssetResponse, error) {
	var err error
	var response = &AssetResponse{}

	if err = validator.Validate(&param); err != nil {
		return nil, err
	}

	resp, err := m.delete(ctx, fmt.Sprintf("files/%s/versions/%s/restore",
		param.FileId, param.VersionId), nil)

	api.SetResponseMeta(resp, response)

	if resp.StatusCode != 200 {
		err = response.ParseError()
	} else {
		err = json.Unmarshal(response.Body(), &response.Data)
	}
	return response, err
}

func (m *API) BulkJobStatus(ctx context.Context, jobId string) (*JobStatusResponse, error) {
	var err error
	var response = &JobStatusResponse{}

	if jobId == "" {
		return nil, errors.New("jobId can not be blank")
	}

	resp, err := m.get(ctx, "bulkJobs/"+jobId)
	defer api.DeferredBodyClose(resp)
	api.SetResponseMeta(resp, response)

	if err != nil {
		return response, err
	}

	if resp.StatusCode != 200 {
		err = response.ParseError()
	} else {
		err = json.Unmarshal(response.Body(), &response.Data)
	}
	return response, err
}
