package test

import (
	"encoding/json"
	"net/url"
	"testing"
	"time"

	dom "github.com/mkenney/go-chrome/tot/cdtp/dom"
	emulation "github.com/mkenney/go-chrome/tot/cdtp/emulation"
	page "github.com/mkenney/go-chrome/tot/cdtp/page"
	runtime "github.com/mkenney/go-chrome/tot/cdtp/runtime"
	"github.com/mkenney/go-chrome/tot/socket"
)

func TestEmulationCanEmulate(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Emulation().CanEmulate()
	mockResult := &emulation.CanEmulateResult{
		Result: true,
	}
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
	if mockResult.Result != result.Result {
		t.Errorf("Expected %v, got %v", mockResult.Result, result.Result)
	}

	resultChan = mockSocket.Emulation().CanEmulate()
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

func TestEmulationClearDeviceMetricsOverride(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Emulation().ClearDeviceMetricsOverride()
	mockResult := &emulation.ClearDeviceMetricsOverrideResult{}
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

	resultChan = mockSocket.Emulation().ClearDeviceMetricsOverride()
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

func TestEmulationClearGeolocationOverride(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Emulation().ClearGeolocationOverride()
	mockResult := &emulation.ClearGeolocationOverrideResult{}
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

	resultChan = mockSocket.Emulation().ClearGeolocationOverride()
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

func TestEmulationResetPageScaleFactor(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Emulation().ResetPageScaleFactor()
	mockResult := &emulation.ResetPageScaleFactorResult{}
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

	resultChan = mockSocket.Emulation().ResetPageScaleFactor()
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

func TestEmulationSetCPUThrottlingRate(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &emulation.SetCPUThrottlingRateParams{
		Rate: 1,
	}
	resultChan := mockSocket.Emulation().SetCPUThrottlingRate(params)
	mockResult := &emulation.SetCPUThrottlingRateResult{}
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

	resultChan = mockSocket.Emulation().SetCPUThrottlingRate(params)
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

func TestEmulationSetDefaultBackgroundColorOverride(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &emulation.SetDefaultBackgroundColorOverrideParams{
		Color: &dom.RGBA{
			R: 1,
			G: 1,
			B: 1,
			A: 1,
		},
	}
	resultChan := mockSocket.Emulation().SetDefaultBackgroundColorOverride(params)
	mockResult := &emulation.SetDefaultBackgroundColorOverrideResult{}
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

	resultChan = mockSocket.Emulation().SetDefaultBackgroundColorOverride(params)
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

func TestEmulationSetDeviceMetricsOverride(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &emulation.SetDeviceMetricsOverrideParams{
		Width:              1,
		Height:             1,
		DeviceScaleFactor:  1,
		Mobile:             true,
		Scale:              1,
		ScreenWidth:        1,
		ScreenHeight:       1,
		PositionX:          1,
		PositionY:          1,
		DontSetVisibleSize: true,
		ScreenOrientation: &emulation.ScreenOrientation{
			Type:  emulation.OrientationType.PortraitPrimary,
			Angle: 1,
		},
		Viewport: &page.Viewport{
			X:      1,
			Y:      1,
			Width:  1,
			Height: 1,
			Scale:  1,
		},
	}
	resultChan := mockSocket.Emulation().SetDeviceMetricsOverride(params)
	mockResult := &emulation.SetDeviceMetricsOverrideResult{}
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

	resultChan = mockSocket.Emulation().SetDeviceMetricsOverride(params)
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

func TestEmulationSetEmitTouchEventsForMouse(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &emulation.SetEmitTouchEventsForMouseParams{
		Enabled:       true,
		Configuration: emulation.Configuration.Mobile,
	}
	resultChan := mockSocket.Emulation().SetEmitTouchEventsForMouse(params)
	mockResult := &emulation.SetEmitTouchEventsForMouseResult{}
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

	resultChan = mockSocket.Emulation().SetEmitTouchEventsForMouse(params)
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

func TestEmulationSetEmulatedMedia(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &emulation.SetEmulatedMediaParams{
		Media: "media",
	}
	resultChan := mockSocket.Emulation().SetEmulatedMedia(params)
	mockResult := &emulation.SetEmulatedMediaResult{}
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

	resultChan = mockSocket.Emulation().SetEmulatedMedia(params)
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

func TestEmulationSetGeolocationOverride(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &emulation.SetGeolocationOverrideParams{
		Latitude:  1,
		Longitude: 1,
		Accuracy:  1,
	}
	resultChan := mockSocket.Emulation().SetGeolocationOverride(params)
	mockResult := &emulation.SetGeolocationOverrideResult{}
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

	resultChan = mockSocket.Emulation().SetGeolocationOverride(params)
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

func TestEmulationSetNavigatorOverrides(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &emulation.SetNavigatorOverridesParams{
		Platform: "platform",
	}
	resultChan := mockSocket.Emulation().SetNavigatorOverrides(params)
	mockResult := &emulation.SetNavigatorOverridesResult{}
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

	resultChan = mockSocket.Emulation().SetNavigatorOverrides(params)
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

func TestEmulationSetPageScaleFactor(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &emulation.SetPageScaleFactorParams{
		PageScaleFactor: 1,
	}
	resultChan := mockSocket.Emulation().SetPageScaleFactor(params)
	mockResult := &emulation.SetPageScaleFactorResult{}
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

	resultChan = mockSocket.Emulation().SetPageScaleFactor(params)
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

func TestEmulationSetScriptExecutionDisabled(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &emulation.SetScriptExecutionDisabledParams{
		Value: true,
	}
	resultChan := mockSocket.Emulation().SetScriptExecutionDisabled(params)
	mockResult := &emulation.SetScriptExecutionDisabledResult{}
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

	resultChan = mockSocket.Emulation().SetScriptExecutionDisabled(params)
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

func TestEmulationSetTouchEmulationEnabled(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &emulation.SetTouchEmulationEnabledParams{
		Enabled:        true,
		MaxTouchPoints: 1,
	}
	resultChan := mockSocket.Emulation().SetTouchEmulationEnabled(params)
	mockResult := &emulation.SetTouchEmulationEnabledResult{}
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

	resultChan = mockSocket.Emulation().SetTouchEmulationEnabled(params)
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

func TestEmulationSetVirtualTimePolicy(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &emulation.SetVirtualTimePolicyParams{
		Policy: emulation.VirtualTimePolicy("policy"),
		Budget: 1,
		MaxVirtualTimeTaskStarvationCount: 1,
	}
	resultChan := mockSocket.Emulation().SetVirtualTimePolicy(params)
	mockResult := &emulation.SetVirtualTimePolicyResult{
		VirtualTimeBase: runtime.Timestamp(time.Now().Unix()),
	}
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
	if mockResult.VirtualTimeBase != result.VirtualTimeBase {
		t.Errorf("Expected %v, got %v", mockResult.VirtualTimeBase, result.VirtualTimeBase)
	}

	resultChan = mockSocket.Emulation().SetVirtualTimePolicy(params)
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

func TestEmulationSetVisibleSize(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &emulation.SetVisibleSizeParams{
		Width:  1,
		Height: 1,
	}
	resultChan := mockSocket.Emulation().SetVisibleSize(params)
	mockResult := &emulation.SetVisibleSizeResult{}
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

	resultChan = mockSocket.Emulation().SetVisibleSize(params)
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

func TestEmulationOnVirtualTimeAdvanced(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *emulation.VirtualTimeAdvancedEvent)
	mockSocket.Emulation().OnVirtualTimeAdvanced(func(eventData *emulation.VirtualTimeAdvancedEvent) {
		resultChan <- eventData
	})
	mockResult := &emulation.VirtualTimeAdvancedEvent{
		VirtualTimeElapsed: 1,
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     0,
		Error:  &socket.Error{},
		Method: "Emulation.virtualTimeAdvanced",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *emulation.VirtualTimeAdvancedEvent)
	mockSocket.Emulation().OnVirtualTimeAdvanced(func(eventData *emulation.VirtualTimeAdvancedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&socket.Response{
		ID: 0,
		Error: &socket.Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Emulation.virtualTimeAdvanced",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestEmulationOnVirtualTimeBudgetExpired(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *emulation.VirtualTimeBudgetExpiredEvent)
	mockSocket.Emulation().OnVirtualTimeBudgetExpired(func(eventData *emulation.VirtualTimeBudgetExpiredEvent) {
		resultChan <- eventData
	})
	mockResult := &emulation.VirtualTimeBudgetExpiredEvent{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     0,
		Error:  &socket.Error{},
		Method: "Emulation.virtualTimeBudgetExpired",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *emulation.VirtualTimeBudgetExpiredEvent)
	mockSocket.Emulation().OnVirtualTimeBudgetExpired(func(eventData *emulation.VirtualTimeBudgetExpiredEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&socket.Response{
		ID: 0,
		Error: &socket.Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Emulation.virtualTimeBudgetExpired",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestEmulationOnVirtualTimePaused(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *emulation.VirtualTimePausedEvent)
	mockSocket.Emulation().OnVirtualTimePaused(func(eventData *emulation.VirtualTimePausedEvent) {
		resultChan <- eventData
	})
	mockResult := &emulation.VirtualTimePausedEvent{
		VirtualTimeElapsed: 1,
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&socket.Response{
		ID:     0,
		Error:  &socket.Error{},
		Method: "Emulation.virtualTimePaused",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	resultChan = make(chan *emulation.VirtualTimePausedEvent)
	mockSocket.Emulation().OnVirtualTimePaused(func(eventData *emulation.VirtualTimePausedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&socket.Response{
		ID: 0,
		Error: &socket.Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Emulation.virtualTimePaused",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
