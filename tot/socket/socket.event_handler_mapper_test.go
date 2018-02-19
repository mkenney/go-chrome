package socket

import (
	"testing"
)

func TestEventHandlerMapper(t *testing.T) {
	var err error

	handlerMap := NewEventHandlerMap()
	handler := NewEventHandler(
		"eventName",
		func(response *Response) {},
	)

	err = handlerMap.Add(handler)
	if nil != err {
		t.Errorf("Expected nil, got error: '%s'", err.Error())
	}

	err = handlerMap.Add(handler)
	if nil == err {
		t.Errorf("Expected error, got nil")
	}

	err = handlerMap.Remove(handler)
	if nil != err {
		t.Errorf("Expected nil, got error: '%s'", err.Error())
	}

	err = handlerMap.Remove(handler)
	if nil == err {
		t.Errorf("Expected error, got nil")
	}

	// no-op
	handlerMap.Delete("eventName")
}
