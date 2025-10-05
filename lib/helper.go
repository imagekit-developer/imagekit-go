// Package lib provides helper utilities for ImageKit SDK
// This file contains custom helper functions - not generated
package lib

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/imagekit-developer/imagekit-go/v2/internal/requestconfig"
	"github.com/imagekit-developer/imagekit-go/v2/option"
	"github.com/imagekit-developer/imagekit-go/v2/packages/param"
	"github.com/imagekit-developer/imagekit-go/v2/shared"
)

// HelperService contains utility methods for ImageKit SDK operations like URL building,
// transformation string generation, and authentication parameter generation.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHelperService] method instead.
type HelperService struct {
	Options    []option.RequestOption
	privateKey string
}

// NewHelperService generates a new helper service that applies the given options.
// It extracts the private key from the options for use in signing and authentication.
func NewHelperService(opts ...option.RequestOption) (r HelperService) {
	r = HelperService{Options: opts}

	// Extract private key from options using a safer approach
	// We'll create a new config and apply only the options that we can safely handle
	cfg := &requestconfig.RequestConfig{}

	// Apply options one by one, skipping any that cause issues
	for _, opt := range opts {
		if opt != nil {
			// Try to apply the option, but catch any panics
			func() {
				defer func() {
					if recovered := recover(); recovered != nil {
						// Skip this option if it causes a panic
					}
				}()
				opt.Apply(cfg)
			}()
		}
	}

	r.privateKey = cfg.PrivateKey
	return
}

// Builds a source URL with the given options.
func (r *HelperService) BuildURL(opts shared.SrcOptionsParam) string {
	// Set defaults
	if opts.URLEndpoint == "" {
		opts.URLEndpoint = ""
	}
	if opts.Src == "" {
		opts.Src = ""
	}
	transformationPosition := opts.TransformationPosition
	if transformationPosition == "" {
		transformationPosition = shared.TransformationPositionQuery
	}

	if opts.Src == "" {
		return ""
	}

	isAbsoluteURL := strings.HasPrefix(opts.Src, "http://") || strings.HasPrefix(opts.Src, "https://")

	var urlObj *url.URL
	var err error
	var isSrcParameterUsedForURL bool

	// Try to parse URL
	if !isAbsoluteURL {
		urlObj, err = url.Parse(opts.URLEndpoint)
		if err != nil {
			return ""
		}
	} else {
		urlObj, err = url.Parse(opts.Src)
		if err != nil {
			return ""
		}
		isSrcParameterUsedForURL = true
	}

	// Add query parameters
	if opts.QueryParameters != nil {
		query := urlObj.Query()
		for key, value := range opts.QueryParameters {
			query.Set(key, value)
		}
		urlObj.RawQuery = query.Encode()
	}

	// Build transformation string
	transformationString := r.buildTransformationStringInternal(opts.Transformation)

	// Determine if transformation will be in query params
	addAsQuery := transformationPosition == shared.TransformationPositionQuery || isSrcParameterUsedForURL

	// Transformation placeholder to avoid URL encoding issues
	const transformationPlaceholder = "PLEASEREPLACEJUSTBEFORESIGN"

	if !isAbsoluteURL {
		// For non-absolute URLs, construct the path: endpoint_path + transformations + src
		endpointURL, _ := url.Parse(opts.URLEndpoint)
		endpointPath := endpointURL.Path
		pathParts := []string{endpointPath}

		if transformationString != "" && !addAsQuery {
			pathParts = append(pathParts, transformationParameter+chainTransformDelimiter+transformationPlaceholder)
		}

		pathParts = append(pathParts, opts.Src)
		urlObj.Path = pathJoin(pathParts)
	}

	// First, build the complete URL with transformations
	finalURL := urlObj.String()

	// Add transformation parameter manually to avoid URL encoding
	// URLSearchParams.set() would encode commas and colons in transformation string,
	// It would work correctly but not very readable e.g., "w-300,h-400" is better than "w-300%2Ch-400"
	if transformationString != "" && addAsQuery {
		separator := "&"
		if urlObj.RawQuery == "" {
			separator = "?"
		}
		finalURL = fmt.Sprintf("%s%s%s=%s", finalURL, separator, transformationParameter, transformationPlaceholder)
	}

	// Replace the placeholder with actual transformation string
	// We don't put actual transformation string before signing to avoid issues with URL encoding
	if transformationString != "" {
		finalURL = strings.Replace(finalURL, transformationPlaceholder, transformationString, -1)
	}

	// Then sign the URL if needed
	shouldSign := (!param.IsOmitted(opts.Signed) && opts.Signed.Value) || (!param.IsOmitted(opts.ExpiresIn) && opts.ExpiresIn.Value > 0)
	if shouldSign {
		expiryTimestamp := getSignatureTimestamp(opts.ExpiresIn)

		urlSignature := r.getSignature(finalURL, opts.URLEndpoint, expiryTimestamp)

		// Add signature parameters to the final URL
		// Use URL object to properly determine if we need ? or & separator
		finalURLObj, _ := url.Parse(finalURL)
		hasExistingParams := finalURLObj.RawQuery != ""
		separator := "&"
		if !hasExistingParams {
			separator = "?"
		}

		if expiryTimestamp != defaultTimestamp {
			finalURL += fmt.Sprintf("%s%s=%d", separator, timestampParameter, expiryTimestamp)
			finalURL += fmt.Sprintf("&%s=%s", signatureParameter, urlSignature)
		} else {
			finalURL += fmt.Sprintf("%s%s=%s", separator, signatureParameter, urlSignature)
		}
	}

	return finalURL
}

