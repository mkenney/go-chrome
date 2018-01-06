package socket

import (
	"net/url"
	"testing"
	"time"

	"github.com/mkenney/go-chrome/cdtp/animation"
	"github.com/mkenney/go-chrome/cdtp/dom"
	"github.com/mkenney/go-chrome/cdtp/runtime"
)

func TestAnimationDisable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Animation().Disable()
	mockSocket.Conn().AddMockData(
		mockSocket.CurCommandID(),
		&Error{},
		"Animation.disable",
		nil,
	)
	result := <-resultChan
	if nil != result.CDTPError {
		t.Errorf("Expected nil, got error: '%s'", result.CDTPError.Error())
	}
}

func TestAnimationEnable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Animation().Enable()
	mockSocket.Conn().AddMockData(
		mockSocket.CurCommandID(),
		&Error{},
		"Animation.enable",
		nil,
	)
	result := <-resultChan
	if nil != result.CDTPError {
		t.Errorf("Expected nil, got error: '%s'", result.CDTPError.Error())
	}
}

func TestAnimationGetCurrentTime(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Animation().GetCurrentTime(
		&animation.GetCurrentTimeParams{},
	)
	mockSocket.Conn().AddMockData(
		mockSocket.CurCommandID(),
		&Error{},
		"Animation.getCurrentTime",
		&animation.GetCurrentTimeResult{},
	)
	result := <-resultChan
	if nil != result.CDTPError {
		t.Errorf("Expected success, got error: '%s'", result.CDTPError.Error())
	}
	if 0 != result.CurrentTime {
		t.Errorf("Expected 0, got %d", result.CurrentTime)
	}

	resultChan = mockSocket.Animation().GetCurrentTime(&animation.GetCurrentTimeParams{
		ID: "animation-id",
	})
	mockResult := &animation.GetCurrentTimeResult{
		CurrentTime: time.Now().Unix(),
	}
	mockSocket.Conn().AddMockData(
		mockSocket.CurCommandID(),
		&Error{},
		"Animation.getCurrentTime",
		mockResult,
	)
	result = <-resultChan
	if nil != result.CDTPError {
		t.Errorf("Expected nil, got error: '%s'", result.CDTPError.Error())
	}
	if result.CurrentTime != mockResult.CurrentTime {
		t.Errorf(
			"Expected %d, received %d",
			mockResult.CurrentTime,
			result.CurrentTime,
		)
	}
}

func TestAnimationGetPlaybackRate(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Animation().GetPlaybackRate()
	mockResult := &animation.GetPlaybackRateResult{
		PlaybackRate: 1.0,
	}
	mockSocket.Conn().AddMockData(
		mockSocket.CurCommandID(),
		&Error{},
		"Animation.getPlaybackRate",
		mockResult,
	)
	result := <-resultChan
	if nil != result.CDTPError {
		t.Errorf("Expected nil, got error: '%s'", result.CDTPError.Error())
	}
	if result.PlaybackRate != mockResult.PlaybackRate {
		t.Errorf(
			"Expected %f, received %f",
			mockResult.PlaybackRate,
			result.PlaybackRate,
		)
	}
}

func TestAnimationReleaseAnimations(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()
	resultChan := mockSocket.Animation().ReleaseAnimations(&animation.ReleaseAnimationsParams{})
	mockSocket.Conn().AddMockData(
		mockSocket.CurCommandID(),
		&Error{},
		"Animation.releaseAnimationsParams",
		nil,
	)
	result := <-resultChan
	if nil != result.CDTPError {
		t.Errorf("Expected nil, got error: '%s'", result.CDTPError.Error())
	}
}

func TestAnimationResolveAnimation(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Animation().ResolveAnimation(&animation.ResolveAnimationParams{})
	mockResult := &animation.ResolveAnimationResult{
		RemoteObject: &runtime.RemoteObject{
			Type:                "array",
			ClassName:           "some-class",
			Value:               nil,
			UnserializableValue: runtime.UnserializableValue(1),
			Description:         "Animation description",
			ObjectID:            "object-id",
			Preview: &runtime.ObjectPreview{
				Type:        "object",
				Subtype:     "weakmap",
				Description: "Preview description",
				Overflow:    true,
				Properties:  []*runtime.PropertyPreview{},
				Entries:     []*runtime.EntryPreview{},
			},
			CustomPreview: &runtime.CustomPreview{
				Header:                     "Some: header",
				HasBody:                    true,
				FormatterObjectID:          runtime.RemoteObjectID("remote-object-id"),
				BindRemoteObjectFunctionID: runtime.RemoteObjectID("bind-remote-object-function-id"),
				ConfigObjectID:             runtime.RemoteObjectID("config-object-id"),
			},
		},
	}
	mockSocket.Conn().AddMockData(
		mockSocket.CurCommandID(),
		&Error{},
		"Animation.resolveAnimation",
		mockResult,
	)
	result := <-resultChan
	if nil != result.CDTPError {
		t.Errorf("Expected nil, got error: '%s'", result.CDTPError.Error())
	}
	if result.RemoteObject.Type != mockResult.RemoteObject.Type {
		t.Errorf(
			"Expected %s, received %s",
			mockResult.RemoteObject.Type,
			result.RemoteObject.Type,
		)
	}
}

