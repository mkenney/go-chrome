package socket

import (
	"testing"
	"time"

	"github.com/mkenney/go-chrome/tot/dom"
	"github.com/mkenney/go-chrome/tot/emulation"
	"github.com/mkenney/go-chrome/tot/page"
	"github.com/mkenney/go-chrome/tot/runtime"
)

func TestEmulationCanEmulate(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &emulation.CanEmulateResult{
		Result: true,
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Emulation().CanEmulate()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Result != result.Result {
		t.Errorf("Expected %v, got %v", mockResult.Result, result.Result)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Emulation().CanEmulate()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestEmulationClearDeviceMetricsOverride(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &emulation.ClearDeviceMetricsOverrideResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Emulation().ClearDeviceMetricsOverride()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Emulation().ClearDeviceMetricsOverride()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestEmulationClearGeolocationOverride(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &emulation.ClearGeolocationOverrideResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Emulation().ClearGeolocationOverride()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Emulation().ClearGeolocationOverride()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestEmulationResetPageScaleFactor(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &emulation.ResetPageScaleFactorResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Emulation().ResetPageScaleFactor()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Emulation().ResetPageScaleFactor()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestEmulationSetCPUThrottlingRate(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &emulation.SetCPUThrottlingRateParams{
		Rate: 1,
	}
	mockResult := &emulation.SetCPUThrottlingRateResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Emulation().SetCPUThrottlingRate(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Emulation().SetCPUThrottlingRate(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestEmulationSetDefaultBackgroundColorOverride(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &emulation.SetDefaultBackgroundColorOverrideParams{
		Color: &dom.RGBA{
			R: 1,
			G: 1,
			B: 1,
			A: 1,
		},
	}
	mockResult := &emulation.SetDefaultBackgroundColorOverrideResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Emulation().SetDefaultBackgroundColorOverride(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Emulation().SetDefaultBackgroundColorOverride(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestEmulationSetDeviceMetricsOverride(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

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
	mockResult := &emulation.SetDeviceMetricsOverrideResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Emulation().SetDeviceMetricsOverride(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Emulation().SetDeviceMetricsOverride(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestEmulationSetEmitTouchEventsForMouse(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &emulation.SetEmitTouchEventsForMouseParams{
		Enabled:       true,
		Configuration: emulation.Configuration.Mobile,
	}
	mockResult := &emulation.SetEmitTouchEventsForMouseResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Emulation().SetEmitTouchEventsForMouse(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Emulation().SetEmitTouchEventsForMouse(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestEmulationSetEmulatedMedia(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &emulation.SetEmulatedMediaParams{
		Media: "media",
	}
	mockResult := &emulation.SetEmulatedMediaResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Emulation().SetEmulatedMedia(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Emulation().SetEmulatedMedia(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestEmulationSetGeolocationOverride(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &emulation.SetGeolocationOverrideParams{
		Latitude:  1,
		Longitude: 1,
		Accuracy:  1,
	}
	mockResult := &emulation.SetGeolocationOverrideResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Emulation().SetGeolocationOverride(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Emulation().SetGeolocationOverride(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestEmulationSetNavigatorOverrides(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &emulation.SetNavigatorOverridesParams{
		Platform: "platform",
	}
	mockResult := &emulation.SetNavigatorOverridesResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Emulation().SetNavigatorOverrides(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Emulation().SetNavigatorOverrides(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestEmulationSetPageScaleFactor(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &emulation.SetPageScaleFactorParams{
		PageScaleFactor: 1,
	}
	mockResult := &emulation.SetPageScaleFactorResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Emulation().SetPageScaleFactor(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Emulation().SetPageScaleFactor(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestEmulationSetScriptExecutionDisabled(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &emulation.SetScriptExecutionDisabledParams{
		Value: true,
	}
	mockResult := &emulation.SetScriptExecutionDisabledResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Emulation().SetScriptExecutionDisabled(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Emulation().SetScriptExecutionDisabled(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestEmulationSetTouchEmulationEnabled(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &emulation.SetTouchEmulationEnabledParams{
		Enabled:        true,
		MaxTouchPoints: 1,
	}
	mockResult := &emulation.SetTouchEmulationEnabledResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Emulation().SetTouchEmulationEnabled(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Emulation().SetTouchEmulationEnabled(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestEmulationSetVirtualTimePolicy(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &emulation.SetVirtualTimePolicyParams{
		Policy:                            emulation.VirtualTimePolicy("policy"),
		Budget:                            1,
		MaxVirtualTimeTaskStarvationCount: 1,
	}
	mockResult := &emulation.SetVirtualTimePolicyResult{
		VirtualTimeBase: runtime.Timestamp(time.Now().Unix()),
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Emulation().SetVirtualTimePolicy(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.VirtualTimeBase != result.VirtualTimeBase {
		t.Errorf("Expected %v, got %v", mockResult.VirtualTimeBase, result.VirtualTimeBase)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Emulation().SetVirtualTimePolicy(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestEmulationSetVisibleSize(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &emulation.SetVisibleSizeParams{
		Width:  1,
		Height: 1,
	}
	mockResult := &emulation.SetVisibleSizeResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Emulation().SetVisibleSize(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Emulation().SetVisibleSize(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestEmulationOnVirtualTimeAdvanced(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *emulation.VirtualTimeAdvancedEvent)
	soc.Emulation().OnVirtualTimeAdvanced(func(eventData *emulation.VirtualTimeAdvancedEvent) {
		resultChan <- eventData
	})

	mockResult := &emulation.VirtualTimeAdvancedEvent{
		VirtualTimeElapsed: 1,
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Emulation.virtualTimeAdvanced",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Emulation.virtualTimeAdvanced",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestEmulationOnVirtualTimeBudgetExpired(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *emulation.VirtualTimeBudgetExpiredEvent)
	soc.Emulation().OnVirtualTimeBudgetExpired(func(eventData *emulation.VirtualTimeBudgetExpiredEvent) {
		resultChan <- eventData
	})

	mockResult := &emulation.VirtualTimeBudgetExpiredEvent{}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Emulation.virtualTimeBudgetExpired",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Emulation.virtualTimeBudgetExpired",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestEmulationOnVirtualTimePaused(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *emulation.VirtualTimePausedEvent)
	soc.Emulation().OnVirtualTimePaused(func(eventData *emulation.VirtualTimePausedEvent) {
		resultChan <- eventData
	})

	mockResult := &emulation.VirtualTimePausedEvent{
		VirtualTimeElapsed: 1,
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Emulation.virtualTimePaused",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Emulation.virtualTimePaused",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
