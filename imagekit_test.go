package imagekit

import (
	"testing"

	imgkiturl "github.com/dhaval070/imagekit-go/url"
)

var imgkit = NewFromParams(NewParams{
	PrivateKey:  "xxx",
	PublicKey:   "yyy",
	EndpointUrl: "https://ik.imagekit.io/dk1m7xkgx/",
})

func TestUrl(t *testing.T) {
	param := imgkiturl.UrlParams{
		Path:           "default-image.jpg",
		UrlEndpoint:    "https://imagekit.io/343534/",
		Transformation: "w-100,rt-90",
	}

	var expected = "https://imagekit.io/343534/tr:w-100,rt-90/default-image.jpg"

	url, err := imgkit.Url(param)

	if err != nil {
		t.Errorf(err.Error())
	}

	if url != expected {
		t.Errorf("Unexpected url " + url + ", expected: " + expected)
	}

	param = imgkiturl.UrlParams{
		Src:            "https://imagekit.io/343534/default-image.jpg",
		Transformation: "w-100,rt-90",
	}
	expected = "https://imagekit.io/343534/default-image.jpg?tr=w-100,rt-90"

	url, err = imgkit.Url(param)

	if err != nil {
		t.Errorf(err.Error())
	}

	if url != expected {
		t.Errorf("Unexpected url " + url + ", expected: " + expected)
	}
}
