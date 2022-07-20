package media

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/imagekit-developer/imagekit-go/api/extension"
	iktest "github.com/imagekit-developer/imagekit-go/test"
)

var ctx = context.Background()

var respBody = `[{"fileId":"6283b04dc82abf6294aee010","name":"beauty_of_nature_12_6S7aNLP3-.jpg","filePath":"/beauty_of_nature_12_6S7aNLP3-.jpg","Tags":null,"AITags":null,"versionInfo":{"id":"6283b04dc82abf6294aee010","name":"Version 2"},"isPrivateFile":false,"customCoordinates":null,"url":"https://ik.imagekit.io/dk1m7xkgi/beauty_of_nature_12_6S7aNLP3-.jpg","thumbnail":"https://ik.imagekit.io/dk1m7xkgi/tr:n-ik_ml_thumbnail/beauty_of_nature_12_6S7aNLP3-.jpg","fileType":"image","mime":"image/png","height":133,"Width":200,"size":26509,"hasAlpha":true,"customMetadata":{"price":10},"embeddedMetadata":{"DateCreated":"2022-06-07T15:20:32.104Z","DateTimeCreated":"2022-06-07T15:20:32.105Z","ImageHeight":133,"ImageWidth":200},"createdAt":"2022-05-17T14:25:17.543Z","updatedAt":"2022-06-07T15:20:32.107Z"}]`

var singleFileResp string
var missingFileIdsBody = `
			{
				"message": "The requested file(s) does not exist.",
				"missingFileIds": [
					"yyy"
				]
			}`

var partialSuccessBody = `
	{
		"successfullyUpdatedFileIds": [
			"xxx"
		],
		"errors": [
			{
				"fileId": "yyy",
				"error": "Error in removing tags"
			}
		]
	}`
var assetsArr []File
var asset File
var mediaApi *API
var testExtenstions = []extension.IExtension{
	&extension.AutoTag{
		Name:          extension.GoogleAutoTag,
		MinConfidence: 50,
		MaxTags:       10,
	},
	extension.NewRemoveBg(extension.RemoveBgOption{
		AddShadow:        true,
		SemiTransparency: true,
		BgColor:          "#000000",
		BgImageUrl:       "http://test/test.jpg",
	}),
}
var customMetadata = map[string]any{
	"brand": "nike",
	"size":  10,
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

	singleFileResp = respBody[1 : len(respBody)-1]
	err = json.Unmarshal([]byte(singleFileResp), &asset)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
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

func TestMedia_Files(t *testing.T) {
	var expected = assetsArr
	var cases = map[string]struct {
		params FilesParam
		result string
	}{
		"default": {
			params: FilesParam{},
			result: "/files",
		},
		"with-params": {
			params: FilesParam{
				Type:        ListFile,
				Sort:        AscName,
				Path:        "/test",
				SearchQuery: `createdAt > "7d" AND name: "file-name"`,
				FileType:    Image,
				Tags:        "tag1,tag2",
				Limit:       100,
				Skip:        10,
			},
			result: "/files?fileType=image&limit=100&path=%2Ftest&searchQuery=createdAt+%3E+%227d%22+AND+name%3A+%22file-name%22&skip=10&sort=ASC_NAME&tags=tag1%2Ctag2&type=file",
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			httpTest := iktest.NewHttp(t)

			ts := httptest.NewServer(httpTest.Handler(200, string(respBody)))
			defer ts.Close()

			mediaApi.Config.API.Prefix = ts.URL + "/"

			resp, err := mediaApi.Files(ctx, tc.params)

			if err != nil {
				t.Error(err)
			}

			if !cmp.Equal(resp.Data, expected) {
				t.Errorf("\n%v\n%v\n", resp.Data, expected)
			}

			httpTest.Test(tc.result, "GET", nil)
		})
	}

	errServer := iktest.NewErrorServer(t)
	mediaApi.Config.API.Prefix = errServer.Url() + "/"
	errServer.TestErrors(func() error {
		_, err := mediaApi.Files(ctx, FilesParam{})
		return err
	})
}

