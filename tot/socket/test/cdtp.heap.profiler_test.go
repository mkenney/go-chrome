package test

import (
	"encoding/json"
	"net/url"
	"testing"

	heapProfiler "github.com/mkenney/go-chrome/tot/cdtp/heap/profiler"
	runtime "github.com/mkenney/go-chrome/tot/cdtp/runtime"
	"github.com/mkenney/go-chrome/tot/socket"
)

func TestHeapProfilerAddInspectedHeapObject(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &heapProfiler.AddInspectedHeapObjectParams{
		HeapObjectID: heapProfiler.HeapSnapshotObjectID("snapshot-id"),
	}
	resultChan := mockSocket.HeapProfiler().AddInspectedHeapObject(params)
	mockResult := &heapProfiler.AddInspectedHeapObjectResult{}
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

	resultChan = mockSocket.HeapProfiler().AddInspectedHeapObject(params)
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

func TestHeapProfilerCollectGarbage(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.HeapProfiler().CollectGarbage()
	mockResult := &heapProfiler.CollectGarbageResult{}
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

	resultChan = mockSocket.HeapProfiler().CollectGarbage()
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

func TestHeapProfilerDisable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.HeapProfiler().Disable()
	mockResult := &heapProfiler.DisableResult{}
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

	resultChan = mockSocket.HeapProfiler().Disable()
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

func TestHeapProfilerEnable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.HeapProfiler().Enable()
	mockResult := &heapProfiler.EnableResult{}
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

	resultChan = mockSocket.HeapProfiler().Enable()
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

func TestHeapProfilerGetHeapObjectID(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &heapProfiler.GetHeapObjectIDParams{
		ObjectID: runtime.RemoteObjectID("remote-object-id"),
	}
	resultChan := mockSocket.HeapProfiler().GetHeapObjectID(params)
	mockResult := &heapProfiler.GetHeapObjectIDResult{
		HeapSnapshotObjectID: heapProfiler.HeapSnapshotObjectID("heap-snapshot-id"),
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
	if mockResult.HeapSnapshotObjectID != result.HeapSnapshotObjectID {
		t.Errorf("Expected '%s', got '%s'", mockResult.HeapSnapshotObjectID, result.HeapSnapshotObjectID)
	}

	resultChan = mockSocket.HeapProfiler().GetHeapObjectID(params)
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

func TestHeapProfilerGetObjectByHeapObjectID(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &heapProfiler.GetObjectByHeapObjectIDParams{
		ObjectID: heapProfiler.HeapSnapshotObjectID("heap-snapshot-id"),
	}
	resultChan := mockSocket.HeapProfiler().GetObjectByHeapObjectID(params)
	mockResult := &heapProfiler.GetObjectByHeapObjectIDResult{
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
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &socket.Error{},
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

func TestHeapProfilerGetSamplingProfile(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &heapProfiler.GetSamplingProfileParams{
		Profile: &heapProfiler.SamplingHeapProfile{
			Head: &heapProfiler.SamplingHeapProfileNode{
				CallFrame: &runtime.CallFrame{
					FunctionName: "funcName",
					ScriptID:     runtime.ScriptID("script-id"),
					URL:          "http://test.url",
					LineNumber:   1,
					ColumnNumber: 1,
				},
				SelfSize: 1,
				Children: []*heapProfiler.SamplingHeapProfileNode{},
			},
		},
	}
	resultChan := mockSocket.HeapProfiler().GetSamplingProfile(params)
	mockResult := &heapProfiler.GetSamplingProfileResult{}
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

	resultChan = mockSocket.HeapProfiler().GetSamplingProfile(params)
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
func TestHeapProfilerStartSampling(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &heapProfiler.StartSamplingParams{
		SamplingInterval: 1,
	}
	resultChan := mockSocket.HeapProfiler().StartSampling(params)
	mockResult := &heapProfiler.StartSamplingResult{}
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

	resultChan = mockSocket.HeapProfiler().StartSampling(params)
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

func TestHeapProfilerStartTrackingHeapObjects(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &heapProfiler.StartTrackingHeapObjectsParams{
		TrackAllocations: true,
	}
	resultChan := mockSocket.HeapProfiler().StartTrackingHeapObjects(params)
	mockResult := &heapProfiler.StartTrackingHeapObjectsResult{}
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

	resultChan = mockSocket.HeapProfiler().StartTrackingHeapObjects(params)
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

func TestHeapProfilerStopSampling(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &heapProfiler.StopSamplingParams{
		Profile: &heapProfiler.SamplingHeapProfile{
			Head: &heapProfiler.SamplingHeapProfileNode{
				CallFrame: &runtime.CallFrame{
					FunctionName: "funcName",
					ScriptID:     runtime.ScriptID("script-id"),
					URL:          "http://test.url",
					LineNumber:   1,
					ColumnNumber: 1,
				},
				SelfSize: 1,
				Children: []*heapProfiler.SamplingHeapProfileNode{},
			},
		},
	}
	resultChan := mockSocket.HeapProfiler().StopSampling(params)
	mockResult := &heapProfiler.StopSamplingResult{}
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

	resultChan = mockSocket.HeapProfiler().StopSampling(params)
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

func TestHeapProfilerStopTrackingHeapObjects(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &heapProfiler.StopTrackingHeapObjectsParams{
		ReportProgress: true,
	}
	resultChan := mockSocket.HeapProfiler().StopTrackingHeapObjects(params)
	mockResult := &heapProfiler.StopTrackingHeapObjectsResult{}
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

	resultChan = mockSocket.HeapProfiler().StopTrackingHeapObjects(params)
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

func TestHeapProfilerTakeHeapSnapshot(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &heapProfiler.TakeHeapSnapshotParams{
		ReportProgress: true,
	}
	resultChan := mockSocket.HeapProfiler().TakeHeapSnapshot(params)
	mockResult := &heapProfiler.TakeHeapSnapshotResult{}
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

	resultChan = mockSocket.HeapProfiler().TakeHeapSnapshot(params)
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

func TestHeapProfilerOnAddHeapSnapshotChunk(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *heapProfiler.AddHeapSnapshotChunkEvent)
	mockSocket.HeapProfiler().OnAddHeapSnapshotChunk(func(eventData *heapProfiler.AddHeapSnapshotChunkEvent) {
		resultChan <- eventData
	})
	mockResult := &heapProfiler.AddHeapSnapshotChunkEvent{
		Chunk: "chunk",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     0,
		Error:  &socket.Error{},
		Method: "HeapProfiler.addHeapSnapshotChunk",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *heapProfiler.AddHeapSnapshotChunkEvent)
	mockSocket.HeapProfiler().OnAddHeapSnapshotChunk(func(eventData *heapProfiler.AddHeapSnapshotChunkEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&socket.Response{
		ID: 0,
		Error: &socket.Error{
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
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *heapProfiler.HeapStatsUpdateEvent)
	mockSocket.HeapProfiler().OnHeapStatsUpdate(func(eventData *heapProfiler.HeapStatsUpdateEvent) {
		resultChan <- eventData
	})
	mockResult := &heapProfiler.HeapStatsUpdateEvent{
		StatsUpdate: []int{1, 2},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     0,
		Error:  &socket.Error{},
		Method: "HeapProfiler.heapStatsUpdate",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *heapProfiler.HeapStatsUpdateEvent)
	mockSocket.HeapProfiler().OnHeapStatsUpdate(func(eventData *heapProfiler.HeapStatsUpdateEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&socket.Response{
		ID: 0,
		Error: &socket.Error{
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
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *heapProfiler.LastSeenObjectIDEvent)
	mockSocket.HeapProfiler().OnLastSeenObjectID(func(eventData *heapProfiler.LastSeenObjectIDEvent) {
		resultChan <- eventData
	})
	mockResult := &heapProfiler.LastSeenObjectIDEvent{
		LastSeenObjectID: 1,
		Timestamp:        1,
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     0,
		Error:  &socket.Error{},
		Method: "HeapProfiler.lastSeenObjectID",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *heapProfiler.LastSeenObjectIDEvent)
	mockSocket.HeapProfiler().OnLastSeenObjectID(func(eventData *heapProfiler.LastSeenObjectIDEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&socket.Response{
		ID: 0,
		Error: &socket.Error{
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
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *heapProfiler.ReportHeapSnapshotProgressEvent)
	mockSocket.HeapProfiler().OnReportHeapSnapshotProgress(func(eventData *heapProfiler.ReportHeapSnapshotProgressEvent) {
		resultChan <- eventData
	})
	mockResult := &heapProfiler.ReportHeapSnapshotProgressEvent{
		Done:     1,
		Total:    1,
		Finished: true,
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     0,
		Error:  &socket.Error{},
		Method: "HeapProfiler.reportHeapSnapshotProgress",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *heapProfiler.ReportHeapSnapshotProgressEvent)
	mockSocket.HeapProfiler().OnReportHeapSnapshotProgress(func(eventData *heapProfiler.ReportHeapSnapshotProgressEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&socket.Response{
		ID: 0,
		Error: &socket.Error{
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
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *heapProfiler.ResetProfilesEvent)
	mockSocket.HeapProfiler().OnResetProfiles(func(eventData *heapProfiler.ResetProfilesEvent) {
		resultChan <- eventData
	})
	mockResult := &heapProfiler.ResetProfilesEvent{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     0,
		Error:  &socket.Error{},
		Method: "HeapProfiler.resetProfiles",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *heapProfiler.ResetProfilesEvent)
	mockSocket.HeapProfiler().OnResetProfiles(func(eventData *heapProfiler.ResetProfilesEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&socket.Response{
		ID: 0,
		Error: &socket.Error{
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
