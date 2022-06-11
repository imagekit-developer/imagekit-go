package media

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/creasty/defaults"
	"github.com/dhaval070/imagekit-go/api"
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
	AscNAME     Sort = "ASC_NAME"
	DescNAME    Sort = "DESC_NAME"
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
	Type        AssetType `default:"file" json:"type"`
	Sort        Sort      `default:"ASC_CREATED" json:"sort"`
	Path        string    `default:"/" json:"path"`
	SearchQuery string    `default:"" json:"searchQuery"`
	FileType    FileType  `default:"all" json:"fileType"`
	Limit       int       `default:"1000" validate:"min=1,max=1000" json:"limit"`
	Skip        int       `default:"0" validate:"min=0" json:"skip"`
}

// AssetByIdParam struct is a parameter type to AssetsById() function to retrieve single asset details.
type AssetByIdParam struct {
	FileId string `validate:"nonzero" json:"fileId"`
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
	AITags            []map[string]string    `json:"AITags"`
	VersionInfo       map[string]string      `json:"versionInfo"`
	IsPrivateFile     *bool                  `json:"isPrivateFile"`
	CustomCoordinates *string                `json:"customCoordinates"`
	Url               string                 `json:"url"`
	Thumbnail         string                 `json:"thumbnail"`
	FileType          FileType               `json:"fileType"`
	Mime              string                 `json:"mime"`
	Height            int                    `json:"height"`
	Width             int                    `json:"Width"`
	Size              uint64                 `json:"size"`
	HasAlpha          bool                   `json:"hasAlpha"`
	CustomMetadata    map[string]interface{} `json:"customMetadata,omitempty"`
	EmbeddedMetadata  map[string]interface{} `json:"embeddedMetadata"`
	CreatedAt         time.Time              `json:"createdAt"`
	UpdatedAt         time.Time              `json:"updatedAt"`
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
	Extensions        map[string]interface{} `json:"extensions,omitempty"`
	Tags              []string               `json:"tags,omitempty"`
	CustomCoordinates string                 `json:"customCoordinates,omitempty"`
	CustomMetadata    map[string]interface{} `json:"customMetadata,omitempty"`
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
	FileIds []string `json:"successfullyDeletedFileIds"`
}

// DeleteAssetsResponse represents response to delete assets api which includes ids of deleted assets.
type DeleteAssetsResponse struct {
	Data DeletedIds
	api.Response
}

// CopyAssetParam represents parameters to copy asset api
type CopyAssetParam struct {
	SourcePath      string `validate:"nonzero" json:"sourceFilePath"`
	DestinationPath string `validate:"nonzero" json:"destinationPath"`
	IncludeVersions bool   `json:"includeVersions"`
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

// CreateFolderParam represents parameter to create folder api
type CreateFolderParam struct {
	FolderName       string `validate:"nonzero" json:"folderName"`
	ParentFolderPath string `validate:"nonzero" json:"parentFolderPath"`
}

type DeleteFolderParam struct {
	FolderPath string `validate:"nonzero" json:"folderPath"`
}

// Assets retrieves media library assets. Filter options can be supplied as AssetsParams.
func (m *API) Assets(ctx context.Context, params AssetsParam) (*AssetsResponse, error) {
	if err := defaults.Set(&params); err != nil {
		return nil, err
	}

	if err := validator.Validate(params); err != nil {
		return nil, err
	}

	values, err := api.StructToParams(params)
	if err != nil {
		return nil, err
	}

	resp, err := m.Get(ctx, "files?"+values.Encode())
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
func (m *API) AssetById(ctx context.Context, params AssetByIdParam) (*AssetResponse, error) {
	if err := validator.Validate(params); err != nil {
		return nil, err
	}

	resp, err := m.Get(ctx, fmt.Sprintf("files/%s/details", params.FileId))

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

	resp, err := m.Get(ctx, strings.Join(parts, "/"))
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

	resp, err := m.Patch(ctx, fmt.Sprintf("files/%s/details", fileId), params)

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

	resp, err := m.Post(ctx, "files/addTags", params)
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
	log.Println(response.ResponseMetaData.StatusCode)
	return response, err
}

// RemoveTags removes tags from bulk files specified by FileIds
func (m *API) RemoveTags(ctx context.Context, params TagsParam) (*TagsResponse, error) {
	response := &TagsResponse{}
	var err error

	resp, err := m.Post(ctx, "files/removeTags", params)
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

// RemoveAITags removes tags from bulk files specified by FileIds
func (m *API) RemoveAITags(ctx context.Context, params AITagsParam) (*TagsResponse, error) {
	response := &TagsResponse{}
	var err error

	resp, err := m.Post(ctx, "files/removeAITags", params)
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

	resp, err := m.Delete(ctx, "files/"+fileId, nil)
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

	resp, err := m.Delete(ctx, fmt.Sprintf("files/%s/versions/%s", fileId, versionId), nil)
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

	resp, err := m.Post(ctx, "files/batch/deleteByFileIds", param)
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

// CopyAsset copies an asset to target path
func (m *API) CopyAsset(ctx context.Context, param CopyAssetParam) (*api.Response, error) {
	var err error

	response := &api.Response{}

	if err = validator.Validate(&param); err != nil {
		return nil, err
	}

	resp, err := m.Post(ctx, "files/copy", &param)
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

	resp, err := m.Post(ctx, "files/move", &param)
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

	resp, err := m.Put(ctx, "files/rename", &param)
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

	resp, err := m.Delete(ctx, fmt.Sprintf("files/%s/versions/%s/restore",
		param.FileId, param.VersionId), nil)

	api.SetResponseMeta(resp, response)

	if resp.StatusCode != 200 {
		err = response.ParseError()
	} else {
		err = json.Unmarshal(response.Body(), &response.Data)
	}
	return response, err
}
