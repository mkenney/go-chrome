package socket

/*
WebSocketer defines the interface for web socket connections
*/
type WebSocketer interface {
	Close() error
	ReadJSON(v interface{}) error
	WriteJSON(v interface{}) error
}
