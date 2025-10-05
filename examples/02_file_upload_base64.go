package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/stainless-sdks/imagekit-go"
	"github.com/stainless-sdks/imagekit-go/option"
	"github.com/stainless-sdks/imagekit-go/packages/param"
)

// Example demonstrates file upload from base64 encoded data
func main() {
	// Initialize the ImageKit client
	client := imagekit.NewClient(
		option.WithPrivateKey(os.Getenv("IMAGEKIT_PRIVATE_KEY")),
	)

	// Read the image file and convert to base64
	imageData, err := os.ReadFile("sample-image.png")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	// Convert to base64 string
	base64String := base64.StdEncoding.EncodeToString(imageData)

	// Decode base64 back to binary data for upload
	decodedData, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		log.Fatal("Error decoding base64:", err)
	}

	// Upload using binary buffer from base64
	ctx := context.Background()
	response, err := client.Files.Upload(ctx, imagekit.FileUploadParams{
		File:     bytes.NewReader(decodedData),
		FileName: "sample-base64-upload.png",
		Folder:   param.NewOpt("go-sdk-testing"), // Optional: organize in folders
		Tags:     []string{"base64", "sample", "test"},
		ResponseFields: []string{
			"tags",
			"customCoordinates",
			"isPrivateFile",
			"embeddedMetadata",
			"customMetadata",
		},
	})

	if err != nil {
		log.Fatal("Error uploading file:", err)
	}

	fmt.Printf("Base64 Upload successful!\n")
	fmt.Printf("File ID: %s\n", response.FileID)
	fmt.Printf("Name: %s\n", response.Name)
	fmt.Printf("URL: %s\n", response.URL)
	fmt.Printf("File Path: %s\n", response.FilePath)
	fmt.Printf("Size: %.0f bytes\n", response.Size)
	fmt.Printf("Tags: %v\n", response.Tags)
	if response.EmbeddedMetadata != nil {
		fmt.Printf("EXIF Data available: Yes\n")
	}
}
