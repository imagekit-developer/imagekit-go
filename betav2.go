// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"github.com/stainless-sdks/imagekit-go/option"
)

// BetaV2Service contains methods and other services that help with interacting
// with the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaV2Service] method instead.
type BetaV2Service struct {
	Options []option.RequestOption
	Files   BetaV2FileService
}

// NewBetaV2Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBetaV2Service(opts ...option.RequestOption) (r BetaV2Service) {
	r = BetaV2Service{}
	r.Options = opts
	r.Files = NewBetaV2FileService(opts...)
	return
}
