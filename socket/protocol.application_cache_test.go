package socket

import (
	"net/url"
	"testing"
	"time"

	applicationCache "github.com/mkenney/go-chrome/protocol/application_cache"
	page "github.com/mkenney/go-chrome/protocol/page"
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

	mockGetForFrameResult := &applicationCache.GetForFrameResult{
		ApplicationCache: &applicationCache.ApplicationCache{
			ManifestURL:  "http://example.com/manifest",
			Size:         1.1,
			CreationTime: float64(time.Now().Unix()),
			UpdateTime:   float64(time.Now().Unix()),
			Resources: []*applicationCache.Resource{{
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
		&applicationCache.GetForFrameParams{FrameID: page.FrameID("mock-frame-id")},
	)
	go func() {
		result, err := mockSocket.ApplicationCache().GetForFrame(&applicationCache.GetForFrameParams{})
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

	mockGetFramesWithManifestsResult := &applicationCache.GetFramesWithManifestsResult{
		FrameIDs: []*applicationCache.FrameWithManifest{{
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

	mockGetManifestForFrameResult := &applicationCache.GetManifestForFrameResult{
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
		result, err := mockSocket.ApplicationCache().GetManifestForFrame(&applicationCache.GetManifestForFrameParams{
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

	mockOnApplicationCacheStatusUpdatedEvent := &applicationCache.StatusUpdatedEvent{
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
	results := make(chan *applicationCache.StatusUpdatedEvent)
	mockSocket.ApplicationCache().OnApplicationCacheStatusUpdated(func(event *applicationCache.StatusUpdatedEvent) {
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

	mockOnNetworkStateUpdatedEvent := &applicationCache.NetworkStateUpdatedEvent{
		IsNowOnline: true,
	}
	mockSocket.Conn().AddMockData(
		0,
		&Error{},
		"ApplicationCache.networkStateUpdated",
		mockOnNetworkStateUpdatedEvent,
		nil,
	)
	results := make(chan *applicationCache.NetworkStateUpdatedEvent)
	mockSocket.ApplicationCache().OnNetworkStateUpdated(func(event *applicationCache.NetworkStateUpdatedEvent) {
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
