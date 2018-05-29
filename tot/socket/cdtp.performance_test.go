package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	performance "github.com/mkenney/go-chrome/tot/cdtp/performance"
)

func TestPerformanceDisable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestPerformanceDisable")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Performance().Disable()
	mockResult := &performance.DisableResult{}
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

	resultChan = mockSocket.Performance().Disable()
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

func TestPerformanceEnable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestPerformanceEnable")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Performance().Enable()
	mockResult := &performance.EnableResult{}
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

	resultChan = mockSocket.Performance().Enable()
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

func TestPerformanceGetMetrics(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestPerformanceGetMetrics")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Performance().GetMetrics()
	mockResult := &performance.GetMetricsResult{
		Metrics: []*performance.Metric{{
			Name:  "metric name",
			Value: 1,
		}},
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
	if mockResult.Metrics[0].Name != result.Metrics[0].Name {
		t.Errorf("Expected %s, got %s", mockResult.Metrics[0].Name, result.Metrics[0].Name)
	}

	resultChan = mockSocket.Performance().GetMetrics()
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

func TestPerformanceOnMetrics(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestPerformanceOnMetrics")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *performance.MetricsEvent)
	mockSocket.Performance().OnMetrics(func(eventData *performance.MetricsEvent) {
		resultChan <- eventData
	})
	mockResult := &performance.MetricsEvent{
		Metrics: []*performance.Metric{{
			Name:  "metric name",
			Value: 1,
		}},
		Title: "event title",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Performance.metrics",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Metrics[0].Name != result.Metrics[0].Name {
		t.Errorf("Expected %s, got %s", mockResult.Metrics[0].Name, result.Metrics[0].Name)
	}

	resultChan = make(chan *performance.MetricsEvent)
	mockSocket.Performance().OnMetrics(func(eventData *performance.MetricsEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Performance.metrics",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
