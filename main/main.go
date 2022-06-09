package main

import (
	"context"
	"log"

	"github.com/dhaval070/imagekit-go"
	"github.com/dhaval070/imagekit-go/api/media"
)

var ctx = context.Background()
var ik, _ = imagekit.New()

func getall() {
	resp, _ := ik.Media.Assets(ctx, media.AssetsParam{})
	log.Println(resp.Data)
}

func main() {
	getall()

}
