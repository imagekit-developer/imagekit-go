[<img width="250" alt="ImageKit.io" src="https://raw.githubusercontent.com/imagekit-developer/imagekit-javascript/master/assets/imagekit-light-logo.svg"/>](https://imagekit.io)

# ImageKit-go
ImageKit.io Go SDK

ImageKit Go SDK allows you to use [image resizing](https://docs.imagekit.io/features/image-transformations), [optimization](https://docs.imagekit.io/features/image-optimization), [file uploading](https://docs.imagekit.io/api-reference/upload-file-api) and other [ImageKit APIs](https://docs.imagekit.io/api-reference/api-introduction) from applications written in the Go language.


Table of contents -
 * [Installation](#Installation)
 * [Initialization](#Initialization)
 * [URL Generation](#URL-generation)
 * [File Upload](#File-Upload)
 * [File Management](#File-Management)
 * [Utility Functions](#Utility-functions)
 * [Support](#Support)
 * [Links](#Links)


## Version Support

| SDK Version | Go > 1.13 |
|-------------|-----------|
| 1.x         | v         |


## Installation
```bash
go get github.com/dhaval070/imagekit-go
```

## Initialization

```go
import (
    "github.com/dhaval070/imagekit-go"
    "github.com/dhaval070/imagekit-go/uploader"
)

// Using environment variables IMAGEKIT_PRIVATE_KEY, IMAGEKIT_PUBLIC_KEY and IMAGEKIT_URL_ENDPOINT
imgkit, _ := ImageKit.New()

// Using keys in argument
imgkit, _ := ImageKit.NewFromParams(privateKey, publicKey, urlEndpoint)
```

## URL-generation
### 1. Using image path and image hostname or endpoint
This method allows you to create a URL using the ```path``` where the image exists and the URL endpoint (```urlEndpoint```) you want to use to access the image. You can refer to the documentation [here](https://docs.imagekit.io/integration/url-endpoints) to read more about URL endpoints in ImageKit and the section about [image origins](https://docs.imagekit.io/integration/configure-origin) to understand about paths with different kinds of origins.

```
url, err := imgkit.Url(url.UrlOptions{
    Path: "/default-image.jpg",
    UrlEndpoint: "https://ik.imagekit.io/your_imagekit_id/endpoint/",
    Transformation: "w-400,h-300:rt-90"
})
```
This results in a URL like:
```
https://ik.imagekit.io/your_imagekit_id/endpoint/tr:h-300,w-400:tr-90/default-image.jpg
```

### 2. This method allows you to add transformation parameters to an existing, complete URL that is already mapped to ImageKit using the ```src``` parameter. Use this method if you have the absolute image URL stored in your database.
```
url, err := imgkit.Url(url.UrlOptions{
    Src: "https://ik.imagekit.io/your_imagekit_id/endpoint/default-image.jpg",
    Transformation: "w-400,h-300:rt-90",
})
```
This results in a URL like:
```
https://ik.imagekit.io/your_imagekit_id/endpoint/default-image.jpg?tr=h-300,w=400:rt-90
```
[See full documentation](https://docs.imagekit.io/features/image-transformations) for transformation options.

## File-Upload
The SDK uploader package provides a simple interface using the ```.upload()``` method to upload files to the ImageKit Media Library. It accepts all the parameters supported by the [ImageKit Upload API](https://docs.imagekit.io/api-reference/upload-file-api/server-side-file-upload).

The upload() method requires at least the ```file``` and the ```fileName``` parameter to upload a file and returns ```file``` object and `err` if any.  You can pass other parameters supported by the ImageKit upload API using the same parameter name as specified in the upload API documentation. For example, to specify tags for a file at the time of upload, use the tags parameter as specified in the [documentation here](https://docs.imagekit.io/api-reference/upload-file-api/server-side-file-upload).

```
file, err := imgkit.Upload.Upload(ctx, uploader.UploadParams{
    File: "htts://example.com/myimage.jpg",
    FileName: "my image.jpg",
})

```

## File-Management
The SDK provides a simple interface for all the [media APIs mentioned here](https://docs.imagekit.io/api-reference/media-api) to manage your files. 

### 1. List & Search Files
```
import (
    "github.com/dhaval070/imagekit-go"
    "github.com/dhaval070/imagekit-go/assets"
)

files, err := imgkit.ListFiles(ctx, assets.ListParams{
    Skip: 10,
    Limit: 500,
    SearchQuery: "createdAt >= "7d" AND size > \"2mb\"",
})
```

### 2. Get File Details
Accepts the file ID and fetches the details as per the [API documentation here](https://docs.imagekit.io/api-reference/media-api/get-file-details).
```
file, err := imgkit.GetFileDetails(ctx, fileId)
```

### 3. Get File Version Details
Get all the details and attributes of any version of a file.
```
file, err := imgkit.GetFileVersionDetails(ctx, fileId)

```

### 4. Get File Versions
Get all the file version details and attributes of a file.
```
files, err := imgkit.GetFileAllVersionDetails(ctx, fileId)
```

### 5. Update File Details
Update parameters associated with the file as per the [API documentation here](https://docs.imagekit.io/api-reference/media-api/update-file-details).
```
resp, err := imgkit.UpdateAsset(ctx, assets.AssetParams{
    FileId: fileId,
    Tags: []string{"tag1", "tag2"},
})
```

### 6. Add Tags (bulk)
Adds given tags to multiple files. Accepts slices of tags and file ids. Returns slice of file ids. [API documentation here](https://docs.imagekit.io/api-reference/media-api/add-tags-bulk)
```
ids, err := imgkit.AddTags(ctx, assets.TagsParams{
    FileIds: []string{"one", "two"},
    Tags: []string{"tag1", "tag2"},
})
```

### 7. Remove Tags (bulk)
Remove tags in bulk API. Returns slice of file ids. [API documentation here](https://docs.imagekit.io/api-reference/media-api/remove-tags-bulk)
```
ids, err := imgkit.RemoveTags(ctx, assets.TagsParams{
    FileIds: []string{"one", "two"},
    Tags: []string{"tag1", "tag2"},
})
```
### 8. Remove AITags (bulk)
Remove AITags in bulk API. Returns slice of file ids. [API documentation here](https://docs.imagekit.io/api-reference/media-api/remove-aitags-bulk)
```
ids, err := imgkit.RemoveAITags(ctx, assets.AITagsParams{
    FileIds: []string{"one", "two"},
    AITags: []string{"tag1", "tag2"},
})
```

### 9. Delete File
Delete a file by fileId. [API documentation here](https://docs.imagekit.io/api-reference/media-api/delete-file).
```
ids, err := imgkit.DeleteAsset(ctx,
    fileId, // string
)
```

### 10. Delete File Version
Deletes given version of the file. [API documentation here](https://docs.imagekit.io/api-reference/media-api/delete-file-version)
```
err := imgkit.DeleteFileVersion(ctx,
    FileId, // string
    VersionId, // string
)
```

### 11. Delete Files (bulk)
Deletes multiple files. [API documentation here](https://docs.imagekit.io/api-reference/media-api/delete-files-bulk).
```
ids, err := imgkit.DeleteAsset(ctx,
    fileId1,
    fileId2,
)
```

### 12. Copy File
This will copy a file from one location to another as per [API documentation here](https://docs.imagekit.io/api-reference/media-api/copy-file).
Accepts the source file's path and destination folder path.
```
err := imgkit.CopyAsset(ctx, assets.CopyParams{
    SourcePath: "/source/a.jpg",
    DestinationPath: "/target/",
    IncludeVersions: true,
})
```

### 13. Move File
This will move a file from one location to another as per [API documentation here](https://docs.imagekit.io/api-reference/media-api/move-file).
Accepts the source file's path and destination folder path.
```
err := imgkit.MoveAsset(ctx, assets.MoveParams{
    SourcePath: "/source/a.jpg",
    DestinationPath: "/target/",
})
```

### 14. Rename File
Renames a file as per [API documentation here](https://docs.imagekit.io/api-reference/media-api/rename-file).
Accepts file path, new name and purge cache option.
```
resp, err := imgkit.RenameAsset(ctx, assets.RenameParams{
    FilePath: "/path/to/file.jpg",
    NewFileName: "newname.jpg",
    PurgeCache: true,
})

```
resp is of type map[string]string in case response is 200 or 207.

### 15. Restore File Version
Restore file version to a different version of a file as per [API documentation here](https://docs.imagekit.io/api-reference/media-api/restore-file-version).
Accepts string type file id and version id.
```
file, err := imgkit.RestoreAssetVersion(ctx, "xxxx", "33534")
```

### 16. Create Folder
Creates a new folder as per [API documentation here](https://docs.imagekit.io/api-reference/media-api/create-folder). ```err``` is not nil when response is not 201.
Accepts string type folder name and parent path.
```
err := imgkit.CreateFolder(ctx, "sample", "/some/parent/folder")
```

### 17. Delete Folder
Deletes the specified folder and all nested files, their versions & folders. This action cannot be undone. Accepts string type folder name to delete. [API documentation here](https://docs.imagekit.io/api-reference/media-api/delete-folder).

```
err := imgkit.DeleteFolder(ctx, "sample")
```

### 18. Copy Folder
Copies given folder to new location with or without versions info as per [API documentation here](https://docs.imagekit.io/api-reference/media-api/copy-folder).
```
err := imgkit.CopyFolder(ctx, "source/path", "destination/path", true)
```

### 19. Move Folder
Movies given folder path to new location as per [API documentation here](https://docs.imagekit.io/api-reference/media-api/move-folder).
```
err := imgkit.MoveFolder(ctx, "source/path", "destination/path")
```

### 20. Bulk Job Status
Get status of a bulk job operation by job id.  Accepts string type job id. [API documentation here](https://docs.imagekit.io/api-reference/media-api/copy-move-folder-status).

```
resp, err := imgkit.BulkJobStatus(ctx, "xxx")
```

### 21. Purge Cache
This will purge given url's CDN and ImageKit.io's internal cache as per [API documentation here](https://docs.imagekit.io/api-reference/media-api/purge-cache).
```
reqId, err := imgkit.PurgeCache(
    ctx,
    "https://ik.imageki.io/your_imagekit_id/rest-of-the-file-path.jpg"
)
```

### 22. Purge Cache Status
Get the status of the submitted purge request. Accepts purge request id. [API documentation here](https://docs.imagekit.io/api-reference/media-api/purge-cache-status).

```
status, err := imgkit.PurgeCacheStatus(ctx, "xxxx")
```

## MetaData API
### Get File Metadata
Accepts the file ID or URL and fetches the metadata as per the [API documentation here](https://docs.imagekit.io/api-reference/metadata-api/get-image-metadata-for-uploaded-media-files).
```
meta, err := imgkit.GetAssetMetadata(ctx, fileId)
```
