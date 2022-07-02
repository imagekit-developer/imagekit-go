package url

type trpos string

const (
	PATH  trpos = "path"
	QUERY trpos = "query"
)

type UrlParam struct {
	Path                string
	Src                 string
	EndpointUrl         string
	Transformations     []Transformation // tr:w-400,h-300
	NamedTransformation string           // n-trname

	Signed                 bool
	ExpireSeconds          int64
	TransformationPosition trpos
	QueryParameters        map[string]string
	UnixTime               func() int64
}
