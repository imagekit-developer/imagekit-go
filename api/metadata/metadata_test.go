package metadata

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	iktest "github.com/dhaval070/imagekit-go/test"
	"github.com/google/go-cmp/cmp"
)

var ctx = context.Background()
var metadataApi *API

func TestMain(m *testing.M) {
	var err error
	metadataApi, err = NewFromConfiguration(iktest.Cfg)

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}
func getHandler(statusCode int, body string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		fmt.Fprintln(w, body)
	}
}

func TestMetadata_FromFile(t *testing.T) {
	var respBody = `{"height":801,"width":597,"size":59718,"format":"jpg","hasColorProfile":true,"quality":0,"density":72,"hasTransparency":false,"exif":{},"pHash":"85d07f1fe4ae8be2"}`

	var err error
	var respObj = &Metadata{}

	if err = json.Unmarshal([]byte(respBody), respObj); err != nil {
		t.Error(err)
	}

	handler := getHandler(200, respBody)
	ts := httptest.NewServer(handler)
	defer ts.Close()

	metadataApi.Config.API.Prefix = ts.URL + "/"

	resp, err := metadataApi.FromAsset(ctx, "3325344545345")

	if err != nil {
		t.Error(err)
	}

	if !cmp.Equal(resp.Data, *respObj) {
		t.Errorf("\n%v\n%v\n", resp.Data, *respObj)
	}
}

func TestMetadata_FromUrl(t *testing.T) {
	var respBody = `{"height":801,"width":597,"size":59718,"format":"jpg","hasColorProfile":true,"quality":0,"density":72,"hasTransparency":false,"exif":{},"pHash":"85d07f1fe4ae8be2"}`

	var err error
	var respObj = &Metadata{}

	if err = json.Unmarshal([]byte(respBody), respObj); err != nil {
		t.Error(err)
	}

	handler := getHandler(200, respBody)
	ts := httptest.NewServer(handler)
	defer ts.Close()

	metadataApi.Config.API.Prefix = ts.URL + "/"

	resp, err := metadataApi.FromUrl(ctx, "https://ik.imagekit.io/xk1m7xkgi/default-image.jpg")

	if err != nil {
		t.Error(err)
	}

	if !cmp.Equal(resp.Data, *respObj) {
		t.Errorf("\n%v\n%v\n", resp.Data, *respObj)
	}
}