// Builds a transformation string from the given transformations.
func (r *HelperService) BuildTransformationString(transformation []shared.TransformationParam) string {
	return r.buildTransformationStringInternal(transformation)
}

// Generates authentication parameters for client-side file uploads using ImageKit's Upload API V1.
// This method creates the required authentication signature that allows secure file uploads directly from the browser or mobile applications without exposing your private API key.
//
// Parameters:
//   - token: A unique token for this upload request. If empty, a UUID will be generated.
//   - expire: Unix timestamp when this authentication should expire. If 0, defaults to 30 minutes from now.
//
// Returns a map containing:
//   - "token": The authentication token
//   - "expire": The expiration timestamp
//   - "signature": The HMAC signature
func (r *HelperService) GetAuthenticationParameters(token string, expire int64) (map[string]interface{}, error) {
	if r.privateKey == "" {
		return nil, fmt.Errorf("private API key is required for authentication parameters generation")
	}

	defaultTimeDiff := int64(60 * 30) // 30 minutes
	defaultExpire := time.Now().Unix() + defaultTimeDiff

	finalToken := token
	if finalToken == "" {
		// Generate 16 random bytes (128 bits of entropy - more than UUID v4's 122 bits)
		tokenBytes := make([]byte, 16)
		if _, err := rand.Read(tokenBytes); err != nil {
			return nil, fmt.Errorf("failed to generate random token: %w", err)
		}
		// Use hex encoding for readability and compatibility
		finalToken = hex.EncodeToString(tokenBytes)
	}

	finalExpire := expire
	if finalExpire == 0 {
		finalExpire = defaultExpire
	}

	return getAuthenticationParameters(finalToken, finalExpire, r.privateKey), nil
}

