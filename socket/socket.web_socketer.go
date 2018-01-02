package socket

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

/*
ChromeWebSocket represents a WebSocketer interface
*/
type ChromeWebSocket struct {
	*websocket.Conn
}

/*
NewWebsocket returns a connected socket connection
*/
func NewWebsocket(socketURL *url.URL) (WebSocketer, error) {
	dialer := &websocket.Dialer{EnableCompression: false}
	header := http.Header{"Origin": []string{socketURL.String()}}

	websocket, response, err := dialer.Dial(socketURL.String(), header)
	if err != nil {
		return nil, fmt.Errorf(
			"Could not create websocket connection. %s responded with '%s'",
			socketURL.String(),
			err.Error(),
		)
	}
	log.Infof("Websocket connection to %s established: %s", socketURL.String(), response.Status)

	return &ChromeWebSocket{websocket}, nil
}

/*
AddMockData implements WebSocketer
*/
func (sock *ChromeWebSocket) AddMockData(id int, error *Error, method string, data ...interface{}) {
	log.Errorf("Attempted to add mock data outside of unit tests")
}
