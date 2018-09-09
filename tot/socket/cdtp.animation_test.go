package socket

import (
	"testing"
	"time"

	"github.com/mkenney/go-chrome/tot/animation"
	"github.com/mkenney/go-chrome/tot/dom"
	"github.com/mkenney/go-chrome/tot/runtime"
)

func TestAnimationDisable(t *testing.T) {
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
	result := <-soc.Animation().Disable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Animation().Disable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestAnimationEnable(t *testing.T) {
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
	result := <-soc.Animation().Enable()
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Animation().Enable()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestAnimationGetCurrentTime(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &animation.GetCurrentTimeParams{
		ID: "animation-id",
	}
	mockResult := &animation.GetCurrentTimeResult{}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Animation().GetCurrentTime(params)
	if nil != result.Err {
		t.Errorf("Expected success, got error: '%s'", result.Err.Error())
	}
	if 0 != result.CurrentTime {
		t.Errorf("Expected 0, got %d", result.CurrentTime)
	}

	mockResult = &animation.GetCurrentTimeResult{
		CurrentTime: time.Now().Unix(),
	}
	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result = <-soc.Animation().GetCurrentTime(&animation.GetCurrentTimeParams{
		ID: "animation-id",
	})
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

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Animation().GetCurrentTime(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestAnimationGetPlaybackRate(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	mockResult := &animation.GetPlaybackRateResult{
		PlaybackRate: 1.0,
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Animation().GetPlaybackRate()
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

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Animation().GetPlaybackRate()
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestAnimationReleaseAnimations(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &animation.ReleaseAnimationsParams{
		Animations: []string{"animation1", "animation2"},
	}

	mockResult := MockData{
		0,
		&Error{},
		nil,
		"",
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Animation().ReleaseAnimations(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Animation().ReleaseAnimations(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestAnimationResolveAnimation(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &animation.ResolveAnimationParams{
		AnimationID: "animation-id",
	}

	mockResult := &animation.ResolveAnimationResult{
		RemoteObject: &runtime.RemoteObject{
			Type:                runtime.ObjectType.Object,
			ClassName:           "some-class",
			Value:               nil,
			UnserializableValue: runtime.UnserializableValue.Infinity,
			Description:         "Animation description",
			ObjectID:            "object-id",
			Preview: &runtime.ObjectPreview{
				Type:        runtime.ObjectType.Object,
				Subtype:     runtime.ObjectSubtype.Array,
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

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Animation().ResolveAnimation(params)
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

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Animation().ResolveAnimation(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestAnimationSeekAnimations(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &animation.SeekAnimationsParams{
		Animations:  []string{"animation1", "animation2"},
		CurrentTime: 1,
	}

	mockResult := MockData{
		0,
		&Error{},
		nil,
		"",
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Animation().SeekAnimations(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Animation().SeekAnimations(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestAnimationSetPaused(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &animation.SetPausedParams{
		Animations: []string{"animation1", "animation2"},
		Paused:     true,
	}

	mockResult := MockData{
		0,
		&Error{},
		nil,
		"",
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Animation().SetPaused(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Animation().SetPaused(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestAnimationSetPlaybackRate(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &animation.SetPlaybackRateParams{
		PlaybackRate: 1,
	}

	mockResult := MockData{
		0,
		&Error{},
		nil,
		"",
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Animation().SetPlaybackRate(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Animation().SetPlaybackRate(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestAnimationSetTiming(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	params := &animation.SetTimingParams{
		AnimationID: "animation-id",
		Duration:    1,
		Delay:       1,
	}

	mockResult := MockData{
		0,
		&Error{},
		nil,
		"",
	}

	chrome.AddData(MockData{0, &Error{}, mockResult, ""})
	result := <-soc.Animation().SetTiming(params)
	if nil != result.Err {
		t.Errorf("Expected nil, got error: '%s'", result.Err.Error())
	}

	chrome.AddData(MockData{0, genericError, nil, ""})
	result = <-soc.Animation().SetTiming(params)
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestAnimationOnAnimationCanceled(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *animation.CanceledEvent)
	soc.Animation().OnAnimationCanceled(func(eventData *animation.CanceledEvent) {
		resultChan <- eventData
	})

	mockResult := &animation.CanceledEvent{
		ID: "event-id",
	}
	chrome.AddData(MockData{0,
		&Error{},
		mockResult,
		"Animation.animationCanceled",
	})
	result := <-resultChan
	if result.ID != mockResult.ID {
		t.Errorf(
			"Expected frame ID '%v', got '%v'",
			mockResult.ID,
			result.ID,
		)
	}

	chrome.AddData(MockData{0,
		genericError,
		nil,
		"Animation.animationCanceled",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestAnimationOnAnimationCreated(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *animation.CreatedEvent)
	soc.Animation().OnAnimationCreated(func(eventData *animation.CreatedEvent) {
		resultChan <- eventData
	})

	mockResult := &animation.CreatedEvent{
		ID: "event-id",
	}
	chrome.AddData(MockData{0,
		&Error{},
		mockResult,
		"Animation.animationCreated",
	})
	result := <-resultChan
	if result.ID != mockResult.ID {
		t.Errorf(
			"Expected frame ID %v, got %v",
			mockResult.ID,
			result.ID,
		)
	}

	chrome.AddData(MockData{0,
		genericError,
		nil,
		"Animation.animationCreated",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}

func TestAnimationOnAnimationStarted(t *testing.T) {
	chrome := NewMockChrome()
	chrome.ListenAndServe()
	chrome.IgnoreInput = true
	defer chrome.Close()
	soc := New(chrome.URL)
	defer soc.Stop()

	resultChan := make(chan *animation.StartedEvent)
	soc.Animation().OnAnimationStarted(func(eventData *animation.StartedEvent) {
		resultChan <- eventData
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
			Type:         animation.Type.CSSTransition,
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
	chrome.AddData(MockData{0,
		&Error{},
		mockResult,
		"Animation.animationStarted",
	})
	result := <-resultChan
	if result.Animation.ID != mockResult.Animation.ID {
		t.Errorf(
			"Expected %s, got %s",
			mockResult.Animation.ID,
			result.Animation.ID,
		)
	}

	chrome.AddData(MockData{0,
		genericError,
		nil,
		"Animation.animationStarted",
	})
	result = <-resultChan
	if nil == result.Err {
		t.Errorf("Expected error, got success")
	}
}
