package socket

import (
	"testing"

	"github.com/mkenney/go-chrome/tot/heap/profiler"
	"github.com/mkenney/go-chrome/tot/runtime"
)

func TestHeapProfilerAddInspectedHeapObject(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &profiler.AddInspectedHeapObjectParams{
		HeapObjectID: profiler.HeapSnapshotObjectID("snapshot-id"),
	}
	mockResult := &profiler.AddInspectedHeapObjectResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.HeapProfiler().AddInspectedHeapObject(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.HeapProfiler().AddInspectedHeapObject(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestHeapProfilerCollectGarbage(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &profiler.CollectGarbageResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.HeapProfiler().CollectGarbage()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.HeapProfiler().CollectGarbage()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestHeapProfilerDisable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &profiler.DisableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.HeapProfiler().Disable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.HeapProfiler().Disable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestHeapProfilerEnable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &profiler.EnableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.HeapProfiler().Enable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.HeapProfiler().Enable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestHeapProfilerGetHeapObjectID(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &profiler.GetHeapObjectIDParams{
		ObjectID: runtime.RemoteObjectID("remote-object-id"),
	}
	mockResult := &profiler.GetHeapObjectIDResult{
		HeapSnapshotObjectID: profiler.HeapSnapshotObjectID("heap-snapshot-id"),
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.HeapProfiler().GetHeapObjectID(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.HeapSnapshotObjectID != result.HeapSnapshotObjectID {
		t.Errorf("Expected '%s', got '%s'", mockResult.HeapSnapshotObjectID, result.HeapSnapshotObjectID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.HeapProfiler().GetHeapObjectID(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestHeapProfilerGetObjectByHeapObjectID(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &profiler.GetObjectByHeapObjectIDParams{
		ObjectID: profiler.HeapSnapshotObjectID("heap-snapshot-id"),
	}
	mockResult := &profiler.GetObjectByHeapObjectIDResult{
		Result: &runtime.RemoteObject{
			Type:                runtime.ObjectType.Object,
			Subtype:             runtime.ObjectSubtype.Array,
			ClassName:           "class-name",
			Value:               &struct{ a string }{a: "somestring"},
			UnserializableValue: runtime.UnserializableValue.Infinity,
			Description:         "description",
			ObjectID:            runtime.RemoteObjectID("remote-object-id"),
			Preview: &runtime.ObjectPreview{
				Type: runtime.ObjectType.Accessor,
			},
			CustomPreview: &runtime.CustomPreview{},
		},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.HeapProfiler().GetObjectByHeapObjectID(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Result.Type != result.Result.Type {
		t.Errorf("Expected '%s', got '%s'", mockResult.Result.Type, result.Result.Type)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.HeapProfiler().GetObjectByHeapObjectID(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestHeapProfilerGetSamplingProfile(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &profiler.GetSamplingProfileParams{
		Profile: &profiler.SamplingHeapProfile{
			Head: &profiler.SamplingHeapProfileNode{
				CallFrame: &runtime.CallFrame{
					FunctionName: "funcName",
					ScriptID:     runtime.ScriptID("script-id"),
					URL:          "http://test.url",
					LineNumber:   1,
					ColumnNumber: 1,
				},
				SelfSize: 1,
				Children: []*profiler.SamplingHeapProfileNode{},
			},
		},
	}
	mockResult := &profiler.GetSamplingProfileResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.HeapProfiler().GetSamplingProfile(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.HeapProfiler().GetSamplingProfile(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
func TestHeapProfilerStartSampling(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &profiler.StartSamplingParams{
		SamplingInterval: 1,
	}
	mockResult := &profiler.StartSamplingResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.HeapProfiler().StartSampling(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.HeapProfiler().StartSampling(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestHeapProfilerStartTrackingHeapObjects(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &profiler.StartTrackingHeapObjectsParams{
		TrackAllocations: true,
	}
	mockResult := &profiler.StartTrackingHeapObjectsResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.HeapProfiler().StartTrackingHeapObjects(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.HeapProfiler().StartTrackingHeapObjects(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestHeapProfilerStopSampling(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &profiler.StopSamplingParams{
		Profile: &profiler.SamplingHeapProfile{
			Head: &profiler.SamplingHeapProfileNode{
				CallFrame: &runtime.CallFrame{
					FunctionName: "funcName",
					ScriptID:     runtime.ScriptID("script-id"),
					URL:          "http://test.url",
					LineNumber:   1,
					ColumnNumber: 1,
				},
				SelfSize: 1,
				Children: []*profiler.SamplingHeapProfileNode{},
			},
		},
	}
	mockResult := &profiler.StopSamplingResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.HeapProfiler().StopSampling(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.HeapProfiler().StopSampling(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestHeapProfilerStopTrackingHeapObjects(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &profiler.StopTrackingHeapObjectsParams{
		ReportProgress: true,
	}
	mockResult := &profiler.StopTrackingHeapObjectsResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.HeapProfiler().StopTrackingHeapObjects(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.HeapProfiler().StopTrackingHeapObjects(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestHeapProfilerTakeHeapSnapshot(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &profiler.TakeHeapSnapshotParams{
		ReportProgress: true,
	}
	mockResult := &profiler.TakeHeapSnapshotResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.HeapProfiler().TakeHeapSnapshot(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.HeapProfiler().TakeHeapSnapshot(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestHeapProfilerOnAddHeapSnapshotChunk(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *profiler.AddHeapSnapshotChunkEvent)
	soc.HeapProfiler().OnAddHeapSnapshotChunk(func(eventData *profiler.AddHeapSnapshotChunkEvent) {
		resultChan <- eventData
	})

	mockResult := &profiler.AddHeapSnapshotChunkEvent{
		Chunk: "chunk",
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "HeapProfiler.addHeapSnapshotChunk",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "HeapProfiler.addHeapSnapshotChunk",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestHeapProfilerOnHeapStatsUpdate(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *profiler.HeapStatsUpdateEvent)
	soc.HeapProfiler().OnHeapStatsUpdate(func(eventData *profiler.HeapStatsUpdateEvent) {
		resultChan <- eventData
	})

	mockResult := &profiler.HeapStatsUpdateEvent{
		StatsUpdate: []int{1, 2},
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "HeapProfiler.heapStatsUpdate",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "HeapProfiler.heapStatsUpdate",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestHeapProfilerOnLastSeenObjectID(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *profiler.LastSeenObjectIDEvent)
	soc.HeapProfiler().OnLastSeenObjectID(func(eventData *profiler.LastSeenObjectIDEvent) {
		resultChan <- eventData
	})

	mockResult := &profiler.LastSeenObjectIDEvent{
		LastSeenObjectID: 1,
		Timestamp:        1,
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "HeapProfiler.lastSeenObjectID",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "HeapProfiler.lastSeenObjectID",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestHeapProfilerOnReportHeapSnapshotProgress(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *profiler.ReportHeapSnapshotProgressEvent)
	soc.HeapProfiler().OnReportHeapSnapshotProgress(func(eventData *profiler.ReportHeapSnapshotProgressEvent) {
		resultChan <- eventData
	})

	mockResult := &profiler.ReportHeapSnapshotProgressEvent{
		Done:     1,
		Total:    1,
		Finished: true,
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "HeapProfiler.reportHeapSnapshotProgress",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "HeapProfiler.reportHeapSnapshotProgress",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestHeapProfilerOnResetProfiles(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *profiler.ResetProfilesEvent)
	soc.HeapProfiler().OnResetProfiles(func(eventData *profiler.ResetProfilesEvent) {
		resultChan <- eventData
	})

	mockResult := &profiler.ResetProfilesEvent{}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "HeapProfiler.resetProfiles",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "HeapProfiler.resetProfiles",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
