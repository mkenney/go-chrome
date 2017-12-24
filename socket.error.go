package Chrome

/*
SocketError represents an error
*/
type SocketError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
