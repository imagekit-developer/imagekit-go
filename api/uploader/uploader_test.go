package uploader

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/imagekit-developer/imagekit-go/api"
	"github.com/imagekit-developer/imagekit-go/api/extension"
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

func TestUploader(t *testing.T) {
	resultJson, err := json.Marshal(file)
	if err != nil {
		t.Error(err)
	}

	reader, err := os.Open(iktest.ImageFilePath)

	if err != nil {
		t.Fatal(err)
	}

	var extensions = []extension.IExtension{
		extension.NewAutoTag(extension.GoogleAutoTag, 50, 10),
		extension.NewRemoveBg(extension.RemoveBgOption{
			AddShadow:        true,
			SemiTransparency: true,
			BgColor:          "#553333",
			BgImageUrl:       "http://test/test.jpg",
		}),
	}

	param := UploadParam{
		FileName:                "new-york-cityscape-buildings_A4zxKJbrL.jpg",
		UseUniqueFileName:       api.Bool(false),
		Tags:                    "tag_1,tag_2",
		Folder:                  "/natural",
		IsPrivateFile:           api.Bool(false),
		CustomCoordinates:       "11,100,400,500",
		ResponseFields:          "tags,customCoordinates,isPrivateFile",
		Extensions:              extensions,
		WebhookUrl:              "http://test/test",
		OverwriteFile:           api.Bool(true),
		OverwriteAITags:         api.Bool(true),
		OverwriteTags:           api.Bool(true),
		OverwriteCustomMetadata: api.Bool(true),
		CustomMetadata: map[string]any{
			"Cost": 100,
		},
	}

	var formStr = `{"customCoordinates":"11,100,400,500","customMetadata":"{\"Cost\":100}","extensions":"[{\"name\":\"google-auto-tagging\",\"minConfidence\":50,\"maxTags\":10},{\"name\":\"remove-bg\",\"options\":{\"add_shadow\":true,\"semitransparency\":true,\"bg_color\":\"#553333\",\"bg_image_url\":\"http://test/test.jpg\"}}]","file":"data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7","fileName":"new-york-cityscape-buildings_A4zxKJbrL.jpg","folder":"/natural","isPrivateFile":"false","overwriteAITags":"true","overwriteCustomMetadata":"true","overwriteFile":"true","overwriteTags":"true","responseFields":"tags,customCoordinates,isPrivateFile","tags":"tag_1,tag_2","useUniqueFileName":"false","webhookUrl":"http://test/test"}`

	var readerForm = `{"customCoordinates":"11,100,400,500","customMetadata":"{\"Cost\":100}","extensions":"[{\"name\":\"google-auto-tagging\",\"minConfidence\":50,\"maxTags\":10},{\"name\":\"remove-bg\",\"options\":{\"add_shadow\":true,\"semitransparency\":true,\"bg_color\":\"#553333\",\"bg_image_url\":\"http://test/test.jpg\"}}]","fileName":"new-york-cityscape-buildings_A4zxKJbrL.jpg","folder":"/natural","isPrivateFile":"false","overwriteAITags":"true","overwriteCustomMetadata":"true","overwriteFile":"true","overwriteTags":"true","responseFields":"tags,customCoordinates,isPrivateFile","tags":"tag_1,tag_2","useUniqueFileName":"false","webhookUrl":"http://test/test"}`

	var cases = map[string]struct {
		file       any
		resp       string
		param      UploadParam
		result     *UploadResult
		shouldFail bool
		formStr    string
	}{
		"base64file": {
			file:    iktest.Base64Image,
			resp:    string(resultJson),
			param:   param,
			result:  file,
			formStr: formStr,
		},
		"io-reader": {
			file:    reader,
			resp:    string(resultJson),
			param:   param,
			result:  file,
			formStr: readerForm,
		},
		"missing-filename": {
			file:       reader,
			resp:       string(resultJson),
			param:      UploadParam{},
			result:     file,
			shouldFail: true,
		},
		"invalid file param": {
			file:       []int{},
			resp:       "",
			param:      UploadParam{},
			result:     nil,
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
				if err == nil {
					t.Error("err is nil")
				} else if err.Error() != "Upload: Filename is required" {
					t.Error("wrong error message", err.Error())
				}
				return
			}

			_, params, err := mime.ParseMediaType(httpTest.Req.Header.Get("Content-Type"))

			if err != nil {
				t.Fatal(err)
			}

			mr := multipart.NewReader(bytes.NewReader(httpTest.Body), params["boundary"])

			form, err := mr.ReadForm(1024 * 2)
			var f = map[string]string{}

			for k, v := range form.Value {
				f[k] = v[0]
			}
			str, err := json.Marshal(f)

			if !cmp.Equal(string(str), test.formStr) {
				t.Error("invalid form values")
				log.Println(test.formStr)
				log.Println(string(str))
			}
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

	errServer := iktest.NewErrorServer(t)
	uploader, err := newUploader(errServer.Url() + "/")
	errServer.TestErrors(func() error {
		_, err := uploader.Upload(ctx, reader, param)
		return err
	})
}

func Test_postFile(t *testing.T) {
	uploader, err := newUploader("/")
	if err != nil {
		t.Fatal(err)
	}

	_, err = uploader.postFile(ctx, 5, url.Values{})

	if err == nil {
		t.Error("expected error")
	}

}
