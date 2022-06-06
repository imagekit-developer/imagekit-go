package uploader

import (
	"bytes"
	"context"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/dhaval070/imagekit-go/api"
	"github.com/dhaval070/imagekit-go/config"
	"github.com/dhaval070/imagekit-go/logger"
)

// API is the upload feature main struct
type API struct {
	Config config.Configuration
	Logger *logger.Logger
	Client api.HttpClient
}

// New creates a new Uploader API instance from the environment variable.
func New() (*API, error) {
	c, err := config.New()
	if err != nil {
		return nil, err
	}

	return NewFromConfiguration(c)
}

// NewFromConfiguration a new Upload API instance with the given Configuration.
func NewFromConfiguration(c *config.Configuration) (*API, error) {
	return &API{
		Config: *c,
		Client: &http.Client{},
		Logger: logger.New(),
	}, nil
}

// postFile uploads file with url.Values parameters
func (u *API) postFile(ctx context.Context, file interface{}, formParams url.Values) (*http.Response, error) {
	uploadEndpoint := api.BuildPath("files", "upload")

	switch fileValue := file.(type) {
	case string:
		if !api.IsLocalFilePath(file) {
			// Can be URL, Base64 encoded string, etc.
			formParams.Add("file", fileValue)

			return u.postForm(ctx, uploadEndpoint, formParams)
		}

		return u.postLocalFile(ctx, uploadEndpoint, fileValue, formParams)
	case io.Reader:
		return u.postIOReader(ctx, uploadEndpoint, fileValue, formParams, map[string]string{})

	default:
		return nil, errors.New("unsupported file type")
	}
}

// postLocalFile creates a new file upload http request with optional extra params.
func (u *API) postLocalFile(ctx context.Context, urlPath string, filePath string, formParams url.Values) (*http.Response, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer api.DeferredClose(file)

	if _, ok := formParams["fileName"]; !ok {
		st, err := file.Stat()
		if err != nil {
			return nil, err
		}
		formParams.Set("fileName", st.Name())
	}

	return u.postIOReader(ctx, urlPath, file, formParams, map[string]string{})
}

// postIOReader uploads file using io.Reader
func (u *API) postIOReader(ctx context.Context, urlPath string, reader io.Reader, formParams url.Values, headers map[string]string) (*http.Response, error) {
	bodyBuf := new(bytes.Buffer)
	formWriter := multipart.NewWriter(bodyBuf)

	headers["Content-Type"] = formWriter.FormDataContentType()

	for key, val := range formParams {
		_ = formWriter.WriteField(key, val[0])
	}

	partWriter, err := formWriter.CreateFormFile("file", formParams.Get("fileName"))
	if err != nil {
		return nil, err
	}

	if _, err = io.Copy(partWriter, reader); err != nil {
		return nil, err
	}

	if err = formWriter.Close(); err != nil {
		return nil, err
	}

	if u.Config.API.UploadTimeout != 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Duration(u.Config.API.UploadTimeout)*time.Second)
		defer cancel()
	}

	return u.postBody(ctx, urlPath, bodyBuf, headers)
}

func (u *API) postBody(ctx context.Context, urlPath string, bodyBuf *bytes.Buffer, headers map[string]string) (*http.Response, error) {

	req, err := http.NewRequest(http.MethodPost,
		u.Config.API.UploadPrefix+urlPath,
		bodyBuf,
	)

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(u.Config.Cloud.PrivateKey, "")

	for key, val := range headers {
		req.Header.Add(key, val)
	}

	req = req.WithContext(ctx)

	return u.Client.Do(req)
}

func (u *API) postForm(ctx context.Context, urlPath string, formParams url.Values) (*http.Response, error) {

	bodyBuf := new(bytes.Buffer)
	writer := multipart.NewWriter(bodyBuf)

	for k, _ := range formParams {
		writer.WriteField(k, formParams.Get(k))
	}
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	h := map[string]string{"Content-Type": writer.FormDataContentType()}
	ctx, cancel := context.WithTimeout(ctx, time.Duration(u.Config.API.Timeout)*time.Second)
	defer cancel()

	return u.postBody(ctx, urlPath, bodyBuf, h)
}
