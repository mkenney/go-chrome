package socket

import (
	"testing"

	"github.com/mkenney/go-chrome/tot/io"
	"github.com/mkenney/go-chrome/tot/tracing"
)

func TestTracingEnd(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer func() { chrome.Close() }()
	soc := New(chrome.URL)
	defer func() { soc.Stop() }()

	mockResult := &tracing.EndResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, "Tracing.end"})
	result := <-soc.Tracing().End()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, "Tracing.end"})
	result = <-soc.Tracing().End()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTracingGetCategories(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer func() { chrome.Close() }()
	soc := New(chrome.URL)
	defer func() { soc.Stop() }()

	mockResult := &tracing.GetCategoriesResult{
		Categories: []string{"cat1", "cat2"},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, "Tracing.getCategories"})
	result := <-soc.Tracing().GetCategories()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if "cat1" != result.Categories[0] {
		t.Errorf("Expected %s, got %s", "cat1", result.Categories[0])
	}

	chrome.AddData(MockData{0, genericError, nil, "Tracing.getCategories"})
	result = <-soc.Tracing().GetCategories()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTracingRecordClockSyncMarker(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer func() { chrome.Close() }()
	soc := New(chrome.URL)
	defer func() { soc.Stop() }()

	params := &tracing.RecordClockSyncMarkerParams{
		SyncID: "SyncID",
	}
	mockResult := &tracing.RecordClockSyncMarkerResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, "Tracing.recordClockSyncMarker"})
	result := <-soc.Tracing().RecordClockSyncMarker(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, "Tracing.recordClockSyncMarker"})
	result = <-soc.Tracing().RecordClockSyncMarker(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTracingRequestMemoryDump(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer func() { chrome.Close() }()
	soc := New(chrome.URL)
	defer func() { soc.Stop() }()

	mockResult := &tracing.RequestMemoryDumpResult{
		DumpGUID: "DumpGUID",
		Success:  true,
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, "Tracing.requestMemoryDump"})
	result := <-soc.Tracing().RequestMemoryDump()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if "DumpGUID" != result.DumpGUID {
		t.Errorf("Expected %s, got %s", "DumpGUID", result.DumpGUID)
	}

	chrome.AddData(MockData{0, genericError, nil, "Tracing.requestMemoryDump"})
	result = <-soc.Tracing().RequestMemoryDump()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTracingStart(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer func() { chrome.Close() }()
	soc := New(chrome.URL)
	defer func() { soc.Stop() }()

	params := &tracing.StartParams{
		Categories:                   "Categories",
		Options:                      "Options",
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
	mockResult := &tracing.StartResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, "Tracing.start"})
	result := <-soc.Tracing().Start(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, "Tracing.start"})
	result = <-soc.Tracing().Start(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTracingOnBufferUsage(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer func() { chrome.Close() }()
	soc := New(chrome.URL)
	defer func() { soc.Stop() }()

	resultChan := make(chan *tracing.BufferUsageEvent)
	soc.Tracing().OnBufferUsage(func(eventData *tracing.BufferUsageEvent) {
		resultChan <- eventData
	})

	chrome.AddData(MockData{
		Err: &Error{},
		Result: &tracing.BufferUsageEvent{
			PercentFull: 1,
			EventCount:  1,
			Value:       1,
		},
		Method: "Tracing.bufferUsage",
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected '%v', got: '%v'", nil, result)
	}
	if float64(1) != result.PercentFull {
		t.Errorf("Expected %f, got %f", float64(1), result.PercentFull)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Tracing.bufferUsage",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTracingOnDataCollected(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer func() { chrome.Close() }()
	soc := New(chrome.URL)
	defer func() { soc.Stop() }()

	resultChan := make(chan *tracing.DataCollectedEvent)
	soc.Tracing().OnDataCollected(func(eventData *tracing.DataCollectedEvent) {
		resultChan <- eventData
	})

	chrome.AddData(MockData{
		Err: &Error{},
		Result: &tracing.DataCollectedEvent{
			Value: []map[string]string{{"key": "value"}},
		},
		Method: "Tracing.dataCollected",
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected '%v', got: '%v'", nil, result)
	}
	if "value" != result.Value[0]["key"] {
		t.Errorf("Expected %s, got %s", "value", result.Value[0]["key"])
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Tracing.dataCollected",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTracingOnTracingComplete(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer func() { chrome.Close() }()
	soc := New(chrome.URL)
	defer func() { soc.Stop() }()

	resultChan := make(chan *tracing.CompleteEvent)
	soc.Tracing().OnTracingComplete(func(eventData *tracing.CompleteEvent) {
		resultChan <- eventData
	})

	chrome.AddData(MockData{
		Err: &Error{},
		Result: &tracing.CompleteEvent{
			Stream: io.StreamHandle("StreamHandle"),
		},
		Method: "Tracing.tracingComplete",
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected '%v', got: '%v'", nil, result.Err)
	}
	if io.StreamHandle("StreamHandle") != result.Stream {
		t.Errorf("Expected %v, got %v", io.StreamHandle("StreamHandle"), result.Stream)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Tracing.tracingComplete",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
