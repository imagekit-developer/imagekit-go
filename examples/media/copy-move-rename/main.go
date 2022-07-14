package main

import (
	"context"
	"log"

	"github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/api/media"
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

	file := assets.UploadFile(ik, "data/nature.jpg")
	file1 := assets.UploadFile(ik, "data/image1.jpg")

	// Copy file
	resp, err := api.CopyFile(ctx, media.CopyFileParam{
		SourcePath:      file.FilePath,
		DestinationPath: "/target/",
	})
	log.Println(resp, err)

	// Move file
	resp, err = api.MoveFile(ctx, media.MoveFileParam{
		SourcePath:      file.FilePath,
		DestinationPath: "/newpath/",
	})
	log.Println(resp, err)

	// Rename file
	renameResp, err := api.RenameFile(ctx, media.RenameFileParam{
		FilePath:    file1.FilePath,
		NewFileName: "sample.jpg",
		PurgeCache:  true, // optionally purge cache
	})
	log.Println(renameResp, err)
}
