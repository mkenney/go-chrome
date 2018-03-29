/*
Package socket allows for tools to instrument, inspect, debug and profile
Chromium, Chrome and other Blink-based browsers. Many existing projects
currently use the protocol. The Chrome DevTools team maintains the protocol API.

See https://chromedevtools.github.io/devtools-protocol/ for details.
*/
package socket

import (
	"encoding/json"
	"fmt"
)

/*
Error represents a socket response error.
*/
type Error struct {
	Code    int             `json:"code"`
	Data    json.RawMessage `json:"data"`
	Message string          `json:"message"`
}

/*
Error implements the error interface for socket response Error structs
*/
func (err Error) Error() string {
	return fmt.Sprintf("code=%d, data=%s, msg=%s", err.Code, err.Data, err.Message)
}

/*
Response represents a socket message.
*/
type Response struct {
	Error  *Error          `json:"error"`
	ID     int             `json:"id"`
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
	Result json.RawMessage `json:"result"`
}

/*
Payload represents a WebSocket JSON payload for a sending a command to the
websocket.
*/
type Payload struct {
	ID     int         `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}
