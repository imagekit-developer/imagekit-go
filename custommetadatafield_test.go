// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/internal/testutil"
	"github.com/imagekit-developer/imagekit-go/option"
)

func TestCustomMetadataFieldNewWithOptionalParams(t *testing.T) {
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
	_, err := client.CustomMetadataFields.New(context.TODO(), imagekit.CustomMetadataFieldNewParams{
		Label: "price",
		Name:  "price",
		Schema: imagekit.CustomMetadataFieldNewParamsSchema{
			Type: "Number",
			DefaultValue: imagekit.CustomMetadataFieldNewParamsSchemaDefaultValueUnion{
				OfString: imagekit.String("string"),
			},
			IsValueRequired: imagekit.Bool(true),
			MaxLength:       imagekit.Float(0),
			MaxValue: imagekit.CustomMetadataFieldNewParamsSchemaMaxValueUnion{
				OfFloat: imagekit.Float(3000),
			},
			MinLength: imagekit.Float(0),
			MinValue: imagekit.CustomMetadataFieldNewParamsSchemaMinValueUnion{
				OfFloat: imagekit.Float(1000),
			},
			SelectOptions: []imagekit.CustomMetadataFieldNewParamsSchemaSelectOptionUnion{{
				OfString: imagekit.String("small"),
			}, {
				OfString: imagekit.String("medium"),
			}, {
				OfString: imagekit.String("large"),
			}, {
				OfFloat: imagekit.Float(30),
			}, {
				OfFloat: imagekit.Float(40),
			}, {
				OfBool: imagekit.Bool(true),
			}},
		},
	})
	if err != nil {
		var apierr *imagekit.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomMetadataFieldUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.CustomMetadataFields.Update(
		context.TODO(),
		"id",
		imagekit.CustomMetadataFieldUpdateParams{
			Label: imagekit.String("price"),
			Schema: imagekit.CustomMetadataFieldUpdateParamsSchema{
				DefaultValue: imagekit.CustomMetadataFieldUpdateParamsSchemaDefaultValueUnion{
					OfString: imagekit.String("string"),
				},
				IsValueRequired: imagekit.Bool(true),
				MaxLength:       imagekit.Float(0),
				MaxValue: imagekit.CustomMetadataFieldUpdateParamsSchemaMaxValueUnion{
					OfFloat: imagekit.Float(3000),
				},
				MinLength: imagekit.Float(0),
				MinValue: imagekit.CustomMetadataFieldUpdateParamsSchemaMinValueUnion{
					OfFloat: imagekit.Float(1000),
				},
				SelectOptions: []imagekit.CustomMetadataFieldUpdateParamsSchemaSelectOptionUnion{{
					OfString: imagekit.String("small"),
				}, {
					OfString: imagekit.String("medium"),
				}, {
					OfString: imagekit.String("large"),
				}, {
					OfFloat: imagekit.Float(30),
				}, {
					OfFloat: imagekit.Float(40),
				}, {
					OfBool: imagekit.Bool(true),
				}},
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

func TestCustomMetadataFieldListWithOptionalParams(t *testing.T) {
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
	_, err := client.CustomMetadataFields.List(context.TODO(), imagekit.CustomMetadataFieldListParams{
		FolderPath:     imagekit.String("folderPath"),
		IncludeDeleted: imagekit.Bool(true),
	})
	if err != nil {
		var apierr *imagekit.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCustomMetadataFieldDelete(t *testing.T) {
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
	_, err := client.CustomMetadataFields.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *imagekit.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
