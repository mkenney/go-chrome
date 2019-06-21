package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/mkenney/go-chrome/tot/dom"
	"github.com/mkenney/go-chrome/tot/dom/snapshot"
	"github.com/mkenney/go-chrome/tot/page"
)

func TestDOMSnapshotDisable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDOMSnapshotDisable")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMSnapshot().Disable()
	mockResult := &snapshot.DisableResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.DOMSnapshot().Disable()
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMSnapshotEnable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDOMSnapshotEnable")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMSnapshot().Enable()
	mockResult := &snapshot.EnableResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.DOMSnapshot().Enable()
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMSnapshotGet(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestDOMSnapshotGet")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMSnapshot().Get(&snapshot.GetParams{
		ComputedStyleWhitelist: []string{"one", "two"},
	})
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
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.DOMNodes[0].NodeType != result.DOMNodes[0].NodeType {
		t.Errorf("Expected '%d', got '%d'", mockResult.DOMNodes[0].NodeType, result.DOMNodes[0].NodeType)
	}

	resultChan = mockSocket.DOMSnapshot().Get(&snapshot.GetParams{
		ComputedStyleWhitelist: []string{"one", "two"},
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
