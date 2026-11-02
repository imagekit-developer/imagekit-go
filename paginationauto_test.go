// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit_test

import (
	"context"
	"os"
	"testing"

	"github.com/imagekit-developer/imagekit-go/v2"
	"github.com/imagekit-developer/imagekit-go/v2/internal/testutil"
	"github.com/imagekit-developer/imagekit-go/v2/option"
)

func TestAutoPagination(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	iter := client.Tags.ListAutoPaging(context.TODO(), imagekit.TagListParams{})
	// The mock server isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		tag := iter.Current()
		t.Logf("%+v\n", tag.AssetCount)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
