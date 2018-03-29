package socket

/*
WebSocketer defines the minimum interface required API for web socket
connections to the Chromium browser.
*/
type WebSocketer interface {
	// Close closes the current websocket connection.
	Close() error

	// ReadJSON listens for the next websocket message and unmarshalls it into
	// the provided variable.
	ReadJSON(v interface{}) error

	// WriteJSON marshalls the provided data as JSON and writes it to the
	// websocket.
	WriteJSON(v interface{}) error
}
