package imagekit_test

import (
	"testing"

	"github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/option"

	"github.com/imagekit-developer/imagekit-go/packages/param"
	"github.com/imagekit-developer/imagekit-go/shared"
)

func TestAITransformations(t *testing.T) {
	privateKey := "My Private API Key"
	client := imagekit.NewClient(option.WithPrivateKey(privateKey))

	// AI Transformation Tests
	t.Run("should generate the correct URL for AI background removal when set to true", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				AIRemoveBackground: true,
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/test_path1.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path1.jpg?tr=e-bgremove"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate the correct URL for external AI background removal when set to true", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				AIRemoveBackgroundExternal: true,
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/test_path1.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path1.jpg?tr=e-removedotbg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate the correct URL when AI drop shadow transformation is set to true", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				AIDropShadow: shared.TransformationAIDropShadowUnionParam{
					OfTransformationAIDropShadowBoolean: param.Opt[bool]{Value: true},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/test_path1.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path1.jpg?tr=e-dropshadow"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate the correct URL when gradient transformation is set to true", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Gradient: shared.TransformationGradientUnionParam{
					OfTransformationGradientBoolean: param.Opt[bool]{Value: true},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/test_path1.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path1.jpg?tr=e-gradient"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	// Additional AI transformation tests based on Node.js SDK
	t.Run("should not apply AI background removal when value is not true", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				// AIRemoveBackground: false, // This would be the case for false values
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/test_path1.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path1.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should not apply external AI background removal when value is not true", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				// AIRemoveBackgroundExternal: false, // This would be the case for false values
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/test_path1.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path1.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should handle AI transformations with parameters", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				AIDropShadow: shared.TransformationAIDropShadowUnionParam{
					OfString: param.Opt[string]{Value: "custom-shadow-params"},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/test_path1.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path1.jpg?tr=e-dropshadow-custom-shadow-params"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should handle gradient with parameters", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Gradient: shared.TransformationGradientUnionParam{
					OfString: param.Opt[string]{Value: "ld-top_from-green_to-00FF0010_sp-1"},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/test_path1.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path1.jpg?tr=e-gradient-ld-top_from-green_to-00FF0010_sp-1"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should combine AI transformations with regular transformations", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Width: shared.TransformationWidthUnionParam{
					OfFloat: param.Opt[float64]{Value: 300},
				},
				Height: shared.TransformationHeightUnionParam{
					OfFloat: param.Opt[float64]{Value: 200},
				},
				AIRemoveBackground: true,
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/test_path1.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path1.jpg?tr=w-300,h-200,e-bgremove"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should handle multiple AI transformations", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				AIRemoveBackground: true,
				AIDropShadow: shared.TransformationAIDropShadowUnionParam{
					OfTransformationAIDropShadowBoolean: param.Opt[bool]{Value: true},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/test_path1.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path1.jpg?tr=e-bgremove,e-dropshadow"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})
}

func TestParameterSpecificTransformations(t *testing.T) {
	privateKey := "My Private API Key"
	client := imagekit.NewClient(option.WithPrivateKey(privateKey))

	// Parameter-specific tests
	t.Run("should generate the correct URL for width transformation when provided with a number value", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Width: shared.TransformationWidthUnionParam{
					OfFloat: param.Opt[float64]{Value: 400},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/test_path1.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path1.jpg?tr=w-400"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate the correct URL for height transformation when provided with a string value", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Height: shared.TransformationHeightUnionParam{
					OfString: param.Opt[string]{Value: "300"},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/test_path1.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path1.jpg?tr=h-300"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate the correct URL for aspectRatio transformation when provided with colon format", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				AspectRatio: shared.TransformationAspectRatioUnionParam{
					OfString: param.Opt[string]{Value: "4:3"},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/test_path1.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path1.jpg?tr=ar-4:3"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate the correct URL for quality transformation when provided with a number value", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Quality: param.Opt[float64]{Value: 80},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/test_path1.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path1.jpg?tr=q-80"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	// Additional parameter validation tests from Node.js SDK
	t.Run("should skip transformation parameters that are undefined or empty", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Width: shared.TransformationWidthUnionParam{
					OfFloat: param.Opt[float64]{Value: 300},
				},
				// Quality is not set - should be ignored
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/test_path1.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path1.jpg?tr=w-300"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should handle boolean transformation values", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Trim: shared.TransformationTrimUnionParam{
					OfTransformationTrimBoolean: param.Opt[bool]{Value: true},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/test_path1.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path1.jpg?tr=t-true"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should handle transformation parameter with empty string value", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				DefaultImage: param.Opt[string]{Value: ""},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/test_path1.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path1.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should handle complex transformation combinations", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Width:   shared.TransformationWidthUnionParam{OfFloat: param.Opt[float64]{Value: 300}},
				Height:  shared.TransformationHeightUnionParam{OfFloat: param.Opt[float64]{Value: 200}},
				Quality: param.Opt[float64]{Value: 85},
				Border:  param.Opt[string]{Value: "5_FF0000"},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/test_path1.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path1.jpg?tr=w-300,h-200,q-85,b-5_FF0000"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate the correct URL with many transformations, including video and AI transforms", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Height: shared.TransformationHeightUnionParam{
					OfFloat: param.Opt[float64]{Value: 300},
				},
				Width: shared.TransformationWidthUnionParam{
					OfFloat: param.Opt[float64]{Value: 400},
				},
				AspectRatio: shared.TransformationAspectRatioUnionParam{
					OfString: param.Opt[string]{Value: "4-3"},
				},
				Quality:  param.Opt[float64]{Value: 40},
				Crop:     shared.TransformationCropForce,
				CropMode: shared.TransformationCropModeExtract,
				Focus:    param.Opt[string]{Value: "left"},
				Format:   shared.TransformationFormatJpeg,
				Radius: shared.TransformationRadiusUnionParam{
					OfFloat: param.Opt[float64]{Value: 50},
				},
				Background: param.Opt[string]{Value: "A94D34"},
				Border:     param.Opt[string]{Value: "5-A94D34"},
				Rotation: shared.TransformationRotationUnionParam{
					OfFloat: param.Opt[float64]{Value: 90},
				},
				Blur:        param.Opt[float64]{Value: 10},
				Named:       param.Opt[string]{Value: "some_name"},
				Progressive: param.Opt[bool]{Value: true},
				Lossless:    param.Opt[bool]{Value: true},
				Trim: shared.TransformationTrimUnionParam{
					OfFloat: param.Opt[float64]{Value: 5},
				},
				Metadata:     param.Opt[bool]{Value: true},
				ColorProfile: param.Opt[bool]{Value: true},
				DefaultImage: param.Opt[string]{Value: "/folder/file.jpg/"},
				Dpr:          param.Opt[float64]{Value: 3},
				X: shared.TransformationXUnionParam{
					OfFloat: param.Opt[float64]{Value: 10},
				},
				Y: shared.TransformationYUnionParam{
					OfFloat: param.Opt[float64]{Value: 20},
				},
				XCenter: shared.TransformationXCenterUnionParam{
					OfFloat: param.Opt[float64]{Value: 30},
				},
				YCenter: shared.TransformationYCenterUnionParam{
					OfFloat: param.Opt[float64]{Value: 40},
				},
				Flip:    shared.TransformationFlipH,
				Opacity: param.Opt[float64]{Value: 0.8},
				Zoom:    param.Opt[float64]{Value: 2},
				// Video transformations
				VideoCodec: shared.TransformationVideoCodecH264,
				AudioCodec: shared.TransformationAudioCodecAac,
				StartOffset: shared.TransformationStartOffsetUnionParam{
					OfFloat: param.Opt[float64]{Value: 5},
				},
				EndOffset: shared.TransformationEndOffsetUnionParam{
					OfFloat: param.Opt[float64]{Value: 15},
				},
				Duration: shared.TransformationDurationUnionParam{
					OfFloat: param.Opt[float64]{Value: 10},
				},
				StreamingResolutions: []shared.StreamingResolution{shared.StreamingResolution1440, shared.StreamingResolution1080},
				// AI transformations
				Grayscale:   true,
				AIUpscale:   true,
				AIRetouch:   true,
				AIVariation: true,
				AIDropShadow: shared.TransformationAIDropShadowUnionParam{
					OfTransformationAIDropShadowBoolean: param.Opt[bool]{Value: true},
				},
				AIChangeBackground: param.Opt[string]{Value: "prompt-car"},
				AIEdit:             param.Opt[string]{Value: "prompt-make it vintage"},
				AIRemoveBackground: true,
				ContrastStretch:    true,
				Shadow: shared.TransformationShadowUnionParam{
					OfString: param.Opt[string]{Value: "bl-15_st-40_x-10_y-N5"},
				},
				Sharpen: shared.TransformationSharpenUnionParam{
					OfFloat: param.Opt[float64]{Value: 10},
				},
				UnsharpMask: shared.TransformationUnsharpMaskUnionParam{
					OfString: param.Opt[string]{Value: "2-2-0.8-0.024"},
				},
				Gradient: shared.TransformationGradientUnionParam{
					OfString: param.Opt[string]{Value: "from-red_to-white"},
				},
				Original: param.Opt[bool]{Value: true},
				Page: shared.TransformationPageUnionParam{
					OfString: param.Opt[string]{Value: "2_4"},
				},
				Raw: param.Opt[string]{Value: "h-200,w-300,l-image,i-logo.png,l-end"},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/test_path.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path.jpg?tr=w-400,h-300,q-40,ar-4-3,c-force,cm-extract,fo-left,f-jpeg,r-50,bg-A94D34,b-5-A94D34,di-folder@@file.jpg,dpr-3,x-10,y-20,xc-30,yc-40,o-0.8,z-2,rt-90,bl-10,n-some_name,pr-true,lo-true,fl-h,t-5,md-true,cp-true,vc-h264,ac-aac,so-5,eo-15,du-10,sr-1440_1080,e-grayscale,e-upscale,e-retouch,e-genvar,e-bgremove,e-contrast,e-dropshadow,e-changebg-prompt-car,e-edit-prompt-make it vintage,e-shadow-bl-15_st-40_x-10_y-N5,e-sharpen-10,e-usm-2-2-0.8-0.024,e-gradient-from-red_to-white,orig-true,pg-2_4,h-200,w-300,l-image,i-logo.png,l-end"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})
}
