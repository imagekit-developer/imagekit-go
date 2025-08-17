// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"github.com/stainless-sdks/imagekit-go/internal/apierror"
	"github.com/stainless-sdks/imagekit-go/packages/param"
	"github.com/stainless-sdks/imagekit-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// Object containing Exif details.
//
// This is an alias to an internal type.
type ExifDetails = shared.ExifDetails

// Object containing EXIF image information.
//
// This is an alias to an internal type.
type ExifImage = shared.ExifImage

// Object containing GPS information.
//
// This is an alias to an internal type.
type Gps = shared.Gps

// JSON object.
//
// This is an alias to an internal type.
type Interoperability = shared.Interoperability

// Object containing Thumbnail information.
//
// This is an alias to an internal type.
type Thumbnail = shared.Thumbnail
