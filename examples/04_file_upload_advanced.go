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

// Example demonstrates advanced file upload with transformations and options
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

	// Upload with advanced options
	ctx := context.Background()
	response, err := client.Files.Upload(ctx, imagekit.FileUploadParams{
		File:                    file,
		FileName:                "advanced-upload.png",
		Folder:                  param.NewOpt("go-sdk-testing"),
		UseUniqueFileName:       param.NewOpt(false),
		IsPrivateFile:           param.NewOpt(true),
		IsPublished:             param.NewOpt(false),
		OverwriteFile:           param.NewOpt(false),
		OverwriteTags:           param.NewOpt(false),
		OverwriteAITags:         param.NewOpt(false),
		OverwriteCustomMetadata: param.NewOpt(false),
		Tags:                    []string{"advanced", "sample", "go-sdk", "demo"},
		CustomCoordinates:       param.NewOpt("10,10,100,100"), // x,y,width,height
		Description:             param.NewOpt("Advanced upload example with transformations"),
		CustomMetadata: map[string]any{
			"rating": 15,
		},
		// Apply transformations during upload
		Transformation: imagekit.FileUploadParamsTransformation{
			Pre: param.NewOpt("w-400,h-300,c-pad_resize,bg-F3F3F3"), // Resize and pad
		},
		// Request additional fields in response
		ResponseFields: []string{
			"tags",
			"customCoordinates",
			"isPrivateFile",
			"embeddedMetadata",
			"isPublished",
			"customMetadata",
			"selectedFieldsSchema",
		},
	})

	if err != nil {
		log.Fatal("Error uploading file:", err)
	}

	fmt.Printf("Advanced Upload successful!\n")
	fmt.Printf("File ID: %s\n", response.FileID)
	fmt.Printf("Name: %s\n", response.Name)
	fmt.Printf("URL: %s\n", response.URL)
	fmt.Printf("File Path: %s\n", response.FilePath)
	fmt.Printf("Size: %.0f bytes\n", response.Size)
	fmt.Printf("File Type: %s\n", response.FileType)
	fmt.Printf("Tags: %v\n", response.Tags)
	fmt.Printf("Custom Coordinates: %s\n", response.CustomCoordinates)
	fmt.Printf("Is Private: %t\n", response.IsPrivateFile)
	fmt.Printf("Is Published: %t\n", response.IsPublished)
	fmt.Printf("Description: %s\n", response.Description)
	if response.CustomMetadata != nil {
		fmt.Printf("Custom Metadata: %+v\n", response.CustomMetadata)
	}
	if response.EmbeddedMetadata != nil {
		fmt.Printf("EXIF Data: Available\n")
	}
	if response.SelectedFieldsSchema != nil {
		fmt.Printf("SelectedFieldsSchema Available\n")
	}
}
