package media

import (
	"context"
	"log"
	"testing"

	iktest "github.com/dhaval070/imagekit-go/test"
)

var ctx = context.Background()

func TestAssets(t *testing.T) {
	assets, err := NewFromConfiguration(iktest.Cfg)

	log.Println(err)
	resp, err := assets.Assets(ctx, AssetsParam{})

	log.Println(resp, err)

}
