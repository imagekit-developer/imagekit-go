package main

import (
	"context"
	"errors"
	"log"

	"github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/api"
	"github.com/imagekit-developer/imagekit-go/api/metadata"
)

var ctx = context.Background()
var ik, _ = imagekit.New()

func createField() {
	resp, err := ik.Metadata.CreateCustomField(ctx, metadata.CreateFieldParam{
		Name:  "speed",
		Label: "Speed",
		Schema: metadata.Schema{
			Type:         "Number",
			DefaultValue: 100,
			MinValue:     1,
			MaxValue:     120,
		},
	})

	log.Println(resp.ResponseMetaData)
	if err != nil {
		if errors.Is(err, api.ErrBadRequest) {
			log.Println("bad request")
		}
		log.Println("got error")
		log.Fatal(err)
	}
	log.Println(err)

}

func main() {

}
