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

	// Copy folder
	folderResp, err := ik.Media.CopyFolder(ctx, media.CopyFolderParam{
		SourceFolderPath:    "/new",
		DestinationPath:     "/sample",
		IncludeFileVersions: true,
	})
	log.Println(folderResp, err)

	// Move Folder
	folderResp, err = ik.Media.MoveFolder(ctx, media.MoveFolderParam{
		SourceFolderPath: "/new",
		DestinationPath:  "/nature",
	})
	log.Println(folderResp, err)

	// Delete Folder
	resp, err = ik.Media.DeleteFolder(ctx, media.DeleteFolderParam{
		FolderPath: "/nature",
	})
	log.Println(resp, err)
}
