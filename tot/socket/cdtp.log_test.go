package socket

import (
	"testing"
	"time"

	"github.com/mkenney/go-chrome/tot/log"
	"github.com/mkenney/go-chrome/tot/network"
	"github.com/mkenney/go-chrome/tot/runtime"
)

func TestLogClear(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &log.ClearResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Log().Clear()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Log().Clear()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestLogDisable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &log.DisableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Log().Disable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Log().Disable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestLogEnable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &log.EnableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Log().Enable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Log().Enable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestLogStartViolationsReport(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &log.StartViolationsReportParams{
		Config: []*log.ViolationSetting{{
			Name:      log.Name.LongTask,
			Threshold: 1,
		}},
	}
	mockResult := &log.StartViolationsReportResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Log().StartViolationsReport(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Log().StartViolationsReport(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestLogStopViolationsReport(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &log.StopViolationsReportResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Log().StopViolationsReport()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Log().StopViolationsReport()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestLogOnEntryAdded(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *log.EntryAddedEvent)
	soc.Log().OnEntryAdded(func(eventData *log.EntryAddedEvent) {
		resultChan <- eventData
	})

	mockResult := &log.EntryAddedEvent{
		Entry: &log.Entry{
			Source:           log.Source.XML,
			Level:            log.Level.Verbose,
			Text:             "entry text",
			Timestamp:        runtime.Timestamp(time.Now().Unix()),
			URL:              "http://some.url",
			LineNumber:       1,
			StackTrace:       &runtime.StackTrace{},
			NetworkRequestID: network.RequestID("request-id"),
			WorkerID:         "worker-id",
			Args:             []*runtime.RemoteObject{{}},
		},
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Log.entryAdded",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Entry.Source != result.Entry.Source {
		t.Errorf("Expected %s, got %s", mockResult.Entry.Source, result.Entry.Source)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Log.entryAdded",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
