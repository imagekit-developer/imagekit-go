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
    file: "htts://example.com/myimage.jpg",
    fileName: "my image.jpg",
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
    Skip: 10, // int
    Limit: 500, // int
    SearchQuery: "createdAt >= "7d" AND size > \"2mb\"", // string
})
```

### 2. Get File Details
Accepts the file ID and fetches the details as per the [API documentation here](https://docs.imagekit.io/api-reference/media-api/get-file-details).
```
file, err := imgkit.GetFileDetails(ctx, fileId)
```
### 3. Get File Metadata
Accepts the file ID or URL and fetches the metadata as per the [API documentation here](https://docs.imagekit.io/api-reference/metadata-api/get-image-metadata-for-uploaded-media-files).
```
meta, err := imgkit.GetAssetMetadata(ctx, fileId)
```
### 4. Update File Details
Update parameters associated with the file as per the [API documentation here](https://docs.imagekit.io/api-reference/media-api/update-file-details).

```
resp, err := imgkit.UpdateAsset(ctx, assets.AssetParams{
    FileId: fileId, // string
    Tags: tags, // strings slice
})
```
### 5. Delete File
Delete a file as per the [API documentation here](https://docs.imagekit.io/api-reference/media-api/delete-file).
```
resp, err := imgkit.DeleteAsset(ctx,
    fileId, // string
)
```
### 6. Bulk Delete Files
Delete multiple files as per the [API documentation here](https://docs.imagekit.io/api-reference/media-api/delete-files-bulk).
```
resp, err := imgkit.BulkDeleteAssets(ctx,
    fileIds, // strings slice
)

```
