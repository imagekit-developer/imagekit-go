package imagekit

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/imagekit-developer/imagekit-go/api/media"
	"github.com/imagekit-developer/imagekit-go/api/metadata"
	"github.com/imagekit-developer/imagekit-go/api/uploader"
	"github.com/imagekit-developer/imagekit-go/config"
	"github.com/imagekit-developer/imagekit-go/logger"
)

// ImageKit main struct
type ImageKit struct {
	Config   config.Configuration
	Logger   *logger.Logger
	Media    *media.API
	Metadata *metadata.API
	Uploader *uploader.API
	getToken func() string
}

// NewParams is a struct to define parameters to imagekit
type NewParams struct {
	PrivateKey  string
	PublicKey   string
	UrlEndpoint string
}

// New returns ImageKit object from environment variables
func New() (*ImageKit, error) {
	cfg, err := config.New()

	if err != nil {
		return nil, err
	}
	return NewFromConfiguration(cfg), nil
}

const DefaultTokenExpire = 60 * 30

func getToken() string {
	uuid := uuid.New()
	return uuid.String()
}

// NewFromParams returns a new ImageKit object from provided parameters
func NewFromParams(params NewParams) *ImageKit {
	return NewFromConfiguration(
		config.NewFromParams(params.PrivateKey, params.PublicKey, params.UrlEndpoint),
	)
}

// NewFromConfiguration returns new ImageKit object from configuration object
func NewFromConfiguration(cfg *config.Configuration) *ImageKit {
	log := logger.New()
	client := &http.Client{}

	return &ImageKit{
		Config: *cfg,
		Logger: log,
		Media: &media.API{
			Config: *cfg,
			Logger: log,
			Client: client,
		},
		Metadata: &metadata.API{
			Config: *cfg,
			Logger: log,
			Client: client,
		},
		Uploader: &uploader.API{
			Config: *cfg,
			Logger: log,
			Client: client,
		},
		getToken: getToken,
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

// SignToken signs given token and expiration timestamp with private key
func (ik *ImageKit) SignToken(param SignTokenParam) SignedToken {
	if param.Token == "" {
		param.Token = ik.getToken()
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
