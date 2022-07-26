[<img width="250" alt="ImageKit.io" src="https://raw.githubusercontent.com/imagekit-developer/imagekit-javascript/master/assets/imagekit-light-logo.svg"/>](https://imagekit.io)

# ImageKit-go
ImageKit.io Go SDK

[![CI](https://github.com/imagekit-developer/imagekit-go/workflows/CI/badge.svg)](https://github.com/imagekit-developer/imagekit-go/)
[![codecov](https://codecov.io/gh/imagekit-developer/imagekit-go/branch/dev/graph/badge.svg)](https://codecov.io/gh/imagekit-developer/imagekit-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Twitter Follow](https://img.shields.io/twitter/follow/imagekitio?label=Follow&style=social)](https://twitter.com/ImagekitIo)

Go SDK for [ImageKit](https://imagekit.io/) implements all the backend API, URL-generation and other utility functions.

ImageKit is complete media storage, optimization, and transformation solution that comes with an [image and video CDN](https://imagekit.io). It can be integrated with your existing infrastructure - storage like AWS S3, web servers, your CDN, and custom domain names, allowing you to deliver optimized images in minutes with minimal code changes.

All methods except url generation and utility functions return response with `ResponseMetaData` which holds raw response header, HTTP status code and raw response body. The Response object also contains `Data` attribtue except when underlying API call is not supposed to return any data such as delete file API.

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

// Using environment variables IMAGEKIT_PRIVATE_KEY, IMAGEKIT_PUBLIC_KEY and IMAGEKIT_URL_ENDPOINT
ik, err := ImageKit.New()

// Using keys in argument
ik, err := ImageKit.NewFromParams(imagekit.NewParams{
    PrivateKey: privateKey,
    PublicKey: publicKey,
    UrlEndpoint: urlEndpoint
})
```

## Response Format
Results returned by functions which call backend api(such as media management, metadata, cache apis) embeds raw response in `ResponseMetaData`, which can be used to get the response HTTP `StatusCode`, `Header` and `Body`. The json resonse body is parsed to appropriate SDK type and assigned to `Data`  attribute.

```
resp, err := ik.Metadata.FromFile(ctx, fileId)
log.Println(resp.ResponseMetaData.Header, resp.Data.Url)

```

Functions which do not get any response body from API, does not include `Data` attribute in response. Such responses are of type `*api.Response` and only ResponseMetaData is available.

## Error Handling
ImageKit API returns non 2xx status code upon error.
SDK defines following errors in api package based on the status code returned:

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
This method allows you to create an URL to access a file using the relative file path and the ImageKit URL endpoint (`urlEndpoint`). The file can be an image, video or any other static file supported by ImageKit.

```
import (
	ikurl "github.com/imagekit-developer/imagekit-go/url"
)

url, err := ik.Url(ikurl.UrlParam{
    Path: "/default-image.jpg",
    UrlEndpoint: "https://ik.imagekit.io/your_imagekit_id/endpoint/",
    Transformations: []ikurl.Transformation{
        {
            Width:  400,
            Height: 300,
            Rotate: 90,
        },
    },
})
```
This results in a URL like:
```
https://ik.imagekit.io/your_imagekit_id/endpoint/tr:h-300,w-400:rt-90/default-image.jpg
```

### 2. Using full image URL**
This method allows you to add transformation parameters to an absolute URL. For example, if you have configured a custom CNAME and have absolute asset URLs in your database or CMS, you will often need this.

```
import (
	ikurl "github.com/imagekit-developer/imagekit-go/url"
)

url, err := ik.Url(ikurl.UrlParam{
    Src: "https://ik.imagekit.io/your_imagekit_id/endpoint/default-image.jpg",
    Transformations: []ikurl.Transformation{
        {
            Width:  400,
            Height: 300,
            Rotate: 90,
        },
    },
})
```

This results in a URL like:

```
https://ik.imagekit.io/your_imagekit_id/endpoint/default-image.jpg?tr=h-300,w=400:rt-90
```


`UrlParam` has following options:

| Option           | Description                    |
| :----------------| :----------------------------- |
| Path             | Conditional. This is the path at which the image exists. For example, `/path/to/image.jpg`. Either the `Path` or `Src` parameter needs to be specified for URL generation. |
| Src              | Conditional. This is the complete URL of an image already mapped to ImageKit. For example, `https://ik.imagekit.io/your_imagekit_id/endpoint/path/to/image.jpg`. Either the `Path` or `Src` parameter needs to be specified for URL generation. |
| UrlEndpoint      | Optional. The base URL to be appended before the path of the image. If not specified, the URL Endpoint specified at the time of SDK initialization is used. For example, https://ik.imagekit.io/your_imagekit_id/endpoint/ |
| Transformations   | Optional. An array of objects specifying the transformation to be applied in the URL. Different steps of a [chained transformation](https://docs.imagekit.io/features/image-transformations/chained-transformations) can be specified as different objects of the array. The complete list of supported transformations in the SDK and some examples of using them are given later. 
| TransformationPosition | Optional. The default value is `Path` that places the transformation string as a path parameter in the URL. It can also be specified as `query`, which adds the transformation string as the URL's query parameter `tr`. If you use the `Src` parameter to create the URL, then the transformation string is always added as a query parameter. |
| NamedTransformation | Optional. Specifies name of a pre defined transformation. |
| QueryParameters  | Optional. These are the other query parameters that you want to add to the final URL. These can be any query parameters and not necessarily related to ImageKit. Especially useful if you want to add some versioning parameter to your URLs. |
| Signed           | Optional. Boolean. Default is `false`. If set to `true`, the SDK generates a signed image URL adding the image signature to the image URL. If you create a URL using the `Src` parameter instead of `Path`, then do correct `UrlEndpoint` for this to work. Otherwise returned URL will have the wrong signature |
| ExpireSeconds    | Optional. Integer. Meant to be used along with the `Signed` parameter to specify the time in seconds from now when the URL should expire. If specified, the URL contains the expiry timestamp in the URL, and the image signature is modified accordingly. |

#### Examples of generating URLs
**1. Chained Transformations as a query parameter**
```go

params := ikurl.UrlParam{
    Path:        "default-image.jpg",
    UrlEndpoint: "https://ik.imagekit.io/demo-id/",
    Transformations: []ikurl.Transformation{
        {
            Height: 300,
            Width:  400,
        },
        {
            Rotate: 90,
        },
    },
    TransformationPosition: ikurl.QUERY,
},

url, err := ik.Url(params)
```

**2. Sharpening and contrast transforms and a progressive JPG image**

There are some transforms like [Sharpening](https://docs.imagekit.io/features/image-transformations/image-enhancement-and-color-manipulation) that can be added to the URL with or without any other value. To use such transforms without specifying a value, specify the value as "-" in the transformation object. Otherwise, specify the value that you want to be added to this transformation.
```go
params := ikurl.UrlParam{
    Path:        "default-image.jpg",
    UrlEndpoint: "https://ik.imagekit.io/demo-id/",
    Transformations: []ikurl.Transformation{
        {
            Sharpen: true,
        },
    },
}
```
#### List of supported transformations
url package defines transformation type:
```go
// Transformation is a struct representing options for transformation parameter to Url()
type Transformation struct {
	Width            float32     // w
	Height           float32     // h
	AspectRatio      AspectRatio // ar
	CropMode         CropMode
	Focus            Focus  // fo-custom|face|auto|left|right|top|bottom|top_left.....
	X                int    // x
	Y                int    // y
	Xc               int    // xc.  Focus with center coordinate Xc and Yc
	Yc               int    // yc
	Quality          int    // q
	Format           Format // f
	Blur             int    // bl
	Dpr              any    // dpr (auto or number 0.1 to 5)
	GrayScale        bool   // e-grayscale
	DefaultImage     string // di (Replaces all forward slashes with @@)
	ProgressiveImage bool   // pr
	Lossless         bool   // lo
	TrimEdges        any    // t (true or number 1-99)
	Border           string // b ("width_hexColorCode")
	ColorProfile     bool   // cp
	ImageMetadata    bool   // md
	Rotate           any    // rt (0 , 90 , 180 , 270 , 360 or auto)
	Radius           any    // r (number or max)
	BgColor          string // bg
	Original         bool   // orig
	Attachment       bool   // ik-attachment
	Contrast         bool
	Sharpen          any // bool or number
	Overlay          *Overlay
	Raw              string
}

// Overlay represents transformation's overlay options
type Overlay struct {
	X          *int         // ox
	Y          *int         // oy
	Height     *int         // oh
	Width      *int         // ow
	Background string       // obg
	Focus      OverlayFocus // ofo
	// Text overlay options
	Text               string // ot
	TextEncoded        string
	TextWidth          *int      // otw
	TextBackground     string    // otbg
	TextPadding        string    // otp ("40" or "40_60" or "40_60_50_10")
	TextInnerAlignment Alignment // otia
	TextColor          string    // otc
	TextFont           string    // otf
	TextSize           *int      // ots
	TextTypography     string    // ott (i, b or ib)=(italic, bold, both)
	Radius             int       // or
	// Image overlay options
	Image            string      // oi (replaces all forward slashes with @@)
	ImageAspectRatio AspectRatio // oiar
	ImageBorder      string      // oib (width_colorcode: 5_red, 10_FF0000)
	ImageDPR         *float32    // oidpr (0.1 to 5)
	ImageQuality     *int        // oiq
	ImageCropping    ImageCropping
	ImageX           *int  //oix
	ImageY           *int  //oiy
	ImageXc          *int  //oixc
	ImageYc          *int  //oixy
	Trim             *bool //oit-true
}
```

[See full documentation](https://docs.imagekit.io/features/image-transformations) for transformation options.

## File-Upload

The SDK uploader package provides a simple interface using the `.upload()` method to upload files to the ImageKit Media Library. It accepts all the parameters supported by the [ImageKit Upload API](https://docs.imagekit.io/api-reference/upload-file-api/server-side-file-upload).

The upload() method accepts file and UploadParam. File param can be base64 encode image, url or io.Reader. This method returns `UploadResponse` object and `err` if any. You can pass other parameters supported by the ImageKit upload API using the same parameter name as specified in the upload API documentation. For example, to specify tags for a file at the time of upload, use the tags parameter as specified in the [documentation here](https://docs.imagekit.io/api-reference/upload-file-api/server-side-file-upload).

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
List files in media library, optionally filter and sort using `FileParams`.

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
Adds given tags to multiple files. Accepts slices of tags and file ids. Returns slice of file ids. [API documentation here](https://docs.imagekit.io/api-reference/media-api/add-tags-bulk).

```
resp, err := ik.Media.AddTags(ctx, media.TagsParam{
    FileIds: []string{"file_id_1", "file_id_2"},
    Tags: []string{"tag_1", "tag_2"},
})
```

### 7. Remove Tags (bulk)
Removes tags from multiple files. Returns slice of file ids updated. [API documentation here](https://docs.imagekit.io/api-reference/media-api/remove-tags-bulk).

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
Deletes given version of the file. [API documentation here](https://docs.imagekit.io/api-reference/media-api/delete-file-version).
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
Restore file version to a different version of a file as per [API documentation here](https://docs.imagekit.io/api-reference/media-api/restore-file-version).
Accepts string type file id and version id.
```
resp, err := ik.Media.RestoreVersion(ctx, media.FileVersionsParam{
    FileId: "file_id",
    VersionId: "version_id",
})
```

### 16. Create Folder
Creates a new folder as per [API documentation here](https://docs.imagekit.io/api-reference/media-api/create-folder). `err` is not nil when response is not 201.

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
Get status of a bulk job operation by job id.  Accepts string type job id. [API documentation here](https://docs.imagekit.io/api-reference/media-api/copy-move-folder-status).

```
resp, err := ik.BulkJobStatus(ctx, "job_id")
```

### 21. Purge Cache
This will purge given url's CDN and ImageKit.io's internal cache as per [API documentation here](https://docs.imagekit.io/api-reference/media-api/purge-cache).

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

### 1. SignToken
This method generates a signature for given token and timestamp using the configured private key. It is useful for client side file upload to authenticate requests. `Token` is a random string. `Expires` is a unix timestamp by which token should expire.  `Token` and `Expires` both are optional parameters. `Token` defaults to auto generated UUID string. `Expires` defaults to a current time + 30 minutes value.

```
// Using auto generated token and expiration
resp := ik.SignToken(imagekit.SignTokenParam{})

// Using specific token and expiration
resp := ik.SignToken(imagekit.SignTokenParam{
    Token: "token-string",
    Expires: 1655379249,
})

```

## Rate Limits
Except for upload API, all ImageKit APIs are rate limited to avoid excessive request rates. 

Whenever backend api returns 429 status code, error of type `ErrTooManyRequests` is returned, which can be tested with `errors.Is`. The rate limit detail can be retrieved from response metadata header. Please sleep/pause for the number of milliseconds specified by the value of `resp.ResponseMetaData.Header["X-RateLimit-Reset"]` property before making additional requests to that endpoint.

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
