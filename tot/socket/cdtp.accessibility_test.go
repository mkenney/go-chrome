package socket

import (
	"encoding/json"
	"testing"

	"github.com/mkenney/go-chrome/tot/accessibility"
	"github.com/mkenney/go-chrome/tot/dom"
)

func TestAccessibilityGetPartialAXTree(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer func() { chrome.Close() }()
	soc := New(chrome.URL)
	defer func() { soc.Stop() }()

	params := &accessibility.PartialAXTreeParams{
		NodeID:         dom.NodeID(1),
		FetchRelatives: true,
	}
	mockResult := &accessibility.PartialAXTreeResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Accessibility().GetPartialAXTree(params)
	if nil != result.Err {
		t.Errorf("Expected success, got error: %s", result.Err)
	}
	if 0 != len(result.Nodes) {
		tmp, _ := json.Marshal(result.Nodes)
		t.Errorf("Expected empty set, got '%s'", tmp)
	}

	mockResult = &accessibility.PartialAXTreeResult{
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
	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result = <-soc.Accessibility().GetPartialAXTree(params)
	if nil != result.Err {
		t.Errorf("Expected success, got error: %s", result.Err)
	}
	if 0 == len(result.Nodes) {
		tmp, _ := json.Marshal(result.Nodes)
		t.Errorf("Expected dataset, got '%s'", tmp)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Accessibility().GetPartialAXTree(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}

	if "" == result.Err.Error() {
		t.Errorf("Expected error message, got empty string")
	}
}
