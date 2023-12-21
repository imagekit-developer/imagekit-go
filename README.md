[<img width="250" alt="ImageKit.io" src="https://raw.githubusercontent.com/imagekit-developer/imagekit-javascript/master/assets/imagekit-light-logo.svg"/>](https://imagekit.io)

# ImageKit.io Go SDK
[![CI](https://github.com/imagekit-developer/imagekit-go/workflows/CI/badge.svg)](https://github.com/imagekit-developer/imagekit-go/)
[![codecov](https://codecov.io/gh/imagekit-developer/imagekit-go/branch/dev/graph/badge.svg)](https://codecov.io/gh/imagekit-developer/imagekit-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Twitter Follow](https://img.shields.io/twitter/follow/imagekitio?label=Follow&style=social)](https://twitter.com/ImagekitIo)

Go SDK for [ImageKit](https://imagekit.io/) implements all the backend API, URL-generation, and other utility functions.

ImageKit is complete media storage, optimization, and transformation solution that comes with an [image and video CDN](https://imagekit.io). It can be integrated with your existing infrastructure - storage like AWS S3, web servers, your CDN, and custom domain names, allowing you to deliver optimized images in minutes with minimal code changes.

All methods except url generation and utility functions return a response with `ResponseMetaData`, which holds the raw response header, HTTP status code, and raw response body. The Response object also contains the `Data` attribute except when the underlying API call is not supposed to return any data such as delete file API.

Table of contents -

 * [Installation](#installation)
 * [Initialization](#initialization)
 * [Response Format](#response-format)
 * [Error Handling](#error-handling)
 * [URL Generation](#url-generation)
 * [File Upload](#file-upload)
 * [File Management](#file-management)
 * [Metadata API](#metadata-api)
 * [Utility Functions](#utility-functions)
 * [Rate Limits](#rate-limits)
 * [Support](#support)
 * [Links](#links)


## Version Support

| SDK Version | Go        |
|-------------|-----------|
| 1.x         | >=1.18     |


## Installation

```bash
go get github.com/imagekit-developer/imagekit-go
```

## Initialization

```go
import (
    "github.com/imagekit-developer/imagekit-go"
)

// Using environment variables IMAGEKIT_PRIVATE_KEY, IMAGEKIT_PUBLIC_KEY and IMAGEKIT_ENDPOINT_URL
ik, err := ImageKit.New()

// Using keys in argument
ik, err := ImageKit.NewFromParams(imagekit.NewParams{
    PrivateKey: privateKey,
    PublicKey: publicKey,
    UrlEndpoint: urlEndpoint
})
```

## Response Format
Results returned by functions that call backend API(such as media management, metadata, cache APIs) embeds raw response in `ResponseMetaData`, which can be used to get the response HTTP `StatusCode`, `Header`, and `Body`. The JSON response body is parsed to the appropriate SDK type and assigned to the `Data`  attribute.

```
resp, err := ik.Metadata.FromFile(ctx, fileId)
log.Println(resp.ResponseMetaData.Header, resp.Data.Url)

```

Functions that do not get any response body from API do not include the `Data` attribute in the response. In such cases, only `ResponseMetaData` is available.

## Error Handling
ImageKit API returns a non-2xx status code upon error.
SDK defines the following errors in the API package based on the status code returned:

```
imagekit-go/api:
400: ErrBadRequest
401: ErrUnauthorized
403: ErrForbidden
404: ErrNotFound
429: ErrTooManyRequests
500, 502, 503, 504: ErrServer
default: "Undefined Error"
```

`err` can be tested using `errors.Is`

```
if errors.Is(err, api.ErrForbidden) {
    log.Println(err.Message)
}
```

[See full documentation](https://docs.imagekit.io/api-reference/api-introduction) for further detail.

## URL-generation

### 1. Using image path and image hostname or endpoint
This method allows you to create an URL to access a file using the relative file path and the ImageKit URL endpoint (`urlEndpoint`). The file can be an image, video, or any other static file supported by ImageKit.

```
import (
    ikurl "github.com/imagekit-developer/imagekit-go/url"
)

url, err := ik.Url(ikurl.UrlParam{
    Path: "/default-image.jpg",
    UrlEndpoint: "https://ik.imagekit.io/your_imagekit_id/endpoint/",
    Transformations: []map[string]any{
        {
            "width":  400,
            "height": 300,
            "rotation": 90,
        },
    },
})
```
This results in a URL like:
```
https://ik.imagekit.io/your_imagekit_id/endpoint/tr:h-300,w-400:rt-90/default-image.jpg
```

### 2. Using full image URL
This method allows you to add transformation parameters to an absolute URL. For example, if you have configured a custom CNAME and have absolute asset URLs in your database or CMS, you will often need this.

```
import (
    ikurl "github.com/imagekit-developer/imagekit-go/url"
)

url, err := ik.Url(ikurl.UrlParam{
    Src: "https://ik.imagekit.io/your_imagekit_id/endpoint/default-image.jpg",
    Transformations: []map[string]any{
        {
            "width":  400,
            "height": 300,
            "rotation": 90,
        },
    },
})
```

This results in a URL like:

```
https://ik.imagekit.io/your_imagekit_id/endpoint/default-image.jpg?tr=h-300,w-400:rt-90
```


`UrlParam` has the following options:

| Option           | Description                    |
| :----------------| :----------------------------- |
| Path             | Conditional. This is the path at which the image exists. For example, `/path/to/image.jpg`. Either the `Path` or `Src` parameter needs to be specified for URL generation. |
| Src              | Conditional. This is the complete URL of an image already mapped to ImageKit. For example, `https://ik.imagekit.io/your_imagekit_id/endpoint/path/to/image.jpg`. Either the `Path` or `Src` parameter needs to be specified for URL generation. |
| UrlEndpoint      | Optional. The base URL to be appended before the path of the image. If not specified, the URL Endpoint specified at the time of SDK initialization is used. For example, https://ik.imagekit.io/your_imagekit_id/endpoint/ |
| Transformations   | Optional. An array of objects specifying the transformation to be applied in the URL. Different steps of a [chained transformation](https://docs.imagekit.io/features/image-transformations/chained-transformations) can be specified as different objects of the array. The complete list of supported transformations in the SDK and some examples of using them are given later. 
| TransformationPosition | Optional. The default value is `Path`, which places the transformation string as a path parameter in the URL. It can also be specified as `query`, which adds the transformation string as the URL's query parameter `tr`. If you use the `Src` parameter to create the URL, then the transformation string is always added as a query parameter. |
| NamedTransformation | Optional. Specifies the name of a pre-defined transformation. |
| QueryParameters  | Optional. These are the other query parameters that you want to add to the final URL. These can be any query parameters and not necessarily related to ImageKit. Especially useful if you want to add some versioning parameters to your URLs. |
| Signed           | Optional. Boolean. Default is `false`. If set to `true`, the SDK generates a signed image URL adding the image signature to the image URL. If you create a URL using the `Src` parameter instead of `Path`, then do correct `UrlEndpoint` for this to work. Otherwise returned URL will have the wrong signature |
| ExpireSeconds    | Optional. Integer. Meant to be used along with the `Signed` parameter to specify the time in seconds from now when the URL should expire. If specified, the URL contains the expiry timestamp in the URL, and the image signature is modified accordingly. |

#### Examples of generating URLs
**1. Chained Transformations as a query parameter**
```go

params := ikurl.UrlParam{
    Path:        "default-image.jpg",
    UrlEndpoint: "https://ik.imagekit.io/demo-id/",
    Transformations: []map[string]any{
        {
            "height": 300,
            "width":  400,
        },
        {
            "rotation": 90,
        },
    },
    TransformationPosition: ikurl.QUERY,
},

url, err := ik.Url(params)
```

**2. Sharpening and contrast transform and a progressive JPG image**

Some transform like [Sharpening](https://docs.imagekit.io/features/image-transformations/image-enhancement-and-color-manipulation) can be added to the URL with or without any other value. To use such transforms without specifying a value, specify the value as "-" in the transformation object. Otherwise, specify the value that you want to be added to this transformation.
```go
params := ikurl.UrlParam{
    Path:        "default-image.jpg",
    UrlEndpoint: "https://ik.imagekit.io/demo-id/",
    Transformations: []map[string]any{
        {
            "effectSharpen": "-",
        },
    },
}
```

**3. Adding overlays**

ImageKit.io enables you to apply overlays to [images](https://docs.imagekit.io/features/image-transformations/overlay-using-layers) and [videos](https://docs.imagekit.io/features/video-transformation/overlay) using the raw parameter with the concept of [layers](https://docs.imagekit.io/features/image-transformations/overlay-using-layers#layers). The raw parameter facilitates incorporating transformations directly in the URL. A layer is a distinct type of transformation that allows you to define an asset to serve as an overlay, along with its positioning and additional transformations.

**Text as overlays**

You can add any text string over a base video or image using a text layer (l-text).

For example:

```go
params := ikurl.UrlParam{
    Path:        "default-image.jpg",
    UrlEndpoint: "https://ik.imagekit.io/demo-id/",
    Transformations: []map[string]any{
        {
            "height": 300,
            "width":  400,
            "raw": "l-text,i-Imagekit,fs-50,l-end"
        },
    },
}
```
**Sample Result URL**
```
https://ik.imagekit.io/demo-id/default-image.jpg?tr=h-300,w-400,l-text,i-Imagekit,fs-50,l-end
```

**Image as overlays**

You can add an image over a base video or image using an image layer (l-image).

For example:

```go
params := ikurl.UrlParam{
    Path:        "default-image.jpg",
    UrlEndpoint: "https://ik.imagekit.io/demo-id/",
    Transformations: []map[string]any{
        {
            "height": 300,
            "width":  400,
            "raw": "l-image,i-default-image.jpg,w-100,b-10_CDDC39,l-end"
        },
    },
}
```
**Sample Result URL**
```
https://ik.imagekit.io/demo-id/default-image.jpg?tr=h-300,w-400,l-image,i-default-image.jpg,w-100,b-10_CDDC39,l-end
```

**Solid color blocks as overlays**

You can add solid color blocks over a base video or image using an image layer (l-image).

For example:

```go
params := ikurl.UrlParam{
    Path:        "img/sample-video.mp4",
    UrlEndpoint: "https://ik.imagekit.io/demo-id/",
    Transformations: []map[string]any{
        {
            "height": 300,
            "width":  400,
            "raw": "l-image,i-ik_canvas,bg-FF0000,w-300,h-100,l-end"
        },
    },
}
```
**Sample Result URL**
```
https://ik.imagekit.io/demo-id/img/sample-video.mp4?tr=h-300,w-400,l-image,i-ik_canvas,bg-FF0000,w-300,h-100,l-end
```
#### List of supported transformations

See the complete list of transformations supported in ImageKit [here](https://docs.imagekit.io/features/image-transformations). The SDK gives a name to each transformation parameter e.g. `height` for `h` and `width` for `w` parameter. It makes your code more readable. If the property does not match any of the following supported options, it is added as it is.

If you want to generate transformations in your application and add them to the URL as it is, use the raw parameter.

| Supported Transformation Name | Translates to parameter |
|-------------------------------|-------------------------|
|height                    |h|
|width                     |w|
|aspectRatio               |ar|
|quality                   |q|
|crop                      |c|
|cropMode                  |cm|
|x                         |x|
|y                         |y|
|xc                        |xc|
|yc                        |yc|
|focus                     |fo|
|format                    |f|
|radius                    |r|
|background                |bg|
|border                    |b|
|rotation                  |rt|
|blur                      |bl|
|named                     |n|
|progressive               |pr|
|lossless                  |lo|
|trim                      |t|
|metadata                  |md|
|colorProfile              |cp|
|defaultImage              |di|
|dpr                       |dpr|
|effectSharpen             |e-sharpen|
|effectUSM                 |e-usm|
|effectContrast            |e-contrast|
|effectGray                |e-grayscale|
|original                  |orig|
|raw                       | `replaced by the parameter value`|


## File-Upload

The SDK uploader package provides a simple interface using the `.upload()` method to upload files to the ImageKit Media Library. It accepts all the parameters supported by the [ImageKit Upload API](https://docs.imagekit.io/api-reference/upload-file-api/server-side-file-upload).

The upload() method accept file and UploadParam. File param can be base64 encoded image, absolute HTTP URL, or io.Reader. This method returns the `UploadResponse` object and `err` if any. In addition, you can pass other parameters supported by the ImageKit upload API using the same parameter name as specified in the upload API documentation. For example, to set tags for a file at the upload time, use the tags parameter as defined in the [documentation here](https://docs.imagekit.io/api-reference/upload-file-api/server-side-file-upload).

```
import "github.com/imagekit-developer/imagekit-go/uploader"

const base64Image = "data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7"
resp, err := ik.Uploader.Upload(ctx, base64Image, uploader.UploadParam{
    FileName: "myimage.jpg"
})

```

## File-Management

The SDK provides a simple interface for all the [media APIs mentioned here](https://docs.imagekit.io/api-reference/media-api) to manage your files. 

### 1. List & Search Files
List files in the media library, optionally filter and sort using `FileParams`.

```
import (
    "github.com/imagekit-developer/imagekit-go"
    "github.com/imagekit-developer/imagekit-go/api/media"
)

resp, err := ik.Media.Files(ctx, media.FilesParam{
    Skip: 10,
    Limit: 500,
    SearchQuery: "createdAt >= \"7d\" AND size > \"2mb\"",
})
```

### 2. Get File Details
Accepts the file ID and fetches the details as per the [API documentation here](https://docs.imagekit.io/api-reference/media-api/get-file-details).

```
resp, err := ik.Media.FileById(ctx, "file_id")
```

### 3. Get File Version Details
Get all the details and attributes of any version of a file as per the [API documentation here](https://docs.imagekit.io/api-reference/media-api/get-file-version-details).

```
resp, err := ik.Media.FileVersions(ctx, media.FileVersionsParam{
    FileId: "file_id",
    VersionId: "version_id",
})

```

### 4. Get File Versions
Get all the file version details and attributes of a file as per the [API documentation here](https://docs.imagekit.io/api-reference/media-api/get-file-versions).

```
resp, err := ik.Media.FileVersions(ctx, media.FileVersionsParam{
    FileId: "file_id",
})
```

### 5. Update File Details
Update parameters associated with the file as per the [API documentation here](https://docs.imagekit.io/api-reference/media-api/update-file-details).

```
resp, err := ik.Media.UpdateFile(ctx, fileId, media.UpdateFileParam{
    Tags: []string{"tag_1", "tag_2"},
    RemoveAITags: []string{"car", "suv"},
})
```

### 6. Add Tags (bulk)
Set tags to multiple files. Accepts slices of tags and file Ids. Returns slice of file ids. [API documentation here](https://docs.imagekit.io/api-reference/media-api/add-tags-bulk).

```
resp, err := ik.Media.AddTags(ctx, media.TagsParam{
    FileIds: []string{"file_id_1", "file_id_2"},
    Tags: []string{"tag_1", "tag_2"},
})
```

### 7. Remove Tags (bulk)
Removes tags from multiple files. Returns slice of file IDs updated. [API documentation here](https://docs.imagekit.io/api-reference/media-api/remove-tags-bulk).

```
resp, err := ik.Media.RemoveTags(ctx, media.TagsParam{
    FileIds: []string{"file_id_1", "file_id_2"},
    Tags: []string{"tag_1", "tag_2"},
})
```
### 8. Remove AITags (bulk)
Remove AITags in bulk API. Returns slice of file ids. [API documentation here](https://docs.imagekit.io/api-reference/media-api/remove-aitags-bulk).

```
resp, err := ik.Media.RemoveAITags(ctx, media.AITagsParam{
    FileIds: []string{"file_id_1", "file_id_2"},
    AITags: []string{"tag_1", "tag_2"},
})
```

### 9. Delete File
Delete a file by fileId. [API documentation here](https://docs.imagekit.io/api-reference/media-api/delete-file).
```
resp, err := ik.Media.DeleteFile(ctx, "file_id")
```

### 10. Delete File Version
Deletes the given version of the file. [API documentation here](https://docs.imagekit.io/api-reference/media-api/delete-file-version).
```
resp, err := ik.Media.DeleteFileVersion(ctx, "file_id", "version_1")
```

### 11. Delete Files (bulk)
Deletes multiple files. [API documentation here](https://docs.imagekit.io/api-reference/media-api/delete-files-bulk).

```
resp, err := ik.Media.DeleteBulkFiles(ctx, media.FileIdsParam{
    FileIds: []string{"file_id1", "file_id2"},
)
```

### 12. Copy File
This will copy a file from one location to another as per [API documentation here](https://docs.imagekit.io/api-reference/media-api/copy-file).

Accepts the source file's path and destination folder path.
```
resp, err := ik.Media.CopyFile(ctx, media.CopyFileParam{
    SourcePath: "/source/a.jpg",
    DestinationPath: "/target/",
    IncludeFileVersions: false,
})
```

### 13. Move File
This will move a file from one location to another as per [API documentation here](https://docs.imagekit.io/api-reference/media-api/move-file).

Accepts the source file's path and destination folder path.
```
resp, err := ik.Media.MoveFile(ctx, media.MoveFileParam{
    SourcePath: "/source/a.jpg",
    DestinationPath: "/target/",
})
```

### 14. Rename File
Renames a file as per [API documentation here](https://docs.imagekit.io/api-reference/media-api/rename-file).
Accepts file path, new name and purge cache option.

```
resp, err := ik.Media.RenameFile(ctx, media.RenameFileParam{
    FilePath: "/path/to/file.jpg",
    NewFileName: "newname.jpg",
    PurgeCache: true,
})

```

### 15. Restore File Version
Restore the file version as per [API documentation here](https://docs.imagekit.io/api-reference/media-api/restore-file-version).
Accepts string type file id and version id.
```
resp, err := ik.Media.RestoreVersion(ctx, media.FileVersionsParam{
    FileId: "file_id",
    VersionId: "version_id",
})
```

### 16. Create Folder
Creates a new folder as per [API documentation here](https://docs.imagekit.io/api-reference/media-api/create-folder). `err` is not nil when the response is not 201.

Accepts string type folder name and parent path.

```
resp, err := ik.Media.CreateFolder(ctx, media.CreateFolderParam{
   FolderName: "nature",
   ParentFolderPath: "/some/pics"
}
```

### 17. Delete Folder
Deletes the specified folder and all nested files, their versions & folders. This action cannot be undone. Accepts string type folder name to delete. [API documentation here](https://docs.imagekit.io/api-reference/media-api/delete-folder).


```
resp, err := ik.Media.DeleteFolder(ctx, media.DeleteFolderParam{
    FolderPath: "/some/pics/nature",
})
```

### 18. Copy Folder
Copies given folder to new location with or without versions info as per [API documentation here](https://docs.imagekit.io/api-reference/media-api/copy-folder).

```
resp, err := ik.Media.CopyFolder(ctx, media.CopyFolderParam{
    SourceFolderPath: "source/path",
    DestinationPath: "destination/",
    IncludeFileVersions: false
})
```

### 19. Move Folder
Moves given folder path to new location as per [API documentation here](https://docs.imagekit.io/api-reference/media-api/move-folder).

```
resp, err := ik.Media.MoveFolder(ctx, media.MoveFolderParam{
    SourceFolderPath: "source/path",
    DestinationPath: "destination/path",
})
```

### 20. Bulk Job Status
Get the status of a bulk job operation by job id. Accepts string type job id. [API documentation here](https://docs.imagekit.io/api-reference/media-api/copy-move-folder-status).

```
resp, err := ik.BulkJobStatus(ctx, "job_id")
```

### 21. Purge Cache
This will purge the CDN and ImageKit internal cache for a given URL. [API documentation here](https://docs.imagekit.io/api-reference/media-api/purge-cache).

```
resp, err := ik.Media.PurgeCache(ctx, media.PurgeCacheParam{
    Url: "https://ik.imageki.io/your_imagekit_id/rest-of-the-file-path.jpg"
})
```

### 22. Purge Cache Status
Get the status of the submitted purge request. Accepts purge request id. [API documentation here](https://docs.imagekit.io/api-reference/media-api/purge-cache-status).


```
resp, err := ik.Media.PurgeCacheStatus(ctx, "request_id")
```

## Metadata API
### 1. Get File Metadata for uploaded media files
Accepts the file ID or URL and fetches the metadata as per the [API documentation here](https://docs.imagekit.io/api-reference/metadata-api/get-image-metadata-for-uploaded-media-files).

```
resp, err := ik.Metadata.FromFile(ctx, "file_id")
```

### 2. Get File Metadata from remote url
Get image EXIF, pHash, and other metadata from ImageKit.io powered remote URL using this API as per the [API documentation here](https://docs.imagekit.io/api-reference/metadata-api/get-image-metadata-from-remote-url).

```
resp, err := ik.Metadata.FromUrl(ctx, "http://domian/a.jpg")
```

## Custom Metadata fields API
Create, Update, Read and Delete custom metadata rules as per the [API documentation here](https://docs.imagekit.io/api-reference/custom-metadata-fields-api).

### 1. Create custom metadata field
```
import "github.com/imagekit-developer/imagekit-go/api/media/metadata"

resp, err := ik.Metadata.CreateCustomField(ctx, metadata.CreateFieldParam{
    Name: "weight",
    Label: "Weight",
    Schema: metadata.Schema{
        Type: "Number",
        MinValue: 1,
        MaxValue: 1000,
    }
})
```

### 2. List custom metadata fields
Accepts context and boolean flag(true|false) to get deleted fields.

```
resp, err := ik.Metadata.CustomFields(ctx, true)

```

### 3. Update custom metadata field

```
resp, err := ik.Metadata.UpdateCustomField(ctx, "field_id", UpdateCustomFieldParam{
    Label: "Cost",
})
```

### 4. Delete custom metadata field
Accepts context and fieldId to delete the custom metadata field.
```
resp, err := ik.Metadata.DeleteCustomField(ctx, "field_id")
```
    
## Utility Functions

We have included the following commonly used utility function in this package.

### 1. Authentication parameter generation
This method generates a signature for a given token and timestamp using the configured private key. It is useful for client-side file upload to authenticate requests. `Token` is a random string. `Expires` is a unix timestamp by which token should expire. `Token` and `Expires` are both optional parameters. `Token` defaults to an auto-generated UUID string. `Expires` defaults to a current time + 30 minutes value.

```
// Using auto-generated token and expiration
resp := ik.SignToken(imagekit.SignTokenParam{})

// Using specific token and expiration
resp := ik.SignToken(imagekit.SignTokenParam{
    Token: "token-string",
    Expires: 1655379249,
})

```

## Rate Limits
Except for upload API, all ImageKit APIs are rate limited to avoid excessive request rates. 

Whenever backend API returns 429 status code, error of type `ErrTooManyRequests` is returned, which can be tested with `errors.Is`. The rate limit detail can be retrieved from the response metadata header. Please sleep/pause for the number of milliseconds specified by the value of `resp.ResponseMetaData.Header["X-RateLimit-Reset"]` property before making additional requests to that endpoint.

```
import (
    "errors"

    "github.com/imagekit-developer/imagekit-go"
    "github.com/imagekit-developer/imagekit-go/metadata"
    "github.com/imagekit-developer/imagekit-go/api"
)
ik, err := ImageKit.New()

resp, err := ik.Metadata.CustomFields(ctx, true)
if errors.Is(err, api.ErrTooManyRequests) {
    log.Println("rate limit exceeded", resp.ResponseMetaData.Header["X-RateLimit-Limit"])
}

```

## Support

For any feedback or to report any issues or general implementation support, please reach out to [support@imagekit.io](mailto:support@imagekit.io)

## Links
* [Documentation](https://docs.imagekit.io)
* [Main website](https://imagekit.io)

## License
Released under the MIT license.
