package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	dom "github.com/mkenney/go-chrome/tot/cdtp/dom"
	layerTree "github.com/mkenney/go-chrome/tot/cdtp/layer/tree"
)

func TestLayerTreeCompositingReasons(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestLayerTreeCompositingReasons")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &layerTree.CompositingReasonsParams{
		LayerID: layerTree.LayerID("layer-id"),
	}
	resultChan := mockSocket.LayerTree().CompositingReasons(params)
	mockResult := &layerTree.CompositingReasonsResult{
		CompositingReasons: []string{"reason1", "reason2"},
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
	if mockResult.CompositingReasons[0] != result.CompositingReasons[0] {
		t.Errorf("Expected %s, got %s", mockResult.CompositingReasons[0], result.CompositingReasons[0])
	}

	resultChan = mockSocket.LayerTree().CompositingReasons(params)
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

func TestLayerTreeDisable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestLayerTreeDisable")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.LayerTree().Disable()
	mockResult := &layerTree.DisableResult{}
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

	resultChan = mockSocket.LayerTree().Disable()
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

func TestLayerTreeEnable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestLayerTreeEnable")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.LayerTree().Enable()
	mockResult := &layerTree.EnableResult{}
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

	resultChan = mockSocket.LayerTree().Enable()
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

func TestLayerTreeLoadSnapshot(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestLayerTreeLoadSnapshot")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &layerTree.LoadSnapshotParams{
		Tiles: []*layerTree.PictureTile{{
			X:       1,
			Y:       1,
			Picture: "picture data",
		}},
	}
	resultChan := mockSocket.LayerTree().LoadSnapshot(params)
	mockResult := &layerTree.LoadSnapshotResult{
		SnapshotID: layerTree.SnapshotID("snapshot-id"),
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
	if mockResult.SnapshotID != result.SnapshotID {
		t.Errorf("Expected %s, got %s", mockResult.SnapshotID, result.SnapshotID)
	}

	resultChan = mockSocket.LayerTree().LoadSnapshot(params)
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

func TestLayerTreeMakeSnapshot(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestLayerTreeMakeSnapshot")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &layerTree.MakeSnapshotParams{
		LayerID: layerTree.LayerID("layer-id"),
	}
	resultChan := mockSocket.LayerTree().MakeSnapshot(params)
	mockResult := &layerTree.MakeSnapshotResult{
		SnapshotID: layerTree.SnapshotID("snapshot-id"),
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
	if mockResult.SnapshotID != result.SnapshotID {
		t.Errorf("Expected %s, got %s", mockResult.SnapshotID, result.SnapshotID)
	}

	resultChan = mockSocket.LayerTree().MakeSnapshot(params)
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

func TestLayerTreeProfileSnapshot(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestLayerTreeProfileSnapshot")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &layerTree.ProfileSnapshotParams{
		SnapshotID:     layerTree.SnapshotID("snapshot-id"),
		MinRepeatCount: 1,
		MinDuration:    1,
		ClipRect: &dom.Rect{
			X:      1,
			Y:      1,
			Width:  1,
			Height: 1,
		},
	}
	resultChan := mockSocket.LayerTree().ProfileSnapshot(params)
	mockResult := &layerTree.ProfileSnapshotResult{
		Timings: []*layerTree.PaintProfile{},
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

	resultChan = mockSocket.LayerTree().ProfileSnapshot(params)
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

func TestLayerTreeReleaseSnapshot(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestLayerTreeReleaseSnapshot")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &layerTree.ReleaseSnapshotParams{
		SnapshotID: layerTree.SnapshotID("snapshot-id"),
	}
	resultChan := mockSocket.LayerTree().ReleaseSnapshot(params)
	mockResult := &layerTree.ReleaseSnapshotResult{}
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

	resultChan = mockSocket.LayerTree().ReleaseSnapshot(params)
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

func TestLayerTreeReplaySnapshot(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestLayerTreeReplaySnapshot")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &layerTree.ReplaySnapshotParams{
		SnapshotID: layerTree.SnapshotID("snapshot-id"),
		FromStep:   1,
		ToStep:     2,
		Scale:      1,
	}
	resultChan := mockSocket.LayerTree().ReplaySnapshot(params)
	mockResult := &layerTree.ReplaySnapshotResult{
		DataURL: "data:urldata",
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
	if mockResult.DataURL != result.DataURL {
		t.Errorf("Expected %s, got %s", mockResult.DataURL, result.DataURL)
	}

	resultChan = mockSocket.LayerTree().ReplaySnapshot(params)
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

func TestLayerTreeSnapshotCommandLog(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestLayerTreeSnapshotCommandLog")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &layerTree.SnapshotCommandLogParams{
		SnapshotID: layerTree.SnapshotID("snapshot-id"),
	}
	resultChan := mockSocket.LayerTree().SnapshotCommandLog(params)
	mockResult := &layerTree.SnapshotCommandLogResult{
		CommandLog: []map[string]string{{"command1": "log1"}},
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
	if mockResult.CommandLog[0]["command1"] != result.CommandLog[0]["command1"] {
		t.Errorf("Expected %s, got %s", mockResult.CommandLog[0]["command1"], result.CommandLog[0]["command1"])
	}

	resultChan = mockSocket.LayerTree().SnapshotCommandLog(params)
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

func TestLayerTreeOnLayerPainted(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestLayerTreeOnLayerPainted")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *layerTree.LayerPaintedEvent)
	mockSocket.LayerTree().OnLayerPainted(func(eventData *layerTree.LayerPaintedEvent) {
		resultChan <- eventData
	})
	mockResult := &layerTree.LayerPaintedEvent{
		LayerID: layerTree.LayerID("layer-id"),
		Clip: &dom.Rect{
			X:      1,
			Y:      1,
			Width:  1,
			Height: 1,
		},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "LayerTree.layerPainted",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.LayerID != result.LayerID {
		t.Errorf("Expected %s, got %s", mockResult.LayerID, result.LayerID)
	}

	resultChan = make(chan *layerTree.LayerPaintedEvent)
	mockSocket.LayerTree().OnLayerPainted(func(eventData *layerTree.LayerPaintedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "LayerTree.layerPainted",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestLayerTreeOnLayerTreeDidChange(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestLayerTreeOnLayerTreeDidChange")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *layerTree.DidChangeEvent)
	mockSocket.LayerTree().OnLayerTreeDidChange(func(eventData *layerTree.DidChangeEvent) {
		resultChan <- eventData
	})
	mockResult := &layerTree.DidChangeEvent{
		Layers: []*layerTree.Layer{{
			LayerID:       layerTree.LayerID("layer-id"),
			ParentLayerID: layerTree.LayerID("parent.layer-id"),
			BackendNodeID: dom.BackendNodeID(1),
			OffsetX:       1,
			OffsetY:       1,
			Width:         1,
			Height:        1,
			Transform:     []float64{1, 2, 3},
			AnchorX:       1,
			AnchorY:       1,
			AnchorZ:       1,
			PaintCount:    1,
			DrawsContent:  true,
			Invisible:     true,
			ScrollRects: []*layerTree.ScrollRect{{
				Rect: &dom.Rect{},
				Type: layerTree.RectType.RepaintsOnScroll,
			}},
			StickyPositionConstraint: &layerTree.StickyPositionConstraint{
				StickyBoxRect:                       &dom.Rect{},
				ContainingBlockRect:                 &dom.Rect{},
				NearestLayerShiftingStickyBox:       layerTree.LayerID("sticky-layer-id"),
				NearestLayerShiftingContainingBlock: layerTree.LayerID("containing-layer-id"),
			},
		}},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "LayerTree.layerTreeDidChange",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Layers[0].LayerID != result.Layers[0].LayerID {
		t.Errorf("Expected %s, got %s", mockResult.Layers[0].LayerID, result.Layers[0].LayerID)
	}

	resultChan = make(chan *layerTree.DidChangeEvent)
	mockSocket.LayerTree().OnLayerTreeDidChange(func(eventData *layerTree.DidChangeEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "LayerTree.layerTreeDidChange",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
