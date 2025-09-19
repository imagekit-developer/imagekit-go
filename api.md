# Shared Params Types

- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go/shared#ExtensionsParam">ExtensionsParam</a>

# CustomMetadataFields

Response Types:

- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#CustomMetadataField">CustomMetadataField</a>
- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#CustomMetadataFieldDeleteResponse">CustomMetadataFieldDeleteResponse</a>

Methods:

- <code title="post /v1/customMetadataFields">client.CustomMetadataFields.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#CustomMetadataFieldService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#CustomMetadataFieldNewParams">CustomMetadataFieldNewParams</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#CustomMetadataField">CustomMetadataField</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/customMetadataFields/{id}">client.CustomMetadataFields.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#CustomMetadataFieldService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#CustomMetadataFieldUpdateParams">CustomMetadataFieldUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#CustomMetadataField">CustomMetadataField</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/customMetadataFields">client.CustomMetadataFields.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#CustomMetadataFieldService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#CustomMetadataFieldListParams">CustomMetadataFieldListParams</a>) ([]<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#CustomMetadataField">CustomMetadataField</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/customMetadataFields/{id}">client.CustomMetadataFields.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#CustomMetadataFieldService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#CustomMetadataFieldDeleteResponse">CustomMetadataFieldDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Files

Params Types:

- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#UpdateFileRequestUnionParam">UpdateFileRequestUnionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#File">File</a>
- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#Folder">Folder</a>
- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#Metadata">Metadata</a>
- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileUpdateResponse">FileUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileCopyResponse">FileCopyResponse</a>
- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileMoveResponse">FileMoveResponse</a>
- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileRenameResponse">FileRenameResponse</a>
- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileUploadResponse">FileUploadResponse</a>

Methods:

- <code title="patch /v1/files/{fileId}/details">client.Files.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileUpdateParams">FileUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileUpdateResponse">FileUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/files/{fileId}">client.Files.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v1/files/copy">client.Files.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileService.Copy">Copy</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileCopyParams">FileCopyParams</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileCopyResponse">FileCopyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/files/{fileId}/details">client.Files.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#File">File</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/files/move">client.Files.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileService.Move">Move</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileMoveParams">FileMoveParams</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileMoveResponse">FileMoveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/files/rename">client.Files.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileService.Rename">Rename</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileRenameParams">FileRenameParams</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileRenameResponse">FileRenameResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /api/v1/files/upload">client.Files.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileService.Upload">Upload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileUploadParams">FileUploadParams</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileUploadResponse">FileUploadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Bulk

Response Types:

- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileBulkDeleteResponse">FileBulkDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileBulkAddTagsResponse">FileBulkAddTagsResponse</a>
- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileBulkRemoveAITagsResponse">FileBulkRemoveAITagsResponse</a>
- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileBulkRemoveTagsResponse">FileBulkRemoveTagsResponse</a>

Methods:

- <code title="post /v1/files/batch/deleteByFileIds">client.Files.Bulk.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileBulkService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileBulkDeleteParams">FileBulkDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileBulkDeleteResponse">FileBulkDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/files/addTags">client.Files.Bulk.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileBulkService.AddTags">AddTags</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileBulkAddTagsParams">FileBulkAddTagsParams</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileBulkAddTagsResponse">FileBulkAddTagsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/files/removeAITags">client.Files.Bulk.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileBulkService.RemoveAITags">RemoveAITags</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileBulkRemoveAITagsParams">FileBulkRemoveAITagsParams</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileBulkRemoveAITagsResponse">FileBulkRemoveAITagsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/files/removeTags">client.Files.Bulk.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileBulkService.RemoveTags">RemoveTags</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileBulkRemoveTagsParams">FileBulkRemoveTagsParams</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileBulkRemoveTagsResponse">FileBulkRemoveTagsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Versions

Response Types:

- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileVersionDeleteResponse">FileVersionDeleteResponse</a>

Methods:

- <code title="get /v1/files/{fileId}/versions">client.Files.Versions.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileVersionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#File">File</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/files/{fileId}/versions/{versionId}">client.Files.Versions.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileVersionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, versionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileVersionDeleteParams">FileVersionDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileVersionDeleteResponse">FileVersionDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/files/{fileId}/versions/{versionId}">client.Files.Versions.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileVersionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, versionID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileVersionGetParams">FileVersionGetParams</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#File">File</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/files/{fileId}/versions/{versionId}/restore">client.Files.Versions.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileVersionService.Restore">Restore</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, versionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileVersionRestoreParams">FileVersionRestoreParams</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#File">File</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Metadata

Methods:

- <code title="get /v1/files/{fileId}/metadata">client.Files.Metadata.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileMetadataService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#Metadata">Metadata</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/files/metadata">client.Files.Metadata.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileMetadataService.GetFromURL">GetFromURL</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FileMetadataGetFromURLParams">FileMetadataGetFromURLParams</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#Metadata">Metadata</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Assets

Response Types:

- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#AssetListResponseUnion">AssetListResponseUnion</a>

Methods:

- <code title="get /v1/files">client.Assets.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#AssetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#AssetListParams">AssetListParams</a>) ([]<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#AssetListResponseUnion">AssetListResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Cache

## Invalidation

Response Types:

- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#CacheInvalidationNewResponse">CacheInvalidationNewResponse</a>
- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#CacheInvalidationGetResponse">CacheInvalidationGetResponse</a>

