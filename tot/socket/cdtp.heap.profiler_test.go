package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/mkenney/go-chrome/tot/heap/profiler"
	"github.com/mkenney/go-chrome/tot/runtime"
)

func TestHeapProfilerAddInspectedHeapObject(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestHeapProfilerAddInspectedHeapObject")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	params := &profiler.AddInspectedHeapObjectParams{
		HeapObjectID: profiler.HeapSnapshotObjectID("snapshot-id"),
	}
	resultChan := mockSocket.HeapProfiler().AddInspectedHeapObject(params)
	mockResult := &profiler.AddInspectedHeapObjectResult{}
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

	resultChan = mockSocket.HeapProfiler().AddInspectedHeapObject(params)
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

func TestHeapProfilerCollectGarbage(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestHeapProfilerCollectGarbage")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	resultChan := mockSocket.HeapProfiler().CollectGarbage()
	mockResult := &profiler.CollectGarbageResult{}
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

	resultChan = mockSocket.HeapProfiler().CollectGarbage()
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

func TestHeapProfilerDisable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestHeapProfilerDisable")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	resultChan := mockSocket.HeapProfiler().Disable()
	mockResult := &profiler.DisableResult{}
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

	resultChan = mockSocket.HeapProfiler().Disable()
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

func TestHeapProfilerEnable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestHeapProfilerEnable")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	resultChan := mockSocket.HeapProfiler().Enable()
	mockResult := &profiler.EnableResult{}
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

	resultChan = mockSocket.HeapProfiler().Enable()
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

func TestHeapProfilerGetHeapObjectID(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestHeapProfilerGetHeapObjectID")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	params := &profiler.GetHeapObjectIDParams{
		ObjectID: runtime.RemoteObjectID("remote-object-id"),
	}
	resultChan := mockSocket.HeapProfiler().GetHeapObjectID(params)
	mockResult := &profiler.GetHeapObjectIDResult{
		HeapSnapshotObjectID: profiler.HeapSnapshotObjectID("heap-snapshot-id"),
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
	if mockResult.HeapSnapshotObjectID != result.HeapSnapshotObjectID {
		t.Errorf("Expected '%s', got '%s'", mockResult.HeapSnapshotObjectID, result.HeapSnapshotObjectID)
	}

	resultChan = mockSocket.HeapProfiler().GetHeapObjectID(params)
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

func TestHeapProfilerGetObjectByHeapObjectID(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestHeapProfilerGetObjectByHeapObjectID")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	params := &profiler.GetObjectByHeapObjectIDParams{
		ObjectID: profiler.HeapSnapshotObjectID("heap-snapshot-id"),
	}
	resultChan := mockSocket.HeapProfiler().GetObjectByHeapObjectID(params)
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
	if mockResult.Result.Type != result.Result.Type {
		t.Errorf("Expected '%s', got '%s'", mockResult.Result.Type, result.Result.Type)
	}

	resultChan = mockSocket.HeapProfiler().GetObjectByHeapObjectID(params)
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

func TestHeapProfilerGetSamplingProfile(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestHeapProfilerGetSamplingProfile")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

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
	resultChan := mockSocket.HeapProfiler().GetSamplingProfile(params)
	mockResult := &profiler.GetSamplingProfileResult{}
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

	resultChan = mockSocket.HeapProfiler().GetSamplingProfile(params)
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
func TestHeapProfilerStartSampling(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestHeapProfilerStartSampling")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	params := &profiler.StartSamplingParams{
		SamplingInterval: 1,
	}
	resultChan := mockSocket.HeapProfiler().StartSampling(params)
	mockResult := &profiler.StartSamplingResult{}
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

	resultChan = mockSocket.HeapProfiler().StartSampling(params)
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

func TestHeapProfilerStartTrackingHeapObjects(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestHeapProfilerStartTrackingHeapObjects")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	params := &profiler.StartTrackingHeapObjectsParams{
		TrackAllocations: true,
	}
	resultChan := mockSocket.HeapProfiler().StartTrackingHeapObjects(params)
	mockResult := &profiler.StartTrackingHeapObjectsResult{}
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

	resultChan = mockSocket.HeapProfiler().StartTrackingHeapObjects(params)
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

func TestHeapProfilerStopSampling(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestHeapProfilerStopSampling")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

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
	resultChan := mockSocket.HeapProfiler().StopSampling(params)
	mockResult := &profiler.StopSamplingResult{}
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

	resultChan = mockSocket.HeapProfiler().StopSampling(params)
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

func TestHeapProfilerStopTrackingHeapObjects(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestHeapProfilerStopTrackingHeapObjects")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	params := &profiler.StopTrackingHeapObjectsParams{
		ReportProgress: true,
	}
	resultChan := mockSocket.HeapProfiler().StopTrackingHeapObjects(params)
	mockResult := &profiler.StopTrackingHeapObjectsResult{}
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

	resultChan = mockSocket.HeapProfiler().StopTrackingHeapObjects(params)
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

func TestHeapProfilerTakeHeapSnapshot(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestHeapProfilerTakeHeapSnapshot")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	params := &profiler.TakeHeapSnapshotParams{
		ReportProgress: true,
	}
	resultChan := mockSocket.HeapProfiler().TakeHeapSnapshot(params)
	mockResult := &profiler.TakeHeapSnapshotResult{}
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

	resultChan = mockSocket.HeapProfiler().TakeHeapSnapshot(params)
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

func TestHeapProfilerOnAddHeapSnapshotChunk(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestHeapProfilerOnAddHeapSnapshotChunk")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	resultChan := make(chan *profiler.AddHeapSnapshotChunkEvent)
	mockSocket.HeapProfiler().OnAddHeapSnapshotChunk(func(eventData *profiler.AddHeapSnapshotChunkEvent) {
		resultChan <- eventData
	})
	mockResult := &profiler.AddHeapSnapshotChunkEvent{
		Chunk: "chunk",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "HeapProfiler.addHeapSnapshotChunk",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *profiler.AddHeapSnapshotChunkEvent)
	mockSocket.HeapProfiler().OnAddHeapSnapshotChunk(func(eventData *profiler.AddHeapSnapshotChunkEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "HeapProfiler.addHeapSnapshotChunk",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestHeapProfilerOnHeapStatsUpdate(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestHeapProfilerOnHeapStatsUpdate")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	resultChan := make(chan *profiler.HeapStatsUpdateEvent)
	mockSocket.HeapProfiler().OnHeapStatsUpdate(func(eventData *profiler.HeapStatsUpdateEvent) {
		resultChan <- eventData
	})
	mockResult := &profiler.HeapStatsUpdateEvent{
		StatsUpdate: []int{1, 2},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "HeapProfiler.heapStatsUpdate",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *profiler.HeapStatsUpdateEvent)
	mockSocket.HeapProfiler().OnHeapStatsUpdate(func(eventData *profiler.HeapStatsUpdateEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "HeapProfiler.heapStatsUpdate",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestHeapProfilerOnLastSeenObjectID(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestHeapProfilerOnLastSeenObjectID")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	resultChan := make(chan *profiler.LastSeenObjectIDEvent)
	mockSocket.HeapProfiler().OnLastSeenObjectID(func(eventData *profiler.LastSeenObjectIDEvent) {
		resultChan <- eventData
	})
	mockResult := &profiler.LastSeenObjectIDEvent{
		LastSeenObjectID: 1,
		Timestamp:        1,
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "HeapProfiler.lastSeenObjectID",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *profiler.LastSeenObjectIDEvent)
	mockSocket.HeapProfiler().OnLastSeenObjectID(func(eventData *profiler.LastSeenObjectIDEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "HeapProfiler.lastSeenObjectID",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestHeapProfilerOnReportHeapSnapshotProgress(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestHeapProfilerOnReportHeapSnapshotProgress")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	resultChan := make(chan *profiler.ReportHeapSnapshotProgressEvent)
	mockSocket.HeapProfiler().OnReportHeapSnapshotProgress(func(eventData *profiler.ReportHeapSnapshotProgressEvent) {
		resultChan <- eventData
	})
	mockResult := &profiler.ReportHeapSnapshotProgressEvent{
		Done:     1,
		Total:    1,
		Finished: true,
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "HeapProfiler.reportHeapSnapshotProgress",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *profiler.ReportHeapSnapshotProgressEvent)
	mockSocket.HeapProfiler().OnReportHeapSnapshotProgress(func(eventData *profiler.ReportHeapSnapshotProgressEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "HeapProfiler.reportHeapSnapshotProgress",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestHeapProfilerOnResetProfiles(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestHeapProfilerOnResetProfiles")
	mockSocket := NewMock(socketURL)
	go func() {_ = mockSocket.Listen()}()
	defer mockSocket.Stop()

	resultChan := make(chan *profiler.ResetProfilesEvent)
	mockSocket.HeapProfiler().OnResetProfiles(func(eventData *profiler.ResetProfilesEvent) {
		resultChan <- eventData
	})
	mockResult := &profiler.ResetProfilesEvent{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "HeapProfiler.resetProfiles",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *profiler.ResetProfilesEvent)
	mockSocket.HeapProfiler().OnResetProfiles(func(eventData *profiler.ResetProfilesEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "HeapProfiler.resetProfiles",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
