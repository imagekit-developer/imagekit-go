package imagekit

import (
	neturl "net/url"
	"os"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/imagekit-developer/imagekit-go/logger"
	ikurl "github.com/imagekit-developer/imagekit-go/url"
)

var imgkit = NewFromParams(NewParams{
	PrivateKey:  "private_",
	PublicKey:   "public_",
	EndpointUrl: "https://ik.imagekit.io/test/",
})

func init() {
	imgkit.Logger.SetLevel(logger.DEBUG)
	os.Setenv("IMAGEKIT_PRIVATE_KEY", "private_")
	os.Setenv("IMAGEKIT_PUBLIC_KEY", "public_")
	os.Setenv("IMAGEKIT_ENDPOINT_URL", "https://ik.imagekit.io/test/")
}

func Test_New(t *testing.T) {
	var ik any
	ik, err := New()
	if err != nil {
		t.Fatal(err)
	}

	if ik == nil {
		t.Error("null")
	}

	if _, ok := ik.(*ImageKit); !ok {
		t.Error("New() did not return *ImageKit")
	}
}

func TestUrl(t *testing.T) {
	type tcase struct {
		name   string
		params ikurl.UrlParam
		url    string
	}

	cases := []tcase{
		{
			name: "path",
			params: ikurl.UrlParam{
				Path:        "default-image.jpg",
				EndpointUrl: "https://imagekit.io/343534/",
				Transformations: []ikurl.Transformation{
					{
						Width:  100,
						Rotate: 90,
					},
				},
			},
			url: "https://imagekit.io/343534/tr:w-100,rt-90/default-image.jpg",
		}, {
			name: "signed-without-transformation",
			params: ikurl.UrlParam{
				Path:          "default-image.jpg",
				Signed:        true,
				ExpireSeconds: 100,
				UnixTime: func() int64 {
					return 1653775828
				},
			},
			url: "https://ik.imagekit.io/test/default-image.jpg?ik-t=1653775928&ik-s=48842eca663c6895331331db6c90f262c601f4e8",
		}, {
			name: "signed-with-transformation",
			params: ikurl.UrlParam{
				Path:        "default-image.jpg",
				EndpointUrl: "https://ik.imagekit.io/test/",
				Transformations: []ikurl.Transformation{
					{
						Width:  200,
						Rotate: 90,
					},
				},
				Signed:        true,
				ExpireSeconds: 100,
				UnixTime: func() int64 {
					return 1653775828
				},
			},
			url: "https://ik.imagekit.io/test/tr:w-200,rt-90/default-image.jpg?ik-t=1653775928&ik-s=6c74b32f5efc0cc39ab5c042718b36553d8a1606",
		}, {
			name: "src-with-transformation",
			params: ikurl.UrlParam{
				Src: "https://imagekit.io/343534/default-image.jpg",
				Transformations: []ikurl.Transformation{
					{
						Width:  100,
						Rotate: 90,
					},
				},
			},
			url: "https://imagekit.io/343534/default-image.jpg?tr=w-100,rt-90",
		}, {
			name: "src-without-transformation",
			params: ikurl.UrlParam{
				Src: "https://imagekit.io/343534/default-image.jpg",
			},
			url: "https://imagekit.io/343534/default-image.jpg",
		}, {
			name: "trquery",
			params: ikurl.UrlParam{
				Path:        "default-image.jpg",
				EndpointUrl: "https://imagekit.io/343534/",
				Transformations: []ikurl.Transformation{
					{
						Width:  100,
						Rotate: 90,
					},
				},
				TransformationPosition: ikurl.QUERY,
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
				t.Errorf("expected url: " + tc.url + "\n got: " + url)
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

func TestSignToken(t *testing.T) {
	var expire int64 = 1655379249 + DefaultTokenExpire
	var unix = func() int64 { return 1655379249 }
	var token = "xxxx-xxxx-xxxxxxxx"

	cases := map[string]struct {
		param  SignTokenParam
		result SignedToken
	}{
		"with param": {
			param:  SignTokenParam{Token: "31c468de-520a-4dc1-8868-de1e0fb93a7b", Expires: 1655379249},
			result: SignedToken{Token: "31c468de-520a-4dc1-8868-de1e0fb93a7b", Expires: 1655379249, Signature: "ed6f1aadeec33eb3509c0576e6a05100861c64c5"},
		},
		"use defaults": {
			param:  SignTokenParam{Token: "31c468de-520a-4dc1-8868-de1e0fb93a7b", unix: unix},
			result: SignedToken{Token: "31c468de-520a-4dc1-8868-de1e0fb93a7b", Expires: expire, Signature: "b3c708386f56dbcdac2a6e650ab00789ec37645e"},
		},
		"autogen-token": {
			param:  SignTokenParam{unix: unix},
			result: SignedToken{Token: token, Expires: expire, Signature: "c46ef585f970b560aea69b90e32cd002c6639515"},
		},
	}

	imgkit.getToken = func() string {
		return token
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			resp := imgkit.SignToken(tc.param)

			if !cmp.Equal(resp, tc.result) {
				t.Errorf("%v\n%v", resp, tc.result)
			}
		})
	}

}
