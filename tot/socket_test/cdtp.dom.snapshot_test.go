package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	dom "github.com/mkenney/go-chrome/tot/cdtp/dom"
	domSnapshot "github.com/mkenney/go-chrome/tot/cdtp/dom/snapshot"
	page "github.com/mkenney/go-chrome/tot/cdtp/page"
	"github.com/mkenney/go-chrome/tot/socket"
)

func TestDOMSnapshotGet(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOMSnapshot().Get(&domSnapshot.GetParams{
		ComputedStyleWhitelist: []string{"one", "two"},
	})
	mockResult := &domSnapshot.GetResult{
		DOMNodes: []*domSnapshot.DOMNode{{
			NodeType:              1,
			NodeName:              "node-name",
			NodeValue:             "node-value",
			TextValue:             "text-value",
			InputValue:            "input-value",
			InputChecked:          true,
			OptionSelected:        true,
			BackendNodeID:         dom.BackendNodeID(1),
			ChildNodeIndexes:      []int64{1, 2, 3},
			Attributes:            []*domSnapshot.NameValue{},
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
		LayoutTreeNodes: []*domSnapshot.LayoutTreeNode{{
			DomNodeIndex: 1,
			BoundingBox: &dom.Rect{
				X:      1,
				Y:      1,
				Width:  1,
				Height: 1,
			},
			LayoutText: "layout text",
			InlineTextNodes: []*domSnapshot.InlineTextBox{{
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
		ComputedStyles: []*domSnapshot.ComputedStyle{{
			Properties: []*domSnapshot.NameValue{{
				Name:  "name",
				Value: "value",
			}},
		}},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &socket.Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.DOMNodes[0].NodeType != result.DOMNodes[0].NodeType {
		t.Errorf("Expected '%d', got '%d'", mockResult.DOMNodes[0].NodeType, result.DOMNodes[0].NodeType)
	}

	resultChan = mockSocket.DOMSnapshot().Get(&domSnapshot.GetParams{
		ComputedStyleWhitelist: []string{"one", "two"},
	})
	mockSocket.Conn().AddMockData(&socket.Response{
		ID: mockSocket.CurCommandID(),
		Error: &socket.Error{
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
