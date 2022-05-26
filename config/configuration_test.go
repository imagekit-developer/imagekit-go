package config_test

import (
	"testing"

	"github.com/dhaval070/imagekit-go/config"
	"github.com/stretchr/testify/assert"
)

var cldURL = "cloudinary://key:secret@test123"
var fakeOAuthToken = "MTQ0NjJkZmQ5OTM2NDE1ZTZjNGZmZjI4"

func TestConfiguration_CreateInstance(t *testing.T) {
	c, _ := config.New()

	c = config.NewFromParams("private", "public", "https://example/nature")

	assert.Equal(t, "private", c.Cloud.PrivateKey)
	assert.Equal(t, "public", c.Cloud.PublicKey)
}
