package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/mkenney/go-chrome/cdtp/accessibility"
	"github.com/mkenney/go-chrome/cdtp/dom"
)

func TestAccessibilityGetPartialAXTree(t *testing.T) {
	socketURL, _ := url.Parse("https://www.example.com/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Accessibility().GetPartialAXTree(
		&accessibility.PartialAXTreeParams{},
	)
	mockSocket.Conn().AddMockData(
		mockSocket.CurCommandID(),
		&Error{},
		"Accessibility.partialAXTreeParams",
		&accessibility.PartialAXTreeResult{},
	)
	result := <-resultChan
	if nil != result.CDTPError {
		t.Errorf("Expected success, got error: %s", result.CDTPError)
	}
	if 0 != len(result.Nodes) {
		tmp, _ := json.Marshal(result.Nodes)
		t.Errorf("Expected empty set, got '%s'", tmp)
	}

	resultChan = mockSocket.Accessibility().GetPartialAXTree(
		&accessibility.PartialAXTreeParams{},
	)
	mockData := accessibility.PartialAXTreeResult{
		Nodes: []*accessibility.AXNode{{
			NodeID:  accessibility.AXNodeID("NodeID"),
			Ignored: false,
			IgnoredReasons: []*accessibility.AXProperty{{
				Name: "PropertyName",
				Value: &accessibility.AXValue{
					Type:  accessibility.AXValueType("ValueType"),
					Value: nil,
					RelatedNodes: []*accessibility.AXRelatedNode{{
						BackendDOMNodeID: dom.BackendNodeID(1),
						IDRef:            "idref",
						Text:             "some text",
					}},
					Sources: []*accessibility.AXValueSource{{
						Type:              accessibility.AXValueSourceType("SourceType"),
						Value:             &accessibility.AXValue{},
						Attribute:         "Attribute",
						AttributeValue:    &accessibility.AXValue{},
						Superseded:        true,
						NativeSource:      accessibility.AXValueNativeSourceType("NativeSourceType"),
						NativeSourceValue: &accessibility.AXValue{},
						Invalid:           true,
						InvalidReason:     "InvalidReason",
					}},
				},
			}},
			Role:             &accessibility.AXValue{},
			Name:             &accessibility.AXValue{},
			Description:      &accessibility.AXValue{},
			Value:            &accessibility.AXValue{},
			Properties:       []*accessibility.AXProperty{},
			ChildIDs:         []accessibility.AXNodeID{accessibility.AXNodeID(1)},
			BackendDOMNodeID: dom.BackendNodeID(1),
		}},
	}
	mockSocket.Conn().AddMockData(
		mockSocket.CurCommandID(),
		&Error{},
		"Accessibility.partialAXTreeParams",
		mockData,
	)
	result = <-resultChan
	if nil != result.CDTPError {
		t.Errorf("Expected success, got error: %s", result.CDTPError)
	}
	if 0 == len(result.Nodes) {
		tmp, _ := json.Marshal(result.Nodes)
		t.Errorf("Expected dataset, got '%s'", tmp)
	}
}
