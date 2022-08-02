package main

import (
	"fmt"
	"log"

	"github.com/imagekit-developer/imagekit-go"
	ikurl "github.com/imagekit-developer/imagekit-go/url"
)

func main() {
	ik, err := imagekit.New()

	if err != nil {
		log.Fatal(err)
	}

	url, err := ik.Url(ikurl.UrlParam{
		Path: "/default-image.jpg",
		Transformations: []map[string]any{
			{
				"height":             300,
				"width":              300,
				"overlayImage":       "test2_5r9tRj4L4.gif",
				"overlayWidth":       100,
				"x":                  0,
				"overlayImageBorder": "10_CDDC39",
			},
		},
	})

	fmt.Println(url, err)
}
