package media

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
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

func TestMedia_UpdateAsset(t *testing.T) {
	var expected = asset
	var mockBody = respBody[1 : len(respBody)-1]

	var cases = map[string]struct {
		result     *Asset
		fileId     string
		body       string
		params     UpdateAssetParam
		statusCode int
		shouldFail bool
	}{
		"update asset": {
			result:     &expected,
			fileId:     "xxx",
			body:       mockBody,
			statusCode: 200,
			shouldFail: false,
			params: UpdateAssetParam{
				RemoveAITags:      []string{"one", "two"},
				WebhookUrl:        "http://example.com/hook",
				Tags:              []string{"abc", "def"},
				CustomCoordinates: "12,11,22,22",
			},
		},
		"require fileid": {
			fileId:     "",
			body:       "",
			statusCode: 400,
			shouldFail: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			var bodyReceived []byte
			var receivedRequest *http.Request

			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				bodyReceived, _ = io.ReadAll(r.Body)
				receivedRequest = r.Clone(ctx)
				w.WriteHeader(tc.statusCode)
				fmt.Fprintln(w, tc.body)
			})

			ts := httptest.NewServer(handler)
			defer ts.Close()

			mediaApi.Config.API.Prefix = ts.URL + "/"
			expectedBody, err := json.Marshal(tc.params)
			response, err := mediaApi.UpdateAsset(ctx, tc.fileId, tc.params)

			if receivedRequest != nil {
				iktest.JsonRequest(receivedRequest, t)
			}

			if tc.shouldFail == false && !cmp.Equal(bodyReceived, expectedBody) {
				t.Error("incorrect request body")
			}

			if tc.shouldFail == true && err == nil {
				t.Error("expected err")
			}

			if tc.shouldFail == false && err != nil {
				t.Error(err)
			}

			if !tc.shouldFail && !cmp.Equal(tc.result, &response.Data) {
				t.Errorf("unexpected response %v\n%v", tc.result, response.Data)
			}
		})
	}
}

func TestMedia_AddTags(t *testing.T) {
	var err error
	var ids = []string{"xxx", "yyy"}
	var tags = []string{"tag1", "tag2"}
	var resp = UpdatedIds{
		FileIds: ids,
	}

	respBody, _ := json.Marshal(&resp)

	handler := getHandler(200, string(respBody))

	ts := httptest.NewServer(handler)
	defer ts.Close()

	mediaApi.Config.API.Prefix = ts.URL + "/"

	params := TagsParam{
		FileIds: ids,
		Tags:    tags,
	}

	response, err := mediaApi.AddTags(ctx, params)

	if err != nil {
		t.Error(err)
	}

	if !cmp.Equal(response.Data, resp) {
		t.Errorf("%v\n%v", response.Data, resp)
	}
}

func TestMedia_RemoveTags(t *testing.T) {
	var err error
	var ids = []string{"xxx", "yyy"}
	var tags = []string{"tag1", "tag2"}
	var resp = UpdatedIds{
		FileIds: ids,
	}

	respBody, _ := json.Marshal(&resp)

	handler := getHandler(200, string(respBody))

	ts := httptest.NewServer(handler)
	defer ts.Close()

	mediaApi.Config.API.Prefix = ts.URL + "/"

	params := TagsParam{
		FileIds: ids,
		Tags:    tags,
	}

	response, err := mediaApi.RemoveTags(ctx, params)

	if err != nil {
		t.Error(err)
	}

	if !cmp.Equal(response.Data, resp) {
		t.Errorf("%v\n%v", response.Data, resp)
	}
}

func TestMedia_RemoveAITags(t *testing.T) {
	var err error
	var ids = []string{"xxx", "yyy"}
	var tags = []string{"tag1", "tag2"}
	var resp = UpdatedIds{
		FileIds: ids,
	}

	respBody, _ := json.Marshal(&resp)

	handler := getHandler(200, string(respBody))

	ts := httptest.NewServer(handler)
	defer ts.Close()

	mediaApi.Config.API.Prefix = ts.URL + "/"

	params := AITagsParam{
		FileIds: ids,
		AITags:  tags,
	}

	response, err := mediaApi.RemoveAITags(ctx, params)

	if err != nil {
		t.Error(err)
	}

	if !cmp.Equal(response.Data, resp) {
		t.Errorf("%v\n%v", response.Data, resp)
	}
}

func TestMedia_DeleteAsset(t *testing.T) {
	var err error

	handler := getHandler(204, string("1"))

	ts := httptest.NewServer(handler)
	defer ts.Close()

	mediaApi.Config.API.Prefix = ts.URL + "/"
	_, err = mediaApi.DeleteAsset(ctx, "xxx")

	if err != nil {
		t.Error(err)
	}

}
