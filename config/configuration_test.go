package config_test

import (
	"testing"

	"github.com/imagekit-developer/imagekit-go/config"
	"github.com/stretchr/testify/assert"
)

func TestConfiguration_CreateInstance(t *testing.T) {
	c, _ := config.New()

	c = config.NewFromParams("private", "public", "https://example/nature")

	assert.Equal(t, "private", c.Cloud.PrivateKey)
	assert.Equal(t, "public", c.Cloud.PublicKey)
}
