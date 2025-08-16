// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/imagekit-go"
	"github.com/stainless-sdks/imagekit-go/internal/testutil"
	"github.com/stainless-sdks/imagekit-go/option"
)

func TestFileDetailGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := imagekit.NewClient(
		option.WithBaseURL(baseURL),
		option.WithPrivateAPIKey("My Private API Key"),
		option.WithPassword("My Password"),
	)
	_, err := client.Files.Details.Get(context.TODO(), "fileId")
	if err != nil {
		var apierr *imagekit.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFileDetailUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := imagekit.NewClient(
		option.WithBaseURL(baseURL),
		option.WithPrivateAPIKey("My Private API Key"),
		option.WithPassword("My Password"),
	)
	_, err := client.Files.Details.Update(
		context.TODO(),
		"fileId",
		imagekit.FileDetailUpdateParams{
			OfUpdateFileDetails: &imagekit.FileDetailUpdateParamsBodyUpdateFileDetails{
				CustomCoordinates: imagekit.String("10,10,100,100"),
				CustomMetadata: map[string]interface{}{
					"brand": "Nike",
					"color": "red",
				},
				Extensions: []imagekit.FileDetailUpdateParamsBodyUpdateFileDetailsExtensionUnion{{
					OfRemoveBackground: &imagekit.FileDetailUpdateParamsBodyUpdateFileDetailsExtensionRemoveBackground{
						Name: "remove-bg",
						Options: imagekit.FileDetailUpdateParamsBodyUpdateFileDetailsExtensionRemoveBackgroundOptions{
							AddShadow:        imagekit.Bool(true),
							BgColor:          imagekit.String("bg_color"),
							BgImageURL:       imagekit.String("bg_image_url"),
							Semitransparency: imagekit.Bool(true),
						},
					},
				}, {
					OfAutoTagging: &imagekit.FileDetailUpdateParamsBodyUpdateFileDetailsExtensionAutoTagging{
						MaxTags:       10,
						MinConfidence: 80,
						Name:          "google-auto-tagging",
					},
				}, {
					OfAutoTagging: &imagekit.FileDetailUpdateParamsBodyUpdateFileDetailsExtensionAutoTagging{
						MaxTags:       10,
						MinConfidence: 80,
						Name:          "aws-auto-tagging",
					},
				}},
				RemoveAITags: imagekit.FileDetailUpdateParamsBodyUpdateFileDetailsRemoveAITagsUnion{
					OfStringArray: []string{"car", "vehicle", "motorsports"},
				},
				Tags:       []string{"tag1", "tag2"},
				WebhookURL: imagekit.String("https://webhook.site/0d6b6c7a-8e5a-4b3a-8b7c-0d6b6c7a8e5a"),
			},
		},
	)
	if err != nil {
		var apierr *imagekit.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
