package main

import (
	"context"
	"log"

	"github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/api/metadata"
)

var ctx = context.Background()

func main() {
	var err error
	ik, err := imagekit.New()

	if err != nil {
		log.Fatal(err)
	}

	api := ik.Metadata

	// Create custom field
	fieldResp, err := api.CreateCustomField(ctx, metadata.CreateFieldParam{
		Name:  "color",
		Label: "Color",
		Schema: metadata.Schema{
			Type:          "SingleSelect",
			SelectOptions: []string{"blue", "red", "green"},
		},
	})

	log.Println(fieldResp, err)
	log.Println(fieldResp.Data.Id)

	// Get custom fields
	resp, err := api.CustomFields(ctx, true)
	log.Println(resp, err)

	// Update custom field
	// Possible to update label or schema only
	updateResp, err := api.UpdateCustomField(ctx, fieldResp.Data.Id, metadata.UpdateCustomFieldParam{
		Label: "Shade",
	})

	log.Println(updateResp, err)

	// Delete custom field
	delResp, err := api.DeleteCustomField(ctx, fieldResp.Data.Id)
	log.Println(delResp, err)
}
