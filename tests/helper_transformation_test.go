package imagekit_test

import (
	"testing"

	"github.com/stainless-sdks/imagekit-go"
	"github.com/stainless-sdks/imagekit-go/option"

	"github.com/stainless-sdks/imagekit-go/packages/param"
	"github.com/stainless-sdks/imagekit-go/shared"
)

func TestBuildTransformationStringDetailed(t *testing.T) {
	client := imagekit.NewClient(option.WithPrivateKey("test-key"))

	t.Run("should return empty string for empty transformation array", func(t *testing.T) {
		result := client.Helper.BuildTransformationString(nil)
		if result != "" {
			t.Errorf("Expected empty string, got %s", result)
		}

		result = client.Helper.BuildTransformationString([]shared.TransformationParam{})
		if result != "" {
			t.Errorf("Expected empty string, got %s", result)
		}
	})

	t.Run("should generate transformation string for width only", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Width: shared.TransformationWidthUnionParam{
					OfFloat: param.Opt[float64]{Value: 300},
				},
			},
		}

		result := client.Helper.BuildTransformationString(transformation)
		expected := "w-300"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("should generate transformation string for multiple parameters", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Width: shared.TransformationWidthUnionParam{
					OfFloat: param.Opt[float64]{Value: 300},
				},
				Height: shared.TransformationHeightUnionParam{
					OfFloat: param.Opt[float64]{Value: 200},
				},
			},
		}

		result := client.Helper.BuildTransformationString(transformation)
		expected := "w-300,h-200"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("should generate transformation string for chained transformations", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Width: shared.TransformationWidthUnionParam{
					OfFloat: param.Opt[float64]{Value: 300},
				},
			},
			{
				Height: shared.TransformationHeightUnionParam{
					OfFloat: param.Opt[float64]{Value: 200},
				},
			},
		}

		result := client.Helper.BuildTransformationString(transformation)
		expected := "w-300:h-200"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	// Additional test cases for comprehensive coverage
	t.Run("should handle empty transformation object", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{}, // Empty transformation
		}

		result := client.Helper.BuildTransformationString(transformation)
		expected := "" // Should return empty string for empty transformations
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("should handle transformation with overlay", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Overlay: shared.OverlayUnionParam{
					OfText: &shared.TextOverlayParam{
						Text: "Hello",
						Type: "text",
					},
				},
			},
		}

		result := client.Helper.BuildTransformationString(transformation)
		expected := "l-text,i-Hello,l-end"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("should handle raw transformation parameter", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Raw: param.Opt[string]{Value: "custom-transform-123"},
			},
		}

		result := client.Helper.BuildTransformationString(transformation)
		expected := "custom-transform-123"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("should handle mixed parameters with raw", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Width: shared.TransformationWidthUnionParam{
					OfFloat: param.Opt[float64]{Value: 300},
				},
				Raw: param.Opt[string]{Value: "custom-param-123"},
			},
		}

		result := client.Helper.BuildTransformationString(transformation)
		expected := "w-300,custom-param-123"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("should handle quality parameter", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Quality: param.Opt[float64]{Value: 80},
			},
		}

		result := client.Helper.BuildTransformationString(transformation)
		expected := "q-80"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("should handle aspect ratio parameter", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				AspectRatio: shared.TransformationAspectRatioUnionParam{
					OfString: param.Opt[string]{Value: "4:3"},
				},
			},
		}

		result := client.Helper.BuildTransformationString(transformation)
		expected := "ar-4:3"
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})
}
