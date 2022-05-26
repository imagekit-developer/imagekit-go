package url

type UrlParams struct {
	Path           string
	Src            string
	UrlEndpoint    string
	Transformation string
	Signed         bool
	ExpireSeconds  int
}