package uploader

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dhaval070/imagekit-go/api"
	iktest "github.com/dhaval070/imagekit-go/test"
	"github.com/google/go-cmp/cmp"
)

var ctx = context.Background()
var header = http.Header{}

func init() {
	header.Add("Content-Type", "text/plain; charset=utf-8")
	header.Add("Content-Length", "600")
	header.Add("Date", "Sat, 04 Jun 2022 06:50:21 GMT")
}

var file = &UploadResult{
	FileId:       "6299dcfb124cdf0dffaf2dfd",
	Name:         "new-york-cityscape-buildings_A4zxKJbrL.jpg",
	Url:          "https://ik.imagekit.io/dk1m7xkgi/new-york-cityscape-buildings_A4zxKJbrL.jpg",
	ThumbnailUrl: "https://ik.imagekit.io/dk1m7xkgi/tr:n-ik_ml_thumbnail/new-york-cityscape-buildings_A4zxKJbrL.jpg",
	Height:       1080,
	Width:        1920,
	Size:         1468522,
	FilePath:     "/new-york-cityscape-buildings_A4zxKJbrL.jpg",
	AITags:       nil,
	VersionInfo: map[string]string{
		"id": "6299dcfb124cdf0dffaf2dfd", "name": "Version 1",
	},
	Response: api.Response{
		ResponseMetaData: api.ResponseMetaData{
			StatusCode: 200,
			Status:     "200 OK",
			Header:     header,
		},
	},
}

func TestUploader_UploadLocalPath(t *testing.T) {
	fileUploadResp, err := json.Marshal(file)
	if err != nil {
		t.Error(err)
	}

	uploader, err := NewFromConfiguration(iktest.Cfg)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, string(fileUploadResp))
	}))

	defer ts.Close()

	uploader.Config.API.UploadPrefix = ts.URL + "/"

	resp, err := uploader.Upload(ctx, iktest.ImageFilePath, UploadParams{})

	log.Println("resp: ", resp)

	if err != nil {
		t.Error(err)
	}

	if resp.Response.StatusCode != 200 {
		t.Error("invalid status code")
	}

	if !cmp.Equal(resp, file) {
		t.Error("file upload response invalid")
	}
}
