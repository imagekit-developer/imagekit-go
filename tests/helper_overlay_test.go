package imagekit_test

import (
	"testing"

	"github.com/imagekit-developer/imagekit-go/v2"
	"github.com/imagekit-developer/imagekit-go/v2/option"

	"github.com/imagekit-developer/imagekit-go/v2/packages/param"
	"github.com/imagekit-developer/imagekit-go/v2/shared"
	"github.com/imagekit-developer/imagekit-go/v2/shared/constant"
)

func TestOverlayTransformations(t *testing.T) {
	privateKey := "My Private API Key"
	client := imagekit.NewClient(option.WithPrivateKey(privateKey))

	t.Run("should ignore overlay when type property is missing", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Width: shared.TransformationWidthUnionParam{
					OfFloat: param.Opt[float64]{Value: 300},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionPath,
			Src:                    "/base-image.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/tr:w-300/base-image.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should ignore text overlay when text property is missing", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Overlay: shared.OverlayUnionParam{
					OfText: &shared.TextOverlayParam{
						Text: "",
						Type: constant.Text("text"),
					},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionPath,
			Src:                    "/base-image.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/base-image.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should ignore image overlay when input property is missing", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Overlay: shared.OverlayUnionParam{
					OfImage: &shared.ImageOverlayParam{
						Input: "",
						Type:  constant.Image("image"),
					},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionPath,
			Src:                    "/base-image.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/base-image.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should ignore video overlay when input property is missing", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Overlay: shared.OverlayUnionParam{
					OfVideo: &shared.VideoOverlayParam{
						Input: "",
						Type:  constant.Video("video"),
					},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionPath,
			Src:                    "/base-image.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/base-image.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should ignore subtitle overlay when input property is missing", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Overlay: shared.OverlayUnionParam{
					OfSubtitle: &shared.SubtitleOverlayParam{
						Input: "",
						Type:  constant.Subtitle("subtitle"),
					},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionPath,
			Src:                    "/base-image.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/base-image.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should ignore solid color overlay when color property is missing", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Overlay: shared.OverlayUnionParam{
					OfSolidColor: &shared.SolidColorOverlayParam{
						Color: "",
						Type:  constant.SolidColor("solidColor"),
					},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionPath,
			Src:                    "/base-image.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/base-image.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate URL with text overlay using URL encoding", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Overlay: shared.OverlayUnionParam{
					OfText: &shared.TextOverlayParam{
						Text: "Minimal Text",
						Type: constant.Text("text"),
					},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionPath,
			Src:                    "/base-image.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/tr:l-text,i-Minimal%20Text,l-end/base-image.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate URL with image overlay from input file", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Overlay: shared.OverlayUnionParam{
					OfImage: &shared.ImageOverlayParam{
						Input: "logo.png",
						Type:  constant.Image("image"),
					},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionPath,
			Src:                    "/base-image.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/tr:l-image,i-logo.png,l-end/base-image.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate URL with video overlay from input file", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Overlay: shared.OverlayUnionParam{
					OfVideo: &shared.VideoOverlayParam{
						Input: "play-pause-loop.mp4",
						Type:  constant.Video("video"),
					},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionPath,
			Src:                    "/base-video.mp4",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/tr:l-video,i-play-pause-loop.mp4,l-end/base-video.mp4"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate URL with subtitle overlay from input file", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Overlay: shared.OverlayUnionParam{
					OfSubtitle: &shared.SubtitleOverlayParam{
						Input: "subtitle.srt",
						Type:  constant.Subtitle("subtitle"),
					},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionPath,
			Src:                    "/base-video.mp4",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/tr:l-subtitle,i-subtitle.srt,l-end/base-video.mp4"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate URL with solid color overlay using background color", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Overlay: shared.OverlayUnionParam{
					OfSolidColor: &shared.SolidColorOverlayParam{
						Color: "FF0000",
						Type:  constant.SolidColor("solidColor"),
					},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionPath,
			Src:                    "/base-image.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/tr:l-image,i-ik_canvas,bg-FF0000,l-end/base-image.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should generate URL with multiple complex overlays including nested transformations", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				// Text overlay
				Overlay: shared.OverlayUnionParam{
					OfText: &shared.TextOverlayParam{
						Text: "Every thing",
						Type: constant.Text("text"),
						BaseOverlayParam: shared.BaseOverlayParam{
							Position: shared.OverlayPositionParam{
								X:     shared.OverlayPositionXUnionParam{OfString: param.Opt[string]{Value: "10"}},
								Y:     shared.OverlayPositionYUnionParam{OfString: param.Opt[string]{Value: "20"}},
								Focus: shared.OverlayPositionFocusCenter,
							},
							Timing: shared.OverlayTimingParam{
								Start:    shared.OverlayTimingStartUnionParam{OfFloat: param.Opt[float64]{Value: 5}},
								Duration: shared.OverlayTimingDurationUnionParam{OfString: param.Opt[string]{Value: "10"}},
								End:      shared.OverlayTimingEndUnionParam{OfFloat: param.Opt[float64]{Value: 15}},
							},
						},
						Transformation: []shared.TextOverlayTransformationParam{
							{
								Width:          shared.TextOverlayTransformationWidthUnionParam{OfString: param.Opt[string]{Value: "bw_mul_0.5"}},
								FontSize:       shared.TextOverlayTransformationFontSizeUnionParam{OfFloat: param.Opt[float64]{Value: 20}},
								FontFamily:     param.Opt[string]{Value: "Arial"},
								FontColor:      param.Opt[string]{Value: "0000ff"},
								InnerAlignment: shared.TextOverlayTransformationInnerAlignmentLeft,
								Padding:        shared.TextOverlayTransformationPaddingUnionParam{OfFloat: param.Opt[float64]{Value: 5}},
								Alpha:          param.Opt[float64]{Value: 7},
								Typography:     param.Opt[string]{Value: "b"},
								Background:     param.Opt[string]{Value: "red"},
								Radius:         shared.TextOverlayTransformationRadiusUnionParam{OfFloat: param.Opt[float64]{Value: 10}},
								Rotation:       shared.TextOverlayTransformationRotationUnionParam{OfString: param.Opt[string]{Value: "N45"}},
								Flip:           shared.TextOverlayTransformationFlipH,
								LineHeight:     shared.TextOverlayTransformationLineHeightUnionParam{OfFloat: param.Opt[float64]{Value: 20}},
							},
						},
					},
				},
			},
			{
				// Image overlay
				Overlay: shared.OverlayUnionParam{
					OfImage: &shared.ImageOverlayParam{
						Input: "logo.png",
						Type:  constant.Image("image"),
						BaseOverlayParam: shared.BaseOverlayParam{
							Position: shared.OverlayPositionParam{
								X:     shared.OverlayPositionXUnionParam{OfString: param.Opt[string]{Value: "10"}},
								Y:     shared.OverlayPositionYUnionParam{OfString: param.Opt[string]{Value: "20"}},
								Focus: shared.OverlayPositionFocusCenter,
							},
							Timing: shared.OverlayTimingParam{
								Start:    shared.OverlayTimingStartUnionParam{OfFloat: param.Opt[float64]{Value: 5}},
								Duration: shared.OverlayTimingDurationUnionParam{OfString: param.Opt[string]{Value: "10"}},
								End:      shared.OverlayTimingEndUnionParam{OfFloat: param.Opt[float64]{Value: 15}},
							},
						},
						Transformation: []shared.TransformationParam{
							{
								Width:    shared.TransformationWidthUnionParam{OfString: param.Opt[string]{Value: "bw_mul_0.5"}},
								Height:   shared.TransformationHeightUnionParam{OfString: param.Opt[string]{Value: "bh_mul_0.5"}},
								Rotation: shared.TransformationRotationUnionParam{OfString: param.Opt[string]{Value: "N45"}},
								Flip:     shared.TransformationFlipH,
								Overlay: shared.OverlayUnionParam{
									OfText: &shared.TextOverlayParam{
										Text: "Nested text overlay",
										Type: constant.Text("text"),
									},
								},
							},
						},
					},
				},
			},
			{
				// Video overlay. Just for URL generation testing, you can't actually overlay a video on an image.
				Overlay: shared.OverlayUnionParam{
					OfVideo: &shared.VideoOverlayParam{
						Input: "play-pause-loop.mp4",
						Type:  constant.Video("video"),
						BaseOverlayParam: shared.BaseOverlayParam{
							Position: shared.OverlayPositionParam{
								X:     shared.OverlayPositionXUnionParam{OfString: param.Opt[string]{Value: "10"}},
								Y:     shared.OverlayPositionYUnionParam{OfString: param.Opt[string]{Value: "20"}},
								Focus: shared.OverlayPositionFocusCenter,
							},
							Timing: shared.OverlayTimingParam{
								Start:    shared.OverlayTimingStartUnionParam{OfFloat: param.Opt[float64]{Value: 5}},
								Duration: shared.OverlayTimingDurationUnionParam{OfString: param.Opt[string]{Value: "10"}},
								End:      shared.OverlayTimingEndUnionParam{OfFloat: param.Opt[float64]{Value: 15}},
							},
						},
						Transformation: []shared.TransformationParam{
							{
								Width:    shared.TransformationWidthUnionParam{OfString: param.Opt[string]{Value: "bw_mul_0.5"}},
								Height:   shared.TransformationHeightUnionParam{OfString: param.Opt[string]{Value: "bh_mul_0.5"}},
								Rotation: shared.TransformationRotationUnionParam{OfString: param.Opt[string]{Value: "N45"}},
								Flip:     shared.TransformationFlipH,
							},
						},
					},
				},
			},
			{
				// Subtitle overlay. Just for URL generation testing, you can't actually overlay a subtitle on an image.
				Overlay: shared.OverlayUnionParam{
					OfSubtitle: &shared.SubtitleOverlayParam{
						Input: "subtitle.srt",
						Type:  constant.Subtitle("subtitle"),
						BaseOverlayParam: shared.BaseOverlayParam{
							Position: shared.OverlayPositionParam{
								X:     shared.OverlayPositionXUnionParam{OfString: param.Opt[string]{Value: "10"}},
								Y:     shared.OverlayPositionYUnionParam{OfString: param.Opt[string]{Value: "20"}},
								Focus: shared.OverlayPositionFocusCenter,
							},
							Timing: shared.OverlayTimingParam{
								Start:    shared.OverlayTimingStartUnionParam{OfFloat: param.Opt[float64]{Value: 5}},
								Duration: shared.OverlayTimingDurationUnionParam{OfString: param.Opt[string]{Value: "10"}},
								End:      shared.OverlayTimingEndUnionParam{OfFloat: param.Opt[float64]{Value: 15}},
							},
						},
						Transformation: []shared.SubtitleOverlayTransformationParam{
							{
								Background:  param.Opt[string]{Value: "red"},
								Color:       param.Opt[string]{Value: "0000ff"},
								FontFamily:  param.Opt[string]{Value: "Arial"},
								FontOutline: param.Opt[string]{Value: "2_A1CCDD50"},
								FontShadow:  param.Opt[string]{Value: "A1CCDD_3"},
							},
						},
					},
				},
			},
			{
				// Solid color overlay
				Overlay: shared.OverlayUnionParam{
					OfSolidColor: &shared.SolidColorOverlayParam{
						Color: "FF0000",
						Type:  constant.SolidColor("solidColor"),
						BaseOverlayParam: shared.BaseOverlayParam{
							Position: shared.OverlayPositionParam{
								X:     shared.OverlayPositionXUnionParam{OfString: param.Opt[string]{Value: "10"}},
								Y:     shared.OverlayPositionYUnionParam{OfString: param.Opt[string]{Value: "20"}},
								Focus: shared.OverlayPositionFocusCenter,
							},
							Timing: shared.OverlayTimingParam{
								Start:    shared.OverlayTimingStartUnionParam{OfFloat: param.Opt[float64]{Value: 5}},
								Duration: shared.OverlayTimingDurationUnionParam{OfString: param.Opt[string]{Value: "10"}},
								End:      shared.OverlayTimingEndUnionParam{OfFloat: param.Opt[float64]{Value: 15}},
							},
						},
						Transformation: []shared.SolidColorOverlayTransformationParam{
							{
								Width:      shared.SolidColorOverlayTransformationWidthUnionParam{OfString: param.Opt[string]{Value: "bw_mul_0.5"}},
								Height:     shared.SolidColorOverlayTransformationHeightUnionParam{OfString: param.Opt[string]{Value: "bh_mul_0.5"}},
								Alpha:      param.Opt[float64]{Value: 0.5},
								Background: param.Opt[string]{Value: "red"},
								Gradient:   shared.SolidColorOverlayTransformationGradientUnionParam{OfSolidColorOverlayTransformationGradientBoolean: param.Opt[bool]{Value: true}},
								Radius:     shared.SolidColorOverlayTransformationRadiusUnionParam{OfMax: constant.Max("max")},
							},
						},
					},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/test_url_endpoint",
			TransformationPosition: shared.TransformationPositionPath,
			Src:                    "/base-image.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/test_url_endpoint/tr:l-text,i-Every%20thing,lx-10,ly-20,lfo-center,lso-5,leo-15,ldu-10,w-bw_mul_0.5,fs-20,ff-Arial,co-0000ff,ia-left,pa-5,al-7,tg-b,bg-red,r-10,rt-N45,fl-h,lh-20,l-end:l-image,i-logo.png,lx-10,ly-20,lfo-center,lso-5,leo-15,ldu-10,w-bw_mul_0.5,h-bh_mul_0.5,rt-N45,fl-h,l-text,i-Nested%20text%20overlay,l-end,l-end:l-video,i-play-pause-loop.mp4,lx-10,ly-20,lfo-center,lso-5,leo-15,ldu-10,w-bw_mul_0.5,h-bh_mul_0.5,rt-N45,fl-h,l-end:l-subtitle,i-subtitle.srt,lx-10,ly-20,lfo-center,lso-5,leo-15,ldu-10,bg-red,co-0000ff,ff-Arial,fol-2_A1CCDD50,fsh-A1CCDD_3,l-end:l-image,i-ik_canvas,bg-FF0000,lx-10,ly-20,lfo-center,lso-5,leo-15,ldu-10,w-bw_mul_0.5,h-bh_mul_0.5,al-0.5,bg-red,e-gradient,r-max,l-end/base-image.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})
}

func TestOverlayEncoding(t *testing.T) {
	privateKey := "My Private API Key"
	client := imagekit.NewClient(option.WithPrivateKey(privateKey))

	t.Run("should use plain encoding for simple image paths with slashes converted to @@", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Overlay: shared.OverlayUnionParam{
					OfImage: &shared.ImageOverlayParam{
						Input: "/customer_logo/nykaa.png",
						Type:  constant.Image("image"),
					},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/demo",
			TransformationPosition: shared.TransformationPositionPath,
			Src:                    "/medium_cafe_B1iTdD0C.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/demo/tr:l-image,i-customer_logo@@nykaa.png,l-end/medium_cafe_B1iTdD0C.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should use base64 encoding for image paths containing special characters", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Overlay: shared.OverlayUnionParam{
					OfImage: &shared.ImageOverlayParam{
						Input: "/customer_logo/Ñykaa.png",
						Type:  constant.Image("image"),
					},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/demo",
			TransformationPosition: shared.TransformationPositionPath,
			Src:                    "/medium_cafe_B1iTdD0C.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/demo/tr:l-image,ie-Y3VzdG9tZXJfbG9nby%2FDkXlrYWEucG5n,l-end/medium_cafe_B1iTdD0C.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should use plain encoding for simple text overlays", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Overlay: shared.OverlayUnionParam{
					OfText: &shared.TextOverlayParam{
						Text: "Manu",
						Type: constant.Text("text"),
					},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/demo",
			TransformationPosition: shared.TransformationPositionPath,
			Src:                    "/medium_cafe_B1iTdD0C.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/demo/tr:l-text,i-Manu,l-end/medium_cafe_B1iTdD0C.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should convert slashes to @@ in fontFamily paths for custom fonts", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Overlay: shared.OverlayUnionParam{
					OfText: &shared.TextOverlayParam{
						Text: "Manu",
						Type: constant.Text("text"),
						Transformation: []shared.TextOverlayTransformationParam{
							{
								FontFamily: param.Opt[string]{Value: "nested-path/Poppins-Regular_Q15GrYWmL.ttf"},
							},
						},
					},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/demo",
			TransformationPosition: shared.TransformationPositionPath,
			Src:                    "/medium_cafe_B1iTdD0C.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/demo/tr:l-text,i-Manu,ff-nested-path@@Poppins-Regular_Q15GrYWmL.ttf,l-end/medium_cafe_B1iTdD0C.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should use URL encoding for text overlays with spaces and safe characters", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Overlay: shared.OverlayUnionParam{
					OfText: &shared.TextOverlayParam{
						Text: "alnum123-._ ",
						Type: constant.Text("text"),
					},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/demo",
			TransformationPosition: shared.TransformationPositionPath,
			Src:                    "/medium_cafe_B1iTdD0C.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/demo/tr:l-text,i-alnum123-._%20,l-end/medium_cafe_B1iTdD0C.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should use base64 encoding for text overlays with special unicode characters", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Overlay: shared.OverlayUnionParam{
					OfText: &shared.TextOverlayParam{
						Text: "Let's use ©, ®, ™, etc",
						Type: constant.Text("text"),
					},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/demo",
			TransformationPosition: shared.TransformationPositionPath,
			Src:                    "/medium_cafe_B1iTdD0C.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/demo/tr:l-text,ie-TGV0J3MgdXNlIMKpLCDCriwg4oSiLCBldGM%3D,l-end/medium_cafe_B1iTdD0C.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should use plain encoding when explicitly specified for text overlay", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Overlay: shared.OverlayUnionParam{
					OfText: &shared.TextOverlayParam{
						Text:     "HelloWorld",
						Type:     constant.Text("text"),
						Encoding: "plain",
					},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/demo",
			TransformationPosition: shared.TransformationPositionPath,
			Src:                    "/sample.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/demo/tr:l-text,i-HelloWorld,l-end/sample.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should use base64 encoding when explicitly specified for text overlay", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Overlay: shared.OverlayUnionParam{
					OfText: &shared.TextOverlayParam{
						Text:     "HelloWorld",
						Type:     constant.Text("text"),
						Encoding: "base64",
					},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/demo",
			TransformationPosition: shared.TransformationPositionPath,
			Src:                    "/sample.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/demo/tr:l-text,ie-SGVsbG9Xb3JsZA%3D%3D,l-end/sample.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should use plain encoding when explicitly specified for image overlay", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Overlay: shared.OverlayUnionParam{
					OfImage: &shared.ImageOverlayParam{
						Input:    "/customer/logo.png",
						Type:     constant.Image("image"),
						Encoding: "plain",
					},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/demo",
			TransformationPosition: shared.TransformationPositionPath,
			Src:                    "/sample.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/demo/tr:l-image,i-customer@@logo.png,l-end/sample.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should use base64 encoding when explicitly specified for image overlay", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Overlay: shared.OverlayUnionParam{
					OfImage: &shared.ImageOverlayParam{
						Input:    "/customer/logo.png",
						Type:     constant.Image("image"),
						Encoding: "base64",
					},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/demo",
			TransformationPosition: shared.TransformationPositionPath,
			Src:                    "/sample.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/demo/tr:l-image,ie-Y3VzdG9tZXIvbG9nby5wbmc%3D,l-end/sample.jpg"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should use base64 encoding when explicitly specified for video overlay", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Overlay: shared.OverlayUnionParam{
					OfVideo: &shared.VideoOverlayParam{
						Input:    "/path/to/video.mp4",
						Type:     constant.Video("video"),
						Encoding: "base64",
					},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/demo",
			TransformationPosition: shared.TransformationPositionPath,
			Src:                    "/sample.mp4",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/demo/tr:l-video,ie-cGF0aC90by92aWRlby5tcDQ%3D,l-end/sample.mp4"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should use plain encoding when explicitly specified for subtitle overlay", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Overlay: shared.OverlayUnionParam{
					OfSubtitle: &shared.SubtitleOverlayParam{
						Input:    "/sub.srt",
						Type:     constant.Subtitle("subtitle"),
						Encoding: "plain",
					},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/demo",
			TransformationPosition: shared.TransformationPositionPath,
			Src:                    "/sample.mp4",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/demo/tr:l-subtitle,i-sub.srt,l-end/sample.mp4"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should use base64 encoding when explicitly specified for subtitle overlay", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Overlay: shared.OverlayUnionParam{
					OfSubtitle: &shared.SubtitleOverlayParam{
						Input:    "sub.srt",
						Type:     constant.Subtitle("subtitle"),
						Encoding: "base64",
					},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/demo",
			TransformationPosition: shared.TransformationPositionPath,
			Src:                    "/sample.mp4",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/demo/tr:l-subtitle,ie-c3ViLnNydA%3D%3D,l-end/sample.mp4"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("should properly encode overlay text when transformations are in query parameters", func(t *testing.T) {
		transformation := []shared.TransformationParam{
			{
				Overlay: shared.OverlayUnionParam{
					OfText: &shared.TextOverlayParam{
						Text: "Minimal Text",
						Type: constant.Text("text"),
					},
				},
			},
		}

		url := client.Helper.BuildURL(shared.SrcOptionsParam{
			URLEndpoint:            "https://ik.imagekit.io/demo",
			TransformationPosition: shared.TransformationPositionQuery,
			Src:                    "/sample.jpg",
			Transformation:         transformation,
		})

		expected := "https://ik.imagekit.io/demo/sample.jpg?tr=l-text,i-Minimal%20Text,l-end"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})
}
