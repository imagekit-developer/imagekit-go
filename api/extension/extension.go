package extension

type TagService string

const (
	GoogleAutoTag TagService = "google-auto-tagging"
	AwsAutoTag    TagService = "aws-auto-tagging"
)

// RemoveBgOptions represents different options for removing bg extension
type RemoveBgOption struct {
	AddShadow        bool   `json:"add_shadow"`
	SemiTransparency bool   `json:"semitransparency"`
	BgColor          string `json:"bg_color,omitempty"`
	BgImageUrl       string `json:"bg_image_url,omitempty"`
}

// AutoTag represents extension struct for auto tagging by providers such as google and aws
type AutoTag struct {
	Name          TagService `json:"name"`
	MinConfidence int        `json:"minConfidence"`
	MaxTags       int        `json:"maxTags,omitempty"`
}

func (e *AutoTag) extName() string {
	return string(e.Name)
}

// NewAutoTag creates an extension parameter for auto tagging
func NewAutoTag(service TagService, minConf int, maxTags int) *AutoTag {
	return &AutoTag{
		Name:          service,
		MinConfidence: minConf,
		MaxTags:       maxTags,
	}
}

// RemoveBg represents extension struct for removing background
type RemoveBg struct {
	Name    string         `json:"name"`
	Options RemoveBgOption `json:"options,omitempty"`
}

func (e *RemoveBg) extName() string {
	return e.Name
}

// NewRemoveBg creates an extension parameter for removing background
func NewRemoveBg(opt RemoveBgOption) *RemoveBg {
	return &RemoveBg{
		Name:    "remove-bg",
		Options: opt,
	}
}

// IExtension prepresents common interface for different extensions.
// This interface is used only as a common type to store different types of extensions in a slice or array.
type IExtension interface {
	extName() string
}
