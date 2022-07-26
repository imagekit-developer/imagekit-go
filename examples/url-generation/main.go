package main

import (
	"fmt"
	"log"

	"github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/api"
	ikurl "github.com/imagekit-developer/imagekit-go/url"
)

func main() {
	ik, err := imagekit.New()

	if err != nil {
		log.Fatal(err)
	}

	url, err := ik.Url(ikurl.UrlParam{
		Path: "/default-image.jpg",
		Transformations: []ikurl.Transformation{
			{
				Height: 300,
				Width:  300,

				Overlay: &ikurl.Overlay{
					Image:       "test2_5r9tRj4L4.gif",
					Width:       api.Int(100),
					X:           api.Int(0),
					ImageBorder: "10_CDDC39",
				},
			},
		},
	})

	fmt.Println(url, err)
}
