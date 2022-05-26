package imagekit

import (
	"github.com/dhaval070/imagekit-go/config"
	"github.com/dhaval070/imagekit-go/logger"
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
