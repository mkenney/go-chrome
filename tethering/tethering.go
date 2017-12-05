package Tethering

/*
BindParams represents Tethering.bind parameters.
*/
type BindParams struct {
	// Port number to bind.
	Port int `json:"port"`
}

/*
UnbindParams represents Tethering.unbind parameters.
*/
type UnbindParams struct {
	// Port number to unbind.
	Port int `json:"port"`
}

/*
AcceptedEvent represents Overlay.accepted event data.
*/
type AcceptedEvent struct {
	// Port number that was successfully bound.
	Port int `json:"port"`

	// Connection ID to be used.
	ConnectionID string `json:"connectionId"`
}
