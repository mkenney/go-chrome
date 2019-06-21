package socket

import (
	"encoding/json"
	"net/url"
	"testing"

	"github.com/mkenney/go-chrome/tot/target"
)

func TestTargetActivateTarget(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestTargetActivateTarget")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &target.ActivateTargetParams{
		ID: target.ID("target-id"),
	}
	resultChan := mockSocket.Target().ActivateTarget(params)
	mockResult := &target.ActivateTargetResult{}
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

	resultChan = mockSocket.Target().ActivateTarget(params)
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

func TestTargetAttachToTarget(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestTargetAttachToTarget")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &target.AttachToTargetParams{
		ID: target.ID("target-id"),
	}
	resultChan := mockSocket.Target().AttachToTarget(params)
	mockResult := &target.AttachToTargetResult{
		SessionID: target.SessionID("session-id"),
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
	if mockResult.SessionID != result.SessionID {
		t.Errorf("Expected %s, got %s", mockResult.SessionID, result.SessionID)
	}

	resultChan = mockSocket.Target().AttachToTarget(params)
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

func TestTargetCloseTarget(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestTargetCloseTarget")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &target.CloseTargetParams{
		ID: target.ID("target-id"),
	}
	resultChan := mockSocket.Target().CloseTarget(params)
	mockResult := &target.CloseTargetResult{
		Success: true,
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
	if mockResult.Success != result.Success {
		t.Errorf("Expected %v, got %v", mockResult.Success, result.Success)
	}

	resultChan = mockSocket.Target().CloseTarget(params)
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

func TestTargetCreateBrowserContext(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestTargetCreateBrowserContext")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Target().CreateBrowserContext()
	mockResult := &target.CreateBrowserContextResult{
		BrowserContextID: target.BrowserContextID("BrowserContextID"),
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
	if mockResult.BrowserContextID != result.BrowserContextID {
		t.Errorf("Expected %s, got %s", mockResult.BrowserContextID, result.BrowserContextID)
	}

	resultChan = mockSocket.Target().CreateBrowserContext()
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

func TestTargetCreateTarget(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestTargetCreateTarget")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &target.CreateTargetParams{
		URL:                     "http://some.url",
		Width:                   1,
		Height:                  1,
		BrowserContextID:        target.BrowserContextID("BrowserContextID"),
		EnableBeginFrameControl: true,
	}
	resultChan := mockSocket.Target().CreateTarget(params)
	mockResult := &target.CreateTargetResult{
		ID: target.ID("target-id"),
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
	if mockResult.ID != result.ID {
		t.Errorf("Expected %s, got %s", mockResult.ID, result.ID)
	}

	resultChan = mockSocket.Target().CreateTarget(params)
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

func TestTargetDetachFromTarget(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestTargetDetachFromTarget")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &target.DetachFromTargetParams{
		SessionID: target.SessionID("SessionID"),
		ID:        target.ID("ID"),
	}
	resultChan := mockSocket.Target().DetachFromTarget(params)
	mockResult := &target.DetachFromTargetResult{}
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

	resultChan = mockSocket.Target().DetachFromTarget(params)
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

func TestTargetDisposeBrowserContext(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestTargetDisposeBrowserContext")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &target.DisposeBrowserContextParams{
		ID: target.ID("ID"),
	}
	resultChan := mockSocket.Target().DisposeBrowserContext(params)
	mockResult := &target.DisposeBrowserContextResult{
		Success: true,
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
	if mockResult.Success != result.Success {
		t.Errorf("Expected %v, got %v", mockResult.Success, result.Success)
	}

	resultChan = mockSocket.Target().DisposeBrowserContext(params)
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

func TestTargetGetTargetInfo(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestTargetGetTargetInfo")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &target.GetTargetInfoParams{
		ID: target.ID("ID"),
	}
	resultChan := mockSocket.Target().GetTargetInfo(params)
	mockResult := &target.GetTargetInfoResult{
		Infos: []*target.Info{{
			ID:       target.ID("ID"),
			Type:     "Type",
			Title:    "Title",
			URL:      "URL",
			Attached: true,
			OpenerID: target.ID("ID"),
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
	if mockResult.Infos[0].ID != result.Infos[0].ID {
		t.Errorf("Expected %v, got %v", mockResult.Infos[0].ID, result.Infos[0].ID)
	}

	resultChan = mockSocket.Target().GetTargetInfo(params)
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

func TestTargetGetTargets(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestTargetGetTargets")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &target.GetTargetsParams{
		Infos: []*target.Info{{
			ID:       target.ID("ID"),
			Type:     "Type",
			Title:    "Title",
			URL:      "URL",
			Attached: true,
			OpenerID: target.ID("ID"),
		}},
	}
	resultChan := mockSocket.Target().GetTargets(params)
	mockResult := &target.GetTargetsResult{}
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

	resultChan = mockSocket.Target().GetTargets(params)
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

func TestTargetSendMessageToTarget(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestTargetSendMessageToTarget")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &target.SendMessageToTargetParams{
		Message:   "message",
		SessionID: target.SessionID("SessionID"),
		ID:        target.ID("ID"),
	}
	resultChan := mockSocket.Target().SendMessageToTarget(params)
	mockResult := &target.SendMessageToTargetResult{}
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

	resultChan = mockSocket.Target().SendMessageToTarget(params)
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

func TestTargetSetAttachToFrames(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestTargetSetAttachToFrames")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &target.SetAttachToFramesParams{
		Value: true,
	}
	resultChan := mockSocket.Target().SetAttachToFrames(params)
	mockResult := &target.SetAttachToFramesResult{}
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

	resultChan = mockSocket.Target().SetAttachToFrames(params)
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

func TestTargetSetAutoAttach(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestTargetSetAutoAttach")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &target.SetAutoAttachParams{
		AutoAttach:             true,
		WaitForDebuggerOnStart: true,
	}
	resultChan := mockSocket.Target().SetAutoAttach(params)
	mockResult := &target.SetAutoAttachResult{}
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

	resultChan = mockSocket.Target().SetAutoAttach(params)
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

func TestTargetSetDiscoverTargets(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestTargetSetDiscoverTargets")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &target.SetDiscoverTargetsParams{
		Discover: true,
	}
	resultChan := mockSocket.Target().SetDiscoverTargets(params)
	mockResult := &target.SetDiscoverTargetsResult{}
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

	resultChan = mockSocket.Target().SetDiscoverTargets(params)
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

func TestTargetSetRemoteLocations(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestTargetSetRemoteLocations")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &target.SetRemoteLocationsParams{
		Locations: []*target.RemoteLocation{{
			Host: "Host",
			Port: 1,
		}},
	}
	resultChan := mockSocket.Target().SetRemoteLocations(params)
	mockResult := &target.SetRemoteLocationsResult{}
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

	resultChan = mockSocket.Target().SetRemoteLocations(params)
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

func TestTargetOnAttachedToTarget(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestTargetOnAttachedToTarget")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *target.AttachedToTargetEvent)
	mockSocket.Target().OnAttachedToTarget(func(eventData *target.AttachedToTargetEvent) {
		resultChan <- eventData
	})
	mockResult := &target.AttachedToTargetEvent{
		SessionID: target.SessionID("SessionID"),
		Info: &target.Info{
			ID:       target.ID("ID"),
			Type:     "Type",
			Title:    "Title",
			URL:      "URL",
			Attached: true,
			OpenerID: target.ID("ID"),
		},
		WaitingForDebugger: true,
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Target.attachedToTarget",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.SessionID != result.SessionID {
		t.Errorf("Expected %s, got %s", mockResult.SessionID, result.SessionID)
	}

	resultChan = make(chan *target.AttachedToTargetEvent)
	mockSocket.Target().OnAttachedToTarget(func(eventData *target.AttachedToTargetEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Target.attachedToTarget",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTargetOnDetachedFromTarget(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestTargetOnDetachedFromTarget")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *target.DetachedFromTargetEvent)
	mockSocket.Target().OnDetachedFromTarget(func(eventData *target.DetachedFromTargetEvent) {
		resultChan <- eventData
	})
	mockResult := &target.DetachedFromTargetEvent{
		SessionID: target.SessionID("SessionID"),
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Target.detachedFromTarget",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.SessionID != result.SessionID {
		t.Errorf("Expected %s, got %s", mockResult.SessionID, result.SessionID)
	}

	resultChan = make(chan *target.DetachedFromTargetEvent)
	mockSocket.Target().OnDetachedFromTarget(func(eventData *target.DetachedFromTargetEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Target.detachedFromTarget",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTargetOnReceivedMessageFromTarget(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestTargetOnReceivedMessageFromTarget")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *target.ReceivedMessageFromTargetEvent)
	mockSocket.Target().OnReceivedMessageFromTarget(func(eventData *target.ReceivedMessageFromTargetEvent) {
		resultChan <- eventData
	})
	mockResult := &target.ReceivedMessageFromTargetEvent{
		SessionID: target.SessionID("SessionID"),
		Message:   "message",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Target.receivedMessageFromTarget",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.SessionID != result.SessionID {
		t.Errorf("Expected %s, got %s", mockResult.SessionID, result.SessionID)
	}

	resultChan = make(chan *target.ReceivedMessageFromTargetEvent)
	mockSocket.Target().OnReceivedMessageFromTarget(func(eventData *target.ReceivedMessageFromTargetEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Target.receivedMessageFromTarget",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTargetOnTargetCreated(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestTargetOnTargetCreated")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *target.CreatedEvent)
	mockSocket.Target().OnTargetCreated(func(eventData *target.CreatedEvent) {
		resultChan <- eventData
	})
	mockResult := &target.CreatedEvent{
		Info: &target.Info{
			ID:       target.ID("ID"),
			Type:     "Type",
			Title:    "Title",
			URL:      "URL",
			Attached: true,
			OpenerID: target.ID("ID"),
		},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Target.targetCreated",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Info.ID != result.Info.ID {
		t.Errorf("Expected %s, got %s", mockResult.Info.ID, result.Info.ID)
	}

	resultChan = make(chan *target.CreatedEvent)
	mockSocket.Target().OnTargetCreated(func(eventData *target.CreatedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Target.targetCreated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTargetOnTargetDestroyed(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestTargetOnTargetDestroyed")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *target.DestroyedEvent)
	mockSocket.Target().OnTargetDestroyed(func(eventData *target.DestroyedEvent) {
		resultChan <- eventData
	})
	mockResult := &target.DestroyedEvent{
		ID: target.ID("ID"),
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Target.targetDestroyed",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.ID != result.ID {
		t.Errorf("Expected %s, got %s", mockResult.ID, result.ID)
	}

	resultChan = make(chan *target.DestroyedEvent)
	mockSocket.Target().OnTargetDestroyed(func(eventData *target.DestroyedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Target.targetDestroyed",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTargetOnTargetInfoChanged(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestTargetOnTargetInfoChanged")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *target.InfoChangedEvent)
	mockSocket.Target().OnTargetInfoChanged(func(eventData *target.InfoChangedEvent) {
		resultChan <- eventData
	})
	mockResult := &target.InfoChangedEvent{
		Info: &target.Info{
			ID:       target.ID("ID"),
			Type:     "Type",
			Title:    "Title",
			URL:      "URL",
			Attached: true,
			OpenerID: target.ID("ID"),
		},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Target.targetInfoChanged",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Info.ID != result.Info.ID {
		t.Errorf("Expected %s, got %s", mockResult.Info.ID, result.Info.ID)
	}

	resultChan = make(chan *target.InfoChangedEvent)
	mockSocket.Target().OnTargetInfoChanged(func(eventData *target.InfoChangedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Target.targetInfoChanged",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
