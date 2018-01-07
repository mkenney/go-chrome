package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	css "github.com/mkenney/go-chrome/cdtp/css"
	dom "github.com/mkenney/go-chrome/cdtp/dom"
	page "github.com/mkenney/go-chrome/cdtp/page"
)

func TestCSSAddRule(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CSS().AddRule(&css.AddRuleParams{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
		RuleText:     "rule text",
		Location: &css.SourceRange{
			StartLine:   10,
			StartColumn: 10,
			EndLine:     10,
			EndColumn:   10,
		},
	})
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
	}}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "CSS.addRule",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Rule.StyleSheetID != result.Rule.StyleSheetID {
		t.Errorf("Expected '%s', got '%s'", mockResult.Rule.StyleSheetID, result.Rule.StyleSheetID)
	}

	resultChan = mockSocket.CSS().AddRule(&css.AddRuleParams{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
		RuleText:     "rule text",
		Location: &css.SourceRange{
			StartLine:   10,
			StartColumn: 10,
			EndLine:     10,
			EndColumn:   10,
		},
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "CSS.addRule",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSCollectClassNames(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CSS().CollectClassNames(&css.CollectClassNamesParams{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
	})
	mockResult := &css.CollectClassNamesResult{ClassNames: []string{
		"class1", "class2",
	}}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "CSS.collectClassNames",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.ClassNames[0] != result.ClassNames[0] {
		t.Errorf("Expected '%s', got '%s'", mockResult.ClassNames[0], result.ClassNames[0])
	}

	resultChan = mockSocket.CSS().CollectClassNames(&css.CollectClassNamesParams{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "CSS.collectClassNames",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSCreateStyleSheet(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CSS().CreateStyleSheet(&css.CreateStyleSheetParams{
		FrameID: page.FrameID("frame-id"),
	})
	mockResult := &css.CreateStyleSheetResult{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "CSS.createStyleSheet",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.StyleSheetID != result.StyleSheetID {
		t.Errorf("Expected '%s', got '%s'", mockResult.StyleSheetID, result.StyleSheetID)
	}

	resultChan = mockSocket.CSS().CreateStyleSheet(&css.CreateStyleSheetParams{
		FrameID: page.FrameID("frame-id"),
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "CSS.createStyleSheet",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSDisable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CSS().Disable()
	mockResult := &css.DisableResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "CSS.disable",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.CSS().Disable()
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "CSS.disable",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSEnable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CSS().Enable()
	mockResult := &css.EnableResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "CSS.enable",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.CSS().Enable()
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "CSS.enable",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSForcePseudoState(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CSS().ForcePseudoState(&css.ForcePseudoStateParams{
		NodeID:              dom.NodeID(1),
		ForcedPseudoClasses: []string{"pseudo-class"},
	})
	mockResult := &css.ForcePseudoStateResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "CSS.forcePseudoState",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.CSS().ForcePseudoState(&css.ForcePseudoStateParams{
		NodeID:              dom.NodeID(1),
		ForcedPseudoClasses: []string{"pseudo-class"},
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "CSS.forcePseudoState",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSGetBackgroundColors(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CSS().GetBackgroundColors(&css.GetBackgroundColorsParams{
		NodeID: dom.NodeID(1),
	})
	mockResult := &css.GetBackgroundColorsResult{
		BackgroundColors:     []string{"blue"},
		ComputedFontSize:     "10pt",
		ComputedFontWeight:   "normal",
		ComputedBodyFontSize: "10pt",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "CSS.getBackgroundColors",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.BackgroundColors[0] != result.BackgroundColors[0] {
		t.Errorf("Expected '%s', got '%s'", mockResult.BackgroundColors[0], result.BackgroundColors[0])
	}

	resultChan = mockSocket.CSS().GetBackgroundColors(&css.GetBackgroundColorsParams{
		NodeID: dom.NodeID(1),
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "CSS.getBackgroundColors",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSGetComputedStyleForNode(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CSS().GetComputedStyleForNode(&css.GetComputedStyleForNodeParams{
		NodeID: dom.NodeID(1),
	})
	mockResult := &css.GetComputedStyleForNodeResult{
		ComputedStyle: []*css.ComputedStyleProperty{{
			Name:  "style-name",
			Value: "style-value",
		}},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "CSS.getComputedStyleForNode",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.ComputedStyle[0].Name != result.ComputedStyle[0].Name {
		t.Errorf("Expected '%s', got '%s'", mockResult.ComputedStyle[0].Name, result.ComputedStyle[0].Name)
	}

	resultChan = mockSocket.CSS().GetComputedStyleForNode(&css.GetComputedStyleForNodeParams{
		NodeID: dom.NodeID(1),
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "CSS.getComputedStyleForNode",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSGetInlineStylesForNode(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CSS().GetInlineStylesForNode(&css.GetInlineStylesForNodeParams{
		NodeID: dom.NodeID(1),
	})
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
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "CSS.getInlineStylesForNode",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.InlineStyle.StyleSheetID != result.InlineStyle.StyleSheetID {
		t.Errorf("Expected '%s', got '%s'", mockResult.InlineStyle.StyleSheetID, result.InlineStyle.StyleSheetID)
	}

	resultChan = mockSocket.CSS().GetInlineStylesForNode(&css.GetInlineStylesForNodeParams{
		NodeID: dom.NodeID(1),
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "CSS.getInlineStylesForNode",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSGetMatchedStylesForNode(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CSS().GetMatchedStylesForNode(&css.GetMatchedStylesForNodeParams{
		NodeID: dom.NodeID(1),
	})
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
				Origin: css.StyleSheetOrigin("origin"),
				Style:  &css.Style{},
			},
		}},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "CSS.getMatchedStylesForNode",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.InlineStyle.StyleSheetID != result.InlineStyle.StyleSheetID {
		t.Errorf("Expected '%s', got '%s'", mockResult.InlineStyle.StyleSheetID, result.InlineStyle.StyleSheetID)
	}

	resultChan = mockSocket.CSS().GetMatchedStylesForNode(&css.GetMatchedStylesForNodeParams{
		NodeID: dom.NodeID(1),
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "CSS.getMatchedStylesForNode",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSGetMediaQueries(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CSS().GetMediaQueries()
	mockResult := &css.GetMediaQueriesResult{
		Medias: []*css.Media{{
			Text:      "media text",
			Source:    "media source",
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
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "CSS.getMediaQueries",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Medias[0].Text != result.Medias[0].Text {
		t.Errorf("Expected '%s', got '%s'", mockResult.Medias[0].Text, result.Medias[0].Text)
	}

	resultChan = mockSocket.CSS().GetMediaQueries()
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "CSS.getMediaQueries",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSGetPlatformFontsForNode(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CSS().GetPlatformFontsForNode(&css.GetPlatformFontsForNodeParams{
		NodeID: dom.NodeID(1),
	})
	mockResult := &css.GetPlatformFontsForNodeResult{
		Fonts: []*css.PlatformFontUsage{{
			FamilyName:   "courier",
			IsCustomFont: true,
			GlyphCount:   1,
		}},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "CSS.getPlatformFontsForNode",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Fonts[0].FamilyName != result.Fonts[0].FamilyName {
		t.Errorf("Expected '%s', got '%s'", mockResult.Fonts[0].FamilyName, result.Fonts[0].FamilyName)
	}

	resultChan = mockSocket.CSS().GetPlatformFontsForNode(&css.GetPlatformFontsForNodeParams{
		NodeID: dom.NodeID(1),
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "CSS.getPlatformFontsForNode",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSGetStyleSheetText(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CSS().GetStyleSheetText(&css.GetStyleSheetTextParams{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
	})
	mockResult := &css.GetStyleSheetTextResult{
		Text: "some text",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "CSS.getStyleSheetText",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Text != result.Text {
		t.Errorf("Expected '%s', got '%s'", mockResult.Text, result.Text)
	}

	resultChan = mockSocket.CSS().GetStyleSheetText(&css.GetStyleSheetTextParams{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "CSS.getStyleSheetText",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSSetEffectivePropertyValueForNode(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CSS().SetEffectivePropertyValueForNode(&css.SetEffectivePropertyValueForNodeParams{
		NodeID:       dom.NodeID(1),
		PropertyName: "property-name",
		Value:        "property-value",
	})
	mockResult := &css.SetEffectivePropertyValueForNodeResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "CSS.setEffectivePropertyValueForNode",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.CSS().SetEffectivePropertyValueForNode(&css.SetEffectivePropertyValueForNodeParams{
		NodeID:       dom.NodeID(1),
		PropertyName: "property-name",
		Value:        "property-value",
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "CSS.setEffectivePropertyValueForNode",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSSetKeyframeKey(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CSS().SetKeyframeKey(&css.SetKeyframeKeyParams{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
		Range: &css.SourceRange{
			StartLine:   10,
			StartColumn: 10,
			EndLine:     10,
			EndColumn:   10,
		},
		Selector: "selector",
	})
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
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "CSS.setKeyframeKey",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.KeyText.Text != result.KeyText.Text {
		t.Errorf("Expected '%s', got '%s'", mockResult.KeyText.Text, result.KeyText.Text)
	}

	resultChan = mockSocket.CSS().SetKeyframeKey(&css.SetKeyframeKeyParams{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
		Range: &css.SourceRange{
			StartLine:   10,
			StartColumn: 10,
			EndLine:     10,
			EndColumn:   10,
		},
		Selector: "selector",
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "CSS.setKeyframeKey",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSSetMediaText(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CSS().SetMediaText(&css.SetMediaTextParams{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
		Range: &css.SourceRange{
			StartLine:   10,
			StartColumn: 10,
			EndLine:     10,
			EndColumn:   10,
		},
		Text: "some text",
	})
	mockResult := &css.SetMediaTextResult{
		Media: &css.Media{
			Text:      "some text",
			Source:    "source",
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
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "CSS.setMediaText",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Media.Text != result.Media.Text {
		t.Errorf("Expected '%s', got '%s'", mockResult.Media.Text, result.Media.Text)
	}

	resultChan = mockSocket.CSS().SetMediaText(&css.SetMediaTextParams{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
		Range: &css.SourceRange{
			StartLine:   10,
			StartColumn: 10,
			EndLine:     10,
			EndColumn:   10,
		},
		Text: "some text",
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "CSS.setMediaText",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSSetRuleSelector(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CSS().SetRuleSelector(&css.SetRuleSelectorParams{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
		Range: &css.SourceRange{
			StartLine:   10,
			StartColumn: 10,
			EndLine:     10,
			EndColumn:   10,
		},
		Selector: "selector",
	})
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
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "CSS.setRuleSelector",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.SelectorList.Selectors[0].Text != result.SelectorList.Selectors[0].Text {
		t.Errorf("Expected '%s', got '%s'", mockResult.SelectorList.Selectors[0].Text, result.SelectorList.Selectors[0].Text)
	}

	resultChan = mockSocket.CSS().SetRuleSelector(&css.SetRuleSelectorParams{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
		Range: &css.SourceRange{
			StartLine:   10,
			StartColumn: 10,
			EndLine:     10,
			EndColumn:   10,
		},
		Selector: "selector",
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "CSS.setRuleSelector",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSSetStyleSheetText(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CSS().SetStyleSheetText(&css.SetStyleSheetTextParams{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
		Text:         "some text",
	})
	mockResult := &css.SetStyleSheetTextResult{
		SourceMapURL: "http://sourcemap-url",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "CSS.setStyleSheetText",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.SourceMapURL != result.SourceMapURL {
		t.Errorf("Expected '%s', got '%s'", mockResult.SourceMapURL, result.SourceMapURL)
	}

	resultChan = mockSocket.CSS().SetStyleSheetText(&css.SetStyleSheetTextParams{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
		Text:         "some text",
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "CSS.setStyleSheetText",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSSetStyleTexts(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CSS().SetStyleTexts(&css.SetStyleTextsParams{
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
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "CSS.setStyleTexts",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Styles[0].StyleSheetID != result.Styles[0].StyleSheetID {
		t.Errorf("Expected '%s', got '%s'", mockResult.Styles[0].StyleSheetID, result.Styles[0].StyleSheetID)
	}

	resultChan = mockSocket.CSS().SetStyleTexts(&css.SetStyleTextsParams{
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
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "CSS.setStyleTexts",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSStartRuleUsageTracking(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CSS().StartRuleUsageTracking()
	mockResult := &css.StartRuleUsageTrackingResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "CSS.startRuleUsageTracking",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.CSS().StartRuleUsageTracking()
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "CSS.startRuleUsageTracking",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSStopRuleUsageTracking(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CSS().StopRuleUsageTracking()
	mockResult := &css.StopRuleUsageTrackingResult{
		RuleUsage: []*css.RuleUsage{{
			StyleSheetID: css.StyleSheetID("stylesheet-id"),
			StartOffset:  1,
			EndOffset:    2,
			Used:         true,
		}},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "CSS.stopRuleUsageTracking",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.RuleUsage[0].StyleSheetID != result.RuleUsage[0].StyleSheetID {
		t.Errorf("Expected '%s', got '%s'", mockResult.RuleUsage[0].StyleSheetID, result.RuleUsage[0].StyleSheetID)
	}

	resultChan = mockSocket.CSS().StopRuleUsageTracking()
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "CSS.stopRuleUsageTracking",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSTakeCoverageDelta(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.CSS().TakeCoverageDelta()
	mockResult := &css.TakeCoverageDeltaResult{
		Coverage: []*css.RuleUsage{{
			StyleSheetID: css.StyleSheetID("stylesheet-id"),
			StartOffset:  1,
			EndOffset:    2,
			Used:         true,
		}},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "CSS.takeCoverageDelta",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Coverage[0].StyleSheetID != result.Coverage[0].StyleSheetID {
		t.Errorf("Expected '%s', got '%s'", mockResult.Coverage[0].StyleSheetID, result.Coverage[0].StyleSheetID)
	}

	resultChan = mockSocket.CSS().TakeCoverageDelta()
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "CSS.takeCoverageDelta",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSOnFontsUpdated(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *css.FontsUpdatedEvent)
	mockSocket.CSS().OnFontsUpdated(func(eventData *css.FontsUpdatedEvent) {
		resultChan <- eventData
	})
	mockResult := &css.FontsUpdatedEvent{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "CSS.fontsUpdated",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *css.FontsUpdatedEvent)
	mockSocket.CSS().OnFontsUpdated(func(eventData *css.FontsUpdatedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "CSS.fontsUpdated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSOnMediaQueryResultChanged(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *css.MediaQueryResultChangedEvent)
	mockSocket.CSS().OnMediaQueryResultChanged(func(eventData *css.MediaQueryResultChangedEvent) {
		resultChan <- eventData
	})
	mockResult := &css.MediaQueryResultChangedEvent{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "CSS.mediaQueryResultChanged",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *css.MediaQueryResultChangedEvent)
	mockSocket.CSS().OnMediaQueryResultChanged(func(eventData *css.MediaQueryResultChangedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "CSS.mediaQueryResultChanged",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSOnStyleSheetAdded(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *css.StyleSheetAddedEvent)
	mockSocket.CSS().OnStyleSheetAdded(func(eventData *css.StyleSheetAddedEvent) {
		resultChan <- eventData
	})
	mockResult := &css.StyleSheetAddedEvent{
		Header: &css.StyleSheetHeader{
			StyleSheetID: css.StyleSheetID("stylesheet-id"),
			FrameID:      page.FrameID("frame-id"),
			SourceURL:    "http://source-url",
			Origin:       css.StyleSheetOrigin("stylesheet-origin"),
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
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "CSS.styleSheetAdded",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected success, got error: '%s'", result.Err.Error())
	}
	if mockResult.Header.StyleSheetID != result.Header.StyleSheetID {
		t.Errorf("Expected '%s', got '%s'", mockResult.Header.StyleSheetID, result.Header.StyleSheetID)
	}

	resultChan = make(chan *css.StyleSheetAddedEvent)
	mockSocket.CSS().OnStyleSheetAdded(func(eventData *css.StyleSheetAddedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "CSS.styleSheetAdded",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSOnStyleSheetChanged(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *css.StyleSheetChangedEvent)
	mockSocket.CSS().OnStyleSheetChanged(func(eventData *css.StyleSheetChangedEvent) {
		resultChan <- eventData
	})
	mockResult := &css.StyleSheetChangedEvent{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "CSS.styleSheetChanged",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected success, got error: '%s'", result.Err.Error())
	}
	if mockResult.StyleSheetID != result.StyleSheetID {
		t.Errorf("Expected '%s', got '%s'", mockResult.StyleSheetID, result.StyleSheetID)
	}

	resultChan = make(chan *css.StyleSheetChangedEvent)
	mockSocket.CSS().OnStyleSheetChanged(func(eventData *css.StyleSheetChangedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "CSS.styleSheetChanged",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestCSSOnStyleSheetRemoved(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *css.StyleSheetRemovedEvent)
	mockSocket.CSS().OnStyleSheetRemoved(func(eventData *css.StyleSheetRemovedEvent) {
		resultChan <- eventData
	})
	mockResult := &css.StyleSheetRemovedEvent{
		StyleSheetID: css.StyleSheetID("stylesheet-id"),
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "CSS.styleSheetRemoved",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected success, got error: '%s'", result.Err.Error())
	}
	if mockResult.StyleSheetID != result.StyleSheetID {
		t.Errorf("Expected '%s', got '%s'", mockResult.StyleSheetID, result.StyleSheetID)
	}

	resultChan = make(chan *css.StyleSheetRemovedEvent)
	mockSocket.CSS().OnStyleSheetRemoved(func(eventData *css.StyleSheetRemovedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "CSS.styleSheetRemoved",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
