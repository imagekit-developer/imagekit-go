package media

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	iktest "github.com/imagekit-developer/imagekit-go/test"
)

/**
REVIEW-COMMENT

Negative test case missing of invalid/missing params, including non 2xx response from backend.
*/
func TestMedia_CreateFolder(t *testing.T) {
	var err error

	var param = CreateFolderParam{
		FolderName:       "testing",
		ParentFolderPath: "/",
	}

	httpTest := iktest.NewHttp(t)

	ts := httptest.NewServer(httpTest.Handler(201, "{}"))
	defer ts.Close()

	mediaApi.Config.API.Prefix = ts.URL + "/"

	_, err = mediaApi.CreateFolder(ctx, param)

	if err != nil {
		t.Error(err)
	}
	httpTest.Test("/folder", "POST", param)
}

/**
REVIEW-COMMENT

Negative test case missing of invalid/missing params, including non 2xx response from backend.
*/
func TestMedia_DeleteFolder(t *testing.T) {
	var err error

	var param = DeleteFolderParam{
		FolderPath: "testing",
	}

	httpTest := iktest.NewHttp(t)

	ts := httptest.NewServer(httpTest.Handler(204, "{}"))
	defer ts.Close()

	mediaApi.Config.API.Prefix = ts.URL + "/"
	_, err = mediaApi.DeleteFolder(ctx, param)

	if err != nil {
		t.Error(err)
	}

	httpTest.Test("/folder", "DELETE", nil)
}

/**
REVIEW-COMMENT

Negative test case missing of invalid/missing params, including non 2xx response from backend.
*/
func TestMedia_MoveFolder(t *testing.T) {
	var err error

	var param = MoveFolderParam{
		SourceFolderPath: "/src",
		DestinationPath:  "dest",
	}

	httpTest := iktest.NewHttp(t)

	ts := httptest.NewServer(httpTest.Handler(200, `{"jobId":"xxx"}`))
	defer ts.Close()

	mediaApi.Config.API.Prefix = ts.URL + "/"
	response, err := mediaApi.MoveFolder(ctx, param)

	if err != nil {
		t.Error(err)
	}
	if !cmp.Equal(response.Data.JobId, "xxx") {
		t.Error(response.Data)
	}
	httpTest.Test("/bulkJobs/moveFolder", "POST", param)
}

/**
REVIEW-COMMENT

Negative test case missing of invalid/missing params, including non 2xx response from backend.
*/
func TestMedia_CopyFolder(t *testing.T) {
	var err error

	var param = CopyFolderParam{
		SourceFolderPath: "/src",
		DestinationPath:  "dest",
	}

	rs := FolderResponse{
		Data: JobIdResponse{
			JobId: "xxx",
		},
	}
	respBody, _ := json.Marshal(&rs.Data)

	httpTest := iktest.NewHttp(t)

	ts := httptest.NewServer(httpTest.Handler(200, string(respBody)))
	defer ts.Close()

	mediaApi.Config.API.Prefix = ts.URL + "/"
	response, err := mediaApi.CopyFolder(ctx, param)

	if err != nil {
		t.Error(err)
	}
	if !cmp.Equal(response.Data.JobId, "xxx") {
		t.Error(response.Data)
	}
	httpTest.Test("/bulkJobs/copyFolder", "POST", param)
}
