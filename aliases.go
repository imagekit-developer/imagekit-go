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

// This is an alias to an internal type.
type AutoDescriptionExtensionParam = shared.AutoDescriptionExtensionParam

// Specifies the auto description extension.
//
// This is an alias to an internal type.
type AutoDescriptionExtensionName = shared.AutoDescriptionExtensionName

// Equals "ai-auto-description"
const AutoDescriptionExtensionNameAIAutoDescription = shared.AutoDescriptionExtensionNameAIAutoDescription

// This is an alias to an internal type.
type AutoTaggingExtensionParam = shared.AutoTaggingExtensionParam

// Specifies the auto-tagging extension used.
//
// This is an alias to an internal type.
type AutoTaggingExtensionName = shared.AutoTaggingExtensionName

// Equals "google-auto-tagging"
const AutoTaggingExtensionNameGoogleAutoTagging = shared.AutoTaggingExtensionNameGoogleAutoTagging

// Equals "aws-auto-tagging"
const AutoTaggingExtensionNameAwsAutoTagging = shared.AutoTaggingExtensionNameAwsAutoTagging

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

// This is an alias to an internal type.
type RemovedotBgExtensionParam = shared.RemovedotBgExtensionParam

// Specifies the background removal extension.
//
// This is an alias to an internal type.
type RemovedotBgExtensionName = shared.RemovedotBgExtensionName

// Equals "remove-bg"
const RemovedotBgExtensionNameRemoveBg = shared.RemovedotBgExtensionNameRemoveBg

// This is an alias to an internal type.
type RemovedotBgExtensionOptionsParam = shared.RemovedotBgExtensionOptionsParam

// Object containing Thumbnail information.
//
// This is an alias to an internal type.
type Thumbnail = shared.Thumbnail
