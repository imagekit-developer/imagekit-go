package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/stainless-sdks/imagekit-go"
	"github.com/stainless-sdks/imagekit-go/option"
	"github.com/stainless-sdks/imagekit-go/packages/param"
)

// Example demonstrates file upload from public URL
func main() {
	// Initialize the ImageKit client
	client := imagekit.NewClient(
		option.WithPrivateKey(os.Getenv("IMAGEKIT_PRIVATE_KEY")),
	)

	// Public image URL to upload from
	imageURL := "https://picsum.photos/800/600"

	// Fetch the image from URL
	resp, err := http.Get(imageURL)
	if err != nil {
		log.Fatal("Error fetching image from URL:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to fetch image: HTTP %d", resp.StatusCode)
	}

	// Upload from URL stream
	ctx := context.Background()
	response, err := client.Files.Upload(ctx, imagekit.FileUploadParams{
		File:     resp.Body,
		FileName: "remote-upload.jpg",
		Folder:   param.NewOpt("go-sdk-testing"),
		Tags:     []string{"remote", "url", "picsum"},
		CustomMetadata: map[string]any{
			"rating": 15,
		},
		ResponseFields: []string{
			"tags",
			"customMetadata",
			"embeddedMetadata",
		},
	})

	if err != nil {
		log.Fatal("Error uploading file:", err)
	}

	fmt.Printf("URL Upload successful!\n")
	fmt.Printf("File ID: %s\n", response.FileID)
	fmt.Printf("Name: %s\n", response.Name)
	fmt.Printf("URL: %s\n", response.URL)
	fmt.Printf("File Path: %s\n", response.FilePath)
	fmt.Printf("Size: %.0f bytes\n", response.Size)
	fmt.Printf("Tags: %v\n", response.Tags)
	if response.CustomMetadata != nil {
		fmt.Printf("Custom Metadata: %+v\n", response.CustomMetadata)
	}
}
