package socket

/*
Conner defines the Socket connection interface required by the API.
*/
type Conner interface {
	// Connect establishes a websocket connection
	Connect() error

	// Connect closes a websocket connection
	Disconnect() error

	// ReadJSON reads data from a websocket connection
	ReadJSON(v interface{}) error

	// WriteJSON writes data to a websocket connection
	WriteJSON(v interface{}) error
}
