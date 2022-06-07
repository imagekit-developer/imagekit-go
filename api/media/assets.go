package media

import (
	"context"
	"log"

	"github.com/creasty/defaults"
	"gopkg.in/validator.v2"
)

type AssetType string

const (
	File        AssetType = "file"
	FileVersion AssetType = "file-version"
	Folder      AssetType = "folder"
)

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

type FileType string

const (
	All      FileType = "all"
	Image    FileType = "image"
	NonImage FileType = "non-image"
)

type AssetsParam struct {
	Type        AssetType `default:"file" json:"type"`
	Sort        Sort      `default:"ASC_CREATED" json:"sort"`
	Path        string    `default:"path" json:"path"`
	SearchQuery string    `default:"SearchQuery" json:"searchQuery"`
	FileType    FileType  `default:"all" json:"fileType"`
	Limit       int       `default:"1000" validate:"min=1,max=1000" json:"limit"`
	Skip        int       `default:0 validate:"min:0", son:"skip"`
}

type AssetsResult struct {
	FileId string `json:"fileId"`
}

// Assets lists media library assets. Filter options can be supplied as AssetsParams.
func (m *API) Assets(ctx context.Context, params AssetsParam) (*AssetsResult, error) {
	var assetsResult = &AssetsResult{}

	if err := defaults.Set(&params); err != nil {
		return nil, err
	}

	log.Println(params)

	if err := validator.Validate(params); err != nil {
		return nil, err
	}

	log.Println(assetsResult)
	return nil, nil
}
