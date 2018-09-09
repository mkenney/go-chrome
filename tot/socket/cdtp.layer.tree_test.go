package socket

import (
	"testing"

	"github.com/mkenney/go-chrome/tot/dom"
	"github.com/mkenney/go-chrome/tot/layer/tree"
)

func TestLayerTreeCompositingReasons(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &tree.CompositingReasonsParams{
		LayerID: tree.LayerID("layer-id"),
	}
	mockResult := &tree.CompositingReasonsResult{
		CompositingReasons: []string{"reason1", "reason2"},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.LayerTree().CompositingReasons(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.CompositingReasons[0] != result.CompositingReasons[0] {
		t.Errorf("Expected %s, got %s", mockResult.CompositingReasons[0], result.CompositingReasons[0])
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.LayerTree().CompositingReasons(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestLayerTreeDisable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &tree.DisableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.LayerTree().Disable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.LayerTree().Disable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestLayerTreeEnable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &tree.EnableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.LayerTree().Enable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.LayerTree().Enable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestLayerTreeLoadSnapshot(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &tree.LoadSnapshotParams{
		Tiles: []*tree.PictureTile{{
			X:       1,
			Y:       1,
			Picture: "picture data",
		}},
	}
	mockResult := &tree.LoadSnapshotResult{
		SnapshotID: tree.SnapshotID("snapshot-id"),
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.LayerTree().LoadSnapshot(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.SnapshotID != result.SnapshotID {
		t.Errorf("Expected %s, got %s", mockResult.SnapshotID, result.SnapshotID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.LayerTree().LoadSnapshot(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestLayerTreeMakeSnapshot(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &tree.MakeSnapshotParams{
		LayerID: tree.LayerID("layer-id"),
	}
	mockResult := &tree.MakeSnapshotResult{
		SnapshotID: tree.SnapshotID("snapshot-id"),
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.LayerTree().MakeSnapshot(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.SnapshotID != result.SnapshotID {
		t.Errorf("Expected %s, got %s", mockResult.SnapshotID, result.SnapshotID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.LayerTree().MakeSnapshot(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestLayerTreeProfileSnapshot(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &tree.ProfileSnapshotParams{
		SnapshotID:     tree.SnapshotID("snapshot-id"),
		MinRepeatCount: 1,
		MinDuration:    1,
		ClipRect: &dom.Rect{
			X:      1,
			Y:      1,
			Width:  1,
			Height: 1,
		},
	}
	mockResult := &tree.ProfileSnapshotResult{
		Timings: []*tree.PaintProfile{},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.LayerTree().ProfileSnapshot(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.LayerTree().ProfileSnapshot(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestLayerTreeReleaseSnapshot(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &tree.ReleaseSnapshotParams{
		SnapshotID: tree.SnapshotID("snapshot-id"),
	}
	mockResult := &tree.ReleaseSnapshotResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.LayerTree().ReleaseSnapshot(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.LayerTree().ReleaseSnapshot(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestLayerTreeReplaySnapshot(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &tree.ReplaySnapshotParams{
		SnapshotID: tree.SnapshotID("snapshot-id"),
		FromStep:   1,
		ToStep:     2,
		Scale:      1,
	}
	mockResult := &tree.ReplaySnapshotResult{
		DataURL: "data:urldata",
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.LayerTree().ReplaySnapshot(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.DataURL != result.DataURL {
		t.Errorf("Expected %s, got %s", mockResult.DataURL, result.DataURL)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.LayerTree().ReplaySnapshot(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestLayerTreeSnapshotCommandLog(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &tree.SnapshotCommandLogParams{
		SnapshotID: tree.SnapshotID("snapshot-id"),
	}
	mockResult := &tree.SnapshotCommandLogResult{
		CommandLog: []map[string]string{{"command1": "log1"}},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.LayerTree().SnapshotCommandLog(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.CommandLog[0]["command1"] != result.CommandLog[0]["command1"] {
		t.Errorf("Expected %s, got %s", mockResult.CommandLog[0]["command1"], result.CommandLog[0]["command1"])
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.LayerTree().SnapshotCommandLog(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestLayerTreeOnLayerPainted(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *tree.LayerPaintedEvent)
	soc.LayerTree().OnLayerPainted(func(eventData *tree.LayerPaintedEvent) {
		resultChan <- eventData
	})

	mockResult := &tree.LayerPaintedEvent{
		LayerID: tree.LayerID("layer-id"),
		Clip: &dom.Rect{
			X:      1,
			Y:      1,
			Width:  1,
			Height: 1,
		},
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "LayerTree.layerPainted",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.LayerID != result.LayerID {
		t.Errorf("Expected %s, got %s", mockResult.LayerID, result.LayerID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "LayerTree.layerPainted",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestLayerTreeOnLayerTreeDidChange(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *tree.DidChangeEvent)
	soc.LayerTree().OnLayerTreeDidChange(func(eventData *tree.DidChangeEvent) {
		resultChan <- eventData
	})

	mockResult := &tree.DidChangeEvent{
		Layers: []*tree.Layer{{
			LayerID:       tree.LayerID("layer-id"),
			ParentLayerID: tree.LayerID("parent.layer-id"),
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
			ScrollRects: []*tree.ScrollRect{{
				Rect: &dom.Rect{},
				Type: tree.RectType.RepaintsOnScroll,
			}},
			StickyPositionConstraint: &tree.StickyPositionConstraint{
				StickyBoxRect:                       &dom.Rect{},
				ContainingBlockRect:                 &dom.Rect{},
				NearestLayerShiftingStickyBox:       tree.LayerID("sticky-layer-id"),
				NearestLayerShiftingContainingBlock: tree.LayerID("containing-layer-id"),
			},
		}},
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "LayerTree.layerTreeDidChange",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Layers[0].LayerID != result.Layers[0].LayerID {
		t.Errorf("Expected %s, got %s", mockResult.Layers[0].LayerID, result.Layers[0].LayerID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "LayerTree.layerTreeDidChange",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
