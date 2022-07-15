package metadata

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	neturl "net/url"

	"github.com/imagekit-developer/imagekit-go/api"
	"github.com/imagekit-developer/imagekit-go/config"
	"github.com/imagekit-developer/imagekit-go/logger"
	"gopkg.in/validator.v2"
)

// API is the main struct for media
type API struct {
	Config config.Configuration
	Logger *logger.Logger
	Client api.HttpClient
}

// New creates a new Media API instance from the environment variable.
func New() (*API, error) {
	c, err := config.New()
	if err != nil {
		return nil, err
	}

	return NewFromConfiguration(c)
}

// NewFromConfiguration a new Media API instance with the given Configuration.
func NewFromConfiguration(c *config.Configuration) (*API, error) {
	return &API{
		Config: *c,
		Client: &http.Client{},
		Logger: logger.New(),
	}, nil
}

// MetadataResponse represents main struct of metadata response of the sdk
type MetadataResponse struct {
	Data Metadata
	api.Response
}

// Metadata represents struct of metadata response from api
type Metadata struct {
	Height          int
	Width           int
	Size            int64
	Format          string
	HasColorProfile bool
	Quality         int
	Density         int
	HasTransparency bool
	PHash           string
	Exif            Mexif
}

type Mexif struct {
	Image            ImageExif
	Thumbnail        ThumbnailExif
	Exif             Exif
	Gps              Gps
	Interoperability Interoperability
	Makernote        map[string]interface{}
}

type ImageExif struct {
	Make             string
	Model            string
	Orientation      string
	XResolution      int
	YResolution      int
	ResolutionUnit   int
	Software         string
	ModifyDate       time.Time
	YCbCrPositioning int
	ExifOffset       int
	GPSInfo          int
}

type ThumbnailExif struct {
	Compression     int
	XResolution     int
	YResolution     int
	ResolutionUnit  int
	ThumbnailOffset int
	ThumbnailLength int
}

type Exif struct {
	ExposureTime             time.Time
	FNumber                  float32
	ExposureProgram          int
	ISO                      int
	ExifVersion              string
	DateTimeOriginal         time.Time
	CreateDate               time.Time
	ShutterSpeedValue        float32
	ApertureValue            float32
	ExposureCompensation     int
	MeteringMode             int
	Flash                    int
	FocalLength              int
	SubSEcTime               string
	SubSecTimeOriginal       string
	FlashpixVersion          string
	ColorSpace               int
	ExifImageWidth           int
	ExifImageHeight          int
	InteropOffset            int
	FocalPlaneXResolution    float32
	FocalPlaneYResolution    float32
	FocalPlaneResolutionUnit int
	CustomRendered           int
	ExposureMode             int
	WhiteBalance             int
	SceneCaptutureType       int
}

type Gps struct {
	GPSVersionID []int
}

type Interoperability struct {
	InteropIndex   string
	InteropVersion string
}

type Schema struct {
	Type string `json:"type"` //Text, Textarea, Number, Date, Boolean, SingleSelect, MultiSelect Date value should be an ISO8601 string

	SelectOptions   interface{} `json:"selectOptions,omitempty"`
	DefaultValue    interface{} `json:"defaultValue,omitempty"`
	IsValueRequired bool        `json:"isValueRequired,omitempty"`
	MinValue        interface{} `json:"minValue,omitempty"`
	MaxValue        interface{} `json:"maxValue,omitempty"`
	MinLength       int         `json:"minLength,omitempty"`
	MaxLength       int         `json:"maxLength,omitempty"`
}

type CreateFieldParam struct {
	Name   string `json:"name"`
	Label  string `json:"label"`
	Schema Schema `json:"schema"`
}

type CustomField struct {
	Id     string
	Name   string
	Label  string
	Schema Schema
}

type CreateFieldResponse struct {
	Data CustomField
	api.Response
}

type UpdateCustomFieldResponse CreateFieldResponse

type UpdateCustomFieldParam struct {
	Label  string `json:"label"`
	Schema Schema `json:"schema"`
}

type CustomFieldsResponse struct {
	Data []CustomField
	api.Response
}

func (m *API) get(ctx context.Context, url string, query map[string]string) (*http.Response, error) {
	var err error
	urlObj, err := neturl.Parse(api.BuildPath(m.Config.API.Prefix, url))
	if err != nil {
		return nil, err
	}

	values := urlObj.Query()
	for k, v := range query {
		values.Set(k, v)
	}

	q := values.Encode()

	sUrl := urlObj.String()
	if q != "" {
		sUrl = sUrl + "?" + values.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, sUrl, nil)

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(m.Config.Cloud.PrivateKey, "")

	return m.Client.Do(req.WithContext(ctx))
}

