// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/imagekit-developer/imagekit-go/v2/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type Abs string                         // Always "abs"
type AIAutoDescription string           // Always "ai-auto-description"
type Akamai string                      // Always "AKAMAI"
type AkeneoPim string                   // Always "AKENEO_PIM"
type All string                         // Always "all"
type AzureBlob string                   // Always "AZURE_BLOB"
type Cloudinary string                  // Always "CLOUDINARY"
type CloudinaryBackup string            // Always "CLOUDINARY_BACKUP"
type Gcs string                         // Always "GCS"
type GifToVideo string                  // Always "gif-to-video"
type Image string                       // Always "image"
type Imgix string                       // Always "IMGIX"
type Max string                         // Always "max"
type RemoveBg string                    // Always "remove-bg"
type S3 string                          // Always "S3"
type S3Compatible string                // Always "S3_COMPATIBLE"
type SolidColor string                  // Always "solidColor"
type Subtitle string                    // Always "subtitle"
type Text string                        // Always "text"
type Thumbnail string                   // Always "thumbnail"
type Transformation string              // Always "transformation"
type UploadPostTransformError string    // Always "upload.post-transform.error"
type UploadPostTransformSuccess string  // Always "upload.post-transform.success"
type UploadPreTransformError string     // Always "upload.pre-transform.error"
type UploadPreTransformSuccess string   // Always "upload.pre-transform.success"
type Video string                       // Always "video"
type VideoTransformationAccepted string // Always "video.transformation.accepted"
type VideoTransformationError string    // Always "video.transformation.error"
type VideoTransformationReady string    // Always "video.transformation.ready"
type WebFolder string                   // Always "WEB_FOLDER"
type WebProxy string                    // Always "WEB_PROXY"

func (c Abs) Default() Abs                             { return "abs" }
func (c AIAutoDescription) Default() AIAutoDescription { return "ai-auto-description" }
func (c Akamai) Default() Akamai                       { return "AKAMAI" }
func (c AkeneoPim) Default() AkeneoPim                 { return "AKENEO_PIM" }
func (c All) Default() All                             { return "all" }
func (c AzureBlob) Default() AzureBlob                 { return "AZURE_BLOB" }
func (c Cloudinary) Default() Cloudinary               { return "CLOUDINARY" }
func (c CloudinaryBackup) Default() CloudinaryBackup   { return "CLOUDINARY_BACKUP" }
func (c Gcs) Default() Gcs                             { return "GCS" }
func (c GifToVideo) Default() GifToVideo               { return "gif-to-video" }
func (c Image) Default() Image                         { return "image" }
func (c Imgix) Default() Imgix                         { return "IMGIX" }
func (c Max) Default() Max                             { return "max" }
func (c RemoveBg) Default() RemoveBg                   { return "remove-bg" }
func (c S3) Default() S3                               { return "S3" }
func (c S3Compatible) Default() S3Compatible           { return "S3_COMPATIBLE" }
func (c SolidColor) Default() SolidColor               { return "solidColor" }
func (c Subtitle) Default() Subtitle                   { return "subtitle" }
func (c Text) Default() Text                           { return "text" }
func (c Thumbnail) Default() Thumbnail                 { return "thumbnail" }
func (c Transformation) Default() Transformation       { return "transformation" }
func (c UploadPostTransformError) Default() UploadPostTransformError {
	return "upload.post-transform.error"
}
func (c UploadPostTransformSuccess) Default() UploadPostTransformSuccess {
	return "upload.post-transform.success"
}
func (c UploadPreTransformError) Default() UploadPreTransformError {
	return "upload.pre-transform.error"
}
func (c UploadPreTransformSuccess) Default() UploadPreTransformSuccess {
	return "upload.pre-transform.success"
}
func (c Video) Default() Video { return "video" }
func (c VideoTransformationAccepted) Default() VideoTransformationAccepted {
	return "video.transformation.accepted"
}
func (c VideoTransformationError) Default() VideoTransformationError {
	return "video.transformation.error"
}
func (c VideoTransformationReady) Default() VideoTransformationReady {
	return "video.transformation.ready"
}
func (c WebFolder) Default() WebFolder { return "WEB_FOLDER" }
func (c WebProxy) Default() WebProxy   { return "WEB_PROXY" }

func (c Abs) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c AIAutoDescription) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c Akamai) MarshalJSON() ([]byte, error)                      { return marshalString(c) }
func (c AkeneoPim) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c All) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c AzureBlob) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c Cloudinary) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c CloudinaryBackup) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c Gcs) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c GifToVideo) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c Image) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c Imgix) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c Max) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c RemoveBg) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c S3) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c S3Compatible) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c SolidColor) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c Subtitle) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c Text) MarshalJSON() ([]byte, error)                        { return marshalString(c) }
func (c Thumbnail) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c Transformation) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c UploadPostTransformError) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c UploadPostTransformSuccess) MarshalJSON() ([]byte, error)  { return marshalString(c) }
func (c UploadPreTransformError) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c UploadPreTransformSuccess) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c Video) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c VideoTransformationAccepted) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c VideoTransformationError) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c VideoTransformationReady) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c WebFolder) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c WebProxy) MarshalJSON() ([]byte, error)                    { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
