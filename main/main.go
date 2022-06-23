package main

import (
	"context"
	"errors"
	"log"

	"github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/api"
	"github.com/imagekit-developer/imagekit-go/api/media"
	"github.com/imagekit-developer/imagekit-go/api/metadata"
	"github.com/imagekit-developer/imagekit-go/api/uploader"
	"github.com/imagekit-developer/imagekit-go/extension"
	"github.com/imagekit-developer/imagekit-go/logger"
)

var ctx = context.Background()
var ik, _ = imagekit.New()

func getall() {
	resp, _ := ik.Media.Assets(ctx, media.AssetsParam{})
	for _, obj := range resp.Data {
		log.Println(obj.FileId, obj.FilePath, obj.Url)
	}
}

func createField() {
	resp, err := ik.Metadata.CreateCustomField(ctx, metadata.CreateFieldParam{
		Name:  "speed",
		Label: "Speed",
		Schema: metadata.Schema{
			Type:         "Number",
			DefaultValue: 100,
			MinValue:     1,
			MaxValue:     120,
		},
	})

	log.Println(resp.ResponseMetaData)
	if err != nil {
		if errors.Is(err, api.ErrBadRequest) {
			log.Println("bad request")
		}
		log.Println("got error")
		log.Fatal(err)
	}
	log.Println(err)

}

func updateFile() {
	fileId := "62a35f6b168aaf1405abf71b"

	resp, err := ik.Media.UpdateAsset(ctx, fileId, media.UpdateAssetParam{
		Extensions: []extension.IExtension{
			extension.NewAutoTag(extension.AwsAutoTag, 0, 10),
			extension.NewRemoveBg(extension.RemoveBgOption{
				BgColor: "green",
			}),
		},
	})

	log.Println(resp.ResponseMetaData)
	if err != nil {
		log.Println("got error", err)
	}
}

func main() {
	var err error
	ik.Logger.SetLevel(logger.DEBUG)
	// getall() return

	filePath := "/home/dhaval/Pictures/mobile/IMG_20170402_101533.jpg"
	resp, err := ik.Uploader.Upload(ctx, filePath, uploader.UploadParam{
		Extensions: []extension.IExtension{
			extension.NewAutoTag(extension.GoogleAutoTag, 0, 10),
			extension.NewRemoveBg(extension.RemoveBgOption{
				AddShadow: true,
			}),
		},
	})

	log.Println(resp.ResponseMetaData)
	if err != nil {
		log.Println("got error", err)
	}

}
