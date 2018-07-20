package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/mkenney/go-chrome/tot/io"
	"github.com/mkenney/go-chrome/tot/tracing"
)

func TestTracingEnd(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestTracingEnd")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Tracing().End()
	mockResult := &tracing.EndResult{}
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

	resultChan = mockSocket.Tracing().End()
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

func TestTracingGetCategories(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestTracingGetCategories")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Tracing().GetCategories()
	mockResult := &tracing.GetCategoriesResult{
		Categories: []string{"cat1", "cat2"},
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
	if mockResult.Categories[0] != result.Categories[0] {
		t.Errorf("Expected %s, got %s", mockResult.Categories[0], result.Categories[0])
	}

	resultChan = mockSocket.Tracing().GetCategories()
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

func TestTracingRecordClockSyncMarker(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestTracingRecordClockSyncMarker")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	params := &tracing.RecordClockSyncMarkerParams{
		SyncID: "SyncID",
	}
	resultChan := mockSocket.Tracing().RecordClockSyncMarker(params)
	mockResult := &tracing.RecordClockSyncMarkerResult{}
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

	resultChan = mockSocket.Tracing().RecordClockSyncMarker(params)
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

func TestTracingRequestMemoryDump(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestTracingRequestMemoryDump")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Tracing().RequestMemoryDump()
	mockResult := &tracing.RequestMemoryDumpResult{
		DumpGUID: "DumpGUID",
		Success:  true,
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
	if mockResult.DumpGUID != result.DumpGUID {
		t.Errorf("Expected %s, got %s", mockResult.DumpGUID, result.DumpGUID)
	}

	resultChan = mockSocket.Tracing().RequestMemoryDump()
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

func TestTracingStart(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestTracingStart")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	params := &tracing.StartParams{
		Categories: "Categories",
		Options:    "Options",
		BufferUsageReportingInterval: 1,
		TransferMode:                 tracing.TransferMode.ReportEvents,
		TraceConfig: &tracing.TraceConfig{
			RecordMode:           tracing.RecordMode.RecordUntilFull,
			EnableSampling:       true,
			EnableSystrace:       true,
			EnableArgumentFilter: true,
			IncludedCategories:   []string{"IncludedCategories1", "IncludedCategories2"},
			ExcludedCategories:   []string{"ExcludedCategories1", "ExcludedCategories2"},
			SyntheticDelays:      []string{"SyntheticDelays1", "SyntheticDelays2"},
			MemoryDumpConfig:     tracing.MemoryDumpConfig{"key": "value"},
		},
	}
	resultChan := mockSocket.Tracing().Start(params)
	mockResult := &tracing.StartResult{}
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

	resultChan = mockSocket.Tracing().Start(params)
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

func TestTracingOnBufferUsage(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestTracingOnBufferUsage")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *tracing.BufferUsageEvent)
	mockSocket.Tracing().OnBufferUsage(func(eventData *tracing.BufferUsageEvent) {
		resultChan <- eventData
	})
	mockResult := &tracing.BufferUsageEvent{
		PercentFull: 1,
		EventCount:  1,
		Value:       1,
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Tracing.bufferUsage",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.PercentFull != result.PercentFull {
		t.Errorf("Expected %f, got %f", mockResult.PercentFull, result.PercentFull)
	}

	resultChan = make(chan *tracing.BufferUsageEvent)
	mockSocket.Tracing().OnBufferUsage(func(eventData *tracing.BufferUsageEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Tracing.bufferUsage",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTracingOnDataCollected(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestTracingOnDataCollected")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *tracing.DataCollectedEvent)
	mockSocket.Tracing().OnDataCollected(func(eventData *tracing.DataCollectedEvent) {
		resultChan <- eventData
	})
	mockResult := &tracing.DataCollectedEvent{
		Value: []map[string]string{{"key": "value"}},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Tracing.dataCollected",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Value[0]["key"] != result.Value[0]["key"] {
		t.Errorf("Expected %s, got %s", mockResult.Value[0]["key"], result.Value[0]["key"])
	}

	resultChan = make(chan *tracing.DataCollectedEvent)
	mockSocket.Tracing().OnDataCollected(func(eventData *tracing.DataCollectedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Tracing.dataCollected",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTracingOnTracingComplete(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestTracingOnTracingComplete")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *tracing.CompleteEvent)
	mockSocket.Tracing().OnTracingComplete(func(eventData *tracing.CompleteEvent) {
		resultChan <- eventData
	})
	mockResult := &tracing.CompleteEvent{
		Stream: io.StreamHandle("StreamHandle"),
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Tracing.tracingComplete",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Stream != result.Stream {
		t.Errorf("Expected %s, got %s", mockResult.Stream, result.Stream)
	}

	resultChan = make(chan *tracing.CompleteEvent)
	mockSocket.Tracing().OnTracingComplete(func(eventData *tracing.CompleteEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Tracing.tracingComplete",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
