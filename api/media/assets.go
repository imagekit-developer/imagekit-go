package media

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
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
	FileId string `json:"fileId"`
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

	url, err := url.Parse(m.Config.API.Prefix + "files?" + values.Encode())
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(m.Config.Cloud.PrivateKey, "")

	resp, err := m.Client.Do(req.WithContext(ctx))
	defer api.DeferredBodyClose(resp)

	assetsResp := &AssetsResponse{}
	api.SetResponseMeta(resp, assetsResp)

	if err != nil {
		return assetsResp, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return assetsResp, err
	}
	err = json.Unmarshal(body, &assetsResp.Data)

	return assetsResp, err
}

// AssetById returns details of single asset by provided id
func (m *API) AssetById(ctx context.Context, params AssetByIdParam) (*AssetByIdResponse, error) {
	url, err := url.Parse(fmt.Sprintf(m.Config.API.Prefix+"files/%s/details", params.FileId))

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(m.Config.Cloud.PrivateKey, "")

	resp, err := m.Client.Do(req.WithContext(ctx))
	defer api.DeferredBodyClose(resp)

	response := &AssetByIdResponse{}
	api.SetResponseMeta(resp, response)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(body, &response.Data)
	return response, err
}
