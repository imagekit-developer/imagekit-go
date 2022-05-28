package url

type trpos string

const (
	PATH  trpos = "path"
	QUERY trpos = "query"
)

type UrlParams struct {
	Path                   string
	Src                    string
	EndpointUrl            string
	Transformation         string
	Signed                 bool
	ExpireSeconds          int64
	TransformationPosition trpos
	QueryParameters        map[string]string
	UnixTime               func() int64
}
