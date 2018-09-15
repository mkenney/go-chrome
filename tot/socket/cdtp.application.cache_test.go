package socket

import (
	"testing"
	"time"

	application_cache "github.com/mkenney/go-chrome/tot/application/cache"
	"github.com/mkenney/go-chrome/tot/page"
)

func TestApplicationCacheEnable(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := MockData{
		0,
		&Error{},
		nil,
		"",
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.ApplicationCache().Enable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.ApplicationCache().Enable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
func TestApplicationCacheGetForFrame(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockParams := &application_cache.GetForFrameParams{FrameID: page.FrameID("mock-frame-id")}
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

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.ApplicationCache().GetForFrame(mockParams)
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

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.ApplicationCache().GetForFrame(mockParams)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestApplicationCacheGetFramesWithManifests(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &application_cache.GetFramesWithManifestsResult{
		FrameIDs: []*application_cache.FrameWithManifest{{
			FrameID:     page.FrameID(1),
			ManifestURL: "http://example.com/manifest",
			Status:      1,
		}},
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.ApplicationCache().GetFramesWithManifests()
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

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.ApplicationCache().GetFramesWithManifests()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestApplicationCacheGetManifestForFrame(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &application_cache.GetManifestForFrameResult{
		ManifestURL: "http://example.com/manifest",
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.ApplicationCache().GetManifestForFrame(&application_cache.GetManifestForFrameParams{
		FrameID: page.FrameID("mock-frame-id"),
	})
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

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.ApplicationCache().GetManifestForFrame(&application_cache.GetManifestForFrameParams{
		FrameID: page.FrameID("mock-frame-id"),
	})
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestApplicationCacheOnApplicationCacheStatusUpdated(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *application_cache.StatusUpdatedEvent)
	soc.ApplicationCache().OnApplicationCacheStatusUpdated(func(eventData *application_cache.StatusUpdatedEvent) {
		resultChan <- eventData
	})

	mockResult := &application_cache.StatusUpdatedEvent{
		FrameID:     page.FrameID("mock-frame-id"),
		ManifestURL: "http://example.com/manifest",
		Status:      1,
	}
	chrome.AddData(MockData{
		0,
		&Error{},
		mockResult,
		"ApplicationCache.applicationCacheStatusUpdated",
	})
	result := <-resultChan
	if result.ManifestURL != mockResult.ManifestURL {
		t.Errorf(
			"Expected frame ID %s, got %s",
			mockResult.ManifestURL,
			result.ManifestURL,
		)
	}

	chrome.AddData(MockData{
		0,
		genericError,
		nil,
		"ApplicationCache.applicationCacheStatusUpdated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestApplicationCacheOnNetworkStateUpdated(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *application_cache.NetworkStateUpdatedEvent)
	soc.ApplicationCache().OnNetworkStateUpdated(func(eventData *application_cache.NetworkStateUpdatedEvent) {
		resultChan <- eventData
	})

	mockResult := &application_cache.NetworkStateUpdatedEvent{
		IsNowOnline: true,
	}
	chrome.AddData(MockData{
		Err:    &Error{},
		Result: mockResult,
		Method: "ApplicationCache.networkStateUpdated",
	})
	result := <-resultChan
	if result.IsNowOnline != mockResult.IsNowOnline {
		t.Errorf(
			"Expected frame ID %v, got %v",
			mockResult.IsNowOnline,
			result.IsNowOnline,
		)
	}

	chrome.AddData(MockData{
		Err:    genericError,
		Result: nil,
		Method: "ApplicationCache.networkStateUpdated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
