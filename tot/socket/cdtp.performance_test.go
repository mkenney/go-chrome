package socket

import (
	"testing"

	"github.com/mkenney/go-chrome/tot/performance"
)

func TestPerformanceDisable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &performance.DisableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Performance().Disable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Performance().Disable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPerformanceEnable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &performance.EnableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Performance().Enable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Performance().Enable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPerformanceGetMetrics(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &performance.GetMetricsResult{
		Metrics: []*performance.Metric{{
			Name:  "metric name",
			Value: 1,
		}},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Performance().GetMetrics()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Metrics[0].Name != result.Metrics[0].Name {
		t.Errorf("Expected %s, got %s", mockResult.Metrics[0].Name, result.Metrics[0].Name)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Performance().GetMetrics()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestPerformanceOnMetrics(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *performance.MetricsEvent)
	soc.Performance().OnMetrics(func(eventData *performance.MetricsEvent) {
		resultChan <- eventData
	})

	mockResult := &performance.MetricsEvent{
		Metrics: []*performance.Metric{{
			Name:  "metric name",
			Value: 1,
		}},
		Title: "event title",
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Performance.metrics",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Metrics[0].Name != result.Metrics[0].Name {
		t.Errorf("Expected %s, got %s", mockResult.Metrics[0].Name, result.Metrics[0].Name)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Performance.metrics",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
