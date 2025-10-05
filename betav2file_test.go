// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/internal/testutil"
	"github.com/imagekit-developer/imagekit-go/option"
	"github.com/imagekit-developer/imagekit-go/shared"
)

func TestBetaV2FileUploadWithOptionalParams(t *testing.T) {
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
		option.WithPrivateKey("My Private Key"),
		option.WithPassword("My Password"),
	)
	_, err := client.Beta.V2.Files.Upload(context.TODO(), imagekit.BetaV2FileUploadParams{
		File:              io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		FileName:          "fileName",
		Token:             imagekit.String("token"),
		Checks:            imagekit.String("\"request.folder\" : \"marketing/\"\n"),
		CustomCoordinates: imagekit.String("customCoordinates"),
		CustomMetadata: map[string]any{
			"brand": "bar",
			"color": "bar",
		},
		Description: imagekit.String("Running shoes"),
		Extensions: shared.ExtensionsParam{shared.ExtensionUnionParam{
			OfRemoveBg: &shared.ExtensionRemoveBgParam{
				Options: shared.ExtensionRemoveBgOptionsParam{
					AddShadow:        imagekit.Bool(true),
					BgColor:          imagekit.String("bg_color"),
					BgImageURL:       imagekit.String("bg_image_url"),
					Semitransparency: imagekit.Bool(true),
				},
			},
		}, shared.ExtensionUnionParam{
			OfAutoTagging: &shared.ExtensionAutoTaggingParam{
				MaxTags:       5,
				MinConfidence: 95,
				Name:          "google-auto-tagging",
			},
		}, shared.ExtensionUnionParam{
			OfAIAutoDescription: &shared.ExtensionAIAutoDescriptionParam{},
		}},
		Folder:                  imagekit.String("folder"),
		IsPrivateFile:           imagekit.Bool(true),
		IsPublished:             imagekit.Bool(true),
		OverwriteAITags:         imagekit.Bool(true),
		OverwriteCustomMetadata: imagekit.Bool(true),
		OverwriteFile:           imagekit.Bool(true),
		OverwriteTags:           imagekit.Bool(true),
		ResponseFields:          []string{"tags", "customCoordinates", "isPrivateFile"},
		Tags:                    []string{"t-shirt", "round-neck", "men"},
		Transformation: imagekit.BetaV2FileUploadParamsTransformation{
			Post: []imagekit.BetaV2FileUploadParamsTransformationPostUnion{{
				OfThumbnail: &imagekit.BetaV2FileUploadParamsTransformationPostThumbnail{
					Value: imagekit.String("w-150,h-150"),
				},
			}, {
				OfAbs: &imagekit.BetaV2FileUploadParamsTransformationPostAbs{
					Protocol: "dash",
					Value:    "sr-240_360_480_720_1080",
				},
			}},
			Pre: imagekit.String("w-300,h-300,q-80"),
		},
		UseUniqueFileName: imagekit.Bool(true),
		WebhookURL:        imagekit.String("https://example.com"),
	})
	if err != nil {
		var apierr *imagekit.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
