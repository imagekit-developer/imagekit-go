package main

import (
	"context"
	"fmt"
	"log"

	"github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/api/media"
	"github.com/imagekit-developer/imagekit-go/api/uploader"
	"github.com/imagekit-developer/imagekit-go/examples/assets"
)

var ctx = context.Background()

func main() {
	var err error
	ik, err := imagekit.New()

	if err != nil {
		log.Fatal(err)
	}

	var api = ik.Media
	var files []uploader.UploadResult

	for _, f := range []string{"data/nature.jpg", "data/nature.jpg"} {
		files = append(files, assets.UploadFile(ik, f))
	}

	// Get all files
	filesResp, err := api.Files(ctx, media.FilesParam{Path: "/"})
	log.Println(string(filesResp.Body()))

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range filesResp.Data {
		fmt.Println("files resp ", file.FileId, file.VersionInfo)
	}

	// Delete a file
	delResp, err := api.DeleteFile(ctx, files[0].FileId)
	log.Println(delResp, err)

	// Delete multiple files
	delBulkResp, err := api.DeleteBulkFiles(ctx, media.FileIdsParam{
		FileIds: []string{files[1].FileId},
	})
	log.Println(delBulkResp, err)
}