func TestMedia_FileById(t *testing.T) {
	var expected = asset
	var mockBody = respBody[1 : len(respBody)-1]

	var cases = map[string]struct {
		fileId     string
		url        string
		result     File
		body       string
		statusCode int
		shouldFail bool
	}{
		"get asset successfully": {
			fileId:     "123",
			url:        "/files/123/details",
			body:       mockBody,
			result:     expected,
			statusCode: 200,
			shouldFail: false,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			httpTest := iktest.NewHttp(t)

			ts := httptest.NewServer(httpTest.Handler(tc.statusCode, tc.body))
			defer ts.Close()

			mediaApi.Config.API.Prefix = ts.URL + "/"

			resp, err := mediaApi.FileById(ctx, tc.fileId)

			if tc.shouldFail && err == nil {
				t.Error("expected error")
			}

			if !tc.shouldFail && err != nil {
				t.Error(err)
			}

			if !cmp.Equal(resp.Data, tc.result) {
				t.Errorf("\n%v\n%v\n", resp.Data, expected)
			}

			httpTest.Test(tc.url, "GET", nil)
		})
	}
	errServer := iktest.NewErrorServer(t)
	mediaApi.Config.API.Prefix = errServer.Url() + "/"

	errServer.TestErrors(func() error {
		_, err := mediaApi.FileById(ctx, "111")
		return err
	})
}

func TestMedia_FileVersions(t *testing.T) {
	var cases = map[string]struct {
		fileId     string
		versionId  string
		body       string
		statusCode int
		shouldFail bool
	}{
		"all versions": {
			fileId:     "6283b04dc82abf6294aee010",
			versionId:  "v123",
			body:       singleFileResp,
			statusCode: 200,
			shouldFail: false,
		},
		"invalid": {
			fileId:     "",
			versionId:  "",
			body:       singleFileResp,
			statusCode: 200,
			shouldFail: true,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			var expectedUrl = "/files/" + tc.fileId + "/versions"

			if tc.versionId != "" {
				expectedUrl = expectedUrl + "/" + tc.versionId
			}

			httpTest := iktest.NewHttp(t)

			ts := httptest.NewServer(httpTest.Handler(tc.statusCode, string(tc.body)))
			defer ts.Close()

			mediaApi.Config.API.Prefix = ts.URL + "/"

			params := FileVersionsParam{
				FileId:    tc.fileId,
				VersionId: tc.versionId,
			}
			_, err := mediaApi.FileVersions(ctx, params)

			if tc.shouldFail && err == nil {
				t.Error("expected error")
			}

			if !tc.shouldFail && err != nil {
				t.Error(err)
			}

			if !tc.shouldFail {
				httpTest.Test(expectedUrl, "GET", nil)
			}
		})
	}
	errServer := iktest.NewErrorServer(t)
	mediaApi.Config.API.Prefix = errServer.Url() + "/"

	errServer.TestErrors(func() error {
		_, err := mediaApi.FileVersions(ctx, FileVersionsParam{FileId: "111", VersionId: "v1"})
		return err
	})
}

