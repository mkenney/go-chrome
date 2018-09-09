package socket

import (
	"testing"
	"time"

	"github.com/mkenney/go-chrome/tot/input"
)

func TestInputDispatchKeyEvent(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &input.DispatchKeyEventParams{
		Type:                  input.KeyEvent.KeyDown,
		Modifiers:             1,
		Timestamp:             input.TimeSinceEpoch(time.Now().Unix()),
		Text:                  "text",
		UnmodifiedText:        "unmodified text",
		KeyIdentifier:         "key-id",
		Code:                  "code",
		Key:                   "key",
		WindowsVirtualKeyCode: 1,
		NativeVirtualKeyCode:  1,
		AutoRepeat:            true,
		IsKeypad:              true,
		IsSystemKey:           true,
		Location:              1,
	}
	mockResult := &input.DispatchKeyEventResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Input().DispatchKeyEvent(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Input().DispatchKeyEvent(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestInputDispatchMouseEvent(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &input.DispatchMouseEventParams{
		Type:       input.MouseEvent.MousePressed,
		X:          1,
		Y:          1,
		Modifiers:  1,
		Timestamp:  input.TimeSinceEpoch(time.Now().Unix()),
		Button:     input.ButtonEvent.None,
		ClickCount: 1,
		DeltaX:     1,
		DeltaY:     1,
	}
	mockResult := &input.DispatchMouseEventResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Input().DispatchMouseEvent(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Input().DispatchMouseEvent(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestInputDispatchTouchEvent(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &input.DispatchTouchEventParams{
		Type: input.TouchEvent.TouchStart,
		TouchPoints: []*input.TouchPoint{{
			X:             1,
			Y:             1,
			RadiusX:       1,
			RadiusY:       1,
			RotationAngle: 1,
			Force:         1,
			ID:            1,
		}},
		Modifiers: 1,
		Timestamp: input.TimeSinceEpoch(time.Now().Unix()),
	}
	mockResult := &input.DispatchTouchEventResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Input().DispatchTouchEvent(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Input().DispatchTouchEvent(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestInputEmulateTouchFromMouseEvent(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &input.EmulateTouchFromMouseEventParams{
		Type:       input.MouseEvent.MousePressed,
		X:          1,
		Y:          1,
		Timestamp:  input.TimeSinceEpoch(time.Now().Unix()),
		Button:     input.ButtonEvent.None,
		DeltaX:     1,
		DeltaY:     1,
		Modifiers:  1,
		ClickCount: 1,
	}
	mockResult := &input.EmulateTouchFromMouseEventResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Input().EmulateTouchFromMouseEvent(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Input().EmulateTouchFromMouseEvent(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestInputSetIgnoreEvents(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &input.SetIgnoreEventsParams{
		Ignore: true,
	}
	mockResult := &input.SetIgnoreEventsResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Input().SetIgnoreEvents(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Input().SetIgnoreEvents(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestInputSynthesizePinchGesture(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &input.SynthesizePinchGestureParams{
		X:                 1,
		Y:                 1,
		ScaleFactor:       1,
		RelativeSpeed:     1,
		GestureSourceType: input.GestureSourceType("gesture-source-type"),
	}
	mockResult := &input.SynthesizePinchGestureResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Input().SynthesizePinchGesture(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Input().SynthesizePinchGesture(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestInputSynthesizeScrollGesture(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &input.SynthesizeScrollGestureParams{
		X:                     1,
		Y:                     1,
		XDistance:             1,
		YDistance:             1,
		XOverscroll:           1,
		YOverscroll:           1,
		PreventFling:          true,
		Speed:                 1,
		GestureSourceType:     input.GestureSourceType("gesture-source-type"),
		RepeatCount:           1,
		RepeatDelayMs:         1,
		InteractionMarkerName: "marker-name",
	}
	mockResult := &input.SynthesizeScrollGestureResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Input().SynthesizeScrollGesture(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Input().SynthesizeScrollGesture(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestInputSynthesizeTapGesture(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &input.SynthesizeTapGestureParams{
		X:                 1,
		Y:                 1,
		Duration:          1,
		TapCount:          1,
		GestureSourceType: input.GestureSourceType("gesture-source-type"),
	}
	mockResult := &input.SynthesizeTapGestureResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Input().SynthesizeTapGesture(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Input().SynthesizeTapGesture(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
