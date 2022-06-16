[<img width="250" alt="ImageKit.io" src="https://raw.githubusercontent.com/imagekit-developer/imagekit-javascript/master/assets/imagekit-light-logo.svg"/>](https://imagekit.io)

# ImageKit-go
ImageKit.io Go SDK

ImageKit Go SDK allows you to use [image resizing](https://docs.imagekit.io/features/image-transformations), [optimization](https://docs.imagekit.io/features/image-optimization), [file uploading](https://docs.imagekit.io/api-reference/upload-file-api) and other [ImageKit APIs](https://docs.imagekit.io/api-reference/api-introduction) from applications written in the Go language.

All features except url generation and utility functions return response with ```ResponseMetaData``` which holds raw response Header, StatusCode and Body. The Response object also contains ```Data``` attribtue except when underlying api call is not supposed to return any data(such as DeleteAsset).

Table of contents -
 * [Installation](#installation)
 * [Initialization](#initialization)
 * [URL Generation](#url-generation)
 * [File Upload](#file-upload)
 * [File Management](#file-management)
 * [Metadata API](#metadata-api)
 * [Utility Functions](#utility-functions)
 * [Rate Limits](#rate-limits)
 * [Support](#support)
 * [Links](#links)


## Version Support

| SDK Version | Go > 1.18 |
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
)

// Using environment variables IMAGEKIT_PRIVATE_KEY, IMAGEKIT_PUBLIC_KEY and IMAGEKIT_URL_ENDPOINT
imgkit, err := ImageKit.New()

// Using keys in argument
imgkit, err := ImageKit.NewFromParams(imagekit.NewParams{
    PrivateKey: privateKey,
    PublicKey: publicKey,
    UrlEndpoint: urlEndpoint
})
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
url, err := imgkit.Url(url.UrlParams{
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

The upload() method accepts file and UploadParams. File param can be local filepath, base64 encode image, url or io.Reader. If File is not local filepath then ```UploadParam``` requires the ```fileName``` parameter. This method returns ```UploadResponse``` object and `err` if any. You can pass other parameters supported by the ImageKit upload API using the same parameter name as specified in the upload API documentation. For example, to specify tags for a file at the time of upload, use the tags parameter as specified in the [documentation here](https://docs.imagekit.io/api-reference/upload-file-api/server-side-file-upload).

```
import "github.com/dhaval070/imagekit-go/uploader"

filePath := "/my/local/file.jpg"
resp, err := imgkit.Upload.Upload(ctx, file, uploader.UploadParams{})

filePath := "http://mydomain/image.jpg"
resp, err := imgkit.Upload.Upload(ctx, file, uploader.UploadParams{
    FileName: "myimage.jpg"
})

```

## File-Management

The SDK provides a simple interface for all the [media APIs mentioned here](https://docs.imagekit.io/api-reference/media-api) to manage your files. 

### 1. List & Search Files
List assets in media library, optionally filter and sort using ```AssetParams```.
```
import (
    "github.com/dhaval070/imagekit-go"
    "github.com/dhaval070/imagekit-go/api/media"
)

resp, err := imgkit.Media.Assets(ctx, media.AssetsParam{
    Skip: 10,
    Limit: 500,
    SearchQuery: "createdAt >= \"7d\" AND size > \"2mb\"",
})
```

### 2. Get File Details
Accepts the file ID and fetches the details as per the [API documentation here](https://docs.imagekit.io/api-reference/media-api/get-file-details).
```
resp, err := imgkit.Media.AssetById(ctx, media.AssetByIdParams{
    FileId: fileId
})
```

### 3. Get File Version Details
Get all the details and attributes of any version of a file as per the [API documentation here](https://docs.imagekit.io/api-reference/media-api/get-file-version-details).
```
resp, err := imgkit.Media.AssetVersions(ctx, media.AssetVersionsParam{
    FileId: fileId,
    VersionId: "version-id",
})

```

### 4. Get File Versions
Get all the file version details and attributes of a file as per the [API documentation here](https://docs.imagekit.io/api-reference/media-api/get-file-versions).
```
resp, err := imgkit.Media.AssetVersions(ctx, media.AssetVersionsParam{
    FileId: fileId,
})
```

### 5. Update File Details
Update parameters associated with the file as per the [API documentation here](https://docs.imagekit.io/api-reference/media-api/update-file-details).
```
resp, err := imgkit.Media.UpdateAsset(ctx, fileId, media.UpdateAssetParam{
    Tags: []string{"tag1", "tag2"},
    RemoveAITags: []string{"car", "suv"},
})
```

### 6. Add Tags (bulk)
Adds given tags to multiple files. Accepts slices of tags and file ids. Returns slice of file ids. [API documentation here](https://docs.imagekit.io/api-reference/media-api/add-tags-bulk)
```
resp, err := imgkit.Media.AddTags(ctx, media.TagsParam{
    FileIds: []string{"one", "two"},
    Tags: []string{"tag1", "tag2"},
})
```

### 7. Remove Tags (bulk)
Removes tags from multiple assets. Returns slice of file ids updated. [API documentation here](https://docs.imagekit.io/api-reference/media-api/remove-tags-bulk)
```
resp, err := imgkit.Media.RemoveTags(ctx, media.TagsParam{
    FileIds: []string{"one", "two"},
    Tags: []string{"tag1", "tag2"},
})
```
### 8. Remove AITags (bulk)
Remove AITags in bulk API. Returns slice of file ids. [API documentation here](https://docs.imagekit.io/api-reference/media-api/remove-aitags-bulk)
```
resp, err := imgkit.Media.RemoveAITags(ctx, media.AITagsParam{
    FileIds: []string{"one", "two"},
    AITags: []string{"tag1", "tag2"},
})
```

### 9. Delete File
Delete a file by fileId. [API documentation here](https://docs.imagekit.io/api-reference/media-api/delete-file).
```
resp, err := imgkit.Media.DeleteAsset(ctx, "32435343334")
```

### 10. Delete File Version
Deletes given version of the file. [API documentation here](https://docs.imagekit.io/api-reference/media-api/delete-file-version)
```
resp, err := imgkit.Media.DeleteAssetVersion(ctx, "32435343334", "version-1")
```

### 11. Delete Files (bulk)
Deletes multiple files. [API documentation here](https://docs.imagekit.io/api-reference/media-api/delete-files-bulk).
```
resp, err := imgkit.Media.DeleteBulkAssets(ctx, media.FileIdsParam{
    FileIds: []string{"324353234", "354332432"},
)
```

### 12. Copy File
This will copy a file from one location to another as per [API documentation here](https://docs.imagekit.io/api-reference/media-api/copy-file).
Accepts the source file's path and destination folder path.
```
resp, err := imgkit.Media.CopyAsset(ctx, media.CopyAssetParam{
    SourcePath: "/source/a.jpg",
    DestinationPath: "/target/",
    IncludeVersions: true,
})
```

### 13. Move File
This will move a file from one location to another as per [API documentation here](https://docs.imagekit.io/api-reference/media-api/move-file).
Accepts the source file's path and destination folder path.
```
resp, err := imgkit.Media.MoveAsset(ctx, media.MoveAssetParam{
    SourcePath: "/source/a.jpg",
    DestinationPath: "/target/",
})
```

### 14. Rename File
Renames a file as per [API documentation here](https://docs.imagekit.io/api-reference/media-api/rename-file).
Accepts file path, new name and purge cache option.
```
resp, err := imgkit.Media.RenameAsset(ctx, media.RenameAssetParam{
    FilePath: "/path/to/file.jpg",
    NewFileName: "newname.jpg",
    PurgeCache: true,
})

```

### 15. Restore File Version
Restore file version to a different version of a file as per [API documentation here](https://docs.imagekit.io/api-reference/media-api/restore-file-version).
Accepts string type file id and version id.
```
file, err := imgkit.Media.RestoreVersion(ctx, media.AssetVersionsParam{
    FileId: "324325334",
    VersionId: "243434",
})
```

### 16. Create Folder
Creates a new folder as per [API documentation here](https://docs.imagekit.io/api-reference/media-api/create-folder). ```err``` is not nil when response is not 201.
Accepts string type folder name and parent path.
```
resp, err := imgkit.Media.CreateFolder(ctx, media.CreateFolderParam{
   FolderName: "nature",
   ParentFolderPath: "/some/pics"
}
```

### 17. Delete Folder
Deletes the specified folder and all nested files, their versions & folders. This action cannot be undone. Accepts string type folder name to delete. [API documentation here](https://docs.imagekit.io/api-reference/media-api/delete-folder).

```
resp, err := imgkit.Media.DeleteFolder(ctx, media.DeleteFolderParam{
    FolderPath: "/some/pics/nature",
})
```

### 18. Copy Folder
Copies given folder to new location with or without versions info as per [API documentation here](https://docs.imagekit.io/api-reference/media-api/copy-folder).
```
err := imgkit.Media.CopyFolder(ctx, media.CopyFolderParam{
    SourceFolderPath: "source/path",
    DestinationPath: "destination/",
    IncludeVersions: true
})
```

### 19. Move Folder
Moves given folder path to new location as per [API documentation here](https://docs.imagekit.io/api-reference/media-api/move-folder).
```
err := imgkit.Media.MoveFolder(ctx, media.MoveFolderParam{
    SourceFolderPath: "source/path",
    DestinationPath: "destination/path",
})
```

### 20. Bulk Job Status
Get status of a bulk job operation by job id.  Accepts string type job id. [API documentation here](https://docs.imagekit.io/api-reference/media-api/copy-move-folder-status).

```
resp, err := imgkit.BulkJobStatus(ctx, "323235")
```

### 21. Purge Cache
This will purge given url's CDN and ImageKit.io's internal cache as per [API documentation here](https://docs.imagekit.io/api-reference/media-api/purge-cache).
```
resp, err := imgkit.Media.PurgeCache(ctx, media.PurgeCacheParam{
    Url: "https://ik.imageki.io/your_imagekit_id/rest-of-the-file-path.jpg"
})
```

### 22. Purge Cache Status
Get the status of the submitted purge request. Accepts purge request id. [API documentation here](https://docs.imagekit.io/api-reference/media-api/purge-cache-status).

```
resp, err := imgkit.Media.PurgeCacheStatus(ctx, "35325532")
```

## Metadata API
### 1. Get File Metadata for uploaded media files
Accepts the file ID or URL and fetches the metadata as per the [API documentation here](https://docs.imagekit.io/api-reference/metadata-api/get-image-metadata-for-uploaded-media-files).
```
fileId := "32432523432433335"
resp, err := imgkit.Metadata.FromAsset(ctx, fileId)
```

### 2. Get File Metadata from remote url
Get image EXIF, pHash, and other metadata from ImageKit.io powered remote URL using this API as per the [API documentation here](https://docs.imagekit.io/api-reference/metadata-api/get-image-metadata-from-remote-url).
```
resp, err := imgkit.Metadata.FromUrl(ctx, "http://domian/a.jpg")
```

## Custom Metadata fields API
Create, Update, Read and Delete custom metadata rules as per the [API documentation here](https://docs.imagekit.io/api-reference/custom-metadata-fields-api).

### 1. Create custom metadata field
```
import "github.com/dhaval070/imagekit-go/api/media/metadata"

resp, err := imgkit.Metadata.CreateCustomField(ctx, metadata.CreateFieldParam{
    Name: "weight",
    Label: "Weight",
    Schema: metadata.Schema{
        Type: "Number",
        MinValue: "1",
        MaxValue: "1000",
    }
})
```

### 2. List custom metadata fields
Accepts context and boolean flag(true|false) to get deleted fields. 
```
resp, err := imgkit.Metadata.CustomFields(ctx, true)

```

### 3. Update custom metadata field
```
resp, err := imgkit.Metadata.UpdateCustomField(ctx, UpdateCustomFieldParam{
    FieldId: "3w3255433",
    Label: "Cost",
})
```

### 4. Delete custom metadata field
Accepts context and fieldId to delete the custom metadata field.
```
resp, err := imgkit.Metadata.DeleteCustomField(ctx, "3325343434")
```
    
## Utility Functions
### 1. SignToken
This method generates a signature for given token and timestamp using the configured private key. It is useful for client side file upload to authenticate requests. ```Token``` is a random string. ```Expires``` is a unix timestamp by which token should expire.  ```Token``` and ```Expires``` both are optional parameters. Token defaults to auto generated uuid string. Expires defaults to a current time + 30 minutes value.
```
// Using auto generated token and expiration
resp := ik.SignToken(imagekit.SignTokenParam{})

// Using specific token and expiration
resp := ik.SignToken(imagekit.SignTokenParam{
    Token: "31c468de-520a-4dc1-8868-de1e0fb93a7b",
    Expires: 1655379249
})

```

## Rate Limits
Except for upload API, all ImageKit APIs are rate limited to avoid excessive request rates. 

Whenever backend api returns 429 status code, error of type ```ErrTooManyRequests``` is returned, which can be tested with ```errors.Is```. The rate limit detail can be retrieved from response metadata header. Please sleep/pause for the number of milliseconds specified by the value of ```resp.ResponseMetaData.Header["X-RateLimit-Reset"]``` property before making additional requests to that endpoint.

```
import (
    "errors"

    "github.com/dhaval070/imagekit-go"
    "github.com/dhaval070/imagekit-go/metadata"
    "github.com/dhaval070/imagekit-go/api"
)
imgkit, err := ImageKit.New()

resp, err := ik.Metadata.CustomFields(ctx, true)
if errors.Is(err, api.ErrTooManyRequests) {
    log.Println("rate limit exceeded", resp.ResponseMetaData.Header["X-RateLimit-Limit"])
}

```
