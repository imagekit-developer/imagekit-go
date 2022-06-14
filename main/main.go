package main

import (
	"context"
	"errors"
	"log"

	"github.com/dhaval070/imagekit-go"
	"github.com/dhaval070/imagekit-go/api"
	"github.com/dhaval070/imagekit-go/api/media"
	"github.com/dhaval070/imagekit-go/api/metadata"
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
func main() {
	//getall() return

	var err error

	resp, err := ik.Metadata.CustomFields(ctx, false)
	log.Println(resp.Data, resp.ResponseMetaData)

	if err != nil {
		log.Println("got error")
	}
	log.Println(err)

	//fileId := "62a2fa9121a9dc2869c8bcb0"
	//	resp, err := ik.Media.AssetVersions(ctx, media.AssetVersionsParam{
	//		FileId: fileId,
	//	})

	//resp, err := ik.Media.DeleteAssetVersion(ctx, fileId, "62a2fa9121a9dc2869c8bcb0")
	//resp, err := ik.Media.DeleteAssetVersion(ctx, fileId, "62a309984a5eb5d7b1f09cab")

}