// buildTransformationStringInternal builds a transformation string with context
func (r *HelperService) buildTransformationStringInternal(transformation []shared.TransformationParam) string {
	if len(transformation) == 0 {
		return ""
	}

	var parsedTransforms []string

	for _, currentTransform := range transformation {
		var parsedTransformStep []string

		// Define transformation mappings in order
		transformMappings := []struct {
			key    string
			getter func() string
		}{
			{"w", func() string {
				return r.getUnionParamValue(currentTransform.Width.OfFloat, currentTransform.Width.OfString)
			}},
			{"h", func() string {
				return r.getUnionParamValue(currentTransform.Height.OfFloat, currentTransform.Height.OfString)
			}},
			{"q", func() string { return r.getOptParamValue(currentTransform.Quality) }},
			{"ar", func() string {
				return r.getUnionParamValue(currentTransform.AspectRatio.OfFloat, currentTransform.AspectRatio.OfString)
			}},
			{"c", func() string { return r.getEnumValue(string(currentTransform.Crop)) }},
			{"cm", func() string { return r.getEnumValue(string(currentTransform.CropMode)) }},
			{"fo", func() string { return r.getOptParamValue(currentTransform.Focus) }},
			{"f", func() string { return r.getEnumValue(string(currentTransform.Format)) }},
			{"r", func() string {
				if !param.IsOmitted(currentTransform.Radius.OfMax) {
					return "max"
				}
				return r.getOptParamValue(currentTransform.Radius.OfFloat)
			}},
			{"bg", func() string { return r.getOptParamValue(currentTransform.Background) }},
			{"b", func() string { return r.getOptParamValue(currentTransform.Border) }},
			{"di", func() string {
				if !param.IsOmitted(currentTransform.DefaultImage) && currentTransform.DefaultImage.Value != "" {
					value := currentTransform.DefaultImage.Value
					value = removeTrailingSlash(removeLeadingSlash(value))
					return strings.ReplaceAll(value, "/", "@@")
				}
				return ""
			}},
			{"dpr", func() string { return r.getOptParamValue(currentTransform.Dpr) }},
			{"x", func() string { return r.getUnionParamValue(currentTransform.X.OfFloat, currentTransform.X.OfString) }},
			{"y", func() string { return r.getUnionParamValue(currentTransform.Y.OfFloat, currentTransform.Y.OfString) }},
			{"xc", func() string {
				return r.getUnionParamValue(currentTransform.XCenter.OfFloat, currentTransform.XCenter.OfString)
			}},
			{"yc", func() string {
				return r.getUnionParamValue(currentTransform.YCenter.OfFloat, currentTransform.YCenter.OfString)
			}},
			{"o", func() string { return r.getOptParamValue(currentTransform.Opacity) }},
			{"z", func() string { return r.getOptParamValue(currentTransform.Zoom) }},
			{"rt", func() string {
				return r.getUnionParamValue(currentTransform.Rotation.OfFloat, currentTransform.Rotation.OfString)
			}},
			{"bl", func() string { return r.getOptParamValue(currentTransform.Blur) }},
			{"n", func() string { return r.getOptParamValue(currentTransform.Named) }},
			{"pr", func() string { return r.getOptParamValue(currentTransform.Progressive) }},
			{"lo", func() string { return r.getOptParamValue(currentTransform.Lossless) }},
			{"fl", func() string { return r.getEnumValue(string(currentTransform.Flip)) }},
			{"t", func() string {
				if !param.IsOmitted(currentTransform.Trim.OfTransformationTrimBoolean) && currentTransform.Trim.OfTransformationTrimBoolean.Value {
					return "true"
				}
				return r.getOptParamValue(currentTransform.Trim.OfFloat)
			}},
			{"md", func() string { return r.getOptParamValue(currentTransform.Metadata) }},
			{"cp", func() string { return r.getOptParamValue(currentTransform.ColorProfile) }},
			{"vc", func() string { return r.getEnumValue(string(currentTransform.VideoCodec)) }},
			{"ac", func() string { return r.getEnumValue(string(currentTransform.AudioCodec)) }},
			{"so", func() string {
				return r.getUnionParamValue(currentTransform.StartOffset.OfFloat, currentTransform.StartOffset.OfString)
			}},
			{"eo", func() string {
				return r.getUnionParamValue(currentTransform.EndOffset.OfFloat, currentTransform.EndOffset.OfString)
			}},
			{"du", func() string {
				return r.getUnionParamValue(currentTransform.Duration.OfFloat, currentTransform.Duration.OfString)
			}},
			{"sr", func() string {
				if len(currentTransform.StreamingResolutions) > 0 {
					var resolutions []string
					for _, res := range currentTransform.StreamingResolutions {
						resolutions = append(resolutions, string(res))
					}
					return strings.Join(resolutions, "_")
				}
				return ""
			}},
		}

		// Process basic transformations
		for _, mapping := range transformMappings {
			if value := mapping.getter(); value != "" {
				parsedTransformStep = append(parsedTransformStep, fmt.Sprintf("%s%s%s", mapping.key, transformKeyValueDelimiter, value))
			}
		}

		// Process AI transformations (boolean flags)
		aiFlags := []struct {
			flag  bool
			value string
		}{
			{currentTransform.Grayscale, "e-grayscale"},
			{currentTransform.AIUpscale, "e-upscale"},
			{currentTransform.AIRetouch, "e-retouch"},
			{currentTransform.AIVariation, "e-genvar"},
			{currentTransform.AIRemoveBackground, "e-bgremove"},
			{currentTransform.AIRemoveBackgroundExternal, "e-removedotbg"},
			{currentTransform.ContrastStretch, "e-contrast"},
		}

		for _, flag := range aiFlags {
			if flag.flag {
				parsedTransformStep = append(parsedTransformStep, flag.value)
			}
		}

		// Process AI transformations with parameters
		if !param.IsOmitted(currentTransform.AIDropShadow.OfTransformationAIDropShadowBoolean) && currentTransform.AIDropShadow.OfTransformationAIDropShadowBoolean.Value {
			parsedTransformStep = append(parsedTransformStep, "e-dropshadow")
		} else if value := r.getOptParamValue(currentTransform.AIDropShadow.OfString); value != "" {
			parsedTransformStep = append(parsedTransformStep, fmt.Sprintf("e-dropshadow%s%s", transformKeyValueDelimiter, value))
		}

		if value := r.getOptParamValue(currentTransform.AIChangeBackground); value != "" {
			parsedTransformStep = append(parsedTransformStep, fmt.Sprintf("e-changebg%s%s", transformKeyValueDelimiter, value))
		}

		if value := r.getOptParamValue(currentTransform.AIEdit); value != "" {
			parsedTransformStep = append(parsedTransformStep, fmt.Sprintf("e-edit%s%s", transformKeyValueDelimiter, value))
		}

		// Process effects
		effectMappings := []struct {
			prefix string
			getter func() (string, bool)
		}{
			{"e-shadow", func() (string, bool) {
				if !param.IsOmitted(currentTransform.Shadow.OfTransformationShadowBoolean) && currentTransform.Shadow.OfTransformationShadowBoolean.Value {
					return "", true
				}
				if value := r.getOptParamValue(currentTransform.Shadow.OfString); value != "" {
					return value, false
				}
				return "", false
			}},
			{"e-sharpen", func() (string, bool) {
				if !param.IsOmitted(currentTransform.Sharpen.OfTransformationSharpenBoolean) && currentTransform.Sharpen.OfTransformationSharpenBoolean.Value {
					return "", true
				}
				if value := r.getOptParamValue(currentTransform.Sharpen.OfFloat); value != "" {
					return value, false
				}
				return "", false
			}},
			{"e-usm", func() (string, bool) {
				if !param.IsOmitted(currentTransform.UnsharpMask.OfTransformationUnsharpMaskBoolean) && currentTransform.UnsharpMask.OfTransformationUnsharpMaskBoolean.Value {
					return "", true
				}
				if value := r.getOptParamValue(currentTransform.UnsharpMask.OfString); value != "" {
					return value, false
				}
				return "", false
			}},
			{"e-gradient", func() (string, bool) {
				if !param.IsOmitted(currentTransform.Gradient.OfTransformationGradientBoolean) && currentTransform.Gradient.OfTransformationGradientBoolean.Value {
					return "", true
				}
				if value := r.getOptParamValue(currentTransform.Gradient.OfString); value != "" {
					return value, false
				}
				return "", false
			}},
		}

		for _, effect := range effectMappings {
			if value, isFlag := effect.getter(); value != "" || isFlag {
				if isFlag {
					parsedTransformStep = append(parsedTransformStep, effect.prefix)
				} else {
					parsedTransformStep = append(parsedTransformStep, fmt.Sprintf("%s%s%s", effect.prefix, transformKeyValueDelimiter, value))
				}
			}
		}

		// Handle Original
		if !param.IsOmitted(currentTransform.Original) && currentTransform.Original.Value {
			parsedTransformStep = append(parsedTransformStep, "orig-true")
		}

		// Handle Page
		if value := r.getUnionParamValue(currentTransform.Page.OfFloat, currentTransform.Page.OfString); value != "" {
			parsedTransformStep = append(parsedTransformStep, fmt.Sprintf("pg%s%s", transformKeyValueDelimiter, value))
		}

		// Handle Overlay
		if overlayString := r.processOverlay(currentTransform.Overlay); overlayString != "" {
			parsedTransformStep = append(parsedTransformStep, overlayString)
		}

		// Handle Raw parameter last
		if !param.IsOmitted(currentTransform.Raw) && currentTransform.Raw.Value != "" {
			parsedTransformStep = append(parsedTransformStep, currentTransform.Raw.Value)
		}

		if len(parsedTransformStep) > 0 {
			parsedTransforms = append(parsedTransforms, strings.Join(parsedTransformStep, transformDelimiter))
		}
	}

	return strings.Join(parsedTransforms, chainTransformDelimiter)
}

