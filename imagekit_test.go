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
	PrivateKey:  "private_XxZH+I8BfOoIsY0M9CQtS4nyNSk=",
	PublicKey:   "public_fGfgv45RjwmkbzGMRy1gKTcHf4M=",
	EndpointUrl: "https://ik.imagekit.io/dk1m7xkgi/",
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
				EndpointUrl:    "https://ik.imagekit.io/dk1m7xkgi/",
				Transformation: "w-200,rt-90",
				Signed:         true,
				ExpireSeconds:  100,
				UnixTime: func() int64 {
					return 1653775828
				},
			},
			url: "https://ik.imagekit.io/dk1m7xkgi/tr:w-200,rt-90/default-image.jpg?ik-t=1653775928&ik-s=889f3563d432c34a668f5b954f8887aec766c2ec",
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
