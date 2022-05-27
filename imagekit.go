package imagekit

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	neturl "net/url"
	"strconv"
	"strings"
	"time"

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

func (ik *ImageKit) Url(params imgkiturl.UrlParams) (string, error) {
	var resultUrl string
	var url *neturl.URL
	var err error

	if params.UrlEndpoint != "" {
		params.UrlEndpoint = strings.TrimRight(params.UrlEndpoint, "/")

		if url, err = neturl.Parse(params.UrlEndpoint); err != nil {
			return "", err
		}

		if params.Transformation != "" {
			url, err = neturl.Parse(url.String() +
				"/tr:" + params.Transformation +
				"/" + strings.TrimLeft(params.Path, "/"))

			resultUrl = url.String()
		}

	} else {
		if url, err = neturl.Parse(params.Src); err != nil {
			return "", err
		}

		resultUrl = url.String()

		if params.Transformation != "" {
			resultUrl = resultUrl + "?tr=" + params.Transformation
		}
	}

	//TODO: write test
	if params.Signed {
		var now int64

		if params.UnixTime == nil {
			now = time.Now().Unix()
		} else {
			now = params.UnixTime()
		}

		var expires = strconv.FormatInt(now+int64(params.ExpireSeconds), 10)
		var path = strings.Replace(resultUrl, params.UrlEndpoint+"/", "", 1)

		path = path + expires
		mac := hmac.New(sha1.New, []byte(ik.Config.Cloud.PrivateKey))
		mac.Write([]byte(path))
		signature := hex.EncodeToString(mac.Sum(nil))
		resultUrl = resultUrl + "?ik-t=" + expires + "&ik-s=" + signature
	}

	return resultUrl, nil
}
