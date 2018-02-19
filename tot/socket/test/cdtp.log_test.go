package test

import (
	"encoding/json"
	"net/url"
	"testing"
	"time"

	log "github.com/mkenney/go-chrome/tot/cdtp/log"
	network "github.com/mkenney/go-chrome/tot/cdtp/network"
	runtime "github.com/mkenney/go-chrome/tot/cdtp/runtime"
	"github.com/mkenney/go-chrome/tot/socket"
)

func TestLogClear(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Log().Clear()
	mockResult := &log.ClearResult{}
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

	resultChan = mockSocket.Log().Clear()
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

func TestLogDisable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Log().Disable()
	mockResult := &log.DisableResult{}
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

	resultChan = mockSocket.Log().Disable()
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

func TestLogEnable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Log().Enable()
	mockResult := &log.EnableResult{}
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

	resultChan = mockSocket.Log().Enable()
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

func TestLogStartViolationsReport(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &log.StartViolationsReportParams{
		Config: []*log.ViolationSetting{{
			Name:      log.Name.LongTask,
			Threshold: 1,
		}},
	}
	resultChan := mockSocket.Log().StartViolationsReport(params)
	mockResult := &log.StartViolationsReportResult{}
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

	resultChan = mockSocket.Log().StartViolationsReport(params)
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

func TestLogStopViolationsReport(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Log().StopViolationsReport()
	mockResult := &log.StopViolationsReportResult{}
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

	resultChan = mockSocket.Log().StopViolationsReport()
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

func TestLogOnEntryAdded(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *log.EntryAddedEvent)
	mockSocket.Log().OnEntryAdded(func(eventData *log.EntryAddedEvent) {
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
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     0,
		Error:  &socket.Error{},
		Method: "Log.entryAdded",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Entry.Source != result.Entry.Source {
		t.Errorf("Expected %s, got %s", mockResult.Entry.Source, result.Entry.Source)
	}

	resultChan = make(chan *log.EntryAddedEvent)
	mockSocket.Log().OnEntryAdded(func(eventData *log.EntryAddedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&socket.Response{
		ID: 0,
		Error: &socket.Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Log.entryAdded",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
