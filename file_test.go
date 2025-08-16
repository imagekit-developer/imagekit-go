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

func TestFileListWithOptionalParams(t *testing.T) {
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
	_, err := client.Files.List(context.TODO(), imagekit.FileListParams{
		FileType:    imagekit.String("fileType"),
		Limit:       imagekit.String("limit"),
		Path:        imagekit.String("path"),
		SearchQuery: imagekit.String("searchQuery"),
		Skip:        imagekit.String("skip"),
		Sort:        imagekit.String("sort"),
		Type:        imagekit.FileListParamsTypeFile,
	})
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

func TestFileAddTags(t *testing.T) {
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
	_, err := client.Files.AddTags(context.TODO(), imagekit.FileAddTagsParams{
		FileIDs: []string{"598821f949c0a938d57563bd", "598821f949c0a938d57563be"},
		Tags:    []string{"t-shirt", "round-neck", "sale2019"},
	})
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

func TestFileRemoveAITags(t *testing.T) {
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
	_, err := client.Files.RemoveAITags(context.TODO(), imagekit.FileRemoveAITagsParams{
		AITags:  []string{"t-shirt", "round-neck", "sale2019"},
		FileIDs: []string{"598821f949c0a938d57563bd", "598821f949c0a938d57563be"},
	})
	if err != nil {
		var apierr *imagekit.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFileRemoveTags(t *testing.T) {
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
	_, err := client.Files.RemoveTags(context.TODO(), imagekit.FileRemoveTagsParams{
		FileIDs: []string{"598821f949c0a938d57563bd", "598821f949c0a938d57563be"},
		Tags:    []string{"t-shirt", "round-neck", "sale2019"},
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

func TestFileUploadV1WithOptionalParams(t *testing.T) {
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
	_, err := client.Files.UploadV1(context.TODO(), imagekit.FileUploadV1Params{
		File:                    "https://www.example.com/rest-of-the-image-path.jpg",
		FileName:                "fileName",
		Token:                   imagekit.String("token"),
		Checks:                  imagekit.String("\"request.folder\" : \"marketing/\"\n"),
		CustomCoordinates:       imagekit.String("customCoordinates"),
		CustomMetadata:          imagekit.String("\"\n  {\n    \"brand\": \"Nike\",\n    \"color\":\"red\"\n  }\n\"\n"),
		Expire:                  imagekit.String("expire"),
		Extensions:              imagekit.String("\"\n[\n  {\"name\":\"remove-bg\",\"options\":{\"add_shadow\":true,\"bg_colour\":\"green\"}},\n  {\"name\":\"google-auto-tagging\",\"maxTags\":5,\"minConfidence\":95}\n]\n\"\n"),
		Folder:                  imagekit.String("folder"),
		IsPrivateFile:           imagekit.FileUploadV1ParamsIsPrivateFileTrue,
		IsPublished:             imagekit.FileUploadV1ParamsIsPublishedTrue,
		OverwriteAITags:         imagekit.FileUploadV1ParamsOverwriteAITagsTrue,
		OverwriteCustomMetadata: imagekit.FileUploadV1ParamsOverwriteCustomMetadataTrue,
		OverwriteFile:           imagekit.String("overwriteFile"),
		OverwriteTags:           imagekit.FileUploadV1ParamsOverwriteTagsTrue,
		PublicKey:               imagekit.String("publicKey"),
		ResponseFields:          imagekit.String("responseFields"),
		Signature:               imagekit.String("signature"),
		Tags:                    imagekit.String("t-shirt,round-neck,men"),
		Transformation:          imagekit.String("'{\"pre\":\"width:300,height:300,quality:80\",\"post\":[{\"type\":\"thumbnail\",\"value\":\"width:100,height:100\"}]}'\n"),
		UseUniqueFileName:       imagekit.FileUploadV1ParamsUseUniqueFileNameTrue,
		WebhookURL:              imagekit.String("webhookUrl"),
	})
	if err != nil {
		var apierr *imagekit.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFileUploadV2WithOptionalParams(t *testing.T) {
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
	_, err := client.Files.UploadV2(context.TODO(), imagekit.FileUploadV2Params{
		File:                    "https://www.example.com/rest-of-the-image-path.jpg",
		FileName:                "fileName",
		Token:                   imagekit.String("token"),
		Checks:                  imagekit.String("\"request.folder\" : \"marketing/\"\n"),
		CustomCoordinates:       imagekit.String("customCoordinates"),
		CustomMetadata:          imagekit.String("\"\n  {\n    \"brand\": \"Nike\",\n    \"color\":\"red\"\n  }\n\"\n"),
		Extensions:              imagekit.String("\"\n[\n  {\"name\":\"remove-bg\",\"options\":{\"add_shadow\":true,\"bg_colour\":\"green\"}},\n  {\"name\":\"google-auto-tagging\",\"maxTags\":5,\"minConfidence\":95}\n]\n\"\n"),
		Folder:                  imagekit.String("folder"),
		IsPrivateFile:           imagekit.FileUploadV2ParamsIsPrivateFileTrue,
		IsPublished:             imagekit.FileUploadV2ParamsIsPublishedTrue,
		OverwriteAITags:         imagekit.FileUploadV2ParamsOverwriteAITagsTrue,
		OverwriteCustomMetadata: imagekit.FileUploadV2ParamsOverwriteCustomMetadataTrue,
		OverwriteFile:           imagekit.String("overwriteFile"),
		OverwriteTags:           imagekit.FileUploadV2ParamsOverwriteTagsTrue,
		ResponseFields:          imagekit.String("responseFields"),
		Tags:                    imagekit.String("t-shirt,round-neck,men"),
		Transformation:          imagekit.String("'{\"pre\":\"width:300,height:300,quality:80\",\"post\":[{\"type\":\"thumbnail\",\"value\":\"width:100,height:100\"}]}'\n"),
		UseUniqueFileName:       imagekit.FileUploadV2ParamsUseUniqueFileNameTrue,
		WebhookURL:              imagekit.String("webhookUrl"),
	})
	if err != nil {
		var apierr *imagekit.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