func TestMedia_UpdateFile(t *testing.T) {
	var expected = asset
	var mockBody = respBody[1 : len(respBody)-1]

	var cases = map[string]struct {
		result     *File
		fileId     string
		body       string
		params     UpdateFileParam
		statusCode int
		shouldFail bool
	}{
		"update asset": {
			result:     &expected,
			fileId:     "file_id",
			body:       mockBody,
			statusCode: 200,
			shouldFail: false,
			params: UpdateFileParam{
				RemoveAITags:      []string{"one", "two"},
				WebhookUrl:        "http://example.com/hook",
				Tags:              []string{"abc", "def"},
				CustomCoordinates: "12,11,22,22",
				Extensions:        testExtenstions,
				CustomMetadata:    customMetadata,
			},
		},
		"missing-file-id": {
			result:     &expected,
			fileId:     "",
			body:       mockBody,
			statusCode: 200,
			shouldFail: true,
			params: UpdateFileParam{
				RemoveAITags:      []string{"one", "two"},
				WebhookUrl:        "http://example.com/hook",
				Tags:              []string{"abc", "def"},
				CustomCoordinates: "12,11,22,22",
				Extensions:        testExtenstions,
				CustomMetadata:    customMetadata,
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			httpTest := iktest.NewHttp(t)
			ts := httptest.NewServer(httpTest.Handler(200, string(tc.body)))
			defer ts.Close()

			mediaApi.Config.API.Prefix = ts.URL + "/"

			response, err := mediaApi.UpdateFile(ctx, tc.fileId, tc.params)

			var expectedUrl = "/files/" + tc.fileId + "/details"

			if tc.shouldFail == false {
				httpTest.Test(expectedUrl, "PATCH", tc.params)
			}

			if tc.shouldFail == true && err == nil {
				t.Error("expected err")
			}

			if tc.shouldFail == false && err != nil {
				t.Error("err not nil" + err.Error())
			}

			if !tc.shouldFail && !cmp.Equal(tc.result, &response.Data) {
				t.Errorf("unexpected response %v\n%v", tc.result, response.Data)
			}

		})
	}
	errServer := iktest.NewErrorServer(t)
	mediaApi.Config.API.Prefix = errServer.Url() + "/"

	errServer.TestErrors(func() error {
		_, err := mediaApi.UpdateFile(ctx, "111", UpdateFileParam{})
		return err
	})
}

func TestMedia_AddTags(t *testing.T) {
	var ids = []string{"xxx", "yyy"}
	var tags = []string{"tag1", "tag2"}
	var resp = UpdatedIds{
		FileIds: ids,
	}

	respBody, _ := json.Marshal(&resp)

	params := TagsParam{
		FileIds: ids,
		Tags:    tags,
	}

	cases := map[string]struct {
		params     TagsParam
		body       string
		statusCode int
	}{
		"add tags": {
			body:       string(respBody),
			statusCode: 200,
		},
		"partial success": {
			body:       partialSuccessBody,
			statusCode: 207,
		},
		"file id not found": {
			body:       missingFileIdsBody,
			statusCode: 404,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			httpTest := iktest.NewHttp(t)

			ts := httptest.NewServer(httpTest.Handler(tc.statusCode, tc.body))
			defer ts.Close()

			mediaApi.Config.API.Prefix = ts.URL + "/"

			response, err := mediaApi.AddTags(ctx, params)

			var url = "/files/addTags"

			switch tc.statusCode {
			case 200:
				if err != nil {
					t.Error("unexpected error", err)
				} else {
					httpTest.Test(url, "POST", params)
					if !cmp.Equal(response.Data, resp) {
						t.Errorf("%v\n%v", response.Data, resp)
					}
				}
				break
			case 207:
				var errPartial *ErrorPartialSuccess
				if !errors.As(err, &errPartial) {
					log.Println(err)
					t.Error("error is not type ErrorPartialSuccess")
				}
				break
			case 404:
				var e *ErrorMissingFileIds
				if !errors.As(err, &e) {
					log.Println(e)
					t.Error("error is not type ErrorMissingFileIds")
				}
				break
			}
		})
	}

	errServer := iktest.NewErrorServer(t)
	mediaApi.Config.API.Prefix = errServer.Url() + "/"
	errServer.TestErrors(func() error {
		_, err := mediaApi.AddTags(ctx, TagsParam{})
		return err
	})
}

