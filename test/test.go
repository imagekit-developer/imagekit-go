package test

import (
	"log"
	"net/http"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/imagekit-developer/imagekit-go/config"
)

var Cfg *config.Configuration

const (
	PrivateKey  = "private_"
	PublicKey   = "public_"
	UrlEndpoint = "https://ik.imagekit.io/tests/"
)

func init() {
	Cfg = config.NewFromParams(PrivateKey, PublicKey, UrlEndpoint)
}

// LogoURL is the URL of the publicly available logo.
const LogoURL = "https://cloudinary-res.cloudinary.com/image/upload/cloudinary_logo.png"

// Base64Image us a base64 encoded test image.
const Base64Image = "data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7"

// ImageFilePath is a full path to the test image file.
var ImageFilePath = TestDataDir() + "new-york-cityscape-buildings.jpg"

// TestDataDir returns the full path to the directory with test files.
func TestDataDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))

	return filepath.Dir(d) + "/test/testdata/"
}

func JsonRequest(r *http.Request, t *testing.T) {
	h := r.Header

	log.Println(h.Get("Authorization"))

	if h.Get("Authorization") != "Basic cHJpdmF0ZV86" {
		t.Error("invalid authorization header")
	}

	if h.Get("Content-Type") != "application/json" {
		t.Error("content type not application/json")
	}
}
