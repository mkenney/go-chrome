package tethering

/*
AcceptedEvent represents Overlay.accepted event data.

https://chromedevtools.github.io/devtools-protocol/tot/Tethering/#event-accepted
*/
type AcceptedEvent struct {
	// Port number that was successfully bound.
	Port int `json:"port"`

	// Connection ID to be used.
	ConnectionID string `json:"connectionId"`
}
