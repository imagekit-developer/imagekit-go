package imagekit

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	neturl "net/url"
	"strconv"
	"strings"
	"time"

	"github.com/dhaval070/imagekit-go/api/media"
	"github.com/dhaval070/imagekit-go/api/metadata"
	"github.com/dhaval070/imagekit-go/config"
	"github.com/dhaval070/imagekit-go/logger"
	ikurl "github.com/dhaval070/imagekit-go/url"
	"github.com/google/uuid"
)

// ImageKit main struct
type ImageKit struct {
	Config   config.Configuration
	Logger   *logger.Logger
	Media    *media.API
	Metadata *metadata.API
}

// NewParams is struct to define parameters to imagekit
type NewParams struct {
	PrivateKey  string
	PublicKey   string
	EndpointUrl string
}

// New returns ImageKit object from environment variables
func New() (*ImageKit, error) {
	cfg, err := config.New()

	if err != nil {
		return nil, err
	}
	return NewFromConfiguration(cfg), nil
}

// NewFromParams return new ImageKit object from provided parameters
func NewFromParams(params NewParams) *ImageKit {
	return NewFromConfiguration(
		config.NewFromParams(params.PrivateKey, params.PublicKey, params.EndpointUrl),
	)
}

// NewFromConfiguration returns new ImageKit object from configuration object
func NewFromConfiguration(cfg *config.Configuration) *ImageKit {
	log := logger.New()

	return &ImageKit{
		Config: *cfg,
		Logger: log,
		Media: &media.API{
			Config: *cfg,
			Logger: log,
			Client: &http.Client{},
		},
		Metadata: &metadata.API{
			Config: *cfg,
			Logger: log,
			Client: &http.Client{},
		},
	}
}

type SignTokenParam struct {
	Token   string
	Expires int64
	unix    func() int64
}

type SignedToken struct {
	Token     string
	Expires   int64
	Signature string
}

const DefaultTokenExpire = 60 * 30

// Url generates url from UrlParams
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

// SignToken signs given token and expiration timestamp with private key
func (ik *ImageKit) SignToken(param SignTokenParam) SignedToken {
	if param.Token == "" {
		uuid := uuid.New()
		param.Token = uuid.String()
	}

	if param.Expires == 0 {
		var e int64

		if param.unix == nil {
			e = time.Now().Unix()
		} else {
			e = param.unix()
		}
		param.Expires = e + DefaultTokenExpire
	}

	log.Println(param)
	mac := hmac.New(sha1.New, []byte(ik.Config.Cloud.PrivateKey))
	mac.Write([]byte(param.Token + strconv.FormatInt(param.Expires, 10)))
	signature := hex.EncodeToString(mac.Sum(nil))
	return SignedToken{Token: param.Token, Expires: param.Expires, Signature: signature}
}
