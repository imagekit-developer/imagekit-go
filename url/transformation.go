package url

import (
	"fmt"
	"strings"
)

// AspectRatio represents transformation aspect ratio
type AspectRatio struct {
	Width  int
	Height int
}

// CropMode is transformation crop mode type
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

// Focus is transformation focus type
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

// Format is transformation image format type
type Format string

const (
	FAuto Format = "f-auto"
	FJpg  Format = "f-jpg"
	FJpeg Format = "f-jpeg"
	FWebp Format = "f-webp"
	FAvif Format = "f-avif"
	FPng  Format = "f-png"
)

// Transformation is a struct representing options for transformation parameter to Url()
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
	Dpr              any    // dpr (auto or number 0.1 to 5)
	GrayScale        bool   // e-grayscale
	DefaultImage     string // di (Replaces all forward slashes with @@)
	ProgressiveImage bool   // pr
	Lossless         bool   // lo
	TrimEdges        any    // t (true or number 1-99)
	Border           string // b ("width_hexColorCode")
	ColorProfile     bool   // cp
	ImageMetadata    bool   // md
	Rotate           any    // rt (0 , 90 , 180 , 270 , 360 or auto)
	Radius           any    // r (number or max)
	BgColor          string // bg
	Original         bool   // orig
	Attachment       bool   // ik-attachment
	Contrast         bool
	Sharpen          any // bool or number
	UnsharpMask      UnsharpMask
	Overlay          *Overlay
	Raw              string
}

// UnsharpMask represents type of UnsharpMask option in Transformation
type UnsharpMask struct {
	Radius    float32
	Sigma     float32
	Amount    float32
	Threshold float32
}

// Overlay represents transformation's overlay options
type Overlay struct {
	X          *int         // ox
	Y          *int         // oy
	Height     *int         // oh
	Width      *int         // ow
	Background string       // obg
	Focus      OverlayFocus // ofo
	// Text overlay options
	Text               string // ot
	TextEncoded        string
	TextWidth          *int      // otw
	TextBackground     string    // otbg
	TextPadding        string    // otp ("40" or "40_60" or "40_60_50_10")
	TextInnerAlignment Alignment // otia
	TextColor          string    // otc
	TextFont           string    // otf
	TextSize           *int      // ots
	TextTypography     string    // ott (i, b or ib)=(italic, bold, both)
	Radius             int       // or
	// Image overlay options
	Image            string      // oi (replaces all forward slashes with @@)
	ImageAspectRatio AspectRatio // oiar
	ImageBorder      string      // oib (width_colorcode: 5_red, 10_FF0000)
	ImageDPR         *float32    // oidpr (0.1 to 5)
	ImageQuality     *int        // oiq
	ImageCropping    ImageCropping
	ImageX           *int  //oix
	ImageY           *int  //oiy
	ImageXc          *int  //oixc
	ImageYc          *int  //oixy
	Trim             *bool //oit-true
}

type Alignment string

const (
	Left   Alignment = "left"
	Right  Alignment = "right"
	Center Alignment = "center"
)

// ImageCropping is image overlay cropping type
type ImageCropping string

const (
	CropAtLeast       ImageCropping = "oic-at_lease"
	CropAtMax         ImageCropping = "oic-at_max"
	CropForce         ImageCropping = "oic-at_lease"
	CropMaintainRatio ImageCropping = "oic-maintain_ratio"
	CropExtract       ImageCropping = "oicm-extract"
	CropPadExtract    ImageCropping = "oicm-pad_extract"
	CropPadResize     ImageCropping = "oic-pad_resize"
	OifoAuto          ImageCropping = "oifo-auto"
	OifoCustom        ImageCropping = "oifo-custom"
	OifoFace          ImageCropping = "oifo-face"
)

type OverlayFocus string

const (
	OfCenter      OverlayFocus = "ofo-center"
	OfTop         OverlayFocus = "ofo-top"
	OfLeft        OverlayFocus = "ofo-left"
	OfBottom      OverlayFocus = "ofo-bottom"
	OfRight       OverlayFocus = "ofo-right"
	OfTopRight    OverlayFocus = "ofo-top_right"
	OfTopLeft     OverlayFocus = "ofo-top_left"
	OfBottomLeft  OverlayFocus = "ofo-bottom_left"
	OfBottomRight OverlayFocus = "ofo-bottom_right"
)

