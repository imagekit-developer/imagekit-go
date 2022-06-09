package media

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

var respBody = `[{"fileId":"6283b04dc82abf6294aee010","name":"beauty_of_nature_12_6S7aNLP3-.jpg","filePath":"/beauty_of_nature_12_6S7aNLP3-.jpg","Tags":null,"AITags":null,"versionInfo":{"id":"6283b04dc82abf6294aee010","name":"Version 2"},"isPrivateFile":false,"customCoordinates":null,"url":"https://ik.imagekit.io/dk1m7xkgi/beauty_of_nature_12_6S7aNLP3-.jpg","thumbnail":"https://ik.imagekit.io/dk1m7xkgi/tr:n-ik_ml_thumbnail/beauty_of_nature_12_6S7aNLP3-.jpg","fileType":"image","mime":"image/png","height":133,"Width":200,"size":26509,"hasAlpha":true,"customMetadata":{"price":10},"embeddedMetadata":{"DateCreated":"2022-06-07T15:20:32.104Z","DateTimeCreated":"2022-06-07T15:20:32.105Z","ImageHeight":133,"ImageWidth":200},"createdAt":"2022-05-17T14:25:17.543Z","updatedAt":"2022-06-07T15:20:32.107Z"}]`

var assetsArr []Asset
var asset Asset
var mediaApi *API

func getHandler(statusCode int, body string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		fmt.Fprintln(w, body)
	}
}

func TestMain(m *testing.M) {
	var err error
	mediaApi, err = NewFromConfiguration(iktest.Cfg)

	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal([]byte(respBody), &assetsArr); err != nil {
		log.Fatal(err)
	}

	var mockBody = respBody[1 : len(respBody)-1]
	err = json.Unmarshal([]byte(mockBody), &asset)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

func TestMedia_Assets(t *testing.T) {
	var err error

	var expected = assetsArr
	err = json.Unmarshal([]byte(respBody), &expected)

	handler := getHandler(200, respBody)

	ts := httptest.NewServer(handler)
	defer ts.Close()

	mediaApi.Config.API.Prefix = ts.URL + "/"

	resp, err := mediaApi.Assets(ctx, AssetsParam{})

	if err != nil {
		t.Error(err)
	}

	if !cmp.Equal(resp.Data, expected) {
		t.Errorf("\n%v\n%v\n", resp.Data, expected)
	}
}

func TestMedia_AssetById(t *testing.T) {
	var expected = asset
	var mockBody = respBody[1 : len(respBody)-1]

	var cases = map[string]struct {
		result     Asset
		body       string
		statusCode int
		shouldFail bool
	}{
		"get asset successfully": {
			body:       mockBody,
			result:     expected,
			statusCode: 200,
			shouldFail: false,
		},
		"check failure": {
			body:       `{"message":"not found"}`,
			result:     Asset{},
			statusCode: 400,
			shouldFail: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			handler := getHandler(tc.statusCode, tc.body)
			ts := httptest.NewServer(handler)
			defer ts.Close()

			mediaApi.Config.API.Prefix = ts.URL + "/"

			resp, err := mediaApi.AssetById(ctx, AssetByIdParam{FileId: expected.FileId})

			if tc.shouldFail && err == nil {
				t.Error("expected error")
			}

			if !tc.shouldFail && err != nil {
				t.Error(err)
			}

			if !cmp.Equal(resp.Data, tc.result) {
				t.Errorf("\n%v\n%v\n", resp.Data, expected)
			}
		})
	}
}

func TestMedia_AssetVersions(t *testing.T) {
	var cases = map[string]struct {
		result     []Asset
		body       string
		statusCode int
		shouldFail bool
	}{
		"all versions": {
			result:     assetsArr,
			body:       respBody,
			statusCode: 200,
			shouldFail: false,
		},
		"should fail": {
			result:     nil,
			body:       `{"message": "not found"}`,
			statusCode: 400,
			shouldFail: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			handler := getHandler(tc.statusCode, tc.body)

			ts := httptest.NewServer(handler)
			defer ts.Close()

			mediaApi.Config.API.Prefix = ts.URL + "/"

			params := AssetVersionsParam{
				FileId: "testxx",
			}
			resp, err := mediaApi.AssetVersions(ctx, params)

			if tc.shouldFail && err == nil {
				t.Error("expected error")
			}

			if !tc.shouldFail && err != nil {
				t.Error(err)
			}

			if !cmp.Equal(resp.Data, tc.result) {
				t.Errorf("\n%v\n%v\n", resp.Data, tc.result)
			}
		})
	}
}