func TestAnimationSeekAnimations(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Animation().SeekAnimations(&animation.SeekAnimationsParams{})
	mockSocket.Conn().AddMockData(
		mockSocket.CurCommandID(),
		&Error{},
		"Animation.seekAnimationsParams",
		nil,
	)
	result := <-resultChan
	if nil != result.CDTPError {
		t.Errorf("Expected nil, got error: '%s'", result.CDTPError.Error())
	}
}

func TestAnimationSetPaused(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Animation().SetPaused(&animation.SetPausedParams{})
	mockSocket.Conn().AddMockData(
		mockSocket.CurCommandID(),
		&Error{},
		"Animation.setPausedParams",
		nil,
	)
	result := <-resultChan
	if nil != result.CDTPError {
		t.Errorf("Expected nil, got error: '%s'", result.CDTPError.Error())
	}
}

func TestAnimationSetPlaybackRate(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Animation().SetPlaybackRate(&animation.SetPlaybackRateParams{})
	mockSocket.Conn().AddMockData(
		mockSocket.CurCommandID(),
		&Error{},
		"Animation.SetPlaybackRateParams",
		nil,
	)
	result := <-resultChan
	if nil != result.CDTPError {
		t.Errorf("Expected nil, got error: '%s'", result.CDTPError.Error())
	}
}

func TestAnimationSetTiming(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Animation().SetTiming(&animation.SetTimingParams{})
	mockSocket.Conn().AddMockData(
		mockSocket.CurCommandID(),
		&Error{},
		"Animation.setTimingParams",
		nil,
	)
	result := <-resultChan
	if nil != result.CDTPError {
		t.Errorf("Expected nil, got error: '%s'", result.CDTPError.Error())
	}
}

func TestAnimationOnAnimationCanceled(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	mockData := &animation.CanceledEvent{
		ID: "event-id",
	}
	mockSocket.Conn().AddMockData(
		0,
		&Error{},
		"Animation.animationCanceled",
		mockData,
	)
	results := make(chan *animation.CanceledEvent)
	mockSocket.Animation().OnAnimationCanceled(func(event *animation.CanceledEvent) {
		results <- mockData
	})
	result := <-results
	if result.ID != mockData.ID {
		t.Errorf(
			"Expected frame ID %v, got %v",
			mockData.ID,
			result.ID,
		)
	}
}

func TestAnimationOnAnimationCreated(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	mockData := &animation.CreatedEvent{
		ID: "event-id",
	}
	mockSocket.Conn().AddMockData(
		0,
		&Error{},
		"Animation.animationCreated",
		mockData,
	)
	results := make(chan *animation.CreatedEvent)
	mockSocket.Animation().OnAnimationCreated(func(event *animation.CreatedEvent) {
		results <- mockData
	})
	result := <-results
	if result.ID != mockData.ID {
		t.Errorf(
			"Expected frame ID %v, got %v",
			mockData.ID,
			result.ID,
		)
	}
}

func TestAnimationOnAnimationStarted(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	mockData := &animation.StartedEvent{
		Animation: &animation.Animation{
			ID:           "animation-id",
			Name:         "Animation name",
			PausedState:  true,
			PlayState:    "play state",
			PlaybackRate: 1.1,
			StartTime:    time.Now().Unix(),
			CurrentTime:  time.Now().Unix(),
			Type:         "animation-type",
			Source: &animation.Effect{
				Delay:          1,
				EndDelay:       1,
				IterationStart: 1,
				Iterations:     1,
				Duration:       1,
				Direction:      "direction",
				Fill:           "fill",
				BackendNodeID:  dom.BackendNodeID(1),
				KeyframesRule: &animation.KeyframesRule{
					Name:      "keyframe-rule-name",
					Keyframes: []*animation.KeyframeStyle{},
				},
				Easing: "easing",
			},
			CSSID: "css-id",
		},
	}
	mockSocket.Conn().AddMockData(
		0,
		&Error{},
		"Animation.animationStarted",
		mockData,
	)
	results := make(chan *animation.StartedEvent)
	mockSocket.Animation().OnAnimationStarted(func(event *animation.StartedEvent) {
		results <- mockData
	})
	result := <-results
	if result.Animation.ID != mockData.Animation.ID {
		t.Errorf(
			"Expected %s, got %s",
			mockData.Animation.ID,
			result.Animation.ID,
		)
	}
}
