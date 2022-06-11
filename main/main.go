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
	for _, obj := range resp.Data {
		log.Println(obj.FileId, obj.FilePath)
	}
}

func main() {
	var err error
	//getall() return

	resp, err := ik.Media.MoveAsset(ctx, media.MoveAssetParam{
		SourcePath:      "/212106665_794340761448854_2038800919838075402_n_LXpeZDSUZO.jpg",
		DestinationPath: "/natural/",
	})

	if err != nil {
		log.Println("got error")
		log.Fatal(err)
	}
	log.Println(resp)

	//fileId := "62a2fa9121a9dc2869c8bcb0"
	//	resp, err := ik.Media.AssetVersions(ctx, media.AssetVersionsParam{
	//		FileId: fileId,
	//	})

	//resp, err := ik.Media.DeleteAssetVersion(ctx, fileId, "62a2fa9121a9dc2869c8bcb0")
	//resp, err := ik.Media.DeleteAssetVersion(ctx, fileId, "62a309984a5eb5d7b1f09cab")

}
