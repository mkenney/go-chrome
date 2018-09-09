package socket

import (
	"testing"

	"github.com/mkenney/go-chrome/tot/css"
	"github.com/mkenney/go-chrome/tot/dom"
	"github.com/mkenney/go-chrome/tot/page"
)

func TestCSSAddRule(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &css.AddRuleResult{Rule: &css.Rule{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
		SelectorList: &css.SelectorList{
			Selectors: []*css.Value{{
				Text: "rule text",
				Range: &css.SourceRange{
					StartLine:   10,
					StartColumn: 10,
					EndLine:     10,
					EndColumn:   10,
				},
			}},
			Text: "rule text",
		},
		Origin: css.StyleSheetOrigin.Injected,
		Style: &css.Style{
			StyleSheetID: css.StyleSheetID("stylesheet-id"),
			Properties: []*css.Property{{
				Name:      "Name",
				Value:     "Value",
				Important: true,
				Implicit:  true,
				Text:      "Text",
				ParsedOk:  true,
				Disabled:  true,
				Range: &css.SourceRange{
					StartLine:   10,
					StartColumn: 10,
					EndLine:     10,
					EndColumn:   10,
				},
			}},
			ShorthandEntries: []*css.ShorthandEntry{{
				Name:      "Name",
				Value:     "Value",
				Important: true,
			}},
			Text: "Text",
			Range: &css.SourceRange{
				StartLine:   1,
				StartColumn: 1,
				EndLine:     1,
				EndColumn:   1,
			},
		},
		Media: []*css.Media{{
			Text:      "media text",
			Source:    css.Source.MediaRule,
			SourceURL: "http://source-url",
			Range: &css.SourceRange{
				StartLine:   10,
				StartColumn: 10,
				EndLine:     10,
				EndColumn:   10,
			},
			StyleSheetID: css.StyleSheetID("stylesheet-id"),
			MediaList: []*css.MediaQuery{{
				Expressions: []*css.MediaQueryExpression{{
					Value:   1,
					Unit:    "px",
					Feature: "feature",
					ValueRange: &css.SourceRange{
						StartLine:   10,
						StartColumn: 10,
						EndLine:     10,
						EndColumn:   10,
					},
					ComputedLength: 1,
				}},
			}},
		}},
	}}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.CSS().AddRule(&css.AddRuleParams{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
		RuleText:     "rule text",
		Location: &css.SourceRange{
			StartLine:   10,
			StartColumn: 10,
			EndLine:     10,
			EndColumn:   10,
		},
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Rule.StyleSheetID != result.Rule.StyleSheetID {
		t.Errorf("Expected '%s', got '%s'", mockResult.Rule.StyleSheetID, result.Rule.StyleSheetID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.CSS().AddRule(&css.AddRuleParams{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
		RuleText:     "rule text",
		Location: &css.SourceRange{
			StartLine:   10,
			StartColumn: 10,
			EndLine:     10,
			EndColumn:   10,
		},
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSCollectClassNames(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &css.CollectClassNamesResult{ClassNames: []string{
		"class1", "class2",
	}}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.CSS().CollectClassNames(&css.CollectClassNamesParams{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.ClassNames[0] != result.ClassNames[0] {
		t.Errorf("Expected '%s', got '%s'", mockResult.ClassNames[0], result.ClassNames[0])
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.CSS().CollectClassNames(&css.CollectClassNamesParams{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSCreateStyleSheet(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &css.CreateStyleSheetResult{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.CSS().CreateStyleSheet(&css.CreateStyleSheetParams{
		FrameID: page.FrameID("frame-id"),
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.StyleSheetID != result.StyleSheetID {
		t.Errorf("Expected '%s', got '%s'", mockResult.StyleSheetID, result.StyleSheetID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.CSS().CreateStyleSheet(&css.CreateStyleSheetParams{
		FrameID: page.FrameID("frame-id"),
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSDisable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &css.DisableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.CSS().Disable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.CSS().Disable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSEnable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &css.EnableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.CSS().Enable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.CSS().Enable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSForcePseudoState(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &css.ForcePseudoStateResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.CSS().ForcePseudoState(&css.ForcePseudoStateParams{
		NodeID:              dom.NodeID(1),
		ForcedPseudoClasses: []css.ForcedPseudoClassesEnum{css.ForcedPseudoClasses.Active},
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.CSS().ForcePseudoState(&css.ForcePseudoStateParams{
		NodeID:              dom.NodeID(1),
		ForcedPseudoClasses: []css.ForcedPseudoClassesEnum{css.ForcedPseudoClasses.Active},
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSGetBackgroundColors(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &css.GetBackgroundColorsResult{
		BackgroundColors:     []string{"blue"},
		ComputedFontSize:     "10pt",
		ComputedFontWeight:   "normal",
		ComputedBodyFontSize: "10pt",
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.CSS().GetBackgroundColors(&css.GetBackgroundColorsParams{
		NodeID: dom.NodeID(1),
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.BackgroundColors[0] != result.BackgroundColors[0] {
		t.Errorf("Expected '%s', got '%s'", mockResult.BackgroundColors[0], result.BackgroundColors[0])
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.CSS().GetBackgroundColors(&css.GetBackgroundColorsParams{
		NodeID: dom.NodeID(1),
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSGetComputedStyleForNode(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &css.GetComputedStyleForNodeResult{
		ComputedStyle: []*css.ComputedStyleProperty{{
			Name:  "style-name",
			Value: "style-value",
		}},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.CSS().GetComputedStyleForNode(&css.GetComputedStyleForNodeParams{
		NodeID: dom.NodeID(1),
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.ComputedStyle[0].Name != result.ComputedStyle[0].Name {
		t.Errorf("Expected '%s', got '%s'", mockResult.ComputedStyle[0].Name, result.ComputedStyle[0].Name)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.CSS().GetComputedStyleForNode(&css.GetComputedStyleForNodeParams{
		NodeID: dom.NodeID(1),
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSGetInlineStylesForNode(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &css.GetInlineStylesForNodeResult{
		InlineStyle: &css.Style{
			StyleSheetID: css.StyleSheetID("stylesheet-id"),
			Properties: []*css.Property{{
				Name:      "property-name",
				Value:     "property-value",
				Important: true,
				Implicit:  true,
				Text:      "style text",
				ParsedOk:  true,
				Disabled:  true,
				Range: &css.SourceRange{
					StartLine:   1,
					StartColumn: 1,
					EndLine:     2,
					EndColumn:   2,
				},
			}},
		},
		AttributesStyle: &css.Style{
			StyleSheetID: css.StyleSheetID("stylesheet-id"),
			Properties: []*css.Property{{
				Name:      "property-name",
				Value:     "property-value",
				Important: true,
				Implicit:  true,
				Text:      "style text",
				ParsedOk:  true,
				Disabled:  true,
				Range: &css.SourceRange{
					StartLine:   1,
					StartColumn: 1,
					EndLine:     2,
					EndColumn:   2,
				},
			}},
		},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.CSS().GetInlineStylesForNode(&css.GetInlineStylesForNodeParams{
		NodeID: dom.NodeID(1),
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.InlineStyle.StyleSheetID != result.InlineStyle.StyleSheetID {
		t.Errorf("Expected '%s', got '%s'", mockResult.InlineStyle.StyleSheetID, result.InlineStyle.StyleSheetID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.CSS().GetInlineStylesForNode(&css.GetInlineStylesForNodeParams{
		NodeID: dom.NodeID(1),
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSGetMatchedStylesForNode(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &css.GetMatchedStylesForNodeResult{
		InlineStyle: &css.Style{
			StyleSheetID: css.StyleSheetID("stylesheet-id"),
			Properties: []*css.Property{{
				Name:      "property-name",
				Value:     "property-value",
				Important: true,
				Implicit:  true,
				Text:      "style text",
				ParsedOk:  true,
				Disabled:  true,
				Range: &css.SourceRange{
					StartLine:   1,
					StartColumn: 1,
					EndLine:     2,
					EndColumn:   2,
				},
			}},
		},
		AttributesStyle: &css.Style{
			StyleSheetID: css.StyleSheetID("stylesheet-id"),
			Properties: []*css.Property{{
				Name:      "property-name",
				Value:     "property-value",
				Important: true,
				Implicit:  true,
				Text:      "style text",
				ParsedOk:  true,
				Disabled:  true,
				Range: &css.SourceRange{
					StartLine:   1,
					StartColumn: 1,
					EndLine:     2,
					EndColumn:   2,
				},
			}},
		},
		MatchedRules: []*css.RuleMatch{{
			Rule: &css.Rule{
				StyleSheetID: css.StyleSheetID("stylesheet-id"),
				SelectorList: &css.SelectorList{
					Selectors: []*css.Value{{
						Text: "some text",
						Range: &css.SourceRange{
							StartLine:   10,
							StartColumn: 10,
							EndLine:     10,
							EndColumn:   10,
						},
					}},
				},
				Origin: css.StyleSheetOrigin.Injected,
				Style:  &css.Style{},
			},
		}},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.CSS().GetMatchedStylesForNode(&css.GetMatchedStylesForNodeParams{
		NodeID: dom.NodeID(1),
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.InlineStyle.StyleSheetID != result.InlineStyle.StyleSheetID {
		t.Errorf("Expected '%s', got '%s'", mockResult.InlineStyle.StyleSheetID, result.InlineStyle.StyleSheetID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.CSS().GetMatchedStylesForNode(&css.GetMatchedStylesForNodeParams{
		NodeID: dom.NodeID(1),
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSGetMediaQueries(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &css.GetMediaQueriesResult{
		Medias: []*css.Media{{
			Text:      "media text",
			Source:    css.Source.MediaRule,
			SourceURL: "http://source-url",
			Range: &css.SourceRange{
				StartLine:   10,
				StartColumn: 10,
				EndLine:     10,
				EndColumn:   10,
			},
			StyleSheetID: css.StyleSheetID("stylesheet-id"),
			MediaList: []*css.MediaQuery{{
				Expressions: []*css.MediaQueryExpression{{
					Value:   1,
					Unit:    "px",
					Feature: "feature",
					ValueRange: &css.SourceRange{
						StartLine:   10,
						StartColumn: 10,
						EndLine:     10,
						EndColumn:   10,
					},
					ComputedLength: 1,
				}},
			}},
		}},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.CSS().GetMediaQueries()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Medias[0].Text != result.Medias[0].Text {
		t.Errorf("Expected '%s', got '%s'", mockResult.Medias[0].Text, result.Medias[0].Text)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.CSS().GetMediaQueries()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSGetPlatformFontsForNode(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &css.GetPlatformFontsForNodeResult{
		Fonts: []*css.PlatformFontUsage{{
			FamilyName:   "courier",
			IsCustomFont: true,
			GlyphCount:   1,
		}},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.CSS().GetPlatformFontsForNode(&css.GetPlatformFontsForNodeParams{
		NodeID: dom.NodeID(1),
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Fonts[0].FamilyName != result.Fonts[0].FamilyName {
		t.Errorf("Expected '%s', got '%s'", mockResult.Fonts[0].FamilyName, result.Fonts[0].FamilyName)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.CSS().GetPlatformFontsForNode(&css.GetPlatformFontsForNodeParams{
		NodeID: dom.NodeID(1),
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSGetStyleSheetText(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &css.GetStyleSheetTextResult{
		Text: "some text",
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.CSS().GetStyleSheetText(&css.GetStyleSheetTextParams{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Text != result.Text {
		t.Errorf("Expected '%s', got '%s'", mockResult.Text, result.Text)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.CSS().GetStyleSheetText(&css.GetStyleSheetTextParams{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSSetEffectivePropertyValueForNode(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &css.SetEffectivePropertyValueForNodeResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.CSS().SetEffectivePropertyValueForNode(&css.SetEffectivePropertyValueForNodeParams{
		NodeID:       dom.NodeID(1),
		PropertyName: "property-name",
		Value:        "property-value",
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.CSS().SetEffectivePropertyValueForNode(&css.SetEffectivePropertyValueForNodeParams{
		NodeID:       dom.NodeID(1),
		PropertyName: "property-name",
		Value:        "property-value",
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSSetKeyframeKey(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &css.SetKeyframeKeyResult{
		KeyText: &css.Value{
			Text: "some text",
			Range: &css.SourceRange{
				StartLine:   10,
				StartColumn: 10,
				EndLine:     10,
				EndColumn:   10,
			},
		},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.CSS().SetKeyframeKey(&css.SetKeyframeKeyParams{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
		Range: &css.SourceRange{
			StartLine:   10,
			StartColumn: 10,
			EndLine:     10,
			EndColumn:   10,
		},
		Selector: "selector",
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.KeyText.Text != result.KeyText.Text {
		t.Errorf("Expected '%s', got '%s'", mockResult.KeyText.Text, result.KeyText.Text)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.CSS().SetKeyframeKey(&css.SetKeyframeKeyParams{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
		Range: &css.SourceRange{
			StartLine:   10,
			StartColumn: 10,
			EndLine:     10,
			EndColumn:   10,
		},
		Selector: "selector",
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSSetMediaText(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &css.SetMediaTextResult{
		Media: &css.Media{
			Text:      "some text",
			Source:    css.Source.MediaRule,
			SourceURL: "http://source-url",
			Range: &css.SourceRange{
				StartLine:   10,
				StartColumn: 10,
				EndLine:     10,
				EndColumn:   10,
			},
			StyleSheetID: css.StyleSheetID("stylesheet-id"),
			MediaList: []*css.MediaQuery{{
				Expressions: []*css.MediaQueryExpression{{
					Value:   1,
					Unit:    "px",
					Feature: "feature",
					ValueRange: &css.SourceRange{
						StartLine:   10,
						StartColumn: 10,
						EndLine:     10,
						EndColumn:   10,
					},
					ComputedLength: 1,
				}},
			}},
		},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.CSS().SetMediaText(&css.SetMediaTextParams{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
		Range: &css.SourceRange{
			StartLine:   10,
			StartColumn: 10,
			EndLine:     10,
			EndColumn:   10,
		},
		Text: "some text",
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Media.Text != result.Media.Text {
		t.Errorf("Expected '%s', got '%s'", mockResult.Media.Text, result.Media.Text)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.CSS().SetMediaText(&css.SetMediaTextParams{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
		Range: &css.SourceRange{
			StartLine:   10,
			StartColumn: 10,
			EndLine:     10,
			EndColumn:   10,
		},
		Text: "some text",
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSSetRuleSelector(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &css.SetRuleSelectorResult{
		SelectorList: &css.SelectorList{
			Selectors: []*css.Value{{
				Text: "rule text",
				Range: &css.SourceRange{
					StartLine:   10,
					StartColumn: 10,
					EndLine:     10,
					EndColumn:   10,
				},
			}},
		},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.CSS().SetRuleSelector(&css.SetRuleSelectorParams{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
		Range: &css.SourceRange{
			StartLine:   10,
			StartColumn: 10,
			EndLine:     10,
			EndColumn:   10,
		},
		Selector: "selector",
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.SelectorList.Selectors[0].Text != result.SelectorList.Selectors[0].Text {
		t.Errorf("Expected '%s', got '%s'", mockResult.SelectorList.Selectors[0].Text, result.SelectorList.Selectors[0].Text)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.CSS().SetRuleSelector(&css.SetRuleSelectorParams{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
		Range: &css.SourceRange{
			StartLine:   10,
			StartColumn: 10,
			EndLine:     10,
			EndColumn:   10,
		},
		Selector: "selector",
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSSetStyleSheetText(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &css.SetStyleSheetTextResult{
		SourceMapURL: "http://sourcemap-url",
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.CSS().SetStyleSheetText(&css.SetStyleSheetTextParams{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
		Text:         "some text",
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.SourceMapURL != result.SourceMapURL {
		t.Errorf("Expected '%s', got '%s'", mockResult.SourceMapURL, result.SourceMapURL)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.CSS().SetStyleSheetText(&css.SetStyleSheetTextParams{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
		Text:         "some text",
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSSetStyleTexts(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &css.SetStyleTextsResult{
		Styles: []*css.Style{{
			StyleSheetID: css.StyleSheetID("stylesheet-id"),
			Properties: []*css.Property{{
				Name:      "property-name",
				Value:     "property-value",
				Important: true,
				Implicit:  true,
				Text:      "style text",
				ParsedOk:  true,
				Disabled:  true,
				Range: &css.SourceRange{
					StartLine:   1,
					StartColumn: 1,
					EndLine:     2,
					EndColumn:   2,
				},
			}},
		}},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.CSS().SetStyleTexts(&css.SetStyleTextsParams{
		Edits: []*css.StyleDeclarationEdit{{
			StyleSheetID: css.StyleSheetID("stylesheet-id"),
			Range: &css.SourceRange{
				StartLine:   10,
				StartColumn: 10,
				EndLine:     10,
				EndColumn:   10,
			},
			Text: "some text",
		}},
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Styles[0].StyleSheetID != result.Styles[0].StyleSheetID {
		t.Errorf("Expected '%s', got '%s'", mockResult.Styles[0].StyleSheetID, result.Styles[0].StyleSheetID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.CSS().SetStyleTexts(&css.SetStyleTextsParams{
		Edits: []*css.StyleDeclarationEdit{{
			StyleSheetID: css.StyleSheetID("stylesheet-id"),
			Range: &css.SourceRange{
				StartLine:   10,
				StartColumn: 10,
				EndLine:     10,
				EndColumn:   10,
			},
			Text: "some text",
		}},
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSStartRuleUsageTracking(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &css.StartRuleUsageTrackingResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.CSS().StartRuleUsageTracking()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.CSS().StartRuleUsageTracking()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSStopRuleUsageTracking(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &css.StopRuleUsageTrackingResult{
		RuleUsage: []*css.RuleUsage{{
			StyleSheetID: css.StyleSheetID("stylesheet-id"),
			StartOffset:  1,
			EndOffset:    2,
			Used:         true,
		}},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.CSS().StopRuleUsageTracking()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.RuleUsage[0].StyleSheetID != result.RuleUsage[0].StyleSheetID {
		t.Errorf("Expected '%s', got '%s'", mockResult.RuleUsage[0].StyleSheetID, result.RuleUsage[0].StyleSheetID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.CSS().StopRuleUsageTracking()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSTakeCoverageDelta(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &css.TakeCoverageDeltaResult{
		Coverage: []*css.RuleUsage{{
			StyleSheetID: css.StyleSheetID("stylesheet-id"),
			StartOffset:  1,
			EndOffset:    2,
			Used:         true,
		}},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.CSS().TakeCoverageDelta()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Coverage[0].StyleSheetID != result.Coverage[0].StyleSheetID {
		t.Errorf("Expected '%s', got '%s'", mockResult.Coverage[0].StyleSheetID, result.Coverage[0].StyleSheetID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.CSS().TakeCoverageDelta()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSOnFontsUpdated(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *css.FontsUpdatedEvent)
	soc.CSS().OnFontsUpdated(func(eventData *css.FontsUpdatedEvent) {
		resultChan <- eventData
	})

	mockResult := &css.FontsUpdatedEvent{}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "CSS.fontsUpdated",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "CSS.fontsUpdated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSOnMediaQueryResultChanged(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *css.MediaQueryResultChangedEvent)
	soc.CSS().OnMediaQueryResultChanged(func(eventData *css.MediaQueryResultChangedEvent) {
		resultChan <- eventData
	})

	mockResult := &css.MediaQueryResultChangedEvent{}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "CSS.mediaQueryResultChanged",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "CSS.mediaQueryResultChanged",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSOnStyleSheetAdded(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *css.StyleSheetAddedEvent)
	soc.CSS().OnStyleSheetAdded(func(eventData *css.StyleSheetAddedEvent) {
		resultChan <- eventData
	})

	mockResult := &css.StyleSheetAddedEvent{
		Header: &css.StyleSheetHeader{
			StyleSheetID: css.StyleSheetID("stylesheet-id"),
			FrameID:      page.FrameID("frame-id"),
			SourceURL:    "http://source-url",
			Origin:       css.StyleSheetOrigin.Injected,
			Title:        "title",
			OwnerNode:    dom.BackendNodeID(1),
			Disabled:     true,
			HasSourceURL: true,
			IsInline:     true,
			StartLine:    1,
			StartColumn:  1,
			Length:       1,
		},
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "CSS.styleSheetAdded",
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected success, got error: '%s'", result.Err.Error())
	}
	if mockResult.Header.StyleSheetID != result.Header.StyleSheetID {
		t.Errorf("Expected '%s', got '%s'", mockResult.Header.StyleSheetID, result.Header.StyleSheetID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "CSS.styleSheetAdded",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSOnStyleSheetChanged(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *css.StyleSheetChangedEvent)
	soc.CSS().OnStyleSheetChanged(func(eventData *css.StyleSheetChangedEvent) {
		resultChan <- eventData
	})

	mockResult := &css.StyleSheetChangedEvent{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "CSS.styleSheetChanged",
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected success, got error: '%s'", result.Err.Error())
	}
	if mockResult.StyleSheetID != result.StyleSheetID {
		t.Errorf("Expected '%s', got '%s'", mockResult.StyleSheetID, result.StyleSheetID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "CSS.styleSheetChanged",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSOnStyleSheetRemoved(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *css.StyleSheetRemovedEvent)
	soc.CSS().OnStyleSheetRemoved(func(eventData *css.StyleSheetRemovedEvent) {
		resultChan <- eventData
	})

	mockResult := &css.StyleSheetRemovedEvent{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "CSS.styleSheetRemoved",
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected success, got error: '%s'", result.Err.Error())
	}
	if mockResult.StyleSheetID != result.StyleSheetID {
		t.Errorf("Expected '%s', got '%s'", mockResult.StyleSheetID, result.StyleSheetID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "CSS.styleSheetRemoved",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
