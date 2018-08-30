package socket

import (
	"testing"

	"github.com/mkenney/go-chrome/tot/io"
	"github.com/mkenney/go-chrome/tot/tracing"
)

func TestTracingEnd(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	chrome.AddData(MockData{
		&Error{},
		&tracing.EndResult{},
		"",
	})
	result := <-soc.Tracing().End()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{
		&Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		nil,
		"",
	})
	result = <-soc.Tracing().End()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTracingGetCategories(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	chrome.AddData(MockData{
		&Error{},
		&tracing.GetCategoriesResult{
			Categories: []string{"cat1", "cat2"},
		},
		"",
	})
	result := <-soc.Tracing().GetCategories()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if "cat1" != result.Categories[0] {
		t.Errorf("Expected %s, got %s", "cat1", result.Categories[0])
	}

	chrome.AddData(MockData{
		&Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		nil,
		"",
	})
	result = <-soc.Tracing().GetCategories()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTracingRecordClockSyncMarker(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &tracing.RecordClockSyncMarkerParams{
		SyncID: "SyncID",
	}

	chrome.AddData(MockData{
		&Error{},
		&tracing.RecordClockSyncMarkerResult{},
		"",
	})
	result := <-soc.Tracing().RecordClockSyncMarker(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{
		&Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		nil,
		"",
	})
	result = <-soc.Tracing().RecordClockSyncMarker(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTracingRequestMemoryDump(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	chrome.AddData(MockData{
		&Error{},
		&tracing.RequestMemoryDumpResult{
			DumpGUID: "DumpGUID",
			Success:  true,
		},
		"",
	})
	result := <-soc.Tracing().RequestMemoryDump()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if "DumpGUID" != result.DumpGUID {
		t.Errorf("Expected %s, got %s", "DumpGUID", result.DumpGUID)
	}

	chrome.AddData(MockData{
		&Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		nil,
		"",
	})
	result = <-soc.Tracing().RequestMemoryDump()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTracingStart(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

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

	chrome.AddData(MockData{
		&Error{},
		&tracing.StartResult{},
		"",
	})
	result := <-soc.Tracing().Start(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{
		&Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		nil,
		"",
	})
	result = <-soc.Tracing().Start(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTracingOnBufferUsage(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()

	soc := New(chrome.URL)
	defer soc.Stop()
	chrome.AddData(MockData{
		Err: &Error{},
		Result: &tracing.BufferUsageEvent{
			PercentFull: 1,
			EventCount:  1,
			Value:       1,
		},
		Method: "Tracing.bufferUsage",
	})
	resultChan := make(chan *tracing.BufferUsageEvent)
	soc.Tracing().OnBufferUsage(func(eventData *tracing.BufferUsageEvent) {
		resultChan <- eventData
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected '%v', got: '%v'", nil, result)
	}
	if float64(1) != result.PercentFull {
		t.Errorf("Expected %f, got %f", float64(1), result.PercentFull)
	}

	soc = New(chrome.URL)
	defer soc.Stop()
	chrome.AddData(MockData{
		Err: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Result: nil,
		Method: "Tracing.bufferUsage",
	})
	resultChan = make(chan *tracing.BufferUsageEvent)
	soc.Tracing().OnBufferUsage(func(eventData *tracing.BufferUsageEvent) {
		resultChan <- eventData
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
	defer chrome.Close()

	soc := New(chrome.URL)
	defer soc.Stop()
	chrome.AddData(MockData{
		Err: &Error{},
		Result: &tracing.DataCollectedEvent{
			Value: []map[string]string{{"key": "value"}},
		},
		Method: "Tracing.dataCollected",
	})
	resultChan := make(chan *tracing.DataCollectedEvent)
	soc.Tracing().OnDataCollected(func(eventData *tracing.DataCollectedEvent) {
		resultChan <- eventData
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected '%v', got: '%v'", nil, result)
	}
	if "value" != result.Value[0]["key"] {
		t.Errorf("Expected %s, got %s", "value", result.Value[0]["key"])
	}

	soc = New(chrome.URL)
	defer soc.Stop()
	chrome.AddData(MockData{
		Err: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Result: nil,
		Method: "Tracing.dataCollected",
	})
	resultChan = make(chan *tracing.DataCollectedEvent)
	soc.Tracing().OnDataCollected(func(eventData *tracing.DataCollectedEvent) {
		resultChan <- eventData
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
	defer chrome.Close()

	soc := New(chrome.URL)
	defer soc.Stop()
	chrome.AddData(MockData{
		Err: &Error{},
		Result: &tracing.CompleteEvent{
			Stream: io.StreamHandle("StreamHandle"),
		},
		Method: "Tracing.tracingComplete",
	})
	resultChan := make(chan *tracing.CompleteEvent)
	soc.Tracing().OnTracingComplete(func(eventData *tracing.CompleteEvent) {
		resultChan <- eventData
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected '%v', got: '%v'", nil, result.Err)
	}
	if io.StreamHandle("StreamHandle") != result.Stream {
		t.Errorf("Expected %s, got %s", io.StreamHandle("StreamHandle"), result.Stream)
	}

	soc = New(chrome.URL)
	defer soc.Stop()
	chrome.AddData(MockData{
		Err: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Result: nil,
		Method: "Tracing.tracingComplete",
	})
	resultChan = make(chan *tracing.CompleteEvent)
	soc.Tracing().OnTracingComplete(func(eventData *tracing.CompleteEvent) {
		resultChan <- eventData
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
