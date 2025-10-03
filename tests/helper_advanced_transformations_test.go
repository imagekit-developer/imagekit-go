package imagekit_test

import (
	"github.com/stainless-sdks/imagekit-go/option"
	"github.com/stainless-sdks/imagekit-go"
	"testing"

	"github.com/stainless-sdks/imagekit-go/packages/param"
	"github.com/stainless-sdks/imagekit-go/shared"
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
					OfString: param.Opt[string]{Value: "linear-top-FF0000-0000FF"},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/test_path1.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path1.jpg?tr=e-gradient-linear-top-FF0000-0000FF"
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
}
