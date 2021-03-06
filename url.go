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

	ikurl "github.com/imagekit-developer/imagekit-go/url"
)

func join(sep string, args ...ikurl.Transformation) string {
	var parts []string

	for _, v := range args {
		parts = append(parts, v.String())
	}
	return strings.Join(parts, sep)
}

// Url generates url from UrlParam
func (ik *ImageKit) Url(params ikurl.UrlParam) (string, error) {
	var resultUrl string
	var url *neturl.URL
	var err error
	var endpoint = params.EndpointUrl

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

		if params.Transformations == nil {
			if url, err = neturl.Parse(endpoint + params.Path); err != nil {
				return "", err
			}
		} else {
			if params.TransformationPosition == ikurl.QUERY {
				params.QueryParameters["tr"] = join(":", params.Transformations...)
				url, err = neturl.Parse(endpoint + params.Path)

			} else {
				url, err = neturl.Parse(url.String() +
					"tr:" + join(":", params.Transformations...) +
					"/" + strings.TrimLeft(params.Path, "/"))
			}
		}
	} else {
		if url, err = neturl.Parse(params.Src); err != nil {
			return "", err
		}

		if params.Transformations != nil {
			params.QueryParameters["tr"] = join(":", params.Transformations...)
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
