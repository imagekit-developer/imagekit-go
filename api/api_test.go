package api

import (
	"errors"
	"testing"
)

func TestApi(t *testing.T) {
	t.Error(errors.New("api fail"))
}
