package socket

/*
WebSocketer defines the minimum interface required API for web socket
connections to the Chromium browser.
*/
type WebSocketer interface {
	Close() error
	ReadJSON(v interface{}) error
	WriteJSON(v interface{}) error
	AddMockData(id int, error *Error, method string, data ...interface{})
}
