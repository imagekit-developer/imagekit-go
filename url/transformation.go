package url

import (
	"fmt"
	"strings"
)

type AspectRatio struct {
	Width  int
	Height int
}

type CropMode string

const (
	CmPadResize     CropMode = "cm-pad_resize"
	CmExtract       CropMode = "cm-extract"
	CmPadExtract    CropMode = "cm-pad_extract"
	CmForcedCrop    CropMode = "c-force"
	CmAtMax         CropMode = "c-at_max"
	CmAtLeast       CropMode = "c-at_least"
	CmMaintainRatio CropMode = "c-maintain_ratio"
)

type Focus string

const (
	FoLeft        Focus = "fo-left"
	FoRight       Focus = "fo-right"
	FoTop         Focus = "fo-top"
	FoBottom      Focus = "fo-bottom"
	FoCustom      Focus = "fo-custom"
	FoCenter      Focus = "fo-center"
	FoTopLeft     Focus = "fo-top_left"
	FoTopRight    Focus = "fo-top_right"
	FoBottomLeft  Focus = "fo-bottom_left"
	FoBottomRight Focus = "fo-bototm_right"
	FoAuto        Focus = "fo-auto"
	FoFace        Focus = "fo-face"
)

type Format string

const (
	FAuto Format = "f-auto"
	FJpg  Format = "f-jpg"
	FJpeg Format = "f-jpeg"
	FWebp Format = "f-webp"
	FAvif Format = "f-avif"
	FPng  Format = "f-png"
)

type Border struct {
	Width   int
	HexCode string
}

type Transformation struct {
	Width            float32     // w
	Height           float32     // h
	AspectRatio      AspectRatio // ar
	CropMode         CropMode
	Focus            Focus  // fo-custom|face|auto|left|right|top|bottom|top_left.....
	X                int    // x
	Y                int    // y
	Xc               int    // xc.  Focus with center coordinate Xc and Yc
	Yc               int    // yc
	Quality          int    // q
	Format           Format // f
	Blur             int    // bl
	Dpr              any    // dpr
	GrayScale        bool   // e-grayscale
	DefaultImage     string // di
	ProgressiveImage bool   // pr
	Lossless         bool   // lo
	TrimEdges        any    // t
	Border           Border // b
	ColorProfile     bool   // cp
	ImageMetadata    bool   // md
	Rotate           any    // rt
	Radius           any    // r
	BgColor          string // bg
	Original         bool   // orig
	Attachment       bool   // ik-attachment
}

func appendBool(parts *[]string, v bool, prefix string) {
	if v == true {
		*parts = append(*parts, prefix+fmt.Sprintf("%v", v))
	}
}

func (t Transformation) String() string {
	var parts []string

	if t.Width > 0 {
		parts = append(parts, "w-"+fmt.Sprintf("%v", t.Width))
	}

	if t.Height > 0 {
		parts = append(parts, "h-"+fmt.Sprintf("%v", t.Height))
	}

	if t.AspectRatio.Height > 0 && t.AspectRatio.Width > 0 {
		parts = append(parts, fmt.Sprintf("ar-%d-%d", t.AspectRatio.Width, t.AspectRatio.Height))
	}

	if t.CropMode != "" {
		parts = append(parts, string(t.CropMode))
	}

	if t.Focus != "" {
		parts = append(parts, string(t.Focus))
	}

	if t.X > 0 && t.Y > 0 {
		parts = append(parts, fmt.Sprintf("x-%d,y=%d", t.X, t.Y))
	} else if t.Xc > 0 && t.Yc > 0 {
		parts = append(parts, fmt.Sprintf("xc-%d,yc=%d", t.X, t.Y))
	}

	if t.Quality > 0 {
		parts = append(parts, fmt.Sprintf("q-%d", t.Quality))
	}

	if t.Format != "" {
		parts = append(parts, string(t.Format))
	}

	if t.Blur > 0 {
		parts = append(parts, fmt.Sprintf("bl-%d", t.Blur))
	}

	if t.Dpr != nil {
		parts = append(parts, fmt.Sprintf("dpr-%v", t.Dpr))
	}

	if t.GrayScale == true {
		parts = append(parts, "e-grayscale")
	}

	if t.DefaultImage != "" {
		parts = append(parts, "di-"+t.DefaultImage)
	}

	appendBool(&parts, t.ProgressiveImage, "pr-")
	//parts = append(parts, "pr-"+fmt.Sprintf("%v", t.ProgressiveImage))
	appendBool(&parts, t.Lossless, "lo-")
	//parts = append(parts, "lo-"+fmt.Sprintf("%v", t.Lossless))

	if t.TrimEdges != nil {
		parts = append(parts, "t-"+fmt.Sprintf("%v", t.TrimEdges))
	}

	if t.Border.Width > 0 {
		parts = append(parts, fmt.Sprintf("b-%d-%s", t.Border.Width, t.Border.HexCode))
	}

	appendBool(&parts, t.ColorProfile, "cp-")
	//parts = append(parts, "cp-"+fmt.Sprintf("%v", t.ColorProfile))

	appendBool(&parts, t.ImageMetadata, "md-")
	//	parts = append(parts, "md-"+fmt.Sprintf("%v", t.ImageMetadata))

	if t.Rotate != nil {
		parts = append(parts, "rt-"+fmt.Sprintf("%v", t.Rotate))
	}

	if t.Radius != nil {
		parts = append(parts, "r-"+fmt.Sprintf("%v", t.Radius))
	}

	if t.BgColor != "" {
		parts = append(parts, "bg-"+t.BgColor)
	}

	if t.Original {
		parts = append(parts, "orig-"+fmt.Sprintf("%v", t.Original))
	}

	if t.Attachment {
		parts = append(parts, "ik-attachment="+fmt.Sprintf("%v", t.Attachment))
	}

	return strings.Join(parts, ",")
}