// Helper methods to reduce repetition
func (r *HelperService) getOptParamValue(opt interface{}) string {
	switch v := opt.(type) {
	case param.Opt[string]:
		if !param.IsOmitted(v) {
			return v.Value
		}
	case param.Opt[float64]:
		if !param.IsOmitted(v) {
			return fmt.Sprintf("%g", v.Value)
		}
	case param.Opt[bool]:
		if !param.IsOmitted(v) {
			return fmt.Sprintf("%t", v.Value)
		}
	}
	return ""
}

func (r *HelperService) getUnionParamValue(floatOpt param.Opt[float64], stringOpt param.Opt[string]) string {
	if !param.IsOmitted(floatOpt) {
		return fmt.Sprintf("%g", floatOpt.Value)
	}
	if !param.IsOmitted(stringOpt) {
		return stringOpt.Value
	}
	return ""
}

func (r *HelperService) getEnumValue(value string) string {
	if value != "" {
		return value
	}
	return ""
}

func (r *HelperService) processOverlay(overlay shared.OverlayUnionParam) string {
	var entries []string
	var baseOverlay *shared.BaseOverlayParam
	var transformationString string

	// Handle text overlay
	if overlay.OfText != nil {
		textOverlay := overlay.OfText
		if textOverlay.Text == "" {
			return ""
		}

		encoding := textOverlay.Encoding
		if encoding == "" {
			encoding = "auto"
		}

		entries = append(entries, "l-text")
		entries = append(entries, processText(textOverlay.Text, encoding))

		baseOverlay = &textOverlay.BaseOverlayParam

		// Process text overlay transformations
		if len(textOverlay.Transformation) > 0 {
			transformationString = r.buildTextOverlayTransformation(textOverlay.Transformation)
		}
	}

	// Handle image overlay
	if overlay.OfImage != nil {
		imageOverlay := overlay.OfImage
		if imageOverlay.Input == "" {
			return ""
		}

		encoding := imageOverlay.Encoding
		if encoding == "" {
			encoding = "auto"
		}

		entries = append(entries, "l-image")
		entries = append(entries, processInputPath(imageOverlay.Input, encoding))
		baseOverlay = &imageOverlay.BaseOverlayParam

		// Process image overlay transformations (regular transformations)
		if len(imageOverlay.Transformation) > 0 {
			transformationString = r.buildTransformationStringInternal(imageOverlay.Transformation)
		}
	}

	// Handle video overlay
	if overlay.OfVideo != nil {
		videoOverlay := overlay.OfVideo
		if videoOverlay.Input == "" {
			return ""
		}

		encoding := videoOverlay.Encoding
		if encoding == "" {
			encoding = "auto"
		}

		entries = append(entries, "l-video")
		entries = append(entries, processInputPath(videoOverlay.Input, encoding))
		baseOverlay = &videoOverlay.BaseOverlayParam

		// Process video overlay transformations (regular transformations)
		if len(videoOverlay.Transformation) > 0 {
			transformationString = r.buildTransformationStringInternal(videoOverlay.Transformation)
		}
	}

	// Handle subtitle overlay
	if overlay.OfSubtitle != nil {
		subtitleOverlay := overlay.OfSubtitle
		if subtitleOverlay.Input == "" {
			return ""
		}

		encoding := subtitleOverlay.Encoding
		if encoding == "" {
			encoding = "auto"
		}

		entries = append(entries, "l-subtitle")
		entries = append(entries, processInputPath(subtitleOverlay.Input, encoding))
		baseOverlay = &subtitleOverlay.BaseOverlayParam

		// Process subtitle overlay transformations
		if len(subtitleOverlay.Transformation) > 0 {
			transformationString = r.buildSubtitleOverlayTransformation(subtitleOverlay.Transformation)
		}
	}

	// Handle solid color overlay
	if overlay.OfSolidColor != nil {
		solidColorOverlay := overlay.OfSolidColor
		if solidColorOverlay.Color == "" {
			return ""
		}

		entries = append(entries, "l-image")
		entries = append(entries, "i-ik_canvas")
		entries = append(entries, fmt.Sprintf("bg-%s", solidColorOverlay.Color))
		baseOverlay = &solidColorOverlay.BaseOverlayParam

		// Process solid color overlay transformations
		if len(solidColorOverlay.Transformation) > 0 {
			transformationString = r.buildSolidColorOverlayTransformation(solidColorOverlay.Transformation)
		}
	}

	// Process position and timing if baseOverlay is set
	if baseOverlay != nil {
		// Position
		position := baseOverlay.Position
		if !param.IsOmitted(position.X.OfFloat) {
			entries = append(entries, fmt.Sprintf("lx-%g", position.X.OfFloat.Value))
		} else if !param.IsOmitted(position.X.OfString) {
			entries = append(entries, fmt.Sprintf("lx-%s", position.X.OfString.Value))
		}

		if !param.IsOmitted(position.Y.OfFloat) {
			entries = append(entries, fmt.Sprintf("ly-%g", position.Y.OfFloat.Value))
		} else if !param.IsOmitted(position.Y.OfString) {
			entries = append(entries, fmt.Sprintf("ly-%s", position.Y.OfString.Value))
		}

		if position.Focus != "" {
			entries = append(entries, fmt.Sprintf("lfo-%s", position.Focus))
		}

		// Timing
		timing := baseOverlay.Timing
		if !param.IsOmitted(timing.Start.OfFloat) {
			entries = append(entries, fmt.Sprintf("lso-%g", timing.Start.OfFloat.Value))
		} else if !param.IsOmitted(timing.Start.OfString) {
			entries = append(entries, fmt.Sprintf("lso-%s", timing.Start.OfString.Value))
		}

		if !param.IsOmitted(timing.End.OfFloat) {
			entries = append(entries, fmt.Sprintf("leo-%g", timing.End.OfFloat.Value))
		} else if !param.IsOmitted(timing.End.OfString) {
			entries = append(entries, fmt.Sprintf("leo-%s", timing.End.OfString.Value))
		}

		if !param.IsOmitted(timing.Duration.OfFloat) {
			entries = append(entries, fmt.Sprintf("ldu-%g", timing.Duration.OfFloat.Value))
		} else if !param.IsOmitted(timing.Duration.OfString) {
			entries = append(entries, fmt.Sprintf("ldu-%s", timing.Duration.OfString.Value))
		}
	}

	// Process overlay transformation if present
	if transformationString != "" {
		entries = append(entries, transformationString)
	}

	// Add l-end if we have any entries
	if len(entries) > 0 {
		entries = append(entries, "l-end")
		return strings.Join(entries, transformDelimiter)
	}

	return ""
}

