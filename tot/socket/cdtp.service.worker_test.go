package socket

import (
	"testing"
	"time"

	"github.com/mkenney/go-chrome/tot/service/worker"
	"github.com/mkenney/go-chrome/tot/target"
)

func TestServiceWorkerDeliverPushMessage(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &worker.DeliverPushMessageParams{
		Origin:         "origin",
		RegistrationID: "registration-id",
		Data:           "data",
	}
	mockResult := &worker.DeliverPushMessageResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.ServiceWorker().DeliverPushMessage(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.ServiceWorker().DeliverPushMessage(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestServiceWorkerDisable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &worker.DisableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.ServiceWorker().Disable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.ServiceWorker().Disable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestServiceWorkerDispatchSyncEvent(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &worker.DispatchSyncEventParams{
		Origin:         "origin",
		RegistrationID: "registration-id",
		Tag:            "tag",
		LastChance:     true,
	}
	mockResult := &worker.DispatchSyncEventResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.ServiceWorker().DispatchSyncEvent(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.ServiceWorker().DispatchSyncEvent(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestServiceWorkerEnable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &worker.EnableResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.ServiceWorker().Enable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.ServiceWorker().Enable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestServiceWorkerInspectWorker(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &worker.InspectWorkerParams{
		VersionID: "version-id",
	}
	mockResult := &worker.InspectWorkerResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.ServiceWorker().InspectWorker(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.ServiceWorker().InspectWorker(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestServiceWorkerSetForceUpdateOnPageLoad(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &worker.SetForceUpdateOnPageLoadParams{
		ForceUpdateOnPageLoad: true,
	}
	mockResult := &worker.SetForceUpdateOnPageLoadResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.ServiceWorker().SetForceUpdateOnPageLoad(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.ServiceWorker().SetForceUpdateOnPageLoad(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestServiceWorkerSkipWaiting(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &worker.SkipWaitingParams{
		ScopeURL: "http://some.url",
	}
	mockResult := &worker.SkipWaitingResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.ServiceWorker().SkipWaiting(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.ServiceWorker().SkipWaiting(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestServiceWorkerStartWorker(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &worker.StartWorkerParams{
		ScopeURL: "http://some.url",
	}
	mockResult := &worker.StartWorkerResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.ServiceWorker().StartWorker(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.ServiceWorker().StartWorker(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestServiceWorkerStopAllWorkers(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &worker.StopAllWorkersResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.ServiceWorker().StopAllWorkers()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.ServiceWorker().StopAllWorkers()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestServiceWorkerStopWorker(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &worker.StopWorkerParams{
		VersionID: "version-id",
	}
	mockResult := &worker.StopWorkerResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.ServiceWorker().StopWorker(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.ServiceWorker().StopWorker(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestServiceWorkerUnregister(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &worker.UnregisterParams{
		ScopeURL: "http://some.url",
	}
	mockResult := &worker.UnregisterResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.ServiceWorker().Unregister(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.ServiceWorker().Unregister(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestServiceWorkerUpdateRegistration(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &worker.UpdateRegistrationParams{
		ScopeURL: "http://some.url",
	}
	mockResult := &worker.UpdateRegistrationResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.ServiceWorker().UpdateRegistration(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.ServiceWorker().UpdateRegistration(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestServiceWorkerOnWorkerErrorReported(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *worker.ErrorReportedEvent)
	soc.ServiceWorker().OnWorkerErrorReported(func(eventData *worker.ErrorReportedEvent) {
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
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "ServiceWorker.workerErrorReported",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.ErrorMessage.ErrorMessage != result.ErrorMessage.ErrorMessage {
		t.Errorf("Expected %s, got %s", mockResult.ErrorMessage.ErrorMessage, result.ErrorMessage.ErrorMessage)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "ServiceWorker.workerErrorReported",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestServiceWorkerOnWorkerRegistrationUpdated(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *worker.RegistrationUpdatedEvent)
	soc.ServiceWorker().OnWorkerRegistrationUpdated(func(eventData *worker.RegistrationUpdatedEvent) {
		resultChan <- eventData
	})

	mockResult := &worker.RegistrationUpdatedEvent{
		Registrations: []*worker.Registration{{
			RegistrationID: "registration-id",
			ScopeURL:       "http://some.url",
			IsDeleted:      true,
		}},
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "ServiceWorker.workerRegistrationUpdated",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Registrations[0].RegistrationID != result.Registrations[0].RegistrationID {
		t.Errorf("Expected %s, got %s", mockResult.Registrations[0].RegistrationID, result.Registrations[0].RegistrationID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "ServiceWorker.workerRegistrationUpdated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestServiceWorkerOnWorkerVersionUpdated(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *worker.VersionUpdatedEvent)
	soc.ServiceWorker().OnWorkerVersionUpdated(func(eventData *worker.VersionUpdatedEvent) {
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
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "ServiceWorker.workerVersionUpdated",
	})
	result := <-resultChan
	if mockResult.Err != result.Err {
		t.Errorf("Expected '%v', got: '%v'", mockResult, result)
	}
	if mockResult.Versions[0].VersionID != result.Versions[0].VersionID {
		t.Errorf("Expected %s, got %s", mockResult.Versions[0].VersionID, result.Versions[0].VersionID)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "ServiceWorker.workerVersionUpdated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
