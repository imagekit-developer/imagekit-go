package uploader

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	iktest "github.com/dhaval070/imagekit-go/test"
)

var ctx = context.Background()
var File = &UploadResult{
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
}

func TestUploader_UploadLocalPath(t *testing.T) {
	fileUploadResp, err := json.Marshal(File)
	if err != nil {
		t.Error(err)
	}

	api, err := NewFromConfiguration(iktest.Cfg)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, string(fileUploadResp))
	}))
	defer ts.Close()

	api.Config.API.UploadPrefix = ts.URL + "/"

	_, err = api.Upload(ctx, iktest.ImageFilePath, UploadParams{})

	if err != nil {
		t.Error(err)
	}
}