func (r *HelperService) getSignature(finalURL, urlEndpoint string, expiryTimestamp int64) string {
	if r.privateKey == "" || finalURL == "" || urlEndpoint == "" {
		return ""
	}

	// Create the string to sign: relative path + expiry timestamp
	urlEndpointWithSlash := addTrailingSlash(urlEndpoint)
	stringToSign := strings.Replace(finalURL, urlEndpointWithSlash, "", 1) + strconv.FormatInt(expiryTimestamp, 10)

	return createHmacSha1(r.privateKey, stringToSign)
}

func (r *HelperService) buildTextOverlayTransformation(transformations []shared.TextOverlayTransformationParam) string {
	if len(transformations) == 0 {
		return ""
	}

	var entries []string

	for _, trans := range transformations {
		// Width
		if !param.IsOmitted(trans.Width.OfFloat) {
			entries = append(entries, fmt.Sprintf("w%s%g", transformKeyValueDelimiter, trans.Width.OfFloat.Value))
		} else if !param.IsOmitted(trans.Width.OfString) {
			entries = append(entries, fmt.Sprintf("w%s%s", transformKeyValueDelimiter, trans.Width.OfString.Value))
		}

		// FontSize
		if !param.IsOmitted(trans.FontSize.OfFloat) {
			entries = append(entries, fmt.Sprintf("fs%s%g", transformKeyValueDelimiter, trans.FontSize.OfFloat.Value))
		} else if !param.IsOmitted(trans.FontSize.OfString) {
			entries = append(entries, fmt.Sprintf("fs%s%s", transformKeyValueDelimiter, trans.FontSize.OfString.Value))
		}

		// FontFamily
		if !param.IsOmitted(trans.FontFamily) && trans.FontFamily.Value != "" {
			fontFamily := removeTrailingSlash(removeLeadingSlash(trans.FontFamily.Value))
			fontFamily = strings.ReplaceAll(fontFamily, "/", "@@")
			entries = append(entries, fmt.Sprintf("ff%s%s", transformKeyValueDelimiter, fontFamily))
		}

		// FontColor
		if !param.IsOmitted(trans.FontColor) && trans.FontColor.Value != "" {
			entries = append(entries, fmt.Sprintf("co%s%s", transformKeyValueDelimiter, trans.FontColor.Value))
		}

		// InnerAlignment
		if trans.InnerAlignment != "" {
			entries = append(entries, fmt.Sprintf("ia%s%s", transformKeyValueDelimiter, trans.InnerAlignment))
		}

		// Padding
		if !param.IsOmitted(trans.Padding.OfFloat) {
			entries = append(entries, fmt.Sprintf("pa%s%g", transformKeyValueDelimiter, trans.Padding.OfFloat.Value))
		} else if !param.IsOmitted(trans.Padding.OfString) {
			entries = append(entries, fmt.Sprintf("pa%s%s", transformKeyValueDelimiter, trans.Padding.OfString.Value))
		}

		// Alpha
		if !param.IsOmitted(trans.Alpha) {
			entries = append(entries, fmt.Sprintf("al%s%g", transformKeyValueDelimiter, trans.Alpha.Value))
		}

		// Typography
		if !param.IsOmitted(trans.Typography) && trans.Typography.Value != "" {
			entries = append(entries, fmt.Sprintf("tg%s%s", transformKeyValueDelimiter, trans.Typography.Value))
		}

		// Background
		if !param.IsOmitted(trans.Background) && trans.Background.Value != "" {
			entries = append(entries, fmt.Sprintf("bg%s%s", transformKeyValueDelimiter, trans.Background.Value))
		}

		// Radius
		if !param.IsOmitted(trans.Radius.OfFloat) {
			entries = append(entries, fmt.Sprintf("r%s%g", transformKeyValueDelimiter, trans.Radius.OfFloat.Value))
		} else if !param.IsOmitted(trans.Radius.OfMax) {
			entries = append(entries, fmt.Sprintf("r%smax", transformKeyValueDelimiter))
		}

		// Rotation
		if !param.IsOmitted(trans.Rotation.OfFloat) {
			entries = append(entries, fmt.Sprintf("rt%s%g", transformKeyValueDelimiter, trans.Rotation.OfFloat.Value))
		} else if !param.IsOmitted(trans.Rotation.OfString) {
			entries = append(entries, fmt.Sprintf("rt%s%s", transformKeyValueDelimiter, trans.Rotation.OfString.Value))
		}

		// Flip
		if trans.Flip != "" {
			entries = append(entries, fmt.Sprintf("fl%s%s", transformKeyValueDelimiter, trans.Flip))
		}

		// LineHeight
		if !param.IsOmitted(trans.LineHeight.OfFloat) {
			entries = append(entries, fmt.Sprintf("lh%s%g", transformKeyValueDelimiter, trans.LineHeight.OfFloat.Value))
		} else if !param.IsOmitted(trans.LineHeight.OfString) {
			entries = append(entries, fmt.Sprintf("lh%s%s", transformKeyValueDelimiter, trans.LineHeight.OfString.Value))
		}
	}

	return strings.Join(entries, transformDelimiter)
}

