package socket

import (
	"testing"
	"time"

	"github.com/mkenney/go-chrome/tot/headless/experimental"
	"github.com/mkenney/go-chrome/tot/runtime"
)

func TestHeadlessExperimentalBeginFrame(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &experimental.BeginFrameParams{
		FrameTime: runtime.Timestamp(time.Now().Unix()),
		Deadline:  runtime.Timestamp(time.Now().Unix()),
		Interval:  1.1,
		Screenshot: &experimental.ScreenshotParams{
			Format:  experimental.Format.Jpeg,
			Quality: 100,
		},
	}
	mockResult := &experimental.BeginFrameResult{
		HasDamage:               true,
		MainFrameContentUpdated: true,
		ScreenshotData:          "data",
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.HeadlessExperimental().BeginFrame(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.HasDamage != result.HasDamage {
		t.Errorf("Expected %v, got %v", mockResult.HasDamage, result.HasDamage)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.HeadlessExperimental().BeginFrame(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestHeadlessExperimentalDisable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &experimental.DisableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.HeadlessExperimental().Disable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.HeadlessExperimental().Disable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestHeadlessExperimentalEnable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &experimental.EnableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.HeadlessExperimental().Enable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.HeadlessExperimental().Enable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestHeadlessExperimentalOnMainFrameReadyForScreenshots(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *experimental.MainFrameReadyForScreenshotsEvent)
	soc.HeadlessExperimental().OnMainFrameReadyForScreenshots(func(eventData *experimental.MainFrameReadyForScreenshotsEvent) {
		resultChan <- eventData
	})

	mockResult := &experimental.MainFrameReadyForScreenshotsEvent{}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "HeadlessExperimental.mainFrameReadyForScreenshots",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "HeadlessExperimental.mainFrameReadyForScreenshots",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestHeadlessExperimentalOnNeedsBeginFramesChanged(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *experimental.NeedsBeginFramesChangedEvent)
	soc.HeadlessExperimental().OnNeedsBeginFramesChanged(func(eventData *experimental.NeedsBeginFramesChangedEvent) {
		resultChan <- eventData
	})

	mockResult := &experimental.NeedsBeginFramesChangedEvent{
		NeedsBeginFrames: true,
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "HeadlessExperimental.needsBeginFramesChanged",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "HeadlessExperimental.needsBeginFramesChanged",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
