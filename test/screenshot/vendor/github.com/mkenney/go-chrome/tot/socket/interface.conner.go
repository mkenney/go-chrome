package socket

/*
Conner defines the Socket connection interface for managing websocket
connections, including reading and writing.
*/
type Conner interface {
	// Conn returns the current web socket pointer.
	Conn() WebSocketer

	// Connect establishes a websocket connection.
	Connect() error

	// Connected returns whether a connection exists.
	Connected() bool

	// Disconnect closes a websocket connection.
	Disconnect() error

	// ReadJSON reads data from a websocket connection.
	ReadJSON(v interface{}) error

	// WriteJSON writes data to a websocket connection.
	WriteJSON(v interface{}) error
}
