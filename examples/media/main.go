package main

import (
	"context"
	"fmt"
	"log"

	"github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/url"
)

var ctx = context.Background()

func main() {
	var err error
	ik, err := imagekit.New()

	if err != nil {
		log.Fatal(err)
	}

	//	uri := "https://ik.imagekit.io/dk1m7xkgi/test1_FDy44QvOp.gif"
	path := "test1_FDy44QvOp.gif"

	s, err := ik.Url(url.UrlParam{
		Path:          path,
		Signed:        true,
		ExpireSeconds: 3600 * 24,
	})

	fmt.Println(s, err)
}
