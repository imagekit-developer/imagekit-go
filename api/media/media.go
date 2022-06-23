package media

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/imagekit-developer/imagekit-go/api"
	"github.com/imagekit-developer/imagekit-go/config"
	"github.com/imagekit-developer/imagekit-go/logger"
)

// API is the main struct for media
type API struct {
	Config config.Configuration
	Logger *logger.Logger
	Client api.HttpClient
}

// New creates a new Media API instance from the environment variable.
func New() (*API, error) {
	c, err := config.New()
	if err != nil {
		return nil, err
	}

	return NewFromConfiguration(c)
}

// NewFromConfiguration a new Media API instance with the given Configuration.
func NewFromConfiguration(c *config.Configuration) (*API, error) {
	return &API{
		Config: *c,
		Client: &http.Client{},
		Logger: logger.New(),
	}, nil
}

func (m *API) post(ctx context.Context, url string, data interface{}) (*http.Response, error) {
	url = api.BuildPath(m.Config.API.Prefix, url)
	var err error
	var body []byte

	if data != nil {
		if body, err = json.Marshal(data); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(m.Config.Cloud.PrivateKey, "")

	return m.Client.Do(req.WithContext(ctx))
}

func (m *API) get(ctx context.Context, url string) (*http.Response, error) {
	url = api.BuildPath(m.Config.API.Prefix, url)
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(m.Config.Cloud.PrivateKey, "")

	return m.Client.Do(req.WithContext(ctx))
}

func (m *API) delete(ctx context.Context, url string, data interface{}) (*http.Response, error) {
	var err error
	url = api.BuildPath(m.Config.API.Prefix, url)
	var body []byte

	if data != nil {
		if body, err = json.Marshal(data); err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(body))
	if data != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(m.Config.Cloud.PrivateKey, "")

	return m.Client.Do(req.WithContext(ctx))
}

func (m *API) patch(ctx context.Context, url string, data interface{}) (*http.Response, error) {
	url = api.BuildPath(m.Config.API.Prefix, url)
	var err error
	var body []byte

	if data != nil {
		if body, err = json.Marshal(data); err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(m.Config.Cloud.PrivateKey, "")

	return m.Client.Do(req.WithContext(ctx))
}

func (m *API) put(ctx context.Context, url string, data interface{}) (*http.Response, error) {
	url = api.BuildPath(m.Config.API.Prefix, url)
	var err error
	var body []byte

	if data != nil {
		if body, err = json.Marshal(data); err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(m.Config.Cloud.PrivateKey, "")

	return m.Client.Do(req.WithContext(ctx))
}
