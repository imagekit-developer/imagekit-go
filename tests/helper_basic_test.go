package imagekit_test

import (
	"testing"

	"github.com/imagekit-developer/imagekit-go/v2"
	"github.com/imagekit-developer/imagekit-go/v2/option"

	"github.com/imagekit-developer/imagekit-go/v2/packages/param"
	"github.com/imagekit-developer/imagekit-go/v2/shared"
)

func TestBuildSrcBasic(t *testing.T) {
	privateKey := "My Private API Key"
	client := imagekit.NewClient(option.WithPrivateKey(privateKey))

	t.Run("should return an empty string when src is not provided", func(t *testing.T) {
		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
		})

		if url != "" {
			t.Errorf("Expected empty string, got %s", url)
		}
	})

	t.Run("should generate a valid URL when src is /", func(t *testing.T) {
		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/",
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate a valid URL when src is provided without transformation", func(t *testing.T) {
		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/test_path.jpg",
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate a valid URL when a src is provided without transformation", func(t *testing.T) {
		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "https://ik.imagekit.io/test_url_endpoint/test_path_alt.jpg",
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path_alt.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate a valid URL when undefined transformation parameters are provided with path", func(t *testing.T) {
		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			Src:                    "/test_path_alt.jpg",
			TransformationPosition: shared.TransformationPositionQuery,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path_alt.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("By default transformationPosition should be query", func(t *testing.T) {
		// Create transformation parameters
		transformation := []shared.TransformationParam{
			{
				Height: shared.TransformationHeightUnionParam{
					OfString: param.Opt[string]{Value: "300"},
				},
				Width: shared.TransformationWidthUnionParam{
					OfString: param.Opt[string]{Value: "400"},
				},
			},
			{
				Rotation: shared.TransformationRotationUnionParam{
					OfFloat: param.Opt[float64]{Value: 90},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:    "https://ik.imagekit.io/test_url_endpoint",
			Src:            "/test_path.jpg",
			Transformation: transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path.jpg?tr=w-400,h-300:rt-90"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate the URL without sdk version", func(t *testing.T) {
		// Create transformation parameters
		transformation := []shared.TransformationParam{
			{
				Height: shared.TransformationHeightUnionParam{
					OfString: param.Opt[string]{Value: "300"},
				},
				Width: shared.TransformationWidthUnionParam{
					OfString: param.Opt[string]{Value: "400"},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			Src:                    "/test_path.jpg",
			Transformation:         transformation,
			TransformationPosition: shared.TransformationPositionPath,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/tr:w-400,h-300/test_path.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate the correct URL with a valid src and transformation", func(t *testing.T) {
		// Create transformation parameters
		transformation := []shared.TransformationParam{
			{
				Height: shared.TransformationHeightUnionParam{
					OfString: param.Opt[string]{Value: "300"},
				},
				Width: shared.TransformationWidthUnionParam{
					OfString: param.Opt[string]{Value: "400"},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/test_path.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path.jpg?tr=w-400,h-300"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should add transformation as query when src has absolute url even if transformationPosition is path", func(t *testing.T) {
		// Create transformation parameters
		transformation := []shared.TransformationParam{
			{
				Height: shared.TransformationHeightUnionParam{
					OfString: param.Opt[string]{Value: "300"},
				},
				Width: shared.TransformationWidthUnionParam{
					OfString: param.Opt[string]{Value: "400"},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionPath,
			Src:                    "https://my.custom.domain.com/test_path.jpg",
			Transformation:         transformation,
		})

		expected := "https://my.custom.domain.com/test_path.jpg?tr=w-400,h-300"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("Handle non-default url-endpoint case", func(t *testing.T) {
		// Create transformation parameters
		transformation := []shared.TransformationParam{
			{
				Height: shared.TransformationHeightUnionParam{
					OfString: param.Opt[string]{Value: "300"},
				},
				Width: shared.TransformationWidthUnionParam{
					OfString: param.Opt[string]{Value: "400"},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:    "https://ik.imagekit.io/imagekit_id/new-endpoint/",
			Src:            "/test_path.jpg",
			Transformation: transformation,
		})

		expected := "https://ik.imagekit.io/imagekit_id/new-endpoint/test_path.jpg?tr=w-400,h-300"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate the correct URL when the provided path contains multiple leading slashes", func(t *testing.T) {
		// Create transformation parameters
		transformation := []shared.TransformationParam{
			{
				Height: shared.TransformationHeightUnionParam{
					OfString: param.Opt[string]{Value: "300"},
				},
				Width: shared.TransformationWidthUnionParam{
					OfString: param.Opt[string]{Value: "400"},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "///test_path.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path.jpg?tr=w-400,h-300"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate the correct URL when the urlEndpoint is overridden", func(t *testing.T) {
		// Create transformation parameters
		transformation := []shared.TransformationParam{
			{
				Height: shared.TransformationHeightUnionParam{
					OfString: param.Opt[string]{Value: "300"},
				},
				Width: shared.TransformationWidthUnionParam{
					OfString: param.Opt[string]{Value: "400"},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint_alt",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/test_path.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint_alt/test_path.jpg?tr=w-400,h-300"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate the correct URL with transformationPosition as query parameter when src is provided", func(t *testing.T) {
		// Create transformation parameters
		transformation := []shared.TransformationParam{
			{
				Height: shared.TransformationHeightUnionParam{
					OfString: param.Opt[string]{Value: "300"},
				},
				Width: shared.TransformationWidthUnionParam{
					OfString: param.Opt[string]{Value: "400"},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			Src:                    "/test_path.jpg",
			TransformationPosition: shared.TransformationPositionQuery,
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path.jpg?tr=w-400,h-300"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	// Additional missing test cases from Node.js SDK
	t.Run("should generate the correct URL with a valid src parameter and transformation", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Height: shared.TransformationHeightUnionParam{
					OfString: param.Opt[string]{Value: "300"},
				},
				Width: shared.TransformationWidthUnionParam{
					OfString: param.Opt[string]{Value: "400"},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "https://ik.imagekit.io/test_url_endpoint/test_path_alt.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path_alt.jpg?tr=w-400,h-300"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should merge query parameters correctly in the generated URL", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Height: shared.TransformationHeightUnionParam{
					OfString: param.Opt[string]{Value: "300"},
				},
				Width: shared.TransformationWidthUnionParam{
					OfString: param.Opt[string]{Value: "400"},
				},
			},
		}

		queryParams := map[string]string{
			"t2": "v2",
			"t3": "v3",
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "https://ik.imagekit.io/test_url_endpoint/test_path_alt.jpg?t1=v1",
			QueryParameters:        queryParams,
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path_alt.jpg?t1=v1&t2=v2&t3=v3&tr=w-400,h-300"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate the correct URL with chained transformations", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Height: shared.TransformationHeightUnionParam{
					OfString: param.Opt[string]{Value: "300"},
				},
				Width: shared.TransformationWidthUnionParam{
					OfString: param.Opt[string]{Value: "400"},
				},
			},
			{
				Rotation: shared.TransformationRotationUnionParam{
					OfString: param.Opt[string]{Value: "90"},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/test_path.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path.jpg?tr=w-400,h-300:rt-90"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate the correct URL with chained transformations including raw transformation", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Height: shared.TransformationHeightUnionParam{
					OfString: param.Opt[string]{Value: "300"},
				},
				Width: shared.TransformationWidthUnionParam{
					OfString: param.Opt[string]{Value: "400"},
				},
			},
			{
				Raw: param.Opt[string]{Value: "rndm_trnsf-abcd"},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/test_path.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path.jpg?tr=w-400,h-300:rndm_trnsf-abcd"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate the correct URL when border transformation is applied", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Height: shared.TransformationHeightUnionParam{
					OfString: param.Opt[string]{Value: "300"},
				},
				Width: shared.TransformationWidthUnionParam{
					OfString: param.Opt[string]{Value: "400"},
				},
				Border: param.Opt[string]{Value: "20_FF0000"},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/test_path.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path.jpg?tr=w-400,h-300,b-20_FF0000"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate the correct URL when transformation has empty key and value", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Raw: param.Opt[string]{Value: ""},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/test_path.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/test_path.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate a valid URL when cname is used", func(t *testing.T) {
		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://custom.domain.com",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/test_path.jpg",
		})

		expected := "https://custom.domain.com/test_path.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate a valid URL when cname is used with a url-pattern", func(t *testing.T) {
		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://custom.domain.com/url-pattern",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/test_path.jpg",
		})

		expected := "https://custom.domain.com/url-pattern/test_path.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})
}
