package media

import (
	"context"

	"github.com/dhaval070/imagekit-go/api"
	"gopkg.in/validator.v2"
)

// CreateFolder creates a new folder in media library
func (m *API) CreateFolder(ctx context.Context, param CreateFolderParam) (*api.Response, error) {
	var err error
	var response = &api.Response{}

	if err = validator.Validate(&param); err != nil {
		return nil, err
	}

	resp, err := m.Post(ctx, "folder", &param)
	defer api.DeferredBodyClose(resp)

	api.SetResponseMeta(resp, response)

	if err != nil {
		return response, err
	}

	if resp.StatusCode != 201 {
		err = response.ParseError()
	}

	return response, err
}

// DeleteFolder removes the folder from media library
func (m *API) DeleteFolder(ctx context.Context, param DeleteFolderParam) (*api.Response, error) {
	var err error
	var response = &api.Response{}

	if err = validator.Validate(&param); err != nil {
		return nil, err
	}

	resp, err := m.Delete(ctx, "folder", &param)

	defer api.DeferredBodyClose(resp)

	api.SetResponseMeta(resp, response)

	if err != nil {
		return response, err
	}
	if resp.StatusCode != 204 {
		err = response.ParseError()
	}
	return response, err
}
