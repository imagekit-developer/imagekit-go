// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"github.com/imagekit-developer/imagekit-go/v2/option"
)

// CacheService contains methods and other services that help with interacting with
// the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCacheService] method instead.
type CacheService struct {
	Options      []option.RequestOption
	Invalidation CacheInvalidationService
}

// NewCacheService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCacheService(opts ...option.RequestOption) (r CacheService) {
	r = CacheService{}
	r.Options = opts
	r.Invalidation = NewCacheInvalidationService(opts...)
	return
}
