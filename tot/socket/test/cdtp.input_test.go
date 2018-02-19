package test

import (
	"encoding/json"
	"net/url"
	"testing"
	"time"

	input "github.com/mkenney/go-chrome/tot/cdtp/input"
	"github.com/mkenney/go-chrome/tot/socket"
)

func TestInputDispatchKeyEvent(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &input.DispatchKeyEventParams{
		Type:           input.KeyEvent.KeyDown,
		Modifiers:      1,
		Timestamp:      input.TimeSinceEpoch(time.Now().Unix()),
		Text:           "text",
		UnmodifiedText: "unmodified text",
		KeyIdentifier:  "key-id",
		Code:           "code",
		Key:            "key",
		WindowsVirtualKeyCode: 1,
		NativeVirtualKeyCode:  1,
		AutoRepeat:            true,
		IsKeypad:              true,
		IsSystemKey:           true,
		Location:              1,
	}
	resultChan := mockSocket.Input().DispatchKeyEvent(params)
	mockResult := &input.DispatchKeyEventResult{}
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

	resultChan = mockSocket.Input().DispatchKeyEvent(params)
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

func TestInputDispatchMouseEvent(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

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
	resultChan := mockSocket.Input().DispatchMouseEvent(params)
	mockResult := &input.DispatchMouseEventResult{}
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

	resultChan = mockSocket.Input().DispatchMouseEvent(params)
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

func TestInputDispatchTouchEvent(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

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
	resultChan := mockSocket.Input().DispatchTouchEvent(params)
	mockResult := &input.DispatchTouchEventResult{}
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

	resultChan = mockSocket.Input().DispatchTouchEvent(params)
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

func TestInputEmulateTouchFromMouseEvent(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

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
	resultChan := mockSocket.Input().EmulateTouchFromMouseEvent(params)
	mockResult := &input.EmulateTouchFromMouseEventResult{}
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

	resultChan = mockSocket.Input().EmulateTouchFromMouseEvent(params)
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

func TestInputSetIgnoreEvents(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &input.SetIgnoreEventsParams{
		Ignore: true,
	}
	resultChan := mockSocket.Input().SetIgnoreEvents(params)
	mockResult := &input.SetIgnoreEventsResult{}
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

	resultChan = mockSocket.Input().SetIgnoreEvents(params)
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

func TestInputSynthesizePinchGesture(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &input.SynthesizePinchGestureParams{
		X:                 1,
		Y:                 1,
		ScaleFactor:       1,
		RelativeSpeed:     1,
		GestureSourceType: input.GestureSourceType("gesture-source-type"),
	}
	resultChan := mockSocket.Input().SynthesizePinchGesture(params)
	mockResult := &input.SynthesizePinchGestureResult{}
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

	resultChan = mockSocket.Input().SynthesizePinchGesture(params)
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

func TestInputSynthesizeScrollGesture(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

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
	resultChan := mockSocket.Input().SynthesizeScrollGesture(params)
	mockResult := &input.SynthesizeScrollGestureResult{}
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

	resultChan = mockSocket.Input().SynthesizeScrollGesture(params)
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

func TestInputSynthesizeTapGesture(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := socket.NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &input.SynthesizeTapGestureParams{
		X:                 1,
		Y:                 1,
		Duration:          1,
		TapCount:          1,
		GestureSourceType: input.GestureSourceType("gesture-source-type"),
	}
	resultChan := mockSocket.Input().SynthesizeTapGesture(params)
	mockResult := &input.SynthesizeTapGestureResult{}
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

	resultChan = mockSocket.Input().SynthesizeTapGesture(params)
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
