package main

import (
	"context"
	"log"

	"github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/examples/assets"
)

var ctx = context.Background()

func main() {
	var err error
	ik, err := imagekit.New()

	if err != nil {
		log.Fatal(err)
	}

	api := ik.Metadata

	file := assets.UploadFile(ik, "data/nature.jpg")

	// metadata from url
	resp, err := api.FromUrl(ctx, file.Url)

	log.Println(resp, err)
	log.Println(resp.Data.Height, resp.Data.Width)

	// metadata from fileId
	resp, err = api.FromFile(ctx, file.FileId)
	log.Println(resp, err)
	log.Println(resp.Data.Height, resp.Data.Width)

}
