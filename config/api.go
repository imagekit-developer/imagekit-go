package config

// API defines the configuration for making requests to the ImageKit.io API.
type API struct {
	UploadPrefix  string `default:"https://upload.imagekit.io/api/v1/"`
	Timeout       int64  `default:"60"` // seconds
	UploadTimeout int64  `upload_timeout"`
}
