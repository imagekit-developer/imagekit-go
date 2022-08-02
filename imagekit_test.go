package imagekit

import (
	neturl "net/url"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/imagekit-developer/imagekit-go/logger"
	ikurl "github.com/imagekit-developer/imagekit-go/url"
)

var imgkit = NewFromParams(NewParams{
	PrivateKey:  "private_",
	PublicKey:   "public_",
	UrlEndpoint: "https://ik.imagekit.io/test/",
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
				UrlEndpoint: "https://imagekit.io/343534/",
				Transformations: []map[string]any{
					{
						"width":    100,
						"rotation": 90,
					},
				},
			},
			url: "https://imagekit.io/343534/tr:w-100,rt-90/default-image.jpg",
		}, {
			name: "signed-url",
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
			name: "src-with-transformation",
			params: ikurl.UrlParam{
				Src: "https://imagekit.io/343534/default-image.jpg",
				Transformations: []map[string]any{
					{
						"width":    100,
						"rotation": 90,
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
				UrlEndpoint: "https://imagekit.io/343534/",
				Transformations: []map[string]any{
					{
						"width":    100,
						"rotation": 90,
					},
				},
				TransformationPosition: ikurl.QUERY,
			},
			url: "https://imagekit.io/343534/default-image.jpg?tr=w-100%2crt-90",
		}, {
			name: "transformations",
			params: ikurl.UrlParam{
				Path:        "default-image.jpg",
				UrlEndpoint: "https://ik.imagekit.io/dk1m7xkgi/",
				Transformations: []map[string]any{
					{
						"width":          200,
						"height":         400,
						"cropMode":       "extract",
						"focus":          "center",
						"x":              100,
						"y":              110,
						"quality":        85,
						"format":         "auto",
						"blur":           5,
						"dpr":            "auto",
						"effectGray":     "-",
						"defaultImage":   "/test2_hBIIEweBy.gif",
						"progressive":    true,
						"lossless":       true,
						"trim":           true,
						"border":         "5_005500",
						"colorProfile":   true,
						"metadata":       true,
						"rotation":       "auto",
						"radius":         40,
						"background":     "344222",
						"effectContrast": "-",
						"effectSharpen":  "-",
						"raw":            "x-1",
					},
				},
			},
			url: "https://ik.imagekit.io/dk1m7xkgi/tr:w-200,h-400,cm-extract,fo-center,x-100,y-110,q-85,f-auto,bl-5,dpr-auto,e-grayscale,di-test2_hBIIEweBy.gif,pr-true,lo-true,t-true,b-5_005500,cp-true,md-true,rt-auto,r-40,bg-344222,e-sharpen,e-contrast,x-1/default-image.jpg",
		}, {
			name: "aspect-ratio-xc-yc",
			params: ikurl.UrlParam{
				Path: "default-image.jpg",
				Transformations: []map[string]any{
					{
						"width":         200,
						"aspectRatio":   "16-9",
						"cropMode":      "extract",
						"focus":         "center",
						"xc":            100,
						"yc":            110,
						"quality":       85,
						"format":        "auto",
						"blur":          50,
						"dpr":           2,
						"rotation":      90,
						"effectSharpen": 40,
					},
				},
			},
			url: "https://ik.imagekit.io/test/tr:w-200,ar-16-9,cm-extract,fo-center,xc-100,yc-110,q-85,f-auto,bl-50,dpr-2,rt-90,e-sharpen-40/default-image.jpg",
		}, {
			name: "unsharp-mask",
			params: ikurl.UrlParam{
				Path: "default-image.jpg",
				Transformations: []map[string]any{
					{
						"effectUSM": "2-2-0.8-0.024",
					},
				},
			},
			url: "https://ik.imagekit.io/test/tr:e-usm-2-2-0.8-0.024/default-image.jpg",
		}, {
			name: "chained transformations",
			params: ikurl.UrlParam{
				Path:        "default-image.jpg",
				UrlEndpoint: "https://ik.imagekit.io/343534/",
				Transformations: []map[string]any{
					{
						"width":  100,
						"height": 200,
					}, {
						"rotation": 90,
					},
				},
			},
			url: "https://ik.imagekit.io/343534/tr:w-100,h-200:rt-90/default-image.jpg",
		}, {

			name: "common-overlay-options",
			params: ikurl.UrlParam{
				Path: "default-image.jpg",
				Transformations: []map[string]any{
					{
						"overlayX":          100,
						"overlayY":          110,
						"overlayHeight":     100,
						"overlayWidth":      90,
						"overlayBackground": "443322",
						"overlayFocus":      "bottom",
					},
				},
			},
			url: "https://ik.imagekit.io/test/tr:ox-100,oy-110,oh-100,ow-90,obg-443322,ofo-bottom/default-image.jpg",
		}, {
			name: "text-overlay-options",
			params: ikurl.UrlParam{
				Path: "default-image.jpg",
				Transformations: []map[string]any{
					{
						"overlayText":               "this is a sample overlay",
						"overlayX":                  100,
						"overlayY":                  110,
						"overlayHeight":             500,
						"overlayWidth":              900,
						"overlayTextPadding":        "20_40",
						"overlayTextBackground":     "ffffff",
						"overlayTextInnerAlignment": "right",
						"overlayTextColor":          "blue",
						"overlayTextFontFamily":     "Arvo",
						"overlayTextFontSize":       40,
						"overlayTextTypography":     "ib",
						"overlayRadius":             20,
					},
				},
			},
			url: "https://ik.imagekit.io/test/tr:ot-this%20is%20a%20sample%20overlay,ox-100,oy-110,oh-500,ow-900,otbg-ffffff,otp-20_40,otia-right,otc-blue,otf-Arvo,ots-40,ott-ib,or-20/default-image.jpg",
		}, {
			name: "image-overlay-options",
			params: ikurl.UrlParam{
				Path: "default-image.jpg",
				Transformations: []map[string]any{
					{
						"overlayImage":         "/test2_hBIIEweBy.gif",
						"overlayX":             100,
						"overlayY":             110,
						"overlayHeight":        200,
						"overlayWidth":         200,
						"overlayImageBorder":   "4_blue",
						"overlayImageDPR":      0.2,
						"overlayImageQuality":  80,
						"overlayImageCropping": "at_max",
						"overlayImageX":        100,
						"overlayImageY":        20,
						"overlayImageTrim":     false,
					},
				},
			},
			url: "https://ik.imagekit.io/test/tr:oi-test2_hBIIEweBy.gif,ox-100,oy-110,oh-200,ow-200,oib-4_blue,oidpr-0.2,oiq-80,oic-at_max,oix-100,oiy-20,oit-false/default-image.jpg",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			url, err := imgkit.Url(tc.params)

			if err != nil {
				t.Errorf(err.Error())
			}

			if strings.Index(tc.url, "tr:") > -1 {
				url, tr := extractTransformation(t, url)
				expectedUrl, expectedTr := extractTransformation(t, tc.url)

				if !urlsEquals(url, expectedUrl) {
					t.Errorf("expected url: %s\ngot: %s", expectedUrl, url)
				}

				if !cmp.Equal(tr, expectedTr) {
					t.Errorf("url: %s, expected tr: %s\ngot tr: %s", tc.url, expectedTr, tr)
				}
			}
		})
	}

}

func extractTransformation(t *testing.T, url string) (urlResult string, trResult []string) {
	re := regexp.MustCompile("tr:(.+)/")
	m := re.FindStringSubmatch(url)

	if m == nil {
		t.Error("transformation not found")
		return
	}

	urlResult = strings.Replace(url, "tr:"+m[1], "", 1)
	trs := strings.Split(m[1], ":")

	for _, tr := range trs {
		parts := strings.Split(tr, ",")
		sort.Strings(parts)
		trResult = append(trResult, parts...)
	}

	return urlResult, trResult
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

func Test_SignToken(t *testing.T) {
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
