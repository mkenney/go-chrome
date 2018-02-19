package socket

import (
	"encoding/json"
	"net/url"
	"testing"
	"time"

	serviceWorker "github.com/mkenney/go-chrome/tot/cdtp/service/worker"
	target "github.com/mkenney/go-chrome/tot/cdtp/target"
)

func TestServiceWorkerDeliverPushMessage(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &serviceWorker.DeliverPushMessageParams{
		Origin:         "origin",
		RegistrationID: "registration-id",
		Data:           "data",
	}
	resultChan := mockSocket.ServiceWorker().DeliverPushMessage(params)
	mockResult := &serviceWorker.DeliverPushMessageResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.ServiceWorker().DeliverPushMessage(params)
	mockSocket.Conn().AddMockData(&Response{
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

func TestServiceWorkerDisable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.ServiceWorker().Disable()
	mockResult := &serviceWorker.DisableResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.ServiceWorker().Disable()
	mockSocket.Conn().AddMockData(&Response{
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

func TestServiceWorkerDispatchSyncEvent(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &serviceWorker.DispatchSyncEventParams{
		Origin:         "origin",
		RegistrationID: "registration-id",
		Tag:            "tag",
		LastChance:     true,
	}
	resultChan := mockSocket.ServiceWorker().DispatchSyncEvent(params)
	mockResult := &serviceWorker.DispatchSyncEventResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.ServiceWorker().DispatchSyncEvent(params)
	mockSocket.Conn().AddMockData(&Response{
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

func TestServiceWorkerEnable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.ServiceWorker().Enable()
	mockResult := &serviceWorker.EnableResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.ServiceWorker().Enable()
	mockSocket.Conn().AddMockData(&Response{
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

func TestServiceWorkerInspectWorker(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &serviceWorker.InspectWorkerParams{
		VersionID: "version-id",
	}
	resultChan := mockSocket.ServiceWorker().InspectWorker(params)
	mockResult := &serviceWorker.InspectWorkerResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.ServiceWorker().InspectWorker(params)
	mockSocket.Conn().AddMockData(&Response{
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

func TestServiceWorkerSetForceUpdateOnPageLoad(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &serviceWorker.SetForceUpdateOnPageLoadParams{
		ForceUpdateOnPageLoad: true,
	}
	resultChan := mockSocket.ServiceWorker().SetForceUpdateOnPageLoad(params)
	mockResult := &serviceWorker.SetForceUpdateOnPageLoadResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.ServiceWorker().SetForceUpdateOnPageLoad(params)
	mockSocket.Conn().AddMockData(&Response{
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

func TestServiceWorkerSkipWaiting(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &serviceWorker.SkipWaitingParams{
		ScopeURL: "http://some.url",
	}
	resultChan := mockSocket.ServiceWorker().SkipWaiting(params)
	mockResult := &serviceWorker.SkipWaitingResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.ServiceWorker().SkipWaiting(params)
	mockSocket.Conn().AddMockData(&Response{
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

func TestServiceWorkerStartWorker(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &serviceWorker.StartWorkerParams{
		ScopeURL: "http://some.url",
	}
	resultChan := mockSocket.ServiceWorker().StartWorker(params)
	mockResult := &serviceWorker.StartWorkerResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.ServiceWorker().StartWorker(params)
	mockSocket.Conn().AddMockData(&Response{
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

func TestServiceWorkerStopAllWorkers(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.ServiceWorker().StopAllWorkers()
	mockResult := &serviceWorker.StopAllWorkersResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.ServiceWorker().StopAllWorkers()
	mockSocket.Conn().AddMockData(&Response{
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

func TestServiceWorkerStopWorker(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &serviceWorker.StopWorkerParams{
		VersionID: "version-id",
	}
	resultChan := mockSocket.ServiceWorker().StopWorker(params)
	mockResult := &serviceWorker.StopWorkerResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.ServiceWorker().StopWorker(params)
	mockSocket.Conn().AddMockData(&Response{
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

func TestServiceWorkerUnregister(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &serviceWorker.UnregisterParams{
		ScopeURL: "http://some.url",
	}
	resultChan := mockSocket.ServiceWorker().Unregister(params)
	mockResult := &serviceWorker.UnregisterResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.ServiceWorker().Unregister(params)
	mockSocket.Conn().AddMockData(&Response{
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

func TestServiceWorkerUpdateRegistration(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	params := &serviceWorker.UpdateRegistrationParams{
		ScopeURL: "http://some.url",
	}
	resultChan := mockSocket.ServiceWorker().UpdateRegistration(params)
	mockResult := &serviceWorker.UpdateRegistrationResult{}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.ServiceWorker().UpdateRegistration(params)
	mockSocket.Conn().AddMockData(&Response{
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

func TestServiceWorkerOnWorkerErrorReported(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *serviceWorker.ErrorReportedEvent)
	mockSocket.ServiceWorker().OnWorkerErrorReported(func(eventData *serviceWorker.ErrorReportedEvent) {
		resultChan <- eventData
	})
	mockResult := &serviceWorker.ErrorReportedEvent{
		ErrorMessage: &serviceWorker.ErrorMessage{
			ErrorMessage:   "error message",
			RegistrationID: "registration-id",
			VersionID:      "version-id",
			SourceURL:      "http://some.url",
			LineNumber:     1,
			ColumnNumber:   1,
		},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "ServiceWorker.workerErrorReported",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.ErrorMessage.ErrorMessage != result.ErrorMessage.ErrorMessage {
		t.Errorf("Expected %s, got %s", mockResult.ErrorMessage.ErrorMessage, result.ErrorMessage.ErrorMessage)
	}

	resultChan = make(chan *serviceWorker.ErrorReportedEvent)
	mockSocket.ServiceWorker().OnWorkerErrorReported(func(eventData *serviceWorker.ErrorReportedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "ServiceWorker.workerErrorReported",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestServiceWorkerOnWorkerRegistrationUpdated(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *serviceWorker.RegistrationUpdatedEvent)
	mockSocket.ServiceWorker().OnWorkerRegistrationUpdated(func(eventData *serviceWorker.RegistrationUpdatedEvent) {
		resultChan <- eventData
	})
	mockResult := &serviceWorker.RegistrationUpdatedEvent{
		Registrations: []*serviceWorker.Registration{{
			RegistrationID: "registration-id",
			ScopeURL:       "http://some.url",
			IsDeleted:      true,
		}},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "ServiceWorker.workerRegistrationUpdated",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Registrations[0].RegistrationID != result.Registrations[0].RegistrationID {
		t.Errorf("Expected %s, got %s", mockResult.Registrations[0].RegistrationID, result.Registrations[0].RegistrationID)
	}

	resultChan = make(chan *serviceWorker.RegistrationUpdatedEvent)
	mockSocket.ServiceWorker().OnWorkerRegistrationUpdated(func(eventData *serviceWorker.RegistrationUpdatedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "ServiceWorker.workerRegistrationUpdated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestServiceWorkerOnWorkerVersionUpdated(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *serviceWorker.VersionUpdatedEvent)
	mockSocket.ServiceWorker().OnWorkerVersionUpdated(func(eventData *serviceWorker.VersionUpdatedEvent) {
		resultChan <- eventData
	})
	mockResult := &serviceWorker.VersionUpdatedEvent{
		Versions: []*serviceWorker.Version{{
			VersionID:          "version-id",
			RegistrationID:     "registration-id",
			ScriptURL:          "http://some.url",
			RunningStatus:      serviceWorker.VersionRunningStatus.Stopped,
			Status:             serviceWorker.VersionStatus.New,
			ScriptLastModified: time.Now().Unix(),
			ScriptResponseTime: time.Now().Unix(),
			ControlledClients:  []target.ID{"target-id"},
			TargetID:           target.ID("target-id"),
		}},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "ServiceWorker.workerVersionUpdated",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Versions[0].VersionID != result.Versions[0].VersionID {
		t.Errorf("Expected %s, got %s", mockResult.Versions[0].VersionID, result.Versions[0].VersionID)
	}

	resultChan = make(chan *serviceWorker.VersionUpdatedEvent)
	mockSocket.ServiceWorker().OnWorkerVersionUpdated(func(eventData *serviceWorker.VersionUpdatedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "ServiceWorker.workerVersionUpdated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
