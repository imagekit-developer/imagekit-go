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

// AssetByIdResponse represents response type of AssetById().
type AssetByIdResponse struct {
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
func (m *API) AssetById(ctx context.Context, params AssetByIdParam) (*AssetByIdResponse, error) {
	if err := validator.Validate(params); err != nil {
		return nil, err
	}

	resp, err := m.Get(ctx, fmt.Sprintf("files/%s/details", params.FileId))

	defer api.DeferredBodyClose(resp)

	response := &AssetByIdResponse{}

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
func (m *API) UpdateAsset(ctx context.Context, fileId string, params UpdateAssetParam) (*AssetByIdResponse, error) {
	response := &AssetByIdResponse{}
	var err error

	if fileId == "" {
		return nil, errors.New("fileId can not be empty")
	}

	body, err := json.Marshal(&params)
	if err != nil {
		return nil, err
	}

	resp, err := m.Patch(ctx, fmt.Sprintf("files/%s/details", fileId), body)

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

	body, err := json.Marshal(&params)
	if err != nil {
		return nil, err
	}

	resp, err := m.Post(ctx, "files/addTags", body)
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

	body, err := json.Marshal(&params)
	if err != nil {
		return nil, err
	}

	resp, err := m.Post(ctx, "files/removeTags", body)
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

	body, err := json.Marshal(&params)
	if err != nil {
		return nil, err
	}

	resp, err := m.Post(ctx, "files/removeAITags", body)
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

	resp, err := m.Delete(ctx, "files/"+fileId)
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
