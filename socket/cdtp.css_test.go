package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	css "github.com/mkenney/go-chrome/cdtp/css"
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
			Selectors: []*css.Value{&css.Value{
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
	if mockResult.ClassNames[0] != mockResult.ClassNames[0] {
		t.Errorf("Expected '%s', got '%s'", mockResult.ClassNames[0], result.ClassNames[0])
	}
}
