package config_test

import (
	"os"
	"testing"

	"github.com/imagekit-developer/imagekit-go/config"
	"github.com/stretchr/testify/assert"
)

func TestConfiguration_CreateInstance(t *testing.T) {
	os.Unsetenv("IMAGEKIT_PRIVATE_KEY")
	os.Unsetenv("IMAGEKIT_PUBLIC_KEY")
	os.Unsetenv("IMAGEKIT_ENDPOINT_URL")

	defer os.Unsetenv("IMAGEKIT_PRIVATE_KEY")
	defer os.Unsetenv("IMAGEKIT_PUBLIC_KEY")
	defer os.Unsetenv("IMAGEKIT_ENDPOINT_URL")

	c, err := config.New()

	if err == nil {
		t.Error("expected error")
	}

	os.Setenv("IMAGEKIT_PRIVATE_KEY", "private_")

	c, err = config.New()

	if err == nil {
		t.Error("expected error")
	}

	os.Setenv("IMAGEKIT_PUBLIC_KEY", "public_")
	c, err = config.New()

	if err == nil {
		t.Error("expected error")
	}

	os.Setenv("IMAGEKIT_PUBLIC_KEY", "public_")
	c, err = config.New()

	if err == nil {
		t.Error("expected error")
	}

	os.Setenv("IMAGEKIT_ENDPOINT_URL", "https://ik.imagekit.io/test/")

	c, err = config.New()

	if err != nil {
		t.Error("Unexpected error")
	}

	c = config.NewFromParams("private", "public", "https://example/nature")

	assert.Equal(t, "private", c.Cloud.PrivateKey)
	assert.Equal(t, "public", c.Cloud.PublicKey)
}
