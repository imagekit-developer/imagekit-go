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
)

func TestFileUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Files.Update(
		context.TODO(),
		"fileId",
		imagekit.FileUpdateParams{
			OfUpdateFileDetails: &imagekit.FileUpdateParamsUpdateUpdateFileDetails{
				CustomCoordinates: imagekit.String("10,10,100,100"),
				CustomMetadata: map[string]any{
					"brand": "bar",
					"color": "bar",
				},
				Description: imagekit.String("description"),
				Extensions: []imagekit.FileUpdateParamsUpdateUpdateFileDetailsExtensionUnion{{
					OfRemoveBg: &imagekit.FileUpdateParamsUpdateUpdateFileDetailsExtensionRemoveBg{
						Options: imagekit.FileUpdateParamsUpdateUpdateFileDetailsExtensionRemoveBgOptions{
							AddShadow:        imagekit.Bool(true),
							BgColor:          imagekit.String("bg_color"),
							BgImageURL:       imagekit.String("bg_image_url"),
							Semitransparency: imagekit.Bool(true),
						},
					},
				}, {
					OfAutoTagging: &imagekit.FileUpdateParamsUpdateUpdateFileDetailsExtensionAutoTagging{
						MaxTags:       10,
						MinConfidence: 80,
						Name:          "google-auto-tagging",
					},
				}, {
					OfAutoTagging: &imagekit.FileUpdateParamsUpdateUpdateFileDetailsExtensionAutoTagging{
						MaxTags:       10,
						MinConfidence: 80,
						Name:          "aws-auto-tagging",
					},
				}, {
					OfAIAutoDescription: &imagekit.FileUpdateParamsUpdateUpdateFileDetailsExtensionAIAutoDescription{},
				}},
				RemoveAITags: imagekit.FileUpdateParamsUpdateUpdateFileDetailsRemoveAITagsUnion{
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

func TestFileDelete(t *testing.T) {
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
	err := client.Files.Delete(context.TODO(), "fileId")
	if err != nil {
		var apierr *imagekit.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFileCopyWithOptionalParams(t *testing.T) {
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
	_, err := client.Files.Copy(context.TODO(), imagekit.FileCopyParams{
		DestinationPath:     "/folder/to/copy/into/",
		SourceFilePath:      "/path/to/file.jpg",
		IncludeFileVersions: imagekit.Bool(false),
	})
	if err != nil {
		var apierr *imagekit.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFileGet(t *testing.T) {
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
	_, err := client.Files.Get(context.TODO(), "fileId")
	if err != nil {
		var apierr *imagekit.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFileMove(t *testing.T) {
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
	_, err := client.Files.Move(context.TODO(), imagekit.FileMoveParams{
		DestinationPath: "/folder/to/move/into/",
		SourceFilePath:  "/path/to/file.jpg",
	})
	if err != nil {
		var apierr *imagekit.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFileRenameWithOptionalParams(t *testing.T) {
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
	_, err := client.Files.Rename(context.TODO(), imagekit.FileRenameParams{
		FilePath:    "/path/to/file.jpg",
		NewFileName: "newFileName.jpg",
		PurgeCache:  imagekit.Bool(true),
	})
	if err != nil {
		var apierr *imagekit.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFileUploadWithOptionalParams(t *testing.T) {
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
	_, err := client.Files.Upload(context.TODO(), imagekit.FileUploadParams{})
	if err != nil {
		var apierr *imagekit.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
