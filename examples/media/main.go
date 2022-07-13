package main

import (
	"context"
	"fmt"
	"log"

	"github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/api/media"
	"github.com/imagekit-developer/imagekit-go/api/uploader"
	"github.com/imagekit-developer/imagekit-go/examples/assets"
	"github.com/imagekit-developer/imagekit-go/extension"
)

var ctx = context.Background()

func uploadFile(ik *imagekit.ImageKit, path string) uploader.UploadResult {
	var err error

	file, err := assets.Fs.Open(path)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	resp, err := ik.Uploader.Upload(ctx, file, uploader.UploadParam{
		FileName: "test1.gif",
	})

	fmt.Println(resp, err)
	return resp.Data
}

func main() {
	var err error
	ik, err := imagekit.New()

	if err != nil {
		log.Fatal(err)
	}

	var api = ik.Media
	var files []uploader.UploadResult

	for _, f := range []string{"data/beauty_of_nature_12.jpg", "data/image.jpg"} {
		files = append(files, uploadFile(ik, f))
	}
	log.Println(files)

	// Get all files
	filesResp, err := api.Files(ctx, media.FilesParam{})

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range filesResp.Data {
		fmt.Println(file.FileId, file.Name)
	}

	file := filesResp.Data[0]

	// Get file details
	detailResp, err := api.FileById(ctx, file.FileId)
	log.Println(detailResp, err)

	// Get all versions of a file
	versionResp, err := api.FileVersions(ctx, media.FileVersionsParam{FileId: file.FileId})
	log.Println(versionResp, err)

	// Get specific version of a file
	versionDetail, err := api.FileVersions(ctx, media.FileVersionsParam{
		FileId:    file.FileId,
		VersionId: file.VersionInfo["id"],
	})

	log.Println(versionDetail, err)

	// Update file details
	updateResp, err := api.UpdateFile(ctx, file.FileId, media.UpdateFileParam{
		Extensions: []extension.IExtension{
			extension.NewAutoTag(extension.GoogleAutoTag, 50, 10),
			extension.NewRemoveBg(extension.RemoveBgOption{}),
		},
	})
	log.Println(updateResp, err)

	// Add tags
	tagsResp, err := api.UpdateFile(ctx, file.FileId, media.UpdateFileParam{
		Tags: []string{"natural", "mountains", "scene", "day"},
	})
	fmt.Println(tagsResp, err)

	// Remove tags
	remTagResp, err := api.RemoveTags(ctx, media.TagsParam{
		FileIds: []string{file.FileId},
		Tags:    []string{"scene", "day"},
	})
	log.Println(remTagResp, err)

	// Remove AI tags
	remTagResp, err = api.RemoveAITags(ctx, media.AITagsParam{
		FileIds: []string{file.FileId},
		AITags:  []string{updateResp.Data.AITags[0]["name"].(string)},
	})

	log.Println(remTagResp, err)

	// Delete a file
	delResp, err := api.DeleteFile(ctx, files[0].FileId)
	log.Println(delResp, err)
}
