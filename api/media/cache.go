package media

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/imagekit-developer/imagekit-go/api"
	"gopkg.in/validator.v2"
)

type RequestId struct {
	RequestId string `json:"requestId"`
}

type PurgeCacheResponse struct {
	Data RequestId
	api.Response
}

type PurgeCacheParam struct {
	Url string `validate:"nonzero" json:"url"`
}

type PurgeCacheStatus struct {
	Status string `json:"status"`
}

type PurgeCacheStatusResponse struct {
	Data PurgeCacheStatus
	api.Response
}

// PurgeCache purges cache and returns requestId in response data.
func (m *API) PurgeCache(ctx context.Context, param PurgeCacheParam) (*PurgeCacheResponse, error) {
	var err error
	var response = &PurgeCacheResponse{}

	if err = validator.Validate(&param); err != nil {
		return nil, err
	}

	resp, err := m.post(ctx, "files/purge", &param, response)

	if err != nil {
		return response, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		err = response.ParseError()
	} else {
		err = json.Unmarshal(response.Body(), &response.Data)
	}
	return response, err
}

// PurgeCacheStatus returns status of purge cache request
func (m *API) PurgeCacheStatus(ctx context.Context, requestId string) (*PurgeCacheStatusResponse, error) {
	var err error
	var response = &PurgeCacheStatusResponse{}

	if requestId == "" {
		return nil, errors.New("requestId can not be empty")
	}

	resp, err := m.get(ctx, "files/purge/"+requestId, response)

	if err != nil {
		return response, err
	}

	if resp.StatusCode != 200 {
		err = response.ParseError()
	} else {
		err = json.Unmarshal(response.Body(), &response.Data)
	}
	return response, err

}
