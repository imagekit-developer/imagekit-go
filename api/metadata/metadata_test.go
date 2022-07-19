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

	"github.com/google/go-cmp/cmp"
	iktest "github.com/imagekit-developer/imagekit-go/test"
)

var ctx = context.Background()
var metadataApi *API

func Test_New(t *testing.T) {
	os.Setenv("IMAGEKIT_PRIVATE_KEY", "private_")
	os.Setenv("IMAGEKIT_PUBLIC_KEY", "public_")
	os.Setenv("IMAGEKIT_ENDPOINT_URL", "https://ik.imagekit.io/test/")

	defer os.Unsetenv("IMAGEKIT_PRIVATE_KEY")
	defer os.Unsetenv("IMAGEKIT_PUBLIC_KEY")
	defer os.Unsetenv("IMAGEKIT_ENDPOINT_URL")

	var api any
	api, err := New()

	if err != nil {
		t.Fatal(err)
	}

	if api == nil {
		t.Error("New() returned null")
	}

	if _, ok := api.(*API); !ok {
		t.Error("New() did not return *API")
	}
}

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

	httpTest := iktest.NewHttp(t)
	ts := httptest.NewServer(httpTest.Handler(200, string(respBody)))
	defer ts.Close()

	metadataApi.Config.API.Prefix = ts.URL + "/"

	resp, err := metadataApi.FromFile(ctx, "file_id")

	if err != nil {
		t.Error(err)
	}

	if !cmp.Equal(resp.Data, *respObj) {
		t.Errorf("\n%v\n%v\n", resp.Data, *respObj)
	}

	httpTest.Test("/files/file_id/metadata", "GET", nil)

	errServer := iktest.NewErrorServer(t)
	metadataApi.Config.API.Prefix = errServer.Url() + "/"

	errServer.TestErrors(func() error {
		_, err := metadataApi.FromFile(ctx, "3325344545345")
		return err
	})

	resp, err = metadataApi.FromFile(ctx, "")

	if err == nil {
		t.Error("expected error")
	}
}

func TestMetadata_FromUrl(t *testing.T) {
	var respBody = `{"height":801,"width":597,"size":59718,"format":"jpg","hasColorProfile":true,"quality":0,"density":72,"hasTransparency":false,"exif":{},"pHash":"85d07f1fe4ae8be2"}`

	var err error
	var respObj = &Metadata{}

	_, err = metadataApi.FromUrl(ctx, "")
	if err == nil {
		t.Error("expected error")
	}

	if err == nil {
		t.Error("expected error")
	}
	if err = json.Unmarshal([]byte(respBody), respObj); err != nil {
		t.Error(err)
	}

	httpTest := iktest.NewHttp(t)

	ts := httptest.NewServer(httpTest.Handler(200, string(respBody)))
	defer ts.Close()

	metadataApi.Config.API.Prefix = ts.URL + "/"

	resp, err := metadataApi.FromUrl(ctx, "https://ik.imagekit.io/xk1m7xkgi/default-image.jpg")

	if err != nil {
		t.Error(err)
	}

	if !cmp.Equal(resp.Data, *respObj) {
		t.Errorf("\n%v\n%v\n", resp.Data, *respObj)
	}

	httpTest.Test("/metadata?url=https%3A%2F%2Fik.imagekit.io%2Fxk1m7xkgi%2Fdefault-image.jpg", "GET", nil)

	errServer := iktest.NewErrorServer(t)
	metadataApi.Config.API.Prefix = errServer.Url() + "/"

	errServer.TestErrors(func() error {
		_, err := metadataApi.FromUrl(ctx, "https://ik.imagekit.io/xk1m7xkgi/default-image.jpg")
		return err
	})
}

func TestMetadata_CreateCustomField(t *testing.T) {
	var respBody = `{"id":"62a8966b663ef736f841fe28","name":"speed","label":"Speed","schema":{"type":"Number","defaultValue":100,"minValue":1,"maxValue":120}}`

	var err error
	var expected = &CustomField{}

	if err = json.Unmarshal([]byte(respBody), expected); err != nil {
		t.Error(err)
	}

	httpTest := iktest.NewHttp(t)

	ts := httptest.NewServer(httpTest.Handler(201, respBody))
	defer ts.Close()

	metadataApi.Config.API.Prefix = ts.URL + "/"

	param := CreateFieldParam{
		Name:  "speed",
		Label: "Speed",
		Schema: Schema{
			Type:         "Number",
			DefaultValue: 100,
			MinValue:     1,
			MaxValue:     120,
		},
	}
	resp, err := metadataApi.CreateCustomField(ctx, param)

	if err != nil {
		t.Error(err)
	}

	if !cmp.Equal(resp.Data, *expected) {
		t.Errorf("\n%v\n%v\n", resp.Data, *expected)
	}

	httpTest.Test("/customMetadataFields", "POST", param)

	errServer := iktest.NewErrorServer(t)
	metadataApi.Config.API.Prefix = errServer.Url() + "/"

	errServer.TestErrors(func() error {
		_, err := metadataApi.CreateCustomField(ctx, param)
		return err
	})
}

