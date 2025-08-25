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

// Object containing details of a file or file version.
//
// This is an alias to an internal type.
type File = shared.File

// This is an alias to an internal type.
type FileAITag = shared.FileAITag

// Type of the asset.
//
// This is an alias to an internal type.
type FileType = shared.FileType

// Equals "file"
const FileTypeFile = shared.FileTypeFile

// Equals "file-version"
const FileTypeFileVersion = shared.FileTypeFileVersion

// An object with details of the file version.
//
// This is an alias to an internal type.
type FileVersionInfo = shared.FileVersionInfo

// This is an alias to an internal type.
type Folder = shared.Folder

// Type of the asset.
//
// This is an alias to an internal type.
type FolderType = shared.FolderType

// Equals "folder"
const FolderTypeFolder = shared.FolderTypeFolder