func (r *HelperService) buildSubtitleOverlayTransformation(transformations []shared.SubtitleOverlayTransformationParam) string {
	if len(transformations) == 0 {
		return ""
	}

	var entries []string

	for _, trans := range transformations {
		// Background
		if !param.IsOmitted(trans.Background) && trans.Background.Value != "" {
			entries = append(entries, fmt.Sprintf("bg%s%s", transformKeyValueDelimiter, trans.Background.Value))
		}

		// Color (font color)
		if !param.IsOmitted(trans.Color) && trans.Color.Value != "" {
			entries = append(entries, fmt.Sprintf("co%s%s", transformKeyValueDelimiter, trans.Color.Value))
		}

		// FontSize
		if !param.IsOmitted(trans.FontSize) {
			entries = append(entries, fmt.Sprintf("fs%s%g", transformKeyValueDelimiter, trans.FontSize.Value))
		}

		// FontFamily
		if !param.IsOmitted(trans.FontFamily) && trans.FontFamily.Value != "" {
			entries = append(entries, fmt.Sprintf("ff%s%s", transformKeyValueDelimiter, trans.FontFamily.Value))
		}

		// FontOutline
		if !param.IsOmitted(trans.FontOutline) && trans.FontOutline.Value != "" {
			entries = append(entries, fmt.Sprintf("fol%s%s", transformKeyValueDelimiter, trans.FontOutline.Value))
		}

		// FontShadow
		if !param.IsOmitted(trans.FontShadow) && trans.FontShadow.Value != "" {
			entries = append(entries, fmt.Sprintf("fsh%s%s", transformKeyValueDelimiter, trans.FontShadow.Value))
		}

		// Typography
		if trans.Typography != "" {
			entries = append(entries, fmt.Sprintf("tg%s%s", transformKeyValueDelimiter, trans.Typography))
		}
	}

	return strings.Join(entries, transformDelimiter)
}

