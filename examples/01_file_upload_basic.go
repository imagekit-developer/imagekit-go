package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/stainless-sdks/imagekit-go"
	"github.com/stainless-sdks/imagekit-go/option"
	"github.com/stainless-sdks/imagekit-go/packages/param"
)

// Example demonstrates basic file upload from local file
func main() {
	// Initialize the ImageKit client
	client := imagekit.NewClient(
		option.WithPrivateKey(os.Getenv("IMAGEKIT_PRIVATE_KEY")),
	)

	// Open local file
	file, err := os.Open("sample-image.png")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	// Upload file with basic parameters
	ctx := context.Background()
	response, err := client.Files.Upload(ctx, imagekit.FileUploadParams{
		File:     file,
		FileName: "sample-upload.png",
		Folder:   param.NewOpt("go-sdk-testing"),
	})

	if err != nil {
		log.Fatal("Error uploading file:", err)
	}

	fmt.Printf("Upload successful!\n")
	fmt.Printf("File ID: %s\n", response.FileID)
	fmt.Printf("Name: %s\n", response.Name)
	fmt.Printf("URL: %s\n", response.URL)
	fmt.Printf("File Path: %s\n", response.FilePath)
	fmt.Printf("Size: %.0f bytes\n", response.Size)
	fmt.Printf("File Type: %s\n", response.FileType)
	fmt.Printf("Width: %.0f\n", response.Width)
	fmt.Printf("Height: %.0f\n", response.Height)
}
