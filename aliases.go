// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit

import (
	"github.com/imagekit-developer/imagekit-go/v2/internal/apierror"
	"github.com/imagekit-developer/imagekit-go/v2/packages/param"
	"github.com/imagekit-developer/imagekit-go/v2/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// This is an alias to an internal type.
type BaseOverlayParam = shared.BaseOverlayParam

// Array of extensions to be applied to the asset. Each extension can be configured
// with specific parameters based on the extension type.
//
// This is an alias to an internal type.
type ExtensionsParam = shared.ExtensionsParam

// This is an alias to an internal type.
type ExtensionUnionParam = shared.ExtensionUnionParam

// This is an alias to an internal type.
type ExtensionRemoveBgParam = shared.ExtensionRemoveBgParam

// This is an alias to an internal type.
type ExtensionRemoveBgOptionsParam = shared.ExtensionRemoveBgOptionsParam

// This is an alias to an internal type.
type ExtensionAutoTaggingParam = shared.ExtensionAutoTaggingParam

// This is an alias to an internal type.
type ExtensionAIAutoDescriptionParam = shared.ExtensionAIAutoDescriptionParam

// Options for generating responsive image attributes including `src`, `srcSet`,
// and `sizes` for HTML `<img>` elements. This schema extends `SrcOptions` to add
// support for responsive image generation with breakpoints.
//
// This is an alias to an internal type.
type GetImageAttributesOptionsParam = shared.GetImageAttributesOptionsParam

// This is an alias to an internal type.
type ImageOverlayParam = shared.ImageOverlayParam

// Specifies an overlay to be applied on the parent image or video. ImageKit
// supports overlays including images, text, videos, subtitles, and solid colors.
// See
// [Overlay using layers](https://imagekit.io/docs/transformations#overlay-using-layers).
//
// This is an alias to an internal type.
type OverlayUnionParam = shared.OverlayUnionParam

// This is an alias to an internal type.
type OverlayPositionParam = shared.OverlayPositionParam

// Specifies the position of the overlay relative to the parent image or video.
// Maps to `lfo` in the URL.
//
// This is an alias to an internal type.
type OverlayPositionFocus = shared.OverlayPositionFocus

// Equals "center"
const OverlayPositionFocusCenter = shared.OverlayPositionFocusCenter

// Equals "top"
const OverlayPositionFocusTop = shared.OverlayPositionFocusTop

// Equals "left"
const OverlayPositionFocusLeft = shared.OverlayPositionFocusLeft

// Equals "bottom"
const OverlayPositionFocusBottom = shared.OverlayPositionFocusBottom

// Equals "right"
const OverlayPositionFocusRight = shared.OverlayPositionFocusRight

// Equals "top_left"
const OverlayPositionFocusTopLeft = shared.OverlayPositionFocusTopLeft

// Equals "top_right"
const OverlayPositionFocusTopRight = shared.OverlayPositionFocusTopRight

// Equals "bottom_left"
const OverlayPositionFocusBottomLeft = shared.OverlayPositionFocusBottomLeft

// Equals "bottom_right"
const OverlayPositionFocusBottomRight = shared.OverlayPositionFocusBottomRight

// Specifies the x-coordinate of the top-left corner of the base asset where the
// overlay's top-left corner will be positioned. It also accepts arithmetic
// expressions such as `bw_mul_0.4` or `bw_sub_cw`. Maps to `lx` in the URL. Learn
// about
// [Arithmetic expressions](https://imagekit.io/docs/arithmetic-expressions-in-transformations).
//
// This is an alias to an internal type.
type OverlayPositionXUnionParam = shared.OverlayPositionXUnionParam

// Specifies the y-coordinate of the top-left corner of the base asset where the
// overlay's top-left corner will be positioned. It also accepts arithmetic
// expressions such as `bh_mul_0.4` or `bh_sub_ch`. Maps to `ly` in the URL. Learn
// about
// [Arithmetic expressions](https://imagekit.io/docs/arithmetic-expressions-in-transformations).
//
// This is an alias to an internal type.
type OverlayPositionYUnionParam = shared.OverlayPositionYUnionParam

// This is an alias to an internal type.
type OverlayTimingParam = shared.OverlayTimingParam

// Specifies the duration (in seconds) during which the overlay should appear on
// the base video. Accepts a positive number up to two decimal places (e.g., `20`
// or `20.50`) and arithmetic expressions such as `bdu_mul_0.4` or `bdu_sub_idu`.
// Applies only if the base asset is a video. Maps to `ldu` in the URL.
//
// This is an alias to an internal type.
type OverlayTimingDurationUnionParam = shared.OverlayTimingDurationUnionParam

// Specifies the end time (in seconds) for when the overlay should disappear from
// the base video. If both end and duration are provided, duration is ignored.
// Accepts a positive number up to two decimal places (e.g., `20` or `20.50`) and
// arithmetic expressions such as `bdu_mul_0.4` or `bdu_sub_idu`. Applies only if
// the base asset is a video. Maps to `leo` in the URL.
//
// This is an alias to an internal type.
type OverlayTimingEndUnionParam = shared.OverlayTimingEndUnionParam

// Specifies the start time (in seconds) for when the overlay should appear on the
// base video. Accepts a positive number up to two decimal places (e.g., `20` or
// `20.50`) and arithmetic expressions such as `bdu_mul_0.4` or `bdu_sub_idu`.
// Applies only if the base asset is a video. Maps to `lso` in the URL.
//
// This is an alias to an internal type.
type OverlayTimingStartUnionParam = shared.OverlayTimingStartUnionParam

// Resulting set of attributes suitable for an HTML `<img>` element. Useful for
// enabling responsive image loading with `srcSet` and `sizes`.
//
// This is an alias to an internal type.
type ResponsiveImageAttributesParam = shared.ResponsiveImageAttributesParam

// This is an alias to an internal type.
type SolidColorOverlayParam = shared.SolidColorOverlayParam

// This is an alias to an internal type.
type SolidColorOverlayTransformationParam = shared.SolidColorOverlayTransformationParam

// Creates a linear gradient with two colors. Pass `true` for a default gradient,
// or provide a string for a custom gradient. Only works if the base asset is an
// image. See
// [gradient](https://imagekit.io/docs/effects-and-enhancements#gradient---e-gradient).
//
// This is an alias to an internal type.
type SolidColorOverlayTransformationGradientUnionParam = shared.SolidColorOverlayTransformationGradientUnionParam

// This is an alias to an internal type.
type SolidColorOverlayTransformationGradientBoolean = shared.SolidColorOverlayTransformationGradientBoolean

// Controls the height of the solid color overlay. Accepts a numeric value or an
// arithmetic expression. Learn about
// [arithmetic expressions](https://imagekit.io/docs/arithmetic-expressions-in-transformations).
//
// This is an alias to an internal type.
type SolidColorOverlayTransformationHeightUnionParam = shared.SolidColorOverlayTransformationHeightUnionParam

// Specifies the corner radius of the solid color overlay. Set to `max` for
// circular or oval shape. See
// [radius](https://imagekit.io/docs/effects-and-enhancements#radius---r).
//
// This is an alias to an internal type.
type SolidColorOverlayTransformationRadiusUnionParam = shared.SolidColorOverlayTransformationRadiusUnionParam

// Controls the width of the solid color overlay. Accepts a numeric value or an
// arithmetic expression (e.g., `bw_mul_0.2` or `bh_div_2`). Learn about
// [arithmetic expressions](https://imagekit.io/docs/arithmetic-expressions-in-transformations).
//
// This is an alias to an internal type.
type SolidColorOverlayTransformationWidthUnionParam = shared.SolidColorOverlayTransformationWidthUnionParam

// Options for generating ImageKit URLs with transformations. See the
// [Transformations guide](https://imagekit.io/docs/transformations).
//
// This is an alias to an internal type.
type SrcOptionsParam = shared.SrcOptionsParam

// Available streaming resolutions for
// [adaptive bitrate streaming](https://imagekit.io/docs/adaptive-bitrate-streaming)
//
// This is an alias to an internal type.
type StreamingResolution = shared.StreamingResolution

// Equals "240"
const StreamingResolution240 = shared.StreamingResolution240

// Equals "360"
const StreamingResolution360 = shared.StreamingResolution360

// Equals "480"
const StreamingResolution480 = shared.StreamingResolution480

// Equals "720"
const StreamingResolution720 = shared.StreamingResolution720

// Equals "1080"
const StreamingResolution1080 = shared.StreamingResolution1080

// Equals "1440"
const StreamingResolution1440 = shared.StreamingResolution1440

// Equals "2160"
const StreamingResolution2160 = shared.StreamingResolution2160

// This is an alias to an internal type.
type SubtitleOverlayParam = shared.SubtitleOverlayParam

// Subtitle styling options.
// [Learn more](https://imagekit.io/docs/add-overlays-on-videos#styling-controls-for-subtitles-layer)
// from the docs.
//
// This is an alias to an internal type.
type SubtitleOverlayTransformationParam = shared.SubtitleOverlayTransformationParam

// Sets the typography style of the subtitle text. Supports values are `b` for
// bold, `i` for italics, and `b_i` for bold with italics.
//
// [Subtitle styling options](https://imagekit.io/docs/add-overlays-on-videos#styling-controls-for-subtitles-layer)
//
// This is an alias to an internal type.
type SubtitleOverlayTransformationTypography = shared.SubtitleOverlayTransformationTypography

// Equals "b"
const SubtitleOverlayTransformationTypographyB = shared.SubtitleOverlayTransformationTypographyB

// Equals "i"
const SubtitleOverlayTransformationTypographyI = shared.SubtitleOverlayTransformationTypographyI

// Equals "b_i"
const SubtitleOverlayTransformationTypographyBI = shared.SubtitleOverlayTransformationTypographyBI

// This is an alias to an internal type.
type TextOverlayParam = shared.TextOverlayParam

// This is an alias to an internal type.
type TextOverlayTransformationParam = shared.TextOverlayTransformationParam

// Flip the text overlay horizontally, vertically, or both.
//
// This is an alias to an internal type.
type TextOverlayTransformationFlip = shared.TextOverlayTransformationFlip

// Equals "h"
const TextOverlayTransformationFlipH = shared.TextOverlayTransformationFlipH

// Equals "v"
const TextOverlayTransformationFlipV = shared.TextOverlayTransformationFlipV

// Equals "h_v"
const TextOverlayTransformationFlipHV = shared.TextOverlayTransformationFlipHV

// Equals "v_h"
const TextOverlayTransformationFlipVH = shared.TextOverlayTransformationFlipVH

// Specifies the font size of the overlaid text. Accepts a numeric value or an
// arithmetic expression.
//
// This is an alias to an internal type.
type TextOverlayTransformationFontSizeUnionParam = shared.TextOverlayTransformationFontSizeUnionParam

// Specifies the inner alignment of the text when width is more than the text
// length.
//
// This is an alias to an internal type.
type TextOverlayTransformationInnerAlignment = shared.TextOverlayTransformationInnerAlignment

// Equals "left"
const TextOverlayTransformationInnerAlignmentLeft = shared.TextOverlayTransformationInnerAlignmentLeft

// Equals "right"
const TextOverlayTransformationInnerAlignmentRight = shared.TextOverlayTransformationInnerAlignmentRight

// Equals "center"
const TextOverlayTransformationInnerAlignmentCenter = shared.TextOverlayTransformationInnerAlignmentCenter

// Specifies the line height of the text overlay. Accepts integer values
// representing line height in points. It can also accept
// [arithmetic expressions](https://imagekit.io/docs/arithmetic-expressions-in-transformations)
// such as `bw_mul_0.2`, or `bh_div_20`.
//
// This is an alias to an internal type.
type TextOverlayTransformationLineHeightUnionParam = shared.TextOverlayTransformationLineHeightUnionParam

// Specifies the padding around the overlaid text. Can be provided as a single
// positive integer or multiple values separated by underscores (following CSS
// shorthand order). Arithmetic expressions are also accepted.
//
// This is an alias to an internal type.
type TextOverlayTransformationPaddingUnionParam = shared.TextOverlayTransformationPaddingUnionParam

// Specifies the corner radius of the text overlay. Set to `max` to achieve a
// circular or oval shape.
//
// This is an alias to an internal type.
type TextOverlayTransformationRadiusUnionParam = shared.TextOverlayTransformationRadiusUnionParam

// Specifies the rotation angle of the text overlay. Accepts a numeric value for
// clockwise rotation or a string prefixed with "N" for counter-clockwise rotation.
//
// This is an alias to an internal type.
type TextOverlayTransformationRotationUnionParam = shared.TextOverlayTransformationRotationUnionParam

// Specifies the maximum width (in pixels) of the overlaid text. The text wraps
// automatically, and arithmetic expressions (e.g., `bw_mul_0.2` or `bh_div_2`) are
// supported. Useful when used in conjunction with the `background`. Learn about
// [Arithmetic expressions](https://imagekit.io/docs/arithmetic-expressions-in-transformations).
//
// This is an alias to an internal type.
type TextOverlayTransformationWidthUnionParam = shared.TextOverlayTransformationWidthUnionParam

// The SDK provides easy-to-use names for transformations. These names are
// converted to the corresponding transformation string before being added to the
// URL. SDKs are updated regularly to support new transformations. If you want to
// use a transformation that is not supported by the SDK, You can use the `raw`
// parameter to pass the transformation string directly. See the
// [Transformations documentation](https://imagekit.io/docs/transformations).
//
// This is an alias to an internal type.
type TransformationParam = shared.TransformationParam

// Adds an AI-based drop shadow around a foreground object on a transparent or
// removed background. Optionally, control the direction, elevation, and saturation
// of the light source (e.g., `az-45` to change light direction). Pass `true` for
// the default drop shadow, or provide a string for a custom drop shadow. Supported
// inside overlay. See
// [AI Drop Shadow](https://imagekit.io/docs/ai-transformations#ai-drop-shadow-e-dropshadow).
//
// This is an alias to an internal type.
type TransformationAIDropShadowUnionParam = shared.TransformationAIDropShadowUnionParam

// This is an alias to an internal type.
type TransformationAIDropShadowBoolean = shared.TransformationAIDropShadowBoolean

// Specifies the aspect ratio for the output, e.g., "ar-4-3". Typically used with
// either width or height (but not both). For example: aspectRatio = `4:3`, `4_3`,
// or an expression like `iar_div_2`. See
// [Image resize and crop – Aspect ratio](https://imagekit.io/docs/image-resize-and-crop#aspect-ratio---ar).
//
// This is an alias to an internal type.
type TransformationAspectRatioUnionParam = shared.TransformationAspectRatioUnionParam

// Specifies the audio codec, e.g., `aac`, `opus`, or `none`. See
// [Audio codec](https://imagekit.io/docs/video-optimization#audio-codec---ac).
//
// This is an alias to an internal type.
type TransformationAudioCodec = shared.TransformationAudioCodec

// Equals "aac"
const TransformationAudioCodecAac = shared.TransformationAudioCodecAac

// Equals "opus"
const TransformationAudioCodecOpus = shared.TransformationAudioCodecOpus

// Equals "none"
const TransformationAudioCodecNone = shared.TransformationAudioCodecNone

// Crop modes for image resizing. See
// [Crop modes & focus](https://imagekit.io/docs/image-resize-and-crop#crop-crop-modes--focus).
//
// This is an alias to an internal type.
type TransformationCrop = shared.TransformationCrop

// Equals "force"
const TransformationCropForce = shared.TransformationCropForce

// Equals "at_max"
const TransformationCropAtMax = shared.TransformationCropAtMax

// Equals "at_max_enlarge"
const TransformationCropAtMaxEnlarge = shared.TransformationCropAtMaxEnlarge

// Equals "at_least"
const TransformationCropAtLeast = shared.TransformationCropAtLeast

// Equals "maintain_ratio"
const TransformationCropMaintainRatio = shared.TransformationCropMaintainRatio

// Additional crop modes for image resizing. See
// [Crop modes & focus](https://imagekit.io/docs/image-resize-and-crop#crop-crop-modes--focus).
//
// This is an alias to an internal type.
type TransformationCropMode = shared.TransformationCropMode

// Equals "pad_resize"
const TransformationCropModePadResize = shared.TransformationCropModePadResize

// Equals "extract"
const TransformationCropModeExtract = shared.TransformationCropModeExtract

// Equals "pad_extract"
const TransformationCropModePadExtract = shared.TransformationCropModePadExtract

// Specifies the duration (in seconds) for trimming videos, e.g., `5` or `10.5`.
// Typically used with startOffset to indicate the length from the start offset.
// Arithmetic expressions are supported. See
// [Trim videos – Duration](https://imagekit.io/docs/trim-videos#duration---du).
//
// This is an alias to an internal type.
type TransformationDurationUnionParam = shared.TransformationDurationUnionParam

// Specifies the end offset (in seconds) for trimming videos, e.g., `5` or `10.5`.
// Typically used with startOffset to define a time window. Arithmetic expressions
// are supported. See
// [Trim videos – End offset](https://imagekit.io/docs/trim-videos#end-offset---eo).
//
// This is an alias to an internal type.
type TransformationEndOffsetUnionParam = shared.TransformationEndOffsetUnionParam

// Flips or mirrors an image either horizontally, vertically, or both. Acceptable
// values: `h` (horizontal), `v` (vertical), `h_v` (horizontal and vertical), or
// `v_h`. See [Flip](https://imagekit.io/docs/effects-and-enhancements#flip---fl).
//
// This is an alias to an internal type.
type TransformationFlip = shared.TransformationFlip

// Equals "h"
const TransformationFlipH = shared.TransformationFlipH

// Equals "v"
const TransformationFlipV = shared.TransformationFlipV

// Equals "h_v"
const TransformationFlipHV = shared.TransformationFlipHV

// Equals "v_h"
const TransformationFlipVH = shared.TransformationFlipVH

// Specifies the output format for images or videos, e.g., `jpg`, `png`, `webp`,
// `mp4`, or `auto`. You can also pass `orig` for images to return the original
// format. ImageKit automatically delivers images and videos in the optimal format
// based on device support unless overridden by the dashboard settings or the
// format parameter. See
// [Image format](https://imagekit.io/docs/image-optimization#format---f) and
// [Video format](https://imagekit.io/docs/video-optimization#format---f).
//
// This is an alias to an internal type.
type TransformationFormat = shared.TransformationFormat

// Equals "auto"
const TransformationFormatAuto = shared.TransformationFormatAuto

// Equals "webp"
const TransformationFormatWebp = shared.TransformationFormatWebp

// Equals "jpg"
const TransformationFormatJpg = shared.TransformationFormatJpg

// Equals "jpeg"
const TransformationFormatJpeg = shared.TransformationFormatJpeg

// Equals "png"
const TransformationFormatPng = shared.TransformationFormatPng

// Equals "gif"
const TransformationFormatGif = shared.TransformationFormatGif

// Equals "svg"
const TransformationFormatSvg = shared.TransformationFormatSvg

// Equals "mp4"
const TransformationFormatMP4 = shared.TransformationFormatMP4

// Equals "webm"
const TransformationFormatWebm = shared.TransformationFormatWebm

// Equals "avif"
const TransformationFormatAvif = shared.TransformationFormatAvif

// Equals "orig"
const TransformationFormatOrig = shared.TransformationFormatOrig

// Creates a linear gradient with two colors. Pass `true` for a default gradient,
// or provide a string for a custom gradient. See
// [Gradient](https://imagekit.io/docs/effects-and-enhancements#gradient---e-gradient).
//
// This is an alias to an internal type.
type TransformationGradientUnionParam = shared.TransformationGradientUnionParam

// This is an alias to an internal type.
type TransformationGradientBoolean = shared.TransformationGradientBoolean

// Specifies the height of the output. If a value between 0 and 1 is provided, it
// is treated as a percentage (e.g., `0.5` represents 50% of the original height).
// You can also supply arithmetic expressions (e.g., `ih_mul_0.5`). Height
// transformation –
// [Images](https://imagekit.io/docs/image-resize-and-crop#height---h) ·
// [Videos](https://imagekit.io/docs/video-resize-and-crop#height---h)
//
// This is an alias to an internal type.
type TransformationHeightUnionParam = shared.TransformationHeightUnionParam

// Extracts a specific page or frame from multi-page or layered files (PDF, PSD,
// AI). For example, specify by number (e.g., `2`), a range (e.g., `3-4` for the
// 2nd and 3rd layers), or by name (e.g., `name-layer-4` for a PSD layer). See
// [Thumbnail extraction](https://imagekit.io/docs/vector-and-animated-images#get-thumbnail-from-psd-pdf-ai-eps-and-animated-files).
//
// This is an alias to an internal type.
type TransformationPageUnionParam = shared.TransformationPageUnionParam

// Specifies the corner radius for rounded corners (e.g., 20) or `max` for circular
// or oval shape. See
// [Radius](https://imagekit.io/docs/effects-and-enhancements#radius---r).
//
// This is an alias to an internal type.
type TransformationRadiusUnionParam = shared.TransformationRadiusUnionParam

// Specifies the rotation angle in degrees. Positive values rotate the image
// clockwise; you can also use, for example, `N40` for counterclockwise rotation or
// `auto` to use the orientation specified in the image's EXIF data. For videos,
// only the following values are supported: 0, 90, 180, 270, or 360. See
// [Rotate](https://imagekit.io/docs/effects-and-enhancements#rotate---rt).
//
// This is an alias to an internal type.
type TransformationRotationUnionParam = shared.TransformationRotationUnionParam

// Adds a shadow beneath solid objects in an image with a transparent background.
// For AI-based drop shadows, refer to aiDropShadow. Pass `true` for a default
// shadow, or provide a string for a custom shadow. See
// [Shadow](https://imagekit.io/docs/effects-and-enhancements#shadow---e-shadow).
//
// This is an alias to an internal type.
type TransformationShadowUnionParam = shared.TransformationShadowUnionParam

// This is an alias to an internal type.
type TransformationShadowBoolean = shared.TransformationShadowBoolean

// Sharpens the input image, highlighting edges and finer details. Pass `true` for
// default sharpening, or provide a numeric value for custom sharpening. See
// [Sharpen](https://imagekit.io/docs/effects-and-enhancements#sharpen---e-sharpen).
//
// This is an alias to an internal type.
type TransformationSharpenUnionParam = shared.TransformationSharpenUnionParam

// This is an alias to an internal type.
type TransformationSharpenBoolean = shared.TransformationSharpenBoolean

// Specifies the start offset (in seconds) for trimming videos, e.g., `5` or
// `10.5`. Arithmetic expressions are also supported. See
// [Trim videos – Start offset](https://imagekit.io/docs/trim-videos#start-offset---so).
//
// This is an alias to an internal type.
type TransformationStartOffsetUnionParam = shared.TransformationStartOffsetUnionParam

// Useful for images with a solid or nearly solid background and a central object.
// This parameter trims the background, leaving only the central object in the
// output image. See
// [Trim edges](https://imagekit.io/docs/effects-and-enhancements#trim-edges---t).
//
// This is an alias to an internal type.
type TransformationTrimUnionParam = shared.TransformationTrimUnionParam

// This is an alias to an internal type.
type TransformationTrimBoolean = shared.TransformationTrimBoolean

// Applies Unsharp Masking (USM), an image sharpening technique. Pass `true` for a
// default unsharp mask, or provide a string for a custom unsharp mask. See
// [Unsharp Mask](https://imagekit.io/docs/effects-and-enhancements#unsharp-mask---e-usm).
//
// This is an alias to an internal type.
type TransformationUnsharpMaskUnionParam = shared.TransformationUnsharpMaskUnionParam

// This is an alias to an internal type.
type TransformationUnsharpMaskBoolean = shared.TransformationUnsharpMaskBoolean

// Specifies the video codec, e.g., `h264`, `vp9`, `av1`, or `none`. See
// [Video codec](https://imagekit.io/docs/video-optimization#video-codec---vc).
//
// This is an alias to an internal type.
type TransformationVideoCodec = shared.TransformationVideoCodec

// Equals "h264"
const TransformationVideoCodecH264 = shared.TransformationVideoCodecH264

// Equals "vp9"
const TransformationVideoCodecVp9 = shared.TransformationVideoCodecVp9

// Equals "av1"
const TransformationVideoCodecAv1 = shared.TransformationVideoCodecAv1

// Equals "none"
const TransformationVideoCodecNone = shared.TransformationVideoCodecNone

// Specifies the width of the output. If a value between 0 and 1 is provided, it is
// treated as a percentage (e.g., `0.4` represents 40% of the original width). You
// can also supply arithmetic expressions (e.g., `iw_div_2`). Width transformation
// – [Images](https://imagekit.io/docs/image-resize-and-crop#width---w) ·
// [Videos](https://imagekit.io/docs/video-resize-and-crop#width---w)
//
// This is an alias to an internal type.
type TransformationWidthUnionParam = shared.TransformationWidthUnionParam

// Focus using cropped image coordinates - X coordinate. See
// [Focus using cropped coordinates](https://imagekit.io/docs/image-resize-and-crop#example---focus-using-cropped-image-coordinates).
//
// This is an alias to an internal type.
type TransformationXUnionParam = shared.TransformationXUnionParam

// Focus using cropped image coordinates - X center coordinate. See
// [Focus using cropped coordinates](https://imagekit.io/docs/image-resize-and-crop#example---focus-using-cropped-image-coordinates).
//
// This is an alias to an internal type.
type TransformationXCenterUnionParam = shared.TransformationXCenterUnionParam

// Focus using cropped image coordinates - Y coordinate. See
// [Focus using cropped coordinates](https://imagekit.io/docs/image-resize-and-crop#example---focus-using-cropped-image-coordinates).
//
// This is an alias to an internal type.
type TransformationYUnionParam = shared.TransformationYUnionParam

// Focus using cropped image coordinates - Y center coordinate. See
// [Focus using cropped coordinates](https://imagekit.io/docs/image-resize-and-crop#example---focus-using-cropped-image-coordinates).
//
// This is an alias to an internal type.
type TransformationYCenterUnionParam = shared.TransformationYCenterUnionParam

// By default, the transformation string is added as a query parameter in the URL,
// e.g., `?tr=w-100,h-100`. If you want to add the transformation string in the
// path of the URL, set this to `path`. Learn more in the
// [Transformations guide](https://imagekit.io/docs/transformations).
//
// This is an alias to an internal type.
type TransformationPosition = shared.TransformationPosition

// Equals "path"
const TransformationPositionPath = shared.TransformationPositionPath

// Equals "query"
const TransformationPositionQuery = shared.TransformationPositionQuery

// This is an alias to an internal type.
type VideoOverlayParam = shared.VideoOverlayParam
