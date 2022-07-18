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

	for _, f := range []string{"data/nature.jpg", "data/image.jpg", "data/image1.jpg", "data/image2.jpg"} {
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

}
