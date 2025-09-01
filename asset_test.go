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

func TestAssetListWithOptionalParams(t *testing.T) {
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
	_, err := client.Assets.List(context.TODO(), imagekit.AssetListParams{
		FileType:    imagekit.AssetListParamsFileTypeAll,
		Limit:       imagekit.Int(1),
		Path:        imagekit.String("path"),
		SearchQuery: imagekit.String("searchQuery"),
		Skip:        imagekit.Int(0),
		Sort:        imagekit.AssetListParamsSortAscName,
		Type:        imagekit.AssetListParamsTypeFile,
	})
	if err != nil {
		var apierr *imagekit.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
