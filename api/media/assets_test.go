package media

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	iktest "github.com/dhaval070/imagekit-go/test"
	"github.com/google/go-cmp/cmp"
)

var ctx = context.Background()

var respBody = `[{"fileId":"6283b04dc82abf6294aee010","name":"beauty_of_nature_12_6S7aNLP3-.jpg","filePath":"/beauty_of_nature_12_6S7aNLP3-.jpg","Tags":null,"AITags":null,"versionInfo":{"id":"6283b04dc82abf6294aee010","name":"Version 2"},"isPrivateFile":false,"customCoordinates":null,"url":"https://ik.imagekit.io/dk1m7xkgi/beauty_of_nature_12_6S7aNLP3-.jpg","thumbnail":"https://ik.imagekit.io/dk1m7xkgi/tr:n-ik_ml_thumbnail/beauty_of_nature_12_6S7aNLP3-.jpg","fileType":"image","mime":"image/png","height":133,"Width":200,"size":26509,"hasAlpha":true,"customMetadata":{"price":10},"embeddedMetadata":{"DateCreated":"2022-06-07T15:20:32.104Z","DateTimeCreated":"2022-06-07T15:20:32.105Z","ImageHeight":133,"ImageWidth":200},"createdAt":"2022-05-17T14:25:17.543Z","updatedAt":"2022-06-07T15:20:32.107Z"}]`

func TestAssets(t *testing.T) {
	var err error
	assets, err := NewFromConfiguration(iktest.Cfg)

	if err != nil {
		t.Error(err)
	}

	var expected []Asset
	err = json.Unmarshal([]byte(respBody), &expected)

	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, respBody)
	}

	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	assets.Config.API.Prefix = ts.URL + "/"

	resp, err := assets.Assets(ctx, AssetsParam{})

	if err != nil {
		t.Error(err)
	}

	if !cmp.Equal(resp.Data, expected) {
		t.Errorf("\n%v\n%v\n", resp.Data, expected)
	}
}
