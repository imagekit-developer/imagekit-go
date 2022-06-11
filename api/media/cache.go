package media

import (
	"context"
	"encoding/json"
	"log"

	"github.com/dhaval070/imagekit-go/api"
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

// PurgeCache purges cache and returns requestId in response data.
func (m *API) PurgeCache(ctx context.Context, param PurgeCacheParam) (*PurgeCacheResponse, error) {
	var err error
	var response = &PurgeCacheResponse{}

	if err = validator.Validate(&param); err != nil {
		return nil, err
	}

	resp, err := m.Post(ctx, "files/purge", &param)
	defer api.DeferredBodyClose(resp)
	log.Println(resp.StatusCode)
	api.SetResponseMeta(resp, response)

	if err != nil {
		return response, err
	}

	if resp.StatusCode != 201 {
		err = response.ParseError()
	} else {
		err = json.Unmarshal(response.Body(), &response.Data)
	}
	return response, err
}
