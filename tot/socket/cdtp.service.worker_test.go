package socket

import (
	"encoding/json"
	"net/url"
	"testing"
	"time"

	"github.com/mkenney/go-chrome/tot/service/worker"
	"github.com/mkenney/go-chrome/tot/target"
)

func TestServiceWorkerDeliverPushMessage(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestServiceWorkerDeliverPushMessage")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	params := &worker.DeliverPushMessageParams{
		Origin:         "origin",
		RegistrationID: "registration-id",
		Data:           "data",
	}
	resultChan := mockSocket.ServiceWorker().DeliverPushMessage(params)
	mockResult := &worker.DeliverPushMessageResult{}
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

	resultChan = mockSocket.ServiceWorker().DeliverPushMessage(params)
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

func TestServiceWorkerDisable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestServiceWorkerDisable")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.ServiceWorker().Disable()
	mockResult := &worker.DisableResult{}
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

	resultChan = mockSocket.ServiceWorker().Disable()
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

func TestServiceWorkerDispatchSyncEvent(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestServiceWorkerDispatchSyncEvent")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	params := &worker.DispatchSyncEventParams{
		Origin:         "origin",
		RegistrationID: "registration-id",
		Tag:            "tag",
		LastChance:     true,
	}
	resultChan := mockSocket.ServiceWorker().DispatchSyncEvent(params)
	mockResult := &worker.DispatchSyncEventResult{}
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

	resultChan = mockSocket.ServiceWorker().DispatchSyncEvent(params)
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

func TestServiceWorkerEnable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestServiceWorkerEnable")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.ServiceWorker().Enable()
	mockResult := &worker.EnableResult{}
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

	resultChan = mockSocket.ServiceWorker().Enable()
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

func TestServiceWorkerInspectWorker(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestServiceWorkerInspectWorker")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	params := &worker.InspectWorkerParams{
		VersionID: "version-id",
	}
	resultChan := mockSocket.ServiceWorker().InspectWorker(params)
	mockResult := &worker.InspectWorkerResult{}
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

	resultChan = mockSocket.ServiceWorker().InspectWorker(params)
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

func TestServiceWorkerSetForceUpdateOnPageLoad(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestServiceWorkerSetForceUpdateOnPageLoad")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	params := &worker.SetForceUpdateOnPageLoadParams{
		ForceUpdateOnPageLoad: true,
	}
	resultChan := mockSocket.ServiceWorker().SetForceUpdateOnPageLoad(params)
	mockResult := &worker.SetForceUpdateOnPageLoadResult{}
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

	resultChan = mockSocket.ServiceWorker().SetForceUpdateOnPageLoad(params)
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

func TestServiceWorkerSkipWaiting(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestServiceWorkerSkipWaiting")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	params := &worker.SkipWaitingParams{
		ScopeURL: "http://some.url",
	}
	resultChan := mockSocket.ServiceWorker().SkipWaiting(params)
	mockResult := &worker.SkipWaitingResult{}
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

	resultChan = mockSocket.ServiceWorker().SkipWaiting(params)
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

func TestServiceWorkerStartWorker(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestServiceWorkerStartWorker")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	params := &worker.StartWorkerParams{
		ScopeURL: "http://some.url",
	}
	resultChan := mockSocket.ServiceWorker().StartWorker(params)
	mockResult := &worker.StartWorkerResult{}
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

	resultChan = mockSocket.ServiceWorker().StartWorker(params)
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

