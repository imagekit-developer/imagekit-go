// Package config defines the ImageKit configuration.
package config

import (
	"errors"
	"os"
)

// Configuration is the main configuration struct.
type Configuration struct {
	API   API
	Cloud Cloud
}

// New returns a new Configuration instance from the environment variables
func New() (*Configuration, error) {
	privateKey := os.Getenv("IMAGEKIT_PRIVATE_KEY")
	publicKey := os.Getenv("IMAGEKIT_PUBLIC_KEY")
	endpointUrl := os.Getenv("IMAGEKIT_ENDPOINT_URL")

	switch {
	case privateKey == "":
		return nil, errors.New("IMAGEKIT_PRIVATE_KEY envvar not set")
	case publicKey == "":
		return nil, errors.New("IMAGEKIT_PUBLIC_KEY envvar not set")
	case endpointUrl == "":
		return nil, errors.New("IMAGEKIT_ENDPOINT_URL envvar not set")
	}

	return NewFromParams(privateKey, publicKey, endpointUrl), nil
}

// NewFromParams returns a new Configuration instance from the provided keys and endpointUrl.
func NewFromParams(privateKey string, publicKey string, endpointUrl string) *Configuration {
	cloudConf := Cloud{
		PrivateKey:  privateKey,
		PublicKey:   publicKey,
		EndpointUrl: endpointUrl,
	}

	return &Configuration{
		Cloud: cloudConf,
		API: API{
			UploadPrefix: "https://upload.imagekit.io/api/v1/",
			Timeout:      60,
		},
	}
}
