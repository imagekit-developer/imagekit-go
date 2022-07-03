package uploader

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

	"github.com/google/go-cmp/cmp"
	"github.com/imagekit-developer/imagekit-go/api"
	iktest "github.com/imagekit-developer/imagekit-go/test"
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
}

func newUploader(url string) (*API, error) {
	uploader, err := NewFromConfiguration(iktest.Cfg)

	if err != nil {
		return nil, err
	}

	uploader.Config.API.UploadPrefix = url
	return uploader, nil

}

func TestUploader(t *testing.T) {
	resultJson, err := json.Marshal(file)
	if err != nil {
		t.Error(err)
	}

	reader, err := os.Open(iktest.ImageFilePath)

	if err != nil {
		t.Fatal(err)
	}

	var cases = map[string]struct {
		file   interface{}
		resp   string
		param  UploadParam
		result *UploadResult
	}{
		"base64file": {
			file: iktest.Base64Image,
			resp: string(resultJson),
			param: UploadParam{
				FileName: "new-york-cityscape-buildings_A4zxKJbrL.jpg",
			},
			result: file,
		},
		"io-reader": {
			file: reader,
			resp: string(resultJson),
			param: UploadParam{
				FileName: "new-york-cityscape-buildings_A4zxKJbrL.jpg",
			},
			result: file,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			submittedFile := []byte{}

			handler := func(w http.ResponseWriter, r *http.Request) {
				reader, _ := r.MultipartReader()

				for {
					part, err := reader.NextPart()
					if err == io.EOF {
						break
					}

					if part.FormName() == "file" {
						log.Println("reading file")
						submittedFile, _ = io.ReadAll(part)
					}
				}
				fmt.Fprintln(w, test.resp)
			}

			ts := httptest.NewServer(http.HandlerFunc(handler))
			defer ts.Close()

			uploader, err := newUploader(ts.URL + "/")
			if err != nil {
				t.Fatal(err)
			}

			resp, err := uploader.Upload(ctx, test.file, test.param)

			if err != nil {
				t.Error(err)
			}

			if resp.Response.StatusCode != 200 {
				t.Error("status code:", resp.Response.StatusCode)
			}

			switch fv := test.file.(type) {
			case string:
				if api.IsLocalFilePath(test.file) {
					fp, _ := os.Open(fv)
					expectedFile, _ := io.ReadAll(fp)

					if !cmp.Equal(expectedFile, submittedFile) {
						log.Println(submittedFile)

						t.Error("expected file not submitted")
					}
				} else {
					if test.file != string(submittedFile) {
						t.Error("unexpected file submitted")
					}
				}
			default:
				reader, _ := os.Open(iktest.ImageFilePath)
				expected, _ := io.ReadAll(reader)
				if !cmp.Equal(expected, submittedFile) {
					t.Error("unexpected file submitted")
				}
			}

			if !cmp.Equal(resp.Data, *test.result) {
				t.Errorf("\n%v\n%v\n", resp.Data, *test.result)
			}
		})
	}
}