func TestMedia_RemoveTags(t *testing.T) {
	var ids = []string{"xxx", "yyy"}
	var tags = []string{"tag1", "tag2"}
	var resp = UpdatedIds{
		FileIds: ids,
	}
	params := TagsParam{
		FileIds: ids,
		Tags:    tags,
	}

	respBody, _ := json.Marshal(&resp)

	var cases = map[string]struct {
		params     TagsParam
		body       string
		statusCode int
	}{
		"remove tags": {
			body:       string(respBody),
			statusCode: 200,
		},
		"partial success": {
			body:       partialSuccessBody,
			statusCode: 207,
		},
		"missing file ids": {
			body:       missingFileIdsBody,
			statusCode: 404,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			httpTest := iktest.NewHttp(t)

			ts := httptest.NewServer(httpTest.Handler(tc.statusCode, tc.body))
			defer ts.Close()

			mediaApi.Config.API.Prefix = ts.URL + "/"

			response, err := mediaApi.RemoveTags(ctx, params)
			var url = "/files/removeTags"
			switch tc.statusCode {
			case 200:
				if err != nil {
					t.Error("unexpected error", err)
				} else {
					httpTest.Test(url, "POST", params)
					if !cmp.Equal(response.Data, resp) {
						t.Errorf("%v\n%v", response.Data, resp)
					}
				}
				break
			case 207:
				var errPartial *ErrorPartialSuccess
				if !errors.As(err, &errPartial) {
					log.Println(err)
					t.Error("error is not type ErrorPartialSuccess")
				}
				break
			case 404:
				var e *ErrorMissingFileIds
				if !errors.As(err, &e) {
					log.Println(e)
					t.Error("error is not type ErrorMissingFileIds")
				}
				break
			}
		})
	}
	errServer := iktest.NewErrorServer(t)
	mediaApi.Config.API.Prefix = errServer.Url() + "/"

	errServer.TestErrors(func() error {
		_, err := mediaApi.RemoveTags(ctx, params)
		return err
	})
}

func TestMedia_RemoveAITags(t *testing.T) {
	var ids = []string{"xxx", "yyy"}
	var tags = []string{"tag1", "tag2"}
	var resp = UpdatedIds{
		FileIds: ids,
	}

	respBody, _ := json.Marshal(&resp)
	params := AITagsParam{
		FileIds: ids,
		AITags:  tags,
	}

	var cases = map[string]struct {
		body       string
		statusCode int
	}{
		"remove tags": {
			body:       string(respBody),
			statusCode: 200,
		},
		"partial success": {
			body:       partialSuccessBody,
			statusCode: 207,
		},
		"missing file ids": {
			body:       missingFileIdsBody,
			statusCode: 404,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			httpTest := iktest.NewHttp(t)

			ts := httptest.NewServer(httpTest.Handler(tc.statusCode, tc.body))
			defer ts.Close()

			mediaApi.Config.API.Prefix = ts.URL + "/"

			response, err := mediaApi.RemoveAITags(ctx, params)

			if tc.statusCode == 200 {
				if err != nil {
					t.Error(err)
				} else if !cmp.Equal(response.Data, resp) {
					t.Errorf("%v\n%v", response.Data, resp)
				}
			}

			httpTest.Test("/files/removeAITags", "POST", params)
		})
	}
	errServer := iktest.NewErrorServer(t)
	mediaApi.Config.API.Prefix = errServer.Url() + "/"

	errServer.TestErrors(func() error {
		_, err := mediaApi.RemoveAITags(ctx, params)
		return err
	})
}

func TestMedia_DeleteFile(t *testing.T) {
	var err error

	httpTest := iktest.NewHttp(t)

	ts := httptest.NewServer(httpTest.Handler(200, "1"))
	defer ts.Close()

	mediaApi.Config.API.Prefix = ts.URL + "/"
	_, err = mediaApi.DeleteFile(ctx, "file_id")

	if err != nil {
		t.Error(err)
	}

	httpTest.Test("/files/file_id", "DELETE", nil)

	_, err = mediaApi.DeleteFile(ctx, "")

	if err == nil {
		t.Error("expected error")
	}

	errServer := iktest.NewErrorServer(t)
	mediaApi.Config.API.Prefix = errServer.Url() + "/"

	errServer.TestErrors(func() error {
		_, err = mediaApi.DeleteFile(ctx, "file_id")
		return err
	})
}

func TestMedia_DeleteFileVersion(t *testing.T) {
	var err error

	httpTest := iktest.NewHttp(t)

	ts := httptest.NewServer(httpTest.Handler(204, "1"))
	defer ts.Close()

	mediaApi.Config.API.Prefix = ts.URL + "/"
	_, err = mediaApi.DeleteFileVersion(ctx, "file_id", "v2")

	if err != nil {
		t.Error(err)
	}

	url := "/files/file_id/versions/v2"

	httpTest.Test(url, "DELETE", nil)

	_, err = mediaApi.DeleteFileVersion(ctx, "", "v2")

	if err == nil {
		t.Error("expected error")
	}

	_, err = mediaApi.DeleteFileVersion(ctx, "file_id", "")

	if err == nil {
		t.Error("expected error")
	}

	errServer := iktest.NewErrorServer(t)
	mediaApi.Config.API.Prefix = errServer.Url() + "/"

	errServer.TestErrors(func() error {
		_, err = mediaApi.DeleteFileVersion(ctx, "file_id", "v2")
		return err
	})
}

