package socket

import (
	"encoding/json"
	"net/url"
	"testing"
	"time"

	application_cache "github.com/mkenney/go-chrome/cdtp/application_cache"
	"github.com/mkenney/go-chrome/cdtp/page"
)

func TestApplicationCacheEnable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.ApplicationCache().Enable()
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "ApplicationCache.enable",
		Result: nil,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.ApplicationCache().Enable()
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "ApplicationCache.enable",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
func TestApplicationCacheGetForFrame(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	mockParams := &application_cache.GetForFrameParams{FrameID: page.FrameID("mock-frame-id")}
	resultChan := mockSocket.ApplicationCache().GetForFrame(mockParams)
	mockResult := &application_cache.GetForFrameResult{
		ApplicationCache: &application_cache.ApplicationCache{
			ManifestURL:  "http://example.com/manifest",
			Size:         1.1,
			CreationTime: float64(time.Now().Unix()),
			UpdateTime:   float64(time.Now().Unix()),
			Resources: []*application_cache.Resource{{
				URL:  "http://example.com/",
				Size: 1,
				Type: "Type",
			}},
		},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockParamsBytes, _ := json.Marshal(mockParams)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "ApplicationCache.getForFrame",
		Params: mockParamsBytes,
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if result.ApplicationCache.CreationTime != mockResult.ApplicationCache.CreationTime {
		t.Errorf(
			"Expected creation time %f, '%f'",
			mockResult.ApplicationCache.CreationTime,
			result.ApplicationCache.CreationTime,
		)
	}

	resultChan = mockSocket.ApplicationCache().GetForFrame(mockParams)
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "ApplicationCache.getForFrame",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestApplicationCacheGetFramesWithManifests(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.ApplicationCache().GetFramesWithManifests()
	mockResult := &application_cache.GetFramesWithManifestsResult{
		FrameIDs: []*application_cache.FrameWithManifest{{
			FrameID:     page.FrameID(1),
			ManifestURL: "http://example.com/manifest",
			Status:      1,
		}},
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "ApplicationCache.getFramesWithManifests",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if result.FrameIDs[0].FrameID != mockResult.FrameIDs[0].FrameID {
		t.Errorf(
			"Expected frame ID %s, got %s",
			mockResult.FrameIDs[0].FrameID,
			result.FrameIDs[0].FrameID,
		)
	}

	resultChan = mockSocket.ApplicationCache().GetFramesWithManifests()
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "ApplicationCache.getFramesWithManifests",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestApplicationCacheGetManifestForFrame(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.ApplicationCache().GetManifestForFrame(&application_cache.GetManifestForFrameParams{
		FrameID: page.FrameID("mock-frame-id"),
	})
	mockResult := &application_cache.GetManifestForFrameResult{
		ManifestURL: "http://example.com/manifest",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "ApplicationCache.getManifestForFrame",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if result.ManifestURL != mockResult.ManifestURL {
		t.Errorf(
			"Expected frame ID %s, got %s",
			mockResult.ManifestURL,
			result.ManifestURL,
		)
	}

	resultChan = mockSocket.ApplicationCache().GetManifestForFrame(&application_cache.GetManifestForFrameParams{
		FrameID: page.FrameID("mock-frame-id"),
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "ApplicationCache.getManifestForFrame",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestApplicationCacheOnApplicationCacheStatusUpdated(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *application_cache.StatusUpdatedEvent)
	mockSocket.ApplicationCache().OnApplicationCacheStatusUpdated(func(eventData *application_cache.StatusUpdatedEvent) {
		resultChan <- eventData
	})
	mockResult := &application_cache.StatusUpdatedEvent{
		FrameID:     page.FrameID("mock-frame-id"),
		ManifestURL: "http://example.com/manifest",
		Status:      1,
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "ApplicationCache.applicationCacheStatusUpdated",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if result.ManifestURL != mockResult.ManifestURL {
		t.Errorf(
			"Expected frame ID %s, got %s",
			mockResult.ManifestURL,
			result.ManifestURL,
		)
	}

	resultChan = make(chan *application_cache.StatusUpdatedEvent)
	mockSocket.ApplicationCache().OnApplicationCacheStatusUpdated(func(eventData *application_cache.StatusUpdatedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "ApplicationCache.applicationCacheStatusUpdated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestApplicationCacheOnNetworkStateUpdated(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *application_cache.NetworkStateUpdatedEvent)
	mockSocket.ApplicationCache().OnNetworkStateUpdated(func(eventData *application_cache.NetworkStateUpdatedEvent) {
		resultChan <- eventData
	})
	mockResult := &application_cache.NetworkStateUpdatedEvent{
		IsNowOnline: true,
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "ApplicationCache.networkStateUpdated",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if result.IsNowOnline != mockResult.IsNowOnline {
		t.Errorf(
			"Expected frame ID %v, got %v",
			mockResult.IsNowOnline,
			result.IsNowOnline,
		)
	}

	resultChan = make(chan *application_cache.NetworkStateUpdatedEvent)
	mockSocket.ApplicationCache().OnNetworkStateUpdated(func(eventData *application_cache.NetworkStateUpdatedEvent) {
		resultChan <- eventData
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: 0,
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "ApplicationCache.networkStateUpdated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