func (r *HelperService) buildSolidColorOverlayTransformation(transformations []shared.SolidColorOverlayTransformationParam) string {
	if len(transformations) == 0 {
		return ""
	}

	var entries []string

	for _, trans := range transformations {
		// Width
		if !param.IsOmitted(trans.Width.OfFloat) {
			entries = append(entries, fmt.Sprintf("w%s%g", transformKeyValueDelimiter, trans.Width.OfFloat.Value))
		} else if !param.IsOmitted(trans.Width.OfString) {
			entries = append(entries, fmt.Sprintf("w%s%s", transformKeyValueDelimiter, trans.Width.OfString.Value))
		}

		// Height
		if !param.IsOmitted(trans.Height.OfFloat) {
			entries = append(entries, fmt.Sprintf("h%s%g", transformKeyValueDelimiter, trans.Height.OfFloat.Value))
		} else if !param.IsOmitted(trans.Height.OfString) {
			entries = append(entries, fmt.Sprintf("h%s%s", transformKeyValueDelimiter, trans.Height.OfString.Value))
		}

		// Alpha
		if !param.IsOmitted(trans.Alpha) {
			entries = append(entries, fmt.Sprintf("al%s%g", transformKeyValueDelimiter, trans.Alpha.Value))
		}

		// Background
		if !param.IsOmitted(trans.Background) && trans.Background.Value != "" {
			entries = append(entries, fmt.Sprintf("bg%s%s", transformKeyValueDelimiter, trans.Background.Value))
		}

		// Gradient
		if !param.IsOmitted(trans.Gradient.OfSolidColorOverlayTransformationGradientBoolean) && trans.Gradient.OfSolidColorOverlayTransformationGradientBoolean.Value {
			entries = append(entries, "e-gradient")
		} else if !param.IsOmitted(trans.Gradient.OfString) {
			entries = append(entries, fmt.Sprintf("e-gradient%s%s", transformKeyValueDelimiter, trans.Gradient.OfString.Value))
		}

		// Radius
		if !param.IsOmitted(trans.Radius.OfFloat) {
			entries = append(entries, fmt.Sprintf("r%s%g", transformKeyValueDelimiter, trans.Radius.OfFloat.Value))
		} else if !param.IsOmitted(trans.Radius.OfMax) {
			entries = append(entries, fmt.Sprintf("r%smax", transformKeyValueDelimiter))
		}
	}

	return strings.Join(entries, transformDelimiter)
}