func TestMedia_DeleteBulkFiles(t *testing.T) {
	var err error
	var param = FileIdsParam{
		FileIds: []string{
			"file_id1", "file_id2",
		},
	}
	var respBody = `{"successfullyDeletedFileIds":["file_id1","file_id2"]}`

	httpTest := iktest.NewHttp(t)

	ts := httptest.NewServer(httpTest.Handler(200, string(respBody)))
	defer ts.Close()

	mediaApi.Config.API.Prefix = ts.URL + "/"

	resp, err := mediaApi.DeleteBulkFiles(ctx, param)

	if err != nil {
		t.Error(err)
	}

	log.Println("delete bulk assets test: ", resp.Data.Errors)

	if !cmp.Equal(resp.Data.FileIds, param.FileIds) {
		t.Errorf("expected: %v, got: %v", param.FileIds, resp.Data.FileIds)
	}
	httpTest.Test("/files/batch/deleteByFileIds", "POST", param)

	resp, err = mediaApi.DeleteBulkFiles(ctx, FileIdsParam{})

	if err == nil {
		t.Error("expected error")
	}

	errServer := iktest.NewErrorServer(t)
	mediaApi.Config.API.Prefix = errServer.Url() + "/"

	errServer.TestErrors(func() error {
		_, err := mediaApi.DeleteBulkFiles(ctx, param)
		return err
	})
}

func TestMedia_CopyFile(t *testing.T) {
	var err error
	var param = CopyFileParam{
		SourcePath:          "/file.jpg",
		DestinationPath:     "/natural/file.jpg",
		IncludeFileVersions: true,
	}

	httpTest := iktest.NewHttp(t)

	ts := httptest.NewServer(httpTest.Handler(204, ""))
	defer ts.Close()

	mediaApi.Config.API.Prefix = ts.URL + "/"

	_, err = mediaApi.CopyFile(ctx, param)
	if err != nil {
		t.Error(err)
	}
	httpTest.Test("/files/copy", "POST", param)

	_, err = mediaApi.CopyFile(ctx, CopyFileParam{})

	if err == nil {
		t.Error(err)
	}
	errServer := iktest.NewErrorServer(t)
	mediaApi.Config.API.Prefix = errServer.Url() + "/"

	errServer.TestErrors(func() error {
		_, err = mediaApi.CopyFile(ctx, param)
		return err
	})
}

func TestMedia_MoveFile(t *testing.T) {
	var err error
	var param = MoveFileParam{
		SourcePath:      "/file.jpg",
		DestinationPath: "/natural/",
	}

	httpTest := iktest.NewHttp(t)

	ts := httptest.NewServer(httpTest.Handler(204, ""))
	defer ts.Close()

	mediaApi.Config.API.Prefix = ts.URL + "/"

	_, err = mediaApi.MoveFile(ctx, param)
	if err != nil {
		t.Error(err)
	}

	httpTest.Test("/files/move", "POST", param)

	_, err = mediaApi.MoveFile(ctx, MoveFileParam{})
	if err == nil {
		t.Error("expected error")
	}

	errServer := iktest.NewErrorServer(t)
	mediaApi.Config.API.Prefix = errServer.Url() + "/"

	errServer.TestErrors(func() error {
		_, err = mediaApi.MoveFile(ctx, param)
		return err
	})
}

