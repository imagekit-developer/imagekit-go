package imagekit

import (
	neturl "net/url"
	"strings"

	"github.com/dhaval070/imagekit-go/config"
	"github.com/dhaval070/imagekit-go/logger"
	imgkiturl "github.com/dhaval070/imagekit-go/url"
)

type ImageKit struct {
	Config config.Configuration
	Logger *logger.Logger
}

type NewParams struct {
	PrivateKey  string
	PublicKey   string
	EndpointUrl string
}

func New() (*ImageKit, error) {
	cfg, err := config.New()

	if err != nil {
		return nil, err
	}
	return NewFromConfiguration(cfg), nil
}

func NewFromParams(params NewParams) *ImageKit {
	return NewFromConfiguration(
		config.NewFromParams(params.PrivateKey, params.PublicKey, params.EndpointUrl),
	)
}

func NewFromConfiguration(cfg *config.Configuration) *ImageKit {
	log := logger.New()

	return &ImageKit{
		Config: *cfg,
		Logger: log,
	}
}

func Url(params imgkiturl.UrlParams) (string, error) {
	var url *neturl.URL
	var err error

	if params.UrlEndpoint != "" {
		params.UrlEndpoint = strings.TrimRight(params.UrlEndpoint, "/")

		if url, err = neturl.Parse(params.UrlEndpoint); err != nil {
			return "", err
		}

		if params.Transformation != "" {
			url, err = neturl.Parse(url.String() + "/tr:" + params.Transformation + "/" + strings.TrimLeft(params.Path, "/"))
		}

	} else {
		if url, err = neturl.Parse(params.Src); err != nil {
			return "", err
		}

		if params.Transformation != "" {
			v := neturl.Values{}
			v.Set("tr", params.Transformation)

			return url.String() + "?" + v.Encode(), nil
		}
	}

	return url.String(), nil
}
