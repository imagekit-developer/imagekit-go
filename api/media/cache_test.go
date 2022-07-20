package media

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	iktest "github.com/imagekit-developer/imagekit-go/test"
)

func TestMedia_PurgeCache(t *testing.T) {
	var param = PurgeCacheParam{
		Url: "https://ik.imagekit.io/dk1m7xkgi/200046731_6088567407851959_2178338512936890186_n_ygW1IEmmc1.jpg?ik-sdk-version=javascript-1.4.3&updatedAt=1654873963613",
	}

	rs := PurgeCacheResponse{
		Data: RequestId{
			RequestId: "xxx",
		},
	}
	respBody, _ := json.Marshal(&rs.Data)

	var cases = map[string]struct {
		param     PurgeCacheParam
		result    string
		expectErr bool
	}{
		"purge-cache": {
			param,
			string(respBody),
			false,
		},
		"empty-url": {
			PurgeCacheParam{},
			"{}",
			true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			httpTest := iktest.NewHttp(t)

			ts := httptest.NewServer(httpTest.Handler(201, tc.result))
			defer ts.Close()

			mediaApi.Config.API.Prefix = ts.URL + "/"
			response, err := mediaApi.PurgeCache(ctx, param)

			if err != nil {
				t.Error(err)
			}

			response, err = mediaApi.PurgeCache(ctx, tc.param)

			if tc.expectErr && err == nil {
				t.Error("expected error")
			}

			if !tc.expectErr && err != nil {
				t.Error("Unexpected error")
			}

			if !tc.expectErr && !cmp.Equal(response.Data.RequestId, "xxx") {
				t.Error(response.Data)
			}

			if name != "empty-url" {
				httpTest.Test("/files/purge", "POST", tc.param)
			}
		})
	}

	errServer := iktest.NewErrorServer(t)
	mediaApi.Config.API.Prefix = errServer.Url() + "/"

	errServer.TestErrors(func() error {
		_, err := mediaApi.PurgeCache(ctx, param)
		return err
	})
}

func TestMedia_PurgeCacheStatus(t *testing.T) {
	reqId := "62a4b"
	var rs = PurgeCacheStatusResponse{
		Data: PurgeCacheStatus{
			Status: "Pending",
		},
	}

	respBody, _ := json.Marshal(&rs.Data)

	httpTest := iktest.NewHttp(t)

	ts := httptest.NewServer(httpTest.Handler(200, string(respBody)))
	defer ts.Close()

	mediaApi.Config.API.Prefix = ts.URL + "/"
	response, err := mediaApi.PurgeCacheStatus(ctx, reqId)

	if err != nil {
		t.Error(err)
	}
	if !cmp.Equal(response.Data, rs.Data) {
		t.Error(response.Data)
	}
	httpTest.Test("/files/purge/"+reqId, "GET", nil)

	response, err = mediaApi.PurgeCacheStatus(ctx, "")
	if err == nil {
		t.Error("expected error")
	}

	errServer := iktest.NewErrorServer(t)
	mediaApi.Config.API.Prefix = errServer.Url() + "/"

	errServer.TestErrors(func() error {
		_, err := mediaApi.PurgeCacheStatus(ctx, reqId)
		return err
	})
}