func TestServiceWorkerStopAllWorkers(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestServiceWorkerStopAllWorkers")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.ServiceWorker().StopAllWorkers()
	mockResult := &worker.StopAllWorkersResult{}
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

	resultChan = mockSocket.ServiceWorker().StopAllWorkers()
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

func TestServiceWorkerStopWorker(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestServiceWorkerStopWorker")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	params := &worker.StopWorkerParams{
		VersionID: "version-id",
	}
	resultChan := mockSocket.ServiceWorker().StopWorker(params)
	mockResult := &worker.StopWorkerResult{}
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

	resultChan = mockSocket.ServiceWorker().StopWorker(params)
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

func TestServiceWorkerUnregister(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestServiceWorkerUnregister")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	params := &worker.UnregisterParams{
		ScopeURL: "http://some.url",
	}
	resultChan := mockSocket.ServiceWorker().Unregister(params)
	mockResult := &worker.UnregisterResult{}
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

	resultChan = mockSocket.ServiceWorker().Unregister(params)
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

func TestServiceWorkerUpdateRegistration(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestServiceWorkerUpdateRegistration")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	params := &worker.UpdateRegistrationParams{
		ScopeURL: "http://some.url",
	}
	resultChan := mockSocket.ServiceWorker().UpdateRegistration(params)
	mockResult := &worker.UpdateRegistrationResult{}
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

	resultChan = mockSocket.ServiceWorker().UpdateRegistration(params)
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

func TestServiceWorkerOnWorkerErrorReported(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/TestServiceWorkerOnWorkerErrorReported")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *worker.ErrorReportedEvent)
	mockSocket.ServiceWorker().OnWorkerErrorReported(func(eventData *worker.ErrorReportedEvent) {
		resultChan <- eventData
	})
	mockResult := &worker.ErrorReportedEvent{
		ErrorMessage: &worker.ErrorMessage{
			ErrorMessage:   "error message",
			RegistrationID: "registration-id",
			VersionID:      "version-id",
			SourceURL:      "http://some.url",
			LineNumber:     1,
			ColumnNumber:   1,
		},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "ServiceWorker.workerErrorReported",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.ErrorMessage.ErrorMessage != result.ErrorMessage.ErrorMessage {
		t.Errorf("Expected %s, got %s", mockResult.ErrorMessage.ErrorMessage, result.ErrorMessage.ErrorMessage)
	}

	resultChan = make(chan *worker.ErrorReportedEvent)
	mockSocket.ServiceWorker().OnWorkerErrorReported(func(eventData *worker.ErrorReportedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
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
	socketURL, _ := url.Parse("https://test:9222/TestServiceWorkerOnWorkerRegistrationUpdated")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *worker.RegistrationUpdatedEvent)
	mockSocket.ServiceWorker().OnWorkerRegistrationUpdated(func(eventData *worker.RegistrationUpdatedEvent) {
		resultChan <- eventData
	})
	mockResult := &worker.RegistrationUpdatedEvent{
		Registrations: []*worker.Registration{{
			RegistrationID: "registration-id",
			ScopeURL:       "http://some.url",
			IsDeleted:      true,
		}},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "ServiceWorker.workerRegistrationUpdated",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Registrations[0].RegistrationID != result.Registrations[0].RegistrationID {
		t.Errorf("Expected %s, got %s", mockResult.Registrations[0].RegistrationID, result.Registrations[0].RegistrationID)
	}

	resultChan = make(chan *worker.RegistrationUpdatedEvent)
	mockSocket.ServiceWorker().OnWorkerRegistrationUpdated(func(eventData *worker.RegistrationUpdatedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
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
	socketURL, _ := url.Parse("https://test:9222/TestServiceWorkerOnWorkerVersionUpdated")
	mockSocket := NewMock(socketURL)
	mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *worker.VersionUpdatedEvent)
	mockSocket.ServiceWorker().OnWorkerVersionUpdated(func(eventData *worker.VersionUpdatedEvent) {
		resultChan <- eventData
	})
	mockResult := &worker.VersionUpdatedEvent{
		Versions: []*worker.Version{{
			VersionID:          "version-id",
			RegistrationID:     "registration-id",
			ScriptURL:          "http://some.url",
			RunningStatus:      worker.VersionRunningStatus.Stopped,
			Status:             worker.VersionStatus.New,
			ScriptLastModified: time.Now().Unix(),
			ScriptResponseTime: time.Now().Unix(),
			ControlledClients:  []target.ID{"target-id"},
			TargetID:           target.ID("target-id"),
		}},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "ServiceWorker.workerVersionUpdated",
		Params: mockResultBytes,
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Versions[0].VersionID != result.Versions[0].VersionID {
		t.Errorf("Expected %s, got %s", mockResult.Versions[0].VersionID, result.Versions[0].VersionID)
	}

	resultChan = make(chan *worker.VersionUpdatedEvent)
	mockSocket.ServiceWorker().OnWorkerVersionUpdated(func(eventData *worker.VersionUpdatedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().(*MockChromeWebSocket).AddMockData(&Response{
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
