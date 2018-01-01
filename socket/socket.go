/*
Package socket allows for tools to instrument, inspect, debug and profile
Chromium, Chrome and other Blink-based browsers. Many existing projects
currently use the protocol. The Chrome DevTools uses this protocol and the team
maintains its API.

See https://chromedevtools.github.io/devtools-protocol/ for more information.
*/
package socket

import (
	"encoding/json"
)

//////////////////////////////////////////////////
// Socket request and response data structs
//////////////////////////////////////////////////

/*
Error represents a socket response error.
*/
type Error struct {
	Code    int             `json:"code"`
	Data    json.RawMessage `json:"data"`
	Message string          `json:"message"`
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
Payload represents a WebSocket JSON payload for a command.
*/
type Payload struct {
	ID     int         `json:"id"`
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}