const (
	transformationParameter    = "tr"
	signatureParameter         = "ik-s"
	timestampParameter         = "ik-t"
	defaultTimestamp           = 9999999999
	chainTransformDelimiter    = ":"
	transformDelimiter         = ","
	transformKeyValueDelimiter = "-"
)

var (
	simpleOverlayPathRegex = regexp.MustCompile(`^[a-zA-Z0-9-._/ ]*$`)
	simpleOverlayTextRegex = regexp.MustCompile(`^[a-zA-Z0-9-._ ]*$`)
)

// Helper functions
func getAuthenticationParameters(token string, expire int64, privateKey string) map[string]interface{} {
	signature := createHmacSha1(privateKey, token+strconv.FormatInt(expire, 10))

	return map[string]interface{}{
		"token":     token,
		"expire":    expire,
		"signature": signature,
	}
}

func createHmacSha1(key, data string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func removeTrailingSlash(str string) string {
	if len(str) > 0 && str[len(str)-1] == '/' {
		return str[:len(str)-1]
	}
	return str
}

func removeLeadingSlash(str string) string {
	if len(str) > 0 && str[0] == '/' {
		return str[1:]
	}
	return str
}

func pathJoin(parts []string) string {
	if len(parts) == 0 {
		return ""
	}

	joined := strings.Join(parts, "/")
	// Replace multiple slashes with single slash
	re := regexp.MustCompile("/+")
	return re.ReplaceAllString(joined, "/")
}

func processInputPath(str, encoding string) string {
	// Remove leading and trailing slashes
	str = removeTrailingSlash(removeLeadingSlash(str))

	if encoding == "plain" {
		return fmt.Sprintf("i-%s", strings.ReplaceAll(str, "/", "@@"))
	}

	if encoding == "base64" {
		base64Str := base64.StdEncoding.EncodeToString([]byte(str))
		return fmt.Sprintf("ie-%s", url.QueryEscape(base64Str))
	}

	if simpleOverlayPathRegex.MatchString(str) {
		return fmt.Sprintf("i-%s", strings.ReplaceAll(str, "/", "@@"))
	}

	base64Str := base64.StdEncoding.EncodeToString([]byte(str))
	return fmt.Sprintf("ie-%s", url.QueryEscape(base64Str))
}

func processText(str, encoding string) string {
	if encoding == "plain" {
		return fmt.Sprintf("i-%s", url.PathEscape(str))
	}

	if encoding == "base64" {
		base64Str := base64.StdEncoding.EncodeToString([]byte(str))
		return fmt.Sprintf("ie-%s", url.QueryEscape(base64Str))
	}

	if simpleOverlayTextRegex.MatchString(str) {
		return fmt.Sprintf("i-%s", url.PathEscape(str))
	}

	// For auto encoding with special characters, use base64
	base64Str := base64.StdEncoding.EncodeToString([]byte(str))
	return fmt.Sprintf("ie-%s", url.QueryEscape(base64Str))
}

func getSignatureTimestamp(expiresIn param.Opt[float64]) int64 {
	if param.IsOmitted(expiresIn) || expiresIn.Value <= 0 {
		return defaultTimestamp
	}

	seconds := int64(expiresIn.Value)
	if seconds <= 0 {
		return defaultTimestamp
	}

	currentTimestamp := time.Now().Unix()
	return currentTimestamp + seconds
}

func addTrailingSlash(str string) string {
	if len(str) > 0 && str[len(str)-1] != '/' {
		return str + "/"
	}
	return str
}
