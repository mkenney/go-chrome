package socket

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/mkenney/go-chrome/protocol/accessibility"
	"github.com/mkenney/go-chrome/protocol/dom"
	log "github.com/sirupsen/logrus"
)

func init() {
	level, err := log.ParseLevel("debug")
	if nil == err {
		log.SetLevel(level)
	}
}
func TestGetPartialAXTree(t *testing.T) {
	MockJSONRead = false
	MockJSONType = "command"
	MockJSONError = false
	MockJSONThrowError = false
	MockJSONData = []byte(`{"nodes":[]}`)

	mockSocket, _ := NewMock("https://www.example.com/")
	go mockSocket.Listen()
	result, err := mockSocket.Accessibility().GetPartialAXTree(
		&accessibility.PartialAXTreeParams{},
	)
	mockSocket.Stop()
	time.Sleep(time.Millisecond * 100) // give it time to stop

	if nil != err {
		t.Errorf("Expected success, got error: %s", err.Error())
	}
	if 0 != len(result.Nodes) {
		tmp, _ := json.Marshal(result.Nodes)
		t.Errorf("Expected empty set, got '%s'", tmp)
	}

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

	MockJSONRead = false
	MockJSONType = "command"
	MockJSONError = false
	MockJSONThrowError = false
	MockJSONData, _ = json.Marshal(mockData)
	go mockSocket.Listen()
	result, err = mockSocket.Accessibility().GetPartialAXTree(
		&accessibility.PartialAXTreeParams{},
	)
	mockSocket.Stop()
	time.Sleep(time.Millisecond * 100) // give it time to stop

	if nil != err {
		t.Errorf("Expected success, got error: %s", err.Error())
	}
	if 0 == len(result.Nodes) {
		tmp, _ := json.Marshal(result.Nodes)
		t.Errorf("Expected dataset, got '%s'", tmp)
	}
}
