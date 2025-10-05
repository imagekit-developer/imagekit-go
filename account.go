// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"github.com/imagekit-developer/imagekit-go/option"
)

// AccountService contains methods and other services that help with interacting
// with the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountService] method instead.
type AccountService struct {
	Options      []option.RequestOption
	Usage        AccountUsageService
	Origins      AccountOriginService
	URLEndpoints AccountURLEndpointService
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r AccountService) {
	r = AccountService{}
	r.Options = opts
	r.Usage = NewAccountUsageService(opts...)
	r.Origins = NewAccountOriginService(opts...)
	r.URLEndpoints = NewAccountURLEndpointService(opts...)
	return
}