func TestMedia_RenameFile(t *testing.T) {
	var cases = map[string]struct {
		param      RenameFileParam
		body       string
		statusCode int
	}{
		"rename asset": {
			param: RenameFileParam{
				FilePath:    "/some/file.jpg",
				NewFileName: "/default.jpg",
				PurgeCache:  true,
			},
			body:       `{"purgeRequestId":"123"}`,
			statusCode: 200,
		},
		"without purge": {
			param: RenameFileParam{
				FilePath:    "/some/file.jpg",
				NewFileName: "/default.jpg",
			},
			body:       `{}`,
			statusCode: 200,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			httpTest := iktest.NewHttp(t)

			ts := httptest.NewServer(httpTest.Handler(tc.statusCode, tc.body))
			defer ts.Close()

			mediaApi.Config.API.Prefix = ts.URL + "/"

			resp, err := mediaApi.RenameFile(ctx, tc.param)
			if err != nil {
				t.Error(err)
			}

			if tc.param.PurgeCache == true && resp.Data.RequestId != "123" {
				t.Error("unexpected request id returned")
			}

			httpTest.Test("/files/rename", "PUT", tc.param)
		})
	}
	httpTest := iktest.NewHttp(t)
	ts := httptest.NewServer(httpTest.Handler(200, ""))
	defer ts.Close()

	mediaApi.Config.API.Prefix = ts.URL + "/"

	_, err := mediaApi.RenameFile(ctx, RenameFileParam{})
	if err == nil {
		t.Error(err)
	}
	errServer := iktest.NewErrorServer(t)
	mediaApi.Config.API.Prefix = errServer.Url() + "/"

	errServer.TestErrors(func() error {
		_, err := mediaApi.RenameFile(ctx, RenameFileParam{
			FilePath:    "/some/file.jpg",
			NewFileName: "/default.jpg",
		})
		return err
	})
}
func TestMedia_RestoreVersion(t *testing.T) {
	var err error

	var param = FileVersionsParam{
		FileId:    "file_id",
		VersionId: "v1",
	}

	httpTest := iktest.NewHttp(t)

	ts := httptest.NewServer(httpTest.Handler(200, singleFileResp))
	defer ts.Close()

	mediaApi.Config.API.Prefix = ts.URL + "/"

	resp, err := mediaApi.RestoreVersion(ctx, param)
	if err != nil {
		t.Error(err)
	}

	if !cmp.Equal(resp.Data, asset) {
		t.Error("unexpected response")
	}

	expectedUrl := fmt.Sprintf("/files/%s/versions/%s/restore",
		param.FileId, param.VersionId)

	httpTest.Test(expectedUrl, "DELETE", param)

	resp, err = mediaApi.RestoreVersion(ctx, FileVersionsParam{})
	if err == nil {
		t.Error("expected error")
	}

	errServer := iktest.NewErrorServer(t)
	mediaApi.Config.API.Prefix = errServer.Url() + "/"

	errServer.TestErrors(func() error {
		_, err := mediaApi.RestoreVersion(ctx, param)
		return err
	})
}

func TestMedia_BulkJobStatus(t *testing.T) {
	var err error
	var mockBody = `{"jobId":"job_id","type":"MOVE_FOLDER","status":"Completed"}`
	var res = JobStatusResponse{
		Data: JobStatus{"job_id", "MOVE_FOLDER", "Completed"},
	}
	_ = json.Unmarshal([]byte(mockBody), &res)
	var jobId = "job_id"
	httpTest := iktest.NewHttp(t)

	ts := httptest.NewServer(httpTest.Handler(200, mockBody))
	defer ts.Close()

	mediaApi.Config.API.Prefix = ts.URL + "/"

	resp, err := mediaApi.BulkJobStatus(ctx, jobId)
	if err != nil {
		t.Error(err)
	}

	if !cmp.Equal(resp.Data, res.Data) {
		t.Error("unexpected response")
	}

	httpTest.Test("/bulkJobs/"+jobId, "GET", nil)

	resp, err = mediaApi.BulkJobStatus(ctx, "")
	if err == nil {
		t.Error("expected error")
	}

	errServer := iktest.NewErrorServer(t)
	mediaApi.Config.API.Prefix = errServer.Url() + "/"

	errServer.TestErrors(func() error {
		_, err := mediaApi.BulkJobStatus(ctx, jobId)
		return err
	})
}
