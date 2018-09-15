package socket

import (
	"testing"

	"github.com/mkenney/go-chrome/tot/dom"
	"github.com/mkenney/go-chrome/tot/page"
	"github.com/mkenney/go-chrome/tot/runtime"
)

func TestDOMCollectClassNamesFromSubtree(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.CollectClassNamesFromSubtreeParams{
		NodeID: dom.NodeID(1),
	}
	mockResult := &dom.CollectClassNamesFromSubtreeResult{
		ClassNames: []string{"class1", "class2"},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().CollectClassNamesFromSubtree(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.ClassNames[0] != result.ClassNames[0] {
		t.Errorf("Expected '%s', got '%s'", mockResult.ClassNames[0], result.ClassNames[0])
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().CollectClassNamesFromSubtree(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMCopyTo(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.CopyToParams{
		NodeID:             dom.NodeID(1),
		TargetNodeID:       dom.NodeID(1),
		InsertBeforeNodeID: dom.NodeID(1),
	}
	mockResult := &dom.CopyToResult{
		NodeID: dom.NodeID(2),
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().CopyTo(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.NodeID != result.NodeID {
		t.Errorf("Expected %d, got %d", mockResult.NodeID, result.NodeID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().CopyTo(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMDescribeNode(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.DescribeNodeParams{
		NodeID:        dom.NodeID(1),
		BackendNodeID: dom.BackendNodeID(1),
		ObjectID:      runtime.RemoteObjectID("remote-object-id"),
		Depth:         1,
		Pierce:        true,
	}
	mockResult := &dom.DescribeNodeResult{
		NodeID: dom.NodeID(1),
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().DescribeNode(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.NodeID != result.NodeID {
		t.Errorf("Expected %d, got %d", mockResult.NodeID, result.NodeID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().DescribeNode(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMDisable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &dom.DisableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().Disable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().Disable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMDiscardSearchResults(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

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
	mockResult := &dom.DiscardSearchResultsResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().DiscardSearchResults(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().DiscardSearchResults(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMEnable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &dom.EnableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().Enable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().Enable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMFocus(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.FocusParams{
		NodeID:        dom.NodeID(1),
		BackendNodeID: dom.BackendNodeID(1),
		ObjectID:      runtime.RemoteObjectID("remote-object-id"),
	}
	mockResult := &dom.FocusResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().Focus(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().Focus(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMGetAttributes(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.GetAttributesParams{
		NodeID: dom.NodeID(1),
	}
	mockResult := &dom.GetAttributesResult{
		Attributes: []string{"attr1", "attr2"},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().GetAttributes(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Attributes[0] != result.Attributes[0] {
		t.Errorf("Expected '%s', got '%s'", mockResult.Attributes[0], result.Attributes[0])
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().GetAttributes(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMGetBoxModel(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.GetBoxModelParams{
		NodeID:        dom.NodeID(1),
		BackendNodeID: dom.BackendNodeID(1),
		ObjectID:      runtime.RemoteObjectID("remote-object-id"),
	}
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

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().GetBoxModel(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Model.Content != result.Model.Content {
		t.Errorf("Expected '%v', got '%v'", mockResult.Model.Content, result.Model.Content)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().GetBoxModel(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMGetDocument(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.GetDocumentParams{
		Depth:  1,
		Pierce: true,
	}
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

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().GetDocument(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Root.NodeID != result.Root.NodeID {
		t.Errorf("Expected %d, got %d", mockResult.Root.NodeID, result.Root.NodeID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().GetDocument(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMGetFlattenedDocument(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.GetFlattenedDocumentParams{
		Depth:  1,
		Pierce: true,
	}
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

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().GetFlattenedDocument(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Nodes[0].NodeID != result.Nodes[0].NodeID {
		t.Errorf("Expected %d, got %d", mockResult.Nodes[0].NodeID, result.Nodes[0].NodeID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().GetFlattenedDocument(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMGetNodeForLocation(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.GetNodeForLocationParams{
		X:                         1,
		Y:                         1,
		IncludeUserAgentShadowDOM: true,
	}
	mockResult := &dom.GetNodeForLocationResult{
		NodeID: dom.NodeID(1),
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().GetNodeForLocation(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.NodeID != result.NodeID {
		t.Errorf("Expected %d, got %d", mockResult.NodeID, result.NodeID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().GetNodeForLocation(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMGetOuterHTML(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.GetOuterHTMLParams{
		NodeID:        dom.NodeID(1),
		BackendNodeID: dom.BackendNodeID(1),
		ObjectID:      runtime.RemoteObjectID("remote-object-id"),
	}
	mockResult := &dom.GetOuterHTMLResult{
		OuterHTML: "<somehtml></somehtml>",
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().GetOuterHTML(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.OuterHTML != result.OuterHTML {
		t.Errorf("Expected '%s', got '%s'", mockResult.OuterHTML, result.OuterHTML)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().GetOuterHTML(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMGetRelayoutBoundary(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.GetRelayoutBoundaryParams{
		NodeID: dom.NodeID(1),
	}
	mockResult := &dom.GetRelayoutBoundaryResult{
		NodeID: dom.NodeID(1),
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().GetRelayoutBoundary(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.NodeID != result.NodeID {
		t.Errorf("Expected %d, got %d", mockResult.NodeID, result.NodeID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().GetRelayoutBoundary(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMGetSearchResults(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.GetSearchResultsParams{
		SearchID:  "search-id",
		FromIndex: 1,
		ToIndex:   2,
	}
	mockResult := &dom.GetSearchResultsResult{
		NodeIDs: []dom.NodeID{1, 2},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().GetSearchResults(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.NodeIDs[0] != result.NodeIDs[0] {
		t.Errorf("Expected %d, got %d", mockResult.NodeIDs[0], result.NodeIDs[0])
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().GetSearchResults(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMMarkUndoableState(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &dom.MarkUndoableStateResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().MarkUndoableState()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().MarkUndoableState()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMMoveTo(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.MoveToParams{
		NodeID:             dom.NodeID(1),
		TargetNodeID:       dom.NodeID(2),
		InsertBeforeNodeID: dom.NodeID(3),
	}
	mockResult := &dom.MoveToResult{
		NodeID: dom.NodeID(1),
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().MoveTo(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.NodeID != result.NodeID {
		t.Errorf("Expected %d, got %d", mockResult.NodeID, result.NodeID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().MoveTo(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMPerformSearch(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.PerformSearchParams{
		Query:                     "search query",
		IncludeUserAgentShadowDOM: true,
	}
	mockResult := &dom.PerformSearchResult{
		SearchID:    "search-id",
		ResultCount: 1,
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().PerformSearch(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.SearchID != result.SearchID {
		t.Errorf("Expected '%s', got '%s'", mockResult.SearchID, result.SearchID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().PerformSearch(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMPushNodeByPathToFrontend(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.PushNodeByPathToFrontendParams{
		Path: "/node/path",
	}
	mockResult := &dom.PushNodeByPathToFrontendResult{
		NodeID: dom.NodeID(1),
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().PushNodeByPathToFrontend(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.NodeID != result.NodeID {
		t.Errorf("Expected %d, got %d", mockResult.NodeID, result.NodeID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().PushNodeByPathToFrontend(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMPushNodesByBackendIDsToFrontend(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.PushNodesByBackendIDsToFrontendParams{
		BackendNodeIDs: []dom.BackendNodeID{1, 2},
	}
	mockResult := &dom.PushNodesByBackendIDsToFrontendResult{
		NodeIDs: []dom.NodeID{1, 2},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().PushNodesByBackendIDsToFrontend(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.NodeIDs[0] != result.NodeIDs[0] {
		t.Errorf("Expected %d, got %d", mockResult.NodeIDs[0], result.NodeIDs[0])
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().PushNodesByBackendIDsToFrontend(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMQuerySelector(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.QuerySelectorParams{
		NodeID:   dom.NodeID(1),
		Selector: "selector",
	}
	mockResult := &dom.QuerySelectorResult{
		NodeID: dom.NodeID(1),
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().QuerySelector(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.NodeID != result.NodeID {
		t.Errorf("Expected %d, got %d", mockResult.NodeID, result.NodeID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().QuerySelector(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMQuerySelectorAll(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.QuerySelectorAllParams{
		NodeID:   dom.NodeID(1),
		Selector: "selector",
	}
	mockResult := &dom.QuerySelectorAllResult{
		NodeIDs: []dom.NodeID{1, 2},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().QuerySelectorAll(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.NodeIDs[0] != result.NodeIDs[0] {
		t.Errorf("Expected %d, got %d", mockResult.NodeIDs[0], result.NodeIDs[0])
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().QuerySelectorAll(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMRedo(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &dom.RedoResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().Redo()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().Redo()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMRemoveAttribute(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.RemoveAttributeParams{
		NodeID: dom.NodeID(1),
		Name:   "attribute-name",
	}
	mockResult := &dom.RemoveAttributeResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().RemoveAttribute(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().RemoveAttribute(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMRemoveNode(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.RemoveNodeParams{
		NodeID: dom.NodeID(1),
	}
	mockResult := &dom.RemoveNodeResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().RemoveNode(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().RemoveNode(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMRequestChildNodes(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.RequestChildNodesParams{
		NodeID: dom.NodeID(1),
		Depth:  1,
		Pierce: true,
	}
	mockResult := &dom.RequestChildNodesResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().RequestChildNodes(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().RequestChildNodes(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMRequestNode(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.RequestNodeParams{
		ObjectID: runtime.RemoteObjectID("remote-object-id"),
	}
	mockResult := &dom.RequestNodeResult{
		NodeID: dom.NodeID(1),
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().RequestNode(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.NodeID != result.NodeID {
		t.Errorf("Expected %d, got %d", mockResult.NodeID, result.NodeID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().RequestNode(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMResolveNode(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.ResolveNodeParams{
		NodeID:        dom.NodeID(1),
		BackendNodeID: dom.BackendNodeID(1),
		ObjectGroup:   "object-group",
	}
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

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().ResolveNode(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Object.Type != result.Object.Type {
		t.Errorf("Expected '%s', got '%s'", mockResult.Object.Type, result.Object.Type)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().ResolveNode(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMSetAttributeValue(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.SetAttributeValueParams{
		NodeID: dom.NodeID(1),
		Name:   "name",
		Value:  "value",
	}
	mockResult := &dom.SetAttributeValueResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().SetAttributeValue(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().SetAttributeValue(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMSetAttributesAsText(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.SetAttributesAsTextParams{
		NodeID: dom.NodeID(1),
		Text:   "some text",
		Name:   "name",
	}
	mockResult := &dom.SetAttributesAsTextResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().SetAttributesAsText(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().SetAttributesAsText(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMSetFileInputFiles(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.SetFileInputFilesParams{
		Files:         []string{"file1", "file2"},
		NodeID:        dom.NodeID(1),
		BackendNodeID: dom.BackendNodeID(1),
		ObjectID:      runtime.RemoteObjectID("remote-object-id"),
	}
	mockResult := &dom.SetFileInputFilesResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().SetFileInputFiles(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().SetFileInputFiles(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMSetInspectedNode(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.SetInspectedNodeParams{
		NodeID: dom.NodeID(1),
	}
	mockResult := &dom.SetInspectedNodeResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().SetInspectedNode(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().SetInspectedNode(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMSetNodeName(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.SetNodeNameParams{
		NodeID: dom.NodeID(1),
		Name:   "new-name",
	}
	mockResult := &dom.SetNodeNameResult{
		NodeID: dom.NodeID(1),
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().SetNodeName(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.NodeID != result.NodeID {
		t.Errorf("Expected %d, got %d", mockResult.NodeID, result.NodeID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().SetNodeName(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMSetNodeValue(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.SetNodeValueParams{
		NodeID: dom.NodeID(1),
		Value:  "new-value",
	}
	mockResult := &dom.SetNodeValueResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().SetNodeValue(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().SetNodeValue(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMSetOuterHTML(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &dom.SetOuterHTMLParams{
		NodeID:    dom.NodeID(1),
		OuterHTML: "<outerhtml></outerhtml>",
	}
	mockResult := &dom.SetOuterHTMLResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().SetOuterHTML(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().SetOuterHTML(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMUndo(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &dom.UndoResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.DOM().Undo()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.DOM().Undo()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMOnAttributeModified(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *dom.AttributeModifiedEvent)
	soc.DOM().OnAttributeModified(func(eventData *dom.AttributeModifiedEvent) {
		resultChan <- eventData
	})

	mockResult := &dom.AttributeModifiedEvent{
		NodeID: dom.NodeID(1),
		Name:   "attribute-name",
		Value:  "attribute-value",
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "DOM.attributeModified",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "DOM.attributeModified",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMOnAttributeRemoved(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *dom.AttributeRemovedEvent)
	soc.DOM().OnAttributeRemoved(func(eventData *dom.AttributeRemovedEvent) {
		resultChan <- eventData
	})

	mockResult := &dom.AttributeRemovedEvent{
		NodeID: dom.NodeID(1),
		Name:   "attribute-name",
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "DOM.attributeRemoved",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "DOM.attributeRemoved",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMOnCharacterDataModified(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *dom.CharacterDataModifiedEvent)
	soc.DOM().OnCharacterDataModified(func(eventData *dom.CharacterDataModifiedEvent) {
		resultChan <- eventData
	})

	mockResult := &dom.CharacterDataModifiedEvent{
		NodeID:        dom.NodeID(1),
		CharacterData: "char-data",
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "DOM.characterDataModified",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "DOM.characterDataModified",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMOnChildNodeCountUpdated(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *dom.ChildNodeCountUpdatedEvent)
	soc.DOM().OnChildNodeCountUpdated(func(eventData *dom.ChildNodeCountUpdatedEvent) {
		resultChan <- eventData
	})

	mockResult := &dom.ChildNodeCountUpdatedEvent{
		ParentNodeID:   dom.NodeID(2),
		PreviousNodeID: dom.NodeID(1),
		Node:           &dom.Node{},
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "DOM.childNodeCountUpdated",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "DOM.childNodeCountUpdated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMOnChildNodeInserted(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *dom.ChildNodeInsertedEvent)
	soc.DOM().OnChildNodeInserted(func(eventData *dom.ChildNodeInsertedEvent) {
		resultChan <- eventData
	})

	mockResult := &dom.ChildNodeInsertedEvent{
		ParentNodeID:   dom.NodeID(1),
		PreviousNodeID: dom.NodeID(1),
		Node:           &dom.Node{},
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "DOM.childNodeInserted",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "DOM.childNodeInserted",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMOnChildNodeRemoved(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *dom.ChildNodeRemovedEvent)
	soc.DOM().OnChildNodeRemoved(func(eventData *dom.ChildNodeRemovedEvent) {
		resultChan <- eventData
	})

	mockResult := &dom.ChildNodeRemovedEvent{
		ParentNodeID: dom.NodeID(1),
		NodeID:       dom.NodeID(1),
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "DOM.childNodeRemoved",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "DOM.childNodeRemoved",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMOnDistributedNodesUpdated(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *dom.DistributedNodesUpdatedEvent)
	soc.DOM().OnDistributedNodesUpdated(func(eventData *dom.DistributedNodesUpdatedEvent) {
		resultChan <- eventData
	})

	mockResult := &dom.DistributedNodesUpdatedEvent{
		InsertionPointID: dom.NodeID(1),
		DistributedNodes: []*dom.BackendNode{},
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "DOM.distributedNodesUpdated",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "DOM.distributedNodesUpdated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMOnDocumentUpdated(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *dom.DocumentUpdatedEvent)
	soc.DOM().OnDocumentUpdated(func(eventData *dom.DocumentUpdatedEvent) {
		resultChan <- eventData
	})

	mockResult := &dom.DocumentUpdatedEvent{}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "DOM.documentUpdated",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "DOM.documentUpdated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMOnInlineStyleInvalidated(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *dom.InlineStyleInvalidatedEvent)
	soc.DOM().OnInlineStyleInvalidated(func(eventData *dom.InlineStyleInvalidatedEvent) {
		resultChan <- eventData
	})

	mockResult := &dom.InlineStyleInvalidatedEvent{
		NodeIDs: []dom.NodeID{1, 2},
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "DOM.inlineStyleInvalidated",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "DOM.inlineStyleInvalidated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMOnPseudoElementAdded(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *dom.PseudoElementAddedEvent)
	soc.DOM().OnPseudoElementAdded(func(eventData *dom.PseudoElementAddedEvent) {
		resultChan <- eventData
	})

	mockResult := &dom.PseudoElementAddedEvent{
		ParentID:      dom.NodeID(1),
		PseudoElement: &dom.Node{},
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "DOM.pseudoElementAdded",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "DOM.pseudoElementAdded",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMOnPseudoElementRemoved(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *dom.PseudoElementRemovedEvent)
	soc.DOM().OnPseudoElementRemoved(func(eventData *dom.PseudoElementRemovedEvent) {
		resultChan <- eventData
	})

	mockResult := &dom.PseudoElementRemovedEvent{
		ParentID:        dom.NodeID(1),
		PseudoElementID: dom.NodeID(1),
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "DOM.pseudoElementRemoved",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "DOM.pseudoElementRemoved",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMOnSetChildNodes(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *dom.SetChildNodesEvent)
	soc.DOM().OnSetChildNodes(func(eventData *dom.SetChildNodesEvent) {
		resultChan <- eventData
	})

	mockResult := &dom.SetChildNodesEvent{
		ParentID: dom.NodeID(1),
		Nodes:    []*dom.Node{},
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "DOM.setChildNodes",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "DOM.setChildNodes",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMOnShadowRootPopped(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *dom.ShadowRootPoppedEvent)
	soc.DOM().OnShadowRootPopped(func(eventData *dom.ShadowRootPoppedEvent) {
		resultChan <- eventData
	})

	mockResult := &dom.ShadowRootPoppedEvent{
		HostID: dom.NodeID(1),
		RootID: dom.NodeID(1),
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "DOM.shadowRootPopped",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "DOM.shadowRootPopped",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestDOMOnShadowRootPushed(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *dom.ShadowRootPushedEvent)
	soc.DOM().OnShadowRootPushed(func(eventData *dom.ShadowRootPushedEvent) {
		resultChan <- eventData
	})

	mockResult := &dom.ShadowRootPushedEvent{
		HostID: dom.NodeID(1),
		Root:   &dom.Node{},
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "DOM.shadowRootPushed",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "DOM.shadowRootPushed",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
