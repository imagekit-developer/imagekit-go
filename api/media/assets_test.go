package media

import (
	"context"
	"log"
	"testing"

	iktest "github.com/dhaval070/imagekit-go/test"
)

var ctx = context.Background()

func TestAssets(t *testing.T) {
	assetsApi, err := NewFromConfiguration(iktest.Cfg)

	log.Println(err)
	assetsApi.Assets(ctx, AssetsParams{})

}
