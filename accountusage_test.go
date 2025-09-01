// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/internal/testutil"
	"github.com/imagekit-developer/imagekit-go/option"
)

func TestAccountUsageGet(t *testing.T) {
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
	_, err := client.Accounts.Usage.Get(context.TODO(), imagekit.AccountUsageGetParams{
		EndDate:   time.Now(),
		StartDate: time.Now(),
	})
	if err != nil {
		var apierr *imagekit.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