func TestMetadata_CustomFields(t *testing.T) {
	var respBody = `[{"id":"629f6b437eb0fe6f1b66d864","name":"price","label":"Price","schema":{"type":"Number","isValueRequired":false,"minValue":1,"maxValue":1000}},{"id":"629f6b6d7eb0fe344f66e1b6","name":"country","label":"Country","schema":{"type":"SingleSelect","isValueRequired":false,"selectOptions":["USA","Canada"]}},{"id":"62a8764d663ef721e93f4ea9","name":"clearance","label":"Clearance","schema":{"type":"MultiSelect","selectOptions":["one","two"]}},{"id":"62a876b1663ef7728f3f5348","name":"mileage","label":"Mileage","schema":{"type":"Number"}},{"id":"62a8966b663ef736f841fe28","name":"speed","label":"Speed","schema":{"type":"Number","defaultValue":100,"minValue":1,"maxValue":120}}]`

	var err error
	var expected = []CustomField{}

	if err = json.Unmarshal([]byte(respBody), &expected); err != nil {
		t.Error(err)
	}

	httpTest := iktest.NewHttp(t)

	ts := httptest.NewServer(httpTest.Handler(200, respBody))
	defer ts.Close()

	metadataApi.Config.API.Prefix = ts.URL + "/"
	resp, err := metadataApi.CustomFields(ctx, false)

	if err != nil {
		t.Error(err)
	}
	if !cmp.Equal(resp.Data, expected) {
		t.Errorf("%v\n%v", resp.Data, expected)
	}

	httpTest.Test("/customMetadataFields?includeDeleted=false", "GET", nil)

	errServer := iktest.NewErrorServer(t)
	metadataApi.Config.API.Prefix = errServer.Url() + "/"

	errServer.TestErrors(func() error {
		_, err := metadataApi.CustomFields(ctx, false)
		return err
	})
}

func TestMetadata_UpdateCustomField(t *testing.T) {
	var respBody = `{"id":"629f6b437eb0fe6f1b66d864","name":"price","label":"Cost","schema":{"type":"Number"}}`

	var err error
	var expected = CustomField{}

	if err = json.Unmarshal([]byte(respBody), &expected); err != nil {
		t.Error(err)
	}

	httpTest := iktest.NewHttp(t)
	ts := httptest.NewServer(httpTest.Handler(200, respBody))
	defer ts.Close()

	metadataApi.Config.API.Prefix = ts.URL + "/"

	param := UpdateCustomFieldParam{
		Label: "Cost",
	}

	resp, err := metadataApi.UpdateCustomField(
		ctx,
		"file_id",
		param,
	)

	if err != nil {
		t.Error(err)
	}
	if !cmp.Equal(resp.Data, expected) {
		t.Errorf("%v\n%v", resp.Data, expected)
	}

	httpTest.Test("/customMetadataFields/file_id", "PATCH", param)

	errServer := iktest.NewErrorServer(t)
	metadataApi.Config.API.Prefix = errServer.Url() + "/"

	errServer.TestErrors(func() error {
		_, err := metadataApi.UpdateCustomField(
			ctx,
			"629f6b437eb0fe6f1b66d864",
			UpdateCustomFieldParam{
				Label: "Cost",
			},
		)
		return err
	})
}

func TestMetadata_DeleteCustomField(t *testing.T) {
	var respBody = ``
	var err error

	if _, err := metadataApi.DeleteCustomField(ctx, ""); err == nil {
		t.Error("expected error")
	}

	httpTest := iktest.NewHttp(t)

	ts := httptest.NewServer(httpTest.Handler(204, respBody))
	defer ts.Close()

	metadataApi.Config.API.Prefix = ts.URL + "/"

	_, err = metadataApi.DeleteCustomField(ctx, "file_id")
	if err != nil {
		log.Println("got error")
		t.Error(err)
	}

	httpTest.Test("/customMetadataFields/file_id", "DELETE", nil)

	errServer := iktest.NewErrorServer(t)
	metadataApi.Config.API.Prefix = errServer.Url() + "/"

	errServer.TestErrors(func() error {
		_, err = metadataApi.DeleteCustomField(ctx, "62a8966b663ef736f841fe28")
		return err
	})
}
