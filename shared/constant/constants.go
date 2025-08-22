// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/stainless-sdks/imagekit-go/internal/encoding/json"
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

type Abs string               // Always "abs"
type AIAutoDescription string // Always "ai-auto-description"
type Akamai string            // Always "AKAMAI"
type Cloudinary string        // Always "CLOUDINARY"
type GifToVideo string        // Always "gif-to-video"
type Imgix string             // Always "IMGIX"
type RemoveBg string          // Always "remove-bg"
type Thumbnail string         // Always "thumbnail"
type Transformation string    // Always "transformation"

func (c Abs) Default() Abs                             { return "abs" }
func (c AIAutoDescription) Default() AIAutoDescription { return "ai-auto-description" }
func (c Akamai) Default() Akamai                       { return "AKAMAI" }
func (c Cloudinary) Default() Cloudinary               { return "CLOUDINARY" }
func (c GifToVideo) Default() GifToVideo               { return "gif-to-video" }
func (c Imgix) Default() Imgix                         { return "IMGIX" }
func (c RemoveBg) Default() RemoveBg                   { return "remove-bg" }
func (c Thumbnail) Default() Thumbnail                 { return "thumbnail" }
func (c Transformation) Default() Transformation       { return "transformation" }

func (c Abs) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c AIAutoDescription) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c Akamai) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c Cloudinary) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c GifToVideo) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Imgix) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c RemoveBg) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c Thumbnail) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c Transformation) MarshalJSON() ([]byte, error)    { return marshalString(c) }

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
