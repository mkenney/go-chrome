package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	dom "github.com/mkenney/go-chrome/tot/cdtp/dom"
	page "github.com/mkenney/go-chrome/tot/cdtp/page"
	runtime "github.com/mkenney/go-chrome/tot/cdtp/runtime"
)

func TestDOMCollectClassNamesFromSubtree(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.CollectClassNamesFromSubtreeParams{
		NodeID: dom.NodeID(1),
	}
	resultChan := mockSocket.DOM().CollectClassNamesFromSubtree(params)
	mockResult := &dom.CollectClassNamesFromSubtreeResult{
		ClassNames: []string{"class1", "class2"},
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
	if mockResult.ClassNames[0] != result.ClassNames[0] {
		t.Errorf("Expected '%s', got '%s'", mockResult.ClassNames[0], result.ClassNames[0])
	}

	resultChan = mockSocket.DOM().CollectClassNamesFromSubtree(params)
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

func TestDOMCopyTo(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.CopyToParams{
		NodeID:             dom.NodeID(1),
		TargetNodeID:       dom.NodeID(1),
		InsertBeforeNodeID: dom.NodeID(1),
	}
	resultChan := mockSocket.DOM().CopyTo(params)
	mockResult := &dom.CopyToResult{
		NodeID: dom.NodeID(2),
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
	if mockResult.NodeID != result.NodeID {
		t.Errorf("Expected %d, got %d", mockResult.NodeID, result.NodeID)
	}

	resultChan = mockSocket.DOM().CopyTo(params)
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

func TestDOMDescribeNode(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.DescribeNodeParams{
		NodeID:        dom.NodeID(1),
		BackendNodeID: dom.BackendNodeID(1),
		ObjectID:      runtime.RemoteObjectID("remote-object-id"),
		Depth:         1,
		Pierce:        true,
	}
	resultChan := mockSocket.DOM().DescribeNode(params)
	mockResult := &dom.DescribeNodeResult{
		NodeID: dom.NodeID(1),
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
	if mockResult.NodeID != result.NodeID {
		t.Errorf("Expected %d, got %d", mockResult.NodeID, result.NodeID)
	}

	resultChan = mockSocket.DOM().DescribeNode(params)
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

func TestDOMDisable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOM().Disable()
	mockResult := &dom.DisableResult{}
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

	resultChan = mockSocket.DOM().Disable()
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

func TestDOMDiscardSearchResults(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.DiscardSearchResultsParams{
		Node: &dom.Node{
			NodeID:           dom.NodeID(1),
			ParentID:         dom.NodeID(2),
			BackendNodeID:    dom.BackendNodeID(1),
			NodeType:         1,
			NodeName:         "node-name",
			LocalName:        "local-name",
			NodeValue:        "node-value",
			ChildNodeCount:   1,
			Children:         []*dom.Node{},
			Attributes:       []string{"attr1", "attr2"},
			DocumentURL:      "http://document.url",
			BaseURL:          "http://base.url",
			PublicID:         "public-id",
			SystemID:         "system-id",
			InternalSubset:   "internal-subset",
			XMLVersion:       "v1",
			Name:             "name",
			Value:            "value",
			PseudoType:       dom.PseudoType("pseudo-type"),
			ShadowRootType:   dom.ShadowRootType("shadow-root-type"),
			FrameID:          page.FrameID("frame-id"),
			ContentDocument:  &dom.Node{},
			ShadowRoots:      []*dom.Node{},
			TemplateContent:  &dom.Node{},
			PseudoElements:   []*dom.Node{},
			ImportedDocument: &dom.Node{},
			DistributedNodes: []*dom.BackendNode{{
				NodeType:      1,
				NodeName:      "node name",
				BackendNodeID: dom.BackendNodeID(1),
			}},
			IsSVG: true,
		},
	}
	resultChan := mockSocket.DOM().DiscardSearchResults(params)
	mockResult := &dom.DiscardSearchResultsResult{}
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

	resultChan = mockSocket.DOM().DiscardSearchResults(params)
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

func TestDOMEnable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOM().Enable()
	mockResult := &dom.EnableResult{}
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

	resultChan = mockSocket.DOM().Enable()
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

func TestDOMFocus(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.FocusParams{
		NodeID:        dom.NodeID(1),
		BackendNodeID: dom.BackendNodeID(1),
		ObjectID:      runtime.RemoteObjectID("remote-object-id"),
	}
	resultChan := mockSocket.DOM().Focus(params)
	mockResult := &dom.FocusResult{}
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

	resultChan = mockSocket.DOM().Focus(params)
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

func TestDOMGetAttributes(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.GetAttributesParams{
		NodeID: dom.NodeID(1),
	}
	resultChan := mockSocket.DOM().GetAttributes(params)
	mockResult := &dom.GetAttributesResult{
		Attributes: []string{"attr1", "attr2"},
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
	if mockResult.Attributes[0] != result.Attributes[0] {
		t.Errorf("Expected '%s', got '%s'", mockResult.Attributes[0], result.Attributes[0])
	}

	resultChan = mockSocket.DOM().GetAttributes(params)
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

func TestDOMGetBoxModel(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.GetBoxModelParams{
		NodeID:        dom.NodeID(1),
		BackendNodeID: dom.BackendNodeID(1),
		ObjectID:      runtime.RemoteObjectID("remote-object-id"),
	}
	resultChan := mockSocket.DOM().GetBoxModel(params)
	mockResult := &dom.GetBoxModelResult{
		Model: &dom.BoxModel{
			Content: dom.Quad{1, 2},
			Padding: dom.Quad{1, 2},
			Border:  dom.Quad{1, 2},
			Margin:  dom.Quad{1, 2},
			Width:   1,
			Height:  1,
			ShapeOutside: &dom.ShapeOutsideInfo{
				Bounds:      dom.Quad{1, 2},
				Shape:       []interface{}{"box"},
				MarginShape: []interface{}{"box"},
			},
		},
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
	if mockResult.Model.Content != result.Model.Content {
		t.Errorf("Expected '%v', got '%v'", mockResult.Model.Content, result.Model.Content)
	}

	resultChan = mockSocket.DOM().GetBoxModel(params)
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

func TestDOMGetDocument(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.GetDocumentParams{
		Depth:  1,
		Pierce: true,
	}
	resultChan := mockSocket.DOM().GetDocument(params)
	mockResult := &dom.GetDocumentResult{
		Root: &dom.Node{
			NodeID:           dom.NodeID(1),
			ParentID:         dom.NodeID(2),
			BackendNodeID:    dom.BackendNodeID(1),
			NodeType:         1,
			NodeName:         "node-name",
			LocalName:        "local-name",
			NodeValue:        "node-value",
			ChildNodeCount:   1,
			Children:         []*dom.Node{},
			Attributes:       []string{"attr1", "attr2"},
			DocumentURL:      "http://document.url",
			BaseURL:          "http://base.url",
			PublicID:         "public-id",
			SystemID:         "system-id",
			InternalSubset:   "internal-subset",
			XMLVersion:       "v1",
			Name:             "name",
			Value:            "value",
			PseudoType:       dom.PseudoType("pseudo-type"),
			ShadowRootType:   dom.ShadowRootType("shadow-root-type"),
			FrameID:          page.FrameID("frame-id"),
			ContentDocument:  &dom.Node{},
			ShadowRoots:      []*dom.Node{},
			TemplateContent:  &dom.Node{},
			PseudoElements:   []*dom.Node{},
			ImportedDocument: &dom.Node{},
			DistributedNodes: []*dom.BackendNode{{
				NodeType:      1,
				NodeName:      "node name",
				BackendNodeID: dom.BackendNodeID(1),
			}},
			IsSVG: true,
		},
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
	if mockResult.Root.NodeID != result.Root.NodeID {
		t.Errorf("Expected %d, got %d", mockResult.Root.NodeID, result.Root.NodeID)
	}

	resultChan = mockSocket.DOM().GetDocument(params)
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

func TestDOMGetFlattenedDocument(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.GetFlattenedDocumentParams{
		Depth:  1,
		Pierce: true,
	}
	resultChan := mockSocket.DOM().GetFlattenedDocument(params)
	mockResult := &dom.GetFlattenedDocumentResult{
		Nodes: []*dom.Node{{
			NodeID:           dom.NodeID(1),
			ParentID:         dom.NodeID(2),
			BackendNodeID:    dom.BackendNodeID(1),
			NodeType:         1,
			NodeName:         "node-name",
			LocalName:        "local-name",
			NodeValue:        "node-value",
			ChildNodeCount:   1,
			Children:         []*dom.Node{},
			Attributes:       []string{"attr1", "attr2"},
			DocumentURL:      "http://document.url",
			BaseURL:          "http://base.url",
			PublicID:         "public-id",
			SystemID:         "system-id",
			InternalSubset:   "internal-subset",
			XMLVersion:       "v1",
			Name:             "name",
			Value:            "value",
			PseudoType:       dom.PseudoType("pseudo-type"),
			ShadowRootType:   dom.ShadowRootType("shadow-root-type"),
			FrameID:          page.FrameID("frame-id"),
			ContentDocument:  &dom.Node{},
			ShadowRoots:      []*dom.Node{},
			TemplateContent:  &dom.Node{},
			PseudoElements:   []*dom.Node{},
			ImportedDocument: &dom.Node{},
			DistributedNodes: []*dom.BackendNode{{
				NodeType:      1,
				NodeName:      "node name",
				BackendNodeID: dom.BackendNodeID(1),
			}},
			IsSVG: true,
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
	if mockResult.Nodes[0].NodeID != result.Nodes[0].NodeID {
		t.Errorf("Expected %d, got %d", mockResult.Nodes[0].NodeID, result.Nodes[0].NodeID)
	}

	resultChan = mockSocket.DOM().GetFlattenedDocument(params)
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

func TestDOMGetNodeForLocation(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.GetNodeForLocationParams{
		X: 1,
		Y: 1,
		IncludeUserAgentShadowDOM: true,
	}
	resultChan := mockSocket.DOM().GetNodeForLocation(params)
	mockResult := &dom.GetNodeForLocationResult{
		NodeID: dom.NodeID(1),
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
	if mockResult.NodeID != result.NodeID {
		t.Errorf("Expected %d, got %d", mockResult.NodeID, result.NodeID)
	}

	resultChan = mockSocket.DOM().GetNodeForLocation(params)
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

func TestDOMGetOuterHTML(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.GetOuterHTMLParams{
		NodeID:        dom.NodeID(1),
		BackendNodeID: dom.BackendNodeID(1),
		ObjectID:      runtime.RemoteObjectID("remote-object-id"),
	}
	resultChan := mockSocket.DOM().GetOuterHTML(params)
	mockResult := &dom.GetOuterHTMLResult{
		OuterHTML: "<somehtml></somehtml>",
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
	if mockResult.OuterHTML != result.OuterHTML {
		t.Errorf("Expected '%s', got '%s'", mockResult.OuterHTML, result.OuterHTML)
	}

	resultChan = mockSocket.DOM().GetOuterHTML(params)
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

func TestDOMGetRelayoutBoundary(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.GetRelayoutBoundaryParams{
		NodeID: dom.NodeID(1),
	}
	resultChan := mockSocket.DOM().GetRelayoutBoundary(params)
	mockResult := &dom.GetRelayoutBoundaryResult{
		NodeID: dom.NodeID(1),
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
	if mockResult.NodeID != result.NodeID {
		t.Errorf("Expected %d, got %d", mockResult.NodeID, result.NodeID)
	}

	resultChan = mockSocket.DOM().GetRelayoutBoundary(params)
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

func TestDOMGetSearchResults(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.GetSearchResultsParams{
		SearchID:  "search-id",
		FromIndex: 1,
		ToIndex:   2,
	}
	resultChan := mockSocket.DOM().GetSearchResults(params)
	mockResult := &dom.GetSearchResultsResult{
		NodeIDs: []dom.NodeID{1, 2},
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
	if mockResult.NodeIDs[0] != result.NodeIDs[0] {
		t.Errorf("Expected %d, got %d", mockResult.NodeIDs[0], result.NodeIDs[0])
	}

	resultChan = mockSocket.DOM().GetSearchResults(params)
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

func TestDOMMarkUndoableState(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOM().MarkUndoableState()
	mockResult := &dom.MarkUndoableStateResult{}
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

	resultChan = mockSocket.DOM().MarkUndoableState()
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

func TestDOMMoveTo(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.MoveToParams{
		NodeID:             dom.NodeID(1),
		TargetNodeID:       dom.NodeID(2),
		InsertBeforeNodeID: dom.NodeID(3),
	}
	resultChan := mockSocket.DOM().MoveTo(params)
	mockResult := &dom.MoveToResult{
		NodeID: dom.NodeID(1),
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
	if mockResult.NodeID != result.NodeID {
		t.Errorf("Expected %d, got %d", mockResult.NodeID, result.NodeID)
	}

	resultChan = mockSocket.DOM().MoveTo(params)
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

func TestDOMPerformSearch(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.PerformSearchParams{
		Query: "search query",
		IncludeUserAgentShadowDOM: true,
	}
	resultChan := mockSocket.DOM().PerformSearch(params)
	mockResult := &dom.PerformSearchResult{
		SearchID:    "search-id",
		ResultCount: 1,
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
	if mockResult.SearchID != result.SearchID {
		t.Errorf("Expected '%s', got '%s'", mockResult.SearchID, result.SearchID)
	}

	resultChan = mockSocket.DOM().PerformSearch(params)
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

func TestDOMPushNodeByPathToFrontend(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.PushNodeByPathToFrontendParams{
		Path: "/node/path",
	}
	resultChan := mockSocket.DOM().PushNodeByPathToFrontend(params)
	mockResult := &dom.PushNodeByPathToFrontendResult{
		NodeID: dom.NodeID(1),
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
	if mockResult.NodeID != result.NodeID {
		t.Errorf("Expected %d, got %d", mockResult.NodeID, result.NodeID)
	}

	resultChan = mockSocket.DOM().PushNodeByPathToFrontend(params)
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

func TestDOMPushNodesByBackendIDsToFrontend(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.PushNodesByBackendIDsToFrontendParams{
		BackendNodeIDs: []dom.BackendNodeID{1, 2},
	}
	resultChan := mockSocket.DOM().PushNodesByBackendIDsToFrontend(params)
	mockResult := &dom.PushNodesByBackendIDsToFrontendResult{
		NodeIDs: []dom.NodeID{1, 2},
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
	if mockResult.NodeIDs[0] != result.NodeIDs[0] {
		t.Errorf("Expected %d, got %d", mockResult.NodeIDs[0], result.NodeIDs[0])
	}

	resultChan = mockSocket.DOM().PushNodesByBackendIDsToFrontend(params)
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

func TestDOMQuerySelector(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.QuerySelectorParams{
		NodeID:   dom.NodeID(1),
		Selector: "selector",
	}
	resultChan := mockSocket.DOM().QuerySelector(params)
	mockResult := &dom.QuerySelectorResult{
		NodeID: dom.NodeID(1),
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
	if mockResult.NodeID != result.NodeID {
		t.Errorf("Expected %d, got %d", mockResult.NodeID, result.NodeID)
	}

	resultChan = mockSocket.DOM().QuerySelector(params)
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

func TestDOMQuerySelectorAll(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.QuerySelectorAllParams{
		NodeID:   dom.NodeID(1),
		Selector: "selector",
	}
	resultChan := mockSocket.DOM().QuerySelectorAll(params)
	mockResult := &dom.QuerySelectorAllResult{
		NodeIDs: []dom.NodeID{1, 2},
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
	if mockResult.NodeIDs[0] != result.NodeIDs[0] {
		t.Errorf("Expected %d, got %d", mockResult.NodeIDs[0], result.NodeIDs[0])
	}

	resultChan = mockSocket.DOM().QuerySelectorAll(params)
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

func TestDOMRedo(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOM().Redo()
	mockResult := &dom.RedoResult{}
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

	resultChan = mockSocket.DOM().Redo()
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

func TestDOMRemoveAttribute(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.RemoveAttributeParams{
		NodeID: dom.NodeID(1),
		Name:   "attribute-name",
	}
	resultChan := mockSocket.DOM().RemoveAttribute(params)
	mockResult := &dom.RemoveAttributeResult{}
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

	resultChan = mockSocket.DOM().RemoveAttribute(params)
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

func TestDOMRemoveNode(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.RemoveNodeParams{
		NodeID: dom.NodeID(1),
	}
	resultChan := mockSocket.DOM().RemoveNode(params)
	mockResult := &dom.RemoveNodeResult{}
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

	resultChan = mockSocket.DOM().RemoveNode(params)
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

func TestDOMRequestChildNodes(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.RequestChildNodesParams{
		NodeID: dom.NodeID(1),
		Depth:  1,
		Pierce: true,
	}
	resultChan := mockSocket.DOM().RequestChildNodes(params)
	mockResult := &dom.RequestChildNodesResult{}
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

	resultChan = mockSocket.DOM().RequestChildNodes(params)
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

func TestDOMRequestNode(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.RequestNodeParams{
		ObjectID: runtime.RemoteObjectID("remote-object-id"),
	}
	resultChan := mockSocket.DOM().RequestNode(params)
	mockResult := &dom.RequestNodeResult{
		NodeID: dom.NodeID(1),
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
	if mockResult.NodeID != result.NodeID {
		t.Errorf("Expected %d, got %d", mockResult.NodeID, result.NodeID)
	}

	resultChan = mockSocket.DOM().RequestNode(params)
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

func TestDOMResolveNode(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.ResolveNodeParams{
		NodeID:        dom.NodeID(1),
		BackendNodeID: dom.BackendNodeID(1),
		ObjectGroup:   "object-group",
	}
	resultChan := mockSocket.DOM().ResolveNode(params)
	mockResult := &dom.ResolveNodeResult{
		Object: &runtime.RemoteObject{
			Type:                runtime.ObjectType.Object,
			Subtype:             runtime.ObjectSubtype.Array,
			ClassName:           "class-name",
			Value:               "some-value",
			UnserializableValue: runtime.UnserializableValue.Infinity,
			Description:         "description",
			ObjectID:            runtime.RemoteObjectID("remote-object-id"),
			Preview: &runtime.ObjectPreview{
				Type:        runtime.ObjectType.Object,
				Subtype:     runtime.ObjectSubtype.Array,
				Description: "description",
				Overflow:    true,
				Properties:  []*runtime.PropertyPreview{},
				Entries:     []*runtime.EntryPreview{},
			},
			CustomPreview: &runtime.CustomPreview{
				Header:                     "header",
				HasBody:                    true,
				FormatterObjectID:          runtime.RemoteObjectID("remote-object-id"),
				BindRemoteObjectFunctionID: runtime.RemoteObjectID("remote-object-id"),
				ConfigObjectID:             runtime.RemoteObjectID("remote-object-id"),
			},
		},
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
	if mockResult.Object.Type != result.Object.Type {
		t.Errorf("Expected '%s', got '%s'", mockResult.Object.Type, result.Object.Type)
	}

	resultChan = mockSocket.DOM().ResolveNode(params)
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

func TestDOMSetAttributeValue(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.SetAttributeValueParams{
		NodeID: dom.NodeID(1),
		Name:   "name",
		Value:  "value",
	}
	resultChan := mockSocket.DOM().SetAttributeValue(params)
	mockResult := &dom.SetAttributeValueResult{}
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

	resultChan = mockSocket.DOM().SetAttributeValue(params)
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

func TestDOMSetAttributesAsText(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.SetAttributesAsTextParams{
		NodeID: dom.NodeID(1),
		Text:   "some text",
		Name:   "name",
	}
	resultChan := mockSocket.DOM().SetAttributesAsText(params)
	mockResult := &dom.SetAttributesAsTextResult{}
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

	resultChan = mockSocket.DOM().SetAttributesAsText(params)
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

func TestDOMSetFileInputFiles(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.SetFileInputFilesParams{
		Files:         []string{"file1", "file2"},
		NodeID:        dom.NodeID(1),
		BackendNodeID: dom.BackendNodeID(1),
		ObjectID:      runtime.RemoteObjectID("remote-object-id"),
	}
	resultChan := mockSocket.DOM().SetFileInputFiles(params)
	mockResult := &dom.SetFileInputFilesResult{}
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

	resultChan = mockSocket.DOM().SetFileInputFiles(params)
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

func TestDOMSetInspectedNode(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.SetInspectedNodeParams{
		NodeID: dom.NodeID(1),
	}
	resultChan := mockSocket.DOM().SetInspectedNode(params)
	mockResult := &dom.SetInspectedNodeResult{}
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

	resultChan = mockSocket.DOM().SetInspectedNode(params)
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

func TestDOMSetNodeName(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.SetNodeNameParams{
		NodeID: dom.NodeID(1),
		Name:   "new-name",
	}
	resultChan := mockSocket.DOM().SetNodeName(params)
	mockResult := &dom.SetNodeNameResult{
		NodeID: dom.NodeID(1),
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
	if mockResult.NodeID != result.NodeID {
		t.Errorf("Expected %d, got %d", mockResult.NodeID, result.NodeID)
	}

	resultChan = mockSocket.DOM().SetNodeName(params)
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

func TestDOMSetNodeValue(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.SetNodeValueParams{
		NodeID: dom.NodeID(1),
		Value:  "new-value",
	}
	resultChan := mockSocket.DOM().SetNodeValue(params)
	mockResult := &dom.SetNodeValueResult{}
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

	resultChan = mockSocket.DOM().SetNodeValue(params)
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

func TestDOMSetOuterHTML(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &dom.SetOuterHTMLParams{
		NodeID:    dom.NodeID(1),
		OuterHTML: "<outerhtml></outerhtml>",
	}
	resultChan := mockSocket.DOM().SetOuterHTML(params)
	mockResult := &dom.SetOuterHTMLResult{}
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

	resultChan = mockSocket.DOM().SetOuterHTML(params)
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

func TestDOMUndo(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.DOM().Undo()
	mockResult := &dom.UndoResult{}
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

	resultChan = mockSocket.DOM().Undo()
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

func TestDOMOnAttributeModified(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *dom.AttributeModifiedEvent)
	mockSocket.DOM().OnAttributeModified(func(eventData *dom.AttributeModifiedEvent) {
		resultChan <- eventData
	})
	mockResult := &dom.AttributeModifiedEvent{
		NodeID: dom.NodeID(1),
		Name:   "attribute-name",
		Value:  "attribute-value",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "DOM.attributeModified",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *dom.AttributeModifiedEvent)
	mockSocket.DOM().OnAttributeModified(func(eventData *dom.AttributeModifiedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "DOM.attributeModified",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMOnAttributeRemoved(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *dom.AttributeRemovedEvent)
	mockSocket.DOM().OnAttributeRemoved(func(eventData *dom.AttributeRemovedEvent) {
		resultChan <- eventData
	})
	mockResult := &dom.AttributeRemovedEvent{
		NodeID: dom.NodeID(1),
		Name:   "attribute-name",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "DOM.attributeRemoved",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *dom.AttributeRemovedEvent)
	mockSocket.DOM().OnAttributeRemoved(func(eventData *dom.AttributeRemovedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "DOM.attributeRemoved",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMOnCharacterDataModified(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *dom.CharacterDataModifiedEvent)
	mockSocket.DOM().OnCharacterDataModified(func(eventData *dom.CharacterDataModifiedEvent) {
		resultChan <- eventData
	})
	mockResult := &dom.CharacterDataModifiedEvent{
		NodeID:        dom.NodeID(1),
		CharacterData: "char-data",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "DOM.characterDataModified",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *dom.CharacterDataModifiedEvent)
	mockSocket.DOM().OnCharacterDataModified(func(eventData *dom.CharacterDataModifiedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "DOM.characterDataModified",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMOnChildNodeCountUpdated(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *dom.ChildNodeCountUpdatedEvent)
	mockSocket.DOM().OnChildNodeCountUpdated(func(eventData *dom.ChildNodeCountUpdatedEvent) {
		resultChan <- eventData
	})
	mockResult := &dom.ChildNodeCountUpdatedEvent{
		ParentNodeID:   dom.NodeID(2),
		PreviousNodeID: dom.NodeID(1),
		Node:           &dom.Node{},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "DOM.childNodeCountUpdated",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *dom.ChildNodeCountUpdatedEvent)
	mockSocket.DOM().OnChildNodeCountUpdated(func(eventData *dom.ChildNodeCountUpdatedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "DOM.childNodeCountUpdated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMOnChildNodeInserted(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *dom.ChildNodeInsertedEvent)
	mockSocket.DOM().OnChildNodeInserted(func(eventData *dom.ChildNodeInsertedEvent) {
		resultChan <- eventData
	})
	mockResult := &dom.ChildNodeInsertedEvent{
		ParentNodeID:   dom.NodeID(1),
		PreviousNodeID: dom.NodeID(1),
		Node:           &dom.Node{},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "DOM.childNodeInserted",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *dom.ChildNodeInsertedEvent)
	mockSocket.DOM().OnChildNodeInserted(func(eventData *dom.ChildNodeInsertedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "DOM.childNodeInserted",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMOnChildNodeRemoved(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *dom.ChildNodeRemovedEvent)
	mockSocket.DOM().OnChildNodeRemoved(func(eventData *dom.ChildNodeRemovedEvent) {
		resultChan <- eventData
	})
	mockResult := &dom.ChildNodeRemovedEvent{
		ParentNodeID: dom.NodeID(1),
		NodeID:       dom.NodeID(1),
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "DOM.childNodeRemoved",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *dom.ChildNodeRemovedEvent)
	mockSocket.DOM().OnChildNodeRemoved(func(eventData *dom.ChildNodeRemovedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "DOM.childNodeRemoved",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMOnDistributedNodesUpdated(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *dom.DistributedNodesUpdatedEvent)
	mockSocket.DOM().OnDistributedNodesUpdated(func(eventData *dom.DistributedNodesUpdatedEvent) {
		resultChan <- eventData
	})
	mockResult := &dom.DistributedNodesUpdatedEvent{
		InsertionPointID: dom.NodeID(1),
		DistributedNodes: []*dom.BackendNode{},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "DOM.distributedNodesUpdated",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *dom.DistributedNodesUpdatedEvent)
	mockSocket.DOM().OnDistributedNodesUpdated(func(eventData *dom.DistributedNodesUpdatedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "DOM.distributedNodesUpdated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMOnDocumentUpdated(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *dom.DocumentUpdatedEvent)
	mockSocket.DOM().OnDocumentUpdated(func(eventData *dom.DocumentUpdatedEvent) {
		resultChan <- eventData
	})
	mockResult := &dom.DocumentUpdatedEvent{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "DOM.documentUpdated",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *dom.DocumentUpdatedEvent)
	mockSocket.DOM().OnDocumentUpdated(func(eventData *dom.DocumentUpdatedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "DOM.documentUpdated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMOnInlineStyleInvalidated(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *dom.InlineStyleInvalidatedEvent)
	mockSocket.DOM().OnInlineStyleInvalidated(func(eventData *dom.InlineStyleInvalidatedEvent) {
		resultChan <- eventData
	})
	mockResult := &dom.InlineStyleInvalidatedEvent{
		NodeIDs: []dom.NodeID{1, 2},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "DOM.inlineStyleInvalidated",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *dom.InlineStyleInvalidatedEvent)
	mockSocket.DOM().OnInlineStyleInvalidated(func(eventData *dom.InlineStyleInvalidatedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "DOM.inlineStyleInvalidated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMOnPseudoElementAdded(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *dom.PseudoElementAddedEvent)
	mockSocket.DOM().OnPseudoElementAdded(func(eventData *dom.PseudoElementAddedEvent) {
		resultChan <- eventData
	})
	mockResult := &dom.PseudoElementAddedEvent{
		ParentID:      dom.NodeID(1),
		PseudoElement: &dom.Node{},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "DOM.pseudoElementAdded",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *dom.PseudoElementAddedEvent)
	mockSocket.DOM().OnPseudoElementAdded(func(eventData *dom.PseudoElementAddedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "DOM.pseudoElementAdded",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMOnPseudoElementRemoved(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *dom.PseudoElementRemovedEvent)
	mockSocket.DOM().OnPseudoElementRemoved(func(eventData *dom.PseudoElementRemovedEvent) {
		resultChan <- eventData
	})
	mockResult := &dom.PseudoElementRemovedEvent{
		ParentID:        dom.NodeID(1),
		PseudoElementID: dom.NodeID(1),
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "DOM.pseudoElementRemoved",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *dom.PseudoElementRemovedEvent)
	mockSocket.DOM().OnPseudoElementRemoved(func(eventData *dom.PseudoElementRemovedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "DOM.pseudoElementRemoved",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMOnSetChildNodes(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *dom.SetChildNodesEvent)
	mockSocket.DOM().OnSetChildNodes(func(eventData *dom.SetChildNodesEvent) {
		resultChan <- eventData
	})
	mockResult := &dom.SetChildNodesEvent{
		ParentID: dom.NodeID(1),
		Nodes:    []*dom.Node{},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "DOM.setChildNodes",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *dom.SetChildNodesEvent)
	mockSocket.DOM().OnSetChildNodes(func(eventData *dom.SetChildNodesEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "DOM.setChildNodes",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMOnShadowRootPopped(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *dom.ShadowRootPoppedEvent)
	mockSocket.DOM().OnShadowRootPopped(func(eventData *dom.ShadowRootPoppedEvent) {
		resultChan <- eventData
	})
	mockResult := &dom.ShadowRootPoppedEvent{
		HostID: dom.NodeID(1),
		RootID: dom.NodeID(1),
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "DOM.shadowRootPopped",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *dom.ShadowRootPoppedEvent)
	mockSocket.DOM().OnShadowRootPopped(func(eventData *dom.ShadowRootPoppedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "DOM.shadowRootPopped",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMOnShadowRootPushed(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *dom.ShadowRootPushedEvent)
	mockSocket.DOM().OnShadowRootPushed(func(eventData *dom.ShadowRootPushedEvent) {
		resultChan <- eventData
	})
	mockResult := &dom.ShadowRootPushedEvent{
		HostID: dom.NodeID(1),
		Root:   &dom.Node{},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "DOM.shadowRootPushed",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *dom.ShadowRootPushedEvent)
	mockSocket.DOM().OnShadowRootPushed(func(eventData *dom.ShadowRootPushedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "DOM.shadowRootPushed",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
