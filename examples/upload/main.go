package main

import (
	"context"
	"fmt"
	"log"

	"github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/api/uploader"
	"github.com/imagekit-developer/imagekit-go/examples/assets"
)

var ctx = context.Background()

const Base64Image = "data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7"
const url = "https://images.pexels.com/photos/247676/pexels-photo-247676.jpeg?auto=compress&cs=tinysrgb&dpr=3&h=750&w=1260"

// Upload base64 encoded file
func uploadBase64(ik *imagekit.ImageKit) {
	resp, err := ik.Uploader.Upload(ctx, Base64Image, uploader.UploadParam{
		FileName: "test.gif",
	})

	fmt.Println(resp, err)
}

// Upload using a reader
func uploadReader(ik *imagekit.ImageKit) {
	var err error

	file, err := assets.Fs.Open("data/nature.jpg")
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}

	resp, err := ik.Uploader.Upload(ctx, file, uploader.UploadParam{
		FileName: "test1.gif",
	})

	fmt.Println(resp, err)
}

// Upload using a url
func uploadUrl(ik *imagekit.ImageKit) {
	var err error

	resp, err := ik.Uploader.Upload(ctx, url, uploader.UploadParam{
		FileName: "test2.gif",
	})

	fmt.Println(resp, err)
}

func main() {
	ik, err := imagekit.New()

	if err != nil {
		log.Fatal(err)
	}

	uploadBase64(ik)
	uploadReader(ik)
	uploadUrl(ik)
}
