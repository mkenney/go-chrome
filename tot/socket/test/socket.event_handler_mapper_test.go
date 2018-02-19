package test

import (
	"testing"

	"github.com/mkenney/go-chrome/tot/socket"
)

func TestEventHandlerMapper(t *testing.T) {
	var err error

	handlerMap := socket.NewEventHandlerMap()
	handler := socket.NewEventHandler(
		"eventName",
		func(response *socket.Response) {},
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
