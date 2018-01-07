package socket

import (
	"encoding/json"
	"net/url"
	"testing"
	"time"

	"github.com/mkenney/go-chrome/cdtp/animation"
	"github.com/mkenney/go-chrome/cdtp/dom"
	"github.com/mkenney/go-chrome/cdtp/runtime"
	log "github.com/sirupsen/logrus"
)

func TestAnimationDisable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Animation().Disable()
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "Animation.disable",
		Result: nil,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Animation().Disable()
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Animation.disable",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestAnimationEnable(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Animation().Enable()
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "Animation.enable",
		Result: nil,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Animation().Enable()
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Animation.enable",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
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
	mockResultBytes, _ := json.Marshal(animation.GetCurrentTimeResult{})
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "Animation.getCurrentTime",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected success, got error: '%s'", result.Err.Error())
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
	mockResultBytes, _ = json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "Animation.getCurrentTime",
		Result: mockResultBytes,
	})
	result = <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if result.CurrentTime != mockResult.CurrentTime {
		t.Errorf(
			"Expected %d, received %d",
			mockResult.CurrentTime,
			result.CurrentTime,
		)
	}

	resultChan = mockSocket.Animation().GetCurrentTime(&animation.GetCurrentTimeParams{
		ID: "animation-id",
	})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Animation.getCurrentTime",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
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
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "Animation.getPlaybackRate",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if result.PlaybackRate != mockResult.PlaybackRate {
		t.Errorf(
			"Expected %f, received %f",
			mockResult.PlaybackRate,
			result.PlaybackRate,
		)
	}

	resultChan = mockSocket.Animation().GetPlaybackRate()
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Animation.getPlaybackRate",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestAnimationReleaseAnimations(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()
	resultChan := mockSocket.Animation().ReleaseAnimations(&animation.ReleaseAnimationsParams{})
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "Animation.releaseAnimationsParams",
		Result: nil,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Animation().ReleaseAnimations(&animation.ReleaseAnimationsParams{})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Animation.releaseAnimations",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
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
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "Animation.resolveAnimation",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}
	if result.RemoteObject.Type != mockResult.RemoteObject.Type {
		t.Errorf(
			"Expected %s, received %s",
			mockResult.RemoteObject.Type,
			result.RemoteObject.Type,
		)
	}

	resultChan = mockSocket.Animation().ResolveAnimation(&animation.ResolveAnimationParams{})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Animation.resolveAnimation",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestAnimationSeekAnimations(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Animation().SeekAnimations(&animation.SeekAnimationsParams{})
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "Animation.seekAnimationsParams",
		Result: nil,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Animation().SeekAnimations(&animation.SeekAnimationsParams{})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Animation.seekAnimations",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestAnimationSetPaused(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Animation().SetPaused(&animation.SetPausedParams{})
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "Animation.setPausedParams",
		Result: nil,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Animation().SetPaused(&animation.SetPausedParams{})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Animation.setPaused",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestAnimationSetPlaybackRate(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Animation().SetPlaybackRate(&animation.SetPlaybackRateParams{})
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "Animation.SetPlaybackRateParams",
		Result: nil,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Animation().SetPlaybackRate(&animation.SetPlaybackRateParams{})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Animation.setPlaybackRate",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestAnimationSetTiming(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := mockSocket.Animation().SetTiming(&animation.SetTimingParams{})
	mockSocket.Conn().AddMockData(&Response{
		ID:     mockSocket.CurCommandID(),
		Error:  &Error{},
		Method: "Animation.setTimingParams",
		Result: nil,
	})
	result := <-resultChan
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	resultChan = mockSocket.Animation().SetTiming(&animation.SetTimingParams{})
	mockSocket.Conn().AddMockData(&Response{
		ID: mockSocket.CurCommandID(),
		Error: &Error{
			Code:    1,
			Data:    []byte(`"error data"`),
			Message: "error message",
		},
		Method: "Animation.setTiming",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestAnimationOnAnimationCanceled(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultChan := make(chan *animation.CanceledEvent)
	mockSocket.Animation().OnAnimationCanceled(func(eventData *animation.CanceledEvent) {
		resultChan <- eventData
	})

	mockResult := &animation.CanceledEvent{
		ID: "event-id",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Animation.animationCanceled",
		Result: mockResultBytes,
	})
	result := <-resultChan
	if result.ID != mockResult.ID {
		t.Errorf(
			"Expected frame ID '%v', got '%v'",
			mockResult.ID,
			result.ID,
		)
	}
}

func TestAnimationOnAnimationCreated(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultsChan := make(chan *animation.CreatedEvent)
	mockSocket.Animation().OnAnimationCreated(func(eventData *animation.CreatedEvent) {
		resultsChan <- eventData
	})

	mockResult := &animation.CreatedEvent{
		ID: "event-id",
	}
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Animation.animationCreated",
		Result: mockResultBytes,
	})
	result := <-resultsChan
	if result.ID != mockResult.ID {
		t.Errorf(
			"Expected frame ID %v, got %v",
			mockResult.ID,
			result.ID,
		)
	}
}

func TestAnimationOnAnimationStarted(t *testing.T) {
	socketURL, _ := url.Parse("https://test:9222/")
	mockSocket := NewMock(socketURL)
	go mockSocket.Listen()
	defer mockSocket.Stop()

	resultsChan := make(chan *animation.StartedEvent)
	mockSocket.Animation().OnAnimationStarted(func(eventData *animation.StartedEvent) {
		resultsChan <- eventData
	})
	mockResult := &animation.StartedEvent{
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
	mockResultBytes, _ := json.Marshal(mockResult)
	mockSocket.Conn().AddMockData(&Response{
		ID:     0,
		Error:  &Error{},
		Method: "Animation.animationStarted",
		Result: mockResultBytes,
	})
	result := <-resultsChan
	tmp, _ := json.Marshal(result)
	log.Debugf("mock: %s", mockResultBytes)
	log.Debugf("result: %s", tmp)
	if result.Animation.ID != mockResult.Animation.ID {
		t.Errorf(
			"Expected %s, got %s",
			mockResult.Animation.ID,
			result.Animation.ID,
		)
	}
}
