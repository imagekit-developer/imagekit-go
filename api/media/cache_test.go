package media

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMedia_PurgeCache(t *testing.T) {
	var err error

	var param = PurgeCacheParam{
		Url: "https://ik.imagekit.io/dk1m7xkgi/200046731_6088567407851959_2178338512936890186_n_ygW1IEmmc1.jpg?ik-sdk-version=javascript-1.4.3&updatedAt=1654873963613",
	}

	rs := PurgeCacheResponse{
		Data: RequestId{
			RequestId: "xxx",
		},
	}
	respBody, _ := json.Marshal(&rs.Data)

	handler := getHandler(201, string(respBody))
	ts := httptest.NewServer(handler)
	defer ts.Close()

	mediaApi.Config.API.Prefix = ts.URL + "/"
	response, err := mediaApi.PurgeCache(ctx, param)

	if err != nil {
		t.Error(err)
	}
	if !cmp.Equal(response.Data.RequestId, "xxx") {
		t.Error(response.Data)
	}
}

func TestMedia_PurgeCacheStatus(t *testing.T) {
	reqId := "62a4b"
	var rs = PurgeCacheStatusResponse{
		Data: PurgeCacheStatus{
			Status: "Pending",
		},
	}

	respBody, _ := json.Marshal(&rs.Data)

	handler := getHandler(200, string(respBody))
	ts := httptest.NewServer(handler)
	defer ts.Close()

	mediaApi.Config.API.Prefix = ts.URL + "/"
	response, err := mediaApi.PurgeCacheStatus(ctx, reqId)

	if err != nil {
		t.Error(err)
	}
	if !cmp.Equal(response.Data, rs.Data) {
		t.Error(response.Data)
	}
}
