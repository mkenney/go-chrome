package socket

import (
	"testing"

	"github.com/mkenney/go-chrome/tot/dom"
	"github.com/mkenney/go-chrome/tot/dom/snapshot"
	"github.com/mkenney/go-chrome/tot/page"
)

func TestDOMSnapshotDisable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &snapshot.DisableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOMSnapshot().Disable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOMSnapshot().Disable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMSnapshotEnable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &snapshot.EnableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOMSnapshot().Enable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOMSnapshot().Enable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMSnapshotGet(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &snapshot.GetResult{
		DOMNodes: []*snapshot.DOMNode{{
			NodeType:              1,
			NodeName:              "node-name",
			NodeValue:             "node-value",
			TextValue:             "text-value",
			InputValue:            "input-value",
			InputChecked:          true,
			OptionSelected:        true,
			BackendNodeID:         dom.BackendNodeID(1),
			ChildNodeIndexes:      []int64{1, 2, 3},
			Attributes:            []*snapshot.NameValue{},
			PseudoElementIndexes:  []int64{1, 2, 3},
			LayoutNodeIndex:       1,
			DocumentURL:           "http://document.url",
			BaseURL:               "http://base.url",
			ContentLanguage:       "language",
			DocumentEncoding:      "encoding",
			PublicID:              "public-id",
			SystemID:              "system-id",
			FrameID:               page.FrameID("frame-id"),
			ContentDocumentIndex:  1,
			ImportedDocumentIndex: 1,
			TemplateContentIndex:  1,
			PseudoType:            dom.PseudoType("pseudo-type"),
			IsClickable:           true,
		}},
		LayoutTreeNodes: []*snapshot.LayoutTreeNode{{
			DomNodeIndex: 1,
			BoundingBox: &dom.Rect{
				X:      1,
				Y:      1,
				Width:  1,
				Height: 1,
			},
			LayoutText: "layout text",
			InlineTextNodes: []*snapshot.InlineTextBox{{
				BoundingBox: &dom.Rect{
					X:      1,
					Y:      1,
					Width:  1,
					Height: 1,
				},
				StartCharacterIndex: 1,
				NumCharacters:       1,
			}},
			StyleIndex: 1,
		}},
		ComputedStyles: []*snapshot.ComputedStyle{{
			Properties: []*snapshot.NameValue{{
				Name:  "name",
				Value: "value",
			}},
		}},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOMSnapshot().Get(&snapshot.GetParams{
		ComputedStyleWhitelist: []string{"one", "two"},
	})
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.DOMNodes[0].NodeType != result.DOMNodes[0].NodeType {
		t.Errorf("Expected '%d', got '%d'", mockResult.DOMNodes[0].NodeType, result.DOMNodes[0].NodeType)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOMSnapshot().Get(&snapshot.GetParams{
		ComputedStyleWhitelist: []string{"one", "two"},
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
