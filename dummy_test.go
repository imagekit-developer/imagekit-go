// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package imagekit_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/imagekit-developer/imagekit-go/v2"
	"github.com/imagekit-developer/imagekit-go/v2/internal/testutil"
	"github.com/imagekit-developer/imagekit-go/v2/option"
	"github.com/imagekit-developer/imagekit-go/v2/shared"
)

func TestDummyNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := imagekit.NewClient(
		option.WithBaseURL(baseURL),
		option.WithPrivateKey("My Private Key"),
		option.WithPassword("My Password"),
	)
	err := client.Dummy.New(context.TODO(), imagekit.DummyNewParams{
		BaseOverlay: shared.BaseOverlayParam{
			LayerMode: shared.BaseOverlayLayerModeMultiply,
			Position: shared.OverlayPositionParam{
				Focus: shared.OverlayPositionFocusCenter,
				X: shared.OverlayPositionXUnionParam{
					OfFloat: imagekit.Float(0),
				},
				Y: shared.OverlayPositionYUnionParam{
					OfFloat: imagekit.Float(0),
				},
			},
			Timing: shared.OverlayTimingParam{
				Duration: shared.OverlayTimingDurationUnionParam{
					OfFloat: imagekit.Float(0),
				},
				End: shared.OverlayTimingEndUnionParam{
					OfFloat: imagekit.Float(0),
				},
				Start: shared.OverlayTimingStartUnionParam{
					OfFloat: imagekit.Float(0),
				},
			},
		},
		ExtensionConfig: shared.ExtensionConfigUnionParam{
			OfRemoveBg: &shared.ExtensionConfigRemoveBgParam{
				Options: shared.ExtensionConfigRemoveBgOptionsParam{
					AddShadow:        imagekit.Bool(true),
					BgColor:          imagekit.String("bg_color"),
					BgImageURL:       imagekit.String("bg_image_url"),
					Semitransparency: imagekit.Bool(true),
				},
			},
		},
		Extensions: shared.ExtensionsParam{shared.ExtensionUnionParam{
			OfRemoveBg: &shared.ExtensionRemoveBgParam{
				Options: shared.ExtensionRemoveBgOptionsParam{
					AddShadow:        imagekit.Bool(true),
					BgColor:          imagekit.String("bg_color"),
					BgImageURL:       imagekit.String("bg_image_url"),
					Semitransparency: imagekit.Bool(true),
				},
			},
		}, shared.ExtensionUnionParam{
			OfAutoTagging: &shared.ExtensionAutoTaggingParam{
				MaxTags:       5,
				MinConfidence: 95,
				Name:          "google-auto-tagging",
			},
		}, shared.ExtensionUnionParam{
			OfAIAutoDescription: &shared.ExtensionAIAutoDescriptionParam{},
		}, shared.ExtensionUnionParam{
			OfAITasks: &shared.ExtensionAITasksParam{
				Tasks: []shared.ExtensionAITasksTaskUnionParam{{
					OfSelectTags: &shared.ExtensionAITasksTaskSelectTagsParam{
						Instruction:   "What types of clothing items are visible in this image?",
						Vocabulary:    []string{"shirt", "tshirt", "dress", "trousers", "jacket"},
						MaxSelections: imagekit.Int(1),
						MinSelections: imagekit.Int(0),
					},
				}, {
					OfYesNo: &shared.ExtensionAITasksTaskYesNoParam{
						Instruction: "Is this a luxury or high-end fashion item?",
						OnNo: shared.ExtensionAITasksTaskYesNoOnNoParam{
							AddTags:    []string{"luxury", "premium"},
							RemoveTags: []string{"budget", "affordable"},
							SetMetadata: []shared.ExtensionAITasksTaskYesNoOnNoSetMetadataParam{{
								Field: "price_range",
								Value: shared.ExtensionAITasksTaskYesNoOnNoSetMetadataValueUnionParam{
									OfString: imagekit.String("premium"),
								},
							}},
							UnsetMetadata: []shared.ExtensionAITasksTaskYesNoOnNoUnsetMetadataParam{{
								Field: "price_range",
							}},
						},
						OnUnknown: shared.ExtensionAITasksTaskYesNoOnUnknownParam{
							AddTags:    []string{"luxury", "premium"},
							RemoveTags: []string{"budget", "affordable"},
							SetMetadata: []shared.ExtensionAITasksTaskYesNoOnUnknownSetMetadataParam{{
								Field: "price_range",
								Value: shared.ExtensionAITasksTaskYesNoOnUnknownSetMetadataValueUnionParam{
									OfString: imagekit.String("premium"),
								},
							}},
							UnsetMetadata: []shared.ExtensionAITasksTaskYesNoOnUnknownUnsetMetadataParam{{
								Field: "price_range",
							}},
						},
						OnYes: shared.ExtensionAITasksTaskYesNoOnYesParam{
							AddTags:    []string{"luxury", "premium"},
							RemoveTags: []string{"budget", "affordable"},
							SetMetadata: []shared.ExtensionAITasksTaskYesNoOnYesSetMetadataParam{{
								Field: "price_range",
								Value: shared.ExtensionAITasksTaskYesNoOnYesSetMetadataValueUnionParam{
									OfString: imagekit.String("premium"),
								},
							}},
							UnsetMetadata: []shared.ExtensionAITasksTaskYesNoOnYesUnsetMetadataParam{{
								Field: "price_range",
							}},
						},
					},
				}},
			},
		}, shared.ExtensionUnionParam{
			OfSavedExtension: &shared.ExtensionSavedExtensionParam{
				ID: "ext_abc123",
			},
		}},
		GetImageAttributesOptions: shared.GetImageAttributesOptionsParam{
			SrcOptionsParam: shared.SrcOptionsParam{
				Src:         "/my-image.jpg",
				URLEndpoint: "https://ik.imagekit.io/demo",
				ExpiresIn:   imagekit.Float(0),
				QueryParameters: map[string]string{
					"foo": "string",
				},
				Signed: imagekit.Bool(true),
				Transformation: []shared.TransformationParam{{
					AIChangeBackground: imagekit.String("aiChangeBackground"),
					AIDropShadow: shared.TransformationAIDropShadowUnionParam{
						OfTransformationAIDropShadowBoolean: imagekit.Bool(true),
					},
					AIEdit:                     imagekit.String("aiEdit"),
					AIRemoveBackground:         true,
					AIRemoveBackgroundExternal: true,
					AIRetouch:                  true,
					AIUpscale:                  true,
					AIVariation:                true,
					AspectRatio: shared.TransformationAspectRatioUnionParam{
						OfString: imagekit.String("4:3"),
					},
					AudioCodec:      shared.TransformationAudioCodecAac,
					Background:      imagekit.String("red"),
					Blur:            imagekit.Float(10),
					Border:          imagekit.String("5_FF0000"),
					ColorProfile:    imagekit.Bool(true),
					ColorReplace:    imagekit.String("colorReplace"),
					ContrastStretch: true,
					Crop:            shared.TransformationCropForce,
					CropMode:        shared.TransformationCropModePadResize,
					DefaultImage:    imagekit.String("defaultImage"),
					Distort:         imagekit.String("distort"),
					Dpr:             imagekit.Float(2),
					Duration: shared.TransformationDurationUnionParam{
						OfFloat: imagekit.Float(0),
					},
					EndOffset: shared.TransformationEndOffsetUnionParam{
						OfFloat: imagekit.Float(0),
					},
					Flip:   shared.TransformationFlipH,
					Focus:  imagekit.String("center"),
					Format: shared.TransformationFormatAuto,
					Gradient: shared.TransformationGradientUnionParam{
						OfTransformationGradientBoolean: imagekit.Bool(true),
					},
					Grayscale: true,
					Height: shared.TransformationHeightUnionParam{
						OfFloat: imagekit.Float(200),
					},
					Lossless: imagekit.Bool(true),
					Metadata: imagekit.Bool(true),
					Named:    imagekit.String("named"),
					Opacity:  imagekit.Float(0),
					Original: imagekit.Bool(true),
					Overlay: shared.OverlayUnionParam{
						OfText: &shared.TextOverlayParam{
							BaseOverlayParam: shared.BaseOverlayParam{
								LayerMode: shared.BaseOverlayLayerModeMultiply,
								Position: shared.OverlayPositionParam{
									Focus: shared.OverlayPositionFocusCenter,
									X: shared.OverlayPositionXUnionParam{
										OfFloat: imagekit.Float(0),
									},
									Y: shared.OverlayPositionYUnionParam{
										OfFloat: imagekit.Float(0),
									},
								},
								Timing: shared.OverlayTimingParam{
									Duration: shared.OverlayTimingDurationUnionParam{
										OfFloat: imagekit.Float(0),
									},
									End: shared.OverlayTimingEndUnionParam{
										OfFloat: imagekit.Float(0),
									},
									Start: shared.OverlayTimingStartUnionParam{
										OfFloat: imagekit.Float(0),
									},
								},
							},
							Text:     "text",
							Encoding: "auto",
							Transformation: []shared.TextOverlayTransformationParam{{
								Alpha:      imagekit.Float(1),
								Background: imagekit.String("background"),
								Flip:       shared.TextOverlayTransformationFlipH,
								FontColor:  imagekit.String("fontColor"),
								FontFamily: imagekit.String("fontFamily"),
								FontSize: shared.TextOverlayTransformationFontSizeUnionParam{
									OfFloat: imagekit.Float(0),
								},
								InnerAlignment: shared.TextOverlayTransformationInnerAlignmentLeft,
								LineHeight: shared.TextOverlayTransformationLineHeightUnionParam{
									OfFloat: imagekit.Float(0),
								},
								Padding: shared.TextOverlayTransformationPaddingUnionParam{
									OfFloat: imagekit.Float(0),
								},
								Radius: shared.TextOverlayTransformationRadiusUnionParam{
									OfFloat: imagekit.Float(0),
								},
								Rotation: shared.TextOverlayTransformationRotationUnionParam{
									OfFloat: imagekit.Float(0),
								},
								Typography: imagekit.String("typography"),
								Width: shared.TextOverlayTransformationWidthUnionParam{
									OfFloat: imagekit.Float(0),
								},
							}},
						},
					},
					Page: shared.TransformationPageUnionParam{
						OfFloat: imagekit.Float(0),
					},
					Progressive: imagekit.Bool(true),
					Quality:     imagekit.Float(80),
					Radius: shared.TransformationRadiusUnionParam{
						OfFloat: imagekit.Float(20),
					},
					Raw: imagekit.String("raw"),
					Rotation: shared.TransformationRotationUnionParam{
						OfFloat: imagekit.Float(90),
					},
					Shadow: shared.TransformationShadowUnionParam{
						OfTransformationShadowBoolean: imagekit.Bool(true),
					},
					Sharpen: shared.TransformationSharpenUnionParam{
						OfTransformationSharpenBoolean: imagekit.Bool(true),
					},
					StartOffset: shared.TransformationStartOffsetUnionParam{
						OfFloat: imagekit.Float(0),
					},
					StreamingResolutions: []shared.StreamingResolution{shared.StreamingResolution240},
					Trim: shared.TransformationTrimUnionParam{
						OfTransformationTrimBoolean: imagekit.Bool(true),
					},
					UnsharpMask: shared.TransformationUnsharpMaskUnionParam{
						OfTransformationUnsharpMaskBoolean: imagekit.Bool(true),
					},
					VideoCodec: shared.TransformationVideoCodecH264,
					Width: shared.TransformationWidthUnionParam{
						OfFloat: imagekit.Float(300),
					},
					X: shared.TransformationXUnionParam{
						OfFloat: imagekit.Float(0),
					},
					XCenter: shared.TransformationXCenterUnionParam{
						OfFloat: imagekit.Float(0),
					},
					Y: shared.TransformationYUnionParam{
						OfFloat: imagekit.Float(0),
					},
					YCenter: shared.TransformationYCenterUnionParam{
						OfFloat: imagekit.Float(0),
					},
					Zoom: imagekit.Float(0),
				}},
				TransformationPosition: shared.TransformationPositionPath,
			},
			DeviceBreakpoints: []float64{640, 750, 828, 1080, 1200, 1920, 2048, 3840},
			ImageBreakpoints:  []float64{16, 32, 48, 64, 96, 128, 256, 384},
			Sizes:             imagekit.String("(min-width: 768px) 50vw, 100vw"),
			Width:             imagekit.Float(400),
		},
		ImageOverlay: shared.ImageOverlayParam{
			BaseOverlayParam: shared.BaseOverlayParam{
				LayerMode: shared.BaseOverlayLayerModeMultiply,
				Position: shared.OverlayPositionParam{
					Focus: shared.OverlayPositionFocusCenter,
					X: shared.OverlayPositionXUnionParam{
						OfFloat: imagekit.Float(0),
					},
					Y: shared.OverlayPositionYUnionParam{
						OfFloat: imagekit.Float(0),
					},
				},
				Timing: shared.OverlayTimingParam{
					Duration: shared.OverlayTimingDurationUnionParam{
						OfFloat: imagekit.Float(0),
					},
					End: shared.OverlayTimingEndUnionParam{
						OfFloat: imagekit.Float(0),
					},
					Start: shared.OverlayTimingStartUnionParam{
						OfFloat: imagekit.Float(0),
					},
				},
			},
			Input:    "input",
			Encoding: "auto",
			Transformation: []shared.TransformationParam{{
				AIChangeBackground: imagekit.String("aiChangeBackground"),
				AIDropShadow: shared.TransformationAIDropShadowUnionParam{
					OfTransformationAIDropShadowBoolean: imagekit.Bool(true),
				},
				AIEdit:                     imagekit.String("aiEdit"),
				AIRemoveBackground:         true,
				AIRemoveBackgroundExternal: true,
				AIRetouch:                  true,
				AIUpscale:                  true,
				AIVariation:                true,
				AspectRatio: shared.TransformationAspectRatioUnionParam{
					OfString: imagekit.String("4:3"),
				},
				AudioCodec:      shared.TransformationAudioCodecAac,
				Background:      imagekit.String("red"),
				Blur:            imagekit.Float(10),
				Border:          imagekit.String("5_FF0000"),
				ColorProfile:    imagekit.Bool(true),
				ColorReplace:    imagekit.String("colorReplace"),
				ContrastStretch: true,
				Crop:            shared.TransformationCropForce,
				CropMode:        shared.TransformationCropModePadResize,
				DefaultImage:    imagekit.String("defaultImage"),
				Distort:         imagekit.String("distort"),
				Dpr:             imagekit.Float(2),
				Duration: shared.TransformationDurationUnionParam{
					OfFloat: imagekit.Float(0),
				},
				EndOffset: shared.TransformationEndOffsetUnionParam{
					OfFloat: imagekit.Float(0),
				},
				Flip:   shared.TransformationFlipH,
				Focus:  imagekit.String("center"),
				Format: shared.TransformationFormatAuto,
				Gradient: shared.TransformationGradientUnionParam{
					OfTransformationGradientBoolean: imagekit.Bool(true),
				},
				Grayscale: true,
				Height: shared.TransformationHeightUnionParam{
					OfFloat: imagekit.Float(200),
				},
				Lossless: imagekit.Bool(true),
				Metadata: imagekit.Bool(true),
				Named:    imagekit.String("named"),
				Opacity:  imagekit.Float(0),
				Original: imagekit.Bool(true),
				Overlay: shared.OverlayUnionParam{
					OfText: &shared.TextOverlayParam{
						BaseOverlayParam: shared.BaseOverlayParam{
							LayerMode: shared.BaseOverlayLayerModeMultiply,
							Position: shared.OverlayPositionParam{
								Focus: shared.OverlayPositionFocusCenter,
								X: shared.OverlayPositionXUnionParam{
									OfFloat: imagekit.Float(0),
								},
								Y: shared.OverlayPositionYUnionParam{
									OfFloat: imagekit.Float(0),
								},
							},
							Timing: shared.OverlayTimingParam{
								Duration: shared.OverlayTimingDurationUnionParam{
									OfFloat: imagekit.Float(0),
								},
								End: shared.OverlayTimingEndUnionParam{
									OfFloat: imagekit.Float(0),
								},
								Start: shared.OverlayTimingStartUnionParam{
									OfFloat: imagekit.Float(0),
								},
							},
						},
						Text:     "text",
						Encoding: "auto",
						Transformation: []shared.TextOverlayTransformationParam{{
							Alpha:      imagekit.Float(1),
							Background: imagekit.String("background"),
							Flip:       shared.TextOverlayTransformationFlipH,
							FontColor:  imagekit.String("fontColor"),
							FontFamily: imagekit.String("fontFamily"),
							FontSize: shared.TextOverlayTransformationFontSizeUnionParam{
								OfFloat: imagekit.Float(0),
							},
							InnerAlignment: shared.TextOverlayTransformationInnerAlignmentLeft,
							LineHeight: shared.TextOverlayTransformationLineHeightUnionParam{
								OfFloat: imagekit.Float(0),
							},
							Padding: shared.TextOverlayTransformationPaddingUnionParam{
								OfFloat: imagekit.Float(0),
							},
							Radius: shared.TextOverlayTransformationRadiusUnionParam{
								OfFloat: imagekit.Float(0),
							},
							Rotation: shared.TextOverlayTransformationRotationUnionParam{
								OfFloat: imagekit.Float(0),
							},
							Typography: imagekit.String("typography"),
							Width: shared.TextOverlayTransformationWidthUnionParam{
								OfFloat: imagekit.Float(0),
							},
						}},
					},
				},
				Page: shared.TransformationPageUnionParam{
					OfFloat: imagekit.Float(0),
				},
				Progressive: imagekit.Bool(true),
				Quality:     imagekit.Float(80),
				Radius: shared.TransformationRadiusUnionParam{
					OfFloat: imagekit.Float(20),
				},
				Raw: imagekit.String("raw"),
				Rotation: shared.TransformationRotationUnionParam{
					OfFloat: imagekit.Float(90),
				},
				Shadow: shared.TransformationShadowUnionParam{
					OfTransformationShadowBoolean: imagekit.Bool(true),
				},
				Sharpen: shared.TransformationSharpenUnionParam{
					OfTransformationSharpenBoolean: imagekit.Bool(true),
				},
				StartOffset: shared.TransformationStartOffsetUnionParam{
					OfFloat: imagekit.Float(0),
				},
				StreamingResolutions: []shared.StreamingResolution{shared.StreamingResolution240},
				Trim: shared.TransformationTrimUnionParam{
					OfTransformationTrimBoolean: imagekit.Bool(true),
				},
				UnsharpMask: shared.TransformationUnsharpMaskUnionParam{
					OfTransformationUnsharpMaskBoolean: imagekit.Bool(true),
				},
				VideoCodec: shared.TransformationVideoCodecH264,
				Width: shared.TransformationWidthUnionParam{
					OfFloat: imagekit.Float(300),
				},
				X: shared.TransformationXUnionParam{
					OfFloat: imagekit.Float(0),
				},
				XCenter: shared.TransformationXCenterUnionParam{
					OfFloat: imagekit.Float(0),
				},
				Y: shared.TransformationYUnionParam{
					OfFloat: imagekit.Float(0),
				},
				YCenter: shared.TransformationYCenterUnionParam{
					OfFloat: imagekit.Float(0),
				},
				Zoom: imagekit.Float(0),
			}},
		},
		Overlay: shared.OverlayUnionParam{
			OfText: &shared.TextOverlayParam{
				BaseOverlayParam: shared.BaseOverlayParam{
					LayerMode: shared.BaseOverlayLayerModeMultiply,
					Position: shared.OverlayPositionParam{
						Focus: shared.OverlayPositionFocusCenter,
						X: shared.OverlayPositionXUnionParam{
							OfFloat: imagekit.Float(0),
						},
						Y: shared.OverlayPositionYUnionParam{
							OfFloat: imagekit.Float(0),
						},
					},
					Timing: shared.OverlayTimingParam{
						Duration: shared.OverlayTimingDurationUnionParam{
							OfFloat: imagekit.Float(0),
						},
						End: shared.OverlayTimingEndUnionParam{
							OfFloat: imagekit.Float(0),
						},
						Start: shared.OverlayTimingStartUnionParam{
							OfFloat: imagekit.Float(0),
						},
					},
				},
				Text:     "text",
				Encoding: "auto",
				Transformation: []shared.TextOverlayTransformationParam{{
					Alpha:      imagekit.Float(1),
					Background: imagekit.String("background"),
					Flip:       shared.TextOverlayTransformationFlipH,
					FontColor:  imagekit.String("fontColor"),
					FontFamily: imagekit.String("fontFamily"),
					FontSize: shared.TextOverlayTransformationFontSizeUnionParam{
						OfFloat: imagekit.Float(0),
					},
					InnerAlignment: shared.TextOverlayTransformationInnerAlignmentLeft,
					LineHeight: shared.TextOverlayTransformationLineHeightUnionParam{
						OfFloat: imagekit.Float(0),
					},
					Padding: shared.TextOverlayTransformationPaddingUnionParam{
						OfFloat: imagekit.Float(0),
					},
					Radius: shared.TextOverlayTransformationRadiusUnionParam{
						OfFloat: imagekit.Float(0),
					},
					Rotation: shared.TextOverlayTransformationRotationUnionParam{
						OfFloat: imagekit.Float(0),
					},
					Typography: imagekit.String("typography"),
					Width: shared.TextOverlayTransformationWidthUnionParam{
						OfFloat: imagekit.Float(0),
					},
				}},
			},
		},
		OverlayPosition: shared.OverlayPositionParam{
			Focus: shared.OverlayPositionFocusCenter,
			X: shared.OverlayPositionXUnionParam{
				OfFloat: imagekit.Float(0),
			},
			Y: shared.OverlayPositionYUnionParam{
				OfFloat: imagekit.Float(0),
			},
		},
		OverlayTiming: shared.OverlayTimingParam{
			Duration: shared.OverlayTimingDurationUnionParam{
				OfFloat: imagekit.Float(0),
			},
			End: shared.OverlayTimingEndUnionParam{
				OfFloat: imagekit.Float(0),
			},
			Start: shared.OverlayTimingStartUnionParam{
				OfFloat: imagekit.Float(0),
			},
		},
		ResponsiveImageAttributes: shared.ResponsiveImageAttributesParam{
			Src:    "https://ik.imagekit.io/demo/image.jpg?tr=w-3840",
			Sizes:  imagekit.String("100vw"),
			SrcSet: imagekit.String("https://ik.imagekit.io/demo/image.jpg?tr=w-640 640w, https://ik.imagekit.io/demo/image.jpg?tr=w-1080 1080w, https://ik.imagekit.io/demo/image.jpg?tr=w-1920 1920w"),
			Width:  imagekit.Float(400),
		},
		SavedExtensions: shared.SavedExtensionParam{
			ID: imagekit.String("ext_abc123"),
			Config: shared.ExtensionConfigUnionParam{
				OfRemoveBg: &shared.ExtensionConfigRemoveBgParam{
					Options: shared.ExtensionConfigRemoveBgOptionsParam{
						AddShadow:        imagekit.Bool(true),
						BgColor:          imagekit.String("bg_color"),
						BgImageURL:       imagekit.String("bg_image_url"),
						Semitransparency: imagekit.Bool(true),
					},
				},
			},
			CreatedAt:   imagekit.Time(time.Now()),
			Description: imagekit.String("Analyzes vehicle images for type, condition, and quality assessment"),
			Name:        imagekit.String("Car Quality Analysis"),
			UpdatedAt:   imagekit.Time(time.Now()),
		},
		SolidColorOverlay: shared.SolidColorOverlayParam{
			BaseOverlayParam: shared.BaseOverlayParam{
				LayerMode: shared.BaseOverlayLayerModeMultiply,
				Position: shared.OverlayPositionParam{
					Focus: shared.OverlayPositionFocusCenter,
					X: shared.OverlayPositionXUnionParam{
						OfFloat: imagekit.Float(0),
					},
					Y: shared.OverlayPositionYUnionParam{
						OfFloat: imagekit.Float(0),
					},
				},
				Timing: shared.OverlayTimingParam{
					Duration: shared.OverlayTimingDurationUnionParam{
						OfFloat: imagekit.Float(0),
					},
					End: shared.OverlayTimingEndUnionParam{
						OfFloat: imagekit.Float(0),
					},
					Start: shared.OverlayTimingStartUnionParam{
						OfFloat: imagekit.Float(0),
					},
				},
			},
			Color: "color",
			Transformation: []shared.SolidColorOverlayTransformationParam{{
				Alpha:      imagekit.Float(1),
				Background: imagekit.String("background"),
				Gradient: shared.SolidColorOverlayTransformationGradientUnionParam{
					OfSolidColorOverlayTransformationGradientBoolean: imagekit.Bool(true),
				},
				Height: shared.SolidColorOverlayTransformationHeightUnionParam{
					OfFloat: imagekit.Float(0),
				},
				Radius: shared.SolidColorOverlayTransformationRadiusUnionParam{
					OfFloat: imagekit.Float(0),
				},
				Width: shared.SolidColorOverlayTransformationWidthUnionParam{
					OfFloat: imagekit.Float(0),
				},
			}},
		},
		SolidColorOverlayTransformation: shared.SolidColorOverlayTransformationParam{
			Alpha:      imagekit.Float(1),
			Background: imagekit.String("background"),
			Gradient: shared.SolidColorOverlayTransformationGradientUnionParam{
				OfSolidColorOverlayTransformationGradientBoolean: imagekit.Bool(true),
			},
			Height: shared.SolidColorOverlayTransformationHeightUnionParam{
				OfFloat: imagekit.Float(0),
			},
			Radius: shared.SolidColorOverlayTransformationRadiusUnionParam{
				OfFloat: imagekit.Float(0),
			},
			Width: shared.SolidColorOverlayTransformationWidthUnionParam{
				OfFloat: imagekit.Float(0),
			},
		},
		SrcOptions: shared.SrcOptionsParam{
			Src:         "/my-image.jpg",
			URLEndpoint: "https://ik.imagekit.io/demo",
			ExpiresIn:   imagekit.Float(0),
			QueryParameters: map[string]string{
				"foo": "string",
			},
			Signed: imagekit.Bool(true),
			Transformation: []shared.TransformationParam{{
				AIChangeBackground: imagekit.String("aiChangeBackground"),
				AIDropShadow: shared.TransformationAIDropShadowUnionParam{
					OfTransformationAIDropShadowBoolean: imagekit.Bool(true),
				},
				AIEdit:                     imagekit.String("aiEdit"),
				AIRemoveBackground:         true,
				AIRemoveBackgroundExternal: true,
				AIRetouch:                  true,
				AIUpscale:                  true,
				AIVariation:                true,
				AspectRatio: shared.TransformationAspectRatioUnionParam{
					OfString: imagekit.String("4:3"),
				},
				AudioCodec:      shared.TransformationAudioCodecAac,
				Background:      imagekit.String("red"),
				Blur:            imagekit.Float(10),
				Border:          imagekit.String("5_FF0000"),
				ColorProfile:    imagekit.Bool(true),
				ColorReplace:    imagekit.String("colorReplace"),
				ContrastStretch: true,
				Crop:            shared.TransformationCropForce,
				CropMode:        shared.TransformationCropModePadResize,
				DefaultImage:    imagekit.String("defaultImage"),
				Distort:         imagekit.String("distort"),
				Dpr:             imagekit.Float(2),
				Duration: shared.TransformationDurationUnionParam{
					OfFloat: imagekit.Float(0),
				},
				EndOffset: shared.TransformationEndOffsetUnionParam{
					OfFloat: imagekit.Float(0),
				},
				Flip:   shared.TransformationFlipH,
				Focus:  imagekit.String("center"),
				Format: shared.TransformationFormatAuto,
				Gradient: shared.TransformationGradientUnionParam{
					OfTransformationGradientBoolean: imagekit.Bool(true),
				},
				Grayscale: true,
				Height: shared.TransformationHeightUnionParam{
					OfFloat: imagekit.Float(200),
				},
				Lossless: imagekit.Bool(true),
				Metadata: imagekit.Bool(true),
				Named:    imagekit.String("named"),
				Opacity:  imagekit.Float(0),
				Original: imagekit.Bool(true),
				Overlay: shared.OverlayUnionParam{
					OfText: &shared.TextOverlayParam{
						BaseOverlayParam: shared.BaseOverlayParam{
							LayerMode: shared.BaseOverlayLayerModeMultiply,
							Position: shared.OverlayPositionParam{
								Focus: shared.OverlayPositionFocusCenter,
								X: shared.OverlayPositionXUnionParam{
									OfFloat: imagekit.Float(0),
								},
								Y: shared.OverlayPositionYUnionParam{
									OfFloat: imagekit.Float(0),
								},
							},
							Timing: shared.OverlayTimingParam{
								Duration: shared.OverlayTimingDurationUnionParam{
									OfFloat: imagekit.Float(0),
								},
								End: shared.OverlayTimingEndUnionParam{
									OfFloat: imagekit.Float(0),
								},
								Start: shared.OverlayTimingStartUnionParam{
									OfFloat: imagekit.Float(0),
								},
							},
						},
						Text:     "text",
						Encoding: "auto",
						Transformation: []shared.TextOverlayTransformationParam{{
							Alpha:      imagekit.Float(1),
							Background: imagekit.String("background"),
							Flip:       shared.TextOverlayTransformationFlipH,
							FontColor:  imagekit.String("fontColor"),
							FontFamily: imagekit.String("fontFamily"),
							FontSize: shared.TextOverlayTransformationFontSizeUnionParam{
								OfFloat: imagekit.Float(0),
							},
							InnerAlignment: shared.TextOverlayTransformationInnerAlignmentLeft,
							LineHeight: shared.TextOverlayTransformationLineHeightUnionParam{
								OfFloat: imagekit.Float(0),
							},
							Padding: shared.TextOverlayTransformationPaddingUnionParam{
								OfFloat: imagekit.Float(0),
							},
							Radius: shared.TextOverlayTransformationRadiusUnionParam{
								OfFloat: imagekit.Float(0),
							},
							Rotation: shared.TextOverlayTransformationRotationUnionParam{
								OfFloat: imagekit.Float(0),
							},
							Typography: imagekit.String("typography"),
							Width: shared.TextOverlayTransformationWidthUnionParam{
								OfFloat: imagekit.Float(0),
							},
						}},
					},
				},
				Page: shared.TransformationPageUnionParam{
					OfFloat: imagekit.Float(0),
				},
				Progressive: imagekit.Bool(true),
				Quality:     imagekit.Float(80),
				Radius: shared.TransformationRadiusUnionParam{
					OfFloat: imagekit.Float(20),
				},
				Raw: imagekit.String("raw"),
				Rotation: shared.TransformationRotationUnionParam{
					OfFloat: imagekit.Float(90),
				},
				Shadow: shared.TransformationShadowUnionParam{
					OfTransformationShadowBoolean: imagekit.Bool(true),
				},
				Sharpen: shared.TransformationSharpenUnionParam{
					OfTransformationSharpenBoolean: imagekit.Bool(true),
				},
				StartOffset: shared.TransformationStartOffsetUnionParam{
					OfFloat: imagekit.Float(0),
				},
				StreamingResolutions: []shared.StreamingResolution{shared.StreamingResolution240},
				Trim: shared.TransformationTrimUnionParam{
					OfTransformationTrimBoolean: imagekit.Bool(true),
				},
				UnsharpMask: shared.TransformationUnsharpMaskUnionParam{
					OfTransformationUnsharpMaskBoolean: imagekit.Bool(true),
				},
				VideoCodec: shared.TransformationVideoCodecH264,
				Width: shared.TransformationWidthUnionParam{
					OfFloat: imagekit.Float(300),
				},
				X: shared.TransformationXUnionParam{
					OfFloat: imagekit.Float(0),
				},
				XCenter: shared.TransformationXCenterUnionParam{
					OfFloat: imagekit.Float(0),
				},
				Y: shared.TransformationYUnionParam{
					OfFloat: imagekit.Float(0),
				},
				YCenter: shared.TransformationYCenterUnionParam{
					OfFloat: imagekit.Float(0),
				},
				Zoom: imagekit.Float(0),
			}},
			TransformationPosition: shared.TransformationPositionPath,
		},
		StreamingResolution: shared.StreamingResolution240,
		SubtitleOverlay: shared.SubtitleOverlayParam{
			BaseOverlayParam: shared.BaseOverlayParam{
				LayerMode: shared.BaseOverlayLayerModeMultiply,
				Position: shared.OverlayPositionParam{
					Focus: shared.OverlayPositionFocusCenter,
					X: shared.OverlayPositionXUnionParam{
						OfFloat: imagekit.Float(0),
					},
					Y: shared.OverlayPositionYUnionParam{
						OfFloat: imagekit.Float(0),
					},
				},
				Timing: shared.OverlayTimingParam{
					Duration: shared.OverlayTimingDurationUnionParam{
						OfFloat: imagekit.Float(0),
					},
					End: shared.OverlayTimingEndUnionParam{
						OfFloat: imagekit.Float(0),
					},
					Start: shared.OverlayTimingStartUnionParam{
						OfFloat: imagekit.Float(0),
					},
				},
			},
			Input:    "input",
			Encoding: "auto",
			Transformation: []shared.SubtitleOverlayTransformationParam{{
				Background:  imagekit.String("background"),
				Color:       imagekit.String("color"),
				FontFamily:  imagekit.String("fontFamily"),
				FontOutline: imagekit.String("fontOutline"),
				FontShadow:  imagekit.String("fontShadow"),
				FontSize:    imagekit.Float(0),
				Typography:  shared.SubtitleOverlayTransformationTypographyB,
			}},
		},
		SubtitleOverlayTransformation: shared.SubtitleOverlayTransformationParam{
			Background:  imagekit.String("background"),
			Color:       imagekit.String("color"),
			FontFamily:  imagekit.String("fontFamily"),
			FontOutline: imagekit.String("fontOutline"),
			FontShadow:  imagekit.String("fontShadow"),
			FontSize:    imagekit.Float(0),
			Typography:  shared.SubtitleOverlayTransformationTypographyB,
		},
		TextOverlay: shared.TextOverlayParam{
			BaseOverlayParam: shared.BaseOverlayParam{
				LayerMode: shared.BaseOverlayLayerModeMultiply,
				Position: shared.OverlayPositionParam{
					Focus: shared.OverlayPositionFocusCenter,
					X: shared.OverlayPositionXUnionParam{
						OfFloat: imagekit.Float(0),
					},
					Y: shared.OverlayPositionYUnionParam{
						OfFloat: imagekit.Float(0),
					},
				},
				Timing: shared.OverlayTimingParam{
					Duration: shared.OverlayTimingDurationUnionParam{
						OfFloat: imagekit.Float(0),
					},
					End: shared.OverlayTimingEndUnionParam{
						OfFloat: imagekit.Float(0),
					},
					Start: shared.OverlayTimingStartUnionParam{
						OfFloat: imagekit.Float(0),
					},
				},
			},
			Text:     "text",
			Encoding: "auto",
			Transformation: []shared.TextOverlayTransformationParam{{
				Alpha:      imagekit.Float(1),
				Background: imagekit.String("background"),
				Flip:       shared.TextOverlayTransformationFlipH,
				FontColor:  imagekit.String("fontColor"),
				FontFamily: imagekit.String("fontFamily"),
				FontSize: shared.TextOverlayTransformationFontSizeUnionParam{
					OfFloat: imagekit.Float(0),
				},
				InnerAlignment: shared.TextOverlayTransformationInnerAlignmentLeft,
				LineHeight: shared.TextOverlayTransformationLineHeightUnionParam{
					OfFloat: imagekit.Float(0),
				},
				Padding: shared.TextOverlayTransformationPaddingUnionParam{
					OfFloat: imagekit.Float(0),
				},
				Radius: shared.TextOverlayTransformationRadiusUnionParam{
					OfFloat: imagekit.Float(0),
				},
				Rotation: shared.TextOverlayTransformationRotationUnionParam{
					OfFloat: imagekit.Float(0),
				},
				Typography: imagekit.String("typography"),
				Width: shared.TextOverlayTransformationWidthUnionParam{
					OfFloat: imagekit.Float(0),
				},
			}},
		},
		TextOverlayTransformation: shared.TextOverlayTransformationParam{
			Alpha:      imagekit.Float(1),
			Background: imagekit.String("background"),
			Flip:       shared.TextOverlayTransformationFlipH,
			FontColor:  imagekit.String("fontColor"),
			FontFamily: imagekit.String("fontFamily"),
			FontSize: shared.TextOverlayTransformationFontSizeUnionParam{
				OfFloat: imagekit.Float(0),
			},
			InnerAlignment: shared.TextOverlayTransformationInnerAlignmentLeft,
			LineHeight: shared.TextOverlayTransformationLineHeightUnionParam{
				OfFloat: imagekit.Float(0),
			},
			Padding: shared.TextOverlayTransformationPaddingUnionParam{
				OfFloat: imagekit.Float(0),
			},
			Radius: shared.TextOverlayTransformationRadiusUnionParam{
				OfFloat: imagekit.Float(0),
			},
			Rotation: shared.TextOverlayTransformationRotationUnionParam{
				OfFloat: imagekit.Float(0),
			},
			Typography: imagekit.String("typography"),
			Width: shared.TextOverlayTransformationWidthUnionParam{
				OfFloat: imagekit.Float(0),
			},
		},
		Transformation: shared.TransformationParam{
			AIChangeBackground: imagekit.String("aiChangeBackground"),
			AIDropShadow: shared.TransformationAIDropShadowUnionParam{
				OfTransformationAIDropShadowBoolean: imagekit.Bool(true),
			},
			AIEdit:                     imagekit.String("aiEdit"),
			AIRemoveBackground:         true,
			AIRemoveBackgroundExternal: true,
			AIRetouch:                  true,
			AIUpscale:                  true,
			AIVariation:                true,
			AspectRatio: shared.TransformationAspectRatioUnionParam{
				OfString: imagekit.String("4:3"),
			},
			AudioCodec:      shared.TransformationAudioCodecAac,
			Background:      imagekit.String("red"),
			Blur:            imagekit.Float(10),
			Border:          imagekit.String("5_FF0000"),
			ColorProfile:    imagekit.Bool(true),
			ColorReplace:    imagekit.String("colorReplace"),
			ContrastStretch: true,
			Crop:            shared.TransformationCropForce,
			CropMode:        shared.TransformationCropModePadResize,
			DefaultImage:    imagekit.String("defaultImage"),
			Distort:         imagekit.String("distort"),
			Dpr:             imagekit.Float(2),
			Duration: shared.TransformationDurationUnionParam{
				OfFloat: imagekit.Float(0),
			},
			EndOffset: shared.TransformationEndOffsetUnionParam{
				OfFloat: imagekit.Float(0),
			},
			Flip:   shared.TransformationFlipH,
			Focus:  imagekit.String("center"),
			Format: shared.TransformationFormatAuto,
			Gradient: shared.TransformationGradientUnionParam{
				OfTransformationGradientBoolean: imagekit.Bool(true),
			},
			Grayscale: true,
			Height: shared.TransformationHeightUnionParam{
				OfFloat: imagekit.Float(200),
			},
			Lossless: imagekit.Bool(true),
			Metadata: imagekit.Bool(true),
			Named:    imagekit.String("named"),
			Opacity:  imagekit.Float(0),
			Original: imagekit.Bool(true),
			Overlay: shared.OverlayUnionParam{
				OfText: &shared.TextOverlayParam{
					BaseOverlayParam: shared.BaseOverlayParam{
						LayerMode: shared.BaseOverlayLayerModeMultiply,
						Position: shared.OverlayPositionParam{
							Focus: shared.OverlayPositionFocusCenter,
							X: shared.OverlayPositionXUnionParam{
								OfFloat: imagekit.Float(0),
							},
							Y: shared.OverlayPositionYUnionParam{
								OfFloat: imagekit.Float(0),
							},
						},
						Timing: shared.OverlayTimingParam{
							Duration: shared.OverlayTimingDurationUnionParam{
								OfFloat: imagekit.Float(0),
							},
							End: shared.OverlayTimingEndUnionParam{
								OfFloat: imagekit.Float(0),
							},
							Start: shared.OverlayTimingStartUnionParam{
								OfFloat: imagekit.Float(0),
							},
						},
					},
					Text:     "text",
					Encoding: "auto",
					Transformation: []shared.TextOverlayTransformationParam{{
						Alpha:      imagekit.Float(1),
						Background: imagekit.String("background"),
						Flip:       shared.TextOverlayTransformationFlipH,
						FontColor:  imagekit.String("fontColor"),
						FontFamily: imagekit.String("fontFamily"),
						FontSize: shared.TextOverlayTransformationFontSizeUnionParam{
							OfFloat: imagekit.Float(0),
						},
						InnerAlignment: shared.TextOverlayTransformationInnerAlignmentLeft,
						LineHeight: shared.TextOverlayTransformationLineHeightUnionParam{
							OfFloat: imagekit.Float(0),
						},
						Padding: shared.TextOverlayTransformationPaddingUnionParam{
							OfFloat: imagekit.Float(0),
						},
						Radius: shared.TextOverlayTransformationRadiusUnionParam{
							OfFloat: imagekit.Float(0),
						},
						Rotation: shared.TextOverlayTransformationRotationUnionParam{
							OfFloat: imagekit.Float(0),
						},
						Typography: imagekit.String("typography"),
						Width: shared.TextOverlayTransformationWidthUnionParam{
							OfFloat: imagekit.Float(0),
						},
					}},
				},
			},
			Page: shared.TransformationPageUnionParam{
				OfFloat: imagekit.Float(0),
			},
			Progressive: imagekit.Bool(true),
			Quality:     imagekit.Float(80),
			Radius: shared.TransformationRadiusUnionParam{
				OfFloat: imagekit.Float(20),
			},
			Raw: imagekit.String("raw"),
			Rotation: shared.TransformationRotationUnionParam{
				OfFloat: imagekit.Float(90),
			},
			Shadow: shared.TransformationShadowUnionParam{
				OfTransformationShadowBoolean: imagekit.Bool(true),
			},
			Sharpen: shared.TransformationSharpenUnionParam{
				OfTransformationSharpenBoolean: imagekit.Bool(true),
			},
			StartOffset: shared.TransformationStartOffsetUnionParam{
				OfFloat: imagekit.Float(0),
			},
			StreamingResolutions: []shared.StreamingResolution{shared.StreamingResolution240},
			Trim: shared.TransformationTrimUnionParam{
				OfTransformationTrimBoolean: imagekit.Bool(true),
			},
			UnsharpMask: shared.TransformationUnsharpMaskUnionParam{
				OfTransformationUnsharpMaskBoolean: imagekit.Bool(true),
			},
			VideoCodec: shared.TransformationVideoCodecH264,
			Width: shared.TransformationWidthUnionParam{
				OfFloat: imagekit.Float(300),
			},
			X: shared.TransformationXUnionParam{
				OfFloat: imagekit.Float(0),
			},
			XCenter: shared.TransformationXCenterUnionParam{
				OfFloat: imagekit.Float(0),
			},
			Y: shared.TransformationYUnionParam{
				OfFloat: imagekit.Float(0),
			},
			YCenter: shared.TransformationYCenterUnionParam{
				OfFloat: imagekit.Float(0),
			},
			Zoom: imagekit.Float(0),
		},
		TransformationPosition: shared.TransformationPositionPath,
		VideoOverlay: shared.VideoOverlayParam{
			BaseOverlayParam: shared.BaseOverlayParam{
				LayerMode: shared.BaseOverlayLayerModeMultiply,
				Position: shared.OverlayPositionParam{
					Focus: shared.OverlayPositionFocusCenter,
					X: shared.OverlayPositionXUnionParam{
						OfFloat: imagekit.Float(0),
					},
					Y: shared.OverlayPositionYUnionParam{
						OfFloat: imagekit.Float(0),
					},
				},
				Timing: shared.OverlayTimingParam{
					Duration: shared.OverlayTimingDurationUnionParam{
						OfFloat: imagekit.Float(0),
					},
					End: shared.OverlayTimingEndUnionParam{
						OfFloat: imagekit.Float(0),
					},
					Start: shared.OverlayTimingStartUnionParam{
						OfFloat: imagekit.Float(0),
					},
				},
			},
			Input:    "input",
			Encoding: "auto",
			Transformation: []shared.TransformationParam{{
				AIChangeBackground: imagekit.String("aiChangeBackground"),
				AIDropShadow: shared.TransformationAIDropShadowUnionParam{
					OfTransformationAIDropShadowBoolean: imagekit.Bool(true),
				},
				AIEdit:                     imagekit.String("aiEdit"),
				AIRemoveBackground:         true,
				AIRemoveBackgroundExternal: true,
				AIRetouch:                  true,
				AIUpscale:                  true,
				AIVariation:                true,
				AspectRatio: shared.TransformationAspectRatioUnionParam{
					OfString: imagekit.String("4:3"),
				},
				AudioCodec:      shared.TransformationAudioCodecAac,
				Background:      imagekit.String("red"),
				Blur:            imagekit.Float(10),
				Border:          imagekit.String("5_FF0000"),
				ColorProfile:    imagekit.Bool(true),
				ColorReplace:    imagekit.String("colorReplace"),
				ContrastStretch: true,
				Crop:            shared.TransformationCropForce,
				CropMode:        shared.TransformationCropModePadResize,
				DefaultImage:    imagekit.String("defaultImage"),
				Distort:         imagekit.String("distort"),
				Dpr:             imagekit.Float(2),
				Duration: shared.TransformationDurationUnionParam{
					OfFloat: imagekit.Float(0),
				},
				EndOffset: shared.TransformationEndOffsetUnionParam{
					OfFloat: imagekit.Float(0),
				},
				Flip:   shared.TransformationFlipH,
				Focus:  imagekit.String("center"),
				Format: shared.TransformationFormatAuto,
				Gradient: shared.TransformationGradientUnionParam{
					OfTransformationGradientBoolean: imagekit.Bool(true),
				},
				Grayscale: true,
				Height: shared.TransformationHeightUnionParam{
					OfFloat: imagekit.Float(200),
				},
				Lossless: imagekit.Bool(true),
				Metadata: imagekit.Bool(true),
				Named:    imagekit.String("named"),
				Opacity:  imagekit.Float(0),
				Original: imagekit.Bool(true),
				Overlay: shared.OverlayUnionParam{
					OfText: &shared.TextOverlayParam{
						BaseOverlayParam: shared.BaseOverlayParam{
							LayerMode: shared.BaseOverlayLayerModeMultiply,
							Position: shared.OverlayPositionParam{
								Focus: shared.OverlayPositionFocusCenter,
								X: shared.OverlayPositionXUnionParam{
									OfFloat: imagekit.Float(0),
								},
								Y: shared.OverlayPositionYUnionParam{
									OfFloat: imagekit.Float(0),
								},
							},
							Timing: shared.OverlayTimingParam{
								Duration: shared.OverlayTimingDurationUnionParam{
									OfFloat: imagekit.Float(0),
								},
								End: shared.OverlayTimingEndUnionParam{
									OfFloat: imagekit.Float(0),
								},
								Start: shared.OverlayTimingStartUnionParam{
									OfFloat: imagekit.Float(0),
								},
							},
						},
						Text:     "text",
						Encoding: "auto",
						Transformation: []shared.TextOverlayTransformationParam{{
							Alpha:      imagekit.Float(1),
							Background: imagekit.String("background"),
							Flip:       shared.TextOverlayTransformationFlipH,
							FontColor:  imagekit.String("fontColor"),
							FontFamily: imagekit.String("fontFamily"),
							FontSize: shared.TextOverlayTransformationFontSizeUnionParam{
								OfFloat: imagekit.Float(0),
							},
							InnerAlignment: shared.TextOverlayTransformationInnerAlignmentLeft,
							LineHeight: shared.TextOverlayTransformationLineHeightUnionParam{
								OfFloat: imagekit.Float(0),
							},
							Padding: shared.TextOverlayTransformationPaddingUnionParam{
								OfFloat: imagekit.Float(0),
							},
							Radius: shared.TextOverlayTransformationRadiusUnionParam{
								OfFloat: imagekit.Float(0),
							},
							Rotation: shared.TextOverlayTransformationRotationUnionParam{
								OfFloat: imagekit.Float(0),
							},
							Typography: imagekit.String("typography"),
							Width: shared.TextOverlayTransformationWidthUnionParam{
								OfFloat: imagekit.Float(0),
							},
						}},
					},
				},
				Page: shared.TransformationPageUnionParam{
					OfFloat: imagekit.Float(0),
				},
				Progressive: imagekit.Bool(true),
				Quality:     imagekit.Float(80),
				Radius: shared.TransformationRadiusUnionParam{
					OfFloat: imagekit.Float(20),
				},
				Raw: imagekit.String("raw"),
				Rotation: shared.TransformationRotationUnionParam{
					OfFloat: imagekit.Float(90),
				},
				Shadow: shared.TransformationShadowUnionParam{
					OfTransformationShadowBoolean: imagekit.Bool(true),
				},
				Sharpen: shared.TransformationSharpenUnionParam{
					OfTransformationSharpenBoolean: imagekit.Bool(true),
				},
				StartOffset: shared.TransformationStartOffsetUnionParam{
					OfFloat: imagekit.Float(0),
				},
				StreamingResolutions: []shared.StreamingResolution{shared.StreamingResolution240},
				Trim: shared.TransformationTrimUnionParam{
					OfTransformationTrimBoolean: imagekit.Bool(true),
				},
				UnsharpMask: shared.TransformationUnsharpMaskUnionParam{
					OfTransformationUnsharpMaskBoolean: imagekit.Bool(true),
				},
				VideoCodec: shared.TransformationVideoCodecH264,
				Width: shared.TransformationWidthUnionParam{
					OfFloat: imagekit.Float(300),
				},
				X: shared.TransformationXUnionParam{
					OfFloat: imagekit.Float(0),
				},
				XCenter: shared.TransformationXCenterUnionParam{
					OfFloat: imagekit.Float(0),
				},
				Y: shared.TransformationYUnionParam{
					OfFloat: imagekit.Float(0),
				},
				YCenter: shared.TransformationYCenterUnionParam{
					OfFloat: imagekit.Float(0),
				},
				Zoom: imagekit.Float(0),
			}},
		},
	})
	if err != nil {
		var apierr *imagekit.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
