// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stainless-sdks/imagekit-go"
	"github.com/stainless-sdks/imagekit-go/internal/testutil"
	"github.com/stainless-sdks/imagekit-go/option"
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
			OfUpdateFileDetails: &imagekit.FileUpdateParamsBodyUpdateFileDetails{
				CustomCoordinates: imagekit.String("10,10,100,100"),
				CustomMetadata: map[string]any{
					"brand": "bar",
					"color": "bar",
				},
				Description: imagekit.String("description"),
				Extensions: []imagekit.FileUpdateParamsBodyUpdateFileDetailsExtensionUnion{{
					OfRemoveBg: &imagekit.FileUpdateParamsBodyUpdateFileDetailsExtensionRemoveBg{
						Options: imagekit.FileUpdateParamsBodyUpdateFileDetailsExtensionRemoveBgOptions{
							AddShadow:        imagekit.Bool(true),
							BgColor:          imagekit.String("bg_color"),
							BgImageURL:       imagekit.String("bg_image_url"),
							Semitransparency: imagekit.Bool(true),
						},
					},
				}, {
					OfAutoTagging: &imagekit.FileUpdateParamsBodyUpdateFileDetailsExtensionAutoTagging{
						MaxTags:       10,
						MinConfidence: 80,
						Name:          "google-auto-tagging",
					},
				}, {
					OfAutoTagging: &imagekit.FileUpdateParamsBodyUpdateFileDetailsExtensionAutoTagging{
						MaxTags:       10,
						MinConfidence: 80,
						Name:          "aws-auto-tagging",
					},
				}, {
					OfAIAutoDescription: &imagekit.FileUpdateParamsBodyUpdateFileDetailsExtensionAIAutoDescription{},
				}},
				RemoveAITags: imagekit.FileUpdateParamsBodyUpdateFileDetailsRemoveAITagsUnion{
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
	_, err := client.Files.Upload(context.TODO(), imagekit.FileUploadParams{
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
		Expire:      imagekit.Int(0),
		Extensions: []imagekit.FileUploadParamsExtensionUnion{{
			OfRemoveBackground: &imagekit.FileUploadParamsExtensionRemoveBackground{
				Name: "remove-bg",
				Options: imagekit.FileUploadParamsExtensionRemoveBackgroundOptions{
					AddShadow:        imagekit.Bool(true),
					BgColor:          imagekit.String("bg_color"),
					BgImageURL:       imagekit.String("bg_image_url"),
					Semitransparency: imagekit.Bool(true),
				},
			},
		}, {
			OfAutoTagging: &imagekit.FileUploadParamsExtensionAutoTagging{
				MaxTags:       5,
				MinConfidence: 95,
				Name:          "google-auto-tagging",
			},
		}},
		Folder:                  imagekit.String("folder"),
		IsPrivateFile:           imagekit.Bool(true),
		IsPublished:             imagekit.Bool(true),
		OverwriteAITags:         imagekit.Bool(true),
		OverwriteCustomMetadata: imagekit.Bool(true),
		OverwriteFile:           imagekit.Bool(true),
		OverwriteTags:           imagekit.Bool(true),
		PublicKey:               imagekit.String("publicKey"),
		ResponseFields:          []string{"tags", "customCoordinates", "isPrivateFile"},
		Signature:               imagekit.String("signature"),
		Tags:                    []string{"t-shirt", "round-neck", "men"},
		Transformation: imagekit.FileUploadParamsTransformation{
			Post: []imagekit.FileUploadParamsTransformationPostUnion{{
				OfGenerateAThumbnail: &imagekit.FileUploadParamsTransformationPostGenerateAThumbnail{
					Type:  "thumbnail",
					Value: imagekit.String("w-150,h-150"),
				},
			}, {
				OfAdaptiveBitrateStreaming: &imagekit.FileUploadParamsTransformationPostAdaptiveBitrateStreaming{
					Protocol: "dash",
					Type:     "abs",
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