Methods:

- <code title="post /v1/files/purge">client.Cache.Invalidation.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#CacheInvalidationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#CacheInvalidationNewParams">CacheInvalidationNewParams</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#CacheInvalidationNewResponse">CacheInvalidationNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/files/purge/{requestId}">client.Cache.Invalidation.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#CacheInvalidationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, requestID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#CacheInvalidationGetResponse">CacheInvalidationGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Folders

Response Types:

- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FolderNewResponse">FolderNewResponse</a>
- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FolderDeleteResponse">FolderDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FolderCopyResponse">FolderCopyResponse</a>
- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FolderMoveResponse">FolderMoveResponse</a>
- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FolderRenameResponse">FolderRenameResponse</a>

Methods:

- <code title="post /v1/folder">client.Folders.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FolderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FolderNewParams">FolderNewParams</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FolderNewResponse">FolderNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/folder">client.Folders.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FolderService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FolderDeleteParams">FolderDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FolderDeleteResponse">FolderDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/bulkJobs/copyFolder">client.Folders.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FolderService.Copy">Copy</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FolderCopyParams">FolderCopyParams</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FolderCopyResponse">FolderCopyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/bulkJobs/moveFolder">client.Folders.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FolderService.Move">Move</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FolderMoveParams">FolderMoveParams</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FolderMoveResponse">FolderMoveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/bulkJobs/renameFolder">client.Folders.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FolderService.Rename">Rename</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FolderRenameParams">FolderRenameParams</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FolderRenameResponse">FolderRenameResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Job

Response Types:

- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FolderJobGetResponse">FolderJobGetResponse</a>

Methods:

- <code title="get /v1/bulkJobs/{jobId}">client.Folders.Job.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FolderJobService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#FolderJobGetResponse">FolderJobGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Accounts

## Usage

Response Types:

- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#AccountUsageGetResponse">AccountUsageGetResponse</a>

Methods:

- <code title="get /v1/accounts/usage">client.Accounts.Usage.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#AccountUsageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#AccountUsageGetParams">AccountUsageGetParams</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#AccountUsageGetResponse">AccountUsageGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Origins

Params Types:

- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#OriginRequestUnionParam">OriginRequestUnionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#OriginResponseUnion">OriginResponseUnion</a>

Methods:

- <code title="post /v1/accounts/origins">client.Accounts.Origins.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#AccountOriginService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#AccountOriginNewParams">AccountOriginNewParams</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#OriginResponseUnion">OriginResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/accounts/origins/{id}">client.Accounts.Origins.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#AccountOriginService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#AccountOriginUpdateParams">AccountOriginUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#OriginResponseUnion">OriginResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/accounts/origins">client.Accounts.Origins.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#AccountOriginService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#OriginResponseUnion">OriginResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/accounts/origins/{id}">client.Accounts.Origins.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#AccountOriginService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/accounts/origins/{id}">client.Accounts.Origins.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#AccountOriginService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#OriginResponseUnion">OriginResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## URLEndpoints

Params Types:

- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#URLEndpointRequestParam">URLEndpointRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#URLEndpointResponse">URLEndpointResponse</a>

Methods:

- <code title="post /v1/accounts/url-endpoints">client.Accounts.URLEndpoints.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#AccountURLEndpointService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#AccountURLEndpointNewParams">AccountURLEndpointNewParams</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#URLEndpointResponse">URLEndpointResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/accounts/url-endpoints/{id}">client.Accounts.URLEndpoints.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#AccountURLEndpointService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#AccountURLEndpointUpdateParams">AccountURLEndpointUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#URLEndpointResponse">URLEndpointResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/accounts/url-endpoints">client.Accounts.URLEndpoints.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#AccountURLEndpointService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#URLEndpointResponse">URLEndpointResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/accounts/url-endpoints/{id}">client.Accounts.URLEndpoints.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#AccountURLEndpointService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/accounts/url-endpoints/{id}">client.Accounts.URLEndpoints.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#AccountURLEndpointService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#URLEndpointResponse">URLEndpointResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Beta

## V2

### Files

Response Types:

- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#BetaV2FileUploadResponse">BetaV2FileUploadResponse</a>

Methods:

- <code title="post /api/v2/files/upload">client.Beta.V2.Files.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#BetaV2FileService.Upload">Upload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#BetaV2FileUploadParams">BetaV2FileUploadParams</a>) (<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#BetaV2FileUploadResponse">BetaV2FileUploadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Webhooks

Response Types:

- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#BaseWebhookEvent">BaseWebhookEvent</a>
- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#UploadPostTransformErrorEvent">UploadPostTransformErrorEvent</a>
- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#UploadPostTransformSuccessEvent">UploadPostTransformSuccessEvent</a>
- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#UploadPreTransformErrorEvent">UploadPreTransformErrorEvent</a>
- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#UploadPreTransformSuccessEvent">UploadPreTransformSuccessEvent</a>
- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#VideoTransformationAcceptedEvent">VideoTransformationAcceptedEvent</a>
- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#VideoTransformationErrorEvent">VideoTransformationErrorEvent</a>
- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#VideoTransformationReadyEvent">VideoTransformationReadyEvent</a>
- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#UnsafeUnwrapWebhookEventUnion">UnsafeUnwrapWebhookEventUnion</a>
- <a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go">imagekit</a>.<a href="https://pkg.go.dev/github.com/imagekit-developer/imagekit-go#UnwrapWebhookEventUnion">UnwrapWebhookEventUnion</a>
