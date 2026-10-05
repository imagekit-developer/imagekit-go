// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/imagekit-developer/imagekit-go/v2/internal/apijson"
	"github.com/imagekit-developer/imagekit-go/v2/internal/requestconfig"
	"github.com/imagekit-developer/imagekit-go/v2/option"
	"github.com/imagekit-developer/imagekit-go/v2/packages/param"
	"github.com/imagekit-developer/imagekit-go/v2/packages/respjson"
)

// AccountWebhookService contains methods and other services that help with
// interacting with the ImageKit API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountWebhookService] method instead.
type AccountWebhookService struct {
	Options []option.RequestOption
}

// NewAccountWebhookService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountWebhookService(opts ...option.RequestOption) (r AccountWebhookService) {
	r = AccountWebhookService{}
	r.Options = opts
	return
}

// Creates a new webhook and returns the created object, including the generated
// signing `secret`.
//
// ImageKit sends a `POST` request to the webhook `endpoint` whenever one of the
// subscribed `events` occurs. Use the `secret` to verify the signature of each
// request. Learn more about [webhooks](https://imagekit.io/docs/webhooks).
//
// You can create up to 3 webhooks per account.
func (r *AccountWebhookService) New(ctx context.Context, body AccountWebhookNewParams, opts ...option.RequestOption) (res *Webhook, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/accounts/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Updates the webhook identified by `id` and returns the updated object. Only the
// fields included in the request body are changed.
//
// When `events` is provided, it replaces the existing list of subscribed events.
// The signing `secret` can't be changed.
func (r *AccountWebhookService) Update(ctx context.Context, id string, body AccountWebhookUpdateParams, opts ...option.RequestOption) (res *Webhook, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/webhooks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns an array of all webhooks configured for your account.
func (r *AccountWebhookService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Webhook, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/accounts/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Permanently deletes the webhook identified by `id`. ImageKit stops sending
// events to its endpoint.
func (r *AccountWebhookService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/accounts/webhooks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Retrieves the webhook identified by `id`.
func (r *AccountWebhookService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Webhook, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/accounts/webhooks/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// A webhook is an HTTP endpoint that ImageKit notifies when subscribed events
// occur in your account. Learn more about
// [webhooks](https://imagekit.io/docs/webhooks).
type Webhook struct {
	// Unique identifier for a webhook.
	ID string `json:"id" api:"required"`
	// ISO 8601 timestamp of when the webhook was created.
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Whether the webhook is enabled. When `false`, ImageKit doesn't send any events
	// to the `endpoint`.
	Enabled bool `json:"enabled" api:"required"`
	// The URL where ImageKit sends webhook events as `POST` requests. Must use the
	// `http` or `https` protocol, and must be unique across all webhooks for your
	// account.
	Endpoint string `json:"endpoint" api:"required" format:"uri"`
	// The event types this webhook is subscribed to. ImageKit sends a request to the
	// `endpoint` only for these events.
	Events []WebhookEventType `json:"events" api:"required"`
	// Signing secret generated by ImageKit, used to verify that webhook requests come
	// from ImageKit. Starts with `whsec_`. This value can't be set or changed through
	// the API. Learn more about
	// [verifying webhook signatures](https://imagekit.io/docs/webhooks#verify-webhook-signature).
	Secret string `json:"secret" api:"required"`
	// ISO 8601 timestamp of when the webhook was last updated.
	UpdatedAt time.Time `json:"updatedAt" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Enabled     respjson.Field
		Endpoint    respjson.Field
		Events      respjson.Field
		Secret      respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Webhook) RawJSON() string { return r.JSON.raw }
func (r *Webhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A webhook event type. Learn more about the payload of each
// [webhook event](https://imagekit.io/docs/webhooks#list-of-events).
type WebhookEventType string

const (
	WebhookEventTypeVideoTransformationAccepted WebhookEventType = "video.transformation.accepted"
	WebhookEventTypeVideoTransformationReady    WebhookEventType = "video.transformation.ready"
	WebhookEventTypeVideoTransformationError    WebhookEventType = "video.transformation.error"
	WebhookEventTypeUploadPreTransformSuccess   WebhookEventType = "upload.pre-transform.success"
	WebhookEventTypeUploadPreTransformError     WebhookEventType = "upload.pre-transform.error"
	WebhookEventTypeUploadPostTransformSuccess  WebhookEventType = "upload.post-transform.success"
	WebhookEventTypeUploadPostTransformError    WebhookEventType = "upload.post-transform.error"
	WebhookEventTypeFileCreated                 WebhookEventType = "file.created"
	WebhookEventTypeFileUpdated                 WebhookEventType = "file.updated"
	WebhookEventTypeFileDeleted                 WebhookEventType = "file.deleted"
	WebhookEventTypeFileVersionCreated          WebhookEventType = "file-version.created"
	WebhookEventTypeFileVersionDeleted          WebhookEventType = "file-version.deleted"
)

type AccountWebhookNewParams struct {
	// The URL where ImageKit sends webhook events as `POST` requests. Must use the
	// `http` or `https` protocol, and must be unique across all webhooks for your
	// account.
	Endpoint string `json:"endpoint" api:"required" format:"uri"`
	// The event types this webhook is subscribed to. ImageKit sends a request to the
	// `endpoint` only for these events.
	Events []WebhookEventType `json:"events,omitzero" api:"required"`
	// Whether the webhook is enabled. When `false`, ImageKit doesn't send any events
	// to the `endpoint`.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	paramObj
}

func (r AccountWebhookNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountWebhookNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountWebhookNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWebhookUpdateParams struct {
	// Whether the webhook is enabled. Omit to leave the current value unchanged.
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// The URL where ImageKit sends webhook events as `POST` requests. Must use the
	// `http` or `https` protocol, and must be unique across all webhooks for your
	// account.
	Endpoint param.Opt[string] `json:"endpoint,omitzero" format:"uri"`
	// The event types this webhook is subscribed to. Replaces the existing list. Omit
	// to leave the current value unchanged.
	Events []WebhookEventType `json:"events,omitzero"`
	paramObj
}

func (r AccountWebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AccountWebhookUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccountWebhookUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
