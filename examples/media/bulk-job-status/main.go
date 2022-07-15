package main

import (
	"context"
	"log"

	"github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/api/media"
)

var ctx = context.Background()

func main() {
	var err error
	ik, err := imagekit.New()

	if err != nil {
		log.Fatal(err)
	}

	// Create folder
	resp, err := ik.Media.CreateFolder(ctx, media.CreateFolderParam{
		FolderName:       "new",
		ParentFolderPath: "/",
	})
	log.Println(resp, err)

	// Move Folder
	moveResp, err := ik.Media.MoveFolder(ctx, media.MoveFolderParam{
		SourceFolderPath: "/new",
		DestinationPath:  "/nature",
	})

	// Get job status
	jobStatus, err := ik.Media.BulkJobStatus(ctx, moveResp.Data.JobId)
	log.Println(jobStatus.Data, err)
}
