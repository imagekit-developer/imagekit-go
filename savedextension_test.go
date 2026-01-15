// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/imagekit-developer/imagekit-go/v2"
	"github.com/imagekit-developer/imagekit-go/v2/internal/testutil"
	"github.com/imagekit-developer/imagekit-go/v2/option"
	"github.com/imagekit-developer/imagekit-go/v2/shared"
)

func TestSavedExtensionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.SavedExtensions.New(context.TODO(), imagekit.SavedExtensionNewParams{
		Config: shared.ExtensionConfigUnionParam{
			OfRemoveBg: &shared.ExtensionConfigRemoveBgParam{
				Options: shared.ExtensionConfigRemoveBgOptionsParam{
					AddShadow:        imagekit.Bool(true),
					BgColor:          imagekit.String("bg_color"),
					BgImageURL:       imagekit.String("bg_image_url"),
					Semitransparency: imagekit.Bool(true),
				},
			},
		},
		Description: "Analyzes vehicle images for type, condition, and quality assessment",
		Name:        "Car Quality Analysis",
	})
	if err != nil {
		var apierr *imagekit.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSavedExtensionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.SavedExtensions.Update(
		context.TODO(),
		"id",
		imagekit.SavedExtensionUpdateParams{
			Config: shared.ExtensionConfigUnionParam{
				OfRemoveBg: &shared.ExtensionConfigRemoveBgParam{
					Options: shared.ExtensionConfigRemoveBgOptionsParam{
						AddShadow:        imagekit.Bool(true),
						BgColor:          imagekit.String("bg_color"),
						BgImageURL:       imagekit.String("bg_image_url"),
						Semitransparency: imagekit.Bool(true),
					},
				},
			},
			Description: imagekit.String("x"),
			Name:        imagekit.String("x"),
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

func TestSavedExtensionList(t *testing.T) {
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
	_, err := client.SavedExtensions.List(context.TODO())
	if err != nil {
		var apierr *imagekit.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSavedExtensionDelete(t *testing.T) {
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
	err := client.SavedExtensions.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *imagekit.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSavedExtensionGet(t *testing.T) {
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
	_, err := client.SavedExtensions.Get(context.TODO(), "id")
	if err != nil {
		var apierr *imagekit.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
