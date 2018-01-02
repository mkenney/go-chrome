package socket

import (
	"net/url"
	"testing"
	"time"

	application_cache "github.com/mkenney/go-chrome/protocol/application_cache"
	"github.com/mkenney/go-chrome/protocol/page"
	log "github.com/sirupsen/logrus"
)

func init() {
	level, err := log.ParseLevel("debug")
	if nil == err {
		log.SetLevel(level)
	}
}

func TestApplicationCacheEnable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	mockSocket.Conn().AddMockData(
		_commandID+1,
		&Error{},
		"ApplicationCache.enable",
		nil,
	)
	go func() {
		err := mockSocket.ApplicationCache().Enable()
		if nil != err {
			t.Errorf("Expected nil, got error: '%s'", err.Error())
		}
	}()
}
func TestApplicationCacheGetForFrame(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	mockGetForFrameResult := &application_cache.GetForFrameResult{
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
	mockSocket.Conn().AddMockData(
		_commandID+2,
		&Error{},
		"ApplicationCache.getForFrame",
		mockGetForFrameResult,
		&application_cache.GetForFrameParams{FrameID: page.FrameID("mock-frame-id")},
	)
	go func() {
		result, err := mockSocket.ApplicationCache().GetForFrame(&application_cache.GetForFrameParams{})
		if nil != err {
			t.Errorf("Expected nil, got error: '%s'", err.Error())
		}
		if result.ApplicationCache.CreationTime != mockGetForFrameResult.ApplicationCache.CreationTime {
			t.Errorf(
				"Expected creation time %f, '%f'",
				mockGetForFrameResult.ApplicationCache.CreationTime,
				result.ApplicationCache.CreationTime,
			)
		}
	}()
}

func TestApplicationCacheGetFramesWithManifests(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	mockGetFramesWithManifestsResult := &application_cache.GetFramesWithManifestsResult{
		FrameIDs: []*application_cache.FrameWithManifest{{
			FrameID:     page.FrameID(1),
			ManifestURL: "http://example.com/manifest",
			Status:      1,
		}},
	}
	mockSocket.Conn().AddMockData(
		_commandID+2,
		&Error{},
		"ApplicationCache.getFramesWithManifests",
		mockGetFramesWithManifestsResult,
		nil,
	)
	go func() {
		result, err := mockSocket.ApplicationCache().GetFramesWithManifests()
		if nil != err {
			t.Errorf("Expected nil, got error: '%s'", err.Error())
		}
		if result.FrameIDs[0].FrameID != mockGetFramesWithManifestsResult.FrameIDs[0].FrameID {
			t.Errorf(
				"Expected frame ID %s, got %s",
				mockGetFramesWithManifestsResult.FrameIDs[0].FrameID,
				result.FrameIDs[0].FrameID,
			)
		}
	}()
}

func TestApplicationCacheGetManifestForFrame(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	mockGetManifestForFrameResult := &application_cache.GetManifestForFrameResult{
		ManifestURL: "http://example.com/manifest",
	}
	mockSocket.Conn().AddMockData(
		_commandID+2,
		&Error{},
		"ApplicationCache.getManifestForFrame",
		mockGetManifestForFrameResult,
		nil,
	)
	go func() {
		result, err := mockSocket.ApplicationCache().GetManifestForFrame(&application_cache.GetManifestForFrameParams{
			FrameID: page.FrameID("mock-frame-id"),
		})
		if nil != err {
			t.Errorf("Expected nil, got error: '%s'", err.Error())
		}
		if result.ManifestURL != mockGetManifestForFrameResult.ManifestURL {
			t.Errorf(
				"Expected frame ID %s, got %s",
				mockGetManifestForFrameResult.ManifestURL,
				result.ManifestURL,
			)
		}
	}()
}

func TestApplicationCacheOnApplicationCacheStatusUpdated(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	mockOnApplicationCacheStatusUpdatedEvent := &application_cache.StatusUpdatedEvent{
		FrameID:     page.FrameID("mock-frame-id"),
		ManifestURL: "http://example.com/manifest",
		Status:      1,
	}
	mockSocket.Conn().AddMockData(
		0,
		&Error{},
		"ApplicationCache.applicationCacheStatusUpdated",
		mockOnApplicationCacheStatusUpdatedEvent,
		nil,
	)
	results := make(chan *application_cache.StatusUpdatedEvent)
	mockSocket.ApplicationCache().OnApplicationCacheStatusUpdated(func(event *application_cache.StatusUpdatedEvent) {
		results <- mockOnApplicationCacheStatusUpdatedEvent
	})
	result := <-results
	if result.ManifestURL != mockOnApplicationCacheStatusUpdatedEvent.ManifestURL {
		t.Errorf(
			"Expected frame ID %s, got %s",
			mockOnApplicationCacheStatusUpdatedEvent.ManifestURL,
			result.ManifestURL,
		)
	}
}

func TestApplicationCacheOnNetworkStateUpdated(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	mockOnNetworkStateUpdatedEvent := &application_cache.NetworkStateUpdatedEvent{
		IsNowOnline: true,
	}
	mockSocket.Conn().AddMockData(
		0,
		&Error{},
		"ApplicationCache.networkStateUpdated",
		mockOnNetworkStateUpdatedEvent,
		nil,
	)
	results := make(chan *application_cache.NetworkStateUpdatedEvent)
	mockSocket.ApplicationCache().OnNetworkStateUpdated(func(event *application_cache.NetworkStateUpdatedEvent) {
		results <- mockOnNetworkStateUpdatedEvent
	})
	result := <-results
	if result.IsNowOnline != mockOnNetworkStateUpdatedEvent.IsNowOnline {
		t.Errorf(
			"Expected frame ID %v, got %v",
			mockOnNetworkStateUpdatedEvent.IsNowOnline,
			result.IsNowOnline,
		)
	}
}
