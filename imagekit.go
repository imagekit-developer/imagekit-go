package imagekit

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	neturl "net/url"
	"strconv"
	"strings"
	"time"

	"github.com/dhaval070/imagekit-go/config"
	"github.com/dhaval070/imagekit-go/logger"
	ikurl "github.com/dhaval070/imagekit-go/url"
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

func (ik *ImageKit) Url(params ikurl.UrlParams) (string, error) {
	var resultUrl string
	var url *neturl.URL
	var err error
	var endpoint = params.EndpointUrl

	if endpoint == "" {
		endpoint = ik.Config.Cloud.EndpointUrl
	}

	if endpoint == "" {
		endpoint = ik.Config.Cloud.EndpointUrl
	}

	endpoint = strings.TrimRight(endpoint, "/") + "/"

	if params.QueryParameters == nil {
		params.QueryParameters = make(map[string]string)
	}

	if params.Src == "" {
		if url, err = neturl.Parse(endpoint); err != nil {
			return "", err
		}

		if params.Transformation != "" {
			if params.TransformationPosition == ikurl.QUERY {
				params.QueryParameters["tr"] = params.Transformation
				url, err = neturl.Parse(endpoint + params.Path)

			} else {
				url, err = neturl.Parse(url.String() +
					"tr:" + params.Transformation +
					"/" + strings.TrimLeft(params.Path, "/"))
			}
		}
	} else {
		if url, err = neturl.Parse(params.Src); err != nil {
			return "", err
		}

		if params.Transformation != "" {
			params.QueryParameters["tr"] = params.Transformation
		}
	}

	if err != nil {
		return "", nil
	}

	query := url.Query()

	for k, v := range params.QueryParameters {
		query.Set(k, v)
	}
	url.RawQuery = query.Encode()
	resultUrl = url.String()

	if params.Signed {
		var now int64

		if params.UnixTime == nil {
			now = time.Now().Unix()
			log.Println(now)
		} else {
			now = params.UnixTime()
		}

		var expires = strconv.FormatInt(now+int64(params.ExpireSeconds), 10)
		var path = strings.Replace(resultUrl, endpoint, "", 1)

		path = path + expires
		mac := hmac.New(sha1.New, []byte(ik.Config.Cloud.PrivateKey))
		mac.Write([]byte(path))
		signature := hex.EncodeToString(mac.Sum(nil))

		if strings.Index(resultUrl, "?") > -1 {
			resultUrl = resultUrl + "&" + fmt.Sprintf("ik-t=%s&ik-s=%s", expires, signature)
		} else {
			resultUrl = resultUrl + "?" + fmt.Sprintf("ik-t=%s&ik-s=%s", expires, signature)
		}
	}

	return resultUrl, nil
}
