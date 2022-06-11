package media

import (
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMedia_CreateFolder(t *testing.T) {
	var err error

	var param = CreateFolderParam{
		FolderName:       "testing",
		ParentFolderPath: "/",
	}

	handler := getHandler(201, "{}")

	ts := httptest.NewServer(handler)
	defer ts.Close()

	mediaApi.Config.API.Prefix = ts.URL + "/"

	_, err = mediaApi.CreateFolder(ctx, param)

	if err != nil {
		t.Error(err)
	}
}

func TestMedia_DeleteFolder(t *testing.T) {
	var err error

	var param = DeleteFolderParam{
		FolderPath: "testing",
	}

	handler := getHandler(204, "{}")

	ts := httptest.NewServer(handler)
	defer ts.Close()

	mediaApi.Config.API.Prefix = ts.URL + "/"
	_, err = mediaApi.DeleteFolder(ctx, param)

	if err != nil {
		t.Error(err)
	}

}

func TestMedia_MoveFolder(t *testing.T) {
	var err error

	var param = MoveFolderParam{
		SourceFolderPath: "/src",
		DestinationPath:  "dest",
	}

	handler := getHandler(200, `{"jobId":"xxx"}`)

	ts := httptest.NewServer(handler)
	defer ts.Close()

	mediaApi.Config.API.Prefix = ts.URL + "/"
	response, err := mediaApi.MoveFolder(ctx, param)

	if err != nil {
		t.Error(err)
	}
	if !cmp.Equal(response.Data.JobId, "xxx") {
		t.Error(response.Data)
	}
}
