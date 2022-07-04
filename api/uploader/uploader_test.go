package uploader

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	iktest "github.com/imagekit-developer/imagekit-go/test"
)

var ctx = context.Background()
var header = http.Header{}
var ImageFileData []byte

func init() {
	header.Add("Content-Type", "text/plain; charset=utf-8")
	header.Add("Content-Length", "600")
	header.Add("Date", "Sat, 04 Jun 2022 06:50:21 GMT")

	file, _ := os.Open(iktest.ImageFilePath)
	ImageFileData, _ = io.ReadAll(file)
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

/**
REVIEW-COMMENT

Multiple test scenearos missing. For example
1. missing requried File and FileName param.
2. passing all parameters that SDK supports and asserting correct conversion from arrays/map to string. For example, 
Tags passsed as array should be converted to comma-seperated stirngs. Extensions passed as array of map should be converted to string. 
CustomMetadata passed as map should be converted. All this needs to be asserted.

*/
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
		file       interface{}
		resp       string
		param      UploadParam
		result     *UploadResult
		shouldFail bool
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
		"missing-filename": {
			file:       reader,
			resp:       string(resultJson),
			param:      UploadParam{},
			result:     file,
			shouldFail: true,
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			httpTest := iktest.NewHttp(t)

			ts := httptest.NewServer(httpTest.Handler(200, string(resultJson)))
			defer ts.Close()

			uploader, err := newUploader(ts.URL + "/")
			if err != nil {
				t.Fatal(err)
			}

			resp, err := uploader.Upload(ctx, test.file, test.param)

			if !test.shouldFail && err != nil {
				t.Error(err)
			}

			if test.shouldFail {
				/**
				REVIEW-COMMENT

				Also assert the content of err

				*/
				if err == nil {
					t.Error("err is nil")
				}
				return
			}

			_, params, err := mime.ParseMediaType(httpTest.Req.Header.Get("Content-Type"))

			if err != nil {
				t.Fatal(err)
			}

			mr := multipart.NewReader(bytes.NewReader(httpTest.Body), params["boundary"])

			form, err := mr.ReadForm(1024 * 2)
			if form.Value["fileName"][0] != test.param.FileName {
				t.Error("invalid filename")
			}

			switch test.file.(type) {
			case string:
				if test.file != form.Value["file"][0] {
					t.Error("unexpected file value")
				}
			default:
				file, err := form.File["file"][0].Open()

				if err != nil {
					t.Fatal(err)
				}
				data, _ := io.ReadAll(file)

				if !cmp.Equal(data, ImageFileData) {
					t.Error(name + ": unexpected file submitted")
				}
			}

			if !cmp.Equal(resp.Data, *test.result) {
				t.Errorf("\n%v\n%v\n", resp.Data, *test.result)
			}
		})
	}
}
