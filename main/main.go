package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dhaval070/imagekit-go"
	"github.com/dhaval070/imagekit-go/api/media"
)

var ctx = context.Background()

func main() {
	ik, err := imagekit.New()

	if err != nil {
		log.Fatal(err)
	}

	resp, err := ik.Media.AssetById(ctx, media.AssetByIdParam{FileId: "6283aed9f2f6a2ee87e1d6b5"})
	s, err := json.Marshal(resp.Data)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(s), err)
}
