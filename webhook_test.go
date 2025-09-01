// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit_test

import (
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/stainless-sdks/imagekit-go"
	"github.com/stainless-sdks/imagekit-go/option"
	standardwebhooks "github.com/standard-webhooks/standard-webhooks/libraries/go"
)

func TestWebhookUnwrap(t *testing.T) {
	client := imagekit.NewClient(
		option.WithWebhookSecret("whsec_c2VjcmV0Cg=="),
		option.WithPrivateAPIKey("My Private API Key"),
		option.WithPassword("My Password"),
	)
	payload := []byte(`{"id":"id","created_at":"2019-12-27T18:11:19.117Z","data":{"asset":{"url":"https://example.com"},"transformation":{"type":"video-transformation","options":{"audio_codec":"aac","auto_rotate":true,"format":"mp4","quality":0,"stream_protocol":"HLS","variants":["string"],"video_codec":"h264"}}},"request":{"url":"https://example.com","x_request_id":"x_request_id","user_agent":"user_agent"},"type":"video.transformation.accepted"}`)
	wh, err := standardwebhooks.NewWebhook("whsec_c2VjcmV0Cg==")
	if err != nil {
		t.Error("Failed to sign test webhook message")
	}
	msgID := "1"
	now := time.Now()
	sig, err := wh.Sign(msgID, now, payload)
	if err != nil {
		t.Error("Failed to sign test webhook message:", err)
	}
	headers := make(http.Header)
	headers.Set("webhook-signature", sig)
	headers.Set("webhook-id", msgID)
	headers.Set("webhook-timestamp", strconv.FormatInt(now.Unix(), 10))
	_, err = client.Webhooks.Unwrap(payload, headers)
	if err != nil {
		t.Error("Failed to unwrap webhook:", err)
	}
}
