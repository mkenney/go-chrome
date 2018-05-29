package socket

import (
	"encoding/json"
	"net/url"
	"testing"
	"time"

	headlessExperimental "github.com/mkenney/go-chrome/tot/cdtp/headless/experimental"
	runtime "github.com/mkenney/go-chrome/tot/cdtp/runtime"
)

func TestHeadlessExperimentalBeginFrame(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestHeadlessExperimentalBeginFrame")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	params := &headlessExperimental.BeginFrameParams{
		FrameTime: runtime.Timestamp(time.Now().Unix()),
		Deadline:  runtime.Timestamp(time.Now().Unix()),
		Interval:  1.1,
		Screenshot: &headlessExperimental.ScreenshotParams{
			Format:  headlessExperimental.Format.Jpeg,
			Quality: 100,
		},
	}
	resultChan := mockSocket.HeadlessExperimental().BeginFrame(params)
	mockResult := &headlessExperimental.BeginFrameResult{
		HasDamage:               true,
		MainFrameContentUpdated: true,
		ScreenshotData:          "data",
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
	if mockResult.HasDamage != result.HasDamage {
		t.Errorf("Expected %v, got %v", mockResult.HasDamage, result.HasDamage)
	}

	resultChan = mockSocket.HeadlessExperimental().BeginFrame(params)
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

func TestHeadlessExperimentalDisable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestHeadlessExperimentalDisable")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.HeadlessExperimental().Disable()
	mockResult := &headlessExperimental.DisableResult{}
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

	resultChan = mockSocket.HeadlessExperimental().Disable()
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

func TestHeadlessExperimentalEnable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestHeadlessExperimentalEnable")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.HeadlessExperimental().Enable()
	mockResult := &headlessExperimental.EnableResult{}
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

	resultChan = mockSocket.HeadlessExperimental().Enable()
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

func TestHeadlessExperimentalOnMainFrameReadyForScreenshots(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestHeadlessExperimentalOnMainFrameReadyForScreenshots")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *headlessExperimental.MainFrameReadyForScreenshotsEvent)
	mockSocket.HeadlessExperimental().OnMainFrameReadyForScreenshots(func(eventData *headlessExperimental.MainFrameReadyForScreenshotsEvent) {
		resultChan <- eventData
	})
	mockResult := &headlessExperimental.MainFrameReadyForScreenshotsEvent{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "HeadlessExperimental.mainFrameReadyForScreenshots",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *headlessExperimental.MainFrameReadyForScreenshotsEvent)
	mockSocket.HeadlessExperimental().OnMainFrameReadyForScreenshots(func(eventData *headlessExperimental.MainFrameReadyForScreenshotsEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "HeadlessExperimental.mainFrameReadyForScreenshots",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestHeadlessExperimentalOnNeedsBeginFramesChanged(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestHeadlessExperimentalOnNeedsBeginFramesChanged")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *headlessExperimental.NeedsBeginFramesChangedEvent)
	mockSocket.HeadlessExperimental().OnNeedsBeginFramesChanged(func(eventData *headlessExperimental.NeedsBeginFramesChangedEvent) {
		resultChan <- eventData
	})
	mockResult := &headlessExperimental.NeedsBeginFramesChangedEvent{
		NeedsBeginFrames: true,
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "HeadlessExperimental.needsBeginFramesChanged",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *headlessExperimental.NeedsBeginFramesChangedEvent)
	mockSocket.HeadlessExperimental().OnNeedsBeginFramesChanged(func(eventData *headlessExperimental.NeedsBeginFramesChangedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "HeadlessExperimental.needsBeginFramesChanged",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
