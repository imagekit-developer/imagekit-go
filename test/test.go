package test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/imagekit-developer/imagekit-go/api"
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

type Http struct {
	T      *testing.T
	Url    string
	Req    *http.Request
	Body   []byte
	Method string
}

func NewHttp(t *testing.T) *Http {
	return &Http{T: t}
}

func (h *Http) Handler(statusCode int, body string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.Url = r.URL.String()
		h.Method = r.Method
		h.Body, _ = io.ReadAll(r.Body)
		h.Req = r.Clone(context.Background())
		w.WriteHeader(statusCode)
		fmt.Fprintln(w, body)
	})
}

func (h *Http) Test(url string, method string, body any) {
	if h.Url != url {
		h.T.Errorf("expected url: %s, got: %s", url, h.Url)
	}

	if h.Method != method {
		h.T.Errorf("expected method: %s, got: %s", method, h.Method)
	}

	if method == "GET" || method == "DELETE" {
		return
	}

	var reqBody []byte

	if d, ok := body.([]byte); ok {
		reqBody = d
	} else {
		reqBody, _ = json.Marshal(body)
	}

	if !cmp.Equal(h.Body, reqBody) {
		h.T.Errorf("expected body: %v, got: %v", body, h.Body)
	}

	if h.Req != nil {
		JsonRequest(h.Req, h.T)
	}

}

type ErrorServer struct {
	t          *testing.T
	statusCode int
	ts         *httptest.Server
}

func NewErrorServer(t *testing.T) *ErrorServer {
	srv := &ErrorServer{t: t}
	srv.ts = httptest.NewServer(http.HandlerFunc(srv.handler))
	return srv
}

func (srv *ErrorServer) Url() string {
	return srv.ts.URL
}

func (srv *ErrorServer) handler(w http.ResponseWriter, r *http.Request) {
	log.Println(srv.statusCode)
	w.WriteHeader(srv.statusCode)
	fmt.Fprintln(w, "{}")
}

func (srv *ErrorServer) TestErrors(fn func() error) {
	var codesToErrors = map[int]error{
		400: api.ErrBadRequest,
		401: api.ErrUnauthorized,
		403: api.ErrForbidden,
		404: api.ErrNotFound,
		429: api.ErrTooManyRequests,
		500: api.ErrServer,
		502: api.ErrServer,
		503: api.ErrServer,
		504: api.ErrServer,
	}

	for code, expectedErr := range codesToErrors {
		srv.statusCode = code
		err := fn()
		if !errors.Is(err, expectedErr) {
			srv.t.Errorf("code %d: expected error %v, got: %v", code, expectedErr, err)
		}
	}
}
