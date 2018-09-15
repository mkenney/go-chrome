package socket

import (
	"testing"

	"github.com/mkenney/go-chrome/tot/target"
)

func TestTargetActivateTarget(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &target.ActivateTargetParams{
		ID: target.ID("target-id"),
	}
	mockResult := &target.ActivateTargetResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Target().ActivateTarget(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Target().ActivateTarget(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTargetAttachToTarget(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &target.AttachToTargetParams{
		ID: target.ID("target-id"),
	}
	mockResult := &target.AttachToTargetResult{
		SessionID: target.SessionID("session-id"),
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Target().AttachToTarget(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.SessionID != result.SessionID {
		t.Errorf("Expected %s, got %s", mockResult.SessionID, result.SessionID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Target().AttachToTarget(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTargetCloseTarget(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &target.CloseTargetParams{
		ID: target.ID("target-id"),
	}
	mockResult := &target.CloseTargetResult{
		Success: true,
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Target().CloseTarget(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Success != result.Success {
		t.Errorf("Expected %v, got %v", mockResult.Success, result.Success)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Target().CloseTarget(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTargetCreateBrowserContext(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &target.CreateBrowserContextResult{
		BrowserContextID: target.BrowserContextID("BrowserContextID"),
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Target().CreateBrowserContext()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.BrowserContextID != result.BrowserContextID {
		t.Errorf("Expected %s, got %s", mockResult.BrowserContextID, result.BrowserContextID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Target().CreateBrowserContext()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTargetCreateTarget(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &target.CreateTargetParams{
		URL:                     "http://some.url",
		Width:                   1,
		Height:                  1,
		BrowserContextID:        target.BrowserContextID("BrowserContextID"),
		EnableBeginFrameControl: true,
	}
	mockResult := &target.CreateTargetResult{
		ID: target.ID("target-id"),
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Target().CreateTarget(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.ID != result.ID {
		t.Errorf("Expected %s, got %s", mockResult.ID, result.ID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Target().CreateTarget(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTargetDetachFromTarget(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &target.DetachFromTargetParams{
		SessionID: target.SessionID("SessionID"),
		ID:        target.ID("ID"),
	}
	mockResult := &target.DetachFromTargetResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Target().DetachFromTarget(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Target().DetachFromTarget(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTargetDisposeBrowserContext(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &target.DisposeBrowserContextParams{
		ID: target.ID("ID"),
	}
	mockResult := &target.DisposeBrowserContextResult{
		Success: true,
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Target().DisposeBrowserContext(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Success != result.Success {
		t.Errorf("Expected %v, got %v", mockResult.Success, result.Success)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Target().DisposeBrowserContext(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTargetGetTargetInfo(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &target.GetTargetInfoParams{
		ID: target.ID("ID"),
	}
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

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Target().GetTargetInfo(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if mockResult.Infos[0].ID != result.Infos[0].ID {
		t.Errorf("Expected %v, got %v", mockResult.Infos[0].ID, result.Infos[0].ID)
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Target().GetTargetInfo(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTargetGetTargets(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

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
	mockResult := &target.GetTargetsResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Target().GetTargets(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Target().GetTargets(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTargetSendMessageToTarget(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &target.SendMessageToTargetParams{
		Message:   "message",
		SessionID: target.SessionID("SessionID"),
		ID:        target.ID("ID"),
	}
	mockResult := &target.SendMessageToTargetResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Target().SendMessageToTarget(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Target().SendMessageToTarget(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTargetSetAttachToFrames(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &target.SetAttachToFramesParams{
		Value: true,
	}
	mockResult := &target.SetAttachToFramesResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Target().SetAttachToFrames(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Target().SetAttachToFrames(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTargetSetAutoAttach(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &target.SetAutoAttachParams{
		AutoAttach:             true,
		WaitForDebuggerOnStart: true,
	}
	mockResult := &target.SetAutoAttachResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Target().SetAutoAttach(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Target().SetAutoAttach(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTargetSetDiscoverTargets(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &target.SetDiscoverTargetsParams{
		Discover: true,
	}
	mockResult := &target.SetDiscoverTargetsResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Target().SetDiscoverTargets(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Target().SetDiscoverTargets(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTargetSetRemoteLocations(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &target.SetRemoteLocationsParams{
		Locations: []*target.RemoteLocation{{
			Host: "Host",
			Port: 1,
		}},
	}
	mockResult := &target.SetRemoteLocationsResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Target().SetRemoteLocations(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Target().SetRemoteLocations(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTargetOnAttachedToTarget(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *target.AttachedToTargetEvent)
	soc.Target().OnAttachedToTarget(func(eventData *target.AttachedToTargetEvent) {
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
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Target.attachedToTarget",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.SessionID != result.SessionID {
		t.Errorf("Expected %s, got %s", mockResult.SessionID, result.SessionID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Target.attachedToTarget",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTargetOnDetachedFromTarget(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *target.DetachedFromTargetEvent)
	soc.Target().OnDetachedFromTarget(func(eventData *target.DetachedFromTargetEvent) {
		resultChan <- eventData
	})

	mockResult := &target.DetachedFromTargetEvent{
		SessionID: target.SessionID("SessionID"),
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Target.detachedFromTarget",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.SessionID != result.SessionID {
		t.Errorf("Expected %s, got %s", mockResult.SessionID, result.SessionID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Target.detachedFromTarget",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTargetOnReceivedMessageFromTarget(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *target.ReceivedMessageFromTargetEvent)
	soc.Target().OnReceivedMessageFromTarget(func(eventData *target.ReceivedMessageFromTargetEvent) {
		resultChan <- eventData
	})

	mockResult := &target.ReceivedMessageFromTargetEvent{
		SessionID: target.SessionID("SessionID"),
		Message:   "message",
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Target.receivedMessageFromTarget",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.SessionID != result.SessionID {
		t.Errorf("Expected %s, got %s", mockResult.SessionID, result.SessionID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Target.receivedMessageFromTarget",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTargetOnTargetCreated(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *target.CreatedEvent)
	soc.Target().OnTargetCreated(func(eventData *target.CreatedEvent) {
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
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Target.targetCreated",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Info.ID != result.Info.ID {
		t.Errorf("Expected %s, got %s", mockResult.Info.ID, result.Info.ID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Target.targetCreated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTargetOnTargetDestroyed(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *target.DestroyedEvent)
	soc.Target().OnTargetDestroyed(func(eventData *target.DestroyedEvent) {
		resultChan <- eventData
	})

	mockResult := &target.DestroyedEvent{
		ID: target.ID("ID"),
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Target.targetDestroyed",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.ID != result.ID {
		t.Errorf("Expected %s, got %s", mockResult.ID, result.ID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Target.targetDestroyed",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestTargetOnTargetInfoChanged(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *target.InfoChangedEvent)
	soc.Target().OnTargetInfoChanged(func(eventData *target.InfoChangedEvent) {
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
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "Target.targetInfoChanged",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Info.ID != result.Info.ID {
		t.Errorf("Expected %s, got %s", mockResult.Info.ID, result.Info.ID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "Target.targetInfoChanged",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
