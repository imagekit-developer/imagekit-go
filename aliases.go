// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"github.com/imagekit-developer/imagekit-go/internal/apierror"
	"github.com/imagekit-developer/imagekit-go/packages/param"
	"github.com/imagekit-developer/imagekit-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// Array of extensions to be applied to the asset. Each extension can be configured
// with specific parameters based on the extension type.
//
// This is an alias to an internal type.
type ExtensionsParam = shared.ExtensionsParam

// This is an alias to an internal type.
type ExtensionUnionParam = shared.ExtensionUnionParam

// This is an alias to an internal type.
type ExtensionRemoveBgParam = shared.ExtensionRemoveBgParam

// This is an alias to an internal type.
type ExtensionRemoveBgOptionsParam = shared.ExtensionRemoveBgOptionsParam

// This is an alias to an internal type.
type ExtensionAutoTaggingParam = shared.ExtensionAutoTaggingParam

// This is an alias to an internal type.
type ExtensionAIAutoDescriptionParam = shared.ExtensionAIAutoDescriptionParam