// Overlay stringer
func (o *Overlay) String() string {
	var parts []string

	if o.Text != "" {
		parts = append(parts, "ot-"+o.Text)
	} else if o.TextEncoded != "" {
		parts = append(parts, "ote-"+o.TextEncoded)
	} else if o.Image != "" {
		parts = append(parts, "oi-"+strings.ReplaceAll(strings.Trim(o.Image, "/"), "/", "@@"))
	}

	appendVar(&parts, o.X, "ox")
	appendVar(&parts, o.Y, "oy")
	appendVar(&parts, o.Height, "oh")
	appendVar(&parts, o.Width, "ow")
	appendVar(&parts, o.TextWidth, "otw")

	if o.TextBackground != "" {
		parts = append(parts, "otbg-"+o.TextBackground)
	}

	if o.TextPadding != "" {
		parts = append(parts, "otp-"+o.TextPadding)
	}

	if o.TextInnerAlignment != "" {
		parts = append(parts, "otia-"+string(o.TextInnerAlignment))
	}

	if o.TextColor != "" {
		parts = append(parts, "otc-"+o.TextColor)
	}

	if o.TextFont != "" {
		parts = append(parts, "otf-"+o.TextFont)
	}

	appendVar(&parts, o.TextSize, "ots")

	if o.TextTypography != "" {
		parts = append(parts, "ott-"+string(o.TextTypography))
	}

	if o.Radius > 0 {
		parts = append(parts, fmt.Sprintf("or-%d", o.Radius))
	}

	if o.Background != "" {
		parts = append(parts, "obg-"+o.Background)
	}

	if o.Focus != "" {
		parts = append(parts, string(o.Focus))
	}

	ar := o.ImageAspectRatio
	if ar.Width > 0 && ar.Height > 0 {
		parts = append(parts, fmt.Sprintf("oiar-%d-%d", ar.Width, ar.Height))
	}

	if o.ImageBorder != "" {
		parts = append(parts, "oib-"+o.ImageBorder)
	}

	appendVar(&parts, o.ImageDPR, "oidpr")
	appendVar(&parts, o.ImageQuality, "oiq")

	if o.ImageCropping != "" {
		parts = append(parts, string(o.ImageCropping))
	}

	appendVar(&parts, o.ImageX, "oix")
	appendVar(&parts, o.ImageY, "oiy")
	appendVar(&parts, o.ImageXc, "oixc")
	appendVar(&parts, o.ImageYc, "oiyc")
	appendVar(&parts, o.Trim, "oit")

	return strings.Join(parts, ",")
}

func appendVar[T int | float32 | bool](parts *[]string, v *T, prefix string) {
	if v != nil {
		*parts = append(*parts, prefix+"-"+fmt.Sprintf("%v", *v))
	}
}

func appendBool(parts *[]string, v bool, prefix string) {
	if v == true {
		*parts = append(*parts, prefix+"-"+fmt.Sprintf("%v", v))
	}
}

// Transformation stringer
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
		parts = append(parts, fmt.Sprintf("x-%d,y-%d", t.X, t.Y))
	} else if t.Xc > 0 && t.Yc > 0 {
		parts = append(parts, fmt.Sprintf("xc-%d,yc-%d", t.Xc, t.Yc))
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
		parts = append(parts, "di-"+strings.ReplaceAll(strings.Trim(t.DefaultImage, "/"), "/", "@@"))
	}

	appendBool(&parts, t.ProgressiveImage, "pr")
	appendBool(&parts, t.Lossless, "lo")

	if t.TrimEdges != nil {
		parts = append(parts, "t-"+fmt.Sprintf("%v", t.TrimEdges))
	}

	if t.Border != "" {
		parts = append(parts, "b-"+t.Border)
	}

	appendBool(&parts, t.ColorProfile, "cp")

	appendBool(&parts, t.ImageMetadata, "md")

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

	switch v := t.Sharpen.(type) {
	case bool:
		parts = append(parts, "e-sharpen")
	case int:
		parts = append(parts, fmt.Sprintf("e-sharpen-%d", v))
	}

	if t.Contrast {
		parts = append(parts, "e-contrast")
	}

	if usm := t.UnsharpMask; (UnsharpMask{} != usm) {
		parts = append(parts, fmt.Sprintf("e-usm-%-.4f-%.4f-%.4f-%.4f",
			usm.Radius, usm.Sigma, usm.Amount, usm.Threshold))
	}

	if t.Overlay != nil {
		parts = append(parts, t.Overlay.String())
	}

	if t.Raw != "" {
		parts = append(parts, t.Raw)
	}
	return strings.Join(parts, ",")

}
