package main

import (
	"context"
	"log"

	"github.com/imagekit-developer/imagekit-go"
)

var ctx = context.Background()

func main() {
	var err error
	ik, err := imagekit.New()

	if err != nil {
		log.Fatal(err)
	}

	var api = ik.Media

	// replace fileId with real value
	fileId := ""

	// Get file detail by id
	fileResp, err := api.FileById(ctx, fileId)
	log.Println(fileResp, err)

}