func (m *API) post(ctx context.Context, url string, data interface{}) (*http.Response, error) {
	url = api.BuildPath(m.Config.API.Prefix, url)
	var err error
	var body []byte

	if data != nil {
		if body, err = json.Marshal(data); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(m.Config.Cloud.PrivateKey, "")

	return m.Client.Do(req.WithContext(ctx))
}

func (m *API) patch(ctx context.Context, url string, data interface{}) (*http.Response, error) {
	url = api.BuildPath(m.Config.API.Prefix, url)
	var err error
	var body []byte

	if data != nil {
		if body, err = json.Marshal(data); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(m.Config.Cloud.PrivateKey, "")

	return m.Client.Do(req.WithContext(ctx))
}

func (m *API) delete(ctx context.Context, url string) (*http.Response, error) {
	var err error
	url = api.BuildPath(m.Config.API.Prefix, url)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	req.SetBasicAuth(m.Config.Cloud.PrivateKey, "")

	return m.Client.Do(req.WithContext(ctx))
}

// FromFile fetches metadata of media library file
func (m *API) FromFile(ctx context.Context, fileId string) (*MetadataResponse, error) {
	if fileId == "" {
		return nil, errors.New("fileId can not be blank")
	}

	var response = &MetadataResponse{}

	resp, err := m.get(ctx, fmt.Sprintf("files/%s/metadata", fileId), nil)
	defer api.DeferredBodyClose(resp)

	api.SetResponseMeta(resp, response)

	if err != nil {
		return response, err
	}

	if resp.StatusCode != 200 {
		err = response.ParseError()
	} else {
		err = json.Unmarshal(response.Body(), &response.Data)
	}

	return response, err
}

func (m *API) FromUrl(ctx context.Context, url string) (*MetadataResponse, error) {
	var err error
	if url == "" {
		return nil, errors.New("url can not be blank")
	}

	var response = &MetadataResponse{}

	if err != nil {
		return nil, err
	}

	resp, err := m.get(ctx, "metadata", map[string]string{"url": url})
	defer api.DeferredBodyClose(resp)

	api.SetResponseMeta(resp, response)

	if err != nil {
		return response, err
	}

	if resp.StatusCode != 200 {
		err = response.ParseError()
	} else {
		err = json.Unmarshal(response.Body(), &response.Data)
	}

	return response, err
}

// CreateField creates new custom metadata field
func (m *API) CreateCustomField(ctx context.Context, param CreateFieldParam) (*CreateFieldResponse, error) {
	var err error
	var response = &CreateFieldResponse{}

	resp, err := m.post(ctx, "customMetadataFields", param)
	defer api.DeferredBodyClose(resp)

	api.SetResponseMeta(resp, response)

	if err != nil {
		return response, err
	}

	if resp.StatusCode != 201 {
		err = response.ParseError()
	} else {
		err = json.Unmarshal(response.Body(), &response.Data)
	}

	return response, err
}

// CustomFields returns all existing custom metadata fields
func (m *API) CustomFields(ctx context.Context, includeDeleted bool) (*CustomFieldsResponse, error) {
	var err error
	var response = &CustomFieldsResponse{}
	var flag string

	if includeDeleted == true {
		flag = "true"
	} else {
		flag = "false"
	}

	resp, err := m.get(ctx, "customMetadataFields", map[string]string{"includeDeleted": flag})
	defer api.DeferredBodyClose(resp)

	api.SetResponseMeta(resp, response)

	if err != nil {
		return response, err
	}

	if resp.StatusCode != 200 {
		err = response.ParseError()
	} else {
		err = json.Unmarshal(response.Body(), &response.Data)
	}

	return response, err
}

// UpdateCustomField updates label or schema attributes of given custom field id
func (m *API) UpdateCustomField(ctx context.Context, fieldId string, param UpdateCustomFieldParam) (*UpdateCustomFieldResponse, error) {
	var err error
	var response = &UpdateCustomFieldResponse{}

	if err = validator.Validate(&param); err != nil {
		return nil, err
	}

	resp, err := m.patch(ctx, "customMetadataFields/"+fieldId, param)
	defer api.DeferredBodyClose(resp)

	api.SetResponseMeta(resp, response)

	if err != nil {
		return response, err
	}

	if resp.StatusCode != 200 {
		err = response.ParseError()
	} else {
		err = json.Unmarshal(response.Body(), &response.Data)
	}

	return response, err
}

// DeleteCustomField deletes custom metadata field by given fieldId
func (m *API) DeleteCustomField(ctx context.Context, fieldId string) (*api.Response, error) {
	if fieldId == "" {
		return nil, errors.New("fieldId can not be blank")
	}
	var err error
	var response = &api.Response{}

	resp, err := m.delete(ctx, "customMetadataFields/"+fieldId)
	defer api.DeferredBodyClose(resp)

	api.SetResponseMeta(resp, response)

	if err != nil {
		return response, err
	}

	if resp.StatusCode != 204 {
		err = response.ParseError()
	}

	return response, err
}
