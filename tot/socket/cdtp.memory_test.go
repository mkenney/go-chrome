package socket

import (
	"testing"

	memory "github.com/mkenney/go-chrome/tot/memory"
)

func TestMemoryGetDOMCounters(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &memory.GetDOMCountersParams{
		Documents:        1,
		Nodes:            1,
		JsEventListeners: 1,
	}
	mockResult := &memory.GetDOMCountersResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Memory().GetDOMCounters(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Memory().GetDOMCounters(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestMemoryPrepareForLeakDetection(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &memory.PrepareForLeakDetectionResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Memory().PrepareForLeakDetection()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Memory().PrepareForLeakDetection()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestMemorySetPressureNotificationsSuppressed(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &memory.SetPressureNotificationsSuppressedParams{
		Suppressed: true,
	}
	mockResult := &memory.SetPressureNotificationsSuppressedResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Memory().SetPressureNotificationsSuppressed(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Memory().SetPressureNotificationsSuppressed(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestMemorySimulatePressureNotification(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &memory.SimulatePressureNotificationParams{
		Level: memory.PressureLevel("pressure-level"),
	}
	mockResult := &memory.SimulatePressureNotificationResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Memory().SimulatePressureNotification(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Memory().SimulatePressureNotification(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
