package imagekit_test

import (
	"strings"
	"testing"

	"github.com/imagekit-developer/imagekit-go/v2"
	"github.com/imagekit-developer/imagekit-go/v2/option"

	"github.com/imagekit-developer/imagekit-go/v2/packages/param"
	"github.com/imagekit-developer/imagekit-go/v2/shared"
)

func TestURLSigning(t *testing.T) {
	privateKey := "dummy-key"
	client := imagekit.NewClient(option.WithPrivateKey(privateKey))

	t.Run("should generate a signed URL when signed is true without expiresIn", func(t *testing.T) {
		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint: "https://ik.imagekit.io/demo/",
			Src:         "sdk-testing-files/future-search.png",
			Signed:      param.Opt[bool]{Value: true},
		})

		expected := "https://ik.imagekit.io/demo/sdk-testing-files/future-search.png?ik-s=32dbbbfc5f945c0403c71b54c38e76896ef2d6b0"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate a signed URL when signed is true with expiresIn", func(t *testing.T) {
		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint: "https://ik.imagekit.io/demo/",
			Src:         "sdk-testing-files/future-search.png",
			Signed:      param.Opt[bool]{Value: true},
			ExpiresIn:   param.Opt[float64]{Value: 3600},
		})

		// Expect ik-t exist in the URL. We don't assert signature because it will keep changing.
		if !strings.Contains(url, "ik-t") {
			t.Errorf("Expected URL to contain 'ik-t', got %s", url)
		}
	})

	t.Run("should generate a signed URL when expiresIn is above 0 and even if signed is false", func(t *testing.T) {
		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint: "https://ik.imagekit.io/demo/",
			Src:         "sdk-testing-files/future-search.png",
			Signed:      param.Opt[bool]{Value: false},
			ExpiresIn:   param.Opt[float64]{Value: 3600},
		})

		// Expect ik-t exist in the URL. We don't assert signature because it will keep changing.
		if !strings.Contains(url, "ik-t") {
			t.Errorf("Expected URL to contain 'ik-t', got %s", url)
		}
	})

	t.Run("should generate signed URL with special characters in filename", func(t *testing.T) {
		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint: "https://ik.imagekit.io/demo/",
			Src:         "sdk-testing-files/हिन्दी.png",
			Signed:      param.Opt[bool]{Value: true},
		})

		expected := "https://ik.imagekit.io/demo/sdk-testing-files/%E0%A4%B9%E0%A4%BF%E0%A4%A8%E0%A5%8D%E0%A4%A6%E0%A5%80.png?ik-s=3fff2f31da1f45e007adcdbe95f88c8c330e743c"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate signed URL with text overlay containing special characters", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Overlay: shared.OverlayUnionParam{
					OfText: &shared.TextOverlayParam{
						Text: "हिन्दी",
						Type: "text",
						Transformation: []shared.TextOverlayTransformationParam{
							{
								FontColor: param.Opt[string]{Value: "red"},
								FontSize: shared.TextOverlayTransformationFontSizeUnionParam{
									OfString: param.Opt[string]{Value: "32"},
								},
								FontFamily: param.Opt[string]{Value: "sdk-testing-files/Poppins-Regular_Q15GrYWmL.ttf"},
							},
						},
					},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:    "https://ik.imagekit.io/demo/",
			Src:            "sdk-testing-files/हिन्दी.png",
			Transformation: transformation,
			Signed:         param.Opt[bool]{Value: true},
		})

		expected := "https://ik.imagekit.io/demo/sdk-testing-files/%E0%A4%B9%E0%A4%BF%E0%A4%A8%E0%A5%8D%E0%A4%A6%E0%A5%80.png?tr=l-text,ie-4KS54KS%2F4KSo4KWN4KSm4KWA,fs-32,ff-sdk-testing-files@@Poppins-Regular_Q15GrYWmL.ttf,co-red,l-end&ik-s=705e41579d368caa6530a4375355325277fcfe5c"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate signed URL with text overlay and special characters using path transformation position", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Overlay: shared.OverlayUnionParam{
					OfText: &shared.TextOverlayParam{
						Text: "हिन्दी",
						Type: "text",
						Transformation: []shared.TextOverlayTransformationParam{
							{
								FontColor: param.Opt[string]{Value: "red"},
								FontSize: shared.TextOverlayTransformationFontSizeUnionParam{
									OfString: param.Opt[string]{Value: "32"},
								},
								FontFamily: param.Opt[string]{Value: "sdk-testing-files/Poppins-Regular_Q15GrYWmL.ttf"},
							},
						},
					},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/demo/",
			Src:                    "sdk-testing-files/हिन्दी.png",
			TransformationPosition: shared.TransformationPositionPath,
			Transformation:         transformation,
			Signed:                 param.Opt[bool]{Value: true},
		})

		expected := "https://ik.imagekit.io/demo/tr:l-text,ie-4KS54KS%2F4KSo4KWN4KSm4KWA,fs-32,ff-sdk-testing-files@@Poppins-Regular_Q15GrYWmL.ttf,co-red,l-end/sdk-testing-files/%E0%A4%B9%E0%A4%BF%E0%A4%A8%E0%A5%8D%E0%A4%A6%E0%A5%80.png?ik-s=20958f6126fd67c90653f55a49f2b7bb938d9d1c"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate signed URL with query parameters", func(t *testing.T) {
		queryParams := map[string]string{
			"version": "1.0",
			"cache":   "false",
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:     "https://ik.imagekit.io/demo/",
			Src:             "sdk-testing-files/future-search.png",
			QueryParameters: queryParams,
			Signed:          param.Opt[bool]{Value: true},
		})

		expected := "https://ik.imagekit.io/demo/sdk-testing-files/future-search.png?cache=false&version=1.0&ik-s=03767bb6f0898c04e42f65714af65d937c696d66"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate signed URL with transformations and query parameters", func(t *testing.T) {
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

		queryParams := map[string]string{
			"version": "2.0",
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:     "https://ik.imagekit.io/demo/",
			Src:             "sdk-testing-files/future-search.png",
			Transformation:  transformation,
			QueryParameters: queryParams,
			Signed:          param.Opt[bool]{Value: true},
		})

		expected := "https://ik.imagekit.io/demo/sdk-testing-files/future-search.png?version=2.0&tr=w-300,h-200&ik-s=601d97a7834b7554f4dabf0d3fc3a219ceeb6b31"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should not sign URL when signed is false", func(t *testing.T) {
		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint: "https://ik.imagekit.io/demo/",
			Src:         "sdk-testing-files/future-search.png",
			Signed:      param.Opt[bool]{Value: false},
		})

		expected := "https://ik.imagekit.io/demo/sdk-testing-files/future-search.png"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
		if strings.Contains(url, "ik-s=") {
			t.Errorf("Expected URL to not contain 'ik-s=', got %s", url)
		}
		if strings.Contains(url, "ik-t=") {
			t.Errorf("Expected URL to not contain 'ik-t=', got %s", url)
		}
	})

	t.Run("should generate signed URL with transformations in path position and query parameters", func(t *testing.T) {
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

		queryParams := map[string]string{
			"version": "2.0",
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/demo/",
			Src:                    "sdk-testing-files/future-search.png",
			Transformation:         transformation,
			TransformationPosition: shared.TransformationPositionPath,
			QueryParameters:        queryParams,
			Signed:                 param.Opt[bool]{Value: true},
		})

		expected := "https://ik.imagekit.io/demo/tr:w-300,h-200/sdk-testing-files/future-search.png?version=2.0&ik-s=dd1ee8f83d019bc59fd57a5fc4674a11eb8a3496"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})
}
