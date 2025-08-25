// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"

	"github.com/stainless-sdks/imagekit-go/internal/apijson"
	"github.com/stainless-sdks/imagekit-go/packages/param"
	"github.com/stainless-sdks/imagekit-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// Object containing details of a file or file version.
type File struct {
	// An array of tags assigned to the file by auto tagging.
	AITags []FileAITag `json:"AITags,nullable"`
	// Date and time when the file was uploaded. The date and time is in ISO8601
	// format.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// An string with custom coordinates of the file.
	CustomCoordinates string `json:"customCoordinates,nullable"`
	// An object with custom metadata for the file.
	CustomMetadata map[string]any `json:"customMetadata"`
	// Optional text to describe the contents of the file. Can be set by the user or
	// the ai-auto-description extension.
	Description string `json:"description"`
	// Unique identifier of the asset.
	FileID string `json:"fileId"`
	// Path of the file. This is the path you would use in the URL to access the file.
	// For example, if the file is at the root of the media library, the path will be
	// `/file.jpg`. If the file is inside a folder named `images`, the path will be
	// `/images/file.jpg`.
	FilePath string `json:"filePath"`
	// Type of the file. Possible values are `image`, `non-image`.
	FileType string `json:"fileType"`
	// Specifies if the image has an alpha channel.
	HasAlpha bool `json:"hasAlpha"`
	// Height of the file.
	Height float64 `json:"height"`
	// Specifies if the file is private or not.
	IsPrivateFile bool `json:"isPrivateFile"`
	// Specifies if the file is published or not.
	IsPublished bool `json:"isPublished"`
	// MIME type of the file.
	Mime string `json:"mime"`
	// Name of the asset.
	Name string `json:"name"`
	// Size of the file in bytes.
	Size float64 `json:"size"`
	// An array of tags assigned to the file. Tags are used to search files in the
	// media library.
	Tags []string `json:"tags,nullable"`
	// URL of the thumbnail image. This URL is used to access the thumbnail image of
	// the file in the media library.
	Thumbnail string `json:"thumbnail" format:"uri"`
	// Type of the asset.
	//
	// Any of "file", "file-version".
	Type FileType `json:"type"`
	// Date and time when the file was last updated. The date and time is in ISO8601
	// format.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// URL of the file.
	URL string `json:"url" format:"uri"`
	// An object with details of the file version.
	VersionInfo FileVersionInfo `json:"versionInfo"`
	// Width of the file.
	Width float64 `json:"width"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
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
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r File) RawJSON() string { return r.JSON.raw }
func (r *File) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileAITag struct {
	// Confidence score of the tag.
	Confidence float64 `json:"confidence"`
	// Name of the tag.
	Name string `json:"name"`
	// Source of the tag. Possible values are `google-auto-tagging` and
	// `aws-auto-tagging`.
	Source string `json:"source"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Confidence  respjson.Field
		Name        respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileAITag) RawJSON() string { return r.JSON.raw }
func (r *FileAITag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the asset.
type FileType string

const (
	FileTypeFile        FileType = "file"
	FileTypeFileVersion FileType = "file-version"
)

// An object with details of the file version.
type FileVersionInfo struct {
	// Unique identifier of the file version.
	ID string `json:"id"`
	// Name of the file version.
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FileVersionInfo) RawJSON() string { return r.JSON.raw }
func (r *FileVersionInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Folder struct {
	// Date and time when the folder was created. The date and time is in ISO8601
	// format.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Unique identifier of the asset.
	FolderID string `json:"folderId"`
	// Path of the folder. This is the path you would use in the URL to access the
	// folder. For example, if the folder is at the root of the media library, the path
	// will be /folder. If the folder is inside another folder named images, the path
	// will be /images/folder.
	FolderPath string `json:"folderPath"`
	// Name of the asset.
	Name string `json:"name"`
	// Type of the asset.
	//
	// Any of "folder".
	Type FolderType `json:"type"`
	// Date and time when the folder was last updated. The date and time is in ISO8601
	// format.
	UpdatedAt time.Time `json:"updatedAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		FolderID    respjson.Field
		FolderPath  respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Folder) RawJSON() string { return r.JSON.raw }
func (r *Folder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the asset.
type FolderType string

const (
	FolderTypeFolder FolderType = "folder"
)
