package imagekit

import (
	neturl "net/url"
	"reflect"
	"testing"

	"github.com/dhaval070/imagekit-go/logger"
	"github.com/dhaval070/imagekit-go/url"
	ikurl "github.com/dhaval070/imagekit-go/url"
)

var imgkit = NewFromParams(NewParams{
	PrivateKey:  "private_",
	PublicKey:   "public_",
	EndpointUrl: "https://ik.imagekit.io/test/",
})

func init() {
	imgkit.Logger.SetLevel(logger.DEBUG)
}

func TestUrl(t *testing.T) {
	type tcase struct {
		name   string
		params ikurl.UrlParams
		url    string
	}

	cases := []tcase{
		{
			name: "path",
			params: ikurl.UrlParams{
				Path:           "default-image.jpg",
				EndpointUrl:    "https://imagekit.io/343534/",
				Transformation: "w-100,rt-90",
			},
			url: "https://imagekit.io/343534/tr:w-100,rt-90/default-image.jpg",
		}, {
			name: "signed",
			params: ikurl.UrlParams{
				Path:           "default-image.jpg",
				EndpointUrl:    "https://ik.imagekit.io/test/",
				Transformation: "w-200,rt-90",
				Signed:         true,
				ExpireSeconds:  100,
				UnixTime: func() int64 {
					return 1653775828
				},
			},
			url: "https://ik.imagekit.io/test/tr:w-200,rt-90/default-image.jpg?ik-t=1653775928&ik-s=6c74b32f5efc0cc39ab5c042718b36553d8a1606",
		}, {
			name: "src",
			params: ikurl.UrlParams{
				Src:            "https://imagekit.io/343534/default-image.jpg",
				Transformation: "w-100,rt-90",
			},
			url: "https://imagekit.io/343534/default-image.jpg?tr=w-100,rt-90",
		}, {
			name: "trquery",
			params: ikurl.UrlParams{
				Path:                   "default-image.jpg",
				EndpointUrl:            "https://imagekit.io/343534/",
				Transformation:         "w-100,rt-90",
				TransformationPosition: url.QUERY,
			},
			url: "https://imagekit.io/343534/default-image.jpg?tr=w-100%2crt-90",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			url, err := imgkit.Url(tc.params)
			imgkit.Logger.Debug(url)

			if err != nil {
				t.Errorf(err.Error())
			}

			if !urlsEquals(url, tc.url) {
				t.Errorf("expected url: " + tc.url + ", got: " + url)
			}
		})
	}

}

func urlsEquals(url1 string, url2 string) bool {
	u1, _ := neturl.Parse(url1)
	u2, _ := neturl.Parse(url2)

	q1 := u1.Query()
	q2 := u2.Query()

	u1.RawQuery = ""
	u2.RawQuery = ""

	if u1.String() != u2.String() {
		return false
	}

	return reflect.DeepEqual(q1, q2)
}
