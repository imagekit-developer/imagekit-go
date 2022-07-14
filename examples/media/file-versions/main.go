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

	var api = ik.Media
	// replace fileId and versionId with real values
	fileId := ""
	versionId := ""

	// Get all versions
	versionsResp, err := api.FileVersions(ctx, media.FileVersionsParam{
		FileId: fileId,
	})

	log.Println(versionsResp, err)

	// Get file version
	versionsResp, err = api.FileVersions(ctx, media.FileVersionsParam{
		FileId:    fileId,
		VersionId: versionId,
	})
	log.Println(versionsResp, err)

	// Delete file version
	resp, err := api.DeleteFileVersion(ctx, fileId, versionId)
	log.Println(resp, err)

	// Delete bulk files
	delBulkResp, err := api.DeleteBulkFiles(ctx, media.FileIdsParam{
		FileIds: []string{"x", "y"}, // replace with actual values
	})
	log.Println(delBulkResp, err)

}
