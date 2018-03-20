package socket

import (
	"net/url"

	log "github.com/sirupsen/logrus"
)

/*
NewMockWebsocket returns a WebSocketer mock for tests that have socket
connection dependencies.
*/
func NewMockWebsocket(socketURL *url.URL) (*ChromeWebSocket, error) {
	log.Infof("Mock websocket connection to %s established", socketURL.String())
	return &ChromeWebSocket{
		mockResponses: make([]*Response, 0),
	}, nil
}
