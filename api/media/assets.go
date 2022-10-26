package media

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/imagekit-developer/imagekit-go/api"
	"github.com/imagekit-developer/imagekit-go/api/extension"
	"gopkg.in/validator.v2"
)

// ListType represents type of media library files in request filter.
type ListType string

const (
	ListFile          ListType = "file"
	ListFTFileVersion ListType = "file-version"
	ListFolder        ListType = "folder"
)

// Sort specifies sort order for ListFiles results data.
type Sort string

const (
	AscName     Sort = "ASC_NAME"
	DescName    Sort = "DESC_NAME"
	AscCreated  Sort = "ASC_CREATED"
	DescCreated Sort = "DESC_CREATED"
	AscUpdated  Sort = "ASC_UPDATED"
	DescUpdated Sort = "DESC_UPDATED"
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

// FilesParam struct is a parameter type to ListFiles() function to search / list media library files.
type FilesParam struct {
	Type        ListType `json:"type,omitempty"`
	Sort        Sort     `json:"sort,omitempty"`
	Path        string   `json:"path,omitempty"`
	SearchQuery string   `json:"searchQuery,omitempty"`
	FileType    FileType `json:"fileType,omitempty"`
	Tags        string   `json:"tags,omitempty"`
	Limit       int      `json:"limit,omitempty"`
	Skip        int      `json:"skip,omitempty"`
}

// FileVersionsParam represents filter for getting file's version
type FileVersionsParam struct {
	FileId    string `validate:"nonzero" json:"fileId"`
	VersionId string `json:"versionId,omitempty"`
}

// File represents media library File details.
type File struct {
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

// FilesResponse represents response type of Files().
type FilesResponse struct {
	Data []File
	api.Response
}

// FileResponse represents response type of FileById().
type FileResponse struct {
	Data File
	api.Response
}

// UpdateFileParam represents file attributes to update
type UpdateFileParam struct {
	RemoveAITags      []string               `json:"removeAITags,omitempty"`
	WebhookUrl        string                 `json:"webhookUrl,omitempty"`
	Extensions        []extension.IExtension `json:"extensions,omitempty"`
	Tags              []string               `json:"tags,omitempty"`
	CustomCoordinates string                 `json:"customCoordinates,omitempty"`
	CustomMetadata    map[string]any         `json:"customMetadata,omitempty"`
}

// TagsParam represents parameters to add tags to bulk files
type TagsParam struct {
	FileIds []string `json:"fileIds"`
	Tags    []string `json:"tags"`
}

// AITagsParam represents parameters to add AI tags to bulk files
type AITagsParam struct {
	FileIds []string `json:"fileIds"`
	AITags  []string `json:"AITags"`
}

// UpdatedIds represents response to tags update calls
type UpdatedIds struct {
	FileIds []string `json:"successfullyUpdatedFileIds"`
}

// TagsResponse represents response to add tags to bulk files. Contains fileIds in Data
type TagsResponse struct {
	Data UpdatedIds
	api.Response
}

// FileIdsParam is a struct to hold slice of file ids to pass as a parameter.
type FileIdsParam struct {
	FileIds []string `validate:"nonzero" json:"fileIds"`
}

// DeletedIds is a struct to hold slice of successfully deleted files ids.
type DeletedIds struct {
	FileIds []string            `json:"successfullyDeletedFileIds"`
	Errors  []map[string]string `json:"errors"`
}

// DeleteFilessResponse represents response to delete files api which includes ids of deleted files.
type DeleteFilesResponse struct {
	Data DeletedIds
	api.Response
}

// CopyFileParam represents parameters to copy files api
type CopyFileParam struct {
	SourcePath          string `validate:"nonzero" json:"sourceFilePath"`
	DestinationPath     string `validate:"nonzero" json:"destinationPath"`
	IncludeFileVersions bool   `json:"includeFileVersions"`
}

// MoveFileParam represents parameters to move file api
type MoveFileParam struct {
	SourcePath      string `validate:"nonzero" json:"sourceFilePath"`
	DestinationPath string `validate:"nonzero" json:"destinationPath"`
}

// RenameFileParam represents parameter to rename file api
type RenameFileParam struct {
	FilePath    string `validate:"nonzero" json:"filePath"`
	NewFileName string `validate:"nonzero" json:"newFileName"`
	PurgeCache  bool   `json:"purgeCache"`
}

// PurgeRequestId contains purge request ids
type PurgeRequestId struct {
	RequestId string `json:"purgeRequestId"`
}

// RenameFileResponse represents response struct of rename File api
type RenameFileResponse struct {
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

// Files retrieves media library files. Filter options can be supplied as FilesParams.
func (m *API) Files(ctx context.Context, params FilesParam) (*FilesResponse, error) {
	values, err := api.StructToParams(params)
	if err != nil {
		return nil, err
	}

	var query = values.Encode()

	if query != "" {
		query = "?" + query
	}

	response := &FilesResponse{}

	resp, err := m.get(ctx, "files"+query, response)

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

// FileById returns details of single file by provided id
func (m *API) FileById(ctx context.Context, fileId string) (*FileResponse, error) {
	response := &FileResponse{}

	resp, err := m.get(ctx, fmt.Sprintf("files/%s/details", fileId), response)

	defer api.DeferredBodyClose(resp)

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

// FileVersions fetches given file version specified by version id or all versions if versionId not supplied
func (m *API) FileVersions(ctx context.Context, params FileVersionsParam) (*FilesResponse, error) {
	parts := []string{"files", params.FileId, "versions"}
	if params.VersionId != "" {
		parts = append(parts, params.VersionId)
	}

	if err := validator.Validate(&params); err != nil {
		return nil, err
	}

	response := &FilesResponse{}

	resp, err := m.get(ctx, strings.Join(parts, "/"), response)

	if err != nil {
		return response, err
	}

	if resp.StatusCode != 200 {
		err = response.ParseError()
	} else {
		if params.VersionId == "" {
			err = json.Unmarshal(response.Body(), &response.Data)
		} else {
			var file = File{}
			if err = json.Unmarshal(response.Body(), &file); err == nil {
				response.Data = []File{file}
			}
		}
	}
	return response, err
}

// UpdateFile updates single file properties specified by UpdateFileParam
func (m *API) UpdateFile(ctx context.Context, fileId string, params UpdateFileParam) (*FileResponse, error) {
	response := &FileResponse{}
	var err error

	if fileId == "" {
		return nil, errors.New("fileId can not be empty")
	}

	resp, err := m.patch(ctx, fmt.Sprintf("files/%s/details", fileId), params, response)

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

	resp, err := m.post(ctx, "files/addTags", params, response)

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

	resp, err := m.post(ctx, "files/removeTags", params, response)

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

	resp, err := m.post(ctx, "files/removeAITags", params, response)

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

// DeleteFile removes file by FileId from media library
func (m *API) DeleteFile(ctx context.Context, fileId string) (*api.Response, error) {
	var err error
	response := &api.Response{}

	if fileId == "" {
		return nil, errors.New("fileId can not be empty")
	}

	resp, err := m.delete(ctx, "files/"+fileId, nil, response)

	if err != nil {
		return response, err
	}

	if resp.StatusCode != 204 {
		err = response.ParseError()
	}
	return response, err
}

// DeleteFileVersion removes given file version
func (m *API) DeleteFileVersion(ctx context.Context, fileId string, versionId string) (*api.Response, error) {
	var err error
	response := &api.Response{}

	if fileId == "" {
		return nil, errors.New("fileId can not be empty")
	}

	if versionId == "" {
		return nil, errors.New("versionId can not be empty")
	}

	resp, err := m.delete(ctx, fmt.Sprintf("files/%s/versions/%s", fileId, versionId), nil, response)

	if err != nil {
		return response, err
	}

	if resp.StatusCode != 204 {
		err = response.ParseError()
	}
	return response, err
}

// DeleteBulkFiles deletes multiple files from media library
func (m *API) DeleteBulkFiles(ctx context.Context, param FileIdsParam) (*DeleteFilesResponse, error) {
	var err error
	response := &DeleteFilesResponse{}

	if err = validator.Validate(&param); err != nil {
		return nil, err
	}

	resp, err := m.post(ctx, "files/batch/deleteByFileIds", param, response)

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

// CopyFile copies a file to target path
func (m *API) CopyFile(ctx context.Context, param CopyFileParam) (*api.Response, error) {
	var err error

	response := &api.Response{}

	if err = validator.Validate(&param); err != nil {
		return nil, err
	}

	resp, err := m.post(ctx, "files/copy", &param, response)

	if err != nil {
		return response, err
	}

	if resp.StatusCode != 204 {
		err = response.ParseError()
	}
	return response, err
}

// MoveFile moves a file to target path
func (m *API) MoveFile(ctx context.Context, param MoveFileParam) (*api.Response, error) {
	var err error

	response := &api.Response{}

	if err = validator.Validate(&param); err != nil {
		return nil, err
	}

	resp, err := m.post(ctx, "files/move", &param, response)

	if err != nil {
		return response, err
	}

	if resp.StatusCode != 204 {
		err = response.ParseError()
	}
	return response, err
}

// RenameFile renames a file to new name as specified in RenameFileParam struct and optionally includes purge request id
func (m *API) RenameFile(ctx context.Context, param RenameFileParam) (*RenameFileResponse, error) {
	var err error
	var response = &RenameFileResponse{}

	if err = validator.Validate(&param); err != nil {
		return nil, err
	}

	resp, err := m.put(ctx, "files/rename", &param, response)
	defer api.DeferredBodyClose(resp)

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

// RestoreVersion sets specified verison of the file as current version
func (m *API) RestoreVersion(ctx context.Context, param FileVersionsParam) (*FileResponse, error) {
	var err error
	var response = &FileResponse{}

	if err = validator.Validate(&param); err != nil {
		return nil, err
	}

	resp, err := m.delete(ctx, fmt.Sprintf("files/%s/versions/%s/restore",
		param.FileId, param.VersionId), nil, response)

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

	resp, err := m.get(ctx, "bulkJobs/"+jobId, response)

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
